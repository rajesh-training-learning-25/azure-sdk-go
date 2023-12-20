//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatalakeanalytics

import "time"

// Account - A Data Lake Analytics account object, containing all information associated with the named Data Lake Analytics
// account.
type Account struct {
	// READ-ONLY; The resource identifier.
	ID *string

	// READ-ONLY; The resource location.
	Location *string

	// READ-ONLY; The resource name.
	Name *string

	// READ-ONLY; The properties defined by Data Lake Analytics all properties are specific to each resource provider.
	Properties *AccountProperties

	// READ-ONLY; The resource tags.
	Tags map[string]*string

	// READ-ONLY; The resource type.
	Type *string
}

// AccountBasic - A Data Lake Analytics account object, containing all information associated with the named Data Lake Analytics
// account.
type AccountBasic struct {
	// READ-ONLY; The resource identifier.
	ID *string

	// READ-ONLY; The resource location.
	Location *string

	// READ-ONLY; The resource name.
	Name *string

	// READ-ONLY; The properties defined by Data Lake Analytics all properties are specific to each resource provider.
	Properties *AccountPropertiesBasic

	// READ-ONLY; The resource tags.
	Tags map[string]*string

	// READ-ONLY; The resource type.
	Type *string
}

// AccountListResult - Data Lake Analytics account list information.
type AccountListResult struct {
	// READ-ONLY; The current number of data lake analytics accounts under this subscription.
	Count *int32

	// READ-ONLY; The link (url) to the next page of results.
	NextLink *string

	// READ-ONLY; The results of the list operation.
	Value []*AccountBasic
}

// AccountProperties - The account specific properties that are associated with an underlying Data Lake Analytics account.
// Returned only when retrieving a specific account.
type AccountProperties struct {
	// The current state of allowing or disallowing IPs originating within Azure through the firewall. If the firewall is disabled,
	// this is not enforced.
	FirewallAllowAzureIPs *FirewallAllowAzureIPsState

	// The current state of the IP address firewall for this account.
	FirewallState *FirewallState

	// The maximum supported degree of parallelism for this account.
	MaxDegreeOfParallelism *int32

	// The maximum supported degree of parallelism per job for this account.
	MaxDegreeOfParallelismPerJob *int32

	// The maximum supported jobs running under the account at the same time.
	MaxJobCount *int32

	// The commitment tier for the next month.
	NewTier *TierType

	// The list of Data Lake Store accounts associated with this account.
	PublicDataLakeStoreAccounts []*DataLakeStoreAccountInformation

	// The number of days that job metadata is retained.
	QueryStoreRetention *int32

	// READ-ONLY; The unique identifier associated with this Data Lake Analytics account.
	AccountID *string

	// READ-ONLY; The list of compute policies associated with this account.
	ComputePolicies []*ComputePolicy

	// READ-ONLY; The account creation time.
	CreationTime *time.Time

	// READ-ONLY; The commitment tier in use for the current month.
	CurrentTier *TierType

	// READ-ONLY; The list of Data Lake Store accounts associated with this account.
	DataLakeStoreAccounts []*DataLakeStoreAccountInformation

	// READ-ONLY; The current state of the DebugDataAccessLevel for this account.
	DebugDataAccessLevel *DebugDataAccessLevel

	// READ-ONLY; The default Data Lake Store account associated with this account.
	DefaultDataLakeStoreAccount *string

	// READ-ONLY; The type of the default Data Lake Store account associated with this account.
	DefaultDataLakeStoreAccountType *string

	// READ-ONLY; The full CName endpoint for this account.
	Endpoint *string

	// READ-ONLY; The list of firewall rules associated with this account.
	FirewallRules []*FirewallRule

	// READ-ONLY; The list of hiveMetastores associated with this account.
	HiveMetastores []*HiveMetastore

	// READ-ONLY; The account last modified time.
	LastModifiedTime *time.Time

	// READ-ONLY; The maximum supported active jobs under the account at the same time.
	MaxActiveJobCountPerUser *int32

	// READ-ONLY; The maximum supported active jobs under the account at the same time.
	MaxJobRunningTimeInMin *int32

	// READ-ONLY; The maximum supported jobs queued under the account at the same time.
	MaxQueuedJobCountPerUser *int32

	// READ-ONLY; The minimum supported priority per job for this account.
	MinPriorityPerJob *int32

	// READ-ONLY; The provisioning status of the Data Lake Analytics account.
	ProvisioningState *DataLakeAnalyticsAccountStatus

	// READ-ONLY; The state of the Data Lake Analytics account.
	State *DataLakeAnalyticsAccountState

	// READ-ONLY; The list of Azure Blob Storage accounts associated with this account.
	StorageAccounts []*StorageAccountInformation

	// READ-ONLY; The system defined maximum supported degree of parallelism for this account, which restricts the maximum value
	// of parallelism the user can set for the account.
	SystemMaxDegreeOfParallelism *int32

	// READ-ONLY; The system defined maximum supported jobs running under the account at the same time, which restricts the maximum
	// number of running jobs the user can set for the account.
	SystemMaxJobCount *int32

	// READ-ONLY; The list of virtualNetwork rules associated with this account.
	VirtualNetworkRules []*VirtualNetworkRule
}

// AccountPropertiesBasic - The basic account specific properties that are associated with an underlying Data Lake Analytics
// account.
type AccountPropertiesBasic struct {
	// READ-ONLY; The unique identifier associated with this Data Lake Analytics account.
	AccountID *string

	// READ-ONLY; The account creation time.
	CreationTime *time.Time

	// READ-ONLY; The full CName endpoint for this account.
	Endpoint *string

	// READ-ONLY; The account last modified time.
	LastModifiedTime *time.Time

	// READ-ONLY; The provisioning status of the Data Lake Analytics account.
	ProvisioningState *DataLakeAnalyticsAccountStatus

	// READ-ONLY; The state of the Data Lake Analytics account.
	State *DataLakeAnalyticsAccountState
}

// AddDataLakeStoreParameters - The parameters used to add a new Data Lake Store account.
type AddDataLakeStoreParameters struct {
	// The Data Lake Store account properties to use when adding a new Data Lake Store account.
	Properties *AddDataLakeStoreProperties
}

// AddDataLakeStoreProperties - The Data Lake Store account properties to use when adding a new Data Lake Store account.
type AddDataLakeStoreProperties struct {
	// The optional suffix for the Data Lake Store account.
	Suffix *string
}

// AddDataLakeStoreWithAccountParameters - The parameters used to add a new Data Lake Store account while creating a new Data
// Lake Analytics account.
type AddDataLakeStoreWithAccountParameters struct {
	// REQUIRED; The unique name of the Data Lake Store account to add.
	Name *string

	// The Data Lake Store account properties to use when adding a new Data Lake Store account.
	Properties *AddDataLakeStoreProperties
}

// AddStorageAccountParameters - The parameters used to add a new Azure Storage account.
type AddStorageAccountParameters struct {
	// REQUIRED; The Azure Storage account properties to use when adding a new Azure Storage account.
	Properties *AddStorageAccountProperties
}

// AddStorageAccountProperties - The Azure Storage account properties to use when adding a new Azure Storage account.
type AddStorageAccountProperties struct {
	// REQUIRED; The access key associated with this Azure Storage account that will be used to connect to it.
	AccessKey *string

	// The optional suffix for the storage account.
	Suffix *string
}

// AddStorageAccountWithAccountParameters - The parameters used to add a new Azure Storage account while creating a new Data
// Lake Analytics account.
type AddStorageAccountWithAccountParameters struct {
	// REQUIRED; The unique name of the Azure Storage account to add.
	Name *string

	// REQUIRED; The Azure Storage account properties to use when adding a new Azure Storage account.
	Properties *AddStorageAccountProperties
}

// CapabilityInformation - Subscription-level properties and limits for Data Lake Analytics.
type CapabilityInformation struct {
	// READ-ONLY; The current number of accounts under this subscription.
	AccountCount *int32

	// READ-ONLY; The maximum supported number of accounts under this subscription.
	MaxAccountCount *int32

	// READ-ONLY; The Boolean value of true or false to indicate the maintenance state.
	MigrationState *bool

	// READ-ONLY; The subscription state.
	State *SubscriptionState

	// READ-ONLY; The subscription credentials that uniquely identifies the subscription.
	SubscriptionID *string
}

// CheckNameAvailabilityParameters - Data Lake Analytics account name availability check parameters.
type CheckNameAvailabilityParameters struct {
	// REQUIRED; The Data Lake Analytics name to check availability for.
	Name *string

	// REQUIRED; The resource type. Note: This should not be set by the user, as the constant value is Microsoft.DataLakeAnalytics/accounts
	Type *CheckNameAvailabilityParametersType
}

// ComputePolicy - Data Lake Analytics compute policy information.
type ComputePolicy struct {
	// READ-ONLY; The resource identifier.
	ID *string

	// READ-ONLY; The resource name.
	Name *string

	// READ-ONLY; The compute policy properties.
	Properties *ComputePolicyProperties

	// READ-ONLY; The resource type.
	Type *string
}

// ComputePolicyListResult - The list of compute policies in the account.
type ComputePolicyListResult struct {
	// READ-ONLY; The link (url) to the next page of results.
	NextLink *string

	// READ-ONLY; The results of the list operation.
	Value []*ComputePolicy
}

// ComputePolicyProperties - The compute policy properties.
type ComputePolicyProperties struct {
	// READ-ONLY; The maximum degree of parallelism per job this user can use to submit jobs.
	MaxDegreeOfParallelismPerJob *int32

	// READ-ONLY; The minimum priority per job this user can use to submit jobs.
	MinPriorityPerJob *int32

	// READ-ONLY; The AAD object identifier for the entity to create a policy for.
	ObjectID *string

	// READ-ONLY; The type of AAD object the object identifier refers to.
	ObjectType *AADObjectType
}

// CreateComputePolicyWithAccountParameters - The parameters used to create a new compute policy while creating a new Data
// Lake Analytics account.
type CreateComputePolicyWithAccountParameters struct {
	// REQUIRED; The unique name of the compute policy to create.
	Name *string

	// REQUIRED; The compute policy properties to use when creating a new compute policy.
	Properties *CreateOrUpdateComputePolicyProperties
}

// CreateDataLakeAnalyticsAccountParameters - The parameters to use for creating a Data Lake Analytics account.
type CreateDataLakeAnalyticsAccountParameters struct {
	// REQUIRED; The resource location.
	Location *string

	// REQUIRED; The Data Lake Analytics account properties to use for creating.
	Properties *CreateDataLakeAnalyticsAccountProperties

	// The resource tags.
	Tags map[string]*string
}

type CreateDataLakeAnalyticsAccountProperties struct {
	// REQUIRED; The list of Data Lake Store accounts associated with this account.
	DataLakeStoreAccounts []*AddDataLakeStoreWithAccountParameters

	// REQUIRED; The default Data Lake Store account associated with this account.
	DefaultDataLakeStoreAccount *string

	// The list of compute policies associated with this account.
	ComputePolicies []*CreateComputePolicyWithAccountParameters

	// The current state of allowing or disallowing IPs originating within Azure through the firewall. If the firewall is disabled,
	// this is not enforced.
	FirewallAllowAzureIPs *FirewallAllowAzureIPsState

	// The list of firewall rules associated with this account.
	FirewallRules []*CreateFirewallRuleWithAccountParameters

	// The current state of the IP address firewall for this account.
	FirewallState *FirewallState

	// The maximum supported degree of parallelism for this account.
	MaxDegreeOfParallelism *int32

	// The maximum supported degree of parallelism per job for this account.
	MaxDegreeOfParallelismPerJob *int32

	// The maximum supported jobs running under the account at the same time.
	MaxJobCount *int32

	// The minimum supported priority per job for this account.
	MinPriorityPerJob *int32

	// The commitment tier for the next month.
	NewTier *TierType

	// The number of days that job metadata is retained.
	QueryStoreRetention *int32

	// The list of Azure Blob Storage accounts associated with this account.
	StorageAccounts []*AddStorageAccountWithAccountParameters
}

// CreateFirewallRuleWithAccountParameters - The parameters used to create a new firewall rule while creating a new Data Lake
// Analytics account.
type CreateFirewallRuleWithAccountParameters struct {
	// REQUIRED; The unique name of the firewall rule to create.
	Name *string

	// REQUIRED; The firewall rule properties to use when creating a new firewall rule.
	Properties *CreateOrUpdateFirewallRuleProperties
}

// CreateOrUpdateComputePolicyParameters - The parameters used to create a new compute policy.
type CreateOrUpdateComputePolicyParameters struct {
	// REQUIRED; The compute policy properties to use when creating a new compute policy.
	Properties *CreateOrUpdateComputePolicyProperties
}

// CreateOrUpdateComputePolicyProperties - The compute policy properties to use when creating a new compute policy.
type CreateOrUpdateComputePolicyProperties struct {
	// REQUIRED; The AAD object identifier for the entity to create a policy for.
	ObjectID *string

	// REQUIRED; The type of AAD object the object identifier refers to.
	ObjectType *AADObjectType

	// The maximum degree of parallelism per job this user can use to submit jobs. This property, the min priority per job property,
	// or both must be passed.
	MaxDegreeOfParallelismPerJob *int32

	// The minimum priority per job this user can use to submit jobs. This property, the max degree of parallelism per job property,
	// or both must be passed.
	MinPriorityPerJob *int32
}

// CreateOrUpdateFirewallRuleParameters - The parameters used to create a new firewall rule.
type CreateOrUpdateFirewallRuleParameters struct {
	// REQUIRED; The firewall rule properties to use when creating a new firewall rule.
	Properties *CreateOrUpdateFirewallRuleProperties
}

// CreateOrUpdateFirewallRuleProperties - The firewall rule properties to use when creating a new firewall rule.
type CreateOrUpdateFirewallRuleProperties struct {
	// REQUIRED; The end IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same
	// protocol.
	EndIPAddress *string

	// REQUIRED; The start IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same
	// protocol.
	StartIPAddress *string
}

// DataLakeStoreAccountInformation - Data Lake Store account information.
type DataLakeStoreAccountInformation struct {
	// The Data Lake Store account properties.
	Properties *DataLakeStoreAccountInformationProperties

	// READ-ONLY; The resource identifier.
	ID *string

	// READ-ONLY; The resource name.
	Name *string

	// READ-ONLY; The resource type.
	Type *string
}

// DataLakeStoreAccountInformationListResult - Data Lake Store account list information.
type DataLakeStoreAccountInformationListResult struct {
	// READ-ONLY; The link (url) to the next page of results.
	NextLink *string

	// READ-ONLY; The results of the list operation.
	Value []*DataLakeStoreAccountInformation
}

// DataLakeStoreAccountInformationProperties - The Data Lake Store account properties.
type DataLakeStoreAccountInformationProperties struct {
	// READ-ONLY; The optional suffix for the Data Lake Store account.
	Suffix *string
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info any

	// READ-ONLY; The additional info type.
	Type *string
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo

	// READ-ONLY; The error code.
	Code *string

	// READ-ONLY; The error details.
	Details []*ErrorDetail

	// READ-ONLY; The error message.
	Message *string

	// READ-ONLY; The error target.
	Target *string
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.).
type ErrorResponse struct {
	// The error object.
	Error *ErrorDetail
}

// FirewallRule - Data Lake Analytics firewall rule information.
type FirewallRule struct {
	// READ-ONLY; The resource identifier.
	ID *string

	// READ-ONLY; The resource name.
	Name *string

	// READ-ONLY; The firewall rule properties.
	Properties *FirewallRuleProperties

	// READ-ONLY; The resource type.
	Type *string
}

// FirewallRuleListResult - Data Lake Analytics firewall rule list information.
type FirewallRuleListResult struct {
	// READ-ONLY; The link (url) to the next page of results.
	NextLink *string

	// READ-ONLY; The results of the list operation.
	Value []*FirewallRule
}

// FirewallRuleProperties - The firewall rule properties.
type FirewallRuleProperties struct {
	// READ-ONLY; The end IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same
	// protocol.
	EndIPAddress *string

	// READ-ONLY; The start IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the
	// same protocol.
	StartIPAddress *string
}

type HiveMetastore struct {
	// READ-ONLY; The resource identifier.
	ID *string

	// READ-ONLY; The resource name.
	Name *string

	// READ-ONLY; The HiveMetastoreProperties rule properties.
	Properties *HiveMetastoreProperties

	// READ-ONLY; The resource type.
	Type *string
}

// HiveMetastoreListResult - Data Lake Analytics HiveMetastore list information.
type HiveMetastoreListResult struct {
	// READ-ONLY; The link (url) to the next page of results.
	NextLink *string

	// READ-ONLY; The results of the list operation.
	Value []*HiveMetastore
}

// HiveMetastoreProperties - The HiveMetastore properties.
type HiveMetastoreProperties struct {
	// READ-ONLY; The databaseName for the Hive MetaStore
	DatabaseName *string

	// READ-ONLY; The current state of the NestedResource
	NestedResourceProvisioningState *NestedResourceProvisioningState

	// READ-ONLY; The password for the Hive MetaStore
	Password *string

	// READ-ONLY; The runtimeVersion for the Hive MetaStore
	RuntimeVersion *string

	// READ-ONLY; The serverUri for the Hive MetaStore
	ServerURI *string

	// READ-ONLY; The userName for the Hive MetaStore
	UserName *string
}

// NameAvailabilityInformation - Data Lake Analytics account name availability result information.
type NameAvailabilityInformation struct {
	// READ-ONLY; The message describing why the Data Lake Analytics account name is not available, if nameAvailable is false.
	Message *string

	// READ-ONLY; The Boolean value of true or false to indicate whether the Data Lake Analytics account name is available or
	// not.
	NameAvailable *bool

	// READ-ONLY; The reason why the Data Lake Analytics account name is not available, if nameAvailable is false.
	Reason *string
}

// Operation - An available operation for Data Lake Analytics.
type Operation struct {
	// READ-ONLY; The display information for the operation.
	Display *OperationDisplay

	// READ-ONLY; The name of the operation.
	Name *string

	// READ-ONLY; The intended executor of the operation.
	Origin *OperationOrigin

	// READ-ONLY; The OperationMetaPropertyInfo for the operation.
	Properties *OperationMetaPropertyInfo
}

// OperationDisplay - The display information for a particular operation.
type OperationDisplay struct {
	// READ-ONLY; A friendly description of the operation.
	Description *string

	// READ-ONLY; A friendly name of the operation.
	Operation *string

	// READ-ONLY; The resource provider of the operation.
	Provider *string

	// READ-ONLY; The resource type of the operation.
	Resource *string
}

// OperationListResult - The list of available operations for Data Lake Analytics.
type OperationListResult struct {
	// READ-ONLY; The link (url) to the next page of results.
	NextLink *string

	// READ-ONLY; The results of the list operation.
	Value []*Operation
}

type OperationMetaLogSpecification struct {
	// The blobDuration for OperationMetaLogSpecification.
	BlobDuration *string

	// The displayName for OperationMetaLogSpecification.
	DisplayName *string

	// The name for OperationMetaLogSpecification.
	Name *string
}

type OperationMetaMetricAvailabilitiesSpecification struct {
	// The blobDuration for OperationMetaMetricAvailabilitiesSpecification.
	BlobDuration *string

	// The timegrain for OperationMetaMetricAvailabilitiesSpecification.
	TimeGrain *string
}

type OperationMetaMetricSpecification struct {
	// The aggregationType for OperationMetaMetricSpecification.
	AggregationType *string

	// The availabilities for OperationMetaMetricSpecification.
	Availabilities []*OperationMetaMetricAvailabilitiesSpecification

	// The displayName for OperationMetaMetricSpecification.
	DisplayDescription *string

	// The displayName for OperationMetaMetricSpecification.
	DisplayName *string

	// The name for OperationMetaMetricSpecification.
	Name *string

	// The unit for OperationMetaMetricSpecification.
	Unit *string
}

type OperationMetaPropertyInfo struct {
	// The operations OperationMetaServiceSpecification.
	ServiceSpecification *OperationMetaServiceSpecification
}

type OperationMetaServiceSpecification struct {
	// The logSpecifications for OperationMetaServiceSpecification.
	LogSpecifications []*OperationMetaLogSpecification

	// The metricSpecifications for OperationMetaServiceSpecification.
	MetricSpecifications []*OperationMetaMetricSpecification
}

// Resource - The resource model definition.
type Resource struct {
	// READ-ONLY; The resource identifier.
	ID *string

	// READ-ONLY; The resource location.
	Location *string

	// READ-ONLY; The resource name.
	Name *string

	// READ-ONLY; The resource tags.
	Tags map[string]*string

	// READ-ONLY; The resource type.
	Type *string
}

// SasTokenInformation - SAS token information.
type SasTokenInformation struct {
	// READ-ONLY; The access token for the associated Azure Storage Container.
	AccessToken *string
}

// SasTokenInformationListResult - The SAS response that contains the storage account, container and associated SAS token
// for connection use.
type SasTokenInformationListResult struct {
	// READ-ONLY; The link (url) to the next page of results.
	NextLink *string

	// READ-ONLY; The results of the list operation.
	Value []*SasTokenInformation
}

// StorageAccountInformation - Azure Storage account information.
type StorageAccountInformation struct {
	// READ-ONLY; The resource identifier.
	ID *string

	// READ-ONLY; The resource name.
	Name *string

	// READ-ONLY; The Azure Storage account properties.
	Properties *StorageAccountInformationProperties

	// READ-ONLY; The resource type.
	Type *string
}

// StorageAccountInformationListResult - Azure Storage account list information.
type StorageAccountInformationListResult struct {
	// READ-ONLY; The link (url) to the next page of results.
	NextLink *string

	// READ-ONLY; The results of the list operation.
	Value []*StorageAccountInformation
}

// StorageAccountInformationProperties - The Azure Storage account properties.
type StorageAccountInformationProperties struct {
	// READ-ONLY; The optional suffix for the storage account.
	Suffix *string
}

// StorageContainer - Azure Storage blob container information.
type StorageContainer struct {
	// READ-ONLY; The resource identifier.
	ID *string

	// READ-ONLY; The resource name.
	Name *string

	// READ-ONLY; The properties of the blob container.
	Properties *StorageContainerProperties

	// READ-ONLY; The resource type.
	Type *string
}

// StorageContainerListResult - The list of blob containers associated with the storage account attached to the Data Lake
// Analytics account.
type StorageContainerListResult struct {
	// READ-ONLY; The link (url) to the next page of results.
	NextLink *string

	// READ-ONLY; The results of the list operation.
	Value []*StorageContainer
}

// StorageContainerProperties - Azure Storage blob container properties information.
type StorageContainerProperties struct {
	// READ-ONLY; The last modified time of the blob container.
	LastModifiedTime *time.Time
}

// SubResource - The resource model definition for a nested resource.
type SubResource struct {
	// READ-ONLY; The resource identifier.
	ID *string

	// READ-ONLY; The resource name.
	Name *string

	// READ-ONLY; The resource type.
	Type *string
}

// UpdateComputePolicyParameters - The parameters used to update a compute policy.
type UpdateComputePolicyParameters struct {
	// The compute policy properties to use when updating a compute policy.
	Properties *UpdateComputePolicyProperties
}

// UpdateComputePolicyProperties - The compute policy properties to use when updating a compute policy.
type UpdateComputePolicyProperties struct {
	// The maximum degree of parallelism per job this user can use to submit jobs. This property, the min priority per job property,
	// or both must be passed.
	MaxDegreeOfParallelismPerJob *int32

	// The minimum priority per job this user can use to submit jobs. This property, the max degree of parallelism per job property,
	// or both must be passed.
	MinPriorityPerJob *int32

	// The AAD object identifier for the entity to create a policy for.
	ObjectID *string

	// The type of AAD object the object identifier refers to.
	ObjectType *AADObjectType
}

// UpdateComputePolicyWithAccountParameters - The parameters used to update a compute policy while updating a Data Lake Analytics
// account.
type UpdateComputePolicyWithAccountParameters struct {
	// REQUIRED; The unique name of the compute policy to update.
	Name *string

	// The compute policy properties to use when updating a compute policy.
	Properties *UpdateComputePolicyProperties
}

// UpdateDataLakeAnalyticsAccountParameters - The parameters that can be used to update an existing Data Lake Analytics account.
type UpdateDataLakeAnalyticsAccountParameters struct {
	// The properties that can be updated in an existing Data Lake Analytics account.
	Properties *UpdateDataLakeAnalyticsAccountProperties

	// The resource tags.
	Tags map[string]*string
}

// UpdateDataLakeAnalyticsAccountProperties - The properties to update that are associated with an underlying Data Lake Analytics
// account.
type UpdateDataLakeAnalyticsAccountProperties struct {
	// The list of compute policies associated with this account.
	ComputePolicies []*UpdateComputePolicyWithAccountParameters

	// The list of Data Lake Store accounts associated with this account.
	DataLakeStoreAccounts []*UpdateDataLakeStoreWithAccountParameters

	// The current state of allowing or disallowing IPs originating within Azure through the firewall. If the firewall is disabled,
	// this is not enforced.
	FirewallAllowAzureIPs *FirewallAllowAzureIPsState

	// The list of firewall rules associated with this account.
	FirewallRules []*UpdateFirewallRuleWithAccountParameters

	// The current state of the IP address firewall for this account. Disabling the firewall does not remove existing rules, they
	// will just be ignored until the firewall is re-enabled.
	FirewallState *FirewallState

	// The maximum supported degree of parallelism for this account.
	MaxDegreeOfParallelism *int32

	// The maximum supported degree of parallelism per job for this account.
	MaxDegreeOfParallelismPerJob *int32

	// The maximum supported jobs running under the account at the same time.
	MaxJobCount *int32

	// The minimum supported priority per job for this account.
	MinPriorityPerJob *int32

	// The commitment tier to use for next month.
	NewTier *TierType

	// The number of days that job metadata is retained.
	QueryStoreRetention *int32

	// The list of Azure Blob storage accounts associated with this account.
	StorageAccounts []*UpdateStorageAccountWithAccountParameters
}

// UpdateDataLakeStoreProperties - The Data Lake Store account properties to use when updating a Data Lake Store account.
type UpdateDataLakeStoreProperties struct {
	// The optional suffix for the Data Lake Store account.
	Suffix *string
}

// UpdateDataLakeStoreWithAccountParameters - The parameters used to update a Data Lake Store account while updating a Data
// Lake Analytics account.
type UpdateDataLakeStoreWithAccountParameters struct {
	// REQUIRED; The unique name of the Data Lake Store account to update.
	Name *string

	// The Data Lake Store account properties to use when updating a Data Lake Store account.
	Properties *UpdateDataLakeStoreProperties
}

// UpdateFirewallRuleParameters - The parameters used to update a firewall rule.
type UpdateFirewallRuleParameters struct {
	// The firewall rule properties to use when updating a firewall rule.
	Properties *UpdateFirewallRuleProperties
}

// UpdateFirewallRuleProperties - The firewall rule properties to use when updating a firewall rule.
type UpdateFirewallRuleProperties struct {
	// The end IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same protocol.
	EndIPAddress *string

	// The start IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same protocol.
	StartIPAddress *string
}

// UpdateFirewallRuleWithAccountParameters - The parameters used to update a firewall rule while updating a Data Lake Analytics
// account.
type UpdateFirewallRuleWithAccountParameters struct {
	// REQUIRED; The unique name of the firewall rule to update.
	Name *string

	// The firewall rule properties to use when updating a firewall rule.
	Properties *UpdateFirewallRuleProperties
}

// UpdateStorageAccountParameters - The parameters used to update an Azure Storage account.
type UpdateStorageAccountParameters struct {
	// The Azure Storage account properties to use when updating an Azure Storage account.
	Properties *UpdateStorageAccountProperties
}

// UpdateStorageAccountProperties - The Azure Storage account properties to use when updating an Azure Storage account.
type UpdateStorageAccountProperties struct {
	// The updated access key associated with this Azure Storage account that will be used to connect to it.
	AccessKey *string

	// The optional suffix for the storage account.
	Suffix *string
}

// UpdateStorageAccountWithAccountParameters - The parameters used to update an Azure Storage account while updating a Data
// Lake Analytics account.
type UpdateStorageAccountWithAccountParameters struct {
	// REQUIRED; The unique name of the Azure Storage account to update.
	Name *string

	// The Azure Storage account properties to use when updating an Azure Storage account.
	Properties *UpdateStorageAccountProperties
}

// VirtualNetworkRule - Data Lake Analytics VirtualNetwork Rule information.
type VirtualNetworkRule struct {
	// READ-ONLY; The resource identifier.
	ID *string

	// READ-ONLY; The resource name.
	Name *string

	// READ-ONLY; The VirtualNetwork rule properties.
	Properties *VirtualNetworkRuleProperties

	// READ-ONLY; The resource type.
	Type *string
}

// VirtualNetworkRuleListResult - Data Lake Analytics VirtualNetwork rule list information.
type VirtualNetworkRuleListResult struct {
	// READ-ONLY; The link (url) to the next page of results.
	NextLink *string

	// READ-ONLY; The results of the list operation.
	Value []*VirtualNetworkRule
}

// VirtualNetworkRuleProperties - The VirtualNetwork Rule properties.
type VirtualNetworkRuleProperties struct {
	// READ-ONLY; The resource identifier for the subnet
	SubnetID *string

	// READ-ONLY; The current state of the VirtualNetwork Rule
	VirtualNetworkRuleState *VirtualNetworkRuleState
}
