//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/preview/2017-05-01-preview/examples/getDiagnosticSetting.json
func ExampleDiagnosticSettingsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewDiagnosticSettingsClient(cred, nil)
	res, err := client.Get(ctx,
		"<resource-uri>",
		"<name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DiagnosticSettingsClientGetResult)
}

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/preview/2017-05-01-preview/examples/createOrUpdateDiagnosticSetting.json
func ExampleDiagnosticSettingsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewDiagnosticSettingsClient(cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-uri>",
		"<name>",
		armmonitor.DiagnosticSettingsResource{
			Properties: &armmonitor.DiagnosticSettings{
				EventHubAuthorizationRuleID: to.StringPtr("<event-hub-authorization-rule-id>"),
				EventHubName:                to.StringPtr("<event-hub-name>"),
				LogAnalyticsDestinationType: to.StringPtr("<log-analytics-destination-type>"),
				Logs: []*armmonitor.LogSettings{
					{
						Category: to.StringPtr("<category>"),
						Enabled:  to.BoolPtr(true),
						RetentionPolicy: &armmonitor.RetentionPolicy{
							Days:    to.Int32Ptr(0),
							Enabled: to.BoolPtr(false),
						},
					}},
				Metrics: []*armmonitor.MetricSettings{
					{
						Category: to.StringPtr("<category>"),
						Enabled:  to.BoolPtr(true),
						RetentionPolicy: &armmonitor.RetentionPolicy{
							Days:    to.Int32Ptr(0),
							Enabled: to.BoolPtr(false),
						},
					}},
				StorageAccountID: to.StringPtr("<storage-account-id>"),
				WorkspaceID:      to.StringPtr("<workspace-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DiagnosticSettingsClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/preview/2017-05-01-preview/examples/deleteDiagnosticSetting.json
func ExampleDiagnosticSettingsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewDiagnosticSettingsClient(cred, nil)
	_, err = client.Delete(ctx,
		"<resource-uri>",
		"<name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/preview/2017-05-01-preview/examples/listDiagnosticSettings.json
func ExampleDiagnosticSettingsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewDiagnosticSettingsClient(cred, nil)
	res, err := client.List(ctx,
		"<resource-uri>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DiagnosticSettingsClientListResult)
}
