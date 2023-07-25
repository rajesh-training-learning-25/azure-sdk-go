//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edd7863aeaad7379a7a8eef424ca11617d8c91eb/specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/DatabaseExtendedAuditingSettingsList.json
func ExampleExtendedDatabaseBlobAuditingPoliciesClient_NewListByDatabasePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExtendedDatabaseBlobAuditingPoliciesClient().NewListByDatabasePager("blobauditingtest-6852", "blobauditingtest-2080", "testdb", nil)
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
		// page.ExtendedDatabaseBlobAuditingPolicyListResult = armsql.ExtendedDatabaseBlobAuditingPolicyListResult{
		// 	Value: []*armsql.ExtendedDatabaseBlobAuditingPolicy{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases/extendedAuditingSettings"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/blobauditingtest-6852/providers/Microsoft.Sql/servers/blobauditingtest-2080/databases/testdb/extendedAuditingSettings/default"),
		// 			Properties: &armsql.ExtendedDatabaseBlobAuditingPolicyProperties{
		// 				AuditActionsAndGroups: []*string{
		// 				},
		// 				IsAzureMonitorTargetEnabled: to.Ptr(false),
		// 				IsManagedIdentityInUse: to.Ptr(false),
		// 				IsStorageSecondaryKeyInUse: to.Ptr(false),
		// 				PredicateExpression: to.Ptr("statement = 'select 1'"),
		// 				RetentionDays: to.Ptr[int32](0),
		// 				State: to.Ptr(armsql.BlobAuditingPolicyStateDisabled),
		// 				StorageAccountSubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				StorageEndpoint: to.Ptr(""),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edd7863aeaad7379a7a8eef424ca11617d8c91eb/specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/ExtendedDatabaseBlobAuditingGet.json
func ExampleExtendedDatabaseBlobAuditingPoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExtendedDatabaseBlobAuditingPoliciesClient().Get(ctx, "blobauditingtest-6852", "blobauditingtest-2080", "testdb", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ExtendedDatabaseBlobAuditingPolicy = armsql.ExtendedDatabaseBlobAuditingPolicy{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/extendedAuditingSettings"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/blobauditingtest-6852/providers/Microsoft.Sql/servers/blobauditingtest-2080/databases/testdb"),
	// 	Properties: &armsql.ExtendedDatabaseBlobAuditingPolicyProperties{
	// 		AuditActionsAndGroups: []*string{
	// 		},
	// 		IsAzureMonitorTargetEnabled: to.Ptr(false),
	// 		IsManagedIdentityInUse: to.Ptr(false),
	// 		IsStorageSecondaryKeyInUse: to.Ptr(false),
	// 		PredicateExpression: to.Ptr("statement = 'select 1'"),
	// 		RetentionDays: to.Ptr[int32](0),
	// 		State: to.Ptr(armsql.BlobAuditingPolicyStateDisabled),
	// 		StorageAccountSubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		StorageEndpoint: to.Ptr(""),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edd7863aeaad7379a7a8eef424ca11617d8c91eb/specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/ExtendedDatabaseAzureMonitorAuditingCreateMin.json
func ExampleExtendedDatabaseBlobAuditingPoliciesClient_CreateOrUpdate_createOrUpdateAnExtendedDatabasesAzureMonitorAuditingPolicyWithMinimalParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExtendedDatabaseBlobAuditingPoliciesClient().CreateOrUpdate(ctx, "blobauditingtest-4799", "blobauditingtest-6440", "testdb", armsql.ExtendedDatabaseBlobAuditingPolicy{
		Properties: &armsql.ExtendedDatabaseBlobAuditingPolicyProperties{
			IsAzureMonitorTargetEnabled: to.Ptr(true),
			State:                       to.Ptr(armsql.BlobAuditingPolicyStateEnabled),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ExtendedDatabaseBlobAuditingPolicy = armsql.ExtendedDatabaseBlobAuditingPolicy{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/extendedAuditingSettings"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/blobauditingtest-4799/providers/Microsoft.Sql/servers/blobauditingtest-6440/databases/testdb"),
	// 	Properties: &armsql.ExtendedDatabaseBlobAuditingPolicyProperties{
	// 		AuditActionsAndGroups: []*string{
	// 			to.Ptr("SUCCESSFUL_DATABASE_AUTHENTICATION_GROUP"),
	// 			to.Ptr("FAILED_DATABASE_AUTHENTICATION_GROUP"),
	// 			to.Ptr("BATCH_COMPLETED_GROUP")},
	// 			IsAzureMonitorTargetEnabled: to.Ptr(true),
	// 			RetentionDays: to.Ptr[int32](0),
	// 			State: to.Ptr(armsql.BlobAuditingPolicyStateEnabled),
	// 			StorageAccountSubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edd7863aeaad7379a7a8eef424ca11617d8c91eb/specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/ExtendedDatabaseBlobAuditingCreateMax.json
func ExampleExtendedDatabaseBlobAuditingPoliciesClient_CreateOrUpdate_createOrUpdateAnExtendedDatabasesBlobAuditingPolicyWithAllParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExtendedDatabaseBlobAuditingPoliciesClient().CreateOrUpdate(ctx, "blobauditingtest-4799", "blobauditingtest-6440", "testdb", armsql.ExtendedDatabaseBlobAuditingPolicy{
		Properties: &armsql.ExtendedDatabaseBlobAuditingPolicyProperties{
			AuditActionsAndGroups: []*string{
				to.Ptr("DATABASE_LOGOUT_GROUP"),
				to.Ptr("DATABASE_ROLE_MEMBER_CHANGE_GROUP"),
				to.Ptr("UPDATE on database::TestDatabaseName by public")},
			IsAzureMonitorTargetEnabled:  to.Ptr(true),
			IsStorageSecondaryKeyInUse:   to.Ptr(false),
			PredicateExpression:          to.Ptr("statement = 'select 1'"),
			QueueDelayMs:                 to.Ptr[int32](4000),
			RetentionDays:                to.Ptr[int32](6),
			State:                        to.Ptr(armsql.BlobAuditingPolicyStateEnabled),
			StorageAccountAccessKey:      to.Ptr("sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD=="),
			StorageAccountSubscriptionID: to.Ptr("00000000-1234-0000-5678-000000000000"),
			StorageEndpoint:              to.Ptr("https://mystorage.blob.core.windows.net"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ExtendedDatabaseBlobAuditingPolicy = armsql.ExtendedDatabaseBlobAuditingPolicy{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/extendedAuditingSettings"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/blobauditingtest-4799/providers/Microsoft.Sql/servers/blobauditingtest-6440/databases/testdb"),
	// 	Properties: &armsql.ExtendedDatabaseBlobAuditingPolicyProperties{
	// 		AuditActionsAndGroups: []*string{
	// 			to.Ptr("DATABASE_LOGOUT_GROUP"),
	// 			to.Ptr("DATABASE_ROLE_MEMBER_CHANGE_GROUP"),
	// 			to.Ptr("UPDATE on database::TestDatabaseName by public")},
	// 			IsAzureMonitorTargetEnabled: to.Ptr(true),
	// 			IsStorageSecondaryKeyInUse: to.Ptr(false),
	// 			PredicateExpression: to.Ptr("statement = 'select 1'"),
	// 			QueueDelayMs: to.Ptr[int32](4000),
	// 			RetentionDays: to.Ptr[int32](0),
	// 			State: to.Ptr(armsql.BlobAuditingPolicyStateEnabled),
	// 			StorageAccountSubscriptionID: to.Ptr("00000000-1234-0000-5678-000000000000"),
	// 			StorageEndpoint: to.Ptr("https://mystorage.blob.core.windows.net"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edd7863aeaad7379a7a8eef424ca11617d8c91eb/specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/ExtendedDatabaseBlobAuditingCreateMin.json
func ExampleExtendedDatabaseBlobAuditingPoliciesClient_CreateOrUpdate_createOrUpdateAnExtendedDatabasesBlobAuditingPolicyWithMinimalParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExtendedDatabaseBlobAuditingPoliciesClient().CreateOrUpdate(ctx, "blobauditingtest-4799", "blobauditingtest-6440", "testdb", armsql.ExtendedDatabaseBlobAuditingPolicy{
		Properties: &armsql.ExtendedDatabaseBlobAuditingPolicyProperties{
			State:                   to.Ptr(armsql.BlobAuditingPolicyStateEnabled),
			StorageAccountAccessKey: to.Ptr("sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD=="),
			StorageEndpoint:         to.Ptr("https://mystorage.blob.core.windows.net"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ExtendedDatabaseBlobAuditingPolicy = armsql.ExtendedDatabaseBlobAuditingPolicy{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/extendedAuditingSettings"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/blobauditingtest-4799/providers/Microsoft.Sql/servers/blobauditingtest-6440/databases/testdb"),
	// 	Properties: &armsql.ExtendedDatabaseBlobAuditingPolicyProperties{
	// 		AuditActionsAndGroups: []*string{
	// 			to.Ptr("SUCCESSFUL_DATABASE_AUTHENTICATION_GROUP"),
	// 			to.Ptr("FAILED_DATABASE_AUTHENTICATION_GROUP"),
	// 			to.Ptr("BATCH_COMPLETED_GROUP")},
	// 			IsAzureMonitorTargetEnabled: to.Ptr(false),
	// 			IsStorageSecondaryKeyInUse: to.Ptr(false),
	// 			RetentionDays: to.Ptr[int32](0),
	// 			State: to.Ptr(armsql.BlobAuditingPolicyStateEnabled),
	// 			StorageAccountSubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			StorageEndpoint: to.Ptr("https://mystorage.blob.core.windows.net"),
	// 		},
	// 	}
}
