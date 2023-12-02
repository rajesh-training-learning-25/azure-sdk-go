//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurity_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/32c63d685a0b03817a504b04be938ce46d06ac19/specification/security/resource-manager/Microsoft.Security/preview/2022-01-01-preview/examples/GovernanceAssignments/ListGovernanceAssignments_example.json
func ExampleGovernanceAssignmentsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGovernanceAssignmentsClient().NewListPager("subscriptions/c32e05d9-7207-4e22-bdf4-4f7d9c72e5fd", "6b9421dd-5555-2251-9b3d-2be58e2f82cd", nil)
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
		// page.GovernanceAssignmentsList = armsecurity.GovernanceAssignmentsList{
		// 	Value: []*armsecurity.GovernanceAssignment{
		// 		{
		// 			Name: to.Ptr("6634ff9f-127b-4bf2-8e6e-b1737f5e789c"),
		// 			Type: to.Ptr("Microsoft.Security/assessments/governanceAssignments"),
		// 			ID: to.Ptr("/subscriptions/c32e05d9-7207-4e22-bdf4-4f7d9c72e5fd/resourceGroups/compute_servers/providers/Microsoft.Compute/virtualMachines/win2012/providers/Microsoft.Security/assessments/6b9421dd-5555-2251-9b3d-2be58e2f82cd/governanceAssignments/6634ff9f-127b-4bf2-8e6e-b1737f5e789c"),
		// 			Properties: &armsecurity.GovernanceAssignmentProperties{
		// 				AdditionalData: &armsecurity.GovernanceAssignmentAdditionalData{
		// 					TicketLink: to.Ptr("https://snow.com"),
		// 					TicketNumber: to.Ptr[int32](123123),
		// 					TicketStatus: to.Ptr("Active"),
		// 				},
		// 				IsGracePeriod: to.Ptr(true),
		// 				Owner: to.Ptr("user@contoso.com"),
		// 				RemediationDueDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-07T13:00:00.000Z"); return t}()),
		// 				RemediationEta: &armsecurity.RemediationEta{
		// 					Eta: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-08T13:00:00.000Z"); return t}()),
		// 					Justification: to.Ptr("Justification of ETA"),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("41fb92a5-43dc-4c00-a969-469c16cef7a7"),
		// 			Type: to.Ptr("Microsoft.Security/assessments/governanceAssignments"),
		// 			ID: to.Ptr("/subscriptions/c32e05d9-7207-4e22-bdf4-4f7d9c72e5fd/resourceGroups/compute_servers/providers/Microsoft.Compute/virtualMachines/win2011/providers/Microsoft.Security/assessments/6b9421dd-5555-2251-9b3d-2be58e2f82cd/governanceAssignments/41fb92a5-43dc-4c00-a969-469c16cef7a7"),
		// 			Properties: &armsecurity.GovernanceAssignmentProperties{
		// 				AdditionalData: &armsecurity.GovernanceAssignmentAdditionalData{
		// 					TicketLink: to.Ptr("https://snow.com"),
		// 					TicketNumber: to.Ptr[int32](653424),
		// 					TicketStatus: to.Ptr("Active"),
		// 				},
		// 				IsGracePeriod: to.Ptr(true),
		// 				Owner: to.Ptr("user2@contoso.com"),
		// 				RemediationDueDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-07T13:00:00.000Z"); return t}()),
		// 				RemediationEta: &armsecurity.RemediationEta{
		// 					Eta: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-08T13:00:00.000Z"); return t}()),
		// 					Justification: to.Ptr("Justification of ETA"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/32c63d685a0b03817a504b04be938ce46d06ac19/specification/security/resource-manager/Microsoft.Security/preview/2022-01-01-preview/examples/GovernanceAssignments/GetGovernanceAssignment_example.json
func ExampleGovernanceAssignmentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGovernanceAssignmentsClient().Get(ctx, "subscriptions/c32e05d9-7207-4e22-bdf4-4f7d9c72e5fd/resourceGroups/compute_servers/providers/Microsoft.Compute/virtualMachines/win2012", "6b9421dd-5555-2251-9b3d-2be58e2f82cd", "6634ff9f-127b-4bf2-8e6e-b1737f5e789c", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GovernanceAssignment = armsecurity.GovernanceAssignment{
	// 	Name: to.Ptr("6634ff9f-127b-4bf2-8e6e-b1737f5e789c"),
	// 	Type: to.Ptr("Microsoft.Security/assessments/governanceAssignments"),
	// 	ID: to.Ptr("/subscriptions/c32e05d9-7207-4e22-bdf4-4f7d9c72e5fd/resourceGroups/compute_servers/providers/Microsoft.Compute/virtualMachines/win2012/providers/Microsoft.Security/assessments/6b9421dd-5555-2251-9b3d-2be58e2f82cd/governanceAssignments/6634ff9f-127b-4bf2-8e6e-b1737f5e789c"),
	// 	Properties: &armsecurity.GovernanceAssignmentProperties{
	// 		AdditionalData: &armsecurity.GovernanceAssignmentAdditionalData{
	// 			TicketLink: to.Ptr("https://snow.com"),
	// 			TicketNumber: to.Ptr[int32](123123),
	// 			TicketStatus: to.Ptr("Active"),
	// 		},
	// 		GovernanceEmailNotification: &armsecurity.GovernanceEmailNotification{
	// 			DisableManagerEmailNotification: to.Ptr(false),
	// 			DisableOwnerEmailNotification: to.Ptr(false),
	// 		},
	// 		IsGracePeriod: to.Ptr(true),
	// 		Owner: to.Ptr("user@contoso.com"),
	// 		RemediationDueDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-07T13:00:00.000Z"); return t}()),
	// 		RemediationEta: &armsecurity.RemediationEta{
	// 			Eta: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-08T13:00:00.000Z"); return t}()),
	// 			Justification: to.Ptr("Justification of ETA"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/32c63d685a0b03817a504b04be938ce46d06ac19/specification/security/resource-manager/Microsoft.Security/preview/2022-01-01-preview/examples/GovernanceAssignments/PutGovernanceAssignment_example.json
func ExampleGovernanceAssignmentsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGovernanceAssignmentsClient().CreateOrUpdate(ctx, "subscriptions/c32e05d9-7207-4e22-bdf4-4f7d9c72e5fd/resourceGroups/compute_servers/providers/Microsoft.Compute/virtualMachines/win2012", "6b9421dd-5555-2251-9b3d-2be58e2f82cd", "6634ff9f-127b-4bf2-8e6e-b1737f5e789c", armsecurity.GovernanceAssignment{
		Properties: &armsecurity.GovernanceAssignmentProperties{
			AdditionalData: &armsecurity.GovernanceAssignmentAdditionalData{
				TicketLink:   to.Ptr("https://snow.com"),
				TicketNumber: to.Ptr[int32](123123),
				TicketStatus: to.Ptr("Active"),
			},
			GovernanceEmailNotification: &armsecurity.GovernanceEmailNotification{
				DisableManagerEmailNotification: to.Ptr(false),
				DisableOwnerEmailNotification:   to.Ptr(false),
			},
			IsGracePeriod:      to.Ptr(true),
			Owner:              to.Ptr("user@contoso.com"),
			RemediationDueDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-07T13:00:00.000Z"); return t }()),
			RemediationEta: &armsecurity.RemediationEta{
				Eta:           to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-08T13:00:00.000Z"); return t }()),
				Justification: to.Ptr("Justification of ETA"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GovernanceAssignment = armsecurity.GovernanceAssignment{
	// 	Name: to.Ptr("6634ff9f-127b-4bf2-8e6e-b1737f5e789c"),
	// 	Type: to.Ptr("Microsoft.Security/assessments/governanceAssignments"),
	// 	ID: to.Ptr("/subscriptions/c32e05d9-7207-4e22-bdf4-4f7d9c72e5fd/resourceGroups/compute_servers/providers/Microsoft.Compute/virtualMachines/win2012/providers/Microsoft.Security/assessments/6b9421dd-5555-2251-9b3d-2be58e2f82cd/governanceAssignments/6634ff9f-127b-4bf2-8e6e-b1737f5e789c"),
	// 	Properties: &armsecurity.GovernanceAssignmentProperties{
	// 		AdditionalData: &armsecurity.GovernanceAssignmentAdditionalData{
	// 			TicketLink: to.Ptr("https://snow.com"),
	// 			TicketNumber: to.Ptr[int32](123123),
	// 			TicketStatus: to.Ptr("Active"),
	// 		},
	// 		GovernanceEmailNotification: &armsecurity.GovernanceEmailNotification{
	// 			DisableManagerEmailNotification: to.Ptr(false),
	// 			DisableOwnerEmailNotification: to.Ptr(false),
	// 		},
	// 		IsGracePeriod: to.Ptr(true),
	// 		Owner: to.Ptr("user@contoso.com"),
	// 		RemediationDueDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-07T13:00:00.000Z"); return t}()),
	// 		RemediationEta: &armsecurity.RemediationEta{
	// 			Eta: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-08T13:00:00.000Z"); return t}()),
	// 			Justification: to.Ptr("Justification of ETA"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/32c63d685a0b03817a504b04be938ce46d06ac19/specification/security/resource-manager/Microsoft.Security/preview/2022-01-01-preview/examples/GovernanceAssignments/DeleteGovernanceAssignment_example.json
func ExampleGovernanceAssignmentsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewGovernanceAssignmentsClient().Delete(ctx, "subscriptions/c32e05d9-7207-4e22-bdf4-4f7d9c72e5fd/resourceGroups/compute_servers/providers/Microsoft.Compute/virtualMachines/win2012", "6b9421dd-5555-2251-9b3d-2be58e2f82cd", "6634ff9f-127b-4bf2-8e6e-b1737f5e789c", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
