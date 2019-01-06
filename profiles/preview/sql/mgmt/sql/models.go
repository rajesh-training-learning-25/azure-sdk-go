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

package sql

import original "github.com/Azure/azure-sdk-for-go/services/sql/mgmt/2014-04-01/sql"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AuthenticationType = original.AuthenticationType

const (
	ADPassword AuthenticationType = original.ADPassword
	SQL        AuthenticationType = original.SQL
)

type CheckNameAvailabilityReason = original.CheckNameAvailabilityReason

const (
	AlreadyExists CheckNameAvailabilityReason = original.AlreadyExists
	Invalid       CheckNameAvailabilityReason = original.Invalid
)

type CreateMode = original.CreateMode

const (
	Copy                           CreateMode = original.Copy
	Default                        CreateMode = original.Default
	NonReadableSecondary           CreateMode = original.NonReadableSecondary
	OnlineSecondary                CreateMode = original.OnlineSecondary
	PointInTimeRestore             CreateMode = original.PointInTimeRestore
	Recovery                       CreateMode = original.Recovery
	Restore                        CreateMode = original.Restore
	RestoreLongTermRetentionBackup CreateMode = original.RestoreLongTermRetentionBackup
)

type DatabaseEdition = original.DatabaseEdition

const (
	Basic         DatabaseEdition = original.Basic
	Business      DatabaseEdition = original.Business
	DataWarehouse DatabaseEdition = original.DataWarehouse
	Free          DatabaseEdition = original.Free
	Premium       DatabaseEdition = original.Premium
	PremiumRS     DatabaseEdition = original.PremiumRS
	Standard      DatabaseEdition = original.Standard
	Stretch       DatabaseEdition = original.Stretch
	System        DatabaseEdition = original.System
	System2       DatabaseEdition = original.System2
	Web           DatabaseEdition = original.Web
)

type ElasticPoolEdition = original.ElasticPoolEdition

const (
	ElasticPoolEditionBasic    ElasticPoolEdition = original.ElasticPoolEditionBasic
	ElasticPoolEditionPremium  ElasticPoolEdition = original.ElasticPoolEditionPremium
	ElasticPoolEditionStandard ElasticPoolEdition = original.ElasticPoolEditionStandard
)

type ElasticPoolState = original.ElasticPoolState

const (
	Creating ElasticPoolState = original.Creating
	Disabled ElasticPoolState = original.Disabled
	Ready    ElasticPoolState = original.Ready
)

type ReadScale = original.ReadScale

const (
	ReadScaleDisabled ReadScale = original.ReadScaleDisabled
	ReadScaleEnabled  ReadScale = original.ReadScaleEnabled
)

type RecommendedIndexAction = original.RecommendedIndexAction

const (
	Create  RecommendedIndexAction = original.Create
	Drop    RecommendedIndexAction = original.Drop
	Rebuild RecommendedIndexAction = original.Rebuild
)

type RecommendedIndexState = original.RecommendedIndexState

const (
	Active        RecommendedIndexState = original.Active
	Blocked       RecommendedIndexState = original.Blocked
	Executing     RecommendedIndexState = original.Executing
	Expired       RecommendedIndexState = original.Expired
	Ignored       RecommendedIndexState = original.Ignored
	Pending       RecommendedIndexState = original.Pending
	PendingRevert RecommendedIndexState = original.PendingRevert
	Reverted      RecommendedIndexState = original.Reverted
	Reverting     RecommendedIndexState = original.Reverting
	Success       RecommendedIndexState = original.Success
	Verifying     RecommendedIndexState = original.Verifying
)

type RecommendedIndexType = original.RecommendedIndexType

const (
	CLUSTERED            RecommendedIndexType = original.CLUSTERED
	CLUSTEREDCOLUMNSTORE RecommendedIndexType = original.CLUSTEREDCOLUMNSTORE
	COLUMNSTORE          RecommendedIndexType = original.COLUMNSTORE
	NONCLUSTERED         RecommendedIndexType = original.NONCLUSTERED
)

type ReplicationRole = original.ReplicationRole

const (
	ReplicationRoleCopy                 ReplicationRole = original.ReplicationRoleCopy
	ReplicationRoleNonReadableSecondary ReplicationRole = original.ReplicationRoleNonReadableSecondary
	ReplicationRolePrimary              ReplicationRole = original.ReplicationRolePrimary
	ReplicationRoleSecondary            ReplicationRole = original.ReplicationRoleSecondary
	ReplicationRoleSource               ReplicationRole = original.ReplicationRoleSource
)

type ReplicationState = original.ReplicationState

const (
	CATCHUP   ReplicationState = original.CATCHUP
	PENDING   ReplicationState = original.PENDING
	SEEDING   ReplicationState = original.SEEDING
	SUSPENDED ReplicationState = original.SUSPENDED
)

type SampleName = original.SampleName

const (
	AdventureWorksLT SampleName = original.AdventureWorksLT
)

type SecurityAlertPolicyEmailAccountAdmins = original.SecurityAlertPolicyEmailAccountAdmins

const (
	SecurityAlertPolicyEmailAccountAdminsDisabled SecurityAlertPolicyEmailAccountAdmins = original.SecurityAlertPolicyEmailAccountAdminsDisabled
	SecurityAlertPolicyEmailAccountAdminsEnabled  SecurityAlertPolicyEmailAccountAdmins = original.SecurityAlertPolicyEmailAccountAdminsEnabled
)

type SecurityAlertPolicyState = original.SecurityAlertPolicyState

const (
	SecurityAlertPolicyStateDisabled SecurityAlertPolicyState = original.SecurityAlertPolicyStateDisabled
	SecurityAlertPolicyStateEnabled  SecurityAlertPolicyState = original.SecurityAlertPolicyStateEnabled
	SecurityAlertPolicyStateNew      SecurityAlertPolicyState = original.SecurityAlertPolicyStateNew
)

type SecurityAlertPolicyUseServerDefault = original.SecurityAlertPolicyUseServerDefault

const (
	SecurityAlertPolicyUseServerDefaultDisabled SecurityAlertPolicyUseServerDefault = original.SecurityAlertPolicyUseServerDefaultDisabled
	SecurityAlertPolicyUseServerDefaultEnabled  SecurityAlertPolicyUseServerDefault = original.SecurityAlertPolicyUseServerDefaultEnabled
)

type ServiceObjectiveName = original.ServiceObjectiveName

const (
	ServiceObjectiveNameBasic       ServiceObjectiveName = original.ServiceObjectiveNameBasic
	ServiceObjectiveNameDS100       ServiceObjectiveName = original.ServiceObjectiveNameDS100
	ServiceObjectiveNameDS1000      ServiceObjectiveName = original.ServiceObjectiveNameDS1000
	ServiceObjectiveNameDS1200      ServiceObjectiveName = original.ServiceObjectiveNameDS1200
	ServiceObjectiveNameDS1500      ServiceObjectiveName = original.ServiceObjectiveNameDS1500
	ServiceObjectiveNameDS200       ServiceObjectiveName = original.ServiceObjectiveNameDS200
	ServiceObjectiveNameDS2000      ServiceObjectiveName = original.ServiceObjectiveNameDS2000
	ServiceObjectiveNameDS300       ServiceObjectiveName = original.ServiceObjectiveNameDS300
	ServiceObjectiveNameDS400       ServiceObjectiveName = original.ServiceObjectiveNameDS400
	ServiceObjectiveNameDS500       ServiceObjectiveName = original.ServiceObjectiveNameDS500
	ServiceObjectiveNameDS600       ServiceObjectiveName = original.ServiceObjectiveNameDS600
	ServiceObjectiveNameDW100       ServiceObjectiveName = original.ServiceObjectiveNameDW100
	ServiceObjectiveNameDW1000      ServiceObjectiveName = original.ServiceObjectiveNameDW1000
	ServiceObjectiveNameDW10000c    ServiceObjectiveName = original.ServiceObjectiveNameDW10000c
	ServiceObjectiveNameDW1000c     ServiceObjectiveName = original.ServiceObjectiveNameDW1000c
	ServiceObjectiveNameDW1200      ServiceObjectiveName = original.ServiceObjectiveNameDW1200
	ServiceObjectiveNameDW1500      ServiceObjectiveName = original.ServiceObjectiveNameDW1500
	ServiceObjectiveNameDW15000c    ServiceObjectiveName = original.ServiceObjectiveNameDW15000c
	ServiceObjectiveNameDW1500c     ServiceObjectiveName = original.ServiceObjectiveNameDW1500c
	ServiceObjectiveNameDW200       ServiceObjectiveName = original.ServiceObjectiveNameDW200
	ServiceObjectiveNameDW2000      ServiceObjectiveName = original.ServiceObjectiveNameDW2000
	ServiceObjectiveNameDW2000c     ServiceObjectiveName = original.ServiceObjectiveNameDW2000c
	ServiceObjectiveNameDW2500c     ServiceObjectiveName = original.ServiceObjectiveNameDW2500c
	ServiceObjectiveNameDW300       ServiceObjectiveName = original.ServiceObjectiveNameDW300
	ServiceObjectiveNameDW3000      ServiceObjectiveName = original.ServiceObjectiveNameDW3000
	ServiceObjectiveNameDW30000c    ServiceObjectiveName = original.ServiceObjectiveNameDW30000c
	ServiceObjectiveNameDW3000c     ServiceObjectiveName = original.ServiceObjectiveNameDW3000c
	ServiceObjectiveNameDW400       ServiceObjectiveName = original.ServiceObjectiveNameDW400
	ServiceObjectiveNameDW500       ServiceObjectiveName = original.ServiceObjectiveNameDW500
	ServiceObjectiveNameDW5000c     ServiceObjectiveName = original.ServiceObjectiveNameDW5000c
	ServiceObjectiveNameDW600       ServiceObjectiveName = original.ServiceObjectiveNameDW600
	ServiceObjectiveNameDW6000      ServiceObjectiveName = original.ServiceObjectiveNameDW6000
	ServiceObjectiveNameDW6000c     ServiceObjectiveName = original.ServiceObjectiveNameDW6000c
	ServiceObjectiveNameDW7500c     ServiceObjectiveName = original.ServiceObjectiveNameDW7500c
	ServiceObjectiveNameElasticPool ServiceObjectiveName = original.ServiceObjectiveNameElasticPool
	ServiceObjectiveNameFree        ServiceObjectiveName = original.ServiceObjectiveNameFree
	ServiceObjectiveNameP1          ServiceObjectiveName = original.ServiceObjectiveNameP1
	ServiceObjectiveNameP11         ServiceObjectiveName = original.ServiceObjectiveNameP11
	ServiceObjectiveNameP15         ServiceObjectiveName = original.ServiceObjectiveNameP15
	ServiceObjectiveNameP2          ServiceObjectiveName = original.ServiceObjectiveNameP2
	ServiceObjectiveNameP3          ServiceObjectiveName = original.ServiceObjectiveNameP3
	ServiceObjectiveNameP4          ServiceObjectiveName = original.ServiceObjectiveNameP4
	ServiceObjectiveNameP6          ServiceObjectiveName = original.ServiceObjectiveNameP6
	ServiceObjectiveNamePRS1        ServiceObjectiveName = original.ServiceObjectiveNamePRS1
	ServiceObjectiveNamePRS2        ServiceObjectiveName = original.ServiceObjectiveNamePRS2
	ServiceObjectiveNamePRS4        ServiceObjectiveName = original.ServiceObjectiveNamePRS4
	ServiceObjectiveNamePRS6        ServiceObjectiveName = original.ServiceObjectiveNamePRS6
	ServiceObjectiveNameS0          ServiceObjectiveName = original.ServiceObjectiveNameS0
	ServiceObjectiveNameS1          ServiceObjectiveName = original.ServiceObjectiveNameS1
	ServiceObjectiveNameS12         ServiceObjectiveName = original.ServiceObjectiveNameS12
	ServiceObjectiveNameS2          ServiceObjectiveName = original.ServiceObjectiveNameS2
	ServiceObjectiveNameS3          ServiceObjectiveName = original.ServiceObjectiveNameS3
	ServiceObjectiveNameS4          ServiceObjectiveName = original.ServiceObjectiveNameS4
	ServiceObjectiveNameS6          ServiceObjectiveName = original.ServiceObjectiveNameS6
	ServiceObjectiveNameS7          ServiceObjectiveName = original.ServiceObjectiveNameS7
	ServiceObjectiveNameS9          ServiceObjectiveName = original.ServiceObjectiveNameS9
	ServiceObjectiveNameSystem      ServiceObjectiveName = original.ServiceObjectiveNameSystem
	ServiceObjectiveNameSystem0     ServiceObjectiveName = original.ServiceObjectiveNameSystem0
	ServiceObjectiveNameSystem1     ServiceObjectiveName = original.ServiceObjectiveNameSystem1
	ServiceObjectiveNameSystem2     ServiceObjectiveName = original.ServiceObjectiveNameSystem2
	ServiceObjectiveNameSystem2L    ServiceObjectiveName = original.ServiceObjectiveNameSystem2L
	ServiceObjectiveNameSystem3     ServiceObjectiveName = original.ServiceObjectiveNameSystem3
	ServiceObjectiveNameSystem3L    ServiceObjectiveName = original.ServiceObjectiveNameSystem3L
	ServiceObjectiveNameSystem4     ServiceObjectiveName = original.ServiceObjectiveNameSystem4
	ServiceObjectiveNameSystem4L    ServiceObjectiveName = original.ServiceObjectiveNameSystem4L
)

type StorageKeyType = original.StorageKeyType

const (
	SharedAccessKey  StorageKeyType = original.SharedAccessKey
	StorageAccessKey StorageKeyType = original.StorageAccessKey
)

type TransparentDataEncryptionActivityStatus = original.TransparentDataEncryptionActivityStatus

const (
	Decrypting TransparentDataEncryptionActivityStatus = original.Decrypting
	Encrypting TransparentDataEncryptionActivityStatus = original.Encrypting
)

type TransparentDataEncryptionStatus = original.TransparentDataEncryptionStatus

const (
	TransparentDataEncryptionStatusDisabled TransparentDataEncryptionStatus = original.TransparentDataEncryptionStatusDisabled
	TransparentDataEncryptionStatusEnabled  TransparentDataEncryptionStatus = original.TransparentDataEncryptionStatusEnabled
)

type BaseClient = original.BaseClient
type CheckNameAvailabilityRequest = original.CheckNameAvailabilityRequest
type CheckNameAvailabilityResponse = original.CheckNameAvailabilityResponse
type Database = original.Database
type DatabaseListResult = original.DatabaseListResult
type DatabaseProperties = original.DatabaseProperties
type DatabaseSecurityAlertPolicy = original.DatabaseSecurityAlertPolicy
type DatabaseSecurityAlertPolicyProperties = original.DatabaseSecurityAlertPolicyProperties
type DatabaseThreatDetectionPoliciesClient = original.DatabaseThreatDetectionPoliciesClient
type DatabaseUpdate = original.DatabaseUpdate
type DatabasesClient = original.DatabasesClient
type DatabasesCreateImportOperationFuture = original.DatabasesCreateImportOperationFuture
type DatabasesCreateOrUpdateFuture = original.DatabasesCreateOrUpdateFuture
type DatabasesExportFuture = original.DatabasesExportFuture
type DatabasesImportFuture = original.DatabasesImportFuture
type DatabasesPauseFuture = original.DatabasesPauseFuture
type DatabasesResumeFuture = original.DatabasesResumeFuture
type DatabasesUpdateFuture = original.DatabasesUpdateFuture
type ElasticPool = original.ElasticPool
type ElasticPoolActivitiesClient = original.ElasticPoolActivitiesClient
type ElasticPoolActivity = original.ElasticPoolActivity
type ElasticPoolActivityListResult = original.ElasticPoolActivityListResult
type ElasticPoolActivityProperties = original.ElasticPoolActivityProperties
type ElasticPoolDatabaseActivitiesClient = original.ElasticPoolDatabaseActivitiesClient
type ElasticPoolDatabaseActivity = original.ElasticPoolDatabaseActivity
type ElasticPoolDatabaseActivityListResult = original.ElasticPoolDatabaseActivityListResult
type ElasticPoolDatabaseActivityProperties = original.ElasticPoolDatabaseActivityProperties
type ElasticPoolListResult = original.ElasticPoolListResult
type ElasticPoolProperties = original.ElasticPoolProperties
type ElasticPoolUpdate = original.ElasticPoolUpdate
type ElasticPoolsClient = original.ElasticPoolsClient
type ElasticPoolsCreateOrUpdateFuture = original.ElasticPoolsCreateOrUpdateFuture
type ElasticPoolsUpdateFuture = original.ElasticPoolsUpdateFuture
type ExportRequest = original.ExportRequest
type FirewallRule = original.FirewallRule
type FirewallRuleListResult = original.FirewallRuleListResult
type FirewallRuleProperties = original.FirewallRuleProperties
type FirewallRulesClient = original.FirewallRulesClient
type ImportExportResponse = original.ImportExportResponse
type ImportExportResponseProperties = original.ImportExportResponseProperties
type ImportExtensionProperties = original.ImportExtensionProperties
type ImportExtensionRequest = original.ImportExtensionRequest
type ImportRequest = original.ImportRequest
type OperationImpact = original.OperationImpact
type ProxyResource = original.ProxyResource
type RecommendedElasticPool = original.RecommendedElasticPool
type RecommendedElasticPoolListMetricsResult = original.RecommendedElasticPoolListMetricsResult
type RecommendedElasticPoolListResult = original.RecommendedElasticPoolListResult
type RecommendedElasticPoolMetric = original.RecommendedElasticPoolMetric
type RecommendedElasticPoolProperties = original.RecommendedElasticPoolProperties
type RecommendedElasticPoolsClient = original.RecommendedElasticPoolsClient
type RecommendedIndex = original.RecommendedIndex
type RecommendedIndexProperties = original.RecommendedIndexProperties
type ReplicationLink = original.ReplicationLink
type ReplicationLinkListResult = original.ReplicationLinkListResult
type ReplicationLinkProperties = original.ReplicationLinkProperties
type ReplicationLinksClient = original.ReplicationLinksClient
type ReplicationLinksFailoverAllowDataLossFuture = original.ReplicationLinksFailoverAllowDataLossFuture
type ReplicationLinksFailoverFuture = original.ReplicationLinksFailoverFuture
type Resource = original.Resource
type ServersClient = original.ServersClient
type ServiceTierAdvisor = original.ServiceTierAdvisor
type ServiceTierAdvisorListResult = original.ServiceTierAdvisorListResult
type ServiceTierAdvisorProperties = original.ServiceTierAdvisorProperties
type ServiceTierAdvisorsClient = original.ServiceTierAdvisorsClient
type SloUsageMetric = original.SloUsageMetric
type TrackedResource = original.TrackedResource
type TransparentDataEncryption = original.TransparentDataEncryption
type TransparentDataEncryptionActivitiesClient = original.TransparentDataEncryptionActivitiesClient
type TransparentDataEncryptionActivity = original.TransparentDataEncryptionActivity
type TransparentDataEncryptionActivityListResult = original.TransparentDataEncryptionActivityListResult
type TransparentDataEncryptionActivityProperties = original.TransparentDataEncryptionActivityProperties
type TransparentDataEncryptionProperties = original.TransparentDataEncryptionProperties
type TransparentDataEncryptionsClient = original.TransparentDataEncryptionsClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewDatabaseThreatDetectionPoliciesClient(subscriptionID string) DatabaseThreatDetectionPoliciesClient {
	return original.NewDatabaseThreatDetectionPoliciesClient(subscriptionID)
}
func NewDatabaseThreatDetectionPoliciesClientWithBaseURI(baseURI string, subscriptionID string) DatabaseThreatDetectionPoliciesClient {
	return original.NewDatabaseThreatDetectionPoliciesClientWithBaseURI(baseURI, subscriptionID)
}
func NewDatabasesClient(subscriptionID string) DatabasesClient {
	return original.NewDatabasesClient(subscriptionID)
}
func NewDatabasesClientWithBaseURI(baseURI string, subscriptionID string) DatabasesClient {
	return original.NewDatabasesClientWithBaseURI(baseURI, subscriptionID)
}
func NewElasticPoolActivitiesClient(subscriptionID string) ElasticPoolActivitiesClient {
	return original.NewElasticPoolActivitiesClient(subscriptionID)
}
func NewElasticPoolActivitiesClientWithBaseURI(baseURI string, subscriptionID string) ElasticPoolActivitiesClient {
	return original.NewElasticPoolActivitiesClientWithBaseURI(baseURI, subscriptionID)
}
func NewElasticPoolDatabaseActivitiesClient(subscriptionID string) ElasticPoolDatabaseActivitiesClient {
	return original.NewElasticPoolDatabaseActivitiesClient(subscriptionID)
}
func NewElasticPoolDatabaseActivitiesClientWithBaseURI(baseURI string, subscriptionID string) ElasticPoolDatabaseActivitiesClient {
	return original.NewElasticPoolDatabaseActivitiesClientWithBaseURI(baseURI, subscriptionID)
}
func NewElasticPoolsClient(subscriptionID string) ElasticPoolsClient {
	return original.NewElasticPoolsClient(subscriptionID)
}
func NewElasticPoolsClientWithBaseURI(baseURI string, subscriptionID string) ElasticPoolsClient {
	return original.NewElasticPoolsClientWithBaseURI(baseURI, subscriptionID)
}
func NewFirewallRulesClient(subscriptionID string) FirewallRulesClient {
	return original.NewFirewallRulesClient(subscriptionID)
}
func NewFirewallRulesClientWithBaseURI(baseURI string, subscriptionID string) FirewallRulesClient {
	return original.NewFirewallRulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewRecommendedElasticPoolsClient(subscriptionID string) RecommendedElasticPoolsClient {
	return original.NewRecommendedElasticPoolsClient(subscriptionID)
}
func NewRecommendedElasticPoolsClientWithBaseURI(baseURI string, subscriptionID string) RecommendedElasticPoolsClient {
	return original.NewRecommendedElasticPoolsClientWithBaseURI(baseURI, subscriptionID)
}
func NewReplicationLinksClient(subscriptionID string) ReplicationLinksClient {
	return original.NewReplicationLinksClient(subscriptionID)
}
func NewReplicationLinksClientWithBaseURI(baseURI string, subscriptionID string) ReplicationLinksClient {
	return original.NewReplicationLinksClientWithBaseURI(baseURI, subscriptionID)
}
func NewServersClient(subscriptionID string) ServersClient {
	return original.NewServersClient(subscriptionID)
}
func NewServersClientWithBaseURI(baseURI string, subscriptionID string) ServersClient {
	return original.NewServersClientWithBaseURI(baseURI, subscriptionID)
}
func NewServiceTierAdvisorsClient(subscriptionID string) ServiceTierAdvisorsClient {
	return original.NewServiceTierAdvisorsClient(subscriptionID)
}
func NewServiceTierAdvisorsClientWithBaseURI(baseURI string, subscriptionID string) ServiceTierAdvisorsClient {
	return original.NewServiceTierAdvisorsClientWithBaseURI(baseURI, subscriptionID)
}
func NewTransparentDataEncryptionActivitiesClient(subscriptionID string) TransparentDataEncryptionActivitiesClient {
	return original.NewTransparentDataEncryptionActivitiesClient(subscriptionID)
}
func NewTransparentDataEncryptionActivitiesClientWithBaseURI(baseURI string, subscriptionID string) TransparentDataEncryptionActivitiesClient {
	return original.NewTransparentDataEncryptionActivitiesClientWithBaseURI(baseURI, subscriptionID)
}
func NewTransparentDataEncryptionsClient(subscriptionID string) TransparentDataEncryptionsClient {
	return original.NewTransparentDataEncryptionsClient(subscriptionID)
}
func NewTransparentDataEncryptionsClientWithBaseURI(baseURI string, subscriptionID string) TransparentDataEncryptionsClient {
	return original.NewTransparentDataEncryptionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAuthenticationTypeValues() []AuthenticationType {
	return original.PossibleAuthenticationTypeValues()
}
func PossibleCheckNameAvailabilityReasonValues() []CheckNameAvailabilityReason {
	return original.PossibleCheckNameAvailabilityReasonValues()
}
func PossibleCreateModeValues() []CreateMode {
	return original.PossibleCreateModeValues()
}
func PossibleDatabaseEditionValues() []DatabaseEdition {
	return original.PossibleDatabaseEditionValues()
}
func PossibleElasticPoolEditionValues() []ElasticPoolEdition {
	return original.PossibleElasticPoolEditionValues()
}
func PossibleElasticPoolStateValues() []ElasticPoolState {
	return original.PossibleElasticPoolStateValues()
}
func PossibleReadScaleValues() []ReadScale {
	return original.PossibleReadScaleValues()
}
func PossibleRecommendedIndexActionValues() []RecommendedIndexAction {
	return original.PossibleRecommendedIndexActionValues()
}
func PossibleRecommendedIndexStateValues() []RecommendedIndexState {
	return original.PossibleRecommendedIndexStateValues()
}
func PossibleRecommendedIndexTypeValues() []RecommendedIndexType {
	return original.PossibleRecommendedIndexTypeValues()
}
func PossibleReplicationRoleValues() []ReplicationRole {
	return original.PossibleReplicationRoleValues()
}
func PossibleReplicationStateValues() []ReplicationState {
	return original.PossibleReplicationStateValues()
}
func PossibleSampleNameValues() []SampleName {
	return original.PossibleSampleNameValues()
}
func PossibleSecurityAlertPolicyEmailAccountAdminsValues() []SecurityAlertPolicyEmailAccountAdmins {
	return original.PossibleSecurityAlertPolicyEmailAccountAdminsValues()
}
func PossibleSecurityAlertPolicyStateValues() []SecurityAlertPolicyState {
	return original.PossibleSecurityAlertPolicyStateValues()
}
func PossibleSecurityAlertPolicyUseServerDefaultValues() []SecurityAlertPolicyUseServerDefault {
	return original.PossibleSecurityAlertPolicyUseServerDefaultValues()
}
func PossibleServiceObjectiveNameValues() []ServiceObjectiveName {
	return original.PossibleServiceObjectiveNameValues()
}
func PossibleStorageKeyTypeValues() []StorageKeyType {
	return original.PossibleStorageKeyTypeValues()
}
func PossibleTransparentDataEncryptionActivityStatusValues() []TransparentDataEncryptionActivityStatus {
	return original.PossibleTransparentDataEncryptionActivityStatusValues()
}
func PossibleTransparentDataEncryptionStatusValues() []TransparentDataEncryptionStatus {
	return original.PossibleTransparentDataEncryptionStatusValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
