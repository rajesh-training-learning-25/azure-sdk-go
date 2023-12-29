//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoDatabasesCheckNameAvailability.json
func ExampleDatabasesClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDatabasesClient().CheckNameAvailability(ctx, "kustorptest", "kustoCluster", armkusto.CheckNameRequest{
		Name: to.Ptr("database1"),
		Type: to.Ptr(armkusto.TypeMicrosoftKustoClustersDatabases),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckNameResult = armkusto.CheckNameResult{
	// 	Name: to.Ptr("database1"),
	// 	Message: to.Ptr("Name 'database1' is already taken. Please specify a different name"),
	// 	NameAvailable: to.Ptr(false),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoDatabasesListByCluster.json
func ExampleDatabasesClient_NewListByClusterPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDatabasesClient().NewListByClusterPager("kustorptest", "kustoCluster", &armkusto.DatabasesClientListByClusterOptions{Top: nil,
		Skiptoken: nil,
	})
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
		// page.DatabaseListResult = armkusto.DatabaseListResult{
		// 	Value: []armkusto.DatabaseClassification{
		// 		&armkusto.ReadWriteDatabase{
		// 			Name: to.Ptr("kustoCluster/KustoDatabase8"),
		// 			Type: to.Ptr("Microsoft.Kusto/Clusters/Databases"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/Databases/KustoDatabase8"),
		// 			Kind: to.Ptr(armkusto.KindReadWrite),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armkusto.ReadWriteDatabaseProperties{
		// 				HotCachePeriod: to.Ptr("P1D"),
		// 				ProvisioningState: to.Ptr(armkusto.ProvisioningStateSucceeded),
		// 				SoftDeletePeriod: to.Ptr("P1D"),
		// 			},
		// 		},
		// 		&armkusto.ReadOnlyFollowingDatabase{
		// 			Name: to.Ptr("kustoCluster/KustoDatabase9"),
		// 			Type: to.Ptr("Microsoft.Kusto/Clusters/Databases"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/Databases/KustoDatabase9"),
		// 			Kind: to.Ptr(armkusto.KindReadOnlyFollowing),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armkusto.ReadOnlyFollowingDatabaseProperties{
		// 				AttachedDatabaseConfigurationName: to.Ptr("attachedDatabaseConfigurationsTest"),
		// 				DatabaseShareOrigin: to.Ptr(armkusto.DatabaseShareOriginDirect),
		// 				HotCachePeriod: to.Ptr("P1D"),
		// 				LeaderClusterResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster2"),
		// 				OriginalDatabaseName: to.Ptr("KustoDatabase9"),
		// 				PrincipalsModificationKind: to.Ptr(armkusto.PrincipalsModificationKindUnion),
		// 				ProvisioningState: to.Ptr(armkusto.ProvisioningStateSucceeded),
		// 				SoftDeletePeriod: to.Ptr("P1D"),
		// 				Statistics: &armkusto.DatabaseStatistics{
		// 					Size: to.Ptr[float32](1024),
		// 				},
		// 				TableLevelSharingProperties: &armkusto.TableLevelSharingProperties{
		// 					ExternalTablesToExclude: []*string{
		// 						to.Ptr("ExternalTable2")},
		// 						ExternalTablesToInclude: []*string{
		// 							to.Ptr("ExternalTable1")},
		// 							MaterializedViewsToExclude: []*string{
		// 								to.Ptr("MaterializedViewTable2")},
		// 								MaterializedViewsToInclude: []*string{
		// 									to.Ptr("MaterializedViewTable1")},
		// 									TablesToExclude: []*string{
		// 										to.Ptr("Table2")},
		// 										TablesToInclude: []*string{
		// 											to.Ptr("Table1")},
		// 										},
		// 									},
		// 								},
		// 								&armkusto.ReadWriteDatabase{
		// 									Name: to.Ptr("kustoCluster/KustoDatabase10"),
		// 									Type: to.Ptr("Microsoft.Kusto/Clusters/Databases"),
		// 									ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/Databases/KustoDatabase10"),
		// 									Kind: to.Ptr(armkusto.KindReadWrite),
		// 									Location: to.Ptr("westus"),
		// 									Properties: &armkusto.ReadWriteDatabaseProperties{
		// 										HotCachePeriod: to.Ptr("P1D"),
		// 										ProvisioningState: to.Ptr(armkusto.ProvisioningStateSucceeded),
		// 										SoftDeletePeriod: to.Ptr("P1D"),
		// 										SuspensionDetails: &armkusto.SuspensionDetails{
		// 											SuspensionStartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-04-16T15:06:54.275Z"); return t}()),
		// 										},
		// 									},
		// 							}},
		// 						}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoDatabasesGet.json
func ExampleDatabasesClient_Get_kustoDatabasesGet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDatabasesClient().Get(ctx, "kustorptest", "kustoCluster", "KustoDatabase8", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armkusto.DatabasesClientGetResponse{
	// 	                            DatabaseClassification: &armkusto.ReadWriteDatabase{
	// 		Name: to.Ptr("kustoCluster/KustoDatabase8"),
	// 		Type: to.Ptr("Microsoft.Kusto/Clusters/Databases"),
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/Databases/KustoDatabase8"),
	// 		Kind: to.Ptr(armkusto.KindReadWrite),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armkusto.ReadWriteDatabaseProperties{
	// 			ProvisioningState: to.Ptr(armkusto.ProvisioningStateSucceeded),
	// 			SoftDeletePeriod: to.Ptr("P1D"),
	// 		},
	// 	},
	// 	                        }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoSuspendedDatabasesGet.json
func ExampleDatabasesClient_Get_kustoSuspendedDatabasesGet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDatabasesClient().Get(ctx, "kustorptest", "kustoCluster", "KustoDatabase9", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armkusto.DatabasesClientGetResponse{
	// 	                            DatabaseClassification: &armkusto.ReadWriteDatabase{
	// 		Name: to.Ptr("kustoCluster/KustoDatabase9"),
	// 		Type: to.Ptr("Microsoft.Kusto/Clusters/Databases"),
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/Databases/KustoDatabase9"),
	// 		Kind: to.Ptr(armkusto.KindReadWrite),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armkusto.ReadWriteDatabaseProperties{
	// 			ProvisioningState: to.Ptr(armkusto.ProvisioningStateSucceeded),
	// 			SoftDeletePeriod: to.Ptr("P1D"),
	// 			SuspensionDetails: &armkusto.SuspensionDetails{
	// 				SuspensionStartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-04-16T15:06:54.275Z"); return t}()),
	// 			},
	// 		},
	// 	},
	// 	                        }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoDatabaseReadonlyUpdate.json
func ExampleDatabasesClient_BeginCreateOrUpdate_kustoReadOnlyDatabaseUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDatabasesClient().BeginCreateOrUpdate(ctx, "kustorptest", "kustoCluster", "kustoReadOnlyDatabase", &armkusto.ReadOnlyFollowingDatabase{
		Kind:     to.Ptr(armkusto.KindReadOnlyFollowing),
		Location: to.Ptr("westus"),
		Properties: &armkusto.ReadOnlyFollowingDatabaseProperties{
			HotCachePeriod: to.Ptr("P1D"),
		},
	}, &armkusto.DatabasesClientBeginCreateOrUpdateOptions{CallerRole: nil})
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
	// res = armkusto.DatabasesClientCreateOrUpdateResponse{
	// 	                            DatabaseClassification: &armkusto.ReadOnlyFollowingDatabase{
	// 		Name: to.Ptr("kustoCluster/kustoReadOnlyDatabase"),
	// 		Type: to.Ptr("Microsoft.Kusto/Clusters/Databases"),
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/Databases/kustoReadOnlyDatabase"),
	// 		Kind: to.Ptr(armkusto.KindReadOnlyFollowing),
	// 		Location: to.Ptr("westus"),
	// 	},
	// 	                        }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoDatabasesCreateOrUpdate.json
func ExampleDatabasesClient_BeginCreateOrUpdate_kustoReadWriteDatabaseCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDatabasesClient().BeginCreateOrUpdate(ctx, "kustorptest", "kustoCluster", "KustoDatabase8", &armkusto.ReadWriteDatabase{
		Kind:     to.Ptr(armkusto.KindReadWrite),
		Location: to.Ptr("westus"),
		Properties: &armkusto.ReadWriteDatabaseProperties{
			SoftDeletePeriod: to.Ptr("P1D"),
		},
	}, &armkusto.DatabasesClientBeginCreateOrUpdateOptions{CallerRole: to.Ptr(armkusto.CallerRoleAdmin)})
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
	// res = armkusto.DatabasesClientCreateOrUpdateResponse{
	// 	                            DatabaseClassification: &armkusto.ReadWriteDatabase{
	// 		Name: to.Ptr("kustoCluster/KustoDatabase8"),
	// 		Type: to.Ptr("Microsoft.Kusto/Clusters/Databases"),
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/Databases/KustoDatabase8"),
	// 		Kind: to.Ptr(armkusto.KindReadWrite),
	// 		Location: to.Ptr("westus"),
	// 	},
	// 	                        }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoDatabasesUpdate.json
func ExampleDatabasesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDatabasesClient().BeginUpdate(ctx, "kustorptest", "kustoCluster", "KustoDatabase8", &armkusto.ReadWriteDatabase{
		Kind: to.Ptr(armkusto.KindReadWrite),
		Properties: &armkusto.ReadWriteDatabaseProperties{
			HotCachePeriod: to.Ptr("P1D"),
		},
	}, &armkusto.DatabasesClientBeginUpdateOptions{CallerRole: nil})
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
	// res = armkusto.DatabasesClientUpdateResponse{
	// 	                            DatabaseClassification: &armkusto.ReadWriteDatabase{
	// 		Name: to.Ptr("kustoCluster/KustoDatabase8"),
	// 		Type: to.Ptr("Microsoft.Kusto/Clusters/Databases"),
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/Databases/KustoDatabase8"),
	// 		Kind: to.Ptr(armkusto.KindReadWrite),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armkusto.ReadWriteDatabaseProperties{
	// 			HotCachePeriod: to.Ptr("P1D"),
	// 			ProvisioningState: to.Ptr(armkusto.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// 	                        }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoDatabasesDelete.json
func ExampleDatabasesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDatabasesClient().BeginDelete(ctx, "kustorptest", "kustoCluster", "KustoDatabase8", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoDatabaseListPrincipals.json
func ExampleDatabasesClient_NewListPrincipalsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDatabasesClient().NewListPrincipalsPager("kustorptest", "kustoCluster", "KustoDatabase8", nil)
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
		// page.DatabasePrincipalListResult = armkusto.DatabasePrincipalListResult{
		// 	Value: []*armkusto.DatabasePrincipal{
		// 		{
		// 			Name: to.Ptr("Some User"),
		// 			Type: to.Ptr(armkusto.DatabasePrincipalTypeUser),
		// 			AppID: to.Ptr(""),
		// 			Email: to.Ptr("user@microsoft.com"),
		// 			Fqn: to.Ptr("aaduser=some_guid"),
		// 			Role: to.Ptr(armkusto.DatabasePrincipalRoleAdmin),
		// 		},
		// 		{
		// 			Name: to.Ptr("Kusto"),
		// 			Type: to.Ptr(armkusto.DatabasePrincipalTypeGroup),
		// 			AppID: to.Ptr(""),
		// 			Email: to.Ptr("kusto@microsoft.com"),
		// 			Fqn: to.Ptr("aadgroup=some_guid"),
		// 			Role: to.Ptr(armkusto.DatabasePrincipalRoleViewer),
		// 		},
		// 		{
		// 			Name: to.Ptr("SomeApp"),
		// 			Type: to.Ptr(armkusto.DatabasePrincipalTypeApp),
		// 			AppID: to.Ptr("some_guid_app_id"),
		// 			Email: to.Ptr(""),
		// 			Fqn: to.Ptr("aadapp=some_guid_app_id"),
		// 			Role: to.Ptr(armkusto.DatabasePrincipalRoleAdmin),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoDatabaseAddPrincipals.json
func ExampleDatabasesClient_AddPrincipals() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDatabasesClient().AddPrincipals(ctx, "kustorptest", "kustoCluster", "KustoDatabase8", armkusto.DatabasePrincipalListRequest{
		Value: []*armkusto.DatabasePrincipal{
			{
				Name:  to.Ptr("Some User"),
				Type:  to.Ptr(armkusto.DatabasePrincipalTypeUser),
				AppID: to.Ptr(""),
				Email: to.Ptr("user@microsoft.com"),
				Fqn:   to.Ptr("aaduser=some_guid"),
				Role:  to.Ptr(armkusto.DatabasePrincipalRoleAdmin),
			},
			{
				Name:  to.Ptr("Kusto"),
				Type:  to.Ptr(armkusto.DatabasePrincipalTypeGroup),
				AppID: to.Ptr(""),
				Email: to.Ptr("kusto@microsoft.com"),
				Fqn:   to.Ptr("aadgroup=some_guid"),
				Role:  to.Ptr(armkusto.DatabasePrincipalRoleViewer),
			},
			{
				Name:  to.Ptr("SomeApp"),
				Type:  to.Ptr(armkusto.DatabasePrincipalTypeApp),
				AppID: to.Ptr("some_guid_app_id"),
				Email: to.Ptr(""),
				Fqn:   to.Ptr("aadapp=some_guid_app_id"),
				Role:  to.Ptr(armkusto.DatabasePrincipalRoleAdmin),
			}},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DatabasePrincipalListResult = armkusto.DatabasePrincipalListResult{
	// 	Value: []*armkusto.DatabasePrincipal{
	// 		{
	// 			Name: to.Ptr("Some User"),
	// 			Type: to.Ptr(armkusto.DatabasePrincipalTypeUser),
	// 			AppID: to.Ptr(""),
	// 			Email: to.Ptr("user@microsoft.com"),
	// 			Fqn: to.Ptr("aaduser=some_guid"),
	// 			Role: to.Ptr(armkusto.DatabasePrincipalRoleAdmin),
	// 		},
	// 		{
	// 			Name: to.Ptr("Kusto"),
	// 			Type: to.Ptr(armkusto.DatabasePrincipalTypeGroup),
	// 			AppID: to.Ptr(""),
	// 			Email: to.Ptr("kusto@microsoft.com"),
	// 			Fqn: to.Ptr("aadgroup=some_guid"),
	// 			Role: to.Ptr(armkusto.DatabasePrincipalRoleViewer),
	// 		},
	// 		{
	// 			Name: to.Ptr("SomeApp"),
	// 			Type: to.Ptr(armkusto.DatabasePrincipalTypeApp),
	// 			AppID: to.Ptr("some_guid_app_id"),
	// 			Email: to.Ptr(""),
	// 			Fqn: to.Ptr("aadapp=some_guid_app_id"),
	// 			Role: to.Ptr(armkusto.DatabasePrincipalRoleAdmin),
	// 	}},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoDatabaseRemovePrincipals.json
func ExampleDatabasesClient_RemovePrincipals() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDatabasesClient().RemovePrincipals(ctx, "kustorptest", "kustoCluster", "KustoDatabase8", armkusto.DatabasePrincipalListRequest{
		Value: []*armkusto.DatabasePrincipal{
			{
				Name:  to.Ptr("Some User"),
				Type:  to.Ptr(armkusto.DatabasePrincipalTypeUser),
				AppID: to.Ptr(""),
				Email: to.Ptr("user@microsoft.com"),
				Fqn:   to.Ptr("aaduser=some_guid"),
				Role:  to.Ptr(armkusto.DatabasePrincipalRoleAdmin),
			},
			{
				Name:  to.Ptr("Kusto"),
				Type:  to.Ptr(armkusto.DatabasePrincipalTypeGroup),
				AppID: to.Ptr(""),
				Email: to.Ptr("kusto@microsoft.com"),
				Fqn:   to.Ptr("aadgroup=some_guid"),
				Role:  to.Ptr(armkusto.DatabasePrincipalRoleViewer),
			},
			{
				Name:  to.Ptr("SomeApp"),
				Type:  to.Ptr(armkusto.DatabasePrincipalTypeApp),
				AppID: to.Ptr("some_guid_app_id"),
				Email: to.Ptr(""),
				Fqn:   to.Ptr("aadapp=some_guid_app_id"),
				Role:  to.Ptr(armkusto.DatabasePrincipalRoleAdmin),
			}},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DatabasePrincipalListResult = armkusto.DatabasePrincipalListResult{
	// 	Value: []*armkusto.DatabasePrincipal{
	// 		{
	// 			Name: to.Ptr("Some User"),
	// 			Type: to.Ptr(armkusto.DatabasePrincipalTypeUser),
	// 			AppID: to.Ptr(""),
	// 			Email: to.Ptr("user@microsoft.com"),
	// 			Fqn: to.Ptr("aaduser=some_guid"),
	// 			Role: to.Ptr(armkusto.DatabasePrincipalRoleAdmin),
	// 		},
	// 		{
	// 			Name: to.Ptr("Kusto"),
	// 			Type: to.Ptr(armkusto.DatabasePrincipalTypeGroup),
	// 			AppID: to.Ptr(""),
	// 			Email: to.Ptr("kusto@microsoft.com"),
	// 			Fqn: to.Ptr("aadgroup=some_guid"),
	// 			Role: to.Ptr(armkusto.DatabasePrincipalRoleViewer),
	// 		},
	// 		{
	// 			Name: to.Ptr("SomeApp"),
	// 			Type: to.Ptr(armkusto.DatabasePrincipalTypeApp),
	// 			AppID: to.Ptr("some_guid_app_id"),
	// 			Email: to.Ptr(""),
	// 			Fqn: to.Ptr("aadapp=some_guid_app_id"),
	// 			Role: to.Ptr(armkusto.DatabasePrincipalRoleAdmin),
	// 	}},
	// }
}
