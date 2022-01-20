//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurestackhci

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

// ExtensionsClient contains the methods for the Extensions group.
// Don't use this type directly, use NewExtensionsClient() instead.
type ExtensionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewExtensionsClient creates a new instance of ExtensionsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewExtensionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ExtensionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &ExtensionsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreate - Create Extension for HCI cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// clusterName - The name of the cluster.
// arcSettingName - The name of the proxy resource holding details of HCI ArcSetting information.
// extensionName - The name of the machine extension.
// extension - Details of the Machine Extension to be created.
// options - ExtensionsClientBeginCreateOptions contains the optional parameters for the ExtensionsClient.BeginCreate method.
func (client *ExtensionsClient) BeginCreate(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, extensionName string, extension Extension, options *ExtensionsClientBeginCreateOptions) (ExtensionsClientCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, clusterName, arcSettingName, extensionName, extension, options)
	if err != nil {
		return ExtensionsClientCreatePollerResponse{}, err
	}
	result := ExtensionsClientCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ExtensionsClient.Create", "azure-async-operation", resp, client.pl)
	if err != nil {
		return ExtensionsClientCreatePollerResponse{}, err
	}
	result.Poller = &ExtensionsClientCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Create Extension for HCI cluster.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ExtensionsClient) create(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, extensionName string, extension Extension, options *ExtensionsClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, clusterName, arcSettingName, extensionName, extension, options)
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

// createCreateRequest creates the Create request.
func (client *ExtensionsClient) createCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, extensionName string, extension Extension, options *ExtensionsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/arcSettings/{arcSettingName}/extensions/{extensionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if arcSettingName == "" {
		return nil, errors.New("parameter arcSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{arcSettingName}", url.PathEscape(arcSettingName))
	if extensionName == "" {
		return nil, errors.New("parameter extensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionName}", url.PathEscape(extensionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, extension)
}

// BeginDelete - Delete particular Arc Extension of HCI Cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// clusterName - The name of the cluster.
// arcSettingName - The name of the proxy resource holding details of HCI ArcSetting information.
// extensionName - The name of the machine extension.
// options - ExtensionsClientBeginDeleteOptions contains the optional parameters for the ExtensionsClient.BeginDelete method.
func (client *ExtensionsClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, extensionName string, options *ExtensionsClientBeginDeleteOptions) (ExtensionsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, clusterName, arcSettingName, extensionName, options)
	if err != nil {
		return ExtensionsClientDeletePollerResponse{}, err
	}
	result := ExtensionsClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ExtensionsClient.Delete", "azure-async-operation", resp, client.pl)
	if err != nil {
		return ExtensionsClientDeletePollerResponse{}, err
	}
	result.Poller = &ExtensionsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Delete particular Arc Extension of HCI Cluster.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ExtensionsClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, extensionName string, options *ExtensionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, arcSettingName, extensionName, options)
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
func (client *ExtensionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, extensionName string, options *ExtensionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/arcSettings/{arcSettingName}/extensions/{extensionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if arcSettingName == "" {
		return nil, errors.New("parameter arcSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{arcSettingName}", url.PathEscape(arcSettingName))
	if extensionName == "" {
		return nil, errors.New("parameter extensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionName}", url.PathEscape(extensionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get particular Arc Extension of HCI Cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// clusterName - The name of the cluster.
// arcSettingName - The name of the proxy resource holding details of HCI ArcSetting information.
// extensionName - The name of the machine extension.
// options - ExtensionsClientGetOptions contains the optional parameters for the ExtensionsClient.Get method.
func (client *ExtensionsClient) Get(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, extensionName string, options *ExtensionsClientGetOptions) (ExtensionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, arcSettingName, extensionName, options)
	if err != nil {
		return ExtensionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExtensionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExtensionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExtensionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, extensionName string, options *ExtensionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/arcSettings/{arcSettingName}/extensions/{extensionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if arcSettingName == "" {
		return nil, errors.New("parameter arcSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{arcSettingName}", url.PathEscape(arcSettingName))
	if extensionName == "" {
		return nil, errors.New("parameter extensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionName}", url.PathEscape(extensionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExtensionsClient) getHandleResponse(resp *http.Response) (ExtensionsClientGetResponse, error) {
	result := ExtensionsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Extension); err != nil {
		return ExtensionsClientGetResponse{}, err
	}
	return result, nil
}

// ListByArcSetting - List all Extensions under ArcSetting resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// clusterName - The name of the cluster.
// arcSettingName - The name of the proxy resource holding details of HCI ArcSetting information.
// options - ExtensionsClientListByArcSettingOptions contains the optional parameters for the ExtensionsClient.ListByArcSetting
// method.
func (client *ExtensionsClient) ListByArcSetting(resourceGroupName string, clusterName string, arcSettingName string, options *ExtensionsClientListByArcSettingOptions) *ExtensionsClientListByArcSettingPager {
	return &ExtensionsClientListByArcSettingPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByArcSettingCreateRequest(ctx, resourceGroupName, clusterName, arcSettingName, options)
		},
		advancer: func(ctx context.Context, resp ExtensionsClientListByArcSettingResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ExtensionList.NextLink)
		},
	}
}

// listByArcSettingCreateRequest creates the ListByArcSetting request.
func (client *ExtensionsClient) listByArcSettingCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, options *ExtensionsClientListByArcSettingOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/arcSettings/{arcSettingName}/extensions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if arcSettingName == "" {
		return nil, errors.New("parameter arcSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{arcSettingName}", url.PathEscape(arcSettingName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByArcSettingHandleResponse handles the ListByArcSetting response.
func (client *ExtensionsClient) listByArcSettingHandleResponse(resp *http.Response) (ExtensionsClientListByArcSettingResponse, error) {
	result := ExtensionsClientListByArcSettingResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtensionList); err != nil {
		return ExtensionsClientListByArcSettingResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update Extension for HCI cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// clusterName - The name of the cluster.
// arcSettingName - The name of the proxy resource holding details of HCI ArcSetting information.
// extensionName - The name of the machine extension.
// extension - Details of the Machine Extension to be created.
// options - ExtensionsClientBeginUpdateOptions contains the optional parameters for the ExtensionsClient.BeginUpdate method.
func (client *ExtensionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, extensionName string, extension Extension, options *ExtensionsClientBeginUpdateOptions) (ExtensionsClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, clusterName, arcSettingName, extensionName, extension, options)
	if err != nil {
		return ExtensionsClientUpdatePollerResponse{}, err
	}
	result := ExtensionsClientUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ExtensionsClient.Update", "original-uri", resp, client.pl)
	if err != nil {
		return ExtensionsClientUpdatePollerResponse{}, err
	}
	result.Poller = &ExtensionsClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Update Extension for HCI cluster.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ExtensionsClient) update(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, extensionName string, extension Extension, options *ExtensionsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, clusterName, arcSettingName, extensionName, extension, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *ExtensionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, extensionName string, extension Extension, options *ExtensionsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/arcSettings/{arcSettingName}/extensions/{extensionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if arcSettingName == "" {
		return nil, errors.New("parameter arcSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{arcSettingName}", url.PathEscape(arcSettingName))
	if extensionName == "" {
		return nil, errors.New("parameter extensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionName}", url.PathEscape(extensionName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, extension)
}
