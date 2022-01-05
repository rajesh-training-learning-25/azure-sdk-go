//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkusto

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DatabasesClient contains the methods for the Databases group.
// Don't use this type directly, use NewDatabasesClient() instead.
type DatabasesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewDatabasesClient creates a new instance of DatabasesClient with the specified values.
func NewDatabasesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *DatabasesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &DatabasesClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// AddPrincipals - Add Database principals permissions.
// If the operation fails it returns the *CloudError error type.
func (client *DatabasesClient) AddPrincipals(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, databasePrincipalsToAdd DatabasePrincipalListRequest, options *DatabasesAddPrincipalsOptions) (DatabasesAddPrincipalsResponse, error) {
	req, err := client.addPrincipalsCreateRequest(ctx, resourceGroupName, clusterName, databaseName, databasePrincipalsToAdd, options)
	if err != nil {
		return DatabasesAddPrincipalsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DatabasesAddPrincipalsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DatabasesAddPrincipalsResponse{}, client.addPrincipalsHandleError(resp)
	}
	return client.addPrincipalsHandleResponse(resp)
}

// addPrincipalsCreateRequest creates the AddPrincipals request.
func (client *DatabasesClient) addPrincipalsCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, databasePrincipalsToAdd DatabasePrincipalListRequest, options *DatabasesAddPrincipalsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/addPrincipals"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, databasePrincipalsToAdd)
}

// addPrincipalsHandleResponse handles the AddPrincipals response.
func (client *DatabasesClient) addPrincipalsHandleResponse(resp *http.Response) (DatabasesAddPrincipalsResponse, error) {
	result := DatabasesAddPrincipalsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabasePrincipalListResult); err != nil {
		return DatabasesAddPrincipalsResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// addPrincipalsHandleError handles the AddPrincipals error response.
func (client *DatabasesClient) addPrincipalsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// CheckNameAvailability - Checks that the databases resource name is valid and is not already in use.
// If the operation fails it returns the *CloudError error type.
func (client *DatabasesClient) CheckNameAvailability(ctx context.Context, resourceGroupName string, clusterName string, resourceName CheckNameRequest, options *DatabasesCheckNameAvailabilityOptions) (DatabasesCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, resourceGroupName, clusterName, resourceName, options)
	if err != nil {
		return DatabasesCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DatabasesCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DatabasesCheckNameAvailabilityResponse{}, client.checkNameAvailabilityHandleError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *DatabasesClient) checkNameAvailabilityCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, resourceName CheckNameRequest, options *DatabasesCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/checkNameAvailability"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, resourceName)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *DatabasesClient) checkNameAvailabilityHandleResponse(resp *http.Response) (DatabasesCheckNameAvailabilityResponse, error) {
	result := DatabasesCheckNameAvailabilityResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameResult); err != nil {
		return DatabasesCheckNameAvailabilityResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// checkNameAvailabilityHandleError handles the CheckNameAvailability error response.
func (client *DatabasesClient) checkNameAvailabilityHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginCreateOrUpdate - Creates or updates a database.
// If the operation fails it returns the *CloudError error type.
func (client *DatabasesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters DatabaseClassification, options *DatabasesBeginCreateOrUpdateOptions) (DatabasesCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, clusterName, databaseName, parameters, options)
	if err != nil {
		return DatabasesCreateOrUpdatePollerResponse{}, err
	}
	result := DatabasesCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DatabasesClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return DatabasesCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &DatabasesCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a database.
// If the operation fails it returns the *CloudError error type.
func (client *DatabasesClient) createOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters DatabaseClassification, options *DatabasesBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, clusterName, databaseName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DatabasesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters DatabaseClassification, options *DatabasesBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *DatabasesClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes the database with the given name.
// If the operation fails it returns the *CloudError error type.
func (client *DatabasesClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, options *DatabasesBeginDeleteOptions) (DatabasesDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, clusterName, databaseName, options)
	if err != nil {
		return DatabasesDeletePollerResponse{}, err
	}
	result := DatabasesDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DatabasesClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return DatabasesDeletePollerResponse{}, err
	}
	result.Poller = &DatabasesDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the database with the given name.
// If the operation fails it returns the *CloudError error type.
func (client *DatabasesClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, options *DatabasesBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, databaseName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DatabasesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, options *DatabasesBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *DatabasesClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Returns a database.
// If the operation fails it returns the *CloudError error type.
func (client *DatabasesClient) Get(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, options *DatabasesGetOptions) (DatabasesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, databaseName, options)
	if err != nil {
		return DatabasesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DatabasesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DatabasesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DatabasesClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, options *DatabasesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DatabasesClient) getHandleResponse(resp *http.Response) (DatabasesGetResponse, error) {
	result := DatabasesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return DatabasesGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *DatabasesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByCluster - Returns the list of databases of the given Kusto cluster.
// If the operation fails it returns the *CloudError error type.
func (client *DatabasesClient) ListByCluster(ctx context.Context, resourceGroupName string, clusterName string, options *DatabasesListByClusterOptions) (DatabasesListByClusterResponse, error) {
	req, err := client.listByClusterCreateRequest(ctx, resourceGroupName, clusterName, options)
	if err != nil {
		return DatabasesListByClusterResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DatabasesListByClusterResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DatabasesListByClusterResponse{}, client.listByClusterHandleError(resp)
	}
	return client.listByClusterHandleResponse(resp)
}

// listByClusterCreateRequest creates the ListByCluster request.
func (client *DatabasesClient) listByClusterCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *DatabasesListByClusterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByClusterHandleResponse handles the ListByCluster response.
func (client *DatabasesClient) listByClusterHandleResponse(resp *http.Response) (DatabasesListByClusterResponse, error) {
	result := DatabasesListByClusterResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabaseListResult); err != nil {
		return DatabasesListByClusterResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByClusterHandleError handles the ListByCluster error response.
func (client *DatabasesClient) listByClusterHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListPrincipals - Returns a list of database principals of the given Kusto cluster and database.
// If the operation fails it returns the *CloudError error type.
func (client *DatabasesClient) ListPrincipals(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, options *DatabasesListPrincipalsOptions) (DatabasesListPrincipalsResponse, error) {
	req, err := client.listPrincipalsCreateRequest(ctx, resourceGroupName, clusterName, databaseName, options)
	if err != nil {
		return DatabasesListPrincipalsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DatabasesListPrincipalsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DatabasesListPrincipalsResponse{}, client.listPrincipalsHandleError(resp)
	}
	return client.listPrincipalsHandleResponse(resp)
}

// listPrincipalsCreateRequest creates the ListPrincipals request.
func (client *DatabasesClient) listPrincipalsCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, options *DatabasesListPrincipalsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/listPrincipals"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listPrincipalsHandleResponse handles the ListPrincipals response.
func (client *DatabasesClient) listPrincipalsHandleResponse(resp *http.Response) (DatabasesListPrincipalsResponse, error) {
	result := DatabasesListPrincipalsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabasePrincipalListResult); err != nil {
		return DatabasesListPrincipalsResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listPrincipalsHandleError handles the ListPrincipals error response.
func (client *DatabasesClient) listPrincipalsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// RemovePrincipals - Remove Database principals permissions.
// If the operation fails it returns the *CloudError error type.
func (client *DatabasesClient) RemovePrincipals(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, databasePrincipalsToRemove DatabasePrincipalListRequest, options *DatabasesRemovePrincipalsOptions) (DatabasesRemovePrincipalsResponse, error) {
	req, err := client.removePrincipalsCreateRequest(ctx, resourceGroupName, clusterName, databaseName, databasePrincipalsToRemove, options)
	if err != nil {
		return DatabasesRemovePrincipalsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DatabasesRemovePrincipalsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DatabasesRemovePrincipalsResponse{}, client.removePrincipalsHandleError(resp)
	}
	return client.removePrincipalsHandleResponse(resp)
}

// removePrincipalsCreateRequest creates the RemovePrincipals request.
func (client *DatabasesClient) removePrincipalsCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, databasePrincipalsToRemove DatabasePrincipalListRequest, options *DatabasesRemovePrincipalsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/removePrincipals"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, databasePrincipalsToRemove)
}

// removePrincipalsHandleResponse handles the RemovePrincipals response.
func (client *DatabasesClient) removePrincipalsHandleResponse(resp *http.Response) (DatabasesRemovePrincipalsResponse, error) {
	result := DatabasesRemovePrincipalsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabasePrincipalListResult); err != nil {
		return DatabasesRemovePrincipalsResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// removePrincipalsHandleError handles the RemovePrincipals error response.
func (client *DatabasesClient) removePrincipalsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginUpdate - Updates a database.
// If the operation fails it returns the *CloudError error type.
func (client *DatabasesClient) BeginUpdate(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters DatabaseClassification, options *DatabasesBeginUpdateOptions) (DatabasesUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, clusterName, databaseName, parameters, options)
	if err != nil {
		return DatabasesUpdatePollerResponse{}, err
	}
	result := DatabasesUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DatabasesClient.Update", "", resp, client.pl, client.updateHandleError)
	if err != nil {
		return DatabasesUpdatePollerResponse{}, err
	}
	result.Poller = &DatabasesUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates a database.
// If the operation fails it returns the *CloudError error type.
func (client *DatabasesClient) update(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters DatabaseClassification, options *DatabasesBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, clusterName, databaseName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *DatabasesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters DatabaseClassification, options *DatabasesBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleError handles the Update error response.
func (client *DatabasesClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
