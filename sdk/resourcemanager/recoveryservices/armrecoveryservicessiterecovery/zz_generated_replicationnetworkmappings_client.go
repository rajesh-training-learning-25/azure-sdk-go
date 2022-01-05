//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicessiterecovery

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

// ReplicationNetworkMappingsClient contains the methods for the ReplicationNetworkMappings group.
// Don't use this type directly, use NewReplicationNetworkMappingsClient() instead.
type ReplicationNetworkMappingsClient struct {
	ep                string
	pl                runtime.Pipeline
	resourceName      string
	resourceGroupName string
	subscriptionID    string
}

// NewReplicationNetworkMappingsClient creates a new instance of ReplicationNetworkMappingsClient with the specified values.
func NewReplicationNetworkMappingsClient(resourceName string, resourceGroupName string, subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ReplicationNetworkMappingsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ReplicationNetworkMappingsClient{resourceName: resourceName, resourceGroupName: resourceGroupName, subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreate - The operation to create an ASR network mapping.
// If the operation fails it returns a generic error.
func (client *ReplicationNetworkMappingsClient) BeginCreate(ctx context.Context, fabricName string, networkName string, networkMappingName string, input CreateNetworkMappingInput, options *ReplicationNetworkMappingsBeginCreateOptions) (ReplicationNetworkMappingsCreatePollerResponse, error) {
	resp, err := client.create(ctx, fabricName, networkName, networkMappingName, input, options)
	if err != nil {
		return ReplicationNetworkMappingsCreatePollerResponse{}, err
	}
	result := ReplicationNetworkMappingsCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ReplicationNetworkMappingsClient.Create", "", resp, client.pl, client.createHandleError)
	if err != nil {
		return ReplicationNetworkMappingsCreatePollerResponse{}, err
	}
	result.Poller = &ReplicationNetworkMappingsCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - The operation to create an ASR network mapping.
// If the operation fails it returns a generic error.
func (client *ReplicationNetworkMappingsClient) create(ctx context.Context, fabricName string, networkName string, networkMappingName string, input CreateNetworkMappingInput, options *ReplicationNetworkMappingsBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, fabricName, networkName, networkMappingName, input, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.createHandleError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *ReplicationNetworkMappingsClient) createCreateRequest(ctx context.Context, fabricName string, networkName string, networkMappingName string, input CreateNetworkMappingInput, options *ReplicationNetworkMappingsBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks/{networkName}/replicationNetworkMappings/{networkMappingName}"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if networkName == "" {
		return nil, errors.New("parameter networkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkName}", url.PathEscape(networkName))
	if networkMappingName == "" {
		return nil, errors.New("parameter networkMappingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkMappingName}", url.PathEscape(networkMappingName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, input)
}

// createHandleError handles the Create error response.
func (client *ReplicationNetworkMappingsClient) createHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginDelete - The operation to delete a network mapping.
// If the operation fails it returns a generic error.
func (client *ReplicationNetworkMappingsClient) BeginDelete(ctx context.Context, fabricName string, networkName string, networkMappingName string, options *ReplicationNetworkMappingsBeginDeleteOptions) (ReplicationNetworkMappingsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, fabricName, networkName, networkMappingName, options)
	if err != nil {
		return ReplicationNetworkMappingsDeletePollerResponse{}, err
	}
	result := ReplicationNetworkMappingsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ReplicationNetworkMappingsClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return ReplicationNetworkMappingsDeletePollerResponse{}, err
	}
	result.Poller = &ReplicationNetworkMappingsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - The operation to delete a network mapping.
// If the operation fails it returns a generic error.
func (client *ReplicationNetworkMappingsClient) deleteOperation(ctx context.Context, fabricName string, networkName string, networkMappingName string, options *ReplicationNetworkMappingsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, fabricName, networkName, networkMappingName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ReplicationNetworkMappingsClient) deleteCreateRequest(ctx context.Context, fabricName string, networkName string, networkMappingName string, options *ReplicationNetworkMappingsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks/{networkName}/replicationNetworkMappings/{networkMappingName}"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if networkName == "" {
		return nil, errors.New("parameter networkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkName}", url.PathEscape(networkName))
	if networkMappingName == "" {
		return nil, errors.New("parameter networkMappingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkMappingName}", url.PathEscape(networkMappingName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ReplicationNetworkMappingsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Gets the details of an ASR network mapping.
// If the operation fails it returns a generic error.
func (client *ReplicationNetworkMappingsClient) Get(ctx context.Context, fabricName string, networkName string, networkMappingName string, options *ReplicationNetworkMappingsGetOptions) (ReplicationNetworkMappingsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, fabricName, networkName, networkMappingName, options)
	if err != nil {
		return ReplicationNetworkMappingsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ReplicationNetworkMappingsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ReplicationNetworkMappingsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ReplicationNetworkMappingsClient) getCreateRequest(ctx context.Context, fabricName string, networkName string, networkMappingName string, options *ReplicationNetworkMappingsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks/{networkName}/replicationNetworkMappings/{networkMappingName}"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if networkName == "" {
		return nil, errors.New("parameter networkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkName}", url.PathEscape(networkName))
	if networkMappingName == "" {
		return nil, errors.New("parameter networkMappingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkMappingName}", url.PathEscape(networkMappingName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ReplicationNetworkMappingsClient) getHandleResponse(resp *http.Response) (ReplicationNetworkMappingsGetResponse, error) {
	result := ReplicationNetworkMappingsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkMapping); err != nil {
		return ReplicationNetworkMappingsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ReplicationNetworkMappingsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// List - Lists all ASR network mappings in the vault.
// If the operation fails it returns a generic error.
func (client *ReplicationNetworkMappingsClient) List(options *ReplicationNetworkMappingsListOptions) *ReplicationNetworkMappingsListPager {
	return &ReplicationNetworkMappingsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ReplicationNetworkMappingsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.NetworkMappingCollection.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ReplicationNetworkMappingsClient) listCreateRequest(ctx context.Context, options *ReplicationNetworkMappingsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationNetworkMappings"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ReplicationNetworkMappingsClient) listHandleResponse(resp *http.Response) (ReplicationNetworkMappingsListResponse, error) {
	result := ReplicationNetworkMappingsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkMappingCollection); err != nil {
		return ReplicationNetworkMappingsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ReplicationNetworkMappingsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByReplicationNetworks - Lists all ASR network mappings for the specified network.
// If the operation fails it returns a generic error.
func (client *ReplicationNetworkMappingsClient) ListByReplicationNetworks(fabricName string, networkName string, options *ReplicationNetworkMappingsListByReplicationNetworksOptions) *ReplicationNetworkMappingsListByReplicationNetworksPager {
	return &ReplicationNetworkMappingsListByReplicationNetworksPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByReplicationNetworksCreateRequest(ctx, fabricName, networkName, options)
		},
		advancer: func(ctx context.Context, resp ReplicationNetworkMappingsListByReplicationNetworksResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.NetworkMappingCollection.NextLink)
		},
	}
}

// listByReplicationNetworksCreateRequest creates the ListByReplicationNetworks request.
func (client *ReplicationNetworkMappingsClient) listByReplicationNetworksCreateRequest(ctx context.Context, fabricName string, networkName string, options *ReplicationNetworkMappingsListByReplicationNetworksOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks/{networkName}/replicationNetworkMappings"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if networkName == "" {
		return nil, errors.New("parameter networkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkName}", url.PathEscape(networkName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByReplicationNetworksHandleResponse handles the ListByReplicationNetworks response.
func (client *ReplicationNetworkMappingsClient) listByReplicationNetworksHandleResponse(resp *http.Response) (ReplicationNetworkMappingsListByReplicationNetworksResponse, error) {
	result := ReplicationNetworkMappingsListByReplicationNetworksResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkMappingCollection); err != nil {
		return ReplicationNetworkMappingsListByReplicationNetworksResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByReplicationNetworksHandleError handles the ListByReplicationNetworks error response.
func (client *ReplicationNetworkMappingsClient) listByReplicationNetworksHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginUpdate - The operation to update an ASR network mapping.
// If the operation fails it returns a generic error.
func (client *ReplicationNetworkMappingsClient) BeginUpdate(ctx context.Context, fabricName string, networkName string, networkMappingName string, input UpdateNetworkMappingInput, options *ReplicationNetworkMappingsBeginUpdateOptions) (ReplicationNetworkMappingsUpdatePollerResponse, error) {
	resp, err := client.update(ctx, fabricName, networkName, networkMappingName, input, options)
	if err != nil {
		return ReplicationNetworkMappingsUpdatePollerResponse{}, err
	}
	result := ReplicationNetworkMappingsUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ReplicationNetworkMappingsClient.Update", "", resp, client.pl, client.updateHandleError)
	if err != nil {
		return ReplicationNetworkMappingsUpdatePollerResponse{}, err
	}
	result.Poller = &ReplicationNetworkMappingsUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - The operation to update an ASR network mapping.
// If the operation fails it returns a generic error.
func (client *ReplicationNetworkMappingsClient) update(ctx context.Context, fabricName string, networkName string, networkMappingName string, input UpdateNetworkMappingInput, options *ReplicationNetworkMappingsBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, fabricName, networkName, networkMappingName, input, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *ReplicationNetworkMappingsClient) updateCreateRequest(ctx context.Context, fabricName string, networkName string, networkMappingName string, input UpdateNetworkMappingInput, options *ReplicationNetworkMappingsBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks/{networkName}/replicationNetworkMappings/{networkMappingName}"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if networkName == "" {
		return nil, errors.New("parameter networkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkName}", url.PathEscape(networkName))
	if networkMappingName == "" {
		return nil, errors.New("parameter networkMappingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkMappingName}", url.PathEscape(networkMappingName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, input)
}

// updateHandleError handles the Update error response.
func (client *ReplicationNetworkMappingsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
