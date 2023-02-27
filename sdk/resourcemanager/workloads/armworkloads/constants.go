//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armworkloads

const (
	moduleName    = "armworkloads"
	moduleVersion = "v0.3.0"
)

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

// ApplicationServerVirtualMachineType - Defines the type of application server VM.
type ApplicationServerVirtualMachineType string

const (
	ApplicationServerVirtualMachineTypeActive  ApplicationServerVirtualMachineType = "Active"
	ApplicationServerVirtualMachineTypeStandby ApplicationServerVirtualMachineType = "Standby"
	ApplicationServerVirtualMachineTypeUnknown ApplicationServerVirtualMachineType = "Unknown"
)

// PossibleApplicationServerVirtualMachineTypeValues returns the possible values for the ApplicationServerVirtualMachineType const type.
func PossibleApplicationServerVirtualMachineTypeValues() []ApplicationServerVirtualMachineType {
	return []ApplicationServerVirtualMachineType{
		ApplicationServerVirtualMachineTypeActive,
		ApplicationServerVirtualMachineTypeStandby,
		ApplicationServerVirtualMachineTypeUnknown,
	}
}

// CentralServerVirtualMachineType - Defines the type of central server VM.
type CentralServerVirtualMachineType string

const (
	CentralServerVirtualMachineTypeASCS        CentralServerVirtualMachineType = "ASCS"
	CentralServerVirtualMachineTypeERS         CentralServerVirtualMachineType = "ERS"
	CentralServerVirtualMachineTypeERSInactive CentralServerVirtualMachineType = "ERSInactive"
	CentralServerVirtualMachineTypePrimary     CentralServerVirtualMachineType = "Primary"
	CentralServerVirtualMachineTypeSecondary   CentralServerVirtualMachineType = "Secondary"
	CentralServerVirtualMachineTypeStandby     CentralServerVirtualMachineType = "Standby"
	CentralServerVirtualMachineTypeUnknown     CentralServerVirtualMachineType = "Unknown"
)

// PossibleCentralServerVirtualMachineTypeValues returns the possible values for the CentralServerVirtualMachineType const type.
func PossibleCentralServerVirtualMachineTypeValues() []CentralServerVirtualMachineType {
	return []CentralServerVirtualMachineType{
		CentralServerVirtualMachineTypeASCS,
		CentralServerVirtualMachineTypeERS,
		CentralServerVirtualMachineTypeERSInactive,
		CentralServerVirtualMachineTypePrimary,
		CentralServerVirtualMachineTypeSecondary,
		CentralServerVirtualMachineTypeStandby,
		CentralServerVirtualMachineTypeUnknown,
	}
}

// ConfigurationType - The type of file share config.
type ConfigurationType string

const (
	ConfigurationTypeCreateAndMount ConfigurationType = "CreateAndMount"
	ConfigurationTypeMount          ConfigurationType = "Mount"
	ConfigurationTypeSkip           ConfigurationType = "Skip"
)

// PossibleConfigurationTypeValues returns the possible values for the ConfigurationType const type.
func PossibleConfigurationTypeValues() []ConfigurationType {
	return []ConfigurationType{
		ConfigurationTypeCreateAndMount,
		ConfigurationTypeMount,
		ConfigurationTypeSkip,
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

// DiskSKUName - Defines the disk sku name.
type DiskSKUName string

const (
	DiskSKUNamePremiumLRS     DiskSKUName = "Premium_LRS"
	DiskSKUNamePremiumV2LRS   DiskSKUName = "PremiumV2_LRS"
	DiskSKUNamePremiumZRS     DiskSKUName = "Premium_ZRS"
	DiskSKUNameStandardLRS    DiskSKUName = "Standard_LRS"
	DiskSKUNameStandardSSDLRS DiskSKUName = "StandardSSD_LRS"
	DiskSKUNameStandardSSDZRS DiskSKUName = "StandardSSD_ZRS"
	DiskSKUNameUltraSSDLRS    DiskSKUName = "UltraSSD_LRS"
)

// PossibleDiskSKUNameValues returns the possible values for the DiskSKUName const type.
func PossibleDiskSKUNameValues() []DiskSKUName {
	return []DiskSKUName{
		DiskSKUNamePremiumLRS,
		DiskSKUNamePremiumV2LRS,
		DiskSKUNamePremiumZRS,
		DiskSKUNameStandardLRS,
		DiskSKUNameStandardSSDLRS,
		DiskSKUNameStandardSSDZRS,
		DiskSKUNameUltraSSDLRS,
	}
}

// EnqueueReplicationServerType - Defines the type of Enqueue Replication Server.
type EnqueueReplicationServerType string

const (
	EnqueueReplicationServerTypeEnqueueReplicator1 EnqueueReplicationServerType = "EnqueueReplicator1"
	EnqueueReplicationServerTypeEnqueueReplicator2 EnqueueReplicationServerType = "EnqueueReplicator2"
)

// PossibleEnqueueReplicationServerTypeValues returns the possible values for the EnqueueReplicationServerType const type.
func PossibleEnqueueReplicationServerTypeValues() []EnqueueReplicationServerType {
	return []EnqueueReplicationServerType{
		EnqueueReplicationServerTypeEnqueueReplicator1,
		EnqueueReplicationServerTypeEnqueueReplicator2,
	}
}

// ManagedServiceIdentityType - Type of managed service identity (only None, UserAssigned types are allowed).
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone         ManagedServiceIdentityType = "None"
	ManagedServiceIdentityTypeUserAssigned ManagedServiceIdentityType = "UserAssigned"
)

// PossibleManagedServiceIdentityTypeValues returns the possible values for the ManagedServiceIdentityType const type.
func PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType {
	return []ManagedServiceIdentityType{
		ManagedServiceIdentityTypeNone,
		ManagedServiceIdentityTypeUserAssigned,
	}
}

// NamingPatternType - The pattern type to be used for resource naming.
type NamingPatternType string

const (
	NamingPatternTypeFullResourceName NamingPatternType = "FullResourceName"
)

// PossibleNamingPatternTypeValues returns the possible values for the NamingPatternType const type.
func PossibleNamingPatternTypeValues() []NamingPatternType {
	return []NamingPatternType{
		NamingPatternTypeFullResourceName,
	}
}

// OSType - The OS Type
type OSType string

const (
	OSTypeLinux   OSType = "Linux"
	OSTypeWindows OSType = "Windows"
)

// PossibleOSTypeValues returns the possible values for the OSType const type.
func PossibleOSTypeValues() []OSType {
	return []OSType{
		OSTypeLinux,
		OSTypeWindows,
	}
}

// OperationProperties - Defines the workload operation origin.
type OperationProperties string

const (
	OperationPropertiesNotSpecified OperationProperties = "NotSpecified"
	OperationPropertiesSystem       OperationProperties = "System"
	OperationPropertiesUser         OperationProperties = "User"
)

// PossibleOperationPropertiesValues returns the possible values for the OperationProperties const type.
func PossibleOperationPropertiesValues() []OperationProperties {
	return []OperationProperties{
		OperationPropertiesNotSpecified,
		OperationPropertiesSystem,
		OperationPropertiesUser,
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

// RoutingPreference - Sets the routing preference of the SAP monitor. By default only RFC1918 traffic is routed to the customer
// VNET.
type RoutingPreference string

const (
	RoutingPreferenceDefault  RoutingPreference = "Default"
	RoutingPreferenceRouteAll RoutingPreference = "RouteAll"
)

// PossibleRoutingPreferenceValues returns the possible values for the RoutingPreference const type.
func PossibleRoutingPreferenceValues() []RoutingPreference {
	return []RoutingPreference{
		RoutingPreferenceDefault,
		RoutingPreferenceRouteAll,
	}
}

// SAPConfigurationType - The configuration Type.
type SAPConfigurationType string

const (
	SAPConfigurationTypeDeployment             SAPConfigurationType = "Deployment"
	SAPConfigurationTypeDeploymentWithOSConfig SAPConfigurationType = "DeploymentWithOSConfig"
	SAPConfigurationTypeDiscovery              SAPConfigurationType = "Discovery"
)

// PossibleSAPConfigurationTypeValues returns the possible values for the SAPConfigurationType const type.
func PossibleSAPConfigurationTypeValues() []SAPConfigurationType {
	return []SAPConfigurationType{
		SAPConfigurationTypeDeployment,
		SAPConfigurationTypeDeploymentWithOSConfig,
		SAPConfigurationTypeDiscovery,
	}
}

// SAPDatabaseScaleMethod - The database scale method.
type SAPDatabaseScaleMethod string

const (
	SAPDatabaseScaleMethodScaleUp SAPDatabaseScaleMethod = "ScaleUp"
)

// PossibleSAPDatabaseScaleMethodValues returns the possible values for the SAPDatabaseScaleMethod const type.
func PossibleSAPDatabaseScaleMethodValues() []SAPDatabaseScaleMethod {
	return []SAPDatabaseScaleMethod{
		SAPDatabaseScaleMethodScaleUp,
	}
}

// SAPDatabaseType - Defines the supported SAP Database types.
type SAPDatabaseType string

const (
	SAPDatabaseTypeDB2  SAPDatabaseType = "DB2"
	SAPDatabaseTypeHANA SAPDatabaseType = "HANA"
)

// PossibleSAPDatabaseTypeValues returns the possible values for the SAPDatabaseType const type.
func PossibleSAPDatabaseTypeValues() []SAPDatabaseType {
	return []SAPDatabaseType{
		SAPDatabaseTypeDB2,
		SAPDatabaseTypeHANA,
	}
}

// SAPDeploymentType - The type of SAP deployment, single server or Three tier.
type SAPDeploymentType string

const (
	SAPDeploymentTypeSingleServer SAPDeploymentType = "SingleServer"
	SAPDeploymentTypeThreeTier    SAPDeploymentType = "ThreeTier"
)

// PossibleSAPDeploymentTypeValues returns the possible values for the SAPDeploymentType const type.
func PossibleSAPDeploymentTypeValues() []SAPDeploymentType {
	return []SAPDeploymentType{
		SAPDeploymentTypeSingleServer,
		SAPDeploymentTypeThreeTier,
	}
}

// SAPEnvironmentType - Defines the environment type - Production/Non Production.
type SAPEnvironmentType string

const (
	SAPEnvironmentTypeNonProd SAPEnvironmentType = "NonProd"
	SAPEnvironmentTypeProd    SAPEnvironmentType = "Prod"
)

// PossibleSAPEnvironmentTypeValues returns the possible values for the SAPEnvironmentType const type.
func PossibleSAPEnvironmentTypeValues() []SAPEnvironmentType {
	return []SAPEnvironmentType{
		SAPEnvironmentTypeNonProd,
		SAPEnvironmentTypeProd,
	}
}

// SAPHealthState - Defines the health of SAP Instances.
type SAPHealthState string

const (
	SAPHealthStateDegraded  SAPHealthState = "Degraded"
	SAPHealthStateHealthy   SAPHealthState = "Healthy"
	SAPHealthStateUnhealthy SAPHealthState = "Unhealthy"
	SAPHealthStateUnknown   SAPHealthState = "Unknown"
)

// PossibleSAPHealthStateValues returns the possible values for the SAPHealthState const type.
func PossibleSAPHealthStateValues() []SAPHealthState {
	return []SAPHealthState{
		SAPHealthStateDegraded,
		SAPHealthStateHealthy,
		SAPHealthStateUnhealthy,
		SAPHealthStateUnknown,
	}
}

// SAPHighAvailabilityType - The high availability Type. AvailabilitySet guarantees 99.95% availability. Availability Zone
// guarantees 99.99% availability.
type SAPHighAvailabilityType string

const (
	SAPHighAvailabilityTypeAvailabilitySet  SAPHighAvailabilityType = "AvailabilitySet"
	SAPHighAvailabilityTypeAvailabilityZone SAPHighAvailabilityType = "AvailabilityZone"
)

// PossibleSAPHighAvailabilityTypeValues returns the possible values for the SAPHighAvailabilityType const type.
func PossibleSAPHighAvailabilityTypeValues() []SAPHighAvailabilityType {
	return []SAPHighAvailabilityType{
		SAPHighAvailabilityTypeAvailabilitySet,
		SAPHighAvailabilityTypeAvailabilityZone,
	}
}

// SAPProductType - Defines the SAP Product type.
type SAPProductType string

const (
	SAPProductTypeECC    SAPProductType = "ECC"
	SAPProductTypeOther  SAPProductType = "Other"
	SAPProductTypeS4HANA SAPProductType = "S4HANA"
)

// PossibleSAPProductTypeValues returns the possible values for the SAPProductType const type.
func PossibleSAPProductTypeValues() []SAPProductType {
	return []SAPProductType{
		SAPProductTypeECC,
		SAPProductTypeOther,
		SAPProductTypeS4HANA,
	}
}

// SAPSoftwareInstallationType - The SAP software installation Type.
type SAPSoftwareInstallationType string

const (
	SAPSoftwareInstallationTypeExternal                  SAPSoftwareInstallationType = "External"
	SAPSoftwareInstallationTypeSAPInstallWithoutOSConfig SAPSoftwareInstallationType = "SAPInstallWithoutOSConfig"
	SAPSoftwareInstallationTypeServiceInitiated          SAPSoftwareInstallationType = "ServiceInitiated"
)

// PossibleSAPSoftwareInstallationTypeValues returns the possible values for the SAPSoftwareInstallationType const type.
func PossibleSAPSoftwareInstallationTypeValues() []SAPSoftwareInstallationType {
	return []SAPSoftwareInstallationType{
		SAPSoftwareInstallationTypeExternal,
		SAPSoftwareInstallationTypeSAPInstallWithoutOSConfig,
		SAPSoftwareInstallationTypeServiceInitiated,
	}
}

// SAPVirtualInstanceState - Defines the Virtual Instance for SAP state.
type SAPVirtualInstanceState string

const (
	SAPVirtualInstanceStateDiscoveryFailed                    SAPVirtualInstanceState = "DiscoveryFailed"
	SAPVirtualInstanceStateDiscoveryInProgress                SAPVirtualInstanceState = "DiscoveryInProgress"
	SAPVirtualInstanceStateDiscoveryPending                   SAPVirtualInstanceState = "DiscoveryPending"
	SAPVirtualInstanceStateInfrastructureDeploymentFailed     SAPVirtualInstanceState = "InfrastructureDeploymentFailed"
	SAPVirtualInstanceStateInfrastructureDeploymentInProgress SAPVirtualInstanceState = "InfrastructureDeploymentInProgress"
	SAPVirtualInstanceStateInfrastructureDeploymentPending    SAPVirtualInstanceState = "InfrastructureDeploymentPending"
	SAPVirtualInstanceStateRegistrationComplete               SAPVirtualInstanceState = "RegistrationComplete"
	SAPVirtualInstanceStateSoftwareDetectionFailed            SAPVirtualInstanceState = "SoftwareDetectionFailed"
	SAPVirtualInstanceStateSoftwareDetectionInProgress        SAPVirtualInstanceState = "SoftwareDetectionInProgress"
	SAPVirtualInstanceStateSoftwareInstallationFailed         SAPVirtualInstanceState = "SoftwareInstallationFailed"
	SAPVirtualInstanceStateSoftwareInstallationInProgress     SAPVirtualInstanceState = "SoftwareInstallationInProgress"
	SAPVirtualInstanceStateSoftwareInstallationPending        SAPVirtualInstanceState = "SoftwareInstallationPending"
)

// PossibleSAPVirtualInstanceStateValues returns the possible values for the SAPVirtualInstanceState const type.
func PossibleSAPVirtualInstanceStateValues() []SAPVirtualInstanceState {
	return []SAPVirtualInstanceState{
		SAPVirtualInstanceStateDiscoveryFailed,
		SAPVirtualInstanceStateDiscoveryInProgress,
		SAPVirtualInstanceStateDiscoveryPending,
		SAPVirtualInstanceStateInfrastructureDeploymentFailed,
		SAPVirtualInstanceStateInfrastructureDeploymentInProgress,
		SAPVirtualInstanceStateInfrastructureDeploymentPending,
		SAPVirtualInstanceStateRegistrationComplete,
		SAPVirtualInstanceStateSoftwareDetectionFailed,
		SAPVirtualInstanceStateSoftwareDetectionInProgress,
		SAPVirtualInstanceStateSoftwareInstallationFailed,
		SAPVirtualInstanceStateSoftwareInstallationInProgress,
		SAPVirtualInstanceStateSoftwareInstallationPending,
	}
}

// SAPVirtualInstanceStatus - Defines the SAP Instance status.
type SAPVirtualInstanceStatus string

const (
	SAPVirtualInstanceStatusOffline          SAPVirtualInstanceStatus = "Offline"
	SAPVirtualInstanceStatusPartiallyRunning SAPVirtualInstanceStatus = "PartiallyRunning"
	SAPVirtualInstanceStatusRunning          SAPVirtualInstanceStatus = "Running"
	SAPVirtualInstanceStatusSoftShutdown     SAPVirtualInstanceStatus = "SoftShutdown"
	SAPVirtualInstanceStatusStarting         SAPVirtualInstanceStatus = "Starting"
	SAPVirtualInstanceStatusStopping         SAPVirtualInstanceStatus = "Stopping"
	SAPVirtualInstanceStatusUnavailable      SAPVirtualInstanceStatus = "Unavailable"
)

// PossibleSAPVirtualInstanceStatusValues returns the possible values for the SAPVirtualInstanceStatus const type.
func PossibleSAPVirtualInstanceStatusValues() []SAPVirtualInstanceStatus {
	return []SAPVirtualInstanceStatus{
		SAPVirtualInstanceStatusOffline,
		SAPVirtualInstanceStatusPartiallyRunning,
		SAPVirtualInstanceStatusRunning,
		SAPVirtualInstanceStatusSoftShutdown,
		SAPVirtualInstanceStatusStarting,
		SAPVirtualInstanceStatusStopping,
		SAPVirtualInstanceStatusUnavailable,
	}
}

// SSLPreference - Gets or sets certificate preference if secure communication is enabled.
type SSLPreference string

const (
	SSLPreferenceDisabled          SSLPreference = "Disabled"
	SSLPreferenceRootCertificate   SSLPreference = "RootCertificate"
	SSLPreferenceServerCertificate SSLPreference = "ServerCertificate"
)

// PossibleSSLPreferenceValues returns the possible values for the SSLPreference const type.
func PossibleSSLPreferenceValues() []SSLPreference {
	return []SSLPreference{
		SSLPreferenceDisabled,
		SSLPreferenceRootCertificate,
		SSLPreferenceServerCertificate,
	}
}

// SapLandscapeMonitorProvisioningState - State of provisioning of the SAP monitor.
type SapLandscapeMonitorProvisioningState string

const (
	SapLandscapeMonitorProvisioningStateAccepted  SapLandscapeMonitorProvisioningState = "Accepted"
	SapLandscapeMonitorProvisioningStateCanceled  SapLandscapeMonitorProvisioningState = "Canceled"
	SapLandscapeMonitorProvisioningStateCreated   SapLandscapeMonitorProvisioningState = "Created"
	SapLandscapeMonitorProvisioningStateFailed    SapLandscapeMonitorProvisioningState = "Failed"
	SapLandscapeMonitorProvisioningStateSucceeded SapLandscapeMonitorProvisioningState = "Succeeded"
)

// PossibleSapLandscapeMonitorProvisioningStateValues returns the possible values for the SapLandscapeMonitorProvisioningState const type.
func PossibleSapLandscapeMonitorProvisioningStateValues() []SapLandscapeMonitorProvisioningState {
	return []SapLandscapeMonitorProvisioningState{
		SapLandscapeMonitorProvisioningStateAccepted,
		SapLandscapeMonitorProvisioningStateCanceled,
		SapLandscapeMonitorProvisioningStateCreated,
		SapLandscapeMonitorProvisioningStateFailed,
		SapLandscapeMonitorProvisioningStateSucceeded,
	}
}

// SapVirtualInstanceProvisioningState - Defines the provisioning states.
type SapVirtualInstanceProvisioningState string

const (
	SapVirtualInstanceProvisioningStateCreating  SapVirtualInstanceProvisioningState = "Creating"
	SapVirtualInstanceProvisioningStateDeleting  SapVirtualInstanceProvisioningState = "Deleting"
	SapVirtualInstanceProvisioningStateFailed    SapVirtualInstanceProvisioningState = "Failed"
	SapVirtualInstanceProvisioningStateSucceeded SapVirtualInstanceProvisioningState = "Succeeded"
	SapVirtualInstanceProvisioningStateUpdating  SapVirtualInstanceProvisioningState = "Updating"
)

// PossibleSapVirtualInstanceProvisioningStateValues returns the possible values for the SapVirtualInstanceProvisioningState const type.
func PossibleSapVirtualInstanceProvisioningStateValues() []SapVirtualInstanceProvisioningState {
	return []SapVirtualInstanceProvisioningState{
		SapVirtualInstanceProvisioningStateCreating,
		SapVirtualInstanceProvisioningStateDeleting,
		SapVirtualInstanceProvisioningStateFailed,
		SapVirtualInstanceProvisioningStateSucceeded,
		SapVirtualInstanceProvisioningStateUpdating,
	}
}

// WorkloadMonitorActionType - Defines the action type of workload operation.
type WorkloadMonitorActionType string

const (
	WorkloadMonitorActionTypeInternal     WorkloadMonitorActionType = "Internal"
	WorkloadMonitorActionTypeNotSpecified WorkloadMonitorActionType = "NotSpecified"
)

// PossibleWorkloadMonitorActionTypeValues returns the possible values for the WorkloadMonitorActionType const type.
func PossibleWorkloadMonitorActionTypeValues() []WorkloadMonitorActionType {
	return []WorkloadMonitorActionType{
		WorkloadMonitorActionTypeInternal,
		WorkloadMonitorActionTypeNotSpecified,
	}
}

// WorkloadMonitorProvisioningState - State of provisioning of the SAP monitor.
type WorkloadMonitorProvisioningState string

const (
	WorkloadMonitorProvisioningStateAccepted  WorkloadMonitorProvisioningState = "Accepted"
	WorkloadMonitorProvisioningStateCreating  WorkloadMonitorProvisioningState = "Creating"
	WorkloadMonitorProvisioningStateDeleting  WorkloadMonitorProvisioningState = "Deleting"
	WorkloadMonitorProvisioningStateFailed    WorkloadMonitorProvisioningState = "Failed"
	WorkloadMonitorProvisioningStateMigrating WorkloadMonitorProvisioningState = "Migrating"
	WorkloadMonitorProvisioningStateSucceeded WorkloadMonitorProvisioningState = "Succeeded"
	WorkloadMonitorProvisioningStateUpdating  WorkloadMonitorProvisioningState = "Updating"
)

// PossibleWorkloadMonitorProvisioningStateValues returns the possible values for the WorkloadMonitorProvisioningState const type.
func PossibleWorkloadMonitorProvisioningStateValues() []WorkloadMonitorProvisioningState {
	return []WorkloadMonitorProvisioningState{
		WorkloadMonitorProvisioningStateAccepted,
		WorkloadMonitorProvisioningStateCreating,
		WorkloadMonitorProvisioningStateDeleting,
		WorkloadMonitorProvisioningStateFailed,
		WorkloadMonitorProvisioningStateMigrating,
		WorkloadMonitorProvisioningStateSucceeded,
		WorkloadMonitorProvisioningStateUpdating,
	}
}
