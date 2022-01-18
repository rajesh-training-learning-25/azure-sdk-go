//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ServerBlobAuditingGet.json
func ExampleServerBlobAuditingPoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsql.NewServerBlobAuditingPoliciesClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<server-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ServerBlobAuditingPoliciesClientGetResult)
}

// x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ServerBlobAuditingCreateMax.json
func ExampleServerBlobAuditingPoliciesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsql.NewServerBlobAuditingPoliciesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<server-name>",
		armsql.ServerBlobAuditingPolicy{
			Properties: &armsql.ServerBlobAuditingPolicyProperties{
				AuditActionsAndGroups: []*string{
					to.StringPtr("SUCCESSFUL_DATABASE_AUTHENTICATION_GROUP"),
					to.StringPtr("FAILED_DATABASE_AUTHENTICATION_GROUP"),
					to.StringPtr("BATCH_COMPLETED_GROUP")},
				IsAzureMonitorTargetEnabled:  to.BoolPtr(true),
				IsStorageSecondaryKeyInUse:   to.BoolPtr(false),
				QueueDelayMs:                 to.Int32Ptr(4000),
				RetentionDays:                to.Int32Ptr(6),
				State:                        armsql.BlobAuditingPolicyStateEnabled.ToPtr(),
				StorageAccountAccessKey:      to.StringPtr("<storage-account-access-key>"),
				StorageAccountSubscriptionID: to.StringPtr("<storage-account-subscription-id>"),
				StorageEndpoint:              to.StringPtr("<storage-endpoint>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ServerBlobAuditingPoliciesClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ServerAuditingSettingsList.json
func ExampleServerBlobAuditingPoliciesClient_ListByServer() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsql.NewServerBlobAuditingPoliciesClient("<subscription-id>", cred, nil)
	pager := client.ListByServer("<resource-group-name>",
		"<server-name>",
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
