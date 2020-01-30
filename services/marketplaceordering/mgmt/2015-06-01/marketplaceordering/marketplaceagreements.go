package marketplaceordering

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

// MarketplaceAgreementsClient is the REST API for MarketplaceOrdering Agreements.
type MarketplaceAgreementsClient struct {
	BaseClient
}

// NewMarketplaceAgreementsClient creates an instance of the MarketplaceAgreementsClient client.
func NewMarketplaceAgreementsClient(subscriptionID string) MarketplaceAgreementsClient {
	return NewMarketplaceAgreementsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewMarketplaceAgreementsClientWithBaseURI creates an instance of the MarketplaceAgreementsClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewMarketplaceAgreementsClientWithBaseURI(baseURI string, subscriptionID string) MarketplaceAgreementsClient {
	return MarketplaceAgreementsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Cancel cancel marketplace terms.
// Parameters:
// publisherID - publisher identifier string of image being deployed.
// offerID - offer identifier string of image being deployed.
// planID - plan identifier string of image being deployed.
func (client MarketplaceAgreementsClient) Cancel(ctx context.Context, publisherID string, offerID string, planID string) (result AgreementTerms, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MarketplaceAgreementsClient.Cancel")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CancelPreparer(ctx, publisherID, offerID, planID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "marketplaceordering.MarketplaceAgreementsClient", "Cancel", nil, "Failure preparing request")
		return
	}

	resp, err := client.CancelSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "marketplaceordering.MarketplaceAgreementsClient", "Cancel", resp, "Failure sending request")
		return
	}

	result, err = client.CancelResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "marketplaceordering.MarketplaceAgreementsClient", "Cancel", resp, "Failure responding to request")
	}

	return
}

// CancelPreparer prepares the Cancel request.
func (client MarketplaceAgreementsClient) CancelPreparer(ctx context.Context, publisherID string, offerID string, planID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"offerId":        autorest.Encode("path", offerID),
		"planId":         autorest.Encode("path", planID),
		"publisherId":    autorest.Encode("path", publisherID),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.MarketplaceOrdering/agreements/{publisherId}/offers/{offerId}/plans/{planId}/cancel", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CancelSender sends the Cancel request. The method will close the
// http.Response Body if it receives an error.
func (client MarketplaceAgreementsClient) CancelSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CancelResponder handles the response to the Cancel request. The method always
// closes the http.Response Body.
func (client MarketplaceAgreementsClient) CancelResponder(resp *http.Response) (result AgreementTerms, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Create save marketplace terms.
// Parameters:
// publisherID - publisher identifier string of image being deployed.
// offerID - offer identifier string of image being deployed.
// planID - plan identifier string of image being deployed.
// parameters - parameters supplied to the Create Marketplace Terms operation.
func (client MarketplaceAgreementsClient) Create(ctx context.Context, publisherID string, offerID string, planID string, parameters AgreementTerms) (result AgreementTerms, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MarketplaceAgreementsClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, publisherID, offerID, planID, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "marketplaceordering.MarketplaceAgreementsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "marketplaceordering.MarketplaceAgreementsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "marketplaceordering.MarketplaceAgreementsClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client MarketplaceAgreementsClient) CreatePreparer(ctx context.Context, publisherID string, offerID string, planID string, parameters AgreementTerms) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"offerId":        autorest.Encode("path", offerID),
		"offerType":      autorest.Encode("path", "virtualmachine"),
		"planId":         autorest.Encode("path", planID),
		"publisherId":    autorest.Encode("path", publisherID),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.MarketplaceOrdering/offerTypes/{offerType}/publishers/{publisherId}/offers/{offerId}/plans/{planId}/agreements/current", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client MarketplaceAgreementsClient) CreateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client MarketplaceAgreementsClient) CreateResponder(resp *http.Response) (result AgreementTerms, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get marketplace terms.
// Parameters:
// publisherID - publisher identifier string of image being deployed.
// offerID - offer identifier string of image being deployed.
// planID - plan identifier string of image being deployed.
func (client MarketplaceAgreementsClient) Get(ctx context.Context, publisherID string, offerID string, planID string) (result AgreementTerms, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MarketplaceAgreementsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, publisherID, offerID, planID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "marketplaceordering.MarketplaceAgreementsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "marketplaceordering.MarketplaceAgreementsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "marketplaceordering.MarketplaceAgreementsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client MarketplaceAgreementsClient) GetPreparer(ctx context.Context, publisherID string, offerID string, planID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"offerId":        autorest.Encode("path", offerID),
		"offerType":      autorest.Encode("path", "virtualmachine"),
		"planId":         autorest.Encode("path", planID),
		"publisherId":    autorest.Encode("path", publisherID),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.MarketplaceOrdering/offerTypes/{offerType}/publishers/{publisherId}/offers/{offerId}/plans/{planId}/agreements/current", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client MarketplaceAgreementsClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client MarketplaceAgreementsClient) GetResponder(resp *http.Response) (result AgreementTerms, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetAgreement get marketplace agreement.
// Parameters:
// publisherID - publisher identifier string of image being deployed.
// offerID - offer identifier string of image being deployed.
// planID - plan identifier string of image being deployed.
func (client MarketplaceAgreementsClient) GetAgreement(ctx context.Context, publisherID string, offerID string, planID string) (result AgreementTerms, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MarketplaceAgreementsClient.GetAgreement")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetAgreementPreparer(ctx, publisherID, offerID, planID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "marketplaceordering.MarketplaceAgreementsClient", "GetAgreement", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAgreementSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "marketplaceordering.MarketplaceAgreementsClient", "GetAgreement", resp, "Failure sending request")
		return
	}

	result, err = client.GetAgreementResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "marketplaceordering.MarketplaceAgreementsClient", "GetAgreement", resp, "Failure responding to request")
	}

	return
}

// GetAgreementPreparer prepares the GetAgreement request.
func (client MarketplaceAgreementsClient) GetAgreementPreparer(ctx context.Context, publisherID string, offerID string, planID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"offerId":        autorest.Encode("path", offerID),
		"planId":         autorest.Encode("path", planID),
		"publisherId":    autorest.Encode("path", publisherID),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.MarketplaceOrdering/agreements/{publisherId}/offers/{offerId}/plans/{planId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetAgreementSender sends the GetAgreement request. The method will close the
// http.Response Body if it receives an error.
func (client MarketplaceAgreementsClient) GetAgreementSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetAgreementResponder handles the response to the GetAgreement request. The method always
// closes the http.Response Body.
func (client MarketplaceAgreementsClient) GetAgreementResponder(resp *http.Response) (result AgreementTerms, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list marketplace agreements in the subscription.
func (client MarketplaceAgreementsClient) List(ctx context.Context) (result ListAgreementTerms, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MarketplaceAgreementsClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "marketplaceordering.MarketplaceAgreementsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "marketplaceordering.MarketplaceAgreementsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "marketplaceordering.MarketplaceAgreementsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client MarketplaceAgreementsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.MarketplaceOrdering/agreements", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client MarketplaceAgreementsClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client MarketplaceAgreementsClient) ListResponder(resp *http.Response) (result ListAgreementTerms, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Sign sign marketplace terms.
// Parameters:
// publisherID - publisher identifier string of image being deployed.
// offerID - offer identifier string of image being deployed.
// planID - plan identifier string of image being deployed.
func (client MarketplaceAgreementsClient) Sign(ctx context.Context, publisherID string, offerID string, planID string) (result AgreementTerms, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MarketplaceAgreementsClient.Sign")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.SignPreparer(ctx, publisherID, offerID, planID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "marketplaceordering.MarketplaceAgreementsClient", "Sign", nil, "Failure preparing request")
		return
	}

	resp, err := client.SignSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "marketplaceordering.MarketplaceAgreementsClient", "Sign", resp, "Failure sending request")
		return
	}

	result, err = client.SignResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "marketplaceordering.MarketplaceAgreementsClient", "Sign", resp, "Failure responding to request")
	}

	return
}

// SignPreparer prepares the Sign request.
func (client MarketplaceAgreementsClient) SignPreparer(ctx context.Context, publisherID string, offerID string, planID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"offerId":        autorest.Encode("path", offerID),
		"planId":         autorest.Encode("path", planID),
		"publisherId":    autorest.Encode("path", publisherID),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.MarketplaceOrdering/agreements/{publisherId}/offers/{offerId}/plans/{planId}/sign", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SignSender sends the Sign request. The method will close the
// http.Response Body if it receives an error.
func (client MarketplaceAgreementsClient) SignSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// SignResponder handles the response to the Sign request. The method always
// closes the http.Response Body.
func (client MarketplaceAgreementsClient) SignResponder(resp *http.Response) (result AgreementTerms, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
