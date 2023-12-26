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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcontainerservice/armhybridcontainerservice"
	"net/http"
	"net/url"
	"regexp"
)

// HybridIdentityMetadataServer is a fake server for instances of the armhybridcontainerservice.HybridIdentityMetadataClient type.
type HybridIdentityMetadataServer struct {
	// BeginDelete is the fake for method HybridIdentityMetadataClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, connectedClusterResourceURI string, options *armhybridcontainerservice.HybridIdentityMetadataClientBeginDeleteOptions) (resp azfake.PollerResponder[armhybridcontainerservice.HybridIdentityMetadataClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method HybridIdentityMetadataClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, connectedClusterResourceURI string, options *armhybridcontainerservice.HybridIdentityMetadataClientGetOptions) (resp azfake.Responder[armhybridcontainerservice.HybridIdentityMetadataClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByClusterPager is the fake for method HybridIdentityMetadataClient.NewListByClusterPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByClusterPager func(connectedClusterResourceURI string, options *armhybridcontainerservice.HybridIdentityMetadataClientListByClusterOptions) (resp azfake.PagerResponder[armhybridcontainerservice.HybridIdentityMetadataClientListByClusterResponse])

	// Put is the fake for method HybridIdentityMetadataClient.Put
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	Put func(ctx context.Context, connectedClusterResourceURI string, body armhybridcontainerservice.HybridIdentityMetadata, options *armhybridcontainerservice.HybridIdentityMetadataClientPutOptions) (resp azfake.Responder[armhybridcontainerservice.HybridIdentityMetadataClientPutResponse], errResp azfake.ErrorResponder)
}

// NewHybridIdentityMetadataServerTransport creates a new instance of HybridIdentityMetadataServerTransport with the provided implementation.
// The returned HybridIdentityMetadataServerTransport instance is connected to an instance of armhybridcontainerservice.HybridIdentityMetadataClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewHybridIdentityMetadataServerTransport(srv *HybridIdentityMetadataServer) *HybridIdentityMetadataServerTransport {
	return &HybridIdentityMetadataServerTransport{
		srv:                   srv,
		beginDelete:           newTracker[azfake.PollerResponder[armhybridcontainerservice.HybridIdentityMetadataClientDeleteResponse]](),
		newListByClusterPager: newTracker[azfake.PagerResponder[armhybridcontainerservice.HybridIdentityMetadataClientListByClusterResponse]](),
	}
}

// HybridIdentityMetadataServerTransport connects instances of armhybridcontainerservice.HybridIdentityMetadataClient to instances of HybridIdentityMetadataServer.
// Don't use this type directly, use NewHybridIdentityMetadataServerTransport instead.
type HybridIdentityMetadataServerTransport struct {
	srv                   *HybridIdentityMetadataServer
	beginDelete           *tracker[azfake.PollerResponder[armhybridcontainerservice.HybridIdentityMetadataClientDeleteResponse]]
	newListByClusterPager *tracker[azfake.PagerResponder[armhybridcontainerservice.HybridIdentityMetadataClientListByClusterResponse]]
}

// Do implements the policy.Transporter interface for HybridIdentityMetadataServerTransport.
func (h *HybridIdentityMetadataServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "HybridIdentityMetadataClient.BeginDelete":
		resp, err = h.dispatchBeginDelete(req)
	case "HybridIdentityMetadataClient.Get":
		resp, err = h.dispatchGet(req)
	case "HybridIdentityMetadataClient.NewListByClusterPager":
		resp, err = h.dispatchNewListByClusterPager(req)
	case "HybridIdentityMetadataClient.Put":
		resp, err = h.dispatchPut(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *HybridIdentityMetadataServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if h.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := h.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/(?P<connectedClusterResourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridContainerService/provisionedClusterInstances/default/hybridIdentityMetadata/default`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		connectedClusterResourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("connectedClusterResourceUri")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := h.srv.BeginDelete(req.Context(), connectedClusterResourceURIParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		h.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		h.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		h.beginDelete.remove(req)
	}

	return resp, nil
}

func (h *HybridIdentityMetadataServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if h.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/(?P<connectedClusterResourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridContainerService/provisionedClusterInstances/default/hybridIdentityMetadata/default`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	connectedClusterResourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("connectedClusterResourceUri")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := h.srv.Get(req.Context(), connectedClusterResourceURIParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).HybridIdentityMetadata, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HybridIdentityMetadataServerTransport) dispatchNewListByClusterPager(req *http.Request) (*http.Response, error) {
	if h.srv.NewListByClusterPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByClusterPager not implemented")}
	}
	newListByClusterPager := h.newListByClusterPager.get(req)
	if newListByClusterPager == nil {
		const regexStr = `/(?P<connectedClusterResourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridContainerService/provisionedClusterInstances/default/hybridIdentityMetadata`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		connectedClusterResourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("connectedClusterResourceUri")])
		if err != nil {
			return nil, err
		}
		resp := h.srv.NewListByClusterPager(connectedClusterResourceURIParam, nil)
		newListByClusterPager = &resp
		h.newListByClusterPager.add(req, newListByClusterPager)
		server.PagerResponderInjectNextLinks(newListByClusterPager, req, func(page *armhybridcontainerservice.HybridIdentityMetadataClientListByClusterResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByClusterPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		h.newListByClusterPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByClusterPager) {
		h.newListByClusterPager.remove(req)
	}
	return resp, nil
}

func (h *HybridIdentityMetadataServerTransport) dispatchPut(req *http.Request) (*http.Response, error) {
	if h.srv.Put == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put not implemented")}
	}
	const regexStr = `/(?P<connectedClusterResourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridContainerService/provisionedClusterInstances/default/hybridIdentityMetadata/default`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armhybridcontainerservice.HybridIdentityMetadata](req)
	if err != nil {
		return nil, err
	}
	connectedClusterResourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("connectedClusterResourceUri")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := h.srv.Put(req.Context(), connectedClusterResourceURIParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).HybridIdentityMetadata, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
