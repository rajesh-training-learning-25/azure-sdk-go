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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto/v2"
	"net/http"
	"net/url"
	"regexp"
)

// ClusterPrincipalAssignmentsServer is a fake server for instances of the armkusto.ClusterPrincipalAssignmentsClient type.
type ClusterPrincipalAssignmentsServer struct {
	// CheckNameAvailability is the fake for method ClusterPrincipalAssignmentsClient.CheckNameAvailability
	// HTTP status codes to indicate success: http.StatusOK
	CheckNameAvailability func(ctx context.Context, resourceGroupName string, clusterName string, principalAssignmentName armkusto.ClusterPrincipalAssignmentCheckNameRequest, options *armkusto.ClusterPrincipalAssignmentsClientCheckNameAvailabilityOptions) (resp azfake.Responder[armkusto.ClusterPrincipalAssignmentsClientCheckNameAvailabilityResponse], errResp azfake.ErrorResponder)

	// BeginCreateOrUpdate is the fake for method ClusterPrincipalAssignmentsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, clusterName string, principalAssignmentName string, parameters armkusto.ClusterPrincipalAssignment, options *armkusto.ClusterPrincipalAssignmentsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armkusto.ClusterPrincipalAssignmentsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ClusterPrincipalAssignmentsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, clusterName string, principalAssignmentName string, options *armkusto.ClusterPrincipalAssignmentsClientBeginDeleteOptions) (resp azfake.PollerResponder[armkusto.ClusterPrincipalAssignmentsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ClusterPrincipalAssignmentsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, clusterName string, principalAssignmentName string, options *armkusto.ClusterPrincipalAssignmentsClientGetOptions) (resp azfake.Responder[armkusto.ClusterPrincipalAssignmentsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ClusterPrincipalAssignmentsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, clusterName string, options *armkusto.ClusterPrincipalAssignmentsClientListOptions) (resp azfake.PagerResponder[armkusto.ClusterPrincipalAssignmentsClientListResponse])
}

// NewClusterPrincipalAssignmentsServerTransport creates a new instance of ClusterPrincipalAssignmentsServerTransport with the provided implementation.
// The returned ClusterPrincipalAssignmentsServerTransport instance is connected to an instance of armkusto.ClusterPrincipalAssignmentsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewClusterPrincipalAssignmentsServerTransport(srv *ClusterPrincipalAssignmentsServer) *ClusterPrincipalAssignmentsServerTransport {
	return &ClusterPrincipalAssignmentsServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armkusto.ClusterPrincipalAssignmentsClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armkusto.ClusterPrincipalAssignmentsClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armkusto.ClusterPrincipalAssignmentsClientListResponse]](),
	}
}

// ClusterPrincipalAssignmentsServerTransport connects instances of armkusto.ClusterPrincipalAssignmentsClient to instances of ClusterPrincipalAssignmentsServer.
// Don't use this type directly, use NewClusterPrincipalAssignmentsServerTransport instead.
type ClusterPrincipalAssignmentsServerTransport struct {
	srv                 *ClusterPrincipalAssignmentsServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armkusto.ClusterPrincipalAssignmentsClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armkusto.ClusterPrincipalAssignmentsClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armkusto.ClusterPrincipalAssignmentsClientListResponse]]
}

// Do implements the policy.Transporter interface for ClusterPrincipalAssignmentsServerTransport.
func (c *ClusterPrincipalAssignmentsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ClusterPrincipalAssignmentsClient.CheckNameAvailability":
		resp, err = c.dispatchCheckNameAvailability(req)
	case "ClusterPrincipalAssignmentsClient.BeginCreateOrUpdate":
		resp, err = c.dispatchBeginCreateOrUpdate(req)
	case "ClusterPrincipalAssignmentsClient.BeginDelete":
		resp, err = c.dispatchBeginDelete(req)
	case "ClusterPrincipalAssignmentsClient.Get":
		resp, err = c.dispatchGet(req)
	case "ClusterPrincipalAssignmentsClient.NewListPager":
		resp, err = c.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *ClusterPrincipalAssignmentsServerTransport) dispatchCheckNameAvailability(req *http.Request) (*http.Response, error) {
	if c.srv.CheckNameAvailability == nil {
		return nil, &nonRetriableError{errors.New("fake for method CheckNameAvailability not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Kusto/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/checkPrincipalAssignmentNameAvailability`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armkusto.ClusterPrincipalAssignmentCheckNameRequest](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.CheckNameAvailability(req.Context(), resourceGroupNameParam, clusterNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CheckNameResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ClusterPrincipalAssignmentsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := c.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Kusto/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/principalAssignments/(?P<principalAssignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armkusto.ClusterPrincipalAssignment](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		principalAssignmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("principalAssignmentName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, clusterNameParam, principalAssignmentNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		c.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		c.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		c.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (c *ClusterPrincipalAssignmentsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := c.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Kusto/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/principalAssignments/(?P<principalAssignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		principalAssignmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("principalAssignmentName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginDelete(req.Context(), resourceGroupNameParam, clusterNameParam, principalAssignmentNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		c.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		c.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		c.beginDelete.remove(req)
	}

	return resp, nil
}

func (c *ClusterPrincipalAssignmentsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Kusto/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/principalAssignments/(?P<principalAssignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
	if err != nil {
		return nil, err
	}
	principalAssignmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("principalAssignmentName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, clusterNameParam, principalAssignmentNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ClusterPrincipalAssignment, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ClusterPrincipalAssignmentsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := c.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Kusto/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/principalAssignments`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		resp := c.srv.NewListPager(resourceGroupNameParam, clusterNameParam, nil)
		newListPager = &resp
		c.newListPager.add(req, newListPager)
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		c.newListPager.remove(req)
	}
	return resp, nil
}
