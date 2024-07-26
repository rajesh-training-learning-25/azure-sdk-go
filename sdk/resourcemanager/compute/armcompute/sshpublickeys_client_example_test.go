//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/sshPublicKeyExamples/SshPublicKey_ListBySubscription_MaximumSet_Gen.json
func ExampleSSHPublicKeysClient_NewListBySubscriptionPager_sshPublicKeyListBySubscriptionMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSSHPublicKeysClient().NewListBySubscriptionPager(nil)
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
		// page.SSHPublicKeysGroupListResult = armcompute.SSHPublicKeysGroupListResult{
		// 	Value: []*armcompute.SSHPublicKeyResource{
		// 		{
		// 			Name: to.Ptr("mySshPublicKeyName"),
		// 			Type: to.Ptr("aaaa"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/sshPublicKeys/mySshPublicKeyName"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"key6396": to.Ptr("aaaaaaaaaaaaa"),
		// 				"key8839": to.Ptr("aaa"),
		// 			},
		// 			Properties: &armcompute.SSHPublicKeyResourceProperties{
		// 				PublicKey: to.Ptr("{ssh-rsa public key}"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/sshPublicKeyExamples/SshPublicKey_ListBySubscription_MinimumSet_Gen.json
func ExampleSSHPublicKeysClient_NewListBySubscriptionPager_sshPublicKeyListBySubscriptionMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSSHPublicKeysClient().NewListBySubscriptionPager(nil)
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
		// page.SSHPublicKeysGroupListResult = armcompute.SSHPublicKeysGroupListResult{
		// 	Value: []*armcompute.SSHPublicKeyResource{
		// 		{
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/rgcompute/providers/Microsoft.Compute/sshPublicKeys/mySshPublicKeyName"),
		// 			Location: to.Ptr("westus"),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/sshPublicKeyExamples/SshPublicKey_ListByResourceGroup_MaximumSet_Gen.json
func ExampleSSHPublicKeysClient_NewListByResourceGroupPager_sshPublicKeyListByResourceGroupMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSSHPublicKeysClient().NewListByResourceGroupPager("rgcompute", nil)
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
		// page.SSHPublicKeysGroupListResult = armcompute.SSHPublicKeysGroupListResult{
		// 	Value: []*armcompute.SSHPublicKeyResource{
		// 		{
		// 			Name: to.Ptr("mySshPublicKeyName"),
		// 			Type: to.Ptr("aaaa"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/sshPublicKeys/mySshPublicKeyName"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"key6396": to.Ptr("aaaaaaaaaaaaa"),
		// 				"key8839": to.Ptr("aaa"),
		// 			},
		// 			Properties: &armcompute.SSHPublicKeyResourceProperties{
		// 				PublicKey: to.Ptr("{ssh-rsa public key}"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/sshPublicKeyExamples/SshPublicKey_ListByResourceGroup_MinimumSet_Gen.json
func ExampleSSHPublicKeysClient_NewListByResourceGroupPager_sshPublicKeyListByResourceGroupMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSSHPublicKeysClient().NewListByResourceGroupPager("rgcompute", nil)
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
		// page.SSHPublicKeysGroupListResult = armcompute.SSHPublicKeysGroupListResult{
		// 	Value: []*armcompute.SSHPublicKeyResource{
		// 		{
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/rgcompute/providers/Microsoft.Compute/sshPublicKeys/mySshPublicKeyName"),
		// 			Location: to.Ptr("westus"),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/sshPublicKeyExamples/SshPublicKey_Create.json
func ExampleSSHPublicKeysClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSSHPublicKeysClient().Create(ctx, "myResourceGroup", "mySshPublicKeyName", armcompute.SSHPublicKeyResource{
		Location: to.Ptr("westus"),
		Properties: &armcompute.SSHPublicKeyResourceProperties{
			PublicKey: to.Ptr("{ssh-rsa public key}"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SSHPublicKeyResource = armcompute.SSHPublicKeyResource{
	// 	Name: to.Ptr("mySshPublicKeyName"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/sshPublicKeys/mySshPublicKeyName"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armcompute.SSHPublicKeyResourceProperties{
	// 		PublicKey: to.Ptr("{ssh-rsa public key}"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/sshPublicKeyExamples/SshPublicKey_Update_MaximumSet_Gen.json
func ExampleSSHPublicKeysClient_Update_sshPublicKeyUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSSHPublicKeysClient().Update(ctx, "rgcompute", "aaaaaaaaaaaa", armcompute.SSHPublicKeyUpdateResource{
		Tags: map[string]*string{
			"key2854": to.Ptr("a"),
		},
		Properties: &armcompute.SSHPublicKeyResourceProperties{
			PublicKey: to.Ptr("{ssh-rsa public key}"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SSHPublicKeyResource = armcompute.SSHPublicKeyResource{
	// 	Name: to.Ptr("mySshPublicKeyName"),
	// 	Type: to.Ptr("aaaa"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/sshPublicKeys/mySshPublicKeyName"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"key6396": to.Ptr("aaaaaaaaaaaaa"),
	// 		"key8839": to.Ptr("aaa"),
	// 	},
	// 	Properties: &armcompute.SSHPublicKeyResourceProperties{
	// 		PublicKey: to.Ptr("{ssh-rsa public key}"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/sshPublicKeyExamples/SshPublicKey_Update_MinimumSet_Gen.json
func ExampleSSHPublicKeysClient_Update_sshPublicKeyUpdateMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSSHPublicKeysClient().Update(ctx, "rgcompute", "aaaaaaaaaaa", armcompute.SSHPublicKeyUpdateResource{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SSHPublicKeyResource = armcompute.SSHPublicKeyResource{
	// 	Location: to.Ptr("westus"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/sshPublicKeyExamples/SshPublicKey_Delete_MaximumSet_Gen.json
func ExampleSSHPublicKeysClient_Delete_sshPublicKeyDeleteMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewSSHPublicKeysClient().Delete(ctx, "rgcompute", "aaaaaaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/sshPublicKeyExamples/SshPublicKey_Delete_MinimumSet_Gen.json
func ExampleSSHPublicKeysClient_Delete_sshPublicKeyDeleteMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewSSHPublicKeysClient().Delete(ctx, "rgcompute", "aaaaaaaaaaaaaaaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/sshPublicKeyExamples/SshPublicKey_Get.json
func ExampleSSHPublicKeysClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSSHPublicKeysClient().Get(ctx, "myResourceGroup", "mySshPublicKeyName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SSHPublicKeyResource = armcompute.SSHPublicKeyResource{
	// 	Name: to.Ptr("mySshPublicKeyName"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/SshPublicKeys/mySshPublicKeyName"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"{tagName}": to.Ptr("{tagValue}"),
	// 	},
	// 	Properties: &armcompute.SSHPublicKeyResourceProperties{
	// 		PublicKey: to.Ptr("{ssh-rsa public key}"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/sshPublicKeyExamples/SshPublicKey_GenerateKeyPair_EncryptionWithEd25519.json
func ExampleSSHPublicKeysClient_GenerateKeyPair_generateAnSshKeyPairWithEd25519Encryption() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSSHPublicKeysClient().GenerateKeyPair(ctx, "myResourceGroup", "mySshPublicKeyName", &armcompute.SSHPublicKeysClientGenerateKeyPairOptions{Parameters: &armcompute.SSHGenerateKeyPairInputParameters{
		EncryptionType: to.Ptr(armcompute.SSHEncryptionTypesRSA),
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SSHPublicKeyGenerateKeyPairResult = armcompute.SSHPublicKeyGenerateKeyPairResult{
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/SshPublicKeys/mySshPublicKeyName"),
	// 	PrivateKey: to.Ptr("{ssh private key}"),
	// 	PublicKey: to.Ptr("{ssh-rsa public key}"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/sshPublicKeyExamples/SshPublicKey_GenerateKeyPair_EncryptionWithRSA.json
func ExampleSSHPublicKeysClient_GenerateKeyPair_generateAnSshKeyPairWithRsaEncryption() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSSHPublicKeysClient().GenerateKeyPair(ctx, "myResourceGroup", "mySshPublicKeyName", &armcompute.SSHPublicKeysClientGenerateKeyPairOptions{Parameters: &armcompute.SSHGenerateKeyPairInputParameters{
		EncryptionType: to.Ptr(armcompute.SSHEncryptionTypesRSA),
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SSHPublicKeyGenerateKeyPairResult = armcompute.SSHPublicKeyGenerateKeyPairResult{
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/SshPublicKeys/mySshPublicKeyName"),
	// 	PrivateKey: to.Ptr("{ed25519 private key}"),
	// 	PublicKey: to.Ptr("{ed25519 public key}"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/sshPublicKeyExamples/SshPublicKey_GenerateKeyPair.json
func ExampleSSHPublicKeysClient_GenerateKeyPair_generateAnSshKeyPair() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSSHPublicKeysClient().GenerateKeyPair(ctx, "myResourceGroup", "mySshPublicKeyName", &armcompute.SSHPublicKeysClientGenerateKeyPairOptions{Parameters: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SSHPublicKeyGenerateKeyPairResult = armcompute.SSHPublicKeyGenerateKeyPairResult{
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/SshPublicKeys/mySshPublicKeyName"),
	// 	PrivateKey: to.Ptr("{ssh private key}"),
	// 	PublicKey: to.Ptr("{ssh-rsa public key}"),
	// }
}
