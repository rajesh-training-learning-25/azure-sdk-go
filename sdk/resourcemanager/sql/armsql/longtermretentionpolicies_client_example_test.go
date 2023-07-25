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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edd7863aeaad7379a7a8eef424ca11617d8c91eb/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/LongTermRetentionPolicyGet.json
func ExampleLongTermRetentionPoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLongTermRetentionPoliciesClient().Get(ctx, "resourceGroup", "testserver", "testDatabase", armsql.LongTermRetentionPolicyNameDefault, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LongTermRetentionPolicy = armsql.LongTermRetentionPolicy{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Sql/resourceGroups/servers/databases/backupLongTermRetentionPolicies"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/resourceGroups/resourceGroup/servers/testserver/databases/testDatabase/backupLongTermRetentionPolicies/default"),
	// 	Properties: &armsql.BaseLongTermRetentionPolicyProperties{
	// 		MonthlyRetention: to.Ptr("P1Y"),
	// 		WeekOfYear: to.Ptr[int32](5),
	// 		WeeklyRetention: to.Ptr("P1M"),
	// 		YearlyRetention: to.Ptr("P5Y"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edd7863aeaad7379a7a8eef424ca11617d8c91eb/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/LongTermRetentionPolicyCreateOrUpdate.json
func ExampleLongTermRetentionPoliciesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLongTermRetentionPoliciesClient().BeginCreateOrUpdate(ctx, "resourceGroup", "testserver", "testDatabase", armsql.LongTermRetentionPolicyNameDefault, armsql.LongTermRetentionPolicy{
		Properties: &armsql.BaseLongTermRetentionPolicyProperties{
			MonthlyRetention: to.Ptr("P1Y"),
			WeekOfYear:       to.Ptr[int32](5),
			WeeklyRetention:  to.Ptr("P1M"),
			YearlyRetention:  to.Ptr("P5Y"),
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
	// res.LongTermRetentionPolicy = armsql.LongTermRetentionPolicy{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Sql/resourceGroups/servers/databases/backupLongTermRetentionPolicies"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/resourceGroups/resourceGroup/servers/testserver/databases/testDatabase/backupLongTermRetentionPolicies/default"),
	// 	Properties: &armsql.BaseLongTermRetentionPolicyProperties{
	// 		MonthlyRetention: to.Ptr("P1Y"),
	// 		WeekOfYear: to.Ptr[int32](5),
	// 		WeeklyRetention: to.Ptr("P1M"),
	// 		YearlyRetention: to.Ptr("P5Y"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edd7863aeaad7379a7a8eef424ca11617d8c91eb/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/LongTermRetentionPolicyListByDatabase.json
func ExampleLongTermRetentionPoliciesClient_NewListByDatabasePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLongTermRetentionPoliciesClient().NewListByDatabasePager("resourceGroup", "testserver", "testDatabase", nil)
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
		// page.LongTermRetentionPolicyListResult = armsql.LongTermRetentionPolicyListResult{
		// 	Value: []*armsql.LongTermRetentionPolicy{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.Sql/resourceGroups/servers/databases/backupLongTermRetentionPolicies"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/resourceGroups/resourceGroup/servers/testserver/databases/testDatabase/backupLongTermRetentionPolicies/default"),
		// 			Properties: &armsql.BaseLongTermRetentionPolicyProperties{
		// 				MonthlyRetention: to.Ptr("P1Y"),
		// 				WeekOfYear: to.Ptr[int32](5),
		// 				WeeklyRetention: to.Ptr("P1M"),
		// 				YearlyRetention: to.Ptr("P5Y"),
		// 			},
		// 	}},
		// }
	}
}
