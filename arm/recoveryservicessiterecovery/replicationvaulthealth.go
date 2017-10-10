package recoveryservicessiterecovery

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

// ReplicationVaultHealthClient is the client for the ReplicationVaultHealth methods of the
// Recoveryservicessiterecovery service.
type ReplicationVaultHealthClient struct {
	ManagementClient
}

// NewReplicationVaultHealthClient creates an instance of the ReplicationVaultHealthClient client.
func NewReplicationVaultHealthClient(subscriptionID string, resourceGroupName string, resourceName string) ReplicationVaultHealthClient {
	return NewReplicationVaultHealthClientWithBaseURI(DefaultBaseURI, subscriptionID, resourceGroupName, resourceName)
}

// NewReplicationVaultHealthClientWithBaseURI creates an instance of the ReplicationVaultHealthClient client.
func NewReplicationVaultHealthClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string, resourceName string) ReplicationVaultHealthClient {
	return ReplicationVaultHealthClient{NewWithBaseURI(baseURI, subscriptionID, resourceGroupName, resourceName)}
}

// NewReplicationVaultHealthClientWithAuthFile creates an instance of the ReplicationVaultHealthClient client.
func NewReplicationVaultHealthClientWithAuthFile(resourceGroupName string, resourceName string) (ReplicationVaultHealthClient, error) {
	c, err := NewWithAuthFile(resourceGroupName, resourceName)
	return ReplicationVaultHealthClient{c}, err
}

// Get gets the health details of the vault.
func (client ReplicationVaultHealthClient) Get() (result VaultHealthDetails, err error) {
	req, err := client.GetPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationVaultHealthClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationVaultHealthClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationVaultHealthClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ReplicationVaultHealthClient) GetPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationVaultHealth", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationVaultHealthClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ReplicationVaultHealthClient) GetResponder(resp *http.Response) (result VaultHealthDetails, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
