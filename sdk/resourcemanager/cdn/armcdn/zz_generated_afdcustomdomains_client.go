//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcdn

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

// AFDCustomDomainsClient contains the methods for the AFDCustomDomains group.
// Don't use this type directly, use NewAFDCustomDomainsClient() instead.
type AFDCustomDomainsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewAFDCustomDomainsClient creates a new instance of AFDCustomDomainsClient with the specified values.
func NewAFDCustomDomainsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *AFDCustomDomainsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &AFDCustomDomainsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreate - Creates a new domain within the specified profile.
// If the operation fails it returns the *AfdErrorResponse error type.
func (client *AFDCustomDomainsClient) BeginCreate(ctx context.Context, resourceGroupName string, profileName string, customDomainName string, customDomain AFDDomain, options *AFDCustomDomainsBeginCreateOptions) (AFDCustomDomainsCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, profileName, customDomainName, customDomain, options)
	if err != nil {
		return AFDCustomDomainsCreatePollerResponse{}, err
	}
	result := AFDCustomDomainsCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("AFDCustomDomainsClient.Create", "azure-async-operation", resp, client.pl, client.createHandleError)
	if err != nil {
		return AFDCustomDomainsCreatePollerResponse{}, err
	}
	result.Poller = &AFDCustomDomainsCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Creates a new domain within the specified profile.
// If the operation fails it returns the *AfdErrorResponse error type.
func (client *AFDCustomDomainsClient) create(ctx context.Context, resourceGroupName string, profileName string, customDomainName string, customDomain AFDDomain, options *AFDCustomDomainsBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, profileName, customDomainName, customDomain, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createHandleError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *AFDCustomDomainsClient) createCreateRequest(ctx context.Context, resourceGroupName string, profileName string, customDomainName string, customDomain AFDDomain, options *AFDCustomDomainsBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/customDomains/{customDomainName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if customDomainName == "" {
		return nil, errors.New("parameter customDomainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customDomainName}", url.PathEscape(customDomainName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, customDomain)
}

// createHandleError handles the Create error response.
func (client *AFDCustomDomainsClient) createHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := AfdErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes an existing AzureFrontDoor domain with the specified domain name under the specified subscription, resource group and profile.
// If the operation fails it returns the *AfdErrorResponse error type.
func (client *AFDCustomDomainsClient) BeginDelete(ctx context.Context, resourceGroupName string, profileName string, customDomainName string, options *AFDCustomDomainsBeginDeleteOptions) (AFDCustomDomainsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, profileName, customDomainName, options)
	if err != nil {
		return AFDCustomDomainsDeletePollerResponse{}, err
	}
	result := AFDCustomDomainsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("AFDCustomDomainsClient.Delete", "azure-async-operation", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return AFDCustomDomainsDeletePollerResponse{}, err
	}
	result.Poller = &AFDCustomDomainsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes an existing AzureFrontDoor domain with the specified domain name under the specified subscription, resource group and profile.
// If the operation fails it returns the *AfdErrorResponse error type.
func (client *AFDCustomDomainsClient) deleteOperation(ctx context.Context, resourceGroupName string, profileName string, customDomainName string, options *AFDCustomDomainsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, profileName, customDomainName, options)
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
func (client *AFDCustomDomainsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, profileName string, customDomainName string, options *AFDCustomDomainsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/customDomains/{customDomainName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if customDomainName == "" {
		return nil, errors.New("parameter customDomainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customDomainName}", url.PathEscape(customDomainName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *AFDCustomDomainsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := AfdErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets an existing AzureFrontDoor domain with the specified domain name under the specified subscription, resource group and profile.
// If the operation fails it returns the *AfdErrorResponse error type.
func (client *AFDCustomDomainsClient) Get(ctx context.Context, resourceGroupName string, profileName string, customDomainName string, options *AFDCustomDomainsGetOptions) (AFDCustomDomainsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, profileName, customDomainName, options)
	if err != nil {
		return AFDCustomDomainsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AFDCustomDomainsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AFDCustomDomainsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AFDCustomDomainsClient) getCreateRequest(ctx context.Context, resourceGroupName string, profileName string, customDomainName string, options *AFDCustomDomainsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/customDomains/{customDomainName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if customDomainName == "" {
		return nil, errors.New("parameter customDomainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customDomainName}", url.PathEscape(customDomainName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AFDCustomDomainsClient) getHandleResponse(resp *http.Response) (AFDCustomDomainsGetResponse, error) {
	result := AFDCustomDomainsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AFDDomain); err != nil {
		return AFDCustomDomainsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *AFDCustomDomainsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := AfdErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByProfile - Lists existing AzureFrontDoor domains.
// If the operation fails it returns the *AfdErrorResponse error type.
func (client *AFDCustomDomainsClient) ListByProfile(resourceGroupName string, profileName string, options *AFDCustomDomainsListByProfileOptions) *AFDCustomDomainsListByProfilePager {
	return &AFDCustomDomainsListByProfilePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByProfileCreateRequest(ctx, resourceGroupName, profileName, options)
		},
		advancer: func(ctx context.Context, resp AFDCustomDomainsListByProfileResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AFDDomainListResult.NextLink)
		},
	}
}

// listByProfileCreateRequest creates the ListByProfile request.
func (client *AFDCustomDomainsClient) listByProfileCreateRequest(ctx context.Context, resourceGroupName string, profileName string, options *AFDCustomDomainsListByProfileOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/customDomains"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByProfileHandleResponse handles the ListByProfile response.
func (client *AFDCustomDomainsClient) listByProfileHandleResponse(resp *http.Response) (AFDCustomDomainsListByProfileResponse, error) {
	result := AFDCustomDomainsListByProfileResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AFDDomainListResult); err != nil {
		return AFDCustomDomainsListByProfileResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByProfileHandleError handles the ListByProfile error response.
func (client *AFDCustomDomainsClient) listByProfileHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := AfdErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginRefreshValidationToken - Updates the domain validation token.
// If the operation fails it returns the *AfdErrorResponse error type.
func (client *AFDCustomDomainsClient) BeginRefreshValidationToken(ctx context.Context, resourceGroupName string, profileName string, customDomainName string, options *AFDCustomDomainsBeginRefreshValidationTokenOptions) (AFDCustomDomainsRefreshValidationTokenPollerResponse, error) {
	resp, err := client.refreshValidationToken(ctx, resourceGroupName, profileName, customDomainName, options)
	if err != nil {
		return AFDCustomDomainsRefreshValidationTokenPollerResponse{}, err
	}
	result := AFDCustomDomainsRefreshValidationTokenPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("AFDCustomDomainsClient.RefreshValidationToken", "azure-async-operation", resp, client.pl, client.refreshValidationTokenHandleError)
	if err != nil {
		return AFDCustomDomainsRefreshValidationTokenPollerResponse{}, err
	}
	result.Poller = &AFDCustomDomainsRefreshValidationTokenPoller{
		pt: pt,
	}
	return result, nil
}

// RefreshValidationToken - Updates the domain validation token.
// If the operation fails it returns the *AfdErrorResponse error type.
func (client *AFDCustomDomainsClient) refreshValidationToken(ctx context.Context, resourceGroupName string, profileName string, customDomainName string, options *AFDCustomDomainsBeginRefreshValidationTokenOptions) (*http.Response, error) {
	req, err := client.refreshValidationTokenCreateRequest(ctx, resourceGroupName, profileName, customDomainName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, client.refreshValidationTokenHandleError(resp)
	}
	return resp, nil
}

// refreshValidationTokenCreateRequest creates the RefreshValidationToken request.
func (client *AFDCustomDomainsClient) refreshValidationTokenCreateRequest(ctx context.Context, resourceGroupName string, profileName string, customDomainName string, options *AFDCustomDomainsBeginRefreshValidationTokenOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/customDomains/{customDomainName}/refreshValidationToken"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if customDomainName == "" {
		return nil, errors.New("parameter customDomainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customDomainName}", url.PathEscape(customDomainName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// refreshValidationTokenHandleError handles the RefreshValidationToken error response.
func (client *AFDCustomDomainsClient) refreshValidationTokenHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := AfdErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginUpdate - Updates an existing domain within a profile.
// If the operation fails it returns the *AfdErrorResponse error type.
func (client *AFDCustomDomainsClient) BeginUpdate(ctx context.Context, resourceGroupName string, profileName string, customDomainName string, customDomainUpdateProperties AFDDomainUpdateParameters, options *AFDCustomDomainsBeginUpdateOptions) (AFDCustomDomainsUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, profileName, customDomainName, customDomainUpdateProperties, options)
	if err != nil {
		return AFDCustomDomainsUpdatePollerResponse{}, err
	}
	result := AFDCustomDomainsUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("AFDCustomDomainsClient.Update", "azure-async-operation", resp, client.pl, client.updateHandleError)
	if err != nil {
		return AFDCustomDomainsUpdatePollerResponse{}, err
	}
	result.Poller = &AFDCustomDomainsUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates an existing domain within a profile.
// If the operation fails it returns the *AfdErrorResponse error type.
func (client *AFDCustomDomainsClient) update(ctx context.Context, resourceGroupName string, profileName string, customDomainName string, customDomainUpdateProperties AFDDomainUpdateParameters, options *AFDCustomDomainsBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, profileName, customDomainName, customDomainUpdateProperties, options)
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
func (client *AFDCustomDomainsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, profileName string, customDomainName string, customDomainUpdateProperties AFDDomainUpdateParameters, options *AFDCustomDomainsBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/customDomains/{customDomainName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if customDomainName == "" {
		return nil, errors.New("parameter customDomainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customDomainName}", url.PathEscape(customDomainName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, customDomainUpdateProperties)
}

// updateHandleError handles the Update error response.
func (client *AFDCustomDomainsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := AfdErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
