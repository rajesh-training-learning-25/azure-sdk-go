//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtrafficmanager

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// GeographicHierarchiesClient contains the methods for the GeographicHierarchies group.
// Don't use this type directly, use NewGeographicHierarchiesClient() instead.
type GeographicHierarchiesClient struct {
	internal *arm.Client
}

// NewGeographicHierarchiesClient creates a new instance of GeographicHierarchiesClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewGeographicHierarchiesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*GeographicHierarchiesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &GeographicHierarchiesClient{
		internal: cl,
	}
	return client, nil
}

// GetDefault - Gets the default Geographic Hierarchy used by the Geographic traffic routing method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-04-01
//   - options - GeographicHierarchiesClientGetDefaultOptions contains the optional parameters for the GeographicHierarchiesClient.GetDefault
//     method.
func (client *GeographicHierarchiesClient) GetDefault(ctx context.Context, options *GeographicHierarchiesClientGetDefaultOptions) (GeographicHierarchiesClientGetDefaultResponse, error) {
	var err error
	const operationName = "GeographicHierarchiesClient.GetDefault"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getDefaultCreateRequest(ctx, options)
	if err != nil {
		return GeographicHierarchiesClientGetDefaultResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GeographicHierarchiesClientGetDefaultResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return GeographicHierarchiesClientGetDefaultResponse{}, err
	}
	resp, err := client.getDefaultHandleResponse(httpResp)
	return resp, err
}

// getDefaultCreateRequest creates the GetDefault request.
func (client *GeographicHierarchiesClient) getDefaultCreateRequest(ctx context.Context, options *GeographicHierarchiesClientGetDefaultOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Network/trafficManagerGeographicHierarchies/default"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDefaultHandleResponse handles the GetDefault response.
func (client *GeographicHierarchiesClient) getDefaultHandleResponse(resp *http.Response) (GeographicHierarchiesClientGetDefaultResponse, error) {
	result := GeographicHierarchiesClientGetDefaultResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GeographicHierarchy); err != nil {
		return GeographicHierarchiesClientGetDefaultResponse{}, err
	}
	return result, nil
}
