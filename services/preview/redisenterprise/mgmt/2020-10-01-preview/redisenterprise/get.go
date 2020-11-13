package redisenterprise

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

// GetClient is the REST API for managing Redis Enterprise resources in Azure.
type GetClient struct {
	BaseClient
}

// NewGetClient creates an instance of the GetClient client.
func NewGetClient(subscriptionID string) GetClient {
	return NewGetClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewGetClientWithBaseURI creates an instance of the GetClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewGetClientWithBaseURI(baseURI string, subscriptionID string) GetClient {
	return GetClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// OperationStatusMethod gets the status of operation.
// Parameters:
// location - the region the operation is in.
// operationID - the operation's unique identifier.
func (client GetClient) OperationStatusMethod(ctx context.Context, location string, operationID string) (result OperationStatus, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GetClient.OperationStatusMethod")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.OperationStatusMethodPreparer(ctx, location, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.GetClient", "OperationStatusMethod", nil, "Failure preparing request")
		return
	}

	resp, err := client.OperationStatusMethodSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "redisenterprise.GetClient", "OperationStatusMethod", resp, "Failure sending request")
		return
	}

	result, err = client.OperationStatusMethodResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.GetClient", "OperationStatusMethod", resp, "Failure responding to request")
	}

	return
}

// OperationStatusMethodPreparer prepares the OperationStatusMethod request.
func (client GetClient) OperationStatusMethodPreparer(ctx context.Context, location string, operationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"operationId":    autorest.Encode("path", operationID),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Cache/locations/{location}/operationsStatus/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// OperationStatusMethodSender sends the OperationStatusMethod request. The method will close the
// http.Response Body if it receives an error.
func (client GetClient) OperationStatusMethodSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// OperationStatusMethodResponder handles the response to the OperationStatusMethod request. The method always
// closes the http.Response Body.
func (client GetClient) OperationStatusMethodResponder(resp *http.Response) (result OperationStatus, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
