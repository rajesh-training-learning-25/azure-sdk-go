//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/AzureWorkload/ProtectionContainers_Get.json
func ExampleProtectionContainersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProtectionContainersClient().Get(ctx, "testVault", "testRg", "Azure", "VMAppContainer;Compute;testRG;testSQL", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProtectionContainerResource = armrecoveryservicesbackup.ProtectionContainerResource{
	// 	Name: to.Ptr("VMAppContainer;Compute;testRG;testSQL"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupFabrics/protectionContainers"),
	// 	ID: to.Ptr("/Subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRg/providers/Microsoft.RecoveryServices/vaults/testVault/backupFabrics/Azure/protectionContainers/VMAppContainer;Compute;testRG;testSQL"),
	// 	Properties: &armrecoveryservicesbackup.AzureVMAppContainerProtectionContainer{
	// 		BackupManagementType: to.Ptr(armrecoveryservicesbackup.BackupManagementTypeAzureWorkload),
	// 		ContainerType: to.Ptr(armrecoveryservicesbackup.ProtectableContainerTypeVMAppContainer),
	// 		FriendlyName: to.Ptr("testSQL"),
	// 		ExtendedInfo: &armrecoveryservicesbackup.AzureWorkloadContainerExtendedInfo{
	// 			HostServerName: to.Ptr("testsql"),
	// 			InquiryInfo: &armrecoveryservicesbackup.InquiryInfo{
	// 				ErrorDetail: &armrecoveryservicesbackup.ErrorDetail{
	// 					Code: to.Ptr("Success"),
	// 					Message: to.Ptr("Not Available"),
	// 					Recommendations: []*string{
	// 						to.Ptr("Not Available")},
	// 					},
	// 					InquiryDetails: []*armrecoveryservicesbackup.WorkloadInquiryDetails{
	// 						{
	// 							Type: to.Ptr("sql"),
	// 							InquiryValidation: &armrecoveryservicesbackup.InquiryValidation{
	// 								ErrorDetail: &armrecoveryservicesbackup.ErrorDetail{
	// 									Code: to.Ptr("Success"),
	// 									Message: to.Ptr("Not Available"),
	// 									Recommendations: []*string{
	// 										to.Ptr("Not Available")},
	// 									},
	// 									Status: to.Ptr("Success"),
	// 								},
	// 								ItemCount: to.Ptr[int64](14),
	// 						}},
	// 						Status: to.Ptr("Success"),
	// 					},
	// 					NodesList: []*armrecoveryservicesbackup.DistributedNodesInfo{
	// 					},
	// 				},
	// 				SourceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.Compute/virtualMachines/testSQL"),
	// 			},
	// 		}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/AzureStorage/ProtectionContainers_Register.json
func ExampleProtectionContainersClient_BeginRegister() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewProtectionContainersClient().BeginRegister(ctx, "swaggertestvault", "SwaggerTestRg", "Azure", "StorageContainer;Storage;SwaggerTestRg;swaggertestsa", armrecoveryservicesbackup.ProtectionContainerResource{
		Properties: &armrecoveryservicesbackup.AzureStorageContainer{
			BackupManagementType:      to.Ptr(armrecoveryservicesbackup.BackupManagementTypeAzureStorage),
			ContainerType:             to.Ptr(armrecoveryservicesbackup.ProtectableContainerTypeStorageContainer),
			FriendlyName:              to.Ptr("swaggertestsa"),
			AcquireStorageAccountLock: to.Ptr(armrecoveryservicesbackup.AcquireStorageAccountLockAcquire),
			SourceResourceID:          to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/SwaggerTestRg/providers/Microsoft.Storage/storageAccounts/swaggertestsa"),
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
	// res.ProtectionContainerResource = armrecoveryservicesbackup.ProtectionContainerResource{
	// 	Name: to.Ptr("StorageContainer;Storage;SwaggerTestRg;swaggertestsa"),
	// 	ID: to.Ptr("/Subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/SwaggerTestRg/providers/Microsoft.RecoveryServices/vaults/swaggertestvault/backupFabrics/Azure/protectionContainers/StorageContainer;Storage;SwaggerTestRg;swaggertestsa"),
	// 	Properties: &armrecoveryservicesbackup.AzureStorageContainer{
	// 		BackupManagementType: to.Ptr(armrecoveryservicesbackup.BackupManagementTypeAzureStorage),
	// 		ContainerType: to.Ptr(armrecoveryservicesbackup.ProtectableContainerTypeStorageContainer),
	// 		FriendlyName: to.Ptr("swaggertestsa"),
	// 		HealthStatus: to.Ptr("Healthy"),
	// 		RegistrationStatus: to.Ptr("Registered"),
	// 		AcquireStorageAccountLock: to.Ptr(armrecoveryservicesbackup.AcquireStorageAccountLockAcquire),
	// 		ProtectedItemCount: to.Ptr[int64](0),
	// 		SourceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/SwaggerTestRg/providers/Microsoft.Storage/storageAccounts/swaggertestsa"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/AzureWorkload/ProtectionContainers_Unregister.json
func ExampleProtectionContainersClient_Unregister() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewProtectionContainersClient().Unregister(ctx, "testVault", "testRg", "Azure", "storagecontainer;Storage;test-rg;teststorage", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/AzureStorage/ProtectionContainers_Inquire.json
func ExampleProtectionContainersClient_Inquire() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewProtectionContainersClient().Inquire(ctx, "testvault", "test-rg", "Azure", "storagecontainer;Storage;test-rg;teststorage", &armrecoveryservicesbackup.ProtectionContainersClientInquireOptions{Filter: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/Common/RefreshContainers.json
func ExampleProtectionContainersClient_Refresh() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewProtectionContainersClient().Refresh(ctx, "NetSDKTestRsVault", "SwaggerTestRg", "Azure", &armrecoveryservicesbackup.ProtectionContainersClientRefreshOptions{Filter: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
