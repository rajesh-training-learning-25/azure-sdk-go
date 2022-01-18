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

// VirtualApplianceSKUsClient contains the methods for the VirtualApplianceSKUs group.
// Don't use this type directly, use NewVirtualApplianceSKUsClient() instead.
type VirtualApplianceSKUsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVirtualApplianceSKUsClient creates a new instance of VirtualApplianceSKUsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVirtualApplianceSKUsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *VirtualApplianceSKUsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &VirtualApplianceSKUsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Get - Retrieves a single available sku for network virtual appliance.
// If the operation fails it returns an *azcore.ResponseError type.
// skuName - Name of the Sku.
// options - VirtualApplianceSKUsClientGetOptions contains the optional parameters for the VirtualApplianceSKUsClient.Get
// method.
func (client *VirtualApplianceSKUsClient) Get(ctx context.Context, skuName string, options *VirtualApplianceSKUsClientGetOptions) (VirtualApplianceSKUsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, skuName, options)
	if err != nil {
		return VirtualApplianceSKUsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualApplianceSKUsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualApplianceSKUsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualApplianceSKUsClient) getCreateRequest(ctx context.Context, skuName string, options *VirtualApplianceSKUsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkVirtualApplianceSkus/{skuName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if skuName == "" {
		return nil, errors.New("parameter skuName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{skuName}", url.PathEscape(skuName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualApplianceSKUsClient) getHandleResponse(resp *http.Response) (VirtualApplianceSKUsClientGetResponse, error) {
	result := VirtualApplianceSKUsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualApplianceSKU); err != nil {
		return VirtualApplianceSKUsClientGetResponse{}, err
	}
	return result, nil
}

// List - List all SKUs available for a virtual appliance.
// If the operation fails it returns an *azcore.ResponseError type.
// options - VirtualApplianceSKUsClientListOptions contains the optional parameters for the VirtualApplianceSKUsClient.List
// method.
func (client *VirtualApplianceSKUsClient) List(options *VirtualApplianceSKUsClientListOptions) *VirtualApplianceSKUsClientListPager {
	return &VirtualApplianceSKUsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp VirtualApplianceSKUsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.VirtualApplianceSKUListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *VirtualApplianceSKUsClient) listCreateRequest(ctx context.Context, options *VirtualApplianceSKUsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkVirtualApplianceSkus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualApplianceSKUsClient) listHandleResponse(resp *http.Response) (VirtualApplianceSKUsClientListResponse, error) {
	result := VirtualApplianceSKUsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualApplianceSKUListResult); err != nil {
		return VirtualApplianceSKUsClientListResponse{}, err
	}
	return result, nil
}
