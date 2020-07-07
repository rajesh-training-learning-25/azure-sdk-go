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
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// PermissionsClient is the billing client provides access to billing resources for Azure subscriptions.
type PermissionsClient struct {
	BaseClient
}

// NewPermissionsClient creates an instance of the PermissionsClient client.
func NewPermissionsClient(subscriptionID string, subscriptionID1 string) PermissionsClient {
	return NewPermissionsClientWithBaseURI(DefaultBaseURI, subscriptionID, subscriptionID1)
}

// NewPermissionsClientWithBaseURI creates an instance of the PermissionsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewPermissionsClientWithBaseURI(baseURI string, subscriptionID string, subscriptionID1 string) PermissionsClient {
	return PermissionsClient{NewWithBaseURI(baseURI, subscriptionID, subscriptionID1)}
}

// ListByBillingAccount lists the billing permissions the caller has on a billing account.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
func (client PermissionsClient) ListByBillingAccount(ctx context.Context, billingAccountName string) (result PermissionsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PermissionsClient.ListByBillingAccount")
		defer func() {
			sc := -1
			if result.plr.Response.Response != nil {
				sc = result.plr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByBillingAccountNextResults
	req, err := client.ListByBillingAccountPreparer(ctx, billingAccountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PermissionsClient", "ListByBillingAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBillingAccountSender(req)
	if err != nil {
		result.plr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.PermissionsClient", "ListByBillingAccount", resp, "Failure sending request")
		return
	}

	result.plr, err = client.ListByBillingAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PermissionsClient", "ListByBillingAccount", resp, "Failure responding to request")
	}

	return
}

// ListByBillingAccountPreparer prepares the ListByBillingAccount request.
func (client PermissionsClient) ListByBillingAccountPreparer(ctx context.Context, billingAccountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingPermissions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByBillingAccountSender sends the ListByBillingAccount request. The method will close the
// http.Response Body if it receives an error.
func (client PermissionsClient) ListByBillingAccountSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByBillingAccountResponder handles the response to the ListByBillingAccount request. The method always
// closes the http.Response Body.
func (client PermissionsClient) ListByBillingAccountResponder(resp *http.Response) (result PermissionsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByBillingAccountNextResults retrieves the next set of results, if any.
func (client PermissionsClient) listByBillingAccountNextResults(ctx context.Context, lastResults PermissionsListResult) (result PermissionsListResult, err error) {
	req, err := lastResults.permissionsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.PermissionsClient", "listByBillingAccountNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByBillingAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.PermissionsClient", "listByBillingAccountNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByBillingAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PermissionsClient", "listByBillingAccountNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByBillingAccountComplete enumerates all values, automatically crossing page boundaries as required.
func (client PermissionsClient) ListByBillingAccountComplete(ctx context.Context, billingAccountName string) (result PermissionsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PermissionsClient.ListByBillingAccount")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByBillingAccount(ctx, billingAccountName)
	return
}

// ListByBillingProfile lists the billing permissions the caller has on a billing profile.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// billingProfileName - the ID that uniquely identifies a billing profile.
func (client PermissionsClient) ListByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string) (result PermissionsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PermissionsClient.ListByBillingProfile")
		defer func() {
			sc := -1
			if result.plr.Response.Response != nil {
				sc = result.plr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByBillingProfileNextResults
	req, err := client.ListByBillingProfilePreparer(ctx, billingAccountName, billingProfileName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PermissionsClient", "ListByBillingProfile", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBillingProfileSender(req)
	if err != nil {
		result.plr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.PermissionsClient", "ListByBillingProfile", resp, "Failure sending request")
		return
	}

	result.plr, err = client.ListByBillingProfileResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PermissionsClient", "ListByBillingProfile", resp, "Failure responding to request")
	}

	return
}

// ListByBillingProfilePreparer prepares the ListByBillingProfile request.
func (client PermissionsClient) ListByBillingProfilePreparer(ctx context.Context, billingAccountName string, billingProfileName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"billingProfileName": autorest.Encode("path", billingProfileName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/billingPermissions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByBillingProfileSender sends the ListByBillingProfile request. The method will close the
// http.Response Body if it receives an error.
func (client PermissionsClient) ListByBillingProfileSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByBillingProfileResponder handles the response to the ListByBillingProfile request. The method always
// closes the http.Response Body.
func (client PermissionsClient) ListByBillingProfileResponder(resp *http.Response) (result PermissionsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByBillingProfileNextResults retrieves the next set of results, if any.
func (client PermissionsClient) listByBillingProfileNextResults(ctx context.Context, lastResults PermissionsListResult) (result PermissionsListResult, err error) {
	req, err := lastResults.permissionsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.PermissionsClient", "listByBillingProfileNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByBillingProfileSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.PermissionsClient", "listByBillingProfileNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByBillingProfileResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PermissionsClient", "listByBillingProfileNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByBillingProfileComplete enumerates all values, automatically crossing page boundaries as required.
func (client PermissionsClient) ListByBillingProfileComplete(ctx context.Context, billingAccountName string, billingProfileName string) (result PermissionsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PermissionsClient.ListByBillingProfile")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByBillingProfile(ctx, billingAccountName, billingProfileName)
	return
}

// ListByCustomer lists the billing permissions the caller has for a customer.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// customerName - the ID that uniquely identifies a customer.
func (client PermissionsClient) ListByCustomer(ctx context.Context, billingAccountName string, customerName string) (result PermissionsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PermissionsClient.ListByCustomer")
		defer func() {
			sc := -1
			if result.plr.Response.Response != nil {
				sc = result.plr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByCustomerNextResults
	req, err := client.ListByCustomerPreparer(ctx, billingAccountName, customerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PermissionsClient", "ListByCustomer", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByCustomerSender(req)
	if err != nil {
		result.plr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.PermissionsClient", "ListByCustomer", resp, "Failure sending request")
		return
	}

	result.plr, err = client.ListByCustomerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PermissionsClient", "ListByCustomer", resp, "Failure responding to request")
	}

	return
}

// ListByCustomerPreparer prepares the ListByCustomer request.
func (client PermissionsClient) ListByCustomerPreparer(ctx context.Context, billingAccountName string, customerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"customerName":       autorest.Encode("path", customerName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/customers/{customerName}/billingPermissions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByCustomerSender sends the ListByCustomer request. The method will close the
// http.Response Body if it receives an error.
func (client PermissionsClient) ListByCustomerSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByCustomerResponder handles the response to the ListByCustomer request. The method always
// closes the http.Response Body.
func (client PermissionsClient) ListByCustomerResponder(resp *http.Response) (result PermissionsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByCustomerNextResults retrieves the next set of results, if any.
func (client PermissionsClient) listByCustomerNextResults(ctx context.Context, lastResults PermissionsListResult) (result PermissionsListResult, err error) {
	req, err := lastResults.permissionsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.PermissionsClient", "listByCustomerNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByCustomerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.PermissionsClient", "listByCustomerNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByCustomerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PermissionsClient", "listByCustomerNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByCustomerComplete enumerates all values, automatically crossing page boundaries as required.
func (client PermissionsClient) ListByCustomerComplete(ctx context.Context, billingAccountName string, customerName string) (result PermissionsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PermissionsClient.ListByCustomer")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByCustomer(ctx, billingAccountName, customerName)
	return
}

// ListByInvoiceSections lists the billing permissions the caller has on an invoice section.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// billingProfileName - the ID that uniquely identifies a billing profile.
// invoiceSectionName - the ID that uniquely identifies an invoice section.
func (client PermissionsClient) ListByInvoiceSections(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string) (result PermissionsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PermissionsClient.ListByInvoiceSections")
		defer func() {
			sc := -1
			if result.plr.Response.Response != nil {
				sc = result.plr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByInvoiceSectionsNextResults
	req, err := client.ListByInvoiceSectionsPreparer(ctx, billingAccountName, billingProfileName, invoiceSectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PermissionsClient", "ListByInvoiceSections", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByInvoiceSectionsSender(req)
	if err != nil {
		result.plr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.PermissionsClient", "ListByInvoiceSections", resp, "Failure sending request")
		return
	}

	result.plr, err = client.ListByInvoiceSectionsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PermissionsClient", "ListByInvoiceSections", resp, "Failure responding to request")
	}

	return
}

// ListByInvoiceSectionsPreparer prepares the ListByInvoiceSections request.
func (client PermissionsClient) ListByInvoiceSectionsPreparer(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"billingProfileName": autorest.Encode("path", billingProfileName),
		"invoiceSectionName": autorest.Encode("path", invoiceSectionName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/billingPermissions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByInvoiceSectionsSender sends the ListByInvoiceSections request. The method will close the
// http.Response Body if it receives an error.
func (client PermissionsClient) ListByInvoiceSectionsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByInvoiceSectionsResponder handles the response to the ListByInvoiceSections request. The method always
// closes the http.Response Body.
func (client PermissionsClient) ListByInvoiceSectionsResponder(resp *http.Response) (result PermissionsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByInvoiceSectionsNextResults retrieves the next set of results, if any.
func (client PermissionsClient) listByInvoiceSectionsNextResults(ctx context.Context, lastResults PermissionsListResult) (result PermissionsListResult, err error) {
	req, err := lastResults.permissionsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.PermissionsClient", "listByInvoiceSectionsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByInvoiceSectionsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.PermissionsClient", "listByInvoiceSectionsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByInvoiceSectionsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PermissionsClient", "listByInvoiceSectionsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByInvoiceSectionsComplete enumerates all values, automatically crossing page boundaries as required.
func (client PermissionsClient) ListByInvoiceSectionsComplete(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string) (result PermissionsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PermissionsClient.ListByInvoiceSections")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByInvoiceSections(ctx, billingAccountName, billingProfileName, invoiceSectionName)
	return
}
