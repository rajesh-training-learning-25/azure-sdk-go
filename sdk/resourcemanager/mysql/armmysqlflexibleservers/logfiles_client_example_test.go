//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmysqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/mysql/resource-manager/Microsoft.DBforMySQL/LogFiles/preview/2021-12-01-preview/examples/LogFilesListByServer.json
func ExampleLogFilesClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLogFilesClient().NewListByServerPager("testrg", "mysqltestsvc1", nil)
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
		// page.LogFileListResult = armmysqlflexibleservers.LogFileListResult{
		// 	Value: []*armmysqlflexibleservers.LogFile{
		// 		{
		// 			Name: to.Ptr("mysql-slow-mysqltestsvc1-2018022823.log"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/logFiles"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/flexibleServers/mysqltestsvc1/logFiles/mysql-slow-mysqltestsvc1-2018022823.log"),
		// 			Properties: &armmysqlflexibleservers.LogFileProperties{
		// 				Type: to.Ptr("slowlog"),
		// 				CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-01T06:09:20.000Z"); return t}()),
		// 				SizeInKB: to.Ptr[int64](1),
		// 				URL: to.Ptr("https://wasd2prodwus1afse42.file.core.windows.net/833c99b2f36c47349e5554b903fe0440/serverlogs/mysql-slow-mysqltestsvc1-2018022823.log?sv=2015-04-05&sr=f&sig=D9Ga4N5Pa%2BPe5Bmjpvs7A0TPD%2FF7IZpk9e4KWR0jgpM%3D&se=2018-03-01T07%3A12%3A13Z&sp=r"),
		// 			},
		// 	}},
		// }
	}
}
