//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmobilenetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ServicesClient contains the methods for the Services group.
// Don't use this type directly, use NewServicesClient() instead.
type ServicesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewServicesClient creates a new instance of ServicesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewServicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServicesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ServicesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a service. Must be created in the same location as its parent mobile network.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// mobileNetworkName - The name of the mobile network.
// serviceName - The name of the service. You must not use any of the following reserved strings - default, requested or service
// parameters - Parameters supplied to the create or update service operation.
// options - ServicesClientBeginCreateOrUpdateOptions contains the optional parameters for the ServicesClient.BeginCreateOrUpdate
// method.
func (client *ServicesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, mobileNetworkName string, serviceName string, parameters Service, options *ServicesClientBeginCreateOrUpdateOptions) (*runtime.Poller[ServicesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, mobileNetworkName, serviceName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[ServicesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ServicesClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a service. Must be created in the same location as its parent mobile network.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01
func (client *ServicesClient) createOrUpdate(ctx context.Context, resourceGroupName string, mobileNetworkName string, serviceName string, parameters Service, options *ServicesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, mobileNetworkName, serviceName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServicesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, mobileNetworkName string, serviceName string, parameters Service, options *ServicesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/mobileNetworks/{mobileNetworkName}/services/{serviceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if mobileNetworkName == "" {
		return nil, errors.New("parameter mobileNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mobileNetworkName}", url.PathEscape(mobileNetworkName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified service.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// mobileNetworkName - The name of the mobile network.
// serviceName - The name of the service. You must not use any of the following reserved strings - default, requested or service
// options - ServicesClientBeginDeleteOptions contains the optional parameters for the ServicesClient.BeginDelete method.
func (client *ServicesClient) BeginDelete(ctx context.Context, resourceGroupName string, mobileNetworkName string, serviceName string, options *ServicesClientBeginDeleteOptions) (*runtime.Poller[ServicesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, mobileNetworkName, serviceName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[ServicesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ServicesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the specified service.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01
func (client *ServicesClient) deleteOperation(ctx context.Context, resourceGroupName string, mobileNetworkName string, serviceName string, options *ServicesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, mobileNetworkName, serviceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ServicesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, mobileNetworkName string, serviceName string, options *ServicesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/mobileNetworks/{mobileNetworkName}/services/{serviceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if mobileNetworkName == "" {
		return nil, errors.New("parameter mobileNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mobileNetworkName}", url.PathEscape(mobileNetworkName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets information about the specified service.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// mobileNetworkName - The name of the mobile network.
// serviceName - The name of the service. You must not use any of the following reserved strings - default, requested or service
// options - ServicesClientGetOptions contains the optional parameters for the ServicesClient.Get method.
func (client *ServicesClient) Get(ctx context.Context, resourceGroupName string, mobileNetworkName string, serviceName string, options *ServicesClientGetOptions) (ServicesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, mobileNetworkName, serviceName, options)
	if err != nil {
		return ServicesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServicesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServicesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServicesClient) getCreateRequest(ctx context.Context, resourceGroupName string, mobileNetworkName string, serviceName string, options *ServicesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/mobileNetworks/{mobileNetworkName}/services/{serviceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if mobileNetworkName == "" {
		return nil, errors.New("parameter mobileNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mobileNetworkName}", url.PathEscape(mobileNetworkName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServicesClient) getHandleResponse(resp *http.Response) (ServicesClientGetResponse, error) {
	result := ServicesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Service); err != nil {
		return ServicesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByMobileNetworkPager - Gets all the services in a mobile network.
// Generated from API version 2022-11-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// mobileNetworkName - The name of the mobile network.
// options - ServicesClientListByMobileNetworkOptions contains the optional parameters for the ServicesClient.ListByMobileNetwork
// method.
func (client *ServicesClient) NewListByMobileNetworkPager(resourceGroupName string, mobileNetworkName string, options *ServicesClientListByMobileNetworkOptions) *runtime.Pager[ServicesClientListByMobileNetworkResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServicesClientListByMobileNetworkResponse]{
		More: func(page ServicesClientListByMobileNetworkResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServicesClientListByMobileNetworkResponse) (ServicesClientListByMobileNetworkResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByMobileNetworkCreateRequest(ctx, resourceGroupName, mobileNetworkName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ServicesClientListByMobileNetworkResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ServicesClientListByMobileNetworkResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServicesClientListByMobileNetworkResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByMobileNetworkHandleResponse(resp)
		},
	})
}

// listByMobileNetworkCreateRequest creates the ListByMobileNetwork request.
func (client *ServicesClient) listByMobileNetworkCreateRequest(ctx context.Context, resourceGroupName string, mobileNetworkName string, options *ServicesClientListByMobileNetworkOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/mobileNetworks/{mobileNetworkName}/services"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if mobileNetworkName == "" {
		return nil, errors.New("parameter mobileNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mobileNetworkName}", url.PathEscape(mobileNetworkName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByMobileNetworkHandleResponse handles the ListByMobileNetwork response.
func (client *ServicesClient) listByMobileNetworkHandleResponse(resp *http.Response) (ServicesClientListByMobileNetworkResponse, error) {
	result := ServicesClientListByMobileNetworkResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceListResult); err != nil {
		return ServicesClientListByMobileNetworkResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates service tags.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// mobileNetworkName - The name of the mobile network.
// serviceName - The name of the service. You must not use any of the following reserved strings - default, requested or service
// parameters - Parameters supplied to update service tags.
// options - ServicesClientUpdateTagsOptions contains the optional parameters for the ServicesClient.UpdateTags method.
func (client *ServicesClient) UpdateTags(ctx context.Context, resourceGroupName string, mobileNetworkName string, serviceName string, parameters TagsObject, options *ServicesClientUpdateTagsOptions) (ServicesClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, mobileNetworkName, serviceName, parameters, options)
	if err != nil {
		return ServicesClientUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServicesClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServicesClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *ServicesClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, mobileNetworkName string, serviceName string, parameters TagsObject, options *ServicesClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/mobileNetworks/{mobileNetworkName}/services/{serviceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if mobileNetworkName == "" {
		return nil, errors.New("parameter mobileNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mobileNetworkName}", url.PathEscape(mobileNetworkName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *ServicesClient) updateTagsHandleResponse(resp *http.Response) (ServicesClientUpdateTagsResponse, error) {
	result := ServicesClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Service); err != nil {
		return ServicesClientUpdateTagsResponse{}, err
	}
	return result, nil
}
