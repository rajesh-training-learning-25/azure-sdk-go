//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armfrontdoor

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	credential     azcore.TokenCredential
	options        *arm.ClientOptions
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(moduleName+".ClientFactory", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID, credential: credential,
		options: options.Clone(),
	}, nil
}

func (c *ClientFactory) NewPoliciesClient() *PoliciesClient {
	subClient, _ := NewPoliciesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewManagedRuleSetsClient() *ManagedRuleSetsClient {
	subClient, _ := NewManagedRuleSetsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewNameAvailabilityClient() *NameAvailabilityClient {
	subClient, _ := NewNameAvailabilityClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewNameAvailabilityWithSubscriptionClient() *NameAvailabilityWithSubscriptionClient {
	subClient, _ := NewNameAvailabilityWithSubscriptionClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewFrontDoorsClient() *FrontDoorsClient {
	subClient, _ := NewFrontDoorsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewFrontendEndpointsClient() *FrontendEndpointsClient {
	subClient, _ := NewFrontendEndpointsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewEndpointsClient() *EndpointsClient {
	subClient, _ := NewEndpointsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRulesEnginesClient() *RulesEnginesClient {
	subClient, _ := NewRulesEnginesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewNetworkExperimentProfilesClient() *NetworkExperimentProfilesClient {
	subClient, _ := NewNetworkExperimentProfilesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPreconfiguredEndpointsClient() *PreconfiguredEndpointsClient {
	subClient, _ := NewPreconfiguredEndpointsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewExperimentsClient() *ExperimentsClient {
	subClient, _ := NewExperimentsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewReportsClient() *ReportsClient {
	subClient, _ := NewReportsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
