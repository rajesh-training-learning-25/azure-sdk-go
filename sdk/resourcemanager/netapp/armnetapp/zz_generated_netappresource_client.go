//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetapp

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

// NetAppResourceClient contains the methods for the NetAppResource group.
// Don't use this type directly, use NewNetAppResourceClient() instead.
type NetAppResourceClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewNetAppResourceClient creates a new instance of NetAppResourceClient with the specified values.
func NewNetAppResourceClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *NetAppResourceClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &NetAppResourceClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CheckFilePathAvailability - Check if a file path is available.
// If the operation fails it returns a generic error.
func (client *NetAppResourceClient) CheckFilePathAvailability(ctx context.Context, location string, body FilePathAvailabilityRequest, options *NetAppResourceCheckFilePathAvailabilityOptions) (NetAppResourceCheckFilePathAvailabilityResponse, error) {
	req, err := client.checkFilePathAvailabilityCreateRequest(ctx, location, body, options)
	if err != nil {
		return NetAppResourceCheckFilePathAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NetAppResourceCheckFilePathAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NetAppResourceCheckFilePathAvailabilityResponse{}, client.checkFilePathAvailabilityHandleError(resp)
	}
	return client.checkFilePathAvailabilityHandleResponse(resp)
}

// checkFilePathAvailabilityCreateRequest creates the CheckFilePathAvailability request.
func (client *NetAppResourceClient) checkFilePathAvailabilityCreateRequest(ctx context.Context, location string, body FilePathAvailabilityRequest, options *NetAppResourceCheckFilePathAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.NetApp/locations/{location}/checkFilePathAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// checkFilePathAvailabilityHandleResponse handles the CheckFilePathAvailability response.
func (client *NetAppResourceClient) checkFilePathAvailabilityHandleResponse(resp *http.Response) (NetAppResourceCheckFilePathAvailabilityResponse, error) {
	result := NetAppResourceCheckFilePathAvailabilityResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckAvailabilityResponse); err != nil {
		return NetAppResourceCheckFilePathAvailabilityResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// checkFilePathAvailabilityHandleError handles the CheckFilePathAvailability error response.
func (client *NetAppResourceClient) checkFilePathAvailabilityHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// CheckNameAvailability - Check if a resource name is available.
// If the operation fails it returns a generic error.
func (client *NetAppResourceClient) CheckNameAvailability(ctx context.Context, location string, body ResourceNameAvailabilityRequest, options *NetAppResourceCheckNameAvailabilityOptions) (NetAppResourceCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, location, body, options)
	if err != nil {
		return NetAppResourceCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NetAppResourceCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NetAppResourceCheckNameAvailabilityResponse{}, client.checkNameAvailabilityHandleError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *NetAppResourceClient) checkNameAvailabilityCreateRequest(ctx context.Context, location string, body ResourceNameAvailabilityRequest, options *NetAppResourceCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.NetApp/locations/{location}/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *NetAppResourceClient) checkNameAvailabilityHandleResponse(resp *http.Response) (NetAppResourceCheckNameAvailabilityResponse, error) {
	result := NetAppResourceCheckNameAvailabilityResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckAvailabilityResponse); err != nil {
		return NetAppResourceCheckNameAvailabilityResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// checkNameAvailabilityHandleError handles the CheckNameAvailability error response.
func (client *NetAppResourceClient) checkNameAvailabilityHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// CheckQuotaAvailability - Check if a quota is available.
// If the operation fails it returns a generic error.
func (client *NetAppResourceClient) CheckQuotaAvailability(ctx context.Context, location string, body QuotaAvailabilityRequest, options *NetAppResourceCheckQuotaAvailabilityOptions) (NetAppResourceCheckQuotaAvailabilityResponse, error) {
	req, err := client.checkQuotaAvailabilityCreateRequest(ctx, location, body, options)
	if err != nil {
		return NetAppResourceCheckQuotaAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NetAppResourceCheckQuotaAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NetAppResourceCheckQuotaAvailabilityResponse{}, client.checkQuotaAvailabilityHandleError(resp)
	}
	return client.checkQuotaAvailabilityHandleResponse(resp)
}

// checkQuotaAvailabilityCreateRequest creates the CheckQuotaAvailability request.
func (client *NetAppResourceClient) checkQuotaAvailabilityCreateRequest(ctx context.Context, location string, body QuotaAvailabilityRequest, options *NetAppResourceCheckQuotaAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.NetApp/locations/{location}/checkQuotaAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// checkQuotaAvailabilityHandleResponse handles the CheckQuotaAvailability response.
func (client *NetAppResourceClient) checkQuotaAvailabilityHandleResponse(resp *http.Response) (NetAppResourceCheckQuotaAvailabilityResponse, error) {
	result := NetAppResourceCheckQuotaAvailabilityResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckAvailabilityResponse); err != nil {
		return NetAppResourceCheckQuotaAvailabilityResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// checkQuotaAvailabilityHandleError handles the CheckQuotaAvailability error response.
func (client *NetAppResourceClient) checkQuotaAvailabilityHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
