//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fba7ffa9cee6453e2a3cf8c857074a323252a12d/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ListJobVersions.json
func ExampleJobVersionsClient_NewListByJobPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewJobVersionsClient().NewListByJobPager("group1", "server1", "agent1", "job1", nil)
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
		// page.JobVersionListResult = armsql.JobVersionListResult{
		// 	Value: []*armsql.JobVersion{
		// 		{
		// 			Name: to.Ptr("1"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/jobAgents/jobs/versions"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/jobs/job1/versions/1"),
		// 		},
		// 		{
		// 			Name: to.Ptr("2"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/jobAgents/jobs/versions"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/jobs/job1/versions/2"),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fba7ffa9cee6453e2a3cf8c857074a323252a12d/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/GetJobVersion.json
func ExampleJobVersionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewJobVersionsClient().Get(ctx, "group1", "server1", "agent1", "job1", 1, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.JobVersion = armsql.JobVersion{
	// 	Name: to.Ptr("1"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/jobAgents/jobs/versions"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/jobs/job1/versions/1"),
	// }
}
