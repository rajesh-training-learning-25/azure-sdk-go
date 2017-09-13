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
// Code generated by Microsoft (R) AutoRest Code Generator 2.2.18.0
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// ReplicationRecoveryServicesProvidersClient is the client for the ReplicationRecoveryServicesProviders methods of the
// Recoveryservicessiterecovery service.
type ReplicationRecoveryServicesProvidersClient struct {
	ManagementClient
}

// NewReplicationRecoveryServicesProvidersClient creates an instance of the ReplicationRecoveryServicesProvidersClient
// client.
func NewReplicationRecoveryServicesProvidersClient(subscriptionID string, resourceGroupName string, resourceName string) ReplicationRecoveryServicesProvidersClient {
	return NewReplicationRecoveryServicesProvidersClientWithBaseURI(DefaultBaseURI, subscriptionID, resourceGroupName, resourceName)
}

// NewReplicationRecoveryServicesProvidersClientWithBaseURI creates an instance of the
// ReplicationRecoveryServicesProvidersClient client.
func NewReplicationRecoveryServicesProvidersClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string, resourceName string) ReplicationRecoveryServicesProvidersClient {
	return ReplicationRecoveryServicesProvidersClient{NewWithBaseURI(baseURI, subscriptionID, resourceGroupName, resourceName)}
}

// Delete the operation to removes/delete(unregister) a recovery services provider from the vault This method may poll
// for completion. Polling can be canceled by passing the cancel channel argument. The channel will be used to cancel
// polling and any outstanding HTTP requests.
//
// fabricName is fabric name. providerName is recovery services provider name.
func (client ReplicationRecoveryServicesProvidersClient) Delete(fabricName string, providerName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
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
		req, err := client.DeletePreparer(fabricName, providerName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClient", "Delete", nil, "Failure preparing request")
			return
		}

		resp, err := client.DeleteSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClient", "Delete", resp, "Failure sending request")
			return
		}

		result, err = client.DeleteResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClient", "Delete", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// DeletePreparer prepares the Delete request.
func (client ReplicationRecoveryServicesProvidersClient) DeletePreparer(fabricName string, providerName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":        autorest.Encode("path", fabricName),
		"providerName":      autorest.Encode("path", providerName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationRecoveryServicesProviders/{providerName}/remove", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationRecoveryServicesProvidersClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ReplicationRecoveryServicesProvidersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the details of registered recovery services provider.
//
// fabricName is fabric name. providerName is recovery services provider name
func (client ReplicationRecoveryServicesProvidersClient) Get(fabricName string, providerName string) (result RecoveryServicesProvider, err error) {
	req, err := client.GetPreparer(fabricName, providerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ReplicationRecoveryServicesProvidersClient) GetPreparer(fabricName string, providerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":        autorest.Encode("path", fabricName),
		"providerName":      autorest.Encode("path", providerName),
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
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationRecoveryServicesProviders/{providerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationRecoveryServicesProvidersClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ReplicationRecoveryServicesProvidersClient) GetResponder(resp *http.Response) (result RecoveryServicesProvider, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists the registered recovery services providers in the vault
func (client ReplicationRecoveryServicesProvidersClient) List() (result RecoveryServicesProviderCollection, err error) {
	req, err := client.ListPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ReplicationRecoveryServicesProvidersClient) ListPreparer() (*http.Request, error) {
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
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationRecoveryServicesProviders", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationRecoveryServicesProvidersClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ReplicationRecoveryServicesProvidersClient) ListResponder(resp *http.Response) (result RecoveryServicesProviderCollection, err error) {
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
func (client ReplicationRecoveryServicesProvidersClient) ListNextResults(lastResults RecoveryServicesProviderCollection) (result RecoveryServicesProviderCollection, err error) {
	req, err := lastResults.RecoveryServicesProviderCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListComplete gets all elements from the list without paging.
func (client ReplicationRecoveryServicesProvidersClient) ListComplete(cancel <-chan struct{}) (<-chan RecoveryServicesProvider, <-chan error) {
	resultChan := make(chan RecoveryServicesProvider)
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

// ListByReplicationFabrics lists the registered recovery services providers for the specified fabric.
//
// fabricName is fabric name
func (client ReplicationRecoveryServicesProvidersClient) ListByReplicationFabrics(fabricName string) (result RecoveryServicesProviderCollection, err error) {
	req, err := client.ListByReplicationFabricsPreparer(fabricName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClient", "ListByReplicationFabrics", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByReplicationFabricsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClient", "ListByReplicationFabrics", resp, "Failure sending request")
		return
	}

	result, err = client.ListByReplicationFabricsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClient", "ListByReplicationFabrics", resp, "Failure responding to request")
	}

	return
}

// ListByReplicationFabricsPreparer prepares the ListByReplicationFabrics request.
func (client ReplicationRecoveryServicesProvidersClient) ListByReplicationFabricsPreparer(fabricName string) (*http.Request, error) {
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
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationRecoveryServicesProviders", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByReplicationFabricsSender sends the ListByReplicationFabrics request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationRecoveryServicesProvidersClient) ListByReplicationFabricsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByReplicationFabricsResponder handles the response to the ListByReplicationFabrics request. The method always
// closes the http.Response Body.
func (client ReplicationRecoveryServicesProvidersClient) ListByReplicationFabricsResponder(resp *http.Response) (result RecoveryServicesProviderCollection, err error) {
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
func (client ReplicationRecoveryServicesProvidersClient) ListByReplicationFabricsNextResults(lastResults RecoveryServicesProviderCollection) (result RecoveryServicesProviderCollection, err error) {
	req, err := lastResults.RecoveryServicesProviderCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClient", "ListByReplicationFabrics", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByReplicationFabricsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClient", "ListByReplicationFabrics", resp, "Failure sending next results request")
	}

	result, err = client.ListByReplicationFabricsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClient", "ListByReplicationFabrics", resp, "Failure responding to next results request")
	}

	return
}

// ListByReplicationFabricsComplete gets all elements from the list without paging.
func (client ReplicationRecoveryServicesProvidersClient) ListByReplicationFabricsComplete(fabricName string, cancel <-chan struct{}) (<-chan RecoveryServicesProvider, <-chan error) {
	resultChan := make(chan RecoveryServicesProvider)
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

// Purge the operation to purge(force delete) a recovery services provider from the vault. This method may poll for
// completion. Polling can be canceled by passing the cancel channel argument. The channel will be used to cancel
// polling and any outstanding HTTP requests.
//
// fabricName is fabric name. providerName is recovery services provider name.
func (client ReplicationRecoveryServicesProvidersClient) Purge(fabricName string, providerName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
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
		req, err := client.PurgePreparer(fabricName, providerName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClient", "Purge", nil, "Failure preparing request")
			return
		}

		resp, err := client.PurgeSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClient", "Purge", resp, "Failure sending request")
			return
		}

		result, err = client.PurgeResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClient", "Purge", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// PurgePreparer prepares the Purge request.
func (client ReplicationRecoveryServicesProvidersClient) PurgePreparer(fabricName string, providerName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":        autorest.Encode("path", fabricName),
		"providerName":      autorest.Encode("path", providerName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationRecoveryServicesProviders/{providerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// PurgeSender sends the Purge request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationRecoveryServicesProvidersClient) PurgeSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// PurgeResponder handles the response to the Purge request. The method always
// closes the http.Response Body.
func (client ReplicationRecoveryServicesProvidersClient) PurgeResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// RefreshProvider the operation to refresh the information from the recovery services provider. This method may poll
// for completion. Polling can be canceled by passing the cancel channel argument. The channel will be used to cancel
// polling and any outstanding HTTP requests.
//
// fabricName is fabric name. providerName is recovery services provider name.
func (client ReplicationRecoveryServicesProvidersClient) RefreshProvider(fabricName string, providerName string, cancel <-chan struct{}) (<-chan RecoveryServicesProvider, <-chan error) {
	resultChan := make(chan RecoveryServicesProvider, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result RecoveryServicesProvider
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.RefreshProviderPreparer(fabricName, providerName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClient", "RefreshProvider", nil, "Failure preparing request")
			return
		}

		resp, err := client.RefreshProviderSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClient", "RefreshProvider", resp, "Failure sending request")
			return
		}

		result, err = client.RefreshProviderResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClient", "RefreshProvider", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// RefreshProviderPreparer prepares the RefreshProvider request.
func (client ReplicationRecoveryServicesProvidersClient) RefreshProviderPreparer(fabricName string, providerName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":        autorest.Encode("path", fabricName),
		"providerName":      autorest.Encode("path", providerName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationRecoveryServicesProviders/{providerName}/refreshProvider", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// RefreshProviderSender sends the RefreshProvider request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationRecoveryServicesProvidersClient) RefreshProviderSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// RefreshProviderResponder handles the response to the RefreshProvider request. The method always
// closes the http.Response Body.
func (client ReplicationRecoveryServicesProvidersClient) RefreshProviderResponder(resp *http.Response) (result RecoveryServicesProvider, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
