package sql

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
	"github.com/satori/go.uuid"
	"net/http"
)

// ManagedDatabasesClient is the the Azure SQL Database management API provides a RESTful set of web services that
// interact with Azure SQL Database services to manage your databases. The API enables you to create, retrieve, update,
// and delete databases.
type ManagedDatabasesClient struct {
	BaseClient
}

// NewManagedDatabasesClient creates an instance of the ManagedDatabasesClient client.
func NewManagedDatabasesClient(subscriptionID string) ManagedDatabasesClient {
	return NewManagedDatabasesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewManagedDatabasesClientWithBaseURI creates an instance of the ManagedDatabasesClient client.
func NewManagedDatabasesClientWithBaseURI(baseURI string, subscriptionID string) ManagedDatabasesClient {
	return ManagedDatabasesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CompleteRestore completes the restore operation on a managed database.
// Parameters:
// locationName - the name of the region where the resource is located.
// operationID - management operation id that this request tries to complete.
// parameters - the definition for completing the restore of this managed database.
func (client ManagedDatabasesClient) CompleteRestore(ctx context.Context, locationName string, operationID uuid.UUID, parameters CompleteDatabaseRestoreDefinition) (result ManagedDatabasesCompleteRestoreFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.LastBackupName", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("sql.ManagedDatabasesClient", "CompleteRestore", err.Error())
	}

	req, err := client.CompleteRestorePreparer(ctx, locationName, operationID, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabasesClient", "CompleteRestore", nil, "Failure preparing request")
		return
	}

	result, err = client.CompleteRestoreSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabasesClient", "CompleteRestore", result.Response(), "Failure sending request")
		return
	}

	return
}

// CompleteRestorePreparer prepares the CompleteRestore request.
func (client ManagedDatabasesClient) CompleteRestorePreparer(ctx context.Context, locationName string, operationID uuid.UUID, parameters CompleteDatabaseRestoreDefinition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"locationName":   autorest.Encode("path", locationName),
		"operationId":    autorest.Encode("path", operationID),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationName}/managedDatabaseRestoreAzureAsyncOperation/{operationId}/completeRestore", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CompleteRestoreSender sends the CompleteRestore request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedDatabasesClient) CompleteRestoreSender(req *http.Request) (future ManagedDatabasesCompleteRestoreFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	err = autorest.Respond(resp, azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CompleteRestoreResponder handles the response to the CompleteRestore request. The method always
// closes the http.Response Body.
func (client ManagedDatabasesClient) CompleteRestoreResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// CreateOrUpdate creates a new database or updates an existing database.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// managedInstanceName - the name of the managed instance.
// databaseName - the name of the database.
// parameters - the requested database resource state.
func (client ManagedDatabasesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, parameters ManagedDatabase) (result ManagedDatabasesCreateOrUpdateFuture, err error) {
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, managedInstanceName, databaseName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabasesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabasesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ManagedDatabasesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, parameters ManagedDatabase) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":        autorest.Encode("path", databaseName),
		"managedInstanceName": autorest.Encode("path", managedInstanceName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedDatabasesClient) CreateOrUpdateSender(req *http.Request) (future ManagedDatabasesCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	err = autorest.Respond(resp, azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ManagedDatabasesClient) CreateOrUpdateResponder(resp *http.Response) (result ManagedDatabase, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the managed database.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// managedInstanceName - the name of the managed instance.
// databaseName - the name of the database.
func (client ManagedDatabasesClient) Delete(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string) (result ManagedDatabasesDeleteFuture, err error) {
	req, err := client.DeletePreparer(ctx, resourceGroupName, managedInstanceName, databaseName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabasesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabasesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ManagedDatabasesClient) DeletePreparer(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":        autorest.Encode("path", databaseName),
		"managedInstanceName": autorest.Encode("path", managedInstanceName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedDatabasesClient) DeleteSender(req *http.Request) (future ManagedDatabasesDeleteFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	err = autorest.Respond(resp, azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ManagedDatabasesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a managed database.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// managedInstanceName - the name of the managed instance.
// databaseName - the name of the database.
func (client ManagedDatabasesClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string) (result ManagedDatabase, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, managedInstanceName, databaseName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabasesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabasesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabasesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ManagedDatabasesClient) GetPreparer(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":        autorest.Encode("path", databaseName),
		"managedInstanceName": autorest.Encode("path", managedInstanceName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedDatabasesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ManagedDatabasesClient) GetResponder(resp *http.Response) (result ManagedDatabase, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByInstance gets a list of managed databases.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// managedInstanceName - the name of the managed instance.
func (client ManagedDatabasesClient) ListByInstance(ctx context.Context, resourceGroupName string, managedInstanceName string) (result ManagedDatabaseListResultPage, err error) {
	result.fn = client.listByInstanceNextResults
	req, err := client.ListByInstancePreparer(ctx, resourceGroupName, managedInstanceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabasesClient", "ListByInstance", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByInstanceSender(req)
	if err != nil {
		result.mdlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabasesClient", "ListByInstance", resp, "Failure sending request")
		return
	}

	result.mdlr, err = client.ListByInstanceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabasesClient", "ListByInstance", resp, "Failure responding to request")
	}

	return
}

// ListByInstancePreparer prepares the ListByInstance request.
func (client ManagedDatabasesClient) ListByInstancePreparer(ctx context.Context, resourceGroupName string, managedInstanceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managedInstanceName": autorest.Encode("path", managedInstanceName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByInstanceSender sends the ListByInstance request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedDatabasesClient) ListByInstanceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByInstanceResponder handles the response to the ListByInstance request. The method always
// closes the http.Response Body.
func (client ManagedDatabasesClient) ListByInstanceResponder(resp *http.Response) (result ManagedDatabaseListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByInstanceNextResults retrieves the next set of results, if any.
func (client ManagedDatabasesClient) listByInstanceNextResults(lastResults ManagedDatabaseListResult) (result ManagedDatabaseListResult, err error) {
	req, err := lastResults.managedDatabaseListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.ManagedDatabasesClient", "listByInstanceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByInstanceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.ManagedDatabasesClient", "listByInstanceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByInstanceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabasesClient", "listByInstanceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByInstanceComplete enumerates all values, automatically crossing page boundaries as required.
func (client ManagedDatabasesClient) ListByInstanceComplete(ctx context.Context, resourceGroupName string, managedInstanceName string) (result ManagedDatabaseListResultIterator, err error) {
	result.page, err = client.ListByInstance(ctx, resourceGroupName, managedInstanceName)
	return
}

// Update updates an existing database.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// managedInstanceName - the name of the managed instance.
// databaseName - the name of the database.
// parameters - the requested database resource state.
func (client ManagedDatabasesClient) Update(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, parameters ManagedDatabaseUpdate) (result ManagedDatabasesUpdateFuture, err error) {
	req, err := client.UpdatePreparer(ctx, resourceGroupName, managedInstanceName, databaseName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabasesClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabasesClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ManagedDatabasesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, parameters ManagedDatabaseUpdate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":        autorest.Encode("path", databaseName),
		"managedInstanceName": autorest.Encode("path", managedInstanceName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedDatabasesClient) UpdateSender(req *http.Request) (future ManagedDatabasesUpdateFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	err = autorest.Respond(resp, azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ManagedDatabasesClient) UpdateResponder(resp *http.Response) (result ManagedDatabase, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
