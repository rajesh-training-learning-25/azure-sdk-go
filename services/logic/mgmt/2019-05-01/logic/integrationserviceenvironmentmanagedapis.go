package logic

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
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

// IntegrationServiceEnvironmentManagedApisClient is the REST API for Azure Logic Apps.
type IntegrationServiceEnvironmentManagedApisClient struct {
	BaseClient
}

// NewIntegrationServiceEnvironmentManagedApisClient creates an instance of the
// IntegrationServiceEnvironmentManagedApisClient client.
func NewIntegrationServiceEnvironmentManagedApisClient(subscriptionID string) IntegrationServiceEnvironmentManagedApisClient {
	return NewIntegrationServiceEnvironmentManagedApisClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewIntegrationServiceEnvironmentManagedApisClientWithBaseURI creates an instance of the
// IntegrationServiceEnvironmentManagedApisClient client using a custom endpoint.  Use this when interacting with an
// Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewIntegrationServiceEnvironmentManagedApisClientWithBaseURI(baseURI string, subscriptionID string) IntegrationServiceEnvironmentManagedApisClient {
	return IntegrationServiceEnvironmentManagedApisClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Delete deletes the integration service environment managed Api.
// Parameters:
// resourceGroup - the resource group.
// integrationServiceEnvironmentName - the integration service environment name.
// APIName - the api name.
func (client IntegrationServiceEnvironmentManagedApisClient) Delete(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, APIName string) (result IntegrationServiceEnvironmentManagedApisDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationServiceEnvironmentManagedApisClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroup, integrationServiceEnvironmentName, APIName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentManagedApisClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentManagedApisClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client IntegrationServiceEnvironmentManagedApisClient) DeletePreparer(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, APIName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"apiName":                           autorest.Encode("path", APIName),
		"integrationServiceEnvironmentName": autorest.Encode("path", integrationServiceEnvironmentName),
		"resourceGroup":                     autorest.Encode("path", resourceGroup),
		"subscriptionId":                    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/managedApis/{apiName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationServiceEnvironmentManagedApisClient) DeleteSender(req *http.Request) (future IntegrationServiceEnvironmentManagedApisDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
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

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client IntegrationServiceEnvironmentManagedApisClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the integration service environment managed Api.
// Parameters:
// resourceGroup - the resource group name.
// integrationServiceEnvironmentName - the integration service environment name.
// APIName - the api name.
func (client IntegrationServiceEnvironmentManagedApisClient) Get(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, APIName string) (result ManagedAPI, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationServiceEnvironmentManagedApisClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroup, integrationServiceEnvironmentName, APIName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentManagedApisClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentManagedApisClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentManagedApisClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client IntegrationServiceEnvironmentManagedApisClient) GetPreparer(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, APIName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"apiName":                           autorest.Encode("path", APIName),
		"integrationServiceEnvironmentName": autorest.Encode("path", integrationServiceEnvironmentName),
		"resourceGroup":                     autorest.Encode("path", resourceGroup),
		"subscriptionId":                    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/managedApis/{apiName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationServiceEnvironmentManagedApisClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client IntegrationServiceEnvironmentManagedApisClient) GetResponder(resp *http.Response) (result ManagedAPI, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets the integration service environment managed Apis.
// Parameters:
// resourceGroup - the resource group.
// integrationServiceEnvironmentName - the integration service environment name.
func (client IntegrationServiceEnvironmentManagedApisClient) List(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string) (result ManagedAPIListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationServiceEnvironmentManagedApisClient.List")
		defer func() {
			sc := -1
			if result.malr.Response.Response != nil {
				sc = result.malr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroup, integrationServiceEnvironmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentManagedApisClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.malr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentManagedApisClient", "List", resp, "Failure sending request")
		return
	}

	result.malr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentManagedApisClient", "List", resp, "Failure responding to request")
		return
	}
	if result.malr.hasNextLink() && result.malr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client IntegrationServiceEnvironmentManagedApisClient) ListPreparer(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationServiceEnvironmentName": autorest.Encode("path", integrationServiceEnvironmentName),
		"resourceGroup":                     autorest.Encode("path", resourceGroup),
		"subscriptionId":                    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/managedApis", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationServiceEnvironmentManagedApisClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client IntegrationServiceEnvironmentManagedApisClient) ListResponder(resp *http.Response) (result ManagedAPIListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client IntegrationServiceEnvironmentManagedApisClient) listNextResults(ctx context.Context, lastResults ManagedAPIListResult) (result ManagedAPIListResult, err error) {
	req, err := lastResults.managedAPIListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentManagedApisClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentManagedApisClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentManagedApisClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client IntegrationServiceEnvironmentManagedApisClient) ListComplete(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string) (result ManagedAPIListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationServiceEnvironmentManagedApisClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroup, integrationServiceEnvironmentName)
	return
}

// Put puts the integration service environment managed Api.
// Parameters:
// resourceGroup - the resource group name.
// integrationServiceEnvironmentName - the integration service environment name.
// APIName - the api name.
func (client IntegrationServiceEnvironmentManagedApisClient) Put(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, APIName string) (result IntegrationServiceEnvironmentManagedApisPutFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationServiceEnvironmentManagedApisClient.Put")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.PutPreparer(ctx, resourceGroup, integrationServiceEnvironmentName, APIName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentManagedApisClient", "Put", nil, "Failure preparing request")
		return
	}

	result, err = client.PutSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentManagedApisClient", "Put", nil, "Failure sending request")
		return
	}

	return
}

// PutPreparer prepares the Put request.
func (client IntegrationServiceEnvironmentManagedApisClient) PutPreparer(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, APIName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"apiName":                           autorest.Encode("path", APIName),
		"integrationServiceEnvironmentName": autorest.Encode("path", integrationServiceEnvironmentName),
		"resourceGroup":                     autorest.Encode("path", resourceGroup),
		"subscriptionId":                    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/managedApis/{apiName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutSender sends the Put request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationServiceEnvironmentManagedApisClient) PutSender(req *http.Request) (future IntegrationServiceEnvironmentManagedApisPutFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
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

// PutResponder handles the response to the Put request. The method always
// closes the http.Response Body.
func (client IntegrationServiceEnvironmentManagedApisClient) PutResponder(resp *http.Response) (result ManagedAPI, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
