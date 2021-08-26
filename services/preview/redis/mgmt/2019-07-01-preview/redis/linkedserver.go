package redis

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

// LinkedServerClient is the REST API for Azure Redis Cache Service.
type LinkedServerClient struct {
	BaseClient
}

// NewLinkedServerClient creates an instance of the LinkedServerClient client.
func NewLinkedServerClient(subscriptionID string) LinkedServerClient {
	return NewLinkedServerClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLinkedServerClientWithBaseURI creates an instance of the LinkedServerClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewLinkedServerClientWithBaseURI(baseURI string, subscriptionID string) LinkedServerClient {
	return LinkedServerClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create adds a linked server to the Redis cache (requires Premium SKU).
// Parameters:
// resourceGroupName - the name of the resource group.
// name - the name of the Redis cache.
// linkedServerName - the name of the linked server that is being added to the Redis cache.
// parameters - parameters supplied to the Create Linked server operation.
func (client LinkedServerClient) Create(ctx context.Context, resourceGroupName string, name string, linkedServerName string, parameters LinkedServerCreateParameters) (result LinkedServerCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LinkedServerClient.Create")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.LinkedServerCreateProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.LinkedServerCreateProperties.LinkedRedisCacheID", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.LinkedServerCreateProperties.LinkedRedisCacheLocation", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("redis.LinkedServerClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, name, linkedServerName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.LinkedServerClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.LinkedServerClient", "Create", nil, "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client LinkedServerClient) CreatePreparer(ctx context.Context, resourceGroupName string, name string, linkedServerName string, parameters LinkedServerCreateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"linkedServerName":  autorest.Encode("path", linkedServerName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/linkedServers/{linkedServerName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client LinkedServerClient) CreateSender(req *http.Request) (future LinkedServerCreateFuture, err error) {
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

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client LinkedServerClient) CreateResponder(resp *http.Response) (result LinkedServerWithProperties, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the linked server from a redis cache (requires Premium SKU).
// Parameters:
// resourceGroupName - the name of the resource group.
// name - the name of the redis cache.
// linkedServerName - the name of the linked server that is being added to the Redis cache.
func (client LinkedServerClient) Delete(ctx context.Context, resourceGroupName string, name string, linkedServerName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LinkedServerClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, name, linkedServerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.LinkedServerClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "redis.LinkedServerClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.LinkedServerClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client LinkedServerClient) DeletePreparer(ctx context.Context, resourceGroupName string, name string, linkedServerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"linkedServerName":  autorest.Encode("path", linkedServerName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/linkedServers/{linkedServerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client LinkedServerClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client LinkedServerClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the detailed information about a linked server of a redis cache (requires Premium SKU).
// Parameters:
// resourceGroupName - the name of the resource group.
// name - the name of the redis cache.
// linkedServerName - the name of the linked server.
func (client LinkedServerClient) Get(ctx context.Context, resourceGroupName string, name string, linkedServerName string) (result LinkedServerWithProperties, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LinkedServerClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, name, linkedServerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.LinkedServerClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "redis.LinkedServerClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.LinkedServerClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client LinkedServerClient) GetPreparer(ctx context.Context, resourceGroupName string, name string, linkedServerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"linkedServerName":  autorest.Encode("path", linkedServerName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/linkedServers/{linkedServerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client LinkedServerClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client LinkedServerClient) GetResponder(resp *http.Response) (result LinkedServerWithProperties, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets the list of linked servers associated with this redis cache (requires Premium SKU).
// Parameters:
// resourceGroupName - the name of the resource group.
// name - the name of the redis cache.
func (client LinkedServerClient) List(ctx context.Context, resourceGroupName string, name string) (result LinkedServerWithPropertiesListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LinkedServerClient.List")
		defer func() {
			sc := -1
			if result.lswpl.Response.Response != nil {
				sc = result.lswpl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.LinkedServerClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.lswpl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "redis.LinkedServerClient", "List", resp, "Failure sending request")
		return
	}

	result.lswpl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.LinkedServerClient", "List", resp, "Failure responding to request")
		return
	}
	if result.lswpl.hasNextLink() && result.lswpl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client LinkedServerClient) ListPreparer(ctx context.Context, resourceGroupName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/linkedServers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client LinkedServerClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client LinkedServerClient) ListResponder(resp *http.Response) (result LinkedServerWithPropertiesList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client LinkedServerClient) listNextResults(ctx context.Context, lastResults LinkedServerWithPropertiesList) (result LinkedServerWithPropertiesList, err error) {
	req, err := lastResults.linkedServerWithPropertiesListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "redis.LinkedServerClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "redis.LinkedServerClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.LinkedServerClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client LinkedServerClient) ListComplete(ctx context.Context, resourceGroupName string, name string) (result LinkedServerWithPropertiesListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LinkedServerClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, name)
	return
}
