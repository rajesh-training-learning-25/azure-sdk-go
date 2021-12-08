//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicessiterecovery

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

// RegisteredIdentitiesClient contains the methods for the RegisteredIdentities group.
// Don't use this type directly, use NewRegisteredIdentitiesClient() instead.
type RegisteredIdentitiesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewRegisteredIdentitiesClient creates a new instance of RegisteredIdentitiesClient with the specified values.
func NewRegisteredIdentitiesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *RegisteredIdentitiesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &RegisteredIdentitiesClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Delete - Unregisters the given container from your Recovery Services vault.
// If the operation fails it returns a generic error.
func (client *RegisteredIdentitiesClient) Delete(ctx context.Context, resourceGroupName string, vaultName string, identityName string, options *RegisteredIdentitiesDeleteOptions) (RegisteredIdentitiesDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, vaultName, identityName, options)
	if err != nil {
		return RegisteredIdentitiesDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RegisteredIdentitiesDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return RegisteredIdentitiesDeleteResponse{}, client.deleteHandleError(resp)
	}
	return RegisteredIdentitiesDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RegisteredIdentitiesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, identityName string, options *RegisteredIdentitiesDeleteOptions) (*policy.Request, error) {
	urlPath := "/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/registeredIdentities/{identityName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if identityName == "" {
		return nil, errors.New("parameter identityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{identityName}", url.PathEscape(identityName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *RegisteredIdentitiesClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
