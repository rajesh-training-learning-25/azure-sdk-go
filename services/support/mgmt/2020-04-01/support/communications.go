package support

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

// CommunicationsClient is the microsoft Azure Support Resource Provider.
type CommunicationsClient struct {
	BaseClient
}

// NewCommunicationsClient creates an instance of the CommunicationsClient client.
func NewCommunicationsClient(subscriptionID string) CommunicationsClient {
	return NewCommunicationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCommunicationsClientWithBaseURI creates an instance of the CommunicationsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewCommunicationsClientWithBaseURI(baseURI string, subscriptionID string) CommunicationsClient {
	return CommunicationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CheckNameAvailability check the availability of a resource name. This API should be used to check the uniqueness of
// the name for adding a new communication to the support ticket.
// Parameters:
// supportTicketName - support ticket name.
// checkNameAvailabilityInput - input to check.
func (client CommunicationsClient) CheckNameAvailability(ctx context.Context, supportTicketName string, checkNameAvailabilityInput CheckNameAvailabilityInput) (result CheckNameAvailabilityOutput, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CommunicationsClient.CheckNameAvailability")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: checkNameAvailabilityInput,
			Constraints: []validation.Constraint{{Target: "checkNameAvailabilityInput.Name", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("support.CommunicationsClient", "CheckNameAvailability", err.Error())
	}

	req, err := client.CheckNameAvailabilityPreparer(ctx, supportTicketName, checkNameAvailabilityInput)
	if err != nil {
		err = autorest.NewErrorWithError(err, "support.CommunicationsClient", "CheckNameAvailability", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckNameAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "support.CommunicationsClient", "CheckNameAvailability", resp, "Failure sending request")
		return
	}

	result, err = client.CheckNameAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "support.CommunicationsClient", "CheckNameAvailability", resp, "Failure responding to request")
		return
	}

	return
}

// CheckNameAvailabilityPreparer prepares the CheckNameAvailability request.
func (client CommunicationsClient) CheckNameAvailabilityPreparer(ctx context.Context, supportTicketName string, checkNameAvailabilityInput CheckNameAvailabilityInput) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"supportTicketName": autorest.Encode("path", supportTicketName),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Support/supportTickets/{supportTicketName}/checkNameAvailability", pathParameters),
		autorest.WithJSON(checkNameAvailabilityInput),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckNameAvailabilitySender sends the CheckNameAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client CommunicationsClient) CheckNameAvailabilitySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CheckNameAvailabilityResponder handles the response to the CheckNameAvailability request. The method always
// closes the http.Response Body.
func (client CommunicationsClient) CheckNameAvailabilityResponder(resp *http.Response) (result CheckNameAvailabilityOutput, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Create adds a new customer communication to an Azure support ticket.
// Parameters:
// supportTicketName - support ticket name.
// communicationName - communication name.
// createCommunicationParameters - communication object.
func (client CommunicationsClient) Create(ctx context.Context, supportTicketName string, communicationName string, createCommunicationParameters CommunicationDetails) (result CommunicationsCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CommunicationsClient.Create")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: createCommunicationParameters,
			Constraints: []validation.Constraint{{Target: "createCommunicationParameters.CommunicationDetailsProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "createCommunicationParameters.CommunicationDetailsProperties.Subject", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "createCommunicationParameters.CommunicationDetailsProperties.Body", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("support.CommunicationsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, supportTicketName, communicationName, createCommunicationParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "support.CommunicationsClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "support.CommunicationsClient", "Create", nil, "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client CommunicationsClient) CreatePreparer(ctx context.Context, supportTicketName string, communicationName string, createCommunicationParameters CommunicationDetails) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"communicationName": autorest.Encode("path", communicationName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"supportTicketName": autorest.Encode("path", supportTicketName),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	createCommunicationParameters.ID = nil
	createCommunicationParameters.Name = nil
	createCommunicationParameters.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Support/supportTickets/{supportTicketName}/communications/{communicationName}", pathParameters),
		autorest.WithJSON(createCommunicationParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client CommunicationsClient) CreateSender(req *http.Request) (future CommunicationsCreateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client CommunicationsClient) (cd CommunicationDetails, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "support.CommunicationsCreateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("support.CommunicationsCreateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		cd.Response.Response, err = future.GetResult(sender)
		if cd.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "support.CommunicationsCreateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && cd.Response.Response.StatusCode != http.StatusNoContent {
			cd, err = client.CreateResponder(cd.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "support.CommunicationsCreateFuture", "Result", cd.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client CommunicationsClient) CreateResponder(resp *http.Response) (result CommunicationDetails, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get returns communication details for a support ticket.
// Parameters:
// supportTicketName - support ticket name.
// communicationName - communication name.
func (client CommunicationsClient) Get(ctx context.Context, supportTicketName string, communicationName string) (result CommunicationDetails, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CommunicationsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, supportTicketName, communicationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "support.CommunicationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "support.CommunicationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "support.CommunicationsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client CommunicationsClient) GetPreparer(ctx context.Context, supportTicketName string, communicationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"communicationName": autorest.Encode("path", communicationName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"supportTicketName": autorest.Encode("path", supportTicketName),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Support/supportTickets/{supportTicketName}/communications/{communicationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client CommunicationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client CommunicationsClient) GetResponder(resp *http.Response) (result CommunicationDetails, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all communications (attachments not included) for a support ticket. <br/></br> You can also filter
// support ticket communications by _CreatedDate_ or _CommunicationType_ using the $filter parameter. The only type of
// communication supported today is _Web_. Output will be a paged result with _nextLink_, using which you can retrieve
// the next set of Communication results. <br/><br/>Support ticket data is available for 18 months after ticket
// creation. If a ticket was created more than 18 months ago, a request for data might cause an error.
// Parameters:
// supportTicketName - support ticket name.
// top - the number of values to return in the collection. Default is 10 and max is 10.
// filter - the filter to apply on the operation. You can filter by communicationType and createdDate
// properties. CommunicationType supports Equals ('eq') operator and createdDate supports Greater Than ('gt')
// and Greater Than or Equals ('ge') operators. You may combine the CommunicationType and CreatedDate filters
// by Logical And ('and') operator.
func (client CommunicationsClient) List(ctx context.Context, supportTicketName string, top *int32, filter string) (result CommunicationsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CommunicationsClient.List")
		defer func() {
			sc := -1
			if result.clr.Response.Response != nil {
				sc = result.clr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, supportTicketName, top, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "support.CommunicationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.clr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "support.CommunicationsClient", "List", resp, "Failure sending request")
		return
	}

	result.clr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "support.CommunicationsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.clr.hasNextLink() && result.clr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client CommunicationsClient) ListPreparer(ctx context.Context, supportTicketName string, top *int32, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"supportTicketName": autorest.Encode("path", supportTicketName),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Support/supportTickets/{supportTicketName}/communications", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client CommunicationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client CommunicationsClient) ListResponder(resp *http.Response) (result CommunicationsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client CommunicationsClient) listNextResults(ctx context.Context, lastResults CommunicationsListResult) (result CommunicationsListResult, err error) {
	req, err := lastResults.communicationsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "support.CommunicationsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "support.CommunicationsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "support.CommunicationsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client CommunicationsClient) ListComplete(ctx context.Context, supportTicketName string, top *int32, filter string) (result CommunicationsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CommunicationsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, supportTicketName, top, filter)
	return
}
