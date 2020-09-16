// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"strings"
)

// MetricNamespacesOperations contains the methods for the MetricNamespaces group.
type MetricNamespacesOperations interface {
	// List - Lists the metric namespaces for the resource.
	List(ctx context.Context, resourceUri string, metricNamespacesListOptions *MetricNamespacesListOptions) (*MetricNamespaceCollectionResponse, error)
}

// MetricNamespacesClient implements the MetricNamespacesOperations interface.
// Don't use this type directly, use NewMetricNamespacesClient() instead.
type MetricNamespacesClient struct {
	*Client
}

// NewMetricNamespacesClient creates a new instance of MetricNamespacesClient with the specified values.
func NewMetricNamespacesClient(c *Client) MetricNamespacesOperations {
	return &MetricNamespacesClient{Client: c}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *MetricNamespacesClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// List - Lists the metric namespaces for the resource.
func (client *MetricNamespacesClient) List(ctx context.Context, resourceUri string, metricNamespacesListOptions *MetricNamespacesListOptions) (*MetricNamespaceCollectionResponse, error) {
	req, err := client.ListCreateRequest(ctx, resourceUri, metricNamespacesListOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListHandleError(resp)
	}
	result, err := client.ListHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListCreateRequest creates the List request.
func (client *MetricNamespacesClient) ListCreateRequest(ctx context.Context, resourceUri string, metricNamespacesListOptions *MetricNamespacesListOptions) (*azcore.Request, error) {
	urlPath := "/{resourceUri}/providers/microsoft.insights/metricNamespaces"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceUri)
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2017-12-01-preview")
	if metricNamespacesListOptions != nil && metricNamespacesListOptions.StartTime != nil {
		query.Set("startTime", *metricNamespacesListOptions.StartTime)
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListHandleResponse handles the List response.
func (client *MetricNamespacesClient) ListHandleResponse(resp *azcore.Response) (*MetricNamespaceCollectionResponse, error) {
	result := MetricNamespaceCollectionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.MetricNamespaceCollection)
}

// ListHandleError handles the List error response.
func (client *MetricNamespacesClient) ListHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
