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

// ExpressRouteCrossConnectionPeeringsOperations contains the methods for the ExpressRouteCrossConnectionPeerings group.
type ExpressRouteCrossConnectionPeeringsOperations interface {
	// BeginCreateOrUpdate - Creates or updates a peering in the specified ExpressRouteCrossConnection.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, peeringParameters ExpressRouteCrossConnectionPeering, options *ExpressRouteCrossConnectionPeeringsCreateOrUpdateOptions) (*ExpressRouteCrossConnectionPeeringPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (ExpressRouteCrossConnectionPeeringPoller, error)
	// BeginDelete - Deletes the specified peering from the ExpressRouteCrossConnection.
	BeginDelete(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, options *ExpressRouteCrossConnectionPeeringsDeleteOptions) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets the specified peering for the ExpressRouteCrossConnection.
	Get(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, options *ExpressRouteCrossConnectionPeeringsGetOptions) (*ExpressRouteCrossConnectionPeeringResponse, error)
	// List - Gets all peerings in a specified ExpressRouteCrossConnection.
	List(resourceGroupName string, crossConnectionName string, options *ExpressRouteCrossConnectionPeeringsListOptions) ExpressRouteCrossConnectionPeeringListPager
}

// ExpressRouteCrossConnectionPeeringsClient implements the ExpressRouteCrossConnectionPeeringsOperations interface.
// Don't use this type directly, use NewExpressRouteCrossConnectionPeeringsClient() instead.
type ExpressRouteCrossConnectionPeeringsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewExpressRouteCrossConnectionPeeringsClient creates a new instance of ExpressRouteCrossConnectionPeeringsClient with the specified values.
func NewExpressRouteCrossConnectionPeeringsClient(con *armcore.Connection, subscriptionID string) ExpressRouteCrossConnectionPeeringsOperations {
	return &ExpressRouteCrossConnectionPeeringsClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *ExpressRouteCrossConnectionPeeringsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

func (client *ExpressRouteCrossConnectionPeeringsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, peeringParameters ExpressRouteCrossConnectionPeering, options *ExpressRouteCrossConnectionPeeringsCreateOrUpdateOptions) (*ExpressRouteCrossConnectionPeeringPollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, crossConnectionName, peeringName, peeringParameters, options)
	if err != nil {
		return nil, err
	}
	result := &ExpressRouteCrossConnectionPeeringPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ExpressRouteCrossConnectionPeeringsClient.CreateOrUpdate", "azure-async-operation", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &expressRouteCrossConnectionPeeringPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ExpressRouteCrossConnectionPeeringResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *ExpressRouteCrossConnectionPeeringsClient) ResumeCreateOrUpdate(token string) (ExpressRouteCrossConnectionPeeringPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRouteCrossConnectionPeeringsClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &expressRouteCrossConnectionPeeringPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates a peering in the specified ExpressRouteCrossConnection.
func (client *ExpressRouteCrossConnectionPeeringsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, peeringParameters ExpressRouteCrossConnectionPeering, options *ExpressRouteCrossConnectionPeeringsCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, crossConnectionName, peeringName, peeringParameters, options)
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
func (client *ExpressRouteCrossConnectionPeeringsClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, peeringParameters ExpressRouteCrossConnectionPeering, options *ExpressRouteCrossConnectionPeeringsCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(peeringParameters)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ExpressRouteCrossConnectionPeeringsClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*ExpressRouteCrossConnectionPeeringResponse, error) {
	result := ExpressRouteCrossConnectionPeeringResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCrossConnectionPeering)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ExpressRouteCrossConnectionPeeringsClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *ExpressRouteCrossConnectionPeeringsClient) BeginDelete(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, options *ExpressRouteCrossConnectionPeeringsDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, crossConnectionName, peeringName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ExpressRouteCrossConnectionPeeringsClient.Delete", "location", resp, client.DeleteHandleError)
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

func (client *ExpressRouteCrossConnectionPeeringsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRouteCrossConnectionPeeringsClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified peering from the ExpressRouteCrossConnection.
func (client *ExpressRouteCrossConnectionPeeringsClient) Delete(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, options *ExpressRouteCrossConnectionPeeringsDeleteOptions) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, crossConnectionName, peeringName, options)
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
func (client *ExpressRouteCrossConnectionPeeringsClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, options *ExpressRouteCrossConnectionPeeringsDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
func (client *ExpressRouteCrossConnectionPeeringsClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified peering for the ExpressRouteCrossConnection.
func (client *ExpressRouteCrossConnectionPeeringsClient) Get(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, options *ExpressRouteCrossConnectionPeeringsGetOptions) (*ExpressRouteCrossConnectionPeeringResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, crossConnectionName, peeringName, options)
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
func (client *ExpressRouteCrossConnectionPeeringsClient) GetCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, options *ExpressRouteCrossConnectionPeeringsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
func (client *ExpressRouteCrossConnectionPeeringsClient) GetHandleResponse(resp *azcore.Response) (*ExpressRouteCrossConnectionPeeringResponse, error) {
	result := ExpressRouteCrossConnectionPeeringResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCrossConnectionPeering)
}

// GetHandleError handles the Get error response.
func (client *ExpressRouteCrossConnectionPeeringsClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets all peerings in a specified ExpressRouteCrossConnection.
func (client *ExpressRouteCrossConnectionPeeringsClient) List(resourceGroupName string, crossConnectionName string, options *ExpressRouteCrossConnectionPeeringsListOptions) ExpressRouteCrossConnectionPeeringListPager {
	return &expressRouteCrossConnectionPeeringListPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, resourceGroupName, crossConnectionName, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *ExpressRouteCrossConnectionPeeringListResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ExpressRouteCrossConnectionPeeringList.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListCreateRequest creates the List request.
func (client *ExpressRouteCrossConnectionPeeringsClient) ListCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, options *ExpressRouteCrossConnectionPeeringsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
func (client *ExpressRouteCrossConnectionPeeringsClient) ListHandleResponse(resp *azcore.Response) (*ExpressRouteCrossConnectionPeeringListResponse, error) {
	result := ExpressRouteCrossConnectionPeeringListResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCrossConnectionPeeringList)
}

// ListHandleError handles the List error response.
func (client *ExpressRouteCrossConnectionPeeringsClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
