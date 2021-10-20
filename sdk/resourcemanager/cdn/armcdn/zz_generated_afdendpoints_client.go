//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcdn

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

// AFDEndpointsClient contains the methods for the AFDEndpoints group.
// Don't use this type directly, use NewAFDEndpointsClient() instead.
type AFDEndpointsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewAFDEndpointsClient creates a new instance of AFDEndpointsClient with the specified values.
func NewAFDEndpointsClient(con *arm.Connection, subscriptionID string) *AFDEndpointsClient {
	return &AFDEndpointsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreate - Creates a new AzureFrontDoor endpoint with the specified endpoint name under the specified subscription, resource group and profile.
// If the operation fails it returns the *AfdErrorResponse error type.
func (client *AFDEndpointsClient) BeginCreate(ctx context.Context, resourceGroupName string, profileName string, endpointName string, endpoint AFDEndpoint, options *AFDEndpointsBeginCreateOptions) (AFDEndpointsCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, profileName, endpointName, endpoint, options)
	if err != nil {
		return AFDEndpointsCreatePollerResponse{}, err
	}
	result := AFDEndpointsCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("AFDEndpointsClient.Create", "azure-async-operation", resp, client.pl, client.createHandleError)
	if err != nil {
		return AFDEndpointsCreatePollerResponse{}, err
	}
	result.Poller = &AFDEndpointsCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Creates a new AzureFrontDoor endpoint with the specified endpoint name under the specified subscription, resource group and profile.
// If the operation fails it returns the *AfdErrorResponse error type.
func (client *AFDEndpointsClient) create(ctx context.Context, resourceGroupName string, profileName string, endpointName string, endpoint AFDEndpoint, options *AFDEndpointsBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, profileName, endpointName, endpoint, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createHandleError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *AFDEndpointsClient) createCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, endpoint AFDEndpoint, options *AFDEndpointsBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/afdEndpoints/{endpointName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, endpoint)
}

// createHandleError handles the Create error response.
func (client *AFDEndpointsClient) createHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := AfdErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes an existing AzureFrontDoor endpoint with the specified endpoint name under the specified subscription, resource group and profile.
// If the operation fails it returns the *AfdErrorResponse error type.
func (client *AFDEndpointsClient) BeginDelete(ctx context.Context, resourceGroupName string, profileName string, endpointName string, options *AFDEndpointsBeginDeleteOptions) (AFDEndpointsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, profileName, endpointName, options)
	if err != nil {
		return AFDEndpointsDeletePollerResponse{}, err
	}
	result := AFDEndpointsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("AFDEndpointsClient.Delete", "azure-async-operation", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return AFDEndpointsDeletePollerResponse{}, err
	}
	result.Poller = &AFDEndpointsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes an existing AzureFrontDoor endpoint with the specified endpoint name under the specified subscription, resource group and profile.
// If the operation fails it returns the *AfdErrorResponse error type.
func (client *AFDEndpointsClient) deleteOperation(ctx context.Context, resourceGroupName string, profileName string, endpointName string, options *AFDEndpointsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, profileName, endpointName, options)
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
func (client *AFDEndpointsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, options *AFDEndpointsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/afdEndpoints/{endpointName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *AFDEndpointsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := AfdErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets an existing AzureFrontDoor endpoint with the specified endpoint name under the specified subscription, resource group and profile.
// If the operation fails it returns the *AfdErrorResponse error type.
func (client *AFDEndpointsClient) Get(ctx context.Context, resourceGroupName string, profileName string, endpointName string, options *AFDEndpointsGetOptions) (AFDEndpointsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, profileName, endpointName, options)
	if err != nil {
		return AFDEndpointsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AFDEndpointsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AFDEndpointsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AFDEndpointsClient) getCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, options *AFDEndpointsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/afdEndpoints/{endpointName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AFDEndpointsClient) getHandleResponse(resp *http.Response) (AFDEndpointsGetResponse, error) {
	result := AFDEndpointsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AFDEndpoint); err != nil {
		return AFDEndpointsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *AFDEndpointsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := AfdErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByProfile - Lists existing AzureFrontDoor endpoints.
// If the operation fails it returns the *AfdErrorResponse error type.
func (client *AFDEndpointsClient) ListByProfile(resourceGroupName string, profileName string, options *AFDEndpointsListByProfileOptions) *AFDEndpointsListByProfilePager {
	return &AFDEndpointsListByProfilePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByProfileCreateRequest(ctx, resourceGroupName, profileName, options)
		},
		advancer: func(ctx context.Context, resp AFDEndpointsListByProfileResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AFDEndpointListResult.NextLink)
		},
	}
}

// listByProfileCreateRequest creates the ListByProfile request.
func (client *AFDEndpointsClient) listByProfileCreateRequest(ctx context.Context, resourceGroupName string, profileName string, options *AFDEndpointsListByProfileOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/afdEndpoints"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByProfileHandleResponse handles the ListByProfile response.
func (client *AFDEndpointsClient) listByProfileHandleResponse(resp *http.Response) (AFDEndpointsListByProfileResponse, error) {
	result := AFDEndpointsListByProfileResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AFDEndpointListResult); err != nil {
		return AFDEndpointsListByProfileResponse{}, err
	}
	return result, nil
}

// listByProfileHandleError handles the ListByProfile error response.
func (client *AFDEndpointsClient) listByProfileHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := AfdErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListResourceUsage - Checks the quota and actual usage of endpoints under the given CDN profile.
// If the operation fails it returns the *AfdErrorResponse error type.
func (client *AFDEndpointsClient) ListResourceUsage(resourceGroupName string, profileName string, endpointName string, options *AFDEndpointsListResourceUsageOptions) *AFDEndpointsListResourceUsagePager {
	return &AFDEndpointsListResourceUsagePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listResourceUsageCreateRequest(ctx, resourceGroupName, profileName, endpointName, options)
		},
		advancer: func(ctx context.Context, resp AFDEndpointsListResourceUsageResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.UsagesListResult.NextLink)
		},
	}
}

// listResourceUsageCreateRequest creates the ListResourceUsage request.
func (client *AFDEndpointsClient) listResourceUsageCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, options *AFDEndpointsListResourceUsageOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/afdEndpoints/{endpointName}/usages"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listResourceUsageHandleResponse handles the ListResourceUsage response.
func (client *AFDEndpointsClient) listResourceUsageHandleResponse(resp *http.Response) (AFDEndpointsListResourceUsageResponse, error) {
	result := AFDEndpointsListResourceUsageResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.UsagesListResult); err != nil {
		return AFDEndpointsListResourceUsageResponse{}, err
	}
	return result, nil
}

// listResourceUsageHandleError handles the ListResourceUsage error response.
func (client *AFDEndpointsClient) listResourceUsageHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := AfdErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginPurgeContent - Removes a content from AzureFrontDoor.
// If the operation fails it returns the *AfdErrorResponse error type.
func (client *AFDEndpointsClient) BeginPurgeContent(ctx context.Context, resourceGroupName string, profileName string, endpointName string, contents AfdPurgeParameters, options *AFDEndpointsBeginPurgeContentOptions) (AFDEndpointsPurgeContentPollerResponse, error) {
	resp, err := client.purgeContent(ctx, resourceGroupName, profileName, endpointName, contents, options)
	if err != nil {
		return AFDEndpointsPurgeContentPollerResponse{}, err
	}
	result := AFDEndpointsPurgeContentPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("AFDEndpointsClient.PurgeContent", "azure-async-operation", resp, client.pl, client.purgeContentHandleError)
	if err != nil {
		return AFDEndpointsPurgeContentPollerResponse{}, err
	}
	result.Poller = &AFDEndpointsPurgeContentPoller{
		pt: pt,
	}
	return result, nil
}

// PurgeContent - Removes a content from AzureFrontDoor.
// If the operation fails it returns the *AfdErrorResponse error type.
func (client *AFDEndpointsClient) purgeContent(ctx context.Context, resourceGroupName string, profileName string, endpointName string, contents AfdPurgeParameters, options *AFDEndpointsBeginPurgeContentOptions) (*http.Response, error) {
	req, err := client.purgeContentCreateRequest(ctx, resourceGroupName, profileName, endpointName, contents, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.purgeContentHandleError(resp)
	}
	return resp, nil
}

// purgeContentCreateRequest creates the PurgeContent request.
func (client *AFDEndpointsClient) purgeContentCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, contents AfdPurgeParameters, options *AFDEndpointsBeginPurgeContentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/afdEndpoints/{endpointName}/purge"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, contents)
}

// purgeContentHandleError handles the PurgeContent error response.
func (client *AFDEndpointsClient) purgeContentHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := AfdErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginUpdate - Updates an existing AzureFrontDoor endpoint with the specified endpoint name under the specified subscription, resource group and profile.
// Only tags can be updated after creating an endpoint. To
// update origins, use the Update Origin operation. To update origin groups, use the Update Origin group operation. To update domains, use the Update Custom
// Domain operation.
// If the operation fails it returns the *AfdErrorResponse error type.
func (client *AFDEndpointsClient) BeginUpdate(ctx context.Context, resourceGroupName string, profileName string, endpointName string, endpointUpdateProperties AFDEndpointUpdateParameters, options *AFDEndpointsBeginUpdateOptions) (AFDEndpointsUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, profileName, endpointName, endpointUpdateProperties, options)
	if err != nil {
		return AFDEndpointsUpdatePollerResponse{}, err
	}
	result := AFDEndpointsUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("AFDEndpointsClient.Update", "azure-async-operation", resp, client.pl, client.updateHandleError)
	if err != nil {
		return AFDEndpointsUpdatePollerResponse{}, err
	}
	result.Poller = &AFDEndpointsUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates an existing AzureFrontDoor endpoint with the specified endpoint name under the specified subscription, resource group and profile. Only
// tags can be updated after creating an endpoint. To
// update origins, use the Update Origin operation. To update origin groups, use the Update Origin group operation. To update domains, use the Update Custom
// Domain operation.
// If the operation fails it returns the *AfdErrorResponse error type.
func (client *AFDEndpointsClient) update(ctx context.Context, resourceGroupName string, profileName string, endpointName string, endpointUpdateProperties AFDEndpointUpdateParameters, options *AFDEndpointsBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, profileName, endpointName, endpointUpdateProperties, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *AFDEndpointsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, endpointUpdateProperties AFDEndpointUpdateParameters, options *AFDEndpointsBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/afdEndpoints/{endpointName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, endpointUpdateProperties)
}

// updateHandleError handles the Update error response.
func (client *AFDEndpointsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := AfdErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ValidateCustomDomain - Validates the custom domain mapping to ensure it maps to the correct CDN endpoint in DNS.
// If the operation fails it returns the *AfdErrorResponse error type.
func (client *AFDEndpointsClient) ValidateCustomDomain(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainProperties ValidateCustomDomainInput, options *AFDEndpointsValidateCustomDomainOptions) (AFDEndpointsValidateCustomDomainResponse, error) {
	req, err := client.validateCustomDomainCreateRequest(ctx, resourceGroupName, profileName, endpointName, customDomainProperties, options)
	if err != nil {
		return AFDEndpointsValidateCustomDomainResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AFDEndpointsValidateCustomDomainResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AFDEndpointsValidateCustomDomainResponse{}, client.validateCustomDomainHandleError(resp)
	}
	return client.validateCustomDomainHandleResponse(resp)
}

// validateCustomDomainCreateRequest creates the ValidateCustomDomain request.
func (client *AFDEndpointsClient) validateCustomDomainCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainProperties ValidateCustomDomainInput, options *AFDEndpointsValidateCustomDomainOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/afdEndpoints/{endpointName}/validateCustomDomain"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, customDomainProperties)
}

// validateCustomDomainHandleResponse handles the ValidateCustomDomain response.
func (client *AFDEndpointsClient) validateCustomDomainHandleResponse(resp *http.Response) (AFDEndpointsValidateCustomDomainResponse, error) {
	result := AFDEndpointsValidateCustomDomainResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ValidateCustomDomainOutput); err != nil {
		return AFDEndpointsValidateCustomDomainResponse{}, err
	}
	return result, nil
}

// validateCustomDomainHandleError handles the ValidateCustomDomain error response.
func (client *AFDEndpointsClient) validateCustomDomainHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := AfdErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
