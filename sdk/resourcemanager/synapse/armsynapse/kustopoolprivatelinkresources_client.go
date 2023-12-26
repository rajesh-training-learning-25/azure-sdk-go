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

// KustoPoolPrivateLinkResourcesClient contains the methods for the KustoPoolPrivateLinkResources group.
// Don't use this type directly, use NewKustoPoolPrivateLinkResourcesClient() instead.
type KustoPoolPrivateLinkResourcesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewKustoPoolPrivateLinkResourcesClient creates a new instance of KustoPoolPrivateLinkResourcesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewKustoPoolPrivateLinkResourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*KustoPoolPrivateLinkResourcesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &KustoPoolPrivateLinkResourcesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Lists all Kusto pool PrivateLinkResources.
//
// Generated from API version 2021-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - kustoPoolName - The name of the Kusto pool.
//   - options - KustoPoolPrivateLinkResourcesClientListOptions contains the optional parameters for the KustoPoolPrivateLinkResourcesClient.NewListPager
//     method.
func (client *KustoPoolPrivateLinkResourcesClient) NewListPager(resourceGroupName string, workspaceName string, kustoPoolName string, options *KustoPoolPrivateLinkResourcesClientListOptions) *runtime.Pager[KustoPoolPrivateLinkResourcesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[KustoPoolPrivateLinkResourcesClientListResponse]{
		More: func(page KustoPoolPrivateLinkResourcesClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *KustoPoolPrivateLinkResourcesClientListResponse) (KustoPoolPrivateLinkResourcesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "KustoPoolPrivateLinkResourcesClient.NewListPager")
			req, err := client.listCreateRequest(ctx, resourceGroupName, workspaceName, kustoPoolName, options)
			if err != nil {
				return KustoPoolPrivateLinkResourcesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return KustoPoolPrivateLinkResourcesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return KustoPoolPrivateLinkResourcesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *KustoPoolPrivateLinkResourcesClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, options *KustoPoolPrivateLinkResourcesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/kustoPools/{kustoPoolName}/privateLinkResources"
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
	if kustoPoolName == "" {
		return nil, errors.New("parameter kustoPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kustoPoolName}", url.PathEscape(kustoPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *KustoPoolPrivateLinkResourcesClient) listHandleResponse(resp *http.Response) (KustoPoolPrivateLinkResourcesClientListResponse, error) {
	result := KustoPoolPrivateLinkResourcesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResources); err != nil {
		return KustoPoolPrivateLinkResourcesClientListResponse{}, err
	}
	return result, nil
}
