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
	"reflect"
	"regexp"
)

// FirewallsServer is a fake server for instances of the armpanngfw.FirewallsClient type.
type FirewallsServer struct {
	// BeginCreateOrUpdate is the fake for method FirewallsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, firewallName string, resource armpanngfw.FirewallResource, options *armpanngfw.FirewallsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armpanngfw.FirewallsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method FirewallsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, firewallName string, options *armpanngfw.FirewallsClientBeginDeleteOptions) (resp azfake.PollerResponder[armpanngfw.FirewallsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method FirewallsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, firewallName string, options *armpanngfw.FirewallsClientGetOptions) (resp azfake.Responder[armpanngfw.FirewallsClientGetResponse], errResp azfake.ErrorResponder)

	// GetGlobalRulestack is the fake for method FirewallsClient.GetGlobalRulestack
	// HTTP status codes to indicate success: http.StatusOK
	GetGlobalRulestack func(ctx context.Context, resourceGroupName string, firewallName string, options *armpanngfw.FirewallsClientGetGlobalRulestackOptions) (resp azfake.Responder[armpanngfw.FirewallsClientGetGlobalRulestackResponse], errResp azfake.ErrorResponder)

	// GetLogProfile is the fake for method FirewallsClient.GetLogProfile
	// HTTP status codes to indicate success: http.StatusOK
	GetLogProfile func(ctx context.Context, resourceGroupName string, firewallName string, options *armpanngfw.FirewallsClientGetLogProfileOptions) (resp azfake.Responder[armpanngfw.FirewallsClientGetLogProfileResponse], errResp azfake.ErrorResponder)

	// GetSupportInfo is the fake for method FirewallsClient.GetSupportInfo
	// HTTP status codes to indicate success: http.StatusOK
	GetSupportInfo func(ctx context.Context, resourceGroupName string, firewallName string, options *armpanngfw.FirewallsClientGetSupportInfoOptions) (resp azfake.Responder[armpanngfw.FirewallsClientGetSupportInfoResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method FirewallsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armpanngfw.FirewallsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armpanngfw.FirewallsClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method FirewallsClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armpanngfw.FirewallsClientListBySubscriptionOptions) (resp azfake.PagerResponder[armpanngfw.FirewallsClientListBySubscriptionResponse])

	// SaveLogProfile is the fake for method FirewallsClient.SaveLogProfile
	// HTTP status codes to indicate success: http.StatusNoContent
	SaveLogProfile func(ctx context.Context, resourceGroupName string, firewallName string, options *armpanngfw.FirewallsClientSaveLogProfileOptions) (resp azfake.Responder[armpanngfw.FirewallsClientSaveLogProfileResponse], errResp azfake.ErrorResponder)

	// Update is the fake for method FirewallsClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, firewallName string, properties armpanngfw.FirewallResourceUpdate, options *armpanngfw.FirewallsClientUpdateOptions) (resp azfake.Responder[armpanngfw.FirewallsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewFirewallsServerTransport creates a new instance of FirewallsServerTransport with the provided implementation.
// The returned FirewallsServerTransport instance is connected to an instance of armpanngfw.FirewallsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewFirewallsServerTransport(srv *FirewallsServer) *FirewallsServerTransport {
	return &FirewallsServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armpanngfw.FirewallsClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armpanngfw.FirewallsClientDeleteResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armpanngfw.FirewallsClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armpanngfw.FirewallsClientListBySubscriptionResponse]](),
	}
}

// FirewallsServerTransport connects instances of armpanngfw.FirewallsClient to instances of FirewallsServer.
// Don't use this type directly, use NewFirewallsServerTransport instead.
type FirewallsServerTransport struct {
	srv                         *FirewallsServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armpanngfw.FirewallsClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armpanngfw.FirewallsClientDeleteResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armpanngfw.FirewallsClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armpanngfw.FirewallsClientListBySubscriptionResponse]]
}

// Do implements the policy.Transporter interface for FirewallsServerTransport.
func (f *FirewallsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "FirewallsClient.BeginCreateOrUpdate":
		resp, err = f.dispatchBeginCreateOrUpdate(req)
	case "FirewallsClient.BeginDelete":
		resp, err = f.dispatchBeginDelete(req)
	case "FirewallsClient.Get":
		resp, err = f.dispatchGet(req)
	case "FirewallsClient.GetGlobalRulestack":
		resp, err = f.dispatchGetGlobalRulestack(req)
	case "FirewallsClient.GetLogProfile":
		resp, err = f.dispatchGetLogProfile(req)
	case "FirewallsClient.GetSupportInfo":
		resp, err = f.dispatchGetSupportInfo(req)
	case "FirewallsClient.NewListByResourceGroupPager":
		resp, err = f.dispatchNewListByResourceGroupPager(req)
	case "FirewallsClient.NewListBySubscriptionPager":
		resp, err = f.dispatchNewListBySubscriptionPager(req)
	case "FirewallsClient.SaveLogProfile":
		resp, err = f.dispatchSaveLogProfile(req)
	case "FirewallsClient.Update":
		resp, err = f.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (f *FirewallsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if f.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := f.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/PaloAltoNetworks\.Cloudngfw/firewalls/(?P<firewallName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armpanngfw.FirewallResource](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		firewallNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("firewallName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := f.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, firewallNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		f.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		f.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		f.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (f *FirewallsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if f.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := f.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/PaloAltoNetworks\.Cloudngfw/firewalls/(?P<firewallName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		firewallNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("firewallName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := f.srv.BeginDelete(req.Context(), resourceGroupNameParam, firewallNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		f.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		f.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		f.beginDelete.remove(req)
	}

	return resp, nil
}

func (f *FirewallsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if f.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/PaloAltoNetworks\.Cloudngfw/firewalls/(?P<firewallName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	firewallNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("firewallName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.Get(req.Context(), resourceGroupNameParam, firewallNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).FirewallResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *FirewallsServerTransport) dispatchGetGlobalRulestack(req *http.Request) (*http.Response, error) {
	if f.srv.GetGlobalRulestack == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetGlobalRulestack not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/PaloAltoNetworks\.Cloudngfw/firewalls/(?P<firewallName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getGlobalRulestack`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	firewallNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("firewallName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.GetGlobalRulestack(req.Context(), resourceGroupNameParam, firewallNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).GlobalRulestackInfo, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *FirewallsServerTransport) dispatchGetLogProfile(req *http.Request) (*http.Response, error) {
	if f.srv.GetLogProfile == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetLogProfile not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/PaloAltoNetworks\.Cloudngfw/firewalls/(?P<firewallName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getLogProfile`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	firewallNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("firewallName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.GetLogProfile(req.Context(), resourceGroupNameParam, firewallNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).LogSettings, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *FirewallsServerTransport) dispatchGetSupportInfo(req *http.Request) (*http.Response, error) {
	if f.srv.GetSupportInfo == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetSupportInfo not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/PaloAltoNetworks\.Cloudngfw/firewalls/(?P<firewallName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getSupportInfo`
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
	firewallNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("firewallName")])
	if err != nil {
		return nil, err
	}
	emailUnescaped, err := url.QueryUnescape(qp.Get("email"))
	if err != nil {
		return nil, err
	}
	emailParam := getOptional(emailUnescaped)
	var options *armpanngfw.FirewallsClientGetSupportInfoOptions
	if emailParam != nil {
		options = &armpanngfw.FirewallsClientGetSupportInfoOptions{
			Email: emailParam,
		}
	}
	respr, errRespr := f.srv.GetSupportInfo(req.Context(), resourceGroupNameParam, firewallNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SupportInfo, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *FirewallsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if f.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := f.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/PaloAltoNetworks\.Cloudngfw/firewalls`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := f.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		f.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armpanngfw.FirewallsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		f.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		f.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (f *FirewallsServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if f.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := f.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/PaloAltoNetworks\.Cloudngfw/firewalls`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := f.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		f.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armpanngfw.FirewallsClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		f.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		f.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (f *FirewallsServerTransport) dispatchSaveLogProfile(req *http.Request) (*http.Response, error) {
	if f.srv.SaveLogProfile == nil {
		return nil, &nonRetriableError{errors.New("fake for method SaveLogProfile not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/PaloAltoNetworks\.Cloudngfw/firewalls/(?P<firewallName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/saveLogProfile`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armpanngfw.LogSettings](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	firewallNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("firewallName")])
	if err != nil {
		return nil, err
	}
	var options *armpanngfw.FirewallsClientSaveLogProfileOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &armpanngfw.FirewallsClientSaveLogProfileOptions{
			LogSettings: &body,
		}
	}
	respr, errRespr := f.srv.SaveLogProfile(req.Context(), resourceGroupNameParam, firewallNameParam, options)
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

func (f *FirewallsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if f.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/PaloAltoNetworks\.Cloudngfw/firewalls/(?P<firewallName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armpanngfw.FirewallResourceUpdate](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	firewallNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("firewallName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.Update(req.Context(), resourceGroupNameParam, firewallNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).FirewallResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
