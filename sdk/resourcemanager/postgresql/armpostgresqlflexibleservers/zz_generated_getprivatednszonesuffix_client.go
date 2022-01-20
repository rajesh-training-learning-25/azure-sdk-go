//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpostgresqlflexibleservers

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// GetPrivateDNSZoneSuffixClient contains the methods for the GetPrivateDNSZoneSuffix group.
// Don't use this type directly, use NewGetPrivateDNSZoneSuffixClient() instead.
type GetPrivateDNSZoneSuffixClient struct {
	host string
	pl   runtime.Pipeline
}

// NewGetPrivateDNSZoneSuffixClient creates a new instance of GetPrivateDNSZoneSuffixClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewGetPrivateDNSZoneSuffixClient(credential azcore.TokenCredential, options *arm.ClientOptions) *GetPrivateDNSZoneSuffixClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &GetPrivateDNSZoneSuffixClient{
		host: string(cp.Endpoint),
		pl:   armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Execute - Get private DNS zone suffix in the cloud
// If the operation fails it returns an *azcore.ResponseError type.
// options - GetPrivateDNSZoneSuffixClientExecuteOptions contains the optional parameters for the GetPrivateDNSZoneSuffixClient.Execute
// method.
func (client *GetPrivateDNSZoneSuffixClient) Execute(ctx context.Context, options *GetPrivateDNSZoneSuffixClientExecuteOptions) (GetPrivateDNSZoneSuffixClientExecuteResponse, error) {
	req, err := client.executeCreateRequest(ctx, options)
	if err != nil {
		return GetPrivateDNSZoneSuffixClientExecuteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GetPrivateDNSZoneSuffixClientExecuteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GetPrivateDNSZoneSuffixClientExecuteResponse{}, runtime.NewResponseError(resp)
	}
	return client.executeHandleResponse(resp)
}

// executeCreateRequest creates the Execute request.
func (client *GetPrivateDNSZoneSuffixClient) executeCreateRequest(ctx context.Context, options *GetPrivateDNSZoneSuffixClientExecuteOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.DBforPostgreSQL/getPrivateDnsZoneSuffix"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// executeHandleResponse handles the Execute response.
func (client *GetPrivateDNSZoneSuffixClient) executeHandleResponse(resp *http.Response) (GetPrivateDNSZoneSuffixClientExecuteResponse, error) {
	result := GetPrivateDNSZoneSuffixClientExecuteResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return GetPrivateDNSZoneSuffixClientExecuteResponse{}, err
	}
	return result, nil
}
