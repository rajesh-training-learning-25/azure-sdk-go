package recoveryservicessiterecovery

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
	"net/http"
)

// ReplicationProtectionContainersClient is the client for the ReplicationProtectionContainers methods of the
// Recoveryservicessiterecovery service.
type ReplicationProtectionContainersClient struct {
	ManagementClient
}

// NewReplicationProtectionContainersClient creates an instance of the ReplicationProtectionContainersClient client.
func NewReplicationProtectionContainersClient(subscriptionID string, resourceGroupName string, resourceName string) ReplicationProtectionContainersClient {
	return NewReplicationProtectionContainersClientWithBaseURI(DefaultBaseURI, subscriptionID, resourceGroupName, resourceName)
}

// NewReplicationProtectionContainersClientWithBaseURI creates an instance of the ReplicationProtectionContainersClient
// client.
func NewReplicationProtectionContainersClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string, resourceName string) ReplicationProtectionContainersClient {
	return ReplicationProtectionContainersClient{NewWithBaseURI(baseURI, subscriptionID, resourceGroupName, resourceName)}
}

// NewReplicationProtectionContainersClientWithAuthFile creates an instance of the
// ReplicationProtectionContainersClient client.
func NewReplicationProtectionContainersClientWithAuthFile(resourceGroupName string, resourceName string) (ReplicationProtectionContainersClient, error) {
	c, err := NewWithAuthFile(resourceGroupName, resourceName)
	return ReplicationProtectionContainersClient{c}, err
}

// Create operation to create a protection container. This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// fabricName is unique fabric ARM name. protectionContainerName is unique protection container ARM name. creationInput
// is creation input.
func (client ReplicationProtectionContainersClient) Create(fabricName string, protectionContainerName string, creationInput CreateProtectionContainerInput, cancel <-chan struct{}) (<-chan ProtectionContainer, <-chan error) {
	resultChan := make(chan ProtectionContainer, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result ProtectionContainer
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreatePreparer(fabricName, protectionContainerName, creationInput, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationProtectionContainersClient", "Create", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationProtectionContainersClient", "Create", resp, "Failure sending request")
			return
		}

		result, err = client.CreateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationProtectionContainersClient", "Create", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreatePreparer prepares the Create request.
func (client ReplicationProtectionContainersClient) CreatePreparer(fabricName string, protectionContainerName string, creationInput CreateProtectionContainerInput, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":              autorest.Encode("path", fabricName),
		"protectionContainerName": autorest.Encode("path", protectionContainerName),
		"resourceGroupName":       autorest.Encode("path", client.ResourceGroupName),
		"resourceName":            autorest.Encode("path", client.ResourceName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}", pathParameters),
		autorest.WithJSON(creationInput),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationProtectionContainersClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ReplicationProtectionContainersClient) CreateResponder(resp *http.Response) (result ProtectionContainer, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete operation to remove a protection container. This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// fabricName is unique fabric ARM name. protectionContainerName is unique protection container ARM name.
func (client ReplicationProtectionContainersClient) Delete(fabricName string, protectionContainerName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
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
		req, err := client.DeletePreparer(fabricName, protectionContainerName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationProtectionContainersClient", "Delete", nil, "Failure preparing request")
			return
		}

		resp, err := client.DeleteSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationProtectionContainersClient", "Delete", resp, "Failure sending request")
			return
		}

		result, err = client.DeleteResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationProtectionContainersClient", "Delete", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// DeletePreparer prepares the Delete request.
func (client ReplicationProtectionContainersClient) DeletePreparer(fabricName string, protectionContainerName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":              autorest.Encode("path", fabricName),
		"protectionContainerName": autorest.Encode("path", protectionContainerName),
		"resourceGroupName":       autorest.Encode("path", client.ResourceGroupName),
		"resourceName":            autorest.Encode("path", client.ResourceName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/remove", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationProtectionContainersClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ReplicationProtectionContainersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DiscoverProtectableItem the operation to a add a protectable item to a protection container(Add physical server.)
// This method may poll for completion. Polling can be canceled by passing the cancel channel argument. The channel
// will be used to cancel polling and any outstanding HTTP requests.
//
// fabricName is the name of the fabric. protectionContainerName is the name of the protection container.
// discoverProtectableItemRequest is the request object to add a protectable item.
func (client ReplicationProtectionContainersClient) DiscoverProtectableItem(fabricName string, protectionContainerName string, discoverProtectableItemRequest DiscoverProtectableItemRequest, cancel <-chan struct{}) (<-chan ProtectionContainer, <-chan error) {
	resultChan := make(chan ProtectionContainer, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result ProtectionContainer
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.DiscoverProtectableItemPreparer(fabricName, protectionContainerName, discoverProtectableItemRequest, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationProtectionContainersClient", "DiscoverProtectableItem", nil, "Failure preparing request")
			return
		}

		resp, err := client.DiscoverProtectableItemSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationProtectionContainersClient", "DiscoverProtectableItem", resp, "Failure sending request")
			return
		}

		result, err = client.DiscoverProtectableItemResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationProtectionContainersClient", "DiscoverProtectableItem", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// DiscoverProtectableItemPreparer prepares the DiscoverProtectableItem request.
func (client ReplicationProtectionContainersClient) DiscoverProtectableItemPreparer(fabricName string, protectionContainerName string, discoverProtectableItemRequest DiscoverProtectableItemRequest, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":              autorest.Encode("path", fabricName),
		"protectionContainerName": autorest.Encode("path", protectionContainerName),
		"resourceGroupName":       autorest.Encode("path", client.ResourceGroupName),
		"resourceName":            autorest.Encode("path", client.ResourceName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/discoverProtectableItem", pathParameters),
		autorest.WithJSON(discoverProtectableItemRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DiscoverProtectableItemSender sends the DiscoverProtectableItem request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationProtectionContainersClient) DiscoverProtectableItemSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DiscoverProtectableItemResponder handles the response to the DiscoverProtectableItem request. The method always
// closes the http.Response Body.
func (client ReplicationProtectionContainersClient) DiscoverProtectableItemResponder(resp *http.Response) (result ProtectionContainer, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets the details of a protection container.
//
// fabricName is fabric name. protectionContainerName is protection container name.
func (client ReplicationProtectionContainersClient) Get(fabricName string, protectionContainerName string) (result ProtectionContainer, err error) {
	req, err := client.GetPreparer(fabricName, protectionContainerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationProtectionContainersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationProtectionContainersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationProtectionContainersClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ReplicationProtectionContainersClient) GetPreparer(fabricName string, protectionContainerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":              autorest.Encode("path", fabricName),
		"protectionContainerName": autorest.Encode("path", protectionContainerName),
		"resourceGroupName":       autorest.Encode("path", client.ResourceGroupName),
		"resourceName":            autorest.Encode("path", client.ResourceName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationProtectionContainersClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ReplicationProtectionContainersClient) GetResponder(resp *http.Response) (result ProtectionContainer, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists the protection containers in a vault.
func (client ReplicationProtectionContainersClient) List() (result ProtectionContainerCollection, err error) {
	req, err := client.ListPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationProtectionContainersClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationProtectionContainersClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationProtectionContainersClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ReplicationProtectionContainersClient) ListPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationProtectionContainers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationProtectionContainersClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ReplicationProtectionContainersClient) ListResponder(resp *http.Response) (result ProtectionContainerCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client ReplicationProtectionContainersClient) ListNextResults(lastResults ProtectionContainerCollection) (result ProtectionContainerCollection, err error) {
	req, err := lastResults.ProtectionContainerCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationProtectionContainersClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationProtectionContainersClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationProtectionContainersClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListComplete gets all elements from the list without paging.
func (client ReplicationProtectionContainersClient) ListComplete(cancel <-chan struct{}) (<-chan ProtectionContainer, <-chan error) {
	resultChan := make(chan ProtectionContainer)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.List()
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
			list, err = client.ListNextResults(list)
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

// ListByReplicationFabrics lists the protection containers in the specified fabric.
//
// fabricName is fabric name.
func (client ReplicationProtectionContainersClient) ListByReplicationFabrics(fabricName string) (result ProtectionContainerCollection, err error) {
	req, err := client.ListByReplicationFabricsPreparer(fabricName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationProtectionContainersClient", "ListByReplicationFabrics", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByReplicationFabricsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationProtectionContainersClient", "ListByReplicationFabrics", resp, "Failure sending request")
		return
	}

	result, err = client.ListByReplicationFabricsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationProtectionContainersClient", "ListByReplicationFabrics", resp, "Failure responding to request")
	}

	return
}

// ListByReplicationFabricsPreparer prepares the ListByReplicationFabrics request.
func (client ReplicationProtectionContainersClient) ListByReplicationFabricsPreparer(fabricName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":        autorest.Encode("path", fabricName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByReplicationFabricsSender sends the ListByReplicationFabrics request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationProtectionContainersClient) ListByReplicationFabricsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByReplicationFabricsResponder handles the response to the ListByReplicationFabrics request. The method always
// closes the http.Response Body.
func (client ReplicationProtectionContainersClient) ListByReplicationFabricsResponder(resp *http.Response) (result ProtectionContainerCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByReplicationFabricsNextResults retrieves the next set of results, if any.
func (client ReplicationProtectionContainersClient) ListByReplicationFabricsNextResults(lastResults ProtectionContainerCollection) (result ProtectionContainerCollection, err error) {
	req, err := lastResults.ProtectionContainerCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationProtectionContainersClient", "ListByReplicationFabrics", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByReplicationFabricsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationProtectionContainersClient", "ListByReplicationFabrics", resp, "Failure sending next results request")
	}

	result, err = client.ListByReplicationFabricsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationProtectionContainersClient", "ListByReplicationFabrics", resp, "Failure responding to next results request")
	}

	return
}

// ListByReplicationFabricsComplete gets all elements from the list without paging.
func (client ReplicationProtectionContainersClient) ListByReplicationFabricsComplete(fabricName string, cancel <-chan struct{}) (<-chan ProtectionContainer, <-chan error) {
	resultChan := make(chan ProtectionContainer)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListByReplicationFabrics(fabricName)
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
			list, err = client.ListByReplicationFabricsNextResults(list)
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

// SwitchProtection operation to switch protection from one container to another or one replication provider to
// another. This method may poll for completion. Polling can be canceled by passing the cancel channel argument. The
// channel will be used to cancel polling and any outstanding HTTP requests.
//
// fabricName is unique fabric name. protectionContainerName is protection container name. switchInput is switch
// protection input.
func (client ReplicationProtectionContainersClient) SwitchProtection(fabricName string, protectionContainerName string, switchInput SwitchProtectionInput, cancel <-chan struct{}) (<-chan ProtectionContainer, <-chan error) {
	resultChan := make(chan ProtectionContainer, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result ProtectionContainer
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.SwitchProtectionPreparer(fabricName, protectionContainerName, switchInput, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationProtectionContainersClient", "SwitchProtection", nil, "Failure preparing request")
			return
		}

		resp, err := client.SwitchProtectionSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationProtectionContainersClient", "SwitchProtection", resp, "Failure sending request")
			return
		}

		result, err = client.SwitchProtectionResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationProtectionContainersClient", "SwitchProtection", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// SwitchProtectionPreparer prepares the SwitchProtection request.
func (client ReplicationProtectionContainersClient) SwitchProtectionPreparer(fabricName string, protectionContainerName string, switchInput SwitchProtectionInput, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":              autorest.Encode("path", fabricName),
		"protectionContainerName": autorest.Encode("path", protectionContainerName),
		"resourceGroupName":       autorest.Encode("path", client.ResourceGroupName),
		"resourceName":            autorest.Encode("path", client.ResourceName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/switchprotection", pathParameters),
		autorest.WithJSON(switchInput),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// SwitchProtectionSender sends the SwitchProtection request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationProtectionContainersClient) SwitchProtectionSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// SwitchProtectionResponder handles the response to the SwitchProtection request. The method always
// closes the http.Response Body.
func (client ReplicationProtectionContainersClient) SwitchProtectionResponder(resp *http.Response) (result ProtectionContainer, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
