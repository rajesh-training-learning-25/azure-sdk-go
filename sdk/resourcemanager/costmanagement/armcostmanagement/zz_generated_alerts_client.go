//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcostmanagement

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AlertsClient contains the methods for the Alerts group.
// Don't use this type directly, use NewAlertsClient() instead.
type AlertsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewAlertsClient creates a new instance of AlertsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAlertsClient(credential azcore.TokenCredential, options *arm.ClientOptions) *AlertsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &AlertsClient{
		host: string(cp.Endpoint),
		pl:   armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Dismiss - Dismisses the specified alert
// If the operation fails it returns an *azcore.ResponseError type.
// scope - The scope associated with alerts operations. This includes '/subscriptions/{subscriptionId}/' for subscription
// scope, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for
// resourceGroup scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope and
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount
// scope,
// '/providers/Microsoft.Management/managementGroups/{managementGroupId} for Management Group scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}'
// for billingProfile scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}'
// for invoiceSection scope, and
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' specific for partners.
// alertID - Alert ID
// parameters - Parameters supplied to the Dismiss Alert operation.
// options - AlertsClientDismissOptions contains the optional parameters for the AlertsClient.Dismiss method.
func (client *AlertsClient) Dismiss(ctx context.Context, scope string, alertID string, parameters DismissAlertPayload, options *AlertsClientDismissOptions) (AlertsClientDismissResponse, error) {
	req, err := client.dismissCreateRequest(ctx, scope, alertID, parameters, options)
	if err != nil {
		return AlertsClientDismissResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AlertsClientDismissResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AlertsClientDismissResponse{}, runtime.NewResponseError(resp)
	}
	return client.dismissHandleResponse(resp)
}

// dismissCreateRequest creates the Dismiss request.
func (client *AlertsClient) dismissCreateRequest(ctx context.Context, scope string, alertID string, parameters DismissAlertPayload, options *AlertsClientDismissOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.CostManagement/alerts/{alertId}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	urlPath = strings.ReplaceAll(urlPath, "{alertId}", alertID)
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// dismissHandleResponse handles the Dismiss response.
func (client *AlertsClient) dismissHandleResponse(resp *http.Response) (AlertsClientDismissResponse, error) {
	result := AlertsClientDismissResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Alert); err != nil {
		return AlertsClientDismissResponse{}, err
	}
	return result, nil
}

// Get - Gets the alert for the scope by alert ID.
// If the operation fails it returns an *azcore.ResponseError type.
// scope - The scope associated with alerts operations. This includes '/subscriptions/{subscriptionId}/' for subscription
// scope, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for
// resourceGroup scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope and
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount
// scope,
// '/providers/Microsoft.Management/managementGroups/{managementGroupId} for Management Group scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}'
// for billingProfile scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}'
// for invoiceSection scope, and
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' specific for partners.
// alertID - Alert ID
// options - AlertsClientGetOptions contains the optional parameters for the AlertsClient.Get method.
func (client *AlertsClient) Get(ctx context.Context, scope string, alertID string, options *AlertsClientGetOptions) (AlertsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, scope, alertID, options)
	if err != nil {
		return AlertsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AlertsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AlertsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AlertsClient) getCreateRequest(ctx context.Context, scope string, alertID string, options *AlertsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.CostManagement/alerts/{alertId}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	urlPath = strings.ReplaceAll(urlPath, "{alertId}", alertID)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AlertsClient) getHandleResponse(resp *http.Response) (AlertsClientGetResponse, error) {
	result := AlertsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Alert); err != nil {
		return AlertsClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists the alerts for scope defined.
// If the operation fails it returns an *azcore.ResponseError type.
// scope - The scope associated with alerts operations. This includes '/subscriptions/{subscriptionId}/' for subscription
// scope, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for
// resourceGroup scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope and
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount
// scope,
// '/providers/Microsoft.Management/managementGroups/{managementGroupId} for Management Group scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}'
// for billingProfile scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}'
// for invoiceSection scope, and
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' specific for partners.
// options - AlertsClientListOptions contains the optional parameters for the AlertsClient.List method.
func (client *AlertsClient) List(ctx context.Context, scope string, options *AlertsClientListOptions) (AlertsClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, scope, options)
	if err != nil {
		return AlertsClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AlertsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AlertsClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *AlertsClient) listCreateRequest(ctx context.Context, scope string, options *AlertsClientListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.CostManagement/alerts"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AlertsClient) listHandleResponse(resp *http.Response) (AlertsClientListResponse, error) {
	result := AlertsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertsResult); err != nil {
		return AlertsClientListResponse{}, err
	}
	return result, nil
}

// ListExternal - Lists the Alerts for external cloud provider type defined.
// If the operation fails it returns an *azcore.ResponseError type.
// externalCloudProviderType - The external cloud provider type associated with dimension/query operations. This includes
// 'externalSubscriptions' for linked account and 'externalBillingAccounts' for consolidated account.
// externalCloudProviderID - This can be '{externalSubscriptionId}' for linked account or '{externalBillingAccountId}' for
// consolidated account used with dimension/query operations.
// options - AlertsClientListExternalOptions contains the optional parameters for the AlertsClient.ListExternal method.
func (client *AlertsClient) ListExternal(ctx context.Context, externalCloudProviderType ExternalCloudProviderType, externalCloudProviderID string, options *AlertsClientListExternalOptions) (AlertsClientListExternalResponse, error) {
	req, err := client.listExternalCreateRequest(ctx, externalCloudProviderType, externalCloudProviderID, options)
	if err != nil {
		return AlertsClientListExternalResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AlertsClientListExternalResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AlertsClientListExternalResponse{}, runtime.NewResponseError(resp)
	}
	return client.listExternalHandleResponse(resp)
}

// listExternalCreateRequest creates the ListExternal request.
func (client *AlertsClient) listExternalCreateRequest(ctx context.Context, externalCloudProviderType ExternalCloudProviderType, externalCloudProviderID string, options *AlertsClientListExternalOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.CostManagement/{externalCloudProviderType}/{externalCloudProviderId}/alerts"
	if externalCloudProviderType == "" {
		return nil, errors.New("parameter externalCloudProviderType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{externalCloudProviderType}", url.PathEscape(string(externalCloudProviderType)))
	if externalCloudProviderID == "" {
		return nil, errors.New("parameter externalCloudProviderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{externalCloudProviderId}", url.PathEscape(externalCloudProviderID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listExternalHandleResponse handles the ListExternal response.
func (client *AlertsClient) listExternalHandleResponse(resp *http.Response) (AlertsClientListExternalResponse, error) {
	result := AlertsClientListExternalResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertsResult); err != nil {
		return AlertsClientListExternalResponse{}, err
	}
	return result, nil
}
