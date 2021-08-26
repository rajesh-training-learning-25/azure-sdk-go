// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// PrivateLinkScopesClient contains the methods for the PrivateLinkScopes group.
// Don't use this type directly, use NewPrivateLinkScopesClient() instead.
type PrivateLinkScopesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewPrivateLinkScopesClient creates a new instance of PrivateLinkScopesClient with the specified values.
func NewPrivateLinkScopesClient(con *armcore.Connection, subscriptionID string) *PrivateLinkScopesClient {
	return &PrivateLinkScopesClient{con: con, subscriptionID: subscriptionID}
}

// CreateOrUpdate - Creates (or updates) a Azure Monitor PrivateLinkScope. Note: You cannot specify a different value for InstrumentationKey nor AppId in
// the Put operation.
// If the operation fails it returns a generic error.
func (client *PrivateLinkScopesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, scopeName string, azureMonitorPrivateLinkScopePayload AzureMonitorPrivateLinkScope, options *PrivateLinkScopesCreateOrUpdateOptions) (AzureMonitorPrivateLinkScopeResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, scopeName, azureMonitorPrivateLinkScopePayload, options)
	if err != nil {
		return AzureMonitorPrivateLinkScopeResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return AzureMonitorPrivateLinkScopeResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return AzureMonitorPrivateLinkScopeResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PrivateLinkScopesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, scopeName string, azureMonitorPrivateLinkScopePayload AzureMonitorPrivateLinkScope, options *PrivateLinkScopesCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/privateLinkScopes/{scopeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if scopeName == "" {
		return nil, errors.New("parameter scopeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scopeName}", url.PathEscape(scopeName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-10-17-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(azureMonitorPrivateLinkScopePayload)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *PrivateLinkScopesClient) createOrUpdateHandleResponse(resp *azcore.Response) (AzureMonitorPrivateLinkScopeResponse, error) {
	var val *AzureMonitorPrivateLinkScope
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return AzureMonitorPrivateLinkScopeResponse{}, err
	}
	return AzureMonitorPrivateLinkScopeResponse{RawResponse: resp.Response, AzureMonitorPrivateLinkScope: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *PrivateLinkScopesClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// BeginDelete - Deletes a Azure Monitor PrivateLinkScope.
// If the operation fails it returns a generic error.
func (client *PrivateLinkScopesClient) BeginDelete(ctx context.Context, resourceGroupName string, scopeName string, options *PrivateLinkScopesBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, scopeName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("PrivateLinkScopesClient.Delete", "", resp, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *PrivateLinkScopesClient) ResumeDelete(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("PrivateLinkScopesClient.Delete", token, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Delete - Deletes a Azure Monitor PrivateLinkScope.
// If the operation fails it returns a generic error.
func (client *PrivateLinkScopesClient) deleteOperation(ctx context.Context, resourceGroupName string, scopeName string, options *PrivateLinkScopesBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, scopeName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PrivateLinkScopesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, scopeName string, options *PrivateLinkScopesBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/privateLinkScopes/{scopeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if scopeName == "" {
		return nil, errors.New("parameter scopeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scopeName}", url.PathEscape(scopeName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-10-17-preview")
	req.URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *PrivateLinkScopesClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Get - Returns a Azure Monitor PrivateLinkScope.
// If the operation fails it returns a generic error.
func (client *PrivateLinkScopesClient) Get(ctx context.Context, resourceGroupName string, scopeName string, options *PrivateLinkScopesGetOptions) (AzureMonitorPrivateLinkScopeResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, scopeName, options)
	if err != nil {
		return AzureMonitorPrivateLinkScopeResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return AzureMonitorPrivateLinkScopeResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return AzureMonitorPrivateLinkScopeResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateLinkScopesClient) getCreateRequest(ctx context.Context, resourceGroupName string, scopeName string, options *PrivateLinkScopesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/privateLinkScopes/{scopeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if scopeName == "" {
		return nil, errors.New("parameter scopeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scopeName}", url.PathEscape(scopeName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-10-17-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateLinkScopesClient) getHandleResponse(resp *azcore.Response) (AzureMonitorPrivateLinkScopeResponse, error) {
	var val *AzureMonitorPrivateLinkScope
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return AzureMonitorPrivateLinkScopeResponse{}, err
	}
	return AzureMonitorPrivateLinkScopeResponse{RawResponse: resp.Response, AzureMonitorPrivateLinkScope: val}, nil
}

// getHandleError handles the Get error response.
func (client *PrivateLinkScopesClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// List - Gets a list of all Azure Monitor PrivateLinkScopes within a subscription.
// If the operation fails it returns a generic error.
func (client *PrivateLinkScopesClient) List(options *PrivateLinkScopesListOptions) AzureMonitorPrivateLinkScopeListResultPager {
	return &azureMonitorPrivateLinkScopeListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp AzureMonitorPrivateLinkScopeListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.AzureMonitorPrivateLinkScopeListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *PrivateLinkScopesClient) listCreateRequest(ctx context.Context, options *PrivateLinkScopesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/microsoft.insights/privateLinkScopes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-10-17-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PrivateLinkScopesClient) listHandleResponse(resp *azcore.Response) (AzureMonitorPrivateLinkScopeListResultResponse, error) {
	var val *AzureMonitorPrivateLinkScopeListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return AzureMonitorPrivateLinkScopeListResultResponse{}, err
	}
	return AzureMonitorPrivateLinkScopeListResultResponse{RawResponse: resp.Response, AzureMonitorPrivateLinkScopeListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *PrivateLinkScopesClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListByResourceGroup - Gets a list of Azure Monitor PrivateLinkScopes within a resource group.
// If the operation fails it returns a generic error.
func (client *PrivateLinkScopesClient) ListByResourceGroup(resourceGroupName string, options *PrivateLinkScopesListByResourceGroupOptions) AzureMonitorPrivateLinkScopeListResultPager {
	return &azureMonitorPrivateLinkScopeListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listByResourceGroupHandleResponse,
		errorer:   client.listByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp AzureMonitorPrivateLinkScopeListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.AzureMonitorPrivateLinkScopeListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *PrivateLinkScopesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *PrivateLinkScopesListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/privateLinkScopes"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-10-17-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *PrivateLinkScopesClient) listByResourceGroupHandleResponse(resp *azcore.Response) (AzureMonitorPrivateLinkScopeListResultResponse, error) {
	var val *AzureMonitorPrivateLinkScopeListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return AzureMonitorPrivateLinkScopeListResultResponse{}, err
	}
	return AzureMonitorPrivateLinkScopeListResultResponse{RawResponse: resp.Response, AzureMonitorPrivateLinkScopeListResult: val}, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *PrivateLinkScopesClient) listByResourceGroupHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// UpdateTags - Updates an existing PrivateLinkScope's tags. To update other fields use the CreateOrUpdate method.
// If the operation fails it returns a generic error.
func (client *PrivateLinkScopesClient) UpdateTags(ctx context.Context, resourceGroupName string, scopeName string, privateLinkScopeTags TagsResource, options *PrivateLinkScopesUpdateTagsOptions) (AzureMonitorPrivateLinkScopeResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, scopeName, privateLinkScopeTags, options)
	if err != nil {
		return AzureMonitorPrivateLinkScopeResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return AzureMonitorPrivateLinkScopeResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return AzureMonitorPrivateLinkScopeResponse{}, client.updateTagsHandleError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *PrivateLinkScopesClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, scopeName string, privateLinkScopeTags TagsResource, options *PrivateLinkScopesUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/privateLinkScopes/{scopeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if scopeName == "" {
		return nil, errors.New("parameter scopeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scopeName}", url.PathEscape(scopeName))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-10-17-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(privateLinkScopeTags)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *PrivateLinkScopesClient) updateTagsHandleResponse(resp *azcore.Response) (AzureMonitorPrivateLinkScopeResponse, error) {
	var val *AzureMonitorPrivateLinkScope
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return AzureMonitorPrivateLinkScopeResponse{}, err
	}
	return AzureMonitorPrivateLinkScopeResponse{RawResponse: resp.Response, AzureMonitorPrivateLinkScope: val}, nil
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *PrivateLinkScopesClient) updateTagsHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
