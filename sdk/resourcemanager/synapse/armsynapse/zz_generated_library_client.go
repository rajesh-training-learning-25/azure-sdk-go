//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse

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

// LibraryClient contains the methods for the Library group.
// Don't use this type directly, use NewLibraryClient() instead.
type LibraryClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewLibraryClient creates a new instance of LibraryClient with the specified values.
func NewLibraryClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *LibraryClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &LibraryClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Get library by name in a workspace.
// If the operation fails it returns the *ErrorResponse error type.
func (client *LibraryClient) Get(ctx context.Context, resourceGroupName string, libraryName string, workspaceName string, options *LibraryGetOptions) (LibraryGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, libraryName, workspaceName, options)
	if err != nil {
		return LibraryGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LibraryGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LibraryGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LibraryClient) getCreateRequest(ctx context.Context, resourceGroupName string, libraryName string, workspaceName string, options *LibraryGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/libraries/{libraryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if libraryName == "" {
		return nil, errors.New("parameter libraryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{libraryName}", url.PathEscape(libraryName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LibraryClient) getHandleResponse(resp *http.Response) (LibraryGetResponse, error) {
	result := LibraryGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.LibraryResource); err != nil {
		return LibraryGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *LibraryClient) getHandleError(resp *http.Response) error {
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
