//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armpostgresql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/DatabaseCreate.json
func ExampleDatabasesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDatabasesClient().BeginCreateOrUpdate(ctx, "TestGroup", "testserver", "db1", armpostgresql.Database{
		Properties: &armpostgresql.DatabaseProperties{
			Charset:   to.Ptr("UTF8"),
			Collation: to.Ptr("English_United States.1252"),
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
	// res.Database = armpostgresql.Database{
	// 	Name: to.Ptr("db1"),
	// 	Type: to.Ptr("Microsoft.DBforPostgreSQL/servers/databases"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/servers/testserver/databases/db1"),
	// 	Properties: &armpostgresql.DatabaseProperties{
	// 		Charset: to.Ptr("UTF8"),
	// 		Collation: to.Ptr("English_United States.1252"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/DatabaseDelete.json
func ExampleDatabasesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDatabasesClient().BeginDelete(ctx, "TestGroup", "testserver", "db1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/DatabaseGet.json
func ExampleDatabasesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDatabasesClient().Get(ctx, "TestGroup", "testserver", "db1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Database = armpostgresql.Database{
	// 	Name: to.Ptr("db1"),
	// 	Type: to.Ptr("Microsoft.DBforPostgreSQL/servers/databases"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/servers/testserver/databases/db1"),
	// 	Properties: &armpostgresql.DatabaseProperties{
	// 		Charset: to.Ptr("UTF8"),
	// 		Collation: to.Ptr("English_United States.1252"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/DatabaseListByServer.json
func ExampleDatabasesClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDatabasesClient().NewListByServerPager("TestGroup", "testserver", nil)
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
		// page.DatabaseListResult = armpostgresql.DatabaseListResult{
		// 	Value: []*armpostgresql.Database{
		// 		{
		// 			Name: to.Ptr("db1"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/servers/databases"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/servers/testserver/databases/db1"),
		// 			Properties: &armpostgresql.DatabaseProperties{
		// 				Charset: to.Ptr("UTF8"),
		// 				Collation: to.Ptr("English_United States.1252"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("db2"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/servers/databases"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/servers/testserver/databases/db2"),
		// 			Properties: &armpostgresql.DatabaseProperties{
		// 				Charset: to.Ptr("UTF8"),
		// 				Collation: to.Ptr("English_United States.1252"),
		// 			},
		// 	}},
		// }
	}
}
