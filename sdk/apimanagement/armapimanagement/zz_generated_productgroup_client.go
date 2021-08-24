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

// ProductGroupClient contains the methods for the ProductGroup group.
// Don't use this type directly, use NewProductGroupClient() instead.
type ProductGroupClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewProductGroupClient creates a new instance of ProductGroupClient with the specified values.
func NewProductGroupClient(con *armcore.Connection, subscriptionID string) *ProductGroupClient {
	return &ProductGroupClient{con: con, subscriptionID: subscriptionID}
}

// CheckEntityExists - Checks that Group entity specified by identifier is associated with the Product entity.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ProductGroupClient) CheckEntityExists(ctx context.Context, resourceGroupName string, serviceName string, productID string, groupID string, options *ProductGroupCheckEntityExistsOptions) (ProductGroupCheckEntityExistsResponse, error) {
	req, err := client.checkEntityExistsCreateRequest(ctx, resourceGroupName, serviceName, productID, groupID, options)
	if err != nil {
		return ProductGroupCheckEntityExistsResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ProductGroupCheckEntityExistsResponse{}, err
	}
	result := ProductGroupCheckEntityExistsResponse{RawResponse: resp.Response}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// checkEntityExistsCreateRequest creates the CheckEntityExists request.
func (client *ProductGroupClient) checkEntityExistsCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, productID string, groupID string, options *ProductGroupCheckEntityExistsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/groups/{groupId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if productID == "" {
		return nil, errors.New("parameter productID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productId}", url.PathEscape(productID))
	if groupID == "" {
		return nil, errors.New("parameter groupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupId}", url.PathEscape(groupID))
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

// CreateOrUpdate - Adds the association between the specified developer group with the specified product.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ProductGroupClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, productID string, groupID string, options *ProductGroupCreateOrUpdateOptions) (ProductGroupCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, productID, groupID, options)
	if err != nil {
		return ProductGroupCreateOrUpdateResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ProductGroupCreateOrUpdateResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return ProductGroupCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ProductGroupClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, productID string, groupID string, options *ProductGroupCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/groups/{groupId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if productID == "" {
		return nil, errors.New("parameter productID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productId}", url.PathEscape(productID))
	if groupID == "" {
		return nil, errors.New("parameter groupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupId}", url.PathEscape(groupID))
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
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ProductGroupClient) createOrUpdateHandleResponse(resp *azcore.Response) (ProductGroupCreateOrUpdateResponse, error) {
	result := ProductGroupCreateOrUpdateResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.GroupContract); err != nil {
		return ProductGroupCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ProductGroupClient) createOrUpdateHandleError(resp *azcore.Response) error {
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

// Delete - Deletes the association between the specified group and product.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ProductGroupClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, productID string, groupID string, options *ProductGroupDeleteOptions) (ProductGroupDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, productID, groupID, options)
	if err != nil {
		return ProductGroupDeleteResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ProductGroupDeleteResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return ProductGroupDeleteResponse{}, client.deleteHandleError(resp)
	}
	return ProductGroupDeleteResponse{RawResponse: resp.Response}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ProductGroupClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, productID string, groupID string, options *ProductGroupDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/groups/{groupId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if productID == "" {
		return nil, errors.New("parameter productID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productId}", url.PathEscape(productID))
	if groupID == "" {
		return nil, errors.New("parameter groupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupId}", url.PathEscape(groupID))
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
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ProductGroupClient) deleteHandleError(resp *azcore.Response) error {
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

// ListByProduct - Lists the collection of developer groups associated with the specified product.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ProductGroupClient) ListByProduct(resourceGroupName string, serviceName string, productID string, options *ProductGroupListByProductOptions) ProductGroupListByProductPager {
	return &productGroupListByProductPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByProductCreateRequest(ctx, resourceGroupName, serviceName, productID, options)
		},
		advancer: func(ctx context.Context, resp ProductGroupListByProductResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.GroupCollection.NextLink)
		},
	}
}

// listByProductCreateRequest creates the ListByProduct request.
func (client *ProductGroupClient) listByProductCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, productID string, options *ProductGroupListByProductOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/groups"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if productID == "" {
		return nil, errors.New("parameter productID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productId}", url.PathEscape(productID))
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

// listByProductHandleResponse handles the ListByProduct response.
func (client *ProductGroupClient) listByProductHandleResponse(resp *azcore.Response) (ProductGroupListByProductResponse, error) {
	result := ProductGroupListByProductResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.GroupCollection); err != nil {
		return ProductGroupListByProductResponse{}, err
	}
	return result, nil
}

// listByProductHandleError handles the ListByProduct error response.
func (client *ProductGroupClient) listByProductHandleError(resp *azcore.Response) error {
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
