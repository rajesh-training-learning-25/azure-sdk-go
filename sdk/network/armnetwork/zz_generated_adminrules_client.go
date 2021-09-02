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
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// AdminRulesClient contains the methods for the AdminRules group.
// Don't use this type directly, use NewAdminRulesClient() instead.
type AdminRulesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewAdminRulesClient creates a new instance of AdminRulesClient with the specified values.
func NewAdminRulesClient(con *arm.Connection, subscriptionID string) *AdminRulesClient {
	return &AdminRulesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// CreateOrUpdate - Creates or updates an admin rule.
// If the operation fails it returns the *CloudErrorAutoGenerated error type.
func (client *AdminRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, ruleCollectionName string, ruleName string, adminRule BaseAdminRuleClassification, options *AdminRulesCreateOrUpdateOptions) (AdminRulesCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, networkManagerName, configurationName, ruleCollectionName, ruleName, adminRule, options)
	if err != nil {
		return AdminRulesCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AdminRulesCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return AdminRulesCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AdminRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, ruleCollectionName string, ruleName string, adminRule BaseAdminRuleClassification, options *AdminRulesCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/securityAdminConfigurations/{configurationName}/ruleCollections/{ruleCollectionName}/rules/{ruleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkManagerName == "" {
		return nil, errors.New("parameter networkManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkManagerName}", url.PathEscape(networkManagerName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	if ruleCollectionName == "" {
		return nil, errors.New("parameter ruleCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleCollectionName}", url.PathEscape(ruleCollectionName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, adminRule)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *AdminRulesClient) createOrUpdateHandleResponse(resp *http.Response) (AdminRulesCreateOrUpdateResponse, error) {
	result := AdminRulesCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return AdminRulesCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *AdminRulesClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudErrorAutoGenerated{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Deletes an admin rule.
// If the operation fails it returns the *CloudErrorAutoGenerated error type.
func (client *AdminRulesClient) Delete(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, ruleCollectionName string, ruleName string, options *AdminRulesDeleteOptions) (AdminRulesDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkManagerName, configurationName, ruleCollectionName, ruleName, options)
	if err != nil {
		return AdminRulesDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AdminRulesDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return AdminRulesDeleteResponse{}, client.deleteHandleError(resp)
	}
	return AdminRulesDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AdminRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, ruleCollectionName string, ruleName string, options *AdminRulesDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/securityAdminConfigurations/{configurationName}/ruleCollections/{ruleCollectionName}/rules/{ruleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkManagerName == "" {
		return nil, errors.New("parameter networkManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkManagerName}", url.PathEscape(networkManagerName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	if ruleCollectionName == "" {
		return nil, errors.New("parameter ruleCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleCollectionName}", url.PathEscape(ruleCollectionName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *AdminRulesClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudErrorAutoGenerated{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets a network manager security configuration admin rule.
// If the operation fails it returns the *CloudErrorAutoGenerated error type.
func (client *AdminRulesClient) Get(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, ruleCollectionName string, ruleName string, options *AdminRulesGetOptions) (AdminRulesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkManagerName, configurationName, ruleCollectionName, ruleName, options)
	if err != nil {
		return AdminRulesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AdminRulesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AdminRulesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AdminRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, ruleCollectionName string, ruleName string, options *AdminRulesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/securityAdminConfigurations/{configurationName}/ruleCollections/{ruleCollectionName}/rules/{ruleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkManagerName == "" {
		return nil, errors.New("parameter networkManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkManagerName}", url.PathEscape(networkManagerName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	if ruleCollectionName == "" {
		return nil, errors.New("parameter ruleCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleCollectionName}", url.PathEscape(ruleCollectionName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AdminRulesClient) getHandleResponse(resp *http.Response) (AdminRulesGetResponse, error) {
	result := AdminRulesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return AdminRulesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *AdminRulesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudErrorAutoGenerated{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - List all network manager security configuration admin rules.
// If the operation fails it returns the *CloudErrorAutoGenerated error type.
func (client *AdminRulesClient) List(resourceGroupName string, networkManagerName string, configurationName string, ruleCollectionName string, options *AdminRulesListOptions) *AdminRulesListPager {
	return &AdminRulesListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, networkManagerName, configurationName, ruleCollectionName, options)
		},
		advancer: func(ctx context.Context, resp AdminRulesListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AdminRuleListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *AdminRulesClient) listCreateRequest(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, ruleCollectionName string, options *AdminRulesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/securityAdminConfigurations/{configurationName}/ruleCollections/{ruleCollectionName}/rules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkManagerName == "" {
		return nil, errors.New("parameter networkManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkManagerName}", url.PathEscape(networkManagerName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	if ruleCollectionName == "" {
		return nil, errors.New("parameter ruleCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleCollectionName}", url.PathEscape(ruleCollectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AdminRulesClient) listHandleResponse(resp *http.Response) (AdminRulesListResponse, error) {
	result := AdminRulesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AdminRuleListResult); err != nil {
		return AdminRulesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *AdminRulesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudErrorAutoGenerated{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
