//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdevcenter

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

// ProjectsClient contains the methods for the Projects group.
// Don't use this type directly, use NewProjectsClient() instead.
type ProjectsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewProjectsClient creates a new instance of ProjectsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewProjectsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ProjectsClient, error) {
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
	client := &ProjectsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a project.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-11-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// projectName - The name of the project.
// body - Represents a project.
// options - ProjectsClientBeginCreateOrUpdateOptions contains the optional parameters for the ProjectsClient.BeginCreateOrUpdate
// method.
func (client *ProjectsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, projectName string, body Project, options *ProjectsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ProjectsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, projectName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[ProjectsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ProjectsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a project.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-11-preview
func (client *ProjectsClient) createOrUpdate(ctx context.Context, resourceGroupName string, projectName string, body Project, options *ProjectsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, projectName, body, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ProjectsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, projectName string, body Project, options *ProjectsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/projects/{projectName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-11-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginDelete - Deletes a project resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-11-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// projectName - The name of the project.
// options - ProjectsClientBeginDeleteOptions contains the optional parameters for the ProjectsClient.BeginDelete method.
func (client *ProjectsClient) BeginDelete(ctx context.Context, resourceGroupName string, projectName string, options *ProjectsClientBeginDeleteOptions) (*runtime.Poller[ProjectsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, projectName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[ProjectsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ProjectsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a project resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-11-preview
func (client *ProjectsClient) deleteOperation(ctx context.Context, resourceGroupName string, projectName string, options *ProjectsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, projectName, options)
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
func (client *ProjectsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, projectName string, options *ProjectsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/projects/{projectName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-11-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a specific project.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-11-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// projectName - The name of the project.
// options - ProjectsClientGetOptions contains the optional parameters for the ProjectsClient.Get method.
func (client *ProjectsClient) Get(ctx context.Context, resourceGroupName string, projectName string, options *ProjectsClientGetOptions) (ProjectsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, projectName, options)
	if err != nil {
		return ProjectsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProjectsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProjectsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ProjectsClient) getCreateRequest(ctx context.Context, resourceGroupName string, projectName string, options *ProjectsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/projects/{projectName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-11-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProjectsClient) getHandleResponse(resp *http.Response) (ProjectsClientGetResponse, error) {
	result := ProjectsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Project); err != nil {
		return ProjectsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists all projects in the resource group.
// Generated from API version 2022-11-11-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - ProjectsClientListByResourceGroupOptions contains the optional parameters for the ProjectsClient.ListByResourceGroup
// method.
func (client *ProjectsClient) NewListByResourceGroupPager(resourceGroupName string, options *ProjectsClientListByResourceGroupOptions) *runtime.Pager[ProjectsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProjectsClientListByResourceGroupResponse]{
		More: func(page ProjectsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ProjectsClientListByResourceGroupResponse) (ProjectsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ProjectsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ProjectsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ProjectsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ProjectsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ProjectsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/projects"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-11-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ProjectsClient) listByResourceGroupHandleResponse(resp *http.Response) (ProjectsClientListByResourceGroupResponse, error) {
	result := ProjectsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProjectListResult); err != nil {
		return ProjectsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Lists all projects in the subscription.
// Generated from API version 2022-11-11-preview
// options - ProjectsClientListBySubscriptionOptions contains the optional parameters for the ProjectsClient.ListBySubscription
// method.
func (client *ProjectsClient) NewListBySubscriptionPager(options *ProjectsClientListBySubscriptionOptions) *runtime.Pager[ProjectsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProjectsClientListBySubscriptionResponse]{
		More: func(page ProjectsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ProjectsClientListBySubscriptionResponse) (ProjectsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ProjectsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ProjectsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ProjectsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ProjectsClient) listBySubscriptionCreateRequest(ctx context.Context, options *ProjectsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DevCenter/projects"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-11-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ProjectsClient) listBySubscriptionHandleResponse(resp *http.Response) (ProjectsClientListBySubscriptionResponse, error) {
	result := ProjectsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProjectListResult); err != nil {
		return ProjectsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Partially updates a project.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-11-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// projectName - The name of the project.
// body - Updatable project properties.
// options - ProjectsClientBeginUpdateOptions contains the optional parameters for the ProjectsClient.BeginUpdate method.
func (client *ProjectsClient) BeginUpdate(ctx context.Context, resourceGroupName string, projectName string, body ProjectUpdate, options *ProjectsClientBeginUpdateOptions) (*runtime.Poller[ProjectsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, projectName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[ProjectsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ProjectsClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Partially updates a project.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-11-preview
func (client *ProjectsClient) update(ctx context.Context, resourceGroupName string, projectName string, body ProjectUpdate, options *ProjectsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, projectName, body, options)
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
func (client *ProjectsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, projectName string, body ProjectUpdate, options *ProjectsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/projects/{projectName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-11-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}
