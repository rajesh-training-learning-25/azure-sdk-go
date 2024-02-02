//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armstoragecache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/73d1e6522dfdedb795f46cf997921c623011caa6/specification/storagecache/resource-manager/Microsoft.StorageCache/preview/2023-11-01-preview/examples/checkAmlFSSubnets.json
func ExampleManagementClient_CheckAmlFSSubnets() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragecache.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewManagementClient().CheckAmlFSSubnets(ctx, &armstoragecache.ManagementClientCheckAmlFSSubnetsOptions{AmlFilesystemSubnetInfo: &armstoragecache.AmlFilesystemSubnetInfo{
		FilesystemSubnet: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.Network/virtualNetworks/scvnet/subnets/fsSub"),
		SKU: &armstoragecache.SKUName{
			Name: to.Ptr("AMLFS-Durable-Premium-125"),
		},
		StorageCapacityTiB: to.Ptr[float32](16),
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/73d1e6522dfdedb795f46cf997921c623011caa6/specification/storagecache/resource-manager/Microsoft.StorageCache/preview/2023-11-01-preview/examples/getRequiredAmlFSSubnetsSize.json
func ExampleManagementClient_GetRequiredAmlFSSubnetsSize() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragecache.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagementClient().GetRequiredAmlFSSubnetsSize(ctx, &armstoragecache.ManagementClientGetRequiredAmlFSSubnetsSizeOptions{RequiredAMLFilesystemSubnetsSizeInfo: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RequiredAmlFilesystemSubnetsSize = armstoragecache.RequiredAmlFilesystemSubnetsSize{
	// 	FilesystemSubnetSize: to.Ptr[int32](24),
	// }
}
