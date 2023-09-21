//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/68d03f91ea7c30e1ab28fb9d35c13f81bc85b724/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/AgentPoolsGet.json
func ExampleAgentPoolsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAgentPoolsClient().Get(ctx, "myResourceGroup", "myRegistry", "myAgentPool", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AgentPool = armcontainerregistry.AgentPool{
	// 	Name: to.Ptr("myAgentPool"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/agentPools"),
	// 	ID: to.Ptr("/subscriptions/f9d7ebed-adbd-4cb4-b973-aaf82c136138/resourceGroups/huanwudfwestgroup/providers/Microsoft.ContainerRegistry/registries/huanglidfwest01/agentPools/testagent26"),
	// 	Location: to.Ptr("WESTUS"),
	// 	Properties: &armcontainerregistry.AgentPoolProperties{
	// 		Count: to.Ptr[int32](1),
	// 		OS: to.Ptr(armcontainerregistry.OSLinux),
	// 		ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		Tier: to.Ptr("S1"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/68d03f91ea7c30e1ab28fb9d35c13f81bc85b724/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/AgentPoolsCreate.json
func ExampleAgentPoolsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAgentPoolsClient().BeginCreate(ctx, "myResourceGroup", "myRegistry", "myAgentPool", armcontainerregistry.AgentPool{
		Location: to.Ptr("WESTUS"),
		Tags: map[string]*string{
			"key": to.Ptr("value"),
		},
		Properties: &armcontainerregistry.AgentPoolProperties{
			Count: to.Ptr[int32](1),
			OS:    to.Ptr(armcontainerregistry.OSLinux),
			Tier:  to.Ptr("S1"),
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
	// res.AgentPool = armcontainerregistry.AgentPool{
	// 	Name: to.Ptr("myAgentPool"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/agentPools"),
	// 	ID: to.Ptr("/subscriptions/f9d7ebed-adbd-4cb4-b973-aaf82c136138/resourceGroups/huanwudfwestgroup/providers/Microsoft.ContainerRegistry/registries/huanglidfwest01/agentPools/testagent26"),
	// 	Location: to.Ptr("WESTUS"),
	// 	Properties: &armcontainerregistry.AgentPoolProperties{
	// 		Count: to.Ptr[int32](1),
	// 		OS: to.Ptr(armcontainerregistry.OSLinux),
	// 		ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		Tier: to.Ptr("S1"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/68d03f91ea7c30e1ab28fb9d35c13f81bc85b724/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/AgentPoolsDelete.json
func ExampleAgentPoolsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAgentPoolsClient().BeginDelete(ctx, "myResourceGroup", "myRegistry", "myAgentPool", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/68d03f91ea7c30e1ab28fb9d35c13f81bc85b724/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/AgentPoolsUpdate.json
func ExampleAgentPoolsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAgentPoolsClient().BeginUpdate(ctx, "myResourceGroup", "myRegistry", "myAgentPool", armcontainerregistry.AgentPoolUpdateParameters{
		Properties: &armcontainerregistry.AgentPoolPropertiesUpdateParameters{
			Count: to.Ptr[int32](1),
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
	// res.AgentPool = armcontainerregistry.AgentPool{
	// 	Name: to.Ptr("myAgentPool"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/agentPools"),
	// 	ID: to.Ptr("/subscriptions/f9d7ebed-adbd-4cb4-b973-aaf82c136138/resourceGroups/huanwudfwestgroup/providers/Microsoft.ContainerRegistry/registries/huanglidfwest01/agentPools/testagent26"),
	// 	Location: to.Ptr("WESTUS"),
	// 	Properties: &armcontainerregistry.AgentPoolProperties{
	// 		Count: to.Ptr[int32](1),
	// 		OS: to.Ptr(armcontainerregistry.OSLinux),
	// 		ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		Tier: to.Ptr("S1"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/68d03f91ea7c30e1ab28fb9d35c13f81bc85b724/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/AgentPoolsList.json
func ExampleAgentPoolsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAgentPoolsClient().NewListPager("myResourceGroup", "myRegistry", nil)
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
		// page.AgentPoolListResult = armcontainerregistry.AgentPoolListResult{
		// 	Value: []*armcontainerregistry.AgentPool{
		// 		{
		// 			Name: to.Ptr("myAgentPool"),
		// 			Type: to.Ptr("Microsoft.ContainerRegistry/registries/agentPools"),
		// 			ID: to.Ptr("/subscriptions/f9d7ebed-adbd-4cb4-b973-aaf82c136138/resourceGroups/huanwudfwestgroup/providers/Microsoft.ContainerRegistry/registries/huanglidfwest01/agentPools/testagent26"),
		// 			Location: to.Ptr("WESTUS"),
		// 			Properties: &armcontainerregistry.AgentPoolProperties{
		// 				Count: to.Ptr[int32](1),
		// 				OS: to.Ptr(armcontainerregistry.OSLinux),
		// 				ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
		// 				Tier: to.Ptr("S1"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/68d03f91ea7c30e1ab28fb9d35c13f81bc85b724/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/AgentPoolsGetQueueStatus.json
func ExampleAgentPoolsClient_GetQueueStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAgentPoolsClient().GetQueueStatus(ctx, "myResourceGroup", "myRegistry", "myAgentPool", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AgentPoolQueueStatus = armcontainerregistry.AgentPoolQueueStatus{
	// 	Count: to.Ptr[int32](10),
	// }
}
