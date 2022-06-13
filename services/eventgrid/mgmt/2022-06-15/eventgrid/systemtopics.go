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

// SystemTopicsClient is the azure EventGrid Management Client
type SystemTopicsClient struct {
	BaseClient
}

// NewSystemTopicsClient creates an instance of the SystemTopicsClient client.
func NewSystemTopicsClient(subscriptionID string) SystemTopicsClient {
	return NewSystemTopicsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSystemTopicsClientWithBaseURI creates an instance of the SystemTopicsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewSystemTopicsClientWithBaseURI(baseURI string, subscriptionID string) SystemTopicsClient {
	return SystemTopicsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate asynchronously creates a new system topic with the specified parameters.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// systemTopicName - name of the system topic.
// systemTopicInfo - system Topic information.
func (client SystemTopicsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, systemTopicName string, systemTopicInfo SystemTopic) (result SystemTopicsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SystemTopicsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, systemTopicName, systemTopicInfo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.SystemTopicsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.SystemTopicsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SystemTopicsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, systemTopicName string, systemTopicInfo SystemTopic) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"systemTopicName":   autorest.Encode("path", systemTopicName),
	}

	const APIVersion = "2022-06-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	systemTopicInfo.SystemData = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics/{systemTopicName}", pathParameters),
		autorest.WithJSON(systemTopicInfo),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SystemTopicsClient) CreateOrUpdateSender(req *http.Request) (future SystemTopicsCreateOrUpdateFuture, err error) {
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
func (client SystemTopicsClient) CreateOrUpdateResponder(resp *http.Response) (result SystemTopic, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete existing system topic.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// systemTopicName - name of the system topic.
func (client SystemTopicsClient) Delete(ctx context.Context, resourceGroupName string, systemTopicName string) (result SystemTopicsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SystemTopicsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, systemTopicName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.SystemTopicsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.SystemTopicsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SystemTopicsClient) DeletePreparer(ctx context.Context, resourceGroupName string, systemTopicName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"systemTopicName":   autorest.Encode("path", systemTopicName),
	}

	const APIVersion = "2022-06-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics/{systemTopicName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SystemTopicsClient) DeleteSender(req *http.Request) (future SystemTopicsDeleteFuture, err error) {
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
func (client SystemTopicsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get properties of a system topic.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// systemTopicName - name of the system topic.
func (client SystemTopicsClient) Get(ctx context.Context, resourceGroupName string, systemTopicName string) (result SystemTopic, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SystemTopicsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, systemTopicName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.SystemTopicsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.SystemTopicsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.SystemTopicsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client SystemTopicsClient) GetPreparer(ctx context.Context, resourceGroupName string, systemTopicName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"systemTopicName":   autorest.Encode("path", systemTopicName),
	}

	const APIVersion = "2022-06-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics/{systemTopicName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SystemTopicsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SystemTopicsClient) GetResponder(resp *http.Response) (result SystemTopic, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup list all the system topics under a resource group.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// filter - the query used to filter the search results using OData syntax. Filtering is permitted on the
// 'name' property only and with limited number of OData operations. These operations are: the 'contains'
// function as well as the following logical operations: not, and, or, eq (for equal), and ne (for not equal).
// No arithmetic operations are supported. The following is a valid filter example: $filter=contains(namE,
// 'PATTERN') and name ne 'PATTERN-1'. The following is not a valid filter example: $filter=location eq
// 'westus'.
// top - the number of results to return per page for the list operation. Valid range for top parameter is 1 to
// 100. If not specified, the default number of results to be returned is 20 items per page.
func (client SystemTopicsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, top *int32) (result SystemTopicsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SystemTopicsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.stlr.Response.Response != nil {
				sc = result.stlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, filter, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.SystemTopicsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.stlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.SystemTopicsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.stlr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.SystemTopicsClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.stlr.hasNextLink() && result.stlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client SystemTopicsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, filter string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-06-15"
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client SystemTopicsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client SystemTopicsClient) ListByResourceGroupResponder(resp *http.Response) (result SystemTopicsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client SystemTopicsClient) listByResourceGroupNextResults(ctx context.Context, lastResults SystemTopicsListResult) (result SystemTopicsListResult, err error) {
	req, err := lastResults.systemTopicsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "eventgrid.SystemTopicsClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "eventgrid.SystemTopicsClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.SystemTopicsClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client SystemTopicsClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string, top *int32) (result SystemTopicsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SystemTopicsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName, filter, top)
	return
}

// ListBySubscription list all the system topics under an Azure subscription.
// Parameters:
// filter - the query used to filter the search results using OData syntax. Filtering is permitted on the
// 'name' property only and with limited number of OData operations. These operations are: the 'contains'
// function as well as the following logical operations: not, and, or, eq (for equal), and ne (for not equal).
// No arithmetic operations are supported. The following is a valid filter example: $filter=contains(namE,
// 'PATTERN') and name ne 'PATTERN-1'. The following is not a valid filter example: $filter=location eq
// 'westus'.
// top - the number of results to return per page for the list operation. Valid range for top parameter is 1 to
// 100. If not specified, the default number of results to be returned is 20 items per page.
func (client SystemTopicsClient) ListBySubscription(ctx context.Context, filter string, top *int32) (result SystemTopicsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SystemTopicsClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.stlr.Response.Response != nil {
				sc = result.stlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx, filter, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.SystemTopicsClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.stlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.SystemTopicsClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.stlr, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.SystemTopicsClient", "ListBySubscription", resp, "Failure responding to request")
		return
	}
	if result.stlr.hasNextLink() && result.stlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client SystemTopicsClient) ListBySubscriptionPreparer(ctx context.Context, filter string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-06-15"
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.EventGrid/systemTopics", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client SystemTopicsClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client SystemTopicsClient) ListBySubscriptionResponder(resp *http.Response) (result SystemTopicsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionNextResults retrieves the next set of results, if any.
func (client SystemTopicsClient) listBySubscriptionNextResults(ctx context.Context, lastResults SystemTopicsListResult) (result SystemTopicsListResult, err error) {
	req, err := lastResults.systemTopicsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "eventgrid.SystemTopicsClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "eventgrid.SystemTopicsClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.SystemTopicsClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client SystemTopicsClient) ListBySubscriptionComplete(ctx context.Context, filter string, top *int32) (result SystemTopicsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SystemTopicsClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySubscription(ctx, filter, top)
	return
}

// Update asynchronously updates a system topic with the specified parameters.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// systemTopicName - name of the system topic.
// systemTopicUpdateParameters - systemTopic update information.
func (client SystemTopicsClient) Update(ctx context.Context, resourceGroupName string, systemTopicName string, systemTopicUpdateParameters SystemTopicUpdateParameters) (result SystemTopicsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SystemTopicsClient.Update")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, systemTopicName, systemTopicUpdateParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.SystemTopicsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.SystemTopicsClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client SystemTopicsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, systemTopicName string, systemTopicUpdateParameters SystemTopicUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"systemTopicName":   autorest.Encode("path", systemTopicName),
	}

	const APIVersion = "2022-06-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics/{systemTopicName}", pathParameters),
		autorest.WithJSON(systemTopicUpdateParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client SystemTopicsClient) UpdateSender(req *http.Request) (future SystemTopicsUpdateFuture, err error) {
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
func (client SystemTopicsClient) UpdateResponder(resp *http.Response) (result SystemTopic, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
