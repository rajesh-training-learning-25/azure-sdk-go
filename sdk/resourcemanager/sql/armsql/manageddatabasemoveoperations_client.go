//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ManagedDatabaseMoveOperationsClient contains the methods for the ManagedDatabaseMoveOperations group.
// Don't use this type directly, use NewManagedDatabaseMoveOperationsClient() instead.
type ManagedDatabaseMoveOperationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewManagedDatabaseMoveOperationsClient creates a new instance of ManagedDatabaseMoveOperationsClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagedDatabaseMoveOperationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagedDatabaseMoveOperationsClient, error) {
	cl, err := arm.NewClient(moduleName+".ManagedDatabaseMoveOperationsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ManagedDatabaseMoveOperationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets a managed database move operation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-05-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - options - ManagedDatabaseMoveOperationsClientGetOptions contains the optional parameters for the ManagedDatabaseMoveOperationsClient.Get
//     method.
func (client *ManagedDatabaseMoveOperationsClient) Get(ctx context.Context, resourceGroupName string, locationName string, operationID string, options *ManagedDatabaseMoveOperationsClientGetOptions) (ManagedDatabaseMoveOperationsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, locationName, operationID, options)
	if err != nil {
		return ManagedDatabaseMoveOperationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedDatabaseMoveOperationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ManagedDatabaseMoveOperationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ManagedDatabaseMoveOperationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, locationName string, operationID string, options *ManagedDatabaseMoveOperationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/locations/{locationName}/managedDatabaseMoveOperationResults/{operationId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagedDatabaseMoveOperationsClient) getHandleResponse(resp *http.Response) (ManagedDatabaseMoveOperationsClientGetResponse, error) {
	result := ManagedDatabaseMoveOperationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedDatabaseMoveOperationResult); err != nil {
		return ManagedDatabaseMoveOperationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByLocationPager - Lists managed database move operations.
//
// Generated from API version 2022-05-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - options - ManagedDatabaseMoveOperationsClientListByLocationOptions contains the optional parameters for the ManagedDatabaseMoveOperationsClient.NewListByLocationPager
//     method.
func (client *ManagedDatabaseMoveOperationsClient) NewListByLocationPager(resourceGroupName string, locationName string, options *ManagedDatabaseMoveOperationsClientListByLocationOptions) *runtime.Pager[ManagedDatabaseMoveOperationsClientListByLocationResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedDatabaseMoveOperationsClientListByLocationResponse]{
		More: func(page ManagedDatabaseMoveOperationsClientListByLocationResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedDatabaseMoveOperationsClientListByLocationResponse) (ManagedDatabaseMoveOperationsClientListByLocationResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByLocationCreateRequest(ctx, resourceGroupName, locationName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ManagedDatabaseMoveOperationsClientListByLocationResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ManagedDatabaseMoveOperationsClientListByLocationResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ManagedDatabaseMoveOperationsClientListByLocationResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByLocationHandleResponse(resp)
		},
	})
}

// listByLocationCreateRequest creates the ListByLocation request.
func (client *ManagedDatabaseMoveOperationsClient) listByLocationCreateRequest(ctx context.Context, resourceGroupName string, locationName string, options *ManagedDatabaseMoveOperationsClientListByLocationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/locations/{locationName}/managedDatabaseMoveOperationResults"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.OnlyLatestPerDatabase != nil {
		reqQP.Set("onlyLatestPerDatabase", strconv.FormatBool(*options.OnlyLatestPerDatabase))
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2022-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByLocationHandleResponse handles the ListByLocation response.
func (client *ManagedDatabaseMoveOperationsClient) listByLocationHandleResponse(resp *http.Response) (ManagedDatabaseMoveOperationsClientListByLocationResponse, error) {
	result := ManagedDatabaseMoveOperationsClientListByLocationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedDatabaseMoveOperationListResult); err != nil {
		return ManagedDatabaseMoveOperationsClientListByLocationResponse{}, err
	}
	return result, nil
}
