//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
	"net/http"
	"regexp"
)

// VirtualMachineScaleSetVMExtensionsServer is a fake server for instances of the armcompute.VirtualMachineScaleSetVMExtensionsClient type.
type VirtualMachineScaleSetVMExtensionsServer struct {
	// BeginCreateOrUpdate is the fake for method VirtualMachineScaleSetVMExtensionsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, vmExtensionName string, extensionParameters armcompute.VirtualMachineScaleSetVMExtension, options *armcompute.VirtualMachineScaleSetVMExtensionsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcompute.VirtualMachineScaleSetVMExtensionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method VirtualMachineScaleSetVMExtensionsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, vmExtensionName string, options *armcompute.VirtualMachineScaleSetVMExtensionsClientBeginDeleteOptions) (resp azfake.PollerResponder[armcompute.VirtualMachineScaleSetVMExtensionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method VirtualMachineScaleSetVMExtensionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, vmExtensionName string, options *armcompute.VirtualMachineScaleSetVMExtensionsClientGetOptions) (resp azfake.Responder[armcompute.VirtualMachineScaleSetVMExtensionsClientGetResponse], errResp azfake.ErrorResponder)

	// List is the fake for method VirtualMachineScaleSetVMExtensionsClient.List
	// HTTP status codes to indicate success: http.StatusOK
	List func(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, options *armcompute.VirtualMachineScaleSetVMExtensionsClientListOptions) (resp azfake.Responder[armcompute.VirtualMachineScaleSetVMExtensionsClientListResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method VirtualMachineScaleSetVMExtensionsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK
	BeginUpdate func(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, vmExtensionName string, extensionParameters armcompute.VirtualMachineScaleSetVMExtensionUpdate, options *armcompute.VirtualMachineScaleSetVMExtensionsClientBeginUpdateOptions) (resp azfake.PollerResponder[armcompute.VirtualMachineScaleSetVMExtensionsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewVirtualMachineScaleSetVMExtensionsServerTransport creates a new instance of VirtualMachineScaleSetVMExtensionsServerTransport with the provided implementation.
// The returned VirtualMachineScaleSetVMExtensionsServerTransport instance is connected to an instance of armcompute.VirtualMachineScaleSetVMExtensionsClient by way of the
// undefined.Transporter field.
func NewVirtualMachineScaleSetVMExtensionsServerTransport(srv *VirtualMachineScaleSetVMExtensionsServer) *VirtualMachineScaleSetVMExtensionsServerTransport {
	return &VirtualMachineScaleSetVMExtensionsServerTransport{srv: srv}
}

// VirtualMachineScaleSetVMExtensionsServerTransport connects instances of armcompute.VirtualMachineScaleSetVMExtensionsClient to instances of VirtualMachineScaleSetVMExtensionsServer.
// Don't use this type directly, use NewVirtualMachineScaleSetVMExtensionsServerTransport instead.
type VirtualMachineScaleSetVMExtensionsServerTransport struct {
	srv                 *VirtualMachineScaleSetVMExtensionsServer
	beginCreateOrUpdate *azfake.PollerResponder[armcompute.VirtualMachineScaleSetVMExtensionsClientCreateOrUpdateResponse]
	beginDelete         *azfake.PollerResponder[armcompute.VirtualMachineScaleSetVMExtensionsClientDeleteResponse]
	beginUpdate         *azfake.PollerResponder[armcompute.VirtualMachineScaleSetVMExtensionsClientUpdateResponse]
}

// Do implements the policy.Transporter interface for VirtualMachineScaleSetVMExtensionsServerTransport.
func (v *VirtualMachineScaleSetVMExtensionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VirtualMachineScaleSetVMExtensionsClient.BeginCreateOrUpdate":
		resp, err = v.dispatchBeginCreateOrUpdate(req)
	case "VirtualMachineScaleSetVMExtensionsClient.BeginDelete":
		resp, err = v.dispatchBeginDelete(req)
	case "VirtualMachineScaleSetVMExtensionsClient.Get":
		resp, err = v.dispatchGet(req)
	case "VirtualMachineScaleSetVMExtensionsClient.List":
		resp, err = v.dispatchList(req)
	case "VirtualMachineScaleSetVMExtensionsClient.BeginUpdate":
		resp, err = v.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VirtualMachineScaleSetVMExtensionsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("method BeginCreateOrUpdate not implemented")}
	}
	if v.beginCreateOrUpdate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/virtualMachineScaleSets/(?P<vmScaleSetName>[a-zA-Z0-9-_]+)/virtualMachines/(?P<instanceId>[a-zA-Z0-9-_]+)/extensions/(?P<vmExtensionName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.VirtualMachineScaleSetVMExtension](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginCreateOrUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("vmScaleSetName")], matches[regex.SubexpIndex("instanceId")], matches[regex.SubexpIndex("vmExtensionName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		v.beginCreateOrUpdate = &respr
	}

	resp, err := server.PollerResponderNext(v.beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(v.beginCreateOrUpdate) {
		v.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (v *VirtualMachineScaleSetVMExtensionsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if v.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("method BeginDelete not implemented")}
	}
	if v.beginDelete == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/virtualMachineScaleSets/(?P<vmScaleSetName>[a-zA-Z0-9-_]+)/virtualMachines/(?P<instanceId>[a-zA-Z0-9-_]+)/extensions/(?P<vmExtensionName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := v.srv.BeginDelete(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("vmScaleSetName")], matches[regex.SubexpIndex("instanceId")], matches[regex.SubexpIndex("vmExtensionName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		v.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(v.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(v.beginDelete) {
		v.beginDelete = nil
	}

	return resp, nil
}

func (v *VirtualMachineScaleSetVMExtensionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/virtualMachineScaleSets/(?P<vmScaleSetName>[a-zA-Z0-9-_]+)/virtualMachines/(?P<instanceId>[a-zA-Z0-9-_]+)/extensions/(?P<vmExtensionName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	expandParam := getOptional(qp.Get("$expand"))
	var options *armcompute.VirtualMachineScaleSetVMExtensionsClientGetOptions
	if expandParam != nil {
		options = &armcompute.VirtualMachineScaleSetVMExtensionsClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := v.srv.Get(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("vmScaleSetName")], matches[regex.SubexpIndex("instanceId")], matches[regex.SubexpIndex("vmExtensionName")], options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualMachineScaleSetVMExtension, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualMachineScaleSetVMExtensionsServerTransport) dispatchList(req *http.Request) (*http.Response, error) {
	if v.srv.List == nil {
		return nil, &nonRetriableError{errors.New("method List not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/virtualMachineScaleSets/(?P<vmScaleSetName>[a-zA-Z0-9-_]+)/virtualMachines/(?P<instanceId>[a-zA-Z0-9-_]+)/extensions"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	expandParam := getOptional(qp.Get("$expand"))
	var options *armcompute.VirtualMachineScaleSetVMExtensionsClientListOptions
	if expandParam != nil {
		options = &armcompute.VirtualMachineScaleSetVMExtensionsClientListOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := v.srv.List(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("vmScaleSetName")], matches[regex.SubexpIndex("instanceId")], options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualMachineScaleSetVMExtensionsListResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualMachineScaleSetVMExtensionsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("method BeginUpdate not implemented")}
	}
	if v.beginUpdate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/virtualMachineScaleSets/(?P<vmScaleSetName>[a-zA-Z0-9-_]+)/virtualMachines/(?P<instanceId>[a-zA-Z0-9-_]+)/extensions/(?P<vmExtensionName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.VirtualMachineScaleSetVMExtensionUpdate](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("vmScaleSetName")], matches[regex.SubexpIndex("instanceId")], matches[regex.SubexpIndex("vmExtensionName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		v.beginUpdate = &respr
	}

	resp, err := server.PollerResponderNext(v.beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PollerResponderMore(v.beginUpdate) {
		v.beginUpdate = nil
	}

	return resp, nil
}
