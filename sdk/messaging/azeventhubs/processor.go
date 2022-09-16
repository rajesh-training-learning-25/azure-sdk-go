// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azeventhubs

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	azlog "github.com/Azure/azure-sdk-for-go/sdk/internal/log"
)

// ProcessorStrategy specifies the load balancing strategy used by the Processor.
type ProcessorStrategy string

const (
	// ProcessorStrategyBalanced will attempt to claim a single partition at a time, until each active
	// owner has an equal share of partitions.
	// This is the default strategy.
	ProcessorStrategyBalanced ProcessorStrategy = "balanced"

	// ProcessorStrategyGreedy will attempt to claim as many partitions at a time as it can, ignoring
	// balance.
	ProcessorStrategyGreedy ProcessorStrategy = "greedy"
)

// ProcessorOptions are the options for the NewProcessor
// function.
type ProcessorOptions struct {
	// LoadBalancingStrategy dictates how concurrent Processor instances distribute
	// ownership of partitions between them.
	// The default strategy is ProcessorStrategyBalanced.
	LoadBalancingStrategy ProcessorStrategy

	// UpdateInterval controls how often attempt to claim partitions.
	// The default value is 10 seconds.
	UpdateInterval time.Duration

	// PartitionExpirationDuration is the amount of time before a partition is considered
	// unowned.
	// The default value is 60 seconds.
	PartitionExpirationDuration time.Duration

	// StartPositions are the default start positions (configurable per partition, or with an overall
	// default value) if a checkpoint is not found in the CheckpointStore.
	// The default position is Latest.
	StartPositions StartPositions

	// OwnerLevel is the priority for partition clients created by this Processor, also known as
	// the 'epoch' level.
	// When used, a partition client with a higher OwnerLevel will take ownership of a partition
	// from partition clients with a lower OwnerLevel.
	// Default is 0.
	OwnerLevel int64
}

// StartPositions are used if there is no checkpoint for a partition in
// the checkpoint store.
type StartPositions struct {
	// PerPartition controls the start position for a specific partition,
	// by partition ID. If a partition is not configured here it will default
	// to Default start position.
	PerPartition map[string]StartPosition

	// Default is used if the partition is not found in the PerPartition map.
	Default StartPosition
}

// Processor uses a ConsumerClient and CheckpointStore to provide automatic
// load balancing between multiple Processor instances, even in separate
// processes or on separate machines.
//
// See [example_processor_test.go] for an example of typical usage or Run
// for a more detailed description of how load balancing works.
//
// [example_processor_test.go]: https://github.com/Azure/azure-sdk-for-go/blob/main/sdk/messaging/azeventhubs/example_processor_test.go
type Processor struct {
	ownershipUpdateInterval time.Duration
	defaultStartPositions   StartPositions
	checkpointStore         CheckpointStore
	ownerLevel              int64

	// consumerClient is actually a *azeventhubs.ConsumerClient
	// it's an interface here to make testing easier.
	consumerClient consumerClientForProcessor

	nextClients           chan *ProcessorPartitionClient
	consumerClientDetails consumerClientDetails

	runCalled chan struct{}
	lb        *processorLoadBalancer
}

type consumerClientForProcessor interface {
	GetEventHubProperties(ctx context.Context, options *GetEventHubPropertiesOptions) (EventHubProperties, error)
	NewPartitionClient(partitionID string, options *PartitionClientOptions) (*PartitionClient, error)
	getDetails() consumerClientDetails
}

// NewProcessor creates a Processor.
//
// More general information can be found on the documentation for the [azeventhubs.Processor] type or the [example_processor_test.go] for an example.
//
// [example_processor_test.go]: https://github.com/Azure/azure-sdk-for-go/blob/main/sdk/messaging/azeventhubs/example_processor_test.go
func NewProcessor(consumerClient *ConsumerClient, checkpointStore CheckpointStore, options *ProcessorOptions) (*Processor, error) {
	return newProcessorImpl(consumerClient, checkpointStore, options)
}

func newProcessorImpl(consumerClient consumerClientForProcessor, checkpointStore CheckpointStore, options *ProcessorOptions) (*Processor, error) {
	if options == nil {
		options = &ProcessorOptions{}
	}

	updateInterval := 10 * time.Second

	if options.UpdateInterval != 0 {
		updateInterval = options.UpdateInterval
	}

	partitionDurationExpiration := time.Minute

	if options.PartitionExpirationDuration != 0 {
		partitionDurationExpiration = options.PartitionExpirationDuration
	}

	startPosPerPartition := map[string]StartPosition{}

	if options.StartPositions.PerPartition != nil {
		for k, v := range options.StartPositions.PerPartition {
			startPosPerPartition[k] = v
		}
	}

	strategy := options.LoadBalancingStrategy

	switch strategy {
	case ProcessorStrategyBalanced:
	case ProcessorStrategyGreedy:
	case "":
		strategy = ProcessorStrategyBalanced
	default:
		return nil, fmt.Errorf("invalid load balancing strategy '%s'", strategy)
	}

	return &Processor{
		ownerLevel:              options.OwnerLevel,
		ownershipUpdateInterval: updateInterval,
		consumerClient:          consumerClient,
		checkpointStore:         checkpointStore,

		defaultStartPositions: StartPositions{
			PerPartition: startPosPerPartition,
			Default:      options.StartPositions.Default,
		},
		consumerClientDetails: consumerClient.getDetails(),
		runCalled:             make(chan struct{}),
		lb:                    newProcessorLoadBalancer(checkpointStore, consumerClient.getDetails(), strategy, partitionDurationExpiration),
		// `nextClients` will be initialized when the user calls Run() since it needs to query the #
		// of partitions on the Event Hub.
	}, nil
}

// NextPartitionClient will get the next owned [PartitionProcessorClient] if one is acquired
// or will block until a new one arrives or [processor.Run] is cancelled.
//
// This function is safe to call before [processor.Run] has been called and will typically
// be executed in a goroutine in a loop.
//
// See [example_processor_test.go] for an example of typical usage.
//
// [example_processor_test.go]: https://github.com/Azure/azure-sdk-for-go/blob/main/sdk/messaging/azeventhubs/example_processor_test.go
func (p *Processor) NextPartitionClient(ctx context.Context) *ProcessorPartitionClient {
	<-p.runCalled

	select {
	case nextClient := <-p.nextClients:
		return nextClient
	case <-ctx.Done():
		return nil
	}
}

// Run handles the load balancing loop, blocking until it encounters fatal errors or the
// passed in context is cancelled.
//
// Run continually checks the passed in CheckpointStore, attempting to distribute partitions
// fairly between itself and any other active consumers. New [ProcessorPartitionClient] instances
// are created and will be returned from [Processor.NextPartitionClient] as they are available.
//
// When partition ownership is lost [ProcessorPartitionClient]'s will return an errors.
//
// Callers will typically create a goroutine, continually calling
// [Processor.NextPartitionClient] in a loop to get [ProcessorPartitionClient]
// instances, and then process those in parallel.
//
// The main thread will call [Processor.Run], blocking until the context passed
// to Run is cancelled.
//

// Run will exit when the passed in context has been cancelled or if there are unrecoverable
// or persistent errors in load balancing.
//
// On cancellation, it will return a nil error.
//
// See [example_processor_test.go] for an example of typical usage.
//
// [example_processor_test.go]: https://github.com/Azure/azure-sdk-for-go/blob/main/sdk/messaging/azeventhubs/example_processor_test.go
func (p *Processor) Run(ctx context.Context) error {
	err := p.runImpl(ctx)

	// the context is the proper way to close down the Run() loop, so it's not
	// an error and doesn't need to be returned.
	if ctx.Err() != nil {
		return nil
	}

	return err
}

func (p *Processor) runImpl(ctx context.Context) error {
	consumers := &sync.Map{}
	defer closeConsumers(ctx, consumers)

	// size the channel to the # of partitions. We can never exceed this size since
	// we'll never reclaim a partition that we already have ownership of.
	eventHubProperties, err := p.initNextClientsCh(ctx)

	if err != nil {
		return err
	}

	// do one dispatch immediately
	if err := p.dispatch(ctx, eventHubProperties, consumers); err != nil {
		return err
	}

	// note randSource is not thread-safe but it's not currently used in a way that requires
	// it to be.
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(calculateUpdateInterval(rnd, p.ownershipUpdateInterval)):
			if err := p.dispatch(ctx, eventHubProperties, consumers); err != nil {
				return err
			}
		}
	}
}

func calculateUpdateInterval(rnd *rand.Rand, updateInterval time.Duration) time.Duration {
	// Introduce some jitter:  [0.0, 1.0) / 2 = [0.0, 0.5) + 0.8 = [0.8, 1.3)
	// (copied from the retry code for calculating jitter)
	return time.Duration(updateInterval.Seconds() * (rnd.Float64()/2 + 0.8) * float64(time.Second))
}

func (p *Processor) initNextClientsCh(ctx context.Context) (EventHubProperties, error) {
	eventHubProperties, err := p.consumerClient.GetEventHubProperties(ctx, nil)

	if err != nil {
		return EventHubProperties{}, err
	}

	p.nextClients = make(chan *ProcessorPartitionClient, len(eventHubProperties.PartitionIDs))
	close(p.runCalled)

	return eventHubProperties, nil
}

// dispatch uses the checkpoint store to figure out which partitions should be processed by this
// instance and starts a PartitionClient, if there isn't one.
// NOTE: due to random number usage in the load balancer, this function is not thread safe.
func (p *Processor) dispatch(ctx context.Context, eventHubProperties EventHubProperties, consumers *sync.Map) error {
	ownerships, err := p.lb.LoadBalance(ctx, eventHubProperties.PartitionIDs)

	if err != nil {
		return err
	}

	checkpoints, err := p.getCheckpointsMap(ctx)

	if err != nil {
		return err
	}

	wg := sync.WaitGroup{}

	for _, ownership := range ownerships {
		wg.Add(1)

		go func(o Ownership) {
			defer wg.Done()

			err := p.addPartitionClient(ctx, o, checkpoints, consumers)

			if err != nil {
				azlog.Writef(EventConsumer, "failed to create partition client for partition '%s': %s", o.PartitionID, err.Error())
			}
		}(ownership)
	}

	wg.Wait()

	return nil
}

// addPartitionClient creates a ProcessorPartitionClient
func (p *Processor) addPartitionClient(ctx context.Context, ownership Ownership, checkpoints map[string]Checkpoint, consumers *sync.Map) error {
	processorPartClient := &ProcessorPartitionClient{
		consumerClientDetails: p.consumerClientDetails,
		checkpointStore:       p.checkpointStore,
		innerClient:           nil,
		partitionID:           ownership.PartitionID,
		cleanupFn: func() {
			consumers.Delete(ownership.PartitionID)
		},
	}

	if _, alreadyExists := consumers.LoadOrStore(ownership.PartitionID, processorPartClient); alreadyExists {
		return nil
	}

	sp, err := p.getStartPosition(checkpoints, ownership)

	if err != nil {
		return err
	}

	partClient, err := p.consumerClient.NewPartitionClient(ownership.PartitionID, &PartitionClientOptions{
		StartPosition: sp,
		OwnerLevel:    &p.ownerLevel,
	})

	if err != nil {
		consumers.Delete(ownership.PartitionID)
		return err
	}

	// make sure we create the link _now_ - if we're stealing we want to stake a claim _now_, rather than
	// later when the user actually calls ReceiveEvents(), since the acquisition of the link is lazy.
	if err := partClient.init(ctx); err != nil {
		consumers.Delete(ownership.PartitionID)
		_ = partClient.Close(ctx)
		return err
	}

	processorPartClient.innerClient = partClient

	select {
	case p.nextClients <- processorPartClient:
		return nil
	default:
		processorPartClient.Close(ctx)
		return fmt.Errorf("partitions channel full, consumer for partition %s could not be returned", ownership.PartitionID)
	}
}

func (p *Processor) getStartPosition(checkpoints map[string]Checkpoint, ownership Ownership) (StartPosition, error) {
	startPosition := p.defaultStartPositions.Default
	cp, hasCheckpoint := checkpoints[ownership.PartitionID]

	if hasCheckpoint {
		if cp.Offset != nil {
			startPosition = StartPosition{
				Offset: cp.Offset,
			}
		} else if cp.SequenceNumber != nil {
			startPosition = StartPosition{
				SequenceNumber: cp.SequenceNumber,
			}
		} else {
			return StartPosition{}, fmt.Errorf("invalid checkpoint for %s, no offset or sequence number", ownership.PartitionID)
		}
	} else if p.defaultStartPositions.PerPartition != nil {
		defaultStartPosition, exists := p.defaultStartPositions.PerPartition[ownership.PartitionID]

		if exists {
			startPosition = defaultStartPosition
		}
	}

	return startPosition, nil
}

func (p *Processor) getCheckpointsMap(ctx context.Context) (map[string]Checkpoint, error) {
	details := p.consumerClient.getDetails()
	checkpoints, err := p.checkpointStore.ListCheckpoints(ctx, details.FullyQualifiedNamespace, details.EventHubName, details.ConsumerGroup, nil)

	if err != nil {
		return nil, err
	}

	m := map[string]Checkpoint{}

	for _, cp := range checkpoints {
		m[cp.PartitionID] = cp
	}

	return m, nil
}

func closeConsumers(ctx context.Context, consumersMap *sync.Map) {
	consumersMap.Range(func(key, value any) bool {
		client := value.(*ProcessorPartitionClient)

		if client != nil {
			client.Close(ctx)
		}

		return true
	})
}
