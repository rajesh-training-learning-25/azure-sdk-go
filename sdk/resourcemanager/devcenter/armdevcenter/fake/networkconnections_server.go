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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// NetworkConnectionsServer is a fake server for instances of the armdevcenter.NetworkConnectionsClient type.
type NetworkConnectionsServer struct {
	// BeginCreateOrUpdate is the fake for method NetworkConnectionsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, networkConnectionName string, body armdevcenter.NetworkConnection, options *armdevcenter.NetworkConnectionsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armdevcenter.NetworkConnectionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method NetworkConnectionsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, networkConnectionName string, options *armdevcenter.NetworkConnectionsClientBeginDeleteOptions) (resp azfake.PollerResponder[armdevcenter.NetworkConnectionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method NetworkConnectionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, networkConnectionName string, options *armdevcenter.NetworkConnectionsClientGetOptions) (resp azfake.Responder[armdevcenter.NetworkConnectionsClientGetResponse], errResp azfake.ErrorResponder)

	// GetHealthDetails is the fake for method NetworkConnectionsClient.GetHealthDetails
	// HTTP status codes to indicate success: http.StatusOK
	GetHealthDetails func(ctx context.Context, resourceGroupName string, networkConnectionName string, options *armdevcenter.NetworkConnectionsClientGetHealthDetailsOptions) (resp azfake.Responder[armdevcenter.NetworkConnectionsClientGetHealthDetailsResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method NetworkConnectionsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armdevcenter.NetworkConnectionsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armdevcenter.NetworkConnectionsClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method NetworkConnectionsClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armdevcenter.NetworkConnectionsClientListBySubscriptionOptions) (resp azfake.PagerResponder[armdevcenter.NetworkConnectionsClientListBySubscriptionResponse])

	// NewListHealthDetailsPager is the fake for method NetworkConnectionsClient.NewListHealthDetailsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListHealthDetailsPager func(resourceGroupName string, networkConnectionName string, options *armdevcenter.NetworkConnectionsClientListHealthDetailsOptions) (resp azfake.PagerResponder[armdevcenter.NetworkConnectionsClientListHealthDetailsResponse])

	// NewListOutboundNetworkDependenciesEndpointsPager is the fake for method NetworkConnectionsClient.NewListOutboundNetworkDependenciesEndpointsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListOutboundNetworkDependenciesEndpointsPager func(resourceGroupName string, networkConnectionName string, options *armdevcenter.NetworkConnectionsClientListOutboundNetworkDependenciesEndpointsOptions) (resp azfake.PagerResponder[armdevcenter.NetworkConnectionsClientListOutboundNetworkDependenciesEndpointsResponse])

	// BeginRunHealthChecks is the fake for method NetworkConnectionsClient.BeginRunHealthChecks
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginRunHealthChecks func(ctx context.Context, resourceGroupName string, networkConnectionName string, options *armdevcenter.NetworkConnectionsClientBeginRunHealthChecksOptions) (resp azfake.PollerResponder[armdevcenter.NetworkConnectionsClientRunHealthChecksResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method NetworkConnectionsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, networkConnectionName string, body armdevcenter.NetworkConnectionUpdate, options *armdevcenter.NetworkConnectionsClientBeginUpdateOptions) (resp azfake.PollerResponder[armdevcenter.NetworkConnectionsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewNetworkConnectionsServerTransport creates a new instance of NetworkConnectionsServerTransport with the provided implementation.
// The returned NetworkConnectionsServerTransport instance is connected to an instance of armdevcenter.NetworkConnectionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewNetworkConnectionsServerTransport(srv *NetworkConnectionsServer) *NetworkConnectionsServerTransport {
	return &NetworkConnectionsServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armdevcenter.NetworkConnectionsClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armdevcenter.NetworkConnectionsClientDeleteResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armdevcenter.NetworkConnectionsClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armdevcenter.NetworkConnectionsClientListBySubscriptionResponse]](),
		newListHealthDetailsPager:   newTracker[azfake.PagerResponder[armdevcenter.NetworkConnectionsClientListHealthDetailsResponse]](),
		newListOutboundNetworkDependenciesEndpointsPager: newTracker[azfake.PagerResponder[armdevcenter.NetworkConnectionsClientListOutboundNetworkDependenciesEndpointsResponse]](),
		beginRunHealthChecks:                             newTracker[azfake.PollerResponder[armdevcenter.NetworkConnectionsClientRunHealthChecksResponse]](),
		beginUpdate:                                      newTracker[azfake.PollerResponder[armdevcenter.NetworkConnectionsClientUpdateResponse]](),
	}
}

// NetworkConnectionsServerTransport connects instances of armdevcenter.NetworkConnectionsClient to instances of NetworkConnectionsServer.
// Don't use this type directly, use NewNetworkConnectionsServerTransport instead.
type NetworkConnectionsServerTransport struct {
	srv                                              *NetworkConnectionsServer
	beginCreateOrUpdate                              *tracker[azfake.PollerResponder[armdevcenter.NetworkConnectionsClientCreateOrUpdateResponse]]
	beginDelete                                      *tracker[azfake.PollerResponder[armdevcenter.NetworkConnectionsClientDeleteResponse]]
	newListByResourceGroupPager                      *tracker[azfake.PagerResponder[armdevcenter.NetworkConnectionsClientListByResourceGroupResponse]]
	newListBySubscriptionPager                       *tracker[azfake.PagerResponder[armdevcenter.NetworkConnectionsClientListBySubscriptionResponse]]
	newListHealthDetailsPager                        *tracker[azfake.PagerResponder[armdevcenter.NetworkConnectionsClientListHealthDetailsResponse]]
	newListOutboundNetworkDependenciesEndpointsPager *tracker[azfake.PagerResponder[armdevcenter.NetworkConnectionsClientListOutboundNetworkDependenciesEndpointsResponse]]
	beginRunHealthChecks                             *tracker[azfake.PollerResponder[armdevcenter.NetworkConnectionsClientRunHealthChecksResponse]]
	beginUpdate                                      *tracker[azfake.PollerResponder[armdevcenter.NetworkConnectionsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for NetworkConnectionsServerTransport.
func (n *NetworkConnectionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "NetworkConnectionsClient.BeginCreateOrUpdate":
		resp, err = n.dispatchBeginCreateOrUpdate(req)
	case "NetworkConnectionsClient.BeginDelete":
		resp, err = n.dispatchBeginDelete(req)
	case "NetworkConnectionsClient.Get":
		resp, err = n.dispatchGet(req)
	case "NetworkConnectionsClient.GetHealthDetails":
		resp, err = n.dispatchGetHealthDetails(req)
	case "NetworkConnectionsClient.NewListByResourceGroupPager":
		resp, err = n.dispatchNewListByResourceGroupPager(req)
	case "NetworkConnectionsClient.NewListBySubscriptionPager":
		resp, err = n.dispatchNewListBySubscriptionPager(req)
	case "NetworkConnectionsClient.NewListHealthDetailsPager":
		resp, err = n.dispatchNewListHealthDetailsPager(req)
	case "NetworkConnectionsClient.NewListOutboundNetworkDependenciesEndpointsPager":
		resp, err = n.dispatchNewListOutboundNetworkDependenciesEndpointsPager(req)
	case "NetworkConnectionsClient.BeginRunHealthChecks":
		resp, err = n.dispatchBeginRunHealthChecks(req)
	case "NetworkConnectionsClient.BeginUpdate":
		resp, err = n.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (n *NetworkConnectionsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if n.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := n.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/networkConnections/(?P<networkConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdevcenter.NetworkConnection](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkConnectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, networkConnectionNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		n.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		n.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		n.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (n *NetworkConnectionsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if n.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := n.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/networkConnections/(?P<networkConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkConnectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginDelete(req.Context(), resourceGroupNameParam, networkConnectionNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		n.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		n.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		n.beginDelete.remove(req)
	}

	return resp, nil
}

func (n *NetworkConnectionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if n.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/networkConnections/(?P<networkConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.Get(req.Context(), resourceGroupNameParam, networkConnectionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NetworkConnection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NetworkConnectionsServerTransport) dispatchGetHealthDetails(req *http.Request) (*http.Response, error) {
	if n.srv.GetHealthDetails == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetHealthDetails not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/networkConnections/(?P<networkConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/healthChecks/latest`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.GetHealthDetails(req.Context(), resourceGroupNameParam, networkConnectionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).HealthCheckStatusDetails, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NetworkConnectionsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if n.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := n.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/networkConnections`
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
		var options *armdevcenter.NetworkConnectionsClientListByResourceGroupOptions
		if topParam != nil {
			options = &armdevcenter.NetworkConnectionsClientListByResourceGroupOptions{
				Top: topParam,
			}
		}
		resp := n.srv.NewListByResourceGroupPager(resourceGroupNameParam, options)
		newListByResourceGroupPager = &resp
		n.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armdevcenter.NetworkConnectionsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		n.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		n.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (n *NetworkConnectionsServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if n.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := n.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/networkConnections`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
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
		var options *armdevcenter.NetworkConnectionsClientListBySubscriptionOptions
		if topParam != nil {
			options = &armdevcenter.NetworkConnectionsClientListBySubscriptionOptions{
				Top: topParam,
			}
		}
		resp := n.srv.NewListBySubscriptionPager(options)
		newListBySubscriptionPager = &resp
		n.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armdevcenter.NetworkConnectionsClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		n.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		n.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (n *NetworkConnectionsServerTransport) dispatchNewListHealthDetailsPager(req *http.Request) (*http.Response, error) {
	if n.srv.NewListHealthDetailsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListHealthDetailsPager not implemented")}
	}
	newListHealthDetailsPager := n.newListHealthDetailsPager.get(req)
	if newListHealthDetailsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/networkConnections/(?P<networkConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/healthChecks`
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
		networkConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkConnectionName")])
		if err != nil {
			return nil, err
		}
		var options *armdevcenter.NetworkConnectionsClientListHealthDetailsOptions
		if topParam != nil {
			options = &armdevcenter.NetworkConnectionsClientListHealthDetailsOptions{
				Top: topParam,
			}
		}
		resp := n.srv.NewListHealthDetailsPager(resourceGroupNameParam, networkConnectionNameParam, options)
		newListHealthDetailsPager = &resp
		n.newListHealthDetailsPager.add(req, newListHealthDetailsPager)
		server.PagerResponderInjectNextLinks(newListHealthDetailsPager, req, func(page *armdevcenter.NetworkConnectionsClientListHealthDetailsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListHealthDetailsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		n.newListHealthDetailsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListHealthDetailsPager) {
		n.newListHealthDetailsPager.remove(req)
	}
	return resp, nil
}

func (n *NetworkConnectionsServerTransport) dispatchNewListOutboundNetworkDependenciesEndpointsPager(req *http.Request) (*http.Response, error) {
	if n.srv.NewListOutboundNetworkDependenciesEndpointsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListOutboundNetworkDependenciesEndpointsPager not implemented")}
	}
	newListOutboundNetworkDependenciesEndpointsPager := n.newListOutboundNetworkDependenciesEndpointsPager.get(req)
	if newListOutboundNetworkDependenciesEndpointsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/networkConnections/(?P<networkConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/outboundNetworkDependenciesEndpoints`
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
		networkConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkConnectionName")])
		if err != nil {
			return nil, err
		}
		var options *armdevcenter.NetworkConnectionsClientListOutboundNetworkDependenciesEndpointsOptions
		if topParam != nil {
			options = &armdevcenter.NetworkConnectionsClientListOutboundNetworkDependenciesEndpointsOptions{
				Top: topParam,
			}
		}
		resp := n.srv.NewListOutboundNetworkDependenciesEndpointsPager(resourceGroupNameParam, networkConnectionNameParam, options)
		newListOutboundNetworkDependenciesEndpointsPager = &resp
		n.newListOutboundNetworkDependenciesEndpointsPager.add(req, newListOutboundNetworkDependenciesEndpointsPager)
		server.PagerResponderInjectNextLinks(newListOutboundNetworkDependenciesEndpointsPager, req, func(page *armdevcenter.NetworkConnectionsClientListOutboundNetworkDependenciesEndpointsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListOutboundNetworkDependenciesEndpointsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		n.newListOutboundNetworkDependenciesEndpointsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListOutboundNetworkDependenciesEndpointsPager) {
		n.newListOutboundNetworkDependenciesEndpointsPager.remove(req)
	}
	return resp, nil
}

func (n *NetworkConnectionsServerTransport) dispatchBeginRunHealthChecks(req *http.Request) (*http.Response, error) {
	if n.srv.BeginRunHealthChecks == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginRunHealthChecks not implemented")}
	}
	beginRunHealthChecks := n.beginRunHealthChecks.get(req)
	if beginRunHealthChecks == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/networkConnections/(?P<networkConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runHealthChecks`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkConnectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginRunHealthChecks(req.Context(), resourceGroupNameParam, networkConnectionNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginRunHealthChecks = &respr
		n.beginRunHealthChecks.add(req, beginRunHealthChecks)
	}

	resp, err := server.PollerResponderNext(beginRunHealthChecks, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		n.beginRunHealthChecks.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginRunHealthChecks) {
		n.beginRunHealthChecks.remove(req)
	}

	return resp, nil
}

func (n *NetworkConnectionsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if n.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := n.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/networkConnections/(?P<networkConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdevcenter.NetworkConnectionUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkConnectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginUpdate(req.Context(), resourceGroupNameParam, networkConnectionNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		n.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		n.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		n.beginUpdate.remove(req)
	}

	return resp, nil
}
