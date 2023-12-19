//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

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

// AccessReviewInstancesClient contains the methods for the AccessReviewInstances group.
// Don't use this type directly, use NewAccessReviewInstancesClient() instead.
type AccessReviewInstancesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAccessReviewInstancesClient creates a new instance of AccessReviewInstancesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAccessReviewInstancesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AccessReviewInstancesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AccessReviewInstancesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Update access review instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01-preview
//   - scheduleDefinitionID - The id of the access review schedule definition.
//   - id - The id of the access review instance.
//   - properties - Access review instance properties.
//   - options - AccessReviewInstancesClientCreateOptions contains the optional parameters for the AccessReviewInstancesClient.Create
//     method.
func (client *AccessReviewInstancesClient) Create(ctx context.Context, scheduleDefinitionID string, id string, properties AccessReviewInstanceProperties, options *AccessReviewInstancesClientCreateOptions) (AccessReviewInstancesClientCreateResponse, error) {
	var err error
	const operationName = "AccessReviewInstancesClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, scheduleDefinitionID, id, properties, options)
	if err != nil {
		return AccessReviewInstancesClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AccessReviewInstancesClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AccessReviewInstancesClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *AccessReviewInstancesClient) createCreateRequest(ctx context.Context, scheduleDefinitionID string, id string, properties AccessReviewInstanceProperties, options *AccessReviewInstancesClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/instances/{id}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *AccessReviewInstancesClient) createHandleResponse(resp *http.Response) (AccessReviewInstancesClientCreateResponse, error) {
	result := AccessReviewInstancesClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewInstance); err != nil {
		return AccessReviewInstancesClientCreateResponse{}, err
	}
	return result, nil
}

// GetByID - Get access review instances
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01-preview
//   - scheduleDefinitionID - The id of the access review schedule definition.
//   - id - The id of the access review instance.
//   - options - AccessReviewInstancesClientGetByIDOptions contains the optional parameters for the AccessReviewInstancesClient.GetByID
//     method.
func (client *AccessReviewInstancesClient) GetByID(ctx context.Context, scheduleDefinitionID string, id string, options *AccessReviewInstancesClientGetByIDOptions) (AccessReviewInstancesClientGetByIDResponse, error) {
	var err error
	const operationName = "AccessReviewInstancesClient.GetByID"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getByIDCreateRequest(ctx, scheduleDefinitionID, id, options)
	if err != nil {
		return AccessReviewInstancesClientGetByIDResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AccessReviewInstancesClientGetByIDResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AccessReviewInstancesClientGetByIDResponse{}, err
	}
	resp, err := client.getByIDHandleResponse(httpResp)
	return resp, err
}

// getByIDCreateRequest creates the GetByID request.
func (client *AccessReviewInstancesClient) getByIDCreateRequest(ctx context.Context, scheduleDefinitionID string, id string, options *AccessReviewInstancesClientGetByIDOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/instances/{id}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByIDHandleResponse handles the GetByID response.
func (client *AccessReviewInstancesClient) getByIDHandleResponse(resp *http.Response) (AccessReviewInstancesClientGetByIDResponse, error) {
	result := AccessReviewInstancesClientGetByIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewInstance); err != nil {
		return AccessReviewInstancesClientGetByIDResponse{}, err
	}
	return result, nil
}

// NewListPager - Get access review instances
//
// Generated from API version 2021-12-01-preview
//   - scheduleDefinitionID - The id of the access review schedule definition.
//   - options - AccessReviewInstancesClientListOptions contains the optional parameters for the AccessReviewInstancesClient.NewListPager
//     method.
func (client *AccessReviewInstancesClient) NewListPager(scheduleDefinitionID string, options *AccessReviewInstancesClientListOptions) *runtime.Pager[AccessReviewInstancesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AccessReviewInstancesClientListResponse]{
		More: func(page AccessReviewInstancesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AccessReviewInstancesClientListResponse) (AccessReviewInstancesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AccessReviewInstancesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, scheduleDefinitionID, options)
			}, nil)
			if err != nil {
				return AccessReviewInstancesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *AccessReviewInstancesClient) listCreateRequest(ctx context.Context, scheduleDefinitionID string, options *AccessReviewInstancesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/instances"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.Raw().URL.RawQuery}
	if options != nil && options.Filter != nil {
		unencodedParams = append(unencodedParams, "$filter="+*options.Filter)
	}
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AccessReviewInstancesClient) listHandleResponse(resp *http.Response) (AccessReviewInstancesClientListResponse, error) {
	result := AccessReviewInstancesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewInstanceListResult); err != nil {
		return AccessReviewInstancesClientListResponse{}, err
	}
	return result, nil
}
