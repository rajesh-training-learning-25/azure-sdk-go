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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub"
	"net/http"
	"net/url"
	"regexp"
)

// DefaultRolloutsServer is a fake server for instances of the armproviderhub.DefaultRolloutsClient type.
type DefaultRolloutsServer struct {
	// BeginCreateOrUpdate is the fake for method DefaultRolloutsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, providerNamespace string, rolloutName string, properties armproviderhub.DefaultRollout, options *armproviderhub.DefaultRolloutsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armproviderhub.DefaultRolloutsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method DefaultRolloutsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, providerNamespace string, rolloutName string, options *armproviderhub.DefaultRolloutsClientDeleteOptions) (resp azfake.Responder[armproviderhub.DefaultRolloutsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method DefaultRolloutsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, providerNamespace string, rolloutName string, options *armproviderhub.DefaultRolloutsClientGetOptions) (resp azfake.Responder[armproviderhub.DefaultRolloutsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByProviderRegistrationPager is the fake for method DefaultRolloutsClient.NewListByProviderRegistrationPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByProviderRegistrationPager func(providerNamespace string, options *armproviderhub.DefaultRolloutsClientListByProviderRegistrationOptions) (resp azfake.PagerResponder[armproviderhub.DefaultRolloutsClientListByProviderRegistrationResponse])

	// Stop is the fake for method DefaultRolloutsClient.Stop
	// HTTP status codes to indicate success: http.StatusOK
	Stop func(ctx context.Context, providerNamespace string, rolloutName string, options *armproviderhub.DefaultRolloutsClientStopOptions) (resp azfake.Responder[armproviderhub.DefaultRolloutsClientStopResponse], errResp azfake.ErrorResponder)
}

// NewDefaultRolloutsServerTransport creates a new instance of DefaultRolloutsServerTransport with the provided implementation.
// The returned DefaultRolloutsServerTransport instance is connected to an instance of armproviderhub.DefaultRolloutsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDefaultRolloutsServerTransport(srv *DefaultRolloutsServer) *DefaultRolloutsServerTransport {
	return &DefaultRolloutsServerTransport{
		srv:                                srv,
		beginCreateOrUpdate:                newTracker[azfake.PollerResponder[armproviderhub.DefaultRolloutsClientCreateOrUpdateResponse]](),
		newListByProviderRegistrationPager: newTracker[azfake.PagerResponder[armproviderhub.DefaultRolloutsClientListByProviderRegistrationResponse]](),
	}
}

// DefaultRolloutsServerTransport connects instances of armproviderhub.DefaultRolloutsClient to instances of DefaultRolloutsServer.
// Don't use this type directly, use NewDefaultRolloutsServerTransport instead.
type DefaultRolloutsServerTransport struct {
	srv                                *DefaultRolloutsServer
	beginCreateOrUpdate                *tracker[azfake.PollerResponder[armproviderhub.DefaultRolloutsClientCreateOrUpdateResponse]]
	newListByProviderRegistrationPager *tracker[azfake.PagerResponder[armproviderhub.DefaultRolloutsClientListByProviderRegistrationResponse]]
}

// Do implements the policy.Transporter interface for DefaultRolloutsServerTransport.
func (d *DefaultRolloutsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DefaultRolloutsClient.BeginCreateOrUpdate":
		resp, err = d.dispatchBeginCreateOrUpdate(req)
	case "DefaultRolloutsClient.Delete":
		resp, err = d.dispatchDelete(req)
	case "DefaultRolloutsClient.Get":
		resp, err = d.dispatchGet(req)
	case "DefaultRolloutsClient.NewListByProviderRegistrationPager":
		resp, err = d.dispatchNewListByProviderRegistrationPager(req)
	case "DefaultRolloutsClient.Stop":
		resp, err = d.dispatchStop(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DefaultRolloutsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := d.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ProviderHub/providerRegistrations/(?P<providerNamespace>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/defaultRollouts/(?P<rolloutName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armproviderhub.DefaultRollout](req)
		if err != nil {
			return nil, err
		}
		providerNamespaceParam, err := url.PathUnescape(matches[regex.SubexpIndex("providerNamespace")])
		if err != nil {
			return nil, err
		}
		rolloutNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("rolloutName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginCreateOrUpdate(req.Context(), providerNamespaceParam, rolloutNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		d.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		d.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		d.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (d *DefaultRolloutsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if d.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ProviderHub/providerRegistrations/(?P<providerNamespace>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/defaultRollouts/(?P<rolloutName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	providerNamespaceParam, err := url.PathUnescape(matches[regex.SubexpIndex("providerNamespace")])
	if err != nil {
		return nil, err
	}
	rolloutNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("rolloutName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Delete(req.Context(), providerNamespaceParam, rolloutNameParam, nil)
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

func (d *DefaultRolloutsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ProviderHub/providerRegistrations/(?P<providerNamespace>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/defaultRollouts/(?P<rolloutName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	providerNamespaceParam, err := url.PathUnescape(matches[regex.SubexpIndex("providerNamespace")])
	if err != nil {
		return nil, err
	}
	rolloutNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("rolloutName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Get(req.Context(), providerNamespaceParam, rolloutNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DefaultRollout, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DefaultRolloutsServerTransport) dispatchNewListByProviderRegistrationPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListByProviderRegistrationPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByProviderRegistrationPager not implemented")}
	}
	newListByProviderRegistrationPager := d.newListByProviderRegistrationPager.get(req)
	if newListByProviderRegistrationPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ProviderHub/providerRegistrations/(?P<providerNamespace>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/defaultRollouts`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		providerNamespaceParam, err := url.PathUnescape(matches[regex.SubexpIndex("providerNamespace")])
		if err != nil {
			return nil, err
		}
		resp := d.srv.NewListByProviderRegistrationPager(providerNamespaceParam, nil)
		newListByProviderRegistrationPager = &resp
		d.newListByProviderRegistrationPager.add(req, newListByProviderRegistrationPager)
		server.PagerResponderInjectNextLinks(newListByProviderRegistrationPager, req, func(page *armproviderhub.DefaultRolloutsClientListByProviderRegistrationResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByProviderRegistrationPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListByProviderRegistrationPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByProviderRegistrationPager) {
		d.newListByProviderRegistrationPager.remove(req)
	}
	return resp, nil
}

func (d *DefaultRolloutsServerTransport) dispatchStop(req *http.Request) (*http.Response, error) {
	if d.srv.Stop == nil {
		return nil, &nonRetriableError{errors.New("fake for method Stop not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ProviderHub/providerRegistrations/(?P<providerNamespace>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/defaultRollouts/(?P<rolloutName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/stop`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	providerNamespaceParam, err := url.PathUnescape(matches[regex.SubexpIndex("providerNamespace")])
	if err != nil {
		return nil, err
	}
	rolloutNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("rolloutName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Stop(req.Context(), providerNamespaceParam, rolloutNameParam, nil)
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
