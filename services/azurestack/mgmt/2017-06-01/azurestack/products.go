package azurestack

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

// ProductsClient is the azure Stack
type ProductsClient struct {
	BaseClient
}

// NewProductsClient creates an instance of the ProductsClient client.
func NewProductsClient(subscriptionID string) ProductsClient {
	return NewProductsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProductsClientWithBaseURI creates an instance of the ProductsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewProductsClientWithBaseURI(baseURI string, subscriptionID string) ProductsClient {
	return ProductsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get returns the specified product.
// Parameters:
// resourceGroup - name of the resource group.
// registrationName - name of the Azure Stack registration.
// productName - name of the product.
func (client ProductsClient) Get(ctx context.Context, resourceGroup string, registrationName string, productName string) (result Product, err error) {
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
	req, err := client.GetPreparer(ctx, resourceGroup, registrationName, productName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestack.ProductsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azurestack.ProductsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestack.ProductsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ProductsClient) GetPreparer(ctx context.Context, resourceGroup string, registrationName string, productName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"productName":      autorest.Encode("path", productName),
		"registrationName": autorest.Encode("path", registrationName),
		"resourceGroup":    autorest.Encode("path", resourceGroup),
		"subscriptionId":   autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}/products/{productName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ProductsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ProductsClient) GetResponder(resp *http.Response) (result Product, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetProduct returns the specified product.
// Parameters:
// resourceGroup - name of the resource group.
// registrationName - name of the Azure Stack registration.
// productName - name of the product.
// deviceConfiguration - device configuration.
func (client ProductsClient) GetProduct(ctx context.Context, resourceGroup string, registrationName string, productName string, deviceConfiguration *DeviceConfiguration) (result Product, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductsClient.GetProduct")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetProductPreparer(ctx, resourceGroup, registrationName, productName, deviceConfiguration)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestack.ProductsClient", "GetProduct", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetProductSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azurestack.ProductsClient", "GetProduct", resp, "Failure sending request")
		return
	}

	result, err = client.GetProductResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestack.ProductsClient", "GetProduct", resp, "Failure responding to request")
		return
	}

	return
}

// GetProductPreparer prepares the GetProduct request.
func (client ProductsClient) GetProductPreparer(ctx context.Context, resourceGroup string, registrationName string, productName string, deviceConfiguration *DeviceConfiguration) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"productName":      autorest.Encode("path", productName),
		"registrationName": autorest.Encode("path", registrationName),
		"resourceGroup":    autorest.Encode("path", resourceGroup),
		"subscriptionId":   autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	deviceConfiguration.DeviceVersion = nil
	deviceConfiguration.IdentitySystem = ""
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}/products/{productName}/getProduct", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if deviceConfiguration != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(deviceConfiguration))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetProductSender sends the GetProduct request. The method will close the
// http.Response Body if it receives an error.
func (client ProductsClient) GetProductSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetProductResponder handles the response to the GetProduct request. The method always
// closes the http.Response Body.
func (client ProductsClient) GetProductResponder(resp *http.Response) (result Product, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetProducts returns a list of products.
// Parameters:
// resourceGroup - name of the resource group.
// registrationName - name of the Azure Stack registration.
// productName - name of the product.
// deviceConfiguration - device configuration.
func (client ProductsClient) GetProducts(ctx context.Context, resourceGroup string, registrationName string, productName string, deviceConfiguration *DeviceConfiguration) (result ProductList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductsClient.GetProducts")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetProductsPreparer(ctx, resourceGroup, registrationName, productName, deviceConfiguration)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestack.ProductsClient", "GetProducts", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetProductsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azurestack.ProductsClient", "GetProducts", resp, "Failure sending request")
		return
	}

	result, err = client.GetProductsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestack.ProductsClient", "GetProducts", resp, "Failure responding to request")
		return
	}

	return
}

// GetProductsPreparer prepares the GetProducts request.
func (client ProductsClient) GetProductsPreparer(ctx context.Context, resourceGroup string, registrationName string, productName string, deviceConfiguration *DeviceConfiguration) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"productName":      autorest.Encode("path", productName),
		"registrationName": autorest.Encode("path", registrationName),
		"resourceGroup":    autorest.Encode("path", resourceGroup),
		"subscriptionId":   autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	deviceConfiguration.DeviceVersion = nil
	deviceConfiguration.IdentitySystem = ""
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}/products/{productName}/getProducts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if deviceConfiguration != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(deviceConfiguration))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetProductsSender sends the GetProducts request. The method will close the
// http.Response Body if it receives an error.
func (client ProductsClient) GetProductsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetProductsResponder handles the response to the GetProducts request. The method always
// closes the http.Response Body.
func (client ProductsClient) GetProductsResponder(resp *http.Response) (result ProductList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List returns a list of products.
// Parameters:
// resourceGroup - name of the resource group.
// registrationName - name of the Azure Stack registration.
func (client ProductsClient) List(ctx context.Context, resourceGroup string, registrationName string) (result ProductListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductsClient.List")
		defer func() {
			sc := -1
			if result.pl.Response.Response != nil {
				sc = result.pl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroup, registrationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestack.ProductsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.pl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azurestack.ProductsClient", "List", resp, "Failure sending request")
		return
	}

	result.pl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestack.ProductsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.pl.hasNextLink() && result.pl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ProductsClient) ListPreparer(ctx context.Context, resourceGroup string, registrationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registrationName": autorest.Encode("path", registrationName),
		"resourceGroup":    autorest.Encode("path", resourceGroup),
		"subscriptionId":   autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}/products", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ProductsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ProductsClient) ListResponder(resp *http.Response) (result ProductList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ProductsClient) listNextResults(ctx context.Context, lastResults ProductList) (result ProductList, err error) {
	req, err := lastResults.productListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "azurestack.ProductsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "azurestack.ProductsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestack.ProductsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProductsClient) ListComplete(ctx context.Context, resourceGroup string, registrationName string) (result ProductListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroup, registrationName)
	return
}

// ListDetails returns the extended properties of a product.
// Parameters:
// resourceGroup - name of the resource group.
// registrationName - name of the Azure Stack registration.
// productName - name of the product.
func (client ProductsClient) ListDetails(ctx context.Context, resourceGroup string, registrationName string, productName string) (result ExtendedProduct, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductsClient.ListDetails")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListDetailsPreparer(ctx, resourceGroup, registrationName, productName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestack.ProductsClient", "ListDetails", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListDetailsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azurestack.ProductsClient", "ListDetails", resp, "Failure sending request")
		return
	}

	result, err = client.ListDetailsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestack.ProductsClient", "ListDetails", resp, "Failure responding to request")
		return
	}

	return
}

// ListDetailsPreparer prepares the ListDetails request.
func (client ProductsClient) ListDetailsPreparer(ctx context.Context, resourceGroup string, registrationName string, productName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"productName":      autorest.Encode("path", productName),
		"registrationName": autorest.Encode("path", registrationName),
		"resourceGroup":    autorest.Encode("path", resourceGroup),
		"subscriptionId":   autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}/products/{productName}/listDetails", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListDetailsSender sends the ListDetails request. The method will close the
// http.Response Body if it receives an error.
func (client ProductsClient) ListDetailsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListDetailsResponder handles the response to the ListDetails request. The method always
// closes the http.Response Body.
func (client ProductsClient) ListDetailsResponder(resp *http.Response) (result ExtendedProduct, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UploadLog returns the specified product.
// Parameters:
// resourceGroup - name of the resource group.
// registrationName - name of the Azure Stack registration.
// productName - name of the product.
// marketplaceProductLogUpdate - update details for product log.
func (client ProductsClient) UploadLog(ctx context.Context, resourceGroup string, registrationName string, productName string, marketplaceProductLogUpdate *MarketplaceProductLogUpdate) (result ProductLog, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductsClient.UploadLog")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UploadLogPreparer(ctx, resourceGroup, registrationName, productName, marketplaceProductLogUpdate)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestack.ProductsClient", "UploadLog", nil, "Failure preparing request")
		return
	}

	resp, err := client.UploadLogSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azurestack.ProductsClient", "UploadLog", resp, "Failure sending request")
		return
	}

	result, err = client.UploadLogResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestack.ProductsClient", "UploadLog", resp, "Failure responding to request")
		return
	}

	return
}

// UploadLogPreparer prepares the UploadLog request.
func (client ProductsClient) UploadLogPreparer(ctx context.Context, resourceGroup string, registrationName string, productName string, marketplaceProductLogUpdate *MarketplaceProductLogUpdate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"productName":      autorest.Encode("path", productName),
		"registrationName": autorest.Encode("path", registrationName),
		"resourceGroup":    autorest.Encode("path", resourceGroup),
		"subscriptionId":   autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	marketplaceProductLogUpdate.Operation = nil
	marketplaceProductLogUpdate.Status = nil
	marketplaceProductLogUpdate.Error = nil
	marketplaceProductLogUpdate.Details = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}/products/{productName}/uploadProductLog", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if marketplaceProductLogUpdate != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(marketplaceProductLogUpdate))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UploadLogSender sends the UploadLog request. The method will close the
// http.Response Body if it receives an error.
func (client ProductsClient) UploadLogSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UploadLogResponder handles the response to the UploadLog request. The method always
// closes the http.Response Body.
func (client ProductsClient) UploadLogResponder(resp *http.Response) (result ProductLog, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
