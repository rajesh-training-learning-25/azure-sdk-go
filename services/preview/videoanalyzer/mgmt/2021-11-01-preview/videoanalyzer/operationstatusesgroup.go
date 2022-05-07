// Deprecated: We’re retiring the Azure Video Analyzer preview service, you're advised to transition your applications off of Video Analyzer by 01 December 2022. This SDK is no longer maintained and won’t work after the service is retired

package videoanalyzer

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

// OperationStatusesGroupClient is the azure Video Analyzer provides a platform for you to build intelligent video
// applications that span the edge and the cloud
type OperationStatusesGroupClient struct {
	BaseClient
}

// NewOperationStatusesGroupClient creates an instance of the OperationStatusesGroupClient client.
func NewOperationStatusesGroupClient(subscriptionID string) OperationStatusesGroupClient {
	return NewOperationStatusesGroupClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewOperationStatusesGroupClientWithBaseURI creates an instance of the OperationStatusesGroupClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewOperationStatusesGroupClientWithBaseURI(baseURI string, subscriptionID string) OperationStatusesGroupClient {
	return OperationStatusesGroupClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get video analyzer operation status.
// Parameters:
// locationName - location name.
// operationID - operation Id.
func (client OperationStatusesGroupClient) Get(ctx context.Context, locationName string, operationID string) (result OperationStatus, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationStatusesGroupClient.Get")
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
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("videoanalyzer.OperationStatusesGroupClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, locationName, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "videoanalyzer.OperationStatusesGroupClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "videoanalyzer.OperationStatusesGroupClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "videoanalyzer.OperationStatusesGroupClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client OperationStatusesGroupClient) GetPreparer(ctx context.Context, locationName string, operationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"locationName":   autorest.Encode("path", locationName),
		"operationId":    autorest.Encode("path", operationID),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Media/locations/{locationName}/videoAnalyzerOperationStatuses/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client OperationStatusesGroupClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client OperationStatusesGroupClient) GetResponder(resp *http.Response) (result OperationStatus, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
