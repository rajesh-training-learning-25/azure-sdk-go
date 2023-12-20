//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmaintenance

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

// ApplyUpdatesClient contains the methods for the ApplyUpdates group.
// Don't use this type directly, use NewApplyUpdatesClient() instead.
type ApplyUpdatesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewApplyUpdatesClient creates a new instance of ApplyUpdatesClient with the specified values.
//   - subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewApplyUpdatesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ApplyUpdatesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ApplyUpdatesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Apply maintenance updates to resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - Resource group name
//   - providerName - Resource provider name
//   - resourceType - Resource type
//   - resourceName - Resource identifier
//   - options - ApplyUpdatesClientCreateOrUpdateOptions contains the optional parameters for the ApplyUpdatesClient.CreateOrUpdate
//     method.
func (client *ApplyUpdatesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, options *ApplyUpdatesClientCreateOrUpdateOptions) (ApplyUpdatesClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "ApplyUpdatesClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, providerName, resourceType, resourceName, options)
	if err != nil {
		return ApplyUpdatesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ApplyUpdatesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ApplyUpdatesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ApplyUpdatesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, options *ApplyUpdatesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/applyUpdates/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	if resourceType == "" {
		return nil, errors.New("parameter resourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", url.PathEscape(resourceType))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ApplyUpdatesClient) createOrUpdateHandleResponse(resp *http.Response) (ApplyUpdatesClientCreateOrUpdateResponse, error) {
	result := ApplyUpdatesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplyUpdate); err != nil {
		return ApplyUpdatesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// CreateOrUpdateParent - Apply maintenance updates to resource with parent
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - Resource group name
//   - providerName - Resource provider name
//   - resourceParentType - Resource parent type
//   - resourceParentName - Resource parent identifier
//   - resourceType - Resource type
//   - resourceName - Resource identifier
//   - options - ApplyUpdatesClientCreateOrUpdateParentOptions contains the optional parameters for the ApplyUpdatesClient.CreateOrUpdateParent
//     method.
func (client *ApplyUpdatesClient) CreateOrUpdateParent(ctx context.Context, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string, options *ApplyUpdatesClientCreateOrUpdateParentOptions) (ApplyUpdatesClientCreateOrUpdateParentResponse, error) {
	var err error
	const operationName = "ApplyUpdatesClient.CreateOrUpdateParent"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateParentCreateRequest(ctx, resourceGroupName, providerName, resourceParentType, resourceParentName, resourceType, resourceName, options)
	if err != nil {
		return ApplyUpdatesClientCreateOrUpdateParentResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ApplyUpdatesClientCreateOrUpdateParentResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ApplyUpdatesClientCreateOrUpdateParentResponse{}, err
	}
	resp, err := client.createOrUpdateParentHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateParentCreateRequest creates the CreateOrUpdateParent request.
func (client *ApplyUpdatesClient) createOrUpdateParentCreateRequest(ctx context.Context, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string, options *ApplyUpdatesClientCreateOrUpdateParentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceParentType}/{resourceParentName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/applyUpdates/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	if resourceParentType == "" {
		return nil, errors.New("parameter resourceParentType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceParentType}", url.PathEscape(resourceParentType))
	if resourceParentName == "" {
		return nil, errors.New("parameter resourceParentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceParentName}", url.PathEscape(resourceParentName))
	if resourceType == "" {
		return nil, errors.New("parameter resourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", url.PathEscape(resourceType))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// createOrUpdateParentHandleResponse handles the CreateOrUpdateParent response.
func (client *ApplyUpdatesClient) createOrUpdateParentHandleResponse(resp *http.Response) (ApplyUpdatesClientCreateOrUpdateParentResponse, error) {
	result := ApplyUpdatesClientCreateOrUpdateParentResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplyUpdate); err != nil {
		return ApplyUpdatesClientCreateOrUpdateParentResponse{}, err
	}
	return result, nil
}

// Get - Track maintenance updates to resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - Resource group name
//   - providerName - Resource provider name
//   - resourceType - Resource type
//   - resourceName - Resource identifier
//   - applyUpdateName - applyUpdate Id
//   - options - ApplyUpdatesClientGetOptions contains the optional parameters for the ApplyUpdatesClient.Get method.
func (client *ApplyUpdatesClient) Get(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, applyUpdateName string, options *ApplyUpdatesClientGetOptions) (ApplyUpdatesClientGetResponse, error) {
	var err error
	const operationName = "ApplyUpdatesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, providerName, resourceType, resourceName, applyUpdateName, options)
	if err != nil {
		return ApplyUpdatesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ApplyUpdatesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ApplyUpdatesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ApplyUpdatesClient) getCreateRequest(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, applyUpdateName string, options *ApplyUpdatesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/applyUpdates/{applyUpdateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	if resourceType == "" {
		return nil, errors.New("parameter resourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", url.PathEscape(resourceType))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if applyUpdateName == "" {
		return nil, errors.New("parameter applyUpdateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applyUpdateName}", url.PathEscape(applyUpdateName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ApplyUpdatesClient) getHandleResponse(resp *http.Response) (ApplyUpdatesClientGetResponse, error) {
	result := ApplyUpdatesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplyUpdate); err != nil {
		return ApplyUpdatesClientGetResponse{}, err
	}
	return result, nil
}

// GetParent - Track maintenance updates to resource with parent
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - Resource group name
//   - providerName - Resource provider name
//   - resourceParentType - Resource parent type
//   - resourceParentName - Resource parent identifier
//   - resourceType - Resource type
//   - resourceName - Resource identifier
//   - applyUpdateName - applyUpdate Id
//   - options - ApplyUpdatesClientGetParentOptions contains the optional parameters for the ApplyUpdatesClient.GetParent method.
func (client *ApplyUpdatesClient) GetParent(ctx context.Context, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string, applyUpdateName string, options *ApplyUpdatesClientGetParentOptions) (ApplyUpdatesClientGetParentResponse, error) {
	var err error
	const operationName = "ApplyUpdatesClient.GetParent"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getParentCreateRequest(ctx, resourceGroupName, providerName, resourceParentType, resourceParentName, resourceType, resourceName, applyUpdateName, options)
	if err != nil {
		return ApplyUpdatesClientGetParentResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ApplyUpdatesClientGetParentResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ApplyUpdatesClientGetParentResponse{}, err
	}
	resp, err := client.getParentHandleResponse(httpResp)
	return resp, err
}

// getParentCreateRequest creates the GetParent request.
func (client *ApplyUpdatesClient) getParentCreateRequest(ctx context.Context, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string, applyUpdateName string, options *ApplyUpdatesClientGetParentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceParentType}/{resourceParentName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/applyUpdates/{applyUpdateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	if resourceParentType == "" {
		return nil, errors.New("parameter resourceParentType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceParentType}", url.PathEscape(resourceParentType))
	if resourceParentName == "" {
		return nil, errors.New("parameter resourceParentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceParentName}", url.PathEscape(resourceParentName))
	if resourceType == "" {
		return nil, errors.New("parameter resourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", url.PathEscape(resourceType))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if applyUpdateName == "" {
		return nil, errors.New("parameter applyUpdateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applyUpdateName}", url.PathEscape(applyUpdateName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getParentHandleResponse handles the GetParent response.
func (client *ApplyUpdatesClient) getParentHandleResponse(resp *http.Response) (ApplyUpdatesClientGetParentResponse, error) {
	result := ApplyUpdatesClientGetParentResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplyUpdate); err != nil {
		return ApplyUpdatesClientGetParentResponse{}, err
	}
	return result, nil
}

// NewListPager - Get Configuration records within a subscription
//
// Generated from API version 2023-04-01
//   - options - ApplyUpdatesClientListOptions contains the optional parameters for the ApplyUpdatesClient.NewListPager method.
func (client *ApplyUpdatesClient) NewListPager(options *ApplyUpdatesClientListOptions) *runtime.Pager[ApplyUpdatesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ApplyUpdatesClientListResponse]{
		More: func(page ApplyUpdatesClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ApplyUpdatesClientListResponse) (ApplyUpdatesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ApplyUpdatesClient.NewListPager")
			req, err := client.listCreateRequest(ctx, options)
			if err != nil {
				return ApplyUpdatesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ApplyUpdatesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ApplyUpdatesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ApplyUpdatesClient) listCreateRequest(ctx context.Context, options *ApplyUpdatesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Maintenance/applyUpdates"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ApplyUpdatesClient) listHandleResponse(resp *http.Response) (ApplyUpdatesClientListResponse, error) {
	result := ApplyUpdatesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListApplyUpdate); err != nil {
		return ApplyUpdatesClientListResponse{}, err
	}
	return result, nil
}
