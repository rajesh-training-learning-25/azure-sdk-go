package security

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// AssessmentsMetadataClient is the API spec for Microsoft.Security (Azure Security Center) resource provider
type AssessmentsMetadataClient struct {
	BaseClient
}

// NewAssessmentsMetadataClient creates an instance of the AssessmentsMetadataClient client.
func NewAssessmentsMetadataClient(subscriptionID string, ascLocation string) AssessmentsMetadataClient {
	return NewAssessmentsMetadataClientWithBaseURI(DefaultBaseURI, subscriptionID, ascLocation)
}

// NewAssessmentsMetadataClientWithBaseURI creates an instance of the AssessmentsMetadataClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewAssessmentsMetadataClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) AssessmentsMetadataClient {
	return AssessmentsMetadataClient{NewWithBaseURI(baseURI, subscriptionID, ascLocation)}
}

// Get get metadata information on an assessment type
// Parameters:
// assessmentMetadataName - the Assessment Key - Unique key for the assessment type
func (client AssessmentsMetadataClient) Get(ctx context.Context, assessmentMetadataName string) (result AssessmentMetadata, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssessmentsMetadataClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, assessmentMetadataName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AssessmentsMetadataClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.AssessmentsMetadataClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AssessmentsMetadataClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client AssessmentsMetadataClient) GetPreparer(ctx context.Context, assessmentMetadataName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"assessmentMetadataName": autorest.Encode("path", assessmentMetadataName),
	}

	const APIVersion = "2020-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Security/assessmentMetadata/{assessmentMetadataName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AssessmentsMetadataClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AssessmentsMetadataClient) GetResponder(resp *http.Response) (result AssessmentMetadata, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List get metadata information on all assessment types
func (client AssessmentsMetadataClient) List(ctx context.Context) (result AssessmentMetadataListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssessmentsMetadataClient.List")
		defer func() {
			sc := -1
			if result.aml.Response.Response != nil {
				sc = result.aml.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AssessmentsMetadataClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.aml.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.AssessmentsMetadataClient", "List", resp, "Failure sending request")
		return
	}

	result.aml, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AssessmentsMetadataClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client AssessmentsMetadataClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2020-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Security/assessmentMetadata"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client AssessmentsMetadataClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client AssessmentsMetadataClient) ListResponder(resp *http.Response) (result AssessmentMetadataList, err error) {
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
func (client AssessmentsMetadataClient) listNextResults(ctx context.Context, lastResults AssessmentMetadataList) (result AssessmentMetadataList, err error) {
	req, err := lastResults.assessmentMetadataListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "security.AssessmentsMetadataClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "security.AssessmentsMetadataClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AssessmentsMetadataClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client AssessmentsMetadataClient) ListComplete(ctx context.Context) (result AssessmentMetadataListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssessmentsMetadataClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}
