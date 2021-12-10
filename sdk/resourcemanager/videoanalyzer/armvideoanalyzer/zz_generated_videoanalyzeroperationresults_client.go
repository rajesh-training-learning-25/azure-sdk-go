//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvideoanalyzer

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

// VideoAnalyzerOperationResultsClient contains the methods for the VideoAnalyzerOperationResults group.
// Don't use this type directly, use NewVideoAnalyzerOperationResultsClient() instead.
type VideoAnalyzerOperationResultsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewVideoAnalyzerOperationResultsClient creates a new instance of VideoAnalyzerOperationResultsClient with the specified values.
func NewVideoAnalyzerOperationResultsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *VideoAnalyzerOperationResultsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &VideoAnalyzerOperationResultsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Get video analyzer operation result.
// If the operation fails it returns the *ErrorResponse error type.
func (client *VideoAnalyzerOperationResultsClient) Get(ctx context.Context, locationName string, operationID string, options *VideoAnalyzerOperationResultsGetOptions) (VideoAnalyzerOperationResultsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, locationName, operationID, options)
	if err != nil {
		return VideoAnalyzerOperationResultsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VideoAnalyzerOperationResultsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return VideoAnalyzerOperationResultsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VideoAnalyzerOperationResultsClient) getCreateRequest(ctx context.Context, locationName string, operationID string, options *VideoAnalyzerOperationResultsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Media/locations/{locationName}/videoAnalyzerOperationResults/{operationId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VideoAnalyzerOperationResultsClient) getHandleResponse(resp *http.Response) (VideoAnalyzerOperationResultsGetResponse, error) {
	result := VideoAnalyzerOperationResultsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VideoAnalyzer); err != nil {
		return VideoAnalyzerOperationResultsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *VideoAnalyzerOperationResultsClient) getHandleError(resp *http.Response) error {
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
