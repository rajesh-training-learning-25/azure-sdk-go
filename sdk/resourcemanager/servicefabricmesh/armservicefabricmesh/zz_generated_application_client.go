//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicefabricmesh

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

// ApplicationClient contains the methods for the Application group.
// Don't use this type directly, use NewApplicationClient() instead.
type ApplicationClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewApplicationClient creates a new instance of ApplicationClient with the specified values.
// subscriptionID - The customer subscription identifier
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewApplicationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ApplicationClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &ApplicationClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Create - Creates an application resource with the specified name, description and properties. If an application resource
// with the same name exists, then it is updated with the specified description and
// properties.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Azure resource group name
// applicationResourceName - The identity of the application.
// applicationResourceDescription - Description for creating a Application resource.
// options - ApplicationClientCreateOptions contains the optional parameters for the ApplicationClient.Create method.
func (client *ApplicationClient) Create(ctx context.Context, resourceGroupName string, applicationResourceName string, applicationResourceDescription ApplicationResourceDescription, options *ApplicationClientCreateOptions) (ApplicationClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, applicationResourceName, applicationResourceDescription, options)
	if err != nil {
		return ApplicationClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApplicationClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return ApplicationClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *ApplicationClient) createCreateRequest(ctx context.Context, resourceGroupName string, applicationResourceName string, applicationResourceDescription ApplicationResourceDescription, options *ApplicationClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/applications/{applicationResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{applicationResourceName}", applicationResourceName)
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, applicationResourceDescription)
}

// createHandleResponse handles the Create response.
func (client *ApplicationClient) createHandleResponse(resp *http.Response) (ApplicationClientCreateResponse, error) {
	result := ApplicationClientCreateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationResourceDescription); err != nil {
		return ApplicationClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the application resource identified by the name.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Azure resource group name
// applicationResourceName - The identity of the application.
// options - ApplicationClientDeleteOptions contains the optional parameters for the ApplicationClient.Delete method.
func (client *ApplicationClient) Delete(ctx context.Context, resourceGroupName string, applicationResourceName string, options *ApplicationClientDeleteOptions) (ApplicationClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, applicationResourceName, options)
	if err != nil {
		return ApplicationClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApplicationClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return ApplicationClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ApplicationClientDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ApplicationClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, applicationResourceName string, options *ApplicationClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/applications/{applicationResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{applicationResourceName}", applicationResourceName)
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the information about the application resource with the given name. The information include the description
// and other properties of the application.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Azure resource group name
// applicationResourceName - The identity of the application.
// options - ApplicationClientGetOptions contains the optional parameters for the ApplicationClient.Get method.
func (client *ApplicationClient) Get(ctx context.Context, resourceGroupName string, applicationResourceName string, options *ApplicationClientGetOptions) (ApplicationClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, applicationResourceName, options)
	if err != nil {
		return ApplicationClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApplicationClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ApplicationClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ApplicationClient) getCreateRequest(ctx context.Context, resourceGroupName string, applicationResourceName string, options *ApplicationClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/applications/{applicationResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{applicationResourceName}", applicationResourceName)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ApplicationClient) getHandleResponse(resp *http.Response) (ApplicationClientGetResponse, error) {
	result := ApplicationClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationResourceDescription); err != nil {
		return ApplicationClientGetResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Gets the information about all application resources in a given resource group. The information include
// the description and other properties of the Application.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Azure resource group name
// options - ApplicationClientListByResourceGroupOptions contains the optional parameters for the ApplicationClient.ListByResourceGroup
// method.
func (client *ApplicationClient) ListByResourceGroup(resourceGroupName string, options *ApplicationClientListByResourceGroupOptions) *ApplicationClientListByResourceGroupPager {
	return &ApplicationClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp ApplicationClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ApplicationResourceDescriptionList.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ApplicationClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ApplicationClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/applications"
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
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ApplicationClient) listByResourceGroupHandleResponse(resp *http.Response) (ApplicationClientListByResourceGroupResponse, error) {
	result := ApplicationClientListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationResourceDescriptionList); err != nil {
		return ApplicationClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListBySubscription - Gets the information about all application resources in a given resource group. The information include
// the description and other properties of the application.
// If the operation fails it returns an *azcore.ResponseError type.
// options - ApplicationClientListBySubscriptionOptions contains the optional parameters for the ApplicationClient.ListBySubscription
// method.
func (client *ApplicationClient) ListBySubscription(options *ApplicationClientListBySubscriptionOptions) *ApplicationClientListBySubscriptionPager {
	return &ApplicationClientListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ApplicationClientListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ApplicationResourceDescriptionList.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ApplicationClient) listBySubscriptionCreateRequest(ctx context.Context, options *ApplicationClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabricMesh/applications"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ApplicationClient) listBySubscriptionHandleResponse(resp *http.Response) (ApplicationClientListBySubscriptionResponse, error) {
	result := ApplicationClientListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationResourceDescriptionList); err != nil {
		return ApplicationClientListBySubscriptionResponse{}, err
	}
	return result, nil
}
