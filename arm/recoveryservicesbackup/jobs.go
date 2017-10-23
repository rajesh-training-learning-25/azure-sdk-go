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
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// JobsClient is the open API 2.0 Specs for Azure RecoveryServices Backup service
type JobsClient struct {
	ManagementClient
}

// NewJobsClient creates an instance of the JobsClient client.
func NewJobsClient(subscriptionID string) JobsClient {
	return NewJobsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewJobsClientWithBaseURI creates an instance of the JobsClient client.
func NewJobsClientWithBaseURI(baseURI string, subscriptionID string) JobsClient {
	return JobsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// NewJobsClientWithAuthFile creates an instance of the JobsClient client.
func NewJobsClientWithAuthFile() (JobsClient, error) {
	c, err := NewWithAuthFile()
	return JobsClient{c}, err
}

// Export triggers export of jobs specified by filters and returns an OperationID to track.
//
// vaultName is the name of the recovery services vault. resourceGroupName is the name of the resource group where the
// recovery services vault is present. filter is oData filter options.
func (client JobsClient) Export(vaultName string, resourceGroupName string, filter string) (result autorest.Response, err error) {
	req, err := client.ExportPreparer(vaultName, resourceGroupName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.JobsClient", "Export", nil, "Failure preparing request")
		return
	}

	resp, err := client.ExportSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.JobsClient", "Export", resp, "Failure sending request")
		return
	}

	result, err = client.ExportResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.JobsClient", "Export", resp, "Failure responding to request")
	}

	return
}

// ExportPreparer prepares the Export request.
func (client JobsClient) ExportPreparer(vaultName string, resourceGroupName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupJobsExport", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ExportSender sends the Export request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) ExportSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ExportResponder handles the response to the Export request. The method always
// closes the http.Response Body.
func (client JobsClient) ExportResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
