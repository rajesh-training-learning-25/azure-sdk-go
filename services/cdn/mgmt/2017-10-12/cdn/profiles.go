package cdn

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
	"net/http"
)

// ProfilesClient is the use these APIs to manage Azure CDN resources through the Azure Resource Manager. You must make
// sure that requests made to these resources are secure.
type ProfilesClient struct {
	BaseClient
}

// NewProfilesClient creates an instance of the ProfilesClient client.
func NewProfilesClient(subscriptionID string, resourceGroupName string) ProfilesClient {
	return NewProfilesClientWithBaseURI(DefaultBaseURI, subscriptionID, resourceGroupName)
}

// NewProfilesClientWithBaseURI creates an instance of the ProfilesClient client.
func NewProfilesClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string) ProfilesClient {
	return ProfilesClient{NewWithBaseURI(baseURI, subscriptionID, resourceGroupName)}
}

// Create creates a new CDN profile with a profile name under the specified subscription and resource group.
//
// profileName is name of the CDN profile which is unique within the resource group. profile is profile properties
// needed to create a new profile.
func (client ProfilesClient) Create(ctx context.Context, profileName string, profile Profile) (result ProfilesCreateFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.ResourceGroupName,
			Constraints: []validation.Constraint{{Target: "client.ResourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "client.ResourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "client.ResourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: profile,
			Constraints: []validation.Constraint{{Target: "profile.Sku", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cdn.ProfilesClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, profileName, profile)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ProfilesClient) CreatePreparer(ctx context.Context, profileName string, profile Profile) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-10-12"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}", pathParameters),
		autorest.WithJSON(profile),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ProfilesClient) CreateSender(req *http.Request) (future ProfilesCreateFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted))
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ProfilesClient) CreateResponder(resp *http.Response) (result Profile, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes an existing CDN profile with the specified parameters. Deleting a profile will result in the deletion
// of all of the sub-resources including endpoints, origins and custom domains.
//
// profileName is name of the CDN profile which is unique within the resource group.
func (client ProfilesClient) Delete(ctx context.Context, profileName string) (result ProfilesDeleteFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.ResourceGroupName,
			Constraints: []validation.Constraint{{Target: "client.ResourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "client.ResourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "client.ResourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cdn.ProfilesClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, profileName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ProfilesClient) DeletePreparer(ctx context.Context, profileName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-10-12"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ProfilesClient) DeleteSender(req *http.Request) (future ProfilesDeleteFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent))
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ProfilesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GenerateSsoURI generates a dynamic SSO URI used to sign in to the CDN supplemental portal. Supplemnetal portal is
// used to configure advanced feature capabilities that are not yet available in the Azure portal, such as core reports
// in a standard profile; rules engine, advanced HTTP reports, and real-time stats and alerts in a premium profile. The
// SSO URI changes approximately every 10 minutes.
//
// profileName is name of the CDN profile which is unique within the resource group.
func (client ProfilesClient) GenerateSsoURI(ctx context.Context, profileName string) (result SsoURI, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.ResourceGroupName,
			Constraints: []validation.Constraint{{Target: "client.ResourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "client.ResourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "client.ResourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cdn.ProfilesClient", "GenerateSsoURI", err.Error())
	}

	req, err := client.GenerateSsoURIPreparer(ctx, profileName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "GenerateSsoURI", nil, "Failure preparing request")
		return
	}

	resp, err := client.GenerateSsoURISender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "GenerateSsoURI", resp, "Failure sending request")
		return
	}

	result, err = client.GenerateSsoURIResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "GenerateSsoURI", resp, "Failure responding to request")
	}

	return
}

// GenerateSsoURIPreparer prepares the GenerateSsoURI request.
func (client ProfilesClient) GenerateSsoURIPreparer(ctx context.Context, profileName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-10-12"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/generateSsoUri", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GenerateSsoURISender sends the GenerateSsoURI request. The method will close the
// http.Response Body if it receives an error.
func (client ProfilesClient) GenerateSsoURISender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GenerateSsoURIResponder handles the response to the GenerateSsoURI request. The method always
// closes the http.Response Body.
func (client ProfilesClient) GenerateSsoURIResponder(resp *http.Response) (result SsoURI, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets a CDN profile with the specified profile name under the specified subscription and resource group.
//
// profileName is name of the CDN profile which is unique within the resource group.
func (client ProfilesClient) Get(ctx context.Context, profileName string) (result Profile, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.ResourceGroupName,
			Constraints: []validation.Constraint{{Target: "client.ResourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "client.ResourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "client.ResourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cdn.ProfilesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, profileName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ProfilesClient) GetPreparer(ctx context.Context, profileName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-10-12"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ProfilesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ProfilesClient) GetResponder(resp *http.Response) (result Profile, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all of the CDN profiles within an Azure subscription.
func (client ProfilesClient) List(ctx context.Context) (result ProfileListResultPage, err error) {
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.plr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "List", resp, "Failure sending request")
		return
	}

	result.plr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ProfilesClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-10-12"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Cdn/profiles", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ProfilesClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ProfilesClient) ListResponder(resp *http.Response) (result ProfileListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ProfilesClient) listNextResults(lastResults ProfileListResult) (result ProfileListResult, err error) {
	req, err := lastResults.profileListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "cdn.ProfilesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "cdn.ProfilesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProfilesClient) ListComplete(ctx context.Context) (result ProfileListResultIterator, err error) {
	result.page, err = client.List(ctx)
	return
}

// ListByResourceGroup lists all of the CDN profiles within a resource group.
func (client ProfilesClient) ListByResourceGroup(ctx context.Context) (result ProfileListResultPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.ResourceGroupName,
			Constraints: []validation.Constraint{{Target: "client.ResourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "client.ResourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "client.ResourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cdn.ProfilesClient", "ListByResourceGroup", err.Error())
	}

	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.plr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.plr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ProfilesClient) ListByResourceGroupPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-10-12"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ProfilesClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ProfilesClient) ListByResourceGroupResponder(resp *http.Response) (result ProfileListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client ProfilesClient) listByResourceGroupNextResults(lastResults ProfileListResult) (result ProfileListResult, err error) {
	req, err := lastResults.profileListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "cdn.ProfilesClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "cdn.ProfilesClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProfilesClient) ListByResourceGroupComplete(ctx context.Context) (result ProfileListResultIterator, err error) {
	result.page, err = client.ListByResourceGroup(ctx)
	return
}

// ListResourceUsage checks the quota and actual usage of endpoints under the given CDN profile.
//
// profileName is name of the CDN profile which is unique within the resource group.
func (client ProfilesClient) ListResourceUsage(ctx context.Context, profileName string) (result ResourceUsageListResultPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.ResourceGroupName,
			Constraints: []validation.Constraint{{Target: "client.ResourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "client.ResourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "client.ResourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cdn.ProfilesClient", "ListResourceUsage", err.Error())
	}

	result.fn = client.listResourceUsageNextResults
	req, err := client.ListResourceUsagePreparer(ctx, profileName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "ListResourceUsage", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListResourceUsageSender(req)
	if err != nil {
		result.rulr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "ListResourceUsage", resp, "Failure sending request")
		return
	}

	result.rulr, err = client.ListResourceUsageResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "ListResourceUsage", resp, "Failure responding to request")
	}

	return
}

// ListResourceUsagePreparer prepares the ListResourceUsage request.
func (client ProfilesClient) ListResourceUsagePreparer(ctx context.Context, profileName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-10-12"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/checkResourceUsage", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListResourceUsageSender sends the ListResourceUsage request. The method will close the
// http.Response Body if it receives an error.
func (client ProfilesClient) ListResourceUsageSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResourceUsageResponder handles the response to the ListResourceUsage request. The method always
// closes the http.Response Body.
func (client ProfilesClient) ListResourceUsageResponder(resp *http.Response) (result ResourceUsageListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listResourceUsageNextResults retrieves the next set of results, if any.
func (client ProfilesClient) listResourceUsageNextResults(lastResults ResourceUsageListResult) (result ResourceUsageListResult, err error) {
	req, err := lastResults.resourceUsageListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "cdn.ProfilesClient", "listResourceUsageNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListResourceUsageSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "cdn.ProfilesClient", "listResourceUsageNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResourceUsageResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "listResourceUsageNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListResourceUsageComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProfilesClient) ListResourceUsageComplete(ctx context.Context, profileName string) (result ResourceUsageListResultIterator, err error) {
	result.page, err = client.ListResourceUsage(ctx, profileName)
	return
}

// ListSupportedOptimizationTypes gets the supported optimization types for the current profile. A user can create an
// endpoint with an optimization type from the listed values.
//
// profileName is name of the CDN profile which is unique within the resource group.
func (client ProfilesClient) ListSupportedOptimizationTypes(ctx context.Context, profileName string) (result SupportedOptimizationTypesListResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.ResourceGroupName,
			Constraints: []validation.Constraint{{Target: "client.ResourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "client.ResourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "client.ResourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cdn.ProfilesClient", "ListSupportedOptimizationTypes", err.Error())
	}

	req, err := client.ListSupportedOptimizationTypesPreparer(ctx, profileName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "ListSupportedOptimizationTypes", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSupportedOptimizationTypesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "ListSupportedOptimizationTypes", resp, "Failure sending request")
		return
	}

	result, err = client.ListSupportedOptimizationTypesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "ListSupportedOptimizationTypes", resp, "Failure responding to request")
	}

	return
}

// ListSupportedOptimizationTypesPreparer prepares the ListSupportedOptimizationTypes request.
func (client ProfilesClient) ListSupportedOptimizationTypesPreparer(ctx context.Context, profileName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-10-12"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/getSupportedOptimizationTypes", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSupportedOptimizationTypesSender sends the ListSupportedOptimizationTypes request. The method will close the
// http.Response Body if it receives an error.
func (client ProfilesClient) ListSupportedOptimizationTypesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListSupportedOptimizationTypesResponder handles the response to the ListSupportedOptimizationTypes request. The method always
// closes the http.Response Body.
func (client ProfilesClient) ListSupportedOptimizationTypesResponder(resp *http.Response) (result SupportedOptimizationTypesListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates an existing CDN profile with the specified profile name under the specified subscription and resource
// group.
//
// profileName is name of the CDN profile which is unique within the resource group. profileUpdateParameters is
// profile properties needed to update an existing profile.
func (client ProfilesClient) Update(ctx context.Context, profileName string, profileUpdateParameters ProfileUpdateParameters) (result ProfilesUpdateFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.ResourceGroupName,
			Constraints: []validation.Constraint{{Target: "client.ResourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "client.ResourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "client.ResourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cdn.ProfilesClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, profileName, profileUpdateParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ProfilesClient) UpdatePreparer(ctx context.Context, profileName string, profileUpdateParameters ProfileUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-10-12"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}", pathParameters),
		autorest.WithJSON(profileUpdateParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ProfilesClient) UpdateSender(req *http.Request) (future ProfilesUpdateFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ProfilesClient) UpdateResponder(resp *http.Response) (result Profile, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
