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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
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
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVPNServerConfigurationsAssociatedWithVirtualWanClient creates a new instance of VPNServerConfigurationsAssociatedWithVirtualWanClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVPNServerConfigurationsAssociatedWithVirtualWanClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *VPNServerConfigurationsAssociatedWithVirtualWanClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &VPNServerConfigurationsAssociatedWithVirtualWanClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginList - Gives the list of VpnServerConfigurations associated with Virtual Wan in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// virtualWANName - The name of the VirtualWAN whose associated VpnServerConfigurations is needed.
// options - VPNServerConfigurationsAssociatedWithVirtualWanClientBeginListOptions contains the optional parameters for the
// VPNServerConfigurationsAssociatedWithVirtualWanClient.BeginList method.
func (client *VPNServerConfigurationsAssociatedWithVirtualWanClient) BeginList(ctx context.Context, resourceGroupName string, virtualWANName string, options *VPNServerConfigurationsAssociatedWithVirtualWanClientBeginListOptions) (VPNServerConfigurationsAssociatedWithVirtualWanClientListPollerResponse, error) {
	resp, err := client.listOperation(ctx, resourceGroupName, virtualWANName, options)
	if err != nil {
		return VPNServerConfigurationsAssociatedWithVirtualWanClientListPollerResponse{}, err
	}
	result := VPNServerConfigurationsAssociatedWithVirtualWanClientListPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VPNServerConfigurationsAssociatedWithVirtualWanClient.List", "location", resp, client.pl)
	if err != nil {
		return VPNServerConfigurationsAssociatedWithVirtualWanClientListPollerResponse{}, err
	}
	result.Poller = &VPNServerConfigurationsAssociatedWithVirtualWanClientListPoller{
		pt: pt,
	}
	return result, nil
}

// List - Gives the list of VpnServerConfigurations associated with Virtual Wan in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VPNServerConfigurationsAssociatedWithVirtualWanClient) listOperation(ctx context.Context, resourceGroupName string, virtualWANName string, options *VPNServerConfigurationsAssociatedWithVirtualWanClientBeginListOptions) (*http.Response, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, virtualWANName, options)
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

// listCreateRequest creates the List request.
func (client *VPNServerConfigurationsAssociatedWithVirtualWanClient) listCreateRequest(ctx context.Context, resourceGroupName string, virtualWANName string, options *VPNServerConfigurationsAssociatedWithVirtualWanClientBeginListOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}
