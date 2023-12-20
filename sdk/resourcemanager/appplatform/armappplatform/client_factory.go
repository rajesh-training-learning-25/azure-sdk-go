//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappplatform

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
//   - subscriptionID - Gets subscription ID which uniquely identify the Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
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

// NewAPIPortalCustomDomainsClient creates a new instance of APIPortalCustomDomainsClient.
func (c *ClientFactory) NewAPIPortalCustomDomainsClient() *APIPortalCustomDomainsClient {
	subClient, _ := NewAPIPortalCustomDomainsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewAPIPortalsClient creates a new instance of APIPortalsClient.
func (c *ClientFactory) NewAPIPortalsClient() *APIPortalsClient {
	subClient, _ := NewAPIPortalsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewApplicationAcceleratorsClient creates a new instance of ApplicationAcceleratorsClient.
func (c *ClientFactory) NewApplicationAcceleratorsClient() *ApplicationAcceleratorsClient {
	subClient, _ := NewApplicationAcceleratorsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewApplicationLiveViewsClient creates a new instance of ApplicationLiveViewsClient.
func (c *ClientFactory) NewApplicationLiveViewsClient() *ApplicationLiveViewsClient {
	subClient, _ := NewApplicationLiveViewsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewAppsClient creates a new instance of AppsClient.
func (c *ClientFactory) NewAppsClient() *AppsClient {
	subClient, _ := NewAppsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewBindingsClient creates a new instance of BindingsClient.
func (c *ClientFactory) NewBindingsClient() *BindingsClient {
	subClient, _ := NewBindingsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewBuildServiceAgentPoolClient creates a new instance of BuildServiceAgentPoolClient.
func (c *ClientFactory) NewBuildServiceAgentPoolClient() *BuildServiceAgentPoolClient {
	subClient, _ := NewBuildServiceAgentPoolClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewBuildServiceBuilderClient creates a new instance of BuildServiceBuilderClient.
func (c *ClientFactory) NewBuildServiceBuilderClient() *BuildServiceBuilderClient {
	subClient, _ := NewBuildServiceBuilderClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewBuildServiceClient creates a new instance of BuildServiceClient.
func (c *ClientFactory) NewBuildServiceClient() *BuildServiceClient {
	subClient, _ := NewBuildServiceClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewBuildpackBindingClient creates a new instance of BuildpackBindingClient.
func (c *ClientFactory) NewBuildpackBindingClient() *BuildpackBindingClient {
	subClient, _ := NewBuildpackBindingClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewCertificatesClient creates a new instance of CertificatesClient.
func (c *ClientFactory) NewCertificatesClient() *CertificatesClient {
	subClient, _ := NewCertificatesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewConfigServersClient creates a new instance of ConfigServersClient.
func (c *ClientFactory) NewConfigServersClient() *ConfigServersClient {
	subClient, _ := NewConfigServersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewConfigurationServicesClient creates a new instance of ConfigurationServicesClient.
func (c *ClientFactory) NewConfigurationServicesClient() *ConfigurationServicesClient {
	subClient, _ := NewConfigurationServicesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewCustomDomainsClient creates a new instance of CustomDomainsClient.
func (c *ClientFactory) NewCustomDomainsClient() *CustomDomainsClient {
	subClient, _ := NewCustomDomainsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewCustomizedAcceleratorsClient creates a new instance of CustomizedAcceleratorsClient.
func (c *ClientFactory) NewCustomizedAcceleratorsClient() *CustomizedAcceleratorsClient {
	subClient, _ := NewCustomizedAcceleratorsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewDeploymentsClient creates a new instance of DeploymentsClient.
func (c *ClientFactory) NewDeploymentsClient() *DeploymentsClient {
	subClient, _ := NewDeploymentsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewDevToolPortalsClient creates a new instance of DevToolPortalsClient.
func (c *ClientFactory) NewDevToolPortalsClient() *DevToolPortalsClient {
	subClient, _ := NewDevToolPortalsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewGatewayCustomDomainsClient creates a new instance of GatewayCustomDomainsClient.
func (c *ClientFactory) NewGatewayCustomDomainsClient() *GatewayCustomDomainsClient {
	subClient, _ := NewGatewayCustomDomainsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewGatewayRouteConfigsClient creates a new instance of GatewayRouteConfigsClient.
func (c *ClientFactory) NewGatewayRouteConfigsClient() *GatewayRouteConfigsClient {
	subClient, _ := NewGatewayRouteConfigsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewGatewaysClient creates a new instance of GatewaysClient.
func (c *ClientFactory) NewGatewaysClient() *GatewaysClient {
	subClient, _ := NewGatewaysClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewMonitoringSettingsClient creates a new instance of MonitoringSettingsClient.
func (c *ClientFactory) NewMonitoringSettingsClient() *MonitoringSettingsClient {
	subClient, _ := NewMonitoringSettingsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

// NewPredefinedAcceleratorsClient creates a new instance of PredefinedAcceleratorsClient.
func (c *ClientFactory) NewPredefinedAcceleratorsClient() *PredefinedAcceleratorsClient {
	subClient, _ := NewPredefinedAcceleratorsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewRuntimeVersionsClient creates a new instance of RuntimeVersionsClient.
func (c *ClientFactory) NewRuntimeVersionsClient() *RuntimeVersionsClient {
	subClient, _ := NewRuntimeVersionsClient(c.credential, c.options)
	return subClient
}

// NewSKUsClient creates a new instance of SKUsClient.
func (c *ClientFactory) NewSKUsClient() *SKUsClient {
	subClient, _ := NewSKUsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewServiceRegistriesClient creates a new instance of ServiceRegistriesClient.
func (c *ClientFactory) NewServiceRegistriesClient() *ServiceRegistriesClient {
	subClient, _ := NewServiceRegistriesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewServicesClient creates a new instance of ServicesClient.
func (c *ClientFactory) NewServicesClient() *ServicesClient {
	subClient, _ := NewServicesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewStoragesClient creates a new instance of StoragesClient.
func (c *ClientFactory) NewStoragesClient() *StoragesClient {
	subClient, _ := NewStoragesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
