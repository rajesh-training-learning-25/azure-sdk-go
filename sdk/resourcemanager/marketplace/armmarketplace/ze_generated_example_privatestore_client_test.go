//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"
)

// x-ms-original-file: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-06-01/examples/GetPrivateStores.json
func ExamplePrivateStoreClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplace.NewPrivateStoreClient(cred, nil)
	pager := client.List(&armmarketplace.PrivateStoreListOptions{UseCache: nil})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("PrivateStore.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-06-01/examples/GetPrivateStore.json
func ExamplePrivateStoreClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplace.NewPrivateStoreClient(cred, nil)
	res, err := client.Get(ctx,
		"<private-store-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("PrivateStore.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-06-01/examples/PrivateStores_update.json
func ExamplePrivateStoreClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplace.NewPrivateStoreClient(cred, nil)
	_, err = client.CreateOrUpdate(ctx,
		"<private-store-id>",
		&armmarketplace.PrivateStoreCreateOrUpdateOptions{Payload: &armmarketplace.PrivateStore{
			Properties: &armmarketplace.PrivateStoreProperties{
				Availability: armmarketplace.AvailabilityDisabled.ToPtr(),
				ETag:         to.StringPtr("<etag>"),
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-06-01/examples/DeletePrivateStore.json
func ExamplePrivateStoreClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplace.NewPrivateStoreClient(cred, nil)
	_, err = client.Delete(ctx,
		"<private-store-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-06-01/examples/QueryOffers.json
func ExamplePrivateStoreClient_QueryOffers() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplace.NewPrivateStoreClient(cred, nil)
	_, err = client.QueryOffers(ctx,
		"<private-store-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-06-01/examples/BillingAccounts.json
func ExamplePrivateStoreClient_BillingAccounts() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplace.NewPrivateStoreClient(cred, nil)
	_, err = client.BillingAccounts(ctx,
		"<private-store-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-06-01/examples/CollectionsToSubscriptionsMapping.json
func ExamplePrivateStoreClient_CollectionsToSubscriptionsMapping() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplace.NewPrivateStoreClient(cred, nil)
	_, err = client.CollectionsToSubscriptionsMapping(ctx,
		"<private-store-id>",
		&armmarketplace.PrivateStoreCollectionsToSubscriptionsMappingOptions{Payload: &armmarketplace.CollectionsToSubscriptionsMappingPayload{
			Properties: &armmarketplace.CollectionsToSubscriptionsMappingProperties{
				SubscriptionIDs: []*string{
					to.StringPtr("b340914e-353d-453a-85fb-8f9b65b51f91"),
					to.StringPtr("f2baa04d-5bfc-461b-b6d8-61b403c9ec48")},
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-06-01/examples/QueryApprovedPlans.json
func ExamplePrivateStoreClient_QueryApprovedPlans() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplace.NewPrivateStoreClient(cred, nil)
	_, err = client.QueryApprovedPlans(ctx,
		"<private-store-id>",
		&armmarketplace.PrivateStoreQueryApprovedPlansOptions{Payload: &armmarketplace.QueryApprovedPlansPayload{
			Properties: &armmarketplace.QueryApprovedPlans{
				OfferID: to.StringPtr("<offer-id>"),
				PlanIDs: []*string{
					to.StringPtr("testPlanA"),
					to.StringPtr("testPlanB"),
					to.StringPtr("testPlanC")},
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-06-01/examples/BulkCollectionsAction.json
func ExamplePrivateStoreClient_BulkCollectionsAction() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplace.NewPrivateStoreClient(cred, nil)
	_, err = client.BulkCollectionsAction(ctx,
		"<private-store-id>",
		&armmarketplace.PrivateStoreBulkCollectionsActionOptions{Payload: &armmarketplace.BulkCollectionsPayload{
			Properties: &armmarketplace.BulkCollectionsDetails{
				Action: to.StringPtr("<action>"),
				CollectionIDs: []*string{
					to.StringPtr("c752f021-1c37-4af5-b82f-74c51c27b44a"),
					to.StringPtr("f47ef1c7-e908-4f39-ae29-db181634ad8d")},
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-06-01/examples/GetApprovalRequestsList.json
func ExamplePrivateStoreClient_GetApprovalRequestsList() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplace.NewPrivateStoreClient(cred, nil)
	_, err = client.GetApprovalRequestsList(ctx,
		"<private-store-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-06-01/examples/GetRequestApproval.json
func ExamplePrivateStoreClient_GetRequestApproval() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplace.NewPrivateStoreClient(cred, nil)
	res, err := client.GetRequestApproval(ctx,
		"<private-store-id>",
		"<request-approval-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("RequestApprovalResource.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-06-01/examples/CreateApprovalRequest.json
func ExamplePrivateStoreClient_CreateApprovalRequest() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplace.NewPrivateStoreClient(cred, nil)
	res, err := client.CreateApprovalRequest(ctx,
		"<private-store-id>",
		"<request-approval-id>",
		&armmarketplace.PrivateStoreCreateApprovalRequestOptions{Payload: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("RequestApprovalResource.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-06-01/examples/QueryRequestApproval.json
func ExamplePrivateStoreClient_QueryRequestApproval() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplace.NewPrivateStoreClient(cred, nil)
	_, err = client.QueryRequestApproval(ctx,
		"<private-store-id>",
		"<request-approval-id>",
		&armmarketplace.PrivateStoreQueryRequestApprovalOptions{Payload: &armmarketplace.QueryRequestApprovalProperties{
			Properties: &armmarketplace.RequestDetails{
				PlanIDs: []*string{
					to.StringPtr("testPlanA"),
					to.StringPtr("testPlanB"),
					to.StringPtr("*")},
				PublisherID: to.StringPtr("<publisher-id>"),
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-06-01/examples/AdminRequestApprovalsList.json
func ExamplePrivateStoreClient_AdminRequestApprovalsList() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplace.NewPrivateStoreClient(cred, nil)
	_, err = client.AdminRequestApprovalsList(ctx,
		"<private-store-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-06-01/examples/GetAdminRequestApproval.json
func ExamplePrivateStoreClient_GetAdminRequestApproval() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplace.NewPrivateStoreClient(cred, nil)
	res, err := client.GetAdminRequestApproval(ctx,
		"<private-store-id>",
		"<admin-request-approval-id>",
		"<publisher-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("AdminRequestApprovalsResource.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-06-01/examples/UpdateAdminRequestApproval.json
func ExamplePrivateStoreClient_UpdateAdminRequestApproval() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplace.NewPrivateStoreClient(cred, nil)
	res, err := client.UpdateAdminRequestApproval(ctx,
		"<private-store-id>",
		"<admin-request-approval-id>",
		&armmarketplace.PrivateStoreUpdateAdminRequestApprovalOptions{Payload: &armmarketplace.AdminRequestApprovalsResource{
			Properties: &armmarketplace.AdminRequestApprovalProperties{
				AdminAction: armmarketplace.AdminActionApproved.ToPtr(),
				ApprovedPlans: []*string{
					to.StringPtr("testPlan")},
				CollectionIDs: []*string{
					to.StringPtr("f8ee227e-85d7-477d-abbf-854d6decaf70"),
					to.StringPtr("39246ad6-c521-4fed-8de7-77dede2e873f")},
				Comment:     to.StringPtr("<comment>"),
				PublisherID: to.StringPtr("<publisher-id>"),
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("AdminRequestApprovalsResource.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-06-01/examples/NotificationsState.json
func ExamplePrivateStoreClient_QueryNotificationsState() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplace.NewPrivateStoreClient(cred, nil)
	_, err = client.QueryNotificationsState(ctx,
		"<private-store-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-06-01/examples/AcknowledgeNotification.json
func ExamplePrivateStoreClient_AcknowledgeOfferNotification() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplace.NewPrivateStoreClient(cred, nil)
	_, err = client.AcknowledgeOfferNotification(ctx,
		"<private-store-id>",
		"<offer-id>",
		&armmarketplace.PrivateStoreAcknowledgeOfferNotificationOptions{Payload: &armmarketplace.AcknowledgeOfferNotificationProperties{
			Properties: &armmarketplace.AcknowledgeOfferNotificationDetails{
				Acknowledge: to.BoolPtr(false),
				AddPlans:    []*string{},
				Dismiss:     to.BoolPtr(false),
				RemoveOffer: to.BoolPtr(false),
				RemovePlans: []*string{
					to.StringPtr("testPlanA")},
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-06-01/examples/WithdrawPlan.json
func ExamplePrivateStoreClient_WithdrawPlan() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplace.NewPrivateStoreClient(cred, nil)
	_, err = client.WithdrawPlan(ctx,
		"<private-store-id>",
		"<request-approval-id>",
		&armmarketplace.PrivateStoreWithdrawPlanOptions{Payload: &armmarketplace.WithdrawProperties{
			Properties: &armmarketplace.WithdrawDetails{
				PlanID:      to.StringPtr("<plan-id>"),
				PublisherID: to.StringPtr("<publisher-id>"),
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
}
