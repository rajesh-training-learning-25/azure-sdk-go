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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
	"net/http"
	"net/url"
	"regexp"
)

// RestorableDroppedDatabasesServer is a fake server for instances of the armsql.RestorableDroppedDatabasesClient type.
type RestorableDroppedDatabasesServer struct {
	// Get is the fake for method RestorableDroppedDatabasesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serverName string, restorableDroppedDatabaseID string, options *armsql.RestorableDroppedDatabasesClientGetOptions) (resp azfake.Responder[armsql.RestorableDroppedDatabasesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByServerPager is the fake for method RestorableDroppedDatabasesClient.NewListByServerPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByServerPager func(resourceGroupName string, serverName string, options *armsql.RestorableDroppedDatabasesClientListByServerOptions) (resp azfake.PagerResponder[armsql.RestorableDroppedDatabasesClientListByServerResponse])
}

// NewRestorableDroppedDatabasesServerTransport creates a new instance of RestorableDroppedDatabasesServerTransport with the provided implementation.
// The returned RestorableDroppedDatabasesServerTransport instance is connected to an instance of armsql.RestorableDroppedDatabasesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRestorableDroppedDatabasesServerTransport(srv *RestorableDroppedDatabasesServer) *RestorableDroppedDatabasesServerTransport {
	return &RestorableDroppedDatabasesServerTransport{
		srv:                  srv,
		newListByServerPager: newTracker[azfake.PagerResponder[armsql.RestorableDroppedDatabasesClientListByServerResponse]](),
	}
}

// RestorableDroppedDatabasesServerTransport connects instances of armsql.RestorableDroppedDatabasesClient to instances of RestorableDroppedDatabasesServer.
// Don't use this type directly, use NewRestorableDroppedDatabasesServerTransport instead.
type RestorableDroppedDatabasesServerTransport struct {
	srv                  *RestorableDroppedDatabasesServer
	newListByServerPager *tracker[azfake.PagerResponder[armsql.RestorableDroppedDatabasesClientListByServerResponse]]
}

// Do implements the policy.Transporter interface for RestorableDroppedDatabasesServerTransport.
func (r *RestorableDroppedDatabasesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "RestorableDroppedDatabasesClient.Get":
		resp, err = r.dispatchGet(req)
	case "RestorableDroppedDatabasesClient.NewListByServerPager":
		resp, err = r.dispatchNewListByServerPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *RestorableDroppedDatabasesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/restorableDroppedDatabases/(?P<restorableDroppedDatabaseId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
	if err != nil {
		return nil, err
	}
	restorableDroppedDatabaseIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("restorableDroppedDatabaseId")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
	if err != nil {
		return nil, err
	}
	filterParam := getOptional(filterUnescaped)
	var options *armsql.RestorableDroppedDatabasesClientGetOptions
	if expandParam != nil || filterParam != nil {
		options = &armsql.RestorableDroppedDatabasesClientGetOptions{
			Expand: expandParam,
			Filter: filterParam,
		}
	}
	respr, errRespr := r.srv.Get(req.Context(), resourceGroupNameParam, serverNameParam, restorableDroppedDatabaseIDParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RestorableDroppedDatabase, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RestorableDroppedDatabasesServerTransport) dispatchNewListByServerPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListByServerPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByServerPager not implemented")}
	}
	newListByServerPager := r.newListByServerPager.get(req)
	if newListByServerPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/restorableDroppedDatabases`
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
		resp := r.srv.NewListByServerPager(resourceGroupNameParam, serverNameParam, nil)
		newListByServerPager = &resp
		r.newListByServerPager.add(req, newListByServerPager)
		server.PagerResponderInjectNextLinks(newListByServerPager, req, func(page *armsql.RestorableDroppedDatabasesClientListByServerResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByServerPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListByServerPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByServerPager) {
		r.newListByServerPager.remove(req)
	}
	return resp, nil
}
