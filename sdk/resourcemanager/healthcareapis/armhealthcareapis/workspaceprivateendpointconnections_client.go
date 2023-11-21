//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhealthcareapis

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// WorkspacePrivateEndpointConnectionsClient contains the methods for the WorkspacePrivateEndpointConnections group.
// Don't use this type directly, use NewWorkspacePrivateEndpointConnectionsClient() instead.
type WorkspacePrivateEndpointConnectionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWorkspacePrivateEndpointConnectionsClient creates a new instance of WorkspacePrivateEndpointConnectionsClient with the specified values.
//   - subscriptionID - The subscription identifier.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWorkspacePrivateEndpointConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkspacePrivateEndpointConnectionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WorkspacePrivateEndpointConnectionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Update the state of the specified private endpoint connection associated with the workspace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01
//   - resourceGroupName - The name of the resource group that contains the service instance.
//   - workspaceName - The name of workspace resource.
//   - privateEndpointConnectionName - The name of the private endpoint connection associated with the Azure resource
//   - properties - The private endpoint connection properties.
//   - options - WorkspacePrivateEndpointConnectionsClientBeginCreateOrUpdateOptions contains the optional parameters for the
//     WorkspacePrivateEndpointConnectionsClient.BeginCreateOrUpdate method.
func (client *WorkspacePrivateEndpointConnectionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, privateEndpointConnectionName string, properties PrivateEndpointConnectionDescription, options *WorkspacePrivateEndpointConnectionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[WorkspacePrivateEndpointConnectionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, privateEndpointConnectionName, properties, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[WorkspacePrivateEndpointConnectionsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[WorkspacePrivateEndpointConnectionsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Update the state of the specified private endpoint connection associated with the workspace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01
func (client *WorkspacePrivateEndpointConnectionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, privateEndpointConnectionName string, properties PrivateEndpointConnectionDescription, options *WorkspacePrivateEndpointConnectionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "WorkspacePrivateEndpointConnectionsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, privateEndpointConnectionName, properties, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WorkspacePrivateEndpointConnectionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, privateEndpointConnectionName string, properties PrivateEndpointConnectionDescription, options *WorkspacePrivateEndpointConnectionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/privateEndpointConnections/{privateEndpointConnectionName}"
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
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a private endpoint connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01
//   - resourceGroupName - The name of the resource group that contains the service instance.
//   - workspaceName - The name of workspace resource.
//   - privateEndpointConnectionName - The name of the private endpoint connection associated with the Azure resource
//   - options - WorkspacePrivateEndpointConnectionsClientBeginDeleteOptions contains the optional parameters for the WorkspacePrivateEndpointConnectionsClient.BeginDelete
//     method.
func (client *WorkspacePrivateEndpointConnectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, workspaceName string, privateEndpointConnectionName string, options *WorkspacePrivateEndpointConnectionsClientBeginDeleteOptions) (*runtime.Poller[WorkspacePrivateEndpointConnectionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, workspaceName, privateEndpointConnectionName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[WorkspacePrivateEndpointConnectionsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[WorkspacePrivateEndpointConnectionsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes a private endpoint connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01
func (client *WorkspacePrivateEndpointConnectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, workspaceName string, privateEndpointConnectionName string, options *WorkspacePrivateEndpointConnectionsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "WorkspacePrivateEndpointConnectionsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, privateEndpointConnectionName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WorkspacePrivateEndpointConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, privateEndpointConnectionName string, options *WorkspacePrivateEndpointConnectionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/privateEndpointConnections/{privateEndpointConnectionName}"
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
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the specified private endpoint connection associated with the workspace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01
//   - resourceGroupName - The name of the resource group that contains the service instance.
//   - workspaceName - The name of workspace resource.
//   - privateEndpointConnectionName - The name of the private endpoint connection associated with the Azure resource
//   - options - WorkspacePrivateEndpointConnectionsClientGetOptions contains the optional parameters for the WorkspacePrivateEndpointConnectionsClient.Get
//     method.
func (client *WorkspacePrivateEndpointConnectionsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, privateEndpointConnectionName string, options *WorkspacePrivateEndpointConnectionsClientGetOptions) (WorkspacePrivateEndpointConnectionsClientGetResponse, error) {
	var err error
	const operationName = "WorkspacePrivateEndpointConnectionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, privateEndpointConnectionName, options)
	if err != nil {
		return WorkspacePrivateEndpointConnectionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspacePrivateEndpointConnectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WorkspacePrivateEndpointConnectionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *WorkspacePrivateEndpointConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, privateEndpointConnectionName string, options *WorkspacePrivateEndpointConnectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/privateEndpointConnections/{privateEndpointConnectionName}"
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
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkspacePrivateEndpointConnectionsClient) getHandleResponse(resp *http.Response) (WorkspacePrivateEndpointConnectionsClientGetResponse, error) {
	result := WorkspacePrivateEndpointConnectionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnectionDescription); err != nil {
		return WorkspacePrivateEndpointConnectionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByWorkspacePager - Lists all private endpoint connections for a workspace.
//
// Generated from API version 2021-11-01
//   - resourceGroupName - The name of the resource group that contains the service instance.
//   - workspaceName - The name of workspace resource.
//   - options - WorkspacePrivateEndpointConnectionsClientListByWorkspaceOptions contains the optional parameters for the WorkspacePrivateEndpointConnectionsClient.NewListByWorkspacePager
//     method.
func (client *WorkspacePrivateEndpointConnectionsClient) NewListByWorkspacePager(resourceGroupName string, workspaceName string, options *WorkspacePrivateEndpointConnectionsClientListByWorkspaceOptions) *runtime.Pager[WorkspacePrivateEndpointConnectionsClientListByWorkspaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[WorkspacePrivateEndpointConnectionsClientListByWorkspaceResponse]{
		More: func(page WorkspacePrivateEndpointConnectionsClientListByWorkspaceResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *WorkspacePrivateEndpointConnectionsClientListByWorkspaceResponse) (WorkspacePrivateEndpointConnectionsClientListByWorkspaceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "WorkspacePrivateEndpointConnectionsClient.NewListByWorkspacePager")
			req, err := client.listByWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, options)
			if err != nil {
				return WorkspacePrivateEndpointConnectionsClientListByWorkspaceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return WorkspacePrivateEndpointConnectionsClientListByWorkspaceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WorkspacePrivateEndpointConnectionsClientListByWorkspaceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByWorkspaceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByWorkspaceCreateRequest creates the ListByWorkspace request.
func (client *WorkspacePrivateEndpointConnectionsClient) listByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspacePrivateEndpointConnectionsClientListByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/privateEndpointConnections"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByWorkspaceHandleResponse handles the ListByWorkspace response.
func (client *WorkspacePrivateEndpointConnectionsClient) listByWorkspaceHandleResponse(resp *http.Response) (WorkspacePrivateEndpointConnectionsClientListByWorkspaceResponse, error) {
	result := WorkspacePrivateEndpointConnectionsClientListByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnectionListResultDescription); err != nil {
		return WorkspacePrivateEndpointConnectionsClientListByWorkspaceResponse{}, err
	}
	return result, nil
}
