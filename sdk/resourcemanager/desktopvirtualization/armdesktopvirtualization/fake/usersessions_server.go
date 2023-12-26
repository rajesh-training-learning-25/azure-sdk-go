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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization/v2"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
)

// UserSessionsServer is a fake server for instances of the armdesktopvirtualization.UserSessionsClient type.
type UserSessionsServer struct {
	// Delete is the fake for method UserSessionsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string, userSessionID string, options *armdesktopvirtualization.UserSessionsClientDeleteOptions) (resp azfake.Responder[armdesktopvirtualization.UserSessionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Disconnect is the fake for method UserSessionsClient.Disconnect
	// HTTP status codes to indicate success: http.StatusOK
	Disconnect func(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string, userSessionID string, options *armdesktopvirtualization.UserSessionsClientDisconnectOptions) (resp azfake.Responder[armdesktopvirtualization.UserSessionsClientDisconnectResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method UserSessionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string, userSessionID string, options *armdesktopvirtualization.UserSessionsClientGetOptions) (resp azfake.Responder[armdesktopvirtualization.UserSessionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method UserSessionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, hostPoolName string, sessionHostName string, options *armdesktopvirtualization.UserSessionsClientListOptions) (resp azfake.PagerResponder[armdesktopvirtualization.UserSessionsClientListResponse])

	// NewListByHostPoolPager is the fake for method UserSessionsClient.NewListByHostPoolPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByHostPoolPager func(resourceGroupName string, hostPoolName string, options *armdesktopvirtualization.UserSessionsClientListByHostPoolOptions) (resp azfake.PagerResponder[armdesktopvirtualization.UserSessionsClientListByHostPoolResponse])

	// SendMessage is the fake for method UserSessionsClient.SendMessage
	// HTTP status codes to indicate success: http.StatusOK
	SendMessage func(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string, userSessionID string, options *armdesktopvirtualization.UserSessionsClientSendMessageOptions) (resp azfake.Responder[armdesktopvirtualization.UserSessionsClientSendMessageResponse], errResp azfake.ErrorResponder)
}

// NewUserSessionsServerTransport creates a new instance of UserSessionsServerTransport with the provided implementation.
// The returned UserSessionsServerTransport instance is connected to an instance of armdesktopvirtualization.UserSessionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewUserSessionsServerTransport(srv *UserSessionsServer) *UserSessionsServerTransport {
	return &UserSessionsServerTransport{
		srv:                    srv,
		newListPager:           newTracker[azfake.PagerResponder[armdesktopvirtualization.UserSessionsClientListResponse]](),
		newListByHostPoolPager: newTracker[azfake.PagerResponder[armdesktopvirtualization.UserSessionsClientListByHostPoolResponse]](),
	}
}

// UserSessionsServerTransport connects instances of armdesktopvirtualization.UserSessionsClient to instances of UserSessionsServer.
// Don't use this type directly, use NewUserSessionsServerTransport instead.
type UserSessionsServerTransport struct {
	srv                    *UserSessionsServer
	newListPager           *tracker[azfake.PagerResponder[armdesktopvirtualization.UserSessionsClientListResponse]]
	newListByHostPoolPager *tracker[azfake.PagerResponder[armdesktopvirtualization.UserSessionsClientListByHostPoolResponse]]
}

// Do implements the policy.Transporter interface for UserSessionsServerTransport.
func (u *UserSessionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "UserSessionsClient.Delete":
		resp, err = u.dispatchDelete(req)
	case "UserSessionsClient.Disconnect":
		resp, err = u.dispatchDisconnect(req)
	case "UserSessionsClient.Get":
		resp, err = u.dispatchGet(req)
	case "UserSessionsClient.NewListPager":
		resp, err = u.dispatchNewListPager(req)
	case "UserSessionsClient.NewListByHostPoolPager":
		resp, err = u.dispatchNewListByHostPoolPager(req)
	case "UserSessionsClient.SendMessage":
		resp, err = u.dispatchSendMessage(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (u *UserSessionsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if u.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DesktopVirtualization/hostPools/(?P<hostPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sessionHosts/(?P<sessionHostName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/userSessions/(?P<userSessionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	hostPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hostPoolName")])
	if err != nil {
		return nil, err
	}
	sessionHostNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sessionHostName")])
	if err != nil {
		return nil, err
	}
	userSessionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("userSessionId")])
	if err != nil {
		return nil, err
	}
	forceUnescaped, err := url.QueryUnescape(qp.Get("force"))
	if err != nil {
		return nil, err
	}
	forceParam, err := parseOptional(forceUnescaped, strconv.ParseBool)
	if err != nil {
		return nil, err
	}
	var options *armdesktopvirtualization.UserSessionsClientDeleteOptions
	if forceParam != nil {
		options = &armdesktopvirtualization.UserSessionsClientDeleteOptions{
			Force: forceParam,
		}
	}
	respr, errRespr := u.srv.Delete(req.Context(), resourceGroupNameParam, hostPoolNameParam, sessionHostNameParam, userSessionIDParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *UserSessionsServerTransport) dispatchDisconnect(req *http.Request) (*http.Response, error) {
	if u.srv.Disconnect == nil {
		return nil, &nonRetriableError{errors.New("fake for method Disconnect not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DesktopVirtualization/hostPools/(?P<hostPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sessionHosts/(?P<sessionHostName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/userSessions/(?P<userSessionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/disconnect`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	hostPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hostPoolName")])
	if err != nil {
		return nil, err
	}
	sessionHostNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sessionHostName")])
	if err != nil {
		return nil, err
	}
	userSessionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("userSessionId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := u.srv.Disconnect(req.Context(), resourceGroupNameParam, hostPoolNameParam, sessionHostNameParam, userSessionIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *UserSessionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if u.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DesktopVirtualization/hostPools/(?P<hostPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sessionHosts/(?P<sessionHostName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/userSessions/(?P<userSessionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	hostPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hostPoolName")])
	if err != nil {
		return nil, err
	}
	sessionHostNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sessionHostName")])
	if err != nil {
		return nil, err
	}
	userSessionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("userSessionId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := u.srv.Get(req.Context(), resourceGroupNameParam, hostPoolNameParam, sessionHostNameParam, userSessionIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).UserSession, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *UserSessionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if u.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := u.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DesktopVirtualization/hostPools/(?P<hostPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sessionHosts/(?P<sessionHostName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/userSessions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		hostPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hostPoolName")])
		if err != nil {
			return nil, err
		}
		sessionHostNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sessionHostName")])
		if err != nil {
			return nil, err
		}
		pageSizeUnescaped, err := url.QueryUnescape(qp.Get("pageSize"))
		if err != nil {
			return nil, err
		}
		pageSizeParam, err := parseOptional(pageSizeUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		isDescendingUnescaped, err := url.QueryUnescape(qp.Get("isDescending"))
		if err != nil {
			return nil, err
		}
		isDescendingParam, err := parseOptional(isDescendingUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		initialSkipUnescaped, err := url.QueryUnescape(qp.Get("initialSkip"))
		if err != nil {
			return nil, err
		}
		initialSkipParam, err := parseOptional(initialSkipUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *armdesktopvirtualization.UserSessionsClientListOptions
		if pageSizeParam != nil || isDescendingParam != nil || initialSkipParam != nil {
			options = &armdesktopvirtualization.UserSessionsClientListOptions{
				PageSize:     pageSizeParam,
				IsDescending: isDescendingParam,
				InitialSkip:  initialSkipParam,
			}
		}
		resp := u.srv.NewListPager(resourceGroupNameParam, hostPoolNameParam, sessionHostNameParam, options)
		newListPager = &resp
		u.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armdesktopvirtualization.UserSessionsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		u.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		u.newListPager.remove(req)
	}
	return resp, nil
}

func (u *UserSessionsServerTransport) dispatchNewListByHostPoolPager(req *http.Request) (*http.Response, error) {
	if u.srv.NewListByHostPoolPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByHostPoolPager not implemented")}
	}
	newListByHostPoolPager := u.newListByHostPoolPager.get(req)
	if newListByHostPoolPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DesktopVirtualization/hostPools/(?P<hostPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/userSessions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		hostPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hostPoolName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		pageSizeUnescaped, err := url.QueryUnescape(qp.Get("pageSize"))
		if err != nil {
			return nil, err
		}
		pageSizeParam, err := parseOptional(pageSizeUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		isDescendingUnescaped, err := url.QueryUnescape(qp.Get("isDescending"))
		if err != nil {
			return nil, err
		}
		isDescendingParam, err := parseOptional(isDescendingUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		initialSkipUnescaped, err := url.QueryUnescape(qp.Get("initialSkip"))
		if err != nil {
			return nil, err
		}
		initialSkipParam, err := parseOptional(initialSkipUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *armdesktopvirtualization.UserSessionsClientListByHostPoolOptions
		if filterParam != nil || pageSizeParam != nil || isDescendingParam != nil || initialSkipParam != nil {
			options = &armdesktopvirtualization.UserSessionsClientListByHostPoolOptions{
				Filter:       filterParam,
				PageSize:     pageSizeParam,
				IsDescending: isDescendingParam,
				InitialSkip:  initialSkipParam,
			}
		}
		resp := u.srv.NewListByHostPoolPager(resourceGroupNameParam, hostPoolNameParam, options)
		newListByHostPoolPager = &resp
		u.newListByHostPoolPager.add(req, newListByHostPoolPager)
		server.PagerResponderInjectNextLinks(newListByHostPoolPager, req, func(page *armdesktopvirtualization.UserSessionsClientListByHostPoolResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByHostPoolPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		u.newListByHostPoolPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByHostPoolPager) {
		u.newListByHostPoolPager.remove(req)
	}
	return resp, nil
}

func (u *UserSessionsServerTransport) dispatchSendMessage(req *http.Request) (*http.Response, error) {
	if u.srv.SendMessage == nil {
		return nil, &nonRetriableError{errors.New("fake for method SendMessage not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DesktopVirtualization/hostPools/(?P<hostPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sessionHosts/(?P<sessionHostName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/userSessions/(?P<userSessionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sendMessage`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdesktopvirtualization.SendMessage](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	hostPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hostPoolName")])
	if err != nil {
		return nil, err
	}
	sessionHostNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sessionHostName")])
	if err != nil {
		return nil, err
	}
	userSessionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("userSessionId")])
	if err != nil {
		return nil, err
	}
	var options *armdesktopvirtualization.UserSessionsClientSendMessageOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &armdesktopvirtualization.UserSessionsClientSendMessageOptions{
			SendMessage: &body,
		}
	}
	respr, errRespr := u.srv.SendMessage(req.Context(), resourceGroupNameParam, hostPoolNameParam, sessionHostNameParam, userSessionIDParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
