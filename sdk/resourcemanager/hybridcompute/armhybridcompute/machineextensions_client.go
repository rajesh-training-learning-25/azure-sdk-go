//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridcompute

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

// MachineExtensionsClient contains the methods for the MachineExtensions group.
// Don't use this type directly, use NewMachineExtensionsClient() instead.
type MachineExtensionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewMachineExtensionsClient creates a new instance of MachineExtensionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewMachineExtensionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MachineExtensionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &MachineExtensionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - The operation to create or update the extension.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-31-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - machineName - The name of the machine where the extension should be created or updated.
//   - extensionName - The name of the machine extension.
//   - extensionParameters - Parameters supplied to the Create Machine Extension operation.
//   - options - MachineExtensionsClientBeginCreateOrUpdateOptions contains the optional parameters for the MachineExtensionsClient.BeginCreateOrUpdate
//     method.
func (client *MachineExtensionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, machineName string, extensionName string, extensionParameters MachineExtension, options *MachineExtensionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[MachineExtensionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, machineName, extensionName, extensionParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[MachineExtensionsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[MachineExtensionsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - The operation to create or update the extension.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-31-preview
func (client *MachineExtensionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, machineName string, extensionName string, extensionParameters MachineExtension, options *MachineExtensionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "MachineExtensionsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, machineName, extensionName, extensionParameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *MachineExtensionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, machineName string, extensionName string, extensionParameters MachineExtension, options *MachineExtensionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/extensions/{extensionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	if extensionName == "" {
		return nil, errors.New("parameter extensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionName}", url.PathEscape(extensionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, extensionParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - The operation to delete the extension.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-31-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - machineName - The name of the machine where the extension should be deleted.
//   - extensionName - The name of the machine extension.
//   - options - MachineExtensionsClientBeginDeleteOptions contains the optional parameters for the MachineExtensionsClient.BeginDelete
//     method.
func (client *MachineExtensionsClient) BeginDelete(ctx context.Context, resourceGroupName string, machineName string, extensionName string, options *MachineExtensionsClientBeginDeleteOptions) (*runtime.Poller[MachineExtensionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, machineName, extensionName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[MachineExtensionsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[MachineExtensionsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - The operation to delete the extension.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-31-preview
func (client *MachineExtensionsClient) deleteOperation(ctx context.Context, resourceGroupName string, machineName string, extensionName string, options *MachineExtensionsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "MachineExtensionsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, machineName, extensionName, options)
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
func (client *MachineExtensionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, machineName string, extensionName string, options *MachineExtensionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/extensions/{extensionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	if extensionName == "" {
		return nil, errors.New("parameter extensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionName}", url.PathEscape(extensionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - The operation to get the extension.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-31-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - machineName - The name of the machine containing the extension.
//   - extensionName - The name of the machine extension.
//   - options - MachineExtensionsClientGetOptions contains the optional parameters for the MachineExtensionsClient.Get method.
func (client *MachineExtensionsClient) Get(ctx context.Context, resourceGroupName string, machineName string, extensionName string, options *MachineExtensionsClientGetOptions) (MachineExtensionsClientGetResponse, error) {
	var err error
	const operationName = "MachineExtensionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, machineName, extensionName, options)
	if err != nil {
		return MachineExtensionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MachineExtensionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MachineExtensionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *MachineExtensionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, machineName string, extensionName string, options *MachineExtensionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/extensions/{extensionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	if extensionName == "" {
		return nil, errors.New("parameter extensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionName}", url.PathEscape(extensionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MachineExtensionsClient) getHandleResponse(resp *http.Response) (MachineExtensionsClientGetResponse, error) {
	result := MachineExtensionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MachineExtension); err != nil {
		return MachineExtensionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - The operation to get all extensions of a non-Azure machine
//
// Generated from API version 2024-07-31-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - machineName - The name of the machine containing the extension.
//   - options - MachineExtensionsClientListOptions contains the optional parameters for the MachineExtensionsClient.NewListPager
//     method.
func (client *MachineExtensionsClient) NewListPager(resourceGroupName string, machineName string, options *MachineExtensionsClientListOptions) *runtime.Pager[MachineExtensionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[MachineExtensionsClientListResponse]{
		More: func(page MachineExtensionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MachineExtensionsClientListResponse) (MachineExtensionsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "MachineExtensionsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, machineName, options)
			}, nil)
			if err != nil {
				return MachineExtensionsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *MachineExtensionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, machineName string, options *MachineExtensionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/extensions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2024-07-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MachineExtensionsClient) listHandleResponse(resp *http.Response) (MachineExtensionsClientListResponse, error) {
	result := MachineExtensionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MachineExtensionsListResult); err != nil {
		return MachineExtensionsClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - The operation to create or update the extension.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-31-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - machineName - The name of the machine where the extension should be created or updated.
//   - extensionName - The name of the machine extension.
//   - extensionParameters - Parameters supplied to the Create Machine Extension operation.
//   - options - MachineExtensionsClientBeginUpdateOptions contains the optional parameters for the MachineExtensionsClient.BeginUpdate
//     method.
func (client *MachineExtensionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, machineName string, extensionName string, extensionParameters MachineExtensionUpdate, options *MachineExtensionsClientBeginUpdateOptions) (*runtime.Poller[MachineExtensionsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, machineName, extensionName, extensionParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[MachineExtensionsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[MachineExtensionsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - The operation to create or update the extension.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-31-preview
func (client *MachineExtensionsClient) update(ctx context.Context, resourceGroupName string, machineName string, extensionName string, extensionParameters MachineExtensionUpdate, options *MachineExtensionsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "MachineExtensionsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, machineName, extensionName, extensionParameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *MachineExtensionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, machineName string, extensionName string, extensionParameters MachineExtensionUpdate, options *MachineExtensionsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/extensions/{extensionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	if extensionName == "" {
		return nil, errors.New("parameter extensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionName}", url.PathEscape(extensionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, extensionParameters); err != nil {
		return nil, err
	}
	return req, nil
}
