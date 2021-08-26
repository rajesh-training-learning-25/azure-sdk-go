// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"strings"
)

// BaselinesClient contains the methods for the Baselines group.
// Don't use this type directly, use NewBaselinesClient() instead.
type BaselinesClient struct {
	con *armcore.Connection
}

// NewBaselinesClient creates a new instance of BaselinesClient with the specified values.
func NewBaselinesClient(con *armcore.Connection) *BaselinesClient {
	return &BaselinesClient{con: con}
}

// List - Lists the metric baseline values for a resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *BaselinesClient) List(ctx context.Context, resourceURI string, options *BaselinesListOptions) (MetricBaselinesResponseResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceURI, options)
	if err != nil {
		return MetricBaselinesResponseResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return MetricBaselinesResponseResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return MetricBaselinesResponseResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *BaselinesClient) listCreateRequest(ctx context.Context, resourceURI string, options *BaselinesListOptions) (*azcore.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Insights/metricBaselines"
	if resourceURI == "" {
		return nil, errors.New("parameter resourceURI cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.Metricnames != nil {
		reqQP.Set("metricnames", *options.Metricnames)
	}
	if options != nil && options.Metricnamespace != nil {
		reqQP.Set("metricnamespace", *options.Metricnamespace)
	}
	if options != nil && options.Timespan != nil {
		reqQP.Set("timespan", *options.Timespan)
	}
	if options != nil && options.Interval != nil {
		reqQP.Set("interval", *options.Interval)
	}
	if options != nil && options.Aggregation != nil {
		reqQP.Set("aggregation", *options.Aggregation)
	}
	if options != nil && options.Sensitivities != nil {
		reqQP.Set("sensitivities", *options.Sensitivities)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.ResultType != nil {
		reqQP.Set("resultType", string(*options.ResultType))
	}
	reqQP.Set("api-version", "2019-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BaselinesClient) listHandleResponse(resp *azcore.Response) (MetricBaselinesResponseResponse, error) {
	var val *MetricBaselinesResponse
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return MetricBaselinesResponseResponse{}, err
	}
	return MetricBaselinesResponseResponse{RawResponse: resp.Response, MetricBaselinesResponse: val}, nil
}

// listHandleError handles the List error response.
func (client *BaselinesClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
