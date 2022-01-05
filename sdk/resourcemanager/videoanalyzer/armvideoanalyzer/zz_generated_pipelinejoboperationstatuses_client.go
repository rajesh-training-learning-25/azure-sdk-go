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

// PipelineJobOperationStatusesClient contains the methods for the PipelineJobOperationStatuses group.
// Don't use this type directly, use NewPipelineJobOperationStatusesClient() instead.
type PipelineJobOperationStatusesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewPipelineJobOperationStatusesClient creates a new instance of PipelineJobOperationStatusesClient with the specified values.
func NewPipelineJobOperationStatusesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *PipelineJobOperationStatusesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &PipelineJobOperationStatusesClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Get the operation status of a pipeline job with the given operationId.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PipelineJobOperationStatusesClient) Get(ctx context.Context, resourceGroupName string, accountName string, pipelineJobName string, operationID string, options *PipelineJobOperationStatusesGetOptions) (PipelineJobOperationStatusesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, pipelineJobName, operationID, options)
	if err != nil {
		return PipelineJobOperationStatusesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PipelineJobOperationStatusesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PipelineJobOperationStatusesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PipelineJobOperationStatusesClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, pipelineJobName string, operationID string, options *PipelineJobOperationStatusesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/videoAnalyzers/{accountName}/pipelineJobs/{pipelineJobName}/operationStatuses/{operationId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if pipelineJobName == "" {
		return nil, errors.New("parameter pipelineJobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineJobName}", url.PathEscape(pipelineJobName))
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
func (client *PipelineJobOperationStatusesClient) getHandleResponse(resp *http.Response) (PipelineJobOperationStatusesGetResponse, error) {
	result := PipelineJobOperationStatusesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PipelineJobOperationStatus); err != nil {
		return PipelineJobOperationStatusesGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *PipelineJobOperationStatusesClient) getHandleError(resp *http.Response) error {
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
