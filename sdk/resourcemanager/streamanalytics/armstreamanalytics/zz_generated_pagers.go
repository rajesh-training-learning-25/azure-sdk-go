//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstreamanalytics

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// ClustersClientListByResourceGroupPager provides operations for iterating over paged responses.
type ClustersClientListByResourceGroupPager struct {
	client    *ClustersClient
	current   ClustersClientListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClustersClientListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ClustersClientListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ClustersClientListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ClusterListResult.NextLink == nil || len(*p.current.ClusterListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ClustersClientListByResourceGroupResponse page.
func (p *ClustersClientListByResourceGroupPager) PageResponse() ClustersClientListByResourceGroupResponse {
	return p.current
}

// ClustersClientListBySubscriptionPager provides operations for iterating over paged responses.
type ClustersClientListBySubscriptionPager struct {
	client    *ClustersClient
	current   ClustersClientListBySubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClustersClientListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ClustersClientListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ClustersClientListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ClusterListResult.NextLink == nil || len(*p.current.ClusterListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listBySubscriptionHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ClustersClientListBySubscriptionResponse page.
func (p *ClustersClientListBySubscriptionPager) PageResponse() ClustersClientListBySubscriptionResponse {
	return p.current
}

// ClustersClientListStreamingJobsPager provides operations for iterating over paged responses.
type ClustersClientListStreamingJobsPager struct {
	client    *ClustersClient
	current   ClustersClientListStreamingJobsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClustersClientListStreamingJobsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ClustersClientListStreamingJobsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ClustersClientListStreamingJobsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ClusterJobListResult.NextLink == nil || len(*p.current.ClusterJobListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listStreamingJobsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ClustersClientListStreamingJobsResponse page.
func (p *ClustersClientListStreamingJobsPager) PageResponse() ClustersClientListStreamingJobsResponse {
	return p.current
}

// FunctionsClientListByStreamingJobPager provides operations for iterating over paged responses.
type FunctionsClientListByStreamingJobPager struct {
	client    *FunctionsClient
	current   FunctionsClientListByStreamingJobResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, FunctionsClientListByStreamingJobResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *FunctionsClientListByStreamingJobPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *FunctionsClientListByStreamingJobPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.FunctionListResult.NextLink == nil || len(*p.current.FunctionListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listByStreamingJobHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current FunctionsClientListByStreamingJobResponse page.
func (p *FunctionsClientListByStreamingJobPager) PageResponse() FunctionsClientListByStreamingJobResponse {
	return p.current
}

// InputsClientListByStreamingJobPager provides operations for iterating over paged responses.
type InputsClientListByStreamingJobPager struct {
	client    *InputsClient
	current   InputsClientListByStreamingJobResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, InputsClientListByStreamingJobResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *InputsClientListByStreamingJobPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *InputsClientListByStreamingJobPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.InputListResult.NextLink == nil || len(*p.current.InputListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listByStreamingJobHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current InputsClientListByStreamingJobResponse page.
func (p *InputsClientListByStreamingJobPager) PageResponse() InputsClientListByStreamingJobResponse {
	return p.current
}

// OperationsClientListPager provides operations for iterating over paged responses.
type OperationsClientListPager struct {
	client    *OperationsClient
	current   OperationsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, OperationsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *OperationsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *OperationsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OperationListResult.NextLink == nil || len(*p.current.OperationListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current OperationsClientListResponse page.
func (p *OperationsClientListPager) PageResponse() OperationsClientListResponse {
	return p.current
}

// OutputsClientListByStreamingJobPager provides operations for iterating over paged responses.
type OutputsClientListByStreamingJobPager struct {
	client    *OutputsClient
	current   OutputsClientListByStreamingJobResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, OutputsClientListByStreamingJobResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *OutputsClientListByStreamingJobPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *OutputsClientListByStreamingJobPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OutputListResult.NextLink == nil || len(*p.current.OutputListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listByStreamingJobHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current OutputsClientListByStreamingJobResponse page.
func (p *OutputsClientListByStreamingJobPager) PageResponse() OutputsClientListByStreamingJobResponse {
	return p.current
}

// PrivateEndpointsClientListByClusterPager provides operations for iterating over paged responses.
type PrivateEndpointsClientListByClusterPager struct {
	client    *PrivateEndpointsClient
	current   PrivateEndpointsClientListByClusterResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PrivateEndpointsClientListByClusterResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PrivateEndpointsClientListByClusterPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PrivateEndpointsClientListByClusterPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PrivateEndpointListResult.NextLink == nil || len(*p.current.PrivateEndpointListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listByClusterHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PrivateEndpointsClientListByClusterResponse page.
func (p *PrivateEndpointsClientListByClusterPager) PageResponse() PrivateEndpointsClientListByClusterResponse {
	return p.current
}

// StreamingJobsClientListByResourceGroupPager provides operations for iterating over paged responses.
type StreamingJobsClientListByResourceGroupPager struct {
	client    *StreamingJobsClient
	current   StreamingJobsClientListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, StreamingJobsClientListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *StreamingJobsClientListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *StreamingJobsClientListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.StreamingJobListResult.NextLink == nil || len(*p.current.StreamingJobListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current StreamingJobsClientListByResourceGroupResponse page.
func (p *StreamingJobsClientListByResourceGroupPager) PageResponse() StreamingJobsClientListByResourceGroupResponse {
	return p.current
}

// StreamingJobsClientListPager provides operations for iterating over paged responses.
type StreamingJobsClientListPager struct {
	client    *StreamingJobsClient
	current   StreamingJobsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, StreamingJobsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *StreamingJobsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *StreamingJobsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.StreamingJobListResult.NextLink == nil || len(*p.current.StreamingJobListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current StreamingJobsClientListResponse page.
func (p *StreamingJobsClientListPager) PageResponse() StreamingJobsClientListResponse {
	return p.current
}
