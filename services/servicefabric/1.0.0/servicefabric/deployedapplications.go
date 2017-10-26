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

// DeployedApplicationsClient is the client for the DeployedApplications methods of the Servicefabric service.
type DeployedApplicationsClient struct {
	ManagementClient
}

// NewDeployedApplicationsClient creates an instance of the DeployedApplicationsClient client.
func NewDeployedApplicationsClient(timeout *int32) DeployedApplicationsClient {
	return NewDeployedApplicationsClientWithBaseURI(DefaultBaseURI, timeout)
}

// NewDeployedApplicationsClientWithBaseURI creates an instance of the DeployedApplicationsClient client.
func NewDeployedApplicationsClientWithBaseURI(baseURI string, timeout *int32) DeployedApplicationsClient {
	return DeployedApplicationsClient{NewWithBaseURI(baseURI, timeout)}
}

// Get get deployed applications
//
// nodeName is the name of the node applicationName is the name of the application
func (client DeployedApplicationsClient) Get(nodeName string, applicationName string) (result DeployedApplication, err error) {
	req, err := client.GetPreparer(nodeName, applicationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.DeployedApplicationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.DeployedApplicationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.DeployedApplicationsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client DeployedApplicationsClient) GetPreparer(nodeName string, applicationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName": autorest.Encode("path", applicationName),
		"nodeName":        autorest.Encode("path", nodeName),
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
		autorest.WithPathParameters("/Nodes/{nodeName}/$/GetApplications/{applicationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DeployedApplicationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DeployedApplicationsClient) GetResponder(resp *http.Response) (result DeployedApplication, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list deployed applications
//
// nodeName is the name of the node
func (client DeployedApplicationsClient) List(nodeName string) (result ListDeployedApplication, err error) {
	req, err := client.ListPreparer(nodeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.DeployedApplicationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.DeployedApplicationsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.DeployedApplicationsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client DeployedApplicationsClient) ListPreparer(nodeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"nodeName": autorest.Encode("path", nodeName),
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
		autorest.WithPathParameters("/Nodes/{nodeName}/$/GetApplications", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DeployedApplicationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DeployedApplicationsClient) ListResponder(resp *http.Response) (result ListDeployedApplication, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
