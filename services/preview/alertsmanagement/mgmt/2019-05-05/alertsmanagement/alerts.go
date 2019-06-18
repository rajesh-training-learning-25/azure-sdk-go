package alertsmanagement

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

// AlertsClient is the alertsManagement Client
type AlertsClient struct {
	BaseClient
}

// NewAlertsClient creates an instance of the AlertsClient client.
func NewAlertsClient(subscriptionID string) AlertsClient {
	return NewAlertsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAlertsClientWithBaseURI creates an instance of the AlertsClient client.
func NewAlertsClientWithBaseURI(baseURI string, subscriptionID string) AlertsClient {
	return AlertsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ChangeState change the state of the alert.
// Parameters:
// alertID - unique ID of an alert object.
// newState - filter by state
func (client AlertsClient) ChangeState(ctx context.Context, alertID string, newState AlertState) (result Alert, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AlertsClient.ChangeState")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ChangeStatePreparer(ctx, alertID, newState)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertsmanagement.AlertsClient", "ChangeState", nil, "Failure preparing request")
		return
	}

	resp, err := client.ChangeStateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "alertsmanagement.AlertsClient", "ChangeState", resp, "Failure sending request")
		return
	}

	result, err = client.ChangeStateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertsmanagement.AlertsClient", "ChangeState", resp, "Failure responding to request")
	}

	return
}

// ChangeStatePreparer prepares the ChangeState request.
func (client AlertsClient) ChangeStatePreparer(ctx context.Context, alertID string, newState AlertState) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"alertId":        autorest.Encode("path", alertID),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-05-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
		"newState":    autorest.Encode("query", newState),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.AlertsManagement/alerts/{alertId}/changestate", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ChangeStateSender sends the ChangeState request. The method will close the
// http.Response Body if it receives an error.
func (client AlertsClient) ChangeStateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ChangeStateResponder handles the response to the ChangeState request. The method always
// closes the http.Response Body.
func (client AlertsClient) ChangeStateResponder(resp *http.Response) (result Alert, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetAll list all the existing alerts, where the results can be selective by passing multiple filter parameters
// including time range and sorted on specific fields.
// Parameters:
// targetResource - filter by target resource
// targetResourceGroup - filter by target resource group name
// targetResourceType - filter by target resource type
// monitorService - filter by monitor service which is the source of the alert object.
// monitorCondition - filter by monitor condition which is the state of the alert at monitor service
// severity - filter by severity
// alertState - filter by state
// smartGroupID - filter by smart Group Id
// includePayload - include payload field content, default value is 'false'.
// pageCount - number of items per page, default value is '25'.
// sortBy - sort the query results by input field, default value is 'lastModifiedDateTime'.
// sortOrder - sort the query results order in either ascending or descending, default value is 'desc' for time
// fields and 'asc' for others.
// timeRange - filter by time range, default value is 1 day
func (client AlertsClient) GetAll(ctx context.Context, targetResource string, targetResourceGroup string, targetResourceType string, monitorService MonitorService, monitorCondition MonitorCondition, severity Severity, alertState AlertState, smartGroupID string, includePayload *bool, pageCount *int32, sortBy AlertsSortByFields, sortOrder string, timeRange TimeRange) (result AlertsListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AlertsClient.GetAll")
		defer func() {
			sc := -1
			if result.al.Response.Response != nil {
				sc = result.al.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.getAllNextResults
	req, err := client.GetAllPreparer(ctx, targetResource, targetResourceGroup, targetResourceType, monitorService, monitorCondition, severity, alertState, smartGroupID, includePayload, pageCount, sortBy, sortOrder, timeRange)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertsmanagement.AlertsClient", "GetAll", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAllSender(req)
	if err != nil {
		result.al.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "alertsmanagement.AlertsClient", "GetAll", resp, "Failure sending request")
		return
	}

	result.al, err = client.GetAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertsmanagement.AlertsClient", "GetAll", resp, "Failure responding to request")
	}

	return
}

// GetAllPreparer prepares the GetAll request.
func (client AlertsClient) GetAllPreparer(ctx context.Context, targetResource string, targetResourceGroup string, targetResourceType string, monitorService MonitorService, monitorCondition MonitorCondition, severity Severity, alertState AlertState, smartGroupID string, includePayload *bool, pageCount *int32, sortBy AlertsSortByFields, sortOrder string, timeRange TimeRange) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-05-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(targetResource) > 0 {
		queryParameters["targetResource"] = autorest.Encode("query", targetResource)
	}
	if len(targetResourceGroup) > 0 {
		queryParameters["targetResourceGroup"] = autorest.Encode("query", targetResourceGroup)
	}
	if len(targetResourceType) > 0 {
		queryParameters["targetResourceType"] = autorest.Encode("query", targetResourceType)
	}
	if len(string(monitorService)) > 0 {
		queryParameters["monitorService"] = autorest.Encode("query", monitorService)
	}
	if len(string(monitorCondition)) > 0 {
		queryParameters["monitorCondition"] = autorest.Encode("query", monitorCondition)
	}
	if len(string(severity)) > 0 {
		queryParameters["severity"] = autorest.Encode("query", severity)
	}
	if len(string(alertState)) > 0 {
		queryParameters["alertState"] = autorest.Encode("query", alertState)
	}
	if len(smartGroupID) > 0 {
		queryParameters["smartGroupId"] = autorest.Encode("query", smartGroupID)
	}
	if includePayload != nil {
		queryParameters["includePayload"] = autorest.Encode("query", *includePayload)
	}
	if pageCount != nil {
		queryParameters["pageCount"] = autorest.Encode("query", *pageCount)
	}
	if len(string(sortBy)) > 0 {
		queryParameters["sortBy"] = autorest.Encode("query", sortBy)
	}
	if len(string(sortOrder)) > 0 {
		queryParameters["sortOrder"] = autorest.Encode("query", sortOrder)
	}
	if len(string(timeRange)) > 0 {
		queryParameters["timeRange"] = autorest.Encode("query", timeRange)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.AlertsManagement/alerts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetAllSender sends the GetAll request. The method will close the
// http.Response Body if it receives an error.
func (client AlertsClient) GetAllSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetAllResponder handles the response to the GetAll request. The method always
// closes the http.Response Body.
func (client AlertsClient) GetAllResponder(resp *http.Response) (result AlertsList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getAllNextResults retrieves the next set of results, if any.
func (client AlertsClient) getAllNextResults(ctx context.Context, lastResults AlertsList) (result AlertsList, err error) {
	req, err := lastResults.alertsListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "alertsmanagement.AlertsClient", "getAllNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "alertsmanagement.AlertsClient", "getAllNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertsmanagement.AlertsClient", "getAllNextResults", resp, "Failure responding to next results request")
	}
	return
}

// GetAllComplete enumerates all values, automatically crossing page boundaries as required.
func (client AlertsClient) GetAllComplete(ctx context.Context, targetResource string, targetResourceGroup string, targetResourceType string, monitorService MonitorService, monitorCondition MonitorCondition, severity Severity, alertState AlertState, smartGroupID string, includePayload *bool, pageCount *int32, sortBy AlertsSortByFields, sortOrder string, timeRange TimeRange) (result AlertsListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AlertsClient.GetAll")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.GetAll(ctx, targetResource, targetResourceGroup, targetResourceType, monitorService, monitorCondition, severity, alertState, smartGroupID, includePayload, pageCount, sortBy, sortOrder, timeRange)
	return
}

// GetByID get information related to a specific alert
// Parameters:
// alertID - unique ID of an alert object.
func (client AlertsClient) GetByID(ctx context.Context, alertID string) (result Alert, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AlertsClient.GetByID")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetByIDPreparer(ctx, alertID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertsmanagement.AlertsClient", "GetByID", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "alertsmanagement.AlertsClient", "GetByID", resp, "Failure sending request")
		return
	}

	result, err = client.GetByIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertsmanagement.AlertsClient", "GetByID", resp, "Failure responding to request")
	}

	return
}

// GetByIDPreparer prepares the GetByID request.
func (client AlertsClient) GetByIDPreparer(ctx context.Context, alertID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"alertId":        autorest.Encode("path", alertID),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-05-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.AlertsManagement/alerts/{alertId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByIDSender sends the GetByID request. The method will close the
// http.Response Body if it receives an error.
func (client AlertsClient) GetByIDSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetByIDResponder handles the response to the GetByID request. The method always
// closes the http.Response Body.
func (client AlertsClient) GetByIDResponder(resp *http.Response) (result Alert, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetHistory get the history of the changes of an alert.
// Parameters:
// alertID - unique ID of an alert object.
func (client AlertsClient) GetHistory(ctx context.Context, alertID string) (result AlertModification, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AlertsClient.GetHistory")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetHistoryPreparer(ctx, alertID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertsmanagement.AlertsClient", "GetHistory", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetHistorySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "alertsmanagement.AlertsClient", "GetHistory", resp, "Failure sending request")
		return
	}

	result, err = client.GetHistoryResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertsmanagement.AlertsClient", "GetHistory", resp, "Failure responding to request")
	}

	return
}

// GetHistoryPreparer prepares the GetHistory request.
func (client AlertsClient) GetHistoryPreparer(ctx context.Context, alertID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"alertId":        autorest.Encode("path", alertID),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-05-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.AlertsManagement/alerts/{alertId}/history", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetHistorySender sends the GetHistory request. The method will close the
// http.Response Body if it receives an error.
func (client AlertsClient) GetHistorySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetHistoryResponder handles the response to the GetHistory request. The method always
// closes the http.Response Body.
func (client AlertsClient) GetHistoryResponder(resp *http.Response) (result AlertModification, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetSummary summary of alerts with the count each severity.
// Parameters:
// targetResourceGroup - filter by target resource group name
// timeRange - filter by time range, default value is 1 day
func (client AlertsClient) GetSummary(ctx context.Context, targetResourceGroup string, timeRange TimeRange) (result AlertsSummary, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AlertsClient.GetSummary")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetSummaryPreparer(ctx, targetResourceGroup, timeRange)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertsmanagement.AlertsClient", "GetSummary", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSummarySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "alertsmanagement.AlertsClient", "GetSummary", resp, "Failure sending request")
		return
	}

	result, err = client.GetSummaryResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertsmanagement.AlertsClient", "GetSummary", resp, "Failure responding to request")
	}

	return
}

// GetSummaryPreparer prepares the GetSummary request.
func (client AlertsClient) GetSummaryPreparer(ctx context.Context, targetResourceGroup string, timeRange TimeRange) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-05-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(targetResourceGroup) > 0 {
		queryParameters["targetResourceGroup"] = autorest.Encode("query", targetResourceGroup)
	}
	if len(string(timeRange)) > 0 {
		queryParameters["timeRange"] = autorest.Encode("query", timeRange)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.AlertsManagement/alertsSummary", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSummarySender sends the GetSummary request. The method will close the
// http.Response Body if it receives an error.
func (client AlertsClient) GetSummarySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetSummaryResponder handles the response to the GetSummary request. The method always
// closes the http.Response Body.
func (client AlertsClient) GetSummaryResponder(resp *http.Response) (result AlertsSummary, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
