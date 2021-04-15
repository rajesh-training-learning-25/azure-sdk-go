// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

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

// RouteFiltersClient contains the methods for the RouteFilters group.
// Don't use this type directly, use NewRouteFiltersClient() instead.
type RouteFiltersClient struct {
	con *armcore.Connection
	subscriptionID string
}

// NewRouteFiltersClient creates a new instance of RouteFiltersClient with the specified values.
func NewRouteFiltersClient(con *armcore.Connection, subscriptionID string) *RouteFiltersClient {
	return &RouteFiltersClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a route filter in a specified resource group.
func (client *RouteFiltersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, routeFilterName string, routeFilterParameters RouteFilter, options *RouteFiltersBeginCreateOrUpdateOptions) (RouteFilterPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, routeFilterName, routeFilterParameters, options)
	if err != nil {
		return RouteFilterPollerResponse{}, err
	}
	result := RouteFilterPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("RouteFiltersClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return RouteFilterPollerResponse{}, err
	}
	poller := &routeFilterPoller{
		pt: pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (RouteFilterResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new RouteFilterPoller from the specified resume token.
// token - The value must come from a previous call to RouteFilterPoller.ResumeToken().
func (client *RouteFiltersClient) ResumeCreateOrUpdate(token string) (RouteFilterPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("RouteFiltersClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &routeFilterPoller{
		pipeline: client.con.Pipeline(),
		pt: pt,
	}, nil
}

// CreateOrUpdate - Creates or updates a route filter in a specified resource group.
func (client *RouteFiltersClient) createOrUpdate(ctx context.Context, resourceGroupName string, routeFilterName string, routeFilterParameters RouteFilter, options *RouteFiltersBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, routeFilterName, routeFilterParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	 return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RouteFiltersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, routeFilterName string, routeFilterParameters RouteFilter, options *RouteFiltersBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if routeFilterName == "" {
		return nil, errors.New("parameter routeFilterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeFilterName}", url.PathEscape(routeFilterName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-07-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(routeFilterParameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *RouteFiltersClient) createOrUpdateHandleResponse(resp *azcore.Response) (RouteFilterResponse, error) {
	var val *RouteFilter
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RouteFilterResponse{}, err
	}
return RouteFilterResponse{RawResponse: resp.Response, RouteFilter: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *RouteFiltersClient) createOrUpdateHandleError(resp *azcore.Response) error {
var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Deletes the specified route filter.
func (client *RouteFiltersClient) BeginDelete(ctx context.Context, resourceGroupName string, routeFilterName string, options *RouteFiltersBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, routeFilterName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("RouteFiltersClient.Delete", "location", resp, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *RouteFiltersClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("RouteFiltersClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt: pt,
	}, nil
}

// Delete - Deletes the specified route filter.
func (client *RouteFiltersClient) deleteOperation(ctx context.Context, resourceGroupName string, routeFilterName string, options *RouteFiltersBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, routeFilterName, options)
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
func (client *RouteFiltersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, routeFilterName string, options *RouteFiltersBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if routeFilterName == "" {
		return nil, errors.New("parameter routeFilterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeFilterName}", url.PathEscape(routeFilterName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-07-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *RouteFiltersClient) deleteHandleError(resp *azcore.Response) error {
var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified route filter.
func (client *RouteFiltersClient) Get(ctx context.Context, resourceGroupName string, routeFilterName string, options *RouteFiltersGetOptions) (RouteFilterResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, routeFilterName, options)
	if err != nil {
		return RouteFilterResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return RouteFilterResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return RouteFilterResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RouteFiltersClient) getCreateRequest(ctx context.Context, resourceGroupName string, routeFilterName string, options *RouteFiltersGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if routeFilterName == "" {
		return nil, errors.New("parameter routeFilterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeFilterName}", url.PathEscape(routeFilterName))
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
	reqQP.Set("api-version", "2020-07-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RouteFiltersClient) getHandleResponse(resp *azcore.Response) (RouteFilterResponse, error) {
	var val *RouteFilter
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RouteFilterResponse{}, err
	}
return RouteFilterResponse{RawResponse: resp.Response, RouteFilter: val}, nil
}

// getHandleError handles the Get error response.
func (client *RouteFiltersClient) getHandleError(resp *azcore.Response) error {
var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets all route filters in a subscription.
func (client *RouteFiltersClient) List(options *RouteFiltersListOptions) (RouteFilterListResultPager) {
	return &routeFilterListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp RouteFilterListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.RouteFilterListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *RouteFiltersClient) listCreateRequest(ctx context.Context, options *RouteFiltersListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/routeFilters"
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
	reqQP.Set("api-version", "2020-07-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RouteFiltersClient) listHandleResponse(resp *azcore.Response) (RouteFilterListResultResponse, error) {
	var val *RouteFilterListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RouteFilterListResultResponse{}, err
	}
return RouteFilterListResultResponse{RawResponse: resp.Response, RouteFilterListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *RouteFiltersClient) listHandleError(resp *azcore.Response) error {
var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListByResourceGroup - Gets all route filters in a resource group.
func (client *RouteFiltersClient) ListByResourceGroup(resourceGroupName string, options *RouteFiltersListByResourceGroupOptions) (RouteFilterListResultPager) {
	return &routeFilterListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listByResourceGroupHandleResponse,
		errorer:   client.listByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp RouteFilterListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.RouteFilterListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *RouteFiltersClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *RouteFiltersListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters"
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
	reqQP.Set("api-version", "2020-07-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *RouteFiltersClient) listByResourceGroupHandleResponse(resp *azcore.Response) (RouteFilterListResultResponse, error) {
	var val *RouteFilterListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RouteFilterListResultResponse{}, err
	}
return RouteFilterListResultResponse{RawResponse: resp.Response, RouteFilterListResult: val}, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *RouteFiltersClient) listByResourceGroupHandleError(resp *azcore.Response) error {
var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// UpdateTags - Updates tags of a route filter.
func (client *RouteFiltersClient) UpdateTags(ctx context.Context, resourceGroupName string, routeFilterName string, parameters TagsObject, options *RouteFiltersUpdateTagsOptions) (RouteFilterResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, routeFilterName, parameters, options)
	if err != nil {
		return RouteFilterResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return RouteFilterResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return RouteFilterResponse{}, client.updateTagsHandleError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *RouteFiltersClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, routeFilterName string, parameters TagsObject, options *RouteFiltersUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if routeFilterName == "" {
		return nil, errors.New("parameter routeFilterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeFilterName}", url.PathEscape(routeFilterName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-07-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *RouteFiltersClient) updateTagsHandleResponse(resp *azcore.Response) (RouteFilterResponse, error) {
	var val *RouteFilter
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RouteFilterResponse{}, err
	}
return RouteFilterResponse{RawResponse: resp.Response, RouteFilter: val}, nil
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *RouteFiltersClient) updateTagsHandleError(resp *azcore.Response) error {
var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

