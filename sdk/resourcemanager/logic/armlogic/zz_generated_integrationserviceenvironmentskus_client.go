//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogic

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

// IntegrationServiceEnvironmentSKUsClient contains the methods for the IntegrationServiceEnvironmentSKUs group.
// Don't use this type directly, use NewIntegrationServiceEnvironmentSKUsClient() instead.
type IntegrationServiceEnvironmentSKUsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewIntegrationServiceEnvironmentSKUsClient creates a new instance of IntegrationServiceEnvironmentSKUsClient with the specified values.
func NewIntegrationServiceEnvironmentSKUsClient(con *arm.Connection, subscriptionID string) *IntegrationServiceEnvironmentSKUsClient {
	return &IntegrationServiceEnvironmentSKUsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// List - Gets a list of integration service environment Skus.
// If the operation fails it returns the *ErrorResponse error type.
func (client *IntegrationServiceEnvironmentSKUsClient) List(resourceGroup string, integrationServiceEnvironmentName string, options *IntegrationServiceEnvironmentSKUsListOptions) *IntegrationServiceEnvironmentSKUsListPager {
	return &IntegrationServiceEnvironmentSKUsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroup, integrationServiceEnvironmentName, options)
		},
		advancer: func(ctx context.Context, resp IntegrationServiceEnvironmentSKUsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.IntegrationServiceEnvironmentSKUList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *IntegrationServiceEnvironmentSKUsClient) listCreateRequest(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, options *IntegrationServiceEnvironmentSKUsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/skus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if integrationServiceEnvironmentName == "" {
		return nil, errors.New("parameter integrationServiceEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationServiceEnvironmentName}", url.PathEscape(integrationServiceEnvironmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IntegrationServiceEnvironmentSKUsClient) listHandleResponse(resp *http.Response) (IntegrationServiceEnvironmentSKUsListResponse, error) {
	result := IntegrationServiceEnvironmentSKUsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationServiceEnvironmentSKUList); err != nil {
		return IntegrationServiceEnvironmentSKUsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *IntegrationServiceEnvironmentSKUsClient) listHandleError(resp *http.Response) error {
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
