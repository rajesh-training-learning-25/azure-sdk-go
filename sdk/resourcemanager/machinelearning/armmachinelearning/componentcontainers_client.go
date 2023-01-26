//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmachinelearning

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

// ComponentContainersClient contains the methods for the ComponentContainers group.
// Don't use this type directly, use NewComponentContainersClient() instead.
type ComponentContainersClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewComponentContainersClient creates a new instance of ComponentContainersClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewComponentContainersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ComponentContainersClient, error) {
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
	client := &ComponentContainersClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update container.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - Name of Azure Machine Learning workspace.
// name - Container name.
// body - Container entity to create or update.
// options - ComponentContainersClientCreateOrUpdateOptions contains the optional parameters for the ComponentContainersClient.CreateOrUpdate
// method.
func (client *ComponentContainersClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, name string, body ComponentContainer, options *ComponentContainersClientCreateOrUpdateOptions) (ComponentContainersClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, name, body, options)
	if err != nil {
		return ComponentContainersClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ComponentContainersClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ComponentContainersClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ComponentContainersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, body ComponentContainer, options *ComponentContainersClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/components/{name}"
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
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ComponentContainersClient) createOrUpdateHandleResponse(resp *http.Response) (ComponentContainersClientCreateOrUpdateResponse, error) {
	result := ComponentContainersClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentContainer); err != nil {
		return ComponentContainersClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete container.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - Name of Azure Machine Learning workspace.
// name - Container name.
// options - ComponentContainersClientDeleteOptions contains the optional parameters for the ComponentContainersClient.Delete
// method.
func (client *ComponentContainersClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *ComponentContainersClientDeleteOptions) (ComponentContainersClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, name, options)
	if err != nil {
		return ComponentContainersClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ComponentContainersClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ComponentContainersClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ComponentContainersClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ComponentContainersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *ComponentContainersClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/components/{name}"
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
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get container.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - Name of Azure Machine Learning workspace.
// name - Container name.
// options - ComponentContainersClientGetOptions contains the optional parameters for the ComponentContainersClient.Get method.
func (client *ComponentContainersClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *ComponentContainersClientGetOptions) (ComponentContainersClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, name, options)
	if err != nil {
		return ComponentContainersClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ComponentContainersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ComponentContainersClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ComponentContainersClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *ComponentContainersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/components/{name}"
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
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ComponentContainersClient) getHandleResponse(resp *http.Response) (ComponentContainersClientGetResponse, error) {
	result := ComponentContainersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentContainer); err != nil {
		return ComponentContainersClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List component containers.
// Generated from API version 2022-10-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - Name of Azure Machine Learning workspace.
// options - ComponentContainersClientListOptions contains the optional parameters for the ComponentContainersClient.List
// method.
func (client *ComponentContainersClient) NewListPager(resourceGroupName string, workspaceName string, options *ComponentContainersClientListOptions) *runtime.Pager[ComponentContainersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ComponentContainersClientListResponse]{
		More: func(page ComponentContainersClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ComponentContainersClientListResponse) (ComponentContainersClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ComponentContainersClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ComponentContainersClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ComponentContainersClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ComponentContainersClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *ComponentContainersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/components"
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
	reqQP.Set("api-version", "2022-10-01")
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", *options.Skip)
	}
	if options != nil && options.ListViewType != nil {
		reqQP.Set("listViewType", string(*options.ListViewType))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ComponentContainersClient) listHandleResponse(resp *http.Response) (ComponentContainersClientListResponse, error) {
	result := ComponentContainersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentContainerResourceArmPaginatedResult); err != nil {
		return ComponentContainersClientListResponse{}, err
	}
	return result, nil
}
