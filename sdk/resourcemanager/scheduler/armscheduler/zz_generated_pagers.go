//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armscheduler

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// JobCollectionsClientListByResourceGroupPager provides operations for iterating over paged responses.
type JobCollectionsClientListByResourceGroupPager struct {
	client    *JobCollectionsClient
	current   JobCollectionsClientListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, JobCollectionsClientListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *JobCollectionsClientListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *JobCollectionsClientListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.JobCollectionListResult.NextLink == nil || len(*p.current.JobCollectionListResult.NextLink) == 0 {
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

// PageResponse returns the current JobCollectionsClientListByResourceGroupResponse page.
func (p *JobCollectionsClientListByResourceGroupPager) PageResponse() JobCollectionsClientListByResourceGroupResponse {
	return p.current
}

// JobCollectionsClientListBySubscriptionPager provides operations for iterating over paged responses.
type JobCollectionsClientListBySubscriptionPager struct {
	client    *JobCollectionsClient
	current   JobCollectionsClientListBySubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, JobCollectionsClientListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *JobCollectionsClientListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *JobCollectionsClientListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.JobCollectionListResult.NextLink == nil || len(*p.current.JobCollectionListResult.NextLink) == 0 {
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

// PageResponse returns the current JobCollectionsClientListBySubscriptionResponse page.
func (p *JobCollectionsClientListBySubscriptionPager) PageResponse() JobCollectionsClientListBySubscriptionResponse {
	return p.current
}

// JobsClientListJobHistoryPager provides operations for iterating over paged responses.
type JobsClientListJobHistoryPager struct {
	client    *JobsClient
	current   JobsClientListJobHistoryResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, JobsClientListJobHistoryResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *JobsClientListJobHistoryPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *JobsClientListJobHistoryPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.JobHistoryListResult.NextLink == nil || len(*p.current.JobHistoryListResult.NextLink) == 0 {
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
	result, err := p.client.listJobHistoryHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current JobsClientListJobHistoryResponse page.
func (p *JobsClientListJobHistoryPager) PageResponse() JobsClientListJobHistoryResponse {
	return p.current
}

// JobsClientListPager provides operations for iterating over paged responses.
type JobsClientListPager struct {
	client    *JobsClient
	current   JobsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, JobsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *JobsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *JobsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.JobListResult.NextLink == nil || len(*p.current.JobListResult.NextLink) == 0 {
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

// PageResponse returns the current JobsClientListResponse page.
func (p *JobsClientListPager) PageResponse() JobsClientListResponse {
	return p.current
}
