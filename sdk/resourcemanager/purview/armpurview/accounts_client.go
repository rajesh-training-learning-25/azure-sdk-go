//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armpurview

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

// AccountsClient contains the methods for the Accounts group.
// Don't use this type directly, use NewAccountsClient() instead.
type AccountsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAccountsClient creates a new instance of AccountsClient with the specified values.
//   - subscriptionID - The subscription identifier
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAccountsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AccountsClient, error) {
	cl, err := arm.NewClient(moduleName+".AccountsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AccountsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// AddRootCollectionAdmin - Add the administrator for root collection associated with this account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01
//   - resourceGroupName - The resource group name.
//   - accountName - The name of the account.
//   - collectionAdminUpdate - The collection admin update payload.
//   - options - AccountsClientAddRootCollectionAdminOptions contains the optional parameters for the AccountsClient.AddRootCollectionAdmin
//     method.
func (client *AccountsClient) AddRootCollectionAdmin(ctx context.Context, resourceGroupName string, accountName string, collectionAdminUpdate CollectionAdminUpdate, options *AccountsClientAddRootCollectionAdminOptions) (AccountsClientAddRootCollectionAdminResponse, error) {
	req, err := client.addRootCollectionAdminCreateRequest(ctx, resourceGroupName, accountName, collectionAdminUpdate, options)
	if err != nil {
		return AccountsClientAddRootCollectionAdminResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AccountsClientAddRootCollectionAdminResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccountsClientAddRootCollectionAdminResponse{}, runtime.NewResponseError(resp)
	}
	return AccountsClientAddRootCollectionAdminResponse{}, nil
}

// addRootCollectionAdminCreateRequest creates the AddRootCollectionAdmin request.
func (client *AccountsClient) addRootCollectionAdminCreateRequest(ctx context.Context, resourceGroupName string, accountName string, collectionAdminUpdate CollectionAdminUpdate, options *AccountsClientAddRootCollectionAdminOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Purview/accounts/{accountName}/addRootCollectionAdmin"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, collectionAdminUpdate)
}

// CheckNameAvailability - Checks if account name is available.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01
//   - checkNameAvailabilityRequest - The check name availability request.
//   - options - AccountsClientCheckNameAvailabilityOptions contains the optional parameters for the AccountsClient.CheckNameAvailability
//     method.
func (client *AccountsClient) CheckNameAvailability(ctx context.Context, checkNameAvailabilityRequest CheckNameAvailabilityRequest, options *AccountsClientCheckNameAvailabilityOptions) (AccountsClientCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, checkNameAvailabilityRequest, options)
	if err != nil {
		return AccountsClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AccountsClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccountsClientCheckNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *AccountsClient) checkNameAvailabilityCreateRequest(ctx context.Context, checkNameAvailabilityRequest CheckNameAvailabilityRequest, options *AccountsClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Purview/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, checkNameAvailabilityRequest)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *AccountsClient) checkNameAvailabilityHandleResponse(resp *http.Response) (AccountsClientCheckNameAvailabilityResponse, error) {
	result := AccountsClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityResult); err != nil {
		return AccountsClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// BeginCreateOrUpdate - Creates or updates an account
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01
//   - resourceGroupName - The resource group name.
//   - accountName - The name of the account.
//   - account - The account.
//   - options - AccountsClientBeginCreateOrUpdateOptions contains the optional parameters for the AccountsClient.BeginCreateOrUpdate
//     method.
func (client *AccountsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, account Account, options *AccountsClientBeginCreateOrUpdateOptions) (*runtime.Poller[AccountsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, accountName, account, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[AccountsClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[AccountsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates an account
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01
func (client *AccountsClient) createOrUpdate(ctx context.Context, resourceGroupName string, accountName string, account Account, options *AccountsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, account, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AccountsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, account Account, options *AccountsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Purview/accounts/{accountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, account)
}

// BeginDelete - Deletes an account resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01
//   - resourceGroupName - The resource group name.
//   - accountName - The name of the account.
//   - options - AccountsClientBeginDeleteOptions contains the optional parameters for the AccountsClient.BeginDelete method.
func (client *AccountsClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, options *AccountsClientBeginDeleteOptions) (*runtime.Poller[AccountsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[AccountsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[AccountsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes an account resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01
func (client *AccountsClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, options *AccountsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AccountsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *AccountsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Purview/accounts/{accountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get an account
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01
//   - resourceGroupName - The resource group name.
//   - accountName - The name of the account.
//   - options - AccountsClientGetOptions contains the optional parameters for the AccountsClient.Get method.
func (client *AccountsClient) Get(ctx context.Context, resourceGroupName string, accountName string, options *AccountsClientGetOptions) (AccountsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return AccountsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AccountsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccountsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AccountsClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *AccountsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Purview/accounts/{accountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AccountsClient) getHandleResponse(resp *http.Response) (AccountsClientGetResponse, error) {
	result := AccountsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Account); err != nil {
		return AccountsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List accounts in ResourceGroup
//
// Generated from API version 2021-12-01
//   - resourceGroupName - The resource group name.
//   - options - AccountsClientListByResourceGroupOptions contains the optional parameters for the AccountsClient.NewListByResourceGroupPager
//     method.
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
			resp, err := client.internal.Pipeline().Do(req)
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
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Purview/accounts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *AccountsClient) listByResourceGroupHandleResponse(resp *http.Response) (AccountsClientListByResourceGroupResponse, error) {
	result := AccountsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccountList); err != nil {
		return AccountsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List accounts in Subscription
//
// Generated from API version 2021-12-01
//   - options - AccountsClientListBySubscriptionOptions contains the optional parameters for the AccountsClient.NewListBySubscriptionPager
//     method.
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
			resp, err := client.internal.Pipeline().Do(req)
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
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Purview/accounts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *AccountsClient) listBySubscriptionHandleResponse(resp *http.Response) (AccountsClientListBySubscriptionResponse, error) {
	result := AccountsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccountList); err != nil {
		return AccountsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// ListKeys - List the authorization keys associated with this account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01
//   - resourceGroupName - The resource group name.
//   - accountName - The name of the account.
//   - options - AccountsClientListKeysOptions contains the optional parameters for the AccountsClient.ListKeys method.
func (client *AccountsClient) ListKeys(ctx context.Context, resourceGroupName string, accountName string, options *AccountsClientListKeysOptions) (AccountsClientListKeysResponse, error) {
	req, err := client.listKeysCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return AccountsClientListKeysResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AccountsClientListKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccountsClientListKeysResponse{}, runtime.NewResponseError(resp)
	}
	return client.listKeysHandleResponse(resp)
}

// listKeysCreateRequest creates the ListKeys request.
func (client *AccountsClient) listKeysCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *AccountsClientListKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Purview/accounts/{accountName}/listkeys"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listKeysHandleResponse handles the ListKeys response.
func (client *AccountsClient) listKeysHandleResponse(resp *http.Response) (AccountsClientListKeysResponse, error) {
	result := AccountsClientListKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessKeys); err != nil {
		return AccountsClientListKeysResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates an account
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01
//   - resourceGroupName - The resource group name.
//   - accountName - The name of the account.
//   - accountUpdateParameters - The account update parameters.
//   - options - AccountsClientBeginUpdateOptions contains the optional parameters for the AccountsClient.BeginUpdate method.
func (client *AccountsClient) BeginUpdate(ctx context.Context, resourceGroupName string, accountName string, accountUpdateParameters AccountUpdateParameters, options *AccountsClientBeginUpdateOptions) (*runtime.Poller[AccountsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, accountName, accountUpdateParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[AccountsClientUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[AccountsClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Updates an account
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01
func (client *AccountsClient) update(ctx context.Context, resourceGroupName string, accountName string, accountUpdateParameters AccountUpdateParameters, options *AccountsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, accountUpdateParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *AccountsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, accountUpdateParameters AccountUpdateParameters, options *AccountsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Purview/accounts/{accountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, accountUpdateParameters)
}
