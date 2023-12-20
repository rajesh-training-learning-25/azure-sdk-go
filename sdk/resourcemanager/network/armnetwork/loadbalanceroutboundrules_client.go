//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

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

// LoadBalancerOutboundRulesClient contains the methods for the LoadBalancerOutboundRules group.
// Don't use this type directly, use NewLoadBalancerOutboundRulesClient() instead.
type LoadBalancerOutboundRulesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewLoadBalancerOutboundRulesClient creates a new instance of LoadBalancerOutboundRulesClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewLoadBalancerOutboundRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LoadBalancerOutboundRulesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &LoadBalancerOutboundRulesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets the specified load balancer outbound rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The name of the resource group.
//   - loadBalancerName - The name of the load balancer.
//   - outboundRuleName - The name of the outbound rule.
//   - options - LoadBalancerOutboundRulesClientGetOptions contains the optional parameters for the LoadBalancerOutboundRulesClient.Get
//     method.
func (client *LoadBalancerOutboundRulesClient) Get(ctx context.Context, resourceGroupName string, loadBalancerName string, outboundRuleName string, options *LoadBalancerOutboundRulesClientGetOptions) (LoadBalancerOutboundRulesClientGetResponse, error) {
	var err error
	const operationName = "LoadBalancerOutboundRulesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, loadBalancerName, outboundRuleName, options)
	if err != nil {
		return LoadBalancerOutboundRulesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LoadBalancerOutboundRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LoadBalancerOutboundRulesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *LoadBalancerOutboundRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, outboundRuleName string, options *LoadBalancerOutboundRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/outboundRules/{outboundRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if outboundRuleName == "" {
		return nil, errors.New("parameter outboundRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{outboundRuleName}", url.PathEscape(outboundRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LoadBalancerOutboundRulesClient) getHandleResponse(resp *http.Response) (LoadBalancerOutboundRulesClientGetResponse, error) {
	result := LoadBalancerOutboundRulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OutboundRule); err != nil {
		return LoadBalancerOutboundRulesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all the outbound rules in a load balancer.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The name of the resource group.
//   - loadBalancerName - The name of the load balancer.
//   - options - LoadBalancerOutboundRulesClientListOptions contains the optional parameters for the LoadBalancerOutboundRulesClient.NewListPager
//     method.
func (client *LoadBalancerOutboundRulesClient) NewListPager(resourceGroupName string, loadBalancerName string, options *LoadBalancerOutboundRulesClientListOptions) *runtime.Pager[LoadBalancerOutboundRulesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[LoadBalancerOutboundRulesClientListResponse]{
		More: func(page LoadBalancerOutboundRulesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LoadBalancerOutboundRulesClientListResponse) (LoadBalancerOutboundRulesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "LoadBalancerOutboundRulesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, loadBalancerName, options)
			}, nil)
			if err != nil {
				return LoadBalancerOutboundRulesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *LoadBalancerOutboundRulesClient) listCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancerOutboundRulesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/outboundRules"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LoadBalancerOutboundRulesClient) listHandleResponse(resp *http.Response) (LoadBalancerOutboundRulesClientListResponse, error) {
	result := LoadBalancerOutboundRulesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LoadBalancerOutboundRuleListResult); err != nil {
		return LoadBalancerOutboundRulesClientListResponse{}, err
	}
	return result, nil
}
