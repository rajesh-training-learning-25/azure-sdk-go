//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package batch

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/batch/mgmt/2021-06-01/batch"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AccountKeyType = original.AccountKeyType

const (
	AccountKeyTypePrimary   AccountKeyType = original.AccountKeyTypePrimary
	AccountKeyTypeSecondary AccountKeyType = original.AccountKeyTypeSecondary
)

type AllocationState = original.AllocationState

const (
	AllocationStateResizing AllocationState = original.AllocationStateResizing
	AllocationStateSteady   AllocationState = original.AllocationStateSteady
	AllocationStateStopping AllocationState = original.AllocationStateStopping
)

type AuthenticationMode = original.AuthenticationMode

const (
	AuthenticationModeAAD                     AuthenticationMode = original.AuthenticationModeAAD
	AuthenticationModeSharedKey               AuthenticationMode = original.AuthenticationModeSharedKey
	AuthenticationModeTaskAuthenticationToken AuthenticationMode = original.AuthenticationModeTaskAuthenticationToken
)

type AutoStorageAuthenticationMode = original.AutoStorageAuthenticationMode

const (
	AutoStorageAuthenticationModeBatchAccountManagedIdentity AutoStorageAuthenticationMode = original.AutoStorageAuthenticationModeBatchAccountManagedIdentity
	AutoStorageAuthenticationModeStorageKeys                 AutoStorageAuthenticationMode = original.AutoStorageAuthenticationModeStorageKeys
)

type AutoUserScope = original.AutoUserScope

const (
	AutoUserScopePool AutoUserScope = original.AutoUserScopePool
	AutoUserScopeTask AutoUserScope = original.AutoUserScopeTask
)

type CachingType = original.CachingType

const (
	CachingTypeNone      CachingType = original.CachingTypeNone
	CachingTypeReadOnly  CachingType = original.CachingTypeReadOnly
	CachingTypeReadWrite CachingType = original.CachingTypeReadWrite
)

type CertificateFormat = original.CertificateFormat

const (
	CertificateFormatCer CertificateFormat = original.CertificateFormatCer
	CertificateFormatPfx CertificateFormat = original.CertificateFormatPfx
)

type CertificateProvisioningState = original.CertificateProvisioningState

const (
	CertificateProvisioningStateDeleting  CertificateProvisioningState = original.CertificateProvisioningStateDeleting
	CertificateProvisioningStateFailed    CertificateProvisioningState = original.CertificateProvisioningStateFailed
	CertificateProvisioningStateSucceeded CertificateProvisioningState = original.CertificateProvisioningStateSucceeded
)

type CertificateStoreLocation = original.CertificateStoreLocation

const (
	CertificateStoreLocationCurrentUser  CertificateStoreLocation = original.CertificateStoreLocationCurrentUser
	CertificateStoreLocationLocalMachine CertificateStoreLocation = original.CertificateStoreLocationLocalMachine
)

type CertificateVisibility = original.CertificateVisibility

const (
	CertificateVisibilityRemoteUser CertificateVisibility = original.CertificateVisibilityRemoteUser
	CertificateVisibilityStartTask  CertificateVisibility = original.CertificateVisibilityStartTask
	CertificateVisibilityTask       CertificateVisibility = original.CertificateVisibilityTask
)

type ComputeNodeDeallocationOption = original.ComputeNodeDeallocationOption

const (
	ComputeNodeDeallocationOptionRequeue        ComputeNodeDeallocationOption = original.ComputeNodeDeallocationOptionRequeue
	ComputeNodeDeallocationOptionRetainedData   ComputeNodeDeallocationOption = original.ComputeNodeDeallocationOptionRetainedData
	ComputeNodeDeallocationOptionTaskCompletion ComputeNodeDeallocationOption = original.ComputeNodeDeallocationOptionTaskCompletion
	ComputeNodeDeallocationOptionTerminate      ComputeNodeDeallocationOption = original.ComputeNodeDeallocationOptionTerminate
)

type ComputeNodeFillType = original.ComputeNodeFillType

const (
	ComputeNodeFillTypePack   ComputeNodeFillType = original.ComputeNodeFillTypePack
	ComputeNodeFillTypeSpread ComputeNodeFillType = original.ComputeNodeFillTypeSpread
)

type ContainerWorkingDirectory = original.ContainerWorkingDirectory

const (
	ContainerWorkingDirectoryContainerImageDefault ContainerWorkingDirectory = original.ContainerWorkingDirectoryContainerImageDefault
	ContainerWorkingDirectoryTaskWorkingDirectory  ContainerWorkingDirectory = original.ContainerWorkingDirectoryTaskWorkingDirectory
)

type DiffDiskPlacement = original.DiffDiskPlacement

const (
	DiffDiskPlacementCacheDisk DiffDiskPlacement = original.DiffDiskPlacementCacheDisk
)

type DiskEncryptionTarget = original.DiskEncryptionTarget

const (
	DiskEncryptionTargetOsDisk        DiskEncryptionTarget = original.DiskEncryptionTargetOsDisk
	DiskEncryptionTargetTemporaryDisk DiskEncryptionTarget = original.DiskEncryptionTargetTemporaryDisk
)

type ElevationLevel = original.ElevationLevel

const (
	ElevationLevelAdmin    ElevationLevel = original.ElevationLevelAdmin
	ElevationLevelNonAdmin ElevationLevel = original.ElevationLevelNonAdmin
)

type IPAddressProvisioningType = original.IPAddressProvisioningType

const (
	IPAddressProvisioningTypeBatchManaged        IPAddressProvisioningType = original.IPAddressProvisioningTypeBatchManaged
	IPAddressProvisioningTypeNoPublicIPAddresses IPAddressProvisioningType = original.IPAddressProvisioningTypeNoPublicIPAddresses
	IPAddressProvisioningTypeUserManaged         IPAddressProvisioningType = original.IPAddressProvisioningTypeUserManaged
)

type InboundEndpointProtocol = original.InboundEndpointProtocol

const (
	InboundEndpointProtocolTCP InboundEndpointProtocol = original.InboundEndpointProtocolTCP
	InboundEndpointProtocolUDP InboundEndpointProtocol = original.InboundEndpointProtocolUDP
)

type InterNodeCommunicationState = original.InterNodeCommunicationState

const (
	InterNodeCommunicationStateDisabled InterNodeCommunicationState = original.InterNodeCommunicationStateDisabled
	InterNodeCommunicationStateEnabled  InterNodeCommunicationState = original.InterNodeCommunicationStateEnabled
)

type KeySource = original.KeySource

const (
	KeySourceMicrosoftBatch    KeySource = original.KeySourceMicrosoftBatch
	KeySourceMicrosoftKeyVault KeySource = original.KeySourceMicrosoftKeyVault
)

type LoginMode = original.LoginMode

const (
	LoginModeBatch       LoginMode = original.LoginModeBatch
	LoginModeInteractive LoginMode = original.LoginModeInteractive
)

type NameAvailabilityReason = original.NameAvailabilityReason

const (
	NameAvailabilityReasonAlreadyExists NameAvailabilityReason = original.NameAvailabilityReasonAlreadyExists
	NameAvailabilityReasonInvalid       NameAvailabilityReason = original.NameAvailabilityReasonInvalid
)

type NetworkSecurityGroupRuleAccess = original.NetworkSecurityGroupRuleAccess

const (
	NetworkSecurityGroupRuleAccessAllow NetworkSecurityGroupRuleAccess = original.NetworkSecurityGroupRuleAccessAllow
	NetworkSecurityGroupRuleAccessDeny  NetworkSecurityGroupRuleAccess = original.NetworkSecurityGroupRuleAccessDeny
)

type NodePlacementPolicyType = original.NodePlacementPolicyType

const (
	NodePlacementPolicyTypeRegional NodePlacementPolicyType = original.NodePlacementPolicyTypeRegional
	NodePlacementPolicyTypeZonal    NodePlacementPolicyType = original.NodePlacementPolicyTypeZonal
)

type PackageState = original.PackageState

const (
	PackageStateActive  PackageState = original.PackageStateActive
	PackageStatePending PackageState = original.PackageStatePending
)

type PoolAllocationMode = original.PoolAllocationMode

const (
	PoolAllocationModeBatchService     PoolAllocationMode = original.PoolAllocationModeBatchService
	PoolAllocationModeUserSubscription PoolAllocationMode = original.PoolAllocationModeUserSubscription
)

type PoolIdentityType = original.PoolIdentityType

const (
	PoolIdentityTypeNone         PoolIdentityType = original.PoolIdentityTypeNone
	PoolIdentityTypeUserAssigned PoolIdentityType = original.PoolIdentityTypeUserAssigned
)

type PoolProvisioningState = original.PoolProvisioningState

const (
	PoolProvisioningStateDeleting  PoolProvisioningState = original.PoolProvisioningStateDeleting
	PoolProvisioningStateSucceeded PoolProvisioningState = original.PoolProvisioningStateSucceeded
)

type PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningState

const (
	PrivateEndpointConnectionProvisioningStateFailed    PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateFailed
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateSucceeded
	PrivateEndpointConnectionProvisioningStateUpdating  PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateUpdating
)

type PrivateLinkServiceConnectionStatus = original.PrivateLinkServiceConnectionStatus

const (
	PrivateLinkServiceConnectionStatusApproved     PrivateLinkServiceConnectionStatus = original.PrivateLinkServiceConnectionStatusApproved
	PrivateLinkServiceConnectionStatusDisconnected PrivateLinkServiceConnectionStatus = original.PrivateLinkServiceConnectionStatusDisconnected
	PrivateLinkServiceConnectionStatusPending      PrivateLinkServiceConnectionStatus = original.PrivateLinkServiceConnectionStatusPending
	PrivateLinkServiceConnectionStatusRejected     PrivateLinkServiceConnectionStatus = original.PrivateLinkServiceConnectionStatusRejected
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateCancelled ProvisioningState = original.ProvisioningStateCancelled
	ProvisioningStateCreating  ProvisioningState = original.ProvisioningStateCreating
	ProvisioningStateDeleting  ProvisioningState = original.ProvisioningStateDeleting
	ProvisioningStateFailed    ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStateInvalid   ProvisioningState = original.ProvisioningStateInvalid
	ProvisioningStateSucceeded ProvisioningState = original.ProvisioningStateSucceeded
)

type PublicNetworkAccessType = original.PublicNetworkAccessType

const (
	PublicNetworkAccessTypeDisabled PublicNetworkAccessType = original.PublicNetworkAccessTypeDisabled
	PublicNetworkAccessTypeEnabled  PublicNetworkAccessType = original.PublicNetworkAccessTypeEnabled
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	ResourceIdentityTypeNone           ResourceIdentityType = original.ResourceIdentityTypeNone
	ResourceIdentityTypeSystemAssigned ResourceIdentityType = original.ResourceIdentityTypeSystemAssigned
	ResourceIdentityTypeUserAssigned   ResourceIdentityType = original.ResourceIdentityTypeUserAssigned
)

type StorageAccountType = original.StorageAccountType

const (
	StorageAccountTypePremiumLRS  StorageAccountType = original.StorageAccountTypePremiumLRS
	StorageAccountTypeStandardLRS StorageAccountType = original.StorageAccountTypeStandardLRS
)

type Account = original.Account
type AccountClient = original.AccountClient
type AccountCreateFuture = original.AccountCreateFuture
type AccountCreateParameters = original.AccountCreateParameters
type AccountCreateProperties = original.AccountCreateProperties
type AccountDeleteFuture = original.AccountDeleteFuture
type AccountIdentity = original.AccountIdentity
type AccountKeys = original.AccountKeys
type AccountListResult = original.AccountListResult
type AccountListResultIterator = original.AccountListResultIterator
type AccountListResultPage = original.AccountListResultPage
type AccountProperties = original.AccountProperties
type AccountRegenerateKeyParameters = original.AccountRegenerateKeyParameters
type AccountUpdateParameters = original.AccountUpdateParameters
type AccountUpdateProperties = original.AccountUpdateProperties
type ActivateApplicationPackageParameters = original.ActivateApplicationPackageParameters
type Application = original.Application
type ApplicationClient = original.ApplicationClient
type ApplicationPackage = original.ApplicationPackage
type ApplicationPackageClient = original.ApplicationPackageClient
type ApplicationPackageProperties = original.ApplicationPackageProperties
type ApplicationPackageReference = original.ApplicationPackageReference
type ApplicationProperties = original.ApplicationProperties
type AutoScaleRun = original.AutoScaleRun
type AutoScaleRunError = original.AutoScaleRunError
type AutoScaleSettings = original.AutoScaleSettings
type AutoStorageBaseProperties = original.AutoStorageBaseProperties
type AutoStorageProperties = original.AutoStorageProperties
type AutoUserSpecification = original.AutoUserSpecification
type AzureBlobFileSystemConfiguration = original.AzureBlobFileSystemConfiguration
type AzureFileShareConfiguration = original.AzureFileShareConfiguration
type BaseClient = original.BaseClient
type CIFSMountConfiguration = original.CIFSMountConfiguration
type Certificate = original.Certificate
type CertificateBaseProperties = original.CertificateBaseProperties
type CertificateClient = original.CertificateClient
type CertificateCreateOrUpdateParameters = original.CertificateCreateOrUpdateParameters
type CertificateCreateOrUpdateProperties = original.CertificateCreateOrUpdateProperties
type CertificateDeleteFuture = original.CertificateDeleteFuture
type CertificateProperties = original.CertificateProperties
type CertificateReference = original.CertificateReference
type CheckNameAvailabilityParameters = original.CheckNameAvailabilityParameters
type CheckNameAvailabilityResult = original.CheckNameAvailabilityResult
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type CloudServiceConfiguration = original.CloudServiceConfiguration
type ComputeNodeIdentityReference = original.ComputeNodeIdentityReference
type ContainerConfiguration = original.ContainerConfiguration
type ContainerRegistry = original.ContainerRegistry
type DataDisk = original.DataDisk
type DeleteCertificateError = original.DeleteCertificateError
type DeploymentConfiguration = original.DeploymentConfiguration
type DiffDiskSettings = original.DiffDiskSettings
type DiskEncryptionConfiguration = original.DiskEncryptionConfiguration
type EncryptionProperties = original.EncryptionProperties
type EndpointDependency = original.EndpointDependency
type EndpointDetail = original.EndpointDetail
type EnvironmentSetting = original.EnvironmentSetting
type FixedScaleSettings = original.FixedScaleSettings
type ImageReference = original.ImageReference
type InboundNatPool = original.InboundNatPool
type KeyVaultProperties = original.KeyVaultProperties
type KeyVaultReference = original.KeyVaultReference
type LinuxUserConfiguration = original.LinuxUserConfiguration
type ListApplicationPackagesResult = original.ListApplicationPackagesResult
type ListApplicationPackagesResultIterator = original.ListApplicationPackagesResultIterator
type ListApplicationPackagesResultPage = original.ListApplicationPackagesResultPage
type ListApplicationsResult = original.ListApplicationsResult
type ListApplicationsResultIterator = original.ListApplicationsResultIterator
type ListApplicationsResultPage = original.ListApplicationsResultPage
type ListCertificatesResult = original.ListCertificatesResult
type ListCertificatesResultIterator = original.ListCertificatesResultIterator
type ListCertificatesResultPage = original.ListCertificatesResultPage
type ListPoolsResult = original.ListPoolsResult
type ListPoolsResultIterator = original.ListPoolsResultIterator
type ListPoolsResultPage = original.ListPoolsResultPage
type ListPrivateEndpointConnectionsResult = original.ListPrivateEndpointConnectionsResult
type ListPrivateEndpointConnectionsResultIterator = original.ListPrivateEndpointConnectionsResultIterator
type ListPrivateEndpointConnectionsResultPage = original.ListPrivateEndpointConnectionsResultPage
type ListPrivateLinkResourcesResult = original.ListPrivateLinkResourcesResult
type ListPrivateLinkResourcesResultIterator = original.ListPrivateLinkResourcesResultIterator
type ListPrivateLinkResourcesResultPage = original.ListPrivateLinkResourcesResultPage
type LocationClient = original.LocationClient
type LocationQuota = original.LocationQuota
type MetadataItem = original.MetadataItem
type MountConfiguration = original.MountConfiguration
type NFSMountConfiguration = original.NFSMountConfiguration
type NetworkConfiguration = original.NetworkConfiguration
type NetworkSecurityGroupRule = original.NetworkSecurityGroupRule
type NodePlacementConfiguration = original.NodePlacementConfiguration
type OSDisk = original.OSDisk
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type OutboundEnvironmentEndpoint = original.OutboundEnvironmentEndpoint
type OutboundEnvironmentEndpointCollection = original.OutboundEnvironmentEndpointCollection
type OutboundEnvironmentEndpointCollectionIterator = original.OutboundEnvironmentEndpointCollectionIterator
type OutboundEnvironmentEndpointCollectionPage = original.OutboundEnvironmentEndpointCollectionPage
type Pool = original.Pool
type PoolClient = original.PoolClient
type PoolDeleteFuture = original.PoolDeleteFuture
type PoolEndpointConfiguration = original.PoolEndpointConfiguration
type PoolIdentity = original.PoolIdentity
type PoolProperties = original.PoolProperties
type PrivateEndpoint = original.PrivateEndpoint
type PrivateEndpointConnection = original.PrivateEndpointConnection
type PrivateEndpointConnectionClient = original.PrivateEndpointConnectionClient
type PrivateEndpointConnectionProperties = original.PrivateEndpointConnectionProperties
type PrivateEndpointConnectionUpdateFuture = original.PrivateEndpointConnectionUpdateFuture
type PrivateLinkResource = original.PrivateLinkResource
type PrivateLinkResourceClient = original.PrivateLinkResourceClient
type PrivateLinkResourceProperties = original.PrivateLinkResourceProperties
type PrivateLinkServiceConnectionState = original.PrivateLinkServiceConnectionState
type ProxyResource = original.ProxyResource
type PublicIPAddressConfiguration = original.PublicIPAddressConfiguration
type ResizeError = original.ResizeError
type ResizeOperationStatus = original.ResizeOperationStatus
type Resource = original.Resource
type ResourceFile = original.ResourceFile
type ScaleSettings = original.ScaleSettings
type SkuCapability = original.SkuCapability
type StartTask = original.StartTask
type SupportedSku = original.SupportedSku
type SupportedSkusResult = original.SupportedSkusResult
type SupportedSkusResultIterator = original.SupportedSkusResultIterator
type SupportedSkusResultPage = original.SupportedSkusResultPage
type TaskContainerSettings = original.TaskContainerSettings
type TaskSchedulingPolicy = original.TaskSchedulingPolicy
type UserAccount = original.UserAccount
type UserAssignedIdentities = original.UserAssignedIdentities
type UserIdentity = original.UserIdentity
type VMExtension = original.VMExtension
type VirtualMachineConfiguration = original.VirtualMachineConfiguration
type VirtualMachineFamilyCoreQuota = original.VirtualMachineFamilyCoreQuota
type WindowsConfiguration = original.WindowsConfiguration
type WindowsUserConfiguration = original.WindowsUserConfiguration

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAccountClient(subscriptionID string) AccountClient {
	return original.NewAccountClient(subscriptionID)
}
func NewAccountClientWithBaseURI(baseURI string, subscriptionID string) AccountClient {
	return original.NewAccountClientWithBaseURI(baseURI, subscriptionID)
}
func NewAccountListResultIterator(page AccountListResultPage) AccountListResultIterator {
	return original.NewAccountListResultIterator(page)
}
func NewAccountListResultPage(cur AccountListResult, getNextPage func(context.Context, AccountListResult) (AccountListResult, error)) AccountListResultPage {
	return original.NewAccountListResultPage(cur, getNextPage)
}
func NewApplicationClient(subscriptionID string) ApplicationClient {
	return original.NewApplicationClient(subscriptionID)
}
func NewApplicationClientWithBaseURI(baseURI string, subscriptionID string) ApplicationClient {
	return original.NewApplicationClientWithBaseURI(baseURI, subscriptionID)
}
func NewApplicationPackageClient(subscriptionID string) ApplicationPackageClient {
	return original.NewApplicationPackageClient(subscriptionID)
}
func NewApplicationPackageClientWithBaseURI(baseURI string, subscriptionID string) ApplicationPackageClient {
	return original.NewApplicationPackageClientWithBaseURI(baseURI, subscriptionID)
}
func NewCertificateClient(subscriptionID string) CertificateClient {
	return original.NewCertificateClient(subscriptionID)
}
func NewCertificateClientWithBaseURI(baseURI string, subscriptionID string) CertificateClient {
	return original.NewCertificateClientWithBaseURI(baseURI, subscriptionID)
}
func NewListApplicationPackagesResultIterator(page ListApplicationPackagesResultPage) ListApplicationPackagesResultIterator {
	return original.NewListApplicationPackagesResultIterator(page)
}
func NewListApplicationPackagesResultPage(cur ListApplicationPackagesResult, getNextPage func(context.Context, ListApplicationPackagesResult) (ListApplicationPackagesResult, error)) ListApplicationPackagesResultPage {
	return original.NewListApplicationPackagesResultPage(cur, getNextPage)
}
func NewListApplicationsResultIterator(page ListApplicationsResultPage) ListApplicationsResultIterator {
	return original.NewListApplicationsResultIterator(page)
}
func NewListApplicationsResultPage(cur ListApplicationsResult, getNextPage func(context.Context, ListApplicationsResult) (ListApplicationsResult, error)) ListApplicationsResultPage {
	return original.NewListApplicationsResultPage(cur, getNextPage)
}
func NewListCertificatesResultIterator(page ListCertificatesResultPage) ListCertificatesResultIterator {
	return original.NewListCertificatesResultIterator(page)
}
func NewListCertificatesResultPage(cur ListCertificatesResult, getNextPage func(context.Context, ListCertificatesResult) (ListCertificatesResult, error)) ListCertificatesResultPage {
	return original.NewListCertificatesResultPage(cur, getNextPage)
}
func NewListPoolsResultIterator(page ListPoolsResultPage) ListPoolsResultIterator {
	return original.NewListPoolsResultIterator(page)
}
func NewListPoolsResultPage(cur ListPoolsResult, getNextPage func(context.Context, ListPoolsResult) (ListPoolsResult, error)) ListPoolsResultPage {
	return original.NewListPoolsResultPage(cur, getNextPage)
}
func NewListPrivateEndpointConnectionsResultIterator(page ListPrivateEndpointConnectionsResultPage) ListPrivateEndpointConnectionsResultIterator {
	return original.NewListPrivateEndpointConnectionsResultIterator(page)
}
func NewListPrivateEndpointConnectionsResultPage(cur ListPrivateEndpointConnectionsResult, getNextPage func(context.Context, ListPrivateEndpointConnectionsResult) (ListPrivateEndpointConnectionsResult, error)) ListPrivateEndpointConnectionsResultPage {
	return original.NewListPrivateEndpointConnectionsResultPage(cur, getNextPage)
}
func NewListPrivateLinkResourcesResultIterator(page ListPrivateLinkResourcesResultPage) ListPrivateLinkResourcesResultIterator {
	return original.NewListPrivateLinkResourcesResultIterator(page)
}
func NewListPrivateLinkResourcesResultPage(cur ListPrivateLinkResourcesResult, getNextPage func(context.Context, ListPrivateLinkResourcesResult) (ListPrivateLinkResourcesResult, error)) ListPrivateLinkResourcesResultPage {
	return original.NewListPrivateLinkResourcesResultPage(cur, getNextPage)
}
func NewLocationClient(subscriptionID string) LocationClient {
	return original.NewLocationClient(subscriptionID)
}
func NewLocationClientWithBaseURI(baseURI string, subscriptionID string) LocationClient {
	return original.NewLocationClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOutboundEnvironmentEndpointCollectionIterator(page OutboundEnvironmentEndpointCollectionPage) OutboundEnvironmentEndpointCollectionIterator {
	return original.NewOutboundEnvironmentEndpointCollectionIterator(page)
}
func NewOutboundEnvironmentEndpointCollectionPage(cur OutboundEnvironmentEndpointCollection, getNextPage func(context.Context, OutboundEnvironmentEndpointCollection) (OutboundEnvironmentEndpointCollection, error)) OutboundEnvironmentEndpointCollectionPage {
	return original.NewOutboundEnvironmentEndpointCollectionPage(cur, getNextPage)
}
func NewPoolClient(subscriptionID string) PoolClient {
	return original.NewPoolClient(subscriptionID)
}
func NewPoolClientWithBaseURI(baseURI string, subscriptionID string) PoolClient {
	return original.NewPoolClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateEndpointConnectionClient(subscriptionID string) PrivateEndpointConnectionClient {
	return original.NewPrivateEndpointConnectionClient(subscriptionID)
}
func NewPrivateEndpointConnectionClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointConnectionClient {
	return original.NewPrivateEndpointConnectionClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateLinkResourceClient(subscriptionID string) PrivateLinkResourceClient {
	return original.NewPrivateLinkResourceClient(subscriptionID)
}
func NewPrivateLinkResourceClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkResourceClient {
	return original.NewPrivateLinkResourceClientWithBaseURI(baseURI, subscriptionID)
}
func NewSupportedSkusResultIterator(page SupportedSkusResultPage) SupportedSkusResultIterator {
	return original.NewSupportedSkusResultIterator(page)
}
func NewSupportedSkusResultPage(cur SupportedSkusResult, getNextPage func(context.Context, SupportedSkusResult) (SupportedSkusResult, error)) SupportedSkusResultPage {
	return original.NewSupportedSkusResultPage(cur, getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAccountKeyTypeValues() []AccountKeyType {
	return original.PossibleAccountKeyTypeValues()
}
func PossibleAllocationStateValues() []AllocationState {
	return original.PossibleAllocationStateValues()
}
func PossibleAuthenticationModeValues() []AuthenticationMode {
	return original.PossibleAuthenticationModeValues()
}
func PossibleAutoStorageAuthenticationModeValues() []AutoStorageAuthenticationMode {
	return original.PossibleAutoStorageAuthenticationModeValues()
}
func PossibleAutoUserScopeValues() []AutoUserScope {
	return original.PossibleAutoUserScopeValues()
}
func PossibleCachingTypeValues() []CachingType {
	return original.PossibleCachingTypeValues()
}
func PossibleCertificateFormatValues() []CertificateFormat {
	return original.PossibleCertificateFormatValues()
}
func PossibleCertificateProvisioningStateValues() []CertificateProvisioningState {
	return original.PossibleCertificateProvisioningStateValues()
}
func PossibleCertificateStoreLocationValues() []CertificateStoreLocation {
	return original.PossibleCertificateStoreLocationValues()
}
func PossibleCertificateVisibilityValues() []CertificateVisibility {
	return original.PossibleCertificateVisibilityValues()
}
func PossibleComputeNodeDeallocationOptionValues() []ComputeNodeDeallocationOption {
	return original.PossibleComputeNodeDeallocationOptionValues()
}
func PossibleComputeNodeFillTypeValues() []ComputeNodeFillType {
	return original.PossibleComputeNodeFillTypeValues()
}
func PossibleContainerWorkingDirectoryValues() []ContainerWorkingDirectory {
	return original.PossibleContainerWorkingDirectoryValues()
}
func PossibleDiffDiskPlacementValues() []DiffDiskPlacement {
	return original.PossibleDiffDiskPlacementValues()
}
func PossibleDiskEncryptionTargetValues() []DiskEncryptionTarget {
	return original.PossibleDiskEncryptionTargetValues()
}
func PossibleElevationLevelValues() []ElevationLevel {
	return original.PossibleElevationLevelValues()
}
func PossibleIPAddressProvisioningTypeValues() []IPAddressProvisioningType {
	return original.PossibleIPAddressProvisioningTypeValues()
}
func PossibleInboundEndpointProtocolValues() []InboundEndpointProtocol {
	return original.PossibleInboundEndpointProtocolValues()
}
func PossibleInterNodeCommunicationStateValues() []InterNodeCommunicationState {
	return original.PossibleInterNodeCommunicationStateValues()
}
func PossibleKeySourceValues() []KeySource {
	return original.PossibleKeySourceValues()
}
func PossibleLoginModeValues() []LoginMode {
	return original.PossibleLoginModeValues()
}
func PossibleNameAvailabilityReasonValues() []NameAvailabilityReason {
	return original.PossibleNameAvailabilityReasonValues()
}
func PossibleNetworkSecurityGroupRuleAccessValues() []NetworkSecurityGroupRuleAccess {
	return original.PossibleNetworkSecurityGroupRuleAccessValues()
}
func PossibleNodePlacementPolicyTypeValues() []NodePlacementPolicyType {
	return original.PossibleNodePlacementPolicyTypeValues()
}
func PossiblePackageStateValues() []PackageState {
	return original.PossiblePackageStateValues()
}
func PossiblePoolAllocationModeValues() []PoolAllocationMode {
	return original.PossiblePoolAllocationModeValues()
}
func PossiblePoolIdentityTypeValues() []PoolIdentityType {
	return original.PossiblePoolIdentityTypeValues()
}
func PossiblePoolProvisioningStateValues() []PoolProvisioningState {
	return original.PossiblePoolProvisioningStateValues()
}
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return original.PossiblePrivateEndpointConnectionProvisioningStateValues()
}
func PossiblePrivateLinkServiceConnectionStatusValues() []PrivateLinkServiceConnectionStatus {
	return original.PossiblePrivateLinkServiceConnectionStatusValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossiblePublicNetworkAccessTypeValues() []PublicNetworkAccessType {
	return original.PossiblePublicNetworkAccessTypeValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func PossibleStorageAccountTypeValues() []StorageAccountType {
	return original.PossibleStorageAccountTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
