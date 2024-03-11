// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azcosmos

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	azruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/errorinfo"
)

type clientRetryPolicy struct {
	gem                    *globalEndpointManager
	useWriteEndpoint       bool
	retryCount             int
	sessionRetryCount      int
	preferredLocationIndex int
}

const maxRetryCount = 120
const defaultBackoff = 1

func (p *clientRetryPolicy) Do(req *policy.Request) (*http.Response, error) {
	p.resetPolicyCounters()
	o := pipelineRequestOptions{}
	if !req.OperationValue(&o) {
		return nil, fmt.Errorf("Failed to obtain request options, please check request being sent: %s", req.Body())
	}
	resolvedEndpoint := p.gem.ResolveServiceEndpoint(p.retryCount, o.isWriteOperation, p.useWriteEndpoint)
	req.Raw().Host = resolvedEndpoint.Host
	req.Raw().URL.Host = resolvedEndpoint.Host
	response, err := req.Next() // err can happen in weird scenarios (connectivity, etc)
	if err != nil {
		return nil, err
	}
	subStatus := response.Header.Get(cosmosHeaderSubstatus)
	if p.shouldRetryStatus(response.StatusCode, subStatus) {
		for {
			p.useWriteEndpoint = false
			subStatus = response.Header.Get(cosmosHeaderSubstatus)
			if p.shouldRetryStatus(response.StatusCode, subStatus) {
				if response.StatusCode == http.StatusForbidden {
					if !p.attemptRetryOnEndpointFailure(req, o.isWriteOperation) {
						return nil, errorinfo.NonRetriableError(azruntime.NewResponseErrorWithErrorCode(response, response.Status))
					}
				} else if response.StatusCode == http.StatusNotFound {
					if !p.attemptRetryOnSessionUnavailable(req, o.isWriteOperation) {
						return nil, errorinfo.NonRetriableError(azruntime.NewResponseErrorWithErrorCode(response, response.Status))
					}
				} else if response.StatusCode == http.StatusServiceUnavailable {
					if !p.attemptRetryOnServiceUnavailable(req, o.isWriteOperation) {
						return nil, errorinfo.NonRetriableError(azruntime.NewResponseErrorWithErrorCode(response, response.Status))
					}
				}
			} else {
				break
			}
			err = req.RewindBody()
			if err != nil {
				return response, err
			}
			resolvedEndpoint := p.gem.ResolveServiceEndpoint(p.retryCount, o.isWriteOperation, p.useWriteEndpoint)
			req.Raw().Host = resolvedEndpoint.Host
			req.Raw().URL.Host = resolvedEndpoint.Host
			p.retryCount += 1
			response, err = req.Next()
		}
	}
	return response, err
}

func (p *clientRetryPolicy) shouldRetryStatus(status int, subStatus string) (shouldRetry bool) {
	if (status == http.StatusForbidden && (subStatus == subStatusWriteForbidden || subStatus == subStatusDatabaseAccountNotFound)) ||
		(status == http.StatusNotFound && subStatus == subStatusReadSessionNotAvailable) ||
		(status == http.StatusServiceUnavailable) {
		return true
	}
	return false
}

func (p *clientRetryPolicy) attemptRetryOnEndpointFailure(req *policy.Request, isWriteOperation bool) bool {
	if (p.retryCount > maxRetryCount) || !p.gem.locationCache.enableCrossRegionRetries {
		return false
	}
	if isWriteOperation {
		p.gem.MarkEndpointUnavailableForWrite(*req.Raw().URL)
		p.gem.Update(req.Raw().Context(), true)
	} else {
		p.gem.MarkEndpointUnavailableForRead(*req.Raw().URL)
	}
	time.Sleep(defaultBackoff * time.Second)
	return true
}

func (p *clientRetryPolicy) attemptRetryOnSessionUnavailable(req *policy.Request, isWriteOperation bool) bool {
	if p.gem.CanUseMultipleWriteLocations() {
		endpoints := []string{}
		if isWriteOperation {
			endpoints = p.gem.locationCache.locationInfo.availWriteLocations
		} else {
			endpoints = p.gem.locationCache.locationInfo.availReadLocations
		}
		if p.sessionRetryCount >= len(endpoints) {
			return false
		}
	} else {
		if p.sessionRetryCount > 0 {
			return false
		}
		p.useWriteEndpoint = true
	}
	p.sessionRetryCount += 1
	return true
}

func (p *clientRetryPolicy) attemptRetryOnServiceUnavailable(req *policy.Request, isWriteOperation bool) bool {
	if !p.gem.locationCache.enableCrossRegionRetries || p.preferredLocationIndex >= len(p.gem.preferredLocations) {
		return false
	}
	if isWriteOperation && !p.gem.CanUseMultipleWriteLocations() {
		return false
	}
	p.preferredLocationIndex += 1
	return true
}

func (p *clientRetryPolicy) resetPolicyCounters() {
	p.retryCount = 0
	p.sessionRetryCount = 0
	p.preferredLocationIndex = 0
}
