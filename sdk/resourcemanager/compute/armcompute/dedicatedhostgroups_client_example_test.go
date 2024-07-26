//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/dedicatedHostExamples/DedicatedHostGroup_CreateOrUpdate_WithUltraSSD.json
func ExampleDedicatedHostGroupsClient_CreateOrUpdate_createOrUpdateADedicatedHostGroupWithUltraSsdSupport() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDedicatedHostGroupsClient().CreateOrUpdate(ctx, "myResourceGroup", "myDedicatedHostGroup", armcompute.DedicatedHostGroup{
		Location: to.Ptr("westus"),
		Tags: map[string]*string{
			"department": to.Ptr("finance"),
		},
		Properties: &armcompute.DedicatedHostGroupProperties{
			AdditionalCapabilities: &armcompute.DedicatedHostGroupPropertiesAdditionalCapabilities{
				UltraSSDEnabled: to.Ptr(true),
			},
			PlatformFaultDomainCount:  to.Ptr[int32](3),
			SupportAutomaticPlacement: to.Ptr(true),
		},
		Zones: []*string{
			to.Ptr("1")},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DedicatedHostGroup = armcompute.DedicatedHostGroup{
	// 	Name: to.Ptr("myDedicatedHostGroup"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/HostGroups/myDedicatedHostGroup"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"department": to.Ptr("finance"),
	// 		"owner": to.Ptr("myCompany"),
	// 	},
	// 	Properties: &armcompute.DedicatedHostGroupProperties{
	// 		AdditionalCapabilities: &armcompute.DedicatedHostGroupPropertiesAdditionalCapabilities{
	// 			UltraSSDEnabled: to.Ptr(true),
	// 		},
	// 		PlatformFaultDomainCount: to.Ptr[int32](3),
	// 		SupportAutomaticPlacement: to.Ptr(true),
	// 	},
	// 	Zones: []*string{
	// 		to.Ptr("1")},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/dedicatedHostExamples/DedicatedHostGroup_CreateOrUpdate.json
func ExampleDedicatedHostGroupsClient_CreateOrUpdate_createOrUpdateADedicatedHostGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDedicatedHostGroupsClient().CreateOrUpdate(ctx, "myResourceGroup", "myDedicatedHostGroup", armcompute.DedicatedHostGroup{
		Location: to.Ptr("westus"),
		Tags: map[string]*string{
			"department": to.Ptr("finance"),
		},
		Properties: &armcompute.DedicatedHostGroupProperties{
			PlatformFaultDomainCount:  to.Ptr[int32](3),
			SupportAutomaticPlacement: to.Ptr(true),
		},
		Zones: []*string{
			to.Ptr("1")},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DedicatedHostGroup = armcompute.DedicatedHostGroup{
	// 	Name: to.Ptr("myDedicatedHostGroup"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/HostGroups/myDedicatedHostGroup"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"department": to.Ptr("finance"),
	// 		"owner": to.Ptr("myCompany"),
	// 	},
	// 	Properties: &armcompute.DedicatedHostGroupProperties{
	// 		PlatformFaultDomainCount: to.Ptr[int32](3),
	// 		SupportAutomaticPlacement: to.Ptr(true),
	// 	},
	// 	Zones: []*string{
	// 		to.Ptr("1")},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/dedicatedHostExamples/DedicatedHostGroup_Update_MaximumSet_Gen.json
func ExampleDedicatedHostGroupsClient_Update_dedicatedHostGroupUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDedicatedHostGroupsClient().Update(ctx, "rgcompute", "aaaa", armcompute.DedicatedHostGroupUpdate{
		Tags: map[string]*string{
			"key9921": to.Ptr("aaaaaaaaaa"),
		},
		Properties: &armcompute.DedicatedHostGroupProperties{
			InstanceView: &armcompute.DedicatedHostGroupInstanceView{
				Hosts: []*armcompute.DedicatedHostInstanceViewWithName{
					{
						AvailableCapacity: &armcompute.DedicatedHostAvailableCapacity{
							AllocatableVMs: []*armcompute.DedicatedHostAllocatableVM{
								{
									Count:  to.Ptr[float64](26),
									VMSize: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
								}},
						},
						Statuses: []*armcompute.InstanceViewStatus{
							{
								Code:          to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
								DisplayStatus: to.Ptr("aaaaaa"),
								Level:         to.Ptr(armcompute.StatusLevelTypesInfo),
								Message:       to.Ptr("a"),
								Time:          to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t }()),
							}},
					}},
			},
			PlatformFaultDomainCount:  to.Ptr[int32](3),
			SupportAutomaticPlacement: to.Ptr(true),
		},
		Zones: []*string{
			to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa")},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DedicatedHostGroup = armcompute.DedicatedHostGroup{
	// 	Name: to.Ptr("myDedicatedHostGroup"),
	// 	Type: to.Ptr("aaaa"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/HostGroups/myDedicatedHostGroup"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armcompute.DedicatedHostGroupProperties{
	// 		Hosts: []*armcompute.SubResourceReadOnly{
	// 			{
	// 				ID: to.Ptr("aaaa"),
	// 		}},
	// 		InstanceView: &armcompute.DedicatedHostGroupInstanceView{
	// 			Hosts: []*armcompute.DedicatedHostInstanceViewWithName{
	// 				{
	// 					AssetID: to.Ptr("aaaa"),
	// 					AvailableCapacity: &armcompute.DedicatedHostAvailableCapacity{
	// 						AllocatableVMs: []*armcompute.DedicatedHostAllocatableVM{
	// 							{
	// 								Count: to.Ptr[float64](26),
	// 								VMSize: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
	// 						}},
	// 					},
	// 					Statuses: []*armcompute.InstanceViewStatus{
	// 						{
	// 							Code: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
	// 							DisplayStatus: to.Ptr("aaaaaa"),
	// 							Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 							Message: to.Ptr("a"),
	// 							Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t}()),
	// 					}},
	// 					Name: to.Ptr("aaaaaaaaaaaaaaaaaa"),
	// 			}},
	// 		},
	// 		PlatformFaultDomainCount: to.Ptr[int32](3),
	// 		SupportAutomaticPlacement: to.Ptr(true),
	// 	},
	// 	Zones: []*string{
	// 		to.Ptr("1")},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/dedicatedHostExamples/DedicatedHostGroup_Update_MinimumSet_Gen.json
func ExampleDedicatedHostGroupsClient_Update_dedicatedHostGroupUpdateMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDedicatedHostGroupsClient().Update(ctx, "rgcompute", "aaaaaaaaaaa", armcompute.DedicatedHostGroupUpdate{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DedicatedHostGroup = armcompute.DedicatedHostGroup{
	// 	Location: to.Ptr("westus"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/dedicatedHostExamples/DedicatedHostGroup_Delete_MaximumSet_Gen.json
func ExampleDedicatedHostGroupsClient_Delete_dedicatedHostGroupDeleteMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewDedicatedHostGroupsClient().Delete(ctx, "rgcompute", "a", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/dedicatedHostExamples/DedicatedHostGroup_Delete_MinimumSet_Gen.json
func ExampleDedicatedHostGroupsClient_Delete_dedicatedHostGroupDeleteMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewDedicatedHostGroupsClient().Delete(ctx, "rgcompute", "aaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/dedicatedHostExamples/DedicatedHostGroup_Get.json
func ExampleDedicatedHostGroupsClient_Get_createADedicatedHostGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDedicatedHostGroupsClient().Get(ctx, "myResourceGroup", "myDedicatedHostGroup", &armcompute.DedicatedHostGroupsClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DedicatedHostGroup = armcompute.DedicatedHostGroup{
	// 	Name: to.Ptr("myDedicatedHostGroup"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/HostGroups/myDedicatedHostGroup"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"{tagName}": to.Ptr("{tagValue}"),
	// 	},
	// 	Properties: &armcompute.DedicatedHostGroupProperties{
	// 		Hosts: []*armcompute.SubResourceReadOnly{
	// 			{
	// 				ID: to.Ptr("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/myDedicatedHostGroup/myHostGroup/Hosts/myHost1"),
	// 			},
	// 			{
	// 				ID: to.Ptr("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/myDedicatedHostGroup/myHostGroup/Hosts/myHost2"),
	// 		}},
	// 		InstanceView: &armcompute.DedicatedHostGroupInstanceView{
	// 			Hosts: []*armcompute.DedicatedHostInstanceViewWithName{
	// 				{
	// 					AssetID: to.Ptr("eb3f58b8-b4e8-4882-b69f-301a01812407"),
	// 					AvailableCapacity: &armcompute.DedicatedHostAvailableCapacity{
	// 						AllocatableVMs: []*armcompute.DedicatedHostAllocatableVM{
	// 							{
	// 								Count: to.Ptr[float64](10),
	// 								VMSize: to.Ptr("Standard_A1"),
	// 						}},
	// 					},
	// 					Statuses: []*armcompute.InstanceViewStatus{
	// 						{
	// 							Code: to.Ptr("ProvisioningState/succeeded"),
	// 							DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 							Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 						},
	// 						{
	// 							Code: to.Ptr("HealthState/available"),
	// 							DisplayStatus: to.Ptr("Host available"),
	// 							Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 					}},
	// 					Name: to.Ptr("myHost1"),
	// 				},
	// 				{
	// 					AssetID: to.Ptr("f293d4ac-5eea-4be2-b0c0-0fcaa09aebf8"),
	// 					AvailableCapacity: &armcompute.DedicatedHostAvailableCapacity{
	// 						AllocatableVMs: []*armcompute.DedicatedHostAllocatableVM{
	// 							{
	// 								Count: to.Ptr[float64](10),
	// 								VMSize: to.Ptr("Standard_A1"),
	// 						}},
	// 					},
	// 					Statuses: []*armcompute.InstanceViewStatus{
	// 						{
	// 							Code: to.Ptr("ProvisioningState/succeeded"),
	// 							DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 							Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 						},
	// 						{
	// 							Code: to.Ptr("HealthState/available"),
	// 							DisplayStatus: to.Ptr("Host available"),
	// 							Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 					}},
	// 					Name: to.Ptr("myHost2"),
	// 			}},
	// 		},
	// 		PlatformFaultDomainCount: to.Ptr[int32](3),
	// 		SupportAutomaticPlacement: to.Ptr(true),
	// 	},
	// 	Zones: []*string{
	// 		to.Ptr("3")},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/dedicatedHostExamples/DedicatedHostGroup_Get_UltraSSDEnabledDedicatedHostGroup.json
func ExampleDedicatedHostGroupsClient_Get_createAnUltraSsdEnabledDedicatedHostGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDedicatedHostGroupsClient().Get(ctx, "myResourceGroup", "myDedicatedHostGroup", &armcompute.DedicatedHostGroupsClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DedicatedHostGroup = armcompute.DedicatedHostGroup{
	// 	Name: to.Ptr("myDedicatedHostGroup"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/HostGroups/myDedicatedHostGroup"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"{tagName}": to.Ptr("{tagValue}"),
	// 	},
	// 	Properties: &armcompute.DedicatedHostGroupProperties{
	// 		AdditionalCapabilities: &armcompute.DedicatedHostGroupPropertiesAdditionalCapabilities{
	// 			UltraSSDEnabled: to.Ptr(true),
	// 		},
	// 		Hosts: []*armcompute.SubResourceReadOnly{
	// 			{
	// 				ID: to.Ptr("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/myDedicatedHostGroup/myHostGroup/Hosts/myHost"),
	// 		}},
	// 		InstanceView: &armcompute.DedicatedHostGroupInstanceView{
	// 			Hosts: []*armcompute.DedicatedHostInstanceViewWithName{
	// 				{
	// 					AssetID: to.Ptr("eb3f58b8-b4e8-4882-b69f-301a01812407"),
	// 					AvailableCapacity: &armcompute.DedicatedHostAvailableCapacity{
	// 						AllocatableVMs: []*armcompute.DedicatedHostAllocatableVM{
	// 							{
	// 								Count: to.Ptr[float64](10),
	// 								VMSize: to.Ptr("Standard_A1"),
	// 						}},
	// 					},
	// 					Statuses: []*armcompute.InstanceViewStatus{
	// 						{
	// 							Code: to.Ptr("ProvisioningState/succeeded"),
	// 							DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 							Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 						},
	// 						{
	// 							Code: to.Ptr("HealthState/available"),
	// 							DisplayStatus: to.Ptr("Host available"),
	// 							Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 					}},
	// 					Name: to.Ptr("myHost"),
	// 			}},
	// 		},
	// 		PlatformFaultDomainCount: to.Ptr[int32](3),
	// 		SupportAutomaticPlacement: to.Ptr(true),
	// 	},
	// 	Zones: []*string{
	// 		to.Ptr("3")},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/dedicatedHostExamples/DedicatedHostGroup_ListByResourceGroup_MaximumSet_Gen.json
func ExampleDedicatedHostGroupsClient_NewListByResourceGroupPager_dedicatedHostGroupListByResourceGroupMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDedicatedHostGroupsClient().NewListByResourceGroupPager("rgcompute", nil)
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
		// page.DedicatedHostGroupListResult = armcompute.DedicatedHostGroupListResult{
		// 	Value: []*armcompute.DedicatedHostGroup{
		// 		{
		// 			Name: to.Ptr("myDedicatedHostGroup"),
		// 			Type: to.Ptr("aaaa"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/HostGroups/myDedicatedHostGroup"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armcompute.DedicatedHostGroupProperties{
		// 				Hosts: []*armcompute.SubResourceReadOnly{
		// 					{
		// 						ID: to.Ptr("aaaa"),
		// 				}},
		// 				InstanceView: &armcompute.DedicatedHostGroupInstanceView{
		// 					Hosts: []*armcompute.DedicatedHostInstanceViewWithName{
		// 						{
		// 							AssetID: to.Ptr("aaaa"),
		// 							AvailableCapacity: &armcompute.DedicatedHostAvailableCapacity{
		// 								AllocatableVMs: []*armcompute.DedicatedHostAllocatableVM{
		// 									{
		// 										Count: to.Ptr[float64](26),
		// 										VMSize: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
		// 								}},
		// 							},
		// 							Statuses: []*armcompute.InstanceViewStatus{
		// 								{
		// 									Code: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
		// 									DisplayStatus: to.Ptr("aaaaaa"),
		// 									Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 									Message: to.Ptr("a"),
		// 									Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t}()),
		// 							}},
		// 							Name: to.Ptr("aaaaaaaaaaaaaaaaaa"),
		// 					}},
		// 				},
		// 				PlatformFaultDomainCount: to.Ptr[int32](3),
		// 				SupportAutomaticPlacement: to.Ptr(true),
		// 			},
		// 			Zones: []*string{
		// 				to.Ptr("1")},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/dedicatedHostExamples/DedicatedHostGroup_ListByResourceGroup_MinimumSet_Gen.json
func ExampleDedicatedHostGroupsClient_NewListByResourceGroupPager_dedicatedHostGroupListByResourceGroupMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDedicatedHostGroupsClient().NewListByResourceGroupPager("rgcompute", nil)
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
		// page.DedicatedHostGroupListResult = armcompute.DedicatedHostGroupListResult{
		// 	Value: []*armcompute.DedicatedHostGroup{
		// 		{
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/rgcompute/providers/Microsoft.Compute/HostGroups/myDedicatedHostGroup"),
		// 			Location: to.Ptr("westus"),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/dedicatedHostExamples/DedicatedHostGroup_ListBySubscription_MaximumSet_Gen.json
func ExampleDedicatedHostGroupsClient_NewListBySubscriptionPager_dedicatedHostGroupListBySubscriptionMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDedicatedHostGroupsClient().NewListBySubscriptionPager(nil)
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
		// page.DedicatedHostGroupListResult = armcompute.DedicatedHostGroupListResult{
		// 	Value: []*armcompute.DedicatedHostGroup{
		// 		{
		// 			Name: to.Ptr("myDedicatedHostGroup"),
		// 			Type: to.Ptr("aaaa"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/HostGroups/myDedicatedHostGroup"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armcompute.DedicatedHostGroupProperties{
		// 				Hosts: []*armcompute.SubResourceReadOnly{
		// 					{
		// 						ID: to.Ptr("aaaa"),
		// 				}},
		// 				InstanceView: &armcompute.DedicatedHostGroupInstanceView{
		// 					Hosts: []*armcompute.DedicatedHostInstanceViewWithName{
		// 						{
		// 							AssetID: to.Ptr("aaaa"),
		// 							AvailableCapacity: &armcompute.DedicatedHostAvailableCapacity{
		// 								AllocatableVMs: []*armcompute.DedicatedHostAllocatableVM{
		// 									{
		// 										Count: to.Ptr[float64](26),
		// 										VMSize: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
		// 								}},
		// 							},
		// 							Statuses: []*armcompute.InstanceViewStatus{
		// 								{
		// 									Code: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
		// 									DisplayStatus: to.Ptr("aaaaaa"),
		// 									Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 									Message: to.Ptr("a"),
		// 									Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t}()),
		// 							}},
		// 							Name: to.Ptr("aaaaaaaaaaaaaaaaaa"),
		// 					}},
		// 				},
		// 				PlatformFaultDomainCount: to.Ptr[int32](3),
		// 				SupportAutomaticPlacement: to.Ptr(true),
		// 			},
		// 			Zones: []*string{
		// 				to.Ptr("1")},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/dedicatedHostExamples/DedicatedHostGroup_ListBySubscription_MinimumSet_Gen.json
func ExampleDedicatedHostGroupsClient_NewListBySubscriptionPager_dedicatedHostGroupListBySubscriptionMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDedicatedHostGroupsClient().NewListBySubscriptionPager(nil)
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
		// page.DedicatedHostGroupListResult = armcompute.DedicatedHostGroupListResult{
		// 	Value: []*armcompute.DedicatedHostGroup{
		// 		{
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/rgcompute/providers/Microsoft.Compute/HostGroups/myDedicatedHostGroup"),
		// 			Location: to.Ptr("westus"),
		// 	}},
		// }
	}
}
