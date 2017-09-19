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
// Code generated by Microsoft (R) AutoRest Code Generator 2.2.21.0
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// VirtualNetworksClient is the network Client
type VirtualNetworksClient struct {
	ManagementClient
}

// NewVirtualNetworksClient creates an instance of the VirtualNetworksClient client.
func NewVirtualNetworksClient(subscriptionID string) VirtualNetworksClient {
	return NewVirtualNetworksClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVirtualNetworksClientWithBaseURI creates an instance of the VirtualNetworksClient client.
func NewVirtualNetworksClientWithBaseURI(baseURI string, subscriptionID string) VirtualNetworksClient {
	return VirtualNetworksClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CheckIPAddressAvailability checks whether a private IP address is available for use.
//
// resourceGroupName is the name of the resource group. virtualNetworkName is the name of the virtual network.
// IPAddress is the private IP address to be verified.
func (client VirtualNetworksClient) CheckIPAddressAvailability(resourceGroupName string, virtualNetworkName string, IPAddress string) (result IPAddressAvailabilityResult, err error) {
	if err := validation.Validate([]validation.Validation{}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "network.VirtualNetworksClient", "CheckIPAddressAvailability")
	}

	req, err := client.CheckIPAddressAvailabilityPreparer(resourceGroupName, virtualNetworkName, IPAddress)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "CheckIPAddressAvailability", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckIPAddressAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "CheckIPAddressAvailability", resp, "Failure sending request")
		return
	}

	result, err = client.CheckIPAddressAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "CheckIPAddressAvailability", resp, "Failure responding to request")
	}

	return
}

// CheckIPAddressAvailabilityPreparer prepares the CheckIPAddressAvailability request.
func (client VirtualNetworksClient) CheckIPAddressAvailabilityPreparer(resourceGroupName string, virtualNetworkName string, IPAddress string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkName": autorest.Encode("path", virtualNetworkName),
	}

	const APIVersion = "2017-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(IPAddress) > 0 {
		queryParameters["ipAddress"] = autorest.Encode("query", IPAddress)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/CheckIPAddressAvailability", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CheckIPAddressAvailabilitySender sends the CheckIPAddressAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworksClient) CheckIPAddressAvailabilitySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CheckIPAddressAvailabilityResponder handles the response to the CheckIPAddressAvailability request. The method always
// closes the http.Response Body.
func (client VirtualNetworksClient) CheckIPAddressAvailabilityResponder(resp *http.Response) (result IPAddressAvailabilityResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdate creates or updates a virtual network in the specified resource group. This method may poll for
// completion. Polling can be canceled by passing the cancel channel argument. The channel will be used to cancel
// polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. virtualNetworkName is the name of the virtual network.
// parameters is parameters supplied to the create or update virtual network operation
func (client VirtualNetworksClient) CreateOrUpdate(resourceGroupName string, virtualNetworkName string, parameters VirtualNetwork, cancel <-chan struct{}) (<-chan VirtualNetwork, <-chan error) {
	resultChan := make(chan VirtualNetwork, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "network.VirtualNetworksClient", "CreateOrUpdate")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result VirtualNetwork
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreateOrUpdatePreparer(resourceGroupName, virtualNetworkName, parameters, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "CreateOrUpdate", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateOrUpdateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "CreateOrUpdate", resp, "Failure sending request")
			return
		}

		result, err = client.CreateOrUpdateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "CreateOrUpdate", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VirtualNetworksClient) CreateOrUpdatePreparer(resourceGroupName string, virtualNetworkName string, parameters VirtualNetwork, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkName": autorest.Encode("path", virtualNetworkName),
	}

	const APIVersion = "2017-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworksClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client VirtualNetworksClient) CreateOrUpdateResponder(resp *http.Response) (result VirtualNetwork, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified virtual network. This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. virtualNetworkName is the name of the virtual network.
func (client VirtualNetworksClient) Delete(resourceGroupName string, virtualNetworkName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	resultChan := make(chan autorest.Response, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "network.VirtualNetworksClient", "Delete")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

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
		req, err := client.DeletePreparer(resourceGroupName, virtualNetworkName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "Delete", nil, "Failure preparing request")
			return
		}

		resp, err := client.DeleteSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "Delete", resp, "Failure sending request")
			return
		}

		result, err = client.DeleteResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "Delete", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// DeletePreparer prepares the Delete request.
func (client VirtualNetworksClient) DeletePreparer(resourceGroupName string, virtualNetworkName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkName": autorest.Encode("path", virtualNetworkName),
	}

	const APIVersion = "2017-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworksClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client VirtualNetworksClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusAccepted, http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the specified virtual network by resource group.
//
// resourceGroupName is the name of the resource group. virtualNetworkName is the name of the virtual network. expand
// is expands referenced resources.
func (client VirtualNetworksClient) Get(resourceGroupName string, virtualNetworkName string, expand string) (result VirtualNetwork, err error) {
	if err := validation.Validate([]validation.Validation{}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "network.VirtualNetworksClient", "Get")
	}

	req, err := client.GetPreparer(resourceGroupName, virtualNetworkName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client VirtualNetworksClient) GetPreparer(resourceGroupName string, virtualNetworkName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkName": autorest.Encode("path", virtualNetworkName),
	}

	const APIVersion = "2017-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworksClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VirtualNetworksClient) GetResponder(resp *http.Response) (result VirtualNetwork, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all virtual networks in a resource group.
//
// resourceGroupName is the name of the resource group.
func (client VirtualNetworksClient) List(resourceGroupName string) (result VirtualNetworkListResult, err error) {
	if err := validation.Validate([]validation.Validation{}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "network.VirtualNetworksClient", "List")
	}

	req, err := client.ListPreparer(resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client VirtualNetworksClient) ListPreparer(resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworksClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VirtualNetworksClient) ListResponder(resp *http.Response) (result VirtualNetworkListResult, err error) {
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
func (client VirtualNetworksClient) ListNextResults(lastResults VirtualNetworkListResult) (result VirtualNetworkListResult, err error) {
	req, err := lastResults.VirtualNetworkListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListComplete gets all elements from the list without paging.
func (client VirtualNetworksClient) ListComplete(resourceGroupName string, cancel <-chan struct{}) (<-chan VirtualNetwork, <-chan error) {
	resultChan := make(chan VirtualNetwork)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.List(resourceGroupName)
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

// ListAll gets all virtual networks in a subscription.
func (client VirtualNetworksClient) ListAll() (result VirtualNetworkListResult, err error) {
	if err := validation.Validate([]validation.Validation{}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "network.VirtualNetworksClient", "ListAll")
	}

	req, err := client.ListAllPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "ListAll", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "ListAll", resp, "Failure sending request")
		return
	}

	result, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "ListAll", resp, "Failure responding to request")
	}

	return
}

// ListAllPreparer prepares the ListAll request.
func (client VirtualNetworksClient) ListAllPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualNetworks", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListAllSender sends the ListAll request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworksClient) ListAllSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListAllResponder handles the response to the ListAll request. The method always
// closes the http.Response Body.
func (client VirtualNetworksClient) ListAllResponder(resp *http.Response) (result VirtualNetworkListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListAllNextResults retrieves the next set of results, if any.
func (client VirtualNetworksClient) ListAllNextResults(lastResults VirtualNetworkListResult) (result VirtualNetworkListResult, err error) {
	req, err := lastResults.VirtualNetworkListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "ListAll", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "ListAll", resp, "Failure sending next results request")
	}

	result, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "ListAll", resp, "Failure responding to next results request")
	}

	return
}

// ListAllComplete gets all elements from the list without paging.
func (client VirtualNetworksClient) ListAllComplete(cancel <-chan struct{}) (<-chan VirtualNetwork, <-chan error) {
	resultChan := make(chan VirtualNetwork)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListAll()
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
			list, err = client.ListAllNextResults(list)
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

// ListUsage lists usage stats.
//
// resourceGroupName is the name of the resource group. virtualNetworkName is the name of the virtual network.
func (client VirtualNetworksClient) ListUsage(resourceGroupName string, virtualNetworkName string) (result VirtualNetworkListUsageResult, err error) {
	if err := validation.Validate([]validation.Validation{}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "network.VirtualNetworksClient", "ListUsage")
	}

	req, err := client.ListUsagePreparer(resourceGroupName, virtualNetworkName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "ListUsage", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListUsageSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "ListUsage", resp, "Failure sending request")
		return
	}

	result, err = client.ListUsageResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "ListUsage", resp, "Failure responding to request")
	}

	return
}

// ListUsagePreparer prepares the ListUsage request.
func (client VirtualNetworksClient) ListUsagePreparer(resourceGroupName string, virtualNetworkName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkName": autorest.Encode("path", virtualNetworkName),
	}

	const APIVersion = "2017-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/usages", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListUsageSender sends the ListUsage request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworksClient) ListUsageSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListUsageResponder handles the response to the ListUsage request. The method always
// closes the http.Response Body.
func (client VirtualNetworksClient) ListUsageResponder(resp *http.Response) (result VirtualNetworkListUsageResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListUsageNextResults retrieves the next set of results, if any.
func (client VirtualNetworksClient) ListUsageNextResults(lastResults VirtualNetworkListUsageResult) (result VirtualNetworkListUsageResult, err error) {
	req, err := lastResults.VirtualNetworkListUsageResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "ListUsage", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListUsageSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "ListUsage", resp, "Failure sending next results request")
	}

	result, err = client.ListUsageResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "ListUsage", resp, "Failure responding to next results request")
	}

	return
}

// ListUsageComplete gets all elements from the list without paging.
func (client VirtualNetworksClient) ListUsageComplete(resourceGroupName string, virtualNetworkName string, cancel <-chan struct{}) (<-chan VirtualNetworkUsage, <-chan error) {
	resultChan := make(chan VirtualNetworkUsage)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListUsage(resourceGroupName, virtualNetworkName)
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
			list, err = client.ListUsageNextResults(list)
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
