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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurearcdata/armazurearcdata"
	"net/http"
	"net/url"
	"regexp"
)

// PostgresInstancesServer is a fake server for instances of the armazurearcdata.PostgresInstancesClient type.
type PostgresInstancesServer struct {
	// BeginCreate is the fake for method PostgresInstancesClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, postgresInstanceName string, resource armazurearcdata.PostgresInstance, options *armazurearcdata.PostgresInstancesClientBeginCreateOptions) (resp azfake.PollerResponder[armazurearcdata.PostgresInstancesClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method PostgresInstancesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, postgresInstanceName string, options *armazurearcdata.PostgresInstancesClientBeginDeleteOptions) (resp azfake.PollerResponder[armazurearcdata.PostgresInstancesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method PostgresInstancesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, postgresInstanceName string, options *armazurearcdata.PostgresInstancesClientGetOptions) (resp azfake.Responder[armazurearcdata.PostgresInstancesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method PostgresInstancesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armazurearcdata.PostgresInstancesClientListOptions) (resp azfake.PagerResponder[armazurearcdata.PostgresInstancesClientListResponse])

	// NewListByResourceGroupPager is the fake for method PostgresInstancesClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armazurearcdata.PostgresInstancesClientListByResourceGroupOptions) (resp azfake.PagerResponder[armazurearcdata.PostgresInstancesClientListByResourceGroupResponse])

	// Update is the fake for method PostgresInstancesClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, postgresInstanceName string, parameters armazurearcdata.PostgresInstanceUpdate, options *armazurearcdata.PostgresInstancesClientUpdateOptions) (resp azfake.Responder[armazurearcdata.PostgresInstancesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewPostgresInstancesServerTransport creates a new instance of PostgresInstancesServerTransport with the provided implementation.
// The returned PostgresInstancesServerTransport instance is connected to an instance of armazurearcdata.PostgresInstancesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPostgresInstancesServerTransport(srv *PostgresInstancesServer) *PostgresInstancesServerTransport {
	return &PostgresInstancesServerTransport{
		srv:                         srv,
		beginCreate:                 newTracker[azfake.PollerResponder[armazurearcdata.PostgresInstancesClientCreateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armazurearcdata.PostgresInstancesClientDeleteResponse]](),
		newListPager:                newTracker[azfake.PagerResponder[armazurearcdata.PostgresInstancesClientListResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armazurearcdata.PostgresInstancesClientListByResourceGroupResponse]](),
	}
}

// PostgresInstancesServerTransport connects instances of armazurearcdata.PostgresInstancesClient to instances of PostgresInstancesServer.
// Don't use this type directly, use NewPostgresInstancesServerTransport instead.
type PostgresInstancesServerTransport struct {
	srv                         *PostgresInstancesServer
	beginCreate                 *tracker[azfake.PollerResponder[armazurearcdata.PostgresInstancesClientCreateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armazurearcdata.PostgresInstancesClientDeleteResponse]]
	newListPager                *tracker[azfake.PagerResponder[armazurearcdata.PostgresInstancesClientListResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armazurearcdata.PostgresInstancesClientListByResourceGroupResponse]]
}

// Do implements the policy.Transporter interface for PostgresInstancesServerTransport.
func (p *PostgresInstancesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PostgresInstancesClient.BeginCreate":
		resp, err = p.dispatchBeginCreate(req)
	case "PostgresInstancesClient.BeginDelete":
		resp, err = p.dispatchBeginDelete(req)
	case "PostgresInstancesClient.Get":
		resp, err = p.dispatchGet(req)
	case "PostgresInstancesClient.NewListPager":
		resp, err = p.dispatchNewListPager(req)
	case "PostgresInstancesClient.NewListByResourceGroupPager":
		resp, err = p.dispatchNewListByResourceGroupPager(req)
	case "PostgresInstancesClient.Update":
		resp, err = p.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PostgresInstancesServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if p.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := p.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureArcData/postgresInstances/(?P<postgresInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armazurearcdata.PostgresInstance](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		postgresInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("postgresInstanceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginCreate(req.Context(), resourceGroupNameParam, postgresInstanceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		p.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		p.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		p.beginCreate.remove(req)
	}

	return resp, nil
}

func (p *PostgresInstancesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := p.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureArcData/postgresInstances/(?P<postgresInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		postgresInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("postgresInstanceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginDelete(req.Context(), resourceGroupNameParam, postgresInstanceNameParam, nil)
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

func (p *PostgresInstancesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureArcData/postgresInstances/(?P<postgresInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	postgresInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("postgresInstanceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, postgresInstanceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PostgresInstance, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PostgresInstancesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := p.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureArcData/postgresInstances`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := p.srv.NewListPager(nil)
		newListPager = &resp
		p.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armazurearcdata.PostgresInstancesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		p.newListPager.remove(req)
	}
	return resp, nil
}

func (p *PostgresInstancesServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := p.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureArcData/postgresInstances`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		p.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armazurearcdata.PostgresInstancesClientListByResourceGroupResponse, createLink func() string) {
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

func (p *PostgresInstancesServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureArcData/postgresInstances/(?P<postgresInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armazurearcdata.PostgresInstanceUpdate](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	postgresInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("postgresInstanceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Update(req.Context(), resourceGroupNameParam, postgresInstanceNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PostgresInstance, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}