//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"
	"net/http"
)

// OrganizationOperationsServer is a fake server for instances of the armconfluent.OrganizationOperationsClient type.
type OrganizationOperationsServer struct {
	// NewListPager is the fake for method OrganizationOperationsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armconfluent.OrganizationOperationsClientListOptions) (resp azfake.PagerResponder[armconfluent.OrganizationOperationsClientListResponse])
}

// NewOrganizationOperationsServerTransport creates a new instance of OrganizationOperationsServerTransport with the provided implementation.
// The returned OrganizationOperationsServerTransport instance is connected to an instance of armconfluent.OrganizationOperationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewOrganizationOperationsServerTransport(srv *OrganizationOperationsServer) *OrganizationOperationsServerTransport {
	return &OrganizationOperationsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armconfluent.OrganizationOperationsClientListResponse]](),
	}
}

// OrganizationOperationsServerTransport connects instances of armconfluent.OrganizationOperationsClient to instances of OrganizationOperationsServer.
// Don't use this type directly, use NewOrganizationOperationsServerTransport instead.
type OrganizationOperationsServerTransport struct {
	srv          *OrganizationOperationsServer
	newListPager *tracker[azfake.PagerResponder[armconfluent.OrganizationOperationsClientListResponse]]
}

// Do implements the policy.Transporter interface for OrganizationOperationsServerTransport.
func (o *OrganizationOperationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "OrganizationOperationsClient.NewListPager":
		resp, err = o.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (o *OrganizationOperationsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if o.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := o.newListPager.get(req)
	if newListPager == nil {
		resp := o.srv.NewListPager(nil)
		newListPager = &resp
		o.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armconfluent.OrganizationOperationsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		o.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		o.newListPager.remove(req)
	}
	return resp, nil
}
