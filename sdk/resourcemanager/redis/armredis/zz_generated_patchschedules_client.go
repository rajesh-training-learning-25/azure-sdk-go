//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armredis

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// PatchSchedulesClient contains the methods for the PatchSchedules group.
// Don't use this type directly, use NewPatchSchedulesClient() instead.
type PatchSchedulesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPatchSchedulesClient creates a new instance of PatchSchedulesClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPatchSchedulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PatchSchedulesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &PatchSchedulesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Create or replace the patching schedule for Redis cache.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// name - The name of the Redis cache.
// defaultParam - Default string modeled as parameter for auto generation to work correctly.
// parameters - Parameters to set the patching schedule for Redis cache.
// options - PatchSchedulesClientCreateOrUpdateOptions contains the optional parameters for the PatchSchedulesClient.CreateOrUpdate
// method.
func (client *PatchSchedulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, name string, defaultParam DefaultName, parameters PatchSchedule, options *PatchSchedulesClientCreateOrUpdateOptions) (PatchSchedulesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, name, defaultParam, parameters, options)
	if err != nil {
		return PatchSchedulesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PatchSchedulesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return PatchSchedulesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PatchSchedulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, name string, defaultParam DefaultName, parameters PatchSchedule, options *PatchSchedulesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/patchSchedules/{default}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if defaultParam == "" {
		return nil, errors.New("parameter defaultParam cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{default}", url.PathEscape(string(defaultParam)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *PatchSchedulesClient) createOrUpdateHandleResponse(resp *http.Response) (PatchSchedulesClientCreateOrUpdateResponse, error) {
	result := PatchSchedulesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PatchSchedule); err != nil {
		return PatchSchedulesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the patching schedule of a redis cache.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// name - The name of the redis cache.
// defaultParam - Default string modeled as parameter for auto generation to work correctly.
// options - PatchSchedulesClientDeleteOptions contains the optional parameters for the PatchSchedulesClient.Delete method.
func (client *PatchSchedulesClient) Delete(ctx context.Context, resourceGroupName string, name string, defaultParam DefaultName, options *PatchSchedulesClientDeleteOptions) (PatchSchedulesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, name, defaultParam, options)
	if err != nil {
		return PatchSchedulesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PatchSchedulesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return PatchSchedulesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return PatchSchedulesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PatchSchedulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, name string, defaultParam DefaultName, options *PatchSchedulesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/patchSchedules/{default}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if defaultParam == "" {
		return nil, errors.New("parameter defaultParam cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{default}", url.PathEscape(string(defaultParam)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the patching schedule of a redis cache.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// name - The name of the redis cache.
// defaultParam - Default string modeled as parameter for auto generation to work correctly.
// options - PatchSchedulesClientGetOptions contains the optional parameters for the PatchSchedulesClient.Get method.
func (client *PatchSchedulesClient) Get(ctx context.Context, resourceGroupName string, name string, defaultParam DefaultName, options *PatchSchedulesClientGetOptions) (PatchSchedulesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, name, defaultParam, options)
	if err != nil {
		return PatchSchedulesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PatchSchedulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PatchSchedulesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PatchSchedulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, name string, defaultParam DefaultName, options *PatchSchedulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/patchSchedules/{default}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if defaultParam == "" {
		return nil, errors.New("parameter defaultParam cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{default}", url.PathEscape(string(defaultParam)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
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
func (client *PatchSchedulesClient) getHandleResponse(resp *http.Response) (PatchSchedulesClientGetResponse, error) {
	result := PatchSchedulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PatchSchedule); err != nil {
		return PatchSchedulesClientGetResponse{}, err
	}
	return result, nil
}

// ListByRedisResource - Gets all patch schedules in the specified redis cache (there is only one).
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// cacheName - The name of the Redis cache.
// options - PatchSchedulesClientListByRedisResourceOptions contains the optional parameters for the PatchSchedulesClient.ListByRedisResource
// method.
func (client *PatchSchedulesClient) ListByRedisResource(resourceGroupName string, cacheName string, options *PatchSchedulesClientListByRedisResourceOptions) *runtime.Pager[PatchSchedulesClientListByRedisResourceResponse] {
	return runtime.NewPager(runtime.PageProcessor[PatchSchedulesClientListByRedisResourceResponse]{
		More: func(page PatchSchedulesClientListByRedisResourceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PatchSchedulesClientListByRedisResourceResponse) (PatchSchedulesClientListByRedisResourceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByRedisResourceCreateRequest(ctx, resourceGroupName, cacheName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PatchSchedulesClientListByRedisResourceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PatchSchedulesClientListByRedisResourceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PatchSchedulesClientListByRedisResourceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByRedisResourceHandleResponse(resp)
		},
	})
}

// listByRedisResourceCreateRequest creates the ListByRedisResource request.
func (client *PatchSchedulesClient) listByRedisResourceCreateRequest(ctx context.Context, resourceGroupName string, cacheName string, options *PatchSchedulesClientListByRedisResourceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{cacheName}/patchSchedules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cacheName == "" {
		return nil, errors.New("parameter cacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cacheName}", url.PathEscape(cacheName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByRedisResourceHandleResponse handles the ListByRedisResource response.
func (client *PatchSchedulesClient) listByRedisResourceHandleResponse(resp *http.Response) (PatchSchedulesClientListByRedisResourceResponse, error) {
	result := PatchSchedulesClientListByRedisResourceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PatchScheduleListResult); err != nil {
		return PatchSchedulesClientListByRedisResourceResponse{}, err
	}
	return result, nil
}
