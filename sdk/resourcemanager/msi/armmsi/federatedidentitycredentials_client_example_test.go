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

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/msi/armmsi"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5e24fa8e30bff0cb321494a0b550b1c1282a8a3c/specification/msi/resource-manager/Microsoft.ManagedIdentity/stable/2023-01-31/examples/FederatedIdentityCredentialList.json
func ExampleFederatedIdentityCredentialsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmsi.NewFederatedIdentityCredentialsClient("c267c0e7-0a73-4789-9e17-d26aeb0904e5", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("rgName", "resourceName", &armmsi.FederatedIdentityCredentialsClientListOptions{Top: nil,
		Skiptoken: nil,
	})
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
		// page.FederatedIdentityCredentialsListResult = armmsi.FederatedIdentityCredentialsListResult{
		// 	Value: []*armmsi.FederatedIdentityCredential{
		// 		{
		// 			Name: to.Ptr("ficResourceName"),
		// 			Type: to.Ptr("Microsoft.ManagedIdentity/userAssignedIdentities/federatedIdentityCredentials"),
		// 			ID: to.Ptr("/subscriptions/c267c0e7-0a73-4789-9e17-d26aeb0904e5/resourcegroups/rgName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identityName/federatedIdentityCredentials/ficResourceName"),
		// 			Properties: &armmsi.FederatedIdentityCredentialProperties{
		// 				Audiences: []*string{
		// 					to.Ptr("api://AzureADTokenExchange")},
		// 					Issuer: to.Ptr("https://oidc.prod-aks.azure.com/TenantGUID/IssuerGUID"),
		// 					Subject: to.Ptr("system:serviceaccount:ns:svcaccount"),
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5e24fa8e30bff0cb321494a0b550b1c1282a8a3c/specification/msi/resource-manager/Microsoft.ManagedIdentity/stable/2023-01-31/examples/FederatedIdentityCredentialCreate.json
func ExampleFederatedIdentityCredentialsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmsi.NewFederatedIdentityCredentialsClient("c267c0e7-0a73-4789-9e17-d26aeb0904e5", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "rgName", "resourceName", "ficResourceName", armmsi.FederatedIdentityCredential{
		Properties: &armmsi.FederatedIdentityCredentialProperties{
			Audiences: []*string{
				to.Ptr("api://AzureADTokenExchange")},
			Issuer:  to.Ptr("https://oidc.prod-aks.azure.com/TenantGUID/IssuerGUID"),
			Subject: to.Ptr("system:serviceaccount:ns:svcaccount"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FederatedIdentityCredential = armmsi.FederatedIdentityCredential{
	// 	Name: to.Ptr("ficResourceName"),
	// 	Type: to.Ptr("Microsoft.ManagedIdentity/userAssignedIdentities/federatedIdentityCredentials"),
	// 	ID: to.Ptr("/subscriptions/c267c0e7-0a73-4789-9e17-d26aeb0904e5/resourcegroups/rgName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identityName/federatedIdentityCredentials/ficResourceName"),
	// 	Properties: &armmsi.FederatedIdentityCredentialProperties{
	// 		Audiences: []*string{
	// 			to.Ptr("api://AzureADTokenExchange")},
	// 			Issuer: to.Ptr("https://oidc.prod-aks.azure.com/TenantGUID/IssuerGUID"),
	// 			Subject: to.Ptr("system:serviceaccount:ns:svcaccount"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5e24fa8e30bff0cb321494a0b550b1c1282a8a3c/specification/msi/resource-manager/Microsoft.ManagedIdentity/stable/2023-01-31/examples/FederatedIdentityCredentialGet.json
func ExampleFederatedIdentityCredentialsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmsi.NewFederatedIdentityCredentialsClient("c267c0e7-0a73-4789-9e17-d26aeb0904e5", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "rgName", "resourceName", "ficResourceName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FederatedIdentityCredential = armmsi.FederatedIdentityCredential{
	// 	Name: to.Ptr("ficResourceName"),
	// 	Type: to.Ptr("Microsoft.ManagedIdentity/userAssignedIdentities/federatedIdentityCredentials"),
	// 	ID: to.Ptr("/subscriptions/c267c0e7-0a73-4789-9e17-d26aeb0904e5/resourcegroups/rgName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identityName/federatedIdentityCredentials/ficResourceName"),
	// 	Properties: &armmsi.FederatedIdentityCredentialProperties{
	// 		Audiences: []*string{
	// 			to.Ptr("api://AzureADTokenExchange")},
	// 			Issuer: to.Ptr("https://oidc.prod-aks.azure.com/TenantGUID/IssuerGUID"),
	// 			Subject: to.Ptr("system:serviceaccount:ns:svcaccount"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5e24fa8e30bff0cb321494a0b550b1c1282a8a3c/specification/msi/resource-manager/Microsoft.ManagedIdentity/stable/2023-01-31/examples/FederatedIdentityCredentialDelete.json
func ExampleFederatedIdentityCredentialsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmsi.NewFederatedIdentityCredentialsClient("c267c0e7-0a73-4789-9e17-d26aeb0904e5", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx, "rgName", "resourceName", "ficResourceName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
