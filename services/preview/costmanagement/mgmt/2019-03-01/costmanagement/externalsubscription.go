package costmanagement

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

// ExternalSubscriptionClient is the client for the ExternalSubscription methods of the Costmanagement service.
type ExternalSubscriptionClient struct {
	BaseClient
}

// NewExternalSubscriptionClient creates an instance of the ExternalSubscriptionClient client.
func NewExternalSubscriptionClient(subscriptionID string) ExternalSubscriptionClient {
	return NewExternalSubscriptionClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewExternalSubscriptionClientWithBaseURI creates an instance of the ExternalSubscriptionClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewExternalSubscriptionClientWithBaseURI(baseURI string, subscriptionID string) ExternalSubscriptionClient {
	return ExternalSubscriptionClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get an ExternalSubscription definition
// Parameters:
// externalSubscriptionName - external Subscription Name. (eg 'aws-{UsageAccountId}')
// expand - may be used to expand the collectionInfo property. By default, collectionInfo is not included.
func (client ExternalSubscriptionClient) Get(ctx context.Context, externalSubscriptionName string, expand string) (result ExternalSubscriptionDefinition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExternalSubscriptionClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, externalSubscriptionName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ExternalSubscriptionClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.ExternalSubscriptionClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ExternalSubscriptionClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ExternalSubscriptionClient) GetPreparer(ctx context.Context, externalSubscriptionName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"externalSubscriptionName": autorest.Encode("path", externalSubscriptionName),
	}

	const APIVersion = "2019-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ExternalSubscriptionClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ExternalSubscriptionClient) GetResponder(resp *http.Response) (result ExternalSubscriptionDefinition, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list all ExternalSubscription definitions
func (client ExternalSubscriptionClient) List(ctx context.Context) (result ExternalSubscriptionDefinitionListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExternalSubscriptionClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ExternalSubscriptionClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.ExternalSubscriptionClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ExternalSubscriptionClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ExternalSubscriptionClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2019-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.CostManagement/externalSubscriptions"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ExternalSubscriptionClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ExternalSubscriptionClient) ListResponder(resp *http.Response) (result ExternalSubscriptionDefinitionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByExternalBillingAccount list all ExternalSubscriptions by ExternalBillingAccount definitions
// Parameters:
// externalBillingAccountName - external Billing Account Name. (eg 'aws-{PayerAccountId}')
func (client ExternalSubscriptionClient) ListByExternalBillingAccount(ctx context.Context, externalBillingAccountName string) (result ExternalSubscriptionDefinitionListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExternalSubscriptionClient.ListByExternalBillingAccount")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListByExternalBillingAccountPreparer(ctx, externalBillingAccountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ExternalSubscriptionClient", "ListByExternalBillingAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByExternalBillingAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.ExternalSubscriptionClient", "ListByExternalBillingAccount", resp, "Failure sending request")
		return
	}

	result, err = client.ListByExternalBillingAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ExternalSubscriptionClient", "ListByExternalBillingAccount", resp, "Failure responding to request")
		return
	}

	return
}

// ListByExternalBillingAccountPreparer prepares the ListByExternalBillingAccount request.
func (client ExternalSubscriptionClient) ListByExternalBillingAccountPreparer(ctx context.Context, externalBillingAccountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"externalBillingAccountName": autorest.Encode("path", externalBillingAccountName),
	}

	const APIVersion = "2019-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}/externalSubscriptions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByExternalBillingAccountSender sends the ListByExternalBillingAccount request. The method will close the
// http.Response Body if it receives an error.
func (client ExternalSubscriptionClient) ListByExternalBillingAccountSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByExternalBillingAccountResponder handles the response to the ListByExternalBillingAccount request. The method always
// closes the http.Response Body.
func (client ExternalSubscriptionClient) ListByExternalBillingAccountResponder(resp *http.Response) (result ExternalSubscriptionDefinitionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByManagementGroup list all ExternalSubscription definitions for Management Group
// Parameters:
// managementGroupID - managementGroup ID
// recurse - the $recurse=true query string parameter allows returning externalSubscriptions associated with
// the specified managementGroup, or any of its descendants.
func (client ExternalSubscriptionClient) ListByManagementGroup(ctx context.Context, managementGroupID string, recurse *bool) (result ExternalSubscriptionDefinitionListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExternalSubscriptionClient.ListByManagementGroup")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListByManagementGroupPreparer(ctx, managementGroupID, recurse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ExternalSubscriptionClient", "ListByManagementGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByManagementGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.ExternalSubscriptionClient", "ListByManagementGroup", resp, "Failure sending request")
		return
	}

	result, err = client.ListByManagementGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ExternalSubscriptionClient", "ListByManagementGroup", resp, "Failure responding to request")
		return
	}

	return
}

// ListByManagementGroupPreparer prepares the ListByManagementGroup request.
func (client ExternalSubscriptionClient) ListByManagementGroupPreparer(ctx context.Context, managementGroupID string, recurse *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementGroupId": autorest.Encode("path", managementGroupID),
	}

	const APIVersion = "2019-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if recurse != nil {
		queryParameters["$recurse"] = autorest.Encode("query", *recurse)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.CostManagement/externalSubscriptions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByManagementGroupSender sends the ListByManagementGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ExternalSubscriptionClient) ListByManagementGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByManagementGroupResponder handles the response to the ListByManagementGroup request. The method always
// closes the http.Response Body.
func (client ExternalSubscriptionClient) ListByManagementGroupResponder(resp *http.Response) (result ExternalSubscriptionDefinitionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UpdateManagementGroup updates the management group of an ExternalSubscription
// Parameters:
// managementGroupID - managementGroup ID
// externalSubscriptionName - external Subscription Name. (eg 'aws-{UsageAccountId}')
func (client ExternalSubscriptionClient) UpdateManagementGroup(ctx context.Context, managementGroupID string, externalSubscriptionName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExternalSubscriptionClient.UpdateManagementGroup")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdateManagementGroupPreparer(ctx, managementGroupID, externalSubscriptionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ExternalSubscriptionClient", "UpdateManagementGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateManagementGroupSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "costmanagement.ExternalSubscriptionClient", "UpdateManagementGroup", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateManagementGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ExternalSubscriptionClient", "UpdateManagementGroup", resp, "Failure responding to request")
		return
	}

	return
}

// UpdateManagementGroupPreparer prepares the UpdateManagementGroup request.
func (client ExternalSubscriptionClient) UpdateManagementGroupPreparer(ctx context.Context, managementGroupID string, externalSubscriptionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"externalSubscriptionName": autorest.Encode("path", externalSubscriptionName),
		"managementGroupId":        autorest.Encode("path", managementGroupID),
	}

	const APIVersion = "2019-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateManagementGroupSender sends the UpdateManagementGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ExternalSubscriptionClient) UpdateManagementGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateManagementGroupResponder handles the response to the UpdateManagementGroup request. The method always
// closes the http.Response Body.
func (client ExternalSubscriptionClient) UpdateManagementGroupResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
