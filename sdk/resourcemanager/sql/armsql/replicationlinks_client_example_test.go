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

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edd7863aeaad7379a7a8eef424ca11617d8c91eb/specification/sql/resource-manager/Microsoft.Sql/preview/2022-02-01-preview/examples/ReplicationLinkListByDatabase.json
func ExampleReplicationLinksClient_NewListByDatabasePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReplicationLinksClient().NewListByDatabasePager("Default", "sourcesvr", "tetha-db", nil)
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
		// page.ReplicationLinkListResult = armsql.ReplicationLinkListResult{
		// 	Value: []*armsql.ReplicationLink{
		// 		{
		// 			Name: to.Ptr("fb92de60-eb87-4a58-b250-3362d0cfdf26"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/sourcesvr/databases/tetha-db/replicationLinks/fb92de60-eb87-4a58-b250-3362d0cfdf26"),
		// 			Properties: &armsql.ReplicationLinkProperties{
		// 				IsTerminationAllowed: to.Ptr(true),
		// 				LinkType: to.Ptr(armsql.ReplicationLinkTypeGEO),
		// 				PartnerDatabase: to.Ptr("tetha-db"),
		// 				PartnerLocation: to.Ptr("Japan East"),
		// 				PartnerRole: to.Ptr(armsql.ReplicationRoleSecondary),
		// 				PartnerServer: to.Ptr("testsvr"),
		// 				PercentComplete: to.Ptr[int32](100),
		// 				ReplicationMode: to.Ptr("ASYNC"),
		// 				ReplicationState: to.Ptr(armsql.ReplicationStateCATCHUP),
		// 				Role: to.Ptr(armsql.ReplicationRolePrimary),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-21T08:11:46.907Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("6ad1eefc-18a2-4fcb-94f3-4b654ba788d7"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/sourcesvr/databases/tetha-db/replicationLinks/6ad1eefc-18a2-4fcb-94f3-4b654ba788d7"),
		// 			Properties: &armsql.ReplicationLinkProperties{
		// 				IsTerminationAllowed: to.Ptr(true),
		// 				LinkType: to.Ptr(armsql.ReplicationLinkTypeGEO),
		// 				PartnerDatabase: to.Ptr("tetha-db"),
		// 				PartnerLocation: to.Ptr("Japan East"),
		// 				PartnerRole: to.Ptr(armsql.ReplicationRoleSecondary),
		// 				PartnerServer: to.Ptr("targetsvr"),
		// 				PercentComplete: to.Ptr[int32](100),
		// 				ReplicationMode: to.Ptr("ASYNC"),
		// 				ReplicationState: to.Ptr(armsql.ReplicationStateCATCHUP),
		// 				Role: to.Ptr(armsql.ReplicationRolePrimary),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-21T08:11:34.423Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edd7863aeaad7379a7a8eef424ca11617d8c91eb/specification/sql/resource-manager/Microsoft.Sql/preview/2022-02-01-preview/examples/ReplicationLinkGet.json
func ExampleReplicationLinksClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReplicationLinksClient().Get(ctx, "Default", "sourcesvr", "gamma-db", "4891ca10-ebd0-47d7-9182-c722651780fb", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ReplicationLink = armsql.ReplicationLink{
	// 	Name: to.Ptr("4891ca10-ebd0-47d7-9182-c722651780fb"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/sourcesvr/databases/gamma-db/replicationLinks/4891ca10-ebd0-47d7-9182-c722651780fb"),
	// 	Properties: &armsql.ReplicationLinkProperties{
	// 		IsTerminationAllowed: to.Ptr(true),
	// 		LinkType: to.Ptr(armsql.ReplicationLinkTypeGEO),
	// 		PartnerDatabase: to.Ptr("gamma-db"),
	// 		PartnerLocation: to.Ptr("Japan East"),
	// 		PartnerRole: to.Ptr(armsql.ReplicationRoleSecondary),
	// 		PartnerServer: to.Ptr("testsvr"),
	// 		PercentComplete: to.Ptr[int32](100),
	// 		ReplicationMode: to.Ptr("ASYNC"),
	// 		ReplicationState: to.Ptr(armsql.ReplicationStateCATCHUP),
	// 		Role: to.Ptr(armsql.ReplicationRolePrimary),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-21T08:12:43.783Z"); return t}()),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edd7863aeaad7379a7a8eef424ca11617d8c91eb/specification/sql/resource-manager/Microsoft.Sql/preview/2022-02-01-preview/examples/ReplicationLinkDelete.json
func ExampleReplicationLinksClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationLinksClient().BeginDelete(ctx, "Default", "sourcesvr", "gamma-db", "4891ca10-ebd0-47d7-9182-c722651780fb", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edd7863aeaad7379a7a8eef424ca11617d8c91eb/specification/sql/resource-manager/Microsoft.Sql/preview/2022-02-01-preview/examples/ReplicationLinkFailover.json
func ExampleReplicationLinksClient_BeginFailover() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationLinksClient().BeginFailover(ctx, "Default", "sourcesvr", "gamma-db", "4891ca10-ebd0-47d7-9182-c722651780fb", nil)
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
	// res.ReplicationLink = armsql.ReplicationLink{
	// 	Name: to.Ptr("4891ca10-ebd0-47d7-9182-c722651780fb"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/sourcesvr/databases/gamma-db/replicationLinks/4891ca10-ebd0-47d7-9182-c722651780fb"),
	// 	Properties: &armsql.ReplicationLinkProperties{
	// 		IsTerminationAllowed: to.Ptr(true),
	// 		LinkType: to.Ptr(armsql.ReplicationLinkTypeGEO),
	// 		PartnerDatabase: to.Ptr("gamma-db"),
	// 		PartnerLocation: to.Ptr("Japan East"),
	// 		PartnerRole: to.Ptr(armsql.ReplicationRoleSecondary),
	// 		PartnerServer: to.Ptr("testsvr"),
	// 		PercentComplete: to.Ptr[int32](100),
	// 		ReplicationMode: to.Ptr("ASYNC"),
	// 		ReplicationState: to.Ptr(armsql.ReplicationStateCATCHUP),
	// 		Role: to.Ptr(armsql.ReplicationRolePrimary),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-21T08:12:43.783Z"); return t}()),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edd7863aeaad7379a7a8eef424ca11617d8c91eb/specification/sql/resource-manager/Microsoft.Sql/preview/2022-02-01-preview/examples/ReplicationLinkFailoverAllowDataLoss.json
func ExampleReplicationLinksClient_BeginFailoverAllowDataLoss() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationLinksClient().BeginFailoverAllowDataLoss(ctx, "Default", "sourcesvr", "gamma-db", "4891ca10-ebd0-47d7-9182-c722651780fb", nil)
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
	// res.ReplicationLink = armsql.ReplicationLink{
	// 	Name: to.Ptr("4891ca10-ebd0-47d7-9182-c722651780fb"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/sourcesvr/databases/gamma-db/replicationLinks/4891ca10-ebd0-47d7-9182-c722651780fb"),
	// 	Properties: &armsql.ReplicationLinkProperties{
	// 		IsTerminationAllowed: to.Ptr(true),
	// 		LinkType: to.Ptr(armsql.ReplicationLinkTypeGEO),
	// 		PartnerDatabase: to.Ptr("gamma-db"),
	// 		PartnerLocation: to.Ptr("Japan East"),
	// 		PartnerRole: to.Ptr(armsql.ReplicationRoleSecondary),
	// 		PartnerServer: to.Ptr("testsvr"),
	// 		PercentComplete: to.Ptr[int32](100),
	// 		ReplicationMode: to.Ptr("ASYNC"),
	// 		ReplicationState: to.Ptr(armsql.ReplicationStateCATCHUP),
	// 		Role: to.Ptr(armsql.ReplicationRolePrimary),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-21T08:12:43.783Z"); return t}()),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edd7863aeaad7379a7a8eef424ca11617d8c91eb/specification/sql/resource-manager/Microsoft.Sql/preview/2022-02-01-preview/examples/ReplicationLinkListByServer.json
func ExampleReplicationLinksClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReplicationLinksClient().NewListByServerPager("Default", "sourcesvr", nil)
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
		// page.ReplicationLinkListResult = armsql.ReplicationLinkListResult{
		// 	Value: []*armsql.ReplicationLink{
		// 		{
		// 			Name: to.Ptr("fb92de60-eb87-4a58-b250-3362d0cfdf26"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-Onebox/providers/Microsoft.Sql/servers/sourcesvr/databases/tetha-db/replicationLinks/fb92de60-eb87-4a58-b250-3362d0cfdf26"),
		// 			Properties: &armsql.ReplicationLinkProperties{
		// 				IsTerminationAllowed: to.Ptr(true),
		// 				LinkType: to.Ptr(armsql.ReplicationLinkTypeGEO),
		// 				PartnerDatabase: to.Ptr("tetha-db"),
		// 				PartnerLocation: to.Ptr("Japan East"),
		// 				PartnerRole: to.Ptr(armsql.ReplicationRoleSecondary),
		// 				PartnerServer: to.Ptr("testsvr"),
		// 				PercentComplete: to.Ptr[int32](100),
		// 				ReplicationMode: to.Ptr("ASYNC"),
		// 				ReplicationState: to.Ptr(armsql.ReplicationStateCATCHUP),
		// 				Role: to.Ptr(armsql.ReplicationRolePrimary),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-21T08:11:46.907Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("6ad1eefc-18a2-4fcb-94f3-4b654ba788d7"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-Onebox/providers/Microsoft.Sql/servers/sourcesvr/databases/tetha-db/replicationLinks/6ad1eefc-18a2-4fcb-94f3-4b654ba788d7"),
		// 			Properties: &armsql.ReplicationLinkProperties{
		// 				IsTerminationAllowed: to.Ptr(true),
		// 				LinkType: to.Ptr(armsql.ReplicationLinkTypeGEO),
		// 				PartnerDatabase: to.Ptr("tetha-db"),
		// 				PartnerLocation: to.Ptr("Japan East"),
		// 				PartnerRole: to.Ptr(armsql.ReplicationRoleSecondary),
		// 				PartnerServer: to.Ptr("targetsvr"),
		// 				PercentComplete: to.Ptr[int32](100),
		// 				ReplicationMode: to.Ptr("ASYNC"),
		// 				ReplicationState: to.Ptr(armsql.ReplicationStateCATCHUP),
		// 				Role: to.Ptr(armsql.ReplicationRolePrimary),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-21T08:11:34.423Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("4891ca10-ebd0-47d7-9182-c722651780fb"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-Onebox/providers/Microsoft.Sql/servers/sourcesvr/databases/gamma-db/replicationLinks/4891ca10-ebd0-47d7-9182-c722651780fb"),
		// 			Properties: &armsql.ReplicationLinkProperties{
		// 				IsTerminationAllowed: to.Ptr(true),
		// 				LinkType: to.Ptr(armsql.ReplicationLinkTypeGEO),
		// 				PartnerDatabase: to.Ptr("gamma-db"),
		// 				PartnerLocation: to.Ptr("Japan East"),
		// 				PartnerRole: to.Ptr(armsql.ReplicationRoleSecondary),
		// 				PartnerServer: to.Ptr("testsvr"),
		// 				PercentComplete: to.Ptr[int32](100),
		// 				ReplicationMode: to.Ptr("ASYNC"),
		// 				ReplicationState: to.Ptr(armsql.ReplicationStateCATCHUP),
		// 				Role: to.Ptr(armsql.ReplicationRolePrimary),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-21T08:12:43.783Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
