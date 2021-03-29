package adhybridhealthservice

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

// AddsServicesUserPreferenceClient is the REST APIs for Azure Active Directory Connect Health
type AddsServicesUserPreferenceClient struct {
	BaseClient
}

// NewAddsServicesUserPreferenceClient creates an instance of the AddsServicesUserPreferenceClient client.
func NewAddsServicesUserPreferenceClient() AddsServicesUserPreferenceClient {
	return NewAddsServicesUserPreferenceClientWithBaseURI(DefaultBaseURI)
}

// NewAddsServicesUserPreferenceClientWithBaseURI creates an instance of the AddsServicesUserPreferenceClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewAddsServicesUserPreferenceClientWithBaseURI(baseURI string) AddsServicesUserPreferenceClient {
	return AddsServicesUserPreferenceClient{NewWithBaseURI(baseURI)}
}

// Add adds the user preferences for a given feature.
// Parameters:
// serviceName - the name of the service.
// featureName - the name of the feature.
// setting - the user preference setting.
func (client AddsServicesUserPreferenceClient) Add(ctx context.Context, serviceName string, featureName string, setting UserPreference) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AddsServicesUserPreferenceClient.Add")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.AddPreparer(ctx, serviceName, featureName, setting)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServicesUserPreferenceClient", "Add", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServicesUserPreferenceClient", "Add", resp, "Failure sending request")
		return
	}

	result, err = client.AddResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServicesUserPreferenceClient", "Add", resp, "Failure responding to request")
		return
	}

	return
}

// AddPreparer prepares the Add request.
func (client AddsServicesUserPreferenceClient) AddPreparer(ctx context.Context, serviceName string, featureName string, setting UserPreference) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"featureName": autorest.Encode("path", featureName),
		"serviceName": autorest.Encode("path", serviceName),
	}

	const APIVersion = "2014-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/features/{featureName}/userpreference", pathParameters),
		autorest.WithJSON(setting),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AddSender sends the Add request. The method will close the
// http.Response Body if it receives an error.
func (client AddsServicesUserPreferenceClient) AddSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// AddResponder handles the response to the Add request. The method always
// closes the http.Response Body.
func (client AddsServicesUserPreferenceClient) AddResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete deletes the user preferences for a given feature.
// Parameters:
// serviceName - the name of the service.
// featureName - the name of the feature.
func (client AddsServicesUserPreferenceClient) Delete(ctx context.Context, serviceName string, featureName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AddsServicesUserPreferenceClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, serviceName, featureName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServicesUserPreferenceClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServicesUserPreferenceClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServicesUserPreferenceClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client AddsServicesUserPreferenceClient) DeletePreparer(ctx context.Context, serviceName string, featureName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"featureName": autorest.Encode("path", featureName),
		"serviceName": autorest.Encode("path", serviceName),
	}

	const APIVersion = "2014-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/features/{featureName}/userpreference", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client AddsServicesUserPreferenceClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client AddsServicesUserPreferenceClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the user preferences for a given feature.
// Parameters:
// serviceName - the name of the service.
// featureName - the name of the feature.
func (client AddsServicesUserPreferenceClient) Get(ctx context.Context, serviceName string, featureName string) (result UserPreference, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AddsServicesUserPreferenceClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, serviceName, featureName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServicesUserPreferenceClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServicesUserPreferenceClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServicesUserPreferenceClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client AddsServicesUserPreferenceClient) GetPreparer(ctx context.Context, serviceName string, featureName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"featureName": autorest.Encode("path", featureName),
		"serviceName": autorest.Encode("path", serviceName),
	}

	const APIVersion = "2014-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/features/{featureName}/userpreference", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AddsServicesUserPreferenceClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AddsServicesUserPreferenceClient) GetResponder(resp *http.Response) (result UserPreference, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
