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

// VirtualNetworkGatewayConnectionsOperations contains the methods for the VirtualNetworkGatewayConnections group.
type VirtualNetworkGatewayConnectionsOperations interface {
	// BeginCreateOrUpdate - Creates or updates a virtual network gateway connection in the specified resource group.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters VirtualNetworkGatewayConnection) (*VirtualNetworkGatewayConnectionPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (VirtualNetworkGatewayConnectionPoller, error)
	// BeginDelete - Deletes the specified virtual network Gateway connection.
	BeginDelete(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets the specified virtual network gateway connection by resource group.
	Get(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string) (*VirtualNetworkGatewayConnectionResponse, error)
	// GetSharedKey - The Get VirtualNetworkGatewayConnectionSharedKey operation retrieves information about the specified virtual network gateway connection shared key through Network resource provider.
	GetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string) (*ConnectionSharedKeyResponse, error)
	// List - The List VirtualNetworkGatewayConnections operation retrieves all the virtual network gateways connections created.
	List(resourceGroupName string) VirtualNetworkGatewayConnectionListResultPager
	// BeginResetSharedKey - The VirtualNetworkGatewayConnectionResetSharedKey operation resets the virtual network gateway connection shared key for passed virtual network gateway connection in the specified resource group through Network resource provider.
	BeginResetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionResetSharedKey) (*ConnectionResetSharedKeyPollerResponse, error)
	// ResumeResetSharedKey - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeResetSharedKey(token string) (ConnectionResetSharedKeyPoller, error)
	// BeginSetSharedKey - The Put VirtualNetworkGatewayConnectionSharedKey operation sets the virtual network gateway connection shared key for passed virtual network gateway connection in the specified resource group through Network resource provider.
	BeginSetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionSharedKey) (*ConnectionSharedKeyPollerResponse, error)
	// ResumeSetSharedKey - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeSetSharedKey(token string) (ConnectionSharedKeyPoller, error)
	// BeginStartPacketCapture - Starts packet capture on virtual network gateway connection in the specified resource group.
	BeginStartPacketCapture(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, virtualNetworkGatewayConnectionsStartPacketCaptureOptions *VirtualNetworkGatewayConnectionsStartPacketCaptureOptions) (*StringPollerResponse, error)
	// ResumeStartPacketCapture - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeStartPacketCapture(token string) (StringPoller, error)
	// BeginStopPacketCapture - Stops packet capture on virtual network gateway connection in the specified resource group.
	BeginStopPacketCapture(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters VpnPacketCaptureStopParameters) (*StringPollerResponse, error)
	// ResumeStopPacketCapture - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeStopPacketCapture(token string) (StringPoller, error)
	// BeginUpdateTags - Updates a virtual network gateway connection tags.
	BeginUpdateTags(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters TagsObject) (*VirtualNetworkGatewayConnectionPollerResponse, error)
	// ResumeUpdateTags - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeUpdateTags(token string) (VirtualNetworkGatewayConnectionPoller, error)
}

// VirtualNetworkGatewayConnectionsClient implements the VirtualNetworkGatewayConnectionsOperations interface.
// Don't use this type directly, use NewVirtualNetworkGatewayConnectionsClient() instead.
type VirtualNetworkGatewayConnectionsClient struct {
	*Client
	subscriptionID string
}

// NewVirtualNetworkGatewayConnectionsClient creates a new instance of VirtualNetworkGatewayConnectionsClient with the specified values.
func NewVirtualNetworkGatewayConnectionsClient(c *Client, subscriptionID string) VirtualNetworkGatewayConnectionsOperations {
	return &VirtualNetworkGatewayConnectionsClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *VirtualNetworkGatewayConnectionsClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

func (client *VirtualNetworkGatewayConnectionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters VirtualNetworkGatewayConnection) (*VirtualNetworkGatewayConnectionPollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, parameters)
	if err != nil {
		return nil, err
	}
	result := &VirtualNetworkGatewayConnectionPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualNetworkGatewayConnectionsClient.CreateOrUpdate", "azure-async-operation", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &virtualNetworkGatewayConnectionPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*VirtualNetworkGatewayConnectionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *VirtualNetworkGatewayConnectionsClient) ResumeCreateOrUpdate(token string) (VirtualNetworkGatewayConnectionPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualNetworkGatewayConnectionsClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &virtualNetworkGatewayConnectionPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates a virtual network gateway connection in the specified resource group.
func (client *VirtualNetworkGatewayConnectionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters VirtualNetworkGatewayConnection) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.CreateOrUpdateHandleError(resp)
	}
	return resp, nil
}

// CreateOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualNetworkGatewayConnectionsClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters VirtualNetworkGatewayConnection) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayConnectionName}", url.PathEscape(virtualNetworkGatewayConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *VirtualNetworkGatewayConnectionsClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*VirtualNetworkGatewayConnectionResponse, error) {
	result := VirtualNetworkGatewayConnectionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualNetworkGatewayConnection)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *VirtualNetworkGatewayConnectionsClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

func (client *VirtualNetworkGatewayConnectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, virtualNetworkGatewayConnectionName)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualNetworkGatewayConnectionsClient.Delete", "location", resp, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *VirtualNetworkGatewayConnectionsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualNetworkGatewayConnectionsClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified virtual network Gateway connection.
func (client *VirtualNetworkGatewayConnectionsClient) Delete(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, virtualNetworkGatewayConnectionName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.DeleteHandleError(resp)
	}
	return resp, nil
}

// DeleteCreateRequest creates the Delete request.
func (client *VirtualNetworkGatewayConnectionsClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayConnectionName}", url.PathEscape(virtualNetworkGatewayConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DeleteHandleError handles the Delete error response.
func (client *VirtualNetworkGatewayConnectionsClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get - Gets the specified virtual network gateway connection by resource group.
func (client *VirtualNetworkGatewayConnectionsClient) Get(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string) (*VirtualNetworkGatewayConnectionResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, virtualNetworkGatewayConnectionName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
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
func (client *VirtualNetworkGatewayConnectionsClient) GetCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayConnectionName}", url.PathEscape(virtualNetworkGatewayConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *VirtualNetworkGatewayConnectionsClient) GetHandleResponse(resp *azcore.Response) (*VirtualNetworkGatewayConnectionResponse, error) {
	result := VirtualNetworkGatewayConnectionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualNetworkGatewayConnection)
}

// GetHandleError handles the Get error response.
func (client *VirtualNetworkGatewayConnectionsClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetSharedKey - The Get VirtualNetworkGatewayConnectionSharedKey operation retrieves information about the specified virtual network gateway connection shared key through Network resource provider.
func (client *VirtualNetworkGatewayConnectionsClient) GetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string) (*ConnectionSharedKeyResponse, error) {
	req, err := client.GetSharedKeyCreateRequest(ctx, resourceGroupName, virtualNetworkGatewayConnectionName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetSharedKeyHandleError(resp)
	}
	result, err := client.GetSharedKeyHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetSharedKeyCreateRequest creates the GetSharedKey request.
func (client *VirtualNetworkGatewayConnectionsClient) GetSharedKeyCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}/sharedkey"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayConnectionName}", url.PathEscape(virtualNetworkGatewayConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetSharedKeyHandleResponse handles the GetSharedKey response.
func (client *VirtualNetworkGatewayConnectionsClient) GetSharedKeyHandleResponse(resp *azcore.Response) (*ConnectionSharedKeyResponse, error) {
	result := ConnectionSharedKeyResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ConnectionSharedKey)
}

// GetSharedKeyHandleError handles the GetSharedKey error response.
func (client *VirtualNetworkGatewayConnectionsClient) GetSharedKeyHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - The List VirtualNetworkGatewayConnections operation retrieves all the virtual network gateways connections created.
func (client *VirtualNetworkGatewayConnectionsClient) List(resourceGroupName string) VirtualNetworkGatewayConnectionListResultPager {
	return &virtualNetworkGatewayConnectionListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, resourceGroupName)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *VirtualNetworkGatewayConnectionListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.VirtualNetworkGatewayConnectionListResult.NextLink)
		},
	}
}

// ListCreateRequest creates the List request.
func (client *VirtualNetworkGatewayConnectionsClient) ListCreateRequest(ctx context.Context, resourceGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListHandleResponse handles the List response.
func (client *VirtualNetworkGatewayConnectionsClient) ListHandleResponse(resp *azcore.Response) (*VirtualNetworkGatewayConnectionListResultResponse, error) {
	result := VirtualNetworkGatewayConnectionListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualNetworkGatewayConnectionListResult)
}

// ListHandleError handles the List error response.
func (client *VirtualNetworkGatewayConnectionsClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

func (client *VirtualNetworkGatewayConnectionsClient) BeginResetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionResetSharedKey) (*ConnectionResetSharedKeyPollerResponse, error) {
	resp, err := client.ResetSharedKey(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, parameters)
	if err != nil {
		return nil, err
	}
	result := &ConnectionResetSharedKeyPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualNetworkGatewayConnectionsClient.ResetSharedKey", "location", resp, client.ResetSharedKeyHandleError)
	if err != nil {
		return nil, err
	}
	poller := &connectionResetSharedKeyPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ConnectionResetSharedKeyResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *VirtualNetworkGatewayConnectionsClient) ResumeResetSharedKey(token string) (ConnectionResetSharedKeyPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualNetworkGatewayConnectionsClient.ResetSharedKey", token, client.ResetSharedKeyHandleError)
	if err != nil {
		return nil, err
	}
	return &connectionResetSharedKeyPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// ResetSharedKey - The VirtualNetworkGatewayConnectionResetSharedKey operation resets the virtual network gateway connection shared key for passed virtual network gateway connection in the specified resource group through Network resource provider.
func (client *VirtualNetworkGatewayConnectionsClient) ResetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionResetSharedKey) (*azcore.Response, error) {
	req, err := client.ResetSharedKeyCreateRequest(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.ResetSharedKeyHandleError(resp)
	}
	return resp, nil
}

// ResetSharedKeyCreateRequest creates the ResetSharedKey request.
func (client *VirtualNetworkGatewayConnectionsClient) ResetSharedKeyCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionResetSharedKey) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}/sharedkey/reset"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayConnectionName}", url.PathEscape(virtualNetworkGatewayConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// ResetSharedKeyHandleResponse handles the ResetSharedKey response.
func (client *VirtualNetworkGatewayConnectionsClient) ResetSharedKeyHandleResponse(resp *azcore.Response) (*ConnectionResetSharedKeyResponse, error) {
	result := ConnectionResetSharedKeyResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ConnectionResetSharedKey)
}

// ResetSharedKeyHandleError handles the ResetSharedKey error response.
func (client *VirtualNetworkGatewayConnectionsClient) ResetSharedKeyHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

func (client *VirtualNetworkGatewayConnectionsClient) BeginSetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionSharedKey) (*ConnectionSharedKeyPollerResponse, error) {
	resp, err := client.SetSharedKey(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, parameters)
	if err != nil {
		return nil, err
	}
	result := &ConnectionSharedKeyPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualNetworkGatewayConnectionsClient.SetSharedKey", "azure-async-operation", resp, client.SetSharedKeyHandleError)
	if err != nil {
		return nil, err
	}
	poller := &connectionSharedKeyPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ConnectionSharedKeyResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *VirtualNetworkGatewayConnectionsClient) ResumeSetSharedKey(token string) (ConnectionSharedKeyPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualNetworkGatewayConnectionsClient.SetSharedKey", token, client.SetSharedKeyHandleError)
	if err != nil {
		return nil, err
	}
	return &connectionSharedKeyPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// SetSharedKey - The Put VirtualNetworkGatewayConnectionSharedKey operation sets the virtual network gateway connection shared key for passed virtual network gateway connection in the specified resource group through Network resource provider.
func (client *VirtualNetworkGatewayConnectionsClient) SetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionSharedKey) (*azcore.Response, error) {
	req, err := client.SetSharedKeyCreateRequest(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.SetSharedKeyHandleError(resp)
	}
	return resp, nil
}

// SetSharedKeyCreateRequest creates the SetSharedKey request.
func (client *VirtualNetworkGatewayConnectionsClient) SetSharedKeyCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionSharedKey) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}/sharedkey"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayConnectionName}", url.PathEscape(virtualNetworkGatewayConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// SetSharedKeyHandleResponse handles the SetSharedKey response.
func (client *VirtualNetworkGatewayConnectionsClient) SetSharedKeyHandleResponse(resp *azcore.Response) (*ConnectionSharedKeyResponse, error) {
	result := ConnectionSharedKeyResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ConnectionSharedKey)
}

// SetSharedKeyHandleError handles the SetSharedKey error response.
func (client *VirtualNetworkGatewayConnectionsClient) SetSharedKeyHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

func (client *VirtualNetworkGatewayConnectionsClient) BeginStartPacketCapture(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, virtualNetworkGatewayConnectionsStartPacketCaptureOptions *VirtualNetworkGatewayConnectionsStartPacketCaptureOptions) (*StringPollerResponse, error) {
	resp, err := client.StartPacketCapture(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, virtualNetworkGatewayConnectionsStartPacketCaptureOptions)
	if err != nil {
		return nil, err
	}
	result := &StringPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualNetworkGatewayConnectionsClient.StartPacketCapture", "location", resp, client.StartPacketCaptureHandleError)
	if err != nil {
		return nil, err
	}
	poller := &stringPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*StringResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *VirtualNetworkGatewayConnectionsClient) ResumeStartPacketCapture(token string) (StringPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualNetworkGatewayConnectionsClient.StartPacketCapture", token, client.StartPacketCaptureHandleError)
	if err != nil {
		return nil, err
	}
	return &stringPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// StartPacketCapture - Starts packet capture on virtual network gateway connection in the specified resource group.
func (client *VirtualNetworkGatewayConnectionsClient) StartPacketCapture(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, virtualNetworkGatewayConnectionsStartPacketCaptureOptions *VirtualNetworkGatewayConnectionsStartPacketCaptureOptions) (*azcore.Response, error) {
	req, err := client.StartPacketCaptureCreateRequest(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, virtualNetworkGatewayConnectionsStartPacketCaptureOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.StartPacketCaptureHandleError(resp)
	}
	return resp, nil
}

// StartPacketCaptureCreateRequest creates the StartPacketCapture request.
func (client *VirtualNetworkGatewayConnectionsClient) StartPacketCaptureCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, virtualNetworkGatewayConnectionsStartPacketCaptureOptions *VirtualNetworkGatewayConnectionsStartPacketCaptureOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}/startPacketCapture"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayConnectionName}", url.PathEscape(virtualNetworkGatewayConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	if virtualNetworkGatewayConnectionsStartPacketCaptureOptions != nil {
		return req, req.MarshalAsJSON(virtualNetworkGatewayConnectionsStartPacketCaptureOptions.Parameters)
	}
	return req, nil
}

// StartPacketCaptureHandleResponse handles the StartPacketCapture response.
func (client *VirtualNetworkGatewayConnectionsClient) StartPacketCaptureHandleResponse(resp *azcore.Response) (*StringResponse, error) {
	result := StringResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// StartPacketCaptureHandleError handles the StartPacketCapture error response.
func (client *VirtualNetworkGatewayConnectionsClient) StartPacketCaptureHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

func (client *VirtualNetworkGatewayConnectionsClient) BeginStopPacketCapture(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters VpnPacketCaptureStopParameters) (*StringPollerResponse, error) {
	resp, err := client.StopPacketCapture(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, parameters)
	if err != nil {
		return nil, err
	}
	result := &StringPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualNetworkGatewayConnectionsClient.StopPacketCapture", "location", resp, client.StopPacketCaptureHandleError)
	if err != nil {
		return nil, err
	}
	poller := &stringPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*StringResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *VirtualNetworkGatewayConnectionsClient) ResumeStopPacketCapture(token string) (StringPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualNetworkGatewayConnectionsClient.StopPacketCapture", token, client.StopPacketCaptureHandleError)
	if err != nil {
		return nil, err
	}
	return &stringPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// StopPacketCapture - Stops packet capture on virtual network gateway connection in the specified resource group.
func (client *VirtualNetworkGatewayConnectionsClient) StopPacketCapture(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters VpnPacketCaptureStopParameters) (*azcore.Response, error) {
	req, err := client.StopPacketCaptureCreateRequest(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.StopPacketCaptureHandleError(resp)
	}
	return resp, nil
}

// StopPacketCaptureCreateRequest creates the StopPacketCapture request.
func (client *VirtualNetworkGatewayConnectionsClient) StopPacketCaptureCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters VpnPacketCaptureStopParameters) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}/stopPacketCapture"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayConnectionName}", url.PathEscape(virtualNetworkGatewayConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// StopPacketCaptureHandleResponse handles the StopPacketCapture response.
func (client *VirtualNetworkGatewayConnectionsClient) StopPacketCaptureHandleResponse(resp *azcore.Response) (*StringResponse, error) {
	result := StringResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// StopPacketCaptureHandleError handles the StopPacketCapture error response.
func (client *VirtualNetworkGatewayConnectionsClient) StopPacketCaptureHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

func (client *VirtualNetworkGatewayConnectionsClient) BeginUpdateTags(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters TagsObject) (*VirtualNetworkGatewayConnectionPollerResponse, error) {
	resp, err := client.UpdateTags(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, parameters)
	if err != nil {
		return nil, err
	}
	result := &VirtualNetworkGatewayConnectionPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualNetworkGatewayConnectionsClient.UpdateTags", "azure-async-operation", resp, client.UpdateTagsHandleError)
	if err != nil {
		return nil, err
	}
	poller := &virtualNetworkGatewayConnectionPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*VirtualNetworkGatewayConnectionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *VirtualNetworkGatewayConnectionsClient) ResumeUpdateTags(token string) (VirtualNetworkGatewayConnectionPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualNetworkGatewayConnectionsClient.UpdateTags", token, client.UpdateTagsHandleError)
	if err != nil {
		return nil, err
	}
	return &virtualNetworkGatewayConnectionPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// UpdateTags - Updates a virtual network gateway connection tags.
func (client *VirtualNetworkGatewayConnectionsClient) UpdateTags(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters TagsObject) (*azcore.Response, error) {
	req, err := client.UpdateTagsCreateRequest(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.UpdateTagsHandleError(resp)
	}
	return resp, nil
}

// UpdateTagsCreateRequest creates the UpdateTags request.
func (client *VirtualNetworkGatewayConnectionsClient) UpdateTagsCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters TagsObject) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayConnectionName}", url.PathEscape(virtualNetworkGatewayConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// UpdateTagsHandleResponse handles the UpdateTags response.
func (client *VirtualNetworkGatewayConnectionsClient) UpdateTagsHandleResponse(resp *azcore.Response) (*VirtualNetworkGatewayConnectionResponse, error) {
	result := VirtualNetworkGatewayConnectionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualNetworkGatewayConnection)
}

// UpdateTagsHandleError handles the UpdateTags error response.
func (client *VirtualNetworkGatewayConnectionsClient) UpdateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
