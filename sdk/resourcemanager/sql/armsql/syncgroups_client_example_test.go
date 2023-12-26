//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SyncGroupGetSyncDatabaseId.json
func ExampleSyncGroupsClient_NewListSyncDatabaseIDsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSyncGroupsClient().NewListSyncDatabaseIDsPager("westus", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.SyncDatabaseIDListResult = armsql.SyncDatabaseIDListResult{
		// 	Value: []*armsql.SyncDatabaseIDProperties{
		// 		{
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/syncgroupcrud-3521/providers/Microsoft.Sql/servers/syncgroupcrud-8475/databases/syncgroupcrud-4328"),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SyncGroupRefreshHubSchema.json
func ExampleSyncGroupsClient_BeginRefreshHubSchema() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSyncGroupsClient().BeginRefreshHubSchema(ctx, "syncgroupcrud-65440", "syncgroupcrud-8475", "syncgroupcrud-4328", "syncgroupcrud-3187", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SyncGroupGetHubSchema.json
func ExampleSyncGroupsClient_NewListHubSchemasPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSyncGroupsClient().NewListHubSchemasPager("syncgroupcrud-65440", "syncgroupcrud-8475", "syncgroupcrud-4328", "syncgroupcrud-3187", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.SyncFullSchemaPropertiesListResult = armsql.SyncFullSchemaPropertiesListResult{
		// 	Value: []*armsql.SyncFullSchemaProperties{
		// 		{
		// 			LastUpdateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-30T07:16:08.210Z"); return t}()),
		// 			Tables: []*armsql.SyncFullSchemaTable{
		// 				{
		// 					Name: to.Ptr("dbo.myTable"),
		// 					Columns: []*armsql.SyncFullSchemaTableColumn{
		// 						{
		// 							Name: to.Ptr("intField"),
		// 							DataSize: to.Ptr("4"),
		// 							DataType: to.Ptr("int"),
		// 							HasError: to.Ptr(false),
		// 							IsPrimaryKey: to.Ptr(false),
		// 							QuotedName: to.Ptr("[intField]"),
		// 						},
		// 						{
		// 							Name: to.Ptr("charField"),
		// 							DataSize: to.Ptr("100"),
		// 							DataType: to.Ptr("nvarchar"),
		// 							HasError: to.Ptr(false),
		// 							IsPrimaryKey: to.Ptr(false),
		// 							QuotedName: to.Ptr("[charField]"),
		// 					}},
		// 					ErrorID: to.Ptr("Schema_TableHasNoPrimaryKey"),
		// 					HasError: to.Ptr(true),
		// 					QuotedName: to.Ptr("[dbo].[myTable]"),
		// 			}},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SyncGroupGetLog.json
func ExampleSyncGroupsClient_NewListLogsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSyncGroupsClient().NewListLogsPager("syncgroupcrud-65440", "syncgroupcrud-8475", "syncgroupcrud-4328", "syncgroupcrud-3187", "2017-01-01T00:00:00", "2017-12-31T00:00:00", armsql.SyncGroupsTypeAll, &armsql.SyncGroupsClientListLogsOptions{ContinuationToken: nil})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.SyncGroupLogListResult = armsql.SyncGroupLogListResult{
		// 	Value: []*armsql.SyncGroupLogProperties{
		// 		{
		// 			Type: to.Ptr(armsql.SyncGroupLogTypeSuccess),
		// 			OperationStatus: to.Ptr("SchemaRefreshSuccess"),
		// 			Source: to.Ptr("syncgroupcrud-8475.database.windows.net/hub"),
		// 			Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-30T07:16:08.250Z"); return t}()),
		// 			TracingID: to.Ptr("c0480c8e-6269-424e-9404-b00efce0ebae"),
		// 			Details: to.Ptr("Schema information obtained successfully."),
		// 		},
		// 		{
		// 			Type: to.Ptr(armsql.SyncGroupLogTypeError),
		// 			OperationStatus: to.Ptr("SchemaRefreshFailure"),
		// 			Source: to.Ptr("syncgroupcrud-8475.database.windows.net/member"),
		// 			Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-30T07:03:37.573Z"); return t}()),
		// 			TracingID: to.Ptr("cd3aa28c-0c31-471f-8a77-f1b21c908cbd"),
		// 			Details: to.Ptr("Getting schema information for the database failed with the exception \"Failed to connect to server .\nInner exception: SqlException Error Code: -2146232060 - SqlError Number:53, Message: A network-related or instance-specific error occurred while establishing a connection to SQL Server. The server was not found or was not accessible. Verify that the instance name is correct and that SQL Server is configured to allow remote connections. (provider: Named Pipes Provider, error: 40 - Could not open a connection to SQL Server) \nInner exception: The network path was not found\n For more information, provide tracing ID ‘cd3aa28c-0c31-471f-8a77-f1b21c908cbd’ to customer support.\""),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SyncGroupCancelSync.json
func ExampleSyncGroupsClient_CancelSync() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewSyncGroupsClient().CancelSync(ctx, "syncgroupcrud-65440", "syncgroupcrud-8475", "syncgroupcrud-4328", "syncgroupcrud-3187", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SyncGroupTriggerSync.json
func ExampleSyncGroupsClient_TriggerSync() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewSyncGroupsClient().TriggerSync(ctx, "syncgroupcrud-65440", "syncgroupcrud-8475", "syncgroupcrud-4328", "syncgroupcrud-3187", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SyncGroupGet.json
func ExampleSyncGroupsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSyncGroupsClient().Get(ctx, "syncgroupcrud-65440", "syncgroupcrud-8475", "syncgroupcrud-4328", "syncgroupcrud-3187", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SyncGroup = armsql.SyncGroup{
	// 	Name: to.Ptr("syncgroupcrud-3187"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/syncGroups"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/syncgroupcrud-3521/providers/Microsoft.Sql/servers/syncgroupcrud-8475/databases/syncgroupcrud-4328/syncGroups/syncgroupcrud-3187"),
	// 	Properties: &armsql.SyncGroupProperties{
	// 		ConflictResolutionPolicy: to.Ptr(armsql.SyncConflictResolutionPolicyHubWin),
	// 		HubDatabaseUserName: to.Ptr("hubUser"),
	// 		Interval: to.Ptr[int32](-1),
	// 		LastSyncTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T08:00:00.000Z"); return t}()),
	// 		PrivateEndpointName: to.Ptr("PE_67FDBBD6-B2D8-4014-9CC6-C68ABBCFD481_syncgroupcrud-3187"),
	// 		SyncDatabaseID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/syncgroupcrud-3521/providers/Microsoft.Sql/servers/syncgroupcrud-8475/databases/syncgroupcrud-4328"),
	// 		SyncState: to.Ptr(armsql.SyncGroupStateNotReady),
	// 		UsePrivateLinkConnection: to.Ptr(true),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SyncGroupCreate.json
func ExampleSyncGroupsClient_BeginCreateOrUpdate_createASyncGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSyncGroupsClient().BeginCreateOrUpdate(ctx, "syncgroupcrud-65440", "syncgroupcrud-8475", "syncgroupcrud-4328", "syncgroupcrud-3187", armsql.SyncGroup{
		Properties: &armsql.SyncGroupProperties{
			ConflictResolutionPolicy: to.Ptr(armsql.SyncConflictResolutionPolicyHubWin),
			HubDatabaseUserName:      to.Ptr("hubUser"),
			Interval:                 to.Ptr[int32](-1),
			SyncDatabaseID:           to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/syncgroupcrud-3521/providers/Microsoft.Sql/servers/syncgroupcrud-8475/databases/syncgroupcrud-4328"),
			UsePrivateLinkConnection: to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SyncGroup = armsql.SyncGroup{
	// 	Name: to.Ptr("syncgroupcrud-3187"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/syncGroups"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/syncgroupcrud-3521/providers/Microsoft.Sql/servers/syncgroupcrud-8475/databases/syncgroupcrud-4328/syncGroups/syncgroupcrud-3187"),
	// 	Properties: &armsql.SyncGroupProperties{
	// 		ConflictResolutionPolicy: to.Ptr(armsql.SyncConflictResolutionPolicyHubWin),
	// 		HubDatabaseUserName: to.Ptr("hubUser"),
	// 		Interval: to.Ptr[int32](-1),
	// 		LastSyncTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T08:00:00.000Z"); return t}()),
	// 		PrivateEndpointName: to.Ptr("PE_67FDBBD6-B2D8-4014-9CC6-C68ABBCFD481_syncgroupcrud-8475"),
	// 		SyncDatabaseID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/syncgroupcrud-3521/providers/Microsoft.Sql/servers/syncgroupcrud-8475/databases/syncgroupcrud-4328"),
	// 		SyncState: to.Ptr(armsql.SyncGroupStateNotReady),
	// 		UsePrivateLinkConnection: to.Ptr(true),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SyncGroupUpdate.json
func ExampleSyncGroupsClient_BeginCreateOrUpdate_updateASyncGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSyncGroupsClient().BeginCreateOrUpdate(ctx, "syncgroupcrud-65440", "syncgroupcrud-8475", "syncgroupcrud-4328", "syncgroupcrud-3187", armsql.SyncGroup{
		Properties: &armsql.SyncGroupProperties{
			ConflictResolutionPolicy: to.Ptr(armsql.SyncConflictResolutionPolicyHubWin),
			HubDatabaseUserName:      to.Ptr("hubUser"),
			Interval:                 to.Ptr[int32](-1),
			SyncDatabaseID:           to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/syncgroupcrud-3521/providers/Microsoft.Sql/servers/syncgroupcrud-8475/databases/syncgroupcrud-4328"),
			UsePrivateLinkConnection: to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SyncGroup = armsql.SyncGroup{
	// 	Name: to.Ptr("syncgroupcrud-3187"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/syncGroups"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/syncgroupcrud-3521/providers/Microsoft.Sql/servers/syncgroupcrud-8475/databases/syncgroupcrud-4328/syncGroups/syncgroupcrud-3187"),
	// 	Properties: &armsql.SyncGroupProperties{
	// 		ConflictResolutionPolicy: to.Ptr(armsql.SyncConflictResolutionPolicyHubWin),
	// 		HubDatabaseUserName: to.Ptr("hubUser"),
	// 		Interval: to.Ptr[int32](-1),
	// 		LastSyncTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T08:00:00.000Z"); return t}()),
	// 		PrivateEndpointName: to.Ptr("PE_67FDBBD6-B2D8-4014-9CC6-C68ABBCFD481_syncgroupcrud-3187"),
	// 		SyncDatabaseID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/syncgroupcrud-3521/providers/Microsoft.Sql/servers/syncgroupcrud-8475/databases/syncgroupcrud-4328"),
	// 		SyncState: to.Ptr(armsql.SyncGroupStateNotReady),
	// 		UsePrivateLinkConnection: to.Ptr(true),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SyncGroupDelete.json
func ExampleSyncGroupsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSyncGroupsClient().BeginDelete(ctx, "syncgroupcrud-65440", "syncgroupcrud-8475", "syncgroupcrud-4328", "syncgroupcrud-3187", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SyncGroupPatch.json
func ExampleSyncGroupsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSyncGroupsClient().BeginUpdate(ctx, "syncgroupcrud-65440", "syncgroupcrud-8475", "syncgroupcrud-4328", "syncgroupcrud-3187", armsql.SyncGroup{
		Properties: &armsql.SyncGroupProperties{
			ConflictResolutionPolicy: to.Ptr(armsql.SyncConflictResolutionPolicyHubWin),
			HubDatabasePassword:      to.Ptr("hubPassword"),
			HubDatabaseUserName:      to.Ptr("hubUser"),
			Interval:                 to.Ptr[int32](-1),
			SyncDatabaseID:           to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/syncgroupcrud-3521/providers/Microsoft.Sql/servers/syncgroupcrud-8475/databases/syncgroupcrud-4328"),
			UsePrivateLinkConnection: to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SyncGroup = armsql.SyncGroup{
	// 	Name: to.Ptr("syncgroup"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/syncGroups"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/syncgroupcrud-3521/providers/Microsoft.Sql/servers/syncgroupcrud-8475/databases/syncgroupcrud-4328/syncGroups/syncgroupcrud-3187"),
	// 	Properties: &armsql.SyncGroupProperties{
	// 		ConflictResolutionPolicy: to.Ptr(armsql.SyncConflictResolutionPolicyHubWin),
	// 		HubDatabaseUserName: to.Ptr("hubUser"),
	// 		Interval: to.Ptr[int32](-1),
	// 		LastSyncTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T08:00:00.000Z"); return t}()),
	// 		PrivateEndpointName: to.Ptr("PE_67FDBBD6-B2D8-4014-9CC6-C68ABBCFD481_syncgroupcrud-3187"),
	// 		SyncDatabaseID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/syncgroupcrud-3521/providers/Microsoft.Sql/servers/syncgroupcrud-8475/databases/syncgroupcrud-4328"),
	// 		SyncState: to.Ptr(armsql.SyncGroupStateNotReady),
	// 		UsePrivateLinkConnection: to.Ptr(true),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SyncGroupListByDatabase.json
func ExampleSyncGroupsClient_NewListByDatabasePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSyncGroupsClient().NewListByDatabasePager("syncgroupcrud-65440", "syncgroupcrud-8475", "syncgroupcrud-4328", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.SyncGroupListResult = armsql.SyncGroupListResult{
		// 	Value: []*armsql.SyncGroup{
		// 		{
		// 			Name: to.Ptr("syncgroupcrud-3187"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases/syncGroups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/syncgroupcrud-3521/providers/Microsoft.Sql/servers/syncgroupcrud-8475/databases/syncgroupcrud-4328/syncGroups/syncgroupcrud-3187"),
		// 			Properties: &armsql.SyncGroupProperties{
		// 				ConflictResolutionPolicy: to.Ptr(armsql.SyncConflictResolutionPolicyHubWin),
		// 				HubDatabaseUserName: to.Ptr("hubUser"),
		// 				Interval: to.Ptr[int32](-1),
		// 				LastSyncTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T08:00:00.000Z"); return t}()),
		// 				PrivateEndpointName: to.Ptr("PE_67FDBBD6-B2D8-4014-9CC6-C68ABBCFD481_syncgroupcrud-3187"),
		// 				SyncDatabaseID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/syncgroupcrud-3521/providers/Microsoft.Sql/servers/syncgroupcrud-8475/databases/syncgroupcrud-4328"),
		// 				SyncState: to.Ptr(armsql.SyncGroupStateNotReady),
		// 				UsePrivateLinkConnection: to.Ptr(true),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("syncgroupcrud-5374"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases/syncGroups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/syncgroupcrud-3521/providers/Microsoft.Sql/servers/syncgroupcrud-8475/databases/syncgroupcrud-4328/syncGroups/syncgroupcrud-5374"),
		// 			Properties: &armsql.SyncGroupProperties{
		// 				ConflictResolutionPolicy: to.Ptr(armsql.SyncConflictResolutionPolicyHubWin),
		// 				HubDatabaseUserName: to.Ptr("hubUser"),
		// 				Interval: to.Ptr[int32](-1),
		// 				LastSyncTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T08:00:00.000Z"); return t}()),
		// 				PrivateEndpointName: to.Ptr("PE_67FDBBD6-B2D8-4014-9CC6-C68ABBCFD481_syncgroupcrud-5374"),
		// 				SyncDatabaseID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/syncgroupcrud-3521/providers/Microsoft.Sql/servers/syncgroupcrud-8475/databases/syncgroupcrud-4328"),
		// 				SyncState: to.Ptr(armsql.SyncGroupStateNotReady),
		// 				UsePrivateLinkConnection: to.Ptr(true),
		// 			},
		// 	}},
		// }
	}
}
