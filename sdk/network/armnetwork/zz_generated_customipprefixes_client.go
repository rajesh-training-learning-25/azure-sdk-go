//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// CustomIPPrefixesClient contains the methods for the CustomIPPrefixes group.
// Don't use this type directly, use NewCustomIPPrefixesClient() instead.
type CustomIPPrefixesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewCustomIPPrefixesClient creates a new instance of CustomIPPrefixesClient with the specified values.
func NewCustomIPPrefixesClient(con *arm.Connection, subscriptionID string) *CustomIPPrefixesClient {
	return &CustomIPPrefixesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a custom IP prefix.
// If the operation fails it returns the *CloudError error type.
func (client *CustomIPPrefixesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, customIPPrefixName string, parameters CustomIPPrefix, options *CustomIPPrefixesBeginCreateOrUpdateOptions) (CustomIPPrefixesCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, customIPPrefixName, parameters, options)
	if err != nil {
		return CustomIPPrefixesCreateOrUpdatePollerResponse{}, err
	}
	result := CustomIPPrefixesCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("CustomIPPrefixesClient.CreateOrUpdate", "location", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return CustomIPPrefixesCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &CustomIPPrefixesCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a custom IP prefix.
// If the operation fails it returns the *CloudError error type.
func (client *CustomIPPrefixesClient) createOrUpdate(ctx context.Context, resourceGroupName string, customIPPrefixName string, parameters CustomIPPrefix, options *CustomIPPrefixesBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, customIPPrefixName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *CustomIPPrefixesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, customIPPrefixName string, parameters CustomIPPrefix, options *CustomIPPrefixesBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/customIpPrefixes/{customIpPrefixName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if customIPPrefixName == "" {
		return nil, errors.New("parameter customIPPrefixName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customIpPrefixName}", url.PathEscape(customIPPrefixName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *CustomIPPrefixesClient) createOrUpdateHandleError(resp *http.Response) error {
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

// BeginDelete - Deletes the specified custom IP prefix.
// If the operation fails it returns the *CloudError error type.
func (client *CustomIPPrefixesClient) BeginDelete(ctx context.Context, resourceGroupName string, customIPPrefixName string, options *CustomIPPrefixesBeginDeleteOptions) (CustomIPPrefixesDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, customIPPrefixName, options)
	if err != nil {
		return CustomIPPrefixesDeletePollerResponse{}, err
	}
	result := CustomIPPrefixesDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("CustomIPPrefixesClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return CustomIPPrefixesDeletePollerResponse{}, err
	}
	result.Poller = &CustomIPPrefixesDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified custom IP prefix.
// If the operation fails it returns the *CloudError error type.
func (client *CustomIPPrefixesClient) deleteOperation(ctx context.Context, resourceGroupName string, customIPPrefixName string, options *CustomIPPrefixesBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, customIPPrefixName, options)
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
func (client *CustomIPPrefixesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, customIPPrefixName string, options *CustomIPPrefixesBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/customIpPrefixes/{customIpPrefixName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if customIPPrefixName == "" {
		return nil, errors.New("parameter customIPPrefixName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customIpPrefixName}", url.PathEscape(customIPPrefixName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *CustomIPPrefixesClient) deleteHandleError(resp *http.Response) error {
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

// Get - Gets the specified custom IP prefix in a specified resource group.
// If the operation fails it returns the *CloudError error type.
func (client *CustomIPPrefixesClient) Get(ctx context.Context, resourceGroupName string, customIPPrefixName string, options *CustomIPPrefixesGetOptions) (CustomIPPrefixesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, customIPPrefixName, options)
	if err != nil {
		return CustomIPPrefixesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CustomIPPrefixesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CustomIPPrefixesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CustomIPPrefixesClient) getCreateRequest(ctx context.Context, resourceGroupName string, customIPPrefixName string, options *CustomIPPrefixesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/customIpPrefixes/{customIpPrefixName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if customIPPrefixName == "" {
		return nil, errors.New("parameter customIPPrefixName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customIpPrefixName}", url.PathEscape(customIPPrefixName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CustomIPPrefixesClient) getHandleResponse(resp *http.Response) (CustomIPPrefixesGetResponse, error) {
	result := CustomIPPrefixesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomIPPrefix); err != nil {
		return CustomIPPrefixesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *CustomIPPrefixesClient) getHandleError(resp *http.Response) error {
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

// List - Gets all custom IP prefixes in a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *CustomIPPrefixesClient) List(resourceGroupName string, options *CustomIPPrefixesListOptions) *CustomIPPrefixesListPager {
	return &CustomIPPrefixesListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp CustomIPPrefixesListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.CustomIPPrefixListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *CustomIPPrefixesClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *CustomIPPrefixesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/customIpPrefixes"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *CustomIPPrefixesClient) listHandleResponse(resp *http.Response) (CustomIPPrefixesListResponse, error) {
	result := CustomIPPrefixesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomIPPrefixListResult); err != nil {
		return CustomIPPrefixesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *CustomIPPrefixesClient) listHandleError(resp *http.Response) error {
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

// ListAll - Gets all the custom IP prefixes in a subscription.
// If the operation fails it returns the *CloudError error type.
func (client *CustomIPPrefixesClient) ListAll(options *CustomIPPrefixesListAllOptions) *CustomIPPrefixesListAllPager {
	return &CustomIPPrefixesListAllPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listAllCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp CustomIPPrefixesListAllResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.CustomIPPrefixListResult.NextLink)
		},
	}
}

// listAllCreateRequest creates the ListAll request.
func (client *CustomIPPrefixesClient) listAllCreateRequest(ctx context.Context, options *CustomIPPrefixesListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/customIpPrefixes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *CustomIPPrefixesClient) listAllHandleResponse(resp *http.Response) (CustomIPPrefixesListAllResponse, error) {
	result := CustomIPPrefixesListAllResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomIPPrefixListResult); err != nil {
		return CustomIPPrefixesListAllResponse{}, err
	}
	return result, nil
}

// listAllHandleError handles the ListAll error response.
func (client *CustomIPPrefixesClient) listAllHandleError(resp *http.Response) error {
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

// UpdateTags - Updates custom IP prefix tags.
// If the operation fails it returns the *CloudError error type.
func (client *CustomIPPrefixesClient) UpdateTags(ctx context.Context, resourceGroupName string, customIPPrefixName string, parameters TagsObject, options *CustomIPPrefixesUpdateTagsOptions) (CustomIPPrefixesUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, customIPPrefixName, parameters, options)
	if err != nil {
		return CustomIPPrefixesUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CustomIPPrefixesUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CustomIPPrefixesUpdateTagsResponse{}, client.updateTagsHandleError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *CustomIPPrefixesClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, customIPPrefixName string, parameters TagsObject, options *CustomIPPrefixesUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/customIpPrefixes/{customIpPrefixName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if customIPPrefixName == "" {
		return nil, errors.New("parameter customIPPrefixName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customIpPrefixName}", url.PathEscape(customIPPrefixName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *CustomIPPrefixesClient) updateTagsHandleResponse(resp *http.Response) (CustomIPPrefixesUpdateTagsResponse, error) {
	result := CustomIPPrefixesUpdateTagsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomIPPrefix); err != nil {
		return CustomIPPrefixesUpdateTagsResponse{}, err
	}
	return result, nil
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *CustomIPPrefixesClient) updateTagsHandleError(resp *http.Response) error {
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
