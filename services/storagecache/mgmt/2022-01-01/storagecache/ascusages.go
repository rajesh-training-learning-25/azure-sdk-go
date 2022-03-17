package storagecache

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

// AscUsagesClient is the a Storage Cache provides scalable caching service for NAS clients, serving data from either
// NFSv3 or Blob at-rest storage (referred to as "Storage Targets"). These operations allow you to manage Caches.
type AscUsagesClient struct {
	BaseClient
}

// NewAscUsagesClient creates an instance of the AscUsagesClient client.
func NewAscUsagesClient(subscriptionID string) AscUsagesClient {
	return NewAscUsagesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAscUsagesClientWithBaseURI creates an instance of the AscUsagesClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewAscUsagesClientWithBaseURI(baseURI string, subscriptionID string) AscUsagesClient {
	return AscUsagesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List gets the quantity used and quota limit for resources
// Parameters:
// location - the name of the region to query for usage information.
func (client AscUsagesClient) List(ctx context.Context, location string) (result ResourceUsagesListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AscUsagesClient.List")
		defer func() {
			sc := -1
			if result.rulr.Response.Response != nil {
				sc = result.rulr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, location)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.AscUsagesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rulr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storagecache.AscUsagesClient", "List", resp, "Failure sending request")
		return
	}

	result.rulr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.AscUsagesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.rulr.hasNextLink() && result.rulr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client AscUsagesClient) ListPreparer(ctx context.Context, location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.StorageCache/locations/{location}/usages", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client AscUsagesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client AscUsagesClient) ListResponder(resp *http.Response) (result ResourceUsagesListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client AscUsagesClient) listNextResults(ctx context.Context, lastResults ResourceUsagesListResult) (result ResourceUsagesListResult, err error) {
	req, err := lastResults.resourceUsagesListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "storagecache.AscUsagesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "storagecache.AscUsagesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.AscUsagesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client AscUsagesClient) ListComplete(ctx context.Context, location string) (result ResourceUsagesListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AscUsagesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, location)
	return
}
