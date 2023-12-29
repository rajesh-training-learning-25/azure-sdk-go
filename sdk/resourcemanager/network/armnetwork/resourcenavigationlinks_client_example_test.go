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

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/639ecfad68419328658bd4cfe7094af4ce472be2/specification/network/resource-manager/Microsoft.Network/stable/2023-06-01/examples/VirtualNetworkGetResourceNavigationLinks.json
func ExampleResourceNavigationLinksClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewResourceNavigationLinksClient().List(ctx, "rg1", "vnet", "subnet", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ResourceNavigationLinksListResult = armnetwork.ResourceNavigationLinksListResult{
	// 	Value: []*armnetwork.ResourceNavigationLink{
	// 		{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet/subnets/subnet/resourceNavigationLinks/redisCache_redis-tester"),
	// 			Name: to.Ptr("redisCache_redis-tester"),
	// 			Type: to.Ptr("Microsoft.Network/virtualNetworks/subnets/resourceNavigationLinks"),
	// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
	// 			Properties: &armnetwork.ResourceNavigationLinkFormat{
	// 				Link: to.Ptr("/subscriptions/subid/resourceGroups/another-rg/providers/Microsoft.Cache/Redis/redis-tester"),
	// 				LinkedResourceType: to.Ptr("Microsoft.Cache/redis"),
	// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			},
	// 	}},
	// }
}
