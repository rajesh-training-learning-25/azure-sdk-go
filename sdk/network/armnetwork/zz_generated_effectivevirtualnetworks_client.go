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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// EffectiveVirtualNetworksClient contains the methods for the EffectiveVirtualNetworks group.
// Don't use this type directly, use NewEffectiveVirtualNetworksClient() instead.
type EffectiveVirtualNetworksClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewEffectiveVirtualNetworksClient creates a new instance of EffectiveVirtualNetworksClient with the specified values.
func NewEffectiveVirtualNetworksClient(con *arm.Connection, subscriptionID string) *EffectiveVirtualNetworksClient {
	return &EffectiveVirtualNetworksClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// ListByNetworkGroup - Lists all effective virtual networks by specified network group.
// If the operation fails it returns the *CloudErrorAutoGenerated error type.
func (client *EffectiveVirtualNetworksClient) ListByNetworkGroup(ctx context.Context, resourceGroupName string, networkManagerName string, networkGroupName string, parameters QueryRequestOptions, options *EffectiveVirtualNetworksListByNetworkGroupOptions) (EffectiveVirtualNetworksListByNetworkGroupResponse, error) {
	req, err := client.listByNetworkGroupCreateRequest(ctx, resourceGroupName, networkManagerName, networkGroupName, parameters, options)
	if err != nil {
		return EffectiveVirtualNetworksListByNetworkGroupResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EffectiveVirtualNetworksListByNetworkGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EffectiveVirtualNetworksListByNetworkGroupResponse{}, client.listByNetworkGroupHandleError(resp)
	}
	return client.listByNetworkGroupHandleResponse(resp)
}

// listByNetworkGroupCreateRequest creates the ListByNetworkGroup request.
func (client *EffectiveVirtualNetworksClient) listByNetworkGroupCreateRequest(ctx context.Context, resourceGroupName string, networkManagerName string, networkGroupName string, parameters QueryRequestOptions, options *EffectiveVirtualNetworksListByNetworkGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/networkGroups/{networkGroupName}/listEffectiveVirtualNetworks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkManagerName == "" {
		return nil, errors.New("parameter networkManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkManagerName}", url.PathEscape(networkManagerName))
	if networkGroupName == "" {
		return nil, errors.New("parameter networkGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkGroupName}", url.PathEscape(networkGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// listByNetworkGroupHandleResponse handles the ListByNetworkGroup response.
func (client *EffectiveVirtualNetworksClient) listByNetworkGroupHandleResponse(resp *http.Response) (EffectiveVirtualNetworksListByNetworkGroupResponse, error) {
	result := EffectiveVirtualNetworksListByNetworkGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.EffectiveVirtualNetworksListResult); err != nil {
		return EffectiveVirtualNetworksListByNetworkGroupResponse{}, err
	}
	return result, nil
}

// listByNetworkGroupHandleError handles the ListByNetworkGroup error response.
func (client *EffectiveVirtualNetworksClient) listByNetworkGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudErrorAutoGenerated{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByNetworkManager - List effective virtual networks in a network manager.
// If the operation fails it returns the *CloudErrorAutoGenerated error type.
func (client *EffectiveVirtualNetworksClient) ListByNetworkManager(ctx context.Context, resourceGroupName string, networkManagerName string, parameters EffectiveVirtualNetworksParameter, options *EffectiveVirtualNetworksListByNetworkManagerOptions) (EffectiveVirtualNetworksListByNetworkManagerResponse, error) {
	req, err := client.listByNetworkManagerCreateRequest(ctx, resourceGroupName, networkManagerName, parameters, options)
	if err != nil {
		return EffectiveVirtualNetworksListByNetworkManagerResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EffectiveVirtualNetworksListByNetworkManagerResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EffectiveVirtualNetworksListByNetworkManagerResponse{}, client.listByNetworkManagerHandleError(resp)
	}
	return client.listByNetworkManagerHandleResponse(resp)
}

// listByNetworkManagerCreateRequest creates the ListByNetworkManager request.
func (client *EffectiveVirtualNetworksClient) listByNetworkManagerCreateRequest(ctx context.Context, resourceGroupName string, networkManagerName string, parameters EffectiveVirtualNetworksParameter, options *EffectiveVirtualNetworksListByNetworkManagerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/listEffectiveVirtualNetworks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkManagerName == "" {
		return nil, errors.New("parameter networkManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkManagerName}", url.PathEscape(networkManagerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// listByNetworkManagerHandleResponse handles the ListByNetworkManager response.
func (client *EffectiveVirtualNetworksClient) listByNetworkManagerHandleResponse(resp *http.Response) (EffectiveVirtualNetworksListByNetworkManagerResponse, error) {
	result := EffectiveVirtualNetworksListByNetworkManagerResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.EffectiveVirtualNetworksListResult); err != nil {
		return EffectiveVirtualNetworksListByNetworkManagerResponse{}, err
	}
	return result, nil
}

// listByNetworkManagerHandleError handles the ListByNetworkManager error response.
func (client *EffectiveVirtualNetworksClient) listByNetworkManagerHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudErrorAutoGenerated{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
