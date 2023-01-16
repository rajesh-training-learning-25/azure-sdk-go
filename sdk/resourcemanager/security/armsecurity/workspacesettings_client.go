//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurity

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

// WorkspaceSettingsClient contains the methods for the WorkspaceSettings group.
// Don't use this type directly, use NewWorkspaceSettingsClient() instead.
type WorkspaceSettingsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewWorkspaceSettingsClient creates a new instance of WorkspaceSettingsClient with the specified values.
// subscriptionID - Azure subscription ID
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewWorkspaceSettingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkspaceSettingsClient, error) {
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
	client := &WorkspaceSettingsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Create - creating settings about where we should store your security data and logs
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-08-01-preview
// workspaceSettingName - Name of the security setting
// workspaceSetting - Security data setting object
// options - WorkspaceSettingsClientCreateOptions contains the optional parameters for the WorkspaceSettingsClient.Create
// method.
func (client *WorkspaceSettingsClient) Create(ctx context.Context, workspaceSettingName string, workspaceSetting WorkspaceSetting, options *WorkspaceSettingsClientCreateOptions) (WorkspaceSettingsClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, workspaceSettingName, workspaceSetting, options)
	if err != nil {
		return WorkspaceSettingsClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspaceSettingsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspaceSettingsClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *WorkspaceSettingsClient) createCreateRequest(ctx context.Context, workspaceSettingName string, workspaceSetting WorkspaceSetting, options *WorkspaceSettingsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/workspaceSettings/{workspaceSettingName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceSettingName == "" {
		return nil, errors.New("parameter workspaceSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceSettingName}", url.PathEscape(workspaceSettingName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, workspaceSetting)
}

// createHandleResponse handles the Create response.
func (client *WorkspaceSettingsClient) createHandleResponse(resp *http.Response) (WorkspaceSettingsClientCreateResponse, error) {
	result := WorkspaceSettingsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceSetting); err != nil {
		return WorkspaceSettingsClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the custom workspace settings for this subscription. new VMs will report to the default workspace
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-08-01-preview
// workspaceSettingName - Name of the security setting
// options - WorkspaceSettingsClientDeleteOptions contains the optional parameters for the WorkspaceSettingsClient.Delete
// method.
func (client *WorkspaceSettingsClient) Delete(ctx context.Context, workspaceSettingName string, options *WorkspaceSettingsClientDeleteOptions) (WorkspaceSettingsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, workspaceSettingName, options)
	if err != nil {
		return WorkspaceSettingsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspaceSettingsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return WorkspaceSettingsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return WorkspaceSettingsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WorkspaceSettingsClient) deleteCreateRequest(ctx context.Context, workspaceSettingName string, options *WorkspaceSettingsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/workspaceSettings/{workspaceSettingName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceSettingName == "" {
		return nil, errors.New("parameter workspaceSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceSettingName}", url.PathEscape(workspaceSettingName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Settings about where we should store your security data and logs. If the result is empty, it means that no custom-workspace
// configuration was set
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-08-01-preview
// workspaceSettingName - Name of the security setting
// options - WorkspaceSettingsClientGetOptions contains the optional parameters for the WorkspaceSettingsClient.Get method.
func (client *WorkspaceSettingsClient) Get(ctx context.Context, workspaceSettingName string, options *WorkspaceSettingsClientGetOptions) (WorkspaceSettingsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, workspaceSettingName, options)
	if err != nil {
		return WorkspaceSettingsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspaceSettingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspaceSettingsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WorkspaceSettingsClient) getCreateRequest(ctx context.Context, workspaceSettingName string, options *WorkspaceSettingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/workspaceSettings/{workspaceSettingName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceSettingName == "" {
		return nil, errors.New("parameter workspaceSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceSettingName}", url.PathEscape(workspaceSettingName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkspaceSettingsClient) getHandleResponse(resp *http.Response) (WorkspaceSettingsClientGetResponse, error) {
	result := WorkspaceSettingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceSetting); err != nil {
		return WorkspaceSettingsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Settings about where we should store your security data and logs. If the result is empty, it means that
// no custom-workspace configuration was set
// Generated from API version 2017-08-01-preview
// options - WorkspaceSettingsClientListOptions contains the optional parameters for the WorkspaceSettingsClient.List method.
func (client *WorkspaceSettingsClient) NewListPager(options *WorkspaceSettingsClientListOptions) *runtime.Pager[WorkspaceSettingsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[WorkspaceSettingsClientListResponse]{
		More: func(page WorkspaceSettingsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WorkspaceSettingsClientListResponse) (WorkspaceSettingsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return WorkspaceSettingsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return WorkspaceSettingsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WorkspaceSettingsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *WorkspaceSettingsClient) listCreateRequest(ctx context.Context, options *WorkspaceSettingsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/workspaceSettings"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *WorkspaceSettingsClient) listHandleResponse(resp *http.Response) (WorkspaceSettingsClientListResponse, error) {
	result := WorkspaceSettingsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceSettingList); err != nil {
		return WorkspaceSettingsClientListResponse{}, err
	}
	return result, nil
}

// Update - Settings about where we should store your security data and logs
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-08-01-preview
// workspaceSettingName - Name of the security setting
// workspaceSetting - Security data setting object
// options - WorkspaceSettingsClientUpdateOptions contains the optional parameters for the WorkspaceSettingsClient.Update
// method.
func (client *WorkspaceSettingsClient) Update(ctx context.Context, workspaceSettingName string, workspaceSetting WorkspaceSetting, options *WorkspaceSettingsClientUpdateOptions) (WorkspaceSettingsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, workspaceSettingName, workspaceSetting, options)
	if err != nil {
		return WorkspaceSettingsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspaceSettingsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspaceSettingsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *WorkspaceSettingsClient) updateCreateRequest(ctx context.Context, workspaceSettingName string, workspaceSetting WorkspaceSetting, options *WorkspaceSettingsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/workspaceSettings/{workspaceSettingName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceSettingName == "" {
		return nil, errors.New("parameter workspaceSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceSettingName}", url.PathEscape(workspaceSettingName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, workspaceSetting)
}

// updateHandleResponse handles the Update response.
func (client *WorkspaceSettingsClient) updateHandleResponse(resp *http.Response) (WorkspaceSettingsClientUpdateResponse, error) {
	result := WorkspaceSettingsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceSetting); err != nil {
		return WorkspaceSettingsClientUpdateResponse{}, err
	}
	return result, nil
}
