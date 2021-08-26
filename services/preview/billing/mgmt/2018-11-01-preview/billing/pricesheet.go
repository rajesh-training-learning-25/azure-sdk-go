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

// PriceSheetClient is the billing client provides access to billing resources for Azure subscriptions.
type PriceSheetClient struct {
	BaseClient
}

// NewPriceSheetClient creates an instance of the PriceSheetClient client.
func NewPriceSheetClient(subscriptionID string) PriceSheetClient {
	return NewPriceSheetClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPriceSheetClientWithBaseURI creates an instance of the PriceSheetClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewPriceSheetClientWithBaseURI(baseURI string, subscriptionID string) PriceSheetClient {
	return PriceSheetClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Download download price sheet for an invoice.
// Parameters:
// billingAccountName - azure Billing Account ID.
// invoiceName - the name of an invoice resource.
func (client PriceSheetClient) Download(ctx context.Context, billingAccountName string, invoiceName string) (result PriceSheetDownloadFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PriceSheetClient.Download")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DownloadPreparer(ctx, billingAccountName, invoiceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PriceSheetClient", "Download", nil, "Failure preparing request")
		return
	}

	result, err = client.DownloadSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PriceSheetClient", "Download", nil, "Failure sending request")
		return
	}

	return
}

// DownloadPreparer prepares the Download request.
func (client PriceSheetClient) DownloadPreparer(ctx context.Context, billingAccountName string, invoiceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"invoiceName":        autorest.Encode("path", invoiceName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/invoices/{invoiceName}/pricesheet/default/download", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DownloadSender sends the Download request. The method will close the
// http.Response Body if it receives an error.
func (client PriceSheetClient) DownloadSender(req *http.Request) (future PriceSheetDownloadFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DownloadResponder handles the response to the Download request. The method always
// closes the http.Response Body.
func (client PriceSheetClient) DownloadResponder(resp *http.Response) (result DownloadURL, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
