//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmarketplace

import "net/http"

// OperationsListResponse contains the response from method Operations.List.
type OperationsListResponse struct {
	OperationsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsListResult contains the result from method Operations.List.
type OperationsListResult struct {
	OperationListResult
}

// PrivateStoreAcknowledgeOfferNotificationResponse contains the response from method PrivateStore.AcknowledgeOfferNotification.
type PrivateStoreAcknowledgeOfferNotificationResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateStoreAdminRequestApprovalsListResponse contains the response from method PrivateStore.AdminRequestApprovalsList.
type PrivateStoreAdminRequestApprovalsListResponse struct {
	PrivateStoreAdminRequestApprovalsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateStoreAdminRequestApprovalsListResult contains the result from method PrivateStore.AdminRequestApprovalsList.
type PrivateStoreAdminRequestApprovalsListResult struct {
	AdminRequestApprovalsList
}

// PrivateStoreBillingAccountsResponse contains the response from method PrivateStore.BillingAccounts.
type PrivateStoreBillingAccountsResponse struct {
	PrivateStoreBillingAccountsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateStoreBillingAccountsResult contains the result from method PrivateStore.BillingAccounts.
type PrivateStoreBillingAccountsResult struct {
	BillingAccountsResponse
}

// PrivateStoreBulkCollectionsActionResponse contains the response from method PrivateStore.BulkCollectionsAction.
type PrivateStoreBulkCollectionsActionResponse struct {
	PrivateStoreBulkCollectionsActionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateStoreBulkCollectionsActionResult contains the result from method PrivateStore.BulkCollectionsAction.
type PrivateStoreBulkCollectionsActionResult struct {
	BulkCollectionsResponse
}

// PrivateStoreCollectionCreateOrUpdateResponse contains the response from method PrivateStoreCollection.CreateOrUpdate.
type PrivateStoreCollectionCreateOrUpdateResponse struct {
	PrivateStoreCollectionCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateStoreCollectionCreateOrUpdateResult contains the result from method PrivateStoreCollection.CreateOrUpdate.
type PrivateStoreCollectionCreateOrUpdateResult struct {
	Collection
}

// PrivateStoreCollectionDeleteResponse contains the response from method PrivateStoreCollection.Delete.
type PrivateStoreCollectionDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateStoreCollectionGetResponse contains the response from method PrivateStoreCollection.Get.
type PrivateStoreCollectionGetResponse struct {
	PrivateStoreCollectionGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateStoreCollectionGetResult contains the result from method PrivateStoreCollection.Get.
type PrivateStoreCollectionGetResult struct {
	Collection
}

// PrivateStoreCollectionListResponse contains the response from method PrivateStoreCollection.List.
type PrivateStoreCollectionListResponse struct {
	PrivateStoreCollectionListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateStoreCollectionListResult contains the result from method PrivateStoreCollection.List.
type PrivateStoreCollectionListResult struct {
	CollectionsList
}

// PrivateStoreCollectionOfferCreateOrUpdateResponse contains the response from method PrivateStoreCollectionOffer.CreateOrUpdate.
type PrivateStoreCollectionOfferCreateOrUpdateResponse struct {
	PrivateStoreCollectionOfferCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateStoreCollectionOfferCreateOrUpdateResult contains the result from method PrivateStoreCollectionOffer.CreateOrUpdate.
type PrivateStoreCollectionOfferCreateOrUpdateResult struct {
	Offer
}

// PrivateStoreCollectionOfferDeleteResponse contains the response from method PrivateStoreCollectionOffer.Delete.
type PrivateStoreCollectionOfferDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateStoreCollectionOfferGetResponse contains the response from method PrivateStoreCollectionOffer.Get.
type PrivateStoreCollectionOfferGetResponse struct {
	PrivateStoreCollectionOfferGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateStoreCollectionOfferGetResult contains the result from method PrivateStoreCollectionOffer.Get.
type PrivateStoreCollectionOfferGetResult struct {
	Offer
}

// PrivateStoreCollectionOfferListResponse contains the response from method PrivateStoreCollectionOffer.List.
type PrivateStoreCollectionOfferListResponse struct {
	PrivateStoreCollectionOfferListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateStoreCollectionOfferListResult contains the result from method PrivateStoreCollectionOffer.List.
type PrivateStoreCollectionOfferListResult struct {
	OfferListResponse
}

// PrivateStoreCollectionOfferPostResponse contains the response from method PrivateStoreCollectionOffer.Post.
type PrivateStoreCollectionOfferPostResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateStoreCollectionPostResponse contains the response from method PrivateStoreCollection.Post.
type PrivateStoreCollectionPostResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateStoreCollectionTransferOffersResponse contains the response from method PrivateStoreCollection.TransferOffers.
type PrivateStoreCollectionTransferOffersResponse struct {
	PrivateStoreCollectionTransferOffersResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateStoreCollectionTransferOffersResult contains the result from method PrivateStoreCollection.TransferOffers.
type PrivateStoreCollectionTransferOffersResult struct {
	TransferOffersResponse
}

// PrivateStoreCollectionsToSubscriptionsMappingResponse contains the response from method PrivateStore.CollectionsToSubscriptionsMapping.
type PrivateStoreCollectionsToSubscriptionsMappingResponse struct {
	PrivateStoreCollectionsToSubscriptionsMappingResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateStoreCollectionsToSubscriptionsMappingResult contains the result from method PrivateStore.CollectionsToSubscriptionsMapping.
type PrivateStoreCollectionsToSubscriptionsMappingResult struct {
	CollectionsToSubscriptionsMappingResponse
}

// PrivateStoreCreateApprovalRequestResponse contains the response from method PrivateStore.CreateApprovalRequest.
type PrivateStoreCreateApprovalRequestResponse struct {
	PrivateStoreCreateApprovalRequestResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateStoreCreateApprovalRequestResult contains the result from method PrivateStore.CreateApprovalRequest.
type PrivateStoreCreateApprovalRequestResult struct {
	RequestApprovalResource
}

// PrivateStoreCreateOrUpdateResponse contains the response from method PrivateStore.CreateOrUpdate.
type PrivateStoreCreateOrUpdateResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateStoreDeleteResponse contains the response from method PrivateStore.Delete.
type PrivateStoreDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateStoreGetAdminRequestApprovalResponse contains the response from method PrivateStore.GetAdminRequestApproval.
type PrivateStoreGetAdminRequestApprovalResponse struct {
	PrivateStoreGetAdminRequestApprovalResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateStoreGetAdminRequestApprovalResult contains the result from method PrivateStore.GetAdminRequestApproval.
type PrivateStoreGetAdminRequestApprovalResult struct {
	AdminRequestApprovalsResource
}

// PrivateStoreGetApprovalRequestsListResponse contains the response from method PrivateStore.GetApprovalRequestsList.
type PrivateStoreGetApprovalRequestsListResponse struct {
	PrivateStoreGetApprovalRequestsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateStoreGetApprovalRequestsListResult contains the result from method PrivateStore.GetApprovalRequestsList.
type PrivateStoreGetApprovalRequestsListResult struct {
	RequestApprovalsList
}

// PrivateStoreGetRequestApprovalResponse contains the response from method PrivateStore.GetRequestApproval.
type PrivateStoreGetRequestApprovalResponse struct {
	PrivateStoreGetRequestApprovalResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateStoreGetRequestApprovalResult contains the result from method PrivateStore.GetRequestApproval.
type PrivateStoreGetRequestApprovalResult struct {
	RequestApprovalResource
}

// PrivateStoreGetResponse contains the response from method PrivateStore.Get.
type PrivateStoreGetResponse struct {
	PrivateStoreGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateStoreGetResult contains the result from method PrivateStore.Get.
type PrivateStoreGetResult struct {
	PrivateStore
}

// PrivateStoreListResponse contains the response from method PrivateStore.List.
type PrivateStoreListResponse struct {
	PrivateStoreListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateStoreListResult contains the result from method PrivateStore.List.
type PrivateStoreListResult struct {
	PrivateStoreList
}

// PrivateStoreQueryApprovedPlansResponse contains the response from method PrivateStore.QueryApprovedPlans.
type PrivateStoreQueryApprovedPlansResponse struct {
	PrivateStoreQueryApprovedPlansResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateStoreQueryApprovedPlansResult contains the result from method PrivateStore.QueryApprovedPlans.
type PrivateStoreQueryApprovedPlansResult struct {
	QueryApprovedPlansResponse
}

// PrivateStoreQueryNotificationsStateResponse contains the response from method PrivateStore.QueryNotificationsState.
type PrivateStoreQueryNotificationsStateResponse struct {
	PrivateStoreQueryNotificationsStateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateStoreQueryNotificationsStateResult contains the result from method PrivateStore.QueryNotificationsState.
type PrivateStoreQueryNotificationsStateResult struct {
	PrivateStoreNotificationsState
}

// PrivateStoreQueryOffersResponse contains the response from method PrivateStore.QueryOffers.
type PrivateStoreQueryOffersResponse struct {
	PrivateStoreQueryOffersResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateStoreQueryOffersResult contains the result from method PrivateStore.QueryOffers.
type PrivateStoreQueryOffersResult struct {
	QueryOffers
}

// PrivateStoreQueryRequestApprovalResponse contains the response from method PrivateStore.QueryRequestApproval.
type PrivateStoreQueryRequestApprovalResponse struct {
	PrivateStoreQueryRequestApprovalResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateStoreQueryRequestApprovalResult contains the result from method PrivateStore.QueryRequestApproval.
type PrivateStoreQueryRequestApprovalResult struct {
	QueryRequestApproval
}

// PrivateStoreUpdateAdminRequestApprovalResponse contains the response from method PrivateStore.UpdateAdminRequestApproval.
type PrivateStoreUpdateAdminRequestApprovalResponse struct {
	PrivateStoreUpdateAdminRequestApprovalResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateStoreUpdateAdminRequestApprovalResult contains the result from method PrivateStore.UpdateAdminRequestApproval.
type PrivateStoreUpdateAdminRequestApprovalResult struct {
	AdminRequestApprovalsResource
}

// PrivateStoreWithdrawPlanResponse contains the response from method PrivateStore.WithdrawPlan.
type PrivateStoreWithdrawPlanResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}
