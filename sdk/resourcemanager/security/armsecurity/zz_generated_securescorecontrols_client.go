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

// SecureScoreControlsClient contains the methods for the SecureScoreControls group.
// Don't use this type directly, use NewSecureScoreControlsClient() instead.
type SecureScoreControlsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewSecureScoreControlsClient creates a new instance of SecureScoreControlsClient with the specified values.
func NewSecureScoreControlsClient(con *arm.Connection, subscriptionID string) *SecureScoreControlsClient {
	return &SecureScoreControlsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// List - Get all security controls within a scope
// If the operation fails it returns the *CloudError error type.
func (client *SecureScoreControlsClient) List(options *SecureScoreControlsListOptions) *SecureScoreControlsListPager {
	return &SecureScoreControlsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp SecureScoreControlsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SecureScoreControlList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *SecureScoreControlsClient) listCreateRequest(ctx context.Context, options *SecureScoreControlsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/secureScoreControls"
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
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SecureScoreControlsClient) listHandleResponse(resp *http.Response) (SecureScoreControlsListResponse, error) {
	result := SecureScoreControlsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecureScoreControlList); err != nil {
		return SecureScoreControlsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *SecureScoreControlsClient) listHandleError(resp *http.Response) error {
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

// ListBySecureScore - Get all security controls for a specific initiative within a scope
// If the operation fails it returns the *CloudError error type.
func (client *SecureScoreControlsClient) ListBySecureScore(secureScoreName string, options *SecureScoreControlsListBySecureScoreOptions) *SecureScoreControlsListBySecureScorePager {
	return &SecureScoreControlsListBySecureScorePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySecureScoreCreateRequest(ctx, secureScoreName, options)
		},
		advancer: func(ctx context.Context, resp SecureScoreControlsListBySecureScoreResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SecureScoreControlList.NextLink)
		},
	}
}

// listBySecureScoreCreateRequest creates the ListBySecureScore request.
func (client *SecureScoreControlsClient) listBySecureScoreCreateRequest(ctx context.Context, secureScoreName string, options *SecureScoreControlsListBySecureScoreOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/secureScores/{secureScoreName}/secureScoreControls"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if secureScoreName == "" {
		return nil, errors.New("parameter secureScoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{secureScoreName}", url.PathEscape(secureScoreName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySecureScoreHandleResponse handles the ListBySecureScore response.
func (client *SecureScoreControlsClient) listBySecureScoreHandleResponse(resp *http.Response) (SecureScoreControlsListBySecureScoreResponse, error) {
	result := SecureScoreControlsListBySecureScoreResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecureScoreControlList); err != nil {
		return SecureScoreControlsListBySecureScoreResponse{}, err
	}
	return result, nil
}

// listBySecureScoreHandleError handles the ListBySecureScore error response.
func (client *SecureScoreControlsClient) listBySecureScoreHandleError(resp *http.Response) error {
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
