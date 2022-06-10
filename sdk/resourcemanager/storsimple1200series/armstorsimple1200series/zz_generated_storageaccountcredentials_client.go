//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorsimple1200series

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

// StorageAccountCredentialsClient contains the methods for the StorageAccountCredentials group.
// Don't use this type directly, use NewStorageAccountCredentialsClient() instead.
type StorageAccountCredentialsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewStorageAccountCredentialsClient creates a new instance of StorageAccountCredentialsClient with the specified values.
// subscriptionID - The subscription id
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewStorageAccountCredentialsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*StorageAccountCredentialsClient, error) {
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
	client := &StorageAccountCredentialsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates the storage account credential
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2016-10-01
// credentialName - The credential name.
// resourceGroupName - The resource group name
// managerName - The manager name
// storageAccount - The storage account credential to be added or updated.
// options - StorageAccountCredentialsClientBeginCreateOrUpdateOptions contains the optional parameters for the StorageAccountCredentialsClient.BeginCreateOrUpdate
// method.
func (client *StorageAccountCredentialsClient) BeginCreateOrUpdate(ctx context.Context, credentialName string, resourceGroupName string, managerName string, storageAccount StorageAccountCredential, options *StorageAccountCredentialsClientBeginCreateOrUpdateOptions) (*runtime.Poller[StorageAccountCredentialsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, credentialName, resourceGroupName, managerName, storageAccount, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[StorageAccountCredentialsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[StorageAccountCredentialsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates the storage account credential
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2016-10-01
func (client *StorageAccountCredentialsClient) createOrUpdate(ctx context.Context, credentialName string, resourceGroupName string, managerName string, storageAccount StorageAccountCredential, options *StorageAccountCredentialsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, credentialName, resourceGroupName, managerName, storageAccount, options)
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
func (client *StorageAccountCredentialsClient) createOrUpdateCreateRequest(ctx context.Context, credentialName string, resourceGroupName string, managerName string, storageAccount StorageAccountCredential, options *StorageAccountCredentialsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageAccountCredentials/{credentialName}"
	if credentialName == "" {
		return nil, errors.New("parameter credentialName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{credentialName}", url.PathEscape(credentialName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managerName == "" {
		return nil, errors.New("parameter managerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", url.PathEscape(managerName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, storageAccount)
}

// BeginDelete - Deletes the storage account credential
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2016-10-01
// credentialName - The name of the storage account credential.
// resourceGroupName - The resource group name
// managerName - The manager name
// options - StorageAccountCredentialsClientBeginDeleteOptions contains the optional parameters for the StorageAccountCredentialsClient.BeginDelete
// method.
func (client *StorageAccountCredentialsClient) BeginDelete(ctx context.Context, credentialName string, resourceGroupName string, managerName string, options *StorageAccountCredentialsClientBeginDeleteOptions) (*runtime.Poller[StorageAccountCredentialsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, credentialName, resourceGroupName, managerName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[StorageAccountCredentialsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[StorageAccountCredentialsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the storage account credential
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2016-10-01
func (client *StorageAccountCredentialsClient) deleteOperation(ctx context.Context, credentialName string, resourceGroupName string, managerName string, options *StorageAccountCredentialsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, credentialName, resourceGroupName, managerName, options)
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
func (client *StorageAccountCredentialsClient) deleteCreateRequest(ctx context.Context, credentialName string, resourceGroupName string, managerName string, options *StorageAccountCredentialsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageAccountCredentials/{credentialName}"
	if credentialName == "" {
		return nil, errors.New("parameter credentialName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{credentialName}", url.PathEscape(credentialName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managerName == "" {
		return nil, errors.New("parameter managerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", url.PathEscape(managerName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Returns the properties of the specified storage account credential name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2016-10-01
// credentialName - The name of storage account credential to be fetched.
// resourceGroupName - The resource group name
// managerName - The manager name
// options - StorageAccountCredentialsClientGetOptions contains the optional parameters for the StorageAccountCredentialsClient.Get
// method.
func (client *StorageAccountCredentialsClient) Get(ctx context.Context, credentialName string, resourceGroupName string, managerName string, options *StorageAccountCredentialsClientGetOptions) (StorageAccountCredentialsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, credentialName, resourceGroupName, managerName, options)
	if err != nil {
		return StorageAccountCredentialsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StorageAccountCredentialsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return StorageAccountCredentialsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *StorageAccountCredentialsClient) getCreateRequest(ctx context.Context, credentialName string, resourceGroupName string, managerName string, options *StorageAccountCredentialsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageAccountCredentials/{credentialName}"
	if credentialName == "" {
		return nil, errors.New("parameter credentialName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{credentialName}", url.PathEscape(credentialName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managerName == "" {
		return nil, errors.New("parameter managerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", url.PathEscape(managerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *StorageAccountCredentialsClient) getHandleResponse(resp *http.Response) (StorageAccountCredentialsClientGetResponse, error) {
	result := StorageAccountCredentialsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageAccountCredential); err != nil {
		return StorageAccountCredentialsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByManagerPager - Retrieves all the storage account credentials in a manager.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2016-10-01
// resourceGroupName - The resource group name
// managerName - The manager name
// options - StorageAccountCredentialsClientListByManagerOptions contains the optional parameters for the StorageAccountCredentialsClient.ListByManager
// method.
func (client *StorageAccountCredentialsClient) NewListByManagerPager(resourceGroupName string, managerName string, options *StorageAccountCredentialsClientListByManagerOptions) *runtime.Pager[StorageAccountCredentialsClientListByManagerResponse] {
	return runtime.NewPager(runtime.PagingHandler[StorageAccountCredentialsClientListByManagerResponse]{
		More: func(page StorageAccountCredentialsClientListByManagerResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *StorageAccountCredentialsClientListByManagerResponse) (StorageAccountCredentialsClientListByManagerResponse, error) {
			req, err := client.listByManagerCreateRequest(ctx, resourceGroupName, managerName, options)
			if err != nil {
				return StorageAccountCredentialsClientListByManagerResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return StorageAccountCredentialsClientListByManagerResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return StorageAccountCredentialsClientListByManagerResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByManagerHandleResponse(resp)
		},
	})
}

// listByManagerCreateRequest creates the ListByManager request.
func (client *StorageAccountCredentialsClient) listByManagerCreateRequest(ctx context.Context, resourceGroupName string, managerName string, options *StorageAccountCredentialsClientListByManagerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageAccountCredentials"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managerName == "" {
		return nil, errors.New("parameter managerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", url.PathEscape(managerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByManagerHandleResponse handles the ListByManager response.
func (client *StorageAccountCredentialsClient) listByManagerHandleResponse(resp *http.Response) (StorageAccountCredentialsClientListByManagerResponse, error) {
	result := StorageAccountCredentialsClientListByManagerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageAccountCredentialList); err != nil {
		return StorageAccountCredentialsClientListByManagerResponse{}, err
	}
	return result, nil
}
