//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbilling

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// PartnerTransfersClient contains the methods for the PartnerTransfers group.
// Don't use this type directly, use NewPartnerTransfersClient() instead.
type PartnerTransfersClient struct {
	internal *arm.Client
}

// NewPartnerTransfersClient creates a new instance of PartnerTransfersClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPartnerTransfersClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*PartnerTransfersClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PartnerTransfersClient{
		internal: cl,
	}
	return client, nil
}

// Cancel - Cancels a transfer request. The operation is supported only for billing accounts with agreement type Microsoft
// Partner Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - billingProfileName - The ID that uniquely identifies a billing profile.
//   - customerName - The ID that uniquely identifies a customer.
//   - transferName - The ID that uniquely identifies a transfer request.
//   - options - PartnerTransfersClientCancelOptions contains the optional parameters for the PartnerTransfersClient.Cancel method.
func (client *PartnerTransfersClient) Cancel(ctx context.Context, billingAccountName string, billingProfileName string, customerName string, transferName string, options *PartnerTransfersClientCancelOptions) (PartnerTransfersClientCancelResponse, error) {
	var err error
	const operationName = "PartnerTransfersClient.Cancel"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.cancelCreateRequest(ctx, billingAccountName, billingProfileName, customerName, transferName, options)
	if err != nil {
		return PartnerTransfersClientCancelResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PartnerTransfersClientCancelResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PartnerTransfersClientCancelResponse{}, err
	}
	resp, err := client.cancelHandleResponse(httpResp)
	return resp, err
}

// cancelCreateRequest creates the Cancel request.
func (client *PartnerTransfersClient) cancelCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, customerName string, transferName string, options *PartnerTransfersClientCancelOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/customers/{customerName}/transfers/{transferName}/cancel"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	if customerName == "" {
		return nil, errors.New("parameter customerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customerName}", url.PathEscape(customerName))
	if transferName == "" {
		return nil, errors.New("parameter transferName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{transferName}", url.PathEscape(transferName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// cancelHandleResponse handles the Cancel response.
func (client *PartnerTransfersClient) cancelHandleResponse(resp *http.Response) (PartnerTransfersClientCancelResponse, error) {
	result := PartnerTransfersClientCancelResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PartnerTransferDetails); err != nil {
		return PartnerTransfersClientCancelResponse{}, err
	}
	return result, nil
}

// Get - Gets a transfer request by ID. The operation is supported only for billing accounts with agreement type Microsoft
// Partner Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - billingProfileName - The ID that uniquely identifies a billing profile.
//   - customerName - The ID that uniquely identifies a customer.
//   - transferName - The ID that uniquely identifies a transfer request.
//   - options - PartnerTransfersClientGetOptions contains the optional parameters for the PartnerTransfersClient.Get method.
func (client *PartnerTransfersClient) Get(ctx context.Context, billingAccountName string, billingProfileName string, customerName string, transferName string, options *PartnerTransfersClientGetOptions) (PartnerTransfersClientGetResponse, error) {
	var err error
	const operationName = "PartnerTransfersClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, billingAccountName, billingProfileName, customerName, transferName, options)
	if err != nil {
		return PartnerTransfersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PartnerTransfersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PartnerTransfersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PartnerTransfersClient) getCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, customerName string, transferName string, options *PartnerTransfersClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/customers/{customerName}/transfers/{transferName}"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	if customerName == "" {
		return nil, errors.New("parameter customerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customerName}", url.PathEscape(customerName))
	if transferName == "" {
		return nil, errors.New("parameter transferName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{transferName}", url.PathEscape(transferName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PartnerTransfersClient) getHandleResponse(resp *http.Response) (PartnerTransfersClientGetResponse, error) {
	result := PartnerTransfersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PartnerTransferDetails); err != nil {
		return PartnerTransfersClientGetResponse{}, err
	}
	return result, nil
}

// Initiate - Sends a request to a user in a customer's billing account to transfer billing ownership of their subscriptions.
// The operation is supported only for billing accounts with agreement type Microsoft
// Partner Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - billingProfileName - The ID that uniquely identifies a billing profile.
//   - customerName - The ID that uniquely identifies a customer.
//   - transferName - The ID that uniquely identifies a transfer request.
//   - parameters - Request parameters that are provided to the initiate transfer operation.
//   - options - PartnerTransfersClientInitiateOptions contains the optional parameters for the PartnerTransfersClient.Initiate
//     method.
func (client *PartnerTransfersClient) Initiate(ctx context.Context, billingAccountName string, billingProfileName string, customerName string, transferName string, parameters PartnerInitiateTransferRequest, options *PartnerTransfersClientInitiateOptions) (PartnerTransfersClientInitiateResponse, error) {
	var err error
	const operationName = "PartnerTransfersClient.Initiate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.initiateCreateRequest(ctx, billingAccountName, billingProfileName, customerName, transferName, parameters, options)
	if err != nil {
		return PartnerTransfersClientInitiateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PartnerTransfersClientInitiateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return PartnerTransfersClientInitiateResponse{}, err
	}
	resp, err := client.initiateHandleResponse(httpResp)
	return resp, err
}

// initiateCreateRequest creates the Initiate request.
func (client *PartnerTransfersClient) initiateCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, customerName string, transferName string, parameters PartnerInitiateTransferRequest, options *PartnerTransfersClientInitiateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/customers/{customerName}/transfers/{transferName}"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	if customerName == "" {
		return nil, errors.New("parameter customerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customerName}", url.PathEscape(customerName))
	if transferName == "" {
		return nil, errors.New("parameter transferName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{transferName}", url.PathEscape(transferName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// initiateHandleResponse handles the Initiate response.
func (client *PartnerTransfersClient) initiateHandleResponse(resp *http.Response) (PartnerTransfersClientInitiateResponse, error) {
	result := PartnerTransfersClientInitiateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PartnerTransferDetails); err != nil {
		return PartnerTransfersClientInitiateResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists the transfer requests sent to a customer. The operation is supported only for billing accounts with
// agreement type Microsoft Partner Agreement.
//
// Generated from API version 2024-04-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - billingProfileName - The ID that uniquely identifies a billing profile.
//   - customerName - The ID that uniquely identifies a customer.
//   - options - PartnerTransfersClientListOptions contains the optional parameters for the PartnerTransfersClient.NewListPager
//     method.
func (client *PartnerTransfersClient) NewListPager(billingAccountName string, billingProfileName string, customerName string, options *PartnerTransfersClientListOptions) *runtime.Pager[PartnerTransfersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[PartnerTransfersClientListResponse]{
		More: func(page PartnerTransfersClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PartnerTransfersClientListResponse) (PartnerTransfersClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PartnerTransfersClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, billingAccountName, billingProfileName, customerName, options)
			}, nil)
			if err != nil {
				return PartnerTransfersClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *PartnerTransfersClient) listCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, customerName string, options *PartnerTransfersClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/customers/{customerName}/transfers"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	if customerName == "" {
		return nil, errors.New("parameter customerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customerName}", url.PathEscape(customerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PartnerTransfersClient) listHandleResponse(resp *http.Response) (PartnerTransfersClientListResponse, error) {
	result := PartnerTransfersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PartnerTransferDetailsListResult); err != nil {
		return PartnerTransfersClientListResponse{}, err
	}
	return result, nil
}