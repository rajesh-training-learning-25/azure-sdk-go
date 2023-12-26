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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v3"
	"net/http"
	"net/url"
	"regexp"
)

// AccessPolicyAssignmentServer is a fake server for instances of the armredis.AccessPolicyAssignmentClient type.
type AccessPolicyAssignmentServer struct {
	// BeginCreateUpdate is the fake for method AccessPolicyAssignmentClient.BeginCreateUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateUpdate func(ctx context.Context, resourceGroupName string, cacheName string, accessPolicyAssignmentName string, parameters armredis.CacheAccessPolicyAssignment, options *armredis.AccessPolicyAssignmentClientBeginCreateUpdateOptions) (resp azfake.PollerResponder[armredis.AccessPolicyAssignmentClientCreateUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method AccessPolicyAssignmentClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, cacheName string, accessPolicyAssignmentName string, options *armredis.AccessPolicyAssignmentClientBeginDeleteOptions) (resp azfake.PollerResponder[armredis.AccessPolicyAssignmentClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method AccessPolicyAssignmentClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, cacheName string, accessPolicyAssignmentName string, options *armredis.AccessPolicyAssignmentClientGetOptions) (resp azfake.Responder[armredis.AccessPolicyAssignmentClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method AccessPolicyAssignmentClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, cacheName string, options *armredis.AccessPolicyAssignmentClientListOptions) (resp azfake.PagerResponder[armredis.AccessPolicyAssignmentClientListResponse])
}

// NewAccessPolicyAssignmentServerTransport creates a new instance of AccessPolicyAssignmentServerTransport with the provided implementation.
// The returned AccessPolicyAssignmentServerTransport instance is connected to an instance of armredis.AccessPolicyAssignmentClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAccessPolicyAssignmentServerTransport(srv *AccessPolicyAssignmentServer) *AccessPolicyAssignmentServerTransport {
	return &AccessPolicyAssignmentServerTransport{
		srv:               srv,
		beginCreateUpdate: newTracker[azfake.PollerResponder[armredis.AccessPolicyAssignmentClientCreateUpdateResponse]](),
		beginDelete:       newTracker[azfake.PollerResponder[armredis.AccessPolicyAssignmentClientDeleteResponse]](),
		newListPager:      newTracker[azfake.PagerResponder[armredis.AccessPolicyAssignmentClientListResponse]](),
	}
}

// AccessPolicyAssignmentServerTransport connects instances of armredis.AccessPolicyAssignmentClient to instances of AccessPolicyAssignmentServer.
// Don't use this type directly, use NewAccessPolicyAssignmentServerTransport instead.
type AccessPolicyAssignmentServerTransport struct {
	srv               *AccessPolicyAssignmentServer
	beginCreateUpdate *tracker[azfake.PollerResponder[armredis.AccessPolicyAssignmentClientCreateUpdateResponse]]
	beginDelete       *tracker[azfake.PollerResponder[armredis.AccessPolicyAssignmentClientDeleteResponse]]
	newListPager      *tracker[azfake.PagerResponder[armredis.AccessPolicyAssignmentClientListResponse]]
}

// Do implements the policy.Transporter interface for AccessPolicyAssignmentServerTransport.
func (a *AccessPolicyAssignmentServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AccessPolicyAssignmentClient.BeginCreateUpdate":
		resp, err = a.dispatchBeginCreateUpdate(req)
	case "AccessPolicyAssignmentClient.BeginDelete":
		resp, err = a.dispatchBeginDelete(req)
	case "AccessPolicyAssignmentClient.Get":
		resp, err = a.dispatchGet(req)
	case "AccessPolicyAssignmentClient.NewListPager":
		resp, err = a.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AccessPolicyAssignmentServerTransport) dispatchBeginCreateUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.BeginCreateUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateUpdate not implemented")}
	}
	beginCreateUpdate := a.beginCreateUpdate.get(req)
	if beginCreateUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cache/redis/(?P<cacheName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/accessPolicyAssignments/(?P<accessPolicyAssignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armredis.CacheAccessPolicyAssignment](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		cacheNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cacheName")])
		if err != nil {
			return nil, err
		}
		accessPolicyAssignmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accessPolicyAssignmentName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginCreateUpdate(req.Context(), resourceGroupNameParam, cacheNameParam, accessPolicyAssignmentNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateUpdate = &respr
		a.beginCreateUpdate.add(req, beginCreateUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		a.beginCreateUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateUpdate) {
		a.beginCreateUpdate.remove(req)
	}

	return resp, nil
}

func (a *AccessPolicyAssignmentServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if a.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := a.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cache/redis/(?P<cacheName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/accessPolicyAssignments/(?P<accessPolicyAssignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		cacheNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cacheName")])
		if err != nil {
			return nil, err
		}
		accessPolicyAssignmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accessPolicyAssignmentName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginDelete(req.Context(), resourceGroupNameParam, cacheNameParam, accessPolicyAssignmentNameParam, nil)
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

func (a *AccessPolicyAssignmentServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cache/redis/(?P<cacheName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/accessPolicyAssignments/(?P<accessPolicyAssignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	cacheNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cacheName")])
	if err != nil {
		return nil, err
	}
	accessPolicyAssignmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accessPolicyAssignmentName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, cacheNameParam, accessPolicyAssignmentNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CacheAccessPolicyAssignment, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AccessPolicyAssignmentServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := a.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cache/redis/(?P<cacheName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/accessPolicyAssignments`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		cacheNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cacheName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListPager(resourceGroupNameParam, cacheNameParam, nil)
		newListPager = &resp
		a.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armredis.AccessPolicyAssignmentClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		a.newListPager.remove(req)
	}
	return resp, nil
}
