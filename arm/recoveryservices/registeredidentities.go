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
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// RegisteredIdentitiesClient is the recovery Services Client
type RegisteredIdentitiesClient struct {
	ManagementClient
}

// NewRegisteredIdentitiesClient creates an instance of the RegisteredIdentitiesClient client.
func NewRegisteredIdentitiesClient(subscriptionID string) RegisteredIdentitiesClient {
	return NewRegisteredIdentitiesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRegisteredIdentitiesClientWithBaseURI creates an instance of the RegisteredIdentitiesClient client.
func NewRegisteredIdentitiesClientWithBaseURI(baseURI string, subscriptionID string) RegisteredIdentitiesClient {
	return RegisteredIdentitiesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// NewRegisteredIdentitiesClientWithAuthFile creates an instance of the RegisteredIdentitiesClient client.
func NewRegisteredIdentitiesClientWithAuthFile() (RegisteredIdentitiesClient, error) {
	c, err := NewWithAuthFile()
	return RegisteredIdentitiesClient{c}, err
}

// Delete unregisters the given container from your Recovery Services vault.
//
// resourceGroupName is the name of the resource group where the recovery services vault is present. vaultName is the
// name of the recovery services vault. identityName is name of the protection container to unregister.
func (client RegisteredIdentitiesClient) Delete(resourceGroupName string, vaultName string, identityName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(resourceGroupName, vaultName, identityName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservices.RegisteredIdentitiesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "recoveryservices.RegisteredIdentitiesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservices.RegisteredIdentitiesClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client RegisteredIdentitiesClient) DeletePreparer(resourceGroupName string, vaultName string, identityName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"identityName":      autorest.Encode("path", identityName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/registeredIdentities/{identityName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client RegisteredIdentitiesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client RegisteredIdentitiesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
