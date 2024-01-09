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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch"
	"net/http"
	"net/url"
	"regexp"
)

// QueryKeysServer is a fake server for instances of the armsearch.QueryKeysClient type.
type QueryKeysServer struct {
	// Create is the fake for method QueryKeysClient.Create
	// HTTP status codes to indicate success: http.StatusOK
	Create func(ctx context.Context, resourceGroupName string, searchServiceName string, name string, searchManagementRequestOptions *armsearch.SearchManagementRequestOptions, options *armsearch.QueryKeysClientCreateOptions) (resp azfake.Responder[armsearch.QueryKeysClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method QueryKeysClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent, http.StatusNotFound
	Delete func(ctx context.Context, resourceGroupName string, searchServiceName string, key string, searchManagementRequestOptions *armsearch.SearchManagementRequestOptions, options *armsearch.QueryKeysClientDeleteOptions) (resp azfake.Responder[armsearch.QueryKeysClientDeleteResponse], errResp azfake.ErrorResponder)

	// NewListBySearchServicePager is the fake for method QueryKeysClient.NewListBySearchServicePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySearchServicePager func(resourceGroupName string, searchServiceName string, searchManagementRequestOptions *armsearch.SearchManagementRequestOptions, options *armsearch.QueryKeysClientListBySearchServiceOptions) (resp azfake.PagerResponder[armsearch.QueryKeysClientListBySearchServiceResponse])
}

// NewQueryKeysServerTransport creates a new instance of QueryKeysServerTransport with the provided implementation.
// The returned QueryKeysServerTransport instance is connected to an instance of armsearch.QueryKeysClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewQueryKeysServerTransport(srv *QueryKeysServer) *QueryKeysServerTransport {
	return &QueryKeysServerTransport{
		srv:                         srv,
		newListBySearchServicePager: newTracker[azfake.PagerResponder[armsearch.QueryKeysClientListBySearchServiceResponse]](),
	}
}

// QueryKeysServerTransport connects instances of armsearch.QueryKeysClient to instances of QueryKeysServer.
// Don't use this type directly, use NewQueryKeysServerTransport instead.
type QueryKeysServerTransport struct {
	srv                         *QueryKeysServer
	newListBySearchServicePager *tracker[azfake.PagerResponder[armsearch.QueryKeysClientListBySearchServiceResponse]]
}

// Do implements the policy.Transporter interface for QueryKeysServerTransport.
func (q *QueryKeysServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "QueryKeysClient.Create":
		resp, err = q.dispatchCreate(req)
	case "QueryKeysClient.Delete":
		resp, err = q.dispatchDelete(req)
	case "QueryKeysClient.NewListBySearchServicePager":
		resp, err = q.dispatchNewListBySearchServicePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (q *QueryKeysServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if q.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Search/searchServices/(?P<searchServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/createQueryKey/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	searchServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("searchServiceName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	clientRequestIDParam := getOptional(getHeaderValue(req.Header, "x-ms-client-request-id"))
	var searchManagementRequestOptions *armsearch.SearchManagementRequestOptions
	if clientRequestIDParam != nil {
		searchManagementRequestOptions = &armsearch.SearchManagementRequestOptions{
			ClientRequestID: clientRequestIDParam,
		}
	}
	respr, errRespr := q.srv.Create(req.Context(), resourceGroupNameParam, searchServiceNameParam, nameParam, searchManagementRequestOptions, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).QueryKey, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (q *QueryKeysServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if q.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Search/searchServices/(?P<searchServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deleteQueryKey/(?P<key>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	searchServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("searchServiceName")])
	if err != nil {
		return nil, err
	}
	keyParam, err := url.PathUnescape(matches[regex.SubexpIndex("key")])
	if err != nil {
		return nil, err
	}
	clientRequestIDParam := getOptional(getHeaderValue(req.Header, "x-ms-client-request-id"))
	var searchManagementRequestOptions *armsearch.SearchManagementRequestOptions
	if clientRequestIDParam != nil {
		searchManagementRequestOptions = &armsearch.SearchManagementRequestOptions{
			ClientRequestID: clientRequestIDParam,
		}
	}
	respr, errRespr := q.srv.Delete(req.Context(), resourceGroupNameParam, searchServiceNameParam, keyParam, searchManagementRequestOptions, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent, http.StatusNotFound}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent, http.StatusNotFound", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (q *QueryKeysServerTransport) dispatchNewListBySearchServicePager(req *http.Request) (*http.Response, error) {
	if q.srv.NewListBySearchServicePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySearchServicePager not implemented")}
	}
	newListBySearchServicePager := q.newListBySearchServicePager.get(req)
	if newListBySearchServicePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Search/searchServices/(?P<searchServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listQueryKeys`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		searchServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("searchServiceName")])
		if err != nil {
			return nil, err
		}
		clientRequestIDParam := getOptional(getHeaderValue(req.Header, "x-ms-client-request-id"))
		var searchManagementRequestOptions *armsearch.SearchManagementRequestOptions
		if clientRequestIDParam != nil {
			searchManagementRequestOptions = &armsearch.SearchManagementRequestOptions{
				ClientRequestID: clientRequestIDParam,
			}
		}
		resp := q.srv.NewListBySearchServicePager(resourceGroupNameParam, searchServiceNameParam, searchManagementRequestOptions, nil)
		newListBySearchServicePager = &resp
		q.newListBySearchServicePager.add(req, newListBySearchServicePager)
		server.PagerResponderInjectNextLinks(newListBySearchServicePager, req, func(page *armsearch.QueryKeysClientListBySearchServiceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySearchServicePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		q.newListBySearchServicePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySearchServicePager) {
		q.newListBySearchServicePager.remove(req)
	}
	return resp, nil
}