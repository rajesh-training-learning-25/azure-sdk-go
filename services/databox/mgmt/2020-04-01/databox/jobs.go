package databox

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

// JobsClient is the client for the Jobs methods of the Databox service.
type JobsClient struct {
	BaseClient
}

// NewJobsClient creates an instance of the JobsClient client.
func NewJobsClient(subscriptionID string) JobsClient {
	return NewJobsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewJobsClientWithBaseURI creates an instance of the JobsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewJobsClientWithBaseURI(baseURI string, subscriptionID string) JobsClient {
	return JobsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// BookShipmentPickUp book shipment pick up.
// Parameters:
// resourceGroupName - the Resource Group Name
// jobName - the name of the job Resource within the specified resource group. job names must be between 3 and
// 24 characters in length and use any alphanumeric and underscore only
// shipmentPickUpRequest - details of shipment pick up request.
func (client JobsClient) BookShipmentPickUp(ctx context.Context, resourceGroupName string, jobName string, shipmentPickUpRequest ShipmentPickUpRequest) (result ShipmentPickUpResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.BookShipmentPickUp")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: jobName,
			Constraints: []validation.Constraint{{Target: "jobName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "jobName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "jobName", Name: validation.Pattern, Rule: `^[-\w\.]+$`, Chain: nil}}},
		{TargetValue: shipmentPickUpRequest,
			Constraints: []validation.Constraint{{Target: "shipmentPickUpRequest.StartTime", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "shipmentPickUpRequest.EndTime", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "shipmentPickUpRequest.ShipmentLocation", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("databox.JobsClient", "BookShipmentPickUp", err.Error())
	}

	req, err := client.BookShipmentPickUpPreparer(ctx, resourceGroupName, jobName, shipmentPickUpRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databox.JobsClient", "BookShipmentPickUp", nil, "Failure preparing request")
		return
	}

	resp, err := client.BookShipmentPickUpSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "databox.JobsClient", "BookShipmentPickUp", resp, "Failure sending request")
		return
	}

	result, err = client.BookShipmentPickUpResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databox.JobsClient", "BookShipmentPickUp", resp, "Failure responding to request")
		return
	}

	return
}

// BookShipmentPickUpPreparer prepares the BookShipmentPickUp request.
func (client JobsClient) BookShipmentPickUpPreparer(ctx context.Context, resourceGroupName string, jobName string, shipmentPickUpRequest ShipmentPickUpRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBox/jobs/{jobName}/bookShipmentPickUp", pathParameters),
		autorest.WithJSON(shipmentPickUpRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// BookShipmentPickUpSender sends the BookShipmentPickUp request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) BookShipmentPickUpSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// BookShipmentPickUpResponder handles the response to the BookShipmentPickUp request. The method always
// closes the http.Response Body.
func (client JobsClient) BookShipmentPickUpResponder(resp *http.Response) (result ShipmentPickUpResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Cancel cancelJob.
// Parameters:
// resourceGroupName - the Resource Group Name
// jobName - the name of the job Resource within the specified resource group. job names must be between 3 and
// 24 characters in length and use any alphanumeric and underscore only
// cancellationReason - reason for cancellation.
func (client JobsClient) Cancel(ctx context.Context, resourceGroupName string, jobName string, cancellationReason CancellationReason) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.Cancel")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: jobName,
			Constraints: []validation.Constraint{{Target: "jobName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "jobName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "jobName", Name: validation.Pattern, Rule: `^[-\w\.]+$`, Chain: nil}}},
		{TargetValue: cancellationReason,
			Constraints: []validation.Constraint{{Target: "cancellationReason.Reason", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("databox.JobsClient", "Cancel", err.Error())
	}

	req, err := client.CancelPreparer(ctx, resourceGroupName, jobName, cancellationReason)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databox.JobsClient", "Cancel", nil, "Failure preparing request")
		return
	}

	resp, err := client.CancelSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "databox.JobsClient", "Cancel", resp, "Failure sending request")
		return
	}

	result, err = client.CancelResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databox.JobsClient", "Cancel", resp, "Failure responding to request")
		return
	}

	return
}

// CancelPreparer prepares the Cancel request.
func (client JobsClient) CancelPreparer(ctx context.Context, resourceGroupName string, jobName string, cancellationReason CancellationReason) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBox/jobs/{jobName}/cancel", pathParameters),
		autorest.WithJSON(cancellationReason),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CancelSender sends the Cancel request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) CancelSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CancelResponder handles the response to the Cancel request. The method always
// closes the http.Response Body.
func (client JobsClient) CancelResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Create creates a new job with the specified parameters. Existing job cannot be updated with this API and should
// instead be updated with the Update job API.
// Parameters:
// resourceGroupName - the Resource Group Name
// jobName - the name of the job Resource within the specified resource group. job names must be between 3 and
// 24 characters in length and use any alphanumeric and underscore only
// jobResource - job details from request body.
func (client JobsClient) Create(ctx context.Context, resourceGroupName string, jobName string, jobResource JobResource) (result JobsCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.Create")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: jobName,
			Constraints: []validation.Constraint{{Target: "jobName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "jobName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "jobName", Name: validation.Pattern, Rule: `^[-\w\.]+$`, Chain: nil}}},
		{TargetValue: jobResource,
			Constraints: []validation.Constraint{{Target: "jobResource.JobProperties", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("databox.JobsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, jobName, jobResource)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databox.JobsClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databox.JobsClient", "Create", nil, "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client JobsClient) CreatePreparer(ctx context.Context, resourceGroupName string, jobName string, jobResource JobResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	jobResource.Name = nil
	jobResource.ID = nil
	jobResource.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBox/jobs/{jobName}", pathParameters),
		autorest.WithJSON(jobResource),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) CreateSender(req *http.Request) (future JobsCreateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
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

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client JobsClient) CreateResponder(resp *http.Response) (result JobResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a job.
// Parameters:
// resourceGroupName - the Resource Group Name
// jobName - the name of the job Resource within the specified resource group. job names must be between 3 and
// 24 characters in length and use any alphanumeric and underscore only
func (client JobsClient) Delete(ctx context.Context, resourceGroupName string, jobName string) (result JobsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: jobName,
			Constraints: []validation.Constraint{{Target: "jobName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "jobName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "jobName", Name: validation.Pattern, Rule: `^[-\w\.]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("databox.JobsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, jobName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databox.JobsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databox.JobsClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client JobsClient) DeletePreparer(ctx context.Context, resourceGroupName string, jobName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBox/jobs/{jobName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) DeleteSender(req *http.Request) (future JobsDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
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
func (client JobsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets information about the specified job.
// Parameters:
// resourceGroupName - the Resource Group Name
// jobName - the name of the job Resource within the specified resource group. job names must be between 3 and
// 24 characters in length and use any alphanumeric and underscore only
// expand - $expand is supported on details parameter for job, which provides details on the job stages.
func (client JobsClient) Get(ctx context.Context, resourceGroupName string, jobName string, expand string) (result JobResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: jobName,
			Constraints: []validation.Constraint{{Target: "jobName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "jobName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "jobName", Name: validation.Pattern, Rule: `^[-\w\.]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("databox.JobsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, jobName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databox.JobsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "databox.JobsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databox.JobsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client JobsClient) GetPreparer(ctx context.Context, resourceGroupName string, jobName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBox/jobs/{jobName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client JobsClient) GetResponder(resp *http.Response) (result JobResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all the jobs available under the subscription.
// Parameters:
// skipToken - $skipToken is supported on Get list of jobs, which provides the next page in the list of jobs.
func (client JobsClient) List(ctx context.Context, skipToken string) (result JobResourceListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.List")
		defer func() {
			sc := -1
			if result.jrl.Response.Response != nil {
				sc = result.jrl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databox.JobsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.jrl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "databox.JobsClient", "List", resp, "Failure sending request")
		return
	}

	result.jrl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databox.JobsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.jrl.hasNextLink() && result.jrl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client JobsClient) ListPreparer(ctx context.Context, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DataBox/jobs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client JobsClient) ListResponder(resp *http.Response) (result JobResourceList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client JobsClient) listNextResults(ctx context.Context, lastResults JobResourceList) (result JobResourceList, err error) {
	req, err := lastResults.jobResourceListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "databox.JobsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "databox.JobsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databox.JobsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client JobsClient) ListComplete(ctx context.Context, skipToken string) (result JobResourceListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, skipToken)
	return
}

// ListByResourceGroup lists all the jobs available under the given resource group.
// Parameters:
// resourceGroupName - the Resource Group Name
// skipToken - $skipToken is supported on Get list of jobs, which provides the next page in the list of jobs.
func (client JobsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, skipToken string) (result JobResourceListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.jrl.Response.Response != nil {
				sc = result.jrl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databox.JobsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.jrl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "databox.JobsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.jrl, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databox.JobsClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.jrl.hasNextLink() && result.jrl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client JobsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBox/jobs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client JobsClient) ListByResourceGroupResponder(resp *http.Response) (result JobResourceList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client JobsClient) listByResourceGroupNextResults(ctx context.Context, lastResults JobResourceList) (result JobResourceList, err error) {
	req, err := lastResults.jobResourceListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "databox.JobsClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "databox.JobsClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databox.JobsClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client JobsClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, skipToken string) (result JobResourceListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName, skipToken)
	return
}

// ListCredentials this method gets the unencrypted secrets related to the job.
// Parameters:
// resourceGroupName - the Resource Group Name
// jobName - the name of the job Resource within the specified resource group. job names must be between 3 and
// 24 characters in length and use any alphanumeric and underscore only
func (client JobsClient) ListCredentials(ctx context.Context, resourceGroupName string, jobName string) (result UnencryptedCredentialsList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.ListCredentials")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: jobName,
			Constraints: []validation.Constraint{{Target: "jobName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "jobName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "jobName", Name: validation.Pattern, Rule: `^[-\w\.]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("databox.JobsClient", "ListCredentials", err.Error())
	}

	req, err := client.ListCredentialsPreparer(ctx, resourceGroupName, jobName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databox.JobsClient", "ListCredentials", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListCredentialsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "databox.JobsClient", "ListCredentials", resp, "Failure sending request")
		return
	}

	result, err = client.ListCredentialsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databox.JobsClient", "ListCredentials", resp, "Failure responding to request")
		return
	}

	return
}

// ListCredentialsPreparer prepares the ListCredentials request.
func (client JobsClient) ListCredentialsPreparer(ctx context.Context, resourceGroupName string, jobName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBox/jobs/{jobName}/listCredentials", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListCredentialsSender sends the ListCredentials request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) ListCredentialsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListCredentialsResponder handles the response to the ListCredentials request. The method always
// closes the http.Response Body.
func (client JobsClient) ListCredentialsResponder(resp *http.Response) (result UnencryptedCredentialsList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates the properties of an existing job.
// Parameters:
// resourceGroupName - the Resource Group Name
// jobName - the name of the job Resource within the specified resource group. job names must be between 3 and
// 24 characters in length and use any alphanumeric and underscore only
// jobResourceUpdateParameter - job update parameters from request body.
// ifMatch - defines the If-Match condition. The patch will be performed only if the ETag of the job on the
// server matches this value.
func (client JobsClient) Update(ctx context.Context, resourceGroupName string, jobName string, jobResourceUpdateParameter JobResourceUpdateParameter, ifMatch string) (result JobsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.Update")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: jobName,
			Constraints: []validation.Constraint{{Target: "jobName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "jobName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "jobName", Name: validation.Pattern, Rule: `^[-\w\.]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("databox.JobsClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, jobName, jobResourceUpdateParameter, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databox.JobsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databox.JobsClient", "Update", nil, "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client JobsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, jobName string, jobResourceUpdateParameter JobResourceUpdateParameter, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBox/jobs/{jobName}", pathParameters),
		autorest.WithJSON(jobResourceUpdateParameter),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) UpdateSender(req *http.Request) (future JobsUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
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

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client JobsClient) UpdateResponder(resp *http.Response) (result JobResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
