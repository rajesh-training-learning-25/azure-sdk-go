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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
	"net/http"
	"net/url"
	"regexp"
)

// WorkItemConfigurationsServer is a fake server for instances of the armapplicationinsights.WorkItemConfigurationsClient type.
type WorkItemConfigurationsServer struct {
	// Create is the fake for method WorkItemConfigurationsClient.Create
	// HTTP status codes to indicate success: http.StatusOK
	Create func(ctx context.Context, resourceGroupName string, resourceName string, workItemConfigurationProperties armapplicationinsights.WorkItemCreateConfiguration, options *armapplicationinsights.WorkItemConfigurationsClientCreateOptions) (resp azfake.Responder[armapplicationinsights.WorkItemConfigurationsClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method WorkItemConfigurationsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK
	Delete func(ctx context.Context, resourceGroupName string, resourceName string, workItemConfigID string, options *armapplicationinsights.WorkItemConfigurationsClientDeleteOptions) (resp azfake.Responder[armapplicationinsights.WorkItemConfigurationsClientDeleteResponse], errResp azfake.ErrorResponder)

	// GetDefault is the fake for method WorkItemConfigurationsClient.GetDefault
	// HTTP status codes to indicate success: http.StatusOK
	GetDefault func(ctx context.Context, resourceGroupName string, resourceName string, options *armapplicationinsights.WorkItemConfigurationsClientGetDefaultOptions) (resp azfake.Responder[armapplicationinsights.WorkItemConfigurationsClientGetDefaultResponse], errResp azfake.ErrorResponder)

	// GetItem is the fake for method WorkItemConfigurationsClient.GetItem
	// HTTP status codes to indicate success: http.StatusOK
	GetItem func(ctx context.Context, resourceGroupName string, resourceName string, workItemConfigID string, options *armapplicationinsights.WorkItemConfigurationsClientGetItemOptions) (resp azfake.Responder[armapplicationinsights.WorkItemConfigurationsClientGetItemResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method WorkItemConfigurationsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, resourceName string, options *armapplicationinsights.WorkItemConfigurationsClientListOptions) (resp azfake.PagerResponder[armapplicationinsights.WorkItemConfigurationsClientListResponse])

	// UpdateItem is the fake for method WorkItemConfigurationsClient.UpdateItem
	// HTTP status codes to indicate success: http.StatusOK
	UpdateItem func(ctx context.Context, resourceGroupName string, resourceName string, workItemConfigID string, workItemConfigurationProperties armapplicationinsights.WorkItemCreateConfiguration, options *armapplicationinsights.WorkItemConfigurationsClientUpdateItemOptions) (resp azfake.Responder[armapplicationinsights.WorkItemConfigurationsClientUpdateItemResponse], errResp azfake.ErrorResponder)
}

// NewWorkItemConfigurationsServerTransport creates a new instance of WorkItemConfigurationsServerTransport with the provided implementation.
// The returned WorkItemConfigurationsServerTransport instance is connected to an instance of armapplicationinsights.WorkItemConfigurationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewWorkItemConfigurationsServerTransport(srv *WorkItemConfigurationsServer) *WorkItemConfigurationsServerTransport {
	return &WorkItemConfigurationsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armapplicationinsights.WorkItemConfigurationsClientListResponse]](),
	}
}

// WorkItemConfigurationsServerTransport connects instances of armapplicationinsights.WorkItemConfigurationsClient to instances of WorkItemConfigurationsServer.
// Don't use this type directly, use NewWorkItemConfigurationsServerTransport instead.
type WorkItemConfigurationsServerTransport struct {
	srv          *WorkItemConfigurationsServer
	newListPager *tracker[azfake.PagerResponder[armapplicationinsights.WorkItemConfigurationsClientListResponse]]
}

// Do implements the policy.Transporter interface for WorkItemConfigurationsServerTransport.
func (w *WorkItemConfigurationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "WorkItemConfigurationsClient.Create":
		resp, err = w.dispatchCreate(req)
	case "WorkItemConfigurationsClient.Delete":
		resp, err = w.dispatchDelete(req)
	case "WorkItemConfigurationsClient.GetDefault":
		resp, err = w.dispatchGetDefault(req)
	case "WorkItemConfigurationsClient.GetItem":
		resp, err = w.dispatchGetItem(req)
	case "WorkItemConfigurationsClient.NewListPager":
		resp, err = w.dispatchNewListPager(req)
	case "WorkItemConfigurationsClient.UpdateItem":
		resp, err = w.dispatchUpdateItem(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w *WorkItemConfigurationsServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if w.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Insights/components/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/WorkItemConfigs`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armapplicationinsights.WorkItemCreateConfiguration](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Create(req.Context(), resourceGroupNameParam, resourceNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).WorkItemConfiguration, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WorkItemConfigurationsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if w.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Insights/components/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/WorkItemConfigs/(?P<workItemConfigId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	workItemConfigIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workItemConfigId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Delete(req.Context(), resourceGroupNameParam, resourceNameParam, workItemConfigIDParam, nil)
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

func (w *WorkItemConfigurationsServerTransport) dispatchGetDefault(req *http.Request) (*http.Response, error) {
	if w.srv.GetDefault == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetDefault not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Insights/components/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/DefaultWorkItemConfig`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.GetDefault(req.Context(), resourceGroupNameParam, resourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).WorkItemConfiguration, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WorkItemConfigurationsServerTransport) dispatchGetItem(req *http.Request) (*http.Response, error) {
	if w.srv.GetItem == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetItem not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Insights/components/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/WorkItemConfigs/(?P<workItemConfigId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	workItemConfigIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workItemConfigId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.GetItem(req.Context(), resourceGroupNameParam, resourceNameParam, workItemConfigIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).WorkItemConfiguration, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WorkItemConfigurationsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if w.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := w.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Insights/components/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/WorkItemConfigs`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		resp := w.srv.NewListPager(resourceGroupNameParam, resourceNameParam, nil)
		newListPager = &resp
		w.newListPager.add(req, newListPager)
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		w.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		w.newListPager.remove(req)
	}
	return resp, nil
}

func (w *WorkItemConfigurationsServerTransport) dispatchUpdateItem(req *http.Request) (*http.Response, error) {
	if w.srv.UpdateItem == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateItem not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Insights/components/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/WorkItemConfigs/(?P<workItemConfigId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armapplicationinsights.WorkItemCreateConfiguration](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	workItemConfigIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workItemConfigId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.UpdateItem(req.Context(), resourceGroupNameParam, resourceNameParam, workItemConfigIDParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).WorkItemConfiguration, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}