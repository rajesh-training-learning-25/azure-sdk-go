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

// BackupUsageSummariesClient contains the methods for the BackupUsageSummaries group.
// Don't use this type directly, use NewBackupUsageSummariesClient() instead.
type BackupUsageSummariesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewBackupUsageSummariesClient creates a new instance of BackupUsageSummariesClient with the specified values.
func NewBackupUsageSummariesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *BackupUsageSummariesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &BackupUsageSummariesClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// List - Fetches the backup management usage summaries of the vault.
// If the operation fails it returns a generic error.
func (client *BackupUsageSummariesClient) List(ctx context.Context, vaultName string, resourceGroupName string, options *BackupUsageSummariesListOptions) (BackupUsageSummariesListResponse, error) {
	req, err := client.listCreateRequest(ctx, vaultName, resourceGroupName, options)
	if err != nil {
		return BackupUsageSummariesListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BackupUsageSummariesListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BackupUsageSummariesListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *BackupUsageSummariesClient) listCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, options *BackupUsageSummariesListOptions) (*policy.Request, error) {
	urlPath := "/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupUsageSummaries"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BackupUsageSummariesClient) listHandleResponse(resp *http.Response) (BackupUsageSummariesListResponse, error) {
	result := BackupUsageSummariesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.BackupManagementUsageList); err != nil {
		return BackupUsageSummariesListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *BackupUsageSummariesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
