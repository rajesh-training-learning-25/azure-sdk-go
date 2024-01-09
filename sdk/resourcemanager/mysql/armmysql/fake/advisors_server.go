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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"
	"net/http"
	"net/url"
	"regexp"
)

// AdvisorsServer is a fake server for instances of the armmysql.AdvisorsClient type.
type AdvisorsServer struct {
	// Get is the fake for method AdvisorsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serverName string, advisorName string, options *armmysql.AdvisorsClientGetOptions) (resp azfake.Responder[armmysql.AdvisorsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByServerPager is the fake for method AdvisorsClient.NewListByServerPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByServerPager func(resourceGroupName string, serverName string, options *armmysql.AdvisorsClientListByServerOptions) (resp azfake.PagerResponder[armmysql.AdvisorsClientListByServerResponse])
}

// NewAdvisorsServerTransport creates a new instance of AdvisorsServerTransport with the provided implementation.
// The returned AdvisorsServerTransport instance is connected to an instance of armmysql.AdvisorsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAdvisorsServerTransport(srv *AdvisorsServer) *AdvisorsServerTransport {
	return &AdvisorsServerTransport{
		srv:                  srv,
		newListByServerPager: newTracker[azfake.PagerResponder[armmysql.AdvisorsClientListByServerResponse]](),
	}
}

// AdvisorsServerTransport connects instances of armmysql.AdvisorsClient to instances of AdvisorsServer.
// Don't use this type directly, use NewAdvisorsServerTransport instead.
type AdvisorsServerTransport struct {
	srv                  *AdvisorsServer
	newListByServerPager *tracker[azfake.PagerResponder[armmysql.AdvisorsClientListByServerResponse]]
}

// Do implements the policy.Transporter interface for AdvisorsServerTransport.
func (a *AdvisorsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AdvisorsClient.Get":
		resp, err = a.dispatchGet(req)
	case "AdvisorsClient.NewListByServerPager":
		resp, err = a.dispatchNewListByServerPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AdvisorsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforMySQL/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/advisors/(?P<advisorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
	if err != nil {
		return nil, err
	}
	advisorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("advisorName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, serverNameParam, advisorNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Advisor, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AdvisorsServerTransport) dispatchNewListByServerPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListByServerPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByServerPager not implemented")}
	}
	newListByServerPager := a.newListByServerPager.get(req)
	if newListByServerPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforMySQL/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/advisors`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListByServerPager(resourceGroupNameParam, serverNameParam, nil)
		newListByServerPager = &resp
		a.newListByServerPager.add(req, newListByServerPager)
		server.PagerResponderInjectNextLinks(newListByServerPager, req, func(page *armmysql.AdvisorsClientListByServerResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByServerPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListByServerPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByServerPager) {
		a.newListByServerPager.remove(req)
	}
	return resp, nil
}