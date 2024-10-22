//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armredisenterprise_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a287afb3721dee0d88f11502ec123470bc52a28/specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2024-09-01-preview/examples/RedisEnterpriseAccessPolicyAssignmentCreateUpdate.json
func ExampleAccessPolicyAssignmentClient_BeginCreateUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredisenterprise.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccessPolicyAssignmentClient().BeginCreateUpdate(ctx, "rg1", "cache1", "default", "defaultTestEntraApp1", armredisenterprise.AccessPolicyAssignment{
		Properties: &armredisenterprise.AccessPolicyAssignmentProperties{
			AccessPolicyName: to.Ptr("default"),
			User: &armredisenterprise.AccessPolicyAssignmentPropertiesUser{
				ObjectID: to.Ptr("6497c918-11ad-41e7-1b0f-7c518a87d0b0"),
			},
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
	// res.AccessPolicyAssignment = armredisenterprise.AccessPolicyAssignment{
	// 	Name: to.Ptr("defaultTestEntraApp1"),
	// 	Type: to.Ptr("Microsoft.Cache/redisEnterprise/accessPolicyAssignments"),
	// 	ID: to.Ptr("subscriptions/e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f/resourceGroupsrg1/providers/Microsoft.Cache/redisEnterprise/cache1/databases/default/accessPolicyAssignments/default-TestEntraApp1"),
	// 	Properties: &armredisenterprise.AccessPolicyAssignmentProperties{
	// 		AccessPolicyName: to.Ptr("default"),
	// 		ProvisioningState: to.Ptr(armredisenterprise.ProvisioningStateSucceeded),
	// 		User: &armredisenterprise.AccessPolicyAssignmentPropertiesUser{
	// 			ObjectID: to.Ptr("6497c918-11ad-41e7-1b0f-7c518a87d0b0"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a287afb3721dee0d88f11502ec123470bc52a28/specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2024-09-01-preview/examples/RedisEnterpriseAccessPolicyAssignmentGet.json
func ExampleAccessPolicyAssignmentClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredisenterprise.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccessPolicyAssignmentClient().Get(ctx, "rg1", "cache1", "default", "accessPolicyAssignmentName1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccessPolicyAssignment = armredisenterprise.AccessPolicyAssignment{
	// 	Name: to.Ptr("accessPolicyAssignmentName1"),
	// 	Type: to.Ptr("Microsoft.Cache/Redis/accessPolicyAssignments"),
	// 	ID: to.Ptr("/subscriptions/e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f/resourceGroups/rg1/providers/Microsoft.Cache/redis/cache1/accessPolicyAssignments/accessPolicyAssignmentName1"),
	// 	Properties: &armredisenterprise.AccessPolicyAssignmentProperties{
	// 		AccessPolicyName: to.Ptr("default"),
	// 		ProvisioningState: to.Ptr(armredisenterprise.ProvisioningStateSucceeded),
	// 		User: &armredisenterprise.AccessPolicyAssignmentPropertiesUser{
	// 			ObjectID: to.Ptr("6497c918-11ad-41e7-1b0f-7c518a87d0b0"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a287afb3721dee0d88f11502ec123470bc52a28/specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2024-09-01-preview/examples/RedisEnterpriseAccessPolicyAssignmentDelete.json
func ExampleAccessPolicyAssignmentClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredisenterprise.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccessPolicyAssignmentClient().BeginDelete(ctx, "rg1", "cache1", "default", "defaultTestEntraApp1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a287afb3721dee0d88f11502ec123470bc52a28/specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2024-09-01-preview/examples/RedisEnterpriseAccessPolicyAssignmentsList.json
func ExampleAccessPolicyAssignmentClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredisenterprise.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccessPolicyAssignmentClient().NewListPager("rg1", "cache1", "default", nil)
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
		// page.AccessPolicyAssignmentList = armredisenterprise.AccessPolicyAssignmentList{
		// 	Value: []*armredisenterprise.AccessPolicyAssignment{
		// 		{
		// 			Name: to.Ptr("accessPolicyAssignmentName1"),
		// 			Type: to.Ptr("Microsoft.Cache/Redis/accessPolicyAssignments"),
		// 			ID: to.Ptr("/subscriptions/e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f/resourceGroups/rg1/providers/Microsoft.Cache/redis/cache1/accessPolicyAssignments/accessPolicyAssignmentName1"),
		// 			Properties: &armredisenterprise.AccessPolicyAssignmentProperties{
		// 				AccessPolicyName: to.Ptr("default"),
		// 				ProvisioningState: to.Ptr(armredisenterprise.ProvisioningStateSucceeded),
		// 				User: &armredisenterprise.AccessPolicyAssignmentPropertiesUser{
		// 					ObjectID: to.Ptr("6497c918-11ad-41e7-1b0f-7c518a87d0b0"),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("accessPolicyAssignmentName2"),
		// 			Type: to.Ptr("Microsoft.Cache/Redis/accessPolicyAssignments"),
		// 			ID: to.Ptr("/subscriptions/e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f/resourceGroups/rg1/providers/Microsoft.Cache/redis/cache1/accessPolicyAssignments/accessPolicyAssignmentName2"),
		// 			Properties: &armredisenterprise.AccessPolicyAssignmentProperties{
		// 				AccessPolicyName: to.Ptr("default"),
		// 				ProvisioningState: to.Ptr(armredisenterprise.ProvisioningStateSucceeded),
		// 				User: &armredisenterprise.AccessPolicyAssignmentPropertiesUser{
		// 					ObjectID: to.Ptr("7497c918-11ad-41e7-1b0f-7c518a87d0b0"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}