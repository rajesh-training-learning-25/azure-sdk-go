//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/createOrUpdatePolicyDefinition.json
func ExampleDefinitionsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpolicy.NewDefinitionsClient("<subscription-id>", cred, nil)
	_, err = client.CreateOrUpdate(ctx,
		"<policy-definition-name>",
		armpolicy.Definition{
			Properties: &armpolicy.DefinitionProperties{
				Description: to.StringPtr("<description>"),
				DisplayName: to.StringPtr("<display-name>"),
				Metadata: map[string]interface{}{
					"category": "Naming",
				},
				Mode: to.StringPtr("<mode>"),
				Parameters: map[string]*armpolicy.ParameterDefinitionsValue{
					"prefix": {
						Type: armpolicy.ParameterType("String").ToPtr(),
						Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
							Description: to.StringPtr("<description>"),
							DisplayName: to.StringPtr("<display-name>"),
						},
					},
					"suffix": {
						Type: armpolicy.ParameterType("String").ToPtr(),
						Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
							Description: to.StringPtr("<description>"),
							DisplayName: to.StringPtr("<display-name>"),
						},
					},
				},
				PolicyRule: map[string]interface{}{
					"if": map[string]interface{}{
						"not": map[string]interface{}{
							"field": "name",
							"like":  "[concat(parameters('prefix'), '*', parameters('suffix'))]",
						},
					},
					"then": map[string]interface{}{
						"effect": "deny",
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/deletePolicyDefinition.json
func ExampleDefinitionsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpolicy.NewDefinitionsClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<policy-definition-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/getPolicyDefinition.json
func ExampleDefinitionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpolicy.NewDefinitionsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<policy-definition-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DefinitionsClientGetResult)
}

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/getBuiltinPolicyDefinition.json
func ExampleDefinitionsClient_GetBuiltIn() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpolicy.NewDefinitionsClient("<subscription-id>", cred, nil)
	res, err := client.GetBuiltIn(ctx,
		"<policy-definition-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DefinitionsClientGetBuiltInResult)
}

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/createOrUpdatePolicyDefinitionAtManagementGroup.json
func ExampleDefinitionsClient_CreateOrUpdateAtManagementGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpolicy.NewDefinitionsClient("<subscription-id>", cred, nil)
	_, err = client.CreateOrUpdateAtManagementGroup(ctx,
		"<policy-definition-name>",
		"<management-group-id>",
		armpolicy.Definition{
			Properties: &armpolicy.DefinitionProperties{
				Description: to.StringPtr("<description>"),
				DisplayName: to.StringPtr("<display-name>"),
				Metadata: map[string]interface{}{
					"category": "Naming",
				},
				Mode: to.StringPtr("<mode>"),
				Parameters: map[string]*armpolicy.ParameterDefinitionsValue{
					"prefix": {
						Type: armpolicy.ParameterType("String").ToPtr(),
						Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
							Description: to.StringPtr("<description>"),
							DisplayName: to.StringPtr("<display-name>"),
						},
					},
					"suffix": {
						Type: armpolicy.ParameterType("String").ToPtr(),
						Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
							Description: to.StringPtr("<description>"),
							DisplayName: to.StringPtr("<display-name>"),
						},
					},
				},
				PolicyRule: map[string]interface{}{
					"if": map[string]interface{}{
						"not": map[string]interface{}{
							"field": "name",
							"like":  "[concat(parameters('prefix'), '*', parameters('suffix'))]",
						},
					},
					"then": map[string]interface{}{
						"effect": "deny",
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/deletePolicyDefinitionAtManagementGroup.json
func ExampleDefinitionsClient_DeleteAtManagementGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpolicy.NewDefinitionsClient("<subscription-id>", cred, nil)
	_, err = client.DeleteAtManagementGroup(ctx,
		"<policy-definition-name>",
		"<management-group-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/getPolicyDefinitionAtManagementGroup.json
func ExampleDefinitionsClient_GetAtManagementGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpolicy.NewDefinitionsClient("<subscription-id>", cred, nil)
	res, err := client.GetAtManagementGroup(ctx,
		"<policy-definition-name>",
		"<management-group-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DefinitionsClientGetAtManagementGroupResult)
}

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/listPolicyDefinitions.json
func ExampleDefinitionsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpolicy.NewDefinitionsClient("<subscription-id>", cred, nil)
	pager := client.List(&armpolicy.DefinitionsClientListOptions{Filter: nil,
		Top: nil,
	})
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

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/listBuiltInPolicyDefinitions.json
func ExampleDefinitionsClient_ListBuiltIn() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpolicy.NewDefinitionsClient("<subscription-id>", cred, nil)
	pager := client.ListBuiltIn(&armpolicy.DefinitionsClientListBuiltInOptions{Filter: nil,
		Top: nil,
	})
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

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/listPolicyDefinitionsByManagementGroup.json
func ExampleDefinitionsClient_ListByManagementGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpolicy.NewDefinitionsClient("<subscription-id>", cred, nil)
	pager := client.ListByManagementGroup("<management-group-id>",
		&armpolicy.DefinitionsClientListByManagementGroupOptions{Filter: nil,
			Top: nil,
		})
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
