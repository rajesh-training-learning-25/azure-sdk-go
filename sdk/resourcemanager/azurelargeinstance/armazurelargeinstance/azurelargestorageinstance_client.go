//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurelargeinstance

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AzureLargeStorageInstanceClient contains the methods for the AzureLargeStorageInstance group.
// Don't use this type directly, use NewAzureLargeStorageInstanceClient() instead.
type AzureLargeStorageInstanceClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAzureLargeStorageInstanceClient creates a new instance of AzureLargeStorageInstanceClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAzureLargeStorageInstanceClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AzureLargeStorageInstanceClient, error) {
	cl, err := arm.NewClient(moduleName+".AzureLargeStorageInstanceClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AzureLargeStorageInstanceClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets an Azure Large Storage instance for the specified subscription, resource group, and instance name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-20-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - azureLargeStorageInstanceName - Name of the AzureLargeStorageInstance.
//   - options - AzureLargeStorageInstanceClientGetOptions contains the optional parameters for the AzureLargeStorageInstanceClient.Get
//     method.
func (client *AzureLargeStorageInstanceClient) Get(ctx context.Context, resourceGroupName string, azureLargeStorageInstanceName string, options *AzureLargeStorageInstanceClientGetOptions) (AzureLargeStorageInstanceClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, azureLargeStorageInstanceName, options)
	if err != nil {
		return AzureLargeStorageInstanceClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AzureLargeStorageInstanceClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AzureLargeStorageInstanceClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AzureLargeStorageInstanceClient) getCreateRequest(ctx context.Context, resourceGroupName string, azureLargeStorageInstanceName string, options *AzureLargeStorageInstanceClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureLargeInstance/azureLargeStorageInstances/{azureLargeStorageInstanceName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureLargeStorageInstanceName == "" {
		return nil, errors.New("parameter azureLargeStorageInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureLargeStorageInstanceName}", url.PathEscape(azureLargeStorageInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-20-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AzureLargeStorageInstanceClient) getHandleResponse(resp *http.Response) (AzureLargeStorageInstanceClientGetResponse, error) {
	result := AzureLargeStorageInstanceClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureLargeStorageInstance); err != nil {
		return AzureLargeStorageInstanceClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets a list of AzureLargeStorageInstances in the specified subscription and resource group.
// The operations returns various properties of each Azure LargeStorage instance.
//
// Generated from API version 2023-07-20-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - AzureLargeStorageInstanceClientListByResourceGroupOptions contains the optional parameters for the AzureLargeStorageInstanceClient.NewListByResourceGroupPager
//     method.
func (client *AzureLargeStorageInstanceClient) NewListByResourceGroupPager(resourceGroupName string, options *AzureLargeStorageInstanceClientListByResourceGroupOptions) *runtime.Pager[AzureLargeStorageInstanceClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[AzureLargeStorageInstanceClientListByResourceGroupResponse]{
		More: func(page AzureLargeStorageInstanceClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AzureLargeStorageInstanceClientListByResourceGroupResponse) (AzureLargeStorageInstanceClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AzureLargeStorageInstanceClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return AzureLargeStorageInstanceClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AzureLargeStorageInstanceClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *AzureLargeStorageInstanceClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *AzureLargeStorageInstanceClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureLargeInstance/azureLargeStorageInstances"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-20-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *AzureLargeStorageInstanceClient) listByResourceGroupHandleResponse(resp *http.Response) (AzureLargeStorageInstanceClientListByResourceGroupResponse, error) {
	result := AzureLargeStorageInstanceClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureLargeStorageInstanceListResult); err != nil {
		return AzureLargeStorageInstanceClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Gets a list of AzureLargeStorageInstances in the specified subscription. The operations returns
// various properties of each Azure LargeStorage instance.
//
// Generated from API version 2023-07-20-preview
//   - options - AzureLargeStorageInstanceClientListBySubscriptionOptions contains the optional parameters for the AzureLargeStorageInstanceClient.NewListBySubscriptionPager
//     method.
func (client *AzureLargeStorageInstanceClient) NewListBySubscriptionPager(options *AzureLargeStorageInstanceClientListBySubscriptionOptions) *runtime.Pager[AzureLargeStorageInstanceClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[AzureLargeStorageInstanceClientListBySubscriptionResponse]{
		More: func(page AzureLargeStorageInstanceClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AzureLargeStorageInstanceClientListBySubscriptionResponse) (AzureLargeStorageInstanceClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AzureLargeStorageInstanceClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return AzureLargeStorageInstanceClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AzureLargeStorageInstanceClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *AzureLargeStorageInstanceClient) listBySubscriptionCreateRequest(ctx context.Context, options *AzureLargeStorageInstanceClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AzureLargeInstance/azureLargeStorageInstances"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-20-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *AzureLargeStorageInstanceClient) listBySubscriptionHandleResponse(resp *http.Response) (AzureLargeStorageInstanceClientListBySubscriptionResponse, error) {
	result := AzureLargeStorageInstanceClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureLargeStorageInstanceListResult); err != nil {
		return AzureLargeStorageInstanceClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Patches the Tags field of a Azure Large Storage Instance for the specified subscription, resource group, and instance
// name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-20-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - azureLargeStorageInstanceName - Name of the AzureLargeStorageInstance.
//   - tagsParameter - Request body that only contains the new Tags field
//   - options - AzureLargeStorageInstanceClientUpdateOptions contains the optional parameters for the AzureLargeStorageInstanceClient.Update
//     method.
func (client *AzureLargeStorageInstanceClient) Update(ctx context.Context, resourceGroupName string, azureLargeStorageInstanceName string, tagsParameter Tags, options *AzureLargeStorageInstanceClientUpdateOptions) (AzureLargeStorageInstanceClientUpdateResponse, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, azureLargeStorageInstanceName, tagsParameter, options)
	if err != nil {
		return AzureLargeStorageInstanceClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AzureLargeStorageInstanceClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AzureLargeStorageInstanceClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *AzureLargeStorageInstanceClient) updateCreateRequest(ctx context.Context, resourceGroupName string, azureLargeStorageInstanceName string, tagsParameter Tags, options *AzureLargeStorageInstanceClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureLargeInstance/azureLargeStorageInstances/{azureLargeStorageInstanceName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureLargeStorageInstanceName == "" {
		return nil, errors.New("parameter azureLargeStorageInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureLargeStorageInstanceName}", url.PathEscape(azureLargeStorageInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-20-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, tagsParameter); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *AzureLargeStorageInstanceClient) updateHandleResponse(resp *http.Response) (AzureLargeStorageInstanceClientUpdateResponse, error) {
	result := AzureLargeStorageInstanceClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureLargeStorageInstance); err != nil {
		return AzureLargeStorageInstanceClientUpdateResponse{}, err
	}
	return result, nil
}
