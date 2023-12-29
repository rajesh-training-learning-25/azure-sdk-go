//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvmwarecloudsimple

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

// CustomizationPoliciesClient contains the methods for the CustomizationPolicies group.
// Don't use this type directly, use NewCustomizationPoliciesClient() instead.
type CustomizationPoliciesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCustomizationPoliciesClient creates a new instance of CustomizationPoliciesClient with the specified values.
//   - subscriptionID - The subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCustomizationPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CustomizationPoliciesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CustomizationPoliciesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Returns customization policy by its name
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-04-01
//   - regionID - The region Id (westus, eastus)
//   - pcName - The private cloud name
//   - customizationPolicyName - customization policy name
//   - options - CustomizationPoliciesClientGetOptions contains the optional parameters for the CustomizationPoliciesClient.Get
//     method.
func (client *CustomizationPoliciesClient) Get(ctx context.Context, regionID string, pcName string, customizationPolicyName string, options *CustomizationPoliciesClientGetOptions) (CustomizationPoliciesClientGetResponse, error) {
	var err error
	const operationName = "CustomizationPoliciesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, regionID, pcName, customizationPolicyName, options)
	if err != nil {
		return CustomizationPoliciesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CustomizationPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CustomizationPoliciesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *CustomizationPoliciesClient) getCreateRequest(ctx context.Context, regionID string, pcName string, customizationPolicyName string, options *CustomizationPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/customizationPolicies/{customizationPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if regionID == "" {
		return nil, errors.New("parameter regionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{regionId}", url.PathEscape(regionID))
	if pcName == "" {
		return nil, errors.New("parameter pcName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pcName}", url.PathEscape(pcName))
	if customizationPolicyName == "" {
		return nil, errors.New("parameter customizationPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customizationPolicyName}", url.PathEscape(customizationPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CustomizationPoliciesClient) getHandleResponse(resp *http.Response) (CustomizationPoliciesClientGetResponse, error) {
	result := CustomizationPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomizationPolicy); err != nil {
		return CustomizationPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Returns list of customization policies in region for private cloud
//
// Generated from API version 2019-04-01
//   - regionID - The region Id (westus, eastus)
//   - pcName - The private cloud name
//   - options - CustomizationPoliciesClientListOptions contains the optional parameters for the CustomizationPoliciesClient.NewListPager
//     method.
func (client *CustomizationPoliciesClient) NewListPager(regionID string, pcName string, options *CustomizationPoliciesClientListOptions) *runtime.Pager[CustomizationPoliciesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[CustomizationPoliciesClientListResponse]{
		More: func(page CustomizationPoliciesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CustomizationPoliciesClientListResponse) (CustomizationPoliciesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CustomizationPoliciesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, regionID, pcName, options)
			}, nil)
			if err != nil {
				return CustomizationPoliciesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *CustomizationPoliciesClient) listCreateRequest(ctx context.Context, regionID string, pcName string, options *CustomizationPoliciesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/customizationPolicies"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if regionID == "" {
		return nil, errors.New("parameter regionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{regionId}", url.PathEscape(regionID))
	if pcName == "" {
		return nil, errors.New("parameter pcName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pcName}", url.PathEscape(pcName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-04-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *CustomizationPoliciesClient) listHandleResponse(resp *http.Response) (CustomizationPoliciesClientListResponse, error) {
	result := CustomizationPoliciesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomizationPoliciesListResponse); err != nil {
		return CustomizationPoliciesClientListResponse{}, err
	}
	return result, nil
}
