package recoveryservices

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

// BackupVaultConfigsClient is the recovery Services Client
type BackupVaultConfigsClient struct {
	ManagementClient
}

// NewBackupVaultConfigsClient creates an instance of the BackupVaultConfigsClient client.
func NewBackupVaultConfigsClient(subscriptionID string) BackupVaultConfigsClient {
	return NewBackupVaultConfigsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewBackupVaultConfigsClientWithBaseURI creates an instance of the BackupVaultConfigsClient client.
func NewBackupVaultConfigsClientWithBaseURI(baseURI string, subscriptionID string) BackupVaultConfigsClient {
	return BackupVaultConfigsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get fetches vault config.
//
// resourceGroupName is the name of the resource group where the recovery services vault is present. vaultName is the
// name of the recovery services vault.
func (client BackupVaultConfigsClient) Get(resourceGroupName string, vaultName string) (result BackupVaultConfig, err error) {
	req, err := client.GetPreparer(resourceGroupName, vaultName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservices.BackupVaultConfigsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "recoveryservices.BackupVaultConfigsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservices.BackupVaultConfigsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client BackupVaultConfigsClient) GetPreparer(resourceGroupName string, vaultName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupconfig/vaultconfig", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client BackupVaultConfigsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client BackupVaultConfigsClient) GetResponder(resp *http.Response) (result BackupVaultConfig, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates vault config model type.
//
// resourceGroupName is the name of the resource group where the recovery services vault is present. vaultName is the
// name of the recovery services vault. backupVaultConfig is backup vault config.
func (client BackupVaultConfigsClient) Update(resourceGroupName string, vaultName string, backupVaultConfig BackupVaultConfig) (result BackupVaultConfig, err error) {
	req, err := client.UpdatePreparer(resourceGroupName, vaultName, backupVaultConfig)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservices.BackupVaultConfigsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "recoveryservices.BackupVaultConfigsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservices.BackupVaultConfigsClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client BackupVaultConfigsClient) UpdatePreparer(resourceGroupName string, vaultName string, backupVaultConfig BackupVaultConfig) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupconfig/vaultconfig", pathParameters),
		autorest.WithJSON(backupVaultConfig),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client BackupVaultConfigsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client BackupVaultConfigsClient) UpdateResponder(resp *http.Response) (result BackupVaultConfig, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
