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

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edd7863aeaad7379a7a8eef424ca11617d8c91eb/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseTableListBySchema.json
func ExampleManagedDatabaseTablesClient_NewListBySchemaPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedDatabaseTablesClient().NewListBySchemaPager("myRG", "myManagedInstanceName", "myDatabase", "dbo", &armsql.ManagedDatabaseTablesClientListBySchemaOptions{Filter: nil})
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
		// page.DatabaseTableListResult = armsql.DatabaseTableListResult{
		// 	Value: []*armsql.DatabaseTable{
		// 		{
		// 			Name: to.Ptr("table1"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/databases/schemas/tables"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myRG/providers/Microsoft.Sql/managedInstances/myManagedInstanceName/databases/myDatabase/schemas/dbo/tables/table1"),
		// 		},
		// 		{
		// 			Name: to.Ptr("table2"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/databases/schemas/tables"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myRG/providers/Microsoft.Sql/managedInstances/myManagedInstanceName/databases/myDatabase/schemas/dbo/tables/table2"),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edd7863aeaad7379a7a8eef424ca11617d8c91eb/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseTableGet.json
func ExampleManagedDatabaseTablesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedDatabaseTablesClient().Get(ctx, "myRG", "myManagedInstanceName", "myDatabase", "dbo", "table1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DatabaseTable = armsql.DatabaseTable{
	// 	Name: to.Ptr("table1"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances/databases/schemas/tables"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myRG/providers/Microsoft.Sql/managedInstances/myManagedInstanceName/databases/myDatabase/schemas/dbo/tables/table1"),
	// }
}
