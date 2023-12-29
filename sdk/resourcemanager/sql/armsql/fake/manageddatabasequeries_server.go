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

// ManagedDatabaseQueriesServer is a fake server for instances of the armsql.ManagedDatabaseQueriesClient type.
type ManagedDatabaseQueriesServer struct {
	// Get is the fake for method ManagedDatabaseQueriesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, queryID string, options *armsql.ManagedDatabaseQueriesClientGetOptions) (resp azfake.Responder[armsql.ManagedDatabaseQueriesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByQueryPager is the fake for method ManagedDatabaseQueriesClient.NewListByQueryPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByQueryPager func(resourceGroupName string, managedInstanceName string, databaseName string, queryID string, options *armsql.ManagedDatabaseQueriesClientListByQueryOptions) (resp azfake.PagerResponder[armsql.ManagedDatabaseQueriesClientListByQueryResponse])
}

// NewManagedDatabaseQueriesServerTransport creates a new instance of ManagedDatabaseQueriesServerTransport with the provided implementation.
// The returned ManagedDatabaseQueriesServerTransport instance is connected to an instance of armsql.ManagedDatabaseQueriesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewManagedDatabaseQueriesServerTransport(srv *ManagedDatabaseQueriesServer) *ManagedDatabaseQueriesServerTransport {
	return &ManagedDatabaseQueriesServerTransport{
		srv:                 srv,
		newListByQueryPager: newTracker[azfake.PagerResponder[armsql.ManagedDatabaseQueriesClientListByQueryResponse]](),
	}
}

// ManagedDatabaseQueriesServerTransport connects instances of armsql.ManagedDatabaseQueriesClient to instances of ManagedDatabaseQueriesServer.
// Don't use this type directly, use NewManagedDatabaseQueriesServerTransport instead.
type ManagedDatabaseQueriesServerTransport struct {
	srv                 *ManagedDatabaseQueriesServer
	newListByQueryPager *tracker[azfake.PagerResponder[armsql.ManagedDatabaseQueriesClientListByQueryResponse]]
}

// Do implements the policy.Transporter interface for ManagedDatabaseQueriesServerTransport.
func (m *ManagedDatabaseQueriesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ManagedDatabaseQueriesClient.Get":
		resp, err = m.dispatchGet(req)
	case "ManagedDatabaseQueriesClient.NewListByQueryPager":
		resp, err = m.dispatchNewListByQueryPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *ManagedDatabaseQueriesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if m.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/queries/(?P<queryId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
	if err != nil {
		return nil, err
	}
	databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
	if err != nil {
		return nil, err
	}
	queryIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("queryId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Get(req.Context(), resourceGroupNameParam, managedInstanceNameParam, databaseNameParam, queryIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ManagedInstanceQuery, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagedDatabaseQueriesServerTransport) dispatchNewListByQueryPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListByQueryPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByQueryPager not implemented")}
	}
	newListByQueryPager := m.newListByQueryPager.get(req)
	if newListByQueryPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/queries/(?P<queryId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/statistics`
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
		managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
		if err != nil {
			return nil, err
		}
		databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
		if err != nil {
			return nil, err
		}
		queryIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("queryId")])
		if err != nil {
			return nil, err
		}
		startTimeUnescaped, err := url.QueryUnescape(qp.Get("startTime"))
		if err != nil {
			return nil, err
		}
		startTimeParam := getOptional(startTimeUnescaped)
		endTimeUnescaped, err := url.QueryUnescape(qp.Get("endTime"))
		if err != nil {
			return nil, err
		}
		endTimeParam := getOptional(endTimeUnescaped)
		intervalUnescaped, err := url.QueryUnescape(qp.Get("interval"))
		if err != nil {
			return nil, err
		}
		intervalParam := getOptional(armsql.QueryTimeGrainType(intervalUnescaped))
		var options *armsql.ManagedDatabaseQueriesClientListByQueryOptions
		if startTimeParam != nil || endTimeParam != nil || intervalParam != nil {
			options = &armsql.ManagedDatabaseQueriesClientListByQueryOptions{
				StartTime: startTimeParam,
				EndTime:   endTimeParam,
				Interval:  intervalParam,
			}
		}
		resp := m.srv.NewListByQueryPager(resourceGroupNameParam, managedInstanceNameParam, databaseNameParam, queryIDParam, options)
		newListByQueryPager = &resp
		m.newListByQueryPager.add(req, newListByQueryPager)
		server.PagerResponderInjectNextLinks(newListByQueryPager, req, func(page *armsql.ManagedDatabaseQueriesClientListByQueryResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByQueryPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListByQueryPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByQueryPager) {
		m.newListByQueryPager.remove(req)
	}
	return resp, nil
}
