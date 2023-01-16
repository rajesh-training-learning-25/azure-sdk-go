//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsdkreleaseplannertest

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AlertsClient contains the methods for the Alerts group.
// Don't use this type directly, use NewAlertsClient() instead.
type AlertsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAlertsClient creates a new instance of AlertsClient with the specified values.
//   - subscriptionID - The subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAlertsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AlertsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &AlertsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Gets an alert by name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-01
//   - deviceName - The device name.
//   - name - The alert name.
//   - resourceGroupName - The resource group name.
//   - options - AlertsClientGetOptions contains the optional parameters for the AlertsClient.Get method.
func (client *AlertsClient) Get(ctx context.Context, deviceName string, name string, resourceGroupName string, options *AlertsClientGetOptions) (AlertsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, deviceName, name, resourceGroupName, options)
	if err != nil {
		return AlertsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AlertsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AlertsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AlertsClient) getCreateRequest(ctx context.Context, deviceName string, name string, resourceGroupName string, options *AlertsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/alerts/{name}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AlertsClient) getHandleResponse(resp *http.Response) (AlertsClientGetResponse, error) {
	result := AlertsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Alert); err != nil {
		return AlertsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByDataBoxEdgeDevicePager - Gets all the alerts for a Data Box Edge/Data Box Gateway device.
//
// Generated from API version 2022-03-01
//   - deviceName - The device name.
//   - resourceGroupName - The resource group name.
//   - options - AlertsClientListByDataBoxEdgeDeviceOptions contains the optional parameters for the AlertsClient.NewListByDataBoxEdgeDevicePager
//     method.
func (client *AlertsClient) NewListByDataBoxEdgeDevicePager(deviceName string, resourceGroupName string, options *AlertsClientListByDataBoxEdgeDeviceOptions) *runtime.Pager[AlertsClientListByDataBoxEdgeDeviceResponse] {
	return runtime.NewPager(runtime.PagingHandler[AlertsClientListByDataBoxEdgeDeviceResponse]{
		More: func(page AlertsClientListByDataBoxEdgeDeviceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AlertsClientListByDataBoxEdgeDeviceResponse) (AlertsClientListByDataBoxEdgeDeviceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByDataBoxEdgeDeviceCreateRequest(ctx, deviceName, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AlertsClientListByDataBoxEdgeDeviceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AlertsClientListByDataBoxEdgeDeviceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AlertsClientListByDataBoxEdgeDeviceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByDataBoxEdgeDeviceHandleResponse(resp)
		},
	})
}

// listByDataBoxEdgeDeviceCreateRequest creates the ListByDataBoxEdgeDevice request.
func (client *AlertsClient) listByDataBoxEdgeDeviceCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, options *AlertsClientListByDataBoxEdgeDeviceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/alerts"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDataBoxEdgeDeviceHandleResponse handles the ListByDataBoxEdgeDevice response.
func (client *AlertsClient) listByDataBoxEdgeDeviceHandleResponse(resp *http.Response) (AlertsClientListByDataBoxEdgeDeviceResponse, error) {
	result := AlertsClientListByDataBoxEdgeDeviceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertList); err != nil {
		return AlertsClientListByDataBoxEdgeDeviceResponse{}, err
	}
	return result, nil
}
