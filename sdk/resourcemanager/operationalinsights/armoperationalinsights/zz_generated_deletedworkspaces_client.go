//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoperationalinsights

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DeletedWorkspacesClient contains the methods for the DeletedWorkspaces group.
// Don't use this type directly, use NewDeletedWorkspacesClient() instead.
type DeletedWorkspacesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDeletedWorkspacesClient creates a new instance of DeletedWorkspacesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDeletedWorkspacesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *DeletedWorkspacesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &DeletedWorkspacesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// List - Gets recently deleted workspaces in a subscription, available for recovery.
// If the operation fails it returns an *azcore.ResponseError type.
// options - DeletedWorkspacesClientListOptions contains the optional parameters for the DeletedWorkspacesClient.List method.
func (client *DeletedWorkspacesClient) List(ctx context.Context, options *DeletedWorkspacesClientListOptions) (DeletedWorkspacesClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return DeletedWorkspacesClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DeletedWorkspacesClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DeletedWorkspacesClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *DeletedWorkspacesClient) listCreateRequest(ctx context.Context, options *DeletedWorkspacesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.OperationalInsights/deletedWorkspaces"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DeletedWorkspacesClient) listHandleResponse(resp *http.Response) (DeletedWorkspacesClientListResponse, error) {
	result := DeletedWorkspacesClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceListResult); err != nil {
		return DeletedWorkspacesClientListResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Gets recently deleted workspaces in a resource group, available for recovery.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - DeletedWorkspacesClientListByResourceGroupOptions contains the optional parameters for the DeletedWorkspacesClient.ListByResourceGroup
// method.
func (client *DeletedWorkspacesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, options *DeletedWorkspacesClientListByResourceGroupOptions) (DeletedWorkspacesClientListByResourceGroupResponse, error) {
	req, err := client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
	if err != nil {
		return DeletedWorkspacesClientListByResourceGroupResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DeletedWorkspacesClientListByResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DeletedWorkspacesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
	}
	return client.listByResourceGroupHandleResponse(resp)
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DeletedWorkspacesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DeletedWorkspacesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/deletedWorkspaces"
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
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DeletedWorkspacesClient) listByResourceGroupHandleResponse(resp *http.Response) (DeletedWorkspacesClientListByResourceGroupResponse, error) {
	result := DeletedWorkspacesClientListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceListResult); err != nil {
		return DeletedWorkspacesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}
