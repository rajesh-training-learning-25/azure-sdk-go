//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmsi_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/msi/armmsi"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5e24fa8e30bff0cb321494a0b550b1c1282a8a3c/specification/msi/resource-manager/Microsoft.ManagedIdentity/stable/2023-01-31/examples/MsiOperationsList.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmsi.NewOperationsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager(nil)
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
		// page.OperationListResult = armmsi.OperationListResult{
		// 	Value: []*armmsi.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.ManagedIdentity/userAssignedIdentities/read"),
		// 			Display: &armmsi.OperationDisplay{
		// 				Description: to.Ptr("Gets an existing user assigned identity"),
		// 				Operation: to.Ptr("Get User Assigned Identity"),
		// 				Provider: to.Ptr("Managed Service Identity"),
		// 				Resource: to.Ptr("User Assigned Identities"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ManagedIdentity/userAssignedIdentities/write"),
		// 			Display: &armmsi.OperationDisplay{
		// 				Description: to.Ptr("Creates a new user assigned identity or updates the tags associated with an existing user assigned identity"),
		// 				Operation: to.Ptr("Create/Update User Assigned Identity"),
		// 				Provider: to.Ptr("Managed Service Identity"),
		// 				Resource: to.Ptr("User Assigned Identities"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ManagedIdentity/userAssignedIdentities/delete"),
		// 			Display: &armmsi.OperationDisplay{
		// 				Description: to.Ptr("Deletes an existing user assigned identity"),
		// 				Operation: to.Ptr("Delete User Assigned Identity"),
		// 				Provider: to.Ptr("Managed Service Identity"),
		// 				Resource: to.Ptr("User Assigned Identities"),
		// 			},
		// 	}},
		// }
	}
}
