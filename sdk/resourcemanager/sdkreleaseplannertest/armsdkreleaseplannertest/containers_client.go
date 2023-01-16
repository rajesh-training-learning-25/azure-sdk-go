//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsdkreleaseplannertest

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

// ContainersClient contains the methods for the Containers group.
// Don't use this type directly, use NewContainersClient() instead.
type ContainersClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewContainersClient creates a new instance of ContainersClient with the specified values.
//   - subscriptionID - The subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewContainersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ContainersClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ContainersClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a new container or updates an existing container on the device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-01
//   - deviceName - The device name.
//   - storageAccountName - The Storage Account Name
//   - containerName - The container name.
//   - resourceGroupName - The resource group name.
//   - containerParam - The container properties.
//   - options - ContainersClientBeginCreateOrUpdateOptions contains the optional parameters for the ContainersClient.BeginCreateOrUpdate
//     method.
func (client *ContainersClient) BeginCreateOrUpdate(ctx context.Context, deviceName string, storageAccountName string, containerName string, resourceGroupName string, containerParam Container, options *ContainersClientBeginCreateOrUpdateOptions) (*runtime.Poller[ContainersClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, deviceName, storageAccountName, containerName, resourceGroupName, containerParam, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ContainersClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ContainersClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates a new container or updates an existing container on the device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-01
func (client *ContainersClient) createOrUpdate(ctx context.Context, deviceName string, storageAccountName string, containerName string, resourceGroupName string, containerParam Container, options *ContainersClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, deviceName, storageAccountName, containerName, resourceGroupName, containerParam, options)
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
func (client *ContainersClient) createOrUpdateCreateRequest(ctx context.Context, deviceName string, storageAccountName string, containerName string, resourceGroupName string, containerParam Container, options *ContainersClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccounts/{storageAccountName}/containers/{containerName}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if storageAccountName == "" {
		return nil, errors.New("parameter storageAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageAccountName}", url.PathEscape(storageAccountName))
	if containerName == "" {
		return nil, errors.New("parameter containerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerName}", url.PathEscape(containerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, containerParam)
}

// BeginDelete - Deletes the container on the Data Box Edge/Data Box Gateway device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-01
//   - deviceName - The device name.
//   - storageAccountName - The Storage Account Name
//   - containerName - The container name.
//   - resourceGroupName - The resource group name.
//   - options - ContainersClientBeginDeleteOptions contains the optional parameters for the ContainersClient.BeginDelete method.
func (client *ContainersClient) BeginDelete(ctx context.Context, deviceName string, storageAccountName string, containerName string, resourceGroupName string, options *ContainersClientBeginDeleteOptions) (*runtime.Poller[ContainersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, deviceName, storageAccountName, containerName, resourceGroupName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ContainersClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ContainersClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the container on the Data Box Edge/Data Box Gateway device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-01
func (client *ContainersClient) deleteOperation(ctx context.Context, deviceName string, storageAccountName string, containerName string, resourceGroupName string, options *ContainersClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, deviceName, storageAccountName, containerName, resourceGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ContainersClient) deleteCreateRequest(ctx context.Context, deviceName string, storageAccountName string, containerName string, resourceGroupName string, options *ContainersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccounts/{storageAccountName}/containers/{containerName}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if storageAccountName == "" {
		return nil, errors.New("parameter storageAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageAccountName}", url.PathEscape(storageAccountName))
	if containerName == "" {
		return nil, errors.New("parameter containerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerName}", url.PathEscape(containerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a container by name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-01
//   - deviceName - The device name.
//   - storageAccountName - The Storage Account Name
//   - containerName - The container Name
//   - resourceGroupName - The resource group name.
//   - options - ContainersClientGetOptions contains the optional parameters for the ContainersClient.Get method.
func (client *ContainersClient) Get(ctx context.Context, deviceName string, storageAccountName string, containerName string, resourceGroupName string, options *ContainersClientGetOptions) (ContainersClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, deviceName, storageAccountName, containerName, resourceGroupName, options)
	if err != nil {
		return ContainersClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContainersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ContainersClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ContainersClient) getCreateRequest(ctx context.Context, deviceName string, storageAccountName string, containerName string, resourceGroupName string, options *ContainersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccounts/{storageAccountName}/containers/{containerName}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if storageAccountName == "" {
		return nil, errors.New("parameter storageAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageAccountName}", url.PathEscape(storageAccountName))
	if containerName == "" {
		return nil, errors.New("parameter containerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerName}", url.PathEscape(containerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ContainersClient) getHandleResponse(resp *http.Response) (ContainersClientGetResponse, error) {
	result := ContainersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Container); err != nil {
		return ContainersClientGetResponse{}, err
	}
	return result, nil
}

// NewListByStorageAccountPager - Lists all the containers of a storage Account in a Data Box Edge/Data Box Gateway device.
//
// Generated from API version 2022-03-01
//   - deviceName - The device name.
//   - storageAccountName - The storage Account name.
//   - resourceGroupName - The resource group name.
//   - options - ContainersClientListByStorageAccountOptions contains the optional parameters for the ContainersClient.NewListByStorageAccountPager
//     method.
func (client *ContainersClient) NewListByStorageAccountPager(deviceName string, storageAccountName string, resourceGroupName string, options *ContainersClientListByStorageAccountOptions) *runtime.Pager[ContainersClientListByStorageAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[ContainersClientListByStorageAccountResponse]{
		More: func(page ContainersClientListByStorageAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ContainersClientListByStorageAccountResponse) (ContainersClientListByStorageAccountResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByStorageAccountCreateRequest(ctx, deviceName, storageAccountName, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ContainersClientListByStorageAccountResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ContainersClientListByStorageAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ContainersClientListByStorageAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByStorageAccountHandleResponse(resp)
		},
	})
}

// listByStorageAccountCreateRequest creates the ListByStorageAccount request.
func (client *ContainersClient) listByStorageAccountCreateRequest(ctx context.Context, deviceName string, storageAccountName string, resourceGroupName string, options *ContainersClientListByStorageAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccounts/{storageAccountName}/containers"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if storageAccountName == "" {
		return nil, errors.New("parameter storageAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageAccountName}", url.PathEscape(storageAccountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByStorageAccountHandleResponse handles the ListByStorageAccount response.
func (client *ContainersClient) listByStorageAccountHandleResponse(resp *http.Response) (ContainersClientListByStorageAccountResponse, error) {
	result := ContainersClientListByStorageAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContainerList); err != nil {
		return ContainersClientListByStorageAccountResponse{}, err
	}
	return result, nil
}

// BeginRefresh - Refreshes the container metadata with the data from the cloud.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-01
//   - deviceName - The device name.
//   - storageAccountName - The Storage Account Name
//   - containerName - The container name.
//   - resourceGroupName - The resource group name.
//   - options - ContainersClientBeginRefreshOptions contains the optional parameters for the ContainersClient.BeginRefresh method.
func (client *ContainersClient) BeginRefresh(ctx context.Context, deviceName string, storageAccountName string, containerName string, resourceGroupName string, options *ContainersClientBeginRefreshOptions) (*runtime.Poller[ContainersClientRefreshResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.refresh(ctx, deviceName, storageAccountName, containerName, resourceGroupName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ContainersClientRefreshResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ContainersClientRefreshResponse](options.ResumeToken, client.pl, nil)
	}
}

// Refresh - Refreshes the container metadata with the data from the cloud.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-01
func (client *ContainersClient) refresh(ctx context.Context, deviceName string, storageAccountName string, containerName string, resourceGroupName string, options *ContainersClientBeginRefreshOptions) (*http.Response, error) {
	req, err := client.refreshCreateRequest(ctx, deviceName, storageAccountName, containerName, resourceGroupName, options)
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

// refreshCreateRequest creates the Refresh request.
func (client *ContainersClient) refreshCreateRequest(ctx context.Context, deviceName string, storageAccountName string, containerName string, resourceGroupName string, options *ContainersClientBeginRefreshOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccounts/{storageAccountName}/containers/{containerName}/refresh"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if storageAccountName == "" {
		return nil, errors.New("parameter storageAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageAccountName}", url.PathEscape(storageAccountName))
	if containerName == "" {
		return nil, errors.New("parameter containerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerName}", url.PathEscape(containerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
