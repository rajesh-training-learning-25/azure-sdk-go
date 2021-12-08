//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridnetwork

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// NetworkFunctionVendorSKUsClient contains the methods for the NetworkFunctionVendorSKUs group.
// Don't use this type directly, use NewNetworkFunctionVendorSKUsClient() instead.
type NetworkFunctionVendorSKUsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewNetworkFunctionVendorSKUsClient creates a new instance of NetworkFunctionVendorSKUsClient with the specified values.
func NewNetworkFunctionVendorSKUsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *NetworkFunctionVendorSKUsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &NetworkFunctionVendorSKUsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// ListBySKU - Lists information about network function vendor sku details.
// If the operation fails it returns the *ErrorResponse error type.
func (client *NetworkFunctionVendorSKUsClient) ListBySKU(vendorName string, vendorSKUName string, options *NetworkFunctionVendorSKUsListBySKUOptions) *NetworkFunctionVendorSKUsListBySKUPager {
	return &NetworkFunctionVendorSKUsListBySKUPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySKUCreateRequest(ctx, vendorName, vendorSKUName, options)
		},
		advancer: func(ctx context.Context, resp NetworkFunctionVendorSKUsListBySKUResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.NetworkFunctionSKUDetails.NextLink)
		},
	}
}

// listBySKUCreateRequest creates the ListBySKU request.
func (client *NetworkFunctionVendorSKUsClient) listBySKUCreateRequest(ctx context.Context, vendorName string, vendorSKUName string, options *NetworkFunctionVendorSKUsListBySKUOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/networkFunctionVendors/{vendorName}/vendorSkus/{vendorSkuName}"
	if vendorName == "" {
		return nil, errors.New("parameter vendorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vendorName}", url.PathEscape(vendorName))
	if vendorSKUName == "" {
		return nil, errors.New("parameter vendorSKUName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vendorSkuName}", url.PathEscape(vendorSKUName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySKUHandleResponse handles the ListBySKU response.
func (client *NetworkFunctionVendorSKUsClient) listBySKUHandleResponse(resp *http.Response) (NetworkFunctionVendorSKUsListBySKUResponse, error) {
	result := NetworkFunctionVendorSKUsListBySKUResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkFunctionSKUDetails); err != nil {
		return NetworkFunctionVendorSKUsListBySKUResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listBySKUHandleError handles the ListBySKU error response.
func (client *NetworkFunctionVendorSKUsClient) listBySKUHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByVendor - Lists all network function vendor sku details in a vendor.
// If the operation fails it returns the *ErrorResponse error type.
func (client *NetworkFunctionVendorSKUsClient) ListByVendor(vendorName string, options *NetworkFunctionVendorSKUsListByVendorOptions) *NetworkFunctionVendorSKUsListByVendorPager {
	return &NetworkFunctionVendorSKUsListByVendorPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByVendorCreateRequest(ctx, vendorName, options)
		},
		advancer: func(ctx context.Context, resp NetworkFunctionVendorSKUsListByVendorResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.NetworkFunctionSKUListResult.NextLink)
		},
	}
}

// listByVendorCreateRequest creates the ListByVendor request.
func (client *NetworkFunctionVendorSKUsClient) listByVendorCreateRequest(ctx context.Context, vendorName string, options *NetworkFunctionVendorSKUsListByVendorOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/networkFunctionVendors/{vendorName}/vendorSkus"
	if vendorName == "" {
		return nil, errors.New("parameter vendorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vendorName}", url.PathEscape(vendorName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByVendorHandleResponse handles the ListByVendor response.
func (client *NetworkFunctionVendorSKUsClient) listByVendorHandleResponse(resp *http.Response) (NetworkFunctionVendorSKUsListByVendorResponse, error) {
	result := NetworkFunctionVendorSKUsListByVendorResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkFunctionSKUListResult); err != nil {
		return NetworkFunctionVendorSKUsListByVendorResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByVendorHandleError handles the ListByVendor error response.
func (client *NetworkFunctionVendorSKUsClient) listByVendorHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
