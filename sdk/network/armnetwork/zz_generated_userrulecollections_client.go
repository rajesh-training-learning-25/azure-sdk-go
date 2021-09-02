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

// UserRuleCollectionsClient contains the methods for the UserRuleCollections group.
// Don't use this type directly, use NewUserRuleCollectionsClient() instead.
type UserRuleCollectionsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewUserRuleCollectionsClient creates a new instance of UserRuleCollectionsClient with the specified values.
func NewUserRuleCollectionsClient(con *arm.Connection, subscriptionID string) *UserRuleCollectionsClient {
	return &UserRuleCollectionsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// CreateOrUpdate - Creates or updates a user rule collection.
// If the operation fails it returns the *CloudErrorAutoGenerated error type.
func (client *UserRuleCollectionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, ruleCollectionName string, userRuleCollection RuleCollection, options *UserRuleCollectionsCreateOrUpdateOptions) (UserRuleCollectionsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, networkManagerName, configurationName, ruleCollectionName, userRuleCollection, options)
	if err != nil {
		return UserRuleCollectionsCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return UserRuleCollectionsCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return UserRuleCollectionsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *UserRuleCollectionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, ruleCollectionName string, userRuleCollection RuleCollection, options *UserRuleCollectionsCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/securityUserConfigurations/{configurationName}/ruleCollections/{ruleCollectionName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, userRuleCollection)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *UserRuleCollectionsClient) createOrUpdateHandleResponse(resp *http.Response) (UserRuleCollectionsCreateOrUpdateResponse, error) {
	result := UserRuleCollectionsCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RuleCollection); err != nil {
		return UserRuleCollectionsCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *UserRuleCollectionsClient) createOrUpdateHandleError(resp *http.Response) error {
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

// Delete - Deletes a user rule collection.
// If the operation fails it returns the *CloudErrorAutoGenerated error type.
func (client *UserRuleCollectionsClient) Delete(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, ruleCollectionName string, options *UserRuleCollectionsDeleteOptions) (UserRuleCollectionsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkManagerName, configurationName, ruleCollectionName, options)
	if err != nil {
		return UserRuleCollectionsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return UserRuleCollectionsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return UserRuleCollectionsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return UserRuleCollectionsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *UserRuleCollectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, ruleCollectionName string, options *UserRuleCollectionsDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/securityUserConfigurations/{configurationName}/ruleCollections/{ruleCollectionName}"
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
func (client *UserRuleCollectionsClient) deleteHandleError(resp *http.Response) error {
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

// Get - Gets a network manager security user configuration rule collection.
// If the operation fails it returns the *CloudErrorAutoGenerated error type.
func (client *UserRuleCollectionsClient) Get(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, ruleCollectionName string, options *UserRuleCollectionsGetOptions) (UserRuleCollectionsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkManagerName, configurationName, ruleCollectionName, options)
	if err != nil {
		return UserRuleCollectionsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return UserRuleCollectionsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return UserRuleCollectionsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *UserRuleCollectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, ruleCollectionName string, options *UserRuleCollectionsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/securityUserConfigurations/{configurationName}/ruleCollections/{ruleCollectionName}"
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
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *UserRuleCollectionsClient) getHandleResponse(resp *http.Response) (UserRuleCollectionsGetResponse, error) {
	result := UserRuleCollectionsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RuleCollection); err != nil {
		return UserRuleCollectionsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *UserRuleCollectionsClient) getHandleError(resp *http.Response) error {
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

// List - Lists all the user rule collections in a security configuration, in a paginated format.
// If the operation fails it returns the *CloudErrorAutoGenerated error type.
func (client *UserRuleCollectionsClient) List(resourceGroupName string, networkManagerName string, configurationName string, options *UserRuleCollectionsListOptions) *UserRuleCollectionsListPager {
	return &UserRuleCollectionsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, networkManagerName, configurationName, options)
		},
		advancer: func(ctx context.Context, resp UserRuleCollectionsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RuleCollectionListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *UserRuleCollectionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, options *UserRuleCollectionsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/securityUserConfigurations/{configurationName}/ruleCollections"
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
func (client *UserRuleCollectionsClient) listHandleResponse(resp *http.Response) (UserRuleCollectionsListResponse, error) {
	result := UserRuleCollectionsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RuleCollectionListResult); err != nil {
		return UserRuleCollectionsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *UserRuleCollectionsClient) listHandleError(resp *http.Response) error {
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
