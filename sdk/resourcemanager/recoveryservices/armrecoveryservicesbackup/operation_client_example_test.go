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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3751704f5318f1175875c94b66af769db917f2d3/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-01-01/examples/AzureIaasVm/ValidateOperation_RestoreDisk.json
func ExampleOperationClient_Validate_validateOperation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicesbackup.NewOperationClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Validate(ctx, "testVault", "testRG", &armrecoveryservicesbackup.ValidateIaasVMRestoreOperationRequest{
		ObjectType: to.Ptr("ValidateIaasVMRestoreOperationRequest"),
		RestoreRequest: &armrecoveryservicesbackup.IaasVMRestoreRequest{
			ObjectType:            to.Ptr("IaasVMRestoreRequest"),
			CreateNewCloudService: to.Ptr(true),
			EncryptionDetails: &armrecoveryservicesbackup.EncryptionDetails{
				EncryptionEnabled: to.Ptr(false),
			},
			IdentityInfo: &armrecoveryservicesbackup.IdentityInfo{
				IsSystemAssignedIdentity:  to.Ptr(false),
				ManagedIdentityResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/asmaskarRG1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/asmaskartestmsi"),
			},
			OriginalStorageAccountOption: to.Ptr(false),
			RecoveryPointID:              to.Ptr("348916168024334"),
			RecoveryType:                 to.Ptr(armrecoveryservicesbackup.RecoveryTypeRestoreDisks),
			Region:                       to.Ptr("southeastasia"),
			SourceResourceID:             to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg/providers/Microsoft.Compute/virtualMachines/netvmtestv2vm1"),
			StorageAccountID:             to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testingRg/providers/Microsoft.Storage/storageAccounts/testAccount"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ValidateOperationsResponse = armrecoveryservicesbackup.ValidateOperationsResponse{
	// 	ValidateOperationResponse: &armrecoveryservicesbackup.ValidateOperationResponse{
	// 		ValidationResults: []*armrecoveryservicesbackup.ErrorDetail{
	// 			{
	// 				Code: to.Ptr("UserErrorCoreCountSubscriptionQuotaReached"),
	// 				Message: to.Ptr("Core Count subscription quota has been reached."),
	// 				Recommendations: []*string{
	// 					to.Ptr("Contact Azure support to increase the limits.")},
	// 			}},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3751704f5318f1175875c94b66af769db917f2d3/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-01-01/examples/AzureIaasVm/ValidateOperation_RestoreDisk_IdentityBasedRestoreDetails.json
func ExampleOperationClient_Validate_validateOperationWithIdentityBasedRestoreDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicesbackup.NewOperationClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Validate(ctx, "testVault", "testRG", &armrecoveryservicesbackup.ValidateIaasVMRestoreOperationRequest{
		ObjectType: to.Ptr("ValidateIaasVMRestoreOperationRequest"),
		RestoreRequest: &armrecoveryservicesbackup.IaasVMRestoreRequest{
			ObjectType:            to.Ptr("IaasVMRestoreRequest"),
			CreateNewCloudService: to.Ptr(true),
			EncryptionDetails: &armrecoveryservicesbackup.EncryptionDetails{
				EncryptionEnabled: to.Ptr(false),
			},
			IdentityBasedRestoreDetails: &armrecoveryservicesbackup.IdentityBasedRestoreDetails{
				TargetStorageAccountID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testingRg/providers/Microsoft.Storage/storageAccounts/testAccount"),
			},
			IdentityInfo: &armrecoveryservicesbackup.IdentityInfo{
				IsSystemAssignedIdentity:  to.Ptr(false),
				ManagedIdentityResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/asmaskarRG1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/asmaskartestmsi"),
			},
			OriginalStorageAccountOption: to.Ptr(false),
			RecoveryPointID:              to.Ptr("348916168024334"),
			RecoveryType:                 to.Ptr(armrecoveryservicesbackup.RecoveryTypeRestoreDisks),
			Region:                       to.Ptr("southeastasia"),
			SourceResourceID:             to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg/providers/Microsoft.Compute/virtualMachines/netvmtestv2vm1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ValidateOperationsResponse = armrecoveryservicesbackup.ValidateOperationsResponse{
	// 	ValidateOperationResponse: &armrecoveryservicesbackup.ValidateOperationResponse{
	// 		ValidationResults: []*armrecoveryservicesbackup.ErrorDetail{
	// 			{
	// 				Code: to.Ptr("UserErrorCoreCountSubscriptionQuotaReached"),
	// 				Message: to.Ptr("Core Count subscription quota has been reached."),
	// 				Recommendations: []*string{
	// 					to.Ptr("Contact Azure support to increase the limits.")},
	// 			}},
	// 		},
	// 	}
}
