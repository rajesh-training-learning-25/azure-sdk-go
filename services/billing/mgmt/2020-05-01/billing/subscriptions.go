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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// SubscriptionsClient is the billing client provides access to billing resources for Azure subscriptions.
type SubscriptionsClient struct {
	BaseClient
}

// NewSubscriptionsClient creates an instance of the SubscriptionsClient client.
func NewSubscriptionsClient(subscriptionID string, subscriptionID1 string) SubscriptionsClient {
	return NewSubscriptionsClientWithBaseURI(DefaultBaseURI, subscriptionID, subscriptionID1)
}

// NewSubscriptionsClientWithBaseURI creates an instance of the SubscriptionsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewSubscriptionsClientWithBaseURI(baseURI string, subscriptionID string, subscriptionID1 string) SubscriptionsClient {
	return SubscriptionsClient{NewWithBaseURI(baseURI, subscriptionID, subscriptionID1)}
}

// Get gets a subscription by its ID. The operation is supported for billing accounts with agreement type Microsoft
// Customer Agreement and Microsoft Partner Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
func (client SubscriptionsClient) Get(ctx context.Context, billingAccountName string) (result Subscription, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, billingAccountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client SubscriptionsClient) GetPreparer(ctx context.Context, billingAccountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingSubscriptions/{subscriptionId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) GetResponder(resp *http.Response) (result Subscription, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByBillingAccount lists the subscriptions for a billing account. The operation is supported for billing accounts
// with agreement type Microsoft Customer Agreement or Microsoft Partner Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
func (client SubscriptionsClient) ListByBillingAccount(ctx context.Context, billingAccountName string) (result SubscriptionsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionsClient.ListByBillingAccount")
		defer func() {
			sc := -1
			if result.slr.Response.Response != nil {
				sc = result.slr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByBillingAccountNextResults
	req, err := client.ListByBillingAccountPreparer(ctx, billingAccountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "ListByBillingAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBillingAccountSender(req)
	if err != nil {
		result.slr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "ListByBillingAccount", resp, "Failure sending request")
		return
	}

	result.slr, err = client.ListByBillingAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "ListByBillingAccount", resp, "Failure responding to request")
	}

	return
}

// ListByBillingAccountPreparer prepares the ListByBillingAccount request.
func (client SubscriptionsClient) ListByBillingAccountPreparer(ctx context.Context, billingAccountName string) (*http.Request, error) {
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
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingSubscriptions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByBillingAccountSender sends the ListByBillingAccount request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionsClient) ListByBillingAccountSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByBillingAccountResponder handles the response to the ListByBillingAccount request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) ListByBillingAccountResponder(resp *http.Response) (result SubscriptionsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByBillingAccountNextResults retrieves the next set of results, if any.
func (client SubscriptionsClient) listByBillingAccountNextResults(ctx context.Context, lastResults SubscriptionsListResult) (result SubscriptionsListResult, err error) {
	req, err := lastResults.subscriptionsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "listByBillingAccountNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByBillingAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "listByBillingAccountNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByBillingAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "listByBillingAccountNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByBillingAccountComplete enumerates all values, automatically crossing page boundaries as required.
func (client SubscriptionsClient) ListByBillingAccountComplete(ctx context.Context, billingAccountName string) (result SubscriptionsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionsClient.ListByBillingAccount")
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

// ListByBillingProfile lists the subscriptions that are billed to a billing profile. The operation is supported for
// billing accounts with agreement type Microsoft Customer Agreement or Microsoft Partner Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// billingProfileName - the ID that uniquely identifies a billing profile.
func (client SubscriptionsClient) ListByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string) (result SubscriptionsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionsClient.ListByBillingProfile")
		defer func() {
			sc := -1
			if result.slr.Response.Response != nil {
				sc = result.slr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByBillingProfileNextResults
	req, err := client.ListByBillingProfilePreparer(ctx, billingAccountName, billingProfileName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "ListByBillingProfile", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBillingProfileSender(req)
	if err != nil {
		result.slr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "ListByBillingProfile", resp, "Failure sending request")
		return
	}

	result.slr, err = client.ListByBillingProfileResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "ListByBillingProfile", resp, "Failure responding to request")
	}

	return
}

// ListByBillingProfilePreparer prepares the ListByBillingProfile request.
func (client SubscriptionsClient) ListByBillingProfilePreparer(ctx context.Context, billingAccountName string, billingProfileName string) (*http.Request, error) {
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
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/billingSubscriptions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByBillingProfileSender sends the ListByBillingProfile request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionsClient) ListByBillingProfileSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByBillingProfileResponder handles the response to the ListByBillingProfile request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) ListByBillingProfileResponder(resp *http.Response) (result SubscriptionsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByBillingProfileNextResults retrieves the next set of results, if any.
func (client SubscriptionsClient) listByBillingProfileNextResults(ctx context.Context, lastResults SubscriptionsListResult) (result SubscriptionsListResult, err error) {
	req, err := lastResults.subscriptionsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "listByBillingProfileNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByBillingProfileSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "listByBillingProfileNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByBillingProfileResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "listByBillingProfileNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByBillingProfileComplete enumerates all values, automatically crossing page boundaries as required.
func (client SubscriptionsClient) ListByBillingProfileComplete(ctx context.Context, billingAccountName string, billingProfileName string) (result SubscriptionsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionsClient.ListByBillingProfile")
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

// ListByCustomer lists the subscriptions for a customer. The operation is supported only for billing accounts with
// agreement type Microsoft Partner Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// customerName - the ID that uniquely identifies a customer.
func (client SubscriptionsClient) ListByCustomer(ctx context.Context, billingAccountName string, customerName string) (result SubscriptionsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionsClient.ListByCustomer")
		defer func() {
			sc := -1
			if result.slr.Response.Response != nil {
				sc = result.slr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByCustomerNextResults
	req, err := client.ListByCustomerPreparer(ctx, billingAccountName, customerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "ListByCustomer", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByCustomerSender(req)
	if err != nil {
		result.slr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "ListByCustomer", resp, "Failure sending request")
		return
	}

	result.slr, err = client.ListByCustomerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "ListByCustomer", resp, "Failure responding to request")
	}

	return
}

// ListByCustomerPreparer prepares the ListByCustomer request.
func (client SubscriptionsClient) ListByCustomerPreparer(ctx context.Context, billingAccountName string, customerName string) (*http.Request, error) {
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
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/customers/{customerName}/billingSubscriptions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByCustomerSender sends the ListByCustomer request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionsClient) ListByCustomerSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByCustomerResponder handles the response to the ListByCustomer request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) ListByCustomerResponder(resp *http.Response) (result SubscriptionsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByCustomerNextResults retrieves the next set of results, if any.
func (client SubscriptionsClient) listByCustomerNextResults(ctx context.Context, lastResults SubscriptionsListResult) (result SubscriptionsListResult, err error) {
	req, err := lastResults.subscriptionsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "listByCustomerNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByCustomerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "listByCustomerNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByCustomerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "listByCustomerNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByCustomerComplete enumerates all values, automatically crossing page boundaries as required.
func (client SubscriptionsClient) ListByCustomerComplete(ctx context.Context, billingAccountName string, customerName string) (result SubscriptionsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionsClient.ListByCustomer")
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

// ListByInvoiceSection lists the subscriptions that are billed to an invoice section. The operation is supported only
// for billing accounts with agreement type Microsoft Customer Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// billingProfileName - the ID that uniquely identifies a billing profile.
// invoiceSectionName - the ID that uniquely identifies an invoice section.
func (client SubscriptionsClient) ListByInvoiceSection(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string) (result SubscriptionsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionsClient.ListByInvoiceSection")
		defer func() {
			sc := -1
			if result.slr.Response.Response != nil {
				sc = result.slr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByInvoiceSectionNextResults
	req, err := client.ListByInvoiceSectionPreparer(ctx, billingAccountName, billingProfileName, invoiceSectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "ListByInvoiceSection", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByInvoiceSectionSender(req)
	if err != nil {
		result.slr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "ListByInvoiceSection", resp, "Failure sending request")
		return
	}

	result.slr, err = client.ListByInvoiceSectionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "ListByInvoiceSection", resp, "Failure responding to request")
	}

	return
}

// ListByInvoiceSectionPreparer prepares the ListByInvoiceSection request.
func (client SubscriptionsClient) ListByInvoiceSectionPreparer(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string) (*http.Request, error) {
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
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/billingSubscriptions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByInvoiceSectionSender sends the ListByInvoiceSection request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionsClient) ListByInvoiceSectionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByInvoiceSectionResponder handles the response to the ListByInvoiceSection request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) ListByInvoiceSectionResponder(resp *http.Response) (result SubscriptionsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByInvoiceSectionNextResults retrieves the next set of results, if any.
func (client SubscriptionsClient) listByInvoiceSectionNextResults(ctx context.Context, lastResults SubscriptionsListResult) (result SubscriptionsListResult, err error) {
	req, err := lastResults.subscriptionsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "listByInvoiceSectionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByInvoiceSectionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "listByInvoiceSectionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByInvoiceSectionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "listByInvoiceSectionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByInvoiceSectionComplete enumerates all values, automatically crossing page boundaries as required.
func (client SubscriptionsClient) ListByInvoiceSectionComplete(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string) (result SubscriptionsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionsClient.ListByInvoiceSection")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByInvoiceSection(ctx, billingAccountName, billingProfileName, invoiceSectionName)
	return
}

// Move moves a subscription's charges to a new invoice section. The new invoice section must belong to the same
// billing profile as the existing invoice section. This operation is supported for billing accounts with agreement
// type Microsoft Customer Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// parameters - request parameters that are provided to the move subscription operation.
func (client SubscriptionsClient) Move(ctx context.Context, billingAccountName string, parameters TransferBillingSubscriptionRequestProperties) (result SubscriptionsMoveFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionsClient.Move")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.DestinationInvoiceSectionID", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("billing.SubscriptionsClient", "Move", err.Error())
	}

	req, err := client.MovePreparer(ctx, billingAccountName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "Move", nil, "Failure preparing request")
		return
	}

	result, err = client.MoveSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "Move", result.Response(), "Failure sending request")
		return
	}

	return
}

// MovePreparer prepares the Move request.
func (client SubscriptionsClient) MovePreparer(ctx context.Context, billingAccountName string, parameters TransferBillingSubscriptionRequestProperties) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingSubscriptions/{subscriptionId}/move", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// MoveSender sends the Move request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionsClient) MoveSender(req *http.Request) (future SubscriptionsMoveFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// MoveResponder handles the response to the Move request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) MoveResponder(resp *http.Response) (result Subscription, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates the properties of a billing subscription. Currently, cost center can be updated. The operation is
// supported only for billing accounts with agreement type Microsoft Customer Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// parameters - request parameters that are provided to the update billing subscription operation.
func (client SubscriptionsClient) Update(ctx context.Context, billingAccountName string, parameters Subscription) (result Subscription, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, billingAccountName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client SubscriptionsClient) UpdatePreparer(ctx context.Context, billingAccountName string, parameters Subscription) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingSubscriptions/{subscriptionId}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) UpdateResponder(resp *http.Response) (result Subscription, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ValidateMove validates if a subscription's charges can be moved to a new invoice section. This operation is
// supported for billing accounts with agreement type Microsoft Customer Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// parameters - request parameters that are provided to the validate move eligibility operation.
func (client SubscriptionsClient) ValidateMove(ctx context.Context, billingAccountName string, parameters TransferBillingSubscriptionRequestProperties) (result ValidateSubscriptionTransferEligibilityResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionsClient.ValidateMove")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.DestinationInvoiceSectionID", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("billing.SubscriptionsClient", "ValidateMove", err.Error())
	}

	req, err := client.ValidateMovePreparer(ctx, billingAccountName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "ValidateMove", nil, "Failure preparing request")
		return
	}

	resp, err := client.ValidateMoveSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "ValidateMove", resp, "Failure sending request")
		return
	}

	result, err = client.ValidateMoveResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsClient", "ValidateMove", resp, "Failure responding to request")
	}

	return
}

// ValidateMovePreparer prepares the ValidateMove request.
func (client SubscriptionsClient) ValidateMovePreparer(ctx context.Context, billingAccountName string, parameters TransferBillingSubscriptionRequestProperties) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingSubscriptions/{subscriptionId}/validateMoveEligibility", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ValidateMoveSender sends the ValidateMove request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionsClient) ValidateMoveSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ValidateMoveResponder handles the response to the ValidateMove request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) ValidateMoveResponder(resp *http.Response) (result ValidateSubscriptionTransferEligibilityResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
