//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappplatform

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

// CertificatesClient contains the methods for the Certificates group.
// Don't use this type directly, use NewCertificatesClient() instead.
type CertificatesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCertificatesClient creates a new instance of CertificatesClient with the specified values.
//   - subscriptionID - Gets subscription ID which uniquely identify the Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCertificatesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CertificatesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CertificatesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update certificate resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - certificateName - The name of the certificate resource.
//   - certificateResource - Parameters for the create or update operation
//   - options - CertificatesClientBeginCreateOrUpdateOptions contains the optional parameters for the CertificatesClient.BeginCreateOrUpdate
//     method.
func (client *CertificatesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, certificateName string, certificateResource CertificateResource, options *CertificatesClientBeginCreateOrUpdateOptions) (*runtime.Poller[CertificatesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serviceName, certificateName, certificateResource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CertificatesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[CertificatesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create or update certificate resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
func (client *CertificatesClient) createOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, certificateName string, certificateResource CertificateResource, options *CertificatesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "CertificatesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, certificateName, certificateResource, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *CertificatesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, certificateName string, certificateResource CertificateResource, options *CertificatesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/certificates/{certificateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, certificateResource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete the certificate resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - certificateName - The name of the certificate resource.
//   - options - CertificatesClientBeginDeleteOptions contains the optional parameters for the CertificatesClient.BeginDelete
//     method.
func (client *CertificatesClient) BeginDelete(ctx context.Context, resourceGroupName string, serviceName string, certificateName string, options *CertificatesClientBeginDeleteOptions) (*runtime.Poller[CertificatesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, serviceName, certificateName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CertificatesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[CertificatesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete the certificate resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
func (client *CertificatesClient) deleteOperation(ctx context.Context, resourceGroupName string, serviceName string, certificateName string, options *CertificatesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "CertificatesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, certificateName, options)
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
func (client *CertificatesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, certificateName string, options *CertificatesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/certificates/{certificateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the certificate resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - certificateName - The name of the certificate resource.
//   - options - CertificatesClientGetOptions contains the optional parameters for the CertificatesClient.Get method.
func (client *CertificatesClient) Get(ctx context.Context, resourceGroupName string, serviceName string, certificateName string, options *CertificatesClientGetOptions) (CertificatesClientGetResponse, error) {
	var err error
	const operationName = "CertificatesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, certificateName, options)
	if err != nil {
		return CertificatesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CertificatesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CertificatesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *CertificatesClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, certificateName string, options *CertificatesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/certificates/{certificateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CertificatesClient) getHandleResponse(resp *http.Response) (CertificatesClientGetResponse, error) {
	result := CertificatesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CertificateResource); err != nil {
		return CertificatesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List all the certificates of one user.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - options - CertificatesClientListOptions contains the optional parameters for the CertificatesClient.NewListPager method.
func (client *CertificatesClient) NewListPager(resourceGroupName string, serviceName string, options *CertificatesClientListOptions) *runtime.Pager[CertificatesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[CertificatesClientListResponse]{
		More: func(page CertificatesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CertificatesClientListResponse) (CertificatesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CertificatesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, serviceName, options)
			}, nil)
			if err != nil {
				return CertificatesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *CertificatesClient) listCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *CertificatesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/certificates"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *CertificatesClient) listHandleResponse(resp *http.Response) (CertificatesClientListResponse, error) {
	result := CertificatesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CertificateResourceCollection); err != nil {
		return CertificatesClientListResponse{}, err
	}
	return result, nil
}
