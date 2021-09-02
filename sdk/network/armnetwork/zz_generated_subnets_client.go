//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

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

// SubnetsClient contains the methods for the Subnets group.
// Don't use this type directly, use NewSubnetsClient() instead.
type SubnetsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewSubnetsClient creates a new instance of SubnetsClient with the specified values.
func NewSubnetsClient(con *arm.Connection, subscriptionID string) *SubnetsClient {
	return &SubnetsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a subnet in the specified virtual network.
// If the operation fails it returns the *CloudError error type.
func (client *SubnetsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, subnetParameters Subnet, options *SubnetsBeginCreateOrUpdateOptions) (SubnetsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, virtualNetworkName, subnetName, subnetParameters, options)
	if err != nil {
		return SubnetsCreateOrUpdatePollerResponse{}, err
	}
	result := SubnetsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SubnetsClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return SubnetsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &SubnetsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a subnet in the specified virtual network.
// If the operation fails it returns the *CloudError error type.
func (client *SubnetsClient) createOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, subnetParameters Subnet, options *SubnetsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, virtualNetworkName, subnetName, subnetParameters, options)
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
func (client *SubnetsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, subnetParameters Subnet, options *SubnetsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	if subnetName == "" {
		return nil, errors.New("parameter subnetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subnetName}", url.PathEscape(subnetName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, subnetParameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *SubnetsClient) createOrUpdateHandleError(resp *http.Response) error {
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

// BeginDelete - Deletes the specified subnet.
// If the operation fails it returns the *CloudError error type.
func (client *SubnetsClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, options *SubnetsBeginDeleteOptions) (SubnetsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, virtualNetworkName, subnetName, options)
	if err != nil {
		return SubnetsDeletePollerResponse{}, err
	}
	result := SubnetsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SubnetsClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return SubnetsDeletePollerResponse{}, err
	}
	result.Poller = &SubnetsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified subnet.
// If the operation fails it returns the *CloudError error type.
func (client *SubnetsClient) deleteOperation(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, options *SubnetsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualNetworkName, subnetName, options)
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
func (client *SubnetsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, options *SubnetsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	if subnetName == "" {
		return nil, errors.New("parameter subnetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subnetName}", url.PathEscape(subnetName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *SubnetsClient) deleteHandleError(resp *http.Response) error {
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

// Get - Gets the specified subnet by virtual network and resource group.
// If the operation fails it returns the *CloudError error type.
func (client *SubnetsClient) Get(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, options *SubnetsGetOptions) (SubnetsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualNetworkName, subnetName, options)
	if err != nil {
		return SubnetsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SubnetsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SubnetsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SubnetsClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, options *SubnetsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	if subnetName == "" {
		return nil, errors.New("parameter subnetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subnetName}", url.PathEscape(subnetName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SubnetsClient) getHandleResponse(resp *http.Response) (SubnetsGetResponse, error) {
	result := SubnetsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Subnet); err != nil {
		return SubnetsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *SubnetsClient) getHandleError(resp *http.Response) error {
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

// List - Gets all subnets in a virtual network.
// If the operation fails it returns the *CloudError error type.
func (client *SubnetsClient) List(resourceGroupName string, virtualNetworkName string, options *SubnetsListOptions) *SubnetsListPager {
	return &SubnetsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, virtualNetworkName, options)
		},
		advancer: func(ctx context.Context, resp SubnetsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SubnetListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *SubnetsClient) listCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, options *SubnetsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SubnetsClient) listHandleResponse(resp *http.Response) (SubnetsListResponse, error) {
	result := SubnetsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubnetListResult); err != nil {
		return SubnetsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *SubnetsClient) listHandleError(resp *http.Response) error {
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

// BeginPrepareNetworkPolicies - Prepares a subnet by applying network intent policies.
// If the operation fails it returns the *CloudError error type.
func (client *SubnetsClient) BeginPrepareNetworkPolicies(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, prepareNetworkPoliciesRequestParameters PrepareNetworkPoliciesRequest, options *SubnetsBeginPrepareNetworkPoliciesOptions) (SubnetsPrepareNetworkPoliciesPollerResponse, error) {
	resp, err := client.prepareNetworkPolicies(ctx, resourceGroupName, virtualNetworkName, subnetName, prepareNetworkPoliciesRequestParameters, options)
	if err != nil {
		return SubnetsPrepareNetworkPoliciesPollerResponse{}, err
	}
	result := SubnetsPrepareNetworkPoliciesPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SubnetsClient.PrepareNetworkPolicies", "location", resp, client.pl, client.prepareNetworkPoliciesHandleError)
	if err != nil {
		return SubnetsPrepareNetworkPoliciesPollerResponse{}, err
	}
	result.Poller = &SubnetsPrepareNetworkPoliciesPoller{
		pt: pt,
	}
	return result, nil
}

// PrepareNetworkPolicies - Prepares a subnet by applying network intent policies.
// If the operation fails it returns the *CloudError error type.
func (client *SubnetsClient) prepareNetworkPolicies(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, prepareNetworkPoliciesRequestParameters PrepareNetworkPoliciesRequest, options *SubnetsBeginPrepareNetworkPoliciesOptions) (*http.Response, error) {
	req, err := client.prepareNetworkPoliciesCreateRequest(ctx, resourceGroupName, virtualNetworkName, subnetName, prepareNetworkPoliciesRequestParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.prepareNetworkPoliciesHandleError(resp)
	}
	return resp, nil
}

// prepareNetworkPoliciesCreateRequest creates the PrepareNetworkPolicies request.
func (client *SubnetsClient) prepareNetworkPoliciesCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, prepareNetworkPoliciesRequestParameters PrepareNetworkPoliciesRequest, options *SubnetsBeginPrepareNetworkPoliciesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}/PrepareNetworkPolicies"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	if subnetName == "" {
		return nil, errors.New("parameter subnetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subnetName}", url.PathEscape(subnetName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, prepareNetworkPoliciesRequestParameters)
}

// prepareNetworkPoliciesHandleError handles the PrepareNetworkPolicies error response.
func (client *SubnetsClient) prepareNetworkPoliciesHandleError(resp *http.Response) error {
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

// BeginUnprepareNetworkPolicies - Unprepares a subnet by removing network intent policies.
// If the operation fails it returns the *CloudError error type.
func (client *SubnetsClient) BeginUnprepareNetworkPolicies(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, unprepareNetworkPoliciesRequestParameters UnprepareNetworkPoliciesRequest, options *SubnetsBeginUnprepareNetworkPoliciesOptions) (SubnetsUnprepareNetworkPoliciesPollerResponse, error) {
	resp, err := client.unprepareNetworkPolicies(ctx, resourceGroupName, virtualNetworkName, subnetName, unprepareNetworkPoliciesRequestParameters, options)
	if err != nil {
		return SubnetsUnprepareNetworkPoliciesPollerResponse{}, err
	}
	result := SubnetsUnprepareNetworkPoliciesPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SubnetsClient.UnprepareNetworkPolicies", "location", resp, client.pl, client.unprepareNetworkPoliciesHandleError)
	if err != nil {
		return SubnetsUnprepareNetworkPoliciesPollerResponse{}, err
	}
	result.Poller = &SubnetsUnprepareNetworkPoliciesPoller{
		pt: pt,
	}
	return result, nil
}

// UnprepareNetworkPolicies - Unprepares a subnet by removing network intent policies.
// If the operation fails it returns the *CloudError error type.
func (client *SubnetsClient) unprepareNetworkPolicies(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, unprepareNetworkPoliciesRequestParameters UnprepareNetworkPoliciesRequest, options *SubnetsBeginUnprepareNetworkPoliciesOptions) (*http.Response, error) {
	req, err := client.unprepareNetworkPoliciesCreateRequest(ctx, resourceGroupName, virtualNetworkName, subnetName, unprepareNetworkPoliciesRequestParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.unprepareNetworkPoliciesHandleError(resp)
	}
	return resp, nil
}

// unprepareNetworkPoliciesCreateRequest creates the UnprepareNetworkPolicies request.
func (client *SubnetsClient) unprepareNetworkPoliciesCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, unprepareNetworkPoliciesRequestParameters UnprepareNetworkPoliciesRequest, options *SubnetsBeginUnprepareNetworkPoliciesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}/UnprepareNetworkPolicies"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	if subnetName == "" {
		return nil, errors.New("parameter subnetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subnetName}", url.PathEscape(subnetName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, unprepareNetworkPoliciesRequestParameters)
}

// unprepareNetworkPoliciesHandleError handles the UnprepareNetworkPolicies error response.
func (client *SubnetsClient) unprepareNetworkPoliciesHandleError(resp *http.Response) error {
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
