// +build go1.9

// Copyright 2018 Microsoft Corporation
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

package fabric

import original "github.com/Azure/azure-sdk-for-go/services/azsadmin/mgmt/2016-05-01/fabric"

type SlbMuxInstancesClient = original.SlbMuxInstancesClient
type MacAddressPoolsClient = original.MacAddressPoolsClient
type StoragePoolsClient = original.StoragePoolsClient
type ComputeFabricOperationsClient = original.ComputeFabricOperationsClient
type LogicalNetworksClient = original.LogicalNetworksClient
type ScaleUnitNodesClient = original.ScaleUnitNodesClient
type IPPoolsClient = original.IPPoolsClient
type InfraRoleInstancesClient = original.InfraRoleInstancesClient
type VolumesClient = original.VolumesClient
type NetworkFabricOperationsClient = original.NetworkFabricOperationsClient
type LogicalSubnetsClient = original.LogicalSubnetsClient
type EdgeGatewayPoolsClient = original.EdgeGatewayPoolsClient
type EdgeGatewaysClient = original.EdgeGatewaysClient
type InfraRoleInstanceState = original.InfraRoleInstanceState

const (
	Running  InfraRoleInstanceState = original.Running
	Starting InfraRoleInstanceState = original.Starting
	Stopped  InfraRoleInstanceState = original.Stopped
	Stopping InfraRoleInstanceState = original.Stopping
)

type PowerState = original.PowerState

const (
	PowerStateRunning  PowerState = original.PowerStateRunning
	PowerStateStarting PowerState = original.PowerStateStarting
	PowerStateStopped  PowerState = original.PowerStateStopped
	PowerStateStopping PowerState = original.PowerStateStopping
)

type ScaleUnitNodeStatus = original.ScaleUnitNodeStatus

const (
	ScaleUnitNodeStatusMaintenance ScaleUnitNodeStatus = original.ScaleUnitNodeStatusMaintenance
	ScaleUnitNodeStatusRunning     ScaleUnitNodeStatus = original.ScaleUnitNodeStatusRunning
	ScaleUnitNodeStatusStopped     ScaleUnitNodeStatus = original.ScaleUnitNodeStatusStopped
)

type ScaleUnitState = original.ScaleUnitState

const (
	ScaleUnitStateCreating  ScaleUnitState = original.ScaleUnitStateCreating
	ScaleUnitStateDeleting  ScaleUnitState = original.ScaleUnitStateDeleting
	ScaleUnitStateRunning   ScaleUnitState = original.ScaleUnitStateRunning
	ScaleUnitStateUnknown   ScaleUnitState = original.ScaleUnitStateUnknown
	ScaleUnitStateUpgrading ScaleUnitState = original.ScaleUnitStateUpgrading
)

type ScaleUnitType = original.ScaleUnitType

const (
	ComputeOnly    ScaleUnitType = original.ComputeOnly
	HyperConverged ScaleUnitType = original.HyperConverged
	StorageOnly    ScaleUnitType = original.StorageOnly
	Unknown        ScaleUnitType = original.Unknown
)

type EdgeGateway = original.EdgeGateway
type EdgeGatewayList = original.EdgeGatewayList
type EdgeGatewayListIterator = original.EdgeGatewayListIterator
type EdgeGatewayListPage = original.EdgeGatewayListPage
type EdgeGatewayModel = original.EdgeGatewayModel
type EdgeGatewayPool = original.EdgeGatewayPool
type EdgeGatewayPoolList = original.EdgeGatewayPoolList
type EdgeGatewayPoolListIterator = original.EdgeGatewayPoolListIterator
type EdgeGatewayPoolListPage = original.EdgeGatewayPoolListPage
type EdgeGatewayPoolModel = original.EdgeGatewayPoolModel
type FileShare = original.FileShare
type FileShareList = original.FileShareList
type FileShareModel = original.FileShareModel
type InfraRole = original.InfraRole
type InfraRoleInstance = original.InfraRoleInstance
type InfraRoleInstanceList = original.InfraRoleInstanceList
type InfraRoleInstanceListIterator = original.InfraRoleInstanceListIterator
type InfraRoleInstanceListPage = original.InfraRoleInstanceListPage
type InfraRoleInstanceModel = original.InfraRoleInstanceModel
type InfraRoleInstanceSize = original.InfraRoleInstanceSize
type InfraRoleInstancesPowerOffFuture = original.InfraRoleInstancesPowerOffFuture
type InfraRoleInstancesPowerOnFuture = original.InfraRoleInstancesPowerOnFuture
type InfraRoleInstancesRebootFuture = original.InfraRoleInstancesRebootFuture
type InfraRoleInstancesShutdownFuture = original.InfraRoleInstancesShutdownFuture
type InfraRoleList = original.InfraRoleList
type InfraRoleListIterator = original.InfraRoleListIterator
type InfraRoleListPage = original.InfraRoleListPage
type InfraRoleModel = original.InfraRoleModel
type IPPool = original.IPPool
type IPPoolList = original.IPPoolList
type IPPoolListIterator = original.IPPoolListIterator
type IPPoolListPage = original.IPPoolListPage
type IPPoolModel = original.IPPoolModel
type IPPoolsCreateFuture = original.IPPoolsCreateFuture
type Location = original.Location
type LocationList = original.LocationList
type LocationListIterator = original.LocationListIterator
type LocationListPage = original.LocationListPage
type LogicalNetwork = original.LogicalNetwork
type LogicalNetworkList = original.LogicalNetworkList
type LogicalNetworkListIterator = original.LogicalNetworkListIterator
type LogicalNetworkListPage = original.LogicalNetworkListPage
type LogicalNetworkModel = original.LogicalNetworkModel
type LogicalSubnet = original.LogicalSubnet
type LogicalSubnetList = original.LogicalSubnetList
type LogicalSubnetListIterator = original.LogicalSubnetListIterator
type LogicalSubnetListPage = original.LogicalSubnetListPage
type LogicalSubnetModel = original.LogicalSubnetModel
type MacAddressPool = original.MacAddressPool
type MacAddressPoolList = original.MacAddressPoolList
type MacAddressPoolListIterator = original.MacAddressPoolListIterator
type MacAddressPoolListPage = original.MacAddressPoolListPage
type MacAddressPoolModel = original.MacAddressPoolModel
type OperationStatus = original.OperationStatus
type OperationStatusLocation = original.OperationStatusLocation
type ProvisioningStateModel = original.ProvisioningStateModel
type Resource = original.Resource
type ScaleUnit = original.ScaleUnit
type ScaleUnitCapacity = original.ScaleUnitCapacity
type ScaleUnitList = original.ScaleUnitList
type ScaleUnitListIterator = original.ScaleUnitListIterator
type ScaleUnitListPage = original.ScaleUnitListPage
type ScaleUnitModel = original.ScaleUnitModel
type ScaleUnitNode = original.ScaleUnitNode
type ScaleUnitNodeList = original.ScaleUnitNodeList
type ScaleUnitNodeListIterator = original.ScaleUnitNodeListIterator
type ScaleUnitNodeListPage = original.ScaleUnitNodeListPage
type ScaleUnitNodeModel = original.ScaleUnitNodeModel
type ScaleUnitNodesPowerOffFuture = original.ScaleUnitNodesPowerOffFuture
type ScaleUnitNodesPowerOnFuture = original.ScaleUnitNodesPowerOnFuture
type ScaleUnitNodesStartMaintenanceModeFuture = original.ScaleUnitNodesStartMaintenanceModeFuture
type ScaleUnitNodesStopMaintenanceModeFuture = original.ScaleUnitNodesStopMaintenanceModeFuture
type SlbMuxInstance = original.SlbMuxInstance
type SlbMuxInstanceList = original.SlbMuxInstanceList
type SlbMuxInstanceListIterator = original.SlbMuxInstanceListIterator
type SlbMuxInstanceListPage = original.SlbMuxInstanceListPage
type SlbMuxInstanceModel = original.SlbMuxInstanceModel
type StoragePool = original.StoragePool
type StoragePoolList = original.StoragePoolList
type StoragePoolListIterator = original.StoragePoolListIterator
type StoragePoolListPage = original.StoragePoolListPage
type StoragePoolModel = original.StoragePoolModel
type StorageSystem = original.StorageSystem
type StorageSystemList = original.StorageSystemList
type StorageSystemListIterator = original.StorageSystemListIterator
type StorageSystemListPage = original.StorageSystemListPage
type StorageSystemModel = original.StorageSystemModel
type Volume = original.Volume
type VolumeList = original.VolumeList
type VolumeListIterator = original.VolumeListIterator
type VolumeListPage = original.VolumeListPage
type VolumeModel = original.VolumeModel
type ScaleUnitsClient = original.ScaleUnitsClient
type InfraRolesClient = original.InfraRolesClient
type LocationsClient = original.LocationsClient
type StorageSystemsClient = original.StorageSystemsClient
type FileSharesClient = original.FileSharesClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient

func NewIPPoolsClient(subscriptionID string) IPPoolsClient {
	return original.NewIPPoolsClient(subscriptionID)
}
func NewIPPoolsClientWithBaseURI(baseURI string, subscriptionID string) IPPoolsClient {
	return original.NewIPPoolsClientWithBaseURI(baseURI, subscriptionID)
}
func NewVolumesClient(subscriptionID string) VolumesClient {
	return original.NewVolumesClient(subscriptionID)
}
func NewVolumesClientWithBaseURI(baseURI string, subscriptionID string) VolumesClient {
	return original.NewVolumesClientWithBaseURI(baseURI, subscriptionID)
}
func NewNetworkFabricOperationsClient(subscriptionID string) NetworkFabricOperationsClient {
	return original.NewNetworkFabricOperationsClient(subscriptionID)
}
func NewNetworkFabricOperationsClientWithBaseURI(baseURI string, subscriptionID string) NetworkFabricOperationsClient {
	return original.NewNetworkFabricOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewLogicalSubnetsClient(subscriptionID string) LogicalSubnetsClient {
	return original.NewLogicalSubnetsClient(subscriptionID)
}
func NewLogicalSubnetsClientWithBaseURI(baseURI string, subscriptionID string) LogicalSubnetsClient {
	return original.NewLogicalSubnetsClientWithBaseURI(baseURI, subscriptionID)
}
func NewEdgeGatewayPoolsClient(subscriptionID string) EdgeGatewayPoolsClient {
	return original.NewEdgeGatewayPoolsClient(subscriptionID)
}
func NewEdgeGatewayPoolsClientWithBaseURI(baseURI string, subscriptionID string) EdgeGatewayPoolsClient {
	return original.NewEdgeGatewayPoolsClientWithBaseURI(baseURI, subscriptionID)
}
func NewInfraRoleInstancesClient(subscriptionID string) InfraRoleInstancesClient {
	return original.NewInfraRoleInstancesClient(subscriptionID)
}
func NewInfraRoleInstancesClientWithBaseURI(baseURI string, subscriptionID string) InfraRoleInstancesClient {
	return original.NewInfraRoleInstancesClientWithBaseURI(baseURI, subscriptionID)
}
func NewScaleUnitsClient(subscriptionID string) ScaleUnitsClient {
	return original.NewScaleUnitsClient(subscriptionID)
}
func NewScaleUnitsClientWithBaseURI(baseURI string, subscriptionID string) ScaleUnitsClient {
	return original.NewScaleUnitsClientWithBaseURI(baseURI, subscriptionID)
}
func NewInfraRolesClient(subscriptionID string) InfraRolesClient {
	return original.NewInfraRolesClient(subscriptionID)
}
func NewInfraRolesClientWithBaseURI(baseURI string, subscriptionID string) InfraRolesClient {
	return original.NewInfraRolesClientWithBaseURI(baseURI, subscriptionID)
}
func NewLocationsClient(subscriptionID string) LocationsClient {
	return original.NewLocationsClient(subscriptionID)
}
func NewLocationsClientWithBaseURI(baseURI string, subscriptionID string) LocationsClient {
	return original.NewLocationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewStorageSystemsClient(subscriptionID string) StorageSystemsClient {
	return original.NewStorageSystemsClient(subscriptionID)
}
func NewStorageSystemsClientWithBaseURI(baseURI string, subscriptionID string) StorageSystemsClient {
	return original.NewStorageSystemsClientWithBaseURI(baseURI, subscriptionID)
}
func NewFileSharesClient(subscriptionID string) FileSharesClient {
	return original.NewFileSharesClient(subscriptionID)
}
func NewFileSharesClientWithBaseURI(baseURI string, subscriptionID string) FileSharesClient {
	return original.NewFileSharesClientWithBaseURI(baseURI, subscriptionID)
}
func NewEdgeGatewaysClient(subscriptionID string) EdgeGatewaysClient {
	return original.NewEdgeGatewaysClient(subscriptionID)
}
func NewEdgeGatewaysClientWithBaseURI(baseURI string, subscriptionID string) EdgeGatewaysClient {
	return original.NewEdgeGatewaysClientWithBaseURI(baseURI, subscriptionID)
}
func PossibleInfraRoleInstanceStateValues() []InfraRoleInstanceState {
	return original.PossibleInfraRoleInstanceStateValues()
}
func PossiblePowerStateValues() []PowerState {
	return original.PossiblePowerStateValues()
}
func PossibleScaleUnitNodeStatusValues() []ScaleUnitNodeStatus {
	return original.PossibleScaleUnitNodeStatusValues()
}
func PossibleScaleUnitStateValues() []ScaleUnitState {
	return original.PossibleScaleUnitStateValues()
}
func PossibleScaleUnitTypeValues() []ScaleUnitType {
	return original.PossibleScaleUnitTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewStoragePoolsClient(subscriptionID string) StoragePoolsClient {
	return original.NewStoragePoolsClient(subscriptionID)
}
func NewStoragePoolsClientWithBaseURI(baseURI string, subscriptionID string) StoragePoolsClient {
	return original.NewStoragePoolsClientWithBaseURI(baseURI, subscriptionID)
}
func NewComputeFabricOperationsClient(subscriptionID string) ComputeFabricOperationsClient {
	return original.NewComputeFabricOperationsClient(subscriptionID)
}
func NewComputeFabricOperationsClientWithBaseURI(baseURI string, subscriptionID string) ComputeFabricOperationsClient {
	return original.NewComputeFabricOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewLogicalNetworksClient(subscriptionID string) LogicalNetworksClient {
	return original.NewLogicalNetworksClient(subscriptionID)
}
func NewLogicalNetworksClientWithBaseURI(baseURI string, subscriptionID string) LogicalNetworksClient {
	return original.NewLogicalNetworksClientWithBaseURI(baseURI, subscriptionID)
}
func NewScaleUnitNodesClient(subscriptionID string) ScaleUnitNodesClient {
	return original.NewScaleUnitNodesClient(subscriptionID)
}
func NewScaleUnitNodesClientWithBaseURI(baseURI string, subscriptionID string) ScaleUnitNodesClient {
	return original.NewScaleUnitNodesClientWithBaseURI(baseURI, subscriptionID)
}
func NewSlbMuxInstancesClient(subscriptionID string) SlbMuxInstancesClient {
	return original.NewSlbMuxInstancesClient(subscriptionID)
}
func NewSlbMuxInstancesClientWithBaseURI(baseURI string, subscriptionID string) SlbMuxInstancesClient {
	return original.NewSlbMuxInstancesClientWithBaseURI(baseURI, subscriptionID)
}
func NewMacAddressPoolsClient(subscriptionID string) MacAddressPoolsClient {
	return original.NewMacAddressPoolsClient(subscriptionID)
}
func NewMacAddressPoolsClientWithBaseURI(baseURI string, subscriptionID string) MacAddressPoolsClient {
	return original.NewMacAddressPoolsClientWithBaseURI(baseURI, subscriptionID)
}
