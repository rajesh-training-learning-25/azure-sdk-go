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
	"github.com/Azure/azure-sdk-for-go/profile/v20200901/internal"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// VirtualNetworksClient contains the methods for the VirtualNetworks group.
// Don't use this type directly, use NewVirtualNetworksClient() instead.
type VirtualNetworksClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVirtualNetworksClient creates a new instance of VirtualNetworksClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVirtualNetworksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualNetworksClient, error) {
	cl, err := arm.NewClient(internal.ModuleName+"/armnetwork.VirtualNetworksClient", internal.ModuleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VirtualNetworksClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CheckIPAddressAvailability - Checks whether a private IP address is available for use.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-11-01
//   - resourceGroupName - The name of the resource group.
//   - virtualNetworkName - The name of the virtual network.
//   - ipAddress - The private IP address to be verified.
//   - options - VirtualNetworksClientCheckIPAddressAvailabilityOptions contains the optional parameters for the VirtualNetworksClient.CheckIPAddressAvailability
//     method.
func (client *VirtualNetworksClient) CheckIPAddressAvailability(ctx context.Context, resourceGroupName string, virtualNetworkName string, ipAddress string, options *VirtualNetworksClientCheckIPAddressAvailabilityOptions) (VirtualNetworksClientCheckIPAddressAvailabilityResponse, error) {
	req, err := client.checkIPAddressAvailabilityCreateRequest(ctx, resourceGroupName, virtualNetworkName, ipAddress, options)
	if err != nil {
		return VirtualNetworksClientCheckIPAddressAvailabilityResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualNetworksClientCheckIPAddressAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualNetworksClientCheckIPAddressAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkIPAddressAvailabilityHandleResponse(resp)
}

// checkIPAddressAvailabilityCreateRequest creates the CheckIPAddressAvailability request.
func (client *VirtualNetworksClient) checkIPAddressAvailabilityCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, ipAddress string, options *VirtualNetworksClientCheckIPAddressAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/CheckIPAddressAvailability"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("ipAddress", ipAddress)
	reqQP.Set("api-version", "2018-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// checkIPAddressAvailabilityHandleResponse handles the CheckIPAddressAvailability response.
func (client *VirtualNetworksClient) checkIPAddressAvailabilityHandleResponse(resp *http.Response) (VirtualNetworksClientCheckIPAddressAvailabilityResponse, error) {
	result := VirtualNetworksClientCheckIPAddressAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IPAddressAvailabilityResult); err != nil {
		return VirtualNetworksClientCheckIPAddressAvailabilityResponse{}, err
	}
	return result, nil
}

// BeginCreateOrUpdate - Creates or updates a virtual network in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-11-01
//   - resourceGroupName - The name of the resource group.
//   - virtualNetworkName - The name of the virtual network.
//   - parameters - Parameters supplied to the create or update virtual network operation
//   - options - VirtualNetworksClientBeginCreateOrUpdateOptions contains the optional parameters for the VirtualNetworksClient.BeginCreateOrUpdate
//     method.
func (client *VirtualNetworksClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkName string, parameters VirtualNetwork, options *VirtualNetworksClientBeginCreateOrUpdateOptions) (*runtime.Poller[VirtualNetworksClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, virtualNetworkName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[VirtualNetworksClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[VirtualNetworksClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates a virtual network in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-11-01
func (client *VirtualNetworksClient) createOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkName string, parameters VirtualNetwork, options *VirtualNetworksClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, virtualNetworkName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualNetworksClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, parameters VirtualNetwork, options *VirtualNetworksClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified virtual network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-11-01
//   - resourceGroupName - The name of the resource group.
//   - virtualNetworkName - The name of the virtual network.
//   - options - VirtualNetworksClientBeginDeleteOptions contains the optional parameters for the VirtualNetworksClient.BeginDelete
//     method.
func (client *VirtualNetworksClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualNetworkName string, options *VirtualNetworksClientBeginDeleteOptions) (*runtime.Poller[VirtualNetworksClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, virtualNetworkName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[VirtualNetworksClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[VirtualNetworksClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the specified virtual network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-11-01
func (client *VirtualNetworksClient) deleteOperation(ctx context.Context, resourceGroupName string, virtualNetworkName string, options *VirtualNetworksClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualNetworkName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VirtualNetworksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, options *VirtualNetworksClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets the specified virtual network by resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-11-01
//   - resourceGroupName - The name of the resource group.
//   - virtualNetworkName - The name of the virtual network.
//   - options - VirtualNetworksClientGetOptions contains the optional parameters for the VirtualNetworksClient.Get method.
func (client *VirtualNetworksClient) Get(ctx context.Context, resourceGroupName string, virtualNetworkName string, options *VirtualNetworksClientGetOptions) (VirtualNetworksClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualNetworkName, options)
	if err != nil {
		return VirtualNetworksClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualNetworksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualNetworksClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualNetworksClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, options *VirtualNetworksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
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
func (client *VirtualNetworksClient) getHandleResponse(resp *http.Response) (VirtualNetworksClientGetResponse, error) {
	result := VirtualNetworksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetwork); err != nil {
		return VirtualNetworksClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all virtual networks in a resource group.
//
// Generated from API version 2018-11-01
//   - resourceGroupName - The name of the resource group.
//   - options - VirtualNetworksClientListOptions contains the optional parameters for the VirtualNetworksClient.NewListPager
//     method.
func (client *VirtualNetworksClient) NewListPager(resourceGroupName string, options *VirtualNetworksClientListOptions) *runtime.Pager[VirtualNetworksClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualNetworksClientListResponse]{
		More: func(page VirtualNetworksClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualNetworksClientListResponse) (VirtualNetworksClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VirtualNetworksClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return VirtualNetworksClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VirtualNetworksClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *VirtualNetworksClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *VirtualNetworksClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
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
func (client *VirtualNetworksClient) listHandleResponse(resp *http.Response) (VirtualNetworksClientListResponse, error) {
	result := VirtualNetworksClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkListResult); err != nil {
		return VirtualNetworksClientListResponse{}, err
	}
	return result, nil
}

// NewListAllPager - Gets all virtual networks in a subscription.
//
// Generated from API version 2018-11-01
//   - options - VirtualNetworksClientListAllOptions contains the optional parameters for the VirtualNetworksClient.NewListAllPager
//     method.
func (client *VirtualNetworksClient) NewListAllPager(options *VirtualNetworksClientListAllOptions) *runtime.Pager[VirtualNetworksClientListAllResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualNetworksClientListAllResponse]{
		More: func(page VirtualNetworksClientListAllResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualNetworksClientListAllResponse) (VirtualNetworksClientListAllResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAllCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VirtualNetworksClientListAllResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return VirtualNetworksClientListAllResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VirtualNetworksClientListAllResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAllHandleResponse(resp)
		},
	})
}

// listAllCreateRequest creates the ListAll request.
func (client *VirtualNetworksClient) listAllCreateRequest(ctx context.Context, options *VirtualNetworksClientListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualNetworks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
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
func (client *VirtualNetworksClient) listAllHandleResponse(resp *http.Response) (VirtualNetworksClientListAllResponse, error) {
	result := VirtualNetworksClientListAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkListResult); err != nil {
		return VirtualNetworksClientListAllResponse{}, err
	}
	return result, nil
}

// NewListUsagePager - Lists usage stats.
//
// Generated from API version 2018-11-01
//   - resourceGroupName - The name of the resource group.
//   - virtualNetworkName - The name of the virtual network.
//   - options - VirtualNetworksClientListUsageOptions contains the optional parameters for the VirtualNetworksClient.NewListUsagePager
//     method.
func (client *VirtualNetworksClient) NewListUsagePager(resourceGroupName string, virtualNetworkName string, options *VirtualNetworksClientListUsageOptions) *runtime.Pager[VirtualNetworksClientListUsageResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualNetworksClientListUsageResponse]{
		More: func(page VirtualNetworksClientListUsageResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualNetworksClientListUsageResponse) (VirtualNetworksClientListUsageResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listUsageCreateRequest(ctx, resourceGroupName, virtualNetworkName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VirtualNetworksClientListUsageResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return VirtualNetworksClientListUsageResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VirtualNetworksClientListUsageResponse{}, runtime.NewResponseError(resp)
			}
			return client.listUsageHandleResponse(resp)
		},
	})
}

// listUsageCreateRequest creates the ListUsage request.
func (client *VirtualNetworksClient) listUsageCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, options *VirtualNetworksClientListUsageOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/usages"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listUsageHandleResponse handles the ListUsage response.
func (client *VirtualNetworksClient) listUsageHandleResponse(resp *http.Response) (VirtualNetworksClientListUsageResponse, error) {
	result := VirtualNetworksClientListUsageResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkListUsageResult); err != nil {
		return VirtualNetworksClientListUsageResponse{}, err
	}
	return result, nil
}

// BeginUpdateTags - Updates a virtual network tags.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-11-01
//   - resourceGroupName - The name of the resource group.
//   - virtualNetworkName - The name of the virtual network.
//   - parameters - Parameters supplied to update virtual network tags.
//   - options - VirtualNetworksClientBeginUpdateTagsOptions contains the optional parameters for the VirtualNetworksClient.BeginUpdateTags
//     method.
func (client *VirtualNetworksClient) BeginUpdateTags(ctx context.Context, resourceGroupName string, virtualNetworkName string, parameters TagsObject, options *VirtualNetworksClientBeginUpdateTagsOptions) (*runtime.Poller[VirtualNetworksClientUpdateTagsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateTags(ctx, resourceGroupName, virtualNetworkName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[VirtualNetworksClientUpdateTagsResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[VirtualNetworksClientUpdateTagsResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// UpdateTags - Updates a virtual network tags.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-11-01
func (client *VirtualNetworksClient) updateTags(ctx context.Context, resourceGroupName string, virtualNetworkName string, parameters TagsObject, options *VirtualNetworksClientBeginUpdateTagsOptions) (*http.Response, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, virtualNetworkName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *VirtualNetworksClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, parameters TagsObject, options *VirtualNetworksClientBeginUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}
