// Package mixedreality implements the Azure ARM Mixedreality service API version 2021-01-01.
//
// Mixed Reality Client
package mixedreality

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

const (
	// DefaultBaseURI is the default URI used for the service Mixedreality
	DefaultBaseURI = "https://management.azure.com"
)

// BaseClient is the base client for Mixedreality.
type BaseClient struct {
	autorest.Client
	BaseURI        string
	SubscriptionID string
}

// New creates an instance of the BaseClient client.
func New(subscriptionID string) BaseClient {
	return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWithBaseURI creates an instance of the BaseClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return BaseClient{
		Client:         autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:        baseURI,
		SubscriptionID: subscriptionID,
	}
}

// CheckNameAvailabilityLocal check Name Availability for local uniqueness
// Parameters:
// location - the location in which uniqueness will be verified.
// checkNameAvailability - check Name Availability Request.
func (client BaseClient) CheckNameAvailabilityLocal(ctx context.Context, location string, checkNameAvailability CheckNameAvailabilityRequest) (result CheckNameAvailabilityResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.CheckNameAvailabilityLocal")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: location,
			Constraints: []validation.Constraint{{Target: "location", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "location", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "location", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: checkNameAvailability,
			Constraints: []validation.Constraint{{Target: "checkNameAvailability.Name", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "checkNameAvailability.Type", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("mixedreality.BaseClient", "CheckNameAvailabilityLocal", err.Error())
	}

	req, err := client.CheckNameAvailabilityLocalPreparer(ctx, location, checkNameAvailability)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mixedreality.BaseClient", "CheckNameAvailabilityLocal", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckNameAvailabilityLocalSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "mixedreality.BaseClient", "CheckNameAvailabilityLocal", resp, "Failure sending request")
		return
	}

	result, err = client.CheckNameAvailabilityLocalResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mixedreality.BaseClient", "CheckNameAvailabilityLocal", resp, "Failure responding to request")
		return
	}

	return
}

// CheckNameAvailabilityLocalPreparer prepares the CheckNameAvailabilityLocal request.
func (client BaseClient) CheckNameAvailabilityLocalPreparer(ctx context.Context, location string, checkNameAvailability CheckNameAvailabilityRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.MixedReality/locations/{location}/checkNameAvailability", pathParameters),
		autorest.WithJSON(checkNameAvailability),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckNameAvailabilityLocalSender sends the CheckNameAvailabilityLocal request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) CheckNameAvailabilityLocalSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CheckNameAvailabilityLocalResponder handles the response to the CheckNameAvailabilityLocal request. The method always
// closes the http.Response Body.
func (client BaseClient) CheckNameAvailabilityLocalResponder(resp *http.Response) (result CheckNameAvailabilityResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
