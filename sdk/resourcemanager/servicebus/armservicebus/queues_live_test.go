//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicebus_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/internal/testutil"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus/v2"
	"github.com/stretchr/testify/suite"
)

type QueuesTestSuite struct {
	suite.Suite

	ctx                   context.Context
	cred                  azcore.TokenCredential
	options               *arm.ClientOptions
	authorizationRuleName string
	namespaceName         string
	queueName             string
	location              string
	resourceGroupName     string
	subscriptionId        string
}

func (testsuite *QueuesTestSuite) SetupSuite() {
	testutil.StartRecording(testsuite.T(), "sdk/resourcemanager/servicebus/armservicebus/testdata")
	testsuite.ctx = context.Background()
	testsuite.cred, testsuite.options = testutil.GetCredAndClientOptions(testsuite.T())
	testsuite.authorizationRuleName = testutil.GenerateAlphaNumericID(testsuite.T(), "queueauthrule", 6)
	testsuite.namespaceName = testutil.GenerateAlphaNumericID(testsuite.T(), "sbnamespaceq", 6)
	testsuite.queueName = testutil.GenerateAlphaNumericID(testsuite.T(), "queuena", 6)
	testsuite.location = testutil.GetEnv("LOCATION", "westus")
	testsuite.resourceGroupName = testutil.GetEnv("RESOURCE_GROUP_NAME", "scenarioTestTempGroup")
	testsuite.subscriptionId = testutil.GetEnv("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")

	resourceGroup, _, err := testutil.CreateResourceGroup(testsuite.ctx, testsuite.subscriptionId, testsuite.cred, testsuite.options, testsuite.location)
	testsuite.Require().NoError(err)
	testsuite.resourceGroupName = *resourceGroup.Name
	testsuite.Prepare()
}

func (testsuite *QueuesTestSuite) TearDownSuite() {
	testsuite.Cleanup()
	_, err := testutil.DeleteResourceGroup(testsuite.ctx, testsuite.subscriptionId, testsuite.cred, testsuite.options, testsuite.resourceGroupName)
	testsuite.Require().NoError(err)
	testutil.StopRecording(testsuite.T())
}

func TestQueuesTestSuite(t *testing.T) {
	suite.Run(t, new(QueuesTestSuite))
}

func (testsuite *QueuesTestSuite) Prepare() {
	var err error
	// From step Namespace_Create
	fmt.Println("Call operation: Namespace_Create")
	namespacesClient, err := armservicebus.NewNamespacesClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	namespacesClientCreateOrUpdateResponsePoller, err := namespacesClient.BeginCreateOrUpdate(testsuite.ctx, testsuite.resourceGroupName, testsuite.namespaceName, armservicebus.SBNamespace{
		Location: to.Ptr(testsuite.location),
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
		SKU: &armservicebus.SBSKU{
			Name: to.Ptr(armservicebus.SKUNameStandard),
			Tier: to.Ptr(armservicebus.SKUTierStandard),
		},
	}, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, namespacesClientCreateOrUpdateResponsePoller)
	testsuite.Require().NoError(err)
}

// Microsoft.ServiceBus/namespaces/queues
func (testsuite *QueuesTestSuite) TestQueue() {
	var err error
	// From step Queue_Create
	fmt.Println("Call operation: Queue_Create")
	queuesClient, err := armservicebus.NewQueuesClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	_, err = queuesClient.CreateOrUpdate(testsuite.ctx, testsuite.resourceGroupName, testsuite.namespaceName, testsuite.queueName, armservicebus.SBQueue{
		Properties: &armservicebus.SBQueueProperties{
			EnablePartitioning: to.Ptr(true),
		},
	}, nil)
	testsuite.Require().NoError(err)

	// From step Queue_Get
	fmt.Println("Call operation: Queue_Get")
	_, err = queuesClient.Get(testsuite.ctx, testsuite.resourceGroupName, testsuite.namespaceName, testsuite.queueName, nil)
	testsuite.Require().NoError(err)

	// From step Queue_ListByNamespace
	fmt.Println("Call operation: Queue_ListByNamespace")
	queuesClientNewListByNamespacePager := queuesClient.NewListByNamespacePager(testsuite.resourceGroupName, testsuite.namespaceName, &armservicebus.QueuesClientListByNamespaceOptions{Skip: nil,
		Top: nil,
	})
	for queuesClientNewListByNamespacePager.More() {
		_, err := queuesClientNewListByNamespacePager.NextPage(testsuite.ctx)
		testsuite.Require().NoError(err)
		break
	}

	// From step Queue_CreateAuthorizationRule
	fmt.Println("Call operation: Queue_CreateAuthorizationRule")
	_, err = queuesClient.CreateOrUpdateAuthorizationRule(testsuite.ctx, testsuite.resourceGroupName, testsuite.namespaceName, testsuite.queueName, testsuite.authorizationRuleName, armservicebus.SBAuthorizationRule{
		Properties: &armservicebus.SBAuthorizationRuleProperties{
			Rights: []*armservicebus.AccessRights{
				to.Ptr(armservicebus.AccessRightsListen),
				to.Ptr(armservicebus.AccessRightsSend)},
		},
	}, nil)
	testsuite.Require().NoError(err)

	// From step Queue_GetAuthorizationRule
	fmt.Println("Call operation: Queue_GetAuthorizationRule")
	_, err = queuesClient.GetAuthorizationRule(testsuite.ctx, testsuite.resourceGroupName, testsuite.namespaceName, testsuite.queueName, testsuite.authorizationRuleName, nil)
	testsuite.Require().NoError(err)

	// From step Queue_ListAuthorizationRules
	fmt.Println("Call operation: Queue_ListAuthorizationRules")
	queuesClientNewListAuthorizationRulesPager := queuesClient.NewListAuthorizationRulesPager(testsuite.resourceGroupName, testsuite.namespaceName, testsuite.queueName, nil)
	for queuesClientNewListAuthorizationRulesPager.More() {
		_, err := queuesClientNewListAuthorizationRulesPager.NextPage(testsuite.ctx)
		testsuite.Require().NoError(err)
		break
	}

	// From step Queue_RegenerateKeys
	fmt.Println("Call operation: Queue_RegenerateKeys")
	_, err = queuesClient.RegenerateKeys(testsuite.ctx, testsuite.resourceGroupName, testsuite.namespaceName, testsuite.queueName, testsuite.authorizationRuleName, armservicebus.RegenerateAccessKeyParameters{
		KeyType: to.Ptr(armservicebus.KeyTypePrimaryKey),
	}, nil)
	testsuite.Require().NoError(err)

	// From step Queue_ListKeys
	fmt.Println("Call operation: Queue_ListKeys")
	_, err = queuesClient.ListKeys(testsuite.ctx, testsuite.resourceGroupName, testsuite.namespaceName, testsuite.queueName, testsuite.authorizationRuleName, nil)
	testsuite.Require().NoError(err)

	// From step Queue_DeleteAuthorizationRule
	fmt.Println("Call operation: Queue_DeleteAuthorizationRule")
	_, err = queuesClient.DeleteAuthorizationRule(testsuite.ctx, testsuite.resourceGroupName, testsuite.namespaceName, testsuite.queueName, testsuite.authorizationRuleName, nil)
	testsuite.Require().NoError(err)

	// From step Queue_Delete
	fmt.Println("Call operation: Queue_Delete")
	_, err = queuesClient.Delete(testsuite.ctx, testsuite.resourceGroupName, testsuite.namespaceName, testsuite.queueName, nil)
	testsuite.Require().NoError(err)
}

func (testsuite *QueuesTestSuite) Cleanup() {
	var err error
	// From step Namespace_Delete
	fmt.Println("Call operation: Namespace_Delete")
	namespacesClient, err := armservicebus.NewNamespacesClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	namespacesClientDeleteResponsePoller, err := namespacesClient.BeginDelete(testsuite.ctx, testsuite.resourceGroupName, testsuite.namespaceName, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, namespacesClientDeleteResponsePoller)
	testsuite.Require().NoError(err)
}
