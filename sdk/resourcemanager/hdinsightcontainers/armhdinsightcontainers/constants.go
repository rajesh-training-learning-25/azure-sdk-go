//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhdinsightcontainers

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsightcontainers/armhdinsightcontainers"
	moduleVersion = "v0.4.0"
)

// Action - A string property that indicates the action to be performed on the Flink job. It can have one of the following
// enum values => NEW, UPDATE, STATELESSUPDATE, STOP, START, CANCEL, SAVEPOINT, LIST
// SAVEPOINT, or DELETE.
type Action string

const (
	ActionCANCEL          Action = "CANCEL"
	ActionDELETE          Action = "DELETE"
	ActionLASTSTATEUPDATE Action = "LAST_STATE_UPDATE"
	ActionLISTSAVEPOINT   Action = "LIST_SAVEPOINT"
	ActionNEW             Action = "NEW"
	ActionRELAUNCH        Action = "RE_LAUNCH"
	ActionSAVEPOINT       Action = "SAVEPOINT"
	ActionSTART           Action = "START"
	ActionSTATELESSUPDATE Action = "STATELESS_UPDATE"
	ActionSTOP            Action = "STOP"
	ActionUPDATE          Action = "UPDATE"
)

// PossibleActionValues returns the possible values for the Action const type.
func PossibleActionValues() []Action {
	return []Action{
		ActionCANCEL,
		ActionDELETE,
		ActionLASTSTATEUPDATE,
		ActionLISTSAVEPOINT,
		ActionNEW,
		ActionRELAUNCH,
		ActionSAVEPOINT,
		ActionSTART,
		ActionSTATELESSUPDATE,
		ActionSTOP,
		ActionUPDATE,
	}
}

// ActionType - Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

// AutoscaleType - User to specify which type of Autoscale to be implemented - Scheduled Based or Load Based.
type AutoscaleType string

const (
	AutoscaleTypeLoadBased     AutoscaleType = "LoadBased"
	AutoscaleTypeScheduleBased AutoscaleType = "ScheduleBased"
)

// PossibleAutoscaleTypeValues returns the possible values for the AutoscaleType const type.
func PossibleAutoscaleTypeValues() []AutoscaleType {
	return []AutoscaleType{
		AutoscaleTypeLoadBased,
		AutoscaleTypeScheduleBased,
	}
}

type Category string

const (
	CategoryCustom     Category = "custom"
	CategoryPredefined Category = "predefined"
)

// PossibleCategoryValues returns the possible values for the Category const type.
func PossibleCategoryValues() []Category {
	return []Category{
		CategoryCustom,
		CategoryPredefined,
	}
}

// ClusterAvailableUpgradeType - Type of upgrade.
type ClusterAvailableUpgradeType string

const (
	ClusterAvailableUpgradeTypeAKSPatchUpgrade     ClusterAvailableUpgradeType = "AKSPatchUpgrade"
	ClusterAvailableUpgradeTypeHotfixUpgrade       ClusterAvailableUpgradeType = "HotfixUpgrade"
	ClusterAvailableUpgradeTypePatchVersionUpgrade ClusterAvailableUpgradeType = "PatchVersionUpgrade"
)

// PossibleClusterAvailableUpgradeTypeValues returns the possible values for the ClusterAvailableUpgradeType const type.
func PossibleClusterAvailableUpgradeTypeValues() []ClusterAvailableUpgradeType {
	return []ClusterAvailableUpgradeType{
		ClusterAvailableUpgradeTypeAKSPatchUpgrade,
		ClusterAvailableUpgradeTypeHotfixUpgrade,
		ClusterAvailableUpgradeTypePatchVersionUpgrade,
	}
}

// ClusterPoolAvailableUpgradeType - Type of upgrade.
type ClusterPoolAvailableUpgradeType string

const (
	ClusterPoolAvailableUpgradeTypeAKSPatchUpgrade ClusterPoolAvailableUpgradeType = "AKSPatchUpgrade"
	ClusterPoolAvailableUpgradeTypeNodeOsUpgrade   ClusterPoolAvailableUpgradeType = "NodeOsUpgrade"
)

// PossibleClusterPoolAvailableUpgradeTypeValues returns the possible values for the ClusterPoolAvailableUpgradeType const type.
func PossibleClusterPoolAvailableUpgradeTypeValues() []ClusterPoolAvailableUpgradeType {
	return []ClusterPoolAvailableUpgradeType{
		ClusterPoolAvailableUpgradeTypeAKSPatchUpgrade,
		ClusterPoolAvailableUpgradeTypeNodeOsUpgrade,
	}
}

// ClusterPoolUpgradeHistoryType - Type of upgrade.
type ClusterPoolUpgradeHistoryType string

const (
	ClusterPoolUpgradeHistoryTypeAKSPatchUpgrade ClusterPoolUpgradeHistoryType = "AKSPatchUpgrade"
	ClusterPoolUpgradeHistoryTypeNodeOsUpgrade   ClusterPoolUpgradeHistoryType = "NodeOsUpgrade"
)

// PossibleClusterPoolUpgradeHistoryTypeValues returns the possible values for the ClusterPoolUpgradeHistoryType const type.
func PossibleClusterPoolUpgradeHistoryTypeValues() []ClusterPoolUpgradeHistoryType {
	return []ClusterPoolUpgradeHistoryType{
		ClusterPoolUpgradeHistoryTypeAKSPatchUpgrade,
		ClusterPoolUpgradeHistoryTypeNodeOsUpgrade,
	}
}

// ClusterPoolUpgradeHistoryUpgradeResultType - Result of this upgrade.
type ClusterPoolUpgradeHistoryUpgradeResultType string

const (
	ClusterPoolUpgradeHistoryUpgradeResultTypeFailed  ClusterPoolUpgradeHistoryUpgradeResultType = "Failed"
	ClusterPoolUpgradeHistoryUpgradeResultTypeSucceed ClusterPoolUpgradeHistoryUpgradeResultType = "Succeed"
)

// PossibleClusterPoolUpgradeHistoryUpgradeResultTypeValues returns the possible values for the ClusterPoolUpgradeHistoryUpgradeResultType const type.
func PossibleClusterPoolUpgradeHistoryUpgradeResultTypeValues() []ClusterPoolUpgradeHistoryUpgradeResultType {
	return []ClusterPoolUpgradeHistoryUpgradeResultType{
		ClusterPoolUpgradeHistoryUpgradeResultTypeFailed,
		ClusterPoolUpgradeHistoryUpgradeResultTypeSucceed,
	}
}

// ClusterPoolUpgradeType - Type of upgrade.
type ClusterPoolUpgradeType string

const (
	ClusterPoolUpgradeTypeAKSPatchUpgrade ClusterPoolUpgradeType = "AKSPatchUpgrade"
	ClusterPoolUpgradeTypeNodeOsUpgrade   ClusterPoolUpgradeType = "NodeOsUpgrade"
)

// PossibleClusterPoolUpgradeTypeValues returns the possible values for the ClusterPoolUpgradeType const type.
func PossibleClusterPoolUpgradeTypeValues() []ClusterPoolUpgradeType {
	return []ClusterPoolUpgradeType{
		ClusterPoolUpgradeTypeAKSPatchUpgrade,
		ClusterPoolUpgradeTypeNodeOsUpgrade,
	}
}

// ClusterUpgradeHistorySeverityType - Severity of this upgrade.
type ClusterUpgradeHistorySeverityType string

const (
	ClusterUpgradeHistorySeverityTypeCritical ClusterUpgradeHistorySeverityType = "critical"
	ClusterUpgradeHistorySeverityTypeHigh     ClusterUpgradeHistorySeverityType = "high"
	ClusterUpgradeHistorySeverityTypeLow      ClusterUpgradeHistorySeverityType = "low"
	ClusterUpgradeHistorySeverityTypeMedium   ClusterUpgradeHistorySeverityType = "medium"
)

// PossibleClusterUpgradeHistorySeverityTypeValues returns the possible values for the ClusterUpgradeHistorySeverityType const type.
func PossibleClusterUpgradeHistorySeverityTypeValues() []ClusterUpgradeHistorySeverityType {
	return []ClusterUpgradeHistorySeverityType{
		ClusterUpgradeHistorySeverityTypeCritical,
		ClusterUpgradeHistorySeverityTypeHigh,
		ClusterUpgradeHistorySeverityTypeLow,
		ClusterUpgradeHistorySeverityTypeMedium,
	}
}

// ClusterUpgradeHistoryType - Type of upgrade.
type ClusterUpgradeHistoryType string

const (
	ClusterUpgradeHistoryTypeAKSPatchUpgrade             ClusterUpgradeHistoryType = "AKSPatchUpgrade"
	ClusterUpgradeHistoryTypeHotfixUpgrade               ClusterUpgradeHistoryType = "HotfixUpgrade"
	ClusterUpgradeHistoryTypeHotfixUpgradeRollback       ClusterUpgradeHistoryType = "HotfixUpgradeRollback"
	ClusterUpgradeHistoryTypePatchVersionUpgrade         ClusterUpgradeHistoryType = "PatchVersionUpgrade"
	ClusterUpgradeHistoryTypePatchVersionUpgradeRollback ClusterUpgradeHistoryType = "PatchVersionUpgradeRollback"
)

// PossibleClusterUpgradeHistoryTypeValues returns the possible values for the ClusterUpgradeHistoryType const type.
func PossibleClusterUpgradeHistoryTypeValues() []ClusterUpgradeHistoryType {
	return []ClusterUpgradeHistoryType{
		ClusterUpgradeHistoryTypeAKSPatchUpgrade,
		ClusterUpgradeHistoryTypeHotfixUpgrade,
		ClusterUpgradeHistoryTypeHotfixUpgradeRollback,
		ClusterUpgradeHistoryTypePatchVersionUpgrade,
		ClusterUpgradeHistoryTypePatchVersionUpgradeRollback,
	}
}

// ClusterUpgradeHistoryUpgradeResultType - Result of this upgrade.
type ClusterUpgradeHistoryUpgradeResultType string

const (
	ClusterUpgradeHistoryUpgradeResultTypeFailed  ClusterUpgradeHistoryUpgradeResultType = "Failed"
	ClusterUpgradeHistoryUpgradeResultTypeSucceed ClusterUpgradeHistoryUpgradeResultType = "Succeed"
)

// PossibleClusterUpgradeHistoryUpgradeResultTypeValues returns the possible values for the ClusterUpgradeHistoryUpgradeResultType const type.
func PossibleClusterUpgradeHistoryUpgradeResultTypeValues() []ClusterUpgradeHistoryUpgradeResultType {
	return []ClusterUpgradeHistoryUpgradeResultType{
		ClusterUpgradeHistoryUpgradeResultTypeFailed,
		ClusterUpgradeHistoryUpgradeResultTypeSucceed,
	}
}

// ClusterUpgradeType - Type of upgrade.
type ClusterUpgradeType string

const (
	ClusterUpgradeTypeAKSPatchUpgrade     ClusterUpgradeType = "AKSPatchUpgrade"
	ClusterUpgradeTypeHotfixUpgrade       ClusterUpgradeType = "HotfixUpgrade"
	ClusterUpgradeTypePatchVersionUpgrade ClusterUpgradeType = "PatchVersionUpgrade"
)

// PossibleClusterUpgradeTypeValues returns the possible values for the ClusterUpgradeType const type.
func PossibleClusterUpgradeTypeValues() []ClusterUpgradeType {
	return []ClusterUpgradeType{
		ClusterUpgradeTypeAKSPatchUpgrade,
		ClusterUpgradeTypeHotfixUpgrade,
		ClusterUpgradeTypePatchVersionUpgrade,
	}
}

// ComparisonOperator - The comparison operator.
type ComparisonOperator string

const (
	ComparisonOperatorGreaterThan        ComparisonOperator = "greaterThan"
	ComparisonOperatorGreaterThanOrEqual ComparisonOperator = "greaterThanOrEqual"
	ComparisonOperatorLessThan           ComparisonOperator = "lessThan"
	ComparisonOperatorLessThanOrEqual    ComparisonOperator = "lessThanOrEqual"
)

// PossibleComparisonOperatorValues returns the possible values for the ComparisonOperator const type.
func PossibleComparisonOperatorValues() []ComparisonOperator {
	return []ComparisonOperator{
		ComparisonOperatorGreaterThan,
		ComparisonOperatorGreaterThanOrEqual,
		ComparisonOperatorLessThan,
		ComparisonOperatorLessThanOrEqual,
	}
}

// ContentEncoding - This property indicates if the content is encoded and is case-insensitive. Please set the value to base64
// if the content is base64 encoded. Set it to none or skip it if the content is plain text.
type ContentEncoding string

const (
	ContentEncodingBase64 ContentEncoding = "Base64"
	ContentEncodingNone   ContentEncoding = "None"
)

// PossibleContentEncodingValues returns the possible values for the ContentEncoding const type.
func PossibleContentEncodingValues() []ContentEncoding {
	return []ContentEncoding{
		ContentEncodingBase64,
		ContentEncodingNone,
	}
}

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

// CurrentClusterAksVersionStatus - Current AKS version's status: whether it is deprecated or supported
type CurrentClusterAksVersionStatus string

const (
	CurrentClusterAksVersionStatusDeprecated CurrentClusterAksVersionStatus = "Deprecated"
	CurrentClusterAksVersionStatusSupported  CurrentClusterAksVersionStatus = "Supported"
)

// PossibleCurrentClusterAksVersionStatusValues returns the possible values for the CurrentClusterAksVersionStatus const type.
func PossibleCurrentClusterAksVersionStatusValues() []CurrentClusterAksVersionStatus {
	return []CurrentClusterAksVersionStatus{
		CurrentClusterAksVersionStatusDeprecated,
		CurrentClusterAksVersionStatusSupported,
	}
}

// CurrentClusterPoolAksVersionStatus - Current AKS version's status: whether it is deprecated or supported
type CurrentClusterPoolAksVersionStatus string

const (
	CurrentClusterPoolAksVersionStatusDeprecated CurrentClusterPoolAksVersionStatus = "Deprecated"
	CurrentClusterPoolAksVersionStatusSupported  CurrentClusterPoolAksVersionStatus = "Supported"
)

// PossibleCurrentClusterPoolAksVersionStatusValues returns the possible values for the CurrentClusterPoolAksVersionStatus const type.
func PossibleCurrentClusterPoolAksVersionStatusValues() []CurrentClusterPoolAksVersionStatus {
	return []CurrentClusterPoolAksVersionStatus{
		CurrentClusterPoolAksVersionStatusDeprecated,
		CurrentClusterPoolAksVersionStatusSupported,
	}
}

// DataDiskType - Managed Disk Type.
type DataDiskType string

const (
	DataDiskTypePremiumSSDLRS   DataDiskType = "Premium_SSD_LRS"
	DataDiskTypePremiumSSDV2LRS DataDiskType = "Premium_SSD_v2_LRS"
	DataDiskTypePremiumSSDZRS   DataDiskType = "Premium_SSD_ZRS"
	DataDiskTypeStandardHDDLRS  DataDiskType = "Standard_HDD_LRS"
	DataDiskTypeStandardSSDLRS  DataDiskType = "Standard_SSD_LRS"
	DataDiskTypeStandardSSDZRS  DataDiskType = "Standard_SSD_ZRS"
)

// PossibleDataDiskTypeValues returns the possible values for the DataDiskType const type.
func PossibleDataDiskTypeValues() []DataDiskType {
	return []DataDiskType{
		DataDiskTypePremiumSSDLRS,
		DataDiskTypePremiumSSDV2LRS,
		DataDiskTypePremiumSSDZRS,
		DataDiskTypeStandardHDDLRS,
		DataDiskTypeStandardSSDLRS,
		DataDiskTypeStandardSSDZRS,
	}
}

// DbConnectionAuthenticationMode - The authentication mode to connect to your Hive metastore database. More details:
// https://learn.microsoft.com/en-us/azure/azure-sql/database/logins-create-manage?view=azuresql#authentication-and-authorization
type DbConnectionAuthenticationMode string

const (
	// DbConnectionAuthenticationModeIdentityAuth - The managed-identity-based authentication to connect to your Hive metastore
	// database.
	DbConnectionAuthenticationModeIdentityAuth DbConnectionAuthenticationMode = "IdentityAuth"
	// DbConnectionAuthenticationModeSQLAuth - The password-based authentication to connect to your Hive metastore database.
	DbConnectionAuthenticationModeSQLAuth DbConnectionAuthenticationMode = "SqlAuth"
)

// PossibleDbConnectionAuthenticationModeValues returns the possible values for the DbConnectionAuthenticationMode const type.
func PossibleDbConnectionAuthenticationModeValues() []DbConnectionAuthenticationMode {
	return []DbConnectionAuthenticationMode{
		DbConnectionAuthenticationModeIdentityAuth,
		DbConnectionAuthenticationModeSQLAuth,
	}
}

// DeploymentMode - A string property that indicates the deployment mode of Flink cluster. It can have one of the following
// enum values => Application, Session. Default value is Session
type DeploymentMode string

const (
	DeploymentModeApplication DeploymentMode = "Application"
	DeploymentModeSession     DeploymentMode = "Session"
)

// PossibleDeploymentModeValues returns the possible values for the DeploymentMode const type.
func PossibleDeploymentModeValues() []DeploymentMode {
	return []DeploymentMode{
		DeploymentModeApplication,
		DeploymentModeSession,
	}
}

// JobType - Type of cluster job.
type JobType string

const (
	JobTypeFlinkJob JobType = "FlinkJob"
)

// PossibleJobTypeValues returns the possible values for the JobType const type.
func PossibleJobTypeValues() []JobType {
	return []JobType{
		JobTypeFlinkJob,
	}
}

// KeyVaultObjectType - Type of key vault object: secret, key or certificate.
type KeyVaultObjectType string

const (
	KeyVaultObjectTypeCertificate KeyVaultObjectType = "Certificate"
	KeyVaultObjectTypeKey         KeyVaultObjectType = "Key"
	KeyVaultObjectTypeSecret      KeyVaultObjectType = "Secret"
)

// PossibleKeyVaultObjectTypeValues returns the possible values for the KeyVaultObjectType const type.
func PossibleKeyVaultObjectTypeValues() []KeyVaultObjectType {
	return []KeyVaultObjectType{
		KeyVaultObjectTypeCertificate,
		KeyVaultObjectTypeKey,
		KeyVaultObjectTypeSecret,
	}
}

// LibraryManagementAction - The library management action.
type LibraryManagementAction string

const (
	LibraryManagementActionInstall   LibraryManagementAction = "Install"
	LibraryManagementActionUninstall LibraryManagementAction = "Uninstall"
)

// PossibleLibraryManagementActionValues returns the possible values for the LibraryManagementAction const type.
func PossibleLibraryManagementActionValues() []LibraryManagementAction {
	return []LibraryManagementAction{
		LibraryManagementActionInstall,
		LibraryManagementActionUninstall,
	}
}

// ManagedIdentityType - The type of managed identity.
type ManagedIdentityType string

const (
	ManagedIdentityTypeCluster  ManagedIdentityType = "cluster"
	ManagedIdentityTypeInternal ManagedIdentityType = "internal"
	ManagedIdentityTypeUser     ManagedIdentityType = "user"
)

// PossibleManagedIdentityTypeValues returns the possible values for the ManagedIdentityType const type.
func PossibleManagedIdentityTypeValues() []ManagedIdentityType {
	return []ManagedIdentityType{
		ManagedIdentityTypeCluster,
		ManagedIdentityTypeInternal,
		ManagedIdentityTypeUser,
	}
}

// MetastoreDbConnectionAuthenticationMode - The authentication mode to connect to your Hive metastore database. More details:
// https://learn.microsoft.com/en-us/azure/azure-sql/database/logins-create-manage?view=azuresql#authentication-and-authorization
type MetastoreDbConnectionAuthenticationMode string

const (
	// MetastoreDbConnectionAuthenticationModeIdentityAuth - The managed-identity-based authentication to connect to your Hive
	// metastore database.
	MetastoreDbConnectionAuthenticationModeIdentityAuth MetastoreDbConnectionAuthenticationMode = "IdentityAuth"
	// MetastoreDbConnectionAuthenticationModeSQLAuth - The password-based authentication to connect to your Hive metastore database.
	MetastoreDbConnectionAuthenticationModeSQLAuth MetastoreDbConnectionAuthenticationMode = "SqlAuth"
)

// PossibleMetastoreDbConnectionAuthenticationModeValues returns the possible values for the MetastoreDbConnectionAuthenticationMode const type.
func PossibleMetastoreDbConnectionAuthenticationModeValues() []MetastoreDbConnectionAuthenticationMode {
	return []MetastoreDbConnectionAuthenticationMode{
		MetastoreDbConnectionAuthenticationModeIdentityAuth,
		MetastoreDbConnectionAuthenticationModeSQLAuth,
	}
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	OriginSystem     Origin = "system"
	OriginUser       Origin = "user"
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// OutboundType - This can only be set at cluster pool creation time and cannot be changed later.
type OutboundType string

const (
	// OutboundTypeLoadBalancer - The load balancer is used for egress through an AKS assigned public IP. This supports Kubernetes
	// services of type 'loadBalancer'.
	OutboundTypeLoadBalancer OutboundType = "loadBalancer"
	// OutboundTypeUserDefinedRouting - Egress paths must be defined by the user. This is an advanced scenario and requires proper
	// network configuration.
	OutboundTypeUserDefinedRouting OutboundType = "userDefinedRouting"
)

// PossibleOutboundTypeValues returns the possible values for the OutboundType const type.
func PossibleOutboundTypeValues() []OutboundType {
	return []OutboundType{
		OutboundTypeLoadBalancer,
		OutboundTypeUserDefinedRouting,
	}
}

// ProvisioningStatus - Provisioning state of the resource.
type ProvisioningStatus string

const (
	ProvisioningStatusAccepted  ProvisioningStatus = "Accepted"
	ProvisioningStatusCanceled  ProvisioningStatus = "Canceled"
	ProvisioningStatusFailed    ProvisioningStatus = "Failed"
	ProvisioningStatusSucceeded ProvisioningStatus = "Succeeded"
)

// PossibleProvisioningStatusValues returns the possible values for the ProvisioningStatus const type.
func PossibleProvisioningStatusValues() []ProvisioningStatus {
	return []ProvisioningStatus{
		ProvisioningStatusAccepted,
		ProvisioningStatusCanceled,
		ProvisioningStatusFailed,
		ProvisioningStatusSucceeded,
	}
}

// RangerUsersyncMode - User & groups can be synced automatically or via a static list that's refreshed.
type RangerUsersyncMode string

const (
	RangerUsersyncModeAutomatic RangerUsersyncMode = "automatic"
	RangerUsersyncModeStatic    RangerUsersyncMode = "static"
)

// PossibleRangerUsersyncModeValues returns the possible values for the RangerUsersyncMode const type.
func PossibleRangerUsersyncModeValues() []RangerUsersyncMode {
	return []RangerUsersyncMode{
		RangerUsersyncModeAutomatic,
		RangerUsersyncModeStatic,
	}
}

// ScaleActionType - The action type.
type ScaleActionType string

const (
	ScaleActionTypeScaledown ScaleActionType = "scaledown"
	ScaleActionTypeScaleup   ScaleActionType = "scaleup"
)

// PossibleScaleActionTypeValues returns the possible values for the ScaleActionType const type.
func PossibleScaleActionTypeValues() []ScaleActionType {
	return []ScaleActionType{
		ScaleActionTypeScaledown,
		ScaleActionTypeScaleup,
	}
}

type ScheduleDay string

const (
	ScheduleDayFriday    ScheduleDay = "Friday"
	ScheduleDayMonday    ScheduleDay = "Monday"
	ScheduleDaySaturday  ScheduleDay = "Saturday"
	ScheduleDaySunday    ScheduleDay = "Sunday"
	ScheduleDayThursday  ScheduleDay = "Thursday"
	ScheduleDayTuesday   ScheduleDay = "Tuesday"
	ScheduleDayWednesday ScheduleDay = "Wednesday"
)

// PossibleScheduleDayValues returns the possible values for the ScheduleDay const type.
func PossibleScheduleDayValues() []ScheduleDay {
	return []ScheduleDay{
		ScheduleDayFriday,
		ScheduleDayMonday,
		ScheduleDaySaturday,
		ScheduleDaySunday,
		ScheduleDayThursday,
		ScheduleDayTuesday,
		ScheduleDayWednesday,
	}
}

// Severity - Severity of this upgrade.
type Severity string

const (
	SeverityCritical Severity = "critical"
	SeverityHigh     Severity = "high"
	SeverityLow      Severity = "low"
	SeverityMedium   Severity = "medium"
)

// PossibleSeverityValues returns the possible values for the Severity const type.
func PossibleSeverityValues() []Severity {
	return []Severity{
		SeverityCritical,
		SeverityHigh,
		SeverityLow,
		SeverityMedium,
	}
}

// Status - Status of the library.
type Status string

const (
	StatusINSTALLED       Status = "INSTALLED"
	StatusINSTALLFAILED   Status = "INSTALL_FAILED"
	StatusINSTALLING      Status = "INSTALLING"
	StatusUNINSTALLFAILED Status = "UNINSTALL_FAILED"
	StatusUNINSTALLING    Status = "UNINSTALLING"
)

// PossibleStatusValues returns the possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{
		StatusINSTALLED,
		StatusINSTALLFAILED,
		StatusINSTALLING,
		StatusUNINSTALLFAILED,
		StatusUNINSTALLING,
	}
}

// Type - Type of the library.
type Type string

const (
	TypeMaven Type = "maven"
	TypePypi  Type = "pypi"
)

// PossibleTypeValues returns the possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{
		TypeMaven,
		TypePypi,
	}
}

// UpgradeMode - A string property that indicates the upgrade mode to be performed on the Flink job. It can have one of the
// following enum values => STATELESSUPDATE, UPDATE, LASTSTATE_UPDATE.
type UpgradeMode string

const (
	UpgradeModeLASTSTATEUPDATE UpgradeMode = "LAST_STATE_UPDATE"
	UpgradeModeSTATELESSUPDATE UpgradeMode = "STATELESS_UPDATE"
	UpgradeModeUPDATE          UpgradeMode = "UPDATE"
)

// PossibleUpgradeModeValues returns the possible values for the UpgradeMode const type.
func PossibleUpgradeModeValues() []UpgradeMode {
	return []UpgradeMode{
		UpgradeModeLASTSTATEUPDATE,
		UpgradeModeSTATELESSUPDATE,
		UpgradeModeUPDATE,
	}
}
