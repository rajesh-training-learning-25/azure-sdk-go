package cdn

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

// CustomDomainsClient is the use these APIs to manage Azure CDN resources through the Azure Resource Manager. You must
// make sure that requests made to these resources are secure.
type CustomDomainsClient struct {
	ManagementClient
}

// NewCustomDomainsClient creates an instance of the CustomDomainsClient client.
func NewCustomDomainsClient(subscriptionID string) CustomDomainsClient {
	return NewCustomDomainsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCustomDomainsClientWithBaseURI creates an instance of the CustomDomainsClient client.
func NewCustomDomainsClientWithBaseURI(baseURI string, subscriptionID string) CustomDomainsClient {
	return CustomDomainsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates a new custom domain within an endpoint. This method may poll for completion. Polling can be canceled
// by passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP
// requests.
//
// resourceGroupName is name of the Resource group within the Azure subscription. profileName is name of the CDN
// profile which is unique within the resource group. endpointName is name of the endpoint under the profile which is
// unique globally. customDomainName is name of the custom domain within an endpoint. customDomainProperties is
// properties required to create a new custom domain.
func (client CustomDomainsClient) Create(resourceGroupName string, profileName string, endpointName string, customDomainName string, customDomainProperties CustomDomainParameters, cancel <-chan struct{}) (<-chan CustomDomain, <-chan error) {
	resultChan := make(chan CustomDomain, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: customDomainProperties,
			Constraints: []validation.Constraint{{Target: "customDomainProperties.CustomDomainPropertiesParameters", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "customDomainProperties.CustomDomainPropertiesParameters.HostName", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "cdn.CustomDomainsClient", "Create")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result CustomDomain
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreatePreparer(resourceGroupName, profileName, endpointName, customDomainName, customDomainProperties, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "Create", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "Create", resp, "Failure sending request")
			return
		}

		result, err = client.CreateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "Create", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreatePreparer prepares the Create request.
func (client CustomDomainsClient) CreatePreparer(resourceGroupName string, profileName string, endpointName string, customDomainName string, customDomainProperties CustomDomainParameters, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"customDomainName":  autorest.Encode("path", customDomainName),
		"endpointName":      autorest.Encode("path", endpointName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/customDomains/{customDomainName}", pathParameters),
		autorest.WithJSON(customDomainProperties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client CustomDomainsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client CustomDomainsClient) CreateResponder(resp *http.Response) (result CustomDomain, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes an existing custom domain within an endpoint. This method may poll for completion. Polling can be
// canceled by passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP
// requests.
//
// resourceGroupName is name of the Resource group within the Azure subscription. profileName is name of the CDN
// profile which is unique within the resource group. endpointName is name of the endpoint under the profile which is
// unique globally. customDomainName is name of the custom domain within an endpoint.
func (client CustomDomainsClient) Delete(resourceGroupName string, profileName string, endpointName string, customDomainName string, cancel <-chan struct{}) (<-chan CustomDomain, <-chan error) {
	resultChan := make(chan CustomDomain, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "cdn.CustomDomainsClient", "Delete")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result CustomDomain
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.DeletePreparer(resourceGroupName, profileName, endpointName, customDomainName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "Delete", nil, "Failure preparing request")
			return
		}

		resp, err := client.DeleteSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "Delete", resp, "Failure sending request")
			return
		}

		result, err = client.DeleteResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "Delete", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// DeletePreparer prepares the Delete request.
func (client CustomDomainsClient) DeletePreparer(resourceGroupName string, profileName string, endpointName string, customDomainName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"customDomainName":  autorest.Encode("path", customDomainName),
		"endpointName":      autorest.Encode("path", endpointName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/customDomains/{customDomainName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client CustomDomainsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client CustomDomainsClient) DeleteResponder(resp *http.Response) (result CustomDomain, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DisableCustomHTTPS disable https delivery of the custom domain.
//
// resourceGroupName is name of the Resource group within the Azure subscription. profileName is name of the CDN
// profile which is unique within the resource group. endpointName is name of the endpoint under the profile which is
// unique globally. customDomainName is name of the custom domain within an endpoint.
func (client CustomDomainsClient) DisableCustomHTTPS(resourceGroupName string, profileName string, endpointName string, customDomainName string) (result CustomDomain, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "cdn.CustomDomainsClient", "DisableCustomHTTPS")
	}

	req, err := client.DisableCustomHTTPSPreparer(resourceGroupName, profileName, endpointName, customDomainName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "DisableCustomHTTPS", nil, "Failure preparing request")
		return
	}

	resp, err := client.DisableCustomHTTPSSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "DisableCustomHTTPS", resp, "Failure sending request")
		return
	}

	result, err = client.DisableCustomHTTPSResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "DisableCustomHTTPS", resp, "Failure responding to request")
	}

	return
}

// DisableCustomHTTPSPreparer prepares the DisableCustomHTTPS request.
func (client CustomDomainsClient) DisableCustomHTTPSPreparer(resourceGroupName string, profileName string, endpointName string, customDomainName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"customDomainName":  autorest.Encode("path", customDomainName),
		"endpointName":      autorest.Encode("path", endpointName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/customDomains/{customDomainName}/disableCustomHttps", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DisableCustomHTTPSSender sends the DisableCustomHTTPS request. The method will close the
// http.Response Body if it receives an error.
func (client CustomDomainsClient) DisableCustomHTTPSSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DisableCustomHTTPSResponder handles the response to the DisableCustomHTTPS request. The method always
// closes the http.Response Body.
func (client CustomDomainsClient) DisableCustomHTTPSResponder(resp *http.Response) (result CustomDomain, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// EnableCustomHTTPS enable https delivery of the custom domain.
//
// resourceGroupName is name of the Resource group within the Azure subscription. profileName is name of the CDN
// profile which is unique within the resource group. endpointName is name of the endpoint under the profile which is
// unique globally. customDomainName is name of the custom domain within an endpoint.
func (client CustomDomainsClient) EnableCustomHTTPS(resourceGroupName string, profileName string, endpointName string, customDomainName string) (result CustomDomain, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "cdn.CustomDomainsClient", "EnableCustomHTTPS")
	}

	req, err := client.EnableCustomHTTPSPreparer(resourceGroupName, profileName, endpointName, customDomainName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "EnableCustomHTTPS", nil, "Failure preparing request")
		return
	}

	resp, err := client.EnableCustomHTTPSSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "EnableCustomHTTPS", resp, "Failure sending request")
		return
	}

	result, err = client.EnableCustomHTTPSResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "EnableCustomHTTPS", resp, "Failure responding to request")
	}

	return
}

// EnableCustomHTTPSPreparer prepares the EnableCustomHTTPS request.
func (client CustomDomainsClient) EnableCustomHTTPSPreparer(resourceGroupName string, profileName string, endpointName string, customDomainName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"customDomainName":  autorest.Encode("path", customDomainName),
		"endpointName":      autorest.Encode("path", endpointName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/customDomains/{customDomainName}/enableCustomHttps", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// EnableCustomHTTPSSender sends the EnableCustomHTTPS request. The method will close the
// http.Response Body if it receives an error.
func (client CustomDomainsClient) EnableCustomHTTPSSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// EnableCustomHTTPSResponder handles the response to the EnableCustomHTTPS request. The method always
// closes the http.Response Body.
func (client CustomDomainsClient) EnableCustomHTTPSResponder(resp *http.Response) (result CustomDomain, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets an exisitng custom domain within an endpoint.
//
// resourceGroupName is name of the Resource group within the Azure subscription. profileName is name of the CDN
// profile which is unique within the resource group. endpointName is name of the endpoint under the profile which is
// unique globally. customDomainName is name of the custom domain within an endpoint.
func (client CustomDomainsClient) Get(resourceGroupName string, profileName string, endpointName string, customDomainName string) (result CustomDomain, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "cdn.CustomDomainsClient", "Get")
	}

	req, err := client.GetPreparer(resourceGroupName, profileName, endpointName, customDomainName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client CustomDomainsClient) GetPreparer(resourceGroupName string, profileName string, endpointName string, customDomainName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"customDomainName":  autorest.Encode("path", customDomainName),
		"endpointName":      autorest.Encode("path", endpointName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/customDomains/{customDomainName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client CustomDomainsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client CustomDomainsClient) GetResponder(resp *http.Response) (result CustomDomain, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByEndpoint lists all of the existing custom domains within an endpoint.
//
// resourceGroupName is name of the Resource group within the Azure subscription. profileName is name of the CDN
// profile which is unique within the resource group. endpointName is name of the endpoint under the profile which is
// unique globally.
func (client CustomDomainsClient) ListByEndpoint(resourceGroupName string, profileName string, endpointName string) (result CustomDomainListResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "cdn.CustomDomainsClient", "ListByEndpoint")
	}

	req, err := client.ListByEndpointPreparer(resourceGroupName, profileName, endpointName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "ListByEndpoint", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByEndpointSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "ListByEndpoint", resp, "Failure sending request")
		return
	}

	result, err = client.ListByEndpointResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "ListByEndpoint", resp, "Failure responding to request")
	}

	return
}

// ListByEndpointPreparer prepares the ListByEndpoint request.
func (client CustomDomainsClient) ListByEndpointPreparer(resourceGroupName string, profileName string, endpointName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/customDomains", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByEndpointSender sends the ListByEndpoint request. The method will close the
// http.Response Body if it receives an error.
func (client CustomDomainsClient) ListByEndpointSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByEndpointResponder handles the response to the ListByEndpoint request. The method always
// closes the http.Response Body.
func (client CustomDomainsClient) ListByEndpointResponder(resp *http.Response) (result CustomDomainListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByEndpointNextResults retrieves the next set of results, if any.
func (client CustomDomainsClient) ListByEndpointNextResults(lastResults CustomDomainListResult) (result CustomDomainListResult, err error) {
	req, err := lastResults.CustomDomainListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "ListByEndpoint", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByEndpointSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "ListByEndpoint", resp, "Failure sending next results request")
	}

	result, err = client.ListByEndpointResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "ListByEndpoint", resp, "Failure responding to next results request")
	}

	return
}

// ListByEndpointComplete gets all elements from the list without paging.
func (client CustomDomainsClient) ListByEndpointComplete(resourceGroupName string, profileName string, endpointName string, cancel <-chan struct{}) (<-chan CustomDomain, <-chan error) {
	resultChan := make(chan CustomDomain)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListByEndpoint(resourceGroupName, profileName, endpointName)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListByEndpointNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}
