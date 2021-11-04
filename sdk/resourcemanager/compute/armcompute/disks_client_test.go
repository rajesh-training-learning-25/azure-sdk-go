//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package armcompute_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/recording"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestDisksClient_CreateOrUpdate(t *testing.T) {
	stop := startTest(t)
	defer stop()

	cred, opt := authenticateTest(t)
	subscriptionID := recording.GetEnvVariable("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	tenantID := recording.GetEnvVariable("AZURE_TENANT_ID", "00000000-0000-0000-0000-000000000000")
	ctx := context.Background()

	// create resource group
	rg,clean := createResourceGroup(t,cred,opt,subscriptionID,"disk","westus")
	rgName := *rg.Name
	defer clean()

	// create vault
	vClient := armkeyvault.NewVaultsClient(subscriptionID,cred,opt)
	vName, err := createRandomName(t, "vaultX")
	require.NoError(t, err)
	vPoller, err := vClient.BeginCreateOrUpdate(
		ctx,
		rgName,
		vName,
		armkeyvault.VaultCreateOrUpdateParameters{
			Location: to.StringPtr("westus"),
			Properties: &armkeyvault.VaultProperties{
				SKU: &armkeyvault.SKU{
					Family: armkeyvault.SKUFamilyA.ToPtr(),
					Name:   armkeyvault.SKUNameStandard.ToPtr(),
				},
				TenantID: to.StringPtr(tenantID),
				AccessPolicies: []*armkeyvault.AccessPolicyEntry{
					{
						TenantID: to.StringPtr(tenantID),
						ObjectID: to.StringPtr("00000000-0000-0000-0000-000000000000"),
						Permissions: &armkeyvault.Permissions{
							Keys: []*armkeyvault.KeyPermissions{
								armkeyvault.KeyPermissionsGet.ToPtr(),
								armkeyvault.KeyPermissionsList.ToPtr(),
								armkeyvault.KeyPermissionsCreate.ToPtr(),
							},
							Secrets: []*armkeyvault.SecretPermissions{
								armkeyvault.SecretPermissionsGet.ToPtr(),
								armkeyvault.SecretPermissionsList.ToPtr(),
							},
							Certificates: []*armkeyvault.CertificatePermissions{
								armkeyvault.CertificatePermissionsGet.ToPtr(),
								armkeyvault.CertificatePermissionsList.ToPtr(),
								armkeyvault.CertificatePermissionsCreate.ToPtr(),
							},
						},
					},
				},
			},
		},
		nil,
	)
	require.NoError(t, err)
	vResp, err := vPoller.PollUntilDone(ctx, 10*time.Second)
	require.NoError(t, err)
	require.Equal(t, *vResp.Name, vName)

	// create key
	keyClient := armkeyvault.NewKeysClient(subscriptionID,cred,nil)
	keyName, err := createRandomName(t, "key")
	require.NoError(t, err)
	keyResp, err := keyClient.CreateIfNotExist(
		ctx,
		rgName,
		vName,
		keyName,
		armkeyvault.KeyCreateParameters{
			Properties: &armkeyvault.KeyProperties{
				Attributes: &armkeyvault.KeyAttributes{
					Enabled: to.BoolPtr(true),
				},
				KeySize: to.Int32Ptr(2048),
				KeyOps: []*armkeyvault.JSONWebKeyOperation{
					armkeyvault.JSONWebKeyOperationEncrypt.ToPtr(),
					armkeyvault.JSONWebKeyOperationDecrypt.ToPtr(),
				},
				Kty: armkeyvault.JSONWebKeyTypeRSA.ToPtr(),
			},
		},
		nil,
	)
	require.NoError(t, err)
	require.Equal(t, *keyResp.Name, keyName)

	// create disk
	diskClient := armcompute.NewDisksClient(subscriptionID,cred,nil)
	diskName, err := createRandomName(t, "disk")
	require.NoError(t, err)
	diskPoller, err := diskClient.BeginCreateOrUpdate(
		ctx,
		rgName,
		diskName,
		armcompute.Disk{
			Resource: armcompute.Resource{
				Location: to.StringPtr("westus"),
			},
			SKU: &armcompute.DiskSKU{
				Name: armcompute.DiskStorageAccountTypesStandardLRS.ToPtr(),
			},
			Properties: &armcompute.DiskProperties{
				CreationData: &armcompute.CreationData{
					CreateOption: armcompute.DiskCreateOptionEmpty.ToPtr(),
				},
				DiskSizeGB: to.Int32Ptr(64),
			},
		},
		nil,
	)
	require.NoError(t, err)
	diskResp, err := diskPoller.PollUntilDone(ctx, 10*time.Second)
	require.NoError(t, err)
	require.Equal(t, *diskResp.Name, diskName)

	// create disk encryption set
	desClient := armcompute.NewDiskEncryptionSetsClient(subscriptionID,cred,nil)
	desName, err := createRandomName(t, "set")
	require.NoError(t, err)
	desPoller, err := desClient.BeginCreateOrUpdate(
		context.Background(),
		rgName,
		desName,
		armcompute.DiskEncryptionSet{
			Resource: armcompute.Resource{
				Location: to.StringPtr("westus"),
			},
			Identity: &armcompute.EncryptionSetIdentity{
				Type: armcompute.DiskEncryptionSetIdentityTypeSystemAssigned.ToPtr(),
			},
			Properties: &armcompute.EncryptionSetProperties{
				ActiveKey: &armcompute.KeyForDiskEncryptionSet{
					SourceVault: &armcompute.SourceVault{
						ID: to.StringPtr(*vResp.ID),
					},
					KeyURL: to.StringPtr(*keyResp.Properties.KeyURIWithVersion),
				},
			},
		},
		nil,
	)
	require.NoError(t, err)
	desResp, err := desPoller.PollUntilDone(context.Background(), 10*time.Second)
	require.NoError(t, err)
	require.Equal(t, *desResp.Name, desName)
}
