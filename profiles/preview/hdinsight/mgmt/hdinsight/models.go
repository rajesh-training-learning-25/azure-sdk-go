//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package hdinsight

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/hdinsight/mgmt/2018-06-01/hdinsight"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AsyncOperationState = original.AsyncOperationState

const (
	AsyncOperationStateFailed     AsyncOperationState = original.AsyncOperationStateFailed
	AsyncOperationStateInProgress AsyncOperationState = original.AsyncOperationStateInProgress
	AsyncOperationStateSucceeded  AsyncOperationState = original.AsyncOperationStateSucceeded
)

type ClusterProvisioningState = original.ClusterProvisioningState

const (
	ClusterProvisioningStateCanceled   ClusterProvisioningState = original.ClusterProvisioningStateCanceled
	ClusterProvisioningStateDeleting   ClusterProvisioningState = original.ClusterProvisioningStateDeleting
	ClusterProvisioningStateFailed     ClusterProvisioningState = original.ClusterProvisioningStateFailed
	ClusterProvisioningStateInProgress ClusterProvisioningState = original.ClusterProvisioningStateInProgress
	ClusterProvisioningStateSucceeded  ClusterProvisioningState = original.ClusterProvisioningStateSucceeded
)

type DaysOfWeek = original.DaysOfWeek

const (
	DaysOfWeekFriday    DaysOfWeek = original.DaysOfWeekFriday
	DaysOfWeekMonday    DaysOfWeek = original.DaysOfWeekMonday
	DaysOfWeekSaturday  DaysOfWeek = original.DaysOfWeekSaturday
	DaysOfWeekSunday    DaysOfWeek = original.DaysOfWeekSunday
	DaysOfWeekThursday  DaysOfWeek = original.DaysOfWeekThursday
	DaysOfWeekTuesday   DaysOfWeek = original.DaysOfWeekTuesday
	DaysOfWeekWednesday DaysOfWeek = original.DaysOfWeekWednesday
)

type DirectoryType = original.DirectoryType

const (
	DirectoryTypeActiveDirectory DirectoryType = original.DirectoryTypeActiveDirectory
)

type FilterMode = original.FilterMode

const (
	FilterModeDefault   FilterMode = original.FilterModeDefault
	FilterModeExclude   FilterMode = original.FilterModeExclude
	FilterModeInclude   FilterMode = original.FilterModeInclude
	FilterModeRecommend FilterMode = original.FilterModeRecommend
)

type JSONWebKeyEncryptionAlgorithm = original.JSONWebKeyEncryptionAlgorithm

const (
	JSONWebKeyEncryptionAlgorithmRSA15      JSONWebKeyEncryptionAlgorithm = original.JSONWebKeyEncryptionAlgorithmRSA15
	JSONWebKeyEncryptionAlgorithmRSAOAEP    JSONWebKeyEncryptionAlgorithm = original.JSONWebKeyEncryptionAlgorithmRSAOAEP
	JSONWebKeyEncryptionAlgorithmRSAOAEP256 JSONWebKeyEncryptionAlgorithm = original.JSONWebKeyEncryptionAlgorithmRSAOAEP256
)

type OSType = original.OSType

const (
	OSTypeLinux   OSType = original.OSTypeLinux
	OSTypeWindows OSType = original.OSTypeWindows
)

type PrivateLink = original.PrivateLink

const (
	PrivateLinkDisabled PrivateLink = original.PrivateLinkDisabled
	PrivateLinkEnabled  PrivateLink = original.PrivateLinkEnabled
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	ResourceIdentityTypeNone                       ResourceIdentityType = original.ResourceIdentityTypeNone
	ResourceIdentityTypeSystemAssigned             ResourceIdentityType = original.ResourceIdentityTypeSystemAssigned
	ResourceIdentityTypeSystemAssignedUserAssigned ResourceIdentityType = original.ResourceIdentityTypeSystemAssignedUserAssigned
	ResourceIdentityTypeUserAssigned               ResourceIdentityType = original.ResourceIdentityTypeUserAssigned
)

type ResourceProviderConnection = original.ResourceProviderConnection

const (
	ResourceProviderConnectionInbound  ResourceProviderConnection = original.ResourceProviderConnectionInbound
	ResourceProviderConnectionOutbound ResourceProviderConnection = original.ResourceProviderConnectionOutbound
)

type Tier = original.Tier

const (
	TierPremium  Tier = original.TierPremium
	TierStandard Tier = original.TierStandard
)

type AaddsResourceDetails = original.AaddsResourceDetails
type Application = original.Application
type ApplicationGetEndpoint = original.ApplicationGetEndpoint
type ApplicationGetHTTPSEndpoint = original.ApplicationGetHTTPSEndpoint
type ApplicationListResult = original.ApplicationListResult
type ApplicationListResultIterator = original.ApplicationListResultIterator
type ApplicationListResultPage = original.ApplicationListResultPage
type ApplicationProperties = original.ApplicationProperties
type ApplicationsClient = original.ApplicationsClient
type ApplicationsCreateFuture = original.ApplicationsCreateFuture
type ApplicationsDeleteFuture = original.ApplicationsDeleteFuture
type AsyncOperationResult = original.AsyncOperationResult
type Autoscale = original.Autoscale
type AutoscaleCapacity = original.AutoscaleCapacity
type AutoscaleConfigurationUpdateParameter = original.AutoscaleConfigurationUpdateParameter
type AutoscaleRecurrence = original.AutoscaleRecurrence
type AutoscaleSchedule = original.AutoscaleSchedule
type AutoscaleTimeAndCapacity = original.AutoscaleTimeAndCapacity
type AzureMonitorRequest = original.AzureMonitorRequest
type AzureMonitorResponse = original.AzureMonitorResponse
type AzureMonitorSelectedConfigurations = original.AzureMonitorSelectedConfigurations
type AzureMonitorTableConfiguration = original.AzureMonitorTableConfiguration
type BaseClient = original.BaseClient
type BillingMeters = original.BillingMeters
type BillingResources = original.BillingResources
type BillingResponseListResult = original.BillingResponseListResult
type CapabilitiesResult = original.CapabilitiesResult
type ClientGroupInfo = original.ClientGroupInfo
type Cluster = original.Cluster
type ClusterConfigurations = original.ClusterConfigurations
type ClusterCreateParametersExtended = original.ClusterCreateParametersExtended
type ClusterCreateProperties = original.ClusterCreateProperties
type ClusterCreateRequestValidationParameters = original.ClusterCreateRequestValidationParameters
type ClusterCreateValidationResult = original.ClusterCreateValidationResult
type ClusterDefinition = original.ClusterDefinition
type ClusterDiskEncryptionParameters = original.ClusterDiskEncryptionParameters
type ClusterGetProperties = original.ClusterGetProperties
type ClusterIdentity = original.ClusterIdentity
type ClusterIdentityUserAssignedIdentitiesValue = original.ClusterIdentityUserAssignedIdentitiesValue
type ClusterListPersistedScriptActionsResult = original.ClusterListPersistedScriptActionsResult
type ClusterListResult = original.ClusterListResult
type ClusterListResultIterator = original.ClusterListResultIterator
type ClusterListResultPage = original.ClusterListResultPage
type ClusterListRuntimeScriptActionDetailResult = original.ClusterListRuntimeScriptActionDetailResult
type ClusterMonitoringRequest = original.ClusterMonitoringRequest
type ClusterMonitoringResponse = original.ClusterMonitoringResponse
type ClusterPatchParameters = original.ClusterPatchParameters
type ClusterResizeParameters = original.ClusterResizeParameters
type ClustersClient = original.ClustersClient
type ClustersCreateFuture = original.ClustersCreateFuture
type ClustersDeleteFuture = original.ClustersDeleteFuture
type ClustersExecuteScriptActionsFuture = original.ClustersExecuteScriptActionsFuture
type ClustersResizeFuture = original.ClustersResizeFuture
type ClustersRotateDiskEncryptionKeyFuture = original.ClustersRotateDiskEncryptionKeyFuture
type ClustersUpdateAutoScaleConfigurationFuture = original.ClustersUpdateAutoScaleConfigurationFuture
type ClustersUpdateGatewaySettingsFuture = original.ClustersUpdateGatewaySettingsFuture
type ClustersUpdateIdentityCertificateFuture = original.ClustersUpdateIdentityCertificateFuture
type ComputeIsolationProperties = original.ComputeIsolationProperties
type ComputeProfile = original.ComputeProfile
type ConfigurationsClient = original.ConfigurationsClient
type ConfigurationsUpdateFuture = original.ConfigurationsUpdateFuture
type ConnectivityEndpoint = original.ConnectivityEndpoint
type DataDisksGroups = original.DataDisksGroups
type Dimension = original.Dimension
type DiskBillingMeters = original.DiskBillingMeters
type DiskEncryptionProperties = original.DiskEncryptionProperties
type EncryptionInTransitProperties = original.EncryptionInTransitProperties
type ErrorResponse = original.ErrorResponse
type Errors = original.Errors
type ExcludedServicesConfig = original.ExcludedServicesConfig
type ExecuteScriptActionParameters = original.ExecuteScriptActionParameters
type Extension = original.Extension
type ExtensionsClient = original.ExtensionsClient
type ExtensionsCreateFuture = original.ExtensionsCreateFuture
type ExtensionsDeleteFuture = original.ExtensionsDeleteFuture
type ExtensionsDisableAzureMonitorFuture = original.ExtensionsDisableAzureMonitorFuture
type ExtensionsDisableMonitoringFuture = original.ExtensionsDisableMonitoringFuture
type ExtensionsEnableAzureMonitorFuture = original.ExtensionsEnableAzureMonitorFuture
type ExtensionsEnableMonitoringFuture = original.ExtensionsEnableMonitoringFuture
type GatewaySettings = original.GatewaySettings
type HardwareProfile = original.HardwareProfile
type HostInfo = original.HostInfo
type KafkaRestProperties = original.KafkaRestProperties
type LinuxOperatingSystemProfile = original.LinuxOperatingSystemProfile
type ListHostInfo = original.ListHostInfo
type LocalizedName = original.LocalizedName
type LocationsClient = original.LocationsClient
type MetricSpecifications = original.MetricSpecifications
type NameAvailabilityCheckRequestParameters = original.NameAvailabilityCheckRequestParameters
type NameAvailabilityCheckResult = original.NameAvailabilityCheckResult
type NetworkProperties = original.NetworkProperties
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationProperties = original.OperationProperties
type OperationsClient = original.OperationsClient
type OsProfile = original.OsProfile
type ProxyResource = original.ProxyResource
type QuotaCapability = original.QuotaCapability
type QuotaInfo = original.QuotaInfo
type RegionalQuotaCapability = original.RegionalQuotaCapability
type RegionsCapability = original.RegionsCapability
type Resource = original.Resource
type Role = original.Role
type RuntimeScriptAction = original.RuntimeScriptAction
type RuntimeScriptActionDetail = original.RuntimeScriptActionDetail
type SSHProfile = original.SSHProfile
type SSHPublicKey = original.SSHPublicKey
type ScriptAction = original.ScriptAction
type ScriptActionExecutionHistoryList = original.ScriptActionExecutionHistoryList
type ScriptActionExecutionHistoryListIterator = original.ScriptActionExecutionHistoryListIterator
type ScriptActionExecutionHistoryListPage = original.ScriptActionExecutionHistoryListPage
type ScriptActionExecutionSummary = original.ScriptActionExecutionSummary
type ScriptActionPersistedGetResponseSpec = original.ScriptActionPersistedGetResponseSpec
type ScriptActionsClient = original.ScriptActionsClient
type ScriptActionsList = original.ScriptActionsList
type ScriptActionsListIterator = original.ScriptActionsListIterator
type ScriptActionsListPage = original.ScriptActionsListPage
type ScriptExecutionHistoryClient = original.ScriptExecutionHistoryClient
type SecurityProfile = original.SecurityProfile
type ServiceSpecification = original.ServiceSpecification
type SetString = original.SetString
type StorageAccount = original.StorageAccount
type StorageProfile = original.StorageProfile
type TrackedResource = original.TrackedResource
type UpdateClusterIdentityCertificateParameters = original.UpdateClusterIdentityCertificateParameters
type UpdateGatewaySettingsParameters = original.UpdateGatewaySettingsParameters
type Usage = original.Usage
type UsagesListResult = original.UsagesListResult
type VMSizeCompatibilityFilter = original.VMSizeCompatibilityFilter
type VMSizeCompatibilityFilterV2 = original.VMSizeCompatibilityFilterV2
type VMSizeProperty = original.VMSizeProperty
type VMSizesCapability = original.VMSizesCapability
type ValidationErrorInfo = original.ValidationErrorInfo
type VersionSpec = original.VersionSpec
type VersionsCapability = original.VersionsCapability
type VirtualMachinesClient = original.VirtualMachinesClient
type VirtualMachinesRestartHostsFuture = original.VirtualMachinesRestartHostsFuture
type VirtualNetworkProfile = original.VirtualNetworkProfile

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewApplicationListResultIterator(page ApplicationListResultPage) ApplicationListResultIterator {
	return original.NewApplicationListResultIterator(page)
}
func NewApplicationListResultPage(cur ApplicationListResult, getNextPage func(context.Context, ApplicationListResult) (ApplicationListResult, error)) ApplicationListResultPage {
	return original.NewApplicationListResultPage(cur, getNextPage)
}
func NewApplicationsClient(subscriptionID string) ApplicationsClient {
	return original.NewApplicationsClient(subscriptionID)
}
func NewApplicationsClientWithBaseURI(baseURI string, subscriptionID string) ApplicationsClient {
	return original.NewApplicationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewClusterListResultIterator(page ClusterListResultPage) ClusterListResultIterator {
	return original.NewClusterListResultIterator(page)
}
func NewClusterListResultPage(cur ClusterListResult, getNextPage func(context.Context, ClusterListResult) (ClusterListResult, error)) ClusterListResultPage {
	return original.NewClusterListResultPage(cur, getNextPage)
}
func NewClustersClient(subscriptionID string) ClustersClient {
	return original.NewClustersClient(subscriptionID)
}
func NewClustersClientWithBaseURI(baseURI string, subscriptionID string) ClustersClient {
	return original.NewClustersClientWithBaseURI(baseURI, subscriptionID)
}
func NewConfigurationsClient(subscriptionID string) ConfigurationsClient {
	return original.NewConfigurationsClient(subscriptionID)
}
func NewConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) ConfigurationsClient {
	return original.NewConfigurationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewExtensionsClient(subscriptionID string) ExtensionsClient {
	return original.NewExtensionsClient(subscriptionID)
}
func NewExtensionsClientWithBaseURI(baseURI string, subscriptionID string) ExtensionsClient {
	return original.NewExtensionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewLocationsClient(subscriptionID string) LocationsClient {
	return original.NewLocationsClient(subscriptionID)
}
func NewLocationsClientWithBaseURI(baseURI string, subscriptionID string) LocationsClient {
	return original.NewLocationsClientWithBaseURI(baseURI, subscriptionID)
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
func NewScriptActionExecutionHistoryListIterator(page ScriptActionExecutionHistoryListPage) ScriptActionExecutionHistoryListIterator {
	return original.NewScriptActionExecutionHistoryListIterator(page)
}
func NewScriptActionExecutionHistoryListPage(cur ScriptActionExecutionHistoryList, getNextPage func(context.Context, ScriptActionExecutionHistoryList) (ScriptActionExecutionHistoryList, error)) ScriptActionExecutionHistoryListPage {
	return original.NewScriptActionExecutionHistoryListPage(cur, getNextPage)
}
func NewScriptActionsClient(subscriptionID string) ScriptActionsClient {
	return original.NewScriptActionsClient(subscriptionID)
}
func NewScriptActionsClientWithBaseURI(baseURI string, subscriptionID string) ScriptActionsClient {
	return original.NewScriptActionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewScriptActionsListIterator(page ScriptActionsListPage) ScriptActionsListIterator {
	return original.NewScriptActionsListIterator(page)
}
func NewScriptActionsListPage(cur ScriptActionsList, getNextPage func(context.Context, ScriptActionsList) (ScriptActionsList, error)) ScriptActionsListPage {
	return original.NewScriptActionsListPage(cur, getNextPage)
}
func NewScriptExecutionHistoryClient(subscriptionID string) ScriptExecutionHistoryClient {
	return original.NewScriptExecutionHistoryClient(subscriptionID)
}
func NewScriptExecutionHistoryClientWithBaseURI(baseURI string, subscriptionID string) ScriptExecutionHistoryClient {
	return original.NewScriptExecutionHistoryClientWithBaseURI(baseURI, subscriptionID)
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
func PossibleAsyncOperationStateValues() []AsyncOperationState {
	return original.PossibleAsyncOperationStateValues()
}
func PossibleClusterProvisioningStateValues() []ClusterProvisioningState {
	return original.PossibleClusterProvisioningStateValues()
}
func PossibleDaysOfWeekValues() []DaysOfWeek {
	return original.PossibleDaysOfWeekValues()
}
func PossibleDirectoryTypeValues() []DirectoryType {
	return original.PossibleDirectoryTypeValues()
}
func PossibleFilterModeValues() []FilterMode {
	return original.PossibleFilterModeValues()
}
func PossibleJSONWebKeyEncryptionAlgorithmValues() []JSONWebKeyEncryptionAlgorithm {
	return original.PossibleJSONWebKeyEncryptionAlgorithmValues()
}
func PossibleOSTypeValues() []OSType {
	return original.PossibleOSTypeValues()
}
func PossiblePrivateLinkValues() []PrivateLink {
	return original.PossiblePrivateLinkValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func PossibleResourceProviderConnectionValues() []ResourceProviderConnection {
	return original.PossibleResourceProviderConnectionValues()
}
func PossibleTierValues() []Tier {
	return original.PossibleTierValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
