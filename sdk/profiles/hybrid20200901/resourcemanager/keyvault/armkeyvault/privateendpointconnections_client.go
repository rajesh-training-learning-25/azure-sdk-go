//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armkeyvault

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/profiles/hybrid20200901/internal"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// PrivateEndpointConnectionsClient contains the methods for the PrivateEndpointConnections group.
// Don't use this type directly, use NewPrivateEndpointConnectionsClient() instead.
type PrivateEndpointConnectionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPrivateEndpointConnectionsClient creates a new instance of PrivateEndpointConnectionsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPrivateEndpointConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateEndpointConnectionsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(internal.ModuleName, internal.ModuleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &PrivateEndpointConnectionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginDelete - Deletes the specified private endpoint connection associated with the key vault.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-09-01
// resourceGroupName - Name of the resource group that contains the key vault.
// vaultName - The name of the key vault.
// privateEndpointConnectionName - Name of the private endpoint connection associated with the key vault.
// options - PrivateEndpointConnectionsClientBeginDeleteOptions contains the optional parameters for the PrivateEndpointConnectionsClient.BeginDelete
// method.
func (client *PrivateEndpointConnectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, vaultName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsClientBeginDeleteOptions) (*runtime.Poller[PrivateEndpointConnectionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, vaultName, privateEndpointConnectionName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[PrivateEndpointConnectionsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[PrivateEndpointConnectionsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the specified private endpoint connection associated with the key vault.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-09-01
func (client *PrivateEndpointConnectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, vaultName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, vaultName, privateEndpointConnectionName, options)
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
func (client *PrivateEndpointConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the specified private endpoint connection associated with the key vault.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-09-01
// resourceGroupName - Name of the resource group that contains the key vault.
// vaultName - The name of the key vault.
// privateEndpointConnectionName - Name of the private endpoint connection associated with the key vault.
// options - PrivateEndpointConnectionsClientGetOptions contains the optional parameters for the PrivateEndpointConnectionsClient.Get
// method.
func (client *PrivateEndpointConnectionsClient) Get(ctx context.Context, resourceGroupName string, vaultName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsClientGetOptions) (PrivateEndpointConnectionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, vaultName, privateEndpointConnectionName, options)
	if err != nil {
		return PrivateEndpointConnectionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateEndpointConnectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return PrivateEndpointConnectionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateEndpointConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateEndpointConnectionsClient) getHandleResponse(resp *http.Response) (PrivateEndpointConnectionsClientGetResponse, error) {
	result := PrivateEndpointConnectionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnection); err != nil {
		return PrivateEndpointConnectionsClientGetResponse{}, err
	}
	return result, nil
}

// Put - Updates the specified private endpoint connection associated with the key vault.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-09-01
// resourceGroupName - Name of the resource group that contains the key vault.
// vaultName - The name of the key vault.
// privateEndpointConnectionName - Name of the private endpoint connection associated with the key vault.
// properties - The intended state of private endpoint connection.
// options - PrivateEndpointConnectionsClientPutOptions contains the optional parameters for the PrivateEndpointConnectionsClient.Put
// method.
func (client *PrivateEndpointConnectionsClient) Put(ctx context.Context, resourceGroupName string, vaultName string, privateEndpointConnectionName string, properties PrivateEndpointConnection, options *PrivateEndpointConnectionsClientPutOptions) (PrivateEndpointConnectionsClientPutResponse, error) {
	req, err := client.putCreateRequest(ctx, resourceGroupName, vaultName, privateEndpointConnectionName, properties, options)
	if err != nil {
		return PrivateEndpointConnectionsClientPutResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateEndpointConnectionsClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateEndpointConnectionsClientPutResponse{}, runtime.NewResponseError(resp)
	}
	return client.putHandleResponse(resp)
}

// putCreateRequest creates the Put request.
func (client *PrivateEndpointConnectionsClient) putCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, privateEndpointConnectionName string, properties PrivateEndpointConnection, options *PrivateEndpointConnectionsClientPutOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, properties)
}

// putHandleResponse handles the Put response.
func (client *PrivateEndpointConnectionsClient) putHandleResponse(resp *http.Response) (PrivateEndpointConnectionsClientPutResponse, error) {
	result := PrivateEndpointConnectionsClientPutResponse{}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return PrivateEndpointConnectionsClientPutResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	if val := resp.Header.Get("Azure-AsyncOperation"); val != "" {
		result.AzureAsyncOperation = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnection); err != nil {
		return PrivateEndpointConnectionsClientPutResponse{}, err
	}
	return result, nil
}
