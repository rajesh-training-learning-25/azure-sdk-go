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

// ChapSettingsClient contains the methods for the ChapSettings group.
// Don't use this type directly, use NewChapSettingsClient() instead.
type ChapSettingsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewChapSettingsClient creates a new instance of ChapSettingsClient with the specified values.
//   - subscriptionID - The subscription id
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewChapSettingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ChapSettingsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ChapSettingsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates the chap setting.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-10-01
//   - deviceName - The device name.
//   - chapUserName - The chap user name.
//   - resourceGroupName - The resource group name
//   - managerName - The manager name
//   - chapSetting - The chap setting to be added or updated.
//   - options - ChapSettingsClientBeginCreateOrUpdateOptions contains the optional parameters for the ChapSettingsClient.BeginCreateOrUpdate
//     method.
func (client *ChapSettingsClient) BeginCreateOrUpdate(ctx context.Context, deviceName string, chapUserName string, resourceGroupName string, managerName string, chapSetting ChapSettings, options *ChapSettingsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ChapSettingsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, deviceName, chapUserName, resourceGroupName, managerName, chapSetting, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ChapSettingsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ChapSettingsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates the chap setting.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-10-01
func (client *ChapSettingsClient) createOrUpdate(ctx context.Context, deviceName string, chapUserName string, resourceGroupName string, managerName string, chapSetting ChapSettings, options *ChapSettingsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ChapSettingsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, deviceName, chapUserName, resourceGroupName, managerName, chapSetting, options)
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
func (client *ChapSettingsClient) createOrUpdateCreateRequest(ctx context.Context, deviceName string, chapUserName string, resourceGroupName string, managerName string, chapSetting ChapSettings, options *ChapSettingsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/chapSettings/{chapUserName}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if chapUserName == "" {
		return nil, errors.New("parameter chapUserName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{chapUserName}", url.PathEscape(chapUserName))
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
	if err := runtime.MarshalAsJSON(req, chapSetting); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the chap setting.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-10-01
//   - deviceName - The device name.
//   - chapUserName - The chap user name.
//   - resourceGroupName - The resource group name
//   - managerName - The manager name
//   - options - ChapSettingsClientBeginDeleteOptions contains the optional parameters for the ChapSettingsClient.BeginDelete
//     method.
func (client *ChapSettingsClient) BeginDelete(ctx context.Context, deviceName string, chapUserName string, resourceGroupName string, managerName string, options *ChapSettingsClientBeginDeleteOptions) (*runtime.Poller[ChapSettingsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, deviceName, chapUserName, resourceGroupName, managerName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ChapSettingsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ChapSettingsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the chap setting.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-10-01
func (client *ChapSettingsClient) deleteOperation(ctx context.Context, deviceName string, chapUserName string, resourceGroupName string, managerName string, options *ChapSettingsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ChapSettingsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, deviceName, chapUserName, resourceGroupName, managerName, options)
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
func (client *ChapSettingsClient) deleteCreateRequest(ctx context.Context, deviceName string, chapUserName string, resourceGroupName string, managerName string, options *ChapSettingsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/chapSettings/{chapUserName}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if chapUserName == "" {
		return nil, errors.New("parameter chapUserName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{chapUserName}", url.PathEscape(chapUserName))
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

// Get - Returns the properties of the specified chap setting name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-10-01
//   - deviceName - The device name.
//   - chapUserName - The user name of chap to be fetched.
//   - resourceGroupName - The resource group name
//   - managerName - The manager name
//   - options - ChapSettingsClientGetOptions contains the optional parameters for the ChapSettingsClient.Get method.
func (client *ChapSettingsClient) Get(ctx context.Context, deviceName string, chapUserName string, resourceGroupName string, managerName string, options *ChapSettingsClientGetOptions) (ChapSettingsClientGetResponse, error) {
	var err error
	const operationName = "ChapSettingsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, deviceName, chapUserName, resourceGroupName, managerName, options)
	if err != nil {
		return ChapSettingsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ChapSettingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ChapSettingsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ChapSettingsClient) getCreateRequest(ctx context.Context, deviceName string, chapUserName string, resourceGroupName string, managerName string, options *ChapSettingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/chapSettings/{chapUserName}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if chapUserName == "" {
		return nil, errors.New("parameter chapUserName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{chapUserName}", url.PathEscape(chapUserName))
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
func (client *ChapSettingsClient) getHandleResponse(resp *http.Response) (ChapSettingsClientGetResponse, error) {
	result := ChapSettingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ChapSettings); err != nil {
		return ChapSettingsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByDevicePager - Retrieves all the chap settings in a device.
//
// Generated from API version 2016-10-01
//   - deviceName - The name of the device.
//   - resourceGroupName - The resource group name
//   - managerName - The manager name
//   - options - ChapSettingsClientListByDeviceOptions contains the optional parameters for the ChapSettingsClient.NewListByDevicePager
//     method.
func (client *ChapSettingsClient) NewListByDevicePager(deviceName string, resourceGroupName string, managerName string, options *ChapSettingsClientListByDeviceOptions) *runtime.Pager[ChapSettingsClientListByDeviceResponse] {
	return runtime.NewPager(runtime.PagingHandler[ChapSettingsClientListByDeviceResponse]{
		More: func(page ChapSettingsClientListByDeviceResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ChapSettingsClientListByDeviceResponse) (ChapSettingsClientListByDeviceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ChapSettingsClient.NewListByDevicePager")
			req, err := client.listByDeviceCreateRequest(ctx, deviceName, resourceGroupName, managerName, options)
			if err != nil {
				return ChapSettingsClientListByDeviceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ChapSettingsClientListByDeviceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ChapSettingsClientListByDeviceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByDeviceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByDeviceCreateRequest creates the ListByDevice request.
func (client *ChapSettingsClient) listByDeviceCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, managerName string, options *ChapSettingsClientListByDeviceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/chapSettings"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
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

// listByDeviceHandleResponse handles the ListByDevice response.
func (client *ChapSettingsClient) listByDeviceHandleResponse(resp *http.Response) (ChapSettingsClientListByDeviceResponse, error) {
	result := ChapSettingsClientListByDeviceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ChapSettingsList); err != nil {
		return ChapSettingsClientListByDeviceResponse{}, err
	}
	return result, nil
}
