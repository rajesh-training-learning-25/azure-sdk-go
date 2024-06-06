//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevopsinfrastructure

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

// ResourceDetailsClient contains the methods for the ResourceDetails group.
// Don't use this type directly, use NewResourceDetailsClient() instead.
type ResourceDetailsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewResourceDetailsClient creates a new instance of ResourceDetailsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewResourceDetailsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ResourceDetailsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ResourceDetailsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListByPoolPager - List ResourceDetailsObject resources by Pool
//
// Generated from API version 2024-04-04-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - poolName - Name of the pool. It needs to be globally unique.
//   - options - ResourceDetailsClientListByPoolOptions contains the optional parameters for the ResourceDetailsClient.NewListByPoolPager
//     method.
func (client *ResourceDetailsClient) NewListByPoolPager(resourceGroupName string, poolName string, options *ResourceDetailsClientListByPoolOptions) *runtime.Pager[ResourceDetailsClientListByPoolResponse] {
	return runtime.NewPager(runtime.PagingHandler[ResourceDetailsClientListByPoolResponse]{
		More: func(page ResourceDetailsClientListByPoolResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ResourceDetailsClientListByPoolResponse) (ResourceDetailsClientListByPoolResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ResourceDetailsClient.NewListByPoolPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByPoolCreateRequest(ctx, resourceGroupName, poolName, options)
			}, nil)
			if err != nil {
				return ResourceDetailsClientListByPoolResponse{}, err
			}
			return client.listByPoolHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByPoolCreateRequest creates the ListByPool request.
func (client *ResourceDetailsClient) listByPoolCreateRequest(ctx context.Context, resourceGroupName string, poolName string, options *ResourceDetailsClientListByPoolOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevOpsInfrastructure/pools/{poolName}/resources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByPoolHandleResponse handles the ListByPool response.
func (client *ResourceDetailsClient) listByPoolHandleResponse(resp *http.Response) (ResourceDetailsClientListByPoolResponse, error) {
	result := ResourceDetailsClientListByPoolResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceDetailsObjectListResult); err != nil {
		return ResourceDetailsClientListByPoolResponse{}, err
	}
	return result, nil
}