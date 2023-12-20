//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstreamanalytics

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

// TransformationsClient contains the methods for the Transformations group.
// Don't use this type directly, use NewTransformationsClient() instead.
type TransformationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewTransformationsClient creates a new instance of TransformationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTransformationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TransformationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TransformationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrReplace - Creates a transformation or replaces an already existing transformation under an existing streaming job.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - jobName - The name of the streaming job.
//   - transformationName - The name of the transformation.
//   - transformation - The definition of the transformation that will be used to create a new transformation or replace the existing
//     one under the streaming job.
//   - options - TransformationsClientCreateOrReplaceOptions contains the optional parameters for the TransformationsClient.CreateOrReplace
//     method.
func (client *TransformationsClient) CreateOrReplace(ctx context.Context, resourceGroupName string, jobName string, transformationName string, transformation Transformation, options *TransformationsClientCreateOrReplaceOptions) (TransformationsClientCreateOrReplaceResponse, error) {
	var err error
	const operationName = "TransformationsClient.CreateOrReplace"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrReplaceCreateRequest(ctx, resourceGroupName, jobName, transformationName, transformation, options)
	if err != nil {
		return TransformationsClientCreateOrReplaceResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TransformationsClientCreateOrReplaceResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return TransformationsClientCreateOrReplaceResponse{}, err
	}
	resp, err := client.createOrReplaceHandleResponse(httpResp)
	return resp, err
}

// createOrReplaceCreateRequest creates the CreateOrReplace request.
func (client *TransformationsClient) createOrReplaceCreateRequest(ctx context.Context, resourceGroupName string, jobName string, transformationName string, transformation Transformation, options *TransformationsClientCreateOrReplaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/transformations/{transformationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	if transformationName == "" {
		return nil, errors.New("parameter transformationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{transformationName}", url.PathEscape(transformationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, transformation); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrReplaceHandleResponse handles the CreateOrReplace response.
func (client *TransformationsClient) createOrReplaceHandleResponse(resp *http.Response) (TransformationsClientCreateOrReplaceResponse, error) {
	result := TransformationsClientCreateOrReplaceResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Transformation); err != nil {
		return TransformationsClientCreateOrReplaceResponse{}, err
	}
	return result, nil
}

// Get - Gets details about the specified transformation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - jobName - The name of the streaming job.
//   - transformationName - The name of the transformation.
//   - options - TransformationsClientGetOptions contains the optional parameters for the TransformationsClient.Get method.
func (client *TransformationsClient) Get(ctx context.Context, resourceGroupName string, jobName string, transformationName string, options *TransformationsClientGetOptions) (TransformationsClientGetResponse, error) {
	var err error
	const operationName = "TransformationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, jobName, transformationName, options)
	if err != nil {
		return TransformationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TransformationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TransformationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *TransformationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, jobName string, transformationName string, options *TransformationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/transformations/{transformationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	if transformationName == "" {
		return nil, errors.New("parameter transformationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{transformationName}", url.PathEscape(transformationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TransformationsClient) getHandleResponse(resp *http.Response) (TransformationsClientGetResponse, error) {
	result := TransformationsClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Transformation); err != nil {
		return TransformationsClientGetResponse{}, err
	}
	return result, nil
}

// Update - Updates an existing transformation under an existing streaming job. This can be used to partially update (ie.
// update one or two properties) a transformation without affecting the rest the job or
// transformation definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - jobName - The name of the streaming job.
//   - transformationName - The name of the transformation.
//   - transformation - A Transformation object. The properties specified here will overwrite the corresponding properties in
//     the existing transformation (ie. Those properties will be updated). Any properties that are set to
//     null here will mean that the corresponding property in the existing transformation will remain the same and not change
//     as a result of this PATCH operation.
//   - options - TransformationsClientUpdateOptions contains the optional parameters for the TransformationsClient.Update method.
func (client *TransformationsClient) Update(ctx context.Context, resourceGroupName string, jobName string, transformationName string, transformation Transformation, options *TransformationsClientUpdateOptions) (TransformationsClientUpdateResponse, error) {
	var err error
	const operationName = "TransformationsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, jobName, transformationName, transformation, options)
	if err != nil {
		return TransformationsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TransformationsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TransformationsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *TransformationsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, jobName string, transformationName string, transformation Transformation, options *TransformationsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/transformations/{transformationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	if transformationName == "" {
		return nil, errors.New("parameter transformationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{transformationName}", url.PathEscape(transformationName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, transformation); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *TransformationsClient) updateHandleResponse(resp *http.Response) (TransformationsClientUpdateResponse, error) {
	result := TransformationsClientUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Transformation); err != nil {
		return TransformationsClientUpdateResponse{}, err
	}
	return result, nil
}
