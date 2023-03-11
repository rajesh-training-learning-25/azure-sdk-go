//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/10c740b7224771c889cfb92f128168f5a0568c26/specification/web/resource-manager/Microsoft.Web/stable/2022-09-01/examples/WorkflowTriggerHistories_List.json
func ExampleWorkflowTriggerHistoriesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappservice.NewWorkflowTriggerHistoriesClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("testResourceGroup", "test-name", "testWorkflowName", "testTriggerName", &armappservice.WorkflowTriggerHistoriesClientListOptions{Top: nil,
		Filter: nil,
	})
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
		// page.WorkflowTriggerHistoryListResult = armappservice.WorkflowTriggerHistoryListResult{
		// 	Value: []*armappservice.WorkflowTriggerHistory{
		// 		{
		// 			ID: to.Ptr("/workflows/testWorkflowName/triggers/testTriggerName/histories/08586676746934337772206998657CU22"),
		// 			Name: to.Ptr("08586676746934337772206998657CU22"),
		// 			Type: to.Ptr("/workflows/triggers/histories"),
		// 			Properties: &armappservice.WorkflowTriggerHistoryProperties{
		// 				Code: to.Ptr("OK"),
		// 				Correlation: &armappservice.Correlation{
		// 					ClientTrackingID: to.Ptr("08586676746934337772206998657CU22"),
		// 				},
		// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.2987996Z"); return t}()),
		// 				Fired: to.Ptr(true),
		// 				Run: &armappservice.ResourceReference{
		// 					Name: to.Ptr("08586676746934337772206998657CU22"),
		// 					Type: to.Ptr("/workflows/runs"),
		// 					ID: to.Ptr("/workflows/testWorkflowName/runs/08586676746934337772206998657CU22"),
		// 				},
		// 				ScheduledTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:31.6344174Z"); return t}()),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.0387927Z"); return t}()),
		// 				Status: to.Ptr(armappservice.WorkflowStatusSucceeded),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/10c740b7224771c889cfb92f128168f5a0568c26/specification/web/resource-manager/Microsoft.Web/stable/2022-09-01/examples/WorkflowTriggerHistories_Get.json
func ExampleWorkflowTriggerHistoriesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappservice.NewWorkflowTriggerHistoriesClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "testResourceGroup", "test-name", "testWorkflowName", "testTriggerName", "08586676746934337772206998657CU22", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkflowTriggerHistory = armappservice.WorkflowTriggerHistory{
	// 	ID: to.Ptr("/workflows/testWorkflowName/triggers/testTriggerName/histories/08586676746934337772206998657CU22"),
	// 	Name: to.Ptr("08586676746934337772206998657CU22"),
	// 	Type: to.Ptr("/workflows/triggers/histories"),
	// 	Properties: &armappservice.WorkflowTriggerHistoryProperties{
	// 		Code: to.Ptr("OK"),
	// 		Correlation: &armappservice.Correlation{
	// 			ClientTrackingID: to.Ptr("08586676746934337772206998657CU22"),
	// 		},
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.2987996Z"); return t}()),
	// 		Fired: to.Ptr(true),
	// 		Run: &armappservice.ResourceReference{
	// 			Name: to.Ptr("08586676746934337772206998657CU22"),
	// 			Type: to.Ptr("/workflows/runs"),
	// 			ID: to.Ptr("/workflows/testWorkflowName/runs/08586676746934337772206998657CU22"),
	// 		},
	// 		ScheduledTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:31.6344174Z"); return t}()),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.0387927Z"); return t}()),
	// 		Status: to.Ptr(armappservice.WorkflowStatusSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/10c740b7224771c889cfb92f128168f5a0568c26/specification/web/resource-manager/Microsoft.Web/stable/2022-09-01/examples/WorkflowTriggerHistories_Resubmit.json
func ExampleWorkflowTriggerHistoriesClient_BeginResubmit() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappservice.NewWorkflowTriggerHistoriesClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginResubmit(ctx, "testResourceGroup", "test-name", "testWorkflowName", "testTriggerName", "testHistoryName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
