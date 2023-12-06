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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/orbital/armorbital/v2"
	"net/http"
	"net/url"
	"regexp"
)

// AvailableGroundStationsServer is a fake server for instances of the armorbital.AvailableGroundStationsClient type.
type AvailableGroundStationsServer struct {
	// NewListByCapabilityPager is the fake for method AvailableGroundStationsClient.NewListByCapabilityPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByCapabilityPager func(capability armorbital.CapabilityParameter, options *armorbital.AvailableGroundStationsClientListByCapabilityOptions) (resp azfake.PagerResponder[armorbital.AvailableGroundStationsClientListByCapabilityResponse])
}

// NewAvailableGroundStationsServerTransport creates a new instance of AvailableGroundStationsServerTransport with the provided implementation.
// The returned AvailableGroundStationsServerTransport instance is connected to an instance of armorbital.AvailableGroundStationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAvailableGroundStationsServerTransport(srv *AvailableGroundStationsServer) *AvailableGroundStationsServerTransport {
	return &AvailableGroundStationsServerTransport{
		srv:                      srv,
		newListByCapabilityPager: newTracker[azfake.PagerResponder[armorbital.AvailableGroundStationsClientListByCapabilityResponse]](),
	}
}

// AvailableGroundStationsServerTransport connects instances of armorbital.AvailableGroundStationsClient to instances of AvailableGroundStationsServer.
// Don't use this type directly, use NewAvailableGroundStationsServerTransport instead.
type AvailableGroundStationsServerTransport struct {
	srv                      *AvailableGroundStationsServer
	newListByCapabilityPager *tracker[azfake.PagerResponder[armorbital.AvailableGroundStationsClientListByCapabilityResponse]]
}

// Do implements the policy.Transporter interface for AvailableGroundStationsServerTransport.
func (a *AvailableGroundStationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AvailableGroundStationsClient.NewListByCapabilityPager":
		resp, err = a.dispatchNewListByCapabilityPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AvailableGroundStationsServerTransport) dispatchNewListByCapabilityPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListByCapabilityPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByCapabilityPager not implemented")}
	}
	newListByCapabilityPager := a.newListByCapabilityPager.get(req)
	if newListByCapabilityPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Orbital/availableGroundStations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		capabilityParam, err := parseWithCast(qp.Get("capability"), func(v string) (armorbital.CapabilityParameter, error) {
			p, unescapeErr := url.QueryUnescape(v)
			if unescapeErr != nil {
				return "", unescapeErr
			}
			return armorbital.CapabilityParameter(p), nil
		})
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListByCapabilityPager(capabilityParam, nil)
		newListByCapabilityPager = &resp
		a.newListByCapabilityPager.add(req, newListByCapabilityPager)
		server.PagerResponderInjectNextLinks(newListByCapabilityPager, req, func(page *armorbital.AvailableGroundStationsClientListByCapabilityResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByCapabilityPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListByCapabilityPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByCapabilityPager) {
		a.newListByCapabilityPager.remove(req)
	}
	return resp, nil
}