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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
	"net/http"
	"net/url"
	"regexp"
)

// TenantLevelAccessReviewInstanceContactedReviewersServer is a fake server for instances of the armauthorization.TenantLevelAccessReviewInstanceContactedReviewersClient type.
type TenantLevelAccessReviewInstanceContactedReviewersServer struct {
	// NewListPager is the fake for method TenantLevelAccessReviewInstanceContactedReviewersClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(scheduleDefinitionID string, id string, options *armauthorization.TenantLevelAccessReviewInstanceContactedReviewersClientListOptions) (resp azfake.PagerResponder[armauthorization.TenantLevelAccessReviewInstanceContactedReviewersClientListResponse])
}

// NewTenantLevelAccessReviewInstanceContactedReviewersServerTransport creates a new instance of TenantLevelAccessReviewInstanceContactedReviewersServerTransport with the provided implementation.
// The returned TenantLevelAccessReviewInstanceContactedReviewersServerTransport instance is connected to an instance of armauthorization.TenantLevelAccessReviewInstanceContactedReviewersClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewTenantLevelAccessReviewInstanceContactedReviewersServerTransport(srv *TenantLevelAccessReviewInstanceContactedReviewersServer) *TenantLevelAccessReviewInstanceContactedReviewersServerTransport {
	return &TenantLevelAccessReviewInstanceContactedReviewersServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armauthorization.TenantLevelAccessReviewInstanceContactedReviewersClientListResponse]](),
	}
}

// TenantLevelAccessReviewInstanceContactedReviewersServerTransport connects instances of armauthorization.TenantLevelAccessReviewInstanceContactedReviewersClient to instances of TenantLevelAccessReviewInstanceContactedReviewersServer.
// Don't use this type directly, use NewTenantLevelAccessReviewInstanceContactedReviewersServerTransport instead.
type TenantLevelAccessReviewInstanceContactedReviewersServerTransport struct {
	srv          *TenantLevelAccessReviewInstanceContactedReviewersServer
	newListPager *tracker[azfake.PagerResponder[armauthorization.TenantLevelAccessReviewInstanceContactedReviewersClientListResponse]]
}

// Do implements the policy.Transporter interface for TenantLevelAccessReviewInstanceContactedReviewersServerTransport.
func (t *TenantLevelAccessReviewInstanceContactedReviewersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "TenantLevelAccessReviewInstanceContactedReviewersClient.NewListPager":
		resp, err = t.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (t *TenantLevelAccessReviewInstanceContactedReviewersServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if t.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := t.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/providers/Microsoft\.Authorization/accessReviewScheduleDefinitions/(?P<scheduleDefinitionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/instances/(?P<id>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/contactedReviewers`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		scheduleDefinitionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("scheduleDefinitionId")])
		if err != nil {
			return nil, err
		}
		idParam, err := url.PathUnescape(matches[regex.SubexpIndex("id")])
		if err != nil {
			return nil, err
		}
		resp := t.srv.NewListPager(scheduleDefinitionIDParam, idParam, nil)
		newListPager = &resp
		t.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armauthorization.TenantLevelAccessReviewInstanceContactedReviewersClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		t.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		t.newListPager.remove(req)
	}
	return resp, nil
}
