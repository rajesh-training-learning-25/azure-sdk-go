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

// CommitmentTiersClient contains the methods for the CommitmentTiers group.
// Don't use this type directly, use NewCommitmentTiersClient() instead.
type CommitmentTiersClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewCommitmentTiersClient creates a new instance of CommitmentTiersClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewCommitmentTiersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CommitmentTiersClient, error) {
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
	client := &CommitmentTiersClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListPager - List Commitment Tiers.
// Generated from API version 2022-10-01
// location - Resource location.
// options - CommitmentTiersClientListOptions contains the optional parameters for the CommitmentTiersClient.List method.
func (client *CommitmentTiersClient) NewListPager(location string, options *CommitmentTiersClientListOptions) *runtime.Pager[CommitmentTiersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[CommitmentTiersClientListResponse]{
		More: func(page CommitmentTiersClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CommitmentTiersClientListResponse) (CommitmentTiersClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, location, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return CommitmentTiersClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return CommitmentTiersClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CommitmentTiersClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *CommitmentTiersClient) listCreateRequest(ctx context.Context, location string, options *CommitmentTiersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.CognitiveServices/locations/{location}/commitmentTiers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
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
func (client *CommitmentTiersClient) listHandleResponse(resp *http.Response) (CommitmentTiersClientListResponse, error) {
	result := CommitmentTiersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CommitmentTierListResult); err != nil {
		return CommitmentTiersClientListResponse{}, err
	}
	return result, nil
}
