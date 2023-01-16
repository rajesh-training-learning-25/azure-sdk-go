//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork

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

// LoadBalancerBackendAddressPoolsClient contains the methods for the LoadBalancerBackendAddressPools group.
// Don't use this type directly, use NewLoadBalancerBackendAddressPoolsClient() instead.
type LoadBalancerBackendAddressPoolsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewLoadBalancerBackendAddressPoolsClient creates a new instance of LoadBalancerBackendAddressPoolsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewLoadBalancerBackendAddressPoolsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LoadBalancerBackendAddressPoolsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &LoadBalancerBackendAddressPoolsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a load balancer backend address pool.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01
// resourceGroupName - The name of the resource group.
// loadBalancerName - The name of the load balancer.
// backendAddressPoolName - The name of the backend address pool.
// parameters - Parameters supplied to the create or update load balancer backend address pool operation.
// options - LoadBalancerBackendAddressPoolsClientBeginCreateOrUpdateOptions contains the optional parameters for the LoadBalancerBackendAddressPoolsClient.BeginCreateOrUpdate
// method.
func (client *LoadBalancerBackendAddressPoolsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, loadBalancerName string, backendAddressPoolName string, parameters BackendAddressPool, options *LoadBalancerBackendAddressPoolsClientBeginCreateOrUpdateOptions) (*runtime.Poller[LoadBalancerBackendAddressPoolsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, loadBalancerName, backendAddressPoolName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[LoadBalancerBackendAddressPoolsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[LoadBalancerBackendAddressPoolsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a load balancer backend address pool.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01
func (client *LoadBalancerBackendAddressPoolsClient) createOrUpdate(ctx context.Context, resourceGroupName string, loadBalancerName string, backendAddressPoolName string, parameters BackendAddressPool, options *LoadBalancerBackendAddressPoolsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, loadBalancerName, backendAddressPoolName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *LoadBalancerBackendAddressPoolsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, backendAddressPoolName string, parameters BackendAddressPool, options *LoadBalancerBackendAddressPoolsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/backendAddressPools/{backendAddressPoolName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if backendAddressPoolName == "" {
		return nil, errors.New("parameter backendAddressPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backendAddressPoolName}", url.PathEscape(backendAddressPoolName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified load balancer backend address pool.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01
// resourceGroupName - The name of the resource group.
// loadBalancerName - The name of the load balancer.
// backendAddressPoolName - The name of the backend address pool.
// options - LoadBalancerBackendAddressPoolsClientBeginDeleteOptions contains the optional parameters for the LoadBalancerBackendAddressPoolsClient.BeginDelete
// method.
func (client *LoadBalancerBackendAddressPoolsClient) BeginDelete(ctx context.Context, resourceGroupName string, loadBalancerName string, backendAddressPoolName string, options *LoadBalancerBackendAddressPoolsClientBeginDeleteOptions) (*runtime.Poller[LoadBalancerBackendAddressPoolsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, loadBalancerName, backendAddressPoolName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[LoadBalancerBackendAddressPoolsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[LoadBalancerBackendAddressPoolsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the specified load balancer backend address pool.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01
func (client *LoadBalancerBackendAddressPoolsClient) deleteOperation(ctx context.Context, resourceGroupName string, loadBalancerName string, backendAddressPoolName string, options *LoadBalancerBackendAddressPoolsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, loadBalancerName, backendAddressPoolName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *LoadBalancerBackendAddressPoolsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, backendAddressPoolName string, options *LoadBalancerBackendAddressPoolsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/backendAddressPools/{backendAddressPoolName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if backendAddressPoolName == "" {
		return nil, errors.New("parameter backendAddressPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backendAddressPoolName}", url.PathEscape(backendAddressPoolName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets load balancer backend address pool.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01
// resourceGroupName - The name of the resource group.
// loadBalancerName - The name of the load balancer.
// backendAddressPoolName - The name of the backend address pool.
// options - LoadBalancerBackendAddressPoolsClientGetOptions contains the optional parameters for the LoadBalancerBackendAddressPoolsClient.Get
// method.
func (client *LoadBalancerBackendAddressPoolsClient) Get(ctx context.Context, resourceGroupName string, loadBalancerName string, backendAddressPoolName string, options *LoadBalancerBackendAddressPoolsClientGetOptions) (LoadBalancerBackendAddressPoolsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, loadBalancerName, backendAddressPoolName, options)
	if err != nil {
		return LoadBalancerBackendAddressPoolsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LoadBalancerBackendAddressPoolsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LoadBalancerBackendAddressPoolsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LoadBalancerBackendAddressPoolsClient) getCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, backendAddressPoolName string, options *LoadBalancerBackendAddressPoolsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/backendAddressPools/{backendAddressPoolName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if backendAddressPoolName == "" {
		return nil, errors.New("parameter backendAddressPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backendAddressPoolName}", url.PathEscape(backendAddressPoolName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LoadBalancerBackendAddressPoolsClient) getHandleResponse(resp *http.Response) (LoadBalancerBackendAddressPoolsClientGetResponse, error) {
	result := LoadBalancerBackendAddressPoolsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BackendAddressPool); err != nil {
		return LoadBalancerBackendAddressPoolsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all the load balancer backed address pools.
// Generated from API version 2022-07-01
// resourceGroupName - The name of the resource group.
// loadBalancerName - The name of the load balancer.
// options - LoadBalancerBackendAddressPoolsClientListOptions contains the optional parameters for the LoadBalancerBackendAddressPoolsClient.List
// method.
func (client *LoadBalancerBackendAddressPoolsClient) NewListPager(resourceGroupName string, loadBalancerName string, options *LoadBalancerBackendAddressPoolsClientListOptions) *runtime.Pager[LoadBalancerBackendAddressPoolsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[LoadBalancerBackendAddressPoolsClientListResponse]{
		More: func(page LoadBalancerBackendAddressPoolsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LoadBalancerBackendAddressPoolsClientListResponse) (LoadBalancerBackendAddressPoolsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, loadBalancerName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return LoadBalancerBackendAddressPoolsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return LoadBalancerBackendAddressPoolsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LoadBalancerBackendAddressPoolsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *LoadBalancerBackendAddressPoolsClient) listCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancerBackendAddressPoolsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/backendAddressPools"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LoadBalancerBackendAddressPoolsClient) listHandleResponse(resp *http.Response) (LoadBalancerBackendAddressPoolsClientListResponse, error) {
	result := LoadBalancerBackendAddressPoolsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LoadBalancerBackendAddressPoolListResult); err != nil {
		return LoadBalancerBackendAddressPoolsClientListResponse{}, err
	}
	return result, nil
}
