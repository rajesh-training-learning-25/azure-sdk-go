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

// TablesClient contains the methods for the Tables group.
// Don't use this type directly, use NewTablesClient() instead.
type TablesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewTablesClient creates a new instance of TablesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewTablesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *TablesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &TablesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Update or Create a Log Analytics workspace table.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// tableName - The name of the table.
// parameters - The parameters required to update table properties.
// options - TablesClientBeginCreateOrUpdateOptions contains the optional parameters for the TablesClient.BeginCreateOrUpdate
// method.
func (client *TablesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, tableName string, parameters Table, options *TablesClientBeginCreateOrUpdateOptions) (TablesClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, tableName, parameters, options)
	if err != nil {
		return TablesClientCreateOrUpdatePollerResponse{}, err
	}
	result := TablesClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("TablesClient.CreateOrUpdate", "azure-async-operation", resp, client.pl)
	if err != nil {
		return TablesClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &TablesClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Update or Create a Log Analytics workspace table.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *TablesClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, tableName string, parameters Table, options *TablesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, tableName, parameters, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *TablesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, tableName string, parameters Table, options *TablesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/tables/{tableName}"
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
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Delete a Log Analytics workspace table.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// tableName - The name of the table.
// options - TablesClientBeginDeleteOptions contains the optional parameters for the TablesClient.BeginDelete method.
func (client *TablesClient) BeginDelete(ctx context.Context, resourceGroupName string, workspaceName string, tableName string, options *TablesClientBeginDeleteOptions) (TablesClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, workspaceName, tableName, options)
	if err != nil {
		return TablesClientDeletePollerResponse{}, err
	}
	result := TablesClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("TablesClient.Delete", "azure-async-operation", resp, client.pl)
	if err != nil {
		return TablesClientDeletePollerResponse{}, err
	}
	result.Poller = &TablesClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Delete a Log Analytics workspace table.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *TablesClient) deleteOperation(ctx context.Context, resourceGroupName string, workspaceName string, tableName string, options *TablesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, tableName, options)
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
func (client *TablesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, tableName string, options *TablesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/tables/{tableName}"
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
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets a Log Analytics workspace table.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// tableName - The name of the table.
// options - TablesClientGetOptions contains the optional parameters for the TablesClient.Get method.
func (client *TablesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, tableName string, options *TablesClientGetOptions) (TablesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, tableName, options)
	if err != nil {
		return TablesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TablesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TablesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *TablesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, tableName string, options *TablesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/tables/{tableName}"
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
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TablesClient) getHandleResponse(resp *http.Response) (TablesClientGetResponse, error) {
	result := TablesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Table); err != nil {
		return TablesClientGetResponse{}, err
	}
	return result, nil
}

// ListByWorkspace - Gets all the tables for the specified Log Analytics workspace.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// options - TablesClientListByWorkspaceOptions contains the optional parameters for the TablesClient.ListByWorkspace method.
func (client *TablesClient) ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string, options *TablesClientListByWorkspaceOptions) (TablesClientListByWorkspaceResponse, error) {
	req, err := client.listByWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, options)
	if err != nil {
		return TablesClientListByWorkspaceResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TablesClientListByWorkspaceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TablesClientListByWorkspaceResponse{}, runtime.NewResponseError(resp)
	}
	return client.listByWorkspaceHandleResponse(resp)
}

// listByWorkspaceCreateRequest creates the ListByWorkspace request.
func (client *TablesClient) listByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *TablesClientListByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/tables"
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
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByWorkspaceHandleResponse handles the ListByWorkspace response.
func (client *TablesClient) listByWorkspaceHandleResponse(resp *http.Response) (TablesClientListByWorkspaceResponse, error) {
	result := TablesClientListByWorkspaceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TablesListResult); err != nil {
		return TablesClientListByWorkspaceResponse{}, err
	}
	return result, nil
}

// Migrate - Migrate a Log Analytics table from support of the Data Collector API and Custom Fields features to support of
// Data Collection Rule-based Custom Logs.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// tableName - The name of the table.
// options - TablesClientMigrateOptions contains the optional parameters for the TablesClient.Migrate method.
func (client *TablesClient) Migrate(ctx context.Context, resourceGroupName string, workspaceName string, tableName string, options *TablesClientMigrateOptions) (TablesClientMigrateResponse, error) {
	req, err := client.migrateCreateRequest(ctx, resourceGroupName, workspaceName, tableName, options)
	if err != nil {
		return TablesClientMigrateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TablesClientMigrateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TablesClientMigrateResponse{}, runtime.NewResponseError(resp)
	}
	return TablesClientMigrateResponse{RawResponse: resp}, nil
}

// migrateCreateRequest creates the Migrate request.
func (client *TablesClient) migrateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, tableName string, options *TablesClientMigrateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/tables/{tableName}/migrate"
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
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginUpdate - Update a Log Analytics workspace table.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// tableName - The name of the table.
// parameters - The parameters required to update table properties.
// options - TablesClientBeginUpdateOptions contains the optional parameters for the TablesClient.BeginUpdate method.
func (client *TablesClient) BeginUpdate(ctx context.Context, resourceGroupName string, workspaceName string, tableName string, parameters Table, options *TablesClientBeginUpdateOptions) (TablesClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, workspaceName, tableName, parameters, options)
	if err != nil {
		return TablesClientUpdatePollerResponse{}, err
	}
	result := TablesClientUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("TablesClient.Update", "azure-async-operation", resp, client.pl)
	if err != nil {
		return TablesClientUpdatePollerResponse{}, err
	}
	result.Poller = &TablesClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Update a Log Analytics workspace table.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *TablesClient) update(ctx context.Context, resourceGroupName string, workspaceName string, tableName string, parameters Table, options *TablesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, workspaceName, tableName, parameters, options)
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

// updateCreateRequest creates the Update request.
func (client *TablesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, tableName string, parameters Table, options *TablesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/tables/{tableName}"
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
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}
