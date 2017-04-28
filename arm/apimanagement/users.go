package apimanagement

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
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// UsersClient is the composite Swagger for ApiManagement Client
type UsersClient struct {
	ManagementClient
}

// NewUsersClient creates an instance of the UsersClient client.
func NewUsersClient(subscriptionID string) UsersClient {
	return NewUsersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewUsersClientWithBaseURI creates an instance of the UsersClient client.
func NewUsersClientWithBaseURI(baseURI string, subscriptionID string) UsersClient {
	return UsersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or Updates a user.
//
// resourceGroupName is the name of the resource group. serviceName is the name
// of the API Management service. uid is user identifier. Must be unique in the
// current API Management service instance. parameters is create or update
// parameters.
func (client UsersClient) CreateOrUpdate(resourceGroupName string, serviceName string, uid string, parameters UserCreateParameters) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: uid,
			Constraints: []validation.Constraint{{Target: "uid", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "uid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "uid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Email", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.Email", Name: validation.MaxLength, Rule: 254, Chain: nil},
					{Target: "parameters.Email", Name: validation.MinLength, Rule: 1, Chain: nil},
				}},
				{Target: "parameters.Password", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.FirstName", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "parameters.FirstName", Name: validation.MaxLength, Rule: 100, Chain: nil},
						{Target: "parameters.FirstName", Name: validation.MinLength, Rule: 1, Chain: nil},
					}},
				{Target: "parameters.LastName", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "parameters.LastName", Name: validation.MaxLength, Rule: 100, Chain: nil},
						{Target: "parameters.LastName", Name: validation.MinLength, Rule: 1, Chain: nil},
					}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "apimanagement.UsersClient", "CreateOrUpdate")
	}

	req, err := client.CreateOrUpdatePreparer(resourceGroupName, serviceName, uid, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client UsersClient) CreateOrUpdatePreparer(resourceGroupName string, serviceName string, uid string, parameters UserCreateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"uid":               autorest.Encode("path", uid),
	}

	const APIVersion = "2016-10-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{uid}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client UsersClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client UsersClient) CreateOrUpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete deletes specific user.
//
// resourceGroupName is the name of the resource group. serviceName is the name
// of the API Management service. uid is user identifier. Must be unique in the
// current API Management service instance. ifMatch is the entity state (Etag)
// version of the user to delete. A value of "*" can be used for If-Match to
// unconditionally apply the operation. deleteSubscriptions is whether to
// delete user's subscription or not.
func (client UsersClient) Delete(resourceGroupName string, serviceName string, uid string, ifMatch string, deleteSubscriptions *bool) (result ErrorBodyContract, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: uid,
			Constraints: []validation.Constraint{{Target: "uid", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "uid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "uid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "apimanagement.UsersClient", "Delete")
	}

	req, err := client.DeletePreparer(resourceGroupName, serviceName, uid, ifMatch, deleteSubscriptions)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client UsersClient) DeletePreparer(resourceGroupName string, serviceName string, uid string, ifMatch string, deleteSubscriptions *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"uid":               autorest.Encode("path", uid),
	}

	const APIVersion = "2016-10-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if deleteSubscriptions != nil {
		queryParameters["deleteSubscriptions"] = autorest.Encode("query", *deleteSubscriptions)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{uid}", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client UsersClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client UsersClient) DeleteResponder(resp *http.Response) (result ErrorBodyContract, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusMethodNotAllowed),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GenerateSsoURL retrieves a redirection URL containing an authentication
// token for signing a given user into the developer portal.
//
// resourceGroupName is the name of the resource group. serviceName is the name
// of the API Management service. uid is user identifier. Must be unique in the
// current API Management service instance.
func (client UsersClient) GenerateSsoURL(resourceGroupName string, serviceName string, uid string) (result GenerateSsoURLResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: uid,
			Constraints: []validation.Constraint{{Target: "uid", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "uid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "uid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "apimanagement.UsersClient", "GenerateSsoURL")
	}

	req, err := client.GenerateSsoURLPreparer(resourceGroupName, serviceName, uid)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "GenerateSsoURL", nil, "Failure preparing request")
		return
	}

	resp, err := client.GenerateSsoURLSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "GenerateSsoURL", resp, "Failure sending request")
		return
	}

	result, err = client.GenerateSsoURLResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "GenerateSsoURL", resp, "Failure responding to request")
	}

	return
}

// GenerateSsoURLPreparer prepares the GenerateSsoURL request.
func (client UsersClient) GenerateSsoURLPreparer(resourceGroupName string, serviceName string, uid string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"uid":               autorest.Encode("path", uid),
	}

	const APIVersion = "2016-10-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{uid}/generateSsoUrl", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GenerateSsoURLSender sends the GenerateSsoURL request. The method will close the
// http.Response Body if it receives an error.
func (client UsersClient) GenerateSsoURLSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GenerateSsoURLResponder handles the response to the GenerateSsoURL request. The method always
// closes the http.Response Body.
func (client UsersClient) GenerateSsoURLResponder(resp *http.Response) (result GenerateSsoURLResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets the details of the user specified by its identifier.
//
// resourceGroupName is the name of the resource group. serviceName is the name
// of the API Management service. uid is user identifier. Must be unique in the
// current API Management service instance.
func (client UsersClient) Get(resourceGroupName string, serviceName string, uid string) (result UserContract, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: uid,
			Constraints: []validation.Constraint{{Target: "uid", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "uid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "uid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "apimanagement.UsersClient", "Get")
	}

	req, err := client.GetPreparer(resourceGroupName, serviceName, uid)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client UsersClient) GetPreparer(resourceGroupName string, serviceName string, uid string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"uid":               autorest.Encode("path", uid),
	}

	const APIVersion = "2016-10-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{uid}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client UsersClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client UsersClient) GetResponder(resp *http.Response) (result UserContract, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByService lists a collection of registered users in the specified
// service instance.
//
// resourceGroupName is the name of the resource group. serviceName is the name
// of the API Management service. filter is | Field            | Supported
// operators    | Supported functions               |
// |------------------|------------------------|-----------------------------------|
// | id               | ge, le, eq, ne, gt, lt | substringof, contains,
// startswith, endswith |
// | firstName        | ge, le, eq, ne, gt, lt | substringof, contains,
// startswith, endswith |
// | lastName         | ge, le, eq, ne, gt, lt | substringof, contains,
// startswith, endswith |
// | email            | ge, le, eq, ne, gt, lt | substringof, contains,
// startswith, endswith |
// | state            | eq                     | N/A
// |
// | registrationDate | ge, le, eq, ne, gt, lt | N/A
// |
// | note             | ge, le, eq, ne, gt, lt | substringof, contains,
// startswith, endswith | top is number of records to return. skip is number of
// records to skip.
func (client UsersClient) ListByService(resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result UserCollection, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil}}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "apimanagement.UsersClient", "ListByService")
	}

	req, err := client.ListByServicePreparer(resourceGroupName, serviceName, filter, top, skip)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "ListByService", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "ListByService", resp, "Failure sending request")
		return
	}

	result, err = client.ListByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "ListByService", resp, "Failure responding to request")
	}

	return
}

// ListByServicePreparer prepares the ListByService request.
func (client UsersClient) ListByServicePreparer(resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-10-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByServiceSender sends the ListByService request. The method will close the
// http.Response Body if it receives an error.
func (client UsersClient) ListByServiceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByServiceResponder handles the response to the ListByService request. The method always
// closes the http.Response Body.
func (client UsersClient) ListByServiceResponder(resp *http.Response) (result UserCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByServiceNextResults retrieves the next set of results, if any.
func (client UsersClient) ListByServiceNextResults(lastResults UserCollection) (result UserCollection, err error) {
	req, err := lastResults.UserCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.UsersClient", "ListByService", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.UsersClient", "ListByService", resp, "Failure sending next results request")
	}

	result, err = client.ListByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "ListByService", resp, "Failure responding to next results request")
	}

	return
}

// Update updates the details of the user specified by its identifier.
//
// resourceGroupName is the name of the resource group. serviceName is the name
// of the API Management service. uid is user identifier. Must be unique in the
// current API Management service instance. parameters is update parameters.
// ifMatch is the entity state (Etag) version of the user to update. A value of
// "*" can be used for If-Match to unconditionally apply the operation.
func (client UsersClient) Update(resourceGroupName string, serviceName string, uid string, parameters UserUpdateParameters, ifMatch string) (result ErrorBodyContract, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: uid,
			Constraints: []validation.Constraint{{Target: "uid", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "uid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "uid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "apimanagement.UsersClient", "Update")
	}

	req, err := client.UpdatePreparer(resourceGroupName, serviceName, uid, parameters, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client UsersClient) UpdatePreparer(resourceGroupName string, serviceName string, uid string, parameters UserUpdateParameters, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"uid":               autorest.Encode("path", uid),
	}

	const APIVersion = "2016-10-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{uid}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	return preparer.Prepare(&http.Request{})
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client UsersClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client UsersClient) UpdateResponder(resp *http.Response) (result ErrorBodyContract, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusMethodNotAllowed),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
