// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpostgresql

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ConfigurationsClient contains the methods for the Configurations group.
// Don't use this type directly, use NewConfigurationsClient() instead.
type ConfigurationsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewConfigurationsClient creates a new instance of ConfigurationsClient with the specified values.
func NewConfigurationsClient(con *armcore.Connection, subscriptionID string) *ConfigurationsClient {
	return &ConfigurationsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Updates a configuration of a server.
// If the operation fails it returns the *CloudError error type.
func (client *ConfigurationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, configurationName string, parameters Configuration, options *ConfigurationsBeginCreateOrUpdateOptions) (ConfigurationsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, configurationName, parameters, options)
	if err != nil {
		return ConfigurationsCreateOrUpdatePollerResponse{}, err
	}
	result := ConfigurationsCreateOrUpdatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ConfigurationsClient.CreateOrUpdate", "", resp, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return ConfigurationsCreateOrUpdatePollerResponse{}, err
	}
	poller := &configurationsCreateOrUpdatePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ConfigurationsCreateOrUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new ConfigurationsCreateOrUpdatePoller from the specified resume token.
// token - The value must come from a previous call to ConfigurationsCreateOrUpdatePoller.ResumeToken().
func (client *ConfigurationsClient) ResumeCreateOrUpdate(ctx context.Context, token string) (ConfigurationsCreateOrUpdatePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("ConfigurationsClient.CreateOrUpdate", token, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return ConfigurationsCreateOrUpdatePollerResponse{}, err
	}
	poller := &configurationsCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ConfigurationsCreateOrUpdatePollerResponse{}, err
	}
	result := ConfigurationsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ConfigurationsCreateOrUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdate - Updates a configuration of a server.
// If the operation fails it returns the *CloudError error type.
func (client *ConfigurationsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, configurationName string, parameters Configuration, options *ConfigurationsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, configurationName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ConfigurationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, configurationName string, parameters Configuration, options *ConfigurationsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/configurations/{configurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2017-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ConfigurationsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Gets information about a configuration of server.
// If the operation fails it returns the *CloudError error type.
func (client *ConfigurationsClient) Get(ctx context.Context, resourceGroupName string, serverName string, configurationName string, options *ConfigurationsGetOptions) (ConfigurationsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, configurationName, options)
	if err != nil {
		return ConfigurationsGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ConfigurationsGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ConfigurationsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ConfigurationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, configurationName string, options *ConfigurationsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/configurations/{configurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2017-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ConfigurationsClient) getHandleResponse(resp *azcore.Response) (ConfigurationsGetResponse, error) {
	result := ConfigurationsGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.Configuration); err != nil {
		return ConfigurationsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ConfigurationsClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListByServer - List all the configurations in a given server.
// If the operation fails it returns the *CloudError error type.
func (client *ConfigurationsClient) ListByServer(ctx context.Context, resourceGroupName string, serverName string, options *ConfigurationsListByServerOptions) (ConfigurationsListByServerResponse, error) {
	req, err := client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
	if err != nil {
		return ConfigurationsListByServerResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ConfigurationsListByServerResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ConfigurationsListByServerResponse{}, client.listByServerHandleError(resp)
	}
	return client.listByServerHandleResponse(resp)
}

// listByServerCreateRequest creates the ListByServer request.
func (client *ConfigurationsClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ConfigurationsListByServerOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/configurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2017-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *ConfigurationsClient) listByServerHandleResponse(resp *azcore.Response) (ConfigurationsListByServerResponse, error) {
	result := ConfigurationsListByServerResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ConfigurationListResult); err != nil {
		return ConfigurationsListByServerResponse{}, err
	}
	return result, nil
}

// listByServerHandleError handles the ListByServer error response.
func (client *ConfigurationsClient) listByServerHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
