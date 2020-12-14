package postgresql

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

// ServerParametersClient is the the Microsoft Azure management API provides create, read, update, and delete
// functionality for Azure PostgreSQL resources including servers, databases, firewall rules, VNET rules, security
// alert policies, log files and configurations with new business model.
type ServerParametersClient struct {
	BaseClient
}

// NewServerParametersClient creates an instance of the ServerParametersClient client.
func NewServerParametersClient(subscriptionID string) ServerParametersClient {
	return NewServerParametersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewServerParametersClientWithBaseURI creates an instance of the ServerParametersClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewServerParametersClientWithBaseURI(baseURI string, subscriptionID string) ServerParametersClient {
	return ServerParametersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListUpdateConfigurations update a list of configurations in a given server.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// serverName - the name of the server.
// value - the parameters for updating a list of server configuration.
func (client ServerParametersClient) ListUpdateConfigurations(ctx context.Context, resourceGroupName string, serverName string, value ConfigurationListResult) (result ServerParametersListUpdateConfigurationsFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServerParametersClient.ListUpdateConfigurations")
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
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("postgresql.ServerParametersClient", "ListUpdateConfigurations", err.Error())
	}

	req, err := client.ListUpdateConfigurationsPreparer(ctx, resourceGroupName, serverName, value)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresql.ServerParametersClient", "ListUpdateConfigurations", nil, "Failure preparing request")
		return
	}

	result, err = client.ListUpdateConfigurationsSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresql.ServerParametersClient", "ListUpdateConfigurations", result.Response(), "Failure sending request")
		return
	}

	return
}

// ListUpdateConfigurationsPreparer prepares the ListUpdateConfigurations request.
func (client ServerParametersClient) ListUpdateConfigurationsPreparer(ctx context.Context, resourceGroupName string, serverName string, value ConfigurationListResult) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBForPostgreSQL/servers/{serverName}/updateConfigurations", pathParameters),
		autorest.WithJSON(value),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListUpdateConfigurationsSender sends the ListUpdateConfigurations request. The method will close the
// http.Response Body if it receives an error.
func (client ServerParametersClient) ListUpdateConfigurationsSender(req *http.Request) (future ServerParametersListUpdateConfigurationsFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// ListUpdateConfigurationsResponder handles the response to the ListUpdateConfigurations request. The method always
// closes the http.Response Body.
func (client ServerParametersClient) ListUpdateConfigurationsResponder(resp *http.Response) (result ConfigurationListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
