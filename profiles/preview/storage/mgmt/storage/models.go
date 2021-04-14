// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package storage

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-02-01/storage"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AccessTier = original.AccessTier

const (
	Cool AccessTier = original.Cool
	Hot  AccessTier = original.Hot
)

type AccountExpand = original.AccountExpand

const (
	AccountExpandBlobRestoreStatus   AccountExpand = original.AccountExpandBlobRestoreStatus
	AccountExpandGeoReplicationStats AccountExpand = original.AccountExpandGeoReplicationStats
)

type AccountStatus = original.AccountStatus

const (
	Available   AccountStatus = original.Available
	Unavailable AccountStatus = original.Unavailable
)

type Action = original.Action

const (
	Allow Action = original.Allow
)

type Action1 = original.Action1

const (
	Acquire Action1 = original.Acquire
	Break   Action1 = original.Break
	Change  Action1 = original.Change
	Release Action1 = original.Release
	Renew   Action1 = original.Renew
)

type BlobRestoreProgressStatus = original.BlobRestoreProgressStatus

const (
	Complete   BlobRestoreProgressStatus = original.Complete
	Failed     BlobRestoreProgressStatus = original.Failed
	InProgress BlobRestoreProgressStatus = original.InProgress
)

type Bypass = original.Bypass

const (
	AzureServices Bypass = original.AzureServices
	Logging       Bypass = original.Logging
	Metrics       Bypass = original.Metrics
	None          Bypass = original.None
)

type CreatedByType = original.CreatedByType

const (
	Application     CreatedByType = original.Application
	Key             CreatedByType = original.Key
	ManagedIdentity CreatedByType = original.ManagedIdentity
	User            CreatedByType = original.User
)

type DefaultAction = original.DefaultAction

const (
	DefaultActionAllow DefaultAction = original.DefaultActionAllow
	DefaultActionDeny  DefaultAction = original.DefaultActionDeny
)

type DirectoryServiceOptions = original.DirectoryServiceOptions

const (
	DirectoryServiceOptionsAADDS DirectoryServiceOptions = original.DirectoryServiceOptionsAADDS
	DirectoryServiceOptionsAD    DirectoryServiceOptions = original.DirectoryServiceOptionsAD
	DirectoryServiceOptionsNone  DirectoryServiceOptions = original.DirectoryServiceOptionsNone
)

type EnabledProtocols = original.EnabledProtocols

const (
	NFS EnabledProtocols = original.NFS
	SMB EnabledProtocols = original.SMB
)

type EncryptionScopeSource = original.EncryptionScopeSource

const (
	MicrosoftKeyVault EncryptionScopeSource = original.MicrosoftKeyVault
	MicrosoftStorage  EncryptionScopeSource = original.MicrosoftStorage
)

type EncryptionScopeState = original.EncryptionScopeState

const (
	Disabled EncryptionScopeState = original.Disabled
	Enabled  EncryptionScopeState = original.Enabled
)

type ExtendedLocationTypes = original.ExtendedLocationTypes

const (
	EdgeZone ExtendedLocationTypes = original.EdgeZone
)

type GeoReplicationStatus = original.GeoReplicationStatus

const (
	GeoReplicationStatusBootstrap   GeoReplicationStatus = original.GeoReplicationStatusBootstrap
	GeoReplicationStatusLive        GeoReplicationStatus = original.GeoReplicationStatusLive
	GeoReplicationStatusUnavailable GeoReplicationStatus = original.GeoReplicationStatusUnavailable
)

type GetShareExpand = original.GetShareExpand

const (
	Stats GetShareExpand = original.Stats
)

type HTTPProtocol = original.HTTPProtocol

const (
	HTTPS     HTTPProtocol = original.HTTPS
	Httpshttp HTTPProtocol = original.Httpshttp
)

type IdentityType = original.IdentityType

const (
	IdentityTypeNone                       IdentityType = original.IdentityTypeNone
	IdentityTypeSystemAssigned             IdentityType = original.IdentityTypeSystemAssigned
	IdentityTypeSystemAssignedUserAssigned IdentityType = original.IdentityTypeSystemAssignedUserAssigned
	IdentityTypeUserAssigned               IdentityType = original.IdentityTypeUserAssigned
)

type ImmutabilityPolicyState = original.ImmutabilityPolicyState

const (
	Locked   ImmutabilityPolicyState = original.Locked
	Unlocked ImmutabilityPolicyState = original.Unlocked
)

type ImmutabilityPolicyUpdateType = original.ImmutabilityPolicyUpdateType

const (
	Extend ImmutabilityPolicyUpdateType = original.Extend
	Lock   ImmutabilityPolicyUpdateType = original.Lock
	Put    ImmutabilityPolicyUpdateType = original.Put
)

type KeyPermission = original.KeyPermission

const (
	Full KeyPermission = original.Full
	Read KeyPermission = original.Read
)

type KeySource = original.KeySource

const (
	KeySourceMicrosoftKeyvault KeySource = original.KeySourceMicrosoftKeyvault
	KeySourceMicrosoftStorage  KeySource = original.KeySourceMicrosoftStorage
)

type KeyType = original.KeyType

const (
	KeyTypeAccount KeyType = original.KeyTypeAccount
	KeyTypeService KeyType = original.KeyTypeService
)

type Kind = original.Kind

const (
	BlobStorage      Kind = original.BlobStorage
	BlockBlobStorage Kind = original.BlockBlobStorage
	FileStorage      Kind = original.FileStorage
	Storage          Kind = original.Storage
	StorageV2        Kind = original.StorageV2
)

type LargeFileSharesState = original.LargeFileSharesState

const (
	LargeFileSharesStateDisabled LargeFileSharesState = original.LargeFileSharesStateDisabled
	LargeFileSharesStateEnabled  LargeFileSharesState = original.LargeFileSharesStateEnabled
)

type LeaseDuration = original.LeaseDuration

const (
	Fixed    LeaseDuration = original.Fixed
	Infinite LeaseDuration = original.Infinite
)

type LeaseState = original.LeaseState

const (
	LeaseStateAvailable LeaseState = original.LeaseStateAvailable
	LeaseStateBreaking  LeaseState = original.LeaseStateBreaking
	LeaseStateBroken    LeaseState = original.LeaseStateBroken
	LeaseStateExpired   LeaseState = original.LeaseStateExpired
	LeaseStateLeased    LeaseState = original.LeaseStateLeased
)

type LeaseStatus = original.LeaseStatus

const (
	LeaseStatusLocked   LeaseStatus = original.LeaseStatusLocked
	LeaseStatusUnlocked LeaseStatus = original.LeaseStatusUnlocked
)

type ListContainersInclude = original.ListContainersInclude

const (
	Deleted ListContainersInclude = original.Deleted
)

type ListKeyExpand = original.ListKeyExpand

const (
	Kerb ListKeyExpand = original.Kerb
)

type ListSharesExpand = original.ListSharesExpand

const (
	ListSharesExpandDeleted   ListSharesExpand = original.ListSharesExpandDeleted
	ListSharesExpandSnapshots ListSharesExpand = original.ListSharesExpandSnapshots
)

type MinimumTLSVersion = original.MinimumTLSVersion

const (
	TLS10 MinimumTLSVersion = original.TLS10
	TLS11 MinimumTLSVersion = original.TLS11
	TLS12 MinimumTLSVersion = original.TLS12
)

type Name = original.Name

const (
	AccessTimeTracking Name = original.AccessTimeTracking
)

type Permissions = original.Permissions

const (
	A Permissions = original.A
	C Permissions = original.C
	D Permissions = original.D
	L Permissions = original.L
	P Permissions = original.P
	R Permissions = original.R
	U Permissions = original.U
	W Permissions = original.W
)

type PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningState

const (
	PrivateEndpointConnectionProvisioningStateCreating  PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateCreating
	PrivateEndpointConnectionProvisioningStateDeleting  PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateDeleting
	PrivateEndpointConnectionProvisioningStateFailed    PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateFailed
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateSucceeded
)

type PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatus

const (
	Approved PrivateEndpointServiceConnectionStatus = original.Approved
	Pending  PrivateEndpointServiceConnectionStatus = original.Pending
	Rejected PrivateEndpointServiceConnectionStatus = original.Rejected
)

type ProvisioningState = original.ProvisioningState

const (
	Creating     ProvisioningState = original.Creating
	ResolvingDNS ProvisioningState = original.ResolvingDNS
	Succeeded    ProvisioningState = original.Succeeded
)

type PublicAccess = original.PublicAccess

const (
	PublicAccessBlob      PublicAccess = original.PublicAccessBlob
	PublicAccessContainer PublicAccess = original.PublicAccessContainer
	PublicAccessNone      PublicAccess = original.PublicAccessNone
)

type PutSharesExpand = original.PutSharesExpand

const (
	Snapshots PutSharesExpand = original.Snapshots
)

type Reason = original.Reason

const (
	AccountNameInvalid Reason = original.AccountNameInvalid
	AlreadyExists      Reason = original.AlreadyExists
)

type ReasonCode = original.ReasonCode

const (
	NotAvailableForSubscription ReasonCode = original.NotAvailableForSubscription
	QuotaID                     ReasonCode = original.QuotaID
)

type RootSquashType = original.RootSquashType

const (
	AllSquash    RootSquashType = original.AllSquash
	NoRootSquash RootSquashType = original.NoRootSquash
	RootSquash   RootSquashType = original.RootSquash
)

type RoutingChoice = original.RoutingChoice

const (
	InternetRouting  RoutingChoice = original.InternetRouting
	MicrosoftRouting RoutingChoice = original.MicrosoftRouting
)

type Services = original.Services

const (
	B Services = original.B
	F Services = original.F
	Q Services = original.Q
	T Services = original.T
)

type ShareAccessTier = original.ShareAccessTier

const (
	ShareAccessTierCool                 ShareAccessTier = original.ShareAccessTierCool
	ShareAccessTierHot                  ShareAccessTier = original.ShareAccessTierHot
	ShareAccessTierPremium              ShareAccessTier = original.ShareAccessTierPremium
	ShareAccessTierTransactionOptimized ShareAccessTier = original.ShareAccessTierTransactionOptimized
)

type SignedResource = original.SignedResource

const (
	SignedResourceB SignedResource = original.SignedResourceB
	SignedResourceC SignedResource = original.SignedResourceC
	SignedResourceF SignedResource = original.SignedResourceF
	SignedResourceS SignedResource = original.SignedResourceS
)

type SignedResourceTypes = original.SignedResourceTypes

const (
	SignedResourceTypesC SignedResourceTypes = original.SignedResourceTypesC
	SignedResourceTypesO SignedResourceTypes = original.SignedResourceTypesO
	SignedResourceTypesS SignedResourceTypes = original.SignedResourceTypesS
)

type SkuName = original.SkuName

const (
	PremiumLRS     SkuName = original.PremiumLRS
	PremiumZRS     SkuName = original.PremiumZRS
	StandardGRS    SkuName = original.StandardGRS
	StandardGZRS   SkuName = original.StandardGZRS
	StandardLRS    SkuName = original.StandardLRS
	StandardRAGRS  SkuName = original.StandardRAGRS
	StandardRAGZRS SkuName = original.StandardRAGZRS
	StandardZRS    SkuName = original.StandardZRS
)

type SkuTier = original.SkuTier

const (
	Premium  SkuTier = original.Premium
	Standard SkuTier = original.Standard
)

type State = original.State

const (
	StateDeprovisioning       State = original.StateDeprovisioning
	StateFailed               State = original.StateFailed
	StateNetworkSourceDeleted State = original.StateNetworkSourceDeleted
	StateProvisioning         State = original.StateProvisioning
	StateSucceeded            State = original.StateSucceeded
)

type UsageUnit = original.UsageUnit

const (
	Bytes           UsageUnit = original.Bytes
	BytesPerSecond  UsageUnit = original.BytesPerSecond
	Count           UsageUnit = original.Count
	CountsPerSecond UsageUnit = original.CountsPerSecond
	Percent         UsageUnit = original.Percent
	Seconds         UsageUnit = original.Seconds
)

type Account = original.Account
type AccountCheckNameAvailabilityParameters = original.AccountCheckNameAvailabilityParameters
type AccountCreateParameters = original.AccountCreateParameters
type AccountInternetEndpoints = original.AccountInternetEndpoints
type AccountKey = original.AccountKey
type AccountListKeysResult = original.AccountListKeysResult
type AccountListResult = original.AccountListResult
type AccountListResultIterator = original.AccountListResultIterator
type AccountListResultPage = original.AccountListResultPage
type AccountMicrosoftEndpoints = original.AccountMicrosoftEndpoints
type AccountProperties = original.AccountProperties
type AccountPropertiesCreateParameters = original.AccountPropertiesCreateParameters
type AccountPropertiesUpdateParameters = original.AccountPropertiesUpdateParameters
type AccountRegenerateKeyParameters = original.AccountRegenerateKeyParameters
type AccountSasParameters = original.AccountSasParameters
type AccountUpdateParameters = original.AccountUpdateParameters
type AccountsClient = original.AccountsClient
type AccountsCreateFuture = original.AccountsCreateFuture
type AccountsFailoverFuture = original.AccountsFailoverFuture
type AccountsRestoreBlobRangesFuture = original.AccountsRestoreBlobRangesFuture
type ActiveDirectoryProperties = original.ActiveDirectoryProperties
type AzureEntityResource = original.AzureEntityResource
type AzureFilesIdentityBasedAuthentication = original.AzureFilesIdentityBasedAuthentication
type BaseClient = original.BaseClient
type BlobContainer = original.BlobContainer
type BlobContainersClient = original.BlobContainersClient
type BlobInventoryPoliciesClient = original.BlobInventoryPoliciesClient
type BlobInventoryPolicy = original.BlobInventoryPolicy
type BlobInventoryPolicyDefinition = original.BlobInventoryPolicyDefinition
type BlobInventoryPolicyFilter = original.BlobInventoryPolicyFilter
type BlobInventoryPolicyProperties = original.BlobInventoryPolicyProperties
type BlobInventoryPolicyRule = original.BlobInventoryPolicyRule
type BlobInventoryPolicySchema = original.BlobInventoryPolicySchema
type BlobRestoreParameters = original.BlobRestoreParameters
type BlobRestoreRange = original.BlobRestoreRange
type BlobRestoreStatus = original.BlobRestoreStatus
type BlobServiceItems = original.BlobServiceItems
type BlobServiceProperties = original.BlobServiceProperties
type BlobServicePropertiesProperties = original.BlobServicePropertiesProperties
type BlobServicesClient = original.BlobServicesClient
type ChangeFeed = original.ChangeFeed
type CheckNameAvailabilityResult = original.CheckNameAvailabilityResult
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type ContainerProperties = original.ContainerProperties
type CorsRule = original.CorsRule
type CorsRules = original.CorsRules
type CustomDomain = original.CustomDomain
type DateAfterCreation = original.DateAfterCreation
type DateAfterModification = original.DateAfterModification
type DeleteRetentionPolicy = original.DeleteRetentionPolicy
type DeletedAccount = original.DeletedAccount
type DeletedAccountListResult = original.DeletedAccountListResult
type DeletedAccountListResultIterator = original.DeletedAccountListResultIterator
type DeletedAccountListResultPage = original.DeletedAccountListResultPage
type DeletedAccountProperties = original.DeletedAccountProperties
type DeletedAccountsClient = original.DeletedAccountsClient
type DeletedShare = original.DeletedShare
type Dimension = original.Dimension
type Encryption = original.Encryption
type EncryptionIdentity = original.EncryptionIdentity
type EncryptionScope = original.EncryptionScope
type EncryptionScopeKeyVaultProperties = original.EncryptionScopeKeyVaultProperties
type EncryptionScopeListResult = original.EncryptionScopeListResult
type EncryptionScopeListResultIterator = original.EncryptionScopeListResultIterator
type EncryptionScopeListResultPage = original.EncryptionScopeListResultPage
type EncryptionScopeProperties = original.EncryptionScopeProperties
type EncryptionScopesClient = original.EncryptionScopesClient
type EncryptionService = original.EncryptionService
type EncryptionServices = original.EncryptionServices
type Endpoints = original.Endpoints
type ErrorResponse = original.ErrorResponse
type ErrorResponseBody = original.ErrorResponseBody
type ExtendedLocation = original.ExtendedLocation
type FileServiceItems = original.FileServiceItems
type FileServiceProperties = original.FileServiceProperties
type FileServicePropertiesProperties = original.FileServicePropertiesProperties
type FileServicesClient = original.FileServicesClient
type FileShare = original.FileShare
type FileShareItem = original.FileShareItem
type FileShareItems = original.FileShareItems
type FileShareItemsIterator = original.FileShareItemsIterator
type FileShareItemsPage = original.FileShareItemsPage
type FileShareProperties = original.FileShareProperties
type FileSharesClient = original.FileSharesClient
type GeoReplicationStats = original.GeoReplicationStats
type IPRule = original.IPRule
type Identity = original.Identity
type ImmutabilityPolicy = original.ImmutabilityPolicy
type ImmutabilityPolicyProperties = original.ImmutabilityPolicyProperties
type ImmutabilityPolicyProperty = original.ImmutabilityPolicyProperty
type KeyCreationTime = original.KeyCreationTime
type KeyPolicy = original.KeyPolicy
type KeyVaultProperties = original.KeyVaultProperties
type LastAccessTimeTrackingPolicy = original.LastAccessTimeTrackingPolicy
type LeaseContainerRequest = original.LeaseContainerRequest
type LeaseContainerResponse = original.LeaseContainerResponse
type LegalHold = original.LegalHold
type LegalHoldProperties = original.LegalHoldProperties
type ListAccountSasResponse = original.ListAccountSasResponse
type ListBlobInventoryPolicy = original.ListBlobInventoryPolicy
type ListContainerItem = original.ListContainerItem
type ListContainerItems = original.ListContainerItems
type ListContainerItemsIterator = original.ListContainerItemsIterator
type ListContainerItemsPage = original.ListContainerItemsPage
type ListQueue = original.ListQueue
type ListQueueProperties = original.ListQueueProperties
type ListQueueResource = original.ListQueueResource
type ListQueueResourceIterator = original.ListQueueResourceIterator
type ListQueueResourcePage = original.ListQueueResourcePage
type ListQueueServices = original.ListQueueServices
type ListServiceSasResponse = original.ListServiceSasResponse
type ListTableResource = original.ListTableResource
type ListTableResourceIterator = original.ListTableResourceIterator
type ListTableResourcePage = original.ListTableResourcePage
type ListTableServices = original.ListTableServices
type ManagementPoliciesClient = original.ManagementPoliciesClient
type ManagementPolicy = original.ManagementPolicy
type ManagementPolicyAction = original.ManagementPolicyAction
type ManagementPolicyBaseBlob = original.ManagementPolicyBaseBlob
type ManagementPolicyDefinition = original.ManagementPolicyDefinition
type ManagementPolicyFilter = original.ManagementPolicyFilter
type ManagementPolicyProperties = original.ManagementPolicyProperties
type ManagementPolicyRule = original.ManagementPolicyRule
type ManagementPolicySchema = original.ManagementPolicySchema
type ManagementPolicySnapShot = original.ManagementPolicySnapShot
type ManagementPolicyVersion = original.ManagementPolicyVersion
type MetricSpecification = original.MetricSpecification
type Multichannel = original.Multichannel
type NetworkRuleSet = original.NetworkRuleSet
type ObjectReplicationPolicies = original.ObjectReplicationPolicies
type ObjectReplicationPoliciesClient = original.ObjectReplicationPoliciesClient
type ObjectReplicationPolicy = original.ObjectReplicationPolicy
type ObjectReplicationPolicyFilter = original.ObjectReplicationPolicyFilter
type ObjectReplicationPolicyProperties = original.ObjectReplicationPolicyProperties
type ObjectReplicationPolicyRule = original.ObjectReplicationPolicyRule
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationProperties = original.OperationProperties
type OperationsClient = original.OperationsClient
type PrivateEndpoint = original.PrivateEndpoint
type PrivateEndpointConnection = original.PrivateEndpointConnection
type PrivateEndpointConnectionListResult = original.PrivateEndpointConnectionListResult
type PrivateEndpointConnectionProperties = original.PrivateEndpointConnectionProperties
type PrivateEndpointConnectionsClient = original.PrivateEndpointConnectionsClient
type PrivateLinkResource = original.PrivateLinkResource
type PrivateLinkResourceListResult = original.PrivateLinkResourceListResult
type PrivateLinkResourceProperties = original.PrivateLinkResourceProperties
type PrivateLinkResourcesClient = original.PrivateLinkResourcesClient
type PrivateLinkServiceConnectionState = original.PrivateLinkServiceConnectionState
type ProtocolSettings = original.ProtocolSettings
type ProxyResource = original.ProxyResource
type Queue = original.Queue
type QueueClient = original.QueueClient
type QueueProperties = original.QueueProperties
type QueueServiceProperties = original.QueueServiceProperties
type QueueServicePropertiesProperties = original.QueueServicePropertiesProperties
type QueueServicesClient = original.QueueServicesClient
type Resource = original.Resource
type ResourceAccessRule = original.ResourceAccessRule
type RestorePolicyProperties = original.RestorePolicyProperties
type Restriction = original.Restriction
type RoutingPreference = original.RoutingPreference
type SKUCapability = original.SKUCapability
type SasPolicy = original.SasPolicy
type ServiceSasParameters = original.ServiceSasParameters
type ServiceSpecification = original.ServiceSpecification
type Sku = original.Sku
type SkuInformation = original.SkuInformation
type SkuListResult = original.SkuListResult
type SkusClient = original.SkusClient
type SmbSetting = original.SmbSetting
type SystemData = original.SystemData
type Table = original.Table
type TableClient = original.TableClient
type TableProperties = original.TableProperties
type TableServiceProperties = original.TableServiceProperties
type TableServicePropertiesProperties = original.TableServicePropertiesProperties
type TableServicesClient = original.TableServicesClient
type TagFilter = original.TagFilter
type TagProperty = original.TagProperty
type TrackedResource = original.TrackedResource
type UpdateHistoryProperty = original.UpdateHistoryProperty
type Usage = original.Usage
type UsageListResult = original.UsageListResult
type UsageName = original.UsageName
type UsagesClient = original.UsagesClient
type UserAssignedIdentity = original.UserAssignedIdentity
type VirtualNetworkRule = original.VirtualNetworkRule

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAccountListResultIterator(page AccountListResultPage) AccountListResultIterator {
	return original.NewAccountListResultIterator(page)
}
func NewAccountListResultPage(cur AccountListResult, getNextPage func(context.Context, AccountListResult) (AccountListResult, error)) AccountListResultPage {
	return original.NewAccountListResultPage(cur, getNextPage)
}
func NewAccountsClient(subscriptionID string) AccountsClient {
	return original.NewAccountsClient(subscriptionID)
}
func NewAccountsClientWithBaseURI(baseURI string, subscriptionID string) AccountsClient {
	return original.NewAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewBlobContainersClient(subscriptionID string) BlobContainersClient {
	return original.NewBlobContainersClient(subscriptionID)
}
func NewBlobContainersClientWithBaseURI(baseURI string, subscriptionID string) BlobContainersClient {
	return original.NewBlobContainersClientWithBaseURI(baseURI, subscriptionID)
}
func NewBlobInventoryPoliciesClient(subscriptionID string) BlobInventoryPoliciesClient {
	return original.NewBlobInventoryPoliciesClient(subscriptionID)
}
func NewBlobInventoryPoliciesClientWithBaseURI(baseURI string, subscriptionID string) BlobInventoryPoliciesClient {
	return original.NewBlobInventoryPoliciesClientWithBaseURI(baseURI, subscriptionID)
}
func NewBlobServicesClient(subscriptionID string) BlobServicesClient {
	return original.NewBlobServicesClient(subscriptionID)
}
func NewBlobServicesClientWithBaseURI(baseURI string, subscriptionID string) BlobServicesClient {
	return original.NewBlobServicesClientWithBaseURI(baseURI, subscriptionID)
}
func NewDeletedAccountListResultIterator(page DeletedAccountListResultPage) DeletedAccountListResultIterator {
	return original.NewDeletedAccountListResultIterator(page)
}
func NewDeletedAccountListResultPage(cur DeletedAccountListResult, getNextPage func(context.Context, DeletedAccountListResult) (DeletedAccountListResult, error)) DeletedAccountListResultPage {
	return original.NewDeletedAccountListResultPage(cur, getNextPage)
}
func NewDeletedAccountsClient(subscriptionID string) DeletedAccountsClient {
	return original.NewDeletedAccountsClient(subscriptionID)
}
func NewDeletedAccountsClientWithBaseURI(baseURI string, subscriptionID string) DeletedAccountsClient {
	return original.NewDeletedAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewEncryptionScopeListResultIterator(page EncryptionScopeListResultPage) EncryptionScopeListResultIterator {
	return original.NewEncryptionScopeListResultIterator(page)
}
func NewEncryptionScopeListResultPage(cur EncryptionScopeListResult, getNextPage func(context.Context, EncryptionScopeListResult) (EncryptionScopeListResult, error)) EncryptionScopeListResultPage {
	return original.NewEncryptionScopeListResultPage(cur, getNextPage)
}
func NewEncryptionScopesClient(subscriptionID string) EncryptionScopesClient {
	return original.NewEncryptionScopesClient(subscriptionID)
}
func NewEncryptionScopesClientWithBaseURI(baseURI string, subscriptionID string) EncryptionScopesClient {
	return original.NewEncryptionScopesClientWithBaseURI(baseURI, subscriptionID)
}
func NewFileServicesClient(subscriptionID string) FileServicesClient {
	return original.NewFileServicesClient(subscriptionID)
}
func NewFileServicesClientWithBaseURI(baseURI string, subscriptionID string) FileServicesClient {
	return original.NewFileServicesClientWithBaseURI(baseURI, subscriptionID)
}
func NewFileShareItemsIterator(page FileShareItemsPage) FileShareItemsIterator {
	return original.NewFileShareItemsIterator(page)
}
func NewFileShareItemsPage(cur FileShareItems, getNextPage func(context.Context, FileShareItems) (FileShareItems, error)) FileShareItemsPage {
	return original.NewFileShareItemsPage(cur, getNextPage)
}
func NewFileSharesClient(subscriptionID string) FileSharesClient {
	return original.NewFileSharesClient(subscriptionID)
}
func NewFileSharesClientWithBaseURI(baseURI string, subscriptionID string) FileSharesClient {
	return original.NewFileSharesClientWithBaseURI(baseURI, subscriptionID)
}
func NewListContainerItemsIterator(page ListContainerItemsPage) ListContainerItemsIterator {
	return original.NewListContainerItemsIterator(page)
}
func NewListContainerItemsPage(cur ListContainerItems, getNextPage func(context.Context, ListContainerItems) (ListContainerItems, error)) ListContainerItemsPage {
	return original.NewListContainerItemsPage(cur, getNextPage)
}
func NewListQueueResourceIterator(page ListQueueResourcePage) ListQueueResourceIterator {
	return original.NewListQueueResourceIterator(page)
}
func NewListQueueResourcePage(cur ListQueueResource, getNextPage func(context.Context, ListQueueResource) (ListQueueResource, error)) ListQueueResourcePage {
	return original.NewListQueueResourcePage(cur, getNextPage)
}
func NewListTableResourceIterator(page ListTableResourcePage) ListTableResourceIterator {
	return original.NewListTableResourceIterator(page)
}
func NewListTableResourcePage(cur ListTableResource, getNextPage func(context.Context, ListTableResource) (ListTableResource, error)) ListTableResourcePage {
	return original.NewListTableResourcePage(cur, getNextPage)
}
func NewManagementPoliciesClient(subscriptionID string) ManagementPoliciesClient {
	return original.NewManagementPoliciesClient(subscriptionID)
}
func NewManagementPoliciesClientWithBaseURI(baseURI string, subscriptionID string) ManagementPoliciesClient {
	return original.NewManagementPoliciesClientWithBaseURI(baseURI, subscriptionID)
}
func NewObjectReplicationPoliciesClient(subscriptionID string) ObjectReplicationPoliciesClient {
	return original.NewObjectReplicationPoliciesClient(subscriptionID)
}
func NewObjectReplicationPoliciesClientWithBaseURI(baseURI string, subscriptionID string) ObjectReplicationPoliciesClient {
	return original.NewObjectReplicationPoliciesClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateEndpointConnectionsClient(subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClient(subscriptionID)
}
func NewPrivateEndpointConnectionsClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateLinkResourcesClient(subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClient(subscriptionID)
}
func NewPrivateLinkResourcesClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClientWithBaseURI(baseURI, subscriptionID)
}
func NewQueueClient(subscriptionID string) QueueClient {
	return original.NewQueueClient(subscriptionID)
}
func NewQueueClientWithBaseURI(baseURI string, subscriptionID string) QueueClient {
	return original.NewQueueClientWithBaseURI(baseURI, subscriptionID)
}
func NewQueueServicesClient(subscriptionID string) QueueServicesClient {
	return original.NewQueueServicesClient(subscriptionID)
}
func NewQueueServicesClientWithBaseURI(baseURI string, subscriptionID string) QueueServicesClient {
	return original.NewQueueServicesClientWithBaseURI(baseURI, subscriptionID)
}
func NewSkusClient(subscriptionID string) SkusClient {
	return original.NewSkusClient(subscriptionID)
}
func NewSkusClientWithBaseURI(baseURI string, subscriptionID string) SkusClient {
	return original.NewSkusClientWithBaseURI(baseURI, subscriptionID)
}
func NewTableClient(subscriptionID string) TableClient {
	return original.NewTableClient(subscriptionID)
}
func NewTableClientWithBaseURI(baseURI string, subscriptionID string) TableClient {
	return original.NewTableClientWithBaseURI(baseURI, subscriptionID)
}
func NewTableServicesClient(subscriptionID string) TableServicesClient {
	return original.NewTableServicesClient(subscriptionID)
}
func NewTableServicesClientWithBaseURI(baseURI string, subscriptionID string) TableServicesClient {
	return original.NewTableServicesClientWithBaseURI(baseURI, subscriptionID)
}
func NewUsagesClient(subscriptionID string) UsagesClient {
	return original.NewUsagesClient(subscriptionID)
}
func NewUsagesClientWithBaseURI(baseURI string, subscriptionID string) UsagesClient {
	return original.NewUsagesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAccessTierValues() []AccessTier {
	return original.PossibleAccessTierValues()
}
func PossibleAccountExpandValues() []AccountExpand {
	return original.PossibleAccountExpandValues()
}
func PossibleAccountStatusValues() []AccountStatus {
	return original.PossibleAccountStatusValues()
}
func PossibleAction1Values() []Action1 {
	return original.PossibleAction1Values()
}
func PossibleActionValues() []Action {
	return original.PossibleActionValues()
}
func PossibleBlobRestoreProgressStatusValues() []BlobRestoreProgressStatus {
	return original.PossibleBlobRestoreProgressStatusValues()
}
func PossibleBypassValues() []Bypass {
	return original.PossibleBypassValues()
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleDefaultActionValues() []DefaultAction {
	return original.PossibleDefaultActionValues()
}
func PossibleDirectoryServiceOptionsValues() []DirectoryServiceOptions {
	return original.PossibleDirectoryServiceOptionsValues()
}
func PossibleEnabledProtocolsValues() []EnabledProtocols {
	return original.PossibleEnabledProtocolsValues()
}
func PossibleEncryptionScopeSourceValues() []EncryptionScopeSource {
	return original.PossibleEncryptionScopeSourceValues()
}
func PossibleEncryptionScopeStateValues() []EncryptionScopeState {
	return original.PossibleEncryptionScopeStateValues()
}
func PossibleExtendedLocationTypesValues() []ExtendedLocationTypes {
	return original.PossibleExtendedLocationTypesValues()
}
func PossibleGeoReplicationStatusValues() []GeoReplicationStatus {
	return original.PossibleGeoReplicationStatusValues()
}
func PossibleGetShareExpandValues() []GetShareExpand {
	return original.PossibleGetShareExpandValues()
}
func PossibleHTTPProtocolValues() []HTTPProtocol {
	return original.PossibleHTTPProtocolValues()
}
func PossibleIdentityTypeValues() []IdentityType {
	return original.PossibleIdentityTypeValues()
}
func PossibleImmutabilityPolicyStateValues() []ImmutabilityPolicyState {
	return original.PossibleImmutabilityPolicyStateValues()
}
func PossibleImmutabilityPolicyUpdateTypeValues() []ImmutabilityPolicyUpdateType {
	return original.PossibleImmutabilityPolicyUpdateTypeValues()
}
func PossibleKeyPermissionValues() []KeyPermission {
	return original.PossibleKeyPermissionValues()
}
func PossibleKeySourceValues() []KeySource {
	return original.PossibleKeySourceValues()
}
func PossibleKeyTypeValues() []KeyType {
	return original.PossibleKeyTypeValues()
}
func PossibleKindValues() []Kind {
	return original.PossibleKindValues()
}
func PossibleLargeFileSharesStateValues() []LargeFileSharesState {
	return original.PossibleLargeFileSharesStateValues()
}
func PossibleLeaseDurationValues() []LeaseDuration {
	return original.PossibleLeaseDurationValues()
}
func PossibleLeaseStateValues() []LeaseState {
	return original.PossibleLeaseStateValues()
}
func PossibleLeaseStatusValues() []LeaseStatus {
	return original.PossibleLeaseStatusValues()
}
func PossibleListContainersIncludeValues() []ListContainersInclude {
	return original.PossibleListContainersIncludeValues()
}
func PossibleListKeyExpandValues() []ListKeyExpand {
	return original.PossibleListKeyExpandValues()
}
func PossibleListSharesExpandValues() []ListSharesExpand {
	return original.PossibleListSharesExpandValues()
}
func PossibleMinimumTLSVersionValues() []MinimumTLSVersion {
	return original.PossibleMinimumTLSVersionValues()
}
func PossibleNameValues() []Name {
	return original.PossibleNameValues()
}
func PossiblePermissionsValues() []Permissions {
	return original.PossiblePermissionsValues()
}
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return original.PossiblePrivateEndpointConnectionProvisioningStateValues()
}
func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return original.PossiblePrivateEndpointServiceConnectionStatusValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossiblePublicAccessValues() []PublicAccess {
	return original.PossiblePublicAccessValues()
}
func PossiblePutSharesExpandValues() []PutSharesExpand {
	return original.PossiblePutSharesExpandValues()
}
func PossibleReasonCodeValues() []ReasonCode {
	return original.PossibleReasonCodeValues()
}
func PossibleReasonValues() []Reason {
	return original.PossibleReasonValues()
}
func PossibleRootSquashTypeValues() []RootSquashType {
	return original.PossibleRootSquashTypeValues()
}
func PossibleRoutingChoiceValues() []RoutingChoice {
	return original.PossibleRoutingChoiceValues()
}
func PossibleServicesValues() []Services {
	return original.PossibleServicesValues()
}
func PossibleShareAccessTierValues() []ShareAccessTier {
	return original.PossibleShareAccessTierValues()
}
func PossibleSignedResourceTypesValues() []SignedResourceTypes {
	return original.PossibleSignedResourceTypesValues()
}
func PossibleSignedResourceValues() []SignedResource {
	return original.PossibleSignedResourceValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}
func PossibleStateValues() []State {
	return original.PossibleStateValues()
}
func PossibleUsageUnitValues() []UsageUnit {
	return original.PossibleUsageUnitValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
