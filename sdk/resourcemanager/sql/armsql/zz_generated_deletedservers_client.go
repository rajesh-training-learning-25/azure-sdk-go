//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DeletedServersClient contains the methods for the DeletedServers group.
// Don't use this type directly, use NewDeletedServersClient() instead.
type DeletedServersClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDeletedServersClient creates a new instance of DeletedServersClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDeletedServersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *DeletedServersClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &DeletedServersClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Get - Gets a deleted server.
// If the operation fails it returns an *azcore.ResponseError type.
// locationName - The name of the region where the resource is located.
// deletedServerName - The name of the deleted server.
// options - DeletedServersClientGetOptions contains the optional parameters for the DeletedServersClient.Get method.
func (client *DeletedServersClient) Get(ctx context.Context, locationName string, deletedServerName string, options *DeletedServersClientGetOptions) (DeletedServersClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, locationName, deletedServerName, options)
	if err != nil {
		return DeletedServersClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DeletedServersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DeletedServersClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DeletedServersClient) getCreateRequest(ctx context.Context, locationName string, deletedServerName string, options *DeletedServersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationName}/deletedServers/{deletedServerName}"
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if deletedServerName == "" {
		return nil, errors.New("parameter deletedServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deletedServerName}", url.PathEscape(deletedServerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DeletedServersClient) getHandleResponse(resp *http.Response) (DeletedServersClientGetResponse, error) {
	result := DeletedServersClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeletedServer); err != nil {
		return DeletedServersClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets a list of all deleted servers in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - DeletedServersClientListOptions contains the optional parameters for the DeletedServersClient.List method.
func (client *DeletedServersClient) List(options *DeletedServersClientListOptions) *DeletedServersClientListPager {
	return &DeletedServersClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp DeletedServersClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DeletedServerListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *DeletedServersClient) listCreateRequest(ctx context.Context, options *DeletedServersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Sql/deletedServers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DeletedServersClient) listHandleResponse(resp *http.Response) (DeletedServersClientListResponse, error) {
	result := DeletedServersClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeletedServerListResult); err != nil {
		return DeletedServersClientListResponse{}, err
	}
	return result, nil
}

// ListByLocation - Gets a list of deleted servers for a location.
// If the operation fails it returns an *azcore.ResponseError type.
// locationName - The name of the region where the resource is located.
// options - DeletedServersClientListByLocationOptions contains the optional parameters for the DeletedServersClient.ListByLocation
// method.
func (client *DeletedServersClient) ListByLocation(locationName string, options *DeletedServersClientListByLocationOptions) *DeletedServersClientListByLocationPager {
	return &DeletedServersClientListByLocationPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByLocationCreateRequest(ctx, locationName, options)
		},
		advancer: func(ctx context.Context, resp DeletedServersClientListByLocationResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DeletedServerListResult.NextLink)
		},
	}
}

// listByLocationCreateRequest creates the ListByLocation request.
func (client *DeletedServersClient) listByLocationCreateRequest(ctx context.Context, locationName string, options *DeletedServersClientListByLocationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationName}/deletedServers"
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByLocationHandleResponse handles the ListByLocation response.
func (client *DeletedServersClient) listByLocationHandleResponse(resp *http.Response) (DeletedServersClientListByLocationResponse, error) {
	result := DeletedServersClientListByLocationResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeletedServerListResult); err != nil {
		return DeletedServersClientListByLocationResponse{}, err
	}
	return result, nil
}

// BeginRecover - Recovers a deleted server.
// If the operation fails it returns an *azcore.ResponseError type.
// locationName - The name of the region where the resource is located.
// deletedServerName - The name of the deleted server.
// options - DeletedServersClientBeginRecoverOptions contains the optional parameters for the DeletedServersClient.BeginRecover
// method.
func (client *DeletedServersClient) BeginRecover(ctx context.Context, locationName string, deletedServerName string, options *DeletedServersClientBeginRecoverOptions) (DeletedServersClientRecoverPollerResponse, error) {
	resp, err := client.recoverOperation(ctx, locationName, deletedServerName, options)
	if err != nil {
		return DeletedServersClientRecoverPollerResponse{}, err
	}
	result := DeletedServersClientRecoverPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DeletedServersClient.Recover", "", resp, client.pl)
	if err != nil {
		return DeletedServersClientRecoverPollerResponse{}, err
	}
	result.Poller = &DeletedServersClientRecoverPoller{
		pt: pt,
	}
	return result, nil
}

// Recover - Recovers a deleted server.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DeletedServersClient) recoverOperation(ctx context.Context, locationName string, deletedServerName string, options *DeletedServersClientBeginRecoverOptions) (*http.Response, error) {
	req, err := client.recoverCreateRequest(ctx, locationName, deletedServerName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// recoverCreateRequest creates the Recover request.
func (client *DeletedServersClient) recoverCreateRequest(ctx context.Context, locationName string, deletedServerName string, options *DeletedServersClientBeginRecoverOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationName}/deletedServers/{deletedServerName}/recover"
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if deletedServerName == "" {
		return nil, errors.New("parameter deletedServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deletedServerName}", url.PathEscape(deletedServerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}
