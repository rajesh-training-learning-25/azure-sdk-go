//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armextendedlocation

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

// ResourceSyncRulesClient contains the methods for the ResourceSyncRules group.
// Don't use this type directly, use NewResourceSyncRulesClient() instead.
type ResourceSyncRulesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewResourceSyncRulesClient creates a new instance of ResourceSyncRulesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewResourceSyncRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ResourceSyncRulesClient, error) {
	cl, err := arm.NewClient(moduleName+".ResourceSyncRulesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ResourceSyncRulesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a Resource Sync Rule in the parent Custom Location, Subscription Id and Resource
// Group
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-31-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - Custom Locations name.
//   - childResourceName - Resource Sync Rule name.
//   - parameters - Parameters supplied to create or update a Resource Sync Rule.
//   - options - ResourceSyncRulesClientBeginCreateOrUpdateOptions contains the optional parameters for the ResourceSyncRulesClient.BeginCreateOrUpdate
//     method.
func (client *ResourceSyncRulesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, parameters ResourceSyncRule, options *ResourceSyncRulesClientBeginCreateOrUpdateOptions) (*runtime.Poller[ResourceSyncRulesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, resourceName, childResourceName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ResourceSyncRulesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ResourceSyncRulesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates a Resource Sync Rule in the parent Custom Location, Subscription Id and Resource Group
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-31-preview
func (client *ResourceSyncRulesClient) createOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, parameters ResourceSyncRule, options *ResourceSyncRulesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, childResourceName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ResourceSyncRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, parameters ResourceSyncRule, options *ResourceSyncRulesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ExtendedLocation/customLocations/{resourceName}/resourceSyncRules/{childResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if childResourceName == "" {
		return nil, errors.New("parameter childResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{childResourceName}", url.PathEscape(childResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// Delete - Deletes the Resource Sync Rule with the specified Resource Sync Rule Name, Custom Location Resource Name, Resource
// Group, and Subscription Id.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-31-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - Custom Locations name.
//   - childResourceName - Resource Sync Rule name.
//   - options - ResourceSyncRulesClientDeleteOptions contains the optional parameters for the ResourceSyncRulesClient.Delete
//     method.
func (client *ResourceSyncRulesClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, options *ResourceSyncRulesClientDeleteOptions) (ResourceSyncRulesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, childResourceName, options)
	if err != nil {
		return ResourceSyncRulesClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ResourceSyncRulesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ResourceSyncRulesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ResourceSyncRulesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ResourceSyncRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, options *ResourceSyncRulesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ExtendedLocation/customLocations/{resourceName}/resourceSyncRules/{childResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if childResourceName == "" {
		return nil, errors.New("parameter childResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{childResourceName}", url.PathEscape(childResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the details of the resourceSyncRule with a specified resource group, subscription id Custom Location resource
// name and Resource Sync Rule name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-31-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - Custom Locations name.
//   - childResourceName - Resource Sync Rule name.
//   - options - ResourceSyncRulesClientGetOptions contains the optional parameters for the ResourceSyncRulesClient.Get method.
func (client *ResourceSyncRulesClient) Get(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, options *ResourceSyncRulesClientGetOptions) (ResourceSyncRulesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, childResourceName, options)
	if err != nil {
		return ResourceSyncRulesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ResourceSyncRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ResourceSyncRulesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ResourceSyncRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, options *ResourceSyncRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ExtendedLocation/customLocations/{resourceName}/resourceSyncRules/{childResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if childResourceName == "" {
		return nil, errors.New("parameter childResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{childResourceName}", url.PathEscape(childResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ResourceSyncRulesClient) getHandleResponse(resp *http.Response) (ResourceSyncRulesClientGetResponse, error) {
	result := ResourceSyncRulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceSyncRule); err != nil {
		return ResourceSyncRulesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByCustomLocationIDPager - Gets a list of Resource Sync Rules in the specified subscription. The operation returns
// properties of each Resource Sync Rule
//
// Generated from API version 2021-08-31-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - Custom Locations name.
//   - options - ResourceSyncRulesClientListByCustomLocationIDOptions contains the optional parameters for the ResourceSyncRulesClient.NewListByCustomLocationIDPager
//     method.
func (client *ResourceSyncRulesClient) NewListByCustomLocationIDPager(resourceGroupName string, resourceName string, options *ResourceSyncRulesClientListByCustomLocationIDOptions) *runtime.Pager[ResourceSyncRulesClientListByCustomLocationIDResponse] {
	return runtime.NewPager(runtime.PagingHandler[ResourceSyncRulesClientListByCustomLocationIDResponse]{
		More: func(page ResourceSyncRulesClientListByCustomLocationIDResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ResourceSyncRulesClientListByCustomLocationIDResponse) (ResourceSyncRulesClientListByCustomLocationIDResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByCustomLocationIDCreateRequest(ctx, resourceGroupName, resourceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ResourceSyncRulesClientListByCustomLocationIDResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ResourceSyncRulesClientListByCustomLocationIDResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ResourceSyncRulesClientListByCustomLocationIDResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByCustomLocationIDHandleResponse(resp)
		},
	})
}

// listByCustomLocationIDCreateRequest creates the ListByCustomLocationID request.
func (client *ResourceSyncRulesClient) listByCustomLocationIDCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *ResourceSyncRulesClientListByCustomLocationIDOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ExtendedLocation/customLocations/{resourceName}/resourceSyncRules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByCustomLocationIDHandleResponse handles the ListByCustomLocationID response.
func (client *ResourceSyncRulesClient) listByCustomLocationIDHandleResponse(resp *http.Response) (ResourceSyncRulesClientListByCustomLocationIDResponse, error) {
	result := ResourceSyncRulesClientListByCustomLocationIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceSyncRuleListResult); err != nil {
		return ResourceSyncRulesClientListByCustomLocationIDResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates a Resource Sync Rule with the specified Resource Sync Rule name in the specified Resource Group,
// Subscription and Custom Location name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-31-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - Custom Locations name.
//   - childResourceName - Resource Sync Rule name.
//   - parameters - The updatable fields of an existing Resource Sync Rule.
//   - options - ResourceSyncRulesClientBeginUpdateOptions contains the optional parameters for the ResourceSyncRulesClient.BeginUpdate
//     method.
func (client *ResourceSyncRulesClient) BeginUpdate(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, parameters PatchableResourceSyncRule, options *ResourceSyncRulesClientBeginUpdateOptions) (*runtime.Poller[ResourceSyncRulesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, resourceName, childResourceName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ResourceSyncRulesClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ResourceSyncRulesClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Updates a Resource Sync Rule with the specified Resource Sync Rule name in the specified Resource Group, Subscription
// and Custom Location name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-31-preview
func (client *ResourceSyncRulesClient) update(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, parameters PatchableResourceSyncRule, options *ResourceSyncRulesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceName, childResourceName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *ResourceSyncRulesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, parameters PatchableResourceSyncRule, options *ResourceSyncRulesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ExtendedLocation/customLocations/{resourceName}/resourceSyncRules/{childResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if childResourceName == "" {
		return nil, errors.New("parameter childResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{childResourceName}", url.PathEscape(childResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}
