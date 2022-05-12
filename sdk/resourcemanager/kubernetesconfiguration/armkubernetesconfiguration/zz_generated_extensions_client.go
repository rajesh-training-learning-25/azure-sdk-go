//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkubernetesconfiguration

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
	"strconv"
	"strings"
)

// ExtensionsClient contains the methods for the Extensions group.
// Don't use this type directly, use NewExtensionsClient() instead.
type ExtensionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewExtensionsClient creates a new instance of ExtensionsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewExtensionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ExtensionsClient, error) {
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
	client := &ExtensionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreate - Create a new Kubernetes Cluster Extension.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// clusterRp - The Kubernetes cluster RP - i.e. Microsoft.ContainerService, Microsoft.Kubernetes, Microsoft.HybridContainerService.
// clusterResourceName - The Kubernetes cluster resource name - i.e. managedClusters, connectedClusters, provisionedClusters.
// clusterName - The name of the kubernetes cluster.
// extensionName - Name of the Extension.
// extension - Properties necessary to Create an Extension.
// options - ExtensionsClientBeginCreateOptions contains the optional parameters for the ExtensionsClient.BeginCreate method.
func (client *ExtensionsClient) BeginCreate(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, extensionName string, extension Extension, options *ExtensionsClientBeginCreateOptions) (*armruntime.Poller[ExtensionsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName, extensionName, extension, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[ExtensionsClientCreateResponse]{
			FinalStateVia: armruntime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[ExtensionsClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Create a new Kubernetes Cluster Extension.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ExtensionsClient) create(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, extensionName string, extension Extension, options *ExtensionsClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName, extensionName, extension, options)
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

// createCreateRequest creates the Create request.
func (client *ExtensionsClient) createCreateRequest(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, extensionName string, extension Extension, options *ExtensionsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/extensions/{extensionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterRp == "" {
		return nil, errors.New("parameter clusterRp cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterRp}", url.PathEscape(clusterRp))
	if clusterResourceName == "" {
		return nil, errors.New("parameter clusterResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterResourceName}", url.PathEscape(clusterResourceName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if extensionName == "" {
		return nil, errors.New("parameter extensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionName}", url.PathEscape(extensionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, extension)
}

// BeginDelete - Delete a Kubernetes Cluster Extension. This will cause the Agent to Uninstall the extension from the cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// clusterRp - The Kubernetes cluster RP - i.e. Microsoft.ContainerService, Microsoft.Kubernetes, Microsoft.HybridContainerService.
// clusterResourceName - The Kubernetes cluster resource name - i.e. managedClusters, connectedClusters, provisionedClusters.
// clusterName - The name of the kubernetes cluster.
// extensionName - Name of the Extension.
// options - ExtensionsClientBeginDeleteOptions contains the optional parameters for the ExtensionsClient.BeginDelete method.
func (client *ExtensionsClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, extensionName string, options *ExtensionsClientBeginDeleteOptions) (*armruntime.Poller[ExtensionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName, extensionName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[ExtensionsClientDeleteResponse]{
			FinalStateVia: armruntime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[ExtensionsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete a Kubernetes Cluster Extension. This will cause the Agent to Uninstall the extension from the cluster.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ExtensionsClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, extensionName string, options *ExtensionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName, extensionName, options)
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
func (client *ExtensionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, extensionName string, options *ExtensionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/extensions/{extensionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterRp == "" {
		return nil, errors.New("parameter clusterRp cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterRp}", url.PathEscape(clusterRp))
	if clusterResourceName == "" {
		return nil, errors.New("parameter clusterResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterResourceName}", url.PathEscape(clusterResourceName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if extensionName == "" {
		return nil, errors.New("parameter extensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionName}", url.PathEscape(extensionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	if options != nil && options.ForceDelete != nil {
		reqQP.Set("forceDelete", strconv.FormatBool(*options.ForceDelete))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets Kubernetes Cluster Extension.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// clusterRp - The Kubernetes cluster RP - i.e. Microsoft.ContainerService, Microsoft.Kubernetes, Microsoft.HybridContainerService.
// clusterResourceName - The Kubernetes cluster resource name - i.e. managedClusters, connectedClusters, provisionedClusters.
// clusterName - The name of the kubernetes cluster.
// extensionName - Name of the Extension.
// options - ExtensionsClientGetOptions contains the optional parameters for the ExtensionsClient.Get method.
func (client *ExtensionsClient) Get(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, extensionName string, options *ExtensionsClientGetOptions) (ExtensionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName, extensionName, options)
	if err != nil {
		return ExtensionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExtensionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExtensionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExtensionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, extensionName string, options *ExtensionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/extensions/{extensionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterRp == "" {
		return nil, errors.New("parameter clusterRp cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterRp}", url.PathEscape(clusterRp))
	if clusterResourceName == "" {
		return nil, errors.New("parameter clusterResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterResourceName}", url.PathEscape(clusterResourceName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if extensionName == "" {
		return nil, errors.New("parameter extensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionName}", url.PathEscape(extensionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExtensionsClient) getHandleResponse(resp *http.Response) (ExtensionsClientGetResponse, error) {
	result := ExtensionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Extension); err != nil {
		return ExtensionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List all Extensions in the cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// clusterRp - The Kubernetes cluster RP - i.e. Microsoft.ContainerService, Microsoft.Kubernetes, Microsoft.HybridContainerService.
// clusterResourceName - The Kubernetes cluster resource name - i.e. managedClusters, connectedClusters, provisionedClusters.
// clusterName - The name of the kubernetes cluster.
// options - ExtensionsClientListOptions contains the optional parameters for the ExtensionsClient.List method.
func (client *ExtensionsClient) NewListPager(resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, options *ExtensionsClientListOptions) *runtime.Pager[ExtensionsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[ExtensionsClientListResponse]{
		More: func(page ExtensionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ExtensionsClientListResponse) (ExtensionsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ExtensionsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ExtensionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ExtensionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ExtensionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, options *ExtensionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/extensions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterRp == "" {
		return nil, errors.New("parameter clusterRp cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterRp}", url.PathEscape(clusterRp))
	if clusterResourceName == "" {
		return nil, errors.New("parameter clusterResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterResourceName}", url.PathEscape(clusterResourceName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ExtensionsClient) listHandleResponse(resp *http.Response) (ExtensionsClientListResponse, error) {
	result := ExtensionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtensionsList); err != nil {
		return ExtensionsClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Patch an existing Kubernetes Cluster Extension.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// clusterRp - The Kubernetes cluster RP - i.e. Microsoft.ContainerService, Microsoft.Kubernetes, Microsoft.HybridContainerService.
// clusterResourceName - The Kubernetes cluster resource name - i.e. managedClusters, connectedClusters, provisionedClusters.
// clusterName - The name of the kubernetes cluster.
// extensionName - Name of the Extension.
// patchExtension - Properties to Patch in an existing Extension.
// options - ExtensionsClientBeginUpdateOptions contains the optional parameters for the ExtensionsClient.BeginUpdate method.
func (client *ExtensionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, extensionName string, patchExtension PatchExtension, options *ExtensionsClientBeginUpdateOptions) (*armruntime.Poller[ExtensionsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName, extensionName, patchExtension, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[ExtensionsClientUpdateResponse]{
			FinalStateVia: armruntime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[ExtensionsClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Patch an existing Kubernetes Cluster Extension.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ExtensionsClient) update(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, extensionName string, patchExtension PatchExtension, options *ExtensionsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName, extensionName, patchExtension, options)
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
func (client *ExtensionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, extensionName string, patchExtension PatchExtension, options *ExtensionsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/extensions/{extensionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterRp == "" {
		return nil, errors.New("parameter clusterRp cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterRp}", url.PathEscape(clusterRp))
	if clusterResourceName == "" {
		return nil, errors.New("parameter clusterResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterResourceName}", url.PathEscape(clusterResourceName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if extensionName == "" {
		return nil, errors.New("parameter extensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionName}", url.PathEscape(extensionName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, patchExtension)
}
