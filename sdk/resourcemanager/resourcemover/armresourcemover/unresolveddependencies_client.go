//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcemover

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

// UnresolvedDependenciesClient contains the methods for the UnresolvedDependencies group.
// Don't use this type directly, use NewUnresolvedDependenciesClient() instead.
type UnresolvedDependenciesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewUnresolvedDependenciesClient creates a new instance of UnresolvedDependenciesClient with the specified values.
//   - subscriptionID - The Subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewUnresolvedDependenciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*UnresolvedDependenciesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &UnresolvedDependenciesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewGetPager - Gets a list of unresolved dependencies.
//
// Generated from API version 2023-08-01
//   - resourceGroupName - The Resource Group Name.
//   - moveCollectionName - The Move Collection Name.
//   - options - UnresolvedDependenciesClientGetOptions contains the optional parameters for the UnresolvedDependenciesClient.NewGetPager
//     method.
func (client *UnresolvedDependenciesClient) NewGetPager(resourceGroupName string, moveCollectionName string, options *UnresolvedDependenciesClientGetOptions) *runtime.Pager[UnresolvedDependenciesClientGetResponse] {
	return runtime.NewPager(runtime.PagingHandler[UnresolvedDependenciesClientGetResponse]{
		More: func(page UnresolvedDependenciesClientGetResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *UnresolvedDependenciesClientGetResponse) (UnresolvedDependenciesClientGetResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "UnresolvedDependenciesClient.NewGetPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.getCreateRequest(ctx, resourceGroupName, moveCollectionName, options)
			}, nil)
			if err != nil {
				return UnresolvedDependenciesClientGetResponse{}, err
			}
			return client.getHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// getCreateRequest creates the Get request.
func (client *UnresolvedDependenciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, moveCollectionName string, options *UnresolvedDependenciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/moveCollections/{moveCollectionName}/unresolvedDependencies"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if moveCollectionName == "" {
		return nil, errors.New("parameter moveCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{moveCollectionName}", url.PathEscape(moveCollectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.DependencyLevel != nil {
		reqQP.Set("dependencyLevel", string(*options.DependencyLevel))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	reqQP.Set("api-version", "2023-08-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *UnresolvedDependenciesClient) getHandleResponse(resp *http.Response) (UnresolvedDependenciesClientGetResponse, error) {
	result := UnresolvedDependenciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UnresolvedDependencyCollection); err != nil {
		return UnresolvedDependenciesClientGetResponse{}, err
	}
	return result, nil
}
