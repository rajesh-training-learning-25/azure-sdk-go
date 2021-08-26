package databricks

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

// VNetPeeringClient is the the Microsoft Azure management APIs allow end users to operate on Azure Databricks
// Workspace resources.
type VNetPeeringClient struct {
	BaseClient
}

// NewVNetPeeringClient creates an instance of the VNetPeeringClient client.
func NewVNetPeeringClient(subscriptionID string) VNetPeeringClient {
	return NewVNetPeeringClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVNetPeeringClientWithBaseURI creates an instance of the VNetPeeringClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewVNetPeeringClientWithBaseURI(baseURI string, subscriptionID string) VNetPeeringClient {
	return VNetPeeringClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates vNet Peering for workspace.
// Parameters:
// virtualNetworkPeeringParameters - parameters supplied to the create workspace vNet Peering.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
// peeringName - the name of the workspace vNet peering.
func (client VNetPeeringClient) CreateOrUpdate(ctx context.Context, virtualNetworkPeeringParameters VirtualNetworkPeering, resourceGroupName string, workspaceName string, peeringName string) (result VNetPeeringCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VNetPeeringClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: virtualNetworkPeeringParameters,
			Constraints: []validation.Constraint{{Target: "virtualNetworkPeeringParameters.VirtualNetworkPeeringPropertiesFormat", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "virtualNetworkPeeringParameters.VirtualNetworkPeeringPropertiesFormat.RemoteVirtualNetwork", Name: validation.Null, Rule: true, Chain: nil}}}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("databricks.VNetPeeringClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, virtualNetworkPeeringParameters, resourceGroupName, workspaceName, peeringName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databricks.VNetPeeringClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databricks.VNetPeeringClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VNetPeeringClient) CreateOrUpdatePreparer(ctx context.Context, virtualNetworkPeeringParameters VirtualNetworkPeering, resourceGroupName string, workspaceName string, peeringName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"peeringName":       autorest.Encode("path", peeringName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2018-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	virtualNetworkPeeringParameters.Name = nil
	virtualNetworkPeeringParameters.ID = nil
	virtualNetworkPeeringParameters.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Databricks/workspaces/{workspaceName}/virtualNetworkPeerings/{peeringName}", pathParameters),
		autorest.WithJSON(virtualNetworkPeeringParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client VNetPeeringClient) CreateOrUpdateSender(req *http.Request) (future VNetPeeringCreateOrUpdateFuture, err error) {
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

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client VNetPeeringClient) CreateOrUpdateResponder(resp *http.Response) (result VirtualNetworkPeering, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the workspace vNetPeering.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
// peeringName - the name of the workspace vNet peering.
func (client VNetPeeringClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, peeringName string) (result VNetPeeringDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VNetPeeringClient.Delete")
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
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("databricks.VNetPeeringClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, workspaceName, peeringName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databricks.VNetPeeringClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databricks.VNetPeeringClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client VNetPeeringClient) DeletePreparer(ctx context.Context, resourceGroupName string, workspaceName string, peeringName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"peeringName":       autorest.Encode("path", peeringName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2018-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Databricks/workspaces/{workspaceName}/virtualNetworkPeerings/{peeringName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VNetPeeringClient) DeleteSender(req *http.Request) (future VNetPeeringDeleteFuture, err error) {
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
func (client VNetPeeringClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the workspace vNet Peering.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
// peeringName - the name of the workspace vNet peering.
func (client VNetPeeringClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, peeringName string) (result VirtualNetworkPeering, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VNetPeeringClient.Get")
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
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("databricks.VNetPeeringClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, workspaceName, peeringName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databricks.VNetPeeringClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "databricks.VNetPeeringClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databricks.VNetPeeringClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client VNetPeeringClient) GetPreparer(ctx context.Context, resourceGroupName string, workspaceName string, peeringName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"peeringName":       autorest.Encode("path", peeringName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2018-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Databricks/workspaces/{workspaceName}/virtualNetworkPeerings/{peeringName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VNetPeeringClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VNetPeeringClient) GetResponder(resp *http.Response) (result VirtualNetworkPeering, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByWorkspace lists the workspace vNet Peerings.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
func (client VNetPeeringClient) ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string) (result VirtualNetworkPeeringListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VNetPeeringClient.ListByWorkspace")
		defer func() {
			sc := -1
			if result.vnpl.Response.Response != nil {
				sc = result.vnpl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("databricks.VNetPeeringClient", "ListByWorkspace", err.Error())
	}

	result.fn = client.listByWorkspaceNextResults
	req, err := client.ListByWorkspacePreparer(ctx, resourceGroupName, workspaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databricks.VNetPeeringClient", "ListByWorkspace", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByWorkspaceSender(req)
	if err != nil {
		result.vnpl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "databricks.VNetPeeringClient", "ListByWorkspace", resp, "Failure sending request")
		return
	}

	result.vnpl, err = client.ListByWorkspaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databricks.VNetPeeringClient", "ListByWorkspace", resp, "Failure responding to request")
		return
	}
	if result.vnpl.hasNextLink() && result.vnpl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByWorkspacePreparer prepares the ListByWorkspace request.
func (client VNetPeeringClient) ListByWorkspacePreparer(ctx context.Context, resourceGroupName string, workspaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2018-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Databricks/workspaces/{workspaceName}/virtualNetworkPeerings", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByWorkspaceSender sends the ListByWorkspace request. The method will close the
// http.Response Body if it receives an error.
func (client VNetPeeringClient) ListByWorkspaceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByWorkspaceResponder handles the response to the ListByWorkspace request. The method always
// closes the http.Response Body.
func (client VNetPeeringClient) ListByWorkspaceResponder(resp *http.Response) (result VirtualNetworkPeeringList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByWorkspaceNextResults retrieves the next set of results, if any.
func (client VNetPeeringClient) listByWorkspaceNextResults(ctx context.Context, lastResults VirtualNetworkPeeringList) (result VirtualNetworkPeeringList, err error) {
	req, err := lastResults.virtualNetworkPeeringListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "databricks.VNetPeeringClient", "listByWorkspaceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByWorkspaceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "databricks.VNetPeeringClient", "listByWorkspaceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByWorkspaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databricks.VNetPeeringClient", "listByWorkspaceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByWorkspaceComplete enumerates all values, automatically crossing page boundaries as required.
func (client VNetPeeringClient) ListByWorkspaceComplete(ctx context.Context, resourceGroupName string, workspaceName string) (result VirtualNetworkPeeringListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VNetPeeringClient.ListByWorkspace")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByWorkspace(ctx, resourceGroupName, workspaceName)
	return
}
