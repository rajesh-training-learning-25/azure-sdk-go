//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armfrontdoor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/frontdoor/resource-manager/Microsoft.Network/stable/2024-02-01/examples/WafListManagedRuleSets.json
func ExampleManagedRuleSetsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfrontdoor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedRuleSetsClient().NewListPager(nil)
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
		// page.ManagedRuleSetDefinitionList = armfrontdoor.ManagedRuleSetDefinitionList{
		// 	Value: []*armfrontdoor.ManagedRuleSetDefinition{
		// 		{
		// 			Name: to.Ptr("DefaultRuleSet_1.0"),
		// 			Type: to.Ptr("Microsoft.Network/frontdoorwebapplicationfirewallmanagedrulesets"),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Network/FrontDoorWebApplicationFirewallManagedRuleSets"),
		// 			Properties: &armfrontdoor.ManagedRuleSetDefinitionProperties{
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				RuleGroups: []*armfrontdoor.ManagedRuleGroupDefinition{
		// 					{
		// 						Description: to.Ptr("SQL injection"),
		// 						RuleGroupName: to.Ptr("SQLI"),
		// 						Rules: []*armfrontdoor.ManagedRuleDefinition{
		// 							{
		// 								Description: to.Ptr("SQL Injection Attack Detected via libinjection"),
		// 								DefaultAction: to.Ptr(armfrontdoor.ActionTypeBlock),
		// 								DefaultState: to.Ptr(armfrontdoor.ManagedRuleEnabledStateEnabled),
		// 								RuleID: to.Ptr("942100"),
		// 							},
		// 							{
		// 								Description: to.Ptr("SQL Injection Attack: Common Injection Testing Detected"),
		// 								DefaultAction: to.Ptr(armfrontdoor.ActionTypeBlock),
		// 								DefaultState: to.Ptr(armfrontdoor.ManagedRuleEnabledStateEnabled),
		// 								RuleID: to.Ptr("942110"),
		// 						}},
		// 					},
		// 					{
		// 						Description: to.Ptr("Cross-site scripting"),
		// 						RuleGroupName: to.Ptr("XSS"),
		// 						Rules: []*armfrontdoor.ManagedRuleDefinition{
		// 							{
		// 								Description: to.Ptr("XSS Attack Detected via libinjection"),
		// 								DefaultAction: to.Ptr(armfrontdoor.ActionTypeBlock),
		// 								DefaultState: to.Ptr(armfrontdoor.ManagedRuleEnabledStateEnabled),
		// 								RuleID: to.Ptr("941100"),
		// 							},
		// 							{
		// 								Description: to.Ptr("XSS Attack Detected via libinjection"),
		// 								DefaultAction: to.Ptr(armfrontdoor.ActionTypeBlock),
		// 								DefaultState: to.Ptr(armfrontdoor.ManagedRuleEnabledStateEnabled),
		// 								RuleID: to.Ptr("941101"),
		// 							},
		// 							{
		// 								Description: to.Ptr("XSS Filter - Category 1: Script Tag Vector"),
		// 								DefaultAction: to.Ptr(armfrontdoor.ActionTypeBlock),
		// 								DefaultState: to.Ptr(armfrontdoor.ManagedRuleEnabledStateEnabled),
		// 								RuleID: to.Ptr("941110"),
		// 						}},
		// 				}},
		// 				RuleSetID: to.Ptr("8125d145-ddc5-4d90-9bc3-24c5f2de69a2"),
		// 				RuleSetType: to.Ptr("DefaultRuleSet"),
		// 				RuleSetVersion: to.Ptr("1.0"),
		// 			},
		// 	}},
		// }
	}
}
