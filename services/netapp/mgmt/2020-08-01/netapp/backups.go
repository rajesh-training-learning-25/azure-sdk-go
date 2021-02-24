package netapp

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
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// BackupsClient is the microsoft NetApp Azure Resource Provider specification
type BackupsClient struct {
	BaseClient
}

// NewBackupsClient creates an instance of the BackupsClient client.
func NewBackupsClient(subscriptionID string) BackupsClient {
	return NewBackupsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewBackupsClientWithBaseURI creates an instance of the BackupsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewBackupsClientWithBaseURI(baseURI string, subscriptionID string) BackupsClient {
	return BackupsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create create a backup for the volume
// Parameters:
// resourceGroupName - the name of the resource group.
// accountName - the name of the NetApp account
// poolName - the name of the capacity pool
// volumeName - the name of the volume
// backupName - the name of the backup
// body - backup object supplied in the body of the operation.
func (client BackupsClient) Create(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, backupName string, body Backup) (result BackupsCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupsClient.Create")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: poolName,
			Constraints: []validation.Constraint{{Target: "poolName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "poolName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "poolName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9\-_]{0,63}$`, Chain: nil}}},
		{TargetValue: volumeName,
			Constraints: []validation.Constraint{{Target: "volumeName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "volumeName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "volumeName", Name: validation.Pattern, Rule: `^[a-zA-Z][a-zA-Z0-9\-_]{0,63}$`, Chain: nil}}},
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.Location", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "body.BackupProperties", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "body.BackupProperties.BackupID", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "body.BackupProperties.BackupID", Name: validation.MaxLength, Rule: 36, Chain: nil},
							{Target: "body.BackupProperties.BackupID", Name: validation.MinLength, Rule: 36, Chain: nil},
							{Target: "body.BackupProperties.BackupID", Name: validation.Pattern, Rule: `^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$`, Chain: nil},
						}},
					}}}}}); err != nil {
		return result, validation.NewError("netapp.BackupsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, accountName, poolName, volumeName, backupName, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netapp.BackupsClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netapp.BackupsClient", "Create", nil, "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client BackupsClient) CreatePreparer(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, backupName string, body Backup) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"backupName":        autorest.Encode("path", backupName),
		"poolName":          autorest.Encode("path", poolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"volumeName":        autorest.Encode("path", volumeName),
	}

	const APIVersion = "2020-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	body.ID = nil
	body.Name = nil
	body.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/backups/{backupName}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client BackupsClient) CreateSender(req *http.Request) (future BackupsCreateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client BackupsClient) (b Backup, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "netapp.BackupsCreateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("netapp.BackupsCreateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		b.Response.Response, err = future.GetResult(sender)
		if b.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "netapp.BackupsCreateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && b.Response.Response.StatusCode != http.StatusNoContent {
			b, err = client.CreateResponder(b.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "netapp.BackupsCreateFuture", "Result", b.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client BackupsClient) CreateResponder(resp *http.Response) (result Backup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a backup of the volume
// Parameters:
// resourceGroupName - the name of the resource group.
// accountName - the name of the NetApp account
// poolName - the name of the capacity pool
// volumeName - the name of the volume
// backupName - the name of the backup
func (client BackupsClient) Delete(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, backupName string) (result BackupsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: poolName,
			Constraints: []validation.Constraint{{Target: "poolName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "poolName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "poolName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9\-_]{0,63}$`, Chain: nil}}},
		{TargetValue: volumeName,
			Constraints: []validation.Constraint{{Target: "volumeName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "volumeName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "volumeName", Name: validation.Pattern, Rule: `^[a-zA-Z][a-zA-Z0-9\-_]{0,63}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("netapp.BackupsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, accountName, poolName, volumeName, backupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netapp.BackupsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netapp.BackupsClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client BackupsClient) DeletePreparer(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, backupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"backupName":        autorest.Encode("path", backupName),
		"poolName":          autorest.Encode("path", poolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"volumeName":        autorest.Encode("path", volumeName),
	}

	const APIVersion = "2020-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/backups/{backupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client BackupsClient) DeleteSender(req *http.Request) (future BackupsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client BackupsClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "netapp.BackupsDeleteFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("netapp.BackupsDeleteFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client BackupsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get a particular backup of the volume
// Parameters:
// resourceGroupName - the name of the resource group.
// accountName - the name of the NetApp account
// poolName - the name of the capacity pool
// volumeName - the name of the volume
// backupName - the name of the backup
func (client BackupsClient) Get(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, backupName string) (result Backup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: poolName,
			Constraints: []validation.Constraint{{Target: "poolName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "poolName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "poolName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9\-_]{0,63}$`, Chain: nil}}},
		{TargetValue: volumeName,
			Constraints: []validation.Constraint{{Target: "volumeName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "volumeName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "volumeName", Name: validation.Pattern, Rule: `^[a-zA-Z][a-zA-Z0-9\-_]{0,63}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("netapp.BackupsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, accountName, poolName, volumeName, backupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netapp.BackupsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "netapp.BackupsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netapp.BackupsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client BackupsClient) GetPreparer(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, backupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"backupName":        autorest.Encode("path", backupName),
		"poolName":          autorest.Encode("path", poolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"volumeName":        autorest.Encode("path", volumeName),
	}

	const APIVersion = "2020-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/backups/{backupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client BackupsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client BackupsClient) GetResponder(resp *http.Response) (result Backup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list all backups for a volume
// Parameters:
// resourceGroupName - the name of the resource group.
// accountName - the name of the NetApp account
// poolName - the name of the capacity pool
// volumeName - the name of the volume
func (client BackupsClient) List(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string) (result BackupsList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupsClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: poolName,
			Constraints: []validation.Constraint{{Target: "poolName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "poolName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "poolName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9\-_]{0,63}$`, Chain: nil}}},
		{TargetValue: volumeName,
			Constraints: []validation.Constraint{{Target: "volumeName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "volumeName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "volumeName", Name: validation.Pattern, Rule: `^[a-zA-Z][a-zA-Z0-9\-_]{0,63}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("netapp.BackupsClient", "List", err.Error())
	}

	req, err := client.ListPreparer(ctx, resourceGroupName, accountName, poolName, volumeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netapp.BackupsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "netapp.BackupsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netapp.BackupsClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client BackupsClient) ListPreparer(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"poolName":          autorest.Encode("path", poolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"volumeName":        autorest.Encode("path", volumeName),
	}

	const APIVersion = "2020-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/backups", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client BackupsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client BackupsClient) ListResponder(resp *http.Response) (result BackupsList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update patch a backup for the volume
// Parameters:
// resourceGroupName - the name of the resource group.
// accountName - the name of the NetApp account
// poolName - the name of the capacity pool
// volumeName - the name of the volume
// backupName - the name of the backup
// body - backup object supplied in the body of the operation.
func (client BackupsClient) Update(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, backupName string, body *BackupPatch) (result Backup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: poolName,
			Constraints: []validation.Constraint{{Target: "poolName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "poolName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "poolName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9\-_]{0,63}$`, Chain: nil}}},
		{TargetValue: volumeName,
			Constraints: []validation.Constraint{{Target: "volumeName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "volumeName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "volumeName", Name: validation.Pattern, Rule: `^[a-zA-Z][a-zA-Z0-9\-_]{0,63}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("netapp.BackupsClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, accountName, poolName, volumeName, backupName, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netapp.BackupsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "netapp.BackupsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netapp.BackupsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client BackupsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, backupName string, body *BackupPatch) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"backupName":        autorest.Encode("path", backupName),
		"poolName":          autorest.Encode("path", poolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"volumeName":        autorest.Encode("path", volumeName),
	}

	const APIVersion = "2020-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/backups/{backupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if body != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(body))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client BackupsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client BackupsClient) UpdateResponder(resp *http.Response) (result Backup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
