//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsupport

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ServicesClient contains the methods for the Services group.
// Don't use this type directly, use NewServicesClient() instead.
type ServicesClient struct {
	ep string
	pl runtime.Pipeline
}

// NewServicesClient creates a new instance of ServicesClient with the specified values.
func NewServicesClient(credential azcore.TokenCredential, options *arm.ClientOptions) *ServicesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ServicesClient{ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Gets a specific Azure service for support ticket creation.
// If the operation fails it returns the *ExceptionResponse error type.
func (client *ServicesClient) Get(ctx context.Context, serviceName string, options *ServicesGetOptions) (ServicesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, serviceName, options)
	if err != nil {
		return ServicesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServicesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServicesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServicesClient) getCreateRequest(ctx context.Context, serviceName string, options *ServicesGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Support/services/{serviceName}"
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServicesClient) getHandleResponse(resp *http.Response) (ServicesGetResponse, error) {
	result := ServicesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Service); err != nil {
		return ServicesGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ServicesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ExceptionResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Lists all the Azure services available for support ticket creation. For Technical issues, select the Service Id that maps to the Azure service/product
// as displayed in the Services drop-down list on
// the Azure portal's New support request [https://portal.azure.com/#blade/MicrosoftAzureSupport/HelpAndSupportBlade/overview] page. Always use the service
// and its corresponding problem classification(s)
// obtained programmatically for support ticket creation. This practice ensures that you always have the most recent set of service and problem classification
// Ids.
// If the operation fails it returns the *ExceptionResponse error type.
func (client *ServicesClient) List(ctx context.Context, options *ServicesListOptions) (ServicesListResponse, error) {
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return ServicesListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServicesListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServicesListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ServicesClient) listCreateRequest(ctx context.Context, options *ServicesListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Support/services"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ServicesClient) listHandleResponse(resp *http.Response) (ServicesListResponse, error) {
	result := ServicesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServicesListResult); err != nil {
		return ServicesListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ServicesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ExceptionResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
