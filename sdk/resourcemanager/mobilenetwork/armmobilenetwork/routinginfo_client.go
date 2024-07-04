//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmobilenetwork

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

// RoutingInfoClient contains the methods for the RoutingInfo group.
// Don't use this type directly, use NewRoutingInfoClient() instead.
type RoutingInfoClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRoutingInfoClient creates a new instance of RoutingInfoClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRoutingInfoClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RoutingInfoClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RoutingInfoClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get the routing information for the packet core.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - packetCoreControlPlaneName - The name of the packet core control plane.
//   - options - RoutingInfoClientGetOptions contains the optional parameters for the RoutingInfoClient.Get method.
func (client *RoutingInfoClient) Get(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, options *RoutingInfoClientGetOptions) (RoutingInfoClientGetResponse, error) {
	var err error
	const operationName = "RoutingInfoClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, packetCoreControlPlaneName, options)
	if err != nil {
		return RoutingInfoClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RoutingInfoClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RoutingInfoClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *RoutingInfoClient) getCreateRequest(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, options *RoutingInfoClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/{packetCoreControlPlaneName}/routingInfo/default"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if packetCoreControlPlaneName == "" {
		return nil, errors.New("parameter packetCoreControlPlaneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCoreControlPlaneName}", url.PathEscape(packetCoreControlPlaneName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
func (client *RoutingInfoClient) getHandleResponse(resp *http.Response) (RoutingInfoClientGetResponse, error) {
	result := RoutingInfoClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoutingInfoModel); err != nil {
		return RoutingInfoClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List all of the routing information for the packet core.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - packetCoreControlPlaneName - The name of the packet core control plane.
//   - options - RoutingInfoClientListOptions contains the optional parameters for the RoutingInfoClient.NewListPager method.
func (client *RoutingInfoClient) NewListPager(resourceGroupName string, packetCoreControlPlaneName string, options *RoutingInfoClientListOptions) *runtime.Pager[RoutingInfoClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[RoutingInfoClientListResponse]{
		More: func(page RoutingInfoClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RoutingInfoClientListResponse) (RoutingInfoClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RoutingInfoClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, packetCoreControlPlaneName, options)
			}, nil)
			if err != nil {
				return RoutingInfoClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *RoutingInfoClient) listCreateRequest(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, options *RoutingInfoClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/{packetCoreControlPlaneName}/routingInfo"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if packetCoreControlPlaneName == "" {
		return nil, errors.New("parameter packetCoreControlPlaneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCoreControlPlaneName}", url.PathEscape(packetCoreControlPlaneName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
func (client *RoutingInfoClient) listHandleResponse(resp *http.Response) (RoutingInfoClientListResponse, error) {
	result := RoutingInfoClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoutingInfoListResult); err != nil {
		return RoutingInfoClientListResponse{}, err
	}
	return result, nil
}