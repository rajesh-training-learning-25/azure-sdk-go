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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare"
	"net/http"
	"net/url"
	"regexp"
)

// EmailRegistrationsServer is a fake server for instances of the armdatashare.EmailRegistrationsClient type.
type EmailRegistrationsServer struct {
	// ActivateEmail is the fake for method EmailRegistrationsClient.ActivateEmail
	// HTTP status codes to indicate success: http.StatusOK
	ActivateEmail func(ctx context.Context, location string, emailRegistration armdatashare.EmailRegistration, options *armdatashare.EmailRegistrationsClientActivateEmailOptions) (resp azfake.Responder[armdatashare.EmailRegistrationsClientActivateEmailResponse], errResp azfake.ErrorResponder)

	// RegisterEmail is the fake for method EmailRegistrationsClient.RegisterEmail
	// HTTP status codes to indicate success: http.StatusOK
	RegisterEmail func(ctx context.Context, location string, options *armdatashare.EmailRegistrationsClientRegisterEmailOptions) (resp azfake.Responder[armdatashare.EmailRegistrationsClientRegisterEmailResponse], errResp azfake.ErrorResponder)
}

// NewEmailRegistrationsServerTransport creates a new instance of EmailRegistrationsServerTransport with the provided implementation.
// The returned EmailRegistrationsServerTransport instance is connected to an instance of armdatashare.EmailRegistrationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewEmailRegistrationsServerTransport(srv *EmailRegistrationsServer) *EmailRegistrationsServerTransport {
	return &EmailRegistrationsServerTransport{srv: srv}
}

// EmailRegistrationsServerTransport connects instances of armdatashare.EmailRegistrationsClient to instances of EmailRegistrationsServer.
// Don't use this type directly, use NewEmailRegistrationsServerTransport instead.
type EmailRegistrationsServerTransport struct {
	srv *EmailRegistrationsServer
}

// Do implements the policy.Transporter interface for EmailRegistrationsServerTransport.
func (e *EmailRegistrationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "EmailRegistrationsClient.ActivateEmail":
		resp, err = e.dispatchActivateEmail(req)
	case "EmailRegistrationsClient.RegisterEmail":
		resp, err = e.dispatchRegisterEmail(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *EmailRegistrationsServerTransport) dispatchActivateEmail(req *http.Request) (*http.Response, error) {
	if e.srv.ActivateEmail == nil {
		return nil, &nonRetriableError{errors.New("fake for method ActivateEmail not implemented")}
	}
	const regexStr = `/providers/Microsoft\.DataShare/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/activateEmail`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdatashare.EmailRegistration](req)
	if err != nil {
		return nil, err
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.ActivateEmail(req.Context(), locationParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).EmailRegistration, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EmailRegistrationsServerTransport) dispatchRegisterEmail(req *http.Request) (*http.Response, error) {
	if e.srv.RegisterEmail == nil {
		return nil, &nonRetriableError{errors.New("fake for method RegisterEmail not implemented")}
	}
	const regexStr = `/providers/Microsoft\.DataShare/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/registerEmail`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.RegisterEmail(req.Context(), locationParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).EmailRegistration, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}