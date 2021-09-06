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
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// AzureFirewallFqdnTagsClient contains the methods for the AzureFirewallFqdnTags group.
// Don't use this type directly, use NewAzureFirewallFqdnTagsClient() instead.
type AzureFirewallFqdnTagsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewAzureFirewallFqdnTagsClient creates a new instance of AzureFirewallFqdnTagsClient with the specified values.
func NewAzureFirewallFqdnTagsClient(con *arm.Connection, subscriptionID string) *AzureFirewallFqdnTagsClient {
	return &AzureFirewallFqdnTagsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// ListAll - Gets all the Azure Firewall FQDN Tags in a subscription.
// If the operation fails it returns the *CloudError error type.
func (client *AzureFirewallFqdnTagsClient) ListAll(options *AzureFirewallFqdnTagsListAllOptions) *AzureFirewallFqdnTagsListAllPager {
	return &AzureFirewallFqdnTagsListAllPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listAllCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp AzureFirewallFqdnTagsListAllResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AzureFirewallFqdnTagListResult.NextLink)
		},
	}
}

// listAllCreateRequest creates the ListAll request.
func (client *AzureFirewallFqdnTagsClient) listAllCreateRequest(ctx context.Context, options *AzureFirewallFqdnTagsListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/azureFirewallFqdnTags"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *AzureFirewallFqdnTagsClient) listAllHandleResponse(resp *http.Response) (AzureFirewallFqdnTagsListAllResponse, error) {
	result := AzureFirewallFqdnTagsListAllResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureFirewallFqdnTagListResult); err != nil {
		return AzureFirewallFqdnTagsListAllResponse{}, err
	}
	return result, nil
}

// listAllHandleError handles the ListAll error response.
func (client *AzureFirewallFqdnTagsClient) listAllHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
