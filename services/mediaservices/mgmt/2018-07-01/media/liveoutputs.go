package media

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

// LiveOutputsClient is the client for the LiveOutputs methods of the Media service.
type LiveOutputsClient struct {
	BaseClient
}

// NewLiveOutputsClient creates an instance of the LiveOutputsClient client.
func NewLiveOutputsClient(subscriptionID string) LiveOutputsClient {
	return NewLiveOutputsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLiveOutputsClientWithBaseURI creates an instance of the LiveOutputsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewLiveOutputsClientWithBaseURI(baseURI string, subscriptionID string) LiveOutputsClient {
	return LiveOutputsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates a Live Output.
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// liveEventName - the name of the Live Event.
// liveOutputName - the name of the Live Output.
// parameters - live Output properties needed for creation.
func (client LiveOutputsClient) Create(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, liveOutputName string, parameters LiveOutput) (result LiveOutputsCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LiveOutputsClient.Create")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: liveEventName,
			Constraints: []validation.Constraint{{Target: "liveEventName", Name: validation.MaxLength, Rule: 32, Chain: nil},
				{Target: "liveEventName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "liveEventName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+(-*[a-zA-Z0-9])*$`, Chain: nil}}},
		{TargetValue: liveOutputName,
			Constraints: []validation.Constraint{{Target: "liveOutputName", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "liveOutputName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "liveOutputName", Name: validation.Pattern, Rule: `^([a-zA-Z0-9])+(-*[a-zA-Z0-9])*$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.LiveOutputProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.LiveOutputProperties.AssetName", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.LiveOutputProperties.ArchiveWindowLength", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("media.LiveOutputsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, accountName, liveEventName, liveOutputName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.LiveOutputsClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.LiveOutputsClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client LiveOutputsClient) CreatePreparer(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, liveOutputName string, parameters LiveOutput) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"liveEventName":     autorest.Encode("path", liveEventName),
		"liveOutputName":    autorest.Encode("path", liveOutputName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}/liveOutputs/{liveOutputName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client LiveOutputsClient) CreateSender(req *http.Request) (future LiveOutputsCreateFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client LiveOutputsClient) CreateResponder(resp *http.Response) (result LiveOutput, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a Live Output.
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// liveEventName - the name of the Live Event.
// liveOutputName - the name of the Live Output.
func (client LiveOutputsClient) Delete(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, liveOutputName string) (result LiveOutputsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LiveOutputsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: liveEventName,
			Constraints: []validation.Constraint{{Target: "liveEventName", Name: validation.MaxLength, Rule: 32, Chain: nil},
				{Target: "liveEventName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "liveEventName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+(-*[a-zA-Z0-9])*$`, Chain: nil}}},
		{TargetValue: liveOutputName,
			Constraints: []validation.Constraint{{Target: "liveOutputName", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "liveOutputName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "liveOutputName", Name: validation.Pattern, Rule: `^([a-zA-Z0-9])+(-*[a-zA-Z0-9])*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("media.LiveOutputsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, accountName, liveEventName, liveOutputName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.LiveOutputsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.LiveOutputsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client LiveOutputsClient) DeletePreparer(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, liveOutputName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"liveEventName":     autorest.Encode("path", liveEventName),
		"liveOutputName":    autorest.Encode("path", liveOutputName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}/liveOutputs/{liveOutputName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client LiveOutputsClient) DeleteSender(req *http.Request) (future LiveOutputsDeleteFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client LiveOutputsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a Live Output.
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// liveEventName - the name of the Live Event.
// liveOutputName - the name of the Live Output.
func (client LiveOutputsClient) Get(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, liveOutputName string) (result LiveOutput, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LiveOutputsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: liveEventName,
			Constraints: []validation.Constraint{{Target: "liveEventName", Name: validation.MaxLength, Rule: 32, Chain: nil},
				{Target: "liveEventName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "liveEventName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+(-*[a-zA-Z0-9])*$`, Chain: nil}}},
		{TargetValue: liveOutputName,
			Constraints: []validation.Constraint{{Target: "liveOutputName", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "liveOutputName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "liveOutputName", Name: validation.Pattern, Rule: `^([a-zA-Z0-9])+(-*[a-zA-Z0-9])*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("media.LiveOutputsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, accountName, liveEventName, liveOutputName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.LiveOutputsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "media.LiveOutputsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.LiveOutputsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client LiveOutputsClient) GetPreparer(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, liveOutputName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"liveEventName":     autorest.Encode("path", liveEventName),
		"liveOutputName":    autorest.Encode("path", liveOutputName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}/liveOutputs/{liveOutputName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client LiveOutputsClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client LiveOutputsClient) GetResponder(resp *http.Response) (result LiveOutput, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists the Live Outputs in the Live Event.
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// liveEventName - the name of the Live Event.
func (client LiveOutputsClient) List(ctx context.Context, resourceGroupName string, accountName string, liveEventName string) (result LiveOutputListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LiveOutputsClient.List")
		defer func() {
			sc := -1
			if result.lolr.Response.Response != nil {
				sc = result.lolr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: liveEventName,
			Constraints: []validation.Constraint{{Target: "liveEventName", Name: validation.MaxLength, Rule: 32, Chain: nil},
				{Target: "liveEventName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "liveEventName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+(-*[a-zA-Z0-9])*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("media.LiveOutputsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, accountName, liveEventName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.LiveOutputsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.lolr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "media.LiveOutputsClient", "List", resp, "Failure sending request")
		return
	}

	result.lolr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.LiveOutputsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client LiveOutputsClient) ListPreparer(ctx context.Context, resourceGroupName string, accountName string, liveEventName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"liveEventName":     autorest.Encode("path", liveEventName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}/liveOutputs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client LiveOutputsClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client LiveOutputsClient) ListResponder(resp *http.Response) (result LiveOutputListResult, err error) {
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
func (client LiveOutputsClient) listNextResults(ctx context.Context, lastResults LiveOutputListResult) (result LiveOutputListResult, err error) {
	req, err := lastResults.liveOutputListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "media.LiveOutputsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "media.LiveOutputsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.LiveOutputsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client LiveOutputsClient) ListComplete(ctx context.Context, resourceGroupName string, accountName string, liveEventName string) (result LiveOutputListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LiveOutputsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, accountName, liveEventName)
	return
}
