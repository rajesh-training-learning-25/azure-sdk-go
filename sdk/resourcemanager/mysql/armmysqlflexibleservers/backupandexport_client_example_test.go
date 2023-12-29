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

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/mysql/resource-manager/Microsoft.DBforMySQL/Backups/preview/2022-09-30-preview/examples/BackupAndExport.json
func ExampleBackupAndExportClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBackupAndExportClient().BeginCreate(ctx, "TestGroup", "mysqltestserver", armmysqlflexibleservers.BackupAndExportRequest{
		BackupSettings: &armmysqlflexibleservers.BackupSettings{
			BackupName: to.Ptr("customer-backup-name"),
		},
		TargetDetails: &armmysqlflexibleservers.FullBackupStoreDetails{
			ObjectType: to.Ptr("FullBackupStoreDetails"),
			SasURIList: []*string{
				to.Ptr("sasuri1"),
				to.Ptr("sasuri2")},
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
	// res.BackupAndExportResponse = armmysqlflexibleservers.BackupAndExportResponse{
	// 	Name: to.Ptr("custom-backup101"),
	// 	Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/backupAndExport"),
	// 	ID: to.Ptr("/subscriptions/cb9d743d-2140-4e73-b871-cd31abab1d2f/resourceGroups/mrgsumitkumatest1/providers/Microsoft.DBforMySQL/flexibleServers/servermysql-01/backupAndExport/custom-backup101"),
	// 	EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-29T07:37:05.640Z"); return t}()),
	// 	Error: &armmysqlflexibleservers.ErrorResponse{
	// 		Code: to.Ptr("AggregateException"),
	// 		Message: to.Ptr("System.AggregateException: One or more errors occurred. (Mismatch in count of number of Commited-Blocks from service.)"),
	// 	},
	// 	PercentComplete: to.Ptr[float64](100),
	// 	Properties: &armmysqlflexibleservers.BackupAndExportResponseProperties{
	// 		BackupMetadata: to.Ptr("{\"key1\":\"value1\",\"key2\":\"value2\"}"),
	// 		DataTransferredInBytes: to.Ptr[int64](1024),
	// 		DatasourceSizeInBytes: to.Ptr[int64](1024),
	// 	},
	// 	StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-29T07:34:02.328Z"); return t}()),
	// 	Status: to.Ptr(armmysqlflexibleservers.OperationStatusFailed),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/mysql/resource-manager/Microsoft.DBforMySQL/Backups/preview/2022-09-30-preview/examples/ValidateBackup.json
func ExampleBackupAndExportClient_ValidateBackup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBackupAndExportClient().ValidateBackup(ctx, "TestGroup", "mysqltestserver", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ValidateBackupResponse = armmysqlflexibleservers.ValidateBackupResponse{
	// 	Properties: &armmysqlflexibleservers.ValidateBackupResponseProperties{
	// 		NumberOfContainers: to.Ptr[int32](1),
	// 	},
	// }
}
