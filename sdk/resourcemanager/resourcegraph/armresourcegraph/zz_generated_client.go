//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcegraph

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// Client contains the methods for the ResourceGraphClient group.
// Don't use this type directly, use NewClient() instead.
type Client struct {
	host string
	pl   runtime.Pipeline
}

// NewClient creates a new instance of Client with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewClient(credential azcore.TokenCredential, options *arm.ClientOptions) *Client {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &Client{
		host: string(cp.Endpoint),
		pl:   armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Resources - Queries the resources managed by Azure Resource Manager for scopes specified in the request.
// If the operation fails it returns an *azcore.ResponseError type.
// query - Request specifying query and its options.
// options - ClientResourcesOptions contains the optional parameters for the Client.Resources method.
func (client *Client) Resources(ctx context.Context, query QueryRequest, options *ClientResourcesOptions) (ClientResourcesResponse, error) {
	req, err := client.resourcesCreateRequest(ctx, query, options)
	if err != nil {
		return ClientResourcesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClientResourcesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClientResourcesResponse{}, runtime.NewResponseError(resp)
	}
	return client.resourcesHandleResponse(resp)
}

// resourcesCreateRequest creates the Resources request.
func (client *Client) resourcesCreateRequest(ctx context.Context, query QueryRequest, options *ClientResourcesOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.ResourceGraph/resources"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, query)
}

// resourcesHandleResponse handles the Resources response.
func (client *Client) resourcesHandleResponse(resp *http.Response) (ClientResourcesResponse, error) {
	result := ClientResourcesResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.QueryResponse); err != nil {
		return ClientResourcesResponse{}, err
	}
	return result, nil
}

// ResourcesHistory - List all snapshots of a resource for a given time interval.
// If the operation fails it returns an *azcore.ResponseError type.
// request - Request specifying the query and its options.
// options - ClientResourcesHistoryOptions contains the optional parameters for the Client.ResourcesHistory method.
func (client *Client) ResourcesHistory(ctx context.Context, request ResourcesHistoryRequest, options *ClientResourcesHistoryOptions) (ClientResourcesHistoryResponse, error) {
	req, err := client.resourcesHistoryCreateRequest(ctx, request, options)
	if err != nil {
		return ClientResourcesHistoryResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClientResourcesHistoryResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClientResourcesHistoryResponse{}, runtime.NewResponseError(resp)
	}
	return client.resourcesHistoryHandleResponse(resp)
}

// resourcesHistoryCreateRequest creates the ResourcesHistory request.
func (client *Client) resourcesHistoryCreateRequest(ctx context.Context, request ResourcesHistoryRequest, options *ClientResourcesHistoryOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.ResourceGraph/resourcesHistory"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, request)
}

// resourcesHistoryHandleResponse handles the ResourcesHistory response.
func (client *Client) resourcesHistoryHandleResponse(resp *http.Response) (ClientResourcesHistoryResponse, error) {
	result := ClientResourcesHistoryResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Object); err != nil {
		return ClientResourcesHistoryResponse{}, err
	}
	return result, nil
}
