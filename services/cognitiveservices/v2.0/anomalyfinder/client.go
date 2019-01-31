// Package anomalyfinder implements the Azure ARM Anomalyfinder service API version 2.0.
//
// The Anomaly Detection Service detects anomalies automatically in time series data. It supports two functionalities,
// one is for detecting the whole series with model trained by the timeseries, another is detecting last point with
// model trained by points before. By using this service, business customers can discover incidents and establish a
// logic flow for root cause analysis. Ensure online service quality is one of the main reasons we develop this
// service. Improving anomaly detection service to provide precise results will always be fundamental work of our team.
package anomalyfinder

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// BaseClient is the base client for Anomalyfinder.
type BaseClient struct {
	autorest.Client
	Endpoint string
}

// New creates an instance of the BaseClient client.
func New(endpoint string) BaseClient {
	return NewWithoutDefaults(endpoint)
}

// NewWithoutDefaults creates an instance of the BaseClient client.
func NewWithoutDefaults(endpoint string) BaseClient {
	return BaseClient{
		Client:   autorest.NewClientWithUserAgent(UserAgent()),
		Endpoint: endpoint,
	}
}

// PostTimeseriesEntireDetect the operation will generate a model using the entire series, each point will be detected
// with the same model. In this method, points before and after a certain point will be used to determine whether it's
// an anomaly. The entire detection can give user an overall status of the time series.
// Parameters:
// body - time series points and period if needed. Advanced model parameters can also be set in the request.
func (client BaseClient) PostTimeseriesEntireDetect(ctx context.Context, body Request) (result EntireDetectResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.PostTimeseriesEntireDetect")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.Series", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "body.Series", Name: validation.MaxItems, Rule: 8640, Chain: nil},
					{Target: "body.Series", Name: validation.MinItems, Rule: 12, Chain: nil},
				}},
				{Target: "body.Period", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "body.Period", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil}}},
				{Target: "body.MaxAnomalyRatio", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "body.MaxAnomalyRatio", Name: validation.ExclusiveMaximum, Rule: 0.5, Chain: nil},
						{Target: "body.MaxAnomalyRatio", Name: validation.ExclusiveMinimum, Rule: 0, Chain: nil},
					}},
				{Target: "body.Sensitivity", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "body.Sensitivity", Name: validation.InclusiveMaximum, Rule: int64(99), Chain: nil},
						{Target: "body.Sensitivity", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil},
					}}}}}); err != nil {
		return result, validation.NewError("anomalyfinder.BaseClient", "PostTimeseriesEntireDetect", err.Error())
	}

	req, err := client.PostTimeseriesEntireDetectPreparer(ctx, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "anomalyfinder.BaseClient", "PostTimeseriesEntireDetect", nil, "Failure preparing request")
		return
	}

	resp, err := client.PostTimeseriesEntireDetectSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "anomalyfinder.BaseClient", "PostTimeseriesEntireDetect", resp, "Failure sending request")
		return
	}

	result, err = client.PostTimeseriesEntireDetectResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "anomalyfinder.BaseClient", "PostTimeseriesEntireDetect", resp, "Failure responding to request")
	}

	return
}

// PostTimeseriesEntireDetectPreparer prepares the PostTimeseriesEntireDetect request.
func (client BaseClient) PostTimeseriesEntireDetectPreparer(ctx context.Context, body Request) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/anomalyfinder/v2.0", urlParameters),
		autorest.WithPath("/timeseries/entire/detect"),
		autorest.WithJSON(body))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PostTimeseriesEntireDetectSender sends the PostTimeseriesEntireDetect request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) PostTimeseriesEntireDetectSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PostTimeseriesEntireDetectResponder handles the response to the PostTimeseriesEntireDetect request. The method always
// closes the http.Response Body.
func (client BaseClient) PostTimeseriesEntireDetectResponder(resp *http.Response) (result EntireDetectResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PostTimeseriesLastDetect the operation will generate a model using points before the latest one, In this method,
// only history points are used for determine whether the target point is an anomaly. Latest point detecting matches
// the scenario of real-time monitoring of business metrics.
// Parameters:
// body - time series points and period if needed. Advanced model parameters can also be set in the request.
func (client BaseClient) PostTimeseriesLastDetect(ctx context.Context, body Request) (result LastDetectResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.PostTimeseriesLastDetect")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.Series", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "body.Series", Name: validation.MaxItems, Rule: 8640, Chain: nil},
					{Target: "body.Series", Name: validation.MinItems, Rule: 12, Chain: nil},
				}},
				{Target: "body.Period", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "body.Period", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil}}},
				{Target: "body.MaxAnomalyRatio", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "body.MaxAnomalyRatio", Name: validation.ExclusiveMaximum, Rule: 0.5, Chain: nil},
						{Target: "body.MaxAnomalyRatio", Name: validation.ExclusiveMinimum, Rule: 0, Chain: nil},
					}},
				{Target: "body.Sensitivity", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "body.Sensitivity", Name: validation.InclusiveMaximum, Rule: int64(99), Chain: nil},
						{Target: "body.Sensitivity", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil},
					}}}}}); err != nil {
		return result, validation.NewError("anomalyfinder.BaseClient", "PostTimeseriesLastDetect", err.Error())
	}

	req, err := client.PostTimeseriesLastDetectPreparer(ctx, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "anomalyfinder.BaseClient", "PostTimeseriesLastDetect", nil, "Failure preparing request")
		return
	}

	resp, err := client.PostTimeseriesLastDetectSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "anomalyfinder.BaseClient", "PostTimeseriesLastDetect", resp, "Failure sending request")
		return
	}

	result, err = client.PostTimeseriesLastDetectResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "anomalyfinder.BaseClient", "PostTimeseriesLastDetect", resp, "Failure responding to request")
	}

	return
}

// PostTimeseriesLastDetectPreparer prepares the PostTimeseriesLastDetect request.
func (client BaseClient) PostTimeseriesLastDetectPreparer(ctx context.Context, body Request) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/anomalyfinder/v2.0", urlParameters),
		autorest.WithPath("/timeseries/last/detect"),
		autorest.WithJSON(body))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PostTimeseriesLastDetectSender sends the PostTimeseriesLastDetect request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) PostTimeseriesLastDetectSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PostTimeseriesLastDetectResponder handles the response to the PostTimeseriesLastDetect request. The method always
// closes the http.Response Body.
func (client BaseClient) PostTimeseriesLastDetectResponder(resp *http.Response) (result LastDetectResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
