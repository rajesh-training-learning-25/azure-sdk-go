//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatafactory_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory"
)

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_ListByFactory.json
func ExampleIntegrationRuntimesClient_ListByFactory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewIntegrationRuntimesClient("<subscription-id>", cred, nil)
	pager := client.ListByFactory("<resource-group-name>",
		"<factory-name>",
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

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_Create.json
func ExampleIntegrationRuntimesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewIntegrationRuntimesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<integration-runtime-name>",
		armdatafactory.IntegrationRuntimeResource{
			Properties: &armdatafactory.SelfHostedIntegrationRuntime{
				Type:        armdatafactory.IntegrationRuntimeType("SelfHosted").ToPtr(),
				Description: to.StringPtr("<description>"),
			},
		},
		&armdatafactory.IntegrationRuntimesClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.IntegrationRuntimesClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_Get.json
func ExampleIntegrationRuntimesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewIntegrationRuntimesClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<integration-runtime-name>",
		&armdatafactory.IntegrationRuntimesClientGetOptions{IfNoneMatch: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.IntegrationRuntimesClientGetResult)
}

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_Update.json
func ExampleIntegrationRuntimesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewIntegrationRuntimesClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<integration-runtime-name>",
		armdatafactory.UpdateIntegrationRuntimeRequest{
			AutoUpdate:        armdatafactory.IntegrationRuntimeAutoUpdate("Off").ToPtr(),
			UpdateDelayOffset: to.StringPtr("<update-delay-offset>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.IntegrationRuntimesClientUpdateResult)
}

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_Delete.json
func ExampleIntegrationRuntimesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewIntegrationRuntimesClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<integration-runtime-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_GetStatus.json
func ExampleIntegrationRuntimesClient_GetStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewIntegrationRuntimesClient("<subscription-id>", cred, nil)
	res, err := client.GetStatus(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<integration-runtime-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.IntegrationRuntimesClientGetStatusResult)
}

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_ListOutboundNetworkDependenciesEndpoints.json
func ExampleIntegrationRuntimesClient_ListOutboundNetworkDependenciesEndpoints() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewIntegrationRuntimesClient("<subscription-id>", cred, nil)
	res, err := client.ListOutboundNetworkDependenciesEndpoints(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<integration-runtime-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.IntegrationRuntimesClientListOutboundNetworkDependenciesEndpointsResult)
}

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_GetConnectionInfo.json
func ExampleIntegrationRuntimesClient_GetConnectionInfo() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewIntegrationRuntimesClient("<subscription-id>", cred, nil)
	res, err := client.GetConnectionInfo(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<integration-runtime-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.IntegrationRuntimesClientGetConnectionInfoResult)
}

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_RegenerateAuthKey.json
func ExampleIntegrationRuntimesClient_RegenerateAuthKey() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewIntegrationRuntimesClient("<subscription-id>", cred, nil)
	res, err := client.RegenerateAuthKey(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<integration-runtime-name>",
		armdatafactory.IntegrationRuntimeRegenerateKeyParameters{
			KeyName: armdatafactory.IntegrationRuntimeAuthKeyName("authKey2").ToPtr(),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.IntegrationRuntimesClientRegenerateAuthKeyResult)
}

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_ListAuthKeys.json
func ExampleIntegrationRuntimesClient_ListAuthKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewIntegrationRuntimesClient("<subscription-id>", cred, nil)
	res, err := client.ListAuthKeys(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<integration-runtime-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.IntegrationRuntimesClientListAuthKeysResult)
}

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_Start.json
func ExampleIntegrationRuntimesClient_BeginStart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewIntegrationRuntimesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginStart(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<integration-runtime-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.IntegrationRuntimesClientStartResult)
}

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_Stop.json
func ExampleIntegrationRuntimesClient_BeginStop() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewIntegrationRuntimesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginStop(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<integration-runtime-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_SyncCredentials.json
func ExampleIntegrationRuntimesClient_SyncCredentials() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewIntegrationRuntimesClient("<subscription-id>", cred, nil)
	_, err = client.SyncCredentials(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<integration-runtime-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_GetMonitoringData.json
func ExampleIntegrationRuntimesClient_GetMonitoringData() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewIntegrationRuntimesClient("<subscription-id>", cred, nil)
	res, err := client.GetMonitoringData(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<integration-runtime-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.IntegrationRuntimesClientGetMonitoringDataResult)
}

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_Upgrade.json
func ExampleIntegrationRuntimesClient_Upgrade() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewIntegrationRuntimesClient("<subscription-id>", cred, nil)
	_, err = client.Upgrade(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<integration-runtime-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_RemoveLinks.json
func ExampleIntegrationRuntimesClient_RemoveLinks() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewIntegrationRuntimesClient("<subscription-id>", cred, nil)
	_, err = client.RemoveLinks(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<integration-runtime-name>",
		armdatafactory.LinkedIntegrationRuntimeRequest{
			LinkedFactoryName: to.StringPtr("<linked-factory-name>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_CreateLinkedIntegrationRuntime.json
func ExampleIntegrationRuntimesClient_CreateLinkedIntegrationRuntime() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewIntegrationRuntimesClient("<subscription-id>", cred, nil)
	res, err := client.CreateLinkedIntegrationRuntime(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<integration-runtime-name>",
		armdatafactory.CreateLinkedIntegrationRuntimeRequest{
			Name:                to.StringPtr("<name>"),
			DataFactoryLocation: to.StringPtr("<data-factory-location>"),
			DataFactoryName:     to.StringPtr("<data-factory-name>"),
			SubscriptionID:      to.StringPtr("<subscription-id>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.IntegrationRuntimesClientCreateLinkedIntegrationRuntimeResult)
}
