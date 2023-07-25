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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edd7863aeaad7379a7a8eef424ca11617d8c91eb/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedColumnsListByDatabaseMax.json
func ExampleManagedDatabaseColumnsClient_NewListByDatabasePager_filterManagedDatabaseColumns() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedDatabaseColumnsClient().NewListByDatabasePager("myRG", "serverName", "myDatabase", &armsql.ManagedDatabaseColumnsClientListByDatabaseOptions{Schema: []string{
		"dbo"},
		Table: []string{
			"customer",
			"address"},
		Column: []string{
			"username"},
		OrderBy: []string{
			"schema asc",
			"table",
			"column desc"},
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
		// page.DatabaseColumnListResult = armsql.DatabaseColumnListResult{
		// 	Value: []*armsql.DatabaseColumn{
		// 		{
		// 			Name: to.Ptr("username"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/databases/schemas/tables/columns"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myRG/providers/Microsoft.Sql/managedInstances/serverName/databases/myDatabase/schemas/dbo/tables/customer/columns/username"),
		// 			Properties: &armsql.DatabaseColumnProperties{
		// 				ColumnType: to.Ptr(armsql.ColumnDataTypeNvarchar),
		// 				IsComputed: to.Ptr(false),
		// 				MemoryOptimized: to.Ptr(false),
		// 				TemporalType: to.Ptr(armsql.TableTemporalTypeNonTemporalTable),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edd7863aeaad7379a7a8eef424ca11617d8c91eb/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedColumnsListByDatabaseMin.json
func ExampleManagedDatabaseColumnsClient_NewListByDatabasePager_listManagedDatabaseColumns() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedDatabaseColumnsClient().NewListByDatabasePager("myRG", "serverName", "myDatabase", &armsql.ManagedDatabaseColumnsClientListByDatabaseOptions{Schema: []string{},
		Table:     []string{},
		Column:    []string{},
		OrderBy:   []string{},
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
		// page.DatabaseColumnListResult = armsql.DatabaseColumnListResult{
		// 	Value: []*armsql.DatabaseColumn{
		// 		{
		// 			Name: to.Ptr("col1"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/databases/schemas/tables/columns"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myRG/providers/Microsoft.Sql/managedInstances/serverName/databases/myDatabase/schemas/dbo/tables/table1/columns/col1"),
		// 			Properties: &armsql.DatabaseColumnProperties{
		// 				ColumnType: to.Ptr(armsql.ColumnDataTypeInt),
		// 				IsComputed: to.Ptr(false),
		// 				MemoryOptimized: to.Ptr(false),
		// 				TemporalType: to.Ptr(armsql.TableTemporalTypeNonTemporalTable),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("col2"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/databases/schemas/tables/columns"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myRG/providers/Microsoft.Sql/managedInstances/serverName/databases/myDatabase/schemas/dbo/tables/table1/columns/col2"),
		// 			Properties: &armsql.DatabaseColumnProperties{
		// 				ColumnType: to.Ptr(armsql.ColumnDataTypeBit),
		// 				IsComputed: to.Ptr(false),
		// 				MemoryOptimized: to.Ptr(false),
		// 				TemporalType: to.Ptr(armsql.TableTemporalTypeSystemVersionedTemporalTable),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edd7863aeaad7379a7a8eef424ca11617d8c91eb/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseColumnListByTable.json
func ExampleManagedDatabaseColumnsClient_NewListByTablePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedDatabaseColumnsClient().NewListByTablePager("myRG", "myManagedInstanceName", "myDatabase", "dbo", "table1", &armsql.ManagedDatabaseColumnsClientListByTableOptions{Filter: nil})
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
		// page.DatabaseColumnListResult = armsql.DatabaseColumnListResult{
		// 	Value: []*armsql.DatabaseColumn{
		// 		{
		// 			Name: to.Ptr("col1"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/databases/schemas/tables/columns"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myRG/providers/Microsoft.Sql/managedInstances/myManagedInstanceName/databases/myDatabase/schemas/dbo/tables/table1/columns/col1"),
		// 			Properties: &armsql.DatabaseColumnProperties{
		// 				ColumnType: to.Ptr(armsql.ColumnDataTypeNvarchar),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("col2"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/databases/schemas/tables/columns"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myRG/providers/Microsoft.Sql/managedInstances/myManagedInstanceName/databases/myDatabase/schemas/dbo/tables/table1/columns/col2"),
		// 			Properties: &armsql.DatabaseColumnProperties{
		// 				ColumnType: to.Ptr(armsql.ColumnDataTypeBit),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edd7863aeaad7379a7a8eef424ca11617d8c91eb/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseColumnGet.json
func ExampleManagedDatabaseColumnsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedDatabaseColumnsClient().Get(ctx, "myRG", "myManagedInstanceName", "myDatabase", "dbo", "table1", "column1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DatabaseColumn = armsql.DatabaseColumn{
	// 	Name: to.Ptr("column1"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances/databases/schemas/tables/columns"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myRG/providers/Microsoft.Sql/managedInstances/myManagedInstanceName/databases/myDatabase/schemas/dbo/tables/table1/columns/column1"),
	// 	Properties: &armsql.DatabaseColumnProperties{
	// 		ColumnType: to.Ptr(armsql.ColumnDataTypeBit),
	// 	},
	// }
}
