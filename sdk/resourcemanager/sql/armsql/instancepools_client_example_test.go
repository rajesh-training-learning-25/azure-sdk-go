//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edd7863aeaad7379a7a8eef424ca11617d8c91eb/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/GetInstancePool.json
func ExampleInstancePoolsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInstancePoolsClient().Get(ctx, "group1", "testIP", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.InstancePool = armsql.InstancePool{
	// 	Name: to.Ptr("testIP"),
	// 	Type: to.Ptr("Microsoft.Sql/instancePools"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/instancePools/testIP"),
	// 	Location: to.Ptr("japaneast"),
	// 	Tags: map[string]*string{
	// 		"a": to.Ptr("b"),
	// 	},
	// 	Properties: &armsql.InstancePoolProperties{
	// 		LicenseType: to.Ptr(armsql.InstancePoolLicenseTypeLicenseIncluded),
	// 		SubnetID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Network/virtualNetwork/myvnet/subnets/mysubnet1"),
	// 		VCores: to.Ptr[int32](8),
	// 	},
	// 	SKU: &armsql.SKU{
	// 		Name: to.Ptr("GP_Gen5"),
	// 		Family: to.Ptr("Gen5"),
	// 		Tier: to.Ptr("GeneralPurpose"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edd7863aeaad7379a7a8eef424ca11617d8c91eb/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/CreateOrUpdateInstancePoolMax.json
func ExampleInstancePoolsClient_BeginCreateOrUpdate_createAnInstancePoolWithAllProperties() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewInstancePoolsClient().BeginCreateOrUpdate(ctx, "group1", "testIP", armsql.InstancePool{
		Location: to.Ptr("japaneast"),
		Tags: map[string]*string{
			"a": to.Ptr("b"),
		},
		Properties: &armsql.InstancePoolProperties{
			LicenseType: to.Ptr(armsql.InstancePoolLicenseTypeLicenseIncluded),
			SubnetID:    to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Network/virtualNetworks/myvnet/subnets/mysubnet1"),
			VCores:      to.Ptr[int32](8),
		},
		SKU: &armsql.SKU{
			Name:   to.Ptr("GP_Gen5"),
			Family: to.Ptr("Gen5"),
			Tier:   to.Ptr("GeneralPurpose"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.InstancePool = armsql.InstancePool{
	// 	Name: to.Ptr("testIP"),
	// 	Type: to.Ptr("Microsoft.Sql/instancePools"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/instancePools/testIP"),
	// 	Location: to.Ptr("japaneast"),
	// 	Tags: map[string]*string{
	// 		"a": to.Ptr("b"),
	// 	},
	// 	Properties: &armsql.InstancePoolProperties{
	// 		LicenseType: to.Ptr(armsql.InstancePoolLicenseTypeLicenseIncluded),
	// 		SubnetID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Network/virtualNetworks/myvnet/subnets/mysubnet1"),
	// 		VCores: to.Ptr[int32](8),
	// 	},
	// 	SKU: &armsql.SKU{
	// 		Name: to.Ptr("GP_Gen5"),
	// 		Family: to.Ptr("Gen5"),
	// 		Tier: to.Ptr("GeneralPurpose"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edd7863aeaad7379a7a8eef424ca11617d8c91eb/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/CreateOrUpdateInstancePoolMin.json
func ExampleInstancePoolsClient_BeginCreateOrUpdate_createAnInstancePoolWithMinProperties() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewInstancePoolsClient().BeginCreateOrUpdate(ctx, "group1", "testIP", armsql.InstancePool{
		Location: to.Ptr("japaneast"),
		Properties: &armsql.InstancePoolProperties{
			LicenseType: to.Ptr(armsql.InstancePoolLicenseTypeLicenseIncluded),
			SubnetID:    to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Network/virtualNetworks/myvnet/subnets/mysubnet1"),
			VCores:      to.Ptr[int32](8),
		},
		SKU: &armsql.SKU{
			Name:   to.Ptr("GP_Gen5"),
			Family: to.Ptr("Gen5"),
			Tier:   to.Ptr("GeneralPurpose"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.InstancePool = armsql.InstancePool{
	// 	Name: to.Ptr("testIP"),
	// 	Type: to.Ptr("Microsoft.Sql/instancePools"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/instancePools/testIP"),
	// 	Location: to.Ptr("japaneast"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armsql.InstancePoolProperties{
	// 		LicenseType: to.Ptr(armsql.InstancePoolLicenseTypeLicenseIncluded),
	// 		SubnetID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Network/virtualNetworks/myvnet/subnets/mysubnet1"),
	// 		VCores: to.Ptr[int32](8),
	// 	},
	// 	SKU: &armsql.SKU{
	// 		Name: to.Ptr("GP_Gen5"),
	// 		Family: to.Ptr("Gen5"),
	// 		Tier: to.Ptr("GeneralPurpose"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edd7863aeaad7379a7a8eef424ca11617d8c91eb/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DeleteInstancePool.json
func ExampleInstancePoolsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewInstancePoolsClient().BeginDelete(ctx, "group1", "testIP", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edd7863aeaad7379a7a8eef424ca11617d8c91eb/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/PatchInstancePool.json
func ExampleInstancePoolsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewInstancePoolsClient().BeginUpdate(ctx, "group1", "testIP", armsql.InstancePoolUpdate{
		Tags: map[string]*string{
			"x": to.Ptr("y"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.InstancePool = armsql.InstancePool{
	// 	Name: to.Ptr("testIP"),
	// 	Type: to.Ptr("Microsoft.Sql/instancePools"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/instancePools/testIP"),
	// 	Location: to.Ptr("japaneast"),
	// 	Tags: map[string]*string{
	// 		"x": to.Ptr("y"),
	// 	},
	// 	Properties: &armsql.InstancePoolProperties{
	// 		LicenseType: to.Ptr(armsql.InstancePoolLicenseTypeLicenseIncluded),
	// 		SubnetID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Network/virtualNetwork/myvnet/subnets/mysubnet1"),
	// 		VCores: to.Ptr[int32](8),
	// 	},
	// 	SKU: &armsql.SKU{
	// 		Name: to.Ptr("GP_Gen5"),
	// 		Family: to.Ptr("Gen5"),
	// 		Tier: to.Ptr("GeneralPurpose"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edd7863aeaad7379a7a8eef424ca11617d8c91eb/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ListInstancePoolsByResourceGroup.json
func ExampleInstancePoolsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewInstancePoolsClient().NewListByResourceGroupPager("group1", nil)
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
		// page.InstancePoolListResult = armsql.InstancePoolListResult{
		// 	Value: []*armsql.InstancePool{
		// 		{
		// 			Name: to.Ptr("testIP"),
		// 			Type: to.Ptr("Microsoft.Sql/instancePools"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/instancePools/testIP"),
		// 			Location: to.Ptr("japaneast"),
		// 			Tags: map[string]*string{
		// 				"a": to.Ptr("b"),
		// 			},
		// 			Properties: &armsql.InstancePoolProperties{
		// 				LicenseType: to.Ptr(armsql.InstancePoolLicenseTypeLicenseIncluded),
		// 				SubnetID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Network/virtualNetwork/myvnet/subnets/mysubnet1"),
		// 				VCores: to.Ptr[int32](8),
		// 			},
		// 			SKU: &armsql.SKU{
		// 				Name: to.Ptr("GP_Gen5"),
		// 				Family: to.Ptr("Gen5"),
		// 				Tier: to.Ptr("GeneralPurpose"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testIP2"),
		// 			Type: to.Ptr("Microsoft.Sql/instancePools"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/instancePools/testIP2"),
		// 			Location: to.Ptr("japaneast"),
		// 			Tags: map[string]*string{
		// 				"a": to.Ptr("b"),
		// 			},
		// 			Properties: &armsql.InstancePoolProperties{
		// 				LicenseType: to.Ptr(armsql.InstancePoolLicenseTypeLicenseIncluded),
		// 				SubnetID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Network/virtualNetwork/myvnet/subnets/mysubnet1"),
		// 				VCores: to.Ptr[int32](8),
		// 			},
		// 			SKU: &armsql.SKU{
		// 				Name: to.Ptr("GP_Gen5"),
		// 				Family: to.Ptr("Gen5"),
		// 				Tier: to.Ptr("GeneralPurpose"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edd7863aeaad7379a7a8eef424ca11617d8c91eb/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ListInstancePoolsBySubscriptionId.json
func ExampleInstancePoolsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewInstancePoolsClient().NewListPager(nil)
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
		// page.InstancePoolListResult = armsql.InstancePoolListResult{
		// 	Value: []*armsql.InstancePool{
		// 		{
		// 			Name: to.Ptr("testIP"),
		// 			Type: to.Ptr("Microsoft.Sql/instancePools"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/instancePools/testIP"),
		// 			Location: to.Ptr("japaneast"),
		// 			Tags: map[string]*string{
		// 				"a": to.Ptr("b"),
		// 			},
		// 			Properties: &armsql.InstancePoolProperties{
		// 				LicenseType: to.Ptr(armsql.InstancePoolLicenseTypeLicenseIncluded),
		// 				SubnetID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Network/virtualNetwork/myvnet/subnets/mysubnet1"),
		// 				VCores: to.Ptr[int32](8),
		// 			},
		// 			SKU: &armsql.SKU{
		// 				Name: to.Ptr("GP_Gen5"),
		// 				Family: to.Ptr("Gen5"),
		// 				Tier: to.Ptr("GeneralPurpose"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testIP2"),
		// 			Type: to.Ptr("Microsoft.Sql/instancePools"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group2/providers/Microsoft.Sql/instancePools/testIP2"),
		// 			Location: to.Ptr("japaneast"),
		// 			Tags: map[string]*string{
		// 				"a": to.Ptr("b"),
		// 			},
		// 			Properties: &armsql.InstancePoolProperties{
		// 				LicenseType: to.Ptr(armsql.InstancePoolLicenseTypeLicenseIncluded),
		// 				SubnetID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group2/providers/Microsoft.Network/virtualNetwork/myvnet/subnets/mysubnet1"),
		// 				VCores: to.Ptr[int32](8),
		// 			},
		// 			SKU: &armsql.SKU{
		// 				Name: to.Ptr("GP_Gen5"),
		// 				Family: to.Ptr("Gen5"),
		// 				Tier: to.Ptr("GeneralPurpose"),
		// 			},
		// 	}},
		// }
	}
}
