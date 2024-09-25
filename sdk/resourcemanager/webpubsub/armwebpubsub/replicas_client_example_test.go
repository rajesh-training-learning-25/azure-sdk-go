//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armwebpubsub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/webpubsub/armwebpubsub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7189fb57f69468c56df76f9a4d68dd9ff04ab100/specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2023-08-01-preview/examples/WebPubSubReplicas_List.json
func ExampleReplicasClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armwebpubsub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReplicasClient().NewListPager("myResourceGroup", "myWebPubSubService", nil)
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
		// page.ReplicaList = armwebpubsub.ReplicaList{
		// 	Value: []*armwebpubsub.Replica{
		// 		{
		// 			Name: to.Ptr("myWebPubSubService-eastus"),
		// 			Type: to.Ptr("Microsoft.SignalRService/WebPubSub"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.SignalRService/WebPubSub/myWebPubSubService/replicas/myWebPubSubService-eastus"),
		// 			SystemData: &armwebpubsub.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-02-03T04:05:06.000Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armwebpubsub.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-02-03T04:05:06.000Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armwebpubsub.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 			},
		// 			Properties: &armwebpubsub.ReplicaProperties{
		// 				ProvisioningState: to.Ptr(armwebpubsub.ProvisioningStateSucceeded),
		// 				RegionEndpointEnabled: to.Ptr("Enabled"),
		// 				ResourceStopped: to.Ptr("false"),
		// 			},
		// 			SKU: &armwebpubsub.ResourceSKU{
		// 				Name: to.Ptr("Premium_P1"),
		// 				Capacity: to.Ptr[int32](1),
		// 				Size: to.Ptr("P1"),
		// 				Tier: to.Ptr(armwebpubsub.WebPubSubSKUTierPremium),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7189fb57f69468c56df76f9a4d68dd9ff04ab100/specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2023-08-01-preview/examples/WebPubSubReplicas_Get.json
func ExampleReplicasClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armwebpubsub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReplicasClient().Get(ctx, "myResourceGroup", "myWebPubSubService", "myWebPubSubService-eastus", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Replica = armwebpubsub.Replica{
	// 	Name: to.Ptr("myWebPubSubService-eastus"),
	// 	Type: to.Ptr("Microsoft.SignalRService/WebPubSub"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.SignalRService/WebPubSub/myWebPubSubService/replicas/myWebPubSubService-eastus"),
	// 	SystemData: &armwebpubsub.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-02-03T04:05:06.000Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armwebpubsub.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-02-03T04:05:06.000Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armwebpubsub.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 	},
	// 	Properties: &armwebpubsub.ReplicaProperties{
	// 		ProvisioningState: to.Ptr(armwebpubsub.ProvisioningStateSucceeded),
	// 		RegionEndpointEnabled: to.Ptr("Enabled"),
	// 		ResourceStopped: to.Ptr("false"),
	// 	},
	// 	SKU: &armwebpubsub.ResourceSKU{
	// 		Name: to.Ptr("Premium_P1"),
	// 		Capacity: to.Ptr[int32](1),
	// 		Size: to.Ptr("P1"),
	// 		Tier: to.Ptr(armwebpubsub.WebPubSubSKUTierPremium),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7189fb57f69468c56df76f9a4d68dd9ff04ab100/specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2023-08-01-preview/examples/WebPubSubReplicas_CreateOrUpdate.json
func ExampleReplicasClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armwebpubsub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicasClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myWebPubSubService", "myWebPubSubService-eastus", armwebpubsub.Replica{
		Location: to.Ptr("eastus"),
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
		},
		Properties: &armwebpubsub.ReplicaProperties{
			ResourceStopped: to.Ptr("false"),
		},
		SKU: &armwebpubsub.ResourceSKU{
			Name:     to.Ptr("Premium_P1"),
			Capacity: to.Ptr[int32](1),
			Tier:     to.Ptr(armwebpubsub.WebPubSubSKUTierPremium),
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
	// res.Replica = armwebpubsub.Replica{
	// 	Name: to.Ptr("myWebPubSubService-eastus"),
	// 	Type: to.Ptr("Microsoft.SignalRService/WebPubSub"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.SignalRService/WebPubSub/myWebPubSubService/replicas/myWebPubSubService-eastus"),
	// 	SystemData: &armwebpubsub.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-02-03T04:05:06.000Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armwebpubsub.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-02-03T04:05:06.000Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armwebpubsub.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 	},
	// 	Properties: &armwebpubsub.ReplicaProperties{
	// 		ProvisioningState: to.Ptr(armwebpubsub.ProvisioningStateSucceeded),
	// 		RegionEndpointEnabled: to.Ptr("Enabled"),
	// 		ResourceStopped: to.Ptr("false"),
	// 	},
	// 	SKU: &armwebpubsub.ResourceSKU{
	// 		Name: to.Ptr("Premium_P1"),
	// 		Capacity: to.Ptr[int32](1),
	// 		Size: to.Ptr("P1"),
	// 		Tier: to.Ptr(armwebpubsub.WebPubSubSKUTierPremium),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7189fb57f69468c56df76f9a4d68dd9ff04ab100/specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2023-08-01-preview/examples/WebPubSubReplicas_Delete.json
func ExampleReplicasClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armwebpubsub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewReplicasClient().Delete(ctx, "myResourceGroup", "myWebPubSubService", "myWebPubSubService-eastus", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7189fb57f69468c56df76f9a4d68dd9ff04ab100/specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2023-08-01-preview/examples/WebPubSubReplicas_Update.json
func ExampleReplicasClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armwebpubsub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicasClient().BeginUpdate(ctx, "myResourceGroup", "myWebPubSubService", "myWebPubSubService-eastus", armwebpubsub.Replica{
		Location: to.Ptr("eastus"),
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
		},
		Properties: &armwebpubsub.ReplicaProperties{
			ResourceStopped: to.Ptr("false"),
		},
		SKU: &armwebpubsub.ResourceSKU{
			Name:     to.Ptr("Premium_P1"),
			Capacity: to.Ptr[int32](1),
			Tier:     to.Ptr(armwebpubsub.WebPubSubSKUTierPremium),
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
	// res.Replica = armwebpubsub.Replica{
	// 	Name: to.Ptr("myWebPubSubService-eastus"),
	// 	Type: to.Ptr("Microsoft.SignalRService/WebPubSub"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.SignalRService/WebPubSub/myWebPubSubService/replicas/myWebPubSubService-eastus"),
	// 	SystemData: &armwebpubsub.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-02-03T04:05:06.000Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armwebpubsub.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-02-03T04:05:06.000Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armwebpubsub.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 	},
	// 	Properties: &armwebpubsub.ReplicaProperties{
	// 		ProvisioningState: to.Ptr(armwebpubsub.ProvisioningStateSucceeded),
	// 		RegionEndpointEnabled: to.Ptr("Enabled"),
	// 		ResourceStopped: to.Ptr("false"),
	// 	},
	// 	SKU: &armwebpubsub.ResourceSKU{
	// 		Name: to.Ptr("Premium_P1"),
	// 		Capacity: to.Ptr[int32](1),
	// 		Size: to.Ptr("P1"),
	// 		Tier: to.Ptr(armwebpubsub.WebPubSubSKUTierPremium),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7189fb57f69468c56df76f9a4d68dd9ff04ab100/specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2023-08-01-preview/examples/WebPubSubReplicas_Restart.json
func ExampleReplicasClient_BeginRestart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armwebpubsub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicasClient().BeginRestart(ctx, "myResourceGroup", "myWebPubSubService", "myWebPubSubService-eastus", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
