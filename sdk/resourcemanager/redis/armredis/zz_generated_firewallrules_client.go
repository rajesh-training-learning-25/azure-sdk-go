//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armredis

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// FirewallRulesClient contains the methods for the FirewallRules group.
// Don't use this type directly, use NewFirewallRulesClient() instead.
type FirewallRulesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewFirewallRulesClient creates a new instance of FirewallRulesClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewFirewallRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*FirewallRulesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &FirewallRulesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update a redis cache firewall rule
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// cacheName - The name of the Redis cache.
// ruleName - The name of the firewall rule.
// parameters - Parameters supplied to the create or update redis firewall rule operation.
// options - FirewallRulesClientCreateOrUpdateOptions contains the optional parameters for the FirewallRulesClient.CreateOrUpdate
// method.
func (client *FirewallRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, cacheName string, ruleName string, parameters FirewallRule, options *FirewallRulesClientCreateOrUpdateOptions) (FirewallRulesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, cacheName, ruleName, parameters, options)
	if err != nil {
		return FirewallRulesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FirewallRulesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return FirewallRulesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *FirewallRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, cacheName string, ruleName string, parameters FirewallRule, options *FirewallRulesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{cacheName}/firewallRules/{ruleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cacheName == "" {
		return nil, errors.New("parameter cacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cacheName}", url.PathEscape(cacheName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *FirewallRulesClient) createOrUpdateHandleResponse(resp *http.Response) (FirewallRulesClientCreateOrUpdateResponse, error) {
	result := FirewallRulesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FirewallRule); err != nil {
		return FirewallRulesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a single firewall rule in a specified redis cache.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// cacheName - The name of the Redis cache.
// ruleName - The name of the firewall rule.
// options - FirewallRulesClientDeleteOptions contains the optional parameters for the FirewallRulesClient.Delete method.
func (client *FirewallRulesClient) Delete(ctx context.Context, resourceGroupName string, cacheName string, ruleName string, options *FirewallRulesClientDeleteOptions) (FirewallRulesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, cacheName, ruleName, options)
	if err != nil {
		return FirewallRulesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FirewallRulesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return FirewallRulesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return FirewallRulesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *FirewallRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, cacheName string, ruleName string, options *FirewallRulesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{cacheName}/firewallRules/{ruleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cacheName == "" {
		return nil, errors.New("parameter cacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cacheName}", url.PathEscape(cacheName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets a single firewall rule in a specified redis cache.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// cacheName - The name of the Redis cache.
// ruleName - The name of the firewall rule.
// options - FirewallRulesClientGetOptions contains the optional parameters for the FirewallRulesClient.Get method.
func (client *FirewallRulesClient) Get(ctx context.Context, resourceGroupName string, cacheName string, ruleName string, options *FirewallRulesClientGetOptions) (FirewallRulesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, cacheName, ruleName, options)
	if err != nil {
		return FirewallRulesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FirewallRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FirewallRulesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *FirewallRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, cacheName string, ruleName string, options *FirewallRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{cacheName}/firewallRules/{ruleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cacheName == "" {
		return nil, errors.New("parameter cacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cacheName}", url.PathEscape(cacheName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FirewallRulesClient) getHandleResponse(resp *http.Response) (FirewallRulesClientGetResponse, error) {
	result := FirewallRulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FirewallRule); err != nil {
		return FirewallRulesClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets all firewall rules in the specified redis cache.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// cacheName - The name of the Redis cache.
// options - FirewallRulesClientListOptions contains the optional parameters for the FirewallRulesClient.List method.
func (client *FirewallRulesClient) List(resourceGroupName string, cacheName string, options *FirewallRulesClientListOptions) *runtime.Pager[FirewallRulesClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[FirewallRulesClientListResponse]{
		More: func(page FirewallRulesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FirewallRulesClientListResponse) (FirewallRulesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, cacheName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return FirewallRulesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return FirewallRulesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return FirewallRulesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *FirewallRulesClient) listCreateRequest(ctx context.Context, resourceGroupName string, cacheName string, options *FirewallRulesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{cacheName}/firewallRules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cacheName == "" {
		return nil, errors.New("parameter cacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cacheName}", url.PathEscape(cacheName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *FirewallRulesClient) listHandleResponse(resp *http.Response) (FirewallRulesClientListResponse, error) {
	result := FirewallRulesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FirewallRuleListResult); err != nil {
		return FirewallRulesClientListResponse{}, err
	}
	return result, nil
}
