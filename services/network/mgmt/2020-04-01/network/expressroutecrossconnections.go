package network

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

// ExpressRouteCrossConnectionsClient is the network Client
type ExpressRouteCrossConnectionsClient struct {
	BaseClient
}

// NewExpressRouteCrossConnectionsClient creates an instance of the ExpressRouteCrossConnectionsClient client.
func NewExpressRouteCrossConnectionsClient(subscriptionID string) ExpressRouteCrossConnectionsClient {
	return NewExpressRouteCrossConnectionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewExpressRouteCrossConnectionsClientWithBaseURI creates an instance of the ExpressRouteCrossConnectionsClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewExpressRouteCrossConnectionsClientWithBaseURI(baseURI string, subscriptionID string) ExpressRouteCrossConnectionsClient {
	return ExpressRouteCrossConnectionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate update the specified ExpressRouteCrossConnection.
// Parameters:
// resourceGroupName - the name of the resource group.
// crossConnectionName - the name of the ExpressRouteCrossConnection.
// parameters - parameters supplied to the update express route crossConnection operation.
func (client ExpressRouteCrossConnectionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, crossConnectionName string, parameters ExpressRouteCrossConnection) (result ExpressRouteCrossConnectionsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteCrossConnectionsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, crossConnectionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ExpressRouteCrossConnectionsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, crossConnectionName string, parameters ExpressRouteCrossConnection) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"crossConnectionName": autorest.Encode("path", crossConnectionName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.Etag = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCrossConnectionsClient) CreateOrUpdateSender(req *http.Request) (future ExpressRouteCrossConnectionsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ExpressRouteCrossConnectionsClient) CreateOrUpdateResponder(resp *http.Response) (result ExpressRouteCrossConnection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets details about the specified ExpressRouteCrossConnection.
// Parameters:
// resourceGroupName - the name of the resource group (peering location of the circuit).
// crossConnectionName - the name of the ExpressRouteCrossConnection (service key of the circuit).
func (client ExpressRouteCrossConnectionsClient) Get(ctx context.Context, resourceGroupName string, crossConnectionName string) (result ExpressRouteCrossConnection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteCrossConnectionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, crossConnectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ExpressRouteCrossConnectionsClient) GetPreparer(ctx context.Context, resourceGroupName string, crossConnectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"crossConnectionName": autorest.Encode("path", crossConnectionName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCrossConnectionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ExpressRouteCrossConnectionsClient) GetResponder(resp *http.Response) (result ExpressRouteCrossConnection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List retrieves all the ExpressRouteCrossConnections in a subscription.
func (client ExpressRouteCrossConnectionsClient) List(ctx context.Context) (result ExpressRouteCrossConnectionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteCrossConnectionsClient.List")
		defer func() {
			sc := -1
			if result.ercclr.Response.Response != nil {
				sc = result.ercclr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.ercclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionsClient", "List", resp, "Failure sending request")
		return
	}

	result.ercclr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ExpressRouteCrossConnectionsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteCrossConnections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCrossConnectionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ExpressRouteCrossConnectionsClient) ListResponder(resp *http.Response) (result ExpressRouteCrossConnectionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ExpressRouteCrossConnectionsClient) listNextResults(ctx context.Context, lastResults ExpressRouteCrossConnectionListResult) (result ExpressRouteCrossConnectionListResult, err error) {
	req, err := lastResults.expressRouteCrossConnectionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ExpressRouteCrossConnectionsClient) ListComplete(ctx context.Context) (result ExpressRouteCrossConnectionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteCrossConnectionsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}

// ListArpTable gets the currently advertised ARP table associated with the express route cross connection in a
// resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// crossConnectionName - the name of the ExpressRouteCrossConnection.
// peeringName - the name of the peering.
// devicePath - the path of the device.
func (client ExpressRouteCrossConnectionsClient) ListArpTable(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string) (result ExpressRouteCrossConnectionsListArpTableFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteCrossConnectionsClient.ListArpTable")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListArpTablePreparer(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionsClient", "ListArpTable", nil, "Failure preparing request")
		return
	}

	result, err = client.ListArpTableSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionsClient", "ListArpTable", result.Response(), "Failure sending request")
		return
	}

	return
}

// ListArpTablePreparer prepares the ListArpTable request.
func (client ExpressRouteCrossConnectionsClient) ListArpTablePreparer(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"crossConnectionName": autorest.Encode("path", crossConnectionName),
		"devicePath":          autorest.Encode("path", devicePath),
		"peeringName":         autorest.Encode("path", peeringName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}/arpTables/{devicePath}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListArpTableSender sends the ListArpTable request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCrossConnectionsClient) ListArpTableSender(req *http.Request) (future ExpressRouteCrossConnectionsListArpTableFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// ListArpTableResponder handles the response to the ListArpTable request. The method always
// closes the http.Response Body.
func (client ExpressRouteCrossConnectionsClient) ListArpTableResponder(resp *http.Response) (result ExpressRouteCircuitsArpTableListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup retrieves all the ExpressRouteCrossConnections in a resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
func (client ExpressRouteCrossConnectionsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result ExpressRouteCrossConnectionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteCrossConnectionsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.ercclr.Response.Response != nil {
				sc = result.ercclr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.ercclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.ercclr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionsClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ExpressRouteCrossConnectionsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCrossConnectionsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ExpressRouteCrossConnectionsClient) ListByResourceGroupResponder(resp *http.Response) (result ExpressRouteCrossConnectionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client ExpressRouteCrossConnectionsClient) listByResourceGroupNextResults(ctx context.Context, lastResults ExpressRouteCrossConnectionListResult) (result ExpressRouteCrossConnectionListResult, err error) {
	req, err := lastResults.expressRouteCrossConnectionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionsClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionsClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionsClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client ExpressRouteCrossConnectionsClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result ExpressRouteCrossConnectionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteCrossConnectionsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// ListRoutesTable gets the currently advertised routes table associated with the express route cross connection in a
// resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// crossConnectionName - the name of the ExpressRouteCrossConnection.
// peeringName - the name of the peering.
// devicePath - the path of the device.
func (client ExpressRouteCrossConnectionsClient) ListRoutesTable(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string) (result ExpressRouteCrossConnectionsListRoutesTableFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteCrossConnectionsClient.ListRoutesTable")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListRoutesTablePreparer(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionsClient", "ListRoutesTable", nil, "Failure preparing request")
		return
	}

	result, err = client.ListRoutesTableSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionsClient", "ListRoutesTable", result.Response(), "Failure sending request")
		return
	}

	return
}

// ListRoutesTablePreparer prepares the ListRoutesTable request.
func (client ExpressRouteCrossConnectionsClient) ListRoutesTablePreparer(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"crossConnectionName": autorest.Encode("path", crossConnectionName),
		"devicePath":          autorest.Encode("path", devicePath),
		"peeringName":         autorest.Encode("path", peeringName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}/routeTables/{devicePath}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListRoutesTableSender sends the ListRoutesTable request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCrossConnectionsClient) ListRoutesTableSender(req *http.Request) (future ExpressRouteCrossConnectionsListRoutesTableFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// ListRoutesTableResponder handles the response to the ListRoutesTable request. The method always
// closes the http.Response Body.
func (client ExpressRouteCrossConnectionsClient) ListRoutesTableResponder(resp *http.Response) (result ExpressRouteCircuitsRoutesTableListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListRoutesTableSummary gets the route table summary associated with the express route cross connection in a resource
// group.
// Parameters:
// resourceGroupName - the name of the resource group.
// crossConnectionName - the name of the ExpressRouteCrossConnection.
// peeringName - the name of the peering.
// devicePath - the path of the device.
func (client ExpressRouteCrossConnectionsClient) ListRoutesTableSummary(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string) (result ExpressRouteCrossConnectionsListRoutesTableSummaryFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteCrossConnectionsClient.ListRoutesTableSummary")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListRoutesTableSummaryPreparer(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionsClient", "ListRoutesTableSummary", nil, "Failure preparing request")
		return
	}

	result, err = client.ListRoutesTableSummarySender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionsClient", "ListRoutesTableSummary", result.Response(), "Failure sending request")
		return
	}

	return
}

// ListRoutesTableSummaryPreparer prepares the ListRoutesTableSummary request.
func (client ExpressRouteCrossConnectionsClient) ListRoutesTableSummaryPreparer(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"crossConnectionName": autorest.Encode("path", crossConnectionName),
		"devicePath":          autorest.Encode("path", devicePath),
		"peeringName":         autorest.Encode("path", peeringName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}/routeTablesSummary/{devicePath}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListRoutesTableSummarySender sends the ListRoutesTableSummary request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCrossConnectionsClient) ListRoutesTableSummarySender(req *http.Request) (future ExpressRouteCrossConnectionsListRoutesTableSummaryFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// ListRoutesTableSummaryResponder handles the response to the ListRoutesTableSummary request. The method always
// closes the http.Response Body.
func (client ExpressRouteCrossConnectionsClient) ListRoutesTableSummaryResponder(resp *http.Response) (result ExpressRouteCrossConnectionsRoutesTableSummaryListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UpdateTags updates an express route cross connection tags.
// Parameters:
// resourceGroupName - the name of the resource group.
// crossConnectionName - the name of the cross connection.
// crossConnectionParameters - parameters supplied to update express route cross connection tags.
func (client ExpressRouteCrossConnectionsClient) UpdateTags(ctx context.Context, resourceGroupName string, crossConnectionName string, crossConnectionParameters TagsObject) (result ExpressRouteCrossConnection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteCrossConnectionsClient.UpdateTags")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdateTagsPreparer(ctx, resourceGroupName, crossConnectionName, crossConnectionParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionsClient", "UpdateTags", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateTagsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionsClient", "UpdateTags", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateTagsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionsClient", "UpdateTags", resp, "Failure responding to request")
	}

	return
}

// UpdateTagsPreparer prepares the UpdateTags request.
func (client ExpressRouteCrossConnectionsClient) UpdateTagsPreparer(ctx context.Context, resourceGroupName string, crossConnectionName string, crossConnectionParameters TagsObject) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"crossConnectionName": autorest.Encode("path", crossConnectionName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}", pathParameters),
		autorest.WithJSON(crossConnectionParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateTagsSender sends the UpdateTags request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCrossConnectionsClient) UpdateTagsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateTagsResponder handles the response to the UpdateTags request. The method always
// closes the http.Response Body.
func (client ExpressRouteCrossConnectionsClient) UpdateTagsResponder(resp *http.Response) (result ExpressRouteCrossConnection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
