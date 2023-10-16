//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/examples/ManagedHsm_CreateOrUpdate.json
func ExampleManagedHsmsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedHsmsClient().BeginCreateOrUpdate(ctx, "hsm-group", "hsm1", armkeyvault.ManagedHsm{
		Location: to.Ptr("westus"),
		SKU: &armkeyvault.ManagedHsmSKU{
			Name:   to.Ptr(armkeyvault.ManagedHsmSKUNameStandardB1),
			Family: to.Ptr(armkeyvault.ManagedHsmSKUFamilyB),
		},
		Tags: map[string]*string{
			"Dept":        to.Ptr("hsm"),
			"Environment": to.Ptr("dogfood"),
		},
		Properties: &armkeyvault.ManagedHsmProperties{
			EnablePurgeProtection: to.Ptr(false),
			EnableSoftDelete:      to.Ptr(true),
			InitialAdminObjectIDs: []*string{
				to.Ptr("00000000-0000-0000-0000-000000000000")},
			SoftDeleteRetentionInDays: to.Ptr[int32](90),
			TenantID:                  to.Ptr("00000000-0000-0000-0000-000000000000"),
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
	// res.ManagedHsm = armkeyvault.ManagedHsm{
	// 	Name: to.Ptr("hsm1"),
	// 	Type: to.Ptr("Microsoft.KeyVault/managedHSMs"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.KeyVault/managedHSMs/hsm1"),
	// 	Location: to.Ptr("westus"),
	// 	SKU: &armkeyvault.ManagedHsmSKU{
	// 		Name: to.Ptr(armkeyvault.ManagedHsmSKUNameStandardB1),
	// 		Family: to.Ptr(armkeyvault.ManagedHsmSKUFamilyB),
	// 	},
	// 	Tags: map[string]*string{
	// 		"Dept": to.Ptr("hsm"),
	// 		"Environment": to.Ptr("dogfood"),
	// 	},
	// 	Properties: &armkeyvault.ManagedHsmProperties{
	// 		EnablePurgeProtection: to.Ptr(false),
	// 		EnableSoftDelete: to.Ptr(true),
	// 		HsmURI: to.Ptr("https://westus.hsm1.managedhsm.azure.net"),
	// 		InitialAdminObjectIDs: []*string{
	// 			to.Ptr("00000000-0000-0000-0000-000000000000")},
	// 			ProvisioningState: to.Ptr(armkeyvault.ProvisioningStateSucceeded),
	// 			SoftDeleteRetentionInDays: to.Ptr[int32](90),
	// 			StatusMessage: to.Ptr("ManagedHsm is functional."),
	// 			TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/examples/ManagedHsm_Update.json
func ExampleManagedHsmsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedHsmsClient().BeginUpdate(ctx, "hsm-group", "hsm1", armkeyvault.ManagedHsm{
		Tags: map[string]*string{
			"Dept":        to.Ptr("hsm"),
			"Environment": to.Ptr("dogfood"),
			"Slice":       to.Ptr("A"),
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
	// res.ManagedHsm = armkeyvault.ManagedHsm{
	// 	Name: to.Ptr("hsm1"),
	// 	Type: to.Ptr("Microsoft.KeyVault/managedHSMs"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.KeyVault/managedHSMs/hsm1"),
	// 	Location: to.Ptr("westus"),
	// 	SKU: &armkeyvault.ManagedHsmSKU{
	// 		Name: to.Ptr(armkeyvault.ManagedHsmSKUNameStandardB1),
	// 		Family: to.Ptr(armkeyvault.ManagedHsmSKUFamilyB),
	// 	},
	// 	Tags: map[string]*string{
	// 		"Dept": to.Ptr("hsm"),
	// 		"Environment": to.Ptr("dogfood"),
	// 		"Slice": to.Ptr("A"),
	// 	},
	// 	Properties: &armkeyvault.ManagedHsmProperties{
	// 		EnablePurgeProtection: to.Ptr(false),
	// 		EnableSoftDelete: to.Ptr(true),
	// 		HsmURI: to.Ptr("https://westus.hsm1.managedhsm.azure.net"),
	// 		InitialAdminObjectIDs: []*string{
	// 			to.Ptr("00000000-0000-0000-0000-000000000000")},
	// 			ProvisioningState: to.Ptr(armkeyvault.ProvisioningStateSucceeded),
	// 			SoftDeleteRetentionInDays: to.Ptr[int32](90),
	// 			StatusMessage: to.Ptr("ManagedHsm is functional."),
	// 			TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/examples/ManagedHsm_Delete.json
func ExampleManagedHsmsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedHsmsClient().BeginDelete(ctx, "hsm-group", "hsm1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/examples/ManagedHsm_Get.json
func ExampleManagedHsmsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedHsmsClient().Get(ctx, "hsm-group", "hsm1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedHsm = armkeyvault.ManagedHsm{
	// 	Name: to.Ptr("hsm1"),
	// 	Type: to.Ptr("Microsoft.KeyVault/managedHSMs"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.KeyVault/managedHSMs/hsm1"),
	// 	Location: to.Ptr("westus"),
	// 	SKU: &armkeyvault.ManagedHsmSKU{
	// 		Name: to.Ptr(armkeyvault.ManagedHsmSKUNameStandardB1),
	// 		Family: to.Ptr(armkeyvault.ManagedHsmSKUFamilyB),
	// 	},
	// 	Tags: map[string]*string{
	// 		"Dept": to.Ptr("hsm"),
	// 		"Environment": to.Ptr("dogfood"),
	// 	},
	// 	Properties: &armkeyvault.ManagedHsmProperties{
	// 		EnablePurgeProtection: to.Ptr(false),
	// 		EnableSoftDelete: to.Ptr(true),
	// 		HsmURI: to.Ptr("https://westus.hsm1.managedhsm.azure.net"),
	// 		InitialAdminObjectIDs: []*string{
	// 			to.Ptr("00000000-0000-0000-0000-000000000000")},
	// 			ProvisioningState: to.Ptr(armkeyvault.ProvisioningStateSucceeded),
	// 			SoftDeleteRetentionInDays: to.Ptr[int32](90),
	// 			StatusMessage: to.Ptr("ManagedHsm is functional."),
	// 			TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/examples/ManagedHsm_ListByResourceGroup.json
func ExampleManagedHsmsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedHsmsClient().NewListByResourceGroupPager("hsm-group", &armkeyvault.ManagedHsmsClientListByResourceGroupOptions{Top: nil})
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
		// page.ManagedHsmListResult = armkeyvault.ManagedHsmListResult{
		// 	Value: []*armkeyvault.ManagedHsm{
		// 		{
		// 			Name: to.Ptr("hsm1"),
		// 			Type: to.Ptr("Microsoft.KeyVault/managedHSMs"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.KeyVault/managedHSMs/hsm1"),
		// 			Location: to.Ptr("westus"),
		// 			SKU: &armkeyvault.ManagedHsmSKU{
		// 				Name: to.Ptr(armkeyvault.ManagedHsmSKUNameStandardB1),
		// 				Family: to.Ptr(armkeyvault.ManagedHsmSKUFamilyB),
		// 			},
		// 			Tags: map[string]*string{
		// 				"Dept": to.Ptr("hsm"),
		// 				"Environment": to.Ptr("dogfood"),
		// 			},
		// 			Properties: &armkeyvault.ManagedHsmProperties{
		// 				EnablePurgeProtection: to.Ptr(false),
		// 				EnableSoftDelete: to.Ptr(true),
		// 				HsmURI: to.Ptr("https://westus.hsm1.managedhsm.azure.net"),
		// 				InitialAdminObjectIDs: []*string{
		// 					to.Ptr("00000000-0000-0000-0000-000000000000")},
		// 					ProvisioningState: to.Ptr(armkeyvault.ProvisioningStateSucceeded),
		// 					SoftDeleteRetentionInDays: to.Ptr[int32](90),
		// 					StatusMessage: to.Ptr("ManagedHsm is functional."),
		// 					TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("hsm2"),
		// 				Type: to.Ptr("Microsoft.KeyVault/managedHSMs"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.KeyVault/managedHSMs/hsm2"),
		// 				Location: to.Ptr("westus"),
		// 				SKU: &armkeyvault.ManagedHsmSKU{
		// 					Name: to.Ptr(armkeyvault.ManagedHsmSKUNameStandardB1),
		// 					Family: to.Ptr(armkeyvault.ManagedHsmSKUFamilyB),
		// 				},
		// 				Tags: map[string]*string{
		// 					"Dept": to.Ptr("hsm"),
		// 					"Environment": to.Ptr("production"),
		// 				},
		// 				Properties: &armkeyvault.ManagedHsmProperties{
		// 					EnablePurgeProtection: to.Ptr(false),
		// 					EnableSoftDelete: to.Ptr(true),
		// 					HsmURI: to.Ptr("https://westus.hsm2.managedhsm.azure.net"),
		// 					InitialAdminObjectIDs: []*string{
		// 						to.Ptr("00000000-0000-0000-0000-000000000000")},
		// 						ProvisioningState: to.Ptr(armkeyvault.ProvisioningStateSucceeded),
		// 						SoftDeleteRetentionInDays: to.Ptr[int32](90),
		// 						StatusMessage: to.Ptr("ManagedHsm is functional."),
		// 						TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					},
		// 			}},
		// 		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/examples/ManagedHsm_ListBySubscription.json
func ExampleManagedHsmsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedHsmsClient().NewListBySubscriptionPager(&armkeyvault.ManagedHsmsClientListBySubscriptionOptions{Top: nil})
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
		// page.ManagedHsmListResult = armkeyvault.ManagedHsmListResult{
		// 	Value: []*armkeyvault.ManagedHsm{
		// 		{
		// 			Name: to.Ptr("hsm1"),
		// 			Type: to.Ptr("Microsoft.KeyVault/managedHSMs"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.KeyVault/managedHSMs/hsm1"),
		// 			Location: to.Ptr("westus"),
		// 			SKU: &armkeyvault.ManagedHsmSKU{
		// 				Name: to.Ptr(armkeyvault.ManagedHsmSKUNameStandardB1),
		// 				Family: to.Ptr(armkeyvault.ManagedHsmSKUFamilyB),
		// 			},
		// 			Tags: map[string]*string{
		// 				"Dept": to.Ptr("hsm"),
		// 				"Environment": to.Ptr("dogfood"),
		// 			},
		// 			Properties: &armkeyvault.ManagedHsmProperties{
		// 				EnablePurgeProtection: to.Ptr(false),
		// 				EnableSoftDelete: to.Ptr(true),
		// 				HsmURI: to.Ptr("https://westus.hsm1.managedhsm.azure.net"),
		// 				InitialAdminObjectIDs: []*string{
		// 					to.Ptr("00000000-0000-0000-0000-000000000000")},
		// 					ProvisioningState: to.Ptr(armkeyvault.ProvisioningStateSucceeded),
		// 					SoftDeleteRetentionInDays: to.Ptr[int32](90),
		// 					StatusMessage: to.Ptr("ManagedHsm is functional."),
		// 					TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("hsm2"),
		// 				Type: to.Ptr("Microsoft.KeyVault/managedHSMs"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.KeyVault/managedHSMs/hsm2"),
		// 				Location: to.Ptr("westus"),
		// 				SKU: &armkeyvault.ManagedHsmSKU{
		// 					Name: to.Ptr(armkeyvault.ManagedHsmSKUNameStandardB1),
		// 					Family: to.Ptr(armkeyvault.ManagedHsmSKUFamilyB),
		// 				},
		// 				Tags: map[string]*string{
		// 					"Dept": to.Ptr("hsm"),
		// 					"Environment": to.Ptr("production"),
		// 				},
		// 				Properties: &armkeyvault.ManagedHsmProperties{
		// 					EnablePurgeProtection: to.Ptr(false),
		// 					EnableSoftDelete: to.Ptr(true),
		// 					HsmURI: to.Ptr("https://westus.hsm2.managedhsm.azure.net"),
		// 					InitialAdminObjectIDs: []*string{
		// 						to.Ptr("00000000-0000-0000-0000-000000000000")},
		// 						ProvisioningState: to.Ptr(armkeyvault.ProvisioningStateSucceeded),
		// 						SoftDeleteRetentionInDays: to.Ptr[int32](90),
		// 						StatusMessage: to.Ptr("ManagedHsm is functional."),
		// 						TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					},
		// 			}},
		// 		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/examples/DeletedManagedHsm_List.json
func ExampleManagedHsmsClient_NewListDeletedPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedHsmsClient().NewListDeletedPager(nil)
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
		// page.DeletedManagedHsmListResult = armkeyvault.DeletedManagedHsmListResult{
		// 	Value: []*armkeyvault.DeletedManagedHsm{
		// 		{
		// 			Name: to.Ptr("vault-agile-drawer-6404"),
		// 			Type: to.Ptr("Microsoft.KeyVault/deletedManagedHSMs"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.KeyVault/locations/westus/deletedManagedHSMs/hsm1"),
		// 			Properties: &armkeyvault.DeletedManagedHsmProperties{
		// 				DeletionDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T00:00:59Z"); return t}()),
		// 				Location: to.Ptr("westus"),
		// 				MhsmID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.KeyVault/managedHSMs/hsm1"),
		// 				PurgeProtectionEnabled: to.Ptr(true),
		// 				ScheduledPurgeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T00:00:59Z"); return t}()),
		// 				Tags: map[string]*string{
		// 					"Dept": to.Ptr("hsm"),
		// 					"Environment": to.Ptr("production"),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("vault-agile-drawer-6404"),
		// 			Type: to.Ptr("Microsoft.KeyVault/deletedManagedHSMs"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.KeyVault/locations/westus/deletedManagedHSMs/hsm2"),
		// 			Properties: &armkeyvault.DeletedManagedHsmProperties{
		// 				DeletionDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T00:00:59Z"); return t}()),
		// 				Location: to.Ptr("westus"),
		// 				MhsmID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.KeyVault/managedHSMs/hsm2"),
		// 				PurgeProtectionEnabled: to.Ptr(true),
		// 				ScheduledPurgeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T00:00:59Z"); return t}()),
		// 				Tags: map[string]*string{
		// 					"Dept": to.Ptr("hsm"),
		// 					"Environment": to.Ptr("production"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/examples/DeletedManagedHsm_Get.json
func ExampleManagedHsmsClient_GetDeleted() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedHsmsClient().GetDeleted(ctx, "hsm1", "westus", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DeletedManagedHsm = armkeyvault.DeletedManagedHsm{
	// 	Name: to.Ptr("vault-agile-drawer-6404"),
	// 	Type: to.Ptr("Microsoft.KeyVault/deletedManagedHSMs"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.KeyVault/locations/westus/deletedManagedHSMs/hsm1"),
	// 	Properties: &armkeyvault.DeletedManagedHsmProperties{
	// 		DeletionDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T00:00:59Z"); return t}()),
	// 		Location: to.Ptr("westus"),
	// 		MhsmID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.KeyVault/managedHSMs/hsm1"),
	// 		PurgeProtectionEnabled: to.Ptr(true),
	// 		ScheduledPurgeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T00:00:59Z"); return t}()),
	// 		Tags: map[string]*string{
	// 			"Dept": to.Ptr("hsm"),
	// 			"Environment": to.Ptr("production"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/examples/DeletedManagedHsm_Purge.json
func ExampleManagedHsmsClient_BeginPurgeDeleted() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedHsmsClient().BeginPurgeDeleted(ctx, "hsm1", "westus", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/examples/ManagedHsm_checkMhsmNameAvailability.json
func ExampleManagedHsmsClient_CheckMhsmNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedHsmsClient().CheckMhsmNameAvailability(ctx, armkeyvault.CheckMhsmNameAvailabilityParameters{
		Name: to.Ptr("sample-mhsm"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckMhsmNameAvailabilityResult = armkeyvault.CheckMhsmNameAvailabilityResult{
	// 	NameAvailable: to.Ptr(true),
	// }
}
