package servermanagement

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

// PowerShellClient is the REST API for Azure Server Management Service.
type PowerShellClient struct {
	ManagementClient
}

// NewPowerShellClient creates an instance of the PowerShellClient client.
func NewPowerShellClient(subscriptionID string) PowerShellClient {
	return NewPowerShellClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPowerShellClientWithBaseURI creates an instance of the PowerShellClient client.
func NewPowerShellClientWithBaseURI(baseURI string, subscriptionID string) PowerShellClient {
	return PowerShellClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// NewPowerShellClientWithAuthFile creates an instance of the PowerShellClient client.
func NewPowerShellClientWithAuthFile() (PowerShellClient, error) {
	c, err := NewWithAuthFile()
	return PowerShellClient{c}, err
}

// CancelCommand cancels a PowerShell command. This method may poll for completion. Polling can be canceled by passing
// the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the resource group name uniquely identifies the resource group within the user subscriptionId.
// nodeName is the node name (256 characters maximum). session is the sessionId from the user. pssession is the
// PowerShell sessionId from the user.
func (client PowerShellClient) CancelCommand(resourceGroupName string, nodeName string, session string, pssession string, cancel <-chan struct{}) (<-chan PowerShellCommandResults, <-chan error) {
	resultChan := make(chan PowerShellCommandResults, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9]+`, Chain: nil}}},
		{TargetValue: nodeName,
			Constraints: []validation.Constraint{{Target: "nodeName", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "nodeName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "nodeName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "servermanagement.PowerShellClient", "CancelCommand")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result PowerShellCommandResults
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CancelCommandPreparer(resourceGroupName, nodeName, session, pssession, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "CancelCommand", nil, "Failure preparing request")
			return
		}

		resp, err := client.CancelCommandSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "CancelCommand", resp, "Failure sending request")
			return
		}

		result, err = client.CancelCommandResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "CancelCommand", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CancelCommandPreparer prepares the CancelCommand request.
func (client PowerShellClient) CancelCommandPreparer(resourceGroupName string, nodeName string, session string, pssession string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"nodeName":          autorest.Encode("path", nodeName),
		"pssession":         autorest.Encode("path", pssession),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"session":           autorest.Encode("path", session),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServerManagement/nodes/{nodeName}/sessions/{session}/features/powerShellConsole/pssessions/{pssession}/cancel", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CancelCommandSender sends the CancelCommand request. The method will close the
// http.Response Body if it receives an error.
func (client PowerShellClient) CancelCommandSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CancelCommandResponder handles the response to the CancelCommand request. The method always
// closes the http.Response Body.
func (client PowerShellClient) CancelCommandResponder(resp *http.Response) (result PowerShellCommandResults, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateSession creates a PowerShell session. This method may poll for completion. Polling can be canceled by passing
// the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the resource group name uniquely identifies the resource group within the user subscriptionId.
// nodeName is the node name (256 characters maximum). session is the sessionId from the user. pssession is the
// PowerShell sessionId from the user.
func (client PowerShellClient) CreateSession(resourceGroupName string, nodeName string, session string, pssession string, cancel <-chan struct{}) (<-chan PowerShellSessionResource, <-chan error) {
	resultChan := make(chan PowerShellSessionResource, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9]+`, Chain: nil}}},
		{TargetValue: nodeName,
			Constraints: []validation.Constraint{{Target: "nodeName", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "nodeName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "nodeName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "servermanagement.PowerShellClient", "CreateSession")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result PowerShellSessionResource
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreateSessionPreparer(resourceGroupName, nodeName, session, pssession, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "CreateSession", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateSessionSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "CreateSession", resp, "Failure sending request")
			return
		}

		result, err = client.CreateSessionResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "CreateSession", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreateSessionPreparer prepares the CreateSession request.
func (client PowerShellClient) CreateSessionPreparer(resourceGroupName string, nodeName string, session string, pssession string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"nodeName":          autorest.Encode("path", nodeName),
		"pssession":         autorest.Encode("path", pssession),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"session":           autorest.Encode("path", session),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServerManagement/nodes/{nodeName}/sessions/{session}/features/powerShellConsole/pssessions/{pssession}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateSessionSender sends the CreateSession request. The method will close the
// http.Response Body if it receives an error.
func (client PowerShellClient) CreateSessionSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateSessionResponder handles the response to the CreateSession request. The method always
// closes the http.Response Body.
func (client PowerShellClient) CreateSessionResponder(resp *http.Response) (result PowerShellSessionResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetCommandStatus gets the status of a command.
//
// resourceGroupName is the resource group name uniquely identifies the resource group within the user subscriptionId.
// nodeName is the node name (256 characters maximum). session is the sessionId from the user. pssession is the
// PowerShell sessionId from the user. expand is gets current output from an ongoing call.
func (client PowerShellClient) GetCommandStatus(resourceGroupName string, nodeName string, session string, pssession string, expand PowerShellExpandOption) (result PowerShellCommandStatus, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9]+`, Chain: nil}}},
		{TargetValue: nodeName,
			Constraints: []validation.Constraint{{Target: "nodeName", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "nodeName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "nodeName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "servermanagement.PowerShellClient", "GetCommandStatus")
	}

	req, err := client.GetCommandStatusPreparer(resourceGroupName, nodeName, session, pssession, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "GetCommandStatus", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetCommandStatusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "GetCommandStatus", resp, "Failure sending request")
		return
	}

	result, err = client.GetCommandStatusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "GetCommandStatus", resp, "Failure responding to request")
	}

	return
}

// GetCommandStatusPreparer prepares the GetCommandStatus request.
func (client PowerShellClient) GetCommandStatusPreparer(resourceGroupName string, nodeName string, session string, pssession string, expand PowerShellExpandOption) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"nodeName":          autorest.Encode("path", nodeName),
		"pssession":         autorest.Encode("path", pssession),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"session":           autorest.Encode("path", session),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(string(expand)) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServerManagement/nodes/{nodeName}/sessions/{session}/features/powerShellConsole/pssessions/{pssession}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetCommandStatusSender sends the GetCommandStatus request. The method will close the
// http.Response Body if it receives an error.
func (client PowerShellClient) GetCommandStatusSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetCommandStatusResponder handles the response to the GetCommandStatus request. The method always
// closes the http.Response Body.
func (client PowerShellClient) GetCommandStatusResponder(resp *http.Response) (result PowerShellCommandStatus, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// InvokeCommand creates a PowerShell script and invokes it. This method may poll for completion. Polling can be
// canceled by passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP
// requests.
//
// resourceGroupName is the resource group name uniquely identifies the resource group within the user subscriptionId.
// nodeName is the node name (256 characters maximum). session is the sessionId from the user. pssession is the
// PowerShell sessionId from the user. powerShellCommandParameters is parameters supplied to the Invoke PowerShell
// Command operation.
func (client PowerShellClient) InvokeCommand(resourceGroupName string, nodeName string, session string, pssession string, powerShellCommandParameters PowerShellCommandParameters, cancel <-chan struct{}) (<-chan PowerShellCommandResults, <-chan error) {
	resultChan := make(chan PowerShellCommandResults, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9]+`, Chain: nil}}},
		{TargetValue: nodeName,
			Constraints: []validation.Constraint{{Target: "nodeName", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "nodeName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "nodeName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "servermanagement.PowerShellClient", "InvokeCommand")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result PowerShellCommandResults
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.InvokeCommandPreparer(resourceGroupName, nodeName, session, pssession, powerShellCommandParameters, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "InvokeCommand", nil, "Failure preparing request")
			return
		}

		resp, err := client.InvokeCommandSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "InvokeCommand", resp, "Failure sending request")
			return
		}

		result, err = client.InvokeCommandResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "InvokeCommand", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// InvokeCommandPreparer prepares the InvokeCommand request.
func (client PowerShellClient) InvokeCommandPreparer(resourceGroupName string, nodeName string, session string, pssession string, powerShellCommandParameters PowerShellCommandParameters, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"nodeName":          autorest.Encode("path", nodeName),
		"pssession":         autorest.Encode("path", pssession),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"session":           autorest.Encode("path", session),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServerManagement/nodes/{nodeName}/sessions/{session}/features/powerShellConsole/pssessions/{pssession}/invokeCommand", pathParameters),
		autorest.WithJSON(powerShellCommandParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// InvokeCommandSender sends the InvokeCommand request. The method will close the
// http.Response Body if it receives an error.
func (client PowerShellClient) InvokeCommandSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// InvokeCommandResponder handles the response to the InvokeCommand request. The method always
// closes the http.Response Body.
func (client PowerShellClient) InvokeCommandResponder(resp *http.Response) (result PowerShellCommandResults, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListSession gets a list of the active sessions.
//
// resourceGroupName is the resource group name uniquely identifies the resource group within the user subscriptionId.
// nodeName is the node name (256 characters maximum). session is the sessionId from the user.
func (client PowerShellClient) ListSession(resourceGroupName string, nodeName string, session string) (result PowerShellSessionResources, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9]+`, Chain: nil}}},
		{TargetValue: nodeName,
			Constraints: []validation.Constraint{{Target: "nodeName", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "nodeName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "nodeName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "servermanagement.PowerShellClient", "ListSession")
	}

	req, err := client.ListSessionPreparer(resourceGroupName, nodeName, session)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "ListSession", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSessionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "ListSession", resp, "Failure sending request")
		return
	}

	result, err = client.ListSessionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "ListSession", resp, "Failure responding to request")
	}

	return
}

// ListSessionPreparer prepares the ListSession request.
func (client PowerShellClient) ListSessionPreparer(resourceGroupName string, nodeName string, session string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"nodeName":          autorest.Encode("path", nodeName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"session":           autorest.Encode("path", session),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServerManagement/nodes/{nodeName}/sessions/{session}/features/powerShellConsole/pssessions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSessionSender sends the ListSession request. The method will close the
// http.Response Body if it receives an error.
func (client PowerShellClient) ListSessionSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListSessionResponder handles the response to the ListSession request. The method always
// closes the http.Response Body.
func (client PowerShellClient) ListSessionResponder(resp *http.Response) (result PowerShellSessionResources, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// TabCompletion gets tab completion values for a command.
//
// resourceGroupName is the resource group name uniquely identifies the resource group within the user subscriptionId.
// nodeName is the node name (256 characters maximum). session is the sessionId from the user. pssession is the
// PowerShell sessionId from the user. powerShellTabCompletionParamters is parameters supplied to the tab completion
// call.
func (client PowerShellClient) TabCompletion(resourceGroupName string, nodeName string, session string, pssession string, powerShellTabCompletionParamters PowerShellTabCompletionParameters) (result PowerShellTabCompletionResults, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9]+`, Chain: nil}}},
		{TargetValue: nodeName,
			Constraints: []validation.Constraint{{Target: "nodeName", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "nodeName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "nodeName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "servermanagement.PowerShellClient", "TabCompletion")
	}

	req, err := client.TabCompletionPreparer(resourceGroupName, nodeName, session, pssession, powerShellTabCompletionParamters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "TabCompletion", nil, "Failure preparing request")
		return
	}

	resp, err := client.TabCompletionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "TabCompletion", resp, "Failure sending request")
		return
	}

	result, err = client.TabCompletionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "TabCompletion", resp, "Failure responding to request")
	}

	return
}

// TabCompletionPreparer prepares the TabCompletion request.
func (client PowerShellClient) TabCompletionPreparer(resourceGroupName string, nodeName string, session string, pssession string, powerShellTabCompletionParamters PowerShellTabCompletionParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"nodeName":          autorest.Encode("path", nodeName),
		"pssession":         autorest.Encode("path", pssession),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"session":           autorest.Encode("path", session),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServerManagement/nodes/{nodeName}/sessions/{session}/features/powerShellConsole/pssessions/{pssession}/tab", pathParameters),
		autorest.WithJSON(powerShellTabCompletionParamters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// TabCompletionSender sends the TabCompletion request. The method will close the
// http.Response Body if it receives an error.
func (client PowerShellClient) TabCompletionSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// TabCompletionResponder handles the response to the TabCompletion request. The method always
// closes the http.Response Body.
func (client PowerShellClient) TabCompletionResponder(resp *http.Response) (result PowerShellTabCompletionResults, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UpdateCommand updates a running PowerShell command with more data. This method may poll for completion. Polling can
// be canceled by passing the cancel channel argument. The channel will be used to cancel polling and any outstanding
// HTTP requests.
//
// resourceGroupName is the resource group name uniquely identifies the resource group within the user subscriptionId.
// nodeName is the node name (256 characters maximum). session is the sessionId from the user. pssession is the
// PowerShell sessionId from the user.
func (client PowerShellClient) UpdateCommand(resourceGroupName string, nodeName string, session string, pssession string, cancel <-chan struct{}) (<-chan PowerShellCommandResults, <-chan error) {
	resultChan := make(chan PowerShellCommandResults, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9]+`, Chain: nil}}},
		{TargetValue: nodeName,
			Constraints: []validation.Constraint{{Target: "nodeName", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "nodeName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "nodeName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "servermanagement.PowerShellClient", "UpdateCommand")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result PowerShellCommandResults
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.UpdateCommandPreparer(resourceGroupName, nodeName, session, pssession, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "UpdateCommand", nil, "Failure preparing request")
			return
		}

		resp, err := client.UpdateCommandSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "UpdateCommand", resp, "Failure sending request")
			return
		}

		result, err = client.UpdateCommandResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "servermanagement.PowerShellClient", "UpdateCommand", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// UpdateCommandPreparer prepares the UpdateCommand request.
func (client PowerShellClient) UpdateCommandPreparer(resourceGroupName string, nodeName string, session string, pssession string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"nodeName":          autorest.Encode("path", nodeName),
		"pssession":         autorest.Encode("path", pssession),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"session":           autorest.Encode("path", session),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServerManagement/nodes/{nodeName}/sessions/{session}/features/powerShellConsole/pssessions/{pssession}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// UpdateCommandSender sends the UpdateCommand request. The method will close the
// http.Response Body if it receives an error.
func (client PowerShellClient) UpdateCommandSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// UpdateCommandResponder handles the response to the UpdateCommand request. The method always
// closes the http.Response Body.
func (client PowerShellClient) UpdateCommandResponder(resp *http.Response) (result PowerShellCommandResults, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
