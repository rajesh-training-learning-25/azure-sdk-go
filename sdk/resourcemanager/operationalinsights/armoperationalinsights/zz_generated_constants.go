//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoperationalinsights

const (
	module  = "armoperationalinsights"
	version = "v0.1.0"
)

// BillingType - Configures whether billing will be only on the cluster or each workspace will be billed by its proportional use. This does not change the
// overall billing, only how it will be distributed. Default
// value is 'Cluster'
type BillingType string

const (
	BillingTypeCluster    BillingType = "Cluster"
	BillingTypeWorkspaces BillingType = "Workspaces"
)

// PossibleBillingTypeValues returns the possible values for the BillingType const type.
func PossibleBillingTypeValues() []BillingType {
	return []BillingType{
		BillingTypeCluster,
		BillingTypeWorkspaces,
	}
}

// ToPtr returns a *BillingType pointing to the current value.
func (c BillingType) ToPtr() *BillingType {
	return &c
}

// Capacity - The capacity value
type Capacity int64

const (
	CapacityFiveHundred  Capacity = 500
	CapacityTenHundred   Capacity = 1000
	CapacityTwoThousand  Capacity = 2000
	CapacityFiveThousand Capacity = 5000
)

// PossibleCapacityValues returns the possible values for the Capacity const type.
func PossibleCapacityValues() []Capacity {
	return []Capacity{
		CapacityFiveHundred,
		CapacityTenHundred,
		CapacityTwoThousand,
		CapacityFiveThousand,
	}
}

// ToPtr returns a *Capacity pointing to the current value.
func (c Capacity) ToPtr() *Capacity {
	return &c
}

// CapacityReservationLevel - The capacity reservation level in GB for this workspace, when CapacityReservation sku is selected.
type CapacityReservationLevel int32

const (
	CapacityReservationLevelOneHundred   CapacityReservationLevel = 100
	CapacityReservationLevelTwoHundred   CapacityReservationLevel = 200
	CapacityReservationLevelThreeHundred CapacityReservationLevel = 300
	CapacityReservationLevelFourHundred  CapacityReservationLevel = 400
	CapacityReservationLevelFiveHundred  CapacityReservationLevel = 500
	CapacityReservationLevelTenHundred   CapacityReservationLevel = 1000
	CapacityReservationLevelTwoThousand  CapacityReservationLevel = 2000
	CapacityReservationLevelFiveThousand CapacityReservationLevel = 5000
)

// PossibleCapacityReservationLevelValues returns the possible values for the CapacityReservationLevel const type.
func PossibleCapacityReservationLevelValues() []CapacityReservationLevel {
	return []CapacityReservationLevel{
		CapacityReservationLevelOneHundred,
		CapacityReservationLevelTwoHundred,
		CapacityReservationLevelThreeHundred,
		CapacityReservationLevelFourHundred,
		CapacityReservationLevelFiveHundred,
		CapacityReservationLevelTenHundred,
		CapacityReservationLevelTwoThousand,
		CapacityReservationLevelFiveThousand,
	}
}

// ToPtr returns a *CapacityReservationLevel pointing to the current value.
func (c CapacityReservationLevel) ToPtr() *CapacityReservationLevel {
	return &c
}

// ClusterEntityStatus - The provisioning state of the cluster.
type ClusterEntityStatus string

const (
	ClusterEntityStatusCanceled            ClusterEntityStatus = "Canceled"
	ClusterEntityStatusCreating            ClusterEntityStatus = "Creating"
	ClusterEntityStatusDeleting            ClusterEntityStatus = "Deleting"
	ClusterEntityStatusFailed              ClusterEntityStatus = "Failed"
	ClusterEntityStatusProvisioningAccount ClusterEntityStatus = "ProvisioningAccount"
	ClusterEntityStatusSucceeded           ClusterEntityStatus = "Succeeded"
	ClusterEntityStatusUpdating            ClusterEntityStatus = "Updating"
)

// PossibleClusterEntityStatusValues returns the possible values for the ClusterEntityStatus const type.
func PossibleClusterEntityStatusValues() []ClusterEntityStatus {
	return []ClusterEntityStatus{
		ClusterEntityStatusCanceled,
		ClusterEntityStatusCreating,
		ClusterEntityStatusDeleting,
		ClusterEntityStatusFailed,
		ClusterEntityStatusProvisioningAccount,
		ClusterEntityStatusSucceeded,
		ClusterEntityStatusUpdating,
	}
}

// ToPtr returns a *ClusterEntityStatus pointing to the current value.
func (c ClusterEntityStatus) ToPtr() *ClusterEntityStatus {
	return &c
}

// ClusterSKUNameEnum - The name of the SKU.
type ClusterSKUNameEnum string

const (
	ClusterSKUNameEnumCapacityReservation ClusterSKUNameEnum = "CapacityReservation"
)

// PossibleClusterSKUNameEnumValues returns the possible values for the ClusterSKUNameEnum const type.
func PossibleClusterSKUNameEnumValues() []ClusterSKUNameEnum {
	return []ClusterSKUNameEnum{
		ClusterSKUNameEnumCapacityReservation,
	}
}

// ToPtr returns a *ClusterSKUNameEnum pointing to the current value.
func (c ClusterSKUNameEnum) ToPtr() *ClusterSKUNameEnum {
	return &c
}

// DataIngestionStatus - The status of data ingestion for this workspace.
type DataIngestionStatus string

const (
	// DataIngestionStatusApproachingQuota - 80% of daily cap quota reached.
	DataIngestionStatusApproachingQuota DataIngestionStatus = "ApproachingQuota"
	// DataIngestionStatusForceOff - Ingestion stopped following service setting change.
	DataIngestionStatusForceOff DataIngestionStatus = "ForceOff"
	// DataIngestionStatusForceOn - Ingestion started following service setting change.
	DataIngestionStatusForceOn DataIngestionStatus = "ForceOn"
	// DataIngestionStatusOverQuota - Reached daily cap quota, ingestion stopped.
	DataIngestionStatusOverQuota DataIngestionStatus = "OverQuota"
	// DataIngestionStatusRespectQuota - Ingestion enabled following daily cap quota reset, or subscription enablement.
	DataIngestionStatusRespectQuota DataIngestionStatus = "RespectQuota"
	// DataIngestionStatusSubscriptionSuspended - Ingestion stopped following suspended subscription.
	DataIngestionStatusSubscriptionSuspended DataIngestionStatus = "SubscriptionSuspended"
)

// PossibleDataIngestionStatusValues returns the possible values for the DataIngestionStatus const type.
func PossibleDataIngestionStatusValues() []DataIngestionStatus {
	return []DataIngestionStatus{
		DataIngestionStatusApproachingQuota,
		DataIngestionStatusForceOff,
		DataIngestionStatusForceOn,
		DataIngestionStatusOverQuota,
		DataIngestionStatusRespectQuota,
		DataIngestionStatusSubscriptionSuspended,
	}
}

// ToPtr returns a *DataIngestionStatus pointing to the current value.
func (c DataIngestionStatus) ToPtr() *DataIngestionStatus {
	return &c
}

// DataSourceKind - The kind of the DataSource.
type DataSourceKind string

const (
	DataSourceKindApplicationInsights                                  DataSourceKind = "ApplicationInsights"
	DataSourceKindAzureActivityLog                                     DataSourceKind = "AzureActivityLog"
	DataSourceKindAzureAuditLog                                        DataSourceKind = "AzureAuditLog"
	DataSourceKindChangeTrackingContentLocation                        DataSourceKind = "ChangeTrackingContentLocation"
	DataSourceKindChangeTrackingCustomPath                             DataSourceKind = "ChangeTrackingCustomPath"
	DataSourceKindChangeTrackingDataTypeConfiguration                  DataSourceKind = "ChangeTrackingDataTypeConfiguration"
	DataSourceKindChangeTrackingDefaultRegistry                        DataSourceKind = "ChangeTrackingDefaultRegistry"
	DataSourceKindChangeTrackingLinuxPath                              DataSourceKind = "ChangeTrackingLinuxPath"
	DataSourceKindChangeTrackingPath                                   DataSourceKind = "ChangeTrackingPath"
	DataSourceKindChangeTrackingRegistry                               DataSourceKind = "ChangeTrackingRegistry"
	DataSourceKindChangeTrackingServices                               DataSourceKind = "ChangeTrackingServices"
	DataSourceKindCustomLog                                            DataSourceKind = "CustomLog"
	DataSourceKindCustomLogCollection                                  DataSourceKind = "CustomLogCollection"
	DataSourceKindDNSAnalytics                                         DataSourceKind = "DnsAnalytics"
	DataSourceKindGenericDataSource                                    DataSourceKind = "GenericDataSource"
	DataSourceKindIISLogs                                              DataSourceKind = "IISLogs"
	DataSourceKindImportComputerGroup                                  DataSourceKind = "ImportComputerGroup"
	DataSourceKindItsm                                                 DataSourceKind = "Itsm"
	DataSourceKindLinuxChangeTrackingPath                              DataSourceKind = "LinuxChangeTrackingPath"
	DataSourceKindLinuxPerformanceCollection                           DataSourceKind = "LinuxPerformanceCollection"
	DataSourceKindLinuxPerformanceObject                               DataSourceKind = "LinuxPerformanceObject"
	DataSourceKindLinuxSyslog                                          DataSourceKind = "LinuxSyslog"
	DataSourceKindLinuxSyslogCollection                                DataSourceKind = "LinuxSyslogCollection"
	DataSourceKindNetworkMonitoring                                    DataSourceKind = "NetworkMonitoring"
	DataSourceKindOffice365                                            DataSourceKind = "Office365"
	DataSourceKindSQLDataClassification                                DataSourceKind = "SqlDataClassification"
	DataSourceKindSecurityCenterSecurityWindowsBaselineConfiguration   DataSourceKind = "SecurityCenterSecurityWindowsBaselineConfiguration"
	DataSourceKindSecurityEventCollectionConfiguration                 DataSourceKind = "SecurityEventCollectionConfiguration"
	DataSourceKindSecurityInsightsSecurityEventCollectionConfiguration DataSourceKind = "SecurityInsightsSecurityEventCollectionConfiguration"
	DataSourceKindSecurityWindowsBaselineConfiguration                 DataSourceKind = "SecurityWindowsBaselineConfiguration"
	DataSourceKindWindowsEvent                                         DataSourceKind = "WindowsEvent"
	DataSourceKindWindowsPerformanceCounter                            DataSourceKind = "WindowsPerformanceCounter"
	DataSourceKindWindowsTelemetry                                     DataSourceKind = "WindowsTelemetry"
)

// PossibleDataSourceKindValues returns the possible values for the DataSourceKind const type.
func PossibleDataSourceKindValues() []DataSourceKind {
	return []DataSourceKind{
		DataSourceKindApplicationInsights,
		DataSourceKindAzureActivityLog,
		DataSourceKindAzureAuditLog,
		DataSourceKindChangeTrackingContentLocation,
		DataSourceKindChangeTrackingCustomPath,
		DataSourceKindChangeTrackingDataTypeConfiguration,
		DataSourceKindChangeTrackingDefaultRegistry,
		DataSourceKindChangeTrackingLinuxPath,
		DataSourceKindChangeTrackingPath,
		DataSourceKindChangeTrackingRegistry,
		DataSourceKindChangeTrackingServices,
		DataSourceKindCustomLog,
		DataSourceKindCustomLogCollection,
		DataSourceKindDNSAnalytics,
		DataSourceKindGenericDataSource,
		DataSourceKindIISLogs,
		DataSourceKindImportComputerGroup,
		DataSourceKindItsm,
		DataSourceKindLinuxChangeTrackingPath,
		DataSourceKindLinuxPerformanceCollection,
		DataSourceKindLinuxPerformanceObject,
		DataSourceKindLinuxSyslog,
		DataSourceKindLinuxSyslogCollection,
		DataSourceKindNetworkMonitoring,
		DataSourceKindOffice365,
		DataSourceKindSQLDataClassification,
		DataSourceKindSecurityCenterSecurityWindowsBaselineConfiguration,
		DataSourceKindSecurityEventCollectionConfiguration,
		DataSourceKindSecurityInsightsSecurityEventCollectionConfiguration,
		DataSourceKindSecurityWindowsBaselineConfiguration,
		DataSourceKindWindowsEvent,
		DataSourceKindWindowsPerformanceCounter,
		DataSourceKindWindowsTelemetry,
	}
}

// ToPtr returns a *DataSourceKind pointing to the current value.
func (c DataSourceKind) ToPtr() *DataSourceKind {
	return &c
}

// DataSourceType - Linked storage accounts type.
type DataSourceType string

const (
	DataSourceTypeCustomLogs  DataSourceType = "CustomLogs"
	DataSourceTypeAzureWatson DataSourceType = "AzureWatson"
	DataSourceTypeQuery       DataSourceType = "Query"
	DataSourceTypeAlerts      DataSourceType = "Alerts"
)

// PossibleDataSourceTypeValues returns the possible values for the DataSourceType const type.
func PossibleDataSourceTypeValues() []DataSourceType {
	return []DataSourceType{
		DataSourceTypeCustomLogs,
		DataSourceTypeAzureWatson,
		DataSourceTypeQuery,
		DataSourceTypeAlerts,
	}
}

// ToPtr returns a *DataSourceType pointing to the current value.
func (c DataSourceType) ToPtr() *DataSourceType {
	return &c
}

// IdentityType - Type of managed service identity.
type IdentityType string

const (
	IdentityTypeSystemAssigned IdentityType = "SystemAssigned"
	IdentityTypeUserAssigned   IdentityType = "UserAssigned"
	IdentityTypeNone           IdentityType = "None"
)

// PossibleIdentityTypeValues returns the possible values for the IdentityType const type.
func PossibleIdentityTypeValues() []IdentityType {
	return []IdentityType{
		IdentityTypeSystemAssigned,
		IdentityTypeUserAssigned,
		IdentityTypeNone,
	}
}

// ToPtr returns a *IdentityType pointing to the current value.
func (c IdentityType) ToPtr() *IdentityType {
	return &c
}

// LinkedServiceEntityStatus - The provisioning state of the linked service.
type LinkedServiceEntityStatus string

const (
	LinkedServiceEntityStatusDeleting            LinkedServiceEntityStatus = "Deleting"
	LinkedServiceEntityStatusProvisioningAccount LinkedServiceEntityStatus = "ProvisioningAccount"
	LinkedServiceEntityStatusSucceeded           LinkedServiceEntityStatus = "Succeeded"
	LinkedServiceEntityStatusUpdating            LinkedServiceEntityStatus = "Updating"
)

// PossibleLinkedServiceEntityStatusValues returns the possible values for the LinkedServiceEntityStatus const type.
func PossibleLinkedServiceEntityStatusValues() []LinkedServiceEntityStatus {
	return []LinkedServiceEntityStatus{
		LinkedServiceEntityStatusDeleting,
		LinkedServiceEntityStatusProvisioningAccount,
		LinkedServiceEntityStatusSucceeded,
		LinkedServiceEntityStatusUpdating,
	}
}

// ToPtr returns a *LinkedServiceEntityStatus pointing to the current value.
func (c LinkedServiceEntityStatus) ToPtr() *LinkedServiceEntityStatus {
	return &c
}

// PublicNetworkAccessType - The network access type for operating on the Log Analytics Workspace. By default it is Enabled
type PublicNetworkAccessType string

const (
	// PublicNetworkAccessTypeDisabled - Disables public connectivity to Log Analytics through public DNS.
	PublicNetworkAccessTypeDisabled PublicNetworkAccessType = "Disabled"
	// PublicNetworkAccessTypeEnabled - Enables connectivity to Log Analytics through public DNS.
	PublicNetworkAccessTypeEnabled PublicNetworkAccessType = "Enabled"
)

// PossiblePublicNetworkAccessTypeValues returns the possible values for the PublicNetworkAccessType const type.
func PossiblePublicNetworkAccessTypeValues() []PublicNetworkAccessType {
	return []PublicNetworkAccessType{
		PublicNetworkAccessTypeDisabled,
		PublicNetworkAccessTypeEnabled,
	}
}

// ToPtr returns a *PublicNetworkAccessType pointing to the current value.
func (c PublicNetworkAccessType) ToPtr() *PublicNetworkAccessType {
	return &c
}

// PurgeState - Status of the operation represented by the requested Id.
type PurgeState string

const (
	PurgeStateCompleted PurgeState = "completed"
	PurgeStatePending   PurgeState = "pending"
)

// PossiblePurgeStateValues returns the possible values for the PurgeState const type.
func PossiblePurgeStateValues() []PurgeState {
	return []PurgeState{
		PurgeStateCompleted,
		PurgeStatePending,
	}
}

// ToPtr returns a *PurgeState pointing to the current value.
func (c PurgeState) ToPtr() *PurgeState {
	return &c
}

// SKUNameEnum - The name of the Service Tier.
type SKUNameEnum string

const (
	SKUNameEnumCapacityReservation SKUNameEnum = "CapacityReservation"
	SKUNameEnumFree                SKUNameEnum = "Free"
	SKUNameEnumPerGB2018           SKUNameEnum = "PerGB2018"
	SKUNameEnumPerNode             SKUNameEnum = "PerNode"
	SKUNameEnumPremium             SKUNameEnum = "Premium"
	SKUNameEnumStandalone          SKUNameEnum = "Standalone"
	SKUNameEnumStandard            SKUNameEnum = "Standard"
)

// PossibleSKUNameEnumValues returns the possible values for the SKUNameEnum const type.
func PossibleSKUNameEnumValues() []SKUNameEnum {
	return []SKUNameEnum{
		SKUNameEnumCapacityReservation,
		SKUNameEnumFree,
		SKUNameEnumPerGB2018,
		SKUNameEnumPerNode,
		SKUNameEnumPremium,
		SKUNameEnumStandalone,
		SKUNameEnumStandard,
	}
}

// ToPtr returns a *SKUNameEnum pointing to the current value.
func (c SKUNameEnum) ToPtr() *SKUNameEnum {
	return &c
}

// SearchSortEnum - The sort order of the search.
type SearchSortEnum string

const (
	SearchSortEnumAsc  SearchSortEnum = "asc"
	SearchSortEnumDesc SearchSortEnum = "desc"
)

// PossibleSearchSortEnumValues returns the possible values for the SearchSortEnum const type.
func PossibleSearchSortEnumValues() []SearchSortEnum {
	return []SearchSortEnum{
		SearchSortEnumAsc,
		SearchSortEnumDesc,
	}
}

// ToPtr returns a *SearchSortEnum pointing to the current value.
func (c SearchSortEnum) ToPtr() *SearchSortEnum {
	return &c
}

// StorageInsightState - The state of the storage insight connection to the workspace
type StorageInsightState string

const (
	StorageInsightStateERROR StorageInsightState = "ERROR"
	StorageInsightStateOK    StorageInsightState = "OK"
)

// PossibleStorageInsightStateValues returns the possible values for the StorageInsightState const type.
func PossibleStorageInsightStateValues() []StorageInsightState {
	return []StorageInsightState{
		StorageInsightStateERROR,
		StorageInsightStateOK,
	}
}

// ToPtr returns a *StorageInsightState pointing to the current value.
func (c StorageInsightState) ToPtr() *StorageInsightState {
	return &c
}

// Type - The type of the destination resource
type Type string

const (
	TypeEventHub       Type = "EventHub"
	TypeStorageAccount Type = "StorageAccount"
)

// PossibleTypeValues returns the possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{
		TypeEventHub,
		TypeStorageAccount,
	}
}

// ToPtr returns a *Type pointing to the current value.
func (c Type) ToPtr() *Type {
	return &c
}

// WorkspaceEntityStatus - The provisioning state of the workspace.
type WorkspaceEntityStatus string

const (
	WorkspaceEntityStatusCanceled            WorkspaceEntityStatus = "Canceled"
	WorkspaceEntityStatusCreating            WorkspaceEntityStatus = "Creating"
	WorkspaceEntityStatusDeleting            WorkspaceEntityStatus = "Deleting"
	WorkspaceEntityStatusFailed              WorkspaceEntityStatus = "Failed"
	WorkspaceEntityStatusProvisioningAccount WorkspaceEntityStatus = "ProvisioningAccount"
	WorkspaceEntityStatusSucceeded           WorkspaceEntityStatus = "Succeeded"
	WorkspaceEntityStatusUpdating            WorkspaceEntityStatus = "Updating"
)

// PossibleWorkspaceEntityStatusValues returns the possible values for the WorkspaceEntityStatus const type.
func PossibleWorkspaceEntityStatusValues() []WorkspaceEntityStatus {
	return []WorkspaceEntityStatus{
		WorkspaceEntityStatusCanceled,
		WorkspaceEntityStatusCreating,
		WorkspaceEntityStatusDeleting,
		WorkspaceEntityStatusFailed,
		WorkspaceEntityStatusProvisioningAccount,
		WorkspaceEntityStatusSucceeded,
		WorkspaceEntityStatusUpdating,
	}
}

// ToPtr returns a *WorkspaceEntityStatus pointing to the current value.
func (c WorkspaceEntityStatus) ToPtr() *WorkspaceEntityStatus {
	return &c
}

// WorkspaceSKUNameEnum - The name of the SKU.
type WorkspaceSKUNameEnum string

const (
	WorkspaceSKUNameEnumCapacityReservation WorkspaceSKUNameEnum = "CapacityReservation"
	WorkspaceSKUNameEnumFree                WorkspaceSKUNameEnum = "Free"
	WorkspaceSKUNameEnumLACluster           WorkspaceSKUNameEnum = "LACluster"
	WorkspaceSKUNameEnumPerGB2018           WorkspaceSKUNameEnum = "PerGB2018"
	WorkspaceSKUNameEnumPerNode             WorkspaceSKUNameEnum = "PerNode"
	WorkspaceSKUNameEnumPremium             WorkspaceSKUNameEnum = "Premium"
	WorkspaceSKUNameEnumStandalone          WorkspaceSKUNameEnum = "Standalone"
	WorkspaceSKUNameEnumStandard            WorkspaceSKUNameEnum = "Standard"
)

// PossibleWorkspaceSKUNameEnumValues returns the possible values for the WorkspaceSKUNameEnum const type.
func PossibleWorkspaceSKUNameEnumValues() []WorkspaceSKUNameEnum {
	return []WorkspaceSKUNameEnum{
		WorkspaceSKUNameEnumCapacityReservation,
		WorkspaceSKUNameEnumFree,
		WorkspaceSKUNameEnumLACluster,
		WorkspaceSKUNameEnumPerGB2018,
		WorkspaceSKUNameEnumPerNode,
		WorkspaceSKUNameEnumPremium,
		WorkspaceSKUNameEnumStandalone,
		WorkspaceSKUNameEnumStandard,
	}
}

// ToPtr returns a *WorkspaceSKUNameEnum pointing to the current value.
func (c WorkspaceSKUNameEnum) ToPtr() *WorkspaceSKUNameEnum {
	return &c
}
