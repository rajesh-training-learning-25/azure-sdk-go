//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armweb

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ResourceHealthMetadataClient contains the methods for the ResourceHealthMetadata group.
// Don't use this type directly, use NewResourceHealthMetadataClient() instead.
type ResourceHealthMetadataClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewResourceHealthMetadataClient creates a new instance of ResourceHealthMetadataClient with the specified values.
func NewResourceHealthMetadataClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ResourceHealthMetadataClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ResourceHealthMetadataClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// GetBySite - Description for Gets the category of ResourceHealthMetadata to use for the given site
// If the operation fails it returns the *DefaultErrorResponse error type.
func (client *ResourceHealthMetadataClient) GetBySite(ctx context.Context, resourceGroupName string, name string, options *ResourceHealthMetadataGetBySiteOptions) (ResourceHealthMetadataGetBySiteResponse, error) {
	req, err := client.getBySiteCreateRequest(ctx, resourceGroupName, name, options)
	if err != nil {
		return ResourceHealthMetadataGetBySiteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ResourceHealthMetadataGetBySiteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ResourceHealthMetadataGetBySiteResponse{}, client.getBySiteHandleError(resp)
	}
	return client.getBySiteHandleResponse(resp)
}

// getBySiteCreateRequest creates the GetBySite request.
func (client *ResourceHealthMetadataClient) getBySiteCreateRequest(ctx context.Context, resourceGroupName string, name string, options *ResourceHealthMetadataGetBySiteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/resourceHealthMetadata/default"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getBySiteHandleResponse handles the GetBySite response.
func (client *ResourceHealthMetadataClient) getBySiteHandleResponse(resp *http.Response) (ResourceHealthMetadataGetBySiteResponse, error) {
	result := ResourceHealthMetadataGetBySiteResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceHealthMetadata); err != nil {
		return ResourceHealthMetadataGetBySiteResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getBySiteHandleError handles the GetBySite error response.
func (client *ResourceHealthMetadataClient) getBySiteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := DefaultErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetBySiteSlot - Description for Gets the category of ResourceHealthMetadata to use for the given site
// If the operation fails it returns the *DefaultErrorResponse error type.
func (client *ResourceHealthMetadataClient) GetBySiteSlot(ctx context.Context, resourceGroupName string, name string, slot string, options *ResourceHealthMetadataGetBySiteSlotOptions) (ResourceHealthMetadataGetBySiteSlotResponse, error) {
	req, err := client.getBySiteSlotCreateRequest(ctx, resourceGroupName, name, slot, options)
	if err != nil {
		return ResourceHealthMetadataGetBySiteSlotResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ResourceHealthMetadataGetBySiteSlotResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ResourceHealthMetadataGetBySiteSlotResponse{}, client.getBySiteSlotHandleError(resp)
	}
	return client.getBySiteSlotHandleResponse(resp)
}

// getBySiteSlotCreateRequest creates the GetBySiteSlot request.
func (client *ResourceHealthMetadataClient) getBySiteSlotCreateRequest(ctx context.Context, resourceGroupName string, name string, slot string, options *ResourceHealthMetadataGetBySiteSlotOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/resourceHealthMetadata/default"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if slot == "" {
		return nil, errors.New("parameter slot cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{slot}", url.PathEscape(slot))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getBySiteSlotHandleResponse handles the GetBySiteSlot response.
func (client *ResourceHealthMetadataClient) getBySiteSlotHandleResponse(resp *http.Response) (ResourceHealthMetadataGetBySiteSlotResponse, error) {
	result := ResourceHealthMetadataGetBySiteSlotResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceHealthMetadata); err != nil {
		return ResourceHealthMetadataGetBySiteSlotResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getBySiteSlotHandleError handles the GetBySiteSlot error response.
func (client *ResourceHealthMetadataClient) getBySiteSlotHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := DefaultErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Description for List all ResourceHealthMetadata for all sites in the subscription.
// If the operation fails it returns the *DefaultErrorResponse error type.
func (client *ResourceHealthMetadataClient) List(options *ResourceHealthMetadataListOptions) *ResourceHealthMetadataListPager {
	return &ResourceHealthMetadataListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ResourceHealthMetadataListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ResourceHealthMetadataCollection.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ResourceHealthMetadataClient) listCreateRequest(ctx context.Context, options *ResourceHealthMetadataListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Web/resourceHealthMetadata"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ResourceHealthMetadataClient) listHandleResponse(resp *http.Response) (ResourceHealthMetadataListResponse, error) {
	result := ResourceHealthMetadataListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceHealthMetadataCollection); err != nil {
		return ResourceHealthMetadataListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ResourceHealthMetadataClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := DefaultErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByResourceGroup - Description for List all ResourceHealthMetadata for all sites in the resource group in the subscription.
// If the operation fails it returns the *DefaultErrorResponse error type.
func (client *ResourceHealthMetadataClient) ListByResourceGroup(resourceGroupName string, options *ResourceHealthMetadataListByResourceGroupOptions) *ResourceHealthMetadataListByResourceGroupPager {
	return &ResourceHealthMetadataListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp ResourceHealthMetadataListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ResourceHealthMetadataCollection.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ResourceHealthMetadataClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ResourceHealthMetadataListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/resourceHealthMetadata"
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
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ResourceHealthMetadataClient) listByResourceGroupHandleResponse(resp *http.Response) (ResourceHealthMetadataListByResourceGroupResponse, error) {
	result := ResourceHealthMetadataListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceHealthMetadataCollection); err != nil {
		return ResourceHealthMetadataListByResourceGroupResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *ResourceHealthMetadataClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := DefaultErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListBySite - Description for Gets the category of ResourceHealthMetadata to use for the given site as a collection
// If the operation fails it returns the *DefaultErrorResponse error type.
func (client *ResourceHealthMetadataClient) ListBySite(resourceGroupName string, name string, options *ResourceHealthMetadataListBySiteOptions) *ResourceHealthMetadataListBySitePager {
	return &ResourceHealthMetadataListBySitePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySiteCreateRequest(ctx, resourceGroupName, name, options)
		},
		advancer: func(ctx context.Context, resp ResourceHealthMetadataListBySiteResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ResourceHealthMetadataCollection.NextLink)
		},
	}
}

// listBySiteCreateRequest creates the ListBySite request.
func (client *ResourceHealthMetadataClient) listBySiteCreateRequest(ctx context.Context, resourceGroupName string, name string, options *ResourceHealthMetadataListBySiteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/resourceHealthMetadata"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySiteHandleResponse handles the ListBySite response.
func (client *ResourceHealthMetadataClient) listBySiteHandleResponse(resp *http.Response) (ResourceHealthMetadataListBySiteResponse, error) {
	result := ResourceHealthMetadataListBySiteResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceHealthMetadataCollection); err != nil {
		return ResourceHealthMetadataListBySiteResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listBySiteHandleError handles the ListBySite error response.
func (client *ResourceHealthMetadataClient) listBySiteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := DefaultErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListBySiteSlot - Description for Gets the category of ResourceHealthMetadata to use for the given site as a collection
// If the operation fails it returns the *DefaultErrorResponse error type.
func (client *ResourceHealthMetadataClient) ListBySiteSlot(resourceGroupName string, name string, slot string, options *ResourceHealthMetadataListBySiteSlotOptions) *ResourceHealthMetadataListBySiteSlotPager {
	return &ResourceHealthMetadataListBySiteSlotPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySiteSlotCreateRequest(ctx, resourceGroupName, name, slot, options)
		},
		advancer: func(ctx context.Context, resp ResourceHealthMetadataListBySiteSlotResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ResourceHealthMetadataCollection.NextLink)
		},
	}
}

// listBySiteSlotCreateRequest creates the ListBySiteSlot request.
func (client *ResourceHealthMetadataClient) listBySiteSlotCreateRequest(ctx context.Context, resourceGroupName string, name string, slot string, options *ResourceHealthMetadataListBySiteSlotOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/resourceHealthMetadata"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if slot == "" {
		return nil, errors.New("parameter slot cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{slot}", url.PathEscape(slot))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySiteSlotHandleResponse handles the ListBySiteSlot response.
func (client *ResourceHealthMetadataClient) listBySiteSlotHandleResponse(resp *http.Response) (ResourceHealthMetadataListBySiteSlotResponse, error) {
	result := ResourceHealthMetadataListBySiteSlotResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceHealthMetadataCollection); err != nil {
		return ResourceHealthMetadataListBySiteSlotResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listBySiteSlotHandleError handles the ListBySiteSlot error response.
func (client *ResourceHealthMetadataClient) listBySiteSlotHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := DefaultErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
