//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevtestlabs

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
	"strconv"
	"strings"
)

// ArmTemplatesClient contains the methods for the ArmTemplates group.
// Don't use this type directly, use NewArmTemplatesClient() instead.
type ArmTemplatesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewArmTemplatesClient creates a new instance of ArmTemplatesClient with the specified values.
// subscriptionID - The subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewArmTemplatesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ArmTemplatesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &ArmTemplatesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Get - Get azure resource manager template.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// labName - The name of the lab.
// artifactSourceName - The name of the artifact source.
// name - The name of the azure resource manager template.
// options - ArmTemplatesClientGetOptions contains the optional parameters for the ArmTemplatesClient.Get method.
func (client *ArmTemplatesClient) Get(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, name string, options *ArmTemplatesClientGetOptions) (ArmTemplatesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, labName, artifactSourceName, name, options)
	if err != nil {
		return ArmTemplatesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ArmTemplatesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ArmTemplatesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ArmTemplatesClient) getCreateRequest(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, name string, options *ArmTemplatesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{artifactSourceName}/armtemplates/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if artifactSourceName == "" {
		return nil, errors.New("parameter artifactSourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{artifactSourceName}", url.PathEscape(artifactSourceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ArmTemplatesClient) getHandleResponse(resp *http.Response) (ArmTemplatesClientGetResponse, error) {
	result := ArmTemplatesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArmTemplate); err != nil {
		return ArmTemplatesClientGetResponse{}, err
	}
	return result, nil
}

// List - List azure resource manager templates in a given artifact source.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// labName - The name of the lab.
// artifactSourceName - The name of the artifact source.
// options - ArmTemplatesClientListOptions contains the optional parameters for the ArmTemplatesClient.List method.
func (client *ArmTemplatesClient) List(resourceGroupName string, labName string, artifactSourceName string, options *ArmTemplatesClientListOptions) *ArmTemplatesClientListPager {
	return &ArmTemplatesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, labName, artifactSourceName, options)
		},
		advancer: func(ctx context.Context, resp ArmTemplatesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ArmTemplateList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ArmTemplatesClient) listCreateRequest(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, options *ArmTemplatesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{artifactSourceName}/armtemplates"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if artifactSourceName == "" {
		return nil, errors.New("parameter artifactSourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{artifactSourceName}", url.PathEscape(artifactSourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ArmTemplatesClient) listHandleResponse(resp *http.Response) (ArmTemplatesClientListResponse, error) {
	result := ArmTemplatesClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArmTemplateList); err != nil {
		return ArmTemplatesClientListResponse{}, err
	}
	return result, nil
}
