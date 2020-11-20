// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/arm/compute/2020-09-30/armcompute"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

func ExampleAvailabilitySetsOperations_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	client := armcompute.NewAvailabilitySetsClient(armcore.NewDefaultConnection(cred, nil), "<subscription ID>")
	resp, err := client.CreateOrUpdate(
		context.Background(),
		"<resource group name>",
		"<availability set name>",
		armcompute.AvailabilitySet{
			Resource: armcompute.Resource{
				Name:     to.StringPtr("<availability set name>"),
				Location: to.StringPtr("<Azure location>"),
			},
			SKU: &armcompute.SKU{
				Name: to.StringPtr(string(armcompute.AvailabilitySetSKUTypesAligned)),
			},
			Properties: &armcompute.AvailabilitySetProperties{
				PlatformFaultDomainCount:  to.Int32Ptr(1),
				PlatformUpdateDomainCount: to.Int32Ptr(1),
			},
		},
		nil,
	)
	if err != nil {
		log.Fatalf("failed to obtain a response: %v", err)
	}
	log.Printf("Availability set ID: %v", *resp.AvailabilitySet.ID)
}

func ExampleAvailabilitySetsOperations_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	client := armcompute.NewAvailabilitySetsClient(armcore.NewDefaultConnection(cred, nil), "<subscription ID>")
	resp, err := client.Get(context.Background(), "<resource group name>", "<availability set name>", nil)
	if err != nil {
		log.Fatalf("failed to obtain a response: %v", err)
	}
	log.Printf("Availability set ID: %s", *resp.AvailabilitySet.ID)
}
