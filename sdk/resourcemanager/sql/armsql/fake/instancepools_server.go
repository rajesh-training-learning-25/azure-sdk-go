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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
	"net/http"
	"net/url"
	"regexp"
)

// InstancePoolsServer is a fake server for instances of the armsql.InstancePoolsClient type.
type InstancePoolsServer struct {
	// BeginCreateOrUpdate is the fake for method InstancePoolsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, instancePoolName string, parameters armsql.InstancePool, options *armsql.InstancePoolsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armsql.InstancePoolsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method InstancePoolsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, instancePoolName string, options *armsql.InstancePoolsClientBeginDeleteOptions) (resp azfake.PollerResponder[armsql.InstancePoolsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method InstancePoolsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, instancePoolName string, options *armsql.InstancePoolsClientGetOptions) (resp azfake.Responder[armsql.InstancePoolsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method InstancePoolsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armsql.InstancePoolsClientListOptions) (resp azfake.PagerResponder[armsql.InstancePoolsClientListResponse])

	// NewListByResourceGroupPager is the fake for method InstancePoolsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armsql.InstancePoolsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armsql.InstancePoolsClientListByResourceGroupResponse])

	// BeginUpdate is the fake for method InstancePoolsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, instancePoolName string, parameters armsql.InstancePoolUpdate, options *armsql.InstancePoolsClientBeginUpdateOptions) (resp azfake.PollerResponder[armsql.InstancePoolsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewInstancePoolsServerTransport creates a new instance of InstancePoolsServerTransport with the provided implementation.
// The returned InstancePoolsServerTransport instance is connected to an instance of armsql.InstancePoolsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewInstancePoolsServerTransport(srv *InstancePoolsServer) *InstancePoolsServerTransport {
	return &InstancePoolsServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armsql.InstancePoolsClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armsql.InstancePoolsClientDeleteResponse]](),
		newListPager:                newTracker[azfake.PagerResponder[armsql.InstancePoolsClientListResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armsql.InstancePoolsClientListByResourceGroupResponse]](),
		beginUpdate:                 newTracker[azfake.PollerResponder[armsql.InstancePoolsClientUpdateResponse]](),
	}
}

// InstancePoolsServerTransport connects instances of armsql.InstancePoolsClient to instances of InstancePoolsServer.
// Don't use this type directly, use NewInstancePoolsServerTransport instead.
type InstancePoolsServerTransport struct {
	srv                         *InstancePoolsServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armsql.InstancePoolsClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armsql.InstancePoolsClientDeleteResponse]]
	newListPager                *tracker[azfake.PagerResponder[armsql.InstancePoolsClientListResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armsql.InstancePoolsClientListByResourceGroupResponse]]
	beginUpdate                 *tracker[azfake.PollerResponder[armsql.InstancePoolsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for InstancePoolsServerTransport.
func (i *InstancePoolsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "InstancePoolsClient.BeginCreateOrUpdate":
		resp, err = i.dispatchBeginCreateOrUpdate(req)
	case "InstancePoolsClient.BeginDelete":
		resp, err = i.dispatchBeginDelete(req)
	case "InstancePoolsClient.Get":
		resp, err = i.dispatchGet(req)
	case "InstancePoolsClient.NewListPager":
		resp, err = i.dispatchNewListPager(req)
	case "InstancePoolsClient.NewListByResourceGroupPager":
		resp, err = i.dispatchNewListByResourceGroupPager(req)
	case "InstancePoolsClient.BeginUpdate":
		resp, err = i.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *InstancePoolsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if i.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := i.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/instancePools/(?P<instancePoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsql.InstancePool](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		instancePoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("instancePoolName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := i.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, instancePoolNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		i.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		i.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		i.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (i *InstancePoolsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if i.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := i.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/instancePools/(?P<instancePoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		instancePoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("instancePoolName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := i.srv.BeginDelete(req.Context(), resourceGroupNameParam, instancePoolNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		i.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		i.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		i.beginDelete.remove(req)
	}

	return resp, nil
}

func (i *InstancePoolsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if i.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/instancePools/(?P<instancePoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	instancePoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("instancePoolName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Get(req.Context(), resourceGroupNameParam, instancePoolNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).InstancePool, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *InstancePoolsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := i.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/instancePools`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := i.srv.NewListPager(nil)
		newListPager = &resp
		i.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armsql.InstancePoolsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		i.newListPager.remove(req)
	}
	return resp, nil
}

func (i *InstancePoolsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := i.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/instancePools`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := i.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		i.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armsql.InstancePoolsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		i.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (i *InstancePoolsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if i.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := i.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/instancePools/(?P<instancePoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsql.InstancePoolUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		instancePoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("instancePoolName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := i.srv.BeginUpdate(req.Context(), resourceGroupNameParam, instancePoolNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		i.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		i.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		i.beginUpdate.remove(req)
	}

	return resp, nil
}
