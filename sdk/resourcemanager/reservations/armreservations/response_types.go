//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armreservations

import "encoding/json"

// AzureReservationAPIClientGetAppliedReservationListResponse contains the response from method AzureReservationAPIClient.GetAppliedReservationList.
type AzureReservationAPIClientGetAppliedReservationListResponse struct {
	AppliedReservations
}

// AzureReservationAPIClientGetCatalogResponse contains the response from method AzureReservationAPIClient.GetCatalog.
type AzureReservationAPIClientGetCatalogResponse struct {
	// Array of Catalog
	CatalogArray []*Catalog
}

// CalculateExchangeClientPostResponse contains the response from method CalculateExchangeClient.Post.
type CalculateExchangeClientPostResponse struct {
	CalculateExchangeOperationResultResponse
}

// CalculateRefundClientPostResponse contains the response from method CalculateRefundClient.Post.
type CalculateRefundClientPostResponse struct {
	CalculateRefundResponse
}

// ExchangeClientPostResponse contains the response from method ExchangeClient.Post.
type ExchangeClientPostResponse struct {
	ExchangeOperationResultResponse
}

// OperationClientListResponse contains the response from method OperationClient.List.
type OperationClientListResponse struct {
	OperationList
}

// QuotaClientCreateOrUpdateResponse contains the response from method QuotaClient.CreateOrUpdate.
type QuotaClientCreateOrUpdateResponse struct {
	CurrentQuotaLimitBase
}

// QuotaClientGetResponse contains the response from method QuotaClient.Get.
type QuotaClientGetResponse struct {
	CurrentQuotaLimitBase
	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// QuotaClientListResponse contains the response from method QuotaClient.List.
type QuotaClientListResponse struct {
	QuotaLimits
	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// QuotaClientUpdateResponse contains the response from method QuotaClient.Update.
type QuotaClientUpdateResponse struct {
	CurrentQuotaLimitBase
}

// QuotaRequestStatusClientGetResponse contains the response from method QuotaRequestStatusClient.Get.
type QuotaRequestStatusClientGetResponse struct {
	QuotaRequestDetails
}

// QuotaRequestStatusClientListResponse contains the response from method QuotaRequestStatusClient.List.
type QuotaRequestStatusClientListResponse struct {
	QuotaRequestDetailsList
}

// ReservationClientArchiveResponse contains the response from method ReservationClient.Archive.
type ReservationClientArchiveResponse struct {
	// placeholder for future response values
}

// ReservationClientAvailableScopesResponse contains the response from method ReservationClient.AvailableScopes.
type ReservationClientAvailableScopesResponse struct {
	AvailableScopeProperties
}

// ReservationClientGetResponse contains the response from method ReservationClient.Get.
type ReservationClientGetResponse struct {
	ReservationResponse
}

// ReservationClientListAllResponse contains the response from method ReservationClient.ListAll.
type ReservationClientListAllResponse struct {
	ListResult
}

// ReservationClientListResponse contains the response from method ReservationClient.List.
type ReservationClientListResponse struct {
	ReservationList
}

// ReservationClientListRevisionsResponse contains the response from method ReservationClient.ListRevisions.
type ReservationClientListRevisionsResponse struct {
	ReservationList
}

// ReservationClientMergeResponse contains the response from method ReservationClient.Merge.
type ReservationClientMergeResponse struct {
	// Array of ReservationResponse
	ReservationResponseArray []*ReservationResponse
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ReservationClientMergeResponse.
func (r *ReservationClientMergeResponse) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ReservationResponseArray)
}

// ReservationClientSplitResponse contains the response from method ReservationClient.Split.
type ReservationClientSplitResponse struct {
	// Array of ReservationResponse
	ReservationResponseArray []*ReservationResponse
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ReservationClientSplitResponse.
func (r *ReservationClientSplitResponse) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ReservationResponseArray)
}

// ReservationClientUnarchiveResponse contains the response from method ReservationClient.Unarchive.
type ReservationClientUnarchiveResponse struct {
	// placeholder for future response values
}

// ReservationClientUpdateResponse contains the response from method ReservationClient.Update.
type ReservationClientUpdateResponse struct {
	ReservationResponse
}

// ReservationOrderClientCalculateResponse contains the response from method ReservationOrderClient.Calculate.
type ReservationOrderClientCalculateResponse struct {
	CalculatePriceResponse
}

// ReservationOrderClientChangeDirectoryResponse contains the response from method ReservationOrderClient.ChangeDirectory.
type ReservationOrderClientChangeDirectoryResponse struct {
	ChangeDirectoryResponse
}

// ReservationOrderClientGetResponse contains the response from method ReservationOrderClient.Get.
type ReservationOrderClientGetResponse struct {
	ReservationOrderResponse
}

// ReservationOrderClientListResponse contains the response from method ReservationOrderClient.List.
type ReservationOrderClientListResponse struct {
	ReservationOrderList
}

// ReservationOrderClientPurchaseResponse contains the response from method ReservationOrderClient.Purchase.
type ReservationOrderClientPurchaseResponse struct {
	ReservationOrderResponse
}

// ReturnClientPostResponse contains the response from method ReturnClient.Post.
type ReturnClientPostResponse struct {
	RefundResponse
	// Location contains the information returned from the Location header response.
	Location *string
}
