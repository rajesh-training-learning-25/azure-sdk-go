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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deploymentmanager/armdeploymentmanager"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
)

// RolloutsServer is a fake server for instances of the armdeploymentmanager.RolloutsClient type.
type RolloutsServer struct {
	// Cancel is the fake for method RolloutsClient.Cancel
	// HTTP status codes to indicate success: http.StatusOK
	Cancel func(ctx context.Context, resourceGroupName string, rolloutName string, options *armdeploymentmanager.RolloutsClientCancelOptions) (resp azfake.Responder[armdeploymentmanager.RolloutsClientCancelResponse], errResp azfake.ErrorResponder)

	// BeginCreateOrUpdate is the fake for method RolloutsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, rolloutName string, options *armdeploymentmanager.RolloutsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armdeploymentmanager.RolloutsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method RolloutsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, rolloutName string, options *armdeploymentmanager.RolloutsClientDeleteOptions) (resp azfake.Responder[armdeploymentmanager.RolloutsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method RolloutsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, rolloutName string, options *armdeploymentmanager.RolloutsClientGetOptions) (resp azfake.Responder[armdeploymentmanager.RolloutsClientGetResponse], errResp azfake.ErrorResponder)

	// List is the fake for method RolloutsClient.List
	// HTTP status codes to indicate success: http.StatusOK
	List func(ctx context.Context, resourceGroupName string, options *armdeploymentmanager.RolloutsClientListOptions) (resp azfake.Responder[armdeploymentmanager.RolloutsClientListResponse], errResp azfake.ErrorResponder)

	// Restart is the fake for method RolloutsClient.Restart
	// HTTP status codes to indicate success: http.StatusOK
	Restart func(ctx context.Context, resourceGroupName string, rolloutName string, options *armdeploymentmanager.RolloutsClientRestartOptions) (resp azfake.Responder[armdeploymentmanager.RolloutsClientRestartResponse], errResp azfake.ErrorResponder)
}

// NewRolloutsServerTransport creates a new instance of RolloutsServerTransport with the provided implementation.
// The returned RolloutsServerTransport instance is connected to an instance of armdeploymentmanager.RolloutsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRolloutsServerTransport(srv *RolloutsServer) *RolloutsServerTransport {
	return &RolloutsServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armdeploymentmanager.RolloutsClientCreateOrUpdateResponse]](),
	}
}

// RolloutsServerTransport connects instances of armdeploymentmanager.RolloutsClient to instances of RolloutsServer.
// Don't use this type directly, use NewRolloutsServerTransport instead.
type RolloutsServerTransport struct {
	srv                 *RolloutsServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armdeploymentmanager.RolloutsClientCreateOrUpdateResponse]]
}

// Do implements the policy.Transporter interface for RolloutsServerTransport.
func (r *RolloutsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "RolloutsClient.Cancel":
		resp, err = r.dispatchCancel(req)
	case "RolloutsClient.BeginCreateOrUpdate":
		resp, err = r.dispatchBeginCreateOrUpdate(req)
	case "RolloutsClient.Delete":
		resp, err = r.dispatchDelete(req)
	case "RolloutsClient.Get":
		resp, err = r.dispatchGet(req)
	case "RolloutsClient.List":
		resp, err = r.dispatchList(req)
	case "RolloutsClient.Restart":
		resp, err = r.dispatchRestart(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *RolloutsServerTransport) dispatchCancel(req *http.Request) (*http.Response, error) {
	if r.srv.Cancel == nil {
		return nil, &nonRetriableError{errors.New("fake for method Cancel not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DeploymentManager/rollouts/(?P<rolloutName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cancel`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	rolloutNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("rolloutName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Cancel(req.Context(), resourceGroupNameParam, rolloutNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Rollout, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RolloutsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if r.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := r.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DeploymentManager/rollouts/(?P<rolloutName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdeploymentmanager.RolloutRequest](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		rolloutNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("rolloutName")])
		if err != nil {
			return nil, err
		}
		var options *armdeploymentmanager.RolloutsClientBeginCreateOrUpdateOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armdeploymentmanager.RolloutsClientBeginCreateOrUpdateOptions{
				RolloutRequest: &body,
			}
		}
		respr, errRespr := r.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, rolloutNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		r.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusCreated}, resp.StatusCode) {
		r.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		r.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (r *RolloutsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if r.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DeploymentManager/rollouts/(?P<rolloutName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	rolloutNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("rolloutName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Delete(req.Context(), resourceGroupNameParam, rolloutNameParam, nil)
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

func (r *RolloutsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DeploymentManager/rollouts/(?P<rolloutName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	rolloutNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("rolloutName")])
	if err != nil {
		return nil, err
	}
	retryAttemptUnescaped, err := url.QueryUnescape(qp.Get("retryAttempt"))
	if err != nil {
		return nil, err
	}
	retryAttemptParam, err := parseOptional(retryAttemptUnescaped, func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	var options *armdeploymentmanager.RolloutsClientGetOptions
	if retryAttemptParam != nil {
		options = &armdeploymentmanager.RolloutsClientGetOptions{
			RetryAttempt: retryAttemptParam,
		}
	}
	respr, errRespr := r.srv.Get(req.Context(), resourceGroupNameParam, rolloutNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Rollout, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RolloutsServerTransport) dispatchList(req *http.Request) (*http.Response, error) {
	if r.srv.List == nil {
		return nil, &nonRetriableError{errors.New("fake for method List not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DeploymentManager/rollouts`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.List(req.Context(), resourceGroupNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RolloutArray, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RolloutsServerTransport) dispatchRestart(req *http.Request) (*http.Response, error) {
	if r.srv.Restart == nil {
		return nil, &nonRetriableError{errors.New("fake for method Restart not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DeploymentManager/rollouts/(?P<rolloutName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/restart`
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
	rolloutNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("rolloutName")])
	if err != nil {
		return nil, err
	}
	skipSucceededUnescaped, err := url.QueryUnescape(qp.Get("skipSucceeded"))
	if err != nil {
		return nil, err
	}
	skipSucceededParam, err := parseOptional(skipSucceededUnescaped, strconv.ParseBool)
	if err != nil {
		return nil, err
	}
	var options *armdeploymentmanager.RolloutsClientRestartOptions
	if skipSucceededParam != nil {
		options = &armdeploymentmanager.RolloutsClientRestartOptions{
			SkipSucceeded: skipSucceededParam,
		}
	}
	respr, errRespr := r.srv.Restart(req.Context(), resourceGroupNameParam, rolloutNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Rollout, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}