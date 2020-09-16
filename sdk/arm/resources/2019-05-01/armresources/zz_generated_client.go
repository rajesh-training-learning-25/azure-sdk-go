// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresources

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

const scope = "https://management.azure.com//.default"
const telemetryInfo = "azsdk-go-armresources/<version>"

// ClientOptions contains configuration settings for the default client's pipeline.
type ClientOptions struct {
	// HTTPClient sets the transport for making HTTP requests.
	HTTPClient azcore.Transport
	// LogOptions configures the built-in request logging policy behavior.
	LogOptions azcore.RequestLogOptions
	// Retry configures the built-in retry policy behavior.
	Retry azcore.RetryOptions
	// Telemetry configures the built-in telemetry policy behavior.
	Telemetry azcore.TelemetryOptions
	// RegisterRPOptions configures the built-in RP registration policy behavior.
	RegisterRPOptions armcore.RegistrationOptions
}

// DefaultClientOptions creates a ClientOptions type initialized with default values.
func DefaultClientOptions() ClientOptions {
	return ClientOptions{
		HTTPClient:        azcore.DefaultHTTPClientTransport(),
		Retry:             azcore.DefaultRetryOptions(),
		RegisterRPOptions: armcore.DefaultRegistrationOptions(),
	}
}

func (c *ClientOptions) telemetryOptions() azcore.TelemetryOptions {
	to := c.Telemetry
	if to.Value == "" {
		to.Value = telemetryInfo
	} else {
		to.Value = fmt.Sprintf("%s %s", telemetryInfo, to.Value)
	}
	return to
}

// Client - Provides operations for working with resources and resource groups.
type Client struct {
	u string
	p azcore.Pipeline
}

// DefaultEndpoint is the default service endpoint.
const DefaultEndpoint = "https://management.azure.com"

// NewDefaultClient creates an instance of the Client type using the DefaultEndpoint.
func NewDefaultClient(cred azcore.Credential, options *ClientOptions) *Client {
	return NewClient(DefaultEndpoint, cred, options)
}

// NewClient creates an instance of the Client type with the specified endpoint.
func NewClient(endpoint string, cred azcore.Credential, options *ClientOptions) *Client {
	if options == nil {
		o := DefaultClientOptions()
		options = &o
	}
	policies := []azcore.Policy{
		azcore.NewTelemetryPolicy(options.telemetryOptions()),
		azcore.NewUniqueRequestIDPolicy(),
	}
	policies = append(policies, armcore.NewRPRegistrationPolicy(cred, &options.RegisterRPOptions))
	policies = append(policies,
		azcore.NewRetryPolicy(&options.Retry),
		cred.AuthenticationPolicy(azcore.AuthenticationPolicyOptions{Options: azcore.TokenRequestOptions{Scopes: []string{scope}}}),
		azcore.NewRequestLogPolicy(options.LogOptions))
	p := azcore.NewPipeline(options.HTTPClient, policies...)
	return NewClientWithPipeline(endpoint, p)
}

// NewClientWithPipeline creates an instance of the Client type with the specified endpoint and pipeline.
func NewClientWithPipeline(endpoint string, p azcore.Pipeline) *Client {
	return &Client{u: endpoint, p: p}
}
