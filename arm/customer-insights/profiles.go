package customerinsights

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

// ProfilesClient is the the Azure Customer Insights management API provides a RESTful set of web services that
// interact with Azure Customer Insights service to manage your resources. The API has entities that capture the
// relationship between an end user and the Azure Customer Insights service.
type ProfilesClient struct {
	ManagementClient
}

// NewProfilesClient creates an instance of the ProfilesClient client.
func NewProfilesClient(subscriptionID string) ProfilesClient {
	return NewProfilesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProfilesClientWithBaseURI creates an instance of the ProfilesClient client.
func NewProfilesClientWithBaseURI(baseURI string, subscriptionID string) ProfilesClient {
	return ProfilesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// NewProfilesClientWithAuthFile creates an instance of the ProfilesClient client.
func NewProfilesClientWithAuthFile() (ProfilesClient, error) {
	c, err := NewWithAuthFile()
	return ProfilesClient{c}, err
}

// CreateOrUpdate creates a profile within a Hub, or updates an existing profile. This method may poll for completion.
// Polling can be canceled by passing the cancel channel argument. The channel will be used to cancel polling and any
// outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. hubName is the name of the hub. profileName is the name of the
// profile. parameters is parameters supplied to the create/delete Profile type operation
func (client ProfilesClient) CreateOrUpdate(resourceGroupName string, hubName string, profileName string, parameters ProfileResourceFormat, cancel <-chan struct{}) (<-chan ProfileResourceFormat, <-chan error) {
	resultChan := make(chan ProfileResourceFormat, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: profileName,
			Constraints: []validation.Constraint{{Target: "profileName", Name: validation.MaxLength, Rule: 128, Chain: nil},
				{Target: "profileName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "profileName", Name: validation.Pattern, Rule: `^[a-zA-Z][a-zA-Z0-9_]+$`, Chain: nil}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "customerinsights.ProfilesClient", "CreateOrUpdate")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result ProfileResourceFormat
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreateOrUpdatePreparer(resourceGroupName, hubName, profileName, parameters, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "customerinsights.ProfilesClient", "CreateOrUpdate", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateOrUpdateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "customerinsights.ProfilesClient", "CreateOrUpdate", resp, "Failure sending request")
			return
		}

		result, err = client.CreateOrUpdateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "customerinsights.ProfilesClient", "CreateOrUpdate", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ProfilesClient) CreateOrUpdatePreparer(resourceGroupName string, hubName string, profileName string, parameters ProfileResourceFormat, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":           autorest.Encode("path", hubName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-26"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/profiles/{profileName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ProfilesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ProfilesClient) CreateOrUpdateResponder(resp *http.Response) (result ProfileResourceFormat, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a profile within a hub This method may poll for completion. Polling can be canceled by passing the
// cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. hubName is the name of the hub. profileName is the name of the
// profile. localeCode is locale of profile to retrieve, default is en-us.
func (client ProfilesClient) Delete(resourceGroupName string, hubName string, profileName string, localeCode string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	resultChan := make(chan autorest.Response, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result autorest.Response
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.DeletePreparer(resourceGroupName, hubName, profileName, localeCode, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "customerinsights.ProfilesClient", "Delete", nil, "Failure preparing request")
			return
		}

		resp, err := client.DeleteSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "customerinsights.ProfilesClient", "Delete", resp, "Failure sending request")
			return
		}

		result, err = client.DeleteResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "customerinsights.ProfilesClient", "Delete", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// DeletePreparer prepares the Delete request.
func (client ProfilesClient) DeletePreparer(resourceGroupName string, hubName string, profileName string, localeCode string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":           autorest.Encode("path", hubName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-26"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(localeCode) > 0 {
		queryParameters["locale-code"] = autorest.Encode("query", localeCode)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/profiles/{profileName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ProfilesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ProfilesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets information about the specified profile.
//
// resourceGroupName is the name of the resource group. hubName is the name of the hub. profileName is the name of the
// profile. localeCode is locale of profile to retrieve, default is en-us.
func (client ProfilesClient) Get(resourceGroupName string, hubName string, profileName string, localeCode string) (result ProfileResourceFormat, err error) {
	req, err := client.GetPreparer(resourceGroupName, hubName, profileName, localeCode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ProfilesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "customerinsights.ProfilesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ProfilesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ProfilesClient) GetPreparer(resourceGroupName string, hubName string, profileName string, localeCode string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":           autorest.Encode("path", hubName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-26"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(localeCode) > 0 {
		queryParameters["locale-code"] = autorest.Encode("query", localeCode)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/profiles/{profileName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ProfilesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ProfilesClient) GetResponder(resp *http.Response) (result ProfileResourceFormat, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetEnrichingKpis gets the KPIs that enrich the profile Type identified by the supplied name. Enrichment happens
// through participants of the Interaction on an Interaction KPI and through Relationships for Profile KPIs.
//
// resourceGroupName is the name of the resource group. hubName is the name of the hub. profileName is the name of the
// profile.
func (client ProfilesClient) GetEnrichingKpis(resourceGroupName string, hubName string, profileName string) (result ListKpiDefinition, err error) {
	req, err := client.GetEnrichingKpisPreparer(resourceGroupName, hubName, profileName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ProfilesClient", "GetEnrichingKpis", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetEnrichingKpisSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "customerinsights.ProfilesClient", "GetEnrichingKpis", resp, "Failure sending request")
		return
	}

	result, err = client.GetEnrichingKpisResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ProfilesClient", "GetEnrichingKpis", resp, "Failure responding to request")
	}

	return
}

// GetEnrichingKpisPreparer prepares the GetEnrichingKpis request.
func (client ProfilesClient) GetEnrichingKpisPreparer(resourceGroupName string, hubName string, profileName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":           autorest.Encode("path", hubName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-26"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/profiles/{profileName}/getEnrichingKpis", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetEnrichingKpisSender sends the GetEnrichingKpis request. The method will close the
// http.Response Body if it receives an error.
func (client ProfilesClient) GetEnrichingKpisSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetEnrichingKpisResponder handles the response to the GetEnrichingKpis request. The method always
// closes the http.Response Body.
func (client ProfilesClient) GetEnrichingKpisResponder(resp *http.Response) (result ListKpiDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByHub gets all profile in the hub.
//
// resourceGroupName is the name of the resource group. hubName is the name of the hub. localeCode is locale of profile
// to retrieve, default is en-us.
func (client ProfilesClient) ListByHub(resourceGroupName string, hubName string, localeCode string) (result ProfileListResult, err error) {
	req, err := client.ListByHubPreparer(resourceGroupName, hubName, localeCode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ProfilesClient", "ListByHub", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByHubSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "customerinsights.ProfilesClient", "ListByHub", resp, "Failure sending request")
		return
	}

	result, err = client.ListByHubResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ProfilesClient", "ListByHub", resp, "Failure responding to request")
	}

	return
}

// ListByHubPreparer prepares the ListByHub request.
func (client ProfilesClient) ListByHubPreparer(resourceGroupName string, hubName string, localeCode string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":           autorest.Encode("path", hubName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-26"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(localeCode) > 0 {
		queryParameters["locale-code"] = autorest.Encode("query", localeCode)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/profiles", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByHubSender sends the ListByHub request. The method will close the
// http.Response Body if it receives an error.
func (client ProfilesClient) ListByHubSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByHubResponder handles the response to the ListByHub request. The method always
// closes the http.Response Body.
func (client ProfilesClient) ListByHubResponder(resp *http.Response) (result ProfileListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByHubNextResults retrieves the next set of results, if any.
func (client ProfilesClient) ListByHubNextResults(lastResults ProfileListResult) (result ProfileListResult, err error) {
	req, err := lastResults.ProfileListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "customerinsights.ProfilesClient", "ListByHub", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByHubSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "customerinsights.ProfilesClient", "ListByHub", resp, "Failure sending next results request")
	}

	result, err = client.ListByHubResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ProfilesClient", "ListByHub", resp, "Failure responding to next results request")
	}

	return
}

// ListByHubComplete gets all elements from the list without paging.
func (client ProfilesClient) ListByHubComplete(resourceGroupName string, hubName string, localeCode string, cancel <-chan struct{}) (<-chan ProfileResourceFormat, <-chan error) {
	resultChan := make(chan ProfileResourceFormat)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListByHub(resourceGroupName, hubName, localeCode)
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
			list, err = client.ListByHubNextResults(list)
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
