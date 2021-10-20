//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armchaos

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// OperationsClient contains the methods for the Operations group.
// Don't use this type directly, use NewOperationsClient() instead.
type OperationsClient struct {
	ep string
	pl runtime.Pipeline
}

// NewOperationsClient creates a new instance of OperationsClient with the specified values.
func NewOperationsClient(con *arm.Connection) *OperationsClient {
	return &OperationsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version)}
}

// ListAll - Get a list all available Operations.
// If the operation fails it returns the *ErrorResponse error type.
func (client *OperationsClient) ListAll(options *OperationsListAllOptions) *OperationsListAllPager {
	return &OperationsListAllPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listAllCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp OperationsListAllResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.OperationListResult.NextLink)
		},
	}
}

// listAllCreateRequest creates the ListAll request.
func (client *OperationsClient) listAllCreateRequest(ctx context.Context, options *OperationsListAllOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Chaos/operations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *OperationsClient) listAllHandleResponse(resp *http.Response) (OperationsListAllResponse, error) {
	result := OperationsListAllResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationListResult); err != nil {
		return OperationsListAllResponse{}, err
	}
	return result, nil
}

// listAllHandleError handles the ListAll error response.
func (client *OperationsClient) listAllHandleError(resp *http.Response) error {
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
