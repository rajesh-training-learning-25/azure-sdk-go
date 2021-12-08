//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armextendedlocation

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

// CustomLocationsClient contains the methods for the CustomLocations group.
// Don't use this type directly, use NewCustomLocationsClient() instead.
type CustomLocationsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewCustomLocationsClient creates a new instance of CustomLocationsClient with the specified values.
func NewCustomLocationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *CustomLocationsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &CustomLocationsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreateOrUpdate - Creates or updates a Custom Location in the specified Subscription and Resource Group
// If the operation fails it returns the *ErrorResponse error type.
func (client *CustomLocationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters CustomLocation, options *CustomLocationsBeginCreateOrUpdateOptions) (CustomLocationsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, resourceName, parameters, options)
	if err != nil {
		return CustomLocationsCreateOrUpdatePollerResponse{}, err
	}
	result := CustomLocationsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("CustomLocationsClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return CustomLocationsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &CustomLocationsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a Custom Location in the specified Subscription and Resource Group
// If the operation fails it returns the *ErrorResponse error type.
func (client *CustomLocationsClient) createOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters CustomLocation, options *CustomLocationsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *CustomLocationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, parameters CustomLocation, options *CustomLocationsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ExtendedLocation/customLocations/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *CustomLocationsClient) createOrUpdateHandleError(resp *http.Response) error {
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

// BeginDelete - Deletes the Custom Location with the specified Resource Name, Resource Group, and Subscription Id.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CustomLocationsClient) BeginDelete(ctx context.Context, resourceGroupName string, resourceName string, options *CustomLocationsBeginDeleteOptions) (CustomLocationsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return CustomLocationsDeletePollerResponse{}, err
	}
	result := CustomLocationsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("CustomLocationsClient.Delete", "azure-async-operation", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return CustomLocationsDeletePollerResponse{}, err
	}
	result.Poller = &CustomLocationsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the Custom Location with the specified Resource Name, Resource Group, and Subscription Id.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CustomLocationsClient) deleteOperation(ctx context.Context, resourceGroupName string, resourceName string, options *CustomLocationsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CustomLocationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *CustomLocationsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ExtendedLocation/customLocations/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *CustomLocationsClient) deleteHandleError(resp *http.Response) error {
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

// Get - Gets the details of the customLocation with a specified resource group and name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CustomLocationsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, options *CustomLocationsGetOptions) (CustomLocationsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return CustomLocationsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CustomLocationsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CustomLocationsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CustomLocationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *CustomLocationsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ExtendedLocation/customLocations/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CustomLocationsClient) getHandleResponse(resp *http.Response) (CustomLocationsGetResponse, error) {
	result := CustomLocationsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomLocation); err != nil {
		return CustomLocationsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *CustomLocationsClient) getHandleError(resp *http.Response) error {
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

// ListByResourceGroup - Gets a list of Custom Locations in the specified subscription and resource group. The operation returns properties of each Custom
// Location.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CustomLocationsClient) ListByResourceGroup(resourceGroupName string, options *CustomLocationsListByResourceGroupOptions) *CustomLocationsListByResourceGroupPager {
	return &CustomLocationsListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp CustomLocationsListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.CustomLocationListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *CustomLocationsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *CustomLocationsListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ExtendedLocation/customLocations"
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
	reqQP.Set("api-version", "2021-08-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *CustomLocationsClient) listByResourceGroupHandleResponse(resp *http.Response) (CustomLocationsListByResourceGroupResponse, error) {
	result := CustomLocationsListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomLocationListResult); err != nil {
		return CustomLocationsListByResourceGroupResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *CustomLocationsClient) listByResourceGroupHandleError(resp *http.Response) error {
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

// ListBySubscription - Gets a list of Custom Locations in the specified subscription. The operation returns properties of each Custom Location
// If the operation fails it returns the *ErrorResponse error type.
func (client *CustomLocationsClient) ListBySubscription(options *CustomLocationsListBySubscriptionOptions) *CustomLocationsListBySubscriptionPager {
	return &CustomLocationsListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp CustomLocationsListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.CustomLocationListResult.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *CustomLocationsClient) listBySubscriptionCreateRequest(ctx context.Context, options *CustomLocationsListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ExtendedLocation/customLocations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *CustomLocationsClient) listBySubscriptionHandleResponse(resp *http.Response) (CustomLocationsListBySubscriptionResponse, error) {
	result := CustomLocationsListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomLocationListResult); err != nil {
		return CustomLocationsListBySubscriptionResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *CustomLocationsClient) listBySubscriptionHandleError(resp *http.Response) error {
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

// ListEnabledResourceTypes - Gets the list of the Enabled Resource Types.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CustomLocationsClient) ListEnabledResourceTypes(resourceGroupName string, resourceName string, options *CustomLocationsListEnabledResourceTypesOptions) *CustomLocationsListEnabledResourceTypesPager {
	return &CustomLocationsListEnabledResourceTypesPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listEnabledResourceTypesCreateRequest(ctx, resourceGroupName, resourceName, options)
		},
		advancer: func(ctx context.Context, resp CustomLocationsListEnabledResourceTypesResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.EnabledResourceTypesListResult.NextLink)
		},
	}
}

// listEnabledResourceTypesCreateRequest creates the ListEnabledResourceTypes request.
func (client *CustomLocationsClient) listEnabledResourceTypesCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *CustomLocationsListEnabledResourceTypesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ExtendedLocation/customLocations/{resourceName}/enabledResourceTypes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listEnabledResourceTypesHandleResponse handles the ListEnabledResourceTypes response.
func (client *CustomLocationsClient) listEnabledResourceTypesHandleResponse(resp *http.Response) (CustomLocationsListEnabledResourceTypesResponse, error) {
	result := CustomLocationsListEnabledResourceTypesResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnabledResourceTypesListResult); err != nil {
		return CustomLocationsListEnabledResourceTypesResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listEnabledResourceTypesHandleError handles the ListEnabledResourceTypes error response.
func (client *CustomLocationsClient) listEnabledResourceTypesHandleError(resp *http.Response) error {
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

// ListOperations - Lists all available Custom Locations operations.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CustomLocationsClient) ListOperations(options *CustomLocationsListOperationsOptions) *CustomLocationsListOperationsPager {
	return &CustomLocationsListOperationsPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listOperationsCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp CustomLocationsListOperationsResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.CustomLocationOperationsList.NextLink)
		},
	}
}

// listOperationsCreateRequest creates the ListOperations request.
func (client *CustomLocationsClient) listOperationsCreateRequest(ctx context.Context, options *CustomLocationsListOperationsOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.ExtendedLocation/operations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listOperationsHandleResponse handles the ListOperations response.
func (client *CustomLocationsClient) listOperationsHandleResponse(resp *http.Response) (CustomLocationsListOperationsResponse, error) {
	result := CustomLocationsListOperationsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomLocationOperationsList); err != nil {
		return CustomLocationsListOperationsResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listOperationsHandleError handles the ListOperations error response.
func (client *CustomLocationsClient) listOperationsHandleError(resp *http.Response) error {
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

// Update - Updates a Custom Location with the specified Resource Name in the specified Resource Group and Subscription.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CustomLocationsClient) Update(ctx context.Context, resourceGroupName string, resourceName string, parameters PatchableCustomLocations, options *CustomLocationsUpdateOptions) (CustomLocationsUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceName, parameters, options)
	if err != nil {
		return CustomLocationsUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CustomLocationsUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CustomLocationsUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *CustomLocationsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, parameters PatchableCustomLocations, options *CustomLocationsUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ExtendedLocation/customLocations/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *CustomLocationsClient) updateHandleResponse(resp *http.Response) (CustomLocationsUpdateResponse, error) {
	result := CustomLocationsUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomLocation); err != nil {
		return CustomLocationsUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *CustomLocationsClient) updateHandleError(resp *http.Response) error {
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
