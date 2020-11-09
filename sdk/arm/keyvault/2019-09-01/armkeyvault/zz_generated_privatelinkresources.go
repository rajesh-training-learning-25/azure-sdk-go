// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkeyvault

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// PrivateLinkResourcesOperations contains the methods for the PrivateLinkResources group.
type PrivateLinkResourcesOperations interface {
	// ListByVault - Gets the private link resources supported for the key vault.
	ListByVault(ctx context.Context, resourceGroupName string, vaultName string, options *PrivateLinkResourcesListByVaultOptions) (*PrivateLinkResourceListResultResponse, error)
}

// PrivateLinkResourcesClient implements the PrivateLinkResourcesOperations interface.
// Don't use this type directly, use NewPrivateLinkResourcesClient() instead.
type PrivateLinkResourcesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewPrivateLinkResourcesClient creates a new instance of PrivateLinkResourcesClient with the specified values.
func NewPrivateLinkResourcesClient(con *armcore.Connection, subscriptionID string) PrivateLinkResourcesOperations {
	return &PrivateLinkResourcesClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *PrivateLinkResourcesClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// ListByVault - Gets the private link resources supported for the key vault.
func (client *PrivateLinkResourcesClient) ListByVault(ctx context.Context, resourceGroupName string, vaultName string, options *PrivateLinkResourcesListByVaultOptions) (*PrivateLinkResourceListResultResponse, error) {
	req, err := client.ListByVaultCreateRequest(ctx, resourceGroupName, vaultName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListByVaultHandleError(resp)
	}
	result, err := client.ListByVaultHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListByVaultCreateRequest creates the ListByVault request.
func (client *PrivateLinkResourcesClient) ListByVaultCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, options *PrivateLinkResourcesListByVaultOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/privateLinkResources"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-09-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListByVaultHandleResponse handles the ListByVault response.
func (client *PrivateLinkResourcesClient) ListByVaultHandleResponse(resp *azcore.Response) (*PrivateLinkResourceListResultResponse, error) {
	result := PrivateLinkResourceListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PrivateLinkResourceListResult)
}

// ListByVaultHandleError handles the ListByVault error response.
func (client *PrivateLinkResourcesClient) ListByVaultHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
