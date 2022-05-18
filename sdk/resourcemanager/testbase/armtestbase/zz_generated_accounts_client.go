//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtestbase

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

// AccountsClient contains the methods for the TestBaseAccounts group.
// Don't use this type directly, use NewAccountsClient() instead.
type AccountsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAccountsClient creates a new instance of AccountsClient with the specified values.
// subscriptionID - The Azure subscription ID. This is a GUID-formatted string.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAccountsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AccountsClient, error) {
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
	client := &AccountsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CheckPackageNameAvailability - Checks that the Test Base Package name and version is valid and is not already in use.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-12-16-preview
// resourceGroupName - The name of the resource group that contains the resource.
// testBaseAccountName - The resource name of the Test Base Account.
// parameters - Parameters supplied to the Test Base Package CheckNameAvailability operation.
// options - AccountsClientCheckPackageNameAvailabilityOptions contains the optional parameters for the AccountsClient.CheckPackageNameAvailability
// method.
func (client *AccountsClient) CheckPackageNameAvailability(ctx context.Context, resourceGroupName string, testBaseAccountName string, parameters PackageCheckNameAvailabilityParameters, options *AccountsClientCheckPackageNameAvailabilityOptions) (AccountsClientCheckPackageNameAvailabilityResponse, error) {
	req, err := client.checkPackageNameAvailabilityCreateRequest(ctx, resourceGroupName, testBaseAccountName, parameters, options)
	if err != nil {
		return AccountsClientCheckPackageNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccountsClientCheckPackageNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccountsClientCheckPackageNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkPackageNameAvailabilityHandleResponse(resp)
}

// checkPackageNameAvailabilityCreateRequest creates the CheckPackageNameAvailability request.
func (client *AccountsClient) checkPackageNameAvailabilityCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, parameters PackageCheckNameAvailabilityParameters, options *AccountsClientCheckPackageNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/checkPackageNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if testBaseAccountName == "" {
		return nil, errors.New("parameter testBaseAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testBaseAccountName}", url.PathEscape(testBaseAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// checkPackageNameAvailabilityHandleResponse handles the CheckPackageNameAvailability response.
func (client *AccountsClient) checkPackageNameAvailabilityHandleResponse(resp *http.Response) (AccountsClientCheckPackageNameAvailabilityResponse, error) {
	result := AccountsClientCheckPackageNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityResult); err != nil {
		return AccountsClientCheckPackageNameAvailabilityResponse{}, err
	}
	return result, nil
}

// BeginCreate - Create or replace (overwrite/recreate, with potential downtime) a Test Base Account in the specified subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-12-16-preview
// resourceGroupName - The name of the resource group that contains the resource.
// testBaseAccountName - The resource name of the Test Base Account.
// parameters - Parameters supplied to create a Test Base Account.
// options - AccountsClientBeginCreateOptions contains the optional parameters for the AccountsClient.BeginCreate method.
func (client *AccountsClient) BeginCreate(ctx context.Context, resourceGroupName string, testBaseAccountName string, parameters AccountResource, options *AccountsClientBeginCreateOptions) (*runtime.Poller[AccountsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, testBaseAccountName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[AccountsClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[AccountsClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Create or replace (overwrite/recreate, with potential downtime) a Test Base Account in the specified subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-12-16-preview
func (client *AccountsClient) create(ctx context.Context, resourceGroupName string, testBaseAccountName string, parameters AccountResource, options *AccountsClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, testBaseAccountName, parameters, options)
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

// createCreateRequest creates the Create request.
func (client *AccountsClient) createCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, parameters AccountResource, options *AccountsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if testBaseAccountName == "" {
		return nil, errors.New("parameter testBaseAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testBaseAccountName}", url.PathEscape(testBaseAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Restore != nil {
		reqQP.Set("restore", strconv.FormatBool(*options.Restore))
	}
	reqQP.Set("api-version", "2020-12-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes a Test Base Account.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-12-16-preview
// resourceGroupName - The name of the resource group that contains the resource.
// testBaseAccountName - The resource name of the Test Base Account.
// options - AccountsClientBeginDeleteOptions contains the optional parameters for the AccountsClient.BeginDelete method.
func (client *AccountsClient) BeginDelete(ctx context.Context, resourceGroupName string, testBaseAccountName string, options *AccountsClientBeginDeleteOptions) (*runtime.Poller[AccountsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, testBaseAccountName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[AccountsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[AccountsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a Test Base Account.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-12-16-preview
func (client *AccountsClient) deleteOperation(ctx context.Context, resourceGroupName string, testBaseAccountName string, options *AccountsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, testBaseAccountName, options)
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
func (client *AccountsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, options *AccountsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if testBaseAccountName == "" {
		return nil, errors.New("parameter testBaseAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testBaseAccountName}", url.PathEscape(testBaseAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a Test Base Account.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-12-16-preview
// resourceGroupName - The name of the resource group that contains the resource.
// testBaseAccountName - The resource name of the Test Base Account.
// options - AccountsClientGetOptions contains the optional parameters for the AccountsClient.Get method.
func (client *AccountsClient) Get(ctx context.Context, resourceGroupName string, testBaseAccountName string, options *AccountsClientGetOptions) (AccountsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, testBaseAccountName, options)
	if err != nil {
		return AccountsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccountsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccountsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AccountsClient) getCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, options *AccountsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if testBaseAccountName == "" {
		return nil, errors.New("parameter testBaseAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testBaseAccountName}", url.PathEscape(testBaseAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AccountsClient) getHandleResponse(resp *http.Response) (AccountsClientGetResponse, error) {
	result := AccountsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccountResource); err != nil {
		return AccountsClientGetResponse{}, err
	}
	return result, nil
}

// GetFileUploadURL - Gets the file upload URL of a Test Base Account.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-12-16-preview
// resourceGroupName - The name of the resource group that contains the resource.
// testBaseAccountName - The resource name of the Test Base Account.
// options - AccountsClientGetFileUploadURLOptions contains the optional parameters for the AccountsClient.GetFileUploadURL
// method.
func (client *AccountsClient) GetFileUploadURL(ctx context.Context, resourceGroupName string, testBaseAccountName string, options *AccountsClientGetFileUploadURLOptions) (AccountsClientGetFileUploadURLResponse, error) {
	req, err := client.getFileUploadURLCreateRequest(ctx, resourceGroupName, testBaseAccountName, options)
	if err != nil {
		return AccountsClientGetFileUploadURLResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccountsClientGetFileUploadURLResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccountsClientGetFileUploadURLResponse{}, runtime.NewResponseError(resp)
	}
	return client.getFileUploadURLHandleResponse(resp)
}

// getFileUploadURLCreateRequest creates the GetFileUploadURL request.
func (client *AccountsClient) getFileUploadURLCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, options *AccountsClientGetFileUploadURLOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/getFileUploadUrl"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if testBaseAccountName == "" {
		return nil, errors.New("parameter testBaseAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testBaseAccountName}", url.PathEscape(testBaseAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Parameters != nil {
		return req, runtime.MarshalAsJSON(req, *options.Parameters)
	}
	return req, nil
}

// getFileUploadURLHandleResponse handles the GetFileUploadURL response.
func (client *AccountsClient) getFileUploadURLHandleResponse(resp *http.Response) (AccountsClientGetFileUploadURLResponse, error) {
	result := AccountsClientGetFileUploadURLResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FileUploadURLResponse); err != nil {
		return AccountsClientGetFileUploadURLResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists all the Test Base Accounts in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-12-16-preview
// resourceGroupName - The name of the resource group that contains the resource.
// options - AccountsClientListByResourceGroupOptions contains the optional parameters for the AccountsClient.ListByResourceGroup
// method.
func (client *AccountsClient) NewListByResourceGroupPager(resourceGroupName string, options *AccountsClientListByResourceGroupOptions) *runtime.Pager[AccountsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[AccountsClientListByResourceGroupResponse]{
		More: func(page AccountsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AccountsClientListByResourceGroupResponse) (AccountsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AccountsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AccountsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AccountsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *AccountsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *AccountsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts"
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
	if options != nil && options.GetDeleted != nil {
		reqQP.Set("getDeleted", strconv.FormatBool(*options.GetDeleted))
	}
	reqQP.Set("api-version", "2020-12-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *AccountsClient) listByResourceGroupHandleResponse(resp *http.Response) (AccountsClientListByResourceGroupResponse, error) {
	result := AccountsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccountListResult); err != nil {
		return AccountsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Lists all the Test Base Accounts in a subscription. This API is required by ARM guidelines.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-12-16-preview
// options - AccountsClientListBySubscriptionOptions contains the optional parameters for the AccountsClient.ListBySubscription
// method.
func (client *AccountsClient) NewListBySubscriptionPager(options *AccountsClientListBySubscriptionOptions) *runtime.Pager[AccountsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[AccountsClientListBySubscriptionResponse]{
		More: func(page AccountsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AccountsClientListBySubscriptionResponse) (AccountsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AccountsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AccountsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AccountsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *AccountsClient) listBySubscriptionCreateRequest(ctx context.Context, options *AccountsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.TestBase/testBaseAccounts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.GetDeleted != nil {
		reqQP.Set("getDeleted", strconv.FormatBool(*options.GetDeleted))
	}
	reqQP.Set("api-version", "2020-12-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *AccountsClient) listBySubscriptionHandleResponse(resp *http.Response) (AccountsClientListBySubscriptionResponse, error) {
	result := AccountsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccountListResult); err != nil {
		return AccountsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginOffboard - Offboard a Test Base Account.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-12-16-preview
// resourceGroupName - The name of the resource group that contains the resource.
// testBaseAccountName - The resource name of the Test Base Account.
// options - AccountsClientBeginOffboardOptions contains the optional parameters for the AccountsClient.BeginOffboard method.
func (client *AccountsClient) BeginOffboard(ctx context.Context, resourceGroupName string, testBaseAccountName string, options *AccountsClientBeginOffboardOptions) (*runtime.Poller[AccountsClientOffboardResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.offboard(ctx, resourceGroupName, testBaseAccountName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[AccountsClientOffboardResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[AccountsClientOffboardResponse](options.ResumeToken, client.pl, nil)
	}
}

// Offboard - Offboard a Test Base Account.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-12-16-preview
func (client *AccountsClient) offboard(ctx context.Context, resourceGroupName string, testBaseAccountName string, options *AccountsClientBeginOffboardOptions) (*http.Response, error) {
	req, err := client.offboardCreateRequest(ctx, resourceGroupName, testBaseAccountName, options)
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

// offboardCreateRequest creates the Offboard request.
func (client *AccountsClient) offboardCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, options *AccountsClientBeginOffboardOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/offboard"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if testBaseAccountName == "" {
		return nil, errors.New("parameter testBaseAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testBaseAccountName}", url.PathEscape(testBaseAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginUpdate - Update an existing Test Base Account.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-12-16-preview
// resourceGroupName - The name of the resource group that contains the resource.
// testBaseAccountName - The resource name of the Test Base Account.
// parameters - Parameters supplied to update a Test Base Account.
// options - AccountsClientBeginUpdateOptions contains the optional parameters for the AccountsClient.BeginUpdate method.
func (client *AccountsClient) BeginUpdate(ctx context.Context, resourceGroupName string, testBaseAccountName string, parameters AccountUpdateParameters, options *AccountsClientBeginUpdateOptions) (*runtime.Poller[AccountsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, testBaseAccountName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[AccountsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[AccountsClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Update an existing Test Base Account.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-12-16-preview
func (client *AccountsClient) update(ctx context.Context, resourceGroupName string, testBaseAccountName string, parameters AccountUpdateParameters, options *AccountsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, testBaseAccountName, parameters, options)
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

// updateCreateRequest creates the Update request.
func (client *AccountsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, parameters AccountUpdateParameters, options *AccountsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if testBaseAccountName == "" {
		return nil, errors.New("parameter testBaseAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testBaseAccountName}", url.PathEscape(testBaseAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}
