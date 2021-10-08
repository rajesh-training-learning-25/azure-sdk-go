package documentdb

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

// CassandraDataCentersClient is the client for the CassandraDataCenters methods of the Documentdb service.
type CassandraDataCentersClient struct {
	BaseClient
}

// NewCassandraDataCentersClient creates an instance of the CassandraDataCentersClient client.
func NewCassandraDataCentersClient(subscriptionID string) CassandraDataCentersClient {
	return NewCassandraDataCentersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCassandraDataCentersClientWithBaseURI creates an instance of the CassandraDataCentersClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewCassandraDataCentersClientWithBaseURI(baseURI string, subscriptionID string) CassandraDataCentersClient {
	return CassandraDataCentersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateUpdate create or update a managed Cassandra data center. When updating, overwrite all properties. To update
// only some properties, use PATCH.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// clusterName - managed Cassandra cluster name.
// dataCenterName - data center name in a managed Cassandra cluster.
// body - parameters specifying the managed Cassandra data center.
func (client CassandraDataCentersClient) CreateUpdate(ctx context.Context, resourceGroupName string, clusterName string, dataCenterName string, body DataCenterResource) (result CassandraDataCentersCreateUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CassandraDataCentersClient.CreateUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: clusterName,
			Constraints: []validation.Constraint{{Target: "clusterName", Name: validation.MaxLength, Rule: 100, Chain: nil},
				{Target: "clusterName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "clusterName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+(-[a-zA-Z0-9]+)*`, Chain: nil}}},
		{TargetValue: dataCenterName,
			Constraints: []validation.Constraint{{Target: "dataCenterName", Name: validation.MaxLength, Rule: 100, Chain: nil},
				{Target: "dataCenterName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "dataCenterName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+(-[a-zA-Z0-9]+)*`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("documentdb.CassandraDataCentersClient", "CreateUpdate", err.Error())
	}

	req, err := client.CreateUpdatePreparer(ctx, resourceGroupName, clusterName, dataCenterName, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.CassandraDataCentersClient", "CreateUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.CassandraDataCentersClient", "CreateUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateUpdatePreparer prepares the CreateUpdate request.
func (client CassandraDataCentersClient) CreateUpdatePreparer(ctx context.Context, resourceGroupName string, clusterName string, dataCenterName string, body DataCenterResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"dataCenterName":    autorest.Encode("path", dataCenterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/cassandraClusters/{clusterName}/dataCenters/{dataCenterName}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateUpdateSender sends the CreateUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client CassandraDataCentersClient) CreateUpdateSender(req *http.Request) (future CassandraDataCentersCreateUpdateFuture, err error) {
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

// CreateUpdateResponder handles the response to the CreateUpdate request. The method always
// closes the http.Response Body.
func (client CassandraDataCentersClient) CreateUpdateResponder(resp *http.Response) (result DataCenterResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a managed Cassandra data center.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// clusterName - managed Cassandra cluster name.
// dataCenterName - data center name in a managed Cassandra cluster.
func (client CassandraDataCentersClient) Delete(ctx context.Context, resourceGroupName string, clusterName string, dataCenterName string) (result CassandraDataCentersDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CassandraDataCentersClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: clusterName,
			Constraints: []validation.Constraint{{Target: "clusterName", Name: validation.MaxLength, Rule: 100, Chain: nil},
				{Target: "clusterName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "clusterName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+(-[a-zA-Z0-9]+)*`, Chain: nil}}},
		{TargetValue: dataCenterName,
			Constraints: []validation.Constraint{{Target: "dataCenterName", Name: validation.MaxLength, Rule: 100, Chain: nil},
				{Target: "dataCenterName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "dataCenterName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+(-[a-zA-Z0-9]+)*`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("documentdb.CassandraDataCentersClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, clusterName, dataCenterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.CassandraDataCentersClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.CassandraDataCentersClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client CassandraDataCentersClient) DeletePreparer(ctx context.Context, resourceGroupName string, clusterName string, dataCenterName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"dataCenterName":    autorest.Encode("path", dataCenterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/cassandraClusters/{clusterName}/dataCenters/{dataCenterName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client CassandraDataCentersClient) DeleteSender(req *http.Request) (future CassandraDataCentersDeleteFuture, err error) {
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
func (client CassandraDataCentersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get the properties of a managed Cassandra data center.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// clusterName - managed Cassandra cluster name.
// dataCenterName - data center name in a managed Cassandra cluster.
func (client CassandraDataCentersClient) Get(ctx context.Context, resourceGroupName string, clusterName string, dataCenterName string) (result DataCenterResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CassandraDataCentersClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: clusterName,
			Constraints: []validation.Constraint{{Target: "clusterName", Name: validation.MaxLength, Rule: 100, Chain: nil},
				{Target: "clusterName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "clusterName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+(-[a-zA-Z0-9]+)*`, Chain: nil}}},
		{TargetValue: dataCenterName,
			Constraints: []validation.Constraint{{Target: "dataCenterName", Name: validation.MaxLength, Rule: 100, Chain: nil},
				{Target: "dataCenterName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "dataCenterName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+(-[a-zA-Z0-9]+)*`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("documentdb.CassandraDataCentersClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, clusterName, dataCenterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.CassandraDataCentersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "documentdb.CassandraDataCentersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.CassandraDataCentersClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client CassandraDataCentersClient) GetPreparer(ctx context.Context, resourceGroupName string, clusterName string, dataCenterName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"dataCenterName":    autorest.Encode("path", dataCenterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/cassandraClusters/{clusterName}/dataCenters/{dataCenterName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client CassandraDataCentersClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client CassandraDataCentersClient) GetResponder(resp *http.Response) (result DataCenterResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list all data centers in a particular managed Cassandra cluster.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// clusterName - managed Cassandra cluster name.
func (client CassandraDataCentersClient) List(ctx context.Context, resourceGroupName string, clusterName string) (result ListDataCenters, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CassandraDataCentersClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: clusterName,
			Constraints: []validation.Constraint{{Target: "clusterName", Name: validation.MaxLength, Rule: 100, Chain: nil},
				{Target: "clusterName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "clusterName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+(-[a-zA-Z0-9]+)*`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("documentdb.CassandraDataCentersClient", "List", err.Error())
	}

	req, err := client.ListPreparer(ctx, resourceGroupName, clusterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.CassandraDataCentersClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "documentdb.CassandraDataCentersClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.CassandraDataCentersClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client CassandraDataCentersClient) ListPreparer(ctx context.Context, resourceGroupName string, clusterName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/cassandraClusters/{clusterName}/dataCenters", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client CassandraDataCentersClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client CassandraDataCentersClient) ListResponder(resp *http.Response) (result ListDataCenters, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update update some of the properties of a managed Cassandra data center.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// clusterName - managed Cassandra cluster name.
// dataCenterName - data center name in a managed Cassandra cluster.
// body - parameters to provide for specifying the managed Cassandra data center.
func (client CassandraDataCentersClient) Update(ctx context.Context, resourceGroupName string, clusterName string, dataCenterName string, body DataCenterResource) (result CassandraDataCentersUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CassandraDataCentersClient.Update")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: clusterName,
			Constraints: []validation.Constraint{{Target: "clusterName", Name: validation.MaxLength, Rule: 100, Chain: nil},
				{Target: "clusterName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "clusterName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+(-[a-zA-Z0-9]+)*`, Chain: nil}}},
		{TargetValue: dataCenterName,
			Constraints: []validation.Constraint{{Target: "dataCenterName", Name: validation.MaxLength, Rule: 100, Chain: nil},
				{Target: "dataCenterName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "dataCenterName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+(-[a-zA-Z0-9]+)*`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("documentdb.CassandraDataCentersClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, clusterName, dataCenterName, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.CassandraDataCentersClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.CassandraDataCentersClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client CassandraDataCentersClient) UpdatePreparer(ctx context.Context, resourceGroupName string, clusterName string, dataCenterName string, body DataCenterResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"dataCenterName":    autorest.Encode("path", dataCenterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/cassandraClusters/{clusterName}/dataCenters/{dataCenterName}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client CassandraDataCentersClient) UpdateSender(req *http.Request) (future CassandraDataCentersUpdateFuture, err error) {
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
func (client CassandraDataCentersClient) UpdateResponder(resp *http.Response) (result DataCenterResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
