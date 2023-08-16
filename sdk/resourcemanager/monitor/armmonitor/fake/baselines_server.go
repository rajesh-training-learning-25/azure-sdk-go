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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"net/http"
	"net/url"
	"regexp"
)

// BaselinesServer is a fake server for instances of the armmonitor.BaselinesClient type.
type BaselinesServer struct {
	// NewListPager is the fake for method BaselinesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceURI string, options *armmonitor.BaselinesClientListOptions) (resp azfake.PagerResponder[armmonitor.BaselinesClientListResponse])
}

// NewBaselinesServerTransport creates a new instance of BaselinesServerTransport with the provided implementation.
// The returned BaselinesServerTransport instance is connected to an instance of armmonitor.BaselinesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBaselinesServerTransport(srv *BaselinesServer) *BaselinesServerTransport {
	return &BaselinesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armmonitor.BaselinesClientListResponse]](),
	}
}

// BaselinesServerTransport connects instances of armmonitor.BaselinesClient to instances of BaselinesServer.
// Don't use this type directly, use NewBaselinesServerTransport instead.
type BaselinesServerTransport struct {
	srv          *BaselinesServer
	newListPager *tracker[azfake.PagerResponder[armmonitor.BaselinesClientListResponse]]
}

// Do implements the policy.Transporter interface for BaselinesServerTransport.
func (b *BaselinesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "BaselinesClient.NewListPager":
		resp, err = b.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b *BaselinesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if b.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := b.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Insights/metricBaselines`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceURIUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
		if err != nil {
			return nil, err
		}
		metricnamesUnescaped, err := url.QueryUnescape(qp.Get("metricnames"))
		if err != nil {
			return nil, err
		}
		metricnamesParam := getOptional(metricnamesUnescaped)
		metricnamespaceUnescaped, err := url.QueryUnescape(qp.Get("metricnamespace"))
		if err != nil {
			return nil, err
		}
		metricnamespaceParam := getOptional(metricnamespaceUnescaped)
		timespanUnescaped, err := url.QueryUnescape(qp.Get("timespan"))
		if err != nil {
			return nil, err
		}
		timespanParam := getOptional(timespanUnescaped)
		intervalUnescaped, err := url.QueryUnescape(qp.Get("interval"))
		if err != nil {
			return nil, err
		}
		intervalParam := getOptional(intervalUnescaped)
		aggregationUnescaped, err := url.QueryUnescape(qp.Get("aggregation"))
		if err != nil {
			return nil, err
		}
		aggregationParam := getOptional(aggregationUnescaped)
		sensitivitiesUnescaped, err := url.QueryUnescape(qp.Get("sensitivities"))
		if err != nil {
			return nil, err
		}
		sensitivitiesParam := getOptional(sensitivitiesUnescaped)
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		resultTypeUnescaped, err := url.QueryUnescape(qp.Get("resultType"))
		if err != nil {
			return nil, err
		}
		resultTypeParam := getOptional(armmonitor.ResultType(resultTypeUnescaped))
		var options *armmonitor.BaselinesClientListOptions
		if metricnamesParam != nil || metricnamespaceParam != nil || timespanParam != nil || intervalParam != nil || aggregationParam != nil || sensitivitiesParam != nil || filterParam != nil || resultTypeParam != nil {
			options = &armmonitor.BaselinesClientListOptions{
				Metricnames:     metricnamesParam,
				Metricnamespace: metricnamespaceParam,
				Timespan:        timespanParam,
				Interval:        intervalParam,
				Aggregation:     aggregationParam,
				Sensitivities:   sensitivitiesParam,
				Filter:          filterParam,
				ResultType:      resultTypeParam,
			}
		}
		resp := b.srv.NewListPager(resourceURIUnescaped, options)
		newListPager = &resp
		b.newListPager.add(req, newListPager)
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		b.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		b.newListPager.remove(req)
	}
	return resp, nil
}
