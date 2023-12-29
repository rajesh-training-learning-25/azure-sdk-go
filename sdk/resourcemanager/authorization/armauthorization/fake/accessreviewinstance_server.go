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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
	"net/http"
	"net/url"
	"regexp"
)

// AccessReviewInstanceServer is a fake server for instances of the armauthorization.AccessReviewInstanceClient type.
type AccessReviewInstanceServer struct {
	// AcceptRecommendations is the fake for method AccessReviewInstanceClient.AcceptRecommendations
	// HTTP status codes to indicate success: http.StatusNoContent
	AcceptRecommendations func(ctx context.Context, scheduleDefinitionID string, id string, options *armauthorization.AccessReviewInstanceClientAcceptRecommendationsOptions) (resp azfake.Responder[armauthorization.AccessReviewInstanceClientAcceptRecommendationsResponse], errResp azfake.ErrorResponder)

	// ApplyDecisions is the fake for method AccessReviewInstanceClient.ApplyDecisions
	// HTTP status codes to indicate success: http.StatusNoContent
	ApplyDecisions func(ctx context.Context, scheduleDefinitionID string, id string, options *armauthorization.AccessReviewInstanceClientApplyDecisionsOptions) (resp azfake.Responder[armauthorization.AccessReviewInstanceClientApplyDecisionsResponse], errResp azfake.ErrorResponder)

	// ResetDecisions is the fake for method AccessReviewInstanceClient.ResetDecisions
	// HTTP status codes to indicate success: http.StatusNoContent
	ResetDecisions func(ctx context.Context, scheduleDefinitionID string, id string, options *armauthorization.AccessReviewInstanceClientResetDecisionsOptions) (resp azfake.Responder[armauthorization.AccessReviewInstanceClientResetDecisionsResponse], errResp azfake.ErrorResponder)

	// SendReminders is the fake for method AccessReviewInstanceClient.SendReminders
	// HTTP status codes to indicate success: http.StatusNoContent
	SendReminders func(ctx context.Context, scheduleDefinitionID string, id string, options *armauthorization.AccessReviewInstanceClientSendRemindersOptions) (resp azfake.Responder[armauthorization.AccessReviewInstanceClientSendRemindersResponse], errResp azfake.ErrorResponder)

	// Stop is the fake for method AccessReviewInstanceClient.Stop
	// HTTP status codes to indicate success: http.StatusNoContent
	Stop func(ctx context.Context, scheduleDefinitionID string, id string, options *armauthorization.AccessReviewInstanceClientStopOptions) (resp azfake.Responder[armauthorization.AccessReviewInstanceClientStopResponse], errResp azfake.ErrorResponder)
}

// NewAccessReviewInstanceServerTransport creates a new instance of AccessReviewInstanceServerTransport with the provided implementation.
// The returned AccessReviewInstanceServerTransport instance is connected to an instance of armauthorization.AccessReviewInstanceClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAccessReviewInstanceServerTransport(srv *AccessReviewInstanceServer) *AccessReviewInstanceServerTransport {
	return &AccessReviewInstanceServerTransport{srv: srv}
}

// AccessReviewInstanceServerTransport connects instances of armauthorization.AccessReviewInstanceClient to instances of AccessReviewInstanceServer.
// Don't use this type directly, use NewAccessReviewInstanceServerTransport instead.
type AccessReviewInstanceServerTransport struct {
	srv *AccessReviewInstanceServer
}

// Do implements the policy.Transporter interface for AccessReviewInstanceServerTransport.
func (a *AccessReviewInstanceServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AccessReviewInstanceClient.AcceptRecommendations":
		resp, err = a.dispatchAcceptRecommendations(req)
	case "AccessReviewInstanceClient.ApplyDecisions":
		resp, err = a.dispatchApplyDecisions(req)
	case "AccessReviewInstanceClient.ResetDecisions":
		resp, err = a.dispatchResetDecisions(req)
	case "AccessReviewInstanceClient.SendReminders":
		resp, err = a.dispatchSendReminders(req)
	case "AccessReviewInstanceClient.Stop":
		resp, err = a.dispatchStop(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AccessReviewInstanceServerTransport) dispatchAcceptRecommendations(req *http.Request) (*http.Response, error) {
	if a.srv.AcceptRecommendations == nil {
		return nil, &nonRetriableError{errors.New("fake for method AcceptRecommendations not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Authorization/accessReviewScheduleDefinitions/(?P<scheduleDefinitionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/instances/(?P<id>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/acceptRecommendations`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
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
	respr, errRespr := a.srv.AcceptRecommendations(req.Context(), scheduleDefinitionIDParam, idParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AccessReviewInstanceServerTransport) dispatchApplyDecisions(req *http.Request) (*http.Response, error) {
	if a.srv.ApplyDecisions == nil {
		return nil, &nonRetriableError{errors.New("fake for method ApplyDecisions not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/accessReviewScheduleDefinitions/(?P<scheduleDefinitionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/instances/(?P<id>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applyDecisions`
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
	respr, errRespr := a.srv.ApplyDecisions(req.Context(), scheduleDefinitionIDParam, idParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AccessReviewInstanceServerTransport) dispatchResetDecisions(req *http.Request) (*http.Response, error) {
	if a.srv.ResetDecisions == nil {
		return nil, &nonRetriableError{errors.New("fake for method ResetDecisions not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/accessReviewScheduleDefinitions/(?P<scheduleDefinitionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/instances/(?P<id>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resetDecisions`
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
	respr, errRespr := a.srv.ResetDecisions(req.Context(), scheduleDefinitionIDParam, idParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AccessReviewInstanceServerTransport) dispatchSendReminders(req *http.Request) (*http.Response, error) {
	if a.srv.SendReminders == nil {
		return nil, &nonRetriableError{errors.New("fake for method SendReminders not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/accessReviewScheduleDefinitions/(?P<scheduleDefinitionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/instances/(?P<id>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sendReminders`
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
	respr, errRespr := a.srv.SendReminders(req.Context(), scheduleDefinitionIDParam, idParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AccessReviewInstanceServerTransport) dispatchStop(req *http.Request) (*http.Response, error) {
	if a.srv.Stop == nil {
		return nil, &nonRetriableError{errors.New("fake for method Stop not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/accessReviewScheduleDefinitions/(?P<scheduleDefinitionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/instances/(?P<id>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/stop`
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
	respr, errRespr := a.srv.Stop(req.Context(), scheduleDefinitionIDParam, idParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
