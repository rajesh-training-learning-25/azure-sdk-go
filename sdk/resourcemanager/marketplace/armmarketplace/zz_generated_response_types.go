//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmarketplace

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}

// PrivateStoreClientAcknowledgeOfferNotificationResponse contains the response from method PrivateStoreClient.AcknowledgeOfferNotification.
type PrivateStoreClientAcknowledgeOfferNotificationResponse struct {
	// placeholder for future response values
}

// PrivateStoreClientAdminRequestApprovalsListResponse contains the response from method PrivateStoreClient.AdminRequestApprovalsList.
type PrivateStoreClientAdminRequestApprovalsListResponse struct {
	AdminRequestApprovalsList
}

// PrivateStoreClientAnyExistingOffersInTheCollectionsResponse contains the response from method PrivateStoreClient.AnyExistingOffersInTheCollections.
type PrivateStoreClientAnyExistingOffersInTheCollectionsResponse struct {
	AnyExistingOffersInTheCollectionsResponse
}

// PrivateStoreClientBillingAccountsResponse contains the response from method PrivateStoreClient.BillingAccounts.
type PrivateStoreClientBillingAccountsResponse struct {
	BillingAccountsResponse
}

// PrivateStoreClientBulkCollectionsActionResponse contains the response from method PrivateStoreClient.BulkCollectionsAction.
type PrivateStoreClientBulkCollectionsActionResponse struct {
	BulkCollectionsResponse
}

// PrivateStoreClientCollectionsToSubscriptionsMappingResponse contains the response from method PrivateStoreClient.CollectionsToSubscriptionsMapping.
type PrivateStoreClientCollectionsToSubscriptionsMappingResponse struct {
	CollectionsToSubscriptionsMappingResponse
}

// PrivateStoreClientCreateApprovalRequestResponse contains the response from method PrivateStoreClient.CreateApprovalRequest.
type PrivateStoreClientCreateApprovalRequestResponse struct {
	RequestApprovalResource
}

// PrivateStoreClientCreateOrUpdateResponse contains the response from method PrivateStoreClient.CreateOrUpdate.
type PrivateStoreClientCreateOrUpdateResponse struct {
	// placeholder for future response values
}

// PrivateStoreClientDeleteResponse contains the response from method PrivateStoreClient.Delete.
type PrivateStoreClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateStoreClientFetchAllSubscriptionsInTenantResponse contains the response from method PrivateStoreClient.FetchAllSubscriptionsInTenant.
type PrivateStoreClientFetchAllSubscriptionsInTenantResponse struct {
	SubscriptionsResponse
}

// PrivateStoreClientGetAdminRequestApprovalResponse contains the response from method PrivateStoreClient.GetAdminRequestApproval.
type PrivateStoreClientGetAdminRequestApprovalResponse struct {
	AdminRequestApprovalsResource
}

// PrivateStoreClientGetApprovalRequestsListResponse contains the response from method PrivateStoreClient.GetApprovalRequestsList.
type PrivateStoreClientGetApprovalRequestsListResponse struct {
	RequestApprovalsList
}

// PrivateStoreClientGetRequestApprovalResponse contains the response from method PrivateStoreClient.GetRequestApproval.
type PrivateStoreClientGetRequestApprovalResponse struct {
	RequestApprovalResource
}

// PrivateStoreClientGetResponse contains the response from method PrivateStoreClient.Get.
type PrivateStoreClientGetResponse struct {
	PrivateStore
}

// PrivateStoreClientListNewPlansNotificationsResponse contains the response from method PrivateStoreClient.ListNewPlansNotifications.
type PrivateStoreClientListNewPlansNotificationsResponse struct {
	NewPlansNotificationsList
}

// PrivateStoreClientListResponse contains the response from method PrivateStoreClient.List.
type PrivateStoreClientListResponse struct {
	PrivateStoreList
}

// PrivateStoreClientListStopSellOffersPlansNotificationsResponse contains the response from method PrivateStoreClient.ListStopSellOffersPlansNotifications.
type PrivateStoreClientListStopSellOffersPlansNotificationsResponse struct {
	StopSellOffersPlansNotificationsList
}

// PrivateStoreClientListSubscriptionsContextResponse contains the response from method PrivateStoreClient.ListSubscriptionsContext.
type PrivateStoreClientListSubscriptionsContextResponse struct {
	SubscriptionsContextList
}

// PrivateStoreClientQueryApprovedPlansResponse contains the response from method PrivateStoreClient.QueryApprovedPlans.
type PrivateStoreClientQueryApprovedPlansResponse struct {
	QueryApprovedPlansResponse
}

// PrivateStoreClientQueryNotificationsStateResponse contains the response from method PrivateStoreClient.QueryNotificationsState.
type PrivateStoreClientQueryNotificationsStateResponse struct {
	PrivateStoreNotificationsState
}

// PrivateStoreClientQueryOffersResponse contains the response from method PrivateStoreClient.QueryOffers.
type PrivateStoreClientQueryOffersResponse struct {
	QueryOffers
}

// PrivateStoreClientQueryRequestApprovalResponse contains the response from method PrivateStoreClient.QueryRequestApproval.
type PrivateStoreClientQueryRequestApprovalResponse struct {
	QueryRequestApproval
}

// PrivateStoreClientQueryUserOffersResponse contains the response from method PrivateStoreClient.QueryUserOffers.
type PrivateStoreClientQueryUserOffersResponse struct {
	QueryOffers
}

// PrivateStoreClientUpdateAdminRequestApprovalResponse contains the response from method PrivateStoreClient.UpdateAdminRequestApproval.
type PrivateStoreClientUpdateAdminRequestApprovalResponse struct {
	AdminRequestApprovalsResource
}

// PrivateStoreClientWithdrawPlanResponse contains the response from method PrivateStoreClient.WithdrawPlan.
type PrivateStoreClientWithdrawPlanResponse struct {
	// placeholder for future response values
}

// PrivateStoreCollectionClientApproveAllItemsResponse contains the response from method PrivateStoreCollectionClient.ApproveAllItems.
type PrivateStoreCollectionClientApproveAllItemsResponse struct {
	Collection
}

// PrivateStoreCollectionClientCreateOrUpdateResponse contains the response from method PrivateStoreCollectionClient.CreateOrUpdate.
type PrivateStoreCollectionClientCreateOrUpdateResponse struct {
	Collection
}

// PrivateStoreCollectionClientDeleteResponse contains the response from method PrivateStoreCollectionClient.Delete.
type PrivateStoreCollectionClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateStoreCollectionClientDisableApproveAllItemsResponse contains the response from method PrivateStoreCollectionClient.DisableApproveAllItems.
type PrivateStoreCollectionClientDisableApproveAllItemsResponse struct {
	Collection
}

// PrivateStoreCollectionClientGetResponse contains the response from method PrivateStoreCollectionClient.Get.
type PrivateStoreCollectionClientGetResponse struct {
	Collection
}

// PrivateStoreCollectionClientListResponse contains the response from method PrivateStoreCollectionClient.List.
type PrivateStoreCollectionClientListResponse struct {
	CollectionsList
}

// PrivateStoreCollectionClientPostResponse contains the response from method PrivateStoreCollectionClient.Post.
type PrivateStoreCollectionClientPostResponse struct {
	// placeholder for future response values
}

// PrivateStoreCollectionClientTransferOffersResponse contains the response from method PrivateStoreCollectionClient.TransferOffers.
type PrivateStoreCollectionClientTransferOffersResponse struct {
	TransferOffersResponse
}

// PrivateStoreCollectionOfferClientCreateOrUpdateResponse contains the response from method PrivateStoreCollectionOfferClient.CreateOrUpdate.
type PrivateStoreCollectionOfferClientCreateOrUpdateResponse struct {
	Offer
}

// PrivateStoreCollectionOfferClientDeleteResponse contains the response from method PrivateStoreCollectionOfferClient.Delete.
type PrivateStoreCollectionOfferClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateStoreCollectionOfferClientGetResponse contains the response from method PrivateStoreCollectionOfferClient.Get.
type PrivateStoreCollectionOfferClientGetResponse struct {
	Offer
}

// PrivateStoreCollectionOfferClientListResponse contains the response from method PrivateStoreCollectionOfferClient.List.
type PrivateStoreCollectionOfferClientListResponse struct {
	OfferListResponse
}

// PrivateStoreCollectionOfferClientPostResponse contains the response from method PrivateStoreCollectionOfferClient.Post.
type PrivateStoreCollectionOfferClientPostResponse struct {
	// placeholder for future response values
}

// PrivateStoreCollectionOfferClientUpsertOfferWithMultiContextResponse contains the response from method PrivateStoreCollectionOfferClient.UpsertOfferWithMultiContext.
type PrivateStoreCollectionOfferClientUpsertOfferWithMultiContextResponse struct {
	Offer
}
