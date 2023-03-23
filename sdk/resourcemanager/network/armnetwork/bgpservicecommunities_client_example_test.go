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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a60468a0c5e2beb054680ae488fb9f92699f0a0d/specification/network/resource-manager/Microsoft.Network/stable/2022-09-01/examples/ServiceCommunityList.json
func ExampleBgpServiceCommunitiesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewBgpServiceCommunitiesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager(nil)
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
		// page.BgpServiceCommunityListResult = armnetwork.BgpServiceCommunityListResult{
		// 	Value: []*armnetwork.BgpServiceCommunity{
		// 		{
		// 			Name: to.Ptr("skype"),
		// 			Type: to.Ptr("Microsoft.Network/bgpServiceCommunities"),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Network/bgpServiceCommunities/skype"),
		// 			Properties: &armnetwork.BgpServiceCommunityPropertiesFormat{
		// 				BgpCommunities: []*armnetwork.BGPCommunity{
		// 					{
		// 						CommunityName: to.Ptr("Skype For Business Online"),
		// 						CommunityPrefixes: []*string{
		// 							to.Ptr("13.67.56.225/32"),
		// 							to.Ptr("13.67.186.105/32")},
		// 							CommunityValue: to.Ptr("12076:5030"),
		// 							IsAuthorizedToUse: to.Ptr(true),
		// 							ServiceGroup: to.Ptr("O365"),
		// 							ServiceSupportedRegion: to.Ptr("Global"),
		// 					}},
		// 					ServiceName: to.Ptr("skype"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("exchange"),
		// 				Type: to.Ptr("Microsoft.Network/bgpServiceCommunities"),
		// 				ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Network/bgpServiceCommunities/exchange"),
		// 				Properties: &armnetwork.BgpServiceCommunityPropertiesFormat{
		// 					BgpCommunities: []*armnetwork.BGPCommunity{
		// 						{
		// 							CommunityName: to.Ptr("Exchange Online"),
		// 							CommunityPrefixes: []*string{
		// 								to.Ptr("13.67.56.225/32"),
		// 								to.Ptr("13.67.186.105/32")},
		// 								CommunityValue: to.Ptr("12076:5040"),
		// 								IsAuthorizedToUse: to.Ptr(true),
		// 								ServiceGroup: to.Ptr("O365"),
		// 								ServiceSupportedRegion: to.Ptr("Global"),
		// 						}},
		// 						ServiceName: to.Ptr("exchange"),
		// 					},
		// 			}},
		// 		}
	}
}
