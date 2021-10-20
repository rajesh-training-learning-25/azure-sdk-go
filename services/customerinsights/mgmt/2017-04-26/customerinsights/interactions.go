package customerinsights

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
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

// InteractionsClient is the the Azure Customer Insights management API provides a RESTful set of web services that
// interact with Azure Customer Insights service to manage your resources. The API has entities that capture the
// relationship between an end user and the Azure Customer Insights service.
type InteractionsClient struct {
	BaseClient
}

// NewInteractionsClient creates an instance of the InteractionsClient client.
func NewInteractionsClient(subscriptionID string) InteractionsClient {
	return NewInteractionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewInteractionsClientWithBaseURI creates an instance of the InteractionsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewInteractionsClientWithBaseURI(baseURI string, subscriptionID string) InteractionsClient {
	return InteractionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates an interaction or updates an existing interaction within a hub.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
// interactionName - the name of the interaction.
// parameters - parameters supplied to the CreateOrUpdate Interaction operation.
func (client InteractionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, hubName string, interactionName string, parameters InteractionResourceFormat) (result InteractionsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InteractionsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: interactionName,
			Constraints: []validation.Constraint{{Target: "interactionName", Name: validation.MaxLength, Rule: 128, Chain: nil},
				{Target: "interactionName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "interactionName", Name: validation.Pattern, Rule: `^[a-zA-Z][a-zA-Z0-9_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("customerinsights.InteractionsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, hubName, interactionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.InteractionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.InteractionsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client InteractionsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, hubName string, interactionName string, parameters InteractionResourceFormat) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":           autorest.Encode("path", hubName),
		"interactionName":   autorest.Encode("path", interactionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-26"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/interactions/{interactionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client InteractionsClient) CreateOrUpdateSender(req *http.Request) (future InteractionsCreateOrUpdateFuture, err error) {
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
func (client InteractionsClient) CreateOrUpdateResponder(resp *http.Response) (result InteractionResourceFormat, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets information about the specified interaction.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
// interactionName - the name of the interaction.
// localeCode - locale of interaction to retrieve, default is en-us.
func (client InteractionsClient) Get(ctx context.Context, resourceGroupName string, hubName string, interactionName string, localeCode string) (result InteractionResourceFormat, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InteractionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, hubName, interactionName, localeCode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.InteractionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "customerinsights.InteractionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.InteractionsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client InteractionsClient) GetPreparer(ctx context.Context, resourceGroupName string, hubName string, interactionName string, localeCode string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":           autorest.Encode("path", hubName),
		"interactionName":   autorest.Encode("path", interactionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-26"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(localeCode) > 0 {
		queryParameters["locale-code"] = autorest.Encode("query", localeCode)
	} else {
		queryParameters["locale-code"] = autorest.Encode("query", "en-us")
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/interactions/{interactionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client InteractionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client InteractionsClient) GetResponder(resp *http.Response) (result InteractionResourceFormat, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByHub gets all interactions in the hub.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
// localeCode - locale of interaction to retrieve, default is en-us.
func (client InteractionsClient) ListByHub(ctx context.Context, resourceGroupName string, hubName string, localeCode string) (result InteractionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InteractionsClient.ListByHub")
		defer func() {
			sc := -1
			if result.ilr.Response.Response != nil {
				sc = result.ilr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByHubNextResults
	req, err := client.ListByHubPreparer(ctx, resourceGroupName, hubName, localeCode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.InteractionsClient", "ListByHub", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByHubSender(req)
	if err != nil {
		result.ilr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "customerinsights.InteractionsClient", "ListByHub", resp, "Failure sending request")
		return
	}

	result.ilr, err = client.ListByHubResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.InteractionsClient", "ListByHub", resp, "Failure responding to request")
		return
	}
	if result.ilr.hasNextLink() && result.ilr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByHubPreparer prepares the ListByHub request.
func (client InteractionsClient) ListByHubPreparer(ctx context.Context, resourceGroupName string, hubName string, localeCode string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":           autorest.Encode("path", hubName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-26"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(localeCode) > 0 {
		queryParameters["locale-code"] = autorest.Encode("query", localeCode)
	} else {
		queryParameters["locale-code"] = autorest.Encode("query", "en-us")
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/interactions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByHubSender sends the ListByHub request. The method will close the
// http.Response Body if it receives an error.
func (client InteractionsClient) ListByHubSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByHubResponder handles the response to the ListByHub request. The method always
// closes the http.Response Body.
func (client InteractionsClient) ListByHubResponder(resp *http.Response) (result InteractionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByHubNextResults retrieves the next set of results, if any.
func (client InteractionsClient) listByHubNextResults(ctx context.Context, lastResults InteractionListResult) (result InteractionListResult, err error) {
	req, err := lastResults.interactionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "customerinsights.InteractionsClient", "listByHubNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByHubSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "customerinsights.InteractionsClient", "listByHubNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByHubResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.InteractionsClient", "listByHubNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByHubComplete enumerates all values, automatically crossing page boundaries as required.
func (client InteractionsClient) ListByHubComplete(ctx context.Context, resourceGroupName string, hubName string, localeCode string) (result InteractionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InteractionsClient.ListByHub")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByHub(ctx, resourceGroupName, hubName, localeCode)
	return
}

// SuggestRelationshipLinks suggests relationships to create relationship links.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
// interactionName - the name of the interaction.
func (client InteractionsClient) SuggestRelationshipLinks(ctx context.Context, resourceGroupName string, hubName string, interactionName string) (result SuggestRelationshipLinksResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InteractionsClient.SuggestRelationshipLinks")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.SuggestRelationshipLinksPreparer(ctx, resourceGroupName, hubName, interactionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.InteractionsClient", "SuggestRelationshipLinks", nil, "Failure preparing request")
		return
	}

	resp, err := client.SuggestRelationshipLinksSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "customerinsights.InteractionsClient", "SuggestRelationshipLinks", resp, "Failure sending request")
		return
	}

	result, err = client.SuggestRelationshipLinksResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.InteractionsClient", "SuggestRelationshipLinks", resp, "Failure responding to request")
		return
	}

	return
}

// SuggestRelationshipLinksPreparer prepares the SuggestRelationshipLinks request.
func (client InteractionsClient) SuggestRelationshipLinksPreparer(ctx context.Context, resourceGroupName string, hubName string, interactionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":           autorest.Encode("path", hubName),
		"interactionName":   autorest.Encode("path", interactionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-26"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/interactions/{interactionName}/suggestRelationshipLinks", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SuggestRelationshipLinksSender sends the SuggestRelationshipLinks request. The method will close the
// http.Response Body if it receives an error.
func (client InteractionsClient) SuggestRelationshipLinksSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// SuggestRelationshipLinksResponder handles the response to the SuggestRelationshipLinks request. The method always
// closes the http.Response Body.
func (client InteractionsClient) SuggestRelationshipLinksResponder(resp *http.Response) (result SuggestRelationshipLinksResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
