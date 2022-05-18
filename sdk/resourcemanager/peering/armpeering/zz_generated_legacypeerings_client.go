//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpeering

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

// LegacyPeeringsClient contains the methods for the LegacyPeerings group.
// Don't use this type directly, use NewLegacyPeeringsClient() instead.
type LegacyPeeringsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewLegacyPeeringsClient creates a new instance of LegacyPeeringsClient with the specified values.
// subscriptionID - The Azure subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewLegacyPeeringsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LegacyPeeringsClient, error) {
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
	client := &LegacyPeeringsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListPager - Lists all of the legacy peerings under the given subscription matching the specified kind and location.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// peeringLocation - The location of the peering.
// kind - The kind of the peering.
// options - LegacyPeeringsClientListOptions contains the optional parameters for the LegacyPeeringsClient.List method.
func (client *LegacyPeeringsClient) NewListPager(peeringLocation string, kind LegacyPeeringsKind, options *LegacyPeeringsClientListOptions) *runtime.Pager[LegacyPeeringsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[LegacyPeeringsClientListResponse]{
		More: func(page LegacyPeeringsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LegacyPeeringsClientListResponse) (LegacyPeeringsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, peeringLocation, kind, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return LegacyPeeringsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return LegacyPeeringsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LegacyPeeringsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *LegacyPeeringsClient) listCreateRequest(ctx context.Context, peeringLocation string, kind LegacyPeeringsKind, options *LegacyPeeringsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Peering/legacyPeerings"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("peeringLocation", peeringLocation)
	reqQP.Set("kind", string(kind))
	if options != nil && options.Asn != nil {
		reqQP.Set("asn", strconv.FormatInt(int64(*options.Asn), 10))
	}
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LegacyPeeringsClient) listHandleResponse(resp *http.Response) (LegacyPeeringsClientListResponse, error) {
	result := LegacyPeeringsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResult); err != nil {
		return LegacyPeeringsClientListResponse{}, err
	}
	return result, nil
}
