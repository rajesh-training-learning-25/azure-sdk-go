//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse

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

// SQLPoolUsagesClient contains the methods for the SQLPoolUsages group.
// Don't use this type directly, use NewSQLPoolUsagesClient() instead.
type SQLPoolUsagesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSQLPoolUsagesClient creates a new instance of SQLPoolUsagesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSQLPoolUsagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SQLPoolUsagesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SQLPoolUsagesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Gets SQL pool usages.
//
// Generated from API version 2021-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - sqlPoolName - SQL pool name
//   - options - SQLPoolUsagesClientListOptions contains the optional parameters for the SQLPoolUsagesClient.NewListPager method.
func (client *SQLPoolUsagesClient) NewListPager(resourceGroupName string, workspaceName string, sqlPoolName string, options *SQLPoolUsagesClientListOptions) *runtime.Pager[SQLPoolUsagesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SQLPoolUsagesClientListResponse]{
		More: func(page SQLPoolUsagesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SQLPoolUsagesClientListResponse) (SQLPoolUsagesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SQLPoolUsagesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, options)
			}, nil)
			if err != nil {
				return SQLPoolUsagesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *SQLPoolUsagesClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, options *SQLPoolUsagesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/usages"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SQLPoolUsagesClient) listHandleResponse(resp *http.Response) (SQLPoolUsagesClientListResponse, error) {
	result := SQLPoolUsagesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLPoolUsageListResult); err != nil {
		return SQLPoolUsagesClientListResponse{}, err
	}
	return result, nil
}
