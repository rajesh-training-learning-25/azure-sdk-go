//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcdn

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

// OriginsClient contains the methods for the Origins group.
// Don't use this type directly, use NewOriginsClient() instead.
type OriginsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewOriginsClient creates a new instance of OriginsClient with the specified values.
// subscriptionID - Azure Subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewOriginsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OriginsClient, error) {
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
	client := &OriginsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreate - Creates a new origin within the specified endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - Name of the Resource group within the Azure subscription.
// profileName - Name of the CDN profile which is unique within the resource group.
// endpointName - Name of the endpoint under the profile which is unique globally.
// originName - Name of the origin that is unique within the endpoint.
// origin - Origin properties
// options - OriginsClientBeginCreateOptions contains the optional parameters for the OriginsClient.BeginCreate method.
func (client *OriginsClient) BeginCreate(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, origin Origin, options *OriginsClientBeginCreateOptions) (*runtime.Poller[OriginsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, profileName, endpointName, originName, origin, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[OriginsClientCreateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[OriginsClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Creates a new origin within the specified endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
func (client *OriginsClient) create(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, origin Origin, options *OriginsClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, profileName, endpointName, originName, origin, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *OriginsClient) createCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, origin Origin, options *OriginsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/origins/{originName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if originName == "" {
		return nil, errors.New("parameter originName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{originName}", url.PathEscape(originName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, origin)
}

// BeginDelete - Deletes an existing origin within an endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - Name of the Resource group within the Azure subscription.
// profileName - Name of the CDN profile which is unique within the resource group.
// endpointName - Name of the endpoint under the profile which is unique globally.
// originName - Name of the origin which is unique within the endpoint.
// options - OriginsClientBeginDeleteOptions contains the optional parameters for the OriginsClient.BeginDelete method.
func (client *OriginsClient) BeginDelete(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, options *OriginsClientBeginDeleteOptions) (*runtime.Poller[OriginsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, profileName, endpointName, originName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[OriginsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[OriginsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes an existing origin within an endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
func (client *OriginsClient) deleteOperation(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, options *OriginsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, profileName, endpointName, originName, options)
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
func (client *OriginsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, options *OriginsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/origins/{originName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if originName == "" {
		return nil, errors.New("parameter originName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{originName}", url.PathEscape(originName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an existing origin within an endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - Name of the Resource group within the Azure subscription.
// profileName - Name of the CDN profile which is unique within the resource group.
// endpointName - Name of the endpoint under the profile which is unique globally.
// originName - Name of the origin which is unique within the endpoint.
// options - OriginsClientGetOptions contains the optional parameters for the OriginsClient.Get method.
func (client *OriginsClient) Get(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, options *OriginsClientGetOptions) (OriginsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, profileName, endpointName, originName, options)
	if err != nil {
		return OriginsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OriginsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OriginsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *OriginsClient) getCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, options *OriginsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/origins/{originName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if originName == "" {
		return nil, errors.New("parameter originName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{originName}", url.PathEscape(originName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OriginsClient) getHandleResponse(resp *http.Response) (OriginsClientGetResponse, error) {
	result := OriginsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Origin); err != nil {
		return OriginsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByEndpointPager - Lists all of the existing origins within an endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - Name of the Resource group within the Azure subscription.
// profileName - Name of the CDN profile which is unique within the resource group.
// endpointName - Name of the endpoint under the profile which is unique globally.
// options - OriginsClientListByEndpointOptions contains the optional parameters for the OriginsClient.ListByEndpoint method.
func (client *OriginsClient) NewListByEndpointPager(resourceGroupName string, profileName string, endpointName string, options *OriginsClientListByEndpointOptions) *runtime.Pager[OriginsClientListByEndpointResponse] {
	return runtime.NewPager(runtime.PagingHandler[OriginsClientListByEndpointResponse]{
		More: func(page OriginsClientListByEndpointResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OriginsClientListByEndpointResponse) (OriginsClientListByEndpointResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByEndpointCreateRequest(ctx, resourceGroupName, profileName, endpointName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return OriginsClientListByEndpointResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return OriginsClientListByEndpointResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return OriginsClientListByEndpointResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByEndpointHandleResponse(resp)
		},
	})
}

// listByEndpointCreateRequest creates the ListByEndpoint request.
func (client *OriginsClient) listByEndpointCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, options *OriginsClientListByEndpointOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/origins"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByEndpointHandleResponse handles the ListByEndpoint response.
func (client *OriginsClient) listByEndpointHandleResponse(resp *http.Response) (OriginsClientListByEndpointResponse, error) {
	result := OriginsClientListByEndpointResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OriginListResult); err != nil {
		return OriginsClientListByEndpointResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates an existing origin within an endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - Name of the Resource group within the Azure subscription.
// profileName - Name of the CDN profile which is unique within the resource group.
// endpointName - Name of the endpoint under the profile which is unique globally.
// originName - Name of the origin which is unique within the endpoint.
// originUpdateProperties - Origin properties
// options - OriginsClientBeginUpdateOptions contains the optional parameters for the OriginsClient.BeginUpdate method.
func (client *OriginsClient) BeginUpdate(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, originUpdateProperties OriginUpdateParameters, options *OriginsClientBeginUpdateOptions) (*runtime.Poller[OriginsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, profileName, endpointName, originName, originUpdateProperties, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[OriginsClientUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[OriginsClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Updates an existing origin within an endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
func (client *OriginsClient) update(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, originUpdateProperties OriginUpdateParameters, options *OriginsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, profileName, endpointName, originName, originUpdateProperties, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *OriginsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, originUpdateProperties OriginUpdateParameters, options *OriginsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/origins/{originName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if originName == "" {
		return nil, errors.New("parameter originName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{originName}", url.PathEscape(originName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, originUpdateProperties)
}
