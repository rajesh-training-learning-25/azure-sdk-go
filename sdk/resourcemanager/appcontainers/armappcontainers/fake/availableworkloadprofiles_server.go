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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v2"
	"net/http"
	"net/url"
	"regexp"
)

// AvailableWorkloadProfilesServer is a fake server for instances of the armappcontainers.AvailableWorkloadProfilesClient type.
type AvailableWorkloadProfilesServer struct {
	// NewGetPager is the fake for method AvailableWorkloadProfilesClient.NewGetPager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetPager func(location string, options *armappcontainers.AvailableWorkloadProfilesClientGetOptions) (resp azfake.PagerResponder[armappcontainers.AvailableWorkloadProfilesClientGetResponse])
}

// NewAvailableWorkloadProfilesServerTransport creates a new instance of AvailableWorkloadProfilesServerTransport with the provided implementation.
// The returned AvailableWorkloadProfilesServerTransport instance is connected to an instance of armappcontainers.AvailableWorkloadProfilesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAvailableWorkloadProfilesServerTransport(srv *AvailableWorkloadProfilesServer) *AvailableWorkloadProfilesServerTransport {
	return &AvailableWorkloadProfilesServerTransport{
		srv:         srv,
		newGetPager: newTracker[azfake.PagerResponder[armappcontainers.AvailableWorkloadProfilesClientGetResponse]](),
	}
}

// AvailableWorkloadProfilesServerTransport connects instances of armappcontainers.AvailableWorkloadProfilesClient to instances of AvailableWorkloadProfilesServer.
// Don't use this type directly, use NewAvailableWorkloadProfilesServerTransport instead.
type AvailableWorkloadProfilesServerTransport struct {
	srv         *AvailableWorkloadProfilesServer
	newGetPager *tracker[azfake.PagerResponder[armappcontainers.AvailableWorkloadProfilesClientGetResponse]]
}

// Do implements the policy.Transporter interface for AvailableWorkloadProfilesServerTransport.
func (a *AvailableWorkloadProfilesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AvailableWorkloadProfilesClient.NewGetPager":
		resp, err = a.dispatchNewGetPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AvailableWorkloadProfilesServerTransport) dispatchNewGetPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewGetPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewGetPager not implemented")}
	}
	newGetPager := a.newGetPager.get(req)
	if newGetPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/availableManagedEnvironmentsWorkloadProfileTypes`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewGetPager(locationParam, nil)
		newGetPager = &resp
		a.newGetPager.add(req, newGetPager)
		server.PagerResponderInjectNextLinks(newGetPager, req, func(page *armappcontainers.AvailableWorkloadProfilesClientGetResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newGetPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newGetPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newGetPager) {
		a.newGetPager.remove(req)
	}
	return resp, nil
}
