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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
	"net/http"
	"net/url"
	"regexp"
)

// AccessReviewInstancesServer is a fake server for instances of the armauthorization.AccessReviewInstancesClient type.
type AccessReviewInstancesServer struct {
	// Create is the fake for method AccessReviewInstancesClient.Create
	// HTTP status codes to indicate success: http.StatusOK
	Create func(ctx context.Context, scheduleDefinitionID string, id string, properties armauthorization.AccessReviewInstanceProperties, options *armauthorization.AccessReviewInstancesClientCreateOptions) (resp azfake.Responder[armauthorization.AccessReviewInstancesClientCreateResponse], errResp azfake.ErrorResponder)

	// GetByID is the fake for method AccessReviewInstancesClient.GetByID
	// HTTP status codes to indicate success: http.StatusOK
	GetByID func(ctx context.Context, scheduleDefinitionID string, id string, options *armauthorization.AccessReviewInstancesClientGetByIDOptions) (resp azfake.Responder[armauthorization.AccessReviewInstancesClientGetByIDResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method AccessReviewInstancesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(scheduleDefinitionID string, options *armauthorization.AccessReviewInstancesClientListOptions) (resp azfake.PagerResponder[armauthorization.AccessReviewInstancesClientListResponse])
}

// NewAccessReviewInstancesServerTransport creates a new instance of AccessReviewInstancesServerTransport with the provided implementation.
// The returned AccessReviewInstancesServerTransport instance is connected to an instance of armauthorization.AccessReviewInstancesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAccessReviewInstancesServerTransport(srv *AccessReviewInstancesServer) *AccessReviewInstancesServerTransport {
	return &AccessReviewInstancesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armauthorization.AccessReviewInstancesClientListResponse]](),
	}
}

// AccessReviewInstancesServerTransport connects instances of armauthorization.AccessReviewInstancesClient to instances of AccessReviewInstancesServer.
// Don't use this type directly, use NewAccessReviewInstancesServerTransport instead.
type AccessReviewInstancesServerTransport struct {
	srv          *AccessReviewInstancesServer
	newListPager *tracker[azfake.PagerResponder[armauthorization.AccessReviewInstancesClientListResponse]]
}

// Do implements the policy.Transporter interface for AccessReviewInstancesServerTransport.
func (a *AccessReviewInstancesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AccessReviewInstancesClient.Create":
		resp, err = a.dispatchCreate(req)
	case "AccessReviewInstancesClient.GetByID":
		resp, err = a.dispatchGetByID(req)
	case "AccessReviewInstancesClient.NewListPager":
		resp, err = a.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AccessReviewInstancesServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if a.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/accessReviewScheduleDefinitions/(?P<scheduleDefinitionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/instances/(?P<id>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armauthorization.AccessReviewInstanceProperties](req)
	if err != nil {
		return nil, err
	}
	scheduleDefinitionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("scheduleDefinitionId")])
	if err != nil {
		return nil, err
	}
	idParam, err := url.PathUnescape(matches[regex.SubexpIndex("id")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Create(req.Context(), scheduleDefinitionIDParam, idParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AccessReviewInstance, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AccessReviewInstancesServerTransport) dispatchGetByID(req *http.Request) (*http.Response, error) {
	if a.srv.GetByID == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetByID not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/accessReviewScheduleDefinitions/(?P<scheduleDefinitionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/instances/(?P<id>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	scheduleDefinitionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("scheduleDefinitionId")])
	if err != nil {
		return nil, err
	}
	idParam, err := url.PathUnescape(matches[regex.SubexpIndex("id")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.GetByID(req.Context(), scheduleDefinitionIDParam, idParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AccessReviewInstance, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AccessReviewInstancesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := a.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/accessReviewScheduleDefinitions/(?P<scheduleDefinitionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/instances`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		scheduleDefinitionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("scheduleDefinitionId")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armauthorization.AccessReviewInstancesClientListOptions
		if filterParam != nil {
			options = &armauthorization.AccessReviewInstancesClientListOptions{
				Filter: filterParam,
			}
		}
		resp := a.srv.NewListPager(scheduleDefinitionIDParam, options)
		newListPager = &resp
		a.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armauthorization.AccessReviewInstancesClientListResponse, createLink func() string) {
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
