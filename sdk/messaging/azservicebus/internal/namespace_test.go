// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package internal

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/Azure/azure-amqp-common-go/v3/auth"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/test"
	"github.com/Azure/go-amqp"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type (
	serviceBusSuite struct {
		test.BaseSuite
	}
)

func TestSB(t *testing.T) {
	if os.Getenv("SERVICEBUS_CONNECTION_STRING") == "" {
		t.Skipf("environment variable SERVICEBUS_CONNECTION_STRING was not set")
	}

	suite.Run(t, new(serviceBusSuite))
}

func (suite *serviceBusSuite) TestCreateNamespaceFromConnectionString() {
	connStr := os.Getenv("SERVICEBUS_CONNECTION_STRING") // `Endpoint=sb://XXXX.servicebus.windows.net/;SharedAccessKeyName=XXXX;SharedAccessKey=XXXX`
	if connStr == "" {
		suite.T().Skipf("environment variable SERVICEBUS_CONNECTION_STRING was not set")
	}

	ns, err := NewNamespace(NamespaceWithConnectionString(connStr))
	if suite.NoError(err) {
		suite.Contains(connStr, ns.Name)
	}
}

func TestNewNamespaceWithAzureEnvironment(t *testing.T) {
	ns, err := NewNamespace(NamespaceWithAzureEnvironment("namespaceName", "AzureGermanCloud"))
	if err != nil {
		t.Fatalf("unexpected error creating namespace: %s", err)
	}
	if ns.Environment != azure.GermanCloud {
		t.Fatalf("expected namespace environment to be %q but was %q", azure.GermanCloud, ns.Environment)
	}
	if !strings.EqualFold(ns.Suffix, azure.GermanCloud.ServiceBusEndpointSuffix) {
		t.Fatalf("expected suffix to be %q but was %q", azure.GermanCloud.ServiceBusEndpointSuffix, ns.Suffix)
	}
	if ns.Name != "namespaceName" {
		t.Fatalf("expected namespace name to be %q but was %q", "namespaceName", ns.Name)
	}
}

// implements `Retrier` interface.
type fakeRetrier struct {
	tryCalled  int
	copyCalled int
}

func (r *fakeRetrier) Copy() Retrier {
	r.copyCalled++

	// NOTE: purposefully not making a copy so I can keep track of the
	// try/copy counts.
	return r
}

func (r *fakeRetrier) Exhausted() bool {
	return false
}

func (r *fakeRetrier) Try(ctx context.Context) bool {
	r.tryCalled++
	return true
}

func TestNamespaceNegotiateClaim(t *testing.T) {
	retrier := &fakeRetrier{}

	ns := &Namespace{
		baseRetrier: retrier,
	}

	cbsNegotiateClaimCalled := 0

	cbsNegotiateClaim := func(ctx context.Context, audience string, conn *amqp.Client, provider auth.TokenProvider) error {
		cbsNegotiateClaimCalled++
		return nil
	}

	getAMQPClientCalled := 0

	getAMQPClient := func(ctx context.Context) (*amqp.Client, error) {
		getAMQPClientCalled++
		return &amqp.Client{}, nil
	}

	var errorsLogged []error

	// fire off a basic negotiate claim. The renewal duration is so long that it won't run - that's a separate test.
	cancel, err := ns.negotiateClaimImpl(context.Background(), "my entity path", time.Hour*24, cbsNegotiateClaim, getAMQPClient, func(err error) {
		errorsLogged = append(errorsLogged, err)
	})
	defer cancel()

	require.NoError(t, err)
	cancel()

	require.EqualValues(t, getAMQPClientCalled, 1)
	require.EqualValues(t, 1, cbsNegotiateClaimCalled)
	require.EqualValues(t, 1, retrier.copyCalled)
	require.EqualValues(t, 1, retrier.tryCalled)
	require.Empty(t, errorsLogged)
}

func TestNamespaceNegotiateClaimRenewal(t *testing.T) {
	retrier := &fakeRetrier{}

	ns := &Namespace{
		baseRetrier: retrier,
	}

	cbsNegotiateClaimCalled := 0

	cbsNegotiateClaim := func(ctx context.Context, audience string, conn *amqp.Client, provider auth.TokenProvider) error {
		cbsNegotiateClaimCalled++
		return nil
	}

	getAMQPClientCalled := 0

	notify := make(chan struct{})

	getAMQPClient := func(ctx context.Context) (*amqp.Client, error) {
		getAMQPClientCalled++

		if getAMQPClientCalled == 3 {
			close(notify)
			<-ctx.Done()
		}

		return &amqp.Client{}, nil
	}

	var errorsLogged []error

	// fire off a basic negotiate claim. The renewal duration is so long that it won't run - that's a separate test.
	cancel, err := ns.negotiateClaimImpl(context.Background(), "my entity path", time.Millisecond, cbsNegotiateClaim, getAMQPClient, func(err error) {
		errorsLogged = append(errorsLogged, err)
	})
	defer cancel()

	require.NoError(t, err)
	time.Sleep(time.Second * 3) // make sure, even with variability, we get at least one renewal

	<-notify

	require.GreaterOrEqual(t, getAMQPClientCalled, 2+1) // that last +1 is when we blocked to prevent us renewing too much for our test!
	require.EqualValues(t, 2, cbsNegotiateClaimCalled)
	require.EqualValues(t, 2, retrier.copyCalled)
	require.EqualValues(t, 2, retrier.tryCalled)
	require.Empty(t, errorsLogged)

	cancel()
}

// TearDownSuite destroys created resources during the run of the suite
func (suite *serviceBusSuite) TearDownSuite() {
	suite.BaseSuite.TearDownSuite()
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	suite.deleteAllTaggedQueues(ctx)
	suite.deleteAllTaggedTopics(ctx)
}

func (suite *serviceBusSuite) deleteAllTaggedQueues(ctx context.Context) {
	ns := suite.getNewSasInstance()
	qm := NewQueueManager(ns.GetHTTPSHostURI(), ns.TokenProvider)

	qs, err := qm.List(ctx)
	if err != nil {
		suite.T().Fatal(err)
	}

	for _, q := range qs {
		if strings.HasSuffix(q.Name, suite.TagID) {
			err := qm.Delete(ctx, q.Name)
			if err != nil {
				suite.T().Fatal(err)
			}
		}
	}
}

func (suite *serviceBusSuite) deleteAllTaggedTopics(ctx context.Context) {
	ns := suite.getNewSasInstance()
	tm := ns.NewTopicManager()

	topics, err := tm.List(ctx)
	if err != nil {
		suite.T().Fatal(err)
	}

	for _, topic := range topics {
		if strings.HasSuffix(topic.Name, suite.TagID) {
			err := tm.Delete(ctx, topic.Name)
			if err != nil {
				suite.T().Fatal(err)
			}
		}
	}
}

func (suite *serviceBusSuite) getNewSasInstance(opts ...NamespaceOption) *Namespace {
	ns, err := NewNamespace(append(opts, NamespaceWithConnectionString(suite.ConnStr))...)
	if err != nil {
		suite.T().Fatal(err)
	}
	return ns
}
