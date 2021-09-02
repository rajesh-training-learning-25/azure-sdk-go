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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ActiveSecurityAdminRulesClient contains the methods for the ActiveSecurityAdminRules group.
// Don't use this type directly, use NewActiveSecurityAdminRulesClient() instead.
type ActiveSecurityAdminRulesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewActiveSecurityAdminRulesClient creates a new instance of ActiveSecurityAdminRulesClient with the specified values.
func NewActiveSecurityAdminRulesClient(con *arm.Connection, subscriptionID string) *ActiveSecurityAdminRulesClient {
	return &ActiveSecurityAdminRulesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// List - Lists active security admin rules in a network manager.
// If the operation fails it returns the *CloudErrorAutoGenerated error type.
func (client *ActiveSecurityAdminRulesClient) List(ctx context.Context, resourceGroupName string, networkManagerName string, parameters interface{}, options *ActiveSecurityAdminRulesListOptions) (ActiveSecurityAdminRulesListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, networkManagerName, parameters, options)
	if err != nil {
		return ActiveSecurityAdminRulesListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ActiveSecurityAdminRulesListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ActiveSecurityAdminRulesListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ActiveSecurityAdminRulesClient) listCreateRequest(ctx context.Context, resourceGroupName string, networkManagerName string, parameters interface{}, options *ActiveSecurityAdminRulesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/listActiveSecurityAdminRules"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ActiveSecurityAdminRulesClient) listHandleResponse(resp *http.Response) (ActiveSecurityAdminRulesListResponse, error) {
	result := ActiveSecurityAdminRulesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ActiveSecurityAdminRulesListResult); err != nil {
		return ActiveSecurityAdminRulesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ActiveSecurityAdminRulesClient) listHandleError(resp *http.Response) error {
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
