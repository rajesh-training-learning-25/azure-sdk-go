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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/58be094c6b365f8d4d73a91e293dfb4818e57cf6/specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/ContainerAppsDiagnostics_List.json
func ExampleContainerAppsDiagnosticsClient_NewListDetectorsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewContainerAppsDiagnosticsClient().NewListDetectorsPager("mikono-workerapp-test-rg", "mikono-capp-stage1", nil)
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
		// page.DiagnosticsCollection = armappcontainers.DiagnosticsCollection{
		// 	Value: []*armappcontainers.Diagnostics{
		// 		{
		// 			Name: to.Ptr("cappContainerAppAvailabilityMetrics"),
		// 			Type: to.Ptr("Microsoft.App/containerapps/detectors"),
		// 			ID: to.Ptr("/subscriptions/f07f3711-b45e-40fe-a941-4e6d93f851e6/resourceGroups/mikono-workerapp-test-rg/providers/Microsoft.App/containerApps/mikono-capp-stage1/detectors/cappContainerAppAvailabilityMetrics"),
		// 			Properties: &armappcontainers.DiagnosticsProperties{
		// 				Dataset: []*armappcontainers.DiagnosticsDataAPIResponse{
		// 				},
		// 				Metadata: &armappcontainers.DiagnosticsDefinition{
		// 					Name: to.Ptr("Availability Metrics for Container Apps"),
		// 					Type: to.Ptr("Analysis"),
		// 					Author: to.Ptr(""),
		// 					Category: to.Ptr("Availability and Performance"),
		// 					ID: to.Ptr("cappContainerAppAvailabilityMetrics"),
		// 					Score: to.Ptr[float32](0),
		// 					SupportTopicList: []*armappcontainers.DiagnosticSupportTopic{
		// 					},
		// 				},
		// 				Status: &armappcontainers.DiagnosticsStatus{
		// 					StatusID: to.Ptr[int32](4),
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/58be094c6b365f8d4d73a91e293dfb4818e57cf6/specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/ContainerAppsDiagnostics_Get.json
func ExampleContainerAppsDiagnosticsClient_GetDetector() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewContainerAppsDiagnosticsClient().GetDetector(ctx, "mikono-workerapp-test-rg", "mikono-capp-stage1", "cappcontainerappnetworkIO", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Diagnostics = armappcontainers.Diagnostics{
	// 	Name: to.Ptr("cappcontainerappnetworkIO"),
	// 	Type: to.Ptr("Microsoft.App/containerapps/detectors"),
	// 	ID: to.Ptr("/subscriptions/f07f3711-b45e-40fe-a941-4e6d93f851e6/resourceGroups/mikono-workerapp-test-rg/providers/Microsoft.App/containerApps/mikono-capp-stage1/detectors/cappcontainerappnetworkIO"),
	// 	Properties: &armappcontainers.DiagnosticsProperties{
	// 		Dataset: []*armappcontainers.DiagnosticsDataAPIResponse{
	// 			{
	// 				RenderingProperties: &armappcontainers.DiagnosticRendering{
	// 					Type: to.Ptr[int32](8),
	// 					Description: to.Ptr(""),
	// 					IsVisible: to.Ptr(true),
	// 					Title: to.Ptr("Container Apps Network Inbound "),
	// 				},
	// 				Table: &armappcontainers.DiagnosticDataTableResponseObject{
	// 					Columns: []*armappcontainers.DiagnosticDataTableResponseColumn{
	// 						{
	// 							ColumnName: to.Ptr("TimeStamp"),
	// 							DataType: to.Ptr("DateTime"),
	// 						},
	// 						{
	// 							ColumnName: to.Ptr("Metric"),
	// 							DataType: to.Ptr("String"),
	// 						},
	// 						{
	// 							ColumnName: to.Ptr("Average"),
	// 							DataType: to.Ptr("Double"),
	// 					}},
	// 					Rows: []any{
	// 						[]any{
	// 							"2022-03-15T21:35:00",
	// 							"RxBytes",
	// 							float64(0),
	// 					}},
	// 					TableName: to.Ptr(""),
	// 				},
	// 		}},
	// 		Metadata: &armappcontainers.DiagnosticsDefinition{
	// 			Name: to.Ptr("Container App Network Inbound and Outbound"),
	// 			Type: to.Ptr("Detector"),
	// 			Description: to.Ptr("This detector shows the Container App Network Inbound and Outbound."),
	// 			Author: to.Ptr(""),
	// 			Category: to.Ptr("Availability and Performance"),
	// 			ID: to.Ptr("cappcontainerappnetworkIO"),
	// 			Score: to.Ptr[float32](0),
	// 			SupportTopicList: []*armappcontainers.DiagnosticSupportTopic{
	// 			},
	// 		},
	// 		Status: &armappcontainers.DiagnosticsStatus{
	// 			StatusID: to.Ptr[int32](3),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/58be094c6b365f8d4d73a91e293dfb4818e57cf6/specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/Revisions_List.json
func ExampleContainerAppsDiagnosticsClient_NewListRevisionsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewContainerAppsDiagnosticsClient().NewListRevisionsPager("rg", "testcontainerApp0", &armappcontainers.ContainerAppsDiagnosticsClientListRevisionsOptions{Filter: nil})
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
		// page.RevisionCollection = armappcontainers.RevisionCollection{
		// 	Value: []*armappcontainers.Revision{
		// 		{
		// 			Name: to.Ptr("testcontainerApp0-pjxhsye"),
		// 			Type: to.Ptr("Microsoft.App/containerApps/revisions"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/containerApps/testcontainerApp0/revisions/testcontainerApp0-pjxhsye"),
		// 			Properties: &armappcontainers.RevisionProperties{
		// 				Active: to.Ptr(true),
		// 				CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-24T21:24:22+00:00"); return t}()),
		// 				Fqdn: to.Ptr("testcontainerApp0-pjxhsye.politehill-ab123456.eastus.azurecontainerapps.io"),
		// 				LastActiveTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-24T21:24:22+00:00"); return t}()),
		// 				Replicas: to.Ptr[int32](1),
		// 				Template: &armappcontainers.Template{
		// 					Containers: []*armappcontainers.Container{
		// 						{
		// 							Name: to.Ptr("testcontainerApp0"),
		// 							Image: to.Ptr("repo/testcontainerApp0:v2"),
		// 							Resources: &armappcontainers.ContainerResources{
		// 								CPU: to.Ptr[float64](0.2),
		// 								Memory: to.Ptr("100Mi"),
		// 							},
		// 					}},
		// 					Scale: &armappcontainers.Scale{
		// 						MaxReplicas: to.Ptr[int32](5),
		// 						MinReplicas: to.Ptr[int32](1),
		// 						Rules: []*armappcontainers.ScaleRule{
		// 							{
		// 								Name: to.Ptr("httpscalingrule"),
		// 								HTTP: &armappcontainers.HTTPScaleRule{
		// 									Metadata: map[string]*string{
		// 										"concurrentRequests": to.Ptr("50"),
		// 									},
		// 								},
		// 						}},
		// 					},
		// 				},
		// 				TrafficWeight: to.Ptr[int32](80),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/58be094c6b365f8d4d73a91e293dfb4818e57cf6/specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/Revisions_Get.json
func ExampleContainerAppsDiagnosticsClient_GetRevision() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewContainerAppsDiagnosticsClient().GetRevision(ctx, "rg", "testcontainerApp0", "testcontainerApp0-pjxhsye", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Revision = armappcontainers.Revision{
	// 	Name: to.Ptr("testcontainerApp0-pjxhsye"),
	// 	Type: to.Ptr("Microsoft.App/containerApps/revisions"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.AppcontainerApps/testcontainerApp0/revisions/testcontainerApp0-pjxhsye"),
	// 	Properties: &armappcontainers.RevisionProperties{
	// 		Active: to.Ptr(true),
	// 		CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-24T21:24:22+00:00"); return t}()),
	// 		Fqdn: to.Ptr("testcontainerApp0-pjxhsye.politehill-ab123456.eastus.azurecontainerapps.io"),
	// 		LastActiveTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-24T21:24:22+00:00"); return t}()),
	// 		Replicas: to.Ptr[int32](1),
	// 		RunningState: to.Ptr(armappcontainers.RevisionRunningStateRunning),
	// 		Template: &armappcontainers.Template{
	// 			Containers: []*armappcontainers.Container{
	// 				{
	// 					Name: to.Ptr("testcontainerApp0"),
	// 					Image: to.Ptr("repo/testcontainerApp0:v2"),
	// 					Resources: &armappcontainers.ContainerResources{
	// 						CPU: to.Ptr[float64](0.2),
	// 						Memory: to.Ptr("100Mi"),
	// 					},
	// 			}},
	// 			Scale: &armappcontainers.Scale{
	// 				MaxReplicas: to.Ptr[int32](5),
	// 				MinReplicas: to.Ptr[int32](1),
	// 				Rules: []*armappcontainers.ScaleRule{
	// 					{
	// 						Name: to.Ptr("httpscalingrule"),
	// 						HTTP: &armappcontainers.HTTPScaleRule{
	// 							Metadata: map[string]*string{
	// 								"concurrentRequests": to.Ptr("50"),
	// 							},
	// 						},
	// 				}},
	// 			},
	// 		},
	// 		TrafficWeight: to.Ptr[int32](80),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/58be094c6b365f8d4d73a91e293dfb4818e57cf6/specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/ContainerApps_Get.json
func ExampleContainerAppsDiagnosticsClient_GetRoot() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewContainerAppsDiagnosticsClient().GetRoot(ctx, "rg", "testcontainerApp0", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ContainerApp = armappcontainers.ContainerApp{
	// 	Name: to.Ptr("testcontainerApp0"),
	// 	Type: to.Ptr("Microsoft.App/containerApps"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/containerApps/testcontainerApp0"),
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armappcontainers.ContainerAppProperties{
	// 		Configuration: &armappcontainers.Configuration{
	// 			Dapr: &armappcontainers.Dapr{
	// 				AppPort: to.Ptr[int32](3000),
	// 				AppProtocol: to.Ptr(armappcontainers.AppProtocolHTTP),
	// 				EnableAPILogging: to.Ptr(true),
	// 				Enabled: to.Ptr(true),
	// 				HTTPMaxRequestSize: to.Ptr[int32](10),
	// 				HTTPReadBufferSize: to.Ptr[int32](30),
	// 				LogLevel: to.Ptr(armappcontainers.LogLevelDebug),
	// 			},
	// 			Ingress: &armappcontainers.Ingress{
	// 				CustomDomains: []*armappcontainers.CustomDomain{
	// 					{
	// 						Name: to.Ptr("www.my-name.com"),
	// 						BindingType: to.Ptr(armappcontainers.BindingTypeSniEnabled),
	// 						CertificateID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube/certificates/my-certificate-for-my-name-dot-com"),
	// 					},
	// 					{
	// 						Name: to.Ptr("www.my--other-name.com"),
	// 						BindingType: to.Ptr(armappcontainers.BindingTypeSniEnabled),
	// 						CertificateID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube/certificates/my-certificate-for-my-other-name-dot-com"),
	// 				}},
	// 				External: to.Ptr(true),
	// 				Fqdn: to.Ptr("testcontainerApp0.demokube-t24clv0g.eastus.containerApps.k4apps.io"),
	// 				IPSecurityRestrictions: []*armappcontainers.IPSecurityRestrictionRule{
	// 					{
	// 						Name: to.Ptr("Allow work IP A subnet"),
	// 						Description: to.Ptr("Allowing all IP's within the subnet below to access containerapp"),
	// 						Action: to.Ptr(armappcontainers.ActionAllow),
	// 						IPAddressRange: to.Ptr("192.168.1.1/32"),
	// 					},
	// 					{
	// 						Name: to.Ptr("Allow work IP B subnet"),
	// 						Description: to.Ptr("Allowing all IP's within the subnet below to access containerapp"),
	// 						Action: to.Ptr(armappcontainers.ActionAllow),
	// 						IPAddressRange: to.Ptr("192.168.1.1/8"),
	// 				}},
	// 				StickySessions: &armappcontainers.IngressStickySessions{
	// 					Affinity: to.Ptr(armappcontainers.AffinitySticky),
	// 				},
	// 				TargetPort: to.Ptr[int32](3000),
	// 				Traffic: []*armappcontainers.TrafficWeight{
	// 					{
	// 						RevisionName: to.Ptr("testcontainerApp0-ab1234"),
	// 						Weight: to.Ptr[int32](80),
	// 					},
	// 					{
	// 						Label: to.Ptr("staging"),
	// 						RevisionName: to.Ptr("testcontainerApp0-ab4321"),
	// 						Weight: to.Ptr[int32](20),
	// 				}},
	// 				Transport: to.Ptr(armappcontainers.IngressTransportMethodAuto),
	// 			},
	// 			MaxInactiveRevisions: to.Ptr[int32](10),
	// 			Service: &armappcontainers.Service{
	// 				Type: to.Ptr("redis"),
	// 			},
	// 		},
	// 		EventStreamEndpoint: to.Ptr("testEndpoint"),
	// 		LatestReadyRevisionName: to.Ptr("testcontainerApp0-pjxhsye"),
	// 		LatestRevisionFqdn: to.Ptr("testcontainerApp0-pjxhsye.demokube-t24clv0g.eastus.containerApps.k4apps.io"),
	// 		ManagedEnvironmentID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube"),
	// 		ProvisioningState: to.Ptr(armappcontainers.ContainerAppProvisioningStateSucceeded),
	// 		Template: &armappcontainers.Template{
	// 			Containers: []*armappcontainers.Container{
	// 				{
	// 					Name: to.Ptr("testcontainerApp0"),
	// 					Image: to.Ptr("repo/testcontainerApp0:v4"),
	// 					Resources: &armappcontainers.ContainerResources{
	// 						CPU: to.Ptr[float64](0.2),
	// 						Memory: to.Ptr("100Mi"),
	// 					},
	// 			}},
	// 			InitContainers: []*armappcontainers.InitContainer{
	// 				{
	// 					Name: to.Ptr("testinitcontainerApp0"),
	// 					Image: to.Ptr("repo/testcontainerApp0:v4"),
	// 					Resources: &armappcontainers.ContainerResources{
	// 						CPU: to.Ptr[float64](0.2),
	// 						Memory: to.Ptr("100Mi"),
	// 					},
	// 			}},
	// 			Scale: &armappcontainers.Scale{
	// 				MaxReplicas: to.Ptr[int32](5),
	// 				MinReplicas: to.Ptr[int32](1),
	// 				Rules: []*armappcontainers.ScaleRule{
	// 					{
	// 						Name: to.Ptr("httpscalingrule"),
	// 						HTTP: &armappcontainers.HTTPScaleRule{
	// 							Metadata: map[string]*string{
	// 								"concurrentRequests": to.Ptr("50"),
	// 							},
	// 						},
	// 				}},
	// 			},
	// 			ServiceBinds: []*armappcontainers.ServiceBind{
	// 				{
	// 					Name: to.Ptr("service"),
	// 					ServiceID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/containerApps/service"),
	// 			}},
	// 		},
	// 		WorkloadProfileName: to.Ptr("My-GP-01"),
	// 	},
	// }
}
