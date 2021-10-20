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
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// ServerAdministratorsClient contains the methods for the ServerAdministrators group.
// Don't use this type directly, use NewServerAdministratorsClient() instead.
type ServerAdministratorsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewServerAdministratorsClient creates a new instance of ServerAdministratorsClient with the specified values.
func NewServerAdministratorsClient(con *arm.Connection, subscriptionID string) *ServerAdministratorsClient {
	return &ServerAdministratorsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or update active directory administrator on an existing server. The update action will overwrite the existing administrator.
// If the operation fails it returns the *CloudError error type.
func (client *ServerAdministratorsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, properties ServerAdministratorResource, options *ServerAdministratorsBeginCreateOrUpdateOptions) (ServerAdministratorsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, properties, options)
	if err != nil {
		return ServerAdministratorsCreateOrUpdatePollerResponse{}, err
	}
	result := ServerAdministratorsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ServerAdministratorsClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return ServerAdministratorsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ServerAdministratorsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or update active directory administrator on an existing server. The update action will overwrite the existing administrator.
// If the operation fails it returns the *CloudError error type.
func (client *ServerAdministratorsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, properties ServerAdministratorResource, options *ServerAdministratorsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, properties, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServerAdministratorsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, properties ServerAdministratorResource, options *ServerAdministratorsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/administrators/activeDirectory"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, properties)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ServerAdministratorsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes server active directory administrator.
// If the operation fails it returns the *CloudError error type.
func (client *ServerAdministratorsClient) BeginDelete(ctx context.Context, resourceGroupName string, serverName string, options *ServerAdministratorsBeginDeleteOptions) (ServerAdministratorsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, serverName, options)
	if err != nil {
		return ServerAdministratorsDeletePollerResponse{}, err
	}
	result := ServerAdministratorsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ServerAdministratorsClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return ServerAdministratorsDeletePollerResponse{}, err
	}
	result.Poller = &ServerAdministratorsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes server active directory administrator.
// If the operation fails it returns the *CloudError error type.
func (client *ServerAdministratorsClient) deleteOperation(ctx context.Context, resourceGroupName string, serverName string, options *ServerAdministratorsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ServerAdministratorsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerAdministratorsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/administrators/activeDirectory"
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ServerAdministratorsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets information about a AAD server administrator.
// If the operation fails it returns the *CloudError error type.
func (client *ServerAdministratorsClient) Get(ctx context.Context, resourceGroupName string, serverName string, options *ServerAdministratorsGetOptions) (ServerAdministratorsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, options)
	if err != nil {
		return ServerAdministratorsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServerAdministratorsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServerAdministratorsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServerAdministratorsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerAdministratorsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/administrators/activeDirectory"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServerAdministratorsClient) getHandleResponse(resp *http.Response) (ServerAdministratorsGetResponse, error) {
	result := ServerAdministratorsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerAdministratorResource); err != nil {
		return ServerAdministratorsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ServerAdministratorsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Returns a list of server Administrators.
// If the operation fails it returns the *CloudError error type.
func (client *ServerAdministratorsClient) List(ctx context.Context, resourceGroupName string, serverName string, options *ServerAdministratorsListOptions) (ServerAdministratorsListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, serverName, options)
	if err != nil {
		return ServerAdministratorsListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServerAdministratorsListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServerAdministratorsListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ServerAdministratorsClient) listCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerAdministratorsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/administrators"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ServerAdministratorsClient) listHandleResponse(resp *http.Response) (ServerAdministratorsListResponse, error) {
	result := ServerAdministratorsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerAdministratorResourceListResult); err != nil {
		return ServerAdministratorsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ServerAdministratorsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
