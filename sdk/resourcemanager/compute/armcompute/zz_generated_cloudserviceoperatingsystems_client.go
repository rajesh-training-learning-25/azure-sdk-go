//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

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

// CloudServiceOperatingSystemsClient contains the methods for the CloudServiceOperatingSystems group.
// Don't use this type directly, use NewCloudServiceOperatingSystemsClient() instead.
type CloudServiceOperatingSystemsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewCloudServiceOperatingSystemsClient creates a new instance of CloudServiceOperatingSystemsClient with the specified values.
func NewCloudServiceOperatingSystemsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *CloudServiceOperatingSystemsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &CloudServiceOperatingSystemsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// GetOSFamily - Gets properties of a guest operating system family that can be specified in the XML service configuration (.cscfg) for a cloud service.
// If the operation fails it returns the *CloudError error type.
func (client *CloudServiceOperatingSystemsClient) GetOSFamily(ctx context.Context, location string, osFamilyName string, options *CloudServiceOperatingSystemsGetOSFamilyOptions) (CloudServiceOperatingSystemsGetOSFamilyResponse, error) {
	req, err := client.getOSFamilyCreateRequest(ctx, location, osFamilyName, options)
	if err != nil {
		return CloudServiceOperatingSystemsGetOSFamilyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CloudServiceOperatingSystemsGetOSFamilyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CloudServiceOperatingSystemsGetOSFamilyResponse{}, client.getOSFamilyHandleError(resp)
	}
	return client.getOSFamilyHandleResponse(resp)
}

// getOSFamilyCreateRequest creates the GetOSFamily request.
func (client *CloudServiceOperatingSystemsClient) getOSFamilyCreateRequest(ctx context.Context, location string, osFamilyName string, options *CloudServiceOperatingSystemsGetOSFamilyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/cloudServiceOsFamilies/{osFamilyName}"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if osFamilyName == "" {
		return nil, errors.New("parameter osFamilyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{osFamilyName}", url.PathEscape(osFamilyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getOSFamilyHandleResponse handles the GetOSFamily response.
func (client *CloudServiceOperatingSystemsClient) getOSFamilyHandleResponse(resp *http.Response) (CloudServiceOperatingSystemsGetOSFamilyResponse, error) {
	result := CloudServiceOperatingSystemsGetOSFamilyResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.OSFamily); err != nil {
		return CloudServiceOperatingSystemsGetOSFamilyResponse{}, err
	}
	return result, nil
}

// getOSFamilyHandleError handles the GetOSFamily error response.
func (client *CloudServiceOperatingSystemsClient) getOSFamilyHandleError(resp *http.Response) error {
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

// GetOSVersion - Gets properties of a guest operating system version that can be specified in the XML service configuration (.cscfg) for a cloud service.
// If the operation fails it returns the *CloudError error type.
func (client *CloudServiceOperatingSystemsClient) GetOSVersion(ctx context.Context, location string, osVersionName string, options *CloudServiceOperatingSystemsGetOSVersionOptions) (CloudServiceOperatingSystemsGetOSVersionResponse, error) {
	req, err := client.getOSVersionCreateRequest(ctx, location, osVersionName, options)
	if err != nil {
		return CloudServiceOperatingSystemsGetOSVersionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CloudServiceOperatingSystemsGetOSVersionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CloudServiceOperatingSystemsGetOSVersionResponse{}, client.getOSVersionHandleError(resp)
	}
	return client.getOSVersionHandleResponse(resp)
}

// getOSVersionCreateRequest creates the GetOSVersion request.
func (client *CloudServiceOperatingSystemsClient) getOSVersionCreateRequest(ctx context.Context, location string, osVersionName string, options *CloudServiceOperatingSystemsGetOSVersionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/cloudServiceOsVersions/{osVersionName}"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if osVersionName == "" {
		return nil, errors.New("parameter osVersionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{osVersionName}", url.PathEscape(osVersionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getOSVersionHandleResponse handles the GetOSVersion response.
func (client *CloudServiceOperatingSystemsClient) getOSVersionHandleResponse(resp *http.Response) (CloudServiceOperatingSystemsGetOSVersionResponse, error) {
	result := CloudServiceOperatingSystemsGetOSVersionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.OSVersion); err != nil {
		return CloudServiceOperatingSystemsGetOSVersionResponse{}, err
	}
	return result, nil
}

// getOSVersionHandleError handles the GetOSVersion error response.
func (client *CloudServiceOperatingSystemsClient) getOSVersionHandleError(resp *http.Response) error {
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

// ListOSFamilies - Gets a list of all guest operating system families available to be specified in the XML service configuration (.cscfg) for a cloud service.
// Use nextLink property in the response to get the next page
// of OS Families. Do this till nextLink is null to fetch all the OS Families.
// If the operation fails it returns the *CloudError error type.
func (client *CloudServiceOperatingSystemsClient) ListOSFamilies(location string, options *CloudServiceOperatingSystemsListOSFamiliesOptions) *CloudServiceOperatingSystemsListOSFamiliesPager {
	return &CloudServiceOperatingSystemsListOSFamiliesPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listOSFamiliesCreateRequest(ctx, location, options)
		},
		advancer: func(ctx context.Context, resp CloudServiceOperatingSystemsListOSFamiliesResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.OSFamilyListResult.NextLink)
		},
	}
}

// listOSFamiliesCreateRequest creates the ListOSFamilies request.
func (client *CloudServiceOperatingSystemsClient) listOSFamiliesCreateRequest(ctx context.Context, location string, options *CloudServiceOperatingSystemsListOSFamiliesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/cloudServiceOsFamilies"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listOSFamiliesHandleResponse handles the ListOSFamilies response.
func (client *CloudServiceOperatingSystemsClient) listOSFamiliesHandleResponse(resp *http.Response) (CloudServiceOperatingSystemsListOSFamiliesResponse, error) {
	result := CloudServiceOperatingSystemsListOSFamiliesResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.OSFamilyListResult); err != nil {
		return CloudServiceOperatingSystemsListOSFamiliesResponse{}, err
	}
	return result, nil
}

// listOSFamiliesHandleError handles the ListOSFamilies error response.
func (client *CloudServiceOperatingSystemsClient) listOSFamiliesHandleError(resp *http.Response) error {
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

// ListOSVersions - Gets a list of all guest operating system versions available to be specified in the XML service configuration (.cscfg) for a cloud service.
// Use nextLink property in the response to get the next page
// of OS versions. Do this till nextLink is null to fetch all the OS versions.
// If the operation fails it returns the *CloudError error type.
func (client *CloudServiceOperatingSystemsClient) ListOSVersions(location string, options *CloudServiceOperatingSystemsListOSVersionsOptions) *CloudServiceOperatingSystemsListOSVersionsPager {
	return &CloudServiceOperatingSystemsListOSVersionsPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listOSVersionsCreateRequest(ctx, location, options)
		},
		advancer: func(ctx context.Context, resp CloudServiceOperatingSystemsListOSVersionsResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.OSVersionListResult.NextLink)
		},
	}
}

// listOSVersionsCreateRequest creates the ListOSVersions request.
func (client *CloudServiceOperatingSystemsClient) listOSVersionsCreateRequest(ctx context.Context, location string, options *CloudServiceOperatingSystemsListOSVersionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/cloudServiceOsVersions"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listOSVersionsHandleResponse handles the ListOSVersions response.
func (client *CloudServiceOperatingSystemsClient) listOSVersionsHandleResponse(resp *http.Response) (CloudServiceOperatingSystemsListOSVersionsResponse, error) {
	result := CloudServiceOperatingSystemsListOSVersionsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.OSVersionListResult); err != nil {
		return CloudServiceOperatingSystemsListOSVersionsResponse{}, err
	}
	return result, nil
}

// listOSVersionsHandleError handles the ListOSVersions error response.
func (client *CloudServiceOperatingSystemsClient) listOSVersionsHandleError(resp *http.Response) error {
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
