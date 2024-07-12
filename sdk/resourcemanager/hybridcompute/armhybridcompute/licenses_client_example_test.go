//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f4c6c8697c59f966db0d1e36b62df3af3bca9065/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-03-31-preview/examples/license/License_CreateOrUpdate.json
func ExampleLicensesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLicensesClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "{licenseName}", armhybridcompute.License{
		Location: to.Ptr("eastus2euap"),
		Properties: &armhybridcompute.LicenseProperties{
			LicenseDetails: &armhybridcompute.LicenseDetails{
				Type:       to.Ptr(armhybridcompute.LicenseCoreTypePCore),
				Edition:    to.Ptr(armhybridcompute.LicenseEditionDatacenter),
				Processors: to.Ptr[int32](6),
				State:      to.Ptr(armhybridcompute.LicenseStateActivated),
				Target:     to.Ptr(armhybridcompute.LicenseTargetWindowsServer2012),
			},
			LicenseType: to.Ptr(armhybridcompute.LicenseTypeESU),
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
	// res.License = armhybridcompute.License{
	// 	Name: to.Ptr("{licenseName}"),
	// 	Type: to.Ptr("Microsoft.HybridCompute/licenses"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/Licenses/{licenseName}"),
	// 	Location: to.Ptr("eastus2euap"),
	// 	Properties: &armhybridcompute.LicenseProperties{
	// 		LicenseDetails: &armhybridcompute.LicenseDetails{
	// 			Type: to.Ptr(armhybridcompute.LicenseCoreTypePCore),
	// 			AssignedLicenses: to.Ptr[int32](2),
	// 			Edition: to.Ptr(armhybridcompute.LicenseEditionDatacenter),
	// 			ImmutableID: to.Ptr("<generated Guid>"),
	// 			Processors: to.Ptr[int32](6),
	// 			State: to.Ptr(armhybridcompute.LicenseStateActivated),
	// 			Target: to.Ptr(armhybridcompute.LicenseTargetWindowsServer2012),
	// 		},
	// 		LicenseType: to.Ptr(armhybridcompute.LicenseTypeESU),
	// 		TenantID: to.Ptr("{tenandId}"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f4c6c8697c59f966db0d1e36b62df3af3bca9065/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-03-31-preview/examples/license/License_Update.json
func ExampleLicensesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLicensesClient().BeginUpdate(ctx, "myResourceGroup", "{licenseName}", armhybridcompute.LicenseUpdate{
		Properties: &armhybridcompute.LicenseUpdateProperties{
			LicenseDetails: &armhybridcompute.LicenseUpdatePropertiesLicenseDetails{
				Type:       to.Ptr(armhybridcompute.LicenseCoreTypePCore),
				Edition:    to.Ptr(armhybridcompute.LicenseEditionDatacenter),
				Processors: to.Ptr[int32](6),
				State:      to.Ptr(armhybridcompute.LicenseStateActivated),
				Target:     to.Ptr(armhybridcompute.LicenseTargetWindowsServer2012),
			},
			LicenseType: to.Ptr(armhybridcompute.LicenseTypeESU),
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
	// res.License = armhybridcompute.License{
	// 	Name: to.Ptr("{licenseName}"),
	// 	Type: to.Ptr("Microsoft.HybridCompute/licenses"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/licenses/{licenseName}"),
	// 	Location: to.Ptr("eastus2euap"),
	// 	Properties: &armhybridcompute.LicenseProperties{
	// 		LicenseDetails: &armhybridcompute.LicenseDetails{
	// 			Type: to.Ptr(armhybridcompute.LicenseCoreTypePCore),
	// 			AssignedLicenses: to.Ptr[int32](8),
	// 			Edition: to.Ptr(armhybridcompute.LicenseEditionDatacenter),
	// 			ImmutableID: to.Ptr("<generated Guid>"),
	// 			Processors: to.Ptr[int32](6),
	// 			State: to.Ptr(armhybridcompute.LicenseStateActivated),
	// 			Target: to.Ptr(armhybridcompute.LicenseTargetWindowsServer2012),
	// 		},
	// 		LicenseType: to.Ptr(armhybridcompute.LicenseTypeESU),
	// 		TenantID: to.Ptr("{tenandId}"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f4c6c8697c59f966db0d1e36b62df3af3bca9065/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-03-31-preview/examples/license/License_Get.json
func ExampleLicensesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLicensesClient().Get(ctx, "myResourceGroup", "{licenseName}", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.License = armhybridcompute.License{
	// 	Name: to.Ptr("{licenseName}"),
	// 	Type: to.Ptr("Microsoft.HybridCompute/licenses"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/Licenses/{licenseName}"),
	// 	Location: to.Ptr("eastus2euap"),
	// 	Properties: &armhybridcompute.LicenseProperties{
	// 		LicenseDetails: &armhybridcompute.LicenseDetails{
	// 			Type: to.Ptr(armhybridcompute.LicenseCoreTypePCore),
	// 			AssignedLicenses: to.Ptr[int32](8),
	// 			Edition: to.Ptr(armhybridcompute.LicenseEditionDatacenter),
	// 			ImmutableID: to.Ptr("<generated Guid>"),
	// 			Processors: to.Ptr[int32](6),
	// 			State: to.Ptr(armhybridcompute.LicenseStateActivated),
	// 			Target: to.Ptr(armhybridcompute.LicenseTargetWindowsServer2012),
	// 		},
	// 		LicenseType: to.Ptr(armhybridcompute.LicenseTypeESU),
	// 		TenantID: to.Ptr("{tenandId}"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f4c6c8697c59f966db0d1e36b62df3af3bca9065/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-03-31-preview/examples/license/License_Delete.json
func ExampleLicensesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLicensesClient().BeginDelete(ctx, "myResourceGroup", "{licenseName}", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f4c6c8697c59f966db0d1e36b62df3af3bca9065/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-03-31-preview/examples/license/License_ListByResourceGroup.json
func ExampleLicensesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLicensesClient().NewListByResourceGroupPager("myResourceGroup", nil)
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
		// page.LicensesListResult = armhybridcompute.LicensesListResult{
		// 	Value: []*armhybridcompute.License{
		// 		{
		// 			Name: to.Ptr("{licenseName}"),
		// 			Type: to.Ptr("Microsoft.HybridCompute/licenses"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/Licenses/{licenseName}"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			Properties: &armhybridcompute.LicenseProperties{
		// 				LicenseDetails: &armhybridcompute.LicenseDetails{
		// 					Type: to.Ptr(armhybridcompute.LicenseCoreTypePCore),
		// 					AssignedLicenses: to.Ptr[int32](8),
		// 					Edition: to.Ptr(armhybridcompute.LicenseEditionDatacenter),
		// 					ImmutableID: to.Ptr("<generated Guid>"),
		// 					Processors: to.Ptr[int32](6),
		// 					State: to.Ptr(armhybridcompute.LicenseStateActivated),
		// 					Target: to.Ptr(armhybridcompute.LicenseTargetWindowsServer2012),
		// 				},
		// 				LicenseType: to.Ptr(armhybridcompute.LicenseTypeESU),
		// 				TenantID: to.Ptr("{tenandId}"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f4c6c8697c59f966db0d1e36b62df3af3bca9065/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-03-31-preview/examples/license/License_ListBySubscription.json
func ExampleLicensesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLicensesClient().NewListBySubscriptionPager(nil)
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
		// page.LicensesListResult = armhybridcompute.LicensesListResult{
		// 	Value: []*armhybridcompute.License{
		// 		{
		// 			Name: to.Ptr("{licenseName}"),
		// 			Type: to.Ptr("Microsoft.HybridCompute/licenses"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/Licenses/{licenseName}"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			Properties: &armhybridcompute.LicenseProperties{
		// 				LicenseDetails: &armhybridcompute.LicenseDetails{
		// 					Type: to.Ptr(armhybridcompute.LicenseCoreTypePCore),
		// 					AssignedLicenses: to.Ptr[int32](8),
		// 					Edition: to.Ptr(armhybridcompute.LicenseEditionDatacenter),
		// 					ImmutableID: to.Ptr("<generated Guid>"),
		// 					Processors: to.Ptr[int32](6),
		// 					State: to.Ptr(armhybridcompute.LicenseStateActivated),
		// 					Target: to.Ptr(armhybridcompute.LicenseTargetWindowsServer2012),
		// 				},
		// 				LicenseType: to.Ptr(armhybridcompute.LicenseTypeESU),
		// 				TenantID: to.Ptr("{tenandId}"),
		// 			},
		// 	}},
		// }
	}
}
