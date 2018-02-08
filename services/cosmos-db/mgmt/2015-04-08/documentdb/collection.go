package documentdb

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
	"net/http"
)

// CollectionClient is the azure Cosmos DB Database Service Resource Provider REST API
type CollectionClient struct {
	BaseClient
}

// NewCollectionClient creates an instance of the CollectionClient client.
func NewCollectionClient(subscriptionID string) CollectionClient {
	return NewCollectionClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCollectionClientWithBaseURI creates an instance of the CollectionClient client.
func NewCollectionClientWithBaseURI(baseURI string, subscriptionID string) CollectionClient {
	return CollectionClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListMetricDefinitions retrieves metric defintions for the given collection.
//
// resourceGroupName is name of an Azure resource group. accountName is cosmos DB database account name. databaseRid is
// cosmos DB database rid. collectionRid is cosmos DB collection rid.
func (client CollectionClient) ListMetricDefinitions(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string) (result MetricDefinitionsListResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "documentdb.CollectionClient", "ListMetricDefinitions")
	}

	req, err := client.ListMetricDefinitionsPreparer(ctx, resourceGroupName, accountName, databaseRid, collectionRid)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.CollectionClient", "ListMetricDefinitions", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListMetricDefinitionsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "documentdb.CollectionClient", "ListMetricDefinitions", resp, "Failure sending request")
		return
	}

	result, err = client.ListMetricDefinitionsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.CollectionClient", "ListMetricDefinitions", resp, "Failure responding to request")
	}

	return
}

// ListMetricDefinitionsPreparer prepares the ListMetricDefinitions request.
func (client CollectionClient) ListMetricDefinitionsPreparer(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"collectionRid":     autorest.Encode("path", collectionRid),
		"databaseRid":       autorest.Encode("path", databaseRid),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-04-08"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/collections/{collectionRid}/metricDefinitions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListMetricDefinitionsSender sends the ListMetricDefinitions request. The method will close the
// http.Response Body if it receives an error.
func (client CollectionClient) ListMetricDefinitionsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListMetricDefinitionsResponder handles the response to the ListMetricDefinitions request. The method always
// closes the http.Response Body.
func (client CollectionClient) ListMetricDefinitionsResponder(resp *http.Response) (result MetricDefinitionsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListMetrics retrieves the metrics determined by the given filter for the given database account and collection.
//
// resourceGroupName is name of an Azure resource group. accountName is cosmos DB database account name. databaseRid is
// cosmos DB database rid. collectionRid is cosmos DB collection rid. filter is an OData filter expression that
// describes a subset of metrics to return. The parameters that can be filtered are name.value (name of the metric, can
// have an or of multiple names), startTime, endTime, and timeGrain. The supported operator is eq.
func (client CollectionClient) ListMetrics(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, filter string) (result MetricListResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "documentdb.CollectionClient", "ListMetrics")
	}

	req, err := client.ListMetricsPreparer(ctx, resourceGroupName, accountName, databaseRid, collectionRid, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.CollectionClient", "ListMetrics", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListMetricsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "documentdb.CollectionClient", "ListMetrics", resp, "Failure sending request")
		return
	}

	result, err = client.ListMetricsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.CollectionClient", "ListMetrics", resp, "Failure responding to request")
	}

	return
}

// ListMetricsPreparer prepares the ListMetrics request.
func (client CollectionClient) ListMetricsPreparer(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"collectionRid":     autorest.Encode("path", collectionRid),
		"databaseRid":       autorest.Encode("path", databaseRid),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-04-08"
	queryParameters := map[string]interface{}{
		"$filter":     autorest.Encode("query", filter),
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/collections/{collectionRid}/metrics", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListMetricsSender sends the ListMetrics request. The method will close the
// http.Response Body if it receives an error.
func (client CollectionClient) ListMetricsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListMetricsResponder handles the response to the ListMetrics request. The method always
// closes the http.Response Body.
func (client CollectionClient) ListMetricsResponder(resp *http.Response) (result MetricListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListUsages retrieves the usages (most recent storage data) for the given collection.
//
// resourceGroupName is name of an Azure resource group. accountName is cosmos DB database account name. databaseRid is
// cosmos DB database rid. collectionRid is cosmos DB collection rid. filter is an OData filter expression that
// describes a subset of usages to return. The supported parameter is name.value (name of the metric, can have an or of
// multiple names).
func (client CollectionClient) ListUsages(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, filter string) (result UsagesResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "documentdb.CollectionClient", "ListUsages")
	}

	req, err := client.ListUsagesPreparer(ctx, resourceGroupName, accountName, databaseRid, collectionRid, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.CollectionClient", "ListUsages", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListUsagesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "documentdb.CollectionClient", "ListUsages", resp, "Failure sending request")
		return
	}

	result, err = client.ListUsagesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.CollectionClient", "ListUsages", resp, "Failure responding to request")
	}

	return
}

// ListUsagesPreparer prepares the ListUsages request.
func (client CollectionClient) ListUsagesPreparer(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"collectionRid":     autorest.Encode("path", collectionRid),
		"databaseRid":       autorest.Encode("path", databaseRid),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-04-08"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/collections/{collectionRid}/usages", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListUsagesSender sends the ListUsages request. The method will close the
// http.Response Body if it receives an error.
func (client CollectionClient) ListUsagesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListUsagesResponder handles the response to the ListUsages request. The method always
// closes the http.Response Body.
func (client CollectionClient) ListUsagesResponder(resp *http.Response) (result UsagesResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
