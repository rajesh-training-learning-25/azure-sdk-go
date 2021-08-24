// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// APITagDescriptionClient contains the methods for the APITagDescription group.
// Don't use this type directly, use NewAPITagDescriptionClient() instead.
type APITagDescriptionClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewAPITagDescriptionClient creates a new instance of APITagDescriptionClient with the specified values.
func NewAPITagDescriptionClient(con *armcore.Connection, subscriptionID string) *APITagDescriptionClient {
	return &APITagDescriptionClient{con: con, subscriptionID: subscriptionID}
}

// CreateOrUpdate - Create/Update tag description in scope of the Api.
// If the operation fails it returns the *ErrorResponse error type.
func (client *APITagDescriptionClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, apiID string, tagDescriptionID string, parameters TagDescriptionCreateParameters, options *APITagDescriptionCreateOrUpdateOptions) (APITagDescriptionCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, apiID, tagDescriptionID, parameters, options)
	if err != nil {
		return APITagDescriptionCreateOrUpdateResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return APITagDescriptionCreateOrUpdateResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return APITagDescriptionCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *APITagDescriptionClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, tagDescriptionID string, parameters TagDescriptionCreateParameters, options *APITagDescriptionCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/tagDescriptions/{tagDescriptionId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if tagDescriptionID == "" {
		return nil, errors.New("parameter tagDescriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tagDescriptionId}", url.PathEscape(tagDescriptionID))
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
	reqQP.Set("api-version", "2021-01-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Header.Set("If-Match", *options.IfMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *APITagDescriptionClient) createOrUpdateHandleResponse(resp *azcore.Response) (APITagDescriptionCreateOrUpdateResponse, error) {
	result := APITagDescriptionCreateOrUpdateResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := resp.UnmarshalAsJSON(&result.TagDescriptionContract); err != nil {
		return APITagDescriptionCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *APITagDescriptionClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Delete - Delete tag description for the Api.
// If the operation fails it returns the *ErrorResponse error type.
func (client *APITagDescriptionClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, apiID string, tagDescriptionID string, ifMatch string, options *APITagDescriptionDeleteOptions) (APITagDescriptionDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, apiID, tagDescriptionID, ifMatch, options)
	if err != nil {
		return APITagDescriptionDeleteResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return APITagDescriptionDeleteResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return APITagDescriptionDeleteResponse{}, client.deleteHandleError(resp)
	}
	return APITagDescriptionDeleteResponse{RawResponse: resp.Response}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *APITagDescriptionClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, tagDescriptionID string, ifMatch string, options *APITagDescriptionDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/tagDescriptions/{tagDescriptionId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if tagDescriptionID == "" {
		return nil, errors.New("parameter tagDescriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tagDescriptionId}", url.PathEscape(tagDescriptionID))
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
	reqQP.Set("api-version", "2021-01-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("If-Match", ifMatch)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *APITagDescriptionClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Get Tag description in scope of API
// If the operation fails it returns the *ErrorResponse error type.
func (client *APITagDescriptionClient) Get(ctx context.Context, resourceGroupName string, serviceName string, apiID string, tagDescriptionID string, options *APITagDescriptionGetOptions) (APITagDescriptionGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, apiID, tagDescriptionID, options)
	if err != nil {
		return APITagDescriptionGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return APITagDescriptionGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return APITagDescriptionGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *APITagDescriptionClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, tagDescriptionID string, options *APITagDescriptionGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/tagDescriptions/{tagDescriptionId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if tagDescriptionID == "" {
		return nil, errors.New("parameter tagDescriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tagDescriptionId}", url.PathEscape(tagDescriptionID))
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
	reqQP.Set("api-version", "2021-01-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *APITagDescriptionClient) getHandleResponse(resp *azcore.Response) (APITagDescriptionGetResponse, error) {
	result := APITagDescriptionGetResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := resp.UnmarshalAsJSON(&result.TagDescriptionContract); err != nil {
		return APITagDescriptionGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *APITagDescriptionClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetEntityTag - Gets the entity state version of the tag specified by its identifier.
// If the operation fails it returns the *ErrorResponse error type.
func (client *APITagDescriptionClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, apiID string, tagDescriptionID string, options *APITagDescriptionGetEntityTagOptions) (APITagDescriptionGetEntityTagResponse, error) {
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, apiID, tagDescriptionID, options)
	if err != nil {
		return APITagDescriptionGetEntityTagResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return APITagDescriptionGetEntityTagResponse{}, err
	}
	return client.getEntityTagHandleResponse(resp)
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *APITagDescriptionClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, tagDescriptionID string, options *APITagDescriptionGetEntityTagOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/tagDescriptions/{tagDescriptionId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if tagDescriptionID == "" {
		return nil, errors.New("parameter tagDescriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tagDescriptionId}", url.PathEscape(tagDescriptionID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodHead, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-01-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getEntityTagHandleResponse handles the GetEntityTag response.
func (client *APITagDescriptionClient) getEntityTagHandleResponse(resp *azcore.Response) (APITagDescriptionGetEntityTagResponse, error) {
	result := APITagDescriptionGetEntityTagResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// ListByService - Lists all Tags descriptions in scope of API. Model similar to swagger - tagDescription is defined on API level but tag may be assigned
// to the Operations
// If the operation fails it returns the *ErrorResponse error type.
func (client *APITagDescriptionClient) ListByService(resourceGroupName string, serviceName string, apiID string, options *APITagDescriptionListByServiceOptions) APITagDescriptionListByServicePager {
	return &apiTagDescriptionListByServicePager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, apiID, options)
		},
		advancer: func(ctx context.Context, resp APITagDescriptionListByServiceResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.TagDescriptionCollection.NextLink)
		},
	}
}

// listByServiceCreateRequest creates the ListByService request.
func (client *APITagDescriptionClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, options *APITagDescriptionListByServiceOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/tagDescriptions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
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
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	reqQP.Set("api-version", "2021-01-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *APITagDescriptionClient) listByServiceHandleResponse(resp *azcore.Response) (APITagDescriptionListByServiceResponse, error) {
	result := APITagDescriptionListByServiceResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.TagDescriptionCollection); err != nil {
		return APITagDescriptionListByServiceResponse{}, err
	}
	return result, nil
}

// listByServiceHandleError handles the ListByService error response.
func (client *APITagDescriptionClient) listByServiceHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
