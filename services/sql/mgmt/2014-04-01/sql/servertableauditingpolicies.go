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

// ServerTableAuditingPoliciesClient is the the Azure SQL Database management API provides a RESTful set of web
// services that interact with Azure SQL Database services to manage your databases. The API enables you to create,
// retrieve, update, and delete databases.
type ServerTableAuditingPoliciesClient struct {
	BaseClient
}

// NewServerTableAuditingPoliciesClient creates an instance of the ServerTableAuditingPoliciesClient client.
func NewServerTableAuditingPoliciesClient(subscriptionID string) ServerTableAuditingPoliciesClient {
	return NewServerTableAuditingPoliciesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewServerTableAuditingPoliciesClientWithBaseURI creates an instance of the ServerTableAuditingPoliciesClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewServerTableAuditingPoliciesClientWithBaseURI(baseURI string, subscriptionID string) ServerTableAuditingPoliciesClient {
	return ServerTableAuditingPoliciesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a server's table auditing policy. Table auditing is deprecated, use blob auditing
// instead.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// parameters - the server table auditing policy.
func (client ServerTableAuditingPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, parameters ServerTableAuditingPolicy) (result ServerTableAuditingPolicy, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServerTableAuditingPoliciesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serverName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServerTableAuditingPoliciesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ServerTableAuditingPoliciesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServerTableAuditingPoliciesClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ServerTableAuditingPoliciesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serverName string, parameters ServerTableAuditingPolicy) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"serverName":              autorest.Encode("path", serverName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
		"tableAuditingPolicyName": autorest.Encode("path", "default"),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.Kind = nil
	parameters.Location = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/auditingPolicies/{tableAuditingPolicyName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ServerTableAuditingPoliciesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ServerTableAuditingPoliciesClient) CreateOrUpdateResponder(resp *http.Response) (result ServerTableAuditingPolicy, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets a server's table auditing policy. Table auditing is deprecated, use blob auditing instead.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
func (client ServerTableAuditingPoliciesClient) Get(ctx context.Context, resourceGroupName string, serverName string) (result ServerTableAuditingPolicy, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServerTableAuditingPoliciesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, serverName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServerTableAuditingPoliciesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ServerTableAuditingPoliciesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServerTableAuditingPoliciesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ServerTableAuditingPoliciesClient) GetPreparer(ctx context.Context, resourceGroupName string, serverName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"serverName":              autorest.Encode("path", serverName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
		"tableAuditingPolicyName": autorest.Encode("path", "default"),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/auditingPolicies/{tableAuditingPolicyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ServerTableAuditingPoliciesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ServerTableAuditingPoliciesClient) GetResponder(resp *http.Response) (result ServerTableAuditingPolicy, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByServer lists a server's table auditing policies. Table auditing is deprecated, use blob auditing instead.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
func (client ServerTableAuditingPoliciesClient) ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result ServerTableAuditingPolicyListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServerTableAuditingPoliciesClient.ListByServer")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListByServerPreparer(ctx, resourceGroupName, serverName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServerTableAuditingPoliciesClient", "ListByServer", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ServerTableAuditingPoliciesClient", "ListByServer", resp, "Failure sending request")
		return
	}

	result, err = client.ListByServerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServerTableAuditingPoliciesClient", "ListByServer", resp, "Failure responding to request")
		return
	}

	return
}

// ListByServerPreparer prepares the ListByServer request.
func (client ServerTableAuditingPoliciesClient) ListByServerPreparer(ctx context.Context, resourceGroupName string, serverName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/auditingPolicies", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByServerSender sends the ListByServer request. The method will close the
// http.Response Body if it receives an error.
func (client ServerTableAuditingPoliciesClient) ListByServerSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByServerResponder handles the response to the ListByServer request. The method always
// closes the http.Response Body.
func (client ServerTableAuditingPoliciesClient) ListByServerResponder(resp *http.Response) (result ServerTableAuditingPolicyListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
