package billing

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

// ProductsClient is the billing client provides access to billing resources for Azure subscriptions.
type ProductsClient struct {
	BaseClient
}

// NewProductsClient creates an instance of the ProductsClient client.
func NewProductsClient(subscriptionID string) ProductsClient {
	return NewProductsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProductsClientWithBaseURI creates an instance of the ProductsClient client.
func NewProductsClientWithBaseURI(baseURI string, subscriptionID string) ProductsClient {
	return ProductsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get a single product by name.
// Parameters:
// billingAccountName - billing Account Id.
// invoiceSectionName - invoiceSection Id.
// productName - invoide Id.
func (client ProductsClient) Get(ctx context.Context, billingAccountName string, invoiceSectionName string, productName string) (result ProductSummary, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, billingAccountName, invoiceSectionName, productName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ProductsClient) GetPreparer(ctx context.Context, billingAccountName string, invoiceSectionName string, productName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"invoiceSectionName": autorest.Encode("path", invoiceSectionName),
		"productName":        autorest.Encode("path", productName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/invoiceSections/{invoiceSectionName}/products/{productName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ProductsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ProductsClient) GetResponder(resp *http.Response) (result ProductSummary, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Transfer the operation to transfer a Product to another InvoiceSection.
// Parameters:
// billingAccountName - billing Account Id.
// invoiceSectionName - invoiceSection Id.
// productName - invoide Id.
// parameters - parameters supplied to the Transfer Product operation.
func (client ProductsClient) Transfer(ctx context.Context, billingAccountName string, invoiceSectionName string, productName string, parameters TransferProductRequestProperties) (result ProductSummary, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductsClient.Transfer")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.TransferPreparer(ctx, billingAccountName, invoiceSectionName, productName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "Transfer", nil, "Failure preparing request")
		return
	}

	resp, err := client.TransferSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "Transfer", resp, "Failure sending request")
		return
	}

	result, err = client.TransferResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "Transfer", resp, "Failure responding to request")
	}

	return
}

// TransferPreparer prepares the Transfer request.
func (client ProductsClient) TransferPreparer(ctx context.Context, billingAccountName string, invoiceSectionName string, productName string, parameters TransferProductRequestProperties) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"invoiceSectionName": autorest.Encode("path", invoiceSectionName),
		"productName":        autorest.Encode("path", productName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/invoiceSections/{invoiceSectionName}/products/{productName}/transfer", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// TransferSender sends the Transfer request. The method will close the
// http.Response Body if it receives an error.
func (client ProductsClient) TransferSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// TransferResponder handles the response to the Transfer request. The method always
// closes the http.Response Body.
func (client ProductsClient) TransferResponder(resp *http.Response) (result ProductSummary, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
