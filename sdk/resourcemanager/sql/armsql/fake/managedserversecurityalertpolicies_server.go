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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
	"net/http"
	"net/url"
	"regexp"
)

// ManagedServerSecurityAlertPoliciesServer is a fake server for instances of the armsql.ManagedServerSecurityAlertPoliciesClient type.
type ManagedServerSecurityAlertPoliciesServer struct {
	// BeginCreateOrUpdate is the fake for method ManagedServerSecurityAlertPoliciesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, managedInstanceName string, securityAlertPolicyName armsql.SecurityAlertPolicyName, parameters armsql.ManagedServerSecurityAlertPolicy, options *armsql.ManagedServerSecurityAlertPoliciesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armsql.ManagedServerSecurityAlertPoliciesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ManagedServerSecurityAlertPoliciesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, managedInstanceName string, securityAlertPolicyName armsql.SecurityAlertPolicyName, options *armsql.ManagedServerSecurityAlertPoliciesClientGetOptions) (resp azfake.Responder[armsql.ManagedServerSecurityAlertPoliciesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByInstancePager is the fake for method ManagedServerSecurityAlertPoliciesClient.NewListByInstancePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByInstancePager func(resourceGroupName string, managedInstanceName string, options *armsql.ManagedServerSecurityAlertPoliciesClientListByInstanceOptions) (resp azfake.PagerResponder[armsql.ManagedServerSecurityAlertPoliciesClientListByInstanceResponse])
}

// NewManagedServerSecurityAlertPoliciesServerTransport creates a new instance of ManagedServerSecurityAlertPoliciesServerTransport with the provided implementation.
// The returned ManagedServerSecurityAlertPoliciesServerTransport instance is connected to an instance of armsql.ManagedServerSecurityAlertPoliciesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewManagedServerSecurityAlertPoliciesServerTransport(srv *ManagedServerSecurityAlertPoliciesServer) *ManagedServerSecurityAlertPoliciesServerTransport {
	return &ManagedServerSecurityAlertPoliciesServerTransport{
		srv:                    srv,
		beginCreateOrUpdate:    newTracker[azfake.PollerResponder[armsql.ManagedServerSecurityAlertPoliciesClientCreateOrUpdateResponse]](),
		newListByInstancePager: newTracker[azfake.PagerResponder[armsql.ManagedServerSecurityAlertPoliciesClientListByInstanceResponse]](),
	}
}

// ManagedServerSecurityAlertPoliciesServerTransport connects instances of armsql.ManagedServerSecurityAlertPoliciesClient to instances of ManagedServerSecurityAlertPoliciesServer.
// Don't use this type directly, use NewManagedServerSecurityAlertPoliciesServerTransport instead.
type ManagedServerSecurityAlertPoliciesServerTransport struct {
	srv                    *ManagedServerSecurityAlertPoliciesServer
	beginCreateOrUpdate    *tracker[azfake.PollerResponder[armsql.ManagedServerSecurityAlertPoliciesClientCreateOrUpdateResponse]]
	newListByInstancePager *tracker[azfake.PagerResponder[armsql.ManagedServerSecurityAlertPoliciesClientListByInstanceResponse]]
}

// Do implements the policy.Transporter interface for ManagedServerSecurityAlertPoliciesServerTransport.
func (m *ManagedServerSecurityAlertPoliciesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ManagedServerSecurityAlertPoliciesClient.BeginCreateOrUpdate":
		resp, err = m.dispatchBeginCreateOrUpdate(req)
	case "ManagedServerSecurityAlertPoliciesClient.Get":
		resp, err = m.dispatchGet(req)
	case "ManagedServerSecurityAlertPoliciesClient.NewListByInstancePager":
		resp, err = m.dispatchNewListByInstancePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *ManagedServerSecurityAlertPoliciesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if m.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := m.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/securityAlertPolicies/(?P<securityAlertPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsql.ManagedServerSecurityAlertPolicy](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
		if err != nil {
			return nil, err
		}
		securityAlertPolicyNameParam, err := parseWithCast(matches[regex.SubexpIndex("securityAlertPolicyName")], func(v string) (armsql.SecurityAlertPolicyName, error) {
			p, unescapeErr := url.PathUnescape(v)
			if unescapeErr != nil {
				return "", unescapeErr
			}
			return armsql.SecurityAlertPolicyName(p), nil
		})
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, managedInstanceNameParam, securityAlertPolicyNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		m.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		m.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		m.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (m *ManagedServerSecurityAlertPoliciesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if m.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/securityAlertPolicies/(?P<securityAlertPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
	if err != nil {
		return nil, err
	}
	securityAlertPolicyNameParam, err := parseWithCast(matches[regex.SubexpIndex("securityAlertPolicyName")], func(v string) (armsql.SecurityAlertPolicyName, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armsql.SecurityAlertPolicyName(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Get(req.Context(), resourceGroupNameParam, managedInstanceNameParam, securityAlertPolicyNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ManagedServerSecurityAlertPolicy, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagedServerSecurityAlertPoliciesServerTransport) dispatchNewListByInstancePager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListByInstancePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByInstancePager not implemented")}
	}
	newListByInstancePager := m.newListByInstancePager.get(req)
	if newListByInstancePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/securityAlertPolicies`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
		if err != nil {
			return nil, err
		}
		resp := m.srv.NewListByInstancePager(resourceGroupNameParam, managedInstanceNameParam, nil)
		newListByInstancePager = &resp
		m.newListByInstancePager.add(req, newListByInstancePager)
		server.PagerResponderInjectNextLinks(newListByInstancePager, req, func(page *armsql.ManagedServerSecurityAlertPoliciesClientListByInstanceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByInstancePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListByInstancePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByInstancePager) {
		m.newListByInstancePager.remove(req)
	}
	return resp, nil
}
