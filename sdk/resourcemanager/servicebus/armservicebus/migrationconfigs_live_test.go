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

type MigrationconfigsTestSuite struct {
	suite.Suite

	ctx               context.Context
	cred              azcore.TokenCredential
	options           *arm.ClientOptions
	namespaceName     string
	namespaceName2    string
	postMigrationName string
	recondNamespaceId string
	location          string
	resourceGroupName string
	subscriptionId    string
}

func (testsuite *MigrationconfigsTestSuite) SetupSuite() {
	testutil.StartRecording(testsuite.T(), "sdk/resourcemanager/servicebus/armservicebus/testdata")
	testsuite.ctx = context.Background()
	testsuite.cred, testsuite.options = testutil.GetCredAndClientOptions(testsuite.T())
	testsuite.namespaceName = testutil.GenerateAlphaNumericID(testsuite.T(), "sbnamespacemc", 6)
	testsuite.namespaceName2 = testutil.GenerateAlphaNumericID(testsuite.T(), "sbnamespacetwomc", 6)
	testsuite.postMigrationName = testutil.GenerateAlphaNumericID(testsuite.T(), "postmigrationna", 6)
	testsuite.location = testutil.GetEnv("LOCATION", "westus")
	testsuite.resourceGroupName = testutil.GetEnv("RESOURCE_GROUP_NAME", "scenarioTestTempGroup")
	testsuite.subscriptionId = testutil.GetEnv("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")

	resourceGroup, _, err := testutil.CreateResourceGroup(testsuite.ctx, testsuite.subscriptionId, testsuite.cred, testsuite.options, testsuite.location)
	testsuite.Require().NoError(err)
	testsuite.resourceGroupName = *resourceGroup.Name
	testsuite.Prepare()
}

func (testsuite *MigrationconfigsTestSuite) TearDownSuite() {
	_, err := testutil.DeleteResourceGroup(testsuite.ctx, testsuite.subscriptionId, testsuite.cred, testsuite.options, testsuite.resourceGroupName)
	testsuite.Require().NoError(err)
	testutil.StopRecording(testsuite.T())
}

func TestMigrationconfigsTestSuite(t *testing.T) {
	suite.Run(t, new(MigrationconfigsTestSuite))
}

func (testsuite *MigrationconfigsTestSuite) Prepare() {
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

	// From step Namespace_Create2
	fmt.Println("Call operation: Namespace_Create")
	namespacesClientCreateOrUpdateResponsePoller, err = namespacesClient.BeginCreateOrUpdate(testsuite.ctx, testsuite.resourceGroupName, testsuite.namespaceName2, armservicebus.SBNamespace{
		Location: to.Ptr("westus2"),
		SKU: &armservicebus.SBSKU{
			Name: to.Ptr(armservicebus.SKUNamePremium),
			Tier: to.Ptr(armservicebus.SKUTierPremium),
		},
	}, nil)
	testsuite.Require().NoError(err)
	var namespacesClientCreateOrUpdateResponse *armservicebus.NamespacesClientCreateOrUpdateResponse
	namespacesClientCreateOrUpdateResponse, err = testutil.PollForTest(testsuite.ctx, namespacesClientCreateOrUpdateResponsePoller)
	testsuite.Require().NoError(err)
	testsuite.recondNamespaceId = *namespacesClientCreateOrUpdateResponse.ID
}

// Microsoft.ServiceBus/namespaces/migrationConfigurations
func (testsuite *MigrationconfigsTestSuite) TestMigrationconfig() {
	var err error
	// From step MigrationConfig_Create
	fmt.Println("Call operation: MigrationConfig_Create")
	migrationConfigsClient, err := armservicebus.NewMigrationConfigsClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	migrationConfigsClientCreateAndStartMigrationResponsePoller, err := migrationConfigsClient.BeginCreateAndStartMigration(testsuite.ctx, testsuite.resourceGroupName, testsuite.namespaceName, armservicebus.MigrationConfigurationNameDefault, armservicebus.MigrationConfigProperties{
		Properties: &armservicebus.MigrationConfigPropertiesProperties{
			PostMigrationName: to.Ptr(testsuite.postMigrationName),
			TargetNamespace:   to.Ptr(testsuite.recondNamespaceId),
		},
	}, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, migrationConfigsClientCreateAndStartMigrationResponsePoller)
	testsuite.Require().NoError(err)

	// From step MigrationConfig_Get
	fmt.Println("Call operation: MigrationConfig_Get")
	_, err = migrationConfigsClient.Get(testsuite.ctx, testsuite.resourceGroupName, testsuite.namespaceName, armservicebus.MigrationConfigurationNameDefault, nil)
	testsuite.Require().NoError(err)

	// From step MigrationConfig_List
	fmt.Println("Call operation: MigrationConfig_List")
	migrationConfigsClientNewListPager := migrationConfigsClient.NewListPager(testsuite.resourceGroupName, testsuite.namespaceName, nil)
	for migrationConfigsClientNewListPager.More() {
		_, err := migrationConfigsClientNewListPager.NextPage(testsuite.ctx)
		testsuite.Require().NoError(err)
		break
	}

	// From step MigrationConfig_Revert
	fmt.Println("Call operation: MigrationConfig_Revert")
	_, err = migrationConfigsClient.Revert(testsuite.ctx, testsuite.resourceGroupName, testsuite.namespaceName, armservicebus.MigrationConfigurationNameDefault, nil)
	testsuite.Require().NoError(err)

	// From step MigrationConfig_CompleteMigration
	fmt.Println("Call operation: MigrationConfig_CompleteMigration")
	_, err = migrationConfigsClient.CompleteMigration(testsuite.ctx, testsuite.resourceGroupName, testsuite.namespaceName, armservicebus.MigrationConfigurationNameDefault, nil)
	testsuite.Require().NoError(err)
}
