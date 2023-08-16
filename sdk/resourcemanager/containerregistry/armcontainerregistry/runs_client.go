//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerregistry

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// RunsClient contains the methods for the Runs group.
// Don't use this type directly, use NewRunsClient() instead.
type RunsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRunsClient creates a new instance of RunsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRunsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RunsClient, error) {
	cl, err := arm.NewClient(moduleName+".RunsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RunsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCancel - Cancel an existing run.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - registryName - The name of the container registry.
//   - runID - The run ID.
//   - options - RunsClientBeginCancelOptions contains the optional parameters for the RunsClient.BeginCancel method.
func (client *RunsClient) BeginCancel(ctx context.Context, resourceGroupName string, registryName string, runID string, options *RunsClientBeginCancelOptions) (*runtime.Poller[RunsClientCancelResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.cancel(ctx, resourceGroupName, registryName, runID, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[RunsClientCancelResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[RunsClientCancelResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Cancel - Cancel an existing run.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
func (client *RunsClient) cancel(ctx context.Context, resourceGroupName string, registryName string, runID string, options *RunsClientBeginCancelOptions) (*http.Response, error) {
	var err error
	const operationName = "RunsClient.BeginCancel"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.cancelCreateRequest(ctx, resourceGroupName, registryName, runID, options)
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

// cancelCreateRequest creates the Cancel request.
func (client *RunsClient) cancelCreateRequest(ctx context.Context, resourceGroupName string, registryName string, runID string, options *RunsClientBeginCancelOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/runs/{runId}/cancel"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if runID == "" {
		return nil, errors.New("parameter runID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runId}", url.PathEscape(runID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the detailed information for a given run.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - registryName - The name of the container registry.
//   - runID - The run ID.
//   - options - RunsClientGetOptions contains the optional parameters for the RunsClient.Get method.
func (client *RunsClient) Get(ctx context.Context, resourceGroupName string, registryName string, runID string, options *RunsClientGetOptions) (RunsClientGetResponse, error) {
	var err error
	const operationName = "RunsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, registryName, runID, options)
	if err != nil {
		return RunsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RunsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RunsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *RunsClient) getCreateRequest(ctx context.Context, resourceGroupName string, registryName string, runID string, options *RunsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/runs/{runId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if runID == "" {
		return nil, errors.New("parameter runID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runId}", url.PathEscape(runID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RunsClient) getHandleResponse(resp *http.Response) (RunsClientGetResponse, error) {
	result := RunsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Run); err != nil {
		return RunsClientGetResponse{}, err
	}
	return result, nil
}

// GetLogSasURL - Gets a link to download the run logs.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - registryName - The name of the container registry.
//   - runID - The run ID.
//   - options - RunsClientGetLogSasURLOptions contains the optional parameters for the RunsClient.GetLogSasURL method.
func (client *RunsClient) GetLogSasURL(ctx context.Context, resourceGroupName string, registryName string, runID string, options *RunsClientGetLogSasURLOptions) (RunsClientGetLogSasURLResponse, error) {
	var err error
	const operationName = "RunsClient.GetLogSasURL"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getLogSasURLCreateRequest(ctx, resourceGroupName, registryName, runID, options)
	if err != nil {
		return RunsClientGetLogSasURLResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RunsClientGetLogSasURLResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RunsClientGetLogSasURLResponse{}, err
	}
	resp, err := client.getLogSasURLHandleResponse(httpResp)
	return resp, err
}

// getLogSasURLCreateRequest creates the GetLogSasURL request.
func (client *RunsClient) getLogSasURLCreateRequest(ctx context.Context, resourceGroupName string, registryName string, runID string, options *RunsClientGetLogSasURLOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/runs/{runId}/listLogSasUrl"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if runID == "" {
		return nil, errors.New("parameter runID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runId}", url.PathEscape(runID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getLogSasURLHandleResponse handles the GetLogSasURL response.
func (client *RunsClient) getLogSasURLHandleResponse(resp *http.Response) (RunsClientGetLogSasURLResponse, error) {
	result := RunsClientGetLogSasURLResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RunGetLogResult); err != nil {
		return RunsClientGetLogSasURLResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all the runs for a registry.
//
// Generated from API version 2019-06-01-preview
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - registryName - The name of the container registry.
//   - options - RunsClientListOptions contains the optional parameters for the RunsClient.NewListPager method.
func (client *RunsClient) NewListPager(resourceGroupName string, registryName string, options *RunsClientListOptions) *runtime.Pager[RunsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[RunsClientListResponse]{
		More: func(page RunsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RunsClientListResponse) (RunsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RunsClient.NewListPager")
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, registryName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RunsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return RunsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RunsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *RunsClient) listCreateRequest(ctx context.Context, resourceGroupName string, registryName string, options *RunsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/runs"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RunsClient) listHandleResponse(resp *http.Response) (RunsClientListResponse, error) {
	result := RunsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RunListResult); err != nil {
		return RunsClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Patch the run properties.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - registryName - The name of the container registry.
//   - runID - The run ID.
//   - runUpdateParameters - The run update properties.
//   - options - RunsClientBeginUpdateOptions contains the optional parameters for the RunsClient.BeginUpdate method.
func (client *RunsClient) BeginUpdate(ctx context.Context, resourceGroupName string, registryName string, runID string, runUpdateParameters RunUpdateParameters, options *RunsClientBeginUpdateOptions) (*runtime.Poller[RunsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, registryName, runID, runUpdateParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[RunsClientUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[RunsClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Patch the run properties.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
func (client *RunsClient) update(ctx context.Context, resourceGroupName string, registryName string, runID string, runUpdateParameters RunUpdateParameters, options *RunsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "RunsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, registryName, runID, runUpdateParameters, options)
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

// updateCreateRequest creates the Update request.
func (client *RunsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, registryName string, runID string, runUpdateParameters RunUpdateParameters, options *RunsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/runs/{runId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if runID == "" {
		return nil, errors.New("parameter runID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runId}", url.PathEscape(runID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, runUpdateParameters); err != nil {
		return nil, err
	}
	return req, nil
}
