//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmysql

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

// QueryTextsClient contains the methods for the QueryTexts group.
// Don't use this type directly, use NewQueryTextsClient() instead.
type QueryTextsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewQueryTextsClient creates a new instance of QueryTextsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewQueryTextsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *QueryTextsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &QueryTextsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Get - Retrieve the Query-Store query texts for the queryId.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serverName - The name of the server.
// queryID - The Query-Store query identifier.
// options - QueryTextsClientGetOptions contains the optional parameters for the QueryTextsClient.Get method.
func (client *QueryTextsClient) Get(ctx context.Context, resourceGroupName string, serverName string, queryID string, options *QueryTextsClientGetOptions) (QueryTextsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, queryID, options)
	if err != nil {
		return QueryTextsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return QueryTextsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return QueryTextsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *QueryTextsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, queryID string, options *QueryTextsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/queryTexts/{queryId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if queryID == "" {
		return nil, errors.New("parameter queryID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queryId}", url.PathEscape(queryID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *QueryTextsClient) getHandleResponse(resp *http.Response) (QueryTextsClientGetResponse, error) {
	result := QueryTextsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.QueryText); err != nil {
		return QueryTextsClientGetResponse{}, err
	}
	return result, nil
}

// ListByServer - Retrieve the Query-Store query texts for specified queryIds.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serverName - The name of the server.
// queryIDs - The query identifiers
// options - QueryTextsClientListByServerOptions contains the optional parameters for the QueryTextsClient.ListByServer method.
func (client *QueryTextsClient) ListByServer(resourceGroupName string, serverName string, queryIDs []string, options *QueryTextsClientListByServerOptions) *QueryTextsClientListByServerPager {
	return &QueryTextsClientListByServerPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByServerCreateRequest(ctx, resourceGroupName, serverName, queryIDs, options)
		},
		advancer: func(ctx context.Context, resp QueryTextsClientListByServerResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.QueryTextsResultList.NextLink)
		},
	}
}

// listByServerCreateRequest creates the ListByServer request.
func (client *QueryTextsClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, queryIDs []string, options *QueryTextsClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/queryTexts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	for _, qv := range queryIDs {
		reqQP.Add("queryIds", qv)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *QueryTextsClient) listByServerHandleResponse(resp *http.Response) (QueryTextsClientListByServerResponse, error) {
	result := QueryTextsClientListByServerResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.QueryTextsResultList); err != nil {
		return QueryTextsClientListByServerResponse{}, err
	}
	return result, nil
}
