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

// VPNServerConfigurationsAssociatedWithVirtualWanClient contains the methods for the VPNServerConfigurationsAssociatedWithVirtualWan group.
// Don't use this type directly, use NewVPNServerConfigurationsAssociatedWithVirtualWanClient() instead.
type VPNServerConfigurationsAssociatedWithVirtualWanClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewVPNServerConfigurationsAssociatedWithVirtualWanClient creates a new instance of VPNServerConfigurationsAssociatedWithVirtualWanClient with the specified values.
func NewVPNServerConfigurationsAssociatedWithVirtualWanClient(con *arm.Connection, subscriptionID string) *VPNServerConfigurationsAssociatedWithVirtualWanClient {
	return &VPNServerConfigurationsAssociatedWithVirtualWanClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginList - Gives the list of VpnServerConfigurations associated with Virtual Wan in a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *VPNServerConfigurationsAssociatedWithVirtualWanClient) BeginList(ctx context.Context, resourceGroupName string, virtualWANName string, options *VPNServerConfigurationsAssociatedWithVirtualWanBeginListOptions) (VPNServerConfigurationsAssociatedWithVirtualWanListPollerResponse, error) {
	resp, err := client.listOperation(ctx, resourceGroupName, virtualWANName, options)
	if err != nil {
		return VPNServerConfigurationsAssociatedWithVirtualWanListPollerResponse{}, err
	}
	result := VPNServerConfigurationsAssociatedWithVirtualWanListPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VPNServerConfigurationsAssociatedWithVirtualWanClient.List", "location", resp, client.pl, client.listHandleError)
	if err != nil {
		return VPNServerConfigurationsAssociatedWithVirtualWanListPollerResponse{}, err
	}
	result.Poller = &VPNServerConfigurationsAssociatedWithVirtualWanListPoller{
		pt: pt,
	}
	return result, nil
}

// List - Gives the list of VpnServerConfigurations associated with Virtual Wan in a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *VPNServerConfigurationsAssociatedWithVirtualWanClient) listOperation(ctx context.Context, resourceGroupName string, virtualWANName string, options *VPNServerConfigurationsAssociatedWithVirtualWanBeginListOptions) (*http.Response, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, virtualWANName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.listHandleError(resp)
	}
	return resp, nil
}

// listCreateRequest creates the List request.
func (client *VPNServerConfigurationsAssociatedWithVirtualWanClient) listCreateRequest(ctx context.Context, resourceGroupName string, virtualWANName string, options *VPNServerConfigurationsAssociatedWithVirtualWanBeginListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{virtualWANName}/vpnServerConfigurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualWANName == "" {
		return nil, errors.New("parameter virtualWANName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualWANName}", url.PathEscape(virtualWANName))
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

// listHandleError handles the List error response.
func (client *VPNServerConfigurationsAssociatedWithVirtualWanClient) listHandleError(resp *http.Response) error {
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
