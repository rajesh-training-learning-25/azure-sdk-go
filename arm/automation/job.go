package automation

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

// JobClient is the automation Client
type JobClient struct {
	ManagementClient
}

// NewJobClient creates an instance of the JobClient client.
func NewJobClient(subscriptionID string) JobClient {
	return NewJobClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewJobClientWithBaseURI creates an instance of the JobClient client.
func NewJobClientWithBaseURI(baseURI string, subscriptionID string) JobClient {
	return JobClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// NewJobClientWithAuthFile creates an instance of the JobClient client.
func NewJobClientWithAuthFile() (JobClient, error) {
	c, err := NewWithAuthFile()
	return JobClient{c}, err
}

// Create create a job of the runbook.
//
// resourceGroupName is the resource group name. automationAccountName is the automation account name. jobID is the job
// id. parameters is the parameters supplied to the create job operation.
func (client JobClient) Create(resourceGroupName string, automationAccountName string, jobID uuid.UUID, parameters JobCreateParameters) (result Job, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.JobCreateProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.JobCreateProperties.Runbook", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "automation.JobClient", "Create")
	}

	req, err := client.CreatePreparer(resourceGroupName, automationAccountName, jobID, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.JobClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client JobClient) CreatePreparer(resourceGroupName string, automationAccountName string, jobID uuid.UUID, parameters JobCreateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"jobId":                 autorest.Encode("path", jobID),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobs/{jobId}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client JobClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client JobClient) CreateResponder(resp *http.Response) (result Job, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get retrieve the job identified by job id.
//
// resourceGroupName is the resource group name. automationAccountName is the automation account name. jobID is the job
// id.
func (client JobClient) Get(resourceGroupName string, automationAccountName string, jobID uuid.UUID) (result Job, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "automation.JobClient", "Get")
	}

	req, err := client.GetPreparer(resourceGroupName, automationAccountName, jobID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.JobClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client JobClient) GetPreparer(resourceGroupName string, automationAccountName string, jobID uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"jobId":                 autorest.Encode("path", jobID),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobs/{jobId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client JobClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client JobClient) GetResponder(resp *http.Response) (result Job, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetOutput retrieve the job output identified by job id.
//
// resourceGroupName is the resource group name. automationAccountName is the automation account name. jobID is the job
// id.
func (client JobClient) GetOutput(resourceGroupName string, automationAccountName string, jobID string) (result ReadCloser, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "automation.JobClient", "GetOutput")
	}

	req, err := client.GetOutputPreparer(resourceGroupName, automationAccountName, jobID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobClient", "GetOutput", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetOutputSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.JobClient", "GetOutput", resp, "Failure sending request")
		return
	}

	result, err = client.GetOutputResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobClient", "GetOutput", resp, "Failure responding to request")
	}

	return
}

// GetOutputPreparer prepares the GetOutput request.
func (client JobClient) GetOutputPreparer(resourceGroupName string, automationAccountName string, jobID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"jobId":                 autorest.Encode("path", jobID),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobs/{jobId}/output", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetOutputSender sends the GetOutput request. The method will close the
// http.Response Body if it receives an error.
func (client JobClient) GetOutputSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetOutputResponder handles the response to the GetOutput request. The method always
// closes the http.Response Body.
func (client JobClient) GetOutputResponder(resp *http.Response) (result ReadCloser, err error) {
	result.Value = &resp.Body
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK))
	result.Response = autorest.Response{Response: resp}
	return
}

// GetRunbookContent retrieve the runbook content of the job identified by job id.
//
// resourceGroupName is the resource group name. automationAccountName is the automation account name. jobID is the job
// id.
func (client JobClient) GetRunbookContent(resourceGroupName string, automationAccountName string, jobID string) (result ReadCloser, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "automation.JobClient", "GetRunbookContent")
	}

	req, err := client.GetRunbookContentPreparer(resourceGroupName, automationAccountName, jobID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobClient", "GetRunbookContent", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetRunbookContentSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.JobClient", "GetRunbookContent", resp, "Failure sending request")
		return
	}

	result, err = client.GetRunbookContentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobClient", "GetRunbookContent", resp, "Failure responding to request")
	}

	return
}

// GetRunbookContentPreparer prepares the GetRunbookContent request.
func (client JobClient) GetRunbookContentPreparer(resourceGroupName string, automationAccountName string, jobID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"jobId":                 autorest.Encode("path", jobID),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobs/{jobId}/runbookContent", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetRunbookContentSender sends the GetRunbookContent request. The method will close the
// http.Response Body if it receives an error.
func (client JobClient) GetRunbookContentSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetRunbookContentResponder handles the response to the GetRunbookContent request. The method always
// closes the http.Response Body.
func (client JobClient) GetRunbookContentResponder(resp *http.Response) (result ReadCloser, err error) {
	result.Value = &resp.Body
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK))
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByAutomationAccount retrieve a list of jobs.
//
// resourceGroupName is the resource group name. automationAccountName is the automation account name. filter is the
// filter to apply on the operation.
func (client JobClient) ListByAutomationAccount(resourceGroupName string, automationAccountName string, filter string) (result JobListResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "automation.JobClient", "ListByAutomationAccount")
	}

	req, err := client.ListByAutomationAccountPreparer(resourceGroupName, automationAccountName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobClient", "ListByAutomationAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByAutomationAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.JobClient", "ListByAutomationAccount", resp, "Failure sending request")
		return
	}

	result, err = client.ListByAutomationAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobClient", "ListByAutomationAccount", resp, "Failure responding to request")
	}

	return
}

// ListByAutomationAccountPreparer prepares the ListByAutomationAccount request.
func (client JobClient) ListByAutomationAccountPreparer(resourceGroupName string, automationAccountName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByAutomationAccountSender sends the ListByAutomationAccount request. The method will close the
// http.Response Body if it receives an error.
func (client JobClient) ListByAutomationAccountSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByAutomationAccountResponder handles the response to the ListByAutomationAccount request. The method always
// closes the http.Response Body.
func (client JobClient) ListByAutomationAccountResponder(resp *http.Response) (result JobListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByAutomationAccountNextResults retrieves the next set of results, if any.
func (client JobClient) ListByAutomationAccountNextResults(lastResults JobListResult) (result JobListResult, err error) {
	req, err := lastResults.JobListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "automation.JobClient", "ListByAutomationAccount", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByAutomationAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "automation.JobClient", "ListByAutomationAccount", resp, "Failure sending next results request")
	}

	result, err = client.ListByAutomationAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobClient", "ListByAutomationAccount", resp, "Failure responding to next results request")
	}

	return
}

// ListByAutomationAccountComplete gets all elements from the list without paging.
func (client JobClient) ListByAutomationAccountComplete(resourceGroupName string, automationAccountName string, filter string, cancel <-chan struct{}) (<-chan Job, <-chan error) {
	resultChan := make(chan Job)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListByAutomationAccount(resourceGroupName, automationAccountName, filter)
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
			list, err = client.ListByAutomationAccountNextResults(list)
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

// Resume resume the job identified by jobId.
//
// resourceGroupName is the resource group name. automationAccountName is the automation account name. jobID is the job
// id.
func (client JobClient) Resume(resourceGroupName string, automationAccountName string, jobID uuid.UUID) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "automation.JobClient", "Resume")
	}

	req, err := client.ResumePreparer(resourceGroupName, automationAccountName, jobID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobClient", "Resume", nil, "Failure preparing request")
		return
	}

	resp, err := client.ResumeSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "automation.JobClient", "Resume", resp, "Failure sending request")
		return
	}

	result, err = client.ResumeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobClient", "Resume", resp, "Failure responding to request")
	}

	return
}

// ResumePreparer prepares the Resume request.
func (client JobClient) ResumePreparer(resourceGroupName string, automationAccountName string, jobID uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"jobId":                 autorest.Encode("path", jobID),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobs/{jobId}/resume", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ResumeSender sends the Resume request. The method will close the
// http.Response Body if it receives an error.
func (client JobClient) ResumeSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ResumeResponder handles the response to the Resume request. The method always
// closes the http.Response Body.
func (client JobClient) ResumeResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Stop stop the job identified by jobId.
//
// resourceGroupName is the resource group name. automationAccountName is the automation account name. jobID is the job
// id.
func (client JobClient) Stop(resourceGroupName string, automationAccountName string, jobID uuid.UUID) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "automation.JobClient", "Stop")
	}

	req, err := client.StopPreparer(resourceGroupName, automationAccountName, jobID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobClient", "Stop", nil, "Failure preparing request")
		return
	}

	resp, err := client.StopSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "automation.JobClient", "Stop", resp, "Failure sending request")
		return
	}

	result, err = client.StopResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobClient", "Stop", resp, "Failure responding to request")
	}

	return
}

// StopPreparer prepares the Stop request.
func (client JobClient) StopPreparer(resourceGroupName string, automationAccountName string, jobID uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"jobId":                 autorest.Encode("path", jobID),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobs/{jobId}/stop", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// StopSender sends the Stop request. The method will close the
// http.Response Body if it receives an error.
func (client JobClient) StopSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// StopResponder handles the response to the Stop request. The method always
// closes the http.Response Body.
func (client JobClient) StopResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Suspend suspend the job identified by jobId.
//
// resourceGroupName is the resource group name. automationAccountName is the automation account name. jobID is the job
// id.
func (client JobClient) Suspend(resourceGroupName string, automationAccountName string, jobID uuid.UUID) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "automation.JobClient", "Suspend")
	}

	req, err := client.SuspendPreparer(resourceGroupName, automationAccountName, jobID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobClient", "Suspend", nil, "Failure preparing request")
		return
	}

	resp, err := client.SuspendSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "automation.JobClient", "Suspend", resp, "Failure sending request")
		return
	}

	result, err = client.SuspendResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobClient", "Suspend", resp, "Failure responding to request")
	}

	return
}

// SuspendPreparer prepares the Suspend request.
func (client JobClient) SuspendPreparer(resourceGroupName string, automationAccountName string, jobID uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"jobId":                 autorest.Encode("path", jobID),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobs/{jobId}/suspend", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// SuspendSender sends the Suspend request. The method will close the
// http.Response Body if it receives an error.
func (client JobClient) SuspendSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// SuspendResponder handles the response to the Suspend request. The method always
// closes the http.Response Body.
func (client JobClient) SuspendResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
