//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnewrelicobservability

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

// AccountsClient contains the methods for the Accounts group.
// Don't use this type directly, use NewAccountsClient() instead.
type AccountsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAccountsClient creates a new instance of AccountsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAccountsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AccountsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &AccountsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListPager - List all the existing accounts
//
// Generated from API version 2022-07-01-preview
//   - userEmail - User Email.
//   - location - Location for NewRelic.
//   - options - AccountsClientListOptions contains the optional parameters for the AccountsClient.NewListPager method.
func (client *AccountsClient) NewListPager(userEmail string, location string, options *AccountsClientListOptions) *runtime.Pager[AccountsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AccountsClientListResponse]{
		More: func(page AccountsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AccountsClientListResponse) (AccountsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, userEmail, location, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AccountsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AccountsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AccountsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *AccountsClient) listCreateRequest(ctx context.Context, userEmail string, location string, options *AccountsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/NewRelic.Observability/accounts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	reqQP.Set("userEmail", userEmail)
	reqQP.Set("location", location)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AccountsClient) listHandleResponse(resp *http.Response) (AccountsClientListResponse, error) {
	result := AccountsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccountsListResponse); err != nil {
		return AccountsClientListResponse{}, err
	}
	return result, nil
}
