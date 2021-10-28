//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

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

// DiskRestorePointClient contains the methods for the DiskRestorePoint group.
// Don't use this type directly, use NewDiskRestorePointClient() instead.
type DiskRestorePointClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewDiskRestorePointClient creates a new instance of DiskRestorePointClient with the specified values.
func NewDiskRestorePointClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *DiskRestorePointClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &DiskRestorePointClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Get disk restorePoint resource
// If the operation fails it returns the *CloudError error type.
func (client *DiskRestorePointClient) Get(ctx context.Context, resourceGroupName string, restorePointCollectionName string, vmRestorePointName string, diskRestorePointName string, options *DiskRestorePointGetOptions) (DiskRestorePointGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, restorePointCollectionName, vmRestorePointName, diskRestorePointName, options)
	if err != nil {
		return DiskRestorePointGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DiskRestorePointGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DiskRestorePointGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DiskRestorePointClient) getCreateRequest(ctx context.Context, resourceGroupName string, restorePointCollectionName string, vmRestorePointName string, diskRestorePointName string, options *DiskRestorePointGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/restorePointCollections/{restorePointCollectionName}/restorePoints/{vmRestorePointName}/diskRestorePoints/{diskRestorePointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if restorePointCollectionName == "" {
		return nil, errors.New("parameter restorePointCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restorePointCollectionName}", url.PathEscape(restorePointCollectionName))
	if vmRestorePointName == "" {
		return nil, errors.New("parameter vmRestorePointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmRestorePointName}", url.PathEscape(vmRestorePointName))
	if diskRestorePointName == "" {
		return nil, errors.New("parameter diskRestorePointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskRestorePointName}", url.PathEscape(diskRestorePointName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DiskRestorePointClient) getHandleResponse(resp *http.Response) (DiskRestorePointGetResponse, error) {
	result := DiskRestorePointGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiskRestorePoint); err != nil {
		return DiskRestorePointGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *DiskRestorePointClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginGrantAccess - Grants access to a diskRestorePoint.
// If the operation fails it returns the *CloudError error type.
func (client *DiskRestorePointClient) BeginGrantAccess(ctx context.Context, resourceGroupName string, restorePointCollectionName string, vmRestorePointName string, diskRestorePointName string, grantAccessData GrantAccessData, options *DiskRestorePointBeginGrantAccessOptions) (DiskRestorePointGrantAccessPollerResponse, error) {
	resp, err := client.grantAccess(ctx, resourceGroupName, restorePointCollectionName, vmRestorePointName, diskRestorePointName, grantAccessData, options)
	if err != nil {
		return DiskRestorePointGrantAccessPollerResponse{}, err
	}
	result := DiskRestorePointGrantAccessPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DiskRestorePointClient.GrantAccess", "location", resp, client.pl, client.grantAccessHandleError)
	if err != nil {
		return DiskRestorePointGrantAccessPollerResponse{}, err
	}
	result.Poller = &DiskRestorePointGrantAccessPoller{
		pt: pt,
	}
	return result, nil
}

// GrantAccess - Grants access to a diskRestorePoint.
// If the operation fails it returns the *CloudError error type.
func (client *DiskRestorePointClient) grantAccess(ctx context.Context, resourceGroupName string, restorePointCollectionName string, vmRestorePointName string, diskRestorePointName string, grantAccessData GrantAccessData, options *DiskRestorePointBeginGrantAccessOptions) (*http.Response, error) {
	req, err := client.grantAccessCreateRequest(ctx, resourceGroupName, restorePointCollectionName, vmRestorePointName, diskRestorePointName, grantAccessData, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.grantAccessHandleError(resp)
	}
	return resp, nil
}

// grantAccessCreateRequest creates the GrantAccess request.
func (client *DiskRestorePointClient) grantAccessCreateRequest(ctx context.Context, resourceGroupName string, restorePointCollectionName string, vmRestorePointName string, diskRestorePointName string, grantAccessData GrantAccessData, options *DiskRestorePointBeginGrantAccessOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/restorePointCollections/{restorePointCollectionName}/restorePoints/{vmRestorePointName}/diskRestorePoints/{diskRestorePointName}/beginGetAccess"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if restorePointCollectionName == "" {
		return nil, errors.New("parameter restorePointCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restorePointCollectionName}", url.PathEscape(restorePointCollectionName))
	if vmRestorePointName == "" {
		return nil, errors.New("parameter vmRestorePointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmRestorePointName}", url.PathEscape(vmRestorePointName))
	if diskRestorePointName == "" {
		return nil, errors.New("parameter diskRestorePointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskRestorePointName}", url.PathEscape(diskRestorePointName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, grantAccessData)
}

// grantAccessHandleError handles the GrantAccess error response.
func (client *DiskRestorePointClient) grantAccessHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByRestorePoint - Lists diskRestorePoints under a vmRestorePoint.
// If the operation fails it returns the *CloudError error type.
func (client *DiskRestorePointClient) ListByRestorePoint(resourceGroupName string, restorePointCollectionName string, vmRestorePointName string, options *DiskRestorePointListByRestorePointOptions) *DiskRestorePointListByRestorePointPager {
	return &DiskRestorePointListByRestorePointPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByRestorePointCreateRequest(ctx, resourceGroupName, restorePointCollectionName, vmRestorePointName, options)
		},
		advancer: func(ctx context.Context, resp DiskRestorePointListByRestorePointResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DiskRestorePointList.NextLink)
		},
	}
}

// listByRestorePointCreateRequest creates the ListByRestorePoint request.
func (client *DiskRestorePointClient) listByRestorePointCreateRequest(ctx context.Context, resourceGroupName string, restorePointCollectionName string, vmRestorePointName string, options *DiskRestorePointListByRestorePointOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/restorePointCollections/{restorePointCollectionName}/restorePoints/{vmRestorePointName}/diskRestorePoints"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if restorePointCollectionName == "" {
		return nil, errors.New("parameter restorePointCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restorePointCollectionName}", url.PathEscape(restorePointCollectionName))
	if vmRestorePointName == "" {
		return nil, errors.New("parameter vmRestorePointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmRestorePointName}", url.PathEscape(vmRestorePointName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByRestorePointHandleResponse handles the ListByRestorePoint response.
func (client *DiskRestorePointClient) listByRestorePointHandleResponse(resp *http.Response) (DiskRestorePointListByRestorePointResponse, error) {
	result := DiskRestorePointListByRestorePointResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiskRestorePointList); err != nil {
		return DiskRestorePointListByRestorePointResponse{}, err
	}
	return result, nil
}

// listByRestorePointHandleError handles the ListByRestorePoint error response.
func (client *DiskRestorePointClient) listByRestorePointHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginRevokeAccess - Revokes access to a diskRestorePoint.
// If the operation fails it returns the *CloudError error type.
func (client *DiskRestorePointClient) BeginRevokeAccess(ctx context.Context, resourceGroupName string, restorePointCollectionName string, vmRestorePointName string, diskRestorePointName string, options *DiskRestorePointBeginRevokeAccessOptions) (DiskRestorePointRevokeAccessPollerResponse, error) {
	resp, err := client.revokeAccess(ctx, resourceGroupName, restorePointCollectionName, vmRestorePointName, diskRestorePointName, options)
	if err != nil {
		return DiskRestorePointRevokeAccessPollerResponse{}, err
	}
	result := DiskRestorePointRevokeAccessPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DiskRestorePointClient.RevokeAccess", "location", resp, client.pl, client.revokeAccessHandleError)
	if err != nil {
		return DiskRestorePointRevokeAccessPollerResponse{}, err
	}
	result.Poller = &DiskRestorePointRevokeAccessPoller{
		pt: pt,
	}
	return result, nil
}

// RevokeAccess - Revokes access to a diskRestorePoint.
// If the operation fails it returns the *CloudError error type.
func (client *DiskRestorePointClient) revokeAccess(ctx context.Context, resourceGroupName string, restorePointCollectionName string, vmRestorePointName string, diskRestorePointName string, options *DiskRestorePointBeginRevokeAccessOptions) (*http.Response, error) {
	req, err := client.revokeAccessCreateRequest(ctx, resourceGroupName, restorePointCollectionName, vmRestorePointName, diskRestorePointName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.revokeAccessHandleError(resp)
	}
	return resp, nil
}

// revokeAccessCreateRequest creates the RevokeAccess request.
func (client *DiskRestorePointClient) revokeAccessCreateRequest(ctx context.Context, resourceGroupName string, restorePointCollectionName string, vmRestorePointName string, diskRestorePointName string, options *DiskRestorePointBeginRevokeAccessOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/restorePointCollections/{restorePointCollectionName}/restorePoints/{vmRestorePointName}/diskRestorePoints/{diskRestorePointName}/endGetAccess"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if restorePointCollectionName == "" {
		return nil, errors.New("parameter restorePointCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restorePointCollectionName}", url.PathEscape(restorePointCollectionName))
	if vmRestorePointName == "" {
		return nil, errors.New("parameter vmRestorePointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmRestorePointName}", url.PathEscape(vmRestorePointName))
	if diskRestorePointName == "" {
		return nil, errors.New("parameter diskRestorePointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskRestorePointName}", url.PathEscape(diskRestorePointName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// revokeAccessHandleError handles the RevokeAccess error response.
func (client *DiskRestorePointClient) revokeAccessHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
