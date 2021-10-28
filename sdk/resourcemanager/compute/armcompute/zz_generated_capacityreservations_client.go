//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

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

// CapacityReservationsClient contains the methods for the CapacityReservations group.
// Don't use this type directly, use NewCapacityReservationsClient() instead.
type CapacityReservationsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewCapacityReservationsClient creates a new instance of CapacityReservationsClient with the specified values.
func NewCapacityReservationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *CapacityReservationsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &CapacityReservationsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreateOrUpdate - The operation to create or update a capacity reservation. Please note some properties can be set only during capacity reservation
// creation. Please refer to https://aka.ms/CapacityReservation for more
// details.
// If the operation fails it returns the *CloudError error type.
func (client *CapacityReservationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, capacityReservationName string, parameters CapacityReservation, options *CapacityReservationsBeginCreateOrUpdateOptions) (CapacityReservationsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, capacityReservationGroupName, capacityReservationName, parameters, options)
	if err != nil {
		return CapacityReservationsCreateOrUpdatePollerResponse{}, err
	}
	result := CapacityReservationsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("CapacityReservationsClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return CapacityReservationsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &CapacityReservationsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - The operation to create or update a capacity reservation. Please note some properties can be set only during capacity reservation creation.
// Please refer to https://aka.ms/CapacityReservation for more
// details.
// If the operation fails it returns the *CloudError error type.
func (client *CapacityReservationsClient) createOrUpdate(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, capacityReservationName string, parameters CapacityReservation, options *CapacityReservationsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, capacityReservationGroupName, capacityReservationName, parameters, options)
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
func (client *CapacityReservationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, capacityReservationName string, parameters CapacityReservation, options *CapacityReservationsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/capacityReservationGroups/{capacityReservationGroupName}/capacityReservations/{capacityReservationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if capacityReservationGroupName == "" {
		return nil, errors.New("parameter capacityReservationGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{capacityReservationGroupName}", url.PathEscape(capacityReservationGroupName))
	if capacityReservationName == "" {
		return nil, errors.New("parameter capacityReservationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{capacityReservationName}", url.PathEscape(capacityReservationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *CapacityReservationsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - The operation to delete a capacity reservation. This operation is allowed only when all the associated resources are disassociated from
// the capacity reservation. Please refer to
// https://aka.ms/CapacityReservation for more details.
// If the operation fails it returns the *CloudError error type.
func (client *CapacityReservationsClient) BeginDelete(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, capacityReservationName string, options *CapacityReservationsBeginDeleteOptions) (CapacityReservationsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, capacityReservationGroupName, capacityReservationName, options)
	if err != nil {
		return CapacityReservationsDeletePollerResponse{}, err
	}
	result := CapacityReservationsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("CapacityReservationsClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return CapacityReservationsDeletePollerResponse{}, err
	}
	result.Poller = &CapacityReservationsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - The operation to delete a capacity reservation. This operation is allowed only when all the associated resources are disassociated from the
// capacity reservation. Please refer to
// https://aka.ms/CapacityReservation for more details.
// If the operation fails it returns the *CloudError error type.
func (client *CapacityReservationsClient) deleteOperation(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, capacityReservationName string, options *CapacityReservationsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, capacityReservationGroupName, capacityReservationName, options)
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
func (client *CapacityReservationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, capacityReservationName string, options *CapacityReservationsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/capacityReservationGroups/{capacityReservationGroupName}/capacityReservations/{capacityReservationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if capacityReservationGroupName == "" {
		return nil, errors.New("parameter capacityReservationGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{capacityReservationGroupName}", url.PathEscape(capacityReservationGroupName))
	if capacityReservationName == "" {
		return nil, errors.New("parameter capacityReservationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{capacityReservationName}", url.PathEscape(capacityReservationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *CapacityReservationsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - The operation that retrieves information about the capacity reservation.
// If the operation fails it returns the *CloudError error type.
func (client *CapacityReservationsClient) Get(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, capacityReservationName string, options *CapacityReservationsGetOptions) (CapacityReservationsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, capacityReservationGroupName, capacityReservationName, options)
	if err != nil {
		return CapacityReservationsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CapacityReservationsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CapacityReservationsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CapacityReservationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, capacityReservationName string, options *CapacityReservationsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/capacityReservationGroups/{capacityReservationGroupName}/capacityReservations/{capacityReservationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if capacityReservationGroupName == "" {
		return nil, errors.New("parameter capacityReservationGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{capacityReservationGroupName}", url.PathEscape(capacityReservationGroupName))
	if capacityReservationName == "" {
		return nil, errors.New("parameter capacityReservationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{capacityReservationName}", url.PathEscape(capacityReservationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CapacityReservationsClient) getHandleResponse(resp *http.Response) (CapacityReservationsGetResponse, error) {
	result := CapacityReservationsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CapacityReservation); err != nil {
		return CapacityReservationsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *CapacityReservationsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByCapacityReservationGroup - Lists all of the capacity reservations in the specified capacity reservation group. Use the nextLink property in the
// response to get the next page of capacity reservations.
// If the operation fails it returns the *CloudError error type.
func (client *CapacityReservationsClient) ListByCapacityReservationGroup(resourceGroupName string, capacityReservationGroupName string, options *CapacityReservationsListByCapacityReservationGroupOptions) *CapacityReservationsListByCapacityReservationGroupPager {
	return &CapacityReservationsListByCapacityReservationGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByCapacityReservationGroupCreateRequest(ctx, resourceGroupName, capacityReservationGroupName, options)
		},
		advancer: func(ctx context.Context, resp CapacityReservationsListByCapacityReservationGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.CapacityReservationListResult.NextLink)
		},
	}
}

// listByCapacityReservationGroupCreateRequest creates the ListByCapacityReservationGroup request.
func (client *CapacityReservationsClient) listByCapacityReservationGroupCreateRequest(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, options *CapacityReservationsListByCapacityReservationGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/capacityReservationGroups/{capacityReservationGroupName}/capacityReservations"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if capacityReservationGroupName == "" {
		return nil, errors.New("parameter capacityReservationGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{capacityReservationGroupName}", url.PathEscape(capacityReservationGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByCapacityReservationGroupHandleResponse handles the ListByCapacityReservationGroup response.
func (client *CapacityReservationsClient) listByCapacityReservationGroupHandleResponse(resp *http.Response) (CapacityReservationsListByCapacityReservationGroupResponse, error) {
	result := CapacityReservationsListByCapacityReservationGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CapacityReservationListResult); err != nil {
		return CapacityReservationsListByCapacityReservationGroupResponse{}, err
	}
	return result, nil
}

// listByCapacityReservationGroupHandleError handles the ListByCapacityReservationGroup error response.
func (client *CapacityReservationsClient) listByCapacityReservationGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginUpdate - The operation to update a capacity reservation.
// If the operation fails it returns the *CloudError error type.
func (client *CapacityReservationsClient) BeginUpdate(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, capacityReservationName string, parameters CapacityReservationUpdate, options *CapacityReservationsBeginUpdateOptions) (CapacityReservationsUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, capacityReservationGroupName, capacityReservationName, parameters, options)
	if err != nil {
		return CapacityReservationsUpdatePollerResponse{}, err
	}
	result := CapacityReservationsUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("CapacityReservationsClient.Update", "", resp, client.pl, client.updateHandleError)
	if err != nil {
		return CapacityReservationsUpdatePollerResponse{}, err
	}
	result.Poller = &CapacityReservationsUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - The operation to update a capacity reservation.
// If the operation fails it returns the *CloudError error type.
func (client *CapacityReservationsClient) update(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, capacityReservationName string, parameters CapacityReservationUpdate, options *CapacityReservationsBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, capacityReservationGroupName, capacityReservationName, parameters, options)
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
func (client *CapacityReservationsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, capacityReservationName string, parameters CapacityReservationUpdate, options *CapacityReservationsBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/capacityReservationGroups/{capacityReservationGroupName}/capacityReservations/{capacityReservationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if capacityReservationGroupName == "" {
		return nil, errors.New("parameter capacityReservationGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{capacityReservationGroupName}", url.PathEscape(capacityReservationGroupName))
	if capacityReservationName == "" {
		return nil, errors.New("parameter capacityReservationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{capacityReservationName}", url.PathEscape(capacityReservationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleError handles the Update error response.
func (client *CapacityReservationsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
