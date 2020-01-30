package search

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
	"github.com/satori/go.uuid"
	"net/http"
)

// AdminKeysClient is the client that can be used to manage Azure Cognitive Search services and API keys.
type AdminKeysClient struct {
	BaseClient
}

// NewAdminKeysClient creates an instance of the AdminKeysClient client.
func NewAdminKeysClient(subscriptionID string) AdminKeysClient {
	return NewAdminKeysClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAdminKeysClientWithBaseURI creates an instance of the AdminKeysClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewAdminKeysClientWithBaseURI(baseURI string, subscriptionID string) AdminKeysClient {
	return AdminKeysClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets the primary and secondary admin API keys for the specified Azure Cognitive Search service.
// Parameters:
// resourceGroupName - the name of the resource group within the current subscription. You can obtain this
// value from the Azure Resource Manager API or the portal.
// searchServiceName - the name of the Azure Cognitive Search service associated with the specified resource
// group.
// clientRequestID - a client-generated GUID value that identifies this request. If specified, this will be
// included in response information as a way to track the request.
func (client AdminKeysClient) Get(ctx context.Context, resourceGroupName string, searchServiceName string, clientRequestID *uuid.UUID) (result AdminKeyResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AdminKeysClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, searchServiceName, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.AdminKeysClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "search.AdminKeysClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.AdminKeysClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client AdminKeysClient) GetPreparer(ctx context.Context, resourceGroupName string, searchServiceName string, clientRequestID *uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"searchServiceName": autorest.Encode("path", searchServiceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-19"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/listAdminKeys", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("x-ms-client-request-id", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AdminKeysClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AdminKeysClient) GetResponder(resp *http.Response) (result AdminKeyResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Regenerate regenerates either the primary or secondary admin API key. You can only regenerate one key at a time.
// Parameters:
// resourceGroupName - the name of the resource group within the current subscription. You can obtain this
// value from the Azure Resource Manager API or the portal.
// searchServiceName - the name of the Azure Cognitive Search service associated with the specified resource
// group.
// keyKind - specifies which key to regenerate. Valid values include 'primary' and 'secondary'.
// clientRequestID - a client-generated GUID value that identifies this request. If specified, this will be
// included in response information as a way to track the request.
func (client AdminKeysClient) Regenerate(ctx context.Context, resourceGroupName string, searchServiceName string, keyKind AdminKeyKind, clientRequestID *uuid.UUID) (result AdminKeyResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AdminKeysClient.Regenerate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RegeneratePreparer(ctx, resourceGroupName, searchServiceName, keyKind, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.AdminKeysClient", "Regenerate", nil, "Failure preparing request")
		return
	}

	resp, err := client.RegenerateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "search.AdminKeysClient", "Regenerate", resp, "Failure sending request")
		return
	}

	result, err = client.RegenerateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.AdminKeysClient", "Regenerate", resp, "Failure responding to request")
	}

	return
}

// RegeneratePreparer prepares the Regenerate request.
func (client AdminKeysClient) RegeneratePreparer(ctx context.Context, resourceGroupName string, searchServiceName string, keyKind AdminKeyKind, clientRequestID *uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"keyKind":           autorest.Encode("path", keyKind),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"searchServiceName": autorest.Encode("path", searchServiceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-19"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/regenerateAdminKey/{keyKind}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("x-ms-client-request-id", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RegenerateSender sends the Regenerate request. The method will close the
// http.Response Body if it receives an error.
func (client AdminKeysClient) RegenerateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// RegenerateResponder handles the response to the Regenerate request. The method always
// closes the http.Response Body.
func (client AdminKeysClient) RegenerateResponder(resp *http.Response) (result AdminKeyResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
