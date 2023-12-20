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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v3"
	"net/http"
	"net/url"
	"regexp"
)

// FetchCrossRegionRestoreJobServer is a fake server for instances of the armdataprotection.FetchCrossRegionRestoreJobClient type.
type FetchCrossRegionRestoreJobServer struct {
	// Get is the fake for method FetchCrossRegionRestoreJobClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, location string, parameters armdataprotection.CrossRegionRestoreJobRequest, options *armdataprotection.FetchCrossRegionRestoreJobClientGetOptions) (resp azfake.Responder[armdataprotection.FetchCrossRegionRestoreJobClientGetResponse], errResp azfake.ErrorResponder)
}

// NewFetchCrossRegionRestoreJobServerTransport creates a new instance of FetchCrossRegionRestoreJobServerTransport with the provided implementation.
// The returned FetchCrossRegionRestoreJobServerTransport instance is connected to an instance of armdataprotection.FetchCrossRegionRestoreJobClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewFetchCrossRegionRestoreJobServerTransport(srv *FetchCrossRegionRestoreJobServer) *FetchCrossRegionRestoreJobServerTransport {
	return &FetchCrossRegionRestoreJobServerTransport{srv: srv}
}

// FetchCrossRegionRestoreJobServerTransport connects instances of armdataprotection.FetchCrossRegionRestoreJobClient to instances of FetchCrossRegionRestoreJobServer.
// Don't use this type directly, use NewFetchCrossRegionRestoreJobServerTransport instead.
type FetchCrossRegionRestoreJobServerTransport struct {
	srv *FetchCrossRegionRestoreJobServer
}

// Do implements the policy.Transporter interface for FetchCrossRegionRestoreJobServerTransport.
func (f *FetchCrossRegionRestoreJobServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "FetchCrossRegionRestoreJobClient.Get":
		resp, err = f.dispatchGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (f *FetchCrossRegionRestoreJobServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if f.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataProtection/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/fetchCrossRegionRestoreJob`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdataprotection.CrossRegionRestoreJobRequest](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.Get(req.Context(), resourceGroupNameParam, locationParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AzureBackupJobResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
