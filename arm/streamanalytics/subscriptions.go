package streamanalytics

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
// Code generated by Microsoft (R) AutoRest Code Generator 2.2.21.0
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// SubscriptionsClient is the stream Analytics Client
type SubscriptionsClient struct {
	ManagementClient
}

// NewSubscriptionsClient creates an instance of the SubscriptionsClient client.
func NewSubscriptionsClient(subscriptionID string) SubscriptionsClient {
	return NewSubscriptionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSubscriptionsClientWithBaseURI creates an instance of the SubscriptionsClient client.
func NewSubscriptionsClientWithBaseURI(baseURI string, subscriptionID string) SubscriptionsClient {
	return SubscriptionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListQuotas retrieves the subscription's current quota information in a particular region.
//
// location is the region in which to retrieve the subscription's quota information. You can find out which regions
// Azure Stream Analytics is supported in here: https://azure.microsoft.com/en-us/regions/
func (client SubscriptionsClient) ListQuotas(location string) (result SubscriptionQuotasListResult, err error) {
	if err := validation.Validate([]validation.Validation{}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "streamanalytics.SubscriptionsClient", "ListQuotas")
	}

	req, err := client.ListQuotasPreparer(location)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.SubscriptionsClient", "ListQuotas", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListQuotasSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "streamanalytics.SubscriptionsClient", "ListQuotas", resp, "Failure sending request")
		return
	}

	result, err = client.ListQuotasResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.SubscriptionsClient", "ListQuotas", resp, "Failure responding to request")
	}

	return
}

// ListQuotasPreparer prepares the ListQuotas request.
func (client SubscriptionsClient) ListQuotasPreparer(location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.StreamAnalytics/locations/{location}/quotas", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListQuotasSender sends the ListQuotas request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionsClient) ListQuotasSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListQuotasResponder handles the response to the ListQuotas request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) ListQuotasResponder(resp *http.Response) (result SubscriptionQuotasListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
