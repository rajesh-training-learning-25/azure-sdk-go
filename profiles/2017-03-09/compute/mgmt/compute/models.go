//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package compute

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2016-03-30/compute"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type CachingTypes = original.CachingTypes

const (
	None      CachingTypes = original.None
	ReadOnly  CachingTypes = original.ReadOnly
	ReadWrite CachingTypes = original.ReadWrite
)

type ComponentNames = original.ComponentNames

const (
	MicrosoftWindowsShellSetup ComponentNames = original.MicrosoftWindowsShellSetup
)

type DiskCreateOptionTypes = original.DiskCreateOptionTypes

const (
	Attach    DiskCreateOptionTypes = original.Attach
	Empty     DiskCreateOptionTypes = original.Empty
	FromImage DiskCreateOptionTypes = original.FromImage
)

type InstanceViewTypes = original.InstanceViewTypes

const (
	InstanceView InstanceViewTypes = original.InstanceView
)

type OperatingSystemTypes = original.OperatingSystemTypes

const (
	Linux   OperatingSystemTypes = original.Linux
	Windows OperatingSystemTypes = original.Windows
)

type PassNames = original.PassNames

const (
	OobeSystem PassNames = original.OobeSystem
)

type ProtocolTypes = original.ProtocolTypes

const (
	HTTP  ProtocolTypes = original.HTTP
	HTTPS ProtocolTypes = original.HTTPS
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	SystemAssigned ResourceIdentityType = original.SystemAssigned
)

type SettingNames = original.SettingNames

const (
	AutoLogon          SettingNames = original.AutoLogon
	FirstLogonCommands SettingNames = original.FirstLogonCommands
)

type StatusLevelTypes = original.StatusLevelTypes

const (
	Error   StatusLevelTypes = original.Error
	Info    StatusLevelTypes = original.Info
	Warning StatusLevelTypes = original.Warning
)

type UpgradeMode = original.UpgradeMode

const (
	Automatic UpgradeMode = original.Automatic
	Manual    UpgradeMode = original.Manual
)

type VirtualMachineScaleSetSkuScaleType = original.VirtualMachineScaleSetSkuScaleType

const (
	VirtualMachineScaleSetSkuScaleTypeAutomatic VirtualMachineScaleSetSkuScaleType = original.VirtualMachineScaleSetSkuScaleTypeAutomatic
	VirtualMachineScaleSetSkuScaleTypeNone      VirtualMachineScaleSetSkuScaleType = original.VirtualMachineScaleSetSkuScaleTypeNone
)

type VirtualMachineSizeTypes = original.VirtualMachineSizeTypes

const (
	BasicA0        VirtualMachineSizeTypes = original.BasicA0
	BasicA1        VirtualMachineSizeTypes = original.BasicA1
	BasicA2        VirtualMachineSizeTypes = original.BasicA2
	BasicA3        VirtualMachineSizeTypes = original.BasicA3
	BasicA4        VirtualMachineSizeTypes = original.BasicA4
	StandardA0     VirtualMachineSizeTypes = original.StandardA0
	StandardA1     VirtualMachineSizeTypes = original.StandardA1
	StandardA10    VirtualMachineSizeTypes = original.StandardA10
	StandardA11    VirtualMachineSizeTypes = original.StandardA11
	StandardA2     VirtualMachineSizeTypes = original.StandardA2
	StandardA3     VirtualMachineSizeTypes = original.StandardA3
	StandardA4     VirtualMachineSizeTypes = original.StandardA4
	StandardA5     VirtualMachineSizeTypes = original.StandardA5
	StandardA6     VirtualMachineSizeTypes = original.StandardA6
	StandardA7     VirtualMachineSizeTypes = original.StandardA7
	StandardA8     VirtualMachineSizeTypes = original.StandardA8
	StandardA9     VirtualMachineSizeTypes = original.StandardA9
	StandardD1     VirtualMachineSizeTypes = original.StandardD1
	StandardD11    VirtualMachineSizeTypes = original.StandardD11
	StandardD11V2  VirtualMachineSizeTypes = original.StandardD11V2
	StandardD12    VirtualMachineSizeTypes = original.StandardD12
	StandardD12V2  VirtualMachineSizeTypes = original.StandardD12V2
	StandardD13    VirtualMachineSizeTypes = original.StandardD13
	StandardD13V2  VirtualMachineSizeTypes = original.StandardD13V2
	StandardD14    VirtualMachineSizeTypes = original.StandardD14
	StandardD14V2  VirtualMachineSizeTypes = original.StandardD14V2
	StandardD15V2  VirtualMachineSizeTypes = original.StandardD15V2
	StandardD1V2   VirtualMachineSizeTypes = original.StandardD1V2
	StandardD2     VirtualMachineSizeTypes = original.StandardD2
	StandardD2V2   VirtualMachineSizeTypes = original.StandardD2V2
	StandardD3     VirtualMachineSizeTypes = original.StandardD3
	StandardD3V2   VirtualMachineSizeTypes = original.StandardD3V2
	StandardD4     VirtualMachineSizeTypes = original.StandardD4
	StandardD4V2   VirtualMachineSizeTypes = original.StandardD4V2
	StandardD5V2   VirtualMachineSizeTypes = original.StandardD5V2
	StandardDS1    VirtualMachineSizeTypes = original.StandardDS1
	StandardDS11   VirtualMachineSizeTypes = original.StandardDS11
	StandardDS11V2 VirtualMachineSizeTypes = original.StandardDS11V2
	StandardDS12   VirtualMachineSizeTypes = original.StandardDS12
	StandardDS12V2 VirtualMachineSizeTypes = original.StandardDS12V2
	StandardDS13   VirtualMachineSizeTypes = original.StandardDS13
	StandardDS13V2 VirtualMachineSizeTypes = original.StandardDS13V2
	StandardDS14   VirtualMachineSizeTypes = original.StandardDS14
	StandardDS14V2 VirtualMachineSizeTypes = original.StandardDS14V2
	StandardDS15V2 VirtualMachineSizeTypes = original.StandardDS15V2
	StandardDS1V2  VirtualMachineSizeTypes = original.StandardDS1V2
	StandardDS2    VirtualMachineSizeTypes = original.StandardDS2
	StandardDS2V2  VirtualMachineSizeTypes = original.StandardDS2V2
	StandardDS3    VirtualMachineSizeTypes = original.StandardDS3
	StandardDS3V2  VirtualMachineSizeTypes = original.StandardDS3V2
	StandardDS4    VirtualMachineSizeTypes = original.StandardDS4
	StandardDS4V2  VirtualMachineSizeTypes = original.StandardDS4V2
	StandardDS5V2  VirtualMachineSizeTypes = original.StandardDS5V2
	StandardG1     VirtualMachineSizeTypes = original.StandardG1
	StandardG2     VirtualMachineSizeTypes = original.StandardG2
	StandardG3     VirtualMachineSizeTypes = original.StandardG3
	StandardG4     VirtualMachineSizeTypes = original.StandardG4
	StandardG5     VirtualMachineSizeTypes = original.StandardG5
	StandardGS1    VirtualMachineSizeTypes = original.StandardGS1
	StandardGS2    VirtualMachineSizeTypes = original.StandardGS2
	StandardGS3    VirtualMachineSizeTypes = original.StandardGS3
	StandardGS4    VirtualMachineSizeTypes = original.StandardGS4
	StandardGS5    VirtualMachineSizeTypes = original.StandardGS5
)

type APIEntityReference = original.APIEntityReference
type APIError = original.APIError
type APIErrorBase = original.APIErrorBase
type AdditionalUnattendContent = original.AdditionalUnattendContent
type AvailabilitySet = original.AvailabilitySet
type AvailabilitySetListResult = original.AvailabilitySetListResult
type AvailabilitySetListResultIterator = original.AvailabilitySetListResultIterator
type AvailabilitySetListResultPage = original.AvailabilitySetListResultPage
type AvailabilitySetProperties = original.AvailabilitySetProperties
type AvailabilitySetsClient = original.AvailabilitySetsClient
type BaseClient = original.BaseClient
type BootDiagnostics = original.BootDiagnostics
type BootDiagnosticsInstanceView = original.BootDiagnosticsInstanceView
type DataDisk = original.DataDisk
type DataDiskImage = original.DataDiskImage
type DiagnosticsProfile = original.DiagnosticsProfile
type DiskEncryptionSettings = original.DiskEncryptionSettings
type DiskInstanceView = original.DiskInstanceView
type HardwareProfile = original.HardwareProfile
type ImageReference = original.ImageReference
type InnerError = original.InnerError
type InstanceViewStatus = original.InstanceViewStatus
type KeyVaultKeyReference = original.KeyVaultKeyReference
type KeyVaultSecretReference = original.KeyVaultSecretReference
type LinuxConfiguration = original.LinuxConfiguration
type ListUsagesResult = original.ListUsagesResult
type ListUsagesResultIterator = original.ListUsagesResultIterator
type ListUsagesResultPage = original.ListUsagesResultPage
type ListVirtualMachineExtensionImage = original.ListVirtualMachineExtensionImage
type ListVirtualMachineImageResource = original.ListVirtualMachineImageResource
type LongRunningOperationProperties = original.LongRunningOperationProperties
type NetworkInterfaceReference = original.NetworkInterfaceReference
type NetworkInterfaceReferenceProperties = original.NetworkInterfaceReferenceProperties
type NetworkProfile = original.NetworkProfile
type OSDisk = original.OSDisk
type OSDiskImage = original.OSDiskImage
type OSProfile = original.OSProfile
type OperationStatusResponse = original.OperationStatusResponse
type Plan = original.Plan
type PurchasePlan = original.PurchasePlan
type Resource = original.Resource
type SSHConfiguration = original.SSHConfiguration
type SSHPublicKey = original.SSHPublicKey
type Sku = original.Sku
type StorageProfile = original.StorageProfile
type SubResource = original.SubResource
type UpdateResource = original.UpdateResource
type UpgradePolicy = original.UpgradePolicy
type Usage = original.Usage
type UsageClient = original.UsageClient
type UsageName = original.UsageName
type VaultCertificate = original.VaultCertificate
type VaultSecretGroup = original.VaultSecretGroup
type VirtualHardDisk = original.VirtualHardDisk
type VirtualMachine = original.VirtualMachine
type VirtualMachineAgentInstanceView = original.VirtualMachineAgentInstanceView
type VirtualMachineCaptureParameters = original.VirtualMachineCaptureParameters
type VirtualMachineCaptureResult = original.VirtualMachineCaptureResult
type VirtualMachineCaptureResultProperties = original.VirtualMachineCaptureResultProperties
type VirtualMachineExtension = original.VirtualMachineExtension
type VirtualMachineExtensionHandlerInstanceView = original.VirtualMachineExtensionHandlerInstanceView
type VirtualMachineExtensionImage = original.VirtualMachineExtensionImage
type VirtualMachineExtensionImageProperties = original.VirtualMachineExtensionImageProperties
type VirtualMachineExtensionImagesClient = original.VirtualMachineExtensionImagesClient
type VirtualMachineExtensionInstanceView = original.VirtualMachineExtensionInstanceView
type VirtualMachineExtensionProperties = original.VirtualMachineExtensionProperties
type VirtualMachineExtensionUpdate = original.VirtualMachineExtensionUpdate
type VirtualMachineExtensionUpdateProperties = original.VirtualMachineExtensionUpdateProperties
type VirtualMachineExtensionsClient = original.VirtualMachineExtensionsClient
type VirtualMachineExtensionsCreateOrUpdateFuture = original.VirtualMachineExtensionsCreateOrUpdateFuture
type VirtualMachineExtensionsDeleteFuture = original.VirtualMachineExtensionsDeleteFuture
type VirtualMachineExtensionsListResult = original.VirtualMachineExtensionsListResult
type VirtualMachineExtensionsUpdateFuture = original.VirtualMachineExtensionsUpdateFuture
type VirtualMachineIdentity = original.VirtualMachineIdentity
type VirtualMachineImage = original.VirtualMachineImage
type VirtualMachineImageProperties = original.VirtualMachineImageProperties
type VirtualMachineImageResource = original.VirtualMachineImageResource
type VirtualMachineImagesClient = original.VirtualMachineImagesClient
type VirtualMachineInstanceView = original.VirtualMachineInstanceView
type VirtualMachineListResult = original.VirtualMachineListResult
type VirtualMachineListResultIterator = original.VirtualMachineListResultIterator
type VirtualMachineListResultPage = original.VirtualMachineListResultPage
type VirtualMachineProperties = original.VirtualMachineProperties
type VirtualMachineScaleSet = original.VirtualMachineScaleSet
type VirtualMachineScaleSetExtension = original.VirtualMachineScaleSetExtension
type VirtualMachineScaleSetExtensionProfile = original.VirtualMachineScaleSetExtensionProfile
type VirtualMachineScaleSetExtensionProperties = original.VirtualMachineScaleSetExtensionProperties
type VirtualMachineScaleSetIPConfiguration = original.VirtualMachineScaleSetIPConfiguration
type VirtualMachineScaleSetIPConfigurationProperties = original.VirtualMachineScaleSetIPConfigurationProperties
type VirtualMachineScaleSetIdentity = original.VirtualMachineScaleSetIdentity
type VirtualMachineScaleSetInstanceView = original.VirtualMachineScaleSetInstanceView
type VirtualMachineScaleSetInstanceViewStatusesSummary = original.VirtualMachineScaleSetInstanceViewStatusesSummary
type VirtualMachineScaleSetListResult = original.VirtualMachineScaleSetListResult
type VirtualMachineScaleSetListResultIterator = original.VirtualMachineScaleSetListResultIterator
type VirtualMachineScaleSetListResultPage = original.VirtualMachineScaleSetListResultPage
type VirtualMachineScaleSetListSkusResult = original.VirtualMachineScaleSetListSkusResult
type VirtualMachineScaleSetListSkusResultIterator = original.VirtualMachineScaleSetListSkusResultIterator
type VirtualMachineScaleSetListSkusResultPage = original.VirtualMachineScaleSetListSkusResultPage
type VirtualMachineScaleSetListWithLinkResult = original.VirtualMachineScaleSetListWithLinkResult
type VirtualMachineScaleSetListWithLinkResultIterator = original.VirtualMachineScaleSetListWithLinkResultIterator
type VirtualMachineScaleSetListWithLinkResultPage = original.VirtualMachineScaleSetListWithLinkResultPage
type VirtualMachineScaleSetNetworkConfiguration = original.VirtualMachineScaleSetNetworkConfiguration
type VirtualMachineScaleSetNetworkConfigurationProperties = original.VirtualMachineScaleSetNetworkConfigurationProperties
type VirtualMachineScaleSetNetworkProfile = original.VirtualMachineScaleSetNetworkProfile
type VirtualMachineScaleSetOSDisk = original.VirtualMachineScaleSetOSDisk
type VirtualMachineScaleSetOSProfile = original.VirtualMachineScaleSetOSProfile
type VirtualMachineScaleSetProperties = original.VirtualMachineScaleSetProperties
type VirtualMachineScaleSetSku = original.VirtualMachineScaleSetSku
type VirtualMachineScaleSetSkuCapacity = original.VirtualMachineScaleSetSkuCapacity
type VirtualMachineScaleSetStorageProfile = original.VirtualMachineScaleSetStorageProfile
type VirtualMachineScaleSetVM = original.VirtualMachineScaleSetVM
type VirtualMachineScaleSetVMExtensionsSummary = original.VirtualMachineScaleSetVMExtensionsSummary
type VirtualMachineScaleSetVMInstanceIDs = original.VirtualMachineScaleSetVMInstanceIDs
type VirtualMachineScaleSetVMInstanceRequiredIDs = original.VirtualMachineScaleSetVMInstanceRequiredIDs
type VirtualMachineScaleSetVMInstanceView = original.VirtualMachineScaleSetVMInstanceView
type VirtualMachineScaleSetVMListResult = original.VirtualMachineScaleSetVMListResult
type VirtualMachineScaleSetVMListResultIterator = original.VirtualMachineScaleSetVMListResultIterator
type VirtualMachineScaleSetVMListResultPage = original.VirtualMachineScaleSetVMListResultPage
type VirtualMachineScaleSetVMProfile = original.VirtualMachineScaleSetVMProfile
type VirtualMachineScaleSetVMProperties = original.VirtualMachineScaleSetVMProperties
type VirtualMachineScaleSetVMsClient = original.VirtualMachineScaleSetVMsClient
type VirtualMachineScaleSetVMsDeallocateFuture = original.VirtualMachineScaleSetVMsDeallocateFuture
type VirtualMachineScaleSetVMsDeleteFuture = original.VirtualMachineScaleSetVMsDeleteFuture
type VirtualMachineScaleSetVMsPowerOffFuture = original.VirtualMachineScaleSetVMsPowerOffFuture
type VirtualMachineScaleSetVMsReimageFuture = original.VirtualMachineScaleSetVMsReimageFuture
type VirtualMachineScaleSetVMsRestartFuture = original.VirtualMachineScaleSetVMsRestartFuture
type VirtualMachineScaleSetVMsStartFuture = original.VirtualMachineScaleSetVMsStartFuture
type VirtualMachineScaleSetsClient = original.VirtualMachineScaleSetsClient
type VirtualMachineScaleSetsCreateOrUpdateFuture = original.VirtualMachineScaleSetsCreateOrUpdateFuture
type VirtualMachineScaleSetsDeallocateFuture = original.VirtualMachineScaleSetsDeallocateFuture
type VirtualMachineScaleSetsDeleteFuture = original.VirtualMachineScaleSetsDeleteFuture
type VirtualMachineScaleSetsDeleteInstancesFuture = original.VirtualMachineScaleSetsDeleteInstancesFuture
type VirtualMachineScaleSetsPowerOffFuture = original.VirtualMachineScaleSetsPowerOffFuture
type VirtualMachineScaleSetsReimageFuture = original.VirtualMachineScaleSetsReimageFuture
type VirtualMachineScaleSetsRestartFuture = original.VirtualMachineScaleSetsRestartFuture
type VirtualMachineScaleSetsStartFuture = original.VirtualMachineScaleSetsStartFuture
type VirtualMachineScaleSetsUpdateInstancesFuture = original.VirtualMachineScaleSetsUpdateInstancesFuture
type VirtualMachineSize = original.VirtualMachineSize
type VirtualMachineSizeListResult = original.VirtualMachineSizeListResult
type VirtualMachineSizesClient = original.VirtualMachineSizesClient
type VirtualMachineStatusCodeCount = original.VirtualMachineStatusCodeCount
type VirtualMachinesCaptureFuture = original.VirtualMachinesCaptureFuture
type VirtualMachinesClient = original.VirtualMachinesClient
type VirtualMachinesCreateOrUpdateFuture = original.VirtualMachinesCreateOrUpdateFuture
type VirtualMachinesDeallocateFuture = original.VirtualMachinesDeallocateFuture
type VirtualMachinesDeleteFuture = original.VirtualMachinesDeleteFuture
type VirtualMachinesPowerOffFuture = original.VirtualMachinesPowerOffFuture
type VirtualMachinesRedeployFuture = original.VirtualMachinesRedeployFuture
type VirtualMachinesRestartFuture = original.VirtualMachinesRestartFuture
type VirtualMachinesStartFuture = original.VirtualMachinesStartFuture
type WinRMConfiguration = original.WinRMConfiguration
type WinRMListener = original.WinRMListener
type WindowsConfiguration = original.WindowsConfiguration

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAvailabilitySetListResultIterator(page AvailabilitySetListResultPage) AvailabilitySetListResultIterator {
	return original.NewAvailabilitySetListResultIterator(page)
}
func NewAvailabilitySetListResultPage(cur AvailabilitySetListResult, getNextPage func(context.Context, AvailabilitySetListResult) (AvailabilitySetListResult, error)) AvailabilitySetListResultPage {
	return original.NewAvailabilitySetListResultPage(cur, getNextPage)
}
func NewAvailabilitySetsClient(subscriptionID string) AvailabilitySetsClient {
	return original.NewAvailabilitySetsClient(subscriptionID)
}
func NewAvailabilitySetsClientWithBaseURI(baseURI string, subscriptionID string) AvailabilitySetsClient {
	return original.NewAvailabilitySetsClientWithBaseURI(baseURI, subscriptionID)
}
func NewListUsagesResultIterator(page ListUsagesResultPage) ListUsagesResultIterator {
	return original.NewListUsagesResultIterator(page)
}
func NewListUsagesResultPage(cur ListUsagesResult, getNextPage func(context.Context, ListUsagesResult) (ListUsagesResult, error)) ListUsagesResultPage {
	return original.NewListUsagesResultPage(cur, getNextPage)
}
func NewUsageClient(subscriptionID string) UsageClient {
	return original.NewUsageClient(subscriptionID)
}
func NewUsageClientWithBaseURI(baseURI string, subscriptionID string) UsageClient {
	return original.NewUsageClientWithBaseURI(baseURI, subscriptionID)
}
func NewVirtualMachineExtensionImagesClient(subscriptionID string) VirtualMachineExtensionImagesClient {
	return original.NewVirtualMachineExtensionImagesClient(subscriptionID)
}
func NewVirtualMachineExtensionImagesClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineExtensionImagesClient {
	return original.NewVirtualMachineExtensionImagesClientWithBaseURI(baseURI, subscriptionID)
}
func NewVirtualMachineExtensionsClient(subscriptionID string) VirtualMachineExtensionsClient {
	return original.NewVirtualMachineExtensionsClient(subscriptionID)
}
func NewVirtualMachineExtensionsClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineExtensionsClient {
	return original.NewVirtualMachineExtensionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewVirtualMachineImagesClient(subscriptionID string) VirtualMachineImagesClient {
	return original.NewVirtualMachineImagesClient(subscriptionID)
}
func NewVirtualMachineImagesClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineImagesClient {
	return original.NewVirtualMachineImagesClientWithBaseURI(baseURI, subscriptionID)
}
func NewVirtualMachineListResultIterator(page VirtualMachineListResultPage) VirtualMachineListResultIterator {
	return original.NewVirtualMachineListResultIterator(page)
}
func NewVirtualMachineListResultPage(cur VirtualMachineListResult, getNextPage func(context.Context, VirtualMachineListResult) (VirtualMachineListResult, error)) VirtualMachineListResultPage {
	return original.NewVirtualMachineListResultPage(cur, getNextPage)
}
func NewVirtualMachineScaleSetListResultIterator(page VirtualMachineScaleSetListResultPage) VirtualMachineScaleSetListResultIterator {
	return original.NewVirtualMachineScaleSetListResultIterator(page)
}
func NewVirtualMachineScaleSetListResultPage(cur VirtualMachineScaleSetListResult, getNextPage func(context.Context, VirtualMachineScaleSetListResult) (VirtualMachineScaleSetListResult, error)) VirtualMachineScaleSetListResultPage {
	return original.NewVirtualMachineScaleSetListResultPage(cur, getNextPage)
}
func NewVirtualMachineScaleSetListSkusResultIterator(page VirtualMachineScaleSetListSkusResultPage) VirtualMachineScaleSetListSkusResultIterator {
	return original.NewVirtualMachineScaleSetListSkusResultIterator(page)
}
func NewVirtualMachineScaleSetListSkusResultPage(cur VirtualMachineScaleSetListSkusResult, getNextPage func(context.Context, VirtualMachineScaleSetListSkusResult) (VirtualMachineScaleSetListSkusResult, error)) VirtualMachineScaleSetListSkusResultPage {
	return original.NewVirtualMachineScaleSetListSkusResultPage(cur, getNextPage)
}
func NewVirtualMachineScaleSetListWithLinkResultIterator(page VirtualMachineScaleSetListWithLinkResultPage) VirtualMachineScaleSetListWithLinkResultIterator {
	return original.NewVirtualMachineScaleSetListWithLinkResultIterator(page)
}
func NewVirtualMachineScaleSetListWithLinkResultPage(cur VirtualMachineScaleSetListWithLinkResult, getNextPage func(context.Context, VirtualMachineScaleSetListWithLinkResult) (VirtualMachineScaleSetListWithLinkResult, error)) VirtualMachineScaleSetListWithLinkResultPage {
	return original.NewVirtualMachineScaleSetListWithLinkResultPage(cur, getNextPage)
}
func NewVirtualMachineScaleSetVMListResultIterator(page VirtualMachineScaleSetVMListResultPage) VirtualMachineScaleSetVMListResultIterator {
	return original.NewVirtualMachineScaleSetVMListResultIterator(page)
}
func NewVirtualMachineScaleSetVMListResultPage(cur VirtualMachineScaleSetVMListResult, getNextPage func(context.Context, VirtualMachineScaleSetVMListResult) (VirtualMachineScaleSetVMListResult, error)) VirtualMachineScaleSetVMListResultPage {
	return original.NewVirtualMachineScaleSetVMListResultPage(cur, getNextPage)
}
func NewVirtualMachineScaleSetVMsClient(subscriptionID string) VirtualMachineScaleSetVMsClient {
	return original.NewVirtualMachineScaleSetVMsClient(subscriptionID)
}
func NewVirtualMachineScaleSetVMsClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineScaleSetVMsClient {
	return original.NewVirtualMachineScaleSetVMsClientWithBaseURI(baseURI, subscriptionID)
}
func NewVirtualMachineScaleSetsClient(subscriptionID string) VirtualMachineScaleSetsClient {
	return original.NewVirtualMachineScaleSetsClient(subscriptionID)
}
func NewVirtualMachineScaleSetsClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineScaleSetsClient {
	return original.NewVirtualMachineScaleSetsClientWithBaseURI(baseURI, subscriptionID)
}
func NewVirtualMachineSizesClient(subscriptionID string) VirtualMachineSizesClient {
	return original.NewVirtualMachineSizesClient(subscriptionID)
}
func NewVirtualMachineSizesClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineSizesClient {
	return original.NewVirtualMachineSizesClientWithBaseURI(baseURI, subscriptionID)
}
func NewVirtualMachinesClient(subscriptionID string) VirtualMachinesClient {
	return original.NewVirtualMachinesClient(subscriptionID)
}
func NewVirtualMachinesClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachinesClient {
	return original.NewVirtualMachinesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleCachingTypesValues() []CachingTypes {
	return original.PossibleCachingTypesValues()
}
func PossibleComponentNamesValues() []ComponentNames {
	return original.PossibleComponentNamesValues()
}
func PossibleDiskCreateOptionTypesValues() []DiskCreateOptionTypes {
	return original.PossibleDiskCreateOptionTypesValues()
}
func PossibleInstanceViewTypesValues() []InstanceViewTypes {
	return original.PossibleInstanceViewTypesValues()
}
func PossibleOperatingSystemTypesValues() []OperatingSystemTypes {
	return original.PossibleOperatingSystemTypesValues()
}
func PossiblePassNamesValues() []PassNames {
	return original.PossiblePassNamesValues()
}
func PossibleProtocolTypesValues() []ProtocolTypes {
	return original.PossibleProtocolTypesValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func PossibleSettingNamesValues() []SettingNames {
	return original.PossibleSettingNamesValues()
}
func PossibleStatusLevelTypesValues() []StatusLevelTypes {
	return original.PossibleStatusLevelTypesValues()
}
func PossibleUpgradeModeValues() []UpgradeMode {
	return original.PossibleUpgradeModeValues()
}
func PossibleVirtualMachineScaleSetSkuScaleTypeValues() []VirtualMachineScaleSetSkuScaleType {
	return original.PossibleVirtualMachineScaleSetSkuScaleTypeValues()
}
func PossibleVirtualMachineSizeTypesValues() []VirtualMachineSizeTypes {
	return original.PossibleVirtualMachineSizeTypesValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/2017-03-09"
}
func Version() string {
	return original.Version()
}
