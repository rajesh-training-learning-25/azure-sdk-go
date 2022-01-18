//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappplatform_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"
)

// x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-01-01-preview/examples/GatewayCustomDomains_Get.json
func ExampleGatewayCustomDomainsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappplatform.NewGatewayCustomDomainsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<gateway-name>",
		"<domain-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.GatewayCustomDomainsClientGetResult)
}

// x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-01-01-preview/examples/GatewayCustomDomains_CreateOrUpdate.json
func ExampleGatewayCustomDomainsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappplatform.NewGatewayCustomDomainsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<gateway-name>",
		"<domain-name>",
		armappplatform.GatewayCustomDomainResource{
			Properties: &armappplatform.GatewayCustomDomainProperties{
				Thumbprint: to.StringPtr("<thumbprint>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.GatewayCustomDomainsClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-01-01-preview/examples/GatewayCustomDomains_Delete.json
func ExampleGatewayCustomDomainsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappplatform.NewGatewayCustomDomainsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<gateway-name>",
		"<domain-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-01-01-preview/examples/GatewayCustomDomains_List.json
func ExampleGatewayCustomDomainsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappplatform.NewGatewayCustomDomainsClient("<subscription-id>", cred, nil)
	pager := client.List("<resource-group-name>",
		"<service-name>",
		"<gateway-name>",
		nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}
