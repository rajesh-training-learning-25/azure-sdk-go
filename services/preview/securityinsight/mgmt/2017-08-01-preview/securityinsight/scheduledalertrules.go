package securityinsight

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

// ScheduledAlertRulesClient is the API spec for Microsoft.SecurityInsights (Azure Security Insights) resource provider
type ScheduledAlertRulesClient struct {
	BaseClient
}

// NewScheduledAlertRulesClient creates an instance of the ScheduledAlertRulesClient client.
func NewScheduledAlertRulesClient(subscriptionID string) ScheduledAlertRulesClient {
	return NewScheduledAlertRulesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewScheduledAlertRulesClientWithBaseURI creates an instance of the ScheduledAlertRulesClient client.
func NewScheduledAlertRulesClientWithBaseURI(baseURI string, subscriptionID string) ScheduledAlertRulesClient {
	return ScheduledAlertRulesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates or updates the alert rule.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// operationalInsightsResourceProvider - the namespace of workspaces resource provider-
// Microsoft.OperationalInsights.
// workspaceName - the name of the workspace.
// ruleID - alert rule ID
// alertScheduledtRule - the alert rule
func (client ScheduledAlertRulesClient) Create(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ruleID string, alertScheduledtRule ScheduledAlertRule) (result ScheduledAlertRule, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ScheduledAlertRulesClient.Create")
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
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: ruleID,
			Constraints: []validation.Constraint{{Target: "ruleID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: alertScheduledtRule,
			Constraints: []validation.Constraint{{Target: "alertScheduledtRule.ScheduledAlertRuleProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "alertScheduledtRule.ScheduledAlertRuleProperties.RuleName", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "alertScheduledtRule.ScheduledAlertRuleProperties.Description", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "alertScheduledtRule.ScheduledAlertRuleProperties.Enabled", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "alertScheduledtRule.ScheduledAlertRuleProperties.Query", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "alertScheduledtRule.ScheduledAlertRuleProperties.QueryFrequency", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "alertScheduledtRule.ScheduledAlertRuleProperties.QueryFrequency", Name: validation.Pattern, Rule: `^(\d+\.)?\d{2}\:\d{2}\:\d{2}$`, Chain: nil}}},
					{Target: "alertScheduledtRule.ScheduledAlertRuleProperties.QueryPeriod", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "alertScheduledtRule.ScheduledAlertRuleProperties.QueryPeriod", Name: validation.Pattern, Rule: `^(\d+\.)?\d{2}\:\d{2}\:\d{2}$`, Chain: nil}}},
					{Target: "alertScheduledtRule.ScheduledAlertRuleProperties.AlertTriggerThreshold", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "alertScheduledtRule.ScheduledAlertRuleProperties.SuppressionEnabled", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "alertScheduledtRule.ScheduledAlertRuleProperties.SuppressionDuration", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "alertScheduledtRule.ScheduledAlertRuleProperties.SuppressionDuration", Name: validation.Pattern, Rule: `^(\d+\.)?\d{2}\:\d{2}\:\d{2}$`, Chain: nil}}},
				}}}}}); err != nil {
		return result, validation.NewError("securityinsight.ScheduledAlertRulesClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName, ruleID, alertScheduledtRule)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.ScheduledAlertRulesClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.ScheduledAlertRulesClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.ScheduledAlertRulesClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ScheduledAlertRulesClient) CreatePreparer(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ruleID string, alertScheduledtRule ScheduledAlertRule) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"operationalInsightsResourceProvider": autorest.Encode("path", operationalInsightsResourceProvider),
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"ruleId":                              autorest.Encode("path", ruleID),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
		"workspaceName":                       autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{operationalInsightsResourceProvider}/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/scheduledAlertRules/{ruleId}", pathParameters),
		autorest.WithJSON(alertScheduledtRule),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ScheduledAlertRulesClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ScheduledAlertRulesClient) CreateResponder(resp *http.Response) (result ScheduledAlertRule, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete the alert rule.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// operationalInsightsResourceProvider - the namespace of workspaces resource provider-
// Microsoft.OperationalInsights.
// workspaceName - the name of the workspace.
// ruleID - alert rule ID
func (client ScheduledAlertRulesClient) Delete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ruleID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ScheduledAlertRulesClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: ruleID,
			Constraints: []validation.Constraint{{Target: "ruleID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.ScheduledAlertRulesClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName, ruleID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.ScheduledAlertRulesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "securityinsight.ScheduledAlertRulesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.ScheduledAlertRulesClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ScheduledAlertRulesClient) DeletePreparer(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ruleID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"operationalInsightsResourceProvider": autorest.Encode("path", operationalInsightsResourceProvider),
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"ruleId":                              autorest.Encode("path", ruleID),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
		"workspaceName":                       autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{operationalInsightsResourceProvider}/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/scheduledAlertRules/{ruleId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ScheduledAlertRulesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ScheduledAlertRulesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a scheduled alert rule.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// operationalInsightsResourceProvider - the namespace of workspaces resource provider-
// Microsoft.OperationalInsights.
// workspaceName - the name of the workspace.
// ruleID - alert rule ID
func (client ScheduledAlertRulesClient) Get(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ruleID string) (result ScheduledAlertRule, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ScheduledAlertRulesClient.Get")
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
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: ruleID,
			Constraints: []validation.Constraint{{Target: "ruleID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.ScheduledAlertRulesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName, ruleID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.ScheduledAlertRulesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.ScheduledAlertRulesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.ScheduledAlertRulesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ScheduledAlertRulesClient) GetPreparer(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ruleID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"operationalInsightsResourceProvider": autorest.Encode("path", operationalInsightsResourceProvider),
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"ruleId":                              autorest.Encode("path", ruleID),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
		"workspaceName":                       autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{operationalInsightsResourceProvider}/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/scheduledAlertRules/{ruleId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ScheduledAlertRulesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ScheduledAlertRulesClient) GetResponder(resp *http.Response) (result ScheduledAlertRule, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all scheduled alert rules.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// operationalInsightsResourceProvider - the namespace of workspaces resource provider-
// Microsoft.OperationalInsights.
// workspaceName - the name of the workspace.
func (client ScheduledAlertRulesClient) List(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string) (result ScheduledAlertRulesListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ScheduledAlertRulesClient.List")
		defer func() {
			sc := -1
			if result.sarl.Response.Response != nil {
				sc = result.sarl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.ScheduledAlertRulesClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.ScheduledAlertRulesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.sarl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.ScheduledAlertRulesClient", "List", resp, "Failure sending request")
		return
	}

	result.sarl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.ScheduledAlertRulesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ScheduledAlertRulesClient) ListPreparer(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"operationalInsightsResourceProvider": autorest.Encode("path", operationalInsightsResourceProvider),
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
		"workspaceName":                       autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{operationalInsightsResourceProvider}/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/scheduledAlertRules", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ScheduledAlertRulesClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ScheduledAlertRulesClient) ListResponder(resp *http.Response) (result ScheduledAlertRulesList, err error) {
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
func (client ScheduledAlertRulesClient) listNextResults(ctx context.Context, lastResults ScheduledAlertRulesList) (result ScheduledAlertRulesList, err error) {
	req, err := lastResults.scheduledAlertRulesListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "securityinsight.ScheduledAlertRulesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "securityinsight.ScheduledAlertRulesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.ScheduledAlertRulesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ScheduledAlertRulesClient) ListComplete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string) (result ScheduledAlertRulesListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ScheduledAlertRulesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName)
	return
}
