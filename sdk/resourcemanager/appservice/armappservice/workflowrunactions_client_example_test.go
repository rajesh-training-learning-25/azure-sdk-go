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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/10c740b7224771c889cfb92f128168f5a0568c26/specification/web/resource-manager/Microsoft.Web/stable/2022-09-01/examples/WorkflowRunActions_List.json
func ExampleWorkflowRunActionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappservice.NewWorkflowRunActionsClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("test-resource-group", "test-name", "test-workflow", "08586676746934337772206998657CU22", &armappservice.WorkflowRunActionsClientListOptions{Top: nil,
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
		// page.WorkflowRunActionListResult = armappservice.WorkflowRunActionListResult{
		// 	Value: []*armappservice.WorkflowRunAction{
		// 		{
		// 			ID: to.Ptr("/workflows/test-workflow/runs/08586676746934337772206998657CU22/actions/HTTP"),
		// 			Name: to.Ptr("HTTP"),
		// 			Type: to.Ptr("/workflows/runs/actions"),
		// 			Properties: &armappservice.WorkflowRunActionProperties{
		// 				Code: to.Ptr("OK"),
		// 				Correlation: &armappservice.RunActionCorrelation{
		// 					ClientTrackingID: to.Ptr("08586676746934337772206998657CU22"),
		// 					ActionTrackingID: to.Ptr("56063357-45dd-4278-9be5-8220ce0cc9ca"),
		// 				},
		// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.5450625Z"); return t}()),
		// 				InputsLink: &armappservice.ContentLink{
		// 					ContentSize: to.Ptr[int64](138),
		// 					URI: to.Ptr("https://test-site.azurewebsites.net:443/runtime/webhooks/workflow/scaleUnits/prod-00/workflows/000012345678900000000056a41a/runs/00000aaaabbbbccccddddeeeeffff/actions/HTTP/contents/ActionInputs?api-version=2018-11-01&code=examplecode=2022-04-13T03%3A00%3A00.0000000Z&sp=%2Fruns%2F00000aaaabbbbccccddddeeeeffff%2Factions%2FHTTP%2Fcontents%2FActionInputs%2Fread&sv=1.0&sig=dJ6F00000aaaabbbbccccddddeeeeffff78co"),
		// 				},
		// 				OutputsLink: &armappservice.ContentLink{
		// 					ContentSize: to.Ptr[int64](15660),
		// 					URI: to.Ptr("https://test-site.azurewebsites.net:443/runtime/webhooks/workflow/scaleUnits/prod-00/workflows/000012345678900000000056a41a/runs/00000aaaabbbbccccddddeeeeffff/actions/HTTP/contents/ActionOutputs?api-version=2018-11-01&code=examplecode=2022-04-13T03%3A00%3A00.0000000Z&sp=%2Fruns%2F00000aaaabbbbccccddddeeeeffff%2Factions%2FHTTP%2Fcontents%2FActionOutputs%2Fread&sv=1.0&sig=fIIgF00000aaaabbbbccccddddeeeeffffWRU0"),
		// 				},
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.305236Z"); return t}()),
		// 				Status: to.Ptr(armappservice.WorkflowStatusSucceeded),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/10c740b7224771c889cfb92f128168f5a0568c26/specification/web/resource-manager/Microsoft.Web/stable/2022-09-01/examples/WorkflowRunActions_Get.json
func ExampleWorkflowRunActionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappservice.NewWorkflowRunActionsClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "test-resource-group", "test-name", "test-workflow", "08586676746934337772206998657CU22", "HTTP", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkflowRunAction = armappservice.WorkflowRunAction{
	// 	ID: to.Ptr("/workflows/test-workflow/runs/08586676746934337772206998657CU22/actions/HTTP"),
	// 	Name: to.Ptr("HTTP"),
	// 	Type: to.Ptr("/workflows/runs/actions"),
	// 	Properties: &armappservice.WorkflowRunActionProperties{
	// 		Code: to.Ptr("OK"),
	// 		Correlation: &armappservice.RunActionCorrelation{
	// 			ClientTrackingID: to.Ptr("08586676746934337772206998657CU22"),
	// 			ActionTrackingID: to.Ptr("56063357-45dd-4278-9be5-8220ce0cc9ca"),
	// 		},
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.5450625Z"); return t}()),
	// 		InputsLink: &armappservice.ContentLink{
	// 			ContentSize: to.Ptr[int64](138),
	// 			URI: to.Ptr("https://test-site.azurewebsites.net:443/runtime/webhooks/workflow/scaleUnits/prod-00/workflows/000012345678900000000056a41a/runs/00000aaaabbbbccccddddeeeeffff/actions/HTTP/contents/ActionInputs?api-version=2018-11-01&code=examplecode=2022-04-13T03%3A00%3A00.0000000Z&sp=%2Fruns%2F00000aaaabbbbccccddddeeeeffff%2Factions%2FHTTP%2Fcontents%2FActionInputs%2Fread&sv=1.0&sig=dJ6F00000aaaabbbbccccddddeeeeffff78co"),
	// 		},
	// 		OutputsLink: &armappservice.ContentLink{
	// 			ContentSize: to.Ptr[int64](15660),
	// 			URI: to.Ptr("https://test-site.azurewebsites.net:443/runtime/webhooks/workflow/scaleUnits/prod-00/workflows/000012345678900000000056a41a/runs/00000aaaabbbbccccddddeeeeffff/actions/HTTP/contents/ActionOutputs?api-version=2018-11-01&code=examplecode=2022-04-13T03%3A00%3A00.0000000Z&sp=%2Fruns%2F00000aaaabbbbccccddddeeeeffff%2Factions%2FHTTP%2Fcontents%2FActionOutputs%2Fread&sv=1.0&sig=fIIgF00000aaaabbbbccccddddeeeeffffWRU0"),
	// 		},
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.305236Z"); return t}()),
	// 		Status: to.Ptr(armappservice.WorkflowStatusSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/10c740b7224771c889cfb92f128168f5a0568c26/specification/web/resource-manager/Microsoft.Web/stable/2022-09-01/examples/WorkflowRunActions_ListExpressionTraces.json
func ExampleWorkflowRunActionsClient_NewListExpressionTracesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappservice.NewWorkflowRunActionsClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListExpressionTracesPager("testResourceGroup", "test-name", "testFlow", "08586776228332053161046300351", "testAction", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Inputs {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ExpressionTraces = armappservice.ExpressionTraces{
		// 	Inputs: []*armappservice.ExpressionRoot{
		// 		{
		// 			Text: to.Ptr("add(4, 4)"),
		// 			Value: float64(8),
		// 			Path: to.Ptr(""),
		// 	}},
		// }
	}
}
