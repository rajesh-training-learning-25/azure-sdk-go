// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// CustomIPPrefixesClient contains the methods for the CustomIPPrefixes group.
// Don't use this type directly, use NewCustomIPPrefixesClient() instead.
type CustomIPPrefixesClient struct {
	con *armcore.Connection
	subscriptionID string
}

// NewCustomIPPrefixesClient creates a new instance of CustomIPPrefixesClient with the specified values.
func NewCustomIPPrefixesClient(con *armcore.Connection, subscriptionID string) *CustomIPPrefixesClient {
	return &CustomIPPrefixesClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a custom IP prefix.
func (client *CustomIPPrefixesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, customIPPrefixName string, parameters CustomIPPrefix, options *CustomIPPrefixesBeginCreateOrUpdateOptions) (CustomIPPrefixPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, customIPPrefixName, parameters, options)
	if err != nil {
		return CustomIPPrefixPollerResponse{}, err
	}
	result := CustomIPPrefixPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("CustomIPPrefixesClient.CreateOrUpdate", "location", resp, client.createOrUpdateHandleError)
	if err != nil {
		return CustomIPPrefixPollerResponse{}, err
	}
	poller := &customIPPrefixPoller{
		pt: pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (CustomIPPrefixResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new CustomIPPrefixPoller from the specified resume token.
// token - The value must come from a previous call to CustomIPPrefixPoller.ResumeToken().
func (client *CustomIPPrefixesClient) ResumeCreateOrUpdate(token string) (CustomIPPrefixPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("CustomIPPrefixesClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &customIPPrefixPoller{
		pipeline: client.con.Pipeline(),
		pt: pt,
	}, nil
}

// CreateOrUpdate - Creates or updates a custom IP prefix.
func (client *CustomIPPrefixesClient) createOrUpdate(ctx context.Context, resourceGroupName string, customIPPrefixName string, parameters CustomIPPrefix, options *CustomIPPrefixesBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, customIPPrefixName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	 return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *CustomIPPrefixesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, customIPPrefixName string, parameters CustomIPPrefix, options *CustomIPPrefixesBeginCreateOrUpdateOptions) (*azcore.Request, error) {
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
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-07-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *CustomIPPrefixesClient) createOrUpdateHandleResponse(resp *azcore.Response) (CustomIPPrefixResponse, error) {
	var val *CustomIPPrefix
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return CustomIPPrefixResponse{}, err
	}
return CustomIPPrefixResponse{RawResponse: resp.Response, CustomIPPrefix: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *CustomIPPrefixesClient) createOrUpdateHandleError(resp *azcore.Response) error {
var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Deletes the specified custom IP prefix.
func (client *CustomIPPrefixesClient) BeginDelete(ctx context.Context, resourceGroupName string, customIPPrefixName string, options *CustomIPPrefixesBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, customIPPrefixName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("CustomIPPrefixesClient.Delete", "location", resp, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *CustomIPPrefixesClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("CustomIPPrefixesClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt: pt,
	}, nil
}

// Delete - Deletes the specified custom IP prefix.
func (client *CustomIPPrefixesClient) deleteOperation(ctx context.Context, resourceGroupName string, customIPPrefixName string, options *CustomIPPrefixesBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, customIPPrefixName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	 return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CustomIPPrefixesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, customIPPrefixName string, options *CustomIPPrefixesBeginDeleteOptions) (*azcore.Request, error) {
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
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-07-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *CustomIPPrefixesClient) deleteHandleError(resp *azcore.Response) error {
var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified custom IP prefix in a specified resource group.
func (client *CustomIPPrefixesClient) Get(ctx context.Context, resourceGroupName string, customIPPrefixName string, options *CustomIPPrefixesGetOptions) (CustomIPPrefixResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, customIPPrefixName, options)
	if err != nil {
		return CustomIPPrefixResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return CustomIPPrefixResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return CustomIPPrefixResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CustomIPPrefixesClient) getCreateRequest(ctx context.Context, resourceGroupName string, customIPPrefixName string, options *CustomIPPrefixesGetOptions) (*azcore.Request, error) {
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
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-07-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CustomIPPrefixesClient) getHandleResponse(resp *azcore.Response) (CustomIPPrefixResponse, error) {
	var val *CustomIPPrefix
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return CustomIPPrefixResponse{}, err
	}
return CustomIPPrefixResponse{RawResponse: resp.Response, CustomIPPrefix: val}, nil
}

// getHandleError handles the Get error response.
func (client *CustomIPPrefixesClient) getHandleError(resp *azcore.Response) error {
var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets all custom IP prefixes in a resource group.
func (client *CustomIPPrefixesClient) List(resourceGroupName string, options *CustomIPPrefixesListOptions) (CustomIPPrefixListResultPager) {
	return &customIPPrefixListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp CustomIPPrefixListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.CustomIPPrefixListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *CustomIPPrefixesClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *CustomIPPrefixesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/customIpPrefixes"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-07-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *CustomIPPrefixesClient) listHandleResponse(resp *azcore.Response) (CustomIPPrefixListResultResponse, error) {
	var val *CustomIPPrefixListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return CustomIPPrefixListResultResponse{}, err
	}
return CustomIPPrefixListResultResponse{RawResponse: resp.Response, CustomIPPrefixListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *CustomIPPrefixesClient) listHandleError(resp *azcore.Response) error {
var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListAll - Gets all the custom IP prefixes in a subscription.
func (client *CustomIPPrefixesClient) ListAll(options *CustomIPPrefixesListAllOptions) (CustomIPPrefixListResultPager) {
	return &customIPPrefixListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listAllCreateRequest(ctx, options)
		},
		responder: client.listAllHandleResponse,
		errorer:   client.listAllHandleError,
		advancer: func(ctx context.Context, resp CustomIPPrefixListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.CustomIPPrefixListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listAllCreateRequest creates the ListAll request.
func (client *CustomIPPrefixesClient) listAllCreateRequest(ctx context.Context, options *CustomIPPrefixesListAllOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/customIpPrefixes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-07-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *CustomIPPrefixesClient) listAllHandleResponse(resp *azcore.Response) (CustomIPPrefixListResultResponse, error) {
	var val *CustomIPPrefixListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return CustomIPPrefixListResultResponse{}, err
	}
return CustomIPPrefixListResultResponse{RawResponse: resp.Response, CustomIPPrefixListResult: val}, nil
}

// listAllHandleError handles the ListAll error response.
func (client *CustomIPPrefixesClient) listAllHandleError(resp *azcore.Response) error {
var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// UpdateTags - Updates custom IP prefix tags.
func (client *CustomIPPrefixesClient) UpdateTags(ctx context.Context, resourceGroupName string, customIPPrefixName string, parameters TagsObject, options *CustomIPPrefixesUpdateTagsOptions) (CustomIPPrefixResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, customIPPrefixName, parameters, options)
	if err != nil {
		return CustomIPPrefixResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return CustomIPPrefixResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return CustomIPPrefixResponse{}, client.updateTagsHandleError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *CustomIPPrefixesClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, customIPPrefixName string, parameters TagsObject, options *CustomIPPrefixesUpdateTagsOptions) (*azcore.Request, error) {
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
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-07-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *CustomIPPrefixesClient) updateTagsHandleResponse(resp *azcore.Response) (CustomIPPrefixResponse, error) {
	var val *CustomIPPrefix
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return CustomIPPrefixResponse{}, err
	}
return CustomIPPrefixResponse{RawResponse: resp.Response, CustomIPPrefix: val}, nil
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *CustomIPPrefixesClient) updateTagsHandleError(resp *azcore.Response) error {
var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

