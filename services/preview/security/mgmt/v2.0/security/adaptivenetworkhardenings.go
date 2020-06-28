package security

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

// AdaptiveNetworkHardeningsClient is the API spec for Microsoft.Security (Azure Security Center) resource provider
type AdaptiveNetworkHardeningsClient struct {
	BaseClient
}

// NewAdaptiveNetworkHardeningsClient creates an instance of the AdaptiveNetworkHardeningsClient client.
func NewAdaptiveNetworkHardeningsClient(subscriptionID string, ascLocation string) AdaptiveNetworkHardeningsClient {
	return NewAdaptiveNetworkHardeningsClientWithBaseURI(DefaultBaseURI, subscriptionID, ascLocation)
}

// NewAdaptiveNetworkHardeningsClientWithBaseURI creates an instance of the AdaptiveNetworkHardeningsClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewAdaptiveNetworkHardeningsClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) AdaptiveNetworkHardeningsClient {
	return AdaptiveNetworkHardeningsClient{NewWithBaseURI(baseURI, subscriptionID, ascLocation)}
}

// Enforce enforces the given rules on the NSG(s) listed in the request
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// resourceNamespace - the Namespace of the resource.
// resourceType - the type of the resource.
// resourceName - name of the resource.
// adaptiveNetworkHardeningResourceName - the name of the Adaptive Network Hardening resource.
func (client AdaptiveNetworkHardeningsClient) Enforce(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, adaptiveNetworkHardeningResourceName string, body AdaptiveNetworkHardeningEnforceRequest) (result AdaptiveNetworkHardeningsEnforceFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AdaptiveNetworkHardeningsClient.Enforce")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.Rules", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "body.NetworkSecurityGroups", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.AdaptiveNetworkHardeningsClient", "Enforce", err.Error())
	}

	req, err := client.EnforcePreparer(ctx, resourceGroupName, resourceNamespace, resourceType, resourceName, adaptiveNetworkHardeningResourceName, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AdaptiveNetworkHardeningsClient", "Enforce", nil, "Failure preparing request")
		return
	}

	result, err = client.EnforceSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AdaptiveNetworkHardeningsClient", "Enforce", result.Response(), "Failure sending request")
		return
	}

	return
}

// EnforcePreparer prepares the Enforce request.
func (client AdaptiveNetworkHardeningsClient) EnforcePreparer(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, adaptiveNetworkHardeningResourceName string, body AdaptiveNetworkHardeningEnforceRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"adaptiveNetworkHardeningEnforceAction": autorest.Encode("path", "enforce"),
		"adaptiveNetworkHardeningResourceName":  autorest.Encode("path", adaptiveNetworkHardeningResourceName),
		"resourceGroupName":                     autorest.Encode("path", resourceGroupName),
		"resourceName":                          autorest.Encode("path", resourceName),
		"resourceNamespace":                     autorest.Encode("path", resourceNamespace),
		"resourceType":                          autorest.Encode("path", resourceType),
		"subscriptionId":                        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceNamespace}/{resourceType}/{resourceName}/providers/Microsoft.Security/adaptiveNetworkHardenings/{adaptiveNetworkHardeningResourceName}/{adaptiveNetworkHardeningEnforceAction}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// EnforceSender sends the Enforce request. The method will close the
// http.Response Body if it receives an error.
func (client AdaptiveNetworkHardeningsClient) EnforceSender(req *http.Request) (future AdaptiveNetworkHardeningsEnforceFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// EnforceResponder handles the response to the Enforce request. The method always
// closes the http.Response Body.
func (client AdaptiveNetworkHardeningsClient) EnforceResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a single Adaptive Network Hardening resource
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// resourceNamespace - the Namespace of the resource.
// resourceType - the type of the resource.
// resourceName - name of the resource.
// adaptiveNetworkHardeningResourceName - the name of the Adaptive Network Hardening resource.
func (client AdaptiveNetworkHardeningsClient) Get(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, adaptiveNetworkHardeningResourceName string) (result AdaptiveNetworkHardening, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AdaptiveNetworkHardeningsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.AdaptiveNetworkHardeningsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, resourceNamespace, resourceType, resourceName, adaptiveNetworkHardeningResourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AdaptiveNetworkHardeningsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.AdaptiveNetworkHardeningsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AdaptiveNetworkHardeningsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client AdaptiveNetworkHardeningsClient) GetPreparer(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, adaptiveNetworkHardeningResourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"adaptiveNetworkHardeningResourceName": autorest.Encode("path", adaptiveNetworkHardeningResourceName),
		"resourceGroupName":                    autorest.Encode("path", resourceGroupName),
		"resourceName":                         autorest.Encode("path", resourceName),
		"resourceNamespace":                    autorest.Encode("path", resourceNamespace),
		"resourceType":                         autorest.Encode("path", resourceType),
		"subscriptionId":                       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceNamespace}/{resourceType}/{resourceName}/providers/Microsoft.Security/adaptiveNetworkHardenings/{adaptiveNetworkHardeningResourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AdaptiveNetworkHardeningsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AdaptiveNetworkHardeningsClient) GetResponder(resp *http.Response) (result AdaptiveNetworkHardening, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByExtendedResource gets a list of Adaptive Network Hardenings resources in scope of an extended resource.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// resourceNamespace - the Namespace of the resource.
// resourceType - the type of the resource.
// resourceName - name of the resource.
func (client AdaptiveNetworkHardeningsClient) ListByExtendedResource(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string) (result AdaptiveNetworkHardeningsListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AdaptiveNetworkHardeningsClient.ListByExtendedResource")
		defer func() {
			sc := -1
			if result.anhl.Response.Response != nil {
				sc = result.anhl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.AdaptiveNetworkHardeningsClient", "ListByExtendedResource", err.Error())
	}

	result.fn = client.listByExtendedResourceNextResults
	req, err := client.ListByExtendedResourcePreparer(ctx, resourceGroupName, resourceNamespace, resourceType, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AdaptiveNetworkHardeningsClient", "ListByExtendedResource", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByExtendedResourceSender(req)
	if err != nil {
		result.anhl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.AdaptiveNetworkHardeningsClient", "ListByExtendedResource", resp, "Failure sending request")
		return
	}

	result.anhl, err = client.ListByExtendedResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AdaptiveNetworkHardeningsClient", "ListByExtendedResource", resp, "Failure responding to request")
	}

	return
}

// ListByExtendedResourcePreparer prepares the ListByExtendedResource request.
func (client AdaptiveNetworkHardeningsClient) ListByExtendedResourcePreparer(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"resourceNamespace": autorest.Encode("path", resourceNamespace),
		"resourceType":      autorest.Encode("path", resourceType),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceNamespace}/{resourceType}/{resourceName}/providers/Microsoft.Security/adaptiveNetworkHardenings", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByExtendedResourceSender sends the ListByExtendedResource request. The method will close the
// http.Response Body if it receives an error.
func (client AdaptiveNetworkHardeningsClient) ListByExtendedResourceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByExtendedResourceResponder handles the response to the ListByExtendedResource request. The method always
// closes the http.Response Body.
func (client AdaptiveNetworkHardeningsClient) ListByExtendedResourceResponder(resp *http.Response) (result AdaptiveNetworkHardeningsList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByExtendedResourceNextResults retrieves the next set of results, if any.
func (client AdaptiveNetworkHardeningsClient) listByExtendedResourceNextResults(ctx context.Context, lastResults AdaptiveNetworkHardeningsList) (result AdaptiveNetworkHardeningsList, err error) {
	req, err := lastResults.adaptiveNetworkHardeningsListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "security.AdaptiveNetworkHardeningsClient", "listByExtendedResourceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByExtendedResourceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "security.AdaptiveNetworkHardeningsClient", "listByExtendedResourceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByExtendedResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AdaptiveNetworkHardeningsClient", "listByExtendedResourceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByExtendedResourceComplete enumerates all values, automatically crossing page boundaries as required.
func (client AdaptiveNetworkHardeningsClient) ListByExtendedResourceComplete(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string) (result AdaptiveNetworkHardeningsListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AdaptiveNetworkHardeningsClient.ListByExtendedResource")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByExtendedResource(ctx, resourceGroupName, resourceNamespace, resourceType, resourceName)
	return
}
