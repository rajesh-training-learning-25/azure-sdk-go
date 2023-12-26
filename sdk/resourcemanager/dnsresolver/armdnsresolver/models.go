//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdnsresolver

import "time"

// DNSForwardingRuleset - Describes a DNS forwarding ruleset.
type DNSForwardingRuleset struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; Properties of the DNS forwarding ruleset.
	Properties *DNSForwardingRulesetProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; ETag of the DNS forwarding ruleset.
	Etag *string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// DNSForwardingRulesetListResult - The response to an enumeration operation on DNS forwarding rulesets.
type DNSForwardingRulesetListResult struct {
	// Enumeration of the DNS forwarding rulesets.
	Value []*DNSForwardingRuleset

	// READ-ONLY; The continuation token for the next page of results.
	NextLink *string
}

// DNSForwardingRulesetPatch - Describes a DNS forwarding ruleset PATCH operation.
type DNSForwardingRulesetPatch struct {
	// The reference to the DNS resolver outbound endpoints that are used to route DNS queries matching the forwarding rules in
	// the ruleset to the target DNS servers.
	DNSResolverOutboundEndpoints []*SubResource

	// Tags for DNS Resolver.
	Tags map[string]*string
}

// DNSForwardingRulesetProperties - Represents the properties of a DNS forwarding ruleset.
type DNSForwardingRulesetProperties struct {
	// REQUIRED; The reference to the DNS resolver outbound endpoints that are used to route DNS queries matching the forwarding
	// rules in the ruleset to the target DNS servers.
	DNSResolverOutboundEndpoints []*SubResource

	// READ-ONLY; The current provisioning state of the DNS forwarding ruleset. This is a read-only property and any attempt to
	// set this value will be ignored.
	ProvisioningState *ProvisioningState

	// READ-ONLY; The resourceGuid for the DNS forwarding ruleset.
	ResourceGUID *string
}

// DNSResolver - Describes a DNS resolver.
type DNSResolver struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; Properties of the DNS resolver.
	Properties *Properties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; ETag of the DNS resolver.
	Etag *string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ForwardingRule - Describes a forwarding rule within a DNS forwarding ruleset.
type ForwardingRule struct {
	// REQUIRED; Properties of the forwarding rule.
	Properties *ForwardingRuleProperties

	// READ-ONLY; ETag of the forwarding rule.
	Etag *string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ForwardingRuleListResult - The response to an enumeration operation on forwarding rules within a DNS forwarding ruleset.
type ForwardingRuleListResult struct {
	// Enumeration of the forwarding rules.
	Value []*ForwardingRule

	// READ-ONLY; The continuation token for the next page of results.
	NextLink *string
}

// ForwardingRulePatch - Describes a forwarding rule for PATCH operation.
type ForwardingRulePatch struct {
	// Updatable properties of the forwarding rule.
	Properties *ForwardingRulePatchProperties
}

// ForwardingRulePatchProperties - Represents the updatable properties of a forwarding rule within a DNS forwarding ruleset.
type ForwardingRulePatchProperties struct {
	// The state of forwarding rule.
	ForwardingRuleState *ForwardingRuleState

	// Metadata attached to the forwarding rule.
	Metadata map[string]*string

	// DNS servers to forward the DNS query to.
	TargetDNSServers []*TargetDNSServer
}

// ForwardingRuleProperties - Represents the properties of a forwarding rule within a DNS forwarding ruleset.
type ForwardingRuleProperties struct {
	// REQUIRED; The domain name for the forwarding rule.
	DomainName *string

	// REQUIRED; DNS servers to forward the DNS query to.
	TargetDNSServers []*TargetDNSServer

	// The state of forwarding rule.
	ForwardingRuleState *ForwardingRuleState

	// Metadata attached to the forwarding rule.
	Metadata map[string]*string

	// READ-ONLY; The current provisioning state of the forwarding rule. This is a read-only property and any attempt to set this
	// value will be ignored.
	ProvisioningState *ProvisioningState
}

// IPConfiguration - IP configuration.
type IPConfiguration struct {
	// REQUIRED; The reference to the subnet bound to the IP configuration.
	Subnet *SubResource

	// Private IP address of the IP configuration.
	PrivateIPAddress *string

	// Private IP address allocation method.
	PrivateIPAllocationMethod *IPAllocationMethod
}

// InboundEndpoint - Describes an inbound endpoint for a DNS resolver.
type InboundEndpoint struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; Properties of the inbound endpoint.
	Properties *InboundEndpointProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; ETag of the inbound endpoint.
	Etag *string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// InboundEndpointListResult - The response to an enumeration operation on inbound endpoints for a DNS resolver.
type InboundEndpointListResult struct {
	// Enumeration of the inbound endpoints for a DNS resolver.
	Value []*InboundEndpoint

	// READ-ONLY; The continuation token for the next page of results.
	NextLink *string
}

// InboundEndpointPatch - Describes an inbound endpoint for a DNS resolver for PATCH operation.
type InboundEndpointPatch struct {
	// Tags for inbound endpoint.
	Tags map[string]*string
}

// InboundEndpointProperties - Represents the properties of an inbound endpoint for a DNS resolver.
type InboundEndpointProperties struct {
	// REQUIRED; IP configurations for the inbound endpoint.
	IPConfigurations []*IPConfiguration

	// READ-ONLY; The current provisioning state of the inbound endpoint. This is a read-only property and any attempt to set
	// this value will be ignored.
	ProvisioningState *ProvisioningState

	// READ-ONLY; The resourceGuid property of the inbound endpoint resource.
	ResourceGUID *string
}

// ListResult - The response to an enumeration operation on DNS resolvers.
type ListResult struct {
	// Enumeration of the DNS resolvers.
	Value []*DNSResolver

	// READ-ONLY; The continuation token for the next page of results.
	NextLink *string
}

// OutboundEndpoint - Describes an outbound endpoint for a DNS resolver.
type OutboundEndpoint struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; Properties of the outbound endpoint.
	Properties *OutboundEndpointProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; ETag of the outbound endpoint.
	Etag *string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// OutboundEndpointListResult - The response to an enumeration operation on outbound endpoints for a DNS resolver.
type OutboundEndpointListResult struct {
	// Enumeration of the outbound endpoints for a DNS resolver.
	Value []*OutboundEndpoint

	// READ-ONLY; The continuation token for the next page of results.
	NextLink *string
}

// OutboundEndpointPatch - Describes an outbound endpoint for a DNS resolver for PATCH operation.
type OutboundEndpointPatch struct {
	// Tags for outbound endpoint.
	Tags map[string]*string
}

// OutboundEndpointProperties - Represents the properties of an outbound endpoint for a DNS resolver.
type OutboundEndpointProperties struct {
	// REQUIRED; The reference to the subnet used for the outbound endpoint.
	Subnet *SubResource

	// READ-ONLY; The current provisioning state of the outbound endpoint. This is a read-only property and any attempt to set
	// this value will be ignored.
	ProvisioningState *ProvisioningState

	// READ-ONLY; The resourceGuid property of the outbound endpoint resource.
	ResourceGUID *string
}

// Patch - Describes a DNS resolver for PATCH operation.
type Patch struct {
	// Tags for DNS Resolver.
	Tags map[string]*string
}

// Properties - Represents the properties of a DNS resolver.
type Properties struct {
	// REQUIRED; The reference to the virtual network. This cannot be changed after creation.
	VirtualNetwork *SubResource

	// READ-ONLY; The current status of the DNS resolver. This is a read-only property and any attempt to set this value will
	// be ignored.
	DNSResolverState *DNSResolverState

	// READ-ONLY; The current provisioning state of the DNS resolver. This is a read-only property and any attempt to set this
	// value will be ignored.
	ProvisioningState *ProvisioningState

	// READ-ONLY; The resourceGuid property of the DNS resolver resource.
	ResourceGUID *string
}

// SubResource - Reference to another ARM resource.
type SubResource struct {
	// REQUIRED; Resource ID.
	ID *string
}

// SubResourceListResult - The response to an enumeration operation on sub-resources.
type SubResourceListResult struct {
	// Enumeration of the sub-resources.
	Value []*SubResource

	// READ-ONLY; The continuation token for the next page of results.
	NextLink *string
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}

// TargetDNSServer - Describes a server to forward the DNS queries to.
type TargetDNSServer struct {
	// REQUIRED; DNS server IP address.
	IPAddress *string

	// DNS server port.
	Port *int32
}

// VirtualNetworkDNSForwardingRuleset - Reference to DNS forwarding ruleset and associated virtual network link.
type VirtualNetworkDNSForwardingRuleset struct {
	// DNS Forwarding Ruleset Resource ID.
	ID *string

	// Properties of the virtual network link sub-resource reference.
	Properties *VirtualNetworkLinkSubResourceProperties
}

// VirtualNetworkDNSForwardingRulesetListResult - The response to an enumeration operation on Virtual Network DNS Forwarding
// Ruleset.
type VirtualNetworkDNSForwardingRulesetListResult struct {
	// Enumeration of the Virtual Network DNS Forwarding Ruleset.
	Value []*VirtualNetworkDNSForwardingRuleset

	// READ-ONLY; The continuation token for the next page of results.
	NextLink *string
}

// VirtualNetworkLink - Describes a virtual network link.
type VirtualNetworkLink struct {
	// REQUIRED; Properties of the virtual network link.
	Properties *VirtualNetworkLinkProperties

	// READ-ONLY; ETag of the virtual network link.
	Etag *string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// VirtualNetworkLinkListResult - The response to an enumeration operation on virtual network links.
type VirtualNetworkLinkListResult struct {
	// Enumeration of the virtual network links.
	Value []*VirtualNetworkLink

	// READ-ONLY; The continuation token for the next page of results.
	NextLink *string
}

// VirtualNetworkLinkPatch - Describes a virtual network link for PATCH operation.
type VirtualNetworkLinkPatch struct {
	// Updatable properties of the virtual network link.
	Properties *VirtualNetworkLinkPatchProperties
}

// VirtualNetworkLinkPatchProperties - Represents the updatable properties of the virtual network link.
type VirtualNetworkLinkPatchProperties struct {
	// Metadata attached to the virtual network link.
	Metadata map[string]*string
}

// VirtualNetworkLinkProperties - Represents the properties of a virtual network link.
type VirtualNetworkLinkProperties struct {
	// REQUIRED; The reference to the virtual network. This cannot be changed after creation.
	VirtualNetwork *SubResource

	// Metadata attached to the virtual network link.
	Metadata map[string]*string

	// READ-ONLY; The current provisioning state of the virtual network link. This is a read-only property and any attempt to
	// set this value will be ignored.
	ProvisioningState *ProvisioningState
}

// VirtualNetworkLinkSubResourceProperties - The reference to the virtual network link that associates between the DNS forwarding
// ruleset and virtual network.
type VirtualNetworkLinkSubResourceProperties struct {
	// The reference to the virtual network link.
	VirtualNetworkLink *SubResource
}
