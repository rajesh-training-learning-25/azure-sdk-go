//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetapp

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ResourceClient contains the methods for the NetAppResource group.
// Don't use this type directly, use NewResourceClient() instead.
type ResourceClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewResourceClient creates a new instance of ResourceClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewResourceClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ResourceClient, error) {
	cl, err := arm.NewClient(moduleName+".ResourceClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ResourceClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CheckFilePathAvailability - Check if a file path is available.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
//   - location - The name of Azure region.
//   - body - File path availability request.
//   - options - ResourceClientCheckFilePathAvailabilityOptions contains the optional parameters for the ResourceClient.CheckFilePathAvailability
//     method.
func (client *ResourceClient) CheckFilePathAvailability(ctx context.Context, location string, body FilePathAvailabilityRequest, options *ResourceClientCheckFilePathAvailabilityOptions) (ResourceClientCheckFilePathAvailabilityResponse, error) {
	req, err := client.checkFilePathAvailabilityCreateRequest(ctx, location, body, options)
	if err != nil {
		return ResourceClientCheckFilePathAvailabilityResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ResourceClientCheckFilePathAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ResourceClientCheckFilePathAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkFilePathAvailabilityHandleResponse(resp)
}

// checkFilePathAvailabilityCreateRequest creates the CheckFilePathAvailability request.
func (client *ResourceClient) checkFilePathAvailabilityCreateRequest(ctx context.Context, location string, body FilePathAvailabilityRequest, options *ResourceClientCheckFilePathAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.NetApp/locations/{location}/checkFilePathAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// checkFilePathAvailabilityHandleResponse handles the CheckFilePathAvailability response.
func (client *ResourceClient) checkFilePathAvailabilityHandleResponse(resp *http.Response) (ResourceClientCheckFilePathAvailabilityResponse, error) {
	result := ResourceClientCheckFilePathAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckAvailabilityResponse); err != nil {
		return ResourceClientCheckFilePathAvailabilityResponse{}, err
	}
	return result, nil
}

// CheckNameAvailability - Check if a resource name is available.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
//   - location - The name of Azure region.
//   - body - Name availability request.
//   - options - ResourceClientCheckNameAvailabilityOptions contains the optional parameters for the ResourceClient.CheckNameAvailability
//     method.
func (client *ResourceClient) CheckNameAvailability(ctx context.Context, location string, body ResourceNameAvailabilityRequest, options *ResourceClientCheckNameAvailabilityOptions) (ResourceClientCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, location, body, options)
	if err != nil {
		return ResourceClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ResourceClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ResourceClientCheckNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *ResourceClient) checkNameAvailabilityCreateRequest(ctx context.Context, location string, body ResourceNameAvailabilityRequest, options *ResourceClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.NetApp/locations/{location}/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *ResourceClient) checkNameAvailabilityHandleResponse(resp *http.Response) (ResourceClientCheckNameAvailabilityResponse, error) {
	result := ResourceClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckAvailabilityResponse); err != nil {
		return ResourceClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// CheckQuotaAvailability - Check if a quota is available.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
//   - location - The name of Azure region.
//   - body - Quota availability request.
//   - options - ResourceClientCheckQuotaAvailabilityOptions contains the optional parameters for the ResourceClient.CheckQuotaAvailability
//     method.
func (client *ResourceClient) CheckQuotaAvailability(ctx context.Context, location string, body QuotaAvailabilityRequest, options *ResourceClientCheckQuotaAvailabilityOptions) (ResourceClientCheckQuotaAvailabilityResponse, error) {
	req, err := client.checkQuotaAvailabilityCreateRequest(ctx, location, body, options)
	if err != nil {
		return ResourceClientCheckQuotaAvailabilityResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ResourceClientCheckQuotaAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ResourceClientCheckQuotaAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkQuotaAvailabilityHandleResponse(resp)
}

// checkQuotaAvailabilityCreateRequest creates the CheckQuotaAvailability request.
func (client *ResourceClient) checkQuotaAvailabilityCreateRequest(ctx context.Context, location string, body QuotaAvailabilityRequest, options *ResourceClientCheckQuotaAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.NetApp/locations/{location}/checkQuotaAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// checkQuotaAvailabilityHandleResponse handles the CheckQuotaAvailability response.
func (client *ResourceClient) checkQuotaAvailabilityHandleResponse(resp *http.Response) (ResourceClientCheckQuotaAvailabilityResponse, error) {
	result := ResourceClientCheckQuotaAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckAvailabilityResponse); err != nil {
		return ResourceClientCheckQuotaAvailabilityResponse{}, err
	}
	return result, nil
}

// QueryRegionInfo - Provides storage to network proximity and logical zone mapping information.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
//   - location - The name of Azure region.
//   - options - ResourceClientQueryRegionInfoOptions contains the optional parameters for the ResourceClient.QueryRegionInfo
//     method.
func (client *ResourceClient) QueryRegionInfo(ctx context.Context, location string, options *ResourceClientQueryRegionInfoOptions) (ResourceClientQueryRegionInfoResponse, error) {
	req, err := client.queryRegionInfoCreateRequest(ctx, location, options)
	if err != nil {
		return ResourceClientQueryRegionInfoResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ResourceClientQueryRegionInfoResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ResourceClientQueryRegionInfoResponse{}, runtime.NewResponseError(resp)
	}
	return client.queryRegionInfoHandleResponse(resp)
}

// queryRegionInfoCreateRequest creates the QueryRegionInfo request.
func (client *ResourceClient) queryRegionInfoCreateRequest(ctx context.Context, location string, options *ResourceClientQueryRegionInfoOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.NetApp/locations/{location}/regionInfo"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// queryRegionInfoHandleResponse handles the QueryRegionInfo response.
func (client *ResourceClient) queryRegionInfoHandleResponse(resp *http.Response) (ResourceClientQueryRegionInfoResponse, error) {
	result := ResourceClientQueryRegionInfoResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RegionInfo); err != nil {
		return ResourceClientQueryRegionInfoResponse{}, err
	}
	return result, nil
}
