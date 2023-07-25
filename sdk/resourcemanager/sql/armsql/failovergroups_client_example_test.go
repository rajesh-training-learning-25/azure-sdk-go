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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edd7863aeaad7379a7a8eef424ca11617d8c91eb/specification/sql/resource-manager/Microsoft.Sql/preview/2022-11-01-preview/examples/FailoverGroupList.json
func ExampleFailoverGroupsClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFailoverGroupsClient().NewListByServerPager("Default", "failover-group-primary-server", nil)
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
		// page.FailoverGroupListResult = armsql.FailoverGroupListResult{
		// 	Value: []*armsql.FailoverGroup{
		// 		{
		// 			Name: to.Ptr("failover-group-test"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/failoverGroups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-primary-server/failoverGroups/failover-group-test"),
		// 			Location: to.Ptr("Japan East"),
		// 			Properties: &armsql.FailoverGroupProperties{
		// 				Databases: []*string{
		// 				},
		// 				PartnerServers: []*armsql.PartnerInfo{
		// 					{
		// 						ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-secondary-server"),
		// 						Location: to.Ptr("Japan West"),
		// 						ReplicationRole: to.Ptr(armsql.FailoverGroupReplicationRoleSecondary),
		// 				}},
		// 				ReadOnlyEndpoint: &armsql.FailoverGroupReadOnlyEndpoint{
		// 					FailoverPolicy: to.Ptr(armsql.ReadOnlyEndpointFailoverPolicyDisabled),
		// 				},
		// 				ReadWriteEndpoint: &armsql.FailoverGroupReadWriteEndpoint{
		// 					FailoverPolicy: to.Ptr(armsql.ReadWriteEndpointFailoverPolicyAutomatic),
		// 					FailoverWithDataLossGracePeriodMinutes: to.Ptr[int32](480),
		// 				},
		// 				ReplicationRole: to.Ptr(armsql.FailoverGroupReplicationRolePrimary),
		// 				ReplicationState: to.Ptr("CATCH_UP"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("failover-group-test-2"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/failoverGroups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-primary-server/failoverGroups/failover-group-test-2"),
		// 			Location: to.Ptr("Japan East"),
		// 			Properties: &armsql.FailoverGroupProperties{
		// 				Databases: []*string{
		// 				},
		// 				PartnerServers: []*armsql.PartnerInfo{
		// 					{
		// 						ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-secondary-server"),
		// 						Location: to.Ptr("Japan West"),
		// 						ReplicationRole: to.Ptr(armsql.FailoverGroupReplicationRoleSecondary),
		// 				}},
		// 				ReadOnlyEndpoint: &armsql.FailoverGroupReadOnlyEndpoint{
		// 					FailoverPolicy: to.Ptr(armsql.ReadOnlyEndpointFailoverPolicyDisabled),
		// 				},
		// 				ReadWriteEndpoint: &armsql.FailoverGroupReadWriteEndpoint{
		// 					FailoverPolicy: to.Ptr(armsql.ReadWriteEndpointFailoverPolicyAutomatic),
		// 					FailoverWithDataLossGracePeriodMinutes: to.Ptr[int32](480),
		// 				},
		// 				ReplicationRole: to.Ptr(armsql.FailoverGroupReplicationRolePrimary),
		// 				ReplicationState: to.Ptr("CATCH_UP"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edd7863aeaad7379a7a8eef424ca11617d8c91eb/specification/sql/resource-manager/Microsoft.Sql/preview/2022-11-01-preview/examples/FailoverGroupGet.json
func ExampleFailoverGroupsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFailoverGroupsClient().Get(ctx, "Default", "failover-group-primary-server", "failover-group-test", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FailoverGroup = armsql.FailoverGroup{
	// 	Name: to.Ptr("failover-group-test"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/failoverGroups"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-primary-server/failoverGroups/failover-group-test"),
	// 	Location: to.Ptr("Japan East"),
	// 	Properties: &armsql.FailoverGroupProperties{
	// 		Databases: []*string{
	// 		},
	// 		PartnerServers: []*armsql.PartnerInfo{
	// 			{
	// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-secondary-server"),
	// 				Location: to.Ptr("Japan West"),
	// 				ReplicationRole: to.Ptr(armsql.FailoverGroupReplicationRoleSecondary),
	// 		}},
	// 		ReadOnlyEndpoint: &armsql.FailoverGroupReadOnlyEndpoint{
	// 			FailoverPolicy: to.Ptr(armsql.ReadOnlyEndpointFailoverPolicyDisabled),
	// 		},
	// 		ReadWriteEndpoint: &armsql.FailoverGroupReadWriteEndpoint{
	// 			FailoverPolicy: to.Ptr(armsql.ReadWriteEndpointFailoverPolicyAutomatic),
	// 			FailoverWithDataLossGracePeriodMinutes: to.Ptr[int32](480),
	// 		},
	// 		ReplicationRole: to.Ptr(armsql.FailoverGroupReplicationRolePrimary),
	// 		ReplicationState: to.Ptr("CATCH_UP"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edd7863aeaad7379a7a8eef424ca11617d8c91eb/specification/sql/resource-manager/Microsoft.Sql/preview/2022-11-01-preview/examples/FailoverGroupCreateOrUpdate.json
func ExampleFailoverGroupsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFailoverGroupsClient().BeginCreateOrUpdate(ctx, "Default", "failover-group-primary-server", "failover-group-test-3", armsql.FailoverGroup{
		Properties: &armsql.FailoverGroupProperties{
			Databases: []*string{
				to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-primary-server/databases/testdb-1"),
				to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-primary-server/databases/testdb-2")},
			PartnerServers: []*armsql.PartnerInfo{
				{
					ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-secondary-server"),
				}},
			ReadOnlyEndpoint: &armsql.FailoverGroupReadOnlyEndpoint{
				FailoverPolicy: to.Ptr(armsql.ReadOnlyEndpointFailoverPolicyDisabled),
			},
			ReadWriteEndpoint: &armsql.FailoverGroupReadWriteEndpoint{
				FailoverPolicy:                         to.Ptr(armsql.ReadWriteEndpointFailoverPolicyAutomatic),
				FailoverWithDataLossGracePeriodMinutes: to.Ptr[int32](480),
			},
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
	// res.FailoverGroup = armsql.FailoverGroup{
	// 	Name: to.Ptr("failover-group-test-3"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/failoverGroups"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-primary-server/failoverGroups/failover-group-test-3"),
	// 	Location: to.Ptr("Japan East"),
	// 	Properties: &armsql.FailoverGroupProperties{
	// 		Databases: []*string{
	// 			to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-primary-server/databases/testdb-1"),
	// 			to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-primary-server/databases/testdb-2")},
	// 			PartnerServers: []*armsql.PartnerInfo{
	// 				{
	// 					ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-secondary-server"),
	// 					Location: to.Ptr("Japan West"),
	// 					ReplicationRole: to.Ptr(armsql.FailoverGroupReplicationRoleSecondary),
	// 			}},
	// 			ReadOnlyEndpoint: &armsql.FailoverGroupReadOnlyEndpoint{
	// 				FailoverPolicy: to.Ptr(armsql.ReadOnlyEndpointFailoverPolicyDisabled),
	// 			},
	// 			ReadWriteEndpoint: &armsql.FailoverGroupReadWriteEndpoint{
	// 				FailoverPolicy: to.Ptr(armsql.ReadWriteEndpointFailoverPolicyAutomatic),
	// 				FailoverWithDataLossGracePeriodMinutes: to.Ptr[int32](480),
	// 			},
	// 			ReplicationRole: to.Ptr(armsql.FailoverGroupReplicationRolePrimary),
	// 			ReplicationState: to.Ptr("CATCH_UP"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edd7863aeaad7379a7a8eef424ca11617d8c91eb/specification/sql/resource-manager/Microsoft.Sql/preview/2022-11-01-preview/examples/FailoverGroupDelete.json
func ExampleFailoverGroupsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFailoverGroupsClient().BeginDelete(ctx, "Default", "failover-group-primary-server", "failover-group-test-1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edd7863aeaad7379a7a8eef424ca11617d8c91eb/specification/sql/resource-manager/Microsoft.Sql/preview/2022-11-01-preview/examples/FailoverGroupUpdate.json
func ExampleFailoverGroupsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFailoverGroupsClient().BeginUpdate(ctx, "Default", "failover-group-primary-server", "failover-group-test-1", armsql.FailoverGroupUpdate{
		Properties: &armsql.FailoverGroupUpdateProperties{
			Databases: []*string{
				to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-primary-server/databases/testdb-1")},
			ReadWriteEndpoint: &armsql.FailoverGroupReadWriteEndpoint{
				FailoverPolicy:                         to.Ptr(armsql.ReadWriteEndpointFailoverPolicyAutomatic),
				FailoverWithDataLossGracePeriodMinutes: to.Ptr[int32](120),
			},
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
	// res.FailoverGroup = armsql.FailoverGroup{
	// 	Name: to.Ptr("failover-group-test-3"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/failoverGroups"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-primary-server/failoverGroups/failover-group-test-3"),
	// 	Location: to.Ptr("Japan East"),
	// 	Properties: &armsql.FailoverGroupProperties{
	// 		Databases: []*string{
	// 			to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-primary-server/databases/testdb-1")},
	// 			PartnerServers: []*armsql.PartnerInfo{
	// 				{
	// 					ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-secondary-server"),
	// 					Location: to.Ptr("Japan West"),
	// 					ReplicationRole: to.Ptr(armsql.FailoverGroupReplicationRoleSecondary),
	// 			}},
	// 			ReadOnlyEndpoint: &armsql.FailoverGroupReadOnlyEndpoint{
	// 				FailoverPolicy: to.Ptr(armsql.ReadOnlyEndpointFailoverPolicyDisabled),
	// 			},
	// 			ReadWriteEndpoint: &armsql.FailoverGroupReadWriteEndpoint{
	// 				FailoverPolicy: to.Ptr(armsql.ReadWriteEndpointFailoverPolicyAutomatic),
	// 				FailoverWithDataLossGracePeriodMinutes: to.Ptr[int32](120),
	// 			},
	// 			ReplicationRole: to.Ptr(armsql.FailoverGroupReplicationRolePrimary),
	// 			ReplicationState: to.Ptr("CATCH_UP"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edd7863aeaad7379a7a8eef424ca11617d8c91eb/specification/sql/resource-manager/Microsoft.Sql/preview/2022-11-01-preview/examples/FailoverGroupFailover.json
func ExampleFailoverGroupsClient_BeginFailover() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFailoverGroupsClient().BeginFailover(ctx, "Default", "failover-group-secondary-server", "failover-group-test-3", nil)
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
	// res.FailoverGroup = armsql.FailoverGroup{
	// 	Name: to.Ptr("failover-group-test-3"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/failoverGroups"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-secondary-server/failoverGroups/failover-group-test-3"),
	// 	Location: to.Ptr("Japan West"),
	// 	Properties: &armsql.FailoverGroupProperties{
	// 		Databases: []*string{
	// 			to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-secondary-server/databases/testdb-1"),
	// 			to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-secondary-server/databases/testdb-2")},
	// 			PartnerServers: []*armsql.PartnerInfo{
	// 				{
	// 					ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-primary-server"),
	// 					Location: to.Ptr("Japan East"),
	// 					ReplicationRole: to.Ptr(armsql.FailoverGroupReplicationRoleSecondary),
	// 			}},
	// 			ReadOnlyEndpoint: &armsql.FailoverGroupReadOnlyEndpoint{
	// 				FailoverPolicy: to.Ptr(armsql.ReadOnlyEndpointFailoverPolicyDisabled),
	// 			},
	// 			ReadWriteEndpoint: &armsql.FailoverGroupReadWriteEndpoint{
	// 				FailoverPolicy: to.Ptr(armsql.ReadWriteEndpointFailoverPolicyAutomatic),
	// 				FailoverWithDataLossGracePeriodMinutes: to.Ptr[int32](120),
	// 			},
	// 			ReplicationRole: to.Ptr(armsql.FailoverGroupReplicationRolePrimary),
	// 			ReplicationState: to.Ptr("CATCH_UP"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edd7863aeaad7379a7a8eef424ca11617d8c91eb/specification/sql/resource-manager/Microsoft.Sql/preview/2022-11-01-preview/examples/FailoverGroupForceFailoverAllowDataLoss.json
func ExampleFailoverGroupsClient_BeginForceFailoverAllowDataLoss() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFailoverGroupsClient().BeginForceFailoverAllowDataLoss(ctx, "Default", "failover-group-secondary-server", "failover-group-test-3", nil)
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
	// res.FailoverGroup = armsql.FailoverGroup{
	// 	Name: to.Ptr("failover-group-test-3"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/failoverGroups"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-secondary-server/failoverGroups/failover-group-test-3"),
	// 	Location: to.Ptr("Japan West"),
	// 	Properties: &armsql.FailoverGroupProperties{
	// 		Databases: []*string{
	// 			to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-secondary-server/databases/testdb-1"),
	// 			to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-secondary-server/databases/testdb-2")},
	// 			PartnerServers: []*armsql.PartnerInfo{
	// 				{
	// 					ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-primary-server"),
	// 					Location: to.Ptr("Japan East"),
	// 					ReplicationRole: to.Ptr(armsql.FailoverGroupReplicationRoleSecondary),
	// 			}},
	// 			ReadOnlyEndpoint: &armsql.FailoverGroupReadOnlyEndpoint{
	// 				FailoverPolicy: to.Ptr(armsql.ReadOnlyEndpointFailoverPolicyDisabled),
	// 			},
	// 			ReadWriteEndpoint: &armsql.FailoverGroupReadWriteEndpoint{
	// 				FailoverPolicy: to.Ptr(armsql.ReadWriteEndpointFailoverPolicyAutomatic),
	// 				FailoverWithDataLossGracePeriodMinutes: to.Ptr[int32](120),
	// 			},
	// 			ReplicationRole: to.Ptr(armsql.FailoverGroupReplicationRolePrimary),
	// 			ReplicationState: to.Ptr("CATCH_UP"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edd7863aeaad7379a7a8eef424ca11617d8c91eb/specification/sql/resource-manager/Microsoft.Sql/preview/2022-11-01-preview/examples/FailoverGroupTryPlannedBeforeForcedFailover.json
func ExampleFailoverGroupsClient_BeginTryPlannedBeforeForcedFailover() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFailoverGroupsClient().BeginTryPlannedBeforeForcedFailover(ctx, "Default", "failovergroupsecondaryserver", "failovergrouptest3", nil)
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
	// res.FailoverGroup = armsql.FailoverGroup{
	// 	Name: to.Ptr("failovergrouptest3"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/failoverGroups"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failovergroupsecondaryserver/failoverGroups/failovergrouptest3"),
	// 	Location: to.Ptr("Japan West"),
	// 	Properties: &armsql.FailoverGroupProperties{
	// 		Databases: []*string{
	// 			to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failovergroupsecondaryserver/databases/testdb1"),
	// 			to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failovergroupsecondaryserver/databases/testdb2")},
	// 			PartnerServers: []*armsql.PartnerInfo{
	// 				{
	// 					ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failovergroupprimaryserver"),
	// 					Location: to.Ptr("Japan East"),
	// 					ReplicationRole: to.Ptr(armsql.FailoverGroupReplicationRoleSecondary),
	// 			}},
	// 			ReadOnlyEndpoint: &armsql.FailoverGroupReadOnlyEndpoint{
	// 				FailoverPolicy: to.Ptr(armsql.ReadOnlyEndpointFailoverPolicyDisabled),
	// 			},
	// 			ReadWriteEndpoint: &armsql.FailoverGroupReadWriteEndpoint{
	// 				FailoverPolicy: to.Ptr(armsql.ReadWriteEndpointFailoverPolicyAutomatic),
	// 				FailoverWithDataLossGracePeriodMinutes: to.Ptr[int32](120),
	// 			},
	// 			ReplicationRole: to.Ptr(armsql.FailoverGroupReplicationRolePrimary),
	// 			ReplicationState: to.Ptr("CATCH_UP"),
	// 		},
	// 	}
}
