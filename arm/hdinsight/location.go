package hdinsight

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

// LocationClient is the hDInsight Management Client
type LocationClient struct {
	ManagementClient
}

// NewLocationClient creates an instance of the LocationClient client.
func NewLocationClient(subscriptionID string) LocationClient {
	return NewLocationClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLocationClientWithBaseURI creates an instance of the LocationClient client.
func NewLocationClientWithBaseURI(baseURI string, subscriptionID string) LocationClient {
	return LocationClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// NewLocationClientWithAuthFile creates an instance of the LocationClient client.
func NewLocationClientWithAuthFile() (LocationClient, error) {
	c, err := NewWithAuthFile()
	return LocationClient{c}, err
}

// GetCapabilities gets the capabilities for the specified location.
//
// location is the location to get capabilities for.
func (client LocationClient) GetCapabilities(location string) (result CapabilitiesResult, err error) {
	req, err := client.GetCapabilitiesPreparer(location)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.LocationClient", "GetCapabilities", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetCapabilitiesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hdinsight.LocationClient", "GetCapabilities", resp, "Failure sending request")
		return
	}

	result, err = client.GetCapabilitiesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.LocationClient", "GetCapabilities", resp, "Failure responding to request")
	}

	return
}

// GetCapabilitiesPreparer prepares the GetCapabilities request.
func (client LocationClient) GetCapabilitiesPreparer(location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.HDInsight/locations/{location}/capabilities", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetCapabilitiesSender sends the GetCapabilities request. The method will close the
// http.Response Body if it receives an error.
func (client LocationClient) GetCapabilitiesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetCapabilitiesResponder handles the response to the GetCapabilities request. The method always
// closes the http.Response Body.
func (client LocationClient) GetCapabilitiesResponder(resp *http.Response) (result CapabilitiesResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
