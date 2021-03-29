package billing

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
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

// AvailableBalancesClient is the billing client provides access to billing resources for Azure subscriptions.
type AvailableBalancesClient struct {
	BaseClient
}

// NewAvailableBalancesClient creates an instance of the AvailableBalancesClient client.
func NewAvailableBalancesClient(subscriptionID string) AvailableBalancesClient {
	return NewAvailableBalancesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAvailableBalancesClientWithBaseURI creates an instance of the AvailableBalancesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewAvailableBalancesClientWithBaseURI(baseURI string, subscriptionID string) AvailableBalancesClient {
	return AvailableBalancesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetByBillingProfile the latest available credit balance for a given billingAccountName and billingProfileName.
// Parameters:
// billingAccountName - billing Account Id.
// billingProfileName - billing Profile Id.
func (client AvailableBalancesClient) GetByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string) (result AvailableBalance, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AvailableBalancesClient.GetByBillingProfile")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetByBillingProfilePreparer(ctx, billingAccountName, billingProfileName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.AvailableBalancesClient", "GetByBillingProfile", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByBillingProfileSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.AvailableBalancesClient", "GetByBillingProfile", resp, "Failure sending request")
		return
	}

	result, err = client.GetByBillingProfileResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.AvailableBalancesClient", "GetByBillingProfile", resp, "Failure responding to request")
		return
	}

	return
}

// GetByBillingProfilePreparer prepares the GetByBillingProfile request.
func (client AvailableBalancesClient) GetByBillingProfilePreparer(ctx context.Context, billingAccountName string, billingProfileName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"billingProfileName": autorest.Encode("path", billingProfileName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/availableBalance/default", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByBillingProfileSender sends the GetByBillingProfile request. The method will close the
// http.Response Body if it receives an error.
func (client AvailableBalancesClient) GetByBillingProfileSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetByBillingProfileResponder handles the response to the GetByBillingProfile request. The method always
// closes the http.Response Body.
func (client AvailableBalancesClient) GetByBillingProfileResponder(resp *http.Response) (result AvailableBalance, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
