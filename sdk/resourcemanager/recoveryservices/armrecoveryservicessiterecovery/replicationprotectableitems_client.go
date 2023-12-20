//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicessiterecovery

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

// ReplicationProtectableItemsClient contains the methods for the ReplicationProtectableItems group.
// Don't use this type directly, use NewReplicationProtectableItemsClient() instead.
type ReplicationProtectableItemsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewReplicationProtectableItemsClient creates a new instance of ReplicationProtectableItemsClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewReplicationProtectableItemsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ReplicationProtectableItemsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ReplicationProtectableItemsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - The operation to get the details of a protectable item.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Fabric name.
//   - protectionContainerName - Protection container name.
//   - protectableItemName - Protectable item name.
//   - options - ReplicationProtectableItemsClientGetOptions contains the optional parameters for the ReplicationProtectableItemsClient.Get
//     method.
func (client *ReplicationProtectableItemsClient) Get(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, protectionContainerName string, protectableItemName string, options *ReplicationProtectableItemsClientGetOptions) (ReplicationProtectableItemsClientGetResponse, error) {
	var err error
	const operationName = "ReplicationProtectableItemsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceName, resourceGroupName, fabricName, protectionContainerName, protectableItemName, options)
	if err != nil {
		return ReplicationProtectableItemsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ReplicationProtectableItemsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ReplicationProtectableItemsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ReplicationProtectableItemsClient) getCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, protectionContainerName string, protectableItemName string, options *ReplicationProtectableItemsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectableItems/{protectableItemName}"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if protectionContainerName == "" {
		return nil, errors.New("parameter protectionContainerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{protectionContainerName}", url.PathEscape(protectionContainerName))
	if protectableItemName == "" {
		return nil, errors.New("parameter protectableItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{protectableItemName}", url.PathEscape(protectableItemName))
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
func (client *ReplicationProtectableItemsClient) getHandleResponse(resp *http.Response) (ReplicationProtectableItemsClientGetResponse, error) {
	result := ReplicationProtectableItemsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProtectableItem); err != nil {
		return ReplicationProtectableItemsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByReplicationProtectionContainersPager - Lists the protectable items in a protection container.
//
// Generated from API version 2023-06-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Fabric name.
//   - protectionContainerName - Protection container name.
//   - options - ReplicationProtectableItemsClientListByReplicationProtectionContainersOptions contains the optional parameters
//     for the ReplicationProtectableItemsClient.NewListByReplicationProtectionContainersPager method.
func (client *ReplicationProtectableItemsClient) NewListByReplicationProtectionContainersPager(resourceName string, resourceGroupName string, fabricName string, protectionContainerName string, options *ReplicationProtectableItemsClientListByReplicationProtectionContainersOptions) *runtime.Pager[ReplicationProtectableItemsClientListByReplicationProtectionContainersResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReplicationProtectableItemsClientListByReplicationProtectionContainersResponse]{
		More: func(page ReplicationProtectableItemsClientListByReplicationProtectionContainersResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReplicationProtectableItemsClientListByReplicationProtectionContainersResponse) (ReplicationProtectableItemsClientListByReplicationProtectionContainersResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ReplicationProtectableItemsClient.NewListByReplicationProtectionContainersPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByReplicationProtectionContainersCreateRequest(ctx, resourceName, resourceGroupName, fabricName, protectionContainerName, options)
			}, nil)
			if err != nil {
				return ReplicationProtectableItemsClientListByReplicationProtectionContainersResponse{}, err
			}
			return client.listByReplicationProtectionContainersHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByReplicationProtectionContainersCreateRequest creates the ListByReplicationProtectionContainers request.
func (client *ReplicationProtectableItemsClient) listByReplicationProtectionContainersCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, protectionContainerName string, options *ReplicationProtectableItemsClientListByReplicationProtectionContainersOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectableItems"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if protectionContainerName == "" {
		return nil, errors.New("parameter protectionContainerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{protectionContainerName}", url.PathEscape(protectionContainerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Take != nil {
		reqQP.Set("$take", *options.Take)
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByReplicationProtectionContainersHandleResponse handles the ListByReplicationProtectionContainers response.
func (client *ReplicationProtectableItemsClient) listByReplicationProtectionContainersHandleResponse(resp *http.Response) (ReplicationProtectableItemsClientListByReplicationProtectionContainersResponse, error) {
	result := ReplicationProtectableItemsClientListByReplicationProtectionContainersResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProtectableItemCollection); err != nil {
		return ReplicationProtectableItemsClientListByReplicationProtectionContainersResponse{}, err
	}
	return result, nil
}
