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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
)

// ProjectsServer is a fake server for instances of the armmigrate.ProjectsClient type.
type ProjectsServer struct {
	// AssessmentOptions is the fake for method ProjectsClient.AssessmentOptions
	// HTTP status codes to indicate success: http.StatusOK
	AssessmentOptions func(ctx context.Context, resourceGroupName string, projectName string, assessmentOptionsName string, options *armmigrate.ProjectsClientAssessmentOptionsOptions) (resp azfake.Responder[armmigrate.ProjectsClientAssessmentOptionsResponse], errResp azfake.ErrorResponder)

	// NewAssessmentOptionsListPager is the fake for method ProjectsClient.NewAssessmentOptionsListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewAssessmentOptionsListPager func(resourceGroupName string, projectName string, options *armmigrate.ProjectsClientAssessmentOptionsListOptions) (resp azfake.PagerResponder[armmigrate.ProjectsClientAssessmentOptionsListResponse])

	// Create is the fake for method ProjectsClient.Create
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	Create func(ctx context.Context, resourceGroupName string, projectName string, options *armmigrate.ProjectsClientCreateOptions) (resp azfake.Responder[armmigrate.ProjectsClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ProjectsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, projectName string, options *armmigrate.ProjectsClientDeleteOptions) (resp azfake.Responder[armmigrate.ProjectsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ProjectsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, projectName string, options *armmigrate.ProjectsClientGetOptions) (resp azfake.Responder[armmigrate.ProjectsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ProjectsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, options *armmigrate.ProjectsClientListOptions) (resp azfake.PagerResponder[armmigrate.ProjectsClientListResponse])

	// NewListBySubscriptionPager is the fake for method ProjectsClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armmigrate.ProjectsClientListBySubscriptionOptions) (resp azfake.PagerResponder[armmigrate.ProjectsClientListBySubscriptionResponse])

	// Update is the fake for method ProjectsClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, projectName string, options *armmigrate.ProjectsClientUpdateOptions) (resp azfake.Responder[armmigrate.ProjectsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewProjectsServerTransport creates a new instance of ProjectsServerTransport with the provided implementation.
// The returned ProjectsServerTransport instance is connected to an instance of armmigrate.ProjectsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewProjectsServerTransport(srv *ProjectsServer) *ProjectsServerTransport {
	return &ProjectsServerTransport{
		srv:                           srv,
		newAssessmentOptionsListPager: newTracker[azfake.PagerResponder[armmigrate.ProjectsClientAssessmentOptionsListResponse]](),
		newListPager:                  newTracker[azfake.PagerResponder[armmigrate.ProjectsClientListResponse]](),
		newListBySubscriptionPager:    newTracker[azfake.PagerResponder[armmigrate.ProjectsClientListBySubscriptionResponse]](),
	}
}

// ProjectsServerTransport connects instances of armmigrate.ProjectsClient to instances of ProjectsServer.
// Don't use this type directly, use NewProjectsServerTransport instead.
type ProjectsServerTransport struct {
	srv                           *ProjectsServer
	newAssessmentOptionsListPager *tracker[azfake.PagerResponder[armmigrate.ProjectsClientAssessmentOptionsListResponse]]
	newListPager                  *tracker[azfake.PagerResponder[armmigrate.ProjectsClientListResponse]]
	newListBySubscriptionPager    *tracker[azfake.PagerResponder[armmigrate.ProjectsClientListBySubscriptionResponse]]
}

// Do implements the policy.Transporter interface for ProjectsServerTransport.
func (p *ProjectsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ProjectsClient.AssessmentOptions":
		resp, err = p.dispatchAssessmentOptions(req)
	case "ProjectsClient.NewAssessmentOptionsListPager":
		resp, err = p.dispatchNewAssessmentOptionsListPager(req)
	case "ProjectsClient.Create":
		resp, err = p.dispatchCreate(req)
	case "ProjectsClient.Delete":
		resp, err = p.dispatchDelete(req)
	case "ProjectsClient.Get":
		resp, err = p.dispatchGet(req)
	case "ProjectsClient.NewListPager":
		resp, err = p.dispatchNewListPager(req)
	case "ProjectsClient.NewListBySubscriptionPager":
		resp, err = p.dispatchNewListBySubscriptionPager(req)
	case "ProjectsClient.Update":
		resp, err = p.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *ProjectsServerTransport) dispatchAssessmentOptions(req *http.Request) (*http.Response, error) {
	if p.srv.AssessmentOptions == nil {
		return nil, &nonRetriableError{errors.New("fake for method AssessmentOptions not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/assessmentOptions/(?P<assessmentOptionsName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	projectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("projectName")])
	if err != nil {
		return nil, err
	}
	assessmentOptionsNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("assessmentOptionsName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.AssessmentOptions(req.Context(), resourceGroupNameParam, projectNameParam, assessmentOptionsNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AssessmentOptions, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).XMSRequestID; val != nil {
		resp.Header.Set("x-ms-request-id", *val)
	}
	return resp, nil
}

func (p *ProjectsServerTransport) dispatchNewAssessmentOptionsListPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewAssessmentOptionsListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewAssessmentOptionsListPager not implemented")}
	}
	newAssessmentOptionsListPager := p.newAssessmentOptionsListPager.get(req)
	if newAssessmentOptionsListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/assessmentOptions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		projectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("projectName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewAssessmentOptionsListPager(resourceGroupNameParam, projectNameParam, nil)
		newAssessmentOptionsListPager = &resp
		p.newAssessmentOptionsListPager.add(req, newAssessmentOptionsListPager)
	}
	resp, err := server.PagerResponderNext(newAssessmentOptionsListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newAssessmentOptionsListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newAssessmentOptionsListPager) {
		p.newAssessmentOptionsListPager.remove(req)
	}
	return resp, nil
}

func (p *ProjectsServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if p.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmigrate.Project](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	projectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("projectName")])
	if err != nil {
		return nil, err
	}
	var options *armmigrate.ProjectsClientCreateOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &armmigrate.ProjectsClientCreateOptions{
			Project: &body,
		}
	}
	respr, errRespr := p.srv.Create(req.Context(), resourceGroupNameParam, projectNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Project, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).XMSRequestID; val != nil {
		resp.Header.Set("x-ms-request-id", *val)
	}
	return resp, nil
}

func (p *ProjectsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if p.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	projectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("projectName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Delete(req.Context(), resourceGroupNameParam, projectNameParam, nil)
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
	if val := server.GetResponse(respr).XMSRequestID; val != nil {
		resp.Header.Set("x-ms-request-id", *val)
	}
	return resp, nil
}

func (p *ProjectsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	projectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("projectName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, projectNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Project, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).XMSRequestID; val != nil {
		resp.Header.Set("x-ms-request-id", *val)
	}
	return resp, nil
}

func (p *ProjectsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := p.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListPager(resourceGroupNameParam, nil)
		newListPager = &resp
		p.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armmigrate.ProjectsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		p.newListPager.remove(req)
	}
	return resp, nil
}

func (p *ProjectsServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := p.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := p.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		p.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armmigrate.ProjectsClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		p.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (p *ProjectsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmigrate.Project](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	projectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("projectName")])
	if err != nil {
		return nil, err
	}
	var options *armmigrate.ProjectsClientUpdateOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &armmigrate.ProjectsClientUpdateOptions{
			Project: &body,
		}
	}
	respr, errRespr := p.srv.Update(req.Context(), resourceGroupNameParam, projectNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Project, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).XMSRequestID; val != nil {
		resp.Header.Set("x-ms-request-id", *val)
	}
	return resp, nil
}
