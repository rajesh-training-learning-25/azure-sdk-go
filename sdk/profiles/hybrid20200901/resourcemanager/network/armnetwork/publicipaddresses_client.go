//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/profiles/hybrid20200901/internal"
	"net/http"
	"net/url"
	"strings"
)

// PublicIPAddressesClient contains the methods for the PublicIPAddresses group.
// Don't use this type directly, use NewPublicIPAddressesClient() instead.
type PublicIPAddressesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPublicIPAddressesClient creates a new instance of PublicIPAddressesClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPublicIPAddressesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PublicIPAddressesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(internal.ModuleName, internal.ModuleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &PublicIPAddressesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a static or dynamic public IP address.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-11-01
// resourceGroupName - The name of the resource group.
// publicIPAddressName - The name of the public IP address.
// parameters - Parameters supplied to the create or update public IP address operation.
// options - PublicIPAddressesClientBeginCreateOrUpdateOptions contains the optional parameters for the PublicIPAddressesClient.BeginCreateOrUpdate
// method.
func (client *PublicIPAddressesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, publicIPAddressName string, parameters PublicIPAddress, options *PublicIPAddressesClientBeginCreateOrUpdateOptions) (*runtime.Poller[PublicIPAddressesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, publicIPAddressName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[PublicIPAddressesClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[PublicIPAddressesClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a static or dynamic public IP address.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-11-01
func (client *PublicIPAddressesClient) createOrUpdate(ctx context.Context, resourceGroupName string, publicIPAddressName string, parameters PublicIPAddress, options *PublicIPAddressesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, publicIPAddressName, parameters, options)
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
func (client *PublicIPAddressesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, publicIPAddressName string, parameters PublicIPAddress, options *PublicIPAddressesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publicIPAddressName == "" {
		return nil, errors.New("parameter publicIPAddressName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publicIpAddressName}", url.PathEscape(publicIPAddressName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified public IP address.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-11-01
// resourceGroupName - The name of the resource group.
// publicIPAddressName - The name of the subnet.
// options - PublicIPAddressesClientBeginDeleteOptions contains the optional parameters for the PublicIPAddressesClient.BeginDelete
// method.
func (client *PublicIPAddressesClient) BeginDelete(ctx context.Context, resourceGroupName string, publicIPAddressName string, options *PublicIPAddressesClientBeginDeleteOptions) (*runtime.Poller[PublicIPAddressesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, publicIPAddressName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[PublicIPAddressesClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[PublicIPAddressesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the specified public IP address.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-11-01
func (client *PublicIPAddressesClient) deleteOperation(ctx context.Context, resourceGroupName string, publicIPAddressName string, options *PublicIPAddressesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, publicIPAddressName, options)
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
func (client *PublicIPAddressesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, publicIPAddressName string, options *PublicIPAddressesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publicIPAddressName == "" {
		return nil, errors.New("parameter publicIPAddressName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publicIpAddressName}", url.PathEscape(publicIPAddressName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets the specified public IP address in a specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-11-01
// resourceGroupName - The name of the resource group.
// publicIPAddressName - The name of the subnet.
// options - PublicIPAddressesClientGetOptions contains the optional parameters for the PublicIPAddressesClient.Get method.
func (client *PublicIPAddressesClient) Get(ctx context.Context, resourceGroupName string, publicIPAddressName string, options *PublicIPAddressesClientGetOptions) (PublicIPAddressesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, publicIPAddressName, options)
	if err != nil {
		return PublicIPAddressesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PublicIPAddressesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PublicIPAddressesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PublicIPAddressesClient) getCreateRequest(ctx context.Context, resourceGroupName string, publicIPAddressName string, options *PublicIPAddressesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publicIPAddressName == "" {
		return nil, errors.New("parameter publicIPAddressName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publicIpAddressName}", url.PathEscape(publicIPAddressName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PublicIPAddressesClient) getHandleResponse(resp *http.Response) (PublicIPAddressesClientGetResponse, error) {
	result := PublicIPAddressesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PublicIPAddress); err != nil {
		return PublicIPAddressesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all public IP addresses in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-11-01
// resourceGroupName - The name of the resource group.
// options - PublicIPAddressesClientListOptions contains the optional parameters for the PublicIPAddressesClient.List method.
func (client *PublicIPAddressesClient) NewListPager(resourceGroupName string, options *PublicIPAddressesClientListOptions) *runtime.Pager[PublicIPAddressesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[PublicIPAddressesClientListResponse]{
		More: func(page PublicIPAddressesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PublicIPAddressesClientListResponse) (PublicIPAddressesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PublicIPAddressesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PublicIPAddressesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PublicIPAddressesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *PublicIPAddressesClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *PublicIPAddressesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PublicIPAddressesClient) listHandleResponse(resp *http.Response) (PublicIPAddressesClientListResponse, error) {
	result := PublicIPAddressesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PublicIPAddressListResult); err != nil {
		return PublicIPAddressesClientListResponse{}, err
	}
	return result, nil
}

// NewListAllPager - Gets all the public IP addresses in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-11-01
// options - PublicIPAddressesClientListAllOptions contains the optional parameters for the PublicIPAddressesClient.ListAll
// method.
func (client *PublicIPAddressesClient) NewListAllPager(options *PublicIPAddressesClientListAllOptions) *runtime.Pager[PublicIPAddressesClientListAllResponse] {
	return runtime.NewPager(runtime.PagingHandler[PublicIPAddressesClientListAllResponse]{
		More: func(page PublicIPAddressesClientListAllResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PublicIPAddressesClientListAllResponse) (PublicIPAddressesClientListAllResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAllCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PublicIPAddressesClientListAllResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PublicIPAddressesClientListAllResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PublicIPAddressesClientListAllResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAllHandleResponse(resp)
		},
	})
}

// listAllCreateRequest creates the ListAll request.
func (client *PublicIPAddressesClient) listAllCreateRequest(ctx context.Context, options *PublicIPAddressesClientListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/publicIPAddresses"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *PublicIPAddressesClient) listAllHandleResponse(resp *http.Response) (PublicIPAddressesClientListAllResponse, error) {
	result := PublicIPAddressesClientListAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PublicIPAddressListResult); err != nil {
		return PublicIPAddressesClientListAllResponse{}, err
	}
	return result, nil
}

// BeginUpdateTags - Updates public IP address tags.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-11-01
// resourceGroupName - The name of the resource group.
// publicIPAddressName - The name of the public IP address.
// parameters - Parameters supplied to update public IP address tags.
// options - PublicIPAddressesClientBeginUpdateTagsOptions contains the optional parameters for the PublicIPAddressesClient.BeginUpdateTags
// method.
func (client *PublicIPAddressesClient) BeginUpdateTags(ctx context.Context, resourceGroupName string, publicIPAddressName string, parameters TagsObject, options *PublicIPAddressesClientBeginUpdateTagsOptions) (*runtime.Poller[PublicIPAddressesClientUpdateTagsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateTags(ctx, resourceGroupName, publicIPAddressName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[PublicIPAddressesClientUpdateTagsResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[PublicIPAddressesClientUpdateTagsResponse](options.ResumeToken, client.pl, nil)
	}
}

// UpdateTags - Updates public IP address tags.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-11-01
func (client *PublicIPAddressesClient) updateTags(ctx context.Context, resourceGroupName string, publicIPAddressName string, parameters TagsObject, options *PublicIPAddressesClientBeginUpdateTagsOptions) (*http.Response, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, publicIPAddressName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *PublicIPAddressesClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, publicIPAddressName string, parameters TagsObject, options *PublicIPAddressesClientBeginUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publicIPAddressName == "" {
		return nil, errors.New("parameter publicIPAddressName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publicIpAddressName}", url.PathEscape(publicIPAddressName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}
