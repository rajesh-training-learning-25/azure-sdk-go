//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.
// Code generated by @autorest/go. DO NOT EDIT.

package armconnectedvmware

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware"
	moduleVersion = "v1.1.1"
)

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// DiskMode - Defines the different types of disk modes.
type DiskMode string

const (
	DiskModeIndependentNonpersistent DiskMode = "independent_nonpersistent"
	DiskModeIndependentPersistent    DiskMode = "independent_persistent"
	DiskModePersistent               DiskMode = "persistent"
)

// PossibleDiskModeValues returns the possible values for the DiskMode const type.
func PossibleDiskModeValues() []DiskMode {
	return []DiskMode{
		DiskModeIndependentNonpersistent,
		DiskModeIndependentPersistent,
		DiskModePersistent,
	}
}

// DiskType - Defines the different types of disks.
type DiskType string

const (
	DiskTypeFlat        DiskType = "flat"
	DiskTypePmem        DiskType = "pmem"
	DiskTypeRawphysical DiskType = "rawphysical"
	DiskTypeRawvirtual  DiskType = "rawvirtual"
	DiskTypeSesparse    DiskType = "sesparse"
	DiskTypeSparse      DiskType = "sparse"
	DiskTypeUnknown     DiskType = "unknown"
)

// PossibleDiskTypeValues returns the possible values for the DiskType const type.
func PossibleDiskTypeValues() []DiskType {
	return []DiskType{
		DiskTypeFlat,
		DiskTypePmem,
		DiskTypeRawphysical,
		DiskTypeRawvirtual,
		DiskTypeSesparse,
		DiskTypeSparse,
		DiskTypeUnknown,
	}
}

// FirmwareType - Firmware type
type FirmwareType string

const (
	FirmwareTypeBios FirmwareType = "bios"
	FirmwareTypeEfi  FirmwareType = "efi"
)

// PossibleFirmwareTypeValues returns the possible values for the FirmwareType const type.
func PossibleFirmwareTypeValues() []FirmwareType {
	return []FirmwareType{
		FirmwareTypeBios,
		FirmwareTypeEfi,
	}
}

// IPAddressAllocationMethod - IP address allocation method.
type IPAddressAllocationMethod string

const (
	IPAddressAllocationMethodDynamic   IPAddressAllocationMethod = "dynamic"
	IPAddressAllocationMethodLinklayer IPAddressAllocationMethod = "linklayer"
	IPAddressAllocationMethodOther     IPAddressAllocationMethod = "other"
	IPAddressAllocationMethodRandom    IPAddressAllocationMethod = "random"
	IPAddressAllocationMethodStatic    IPAddressAllocationMethod = "static"
	IPAddressAllocationMethodUnset     IPAddressAllocationMethod = "unset"
)

// PossibleIPAddressAllocationMethodValues returns the possible values for the IPAddressAllocationMethod const type.
func PossibleIPAddressAllocationMethodValues() []IPAddressAllocationMethod {
	return []IPAddressAllocationMethod{
		IPAddressAllocationMethodDynamic,
		IPAddressAllocationMethodLinklayer,
		IPAddressAllocationMethodOther,
		IPAddressAllocationMethodRandom,
		IPAddressAllocationMethodStatic,
		IPAddressAllocationMethodUnset,
	}
}

// InventoryType - The inventory type.
type InventoryType string

const (
	InventoryTypeCluster                InventoryType = "Cluster"
	InventoryTypeDatastore              InventoryType = "Datastore"
	InventoryTypeHost                   InventoryType = "Host"
	InventoryTypeResourcePool           InventoryType = "ResourcePool"
	InventoryTypeVirtualMachine         InventoryType = "VirtualMachine"
	InventoryTypeVirtualMachineTemplate InventoryType = "VirtualMachineTemplate"
	InventoryTypeVirtualNetwork         InventoryType = "VirtualNetwork"
)

// PossibleInventoryTypeValues returns the possible values for the InventoryType const type.
func PossibleInventoryTypeValues() []InventoryType {
	return []InventoryType{
		InventoryTypeCluster,
		InventoryTypeDatastore,
		InventoryTypeHost,
		InventoryTypeResourcePool,
		InventoryTypeVirtualMachine,
		InventoryTypeVirtualMachineTemplate,
		InventoryTypeVirtualNetwork,
	}
}

// NICType - NIC type
type NICType string

const (
	NICTypeE1000   NICType = "e1000"
	NICTypeE1000E  NICType = "e1000e"
	NICTypePcnet32 NICType = "pcnet32"
	NICTypeVmxnet  NICType = "vmxnet"
	NICTypeVmxnet2 NICType = "vmxnet2"
	NICTypeVmxnet3 NICType = "vmxnet3"
)

// PossibleNICTypeValues returns the possible values for the NICType const type.
func PossibleNICTypeValues() []NICType {
	return []NICType{
		NICTypeE1000,
		NICTypeE1000E,
		NICTypePcnet32,
		NICTypeVmxnet,
		NICTypeVmxnet2,
		NICTypeVmxnet3,
	}
}

// OsType - Defines the different types of VM guest operating systems.
type OsType string

const (
	OsTypeLinux   OsType = "Linux"
	OsTypeOther   OsType = "Other"
	OsTypeWindows OsType = "Windows"
)

// PossibleOsTypeValues returns the possible values for the OsType const type.
func PossibleOsTypeValues() []OsType {
	return []OsType{
		OsTypeLinux,
		OsTypeOther,
		OsTypeWindows,
	}
}

// PowerOnBootOption - Defines the options for power on boot.
type PowerOnBootOption string

const (
	PowerOnBootOptionDisabled PowerOnBootOption = "disabled"
	PowerOnBootOptionEnabled  PowerOnBootOption = "enabled"
)

// PossiblePowerOnBootOptionValues returns the possible values for the PowerOnBootOption const type.
func PossiblePowerOnBootOptionValues() []PowerOnBootOption {
	return []PowerOnBootOption{
		PowerOnBootOptionDisabled,
		PowerOnBootOptionEnabled,
	}
}

// ProvisioningAction - Defines the different types of operations for guest agent.
type ProvisioningAction string

const (
	ProvisioningActionInstall   ProvisioningAction = "install"
	ProvisioningActionRepair    ProvisioningAction = "repair"
	ProvisioningActionUninstall ProvisioningAction = "uninstall"
)

// PossibleProvisioningActionValues returns the possible values for the ProvisioningAction const type.
func PossibleProvisioningActionValues() []ProvisioningAction {
	return []ProvisioningAction{
		ProvisioningActionInstall,
		ProvisioningActionRepair,
		ProvisioningActionUninstall,
	}
}

// ProvisioningState - The current deployment state of resource.
type ProvisioningState string

const (
	ProvisioningStateAccepted     ProvisioningState = "Accepted"
	ProvisioningStateCanceled     ProvisioningState = "Canceled"
	ProvisioningStateCreated      ProvisioningState = "Created"
	ProvisioningStateDeleting     ProvisioningState = "Deleting"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
	ProvisioningStateUpdating     ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateCreated,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateProvisioning,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// SCSIControllerType - Defines the different types of SCSI controllers.
type SCSIControllerType string

const (
	SCSIControllerTypeBuslogic    SCSIControllerType = "buslogic"
	SCSIControllerTypeLsilogic    SCSIControllerType = "lsilogic"
	SCSIControllerTypeLsilogicsas SCSIControllerType = "lsilogicsas"
	SCSIControllerTypePvscsi      SCSIControllerType = "pvscsi"
)

// PossibleSCSIControllerTypeValues returns the possible values for the SCSIControllerType const type.
func PossibleSCSIControllerTypeValues() []SCSIControllerType {
	return []SCSIControllerType{
		SCSIControllerTypeBuslogic,
		SCSIControllerTypeLsilogic,
		SCSIControllerTypeLsilogicsas,
		SCSIControllerTypePvscsi,
	}
}

// VirtualSCSISharing - Defines the sharing mode for sharing the SCSI bus.
type VirtualSCSISharing string

const (
	VirtualSCSISharingNoSharing       VirtualSCSISharing = "noSharing"
	VirtualSCSISharingPhysicalSharing VirtualSCSISharing = "physicalSharing"
	VirtualSCSISharingVirtualSharing  VirtualSCSISharing = "virtualSharing"
)

// PossibleVirtualSCSISharingValues returns the possible values for the VirtualSCSISharing const type.
func PossibleVirtualSCSISharingValues() []VirtualSCSISharing {
	return []VirtualSCSISharing{
		VirtualSCSISharingNoSharing,
		VirtualSCSISharingPhysicalSharing,
		VirtualSCSISharingVirtualSharing,
	}
}
