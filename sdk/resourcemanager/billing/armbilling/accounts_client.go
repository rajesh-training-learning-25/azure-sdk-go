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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AccountsClient contains the methods for the BillingAccounts group.
// Don't use this type directly, use NewAccountsClient() instead.
type AccountsClient struct {
	internal *arm.Client
}

// NewAccountsClient creates a new instance of AccountsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAccountsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*AccountsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AccountsClient{
		internal: cl,
	}
	return client, nil
}

// Get - Gets a billing account by its ID.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-05-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - options - AccountsClientGetOptions contains the optional parameters for the AccountsClient.Get method.
func (client *AccountsClient) Get(ctx context.Context, billingAccountName string, options *AccountsClientGetOptions) (AccountsClientGetResponse, error) {
	var err error
	const operationName = "AccountsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, billingAccountName, options)
	if err != nil {
		return AccountsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AccountsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AccountsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AccountsClient) getCreateRequest(ctx context.Context, billingAccountName string, options *AccountsClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AccountsClient) getHandleResponse(resp *http.Response) (AccountsClientGetResponse, error) {
	result := AccountsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Account); err != nil {
		return AccountsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists the billing accounts that a user has access to.
//
// Generated from API version 2020-05-01
//   - options - AccountsClientListOptions contains the optional parameters for the AccountsClient.NewListPager method.
func (client *AccountsClient) NewListPager(options *AccountsClientListOptions) *runtime.Pager[AccountsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AccountsClientListResponse]{
		More: func(page AccountsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AccountsClientListResponse) (AccountsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AccountsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return AccountsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *AccountsClient) listCreateRequest(ctx context.Context, options *AccountsClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AccountsClient) listHandleResponse(resp *http.Response) (AccountsClientListResponse, error) {
	result := AccountsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccountListResult); err != nil {
		return AccountsClientListResponse{}, err
	}
	return result, nil
}

// NewListInvoiceSectionsByCreateSubscriptionPermissionPager - Lists the invoice sections for which the user has permission
// to create Azure subscriptions. The operation is supported only for billing accounts with agreement type Microsoft Customer
// Agreement.
//
// Generated from API version 2020-05-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - options - AccountsClientListInvoiceSectionsByCreateSubscriptionPermissionOptions contains the optional parameters for the
//     AccountsClient.NewListInvoiceSectionsByCreateSubscriptionPermissionPager method.
func (client *AccountsClient) NewListInvoiceSectionsByCreateSubscriptionPermissionPager(billingAccountName string, options *AccountsClientListInvoiceSectionsByCreateSubscriptionPermissionOptions) *runtime.Pager[AccountsClientListInvoiceSectionsByCreateSubscriptionPermissionResponse] {
	return runtime.NewPager(runtime.PagingHandler[AccountsClientListInvoiceSectionsByCreateSubscriptionPermissionResponse]{
		More: func(page AccountsClientListInvoiceSectionsByCreateSubscriptionPermissionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AccountsClientListInvoiceSectionsByCreateSubscriptionPermissionResponse) (AccountsClientListInvoiceSectionsByCreateSubscriptionPermissionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AccountsClient.NewListInvoiceSectionsByCreateSubscriptionPermissionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listInvoiceSectionsByCreateSubscriptionPermissionCreateRequest(ctx, billingAccountName, options)
			}, nil)
			if err != nil {
				return AccountsClientListInvoiceSectionsByCreateSubscriptionPermissionResponse{}, err
			}
			return client.listInvoiceSectionsByCreateSubscriptionPermissionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listInvoiceSectionsByCreateSubscriptionPermissionCreateRequest creates the ListInvoiceSectionsByCreateSubscriptionPermission request.
func (client *AccountsClient) listInvoiceSectionsByCreateSubscriptionPermissionCreateRequest(ctx context.Context, billingAccountName string, options *AccountsClientListInvoiceSectionsByCreateSubscriptionPermissionOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/listInvoiceSectionsWithCreateSubscriptionPermission"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listInvoiceSectionsByCreateSubscriptionPermissionHandleResponse handles the ListInvoiceSectionsByCreateSubscriptionPermission response.
func (client *AccountsClient) listInvoiceSectionsByCreateSubscriptionPermissionHandleResponse(resp *http.Response) (AccountsClientListInvoiceSectionsByCreateSubscriptionPermissionResponse, error) {
	result := AccountsClientListInvoiceSectionsByCreateSubscriptionPermissionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InvoiceSectionListWithCreateSubPermissionResult); err != nil {
		return AccountsClientListInvoiceSectionsByCreateSubscriptionPermissionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates the properties of a billing account. Currently, displayName and address can be updated. The operation
// is supported only for billing accounts with agreement type Microsoft Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-05-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - parameters - Request parameters that are provided to the update billing account operation.
//   - options - AccountsClientBeginUpdateOptions contains the optional parameters for the AccountsClient.BeginUpdate method.
func (client *AccountsClient) BeginUpdate(ctx context.Context, billingAccountName string, parameters AccountUpdateRequest, options *AccountsClientBeginUpdateOptions) (*runtime.Poller[AccountsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, billingAccountName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AccountsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AccountsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Updates the properties of a billing account. Currently, displayName and address can be updated. The operation
// is supported only for billing accounts with agreement type Microsoft Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-05-01
func (client *AccountsClient) update(ctx context.Context, billingAccountName string, parameters AccountUpdateRequest, options *AccountsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "AccountsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, billingAccountName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *AccountsClient) updateCreateRequest(ctx context.Context, billingAccountName string, parameters AccountUpdateRequest, options *AccountsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}
