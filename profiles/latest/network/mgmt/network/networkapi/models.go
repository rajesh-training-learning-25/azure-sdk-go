// +build go1.9

// Copyright 2019 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package networkapi

import original "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-09-01/network/networkapi"

type ApplicationGatewaysClientAPI = original.ApplicationGatewaysClientAPI
type ApplicationSecurityGroupsClientAPI = original.ApplicationSecurityGroupsClientAPI
type AvailableDelegationsClientAPI = original.AvailableDelegationsClientAPI
type AvailableEndpointServicesClientAPI = original.AvailableEndpointServicesClientAPI
type AvailablePrivateEndpointTypesClientAPI = original.AvailablePrivateEndpointTypesClientAPI
type AvailableResourceGroupDelegationsClientAPI = original.AvailableResourceGroupDelegationsClientAPI
type AvailableServiceAliasesClientAPI = original.AvailableServiceAliasesClientAPI
type AzureFirewallFqdnTagsClientAPI = original.AzureFirewallFqdnTagsClientAPI
type AzureFirewallsClientAPI = original.AzureFirewallsClientAPI
type BaseClientAPI = original.BaseClientAPI
type BastionHostsClientAPI = original.BastionHostsClientAPI
type BgpServiceCommunitiesClientAPI = original.BgpServiceCommunitiesClientAPI
type ConnectionMonitorsClientAPI = original.ConnectionMonitorsClientAPI
type DdosCustomPoliciesClientAPI = original.DdosCustomPoliciesClientAPI
type DdosProtectionPlansClientAPI = original.DdosProtectionPlansClientAPI
type DefaultSecurityRulesClientAPI = original.DefaultSecurityRulesClientAPI
type ExpressRouteCircuitAuthorizationsClientAPI = original.ExpressRouteCircuitAuthorizationsClientAPI
type ExpressRouteCircuitConnectionsClientAPI = original.ExpressRouteCircuitConnectionsClientAPI
type ExpressRouteCircuitPeeringsClientAPI = original.ExpressRouteCircuitPeeringsClientAPI
type ExpressRouteCircuitsClientAPI = original.ExpressRouteCircuitsClientAPI
type ExpressRouteConnectionsClientAPI = original.ExpressRouteConnectionsClientAPI
type ExpressRouteCrossConnectionPeeringsClientAPI = original.ExpressRouteCrossConnectionPeeringsClientAPI
type ExpressRouteCrossConnectionsClientAPI = original.ExpressRouteCrossConnectionsClientAPI
type ExpressRouteGatewaysClientAPI = original.ExpressRouteGatewaysClientAPI
type ExpressRouteLinksClientAPI = original.ExpressRouteLinksClientAPI
type ExpressRoutePortsClientAPI = original.ExpressRoutePortsClientAPI
type ExpressRoutePortsLocationsClientAPI = original.ExpressRoutePortsLocationsClientAPI
type ExpressRouteServiceProvidersClientAPI = original.ExpressRouteServiceProvidersClientAPI
type FirewallPoliciesClientAPI = original.FirewallPoliciesClientAPI
type FirewallPolicyRuleGroupsClientAPI = original.FirewallPolicyRuleGroupsClientAPI
type HubVirtualNetworkConnectionsClientAPI = original.HubVirtualNetworkConnectionsClientAPI
type IPGroupsClientAPI = original.IPGroupsClientAPI
type InboundNatRulesClientAPI = original.InboundNatRulesClientAPI
type InterfaceIPConfigurationsClientAPI = original.InterfaceIPConfigurationsClientAPI
type InterfaceLoadBalancersClientAPI = original.InterfaceLoadBalancersClientAPI
type InterfaceTapConfigurationsClientAPI = original.InterfaceTapConfigurationsClientAPI
type InterfacesClientAPI = original.InterfacesClientAPI
type LoadBalancerBackendAddressPoolsClientAPI = original.LoadBalancerBackendAddressPoolsClientAPI
type LoadBalancerFrontendIPConfigurationsClientAPI = original.LoadBalancerFrontendIPConfigurationsClientAPI
type LoadBalancerLoadBalancingRulesClientAPI = original.LoadBalancerLoadBalancingRulesClientAPI
type LoadBalancerNetworkInterfacesClientAPI = original.LoadBalancerNetworkInterfacesClientAPI
type LoadBalancerOutboundRulesClientAPI = original.LoadBalancerOutboundRulesClientAPI
type LoadBalancerProbesClientAPI = original.LoadBalancerProbesClientAPI
type LoadBalancersClientAPI = original.LoadBalancersClientAPI
type LocalNetworkGatewaysClientAPI = original.LocalNetworkGatewaysClientAPI
type NatGatewaysClientAPI = original.NatGatewaysClientAPI
type OperationsClientAPI = original.OperationsClientAPI
type P2sVpnGatewaysClientAPI = original.P2sVpnGatewaysClientAPI
type PacketCapturesClientAPI = original.PacketCapturesClientAPI
type PeerExpressRouteCircuitConnectionsClientAPI = original.PeerExpressRouteCircuitConnectionsClientAPI
type PrivateEndpointsClientAPI = original.PrivateEndpointsClientAPI
type PrivateLinkServicesClientAPI = original.PrivateLinkServicesClientAPI
type ProfilesClientAPI = original.ProfilesClientAPI
type PublicIPAddressesClientAPI = original.PublicIPAddressesClientAPI
type PublicIPPrefixesClientAPI = original.PublicIPPrefixesClientAPI
type ResourceNavigationLinksClientAPI = original.ResourceNavigationLinksClientAPI
type RouteFilterRulesClientAPI = original.RouteFilterRulesClientAPI
type RouteFiltersClientAPI = original.RouteFiltersClientAPI
type RouteTablesClientAPI = original.RouteTablesClientAPI
type RoutesClientAPI = original.RoutesClientAPI
type SecurityGroupsClientAPI = original.SecurityGroupsClientAPI
type SecurityRulesClientAPI = original.SecurityRulesClientAPI
type ServiceAssociationLinksClientAPI = original.ServiceAssociationLinksClientAPI
type ServiceEndpointPoliciesClientAPI = original.ServiceEndpointPoliciesClientAPI
type ServiceEndpointPolicyDefinitionsClientAPI = original.ServiceEndpointPolicyDefinitionsClientAPI
type ServiceTagsClientAPI = original.ServiceTagsClientAPI
type SubnetsClientAPI = original.SubnetsClientAPI
type UsagesClientAPI = original.UsagesClientAPI
type VirtualHubRouteTableV2sClientAPI = original.VirtualHubRouteTableV2sClientAPI
type VirtualHubsClientAPI = original.VirtualHubsClientAPI
type VirtualNetworkGatewayConnectionsClientAPI = original.VirtualNetworkGatewayConnectionsClientAPI
type VirtualNetworkGatewaysClientAPI = original.VirtualNetworkGatewaysClientAPI
type VirtualNetworkPeeringsClientAPI = original.VirtualNetworkPeeringsClientAPI
type VirtualNetworkTapsClientAPI = original.VirtualNetworkTapsClientAPI
type VirtualNetworksClientAPI = original.VirtualNetworksClientAPI
type VirtualRouterPeeringsClientAPI = original.VirtualRouterPeeringsClientAPI
type VirtualRoutersClientAPI = original.VirtualRoutersClientAPI
type VirtualWansClientAPI = original.VirtualWansClientAPI
type VpnConnectionsClientAPI = original.VpnConnectionsClientAPI
type VpnGatewaysClientAPI = original.VpnGatewaysClientAPI
type VpnLinkConnectionsClientAPI = original.VpnLinkConnectionsClientAPI
type VpnServerConfigurationsAssociatedWithVirtualWanClientAPI = original.VpnServerConfigurationsAssociatedWithVirtualWanClientAPI
type VpnServerConfigurationsClientAPI = original.VpnServerConfigurationsClientAPI
type VpnSiteLinkConnectionsClientAPI = original.VpnSiteLinkConnectionsClientAPI
type VpnSiteLinksClientAPI = original.VpnSiteLinksClientAPI
type VpnSitesClientAPI = original.VpnSitesClientAPI
type VpnSitesConfigurationClientAPI = original.VpnSitesConfigurationClientAPI
type WatchersClientAPI = original.WatchersClientAPI
type WebApplicationFirewallPoliciesClientAPI = original.WebApplicationFirewallPoliciesClientAPI
