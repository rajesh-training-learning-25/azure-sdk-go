package mixedreality

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

// ObjectAnchorsAccountsClient is the mixed Reality Client
type ObjectAnchorsAccountsClient struct {
	BaseClient
}

// NewObjectAnchorsAccountsClient creates an instance of the ObjectAnchorsAccountsClient client.
func NewObjectAnchorsAccountsClient(subscriptionID string) ObjectAnchorsAccountsClient {
	return NewObjectAnchorsAccountsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewObjectAnchorsAccountsClientWithBaseURI creates an instance of the ObjectAnchorsAccountsClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewObjectAnchorsAccountsClientWithBaseURI(baseURI string, subscriptionID string) ObjectAnchorsAccountsClient {
	return ObjectAnchorsAccountsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creating or Updating an object anchors Account.
// Parameters:
// resourceGroupName - name of an Azure resource group.
// accountName - name of an Mixed Reality Account.
// objectAnchorsAccount - object Anchors Account parameter.
func (client ObjectAnchorsAccountsClient) Create(ctx context.Context, resourceGroupName string, accountName string, objectAnchorsAccount ObjectAnchorsAccount) (result ObjectAnchorsAccount, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ObjectAnchorsAccountsClient.Create")
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
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: objectAnchorsAccount,
			Constraints: []validation.Constraint{{Target: "objectAnchorsAccount.Sku", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "objectAnchorsAccount.Sku.Name", Name: validation.Null, Rule: true, Chain: nil}}},
				{Target: "objectAnchorsAccount.Kind", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "objectAnchorsAccount.Kind.Name", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("mixedreality.ObjectAnchorsAccountsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, accountName, objectAnchorsAccount)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mixedreality.ObjectAnchorsAccountsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "mixedreality.ObjectAnchorsAccountsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mixedreality.ObjectAnchorsAccountsClient", "Create", resp, "Failure responding to request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ObjectAnchorsAccountsClient) CreatePreparer(ctx context.Context, resourceGroupName string, accountName string, objectAnchorsAccount ObjectAnchorsAccount) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	objectAnchorsAccount.SystemData = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MixedReality/objectAnchorsAccounts/{accountName}", pathParameters),
		autorest.WithJSON(objectAnchorsAccount),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ObjectAnchorsAccountsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ObjectAnchorsAccountsClient) CreateResponder(resp *http.Response) (result ObjectAnchorsAccount, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete an Object Anchors Account.
// Parameters:
// resourceGroupName - name of an Azure resource group.
// accountName - name of an Mixed Reality Account.
func (client ObjectAnchorsAccountsClient) Delete(ctx context.Context, resourceGroupName string, accountName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ObjectAnchorsAccountsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
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
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("mixedreality.ObjectAnchorsAccountsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, accountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mixedreality.ObjectAnchorsAccountsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "mixedreality.ObjectAnchorsAccountsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mixedreality.ObjectAnchorsAccountsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ObjectAnchorsAccountsClient) DeletePreparer(ctx context.Context, resourceGroupName string, accountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MixedReality/objectAnchorsAccounts/{accountName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ObjectAnchorsAccountsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ObjectAnchorsAccountsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieve an Object Anchors Account.
// Parameters:
// resourceGroupName - name of an Azure resource group.
// accountName - name of an Mixed Reality Account.
func (client ObjectAnchorsAccountsClient) Get(ctx context.Context, resourceGroupName string, accountName string) (result ObjectAnchorsAccount, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ObjectAnchorsAccountsClient.Get")
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
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("mixedreality.ObjectAnchorsAccountsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, accountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mixedreality.ObjectAnchorsAccountsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "mixedreality.ObjectAnchorsAccountsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mixedreality.ObjectAnchorsAccountsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ObjectAnchorsAccountsClient) GetPreparer(ctx context.Context, resourceGroupName string, accountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MixedReality/objectAnchorsAccounts/{accountName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ObjectAnchorsAccountsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ObjectAnchorsAccountsClient) GetResponder(resp *http.Response) (result ObjectAnchorsAccount, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup list Resources by Resource Group
// Parameters:
// resourceGroupName - name of an Azure resource group.
func (client ObjectAnchorsAccountsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result ObjectAnchorsAccountPagePage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ObjectAnchorsAccountsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.oaap.Response.Response != nil {
				sc = result.oaap.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("mixedreality.ObjectAnchorsAccountsClient", "ListByResourceGroup", err.Error())
	}

	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mixedreality.ObjectAnchorsAccountsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.oaap.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "mixedreality.ObjectAnchorsAccountsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.oaap, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mixedreality.ObjectAnchorsAccountsClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.oaap.hasNextLink() && result.oaap.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ObjectAnchorsAccountsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MixedReality/objectAnchorsAccounts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ObjectAnchorsAccountsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ObjectAnchorsAccountsClient) ListByResourceGroupResponder(resp *http.Response) (result ObjectAnchorsAccountPage, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client ObjectAnchorsAccountsClient) listByResourceGroupNextResults(ctx context.Context, lastResults ObjectAnchorsAccountPage) (result ObjectAnchorsAccountPage, err error) {
	req, err := lastResults.objectAnchorsAccountPagePreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "mixedreality.ObjectAnchorsAccountsClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "mixedreality.ObjectAnchorsAccountsClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mixedreality.ObjectAnchorsAccountsClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client ObjectAnchorsAccountsClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result ObjectAnchorsAccountPageIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ObjectAnchorsAccountsClient.ListByResourceGroup")
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

// ListBySubscription list Object Anchors Accounts by Subscription
func (client ObjectAnchorsAccountsClient) ListBySubscription(ctx context.Context) (result ObjectAnchorsAccountPagePage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ObjectAnchorsAccountsClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.oaap.Response.Response != nil {
				sc = result.oaap.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mixedreality.ObjectAnchorsAccountsClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.oaap.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "mixedreality.ObjectAnchorsAccountsClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.oaap, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mixedreality.ObjectAnchorsAccountsClient", "ListBySubscription", resp, "Failure responding to request")
		return
	}
	if result.oaap.hasNextLink() && result.oaap.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client ObjectAnchorsAccountsClient) ListBySubscriptionPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.MixedReality/objectAnchorsAccounts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client ObjectAnchorsAccountsClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client ObjectAnchorsAccountsClient) ListBySubscriptionResponder(resp *http.Response) (result ObjectAnchorsAccountPage, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionNextResults retrieves the next set of results, if any.
func (client ObjectAnchorsAccountsClient) listBySubscriptionNextResults(ctx context.Context, lastResults ObjectAnchorsAccountPage) (result ObjectAnchorsAccountPage, err error) {
	req, err := lastResults.objectAnchorsAccountPagePreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "mixedreality.ObjectAnchorsAccountsClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "mixedreality.ObjectAnchorsAccountsClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mixedreality.ObjectAnchorsAccountsClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client ObjectAnchorsAccountsClient) ListBySubscriptionComplete(ctx context.Context) (result ObjectAnchorsAccountPageIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ObjectAnchorsAccountsClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySubscription(ctx)
	return
}

// ListKeys list Both of the 2 Keys of an object anchors Account
// Parameters:
// resourceGroupName - name of an Azure resource group.
// accountName - name of an Mixed Reality Account.
func (client ObjectAnchorsAccountsClient) ListKeys(ctx context.Context, resourceGroupName string, accountName string) (result AccountKeys, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ObjectAnchorsAccountsClient.ListKeys")
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
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("mixedreality.ObjectAnchorsAccountsClient", "ListKeys", err.Error())
	}

	req, err := client.ListKeysPreparer(ctx, resourceGroupName, accountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mixedreality.ObjectAnchorsAccountsClient", "ListKeys", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListKeysSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "mixedreality.ObjectAnchorsAccountsClient", "ListKeys", resp, "Failure sending request")
		return
	}

	result, err = client.ListKeysResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mixedreality.ObjectAnchorsAccountsClient", "ListKeys", resp, "Failure responding to request")
		return
	}

	return
}

// ListKeysPreparer prepares the ListKeys request.
func (client ObjectAnchorsAccountsClient) ListKeysPreparer(ctx context.Context, resourceGroupName string, accountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MixedReality/objectAnchorsAccounts/{accountName}/listKeys", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListKeysSender sends the ListKeys request. The method will close the
// http.Response Body if it receives an error.
func (client ObjectAnchorsAccountsClient) ListKeysSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListKeysResponder handles the response to the ListKeys request. The method always
// closes the http.Response Body.
func (client ObjectAnchorsAccountsClient) ListKeysResponder(resp *http.Response) (result AccountKeys, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// RegenerateKeys regenerate specified Key of an object anchors Account
// Parameters:
// resourceGroupName - name of an Azure resource group.
// accountName - name of an Mixed Reality Account.
// regenerate - required information for key regeneration.
func (client ObjectAnchorsAccountsClient) RegenerateKeys(ctx context.Context, resourceGroupName string, accountName string, regenerate AccountKeyRegenerateRequest) (result AccountKeys, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ObjectAnchorsAccountsClient.RegenerateKeys")
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
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("mixedreality.ObjectAnchorsAccountsClient", "RegenerateKeys", err.Error())
	}

	req, err := client.RegenerateKeysPreparer(ctx, resourceGroupName, accountName, regenerate)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mixedreality.ObjectAnchorsAccountsClient", "RegenerateKeys", nil, "Failure preparing request")
		return
	}

	resp, err := client.RegenerateKeysSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "mixedreality.ObjectAnchorsAccountsClient", "RegenerateKeys", resp, "Failure sending request")
		return
	}

	result, err = client.RegenerateKeysResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mixedreality.ObjectAnchorsAccountsClient", "RegenerateKeys", resp, "Failure responding to request")
		return
	}

	return
}

// RegenerateKeysPreparer prepares the RegenerateKeys request.
func (client ObjectAnchorsAccountsClient) RegenerateKeysPreparer(ctx context.Context, resourceGroupName string, accountName string, regenerate AccountKeyRegenerateRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MixedReality/objectAnchorsAccounts/{accountName}/regenerateKeys", pathParameters),
		autorest.WithJSON(regenerate),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RegenerateKeysSender sends the RegenerateKeys request. The method will close the
// http.Response Body if it receives an error.
func (client ObjectAnchorsAccountsClient) RegenerateKeysSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// RegenerateKeysResponder handles the response to the RegenerateKeys request. The method always
// closes the http.Response Body.
func (client ObjectAnchorsAccountsClient) RegenerateKeysResponder(resp *http.Response) (result AccountKeys, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updating an Object Anchors Account
// Parameters:
// resourceGroupName - name of an Azure resource group.
// accountName - name of an Mixed Reality Account.
// objectAnchorsAccount - object Anchors Account parameter.
func (client ObjectAnchorsAccountsClient) Update(ctx context.Context, resourceGroupName string, accountName string, objectAnchorsAccount ObjectAnchorsAccount) (result ObjectAnchorsAccount, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ObjectAnchorsAccountsClient.Update")
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
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("mixedreality.ObjectAnchorsAccountsClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, accountName, objectAnchorsAccount)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mixedreality.ObjectAnchorsAccountsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "mixedreality.ObjectAnchorsAccountsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mixedreality.ObjectAnchorsAccountsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ObjectAnchorsAccountsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, accountName string, objectAnchorsAccount ObjectAnchorsAccount) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	objectAnchorsAccount.SystemData = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MixedReality/objectAnchorsAccounts/{accountName}", pathParameters),
		autorest.WithJSON(objectAnchorsAccount),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ObjectAnchorsAccountsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ObjectAnchorsAccountsClient) UpdateResponder(resp *http.Response) (result ObjectAnchorsAccount, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
