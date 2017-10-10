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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// FailoverGroupsClient is the the Azure SQL Database management API provides a RESTful set of web services that
// interact with Azure SQL Database services to manage your databases. The API enables you to create, retrieve, update,
// and delete databases.
type FailoverGroupsClient struct {
	ManagementClient
}

// NewFailoverGroupsClient creates an instance of the FailoverGroupsClient client.
func NewFailoverGroupsClient(subscriptionID string) FailoverGroupsClient {
	return NewFailoverGroupsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewFailoverGroupsClientWithBaseURI creates an instance of the FailoverGroupsClient client.
func NewFailoverGroupsClientWithBaseURI(baseURI string, subscriptionID string) FailoverGroupsClient {
	return FailoverGroupsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// NewFailoverGroupsClientWithAuthFile creates an instance of the FailoverGroupsClient client.
func NewFailoverGroupsClientWithAuthFile() (FailoverGroupsClient, error) {
	c, err := NewWithAuthFile()
	return FailoverGroupsClient{c}, err
}

// CreateOrUpdate creates or updates a failover group. This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from the
// Azure Resource Manager API or the portal. serverName is the name of the server containing the failover group.
// failoverGroupName is the name of the failover group. parameters is the failover group parameters.
func (client FailoverGroupsClient) CreateOrUpdate(resourceGroupName string, serverName string, failoverGroupName string, parameters FailoverGroup, cancel <-chan struct{}) (<-chan FailoverGroup, <-chan error) {
	resultChan := make(chan FailoverGroup, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.FailoverGroupProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.FailoverGroupProperties.ReadWriteEndpoint", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.FailoverGroupProperties.PartnerServers", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "sql.FailoverGroupsClient", "CreateOrUpdate")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result FailoverGroup
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreateOrUpdatePreparer(resourceGroupName, serverName, failoverGroupName, parameters, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "CreateOrUpdate", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateOrUpdateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "CreateOrUpdate", resp, "Failure sending request")
			return
		}

		result, err = client.CreateOrUpdateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "CreateOrUpdate", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client FailoverGroupsClient) CreateOrUpdatePreparer(resourceGroupName string, serverName string, failoverGroupName string, parameters FailoverGroup, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"failoverGroupName": autorest.Encode("path", failoverGroupName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/failoverGroups/{failoverGroupName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client FailoverGroupsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client FailoverGroupsClient) CreateOrUpdateResponder(resp *http.Response) (result FailoverGroup, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a failover group. This method may poll for completion. Polling can be canceled by passing the cancel
// channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from the
// Azure Resource Manager API or the portal. serverName is the name of the server containing the failover group.
// failoverGroupName is the name of the failover group.
func (client FailoverGroupsClient) Delete(resourceGroupName string, serverName string, failoverGroupName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
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
		req, err := client.DeletePreparer(resourceGroupName, serverName, failoverGroupName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "Delete", nil, "Failure preparing request")
			return
		}

		resp, err := client.DeleteSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "Delete", resp, "Failure sending request")
			return
		}

		result, err = client.DeleteResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "Delete", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// DeletePreparer prepares the Delete request.
func (client FailoverGroupsClient) DeletePreparer(resourceGroupName string, serverName string, failoverGroupName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"failoverGroupName": autorest.Encode("path", failoverGroupName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/failoverGroups/{failoverGroupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client FailoverGroupsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client FailoverGroupsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Failover fails over from the current primary server to this server. This method may poll for completion. Polling can
// be canceled by passing the cancel channel argument. The channel will be used to cancel polling and any outstanding
// HTTP requests.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from the
// Azure Resource Manager API or the portal. serverName is the name of the server containing the failover group.
// failoverGroupName is the name of the failover group.
func (client FailoverGroupsClient) Failover(resourceGroupName string, serverName string, failoverGroupName string, cancel <-chan struct{}) (<-chan FailoverGroup, <-chan error) {
	resultChan := make(chan FailoverGroup, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result FailoverGroup
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.FailoverPreparer(resourceGroupName, serverName, failoverGroupName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "Failover", nil, "Failure preparing request")
			return
		}

		resp, err := client.FailoverSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "Failover", resp, "Failure sending request")
			return
		}

		result, err = client.FailoverResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "Failover", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// FailoverPreparer prepares the Failover request.
func (client FailoverGroupsClient) FailoverPreparer(resourceGroupName string, serverName string, failoverGroupName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"failoverGroupName": autorest.Encode("path", failoverGroupName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/failoverGroups/{failoverGroupName}/failover", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// FailoverSender sends the Failover request. The method will close the
// http.Response Body if it receives an error.
func (client FailoverGroupsClient) FailoverSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// FailoverResponder handles the response to the Failover request. The method always
// closes the http.Response Body.
func (client FailoverGroupsClient) FailoverResponder(resp *http.Response) (result FailoverGroup, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ForceFailoverAllowDataLoss fails over from the current primary server to this server. This operation might result in
// data loss. This method may poll for completion. Polling can be canceled by passing the cancel channel argument. The
// channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from the
// Azure Resource Manager API or the portal. serverName is the name of the server containing the failover group.
// failoverGroupName is the name of the failover group.
func (client FailoverGroupsClient) ForceFailoverAllowDataLoss(resourceGroupName string, serverName string, failoverGroupName string, cancel <-chan struct{}) (<-chan FailoverGroup, <-chan error) {
	resultChan := make(chan FailoverGroup, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result FailoverGroup
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.ForceFailoverAllowDataLossPreparer(resourceGroupName, serverName, failoverGroupName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "ForceFailoverAllowDataLoss", nil, "Failure preparing request")
			return
		}

		resp, err := client.ForceFailoverAllowDataLossSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "ForceFailoverAllowDataLoss", resp, "Failure sending request")
			return
		}

		result, err = client.ForceFailoverAllowDataLossResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "ForceFailoverAllowDataLoss", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// ForceFailoverAllowDataLossPreparer prepares the ForceFailoverAllowDataLoss request.
func (client FailoverGroupsClient) ForceFailoverAllowDataLossPreparer(resourceGroupName string, serverName string, failoverGroupName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"failoverGroupName": autorest.Encode("path", failoverGroupName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/failoverGroups/{failoverGroupName}/forceFailoverAllowDataLoss", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// ForceFailoverAllowDataLossSender sends the ForceFailoverAllowDataLoss request. The method will close the
// http.Response Body if it receives an error.
func (client FailoverGroupsClient) ForceFailoverAllowDataLossSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// ForceFailoverAllowDataLossResponder handles the response to the ForceFailoverAllowDataLoss request. The method always
// closes the http.Response Body.
func (client FailoverGroupsClient) ForceFailoverAllowDataLossResponder(resp *http.Response) (result FailoverGroup, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets a failover group.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from the
// Azure Resource Manager API or the portal. serverName is the name of the server containing the failover group.
// failoverGroupName is the name of the failover group.
func (client FailoverGroupsClient) Get(resourceGroupName string, serverName string, failoverGroupName string) (result FailoverGroup, err error) {
	req, err := client.GetPreparer(resourceGroupName, serverName, failoverGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client FailoverGroupsClient) GetPreparer(resourceGroupName string, serverName string, failoverGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"failoverGroupName": autorest.Encode("path", failoverGroupName),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/failoverGroups/{failoverGroupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client FailoverGroupsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client FailoverGroupsClient) GetResponder(resp *http.Response) (result FailoverGroup, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByServer lists the failover groups in a server.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from the
// Azure Resource Manager API or the portal. serverName is the name of the server containing the failover group.
func (client FailoverGroupsClient) ListByServer(resourceGroupName string, serverName string) (result FailoverGroupListResult, err error) {
	req, err := client.ListByServerPreparer(resourceGroupName, serverName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "ListByServer", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "ListByServer", resp, "Failure sending request")
		return
	}

	result, err = client.ListByServerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "ListByServer", resp, "Failure responding to request")
	}

	return
}

// ListByServerPreparer prepares the ListByServer request.
func (client FailoverGroupsClient) ListByServerPreparer(resourceGroupName string, serverName string) (*http.Request, error) {
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/failoverGroups", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByServerSender sends the ListByServer request. The method will close the
// http.Response Body if it receives an error.
func (client FailoverGroupsClient) ListByServerSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByServerResponder handles the response to the ListByServer request. The method always
// closes the http.Response Body.
func (client FailoverGroupsClient) ListByServerResponder(resp *http.Response) (result FailoverGroupListResult, err error) {
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
func (client FailoverGroupsClient) ListByServerNextResults(lastResults FailoverGroupListResult) (result FailoverGroupListResult, err error) {
	req, err := lastResults.FailoverGroupListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "ListByServer", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByServerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "ListByServer", resp, "Failure sending next results request")
	}

	result, err = client.ListByServerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "ListByServer", resp, "Failure responding to next results request")
	}

	return
}

// ListByServerComplete gets all elements from the list without paging.
func (client FailoverGroupsClient) ListByServerComplete(resourceGroupName string, serverName string, cancel <-chan struct{}) (<-chan FailoverGroup, <-chan error) {
	resultChan := make(chan FailoverGroup)
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

// Update updates a failover group. This method may poll for completion. Polling can be canceled by passing the cancel
// channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from the
// Azure Resource Manager API or the portal. serverName is the name of the server containing the failover group.
// failoverGroupName is the name of the failover group. parameters is the failover group parameters.
func (client FailoverGroupsClient) Update(resourceGroupName string, serverName string, failoverGroupName string, parameters FailoverGroupUpdate, cancel <-chan struct{}) (<-chan FailoverGroup, <-chan error) {
	resultChan := make(chan FailoverGroup, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result FailoverGroup
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.UpdatePreparer(resourceGroupName, serverName, failoverGroupName, parameters, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "Update", nil, "Failure preparing request")
			return
		}

		resp, err := client.UpdateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "Update", resp, "Failure sending request")
			return
		}

		result, err = client.UpdateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "Update", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// UpdatePreparer prepares the Update request.
func (client FailoverGroupsClient) UpdatePreparer(resourceGroupName string, serverName string, failoverGroupName string, parameters FailoverGroupUpdate, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"failoverGroupName": autorest.Encode("path", failoverGroupName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/failoverGroups/{failoverGroupName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client FailoverGroupsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client FailoverGroupsClient) UpdateResponder(resp *http.Response) (result FailoverGroup, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
