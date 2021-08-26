package network

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

// HubRouteTablesClient is the network Client
type HubRouteTablesClient struct {
	BaseClient
}

// NewHubRouteTablesClient creates an instance of the HubRouteTablesClient client.
func NewHubRouteTablesClient(subscriptionID string) HubRouteTablesClient {
	return NewHubRouteTablesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewHubRouteTablesClientWithBaseURI creates an instance of the HubRouteTablesClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewHubRouteTablesClientWithBaseURI(baseURI string, subscriptionID string) HubRouteTablesClient {
	return HubRouteTablesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a RouteTable resource if it doesn't exist else updates the existing RouteTable.
// Parameters:
// resourceGroupName - the resource group name of the VirtualHub.
// virtualHubName - the name of the VirtualHub.
// routeTableName - the name of the RouteTable.
// routeTableParameters - parameters supplied to create or update RouteTable.
func (client HubRouteTablesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, routeTableParameters HubRouteTable) (result HubRouteTablesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HubRouteTablesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, virtualHubName, routeTableName, routeTableParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.HubRouteTablesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.HubRouteTablesClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client HubRouteTablesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, routeTableParameters HubRouteTable) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"routeTableName":    autorest.Encode("path", routeTableName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"virtualHubName":    autorest.Encode("path", virtualHubName),
	}

	const APIVersion = "2020-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	routeTableParameters.Etag = nil
	routeTableParameters.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubRouteTables/{routeTableName}", pathParameters),
		autorest.WithJSON(routeTableParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client HubRouteTablesClient) CreateOrUpdateSender(req *http.Request) (future HubRouteTablesCreateOrUpdateFuture, err error) {
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
func (client HubRouteTablesClient) CreateOrUpdateResponder(resp *http.Response) (result HubRouteTable, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a RouteTable.
// Parameters:
// resourceGroupName - the resource group name of the RouteTable.
// virtualHubName - the name of the VirtualHub.
// routeTableName - the name of the RouteTable.
func (client HubRouteTablesClient) Delete(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string) (result HubRouteTablesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HubRouteTablesClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, virtualHubName, routeTableName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.HubRouteTablesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.HubRouteTablesClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client HubRouteTablesClient) DeletePreparer(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"routeTableName":    autorest.Encode("path", routeTableName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"virtualHubName":    autorest.Encode("path", virtualHubName),
	}

	const APIVersion = "2020-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubRouteTables/{routeTableName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client HubRouteTablesClient) DeleteSender(req *http.Request) (future HubRouteTablesDeleteFuture, err error) {
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
func (client HubRouteTablesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieves the details of a RouteTable.
// Parameters:
// resourceGroupName - the resource group name of the VirtualHub.
// virtualHubName - the name of the VirtualHub.
// routeTableName - the name of the RouteTable.
func (client HubRouteTablesClient) Get(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string) (result HubRouteTable, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HubRouteTablesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, virtualHubName, routeTableName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.HubRouteTablesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.HubRouteTablesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.HubRouteTablesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client HubRouteTablesClient) GetPreparer(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"routeTableName":    autorest.Encode("path", routeTableName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"virtualHubName":    autorest.Encode("path", virtualHubName),
	}

	const APIVersion = "2020-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubRouteTables/{routeTableName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client HubRouteTablesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client HubRouteTablesClient) GetResponder(resp *http.Response) (result HubRouteTable, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List retrieves the details of all RouteTables.
// Parameters:
// resourceGroupName - the resource group name of the VirtualHub.
// virtualHubName - the name of the VirtualHub.
func (client HubRouteTablesClient) List(ctx context.Context, resourceGroupName string, virtualHubName string) (result ListHubRouteTablesResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HubRouteTablesClient.List")
		defer func() {
			sc := -1
			if result.lhrtr.Response.Response != nil {
				sc = result.lhrtr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, virtualHubName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.HubRouteTablesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.lhrtr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.HubRouteTablesClient", "List", resp, "Failure sending request")
		return
	}

	result.lhrtr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.HubRouteTablesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.lhrtr.hasNextLink() && result.lhrtr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client HubRouteTablesClient) ListPreparer(ctx context.Context, resourceGroupName string, virtualHubName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"virtualHubName":    autorest.Encode("path", virtualHubName),
	}

	const APIVersion = "2020-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubRouteTables", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client HubRouteTablesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client HubRouteTablesClient) ListResponder(resp *http.Response) (result ListHubRouteTablesResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client HubRouteTablesClient) listNextResults(ctx context.Context, lastResults ListHubRouteTablesResult) (result ListHubRouteTablesResult, err error) {
	req, err := lastResults.listHubRouteTablesResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.HubRouteTablesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.HubRouteTablesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.HubRouteTablesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client HubRouteTablesClient) ListComplete(ctx context.Context, resourceGroupName string, virtualHubName string) (result ListHubRouteTablesResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HubRouteTablesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, virtualHubName)
	return
}
