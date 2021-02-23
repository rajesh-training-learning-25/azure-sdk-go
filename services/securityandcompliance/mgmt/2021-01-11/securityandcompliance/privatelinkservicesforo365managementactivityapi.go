package securityandcompliance

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

// PrivateLinkServicesForO365ManagementActivityAPIClient is the use this API to manage Microsoft SecurityAndCompliance
// services in your Azure subscription.
type PrivateLinkServicesForO365ManagementActivityAPIClient struct {
	BaseClient
}

// NewPrivateLinkServicesForO365ManagementActivityAPIClient creates an instance of the
// PrivateLinkServicesForO365ManagementActivityAPIClient client.
func NewPrivateLinkServicesForO365ManagementActivityAPIClient(subscriptionID string) PrivateLinkServicesForO365ManagementActivityAPIClient {
	return NewPrivateLinkServicesForO365ManagementActivityAPIClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPrivateLinkServicesForO365ManagementActivityAPIClientWithBaseURI creates an instance of the
// PrivateLinkServicesForO365ManagementActivityAPIClient client using a custom endpoint.  Use this when interacting
// with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewPrivateLinkServicesForO365ManagementActivityAPIClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkServicesForO365ManagementActivityAPIClient {
	return PrivateLinkServicesForO365ManagementActivityAPIClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or update the metadata of a privateLinkServicesForO365ManagementActivityAPI instance.
// Parameters:
// resourceGroupName - the name of the resource group that contains the service instance.
// resourceName - the name of the service instance.
// privateLinkServicesForO365ManagementActivityAPIDescription - the service instance metadata.
func (client PrivateLinkServicesForO365ManagementActivityAPIClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, privateLinkServicesForO365ManagementActivityAPIDescription PrivateLinkServicesForO365ManagementActivityAPIDescription) (result PrivateLinkServicesForO365ManagementActivityAPICreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkServicesForO365ManagementActivityAPIClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: privateLinkServicesForO365ManagementActivityAPIDescription,
			Constraints: []validation.Constraint{{Target: "privateLinkServicesForO365ManagementActivityAPIDescription.Properties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "privateLinkServicesForO365ManagementActivityAPIDescription.Properties.CosmosDbConfiguration", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "privateLinkServicesForO365ManagementActivityAPIDescription.Properties.CosmosDbConfiguration.OfferThroughput", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "privateLinkServicesForO365ManagementActivityAPIDescription.Properties.CosmosDbConfiguration.OfferThroughput", Name: validation.InclusiveMaximum, Rule: int64(10000), Chain: nil},
							{Target: "privateLinkServicesForO365ManagementActivityAPIDescription.Properties.CosmosDbConfiguration.OfferThroughput", Name: validation.InclusiveMinimum, Rule: int64(400), Chain: nil},
						}},
					}},
					{Target: "privateLinkServicesForO365ManagementActivityAPIDescription.Properties.CorsConfiguration", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "privateLinkServicesForO365ManagementActivityAPIDescription.Properties.CorsConfiguration.MaxAge", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "privateLinkServicesForO365ManagementActivityAPIDescription.Properties.CorsConfiguration.MaxAge", Name: validation.InclusiveMaximum, Rule: int64(99999), Chain: nil},
								{Target: "privateLinkServicesForO365ManagementActivityAPIDescription.Properties.CorsConfiguration.MaxAge", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil},
							}},
						}},
				}}}}}); err != nil {
		return result, validation.NewError("securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPIClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, resourceName, privateLinkServicesForO365ManagementActivityAPIDescription)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPIClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPIClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client PrivateLinkServicesForO365ManagementActivityAPIClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, resourceName string, privateLinkServicesForO365ManagementActivityAPIDescription PrivateLinkServicesForO365ManagementActivityAPIDescription) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-01-11"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForO365ManagementActivityAPI/{resourceName}", pathParameters),
		autorest.WithJSON(privateLinkServicesForO365ManagementActivityAPIDescription),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateLinkServicesForO365ManagementActivityAPIClient) CreateOrUpdateSender(req *http.Request) (future PrivateLinkServicesForO365ManagementActivityAPICreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client PrivateLinkServicesForO365ManagementActivityAPIClient) (plsfomaad PrivateLinkServicesForO365ManagementActivityAPIDescription, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPICreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPICreateOrUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		plsfomaad.Response.Response, err = future.GetResult(sender)
		if plsfomaad.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPICreateOrUpdateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && plsfomaad.Response.Response.StatusCode != http.StatusNoContent {
			plsfomaad, err = client.CreateOrUpdateResponder(plsfomaad.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPICreateOrUpdateFuture", "Result", plsfomaad.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client PrivateLinkServicesForO365ManagementActivityAPIClient) CreateOrUpdateResponder(resp *http.Response) (result PrivateLinkServicesForO365ManagementActivityAPIDescription, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a service instance.
// Parameters:
// resourceGroupName - the name of the resource group that contains the service instance.
// resourceName - the name of the service instance.
func (client PrivateLinkServicesForO365ManagementActivityAPIClient) Delete(ctx context.Context, resourceGroupName string, resourceName string) (result PrivateLinkServicesForO365ManagementActivityAPIDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkServicesForO365ManagementActivityAPIClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPIClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPIClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPIClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client PrivateLinkServicesForO365ManagementActivityAPIClient) DeletePreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-01-11"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForO365ManagementActivityAPI/{resourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateLinkServicesForO365ManagementActivityAPIClient) DeleteSender(req *http.Request) (future PrivateLinkServicesForO365ManagementActivityAPIDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client PrivateLinkServicesForO365ManagementActivityAPIClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPIDeleteFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPIDeleteFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client PrivateLinkServicesForO365ManagementActivityAPIClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get the metadata of a privateLinkServicesForO365ManagementActivityAPI resource.
// Parameters:
// resourceGroupName - the name of the resource group that contains the service instance.
// resourceName - the name of the service instance.
func (client PrivateLinkServicesForO365ManagementActivityAPIClient) Get(ctx context.Context, resourceGroupName string, resourceName string) (result PrivateLinkServicesForO365ManagementActivityAPIDescription, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkServicesForO365ManagementActivityAPIClient.Get")
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
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPIClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPIClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPIClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPIClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client PrivateLinkServicesForO365ManagementActivityAPIClient) GetPreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-01-11"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForO365ManagementActivityAPI/{resourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateLinkServicesForO365ManagementActivityAPIClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PrivateLinkServicesForO365ManagementActivityAPIClient) GetResponder(resp *http.Response) (result PrivateLinkServicesForO365ManagementActivityAPIDescription, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List get all the privateLinkServicesForO365ManagementActivityAPI instances in a subscription.
func (client PrivateLinkServicesForO365ManagementActivityAPIClient) List(ctx context.Context) (result PrivateLinkServicesForO365ManagementActivityAPIDescriptionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkServicesForO365ManagementActivityAPIClient.List")
		defer func() {
			sc := -1
			if result.plsfomaadlr.Response.Response != nil {
				sc = result.plsfomaadlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPIClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.plsfomaadlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPIClient", "List", resp, "Failure sending request")
		return
	}

	result.plsfomaadlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPIClient", "List", resp, "Failure responding to request")
		return
	}
	if result.plsfomaadlr.hasNextLink() && result.plsfomaadlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client PrivateLinkServicesForO365ManagementActivityAPIClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-01-11"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForO365ManagementActivityAPI", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateLinkServicesForO365ManagementActivityAPIClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client PrivateLinkServicesForO365ManagementActivityAPIClient) ListResponder(resp *http.Response) (result PrivateLinkServicesForO365ManagementActivityAPIDescriptionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client PrivateLinkServicesForO365ManagementActivityAPIClient) listNextResults(ctx context.Context, lastResults PrivateLinkServicesForO365ManagementActivityAPIDescriptionListResult) (result PrivateLinkServicesForO365ManagementActivityAPIDescriptionListResult, err error) {
	req, err := lastResults.privateLinkServicesForO365ManagementActivityAPIDescriptionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPIClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPIClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPIClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client PrivateLinkServicesForO365ManagementActivityAPIClient) ListComplete(ctx context.Context) (result PrivateLinkServicesForO365ManagementActivityAPIDescriptionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkServicesForO365ManagementActivityAPIClient.List")
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

// ListByResourceGroup get all the service instances in a resource group.
// Parameters:
// resourceGroupName - the name of the resource group that contains the service instance.
func (client PrivateLinkServicesForO365ManagementActivityAPIClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result PrivateLinkServicesForO365ManagementActivityAPIDescriptionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkServicesForO365ManagementActivityAPIClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.plsfomaadlr.Response.Response != nil {
				sc = result.plsfomaadlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPIClient", "ListByResourceGroup", err.Error())
	}

	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPIClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.plsfomaadlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPIClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.plsfomaadlr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPIClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.plsfomaadlr.hasNextLink() && result.plsfomaadlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client PrivateLinkServicesForO365ManagementActivityAPIClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-01-11"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForO365ManagementActivityAPI", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateLinkServicesForO365ManagementActivityAPIClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client PrivateLinkServicesForO365ManagementActivityAPIClient) ListByResourceGroupResponder(resp *http.Response) (result PrivateLinkServicesForO365ManagementActivityAPIDescriptionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client PrivateLinkServicesForO365ManagementActivityAPIClient) listByResourceGroupNextResults(ctx context.Context, lastResults PrivateLinkServicesForO365ManagementActivityAPIDescriptionListResult) (result PrivateLinkServicesForO365ManagementActivityAPIDescriptionListResult, err error) {
	req, err := lastResults.privateLinkServicesForO365ManagementActivityAPIDescriptionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPIClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPIClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPIClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client PrivateLinkServicesForO365ManagementActivityAPIClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result PrivateLinkServicesForO365ManagementActivityAPIDescriptionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkServicesForO365ManagementActivityAPIClient.ListByResourceGroup")
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

// Update update the metadata of a privateLinkServicesForO365ManagementActivityAPI instance.
// Parameters:
// resourceGroupName - the name of the resource group that contains the service instance.
// resourceName - the name of the service instance.
// servicePatchDescription - the service instance metadata and security metadata.
func (client PrivateLinkServicesForO365ManagementActivityAPIClient) Update(ctx context.Context, resourceGroupName string, resourceName string, servicePatchDescription ServicesPatchDescription) (result PrivateLinkServicesForO365ManagementActivityAPIUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkServicesForO365ManagementActivityAPIClient.Update")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPIClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, resourceName, servicePatchDescription)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPIClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPIClient", "Update", nil, "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client PrivateLinkServicesForO365ManagementActivityAPIClient) UpdatePreparer(ctx context.Context, resourceGroupName string, resourceName string, servicePatchDescription ServicesPatchDescription) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-01-11"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForO365ManagementActivityAPI/{resourceName}", pathParameters),
		autorest.WithJSON(servicePatchDescription),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateLinkServicesForO365ManagementActivityAPIClient) UpdateSender(req *http.Request) (future PrivateLinkServicesForO365ManagementActivityAPIUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client PrivateLinkServicesForO365ManagementActivityAPIClient) (plsfomaad PrivateLinkServicesForO365ManagementActivityAPIDescription, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPIUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPIUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		plsfomaad.Response.Response, err = future.GetResult(sender)
		if plsfomaad.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPIUpdateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && plsfomaad.Response.Response.StatusCode != http.StatusNoContent {
			plsfomaad, err = client.UpdateResponder(plsfomaad.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPIUpdateFuture", "Result", plsfomaad.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client PrivateLinkServicesForO365ManagementActivityAPIClient) UpdateResponder(resp *http.Response) (result PrivateLinkServicesForO365ManagementActivityAPIDescription, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
