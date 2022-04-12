//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetapp

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

// SubvolumesClient contains the methods for the Subvolumes group.
// Don't use this type directly, use NewSubvolumesClient() instead.
type SubvolumesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSubvolumesClient creates a new instance of SubvolumesClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSubvolumesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SubvolumesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &SubvolumesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreate - Creates a subvolume in the path or clones the subvolume mentioned in the parentPath
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// accountName - The name of the NetApp account
// poolName - The name of the capacity pool
// volumeName - The name of the volume
// subvolumeName - The name of the subvolume.
// body - Subvolume object supplied in the body of the operation.
// options - SubvolumesClientBeginCreateOptions contains the optional parameters for the SubvolumesClient.BeginCreate method.
func (client *SubvolumesClient) BeginCreate(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, subvolumeName string, body SubvolumeInfo, options *SubvolumesClientBeginCreateOptions) (*armruntime.Poller[SubvolumesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, accountName, poolName, volumeName, subvolumeName, body, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[SubvolumesClientCreateResponse]{
			FinalStateVia: armruntime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[SubvolumesClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Creates a subvolume in the path or clones the subvolume mentioned in the parentPath
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SubvolumesClient) create(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, subvolumeName string, body SubvolumeInfo, options *SubvolumesClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, subvolumeName, body, options)
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
func (client *SubvolumesClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, subvolumeName string, body SubvolumeInfo, options *SubvolumesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/subvolumes/{subvolumeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	if subvolumeName == "" {
		return nil, errors.New("parameter subvolumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subvolumeName}", url.PathEscape(subvolumeName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginDelete - Delete subvolume
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// accountName - The name of the NetApp account
// poolName - The name of the capacity pool
// volumeName - The name of the volume
// subvolumeName - The name of the subvolume.
// options - SubvolumesClientBeginDeleteOptions contains the optional parameters for the SubvolumesClient.BeginDelete method.
func (client *SubvolumesClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, subvolumeName string, options *SubvolumesClientBeginDeleteOptions) (*armruntime.Poller[SubvolumesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, poolName, volumeName, subvolumeName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[SubvolumesClientDeleteResponse]{
			FinalStateVia: armruntime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[SubvolumesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete subvolume
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SubvolumesClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, subvolumeName string, options *SubvolumesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, subvolumeName, options)
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
func (client *SubvolumesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, subvolumeName string, options *SubvolumesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/subvolumes/{subvolumeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	if subvolumeName == "" {
		return nil, errors.New("parameter subvolumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subvolumeName}", url.PathEscape(subvolumeName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Returns the path associated with the subvolumeName provided
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// accountName - The name of the NetApp account
// poolName - The name of the capacity pool
// volumeName - The name of the volume
// subvolumeName - The name of the subvolume.
// options - SubvolumesClientGetOptions contains the optional parameters for the SubvolumesClient.Get method.
func (client *SubvolumesClient) Get(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, subvolumeName string, options *SubvolumesClientGetOptions) (SubvolumesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, subvolumeName, options)
	if err != nil {
		return SubvolumesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SubvolumesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SubvolumesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SubvolumesClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, subvolumeName string, options *SubvolumesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/subvolumes/{subvolumeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	if subvolumeName == "" {
		return nil, errors.New("parameter subvolumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subvolumeName}", url.PathEscape(subvolumeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SubvolumesClient) getHandleResponse(resp *http.Response) (SubvolumesClientGetResponse, error) {
	result := SubvolumesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubvolumeInfo); err != nil {
		return SubvolumesClientGetResponse{}, err
	}
	return result, nil
}

// BeginGetMetadata - Get details of the specified subvolume
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// accountName - The name of the NetApp account
// poolName - The name of the capacity pool
// volumeName - The name of the volume
// subvolumeName - The name of the subvolume.
// options - SubvolumesClientBeginGetMetadataOptions contains the optional parameters for the SubvolumesClient.BeginGetMetadata
// method.
func (client *SubvolumesClient) BeginGetMetadata(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, subvolumeName string, options *SubvolumesClientBeginGetMetadataOptions) (*armruntime.Poller[SubvolumesClientGetMetadataResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.getMetadata(ctx, resourceGroupName, accountName, poolName, volumeName, subvolumeName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[SubvolumesClientGetMetadataResponse]{
			FinalStateVia: armruntime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[SubvolumesClientGetMetadataResponse](options.ResumeToken, client.pl, nil)
	}
}

// GetMetadata - Get details of the specified subvolume
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SubvolumesClient) getMetadata(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, subvolumeName string, options *SubvolumesClientBeginGetMetadataOptions) (*http.Response, error) {
	req, err := client.getMetadataCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, subvolumeName, options)
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

// getMetadataCreateRequest creates the GetMetadata request.
func (client *SubvolumesClient) getMetadataCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, subvolumeName string, options *SubvolumesClientBeginGetMetadataOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/subvolumes/{subvolumeName}/getMetadata"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	if subvolumeName == "" {
		return nil, errors.New("parameter subvolumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subvolumeName}", url.PathEscape(subvolumeName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// ListByVolume - Returns a list of the subvolumes in the volume
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// accountName - The name of the NetApp account
// poolName - The name of the capacity pool
// volumeName - The name of the volume
// options - SubvolumesClientListByVolumeOptions contains the optional parameters for the SubvolumesClient.ListByVolume method.
func (client *SubvolumesClient) ListByVolume(resourceGroupName string, accountName string, poolName string, volumeName string, options *SubvolumesClientListByVolumeOptions) *runtime.Pager[SubvolumesClientListByVolumeResponse] {
	return runtime.NewPager(runtime.PageProcessor[SubvolumesClientListByVolumeResponse]{
		More: func(page SubvolumesClientListByVolumeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SubvolumesClientListByVolumeResponse) (SubvolumesClientListByVolumeResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByVolumeCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SubvolumesClientListByVolumeResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SubvolumesClientListByVolumeResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SubvolumesClientListByVolumeResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByVolumeHandleResponse(resp)
		},
	})
}

// listByVolumeCreateRequest creates the ListByVolume request.
func (client *SubvolumesClient) listByVolumeCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, options *SubvolumesClientListByVolumeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/subvolumes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByVolumeHandleResponse handles the ListByVolume response.
func (client *SubvolumesClient) listByVolumeHandleResponse(resp *http.Response) (SubvolumesClientListByVolumeResponse, error) {
	result := SubvolumesClientListByVolumeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubvolumesList); err != nil {
		return SubvolumesClientListByVolumeResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Patch a subvolume
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// accountName - The name of the NetApp account
// poolName - The name of the capacity pool
// volumeName - The name of the volume
// subvolumeName - The name of the subvolume.
// body - Subvolume object supplied in the body of the operation.
// options - SubvolumesClientBeginUpdateOptions contains the optional parameters for the SubvolumesClient.BeginUpdate method.
func (client *SubvolumesClient) BeginUpdate(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, subvolumeName string, body SubvolumePatchRequest, options *SubvolumesClientBeginUpdateOptions) (*armruntime.Poller[SubvolumesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, accountName, poolName, volumeName, subvolumeName, body, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[SubvolumesClientUpdateResponse]{
			FinalStateVia: armruntime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[SubvolumesClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Patch a subvolume
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SubvolumesClient) update(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, subvolumeName string, body SubvolumePatchRequest, options *SubvolumesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, subvolumeName, body, options)
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
func (client *SubvolumesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, subvolumeName string, body SubvolumePatchRequest, options *SubvolumesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/subvolumes/{subvolumeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	if subvolumeName == "" {
		return nil, errors.New("parameter subvolumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subvolumeName}", url.PathEscape(subvolumeName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}
