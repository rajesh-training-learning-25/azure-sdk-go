package azservicebus

type ReceiveMode int

const (
	ReceiveModePeekLock         ReceiveMode = 0
	ReceiveModeReceiveAndDelete ReceiveMode = 1
)

type SubQueue string

const (
	SubQueueDeadLetter = "deadLetter"
	SubQueueTransfer   = "transferDeadLetter"
)

type ServiceBusReceiver struct{}

// used for batch APIs

// TODO: needs manual credit management.
// func (r *ServiceBusReceiver) ReceiveMessages(ctx context.Context, numMessages int) ([]*ServiceBusReceivedMessage, error) {
// 	return nil, nil
// }
