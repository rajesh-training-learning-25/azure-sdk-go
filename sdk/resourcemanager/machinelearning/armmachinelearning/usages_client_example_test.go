//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/aafb0944f7ab936e8cfbad8969bd5eb32263fb4f/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Usage/list.json
func ExampleUsagesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewUsagesClient().NewListPager("eastus", nil)
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
		// page.ListUsagesResult = armmachinelearning.ListUsagesResult{
		// 	Value: []*armmachinelearning.Usage{
		// 		{
		// 			Name: &armmachinelearning.UsageName{
		// 				LocalizedValue: to.Ptr("Clusters"),
		// 				Value: to.Ptr("Clusters"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/totalCores/usages"),
		// 			CurrentValue: to.Ptr[int64](7),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/usages"),
		// 			Limit: to.Ptr[int64](100),
		// 			Unit: to.Ptr(armmachinelearning.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.UsageName{
		// 				LocalizedValue: to.Ptr("Total Cluster Dedicated Regional vCPUs"),
		// 				Value: to.Ptr("Total Cluster Dedicated Regional vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/dedicatedCores/usages"),
		// 			CurrentValue: to.Ptr[int64](14),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/usages"),
		// 			Limit: to.Ptr[int64](24),
		// 			Unit: to.Ptr(armmachinelearning.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.UsageName{
		// 				LocalizedValue: to.Ptr("Standard D Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard D Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/vmFamily/usages"),
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/usages/Standard_D_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](48),
		// 			Unit: to.Ptr(armmachinelearning.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.UsageName{
		// 				LocalizedValue: to.Ptr("Standard DSv2 Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard DSv2 Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/vmFamily/usages"),
		// 			CurrentValue: to.Ptr[int64](2),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/usages/Standard_DSv2_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](24),
		// 			Unit: to.Ptr(armmachinelearning.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.UsageName{
		// 				LocalizedValue: to.Ptr("Standard DSv2 Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard DSv2 Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/usages"),
		// 			CurrentValue: to.Ptr[int64](2),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace1/usages/Standard_DSv2_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](24),
		// 			Unit: to.Ptr(armmachinelearning.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.UsageName{
		// 				LocalizedValue: to.Ptr("Standard DSv2 Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard DSv2 Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/computes/usages"),
		// 			CurrentValue: to.Ptr[int64](2),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace1/computes/demo_cluster1_dsv2/usages/Standard_DSv2_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](24),
		// 			Unit: to.Ptr(armmachinelearning.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.UsageName{
		// 				LocalizedValue: to.Ptr("Standard DSv2 Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard DSv2 Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/computes/usages"),
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace1/computes/demo_cluster2_dsv2/usages/Standard_DSv2_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](24),
		// 			Unit: to.Ptr(armmachinelearning.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.UsageName{
		// 				LocalizedValue: to.Ptr("Standard Dv2 Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard Dv2 Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/vmFamily/usages"),
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/usages/Standard_Dv2_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](24),
		// 			Unit: to.Ptr(armmachinelearning.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.UsageName{
		// 				LocalizedValue: to.Ptr("Standard FSv2 Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard FSv2 Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/vmFamily/usages"),
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/usages/Standard_FSv2_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](24),
		// 			Unit: to.Ptr(armmachinelearning.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.UsageName{
		// 				LocalizedValue: to.Ptr("Standard NC Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard NC Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/vmFamily/usages"),
		// 			CurrentValue: to.Ptr[int64](12),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/usages/Standard_NC_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](24),
		// 			Unit: to.Ptr(armmachinelearning.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.UsageName{
		// 				LocalizedValue: to.Ptr("Standard NC Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard NC Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspace/usages"),
		// 			CurrentValue: to.Ptr[int64](6),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace1/usages/Standard_NC_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](24),
		// 			Unit: to.Ptr(armmachinelearning.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.UsageName{
		// 				LocalizedValue: to.Ptr("Standard NC Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard NC Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspace/computes/usages"),
		// 			CurrentValue: to.Ptr[int64](6),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace1/computes/demo_cluster1_nc/usages/Standard_NC_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](24),
		// 			Unit: to.Ptr(armmachinelearning.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.UsageName{
		// 				LocalizedValue: to.Ptr("Standard NC Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard NC Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/usages"),
		// 			CurrentValue: to.Ptr[int64](6),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace2/usages/Standard_NC_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](24),
		// 			Unit: to.Ptr(armmachinelearning.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.UsageName{
		// 				LocalizedValue: to.Ptr("Standard NC Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard NC Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/computes/usages"),
		// 			CurrentValue: to.Ptr[int64](6),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace2/computes/demo_cluser1_nc/usages/Standard_NC_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](24),
		// 			Unit: to.Ptr(armmachinelearning.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.UsageName{
		// 				LocalizedValue: to.Ptr("Standard NCv2 Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard NCv2 Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/vmFamily/usages"),
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/usages/Standard_NCv2_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](0),
		// 			Unit: to.Ptr(armmachinelearning.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.UsageName{
		// 				LocalizedValue: to.Ptr("Standard NCv3 Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard NCv3 Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/vmFamily/usages"),
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/usages/Standard_NCv3_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](0),
		// 			Unit: to.Ptr(armmachinelearning.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.UsageName{
		// 				LocalizedValue: to.Ptr("Standard ND Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard ND Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/vmFamily/usages"),
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/usages/Standard_ND_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](0),
		// 			Unit: to.Ptr(armmachinelearning.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.UsageName{
		// 				LocalizedValue: to.Ptr("Standard NDv2 Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard NDv2 Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/vmFamily/usages"),
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/usages/Standard_NDv2_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](0),
		// 			Unit: to.Ptr(armmachinelearning.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.UsageName{
		// 				LocalizedValue: to.Ptr("Standard NV Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard NV Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/vmFamily/usages"),
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/usages/Standard_NV_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](24),
		// 			Unit: to.Ptr(armmachinelearning.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.UsageName{
		// 				LocalizedValue: to.Ptr("Total Cluster LowPriority Regional vCPUs"),
		// 				Value: to.Ptr("Total Cluster LowPriority Regional vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/lowPriorityCores/usages"),
		// 			CurrentValue: to.Ptr[int64](18),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/usages"),
		// 			Limit: to.Ptr[int64](50),
		// 			Unit: to.Ptr(armmachinelearning.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.UsageName{
		// 				LocalizedValue: to.Ptr("Standard D Family Cluster LowPriority vCPUs"),
		// 				Value: to.Ptr("Standard D Family Cluster LowPriority vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/vmFamily/usages"),
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/usages/Standard_D_Family_Cluster_LowPriority_vCPUs"),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armmachinelearning.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.UsageName{
		// 				LocalizedValue: to.Ptr("Standard DSv2 Family Cluster LowPriority vCPUs"),
		// 				Value: to.Ptr("Standard DSv2 Family Cluster LowPriority vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/vmFamily/usages"),
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/usages/Standard_DSv2_Family_Cluster_LowPriority_vCPUs"),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armmachinelearning.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.UsageName{
		// 				LocalizedValue: to.Ptr("Standard Dv2 Family Cluster LowPriority vCPUs"),
		// 				Value: to.Ptr("Standard Dv2 Family Cluster LowPriority vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/vmFamily/usages"),
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/usages/Standard_Dv2_Family_Cluster_LowPriority_vCPUs"),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armmachinelearning.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.UsageName{
		// 				LocalizedValue: to.Ptr("Standard FSv2 Family Cluster LowPriority vCPUs"),
		// 				Value: to.Ptr("Standard FSv2 Family Cluster LowPriority vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/vmFamily/usages"),
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/usages/Standard_FSv2_Family_Cluster_LowPriority_vCPUs"),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armmachinelearning.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.UsageName{
		// 				LocalizedValue: to.Ptr("Standard NC Family Cluster LowPriority vCPUs"),
		// 				Value: to.Ptr("Standard NC Family Cluster LowPriority vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/vmFamily/usages"),
		// 			CurrentValue: to.Ptr[int64](18),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/usages/Standard_NC_Family_Cluster_LowPriority_vCPUs"),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armmachinelearning.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.UsageName{
		// 				LocalizedValue: to.Ptr("Standard NC Family Cluster LowPriority vCPUs"),
		// 				Value: to.Ptr("Standard NC Family Cluster LowPriority vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspace/usages"),
		// 			CurrentValue: to.Ptr[int64](6),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace1/usages/Standard_NC_Family_Cluster_LowPriority_vCPUs"),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armmachinelearning.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.UsageName{
		// 				LocalizedValue: to.Ptr("Standard NC Family Cluster LowPriority vCPUs"),
		// 				Value: to.Ptr("Standard NC Family Cluster LowPriority vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspace/computes/usages"),
		// 			CurrentValue: to.Ptr[int64](6),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace1/computes/demo_cluster1_lowPriority_nc/usages/Standard_NC_Family_Cluster_LowPriority_vCPUs"),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armmachinelearning.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.UsageName{
		// 				LocalizedValue: to.Ptr("Standard NC Family Cluster LowPriority vCPUs"),
		// 				Value: to.Ptr("Standard NC Family Cluster LowPriority vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspace/usages"),
		// 			CurrentValue: to.Ptr[int64](12),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace2/usages/Standard_NC_Family_Cluster_LowPriority_vCPUs"),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armmachinelearning.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.UsageName{
		// 				LocalizedValue: to.Ptr("Standard NC Family Cluster LowPriority vCPUs"),
		// 				Value: to.Ptr("Standard NC Family Cluster LowPriority vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspace/computes/usages"),
		// 			CurrentValue: to.Ptr[int64](6),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace2/computes/demo_cluster2_lowPriority_nc/usages/Standard_NC_Family_Cluster_LowPriority_vCPUs"),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armmachinelearning.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.UsageName{
		// 				LocalizedValue: to.Ptr("Standard NC Family Cluster LowPriority vCPUs"),
		// 				Value: to.Ptr("Standard NC Family Cluster LowPriority vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspace/computes/usages"),
		// 			CurrentValue: to.Ptr[int64](6),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace2/computes/demo_cluster3_lowPriority_nc/usages/Standard_NC_Family_Cluster_LowPriority_vCPUs"),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armmachinelearning.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.UsageName{
		// 				LocalizedValue: to.Ptr("Standard NCv2 Family Cluster LowPriority vCPUs"),
		// 				Value: to.Ptr("Standard NCv2 Family Cluster LowPriority vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/vmFamily/usages"),
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/usages/Standard_NCv2_Family_Cluster_LowPriority_vCPUs"),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armmachinelearning.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.UsageName{
		// 				LocalizedValue: to.Ptr("Standard NCv3 Family Cluster LowPriority vCPUs"),
		// 				Value: to.Ptr("Standard NCv3 Family Cluster LowPriority vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/vmFamily/usages"),
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/usages/Standard_NCv3_Family_Cluster_LowPriority_vCPUs"),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armmachinelearning.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.UsageName{
		// 				LocalizedValue: to.Ptr("Standard ND Family Cluster LowPriority vCPUs"),
		// 				Value: to.Ptr("Standard ND Family Cluster LowPriority vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/vmFamily/usages"),
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/usages/Standard_ND_Family_Cluster_LowPriority_vCPUs"),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armmachinelearning.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.UsageName{
		// 				LocalizedValue: to.Ptr("Standard NDv2 Family Cluster LowPriority vCPUs"),
		// 				Value: to.Ptr("Standard NDv2 Family Cluster LowPriority vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/vmFamily/usages"),
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/usages/Standard_NDv2_Family_Cluster_LowPriority_vCPUs"),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armmachinelearning.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.UsageName{
		// 				LocalizedValue: to.Ptr("Standard NV Family Cluster LowPriority vCPUs"),
		// 				Value: to.Ptr("Standard NV Family Cluster LowPriority vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/vmFamily/usages"),
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/usages/Standard_NV_Family_Cluster_LowPriority_vCPUs"),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armmachinelearning.UsageUnitCount),
		// 	}},
		// }
	}
}
