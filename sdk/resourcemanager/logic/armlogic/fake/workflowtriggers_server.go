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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// WorkflowTriggersServer is a fake server for instances of the armlogic.WorkflowTriggersClient type.
type WorkflowTriggersServer struct {
	// Get is the fake for method WorkflowTriggersClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, workflowName string, triggerName string, options *armlogic.WorkflowTriggersClientGetOptions) (resp azfake.Responder[armlogic.WorkflowTriggersClientGetResponse], errResp azfake.ErrorResponder)

	// GetSchemaJSON is the fake for method WorkflowTriggersClient.GetSchemaJSON
	// HTTP status codes to indicate success: http.StatusOK
	GetSchemaJSON func(ctx context.Context, resourceGroupName string, workflowName string, triggerName string, options *armlogic.WorkflowTriggersClientGetSchemaJSONOptions) (resp azfake.Responder[armlogic.WorkflowTriggersClientGetSchemaJSONResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method WorkflowTriggersClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, workflowName string, options *armlogic.WorkflowTriggersClientListOptions) (resp azfake.PagerResponder[armlogic.WorkflowTriggersClientListResponse])

	// ListCallbackURL is the fake for method WorkflowTriggersClient.ListCallbackURL
	// HTTP status codes to indicate success: http.StatusOK
	ListCallbackURL func(ctx context.Context, resourceGroupName string, workflowName string, triggerName string, options *armlogic.WorkflowTriggersClientListCallbackURLOptions) (resp azfake.Responder[armlogic.WorkflowTriggersClientListCallbackURLResponse], errResp azfake.ErrorResponder)

	// Reset is the fake for method WorkflowTriggersClient.Reset
	// HTTP status codes to indicate success: http.StatusOK
	Reset func(ctx context.Context, resourceGroupName string, workflowName string, triggerName string, options *armlogic.WorkflowTriggersClientResetOptions) (resp azfake.Responder[armlogic.WorkflowTriggersClientResetResponse], errResp azfake.ErrorResponder)

	// Run is the fake for method WorkflowTriggersClient.Run
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	Run func(ctx context.Context, resourceGroupName string, workflowName string, triggerName string, options *armlogic.WorkflowTriggersClientRunOptions) (resp azfake.Responder[armlogic.WorkflowTriggersClientRunResponse], errResp azfake.ErrorResponder)

	// SetState is the fake for method WorkflowTriggersClient.SetState
	// HTTP status codes to indicate success: http.StatusOK
	SetState func(ctx context.Context, resourceGroupName string, workflowName string, triggerName string, setState armlogic.SetTriggerStateActionDefinition, options *armlogic.WorkflowTriggersClientSetStateOptions) (resp azfake.Responder[armlogic.WorkflowTriggersClientSetStateResponse], errResp azfake.ErrorResponder)
}

// NewWorkflowTriggersServerTransport creates a new instance of WorkflowTriggersServerTransport with the provided implementation.
// The returned WorkflowTriggersServerTransport instance is connected to an instance of armlogic.WorkflowTriggersClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewWorkflowTriggersServerTransport(srv *WorkflowTriggersServer) *WorkflowTriggersServerTransport {
	return &WorkflowTriggersServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armlogic.WorkflowTriggersClientListResponse]](),
	}
}

// WorkflowTriggersServerTransport connects instances of armlogic.WorkflowTriggersClient to instances of WorkflowTriggersServer.
// Don't use this type directly, use NewWorkflowTriggersServerTransport instead.
type WorkflowTriggersServerTransport struct {
	srv          *WorkflowTriggersServer
	newListPager *tracker[azfake.PagerResponder[armlogic.WorkflowTriggersClientListResponse]]
}

// Do implements the policy.Transporter interface for WorkflowTriggersServerTransport.
func (w *WorkflowTriggersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "WorkflowTriggersClient.Get":
		resp, err = w.dispatchGet(req)
	case "WorkflowTriggersClient.GetSchemaJSON":
		resp, err = w.dispatchGetSchemaJSON(req)
	case "WorkflowTriggersClient.NewListPager":
		resp, err = w.dispatchNewListPager(req)
	case "WorkflowTriggersClient.ListCallbackURL":
		resp, err = w.dispatchListCallbackURL(req)
	case "WorkflowTriggersClient.Reset":
		resp, err = w.dispatchReset(req)
	case "WorkflowTriggersClient.Run":
		resp, err = w.dispatchRun(req)
	case "WorkflowTriggersClient.SetState":
		resp, err = w.dispatchSetState(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w *WorkflowTriggersServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if w.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Logic/workflows/(?P<workflowName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/triggers/(?P<triggerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workflowNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workflowName")])
	if err != nil {
		return nil, err
	}
	triggerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("triggerName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Get(req.Context(), resourceGroupNameParam, workflowNameParam, triggerNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).WorkflowTrigger, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WorkflowTriggersServerTransport) dispatchGetSchemaJSON(req *http.Request) (*http.Response, error) {
	if w.srv.GetSchemaJSON == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetSchemaJSON not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Logic/workflows/(?P<workflowName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/triggers/(?P<triggerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schemas/json`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workflowNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workflowName")])
	if err != nil {
		return nil, err
	}
	triggerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("triggerName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.GetSchemaJSON(req.Context(), resourceGroupNameParam, workflowNameParam, triggerNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).JSONSchema, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WorkflowTriggersServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if w.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := w.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Logic/workflows/(?P<workflowName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/triggers`
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
		workflowNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workflowName")])
		if err != nil {
			return nil, err
		}
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armlogic.WorkflowTriggersClientListOptions
		if topParam != nil || filterParam != nil {
			options = &armlogic.WorkflowTriggersClientListOptions{
				Top:    topParam,
				Filter: filterParam,
			}
		}
		resp := w.srv.NewListPager(resourceGroupNameParam, workflowNameParam, options)
		newListPager = &resp
		w.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armlogic.WorkflowTriggersClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
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

func (w *WorkflowTriggersServerTransport) dispatchListCallbackURL(req *http.Request) (*http.Response, error) {
	if w.srv.ListCallbackURL == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListCallbackURL not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Logic/workflows/(?P<workflowName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/triggers/(?P<triggerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listCallbackUrl`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workflowNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workflowName")])
	if err != nil {
		return nil, err
	}
	triggerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("triggerName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.ListCallbackURL(req.Context(), resourceGroupNameParam, workflowNameParam, triggerNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).WorkflowTriggerCallbackURL, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WorkflowTriggersServerTransport) dispatchReset(req *http.Request) (*http.Response, error) {
	if w.srv.Reset == nil {
		return nil, &nonRetriableError{errors.New("fake for method Reset not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Logic/workflows/(?P<workflowName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/triggers/(?P<triggerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reset`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workflowNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workflowName")])
	if err != nil {
		return nil, err
	}
	triggerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("triggerName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Reset(req.Context(), resourceGroupNameParam, workflowNameParam, triggerNameParam, nil)
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

func (w *WorkflowTriggersServerTransport) dispatchRun(req *http.Request) (*http.Response, error) {
	if w.srv.Run == nil {
		return nil, &nonRetriableError{errors.New("fake for method Run not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Logic/workflows/(?P<workflowName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/triggers/(?P<triggerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/run`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workflowNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workflowName")])
	if err != nil {
		return nil, err
	}
	triggerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("triggerName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Run(req.Context(), resourceGroupNameParam, workflowNameParam, triggerNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusAccepted}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WorkflowTriggersServerTransport) dispatchSetState(req *http.Request) (*http.Response, error) {
	if w.srv.SetState == nil {
		return nil, &nonRetriableError{errors.New("fake for method SetState not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Logic/workflows/(?P<workflowName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/triggers/(?P<triggerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/setState`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armlogic.SetTriggerStateActionDefinition](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workflowNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workflowName")])
	if err != nil {
		return nil, err
	}
	triggerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("triggerName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.SetState(req.Context(), resourceGroupNameParam, workflowNameParam, triggerNameParam, body, nil)
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
