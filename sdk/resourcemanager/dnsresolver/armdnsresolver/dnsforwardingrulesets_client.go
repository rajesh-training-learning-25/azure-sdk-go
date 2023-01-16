//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdnsresolver

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

// DNSForwardingRulesetsClient contains the methods for the DNSForwardingRulesets group.
// Don't use this type directly, use NewDNSForwardingRulesetsClient() instead.
type DNSForwardingRulesetsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDNSForwardingRulesetsClient creates a new instance of DNSForwardingRulesetsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDNSForwardingRulesetsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DNSForwardingRulesetsClient, error) {
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
	client := &DNSForwardingRulesetsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a DNS forwarding ruleset.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// dnsForwardingRulesetName - The name of the DNS forwarding ruleset.
// parameters - Parameters supplied to the CreateOrUpdate operation.
// options - DNSForwardingRulesetsClientBeginCreateOrUpdateOptions contains the optional parameters for the DNSForwardingRulesetsClient.BeginCreateOrUpdate
// method.
func (client *DNSForwardingRulesetsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, dnsForwardingRulesetName string, parameters DNSForwardingRuleset, options *DNSForwardingRulesetsClientBeginCreateOrUpdateOptions) (*runtime.Poller[DNSForwardingRulesetsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, dnsForwardingRulesetName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DNSForwardingRulesetsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[DNSForwardingRulesetsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a DNS forwarding ruleset.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01
func (client *DNSForwardingRulesetsClient) createOrUpdate(ctx context.Context, resourceGroupName string, dnsForwardingRulesetName string, parameters DNSForwardingRuleset, options *DNSForwardingRulesetsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, dnsForwardingRulesetName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DNSForwardingRulesetsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, dnsForwardingRulesetName string, parameters DNSForwardingRuleset, options *DNSForwardingRulesetsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsForwardingRulesets/{dnsForwardingRulesetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dnsForwardingRulesetName == "" {
		return nil, errors.New("parameter dnsForwardingRulesetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dnsForwardingRulesetName}", url.PathEscape(dnsForwardingRulesetName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes a DNS forwarding ruleset. WARNING: This operation cannot be undone. All forwarding rules within the
// ruleset will be deleted.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// dnsForwardingRulesetName - The name of the DNS forwarding ruleset.
// options - DNSForwardingRulesetsClientBeginDeleteOptions contains the optional parameters for the DNSForwardingRulesetsClient.BeginDelete
// method.
func (client *DNSForwardingRulesetsClient) BeginDelete(ctx context.Context, resourceGroupName string, dnsForwardingRulesetName string, options *DNSForwardingRulesetsClientBeginDeleteOptions) (*runtime.Poller[DNSForwardingRulesetsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, dnsForwardingRulesetName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DNSForwardingRulesetsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[DNSForwardingRulesetsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a DNS forwarding ruleset. WARNING: This operation cannot be undone. All forwarding rules within the ruleset
// will be deleted.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01
func (client *DNSForwardingRulesetsClient) deleteOperation(ctx context.Context, resourceGroupName string, dnsForwardingRulesetName string, options *DNSForwardingRulesetsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, dnsForwardingRulesetName, options)
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
func (client *DNSForwardingRulesetsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, dnsForwardingRulesetName string, options *DNSForwardingRulesetsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsForwardingRulesets/{dnsForwardingRulesetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dnsForwardingRulesetName == "" {
		return nil, errors.New("parameter dnsForwardingRulesetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dnsForwardingRulesetName}", url.PathEscape(dnsForwardingRulesetName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a DNS forwarding ruleset properties.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// dnsForwardingRulesetName - The name of the DNS forwarding ruleset.
// options - DNSForwardingRulesetsClientGetOptions contains the optional parameters for the DNSForwardingRulesetsClient.Get
// method.
func (client *DNSForwardingRulesetsClient) Get(ctx context.Context, resourceGroupName string, dnsForwardingRulesetName string, options *DNSForwardingRulesetsClientGetOptions) (DNSForwardingRulesetsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, dnsForwardingRulesetName, options)
	if err != nil {
		return DNSForwardingRulesetsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DNSForwardingRulesetsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DNSForwardingRulesetsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DNSForwardingRulesetsClient) getCreateRequest(ctx context.Context, resourceGroupName string, dnsForwardingRulesetName string, options *DNSForwardingRulesetsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsForwardingRulesets/{dnsForwardingRulesetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dnsForwardingRulesetName == "" {
		return nil, errors.New("parameter dnsForwardingRulesetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dnsForwardingRulesetName}", url.PathEscape(dnsForwardingRulesetName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DNSForwardingRulesetsClient) getHandleResponse(resp *http.Response) (DNSForwardingRulesetsClientGetResponse, error) {
	result := DNSForwardingRulesetsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DNSForwardingRuleset); err != nil {
		return DNSForwardingRulesetsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists DNS forwarding rulesets in all resource groups of a subscription.
// Generated from API version 2022-07-01
// options - DNSForwardingRulesetsClientListOptions contains the optional parameters for the DNSForwardingRulesetsClient.List
// method.
func (client *DNSForwardingRulesetsClient) NewListPager(options *DNSForwardingRulesetsClientListOptions) *runtime.Pager[DNSForwardingRulesetsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DNSForwardingRulesetsClientListResponse]{
		More: func(page DNSForwardingRulesetsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DNSForwardingRulesetsClientListResponse) (DNSForwardingRulesetsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DNSForwardingRulesetsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DNSForwardingRulesetsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DNSForwardingRulesetsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *DNSForwardingRulesetsClient) listCreateRequest(ctx context.Context, options *DNSForwardingRulesetsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/dnsForwardingRulesets"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DNSForwardingRulesetsClient) listHandleResponse(resp *http.Response) (DNSForwardingRulesetsClientListResponse, error) {
	result := DNSForwardingRulesetsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DNSForwardingRulesetListResult); err != nil {
		return DNSForwardingRulesetsClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists DNS forwarding rulesets within a resource group.
// Generated from API version 2022-07-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - DNSForwardingRulesetsClientListByResourceGroupOptions contains the optional parameters for the DNSForwardingRulesetsClient.ListByResourceGroup
// method.
func (client *DNSForwardingRulesetsClient) NewListByResourceGroupPager(resourceGroupName string, options *DNSForwardingRulesetsClientListByResourceGroupOptions) *runtime.Pager[DNSForwardingRulesetsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[DNSForwardingRulesetsClientListByResourceGroupResponse]{
		More: func(page DNSForwardingRulesetsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DNSForwardingRulesetsClientListByResourceGroupResponse) (DNSForwardingRulesetsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DNSForwardingRulesetsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DNSForwardingRulesetsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DNSForwardingRulesetsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DNSForwardingRulesetsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DNSForwardingRulesetsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsForwardingRulesets"
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
	reqQP.Set("api-version", "2022-07-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DNSForwardingRulesetsClient) listByResourceGroupHandleResponse(resp *http.Response) (DNSForwardingRulesetsClientListByResourceGroupResponse, error) {
	result := DNSForwardingRulesetsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DNSForwardingRulesetListResult); err != nil {
		return DNSForwardingRulesetsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListByVirtualNetworkPager - Lists DNS forwarding ruleset resource IDs attached to a virtual network.
// Generated from API version 2022-07-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// virtualNetworkName - The name of the virtual network.
// options - DNSForwardingRulesetsClientListByVirtualNetworkOptions contains the optional parameters for the DNSForwardingRulesetsClient.ListByVirtualNetwork
// method.
func (client *DNSForwardingRulesetsClient) NewListByVirtualNetworkPager(resourceGroupName string, virtualNetworkName string, options *DNSForwardingRulesetsClientListByVirtualNetworkOptions) *runtime.Pager[DNSForwardingRulesetsClientListByVirtualNetworkResponse] {
	return runtime.NewPager(runtime.PagingHandler[DNSForwardingRulesetsClientListByVirtualNetworkResponse]{
		More: func(page DNSForwardingRulesetsClientListByVirtualNetworkResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DNSForwardingRulesetsClientListByVirtualNetworkResponse) (DNSForwardingRulesetsClientListByVirtualNetworkResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByVirtualNetworkCreateRequest(ctx, resourceGroupName, virtualNetworkName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DNSForwardingRulesetsClientListByVirtualNetworkResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DNSForwardingRulesetsClientListByVirtualNetworkResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DNSForwardingRulesetsClientListByVirtualNetworkResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByVirtualNetworkHandleResponse(resp)
		},
	})
}

// listByVirtualNetworkCreateRequest creates the ListByVirtualNetwork request.
func (client *DNSForwardingRulesetsClient) listByVirtualNetworkCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, options *DNSForwardingRulesetsClientListByVirtualNetworkOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/listDnsForwardingRulesets"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByVirtualNetworkHandleResponse handles the ListByVirtualNetwork response.
func (client *DNSForwardingRulesetsClient) listByVirtualNetworkHandleResponse(resp *http.Response) (DNSForwardingRulesetsClientListByVirtualNetworkResponse, error) {
	result := DNSForwardingRulesetsClientListByVirtualNetworkResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkDNSForwardingRulesetListResult); err != nil {
		return DNSForwardingRulesetsClientListByVirtualNetworkResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates a DNS forwarding ruleset.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// dnsForwardingRulesetName - The name of the DNS forwarding ruleset.
// parameters - Parameters supplied to the Update operation.
// options - DNSForwardingRulesetsClientBeginUpdateOptions contains the optional parameters for the DNSForwardingRulesetsClient.BeginUpdate
// method.
func (client *DNSForwardingRulesetsClient) BeginUpdate(ctx context.Context, resourceGroupName string, dnsForwardingRulesetName string, parameters DNSForwardingRulesetPatch, options *DNSForwardingRulesetsClientBeginUpdateOptions) (*runtime.Poller[DNSForwardingRulesetsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, dnsForwardingRulesetName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DNSForwardingRulesetsClientUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[DNSForwardingRulesetsClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Updates a DNS forwarding ruleset.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01
func (client *DNSForwardingRulesetsClient) update(ctx context.Context, resourceGroupName string, dnsForwardingRulesetName string, parameters DNSForwardingRulesetPatch, options *DNSForwardingRulesetsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, dnsForwardingRulesetName, parameters, options)
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
func (client *DNSForwardingRulesetsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, dnsForwardingRulesetName string, parameters DNSForwardingRulesetPatch, options *DNSForwardingRulesetsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsForwardingRulesets/{dnsForwardingRulesetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dnsForwardingRulesetName == "" {
		return nil, errors.New("parameter dnsForwardingRulesetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dnsForwardingRulesetName}", url.PathEscape(dnsForwardingRulesetName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}
