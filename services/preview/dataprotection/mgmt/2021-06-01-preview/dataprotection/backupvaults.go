package dataprotection

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
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

// BackupVaultsClient is the open API 2.0 Specs for Azure Data Protection service
type BackupVaultsClient struct {
	BaseClient
}

// NewBackupVaultsClient creates an instance of the BackupVaultsClient client.
func NewBackupVaultsClient(subscriptionID string) BackupVaultsClient {
	return NewBackupVaultsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewBackupVaultsClientWithBaseURI creates an instance of the BackupVaultsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewBackupVaultsClientWithBaseURI(baseURI string, subscriptionID string) BackupVaultsClient {
	return BackupVaultsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CheckNameAvailability sends the check name availability request.
// Parameters:
// resourceGroupName - the name of the resource group where the backup vault is present.
// location - the location in which uniqueness will be verified.
// parameters - check name availability request
func (client BackupVaultsClient) CheckNameAvailability(ctx context.Context, resourceGroupName string, location string, parameters CheckNameAvailabilityRequest) (result CheckNameAvailabilityResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupVaultsClient.CheckNameAvailability")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CheckNameAvailabilityPreparer(ctx, resourceGroupName, location, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BackupVaultsClient", "CheckNameAvailability", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckNameAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dataprotection.BackupVaultsClient", "CheckNameAvailability", resp, "Failure sending request")
		return
	}

	result, err = client.CheckNameAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BackupVaultsClient", "CheckNameAvailability", resp, "Failure responding to request")
		return
	}

	return
}

// CheckNameAvailabilityPreparer prepares the CheckNameAvailability request.
func (client BackupVaultsClient) CheckNameAvailabilityPreparer(ctx context.Context, resourceGroupName string, location string, parameters CheckNameAvailabilityRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":          autorest.Encode("path", location),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/locations/{location}/checkNameAvailability", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckNameAvailabilitySender sends the CheckNameAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client BackupVaultsClient) CheckNameAvailabilitySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CheckNameAvailabilityResponder handles the response to the CheckNameAvailability request. The method always
// closes the http.Response Body.
func (client BackupVaultsClient) CheckNameAvailabilityResponder(resp *http.Response) (result CheckNameAvailabilityResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdate creates or updates a BackupVault resource belonging to a resource group.
// Parameters:
// vaultName - the name of the backup vault.
// resourceGroupName - the name of the resource group where the backup vault is present.
// parameters - request body for operation
func (client BackupVaultsClient) CreateOrUpdate(ctx context.Context, vaultName string, resourceGroupName string, parameters BackupVaultResource) (result BackupVaultsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupVaultsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Properties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.Properties.StorageSettings", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("dataprotection.BackupVaultsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, vaultName, resourceGroupName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BackupVaultsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BackupVaultsClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client BackupVaultsClient) CreateOrUpdatePreparer(ctx context.Context, vaultName string, resourceGroupName string, parameters BackupVaultResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2021-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client BackupVaultsClient) CreateOrUpdateSender(req *http.Request) (future BackupVaultsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client BackupVaultsClient) CreateOrUpdateResponder(resp *http.Response) (result BackupVaultResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a BackupVault resource from the resource group.
// Parameters:
// vaultName - the name of the backup vault.
// resourceGroupName - the name of the resource group where the backup vault is present.
func (client BackupVaultsClient) Delete(ctx context.Context, vaultName string, resourceGroupName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupVaultsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, vaultName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BackupVaultsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "dataprotection.BackupVaultsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BackupVaultsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client BackupVaultsClient) DeletePreparer(ctx context.Context, vaultName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2021-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client BackupVaultsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client BackupVaultsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get returns a resource belonging to a resource group.
// Parameters:
// vaultName - the name of the backup vault.
// resourceGroupName - the name of the resource group where the backup vault is present.
func (client BackupVaultsClient) Get(ctx context.Context, vaultName string, resourceGroupName string) (result BackupVaultResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupVaultsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, vaultName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BackupVaultsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dataprotection.BackupVaultsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BackupVaultsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client BackupVaultsClient) GetPreparer(ctx context.Context, vaultName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2021-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client BackupVaultsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client BackupVaultsClient) GetResponder(resp *http.Response) (result BackupVaultResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetInResourceGroup returns resource collection belonging to a resource group.
// Parameters:
// resourceGroupName - the name of the resource group where the backup vault is present.
func (client BackupVaultsClient) GetInResourceGroup(ctx context.Context, resourceGroupName string) (result BackupVaultResourceListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupVaultsClient.GetInResourceGroup")
		defer func() {
			sc := -1
			if result.bvrl.Response.Response != nil {
				sc = result.bvrl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.getInResourceGroupNextResults
	req, err := client.GetInResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BackupVaultsClient", "GetInResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetInResourceGroupSender(req)
	if err != nil {
		result.bvrl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dataprotection.BackupVaultsClient", "GetInResourceGroup", resp, "Failure sending request")
		return
	}

	result.bvrl, err = client.GetInResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BackupVaultsClient", "GetInResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.bvrl.hasNextLink() && result.bvrl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// GetInResourceGroupPreparer prepares the GetInResourceGroup request.
func (client BackupVaultsClient) GetInResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetInResourceGroupSender sends the GetInResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client BackupVaultsClient) GetInResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetInResourceGroupResponder handles the response to the GetInResourceGroup request. The method always
// closes the http.Response Body.
func (client BackupVaultsClient) GetInResourceGroupResponder(resp *http.Response) (result BackupVaultResourceList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getInResourceGroupNextResults retrieves the next set of results, if any.
func (client BackupVaultsClient) getInResourceGroupNextResults(ctx context.Context, lastResults BackupVaultResourceList) (result BackupVaultResourceList, err error) {
	req, err := lastResults.backupVaultResourceListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "dataprotection.BackupVaultsClient", "getInResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetInResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "dataprotection.BackupVaultsClient", "getInResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetInResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BackupVaultsClient", "getInResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// GetInResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client BackupVaultsClient) GetInResourceGroupComplete(ctx context.Context, resourceGroupName string) (result BackupVaultResourceListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupVaultsClient.GetInResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.GetInResourceGroup(ctx, resourceGroupName)
	return
}

// GetInSubscription returns resource collection belonging to a subscription.
func (client BackupVaultsClient) GetInSubscription(ctx context.Context) (result BackupVaultResourceListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupVaultsClient.GetInSubscription")
		defer func() {
			sc := -1
			if result.bvrl.Response.Response != nil {
				sc = result.bvrl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.getInSubscriptionNextResults
	req, err := client.GetInSubscriptionPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BackupVaultsClient", "GetInSubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetInSubscriptionSender(req)
	if err != nil {
		result.bvrl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dataprotection.BackupVaultsClient", "GetInSubscription", resp, "Failure sending request")
		return
	}

	result.bvrl, err = client.GetInSubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BackupVaultsClient", "GetInSubscription", resp, "Failure responding to request")
		return
	}
	if result.bvrl.hasNextLink() && result.bvrl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// GetInSubscriptionPreparer prepares the GetInSubscription request.
func (client BackupVaultsClient) GetInSubscriptionPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DataProtection/backupVaults", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetInSubscriptionSender sends the GetInSubscription request. The method will close the
// http.Response Body if it receives an error.
func (client BackupVaultsClient) GetInSubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetInSubscriptionResponder handles the response to the GetInSubscription request. The method always
// closes the http.Response Body.
func (client BackupVaultsClient) GetInSubscriptionResponder(resp *http.Response) (result BackupVaultResourceList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getInSubscriptionNextResults retrieves the next set of results, if any.
func (client BackupVaultsClient) getInSubscriptionNextResults(ctx context.Context, lastResults BackupVaultResourceList) (result BackupVaultResourceList, err error) {
	req, err := lastResults.backupVaultResourceListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "dataprotection.BackupVaultsClient", "getInSubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetInSubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "dataprotection.BackupVaultsClient", "getInSubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetInSubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BackupVaultsClient", "getInSubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// GetInSubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client BackupVaultsClient) GetInSubscriptionComplete(ctx context.Context) (result BackupVaultResourceListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupVaultsClient.GetInSubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.GetInSubscription(ctx)
	return
}

// Update updates a BackupVault resource belonging to a resource group. For example, updating tags for a resource.
// Parameters:
// vaultName - the name of the backup vault.
// resourceGroupName - the name of the resource group where the backup vault is present.
// parameters - request body for operation
func (client BackupVaultsClient) Update(ctx context.Context, vaultName string, resourceGroupName string, parameters PatchResourceRequestInput) (result BackupVaultsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupVaultsClient.Update")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, vaultName, resourceGroupName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BackupVaultsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BackupVaultsClient", "Update", nil, "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client BackupVaultsClient) UpdatePreparer(ctx context.Context, vaultName string, resourceGroupName string, parameters PatchResourceRequestInput) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2021-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client BackupVaultsClient) UpdateSender(req *http.Request) (future BackupVaultsUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client BackupVaultsClient) UpdateResponder(resp *http.Response) (result BackupVaultResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
