//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// SubscriptionUsagesClient contains the methods for the SubscriptionUsages group.
// Don't use this type directly, use NewSubscriptionUsagesClient() instead.
type SubscriptionUsagesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSubscriptionUsagesClient creates a new instance of SubscriptionUsagesClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSubscriptionUsagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SubscriptionUsagesClient, error) {
	cl, err := arm.NewClient(moduleName+".SubscriptionUsagesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SubscriptionUsagesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets a subscription usage metric.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
//   - locationName - The name of the region where the resource is located.
//   - usageName - Name of usage metric to return.
//   - options - SubscriptionUsagesClientGetOptions contains the optional parameters for the SubscriptionUsagesClient.Get method.
func (client *SubscriptionUsagesClient) Get(ctx context.Context, locationName string, usageName string, options *SubscriptionUsagesClientGetOptions) (SubscriptionUsagesClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, locationName, usageName, options)
	if err != nil {
		return SubscriptionUsagesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SubscriptionUsagesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SubscriptionUsagesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SubscriptionUsagesClient) getCreateRequest(ctx context.Context, locationName string, usageName string, options *SubscriptionUsagesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationName}/usages/{usageName}"
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if usageName == "" {
		return nil, errors.New("parameter usageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{usageName}", url.PathEscape(usageName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SubscriptionUsagesClient) getHandleResponse(resp *http.Response) (SubscriptionUsagesClientGetResponse, error) {
	result := SubscriptionUsagesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubscriptionUsage); err != nil {
		return SubscriptionUsagesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByLocationPager - Gets all subscription usage metrics in a given location.
//
// Generated from API version 2020-11-01-preview
//   - locationName - The name of the region where the resource is located.
//   - options - SubscriptionUsagesClientListByLocationOptions contains the optional parameters for the SubscriptionUsagesClient.NewListByLocationPager
//     method.
func (client *SubscriptionUsagesClient) NewListByLocationPager(locationName string, options *SubscriptionUsagesClientListByLocationOptions) *runtime.Pager[SubscriptionUsagesClientListByLocationResponse] {
	return runtime.NewPager(runtime.PagingHandler[SubscriptionUsagesClientListByLocationResponse]{
		More: func(page SubscriptionUsagesClientListByLocationResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SubscriptionUsagesClientListByLocationResponse) (SubscriptionUsagesClientListByLocationResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByLocationCreateRequest(ctx, locationName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SubscriptionUsagesClientListByLocationResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return SubscriptionUsagesClientListByLocationResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SubscriptionUsagesClientListByLocationResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByLocationHandleResponse(resp)
		},
	})
}

// listByLocationCreateRequest creates the ListByLocation request.
func (client *SubscriptionUsagesClient) listByLocationCreateRequest(ctx context.Context, locationName string, options *SubscriptionUsagesClientListByLocationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationName}/usages"
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByLocationHandleResponse handles the ListByLocation response.
func (client *SubscriptionUsagesClient) listByLocationHandleResponse(resp *http.Response) (SubscriptionUsagesClientListByLocationResponse, error) {
	result := SubscriptionUsagesClientListByLocationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubscriptionUsageListResult); err != nil {
		return SubscriptionUsagesClientListByLocationResponse{}, err
	}
	return result, nil
}
