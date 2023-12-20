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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v2"
	"net/http"
	"net/url"
	"regexp"
)

// AFDOriginGroupsServer is a fake server for instances of the armcdn.AFDOriginGroupsClient type.
type AFDOriginGroupsServer struct {
	// BeginCreate is the fake for method AFDOriginGroupsClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreate func(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, originGroup armcdn.AFDOriginGroup, options *armcdn.AFDOriginGroupsClientBeginCreateOptions) (resp azfake.PollerResponder[armcdn.AFDOriginGroupsClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method AFDOriginGroupsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, options *armcdn.AFDOriginGroupsClientBeginDeleteOptions) (resp azfake.PollerResponder[armcdn.AFDOriginGroupsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method AFDOriginGroupsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, options *armcdn.AFDOriginGroupsClientGetOptions) (resp azfake.Responder[armcdn.AFDOriginGroupsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByProfilePager is the fake for method AFDOriginGroupsClient.NewListByProfilePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByProfilePager func(resourceGroupName string, profileName string, options *armcdn.AFDOriginGroupsClientListByProfileOptions) (resp azfake.PagerResponder[armcdn.AFDOriginGroupsClientListByProfileResponse])

	// NewListResourceUsagePager is the fake for method AFDOriginGroupsClient.NewListResourceUsagePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListResourceUsagePager func(resourceGroupName string, profileName string, originGroupName string, options *armcdn.AFDOriginGroupsClientListResourceUsageOptions) (resp azfake.PagerResponder[armcdn.AFDOriginGroupsClientListResourceUsageResponse])

	// BeginUpdate is the fake for method AFDOriginGroupsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, originGroupUpdateProperties armcdn.AFDOriginGroupUpdateParameters, options *armcdn.AFDOriginGroupsClientBeginUpdateOptions) (resp azfake.PollerResponder[armcdn.AFDOriginGroupsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewAFDOriginGroupsServerTransport creates a new instance of AFDOriginGroupsServerTransport with the provided implementation.
// The returned AFDOriginGroupsServerTransport instance is connected to an instance of armcdn.AFDOriginGroupsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAFDOriginGroupsServerTransport(srv *AFDOriginGroupsServer) *AFDOriginGroupsServerTransport {
	return &AFDOriginGroupsServerTransport{
		srv:                       srv,
		beginCreate:               newTracker[azfake.PollerResponder[armcdn.AFDOriginGroupsClientCreateResponse]](),
		beginDelete:               newTracker[azfake.PollerResponder[armcdn.AFDOriginGroupsClientDeleteResponse]](),
		newListByProfilePager:     newTracker[azfake.PagerResponder[armcdn.AFDOriginGroupsClientListByProfileResponse]](),
		newListResourceUsagePager: newTracker[azfake.PagerResponder[armcdn.AFDOriginGroupsClientListResourceUsageResponse]](),
		beginUpdate:               newTracker[azfake.PollerResponder[armcdn.AFDOriginGroupsClientUpdateResponse]](),
	}
}

// AFDOriginGroupsServerTransport connects instances of armcdn.AFDOriginGroupsClient to instances of AFDOriginGroupsServer.
// Don't use this type directly, use NewAFDOriginGroupsServerTransport instead.
type AFDOriginGroupsServerTransport struct {
	srv                       *AFDOriginGroupsServer
	beginCreate               *tracker[azfake.PollerResponder[armcdn.AFDOriginGroupsClientCreateResponse]]
	beginDelete               *tracker[azfake.PollerResponder[armcdn.AFDOriginGroupsClientDeleteResponse]]
	newListByProfilePager     *tracker[azfake.PagerResponder[armcdn.AFDOriginGroupsClientListByProfileResponse]]
	newListResourceUsagePager *tracker[azfake.PagerResponder[armcdn.AFDOriginGroupsClientListResourceUsageResponse]]
	beginUpdate               *tracker[azfake.PollerResponder[armcdn.AFDOriginGroupsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for AFDOriginGroupsServerTransport.
func (a *AFDOriginGroupsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AFDOriginGroupsClient.BeginCreate":
		resp, err = a.dispatchBeginCreate(req)
	case "AFDOriginGroupsClient.BeginDelete":
		resp, err = a.dispatchBeginDelete(req)
	case "AFDOriginGroupsClient.Get":
		resp, err = a.dispatchGet(req)
	case "AFDOriginGroupsClient.NewListByProfilePager":
		resp, err = a.dispatchNewListByProfilePager(req)
	case "AFDOriginGroupsClient.NewListResourceUsagePager":
		resp, err = a.dispatchNewListResourceUsagePager(req)
	case "AFDOriginGroupsClient.BeginUpdate":
		resp, err = a.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AFDOriginGroupsServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if a.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := a.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cdn/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/originGroups/(?P<originGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcdn.AFDOriginGroup](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
		if err != nil {
			return nil, err
		}
		originGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("originGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginCreate(req.Context(), resourceGroupNameParam, profileNameParam, originGroupNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		a.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		a.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		a.beginCreate.remove(req)
	}

	return resp, nil
}

func (a *AFDOriginGroupsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if a.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := a.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cdn/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/originGroups/(?P<originGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
		if err != nil {
			return nil, err
		}
		originGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("originGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginDelete(req.Context(), resourceGroupNameParam, profileNameParam, originGroupNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		a.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		a.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		a.beginDelete.remove(req)
	}

	return resp, nil
}

func (a *AFDOriginGroupsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cdn/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/originGroups/(?P<originGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
	if err != nil {
		return nil, err
	}
	originGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("originGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, profileNameParam, originGroupNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AFDOriginGroup, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AFDOriginGroupsServerTransport) dispatchNewListByProfilePager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListByProfilePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByProfilePager not implemented")}
	}
	newListByProfilePager := a.newListByProfilePager.get(req)
	if newListByProfilePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cdn/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/originGroups`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListByProfilePager(resourceGroupNameParam, profileNameParam, nil)
		newListByProfilePager = &resp
		a.newListByProfilePager.add(req, newListByProfilePager)
		server.PagerResponderInjectNextLinks(newListByProfilePager, req, func(page *armcdn.AFDOriginGroupsClientListByProfileResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByProfilePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListByProfilePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByProfilePager) {
		a.newListByProfilePager.remove(req)
	}
	return resp, nil
}

func (a *AFDOriginGroupsServerTransport) dispatchNewListResourceUsagePager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListResourceUsagePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListResourceUsagePager not implemented")}
	}
	newListResourceUsagePager := a.newListResourceUsagePager.get(req)
	if newListResourceUsagePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cdn/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/originGroups/(?P<originGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/usages`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
		if err != nil {
			return nil, err
		}
		originGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("originGroupName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListResourceUsagePager(resourceGroupNameParam, profileNameParam, originGroupNameParam, nil)
		newListResourceUsagePager = &resp
		a.newListResourceUsagePager.add(req, newListResourceUsagePager)
		server.PagerResponderInjectNextLinks(newListResourceUsagePager, req, func(page *armcdn.AFDOriginGroupsClientListResourceUsageResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListResourceUsagePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListResourceUsagePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListResourceUsagePager) {
		a.newListResourceUsagePager.remove(req)
	}
	return resp, nil
}

func (a *AFDOriginGroupsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := a.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cdn/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/originGroups/(?P<originGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcdn.AFDOriginGroupUpdateParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
		if err != nil {
			return nil, err
		}
		originGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("originGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginUpdate(req.Context(), resourceGroupNameParam, profileNameParam, originGroupNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		a.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		a.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		a.beginUpdate.remove(req)
	}

	return resp, nil
}
