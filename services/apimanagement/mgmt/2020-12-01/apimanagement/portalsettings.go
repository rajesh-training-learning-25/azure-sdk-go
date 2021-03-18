package apimanagement

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

// PortalSettingsClient is the apiManagement Client
type PortalSettingsClient struct {
	BaseClient
}

// NewPortalSettingsClient creates an instance of the PortalSettingsClient client.
func NewPortalSettingsClient(subscriptionID string) PortalSettingsClient {
	return NewPortalSettingsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPortalSettingsClientWithBaseURI creates an instance of the PortalSettingsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewPortalSettingsClientWithBaseURI(baseURI string, subscriptionID string) PortalSettingsClient {
	return PortalSettingsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListByService lists a collection of portalsettings defined within a service instance..
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
func (client PortalSettingsClient) ListByService(ctx context.Context, resourceGroupName string, serviceName string) (result PortalSettingsCollection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PortalSettingsClient.ListByService")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.PortalSettingsClient", "ListByService", err.Error())
	}

	req, err := client.ListByServicePreparer(ctx, resourceGroupName, serviceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.PortalSettingsClient", "ListByService", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.PortalSettingsClient", "ListByService", resp, "Failure sending request")
		return
	}

	result, err = client.ListByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.PortalSettingsClient", "ListByService", resp, "Failure responding to request")
		return
	}

	return
}

// ListByServicePreparer prepares the ListByService request.
func (client PortalSettingsClient) ListByServicePreparer(ctx context.Context, resourceGroupName string, serviceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByServiceSender sends the ListByService request. The method will close the
// http.Response Body if it receives an error.
func (client PortalSettingsClient) ListByServiceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByServiceResponder handles the response to the ListByService request. The method always
// closes the http.Response Body.
func (client PortalSettingsClient) ListByServiceResponder(resp *http.Response) (result PortalSettingsCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
