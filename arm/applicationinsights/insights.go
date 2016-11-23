package applicationinsights

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
// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/azure"
    "net/http"
)

// InsightsClient is the azure Application Insights client
type InsightsClient struct {
    ManagementClient
}
// NewInsightsClient creates an instance of the InsightsClient client.
func NewInsightsClient(subscriptionID string, resourceGroupName string, name string) InsightsClient {
    return NewInsightsClientWithBaseURI(DefaultBaseURI, subscriptionID, resourceGroupName, name)
}

// NewInsightsClientWithBaseURI creates an instance of the InsightsClient
// client.
func NewInsightsClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string, name string) InsightsClient {
    return InsightsClient{NewWithBaseURI(baseURI, subscriptionID, resourceGroupName, name)}
}

// CreateUpdate creates or Updates an Application Insights instance
func (client InsightsClient) CreateUpdate() (result Resource, err error) {
    req, err := client.CreateUpdatePreparer()
    if err != nil {
        return result, autorest.NewErrorWithError(err, "applicationinsights.InsightsClient", "CreateUpdate", nil , "Failure preparing request")
    }

    resp, err := client.CreateUpdateSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        return result, autorest.NewErrorWithError(err, "applicationinsights.InsightsClient", "CreateUpdate", resp, "Failure sending request")
    }

    result, err = client.CreateUpdateResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "applicationinsights.InsightsClient", "CreateUpdate", resp, "Failure responding to request")
    }

    return
}

// CreateUpdatePreparer prepares the CreateUpdate request.
func (client InsightsClient) CreateUpdatePreparer() (*http.Request, error) {
    pathParameters := map[string]interface{} {
    "name": autorest.Encode("path",client.Name),
    "resourceGroupName": autorest.Encode("path",client.ResourceGroupName),
    "subscriptionId": autorest.Encode("path",client.SubscriptionID),
    }

    queryParameters := map[string]interface{} {
    "api-version": client.APIVersion,
    }

    preparer := autorest.CreatePreparer(
                        autorest.AsPost(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{name}",pathParameters),
                        autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare(&http.Request{})
}

// CreateUpdateSender sends the CreateUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client InsightsClient) CreateUpdateSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// CreateUpdateResponder handles the response to the CreateUpdate request. The method always
// closes the http.Response Body.
func (client InsightsClient) CreateUpdateResponder(resp *http.Response) (result Resource, err error) { 
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNotFound),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
    return
}

// CreateUpdateNextResults retrieves the next set of results, if any.
func (client InsightsClient) CreateUpdateNextResults(lastResults Resource) (result Resource, err error) {
    req, err := lastResults.ResourcePreparer()
    if err != nil {
        return result, autorest.NewErrorWithError(err, "applicationinsights.InsightsClient", "CreateUpdate", nil , "Failure preparing next results request")
    }
    if req == nil {
        return
    }

    resp, err := client.CreateUpdateSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        return result, autorest.NewErrorWithError(err, "applicationinsights.InsightsClient", "CreateUpdate", resp, "Failure sending next results request")
    }

    result, err = client.CreateUpdateResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "applicationinsights.InsightsClient", "CreateUpdate", resp, "Failure responding to next results request")
    }

    return
}

// Delete deletes an Application Insights instance
func (client InsightsClient) Delete() (result autorest.Response, err error) {
    req, err := client.DeletePreparer()
    if err != nil {
        return result, autorest.NewErrorWithError(err, "applicationinsights.InsightsClient", "Delete", nil , "Failure preparing request")
    }

    resp, err := client.DeleteSender(req)
    if err != nil {
        result.Response = resp
        return result, autorest.NewErrorWithError(err, "applicationinsights.InsightsClient", "Delete", resp, "Failure sending request")
    }

    result, err = client.DeleteResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "applicationinsights.InsightsClient", "Delete", resp, "Failure responding to request")
    }

    return
}

// DeletePreparer prepares the Delete request.
func (client InsightsClient) DeletePreparer() (*http.Request, error) {
    pathParameters := map[string]interface{} {
    "name": autorest.Encode("path",client.Name),
    "resourceGroupName": autorest.Encode("path",client.ResourceGroupName),
    "subscriptionId": autorest.Encode("path",client.SubscriptionID),
    }

    queryParameters := map[string]interface{} {
    "api-version": client.APIVersion,
    }

    preparer := autorest.CreatePreparer(
                        autorest.AsDelete(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{name}",pathParameters),
                        autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client InsightsClient) DeleteSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client InsightsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) { 
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNotFound),
            autorest.ByClosing())
    result.Response = resp
    return
}

// Get returns an Application Insights instance
func (client InsightsClient) Get() (result Resource, err error) {
    req, err := client.GetPreparer()
    if err != nil {
        return result, autorest.NewErrorWithError(err, "applicationinsights.InsightsClient", "Get", nil , "Failure preparing request")
    }

    resp, err := client.GetSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        return result, autorest.NewErrorWithError(err, "applicationinsights.InsightsClient", "Get", resp, "Failure sending request")
    }

    result, err = client.GetResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "applicationinsights.InsightsClient", "Get", resp, "Failure responding to request")
    }

    return
}

// GetPreparer prepares the Get request.
func (client InsightsClient) GetPreparer() (*http.Request, error) {
    pathParameters := map[string]interface{} {
    "name": autorest.Encode("path",client.Name),
    "resourceGroupName": autorest.Encode("path",client.ResourceGroupName),
    "subscriptionId": autorest.Encode("path",client.SubscriptionID),
    }

    queryParameters := map[string]interface{} {
    "api-version": client.APIVersion,
    }

    preparer := autorest.CreatePreparer(
                        autorest.AsGet(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{name}",pathParameters),
                        autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client InsightsClient) GetSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client InsightsClient) GetResponder(resp *http.Response) (result Resource, err error) { 
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
    return
}

// GetNextResults retrieves the next set of results, if any.
func (client InsightsClient) GetNextResults(lastResults Resource) (result Resource, err error) {
    req, err := lastResults.ResourcePreparer()
    if err != nil {
        return result, autorest.NewErrorWithError(err, "applicationinsights.InsightsClient", "Get", nil , "Failure preparing next results request")
    }
    if req == nil {
        return
    }

    resp, err := client.GetSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        return result, autorest.NewErrorWithError(err, "applicationinsights.InsightsClient", "Get", resp, "Failure sending next results request")
    }

    result, err = client.GetResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "applicationinsights.InsightsClient", "Get", resp, "Failure responding to next results request")
    }

    return
}

