//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorsimple1200series

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

// AccessControlRecordsClient contains the methods for the AccessControlRecords group.
// Don't use this type directly, use NewAccessControlRecordsClient() instead.
type AccessControlRecordsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAccessControlRecordsClient creates a new instance of AccessControlRecordsClient with the specified values.
//   - subscriptionID - The subscription id
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAccessControlRecordsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AccessControlRecordsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AccessControlRecordsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or Updates an access control record.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-10-01
//   - accessControlRecordName - The name of the access control record.
//   - resourceGroupName - The resource group name
//   - managerName - The manager name
//   - accessControlRecord - The access control record to be added or updated.
//   - options - AccessControlRecordsClientBeginCreateOrUpdateOptions contains the optional parameters for the AccessControlRecordsClient.BeginCreateOrUpdate
//     method.
func (client *AccessControlRecordsClient) BeginCreateOrUpdate(ctx context.Context, accessControlRecordName string, resourceGroupName string, managerName string, accessControlRecord AccessControlRecord, options *AccessControlRecordsClientBeginCreateOrUpdateOptions) (*runtime.Poller[AccessControlRecordsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, accessControlRecordName, resourceGroupName, managerName, accessControlRecord, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AccessControlRecordsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AccessControlRecordsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or Updates an access control record.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-10-01
func (client *AccessControlRecordsClient) createOrUpdate(ctx context.Context, accessControlRecordName string, resourceGroupName string, managerName string, accessControlRecord AccessControlRecord, options *AccessControlRecordsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "AccessControlRecordsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, accessControlRecordName, resourceGroupName, managerName, accessControlRecord, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AccessControlRecordsClient) createOrUpdateCreateRequest(ctx context.Context, accessControlRecordName string, resourceGroupName string, managerName string, accessControlRecord AccessControlRecord, options *AccessControlRecordsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/accessControlRecords/{accessControlRecordName}"
	if accessControlRecordName == "" {
		return nil, errors.New("parameter accessControlRecordName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accessControlRecordName}", url.PathEscape(accessControlRecordName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managerName == "" {
		return nil, errors.New("parameter managerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", url.PathEscape(managerName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, accessControlRecord); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the access control record.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-10-01
//   - accessControlRecordName - The name of the access control record to delete.
//   - resourceGroupName - The resource group name
//   - managerName - The manager name
//   - options - AccessControlRecordsClientBeginDeleteOptions contains the optional parameters for the AccessControlRecordsClient.BeginDelete
//     method.
func (client *AccessControlRecordsClient) BeginDelete(ctx context.Context, accessControlRecordName string, resourceGroupName string, managerName string, options *AccessControlRecordsClientBeginDeleteOptions) (*runtime.Poller[AccessControlRecordsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, accessControlRecordName, resourceGroupName, managerName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AccessControlRecordsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AccessControlRecordsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the access control record.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-10-01
func (client *AccessControlRecordsClient) deleteOperation(ctx context.Context, accessControlRecordName string, resourceGroupName string, managerName string, options *AccessControlRecordsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "AccessControlRecordsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, accessControlRecordName, resourceGroupName, managerName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AccessControlRecordsClient) deleteCreateRequest(ctx context.Context, accessControlRecordName string, resourceGroupName string, managerName string, options *AccessControlRecordsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/accessControlRecords/{accessControlRecordName}"
	if accessControlRecordName == "" {
		return nil, errors.New("parameter accessControlRecordName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accessControlRecordName}", url.PathEscape(accessControlRecordName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managerName == "" {
		return nil, errors.New("parameter managerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", url.PathEscape(managerName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Returns the properties of the specified access control record name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-10-01
//   - accessControlRecordName - Name of access control record to be fetched.
//   - resourceGroupName - The resource group name
//   - managerName - The manager name
//   - options - AccessControlRecordsClientGetOptions contains the optional parameters for the AccessControlRecordsClient.Get
//     method.
func (client *AccessControlRecordsClient) Get(ctx context.Context, accessControlRecordName string, resourceGroupName string, managerName string, options *AccessControlRecordsClientGetOptions) (AccessControlRecordsClientGetResponse, error) {
	var err error
	const operationName = "AccessControlRecordsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, accessControlRecordName, resourceGroupName, managerName, options)
	if err != nil {
		return AccessControlRecordsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AccessControlRecordsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AccessControlRecordsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AccessControlRecordsClient) getCreateRequest(ctx context.Context, accessControlRecordName string, resourceGroupName string, managerName string, options *AccessControlRecordsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/accessControlRecords/{accessControlRecordName}"
	if accessControlRecordName == "" {
		return nil, errors.New("parameter accessControlRecordName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accessControlRecordName}", url.PathEscape(accessControlRecordName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managerName == "" {
		return nil, errors.New("parameter managerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", url.PathEscape(managerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AccessControlRecordsClient) getHandleResponse(resp *http.Response) (AccessControlRecordsClientGetResponse, error) {
	result := AccessControlRecordsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessControlRecord); err != nil {
		return AccessControlRecordsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByManagerPager - Retrieves all the access control records in a manager.
//
// Generated from API version 2016-10-01
//   - resourceGroupName - The resource group name
//   - managerName - The manager name
//   - options - AccessControlRecordsClientListByManagerOptions contains the optional parameters for the AccessControlRecordsClient.NewListByManagerPager
//     method.
func (client *AccessControlRecordsClient) NewListByManagerPager(resourceGroupName string, managerName string, options *AccessControlRecordsClientListByManagerOptions) *runtime.Pager[AccessControlRecordsClientListByManagerResponse] {
	return runtime.NewPager(runtime.PagingHandler[AccessControlRecordsClientListByManagerResponse]{
		More: func(page AccessControlRecordsClientListByManagerResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *AccessControlRecordsClientListByManagerResponse) (AccessControlRecordsClientListByManagerResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AccessControlRecordsClient.NewListByManagerPager")
			req, err := client.listByManagerCreateRequest(ctx, resourceGroupName, managerName, options)
			if err != nil {
				return AccessControlRecordsClientListByManagerResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return AccessControlRecordsClientListByManagerResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AccessControlRecordsClientListByManagerResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByManagerHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByManagerCreateRequest creates the ListByManager request.
func (client *AccessControlRecordsClient) listByManagerCreateRequest(ctx context.Context, resourceGroupName string, managerName string, options *AccessControlRecordsClientListByManagerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/accessControlRecords"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managerName == "" {
		return nil, errors.New("parameter managerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", url.PathEscape(managerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByManagerHandleResponse handles the ListByManager response.
func (client *AccessControlRecordsClient) listByManagerHandleResponse(resp *http.Response) (AccessControlRecordsClientListByManagerResponse, error) {
	result := AccessControlRecordsClientListByManagerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessControlRecordList); err != nil {
		return AccessControlRecordsClientListByManagerResponse{}, err
	}
	return result, nil
}
