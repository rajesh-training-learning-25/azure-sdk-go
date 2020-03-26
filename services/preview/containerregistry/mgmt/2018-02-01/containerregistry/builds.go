package containerregistry

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
    "context"
    "github.com/Azure/go-autorest/tracing"
    "github.com/Azure/go-autorest/autorest/validation"
)

// BuildsClient is the client for the Builds methods of the Containerregistry service.
type BuildsClient struct {
    BaseClient
}
// NewBuildsClient creates an instance of the BuildsClient client.
func NewBuildsClient(subscriptionID string) BuildsClient {
    return NewBuildsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewBuildsClientWithBaseURI creates an instance of the BuildsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewBuildsClientWithBaseURI(baseURI string, subscriptionID string) BuildsClient {
        return BuildsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// Cancel cancel an existing build.
    // Parameters:
        // resourceGroupName - the name of the resource group to which the container registry belongs.
        // registryName - the name of the container registry.
        // buildID - the build ID.
func (client BuildsClient) Cancel(ctx context.Context, resourceGroupName string, registryName string, buildID string) (result BuildsCancelFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/BuildsClient.Cancel")
        defer func() {
            sc := -1
            if result.Response() != nil {
                sc = result.Response().StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: registryName,
             Constraints: []validation.Constraint{	{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil },
            	{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil },
            	{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("containerregistry.BuildsClient", "Cancel", err.Error())
            }

                req, err := client.CancelPreparer(ctx, resourceGroupName, registryName, buildID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "containerregistry.BuildsClient", "Cancel", nil , "Failure preparing request")
    return
    }

            result, err = client.CancelSender(req)
            if err != nil {
            err = autorest.NewErrorWithError(err, "containerregistry.BuildsClient", "Cancel", result.Response(), "Failure sending request")
            return
            }

    return
    }

    // CancelPreparer prepares the Cancel request.
    func (client BuildsClient) CancelPreparer(ctx context.Context, resourceGroupName string, registryName string, buildID string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "buildId": autorest.Encode("path",buildID),
            "registryName": autorest.Encode("path",registryName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-02-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/builds/{buildId}/cancel",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CancelSender sends the Cancel request. The method will close the
    // http.Response Body if it receives an error.
    func (client BuildsClient) CancelSender(req *http.Request) (future BuildsCancelFuture, err error) {
            var resp *http.Response
            resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
            if err != nil {
            return
            }
            future.Future, err = azure.NewFutureFromResponse(resp)
            return
            }

// CancelResponder handles the response to the Cancel request. The method always
// closes the http.Response Body.
func (client BuildsClient) CancelResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Get gets the detailed information for a given build.
    // Parameters:
        // resourceGroupName - the name of the resource group to which the container registry belongs.
        // registryName - the name of the container registry.
        // buildID - the build ID.
func (client BuildsClient) Get(ctx context.Context, resourceGroupName string, registryName string, buildID string) (result Build, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/BuildsClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: registryName,
             Constraints: []validation.Constraint{	{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil },
            	{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil },
            	{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("containerregistry.BuildsClient", "Get", err.Error())
            }

                req, err := client.GetPreparer(ctx, resourceGroupName, registryName, buildID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "containerregistry.BuildsClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "containerregistry.BuildsClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "containerregistry.BuildsClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client BuildsClient) GetPreparer(ctx context.Context, resourceGroupName string, registryName string, buildID string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "buildId": autorest.Encode("path",buildID),
            "registryName": autorest.Encode("path",registryName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-02-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/builds/{buildId}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client BuildsClient) GetSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client BuildsClient) GetResponder(resp *http.Response) (result Build, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// GetLogLink gets a link to download the build logs.
    // Parameters:
        // resourceGroupName - the name of the resource group to which the container registry belongs.
        // registryName - the name of the container registry.
        // buildID - the build ID.
func (client BuildsClient) GetLogLink(ctx context.Context, resourceGroupName string, registryName string, buildID string) (result BuildGetLogResult, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/BuildsClient.GetLogLink")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: registryName,
             Constraints: []validation.Constraint{	{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil },
            	{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil },
            	{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("containerregistry.BuildsClient", "GetLogLink", err.Error())
            }

                req, err := client.GetLogLinkPreparer(ctx, resourceGroupName, registryName, buildID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "containerregistry.BuildsClient", "GetLogLink", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetLogLinkSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "containerregistry.BuildsClient", "GetLogLink", resp, "Failure sending request")
            return
            }

            result, err = client.GetLogLinkResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "containerregistry.BuildsClient", "GetLogLink", resp, "Failure responding to request")
            }

    return
    }

    // GetLogLinkPreparer prepares the GetLogLink request.
    func (client BuildsClient) GetLogLinkPreparer(ctx context.Context, resourceGroupName string, registryName string, buildID string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "buildId": autorest.Encode("path",buildID),
            "registryName": autorest.Encode("path",registryName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-02-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/builds/{buildId}/getLogLink",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetLogLinkSender sends the GetLogLink request. The method will close the
    // http.Response Body if it receives an error.
    func (client BuildsClient) GetLogLinkSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// GetLogLinkResponder handles the response to the GetLogLink request. The method always
// closes the http.Response Body.
func (client BuildsClient) GetLogLinkResponder(resp *http.Response) (result BuildGetLogResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// List gets all the builds for a registry.
    // Parameters:
        // resourceGroupName - the name of the resource group to which the container registry belongs.
        // registryName - the name of the container registry.
        // filter - the builds filter to apply on the operation.
        // top - $top is supported for get list of builds, which limits the maximum number of builds to return.
        // skipToken - $skipToken is supported on get list of builds, which provides the next page in the list of
        // builds.
func (client BuildsClient) List(ctx context.Context, resourceGroupName string, registryName string, filter string, top *int32, skipToken string) (result BuildListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/BuildsClient.List")
        defer func() {
            sc := -1
            if result.blr.Response.Response != nil {
                sc = result.blr.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: registryName,
             Constraints: []validation.Constraint{	{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil },
            	{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil },
            	{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("containerregistry.BuildsClient", "List", err.Error())
            }

                        result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, resourceGroupName, registryName, filter, top, skipToken)
    if err != nil {
    err = autorest.NewErrorWithError(err, "containerregistry.BuildsClient", "List", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListSender(req)
            if err != nil {
            result.blr.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "containerregistry.BuildsClient", "List", resp, "Failure sending request")
            return
            }

            result.blr, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "containerregistry.BuildsClient", "List", resp, "Failure responding to request")
            }

    return
    }

    // ListPreparer prepares the List request.
    func (client BuildsClient) ListPreparer(ctx context.Context, resourceGroupName string, registryName string, filter string, top *int32, skipToken string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "registryName": autorest.Encode("path",registryName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-02-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }
            if len(filter) > 0 {
            queryParameters["$filter"] = autorest.Encode("query",filter)
            }
            if top != nil {
            queryParameters["$top"] = autorest.Encode("query",*top)
            }
            if len(skipToken) > 0 {
            queryParameters["$skipToken"] = autorest.Encode("query",skipToken)
            }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/builds",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client BuildsClient) ListSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client BuildsClient) ListResponder(resp *http.Response) (result BuildListResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

            // listNextResults retrieves the next set of results, if any.
            func (client BuildsClient) listNextResults(ctx context.Context, lastResults BuildListResult) (result BuildListResult, err error) {
            req, err := lastResults.buildListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "containerregistry.BuildsClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "containerregistry.BuildsClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "containerregistry.BuildsClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListComplete enumerates all values, automatically crossing page boundaries as required.
    func (client BuildsClient) ListComplete(ctx context.Context, resourceGroupName string, registryName string, filter string, top *int32, skipToken string) (result BuildListResultIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/BuildsClient.List")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.List(ctx, resourceGroupName, registryName, filter, top, skipToken)
                return
        }

// Update patch the build properties.
    // Parameters:
        // resourceGroupName - the name of the resource group to which the container registry belongs.
        // registryName - the name of the container registry.
        // buildID - the build ID.
        // buildUpdateParameters - the build update properties.
func (client BuildsClient) Update(ctx context.Context, resourceGroupName string, registryName string, buildID string, buildUpdateParameters BuildUpdateParameters) (result BuildsUpdateFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/BuildsClient.Update")
        defer func() {
            sc := -1
            if result.Response() != nil {
                sc = result.Response().StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: registryName,
             Constraints: []validation.Constraint{	{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil },
            	{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil },
            	{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("containerregistry.BuildsClient", "Update", err.Error())
            }

                req, err := client.UpdatePreparer(ctx, resourceGroupName, registryName, buildID, buildUpdateParameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "containerregistry.BuildsClient", "Update", nil , "Failure preparing request")
    return
    }

            result, err = client.UpdateSender(req)
            if err != nil {
            err = autorest.NewErrorWithError(err, "containerregistry.BuildsClient", "Update", result.Response(), "Failure sending request")
            return
            }

    return
    }

    // UpdatePreparer prepares the Update request.
    func (client BuildsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, registryName string, buildID string, buildUpdateParameters BuildUpdateParameters) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "buildId": autorest.Encode("path",buildID),
            "registryName": autorest.Encode("path",registryName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-02-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPatch(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/builds/{buildId}",pathParameters),
    autorest.WithJSON(buildUpdateParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // UpdateSender sends the Update request. The method will close the
    // http.Response Body if it receives an error.
    func (client BuildsClient) UpdateSender(req *http.Request) (future BuildsUpdateFuture, err error) {
            var resp *http.Response
            resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
            if err != nil {
            return
            }
            future.Future, err = azure.NewFutureFromResponse(resp)
            return
            }

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client BuildsClient) UpdateResponder(resp *http.Response) (result Build, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

