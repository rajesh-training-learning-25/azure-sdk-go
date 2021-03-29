package vmwarecloudsimple

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

// CustomizationPoliciesClient is the description of the new service
type CustomizationPoliciesClient struct {
	BaseClient
}

// NewCustomizationPoliciesClient creates an instance of the CustomizationPoliciesClient client.
func NewCustomizationPoliciesClient(subscriptionID string, referer string) CustomizationPoliciesClient {
	return NewCustomizationPoliciesClientWithBaseURI(DefaultBaseURI, subscriptionID, referer)
}

// NewCustomizationPoliciesClientWithBaseURI creates an instance of the CustomizationPoliciesClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewCustomizationPoliciesClientWithBaseURI(baseURI string, subscriptionID string, referer string) CustomizationPoliciesClient {
	return CustomizationPoliciesClient{NewWithBaseURI(baseURI, subscriptionID, referer)}
}

// Get returns customization policy by its name
// Parameters:
// regionID - the region Id (westus, eastus)
// pcName - the private cloud name
// customizationPolicyName - customization policy name
func (client CustomizationPoliciesClient) Get(ctx context.Context, regionID string, pcName string, customizationPolicyName string) (result CustomizationPolicy, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomizationPoliciesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, regionID, pcName, customizationPolicyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.CustomizationPoliciesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.CustomizationPoliciesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.CustomizationPoliciesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client CustomizationPoliciesClient) GetPreparer(ctx context.Context, regionID string, pcName string, customizationPolicyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"customizationPolicyName": autorest.Encode("path", customizationPolicyName),
		"pcName":                  autorest.Encode("path", pcName),
		"regionId":                autorest.Encode("path", regionID),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/customizationPolicies/{customizationPolicyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client CustomizationPoliciesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client CustomizationPoliciesClient) GetResponder(resp *http.Response) (result CustomizationPolicy, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List returns list of customization policies in region for private cloud
// Parameters:
// regionID - the region Id (westus, eastus)
// pcName - the private cloud name
// filter - the filter to apply on the list operation. only type is allowed here as a filter e.g. $filter=type
// eq 'xxxx'
func (client CustomizationPoliciesClient) List(ctx context.Context, regionID string, pcName string, filter string) (result CustomizationPoliciesListResponsePage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomizationPoliciesClient.List")
		defer func() {
			sc := -1
			if result.cplr.Response.Response != nil {
				sc = result.cplr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, regionID, pcName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.CustomizationPoliciesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.cplr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.CustomizationPoliciesClient", "List", resp, "Failure sending request")
		return
	}

	result.cplr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.CustomizationPoliciesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.cplr.hasNextLink() && result.cplr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client CustomizationPoliciesClient) ListPreparer(ctx context.Context, regionID string, pcName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"pcName":         autorest.Encode("path", pcName),
		"regionId":       autorest.Encode("path", regionID),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/customizationPolicies", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client CustomizationPoliciesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client CustomizationPoliciesClient) ListResponder(resp *http.Response) (result CustomizationPoliciesListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client CustomizationPoliciesClient) listNextResults(ctx context.Context, lastResults CustomizationPoliciesListResponse) (result CustomizationPoliciesListResponse, err error) {
	req, err := lastResults.customizationPoliciesListResponsePreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "vmwarecloudsimple.CustomizationPoliciesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "vmwarecloudsimple.CustomizationPoliciesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.CustomizationPoliciesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client CustomizationPoliciesClient) ListComplete(ctx context.Context, regionID string, pcName string, filter string) (result CustomizationPoliciesListResponseIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomizationPoliciesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, regionID, pcName, filter)
	return
}
