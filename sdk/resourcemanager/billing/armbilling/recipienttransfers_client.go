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

// RecipientTransfersClient contains the methods for the RecipientTransfers group.
// Don't use this type directly, use NewRecipientTransfersClient() instead.
type RecipientTransfersClient struct {
	internal *arm.Client
}

// NewRecipientTransfersClient creates a new instance of RecipientTransfersClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRecipientTransfersClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*RecipientTransfersClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RecipientTransfersClient{
		internal: cl,
	}
	return client, nil
}

// Accept - Accepts a transfer request.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - transferName - The ID that uniquely identifies a transfer request.
//   - parameters - Request parameters that are provided to the accept transfer operation.
//   - options - RecipientTransfersClientAcceptOptions contains the optional parameters for the RecipientTransfersClient.Accept
//     method.
func (client *RecipientTransfersClient) Accept(ctx context.Context, transferName string, parameters AcceptTransferRequest, options *RecipientTransfersClientAcceptOptions) (RecipientTransfersClientAcceptResponse, error) {
	var err error
	const operationName = "RecipientTransfersClient.Accept"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.acceptCreateRequest(ctx, transferName, parameters, options)
	if err != nil {
		return RecipientTransfersClientAcceptResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RecipientTransfersClientAcceptResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RecipientTransfersClientAcceptResponse{}, err
	}
	resp, err := client.acceptHandleResponse(httpResp)
	return resp, err
}

// acceptCreateRequest creates the Accept request.
func (client *RecipientTransfersClient) acceptCreateRequest(ctx context.Context, transferName string, parameters AcceptTransferRequest, options *RecipientTransfersClientAcceptOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/transfers/{transferName}/accept"
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
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// acceptHandleResponse handles the Accept response.
func (client *RecipientTransfersClient) acceptHandleResponse(resp *http.Response) (RecipientTransfersClientAcceptResponse, error) {
	result := RecipientTransfersClientAcceptResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecipientTransferDetails); err != nil {
		return RecipientTransfersClientAcceptResponse{}, err
	}
	return result, nil
}

// Decline - Declines a transfer request.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - transferName - The ID that uniquely identifies a transfer request.
//   - options - RecipientTransfersClientDeclineOptions contains the optional parameters for the RecipientTransfersClient.Decline
//     method.
func (client *RecipientTransfersClient) Decline(ctx context.Context, transferName string, options *RecipientTransfersClientDeclineOptions) (RecipientTransfersClientDeclineResponse, error) {
	var err error
	const operationName = "RecipientTransfersClient.Decline"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.declineCreateRequest(ctx, transferName, options)
	if err != nil {
		return RecipientTransfersClientDeclineResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RecipientTransfersClientDeclineResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RecipientTransfersClientDeclineResponse{}, err
	}
	resp, err := client.declineHandleResponse(httpResp)
	return resp, err
}

// declineCreateRequest creates the Decline request.
func (client *RecipientTransfersClient) declineCreateRequest(ctx context.Context, transferName string, options *RecipientTransfersClientDeclineOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/transfers/{transferName}/decline"
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

// declineHandleResponse handles the Decline response.
func (client *RecipientTransfersClient) declineHandleResponse(resp *http.Response) (RecipientTransfersClientDeclineResponse, error) {
	result := RecipientTransfersClientDeclineResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecipientTransferDetails); err != nil {
		return RecipientTransfersClientDeclineResponse{}, err
	}
	return result, nil
}

// Get - Gets a transfer request by ID. The caller must be the recipient of the transfer request.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - transferName - The ID that uniquely identifies a transfer request.
//   - options - RecipientTransfersClientGetOptions contains the optional parameters for the RecipientTransfersClient.Get method.
func (client *RecipientTransfersClient) Get(ctx context.Context, transferName string, options *RecipientTransfersClientGetOptions) (RecipientTransfersClientGetResponse, error) {
	var err error
	const operationName = "RecipientTransfersClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, transferName, options)
	if err != nil {
		return RecipientTransfersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RecipientTransfersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RecipientTransfersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *RecipientTransfersClient) getCreateRequest(ctx context.Context, transferName string, options *RecipientTransfersClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/transfers/{transferName}"
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
func (client *RecipientTransfersClient) getHandleResponse(resp *http.Response) (RecipientTransfersClientGetResponse, error) {
	result := RecipientTransfersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecipientTransferDetails); err != nil {
		return RecipientTransfersClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists the transfer requests received by the caller.
//
// Generated from API version 2024-04-01
//   - options - RecipientTransfersClientListOptions contains the optional parameters for the RecipientTransfersClient.NewListPager
//     method.
func (client *RecipientTransfersClient) NewListPager(options *RecipientTransfersClientListOptions) *runtime.Pager[RecipientTransfersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[RecipientTransfersClientListResponse]{
		More: func(page RecipientTransfersClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RecipientTransfersClientListResponse) (RecipientTransfersClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RecipientTransfersClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return RecipientTransfersClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *RecipientTransfersClient) listCreateRequest(ctx context.Context, options *RecipientTransfersClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/transfers"
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
func (client *RecipientTransfersClient) listHandleResponse(resp *http.Response) (RecipientTransfersClientListResponse, error) {
	result := RecipientTransfersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecipientTransferDetailsListResult); err != nil {
		return RecipientTransfersClientListResponse{}, err
	}
	return result, nil
}

// Validate - Validates if a subscription or a reservation can be transferred. Use this operation to validate your subscriptions
// or reservation before using the accept transfer operation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - transferName - The ID that uniquely identifies a transfer request.
//   - parameters - Request parameters that are provided to the validate transfer operation.
//   - options - RecipientTransfersClientValidateOptions contains the optional parameters for the RecipientTransfersClient.Validate
//     method.
func (client *RecipientTransfersClient) Validate(ctx context.Context, transferName string, parameters AcceptTransferRequest, options *RecipientTransfersClientValidateOptions) (RecipientTransfersClientValidateResponse, error) {
	var err error
	const operationName = "RecipientTransfersClient.Validate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.validateCreateRequest(ctx, transferName, parameters, options)
	if err != nil {
		return RecipientTransfersClientValidateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RecipientTransfersClientValidateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RecipientTransfersClientValidateResponse{}, err
	}
	resp, err := client.validateHandleResponse(httpResp)
	return resp, err
}

// validateCreateRequest creates the Validate request.
func (client *RecipientTransfersClient) validateCreateRequest(ctx context.Context, transferName string, parameters AcceptTransferRequest, options *RecipientTransfersClientValidateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/transfers/{transferName}/validate"
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
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// validateHandleResponse handles the Validate response.
func (client *RecipientTransfersClient) validateHandleResponse(resp *http.Response) (RecipientTransfersClientValidateResponse, error) {
	result := RecipientTransfersClientValidateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ValidateTransferListResponse); err != nil {
		return RecipientTransfersClientValidateResponse{}, err
	}
	return result, nil
}