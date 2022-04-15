//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridnetwork

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

// VendorNetworkFunctionsClient contains the methods for the VendorNetworkFunctions group.
// Don't use this type directly, use NewVendorNetworkFunctionsClient() instead.
type VendorNetworkFunctionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVendorNetworkFunctionsClient creates a new instance of VendorNetworkFunctionsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVendorNetworkFunctionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VendorNetworkFunctionsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &VendorNetworkFunctionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a vendor network function. This operation can take up to 6 hours to complete.
// This is expected service behavior.
// If the operation fails it returns an *azcore.ResponseError type.
// locationName - The Azure region where the network function resource was created by the customer.
// vendorName - The name of the vendor.
// serviceKey - The GUID for the vendor network function.
// parameters - Parameters supplied to the create or update vendor network function operation.
// options - VendorNetworkFunctionsClientBeginCreateOrUpdateOptions contains the optional parameters for the VendorNetworkFunctionsClient.BeginCreateOrUpdate
// method.
func (client *VendorNetworkFunctionsClient) BeginCreateOrUpdate(ctx context.Context, locationName string, vendorName string, serviceKey string, parameters VendorNetworkFunction, options *VendorNetworkFunctionsClientBeginCreateOrUpdateOptions) (*armruntime.Poller[VendorNetworkFunctionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, locationName, vendorName, serviceKey, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[VendorNetworkFunctionsClientCreateOrUpdateResponse]{
			FinalStateVia: armruntime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[VendorNetworkFunctionsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a vendor network function. This operation can take up to 6 hours to complete. This
// is expected service behavior.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VendorNetworkFunctionsClient) createOrUpdate(ctx context.Context, locationName string, vendorName string, serviceKey string, parameters VendorNetworkFunction, options *VendorNetworkFunctionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, locationName, vendorName, serviceKey, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VendorNetworkFunctionsClient) createOrUpdateCreateRequest(ctx context.Context, locationName string, vendorName string, serviceKey string, parameters VendorNetworkFunction, options *VendorNetworkFunctionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/locations/{locationName}/vendors/{vendorName}/networkFunctions/{serviceKey}"
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if vendorName == "" {
		return nil, errors.New("parameter vendorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vendorName}", url.PathEscape(vendorName))
	if serviceKey == "" {
		return nil, errors.New("parameter serviceKey cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceKey}", url.PathEscape(serviceKey))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// Get - Gets information about the specified vendor network function.
// If the operation fails it returns an *azcore.ResponseError type.
// locationName - The Azure region where the network function resource was created by the customer.
// vendorName - The name of the vendor.
// serviceKey - The GUID for the vendor network function.
// options - VendorNetworkFunctionsClientGetOptions contains the optional parameters for the VendorNetworkFunctionsClient.Get
// method.
func (client *VendorNetworkFunctionsClient) Get(ctx context.Context, locationName string, vendorName string, serviceKey string, options *VendorNetworkFunctionsClientGetOptions) (VendorNetworkFunctionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, locationName, vendorName, serviceKey, options)
	if err != nil {
		return VendorNetworkFunctionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VendorNetworkFunctionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VendorNetworkFunctionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VendorNetworkFunctionsClient) getCreateRequest(ctx context.Context, locationName string, vendorName string, serviceKey string, options *VendorNetworkFunctionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/locations/{locationName}/vendors/{vendorName}/networkFunctions/{serviceKey}"
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if vendorName == "" {
		return nil, errors.New("parameter vendorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vendorName}", url.PathEscape(vendorName))
	if serviceKey == "" {
		return nil, errors.New("parameter serviceKey cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceKey}", url.PathEscape(serviceKey))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VendorNetworkFunctionsClient) getHandleResponse(resp *http.Response) (VendorNetworkFunctionsClientGetResponse, error) {
	result := VendorNetworkFunctionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VendorNetworkFunction); err != nil {
		return VendorNetworkFunctionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all the vendor network function sub resources in an Azure region, filtered by skuType, skuName, vendorProvisioningState.
// If the operation fails it returns an *azcore.ResponseError type.
// locationName - The Azure region where the network function resource was created by the customer.
// vendorName - The name of the vendor.
// options - VendorNetworkFunctionsClientListOptions contains the optional parameters for the VendorNetworkFunctionsClient.List
// method.
func (client *VendorNetworkFunctionsClient) NewListPager(locationName string, vendorName string, options *VendorNetworkFunctionsClientListOptions) *runtime.Pager[VendorNetworkFunctionsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[VendorNetworkFunctionsClientListResponse]{
		More: func(page VendorNetworkFunctionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VendorNetworkFunctionsClientListResponse) (VendorNetworkFunctionsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, locationName, vendorName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VendorNetworkFunctionsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return VendorNetworkFunctionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VendorNetworkFunctionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *VendorNetworkFunctionsClient) listCreateRequest(ctx context.Context, locationName string, vendorName string, options *VendorNetworkFunctionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/locations/{locationName}/vendors/{vendorName}/networkFunctions"
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if vendorName == "" {
		return nil, errors.New("parameter vendorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vendorName}", url.PathEscape(vendorName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VendorNetworkFunctionsClient) listHandleResponse(resp *http.Response) (VendorNetworkFunctionsClientListResponse, error) {
	result := VendorNetworkFunctionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VendorNetworkFunctionListResult); err != nil {
		return VendorNetworkFunctionsClientListResponse{}, err
	}
	return result, nil
}
