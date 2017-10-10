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
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// VirtualMachineRunCommandsClient is the compute Client
type VirtualMachineRunCommandsClient struct {
	ManagementClient
}

// NewVirtualMachineRunCommandsClient creates an instance of the VirtualMachineRunCommandsClient client.
func NewVirtualMachineRunCommandsClient(subscriptionID string) VirtualMachineRunCommandsClient {
	return NewVirtualMachineRunCommandsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVirtualMachineRunCommandsClientWithBaseURI creates an instance of the VirtualMachineRunCommandsClient client.
func NewVirtualMachineRunCommandsClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineRunCommandsClient {
	return VirtualMachineRunCommandsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// NewVirtualMachineRunCommandsClientWithAuthFile creates an instance of the VirtualMachineRunCommandsClient
// client.
func NewVirtualMachineRunCommandsClientWithAuthFile() (VirtualMachineRunCommandsClient, error) {
	c, err := NewWithAuthFile()
	return VirtualMachineRunCommandsClient{c}, err
}

// Get gets specific run command for a subscription in a location.
//
// location is the location upon which run commands is queried. commandID is the command id.
func (client VirtualMachineRunCommandsClient) Get(location string, commandID string) (result RunCommandDocument, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: location,
			Constraints: []validation.Constraint{{Target: "location", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "compute.VirtualMachineRunCommandsClient", "Get")
	}

	req, err := client.GetPreparer(location, commandID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineRunCommandsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineRunCommandsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineRunCommandsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client VirtualMachineRunCommandsClient) GetPreparer(location string, commandID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"commandId":      autorest.Encode("path", commandID),
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/runCommands/{commandId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineRunCommandsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VirtualMachineRunCommandsClient) GetResponder(resp *http.Response) (result RunCommandDocument, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all available run commands for a subscription in a location.
//
// location is the location upon which run commands is queried.
func (client VirtualMachineRunCommandsClient) List(location string) (result RunCommandListResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: location,
			Constraints: []validation.Constraint{{Target: "location", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "compute.VirtualMachineRunCommandsClient", "List")
	}

	req, err := client.ListPreparer(location)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineRunCommandsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineRunCommandsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineRunCommandsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client VirtualMachineRunCommandsClient) ListPreparer(location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/runCommands", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineRunCommandsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VirtualMachineRunCommandsClient) ListResponder(resp *http.Response) (result RunCommandListResult, err error) {
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
func (client VirtualMachineRunCommandsClient) ListNextResults(lastResults RunCommandListResult) (result RunCommandListResult, err error) {
	req, err := lastResults.RunCommandListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachineRunCommandsClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachineRunCommandsClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineRunCommandsClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListComplete gets all elements from the list without paging.
func (client VirtualMachineRunCommandsClient) ListComplete(location string, cancel <-chan struct{}) (<-chan RunCommandDocumentBase, <-chan error) {
	resultChan := make(chan RunCommandDocumentBase)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.List(location)
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
