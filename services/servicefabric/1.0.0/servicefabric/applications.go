package servicefabric

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

// ApplicationsClient is the client for the Applications methods of the Servicefabric service.
type ApplicationsClient struct {
	ManagementClient
}

// NewApplicationsClient creates an instance of the ApplicationsClient client.
func NewApplicationsClient(timeout *int32) ApplicationsClient {
	return NewApplicationsClientWithBaseURI(DefaultBaseURI, timeout)
}

// NewApplicationsClientWithBaseURI creates an instance of the ApplicationsClient client.
func NewApplicationsClientWithBaseURI(baseURI string, timeout *int32) ApplicationsClient {
	return ApplicationsClient{NewWithBaseURI(baseURI, timeout)}
}

// Create create applications
//
// applicationDescription is the description of the application
func (client ApplicationsClient) Create(applicationDescription ApplicationDescription) (result String, err error) {
	req, err := client.CreatePreparer(applicationDescription)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationsClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ApplicationsClient) CreatePreparer(applicationDescription ApplicationDescription) (*http.Request, error) {
	const APIVersion = "1.0.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if client.Timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *client.Timeout)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/Applications/$/Create"),
		autorest.WithJSON(applicationDescription),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) CreateResponder(resp *http.Response) (result String, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get applications
//
// applicationName is the name of the application
func (client ApplicationsClient) Get(applicationName string) (result Application, err error) {
	req, err := client.GetPreparer(applicationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ApplicationsClient) GetPreparer(applicationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName": applicationName,
	}

	const APIVersion = "1.0.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if client.Timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *client.Timeout)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Applications/{applicationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) GetResponder(resp *http.Response) (result Application, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list applications
//
// continuationToken is the token of the continuation
func (client ApplicationsClient) List(continuationToken string) (result ApplicationList, err error) {
	req, err := client.ListPreparer(continuationToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ApplicationsClient) ListPreparer(continuationToken string) (*http.Request, error) {
	const APIVersion = "1.0.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if client.Timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *client.Timeout)
	}
	if len(continuationToken) > 0 {
		queryParameters["continuation-token"] = autorest.Encode("query", continuationToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/Applications"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) ListResponder(resp *http.Response) (result ApplicationList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Remove remove applications
//
// applicationName is the name of the application forceRemove is the force remove flag to skip services check
func (client ApplicationsClient) Remove(applicationName string, forceRemove *bool) (result String, err error) {
	req, err := client.RemovePreparer(applicationName, forceRemove)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationsClient", "Remove", nil, "Failure preparing request")
		return
	}

	resp, err := client.RemoveSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationsClient", "Remove", resp, "Failure sending request")
		return
	}

	result, err = client.RemoveResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationsClient", "Remove", resp, "Failure responding to request")
	}

	return
}

// RemovePreparer prepares the Remove request.
func (client ApplicationsClient) RemovePreparer(applicationName string, forceRemove *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName": applicationName,
	}

	const APIVersion = "1.0.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if forceRemove != nil {
		queryParameters["ForceRemove"] = autorest.Encode("query", *forceRemove)
	}
	if client.Timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *client.Timeout)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Applications/{applicationName}/$/Delete", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// RemoveSender sends the Remove request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) RemoveSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// RemoveResponder handles the response to the Remove request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) RemoveResponder(resp *http.Response) (result String, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
