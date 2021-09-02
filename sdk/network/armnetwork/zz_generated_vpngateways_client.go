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
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// VPNGatewaysClient contains the methods for the VPNGateways group.
// Don't use this type directly, use NewVPNGatewaysClient() instead.
type VPNGatewaysClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewVPNGatewaysClient creates a new instance of VPNGatewaysClient with the specified values.
func NewVPNGatewaysClient(con *arm.Connection, subscriptionID string) *VPNGatewaysClient {
	return &VPNGatewaysClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates a virtual wan vpn gateway if it doesn't exist else updates the existing gateway.
// If the operation fails it returns the *CloudError error type.
func (client *VPNGatewaysClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, vpnGatewayParameters VPNGateway, options *VPNGatewaysBeginCreateOrUpdateOptions) (VPNGatewaysCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, gatewayName, vpnGatewayParameters, options)
	if err != nil {
		return VPNGatewaysCreateOrUpdatePollerResponse{}, err
	}
	result := VPNGatewaysCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VPNGatewaysClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return VPNGatewaysCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &VPNGatewaysCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates a virtual wan vpn gateway if it doesn't exist else updates the existing gateway.
// If the operation fails it returns the *CloudError error type.
func (client *VPNGatewaysClient) createOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, vpnGatewayParameters VPNGateway, options *VPNGatewaysBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, gatewayName, vpnGatewayParameters, options)
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
func (client *VPNGatewaysClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, vpnGatewayParameters VPNGateway, options *VPNGatewaysBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, vpnGatewayParameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *VPNGatewaysClient) createOrUpdateHandleError(resp *http.Response) error {
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

// BeginDelete - Deletes a virtual wan vpn gateway.
// If the operation fails it returns the *CloudError error type.
func (client *VPNGatewaysClient) BeginDelete(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysBeginDeleteOptions) (VPNGatewaysDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, gatewayName, options)
	if err != nil {
		return VPNGatewaysDeletePollerResponse{}, err
	}
	result := VPNGatewaysDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VPNGatewaysClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return VPNGatewaysDeletePollerResponse{}, err
	}
	result.Poller = &VPNGatewaysDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a virtual wan vpn gateway.
// If the operation fails it returns the *CloudError error type.
func (client *VPNGatewaysClient) deleteOperation(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, gatewayName, options)
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
func (client *VPNGatewaysClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *VPNGatewaysClient) deleteHandleError(resp *http.Response) error {
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

// Get - Retrieves the details of a virtual wan vpn gateway.
// If the operation fails it returns the *CloudError error type.
func (client *VPNGatewaysClient) Get(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysGetOptions) (VPNGatewaysGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, gatewayName, options)
	if err != nil {
		return VPNGatewaysGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VPNGatewaysGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VPNGatewaysGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VPNGatewaysClient) getCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VPNGatewaysClient) getHandleResponse(resp *http.Response) (VPNGatewaysGetResponse, error) {
	result := VPNGatewaysGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VPNGateway); err != nil {
		return VPNGatewaysGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *VPNGatewaysClient) getHandleError(resp *http.Response) error {
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

// List - Lists all the VpnGateways in a subscription.
// If the operation fails it returns the *CloudError error type.
func (client *VPNGatewaysClient) List(options *VPNGatewaysListOptions) *VPNGatewaysListPager {
	return &VPNGatewaysListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp VPNGatewaysListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListVPNGatewaysResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *VPNGatewaysClient) listCreateRequest(ctx context.Context, options *VPNGatewaysListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/vpnGateways"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VPNGatewaysClient) listHandleResponse(resp *http.Response) (VPNGatewaysListResponse, error) {
	result := VPNGatewaysListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListVPNGatewaysResult); err != nil {
		return VPNGatewaysListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *VPNGatewaysClient) listHandleError(resp *http.Response) error {
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

// ListByResourceGroup - Lists all the VpnGateways in a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *VPNGatewaysClient) ListByResourceGroup(resourceGroupName string, options *VPNGatewaysListByResourceGroupOptions) *VPNGatewaysListByResourceGroupPager {
	return &VPNGatewaysListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp VPNGatewaysListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListVPNGatewaysResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *VPNGatewaysClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *VPNGatewaysListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *VPNGatewaysClient) listByResourceGroupHandleResponse(resp *http.Response) (VPNGatewaysListByResourceGroupResponse, error) {
	result := VPNGatewaysListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListVPNGatewaysResult); err != nil {
		return VPNGatewaysListByResourceGroupResponse{}, err
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *VPNGatewaysClient) listByResourceGroupHandleError(resp *http.Response) error {
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

// BeginReset - Resets the primary of the vpn gateway in the specified resource group.
// If the operation fails it returns the *CloudError error type.
func (client *VPNGatewaysClient) BeginReset(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysBeginResetOptions) (VPNGatewaysResetPollerResponse, error) {
	resp, err := client.reset(ctx, resourceGroupName, gatewayName, options)
	if err != nil {
		return VPNGatewaysResetPollerResponse{}, err
	}
	result := VPNGatewaysResetPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VPNGatewaysClient.Reset", "location", resp, client.pl, client.resetHandleError)
	if err != nil {
		return VPNGatewaysResetPollerResponse{}, err
	}
	result.Poller = &VPNGatewaysResetPoller{
		pt: pt,
	}
	return result, nil
}

// Reset - Resets the primary of the vpn gateway in the specified resource group.
// If the operation fails it returns the *CloudError error type.
func (client *VPNGatewaysClient) reset(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysBeginResetOptions) (*http.Response, error) {
	req, err := client.resetCreateRequest(ctx, resourceGroupName, gatewayName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.resetHandleError(resp)
	}
	return resp, nil
}

// resetCreateRequest creates the Reset request.
func (client *VPNGatewaysClient) resetCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysBeginResetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/reset"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// resetHandleError handles the Reset error response.
func (client *VPNGatewaysClient) resetHandleError(resp *http.Response) error {
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

// BeginStartPacketCapture - Starts packet capture on vpn gateway in the specified resource group.
// If the operation fails it returns the *CloudError error type.
func (client *VPNGatewaysClient) BeginStartPacketCapture(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysBeginStartPacketCaptureOptions) (VPNGatewaysStartPacketCapturePollerResponse, error) {
	resp, err := client.startPacketCapture(ctx, resourceGroupName, gatewayName, options)
	if err != nil {
		return VPNGatewaysStartPacketCapturePollerResponse{}, err
	}
	result := VPNGatewaysStartPacketCapturePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VPNGatewaysClient.StartPacketCapture", "location", resp, client.pl, client.startPacketCaptureHandleError)
	if err != nil {
		return VPNGatewaysStartPacketCapturePollerResponse{}, err
	}
	result.Poller = &VPNGatewaysStartPacketCapturePoller{
		pt: pt,
	}
	return result, nil
}

// StartPacketCapture - Starts packet capture on vpn gateway in the specified resource group.
// If the operation fails it returns the *CloudError error type.
func (client *VPNGatewaysClient) startPacketCapture(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysBeginStartPacketCaptureOptions) (*http.Response, error) {
	req, err := client.startPacketCaptureCreateRequest(ctx, resourceGroupName, gatewayName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.startPacketCaptureHandleError(resp)
	}
	return resp, nil
}

// startPacketCaptureCreateRequest creates the StartPacketCapture request.
func (client *VPNGatewaysClient) startPacketCaptureCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysBeginStartPacketCaptureOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/startpacketcapture"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Parameters != nil {
		return req, runtime.MarshalAsJSON(req, *options.Parameters)
	}
	return req, nil
}

// startPacketCaptureHandleError handles the StartPacketCapture error response.
func (client *VPNGatewaysClient) startPacketCaptureHandleError(resp *http.Response) error {
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

// BeginStopPacketCapture - Stops packet capture on vpn gateway in the specified resource group.
// If the operation fails it returns the *CloudError error type.
func (client *VPNGatewaysClient) BeginStopPacketCapture(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysBeginStopPacketCaptureOptions) (VPNGatewaysStopPacketCapturePollerResponse, error) {
	resp, err := client.stopPacketCapture(ctx, resourceGroupName, gatewayName, options)
	if err != nil {
		return VPNGatewaysStopPacketCapturePollerResponse{}, err
	}
	result := VPNGatewaysStopPacketCapturePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VPNGatewaysClient.StopPacketCapture", "location", resp, client.pl, client.stopPacketCaptureHandleError)
	if err != nil {
		return VPNGatewaysStopPacketCapturePollerResponse{}, err
	}
	result.Poller = &VPNGatewaysStopPacketCapturePoller{
		pt: pt,
	}
	return result, nil
}

// StopPacketCapture - Stops packet capture on vpn gateway in the specified resource group.
// If the operation fails it returns the *CloudError error type.
func (client *VPNGatewaysClient) stopPacketCapture(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysBeginStopPacketCaptureOptions) (*http.Response, error) {
	req, err := client.stopPacketCaptureCreateRequest(ctx, resourceGroupName, gatewayName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.stopPacketCaptureHandleError(resp)
	}
	return resp, nil
}

// stopPacketCaptureCreateRequest creates the StopPacketCapture request.
func (client *VPNGatewaysClient) stopPacketCaptureCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysBeginStopPacketCaptureOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/stoppacketcapture"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Parameters != nil {
		return req, runtime.MarshalAsJSON(req, *options.Parameters)
	}
	return req, nil
}

// stopPacketCaptureHandleError handles the StopPacketCapture error response.
func (client *VPNGatewaysClient) stopPacketCaptureHandleError(resp *http.Response) error {
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

// BeginUpdateTags - Updates virtual wan vpn gateway tags.
// If the operation fails it returns the *CloudError error type.
func (client *VPNGatewaysClient) BeginUpdateTags(ctx context.Context, resourceGroupName string, gatewayName string, vpnGatewayParameters TagsObject, options *VPNGatewaysBeginUpdateTagsOptions) (VPNGatewaysUpdateTagsPollerResponse, error) {
	resp, err := client.updateTags(ctx, resourceGroupName, gatewayName, vpnGatewayParameters, options)
	if err != nil {
		return VPNGatewaysUpdateTagsPollerResponse{}, err
	}
	result := VPNGatewaysUpdateTagsPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VPNGatewaysClient.UpdateTags", "azure-async-operation", resp, client.pl, client.updateTagsHandleError)
	if err != nil {
		return VPNGatewaysUpdateTagsPollerResponse{}, err
	}
	result.Poller = &VPNGatewaysUpdateTagsPoller{
		pt: pt,
	}
	return result, nil
}

// UpdateTags - Updates virtual wan vpn gateway tags.
// If the operation fails it returns the *CloudError error type.
func (client *VPNGatewaysClient) updateTags(ctx context.Context, resourceGroupName string, gatewayName string, vpnGatewayParameters TagsObject, options *VPNGatewaysBeginUpdateTagsOptions) (*http.Response, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, gatewayName, vpnGatewayParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.updateTagsHandleError(resp)
	}
	return resp, nil
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *VPNGatewaysClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, vpnGatewayParameters TagsObject, options *VPNGatewaysBeginUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, vpnGatewayParameters)
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *VPNGatewaysClient) updateTagsHandleError(resp *http.Response) error {
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
