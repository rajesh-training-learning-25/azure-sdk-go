// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// HubRouteTablesOperations contains the methods for the HubRouteTables group.
type HubRouteTablesOperations interface {
	// BeginCreateOrUpdate - Creates a RouteTable resource if it doesn't exist else updates the existing RouteTable.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, routeTableParameters HubRouteTable, options *HubRouteTablesCreateOrUpdateOptions) (*HubRouteTablePollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (HubRouteTablePoller, error)
	// BeginDelete - Deletes a RouteTable.
	BeginDelete(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, options *HubRouteTablesDeleteOptions) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Retrieves the details of a RouteTable.
	Get(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, options *HubRouteTablesGetOptions) (*HubRouteTableResponse, error)
	// List - Retrieves the details of all RouteTables.
	List(resourceGroupName string, virtualHubName string, options *HubRouteTablesListOptions) ListHubRouteTablesResultPager
}

// HubRouteTablesClient implements the HubRouteTablesOperations interface.
// Don't use this type directly, use NewHubRouteTablesClient() instead.
type HubRouteTablesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewHubRouteTablesClient creates a new instance of HubRouteTablesClient with the specified values.
func NewHubRouteTablesClient(con *armcore.Connection, subscriptionID string) HubRouteTablesOperations {
	return &HubRouteTablesClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *HubRouteTablesClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

func (client *HubRouteTablesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, routeTableParameters HubRouteTable, options *HubRouteTablesCreateOrUpdateOptions) (*HubRouteTablePollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, virtualHubName, routeTableName, routeTableParameters, options)
	if err != nil {
		return nil, err
	}
	result := &HubRouteTablePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("HubRouteTablesClient.CreateOrUpdate", "azure-async-operation", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &hubRouteTablePoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*HubRouteTableResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *HubRouteTablesClient) ResumeCreateOrUpdate(token string) (HubRouteTablePoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("HubRouteTablesClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &hubRouteTablePoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates a RouteTable resource if it doesn't exist else updates the existing RouteTable.
func (client *HubRouteTablesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, routeTableParameters HubRouteTable, options *HubRouteTablesCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, virtualHubName, routeTableName, routeTableParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.CreateOrUpdateHandleError(resp)
	}
	return resp, nil
}

// CreateOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *HubRouteTablesClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, routeTableParameters HubRouteTable, options *HubRouteTablesCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubRouteTables/{routeTableName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	urlPath = strings.ReplaceAll(urlPath, "{routeTableName}", url.PathEscape(routeTableName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(routeTableParameters)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *HubRouteTablesClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*HubRouteTableResponse, error) {
	result := HubRouteTableResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.HubRouteTable)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *HubRouteTablesClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *HubRouteTablesClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, options *HubRouteTablesDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, virtualHubName, routeTableName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("HubRouteTablesClient.Delete", "location", resp, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *HubRouteTablesClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("HubRouteTablesClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes a RouteTable.
func (client *HubRouteTablesClient) Delete(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, options *HubRouteTablesDeleteOptions) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, virtualHubName, routeTableName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.DeleteHandleError(resp)
	}
	return resp, nil
}

// DeleteCreateRequest creates the Delete request.
func (client *HubRouteTablesClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, options *HubRouteTablesDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubRouteTables/{routeTableName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	urlPath = strings.ReplaceAll(urlPath, "{routeTableName}", url.PathEscape(routeTableName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DeleteHandleError handles the Delete error response.
func (client *HubRouteTablesClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Retrieves the details of a RouteTable.
func (client *HubRouteTablesClient) Get(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, options *HubRouteTablesGetOptions) (*HubRouteTableResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, virtualHubName, routeTableName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetHandleError(resp)
	}
	result, err := client.GetHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCreateRequest creates the Get request.
func (client *HubRouteTablesClient) GetCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, options *HubRouteTablesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubRouteTables/{routeTableName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	urlPath = strings.ReplaceAll(urlPath, "{routeTableName}", url.PathEscape(routeTableName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *HubRouteTablesClient) GetHandleResponse(resp *azcore.Response) (*HubRouteTableResponse, error) {
	result := HubRouteTableResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.HubRouteTable)
}

// GetHandleError handles the Get error response.
func (client *HubRouteTablesClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Retrieves the details of all RouteTables.
func (client *HubRouteTablesClient) List(resourceGroupName string, virtualHubName string, options *HubRouteTablesListOptions) ListHubRouteTablesResultPager {
	return &listHubRouteTablesResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, resourceGroupName, virtualHubName, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *ListHubRouteTablesResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ListHubRouteTablesResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListCreateRequest creates the List request.
func (client *HubRouteTablesClient) ListCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, options *HubRouteTablesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubRouteTables"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListHandleResponse handles the List response.
func (client *HubRouteTablesClient) ListHandleResponse(resp *azcore.Response) (*ListHubRouteTablesResultResponse, error) {
	result := ListHubRouteTablesResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ListHubRouteTablesResult)
}

// ListHandleError handles the List error response.
func (client *HubRouteTablesClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
