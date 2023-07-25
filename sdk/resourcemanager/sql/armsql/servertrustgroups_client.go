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

// ServerTrustGroupsClient contains the methods for the ServerTrustGroups group.
// Don't use this type directly, use NewServerTrustGroupsClient() instead.
type ServerTrustGroupsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewServerTrustGroupsClient creates a new instance of ServerTrustGroupsClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewServerTrustGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServerTrustGroupsClient, error) {
	cl, err := arm.NewClient(moduleName+".ServerTrustGroupsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ServerTrustGroupsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a server trust group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - locationName - The name of the region where the resource is located.
//   - serverTrustGroupName - The name of the server trust group.
//   - parameters - The server trust group parameters.
//   - options - ServerTrustGroupsClientBeginCreateOrUpdateOptions contains the optional parameters for the ServerTrustGroupsClient.BeginCreateOrUpdate
//     method.
func (client *ServerTrustGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, locationName string, serverTrustGroupName string, parameters ServerTrustGroup, options *ServerTrustGroupsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ServerTrustGroupsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, locationName, serverTrustGroupName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[ServerTrustGroupsClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ServerTrustGroupsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates a server trust group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
func (client *ServerTrustGroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, locationName string, serverTrustGroupName string, parameters ServerTrustGroup, options *ServerTrustGroupsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, locationName, serverTrustGroupName, parameters, options)
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
func (client *ServerTrustGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, locationName string, serverTrustGroupName string, parameters ServerTrustGroup, options *ServerTrustGroupsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/locations/{locationName}/serverTrustGroups/{serverTrustGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if serverTrustGroupName == "" {
		return nil, errors.New("parameter serverTrustGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverTrustGroupName}", url.PathEscape(serverTrustGroupName))
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

// BeginDelete - Deletes a server trust group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - locationName - The name of the region where the resource is located.
//   - serverTrustGroupName - The name of the server trust group.
//   - options - ServerTrustGroupsClientBeginDeleteOptions contains the optional parameters for the ServerTrustGroupsClient.BeginDelete
//     method.
func (client *ServerTrustGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, locationName string, serverTrustGroupName string, options *ServerTrustGroupsClientBeginDeleteOptions) (*runtime.Poller[ServerTrustGroupsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, locationName, serverTrustGroupName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[ServerTrustGroupsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ServerTrustGroupsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes a server trust group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
func (client *ServerTrustGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, locationName string, serverTrustGroupName string, options *ServerTrustGroupsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, locationName, serverTrustGroupName, options)
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
func (client *ServerTrustGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, locationName string, serverTrustGroupName string, options *ServerTrustGroupsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/locations/{locationName}/serverTrustGroups/{serverTrustGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if serverTrustGroupName == "" {
		return nil, errors.New("parameter serverTrustGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverTrustGroupName}", url.PathEscape(serverTrustGroupName))
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

// Get - Gets a server trust group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - locationName - The name of the region where the resource is located.
//   - serverTrustGroupName - The name of the server trust group.
//   - options - ServerTrustGroupsClientGetOptions contains the optional parameters for the ServerTrustGroupsClient.Get method.
func (client *ServerTrustGroupsClient) Get(ctx context.Context, resourceGroupName string, locationName string, serverTrustGroupName string, options *ServerTrustGroupsClientGetOptions) (ServerTrustGroupsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, locationName, serverTrustGroupName, options)
	if err != nil {
		return ServerTrustGroupsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServerTrustGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServerTrustGroupsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ServerTrustGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, locationName string, serverTrustGroupName string, options *ServerTrustGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/locations/{locationName}/serverTrustGroups/{serverTrustGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if serverTrustGroupName == "" {
		return nil, errors.New("parameter serverTrustGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverTrustGroupName}", url.PathEscape(serverTrustGroupName))
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
func (client *ServerTrustGroupsClient) getHandleResponse(resp *http.Response) (ServerTrustGroupsClientGetResponse, error) {
	result := ServerTrustGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerTrustGroup); err != nil {
		return ServerTrustGroupsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByInstancePager - Gets a server trust groups by instance name.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - options - ServerTrustGroupsClientListByInstanceOptions contains the optional parameters for the ServerTrustGroupsClient.NewListByInstancePager
//     method.
func (client *ServerTrustGroupsClient) NewListByInstancePager(resourceGroupName string, managedInstanceName string, options *ServerTrustGroupsClientListByInstanceOptions) *runtime.Pager[ServerTrustGroupsClientListByInstanceResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServerTrustGroupsClientListByInstanceResponse]{
		More: func(page ServerTrustGroupsClientListByInstanceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServerTrustGroupsClientListByInstanceResponse) (ServerTrustGroupsClientListByInstanceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByInstanceCreateRequest(ctx, resourceGroupName, managedInstanceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ServerTrustGroupsClientListByInstanceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ServerTrustGroupsClientListByInstanceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServerTrustGroupsClientListByInstanceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByInstanceHandleResponse(resp)
		},
	})
}

// listByInstanceCreateRequest creates the ListByInstance request.
func (client *ServerTrustGroupsClient) listByInstanceCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ServerTrustGroupsClientListByInstanceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/serverTrustGroups"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
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

// listByInstanceHandleResponse handles the ListByInstance response.
func (client *ServerTrustGroupsClient) listByInstanceHandleResponse(resp *http.Response) (ServerTrustGroupsClientListByInstanceResponse, error) {
	result := ServerTrustGroupsClientListByInstanceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerTrustGroupListResult); err != nil {
		return ServerTrustGroupsClientListByInstanceResponse{}, err
	}
	return result, nil
}

// NewListByLocationPager - Lists a server trust group.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - locationName - The name of the region where the resource is located.
//   - options - ServerTrustGroupsClientListByLocationOptions contains the optional parameters for the ServerTrustGroupsClient.NewListByLocationPager
//     method.
func (client *ServerTrustGroupsClient) NewListByLocationPager(resourceGroupName string, locationName string, options *ServerTrustGroupsClientListByLocationOptions) *runtime.Pager[ServerTrustGroupsClientListByLocationResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServerTrustGroupsClientListByLocationResponse]{
		More: func(page ServerTrustGroupsClientListByLocationResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServerTrustGroupsClientListByLocationResponse) (ServerTrustGroupsClientListByLocationResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByLocationCreateRequest(ctx, resourceGroupName, locationName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ServerTrustGroupsClientListByLocationResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ServerTrustGroupsClientListByLocationResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServerTrustGroupsClientListByLocationResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByLocationHandleResponse(resp)
		},
	})
}

// listByLocationCreateRequest creates the ListByLocation request.
func (client *ServerTrustGroupsClient) listByLocationCreateRequest(ctx context.Context, resourceGroupName string, locationName string, options *ServerTrustGroupsClientListByLocationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/locations/{locationName}/serverTrustGroups"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
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

// listByLocationHandleResponse handles the ListByLocation response.
func (client *ServerTrustGroupsClient) listByLocationHandleResponse(resp *http.Response) (ServerTrustGroupsClientListByLocationResponse, error) {
	result := ServerTrustGroupsClientListByLocationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerTrustGroupListResult); err != nil {
		return ServerTrustGroupsClientListByLocationResponse{}, err
	}
	return result, nil
}
