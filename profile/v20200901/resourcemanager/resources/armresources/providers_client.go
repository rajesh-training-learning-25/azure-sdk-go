//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armresources

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/profile/v20200901/internal"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ProvidersClient contains the methods for the Providers group.
// Don't use this type directly, use NewProvidersClient() instead.
type ProvidersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewProvidersClient creates a new instance of ProvidersClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewProvidersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ProvidersClient, error) {
	cl, err := arm.NewClient(internal.ModuleName+"/armresources.ProvidersClient", internal.ModuleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ProvidersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets the specified resource provider.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-10-01
//   - resourceProviderNamespace - The namespace of the resource provider.
//   - options - ProvidersClientGetOptions contains the optional parameters for the ProvidersClient.Get method.
func (client *ProvidersClient) Get(ctx context.Context, resourceProviderNamespace string, options *ProvidersClientGetOptions) (ProvidersClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceProviderNamespace, options)
	if err != nil {
		return ProvidersClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProvidersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProvidersClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ProvidersClient) getCreateRequest(ctx context.Context, resourceProviderNamespace string, options *ProvidersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/{resourceProviderNamespace}"
	if resourceProviderNamespace == "" {
		return nil, errors.New("parameter resourceProviderNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderNamespace}", url.PathEscape(resourceProviderNamespace))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProvidersClient) getHandleResponse(resp *http.Response) (ProvidersClientGetResponse, error) {
	result := ProvidersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Provider); err != nil {
		return ProvidersClientGetResponse{}, err
	}
	return result, nil
}

// GetAtTenantScope - Gets the specified resource provider at the tenant level.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-10-01
//   - resourceProviderNamespace - The namespace of the resource provider.
//   - options - ProvidersClientGetAtTenantScopeOptions contains the optional parameters for the ProvidersClient.GetAtTenantScope
//     method.
func (client *ProvidersClient) GetAtTenantScope(ctx context.Context, resourceProviderNamespace string, options *ProvidersClientGetAtTenantScopeOptions) (ProvidersClientGetAtTenantScopeResponse, error) {
	req, err := client.getAtTenantScopeCreateRequest(ctx, resourceProviderNamespace, options)
	if err != nil {
		return ProvidersClientGetAtTenantScopeResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProvidersClientGetAtTenantScopeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProvidersClientGetAtTenantScopeResponse{}, runtime.NewResponseError(resp)
	}
	return client.getAtTenantScopeHandleResponse(resp)
}

// getAtTenantScopeCreateRequest creates the GetAtTenantScope request.
func (client *ProvidersClient) getAtTenantScopeCreateRequest(ctx context.Context, resourceProviderNamespace string, options *ProvidersClientGetAtTenantScopeOptions) (*policy.Request, error) {
	urlPath := "/providers/{resourceProviderNamespace}"
	if resourceProviderNamespace == "" {
		return nil, errors.New("parameter resourceProviderNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderNamespace}", url.PathEscape(resourceProviderNamespace))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAtTenantScopeHandleResponse handles the GetAtTenantScope response.
func (client *ProvidersClient) getAtTenantScopeHandleResponse(resp *http.Response) (ProvidersClientGetAtTenantScopeResponse, error) {
	result := ProvidersClientGetAtTenantScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Provider); err != nil {
		return ProvidersClientGetAtTenantScopeResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all resource providers for a subscription.
//
// Generated from API version 2019-10-01
//   - options - ProvidersClientListOptions contains the optional parameters for the ProvidersClient.NewListPager method.
func (client *ProvidersClient) NewListPager(options *ProvidersClientListOptions) *runtime.Pager[ProvidersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProvidersClientListResponse]{
		More: func(page ProvidersClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ProvidersClientListResponse) (ProvidersClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ProvidersClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ProvidersClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ProvidersClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ProvidersClient) listCreateRequest(ctx context.Context, options *ProvidersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ProvidersClient) listHandleResponse(resp *http.Response) (ProvidersClientListResponse, error) {
	result := ProvidersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProviderListResult); err != nil {
		return ProvidersClientListResponse{}, err
	}
	return result, nil
}

// NewListAtTenantScopePager - Gets all resource providers for the tenant.
//
// Generated from API version 2019-10-01
//   - options - ProvidersClientListAtTenantScopeOptions contains the optional parameters for the ProvidersClient.NewListAtTenantScopePager
//     method.
func (client *ProvidersClient) NewListAtTenantScopePager(options *ProvidersClientListAtTenantScopeOptions) *runtime.Pager[ProvidersClientListAtTenantScopeResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProvidersClientListAtTenantScopeResponse]{
		More: func(page ProvidersClientListAtTenantScopeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ProvidersClientListAtTenantScopeResponse) (ProvidersClientListAtTenantScopeResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAtTenantScopeCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ProvidersClientListAtTenantScopeResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ProvidersClientListAtTenantScopeResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ProvidersClientListAtTenantScopeResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAtTenantScopeHandleResponse(resp)
		},
	})
}

// listAtTenantScopeCreateRequest creates the ListAtTenantScope request.
func (client *ProvidersClient) listAtTenantScopeCreateRequest(ctx context.Context, options *ProvidersClientListAtTenantScopeOptions) (*policy.Request, error) {
	urlPath := "/providers"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAtTenantScopeHandleResponse handles the ListAtTenantScope response.
func (client *ProvidersClient) listAtTenantScopeHandleResponse(resp *http.Response) (ProvidersClientListAtTenantScopeResponse, error) {
	result := ProvidersClientListAtTenantScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProviderListResult); err != nil {
		return ProvidersClientListAtTenantScopeResponse{}, err
	}
	return result, nil
}

// Register - Registers a subscription with a resource provider.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-10-01
//   - resourceProviderNamespace - The namespace of the resource provider to register.
//   - options - ProvidersClientRegisterOptions contains the optional parameters for the ProvidersClient.Register method.
func (client *ProvidersClient) Register(ctx context.Context, resourceProviderNamespace string, options *ProvidersClientRegisterOptions) (ProvidersClientRegisterResponse, error) {
	req, err := client.registerCreateRequest(ctx, resourceProviderNamespace, options)
	if err != nil {
		return ProvidersClientRegisterResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProvidersClientRegisterResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProvidersClientRegisterResponse{}, runtime.NewResponseError(resp)
	}
	return client.registerHandleResponse(resp)
}

// registerCreateRequest creates the Register request.
func (client *ProvidersClient) registerCreateRequest(ctx context.Context, resourceProviderNamespace string, options *ProvidersClientRegisterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/{resourceProviderNamespace}/register"
	if resourceProviderNamespace == "" {
		return nil, errors.New("parameter resourceProviderNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderNamespace}", url.PathEscape(resourceProviderNamespace))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// registerHandleResponse handles the Register response.
func (client *ProvidersClient) registerHandleResponse(resp *http.Response) (ProvidersClientRegisterResponse, error) {
	result := ProvidersClientRegisterResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Provider); err != nil {
		return ProvidersClientRegisterResponse{}, err
	}
	return result, nil
}

// Unregister - Unregisters a subscription from a resource provider.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-10-01
//   - resourceProviderNamespace - The namespace of the resource provider to unregister.
//   - options - ProvidersClientUnregisterOptions contains the optional parameters for the ProvidersClient.Unregister method.
func (client *ProvidersClient) Unregister(ctx context.Context, resourceProviderNamespace string, options *ProvidersClientUnregisterOptions) (ProvidersClientUnregisterResponse, error) {
	req, err := client.unregisterCreateRequest(ctx, resourceProviderNamespace, options)
	if err != nil {
		return ProvidersClientUnregisterResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProvidersClientUnregisterResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProvidersClientUnregisterResponse{}, runtime.NewResponseError(resp)
	}
	return client.unregisterHandleResponse(resp)
}

// unregisterCreateRequest creates the Unregister request.
func (client *ProvidersClient) unregisterCreateRequest(ctx context.Context, resourceProviderNamespace string, options *ProvidersClientUnregisterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/{resourceProviderNamespace}/unregister"
	if resourceProviderNamespace == "" {
		return nil, errors.New("parameter resourceProviderNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderNamespace}", url.PathEscape(resourceProviderNamespace))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// unregisterHandleResponse handles the Unregister response.
func (client *ProvidersClient) unregisterHandleResponse(resp *http.Response) (ProvidersClientUnregisterResponse, error) {
	result := ProvidersClientUnregisterResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Provider); err != nil {
		return ProvidersClientUnregisterResponse{}, err
	}
	return result, nil
}
