//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcognitiveservices

import (
	"context"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// AccountsCreatePoller provides polling facilities until the operation reaches a terminal state.
type AccountsCreatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *AccountsCreatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *AccountsCreatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final AccountsCreateResponse will be returned.
func (p *AccountsCreatePoller) FinalResponse(ctx context.Context) (AccountsCreateResponse, error) {
	respType := AccountsCreateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.Account)
	if err != nil {
		return AccountsCreateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *AccountsCreatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// AccountsDeletePoller provides polling facilities until the operation reaches a terminal state.
type AccountsDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *AccountsDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *AccountsDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final AccountsDeleteResponse will be returned.
func (p *AccountsDeletePoller) FinalResponse(ctx context.Context) (AccountsDeleteResponse, error) {
	respType := AccountsDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return AccountsDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *AccountsDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// AccountsUpdatePoller provides polling facilities until the operation reaches a terminal state.
type AccountsUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *AccountsUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *AccountsUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final AccountsUpdateResponse will be returned.
func (p *AccountsUpdatePoller) FinalResponse(ctx context.Context) (AccountsUpdateResponse, error) {
	respType := AccountsUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.Account)
	if err != nil {
		return AccountsUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *AccountsUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// DeletedAccountsPurgePoller provides polling facilities until the operation reaches a terminal state.
type DeletedAccountsPurgePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *DeletedAccountsPurgePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *DeletedAccountsPurgePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final DeletedAccountsPurgeResponse will be returned.
func (p *DeletedAccountsPurgePoller) FinalResponse(ctx context.Context) (DeletedAccountsPurgeResponse, error) {
	respType := DeletedAccountsPurgeResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return DeletedAccountsPurgeResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *DeletedAccountsPurgePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// PrivateEndpointConnectionsCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type PrivateEndpointConnectionsCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *PrivateEndpointConnectionsCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *PrivateEndpointConnectionsCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final PrivateEndpointConnectionsCreateOrUpdateResponse will be returned.
func (p *PrivateEndpointConnectionsCreateOrUpdatePoller) FinalResponse(ctx context.Context) (PrivateEndpointConnectionsCreateOrUpdateResponse, error) {
	respType := PrivateEndpointConnectionsCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.PrivateEndpointConnection)
	if err != nil {
		return PrivateEndpointConnectionsCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *PrivateEndpointConnectionsCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// PrivateEndpointConnectionsDeletePoller provides polling facilities until the operation reaches a terminal state.
type PrivateEndpointConnectionsDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *PrivateEndpointConnectionsDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *PrivateEndpointConnectionsDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final PrivateEndpointConnectionsDeleteResponse will be returned.
func (p *PrivateEndpointConnectionsDeletePoller) FinalResponse(ctx context.Context) (PrivateEndpointConnectionsDeleteResponse, error) {
	respType := PrivateEndpointConnectionsDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return PrivateEndpointConnectionsDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *PrivateEndpointConnectionsDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}
