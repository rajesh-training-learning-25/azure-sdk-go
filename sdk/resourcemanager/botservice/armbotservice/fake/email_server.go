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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/botservice/armbotservice"
	"net/http"
	"net/url"
	"regexp"
)

// EmailServer is a fake server for instances of the armbotservice.EmailClient type.
type EmailServer struct {
	// CreateSignInURL is the fake for method EmailClient.CreateSignInURL
	// HTTP status codes to indicate success: http.StatusOK
	CreateSignInURL func(ctx context.Context, resourceGroupName string, resourceName string, options *armbotservice.EmailClientCreateSignInURLOptions) (resp azfake.Responder[armbotservice.EmailClientCreateSignInURLResponse], errResp azfake.ErrorResponder)
}

// NewEmailServerTransport creates a new instance of EmailServerTransport with the provided implementation.
// The returned EmailServerTransport instance is connected to an instance of armbotservice.EmailClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewEmailServerTransport(srv *EmailServer) *EmailServerTransport {
	return &EmailServerTransport{srv: srv}
}

// EmailServerTransport connects instances of armbotservice.EmailClient to instances of EmailServer.
// Don't use this type directly, use NewEmailServerTransport instead.
type EmailServerTransport struct {
	srv *EmailServer
}

// Do implements the policy.Transporter interface for EmailServerTransport.
func (e *EmailServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "EmailClient.CreateSignInURL":
		resp, err = e.dispatchCreateSignInURL(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *EmailServerTransport) dispatchCreateSignInURL(req *http.Request) (*http.Response, error) {
	if e.srv.CreateSignInURL == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateSignInURL not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.BotService/botServices/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/createEmailSignInUrl`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.CreateSignInURL(req.Context(), resourceGroupNameParam, resourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CreateEmailSignInURLResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
