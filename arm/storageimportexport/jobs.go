package storageimportexport

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
// Code generated by Microsoft (R) AutoRest Code Generator 2.2.18.0
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// JobsClient is the the Microsoft Azure Storage Import/Export Resource Provider API.
type JobsClient struct {
	ManagementClient
}

// NewJobsClient creates an instance of the JobsClient client.
func NewJobsClient(subscriptionID string) JobsClient {
	return NewJobsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewJobsClientWithBaseURI creates an instance of the JobsClient client.
func NewJobsClientWithBaseURI(baseURI string, subscriptionID string) JobsClient {
	return JobsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a new import/export job or updates an existing import/export job in the specified
// subscription.
//
// resourceGroupName is the resource group name uniquely identifies the resource group within the user subscription.
// jobName is the name of the import/export job. jobProperties is properties of the import/export job that need to be
// specified during creation.
func (client JobsClient) CreateOrUpdate(resourceGroupName string, jobName string, jobProperties Job) (result Job, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: jobProperties,
			Constraints: []validation.Constraint{{Target: "jobProperties.JobProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "jobProperties.JobProperties.StorageAccountID", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "jobProperties.JobProperties.ReturnAddress", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "jobProperties.JobProperties.ReturnAddress.RecipientName", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "jobProperties.JobProperties.ReturnAddress.StreetAddress1", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "jobProperties.JobProperties.ReturnAddress.City", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "jobProperties.JobProperties.ReturnAddress.PostalCode", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "jobProperties.JobProperties.ReturnAddress.CountryOrRegion", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "jobProperties.JobProperties.ReturnAddress.Phone", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "jobProperties.JobProperties.ReturnAddress.Email", Name: validation.Null, Rule: true, Chain: nil},
						}},
					{Target: "jobProperties.JobProperties.ReturnShipping", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "jobProperties.JobProperties.ReturnShipping.CarrierName", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "jobProperties.JobProperties.ReturnShipping.CarrierAccountNumber", Name: validation.Null, Rule: true, Chain: nil},
						}},
					{Target: "jobProperties.JobProperties.ShippingInformation", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "jobProperties.JobProperties.ShippingInformation.Name", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "jobProperties.JobProperties.ShippingInformation.Address", Name: validation.Null, Rule: true, Chain: nil},
						}},
					{Target: "jobProperties.JobProperties.DeliveryPackage", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "jobProperties.JobProperties.DeliveryPackage.CarrierName", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "jobProperties.JobProperties.DeliveryPackage.TrackingNumber", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "jobProperties.JobProperties.DeliveryPackage.DriveCount", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "jobProperties.JobProperties.DeliveryPackage.ShipDate", Name: validation.Null, Rule: true, Chain: nil},
						}},
					{Target: "jobProperties.JobProperties.ReturnPackage", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "jobProperties.JobProperties.ReturnPackage.CarrierName", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "jobProperties.JobProperties.ReturnPackage.TrackingNumber", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "jobProperties.JobProperties.ReturnPackage.DriveCount", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "jobProperties.JobProperties.ReturnPackage.ShipDate", Name: validation.Null, Rule: true, Chain: nil},
						}},
					{Target: "jobProperties.JobProperties.DiagnosticsPath", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "jobProperties.JobProperties.DriveList", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "jobProperties.JobProperties.DriveList", Name: validation.MaxItems, Rule: 10, Chain: nil},
							{Target: "jobProperties.JobProperties.DriveList", Name: validation.MinItems, Rule: 0, Chain: nil},
						}},
				}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "storageimportexport.JobsClient", "CreateOrUpdate")
	}

	req, err := client.CreateOrUpdatePreparer(resourceGroupName, jobName, jobProperties)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client JobsClient) CreateOrUpdatePreparer(resourceGroupName string, jobName string, jobProperties Job) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ImportExport/jobs/{jobName}", pathParameters),
		autorest.WithJSON(jobProperties),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("Accept-Language", client.AcceptLanguage))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client JobsClient) CreateOrUpdateResponder(resp *http.Response) (result Job, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes an existing import/export job. Only import/export jobs in the Creating or Completed states can be
// deleted.
//
// resourceGroupName is the resource group name uniquely identifies the resource group within the user subscription.
// jobName is the name of the import/export job.
func (client JobsClient) Delete(resourceGroupName string, jobName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(resourceGroupName, jobName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client JobsClient) DeletePreparer(resourceGroupName string, jobName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ImportExport/jobs/{jobName}", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("Accept-Language", client.AcceptLanguage))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client JobsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets information about an existing import/export job.
//
// resourceGroupName is the resource group name uniquely identifies the resource group within the user subscription.
// jobName is the name of the import/export job.
func (client JobsClient) Get(resourceGroupName string, jobName string) (result Job, err error) {
	req, err := client.GetPreparer(resourceGroupName, jobName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client JobsClient) GetPreparer(resourceGroupName string, jobName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ImportExport/jobs/{jobName}", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("Accept-Language", client.AcceptLanguage))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client JobsClient) GetResponder(resp *http.Response) (result Job, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all the active and completed import/export jobs in a subscription.
//
// top is an integer value that specifies how many jobs at most should be returned. The value cannot exceed 100. filter
// is can be used to restrict the results to certain conditions. The following possible values can be used with
// $filter: 1) $filter=type eq '{type}'; 2) $filter=trackingnumber eq '{trackingnumber}'; 3) $filter=state eq
// '{state}'; 4) Logical and combination of the above, for example: $filter=type eq 'Import' and state eq
// 'Transferring'. Valid values for type are Import and Export. Valid values for state are Creating, Shipping,
// Received, Transferring, Packaging, Closed, and Completed.
func (client JobsClient) List(top *int32, filter string) (result JobListResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMaximum, Rule: 100, Chain: nil},
					{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "storageimportexport.JobsClient", "List")
	}

	req, err := client.ListPreparer(top, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client JobsClient) ListPreparer(top *int32, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.ImportExport/jobs", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("Accept-Language", client.AcceptLanguage))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client JobsClient) ListResponder(resp *http.Response) (result JobListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client JobsClient) ListNextResults(lastResults JobListResult) (result JobListResult, err error) {
	req, err := lastResults.JobListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListComplete gets all elements from the list without paging.
func (client JobsClient) ListComplete(top *int32, filter string, cancel <-chan struct{}) (<-chan Job, <-chan error) {
	resultChan := make(chan Job)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.List(top, filter)
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
			list, err = client.ListNextResults(list)
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

// ListBitLockerKeys lists the BitLocker keys for all drives in the specified import/export job.
//
// resourceGroupName is the resource group name uniquely identifies the resource group within the user subscription.
// jobName is the name of the import/export job.
func (client JobsClient) ListBitLockerKeys(resourceGroupName string, jobName string) (result BitLockerKeysListResult, err error) {
	req, err := client.ListBitLockerKeysPreparer(resourceGroupName, jobName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "ListBitLockerKeys", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBitLockerKeysSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "ListBitLockerKeys", resp, "Failure sending request")
		return
	}

	result, err = client.ListBitLockerKeysResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "ListBitLockerKeys", resp, "Failure responding to request")
	}

	return
}

// ListBitLockerKeysPreparer prepares the ListBitLockerKeys request.
func (client JobsClient) ListBitLockerKeysPreparer(resourceGroupName string, jobName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ImportExport/jobs/{jobName}/listBitLockerKeys", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("Accept-Language", client.AcceptLanguage))
	return preparer.Prepare(&http.Request{})
}

// ListBitLockerKeysSender sends the ListBitLockerKeys request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) ListBitLockerKeysSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListBitLockerKeysResponder handles the response to the ListBitLockerKeys request. The method always
// closes the http.Response Body.
func (client JobsClient) ListBitLockerKeysResponder(resp *http.Response) (result BitLockerKeysListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup returns all active and completed import/export jobs in a resource group.
//
// resourceGroupName is the resource group name uniquely identifies the resource group within the user subscription.
// top is an integer value that specifies how many jobs at most should be returned. The value cannot exceed 100. filter
// is can be used to restrict the results to certain conditions. The following possible values can be used with
// $filter: 1) $filter=type eq '{type}'; 2) $filter=trackingnumber eq '{trackingnumber}'; 3) $filter=state eq
// '{state}'; 4) Logical and combination of the above, for example: $filter=type eq 'Import' and state eq
// 'Transferring'. Valid values for type are Import and Export. Valid values for state are Creating, Shipping,
// Received, Transferring, Packaging, Closed, and Completed.
func (client JobsClient) ListByResourceGroup(resourceGroupName string, top *int32, filter string) (result JobListResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMaximum, Rule: 100, Chain: nil},
					{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "storageimportexport.JobsClient", "ListByResourceGroup")
	}

	req, err := client.ListByResourceGroupPreparer(resourceGroupName, top, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client JobsClient) ListByResourceGroupPreparer(resourceGroupName string, top *int32, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ImportExport/jobs", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("Accept-Language", client.AcceptLanguage))
	return preparer.Prepare(&http.Request{})
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client JobsClient) ListByResourceGroupResponder(resp *http.Response) (result JobListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroupNextResults retrieves the next set of results, if any.
func (client JobsClient) ListByResourceGroupNextResults(lastResults JobListResult) (result JobListResult, err error) {
	req, err := lastResults.JobListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "ListByResourceGroup", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "ListByResourceGroup", resp, "Failure sending next results request")
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "ListByResourceGroup", resp, "Failure responding to next results request")
	}

	return
}

// ListByResourceGroupComplete gets all elements from the list without paging.
func (client JobsClient) ListByResourceGroupComplete(resourceGroupName string, top *int32, filter string, cancel <-chan struct{}) (<-chan Job, <-chan error) {
	resultChan := make(chan Job)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListByResourceGroup(resourceGroupName, top, filter)
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
			list, err = client.ListByResourceGroupNextResults(list)
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

// Move moves the specified import/export jobs from the resource group to a target resource group. The target resource
// group may be in a different subscription.
//
// resourceGroupName is the resource group name uniquely identifies the resource group within the user subscription.
// moveJobsParameters is parameters to be provided to move a job from one resource group to another.
func (client JobsClient) Move(resourceGroupName string, moveJobsParameters MoveJobParameters) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: moveJobsParameters,
			Constraints: []validation.Constraint{{Target: "moveJobsParameters.TargetResourceGroup", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "moveJobsParameters.Resources", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "storageimportexport.JobsClient", "Move")
	}

	req, err := client.MovePreparer(resourceGroupName, moveJobsParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "Move", nil, "Failure preparing request")
		return
	}

	resp, err := client.MoveSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "Move", resp, "Failure sending request")
		return
	}

	result, err = client.MoveResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "Move", resp, "Failure responding to request")
	}

	return
}

// MovePreparer prepares the Move request.
func (client JobsClient) MovePreparer(resourceGroupName string, moveJobsParameters MoveJobParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ImportExport/moveResources", pathParameters),
		autorest.WithJSON(moveJobsParameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("Accept-Language", client.AcceptLanguage))
	return preparer.Prepare(&http.Request{})
}

// MoveSender sends the Move request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) MoveSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// MoveResponder handles the response to the Move request. The method always
// closes the http.Response Body.
func (client JobsClient) MoveResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update updates specific properties of the import/export job. You can call this operation to notify the Import/Export
// service that the hard drives comprising the import or export job have been shipped to the Microsoft data center. It
// can also be used to cancel an existing job.
//
// resourceGroupName is the resource group name uniquely identifies the resource group within the user subscription.
// jobName is the name of the import/export job. jobProperties is import/export job properties that need to be updated.
func (client JobsClient) Update(resourceGroupName string, jobName string, jobProperties MutableJob) (result Job, err error) {
	req, err := client.UpdatePreparer(resourceGroupName, jobName, jobProperties)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client JobsClient) UpdatePreparer(resourceGroupName string, jobName string, jobProperties MutableJob) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ImportExport/jobs/{jobName}", pathParameters),
		autorest.WithJSON(jobProperties),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("Accept-Language", client.AcceptLanguage))
	return preparer.Prepare(&http.Request{})
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client JobsClient) UpdateResponder(resp *http.Response) (result Job, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
