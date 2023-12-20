//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// ServerAzureADOnlyAuthenticationsClient contains the methods for the ServerAzureADOnlyAuthentications group.
// Don't use this type directly, use NewServerAzureADOnlyAuthenticationsClient() instead.
type ServerAzureADOnlyAuthenticationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewServerAzureADOnlyAuthenticationsClient creates a new instance of ServerAzureADOnlyAuthenticationsClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewServerAzureADOnlyAuthenticationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServerAzureADOnlyAuthenticationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ServerAzureADOnlyAuthenticationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Sets Server Active Directory only authentication property or updates an existing server Active Directory
// only authentication property.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - authenticationName - The name of server azure active directory only authentication.
//   - parameters - The required parameters for creating or updating an Active Directory only authentication property.
//   - options - ServerAzureADOnlyAuthenticationsClientBeginCreateOrUpdateOptions contains the optional parameters for the ServerAzureADOnlyAuthenticationsClient.BeginCreateOrUpdate
//     method.
func (client *ServerAzureADOnlyAuthenticationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, authenticationName AuthenticationName, parameters ServerAzureADOnlyAuthentication, options *ServerAzureADOnlyAuthenticationsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ServerAzureADOnlyAuthenticationsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, authenticationName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ServerAzureADOnlyAuthenticationsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ServerAzureADOnlyAuthenticationsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Sets Server Active Directory only authentication property or updates an existing server Active Directory
// only authentication property.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
func (client *ServerAzureADOnlyAuthenticationsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, authenticationName AuthenticationName, parameters ServerAzureADOnlyAuthentication, options *ServerAzureADOnlyAuthenticationsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ServerAzureADOnlyAuthenticationsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, authenticationName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServerAzureADOnlyAuthenticationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, authenticationName AuthenticationName, parameters ServerAzureADOnlyAuthentication, options *ServerAzureADOnlyAuthenticationsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/azureADOnlyAuthentications/{authenticationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if authenticationName == "" {
		return nil, errors.New("parameter authenticationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authenticationName}", url.PathEscape(string(authenticationName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes an existing server Active Directory only authentication property.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - authenticationName - The name of server azure active directory only authentication.
//   - options - ServerAzureADOnlyAuthenticationsClientBeginDeleteOptions contains the optional parameters for the ServerAzureADOnlyAuthenticationsClient.BeginDelete
//     method.
func (client *ServerAzureADOnlyAuthenticationsClient) BeginDelete(ctx context.Context, resourceGroupName string, serverName string, authenticationName AuthenticationName, options *ServerAzureADOnlyAuthenticationsClientBeginDeleteOptions) (*runtime.Poller[ServerAzureADOnlyAuthenticationsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, serverName, authenticationName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ServerAzureADOnlyAuthenticationsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ServerAzureADOnlyAuthenticationsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes an existing server Active Directory only authentication property.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
func (client *ServerAzureADOnlyAuthenticationsClient) deleteOperation(ctx context.Context, resourceGroupName string, serverName string, authenticationName AuthenticationName, options *ServerAzureADOnlyAuthenticationsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ServerAzureADOnlyAuthenticationsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, authenticationName, options)
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
func (client *ServerAzureADOnlyAuthenticationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, authenticationName AuthenticationName, options *ServerAzureADOnlyAuthenticationsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/azureADOnlyAuthentications/{authenticationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if authenticationName == "" {
		return nil, errors.New("parameter authenticationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authenticationName}", url.PathEscape(string(authenticationName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets a specific Azure Active Directory only authentication property.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - authenticationName - The name of server azure active directory only authentication.
//   - options - ServerAzureADOnlyAuthenticationsClientGetOptions contains the optional parameters for the ServerAzureADOnlyAuthenticationsClient.Get
//     method.
func (client *ServerAzureADOnlyAuthenticationsClient) Get(ctx context.Context, resourceGroupName string, serverName string, authenticationName AuthenticationName, options *ServerAzureADOnlyAuthenticationsClientGetOptions) (ServerAzureADOnlyAuthenticationsClientGetResponse, error) {
	var err error
	const operationName = "ServerAzureADOnlyAuthenticationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, authenticationName, options)
	if err != nil {
		return ServerAzureADOnlyAuthenticationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServerAzureADOnlyAuthenticationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServerAzureADOnlyAuthenticationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ServerAzureADOnlyAuthenticationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, authenticationName AuthenticationName, options *ServerAzureADOnlyAuthenticationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/azureADOnlyAuthentications/{authenticationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if authenticationName == "" {
		return nil, errors.New("parameter authenticationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authenticationName}", url.PathEscape(string(authenticationName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServerAzureADOnlyAuthenticationsClient) getHandleResponse(resp *http.Response) (ServerAzureADOnlyAuthenticationsClientGetResponse, error) {
	result := ServerAzureADOnlyAuthenticationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerAzureADOnlyAuthentication); err != nil {
		return ServerAzureADOnlyAuthenticationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServerPager - Gets a list of server Azure Active Directory only authentications.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - options - ServerAzureADOnlyAuthenticationsClientListByServerOptions contains the optional parameters for the ServerAzureADOnlyAuthenticationsClient.NewListByServerPager
//     method.
func (client *ServerAzureADOnlyAuthenticationsClient) NewListByServerPager(resourceGroupName string, serverName string, options *ServerAzureADOnlyAuthenticationsClientListByServerOptions) *runtime.Pager[ServerAzureADOnlyAuthenticationsClientListByServerResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServerAzureADOnlyAuthenticationsClientListByServerResponse]{
		More: func(page ServerAzureADOnlyAuthenticationsClientListByServerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServerAzureADOnlyAuthenticationsClientListByServerResponse) (ServerAzureADOnlyAuthenticationsClientListByServerResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ServerAzureADOnlyAuthenticationsClient.NewListByServerPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
			}, nil)
			if err != nil {
				return ServerAzureADOnlyAuthenticationsClientListByServerResponse{}, err
			}
			return client.listByServerHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByServerCreateRequest creates the ListByServer request.
func (client *ServerAzureADOnlyAuthenticationsClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerAzureADOnlyAuthenticationsClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/azureADOnlyAuthentications"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *ServerAzureADOnlyAuthenticationsClient) listByServerHandleResponse(resp *http.Response) (ServerAzureADOnlyAuthenticationsClientListByServerResponse, error) {
	result := ServerAzureADOnlyAuthenticationsClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureADOnlyAuthListResult); err != nil {
		return ServerAzureADOnlyAuthenticationsClientListByServerResponse{}, err
	}
	return result, nil
}
