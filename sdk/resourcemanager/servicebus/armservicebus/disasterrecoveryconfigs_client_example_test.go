//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5e24fa8e30bff0cb321494a0b550b1c1282a8a3c/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/disasterRecoveryConfigs/SBAliasList.json
func ExampleDisasterRecoveryConfigsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armservicebus.NewDisasterRecoveryConfigsClient("5f750a97-50d9-4e36-8081-c9ee4c0210d4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("ardsouzatestRG", "sdk-Namespace-8860", nil)
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
		// page.ArmDisasterRecoveryListResult = armservicebus.ArmDisasterRecoveryListResult{
		// 	Value: []*armservicebus.ArmDisasterRecovery{
		// 		{
		// 			Name: to.Ptr("sdk-DisasterRecovery-3814"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces/DisasterRecoveryConfig"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/ardsouzatestRG/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-8860/disasterRecoveryConfig/sdk-DisasterRecovery-3814"),
		// 			Properties: &armservicebus.ArmDisasterRecoveryProperties{
		// 				PartnerNamespace: to.Ptr("sdk-Namespace-37"),
		// 				ProvisioningState: to.Ptr(armservicebus.ProvisioningStateDRSucceeded),
		// 				Role: to.Ptr(armservicebus.RoleDisasterRecoveryPrimary),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5e24fa8e30bff0cb321494a0b550b1c1282a8a3c/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/disasterRecoveryConfigs/SBAliasCreate.json
func ExampleDisasterRecoveryConfigsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armservicebus.NewDisasterRecoveryConfigsClient("5f750a97-50d9-4e36-8081-c9ee4c0210d4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "ardsouzatestRG", "sdk-Namespace-8860", "sdk-Namespace-8860", armservicebus.ArmDisasterRecovery{
		Properties: &armservicebus.ArmDisasterRecoveryProperties{
			AlternateName:    to.Ptr("alternameforAlias-Namespace-8860"),
			PartnerNamespace: to.Ptr("sdk-Namespace-37"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ArmDisasterRecovery = armservicebus.ArmDisasterRecovery{
	// 	Name: to.Ptr("sdk-Namespace-8860"),
	// 	Type: to.Ptr("Microsoft.ServiceBus/Namespaces/DisasterRecoveryConfig"),
	// 	ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/ardsouzatestRG/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-8860/disasterRecoveryConfig/sdk-Namespace-8860"),
	// 	Properties: &armservicebus.ArmDisasterRecoveryProperties{
	// 		AlternateName: to.Ptr("alternameforAlias-Namespace-8860"),
	// 		PartnerNamespace: to.Ptr("sdk-Namespace-37"),
	// 		ProvisioningState: to.Ptr(armservicebus.ProvisioningStateDRSucceeded),
	// 		Role: to.Ptr(armservicebus.RoleDisasterRecoveryPrimary),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5e24fa8e30bff0cb321494a0b550b1c1282a8a3c/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/disasterRecoveryConfigs/SBAliasDelete.json
func ExampleDisasterRecoveryConfigsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armservicebus.NewDisasterRecoveryConfigsClient("5f750a97-50d9-4e36-8081-c9ee4c0210d4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx, "SouthCentralUS", "sdk-Namespace-8860", "sdk-DisasterRecovery-3814", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5e24fa8e30bff0cb321494a0b550b1c1282a8a3c/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/disasterRecoveryConfigs/SBAliasGet.json
func ExampleDisasterRecoveryConfigsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armservicebus.NewDisasterRecoveryConfigsClient("5f750a97-50d9-4e36-8081-c9ee4c0210d4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "ardsouzatestRG", "sdk-Namespace-8860", "sdk-DisasterRecovery-3814", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ArmDisasterRecovery = armservicebus.ArmDisasterRecovery{
	// 	Name: to.Ptr("sdk-DisasterRecovery-3814"),
	// 	Type: to.Ptr("Microsoft.ServiceBus/Namespaces/DisasterRecoveryConfig"),
	// 	ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/ardsouzatestRG/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-37/disasterRecoveryConfig/sdk-DisasterRecovery-3814"),
	// 	Properties: &armservicebus.ArmDisasterRecoveryProperties{
	// 		PartnerNamespace: to.Ptr("sdk-Namespace-8860"),
	// 		PendingReplicationOperationsCount: to.Ptr[int64](0),
	// 		ProvisioningState: to.Ptr(armservicebus.ProvisioningStateDRSucceeded),
	// 		Role: to.Ptr(armservicebus.RoleDisasterRecoverySecondary),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5e24fa8e30bff0cb321494a0b550b1c1282a8a3c/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/disasterRecoveryConfigs/SBEHAliasBreakPairing.json
func ExampleDisasterRecoveryConfigsClient_BreakPairing() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armservicebus.NewDisasterRecoveryConfigsClient("5f750a97-50d9-4e36-8081-c9ee4c0210d4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.BreakPairing(ctx, "ardsouzatestRG", "sdk-Namespace-8860", "sdk-DisasterRecovery-3814", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5e24fa8e30bff0cb321494a0b550b1c1282a8a3c/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/disasterRecoveryConfigs/SBAliasFailOver.json
func ExampleDisasterRecoveryConfigsClient_FailOver() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armservicebus.NewDisasterRecoveryConfigsClient("5f750a97-50d9-4e36-8081-c9ee4c0210d4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.FailOver(ctx, "ardsouzatestRG", "sdk-Namespace-8860", "sdk-DisasterRecovery-3814", &armservicebus.DisasterRecoveryConfigsClientFailOverOptions{Parameters: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5e24fa8e30bff0cb321494a0b550b1c1282a8a3c/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/disasterRecoveryConfigs/SBAliasAuthorizationRuleListAll.json
func ExampleDisasterRecoveryConfigsClient_NewListAuthorizationRulesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armservicebus.NewDisasterRecoveryConfigsClient("exampleSubscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListAuthorizationRulesPager("exampleResourceGroup", "sdk-Namespace-9080", "sdk-DisasterRecovery-4047", nil)
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
		// page.SBAuthorizationRuleListResult = armservicebus.SBAuthorizationRuleListResult{
		// 	Value: []*armservicebus.SBAuthorizationRule{
		// 		{
		// 			Name: to.Ptr("RootManageSharedAccessKey"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces/DisasterRecoveryConfig/AuthorizationRules"),
		// 			ID: to.Ptr("/subscriptions/exampleSubscriptionId/resourceGroups/exampleResourceGroup/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-9080/disasterRecoveryConfigs/sdk-DisasterRecovery-4047/AuthorizationRules/RootManageSharedAccessKey"),
		// 			Properties: &armservicebus.SBAuthorizationRuleProperties{
		// 				Rights: []*armservicebus.AccessRights{
		// 					to.Ptr(armservicebus.AccessRightsListen),
		// 					to.Ptr(armservicebus.AccessRightsManage),
		// 					to.Ptr(armservicebus.AccessRightsSend)},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("sdk-Authrules-1067"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces/DisasterRecoveryConfig/AuthorizationRules"),
		// 				ID: to.Ptr("/subscriptions/exampleSubscriptionId/resourceGroups/exampleResourceGroup/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-9080/disasterRecoveryConfigs/sdk-DisasterRecovery-4047/AuthorizationRules/sdk-Authrules-1067"),
		// 				Properties: &armservicebus.SBAuthorizationRuleProperties{
		// 					Rights: []*armservicebus.AccessRights{
		// 						to.Ptr(armservicebus.AccessRightsListen),
		// 						to.Ptr(armservicebus.AccessRightsSend)},
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("sdk-Authrules-1684"),
		// 					Type: to.Ptr("Microsoft.ServiceBus/Namespaces/DisasterRecoveryConfig/AuthorizationRules"),
		// 					ID: to.Ptr("/subscriptions/exampleSubscriptionId/resourceGroups/exampleResourceGroup/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-9080/disasterRecoveryConfigs/sdk-DisasterRecovery-4047/AuthorizationRules/sdk-Authrules-1684"),
		// 					Properties: &armservicebus.SBAuthorizationRuleProperties{
		// 						Rights: []*armservicebus.AccessRights{
		// 							to.Ptr(armservicebus.AccessRightsListen),
		// 							to.Ptr(armservicebus.AccessRightsSend)},
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("sdk-Authrules-4879"),
		// 						Type: to.Ptr("Microsoft.ServiceBus/Namespaces/DisasterRecoveryConfig/AuthorizationRules"),
		// 						ID: to.Ptr("/subscriptions/exampleSubscriptionId/resourceGroups/exampleResourceGroup/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-9080/disasterRecoveryConfigs/sdk-DisasterRecovery-4047/AuthorizationRules/sdk-Authrules-4879"),
		// 						Properties: &armservicebus.SBAuthorizationRuleProperties{
		// 							Rights: []*armservicebus.AccessRights{
		// 								to.Ptr(armservicebus.AccessRightsListen),
		// 								to.Ptr(armservicebus.AccessRightsSend)},
		// 							},
		// 					}},
		// 				}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5e24fa8e30bff0cb321494a0b550b1c1282a8a3c/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/disasterRecoveryConfigs/SBAliasAuthorizationRuleGet.json
func ExampleDisasterRecoveryConfigsClient_GetAuthorizationRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armservicebus.NewDisasterRecoveryConfigsClient("exampleSubscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetAuthorizationRule(ctx, "exampleResourceGroup", "sdk-Namespace-9080", "sdk-DisasterRecovery-4879", "sdk-Authrules-4879", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SBAuthorizationRule = armservicebus.SBAuthorizationRule{
	// 	Name: to.Ptr("sdk-Authrules-4879"),
	// 	Type: to.Ptr("Microsoft.ServiceBus/Namespaces/DisasterRecoveryConfig/AuthorizationRules"),
	// 	ID: to.Ptr("/subscriptions/exampleSubscriptionId/resourceGroups/exampleResourceGroup/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-9080/disasterRecoveryConfigs/sdk-DisasterRecovery-4047/AuthorizationRules/sdk-Authrules-4879"),
	// 	Properties: &armservicebus.SBAuthorizationRuleProperties{
	// 		Rights: []*armservicebus.AccessRights{
	// 			to.Ptr(armservicebus.AccessRightsListen),
	// 			to.Ptr(armservicebus.AccessRightsSend)},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5e24fa8e30bff0cb321494a0b550b1c1282a8a3c/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/disasterRecoveryConfigs/SBAliasAuthorizationRuleListKey.json
func ExampleDisasterRecoveryConfigsClient_ListKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armservicebus.NewDisasterRecoveryConfigsClient("exampleSubscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ListKeys(ctx, "exampleResourceGroup", "sdk-Namespace-2702", "sdk-DisasterRecovery-4047", "sdk-Authrules-1746", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccessKeys = armservicebus.AccessKeys{
	// 	AliasPrimaryConnectionString: to.Ptr("Endpoint=sb://sdk-disasterrecovery-4047.servicebus.windows-int.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=############################################"),
	// 	AliasSecondaryConnectionString: to.Ptr("Endpoint=sb://sdk-disasterrecovery-4047.servicebus.windows-int.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=############################################"),
	// 	KeyName: to.Ptr("sdk-Authrules-1746"),
	// 	PrimaryKey: to.Ptr("############################################"),
	// 	SecondaryKey: to.Ptr("############################################"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5e24fa8e30bff0cb321494a0b550b1c1282a8a3c/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/disasterRecoveryConfigs/SBAliasCheckNameAvailability.json
func ExampleDisasterRecoveryConfigsClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armservicebus.NewDisasterRecoveryConfigsClient("exampleSubscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CheckNameAvailability(ctx, "exampleResourceGroup", "sdk-Namespace-9080", armservicebus.CheckNameAvailability{
		Name: to.Ptr("sdk-DisasterRecovery-9474"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckNameAvailabilityResult = armservicebus.CheckNameAvailabilityResult{
	// 	Message: to.Ptr(""),
	// 	NameAvailable: to.Ptr(true),
	// 	Reason: to.Ptr(armservicebus.UnavailableReasonNone),
	// }
}
