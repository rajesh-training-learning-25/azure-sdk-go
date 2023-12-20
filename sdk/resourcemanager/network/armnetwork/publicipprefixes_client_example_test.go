//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/639ecfad68419328658bd4cfe7094af4ce472be2/specification/network/resource-manager/Microsoft.Network/stable/2023-06-01/examples/PublicIpPrefixDelete.json
func ExamplePublicIPPrefixesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPublicIPPrefixesClient().BeginDelete(ctx, "rg1", "test-ipprefix", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/639ecfad68419328658bd4cfe7094af4ce472be2/specification/network/resource-manager/Microsoft.Network/stable/2023-06-01/examples/PublicIpPrefixGet.json
func ExamplePublicIPPrefixesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPublicIPPrefixesClient().Get(ctx, "rg1", "test-ipprefix", &armnetwork.PublicIPPrefixesClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PublicIPPrefix = armnetwork.PublicIPPrefix{
	// 	Name: to.Ptr("test-ipprefix"),
	// 	Type: to.Ptr("Microsoft.Network/publicIPPrefixes"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPPrefixes/test-ipprefix"),
	// 	Location: to.Ptr("westus"),
	// 	Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
	// 	Properties: &armnetwork.PublicIPPrefixPropertiesFormat{
	// 		IPPrefix: to.Ptr("192.168.254.2/30"),
	// 		IPTags: []*armnetwork.IPTag{
	// 		},
	// 		PrefixLength: to.Ptr[int32](30),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		PublicIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
	// 		PublicIPAddresses: []*armnetwork.ReferencedPublicIPAddress{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/PublicIpAddress1"),
	// 		}},
	// 		ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
	// 	},
	// 	SKU: &armnetwork.PublicIPPrefixSKU{
	// 		Name: to.Ptr(armnetwork.PublicIPPrefixSKUNameStandard),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/639ecfad68419328658bd4cfe7094af4ce472be2/specification/network/resource-manager/Microsoft.Network/stable/2023-06-01/examples/PublicIpPrefixCreateCustomizedValues.json
func ExamplePublicIPPrefixesClient_BeginCreateOrUpdate_createPublicIpPrefixAllocationMethod() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPublicIPPrefixesClient().BeginCreateOrUpdate(ctx, "rg1", "test-ipprefix", armnetwork.PublicIPPrefix{
		Location: to.Ptr("westus"),
		Properties: &armnetwork.PublicIPPrefixPropertiesFormat{
			PrefixLength:           to.Ptr[int32](30),
			PublicIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
		},
		SKU: &armnetwork.PublicIPPrefixSKU{
			Name: to.Ptr(armnetwork.PublicIPPrefixSKUNameStandard),
			Tier: to.Ptr(armnetwork.PublicIPPrefixSKUTierRegional),
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
	// res.PublicIPPrefix = armnetwork.PublicIPPrefix{
	// 	Name: to.Ptr("test-ipprefix"),
	// 	Type: to.Ptr("Microsoft.Network/publicIPPrefixes"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPPrefixes/test-ipprefix"),
	// 	Location: to.Ptr("westus"),
	// 	Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
	// 	Properties: &armnetwork.PublicIPPrefixPropertiesFormat{
	// 		IPPrefix: to.Ptr("192.168.254.2/30"),
	// 		IPTags: []*armnetwork.IPTag{
	// 		},
	// 		PrefixLength: to.Ptr[int32](30),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		PublicIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
	// 		ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
	// 	},
	// 	SKU: &armnetwork.PublicIPPrefixSKU{
	// 		Name: to.Ptr(armnetwork.PublicIPPrefixSKUNameStandard),
	// 		Tier: to.Ptr(armnetwork.PublicIPPrefixSKUTierRegional),
	// 	},
	// 	Zones: []*string{
	// 		to.Ptr("1")},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/639ecfad68419328658bd4cfe7094af4ce472be2/specification/network/resource-manager/Microsoft.Network/stable/2023-06-01/examples/PublicIpPrefixCreateDefaults.json
func ExamplePublicIPPrefixesClient_BeginCreateOrUpdate_createPublicIpPrefixDefaults() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPublicIPPrefixesClient().BeginCreateOrUpdate(ctx, "rg1", "test-ipprefix", armnetwork.PublicIPPrefix{
		Location: to.Ptr("westus"),
		Properties: &armnetwork.PublicIPPrefixPropertiesFormat{
			PrefixLength: to.Ptr[int32](30),
		},
		SKU: &armnetwork.PublicIPPrefixSKU{
			Name: to.Ptr(armnetwork.PublicIPPrefixSKUNameStandard),
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
	// res.PublicIPPrefix = armnetwork.PublicIPPrefix{
	// 	Name: to.Ptr("test-ipprefix"),
	// 	Type: to.Ptr("Microsoft.Network/publicIPPrefixes"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPPrefixes/test-ipprefix"),
	// 	Location: to.Ptr("westus"),
	// 	Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
	// 	Properties: &armnetwork.PublicIPPrefixPropertiesFormat{
	// 		IPPrefix: to.Ptr("192.168.254.2/30"),
	// 		IPTags: []*armnetwork.IPTag{
	// 		},
	// 		PrefixLength: to.Ptr[int32](30),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		PublicIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
	// 		ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
	// 	},
	// 	SKU: &armnetwork.PublicIPPrefixSKU{
	// 		Name: to.Ptr(armnetwork.PublicIPPrefixSKUNameStandard),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/639ecfad68419328658bd4cfe7094af4ce472be2/specification/network/resource-manager/Microsoft.Network/stable/2023-06-01/examples/PublicIpPrefixUpdateTags.json
func ExamplePublicIPPrefixesClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPublicIPPrefixesClient().UpdateTags(ctx, "rg1", "test-ipprefix", armnetwork.TagsObject{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PublicIPPrefix = armnetwork.PublicIPPrefix{
	// 	Name: to.Ptr("test-ipprefix"),
	// 	Type: to.Ptr("Microsoft.Network/publicIPPrefixes"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPPrefixes/test-ipprefix"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
	// 	Properties: &armnetwork.PublicIPPrefixPropertiesFormat{
	// 		IPPrefix: to.Ptr("40.85.154.247/30"),
	// 		IPTags: []*armnetwork.IPTag{
	// 		},
	// 		PrefixLength: to.Ptr[int32](30),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		PublicIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
	// 		ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
	// 	},
	// 	SKU: &armnetwork.PublicIPPrefixSKU{
	// 		Name: to.Ptr(armnetwork.PublicIPPrefixSKUNameStandard),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/639ecfad68419328658bd4cfe7094af4ce472be2/specification/network/resource-manager/Microsoft.Network/stable/2023-06-01/examples/PublicIpPrefixListAll.json
func ExamplePublicIPPrefixesClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPublicIPPrefixesClient().NewListAllPager(nil)
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
		// page.PublicIPPrefixListResult = armnetwork.PublicIPPrefixListResult{
		// 	Value: []*armnetwork.PublicIPPrefix{
		// 		{
		// 			Name: to.Ptr("test-ipprefix"),
		// 			Type: to.Ptr("Microsoft.Network/publicIPPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPPrefixes/test-ipprefix"),
		// 			Location: to.Ptr("westus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.PublicIPPrefixPropertiesFormat{
		// 				IPPrefix: to.Ptr("41.85.154.247/30"),
		// 				IPTags: []*armnetwork.IPTag{
		// 				},
		// 				PrefixLength: to.Ptr[int32](30),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
		// 				PublicIPAddresses: []*armnetwork.ReferencedPublicIPAddress{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/PublicIpAddress1"),
		// 				}},
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 			},
		// 			SKU: &armnetwork.PublicIPPrefixSKU{
		// 				Name: to.Ptr(armnetwork.PublicIPPrefixSKUNameStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ipprefix01"),
		// 			Type: to.Ptr("Microsoft.Network/publicIPPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/publicIPPrefixes/ipprefix01"),
		// 			Location: to.Ptr("westus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.PublicIPPrefixPropertiesFormat{
		// 				IPPrefix: to.Ptr("40.85.154.247/30"),
		// 				IPTags: []*armnetwork.IPTag{
		// 				},
		// 				PrefixLength: to.Ptr[int32](30),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 			},
		// 			SKU: &armnetwork.PublicIPPrefixSKU{
		// 				Name: to.Ptr(armnetwork.PublicIPPrefixSKUNameStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("pfx"),
		// 			Type: to.Ptr("Microsoft.Network/publicIPPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg3/providers/Microsoft.Network/publicIPPrefixes/pfx"),
		// 			Location: to.Ptr("westus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.PublicIPPrefixPropertiesFormat{
		// 				IPPrefix: to.Ptr("25.101.84.16/30"),
		// 				IPTags: []*armnetwork.IPTag{
		// 				},
		// 				LoadBalancerFrontendIPConfiguration: &armnetwork.SubResource{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg3/providers/Microsoft.Network/loadBalancers/lb-pfx/frontendIPConfigurations/ipconfig1"),
		// 				},
		// 				PrefixLength: to.Ptr[int32](30),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 			},
		// 			SKU: &armnetwork.PublicIPPrefixSKU{
		// 				Name: to.Ptr(armnetwork.PublicIPPrefixSKUNameStandard),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/639ecfad68419328658bd4cfe7094af4ce472be2/specification/network/resource-manager/Microsoft.Network/stable/2023-06-01/examples/PublicIpPrefixList.json
func ExamplePublicIPPrefixesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPublicIPPrefixesClient().NewListPager("rg1", nil)
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
		// page.PublicIPPrefixListResult = armnetwork.PublicIPPrefixListResult{
		// 	Value: []*armnetwork.PublicIPPrefix{
		// 		{
		// 			Name: to.Ptr("test-ipprefix"),
		// 			Type: to.Ptr("Microsoft.Network/publicIPPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPPrefixes/test-ipprefix"),
		// 			Location: to.Ptr("westus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.PublicIPPrefixPropertiesFormat{
		// 				IPPrefix: to.Ptr("40.85.154.2/30"),
		// 				IPTags: []*armnetwork.IPTag{
		// 					{
		// 						IPTagType: to.Ptr("FirstPartyUsage"),
		// 						Tag: to.Ptr("SQL"),
		// 				}},
		// 				PrefixLength: to.Ptr[int32](30),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 			},
		// 			SKU: &armnetwork.PublicIPPrefixSKU{
		// 				Name: to.Ptr(armnetwork.PublicIPPrefixSKUNameStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ipprefix03"),
		// 			Type: to.Ptr("Microsoft.Network/publicIPPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPPrefixes/ipprefix03"),
		// 			Location: to.Ptr("westus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.PublicIPPrefixPropertiesFormat{
		// 				IPPrefix: to.Ptr("40.85.153.2/31"),
		// 				IPTags: []*armnetwork.IPTag{
		// 				},
		// 				PrefixLength: to.Ptr[int32](31),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 			},
		// 			SKU: &armnetwork.PublicIPPrefixSKU{
		// 				Name: to.Ptr(armnetwork.PublicIPPrefixSKUNameStandard),
		// 			},
		// 	}},
		// }
	}
}
