package webpubsub

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

// PrivateEndpointConnectionsClient is the REST API for Azure WebPubSub Service
type PrivateEndpointConnectionsClient struct {
	BaseClient
}

// NewPrivateEndpointConnectionsClient creates an instance of the PrivateEndpointConnectionsClient client.
func NewPrivateEndpointConnectionsClient(subscriptionID string) PrivateEndpointConnectionsClient {
	return NewPrivateEndpointConnectionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPrivateEndpointConnectionsClientWithBaseURI creates an instance of the PrivateEndpointConnectionsClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewPrivateEndpointConnectionsClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointConnectionsClient {
	return PrivateEndpointConnectionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Delete delete the specified private endpoint connection
// Parameters:
// privateEndpointConnectionName - the name of the private endpoint connection
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// resourceName - the name of the resource.
func (client PrivateEndpointConnectionsClient) Delete(ctx context.Context, privateEndpointConnectionName string, resourceGroupName string, resourceName string) (result PrivateEndpointConnectionsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndpointConnectionsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, privateEndpointConnectionName, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webpubsub.PrivateEndpointConnectionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webpubsub.PrivateEndpointConnectionsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client PrivateEndpointConnectionsClient) DeletePreparer(ctx context.Context, privateEndpointConnectionName string, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateEndpointConnectionName": autorest.Encode("path", privateEndpointConnectionName),
		"resourceGroupName":             autorest.Encode("path", resourceGroupName),
		"resourceName":                  autorest.Encode("path", resourceName),
		"subscriptionId":                autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/webPubSub/{resourceName}/privateEndpointConnections/{privateEndpointConnectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateEndpointConnectionsClient) DeleteSender(req *http.Request) (future PrivateEndpointConnectionsDeleteFuture, err error) {
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
func (client PrivateEndpointConnectionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get the specified private endpoint connection
// Parameters:
// privateEndpointConnectionName - the name of the private endpoint connection
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// resourceName - the name of the resource.
func (client PrivateEndpointConnectionsClient) Get(ctx context.Context, privateEndpointConnectionName string, resourceGroupName string, resourceName string) (result PrivateEndpointConnection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndpointConnectionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, privateEndpointConnectionName, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webpubsub.PrivateEndpointConnectionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "webpubsub.PrivateEndpointConnectionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webpubsub.PrivateEndpointConnectionsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client PrivateEndpointConnectionsClient) GetPreparer(ctx context.Context, privateEndpointConnectionName string, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateEndpointConnectionName": autorest.Encode("path", privateEndpointConnectionName),
		"resourceGroupName":             autorest.Encode("path", resourceGroupName),
		"resourceName":                  autorest.Encode("path", resourceName),
		"subscriptionId":                autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/webPubSub/{resourceName}/privateEndpointConnections/{privateEndpointConnectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateEndpointConnectionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PrivateEndpointConnectionsClient) GetResponder(resp *http.Response) (result PrivateEndpointConnection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list private endpoint connections
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// resourceName - the name of the resource.
func (client PrivateEndpointConnectionsClient) List(ctx context.Context, resourceGroupName string, resourceName string) (result PrivateEndpointConnectionListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndpointConnectionsClient.List")
		defer func() {
			sc := -1
			if result.pecl.Response.Response != nil {
				sc = result.pecl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webpubsub.PrivateEndpointConnectionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.pecl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "webpubsub.PrivateEndpointConnectionsClient", "List", resp, "Failure sending request")
		return
	}

	result.pecl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webpubsub.PrivateEndpointConnectionsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.pecl.hasNextLink() && result.pecl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client PrivateEndpointConnectionsClient) ListPreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/webPubSub/{resourceName}/privateEndpointConnections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateEndpointConnectionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client PrivateEndpointConnectionsClient) ListResponder(resp *http.Response) (result PrivateEndpointConnectionList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client PrivateEndpointConnectionsClient) listNextResults(ctx context.Context, lastResults PrivateEndpointConnectionList) (result PrivateEndpointConnectionList, err error) {
	req, err := lastResults.privateEndpointConnectionListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "webpubsub.PrivateEndpointConnectionsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "webpubsub.PrivateEndpointConnectionsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webpubsub.PrivateEndpointConnectionsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client PrivateEndpointConnectionsClient) ListComplete(ctx context.Context, resourceGroupName string, resourceName string) (result PrivateEndpointConnectionListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndpointConnectionsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, resourceName)
	return
}

// Update update the state of specified private endpoint connection
// Parameters:
// privateEndpointConnectionName - the name of the private endpoint connection
// parameters - the resource of private endpoint and its properties
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// resourceName - the name of the resource.
func (client PrivateEndpointConnectionsClient) Update(ctx context.Context, privateEndpointConnectionName string, parameters PrivateEndpointConnection, resourceGroupName string, resourceName string) (result PrivateEndpointConnection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndpointConnectionsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, privateEndpointConnectionName, parameters, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webpubsub.PrivateEndpointConnectionsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "webpubsub.PrivateEndpointConnectionsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webpubsub.PrivateEndpointConnectionsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client PrivateEndpointConnectionsClient) UpdatePreparer(ctx context.Context, privateEndpointConnectionName string, parameters PrivateEndpointConnection, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateEndpointConnectionName": autorest.Encode("path", privateEndpointConnectionName),
		"resourceGroupName":             autorest.Encode("path", resourceGroupName),
		"resourceName":                  autorest.Encode("path", resourceName),
		"subscriptionId":                autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.SystemData = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/webPubSub/{resourceName}/privateEndpointConnections/{privateEndpointConnectionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateEndpointConnectionsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client PrivateEndpointConnectionsClient) UpdateResponder(resp *http.Response) (result PrivateEndpointConnection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
