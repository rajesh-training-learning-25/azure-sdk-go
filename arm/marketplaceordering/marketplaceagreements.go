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
// Code generated by Microsoft (R) AutoRest Code Generator 2.2.21.0
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// MarketplaceAgreementsClient is the REST API for MarketplaceOrdering Agreements.
type MarketplaceAgreementsClient struct {
	ManagementClient
}

// NewMarketplaceAgreementsClient creates an instance of the MarketplaceAgreementsClient client.
func NewMarketplaceAgreementsClient(subscriptionID string) MarketplaceAgreementsClient {
	return NewMarketplaceAgreementsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewMarketplaceAgreementsClientWithBaseURI creates an instance of the MarketplaceAgreementsClient client.
func NewMarketplaceAgreementsClientWithBaseURI(baseURI string, subscriptionID string) MarketplaceAgreementsClient {
	return MarketplaceAgreementsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create save marketplace terms.
//
// offerType is offer Type, currently only virtualmachine type is supported. publisherID is publisher identifier string
// of image being deployed. offerID is offer identifier string of image being deployed. planID is plan identifier
// string of image being deployed. parameters is parameters supplied to the Create Marketplace Terms operation.
func (client MarketplaceAgreementsClient) Create(offerType string, publisherID string, offerID string, planID string, parameters AgreementTerms) (result AgreementTerms, err error) {
	if err := validation.Validate([]validation.Validation{}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "marketplaceordering.MarketplaceAgreementsClient", "Create")
	}

	req, err := client.CreatePreparer(offerType, publisherID, offerID, planID, parameters)
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
func (client MarketplaceAgreementsClient) CreatePreparer(offerType string, publisherID string, offerID string, planID string, parameters AgreementTerms) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"offerId":        autorest.Encode("path", offerID),
		"offerType":      autorest.Encode("path", offerType),
		"planId":         autorest.Encode("path", planID),
		"publisherId":    autorest.Encode("path", publisherID),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.MarketplaceOrdering/offerTypes/{offerType}/publishers/{publisherId}/offers/{offerId}/plans/{planId}/agreements/current", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client MarketplaceAgreementsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
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
//
// offerType is offer Type, currently only virtualmachine type is supported. publisherID is publisher identifier string
// of image being deployed. offerID is offer identifier string of image being deployed. planID is plan identifier
// string of image being deployed.
func (client MarketplaceAgreementsClient) Get(offerType string, publisherID string, offerID string, planID string) (result AgreementTerms, err error) {
	if err := validation.Validate([]validation.Validation{}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "marketplaceordering.MarketplaceAgreementsClient", "Get")
	}

	req, err := client.GetPreparer(offerType, publisherID, offerID, planID)
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
func (client MarketplaceAgreementsClient) GetPreparer(offerType string, publisherID string, offerID string, planID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"offerId":        autorest.Encode("path", offerID),
		"offerType":      autorest.Encode("path", offerType),
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
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client MarketplaceAgreementsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
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
