//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armworkloads

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
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID, credential: credential,
		options: options.Clone(),
	}, nil
}

// NewClient creates a new instance of Client.
func (c *ClientFactory) NewClient() *Client {
	subClient, _ := NewClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewMonitorsClient creates a new instance of MonitorsClient.
func (c *ClientFactory) NewMonitorsClient() *MonitorsClient {
	subClient, _ := NewMonitorsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

// NewProviderInstancesClient creates a new instance of ProviderInstancesClient.
func (c *ClientFactory) NewProviderInstancesClient() *ProviderInstancesClient {
	subClient, _ := NewProviderInstancesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewSAPApplicationServerInstancesClient creates a new instance of SAPApplicationServerInstancesClient.
func (c *ClientFactory) NewSAPApplicationServerInstancesClient() *SAPApplicationServerInstancesClient {
	subClient, _ := NewSAPApplicationServerInstancesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewSAPCentralInstancesClient creates a new instance of SAPCentralInstancesClient.
func (c *ClientFactory) NewSAPCentralInstancesClient() *SAPCentralInstancesClient {
	subClient, _ := NewSAPCentralInstancesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewSAPDatabaseInstancesClient creates a new instance of SAPDatabaseInstancesClient.
func (c *ClientFactory) NewSAPDatabaseInstancesClient() *SAPDatabaseInstancesClient {
	subClient, _ := NewSAPDatabaseInstancesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewSAPVirtualInstancesClient creates a new instance of SAPVirtualInstancesClient.
func (c *ClientFactory) NewSAPVirtualInstancesClient() *SAPVirtualInstancesClient {
	subClient, _ := NewSAPVirtualInstancesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewSapLandscapeMonitorClient creates a new instance of SapLandscapeMonitorClient.
func (c *ClientFactory) NewSapLandscapeMonitorClient() *SapLandscapeMonitorClient {
	subClient, _ := NewSapLandscapeMonitorClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
