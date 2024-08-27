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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters"
	"net/http"
	"net/url"
	"regexp"
)

// ManagedAzResiliencyStatusServer is a fake server for instances of the armservicefabricmanagedclusters.ManagedAzResiliencyStatusClient type.
type ManagedAzResiliencyStatusServer struct {
	// Get is the fake for method ManagedAzResiliencyStatusClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, clusterName string, options *armservicefabricmanagedclusters.ManagedAzResiliencyStatusClientGetOptions) (resp azfake.Responder[armservicefabricmanagedclusters.ManagedAzResiliencyStatusClientGetResponse], errResp azfake.ErrorResponder)
}

// NewManagedAzResiliencyStatusServerTransport creates a new instance of ManagedAzResiliencyStatusServerTransport with the provided implementation.
// The returned ManagedAzResiliencyStatusServerTransport instance is connected to an instance of armservicefabricmanagedclusters.ManagedAzResiliencyStatusClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewManagedAzResiliencyStatusServerTransport(srv *ManagedAzResiliencyStatusServer) *ManagedAzResiliencyStatusServerTransport {
	return &ManagedAzResiliencyStatusServerTransport{srv: srv}
}

// ManagedAzResiliencyStatusServerTransport connects instances of armservicefabricmanagedclusters.ManagedAzResiliencyStatusClient to instances of ManagedAzResiliencyStatusServer.
// Don't use this type directly, use NewManagedAzResiliencyStatusServerTransport instead.
type ManagedAzResiliencyStatusServerTransport struct {
	srv *ManagedAzResiliencyStatusServer
}

// Do implements the policy.Transporter interface for ManagedAzResiliencyStatusServerTransport.
func (m *ManagedAzResiliencyStatusServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ManagedAzResiliencyStatusClient.Get":
		resp, err = m.dispatchGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *ManagedAzResiliencyStatusServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if m.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceFabric/managedClusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getazresiliencystatus`
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
	respr, errRespr := m.srv.Get(req.Context(), resourceGroupNameParam, clusterNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ManagedAzResiliencyStatus, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}