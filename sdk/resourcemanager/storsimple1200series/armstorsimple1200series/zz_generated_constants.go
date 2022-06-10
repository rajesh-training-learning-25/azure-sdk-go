//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorsimple1200series

const (
	moduleName    = "armstorsimple1200series"
	moduleVersion = "v0.1.0"
)

// AlertEmailNotificationStatus - Value indicating whether user/admins will receive emails when an alert condition occurs
// on the system
type AlertEmailNotificationStatus string

const (
	AlertEmailNotificationStatusEnabled  AlertEmailNotificationStatus = "Enabled"
	AlertEmailNotificationStatusDisabled AlertEmailNotificationStatus = "Disabled"
)

// PossibleAlertEmailNotificationStatusValues returns the possible values for the AlertEmailNotificationStatus const type.
func PossibleAlertEmailNotificationStatusValues() []AlertEmailNotificationStatus {
	return []AlertEmailNotificationStatus{
		AlertEmailNotificationStatusEnabled,
		AlertEmailNotificationStatusDisabled,
	}
}

// AlertScope - Device or Resource alert
type AlertScope string

const (
	AlertScopeResource AlertScope = "Resource"
	AlertScopeDevice   AlertScope = "Device"
)

// PossibleAlertScopeValues returns the possible values for the AlertScope const type.
func PossibleAlertScopeValues() []AlertScope {
	return []AlertScope{
		AlertScopeResource,
		AlertScopeDevice,
	}
}

// AlertSeverity - Severity of the alert
type AlertSeverity string

const (
	AlertSeverityInformational AlertSeverity = "Informational"
	AlertSeverityWarning       AlertSeverity = "Warning"
	AlertSeverityCritical      AlertSeverity = "Critical"
)

// PossibleAlertSeverityValues returns the possible values for the AlertSeverity const type.
func PossibleAlertSeverityValues() []AlertSeverity {
	return []AlertSeverity{
		AlertSeverityInformational,
		AlertSeverityWarning,
		AlertSeverityCritical,
	}
}

// AlertSourceType - Source of the alert
type AlertSourceType string

const (
	AlertSourceTypeResource AlertSourceType = "Resource"
	AlertSourceTypeDevice   AlertSourceType = "Device"
)

// PossibleAlertSourceTypeValues returns the possible values for the AlertSourceType const type.
func PossibleAlertSourceTypeValues() []AlertSourceType {
	return []AlertSourceType{
		AlertSourceTypeResource,
		AlertSourceTypeDevice,
	}
}

// AlertStatus - Status of the alert
type AlertStatus string

const (
	AlertStatusActive  AlertStatus = "Active"
	AlertStatusCleared AlertStatus = "Cleared"
)

// PossibleAlertStatusValues returns the possible values for the AlertStatus const type.
func PossibleAlertStatusValues() []AlertStatus {
	return []AlertStatus{
		AlertStatusActive,
		AlertStatusCleared,
	}
}

// AuthType - Specify the Authentication type
type AuthType string

const (
	AuthTypeInvalid              AuthType = "Invalid"
	AuthTypeAccessControlService AuthType = "AccessControlService"
	AuthTypeAzureActiveDirectory AuthType = "AzureActiveDirectory"
)

// PossibleAuthTypeValues returns the possible values for the AuthType const type.
func PossibleAuthTypeValues() []AuthType {
	return []AuthType{
		AuthTypeInvalid,
		AuthTypeAccessControlService,
		AuthTypeAzureActiveDirectory,
	}
}

// CloudType - The cloud service provider
type CloudType string

const (
	CloudTypeAzure     CloudType = "Azure"
	CloudTypeS3        CloudType = "S3"
	CloudTypeS3RRS     CloudType = "S3_RRS"
	CloudTypeOpenStack CloudType = "OpenStack"
	CloudTypeHP        CloudType = "HP"
)

// PossibleCloudTypeValues returns the possible values for the CloudType const type.
func PossibleCloudTypeValues() []CloudType {
	return []CloudType{
		CloudTypeAzure,
		CloudTypeS3,
		CloudTypeS3RRS,
		CloudTypeOpenStack,
		CloudTypeHP,
	}
}

// ContractVersions - Gets ContractVersion
type ContractVersions string

const (
	ContractVersionsInvalidVersion ContractVersions = "InvalidVersion"
	ContractVersionsV201109        ContractVersions = "V2011_09"
	ContractVersionsV201202        ContractVersions = "V2012_02"
	ContractVersionsV201205        ContractVersions = "V2012_05"
	ContractVersionsV201212        ContractVersions = "V2012_12"
	ContractVersionsV201304        ContractVersions = "V2013_04"
	ContractVersionsV201310        ContractVersions = "V2013_10"
	ContractVersionsV201311        ContractVersions = "V2013_11"
	ContractVersionsV201404        ContractVersions = "V2014_04"
	ContractVersionsV201406        ContractVersions = "V2014_06"
	ContractVersionsV201407        ContractVersions = "V2014_07"
	ContractVersionsV201409        ContractVersions = "V2014_09"
	ContractVersionsV201410        ContractVersions = "V2014_10"
	ContractVersionsV201412        ContractVersions = "V2014_12"
	ContractVersionsV201501        ContractVersions = "V2015_01"
	ContractVersionsV201502        ContractVersions = "V2015_02"
	ContractVersionsV201504        ContractVersions = "V2015_04"
	ContractVersionsV201505        ContractVersions = "V2015_05"
	ContractVersionsV201506        ContractVersions = "V2015_06"
	ContractVersionsV201507        ContractVersions = "V2015_07"
	ContractVersionsV201508        ContractVersions = "V2015_08"
	ContractVersionsV201510        ContractVersions = "V2015_10"
	ContractVersionsV201512        ContractVersions = "V2015_12"
	ContractVersionsV201601        ContractVersions = "V2016_01"
	ContractVersionsV201602        ContractVersions = "V2016_02"
	ContractVersionsV201604        ContractVersions = "V2016_04"
	ContractVersionsV201605        ContractVersions = "V2016_05"
	ContractVersionsV201607        ContractVersions = "V2016_07"
	ContractVersionsV201608        ContractVersions = "V2016_08"
)

// PossibleContractVersionsValues returns the possible values for the ContractVersions const type.
func PossibleContractVersionsValues() []ContractVersions {
	return []ContractVersions{
		ContractVersionsInvalidVersion,
		ContractVersionsV201109,
		ContractVersionsV201202,
		ContractVersionsV201205,
		ContractVersionsV201212,
		ContractVersionsV201304,
		ContractVersionsV201310,
		ContractVersionsV201311,
		ContractVersionsV201404,
		ContractVersionsV201406,
		ContractVersionsV201407,
		ContractVersionsV201409,
		ContractVersionsV201410,
		ContractVersionsV201412,
		ContractVersionsV201501,
		ContractVersionsV201502,
		ContractVersionsV201504,
		ContractVersionsV201505,
		ContractVersionsV201506,
		ContractVersionsV201507,
		ContractVersionsV201508,
		ContractVersionsV201510,
		ContractVersionsV201512,
		ContractVersionsV201601,
		ContractVersionsV201602,
		ContractVersionsV201604,
		ContractVersionsV201605,
		ContractVersionsV201607,
		ContractVersionsV201608,
	}
}

// DataPolicy - The data policy of backed up endpoint.
type DataPolicy string

const (
	DataPolicyInvalid DataPolicy = "Invalid"
	DataPolicyLocal   DataPolicy = "Local"
	DataPolicyTiered  DataPolicy = "Tiered"
	DataPolicyCloud   DataPolicy = "Cloud"
)

// PossibleDataPolicyValues returns the possible values for the DataPolicy const type.
func PossibleDataPolicyValues() []DataPolicy {
	return []DataPolicy{
		DataPolicyInvalid,
		DataPolicyLocal,
		DataPolicyTiered,
		DataPolicyCloud,
	}
}

// DeviceConfigurationStatus - "Complete" if the device has been successfully registered as File/IscsiServer and the creation
// of share/volume is complete, "Pending" if the device is only registered but the creation of share/volume
// is complete is still pending
type DeviceConfigurationStatus string

const (
	DeviceConfigurationStatusComplete DeviceConfigurationStatus = "Complete"
	DeviceConfigurationStatusPending  DeviceConfigurationStatus = "Pending"
)

// PossibleDeviceConfigurationStatusValues returns the possible values for the DeviceConfigurationStatus const type.
func PossibleDeviceConfigurationStatusValues() []DeviceConfigurationStatus {
	return []DeviceConfigurationStatus{
		DeviceConfigurationStatusComplete,
		DeviceConfigurationStatusPending,
	}
}

type DeviceOperation string

const (
	DeviceOperationNone              DeviceOperation = "None"
	DeviceOperationDelete            DeviceOperation = "Delete"
	DeviceOperationDeleteWithWarning DeviceOperation = "DeleteWithWarning"
	DeviceOperationDRSource          DeviceOperation = "DRSource"
	DeviceOperationDRTarget          DeviceOperation = "DRTarget"
	DeviceOperationBrowsable         DeviceOperation = "Browsable"
	DeviceOperationReadOnlyForDR     DeviceOperation = "ReadOnlyForDR"
	DeviceOperationDeactivate        DeviceOperation = "Deactivate"
)

// PossibleDeviceOperationValues returns the possible values for the DeviceOperation const type.
func PossibleDeviceOperationValues() []DeviceOperation {
	return []DeviceOperation{
		DeviceOperationNone,
		DeviceOperationDelete,
		DeviceOperationDeleteWithWarning,
		DeviceOperationDRSource,
		DeviceOperationDRTarget,
		DeviceOperationBrowsable,
		DeviceOperationReadOnlyForDR,
		DeviceOperationDeactivate,
	}
}

// DeviceStatus - Current status of the device
type DeviceStatus string

const (
	DeviceStatusUnknown           DeviceStatus = "Unknown"
	DeviceStatusOnline            DeviceStatus = "Online"
	DeviceStatusOffline           DeviceStatus = "Offline"
	DeviceStatusRequiresAttention DeviceStatus = "RequiresAttention"
	DeviceStatusMaintenanceMode   DeviceStatus = "MaintenanceMode"
	DeviceStatusCreating          DeviceStatus = "Creating"
	DeviceStatusProvisioning      DeviceStatus = "Provisioning"
	DeviceStatusDeleted           DeviceStatus = "Deleted"
	DeviceStatusReadyToSetup      DeviceStatus = "ReadyToSetup"
	DeviceStatusDeactivated       DeviceStatus = "Deactivated"
	DeviceStatusDeactivating      DeviceStatus = "Deactivating"
)

// PossibleDeviceStatusValues returns the possible values for the DeviceStatus const type.
func PossibleDeviceStatusValues() []DeviceStatus {
	return []DeviceStatus{
		DeviceStatusUnknown,
		DeviceStatusOnline,
		DeviceStatusOffline,
		DeviceStatusRequiresAttention,
		DeviceStatusMaintenanceMode,
		DeviceStatusCreating,
		DeviceStatusProvisioning,
		DeviceStatusDeleted,
		DeviceStatusReadyToSetup,
		DeviceStatusDeactivated,
		DeviceStatusDeactivating,
	}
}

// DeviceType - Type of the device
type DeviceType string

const (
	DeviceTypeInvalid                          DeviceType = "Invalid"
	DeviceTypeAppliance                        DeviceType = "Appliance"
	DeviceTypeVirtualAppliance                 DeviceType = "VirtualAppliance"
	DeviceTypeSeries9000OnPremVirtualAppliance DeviceType = "Series9000OnPremVirtualAppliance"
	DeviceTypeSeries9000VirtualAppliance       DeviceType = "Series9000VirtualAppliance"
	DeviceTypeSeries9000PhysicalAppliance      DeviceType = "Series9000PhysicalAppliance"
)

// PossibleDeviceTypeValues returns the possible values for the DeviceType const type.
func PossibleDeviceTypeValues() []DeviceType {
	return []DeviceType{
		DeviceTypeInvalid,
		DeviceTypeAppliance,
		DeviceTypeVirtualAppliance,
		DeviceTypeSeries9000OnPremVirtualAppliance,
		DeviceTypeSeries9000VirtualAppliance,
		DeviceTypeSeries9000PhysicalAppliance,
	}
}

// DhcpStatus - Represents state of DHCP.
type DhcpStatus string

const (
	DhcpStatusEnabled  DhcpStatus = "Enabled"
	DhcpStatusDisabled DhcpStatus = "Disabled"
)

// PossibleDhcpStatusValues returns the possible values for the DhcpStatus const type.
func PossibleDhcpStatusValues() []DhcpStatus {
	return []DhcpStatus{
		DhcpStatusEnabled,
		DhcpStatusDisabled,
	}
}

// DiskStatus - The disk status.
type DiskStatus string

const (
	DiskStatusOnline  DiskStatus = "Online"
	DiskStatusOffline DiskStatus = "Offline"
)

// PossibleDiskStatusValues returns the possible values for the DiskStatus const type.
func PossibleDiskStatusValues() []DiskStatus {
	return []DiskStatus{
		DiskStatusOnline,
		DiskStatusOffline,
	}
}

// DownloadPhase - The download phase.
type DownloadPhase string

const (
	DownloadPhaseUnknown      DownloadPhase = "Unknown"
	DownloadPhaseInitializing DownloadPhase = "Initializing"
	DownloadPhaseDownloading  DownloadPhase = "Downloading"
	DownloadPhaseVerifying    DownloadPhase = "Verifying"
)

// PossibleDownloadPhaseValues returns the possible values for the DownloadPhase const type.
func PossibleDownloadPhaseValues() []DownloadPhase {
	return []DownloadPhase{
		DownloadPhaseUnknown,
		DownloadPhaseInitializing,
		DownloadPhaseDownloading,
		DownloadPhaseVerifying,
	}
}

// EncryptionAlgorithm - Algorithm used to encrypt "Value"
type EncryptionAlgorithm string

const (
	EncryptionAlgorithmNone          EncryptionAlgorithm = "None"
	EncryptionAlgorithmAES256        EncryptionAlgorithm = "AES256"
	EncryptionAlgorithmRSAESPKCS1V15 EncryptionAlgorithm = "RSAES_PKCS1_v_1_5"
)

// PossibleEncryptionAlgorithmValues returns the possible values for the EncryptionAlgorithm const type.
func PossibleEncryptionAlgorithmValues() []EncryptionAlgorithm {
	return []EncryptionAlgorithm{
		EncryptionAlgorithmNone,
		EncryptionAlgorithmAES256,
		EncryptionAlgorithmRSAESPKCS1V15,
	}
}

// EncryptionStatus - The encryption status which indicates if encryption is enabled or not.
type EncryptionStatus string

const (
	EncryptionStatusEnabled  EncryptionStatus = "Enabled"
	EncryptionStatusDisabled EncryptionStatus = "Disabled"
)

// PossibleEncryptionStatusValues returns the possible values for the EncryptionStatus const type.
func PossibleEncryptionStatusValues() []EncryptionStatus {
	return []EncryptionStatus{
		EncryptionStatusEnabled,
		EncryptionStatusDisabled,
	}
}

// InitiatedBy - Gets or sets InitiatedBy
type InitiatedBy string

const (
	InitiatedByManual    InitiatedBy = "Manual"
	InitiatedByScheduled InitiatedBy = "Scheduled"
)

// PossibleInitiatedByValues returns the possible values for the InitiatedBy const type.
func PossibleInitiatedByValues() []InitiatedBy {
	return []InitiatedBy{
		InitiatedByManual,
		InitiatedByScheduled,
	}
}

// JobStatus - Current status of the job
type JobStatus string

const (
	JobStatusInvalid   JobStatus = "Invalid"
	JobStatusRunning   JobStatus = "Running"
	JobStatusSucceeded JobStatus = "Succeeded"
	JobStatusFailed    JobStatus = "Failed"
	JobStatusCanceled  JobStatus = "Canceled"
	JobStatusPaused    JobStatus = "Paused"
	JobStatusScheduled JobStatus = "Scheduled"
)

// PossibleJobStatusValues returns the possible values for the JobStatus const type.
func PossibleJobStatusValues() []JobStatus {
	return []JobStatus{
		JobStatusInvalid,
		JobStatusRunning,
		JobStatusSucceeded,
		JobStatusFailed,
		JobStatusCanceled,
		JobStatusPaused,
		JobStatusScheduled,
	}
}

// JobType - The job type.
type JobType string

const (
	JobTypeBackup          JobType = "Backup"
	JobTypeClone           JobType = "Clone"
	JobTypeFailover        JobType = "Failover"
	JobTypeDownloadUpdates JobType = "DownloadUpdates"
	JobTypeInstallUpdates  JobType = "InstallUpdates"
)

// PossibleJobTypeValues returns the possible values for the JobType const type.
func PossibleJobTypeValues() []JobType {
	return []JobType{
		JobTypeBackup,
		JobTypeClone,
		JobTypeFailover,
		JobTypeDownloadUpdates,
		JobTypeInstallUpdates,
	}
}

// KeyRolloverStatus - The key rollover status which indicates if key rollover is required or not. If secrets encryption has
// been upgraded, then it requires key rollover.
type KeyRolloverStatus string

const (
	KeyRolloverStatusRequired    KeyRolloverStatus = "Required"
	KeyRolloverStatusNotRequired KeyRolloverStatus = "NotRequired"
)

// PossibleKeyRolloverStatusValues returns the possible values for the KeyRolloverStatus const type.
func PossibleKeyRolloverStatusValues() []KeyRolloverStatus {
	return []KeyRolloverStatus{
		KeyRolloverStatusRequired,
		KeyRolloverStatusNotRequired,
	}
}

// ManagerType - Refers to the type of the StorSimple Manager
type ManagerType string

const (
	ManagerTypeGardaV1    ManagerType = "GardaV1"
	ManagerTypeHelsinkiV1 ManagerType = "HelsinkiV1"
)

// PossibleManagerTypeValues returns the possible values for the ManagerType const type.
func PossibleManagerTypeValues() []ManagerType {
	return []ManagerType{
		ManagerTypeGardaV1,
		ManagerTypeHelsinkiV1,
	}
}

// MetricAggregationType - The metric aggregation type
type MetricAggregationType string

const (
	MetricAggregationTypeAverage MetricAggregationType = "Average"
	MetricAggregationTypeLast    MetricAggregationType = "Last"
	MetricAggregationTypeMaximum MetricAggregationType = "Maximum"
	MetricAggregationTypeMinimum MetricAggregationType = "Minimum"
	MetricAggregationTypeNone    MetricAggregationType = "None"
	MetricAggregationTypeTotal   MetricAggregationType = "Total"
)

// PossibleMetricAggregationTypeValues returns the possible values for the MetricAggregationType const type.
func PossibleMetricAggregationTypeValues() []MetricAggregationType {
	return []MetricAggregationType{
		MetricAggregationTypeAverage,
		MetricAggregationTypeLast,
		MetricAggregationTypeMaximum,
		MetricAggregationTypeMinimum,
		MetricAggregationTypeNone,
		MetricAggregationTypeTotal,
	}
}

// MetricUnit - The metric unit
type MetricUnit string

const (
	MetricUnitBytes          MetricUnit = "Bytes"
	MetricUnitBytesPerSecond MetricUnit = "BytesPerSecond"
	MetricUnitCount          MetricUnit = "Count"
	MetricUnitCountPerSecond MetricUnit = "CountPerSecond"
	MetricUnitPercent        MetricUnit = "Percent"
	MetricUnitSeconds        MetricUnit = "Seconds"
)

// PossibleMetricUnitValues returns the possible values for the MetricUnit const type.
func PossibleMetricUnitValues() []MetricUnit {
	return []MetricUnit{
		MetricUnitBytes,
		MetricUnitBytesPerSecond,
		MetricUnitCount,
		MetricUnitCountPerSecond,
		MetricUnitPercent,
		MetricUnitSeconds,
	}
}

// MonitoringStatus - The monitoring status
type MonitoringStatus string

const (
	MonitoringStatusEnabled  MonitoringStatus = "Enabled"
	MonitoringStatusDisabled MonitoringStatus = "Disabled"
)

// PossibleMonitoringStatusValues returns the possible values for the MonitoringStatus const type.
func PossibleMonitoringStatusValues() []MonitoringStatus {
	return []MonitoringStatus{
		MonitoringStatusEnabled,
		MonitoringStatusDisabled,
	}
}

// SSLStatus - SSL needs to be enabled or not
type SSLStatus string

const (
	SSLStatusEnabled  SSLStatus = "Enabled"
	SSLStatusDisabled SSLStatus = "Disabled"
)

// PossibleSSLStatusValues returns the possible values for the SSLStatus const type.
func PossibleSSLStatusValues() []SSLStatus {
	return []SSLStatus{
		SSLStatusEnabled,
		SSLStatusDisabled,
	}
}

// ServiceOwnersAlertNotificationStatus - Value indicating whether service owners will receive emails when an alert condition
// occurs on the system. Applicable only if emailNotification flag is Enabled.
type ServiceOwnersAlertNotificationStatus string

const (
	ServiceOwnersAlertNotificationStatusEnabled  ServiceOwnersAlertNotificationStatus = "Enabled"
	ServiceOwnersAlertNotificationStatusDisabled ServiceOwnersAlertNotificationStatus = "Disabled"
)

// PossibleServiceOwnersAlertNotificationStatusValues returns the possible values for the ServiceOwnersAlertNotificationStatus const type.
func PossibleServiceOwnersAlertNotificationStatusValues() []ServiceOwnersAlertNotificationStatus {
	return []ServiceOwnersAlertNotificationStatus{
		ServiceOwnersAlertNotificationStatusEnabled,
		ServiceOwnersAlertNotificationStatusDisabled,
	}
}

// ShareStatus - The Share Status
type ShareStatus string

const (
	ShareStatusOnline  ShareStatus = "Online"
	ShareStatusOffline ShareStatus = "Offline"
)

// PossibleShareStatusValues returns the possible values for the ShareStatus const type.
func PossibleShareStatusValues() []ShareStatus {
	return []ShareStatus{
		ShareStatusOnline,
		ShareStatusOffline,
	}
}

type SupportedDeviceCapabilities string

const (
	SupportedDeviceCapabilitiesInvalid     SupportedDeviceCapabilities = "Invalid"
	SupportedDeviceCapabilitiesFileServer  SupportedDeviceCapabilities = "FileServer"
	SupportedDeviceCapabilitiesIscsiServer SupportedDeviceCapabilities = "IscsiServer"
)

// PossibleSupportedDeviceCapabilitiesValues returns the possible values for the SupportedDeviceCapabilities const type.
func PossibleSupportedDeviceCapabilitiesValues() []SupportedDeviceCapabilities {
	return []SupportedDeviceCapabilities{
		SupportedDeviceCapabilitiesInvalid,
		SupportedDeviceCapabilitiesFileServer,
		SupportedDeviceCapabilitiesIscsiServer,
	}
}

// TargetType - The target type of the backup.
type TargetType string

const (
	TargetTypeFileServer TargetType = "FileServer"
	TargetTypeDiskServer TargetType = "DiskServer"
)

// PossibleTargetTypeValues returns the possible values for the TargetType const type.
func PossibleTargetTypeValues() []TargetType {
	return []TargetType{
		TargetTypeFileServer,
		TargetTypeDiskServer,
	}
}

// UpdateOperation - The current update operation.
type UpdateOperation string

const (
	UpdateOperationIdle        UpdateOperation = "Idle"
	UpdateOperationScanning    UpdateOperation = "Scanning"
	UpdateOperationDownloading UpdateOperation = "Downloading"
	UpdateOperationInstalling  UpdateOperation = "Installing"
)

// PossibleUpdateOperationValues returns the possible values for the UpdateOperation const type.
func PossibleUpdateOperationValues() []UpdateOperation {
	return []UpdateOperation{
		UpdateOperationIdle,
		UpdateOperationScanning,
		UpdateOperationDownloading,
		UpdateOperationInstalling,
	}
}
