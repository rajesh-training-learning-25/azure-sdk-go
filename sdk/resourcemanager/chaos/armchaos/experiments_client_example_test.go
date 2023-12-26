//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/chaos/resource-manager/Microsoft.Chaos/stable/2023-11-01/examples/ListExperimentsInASubscription.json
func ExampleExperimentsClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExperimentsClient().NewListAllPager(&armchaos.ExperimentsClientListAllOptions{Running: nil,
		ContinuationToken: nil,
	})
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
		// page.ExperimentListResult = armchaos.ExperimentListResult{
		// 	Value: []*armchaos.Experiment{
		// 		{
		// 			Name: to.Ptr("exampleExperiment"),
		// 			Type: to.Ptr("Microsoft.Chaos/experiments"),
		// 			ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/experiments/exampleExperiment"),
		// 			Location: to.Ptr("centraluseuap"),
		// 			Identity: &armchaos.ResourceIdentity{
		// 				Type: to.Ptr(armchaos.ResourceIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("d04ab567-2c07-43ef-a7f4-4527626b7f56"),
		// 				TenantID: to.Ptr("8c3e2fb2-fe7a-4bf1-b779-d73990782fe6"),
		// 			},
		// 			Properties: &armchaos.ExperimentProperties{
		// 				Selectors: []armchaos.TargetSelectorClassification{
		// 					&armchaos.TargetListSelector{
		// 						Type: to.Ptr(armchaos.SelectorTypeList),
		// 						ID: to.Ptr("selector1"),
		// 						Targets: []*armchaos.TargetReference{
		// 							{
		// 								Type: to.Ptr(armchaos.TargetReferenceTypeChaosTarget),
		// 								ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/exampleVM/providers/Microsoft.Chaos/targets/Microsoft-VirtualMachine"),
		// 						}},
		// 				}},
		// 				Steps: []*armchaos.ExperimentStep{
		// 					{
		// 						Name: to.Ptr("step1"),
		// 						Branches: []*armchaos.ExperimentBranch{
		// 							{
		// 								Name: to.Ptr("branch1"),
		// 								Actions: []armchaos.ExperimentActionClassification{
		// 									&armchaos.ContinuousAction{
		// 										Name: to.Ptr("urn:csci:provider:providername:Shutdown/1.0"),
		// 										Type: to.Ptr("continuous"),
		// 										Duration: to.Ptr("PT10M"),
		// 										Parameters: []*armchaos.KeyValuePair{
		// 											{
		// 												Key: to.Ptr("abruptShutdown"),
		// 												Value: to.Ptr("false"),
		// 										}},
		// 										SelectorID: to.Ptr("selector1"),
		// 								}},
		// 						}},
		// 				}},
		// 			},
		// 			SystemData: &armchaos.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-01T00:00:00.000Z"); return t}()),
		// 				CreatedBy: to.Ptr("User"),
		// 				CreatedByType: to.Ptr(armchaos.CreatedByType("b3a41dba-4415-4d36-9ee8-e5eaa86db976")),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-01T00:00:00.000Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("User"),
		// 				LastModifiedByType: to.Ptr(armchaos.CreatedByType("b3a41dba-4415-4d36-9ee8-e5eaa86db976")),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/chaos/resource-manager/Microsoft.Chaos/stable/2023-11-01/examples/ListExperimentsInAResourceGroup.json
func ExampleExperimentsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExperimentsClient().NewListPager("exampleRG", &armchaos.ExperimentsClientListOptions{Running: nil,
		ContinuationToken: nil,
	})
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
		// page.ExperimentListResult = armchaos.ExperimentListResult{
		// 	Value: []*armchaos.Experiment{
		// 		{
		// 			Name: to.Ptr("exampleExperiment"),
		// 			Type: to.Ptr("Microsoft.Chaos/experiments"),
		// 			ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/experiments/exampleExperiment"),
		// 			Location: to.Ptr("centraluseuap"),
		// 			Identity: &armchaos.ResourceIdentity{
		// 				Type: to.Ptr(armchaos.ResourceIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("d04ab567-2c07-43ef-a7f4-4527626b7f56"),
		// 				TenantID: to.Ptr("8c3e2fb2-fe7a-4bf1-b779-d73990782fe6"),
		// 			},
		// 			Properties: &armchaos.ExperimentProperties{
		// 				Selectors: []armchaos.TargetSelectorClassification{
		// 					&armchaos.TargetListSelector{
		// 						Type: to.Ptr(armchaos.SelectorTypeList),
		// 						ID: to.Ptr("selector1"),
		// 						Targets: []*armchaos.TargetReference{
		// 							{
		// 								Type: to.Ptr(armchaos.TargetReferenceTypeChaosTarget),
		// 								ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/exampleVM/providers/Microsoft.Chaos/targets/Microsoft-VirtualMachine"),
		// 						}},
		// 				}},
		// 				Steps: []*armchaos.ExperimentStep{
		// 					{
		// 						Name: to.Ptr("step1"),
		// 						Branches: []*armchaos.ExperimentBranch{
		// 							{
		// 								Name: to.Ptr("branch1"),
		// 								Actions: []armchaos.ExperimentActionClassification{
		// 									&armchaos.ContinuousAction{
		// 										Name: to.Ptr("urn:csci:provider:providername:Shutdown/1.0"),
		// 										Type: to.Ptr("continuous"),
		// 										Duration: to.Ptr("PT10M"),
		// 										Parameters: []*armchaos.KeyValuePair{
		// 											{
		// 												Key: to.Ptr("abruptShutdown"),
		// 												Value: to.Ptr("false"),
		// 										}},
		// 										SelectorID: to.Ptr("selector1"),
		// 								}},
		// 						}},
		// 				}},
		// 			},
		// 			SystemData: &armchaos.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-01T00:00:00.000Z"); return t}()),
		// 				CreatedBy: to.Ptr("User"),
		// 				CreatedByType: to.Ptr(armchaos.CreatedByType("b3a41dba-4415-4d36-9ee8-e5eaa86db976")),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-01T00:00:00.000Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("User"),
		// 				LastModifiedByType: to.Ptr(armchaos.CreatedByType("b3a41dba-4415-4d36-9ee8-e5eaa86db976")),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/chaos/resource-manager/Microsoft.Chaos/stable/2023-11-01/examples/DeleteExperiment.json
func ExampleExperimentsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExperimentsClient().BeginDelete(ctx, "exampleRG", "exampleExperiment", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/chaos/resource-manager/Microsoft.Chaos/stable/2023-11-01/examples/GetExperiment.json
func ExampleExperimentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExperimentsClient().Get(ctx, "exampleRG", "exampleExperiment", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Experiment = armchaos.Experiment{
	// 	Name: to.Ptr("exampleExperiment"),
	// 	Type: to.Ptr("Microsoft.Chaos/experiments"),
	// 	ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/experiments/exampleExperiment"),
	// 	Location: to.Ptr("centraluseuap"),
	// 	Identity: &armchaos.ResourceIdentity{
	// 		Type: to.Ptr(armchaos.ResourceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("d04ab567-2c07-43ef-a7f4-4527626b7f56"),
	// 		TenantID: to.Ptr("8c3e2fb2-fe7a-4bf1-b779-d73990782fe6"),
	// 	},
	// 	Properties: &armchaos.ExperimentProperties{
	// 		Selectors: []armchaos.TargetSelectorClassification{
	// 			&armchaos.TargetListSelector{
	// 				Type: to.Ptr(armchaos.SelectorTypeList),
	// 				ID: to.Ptr("selector1"),
	// 				Targets: []*armchaos.TargetReference{
	// 					{
	// 						Type: to.Ptr(armchaos.TargetReferenceTypeChaosTarget),
	// 						ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/exampleVM/providers/Microsoft.Chaos/targets/Microsoft-VirtualMachine"),
	// 				}},
	// 		}},
	// 		Steps: []*armchaos.ExperimentStep{
	// 			{
	// 				Name: to.Ptr("step1"),
	// 				Branches: []*armchaos.ExperimentBranch{
	// 					{
	// 						Name: to.Ptr("branch1"),
	// 						Actions: []armchaos.ExperimentActionClassification{
	// 							&armchaos.ContinuousAction{
	// 								Name: to.Ptr("urn:csci:provider:providername:Shutdown/1.0"),
	// 								Type: to.Ptr("continuous"),
	// 								Duration: to.Ptr("PT10M"),
	// 								Parameters: []*armchaos.KeyValuePair{
	// 									{
	// 										Key: to.Ptr("abruptShutdown"),
	// 										Value: to.Ptr("false"),
	// 								}},
	// 								SelectorID: to.Ptr("selector1"),
	// 						}},
	// 				}},
	// 		}},
	// 	},
	// 	SystemData: &armchaos.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-01T00:00:00.000Z"); return t}()),
	// 		CreatedBy: to.Ptr("User"),
	// 		CreatedByType: to.Ptr(armchaos.CreatedByType("b3a41dba-4415-4d36-9ee8-e5eaa86db976")),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-01T00:00:00.000Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("User"),
	// 		LastModifiedByType: to.Ptr(armchaos.CreatedByType("b3a41dba-4415-4d36-9ee8-e5eaa86db976")),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/chaos/resource-manager/Microsoft.Chaos/stable/2023-11-01/examples/CreateUpdateExperiment.json
func ExampleExperimentsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExperimentsClient().BeginCreateOrUpdate(ctx, "exampleRG", "exampleExperiment", armchaos.Experiment{
		Location: to.Ptr("eastus2euap"),
		Identity: &armchaos.ResourceIdentity{
			Type: to.Ptr(armchaos.ResourceIdentityTypeSystemAssigned),
		},
		Properties: &armchaos.ExperimentProperties{
			Selectors: []armchaos.TargetSelectorClassification{
				&armchaos.TargetListSelector{
					Type: to.Ptr(armchaos.SelectorTypeList),
					ID:   to.Ptr("selector1"),
					Targets: []*armchaos.TargetReference{
						{
							Type: to.Ptr(armchaos.TargetReferenceTypeChaosTarget),
							ID:   to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/exampleVM/providers/Microsoft.Chaos/targets/Microsoft-VirtualMachine"),
						}},
				}},
			Steps: []*armchaos.ExperimentStep{
				{
					Name: to.Ptr("step1"),
					Branches: []*armchaos.ExperimentBranch{
						{
							Name: to.Ptr("branch1"),
							Actions: []armchaos.ExperimentActionClassification{
								&armchaos.ContinuousAction{
									Name:     to.Ptr("urn:csci:microsoft:virtualMachine:shutdown/1.0"),
									Type:     to.Ptr("continuous"),
									Duration: to.Ptr("PT10M"),
									Parameters: []*armchaos.KeyValuePair{
										{
											Key:   to.Ptr("abruptShutdown"),
											Value: to.Ptr("false"),
										}},
									SelectorID: to.Ptr("selector1"),
								}},
						}},
				}},
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
	// res.Experiment = armchaos.Experiment{
	// 	Name: to.Ptr("exampleExperiment"),
	// 	Type: to.Ptr("Microsoft.Chaos/experiments"),
	// 	ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/experiments/exampleExperiment"),
	// 	Location: to.Ptr("eastus2euap"),
	// 	Identity: &armchaos.ResourceIdentity{
	// 		Type: to.Ptr(armchaos.ResourceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("d04ab567-2c07-43ef-a7f4-4527626b7f56"),
	// 		TenantID: to.Ptr("8c3e2fb2-fe7a-4bf1-b779-d73990782fe6"),
	// 	},
	// 	Properties: &armchaos.ExperimentProperties{
	// 		ProvisioningState: to.Ptr(armchaos.ProvisioningStateSucceeded),
	// 		Selectors: []armchaos.TargetSelectorClassification{
	// 			&armchaos.TargetListSelector{
	// 				Type: to.Ptr(armchaos.SelectorTypeList),
	// 				ID: to.Ptr("selector1"),
	// 				Targets: []*armchaos.TargetReference{
	// 					{
	// 						Type: to.Ptr(armchaos.TargetReferenceTypeChaosTarget),
	// 						ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/exampleVM/providers/Microsoft.Chaos/targets/Microsoft-VirtualMachine"),
	// 				}},
	// 		}},
	// 		Steps: []*armchaos.ExperimentStep{
	// 			{
	// 				Name: to.Ptr("step1"),
	// 				Branches: []*armchaos.ExperimentBranch{
	// 					{
	// 						Name: to.Ptr("branch1"),
	// 						Actions: []armchaos.ExperimentActionClassification{
	// 							&armchaos.ContinuousAction{
	// 								Name: to.Ptr("urn:csci:microsoft:virtualMachine:shutdown/1.0"),
	// 								Type: to.Ptr("continuous"),
	// 								Duration: to.Ptr("PT10M"),
	// 								Parameters: []*armchaos.KeyValuePair{
	// 									{
	// 										Key: to.Ptr("abruptShutdown"),
	// 										Value: to.Ptr("false"),
	// 								}},
	// 								SelectorID: to.Ptr("selector1"),
	// 						}},
	// 				}},
	// 		}},
	// 	},
	// 	SystemData: &armchaos.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-01T00:00:00.000Z"); return t}()),
	// 		CreatedBy: to.Ptr("User"),
	// 		CreatedByType: to.Ptr(armchaos.CreatedByType("b3a41dba-4415-4d36-9ee8-e5eaa86db976")),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-01T00:00:00.000Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("User"),
	// 		LastModifiedByType: to.Ptr(armchaos.CreatedByType("b3a41dba-4415-4d36-9ee8-e5eaa86db976")),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/chaos/resource-manager/Microsoft.Chaos/stable/2023-11-01/examples/UpdateExperiment.json
func ExampleExperimentsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExperimentsClient().BeginUpdate(ctx, "exampleRG", "exampleExperiment", armchaos.ExperimentUpdate{
		Identity: &armchaos.ResourceIdentity{
			Type: to.Ptr(armchaos.ResourceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armchaos.UserAssignedIdentity{
				"/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.ManagedIdentity/userAssignedIdentity/exampleUMI": {},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/chaos/resource-manager/Microsoft.Chaos/stable/2023-11-01/examples/CancelExperiment.json
func ExampleExperimentsClient_BeginCancel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExperimentsClient().BeginCancel(ctx, "exampleRG", "exampleExperiment", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/chaos/resource-manager/Microsoft.Chaos/stable/2023-11-01/examples/StartExperiment.json
func ExampleExperimentsClient_BeginStart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExperimentsClient().BeginStart(ctx, "exampleRG", "exampleExperiment", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/chaos/resource-manager/Microsoft.Chaos/stable/2023-11-01/examples/ListExperimentExecutions.json
func ExampleExperimentsClient_NewListAllExecutionsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExperimentsClient().NewListAllExecutionsPager("exampleRG", "exampleExperiment", nil)
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
		// page.ExperimentExecutionListResult = armchaos.ExperimentExecutionListResult{
		// 	Value: []*armchaos.ExperimentExecution{
		// 		{
		// 			Name: to.Ptr("f24500ad-744e-4a26-864b-b76199eac333"),
		// 			Type: to.Ptr("Microsoft.Chaos/experiments/executions"),
		// 			ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/experiments/exampleExperiment/executions/"),
		// 			Properties: &armchaos.ExperimentExecutionProperties{
		// 				StartedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T21:52:52.255Z"); return t}()),
		// 				Status: to.Ptr("failed"),
		// 				StoppedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T21:56:18.928Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("14d98367-52ef-4596-be4f-53fc81bbfc33"),
		// 			Type: to.Ptr("Microsoft.Chaos/experiments/executions"),
		// 			ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/experiments/exampleExperiment/executionDetails/14d98367-52ef-4596-be4f-53fc81bbfc33"),
		// 			Properties: &armchaos.ExperimentExecutionProperties{
		// 				StartedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T21:52:52.255Z"); return t}()),
		// 				Status: to.Ptr("success"),
		// 				StoppedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T21:56:18.928Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/chaos/resource-manager/Microsoft.Chaos/stable/2023-11-01/examples/GetExperimentExecution.json
func ExampleExperimentsClient_GetExecution() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExperimentsClient().GetExecution(ctx, "exampleRG", "exampleExperiment", "f24500ad-744e-4a26-864b-b76199eac333", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ExperimentExecution = armchaos.ExperimentExecution{
	// 	Name: to.Ptr("f24500ad-744e-4a26-864b-b76199eac333"),
	// 	Type: to.Ptr("Microsoft.Chaos/experiments/executions"),
	// 	ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/experiments/exampleExperiment/executions/f24500ad-744e-4a26-864b-b76199eac333"),
	// 	Properties: &armchaos.ExperimentExecutionProperties{
	// 		StartedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T21:52:52.255Z"); return t}()),
	// 		Status: to.Ptr("failed"),
	// 		StoppedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T21:56:18.928Z"); return t}()),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/chaos/resource-manager/Microsoft.Chaos/stable/2023-11-01/examples/DetailsExperiment.json
func ExampleExperimentsClient_ExecutionDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExperimentsClient().ExecutionDetails(ctx, "exampleRG", "exampleExperiment", "f24500ad-744e-4a26-864b-b76199eac333", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ExperimentExecutionDetails = armchaos.ExperimentExecutionDetails{
	// 	Name: to.Ptr("f24500ad-744e-4a26-864b-b76199eac333"),
	// 	Type: to.Ptr("Microsoft.Chaos/experiments/executions/getExecutionDetails"),
	// 	ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/experiments/exampleExperiment/executions/f24500ad-744e-4a26-864b-b76199eac333/getExecutionDetails"),
	// 	Properties: &armchaos.ExperimentExecutionDetailsProperties{
	// 		StartedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T21:52:52.255Z"); return t}()),
	// 		Status: to.Ptr("failed"),
	// 		StoppedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T21:56:18.928Z"); return t}()),
	// 		FailureReason: to.Ptr("Dependency failure"),
	// 		LastActionAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T21:52:52.255Z"); return t}()),
	// 		RunInformation: &armchaos.ExperimentExecutionDetailsPropertiesRunInformation{
	// 			Steps: []*armchaos.StepStatus{
	// 				{
	// 					Branches: []*armchaos.BranchStatus{
	// 						{
	// 							Actions: []*armchaos.ActionStatus{
	// 								{
	// 									ActionID: to.Ptr("59499d33-6751-4b6e-a1f6-58f4d56a040a"),
	// 									ActionName: to.Ptr("urn:provider:agent-v2:Microsoft.Azure.Chaos.Fault.CPUPressureAllProcessors"),
	// 									EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T21:56:13.627Z"); return t}()),
	// 									StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T21:56:13.627Z"); return t}()),
	// 									Status: to.Ptr("failed"),
	// 									Targets: []*armchaos.ExperimentExecutionActionTargetDetailsProperties{
	// 										{
	// 											Status: to.Ptr("succeeded"),
	// 											Target: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/VM1"),
	// 											TargetCompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T17:30:55.000Z"); return t}()),
	// 											TargetFailedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T16:30:55.000Z"); return t}()),
	// 										},
	// 										{
	// 											Status: to.Ptr("failed"),
	// 											Target: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/VM1"),
	// 											TargetCompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T17:30:55.000Z"); return t}()),
	// 											TargetFailedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T16:30:55.000Z"); return t}()),
	// 									}},
	// 							}},
	// 							BranchID: to.Ptr("FirstBranch"),
	// 							BranchName: to.Ptr("FirstBranch"),
	// 							Status: to.Ptr("failed"),
	// 					}},
	// 					Status: to.Ptr("failed"),
	// 					StepID: to.Ptr("FirstStep"),
	// 					StepName: to.Ptr("FirstStep"),
	// 			}},
	// 		},
	// 	},
	// }
}
