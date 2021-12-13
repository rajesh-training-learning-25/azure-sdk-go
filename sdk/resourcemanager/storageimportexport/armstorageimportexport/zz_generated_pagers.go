//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorageimportexport

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// JobsListByResourceGroupPager provides operations for iterating over paged responses.
type JobsListByResourceGroupPager struct {
	client    *JobsClient
	current   JobsListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, JobsListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *JobsListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *JobsListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ListJobsResponse.NextLink == nil || len(*p.current.ListJobsResponse.NextLink) == 0 {
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
		p.err = p.client.listByResourceGroupHandleError(resp)
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

// PageResponse returns the current JobsListByResourceGroupResponse page.
func (p *JobsListByResourceGroupPager) PageResponse() JobsListByResourceGroupResponse {
	return p.current
}

// JobsListBySubscriptionPager provides operations for iterating over paged responses.
type JobsListBySubscriptionPager struct {
	client    *JobsClient
	current   JobsListBySubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, JobsListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *JobsListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *JobsListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ListJobsResponse.NextLink == nil || len(*p.current.ListJobsResponse.NextLink) == 0 {
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
		p.err = p.client.listBySubscriptionHandleError(resp)
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

// PageResponse returns the current JobsListBySubscriptionResponse page.
func (p *JobsListBySubscriptionPager) PageResponse() JobsListBySubscriptionResponse {
	return p.current
}
