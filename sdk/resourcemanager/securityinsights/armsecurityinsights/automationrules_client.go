//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsights

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AutomationRulesClient contains the methods for the AutomationRules group.
// Don't use this type directly, use NewAutomationRulesClient() instead.
type AutomationRulesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAutomationRulesClient creates a new instance of AutomationRulesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAutomationRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AutomationRulesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AutomationRulesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates the automation rule
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - automationRuleID - Automation rule ID
//   - options - AutomationRulesClientCreateOrUpdateOptions contains the optional parameters for the AutomationRulesClient.CreateOrUpdate
//     method.
func (client *AutomationRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, automationRuleID string, options *AutomationRulesClientCreateOrUpdateOptions) (AutomationRulesClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "AutomationRulesClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, automationRuleID, options)
	if err != nil {
		return AutomationRulesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AutomationRulesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return AutomationRulesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AutomationRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, automationRuleID string, options *AutomationRulesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/automationRules/{automationRuleId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if automationRuleID == "" {
		return nil, errors.New("parameter automationRuleID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationRuleId}", url.PathEscape(automationRuleID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.AutomationRuleToUpsert != nil {
		if err := runtime.MarshalAsJSON(req, *options.AutomationRuleToUpsert); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *AutomationRulesClient) createOrUpdateHandleResponse(resp *http.Response) (AutomationRulesClientCreateOrUpdateResponse, error) {
	result := AutomationRulesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AutomationRule); err != nil {
		return AutomationRulesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete the automation rule
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - automationRuleID - Automation rule ID
//   - options - AutomationRulesClientDeleteOptions contains the optional parameters for the AutomationRulesClient.Delete method.
func (client *AutomationRulesClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, automationRuleID string, options *AutomationRulesClientDeleteOptions) (AutomationRulesClientDeleteResponse, error) {
	var err error
	const operationName = "AutomationRulesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, automationRuleID, options)
	if err != nil {
		return AutomationRulesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AutomationRulesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return AutomationRulesClientDeleteResponse{}, err
	}
	resp, err := client.deleteHandleResponse(httpResp)
	return resp, err
}

// deleteCreateRequest creates the Delete request.
func (client *AutomationRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, automationRuleID string, options *AutomationRulesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/automationRules/{automationRuleId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if automationRuleID == "" {
		return nil, errors.New("parameter automationRuleID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationRuleId}", url.PathEscape(automationRuleID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *AutomationRulesClient) deleteHandleResponse(resp *http.Response) (AutomationRulesClientDeleteResponse, error) {
	result := AutomationRulesClientDeleteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Interface); err != nil {
		return AutomationRulesClientDeleteResponse{}, err
	}
	return result, nil
}

// Get - Gets the automation rule
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - automationRuleID - Automation rule ID
//   - options - AutomationRulesClientGetOptions contains the optional parameters for the AutomationRulesClient.Get method.
func (client *AutomationRulesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, automationRuleID string, options *AutomationRulesClientGetOptions) (AutomationRulesClientGetResponse, error) {
	var err error
	const operationName = "AutomationRulesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, automationRuleID, options)
	if err != nil {
		return AutomationRulesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AutomationRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AutomationRulesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AutomationRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, automationRuleID string, options *AutomationRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/automationRules/{automationRuleId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if automationRuleID == "" {
		return nil, errors.New("parameter automationRuleID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationRuleId}", url.PathEscape(automationRuleID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AutomationRulesClient) getHandleResponse(resp *http.Response) (AutomationRulesClientGetResponse, error) {
	result := AutomationRulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AutomationRule); err != nil {
		return AutomationRulesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all automation rules
//
// Generated from API version 2021-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - options - AutomationRulesClientListOptions contains the optional parameters for the AutomationRulesClient.NewListPager
//     method.
func (client *AutomationRulesClient) NewListPager(resourceGroupName string, workspaceName string, options *AutomationRulesClientListOptions) *runtime.Pager[AutomationRulesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AutomationRulesClientListResponse]{
		More: func(page AutomationRulesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AutomationRulesClientListResponse) (AutomationRulesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AutomationRulesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			}, nil)
			if err != nil {
				return AutomationRulesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *AutomationRulesClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *AutomationRulesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/automationRules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AutomationRulesClient) listHandleResponse(resp *http.Response) (AutomationRulesClientListResponse, error) {
	result := AutomationRulesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AutomationRulesList); err != nil {
		return AutomationRulesClientListResponse{}, err
	}
	return result, nil
}
