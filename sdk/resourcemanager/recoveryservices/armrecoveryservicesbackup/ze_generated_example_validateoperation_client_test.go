//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
)

// x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-10-01/examples/AzureIaasVm/TriggerValidateOperation_RestoreDisk.json
func ExampleValidateOperationClient_BeginTrigger() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicesbackup.NewValidateOperationClient("<subscription-id>", cred, nil)
	poller, err := client.BeginTrigger(ctx,
		"<vault-name>",
		"<resource-group-name>",
		&armrecoveryservicesbackup.ValidateIaasVMRestoreOperationRequest{
			ObjectType: to.StringPtr("<object-type>"),
			RestoreRequest: &armrecoveryservicesbackup.IaasVMRestoreRequest{
				ObjectType:            to.StringPtr("<object-type>"),
				CreateNewCloudService: to.BoolPtr(true),
				EncryptionDetails: &armrecoveryservicesbackup.EncryptionDetails{
					EncryptionEnabled: to.BoolPtr(false),
				},
				IdentityInfo: &armrecoveryservicesbackup.IdentityInfo{
					IsSystemAssignedIdentity:  to.BoolPtr(false),
					ManagedIdentityResourceID: to.StringPtr("<managed-identity-resource-id>"),
				},
				OriginalStorageAccountOption: to.BoolPtr(false),
				RecoveryPointID:              to.StringPtr("<recovery-point-id>"),
				RecoveryType:                 armrecoveryservicesbackup.RecoveryType("RestoreDisks").ToPtr(),
				Region:                       to.StringPtr("<region>"),
				SourceResourceID:             to.StringPtr("<source-resource-id>"),
				StorageAccountID:             to.StringPtr("<storage-account-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
