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
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// OriginsClient is the use these APIs to manage Azure CDN resources through the Azure Resource Manager. You must make
// sure that requests made to these resources are secure. For more information, see
// https://msdn.microsoft.com/en-us/library/azure/dn790557.aspx.
type OriginsClient struct {
	ManagementClient
}

// NewOriginsClient creates an instance of the OriginsClient client.
func NewOriginsClient(subscriptionID string) OriginsClient {
	return NewOriginsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewOriginsClientWithBaseURI creates an instance of the OriginsClient client.
func NewOriginsClientWithBaseURI(baseURI string, subscriptionID string) OriginsClient {
	return OriginsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create sends the create request. This method may poll for completion. Polling can be canceled by passing the cancel
// channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// originName is name of the origin, an arbitrary value but it needs to be unique under endpoint originProperties is
// origin properties endpointName is name of the endpoint within the CDN profile. profileName is name of the CDN
// profile within the resource group. resourceGroupName is name of the resource group within the Azure subscription.
func (client OriginsClient) Create(originName string, originProperties OriginParameters, endpointName string, profileName string, resourceGroupName string, cancel <-chan struct{}) (<-chan Origin, <-chan error) {
	resultChan := make(chan Origin, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: originProperties,
			Constraints: []validation.Constraint{{Target: "originProperties.OriginPropertiesParameters", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "originProperties.OriginPropertiesParameters.HostName", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "cdn.OriginsClient", "Create")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result Origin
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreatePreparer(originName, originProperties, endpointName, profileName, resourceGroupName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "cdn.OriginsClient", "Create", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "cdn.OriginsClient", "Create", resp, "Failure sending request")
			return
		}

		result, err = client.CreateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "cdn.OriginsClient", "Create", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreatePreparer prepares the Create request.
func (client OriginsClient) CreatePreparer(originName string, originProperties OriginParameters, endpointName string, profileName string, resourceGroupName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"originName":        autorest.Encode("path", originName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-04-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/origins/{originName}", pathParameters),
		autorest.WithJSON(originProperties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client OriginsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client OriginsClient) CreateResponder(resp *http.Response) (result Origin, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteIfExists sends the delete if exists request. This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// originName is name of the origin. Must be unique within endpoint. endpointName is name of the endpoint within the
// CDN profile. profileName is name of the CDN profile within the resource group. resourceGroupName is name of the
// resource group within the Azure subscription.
func (client OriginsClient) DeleteIfExists(originName string, endpointName string, profileName string, resourceGroupName string, cancel <-chan struct{}) (<-chan Origin, <-chan error) {
	resultChan := make(chan Origin, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result Origin
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.DeleteIfExistsPreparer(originName, endpointName, profileName, resourceGroupName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "cdn.OriginsClient", "DeleteIfExists", nil, "Failure preparing request")
			return
		}

		resp, err := client.DeleteIfExistsSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "cdn.OriginsClient", "DeleteIfExists", resp, "Failure sending request")
			return
		}

		result, err = client.DeleteIfExistsResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "cdn.OriginsClient", "DeleteIfExists", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// DeleteIfExistsPreparer prepares the DeleteIfExists request.
func (client OriginsClient) DeleteIfExistsPreparer(originName string, endpointName string, profileName string, resourceGroupName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"originName":        autorest.Encode("path", originName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-04-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/origins/{originName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteIfExistsSender sends the DeleteIfExists request. The method will close the
// http.Response Body if it receives an error.
func (client OriginsClient) DeleteIfExistsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteIfExistsResponder handles the response to the DeleteIfExists request. The method always
// closes the http.Response Body.
func (client OriginsClient) DeleteIfExistsResponder(resp *http.Response) (result Origin, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get sends the get request.
//
// originName is name of the origin, an arbitrary value but it needs to be unique under endpoint endpointName is name
// of the endpoint within the CDN profile. profileName is name of the CDN profile within the resource group.
// resourceGroupName is name of the resource group within the Azure subscription.
func (client OriginsClient) Get(originName string, endpointName string, profileName string, resourceGroupName string) (result Origin, err error) {
	req, err := client.GetPreparer(originName, endpointName, profileName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.OriginsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cdn.OriginsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.OriginsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client OriginsClient) GetPreparer(originName string, endpointName string, profileName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"originName":        autorest.Encode("path", originName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-04-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/origins/{originName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client OriginsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client OriginsClient) GetResponder(resp *http.Response) (result Origin, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByEndpoint sends the list by endpoint request.
//
// endpointName is name of the endpoint within the CDN profile. profileName is name of the CDN profile within the
// resource group. resourceGroupName is name of the resource group within the Azure subscription.
func (client OriginsClient) ListByEndpoint(endpointName string, profileName string, resourceGroupName string) (result OriginListResult, err error) {
	req, err := client.ListByEndpointPreparer(endpointName, profileName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.OriginsClient", "ListByEndpoint", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByEndpointSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cdn.OriginsClient", "ListByEndpoint", resp, "Failure sending request")
		return
	}

	result, err = client.ListByEndpointResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.OriginsClient", "ListByEndpoint", resp, "Failure responding to request")
	}

	return
}

// ListByEndpointPreparer prepares the ListByEndpoint request.
func (client OriginsClient) ListByEndpointPreparer(endpointName string, profileName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-04-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/origins", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByEndpointSender sends the ListByEndpoint request. The method will close the
// http.Response Body if it receives an error.
func (client OriginsClient) ListByEndpointSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByEndpointResponder handles the response to the ListByEndpoint request. The method always
// closes the http.Response Body.
func (client OriginsClient) ListByEndpointResponder(resp *http.Response) (result OriginListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update sends the update request. This method may poll for completion. Polling can be canceled by passing the cancel
// channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// originName is name of the origin. Must be unique within endpoint. originProperties is origin properties endpointName
// is name of the endpoint within the CDN profile. profileName is name of the CDN profile within the resource group.
// resourceGroupName is name of the resource group within the Azure subscription.
func (client OriginsClient) Update(originName string, originProperties OriginParameters, endpointName string, profileName string, resourceGroupName string, cancel <-chan struct{}) (<-chan Origin, <-chan error) {
	resultChan := make(chan Origin, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result Origin
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.UpdatePreparer(originName, originProperties, endpointName, profileName, resourceGroupName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "cdn.OriginsClient", "Update", nil, "Failure preparing request")
			return
		}

		resp, err := client.UpdateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "cdn.OriginsClient", "Update", resp, "Failure sending request")
			return
		}

		result, err = client.UpdateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "cdn.OriginsClient", "Update", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// UpdatePreparer prepares the Update request.
func (client OriginsClient) UpdatePreparer(originName string, originProperties OriginParameters, endpointName string, profileName string, resourceGroupName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"originName":        autorest.Encode("path", originName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-04-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/origins/{originName}", pathParameters),
		autorest.WithJSON(originProperties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client OriginsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client OriginsClient) UpdateResponder(resp *http.Response) (result Origin, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
