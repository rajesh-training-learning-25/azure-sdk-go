//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armstoragemover_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagemover/armstoragemover"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3d7a3848106b831a4a7f46976fe38aa605c4f44d/specification/storagemover/resource-manager/Microsoft.StorageMover/preview/2022-07-01-preview/examples/Agents_List.json
func ExampleAgentsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstoragemover.NewAgentsClient("11111111-2222-3333-4444-555555555555", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("examples-rg", "examples-storageMoverName", nil)
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
		// page.AgentList = armstoragemover.AgentList{
		// 	Value: []*armstoragemover.Agent{
		// 		{
		// 			Name: to.Ptr("examples-agentName1"),
		// 			Type: to.Ptr("Microsoft.StorageMover/storageMovers/agents"),
		// 			ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/agents/examples-agentName1"),
		// 			Properties: &armstoragemover.AgentProperties{
		// 				Description: to.Ptr("Example Agent 1 Description"),
		// 				AgentStatus: to.Ptr(armstoragemover.AgentStatusOnline),
		// 				AgentVersion: to.Ptr("1.0.0"),
		// 				ArcResourceID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.HybridCompute/machines/examples-hybridComputeName1"),
		// 				ArcVMUUID: to.Ptr("3bb2c024-eba9-4d18-9e7a-1d772fcc5fe9"),
		// 				LastStatusUpdate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.1075056Z"); return t}()),
		// 				LocalIPAddress: to.Ptr("192.168.0.0"),
		// 				MemoryInMB: to.Ptr[int64](100024),
		// 				NumberOfCores: to.Ptr[int64](32),
		// 				UptimeInSeconds: to.Ptr[int64](522),
		// 			},
		// 			SystemData: &armstoragemover.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.1075056Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armstoragemover.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.1075056Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armstoragemover.CreatedByTypeUser),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("examples-agentName2"),
		// 			Type: to.Ptr("Microsoft.StorageMover/storageMovers/agents"),
		// 			ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/agents/examples-agentName2"),
		// 			Properties: &armstoragemover.AgentProperties{
		// 				Description: to.Ptr("Example Agent 2 Description"),
		// 				AgentStatus: to.Ptr(armstoragemover.AgentStatusOnline),
		// 				AgentVersion: to.Ptr("1.0.0"),
		// 				ArcResourceID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.HybridCompute/machines/examples-hybridComputeName2"),
		// 				ArcVMUUID: to.Ptr("147a1f84-7bbf-4e99-9a6a-a1735a91dfd5"),
		// 				LastStatusUpdate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.1075056Z"); return t}()),
		// 				LocalIPAddress: to.Ptr("192.168.0.0"),
		// 				MemoryInMB: to.Ptr[int64](100024),
		// 				NumberOfCores: to.Ptr[int64](32),
		// 				UptimeInSeconds: to.Ptr[int64](522),
		// 			},
		// 			SystemData: &armstoragemover.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.1075056Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armstoragemover.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.1075056Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armstoragemover.CreatedByTypeUser),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("examples-agentName3"),
		// 			Type: to.Ptr("Microsoft.StorageMover/storageMovers/agents"),
		// 			ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/agents/examples-agentName3"),
		// 			Properties: &armstoragemover.AgentProperties{
		// 				Description: to.Ptr("Example Agent 3 Description"),
		// 				AgentStatus: to.Ptr(armstoragemover.AgentStatusOnline),
		// 				AgentVersion: to.Ptr("1.0.0"),
		// 				ArcResourceID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.HybridCompute/machines/examples-hybridComputeName3"),
		// 				ArcVMUUID: to.Ptr("648a7958-f99e-4268-b20e-94c96558dc0d"),
		// 				ErrorDetails: &armstoragemover.AgentPropertiesErrorDetails{
		// 					Code: to.Ptr("SampleErrorCode"),
		// 					Message: to.Ptr("Detailed sample error message."),
		// 				},
		// 				LastStatusUpdate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.1075056Z"); return t}()),
		// 				LocalIPAddress: to.Ptr("192.168.0.0"),
		// 				MemoryInMB: to.Ptr[int64](100024),
		// 				NumberOfCores: to.Ptr[int64](32),
		// 				UptimeInSeconds: to.Ptr[int64](522),
		// 			},
		// 			SystemData: &armstoragemover.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.1075056Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armstoragemover.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.1075056Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armstoragemover.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3d7a3848106b831a4a7f46976fe38aa605c4f44d/specification/storagemover/resource-manager/Microsoft.StorageMover/preview/2022-07-01-preview/examples/Agents_Get.json
func ExampleAgentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstoragemover.NewAgentsClient("11111111-2222-3333-4444-555555555555", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "examples-rg", "examples-storageMoverName", "examples-agentName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Agent = armstoragemover.Agent{
	// 	Name: to.Ptr("examples-agentName"),
	// 	Type: to.Ptr("Microsoft.StorageMover/storageMovers/agents"),
	// 	ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/agents/examples-agentName"),
	// 	Properties: &armstoragemover.AgentProperties{
	// 		Description: to.Ptr("Example Agent Description"),
	// 		AgentStatus: to.Ptr(armstoragemover.AgentStatusOnline),
	// 		AgentVersion: to.Ptr("1.0.0"),
	// 		ArcResourceID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.HybridCompute/machines/examples-hybridComputeName"),
	// 		ArcVMUUID: to.Ptr("3bb2c024-eba9-4d18-9e7a-1d772fcc5fe9"),
	// 		LastStatusUpdate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.1075056Z"); return t}()),
	// 		LocalIPAddress: to.Ptr("192.168.0.0"),
	// 		MemoryInMB: to.Ptr[int64](100024),
	// 		NumberOfCores: to.Ptr[int64](32),
	// 		UptimeInSeconds: to.Ptr[int64](522),
	// 	},
	// 	SystemData: &armstoragemover.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.1075056Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armstoragemover.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.1075056Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armstoragemover.CreatedByTypeUser),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3d7a3848106b831a4a7f46976fe38aa605c4f44d/specification/storagemover/resource-manager/Microsoft.StorageMover/preview/2022-07-01-preview/examples/Agents_CreateOrUpdate.json
func ExampleAgentsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstoragemover.NewAgentsClient("11111111-2222-3333-4444-555555555555", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "examples-rg", "examples-storageMoverName", "examples-agentName", armstoragemover.Agent{
		Properties: &armstoragemover.AgentProperties{
			Description:   to.Ptr("Example Agent Description"),
			ArcResourceID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.HybridCompute/machines/examples-hybridComputeName"),
			ArcVMUUID:     to.Ptr("3bb2c024-eba9-4d18-9e7a-1d772fcc5fe9"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Agent = armstoragemover.Agent{
	// 	Name: to.Ptr("examples-agentName"),
	// 	Type: to.Ptr("Microsoft.StorageMover/storageMovers/agents"),
	// 	ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/agents/examples-agentName"),
	// 	Properties: &armstoragemover.AgentProperties{
	// 		Description: to.Ptr("Example Agent Description"),
	// 		AgentStatus: to.Ptr(armstoragemover.AgentStatusRegistering),
	// 		AgentVersion: to.Ptr("1.0.0"),
	// 		ArcResourceID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.HybridCompute/machines/examples-hybridComputeName"),
	// 		ArcVMUUID: to.Ptr("3bb2c024-eba9-4d18-9e7a-1d772fcc5fe9"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3d7a3848106b831a4a7f46976fe38aa605c4f44d/specification/storagemover/resource-manager/Microsoft.StorageMover/preview/2022-07-01-preview/examples/Agents_Update.json
func ExampleAgentsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstoragemover.NewAgentsClient("11111111-2222-3333-4444-555555555555", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx, "examples-rg", "examples-storageMoverName", "examples-agentName", armstoragemover.AgentUpdateParameters{
		Properties: &armstoragemover.AgentUpdateProperties{
			Description: to.Ptr("Updated Agent Description"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Agent = armstoragemover.Agent{
	// 	Name: to.Ptr("examples-agentName"),
	// 	Type: to.Ptr("Microsoft.StorageMover/storageMovers/agents"),
	// 	ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/agents/examples-agentName"),
	// 	Properties: &armstoragemover.AgentProperties{
	// 		Description: to.Ptr("Updated Agent Description"),
	// 		AgentStatus: to.Ptr(armstoragemover.AgentStatusOnline),
	// 		AgentVersion: to.Ptr("1.0.0"),
	// 		ArcResourceID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.HybridCompute/machines/examples-hybridComputeName"),
	// 		ArcVMUUID: to.Ptr("3bb2c024-eba9-4d18-9e7a-1d772fcc5fe9"),
	// 		LastStatusUpdate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.1075056Z"); return t}()),
	// 		LocalIPAddress: to.Ptr("192.168.0.0"),
	// 		MemoryInMB: to.Ptr[int64](100024),
	// 		NumberOfCores: to.Ptr[int64](32),
	// 		UptimeInSeconds: to.Ptr[int64](522),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3d7a3848106b831a4a7f46976fe38aa605c4f44d/specification/storagemover/resource-manager/Microsoft.StorageMover/preview/2022-07-01-preview/examples/Agents_Delete.json
func ExampleAgentsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstoragemover.NewAgentsClient("11111111-2222-3333-4444-555555555555", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx, "examples-rg", "examples-storageMoverName", "examples-agentName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
