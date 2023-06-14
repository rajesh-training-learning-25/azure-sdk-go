//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

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

// SSHPublicKeysClient contains the methods for the SSHPublicKeys group.
// Don't use this type directly, use NewSSHPublicKeysClient() instead.
type SSHPublicKeysClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSSHPublicKeysClient creates a new instance of SSHPublicKeysClient with the specified values.
//   - subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSSHPublicKeysClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SSHPublicKeysClient, error) {
	cl, err := arm.NewClient(moduleName+".SSHPublicKeysClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SSHPublicKeysClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Creates a new SSH public key resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group.
//   - sshPublicKeyName - The name of the SSH public key.
//   - parameters - Parameters supplied to create the SSH public key.
//   - options - SSHPublicKeysClientCreateOptions contains the optional parameters for the SSHPublicKeysClient.Create method.
func (client *SSHPublicKeysClient) Create(ctx context.Context, resourceGroupName string, sshPublicKeyName string, parameters SSHPublicKeyResource, options *SSHPublicKeysClientCreateOptions) (SSHPublicKeysClientCreateResponse, error) {
	var err error
	const operationName = "SSHPublicKeysClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, sshPublicKeyName, parameters, options)
	if err != nil {
		return SSHPublicKeysClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SSHPublicKeysClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return SSHPublicKeysClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *SSHPublicKeysClient) createCreateRequest(ctx context.Context, resourceGroupName string, sshPublicKeyName string, parameters SSHPublicKeyResource, options *SSHPublicKeysClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/sshPublicKeys/{sshPublicKeyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sshPublicKeyName == "" {
		return nil, errors.New("parameter sshPublicKeyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sshPublicKeyName}", url.PathEscape(sshPublicKeyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *SSHPublicKeysClient) createHandleResponse(resp *http.Response) (SSHPublicKeysClientCreateResponse, error) {
	result := SSHPublicKeysClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SSHPublicKeyResource); err != nil {
		return SSHPublicKeysClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Delete an SSH public key.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group.
//   - sshPublicKeyName - The name of the SSH public key.
//   - options - SSHPublicKeysClientDeleteOptions contains the optional parameters for the SSHPublicKeysClient.Delete method.
func (client *SSHPublicKeysClient) Delete(ctx context.Context, resourceGroupName string, sshPublicKeyName string, options *SSHPublicKeysClientDeleteOptions) (SSHPublicKeysClientDeleteResponse, error) {
	var err error
	const operationName = "SSHPublicKeysClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, sshPublicKeyName, options)
	if err != nil {
		return SSHPublicKeysClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SSHPublicKeysClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return SSHPublicKeysClientDeleteResponse{}, err
	}
	return SSHPublicKeysClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SSHPublicKeysClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, sshPublicKeyName string, options *SSHPublicKeysClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/sshPublicKeys/{sshPublicKeyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sshPublicKeyName == "" {
		return nil, errors.New("parameter sshPublicKeyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sshPublicKeyName}", url.PathEscape(sshPublicKeyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GenerateKeyPair - Generates and returns a public/private key pair and populates the SSH public key resource with the public
// key. The length of the key will be 3072 bits. This operation can only be performed once per
// SSH public key resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group.
//   - sshPublicKeyName - The name of the SSH public key.
//   - options - SSHPublicKeysClientGenerateKeyPairOptions contains the optional parameters for the SSHPublicKeysClient.GenerateKeyPair
//     method.
func (client *SSHPublicKeysClient) GenerateKeyPair(ctx context.Context, resourceGroupName string, sshPublicKeyName string, options *SSHPublicKeysClientGenerateKeyPairOptions) (SSHPublicKeysClientGenerateKeyPairResponse, error) {
	var err error
	const operationName = "SSHPublicKeysClient.GenerateKeyPair"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.generateKeyPairCreateRequest(ctx, resourceGroupName, sshPublicKeyName, options)
	if err != nil {
		return SSHPublicKeysClientGenerateKeyPairResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SSHPublicKeysClientGenerateKeyPairResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SSHPublicKeysClientGenerateKeyPairResponse{}, err
	}
	resp, err := client.generateKeyPairHandleResponse(httpResp)
	return resp, err
}

// generateKeyPairCreateRequest creates the GenerateKeyPair request.
func (client *SSHPublicKeysClient) generateKeyPairCreateRequest(ctx context.Context, resourceGroupName string, sshPublicKeyName string, options *SSHPublicKeysClientGenerateKeyPairOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/sshPublicKeys/{sshPublicKeyName}/generateKeyPair"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sshPublicKeyName == "" {
		return nil, errors.New("parameter sshPublicKeyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sshPublicKeyName}", url.PathEscape(sshPublicKeyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// generateKeyPairHandleResponse handles the GenerateKeyPair response.
func (client *SSHPublicKeysClient) generateKeyPairHandleResponse(resp *http.Response) (SSHPublicKeysClientGenerateKeyPairResponse, error) {
	result := SSHPublicKeysClientGenerateKeyPairResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SSHPublicKeyGenerateKeyPairResult); err != nil {
		return SSHPublicKeysClientGenerateKeyPairResponse{}, err
	}
	return result, nil
}

// Get - Retrieves information about an SSH public key.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group.
//   - sshPublicKeyName - The name of the SSH public key.
//   - options - SSHPublicKeysClientGetOptions contains the optional parameters for the SSHPublicKeysClient.Get method.
func (client *SSHPublicKeysClient) Get(ctx context.Context, resourceGroupName string, sshPublicKeyName string, options *SSHPublicKeysClientGetOptions) (SSHPublicKeysClientGetResponse, error) {
	var err error
	const operationName = "SSHPublicKeysClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, sshPublicKeyName, options)
	if err != nil {
		return SSHPublicKeysClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SSHPublicKeysClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SSHPublicKeysClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SSHPublicKeysClient) getCreateRequest(ctx context.Context, resourceGroupName string, sshPublicKeyName string, options *SSHPublicKeysClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/sshPublicKeys/{sshPublicKeyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sshPublicKeyName == "" {
		return nil, errors.New("parameter sshPublicKeyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sshPublicKeyName}", url.PathEscape(sshPublicKeyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SSHPublicKeysClient) getHandleResponse(resp *http.Response) (SSHPublicKeysClientGetResponse, error) {
	result := SSHPublicKeysClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SSHPublicKeyResource); err != nil {
		return SSHPublicKeysClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists all of the SSH public keys in the specified resource group. Use the nextLink property
// in the response to get the next page of SSH public keys.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group.
//   - options - SSHPublicKeysClientListByResourceGroupOptions contains the optional parameters for the SSHPublicKeysClient.NewListByResourceGroupPager
//     method.
func (client *SSHPublicKeysClient) NewListByResourceGroupPager(resourceGroupName string, options *SSHPublicKeysClientListByResourceGroupOptions) *runtime.Pager[SSHPublicKeysClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[SSHPublicKeysClientListByResourceGroupResponse]{
		More: func(page SSHPublicKeysClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SSHPublicKeysClientListByResourceGroupResponse) (SSHPublicKeysClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SSHPublicKeysClient.NewListByResourceGroupPager")
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SSHPublicKeysClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return SSHPublicKeysClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SSHPublicKeysClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *SSHPublicKeysClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *SSHPublicKeysClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/sshPublicKeys"
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
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *SSHPublicKeysClient) listByResourceGroupHandleResponse(resp *http.Response) (SSHPublicKeysClientListByResourceGroupResponse, error) {
	result := SSHPublicKeysClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SSHPublicKeysGroupListResult); err != nil {
		return SSHPublicKeysClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Lists all of the SSH public keys in the subscription. Use the nextLink property in the response
// to get the next page of SSH public keys.
//
// Generated from API version 2023-03-01
//   - options - SSHPublicKeysClientListBySubscriptionOptions contains the optional parameters for the SSHPublicKeysClient.NewListBySubscriptionPager
//     method.
func (client *SSHPublicKeysClient) NewListBySubscriptionPager(options *SSHPublicKeysClientListBySubscriptionOptions) *runtime.Pager[SSHPublicKeysClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[SSHPublicKeysClientListBySubscriptionResponse]{
		More: func(page SSHPublicKeysClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SSHPublicKeysClientListBySubscriptionResponse) (SSHPublicKeysClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SSHPublicKeysClient.NewListBySubscriptionPager")
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SSHPublicKeysClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return SSHPublicKeysClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SSHPublicKeysClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *SSHPublicKeysClient) listBySubscriptionCreateRequest(ctx context.Context, options *SSHPublicKeysClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/sshPublicKeys"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *SSHPublicKeysClient) listBySubscriptionHandleResponse(resp *http.Response) (SSHPublicKeysClientListBySubscriptionResponse, error) {
	result := SSHPublicKeysClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SSHPublicKeysGroupListResult); err != nil {
		return SSHPublicKeysClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Updates a new SSH public key resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group.
//   - sshPublicKeyName - The name of the SSH public key.
//   - parameters - Parameters supplied to update the SSH public key.
//   - options - SSHPublicKeysClientUpdateOptions contains the optional parameters for the SSHPublicKeysClient.Update method.
func (client *SSHPublicKeysClient) Update(ctx context.Context, resourceGroupName string, sshPublicKeyName string, parameters SSHPublicKeyUpdateResource, options *SSHPublicKeysClientUpdateOptions) (SSHPublicKeysClientUpdateResponse, error) {
	var err error
	const operationName = "SSHPublicKeysClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, sshPublicKeyName, parameters, options)
	if err != nil {
		return SSHPublicKeysClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SSHPublicKeysClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SSHPublicKeysClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *SSHPublicKeysClient) updateCreateRequest(ctx context.Context, resourceGroupName string, sshPublicKeyName string, parameters SSHPublicKeyUpdateResource, options *SSHPublicKeysClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/sshPublicKeys/{sshPublicKeyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sshPublicKeyName == "" {
		return nil, errors.New("parameter sshPublicKeyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sshPublicKeyName}", url.PathEscape(sshPublicKeyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *SSHPublicKeysClient) updateHandleResponse(resp *http.Response) (SSHPublicKeysClientUpdateResponse, error) {
	result := SSHPublicKeysClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SSHPublicKeyResource); err != nil {
		return SSHPublicKeysClientUpdateResponse{}, err
	}
	return result, nil
}
