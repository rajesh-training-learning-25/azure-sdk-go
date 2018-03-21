package apimanagement

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
	"net/http"
)

// IdentityProviderClient is the client for the IdentityProvider methods of the Apimanagement service.
type IdentityProviderClient struct {
	BaseClient
}

// NewIdentityProviderClient creates an instance of the IdentityProviderClient client.
func NewIdentityProviderClient() IdentityProviderClient {
	return IdentityProviderClient{New()}
}

// CreateOrUpdate creates or Updates the IdentityProvider configuration.
//
// apimBaseURL is the management endpoint of the API Management service, for example
// https://myapimservice.management.azure-api.net. identityProviderName is identity Provider Type identifier.
// parameters is create parameters.
func (client IdentityProviderClient) CreateOrUpdate(ctx context.Context, apimBaseURL string, identityProviderName IdentityProviderType, parameters IdentityProviderContract) (result IdentityProviderContract, err error) {
	req, err := client.CreateOrUpdatePreparer(ctx, apimBaseURL, identityProviderName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.IdentityProviderClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.IdentityProviderClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.IdentityProviderClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client IdentityProviderClient) CreateOrUpdatePreparer(ctx context.Context, apimBaseURL string, identityProviderName IdentityProviderType, parameters IdentityProviderContract) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"apimBaseUrl": apimBaseURL,
	}

	pathParameters := map[string]interface{}{
		"identityProviderName": autorest.Encode("path", identityProviderName),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithCustomBaseURL("{apimBaseUrl}", urlParameters),
		autorest.WithPathParameters("/identityProviders/{identityProviderName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client IdentityProviderClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client IdentityProviderClient) CreateOrUpdateResponder(resp *http.Response) (result IdentityProviderContract, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified identity provider configuration.
//
// apimBaseURL is the management endpoint of the API Management service, for example
// https://myapimservice.management.azure-api.net. identityProviderName is identity Provider Type identifier.
// ifMatch is the entity state (Etag) version of the backend to delete. A value of "*" can be used for If-Match to
// unconditionally apply the operation.
func (client IdentityProviderClient) Delete(ctx context.Context, apimBaseURL string, identityProviderName IdentityProviderType, ifMatch string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(ctx, apimBaseURL, identityProviderName, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.IdentityProviderClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.IdentityProviderClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.IdentityProviderClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client IdentityProviderClient) DeletePreparer(ctx context.Context, apimBaseURL string, identityProviderName IdentityProviderType, ifMatch string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"apimBaseUrl": apimBaseURL,
	}

	pathParameters := map[string]interface{}{
		"identityProviderName": autorest.Encode("path", identityProviderName),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("{apimBaseUrl}", urlParameters),
		autorest.WithPathParameters("/identityProviders/{identityProviderName}", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client IdentityProviderClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client IdentityProviderClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the configuration details of the identity Provider configured in specified service instance.
//
// apimBaseURL is the management endpoint of the API Management service, for example
// https://myapimservice.management.azure-api.net. identityProviderName is identity Provider Type identifier.
func (client IdentityProviderClient) Get(ctx context.Context, apimBaseURL string, identityProviderName IdentityProviderType) (result IdentityProviderContract, err error) {
	req, err := client.GetPreparer(ctx, apimBaseURL, identityProviderName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.IdentityProviderClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.IdentityProviderClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.IdentityProviderClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client IdentityProviderClient) GetPreparer(ctx context.Context, apimBaseURL string, identityProviderName IdentityProviderType) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"apimBaseUrl": apimBaseURL,
	}

	pathParameters := map[string]interface{}{
		"identityProviderName": autorest.Encode("path", identityProviderName),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{apimBaseUrl}", urlParameters),
		autorest.WithPathParameters("/identityProviders/{identityProviderName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client IdentityProviderClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client IdentityProviderClient) GetResponder(resp *http.Response) (result IdentityProviderContract, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists a collection of Identity Provider configured in the specified service instance.
//
// apimBaseURL is the management endpoint of the API Management service, for example
// https://myapimservice.management.azure-api.net.
func (client IdentityProviderClient) List(ctx context.Context, apimBaseURL string) (result IdentityProviderList, err error) {
	req, err := client.ListPreparer(ctx, apimBaseURL)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.IdentityProviderClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.IdentityProviderClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.IdentityProviderClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client IdentityProviderClient) ListPreparer(ctx context.Context, apimBaseURL string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"apimBaseUrl": apimBaseURL,
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{apimBaseUrl}", urlParameters),
		autorest.WithPath("/identityProviders"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client IdentityProviderClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client IdentityProviderClient) ListResponder(resp *http.Response) (result IdentityProviderList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates an existing IdentityProvider configuration.
//
// apimBaseURL is the management endpoint of the API Management service, for example
// https://myapimservice.management.azure-api.net. identityProviderName is identity Provider Type identifier.
// parameters is update parameters. ifMatch is the entity state (Etag) version of the identity provider
// configuration to update. A value of "*" can be used for If-Match to unconditionally apply the operation.
func (client IdentityProviderClient) Update(ctx context.Context, apimBaseURL string, identityProviderName IdentityProviderType, parameters IdentityProviderUpdateParameters, ifMatch string) (result autorest.Response, err error) {
	req, err := client.UpdatePreparer(ctx, apimBaseURL, identityProviderName, parameters, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.IdentityProviderClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.IdentityProviderClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.IdentityProviderClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client IdentityProviderClient) UpdatePreparer(ctx context.Context, apimBaseURL string, identityProviderName IdentityProviderType, parameters IdentityProviderUpdateParameters, ifMatch string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"apimBaseUrl": apimBaseURL,
	}

	pathParameters := map[string]interface{}{
		"identityProviderName": autorest.Encode("path", identityProviderName),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithCustomBaseURL("{apimBaseUrl}", urlParameters),
		autorest.WithPathParameters("/identityProviders/{identityProviderName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client IdentityProviderClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client IdentityProviderClient) UpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
