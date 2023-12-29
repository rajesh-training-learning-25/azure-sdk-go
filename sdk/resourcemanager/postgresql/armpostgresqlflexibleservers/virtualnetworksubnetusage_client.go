//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpostgresqlflexibleservers

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// VirtualNetworkSubnetUsageClient contains the methods for the VirtualNetworkSubnetUsage group.
// Don't use this type directly, use NewVirtualNetworkSubnetUsageClient() instead.
type VirtualNetworkSubnetUsageClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVirtualNetworkSubnetUsageClient creates a new instance of VirtualNetworkSubnetUsageClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVirtualNetworkSubnetUsageClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualNetworkSubnetUsageClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VirtualNetworkSubnetUsageClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Execute - Get virtual network subnet usage for a given vNet resource id.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - locationName - The name of the location.
//   - parameters - The required parameters for creating or updating a server.
//   - options - VirtualNetworkSubnetUsageClientExecuteOptions contains the optional parameters for the VirtualNetworkSubnetUsageClient.Execute
//     method.
func (client *VirtualNetworkSubnetUsageClient) Execute(ctx context.Context, locationName string, parameters VirtualNetworkSubnetUsageParameter, options *VirtualNetworkSubnetUsageClientExecuteOptions) (VirtualNetworkSubnetUsageClientExecuteResponse, error) {
	var err error
	const operationName = "VirtualNetworkSubnetUsageClient.Execute"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.executeCreateRequest(ctx, locationName, parameters, options)
	if err != nil {
		return VirtualNetworkSubnetUsageClientExecuteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualNetworkSubnetUsageClientExecuteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VirtualNetworkSubnetUsageClientExecuteResponse{}, err
	}
	resp, err := client.executeHandleResponse(httpResp)
	return resp, err
}

// executeCreateRequest creates the Execute request.
func (client *VirtualNetworkSubnetUsageClient) executeCreateRequest(ctx context.Context, locationName string, parameters VirtualNetworkSubnetUsageParameter, options *VirtualNetworkSubnetUsageClientExecuteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DBforPostgreSQL/locations/{locationName}/checkVirtualNetworkSubnetUsage"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// executeHandleResponse handles the Execute response.
func (client *VirtualNetworkSubnetUsageClient) executeHandleResponse(resp *http.Response) (VirtualNetworkSubnetUsageClientExecuteResponse, error) {
	result := VirtualNetworkSubnetUsageClientExecuteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkSubnetUsageResult); err != nil {
		return VirtualNetworkSubnetUsageClientExecuteResponse{}, err
	}
	return result, nil
}
