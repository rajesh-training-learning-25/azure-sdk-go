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
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// BackupLongTermRetentionPoliciesClient is the the Azure SQL Database management API provides a RESTful set of web
// services that interact with Azure SQL Database services to manage your databases. The API enables you to create,
// retrieve, update, and delete databases.
type BackupLongTermRetentionPoliciesClient struct {
	ManagementClient
}

// NewBackupLongTermRetentionPoliciesClient creates an instance of the BackupLongTermRetentionPoliciesClient client.
func NewBackupLongTermRetentionPoliciesClient(subscriptionID string) BackupLongTermRetentionPoliciesClient {
	return NewBackupLongTermRetentionPoliciesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewBackupLongTermRetentionPoliciesClientWithBaseURI creates an instance of the BackupLongTermRetentionPoliciesClient
// client.
func NewBackupLongTermRetentionPoliciesClientWithBaseURI(baseURI string, subscriptionID string) BackupLongTermRetentionPoliciesClient {
	return BackupLongTermRetentionPoliciesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a database backup long term retention policy This method may poll for completion.
// Polling can be canceled by passing the cancel channel argument. The channel will be used to cancel polling and any
// outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from the
// Azure Resource Manager API or the portal. serverName is the name of the server. databaseName is the name of the
// database backupLongTermRetentionPolicyName is the name of the backup long term retention policy parameters is the
// required parameters to update a backup long term retention policy
func (client BackupLongTermRetentionPoliciesClient) CreateOrUpdate(resourceGroupName string, serverName string, databaseName string, backupLongTermRetentionPolicyName string, parameters BackupLongTermRetentionPolicy, cancel <-chan struct{}) (<-chan BackupLongTermRetentionPolicy, <-chan error) {
	resultChan := make(chan BackupLongTermRetentionPolicy, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.BackupLongTermRetentionPolicyProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.BackupLongTermRetentionPolicyProperties.RecoveryServicesBackupPolicyResourceID", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "sql.BackupLongTermRetentionPoliciesClient", "CreateOrUpdate")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result BackupLongTermRetentionPolicy
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreateOrUpdatePreparer(resourceGroupName, serverName, databaseName, backupLongTermRetentionPolicyName, parameters, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "sql.BackupLongTermRetentionPoliciesClient", "CreateOrUpdate", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateOrUpdateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "sql.BackupLongTermRetentionPoliciesClient", "CreateOrUpdate", resp, "Failure sending request")
			return
		}

		result, err = client.CreateOrUpdateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "sql.BackupLongTermRetentionPoliciesClient", "CreateOrUpdate", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client BackupLongTermRetentionPoliciesClient) CreateOrUpdatePreparer(resourceGroupName string, serverName string, databaseName string, backupLongTermRetentionPolicyName string, parameters BackupLongTermRetentionPolicy, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backupLongTermRetentionPolicyName": autorest.Encode("path", backupLongTermRetentionPolicyName),
		"databaseName":                      autorest.Encode("path", databaseName),
		"resourceGroupName":                 autorest.Encode("path", resourceGroupName),
		"serverName":                        autorest.Encode("path", serverName),
		"subscriptionId":                    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/backupLongTermRetentionPolicies/{backupLongTermRetentionPolicyName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client BackupLongTermRetentionPoliciesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client BackupLongTermRetentionPoliciesClient) CreateOrUpdateResponder(resp *http.Response) (result BackupLongTermRetentionPolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get returns a database backup long term retention policy
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from the
// Azure Resource Manager API or the portal. serverName is the name of the server. databaseName is the name of the
// database. backupLongTermRetentionPolicyName is the name of the backup long term retention policy
func (client BackupLongTermRetentionPoliciesClient) Get(resourceGroupName string, serverName string, databaseName string, backupLongTermRetentionPolicyName string) (result BackupLongTermRetentionPolicy, err error) {
	req, err := client.GetPreparer(resourceGroupName, serverName, databaseName, backupLongTermRetentionPolicyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.BackupLongTermRetentionPoliciesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.BackupLongTermRetentionPoliciesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.BackupLongTermRetentionPoliciesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client BackupLongTermRetentionPoliciesClient) GetPreparer(resourceGroupName string, serverName string, databaseName string, backupLongTermRetentionPolicyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backupLongTermRetentionPolicyName": autorest.Encode("path", backupLongTermRetentionPolicyName),
		"databaseName":                      autorest.Encode("path", databaseName),
		"resourceGroupName":                 autorest.Encode("path", resourceGroupName),
		"serverName":                        autorest.Encode("path", serverName),
		"subscriptionId":                    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/backupLongTermRetentionPolicies/{backupLongTermRetentionPolicyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client BackupLongTermRetentionPoliciesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client BackupLongTermRetentionPoliciesClient) GetResponder(resp *http.Response) (result BackupLongTermRetentionPolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
