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

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2017-04-01/examples/createOrUpdateActivityLogAlert.json
func ExampleActivityLogAlertsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewActivityLogAlertsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<activity-log-alert-name>",
		armmonitor.ActivityLogAlertResource{
			Resource: armmonitor.Resource{
				Location: to.StringPtr("<location>"),
				Tags:     map[string]*string{},
			},
			Properties: &armmonitor.ActivityLogAlert{
				Description: to.StringPtr("<description>"),
				Actions: &armmonitor.ActivityLogAlertActionList{
					ActionGroups: []*armmonitor.ActivityLogAlertActionGroup{
						{
							ActionGroupID: to.StringPtr("<action-group-id>"),
							WebhookProperties: map[string]*string{
								"sampleWebhookProperty": to.StringPtr("samplePropertyValue"),
							},
						}},
				},
				Condition: &armmonitor.ActivityLogAlertAllOfCondition{
					AllOf: []*armmonitor.ActivityLogAlertLeafCondition{
						{
							Equals: to.StringPtr("<equals>"),
							Field:  to.StringPtr("<field>"),
						},
						{
							Equals: to.StringPtr("<equals>"),
							Field:  to.StringPtr("<field>"),
						}},
				},
				Enabled: to.BoolPtr(true),
				Scopes: []*string{
					to.StringPtr("subscriptions/187f412d-1758-44d9-b052-169e2564721d")},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ActivityLogAlertResource.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2017-04-01/examples/getActivityLogAlert.json
func ExampleActivityLogAlertsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewActivityLogAlertsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<activity-log-alert-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ActivityLogAlertResource.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2017-04-01/examples/deleteActivityLogAlert.json
func ExampleActivityLogAlertsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewActivityLogAlertsClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<activity-log-alert-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2017-04-01/examples/patchActivityLogAlert.json
func ExampleActivityLogAlertsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewActivityLogAlertsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<activity-log-alert-name>",
		armmonitor.ActivityLogAlertPatchBody{
			Properties: &armmonitor.ActivityLogAlertPatch{
				Enabled: to.BoolPtr(false),
			},
			Tags: map[string]*string{
				"key1": to.StringPtr("value1"),
				"key2": to.StringPtr("value2"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ActivityLogAlertResource.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2017-04-01/examples/listActivityLogAlerts.json
func ExampleActivityLogAlertsClient_ListBySubscriptionID() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewActivityLogAlertsClient("<subscription-id>", cred, nil)
	_, err = client.ListBySubscriptionID(ctx,
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2017-04-01/examples/listActivityLogAlerts.json
func ExampleActivityLogAlertsClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewActivityLogAlertsClient("<subscription-id>", cred, nil)
	_, err = client.ListByResourceGroup(ctx,
		"<resource-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
