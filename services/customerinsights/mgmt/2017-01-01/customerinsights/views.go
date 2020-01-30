package customerinsights

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

// ViewsClient is the the Azure Customer Insights management API provides a RESTful set of web services that interact
// with Azure Customer Insights service to manage your resources. The API has entities that capture the relationship
// between an end user and the Azure Customer Insights service.
type ViewsClient struct {
	BaseClient
}

// NewViewsClient creates an instance of the ViewsClient client.
func NewViewsClient(subscriptionID string) ViewsClient {
	return NewViewsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewViewsClientWithBaseURI creates an instance of the ViewsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewViewsClientWithBaseURI(baseURI string, subscriptionID string) ViewsClient {
	return ViewsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a view or updates an existing view in the hub.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
// viewName - the name of the view.
// parameters - parameters supplied to the CreateOrUpdate View operation.
func (client ViewsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, hubName string, viewName string, parameters ViewResourceFormat) (result ViewResourceFormat, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ViewsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: viewName,
			Constraints: []validation.Constraint{{Target: "viewName", Name: validation.MaxLength, Rule: 512, Chain: nil},
				{Target: "viewName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.View", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.View.Definition", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("customerinsights.ViewsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, hubName, viewName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ViewsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "customerinsights.ViewsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ViewsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ViewsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, hubName string, viewName string, parameters ViewResourceFormat) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":           autorest.Encode("path", hubName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"viewName":          autorest.Encode("path", viewName),
	}

	const APIVersion = "2017-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/views/{viewName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ViewsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ViewsClient) CreateOrUpdateResponder(resp *http.Response) (result ViewResourceFormat, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a view in the specified hub.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
// viewName - the name of the view.
// userID - the user ID. Use * to retrieve hub level view.
func (client ViewsClient) Delete(ctx context.Context, resourceGroupName string, hubName string, viewName string, userID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ViewsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, hubName, viewName, userID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ViewsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "customerinsights.ViewsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ViewsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ViewsClient) DeletePreparer(ctx context.Context, resourceGroupName string, hubName string, viewName string, userID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":           autorest.Encode("path", hubName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"viewName":          autorest.Encode("path", viewName),
	}

	const APIVersion = "2017-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
		"userId":      autorest.Encode("query", userID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/views/{viewName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ViewsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ViewsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a view in the hub.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
// viewName - the name of the view.
// userID - the user ID. Use * to retrieve hub level view.
func (client ViewsClient) Get(ctx context.Context, resourceGroupName string, hubName string, viewName string, userID string) (result ViewResourceFormat, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ViewsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, hubName, viewName, userID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ViewsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "customerinsights.ViewsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ViewsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ViewsClient) GetPreparer(ctx context.Context, resourceGroupName string, hubName string, viewName string, userID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":           autorest.Encode("path", hubName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"viewName":          autorest.Encode("path", viewName),
	}

	const APIVersion = "2017-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
		"userId":      autorest.Encode("query", userID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/views/{viewName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ViewsClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ViewsClient) GetResponder(resp *http.Response) (result ViewResourceFormat, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByHub gets all available views for given user in the specified hub.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
// userID - the user ID. Use * to retrieve hub level views.
func (client ViewsClient) ListByHub(ctx context.Context, resourceGroupName string, hubName string, userID string) (result ViewListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ViewsClient.ListByHub")
		defer func() {
			sc := -1
			if result.vlr.Response.Response != nil {
				sc = result.vlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByHubNextResults
	req, err := client.ListByHubPreparer(ctx, resourceGroupName, hubName, userID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ViewsClient", "ListByHub", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByHubSender(req)
	if err != nil {
		result.vlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "customerinsights.ViewsClient", "ListByHub", resp, "Failure sending request")
		return
	}

	result.vlr, err = client.ListByHubResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ViewsClient", "ListByHub", resp, "Failure responding to request")
	}

	return
}

// ListByHubPreparer prepares the ListByHub request.
func (client ViewsClient) ListByHubPreparer(ctx context.Context, resourceGroupName string, hubName string, userID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":           autorest.Encode("path", hubName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
		"userId":      autorest.Encode("query", userID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/views", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByHubSender sends the ListByHub request. The method will close the
// http.Response Body if it receives an error.
func (client ViewsClient) ListByHubSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByHubResponder handles the response to the ListByHub request. The method always
// closes the http.Response Body.
func (client ViewsClient) ListByHubResponder(resp *http.Response) (result ViewListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByHubNextResults retrieves the next set of results, if any.
func (client ViewsClient) listByHubNextResults(ctx context.Context, lastResults ViewListResult) (result ViewListResult, err error) {
	req, err := lastResults.viewListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "customerinsights.ViewsClient", "listByHubNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByHubSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "customerinsights.ViewsClient", "listByHubNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByHubResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ViewsClient", "listByHubNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByHubComplete enumerates all values, automatically crossing page boundaries as required.
func (client ViewsClient) ListByHubComplete(ctx context.Context, resourceGroupName string, hubName string, userID string) (result ViewListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ViewsClient.ListByHub")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByHub(ctx, resourceGroupName, hubName, userID)
	return
}
