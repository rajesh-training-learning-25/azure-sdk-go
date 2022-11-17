//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/FileServicesList.json
func ExampleFileServicesClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewFileServicesClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.List(ctx, "res9290", "sto1590", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/FileServicesPut.json
func ExampleFileServicesClient_SetServiceProperties_putFileServices() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewFileServicesClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.SetServiceProperties(ctx, "res4410", "sto8607", armstorage.FileServiceProperties{
		FileServiceProperties: &armstorage.FileServicePropertiesProperties{
			Cors: &armstorage.CorsRules{
				CorsRules: []*armstorage.CorsRule{
					{
						AllowedHeaders: []*string{
							to.Ptr("x-ms-meta-abc"),
							to.Ptr("x-ms-meta-data*"),
							to.Ptr("x-ms-meta-target*")},
						AllowedMethods: []*armstorage.CorsRuleAllowedMethodsItem{
							to.Ptr(armstorage.CorsRuleAllowedMethodsItemGET),
							to.Ptr(armstorage.CorsRuleAllowedMethodsItemHEAD),
							to.Ptr(armstorage.CorsRuleAllowedMethodsItemPOST),
							to.Ptr(armstorage.CorsRuleAllowedMethodsItemOPTIONS),
							to.Ptr(armstorage.CorsRuleAllowedMethodsItemMERGE),
							to.Ptr(armstorage.CorsRuleAllowedMethodsItemPUT)},
						AllowedOrigins: []*string{
							to.Ptr("http://www.contoso.com"),
							to.Ptr("http://www.fabrikam.com")},
						ExposedHeaders: []*string{
							to.Ptr("x-ms-meta-*")},
						MaxAgeInSeconds: to.Ptr[int32](100),
					},
					{
						AllowedHeaders: []*string{
							to.Ptr("*")},
						AllowedMethods: []*armstorage.CorsRuleAllowedMethodsItem{
							to.Ptr(armstorage.CorsRuleAllowedMethodsItemGET)},
						AllowedOrigins: []*string{
							to.Ptr("*")},
						ExposedHeaders: []*string{
							to.Ptr("*")},
						MaxAgeInSeconds: to.Ptr[int32](2),
					},
					{
						AllowedHeaders: []*string{
							to.Ptr("x-ms-meta-12345675754564*")},
						AllowedMethods: []*armstorage.CorsRuleAllowedMethodsItem{
							to.Ptr(armstorage.CorsRuleAllowedMethodsItemGET),
							to.Ptr(armstorage.CorsRuleAllowedMethodsItemPUT)},
						AllowedOrigins: []*string{
							to.Ptr("http://www.abc23.com"),
							to.Ptr("https://www.fabrikam.com/*")},
						ExposedHeaders: []*string{
							to.Ptr("x-ms-meta-abc"),
							to.Ptr("x-ms-meta-data*"),
							to.Ptr("x-ms-meta-target*")},
						MaxAgeInSeconds: to.Ptr[int32](2000),
					}},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/FileServicesPut_EnableSMBMultichannel.json
func ExampleFileServicesClient_SetServiceProperties_putFileServicesEnableSmbMultichannel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewFileServicesClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.SetServiceProperties(ctx, "res4410", "sto8607", armstorage.FileServiceProperties{
		FileServiceProperties: &armstorage.FileServicePropertiesProperties{
			ProtocolSettings: &armstorage.ProtocolSettings{
				Smb: &armstorage.SmbSetting{
					Multichannel: &armstorage.Multichannel{
						Enabled: to.Ptr(true),
					},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/FileServicesPut_EnableSecureSmbFeatures.json
func ExampleFileServicesClient_SetServiceProperties_putFileServicesEnableSecureSmbFeatures() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewFileServicesClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.SetServiceProperties(ctx, "res4410", "sto8607", armstorage.FileServiceProperties{
		FileServiceProperties: &armstorage.FileServicePropertiesProperties{
			ProtocolSettings: &armstorage.ProtocolSettings{
				Smb: &armstorage.SmbSetting{
					AuthenticationMethods:    to.Ptr("NTLMv2;Kerberos"),
					ChannelEncryption:        to.Ptr("AES-128-CCM;AES-128-GCM;AES-256-GCM"),
					KerberosTicketEncryption: to.Ptr("RC4-HMAC;AES-256"),
					Versions:                 to.Ptr("SMB2.1;SMB3.0;SMB3.1.1"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/FileServicesGet.json
func ExampleFileServicesClient_GetServiceProperties() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewFileServicesClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetServiceProperties(ctx, "res4410", "sto8607", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
