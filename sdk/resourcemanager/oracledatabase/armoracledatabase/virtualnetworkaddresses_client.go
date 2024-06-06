//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoracledatabase

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

// VirtualNetworkAddressesClient contains the methods for the VirtualNetworkAddresses group.
// Don't use this type directly, use NewVirtualNetworkAddressesClient() instead.
type VirtualNetworkAddressesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVirtualNetworkAddressesClient creates a new instance of VirtualNetworkAddressesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVirtualNetworkAddressesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualNetworkAddressesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VirtualNetworkAddressesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a VirtualNetworkAddress
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - cloudvmclustername - CloudVmCluster name
//   - virtualnetworkaddressname - Virtual IP address hostname.
//   - resource - Resource create parameters.
//   - options - VirtualNetworkAddressesClientBeginCreateOrUpdateOptions contains the optional parameters for the VirtualNetworkAddressesClient.BeginCreateOrUpdate
//     method.
func (client *VirtualNetworkAddressesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, cloudvmclustername string, virtualnetworkaddressname string, resource VirtualNetworkAddress, options *VirtualNetworkAddressesClientBeginCreateOrUpdateOptions) (*runtime.Poller[VirtualNetworkAddressesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, cloudvmclustername, virtualnetworkaddressname, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualNetworkAddressesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VirtualNetworkAddressesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a VirtualNetworkAddress
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
func (client *VirtualNetworkAddressesClient) createOrUpdate(ctx context.Context, resourceGroupName string, cloudvmclustername string, virtualnetworkaddressname string, resource VirtualNetworkAddress, options *VirtualNetworkAddressesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "VirtualNetworkAddressesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, cloudvmclustername, virtualnetworkaddressname, resource, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualNetworkAddressesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, cloudvmclustername string, virtualnetworkaddressname string, resource VirtualNetworkAddress, options *VirtualNetworkAddressesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Oracle.Database/cloudVmClusters/{cloudvmclustername}/virtualNetworkAddresses/{virtualnetworkaddressname}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudvmclustername == "" {
		return nil, errors.New("parameter cloudvmclustername cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudvmclustername}", url.PathEscape(cloudvmclustername))
	if virtualnetworkaddressname == "" {
		return nil, errors.New("parameter virtualnetworkaddressname cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualnetworkaddressname}", url.PathEscape(virtualnetworkaddressname))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a VirtualNetworkAddress
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - cloudvmclustername - CloudVmCluster name
//   - virtualnetworkaddressname - Virtual IP address hostname.
//   - options - VirtualNetworkAddressesClientBeginDeleteOptions contains the optional parameters for the VirtualNetworkAddressesClient.BeginDelete
//     method.
func (client *VirtualNetworkAddressesClient) BeginDelete(ctx context.Context, resourceGroupName string, cloudvmclustername string, virtualnetworkaddressname string, options *VirtualNetworkAddressesClientBeginDeleteOptions) (*runtime.Poller[VirtualNetworkAddressesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, cloudvmclustername, virtualnetworkaddressname, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualNetworkAddressesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VirtualNetworkAddressesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a VirtualNetworkAddress
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
func (client *VirtualNetworkAddressesClient) deleteOperation(ctx context.Context, resourceGroupName string, cloudvmclustername string, virtualnetworkaddressname string, options *VirtualNetworkAddressesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "VirtualNetworkAddressesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, cloudvmclustername, virtualnetworkaddressname, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VirtualNetworkAddressesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, cloudvmclustername string, virtualnetworkaddressname string, options *VirtualNetworkAddressesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Oracle.Database/cloudVmClusters/{cloudvmclustername}/virtualNetworkAddresses/{virtualnetworkaddressname}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudvmclustername == "" {
		return nil, errors.New("parameter cloudvmclustername cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudvmclustername}", url.PathEscape(cloudvmclustername))
	if virtualnetworkaddressname == "" {
		return nil, errors.New("parameter virtualnetworkaddressname cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualnetworkaddressname}", url.PathEscape(virtualnetworkaddressname))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a VirtualNetworkAddress
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - cloudvmclustername - CloudVmCluster name
//   - virtualnetworkaddressname - Virtual IP address hostname.
//   - options - VirtualNetworkAddressesClientGetOptions contains the optional parameters for the VirtualNetworkAddressesClient.Get
//     method.
func (client *VirtualNetworkAddressesClient) Get(ctx context.Context, resourceGroupName string, cloudvmclustername string, virtualnetworkaddressname string, options *VirtualNetworkAddressesClientGetOptions) (VirtualNetworkAddressesClientGetResponse, error) {
	var err error
	const operationName = "VirtualNetworkAddressesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, cloudvmclustername, virtualnetworkaddressname, options)
	if err != nil {
		return VirtualNetworkAddressesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualNetworkAddressesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VirtualNetworkAddressesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *VirtualNetworkAddressesClient) getCreateRequest(ctx context.Context, resourceGroupName string, cloudvmclustername string, virtualnetworkaddressname string, options *VirtualNetworkAddressesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Oracle.Database/cloudVmClusters/{cloudvmclustername}/virtualNetworkAddresses/{virtualnetworkaddressname}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudvmclustername == "" {
		return nil, errors.New("parameter cloudvmclustername cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudvmclustername}", url.PathEscape(cloudvmclustername))
	if virtualnetworkaddressname == "" {
		return nil, errors.New("parameter virtualnetworkaddressname cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualnetworkaddressname}", url.PathEscape(virtualnetworkaddressname))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualNetworkAddressesClient) getHandleResponse(resp *http.Response) (VirtualNetworkAddressesClientGetResponse, error) {
	result := VirtualNetworkAddressesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkAddress); err != nil {
		return VirtualNetworkAddressesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByCloudVMClusterPager - List VirtualNetworkAddress resources by CloudVmCluster
//
// Generated from API version 2023-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - cloudvmclustername - CloudVmCluster name
//   - options - VirtualNetworkAddressesClientListByCloudVMClusterOptions contains the optional parameters for the VirtualNetworkAddressesClient.NewListByCloudVMClusterPager
//     method.
func (client *VirtualNetworkAddressesClient) NewListByCloudVMClusterPager(resourceGroupName string, cloudvmclustername string, options *VirtualNetworkAddressesClientListByCloudVMClusterOptions) *runtime.Pager[VirtualNetworkAddressesClientListByCloudVMClusterResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualNetworkAddressesClientListByCloudVMClusterResponse]{
		More: func(page VirtualNetworkAddressesClientListByCloudVMClusterResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualNetworkAddressesClientListByCloudVMClusterResponse) (VirtualNetworkAddressesClientListByCloudVMClusterResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "VirtualNetworkAddressesClient.NewListByCloudVMClusterPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByCloudVMClusterCreateRequest(ctx, resourceGroupName, cloudvmclustername, options)
			}, nil)
			if err != nil {
				return VirtualNetworkAddressesClientListByCloudVMClusterResponse{}, err
			}
			return client.listByCloudVMClusterHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByCloudVMClusterCreateRequest creates the ListByCloudVMCluster request.
func (client *VirtualNetworkAddressesClient) listByCloudVMClusterCreateRequest(ctx context.Context, resourceGroupName string, cloudvmclustername string, options *VirtualNetworkAddressesClientListByCloudVMClusterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Oracle.Database/cloudVmClusters/{cloudvmclustername}/virtualNetworkAddresses"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudvmclustername == "" {
		return nil, errors.New("parameter cloudvmclustername cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudvmclustername}", url.PathEscape(cloudvmclustername))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByCloudVMClusterHandleResponse handles the ListByCloudVMCluster response.
func (client *VirtualNetworkAddressesClient) listByCloudVMClusterHandleResponse(resp *http.Response) (VirtualNetworkAddressesClientListByCloudVMClusterResponse, error) {
	result := VirtualNetworkAddressesClientListByCloudVMClusterResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkAddressListResult); err != nil {
		return VirtualNetworkAddressesClientListByCloudVMClusterResponse{}, err
	}
	return result, nil
}