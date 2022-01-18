//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// ServerDNSAliasesClient contains the methods for the ServerDNSAliases group.
// Don't use this type directly, use NewServerDNSAliasesClient() instead.
type ServerDNSAliasesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewServerDNSAliasesClient creates a new instance of ServerDNSAliasesClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewServerDNSAliasesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ServerDNSAliasesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &ServerDNSAliasesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginAcquire - Acquires server DNS alias from another server.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server that the alias is pointing to.
// dnsAliasName - The name of the server dns alias.
// options - ServerDNSAliasesClientBeginAcquireOptions contains the optional parameters for the ServerDNSAliasesClient.BeginAcquire
// method.
func (client *ServerDNSAliasesClient) BeginAcquire(ctx context.Context, resourceGroupName string, serverName string, dnsAliasName string, parameters ServerDNSAliasAcquisition, options *ServerDNSAliasesClientBeginAcquireOptions) (ServerDNSAliasesClientAcquirePollerResponse, error) {
	resp, err := client.acquire(ctx, resourceGroupName, serverName, dnsAliasName, parameters, options)
	if err != nil {
		return ServerDNSAliasesClientAcquirePollerResponse{}, err
	}
	result := ServerDNSAliasesClientAcquirePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ServerDNSAliasesClient.Acquire", "", resp, client.pl)
	if err != nil {
		return ServerDNSAliasesClientAcquirePollerResponse{}, err
	}
	result.Poller = &ServerDNSAliasesClientAcquirePoller{
		pt: pt,
	}
	return result, nil
}

// Acquire - Acquires server DNS alias from another server.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ServerDNSAliasesClient) acquire(ctx context.Context, resourceGroupName string, serverName string, dnsAliasName string, parameters ServerDNSAliasAcquisition, options *ServerDNSAliasesClientBeginAcquireOptions) (*http.Response, error) {
	req, err := client.acquireCreateRequest(ctx, resourceGroupName, serverName, dnsAliasName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// acquireCreateRequest creates the Acquire request.
func (client *ServerDNSAliasesClient) acquireCreateRequest(ctx context.Context, resourceGroupName string, serverName string, dnsAliasName string, parameters ServerDNSAliasAcquisition, options *ServerDNSAliasesClientBeginAcquireOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/dnsAliases/{dnsAliasName}/acquire"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if dnsAliasName == "" {
		return nil, errors.New("parameter dnsAliasName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dnsAliasName}", url.PathEscape(dnsAliasName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginCreateOrUpdate - Creates a server DNS alias.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server that the alias is pointing to.
// dnsAliasName - The name of the server dns alias.
// options - ServerDNSAliasesClientBeginCreateOrUpdateOptions contains the optional parameters for the ServerDNSAliasesClient.BeginCreateOrUpdate
// method.
func (client *ServerDNSAliasesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, dnsAliasName string, options *ServerDNSAliasesClientBeginCreateOrUpdateOptions) (ServerDNSAliasesClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, dnsAliasName, options)
	if err != nil {
		return ServerDNSAliasesClientCreateOrUpdatePollerResponse{}, err
	}
	result := ServerDNSAliasesClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ServerDNSAliasesClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return ServerDNSAliasesClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ServerDNSAliasesClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates a server DNS alias.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ServerDNSAliasesClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, dnsAliasName string, options *ServerDNSAliasesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, dnsAliasName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServerDNSAliasesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, dnsAliasName string, options *ServerDNSAliasesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/dnsAliases/{dnsAliasName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if dnsAliasName == "" {
		return nil, errors.New("parameter dnsAliasName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dnsAliasName}", url.PathEscape(dnsAliasName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginDelete - Deletes the server DNS alias with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server that the alias is pointing to.
// dnsAliasName - The name of the server dns alias.
// options - ServerDNSAliasesClientBeginDeleteOptions contains the optional parameters for the ServerDNSAliasesClient.BeginDelete
// method.
func (client *ServerDNSAliasesClient) BeginDelete(ctx context.Context, resourceGroupName string, serverName string, dnsAliasName string, options *ServerDNSAliasesClientBeginDeleteOptions) (ServerDNSAliasesClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, serverName, dnsAliasName, options)
	if err != nil {
		return ServerDNSAliasesClientDeletePollerResponse{}, err
	}
	result := ServerDNSAliasesClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ServerDNSAliasesClient.Delete", "", resp, client.pl)
	if err != nil {
		return ServerDNSAliasesClientDeletePollerResponse{}, err
	}
	result.Poller = &ServerDNSAliasesClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the server DNS alias with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ServerDNSAliasesClient) deleteOperation(ctx context.Context, resourceGroupName string, serverName string, dnsAliasName string, options *ServerDNSAliasesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, dnsAliasName, options)
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
func (client *ServerDNSAliasesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, dnsAliasName string, options *ServerDNSAliasesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/dnsAliases/{dnsAliasName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if dnsAliasName == "" {
		return nil, errors.New("parameter dnsAliasName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dnsAliasName}", url.PathEscape(dnsAliasName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets a server DNS alias.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server that the alias is pointing to.
// dnsAliasName - The name of the server dns alias.
// options - ServerDNSAliasesClientGetOptions contains the optional parameters for the ServerDNSAliasesClient.Get method.
func (client *ServerDNSAliasesClient) Get(ctx context.Context, resourceGroupName string, serverName string, dnsAliasName string, options *ServerDNSAliasesClientGetOptions) (ServerDNSAliasesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, dnsAliasName, options)
	if err != nil {
		return ServerDNSAliasesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServerDNSAliasesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServerDNSAliasesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServerDNSAliasesClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, dnsAliasName string, options *ServerDNSAliasesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/dnsAliases/{dnsAliasName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if dnsAliasName == "" {
		return nil, errors.New("parameter dnsAliasName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dnsAliasName}", url.PathEscape(dnsAliasName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServerDNSAliasesClient) getHandleResponse(resp *http.Response) (ServerDNSAliasesClientGetResponse, error) {
	result := ServerDNSAliasesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerDNSAlias); err != nil {
		return ServerDNSAliasesClientGetResponse{}, err
	}
	return result, nil
}

// ListByServer - Gets a list of server DNS aliases for a server.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server that the alias is pointing to.
// options - ServerDNSAliasesClientListByServerOptions contains the optional parameters for the ServerDNSAliasesClient.ListByServer
// method.
func (client *ServerDNSAliasesClient) ListByServer(resourceGroupName string, serverName string, options *ServerDNSAliasesClientListByServerOptions) *ServerDNSAliasesClientListByServerPager {
	return &ServerDNSAliasesClientListByServerPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
		},
		advancer: func(ctx context.Context, resp ServerDNSAliasesClientListByServerResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ServerDNSAliasListResult.NextLink)
		},
	}
}

// listByServerCreateRequest creates the ListByServer request.
func (client *ServerDNSAliasesClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerDNSAliasesClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/dnsAliases"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *ServerDNSAliasesClient) listByServerHandleResponse(resp *http.Response) (ServerDNSAliasesClientListByServerResponse, error) {
	result := ServerDNSAliasesClientListByServerResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerDNSAliasListResult); err != nil {
		return ServerDNSAliasesClientListByServerResponse{}, err
	}
	return result, nil
}
