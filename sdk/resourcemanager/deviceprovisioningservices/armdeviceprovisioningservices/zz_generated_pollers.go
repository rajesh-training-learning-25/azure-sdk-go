//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdeviceprovisioningservices

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// IotDpsResourceCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type IotDpsResourceCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *IotDpsResourceCreateOrUpdatePoller) Done() bool {
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
func (p *IotDpsResourceCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final IotDpsResourceCreateOrUpdateResponse will be returned.
func (p *IotDpsResourceCreateOrUpdatePoller) FinalResponse(ctx context.Context) (IotDpsResourceCreateOrUpdateResponse, error) {
	respType := IotDpsResourceCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.ProvisioningServiceDescription)
	if err != nil {
		return IotDpsResourceCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *IotDpsResourceCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// IotDpsResourceCreateOrUpdatePrivateEndpointConnectionPoller provides polling facilities until the operation reaches a terminal state.
type IotDpsResourceCreateOrUpdatePrivateEndpointConnectionPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *IotDpsResourceCreateOrUpdatePrivateEndpointConnectionPoller) Done() bool {
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
func (p *IotDpsResourceCreateOrUpdatePrivateEndpointConnectionPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final IotDpsResourceCreateOrUpdatePrivateEndpointConnectionResponse will be returned.
func (p *IotDpsResourceCreateOrUpdatePrivateEndpointConnectionPoller) FinalResponse(ctx context.Context) (IotDpsResourceCreateOrUpdatePrivateEndpointConnectionResponse, error) {
	respType := IotDpsResourceCreateOrUpdatePrivateEndpointConnectionResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.PrivateEndpointConnection)
	if err != nil {
		return IotDpsResourceCreateOrUpdatePrivateEndpointConnectionResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *IotDpsResourceCreateOrUpdatePrivateEndpointConnectionPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// IotDpsResourceDeletePoller provides polling facilities until the operation reaches a terminal state.
type IotDpsResourceDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *IotDpsResourceDeletePoller) Done() bool {
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
func (p *IotDpsResourceDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final IotDpsResourceDeleteResponse will be returned.
func (p *IotDpsResourceDeletePoller) FinalResponse(ctx context.Context) (IotDpsResourceDeleteResponse, error) {
	respType := IotDpsResourceDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return IotDpsResourceDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *IotDpsResourceDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// IotDpsResourceDeletePrivateEndpointConnectionPoller provides polling facilities until the operation reaches a terminal state.
type IotDpsResourceDeletePrivateEndpointConnectionPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *IotDpsResourceDeletePrivateEndpointConnectionPoller) Done() bool {
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
func (p *IotDpsResourceDeletePrivateEndpointConnectionPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final IotDpsResourceDeletePrivateEndpointConnectionResponse will be returned.
func (p *IotDpsResourceDeletePrivateEndpointConnectionPoller) FinalResponse(ctx context.Context) (IotDpsResourceDeletePrivateEndpointConnectionResponse, error) {
	respType := IotDpsResourceDeletePrivateEndpointConnectionResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.PrivateEndpointConnection)
	if err != nil {
		return IotDpsResourceDeletePrivateEndpointConnectionResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *IotDpsResourceDeletePrivateEndpointConnectionPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// IotDpsResourceUpdatePoller provides polling facilities until the operation reaches a terminal state.
type IotDpsResourceUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *IotDpsResourceUpdatePoller) Done() bool {
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
func (p *IotDpsResourceUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final IotDpsResourceUpdateResponse will be returned.
func (p *IotDpsResourceUpdatePoller) FinalResponse(ctx context.Context) (IotDpsResourceUpdateResponse, error) {
	respType := IotDpsResourceUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.ProvisioningServiceDescription)
	if err != nil {
		return IotDpsResourceUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *IotDpsResourceUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}
