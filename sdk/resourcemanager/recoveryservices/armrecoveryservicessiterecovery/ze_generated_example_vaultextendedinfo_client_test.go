//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

// x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2021-08-01/examples/GETVaultExtendedInfo.json
func ExampleVaultExtendedInfoClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewVaultExtendedInfoClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<vault-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("VaultExtendedInfoResource.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2021-08-01/examples/UpdateVaultExtendedInfo.json
func ExampleVaultExtendedInfoClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewVaultExtendedInfoClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<vault-name>",
		armrecoveryservicessiterecovery.VaultExtendedInfoResource{},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("VaultExtendedInfoResource.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2021-08-01/examples/UpdateVaultExtendedInfo.json
func ExampleVaultExtendedInfoClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewVaultExtendedInfoClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<vault-name>",
		armrecoveryservicessiterecovery.VaultExtendedInfoResource{},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("VaultExtendedInfoResource.ID: %s\n", *res.ID)
}
