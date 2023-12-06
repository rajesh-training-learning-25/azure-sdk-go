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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthbot/armhealthbot"
	"net/http"
	"net/url"
	"regexp"
)

// BotsServer is a fake server for instances of the armhealthbot.BotsClient type.
type BotsServer struct {
	// BeginCreate is the fake for method BotsClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, botName string, parameters armhealthbot.HealthBot, options *armhealthbot.BotsClientBeginCreateOptions) (resp azfake.PollerResponder[armhealthbot.BotsClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method BotsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, botName string, options *armhealthbot.BotsClientBeginDeleteOptions) (resp azfake.PollerResponder[armhealthbot.BotsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method BotsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, botName string, options *armhealthbot.BotsClientGetOptions) (resp azfake.Responder[armhealthbot.BotsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method BotsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armhealthbot.BotsClientListOptions) (resp azfake.PagerResponder[armhealthbot.BotsClientListResponse])

	// NewListByResourceGroupPager is the fake for method BotsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armhealthbot.BotsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armhealthbot.BotsClientListByResourceGroupResponse])

	// Update is the fake for method BotsClient.Update
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	Update func(ctx context.Context, resourceGroupName string, botName string, parameters armhealthbot.UpdateParameters, options *armhealthbot.BotsClientUpdateOptions) (resp azfake.Responder[armhealthbot.BotsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewBotsServerTransport creates a new instance of BotsServerTransport with the provided implementation.
// The returned BotsServerTransport instance is connected to an instance of armhealthbot.BotsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBotsServerTransport(srv *BotsServer) *BotsServerTransport {
	return &BotsServerTransport{
		srv:                         srv,
		beginCreate:                 newTracker[azfake.PollerResponder[armhealthbot.BotsClientCreateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armhealthbot.BotsClientDeleteResponse]](),
		newListPager:                newTracker[azfake.PagerResponder[armhealthbot.BotsClientListResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armhealthbot.BotsClientListByResourceGroupResponse]](),
	}
}

// BotsServerTransport connects instances of armhealthbot.BotsClient to instances of BotsServer.
// Don't use this type directly, use NewBotsServerTransport instead.
type BotsServerTransport struct {
	srv                         *BotsServer
	beginCreate                 *tracker[azfake.PollerResponder[armhealthbot.BotsClientCreateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armhealthbot.BotsClientDeleteResponse]]
	newListPager                *tracker[azfake.PagerResponder[armhealthbot.BotsClientListResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armhealthbot.BotsClientListByResourceGroupResponse]]
}

// Do implements the policy.Transporter interface for BotsServerTransport.
func (b *BotsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "BotsClient.BeginCreate":
		resp, err = b.dispatchBeginCreate(req)
	case "BotsClient.BeginDelete":
		resp, err = b.dispatchBeginDelete(req)
	case "BotsClient.Get":
		resp, err = b.dispatchGet(req)
	case "BotsClient.NewListPager":
		resp, err = b.dispatchNewListPager(req)
	case "BotsClient.NewListByResourceGroupPager":
		resp, err = b.dispatchNewListByResourceGroupPager(req)
	case "BotsClient.Update":
		resp, err = b.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b *BotsServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if b.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := b.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HealthBot/healthBots/(?P<botName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armhealthbot.HealthBot](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		botNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("botName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := b.srv.BeginCreate(req.Context(), resourceGroupNameParam, botNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		b.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		b.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		b.beginCreate.remove(req)
	}

	return resp, nil
}

func (b *BotsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if b.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := b.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HealthBot/healthBots/(?P<botName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		botNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("botName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := b.srv.BeginDelete(req.Context(), resourceGroupNameParam, botNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		b.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		b.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		b.beginDelete.remove(req)
	}

	return resp, nil
}

func (b *BotsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if b.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HealthBot/healthBots/(?P<botName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	botNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("botName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Get(req.Context(), resourceGroupNameParam, botNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).HealthBot, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BotsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if b.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := b.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HealthBot/healthBots`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := b.srv.NewListPager(nil)
		newListPager = &resp
		b.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armhealthbot.BotsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		b.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		b.newListPager.remove(req)
	}
	return resp, nil
}

func (b *BotsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if b.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := b.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HealthBot/healthBots`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := b.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		b.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armhealthbot.BotsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		b.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		b.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (b *BotsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if b.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HealthBot/healthBots/(?P<botName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armhealthbot.UpdateParameters](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	botNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("botName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Update(req.Context(), resourceGroupNameParam, botNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).HealthBot, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}