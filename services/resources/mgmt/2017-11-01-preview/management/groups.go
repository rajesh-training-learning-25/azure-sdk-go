package management

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
	"net/http"
)

// GroupsClient is the the Azure Management Groups API enables consolidation of multiple
// subscriptions/resources into an organizational hierarchy and centrally
// manage access control, policies, alerting and reporting for those resources.
type GroupsClient struct {
	BaseClient
}

// NewGroupsClient creates an instance of the GroupsClient client.
func NewGroupsClient() GroupsClient {
	return NewGroupsClientWithBaseURI(DefaultBaseURI)
}

// NewGroupsClientWithBaseURI creates an instance of the GroupsClient client.
func NewGroupsClientWithBaseURI(baseURI string) GroupsClient {
	return GroupsClient{NewWithBaseURI(baseURI)}
}

// CreateOrUpdate create or update a management group.
// If a management group is already created and a subsequent create request is issued with different properties, the
// management group properties will be updated.
//
// groupID is management Group ID. createGroupRequest is management group creation parameters. cacheControl is
// indicates that the request shouldn't utilize any caches.
func (client GroupsClient) CreateOrUpdate(ctx context.Context, groupID string, createGroupRequest CreateGroupRequest, cacheControl string) (result Group, err error) {
	req, err := client.CreateOrUpdatePreparer(ctx, groupID, createGroupRequest, cacheControl)
	if err != nil {
		err = autorest.NewErrorWithError(err, "management.GroupsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "management.GroupsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "management.GroupsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client GroupsClient) CreateOrUpdatePreparer(ctx context.Context, groupID string, createGroupRequest CreateGroupRequest, cacheControl string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupId": autorest.Encode("path", groupID),
	}

	const APIVersion = "2017-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementGroups/{groupId}", pathParameters),
		autorest.WithJSON(createGroupRequest),
		autorest.WithQueryParameters(queryParameters))
	if len(cacheControl) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Cache-Control", autorest.String(cacheControl)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client GroupsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client GroupsClient) CreateOrUpdateResponder(resp *http.Response) (result Group, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete management group.
// If a management group contains child resources, the request will fail.
//
// groupID is management Group ID. cacheControl is indicates that the request shouldn't utilize any caches.
func (client GroupsClient) Delete(ctx context.Context, groupID string, cacheControl string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(ctx, groupID, cacheControl)
	if err != nil {
		err = autorest.NewErrorWithError(err, "management.GroupsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "management.GroupsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "management.GroupsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client GroupsClient) DeletePreparer(ctx context.Context, groupID string, cacheControl string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupId": autorest.Encode("path", groupID),
	}

	const APIVersion = "2017-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementGroups/{groupId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(cacheControl) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Cache-Control", autorest.String(cacheControl)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client GroupsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client GroupsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get the details of the management group.
//
// groupID is management Group ID. expand is the $expand=children query string parameter allows clients to request
// inclusion of children in the response payload. recurse is the $recurse=true query string parameter allows clients to
// request inclusion of entire hierarchy in the response payload. cacheControl is indicates that the request shouldn't
// utilize any caches.
func (client GroupsClient) Get(ctx context.Context, groupID string, expand string, recurse *bool, cacheControl string) (result Group, err error) {
	req, err := client.GetPreparer(ctx, groupID, expand, recurse, cacheControl)
	if err != nil {
		err = autorest.NewErrorWithError(err, "management.GroupsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "management.GroupsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "management.GroupsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client GroupsClient) GetPreparer(ctx context.Context, groupID string, expand string, recurse *bool, cacheControl string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupId": autorest.Encode("path", groupID),
	}

	const APIVersion = "2017-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(string(expand)) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if recurse != nil {
		queryParameters["$recurse"] = autorest.Encode("query", *recurse)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementGroups/{groupId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(cacheControl) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Cache-Control", autorest.String(cacheControl)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client GroupsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client GroupsClient) GetResponder(resp *http.Response) (result Group, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list management groups for the authenticated user.
//
// cacheControl is indicates that the request shouldn't utilize any caches. skiptoken is page continuation token is
// only used if a previous operation returned a partial result.
// If a previous response contains a nextLink element, the value of the nextLink element will include a token parameter
// that specifies a starting point to use for subsequent calls.
func (client GroupsClient) List(ctx context.Context, cacheControl string, skiptoken string) (result GroupListResultPage, err error) {
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, cacheControl, skiptoken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "management.GroupsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.glr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "management.GroupsClient", "List", resp, "Failure sending request")
		return
	}

	result.glr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "management.GroupsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client GroupsClient) ListPreparer(ctx context.Context, cacheControl string, skiptoken string) (*http.Request, error) {
	const APIVersion = "2017-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", skiptoken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Management/managementGroups"),
		autorest.WithQueryParameters(queryParameters))
	if len(cacheControl) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Cache-Control", autorest.String(cacheControl)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client GroupsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client GroupsClient) ListResponder(resp *http.Response) (result GroupListResult, err error) {
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
func (client GroupsClient) listNextResults(lastResults GroupListResult) (result GroupListResult, err error) {
	req, err := lastResults.groupListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "management.GroupsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "management.GroupsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "management.GroupsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client GroupsClient) ListComplete(ctx context.Context, cacheControl string, skiptoken string) (result GroupListResultIterator, err error) {
	result.page, err = client.List(ctx, cacheControl, skiptoken)
	return
}

// Update update a management group.
//
// groupID is management Group ID. createGroupRequest is management group creation parameters. cacheControl is
// indicates that the request shouldn't utilize any caches.
func (client GroupsClient) Update(ctx context.Context, groupID string, createGroupRequest CreateGroupRequest, cacheControl string) (result Group, err error) {
	req, err := client.UpdatePreparer(ctx, groupID, createGroupRequest, cacheControl)
	if err != nil {
		err = autorest.NewErrorWithError(err, "management.GroupsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "management.GroupsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "management.GroupsClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client GroupsClient) UpdatePreparer(ctx context.Context, groupID string, createGroupRequest CreateGroupRequest, cacheControl string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupId": autorest.Encode("path", groupID),
	}

	const APIVersion = "2017-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementGroups/{groupId}", pathParameters),
		autorest.WithJSON(createGroupRequest),
		autorest.WithQueryParameters(queryParameters))
	if len(cacheControl) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Cache-Control", autorest.String(cacheControl)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client GroupsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client GroupsClient) UpdateResponder(resp *http.Response) (result Group, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
