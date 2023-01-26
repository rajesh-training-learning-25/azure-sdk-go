//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos_test

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/internal/testutil"
	"github.com/stretchr/testify/suite"
)

type MongodbresourcesTestSuite struct {
	suite.Suite

	ctx               context.Context
	cred              azcore.TokenCredential
	options           *arm.ClientOptions
	accountName       string
	location          string
	resourceGroupName string
	subscriptionId    string
}

func (testsuite *MongodbresourcesTestSuite) SetupSuite() {
	testsuite.ctx = context.Background()
	testsuite.cred, testsuite.options = testutil.GetCredAndClientOptions(testsuite.T())
	testsuite.accountName = "databaseaccountmongodb"
	testsuite.location = testutil.GetEnv("LOCATION", "westus")
	testsuite.resourceGroupName = testutil.GetEnv("RESOURCE_GROUP_NAME", "scenarioTestTempGroup")
	testsuite.subscriptionId = testutil.GetEnv("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")

	testutil.StartRecording(testsuite.T(), "sdk/resourcemanager/cosmos/armcosmos/testdata")
	resourceGroup, _, err := testutil.CreateResourceGroup(testsuite.ctx, testsuite.subscriptionId, testsuite.cred, testsuite.options, testsuite.location)
	testsuite.Require().NoError(err)
	testsuite.resourceGroupName = *resourceGroup.Name
	testsuite.Prepare()
}

func (testsuite *MongodbresourcesTestSuite) TearDownSuite() {
	_, err := testutil.DeleteResourceGroup(testsuite.ctx, testsuite.subscriptionId, testsuite.cred, testsuite.options, testsuite.resourceGroupName)
	testsuite.Require().NoError(err)
	testutil.StopRecording(testsuite.T())
}

func TestMongodbresourcesTestSuite(t *testing.T) {
	suite.Run(t, new(MongodbresourcesTestSuite))
}

func (testsuite *MongodbresourcesTestSuite) Prepare() {
	var err error
	// From step DatabaseAccount_Create
	databaseAccountsClient, err := armcosmos.NewDatabaseAccountsClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	databaseAccountsClientCreateOrUpdateResponsePoller, err := databaseAccountsClient.BeginCreateOrUpdate(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, armcosmos.DatabaseAccountCreateUpdateParameters{
		Location: to.Ptr(testsuite.location),
		Kind:     to.Ptr(armcosmos.DatabaseAccountKindMongoDB),
		Properties: &armcosmos.DatabaseAccountCreateUpdateProperties{
			CreateMode:               to.Ptr(armcosmos.CreateModeDefault),
			DatabaseAccountOfferType: to.Ptr("Standard"),
			Locations: []*armcosmos.Location{
				{
					FailoverPriority: to.Ptr[int32](0),
					IsZoneRedundant:  to.Ptr(false),
					LocationName:     to.Ptr("eastus"),
				}},
		},
	}, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, databaseAccountsClientCreateOrUpdateResponsePoller)
	testsuite.Require().NoError(err)
}

// Microsoft.DocumentDB/databaseAccounts/mongodbDatabases
func (testsuite *MongodbresourcesTestSuite) TestMongodbresources() {
	collectionName := "mongodb_collection"
	databaseName := "scenario_mongodb"
	var err error
	// From step MongoDBResources_CreateMongoDBDatabase
	mongoDBResourcesClient, err := armcosmos.NewMongoDBResourcesClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	mongoDBResourcesClientCreateUpdateMongoDBDatabaseResponsePoller, err := mongoDBResourcesClient.BeginCreateUpdateMongoDBDatabase(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, databaseName, armcosmos.MongoDBDatabaseCreateUpdateParameters{
		Location: to.Ptr(testsuite.location),
		Tags:     map[string]*string{},
		Properties: &armcosmos.MongoDBDatabaseCreateUpdateProperties{
			Options: &armcosmos.CreateUpdateOptions{
				Throughput: to.Ptr[int32](2000),
			},
			Resource: &armcosmos.MongoDBDatabaseResource{
				ID: to.Ptr(databaseName),
			},
		},
	}, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, mongoDBResourcesClientCreateUpdateMongoDBDatabaseResponsePoller)
	testsuite.Require().NoError(err)

	// From step MongoDBResources_GetMongoDBDatabase
	_, err = mongoDBResourcesClient.GetMongoDBDatabase(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, databaseName, nil)
	testsuite.Require().NoError(err)

	// From step MongoDBResources_ListMongoDBDatabases
	mongoDBResourcesClientNewListMongoDBDatabasesPager := mongoDBResourcesClient.NewListMongoDBDatabasesPager(testsuite.resourceGroupName, testsuite.accountName, nil)
	for mongoDBResourcesClientNewListMongoDBDatabasesPager.More() {
		_, err := mongoDBResourcesClientNewListMongoDBDatabasesPager.NextPage(testsuite.ctx)
		testsuite.Require().NoError(err)
		break
	}

	// From step MongoDBResources_UpdateMongoDBDatabaseThroughput
	mongoDBResourcesClientUpdateMongoDBDatabaseThroughputResponsePoller, err := mongoDBResourcesClient.BeginUpdateMongoDBDatabaseThroughput(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, databaseName, armcosmos.ThroughputSettingsUpdateParameters{
		Location: to.Ptr(testsuite.location),
		Tags:     map[string]*string{},
		Properties: &armcosmos.ThroughputSettingsUpdateProperties{
			Resource: &armcosmos.ThroughputSettingsResource{
				Throughput: to.Ptr[int32](400),
			},
		},
	}, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, mongoDBResourcesClientUpdateMongoDBDatabaseThroughputResponsePoller)
	testsuite.Require().NoError(err)

	// From step MongoDBResources_GetMongoDBDatabaseThroughput
	_, err = mongoDBResourcesClient.GetMongoDBDatabaseThroughput(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, databaseName, nil)
	testsuite.Require().NoError(err)

	// From step MongoDBResources_MigrateMongoDBDatabaseToAutoscale
	mongoDBResourcesClientMigrateMongoDBDatabaseToAutoscaleResponsePoller, err := mongoDBResourcesClient.BeginMigrateMongoDBDatabaseToAutoscale(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, databaseName, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, mongoDBResourcesClientMigrateMongoDBDatabaseToAutoscaleResponsePoller)
	testsuite.Require().NoError(err)

	// From step MongoDBResources_MigrateMongoDBDatabaseToManualThroughput
	mongoDBResourcesClientMigrateMongoDBDatabaseToManualThroughputResponsePoller, err := mongoDBResourcesClient.BeginMigrateMongoDBDatabaseToManualThroughput(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, databaseName, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, mongoDBResourcesClientMigrateMongoDBDatabaseToManualThroughputResponsePoller)
	testsuite.Require().NoError(err)

	// From step MongoDBResources_CreateMongoDBCollection
	mongoDBResourcesClientCreateUpdateMongoDBCollectionResponsePoller, err := mongoDBResourcesClient.BeginCreateUpdateMongoDBCollection(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, databaseName, collectionName, armcosmos.MongoDBCollectionCreateUpdateParameters{
		Location: to.Ptr(testsuite.location),
		Tags:     map[string]*string{},
		Properties: &armcosmos.MongoDBCollectionCreateUpdateProperties{
			Options: &armcosmos.CreateUpdateOptions{
				Throughput: to.Ptr[int32](2000),
			},
			Resource: &armcosmos.MongoDBCollectionResource{
				ID: to.Ptr(collectionName),
				ShardKey: map[string]*string{
					"testKey": to.Ptr("Hash"),
				},
			},
		},
	}, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, mongoDBResourcesClientCreateUpdateMongoDBCollectionResponsePoller)
	testsuite.Require().NoError(err)

	// From step MongoDBResources_GetMongoDBCollection
	_, err = mongoDBResourcesClient.GetMongoDBCollection(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, databaseName, collectionName, nil)
	testsuite.Require().NoError(err)

	// From step MongoDBResources_ListMongoDBCollections
	mongoDBResourcesClientNewListMongoDBCollectionsPager := mongoDBResourcesClient.NewListMongoDBCollectionsPager(testsuite.resourceGroupName, testsuite.accountName, databaseName, nil)
	for mongoDBResourcesClientNewListMongoDBCollectionsPager.More() {
		_, err := mongoDBResourcesClientNewListMongoDBCollectionsPager.NextPage(testsuite.ctx)
		testsuite.Require().NoError(err)
		break
	}

	// From step MongoDBResources_UpdateMongoDBCollectionThroughput
	mongoDBResourcesClientUpdateMongoDBCollectionThroughputResponsePoller, err := mongoDBResourcesClient.BeginUpdateMongoDBCollectionThroughput(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, databaseName, collectionName, armcosmos.ThroughputSettingsUpdateParameters{
		Location: to.Ptr(testsuite.location),
		Tags:     map[string]*string{},
		Properties: &armcosmos.ThroughputSettingsUpdateProperties{
			Resource: &armcosmos.ThroughputSettingsResource{
				Throughput: to.Ptr[int32](400),
			},
		},
	}, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, mongoDBResourcesClientUpdateMongoDBCollectionThroughputResponsePoller)
	testsuite.Require().NoError(err)

	// From step MongoDBResources_GetMongoDBCollectionThroughput
	_, err = mongoDBResourcesClient.GetMongoDBCollectionThroughput(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, databaseName, collectionName, nil)
	testsuite.Require().NoError(err)

	// From step MongoDBResources_MigrateMongoDBCollectionToAutoscale
	mongoDBResourcesClientMigrateMongoDBCollectionToAutoscaleResponsePoller, err := mongoDBResourcesClient.BeginMigrateMongoDBCollectionToAutoscale(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, databaseName, collectionName, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, mongoDBResourcesClientMigrateMongoDBCollectionToAutoscaleResponsePoller)
	testsuite.Require().NoError(err)

	// From step MongoDBResources_MigrateMongoDBCollectionToManualThroughput
	mongoDBResourcesClientMigrateMongoDBCollectionToManualThroughputResponsePoller, err := mongoDBResourcesClient.BeginMigrateMongoDBCollectionToManualThroughput(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, databaseName, collectionName, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, mongoDBResourcesClientMigrateMongoDBCollectionToManualThroughputResponsePoller)
	testsuite.Require().NoError(err)

	// From step MongoDBResources_DeleteMongoDBCollection
	mongoDBResourcesClientDeleteMongoDBCollectionResponsePoller, err := mongoDBResourcesClient.BeginDeleteMongoDBCollection(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, databaseName, collectionName, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, mongoDBResourcesClientDeleteMongoDBCollectionResponsePoller)
	testsuite.Require().NoError(err)

	// From step MongoDBResources_DeleteMongoDBDatabase
	mongoDBResourcesClientDeleteMongoDBDatabaseResponsePoller, err := mongoDBResourcesClient.BeginDeleteMongoDBDatabase(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, databaseName, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, mongoDBResourcesClientDeleteMongoDBDatabaseResponsePoller)
	testsuite.Require().NoError(err)
}
