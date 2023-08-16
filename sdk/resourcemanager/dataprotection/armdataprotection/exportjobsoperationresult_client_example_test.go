//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-05-01/examples/JobCRUD/GetExportJobsOperationResult.json
func ExampleExportJobsOperationResultClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExportJobsOperationResultClient().Get(ctx, "SwaggerTestRg", "NetSDKTestRsVault", "00000000-0000-0000-0000-000000000000", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ExportJobsResult = armdataprotection.ExportJobsResult{
	// 	BlobSasKey: to.Ptr("someKey"),
	// 	BlobURL: to.Ptr("https://azureblob.blob.core.windows.net/reportcontainer/exportjobsreport00000000-0000-0000-0000-000000000000"),
	// 	ExcelFileBlobSasKey: to.Ptr("someKey"),
	// 	ExcelFileBlobURL: to.Ptr("https://azureblob.blob.core.windows.net/reportcontainer/exportjobsreport00000000-0000-0000-0000-000000000000_ExcelFile.xlsx"),
	// }
}
