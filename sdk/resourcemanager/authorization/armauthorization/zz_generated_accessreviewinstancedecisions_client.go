//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

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

// AccessReviewInstanceDecisionsClient contains the methods for the AccessReviewInstanceDecisions group.
// Don't use this type directly, use NewAccessReviewInstanceDecisionsClient() instead.
type AccessReviewInstanceDecisionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAccessReviewInstanceDecisionsClient creates a new instance of AccessReviewInstanceDecisionsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAccessReviewInstanceDecisionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AccessReviewInstanceDecisionsClient, error) {
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
	client := &AccessReviewInstanceDecisionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListPager - Get access review instance decisions
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-16-preview
// scheduleDefinitionID - The id of the access review schedule definition.
// id - The id of the access review instance.
// options - AccessReviewInstanceDecisionsClientListOptions contains the optional parameters for the AccessReviewInstanceDecisionsClient.List
// method.
func (client *AccessReviewInstanceDecisionsClient) NewListPager(scheduleDefinitionID string, id string, options *AccessReviewInstanceDecisionsClientListOptions) *runtime.Pager[AccessReviewInstanceDecisionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AccessReviewInstanceDecisionsClientListResponse]{
		More: func(page AccessReviewInstanceDecisionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AccessReviewInstanceDecisionsClientListResponse) (AccessReviewInstanceDecisionsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, scheduleDefinitionID, id, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AccessReviewInstanceDecisionsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AccessReviewInstanceDecisionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AccessReviewInstanceDecisionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *AccessReviewInstanceDecisionsClient) listCreateRequest(ctx context.Context, scheduleDefinitionID string, id string, options *AccessReviewInstanceDecisionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/instances/{id}/decisions"
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-16-preview")
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
func (client *AccessReviewInstanceDecisionsClient) listHandleResponse(resp *http.Response) (AccessReviewInstanceDecisionsClientListResponse, error) {
	result := AccessReviewInstanceDecisionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewDecisionListResult); err != nil {
		return AccessReviewInstanceDecisionsClientListResponse{}, err
	}
	return result, nil
}
