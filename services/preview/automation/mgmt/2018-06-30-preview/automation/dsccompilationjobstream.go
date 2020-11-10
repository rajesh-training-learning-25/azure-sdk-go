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
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"github.com/gofrs/uuid"
	"net/http"
)

// DscCompilationJobStreamClient is the automation Client
type DscCompilationJobStreamClient struct {
	BaseClient
}

// NewDscCompilationJobStreamClient creates an instance of the DscCompilationJobStreamClient client.
func NewDscCompilationJobStreamClient(subscriptionID string) DscCompilationJobStreamClient {
	return NewDscCompilationJobStreamClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDscCompilationJobStreamClientWithBaseURI creates an instance of the DscCompilationJobStreamClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewDscCompilationJobStreamClientWithBaseURI(baseURI string, subscriptionID string) DscCompilationJobStreamClient {
	return DscCompilationJobStreamClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListByJob retrieve all the job streams for the compilation Job.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// jobID - the job id.
func (client DscCompilationJobStreamClient) ListByJob(ctx context.Context, resourceGroupName string, automationAccountName string, jobID uuid.UUID) (result JobStreamListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DscCompilationJobStreamClient.ListByJob")
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
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.DscCompilationJobStreamClient", "ListByJob", err.Error())
	}

	req, err := client.ListByJobPreparer(ctx, resourceGroupName, automationAccountName, jobID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.DscCompilationJobStreamClient", "ListByJob", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByJobSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.DscCompilationJobStreamClient", "ListByJob", resp, "Failure sending request")
		return
	}

	result, err = client.ListByJobResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.DscCompilationJobStreamClient", "ListByJob", resp, "Failure responding to request")
	}

	return
}

// ListByJobPreparer prepares the ListByJob request.
func (client DscCompilationJobStreamClient) ListByJobPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, jobID uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"jobId":                 autorest.Encode("path", jobID),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/compilationjobs/{jobId}/streams", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByJobSender sends the ListByJob request. The method will close the
// http.Response Body if it receives an error.
func (client DscCompilationJobStreamClient) ListByJobSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByJobResponder handles the response to the ListByJob request. The method always
// closes the http.Response Body.
func (client DscCompilationJobStreamClient) ListByJobResponder(resp *http.Response) (result JobStreamListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
