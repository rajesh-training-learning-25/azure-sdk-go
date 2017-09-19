package billing

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

// PeriodsClient is the billing client provides access to billing resources for Azure Web-Direct subscriptions. Other
// subscription types which were not purchased directly through the Azure web portal are not supported through this
// preview API.
type PeriodsClient struct {
	ManagementClient
}

// NewPeriodsClient creates an instance of the PeriodsClient client.
func NewPeriodsClient(subscriptionID string) PeriodsClient {
	return NewPeriodsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPeriodsClientWithBaseURI creates an instance of the PeriodsClient client.
func NewPeriodsClientWithBaseURI(baseURI string, subscriptionID string) PeriodsClient {
	return PeriodsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets a named billing period.
//
// billingPeriodName is the name of a BillingPeriod resource.
func (client PeriodsClient) Get(billingPeriodName string) (result Period, err error) {
	if err := validation.Validate([]validation.Validation{}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "billing.PeriodsClient", "Get")
	}

	req, err := client.GetPreparer(billingPeriodName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PeriodsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.PeriodsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PeriodsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client PeriodsClient) GetPreparer(billingPeriodName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingPeriodName": autorest.Encode("path", billingPeriodName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-24-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Billing/billingPeriods/{billingPeriodName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PeriodsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PeriodsClient) GetResponder(resp *http.Response) (result Period, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists the available billing periods for a subscription in reverse chronological order.
//
// filter is may be used to filter billing periods by billingPeriodEndDate. The filter supports 'eq', 'lt', 'gt', 'le',
// 'ge', and 'and'. It does not currently support 'ne', 'or', or 'not'. skiptoken is skiptoken is only used if a
// previous operation returned a partial result. If a previous response contains a nextLink element, the value of the
// nextLink element will include a skiptoken parameter that specifies a starting point to use for subsequent calls. top
// is may be used to limit the number of results to the most recent N billing periods.
func (client PeriodsClient) List(filter string, skiptoken string, top *int32) (result PeriodsListResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMaximum, Rule: 100, Chain: nil},
					{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "billing.PeriodsClient", "List")
	}

	req, err := client.ListPreparer(filter, skiptoken, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PeriodsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.PeriodsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PeriodsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client PeriodsClient) ListPreparer(filter string, skiptoken string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-24-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", skiptoken)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Billing/billingPeriods", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client PeriodsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client PeriodsClient) ListResponder(resp *http.Response) (result PeriodsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client PeriodsClient) ListNextResults(lastResults PeriodsListResult) (result PeriodsListResult, err error) {
	req, err := lastResults.PeriodsListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.PeriodsClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.PeriodsClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PeriodsClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListComplete gets all elements from the list without paging.
func (client PeriodsClient) ListComplete(filter string, skiptoken string, top *int32, cancel <-chan struct{}) (<-chan Period, <-chan error) {
	resultChan := make(chan Period)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.List(filter, skiptoken, top)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}
