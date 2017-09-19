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
// Code generated by Microsoft (R) AutoRest Code Generator 2.2.21.0
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// GeoBackupPoliciesClient is the the Azure SQL Database management API provides a RESTful set of web services that
// interact with Azure SQL Database services to manage your databases. The API enables you to create, retrieve, update,
// and delete databases.
type GeoBackupPoliciesClient struct {
	ManagementClient
}

// NewGeoBackupPoliciesClient creates an instance of the GeoBackupPoliciesClient client.
func NewGeoBackupPoliciesClient(subscriptionID string) GeoBackupPoliciesClient {
	return NewGeoBackupPoliciesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewGeoBackupPoliciesClientWithBaseURI creates an instance of the GeoBackupPoliciesClient client.
func NewGeoBackupPoliciesClientWithBaseURI(baseURI string, subscriptionID string) GeoBackupPoliciesClient {
	return GeoBackupPoliciesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate updates a database geo backup policy.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from the
// Azure Resource Manager API or the portal. serverName is the name of the server. databaseName is the name of the
// database. geoBackupPolicyName is the name of the geo backup policy. parameters is the required parameters for
// creating or updating the geo backup policy.
func (client GeoBackupPoliciesClient) CreateOrUpdate(resourceGroupName string, serverName string, databaseName string, geoBackupPolicyName string, parameters GeoBackupPolicy) (result GeoBackupPolicy, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.GeoBackupPolicyProperties", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "sql.GeoBackupPoliciesClient", "CreateOrUpdate")
	}

	req, err := client.CreateOrUpdatePreparer(resourceGroupName, serverName, databaseName, geoBackupPolicyName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.GeoBackupPoliciesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.GeoBackupPoliciesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.GeoBackupPoliciesClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client GeoBackupPoliciesClient) CreateOrUpdatePreparer(resourceGroupName string, serverName string, databaseName string, geoBackupPolicyName string, parameters GeoBackupPolicy) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":        autorest.Encode("path", databaseName),
		"geoBackupPolicyName": autorest.Encode("path", geoBackupPolicyName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"serverName":          autorest.Encode("path", serverName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/geoBackupPolicies/{geoBackupPolicyName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client GeoBackupPoliciesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client GeoBackupPoliciesClient) CreateOrUpdateResponder(resp *http.Response) (result GeoBackupPolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets a geo backup policy.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from the
// Azure Resource Manager API or the portal. serverName is the name of the server. databaseName is the name of the
// database. geoBackupPolicyName is the name of the geo backup policy.
func (client GeoBackupPoliciesClient) Get(resourceGroupName string, serverName string, databaseName string, geoBackupPolicyName string) (result GeoBackupPolicy, err error) {
	if err := validation.Validate([]validation.Validation{}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "sql.GeoBackupPoliciesClient", "Get")
	}

	req, err := client.GetPreparer(resourceGroupName, serverName, databaseName, geoBackupPolicyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.GeoBackupPoliciesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.GeoBackupPoliciesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.GeoBackupPoliciesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client GeoBackupPoliciesClient) GetPreparer(resourceGroupName string, serverName string, databaseName string, geoBackupPolicyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":        autorest.Encode("path", databaseName),
		"geoBackupPolicyName": autorest.Encode("path", geoBackupPolicyName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"serverName":          autorest.Encode("path", serverName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/geoBackupPolicies/{geoBackupPolicyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client GeoBackupPoliciesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client GeoBackupPoliciesClient) GetResponder(resp *http.Response) (result GeoBackupPolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByDatabase returns a list of geo backup policies.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from the
// Azure Resource Manager API or the portal. serverName is the name of the server. databaseName is the name of the
// database.
func (client GeoBackupPoliciesClient) ListByDatabase(resourceGroupName string, serverName string, databaseName string) (result GeoBackupPolicyListResult, err error) {
	if err := validation.Validate([]validation.Validation{}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "sql.GeoBackupPoliciesClient", "ListByDatabase")
	}

	req, err := client.ListByDatabasePreparer(resourceGroupName, serverName, databaseName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.GeoBackupPoliciesClient", "ListByDatabase", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByDatabaseSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.GeoBackupPoliciesClient", "ListByDatabase", resp, "Failure sending request")
		return
	}

	result, err = client.ListByDatabaseResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.GeoBackupPoliciesClient", "ListByDatabase", resp, "Failure responding to request")
	}

	return
}

// ListByDatabasePreparer prepares the ListByDatabase request.
func (client GeoBackupPoliciesClient) ListByDatabasePreparer(resourceGroupName string, serverName string, databaseName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/geoBackupPolicies", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByDatabaseSender sends the ListByDatabase request. The method will close the
// http.Response Body if it receives an error.
func (client GeoBackupPoliciesClient) ListByDatabaseSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByDatabaseResponder handles the response to the ListByDatabase request. The method always
// closes the http.Response Body.
func (client GeoBackupPoliciesClient) ListByDatabaseResponder(resp *http.Response) (result GeoBackupPolicyListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
