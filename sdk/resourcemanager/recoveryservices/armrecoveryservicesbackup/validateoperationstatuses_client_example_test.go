//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a4ddec441435d1ef766c4f160eda658a69cc5dc2/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/AzureIaasVm/ValidateOperationStatus.json
func ExampleValidateOperationStatusesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewValidateOperationStatusesClient().Get(ctx, "NetSDKTestRsVault", "SwaggerTestRg", "00000000-0000-0000-0000-000000000000", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationStatus = armrecoveryservicesbackup.OperationStatus{
	// 	Name: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 	EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-29T06:04:18.207Z"); return t}()),
	// 	ID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 	Properties: &armrecoveryservicesbackup.OperationStatusValidateOperationExtendedInfo{
	// 		ObjectType: to.Ptr("OperationStatusValidateOperationExtendedInfo"),
	// 		ValidateOperationResponse: &armrecoveryservicesbackup.ValidateOperationResponse{
	// 			ValidationResults: []*armrecoveryservicesbackup.ErrorDetail{
	// 				{
	// 					Code: to.Ptr("UserErrorCoreCountSubscriptionQuotaReached"),
	// 					Message: to.Ptr("Core Count subscription quota has been reached."),
	// 					Recommendations: []*string{
	// 						to.Ptr("Contact Azure support to increase the limits.")},
	// 				}},
	// 			},
	// 		},
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-29T06:04:18.207Z"); return t}()),
	// 		Status: to.Ptr(armrecoveryservicesbackup.OperationStatusValuesSucceeded),
	// 	}
}
