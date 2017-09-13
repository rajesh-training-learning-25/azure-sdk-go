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
// Code generated by Microsoft (R) AutoRest Code Generator 2.2.18.0
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// SyncAgentsClient is the the Azure SQL Database management API provides a RESTful set of web services that interact
// with Azure SQL Database services to manage your databases. The API enables you to create, retrieve, update, and
// delete databases.
type SyncAgentsClient struct {
	ManagementClient
}

// NewSyncAgentsClient creates an instance of the SyncAgentsClient client.
func NewSyncAgentsClient(subscriptionID string) SyncAgentsClient {
	return NewSyncAgentsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSyncAgentsClientWithBaseURI creates an instance of the SyncAgentsClient client.
func NewSyncAgentsClientWithBaseURI(baseURI string, subscriptionID string) SyncAgentsClient {
	return SyncAgentsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a sync agent. This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from the
// Azure Resource Manager API or the portal. serverName is the name of the server on which the sync agent is hosted.
// syncAgentName is the name of the sync agent. parameters is the requested sync agent resource state.
func (client SyncAgentsClient) CreateOrUpdate(resourceGroupName string, serverName string, syncAgentName string, parameters SyncAgent, cancel <-chan struct{}) (<-chan SyncAgent, <-chan error) {
	resultChan := make(chan SyncAgent, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result SyncAgent
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreateOrUpdatePreparer(resourceGroupName, serverName, syncAgentName, parameters, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "CreateOrUpdate", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateOrUpdateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "CreateOrUpdate", resp, "Failure sending request")
			return
		}

		result, err = client.CreateOrUpdateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "CreateOrUpdate", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SyncAgentsClient) CreateOrUpdatePreparer(resourceGroupName string, serverName string, syncAgentName string, parameters SyncAgent, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"syncAgentName":     autorest.Encode("path", syncAgentName),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/syncAgents/{syncAgentName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SyncAgentsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SyncAgentsClient) CreateOrUpdateResponder(resp *http.Response) (result SyncAgent, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a sync agent. This method may poll for completion. Polling can be canceled by passing the cancel
// channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from the
// Azure Resource Manager API or the portal. serverName is the name of the server on which the sync agent is hosted.
// syncAgentName is the name of the sync agent.
func (client SyncAgentsClient) Delete(resourceGroupName string, serverName string, syncAgentName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	resultChan := make(chan autorest.Response, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result autorest.Response
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.DeletePreparer(resourceGroupName, serverName, syncAgentName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "Delete", nil, "Failure preparing request")
			return
		}

		resp, err := client.DeleteSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "Delete", resp, "Failure sending request")
			return
		}

		result, err = client.DeleteResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "Delete", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// DeletePreparer prepares the Delete request.
func (client SyncAgentsClient) DeletePreparer(resourceGroupName string, serverName string, syncAgentName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"syncAgentName":     autorest.Encode("path", syncAgentName),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/syncAgents/{syncAgentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SyncAgentsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SyncAgentsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GenerateKey generates a sync agent key.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from the
// Azure Resource Manager API or the portal. serverName is the name of the server on which the sync agent is hosted.
// syncAgentName is the name of the sync agent.
func (client SyncAgentsClient) GenerateKey(resourceGroupName string, serverName string, syncAgentName string) (result SyncAgentKeyProperties, err error) {
	req, err := client.GenerateKeyPreparer(resourceGroupName, serverName, syncAgentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "GenerateKey", nil, "Failure preparing request")
		return
	}

	resp, err := client.GenerateKeySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "GenerateKey", resp, "Failure sending request")
		return
	}

	result, err = client.GenerateKeyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "GenerateKey", resp, "Failure responding to request")
	}

	return
}

// GenerateKeyPreparer prepares the GenerateKey request.
func (client SyncAgentsClient) GenerateKeyPreparer(resourceGroupName string, serverName string, syncAgentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"syncAgentName":     autorest.Encode("path", syncAgentName),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/syncAgents/{syncAgentName}/generateKey", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GenerateKeySender sends the GenerateKey request. The method will close the
// http.Response Body if it receives an error.
func (client SyncAgentsClient) GenerateKeySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GenerateKeyResponder handles the response to the GenerateKey request. The method always
// closes the http.Response Body.
func (client SyncAgentsClient) GenerateKeyResponder(resp *http.Response) (result SyncAgentKeyProperties, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets a sync agent.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from the
// Azure Resource Manager API or the portal. serverName is the name of the server on which the sync agent is hosted.
// syncAgentName is the name of the sync agent.
func (client SyncAgentsClient) Get(resourceGroupName string, serverName string, syncAgentName string) (result SyncAgent, err error) {
	req, err := client.GetPreparer(resourceGroupName, serverName, syncAgentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client SyncAgentsClient) GetPreparer(resourceGroupName string, serverName string, syncAgentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"syncAgentName":     autorest.Encode("path", syncAgentName),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/syncAgents/{syncAgentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SyncAgentsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SyncAgentsClient) GetResponder(resp *http.Response) (result SyncAgent, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByServer lists sync agents in a server.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from the
// Azure Resource Manager API or the portal. serverName is the name of the server on which the sync agent is hosted.
func (client SyncAgentsClient) ListByServer(resourceGroupName string, serverName string) (result SyncAgentListResult, err error) {
	req, err := client.ListByServerPreparer(resourceGroupName, serverName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "ListByServer", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "ListByServer", resp, "Failure sending request")
		return
	}

	result, err = client.ListByServerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "ListByServer", resp, "Failure responding to request")
	}

	return
}

// ListByServerPreparer prepares the ListByServer request.
func (client SyncAgentsClient) ListByServerPreparer(resourceGroupName string, serverName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/syncAgents", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByServerSender sends the ListByServer request. The method will close the
// http.Response Body if it receives an error.
func (client SyncAgentsClient) ListByServerSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByServerResponder handles the response to the ListByServer request. The method always
// closes the http.Response Body.
func (client SyncAgentsClient) ListByServerResponder(resp *http.Response) (result SyncAgentListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByServerNextResults retrieves the next set of results, if any.
func (client SyncAgentsClient) ListByServerNextResults(lastResults SyncAgentListResult) (result SyncAgentListResult, err error) {
	req, err := lastResults.SyncAgentListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "ListByServer", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByServerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "ListByServer", resp, "Failure sending next results request")
	}

	result, err = client.ListByServerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "ListByServer", resp, "Failure responding to next results request")
	}

	return
}

// ListByServerComplete gets all elements from the list without paging.
func (client SyncAgentsClient) ListByServerComplete(resourceGroupName string, serverName string, cancel <-chan struct{}) (<-chan SyncAgent, <-chan error) {
	resultChan := make(chan SyncAgent)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListByServer(resourceGroupName, serverName)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListByServerNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}

// ListLinkedDatabases lists databases linked to a sync agent.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from the
// Azure Resource Manager API or the portal. serverName is the name of the server on which the sync agent is hosted.
// syncAgentName is the name of the sync agent.
func (client SyncAgentsClient) ListLinkedDatabases(resourceGroupName string, serverName string, syncAgentName string) (result SyncAgentLinkedDatabaseListResult, err error) {
	req, err := client.ListLinkedDatabasesPreparer(resourceGroupName, serverName, syncAgentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "ListLinkedDatabases", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListLinkedDatabasesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "ListLinkedDatabases", resp, "Failure sending request")
		return
	}

	result, err = client.ListLinkedDatabasesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "ListLinkedDatabases", resp, "Failure responding to request")
	}

	return
}

// ListLinkedDatabasesPreparer prepares the ListLinkedDatabases request.
func (client SyncAgentsClient) ListLinkedDatabasesPreparer(resourceGroupName string, serverName string, syncAgentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"syncAgentName":     autorest.Encode("path", syncAgentName),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/syncAgents/{syncAgentName}/linkedDatabases", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListLinkedDatabasesSender sends the ListLinkedDatabases request. The method will close the
// http.Response Body if it receives an error.
func (client SyncAgentsClient) ListLinkedDatabasesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListLinkedDatabasesResponder handles the response to the ListLinkedDatabases request. The method always
// closes the http.Response Body.
func (client SyncAgentsClient) ListLinkedDatabasesResponder(resp *http.Response) (result SyncAgentLinkedDatabaseListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListLinkedDatabasesNextResults retrieves the next set of results, if any.
func (client SyncAgentsClient) ListLinkedDatabasesNextResults(lastResults SyncAgentLinkedDatabaseListResult) (result SyncAgentLinkedDatabaseListResult, err error) {
	req, err := lastResults.SyncAgentLinkedDatabaseListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "ListLinkedDatabases", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListLinkedDatabasesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "ListLinkedDatabases", resp, "Failure sending next results request")
	}

	result, err = client.ListLinkedDatabasesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "ListLinkedDatabases", resp, "Failure responding to next results request")
	}

	return
}

// ListLinkedDatabasesComplete gets all elements from the list without paging.
func (client SyncAgentsClient) ListLinkedDatabasesComplete(resourceGroupName string, serverName string, syncAgentName string, cancel <-chan struct{}) (<-chan SyncAgentLinkedDatabase, <-chan error) {
	resultChan := make(chan SyncAgentLinkedDatabase)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListLinkedDatabases(resourceGroupName, serverName, syncAgentName)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListLinkedDatabasesNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}
