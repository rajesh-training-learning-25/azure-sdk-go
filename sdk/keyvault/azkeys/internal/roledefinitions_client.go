//go:build go1.16
// +build go1.16

// Code generated by Microsoft (R) AutoRest Code Generator (autorest: 3.6.2, generator: @autorest/go@4.0.0-preview.29)
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package internal

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

type roleDefinitionsClient struct {
	con *connection
}

// CreateOrUpdate - Creates or updates a custom role definition.
// If the operation fails it returns the *KeyVaultError error type.
func (client *roleDefinitionsClient) CreateOrUpdate(ctx context.Context, vaultBaseURL string, scope string, roleDefinitionName string, parameters RoleDefinitionCreateParameters, options *RoleDefinitionsCreateOrUpdateOptions) (RoleDefinitionsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, vaultBaseURL, scope, roleDefinitionName, parameters, options)
	if err != nil {
		return RoleDefinitionsCreateOrUpdateResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return RoleDefinitionsCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return RoleDefinitionsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *roleDefinitionsClient) createOrUpdateCreateRequest(ctx context.Context, vaultBaseURL string, scope string, roleDefinitionName string, parameters RoleDefinitionCreateParameters, options *RoleDefinitionsCreateOrUpdateOptions) (*policy.Request, error) {
	host := "{vaultBaseUrl}"
	host = strings.ReplaceAll(host, "{vaultBaseUrl}", vaultBaseURL)
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleDefinitions/{roleDefinitionName}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleDefinitionName == "" {
		return nil, errors.New("parameter roleDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleDefinitionName}", url.PathEscape(roleDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "7.2")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *roleDefinitionsClient) createOrUpdateHandleResponse(resp *http.Response) (RoleDefinitionsCreateOrUpdateResponse, error) {
	result := RoleDefinitionsCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleDefinition); err != nil {
		return RoleDefinitionsCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *roleDefinitionsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := KeyVaultError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Deletes a custom role definition.
// If the operation fails it returns the *KeyVaultError error type.
func (client *roleDefinitionsClient) Delete(ctx context.Context, vaultBaseURL string, scope string, roleDefinitionName string, options *RoleDefinitionsDeleteOptions) (RoleDefinitionsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, vaultBaseURL, scope, roleDefinitionName, options)
	if err != nil {
		return RoleDefinitionsDeleteResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return RoleDefinitionsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoleDefinitionsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return client.deleteHandleResponse(resp)
}

// deleteCreateRequest creates the Delete request.
func (client *roleDefinitionsClient) deleteCreateRequest(ctx context.Context, vaultBaseURL string, scope string, roleDefinitionName string, options *RoleDefinitionsDeleteOptions) (*policy.Request, error) {
	host := "{vaultBaseUrl}"
	host = strings.ReplaceAll(host, "{vaultBaseUrl}", vaultBaseURL)
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleDefinitions/{roleDefinitionName}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleDefinitionName == "" {
		return nil, errors.New("parameter roleDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleDefinitionName}", url.PathEscape(roleDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "7.2")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *roleDefinitionsClient) deleteHandleResponse(resp *http.Response) (RoleDefinitionsDeleteResponse, error) {
	result := RoleDefinitionsDeleteResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleDefinition); err != nil {
		return RoleDefinitionsDeleteResponse{}, err
	}
	return result, nil
}

// deleteHandleError handles the Delete error response.
func (client *roleDefinitionsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := KeyVaultError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Get the specified role definition.
// If the operation fails it returns the *KeyVaultError error type.
func (client *roleDefinitionsClient) Get(ctx context.Context, vaultBaseURL string, scope string, roleDefinitionName string, options *RoleDefinitionsGetOptions) (RoleDefinitionsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, vaultBaseURL, scope, roleDefinitionName, options)
	if err != nil {
		return RoleDefinitionsGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return RoleDefinitionsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoleDefinitionsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *roleDefinitionsClient) getCreateRequest(ctx context.Context, vaultBaseURL string, scope string, roleDefinitionName string, options *RoleDefinitionsGetOptions) (*policy.Request, error) {
	host := "{vaultBaseUrl}"
	host = strings.ReplaceAll(host, "{vaultBaseUrl}", vaultBaseURL)
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleDefinitions/{roleDefinitionName}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleDefinitionName == "" {
		return nil, errors.New("parameter roleDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleDefinitionName}", url.PathEscape(roleDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "7.2")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *roleDefinitionsClient) getHandleResponse(resp *http.Response) (RoleDefinitionsGetResponse, error) {
	result := RoleDefinitionsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleDefinition); err != nil {
		return RoleDefinitionsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *roleDefinitionsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := KeyVaultError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Get all role definitions that are applicable at scope and above.
// If the operation fails it returns the *KeyVaultError error type.
func (client *roleDefinitionsClient) List(vaultBaseURL string, scope string, options *RoleDefinitionsListOptions) *RoleDefinitionsListPager {
	return &RoleDefinitionsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, vaultBaseURL, scope, options)
		},
		advancer: func(ctx context.Context, resp RoleDefinitionsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RoleDefinitionListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *roleDefinitionsClient) listCreateRequest(ctx context.Context, vaultBaseURL string, scope string, options *RoleDefinitionsListOptions) (*policy.Request, error) {
	host := "{vaultBaseUrl}"
	host = strings.ReplaceAll(host, "{vaultBaseUrl}", vaultBaseURL)
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleDefinitions"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "7.2")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *roleDefinitionsClient) listHandleResponse(resp *http.Response) (RoleDefinitionsListResponse, error) {
	result := RoleDefinitionsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleDefinitionListResult); err != nil {
		return RoleDefinitionsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *roleDefinitionsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := KeyVaultError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
