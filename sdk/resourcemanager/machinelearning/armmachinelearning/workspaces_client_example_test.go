//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Workspace/get.json
func ExampleWorkspacesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewWorkspacesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "workspace-1234", "testworkspace", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Workspace/create.json
func ExampleWorkspacesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewWorkspacesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "workspace-1234", "testworkspace", armmachinelearning.Workspace{
		Identity: &armmachinelearning.ManagedServiceIdentity{
			Type: to.Ptr(armmachinelearning.ManagedServiceIdentityTypeSystemAssignedUserAssigned),
			UserAssignedIdentities: map[string]*armmachinelearning.UserAssignedIdentity{
				"/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testuai": {},
			},
		},
		Location: to.Ptr("eastus2euap"),
		Properties: &armmachinelearning.WorkspaceProperties{
			Description:         to.Ptr("test description"),
			ApplicationInsights: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/microsoft.insights/components/testinsights"),
			ContainerRegistry:   to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.ContainerRegistry/registries/testRegistry"),
			Encryption: &armmachinelearning.EncryptionProperty{
				Identity: &armmachinelearning.IdentityForCmk{
					UserAssignedIdentity: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testuai"),
				},
				KeyVaultProperties: &armmachinelearning.EncryptionKeyVaultProperties{
					IdentityClientID: to.Ptr(""),
					KeyIdentifier:    to.Ptr("https://testkv.vault.azure.net/keys/testkey/aabbccddee112233445566778899aabb"),
					KeyVaultArmID:    to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.KeyVault/vaults/testkv"),
				},
				Status: to.Ptr(armmachinelearning.EncryptionStatusEnabled),
			},
			FriendlyName: to.Ptr("HelloName"),
			HbiWorkspace: to.Ptr(false),
			KeyVault:     to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.KeyVault/vaults/testkv"),
			SharedPrivateLinkResources: []*armmachinelearning.SharedPrivateLinkResource{
				{
					Name: to.Ptr("testdbresource"),
					Properties: &armmachinelearning.SharedPrivateLinkResourceProperty{
						GroupID:               to.Ptr("Sql"),
						PrivateLinkResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.DocumentDB/databaseAccounts/testdbresource/privateLinkResources/Sql"),
						RequestMessage:        to.Ptr("Please approve"),
						Status:                to.Ptr(armmachinelearning.PrivateEndpointServiceConnectionStatusApproved),
					},
				}},
			StorageAccount: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/accountcrud-1234/providers/Microsoft.Storage/storageAccounts/testStorageAccount"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Workspace/delete.json
func ExampleWorkspacesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewWorkspacesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx, "workspace-1234", "testworkspace", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Workspace/update.json
func ExampleWorkspacesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewWorkspacesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx, "workspace-1234", "testworkspace", armmachinelearning.WorkspaceUpdateParameters{
		Properties: &armmachinelearning.WorkspacePropertiesUpdateParameters{
			Description:         to.Ptr("new description"),
			FriendlyName:        to.Ptr("New friendly name"),
			PublicNetworkAccess: to.Ptr(armmachinelearning.PublicNetworkAccessDisabled),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Workspace/listByResourceGroup.json
func ExampleWorkspacesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewWorkspacesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByResourceGroupPager("workspace-1234", &armmachinelearning.WorkspacesClientListByResourceGroupOptions{Skip: nil})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Workspace/diagnose.json
func ExampleWorkspacesClient_BeginDiagnose() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewWorkspacesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDiagnose(ctx, "workspace-1234", "testworkspace", &armmachinelearning.WorkspacesClientBeginDiagnoseOptions{Parameters: &armmachinelearning.DiagnoseWorkspaceParameters{
		Value: &armmachinelearning.DiagnoseRequestProperties{
			ApplicationInsights: map[string]interface{}{},
			ContainerRegistry:   map[string]interface{}{},
			DNSResolution:       map[string]interface{}{},
			KeyVault:            map[string]interface{}{},
			Nsg:                 map[string]interface{}{},
			Others:              map[string]interface{}{},
			ResourceLock:        map[string]interface{}{},
			StorageAccount:      map[string]interface{}{},
			Udr:                 map[string]interface{}{},
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Workspace/listKeys.json
func ExampleWorkspacesClient_ListKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewWorkspacesClient("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ListKeys(ctx, "testrg123", "workspaces123", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Workspace/resyncKeys.json
func ExampleWorkspacesClient_BeginResyncKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewWorkspacesClient("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginResyncKeys(ctx, "testrg123", "workspaces123", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Workspace/listBySubscription.json
func ExampleWorkspacesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewWorkspacesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListBySubscriptionPager(&armmachinelearning.WorkspacesClientListBySubscriptionOptions{Skip: nil})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Workspace/listNotebookAccessToken.json
func ExampleWorkspacesClient_ListNotebookAccessToken() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewWorkspacesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ListNotebookAccessToken(ctx, "workspace-1234", "testworkspace", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Notebook/prepare.json
func ExampleWorkspacesClient_BeginPrepareNotebook() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewWorkspacesClient("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginPrepareNotebook(ctx, "testrg123", "workspaces123", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Workspace/listStorageAccountKeys.json
func ExampleWorkspacesClient_ListStorageAccountKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewWorkspacesClient("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ListStorageAccountKeys(ctx, "testrg123", "workspaces123", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Notebook/listKeys.json
func ExampleWorkspacesClient_ListNotebookKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewWorkspacesClient("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ListNotebookKeys(ctx, "testrg123", "workspaces123", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/ExternalFQDN/get.json
func ExampleWorkspacesClient_ListOutboundNetworkDependenciesEndpoints() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewWorkspacesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ListOutboundNetworkDependenciesEndpoints(ctx, "workspace-1234", "testworkspace", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
