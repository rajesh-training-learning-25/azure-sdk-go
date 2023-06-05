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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/AzureWebCategoryGet.json
func ExampleWebCategoriesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWebCategoriesClient().Get(ctx, "Arts", &armnetwork.WebCategoriesClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AzureWebCategory = armnetwork.AzureWebCategory{
	// 	Name: to.Ptr("Arts"),
	// 	Type: to.Ptr("Microsoft.Network/azureWebCategories"),
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	ID: to.Ptr("/subscriptions/4de8428a-4a92-4cea-90ff-b47128b8cab8/providers/Microsoft.Network/azureWebCategories/Arts"),
	// 	Properties: &armnetwork.AzureWebCategoryPropertiesFormat{
	// 		Group: to.Ptr("General"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/AzureWebCategoriesListBySubscription.json
func ExampleWebCategoriesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWebCategoriesClient().NewListBySubscriptionPager(nil)
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
		// page.AzureWebCategoryListResult = armnetwork.AzureWebCategoryListResult{
		// 	Value: []*armnetwork.AzureWebCategory{
		// 		{
		// 			Name: to.Ptr("Arts"),
		// 			Type: to.Ptr("Microsoft.Network/azureWebCategories"),
		// 			Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 			ID: to.Ptr("/subscriptions/4de8428a-4a92-4cea-90ff-b47128b8cab8/providers/Microsoft.Network/azureWebCategories/Arts"),
		// 			Properties: &armnetwork.AzureWebCategoryPropertiesFormat{
		// 				Group: to.Ptr("General"),
		// 			},
		// 	}},
		// }
	}
}
