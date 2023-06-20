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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fba7ffa9cee6453e2a3cf8c857074a323252a12d/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/GetMaintenanceWindowOptions.json
func ExampleMaintenanceWindowOptionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMaintenanceWindowOptionsClient().Get(ctx, "Default-SQL-SouthEastAsia", "testsvr", "testdb", "current", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MaintenanceWindowOptions = armsql.MaintenanceWindowOptions{
	// 	Name: to.Ptr("current"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/maintenancewindows"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdb/maintenancewindowoptions/current"),
	// 	Properties: &armsql.MaintenanceWindowOptionsProperties{
	// 		AllowMultipleMaintenanceWindowsPerCycle: to.Ptr(true),
	// 		DefaultDurationInMinutes: to.Ptr[int32](120),
	// 		IsEnabled: to.Ptr(true),
	// 		MaintenanceWindowCycles: []*armsql.MaintenanceWindowTimeRange{
	// 			{
	// 				DayOfWeek: to.Ptr(armsql.DayOfWeekSaturday),
	// 				Duration: to.Ptr("PT60M"),
	// 				StartTime: to.Ptr("00:00:00"),
	// 		}},
	// 		MinCycles: to.Ptr[int32](2),
	// 		MinDurationInMinutes: to.Ptr[int32](60),
	// 		TimeGranularityInMinutes: to.Ptr[int32](5),
	// 	},
	// }
}
