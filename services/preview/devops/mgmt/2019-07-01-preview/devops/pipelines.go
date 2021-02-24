package devops

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

// PipelinesClient is the azure DevOps Resource Provider
type PipelinesClient struct {
	BaseClient
}

// NewPipelinesClient creates an instance of the PipelinesClient client.
func NewPipelinesClient(subscriptionID string) PipelinesClient {
	return NewPipelinesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPipelinesClientWithBaseURI creates an instance of the PipelinesClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewPipelinesClientWithBaseURI(baseURI string, subscriptionID string) PipelinesClient {
	return PipelinesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates an Azure Pipeline.
// Parameters:
// resourceGroupName - name of the resource group within the Azure subscription.
// pipelineName - the name of the Azure Pipeline resource in ARM.
// createOperationParameters - the request payload to create the Azure Pipeline.
func (client PipelinesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, pipelineName string, createOperationParameters Pipeline) (result PipelinesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PipelinesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: createOperationParameters,
			Constraints: []validation.Constraint{{Target: "createOperationParameters.PipelineProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "createOperationParameters.PipelineProperties.Organization", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "createOperationParameters.PipelineProperties.Organization.Name", Name: validation.Null, Rule: true, Chain: nil}}},
					{Target: "createOperationParameters.PipelineProperties.Project", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "createOperationParameters.PipelineProperties.Project.Name", Name: validation.Null, Rule: true, Chain: nil}}},
					{Target: "createOperationParameters.PipelineProperties.BootstrapConfiguration", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "createOperationParameters.PipelineProperties.BootstrapConfiguration.Repository", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "createOperationParameters.PipelineProperties.BootstrapConfiguration.Repository.ID", Name: validation.Null, Rule: true, Chain: nil},
								{Target: "createOperationParameters.PipelineProperties.BootstrapConfiguration.Repository.DefaultBranch", Name: validation.Null, Rule: true, Chain: nil},
								{Target: "createOperationParameters.PipelineProperties.BootstrapConfiguration.Repository.Authorization", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "createOperationParameters.PipelineProperties.BootstrapConfiguration.Repository.Authorization.AuthorizationType", Name: validation.Null, Rule: true, Chain: nil}}},
							}},
							{Target: "createOperationParameters.PipelineProperties.BootstrapConfiguration.Template", Name: validation.Null, Rule: true,
								Chain: []validation.Constraint{{Target: "createOperationParameters.PipelineProperties.BootstrapConfiguration.Template.ID", Name: validation.Null, Rule: true, Chain: nil}}},
						}},
				}}}}}); err != nil {
		return result, validation.NewError("devops.PipelinesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, pipelineName, createOperationParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devops.PipelinesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devops.PipelinesClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client PipelinesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, pipelineName string, createOperationParameters Pipeline) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"pipelineName":      autorest.Encode("path", pipelineName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevOps/pipelines/{pipelineName}", pathParameters),
		autorest.WithJSON(createOperationParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client PipelinesClient) CreateOrUpdateSender(req *http.Request) (future PipelinesCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client PipelinesClient) (p Pipeline, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "devops.PipelinesCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("devops.PipelinesCreateOrUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		p.Response.Response, err = future.GetResult(sender)
		if p.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "devops.PipelinesCreateOrUpdateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && p.Response.Response.StatusCode != http.StatusNoContent {
			p, err = client.CreateOrUpdateResponder(p.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "devops.PipelinesCreateOrUpdateFuture", "Result", p.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client PipelinesClient) CreateOrUpdateResponder(resp *http.Response) (result Pipeline, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes an Azure Pipeline.
// Parameters:
// resourceGroupName - name of the resource group within the Azure subscription.
// pipelineName - the name of the Azure Pipeline resource.
func (client PipelinesClient) Delete(ctx context.Context, resourceGroupName string, pipelineName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PipelinesClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, pipelineName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devops.PipelinesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "devops.PipelinesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devops.PipelinesClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client PipelinesClient) DeletePreparer(ctx context.Context, resourceGroupName string, pipelineName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"pipelineName":      autorest.Encode("path", pipelineName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevOps/pipelines/{pipelineName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client PipelinesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client PipelinesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets an existing Azure Pipeline.
// Parameters:
// resourceGroupName - name of the resource group within the Azure subscription.
// pipelineName - the name of the Azure Pipeline resource in ARM.
func (client PipelinesClient) Get(ctx context.Context, resourceGroupName string, pipelineName string) (result Pipeline, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PipelinesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, pipelineName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devops.PipelinesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "devops.PipelinesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devops.PipelinesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client PipelinesClient) GetPreparer(ctx context.Context, resourceGroupName string, pipelineName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"pipelineName":      autorest.Encode("path", pipelineName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevOps/pipelines/{pipelineName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PipelinesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PipelinesClient) GetResponder(resp *http.Response) (result Pipeline, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup lists all Azure Pipelines under the specified resource group.
// Parameters:
// resourceGroupName - name of the resource group within the Azure subscription.
func (client PipelinesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result PipelineListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PipelinesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.plr.Response.Response != nil {
				sc = result.plr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devops.PipelinesClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.plr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "devops.PipelinesClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.plr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devops.PipelinesClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.plr.hasNextLink() && result.plr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client PipelinesClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevOps/pipelines", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client PipelinesClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client PipelinesClient) ListByResourceGroupResponder(resp *http.Response) (result PipelineListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client PipelinesClient) listByResourceGroupNextResults(ctx context.Context, lastResults PipelineListResult) (result PipelineListResult, err error) {
	req, err := lastResults.pipelineListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "devops.PipelinesClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "devops.PipelinesClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devops.PipelinesClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client PipelinesClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result PipelineListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PipelinesClient.ListByResourceGroup")
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

// ListBySubscription lists all Azure Pipelines under the specified subscription.
func (client PipelinesClient) ListBySubscription(ctx context.Context) (result PipelineListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PipelinesClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.plr.Response.Response != nil {
				sc = result.plr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devops.PipelinesClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.plr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "devops.PipelinesClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.plr, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devops.PipelinesClient", "ListBySubscription", resp, "Failure responding to request")
		return
	}
	if result.plr.hasNextLink() && result.plr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client PipelinesClient) ListBySubscriptionPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DevOps/pipelines", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client PipelinesClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client PipelinesClient) ListBySubscriptionResponder(resp *http.Response) (result PipelineListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionNextResults retrieves the next set of results, if any.
func (client PipelinesClient) listBySubscriptionNextResults(ctx context.Context, lastResults PipelineListResult) (result PipelineListResult, err error) {
	req, err := lastResults.pipelineListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "devops.PipelinesClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "devops.PipelinesClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devops.PipelinesClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client PipelinesClient) ListBySubscriptionComplete(ctx context.Context) (result PipelineListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PipelinesClient.ListBySubscription")
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

// Update updates the properties of an Azure Pipeline. Currently, only tags can be updated.
// Parameters:
// resourceGroupName - name of the resource group within the Azure subscription.
// pipelineName - the name of the Azure Pipeline resource.
// updateOperationParameters - the request payload containing the properties to update in the Azure Pipeline.
func (client PipelinesClient) Update(ctx context.Context, resourceGroupName string, pipelineName string, updateOperationParameters PipelineUpdateParameters) (result Pipeline, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PipelinesClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, pipelineName, updateOperationParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devops.PipelinesClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "devops.PipelinesClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devops.PipelinesClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client PipelinesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, pipelineName string, updateOperationParameters PipelineUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"pipelineName":      autorest.Encode("path", pipelineName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevOps/pipelines/{pipelineName}", pathParameters),
		autorest.WithJSON(updateOperationParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client PipelinesClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client PipelinesClient) UpdateResponder(resp *http.Response) (result Pipeline, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
