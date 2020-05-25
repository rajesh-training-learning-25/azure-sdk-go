package migrate

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

// VCenterClient is the discover your workloads for Azure.
type VCenterClient struct {
	BaseClient
}

// NewVCenterClient creates an instance of the VCenterClient client.
func NewVCenterClient() VCenterClient {
	return NewVCenterClientWithBaseURI(DefaultBaseURI)
}

// NewVCenterClientWithBaseURI creates an instance of the VCenterClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewVCenterClientWithBaseURI(baseURI string) VCenterClient {
	return VCenterClient{NewWithBaseURI(baseURI)}
}

// DeleteVCenter sends the delete v center request.
// Parameters:
// subscriptionID - the ID of the target subscription.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// siteName - site name.
// vcenterName - vCenter ARM name.
// APIVersion - the API version to use for this operation.
func (client VCenterClient) DeleteVCenter(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, vcenterName string, APIVersion string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VCenterClient.DeleteVCenter")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteVCenterPreparer(ctx, subscriptionID, resourceGroupName, siteName, vcenterName, APIVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.VCenterClient", "DeleteVCenter", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteVCenterSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "migrate.VCenterClient", "DeleteVCenter", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteVCenterResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.VCenterClient", "DeleteVCenter", resp, "Failure responding to request")
	}

	return
}

// DeleteVCenterPreparer prepares the DeleteVCenter request.
func (client VCenterClient) DeleteVCenterPreparer(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, vcenterName string, APIVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"siteName":          autorest.Encode("path", siteName),
		"subscriptionId":    autorest.Encode("path", subscriptionID),
		"vcenterName":       autorest.Encode("path", vcenterName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/VMwareSites/{siteName}/vCenters/{vcenterName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteVCenterSender sends the DeleteVCenter request. The method will close the
// http.Response Body if it receives an error.
func (client VCenterClient) DeleteVCenterSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteVCenterResponder handles the response to the DeleteVCenter request. The method always
// closes the http.Response Body.
func (client VCenterClient) DeleteVCenterResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetAllVCentersInSite sends the get all v centers in site request.
// Parameters:
// subscriptionID - the ID of the target subscription.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// siteName - site name.
// APIVersion - the API version to use for this operation.
func (client VCenterClient) GetAllVCentersInSite(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string, filter string) (result VCenterCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VCenterClient.GetAllVCentersInSite")
		defer func() {
			sc := -1
			if result.vcc.Response.Response != nil {
				sc = result.vcc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.getAllVCentersInSiteNextResults
	req, err := client.GetAllVCentersInSitePreparer(ctx, subscriptionID, resourceGroupName, siteName, APIVersion, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.VCenterClient", "GetAllVCentersInSite", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAllVCentersInSiteSender(req)
	if err != nil {
		result.vcc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.VCenterClient", "GetAllVCentersInSite", resp, "Failure sending request")
		return
	}

	result.vcc, err = client.GetAllVCentersInSiteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.VCenterClient", "GetAllVCentersInSite", resp, "Failure responding to request")
	}

	return
}

// GetAllVCentersInSitePreparer prepares the GetAllVCentersInSite request.
func (client VCenterClient) GetAllVCentersInSitePreparer(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"siteName":          autorest.Encode("path", siteName),
		"subscriptionId":    autorest.Encode("path", subscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/VMwareSites/{siteName}/vCenters", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetAllVCentersInSiteSender sends the GetAllVCentersInSite request. The method will close the
// http.Response Body if it receives an error.
func (client VCenterClient) GetAllVCentersInSiteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetAllVCentersInSiteResponder handles the response to the GetAllVCentersInSite request. The method always
// closes the http.Response Body.
func (client VCenterClient) GetAllVCentersInSiteResponder(resp *http.Response) (result VCenterCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getAllVCentersInSiteNextResults retrieves the next set of results, if any.
func (client VCenterClient) getAllVCentersInSiteNextResults(ctx context.Context, lastResults VCenterCollection) (result VCenterCollection, err error) {
	req, err := lastResults.vCenterCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "migrate.VCenterClient", "getAllVCentersInSiteNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetAllVCentersInSiteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "migrate.VCenterClient", "getAllVCentersInSiteNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetAllVCentersInSiteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.VCenterClient", "getAllVCentersInSiteNextResults", resp, "Failure responding to next results request")
	}
	return
}

// GetAllVCentersInSiteComplete enumerates all values, automatically crossing page boundaries as required.
func (client VCenterClient) GetAllVCentersInSiteComplete(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string, filter string) (result VCenterCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VCenterClient.GetAllVCentersInSite")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.GetAllVCentersInSite(ctx, subscriptionID, resourceGroupName, siteName, APIVersion, filter)
	return
}

// GetVCenter sends the get v center request.
// Parameters:
// subscriptionID - the ID of the target subscription.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// siteName - site name.
// vcenterName - vCenter ARM name.
// APIVersion - the API version to use for this operation.
func (client VCenterClient) GetVCenter(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, vcenterName string, APIVersion string) (result VCenter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VCenterClient.GetVCenter")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetVCenterPreparer(ctx, subscriptionID, resourceGroupName, siteName, vcenterName, APIVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.VCenterClient", "GetVCenter", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetVCenterSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.VCenterClient", "GetVCenter", resp, "Failure sending request")
		return
	}

	result, err = client.GetVCenterResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.VCenterClient", "GetVCenter", resp, "Failure responding to request")
	}

	return
}

// GetVCenterPreparer prepares the GetVCenter request.
func (client VCenterClient) GetVCenterPreparer(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, vcenterName string, APIVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"siteName":          autorest.Encode("path", siteName),
		"subscriptionId":    autorest.Encode("path", subscriptionID),
		"vcenterName":       autorest.Encode("path", vcenterName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/VMwareSites/{siteName}/vCenters/{vcenterName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetVCenterSender sends the GetVCenter request. The method will close the
// http.Response Body if it receives an error.
func (client VCenterClient) GetVCenterSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetVCenterResponder handles the response to the GetVCenter request. The method always
// closes the http.Response Body.
func (client VCenterClient) GetVCenterResponder(resp *http.Response) (result VCenter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PutVCenter sends the put v center request.
// Parameters:
// subscriptionID - the ID of the target subscription.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// siteName - site name.
// vcenterName - vCenter ARM name.
// body - put vCenter body.
// APIVersion - the API version to use for this operation.
func (client VCenterClient) PutVCenter(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, vcenterName string, body VCenter, APIVersion string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VCenterClient.PutVCenter")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.PutVCenterPreparer(ctx, subscriptionID, resourceGroupName, siteName, vcenterName, body, APIVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.VCenterClient", "PutVCenter", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutVCenterSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "migrate.VCenterClient", "PutVCenter", resp, "Failure sending request")
		return
	}

	result, err = client.PutVCenterResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.VCenterClient", "PutVCenter", resp, "Failure responding to request")
	}

	return
}

// PutVCenterPreparer prepares the PutVCenter request.
func (client VCenterClient) PutVCenterPreparer(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, vcenterName string, body VCenter, APIVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"siteName":          autorest.Encode("path", siteName),
		"subscriptionId":    autorest.Encode("path", subscriptionID),
		"vcenterName":       autorest.Encode("path", vcenterName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	body.ID = nil
	body.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/VMwareSites/{siteName}/vCenters/{vcenterName}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutVCenterSender sends the PutVCenter request. The method will close the
// http.Response Body if it receives an error.
func (client VCenterClient) PutVCenterSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// PutVCenterResponder handles the response to the PutVCenter request. The method always
// closes the http.Response Body.
func (client VCenterClient) PutVCenterResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
