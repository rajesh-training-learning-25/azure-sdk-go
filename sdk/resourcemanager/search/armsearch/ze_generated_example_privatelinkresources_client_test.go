//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsearch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch"
)

// x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2020-08-01/examples/ListSupportedPrivateLinkResources.json
func ExamplePrivateLinkResourcesClient_ListSupported() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsearch.NewPrivateLinkResourcesClient("<subscription-id>", cred, nil)
	res, err := client.ListSupported(ctx,
		"<resource-group-name>",
		"<search-service-name>",
		&armsearch.SearchManagementRequestOptions{ClientRequestID: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PrivateLinkResourcesClientListSupportedResult)
}
