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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
	"net/http"
	"net/url"
	"regexp"
)

// ReservationRecommendationDetailsServer is a fake server for instances of the armconsumption.ReservationRecommendationDetailsClient type.
type ReservationRecommendationDetailsServer struct {
	// Get is the fake for method ReservationRecommendationDetailsClient.Get
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Get func(ctx context.Context, resourceScope string, scope armconsumption.Scope, region string, term armconsumption.Term, lookBackPeriod armconsumption.LookBackPeriod, product string, options *armconsumption.ReservationRecommendationDetailsClientGetOptions) (resp azfake.Responder[armconsumption.ReservationRecommendationDetailsClientGetResponse], errResp azfake.ErrorResponder)
}

// NewReservationRecommendationDetailsServerTransport creates a new instance of ReservationRecommendationDetailsServerTransport with the provided implementation.
// The returned ReservationRecommendationDetailsServerTransport instance is connected to an instance of armconsumption.ReservationRecommendationDetailsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewReservationRecommendationDetailsServerTransport(srv *ReservationRecommendationDetailsServer) *ReservationRecommendationDetailsServerTransport {
	return &ReservationRecommendationDetailsServerTransport{srv: srv}
}

// ReservationRecommendationDetailsServerTransport connects instances of armconsumption.ReservationRecommendationDetailsClient to instances of ReservationRecommendationDetailsServer.
// Don't use this type directly, use NewReservationRecommendationDetailsServerTransport instead.
type ReservationRecommendationDetailsServerTransport struct {
	srv *ReservationRecommendationDetailsServer
}

// Do implements the policy.Transporter interface for ReservationRecommendationDetailsServerTransport.
func (r *ReservationRecommendationDetailsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ReservationRecommendationDetailsClient.Get":
		resp, err = r.dispatchGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *ReservationRecommendationDetailsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/(?P<resourceScope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Consumption/reservationRecommendationDetails`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceScopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceScope")])
	if err != nil {
		return nil, err
	}
	scopeParam, err := parseWithCast(qp.Get("scope"), func(v string) (armconsumption.Scope, error) {
		p, unescapeErr := url.QueryUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armconsumption.Scope(p), nil
	})
	if err != nil {
		return nil, err
	}
	regionParam, err := url.QueryUnescape(qp.Get("region"))
	if err != nil {
		return nil, err
	}
	termParam, err := parseWithCast(qp.Get("term"), func(v string) (armconsumption.Term, error) {
		p, unescapeErr := url.QueryUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armconsumption.Term(p), nil
	})
	if err != nil {
		return nil, err
	}
	lookBackPeriodParam, err := parseWithCast(qp.Get("lookBackPeriod"), func(v string) (armconsumption.LookBackPeriod, error) {
		p, unescapeErr := url.QueryUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armconsumption.LookBackPeriod(p), nil
	})
	if err != nil {
		return nil, err
	}
	productParam, err := url.QueryUnescape(qp.Get("product"))
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), resourceScopeParam, scopeParam, regionParam, termParam, lookBackPeriodParam, productParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ReservationRecommendationDetailsModel, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
