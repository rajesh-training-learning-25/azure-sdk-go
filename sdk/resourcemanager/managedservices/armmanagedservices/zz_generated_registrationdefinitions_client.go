//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagedservices

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// RegistrationDefinitionsClient contains the methods for the RegistrationDefinitions group.
// Don't use this type directly, use NewRegistrationDefinitionsClient() instead.
type RegistrationDefinitionsClient struct {
	ep string
	pl runtime.Pipeline
}

// NewRegistrationDefinitionsClient creates a new instance of RegistrationDefinitionsClient with the specified values.
func NewRegistrationDefinitionsClient(credential azcore.TokenCredential, options *arm.ClientOptions) *RegistrationDefinitionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &RegistrationDefinitionsClient{ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreateOrUpdate - Creates or updates a registration definition.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RegistrationDefinitionsClient) BeginCreateOrUpdate(ctx context.Context, registrationDefinitionID string, scope string, requestBody RegistrationDefinition, options *RegistrationDefinitionsBeginCreateOrUpdateOptions) (RegistrationDefinitionsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, registrationDefinitionID, scope, requestBody, options)
	if err != nil {
		return RegistrationDefinitionsCreateOrUpdatePollerResponse{}, err
	}
	result := RegistrationDefinitionsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("RegistrationDefinitionsClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return RegistrationDefinitionsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &RegistrationDefinitionsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a registration definition.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RegistrationDefinitionsClient) createOrUpdate(ctx context.Context, registrationDefinitionID string, scope string, requestBody RegistrationDefinition, options *RegistrationDefinitionsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, registrationDefinitionID, scope, requestBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RegistrationDefinitionsClient) createOrUpdateCreateRequest(ctx context.Context, registrationDefinitionID string, scope string, requestBody RegistrationDefinition, options *RegistrationDefinitionsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.ManagedServices/registrationDefinitions/{registrationDefinitionId}"
	if registrationDefinitionID == "" {
		return nil, errors.New("parameter registrationDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registrationDefinitionId}", url.PathEscape(registrationDefinitionID))
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, requestBody)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *RegistrationDefinitionsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Deletes the registration definition.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RegistrationDefinitionsClient) Delete(ctx context.Context, registrationDefinitionID string, scope string, options *RegistrationDefinitionsDeleteOptions) (RegistrationDefinitionsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, registrationDefinitionID, scope, options)
	if err != nil {
		return RegistrationDefinitionsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RegistrationDefinitionsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return RegistrationDefinitionsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return RegistrationDefinitionsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RegistrationDefinitionsClient) deleteCreateRequest(ctx context.Context, registrationDefinitionID string, scope string, options *RegistrationDefinitionsDeleteOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.ManagedServices/registrationDefinitions/{registrationDefinitionId}"
	if registrationDefinitionID == "" {
		return nil, errors.New("parameter registrationDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registrationDefinitionId}", url.PathEscape(registrationDefinitionID))
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *RegistrationDefinitionsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets the registration definition details.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RegistrationDefinitionsClient) Get(ctx context.Context, scope string, registrationDefinitionID string, options *RegistrationDefinitionsGetOptions) (RegistrationDefinitionsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, scope, registrationDefinitionID, options)
	if err != nil {
		return RegistrationDefinitionsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RegistrationDefinitionsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RegistrationDefinitionsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RegistrationDefinitionsClient) getCreateRequest(ctx context.Context, scope string, registrationDefinitionID string, options *RegistrationDefinitionsGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.ManagedServices/registrationDefinitions/{registrationDefinitionId}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if registrationDefinitionID == "" {
		return nil, errors.New("parameter registrationDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registrationDefinitionId}", url.PathEscape(registrationDefinitionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RegistrationDefinitionsClient) getHandleResponse(resp *http.Response) (RegistrationDefinitionsGetResponse, error) {
	result := RegistrationDefinitionsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RegistrationDefinition); err != nil {
		return RegistrationDefinitionsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *RegistrationDefinitionsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Gets a list of the registration definitions.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RegistrationDefinitionsClient) List(scope string, options *RegistrationDefinitionsListOptions) *RegistrationDefinitionsListPager {
	return &RegistrationDefinitionsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, scope, options)
		},
		advancer: func(ctx context.Context, resp RegistrationDefinitionsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RegistrationDefinitionList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *RegistrationDefinitionsClient) listCreateRequest(ctx context.Context, scope string, options *RegistrationDefinitionsListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.ManagedServices/registrationDefinitions"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RegistrationDefinitionsClient) listHandleResponse(resp *http.Response) (RegistrationDefinitionsListResponse, error) {
	result := RegistrationDefinitionsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RegistrationDefinitionList); err != nil {
		return RegistrationDefinitionsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *RegistrationDefinitionsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
