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

// CapabilitiesClient is the the Azure SQL Database management API provides a RESTful set of web services that interact
// with Azure SQL Database services to manage your databases. The API enables you to create, retrieve, update, and
// delete databases.
type CapabilitiesClient struct {
	ManagementClient
}

// NewCapabilitiesClient creates an instance of the CapabilitiesClient client.
func NewCapabilitiesClient(subscriptionID string) CapabilitiesClient {
	return NewCapabilitiesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCapabilitiesClientWithBaseURI creates an instance of the CapabilitiesClient client.
func NewCapabilitiesClientWithBaseURI(baseURI string, subscriptionID string) CapabilitiesClient {
	return CapabilitiesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListByLocation gets the capabilities available for the specified location.
//
// locationID is the location id whose capabilities are retrieved.
func (client CapabilitiesClient) ListByLocation(locationID string) (result LocationCapabilities, err error) {
	req, err := client.ListByLocationPreparer(locationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.CapabilitiesClient", "ListByLocation", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByLocationSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.CapabilitiesClient", "ListByLocation", resp, "Failure sending request")
		return
	}

	result, err = client.ListByLocationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.CapabilitiesClient", "ListByLocation", resp, "Failure responding to request")
	}

	return
}

// ListByLocationPreparer prepares the ListByLocation request.
func (client CapabilitiesClient) ListByLocationPreparer(locationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"locationId":     autorest.Encode("path", locationID),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationId}/capabilities", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByLocationSender sends the ListByLocation request. The method will close the
// http.Response Body if it receives an error.
func (client CapabilitiesClient) ListByLocationSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByLocationResponder handles the response to the ListByLocation request. The method always
// closes the http.Response Body.
func (client CapabilitiesClient) ListByLocationResponder(resp *http.Response) (result LocationCapabilities, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
