package eventgrid

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
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

// SystemTopicEventSubscriptionsClient is the azure EventGrid Management Client
type SystemTopicEventSubscriptionsClient struct {
	BaseClient
}

// NewSystemTopicEventSubscriptionsClient creates an instance of the SystemTopicEventSubscriptionsClient client.
func NewSystemTopicEventSubscriptionsClient(subscriptionID string) SystemTopicEventSubscriptionsClient {
	return NewSystemTopicEventSubscriptionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSystemTopicEventSubscriptionsClientWithBaseURI creates an instance of the SystemTopicEventSubscriptionsClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewSystemTopicEventSubscriptionsClientWithBaseURI(baseURI string, subscriptionID string) SystemTopicEventSubscriptionsClient {
	return SystemTopicEventSubscriptionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate asynchronously creates or updates an event subscription with the specified parameters. Existing event
// subscriptions will be updated with this API.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// systemTopicName - name of the system topic.
// eventSubscriptionName - name of the event subscription to be created. Event subscription names must be
// between 3 and 100 characters in length and use alphanumeric letters only.
// eventSubscriptionInfo - event subscription properties containing the destination and filter information.
func (client SystemTopicEventSubscriptionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string, eventSubscriptionInfo EventSubscription) (result SystemTopicEventSubscriptionsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SystemTopicEventSubscriptionsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, systemTopicName, eventSubscriptionName, eventSubscriptionInfo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.SystemTopicEventSubscriptionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.SystemTopicEventSubscriptionsClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SystemTopicEventSubscriptionsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string, eventSubscriptionInfo EventSubscription) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"eventSubscriptionName": autorest.Encode("path", eventSubscriptionName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
		"systemTopicName":       autorest.Encode("path", systemTopicName),
	}

	const APIVersion = "2020-10-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	eventSubscriptionInfo.SystemData = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics/{systemTopicName}/eventSubscriptions/{eventSubscriptionName}", pathParameters),
		autorest.WithJSON(eventSubscriptionInfo),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SystemTopicEventSubscriptionsClient) CreateOrUpdateSender(req *http.Request) (future SystemTopicEventSubscriptionsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SystemTopicEventSubscriptionsClient) CreateOrUpdateResponder(resp *http.Response) (result EventSubscription, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete an event subscription of a system topic.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// systemTopicName - name of the system topic.
// eventSubscriptionName - name of the event subscription to be created. Event subscription names must be
// between 3 and 100 characters in length and use alphanumeric letters only.
func (client SystemTopicEventSubscriptionsClient) Delete(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string) (result SystemTopicEventSubscriptionsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SystemTopicEventSubscriptionsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, systemTopicName, eventSubscriptionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.SystemTopicEventSubscriptionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.SystemTopicEventSubscriptionsClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SystemTopicEventSubscriptionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"eventSubscriptionName": autorest.Encode("path", eventSubscriptionName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
		"systemTopicName":       autorest.Encode("path", systemTopicName),
	}

	const APIVersion = "2020-10-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics/{systemTopicName}/eventSubscriptions/{eventSubscriptionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SystemTopicEventSubscriptionsClient) DeleteSender(req *http.Request) (future SystemTopicEventSubscriptionsDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SystemTopicEventSubscriptionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get an event subscription.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// systemTopicName - name of the system topic.
// eventSubscriptionName - name of the event subscription to be created. Event subscription names must be
// between 3 and 100 characters in length and use alphanumeric letters only.
func (client SystemTopicEventSubscriptionsClient) Get(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string) (result EventSubscription, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SystemTopicEventSubscriptionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, systemTopicName, eventSubscriptionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.SystemTopicEventSubscriptionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.SystemTopicEventSubscriptionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.SystemTopicEventSubscriptionsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client SystemTopicEventSubscriptionsClient) GetPreparer(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"eventSubscriptionName": autorest.Encode("path", eventSubscriptionName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
		"systemTopicName":       autorest.Encode("path", systemTopicName),
	}

	const APIVersion = "2020-10-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics/{systemTopicName}/eventSubscriptions/{eventSubscriptionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SystemTopicEventSubscriptionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SystemTopicEventSubscriptionsClient) GetResponder(resp *http.Response) (result EventSubscription, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetDeliveryAttributes get all delivery attributes for an event subscription.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// systemTopicName - name of the system topic.
// eventSubscriptionName - name of the event subscription to be created. Event subscription names must be
// between 3 and 100 characters in length and use alphanumeric letters only.
func (client SystemTopicEventSubscriptionsClient) GetDeliveryAttributes(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string) (result DeliveryAttributeListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SystemTopicEventSubscriptionsClient.GetDeliveryAttributes")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetDeliveryAttributesPreparer(ctx, resourceGroupName, systemTopicName, eventSubscriptionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.SystemTopicEventSubscriptionsClient", "GetDeliveryAttributes", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetDeliveryAttributesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.SystemTopicEventSubscriptionsClient", "GetDeliveryAttributes", resp, "Failure sending request")
		return
	}

	result, err = client.GetDeliveryAttributesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.SystemTopicEventSubscriptionsClient", "GetDeliveryAttributes", resp, "Failure responding to request")
		return
	}

	return
}

// GetDeliveryAttributesPreparer prepares the GetDeliveryAttributes request.
func (client SystemTopicEventSubscriptionsClient) GetDeliveryAttributesPreparer(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"eventSubscriptionName": autorest.Encode("path", eventSubscriptionName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
		"systemTopicName":       autorest.Encode("path", systemTopicName),
	}

	const APIVersion = "2020-10-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics/{systemTopicName}/eventSubscriptions/{eventSubscriptionName}/getDeliveryAttributes", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetDeliveryAttributesSender sends the GetDeliveryAttributes request. The method will close the
// http.Response Body if it receives an error.
func (client SystemTopicEventSubscriptionsClient) GetDeliveryAttributesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetDeliveryAttributesResponder handles the response to the GetDeliveryAttributes request. The method always
// closes the http.Response Body.
func (client SystemTopicEventSubscriptionsClient) GetDeliveryAttributesResponder(resp *http.Response) (result DeliveryAttributeListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetFullURL get the full endpoint URL for an event subscription of a system topic.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// systemTopicName - name of the system topic.
// eventSubscriptionName - name of the event subscription to be created. Event subscription names must be
// between 3 and 100 characters in length and use alphanumeric letters only.
func (client SystemTopicEventSubscriptionsClient) GetFullURL(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string) (result EventSubscriptionFullURL, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SystemTopicEventSubscriptionsClient.GetFullURL")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetFullURLPreparer(ctx, resourceGroupName, systemTopicName, eventSubscriptionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.SystemTopicEventSubscriptionsClient", "GetFullURL", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetFullURLSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.SystemTopicEventSubscriptionsClient", "GetFullURL", resp, "Failure sending request")
		return
	}

	result, err = client.GetFullURLResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.SystemTopicEventSubscriptionsClient", "GetFullURL", resp, "Failure responding to request")
		return
	}

	return
}

// GetFullURLPreparer prepares the GetFullURL request.
func (client SystemTopicEventSubscriptionsClient) GetFullURLPreparer(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"eventSubscriptionName": autorest.Encode("path", eventSubscriptionName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
		"systemTopicName":       autorest.Encode("path", systemTopicName),
	}

	const APIVersion = "2020-10-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics/{systemTopicName}/eventSubscriptions/{eventSubscriptionName}/getFullUrl", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetFullURLSender sends the GetFullURL request. The method will close the
// http.Response Body if it receives an error.
func (client SystemTopicEventSubscriptionsClient) GetFullURLSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetFullURLResponder handles the response to the GetFullURL request. The method always
// closes the http.Response Body.
func (client SystemTopicEventSubscriptionsClient) GetFullURLResponder(resp *http.Response) (result EventSubscriptionFullURL, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListBySystemTopic list event subscriptions that belong to a specific system topic.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// systemTopicName - name of the system topic.
// filter - the query used to filter the search results using OData syntax. Filtering is permitted on the
// 'name' property only and with limited number of OData operations. These operations are: the 'contains'
// function as well as the following logical operations: not, and, or, eq (for equal), and ne (for not equal).
// No arithmetic operations are supported. The following is a valid filter example: $filter=contains(namE,
// 'PATTERN') and name ne 'PATTERN-1'. The following is not a valid filter example: $filter=location eq
// 'westus'.
// top - the number of results to return per page for the list operation. Valid range for top parameter is 1 to
// 100. If not specified, the default number of results to be returned is 20 items per page.
func (client SystemTopicEventSubscriptionsClient) ListBySystemTopic(ctx context.Context, resourceGroupName string, systemTopicName string, filter string, top *int32) (result EventSubscriptionsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SystemTopicEventSubscriptionsClient.ListBySystemTopic")
		defer func() {
			sc := -1
			if result.eslr.Response.Response != nil {
				sc = result.eslr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listBySystemTopicNextResults
	req, err := client.ListBySystemTopicPreparer(ctx, resourceGroupName, systemTopicName, filter, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.SystemTopicEventSubscriptionsClient", "ListBySystemTopic", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySystemTopicSender(req)
	if err != nil {
		result.eslr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.SystemTopicEventSubscriptionsClient", "ListBySystemTopic", resp, "Failure sending request")
		return
	}

	result.eslr, err = client.ListBySystemTopicResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.SystemTopicEventSubscriptionsClient", "ListBySystemTopic", resp, "Failure responding to request")
		return
	}
	if result.eslr.hasNextLink() && result.eslr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListBySystemTopicPreparer prepares the ListBySystemTopic request.
func (client SystemTopicEventSubscriptionsClient) ListBySystemTopicPreparer(ctx context.Context, resourceGroupName string, systemTopicName string, filter string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"systemTopicName":   autorest.Encode("path", systemTopicName),
	}

	const APIVersion = "2020-10-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics/{systemTopicName}/eventSubscriptions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySystemTopicSender sends the ListBySystemTopic request. The method will close the
// http.Response Body if it receives an error.
func (client SystemTopicEventSubscriptionsClient) ListBySystemTopicSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySystemTopicResponder handles the response to the ListBySystemTopic request. The method always
// closes the http.Response Body.
func (client SystemTopicEventSubscriptionsClient) ListBySystemTopicResponder(resp *http.Response) (result EventSubscriptionsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySystemTopicNextResults retrieves the next set of results, if any.
func (client SystemTopicEventSubscriptionsClient) listBySystemTopicNextResults(ctx context.Context, lastResults EventSubscriptionsListResult) (result EventSubscriptionsListResult, err error) {
	req, err := lastResults.eventSubscriptionsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "eventgrid.SystemTopicEventSubscriptionsClient", "listBySystemTopicNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySystemTopicSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "eventgrid.SystemTopicEventSubscriptionsClient", "listBySystemTopicNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySystemTopicResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.SystemTopicEventSubscriptionsClient", "listBySystemTopicNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySystemTopicComplete enumerates all values, automatically crossing page boundaries as required.
func (client SystemTopicEventSubscriptionsClient) ListBySystemTopicComplete(ctx context.Context, resourceGroupName string, systemTopicName string, filter string, top *int32) (result EventSubscriptionsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SystemTopicEventSubscriptionsClient.ListBySystemTopic")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySystemTopic(ctx, resourceGroupName, systemTopicName, filter, top)
	return
}

// Update update event subscription of a system topic.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// systemTopicName - name of the system topic.
// eventSubscriptionName - name of the event subscription to be created. Event subscription names must be
// between 3 and 100 characters in length and use alphanumeric letters only.
// eventSubscriptionUpdateParameters - updated event subscription information.
func (client SystemTopicEventSubscriptionsClient) Update(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string, eventSubscriptionUpdateParameters EventSubscriptionUpdateParameters) (result SystemTopicEventSubscriptionsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SystemTopicEventSubscriptionsClient.Update")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, systemTopicName, eventSubscriptionName, eventSubscriptionUpdateParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.SystemTopicEventSubscriptionsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.SystemTopicEventSubscriptionsClient", "Update", nil, "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client SystemTopicEventSubscriptionsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string, eventSubscriptionUpdateParameters EventSubscriptionUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"eventSubscriptionName": autorest.Encode("path", eventSubscriptionName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
		"systemTopicName":       autorest.Encode("path", systemTopicName),
	}

	const APIVersion = "2020-10-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics/{systemTopicName}/eventSubscriptions/{eventSubscriptionName}", pathParameters),
		autorest.WithJSON(eventSubscriptionUpdateParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client SystemTopicEventSubscriptionsClient) UpdateSender(req *http.Request) (future SystemTopicEventSubscriptionsUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client SystemTopicEventSubscriptionsClient) UpdateResponder(resp *http.Response) (result EventSubscription, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
