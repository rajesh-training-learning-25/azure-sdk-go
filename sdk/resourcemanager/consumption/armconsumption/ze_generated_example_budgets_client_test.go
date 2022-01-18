//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconsumption_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
)

// x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/BudgetsList.json
func ExampleBudgetsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armconsumption.NewBudgetsClient(cred, nil)
	pager := client.List("<scope>",
		nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/Budget.json
func ExampleBudgetsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armconsumption.NewBudgetsClient(cred, nil)
	res, err := client.Get(ctx,
		"<scope>",
		"<budget-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.BudgetsClientGetResult)
}

// x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/CreateOrUpdateBudget.json
func ExampleBudgetsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armconsumption.NewBudgetsClient(cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<scope>",
		"<budget-name>",
		armconsumption.Budget{
			ETag: to.StringPtr("<etag>"),
			Properties: &armconsumption.BudgetProperties{
				Amount:   to.Float64Ptr(100.65),
				Category: armconsumption.CategoryType("Cost").ToPtr(),
				Filter: &armconsumption.BudgetFilter{
					And: []*armconsumption.BudgetFilterProperties{
						{
							Dimensions: &armconsumption.BudgetComparisonExpression{
								Name:     to.StringPtr("<name>"),
								Operator: armconsumption.BudgetOperatorType("In").ToPtr(),
								Values: []*string{
									to.StringPtr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Compute/virtualMachines/MSVM2"),
									to.StringPtr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Compute/virtualMachines/platformcloudplatformGeneric1")},
							},
						},
						{
							Tags: &armconsumption.BudgetComparisonExpression{
								Name:     to.StringPtr("<name>"),
								Operator: armconsumption.BudgetOperatorType("In").ToPtr(),
								Values: []*string{
									to.StringPtr("Dev"),
									to.StringPtr("Prod")},
							},
						},
						{
							Tags: &armconsumption.BudgetComparisonExpression{
								Name:     to.StringPtr("<name>"),
								Operator: armconsumption.BudgetOperatorType("In").ToPtr(),
								Values: []*string{
									to.StringPtr("engineering"),
									to.StringPtr("sales")},
							},
						}},
				},
				Notifications: map[string]*armconsumption.Notification{
					"Actual_GreaterThan_80_Percent": {
						ContactEmails: []*string{
							to.StringPtr("johndoe@contoso.com"),
							to.StringPtr("janesmith@contoso.com")},
						ContactGroups: []*string{
							to.StringPtr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/microsoft.insights/actionGroups/SampleActionGroup")},
						ContactRoles: []*string{
							to.StringPtr("Contributor"),
							to.StringPtr("Reader")},
						Enabled:       to.BoolPtr(true),
						Locale:        armconsumption.CultureCode("en-us").ToPtr(),
						Operator:      armconsumption.OperatorType("GreaterThan").ToPtr(),
						Threshold:     to.Float64Ptr(80),
						ThresholdType: armconsumption.ThresholdType("Actual").ToPtr(),
					},
				},
				TimeGrain: armconsumption.TimeGrainType("Monthly").ToPtr(),
				TimePeriod: &armconsumption.BudgetTimePeriod{
					EndDate:   to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-31T00:00:00Z"); return t }()),
					StartDate: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-01T00:00:00Z"); return t }()),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.BudgetsClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/DeleteBudget.json
func ExampleBudgetsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armconsumption.NewBudgetsClient(cred, nil)
	_, err = client.Delete(ctx,
		"<scope>",
		"<budget-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
