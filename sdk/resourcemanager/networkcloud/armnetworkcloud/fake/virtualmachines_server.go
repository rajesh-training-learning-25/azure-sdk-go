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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
)

// VirtualMachinesServer is a fake server for instances of the armnetworkcloud.VirtualMachinesClient type.
type VirtualMachinesServer struct {
	// BeginCreateOrUpdate is the fake for method VirtualMachinesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, virtualMachineName string, virtualMachineParameters armnetworkcloud.VirtualMachine, options *armnetworkcloud.VirtualMachinesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetworkcloud.VirtualMachinesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method VirtualMachinesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, virtualMachineName string, options *armnetworkcloud.VirtualMachinesClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetworkcloud.VirtualMachinesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method VirtualMachinesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, virtualMachineName string, options *armnetworkcloud.VirtualMachinesClientGetOptions) (resp azfake.Responder[armnetworkcloud.VirtualMachinesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method VirtualMachinesClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armnetworkcloud.VirtualMachinesClientListByResourceGroupOptions) (resp azfake.PagerResponder[armnetworkcloud.VirtualMachinesClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method VirtualMachinesClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armnetworkcloud.VirtualMachinesClientListBySubscriptionOptions) (resp azfake.PagerResponder[armnetworkcloud.VirtualMachinesClientListBySubscriptionResponse])

	// BeginPowerOff is the fake for method VirtualMachinesClient.BeginPowerOff
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginPowerOff func(ctx context.Context, resourceGroupName string, virtualMachineName string, options *armnetworkcloud.VirtualMachinesClientBeginPowerOffOptions) (resp azfake.PollerResponder[armnetworkcloud.VirtualMachinesClientPowerOffResponse], errResp azfake.ErrorResponder)

	// BeginReimage is the fake for method VirtualMachinesClient.BeginReimage
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginReimage func(ctx context.Context, resourceGroupName string, virtualMachineName string, options *armnetworkcloud.VirtualMachinesClientBeginReimageOptions) (resp azfake.PollerResponder[armnetworkcloud.VirtualMachinesClientReimageResponse], errResp azfake.ErrorResponder)

	// BeginRestart is the fake for method VirtualMachinesClient.BeginRestart
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginRestart func(ctx context.Context, resourceGroupName string, virtualMachineName string, options *armnetworkcloud.VirtualMachinesClientBeginRestartOptions) (resp azfake.PollerResponder[armnetworkcloud.VirtualMachinesClientRestartResponse], errResp azfake.ErrorResponder)

	// BeginStart is the fake for method VirtualMachinesClient.BeginStart
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginStart func(ctx context.Context, resourceGroupName string, virtualMachineName string, options *armnetworkcloud.VirtualMachinesClientBeginStartOptions) (resp azfake.PollerResponder[armnetworkcloud.VirtualMachinesClientStartResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method VirtualMachinesClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, virtualMachineName string, virtualMachineUpdateParameters armnetworkcloud.VirtualMachinePatchParameters, options *armnetworkcloud.VirtualMachinesClientBeginUpdateOptions) (resp azfake.PollerResponder[armnetworkcloud.VirtualMachinesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewVirtualMachinesServerTransport creates a new instance of VirtualMachinesServerTransport with the provided implementation.
// The returned VirtualMachinesServerTransport instance is connected to an instance of armnetworkcloud.VirtualMachinesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVirtualMachinesServerTransport(srv *VirtualMachinesServer) *VirtualMachinesServerTransport {
	return &VirtualMachinesServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armnetworkcloud.VirtualMachinesClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armnetworkcloud.VirtualMachinesClientDeleteResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armnetworkcloud.VirtualMachinesClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armnetworkcloud.VirtualMachinesClientListBySubscriptionResponse]](),
		beginPowerOff:               newTracker[azfake.PollerResponder[armnetworkcloud.VirtualMachinesClientPowerOffResponse]](),
		beginReimage:                newTracker[azfake.PollerResponder[armnetworkcloud.VirtualMachinesClientReimageResponse]](),
		beginRestart:                newTracker[azfake.PollerResponder[armnetworkcloud.VirtualMachinesClientRestartResponse]](),
		beginStart:                  newTracker[azfake.PollerResponder[armnetworkcloud.VirtualMachinesClientStartResponse]](),
		beginUpdate:                 newTracker[azfake.PollerResponder[armnetworkcloud.VirtualMachinesClientUpdateResponse]](),
	}
}

// VirtualMachinesServerTransport connects instances of armnetworkcloud.VirtualMachinesClient to instances of VirtualMachinesServer.
// Don't use this type directly, use NewVirtualMachinesServerTransport instead.
type VirtualMachinesServerTransport struct {
	srv                         *VirtualMachinesServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armnetworkcloud.VirtualMachinesClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armnetworkcloud.VirtualMachinesClientDeleteResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armnetworkcloud.VirtualMachinesClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armnetworkcloud.VirtualMachinesClientListBySubscriptionResponse]]
	beginPowerOff               *tracker[azfake.PollerResponder[armnetworkcloud.VirtualMachinesClientPowerOffResponse]]
	beginReimage                *tracker[azfake.PollerResponder[armnetworkcloud.VirtualMachinesClientReimageResponse]]
	beginRestart                *tracker[azfake.PollerResponder[armnetworkcloud.VirtualMachinesClientRestartResponse]]
	beginStart                  *tracker[azfake.PollerResponder[armnetworkcloud.VirtualMachinesClientStartResponse]]
	beginUpdate                 *tracker[azfake.PollerResponder[armnetworkcloud.VirtualMachinesClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for VirtualMachinesServerTransport.
func (v *VirtualMachinesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VirtualMachinesClient.BeginCreateOrUpdate":
		resp, err = v.dispatchBeginCreateOrUpdate(req)
	case "VirtualMachinesClient.BeginDelete":
		resp, err = v.dispatchBeginDelete(req)
	case "VirtualMachinesClient.Get":
		resp, err = v.dispatchGet(req)
	case "VirtualMachinesClient.NewListByResourceGroupPager":
		resp, err = v.dispatchNewListByResourceGroupPager(req)
	case "VirtualMachinesClient.NewListBySubscriptionPager":
		resp, err = v.dispatchNewListBySubscriptionPager(req)
	case "VirtualMachinesClient.BeginPowerOff":
		resp, err = v.dispatchBeginPowerOff(req)
	case "VirtualMachinesClient.BeginReimage":
		resp, err = v.dispatchBeginReimage(req)
	case "VirtualMachinesClient.BeginRestart":
		resp, err = v.dispatchBeginRestart(req)
	case "VirtualMachinesClient.BeginStart":
		resp, err = v.dispatchBeginStart(req)
	case "VirtualMachinesClient.BeginUpdate":
		resp, err = v.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VirtualMachinesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := v.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/virtualMachines/(?P<virtualMachineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetworkcloud.VirtualMachine](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualMachineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualMachineName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, virtualMachineNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		v.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		v.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		v.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (v *VirtualMachinesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if v.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := v.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/virtualMachines/(?P<virtualMachineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualMachineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualMachineName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginDelete(req.Context(), resourceGroupNameParam, virtualMachineNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		v.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		v.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		v.beginDelete.remove(req)
	}

	return resp, nil
}

func (v *VirtualMachinesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/virtualMachines/(?P<virtualMachineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	virtualMachineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualMachineName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Get(req.Context(), resourceGroupNameParam, virtualMachineNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualMachine, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualMachinesServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := v.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/virtualMachines`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := v.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		v.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armnetworkcloud.VirtualMachinesClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		v.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		v.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (v *VirtualMachinesServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := v.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/virtualMachines`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := v.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		v.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armnetworkcloud.VirtualMachinesClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		v.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		v.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (v *VirtualMachinesServerTransport) dispatchBeginPowerOff(req *http.Request) (*http.Response, error) {
	if v.srv.BeginPowerOff == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPowerOff not implemented")}
	}
	beginPowerOff := v.beginPowerOff.get(req)
	if beginPowerOff == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/virtualMachines/(?P<virtualMachineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/powerOff`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetworkcloud.VirtualMachinePowerOffParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualMachineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualMachineName")])
		if err != nil {
			return nil, err
		}
		var options *armnetworkcloud.VirtualMachinesClientBeginPowerOffOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armnetworkcloud.VirtualMachinesClientBeginPowerOffOptions{
				VirtualMachinePowerOffParameters: &body,
			}
		}
		respr, errRespr := v.srv.BeginPowerOff(req.Context(), resourceGroupNameParam, virtualMachineNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginPowerOff = &respr
		v.beginPowerOff.add(req, beginPowerOff)
	}

	resp, err := server.PollerResponderNext(beginPowerOff, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		v.beginPowerOff.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginPowerOff) {
		v.beginPowerOff.remove(req)
	}

	return resp, nil
}

func (v *VirtualMachinesServerTransport) dispatchBeginReimage(req *http.Request) (*http.Response, error) {
	if v.srv.BeginReimage == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginReimage not implemented")}
	}
	beginReimage := v.beginReimage.get(req)
	if beginReimage == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/virtualMachines/(?P<virtualMachineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reimage`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualMachineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualMachineName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginReimage(req.Context(), resourceGroupNameParam, virtualMachineNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginReimage = &respr
		v.beginReimage.add(req, beginReimage)
	}

	resp, err := server.PollerResponderNext(beginReimage, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		v.beginReimage.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginReimage) {
		v.beginReimage.remove(req)
	}

	return resp, nil
}

func (v *VirtualMachinesServerTransport) dispatchBeginRestart(req *http.Request) (*http.Response, error) {
	if v.srv.BeginRestart == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginRestart not implemented")}
	}
	beginRestart := v.beginRestart.get(req)
	if beginRestart == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/virtualMachines/(?P<virtualMachineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/restart`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualMachineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualMachineName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginRestart(req.Context(), resourceGroupNameParam, virtualMachineNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginRestart = &respr
		v.beginRestart.add(req, beginRestart)
	}

	resp, err := server.PollerResponderNext(beginRestart, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		v.beginRestart.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginRestart) {
		v.beginRestart.remove(req)
	}

	return resp, nil
}

func (v *VirtualMachinesServerTransport) dispatchBeginStart(req *http.Request) (*http.Response, error) {
	if v.srv.BeginStart == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginStart not implemented")}
	}
	beginStart := v.beginStart.get(req)
	if beginStart == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/virtualMachines/(?P<virtualMachineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/start`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualMachineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualMachineName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginStart(req.Context(), resourceGroupNameParam, virtualMachineNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginStart = &respr
		v.beginStart.add(req, beginStart)
	}

	resp, err := server.PollerResponderNext(beginStart, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		v.beginStart.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginStart) {
		v.beginStart.remove(req)
	}

	return resp, nil
}

func (v *VirtualMachinesServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := v.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/virtualMachines/(?P<virtualMachineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetworkcloud.VirtualMachinePatchParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualMachineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualMachineName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginUpdate(req.Context(), resourceGroupNameParam, virtualMachineNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		v.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		v.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		v.beginUpdate.remove(req)
	}

	return resp, nil
}
