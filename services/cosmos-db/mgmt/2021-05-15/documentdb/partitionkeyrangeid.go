package documentdb

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

// PartitionKeyRangeIDClient is the client for the PartitionKeyRangeID methods of the Documentdb service.
type PartitionKeyRangeIDClient struct {
	BaseClient
}

// NewPartitionKeyRangeIDClient creates an instance of the PartitionKeyRangeIDClient client.
func NewPartitionKeyRangeIDClient(subscriptionID string) PartitionKeyRangeIDClient {
	return NewPartitionKeyRangeIDClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPartitionKeyRangeIDClientWithBaseURI creates an instance of the PartitionKeyRangeIDClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewPartitionKeyRangeIDClientWithBaseURI(baseURI string, subscriptionID string) PartitionKeyRangeIDClient {
	return PartitionKeyRangeIDClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListMetrics retrieves the metrics determined by the given filter for the given partition key range id.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// accountName - cosmos DB database account name.
// databaseRid - cosmos DB database rid.
// collectionRid - cosmos DB collection rid.
// partitionKeyRangeID - partition Key Range Id for which to get data.
// filter - an OData filter expression that describes a subset of metrics to return. The parameters that can be
// filtered are name.value (name of the metric, can have an or of multiple names), startTime, endTime, and
// timeGrain. The supported operator is eq.
func (client PartitionKeyRangeIDClient) ListMetrics(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, partitionKeyRangeID string, filter string) (result PartitionMetricListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PartitionKeyRangeIDClient.ListMetrics")
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
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[a-z0-9]+(-[a-z0-9]+)*`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("documentdb.PartitionKeyRangeIDClient", "ListMetrics", err.Error())
	}

	req, err := client.ListMetricsPreparer(ctx, resourceGroupName, accountName, databaseRid, collectionRid, partitionKeyRangeID, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.PartitionKeyRangeIDClient", "ListMetrics", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListMetricsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "documentdb.PartitionKeyRangeIDClient", "ListMetrics", resp, "Failure sending request")
		return
	}

	result, err = client.ListMetricsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.PartitionKeyRangeIDClient", "ListMetrics", resp, "Failure responding to request")
		return
	}

	return
}

// ListMetricsPreparer prepares the ListMetrics request.
func (client PartitionKeyRangeIDClient) ListMetricsPreparer(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, partitionKeyRangeID string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":         autorest.Encode("path", accountName),
		"collectionRid":       autorest.Encode("path", collectionRid),
		"databaseRid":         autorest.Encode("path", databaseRid),
		"partitionKeyRangeId": autorest.Encode("path", partitionKeyRangeID),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-05-15"
	queryParameters := map[string]interface{}{
		"$filter":     autorest.Encode("query", filter),
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/collections/{collectionRid}/partitionKeyRangeId/{partitionKeyRangeId}/metrics", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListMetricsSender sends the ListMetrics request. The method will close the
// http.Response Body if it receives an error.
func (client PartitionKeyRangeIDClient) ListMetricsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListMetricsResponder handles the response to the ListMetrics request. The method always
// closes the http.Response Body.
func (client PartitionKeyRangeIDClient) ListMetricsResponder(resp *http.Response) (result PartitionMetricListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
