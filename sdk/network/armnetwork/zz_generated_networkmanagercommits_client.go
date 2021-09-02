//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// NetworkManagerCommitsClient contains the methods for the NetworkManagerCommits group.
// Don't use this type directly, use NewNetworkManagerCommitsClient() instead.
type NetworkManagerCommitsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewNetworkManagerCommitsClient creates a new instance of NetworkManagerCommitsClient with the specified values.
func NewNetworkManagerCommitsClient(con *arm.Connection, subscriptionID string) *NetworkManagerCommitsClient {
	return &NetworkManagerCommitsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Post - Post a Network Manager Commit.
// If the operation fails it returns the *CloudErrorAutoGenerated error type.
func (client *NetworkManagerCommitsClient) Post(ctx context.Context, resourceGroupName string, networkManagerName string, parameters NetworkManagerCommit, options *NetworkManagerCommitsPostOptions) (NetworkManagerCommitsPostResponse, error) {
	req, err := client.postCreateRequest(ctx, resourceGroupName, networkManagerName, parameters, options)
	if err != nil {
		return NetworkManagerCommitsPostResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NetworkManagerCommitsPostResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return NetworkManagerCommitsPostResponse{}, client.postHandleError(resp)
	}
	return client.postHandleResponse(resp)
}

// postCreateRequest creates the Post request.
func (client *NetworkManagerCommitsClient) postCreateRequest(ctx context.Context, resourceGroupName string, networkManagerName string, parameters NetworkManagerCommit, options *NetworkManagerCommitsPostOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/commit"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkManagerName == "" {
		return nil, errors.New("parameter networkManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkManagerName}", url.PathEscape(networkManagerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// postHandleResponse handles the Post response.
func (client *NetworkManagerCommitsClient) postHandleResponse(resp *http.Response) (NetworkManagerCommitsPostResponse, error) {
	result := NetworkManagerCommitsPostResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkManagerCommit); err != nil {
		return NetworkManagerCommitsPostResponse{}, err
	}
	return result, nil
}

// postHandleError handles the Post error response.
func (client *NetworkManagerCommitsClient) postHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudErrorAutoGenerated{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
