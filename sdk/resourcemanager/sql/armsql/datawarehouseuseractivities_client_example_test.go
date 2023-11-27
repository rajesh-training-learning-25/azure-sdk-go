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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/GetDataWarehouseUserActivities.json
func ExampleDataWarehouseUserActivitiesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataWarehouseUserActivitiesClient().Get(ctx, "Default-SQL-SouthEastAsia", "testsvr", "testdb", armsql.DataWarehouseUserActivityNameCurrent, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DataWarehouseUserActivities = armsql.DataWarehouseUserActivities{
	// 	Name: to.Ptr("current"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/dataWarehouseUserActivities"),
	// 	ID: to.Ptr("subscriptions/326affc3-21f4-4471-a545-e37430b70113/resourceGroups/Default-SQL-Onebox/providers/Microsoft.Sql/servers/testsvr/databases/dwdb01/dataWarehouseUserActivities/current"),
	// 	Properties: &armsql.DataWarehouseUserActivitiesProperties{
	// 		ActiveQueriesCount: to.Ptr[int32](0),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ListDataWarehouseUserActivities.json
func ExampleDataWarehouseUserActivitiesClient_NewListByDatabasePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDataWarehouseUserActivitiesClient().NewListByDatabasePager("Default-SQL-SouthEastAsia", "testsvr", "testdb", nil)
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
		// page.DataWarehouseUserActivitiesListResult = armsql.DataWarehouseUserActivitiesListResult{
		// 	Value: []*armsql.DataWarehouseUserActivities{
		// 		{
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases/dataWarehouseUserActivities"),
		// 			ID: to.Ptr("subscriptions/326affc3-21f4-4471-a545-e37430b70113/resourceGroups/Default-SQL-Onebox/providers/Microsoft.Sql/servers/testsvr/databases/dwdb01/dataWarehouseUserActivities/current"),
		// 			Properties: &armsql.DataWarehouseUserActivitiesProperties{
		// 				ActiveQueriesCount: to.Ptr[int32](0),
		// 			},
		// 	}},
		// }
	}
}
