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

// PublicIPPrefixesClient contains the methods for the PublicIPPrefixes group.
// Don't use this type directly, use NewPublicIPPrefixesClient() instead.
type PublicIPPrefixesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewPublicIPPrefixesClient creates a new instance of PublicIPPrefixesClient with the specified values.
func NewPublicIPPrefixesClient(con *arm.Connection, subscriptionID string) *PublicIPPrefixesClient {
	return &PublicIPPrefixesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a static or dynamic public IP prefix.
// If the operation fails it returns the *CloudError error type.
func (client *PublicIPPrefixesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, publicIPPrefixName string, parameters PublicIPPrefix, options *PublicIPPrefixesBeginCreateOrUpdateOptions) (PublicIPPrefixesCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, publicIPPrefixName, parameters, options)
	if err != nil {
		return PublicIPPrefixesCreateOrUpdatePollerResponse{}, err
	}
	result := PublicIPPrefixesCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("PublicIPPrefixesClient.CreateOrUpdate", "location", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return PublicIPPrefixesCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &PublicIPPrefixesCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a static or dynamic public IP prefix.
// If the operation fails it returns the *CloudError error type.
func (client *PublicIPPrefixesClient) createOrUpdate(ctx context.Context, resourceGroupName string, publicIPPrefixName string, parameters PublicIPPrefix, options *PublicIPPrefixesBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, publicIPPrefixName, parameters, options)
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
func (client *PublicIPPrefixesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, publicIPPrefixName string, parameters PublicIPPrefix, options *PublicIPPrefixesBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIpPrefixName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publicIPPrefixName == "" {
		return nil, errors.New("parameter publicIPPrefixName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publicIpPrefixName}", url.PathEscape(publicIPPrefixName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *PublicIPPrefixesClient) createOrUpdateHandleError(resp *http.Response) error {
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

// BeginDelete - Deletes the specified public IP prefix.
// If the operation fails it returns the *CloudError error type.
func (client *PublicIPPrefixesClient) BeginDelete(ctx context.Context, resourceGroupName string, publicIPPrefixName string, options *PublicIPPrefixesBeginDeleteOptions) (PublicIPPrefixesDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, publicIPPrefixName, options)
	if err != nil {
		return PublicIPPrefixesDeletePollerResponse{}, err
	}
	result := PublicIPPrefixesDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("PublicIPPrefixesClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return PublicIPPrefixesDeletePollerResponse{}, err
	}
	result.Poller = &PublicIPPrefixesDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified public IP prefix.
// If the operation fails it returns the *CloudError error type.
func (client *PublicIPPrefixesClient) deleteOperation(ctx context.Context, resourceGroupName string, publicIPPrefixName string, options *PublicIPPrefixesBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, publicIPPrefixName, options)
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
func (client *PublicIPPrefixesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, publicIPPrefixName string, options *PublicIPPrefixesBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIpPrefixName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publicIPPrefixName == "" {
		return nil, errors.New("parameter publicIPPrefixName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publicIpPrefixName}", url.PathEscape(publicIPPrefixName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *PublicIPPrefixesClient) deleteHandleError(resp *http.Response) error {
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

// Get - Gets the specified public IP prefix in a specified resource group.
// If the operation fails it returns the *CloudError error type.
func (client *PublicIPPrefixesClient) Get(ctx context.Context, resourceGroupName string, publicIPPrefixName string, options *PublicIPPrefixesGetOptions) (PublicIPPrefixesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, publicIPPrefixName, options)
	if err != nil {
		return PublicIPPrefixesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PublicIPPrefixesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PublicIPPrefixesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PublicIPPrefixesClient) getCreateRequest(ctx context.Context, resourceGroupName string, publicIPPrefixName string, options *PublicIPPrefixesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIpPrefixName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publicIPPrefixName == "" {
		return nil, errors.New("parameter publicIPPrefixName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publicIpPrefixName}", url.PathEscape(publicIPPrefixName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PublicIPPrefixesClient) getHandleResponse(resp *http.Response) (PublicIPPrefixesGetResponse, error) {
	result := PublicIPPrefixesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PublicIPPrefix); err != nil {
		return PublicIPPrefixesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *PublicIPPrefixesClient) getHandleError(resp *http.Response) error {
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

// List - Gets all public IP prefixes in a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *PublicIPPrefixesClient) List(resourceGroupName string, options *PublicIPPrefixesListOptions) *PublicIPPrefixesListPager {
	return &PublicIPPrefixesListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp PublicIPPrefixesListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PublicIPPrefixListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *PublicIPPrefixesClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *PublicIPPrefixesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes"
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
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PublicIPPrefixesClient) listHandleResponse(resp *http.Response) (PublicIPPrefixesListResponse, error) {
	result := PublicIPPrefixesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PublicIPPrefixListResult); err != nil {
		return PublicIPPrefixesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *PublicIPPrefixesClient) listHandleError(resp *http.Response) error {
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

// ListAll - Gets all the public IP prefixes in a subscription.
// If the operation fails it returns the *CloudError error type.
func (client *PublicIPPrefixesClient) ListAll(options *PublicIPPrefixesListAllOptions) *PublicIPPrefixesListAllPager {
	return &PublicIPPrefixesListAllPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listAllCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp PublicIPPrefixesListAllResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PublicIPPrefixListResult.NextLink)
		},
	}
}

// listAllCreateRequest creates the ListAll request.
func (client *PublicIPPrefixesClient) listAllCreateRequest(ctx context.Context, options *PublicIPPrefixesListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/publicIPPrefixes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *PublicIPPrefixesClient) listAllHandleResponse(resp *http.Response) (PublicIPPrefixesListAllResponse, error) {
	result := PublicIPPrefixesListAllResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PublicIPPrefixListResult); err != nil {
		return PublicIPPrefixesListAllResponse{}, err
	}
	return result, nil
}

// listAllHandleError handles the ListAll error response.
func (client *PublicIPPrefixesClient) listAllHandleError(resp *http.Response) error {
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

// UpdateTags - Updates public IP prefix tags.
// If the operation fails it returns the *CloudError error type.
func (client *PublicIPPrefixesClient) UpdateTags(ctx context.Context, resourceGroupName string, publicIPPrefixName string, parameters TagsObject, options *PublicIPPrefixesUpdateTagsOptions) (PublicIPPrefixesUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, publicIPPrefixName, parameters, options)
	if err != nil {
		return PublicIPPrefixesUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PublicIPPrefixesUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PublicIPPrefixesUpdateTagsResponse{}, client.updateTagsHandleError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *PublicIPPrefixesClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, publicIPPrefixName string, parameters TagsObject, options *PublicIPPrefixesUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIpPrefixName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publicIPPrefixName == "" {
		return nil, errors.New("parameter publicIPPrefixName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publicIpPrefixName}", url.PathEscape(publicIPPrefixName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *PublicIPPrefixesClient) updateTagsHandleResponse(resp *http.Response) (PublicIPPrefixesUpdateTagsResponse, error) {
	result := PublicIPPrefixesUpdateTagsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PublicIPPrefix); err != nil {
		return PublicIPPrefixesUpdateTagsResponse{}, err
	}
	return result, nil
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *PublicIPPrefixesClient) updateTagsHandleError(resp *http.Response) error {
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
