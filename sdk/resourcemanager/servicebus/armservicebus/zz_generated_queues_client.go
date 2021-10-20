//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicebus

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

// QueuesClient contains the methods for the Queues group.
// Don't use this type directly, use NewQueuesClient() instead.
type QueuesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewQueuesClient creates a new instance of QueuesClient with the specified values.
func NewQueuesClient(con *arm.Connection, subscriptionID string) *QueuesClient {
	return &QueuesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// CreateOrUpdate - Creates or updates a Service Bus queue. This operation is idempotent.
// If the operation fails it returns the *ErrorResponse error type.
func (client *QueuesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, parameters SBQueue, options *QueuesCreateOrUpdateOptions) (QueuesCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, namespaceName, queueName, parameters, options)
	if err != nil {
		return QueuesCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return QueuesCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return QueuesCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *QueuesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, parameters SBQueue, options *QueuesCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if queueName == "" {
		return nil, errors.New("parameter queueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queueName}", url.PathEscape(queueName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *QueuesClient) createOrUpdateHandleResponse(resp *http.Response) (QueuesCreateOrUpdateResponse, error) {
	result := QueuesCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SBQueue); err != nil {
		return QueuesCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *QueuesClient) createOrUpdateHandleError(resp *http.Response) error {
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

// CreateOrUpdateAuthorizationRule - Creates an authorization rule for a queue.
// If the operation fails it returns the *ErrorResponse error type.
func (client *QueuesClient) CreateOrUpdateAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, authorizationRuleName string, parameters SBAuthorizationRule, options *QueuesCreateOrUpdateAuthorizationRuleOptions) (QueuesCreateOrUpdateAuthorizationRuleResponse, error) {
	req, err := client.createOrUpdateAuthorizationRuleCreateRequest(ctx, resourceGroupName, namespaceName, queueName, authorizationRuleName, parameters, options)
	if err != nil {
		return QueuesCreateOrUpdateAuthorizationRuleResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return QueuesCreateOrUpdateAuthorizationRuleResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return QueuesCreateOrUpdateAuthorizationRuleResponse{}, client.createOrUpdateAuthorizationRuleHandleError(resp)
	}
	return client.createOrUpdateAuthorizationRuleHandleResponse(resp)
}

// createOrUpdateAuthorizationRuleCreateRequest creates the CreateOrUpdateAuthorizationRule request.
func (client *QueuesClient) createOrUpdateAuthorizationRuleCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, authorizationRuleName string, parameters SBAuthorizationRule, options *QueuesCreateOrUpdateAuthorizationRuleOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}/authorizationRules/{authorizationRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if queueName == "" {
		return nil, errors.New("parameter queueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queueName}", url.PathEscape(queueName))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateAuthorizationRuleHandleResponse handles the CreateOrUpdateAuthorizationRule response.
func (client *QueuesClient) createOrUpdateAuthorizationRuleHandleResponse(resp *http.Response) (QueuesCreateOrUpdateAuthorizationRuleResponse, error) {
	result := QueuesCreateOrUpdateAuthorizationRuleResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SBAuthorizationRule); err != nil {
		return QueuesCreateOrUpdateAuthorizationRuleResponse{}, err
	}
	return result, nil
}

// createOrUpdateAuthorizationRuleHandleError handles the CreateOrUpdateAuthorizationRule error response.
func (client *QueuesClient) createOrUpdateAuthorizationRuleHandleError(resp *http.Response) error {
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

// Delete - Deletes a queue from the specified namespace in a resource group.
// If the operation fails it returns the *ErrorResponse error type.
func (client *QueuesClient) Delete(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, options *QueuesDeleteOptions) (QueuesDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, namespaceName, queueName, options)
	if err != nil {
		return QueuesDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return QueuesDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return QueuesDeleteResponse{}, client.deleteHandleError(resp)
	}
	return QueuesDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *QueuesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, options *QueuesDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if queueName == "" {
		return nil, errors.New("parameter queueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queueName}", url.PathEscape(queueName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *QueuesClient) deleteHandleError(resp *http.Response) error {
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

// DeleteAuthorizationRule - Deletes a queue authorization rule.
// If the operation fails it returns the *ErrorResponse error type.
func (client *QueuesClient) DeleteAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, authorizationRuleName string, options *QueuesDeleteAuthorizationRuleOptions) (QueuesDeleteAuthorizationRuleResponse, error) {
	req, err := client.deleteAuthorizationRuleCreateRequest(ctx, resourceGroupName, namespaceName, queueName, authorizationRuleName, options)
	if err != nil {
		return QueuesDeleteAuthorizationRuleResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return QueuesDeleteAuthorizationRuleResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return QueuesDeleteAuthorizationRuleResponse{}, client.deleteAuthorizationRuleHandleError(resp)
	}
	return QueuesDeleteAuthorizationRuleResponse{RawResponse: resp}, nil
}

// deleteAuthorizationRuleCreateRequest creates the DeleteAuthorizationRule request.
func (client *QueuesClient) deleteAuthorizationRuleCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, authorizationRuleName string, options *QueuesDeleteAuthorizationRuleOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}/authorizationRules/{authorizationRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if queueName == "" {
		return nil, errors.New("parameter queueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queueName}", url.PathEscape(queueName))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteAuthorizationRuleHandleError handles the DeleteAuthorizationRule error response.
func (client *QueuesClient) deleteAuthorizationRuleHandleError(resp *http.Response) error {
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

// Get - Returns a description for the specified queue.
// If the operation fails it returns the *ErrorResponse error type.
func (client *QueuesClient) Get(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, options *QueuesGetOptions) (QueuesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, namespaceName, queueName, options)
	if err != nil {
		return QueuesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return QueuesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return QueuesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *QueuesClient) getCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, options *QueuesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if queueName == "" {
		return nil, errors.New("parameter queueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queueName}", url.PathEscape(queueName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *QueuesClient) getHandleResponse(resp *http.Response) (QueuesGetResponse, error) {
	result := QueuesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SBQueue); err != nil {
		return QueuesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *QueuesClient) getHandleError(resp *http.Response) error {
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

// GetAuthorizationRule - Gets an authorization rule for a queue by rule name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *QueuesClient) GetAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, authorizationRuleName string, options *QueuesGetAuthorizationRuleOptions) (QueuesGetAuthorizationRuleResponse, error) {
	req, err := client.getAuthorizationRuleCreateRequest(ctx, resourceGroupName, namespaceName, queueName, authorizationRuleName, options)
	if err != nil {
		return QueuesGetAuthorizationRuleResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return QueuesGetAuthorizationRuleResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return QueuesGetAuthorizationRuleResponse{}, client.getAuthorizationRuleHandleError(resp)
	}
	return client.getAuthorizationRuleHandleResponse(resp)
}

// getAuthorizationRuleCreateRequest creates the GetAuthorizationRule request.
func (client *QueuesClient) getAuthorizationRuleCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, authorizationRuleName string, options *QueuesGetAuthorizationRuleOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}/authorizationRules/{authorizationRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if queueName == "" {
		return nil, errors.New("parameter queueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queueName}", url.PathEscape(queueName))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getAuthorizationRuleHandleResponse handles the GetAuthorizationRule response.
func (client *QueuesClient) getAuthorizationRuleHandleResponse(resp *http.Response) (QueuesGetAuthorizationRuleResponse, error) {
	result := QueuesGetAuthorizationRuleResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SBAuthorizationRule); err != nil {
		return QueuesGetAuthorizationRuleResponse{}, err
	}
	return result, nil
}

// getAuthorizationRuleHandleError handles the GetAuthorizationRule error response.
func (client *QueuesClient) getAuthorizationRuleHandleError(resp *http.Response) error {
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

// ListAuthorizationRules - Gets all authorization rules for a queue.
// If the operation fails it returns the *ErrorResponse error type.
func (client *QueuesClient) ListAuthorizationRules(resourceGroupName string, namespaceName string, queueName string, options *QueuesListAuthorizationRulesOptions) *QueuesListAuthorizationRulesPager {
	return &QueuesListAuthorizationRulesPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listAuthorizationRulesCreateRequest(ctx, resourceGroupName, namespaceName, queueName, options)
		},
		advancer: func(ctx context.Context, resp QueuesListAuthorizationRulesResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SBAuthorizationRuleListResult.NextLink)
		},
	}
}

// listAuthorizationRulesCreateRequest creates the ListAuthorizationRules request.
func (client *QueuesClient) listAuthorizationRulesCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, options *QueuesListAuthorizationRulesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}/authorizationRules"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if queueName == "" {
		return nil, errors.New("parameter queueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queueName}", url.PathEscape(queueName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listAuthorizationRulesHandleResponse handles the ListAuthorizationRules response.
func (client *QueuesClient) listAuthorizationRulesHandleResponse(resp *http.Response) (QueuesListAuthorizationRulesResponse, error) {
	result := QueuesListAuthorizationRulesResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SBAuthorizationRuleListResult); err != nil {
		return QueuesListAuthorizationRulesResponse{}, err
	}
	return result, nil
}

// listAuthorizationRulesHandleError handles the ListAuthorizationRules error response.
func (client *QueuesClient) listAuthorizationRulesHandleError(resp *http.Response) error {
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

// ListByNamespace - Gets the queues within a namespace.
// If the operation fails it returns the *ErrorResponse error type.
func (client *QueuesClient) ListByNamespace(resourceGroupName string, namespaceName string, options *QueuesListByNamespaceOptions) *QueuesListByNamespacePager {
	return &QueuesListByNamespacePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByNamespaceCreateRequest(ctx, resourceGroupName, namespaceName, options)
		},
		advancer: func(ctx context.Context, resp QueuesListByNamespaceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SBQueueListResult.NextLink)
		},
	}
}

// listByNamespaceCreateRequest creates the ListByNamespace request.
func (client *QueuesClient) listByNamespaceCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, options *QueuesListByNamespaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByNamespaceHandleResponse handles the ListByNamespace response.
func (client *QueuesClient) listByNamespaceHandleResponse(resp *http.Response) (QueuesListByNamespaceResponse, error) {
	result := QueuesListByNamespaceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SBQueueListResult); err != nil {
		return QueuesListByNamespaceResponse{}, err
	}
	return result, nil
}

// listByNamespaceHandleError handles the ListByNamespace error response.
func (client *QueuesClient) listByNamespaceHandleError(resp *http.Response) error {
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

// ListKeys - Primary and secondary connection strings to the queue.
// If the operation fails it returns the *ErrorResponse error type.
func (client *QueuesClient) ListKeys(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, authorizationRuleName string, options *QueuesListKeysOptions) (QueuesListKeysResponse, error) {
	req, err := client.listKeysCreateRequest(ctx, resourceGroupName, namespaceName, queueName, authorizationRuleName, options)
	if err != nil {
		return QueuesListKeysResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return QueuesListKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return QueuesListKeysResponse{}, client.listKeysHandleError(resp)
	}
	return client.listKeysHandleResponse(resp)
}

// listKeysCreateRequest creates the ListKeys request.
func (client *QueuesClient) listKeysCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, authorizationRuleName string, options *QueuesListKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}/authorizationRules/{authorizationRuleName}/ListKeys"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if queueName == "" {
		return nil, errors.New("parameter queueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queueName}", url.PathEscape(queueName))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listKeysHandleResponse handles the ListKeys response.
func (client *QueuesClient) listKeysHandleResponse(resp *http.Response) (QueuesListKeysResponse, error) {
	result := QueuesListKeysResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessKeys); err != nil {
		return QueuesListKeysResponse{}, err
	}
	return result, nil
}

// listKeysHandleError handles the ListKeys error response.
func (client *QueuesClient) listKeysHandleError(resp *http.Response) error {
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

// RegenerateKeys - Regenerates the primary or secondary connection strings to the queue.
// If the operation fails it returns the *ErrorResponse error type.
func (client *QueuesClient) RegenerateKeys(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, authorizationRuleName string, parameters RegenerateAccessKeyParameters, options *QueuesRegenerateKeysOptions) (QueuesRegenerateKeysResponse, error) {
	req, err := client.regenerateKeysCreateRequest(ctx, resourceGroupName, namespaceName, queueName, authorizationRuleName, parameters, options)
	if err != nil {
		return QueuesRegenerateKeysResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return QueuesRegenerateKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return QueuesRegenerateKeysResponse{}, client.regenerateKeysHandleError(resp)
	}
	return client.regenerateKeysHandleResponse(resp)
}

// regenerateKeysCreateRequest creates the RegenerateKeys request.
func (client *QueuesClient) regenerateKeysCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, authorizationRuleName string, parameters RegenerateAccessKeyParameters, options *QueuesRegenerateKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}/authorizationRules/{authorizationRuleName}/regenerateKeys"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if queueName == "" {
		return nil, errors.New("parameter queueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queueName}", url.PathEscape(queueName))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// regenerateKeysHandleResponse handles the RegenerateKeys response.
func (client *QueuesClient) regenerateKeysHandleResponse(resp *http.Response) (QueuesRegenerateKeysResponse, error) {
	result := QueuesRegenerateKeysResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessKeys); err != nil {
		return QueuesRegenerateKeysResponse{}, err
	}
	return result, nil
}

// regenerateKeysHandleError handles the RegenerateKeys error response.
func (client *QueuesClient) regenerateKeysHandleError(resp *http.Response) error {
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
