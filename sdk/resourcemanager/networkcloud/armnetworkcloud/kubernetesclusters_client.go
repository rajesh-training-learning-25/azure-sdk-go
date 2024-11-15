//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetworkcloud

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

// KubernetesClustersClient contains the methods for the KubernetesClusters group.
// Don't use this type directly, use NewKubernetesClustersClient() instead.
type KubernetesClustersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewKubernetesClustersClient creates a new instance of KubernetesClustersClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewKubernetesClustersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*KubernetesClustersClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &KubernetesClustersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a new Kubernetes cluster or update the properties of the existing one.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - kubernetesClusterName - The name of the Kubernetes cluster.
//   - kubernetesClusterParameters - The request body.
//   - options - KubernetesClustersClientBeginCreateOrUpdateOptions contains the optional parameters for the KubernetesClustersClient.BeginCreateOrUpdate
//     method.
func (client *KubernetesClustersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, kubernetesClusterName string, kubernetesClusterParameters KubernetesCluster, options *KubernetesClustersClientBeginCreateOrUpdateOptions) (*runtime.Poller[KubernetesClustersClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, kubernetesClusterName, kubernetesClusterParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[KubernetesClustersClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[KubernetesClustersClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a new Kubernetes cluster or update the properties of the existing one.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
func (client *KubernetesClustersClient) createOrUpdate(ctx context.Context, resourceGroupName string, kubernetesClusterName string, kubernetesClusterParameters KubernetesCluster, options *KubernetesClustersClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "KubernetesClustersClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, kubernetesClusterName, kubernetesClusterParameters, options)
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
func (client *KubernetesClustersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, kubernetesClusterName string, kubernetesClusterParameters KubernetesCluster, options *KubernetesClustersClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/kubernetesClusters/{kubernetesClusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if kubernetesClusterName == "" {
		return nil, errors.New("parameter kubernetesClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kubernetesClusterName}", url.PathEscape(kubernetesClusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, kubernetesClusterParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete the provided Kubernetes cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - kubernetesClusterName - The name of the Kubernetes cluster.
//   - options - KubernetesClustersClientBeginDeleteOptions contains the optional parameters for the KubernetesClustersClient.BeginDelete
//     method.
func (client *KubernetesClustersClient) BeginDelete(ctx context.Context, resourceGroupName string, kubernetesClusterName string, options *KubernetesClustersClientBeginDeleteOptions) (*runtime.Poller[KubernetesClustersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, kubernetesClusterName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[KubernetesClustersClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[KubernetesClustersClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete the provided Kubernetes cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
func (client *KubernetesClustersClient) deleteOperation(ctx context.Context, resourceGroupName string, kubernetesClusterName string, options *KubernetesClustersClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "KubernetesClustersClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, kubernetesClusterName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *KubernetesClustersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, kubernetesClusterName string, options *KubernetesClustersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/kubernetesClusters/{kubernetesClusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if kubernetesClusterName == "" {
		return nil, errors.New("parameter kubernetesClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kubernetesClusterName}", url.PathEscape(kubernetesClusterName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get properties of the provided the Kubernetes cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - kubernetesClusterName - The name of the Kubernetes cluster.
//   - options - KubernetesClustersClientGetOptions contains the optional parameters for the KubernetesClustersClient.Get method.
func (client *KubernetesClustersClient) Get(ctx context.Context, resourceGroupName string, kubernetesClusterName string, options *KubernetesClustersClientGetOptions) (KubernetesClustersClientGetResponse, error) {
	var err error
	const operationName = "KubernetesClustersClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, kubernetesClusterName, options)
	if err != nil {
		return KubernetesClustersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return KubernetesClustersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return KubernetesClustersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *KubernetesClustersClient) getCreateRequest(ctx context.Context, resourceGroupName string, kubernetesClusterName string, options *KubernetesClustersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/kubernetesClusters/{kubernetesClusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if kubernetesClusterName == "" {
		return nil, errors.New("parameter kubernetesClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kubernetesClusterName}", url.PathEscape(kubernetesClusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *KubernetesClustersClient) getHandleResponse(resp *http.Response) (KubernetesClustersClientGetResponse, error) {
	result := KubernetesClustersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.KubernetesCluster); err != nil {
		return KubernetesClustersClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Get a list of Kubernetes clusters in the provided resource group.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - KubernetesClustersClientListByResourceGroupOptions contains the optional parameters for the KubernetesClustersClient.NewListByResourceGroupPager
//     method.
func (client *KubernetesClustersClient) NewListByResourceGroupPager(resourceGroupName string, options *KubernetesClustersClientListByResourceGroupOptions) *runtime.Pager[KubernetesClustersClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[KubernetesClustersClientListByResourceGroupResponse]{
		More: func(page KubernetesClustersClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *KubernetesClustersClientListByResourceGroupResponse) (KubernetesClustersClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "KubernetesClustersClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return KubernetesClustersClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *KubernetesClustersClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *KubernetesClustersClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/kubernetesClusters"
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
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *KubernetesClustersClient) listByResourceGroupHandleResponse(resp *http.Response) (KubernetesClustersClientListByResourceGroupResponse, error) {
	result := KubernetesClustersClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.KubernetesClusterList); err != nil {
		return KubernetesClustersClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Get a list of Kubernetes clusters in the provided subscription.
//
// Generated from API version 2024-06-01-preview
//   - options - KubernetesClustersClientListBySubscriptionOptions contains the optional parameters for the KubernetesClustersClient.NewListBySubscriptionPager
//     method.
func (client *KubernetesClustersClient) NewListBySubscriptionPager(options *KubernetesClustersClientListBySubscriptionOptions) *runtime.Pager[KubernetesClustersClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[KubernetesClustersClientListBySubscriptionResponse]{
		More: func(page KubernetesClustersClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *KubernetesClustersClientListBySubscriptionResponse) (KubernetesClustersClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "KubernetesClustersClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return KubernetesClustersClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *KubernetesClustersClient) listBySubscriptionCreateRequest(ctx context.Context, options *KubernetesClustersClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.NetworkCloud/kubernetesClusters"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *KubernetesClustersClient) listBySubscriptionHandleResponse(resp *http.Response) (KubernetesClustersClientListBySubscriptionResponse, error) {
	result := KubernetesClustersClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.KubernetesClusterList); err != nil {
		return KubernetesClustersClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginRestartNode - Restart a targeted node of a Kubernetes cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - kubernetesClusterName - The name of the Kubernetes cluster.
//   - kubernetesClusterRestartNodeParameters - The request body.
//   - options - KubernetesClustersClientBeginRestartNodeOptions contains the optional parameters for the KubernetesClustersClient.BeginRestartNode
//     method.
func (client *KubernetesClustersClient) BeginRestartNode(ctx context.Context, resourceGroupName string, kubernetesClusterName string, kubernetesClusterRestartNodeParameters KubernetesClusterRestartNodeParameters, options *KubernetesClustersClientBeginRestartNodeOptions) (*runtime.Poller[KubernetesClustersClientRestartNodeResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.restartNode(ctx, resourceGroupName, kubernetesClusterName, kubernetesClusterRestartNodeParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[KubernetesClustersClientRestartNodeResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[KubernetesClustersClientRestartNodeResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// RestartNode - Restart a targeted node of a Kubernetes cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
func (client *KubernetesClustersClient) restartNode(ctx context.Context, resourceGroupName string, kubernetesClusterName string, kubernetesClusterRestartNodeParameters KubernetesClusterRestartNodeParameters, options *KubernetesClustersClientBeginRestartNodeOptions) (*http.Response, error) {
	var err error
	const operationName = "KubernetesClustersClient.BeginRestartNode"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.restartNodeCreateRequest(ctx, resourceGroupName, kubernetesClusterName, kubernetesClusterRestartNodeParameters, options)
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

// restartNodeCreateRequest creates the RestartNode request.
func (client *KubernetesClustersClient) restartNodeCreateRequest(ctx context.Context, resourceGroupName string, kubernetesClusterName string, kubernetesClusterRestartNodeParameters KubernetesClusterRestartNodeParameters, options *KubernetesClustersClientBeginRestartNodeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/kubernetesClusters/{kubernetesClusterName}/restartNode"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if kubernetesClusterName == "" {
		return nil, errors.New("parameter kubernetesClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kubernetesClusterName}", url.PathEscape(kubernetesClusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, kubernetesClusterRestartNodeParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginUpdate - Patch the properties of the provided Kubernetes cluster, or update the tags associated with the Kubernetes
// cluster. Properties and tag updates can be done independently.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - kubernetesClusterName - The name of the Kubernetes cluster.
//   - kubernetesClusterUpdateParameters - The request body.
//   - options - KubernetesClustersClientBeginUpdateOptions contains the optional parameters for the KubernetesClustersClient.BeginUpdate
//     method.
func (client *KubernetesClustersClient) BeginUpdate(ctx context.Context, resourceGroupName string, kubernetesClusterName string, kubernetesClusterUpdateParameters KubernetesClusterPatchParameters, options *KubernetesClustersClientBeginUpdateOptions) (*runtime.Poller[KubernetesClustersClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, kubernetesClusterName, kubernetesClusterUpdateParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[KubernetesClustersClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[KubernetesClustersClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Patch the properties of the provided Kubernetes cluster, or update the tags associated with the Kubernetes cluster.
// Properties and tag updates can be done independently.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
func (client *KubernetesClustersClient) update(ctx context.Context, resourceGroupName string, kubernetesClusterName string, kubernetesClusterUpdateParameters KubernetesClusterPatchParameters, options *KubernetesClustersClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "KubernetesClustersClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, kubernetesClusterName, kubernetesClusterUpdateParameters, options)
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
func (client *KubernetesClustersClient) updateCreateRequest(ctx context.Context, resourceGroupName string, kubernetesClusterName string, kubernetesClusterUpdateParameters KubernetesClusterPatchParameters, options *KubernetesClustersClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/kubernetesClusters/{kubernetesClusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if kubernetesClusterName == "" {
		return nil, errors.New("parameter kubernetesClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kubernetesClusterName}", url.PathEscape(kubernetesClusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, kubernetesClusterUpdateParameters); err != nil {
		return nil, err
	}
	return req, nil
}
