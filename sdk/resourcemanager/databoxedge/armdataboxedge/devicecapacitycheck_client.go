//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdataboxedge

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

// DeviceCapacityCheckClient contains the methods for the DeviceCapacityCheck group.
// Don't use this type directly, use NewDeviceCapacityCheckClient() instead.
type DeviceCapacityCheckClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDeviceCapacityCheckClient creates a new instance of DeviceCapacityCheckClient with the specified values.
//   - subscriptionID - The subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDeviceCapacityCheckClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DeviceCapacityCheckClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DeviceCapacityCheckClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCheckResourceCreationFeasibility - Posts the device capacity request info to check feasibility.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-01
//   - resourceGroupName - The resource group name.
//   - deviceName - The device name.
//   - deviceCapacityRequestInfo - The device capacity request info.
//   - options - DeviceCapacityCheckClientBeginCheckResourceCreationFeasibilityOptions contains the optional parameters for the
//     DeviceCapacityCheckClient.BeginCheckResourceCreationFeasibility method.
func (client *DeviceCapacityCheckClient) BeginCheckResourceCreationFeasibility(ctx context.Context, resourceGroupName string, deviceName string, deviceCapacityRequestInfo DeviceCapacityRequestInfo, options *DeviceCapacityCheckClientBeginCheckResourceCreationFeasibilityOptions) (*runtime.Poller[DeviceCapacityCheckClientCheckResourceCreationFeasibilityResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.checkResourceCreationFeasibility(ctx, resourceGroupName, deviceName, deviceCapacityRequestInfo, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DeviceCapacityCheckClientCheckResourceCreationFeasibilityResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DeviceCapacityCheckClientCheckResourceCreationFeasibilityResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CheckResourceCreationFeasibility - Posts the device capacity request info to check feasibility.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-01
func (client *DeviceCapacityCheckClient) checkResourceCreationFeasibility(ctx context.Context, resourceGroupName string, deviceName string, deviceCapacityRequestInfo DeviceCapacityRequestInfo, options *DeviceCapacityCheckClientBeginCheckResourceCreationFeasibilityOptions) (*http.Response, error) {
	var err error
	const operationName = "DeviceCapacityCheckClient.BeginCheckResourceCreationFeasibility"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.checkResourceCreationFeasibilityCreateRequest(ctx, resourceGroupName, deviceName, deviceCapacityRequestInfo, options)
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

// checkResourceCreationFeasibilityCreateRequest creates the CheckResourceCreationFeasibility request.
func (client *DeviceCapacityCheckClient) checkResourceCreationFeasibilityCreateRequest(ctx context.Context, resourceGroupName string, deviceName string, deviceCapacityRequestInfo DeviceCapacityRequestInfo, options *DeviceCapacityCheckClientBeginCheckResourceCreationFeasibilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/deviceCapacityCheck"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	if options != nil && options.CapacityName != nil {
		reqQP.Set("capacityName", *options.CapacityName)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, deviceCapacityRequestInfo); err != nil {
		return nil, err
	}
	return req, nil
}
