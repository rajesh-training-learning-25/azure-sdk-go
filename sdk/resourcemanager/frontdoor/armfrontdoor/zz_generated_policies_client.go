//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfrontdoor

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// PoliciesClient contains the methods for the Policies group.
// Don't use this type directly, use NewPoliciesClient() instead.
type PoliciesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPoliciesClient creates a new instance of PoliciesClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *PoliciesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &PoliciesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Create or update policy with specified rule set name within a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// policyName - The name of the Web Application Firewall Policy.
// parameters - Policy to be created.
// options - PoliciesClientBeginCreateOrUpdateOptions contains the optional parameters for the PoliciesClient.BeginCreateOrUpdate
// method.
func (client *PoliciesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, policyName string, parameters WebApplicationFirewallPolicy, options *PoliciesClientBeginCreateOrUpdateOptions) (PoliciesClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, policyName, parameters, options)
	if err != nil {
		return PoliciesClientCreateOrUpdatePollerResponse{}, err
	}
	result := PoliciesClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("PoliciesClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return PoliciesClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &PoliciesClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Create or update policy with specified rule set name within a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *PoliciesClient) createOrUpdate(ctx context.Context, resourceGroupName string, policyName string, parameters WebApplicationFirewallPolicy, options *PoliciesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, policyName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, policyName string, parameters WebApplicationFirewallPolicy, options *PoliciesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/FrontDoorWebApplicationFirewallPolicies/{policyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if policyName == "" {
		return nil, errors.New("parameter policyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(policyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes Policy
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// policyName - The name of the Web Application Firewall Policy.
// options - PoliciesClientBeginDeleteOptions contains the optional parameters for the PoliciesClient.BeginDelete method.
func (client *PoliciesClient) BeginDelete(ctx context.Context, resourceGroupName string, policyName string, options *PoliciesClientBeginDeleteOptions) (PoliciesClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, policyName, options)
	if err != nil {
		return PoliciesClientDeletePollerResponse{}, err
	}
	result := PoliciesClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("PoliciesClient.Delete", "", resp, client.pl)
	if err != nil {
		return PoliciesClientDeletePollerResponse{}, err
	}
	result.Poller = &PoliciesClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes Policy
// If the operation fails it returns an *azcore.ResponseError type.
func (client *PoliciesClient) deleteOperation(ctx context.Context, resourceGroupName string, policyName string, options *PoliciesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, policyName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PoliciesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, policyName string, options *PoliciesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/FrontDoorWebApplicationFirewallPolicies/{policyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if policyName == "" {
		return nil, errors.New("parameter policyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(policyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Retrieve protection policy with specified name within a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// policyName - The name of the Web Application Firewall Policy.
// options - PoliciesClientGetOptions contains the optional parameters for the PoliciesClient.Get method.
func (client *PoliciesClient) Get(ctx context.Context, resourceGroupName string, policyName string, options *PoliciesClientGetOptions) (PoliciesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, policyName, options)
	if err != nil {
		return PoliciesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PoliciesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, policyName string, options *PoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/FrontDoorWebApplicationFirewallPolicies/{policyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if policyName == "" {
		return nil, errors.New("parameter policyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(policyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PoliciesClient) getHandleResponse(resp *http.Response) (PoliciesClientGetResponse, error) {
	result := PoliciesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebApplicationFirewallPolicy); err != nil {
		return PoliciesClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists all of the protection policies within a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// options - PoliciesClientListOptions contains the optional parameters for the PoliciesClient.List method.
func (client *PoliciesClient) List(resourceGroupName string, options *PoliciesClientListOptions) *PoliciesClientListPager {
	return &PoliciesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp PoliciesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.WebApplicationFirewallPolicyList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *PoliciesClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *PoliciesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoorWebApplicationFirewallPolicies"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
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
func (client *PoliciesClient) listHandleResponse(resp *http.Response) (PoliciesClientListResponse, error) {
	result := PoliciesClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebApplicationFirewallPolicyList); err != nil {
		return PoliciesClientListResponse{}, err
	}
	return result, nil
}
