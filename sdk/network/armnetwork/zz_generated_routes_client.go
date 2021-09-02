//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// RoutesClient contains the methods for the Routes group.
// Don't use this type directly, use NewRoutesClient() instead.
type RoutesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewRoutesClient creates a new instance of RoutesClient with the specified values.
func NewRoutesClient(con *arm.Connection, subscriptionID string) *RoutesClient {
	return &RoutesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a route in the specified route table.
// If the operation fails it returns the *CloudError error type.
func (client *RoutesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, routeTableName string, routeName string, routeParameters Route, options *RoutesBeginCreateOrUpdateOptions) (RoutesCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, routeTableName, routeName, routeParameters, options)
	if err != nil {
		return RoutesCreateOrUpdatePollerResponse{}, err
	}
	result := RoutesCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("RoutesClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return RoutesCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &RoutesCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a route in the specified route table.
// If the operation fails it returns the *CloudError error type.
func (client *RoutesClient) createOrUpdate(ctx context.Context, resourceGroupName string, routeTableName string, routeName string, routeParameters Route, options *RoutesBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, routeTableName, routeName, routeParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RoutesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, routeTableName string, routeName string, routeParameters Route, options *RoutesBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}/routes/{routeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if routeTableName == "" {
		return nil, errors.New("parameter routeTableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeTableName}", url.PathEscape(routeTableName))
	if routeName == "" {
		return nil, errors.New("parameter routeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeName}", url.PathEscape(routeName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, routeParameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *RoutesClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes the specified route from a route table.
// If the operation fails it returns the *CloudError error type.
func (client *RoutesClient) BeginDelete(ctx context.Context, resourceGroupName string, routeTableName string, routeName string, options *RoutesBeginDeleteOptions) (RoutesDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, routeTableName, routeName, options)
	if err != nil {
		return RoutesDeletePollerResponse{}, err
	}
	result := RoutesDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("RoutesClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return RoutesDeletePollerResponse{}, err
	}
	result.Poller = &RoutesDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified route from a route table.
// If the operation fails it returns the *CloudError error type.
func (client *RoutesClient) deleteOperation(ctx context.Context, resourceGroupName string, routeTableName string, routeName string, options *RoutesBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, routeTableName, routeName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RoutesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, routeTableName string, routeName string, options *RoutesBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}/routes/{routeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if routeTableName == "" {
		return nil, errors.New("parameter routeTableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeTableName}", url.PathEscape(routeTableName))
	if routeName == "" {
		return nil, errors.New("parameter routeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeName}", url.PathEscape(routeName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *RoutesClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets the specified route from a route table.
// If the operation fails it returns the *CloudError error type.
func (client *RoutesClient) Get(ctx context.Context, resourceGroupName string, routeTableName string, routeName string, options *RoutesGetOptions) (RoutesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, routeTableName, routeName, options)
	if err != nil {
		return RoutesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoutesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoutesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RoutesClient) getCreateRequest(ctx context.Context, resourceGroupName string, routeTableName string, routeName string, options *RoutesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}/routes/{routeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if routeTableName == "" {
		return nil, errors.New("parameter routeTableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeTableName}", url.PathEscape(routeTableName))
	if routeName == "" {
		return nil, errors.New("parameter routeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeName}", url.PathEscape(routeName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RoutesClient) getHandleResponse(resp *http.Response) (RoutesGetResponse, error) {
	result := RoutesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Route); err != nil {
		return RoutesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *RoutesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Gets all routes in a route table.
// If the operation fails it returns the *CloudError error type.
func (client *RoutesClient) List(resourceGroupName string, routeTableName string, options *RoutesListOptions) *RoutesListPager {
	return &RoutesListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, routeTableName, options)
		},
		advancer: func(ctx context.Context, resp RoutesListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RouteListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *RoutesClient) listCreateRequest(ctx context.Context, resourceGroupName string, routeTableName string, options *RoutesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}/routes"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if routeTableName == "" {
		return nil, errors.New("parameter routeTableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeTableName}", url.PathEscape(routeTableName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RoutesClient) listHandleResponse(resp *http.Response) (RoutesListResponse, error) {
	result := RoutesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RouteListResult); err != nil {
		return RoutesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *RoutesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
