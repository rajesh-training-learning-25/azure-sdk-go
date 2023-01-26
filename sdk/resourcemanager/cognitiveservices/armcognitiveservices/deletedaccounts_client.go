//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcognitiveservices

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

// DeletedAccountsClient contains the methods for the DeletedAccounts group.
// Don't use this type directly, use NewDeletedAccountsClient() instead.
type DeletedAccountsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDeletedAccountsClient creates a new instance of DeletedAccountsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDeletedAccountsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DeletedAccountsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &DeletedAccountsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Returns a Cognitive Services account specified by the parameters.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01
// location - Resource location.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The name of Cognitive Services account.
// options - DeletedAccountsClientGetOptions contains the optional parameters for the DeletedAccountsClient.Get method.
func (client *DeletedAccountsClient) Get(ctx context.Context, location string, resourceGroupName string, accountName string, options *DeletedAccountsClientGetOptions) (DeletedAccountsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, location, resourceGroupName, accountName, options)
	if err != nil {
		return DeletedAccountsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DeletedAccountsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DeletedAccountsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DeletedAccountsClient) getCreateRequest(ctx context.Context, location string, resourceGroupName string, accountName string, options *DeletedAccountsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.CognitiveServices/locations/{location}/resourceGroups/{resourceGroupName}/deletedAccounts/{accountName}"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DeletedAccountsClient) getHandleResponse(resp *http.Response) (DeletedAccountsClientGetResponse, error) {
	result := DeletedAccountsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Account); err != nil {
		return DeletedAccountsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Returns all the resources of a particular type belonging to a subscription.
// Generated from API version 2022-10-01
// options - DeletedAccountsClientListOptions contains the optional parameters for the DeletedAccountsClient.List method.
func (client *DeletedAccountsClient) NewListPager(options *DeletedAccountsClientListOptions) *runtime.Pager[DeletedAccountsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DeletedAccountsClientListResponse]{
		More: func(page DeletedAccountsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DeletedAccountsClientListResponse) (DeletedAccountsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DeletedAccountsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DeletedAccountsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DeletedAccountsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *DeletedAccountsClient) listCreateRequest(ctx context.Context, options *DeletedAccountsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.CognitiveServices/deletedAccounts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DeletedAccountsClient) listHandleResponse(resp *http.Response) (DeletedAccountsClientListResponse, error) {
	result := DeletedAccountsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccountListResult); err != nil {
		return DeletedAccountsClientListResponse{}, err
	}
	return result, nil
}

// BeginPurge - Deletes a Cognitive Services account from the resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01
// location - Resource location.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The name of Cognitive Services account.
// options - DeletedAccountsClientBeginPurgeOptions contains the optional parameters for the DeletedAccountsClient.BeginPurge
// method.
func (client *DeletedAccountsClient) BeginPurge(ctx context.Context, location string, resourceGroupName string, accountName string, options *DeletedAccountsClientBeginPurgeOptions) (*runtime.Poller[DeletedAccountsClientPurgeResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.purge(ctx, location, resourceGroupName, accountName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DeletedAccountsClientPurgeResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[DeletedAccountsClientPurgeResponse](options.ResumeToken, client.pl, nil)
	}
}

// Purge - Deletes a Cognitive Services account from the resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01
func (client *DeletedAccountsClient) purge(ctx context.Context, location string, resourceGroupName string, accountName string, options *DeletedAccountsClientBeginPurgeOptions) (*http.Response, error) {
	req, err := client.purgeCreateRequest(ctx, location, resourceGroupName, accountName, options)
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

// purgeCreateRequest creates the Purge request.
func (client *DeletedAccountsClient) purgeCreateRequest(ctx context.Context, location string, resourceGroupName string, accountName string, options *DeletedAccountsClientBeginPurgeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.CognitiveServices/locations/{location}/resourceGroups/{resourceGroupName}/deletedAccounts/{accountName}"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
