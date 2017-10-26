package web

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
	"net/http"
)

// CertificatesClient is the webSite Management Client
type CertificatesClient struct {
	ManagementClient
}

// NewCertificatesClient creates an instance of the CertificatesClient client.
func NewCertificatesClient(subscriptionID string) CertificatesClient {
	return NewCertificatesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCertificatesClientWithBaseURI creates an instance of the CertificatesClient client.
func NewCertificatesClientWithBaseURI(baseURI string, subscriptionID string) CertificatesClient {
	return CertificatesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdateCertificate sends the create or update certificate request.
//
// resourceGroupName is name of the resource group name is name of the certificate. certificateEnvelope is details of
// certificate if it exists already.
func (client CertificatesClient) CreateOrUpdateCertificate(resourceGroupName string, name string, certificateEnvelope Certificate) (result Certificate, err error) {
	req, err := client.CreateOrUpdateCertificatePreparer(resourceGroupName, name, certificateEnvelope)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.CertificatesClient", "CreateOrUpdateCertificate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateCertificateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.CertificatesClient", "CreateOrUpdateCertificate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateCertificateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.CertificatesClient", "CreateOrUpdateCertificate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdateCertificatePreparer prepares the CreateOrUpdateCertificate request.
func (client CertificatesClient) CreateOrUpdateCertificatePreparer(resourceGroupName string, name string, certificateEnvelope Certificate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/certificates/{name}", pathParameters),
		autorest.WithJSON(certificateEnvelope),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateCertificateSender sends the CreateOrUpdateCertificate request. The method will close the
// http.Response Body if it receives an error.
func (client CertificatesClient) CreateOrUpdateCertificateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateCertificateResponder handles the response to the CreateOrUpdateCertificate request. The method always
// closes the http.Response Body.
func (client CertificatesClient) CreateOrUpdateCertificateResponder(resp *http.Response) (result Certificate, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdateCsr sends the create or update csr request.
//
// resourceGroupName is name of the resource group name is name of the certificate. csrEnvelope is details of
// certificate signing request if it exists already.
func (client CertificatesClient) CreateOrUpdateCsr(resourceGroupName string, name string, csrEnvelope Csr) (result Csr, err error) {
	req, err := client.CreateOrUpdateCsrPreparer(resourceGroupName, name, csrEnvelope)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.CertificatesClient", "CreateOrUpdateCsr", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateCsrSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.CertificatesClient", "CreateOrUpdateCsr", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateCsrResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.CertificatesClient", "CreateOrUpdateCsr", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdateCsrPreparer prepares the CreateOrUpdateCsr request.
func (client CertificatesClient) CreateOrUpdateCsrPreparer(resourceGroupName string, name string, csrEnvelope Csr) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/csrs/{name}", pathParameters),
		autorest.WithJSON(csrEnvelope),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateCsrSender sends the CreateOrUpdateCsr request. The method will close the
// http.Response Body if it receives an error.
func (client CertificatesClient) CreateOrUpdateCsrSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateCsrResponder handles the response to the CreateOrUpdateCsr request. The method always
// closes the http.Response Body.
func (client CertificatesClient) CreateOrUpdateCsrResponder(resp *http.Response) (result Csr, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteCertificate sends the delete certificate request.
//
// resourceGroupName is name of the resource group name is name of the certificate to be deleted.
func (client CertificatesClient) DeleteCertificate(resourceGroupName string, name string) (result SetObject, err error) {
	req, err := client.DeleteCertificatePreparer(resourceGroupName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.CertificatesClient", "DeleteCertificate", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteCertificateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.CertificatesClient", "DeleteCertificate", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteCertificateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.CertificatesClient", "DeleteCertificate", resp, "Failure responding to request")
	}

	return
}

// DeleteCertificatePreparer prepares the DeleteCertificate request.
func (client CertificatesClient) DeleteCertificatePreparer(resourceGroupName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/certificates/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteCertificateSender sends the DeleteCertificate request. The method will close the
// http.Response Body if it receives an error.
func (client CertificatesClient) DeleteCertificateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteCertificateResponder handles the response to the DeleteCertificate request. The method always
// closes the http.Response Body.
func (client CertificatesClient) DeleteCertificateResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteCsr sends the delete csr request.
//
// resourceGroupName is name of the resource group name is name of the certificate signing request.
func (client CertificatesClient) DeleteCsr(resourceGroupName string, name string) (result SetObject, err error) {
	req, err := client.DeleteCsrPreparer(resourceGroupName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.CertificatesClient", "DeleteCsr", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteCsrSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.CertificatesClient", "DeleteCsr", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteCsrResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.CertificatesClient", "DeleteCsr", resp, "Failure responding to request")
	}

	return
}

// DeleteCsrPreparer prepares the DeleteCsr request.
func (client CertificatesClient) DeleteCsrPreparer(resourceGroupName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/csrs/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteCsrSender sends the DeleteCsr request. The method will close the
// http.Response Body if it receives an error.
func (client CertificatesClient) DeleteCsrSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteCsrResponder handles the response to the DeleteCsr request. The method always
// closes the http.Response Body.
func (client CertificatesClient) DeleteCsrResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetCertificate sends the get certificate request.
//
// resourceGroupName is name of the resource group name is name of the certificate.
func (client CertificatesClient) GetCertificate(resourceGroupName string, name string) (result Certificate, err error) {
	req, err := client.GetCertificatePreparer(resourceGroupName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.CertificatesClient", "GetCertificate", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetCertificateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.CertificatesClient", "GetCertificate", resp, "Failure sending request")
		return
	}

	result, err = client.GetCertificateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.CertificatesClient", "GetCertificate", resp, "Failure responding to request")
	}

	return
}

// GetCertificatePreparer prepares the GetCertificate request.
func (client CertificatesClient) GetCertificatePreparer(resourceGroupName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/certificates/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetCertificateSender sends the GetCertificate request. The method will close the
// http.Response Body if it receives an error.
func (client CertificatesClient) GetCertificateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetCertificateResponder handles the response to the GetCertificate request. The method always
// closes the http.Response Body.
func (client CertificatesClient) GetCertificateResponder(resp *http.Response) (result Certificate, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetCertificates sends the get certificates request.
//
// resourceGroupName is name of the resource group
func (client CertificatesClient) GetCertificates(resourceGroupName string) (result CertificateCollection, err error) {
	req, err := client.GetCertificatesPreparer(resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.CertificatesClient", "GetCertificates", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetCertificatesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.CertificatesClient", "GetCertificates", resp, "Failure sending request")
		return
	}

	result, err = client.GetCertificatesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.CertificatesClient", "GetCertificates", resp, "Failure responding to request")
	}

	return
}

// GetCertificatesPreparer prepares the GetCertificates request.
func (client CertificatesClient) GetCertificatesPreparer(resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/certificates", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetCertificatesSender sends the GetCertificates request. The method will close the
// http.Response Body if it receives an error.
func (client CertificatesClient) GetCertificatesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetCertificatesResponder handles the response to the GetCertificates request. The method always
// closes the http.Response Body.
func (client CertificatesClient) GetCertificatesResponder(resp *http.Response) (result CertificateCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetCertificatesNextResults retrieves the next set of results, if any.
func (client CertificatesClient) GetCertificatesNextResults(lastResults CertificateCollection) (result CertificateCollection, err error) {
	req, err := lastResults.CertificateCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.CertificatesClient", "GetCertificates", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.GetCertificatesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.CertificatesClient", "GetCertificates", resp, "Failure sending next results request")
	}

	result, err = client.GetCertificatesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.CertificatesClient", "GetCertificates", resp, "Failure responding to next results request")
	}

	return
}

// GetCertificatesComplete gets all elements from the list without paging.
func (client CertificatesClient) GetCertificatesComplete(resourceGroupName string, cancel <-chan struct{}) (<-chan Certificate, <-chan error) {
	resultChan := make(chan Certificate)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.GetCertificates(resourceGroupName)
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
			list, err = client.GetCertificatesNextResults(list)
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

// GetCsr sends the get csr request.
//
// resourceGroupName is name of the resource group name is name of the certificate.
func (client CertificatesClient) GetCsr(resourceGroupName string, name string) (result Csr, err error) {
	req, err := client.GetCsrPreparer(resourceGroupName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.CertificatesClient", "GetCsr", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetCsrSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.CertificatesClient", "GetCsr", resp, "Failure sending request")
		return
	}

	result, err = client.GetCsrResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.CertificatesClient", "GetCsr", resp, "Failure responding to request")
	}

	return
}

// GetCsrPreparer prepares the GetCsr request.
func (client CertificatesClient) GetCsrPreparer(resourceGroupName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/csrs/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetCsrSender sends the GetCsr request. The method will close the
// http.Response Body if it receives an error.
func (client CertificatesClient) GetCsrSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetCsrResponder handles the response to the GetCsr request. The method always
// closes the http.Response Body.
func (client CertificatesClient) GetCsrResponder(resp *http.Response) (result Csr, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetCsrs sends the get csrs request.
//
// resourceGroupName is name of the resource group
func (client CertificatesClient) GetCsrs(resourceGroupName string) (result ListCsr, err error) {
	req, err := client.GetCsrsPreparer(resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.CertificatesClient", "GetCsrs", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetCsrsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.CertificatesClient", "GetCsrs", resp, "Failure sending request")
		return
	}

	result, err = client.GetCsrsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.CertificatesClient", "GetCsrs", resp, "Failure responding to request")
	}

	return
}

// GetCsrsPreparer prepares the GetCsrs request.
func (client CertificatesClient) GetCsrsPreparer(resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/csrs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetCsrsSender sends the GetCsrs request. The method will close the
// http.Response Body if it receives an error.
func (client CertificatesClient) GetCsrsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetCsrsResponder handles the response to the GetCsrs request. The method always
// closes the http.Response Body.
func (client CertificatesClient) GetCsrsResponder(resp *http.Response) (result ListCsr, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UpdateCertificate sends the update certificate request.
//
// resourceGroupName is name of the resource group name is name of the certificate. certificateEnvelope is details of
// certificate if it exists already.
func (client CertificatesClient) UpdateCertificate(resourceGroupName string, name string, certificateEnvelope Certificate) (result Certificate, err error) {
	req, err := client.UpdateCertificatePreparer(resourceGroupName, name, certificateEnvelope)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.CertificatesClient", "UpdateCertificate", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateCertificateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.CertificatesClient", "UpdateCertificate", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateCertificateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.CertificatesClient", "UpdateCertificate", resp, "Failure responding to request")
	}

	return
}

// UpdateCertificatePreparer prepares the UpdateCertificate request.
func (client CertificatesClient) UpdateCertificatePreparer(resourceGroupName string, name string, certificateEnvelope Certificate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/certificates/{name}", pathParameters),
		autorest.WithJSON(certificateEnvelope),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// UpdateCertificateSender sends the UpdateCertificate request. The method will close the
// http.Response Body if it receives an error.
func (client CertificatesClient) UpdateCertificateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// UpdateCertificateResponder handles the response to the UpdateCertificate request. The method always
// closes the http.Response Body.
func (client CertificatesClient) UpdateCertificateResponder(resp *http.Response) (result Certificate, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UpdateCsr sends the update csr request.
//
// resourceGroupName is name of the resource group name is name of the certificate. csrEnvelope is details of
// certificate signing request if it exists already.
func (client CertificatesClient) UpdateCsr(resourceGroupName string, name string, csrEnvelope Csr) (result Csr, err error) {
	req, err := client.UpdateCsrPreparer(resourceGroupName, name, csrEnvelope)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.CertificatesClient", "UpdateCsr", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateCsrSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.CertificatesClient", "UpdateCsr", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateCsrResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.CertificatesClient", "UpdateCsr", resp, "Failure responding to request")
	}

	return
}

// UpdateCsrPreparer prepares the UpdateCsr request.
func (client CertificatesClient) UpdateCsrPreparer(resourceGroupName string, name string, csrEnvelope Csr) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/csrs/{name}", pathParameters),
		autorest.WithJSON(csrEnvelope),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// UpdateCsrSender sends the UpdateCsr request. The method will close the
// http.Response Body if it receives an error.
func (client CertificatesClient) UpdateCsrSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// UpdateCsrResponder handles the response to the UpdateCsr request. The method always
// closes the http.Response Body.
func (client CertificatesClient) UpdateCsrResponder(resp *http.Response) (result Csr, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
