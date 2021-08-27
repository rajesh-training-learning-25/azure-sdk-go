package guestconfiguration

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

// AssignmentReportsClient is the guest Configuration Client
type AssignmentReportsClient struct {
	BaseClient
}

// NewAssignmentReportsClient creates an instance of the AssignmentReportsClient client.
func NewAssignmentReportsClient(subscriptionID string) AssignmentReportsClient {
	return NewAssignmentReportsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAssignmentReportsClientWithBaseURI creates an instance of the AssignmentReportsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewAssignmentReportsClientWithBaseURI(baseURI string, subscriptionID string) AssignmentReportsClient {
	return AssignmentReportsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get a report for the guest configuration assignment, by reportId.
// Parameters:
// resourceGroupName - the resource group name.
// guestConfigurationAssignmentName - the guest configuration assignment name.
// reportID - the GUID for the guest configuration assignment report.
// VMName - the name of the virtual machine.
func (client AssignmentReportsClient) Get(ctx context.Context, resourceGroupName string, guestConfigurationAssignmentName string, reportID string, VMName string) (result AssignmentReportType, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssignmentReportsClient.Get")
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
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("guestconfiguration.AssignmentReportsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, guestConfigurationAssignmentName, reportID, VMName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "guestconfiguration.AssignmentReportsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "guestconfiguration.AssignmentReportsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "guestconfiguration.AssignmentReportsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client AssignmentReportsClient) GetPreparer(ctx context.Context, resourceGroupName string, guestConfigurationAssignmentName string, reportID string, VMName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"guestConfigurationAssignmentName": autorest.Encode("path", guestConfigurationAssignmentName),
		"reportId":                         autorest.Encode("path", reportID),
		"resourceGroupName":                autorest.Encode("path", resourceGroupName),
		"subscriptionId":                   autorest.Encode("path", client.SubscriptionID),
		"vmName":                           autorest.Encode("path", VMName),
	}

	const APIVersion = "2021-01-25"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{guestConfigurationAssignmentName}/reports/{reportId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AssignmentReportsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AssignmentReportsClient) GetResponder(resp *http.Response) (result AssignmentReportType, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list all reports for the guest configuration assignment, latest report first.
// Parameters:
// resourceGroupName - the resource group name.
// guestConfigurationAssignmentName - the guest configuration assignment name.
// VMName - the name of the virtual machine.
func (client AssignmentReportsClient) List(ctx context.Context, resourceGroupName string, guestConfigurationAssignmentName string, VMName string) (result AssignmentReportList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssignmentReportsClient.List")
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
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("guestconfiguration.AssignmentReportsClient", "List", err.Error())
	}

	req, err := client.ListPreparer(ctx, resourceGroupName, guestConfigurationAssignmentName, VMName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "guestconfiguration.AssignmentReportsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "guestconfiguration.AssignmentReportsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "guestconfiguration.AssignmentReportsClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client AssignmentReportsClient) ListPreparer(ctx context.Context, resourceGroupName string, guestConfigurationAssignmentName string, VMName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"guestConfigurationAssignmentName": autorest.Encode("path", guestConfigurationAssignmentName),
		"resourceGroupName":                autorest.Encode("path", resourceGroupName),
		"subscriptionId":                   autorest.Encode("path", client.SubscriptionID),
		"vmName":                           autorest.Encode("path", VMName),
	}

	const APIVersion = "2021-01-25"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{guestConfigurationAssignmentName}/reports", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client AssignmentReportsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client AssignmentReportsClient) ListResponder(resp *http.Response) (result AssignmentReportList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
