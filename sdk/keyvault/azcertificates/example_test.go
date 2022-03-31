//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azcertificates_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azcertificates"
)

func ExampleNewClient() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		panic(err)
	}

	vaultURL, ok := os.LookupEnv("AZURE_KEYVAULT_URL")
	if !ok {
		log.Fatalf("Could not find 'AZURE_KEYVAULT_URL' in environment variables")
	}

	client, err := azcertificates.NewClient(vaultURL, cred, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(client)
}

func ExampleClient_BeginCreateCertificate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		panic(err)
	}

	vaultURL, ok := os.LookupEnv("AZURE_KEYVAULT_URL")
	if !ok {
		log.Fatalf("Could not find 'AZURE_KEYVAULT_URL' in environment variables")
	}

	client, err := azcertificates.NewClient(vaultURL, cred, nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.BeginCreateCertificate(context.TODO(), "certificateName", azcertificates.Policy{
		IssuerParameters: &azcertificates.IssuerParameters{
			Name: to.Ptr("Self"),
		},
		X509CertificateProperties: &azcertificates.X509CertificateProperties{
			Subject: to.Ptr("CN=DefaultPolicy"),
		},
	}, nil)
	if err != nil {
		panic(err)
	}

	finalResponse, err := resp.PollUntilDone(context.TODO(), time.Second)
	if err != nil {
		panic(err)
	}

	fmt.Println("Created a certificate with ID: ", *finalResponse.ID)
}

func ExampleClient_GetCertificate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		panic(err)
	}

	vaultURL, ok := os.LookupEnv("AZURE_KEYVAULT_URL")
	if !ok {
		log.Fatalf("Could not find 'AZURE_KEYVAULT_URL' in environment variables")
	}

	client, err := azcertificates.NewClient(vaultURL, cred, nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.GetCertificate(context.TODO(), "myCertName", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(*resp.ID)

	// optionally you can get a specific version
	resp, err = client.GetCertificate(context.TODO(), "myCertName", &azcertificates.GetCertificateOptions{Version: "myCertVersion"})
	if err != nil {
		panic(err)
	}
}
func ExampleClient_UpdateCertificateProperties() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		panic(err)
	}

	vaultURL, ok := os.LookupEnv("AZURE_KEYVAULT_URL")
	if !ok {
		log.Fatalf("Could not find 'AZURE_KEYVAULT_URL' in environment variables")
	}

	client, err := azcertificates.NewClient(vaultURL, cred, nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.UpdateCertificateProperties(context.TODO(), "myCertName", &azcertificates.UpdateCertificatePropertiesOptions{
		Version: "myNewVersion",
		Properties: &azcertificates.Properties{
			Enabled: to.Ptr(false),
			Expires: to.Ptr(time.Now().Add(72 * time.Hour)),
			Tags:    map[string]string{"Tag1": "Val1"},
		},
		CertificatePolicy: &azcertificates.Policy{
			IssuerParameters: &azcertificates.IssuerParameters{
				Name: to.Ptr("Self"),
			},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(*resp.ID)
	fmt.Println(*resp.Certificate.Properties.Enabled)
	fmt.Println(resp.Properties.Tags)
}

func ExampleClient_ListCertificates() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		panic(err)
	}

	vaultURL, ok := os.LookupEnv("AZURE_KEYVAULT_URL")
	if !ok {
		log.Fatalf("Could not find 'AZURE_KEYVAULT_URL' in environment variables")
	}

	client, err := azcertificates.NewClient(vaultURL, cred, nil)
	if err != nil {
		panic(err)
	}

	pager := client.ListCertificates(nil)
	for pager.More() {
		page, err := pager.NextPage(context.TODO())
		if err != nil {
			panic(err)
		}
		for _, cert := range page.Certificates {
			fmt.Println(*cert.ID)
		}
	}
}

func ExampleClient_BeginDeleteCertificate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		panic(err)
	}

	vaultURL, ok := os.LookupEnv("AZURE_KEYVAULT_URL")
	if !ok {
		log.Fatalf("Could not find 'AZURE_KEYVAULT_URL' in environment variables")
	}

	client, err := azcertificates.NewClient(vaultURL, cred, nil)
	if err != nil {
		panic(err)
	}

	pollerResp, err := client.BeginDeleteCertificate(context.TODO(), "certToDelete", nil)
	if err != nil {
		panic(err)
	}
	finalResp, err := pollerResp.PollUntilDone(context.TODO(), time.Second)
	if err != nil {
		panic(err)
	}

	fmt.Println("Deleted certificate with ID: ", *finalResp.ID)
}
