//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
)

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/GetFunctionAppStacks.json
func ExampleProviderClient_GetFunctionAppStacks() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewProviderClient("<subscription-id>", cred, nil)
	pager := client.GetFunctionAppStacks(&armappservice.ProviderGetFunctionAppStacksOptions{StackOsType: nil})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("FunctionAppStack.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/GetFunctionAppStacksForLocation.json
func ExampleProviderClient_GetFunctionAppStacksForLocation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewProviderClient("<subscription-id>", cred, nil)
	pager := client.GetFunctionAppStacksForLocation("<location>",
		&armappservice.ProviderGetFunctionAppStacksForLocationOptions{StackOsType: nil})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("FunctionAppStack.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/GetWebAppStacksForLocation.json
func ExampleProviderClient_GetWebAppStacksForLocation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewProviderClient("<subscription-id>", cred, nil)
	pager := client.GetWebAppStacksForLocation("<location>",
		&armappservice.ProviderGetWebAppStacksForLocationOptions{StackOsType: nil})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("WebAppStack.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/ListOperations.json
func ExampleProviderClient_ListOperations() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewProviderClient("<subscription-id>", cred, nil)
	pager := client.ListOperations(nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
	}
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/GetWebAppStacks.json
func ExampleProviderClient_GetWebAppStacks() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewProviderClient("<subscription-id>", cred, nil)
	pager := client.GetWebAppStacks(&armappservice.ProviderGetWebAppStacksOptions{StackOsType: nil})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("WebAppStack.ID: %s\n", *v.ID)
		}
	}
}
