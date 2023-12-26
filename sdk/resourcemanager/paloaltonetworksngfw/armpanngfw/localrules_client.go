//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpanngfw

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

// LocalRulesClient contains the methods for the LocalRules group.
// Don't use this type directly, use NewLocalRulesClient() instead.
type LocalRulesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewLocalRulesClient creates a new instance of LocalRulesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewLocalRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LocalRulesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &LocalRulesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a LocalRulesResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - localRulestackName - LocalRulestack resource name
//   - priority - Local Rule priority
//   - resource - Resource create parameters.
//   - options - LocalRulesClientBeginCreateOrUpdateOptions contains the optional parameters for the LocalRulesClient.BeginCreateOrUpdate
//     method.
func (client *LocalRulesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, localRulestackName string, priority string, resource LocalRulesResource, options *LocalRulesClientBeginCreateOrUpdateOptions) (*runtime.Poller[LocalRulesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, localRulestackName, priority, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[LocalRulesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[LocalRulesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a LocalRulesResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
func (client *LocalRulesClient) createOrUpdate(ctx context.Context, resourceGroupName string, localRulestackName string, priority string, resource LocalRulesResource, options *LocalRulesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "LocalRulesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, localRulestackName, priority, resource, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *LocalRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, localRulestackName string, priority string, resource LocalRulesResource, options *LocalRulesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/{localRulestackName}/localRules/{priority}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if localRulestackName == "" {
		return nil, errors.New("parameter localRulestackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{localRulestackName}", url.PathEscape(localRulestackName))
	if priority == "" {
		return nil, errors.New("parameter priority cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{priority}", url.PathEscape(priority))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a LocalRulesResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - localRulestackName - LocalRulestack resource name
//   - priority - Local Rule priority
//   - options - LocalRulesClientBeginDeleteOptions contains the optional parameters for the LocalRulesClient.BeginDelete method.
func (client *LocalRulesClient) BeginDelete(ctx context.Context, resourceGroupName string, localRulestackName string, priority string, options *LocalRulesClientBeginDeleteOptions) (*runtime.Poller[LocalRulesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, localRulestackName, priority, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[LocalRulesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[LocalRulesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a LocalRulesResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
func (client *LocalRulesClient) deleteOperation(ctx context.Context, resourceGroupName string, localRulestackName string, priority string, options *LocalRulesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "LocalRulesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, localRulestackName, priority, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *LocalRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, localRulestackName string, priority string, options *LocalRulesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/{localRulestackName}/localRules/{priority}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if localRulestackName == "" {
		return nil, errors.New("parameter localRulestackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{localRulestackName}", url.PathEscape(localRulestackName))
	if priority == "" {
		return nil, errors.New("parameter priority cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{priority}", url.PathEscape(priority))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a LocalRulesResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - localRulestackName - LocalRulestack resource name
//   - priority - Local Rule priority
//   - options - LocalRulesClientGetOptions contains the optional parameters for the LocalRulesClient.Get method.
func (client *LocalRulesClient) Get(ctx context.Context, resourceGroupName string, localRulestackName string, priority string, options *LocalRulesClientGetOptions) (LocalRulesClientGetResponse, error) {
	var err error
	const operationName = "LocalRulesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, localRulestackName, priority, options)
	if err != nil {
		return LocalRulesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LocalRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LocalRulesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *LocalRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, localRulestackName string, priority string, options *LocalRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/{localRulestackName}/localRules/{priority}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if localRulestackName == "" {
		return nil, errors.New("parameter localRulestackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{localRulestackName}", url.PathEscape(localRulestackName))
	if priority == "" {
		return nil, errors.New("parameter priority cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{priority}", url.PathEscape(priority))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LocalRulesClient) getHandleResponse(resp *http.Response) (LocalRulesClientGetResponse, error) {
	result := LocalRulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LocalRulesResource); err != nil {
		return LocalRulesClientGetResponse{}, err
	}
	return result, nil
}

// GetCounters - Get counters
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - localRulestackName - LocalRulestack resource name
//   - priority - Local Rule priority
//   - options - LocalRulesClientGetCountersOptions contains the optional parameters for the LocalRulesClient.GetCounters method.
func (client *LocalRulesClient) GetCounters(ctx context.Context, resourceGroupName string, localRulestackName string, priority string, options *LocalRulesClientGetCountersOptions) (LocalRulesClientGetCountersResponse, error) {
	var err error
	const operationName = "LocalRulesClient.GetCounters"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCountersCreateRequest(ctx, resourceGroupName, localRulestackName, priority, options)
	if err != nil {
		return LocalRulesClientGetCountersResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LocalRulesClientGetCountersResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LocalRulesClientGetCountersResponse{}, err
	}
	resp, err := client.getCountersHandleResponse(httpResp)
	return resp, err
}

// getCountersCreateRequest creates the GetCounters request.
func (client *LocalRulesClient) getCountersCreateRequest(ctx context.Context, resourceGroupName string, localRulestackName string, priority string, options *LocalRulesClientGetCountersOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/{localRulestackName}/localRules/{priority}/getCounters"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if localRulestackName == "" {
		return nil, errors.New("parameter localRulestackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{localRulestackName}", url.PathEscape(localRulestackName))
	if priority == "" {
		return nil, errors.New("parameter priority cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{priority}", url.PathEscape(priority))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	if options != nil && options.FirewallName != nil {
		reqQP.Set("firewallName", *options.FirewallName)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getCountersHandleResponse handles the GetCounters response.
func (client *LocalRulesClient) getCountersHandleResponse(resp *http.Response) (LocalRulesClientGetCountersResponse, error) {
	result := LocalRulesClientGetCountersResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RuleCounter); err != nil {
		return LocalRulesClientGetCountersResponse{}, err
	}
	return result, nil
}

// NewListByLocalRulestacksPager - List LocalRulesResource resources by LocalRulestacks
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - localRulestackName - LocalRulestack resource name
//   - options - LocalRulesClientListByLocalRulestacksOptions contains the optional parameters for the LocalRulesClient.NewListByLocalRulestacksPager
//     method.
func (client *LocalRulesClient) NewListByLocalRulestacksPager(resourceGroupName string, localRulestackName string, options *LocalRulesClientListByLocalRulestacksOptions) *runtime.Pager[LocalRulesClientListByLocalRulestacksResponse] {
	return runtime.NewPager(runtime.PagingHandler[LocalRulesClientListByLocalRulestacksResponse]{
		More: func(page LocalRulesClientListByLocalRulestacksResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LocalRulesClientListByLocalRulestacksResponse) (LocalRulesClientListByLocalRulestacksResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "LocalRulesClient.NewListByLocalRulestacksPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByLocalRulestacksCreateRequest(ctx, resourceGroupName, localRulestackName, options)
			}, nil)
			if err != nil {
				return LocalRulesClientListByLocalRulestacksResponse{}, err
			}
			return client.listByLocalRulestacksHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByLocalRulestacksCreateRequest creates the ListByLocalRulestacks request.
func (client *LocalRulesClient) listByLocalRulestacksCreateRequest(ctx context.Context, resourceGroupName string, localRulestackName string, options *LocalRulesClientListByLocalRulestacksOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/{localRulestackName}/localRules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if localRulestackName == "" {
		return nil, errors.New("parameter localRulestackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{localRulestackName}", url.PathEscape(localRulestackName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByLocalRulestacksHandleResponse handles the ListByLocalRulestacks response.
func (client *LocalRulesClient) listByLocalRulestacksHandleResponse(resp *http.Response) (LocalRulesClientListByLocalRulestacksResponse, error) {
	result := LocalRulesClientListByLocalRulestacksResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LocalRulesResourceListResult); err != nil {
		return LocalRulesClientListByLocalRulestacksResponse{}, err
	}
	return result, nil
}

// RefreshCounters - Refresh counters
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - localRulestackName - LocalRulestack resource name
//   - priority - Local Rule priority
//   - options - LocalRulesClientRefreshCountersOptions contains the optional parameters for the LocalRulesClient.RefreshCounters
//     method.
func (client *LocalRulesClient) RefreshCounters(ctx context.Context, resourceGroupName string, localRulestackName string, priority string, options *LocalRulesClientRefreshCountersOptions) (LocalRulesClientRefreshCountersResponse, error) {
	var err error
	const operationName = "LocalRulesClient.RefreshCounters"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.refreshCountersCreateRequest(ctx, resourceGroupName, localRulestackName, priority, options)
	if err != nil {
		return LocalRulesClientRefreshCountersResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LocalRulesClientRefreshCountersResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return LocalRulesClientRefreshCountersResponse{}, err
	}
	return LocalRulesClientRefreshCountersResponse{}, nil
}

// refreshCountersCreateRequest creates the RefreshCounters request.
func (client *LocalRulesClient) refreshCountersCreateRequest(ctx context.Context, resourceGroupName string, localRulestackName string, priority string, options *LocalRulesClientRefreshCountersOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/{localRulestackName}/localRules/{priority}/refreshCounters"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if localRulestackName == "" {
		return nil, errors.New("parameter localRulestackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{localRulestackName}", url.PathEscape(localRulestackName))
	if priority == "" {
		return nil, errors.New("parameter priority cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{priority}", url.PathEscape(priority))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	if options != nil && options.FirewallName != nil {
		reqQP.Set("firewallName", *options.FirewallName)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// ResetCounters - Reset counters
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - localRulestackName - LocalRulestack resource name
//   - priority - Local Rule priority
//   - options - LocalRulesClientResetCountersOptions contains the optional parameters for the LocalRulesClient.ResetCounters
//     method.
func (client *LocalRulesClient) ResetCounters(ctx context.Context, resourceGroupName string, localRulestackName string, priority string, options *LocalRulesClientResetCountersOptions) (LocalRulesClientResetCountersResponse, error) {
	var err error
	const operationName = "LocalRulesClient.ResetCounters"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.resetCountersCreateRequest(ctx, resourceGroupName, localRulestackName, priority, options)
	if err != nil {
		return LocalRulesClientResetCountersResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LocalRulesClientResetCountersResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LocalRulesClientResetCountersResponse{}, err
	}
	resp, err := client.resetCountersHandleResponse(httpResp)
	return resp, err
}

// resetCountersCreateRequest creates the ResetCounters request.
func (client *LocalRulesClient) resetCountersCreateRequest(ctx context.Context, resourceGroupName string, localRulestackName string, priority string, options *LocalRulesClientResetCountersOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/{localRulestackName}/localRules/{priority}/resetCounters"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if localRulestackName == "" {
		return nil, errors.New("parameter localRulestackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{localRulestackName}", url.PathEscape(localRulestackName))
	if priority == "" {
		return nil, errors.New("parameter priority cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{priority}", url.PathEscape(priority))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	if options != nil && options.FirewallName != nil {
		reqQP.Set("firewallName", *options.FirewallName)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// resetCountersHandleResponse handles the ResetCounters response.
func (client *LocalRulesClient) resetCountersHandleResponse(resp *http.Response) (LocalRulesClientResetCountersResponse, error) {
	result := LocalRulesClientResetCountersResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RuleCounterReset); err != nil {
		return LocalRulesClientResetCountersResponse{}, err
	}
	return result, nil
}
