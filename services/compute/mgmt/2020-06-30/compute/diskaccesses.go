package compute

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// DiskAccessesClient is the compute Client
type DiskAccessesClient struct {
	BaseClient
}

// NewDiskAccessesClient creates an instance of the DiskAccessesClient client.
func NewDiskAccessesClient(subscriptionID string) DiskAccessesClient {
	return NewDiskAccessesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDiskAccessesClientWithBaseURI creates an instance of the DiskAccessesClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewDiskAccessesClientWithBaseURI(baseURI string, subscriptionID string) DiskAccessesClient {
	return DiskAccessesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a disk access resource
// Parameters:
// resourceGroupName - the name of the resource group.
// diskAccessName - the name of the disk access resource that is being created. The name can't be changed after
// the disk encryption set is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The maximum
// name length is 80 characters.
// diskAccess - disk access object supplied in the body of the Put disk access operation.
func (client DiskAccessesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, diskAccessName string, diskAccess DiskAccess) (result DiskAccessesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DiskAccessesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, diskAccessName, diskAccess)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskAccessesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskAccessesClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DiskAccessesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, diskAccessName string, diskAccess DiskAccess) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"diskAccessName":    autorest.Encode("path", diskAccessName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-06-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskAccesses/{diskAccessName}", pathParameters),
		autorest.WithJSON(diskAccess),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client DiskAccessesClient) CreateOrUpdateSender(req *http.Request) (future DiskAccessesCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client DiskAccessesClient) (da DiskAccess, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "compute.DiskAccessesCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("compute.DiskAccessesCreateOrUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		da.Response.Response, err = future.GetResult(sender)
		if da.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "compute.DiskAccessesCreateOrUpdateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && da.Response.Response.StatusCode != http.StatusNoContent {
			da, err = client.CreateOrUpdateResponder(da.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "compute.DiskAccessesCreateOrUpdateFuture", "Result", da.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client DiskAccessesClient) CreateOrUpdateResponder(resp *http.Response) (result DiskAccess, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a disk access resource.
// Parameters:
// resourceGroupName - the name of the resource group.
// diskAccessName - the name of the disk access resource that is being created. The name can't be changed after
// the disk encryption set is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The maximum
// name length is 80 characters.
func (client DiskAccessesClient) Delete(ctx context.Context, resourceGroupName string, diskAccessName string) (result DiskAccessesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DiskAccessesClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, diskAccessName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskAccessesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskAccessesClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client DiskAccessesClient) DeletePreparer(ctx context.Context, resourceGroupName string, diskAccessName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"diskAccessName":    autorest.Encode("path", diskAccessName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-06-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskAccesses/{diskAccessName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DiskAccessesClient) DeleteSender(req *http.Request) (future DiskAccessesDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client DiskAccessesClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "compute.DiskAccessesDeleteFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("compute.DiskAccessesDeleteFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client DiskAccessesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets information about a disk access resource.
// Parameters:
// resourceGroupName - the name of the resource group.
// diskAccessName - the name of the disk access resource that is being created. The name can't be changed after
// the disk encryption set is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The maximum
// name length is 80 characters.
func (client DiskAccessesClient) Get(ctx context.Context, resourceGroupName string, diskAccessName string) (result DiskAccess, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DiskAccessesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, diskAccessName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskAccessesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.DiskAccessesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskAccessesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client DiskAccessesClient) GetPreparer(ctx context.Context, resourceGroupName string, diskAccessName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"diskAccessName":    autorest.Encode("path", diskAccessName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-06-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskAccesses/{diskAccessName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DiskAccessesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DiskAccessesClient) GetResponder(resp *http.Response) (result DiskAccess, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetPrivateLinkResources gets the private link resources possible under disk access resource
// Parameters:
// resourceGroupName - the name of the resource group.
// diskAccessName - the name of the disk access resource that is being created. The name can't be changed after
// the disk encryption set is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The maximum
// name length is 80 characters.
func (client DiskAccessesClient) GetPrivateLinkResources(ctx context.Context, resourceGroupName string, diskAccessName string) (result PrivateLinkResourceListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DiskAccessesClient.GetPrivateLinkResources")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPrivateLinkResourcesPreparer(ctx, resourceGroupName, diskAccessName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskAccessesClient", "GetPrivateLinkResources", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetPrivateLinkResourcesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.DiskAccessesClient", "GetPrivateLinkResources", resp, "Failure sending request")
		return
	}

	result, err = client.GetPrivateLinkResourcesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskAccessesClient", "GetPrivateLinkResources", resp, "Failure responding to request")
		return
	}

	return
}

// GetPrivateLinkResourcesPreparer prepares the GetPrivateLinkResources request.
func (client DiskAccessesClient) GetPrivateLinkResourcesPreparer(ctx context.Context, resourceGroupName string, diskAccessName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"diskAccessName":    autorest.Encode("path", diskAccessName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-06-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskAccesses/{diskAccessName}/privateLinkResources", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetPrivateLinkResourcesSender sends the GetPrivateLinkResources request. The method will close the
// http.Response Body if it receives an error.
func (client DiskAccessesClient) GetPrivateLinkResourcesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetPrivateLinkResourcesResponder handles the response to the GetPrivateLinkResources request. The method always
// closes the http.Response Body.
func (client DiskAccessesClient) GetPrivateLinkResourcesResponder(resp *http.Response) (result PrivateLinkResourceListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all the disk access resources under a subscription.
func (client DiskAccessesClient) List(ctx context.Context) (result DiskAccessListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DiskAccessesClient.List")
		defer func() {
			sc := -1
			if result.dal.Response.Response != nil {
				sc = result.dal.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskAccessesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.dal.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.DiskAccessesClient", "List", resp, "Failure sending request")
		return
	}

	result.dal, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskAccessesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.dal.hasNextLink() && result.dal.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client DiskAccessesClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-06-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/diskAccesses", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DiskAccessesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DiskAccessesClient) ListResponder(resp *http.Response) (result DiskAccessList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client DiskAccessesClient) listNextResults(ctx context.Context, lastResults DiskAccessList) (result DiskAccessList, err error) {
	req, err := lastResults.diskAccessListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.DiskAccessesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.DiskAccessesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskAccessesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client DiskAccessesClient) ListComplete(ctx context.Context) (result DiskAccessListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DiskAccessesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}

// ListByResourceGroup lists all the disk access resources under a resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
func (client DiskAccessesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result DiskAccessListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DiskAccessesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.dal.Response.Response != nil {
				sc = result.dal.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskAccessesClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.dal.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.DiskAccessesClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.dal, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskAccessesClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.dal.hasNextLink() && result.dal.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client DiskAccessesClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-06-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskAccesses", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client DiskAccessesClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client DiskAccessesClient) ListByResourceGroupResponder(resp *http.Response) (result DiskAccessList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client DiskAccessesClient) listByResourceGroupNextResults(ctx context.Context, lastResults DiskAccessList) (result DiskAccessList, err error) {
	req, err := lastResults.diskAccessListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.DiskAccessesClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.DiskAccessesClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskAccessesClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client DiskAccessesClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result DiskAccessListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DiskAccessesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// Update updates (patches) a disk access resource.
// Parameters:
// resourceGroupName - the name of the resource group.
// diskAccessName - the name of the disk access resource that is being created. The name can't be changed after
// the disk encryption set is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The maximum
// name length is 80 characters.
// diskAccess - disk access object supplied in the body of the Patch disk access operation.
func (client DiskAccessesClient) Update(ctx context.Context, resourceGroupName string, diskAccessName string, diskAccess DiskAccessUpdate) (result DiskAccessesUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DiskAccessesClient.Update")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, diskAccessName, diskAccess)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskAccessesClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskAccessesClient", "Update", nil, "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client DiskAccessesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, diskAccessName string, diskAccess DiskAccessUpdate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"diskAccessName":    autorest.Encode("path", diskAccessName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-06-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskAccesses/{diskAccessName}", pathParameters),
		autorest.WithJSON(diskAccess),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client DiskAccessesClient) UpdateSender(req *http.Request) (future DiskAccessesUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client DiskAccessesClient) (da DiskAccess, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "compute.DiskAccessesUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("compute.DiskAccessesUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		da.Response.Response, err = future.GetResult(sender)
		if da.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "compute.DiskAccessesUpdateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && da.Response.Response.StatusCode != http.StatusNoContent {
			da, err = client.UpdateResponder(da.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "compute.DiskAccessesUpdateFuture", "Result", da.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client DiskAccessesClient) UpdateResponder(resp *http.Response) (result DiskAccess, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
