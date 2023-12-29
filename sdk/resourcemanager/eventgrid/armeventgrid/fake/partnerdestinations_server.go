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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// PartnerDestinationsServer is a fake server for instances of the armeventgrid.PartnerDestinationsClient type.
type PartnerDestinationsServer struct {
	// Activate is the fake for method PartnerDestinationsClient.Activate
	// HTTP status codes to indicate success: http.StatusOK
	Activate func(ctx context.Context, resourceGroupName string, partnerDestinationName string, options *armeventgrid.PartnerDestinationsClientActivateOptions) (resp azfake.Responder[armeventgrid.PartnerDestinationsClientActivateResponse], errResp azfake.ErrorResponder)

	// BeginCreateOrUpdate is the fake for method PartnerDestinationsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, partnerDestinationName string, partnerDestination armeventgrid.PartnerDestination, options *armeventgrid.PartnerDestinationsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armeventgrid.PartnerDestinationsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method PartnerDestinationsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, partnerDestinationName string, options *armeventgrid.PartnerDestinationsClientBeginDeleteOptions) (resp azfake.PollerResponder[armeventgrid.PartnerDestinationsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method PartnerDestinationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, partnerDestinationName string, options *armeventgrid.PartnerDestinationsClientGetOptions) (resp azfake.Responder[armeventgrid.PartnerDestinationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method PartnerDestinationsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armeventgrid.PartnerDestinationsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armeventgrid.PartnerDestinationsClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method PartnerDestinationsClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armeventgrid.PartnerDestinationsClientListBySubscriptionOptions) (resp azfake.PagerResponder[armeventgrid.PartnerDestinationsClientListBySubscriptionResponse])

	// BeginUpdate is the fake for method PartnerDestinationsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, partnerDestinationName string, partnerDestinationUpdateParameters armeventgrid.PartnerDestinationUpdateParameters, options *armeventgrid.PartnerDestinationsClientBeginUpdateOptions) (resp azfake.PollerResponder[armeventgrid.PartnerDestinationsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewPartnerDestinationsServerTransport creates a new instance of PartnerDestinationsServerTransport with the provided implementation.
// The returned PartnerDestinationsServerTransport instance is connected to an instance of armeventgrid.PartnerDestinationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPartnerDestinationsServerTransport(srv *PartnerDestinationsServer) *PartnerDestinationsServerTransport {
	return &PartnerDestinationsServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armeventgrid.PartnerDestinationsClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armeventgrid.PartnerDestinationsClientDeleteResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armeventgrid.PartnerDestinationsClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armeventgrid.PartnerDestinationsClientListBySubscriptionResponse]](),
		beginUpdate:                 newTracker[azfake.PollerResponder[armeventgrid.PartnerDestinationsClientUpdateResponse]](),
	}
}

// PartnerDestinationsServerTransport connects instances of armeventgrid.PartnerDestinationsClient to instances of PartnerDestinationsServer.
// Don't use this type directly, use NewPartnerDestinationsServerTransport instead.
type PartnerDestinationsServerTransport struct {
	srv                         *PartnerDestinationsServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armeventgrid.PartnerDestinationsClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armeventgrid.PartnerDestinationsClientDeleteResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armeventgrid.PartnerDestinationsClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armeventgrid.PartnerDestinationsClientListBySubscriptionResponse]]
	beginUpdate                 *tracker[azfake.PollerResponder[armeventgrid.PartnerDestinationsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for PartnerDestinationsServerTransport.
func (p *PartnerDestinationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PartnerDestinationsClient.Activate":
		resp, err = p.dispatchActivate(req)
	case "PartnerDestinationsClient.BeginCreateOrUpdate":
		resp, err = p.dispatchBeginCreateOrUpdate(req)
	case "PartnerDestinationsClient.BeginDelete":
		resp, err = p.dispatchBeginDelete(req)
	case "PartnerDestinationsClient.Get":
		resp, err = p.dispatchGet(req)
	case "PartnerDestinationsClient.NewListByResourceGroupPager":
		resp, err = p.dispatchNewListByResourceGroupPager(req)
	case "PartnerDestinationsClient.NewListBySubscriptionPager":
		resp, err = p.dispatchNewListBySubscriptionPager(req)
	case "PartnerDestinationsClient.BeginUpdate":
		resp, err = p.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PartnerDestinationsServerTransport) dispatchActivate(req *http.Request) (*http.Response, error) {
	if p.srv.Activate == nil {
		return nil, &nonRetriableError{errors.New("fake for method Activate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventGrid/partnerDestinations/(?P<partnerDestinationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/activate`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	partnerDestinationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("partnerDestinationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Activate(req.Context(), resourceGroupNameParam, partnerDestinationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PartnerDestination, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PartnerDestinationsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := p.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventGrid/partnerDestinations/(?P<partnerDestinationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armeventgrid.PartnerDestination](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		partnerDestinationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("partnerDestinationName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, partnerDestinationNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		p.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		p.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		p.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (p *PartnerDestinationsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := p.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventGrid/partnerDestinations/(?P<partnerDestinationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		partnerDestinationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("partnerDestinationName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginDelete(req.Context(), resourceGroupNameParam, partnerDestinationNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		p.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		p.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		p.beginDelete.remove(req)
	}

	return resp, nil
}

func (p *PartnerDestinationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventGrid/partnerDestinations/(?P<partnerDestinationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	partnerDestinationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("partnerDestinationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, partnerDestinationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PartnerDestination, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PartnerDestinationsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := p.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventGrid/partnerDestinations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
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
		var options *armeventgrid.PartnerDestinationsClientListByResourceGroupOptions
		if filterParam != nil || topParam != nil {
			options = &armeventgrid.PartnerDestinationsClientListByResourceGroupOptions{
				Filter: filterParam,
				Top:    topParam,
			}
		}
		resp := p.srv.NewListByResourceGroupPager(resourceGroupNameParam, options)
		newListByResourceGroupPager = &resp
		p.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armeventgrid.PartnerDestinationsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		p.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (p *PartnerDestinationsServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := p.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventGrid/partnerDestinations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
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
		var options *armeventgrid.PartnerDestinationsClientListBySubscriptionOptions
		if filterParam != nil || topParam != nil {
			options = &armeventgrid.PartnerDestinationsClientListBySubscriptionOptions{
				Filter: filterParam,
				Top:    topParam,
			}
		}
		resp := p.srv.NewListBySubscriptionPager(options)
		newListBySubscriptionPager = &resp
		p.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armeventgrid.PartnerDestinationsClientListBySubscriptionResponse, createLink func() string) {
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

func (p *PartnerDestinationsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := p.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventGrid/partnerDestinations/(?P<partnerDestinationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armeventgrid.PartnerDestinationUpdateParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		partnerDestinationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("partnerDestinationName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginUpdate(req.Context(), resourceGroupNameParam, partnerDestinationNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		p.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		p.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		p.beginUpdate.remove(req)
	}

	return resp, nil
}
