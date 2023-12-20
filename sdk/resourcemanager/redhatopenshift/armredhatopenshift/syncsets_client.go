//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armredhatopenshift

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

// SyncSetsClient contains the methods for the SyncSets group.
// Don't use this type directly, use NewSyncSetsClient() instead.
type SyncSetsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSyncSetsClient creates a new instance of SyncSetsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSyncSetsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SyncSetsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SyncSetsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - The operation returns properties of a SyncSet.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-04
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the OpenShift cluster resource.
//   - childResourceName - The name of the SyncSet resource.
//   - parameters - The SyncSet resource.
//   - options - SyncSetsClientCreateOrUpdateOptions contains the optional parameters for the SyncSetsClient.CreateOrUpdate method.
func (client *SyncSetsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, parameters SyncSet, options *SyncSetsClientCreateOrUpdateOptions) (SyncSetsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "SyncSetsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, childResourceName, parameters, options)
	if err != nil {
		return SyncSetsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SyncSetsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return SyncSetsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SyncSetsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, parameters SyncSet, options *SyncSetsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RedHatOpenShift/openshiftclusters/{resourceName}/syncSet/{childResourceName}"
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
	reqQP.Set("api-version", "2023-09-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *SyncSetsClient) createOrUpdateHandleResponse(resp *http.Response) (SyncSetsClientCreateOrUpdateResponse, error) {
	result := SyncSetsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SyncSet); err != nil {
		return SyncSetsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - The operation returns nothing.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-04
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the OpenShift cluster resource.
//   - childResourceName - The name of the SyncSet resource.
//   - options - SyncSetsClientDeleteOptions contains the optional parameters for the SyncSetsClient.Delete method.
func (client *SyncSetsClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, options *SyncSetsClientDeleteOptions) (SyncSetsClientDeleteResponse, error) {
	var err error
	const operationName = "SyncSetsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, childResourceName, options)
	if err != nil {
		return SyncSetsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SyncSetsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return SyncSetsClientDeleteResponse{}, err
	}
	return SyncSetsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SyncSetsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, options *SyncSetsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RedHatOpenShift/openshiftclusters/{resourceName}/syncSet/{childResourceName}"
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
	reqQP.Set("api-version", "2023-09-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - The operation returns properties of a SyncSet.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-04
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the OpenShift cluster resource.
//   - childResourceName - The name of the SyncSet resource.
//   - options - SyncSetsClientGetOptions contains the optional parameters for the SyncSetsClient.Get method.
func (client *SyncSetsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, options *SyncSetsClientGetOptions) (SyncSetsClientGetResponse, error) {
	var err error
	const operationName = "SyncSetsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, childResourceName, options)
	if err != nil {
		return SyncSetsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SyncSetsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SyncSetsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SyncSetsClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, options *SyncSetsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RedHatOpenShift/openshiftclusters/{resourceName}/syncSet/{childResourceName}"
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
	reqQP.Set("api-version", "2023-09-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SyncSetsClient) getHandleResponse(resp *http.Response) (SyncSetsClientGetResponse, error) {
	result := SyncSetsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SyncSet); err != nil {
		return SyncSetsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - The operation returns properties of each SyncSet.
//
// Generated from API version 2023-09-04
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the OpenShift cluster resource.
//   - options - SyncSetsClientListOptions contains the optional parameters for the SyncSetsClient.NewListPager method.
func (client *SyncSetsClient) NewListPager(resourceGroupName string, resourceName string, options *SyncSetsClientListOptions) *runtime.Pager[SyncSetsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SyncSetsClientListResponse]{
		More: func(page SyncSetsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SyncSetsClientListResponse) (SyncSetsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SyncSetsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, resourceName, options)
			}, nil)
			if err != nil {
				return SyncSetsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *SyncSetsClient) listCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *SyncSetsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RedHatOpenShift/openShiftCluster/{resourceName}/syncSets"
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
	reqQP.Set("api-version", "2023-09-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SyncSetsClient) listHandleResponse(resp *http.Response) (SyncSetsClientListResponse, error) {
	result := SyncSetsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SyncSetList); err != nil {
		return SyncSetsClientListResponse{}, err
	}
	return result, nil
}

// Update - The operation returns properties of a SyncSet.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-04
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the OpenShift cluster resource.
//   - childResourceName - The name of the SyncSet resource.
//   - parameters - The SyncSet resource.
//   - options - SyncSetsClientUpdateOptions contains the optional parameters for the SyncSetsClient.Update method.
func (client *SyncSetsClient) Update(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, parameters SyncSetUpdate, options *SyncSetsClientUpdateOptions) (SyncSetsClientUpdateResponse, error) {
	var err error
	const operationName = "SyncSetsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceName, childResourceName, parameters, options)
	if err != nil {
		return SyncSetsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SyncSetsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SyncSetsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *SyncSetsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, parameters SyncSetUpdate, options *SyncSetsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RedHatOpenShift/openshiftclusters/{resourceName}/syncSet/{childResourceName}"
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
	reqQP.Set("api-version", "2023-09-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *SyncSetsClient) updateHandleResponse(resp *http.Response) (SyncSetsClientUpdateResponse, error) {
	result := SyncSetsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SyncSet); err != nil {
		return SyncSetsClientUpdateResponse{}, err
	}
	return result, nil
}
