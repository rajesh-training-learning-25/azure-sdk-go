//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// ActivityLogAlertsClient contains the methods for the ActivityLogAlerts group.
// Don't use this type directly, use NewActivityLogAlertsClient() instead.
type ActivityLogAlertsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewActivityLogAlertsClient creates a new instance of ActivityLogAlertsClient with the specified values.
func NewActivityLogAlertsClient(con *arm.Connection, subscriptionID string) *ActivityLogAlertsClient {
	return &ActivityLogAlertsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// CreateOrUpdate - Create a new activity log alert or update an existing one.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ActivityLogAlertsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, activityLogAlertName string, activityLogAlert ActivityLogAlertResource, options *ActivityLogAlertsCreateOrUpdateOptions) (ActivityLogAlertsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, activityLogAlertName, activityLogAlert, options)
	if err != nil {
		return ActivityLogAlertsCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ActivityLogAlertsCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ActivityLogAlertsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ActivityLogAlertsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, activityLogAlertName string, activityLogAlert ActivityLogAlertResource, options *ActivityLogAlertsCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/activityLogAlerts/{activityLogAlertName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if activityLogAlertName == "" {
		return nil, errors.New("parameter activityLogAlertName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{activityLogAlertName}", url.PathEscape(activityLogAlertName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, activityLogAlert)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ActivityLogAlertsClient) createOrUpdateHandleResponse(resp *http.Response) (ActivityLogAlertsCreateOrUpdateResponse, error) {
	result := ActivityLogAlertsCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ActivityLogAlertResource); err != nil {
		return ActivityLogAlertsCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ActivityLogAlertsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Delete an activity log alert.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ActivityLogAlertsClient) Delete(ctx context.Context, resourceGroupName string, activityLogAlertName string, options *ActivityLogAlertsDeleteOptions) (ActivityLogAlertsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, activityLogAlertName, options)
	if err != nil {
		return ActivityLogAlertsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ActivityLogAlertsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ActivityLogAlertsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return ActivityLogAlertsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ActivityLogAlertsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, activityLogAlertName string, options *ActivityLogAlertsDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/activityLogAlerts/{activityLogAlertName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if activityLogAlertName == "" {
		return nil, errors.New("parameter activityLogAlertName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{activityLogAlertName}", url.PathEscape(activityLogAlertName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ActivityLogAlertsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Get an activity log alert.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ActivityLogAlertsClient) Get(ctx context.Context, resourceGroupName string, activityLogAlertName string, options *ActivityLogAlertsGetOptions) (ActivityLogAlertsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, activityLogAlertName, options)
	if err != nil {
		return ActivityLogAlertsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ActivityLogAlertsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ActivityLogAlertsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ActivityLogAlertsClient) getCreateRequest(ctx context.Context, resourceGroupName string, activityLogAlertName string, options *ActivityLogAlertsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/activityLogAlerts/{activityLogAlertName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if activityLogAlertName == "" {
		return nil, errors.New("parameter activityLogAlertName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{activityLogAlertName}", url.PathEscape(activityLogAlertName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ActivityLogAlertsClient) getHandleResponse(resp *http.Response) (ActivityLogAlertsGetResponse, error) {
	result := ActivityLogAlertsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ActivityLogAlertResource); err != nil {
		return ActivityLogAlertsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ActivityLogAlertsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByResourceGroup - Get a list of all activity log alerts in a resource group.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ActivityLogAlertsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, options *ActivityLogAlertsListByResourceGroupOptions) (ActivityLogAlertsListByResourceGroupResponse, error) {
	req, err := client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
	if err != nil {
		return ActivityLogAlertsListByResourceGroupResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ActivityLogAlertsListByResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ActivityLogAlertsListByResourceGroupResponse{}, client.listByResourceGroupHandleError(resp)
	}
	return client.listByResourceGroupHandleResponse(resp)
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ActivityLogAlertsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ActivityLogAlertsListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/activityLogAlerts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ActivityLogAlertsClient) listByResourceGroupHandleResponse(resp *http.Response) (ActivityLogAlertsListByResourceGroupResponse, error) {
	result := ActivityLogAlertsListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ActivityLogAlertList); err != nil {
		return ActivityLogAlertsListByResourceGroupResponse{}, err
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *ActivityLogAlertsClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListBySubscriptionID - Get a list of all activity log alerts in a subscription.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ActivityLogAlertsClient) ListBySubscriptionID(ctx context.Context, options *ActivityLogAlertsListBySubscriptionIDOptions) (ActivityLogAlertsListBySubscriptionIDResponse, error) {
	req, err := client.listBySubscriptionIDCreateRequest(ctx, options)
	if err != nil {
		return ActivityLogAlertsListBySubscriptionIDResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ActivityLogAlertsListBySubscriptionIDResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ActivityLogAlertsListBySubscriptionIDResponse{}, client.listBySubscriptionIDHandleError(resp)
	}
	return client.listBySubscriptionIDHandleResponse(resp)
}

// listBySubscriptionIDCreateRequest creates the ListBySubscriptionID request.
func (client *ActivityLogAlertsClient) listBySubscriptionIDCreateRequest(ctx context.Context, options *ActivityLogAlertsListBySubscriptionIDOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/microsoft.insights/activityLogAlerts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionIDHandleResponse handles the ListBySubscriptionID response.
func (client *ActivityLogAlertsClient) listBySubscriptionIDHandleResponse(resp *http.Response) (ActivityLogAlertsListBySubscriptionIDResponse, error) {
	result := ActivityLogAlertsListBySubscriptionIDResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ActivityLogAlertList); err != nil {
		return ActivityLogAlertsListBySubscriptionIDResponse{}, err
	}
	return result, nil
}

// listBySubscriptionIDHandleError handles the ListBySubscriptionID error response.
func (client *ActivityLogAlertsClient) listBySubscriptionIDHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Update - Updates an existing ActivityLogAlertResource's tags. To update other fields use the CreateOrUpdate method.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ActivityLogAlertsClient) Update(ctx context.Context, resourceGroupName string, activityLogAlertName string, activityLogAlertPatch ActivityLogAlertPatchBody, options *ActivityLogAlertsUpdateOptions) (ActivityLogAlertsUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, activityLogAlertName, activityLogAlertPatch, options)
	if err != nil {
		return ActivityLogAlertsUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ActivityLogAlertsUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ActivityLogAlertsUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ActivityLogAlertsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, activityLogAlertName string, activityLogAlertPatch ActivityLogAlertPatchBody, options *ActivityLogAlertsUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/activityLogAlerts/{activityLogAlertName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if activityLogAlertName == "" {
		return nil, errors.New("parameter activityLogAlertName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{activityLogAlertName}", url.PathEscape(activityLogAlertName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, activityLogAlertPatch)
}

// updateHandleResponse handles the Update response.
func (client *ActivityLogAlertsClient) updateHandleResponse(resp *http.Response) (ActivityLogAlertsUpdateResponse, error) {
	result := ActivityLogAlertsUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ActivityLogAlertResource); err != nil {
		return ActivityLogAlertsUpdateResponse{}, err
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *ActivityLogAlertsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
