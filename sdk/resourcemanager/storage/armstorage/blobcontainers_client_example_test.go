//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/BlobContainersList.json
func ExampleBlobContainersClient_NewListPager_listContainers() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewBlobContainersClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("res9290", "sto1590", &armstorage.BlobContainersClientListOptions{Maxpagesize: nil,
		Filter:  nil,
		Include: nil,
	})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/DeletedBlobContainersList.json
func ExampleBlobContainersClient_NewListPager_listDeletedContainers() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewBlobContainersClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("res9290", "sto1590", &armstorage.BlobContainersClientListOptions{Maxpagesize: nil,
		Filter:  nil,
		Include: to.Ptr(armstorage.ListContainersIncludeDeleted),
	})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/BlobContainersPutDefaultEncryptionScope.json
func ExampleBlobContainersClient_Create_putContainerWithDefaultEncryptionScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewBlobContainersClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx, "res3376", "sto328", "container6185", armstorage.BlobContainer{
		ContainerProperties: &armstorage.ContainerProperties{
			DefaultEncryptionScope:      to.Ptr("encryptionscope185"),
			DenyEncryptionScopeOverride: to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/BlobContainersPutObjectLevelWorm.json
func ExampleBlobContainersClient_Create_putContainerWithObjectLevelWorm() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewBlobContainersClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx, "res3376", "sto328", "container6185", armstorage.BlobContainer{
		ContainerProperties: &armstorage.ContainerProperties{
			ImmutableStorageWithVersioning: &armstorage.ImmutableStorageWithVersioning{
				Enabled: to.Ptr(true),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/BlobContainersPut.json
func ExampleBlobContainersClient_Create_putContainers() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewBlobContainersClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx, "res3376", "sto328", "container6185", armstorage.BlobContainer{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/BlobContainersPatch.json
func ExampleBlobContainersClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewBlobContainersClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx, "res3376", "sto328", "container6185", armstorage.BlobContainer{
		ContainerProperties: &armstorage.ContainerProperties{
			Metadata: map[string]*string{
				"metadata": to.Ptr("true"),
			},
			PublicAccess: to.Ptr(armstorage.PublicAccessContainer),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/BlobContainersGetWithAllowProtectedAppendWritesAll.json
func ExampleBlobContainersClient_Get_getBlobContainersGetWithAllowProtectedAppendWritesAll() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewBlobContainersClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "res9871", "sto6217", "container1634", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/BlobContainersGet.json
func ExampleBlobContainersClient_Get_getContainers() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewBlobContainersClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "res9871", "sto6217", "container1634", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/BlobContainersDelete.json
func ExampleBlobContainersClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewBlobContainersClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx, "res4079", "sto4506", "container9689", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/BlobContainersSetLegalHold.json
func ExampleBlobContainersClient_SetLegalHold_setLegalHoldContainers() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewBlobContainersClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.SetLegalHold(ctx, "res4303", "sto7280", "container8723", armstorage.LegalHold{
		Tags: []*string{
			to.Ptr("tag1"),
			to.Ptr("tag2"),
			to.Ptr("tag3")},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/BlobContainersSetLegalHoldAllowProtectedAppendWritesAll.json
func ExampleBlobContainersClient_SetLegalHold_setLegalHoldContainersWithAllowProtectedAppendWritesAll() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewBlobContainersClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.SetLegalHold(ctx, "res4303", "sto7280", "container8723", armstorage.LegalHold{
		AllowProtectedAppendWritesAll: to.Ptr(true),
		Tags: []*string{
			to.Ptr("tag1"),
			to.Ptr("tag2"),
			to.Ptr("tag3")},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/BlobContainersClearLegalHold.json
func ExampleBlobContainersClient_ClearLegalHold() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewBlobContainersClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ClearLegalHold(ctx, "res4303", "sto7280", "container8723", armstorage.LegalHold{
		Tags: []*string{
			to.Ptr("tag1"),
			to.Ptr("tag2"),
			to.Ptr("tag3")},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/BlobContainersPutImmutabilityPolicy.json
func ExampleBlobContainersClient_CreateOrUpdateImmutabilityPolicy_createOrUpdateImmutabilityPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewBlobContainersClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdateImmutabilityPolicy(ctx, "res1782", "sto7069", "container6397", &armstorage.BlobContainersClientCreateOrUpdateImmutabilityPolicyOptions{IfMatch: nil,
		Parameters: &armstorage.ImmutabilityPolicy{
			Properties: &armstorage.ImmutabilityPolicyProperty{
				AllowProtectedAppendWrites:            to.Ptr(true),
				ImmutabilityPeriodSinceCreationInDays: to.Ptr[int32](3),
			},
		},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/BlobContainersPutImmutabilityPolicyAllowProtectedAppendWritesAll.json
func ExampleBlobContainersClient_CreateOrUpdateImmutabilityPolicy_createOrUpdateImmutabilityPolicyWithAllowProtectedAppendWritesAll() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewBlobContainersClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdateImmutabilityPolicy(ctx, "res1782", "sto7069", "container6397", &armstorage.BlobContainersClientCreateOrUpdateImmutabilityPolicyOptions{IfMatch: nil,
		Parameters: &armstorage.ImmutabilityPolicy{
			Properties: &armstorage.ImmutabilityPolicyProperty{
				AllowProtectedAppendWritesAll:         to.Ptr(true),
				ImmutabilityPeriodSinceCreationInDays: to.Ptr[int32](3),
			},
		},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/BlobContainersGetImmutabilityPolicy.json
func ExampleBlobContainersClient_GetImmutabilityPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewBlobContainersClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetImmutabilityPolicy(ctx, "res5221", "sto9177", "container3489", &armstorage.BlobContainersClientGetImmutabilityPolicyOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/BlobContainersDeleteImmutabilityPolicy.json
func ExampleBlobContainersClient_DeleteImmutabilityPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewBlobContainersClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.DeleteImmutabilityPolicy(ctx, "res1581", "sto9621", "container4910", "\"8d59f81a7fa7be0\"", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/BlobContainersLockImmutabilityPolicy.json
func ExampleBlobContainersClient_LockImmutabilityPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewBlobContainersClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.LockImmutabilityPolicy(ctx, "res2702", "sto5009", "container1631", "\"8d59f825b721dd3\"", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/BlobContainersExtendImmutabilityPolicy.json
func ExampleBlobContainersClient_ExtendImmutabilityPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewBlobContainersClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ExtendImmutabilityPolicy(ctx, "res6238", "sto232", "container5023", "\"8d59f830d0c3bf9\"", &armstorage.BlobContainersClientExtendImmutabilityPolicyOptions{Parameters: &armstorage.ImmutabilityPolicy{
		Properties: &armstorage.ImmutabilityPolicyProperty{
			ImmutabilityPeriodSinceCreationInDays: to.Ptr[int32](100),
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/BlobContainersLease_Acquire.json
func ExampleBlobContainersClient_Lease_acquireALeaseOnAContainer() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewBlobContainersClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Lease(ctx, "res3376", "sto328", "container6185", &armstorage.BlobContainersClientLeaseOptions{Parameters: &armstorage.LeaseContainerRequest{
		Action:        to.Ptr(armstorage.LeaseContainerRequestActionAcquire),
		LeaseDuration: to.Ptr[int32](-1),
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/BlobContainersLease_Break.json
func ExampleBlobContainersClient_Lease_breakALeaseOnAContainer() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewBlobContainersClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Lease(ctx, "res3376", "sto328", "container6185", &armstorage.BlobContainersClientLeaseOptions{Parameters: &armstorage.LeaseContainerRequest{
		Action:  to.Ptr(armstorage.LeaseContainerRequestActionBreak),
		LeaseID: to.Ptr("8698f513-fa75-44a1-b8eb-30ba336af27d"),
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/ObjectLevelWormContainerMigration.json
func ExampleBlobContainersClient_BeginObjectLevelWorm() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewBlobContainersClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginObjectLevelWorm(ctx, "res1782", "sto7069", "container6397", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
