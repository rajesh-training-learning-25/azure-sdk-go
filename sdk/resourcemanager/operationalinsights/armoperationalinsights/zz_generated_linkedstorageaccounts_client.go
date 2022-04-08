//go:build go1.16
// +build go1.16

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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// LinkedStorageAccountsClient contains the methods for the LinkedStorageAccounts group.
// Don't use this type directly, use NewLinkedStorageAccountsClient() instead.
type LinkedStorageAccountsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewLinkedStorageAccountsClient creates a new instance of LinkedStorageAccountsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewLinkedStorageAccountsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *LinkedStorageAccountsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &LinkedStorageAccountsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrUpdate - Create or Update a link relation between current workspace and a group of storage accounts of a specific
// data source type.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// dataSourceType - Linked storage accounts type.
// parameters - The parameters required to create or update linked storage accounts.
// options - LinkedStorageAccountsClientCreateOrUpdateOptions contains the optional parameters for the LinkedStorageAccountsClient.CreateOrUpdate
// method.
func (client *LinkedStorageAccountsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceType DataSourceType, parameters LinkedStorageAccountsResource, options *LinkedStorageAccountsClientCreateOrUpdateOptions) (LinkedStorageAccountsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, dataSourceType, parameters, options)
	if err != nil {
		return LinkedStorageAccountsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LinkedStorageAccountsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LinkedStorageAccountsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *LinkedStorageAccountsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceType DataSourceType, parameters LinkedStorageAccountsResource, options *LinkedStorageAccountsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/linkedStorageAccounts/{dataSourceType}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if dataSourceType == "" {
		return nil, errors.New("parameter dataSourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataSourceType}", url.PathEscape(string(dataSourceType)))
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
func (client *LinkedStorageAccountsClient) createOrUpdateHandleResponse(resp *http.Response) (LinkedStorageAccountsClientCreateOrUpdateResponse, error) {
	result := LinkedStorageAccountsClientCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkedStorageAccountsResource); err != nil {
		return LinkedStorageAccountsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes all linked storage accounts of a specific data source type associated with the specified workspace.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// dataSourceType - Linked storage accounts type.
// options - LinkedStorageAccountsClientDeleteOptions contains the optional parameters for the LinkedStorageAccountsClient.Delete
// method.
func (client *LinkedStorageAccountsClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceType DataSourceType, options *LinkedStorageAccountsClientDeleteOptions) (LinkedStorageAccountsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, dataSourceType, options)
	if err != nil {
		return LinkedStorageAccountsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LinkedStorageAccountsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LinkedStorageAccountsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return LinkedStorageAccountsClientDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *LinkedStorageAccountsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceType DataSourceType, options *LinkedStorageAccountsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/linkedStorageAccounts/{dataSourceType}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if dataSourceType == "" {
		return nil, errors.New("parameter dataSourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataSourceType}", url.PathEscape(string(dataSourceType)))
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

// Get - Gets all linked storage account of a specific data source type associated with the specified workspace.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// dataSourceType - Linked storage accounts type.
// options - LinkedStorageAccountsClientGetOptions contains the optional parameters for the LinkedStorageAccountsClient.Get
// method.
func (client *LinkedStorageAccountsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceType DataSourceType, options *LinkedStorageAccountsClientGetOptions) (LinkedStorageAccountsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, dataSourceType, options)
	if err != nil {
		return LinkedStorageAccountsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LinkedStorageAccountsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LinkedStorageAccountsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LinkedStorageAccountsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceType DataSourceType, options *LinkedStorageAccountsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/linkedStorageAccounts/{dataSourceType}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if dataSourceType == "" {
		return nil, errors.New("parameter dataSourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataSourceType}", url.PathEscape(string(dataSourceType)))
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
func (client *LinkedStorageAccountsClient) getHandleResponse(resp *http.Response) (LinkedStorageAccountsClientGetResponse, error) {
	result := LinkedStorageAccountsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkedStorageAccountsResource); err != nil {
		return LinkedStorageAccountsClientGetResponse{}, err
	}
	return result, nil
}

// ListByWorkspace - Gets all linked storage accounts associated with the specified workspace, storage accounts will be sorted
// by their data source type.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// options - LinkedStorageAccountsClientListByWorkspaceOptions contains the optional parameters for the LinkedStorageAccountsClient.ListByWorkspace
// method.
func (client *LinkedStorageAccountsClient) ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string, options *LinkedStorageAccountsClientListByWorkspaceOptions) (LinkedStorageAccountsClientListByWorkspaceResponse, error) {
	req, err := client.listByWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, options)
	if err != nil {
		return LinkedStorageAccountsClientListByWorkspaceResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LinkedStorageAccountsClientListByWorkspaceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LinkedStorageAccountsClientListByWorkspaceResponse{}, runtime.NewResponseError(resp)
	}
	return client.listByWorkspaceHandleResponse(resp)
}

// listByWorkspaceCreateRequest creates the ListByWorkspace request.
func (client *LinkedStorageAccountsClient) listByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *LinkedStorageAccountsClientListByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/linkedStorageAccounts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
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
func (client *LinkedStorageAccountsClient) listByWorkspaceHandleResponse(resp *http.Response) (LinkedStorageAccountsClientListByWorkspaceResponse, error) {
	result := LinkedStorageAccountsClientListByWorkspaceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkedStorageAccountsListResult); err != nil {
		return LinkedStorageAccountsClientListByWorkspaceResponse{}, err
	}
	return result, nil
}
