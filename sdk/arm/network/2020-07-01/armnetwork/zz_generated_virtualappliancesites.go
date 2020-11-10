// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// VirtualApplianceSitesOperations contains the methods for the VirtualApplianceSites group.
type VirtualApplianceSitesOperations interface {
	// BeginCreateOrUpdate - Creates or updates the specified Network Virtual Appliance Site.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, siteName string, parameters VirtualApplianceSite, options *VirtualApplianceSitesCreateOrUpdateOptions) (*VirtualApplianceSitePollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (VirtualApplianceSitePoller, error)
	// BeginDelete - Deletes the specified site from a Virtual Appliance.
	BeginDelete(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, siteName string, options *VirtualApplianceSitesDeleteOptions) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets the specified Virtual Appliance Site.
	Get(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, siteName string, options *VirtualApplianceSitesGetOptions) (*VirtualApplianceSiteResponse, error)
	// List - Lists all Network Virtual Appliance Sites in a Network Virtual Appliance resource.
	List(resourceGroupName string, networkVirtualApplianceName string, options *VirtualApplianceSitesListOptions) NetworkVirtualApplianceSiteListResultPager
}

// VirtualApplianceSitesClient implements the VirtualApplianceSitesOperations interface.
// Don't use this type directly, use NewVirtualApplianceSitesClient() instead.
type VirtualApplianceSitesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVirtualApplianceSitesClient creates a new instance of VirtualApplianceSitesClient with the specified values.
func NewVirtualApplianceSitesClient(con *armcore.Connection, subscriptionID string) VirtualApplianceSitesOperations {
	return &VirtualApplianceSitesClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *VirtualApplianceSitesClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

func (client *VirtualApplianceSitesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, siteName string, parameters VirtualApplianceSite, options *VirtualApplianceSitesCreateOrUpdateOptions) (*VirtualApplianceSitePollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, networkVirtualApplianceName, siteName, parameters, options)
	if err != nil {
		return nil, err
	}
	result := &VirtualApplianceSitePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualApplianceSitesClient.CreateOrUpdate", "azure-async-operation", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &virtualApplianceSitePoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*VirtualApplianceSiteResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *VirtualApplianceSitesClient) ResumeCreateOrUpdate(token string) (VirtualApplianceSitePoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualApplianceSitesClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &virtualApplianceSitePoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates the specified Network Virtual Appliance Site.
func (client *VirtualApplianceSitesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, siteName string, parameters VirtualApplianceSite, options *VirtualApplianceSitesCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, networkVirtualApplianceName, siteName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.CreateOrUpdateHandleError(resp)
	}
	return resp, nil
}

// CreateOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualApplianceSitesClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, siteName string, parameters VirtualApplianceSite, options *VirtualApplianceSitesCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}/virtualApplianceSites/{siteName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkVirtualApplianceName}", url.PathEscape(networkVirtualApplianceName))
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *VirtualApplianceSitesClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*VirtualApplianceSiteResponse, error) {
	result := VirtualApplianceSiteResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualApplianceSite)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *VirtualApplianceSitesClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *VirtualApplianceSitesClient) BeginDelete(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, siteName string, options *VirtualApplianceSitesDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, networkVirtualApplianceName, siteName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualApplianceSitesClient.Delete", "location", resp, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *VirtualApplianceSitesClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualApplianceSitesClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified site from a Virtual Appliance.
func (client *VirtualApplianceSitesClient) Delete(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, siteName string, options *VirtualApplianceSitesDeleteOptions) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, networkVirtualApplianceName, siteName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.DeleteHandleError(resp)
	}
	return resp, nil
}

// DeleteCreateRequest creates the Delete request.
func (client *VirtualApplianceSitesClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, siteName string, options *VirtualApplianceSitesDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}/virtualApplianceSites/{siteName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkVirtualApplianceName}", url.PathEscape(networkVirtualApplianceName))
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DeleteHandleError handles the Delete error response.
func (client *VirtualApplianceSitesClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified Virtual Appliance Site.
func (client *VirtualApplianceSitesClient) Get(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, siteName string, options *VirtualApplianceSitesGetOptions) (*VirtualApplianceSiteResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, networkVirtualApplianceName, siteName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetHandleError(resp)
	}
	result, err := client.GetHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCreateRequest creates the Get request.
func (client *VirtualApplianceSitesClient) GetCreateRequest(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, siteName string, options *VirtualApplianceSitesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}/virtualApplianceSites/{siteName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkVirtualApplianceName}", url.PathEscape(networkVirtualApplianceName))
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *VirtualApplianceSitesClient) GetHandleResponse(resp *azcore.Response) (*VirtualApplianceSiteResponse, error) {
	result := VirtualApplianceSiteResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualApplianceSite)
}

// GetHandleError handles the Get error response.
func (client *VirtualApplianceSitesClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Lists all Network Virtual Appliance Sites in a Network Virtual Appliance resource.
func (client *VirtualApplianceSitesClient) List(resourceGroupName string, networkVirtualApplianceName string, options *VirtualApplianceSitesListOptions) NetworkVirtualApplianceSiteListResultPager {
	return &networkVirtualApplianceSiteListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, resourceGroupName, networkVirtualApplianceName, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *NetworkVirtualApplianceSiteListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.NetworkVirtualApplianceSiteListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListCreateRequest creates the List request.
func (client *VirtualApplianceSitesClient) ListCreateRequest(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, options *VirtualApplianceSitesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}/virtualApplianceSites"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkVirtualApplianceName}", url.PathEscape(networkVirtualApplianceName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListHandleResponse handles the List response.
func (client *VirtualApplianceSitesClient) ListHandleResponse(resp *azcore.Response) (*NetworkVirtualApplianceSiteListResultResponse, error) {
	result := NetworkVirtualApplianceSiteListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NetworkVirtualApplianceSiteListResult)
}

// ListHandleError handles the List error response.
func (client *VirtualApplianceSitesClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
