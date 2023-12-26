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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw"
	"net/http"
	"net/url"
	"regexp"
)

// LocalRulesServer is a fake server for instances of the armpanngfw.LocalRulesClient type.
type LocalRulesServer struct {
	// BeginCreateOrUpdate is the fake for method LocalRulesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, localRulestackName string, priority string, resource armpanngfw.LocalRulesResource, options *armpanngfw.LocalRulesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armpanngfw.LocalRulesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method LocalRulesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, localRulestackName string, priority string, options *armpanngfw.LocalRulesClientBeginDeleteOptions) (resp azfake.PollerResponder[armpanngfw.LocalRulesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method LocalRulesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, localRulestackName string, priority string, options *armpanngfw.LocalRulesClientGetOptions) (resp azfake.Responder[armpanngfw.LocalRulesClientGetResponse], errResp azfake.ErrorResponder)

	// GetCounters is the fake for method LocalRulesClient.GetCounters
	// HTTP status codes to indicate success: http.StatusOK
	GetCounters func(ctx context.Context, resourceGroupName string, localRulestackName string, priority string, options *armpanngfw.LocalRulesClientGetCountersOptions) (resp azfake.Responder[armpanngfw.LocalRulesClientGetCountersResponse], errResp azfake.ErrorResponder)

	// NewListByLocalRulestacksPager is the fake for method LocalRulesClient.NewListByLocalRulestacksPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByLocalRulestacksPager func(resourceGroupName string, localRulestackName string, options *armpanngfw.LocalRulesClientListByLocalRulestacksOptions) (resp azfake.PagerResponder[armpanngfw.LocalRulesClientListByLocalRulestacksResponse])

	// RefreshCounters is the fake for method LocalRulesClient.RefreshCounters
	// HTTP status codes to indicate success: http.StatusNoContent
	RefreshCounters func(ctx context.Context, resourceGroupName string, localRulestackName string, priority string, options *armpanngfw.LocalRulesClientRefreshCountersOptions) (resp azfake.Responder[armpanngfw.LocalRulesClientRefreshCountersResponse], errResp azfake.ErrorResponder)

	// ResetCounters is the fake for method LocalRulesClient.ResetCounters
	// HTTP status codes to indicate success: http.StatusOK
	ResetCounters func(ctx context.Context, resourceGroupName string, localRulestackName string, priority string, options *armpanngfw.LocalRulesClientResetCountersOptions) (resp azfake.Responder[armpanngfw.LocalRulesClientResetCountersResponse], errResp azfake.ErrorResponder)
}

// NewLocalRulesServerTransport creates a new instance of LocalRulesServerTransport with the provided implementation.
// The returned LocalRulesServerTransport instance is connected to an instance of armpanngfw.LocalRulesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewLocalRulesServerTransport(srv *LocalRulesServer) *LocalRulesServerTransport {
	return &LocalRulesServerTransport{
		srv:                           srv,
		beginCreateOrUpdate:           newTracker[azfake.PollerResponder[armpanngfw.LocalRulesClientCreateOrUpdateResponse]](),
		beginDelete:                   newTracker[azfake.PollerResponder[armpanngfw.LocalRulesClientDeleteResponse]](),
		newListByLocalRulestacksPager: newTracker[azfake.PagerResponder[armpanngfw.LocalRulesClientListByLocalRulestacksResponse]](),
	}
}

// LocalRulesServerTransport connects instances of armpanngfw.LocalRulesClient to instances of LocalRulesServer.
// Don't use this type directly, use NewLocalRulesServerTransport instead.
type LocalRulesServerTransport struct {
	srv                           *LocalRulesServer
	beginCreateOrUpdate           *tracker[azfake.PollerResponder[armpanngfw.LocalRulesClientCreateOrUpdateResponse]]
	beginDelete                   *tracker[azfake.PollerResponder[armpanngfw.LocalRulesClientDeleteResponse]]
	newListByLocalRulestacksPager *tracker[azfake.PagerResponder[armpanngfw.LocalRulesClientListByLocalRulestacksResponse]]
}

// Do implements the policy.Transporter interface for LocalRulesServerTransport.
func (l *LocalRulesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "LocalRulesClient.BeginCreateOrUpdate":
		resp, err = l.dispatchBeginCreateOrUpdate(req)
	case "LocalRulesClient.BeginDelete":
		resp, err = l.dispatchBeginDelete(req)
	case "LocalRulesClient.Get":
		resp, err = l.dispatchGet(req)
	case "LocalRulesClient.GetCounters":
		resp, err = l.dispatchGetCounters(req)
	case "LocalRulesClient.NewListByLocalRulestacksPager":
		resp, err = l.dispatchNewListByLocalRulestacksPager(req)
	case "LocalRulesClient.RefreshCounters":
		resp, err = l.dispatchRefreshCounters(req)
	case "LocalRulesClient.ResetCounters":
		resp, err = l.dispatchResetCounters(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (l *LocalRulesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if l.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := l.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/PaloAltoNetworks\.Cloudngfw/localRulestacks/(?P<localRulestackName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/localRules/(?P<priority>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armpanngfw.LocalRulesResource](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		localRulestackNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("localRulestackName")])
		if err != nil {
			return nil, err
		}
		priorityParam, err := url.PathUnescape(matches[regex.SubexpIndex("priority")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, localRulestackNameParam, priorityParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		l.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		l.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		l.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (l *LocalRulesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if l.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := l.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/PaloAltoNetworks\.Cloudngfw/localRulestacks/(?P<localRulestackName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/localRules/(?P<priority>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		localRulestackNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("localRulestackName")])
		if err != nil {
			return nil, err
		}
		priorityParam, err := url.PathUnescape(matches[regex.SubexpIndex("priority")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginDelete(req.Context(), resourceGroupNameParam, localRulestackNameParam, priorityParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		l.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		l.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		l.beginDelete.remove(req)
	}

	return resp, nil
}

func (l *LocalRulesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if l.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/PaloAltoNetworks\.Cloudngfw/localRulestacks/(?P<localRulestackName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/localRules/(?P<priority>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	localRulestackNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("localRulestackName")])
	if err != nil {
		return nil, err
	}
	priorityParam, err := url.PathUnescape(matches[regex.SubexpIndex("priority")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := l.srv.Get(req.Context(), resourceGroupNameParam, localRulestackNameParam, priorityParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).LocalRulesResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LocalRulesServerTransport) dispatchGetCounters(req *http.Request) (*http.Response, error) {
	if l.srv.GetCounters == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetCounters not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/PaloAltoNetworks\.Cloudngfw/localRulestacks/(?P<localRulestackName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/localRules/(?P<priority>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getCounters`
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
	localRulestackNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("localRulestackName")])
	if err != nil {
		return nil, err
	}
	priorityParam, err := url.PathUnescape(matches[regex.SubexpIndex("priority")])
	if err != nil {
		return nil, err
	}
	firewallNameUnescaped, err := url.QueryUnescape(qp.Get("firewallName"))
	if err != nil {
		return nil, err
	}
	firewallNameParam := getOptional(firewallNameUnescaped)
	var options *armpanngfw.LocalRulesClientGetCountersOptions
	if firewallNameParam != nil {
		options = &armpanngfw.LocalRulesClientGetCountersOptions{
			FirewallName: firewallNameParam,
		}
	}
	respr, errRespr := l.srv.GetCounters(req.Context(), resourceGroupNameParam, localRulestackNameParam, priorityParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RuleCounter, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LocalRulesServerTransport) dispatchNewListByLocalRulestacksPager(req *http.Request) (*http.Response, error) {
	if l.srv.NewListByLocalRulestacksPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByLocalRulestacksPager not implemented")}
	}
	newListByLocalRulestacksPager := l.newListByLocalRulestacksPager.get(req)
	if newListByLocalRulestacksPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/PaloAltoNetworks\.Cloudngfw/localRulestacks/(?P<localRulestackName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/localRules`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		localRulestackNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("localRulestackName")])
		if err != nil {
			return nil, err
		}
		resp := l.srv.NewListByLocalRulestacksPager(resourceGroupNameParam, localRulestackNameParam, nil)
		newListByLocalRulestacksPager = &resp
		l.newListByLocalRulestacksPager.add(req, newListByLocalRulestacksPager)
		server.PagerResponderInjectNextLinks(newListByLocalRulestacksPager, req, func(page *armpanngfw.LocalRulesClientListByLocalRulestacksResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByLocalRulestacksPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		l.newListByLocalRulestacksPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByLocalRulestacksPager) {
		l.newListByLocalRulestacksPager.remove(req)
	}
	return resp, nil
}

func (l *LocalRulesServerTransport) dispatchRefreshCounters(req *http.Request) (*http.Response, error) {
	if l.srv.RefreshCounters == nil {
		return nil, &nonRetriableError{errors.New("fake for method RefreshCounters not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/PaloAltoNetworks\.Cloudngfw/localRulestacks/(?P<localRulestackName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/localRules/(?P<priority>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/refreshCounters`
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
	localRulestackNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("localRulestackName")])
	if err != nil {
		return nil, err
	}
	priorityParam, err := url.PathUnescape(matches[regex.SubexpIndex("priority")])
	if err != nil {
		return nil, err
	}
	firewallNameUnescaped, err := url.QueryUnescape(qp.Get("firewallName"))
	if err != nil {
		return nil, err
	}
	firewallNameParam := getOptional(firewallNameUnescaped)
	var options *armpanngfw.LocalRulesClientRefreshCountersOptions
	if firewallNameParam != nil {
		options = &armpanngfw.LocalRulesClientRefreshCountersOptions{
			FirewallName: firewallNameParam,
		}
	}
	respr, errRespr := l.srv.RefreshCounters(req.Context(), resourceGroupNameParam, localRulestackNameParam, priorityParam, options)
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

func (l *LocalRulesServerTransport) dispatchResetCounters(req *http.Request) (*http.Response, error) {
	if l.srv.ResetCounters == nil {
		return nil, &nonRetriableError{errors.New("fake for method ResetCounters not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/PaloAltoNetworks\.Cloudngfw/localRulestacks/(?P<localRulestackName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/localRules/(?P<priority>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resetCounters`
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
	localRulestackNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("localRulestackName")])
	if err != nil {
		return nil, err
	}
	priorityParam, err := url.PathUnescape(matches[regex.SubexpIndex("priority")])
	if err != nil {
		return nil, err
	}
	firewallNameUnescaped, err := url.QueryUnescape(qp.Get("firewallName"))
	if err != nil {
		return nil, err
	}
	firewallNameParam := getOptional(firewallNameUnescaped)
	var options *armpanngfw.LocalRulesClientResetCountersOptions
	if firewallNameParam != nil {
		options = &armpanngfw.LocalRulesClientResetCountersOptions{
			FirewallName: firewallNameParam,
		}
	}
	respr, errRespr := l.srv.ResetCounters(req.Context(), resourceGroupNameParam, localRulestackNameParam, priorityParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RuleCounterReset, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
