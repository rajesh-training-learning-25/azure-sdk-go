//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicefabric

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

// ClusterVersionsClient contains the methods for the ClusterVersions group.
// Don't use this type directly, use NewClusterVersionsClient() instead.
type ClusterVersionsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewClusterVersionsClient creates a new instance of ClusterVersionsClient with the specified values.
func NewClusterVersionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ClusterVersionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ClusterVersionsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Gets information about an available Service Fabric cluster code version.
// If the operation fails it returns the *ErrorModel error type.
func (client *ClusterVersionsClient) Get(ctx context.Context, location string, clusterVersion string, options *ClusterVersionsGetOptions) (ClusterVersionsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, location, clusterVersion, options)
	if err != nil {
		return ClusterVersionsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClusterVersionsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClusterVersionsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ClusterVersionsClient) getCreateRequest(ctx context.Context, location string, clusterVersion string, options *ClusterVersionsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/clusterVersions/{clusterVersion}"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if clusterVersion == "" {
		return nil, errors.New("parameter clusterVersion cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterVersion}", url.PathEscape(clusterVersion))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ClusterVersionsClient) getHandleResponse(resp *http.Response) (ClusterVersionsGetResponse, error) {
	result := ClusterVersionsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterCodeVersionsListResult); err != nil {
		return ClusterVersionsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ClusterVersionsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorModel{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetByEnvironment - Gets information about an available Service Fabric cluster code version by environment.
// If the operation fails it returns the *ErrorModel error type.
func (client *ClusterVersionsClient) GetByEnvironment(ctx context.Context, location string, environment Enum14, clusterVersion string, options *ClusterVersionsGetByEnvironmentOptions) (ClusterVersionsGetByEnvironmentResponse, error) {
	req, err := client.getByEnvironmentCreateRequest(ctx, location, environment, clusterVersion, options)
	if err != nil {
		return ClusterVersionsGetByEnvironmentResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClusterVersionsGetByEnvironmentResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClusterVersionsGetByEnvironmentResponse{}, client.getByEnvironmentHandleError(resp)
	}
	return client.getByEnvironmentHandleResponse(resp)
}

// getByEnvironmentCreateRequest creates the GetByEnvironment request.
func (client *ClusterVersionsClient) getByEnvironmentCreateRequest(ctx context.Context, location string, environment Enum14, clusterVersion string, options *ClusterVersionsGetByEnvironmentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/environments/{environment}/clusterVersions/{clusterVersion}"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if environment == "" {
		return nil, errors.New("parameter environment cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environment}", url.PathEscape(string(environment)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if clusterVersion == "" {
		return nil, errors.New("parameter clusterVersion cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterVersion}", url.PathEscape(clusterVersion))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getByEnvironmentHandleResponse handles the GetByEnvironment response.
func (client *ClusterVersionsClient) getByEnvironmentHandleResponse(resp *http.Response) (ClusterVersionsGetByEnvironmentResponse, error) {
	result := ClusterVersionsGetByEnvironmentResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterCodeVersionsListResult); err != nil {
		return ClusterVersionsGetByEnvironmentResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getByEnvironmentHandleError handles the GetByEnvironment error response.
func (client *ClusterVersionsClient) getByEnvironmentHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorModel{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Gets all available code versions for Service Fabric cluster resources by location.
// If the operation fails it returns the *ErrorModel error type.
func (client *ClusterVersionsClient) List(ctx context.Context, location string, options *ClusterVersionsListOptions) (ClusterVersionsListResponse, error) {
	req, err := client.listCreateRequest(ctx, location, options)
	if err != nil {
		return ClusterVersionsListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClusterVersionsListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClusterVersionsListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ClusterVersionsClient) listCreateRequest(ctx context.Context, location string, options *ClusterVersionsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/clusterVersions"
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
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ClusterVersionsClient) listHandleResponse(resp *http.Response) (ClusterVersionsListResponse, error) {
	result := ClusterVersionsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterCodeVersionsListResult); err != nil {
		return ClusterVersionsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ClusterVersionsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorModel{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByEnvironment - Gets all available code versions for Service Fabric cluster resources by environment.
// If the operation fails it returns the *ErrorModel error type.
func (client *ClusterVersionsClient) ListByEnvironment(ctx context.Context, location string, environment Enum14, options *ClusterVersionsListByEnvironmentOptions) (ClusterVersionsListByEnvironmentResponse, error) {
	req, err := client.listByEnvironmentCreateRequest(ctx, location, environment, options)
	if err != nil {
		return ClusterVersionsListByEnvironmentResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClusterVersionsListByEnvironmentResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClusterVersionsListByEnvironmentResponse{}, client.listByEnvironmentHandleError(resp)
	}
	return client.listByEnvironmentHandleResponse(resp)
}

// listByEnvironmentCreateRequest creates the ListByEnvironment request.
func (client *ClusterVersionsClient) listByEnvironmentCreateRequest(ctx context.Context, location string, environment Enum14, options *ClusterVersionsListByEnvironmentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/environments/{environment}/clusterVersions"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if environment == "" {
		return nil, errors.New("parameter environment cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environment}", url.PathEscape(string(environment)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByEnvironmentHandleResponse handles the ListByEnvironment response.
func (client *ClusterVersionsClient) listByEnvironmentHandleResponse(resp *http.Response) (ClusterVersionsListByEnvironmentResponse, error) {
	result := ClusterVersionsListByEnvironmentResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterCodeVersionsListResult); err != nil {
		return ClusterVersionsListByEnvironmentResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByEnvironmentHandleError handles the ListByEnvironment error response.
func (client *ClusterVersionsClient) listByEnvironmentHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorModel{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
