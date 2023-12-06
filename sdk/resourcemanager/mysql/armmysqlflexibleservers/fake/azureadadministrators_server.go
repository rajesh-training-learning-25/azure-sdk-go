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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers/v2"
	"net/http"
	"net/url"
	"regexp"
)

// AzureADAdministratorsServer is a fake server for instances of the armmysqlflexibleservers.AzureADAdministratorsClient type.
type AzureADAdministratorsServer struct {
	// BeginCreateOrUpdate is the fake for method AzureADAdministratorsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, serverName string, administratorName armmysqlflexibleservers.AdministratorName, parameters armmysqlflexibleservers.AzureADAdministrator, options *armmysqlflexibleservers.AzureADAdministratorsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armmysqlflexibleservers.AzureADAdministratorsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method AzureADAdministratorsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, serverName string, administratorName armmysqlflexibleservers.AdministratorName, options *armmysqlflexibleservers.AzureADAdministratorsClientBeginDeleteOptions) (resp azfake.PollerResponder[armmysqlflexibleservers.AzureADAdministratorsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method AzureADAdministratorsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serverName string, administratorName armmysqlflexibleservers.AdministratorName, options *armmysqlflexibleservers.AzureADAdministratorsClientGetOptions) (resp azfake.Responder[armmysqlflexibleservers.AzureADAdministratorsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByServerPager is the fake for method AzureADAdministratorsClient.NewListByServerPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByServerPager func(resourceGroupName string, serverName string, options *armmysqlflexibleservers.AzureADAdministratorsClientListByServerOptions) (resp azfake.PagerResponder[armmysqlflexibleservers.AzureADAdministratorsClientListByServerResponse])
}

// NewAzureADAdministratorsServerTransport creates a new instance of AzureADAdministratorsServerTransport with the provided implementation.
// The returned AzureADAdministratorsServerTransport instance is connected to an instance of armmysqlflexibleservers.AzureADAdministratorsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAzureADAdministratorsServerTransport(srv *AzureADAdministratorsServer) *AzureADAdministratorsServerTransport {
	return &AzureADAdministratorsServerTransport{
		srv:                  srv,
		beginCreateOrUpdate:  newTracker[azfake.PollerResponder[armmysqlflexibleservers.AzureADAdministratorsClientCreateOrUpdateResponse]](),
		beginDelete:          newTracker[azfake.PollerResponder[armmysqlflexibleservers.AzureADAdministratorsClientDeleteResponse]](),
		newListByServerPager: newTracker[azfake.PagerResponder[armmysqlflexibleservers.AzureADAdministratorsClientListByServerResponse]](),
	}
}

// AzureADAdministratorsServerTransport connects instances of armmysqlflexibleservers.AzureADAdministratorsClient to instances of AzureADAdministratorsServer.
// Don't use this type directly, use NewAzureADAdministratorsServerTransport instead.
type AzureADAdministratorsServerTransport struct {
	srv                  *AzureADAdministratorsServer
	beginCreateOrUpdate  *tracker[azfake.PollerResponder[armmysqlflexibleservers.AzureADAdministratorsClientCreateOrUpdateResponse]]
	beginDelete          *tracker[azfake.PollerResponder[armmysqlflexibleservers.AzureADAdministratorsClientDeleteResponse]]
	newListByServerPager *tracker[azfake.PagerResponder[armmysqlflexibleservers.AzureADAdministratorsClientListByServerResponse]]
}

// Do implements the policy.Transporter interface for AzureADAdministratorsServerTransport.
func (a *AzureADAdministratorsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AzureADAdministratorsClient.BeginCreateOrUpdate":
		resp, err = a.dispatchBeginCreateOrUpdate(req)
	case "AzureADAdministratorsClient.BeginDelete":
		resp, err = a.dispatchBeginDelete(req)
	case "AzureADAdministratorsClient.Get":
		resp, err = a.dispatchGet(req)
	case "AzureADAdministratorsClient.NewListByServerPager":
		resp, err = a.dispatchNewListByServerPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AzureADAdministratorsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := a.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforMySQL/flexibleServers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/administrators/(?P<administratorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmysqlflexibleservers.AzureADAdministrator](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		administratorNameParam, err := parseWithCast(matches[regex.SubexpIndex("administratorName")], func(v string) (armmysqlflexibleservers.AdministratorName, error) {
			p, unescapeErr := url.PathUnescape(v)
			if unescapeErr != nil {
				return "", unescapeErr
			}
			return armmysqlflexibleservers.AdministratorName(p), nil
		})
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, serverNameParam, administratorNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		a.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		a.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		a.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (a *AzureADAdministratorsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if a.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := a.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforMySQL/flexibleServers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/administrators/(?P<administratorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		administratorNameParam, err := parseWithCast(matches[regex.SubexpIndex("administratorName")], func(v string) (armmysqlflexibleservers.AdministratorName, error) {
			p, unescapeErr := url.PathUnescape(v)
			if unescapeErr != nil {
				return "", unescapeErr
			}
			return armmysqlflexibleservers.AdministratorName(p), nil
		})
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginDelete(req.Context(), resourceGroupNameParam, serverNameParam, administratorNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		a.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		a.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		a.beginDelete.remove(req)
	}

	return resp, nil
}

func (a *AzureADAdministratorsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforMySQL/flexibleServers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/administrators/(?P<administratorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
	if err != nil {
		return nil, err
	}
	administratorNameParam, err := parseWithCast(matches[regex.SubexpIndex("administratorName")], func(v string) (armmysqlflexibleservers.AdministratorName, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armmysqlflexibleservers.AdministratorName(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, serverNameParam, administratorNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AzureADAdministrator, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AzureADAdministratorsServerTransport) dispatchNewListByServerPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListByServerPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByServerPager not implemented")}
	}
	newListByServerPager := a.newListByServerPager.get(req)
	if newListByServerPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforMySQL/flexibleServers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/administrators`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListByServerPager(resourceGroupNameParam, serverNameParam, nil)
		newListByServerPager = &resp
		a.newListByServerPager.add(req, newListByServerPager)
		server.PagerResponderInjectNextLinks(newListByServerPager, req, func(page *armmysqlflexibleservers.AzureADAdministratorsClientListByServerResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByServerPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListByServerPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByServerPager) {
		a.newListByServerPager.remove(req)
	}
	return resp, nil
}