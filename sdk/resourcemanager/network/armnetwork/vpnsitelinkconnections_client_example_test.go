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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/network/resource-manager/Microsoft.Network/stable/2023-09-01/examples/VpnSiteLinkConnectionGet.json
func ExampleVPNSiteLinkConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVPNSiteLinkConnectionsClient().Get(ctx, "rg1", "gateway1", "vpnConnection1", "Connection-Link1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VPNSiteLinkConnection = armnetwork.VPNSiteLinkConnection{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnGateways/gateway1/vpnConnections/vpnConnection1/VpnSiteLinkConnections/Connection-Link1"),
	// 	Name: to.Ptr("Connection-Link1"),
	// 	Type: to.Ptr("Microsoft.Network/vpnGateways/vpnConnections/VpnSiteLinkConnections"),
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.VPNSiteLinkConnectionProperties{
	// 		ConnectionBandwidth: to.Ptr[int32](200),
	// 		EgressBytesTransferred: to.Ptr[int64](0),
	// 		EnableBgp: to.Ptr(false),
	// 		EnableRateLimiting: to.Ptr(false),
	// 		IngressBytesTransferred: to.Ptr[int64](0),
	// 		IPSecPolicies: []*armnetwork.IPSecPolicy{
	// 		},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		RoutingWeight: to.Ptr[int32](0),
	// 		SharedKey: to.Ptr("key"),
	// 		UseLocalAzureIPAddress: to.Ptr(false),
	// 		UsePolicyBasedTrafficSelectors: to.Ptr(false),
	// 		VPNConnectionProtocolType: to.Ptr(armnetwork.VirtualNetworkGatewayConnectionProtocolIKEv2),
	// 		VPNGatewayCustomBgpAddresses: []*armnetwork.GatewayCustomBgpIPAddressIPConfiguration{
	// 			{
	// 				CustomBgpIPAddress: to.Ptr("169.254.21.1"),
	// 				IPConfigurationID: to.Ptr("Instance0"),
	// 			},
	// 			{
	// 				CustomBgpIPAddress: to.Ptr("169.254.21.3"),
	// 				IPConfigurationID: to.Ptr("Instance1"),
	// 		}},
	// 		VPNSiteLink: &armnetwork.SubResource{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnSites/vpnSite1/vpnSiteLinks/siteLink1"),
	// 		},
	// 	},
	// }
}
