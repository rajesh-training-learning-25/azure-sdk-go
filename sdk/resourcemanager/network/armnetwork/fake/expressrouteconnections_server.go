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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
	"net/http"
	"net/url"
	"regexp"
)

// ExpressRouteConnectionsServer is a fake server for instances of the armnetwork.ExpressRouteConnectionsClient type.
type ExpressRouteConnectionsServer struct {
	// BeginCreateOrUpdate is the fake for method ExpressRouteConnectionsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, connectionName string, putExpressRouteConnectionParameters armnetwork.ExpressRouteConnection, options *armnetwork.ExpressRouteConnectionsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.ExpressRouteConnectionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ExpressRouteConnectionsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, connectionName string, options *armnetwork.ExpressRouteConnectionsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.ExpressRouteConnectionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ExpressRouteConnectionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, connectionName string, options *armnetwork.ExpressRouteConnectionsClientGetOptions) (resp azfake.Responder[armnetwork.ExpressRouteConnectionsClientGetResponse], errResp azfake.ErrorResponder)

	// List is the fake for method ExpressRouteConnectionsClient.List
	// HTTP status codes to indicate success: http.StatusOK
	List func(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, options *armnetwork.ExpressRouteConnectionsClientListOptions) (resp azfake.Responder[armnetwork.ExpressRouteConnectionsClientListResponse], errResp azfake.ErrorResponder)
}

// NewExpressRouteConnectionsServerTransport creates a new instance of ExpressRouteConnectionsServerTransport with the provided implementation.
// The returned ExpressRouteConnectionsServerTransport instance is connected to an instance of armnetwork.ExpressRouteConnectionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewExpressRouteConnectionsServerTransport(srv *ExpressRouteConnectionsServer) *ExpressRouteConnectionsServerTransport {
	return &ExpressRouteConnectionsServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armnetwork.ExpressRouteConnectionsClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armnetwork.ExpressRouteConnectionsClientDeleteResponse]](),
	}
}

// ExpressRouteConnectionsServerTransport connects instances of armnetwork.ExpressRouteConnectionsClient to instances of ExpressRouteConnectionsServer.
// Don't use this type directly, use NewExpressRouteConnectionsServerTransport instead.
type ExpressRouteConnectionsServerTransport struct {
	srv                 *ExpressRouteConnectionsServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armnetwork.ExpressRouteConnectionsClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armnetwork.ExpressRouteConnectionsClientDeleteResponse]]
}

// Do implements the policy.Transporter interface for ExpressRouteConnectionsServerTransport.
func (e *ExpressRouteConnectionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ExpressRouteConnectionsClient.BeginCreateOrUpdate":
		resp, err = e.dispatchBeginCreateOrUpdate(req)
	case "ExpressRouteConnectionsClient.BeginDelete":
		resp, err = e.dispatchBeginDelete(req)
	case "ExpressRouteConnectionsClient.Get":
		resp, err = e.dispatchGet(req)
	case "ExpressRouteConnectionsClient.List":
		resp, err = e.dispatchList(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *ExpressRouteConnectionsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if e.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := e.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/expressRouteGateways/(?P<expressRouteGatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/expressRouteConnections/(?P<connectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.ExpressRouteConnection](req)
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
		connectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("connectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameUnescaped, expressRouteGatewayNameUnescaped, connectionNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		e.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		e.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		e.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (e *ExpressRouteConnectionsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if e.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := e.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/expressRouteGateways/(?P<expressRouteGatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/expressRouteConnections/(?P<connectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
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
		connectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("connectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, expressRouteGatewayNameUnescaped, connectionNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		e.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		e.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		e.beginDelete.remove(req)
	}

	return resp, nil
}

func (e *ExpressRouteConnectionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if e.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/expressRouteGateways/(?P<expressRouteGatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/expressRouteConnections/(?P<connectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
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
	connectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("connectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Get(req.Context(), resourceGroupNameUnescaped, expressRouteGatewayNameUnescaped, connectionNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ExpressRouteConnection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExpressRouteConnectionsServerTransport) dispatchList(req *http.Request) (*http.Response, error) {
	if e.srv.List == nil {
		return nil, &nonRetriableError{errors.New("fake for method List not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/expressRouteGateways/(?P<expressRouteGatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/expressRouteConnections`
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
	respr, errRespr := e.srv.List(req.Context(), resourceGroupNameUnescaped, expressRouteGatewayNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ExpressRouteConnectionList, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
