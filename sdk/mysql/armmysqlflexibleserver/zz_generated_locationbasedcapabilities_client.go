//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmysqlflexibleserver

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

// LocationBasedCapabilitiesClient contains the methods for the LocationBasedCapabilities group.
// Don't use this type directly, use NewLocationBasedCapabilitiesClient() instead.
type LocationBasedCapabilitiesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewLocationBasedCapabilitiesClient creates a new instance of LocationBasedCapabilitiesClient with the specified values.
func NewLocationBasedCapabilitiesClient(con *arm.Connection, subscriptionID string) *LocationBasedCapabilitiesClient {
	return &LocationBasedCapabilitiesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// List - Get capabilities at specified location in a given subscription.
// If the operation fails it returns the *CloudError error type.
func (client *LocationBasedCapabilitiesClient) List(locationName string, options *LocationBasedCapabilitiesListOptions) *LocationBasedCapabilitiesListPager {
	return &LocationBasedCapabilitiesListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, locationName, options)
		},
		advancer: func(ctx context.Context, resp LocationBasedCapabilitiesListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.CapabilitiesListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *LocationBasedCapabilitiesClient) listCreateRequest(ctx context.Context, locationName string, options *LocationBasedCapabilitiesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DBforMySQL/locations/{locationName}/capabilities"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LocationBasedCapabilitiesClient) listHandleResponse(resp *http.Response) (LocationBasedCapabilitiesListResponse, error) {
	result := LocationBasedCapabilitiesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CapabilitiesListResult); err != nil {
		return LocationBasedCapabilitiesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *LocationBasedCapabilitiesClient) listHandleError(resp *http.Response) error {
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
