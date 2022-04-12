//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoperationalinsights

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

// StorageInsightConfigsClient contains the methods for the StorageInsightConfigs group.
// Don't use this type directly, use NewStorageInsightConfigsClient() instead.
type StorageInsightConfigsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewStorageInsightConfigsClient creates a new instance of StorageInsightConfigsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewStorageInsightConfigsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*StorageInsightConfigsClient, error) {
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
	client := &StorageInsightConfigsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update a storage insight.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// storageInsightName - Name of the storageInsightsConfigs resource
// parameters - The parameters required to create or update a storage insight.
// options - StorageInsightConfigsClientCreateOrUpdateOptions contains the optional parameters for the StorageInsightConfigsClient.CreateOrUpdate
// method.
func (client *StorageInsightConfigsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, storageInsightName string, parameters StorageInsight, options *StorageInsightConfigsClientCreateOrUpdateOptions) (StorageInsightConfigsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, storageInsightName, parameters, options)
	if err != nil {
		return StorageInsightConfigsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StorageInsightConfigsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return StorageInsightConfigsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *StorageInsightConfigsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, storageInsightName string, parameters StorageInsight, options *StorageInsightConfigsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/storageInsightConfigs/{storageInsightName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if storageInsightName == "" {
		return nil, errors.New("parameter storageInsightName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageInsightName}", url.PathEscape(storageInsightName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *StorageInsightConfigsClient) createOrUpdateHandleResponse(resp *http.Response) (StorageInsightConfigsClientCreateOrUpdateResponse, error) {
	result := StorageInsightConfigsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageInsight); err != nil {
		return StorageInsightConfigsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a storageInsightsConfigs resource
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// storageInsightName - Name of the storageInsightsConfigs resource
// options - StorageInsightConfigsClientDeleteOptions contains the optional parameters for the StorageInsightConfigsClient.Delete
// method.
func (client *StorageInsightConfigsClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, storageInsightName string, options *StorageInsightConfigsClientDeleteOptions) (StorageInsightConfigsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, storageInsightName, options)
	if err != nil {
		return StorageInsightConfigsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StorageInsightConfigsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return StorageInsightConfigsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return StorageInsightConfigsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *StorageInsightConfigsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, storageInsightName string, options *StorageInsightConfigsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/storageInsightConfigs/{storageInsightName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if storageInsightName == "" {
		return nil, errors.New("parameter storageInsightName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageInsightName}", url.PathEscape(storageInsightName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets a storage insight instance.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// storageInsightName - Name of the storageInsightsConfigs resource
// options - StorageInsightConfigsClientGetOptions contains the optional parameters for the StorageInsightConfigsClient.Get
// method.
func (client *StorageInsightConfigsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, storageInsightName string, options *StorageInsightConfigsClientGetOptions) (StorageInsightConfigsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, storageInsightName, options)
	if err != nil {
		return StorageInsightConfigsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StorageInsightConfigsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return StorageInsightConfigsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *StorageInsightConfigsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, storageInsightName string, options *StorageInsightConfigsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/storageInsightConfigs/{storageInsightName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if storageInsightName == "" {
		return nil, errors.New("parameter storageInsightName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageInsightName}", url.PathEscape(storageInsightName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *StorageInsightConfigsClient) getHandleResponse(resp *http.Response) (StorageInsightConfigsClientGetResponse, error) {
	result := StorageInsightConfigsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageInsight); err != nil {
		return StorageInsightConfigsClientGetResponse{}, err
	}
	return result, nil
}

// ListByWorkspace - Lists the storage insight instances within a workspace
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// options - StorageInsightConfigsClientListByWorkspaceOptions contains the optional parameters for the StorageInsightConfigsClient.ListByWorkspace
// method.
func (client *StorageInsightConfigsClient) ListByWorkspace(resourceGroupName string, workspaceName string, options *StorageInsightConfigsClientListByWorkspaceOptions) *runtime.Pager[StorageInsightConfigsClientListByWorkspaceResponse] {
	return runtime.NewPager(runtime.PageProcessor[StorageInsightConfigsClientListByWorkspaceResponse]{
		More: func(page StorageInsightConfigsClientListByWorkspaceResponse) bool {
			return page.ODataNextLink != nil && len(*page.ODataNextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *StorageInsightConfigsClientListByWorkspaceResponse) (StorageInsightConfigsClientListByWorkspaceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.ODataNextLink)
			}
			if err != nil {
				return StorageInsightConfigsClientListByWorkspaceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return StorageInsightConfigsClientListByWorkspaceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return StorageInsightConfigsClientListByWorkspaceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByWorkspaceHandleResponse(resp)
		},
	})
}

// listByWorkspaceCreateRequest creates the ListByWorkspace request.
func (client *StorageInsightConfigsClient) listByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *StorageInsightConfigsClientListByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/storageInsightConfigs"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByWorkspaceHandleResponse handles the ListByWorkspace response.
func (client *StorageInsightConfigsClient) listByWorkspaceHandleResponse(resp *http.Response) (StorageInsightConfigsClientListByWorkspaceResponse, error) {
	result := StorageInsightConfigsClientListByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageInsightListResult); err != nil {
		return StorageInsightConfigsClientListByWorkspaceResponse{}, err
	}
	return result, nil
}
