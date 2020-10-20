package hybridnetwork

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// DeviceType enumerates the values for device type.
type DeviceType string

const (
	// DeviceTypeDevicePropertiesFormat ...
	DeviceTypeDevicePropertiesFormat DeviceType = "DevicePropertiesFormat"
)

// PossibleDeviceTypeValues returns an array of possible values for the DeviceType const type.
func PossibleDeviceTypeValues() []DeviceType {
	return []DeviceType{DeviceTypeDevicePropertiesFormat}
}

// DiskCreateOptionTypes enumerates the values for disk create option types.
type DiskCreateOptionTypes string

const (
	// Empty ...
	Empty DiskCreateOptionTypes = "Empty"
	// Unknown ...
	Unknown DiskCreateOptionTypes = "Unknown"
)

// PossibleDiskCreateOptionTypesValues returns an array of possible values for the DiskCreateOptionTypes const type.
func PossibleDiskCreateOptionTypesValues() []DiskCreateOptionTypes {
	return []DiskCreateOptionTypes{Empty, Unknown}
}

// IPAllocationMethod enumerates the values for ip allocation method.
type IPAllocationMethod string

const (
	// IPAllocationMethodDynamic ...
	IPAllocationMethodDynamic IPAllocationMethod = "Dynamic"
	// IPAllocationMethodStatic ...
	IPAllocationMethodStatic IPAllocationMethod = "Static"
	// IPAllocationMethodUnknown ...
	IPAllocationMethodUnknown IPAllocationMethod = "Unknown"
)

// PossibleIPAllocationMethodValues returns an array of possible values for the IPAllocationMethod const type.
func PossibleIPAllocationMethodValues() []IPAllocationMethod {
	return []IPAllocationMethod{IPAllocationMethodDynamic, IPAllocationMethodStatic, IPAllocationMethodUnknown}
}

// IPVersion enumerates the values for ip version.
type IPVersion string

const (
	// IPVersionIPv4 ...
	IPVersionIPv4 IPVersion = "IPv4"
	// IPVersionUnknown ...
	IPVersionUnknown IPVersion = "Unknown"
)

// PossibleIPVersionValues returns an array of possible values for the IPVersion const type.
func PossibleIPVersionValues() []IPVersion {
	return []IPVersion{IPVersionIPv4, IPVersionUnknown}
}

// NetworkFunctionRoleConfigurationType enumerates the values for network function role configuration type.
type NetworkFunctionRoleConfigurationType string

const (
	// NetworkFunctionRoleConfigurationTypeUnknown ...
	NetworkFunctionRoleConfigurationTypeUnknown NetworkFunctionRoleConfigurationType = "Unknown"
	// NetworkFunctionRoleConfigurationTypeVirtualMachine ...
	NetworkFunctionRoleConfigurationTypeVirtualMachine NetworkFunctionRoleConfigurationType = "VirtualMachine"
)

// PossibleNetworkFunctionRoleConfigurationTypeValues returns an array of possible values for the NetworkFunctionRoleConfigurationType const type.
func PossibleNetworkFunctionRoleConfigurationTypeValues() []NetworkFunctionRoleConfigurationType {
	return []NetworkFunctionRoleConfigurationType{NetworkFunctionRoleConfigurationTypeUnknown, NetworkFunctionRoleConfigurationTypeVirtualMachine}
}

// OperatingSystemTypes enumerates the values for operating system types.
type OperatingSystemTypes string

const (
	// OperatingSystemTypesLinux ...
	OperatingSystemTypesLinux OperatingSystemTypes = "Linux"
	// OperatingSystemTypesUnknown ...
	OperatingSystemTypesUnknown OperatingSystemTypes = "Unknown"
	// OperatingSystemTypesWindows ...
	OperatingSystemTypesWindows OperatingSystemTypes = "Windows"
)

// PossibleOperatingSystemTypesValues returns an array of possible values for the OperatingSystemTypes const type.
func PossibleOperatingSystemTypesValues() []OperatingSystemTypes {
	return []OperatingSystemTypes{OperatingSystemTypesLinux, OperatingSystemTypesUnknown, OperatingSystemTypesWindows}
}

// OperationalState enumerates the values for operational state.
type OperationalState string

const (
	// OperationalStateRunning ...
	OperationalStateRunning OperationalState = "Running"
	// OperationalStateStarting ...
	OperationalStateStarting OperationalState = "Starting"
	// OperationalStateStopped ...
	OperationalStateStopped OperationalState = "Stopped"
	// OperationalStateStopping ...
	OperationalStateStopping OperationalState = "Stopping"
	// OperationalStateUnknown ...
	OperationalStateUnknown OperationalState = "Unknown"
)

// PossibleOperationalStateValues returns an array of possible values for the OperationalState const type.
func PossibleOperationalStateValues() []OperationalState {
	return []OperationalState{OperationalStateRunning, OperationalStateStarting, OperationalStateStopped, OperationalStateStopping, OperationalStateUnknown}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// ProvisioningStateAccepted ...
	ProvisioningStateAccepted ProvisioningState = "Accepted"
	// ProvisioningStateCanceled ...
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateDeleted ...
	ProvisioningStateDeleted ProvisioningState = "Deleted"
	// ProvisioningStateDeleting ...
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed ...
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateSucceeded ...
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateUnknown ...
	ProvisioningStateUnknown ProvisioningState = "Unknown"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{ProvisioningStateAccepted, ProvisioningStateCanceled, ProvisioningStateDeleted, ProvisioningStateDeleting, ProvisioningStateFailed, ProvisioningStateSucceeded, ProvisioningStateUnknown}
}

// SkuDeploymentMode enumerates the values for sku deployment mode.
type SkuDeploymentMode string

const (
	// SkuDeploymentModeAzure ...
	SkuDeploymentModeAzure SkuDeploymentMode = "Azure"
	// SkuDeploymentModePrivateEdgeZone ...
	SkuDeploymentModePrivateEdgeZone SkuDeploymentMode = "PrivateEdgeZone"
	// SkuDeploymentModeUnknown ...
	SkuDeploymentModeUnknown SkuDeploymentMode = "Unknown"
)

// PossibleSkuDeploymentModeValues returns an array of possible values for the SkuDeploymentMode const type.
func PossibleSkuDeploymentModeValues() []SkuDeploymentMode {
	return []SkuDeploymentMode{SkuDeploymentModeAzure, SkuDeploymentModePrivateEdgeZone, SkuDeploymentModeUnknown}
}

// SkuType enumerates the values for sku type.
type SkuType string

const (
	// SkuTypeEvolvedPacketCore ...
	SkuTypeEvolvedPacketCore SkuType = "EvolvedPacketCore"
	// SkuTypeFirewall ...
	SkuTypeFirewall SkuType = "Firewall"
	// SkuTypeSDWAN ...
	SkuTypeSDWAN SkuType = "SDWAN"
	// SkuTypeUnknown ...
	SkuTypeUnknown SkuType = "Unknown"
)

// PossibleSkuTypeValues returns an array of possible values for the SkuType const type.
func PossibleSkuTypeValues() []SkuType {
	return []SkuType{SkuTypeEvolvedPacketCore, SkuTypeFirewall, SkuTypeSDWAN, SkuTypeUnknown}
}

// Status enumerates the values for status.
type Status string

const (
	// StatusDeleted ...
	StatusDeleted Status = "Deleted"
	// StatusNotRegistered ...
	StatusNotRegistered Status = "NotRegistered"
	// StatusRegistered ...
	StatusRegistered Status = "Registered"
	// StatusUnknown ...
	StatusUnknown Status = "Unknown"
)

// PossibleStatusValues returns an array of possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{StatusDeleted, StatusNotRegistered, StatusRegistered, StatusUnknown}
}

// VendorProvisioningState enumerates the values for vendor provisioning state.
type VendorProvisioningState string

const (
	// VendorProvisioningStateDeprovisioned ...
	VendorProvisioningStateDeprovisioned VendorProvisioningState = "Deprovisioned"
	// VendorProvisioningStateNotProvisioned ...
	VendorProvisioningStateNotProvisioned VendorProvisioningState = "NotProvisioned"
	// VendorProvisioningStateProvisioned ...
	VendorProvisioningStateProvisioned VendorProvisioningState = "Provisioned"
	// VendorProvisioningStateProvisioning ...
	VendorProvisioningStateProvisioning VendorProvisioningState = "Provisioning"
	// VendorProvisioningStateUnknown ...
	VendorProvisioningStateUnknown VendorProvisioningState = "Unknown"
	// VendorProvisioningStateUserDataValidationFailed ...
	VendorProvisioningStateUserDataValidationFailed VendorProvisioningState = "UserDataValidationFailed"
)

// PossibleVendorProvisioningStateValues returns an array of possible values for the VendorProvisioningState const type.
func PossibleVendorProvisioningStateValues() []VendorProvisioningState {
	return []VendorProvisioningState{VendorProvisioningStateDeprovisioned, VendorProvisioningStateNotProvisioned, VendorProvisioningStateProvisioned, VendorProvisioningStateProvisioning, VendorProvisioningStateUnknown, VendorProvisioningStateUserDataValidationFailed}
}

// VirtualMachineSizeTypes enumerates the values for virtual machine size types.
type VirtualMachineSizeTypes string

const (
	// VirtualMachineSizeTypesStandardD11V2 ...
	VirtualMachineSizeTypesStandardD11V2 VirtualMachineSizeTypes = "Standard_D11_v2"
	// VirtualMachineSizeTypesStandardD12V2 ...
	VirtualMachineSizeTypesStandardD12V2 VirtualMachineSizeTypes = "Standard_D12_v2"
	// VirtualMachineSizeTypesStandardD13V2 ...
	VirtualMachineSizeTypesStandardD13V2 VirtualMachineSizeTypes = "Standard_D13_v2"
	// VirtualMachineSizeTypesStandardD1V2 ...
	VirtualMachineSizeTypesStandardD1V2 VirtualMachineSizeTypes = "Standard_D1_v2"
	// VirtualMachineSizeTypesStandardD2V2 ...
	VirtualMachineSizeTypesStandardD2V2 VirtualMachineSizeTypes = "Standard_D2_v2"
	// VirtualMachineSizeTypesStandardD3V2 ...
	VirtualMachineSizeTypesStandardD3V2 VirtualMachineSizeTypes = "Standard_D3_v2"
	// VirtualMachineSizeTypesStandardD4V2 ...
	VirtualMachineSizeTypesStandardD4V2 VirtualMachineSizeTypes = "Standard_D4_v2"
	// VirtualMachineSizeTypesStandardD5V2 ...
	VirtualMachineSizeTypesStandardD5V2 VirtualMachineSizeTypes = "Standard_D5_v2"
	// VirtualMachineSizeTypesStandardDS11V2 ...
	VirtualMachineSizeTypesStandardDS11V2 VirtualMachineSizeTypes = "Standard_DS11_v2"
	// VirtualMachineSizeTypesStandardDS12V2 ...
	VirtualMachineSizeTypesStandardDS12V2 VirtualMachineSizeTypes = "Standard_DS12_v2"
	// VirtualMachineSizeTypesStandardDS13V2 ...
	VirtualMachineSizeTypesStandardDS13V2 VirtualMachineSizeTypes = "Standard_DS13_v2"
	// VirtualMachineSizeTypesStandardDS1V2 ...
	VirtualMachineSizeTypesStandardDS1V2 VirtualMachineSizeTypes = "Standard_DS1_v2"
	// VirtualMachineSizeTypesStandardDS2V2 ...
	VirtualMachineSizeTypesStandardDS2V2 VirtualMachineSizeTypes = "Standard_DS2_v2"
	// VirtualMachineSizeTypesStandardDS3V2 ...
	VirtualMachineSizeTypesStandardDS3V2 VirtualMachineSizeTypes = "Standard_DS3_v2"
	// VirtualMachineSizeTypesStandardDS4V2 ...
	VirtualMachineSizeTypesStandardDS4V2 VirtualMachineSizeTypes = "Standard_DS4_v2"
	// VirtualMachineSizeTypesStandardDS5V2 ...
	VirtualMachineSizeTypesStandardDS5V2 VirtualMachineSizeTypes = "Standard_DS5_v2"
	// VirtualMachineSizeTypesStandardF1 ...
	VirtualMachineSizeTypesStandardF1 VirtualMachineSizeTypes = "Standard_F1"
	// VirtualMachineSizeTypesStandardF16 ...
	VirtualMachineSizeTypesStandardF16 VirtualMachineSizeTypes = "Standard_F16"
	// VirtualMachineSizeTypesStandardF16s ...
	VirtualMachineSizeTypesStandardF16s VirtualMachineSizeTypes = "Standard_F16s"
	// VirtualMachineSizeTypesStandardF1s ...
	VirtualMachineSizeTypesStandardF1s VirtualMachineSizeTypes = "Standard_F1s"
	// VirtualMachineSizeTypesStandardF2 ...
	VirtualMachineSizeTypesStandardF2 VirtualMachineSizeTypes = "Standard_F2"
	// VirtualMachineSizeTypesStandardF2s ...
	VirtualMachineSizeTypesStandardF2s VirtualMachineSizeTypes = "Standard_F2s"
	// VirtualMachineSizeTypesStandardF4 ...
	VirtualMachineSizeTypesStandardF4 VirtualMachineSizeTypes = "Standard_F4"
	// VirtualMachineSizeTypesStandardF4s ...
	VirtualMachineSizeTypesStandardF4s VirtualMachineSizeTypes = "Standard_F4s"
	// VirtualMachineSizeTypesStandardF8 ...
	VirtualMachineSizeTypesStandardF8 VirtualMachineSizeTypes = "Standard_F8"
	// VirtualMachineSizeTypesStandardF8s ...
	VirtualMachineSizeTypesStandardF8s VirtualMachineSizeTypes = "Standard_F8s"
	// VirtualMachineSizeTypesUnknown ...
	VirtualMachineSizeTypesUnknown VirtualMachineSizeTypes = "Unknown"
)

// PossibleVirtualMachineSizeTypesValues returns an array of possible values for the VirtualMachineSizeTypes const type.
func PossibleVirtualMachineSizeTypesValues() []VirtualMachineSizeTypes {
	return []VirtualMachineSizeTypes{VirtualMachineSizeTypesStandardD11V2, VirtualMachineSizeTypesStandardD12V2, VirtualMachineSizeTypesStandardD13V2, VirtualMachineSizeTypesStandardD1V2, VirtualMachineSizeTypesStandardD2V2, VirtualMachineSizeTypesStandardD3V2, VirtualMachineSizeTypesStandardD4V2, VirtualMachineSizeTypesStandardD5V2, VirtualMachineSizeTypesStandardDS11V2, VirtualMachineSizeTypesStandardDS12V2, VirtualMachineSizeTypesStandardDS13V2, VirtualMachineSizeTypesStandardDS1V2, VirtualMachineSizeTypesStandardDS2V2, VirtualMachineSizeTypesStandardDS3V2, VirtualMachineSizeTypesStandardDS4V2, VirtualMachineSizeTypesStandardDS5V2, VirtualMachineSizeTypesStandardF1, VirtualMachineSizeTypesStandardF16, VirtualMachineSizeTypesStandardF16s, VirtualMachineSizeTypesStandardF1s, VirtualMachineSizeTypesStandardF2, VirtualMachineSizeTypesStandardF2s, VirtualMachineSizeTypesStandardF4, VirtualMachineSizeTypesStandardF4s, VirtualMachineSizeTypesStandardF8, VirtualMachineSizeTypesStandardF8s, VirtualMachineSizeTypesUnknown}
}

// VMSwitchType enumerates the values for vm switch type.
type VMSwitchType string

const (
	// VMSwitchTypeLan ...
	VMSwitchTypeLan VMSwitchType = "Lan"
	// VMSwitchTypeManagement ...
	VMSwitchTypeManagement VMSwitchType = "Management"
	// VMSwitchTypeUnknown ...
	VMSwitchTypeUnknown VMSwitchType = "Unknown"
	// VMSwitchTypeWan ...
	VMSwitchTypeWan VMSwitchType = "Wan"
)

// PossibleVMSwitchTypeValues returns an array of possible values for the VMSwitchType const type.
func PossibleVMSwitchTypeValues() []VMSwitchType {
	return []VMSwitchType{VMSwitchTypeLan, VMSwitchTypeManagement, VMSwitchTypeUnknown, VMSwitchTypeWan}
}
