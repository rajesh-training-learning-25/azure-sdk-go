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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
)

// IntegrationRuntimeObjectMetadataServer is a fake server for instances of the armsynapse.IntegrationRuntimeObjectMetadataClient type.
type IntegrationRuntimeObjectMetadataServer struct {
	// List is the fake for method IntegrationRuntimeObjectMetadataClient.List
	// HTTP status codes to indicate success: http.StatusOK
	List func(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, options *armsynapse.IntegrationRuntimeObjectMetadataClientListOptions) (resp azfake.Responder[armsynapse.IntegrationRuntimeObjectMetadataClientListResponse], errResp azfake.ErrorResponder)

	// BeginRefresh is the fake for method IntegrationRuntimeObjectMetadataClient.BeginRefresh
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginRefresh func(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, options *armsynapse.IntegrationRuntimeObjectMetadataClientBeginRefreshOptions) (resp azfake.PollerResponder[armsynapse.IntegrationRuntimeObjectMetadataClientRefreshResponse], errResp azfake.ErrorResponder)
}

// NewIntegrationRuntimeObjectMetadataServerTransport creates a new instance of IntegrationRuntimeObjectMetadataServerTransport with the provided implementation.
// The returned IntegrationRuntimeObjectMetadataServerTransport instance is connected to an instance of armsynapse.IntegrationRuntimeObjectMetadataClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewIntegrationRuntimeObjectMetadataServerTransport(srv *IntegrationRuntimeObjectMetadataServer) *IntegrationRuntimeObjectMetadataServerTransport {
	return &IntegrationRuntimeObjectMetadataServerTransport{
		srv:          srv,
		beginRefresh: newTracker[azfake.PollerResponder[armsynapse.IntegrationRuntimeObjectMetadataClientRefreshResponse]](),
	}
}

// IntegrationRuntimeObjectMetadataServerTransport connects instances of armsynapse.IntegrationRuntimeObjectMetadataClient to instances of IntegrationRuntimeObjectMetadataServer.
// Don't use this type directly, use NewIntegrationRuntimeObjectMetadataServerTransport instead.
type IntegrationRuntimeObjectMetadataServerTransport struct {
	srv          *IntegrationRuntimeObjectMetadataServer
	beginRefresh *tracker[azfake.PollerResponder[armsynapse.IntegrationRuntimeObjectMetadataClientRefreshResponse]]
}

// Do implements the policy.Transporter interface for IntegrationRuntimeObjectMetadataServerTransport.
func (i *IntegrationRuntimeObjectMetadataServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "IntegrationRuntimeObjectMetadataClient.List":
		resp, err = i.dispatchList(req)
	case "IntegrationRuntimeObjectMetadataClient.BeginRefresh":
		resp, err = i.dispatchBeginRefresh(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *IntegrationRuntimeObjectMetadataServerTransport) dispatchList(req *http.Request) (*http.Response, error) {
	if i.srv.List == nil {
		return nil, &nonRetriableError{errors.New("fake for method List not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Synapse/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/integrationRuntimes/(?P<integrationRuntimeName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getObjectMetadata`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armsynapse.GetSsisObjectMetadataRequest](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	integrationRuntimeNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("integrationRuntimeName")])
	if err != nil {
		return nil, err
	}
	var options *armsynapse.IntegrationRuntimeObjectMetadataClientListOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &armsynapse.IntegrationRuntimeObjectMetadataClientListOptions{
			GetMetadataRequest: &body,
		}
	}
	respr, errRespr := i.srv.List(req.Context(), resourceGroupNameParam, workspaceNameParam, integrationRuntimeNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SsisObjectMetadataListResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *IntegrationRuntimeObjectMetadataServerTransport) dispatchBeginRefresh(req *http.Request) (*http.Response, error) {
	if i.srv.BeginRefresh == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginRefresh not implemented")}
	}
	beginRefresh := i.beginRefresh.get(req)
	if beginRefresh == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Synapse/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/integrationRuntimes/(?P<integrationRuntimeName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/refreshObjectMetadata`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		integrationRuntimeNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("integrationRuntimeName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := i.srv.BeginRefresh(req.Context(), resourceGroupNameParam, workspaceNameParam, integrationRuntimeNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginRefresh = &respr
		i.beginRefresh.add(req, beginRefresh)
	}

	resp, err := server.PollerResponderNext(beginRefresh, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		i.beginRefresh.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginRefresh) {
		i.beginRefresh.remove(req)
	}

	return resp, nil
}
