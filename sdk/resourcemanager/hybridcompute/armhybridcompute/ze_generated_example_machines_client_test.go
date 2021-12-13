//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute"
)

// x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2021-06-10-preview/examples/Machines_CreateOrUpdate.json
func ExampleMachinesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybridcompute.NewMachinesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<machine-name>",
		armhybridcompute.Machine{
			TrackedResource: armhybridcompute.TrackedResource{
				Location: to.StringPtr("<location>"),
			},
			Identity: &armhybridcompute.Identity{
				Type: to.StringPtr("<type>"),
			},
			Properties: &armhybridcompute.MachineProperties{
				ClientPublicKey: to.StringPtr("<client-public-key>"),
				LocationData: &armhybridcompute.LocationData{
					Name: to.StringPtr("<name>"),
				},
				ParentClusterResourceID:    to.StringPtr("<parent-cluster-resource-id>"),
				PrivateLinkScopeResourceID: to.StringPtr("<private-link-scope-resource-id>"),
				VMID:                       to.StringPtr("<vmid>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Machine.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2021-06-10-preview/examples/Machines_Update.json
func ExampleMachinesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybridcompute.NewMachinesClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<machine-name>",
		armhybridcompute.MachineUpdate{
			Identity: &armhybridcompute.Identity{
				Type: to.StringPtr("<type>"),
			},
			Properties: &armhybridcompute.MachineUpdateProperties{
				LocationData: &armhybridcompute.LocationData{
					Name: to.StringPtr("<name>"),
				},
				OSProfile: &armhybridcompute.OSProfile{
					LinuxConfiguration: &armhybridcompute.OSProfileLinuxConfiguration{
						PatchSettings: &armhybridcompute.PatchSettings{
							AssessmentMode: to.StringPtr("<assessment-mode>"),
						},
					},
					WindowsConfiguration: &armhybridcompute.OSProfileWindowsConfiguration{
						PatchSettings: &armhybridcompute.PatchSettings{
							AssessmentMode: to.StringPtr("<assessment-mode>"),
						},
					},
				},
				ParentClusterResourceID:    to.StringPtr("<parent-cluster-resource-id>"),
				PrivateLinkScopeResourceID: to.StringPtr("<private-link-scope-resource-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Machine.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2021-06-10-preview/examples/Machines_Delete.json
func ExampleMachinesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybridcompute.NewMachinesClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<machine-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2021-06-10-preview/examples/Machines_Get.json
func ExampleMachinesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybridcompute.NewMachinesClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<machine-name>",
		&armhybridcompute.MachinesGetOptions{Expand: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Machine.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2021-06-10-preview/examples/Machines_ListByResourceGroup.json
func ExampleMachinesClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybridcompute.NewMachinesClient("<subscription-id>", cred, nil)
	pager := client.ListByResourceGroup("<resource-group-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Machine.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2021-06-10-preview/examples/Machines_ListBySubscription.json
func ExampleMachinesClient_ListBySubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybridcompute.NewMachinesClient("<subscription-id>", cred, nil)
	pager := client.ListBySubscription(nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Machine.ID: %s\n", *v.ID)
		}
	}
}
