package peering

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

// RegisteredAsnsClient is the peering Client
type RegisteredAsnsClient struct {
	BaseClient
}

// NewRegisteredAsnsClient creates an instance of the RegisteredAsnsClient client.
func NewRegisteredAsnsClient(subscriptionID string) RegisteredAsnsClient {
	return NewRegisteredAsnsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRegisteredAsnsClientWithBaseURI creates an instance of the RegisteredAsnsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewRegisteredAsnsClientWithBaseURI(baseURI string, subscriptionID string) RegisteredAsnsClient {
	return RegisteredAsnsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a new registered ASN with the specified name under the given subscription, resource group and
// peering.
// Parameters:
// resourceGroupName - the name of the resource group.
// peeringName - the name of the peering.
// registeredAsnName - the name of the ASN.
// registeredAsn - the properties needed to create a registered ASN.
func (client RegisteredAsnsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, peeringName string, registeredAsnName string, registeredAsn RegisteredAsn) (result RegisteredAsn, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegisteredAsnsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, peeringName, registeredAsnName, registeredAsn)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.RegisteredAsnsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "peering.RegisteredAsnsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.RegisteredAsnsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client RegisteredAsnsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, peeringName string, registeredAsnName string, registeredAsn RegisteredAsn) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"peeringName":       autorest.Encode("path", peeringName),
		"registeredAsnName": autorest.Encode("path", registeredAsnName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}/registeredAsns/{registeredAsnName}", pathParameters),
		autorest.WithJSON(registeredAsn),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client RegisteredAsnsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client RegisteredAsnsClient) CreateOrUpdateResponder(resp *http.Response) (result RegisteredAsn, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes an existing registered ASN with the specified name under the given subscription, resource group and
// peering.
// Parameters:
// resourceGroupName - the name of the resource group.
// peeringName - the name of the peering.
// registeredAsnName - the name of the registered ASN.
func (client RegisteredAsnsClient) Delete(ctx context.Context, resourceGroupName string, peeringName string, registeredAsnName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegisteredAsnsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, peeringName, registeredAsnName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.RegisteredAsnsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "peering.RegisteredAsnsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.RegisteredAsnsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client RegisteredAsnsClient) DeletePreparer(ctx context.Context, resourceGroupName string, peeringName string, registeredAsnName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"peeringName":       autorest.Encode("path", peeringName),
		"registeredAsnName": autorest.Encode("path", registeredAsnName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}/registeredAsns/{registeredAsnName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client RegisteredAsnsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client RegisteredAsnsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets an existing registered ASN with the specified name under the given subscription, resource group and
// peering.
// Parameters:
// resourceGroupName - the name of the resource group.
// peeringName - the name of the peering.
// registeredAsnName - the name of the registered ASN.
func (client RegisteredAsnsClient) Get(ctx context.Context, resourceGroupName string, peeringName string, registeredAsnName string) (result RegisteredAsn, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegisteredAsnsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, peeringName, registeredAsnName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.RegisteredAsnsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "peering.RegisteredAsnsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.RegisteredAsnsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client RegisteredAsnsClient) GetPreparer(ctx context.Context, resourceGroupName string, peeringName string, registeredAsnName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"peeringName":       autorest.Encode("path", peeringName),
		"registeredAsnName": autorest.Encode("path", registeredAsnName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}/registeredAsns/{registeredAsnName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RegisteredAsnsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RegisteredAsnsClient) GetResponder(resp *http.Response) (result RegisteredAsn, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByPeering lists all registered ASNs under the given subscription, resource group and peering.
// Parameters:
// resourceGroupName - the name of the resource group.
// peeringName - the name of the peering.
func (client RegisteredAsnsClient) ListByPeering(ctx context.Context, resourceGroupName string, peeringName string) (result RegisteredAsnListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegisteredAsnsClient.ListByPeering")
		defer func() {
			sc := -1
			if result.ralr.Response.Response != nil {
				sc = result.ralr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByPeeringNextResults
	req, err := client.ListByPeeringPreparer(ctx, resourceGroupName, peeringName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.RegisteredAsnsClient", "ListByPeering", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByPeeringSender(req)
	if err != nil {
		result.ralr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "peering.RegisteredAsnsClient", "ListByPeering", resp, "Failure sending request")
		return
	}

	result.ralr, err = client.ListByPeeringResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.RegisteredAsnsClient", "ListByPeering", resp, "Failure responding to request")
	}

	return
}

// ListByPeeringPreparer prepares the ListByPeering request.
func (client RegisteredAsnsClient) ListByPeeringPreparer(ctx context.Context, resourceGroupName string, peeringName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"peeringName":       autorest.Encode("path", peeringName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}/registeredAsns", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByPeeringSender sends the ListByPeering request. The method will close the
// http.Response Body if it receives an error.
func (client RegisteredAsnsClient) ListByPeeringSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByPeeringResponder handles the response to the ListByPeering request. The method always
// closes the http.Response Body.
func (client RegisteredAsnsClient) ListByPeeringResponder(resp *http.Response) (result RegisteredAsnListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByPeeringNextResults retrieves the next set of results, if any.
func (client RegisteredAsnsClient) listByPeeringNextResults(ctx context.Context, lastResults RegisteredAsnListResult) (result RegisteredAsnListResult, err error) {
	req, err := lastResults.registeredAsnListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "peering.RegisteredAsnsClient", "listByPeeringNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByPeeringSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "peering.RegisteredAsnsClient", "listByPeeringNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByPeeringResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.RegisteredAsnsClient", "listByPeeringNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByPeeringComplete enumerates all values, automatically crossing page boundaries as required.
func (client RegisteredAsnsClient) ListByPeeringComplete(ctx context.Context, resourceGroupName string, peeringName string) (result RegisteredAsnListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegisteredAsnsClient.ListByPeering")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByPeering(ctx, resourceGroupName, peeringName)
	return
}
