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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
	"net/http"
	"net/url"
	"regexp"
)

// IntegrationAccountBatchConfigurationsServer is a fake server for instances of the armlogic.IntegrationAccountBatchConfigurationsClient type.
type IntegrationAccountBatchConfigurationsServer struct {
	// CreateOrUpdate is the fake for method IntegrationAccountBatchConfigurationsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, integrationAccountName string, batchConfigurationName string, batchConfiguration armlogic.BatchConfiguration, options *armlogic.IntegrationAccountBatchConfigurationsClientCreateOrUpdateOptions) (resp azfake.Responder[armlogic.IntegrationAccountBatchConfigurationsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method IntegrationAccountBatchConfigurationsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, integrationAccountName string, batchConfigurationName string, options *armlogic.IntegrationAccountBatchConfigurationsClientDeleteOptions) (resp azfake.Responder[armlogic.IntegrationAccountBatchConfigurationsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method IntegrationAccountBatchConfigurationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, integrationAccountName string, batchConfigurationName string, options *armlogic.IntegrationAccountBatchConfigurationsClientGetOptions) (resp azfake.Responder[armlogic.IntegrationAccountBatchConfigurationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method IntegrationAccountBatchConfigurationsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, integrationAccountName string, options *armlogic.IntegrationAccountBatchConfigurationsClientListOptions) (resp azfake.PagerResponder[armlogic.IntegrationAccountBatchConfigurationsClientListResponse])
}

// NewIntegrationAccountBatchConfigurationsServerTransport creates a new instance of IntegrationAccountBatchConfigurationsServerTransport with the provided implementation.
// The returned IntegrationAccountBatchConfigurationsServerTransport instance is connected to an instance of armlogic.IntegrationAccountBatchConfigurationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewIntegrationAccountBatchConfigurationsServerTransport(srv *IntegrationAccountBatchConfigurationsServer) *IntegrationAccountBatchConfigurationsServerTransport {
	return &IntegrationAccountBatchConfigurationsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armlogic.IntegrationAccountBatchConfigurationsClientListResponse]](),
	}
}

// IntegrationAccountBatchConfigurationsServerTransport connects instances of armlogic.IntegrationAccountBatchConfigurationsClient to instances of IntegrationAccountBatchConfigurationsServer.
// Don't use this type directly, use NewIntegrationAccountBatchConfigurationsServerTransport instead.
type IntegrationAccountBatchConfigurationsServerTransport struct {
	srv          *IntegrationAccountBatchConfigurationsServer
	newListPager *tracker[azfake.PagerResponder[armlogic.IntegrationAccountBatchConfigurationsClientListResponse]]
}

// Do implements the policy.Transporter interface for IntegrationAccountBatchConfigurationsServerTransport.
func (i *IntegrationAccountBatchConfigurationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "IntegrationAccountBatchConfigurationsClient.CreateOrUpdate":
		resp, err = i.dispatchCreateOrUpdate(req)
	case "IntegrationAccountBatchConfigurationsClient.Delete":
		resp, err = i.dispatchDelete(req)
	case "IntegrationAccountBatchConfigurationsClient.Get":
		resp, err = i.dispatchGet(req)
	case "IntegrationAccountBatchConfigurationsClient.NewListPager":
		resp, err = i.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *IntegrationAccountBatchConfigurationsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if i.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Logic/integrationAccounts/(?P<integrationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/batchConfigurations/(?P<batchConfigurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armlogic.BatchConfiguration](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	integrationAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("integrationAccountName")])
	if err != nil {
		return nil, err
	}
	batchConfigurationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("batchConfigurationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, integrationAccountNameParam, batchConfigurationNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).BatchConfiguration, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *IntegrationAccountBatchConfigurationsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if i.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Logic/integrationAccounts/(?P<integrationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/batchConfigurations/(?P<batchConfigurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	integrationAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("integrationAccountName")])
	if err != nil {
		return nil, err
	}
	batchConfigurationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("batchConfigurationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Delete(req.Context(), resourceGroupNameParam, integrationAccountNameParam, batchConfigurationNameParam, nil)
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

func (i *IntegrationAccountBatchConfigurationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if i.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Logic/integrationAccounts/(?P<integrationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/batchConfigurations/(?P<batchConfigurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	integrationAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("integrationAccountName")])
	if err != nil {
		return nil, err
	}
	batchConfigurationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("batchConfigurationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Get(req.Context(), resourceGroupNameParam, integrationAccountNameParam, batchConfigurationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).BatchConfiguration, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *IntegrationAccountBatchConfigurationsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := i.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Logic/integrationAccounts/(?P<integrationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/batchConfigurations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		integrationAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("integrationAccountName")])
		if err != nil {
			return nil, err
		}
		resp := i.srv.NewListPager(resourceGroupNameParam, integrationAccountNameParam, nil)
		newListPager = &resp
		i.newListPager.add(req, newListPager)
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
