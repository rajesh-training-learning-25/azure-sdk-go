package sql

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

// ManagedServerSecurityAlertPoliciesClient is the the Azure SQL Database management API provides a RESTful set of web
// services that interact with Azure SQL Database services to manage your databases. The API enables you to create,
// retrieve, update, and delete databases.
type ManagedServerSecurityAlertPoliciesClient struct {
	BaseClient
}

// NewManagedServerSecurityAlertPoliciesClient creates an instance of the ManagedServerSecurityAlertPoliciesClient
// client.
func NewManagedServerSecurityAlertPoliciesClient(subscriptionID string) ManagedServerSecurityAlertPoliciesClient {
	return NewManagedServerSecurityAlertPoliciesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewManagedServerSecurityAlertPoliciesClientWithBaseURI creates an instance of the
// ManagedServerSecurityAlertPoliciesClient client using a custom endpoint.  Use this when interacting with an Azure
// cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewManagedServerSecurityAlertPoliciesClientWithBaseURI(baseURI string, subscriptionID string) ManagedServerSecurityAlertPoliciesClient {
	return ManagedServerSecurityAlertPoliciesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a threat detection policy.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// managedInstanceName - the name of the managed instance.
// parameters - the managed server security alert policy.
func (client ManagedServerSecurityAlertPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters ManagedServerSecurityAlertPolicy) (result ManagedServerSecurityAlertPoliciesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedServerSecurityAlertPoliciesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, managedInstanceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedServerSecurityAlertPoliciesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedServerSecurityAlertPoliciesClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ManagedServerSecurityAlertPoliciesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters ManagedServerSecurityAlertPolicy) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managedInstanceName":     autorest.Encode("path", managedInstanceName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"securityAlertPolicyName": autorest.Encode("path", "Default"),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.SystemData = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/securityAlertPolicies/{securityAlertPolicyName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedServerSecurityAlertPoliciesClient) CreateOrUpdateSender(req *http.Request) (future ManagedServerSecurityAlertPoliciesCreateOrUpdateFuture, err error) {
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

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ManagedServerSecurityAlertPoliciesClient) CreateOrUpdateResponder(resp *http.Response) (result ManagedServerSecurityAlertPolicy, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get a managed server's threat detection policy.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// managedInstanceName - the name of the managed instance.
func (client ManagedServerSecurityAlertPoliciesClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string) (result ManagedServerSecurityAlertPolicy, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedServerSecurityAlertPoliciesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, managedInstanceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedServerSecurityAlertPoliciesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ManagedServerSecurityAlertPoliciesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedServerSecurityAlertPoliciesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ManagedServerSecurityAlertPoliciesClient) GetPreparer(ctx context.Context, resourceGroupName string, managedInstanceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managedInstanceName":     autorest.Encode("path", managedInstanceName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"securityAlertPolicyName": autorest.Encode("path", "Default"),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/securityAlertPolicies/{securityAlertPolicyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedServerSecurityAlertPoliciesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ManagedServerSecurityAlertPoliciesClient) GetResponder(resp *http.Response) (result ManagedServerSecurityAlertPolicy, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByInstance get the managed server's threat detection policies.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// managedInstanceName - the name of the managed instance.
func (client ManagedServerSecurityAlertPoliciesClient) ListByInstance(ctx context.Context, resourceGroupName string, managedInstanceName string) (result ManagedServerSecurityAlertPolicyListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedServerSecurityAlertPoliciesClient.ListByInstance")
		defer func() {
			sc := -1
			if result.mssaplr.Response.Response != nil {
				sc = result.mssaplr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByInstanceNextResults
	req, err := client.ListByInstancePreparer(ctx, resourceGroupName, managedInstanceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedServerSecurityAlertPoliciesClient", "ListByInstance", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByInstanceSender(req)
	if err != nil {
		result.mssaplr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ManagedServerSecurityAlertPoliciesClient", "ListByInstance", resp, "Failure sending request")
		return
	}

	result.mssaplr, err = client.ListByInstanceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedServerSecurityAlertPoliciesClient", "ListByInstance", resp, "Failure responding to request")
		return
	}
	if result.mssaplr.hasNextLink() && result.mssaplr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByInstancePreparer prepares the ListByInstance request.
func (client ManagedServerSecurityAlertPoliciesClient) ListByInstancePreparer(ctx context.Context, resourceGroupName string, managedInstanceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managedInstanceName": autorest.Encode("path", managedInstanceName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/securityAlertPolicies", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByInstanceSender sends the ListByInstance request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedServerSecurityAlertPoliciesClient) ListByInstanceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByInstanceResponder handles the response to the ListByInstance request. The method always
// closes the http.Response Body.
func (client ManagedServerSecurityAlertPoliciesClient) ListByInstanceResponder(resp *http.Response) (result ManagedServerSecurityAlertPolicyListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByInstanceNextResults retrieves the next set of results, if any.
func (client ManagedServerSecurityAlertPoliciesClient) listByInstanceNextResults(ctx context.Context, lastResults ManagedServerSecurityAlertPolicyListResult) (result ManagedServerSecurityAlertPolicyListResult, err error) {
	req, err := lastResults.managedServerSecurityAlertPolicyListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.ManagedServerSecurityAlertPoliciesClient", "listByInstanceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByInstanceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.ManagedServerSecurityAlertPoliciesClient", "listByInstanceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByInstanceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedServerSecurityAlertPoliciesClient", "listByInstanceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByInstanceComplete enumerates all values, automatically crossing page boundaries as required.
func (client ManagedServerSecurityAlertPoliciesClient) ListByInstanceComplete(ctx context.Context, resourceGroupName string, managedInstanceName string) (result ManagedServerSecurityAlertPolicyListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedServerSecurityAlertPoliciesClient.ListByInstance")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByInstance(ctx, resourceGroupName, managedInstanceName)
	return
}
