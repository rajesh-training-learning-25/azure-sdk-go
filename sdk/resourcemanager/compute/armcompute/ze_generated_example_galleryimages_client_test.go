//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/gallery/CreateOrUpdateASimpleGalleryImage.json
func ExampleGalleryImagesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewGalleryImagesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<gallery-name>",
		"<gallery-image-name>",
		armcompute.GalleryImage{
			Location: to.StringPtr("<location>"),
			Properties: &armcompute.GalleryImageProperties{
				HyperVGeneration: armcompute.HyperVGeneration("V1").ToPtr(),
				Identifier: &armcompute.GalleryImageIdentifier{
					Offer:     to.StringPtr("<offer>"),
					Publisher: to.StringPtr("<publisher>"),
					SKU:       to.StringPtr("<sku>"),
				},
				OSState: armcompute.OperatingSystemStateTypesGeneralized.ToPtr(),
				OSType:  armcompute.OperatingSystemTypesWindows.ToPtr(),
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
	log.Printf("Response result: %#v\n", res.GalleryImagesClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/gallery/UpdateASimpleGalleryImage.json
func ExampleGalleryImagesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewGalleryImagesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<gallery-name>",
		"<gallery-image-name>",
		armcompute.GalleryImageUpdate{
			Properties: &armcompute.GalleryImageProperties{
				HyperVGeneration: armcompute.HyperVGeneration("V1").ToPtr(),
				Identifier: &armcompute.GalleryImageIdentifier{
					Offer:     to.StringPtr("<offer>"),
					Publisher: to.StringPtr("<publisher>"),
					SKU:       to.StringPtr("<sku>"),
				},
				OSState: armcompute.OperatingSystemStateTypesGeneralized.ToPtr(),
				OSType:  armcompute.OperatingSystemTypesWindows.ToPtr(),
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
	log.Printf("Response result: %#v\n", res.GalleryImagesClientUpdateResult)
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/gallery/GetAGalleryImage.json
func ExampleGalleryImagesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewGalleryImagesClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<gallery-name>",
		"<gallery-image-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.GalleryImagesClientGetResult)
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/gallery/DeleteAGalleryImage.json
func ExampleGalleryImagesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewGalleryImagesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<gallery-name>",
		"<gallery-image-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/gallery/ListGalleryImagesInAGallery.json
func ExampleGalleryImagesClient_ListByGallery() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewGalleryImagesClient("<subscription-id>", cred, nil)
	pager := client.ListByGallery("<resource-group-name>",
		"<gallery-name>",
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
