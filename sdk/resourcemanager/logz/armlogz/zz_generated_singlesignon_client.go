//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogz

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

// SingleSignOnClient contains the methods for the SingleSignOn group.
// Don't use this type directly, use NewSingleSignOnClient() instead.
type SingleSignOnClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSingleSignOnClient creates a new instance of SingleSignOnClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSingleSignOnClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SingleSignOnClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &SingleSignOnClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Configures single-sign-on for this resource. This operation can take upto 10 minutes to complete.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// monitorName - Monitor resource name
// options - SingleSignOnClientBeginCreateOrUpdateOptions contains the optional parameters for the SingleSignOnClient.BeginCreateOrUpdate
// method.
func (client *SingleSignOnClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, monitorName string, configurationName string, options *SingleSignOnClientBeginCreateOrUpdateOptions) (SingleSignOnClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, monitorName, configurationName, options)
	if err != nil {
		return SingleSignOnClientCreateOrUpdatePollerResponse{}, err
	}
	result := SingleSignOnClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SingleSignOnClient.CreateOrUpdate", "azure-async-operation", resp, client.pl)
	if err != nil {
		return SingleSignOnClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &SingleSignOnClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Configures single-sign-on for this resource. This operation can take upto 10 minutes to complete.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SingleSignOnClient) createOrUpdate(ctx context.Context, resourceGroupName string, monitorName string, configurationName string, options *SingleSignOnClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, monitorName, configurationName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SingleSignOnClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, configurationName string, options *SingleSignOnClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logz/monitors/{monitorName}/singleSignOnConfigurations/{configurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Body != nil {
		return req, runtime.MarshalAsJSON(req, *options.Body)
	}
	return req, nil
}

// Get - Gets the Logz single sign-on resource for the given Monitor.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// monitorName - Monitor resource name
// options - SingleSignOnClientGetOptions contains the optional parameters for the SingleSignOnClient.Get method.
func (client *SingleSignOnClient) Get(ctx context.Context, resourceGroupName string, monitorName string, configurationName string, options *SingleSignOnClientGetOptions) (SingleSignOnClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, monitorName, configurationName, options)
	if err != nil {
		return SingleSignOnClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SingleSignOnClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SingleSignOnClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SingleSignOnClient) getCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, configurationName string, options *SingleSignOnClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logz/monitors/{monitorName}/singleSignOnConfigurations/{configurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SingleSignOnClient) getHandleResponse(resp *http.Response) (SingleSignOnClientGetResponse, error) {
	result := SingleSignOnClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SingleSignOnResource); err != nil {
		return SingleSignOnClientGetResponse{}, err
	}
	return result, nil
}

// List - List the single sign-on configurations for a given monitor resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// monitorName - Monitor resource name
// options - SingleSignOnClientListOptions contains the optional parameters for the SingleSignOnClient.List method.
func (client *SingleSignOnClient) List(resourceGroupName string, monitorName string, options *SingleSignOnClientListOptions) *SingleSignOnClientListPager {
	return &SingleSignOnClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, monitorName, options)
		},
		advancer: func(ctx context.Context, resp SingleSignOnClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SingleSignOnResourceListResponse.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *SingleSignOnClient) listCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *SingleSignOnClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logz/monitors/{monitorName}/singleSignOnConfigurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SingleSignOnClient) listHandleResponse(resp *http.Response) (SingleSignOnClientListResponse, error) {
	result := SingleSignOnClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SingleSignOnResourceListResponse); err != nil {
		return SingleSignOnClientListResponse{}, err
	}
	return result, nil
}
