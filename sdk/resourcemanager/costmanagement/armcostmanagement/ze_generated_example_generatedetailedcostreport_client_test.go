//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcostmanagement_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement"
)

// x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2021-10-01/examples/GenerateDetailedCostReportByBillingAccountLegacyAndBillingPeriod.json
func ExampleGenerateDetailedCostReportClient_BeginCreateOperation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcostmanagement.NewGenerateDetailedCostReportClient(cred, nil)
	poller, err := client.BeginCreateOperation(ctx,
		"<scope>",
		armcostmanagement.GenerateDetailedCostReportDefinition{
			BillingPeriod: to.StringPtr("<billing-period>"),
			Metric:        armcostmanagement.GenerateDetailedCostReportMetricType("ActualCost").ToPtr(),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.GenerateDetailedCostReportClientCreateOperationResult)
}
