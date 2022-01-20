//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmixedreality

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// ObjectAnchorsAccountsClientListByResourceGroupPager provides operations for iterating over paged responses.
type ObjectAnchorsAccountsClientListByResourceGroupPager struct {
	client    *ObjectAnchorsAccountsClient
	current   ObjectAnchorsAccountsClientListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ObjectAnchorsAccountsClientListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ObjectAnchorsAccountsClientListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ObjectAnchorsAccountsClientListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ObjectAnchorsAccountPage.NextLink == nil || len(*p.current.ObjectAnchorsAccountPage.NextLink) == 0 {
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

// PageResponse returns the current ObjectAnchorsAccountsClientListByResourceGroupResponse page.
func (p *ObjectAnchorsAccountsClientListByResourceGroupPager) PageResponse() ObjectAnchorsAccountsClientListByResourceGroupResponse {
	return p.current
}

// ObjectAnchorsAccountsClientListBySubscriptionPager provides operations for iterating over paged responses.
type ObjectAnchorsAccountsClientListBySubscriptionPager struct {
	client    *ObjectAnchorsAccountsClient
	current   ObjectAnchorsAccountsClientListBySubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ObjectAnchorsAccountsClientListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ObjectAnchorsAccountsClientListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ObjectAnchorsAccountsClientListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ObjectAnchorsAccountPage.NextLink == nil || len(*p.current.ObjectAnchorsAccountPage.NextLink) == 0 {
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

// PageResponse returns the current ObjectAnchorsAccountsClientListBySubscriptionResponse page.
func (p *ObjectAnchorsAccountsClientListBySubscriptionPager) PageResponse() ObjectAnchorsAccountsClientListBySubscriptionResponse {
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
		if p.current.OperationPage.NextLink == nil || len(*p.current.OperationPage.NextLink) == 0 {
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

// RemoteRenderingAccountsClientListByResourceGroupPager provides operations for iterating over paged responses.
type RemoteRenderingAccountsClientListByResourceGroupPager struct {
	client    *RemoteRenderingAccountsClient
	current   RemoteRenderingAccountsClientListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, RemoteRenderingAccountsClientListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *RemoteRenderingAccountsClientListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *RemoteRenderingAccountsClientListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RemoteRenderingAccountPage.NextLink == nil || len(*p.current.RemoteRenderingAccountPage.NextLink) == 0 {
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

// PageResponse returns the current RemoteRenderingAccountsClientListByResourceGroupResponse page.
func (p *RemoteRenderingAccountsClientListByResourceGroupPager) PageResponse() RemoteRenderingAccountsClientListByResourceGroupResponse {
	return p.current
}

// RemoteRenderingAccountsClientListBySubscriptionPager provides operations for iterating over paged responses.
type RemoteRenderingAccountsClientListBySubscriptionPager struct {
	client    *RemoteRenderingAccountsClient
	current   RemoteRenderingAccountsClientListBySubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, RemoteRenderingAccountsClientListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *RemoteRenderingAccountsClientListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *RemoteRenderingAccountsClientListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RemoteRenderingAccountPage.NextLink == nil || len(*p.current.RemoteRenderingAccountPage.NextLink) == 0 {
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

// PageResponse returns the current RemoteRenderingAccountsClientListBySubscriptionResponse page.
func (p *RemoteRenderingAccountsClientListBySubscriptionPager) PageResponse() RemoteRenderingAccountsClientListBySubscriptionResponse {
	return p.current
}

// SpatialAnchorsAccountsClientListByResourceGroupPager provides operations for iterating over paged responses.
type SpatialAnchorsAccountsClientListByResourceGroupPager struct {
	client    *SpatialAnchorsAccountsClient
	current   SpatialAnchorsAccountsClientListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SpatialAnchorsAccountsClientListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SpatialAnchorsAccountsClientListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SpatialAnchorsAccountsClientListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SpatialAnchorsAccountPage.NextLink == nil || len(*p.current.SpatialAnchorsAccountPage.NextLink) == 0 {
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

// PageResponse returns the current SpatialAnchorsAccountsClientListByResourceGroupResponse page.
func (p *SpatialAnchorsAccountsClientListByResourceGroupPager) PageResponse() SpatialAnchorsAccountsClientListByResourceGroupResponse {
	return p.current
}

// SpatialAnchorsAccountsClientListBySubscriptionPager provides operations for iterating over paged responses.
type SpatialAnchorsAccountsClientListBySubscriptionPager struct {
	client    *SpatialAnchorsAccountsClient
	current   SpatialAnchorsAccountsClientListBySubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SpatialAnchorsAccountsClientListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SpatialAnchorsAccountsClientListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SpatialAnchorsAccountsClientListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SpatialAnchorsAccountPage.NextLink == nil || len(*p.current.SpatialAnchorsAccountPage.NextLink) == 0 {
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

// PageResponse returns the current SpatialAnchorsAccountsClientListBySubscriptionResponse page.
func (p *SpatialAnchorsAccountsClientListBySubscriptionPager) PageResponse() SpatialAnchorsAccountsClientListBySubscriptionResponse {
	return p.current
}
