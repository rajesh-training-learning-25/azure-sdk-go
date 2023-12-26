//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armazurearcdata_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurearcdata/armazurearcdata"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/azurearcdata/resource-manager/Microsoft.AzureArcData/preview/2022-03-01-preview/examples/ListSubscriptionDataController.json
func ExampleDataControllersClient_NewListInSubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurearcdata.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDataControllersClient().NewListInSubscriptionPager(nil)
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
		// page.PageOfDataControllerResource = armazurearcdata.PageOfDataControllerResource{
		// 	Value: []*armazurearcdata.DataControllerResource{
		// 		{
		// 			Name: to.Ptr("testdataController1"),
		// 			Type: to.Ptr("Microsoft.AzureArcData/dataControllers"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.AzureArcData/dataControllers/testdataController1"),
		// 			SystemData: &armazurearcdata.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armazurearcdata.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armazurearcdata.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("northeurope"),
		// 			Tags: map[string]*string{
		// 				"mytag": to.Ptr("myval"),
		// 			},
		// 			ExtendedLocation: &armazurearcdata.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.ExtendedLocation/customLocations/arclocation"),
		// 				Type: to.Ptr(armazurearcdata.ExtendedLocationTypesCustomLocation),
		// 			},
		// 			Properties: &armazurearcdata.DataControllerProperties{
		// 				BasicLoginInformation: &armazurearcdata.BasicLoginInformation{
		// 					Username: to.Ptr("username"),
		// 				},
		// 				ClusterID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/connectedk8s"),
		// 				ExtensionID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/connectedk8s/providers/Microsoft.KubernetesConfiguration/extensions/extension"),
		// 				Infrastructure: to.Ptr(armazurearcdata.InfrastructureOnpremises),
		// 				LogAnalyticsWorkspaceConfig: &armazurearcdata.LogAnalyticsWorkspaceConfig{
		// 					WorkspaceID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 				},
		// 				LogsDashboardCredential: &armazurearcdata.BasicLoginInformation{
		// 					Username: to.Ptr("username"),
		// 				},
		// 				MetricsDashboardCredential: &armazurearcdata.BasicLoginInformation{
		// 					Username: to.Ptr("username"),
		// 				},
		// 				OnPremiseProperty: &armazurearcdata.OnPremiseProperty{
		// 					ID: to.Ptr("12345678-1234-1234-ab12-1a2b3c4d5e6f"),
		// 					PublicSigningKey: to.Ptr("publicOnPremSigningKey"),
		// 					SigningCertificateThumbprint: to.Ptr("Unique thumbprint returned to customer to verify the certificate they uploaded"),
		// 				},
		// 				UploadServicePrincipal: &armazurearcdata.UploadServicePrincipal{
		// 					Authority: to.Ptr("https://login.microsoftonline.com/"),
		// 					ClientID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 					TenantID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 				},
		// 				UploadWatermark: &armazurearcdata.UploadWatermark{
		// 					Logs: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 					Metrics: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 					Usages: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testdataController2"),
		// 			Type: to.Ptr("Microsoft.AzureArcData/dataControllers"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.AzureArcData/dataControllers/testdataController2"),
		// 			SystemData: &armazurearcdata.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armazurearcdata.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armazurearcdata.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("northeurope"),
		// 			Tags: map[string]*string{
		// 				"mytag": to.Ptr("myval"),
		// 			},
		// 			ExtendedLocation: &armazurearcdata.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.ExtendedLocation/customLocations/arclocation"),
		// 				Type: to.Ptr(armazurearcdata.ExtendedLocationTypesCustomLocation),
		// 			},
		// 			Properties: &armazurearcdata.DataControllerProperties{
		// 				BasicLoginInformation: &armazurearcdata.BasicLoginInformation{
		// 					Username: to.Ptr("username"),
		// 				},
		// 				ClusterID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/connectedk8s"),
		// 				ExtensionID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/connectedk8s/providers/Microsoft.KubernetesConfiguration/extensions/extension"),
		// 				LogAnalyticsWorkspaceConfig: &armazurearcdata.LogAnalyticsWorkspaceConfig{
		// 					WorkspaceID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 				},
		// 				LogsDashboardCredential: &armazurearcdata.BasicLoginInformation{
		// 					Username: to.Ptr("username"),
		// 				},
		// 				MetricsDashboardCredential: &armazurearcdata.BasicLoginInformation{
		// 					Username: to.Ptr("username"),
		// 				},
		// 				OnPremiseProperty: &armazurearcdata.OnPremiseProperty{
		// 					ID: to.Ptr("12345678-1234-1234-ab12-1a2b3c4d5e6f"),
		// 					PublicSigningKey: to.Ptr("publicOnPremSigningKey"),
		// 					SigningCertificateThumbprint: to.Ptr("Unique thumbprint returned to customer to verify the certificate they uploaded"),
		// 				},
		// 				UploadServicePrincipal: &armazurearcdata.UploadServicePrincipal{
		// 					Authority: to.Ptr("https://login.microsoftonline.com/"),
		// 					ClientID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 					TenantID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 				},
		// 				UploadWatermark: &armazurearcdata.UploadWatermark{
		// 					Logs: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 					Metrics: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 					Usages: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/azurearcdata/resource-manager/Microsoft.AzureArcData/preview/2022-03-01-preview/examples/ListByResourceGroupDataController.json
func ExampleDataControllersClient_NewListInGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurearcdata.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDataControllersClient().NewListInGroupPager("testrg", nil)
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
		// page.PageOfDataControllerResource = armazurearcdata.PageOfDataControllerResource{
		// 	Value: []*armazurearcdata.DataControllerResource{
		// 		{
		// 			Name: to.Ptr("testdataController1"),
		// 			Type: to.Ptr("Microsoft.AzureArcData/dataControllers"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.AzureArcData/dataControllers/testdataController1"),
		// 			SystemData: &armazurearcdata.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armazurearcdata.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armazurearcdata.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("northeurope"),
		// 			Tags: map[string]*string{
		// 				"mytag": to.Ptr("myval"),
		// 			},
		// 			ExtendedLocation: &armazurearcdata.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.ExtendedLocation/customLocations/arclocation"),
		// 				Type: to.Ptr(armazurearcdata.ExtendedLocationTypesCustomLocation),
		// 			},
		// 			Properties: &armazurearcdata.DataControllerProperties{
		// 				BasicLoginInformation: &armazurearcdata.BasicLoginInformation{
		// 					Username: to.Ptr("username"),
		// 				},
		// 				ClusterID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/connectedk8s"),
		// 				ExtensionID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/connectedk8s/providers/Microsoft.KubernetesConfiguration/extensions/extension"),
		// 				Infrastructure: to.Ptr(armazurearcdata.InfrastructureOnpremises),
		// 				LogAnalyticsWorkspaceConfig: &armazurearcdata.LogAnalyticsWorkspaceConfig{
		// 					WorkspaceID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 				},
		// 				LogsDashboardCredential: &armazurearcdata.BasicLoginInformation{
		// 					Username: to.Ptr("username"),
		// 				},
		// 				MetricsDashboardCredential: &armazurearcdata.BasicLoginInformation{
		// 					Username: to.Ptr("username"),
		// 				},
		// 				OnPremiseProperty: &armazurearcdata.OnPremiseProperty{
		// 					ID: to.Ptr("12345678-1234-1234-ab12-1a2b3c4d5e6f"),
		// 					PublicSigningKey: to.Ptr("publicOnPremSigningKey"),
		// 					SigningCertificateThumbprint: to.Ptr("Unique thumbprint returned to customer to verify the certificate they uploaded"),
		// 				},
		// 				UploadServicePrincipal: &armazurearcdata.UploadServicePrincipal{
		// 					Authority: to.Ptr("https://login.microsoftonline.com/"),
		// 					ClientID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 					TenantID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 				},
		// 				UploadWatermark: &armazurearcdata.UploadWatermark{
		// 					Logs: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 					Metrics: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 					Usages: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testdataController2"),
		// 			Type: to.Ptr("Microsoft.AzureArcData/dataControllers"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.AzureArcData/dataControllers/testdataController2"),
		// 			SystemData: &armazurearcdata.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armazurearcdata.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armazurearcdata.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("northeurope"),
		// 			Tags: map[string]*string{
		// 				"mytag": to.Ptr("myval"),
		// 			},
		// 			ExtendedLocation: &armazurearcdata.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.ExtendedLocation/customLocations/arclocation"),
		// 				Type: to.Ptr(armazurearcdata.ExtendedLocationTypesCustomLocation),
		// 			},
		// 			Properties: &armazurearcdata.DataControllerProperties{
		// 				BasicLoginInformation: &armazurearcdata.BasicLoginInformation{
		// 					Username: to.Ptr("username"),
		// 				},
		// 				ClusterID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/connectedk8s"),
		// 				ExtensionID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/connectedk8s/providers/Microsoft.KubernetesConfiguration/extensions/extension"),
		// 				Infrastructure: to.Ptr(armazurearcdata.InfrastructureOnpremises),
		// 				LogAnalyticsWorkspaceConfig: &armazurearcdata.LogAnalyticsWorkspaceConfig{
		// 					WorkspaceID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 				},
		// 				OnPremiseProperty: &armazurearcdata.OnPremiseProperty{
		// 					ID: to.Ptr("12345678-1234-1234-ab12-1a2b3c4d5e6f"),
		// 					PublicSigningKey: to.Ptr("publicOnPremSigningKey"),
		// 					SigningCertificateThumbprint: to.Ptr("Unique thumbprint returned to customer to verify the certificate they uploaded"),
		// 				},
		// 				UploadServicePrincipal: &armazurearcdata.UploadServicePrincipal{
		// 					Authority: to.Ptr("https://login.microsoftonline.com/"),
		// 					ClientID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 					TenantID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 				},
		// 				UploadWatermark: &armazurearcdata.UploadWatermark{
		// 					Logs: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 					Metrics: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 					Usages: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/azurearcdata/resource-manager/Microsoft.AzureArcData/preview/2022-03-01-preview/examples/CreateOrUpdateDataController.json
func ExampleDataControllersClient_BeginPutDataController() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurearcdata.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDataControllersClient().BeginPutDataController(ctx, "testrg", "testdataController", armazurearcdata.DataControllerResource{
		Location: to.Ptr("northeurope"),
		Tags: map[string]*string{
			"mytag": to.Ptr("myval"),
		},
		ExtendedLocation: &armazurearcdata.ExtendedLocation{
			Name: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.ExtendedLocation/customLocations/arclocation"),
			Type: to.Ptr(armazurearcdata.ExtendedLocationTypesCustomLocation),
		},
		Properties: &armazurearcdata.DataControllerProperties{
			BasicLoginInformation: &armazurearcdata.BasicLoginInformation{
				Password: to.Ptr("********"),
				Username: to.Ptr("username"),
			},
			ClusterID:      to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/connectedk8s"),
			ExtensionID:    to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/connectedk8s/providers/Microsoft.KubernetesConfiguration/extensions/extension"),
			Infrastructure: to.Ptr(armazurearcdata.InfrastructureOnpremises),
			LogAnalyticsWorkspaceConfig: &armazurearcdata.LogAnalyticsWorkspaceConfig{
				PrimaryKey:  to.Ptr("********"),
				WorkspaceID: to.Ptr("00000000-1111-2222-3333-444444444444"),
			},
			LogsDashboardCredential: &armazurearcdata.BasicLoginInformation{
				Password: to.Ptr("********"),
				Username: to.Ptr("username"),
			},
			MetricsDashboardCredential: &armazurearcdata.BasicLoginInformation{
				Password: to.Ptr("********"),
				Username: to.Ptr("username"),
			},
			OnPremiseProperty: &armazurearcdata.OnPremiseProperty{
				ID:               to.Ptr("12345678-1234-1234-ab12-1a2b3c4d5e6f"),
				PublicSigningKey: to.Ptr("publicOnPremSigningKey"),
			},
			UploadServicePrincipal: &armazurearcdata.UploadServicePrincipal{
				Authority:    to.Ptr("https://login.microsoftonline.com/"),
				ClientID:     to.Ptr("00000000-1111-2222-3333-444444444444"),
				ClientSecret: to.Ptr("********"),
				TenantID:     to.Ptr("00000000-1111-2222-3333-444444444444"),
			},
			UploadWatermark: &armazurearcdata.UploadWatermark{
				Logs:    to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t }()),
				Metrics: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t }()),
				Usages:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t }()),
			},
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
	// res.DataControllerResource = armazurearcdata.DataControllerResource{
	// 	Name: to.Ptr("testdataController"),
	// 	Type: to.Ptr("Microsoft.AzureArcData/dataControllers"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.AzureArcData/dataControllers/testdataController"),
	// 	SystemData: &armazurearcdata.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armazurearcdata.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armazurearcdata.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("northeurope"),
	// 	Tags: map[string]*string{
	// 		"mytag": to.Ptr("myval"),
	// 	},
	// 	ExtendedLocation: &armazurearcdata.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.ExtendedLocation/customLocations/arclocation"),
	// 		Type: to.Ptr(armazurearcdata.ExtendedLocationTypesCustomLocation),
	// 	},
	// 	Properties: &armazurearcdata.DataControllerProperties{
	// 		BasicLoginInformation: &armazurearcdata.BasicLoginInformation{
	// 			Username: to.Ptr("username"),
	// 		},
	// 		ClusterID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/connectedk8s"),
	// 		ExtensionID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/connectedk8s/providers/Microsoft.KubernetesConfiguration/extensions/extension"),
	// 		Infrastructure: to.Ptr(armazurearcdata.InfrastructureOnpremises),
	// 		LogAnalyticsWorkspaceConfig: &armazurearcdata.LogAnalyticsWorkspaceConfig{
	// 			WorkspaceID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 		},
	// 		LogsDashboardCredential: &armazurearcdata.BasicLoginInformation{
	// 			Username: to.Ptr("username"),
	// 		},
	// 		MetricsDashboardCredential: &armazurearcdata.BasicLoginInformation{
	// 			Username: to.Ptr("username"),
	// 		},
	// 		OnPremiseProperty: &armazurearcdata.OnPremiseProperty{
	// 			ID: to.Ptr("12345678-1234-1234-ab12-1a2b3c4d5e6f"),
	// 			PublicSigningKey: to.Ptr("publicOnPremSigningKey"),
	// 			SigningCertificateThumbprint: to.Ptr("Unique thumbprint returned to customer to verify the certificate being uploaded"),
	// 		},
	// 		UploadServicePrincipal: &armazurearcdata.UploadServicePrincipal{
	// 			Authority: to.Ptr("https://login.microsoftonline.com/"),
	// 			ClientID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 			TenantID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 		},
	// 		UploadWatermark: &armazurearcdata.UploadWatermark{
	// 			Logs: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 			Metrics: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 			Usages: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/azurearcdata/resource-manager/Microsoft.AzureArcData/preview/2022-03-01-preview/examples/DeleteDataController.json
func ExampleDataControllersClient_BeginDeleteDataController() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurearcdata.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDataControllersClient().BeginDeleteDataController(ctx, "testrg", "testdataController", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/azurearcdata/resource-manager/Microsoft.AzureArcData/preview/2022-03-01-preview/examples/GetDataController.json
func ExampleDataControllersClient_GetDataController() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurearcdata.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataControllersClient().GetDataController(ctx, "testrg", "testdataController", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DataControllerResource = armazurearcdata.DataControllerResource{
	// 	Name: to.Ptr("testdataController"),
	// 	Type: to.Ptr("Microsoft.AzureArcData/dataControllers"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.AzureArcData/dataControllers/testdataController"),
	// 	SystemData: &armazurearcdata.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armazurearcdata.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armazurearcdata.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("northeurope"),
	// 	Tags: map[string]*string{
	// 		"mytag": to.Ptr("myval"),
	// 	},
	// 	ExtendedLocation: &armazurearcdata.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.ExtendedLocation/customLocations/arclocation"),
	// 		Type: to.Ptr(armazurearcdata.ExtendedLocationTypesCustomLocation),
	// 	},
	// 	Properties: &armazurearcdata.DataControllerProperties{
	// 		BasicLoginInformation: &armazurearcdata.BasicLoginInformation{
	// 			Username: to.Ptr("username"),
	// 		},
	// 		ClusterID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/connectedk8s"),
	// 		ExtensionID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/connectedk8s/providers/Microsoft.KubernetesConfiguration/extensions/extension"),
	// 		Infrastructure: to.Ptr(armazurearcdata.InfrastructureOnpremises),
	// 		LogAnalyticsWorkspaceConfig: &armazurearcdata.LogAnalyticsWorkspaceConfig{
	// 			WorkspaceID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 		},
	// 		LogsDashboardCredential: &armazurearcdata.BasicLoginInformation{
	// 			Username: to.Ptr("username"),
	// 		},
	// 		MetricsDashboardCredential: &armazurearcdata.BasicLoginInformation{
	// 			Username: to.Ptr("username"),
	// 		},
	// 		OnPremiseProperty: &armazurearcdata.OnPremiseProperty{
	// 			ID: to.Ptr("12345678-1234-1234-ab12-1a2b3c4d5e6f"),
	// 			PublicSigningKey: to.Ptr("publicOnPremSigningKey"),
	// 			SigningCertificateThumbprint: to.Ptr("Unique thumbprint returned to customer to verify the certificate they uploaded"),
	// 		},
	// 		UploadServicePrincipal: &armazurearcdata.UploadServicePrincipal{
	// 			Authority: to.Ptr("https://login.microsoftonline.com/"),
	// 			ClientID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 			TenantID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 		},
	// 		UploadWatermark: &armazurearcdata.UploadWatermark{
	// 			Logs: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 			Metrics: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 			Usages: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/azurearcdata/resource-manager/Microsoft.AzureArcData/preview/2022-03-01-preview/examples/UpdateDataController.json
func ExampleDataControllersClient_BeginPatchDataController() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurearcdata.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDataControllersClient().BeginPatchDataController(ctx, "testrg", "testdataController1", armazurearcdata.DataControllerUpdate{
		Tags: map[string]*string{
			"mytag": to.Ptr("myval"),
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
	// res.DataControllerResource = armazurearcdata.DataControllerResource{
	// 	Name: to.Ptr("testdataController1"),
	// 	Type: to.Ptr("Microsoft.AzureArcData/dataControllers"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.AzureArcData/dataControllers/testdataController1"),
	// 	SystemData: &armazurearcdata.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armazurearcdata.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armazurearcdata.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("northeurope"),
	// 	Tags: map[string]*string{
	// 		"mytag": to.Ptr("myval"),
	// 	},
	// 	ExtendedLocation: &armazurearcdata.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.ExtendedLocation/customLocations/arclocation"),
	// 		Type: to.Ptr(armazurearcdata.ExtendedLocationTypesCustomLocation),
	// 	},
	// 	Properties: &armazurearcdata.DataControllerProperties{
	// 		BasicLoginInformation: &armazurearcdata.BasicLoginInformation{
	// 			Username: to.Ptr("username"),
	// 		},
	// 		ClusterID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/connectedk8s"),
	// 		ExtensionID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/connectedk8s/providers/Microsoft.KubernetesConfiguration/extensions/extension"),
	// 		Infrastructure: to.Ptr(armazurearcdata.InfrastructureOnpremises),
	// 		LogAnalyticsWorkspaceConfig: &armazurearcdata.LogAnalyticsWorkspaceConfig{
	// 			WorkspaceID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 		},
	// 		LogsDashboardCredential: &armazurearcdata.BasicLoginInformation{
	// 			Username: to.Ptr("username"),
	// 		},
	// 		MetricsDashboardCredential: &armazurearcdata.BasicLoginInformation{
	// 			Username: to.Ptr("username"),
	// 		},
	// 		OnPremiseProperty: &armazurearcdata.OnPremiseProperty{
	// 			ID: to.Ptr("12345678-1234-1234-ab12-1a2b3c4d5e6f"),
	// 			PublicSigningKey: to.Ptr("publicOnPremSigningKey"),
	// 			SigningCertificateThumbprint: to.Ptr("Unique thumbprint returned to customer to verify the certificate they uploaded"),
	// 		},
	// 		UploadServicePrincipal: &armazurearcdata.UploadServicePrincipal{
	// 			Authority: to.Ptr("https://login.microsoftonline.com/"),
	// 			ClientID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 			TenantID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 		},
	// 		UploadWatermark: &armazurearcdata.UploadWatermark{
	// 			Logs: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 			Metrics: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 			Usages: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 		},
	// 	},
	// }
}
