//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armguestconfiguration

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

// AssignmentReportsVMSSClient contains the methods for the GuestConfigurationAssignmentReportsVMSS group.
// Don't use this type directly, use NewAssignmentReportsVMSSClient() instead.
type AssignmentReportsVMSSClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAssignmentReportsVMSSClient creates a new instance of AssignmentReportsVMSSClient with the specified values.
//   - subscriptionID - Subscription ID which uniquely identify Microsoft Azure subscription. The subscription ID forms part of
//     the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAssignmentReportsVMSSClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AssignmentReportsVMSSClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AssignmentReportsVMSSClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get a report for the VMSS guest configuration assignment, by reportId.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-01-25
//   - resourceGroupName - The resource group name.
//   - vmssName - The name of the virtual machine scale set.
//   - name - The guest configuration assignment name.
//   - id - The GUID for the guest configuration assignment report.
//   - options - AssignmentReportsVMSSClientGetOptions contains the optional parameters for the AssignmentReportsVMSSClient.Get
//     method.
func (client *AssignmentReportsVMSSClient) Get(ctx context.Context, resourceGroupName string, vmssName string, name string, id string, options *AssignmentReportsVMSSClientGetOptions) (AssignmentReportsVMSSClientGetResponse, error) {
	var err error
	const operationName = "AssignmentReportsVMSSClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, vmssName, name, id, options)
	if err != nil {
		return AssignmentReportsVMSSClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AssignmentReportsVMSSClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AssignmentReportsVMSSClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AssignmentReportsVMSSClient) getCreateRequest(ctx context.Context, resourceGroupName string, vmssName string, name string, id string, options *AssignmentReportsVMSSClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmssName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{name}/reports/{id}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmssName == "" {
		return nil, errors.New("parameter vmssName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmssName}", url.PathEscape(vmssName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-25")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AssignmentReportsVMSSClient) getHandleResponse(resp *http.Response) (AssignmentReportsVMSSClientGetResponse, error) {
	result := AssignmentReportsVMSSClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssignmentReport); err != nil {
		return AssignmentReportsVMSSClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List all reports for the VMSS guest configuration assignment, latest report first.
//
// Generated from API version 2022-01-25
//   - resourceGroupName - The resource group name.
//   - vmssName - The name of the virtual machine scale set.
//   - name - The guest configuration assignment name.
//   - options - AssignmentReportsVMSSClientListOptions contains the optional parameters for the AssignmentReportsVMSSClient.NewListPager
//     method.
func (client *AssignmentReportsVMSSClient) NewListPager(resourceGroupName string, vmssName string, name string, options *AssignmentReportsVMSSClientListOptions) *runtime.Pager[AssignmentReportsVMSSClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AssignmentReportsVMSSClientListResponse]{
		More: func(page AssignmentReportsVMSSClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *AssignmentReportsVMSSClientListResponse) (AssignmentReportsVMSSClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AssignmentReportsVMSSClient.NewListPager")
			req, err := client.listCreateRequest(ctx, resourceGroupName, vmssName, name, options)
			if err != nil {
				return AssignmentReportsVMSSClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return AssignmentReportsVMSSClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AssignmentReportsVMSSClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *AssignmentReportsVMSSClient) listCreateRequest(ctx context.Context, resourceGroupName string, vmssName string, name string, options *AssignmentReportsVMSSClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmssName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{name}/reports"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmssName == "" {
		return nil, errors.New("parameter vmssName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmssName}", url.PathEscape(vmssName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-25")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AssignmentReportsVMSSClient) listHandleResponse(resp *http.Response) (AssignmentReportsVMSSClientListResponse, error) {
	result := AssignmentReportsVMSSClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssignmentReportList); err != nil {
		return AssignmentReportsVMSSClientListResponse{}, err
	}
	return result, nil
}
