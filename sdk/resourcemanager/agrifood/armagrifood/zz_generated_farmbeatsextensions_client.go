//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armagrifood

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// FarmBeatsExtensionsClient contains the methods for the FarmBeatsExtensions group.
// Don't use this type directly, use NewFarmBeatsExtensionsClient() instead.
type FarmBeatsExtensionsClient struct {
	ep string
	pl runtime.Pipeline
}

// NewFarmBeatsExtensionsClient creates a new instance of FarmBeatsExtensionsClient with the specified values.
func NewFarmBeatsExtensionsClient(con *arm.Connection) *FarmBeatsExtensionsClient {
	return &FarmBeatsExtensionsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version)}
}

// Get - Get farmBeats extension.
// If the operation fails it returns the *ErrorResponse error type.
func (client *FarmBeatsExtensionsClient) Get(ctx context.Context, farmBeatsExtensionID string, options *FarmBeatsExtensionsGetOptions) (FarmBeatsExtensionsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, farmBeatsExtensionID, options)
	if err != nil {
		return FarmBeatsExtensionsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FarmBeatsExtensionsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FarmBeatsExtensionsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *FarmBeatsExtensionsClient) getCreateRequest(ctx context.Context, farmBeatsExtensionID string, options *FarmBeatsExtensionsGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AgFoodPlatform/farmBeatsExtensionDefinitions/{farmBeatsExtensionId}"
	if farmBeatsExtensionID == "" {
		return nil, errors.New("parameter farmBeatsExtensionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{farmBeatsExtensionId}", url.PathEscape(farmBeatsExtensionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-12-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FarmBeatsExtensionsClient) getHandleResponse(resp *http.Response) (FarmBeatsExtensionsGetResponse, error) {
	result := FarmBeatsExtensionsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.FarmBeatsExtension); err != nil {
		return FarmBeatsExtensionsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *FarmBeatsExtensionsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Get list of farmBeats extension.
// If the operation fails it returns the *ErrorResponse error type.
func (client *FarmBeatsExtensionsClient) List(options *FarmBeatsExtensionsListOptions) *FarmBeatsExtensionsListPager {
	return &FarmBeatsExtensionsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp FarmBeatsExtensionsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.FarmBeatsExtensionListResponse.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *FarmBeatsExtensionsClient) listCreateRequest(ctx context.Context, options *FarmBeatsExtensionsListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AgFoodPlatform/farmBeatsExtensionDefinitions"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.FarmBeatsExtensionIDs != nil {
		for _, qv := range options.FarmBeatsExtensionIDs {
			reqQP.Add("farmBeatsExtensionIds", qv)
		}
	}
	if options != nil && options.FarmBeatsExtensionNames != nil {
		for _, qv := range options.FarmBeatsExtensionNames {
			reqQP.Add("farmBeatsExtensionNames", qv)
		}
	}
	if options != nil && options.ExtensionCategories != nil {
		for _, qv := range options.ExtensionCategories {
			reqQP.Add("extensionCategories", qv)
		}
	}
	if options != nil && options.PublisherIDs != nil {
		for _, qv := range options.PublisherIDs {
			reqQP.Add("publisherIds", qv)
		}
	}
	if options != nil && options.MaxPageSize != nil {
		reqQP.Set("$maxPageSize", strconv.FormatInt(int64(*options.MaxPageSize), 10))
	}
	reqQP.Set("api-version", "2020-05-12-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *FarmBeatsExtensionsClient) listHandleResponse(resp *http.Response) (FarmBeatsExtensionsListResponse, error) {
	result := FarmBeatsExtensionsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.FarmBeatsExtensionListResponse); err != nil {
		return FarmBeatsExtensionsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *FarmBeatsExtensionsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
