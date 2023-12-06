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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork/v2"
	"net/http"
	"net/url"
	"regexp"
)

// ConfigurationGroupSchemasServer is a fake server for instances of the armhybridnetwork.ConfigurationGroupSchemasClient type.
type ConfigurationGroupSchemasServer struct {
	// BeginCreateOrUpdate is the fake for method ConfigurationGroupSchemasClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, publisherName string, configurationGroupSchemaName string, parameters armhybridnetwork.ConfigurationGroupSchema, options *armhybridnetwork.ConfigurationGroupSchemasClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armhybridnetwork.ConfigurationGroupSchemasClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ConfigurationGroupSchemasClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, publisherName string, configurationGroupSchemaName string, options *armhybridnetwork.ConfigurationGroupSchemasClientBeginDeleteOptions) (resp azfake.PollerResponder[armhybridnetwork.ConfigurationGroupSchemasClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ConfigurationGroupSchemasClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, publisherName string, configurationGroupSchemaName string, options *armhybridnetwork.ConfigurationGroupSchemasClientGetOptions) (resp azfake.Responder[armhybridnetwork.ConfigurationGroupSchemasClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByPublisherPager is the fake for method ConfigurationGroupSchemasClient.NewListByPublisherPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByPublisherPager func(resourceGroupName string, publisherName string, options *armhybridnetwork.ConfigurationGroupSchemasClientListByPublisherOptions) (resp azfake.PagerResponder[armhybridnetwork.ConfigurationGroupSchemasClientListByPublisherResponse])

	// Update is the fake for method ConfigurationGroupSchemasClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, publisherName string, configurationGroupSchemaName string, parameters armhybridnetwork.TagsObject, options *armhybridnetwork.ConfigurationGroupSchemasClientUpdateOptions) (resp azfake.Responder[armhybridnetwork.ConfigurationGroupSchemasClientUpdateResponse], errResp azfake.ErrorResponder)

	// BeginUpdateState is the fake for method ConfigurationGroupSchemasClient.BeginUpdateState
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdateState func(ctx context.Context, resourceGroupName string, publisherName string, configurationGroupSchemaName string, parameters armhybridnetwork.ConfigurationGroupSchemaVersionUpdateState, options *armhybridnetwork.ConfigurationGroupSchemasClientBeginUpdateStateOptions) (resp azfake.PollerResponder[armhybridnetwork.ConfigurationGroupSchemasClientUpdateStateResponse], errResp azfake.ErrorResponder)
}

// NewConfigurationGroupSchemasServerTransport creates a new instance of ConfigurationGroupSchemasServerTransport with the provided implementation.
// The returned ConfigurationGroupSchemasServerTransport instance is connected to an instance of armhybridnetwork.ConfigurationGroupSchemasClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewConfigurationGroupSchemasServerTransport(srv *ConfigurationGroupSchemasServer) *ConfigurationGroupSchemasServerTransport {
	return &ConfigurationGroupSchemasServerTransport{
		srv:                     srv,
		beginCreateOrUpdate:     newTracker[azfake.PollerResponder[armhybridnetwork.ConfigurationGroupSchemasClientCreateOrUpdateResponse]](),
		beginDelete:             newTracker[azfake.PollerResponder[armhybridnetwork.ConfigurationGroupSchemasClientDeleteResponse]](),
		newListByPublisherPager: newTracker[azfake.PagerResponder[armhybridnetwork.ConfigurationGroupSchemasClientListByPublisherResponse]](),
		beginUpdateState:        newTracker[azfake.PollerResponder[armhybridnetwork.ConfigurationGroupSchemasClientUpdateStateResponse]](),
	}
}

// ConfigurationGroupSchemasServerTransport connects instances of armhybridnetwork.ConfigurationGroupSchemasClient to instances of ConfigurationGroupSchemasServer.
// Don't use this type directly, use NewConfigurationGroupSchemasServerTransport instead.
type ConfigurationGroupSchemasServerTransport struct {
	srv                     *ConfigurationGroupSchemasServer
	beginCreateOrUpdate     *tracker[azfake.PollerResponder[armhybridnetwork.ConfigurationGroupSchemasClientCreateOrUpdateResponse]]
	beginDelete             *tracker[azfake.PollerResponder[armhybridnetwork.ConfigurationGroupSchemasClientDeleteResponse]]
	newListByPublisherPager *tracker[azfake.PagerResponder[armhybridnetwork.ConfigurationGroupSchemasClientListByPublisherResponse]]
	beginUpdateState        *tracker[azfake.PollerResponder[armhybridnetwork.ConfigurationGroupSchemasClientUpdateStateResponse]]
}

// Do implements the policy.Transporter interface for ConfigurationGroupSchemasServerTransport.
func (c *ConfigurationGroupSchemasServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ConfigurationGroupSchemasClient.BeginCreateOrUpdate":
		resp, err = c.dispatchBeginCreateOrUpdate(req)
	case "ConfigurationGroupSchemasClient.BeginDelete":
		resp, err = c.dispatchBeginDelete(req)
	case "ConfigurationGroupSchemasClient.Get":
		resp, err = c.dispatchGet(req)
	case "ConfigurationGroupSchemasClient.NewListByPublisherPager":
		resp, err = c.dispatchNewListByPublisherPager(req)
	case "ConfigurationGroupSchemasClient.Update":
		resp, err = c.dispatchUpdate(req)
	case "ConfigurationGroupSchemasClient.BeginUpdateState":
		resp, err = c.dispatchBeginUpdateState(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *ConfigurationGroupSchemasServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := c.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridNetwork/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/configurationGroupSchemas/(?P<configurationGroupSchemaName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armhybridnetwork.ConfigurationGroupSchema](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		publisherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
		if err != nil {
			return nil, err
		}
		configurationGroupSchemaNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("configurationGroupSchemaName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, publisherNameParam, configurationGroupSchemaNameParam, body, nil)
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

func (c *ConfigurationGroupSchemasServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := c.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridNetwork/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/configurationGroupSchemas/(?P<configurationGroupSchemaName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		publisherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
		if err != nil {
			return nil, err
		}
		configurationGroupSchemaNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("configurationGroupSchemaName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginDelete(req.Context(), resourceGroupNameParam, publisherNameParam, configurationGroupSchemaNameParam, nil)
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

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		c.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		c.beginDelete.remove(req)
	}

	return resp, nil
}

func (c *ConfigurationGroupSchemasServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridNetwork/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/configurationGroupSchemas/(?P<configurationGroupSchemaName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	publisherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
	if err != nil {
		return nil, err
	}
	configurationGroupSchemaNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("configurationGroupSchemaName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, publisherNameParam, configurationGroupSchemaNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ConfigurationGroupSchema, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ConfigurationGroupSchemasServerTransport) dispatchNewListByPublisherPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListByPublisherPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByPublisherPager not implemented")}
	}
	newListByPublisherPager := c.newListByPublisherPager.get(req)
	if newListByPublisherPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridNetwork/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/configurationGroupSchemas`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		publisherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
		if err != nil {
			return nil, err
		}
		resp := c.srv.NewListByPublisherPager(resourceGroupNameParam, publisherNameParam, nil)
		newListByPublisherPager = &resp
		c.newListByPublisherPager.add(req, newListByPublisherPager)
		server.PagerResponderInjectNextLinks(newListByPublisherPager, req, func(page *armhybridnetwork.ConfigurationGroupSchemasClientListByPublisherResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByPublisherPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListByPublisherPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByPublisherPager) {
		c.newListByPublisherPager.remove(req)
	}
	return resp, nil
}

func (c *ConfigurationGroupSchemasServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridNetwork/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/configurationGroupSchemas/(?P<configurationGroupSchemaName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armhybridnetwork.TagsObject](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	publisherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
	if err != nil {
		return nil, err
	}
	configurationGroupSchemaNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("configurationGroupSchemaName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Update(req.Context(), resourceGroupNameParam, publisherNameParam, configurationGroupSchemaNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ConfigurationGroupSchema, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ConfigurationGroupSchemasServerTransport) dispatchBeginUpdateState(req *http.Request) (*http.Response, error) {
	if c.srv.BeginUpdateState == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdateState not implemented")}
	}
	beginUpdateState := c.beginUpdateState.get(req)
	if beginUpdateState == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridNetwork/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/configurationGroupSchemas/(?P<configurationGroupSchemaName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/updateState`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armhybridnetwork.ConfigurationGroupSchemaVersionUpdateState](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		publisherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
		if err != nil {
			return nil, err
		}
		configurationGroupSchemaNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("configurationGroupSchemaName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginUpdateState(req.Context(), resourceGroupNameParam, publisherNameParam, configurationGroupSchemaNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdateState = &respr
		c.beginUpdateState.add(req, beginUpdateState)
	}

	resp, err := server.PollerResponderNext(beginUpdateState, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginUpdateState.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdateState) {
		c.beginUpdateState.remove(req)
	}

	return resp, nil
}