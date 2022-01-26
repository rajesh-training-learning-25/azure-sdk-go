//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

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

// AlertsSuppressionRulesClient contains the methods for the AlertsSuppressionRules group.
// Don't use this type directly, use NewAlertsSuppressionRulesClient() instead.
type AlertsSuppressionRulesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAlertsSuppressionRulesClient creates a new instance of AlertsSuppressionRulesClient with the specified values.
// subscriptionID - Azure subscription ID
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAlertsSuppressionRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *AlertsSuppressionRulesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &AlertsSuppressionRulesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Delete - Delete dismiss alert rule for this subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// alertsSuppressionRuleName - The unique name of the suppression alert rule
// options - AlertsSuppressionRulesClientDeleteOptions contains the optional parameters for the AlertsSuppressionRulesClient.Delete
// method.
func (client *AlertsSuppressionRulesClient) Delete(ctx context.Context, alertsSuppressionRuleName string, options *AlertsSuppressionRulesClientDeleteOptions) (AlertsSuppressionRulesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, alertsSuppressionRuleName, options)
	if err != nil {
		return AlertsSuppressionRulesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AlertsSuppressionRulesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return AlertsSuppressionRulesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return AlertsSuppressionRulesClientDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AlertsSuppressionRulesClient) deleteCreateRequest(ctx context.Context, alertsSuppressionRuleName string, options *AlertsSuppressionRulesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/alertsSuppressionRules/{alertsSuppressionRuleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if alertsSuppressionRuleName == "" {
		return nil, errors.New("parameter alertsSuppressionRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alertsSuppressionRuleName}", url.PathEscape(alertsSuppressionRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get dismiss rule, with name: {alertsSuppressionRuleName}, for the given subscription
// If the operation fails it returns an *azcore.ResponseError type.
// alertsSuppressionRuleName - The unique name of the suppression alert rule
// options - AlertsSuppressionRulesClientGetOptions contains the optional parameters for the AlertsSuppressionRulesClient.Get
// method.
func (client *AlertsSuppressionRulesClient) Get(ctx context.Context, alertsSuppressionRuleName string, options *AlertsSuppressionRulesClientGetOptions) (AlertsSuppressionRulesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, alertsSuppressionRuleName, options)
	if err != nil {
		return AlertsSuppressionRulesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AlertsSuppressionRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AlertsSuppressionRulesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AlertsSuppressionRulesClient) getCreateRequest(ctx context.Context, alertsSuppressionRuleName string, options *AlertsSuppressionRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/alertsSuppressionRules/{alertsSuppressionRuleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if alertsSuppressionRuleName == "" {
		return nil, errors.New("parameter alertsSuppressionRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alertsSuppressionRuleName}", url.PathEscape(alertsSuppressionRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AlertsSuppressionRulesClient) getHandleResponse(resp *http.Response) (AlertsSuppressionRulesClientGetResponse, error) {
	result := AlertsSuppressionRulesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertsSuppressionRule); err != nil {
		return AlertsSuppressionRulesClientGetResponse{}, err
	}
	return result, nil
}

// List - List of all the dismiss rules for the given subscription
// If the operation fails it returns an *azcore.ResponseError type.
// options - AlertsSuppressionRulesClientListOptions contains the optional parameters for the AlertsSuppressionRulesClient.List
// method.
func (client *AlertsSuppressionRulesClient) List(options *AlertsSuppressionRulesClientListOptions) *AlertsSuppressionRulesClientListPager {
	return &AlertsSuppressionRulesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp AlertsSuppressionRulesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AlertsSuppressionRulesList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *AlertsSuppressionRulesClient) listCreateRequest(ctx context.Context, options *AlertsSuppressionRulesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/alertsSuppressionRules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-01-01-preview")
	if options != nil && options.AlertType != nil {
		reqQP.Set("AlertType", *options.AlertType)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AlertsSuppressionRulesClient) listHandleResponse(resp *http.Response) (AlertsSuppressionRulesClientListResponse, error) {
	result := AlertsSuppressionRulesClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertsSuppressionRulesList); err != nil {
		return AlertsSuppressionRulesClientListResponse{}, err
	}
	return result, nil
}

// Update - Update existing rule or create new rule if it doesn't exist
// If the operation fails it returns an *azcore.ResponseError type.
// alertsSuppressionRuleName - The unique name of the suppression alert rule
// alertsSuppressionRule - Suppression rule object
// options - AlertsSuppressionRulesClientUpdateOptions contains the optional parameters for the AlertsSuppressionRulesClient.Update
// method.
func (client *AlertsSuppressionRulesClient) Update(ctx context.Context, alertsSuppressionRuleName string, alertsSuppressionRule AlertsSuppressionRule, options *AlertsSuppressionRulesClientUpdateOptions) (AlertsSuppressionRulesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, alertsSuppressionRuleName, alertsSuppressionRule, options)
	if err != nil {
		return AlertsSuppressionRulesClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AlertsSuppressionRulesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AlertsSuppressionRulesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *AlertsSuppressionRulesClient) updateCreateRequest(ctx context.Context, alertsSuppressionRuleName string, alertsSuppressionRule AlertsSuppressionRule, options *AlertsSuppressionRulesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/alertsSuppressionRules/{alertsSuppressionRuleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if alertsSuppressionRuleName == "" {
		return nil, errors.New("parameter alertsSuppressionRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alertsSuppressionRuleName}", url.PathEscape(alertsSuppressionRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, alertsSuppressionRule)
}

// updateHandleResponse handles the Update response.
func (client *AlertsSuppressionRulesClient) updateHandleResponse(resp *http.Response) (AlertsSuppressionRulesClientUpdateResponse, error) {
	result := AlertsSuppressionRulesClientUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertsSuppressionRule); err != nil {
		return AlertsSuppressionRulesClientUpdateResponse{}, err
	}
	return result, nil
}
