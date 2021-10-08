package batchai

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ClustersClient is the the Azure BatchAI Management API.
type ClustersClient struct {
	BaseClient
}

// NewClustersClient creates an instance of the ClustersClient client.
func NewClustersClient(subscriptionID string) ClustersClient {
	return NewClustersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewClustersClientWithBaseURI creates an instance of the ClustersClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewClustersClientWithBaseURI(baseURI string, subscriptionID string) ClustersClient {
	return ClustersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates a Cluster in the given Workspace.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// workspaceName - the name of the workspace. Workspace names can only contain a combination of alphanumeric
// characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
// clusterName - the name of the cluster within the specified resource group. Cluster names can only contain a
// combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1
// through 64 characters long.
// parameters - the parameters to provide for the Cluster creation.
func (client ClustersClient) Create(ctx context.Context, resourceGroupName string, workspaceName string, clusterName string, parameters ClusterCreateParameters) (result ClustersCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClustersClient.Create")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `^[-\w_]+$`, Chain: nil}}},
		{TargetValue: clusterName,
			Constraints: []validation.Constraint{{Target: "clusterName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "clusterName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "clusterName", Name: validation.Pattern, Rule: `^[-\w_]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.ClusterBaseProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.VMSize", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.ClusterBaseProperties.ScaleSettings", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.ScaleSettings.Manual", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.ScaleSettings.Manual.TargetNodeCount", Name: validation.Null, Rule: true, Chain: nil}}},
							{Target: "parameters.ClusterBaseProperties.ScaleSettings.AutoScale", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.ScaleSettings.AutoScale.MinimumNodeCount", Name: validation.Null, Rule: true, Chain: nil},
									{Target: "parameters.ClusterBaseProperties.ScaleSettings.AutoScale.MaximumNodeCount", Name: validation.Null, Rule: true, Chain: nil},
								}},
						}},
					{Target: "parameters.ClusterBaseProperties.VirtualMachineConfiguration", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.VirtualMachineConfiguration.ImageReference", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.VirtualMachineConfiguration.ImageReference.Publisher", Name: validation.Null, Rule: true, Chain: nil},
								{Target: "parameters.ClusterBaseProperties.VirtualMachineConfiguration.ImageReference.Offer", Name: validation.Null, Rule: true, Chain: nil},
								{Target: "parameters.ClusterBaseProperties.VirtualMachineConfiguration.ImageReference.Sku", Name: validation.Null, Rule: true, Chain: nil},
							}},
						}},
					{Target: "parameters.ClusterBaseProperties.NodeSetup", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.NodeSetup.SetupTask", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.NodeSetup.SetupTask.CommandLine", Name: validation.Null, Rule: true, Chain: nil},
								{Target: "parameters.ClusterBaseProperties.NodeSetup.SetupTask.StdOutErrPathPrefix", Name: validation.Null, Rule: true, Chain: nil},
							}},
							{Target: "parameters.ClusterBaseProperties.NodeSetup.PerformanceCountersSettings", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.NodeSetup.PerformanceCountersSettings.AppInsightsReference", Name: validation.Null, Rule: true,
									Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.NodeSetup.PerformanceCountersSettings.AppInsightsReference.Component", Name: validation.Null, Rule: true,
										Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.NodeSetup.PerformanceCountersSettings.AppInsightsReference.Component.ID", Name: validation.Null, Rule: true, Chain: nil}}},
										{Target: "parameters.ClusterBaseProperties.NodeSetup.PerformanceCountersSettings.AppInsightsReference.InstrumentationKeySecretReference", Name: validation.Null, Rule: false,
											Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.NodeSetup.PerformanceCountersSettings.AppInsightsReference.InstrumentationKeySecretReference.SourceVault", Name: validation.Null, Rule: true,
												Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.NodeSetup.PerformanceCountersSettings.AppInsightsReference.InstrumentationKeySecretReference.SourceVault.ID", Name: validation.Null, Rule: true, Chain: nil}}},
												{Target: "parameters.ClusterBaseProperties.NodeSetup.PerformanceCountersSettings.AppInsightsReference.InstrumentationKeySecretReference.SecretURL", Name: validation.Null, Rule: true, Chain: nil},
											}},
									}},
								}},
						}},
					{Target: "parameters.ClusterBaseProperties.UserAccountSettings", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.UserAccountSettings.AdminUserName", Name: validation.Null, Rule: true, Chain: nil}}},
					{Target: "parameters.ClusterBaseProperties.Subnet", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.Subnet.ID", Name: validation.Null, Rule: true, Chain: nil}}},
				}}}}}); err != nil {
		return result, validation.NewError("batchai.ClustersClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, workspaceName, clusterName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ClustersClient) CreatePreparer(ctx context.Context, resourceGroupName string, workspaceName string, clusterName string, parameters ClusterCreateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2018-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/clusters/{clusterName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ClustersClient) CreateSender(req *http.Request) (future ClustersCreateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ClustersClient) CreateResponder(resp *http.Response) (result Cluster, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a Cluster.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// workspaceName - the name of the workspace. Workspace names can only contain a combination of alphanumeric
// characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
// clusterName - the name of the cluster within the specified resource group. Cluster names can only contain a
// combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1
// through 64 characters long.
func (client ClustersClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, clusterName string) (result ClustersDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClustersClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `^[-\w_]+$`, Chain: nil}}},
		{TargetValue: clusterName,
			Constraints: []validation.Constraint{{Target: "clusterName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "clusterName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "clusterName", Name: validation.Pattern, Rule: `^[-\w_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batchai.ClustersClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, workspaceName, clusterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ClustersClient) DeletePreparer(ctx context.Context, resourceGroupName string, workspaceName string, clusterName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2018-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/clusters/{clusterName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ClustersClient) DeleteSender(req *http.Request) (future ClustersDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ClustersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets information about a Cluster.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// workspaceName - the name of the workspace. Workspace names can only contain a combination of alphanumeric
// characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
// clusterName - the name of the cluster within the specified resource group. Cluster names can only contain a
// combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1
// through 64 characters long.
func (client ClustersClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, clusterName string) (result Cluster, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClustersClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `^[-\w_]+$`, Chain: nil}}},
		{TargetValue: clusterName,
			Constraints: []validation.Constraint{{Target: "clusterName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "clusterName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "clusterName", Name: validation.Pattern, Rule: `^[-\w_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batchai.ClustersClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, workspaceName, clusterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ClustersClient) GetPreparer(ctx context.Context, resourceGroupName string, workspaceName string, clusterName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2018-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/clusters/{clusterName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ClustersClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ClustersClient) GetResponder(resp *http.Response) (result Cluster, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByWorkspace gets information about Clusters associated with the given Workspace.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// workspaceName - the name of the workspace. Workspace names can only contain a combination of alphanumeric
// characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
// maxResults - the maximum number of items to return in the response. A maximum of 1000 files can be returned.
func (client ClustersClient) ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string, maxResults *int32) (result ClusterListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClustersClient.ListByWorkspace")
		defer func() {
			sc := -1
			if result.clr.Response.Response != nil {
				sc = result.clr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `^[-\w_]+$`, Chain: nil}}},
		{TargetValue: maxResults,
			Constraints: []validation.Constraint{{Target: "maxResults", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "maxResults", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
					{Target: "maxResults", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("batchai.ClustersClient", "ListByWorkspace", err.Error())
	}

	result.fn = client.listByWorkspaceNextResults
	req, err := client.ListByWorkspacePreparer(ctx, resourceGroupName, workspaceName, maxResults)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "ListByWorkspace", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByWorkspaceSender(req)
	if err != nil {
		result.clr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "ListByWorkspace", resp, "Failure sending request")
		return
	}

	result.clr, err = client.ListByWorkspaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "ListByWorkspace", resp, "Failure responding to request")
		return
	}
	if result.clr.hasNextLink() && result.clr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByWorkspacePreparer prepares the ListByWorkspace request.
func (client ClustersClient) ListByWorkspacePreparer(ctx context.Context, resourceGroupName string, workspaceName string, maxResults *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2018-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if maxResults != nil {
		queryParameters["maxresults"] = autorest.Encode("query", *maxResults)
	} else {
		queryParameters["maxresults"] = autorest.Encode("query", 1000)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/clusters", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByWorkspaceSender sends the ListByWorkspace request. The method will close the
// http.Response Body if it receives an error.
func (client ClustersClient) ListByWorkspaceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByWorkspaceResponder handles the response to the ListByWorkspace request. The method always
// closes the http.Response Body.
func (client ClustersClient) ListByWorkspaceResponder(resp *http.Response) (result ClusterListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByWorkspaceNextResults retrieves the next set of results, if any.
func (client ClustersClient) listByWorkspaceNextResults(ctx context.Context, lastResults ClusterListResult) (result ClusterListResult, err error) {
	req, err := lastResults.clusterListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "batchai.ClustersClient", "listByWorkspaceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByWorkspaceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "batchai.ClustersClient", "listByWorkspaceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByWorkspaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "listByWorkspaceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByWorkspaceComplete enumerates all values, automatically crossing page boundaries as required.
func (client ClustersClient) ListByWorkspaceComplete(ctx context.Context, resourceGroupName string, workspaceName string, maxResults *int32) (result ClusterListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClustersClient.ListByWorkspace")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByWorkspace(ctx, resourceGroupName, workspaceName, maxResults)
	return
}

// ListRemoteLoginInformation get the IP address, port of all the compute nodes in the Cluster.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// workspaceName - the name of the workspace. Workspace names can only contain a combination of alphanumeric
// characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
// clusterName - the name of the cluster within the specified resource group. Cluster names can only contain a
// combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1
// through 64 characters long.
func (client ClustersClient) ListRemoteLoginInformation(ctx context.Context, resourceGroupName string, workspaceName string, clusterName string) (result RemoteLoginInformationListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClustersClient.ListRemoteLoginInformation")
		defer func() {
			sc := -1
			if result.rlilr.Response.Response != nil {
				sc = result.rlilr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `^[-\w_]+$`, Chain: nil}}},
		{TargetValue: clusterName,
			Constraints: []validation.Constraint{{Target: "clusterName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "clusterName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "clusterName", Name: validation.Pattern, Rule: `^[-\w_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batchai.ClustersClient", "ListRemoteLoginInformation", err.Error())
	}

	result.fn = client.listRemoteLoginInformationNextResults
	req, err := client.ListRemoteLoginInformationPreparer(ctx, resourceGroupName, workspaceName, clusterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "ListRemoteLoginInformation", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListRemoteLoginInformationSender(req)
	if err != nil {
		result.rlilr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "ListRemoteLoginInformation", resp, "Failure sending request")
		return
	}

	result.rlilr, err = client.ListRemoteLoginInformationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "ListRemoteLoginInformation", resp, "Failure responding to request")
		return
	}
	if result.rlilr.hasNextLink() && result.rlilr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListRemoteLoginInformationPreparer prepares the ListRemoteLoginInformation request.
func (client ClustersClient) ListRemoteLoginInformationPreparer(ctx context.Context, resourceGroupName string, workspaceName string, clusterName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2018-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/clusters/{clusterName}/listRemoteLoginInformation", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListRemoteLoginInformationSender sends the ListRemoteLoginInformation request. The method will close the
// http.Response Body if it receives an error.
func (client ClustersClient) ListRemoteLoginInformationSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListRemoteLoginInformationResponder handles the response to the ListRemoteLoginInformation request. The method always
// closes the http.Response Body.
func (client ClustersClient) ListRemoteLoginInformationResponder(resp *http.Response) (result RemoteLoginInformationListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listRemoteLoginInformationNextResults retrieves the next set of results, if any.
func (client ClustersClient) listRemoteLoginInformationNextResults(ctx context.Context, lastResults RemoteLoginInformationListResult) (result RemoteLoginInformationListResult, err error) {
	req, err := lastResults.remoteLoginInformationListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "batchai.ClustersClient", "listRemoteLoginInformationNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListRemoteLoginInformationSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "batchai.ClustersClient", "listRemoteLoginInformationNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListRemoteLoginInformationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "listRemoteLoginInformationNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListRemoteLoginInformationComplete enumerates all values, automatically crossing page boundaries as required.
func (client ClustersClient) ListRemoteLoginInformationComplete(ctx context.Context, resourceGroupName string, workspaceName string, clusterName string) (result RemoteLoginInformationListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClustersClient.ListRemoteLoginInformation")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListRemoteLoginInformation(ctx, resourceGroupName, workspaceName, clusterName)
	return
}

// Update updates properties of a Cluster.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// workspaceName - the name of the workspace. Workspace names can only contain a combination of alphanumeric
// characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
// clusterName - the name of the cluster within the specified resource group. Cluster names can only contain a
// combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1
// through 64 characters long.
// parameters - additional parameters for cluster update.
func (client ClustersClient) Update(ctx context.Context, resourceGroupName string, workspaceName string, clusterName string, parameters ClusterUpdateParameters) (result Cluster, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClustersClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `^[-\w_]+$`, Chain: nil}}},
		{TargetValue: clusterName,
			Constraints: []validation.Constraint{{Target: "clusterName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "clusterName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "clusterName", Name: validation.Pattern, Rule: `^[-\w_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batchai.ClustersClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, workspaceName, clusterName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ClustersClient) UpdatePreparer(ctx context.Context, resourceGroupName string, workspaceName string, clusterName string, parameters ClusterUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2018-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/clusters/{clusterName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ClustersClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ClustersClient) UpdateResponder(resp *http.Response) (result Cluster, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
