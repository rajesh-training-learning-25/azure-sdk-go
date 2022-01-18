//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsight/armsecurityinsight"
)

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/entityQueries/GetEntityQueries.json
func ExampleEntityQueriesClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsight.NewEntityQueriesClient("<subscription-id>", cred, nil)
	pager := client.List("<resource-group-name>",
		"<workspace-name>",
		&armsecurityinsight.EntityQueriesClientListOptions{Kind: armsecurityinsight.Enum8("Expansion").ToPtr()})
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

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/entityQueries/GetActivityEntityQueryById.json
func ExampleEntityQueriesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsight.NewEntityQueriesClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<entity-query-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.EntityQueriesClientGetResult)
}

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/entityQueries/CreateEntityQueryActivity.json
func ExampleEntityQueriesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsight.NewEntityQueriesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<entity-query-id>",
		&armsecurityinsight.ActivityCustomEntityQuery{
			Etag: to.StringPtr("<etag>"),
			Kind: armsecurityinsight.CustomEntityQueryKind("Activity").ToPtr(),
			Properties: &armsecurityinsight.ActivityEntityQueriesProperties{
				Description: to.StringPtr("<description>"),
				Content:     to.StringPtr("<content>"),
				Enabled:     to.BoolPtr(true),
				EntitiesFilter: map[string][]*string{
					"Host_OsFamily": {
						to.StringPtr("Windows")},
				},
				InputEntityType: armsecurityinsight.EntityType("Host").ToPtr(),
				QueryDefinitions: &armsecurityinsight.ActivityEntityQueriesPropertiesQueryDefinitions{
					Query: to.StringPtr("<query>"),
				},
				RequiredInputFieldsSets: [][]*string{
					{
						to.StringPtr("Host_HostName"),
						to.StringPtr("Host_NTDomain")},
					{
						to.StringPtr("Host_HostName"),
						to.StringPtr("Host_DnsDomain")},
					{
						to.StringPtr("Host_AzureID")},
					{
						to.StringPtr("Host_OMSAgentID")}},
				Title: to.StringPtr("<title>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.EntityQueriesClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/entityQueries/DeleteEntityQuery.json
func ExampleEntityQueriesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsight.NewEntityQueriesClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<entity-query-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
