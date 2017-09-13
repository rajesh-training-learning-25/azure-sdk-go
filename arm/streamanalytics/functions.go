package streamanalytics

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
// Code generated by Microsoft (R) AutoRest Code Generator 2.2.18.0
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// FunctionsClient is the stream Analytics Client
type FunctionsClient struct {
	ManagementClient
}

// NewFunctionsClient creates an instance of the FunctionsClient client.
func NewFunctionsClient(subscriptionID string) FunctionsClient {
	return NewFunctionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewFunctionsClientWithBaseURI creates an instance of the FunctionsClient client.
func NewFunctionsClientWithBaseURI(baseURI string, subscriptionID string) FunctionsClient {
	return FunctionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrReplace creates a function or replaces an already existing function under an existing streaming job.
//
// function is the definition of the function that will be used to create a new function or replace the existing one
// under the streaming job. resourceGroupName is the name of the resource group that contains the resource. You can
// obtain this value from the Azure Resource Manager API or the portal. jobName is the name of the streaming job.
// functionName is the name of the function. ifMatch is the ETag of the function. Omit this value to always overwrite
// the current function. Specify the last-seen ETag value to prevent accidentally overwritting concurrent changes.
// ifNoneMatch is set to '*' to allow a new function to be created, but to prevent updating an existing function. Other
// values will result in a 412 Pre-condition Failed response.
func (client FunctionsClient) CreateOrReplace(function Function, resourceGroupName string, jobName string, functionName string, ifMatch string, ifNoneMatch string) (result Function, err error) {
	req, err := client.CreateOrReplacePreparer(function, resourceGroupName, jobName, functionName, ifMatch, ifNoneMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "CreateOrReplace", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrReplaceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "CreateOrReplace", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrReplaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "CreateOrReplace", resp, "Failure responding to request")
	}

	return
}

// CreateOrReplacePreparer prepares the CreateOrReplace request.
func (client FunctionsClient) CreateOrReplacePreparer(function Function, resourceGroupName string, jobName string, functionName string, ifMatch string, ifNoneMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"functionName":      autorest.Encode("path", functionName),
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/functions/{functionName}", pathParameters),
		autorest.WithJSON(function),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	if len(ifNoneMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-None-Match", autorest.String(ifNoneMatch)))
	}
	return preparer.Prepare(&http.Request{})
}

// CreateOrReplaceSender sends the CreateOrReplace request. The method will close the
// http.Response Body if it receives an error.
func (client FunctionsClient) CreateOrReplaceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrReplaceResponder handles the response to the CreateOrReplace request. The method always
// closes the http.Response Body.
func (client FunctionsClient) CreateOrReplaceResponder(resp *http.Response) (result Function, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a function from the streaming job.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from the
// Azure Resource Manager API or the portal. jobName is the name of the streaming job. functionName is the name of the
// function.
func (client FunctionsClient) Delete(resourceGroupName string, jobName string, functionName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(resourceGroupName, jobName, functionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client FunctionsClient) DeletePreparer(resourceGroupName string, jobName string, functionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"functionName":      autorest.Encode("path", functionName),
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/functions/{functionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client FunctionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client FunctionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets details about the specified function.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from the
// Azure Resource Manager API or the portal. jobName is the name of the streaming job. functionName is the name of the
// function.
func (client FunctionsClient) Get(resourceGroupName string, jobName string, functionName string) (result Function, err error) {
	req, err := client.GetPreparer(resourceGroupName, jobName, functionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client FunctionsClient) GetPreparer(resourceGroupName string, jobName string, functionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"functionName":      autorest.Encode("path", functionName),
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/functions/{functionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client FunctionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client FunctionsClient) GetResponder(resp *http.Response) (result Function, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByStreamingJob lists all of the functions under the specified streaming job.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from the
// Azure Resource Manager API or the portal. jobName is the name of the streaming job. selectParameter is the $select
// OData query parameter. This is a comma-separated list of structural properties to include in the response, or “*” to
// include all properties. By default, all properties are returned except diagnostics. Currently only accepts '*' as a
// valid value.
func (client FunctionsClient) ListByStreamingJob(resourceGroupName string, jobName string, selectParameter string) (result FunctionListResult, err error) {
	req, err := client.ListByStreamingJobPreparer(resourceGroupName, jobName, selectParameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "ListByStreamingJob", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByStreamingJobSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "ListByStreamingJob", resp, "Failure sending request")
		return
	}

	result, err = client.ListByStreamingJobResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "ListByStreamingJob", resp, "Failure responding to request")
	}

	return
}

// ListByStreamingJobPreparer prepares the ListByStreamingJob request.
func (client FunctionsClient) ListByStreamingJobPreparer(resourceGroupName string, jobName string, selectParameter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/functions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByStreamingJobSender sends the ListByStreamingJob request. The method will close the
// http.Response Body if it receives an error.
func (client FunctionsClient) ListByStreamingJobSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByStreamingJobResponder handles the response to the ListByStreamingJob request. The method always
// closes the http.Response Body.
func (client FunctionsClient) ListByStreamingJobResponder(resp *http.Response) (result FunctionListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByStreamingJobNextResults retrieves the next set of results, if any.
func (client FunctionsClient) ListByStreamingJobNextResults(lastResults FunctionListResult) (result FunctionListResult, err error) {
	req, err := lastResults.FunctionListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "ListByStreamingJob", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByStreamingJobSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "ListByStreamingJob", resp, "Failure sending next results request")
	}

	result, err = client.ListByStreamingJobResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "ListByStreamingJob", resp, "Failure responding to next results request")
	}

	return
}

// ListByStreamingJobComplete gets all elements from the list without paging.
func (client FunctionsClient) ListByStreamingJobComplete(resourceGroupName string, jobName string, selectParameter string, cancel <-chan struct{}) (<-chan Function, <-chan error) {
	resultChan := make(chan Function)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListByStreamingJob(resourceGroupName, jobName, selectParameter)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListByStreamingJobNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}

// RetrieveDefaultDefinition retrieves the default definition of a function based on the parameters specified.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from the
// Azure Resource Manager API or the portal. jobName is the name of the streaming job. functionName is the name of the
// function. functionRetrieveDefaultDefinitionParameters is parameters used to specify the type of function to retrieve
// the default definition for.
func (client FunctionsClient) RetrieveDefaultDefinition(resourceGroupName string, jobName string, functionName string, functionRetrieveDefaultDefinitionParameters *FunctionRetrieveDefaultDefinitionParameters) (result Function, err error) {
	req, err := client.RetrieveDefaultDefinitionPreparer(resourceGroupName, jobName, functionName, functionRetrieveDefaultDefinitionParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "RetrieveDefaultDefinition", nil, "Failure preparing request")
		return
	}

	resp, err := client.RetrieveDefaultDefinitionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "RetrieveDefaultDefinition", resp, "Failure sending request")
		return
	}

	result, err = client.RetrieveDefaultDefinitionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "RetrieveDefaultDefinition", resp, "Failure responding to request")
	}

	return
}

// RetrieveDefaultDefinitionPreparer prepares the RetrieveDefaultDefinition request.
func (client FunctionsClient) RetrieveDefaultDefinitionPreparer(resourceGroupName string, jobName string, functionName string, functionRetrieveDefaultDefinitionParameters *FunctionRetrieveDefaultDefinitionParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"functionName":      autorest.Encode("path", functionName),
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/functions/{functionName}/RetrieveDefaultDefinition", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if functionRetrieveDefaultDefinitionParameters != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(functionRetrieveDefaultDefinitionParameters))
	}
	return preparer.Prepare(&http.Request{})
}

// RetrieveDefaultDefinitionSender sends the RetrieveDefaultDefinition request. The method will close the
// http.Response Body if it receives an error.
func (client FunctionsClient) RetrieveDefaultDefinitionSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// RetrieveDefaultDefinitionResponder handles the response to the RetrieveDefaultDefinition request. The method always
// closes the http.Response Body.
func (client FunctionsClient) RetrieveDefaultDefinitionResponder(resp *http.Response) (result Function, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Test tests if the information provided for a function is valid. This can range from testing the connection to the
// underlying web service behind the function or making sure the function code provided is syntactically correct. This
// method may poll for completion. Polling can be canceled by passing the cancel channel argument. The channel will be
// used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from the
// Azure Resource Manager API or the portal. jobName is the name of the streaming job. functionName is the name of the
// function. function is if the function specified does not already exist, this parameter must contain the full
// function definition intended to be tested. If the function specified already exists, this parameter can be left null
// to test the existing function as is or if specified, the properties specified will overwrite the corresponding
// properties in the existing function (exactly like a PATCH operation) and the resulting function will be tested.
func (client FunctionsClient) Test(resourceGroupName string, jobName string, functionName string, function *Function, cancel <-chan struct{}) (<-chan ResourceTestStatus, <-chan error) {
	resultChan := make(chan ResourceTestStatus, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result ResourceTestStatus
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.TestPreparer(resourceGroupName, jobName, functionName, function, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "Test", nil, "Failure preparing request")
			return
		}

		resp, err := client.TestSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "Test", resp, "Failure sending request")
			return
		}

		result, err = client.TestResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "Test", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// TestPreparer prepares the Test request.
func (client FunctionsClient) TestPreparer(resourceGroupName string, jobName string, functionName string, function *Function, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"functionName":      autorest.Encode("path", functionName),
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/functions/{functionName}/test", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if function != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(function))
	}
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// TestSender sends the Test request. The method will close the
// http.Response Body if it receives an error.
func (client FunctionsClient) TestSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// TestResponder handles the response to the Test request. The method always
// closes the http.Response Body.
func (client FunctionsClient) TestResponder(resp *http.Response) (result ResourceTestStatus, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates an existing function under an existing streaming job. This can be used to partially update (ie.
// update one or two properties) a function without affecting the rest the job or function definition.
//
// function is a function object. The properties specified here will overwrite the corresponding properties in the
// existing function (ie. Those properties will be updated). Any properties that are set to null here will mean that
// the corresponding property in the existing function will remain the same and not change as a result of this PATCH
// operation. resourceGroupName is the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal. jobName is the name of the streaming job. functionName is the
// name of the function. ifMatch is the ETag of the function. Omit this value to always overwrite the current function.
// Specify the last-seen ETag value to prevent accidentally overwritting concurrent changes.
func (client FunctionsClient) Update(function Function, resourceGroupName string, jobName string, functionName string, ifMatch string) (result Function, err error) {
	req, err := client.UpdatePreparer(function, resourceGroupName, jobName, functionName, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client FunctionsClient) UpdatePreparer(function Function, resourceGroupName string, jobName string, functionName string, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"functionName":      autorest.Encode("path", functionName),
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/functions/{functionName}", pathParameters),
		autorest.WithJSON(function),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare(&http.Request{})
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client FunctionsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client FunctionsClient) UpdateResponder(resp *http.Response) (result Function, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
