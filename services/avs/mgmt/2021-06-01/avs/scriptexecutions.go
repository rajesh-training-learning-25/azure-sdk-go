package avs

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

// ScriptExecutionsClient is the azure VMware Solution API
type ScriptExecutionsClient struct {
	BaseClient
}

// NewScriptExecutionsClient creates an instance of the ScriptExecutionsClient client.
func NewScriptExecutionsClient(subscriptionID string) ScriptExecutionsClient {
	return NewScriptExecutionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewScriptExecutionsClientWithBaseURI creates an instance of the ScriptExecutionsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewScriptExecutionsClientWithBaseURI(baseURI string, subscriptionID string) ScriptExecutionsClient {
	return ScriptExecutionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate sends the create or update request.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// privateCloudName - the name of the private cloud.
// scriptExecutionName - name of the user-invoked script execution resource
// scriptExecution - a script running in the private cloud
func (client ScriptExecutionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, privateCloudName string, scriptExecutionName string, scriptExecution ScriptExecution) (result ScriptExecutionsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ScriptExecutionsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
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
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: scriptExecution,
			Constraints: []validation.Constraint{{Target: "scriptExecution.ScriptExecutionProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "scriptExecution.ScriptExecutionProperties.Timeout", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("avs.ScriptExecutionsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, privateCloudName, scriptExecutionName, scriptExecution)
	if err != nil {
		err = autorest.NewErrorWithError(err, "avs.ScriptExecutionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "avs.ScriptExecutionsClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ScriptExecutionsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, privateCloudName string, scriptExecutionName string, scriptExecution ScriptExecution) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateCloudName":    autorest.Encode("path", privateCloudName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"scriptExecutionName": autorest.Encode("path", scriptExecutionName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/scriptExecutions/{scriptExecutionName}", pathParameters),
		autorest.WithJSON(scriptExecution),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ScriptExecutionsClient) CreateOrUpdateSender(req *http.Request) (future ScriptExecutionsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ScriptExecutionsClient) CreateOrUpdateResponder(resp *http.Response) (result ScriptExecution, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete sends the delete request.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// privateCloudName - name of the private cloud
// scriptExecutionName - name of the user-invoked script execution resource
func (client ScriptExecutionsClient) Delete(ctx context.Context, resourceGroupName string, privateCloudName string, scriptExecutionName string) (result ScriptExecutionsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ScriptExecutionsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
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
		return result, validation.NewError("avs.ScriptExecutionsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, privateCloudName, scriptExecutionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "avs.ScriptExecutionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "avs.ScriptExecutionsClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ScriptExecutionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, privateCloudName string, scriptExecutionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateCloudName":    autorest.Encode("path", privateCloudName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"scriptExecutionName": autorest.Encode("path", scriptExecutionName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/scriptExecutions/{scriptExecutionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ScriptExecutionsClient) DeleteSender(req *http.Request) (future ScriptExecutionsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ScriptExecutionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get sends the get request.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// privateCloudName - name of the private cloud
// scriptExecutionName - name of the user-invoked script execution resource
func (client ScriptExecutionsClient) Get(ctx context.Context, resourceGroupName string, privateCloudName string, scriptExecutionName string) (result ScriptExecution, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ScriptExecutionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
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
		return result, validation.NewError("avs.ScriptExecutionsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, privateCloudName, scriptExecutionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "avs.ScriptExecutionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "avs.ScriptExecutionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "avs.ScriptExecutionsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ScriptExecutionsClient) GetPreparer(ctx context.Context, resourceGroupName string, privateCloudName string, scriptExecutionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateCloudName":    autorest.Encode("path", privateCloudName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"scriptExecutionName": autorest.Encode("path", scriptExecutionName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/scriptExecutions/{scriptExecutionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ScriptExecutionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ScriptExecutionsClient) GetResponder(resp *http.Response) (result ScriptExecution, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetExecutionLogs return the logs for a script execution resource
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// privateCloudName - name of the private cloud
// scriptExecutionName - name of the user-invoked script execution resource
// scriptOutputStreamType - name of the desired output stream to return. If not provided, will return all. An
// empty array will return nothing
func (client ScriptExecutionsClient) GetExecutionLogs(ctx context.Context, resourceGroupName string, privateCloudName string, scriptExecutionName string, scriptOutputStreamType []ScriptOutputStreamType) (result ScriptExecution, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ScriptExecutionsClient.GetExecutionLogs")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
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
		return result, validation.NewError("avs.ScriptExecutionsClient", "GetExecutionLogs", err.Error())
	}

	req, err := client.GetExecutionLogsPreparer(ctx, resourceGroupName, privateCloudName, scriptExecutionName, scriptOutputStreamType)
	if err != nil {
		err = autorest.NewErrorWithError(err, "avs.ScriptExecutionsClient", "GetExecutionLogs", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetExecutionLogsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "avs.ScriptExecutionsClient", "GetExecutionLogs", resp, "Failure sending request")
		return
	}

	result, err = client.GetExecutionLogsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "avs.ScriptExecutionsClient", "GetExecutionLogs", resp, "Failure responding to request")
		return
	}

	return
}

// GetExecutionLogsPreparer prepares the GetExecutionLogs request.
func (client ScriptExecutionsClient) GetExecutionLogsPreparer(ctx context.Context, resourceGroupName string, privateCloudName string, scriptExecutionName string, scriptOutputStreamType []ScriptOutputStreamType) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateCloudName":    autorest.Encode("path", privateCloudName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"scriptExecutionName": autorest.Encode("path", scriptExecutionName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/scriptExecutions/{scriptExecutionName}/getExecutionLogs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if scriptOutputStreamType != nil && len(scriptOutputStreamType) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(scriptOutputStreamType))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetExecutionLogsSender sends the GetExecutionLogs request. The method will close the
// http.Response Body if it receives an error.
func (client ScriptExecutionsClient) GetExecutionLogsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetExecutionLogsResponder handles the response to the GetExecutionLogs request. The method always
// closes the http.Response Body.
func (client ScriptExecutionsClient) GetExecutionLogsResponder(resp *http.Response) (result ScriptExecution, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List sends the list request.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// privateCloudName - name of the private cloud
func (client ScriptExecutionsClient) List(ctx context.Context, resourceGroupName string, privateCloudName string) (result ScriptExecutionsListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ScriptExecutionsClient.List")
		defer func() {
			sc := -1
			if result.sel.Response.Response != nil {
				sc = result.sel.Response.Response.StatusCode
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
		return result, validation.NewError("avs.ScriptExecutionsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, privateCloudName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "avs.ScriptExecutionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.sel.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "avs.ScriptExecutionsClient", "List", resp, "Failure sending request")
		return
	}

	result.sel, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "avs.ScriptExecutionsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.sel.hasNextLink() && result.sel.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ScriptExecutionsClient) ListPreparer(ctx context.Context, resourceGroupName string, privateCloudName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateCloudName":  autorest.Encode("path", privateCloudName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/scriptExecutions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ScriptExecutionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ScriptExecutionsClient) ListResponder(resp *http.Response) (result ScriptExecutionsList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ScriptExecutionsClient) listNextResults(ctx context.Context, lastResults ScriptExecutionsList) (result ScriptExecutionsList, err error) {
	req, err := lastResults.scriptExecutionsListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "avs.ScriptExecutionsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "avs.ScriptExecutionsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "avs.ScriptExecutionsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ScriptExecutionsClient) ListComplete(ctx context.Context, resourceGroupName string, privateCloudName string) (result ScriptExecutionsListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ScriptExecutionsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, privateCloudName)
	return
}
