//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/58be094c6b365f8d4d73a91e293dfb4818e57cf6/specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/Replicas_Get.json
func ExampleContainerAppsRevisionReplicasClient_GetReplica() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewContainerAppsRevisionReplicasClient().GetReplica(ctx, "workerapps-rg-xj", "myapp", "myapp--0wlqy09", "myapp--0wlqy09-5d9774cff-5wnd8", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Replica = armappcontainers.Replica{
	// 	Name: to.Ptr("myapp--0wlqy09-5d9774cff-5wnd8"),
	// 	Type: to.Ptr("Microsoft.Web/containerapps/revisions/replicas"),
	// 	ID: to.Ptr("/subscriptions/651f8027-33e8-4ec4-97b4-f6e9f3dc8744/resourceGroups/workerapps-rg-xj/providers/Microsoft.App/containerApps/myapp/revisions/myapp--0wlqy09/replicas/myapp--0wlqy09-5d9774cff-5wnd8"),
	// 	Properties: &armappcontainers.ReplicaProperties{
	// 		Containers: []*armappcontainers.ReplicaContainer{
	// 			{
	// 				Name: to.Ptr("hello92"),
	// 				ContainerID: to.Ptr("containerd://6bac7bb3afed1c704b5fe563c34c0ecf59ac30c766bb73488f7fa552dc42ee54"),
	// 				ExecEndpoint: to.Ptr("testExecEndpoint"),
	// 				LogStreamEndpoint: to.Ptr("testLogStreamEndpoint"),
	// 				Ready: to.Ptr(true),
	// 				RestartCount: to.Ptr[int32](0),
	// 				RunningState: to.Ptr(armappcontainers.ContainerAppContainerRunningStateRunning),
	// 				RunningStateDetails: to.Ptr("testDetail"),
	// 				Started: to.Ptr(true),
	// 		}},
	// 		CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-25T19:42:45Z"); return t}()),
	// 		InitContainers: []*armappcontainers.ReplicaContainer{
	// 		},
	// 		RunningState: to.Ptr(armappcontainers.ContainerAppReplicaRunningStateRunning),
	// 		RunningStateDetails: to.Ptr("testDetail"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/58be094c6b365f8d4d73a91e293dfb4818e57cf6/specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/Replicas_List.json
func ExampleContainerAppsRevisionReplicasClient_ListReplicas() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewContainerAppsRevisionReplicasClient().ListReplicas(ctx, "workerapps-rg-xj", "myapp", "myapp--0wlqy09", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ReplicaCollection = armappcontainers.ReplicaCollection{
	// 	Value: []*armappcontainers.Replica{
	// 		{
	// 			Name: to.Ptr("myapp--0wlqy09-5d9774cff-5wnd8"),
	// 			Type: to.Ptr("Microsoft.Web/containerapps/revisions/replicas"),
	// 			ID: to.Ptr("/subscriptions/651f8027-33e8-4ec4-97b4-f6e9f3dc8744/resourceGroups/workerapps-rg-xj/providers/Microsoft.App/containerApps/myapp/revisions/myapp--0wlqy09/replicas/myapp--0wlqy09-5d9774cff-5wnd8"),
	// 			Properties: &armappcontainers.ReplicaProperties{
	// 				Containers: []*armappcontainers.ReplicaContainer{
	// 					{
	// 						Name: to.Ptr("hello92"),
	// 						ContainerID: to.Ptr("containerd://6bac7bb3afed1c704b5fe563c34c0ecf59ac30c766bb73488f7fa552dc42ee54"),
	// 						ExecEndpoint: to.Ptr("testExecEndpoint"),
	// 						LogStreamEndpoint: to.Ptr("testLogStreamEndpoint"),
	// 						Ready: to.Ptr(true),
	// 						RestartCount: to.Ptr[int32](0),
	// 						RunningState: to.Ptr(armappcontainers.ContainerAppContainerRunningStateRunning),
	// 						RunningStateDetails: to.Ptr("testDetail"),
	// 						Started: to.Ptr(true),
	// 				}},
	// 				CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-25T19:42:45Z"); return t}()),
	// 				InitContainers: []*armappcontainers.ReplicaContainer{
	// 				},
	// 				RunningState: to.Ptr(armappcontainers.ContainerAppReplicaRunningStateRunning),
	// 				RunningStateDetails: to.Ptr("testDetail"),
	// 			},
	// 	}},
	// }
}
