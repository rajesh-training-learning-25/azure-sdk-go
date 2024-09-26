// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armstandbypool_test

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/standbypool/armstandbypool"
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"log"
)

// Generated from example definition: 2024-03-01/StandbyVirtualMachinePools_CreateOrUpdate.json
func ExampleStandbyVirtualMachinePoolsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstandbypool.NewClientFactory("00000000-0000-0000-0000-000000000009", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewStandbyVirtualMachinePoolsClient().BeginCreateOrUpdate(ctx, "rgstandbypool", "pool", armstandbypool.StandbyVirtualMachinePoolResource{
		Properties: &armstandbypool.StandbyVirtualMachinePoolResourceProperties{
			ElasticityProfile: &armstandbypool.StandbyVirtualMachinePoolElasticityProfile{
				MaxReadyCapacity: to.Ptr[int64](304),
				MinReadyCapacity: to.Ptr[int64](300),
			},
			VirtualMachineState:              to.Ptr(armstandbypool.VirtualMachineStateRunning),
			AttachedVirtualMachineScaleSetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000009/resourceGroups/rgstandbypool/providers/Microsoft.Compute/virtualMachineScaleSets/myVmss"),
		},
		Tags:     map[string]*string{},
		Location: to.Ptr("West US"),
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
	// res = armstandbypool.StandbyVirtualMachinePoolsClientCreateOrUpdateResponse{
	// 	StandbyVirtualMachinePoolResource: &armstandbypool.StandbyVirtualMachinePoolResource{
	// 		Properties: &armstandbypool.StandbyVirtualMachinePoolResourceProperties{
	// 			ElasticityProfile: &armstandbypool.StandbyVirtualMachinePoolElasticityProfile{
	// 				MaxReadyCapacity: to.Ptr[int64](304),
	// 				MinReadyCapacity: to.Ptr[int64](300),
	// 			},
	// 			VirtualMachineState: to.Ptr(armstandbypool.VirtualMachineStateRunning),
	// 			AttachedVirtualMachineScaleSetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000009/resourceGroups/rgstandbypool/providers/Microsoft.Compute/virtualMachineScaleSets/myVmss"),
	// 			ProvisioningState: to.Ptr(armstandbypool.ProvisioningStateSucceeded),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 		Location: to.Ptr("West US"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000009/resourceGroups/rgstandbypool/providers/Microsoft.StandbyPool/standbyVirtualMachinePools/pool"),
	// 		Name: to.Ptr("pool"),
	// 		Type: to.Ptr("Microsoft.StandbyPool/standbyVirtualMachinePools"),
	// 		SystemData: &armstandbypool.SystemData{
	// 			CreatedBy: to.Ptr("pooluser@microsoft.com"),
	// 			CreatedByType: to.Ptr(armstandbypool.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-07T16:33:22.210Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("pooluser@microsoft.com"),
	// 			LastModifiedByType: to.Ptr(armstandbypool.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-07T16:33:22.210Z"); return t}()),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2024-03-01/StandbyVirtualMachinePools_Delete.json
func ExampleStandbyVirtualMachinePoolsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstandbypool.NewClientFactory("00000000-0000-0000-0000-000000000009", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewStandbyVirtualMachinePoolsClient().BeginDelete(ctx, "rgstandbypool", "pool", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: 2024-03-01/StandbyVirtualMachinePools_Get.json
func ExampleStandbyVirtualMachinePoolsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstandbypool.NewClientFactory("00000000-0000-0000-0000-000000000009", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStandbyVirtualMachinePoolsClient().Get(ctx, "rgstandbypool", "pool", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armstandbypool.StandbyVirtualMachinePoolsClientGetResponse{
	// 	StandbyVirtualMachinePoolResource: &armstandbypool.StandbyVirtualMachinePoolResource{
	// 		Properties: &armstandbypool.StandbyVirtualMachinePoolResourceProperties{
	// 			ElasticityProfile: &armstandbypool.StandbyVirtualMachinePoolElasticityProfile{
	// 				MaxReadyCapacity: to.Ptr[int64](304),
	// 				MinReadyCapacity: to.Ptr[int64](300),
	// 			},
	// 			VirtualMachineState: to.Ptr(armstandbypool.VirtualMachineStateRunning),
	// 			AttachedVirtualMachineScaleSetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000009/resourceGroups/rgstandbypool/providers/Microsoft.Compute/virtualMachineScaleSets/myVmss"),
	// 			ProvisioningState: to.Ptr(armstandbypool.ProvisioningStateSucceeded),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 		Location: to.Ptr("West US"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000009/resourceGroups/rgstandbypool/providers/Microsoft.StandbyPool/standbyVirtualMachinePools/pool"),
	// 		Name: to.Ptr("pool"),
	// 		Type: to.Ptr("Microsoft.StandbyPool/standbyVirtualMachinePools"),
	// 		SystemData: &armstandbypool.SystemData{
	// 			CreatedBy: to.Ptr("pooluser@microsoft.com"),
	// 			CreatedByType: to.Ptr(armstandbypool.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-07T16:33:22.210Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("pooluser@microsoft.com"),
	// 			LastModifiedByType: to.Ptr(armstandbypool.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-07T16:33:22.210Z"); return t}()),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2024-03-01/StandbyVirtualMachinePools_ListByResourceGroup.json
func ExampleStandbyVirtualMachinePoolsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstandbypool.NewClientFactory("00000000-0000-0000-0000-000000000009", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewStandbyVirtualMachinePoolsClient().NewListByResourceGroupPager("rgstandbypool", nil)
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
		// page = armstandbypool.StandbyVirtualMachinePoolsClientListByResourceGroupResponse{
		// 	StandbyVirtualMachinePoolResourceListResult: armstandbypool.StandbyVirtualMachinePoolResourceListResult{
		// 		Value: []*armstandbypool.StandbyVirtualMachinePoolResource{
		// 			{
		// 				Properties: &armstandbypool.StandbyVirtualMachinePoolResourceProperties{
		// 					ElasticityProfile: &armstandbypool.StandbyVirtualMachinePoolElasticityProfile{
		// 						MaxReadyCapacity: to.Ptr[int64](304),
		// 					},
		// 					VirtualMachineState: to.Ptr(armstandbypool.VirtualMachineStateRunning),
		// 					AttachedVirtualMachineScaleSetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000009/resourceGroups/rgstandbypool/providers/Microsoft.Compute/virtualMachineScaleSets/myVmss"),
		// 					ProvisioningState: to.Ptr(armstandbypool.ProvisioningStateSucceeded),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 				Location: to.Ptr("West US"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000009/resourceGroups/rgstandbypool/providers/Microsoft.StandbyPool/standbyVirtualMachinePools/pool"),
		// 				Name: to.Ptr("pool"),
		// 				Type: to.Ptr("Microsoft.StandbyPool/standbyVirtualMachinePools"),
		// 				SystemData: &armstandbypool.SystemData{
		// 					CreatedBy: to.Ptr("pooluser@microsoft.com"),
		// 					CreatedByType: to.Ptr(armstandbypool.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-07T16:33:22.210Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("pooluser@microsoft.com"),
		// 					LastModifiedByType: to.Ptr(armstandbypool.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-07T16:33:22.210Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://example.com/subscriptions/00000000-0000-0000-0000-000000000009/resourceGroups/rgstandbypool/providers/Microsoft.StandbyPool/standbyVirtualMachinePools"),
		// 	},
		// }
	}
}

// Generated from example definition: 2024-03-01/StandbyVirtualMachinePools_ListBySubscription.json
func ExampleStandbyVirtualMachinePoolsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstandbypool.NewClientFactory("00000000-0000-0000-0000-000000000009", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewStandbyVirtualMachinePoolsClient().NewListBySubscriptionPager(nil)
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
		// page = armstandbypool.StandbyVirtualMachinePoolsClientListBySubscriptionResponse{
		// 	StandbyVirtualMachinePoolResourceListResult: armstandbypool.StandbyVirtualMachinePoolResourceListResult{
		// 		Value: []*armstandbypool.StandbyVirtualMachinePoolResource{
		// 			{
		// 				Properties: &armstandbypool.StandbyVirtualMachinePoolResourceProperties{
		// 					ElasticityProfile: &armstandbypool.StandbyVirtualMachinePoolElasticityProfile{
		// 						MaxReadyCapacity: to.Ptr[int64](304),
		// 						MinReadyCapacity: to.Ptr[int64](300),
		// 					},
		// 					VirtualMachineState: to.Ptr(armstandbypool.VirtualMachineStateRunning),
		// 					AttachedVirtualMachineScaleSetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000009/resourceGroups/rgstandbypool/providers/Microsoft.Compute/virtualMachineScaleSets/myVmss"),
		// 					ProvisioningState: to.Ptr(armstandbypool.ProvisioningStateSucceeded),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 				Location: to.Ptr("West US"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000009/resourceGroups/rgstandbypool/providers/Microsoft.StandbyPool/standbyVirtualMachinePools/pool"),
		// 				Name: to.Ptr("pool"),
		// 				Type: to.Ptr("Microsoft.StandbyPool/standbyVirtualMachinePools"),
		// 				SystemData: &armstandbypool.SystemData{
		// 					CreatedBy: to.Ptr("pooluser@microsoft.com"),
		// 					CreatedByType: to.Ptr(armstandbypool.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-07T16:33:22.210Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("pooluser@microsoft.com"),
		// 					LastModifiedByType: to.Ptr(armstandbypool.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-07T16:33:22.210Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://example.com/subscriptions/00000000-0000-0000-0000-000000000009/resourceGroups/rgstandbypool/providers/Microsoft.StandbyPool/standbyVirtualMachinePools"),
		// 	},
		// }
	}
}

// Generated from example definition: 2024-03-01/StandbyVirtualMachinePools_Update.json
func ExampleStandbyVirtualMachinePoolsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstandbypool.NewClientFactory("00000000-0000-0000-0000-000000000009", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStandbyVirtualMachinePoolsClient().Update(ctx, "rgstandbypool", "pool", armstandbypool.StandbyVirtualMachinePoolResourceUpdate{
		Tags: map[string]*string{},
		Properties: &armstandbypool.StandbyVirtualMachinePoolResourceUpdateProperties{
			ElasticityProfile: &armstandbypool.StandbyVirtualMachinePoolElasticityProfile{
				MaxReadyCapacity: to.Ptr[int64](304),
				MinReadyCapacity: to.Ptr[int64](300),
			},
			VirtualMachineState:              to.Ptr(armstandbypool.VirtualMachineStateRunning),
			AttachedVirtualMachineScaleSetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000009/resourceGroups/rgstandbypool/providers/Microsoft.Compute/virtualMachineScaleSets/myVmss"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armstandbypool.StandbyVirtualMachinePoolsClientUpdateResponse{
	// 	StandbyVirtualMachinePoolResource: &armstandbypool.StandbyVirtualMachinePoolResource{
	// 		Properties: &armstandbypool.StandbyVirtualMachinePoolResourceProperties{
	// 			ElasticityProfile: &armstandbypool.StandbyVirtualMachinePoolElasticityProfile{
	// 				MaxReadyCapacity: to.Ptr[int64](304),
	// 				MinReadyCapacity: to.Ptr[int64](300),
	// 			},
	// 			VirtualMachineState: to.Ptr(armstandbypool.VirtualMachineStateRunning),
	// 			AttachedVirtualMachineScaleSetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000009/resourceGroups/rgstandbypool/providers/Microsoft.Compute/virtualMachineScaleSets/myVmss"),
	// 			ProvisioningState: to.Ptr(armstandbypool.ProvisioningStateSucceeded),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 		Location: to.Ptr("West US"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000009/resourceGroups/rgstandbypool/providers/Microsoft.StandbyPool/standbyVirtualMachinePools/pool"),
	// 		Name: to.Ptr("pool"),
	// 		Type: to.Ptr("Microsoft.StandbyPool/standbyVirtualMachinePools"),
	// 		SystemData: &armstandbypool.SystemData{
	// 			CreatedBy: to.Ptr("pooluser@microsoft.com"),
	// 			CreatedByType: to.Ptr(armstandbypool.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-07T16:33:22.210Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("pooluser@microsoft.com"),
	// 			LastModifiedByType: to.Ptr(armstandbypool.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-07T16:33:22.210Z"); return t}()),
	// 		},
	// 	},
	// }
}
