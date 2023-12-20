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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/m365securityandcompliance/armm365securityandcompliance"
	"net/http"
	"net/url"
	"regexp"
)

// PrivateLinkResourcesForSCCPowershellServer is a fake server for instances of the armm365securityandcompliance.PrivateLinkResourcesForSCCPowershellClient type.
type PrivateLinkResourcesForSCCPowershellServer struct {
	// Get is the fake for method PrivateLinkResourcesForSCCPowershellClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, resourceName string, groupName string, options *armm365securityandcompliance.PrivateLinkResourcesForSCCPowershellClientGetOptions) (resp azfake.Responder[armm365securityandcompliance.PrivateLinkResourcesForSCCPowershellClientGetResponse], errResp azfake.ErrorResponder)

	// ListByService is the fake for method PrivateLinkResourcesForSCCPowershellClient.ListByService
	// HTTP status codes to indicate success: http.StatusOK
	ListByService func(ctx context.Context, resourceGroupName string, resourceName string, options *armm365securityandcompliance.PrivateLinkResourcesForSCCPowershellClientListByServiceOptions) (resp azfake.Responder[armm365securityandcompliance.PrivateLinkResourcesForSCCPowershellClientListByServiceResponse], errResp azfake.ErrorResponder)
}

// NewPrivateLinkResourcesForSCCPowershellServerTransport creates a new instance of PrivateLinkResourcesForSCCPowershellServerTransport with the provided implementation.
// The returned PrivateLinkResourcesForSCCPowershellServerTransport instance is connected to an instance of armm365securityandcompliance.PrivateLinkResourcesForSCCPowershellClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPrivateLinkResourcesForSCCPowershellServerTransport(srv *PrivateLinkResourcesForSCCPowershellServer) *PrivateLinkResourcesForSCCPowershellServerTransport {
	return &PrivateLinkResourcesForSCCPowershellServerTransport{srv: srv}
}

// PrivateLinkResourcesForSCCPowershellServerTransport connects instances of armm365securityandcompliance.PrivateLinkResourcesForSCCPowershellClient to instances of PrivateLinkResourcesForSCCPowershellServer.
// Don't use this type directly, use NewPrivateLinkResourcesForSCCPowershellServerTransport instead.
type PrivateLinkResourcesForSCCPowershellServerTransport struct {
	srv *PrivateLinkResourcesForSCCPowershellServer
}

// Do implements the policy.Transporter interface for PrivateLinkResourcesForSCCPowershellServerTransport.
func (p *PrivateLinkResourcesForSCCPowershellServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PrivateLinkResourcesForSCCPowershellClient.Get":
		resp, err = p.dispatchGet(req)
	case "PrivateLinkResourcesForSCCPowershellClient.ListByService":
		resp, err = p.dispatchListByService(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PrivateLinkResourcesForSCCPowershellServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.M365SecurityAndCompliance/privateLinkServicesForSCCPowershell/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateLinkResources/(?P<groupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	groupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("groupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, resourceNameParam, groupNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateLinkResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrivateLinkResourcesForSCCPowershellServerTransport) dispatchListByService(req *http.Request) (*http.Response, error) {
	if p.srv.ListByService == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListByService not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.M365SecurityAndCompliance/privateLinkServicesForSCCPowershell/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateLinkResources`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.ListByService(req.Context(), resourceGroupNameParam, resourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateLinkResourceListResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
