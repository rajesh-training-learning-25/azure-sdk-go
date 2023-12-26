//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetworkanalytics

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

// DataProductsCatalogsClient contains the methods for the DataProductsCatalogs group.
// Don't use this type directly, use NewDataProductsCatalogsClient() instead.
type DataProductsCatalogsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDataProductsCatalogsClient creates a new instance of DataProductsCatalogsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDataProductsCatalogsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DataProductsCatalogsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DataProductsCatalogsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Retrieve data type resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - DataProductsCatalogsClientGetOptions contains the optional parameters for the DataProductsCatalogsClient.Get
//     method.
func (client *DataProductsCatalogsClient) Get(ctx context.Context, resourceGroupName string, options *DataProductsCatalogsClientGetOptions) (DataProductsCatalogsClientGetResponse, error) {
	var err error
	const operationName = "DataProductsCatalogsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, options)
	if err != nil {
		return DataProductsCatalogsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DataProductsCatalogsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DataProductsCatalogsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DataProductsCatalogsClient) getCreateRequest(ctx context.Context, resourceGroupName string, options *DataProductsCatalogsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkAnalytics/dataProductsCatalogs/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DataProductsCatalogsClient) getHandleResponse(resp *http.Response) (DataProductsCatalogsClientGetResponse, error) {
	result := DataProductsCatalogsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataProductsCatalog); err != nil {
		return DataProductsCatalogsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List data catalog by resource group.
//
// Generated from API version 2023-11-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - DataProductsCatalogsClientListByResourceGroupOptions contains the optional parameters for the DataProductsCatalogsClient.NewListByResourceGroupPager
//     method.
func (client *DataProductsCatalogsClient) NewListByResourceGroupPager(resourceGroupName string, options *DataProductsCatalogsClientListByResourceGroupOptions) *runtime.Pager[DataProductsCatalogsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[DataProductsCatalogsClientListByResourceGroupResponse]{
		More: func(page DataProductsCatalogsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DataProductsCatalogsClientListByResourceGroupResponse) (DataProductsCatalogsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DataProductsCatalogsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return DataProductsCatalogsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DataProductsCatalogsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DataProductsCatalogsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkAnalytics/dataProductsCatalogs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DataProductsCatalogsClient) listByResourceGroupHandleResponse(resp *http.Response) (DataProductsCatalogsClientListByResourceGroupResponse, error) {
	result := DataProductsCatalogsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataProductsCatalogListResult); err != nil {
		return DataProductsCatalogsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List data catalog by subscription.
//
// Generated from API version 2023-11-15
//   - options - DataProductsCatalogsClientListBySubscriptionOptions contains the optional parameters for the DataProductsCatalogsClient.NewListBySubscriptionPager
//     method.
func (client *DataProductsCatalogsClient) NewListBySubscriptionPager(options *DataProductsCatalogsClientListBySubscriptionOptions) *runtime.Pager[DataProductsCatalogsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[DataProductsCatalogsClientListBySubscriptionResponse]{
		More: func(page DataProductsCatalogsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DataProductsCatalogsClientListBySubscriptionResponse) (DataProductsCatalogsClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DataProductsCatalogsClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return DataProductsCatalogsClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *DataProductsCatalogsClient) listBySubscriptionCreateRequest(ctx context.Context, options *DataProductsCatalogsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.NetworkAnalytics/dataProductsCatalogs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *DataProductsCatalogsClient) listBySubscriptionHandleResponse(resp *http.Response) (DataProductsCatalogsClientListBySubscriptionResponse, error) {
	result := DataProductsCatalogsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataProductsCatalogListResult); err != nil {
		return DataProductsCatalogsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}
