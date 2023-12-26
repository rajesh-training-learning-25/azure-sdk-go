//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

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

// IPGroupsClient contains the methods for the IPGroups group.
// Don't use this type directly, use NewIPGroupsClient() instead.
type IPGroupsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewIPGroupsClient creates a new instance of IPGroupsClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewIPGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IPGroupsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &IPGroupsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates an ipGroups in a specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The name of the resource group.
//   - ipGroupsName - The name of the ipGroups.
//   - parameters - Parameters supplied to the create or update IpGroups operation.
//   - options - IPGroupsClientBeginCreateOrUpdateOptions contains the optional parameters for the IPGroupsClient.BeginCreateOrUpdate
//     method.
func (client *IPGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, ipGroupsName string, parameters IPGroup, options *IPGroupsClientBeginCreateOrUpdateOptions) (*runtime.Poller[IPGroupsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, ipGroupsName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[IPGroupsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[IPGroupsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates an ipGroups in a specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
func (client *IPGroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, ipGroupsName string, parameters IPGroup, options *IPGroupsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "IPGroupsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, ipGroupsName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *IPGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, ipGroupsName string, parameters IPGroup, options *IPGroupsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ipGroups/{ipGroupsName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ipGroupsName == "" {
		return nil, errors.New("parameter ipGroupsName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipGroupsName}", url.PathEscape(ipGroupsName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the specified ipGroups.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The name of the resource group.
//   - ipGroupsName - The name of the ipGroups.
//   - options - IPGroupsClientBeginDeleteOptions contains the optional parameters for the IPGroupsClient.BeginDelete method.
func (client *IPGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, ipGroupsName string, options *IPGroupsClientBeginDeleteOptions) (*runtime.Poller[IPGroupsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, ipGroupsName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[IPGroupsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[IPGroupsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the specified ipGroups.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
func (client *IPGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, ipGroupsName string, options *IPGroupsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "IPGroupsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, ipGroupsName, options)
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
func (client *IPGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, ipGroupsName string, options *IPGroupsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ipGroups/{ipGroupsName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ipGroupsName == "" {
		return nil, errors.New("parameter ipGroupsName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipGroupsName}", url.PathEscape(ipGroupsName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the specified ipGroups.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The name of the resource group.
//   - ipGroupsName - The name of the ipGroups.
//   - options - IPGroupsClientGetOptions contains the optional parameters for the IPGroupsClient.Get method.
func (client *IPGroupsClient) Get(ctx context.Context, resourceGroupName string, ipGroupsName string, options *IPGroupsClientGetOptions) (IPGroupsClientGetResponse, error) {
	var err error
	const operationName = "IPGroupsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, ipGroupsName, options)
	if err != nil {
		return IPGroupsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IPGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IPGroupsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *IPGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, ipGroupsName string, options *IPGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ipGroups/{ipGroupsName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ipGroupsName == "" {
		return nil, errors.New("parameter ipGroupsName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipGroupsName}", url.PathEscape(ipGroupsName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IPGroupsClient) getHandleResponse(resp *http.Response) (IPGroupsClientGetResponse, error) {
	result := IPGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IPGroup); err != nil {
		return IPGroupsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all IpGroups in a subscription.
//
// Generated from API version 2023-06-01
//   - options - IPGroupsClientListOptions contains the optional parameters for the IPGroupsClient.NewListPager method.
func (client *IPGroupsClient) NewListPager(options *IPGroupsClientListOptions) *runtime.Pager[IPGroupsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[IPGroupsClientListResponse]{
		More: func(page IPGroupsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *IPGroupsClientListResponse) (IPGroupsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "IPGroupsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return IPGroupsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *IPGroupsClient) listCreateRequest(ctx context.Context, options *IPGroupsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/ipGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IPGroupsClient) listHandleResponse(resp *http.Response) (IPGroupsClientListResponse, error) {
	result := IPGroupsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IPGroupListResult); err != nil {
		return IPGroupsClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets all IpGroups in a resource group.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The name of the resource group.
//   - options - IPGroupsClientListByResourceGroupOptions contains the optional parameters for the IPGroupsClient.NewListByResourceGroupPager
//     method.
func (client *IPGroupsClient) NewListByResourceGroupPager(resourceGroupName string, options *IPGroupsClientListByResourceGroupOptions) *runtime.Pager[IPGroupsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[IPGroupsClientListByResourceGroupResponse]{
		More: func(page IPGroupsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *IPGroupsClientListByResourceGroupResponse) (IPGroupsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "IPGroupsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return IPGroupsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *IPGroupsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *IPGroupsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ipGroups"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *IPGroupsClient) listByResourceGroupHandleResponse(resp *http.Response) (IPGroupsClientListByResourceGroupResponse, error) {
	result := IPGroupsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IPGroupListResult); err != nil {
		return IPGroupsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// UpdateGroups - Updates tags of an IpGroups resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The name of the resource group.
//   - ipGroupsName - The name of the ipGroups.
//   - parameters - Parameters supplied to the update ipGroups operation.
//   - options - IPGroupsClientUpdateGroupsOptions contains the optional parameters for the IPGroupsClient.UpdateGroups method.
func (client *IPGroupsClient) UpdateGroups(ctx context.Context, resourceGroupName string, ipGroupsName string, parameters TagsObject, options *IPGroupsClientUpdateGroupsOptions) (IPGroupsClientUpdateGroupsResponse, error) {
	var err error
	const operationName = "IPGroupsClient.UpdateGroups"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateGroupsCreateRequest(ctx, resourceGroupName, ipGroupsName, parameters, options)
	if err != nil {
		return IPGroupsClientUpdateGroupsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IPGroupsClientUpdateGroupsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IPGroupsClientUpdateGroupsResponse{}, err
	}
	resp, err := client.updateGroupsHandleResponse(httpResp)
	return resp, err
}

// updateGroupsCreateRequest creates the UpdateGroups request.
func (client *IPGroupsClient) updateGroupsCreateRequest(ctx context.Context, resourceGroupName string, ipGroupsName string, parameters TagsObject, options *IPGroupsClientUpdateGroupsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ipGroups/{ipGroupsName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ipGroupsName == "" {
		return nil, errors.New("parameter ipGroupsName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipGroupsName}", url.PathEscape(ipGroupsName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateGroupsHandleResponse handles the UpdateGroups response.
func (client *IPGroupsClient) updateGroupsHandleResponse(resp *http.Response) (IPGroupsClientUpdateGroupsResponse, error) {
	result := IPGroupsClientUpdateGroupsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IPGroup); err != nil {
		return IPGroupsClientUpdateGroupsResponse{}, err
	}
	return result, nil
}
