package compute

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

// VirtualMachineScaleSetExtensionsClient is the compute Client
type VirtualMachineScaleSetExtensionsClient struct {
	ManagementClient
}

// NewVirtualMachineScaleSetExtensionsClient creates an instance of the VirtualMachineScaleSetExtensionsClient client.
func NewVirtualMachineScaleSetExtensionsClient(subscriptionID string) VirtualMachineScaleSetExtensionsClient {
	return NewVirtualMachineScaleSetExtensionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVirtualMachineScaleSetExtensionsClientWithBaseURI creates an instance of the
// VirtualMachineScaleSetExtensionsClient client.
func NewVirtualMachineScaleSetExtensionsClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineScaleSetExtensionsClient {
	return VirtualMachineScaleSetExtensionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// NewVirtualMachineScaleSetExtensionsClientWithAuthFile creates an instance of the
// VirtualMachineScaleSetExtensionsClient client.
func NewVirtualMachineScaleSetExtensionsClientWithAuthFile() (VirtualMachineScaleSetExtensionsClient, error) {
	c, err := NewWithAuthFile()
	return VirtualMachineScaleSetExtensionsClient{c}, err
}

// CreateOrUpdate the operation to create or update an extension. This method may poll for completion. Polling can be
// canceled by passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP
// requests.
//
// resourceGroupName is the name of the resource group. VMScaleSetName is the name of the VM scale set where the
// extension should be create or updated. vmssExtensionName is the name of the VM scale set extension.
// extensionParameters is parameters supplied to the Create VM scale set Extension operation.
func (client VirtualMachineScaleSetExtensionsClient) CreateOrUpdate(resourceGroupName string, VMScaleSetName string, vmssExtensionName string, extensionParameters VirtualMachineScaleSetExtension, cancel <-chan struct{}) (<-chan VirtualMachineScaleSetExtension, <-chan error) {
	resultChan := make(chan VirtualMachineScaleSetExtension, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result VirtualMachineScaleSetExtension
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreateOrUpdatePreparer(resourceGroupName, VMScaleSetName, vmssExtensionName, extensionParameters, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetExtensionsClient", "CreateOrUpdate", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateOrUpdateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetExtensionsClient", "CreateOrUpdate", resp, "Failure sending request")
			return
		}

		result, err = client.CreateOrUpdateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetExtensionsClient", "CreateOrUpdate", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VirtualMachineScaleSetExtensionsClient) CreateOrUpdatePreparer(resourceGroupName string, VMScaleSetName string, vmssExtensionName string, extensionParameters VirtualMachineScaleSetExtension, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmScaleSetName":    autorest.Encode("path", VMScaleSetName),
		"vmssExtensionName": autorest.Encode("path", vmssExtensionName),
	}

	const APIVersion = "2017-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/extensions/{vmssExtensionName}", pathParameters),
		autorest.WithJSON(extensionParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineScaleSetExtensionsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client VirtualMachineScaleSetExtensionsClient) CreateOrUpdateResponder(resp *http.Response) (result VirtualMachineScaleSetExtension, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete the operation to delete the extension. This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. VMScaleSetName is the name of the VM scale set where the
// extension should be deleted. vmssExtensionName is the name of the VM scale set extension.
func (client VirtualMachineScaleSetExtensionsClient) Delete(resourceGroupName string, VMScaleSetName string, vmssExtensionName string, cancel <-chan struct{}) (<-chan OperationStatusResponse, <-chan error) {
	resultChan := make(chan OperationStatusResponse, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result OperationStatusResponse
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.DeletePreparer(resourceGroupName, VMScaleSetName, vmssExtensionName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetExtensionsClient", "Delete", nil, "Failure preparing request")
			return
		}

		resp, err := client.DeleteSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetExtensionsClient", "Delete", resp, "Failure sending request")
			return
		}

		result, err = client.DeleteResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetExtensionsClient", "Delete", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// DeletePreparer prepares the Delete request.
func (client VirtualMachineScaleSetExtensionsClient) DeletePreparer(resourceGroupName string, VMScaleSetName string, vmssExtensionName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmScaleSetName":    autorest.Encode("path", VMScaleSetName),
		"vmssExtensionName": autorest.Encode("path", vmssExtensionName),
	}

	const APIVersion = "2017-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/extensions/{vmssExtensionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineScaleSetExtensionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client VirtualMachineScaleSetExtensionsClient) DeleteResponder(resp *http.Response) (result OperationStatusResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get the operation to get the extension.
//
// resourceGroupName is the name of the resource group. VMScaleSetName is the name of the VM scale set containing the
// extension. vmssExtensionName is the name of the VM scale set extension. expand is the expand expression to apply on
// the operation.
func (client VirtualMachineScaleSetExtensionsClient) Get(resourceGroupName string, VMScaleSetName string, vmssExtensionName string, expand string) (result VirtualMachineScaleSetExtension, err error) {
	req, err := client.GetPreparer(resourceGroupName, VMScaleSetName, vmssExtensionName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetExtensionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetExtensionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetExtensionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client VirtualMachineScaleSetExtensionsClient) GetPreparer(resourceGroupName string, VMScaleSetName string, vmssExtensionName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmScaleSetName":    autorest.Encode("path", VMScaleSetName),
		"vmssExtensionName": autorest.Encode("path", vmssExtensionName),
	}

	const APIVersion = "2017-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/extensions/{vmssExtensionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineScaleSetExtensionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VirtualMachineScaleSetExtensionsClient) GetResponder(resp *http.Response) (result VirtualMachineScaleSetExtension, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list of all extensions in a VM scale set.
//
// resourceGroupName is the name of the resource group. VMScaleSetName is the name of the VM scale set containing the
// extension.
func (client VirtualMachineScaleSetExtensionsClient) List(resourceGroupName string, VMScaleSetName string) (result VirtualMachineScaleSetExtensionListResult, err error) {
	req, err := client.ListPreparer(resourceGroupName, VMScaleSetName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetExtensionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetExtensionsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetExtensionsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client VirtualMachineScaleSetExtensionsClient) ListPreparer(resourceGroupName string, VMScaleSetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmScaleSetName":    autorest.Encode("path", VMScaleSetName),
	}

	const APIVersion = "2017-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/extensions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineScaleSetExtensionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VirtualMachineScaleSetExtensionsClient) ListResponder(resp *http.Response) (result VirtualMachineScaleSetExtensionListResult, err error) {
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
func (client VirtualMachineScaleSetExtensionsClient) ListNextResults(lastResults VirtualMachineScaleSetExtensionListResult) (result VirtualMachineScaleSetExtensionListResult, err error) {
	req, err := lastResults.VirtualMachineScaleSetExtensionListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetExtensionsClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetExtensionsClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetExtensionsClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListComplete gets all elements from the list without paging.
func (client VirtualMachineScaleSetExtensionsClient) ListComplete(resourceGroupName string, VMScaleSetName string, cancel <-chan struct{}) (<-chan VirtualMachineScaleSetExtension, <-chan error) {
	resultChan := make(chan VirtualMachineScaleSetExtension)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.List(resourceGroupName, VMScaleSetName)
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
