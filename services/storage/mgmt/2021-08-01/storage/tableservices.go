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

// TableServicesClient is the the Azure Storage Management API.
type TableServicesClient struct {
	BaseClient
}

// NewTableServicesClient creates an instance of the TableServicesClient client.
func NewTableServicesClient(subscriptionID string) TableServicesClient {
	return NewTableServicesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewTableServicesClientWithBaseURI creates an instance of the TableServicesClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewTableServicesClientWithBaseURI(baseURI string, subscriptionID string) TableServicesClient {
	return TableServicesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetServiceProperties gets the properties of a storage account’s Table service, including properties for Storage
// Analytics and CORS (Cross-Origin Resource Sharing) rules.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// accountName - the name of the storage account within the specified resource group. Storage account names
// must be between 3 and 24 characters in length and use numbers and lower-case letters only.
func (client TableServicesClient) GetServiceProperties(ctx context.Context, resourceGroupName string, accountName string) (result TableServiceProperties, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TableServicesClient.GetServiceProperties")
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
		return result, validation.NewError("storage.TableServicesClient", "GetServiceProperties", err.Error())
	}

	req, err := client.GetServicePropertiesPreparer(ctx, resourceGroupName, accountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storage.TableServicesClient", "GetServiceProperties", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetServicePropertiesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storage.TableServicesClient", "GetServiceProperties", resp, "Failure sending request")
		return
	}

	result, err = client.GetServicePropertiesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storage.TableServicesClient", "GetServiceProperties", resp, "Failure responding to request")
		return
	}

	return
}

// GetServicePropertiesPreparer prepares the GetServiceProperties request.
func (client TableServicesClient) GetServicePropertiesPreparer(ctx context.Context, resourceGroupName string, accountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"tableServiceName":  autorest.Encode("path", "default"),
	}

	const APIVersion = "2021-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/tableServices/{tableServiceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetServicePropertiesSender sends the GetServiceProperties request. The method will close the
// http.Response Body if it receives an error.
func (client TableServicesClient) GetServicePropertiesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetServicePropertiesResponder handles the response to the GetServiceProperties request. The method always
// closes the http.Response Body.
func (client TableServicesClient) GetServicePropertiesResponder(resp *http.Response) (result TableServiceProperties, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list all table services for the storage account.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// accountName - the name of the storage account within the specified resource group. Storage account names
// must be between 3 and 24 characters in length and use numbers and lower-case letters only.
func (client TableServicesClient) List(ctx context.Context, resourceGroupName string, accountName string) (result ListTableServices, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TableServicesClient.List")
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
		return result, validation.NewError("storage.TableServicesClient", "List", err.Error())
	}

	req, err := client.ListPreparer(ctx, resourceGroupName, accountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storage.TableServicesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storage.TableServicesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storage.TableServicesClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client TableServicesClient) ListPreparer(ctx context.Context, resourceGroupName string, accountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/tableServices", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client TableServicesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client TableServicesClient) ListResponder(resp *http.Response) (result ListTableServices, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// SetServiceProperties sets the properties of a storage account’s Table service, including properties for Storage
// Analytics and CORS (Cross-Origin Resource Sharing) rules.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// accountName - the name of the storage account within the specified resource group. Storage account names
// must be between 3 and 24 characters in length and use numbers and lower-case letters only.
// parameters - the properties of a storage account’s Table service, only properties for Storage Analytics and
// CORS (Cross-Origin Resource Sharing) rules can be specified.
func (client TableServicesClient) SetServiceProperties(ctx context.Context, resourceGroupName string, accountName string, parameters TableServiceProperties) (result TableServiceProperties, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TableServicesClient.SetServiceProperties")
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
		return result, validation.NewError("storage.TableServicesClient", "SetServiceProperties", err.Error())
	}

	req, err := client.SetServicePropertiesPreparer(ctx, resourceGroupName, accountName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storage.TableServicesClient", "SetServiceProperties", nil, "Failure preparing request")
		return
	}

	resp, err := client.SetServicePropertiesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storage.TableServicesClient", "SetServiceProperties", resp, "Failure sending request")
		return
	}

	result, err = client.SetServicePropertiesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storage.TableServicesClient", "SetServiceProperties", resp, "Failure responding to request")
		return
	}

	return
}

// SetServicePropertiesPreparer prepares the SetServiceProperties request.
func (client TableServicesClient) SetServicePropertiesPreparer(ctx context.Context, resourceGroupName string, accountName string, parameters TableServiceProperties) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"tableServiceName":  autorest.Encode("path", "default"),
	}

	const APIVersion = "2021-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/tableServices/{tableServiceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SetServicePropertiesSender sends the SetServiceProperties request. The method will close the
// http.Response Body if it receives an error.
func (client TableServicesClient) SetServicePropertiesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// SetServicePropertiesResponder handles the response to the SetServiceProperties request. The method always
// closes the http.Response Body.
func (client TableServicesClient) SetServicePropertiesResponder(resp *http.Response) (result TableServiceProperties, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
