// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armcomputefleet

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

// FleetsClient contains the methods for the Fleets group.
// Don't use this type directly, use NewFleetsClient() instead.
type FleetsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewFleetsClient creates a new instance of FleetsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewFleetsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*FleetsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &FleetsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a Fleet
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - fleetName - The name of the Compute Fleet
//   - resource - Resource create parameters.
//   - options - FleetsClientBeginCreateOrUpdateOptions contains the optional parameters for the FleetsClient.BeginCreateOrUpdate
//     method.
func (client *FleetsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, fleetName string, resource Fleet, options *FleetsClientBeginCreateOrUpdateOptions) (*runtime.Poller[FleetsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, fleetName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[FleetsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[FleetsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a Fleet
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
func (client *FleetsClient) createOrUpdate(ctx context.Context, resourceGroupName string, fleetName string, resource Fleet, options *FleetsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "FleetsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, fleetName, resource, options)
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
func (client *FleetsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, fleetName string, resource Fleet, _ *FleetsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureFleet/fleets/{fleetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if fleetName == "" {
		return nil, errors.New("parameter fleetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fleetName}", url.PathEscape(fleetName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a Fleet
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - fleetName - The name of the Compute Fleet
//   - options - FleetsClientBeginDeleteOptions contains the optional parameters for the FleetsClient.BeginDelete method.
func (client *FleetsClient) BeginDelete(ctx context.Context, resourceGroupName string, fleetName string, options *FleetsClientBeginDeleteOptions) (*runtime.Poller[FleetsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, fleetName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[FleetsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[FleetsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a Fleet
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
func (client *FleetsClient) deleteOperation(ctx context.Context, resourceGroupName string, fleetName string, options *FleetsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "FleetsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, fleetName, options)
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
func (client *FleetsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, fleetName string, _ *FleetsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureFleet/fleets/{fleetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if fleetName == "" {
		return nil, errors.New("parameter fleetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fleetName}", url.PathEscape(fleetName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a Fleet
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - fleetName - The name of the Compute Fleet
//   - options - FleetsClientGetOptions contains the optional parameters for the FleetsClient.Get method.
func (client *FleetsClient) Get(ctx context.Context, resourceGroupName string, fleetName string, options *FleetsClientGetOptions) (FleetsClientGetResponse, error) {
	var err error
	const operationName = "FleetsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, fleetName, options)
	if err != nil {
		return FleetsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FleetsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return FleetsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *FleetsClient) getCreateRequest(ctx context.Context, resourceGroupName string, fleetName string, _ *FleetsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureFleet/fleets/{fleetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if fleetName == "" {
		return nil, errors.New("parameter fleetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fleetName}", url.PathEscape(fleetName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FleetsClient) getHandleResponse(resp *http.Response) (FleetsClientGetResponse, error) {
	result := FleetsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Fleet); err != nil {
		return FleetsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List Fleet resources by resource group
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - FleetsClientListByResourceGroupOptions contains the optional parameters for the FleetsClient.NewListByResourceGroupPager
//     method.
func (client *FleetsClient) NewListByResourceGroupPager(resourceGroupName string, options *FleetsClientListByResourceGroupOptions) *runtime.Pager[FleetsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[FleetsClientListByResourceGroupResponse]{
		More: func(page FleetsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FleetsClientListByResourceGroupResponse) (FleetsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "FleetsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return FleetsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *FleetsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, _ *FleetsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureFleet/fleets"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
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
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *FleetsClient) listByResourceGroupHandleResponse(resp *http.Response) (FleetsClientListByResourceGroupResponse, error) {
	result := FleetsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FleetListResult); err != nil {
		return FleetsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List Fleet resources by subscription ID
//
// Generated from API version 2024-11-01
//   - options - FleetsClientListBySubscriptionOptions contains the optional parameters for the FleetsClient.NewListBySubscriptionPager
//     method.
func (client *FleetsClient) NewListBySubscriptionPager(options *FleetsClientListBySubscriptionOptions) *runtime.Pager[FleetsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[FleetsClientListBySubscriptionResponse]{
		More: func(page FleetsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FleetsClientListBySubscriptionResponse) (FleetsClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "FleetsClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return FleetsClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *FleetsClient) listBySubscriptionCreateRequest(ctx context.Context, _ *FleetsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AzureFleet/fleets"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *FleetsClient) listBySubscriptionHandleResponse(resp *http.Response) (FleetsClientListBySubscriptionResponse, error) {
	result := FleetsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FleetListResult); err != nil {
		return FleetsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// NewListVirtualMachineScaleSetsPager - List VirtualMachineScaleSet resources by Fleet
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - name - The name of the Fleet
//   - options - FleetsClientListVirtualMachineScaleSetsOptions contains the optional parameters for the FleetsClient.NewListVirtualMachineScaleSetsPager
//     method.
func (client *FleetsClient) NewListVirtualMachineScaleSetsPager(resourceGroupName string, name string, options *FleetsClientListVirtualMachineScaleSetsOptions) *runtime.Pager[FleetsClientListVirtualMachineScaleSetsResponse] {
	return runtime.NewPager(runtime.PagingHandler[FleetsClientListVirtualMachineScaleSetsResponse]{
		More: func(page FleetsClientListVirtualMachineScaleSetsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FleetsClientListVirtualMachineScaleSetsResponse) (FleetsClientListVirtualMachineScaleSetsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "FleetsClient.NewListVirtualMachineScaleSetsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listVirtualMachineScaleSetsCreateRequest(ctx, resourceGroupName, name, options)
			}, nil)
			if err != nil {
				return FleetsClientListVirtualMachineScaleSetsResponse{}, err
			}
			return client.listVirtualMachineScaleSetsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listVirtualMachineScaleSetsCreateRequest creates the ListVirtualMachineScaleSets request.
func (client *FleetsClient) listVirtualMachineScaleSetsCreateRequest(ctx context.Context, resourceGroupName string, name string, _ *FleetsClientListVirtualMachineScaleSetsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureFleet/fleets/{name}/virtualMachineScaleSets"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listVirtualMachineScaleSetsHandleResponse handles the ListVirtualMachineScaleSets response.
func (client *FleetsClient) listVirtualMachineScaleSetsHandleResponse(resp *http.Response) (FleetsClientListVirtualMachineScaleSetsResponse, error) {
	result := FleetsClientListVirtualMachineScaleSetsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineScaleSetListResult); err != nil {
		return FleetsClientListVirtualMachineScaleSetsResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a Fleet
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - fleetName - The name of the Compute Fleet
//   - properties - The resource properties to be updated.
//   - options - FleetsClientBeginUpdateOptions contains the optional parameters for the FleetsClient.BeginUpdate method.
func (client *FleetsClient) BeginUpdate(ctx context.Context, resourceGroupName string, fleetName string, properties FleetUpdate, options *FleetsClientBeginUpdateOptions) (*runtime.Poller[FleetsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, fleetName, properties, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[FleetsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[FleetsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update a Fleet
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
func (client *FleetsClient) update(ctx context.Context, resourceGroupName string, fleetName string, properties FleetUpdate, options *FleetsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "FleetsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, fleetName, properties, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *FleetsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, fleetName string, properties FleetUpdate, _ *FleetsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureFleet/fleets/{fleetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if fleetName == "" {
		return nil, errors.New("parameter fleetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fleetName}", url.PathEscape(fleetName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}
