package dns

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

// RecordSetsClient is the the DNS Management Client.
type RecordSetsClient struct {
	ManagementClient
}

// NewRecordSetsClient creates an instance of the RecordSetsClient client.
func NewRecordSetsClient(subscriptionID string) RecordSetsClient {
	return NewRecordSetsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRecordSetsClientWithBaseURI creates an instance of the RecordSetsClient client.
func NewRecordSetsClientWithBaseURI(baseURI string, subscriptionID string) RecordSetsClient {
	return RecordSetsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// NewRecordSetsClientWithAuthFile creates an instance of the RecordSetsClient client.
func NewRecordSetsClientWithAuthFile() (RecordSetsClient, error) {
	c, err := NewWithAuthFile()
	return RecordSetsClient{c}, err
}

// CreateOrUpdate creates or updates a record set within a DNS zone.
//
// resourceGroupName is the name of the resource group. zoneName is the name of the DNS zone (without a terminating
// dot). relativeRecordSetName is the name of the record set, relative to the name of the zone. recordType is the type
// of DNS record in this record set. Record sets of type SOA can be updated but not created (they are created when the
// DNS zone is created). parameters is parameters supplied to the CreateOrUpdate operation. ifMatch is the etag of the
// record set. Omit this value to always overwrite the current record set. Specify the last-seen etag value to prevent
// accidentally overwritting any concurrent changes. ifNoneMatch is set to '*' to allow a new record set to be created,
// but to prevent updating an existing record set. Other values will be ignored.
func (client RecordSetsClient) CreateOrUpdate(resourceGroupName string, zoneName string, relativeRecordSetName string, recordType RecordType, parameters RecordSet, ifMatch string, ifNoneMatch string) (result RecordSet, err error) {
	req, err := client.CreateOrUpdatePreparer(resourceGroupName, zoneName, relativeRecordSetName, recordType, parameters, ifMatch, ifNoneMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client RecordSetsClient) CreateOrUpdatePreparer(resourceGroupName string, zoneName string, relativeRecordSetName string, recordType RecordType, parameters RecordSet, ifMatch string, ifNoneMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"recordType":            autorest.Encode("path", recordType),
		"relativeRecordSetName": relativeRecordSetName,
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
		"zoneName":              autorest.Encode("path", zoneName),
	}

	const APIVersion = "2016-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/{recordType}/{relativeRecordSetName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	if len(ifNoneMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-None-Match", autorest.String(ifNoneMatch)))
	}
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client RecordSetsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client RecordSetsClient) CreateOrUpdateResponder(resp *http.Response) (result RecordSet, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a record set from a DNS zone. This operation cannot be undone.
//
// resourceGroupName is the name of the resource group. zoneName is the name of the DNS zone (without a terminating
// dot). relativeRecordSetName is the name of the record set, relative to the name of the zone. recordType is the type
// of DNS record in this record set. Record sets of type SOA cannot be deleted (they are deleted when the DNS zone is
// deleted). ifMatch is the etag of the record set. Omit this value to always delete the current record set. Specify
// the last-seen etag value to prevent accidentally deleting any concurrent changes.
func (client RecordSetsClient) Delete(resourceGroupName string, zoneName string, relativeRecordSetName string, recordType RecordType, ifMatch string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(resourceGroupName, zoneName, relativeRecordSetName, recordType, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client RecordSetsClient) DeletePreparer(resourceGroupName string, zoneName string, relativeRecordSetName string, recordType RecordType, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"recordType":            autorest.Encode("path", recordType),
		"relativeRecordSetName": relativeRecordSetName,
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
		"zoneName":              autorest.Encode("path", zoneName),
	}

	const APIVersion = "2016-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/{recordType}/{relativeRecordSetName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client RecordSetsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client RecordSetsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a record set.
//
// resourceGroupName is the name of the resource group. zoneName is the name of the DNS zone (without a terminating
// dot). relativeRecordSetName is the name of the record set, relative to the name of the zone. recordType is the type
// of DNS record in this record set.
func (client RecordSetsClient) Get(resourceGroupName string, zoneName string, relativeRecordSetName string, recordType RecordType) (result RecordSet, err error) {
	req, err := client.GetPreparer(resourceGroupName, zoneName, relativeRecordSetName, recordType)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client RecordSetsClient) GetPreparer(resourceGroupName string, zoneName string, relativeRecordSetName string, recordType RecordType) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"recordType":            autorest.Encode("path", recordType),
		"relativeRecordSetName": relativeRecordSetName,
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
		"zoneName":              autorest.Encode("path", zoneName),
	}

	const APIVersion = "2016-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/{recordType}/{relativeRecordSetName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RecordSetsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RecordSetsClient) GetResponder(resp *http.Response) (result RecordSet, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByDNSZone lists all record sets in a DNS zone.
//
// resourceGroupName is the name of the resource group. zoneName is the name of the DNS zone (without a terminating
// dot). top is the maximum number of record sets to return. If not specified, returns up to 100 record sets.
// recordsetnamesuffix is the suffix label of the record set name that has to be used to filter the record set
// enumerations. If this parameter is specified, Enumeration will return only records that end with
// .<recordSetNameSuffix>
func (client RecordSetsClient) ListByDNSZone(resourceGroupName string, zoneName string, top *int32, recordsetnamesuffix string) (result RecordSetListResult, err error) {
	req, err := client.ListByDNSZonePreparer(resourceGroupName, zoneName, top, recordsetnamesuffix)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "ListByDNSZone", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByDNSZoneSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "ListByDNSZone", resp, "Failure sending request")
		return
	}

	result, err = client.ListByDNSZoneResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "ListByDNSZone", resp, "Failure responding to request")
	}

	return
}

// ListByDNSZonePreparer prepares the ListByDNSZone request.
func (client RecordSetsClient) ListByDNSZonePreparer(resourceGroupName string, zoneName string, top *int32, recordsetnamesuffix string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"zoneName":          autorest.Encode("path", zoneName),
	}

	const APIVersion = "2016-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(recordsetnamesuffix) > 0 {
		queryParameters["$recordsetnamesuffix"] = autorest.Encode("query", recordsetnamesuffix)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/recordsets", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByDNSZoneSender sends the ListByDNSZone request. The method will close the
// http.Response Body if it receives an error.
func (client RecordSetsClient) ListByDNSZoneSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByDNSZoneResponder handles the response to the ListByDNSZone request. The method always
// closes the http.Response Body.
func (client RecordSetsClient) ListByDNSZoneResponder(resp *http.Response) (result RecordSetListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByDNSZoneNextResults retrieves the next set of results, if any.
func (client RecordSetsClient) ListByDNSZoneNextResults(lastResults RecordSetListResult) (result RecordSetListResult, err error) {
	req, err := lastResults.RecordSetListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "dns.RecordSetsClient", "ListByDNSZone", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByDNSZoneSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "dns.RecordSetsClient", "ListByDNSZone", resp, "Failure sending next results request")
	}

	result, err = client.ListByDNSZoneResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "ListByDNSZone", resp, "Failure responding to next results request")
	}

	return
}

// ListByDNSZoneComplete gets all elements from the list without paging.
func (client RecordSetsClient) ListByDNSZoneComplete(resourceGroupName string, zoneName string, top *int32, recordsetnamesuffix string, cancel <-chan struct{}) (<-chan RecordSet, <-chan error) {
	resultChan := make(chan RecordSet)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListByDNSZone(resourceGroupName, zoneName, top, recordsetnamesuffix)
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
			list, err = client.ListByDNSZoneNextResults(list)
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

// ListByType lists the record sets of a specified type in a DNS zone.
//
// resourceGroupName is the name of the resource group. zoneName is the name of the DNS zone (without a terminating
// dot). recordType is the type of record sets to enumerate. top is the maximum number of record sets to return. If not
// specified, returns up to 100 record sets. recordsetnamesuffix is the suffix label of the record set name that has to
// be used to filter the record set enumerations. If this parameter is specified, Enumeration will return only records
// that end with .<recordSetNameSuffix>
func (client RecordSetsClient) ListByType(resourceGroupName string, zoneName string, recordType RecordType, top *int32, recordsetnamesuffix string) (result RecordSetListResult, err error) {
	req, err := client.ListByTypePreparer(resourceGroupName, zoneName, recordType, top, recordsetnamesuffix)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "ListByType", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByTypeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "ListByType", resp, "Failure sending request")
		return
	}

	result, err = client.ListByTypeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "ListByType", resp, "Failure responding to request")
	}

	return
}

// ListByTypePreparer prepares the ListByType request.
func (client RecordSetsClient) ListByTypePreparer(resourceGroupName string, zoneName string, recordType RecordType, top *int32, recordsetnamesuffix string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"recordType":        autorest.Encode("path", recordType),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"zoneName":          autorest.Encode("path", zoneName),
	}

	const APIVersion = "2016-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(recordsetnamesuffix) > 0 {
		queryParameters["$recordsetnamesuffix"] = autorest.Encode("query", recordsetnamesuffix)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/{recordType}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByTypeSender sends the ListByType request. The method will close the
// http.Response Body if it receives an error.
func (client RecordSetsClient) ListByTypeSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByTypeResponder handles the response to the ListByType request. The method always
// closes the http.Response Body.
func (client RecordSetsClient) ListByTypeResponder(resp *http.Response) (result RecordSetListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByTypeNextResults retrieves the next set of results, if any.
func (client RecordSetsClient) ListByTypeNextResults(lastResults RecordSetListResult) (result RecordSetListResult, err error) {
	req, err := lastResults.RecordSetListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "dns.RecordSetsClient", "ListByType", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByTypeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "dns.RecordSetsClient", "ListByType", resp, "Failure sending next results request")
	}

	result, err = client.ListByTypeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "ListByType", resp, "Failure responding to next results request")
	}

	return
}

// ListByTypeComplete gets all elements from the list without paging.
func (client RecordSetsClient) ListByTypeComplete(resourceGroupName string, zoneName string, recordType RecordType, top *int32, recordsetnamesuffix string, cancel <-chan struct{}) (<-chan RecordSet, <-chan error) {
	resultChan := make(chan RecordSet)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListByType(resourceGroupName, zoneName, recordType, top, recordsetnamesuffix)
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
			list, err = client.ListByTypeNextResults(list)
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

// Update updates a record set within a DNS zone.
//
// resourceGroupName is the name of the resource group. zoneName is the name of the DNS zone (without a terminating
// dot). relativeRecordSetName is the name of the record set, relative to the name of the zone. recordType is the type
// of DNS record in this record set. parameters is parameters supplied to the Update operation. ifMatch is the etag of
// the record set. Omit this value to always overwrite the current record set. Specify the last-seen etag value to
// prevent accidentally overwritting concurrent changes.
func (client RecordSetsClient) Update(resourceGroupName string, zoneName string, relativeRecordSetName string, recordType RecordType, parameters RecordSet, ifMatch string) (result RecordSet, err error) {
	req, err := client.UpdatePreparer(resourceGroupName, zoneName, relativeRecordSetName, recordType, parameters, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client RecordSetsClient) UpdatePreparer(resourceGroupName string, zoneName string, relativeRecordSetName string, recordType RecordType, parameters RecordSet, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"recordType":            autorest.Encode("path", recordType),
		"relativeRecordSetName": relativeRecordSetName,
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
		"zoneName":              autorest.Encode("path", zoneName),
	}

	const APIVersion = "2016-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/{recordType}/{relativeRecordSetName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare(&http.Request{})
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client RecordSetsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client RecordSetsClient) UpdateResponder(resp *http.Response) (result RecordSet, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
