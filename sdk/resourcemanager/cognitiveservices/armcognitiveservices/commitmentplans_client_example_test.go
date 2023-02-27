//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-12-01/examples/ListCommitmentPlans.json
func ExampleCommitmentPlansClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcognitiveservices.NewCommitmentPlansClient("subscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("resourceGroupName", "accountName", nil)
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
		// page.CommitmentPlanListResult = armcognitiveservices.CommitmentPlanListResult{
		// 	Value: []*armcognitiveservices.CommitmentPlan{
		// 		{
		// 			Name: to.Ptr("commitmentPlanName"),
		// 			Type: to.Ptr("Microsoft.CognitiveServices/accounts/commitmentPlans"),
		// 			ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/commitmentPlans/commitmentPlanName"),
		// 			Properties: &armcognitiveservices.CommitmentPlanProperties{
		// 				AutoRenew: to.Ptr(true),
		// 				Current: &armcognitiveservices.CommitmentPeriod{
		// 					Tier: to.Ptr("T1"),
		// 				},
		// 				HostingModel: to.Ptr(armcognitiveservices.HostingModelWeb),
		// 				PlanType: to.Ptr("Speech2Text"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-12-01/examples/GetCommitmentPlan.json
func ExampleCommitmentPlansClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcognitiveservices.NewCommitmentPlansClient("subscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "resourceGroupName", "accountName", "commitmentPlanName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CommitmentPlan = armcognitiveservices.CommitmentPlan{
	// 	Name: to.Ptr("commitmentPlanName"),
	// 	Type: to.Ptr("Microsoft.CognitiveServices/accounts/commitmentPlans"),
	// 	ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/commitmentPlans/commitmentPlanName"),
	// 	Properties: &armcognitiveservices.CommitmentPlanProperties{
	// 		AutoRenew: to.Ptr(true),
	// 		Current: &armcognitiveservices.CommitmentPeriod{
	// 			Tier: to.Ptr("T1"),
	// 		},
	// 		HostingModel: to.Ptr(armcognitiveservices.HostingModelWeb),
	// 		PlanType: to.Ptr("Speech2Text"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-12-01/examples/PutCommitmentPlan.json
func ExampleCommitmentPlansClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcognitiveservices.NewCommitmentPlansClient("subscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "resourceGroupName", "accountName", "commitmentPlanName", armcognitiveservices.CommitmentPlan{
		Properties: &armcognitiveservices.CommitmentPlanProperties{
			AutoRenew: to.Ptr(true),
			Current: &armcognitiveservices.CommitmentPeriod{
				Tier: to.Ptr("T1"),
			},
			HostingModel: to.Ptr(armcognitiveservices.HostingModelWeb),
			PlanType:     to.Ptr("Speech2Text"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CommitmentPlan = armcognitiveservices.CommitmentPlan{
	// 	Name: to.Ptr("commitmentPlanName"),
	// 	Type: to.Ptr("Microsoft.CognitiveServices/accounts/commitmentPlans"),
	// 	ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/commitmentPlans/commitmentPlanName"),
	// 	Properties: &armcognitiveservices.CommitmentPlanProperties{
	// 		AutoRenew: to.Ptr(true),
	// 		Current: &armcognitiveservices.CommitmentPeriod{
	// 			Tier: to.Ptr("T1"),
	// 		},
	// 		HostingModel: to.Ptr(armcognitiveservices.HostingModelWeb),
	// 		PlanType: to.Ptr("Speech2Text"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-12-01/examples/DeleteCommitmentPlan.json
func ExampleCommitmentPlansClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcognitiveservices.NewCommitmentPlansClient("subscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx, "resourceGroupName", "accountName", "commitmentPlanName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-12-01/examples/CreateSharedCommitmentPlan.json
func ExampleCommitmentPlansClient_BeginCreateOrUpdatePlan() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcognitiveservices.NewCommitmentPlansClient("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdatePlan(ctx, "resourceGroupName", "commitmentPlanName", armcognitiveservices.CommitmentPlan{
		Kind:     to.Ptr("SpeechServices"),
		Location: to.Ptr("West US"),
		Properties: &armcognitiveservices.CommitmentPlanProperties{
			AutoRenew: to.Ptr(true),
			Current: &armcognitiveservices.CommitmentPeriod{
				Tier: to.Ptr("T1"),
			},
			HostingModel: to.Ptr(armcognitiveservices.HostingModelWeb),
			PlanType:     to.Ptr("STT"),
		},
		SKU: &armcognitiveservices.SKU{
			Name: to.Ptr("S0"),
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
	// res.CommitmentPlan = armcognitiveservices.CommitmentPlan{
	// 	Name: to.Ptr("commitmentPlanName"),
	// 	Type: to.Ptr("Microsoft.CognitiveServices/commitmentPlans"),
	// 	ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/commitmentPlans/commitmentPlanName"),
	// 	Kind: to.Ptr("SpeechServices"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armcognitiveservices.CommitmentPlanProperties{
	// 		AutoRenew: to.Ptr(true),
	// 		Current: &armcognitiveservices.CommitmentPeriod{
	// 			Tier: to.Ptr("T1"),
	// 		},
	// 		HostingModel: to.Ptr(armcognitiveservices.HostingModelWeb),
	// 		PlanType: to.Ptr("STT"),
	// 		ProvisioningState: to.Ptr(armcognitiveservices.CommitmentPlanProvisioningStateSucceeded),
	// 	},
	// 	SKU: &armcognitiveservices.SKU{
	// 		Name: to.Ptr("S0"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-12-01/examples/UpdateSharedCommitmentPlan.json
func ExampleCommitmentPlansClient_BeginUpdatePlan() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcognitiveservices.NewCommitmentPlansClient("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdatePlan(ctx, "resourceGroupName", "commitmentPlanName", armcognitiveservices.PatchResourceTagsAndSKU{
		Tags: map[string]*string{
			"name": to.Ptr("value"),
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
	// res.CommitmentPlan = armcognitiveservices.CommitmentPlan{
	// 	Name: to.Ptr("commitmentPlanName"),
	// 	Type: to.Ptr("Microsoft.CognitiveServices/commitmentPlans"),
	// 	ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/commitmentPlans/commitmentPlanName"),
	// 	Kind: to.Ptr("SpeechServices"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armcognitiveservices.CommitmentPlanProperties{
	// 		AutoRenew: to.Ptr(true),
	// 		Current: &armcognitiveservices.CommitmentPeriod{
	// 			Tier: to.Ptr("T1"),
	// 		},
	// 		HostingModel: to.Ptr(armcognitiveservices.HostingModelWeb),
	// 		PlanType: to.Ptr("STT"),
	// 		ProvisioningState: to.Ptr(armcognitiveservices.CommitmentPlanProvisioningStateSucceeded),
	// 	},
	// 	SKU: &armcognitiveservices.SKU{
	// 		Name: to.Ptr("S0"),
	// 	},
	// 	Tags: map[string]*string{
	// 		"name": to.Ptr("value"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-12-01/examples/DeleteSharedCommitmentPlan.json
func ExampleCommitmentPlansClient_BeginDeletePlan() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcognitiveservices.NewCommitmentPlansClient("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDeletePlan(ctx, "resourceGroupName", "commitmentPlanName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-12-01/examples/GetSharedCommitmentPlan.json
func ExampleCommitmentPlansClient_GetPlan() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcognitiveservices.NewCommitmentPlansClient("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetPlan(ctx, "resourceGroupName", "commitmentPlanName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CommitmentPlan = armcognitiveservices.CommitmentPlan{
	// 	Name: to.Ptr("commitmentPlanName"),
	// 	Type: to.Ptr("Microsoft.CognitiveServices/commitmentPlans"),
	// 	ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/commitmentPlans/commitmentPlanName"),
	// 	Kind: to.Ptr("SpeechServices"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armcognitiveservices.CommitmentPlanProperties{
	// 		AutoRenew: to.Ptr(true),
	// 		Current: &armcognitiveservices.CommitmentPeriod{
	// 			Tier: to.Ptr("T1"),
	// 		},
	// 		HostingModel: to.Ptr(armcognitiveservices.HostingModelWeb),
	// 		PlanType: to.Ptr("STT"),
	// 		ProvisioningState: to.Ptr(armcognitiveservices.CommitmentPlanProvisioningStateSucceeded),
	// 	},
	// 	SKU: &armcognitiveservices.SKU{
	// 		Name: to.Ptr("S0"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-12-01/examples/ListSharedCommitmentPlansByResourceGroup.json
func ExampleCommitmentPlansClient_NewListPlansByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcognitiveservices.NewCommitmentPlansClient("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPlansByResourceGroupPager("resourceGroupName", nil)
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
		// page.CommitmentPlanListResult = armcognitiveservices.CommitmentPlanListResult{
		// 	Value: []*armcognitiveservices.CommitmentPlan{
		// 		{
		// 			Name: to.Ptr("commitmentPlanName"),
		// 			Type: to.Ptr("Microsoft.CognitiveServices/commitmentPlans"),
		// 			ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/commitmentPlans/commitmentPlanName"),
		// 			Kind: to.Ptr("SpeechServices"),
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armcognitiveservices.CommitmentPlanProperties{
		// 				AutoRenew: to.Ptr(true),
		// 				Current: &armcognitiveservices.CommitmentPeriod{
		// 					Tier: to.Ptr("T1"),
		// 				},
		// 				HostingModel: to.Ptr(armcognitiveservices.HostingModelWeb),
		// 				PlanType: to.Ptr("STT"),
		// 				ProvisioningState: to.Ptr(armcognitiveservices.CommitmentPlanProvisioningStateSucceeded),
		// 			},
		// 			SKU: &armcognitiveservices.SKU{
		// 				Name: to.Ptr("S0"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-12-01/examples/ListSharedCommitmentPlansBySubscription.json
func ExampleCommitmentPlansClient_NewListPlansBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcognitiveservices.NewCommitmentPlansClient("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPlansBySubscriptionPager(nil)
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
		// page.CommitmentPlanListResult = armcognitiveservices.CommitmentPlanListResult{
		// 	Value: []*armcognitiveservices.CommitmentPlan{
		// 		{
		// 			Name: to.Ptr("commitmentPlanName"),
		// 			Type: to.Ptr("Microsoft.CognitiveServices/commitmentPlans"),
		// 			ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/commitmentPlans/commitmentPlanName"),
		// 			Kind: to.Ptr("SpeechServices"),
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armcognitiveservices.CommitmentPlanProperties{
		// 				AutoRenew: to.Ptr(true),
		// 				Current: &armcognitiveservices.CommitmentPeriod{
		// 					Tier: to.Ptr("T1"),
		// 				},
		// 				HostingModel: to.Ptr(armcognitiveservices.HostingModelWeb),
		// 				PlanType: to.Ptr("STT"),
		// 				ProvisioningState: to.Ptr(armcognitiveservices.CommitmentPlanProvisioningStateSucceeded),
		// 			},
		// 			SKU: &armcognitiveservices.SKU{
		// 				Name: to.Ptr("S0"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-12-01/examples/ListSharedCommitmentPlanAssociations.json
func ExampleCommitmentPlansClient_NewListAssociationsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcognitiveservices.NewCommitmentPlansClient("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListAssociationsPager("resourceGroupName", "commitmentPlanName", nil)
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
		// page.CommitmentPlanAccountAssociationListResult = armcognitiveservices.CommitmentPlanAccountAssociationListResult{
		// 	Value: []*armcognitiveservices.CommitmentPlanAccountAssociation{
		// 		{
		// 			Name: to.Ptr("accountAssociationName"),
		// 			Type: to.Ptr("Microsoft.CognitiveServices/commitmentPlans/accountAssociations"),
		// 			ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/commitmentPlans/commitmentPlanName/accountAssociations/accountAssociationName"),
		// 			Properties: &armcognitiveservices.CommitmentPlanAccountAssociationProperties{
		// 				AccountID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-12-01/examples/GetSharedCommitmentPlanAssociation.json
func ExampleCommitmentPlansClient_GetAssociation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcognitiveservices.NewCommitmentPlansClient("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetAssociation(ctx, "resourceGroupName", "commitmentPlanName", "commitmentPlanAssociationName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CommitmentPlanAccountAssociation = armcognitiveservices.CommitmentPlanAccountAssociation{
	// 	Name: to.Ptr("commitmentPlanAssociationName"),
	// 	Type: to.Ptr("Microsoft.CognitiveServices/commitmentPlans/accountAssociations"),
	// 	ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/commitmentPlans/commitmentPlanName/accountAssociations/commitmentPlanAssociationName"),
	// 	Properties: &armcognitiveservices.CommitmentPlanAccountAssociationProperties{
	// 		AccountID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-12-01/examples/CreateSharedCommitmentPlanAssociation.json
func ExampleCommitmentPlansClient_BeginCreateOrUpdateAssociation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcognitiveservices.NewCommitmentPlansClient("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdateAssociation(ctx, "resourceGroupName", "commitmentPlanName", "commitmentPlanAssociationName", armcognitiveservices.CommitmentPlanAccountAssociation{
		Properties: &armcognitiveservices.CommitmentPlanAccountAssociationProperties{
			AccountID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName"),
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
	// res.CommitmentPlanAccountAssociation = armcognitiveservices.CommitmentPlanAccountAssociation{
	// 	Name: to.Ptr("commitmentPlanAssociationName"),
	// 	Type: to.Ptr("Microsoft.CognitiveServices/commitmentPlans/accountAssociations"),
	// 	ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/commitmentPlans/commitmentPlanName/accountAssociations/commitmentPlanAssociationName"),
	// 	Properties: &armcognitiveservices.CommitmentPlanAccountAssociationProperties{
	// 		AccountID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-12-01/examples/DeleteSharedCommitmentPlanAssociation.json
func ExampleCommitmentPlansClient_BeginDeleteAssociation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcognitiveservices.NewCommitmentPlansClient("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDeleteAssociation(ctx, "resourceGroupName", "commitmentPlanName", "commitmentPlanAssociationName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
