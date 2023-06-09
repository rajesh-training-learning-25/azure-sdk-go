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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v3"
	"net/http"
	"net/url"
	"regexp"
)

// ExpressRouteGatewaysServer is a fake server for instances of the armnetwork.ExpressRouteGatewaysClient type.
type ExpressRouteGatewaysServer struct {
	// BeginCreateOrUpdate is the fake for method ExpressRouteGatewaysClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, putExpressRouteGatewayParameters armnetwork.ExpressRouteGateway, options *armnetwork.ExpressRouteGatewaysClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.ExpressRouteGatewaysClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ExpressRouteGatewaysClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, options *armnetwork.ExpressRouteGatewaysClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.ExpressRouteGatewaysClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ExpressRouteGatewaysClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, options *armnetwork.ExpressRouteGatewaysClientGetOptions) (resp azfake.Responder[armnetwork.ExpressRouteGatewaysClientGetResponse], errResp azfake.ErrorResponder)

	// ListByResourceGroup is the fake for method ExpressRouteGatewaysClient.ListByResourceGroup
	// HTTP status codes to indicate success: http.StatusOK
	ListByResourceGroup func(ctx context.Context, resourceGroupName string, options *armnetwork.ExpressRouteGatewaysClientListByResourceGroupOptions) (resp azfake.Responder[armnetwork.ExpressRouteGatewaysClientListByResourceGroupResponse], errResp azfake.ErrorResponder)

	// ListBySubscription is the fake for method ExpressRouteGatewaysClient.ListBySubscription
	// HTTP status codes to indicate success: http.StatusOK
	ListBySubscription func(ctx context.Context, options *armnetwork.ExpressRouteGatewaysClientListBySubscriptionOptions) (resp azfake.Responder[armnetwork.ExpressRouteGatewaysClientListBySubscriptionResponse], errResp azfake.ErrorResponder)

	// BeginUpdateTags is the fake for method ExpressRouteGatewaysClient.BeginUpdateTags
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdateTags func(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, expressRouteGatewayParameters armnetwork.TagsObject, options *armnetwork.ExpressRouteGatewaysClientBeginUpdateTagsOptions) (resp azfake.PollerResponder[armnetwork.ExpressRouteGatewaysClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewExpressRouteGatewaysServerTransport creates a new instance of ExpressRouteGatewaysServerTransport with the provided implementation.
// The returned ExpressRouteGatewaysServerTransport instance is connected to an instance of armnetwork.ExpressRouteGatewaysClient by way of the
// undefined.Transporter field.
func NewExpressRouteGatewaysServerTransport(srv *ExpressRouteGatewaysServer) *ExpressRouteGatewaysServerTransport {
	return &ExpressRouteGatewaysServerTransport{srv: srv}
}

// ExpressRouteGatewaysServerTransport connects instances of armnetwork.ExpressRouteGatewaysClient to instances of ExpressRouteGatewaysServer.
// Don't use this type directly, use NewExpressRouteGatewaysServerTransport instead.
type ExpressRouteGatewaysServerTransport struct {
	srv                 *ExpressRouteGatewaysServer
	beginCreateOrUpdate *azfake.PollerResponder[armnetwork.ExpressRouteGatewaysClientCreateOrUpdateResponse]
	beginDelete         *azfake.PollerResponder[armnetwork.ExpressRouteGatewaysClientDeleteResponse]
	beginUpdateTags     *azfake.PollerResponder[armnetwork.ExpressRouteGatewaysClientUpdateTagsResponse]
}

// Do implements the policy.Transporter interface for ExpressRouteGatewaysServerTransport.
func (e *ExpressRouteGatewaysServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ExpressRouteGatewaysClient.BeginCreateOrUpdate":
		resp, err = e.dispatchBeginCreateOrUpdate(req)
	case "ExpressRouteGatewaysClient.BeginDelete":
		resp, err = e.dispatchBeginDelete(req)
	case "ExpressRouteGatewaysClient.Get":
		resp, err = e.dispatchGet(req)
	case "ExpressRouteGatewaysClient.ListByResourceGroup":
		resp, err = e.dispatchListByResourceGroup(req)
	case "ExpressRouteGatewaysClient.ListBySubscription":
		resp, err = e.dispatchListBySubscription(req)
	case "ExpressRouteGatewaysClient.BeginUpdateTags":
		resp, err = e.dispatchBeginUpdateTags(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *ExpressRouteGatewaysServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if e.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	if e.beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/expressRouteGateways/(?P<expressRouteGatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.ExpressRouteGateway](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		expressRouteGatewayNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("expressRouteGatewayName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameUnescaped, expressRouteGatewayNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		e.beginCreateOrUpdate = &respr
	}

	resp, err := server.PollerResponderNext(e.beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(e.beginCreateOrUpdate) {
		e.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (e *ExpressRouteGatewaysServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if e.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	if e.beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/expressRouteGateways/(?P<expressRouteGatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		expressRouteGatewayNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("expressRouteGatewayName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, expressRouteGatewayNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		e.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(e.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(e.beginDelete) {
		e.beginDelete = nil
	}

	return resp, nil
}

func (e *ExpressRouteGatewaysServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if e.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/expressRouteGateways/(?P<expressRouteGatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	expressRouteGatewayNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("expressRouteGatewayName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Get(req.Context(), resourceGroupNameUnescaped, expressRouteGatewayNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ExpressRouteGateway, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExpressRouteGatewaysServerTransport) dispatchListByResourceGroup(req *http.Request) (*http.Response, error) {
	if e.srv.ListByResourceGroup == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListByResourceGroup not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/expressRouteGateways`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.ListByResourceGroup(req.Context(), resourceGroupNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ExpressRouteGatewayList, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExpressRouteGatewaysServerTransport) dispatchListBySubscription(req *http.Request) (*http.Response, error) {
	if e.srv.ListBySubscription == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListBySubscription not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/expressRouteGateways`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := e.srv.ListBySubscription(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ExpressRouteGatewayList, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExpressRouteGatewaysServerTransport) dispatchBeginUpdateTags(req *http.Request) (*http.Response, error) {
	if e.srv.BeginUpdateTags == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdateTags not implemented")}
	}
	if e.beginUpdateTags == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/expressRouteGateways/(?P<expressRouteGatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.TagsObject](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		expressRouteGatewayNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("expressRouteGatewayName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginUpdateTags(req.Context(), resourceGroupNameUnescaped, expressRouteGatewayNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		e.beginUpdateTags = &respr
	}

	resp, err := server.PollerResponderNext(e.beginUpdateTags, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(e.beginUpdateTags) {
		e.beginUpdateTags = nil
	}

	return resp, nil
}
