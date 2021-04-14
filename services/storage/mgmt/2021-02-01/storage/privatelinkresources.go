package storage

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
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

// PrivateLinkResourcesClient is the the Azure Storage Management API.
type PrivateLinkResourcesClient struct {
	BaseClient
}

// NewPrivateLinkResourcesClient creates an instance of the PrivateLinkResourcesClient client.
func NewPrivateLinkResourcesClient(subscriptionID string) PrivateLinkResourcesClient {
	return NewPrivateLinkResourcesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPrivateLinkResourcesClientWithBaseURI creates an instance of the PrivateLinkResourcesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewPrivateLinkResourcesClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkResourcesClient {
	return PrivateLinkResourcesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListByStorageAccount gets the private link resources that need to be created for a storage account.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// accountName - the name of the storage account within the specified resource group. Storage account names
// must be between 3 and 24 characters in length and use numbers and lower-case letters only.
func (client PrivateLinkResourcesClient) ListByStorageAccount(ctx context.Context, resourceGroupName string, accountName string) (result PrivateLinkResourceListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkResourcesClient.ListByStorageAccount")
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
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storage.PrivateLinkResourcesClient", "ListByStorageAccount", err.Error())
	}

	req, err := client.ListByStorageAccountPreparer(ctx, resourceGroupName, accountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storage.PrivateLinkResourcesClient", "ListByStorageAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByStorageAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storage.PrivateLinkResourcesClient", "ListByStorageAccount", resp, "Failure sending request")
		return
	}

	result, err = client.ListByStorageAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storage.PrivateLinkResourcesClient", "ListByStorageAccount", resp, "Failure responding to request")
		return
	}

	return
}

// ListByStorageAccountPreparer prepares the ListByStorageAccount request.
func (client PrivateLinkResourcesClient) ListByStorageAccountPreparer(ctx context.Context, resourceGroupName string, accountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/privateLinkResources", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByStorageAccountSender sends the ListByStorageAccount request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateLinkResourcesClient) ListByStorageAccountSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByStorageAccountResponder handles the response to the ListByStorageAccount request. The method always
// closes the http.Response Body.
func (client PrivateLinkResourcesClient) ListByStorageAccountResponder(resp *http.Response) (result PrivateLinkResourceListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
