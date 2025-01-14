//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armplaywrighttesting_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/playwrighttesting/armplaywrighttesting"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/2bde125befabb21807a2021765901f20e3e74ec8/specification/playwrighttesting/resource-manager/Microsoft.AzurePlaywrightService/stable/2024-12-01/examples/Quotas_ListBySubscription.json
func ExampleQuotasClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armplaywrighttesting.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewQuotasClient().NewListBySubscriptionPager("eastus", nil)
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
		// page.QuotaListResult = armplaywrighttesting.QuotaListResult{
		// 	Value: []*armplaywrighttesting.Quota{
		// 		{
		// 			Name: to.Ptr("ScalableExecution"),
		// 			Type: to.Ptr("Microsoft.AzurePlaywrightService/Locations/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.AzurePlaywrightService/locations/eastus/quotas/ScalableExecution"),
		// 			Properties: &armplaywrighttesting.QuotaProperties{
		// 				FreeTrial: &armplaywrighttesting.FreeTrialProperties{
		// 					AccountID: to.Ptr("3fa85f64-5717-4562-b3fc-2c963f66afa6"),
		// 					State: to.Ptr(armplaywrighttesting.FreeTrialStateActive),
		// 				},
		// 				OfferingType: to.Ptr(armplaywrighttesting.OfferingTypeNotApplicable),
		// 				ProvisioningState: to.Ptr(armplaywrighttesting.ProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Reporting"),
		// 			Type: to.Ptr("Microsoft.AzurePlaywrightService/Locations/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.AzurePlaywrightService/locations/eastus/quotas/Reporting"),
		// 			Properties: &armplaywrighttesting.QuotaProperties{
		// 				FreeTrial: &armplaywrighttesting.FreeTrialProperties{
		// 					AccountID: to.Ptr("3fa85f64-5717-4562-b3fc-2c963f66afa6"),
		// 					State: to.Ptr(armplaywrighttesting.FreeTrialStateActive),
		// 				},
		// 				OfferingType: to.Ptr(armplaywrighttesting.OfferingTypePrivatePreview),
		// 				ProvisioningState: to.Ptr(armplaywrighttesting.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/2bde125befabb21807a2021765901f20e3e74ec8/specification/playwrighttesting/resource-manager/Microsoft.AzurePlaywrightService/stable/2024-12-01/examples/Quotas_Get.json
func ExampleQuotasClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armplaywrighttesting.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewQuotasClient().Get(ctx, "eastus", armplaywrighttesting.QuotaNamesScalableExecution, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Quota = armplaywrighttesting.Quota{
	// 	Name: to.Ptr("ScalableExecution"),
	// 	Type: to.Ptr("Microsoft.AzurePlaywrightService/Locations/Quotas"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.AzurePlaywrightService/locations/eastus/quotas/ScalableExecution"),
	// 	Properties: &armplaywrighttesting.QuotaProperties{
	// 		FreeTrial: &armplaywrighttesting.FreeTrialProperties{
	// 			AccountID: to.Ptr("3fa85f64-5717-4562-b3fc-2c963f66afa6"),
	// 			State: to.Ptr(armplaywrighttesting.FreeTrialStateActive),
	// 		},
	// 		OfferingType: to.Ptr(armplaywrighttesting.OfferingTypeNotApplicable),
	// 		ProvisioningState: to.Ptr(armplaywrighttesting.ProvisioningStateSucceeded),
	// 	},
	// }
}
