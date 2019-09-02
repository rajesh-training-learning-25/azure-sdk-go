package security

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

// RegulatoryComplianceStandardsClient is the API spec for Microsoft.Security (Azure Security Center) resource provider
type RegulatoryComplianceStandardsClient struct {
	BaseClient
}

// NewRegulatoryComplianceStandardsClient creates an instance of the RegulatoryComplianceStandardsClient client.
func NewRegulatoryComplianceStandardsClient(subscriptionID string, ascLocation string) RegulatoryComplianceStandardsClient {
	return NewRegulatoryComplianceStandardsClientWithBaseURI(DefaultBaseURI, subscriptionID, ascLocation)
}

// NewRegulatoryComplianceStandardsClientWithBaseURI creates an instance of the RegulatoryComplianceStandardsClient
// client.
func NewRegulatoryComplianceStandardsClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) RegulatoryComplianceStandardsClient {
	return RegulatoryComplianceStandardsClient{NewWithBaseURI(baseURI, subscriptionID, ascLocation)}
}

// Get supported regulatory compliance details state for selected standard
// Parameters:
// regulatoryComplianceStandardName - name of the regulatory compliance standard object
func (client RegulatoryComplianceStandardsClient) Get(ctx context.Context, regulatoryComplianceStandardName string) (result RegulatoryComplianceStandard, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegulatoryComplianceStandardsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.RegulatoryComplianceStandardsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, regulatoryComplianceStandardName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.RegulatoryComplianceStandardsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.RegulatoryComplianceStandardsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.RegulatoryComplianceStandardsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client RegulatoryComplianceStandardsClient) GetPreparer(ctx context.Context, regulatoryComplianceStandardName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"regulatoryComplianceStandardName": autorest.Encode("path", regulatoryComplianceStandardName),
		"subscriptionId":                   autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/regulatoryComplianceStandards/{regulatoryComplianceStandardName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RegulatoryComplianceStandardsClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RegulatoryComplianceStandardsClient) GetResponder(resp *http.Response) (result RegulatoryComplianceStandard, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List supported regulatory compliance standards details and state
// Parameters:
// filter - oData filter. Optional.
func (client RegulatoryComplianceStandardsClient) List(ctx context.Context, filter string) (result RegulatoryComplianceStandardListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegulatoryComplianceStandardsClient.List")
		defer func() {
			sc := -1
			if result.rcsl.Response.Response != nil {
				sc = result.rcsl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.RegulatoryComplianceStandardsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.RegulatoryComplianceStandardsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rcsl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.RegulatoryComplianceStandardsClient", "List", resp, "Failure sending request")
		return
	}

	result.rcsl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.RegulatoryComplianceStandardsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client RegulatoryComplianceStandardsClient) ListPreparer(ctx context.Context, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/regulatoryComplianceStandards", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client RegulatoryComplianceStandardsClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client RegulatoryComplianceStandardsClient) ListResponder(resp *http.Response) (result RegulatoryComplianceStandardList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client RegulatoryComplianceStandardsClient) listNextResults(ctx context.Context, lastResults RegulatoryComplianceStandardList) (result RegulatoryComplianceStandardList, err error) {
	req, err := lastResults.regulatoryComplianceStandardListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "security.RegulatoryComplianceStandardsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "security.RegulatoryComplianceStandardsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.RegulatoryComplianceStandardsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client RegulatoryComplianceStandardsClient) ListComplete(ctx context.Context, filter string) (result RegulatoryComplianceStandardListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegulatoryComplianceStandardsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, filter)
	return
}
