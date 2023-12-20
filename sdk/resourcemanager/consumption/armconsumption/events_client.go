//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconsumption

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

// EventsClient contains the methods for the Events group.
// Don't use this type directly, use NewEventsClient() instead.
type EventsClient struct {
	internal *arm.Client
}

// NewEventsClient creates a new instance of EventsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewEventsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*EventsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &EventsClient{
		internal: cl,
	}
	return client, nil
}

// NewListByBillingAccountPager - Lists the events that decrements Azure credits or Microsoft Azure consumption commitment
// for a billing account or a billing profile for a given start and end date.
//
// Generated from API version 2021-10-01
//   - billingAccountID - BillingAccount ID
//   - options - EventsClientListByBillingAccountOptions contains the optional parameters for the EventsClient.NewListByBillingAccountPager
//     method.
func (client *EventsClient) NewListByBillingAccountPager(billingAccountID string, options *EventsClientListByBillingAccountOptions) *runtime.Pager[EventsClientListByBillingAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[EventsClientListByBillingAccountResponse]{
		More: func(page EventsClientListByBillingAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EventsClientListByBillingAccountResponse) (EventsClientListByBillingAccountResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "EventsClient.NewListByBillingAccountPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByBillingAccountCreateRequest(ctx, billingAccountID, options)
			}, nil)
			if err != nil {
				return EventsClientListByBillingAccountResponse{}, err
			}
			return client.listByBillingAccountHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByBillingAccountCreateRequest creates the ListByBillingAccount request.
func (client *EventsClient) listByBillingAccountCreateRequest(ctx context.Context, billingAccountID string, options *EventsClientListByBillingAccountOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/providers/Microsoft.Consumption/events"
	if billingAccountID == "" {
		return nil, errors.New("parameter billingAccountID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountId}", url.PathEscape(billingAccountID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByBillingAccountHandleResponse handles the ListByBillingAccount response.
func (client *EventsClient) listByBillingAccountHandleResponse(resp *http.Response) (EventsClientListByBillingAccountResponse, error) {
	result := EventsClientListByBillingAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Events); err != nil {
		return EventsClientListByBillingAccountResponse{}, err
	}
	return result, nil
}

// NewListByBillingProfilePager - Lists the events that decrements Azure credits or Microsoft Azure consumption commitment
// for a billing account or a billing profile for a given start and end date.
//
// Generated from API version 2021-10-01
//   - billingAccountID - BillingAccount ID
//   - billingProfileID - Azure Billing Profile ID.
//   - startDate - Start date
//   - endDate - End date
//   - options - EventsClientListByBillingProfileOptions contains the optional parameters for the EventsClient.NewListByBillingProfilePager
//     method.
func (client *EventsClient) NewListByBillingProfilePager(billingAccountID string, billingProfileID string, startDate string, endDate string, options *EventsClientListByBillingProfileOptions) *runtime.Pager[EventsClientListByBillingProfileResponse] {
	return runtime.NewPager(runtime.PagingHandler[EventsClientListByBillingProfileResponse]{
		More: func(page EventsClientListByBillingProfileResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EventsClientListByBillingProfileResponse) (EventsClientListByBillingProfileResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "EventsClient.NewListByBillingProfilePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByBillingProfileCreateRequest(ctx, billingAccountID, billingProfileID, startDate, endDate, options)
			}, nil)
			if err != nil {
				return EventsClientListByBillingProfileResponse{}, err
			}
			return client.listByBillingProfileHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByBillingProfileCreateRequest creates the ListByBillingProfile request.
func (client *EventsClient) listByBillingProfileCreateRequest(ctx context.Context, billingAccountID string, billingProfileID string, startDate string, endDate string, options *EventsClientListByBillingProfileOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/providers/Microsoft.Consumption/events"
	if billingAccountID == "" {
		return nil, errors.New("parameter billingAccountID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountId}", url.PathEscape(billingAccountID))
	if billingProfileID == "" {
		return nil, errors.New("parameter billingProfileID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileId}", url.PathEscape(billingProfileID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	reqQP.Set("startDate", startDate)
	reqQP.Set("endDate", endDate)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByBillingProfileHandleResponse handles the ListByBillingProfile response.
func (client *EventsClient) listByBillingProfileHandleResponse(resp *http.Response) (EventsClientListByBillingProfileResponse, error) {
	result := EventsClientListByBillingProfileResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Events); err != nil {
		return EventsClientListByBillingProfileResponse{}, err
	}
	return result, nil
}
