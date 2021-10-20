//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SecuritySolutionsClient contains the methods for the SecuritySolutions group.
// Don't use this type directly, use NewSecuritySolutionsClient() instead.
type SecuritySolutionsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
	ascLocation    string
}

// NewSecuritySolutionsClient creates a new instance of SecuritySolutionsClient with the specified values.
func NewSecuritySolutionsClient(con *arm.Connection, subscriptionID string, ascLocation string) *SecuritySolutionsClient {
	return &SecuritySolutionsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID, ascLocation: ascLocation}
}

// Get - Gets a specific Security Solution.
// If the operation fails it returns the *CloudError error type.
func (client *SecuritySolutionsClient) Get(ctx context.Context, resourceGroupName string, securitySolutionName string, options *SecuritySolutionsGetOptions) (SecuritySolutionsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, securitySolutionName, options)
	if err != nil {
		return SecuritySolutionsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SecuritySolutionsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SecuritySolutionsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SecuritySolutionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, securitySolutionName string, options *SecuritySolutionsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/locations/{ascLocation}/securitySolutions/{securitySolutionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.ascLocation == "" {
		return nil, errors.New("parameter client.ascLocation cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ascLocation}", url.PathEscape(client.ascLocation))
	if securitySolutionName == "" {
		return nil, errors.New("parameter securitySolutionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securitySolutionName}", url.PathEscape(securitySolutionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SecuritySolutionsClient) getHandleResponse(resp *http.Response) (SecuritySolutionsGetResponse, error) {
	result := SecuritySolutionsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecuritySolution); err != nil {
		return SecuritySolutionsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *SecuritySolutionsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Gets a list of Security Solutions for the subscription.
// If the operation fails it returns the *CloudError error type.
func (client *SecuritySolutionsClient) List(options *SecuritySolutionsListOptions) *SecuritySolutionsListPager {
	return &SecuritySolutionsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp SecuritySolutionsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SecuritySolutionList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *SecuritySolutionsClient) listCreateRequest(ctx context.Context, options *SecuritySolutionsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/securitySolutions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SecuritySolutionsClient) listHandleResponse(resp *http.Response) (SecuritySolutionsListResponse, error) {
	result := SecuritySolutionsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecuritySolutionList); err != nil {
		return SecuritySolutionsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *SecuritySolutionsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
