//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicesbackup

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// BMSPrepareDataMoveOperationResultClient contains the methods for the BMSPrepareDataMoveOperationResult group.
// Don't use this type directly, use NewBMSPrepareDataMoveOperationResultClient() instead.
type BMSPrepareDataMoveOperationResultClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewBMSPrepareDataMoveOperationResultClient creates a new instance of BMSPrepareDataMoveOperationResultClient with the specified values.
// subscriptionID - The subscription Id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewBMSPrepareDataMoveOperationResultClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *BMSPrepareDataMoveOperationResultClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &BMSPrepareDataMoveOperationResultClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Get - Fetches Operation Result for Prepare Data Move
// If the operation fails it returns an *azcore.ResponseError type.
// vaultName - The name of the recovery services vault.
// resourceGroupName - The name of the resource group where the recovery services vault is present.
// options - BMSPrepareDataMoveOperationResultClientGetOptions contains the optional parameters for the BMSPrepareDataMoveOperationResultClient.Get
// method.
func (client *BMSPrepareDataMoveOperationResultClient) Get(ctx context.Context, vaultName string, resourceGroupName string, operationID string, options *BMSPrepareDataMoveOperationResultClientGetOptions) (BMSPrepareDataMoveOperationResultClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, vaultName, resourceGroupName, operationID, options)
	if err != nil {
		return BMSPrepareDataMoveOperationResultClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BMSPrepareDataMoveOperationResultClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return BMSPrepareDataMoveOperationResultClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *BMSPrepareDataMoveOperationResultClient) getCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, operationID string, options *BMSPrepareDataMoveOperationResultClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupstorageconfig/vaultstorageconfig/operationResults/{operationId}"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BMSPrepareDataMoveOperationResultClient) getHandleResponse(resp *http.Response) (BMSPrepareDataMoveOperationResultClientGetResponse, error) {
	result := BMSPrepareDataMoveOperationResultClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return BMSPrepareDataMoveOperationResultClientGetResponse{}, err
	}
	return result, nil
}
