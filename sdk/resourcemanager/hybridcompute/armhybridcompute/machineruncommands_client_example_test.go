//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b574e2a41acda14a90ef237006e8bbdda2b63c63/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-05-20-preview/examples/runCommand/RunCommands_CreateOrUpdate.json
func ExampleMachineRunCommandsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMachineRunCommandsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myMachine", "myRunCommand", armhybridcompute.MachineRunCommand{
		Location: to.Ptr("eastus2"),
		Properties: &armhybridcompute.MachineRunCommandProperties{
			AsyncExecution: to.Ptr(false),
			ErrorBlobURI:   to.Ptr("https://mystorageaccount.blob.core.windows.net/mycontainer/MyScriptError.txt"),
			OutputBlobURI:  to.Ptr("https://mystorageaccount.blob.core.windows.net/myscriptoutputcontainer/MyScriptoutput.txt"),
			Parameters: []*armhybridcompute.RunCommandInputParameter{
				{
					Name:  to.Ptr("param1"),
					Value: to.Ptr("value1"),
				},
				{
					Name:  to.Ptr("param2"),
					Value: to.Ptr("value2"),
				}},
			RunAsPassword: to.Ptr("<runAsPassword>"),
			RunAsUser:     to.Ptr("user1"),
			Source: &armhybridcompute.MachineRunCommandScriptSource{
				Script: to.Ptr("Write-Host Hello World!"),
			},
			TimeoutInSeconds: to.Ptr[int32](3600),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MachineRunCommand = armhybridcompute.MachineRunCommand{
	// 	Name: to.Ptr("myRunCommand"),
	// 	Type: to.Ptr("Microsoft.HybridCompute/machines/runcommands"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/Machines/myMachine/runcommands/myRunCommand"),
	// 	Location: to.Ptr("eastus2"),
	// 	Properties: &armhybridcompute.MachineRunCommandProperties{
	// 		AsyncExecution: to.Ptr(false),
	// 		ErrorBlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/mycontainer/MyScriptError.txt"),
	// 		OutputBlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/myscriptoutputcontainer/MyScriptoutput.txt"),
	// 		Parameters: []*armhybridcompute.RunCommandInputParameter{
	// 			{
	// 				Name: to.Ptr("param1"),
	// 				Value: to.Ptr("value1"),
	// 			},
	// 			{
	// 				Name: to.Ptr("param2"),
	// 				Value: to.Ptr("value2"),
	// 		}},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		RunAsUser: to.Ptr("user1"),
	// 		Source: &armhybridcompute.MachineRunCommandScriptSource{
	// 			Script: to.Ptr("Write-Host Hello World!"),
	// 		},
	// 		TimeoutInSeconds: to.Ptr[int32](3600),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b574e2a41acda14a90ef237006e8bbdda2b63c63/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-05-20-preview/examples/runCommand/RunCommands_Delete.json
func ExampleMachineRunCommandsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMachineRunCommandsClient().BeginDelete(ctx, "myResourceGroup", "myMachine", "myRunCommand", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b574e2a41acda14a90ef237006e8bbdda2b63c63/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-05-20-preview/examples/runCommand/RunCommands_Get.json
func ExampleMachineRunCommandsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMachineRunCommandsClient().Get(ctx, "myResourceGroup", "myMachine", "myRunCommand", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MachineRunCommand = armhybridcompute.MachineRunCommand{
	// 	Name: to.Ptr("myRunCommand"),
	// 	Type: to.Ptr("Microsoft.HybridCompute/machines/runcommands"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/Machines/myMachine/runcommands/myRunCommand"),
	// 	Location: to.Ptr("eastus2"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armhybridcompute.MachineRunCommandProperties{
	// 		AsyncExecution: to.Ptr(false),
	// 		InstanceView: &armhybridcompute.MachineRunCommandInstanceView{
	// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-15T20:48:41.464Z"); return t}()),
	// 			Error: to.Ptr(""),
	// 			ExecutionMessage: to.Ptr(""),
	// 			ExecutionState: to.Ptr(armhybridcompute.ExecutionStateSucceeded),
	// 			ExitCode: to.Ptr[int32](0),
	// 			Output: to.Ptr("Hello World"),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-15T20:48:41.464Z"); return t}()),
	// 		},
	// 		Parameters: []*armhybridcompute.RunCommandInputParameter{
	// 			{
	// 				Name: to.Ptr("param1"),
	// 				Value: to.Ptr("value1"),
	// 			},
	// 			{
	// 				Name: to.Ptr("param2"),
	// 				Value: to.Ptr("value2"),
	// 		}},
	// 		ProtectedParameters: []*armhybridcompute.RunCommandInputParameter{
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		RunAsUser: to.Ptr("user1"),
	// 		Source: &armhybridcompute.MachineRunCommandScriptSource{
	// 			Script: to.Ptr("Write-Host Hello World!"),
	// 		},
	// 		TimeoutInSeconds: to.Ptr[int32](3600),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b574e2a41acda14a90ef237006e8bbdda2b63c63/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-05-20-preview/examples/runCommand/RunCommands_List.json
func ExampleMachineRunCommandsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMachineRunCommandsClient().NewListPager("myResourceGroup", "myMachine", &armhybridcompute.MachineRunCommandsClientListOptions{Expand: nil})
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
		// page.MachineRunCommandsListResult = armhybridcompute.MachineRunCommandsListResult{
		// 	Value: []*armhybridcompute.MachineRunCommand{
		// 		{
		// 			Name: to.Ptr("myRunCommand_1"),
		// 			Type: to.Ptr("Microsoft.HybridCompute/machines/runcommands"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/Machines/myMachine/runcommands/myRunCommand_1"),
		// 			Location: to.Ptr("eastus2"),
		// 			Properties: &armhybridcompute.MachineRunCommandProperties{
		// 				AsyncExecution: to.Ptr(false),
		// 				InstanceView: &armhybridcompute.MachineRunCommandInstanceView{
		// 					EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-15T20:48:41.464Z"); return t}()),
		// 					Error: to.Ptr(""),
		// 					ExecutionMessage: to.Ptr(""),
		// 					ExecutionState: to.Ptr(armhybridcompute.ExecutionStateSucceeded),
		// 					ExitCode: to.Ptr[int32](0),
		// 					Output: to.Ptr("Hello World"),
		// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-15T20:48:41.464Z"); return t}()),
		// 				},
		// 				Parameters: []*armhybridcompute.RunCommandInputParameter{
		// 					{
		// 						Name: to.Ptr("param1"),
		// 						Value: to.Ptr("value1"),
		// 					},
		// 					{
		// 						Name: to.Ptr("param2"),
		// 						Value: to.Ptr("value2"),
		// 				}},
		// 				ProtectedParameters: []*armhybridcompute.RunCommandInputParameter{
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				RunAsUser: to.Ptr("user1"),
		// 				Source: &armhybridcompute.MachineRunCommandScriptSource{
		// 					Script: to.Ptr("Write-Host Hello World!"),
		// 				},
		// 				TimeoutInSeconds: to.Ptr[int32](3600),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myRunCommand_2"),
		// 			Type: to.Ptr("Microsoft.HybridCompute/machines/runcommands"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/Machines/myMachine/runcommands/myRunCommand_2"),
		// 			Location: to.Ptr("eastus2"),
		// 			Properties: &armhybridcompute.MachineRunCommandProperties{
		// 				AsyncExecution: to.Ptr(false),
		// 				InstanceView: &armhybridcompute.MachineRunCommandInstanceView{
		// 					EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-15T20:48:41.464Z"); return t}()),
		// 					Error: to.Ptr(""),
		// 					ExecutionMessage: to.Ptr(""),
		// 					ExecutionState: to.Ptr(armhybridcompute.ExecutionStateSucceeded),
		// 					ExitCode: to.Ptr[int32](0),
		// 					Output: to.Ptr("<some output>"),
		// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-15T20:48:41.464Z"); return t}()),
		// 				},
		// 				Parameters: []*armhybridcompute.RunCommandInputParameter{
		// 				},
		// 				ProtectedParameters: []*armhybridcompute.RunCommandInputParameter{
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				RunAsUser: to.Ptr("userA"),
		// 				Source: &armhybridcompute.MachineRunCommandScriptSource{
		// 					Script: to.Ptr("Get-Process | Where-Object { $_.CPU -gt 10000 }"),
		// 				},
		// 				TimeoutInSeconds: to.Ptr[int32](100),
		// 			},
		// 	}},
		// }
	}
}
