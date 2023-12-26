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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbiprivatelinks/armpowerbiprivatelinks/v2"
	"net/http"
	"regexp"
)

// PrivateLinkServicesForPowerBIServer is a fake server for instances of the armpowerbiprivatelinks.PrivateLinkServicesForPowerBIClient type.
type PrivateLinkServicesForPowerBIServer struct {
	// ListBySubscriptionID is the fake for method PrivateLinkServicesForPowerBIClient.ListBySubscriptionID
	// HTTP status codes to indicate success: http.StatusOK
	ListBySubscriptionID func(ctx context.Context, options *armpowerbiprivatelinks.PrivateLinkServicesForPowerBIClientListBySubscriptionIDOptions) (resp azfake.Responder[armpowerbiprivatelinks.PrivateLinkServicesForPowerBIClientListBySubscriptionIDResponse], errResp azfake.ErrorResponder)
}

// NewPrivateLinkServicesForPowerBIServerTransport creates a new instance of PrivateLinkServicesForPowerBIServerTransport with the provided implementation.
// The returned PrivateLinkServicesForPowerBIServerTransport instance is connected to an instance of armpowerbiprivatelinks.PrivateLinkServicesForPowerBIClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPrivateLinkServicesForPowerBIServerTransport(srv *PrivateLinkServicesForPowerBIServer) *PrivateLinkServicesForPowerBIServerTransport {
	return &PrivateLinkServicesForPowerBIServerTransport{srv: srv}
}

// PrivateLinkServicesForPowerBIServerTransport connects instances of armpowerbiprivatelinks.PrivateLinkServicesForPowerBIClient to instances of PrivateLinkServicesForPowerBIServer.
// Don't use this type directly, use NewPrivateLinkServicesForPowerBIServerTransport instead.
type PrivateLinkServicesForPowerBIServerTransport struct {
	srv *PrivateLinkServicesForPowerBIServer
}

// Do implements the policy.Transporter interface for PrivateLinkServicesForPowerBIServerTransport.
func (p *PrivateLinkServicesForPowerBIServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PrivateLinkServicesForPowerBIClient.ListBySubscriptionID":
		resp, err = p.dispatchListBySubscriptionID(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PrivateLinkServicesForPowerBIServerTransport) dispatchListBySubscriptionID(req *http.Request) (*http.Response, error) {
	if p.srv.ListBySubscriptionID == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListBySubscriptionID not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.PowerBI/privateLinkServicesForPowerBI`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := p.srv.ListBySubscriptionID(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).TenantResourceArray, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
