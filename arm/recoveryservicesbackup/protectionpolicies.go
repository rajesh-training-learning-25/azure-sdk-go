package recoveryservicesbackup

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

// ProtectionPoliciesClient is the open API 2.0 Specs for Azure RecoveryServices Backup service
type ProtectionPoliciesClient struct {
	ManagementClient
}

// NewProtectionPoliciesClient creates an instance of the ProtectionPoliciesClient client.
func NewProtectionPoliciesClient(subscriptionID string) ProtectionPoliciesClient {
	return NewProtectionPoliciesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProtectionPoliciesClientWithBaseURI creates an instance of the ProtectionPoliciesClient client.
func NewProtectionPoliciesClientWithBaseURI(baseURI string, subscriptionID string) ProtectionPoliciesClient {
	return ProtectionPoliciesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or modifies a backup policy. This is an asynchronous operation. Status of the operation can
// be fetched using GetPolicyOperationResult API.
//
// vaultName is the name of the recovery services vault. resourceGroupName is the name of the resource group where the
// recovery services vault is present. policyName is backup policy to be created. parameters is resource backup policy
func (client ProtectionPoliciesClient) CreateOrUpdate(vaultName string, resourceGroupName string, policyName string, parameters ProtectionPolicyResource) (result ProtectionPolicyResource, err error) {
	if err := validation.Validate([]validation.Validation{}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "recoveryservicesbackup.ProtectionPoliciesClient", "CreateOrUpdate")
	}

	req, err := client.CreateOrUpdatePreparer(vaultName, resourceGroupName, policyName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.ProtectionPoliciesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.ProtectionPoliciesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.ProtectionPoliciesClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ProtectionPoliciesClient) CreateOrUpdatePreparer(vaultName string, resourceGroupName string, policyName string, parameters ProtectionPolicyResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyName":        autorest.Encode("path", policyName),
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
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupPolicies/{policyName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ProtectionPoliciesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ProtectionPoliciesClient) CreateOrUpdateResponder(resp *http.Response) (result ProtectionPolicyResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes specified backup policy from your Recovery Services Vault. This is an asynchronous operation. Status
// of the operation can be fetched using GetPolicyOperationResult API.
//
// vaultName is the name of the recovery services vault. resourceGroupName is the name of the resource group where the
// recovery services vault is present. policyName is backup policy to be deleted.
func (client ProtectionPoliciesClient) Delete(vaultName string, resourceGroupName string, policyName string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "recoveryservicesbackup.ProtectionPoliciesClient", "Delete")
	}

	req, err := client.DeletePreparer(vaultName, resourceGroupName, policyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.ProtectionPoliciesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.ProtectionPoliciesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.ProtectionPoliciesClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ProtectionPoliciesClient) DeletePreparer(vaultName string, resourceGroupName string, policyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyName":        autorest.Encode("path", policyName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupPolicies/{policyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ProtectionPoliciesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ProtectionPoliciesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get provides the details of the backup policies associated to Recovery Services Vault. This is an asynchronous
// operation. Status of the operation can be fetched using GetPolicyOperationResult API.
//
// vaultName is the name of the recovery services vault. resourceGroupName is the name of the resource group where the
// recovery services vault is present. policyName is backup policy information to be fetched.
func (client ProtectionPoliciesClient) Get(vaultName string, resourceGroupName string, policyName string) (result ProtectionPolicyResource, err error) {
	if err := validation.Validate([]validation.Validation{}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "recoveryservicesbackup.ProtectionPoliciesClient", "Get")
	}

	req, err := client.GetPreparer(vaultName, resourceGroupName, policyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.ProtectionPoliciesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.ProtectionPoliciesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.ProtectionPoliciesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ProtectionPoliciesClient) GetPreparer(vaultName string, resourceGroupName string, policyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyName":        autorest.Encode("path", policyName),
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
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupPolicies/{policyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ProtectionPoliciesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ProtectionPoliciesClient) GetResponder(resp *http.Response) (result ProtectionPolicyResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
