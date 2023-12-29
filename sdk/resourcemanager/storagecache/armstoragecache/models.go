//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragecache

import "time"

// APIOperation - REST API operation description: see https://github.com/Azure/azure-rest-api-specs/blob/master/documentation/openapi-authoring-automated-guidelines.md#r3023-operationsapiimplementation
type APIOperation struct {
	// The object that represents the operation.
	Display *APIOperationDisplay

	// The flag that indicates whether the operation applies to data plane.
	IsDataAction *bool

	// Operation name: {provider}/{resource}/{operation}
	Name *string

	// Origin of the operation.
	Origin *string

	// Additional details about an operation.
	Properties *APIOperationProperties
}

// APIOperationDisplay - The object that represents the operation.
type APIOperationDisplay struct {
	// The description of the operation
	Description *string

	// Operation type: Read, write, delete, etc.
	Operation *string

	// Service provider: Microsoft.StorageCache
	Provider *string

	// Resource on which the operation is performed: Cache, etc.
	Resource *string
}

// APIOperationListResult - Result of the request to list Resource Provider operations. It contains a list of operations and
// a URL link to get the next set of results.
type APIOperationListResult struct {
	// URL to get the next set of operation list results if there are any.
	NextLink *string

	// List of Resource Provider operations supported by the Microsoft.StorageCache resource provider.
	Value []*APIOperation
}

// APIOperationProperties - Additional details about an operation.
type APIOperationProperties struct {
	// Specification of the all the metrics provided for a resource type.
	ServiceSpecification *APIOperationPropertiesServiceSpecification
}

// APIOperationPropertiesServiceSpecification - Specification of the all the metrics provided for a resource type.
type APIOperationPropertiesServiceSpecification struct {
	// Details about operations related to logs.
	LogSpecifications []*LogSpecification

	// Details about operations related to metrics.
	MetricSpecifications []*MetricSpecification
}

// AmlFilesystem - An AML file system instance. Follows Azure Resource Manager standards: https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/resource-api-reference.md
type AmlFilesystem struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// The managed identity used by the AML file system, if configured.
	Identity *AmlFilesystemIdentity

	// Properties of the AML file system.
	Properties *AmlFilesystemProperties

	// SKU for the resource.
	SKU *SKUName

	// Resource tags.
	Tags map[string]*string

	// Availability zones for resources. This field should only contain a single element in the array.
	Zones []*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// AmlFilesystemArchive - Information about the AML file system archive
type AmlFilesystemArchive struct {
	// READ-ONLY; Lustre file system path to archive relative to the file system root. Specify '/' to archive all modified data.
	FilesystemPath *string

	// READ-ONLY; The status of the archive
	Status *AmlFilesystemArchiveStatus
}

// AmlFilesystemArchiveInfo - Information required to execute the archive operation
type AmlFilesystemArchiveInfo struct {
	// Lustre file system path to archive relative to the file system root. Specify '/' to archive all modified data.
	FilesystemPath *string
}

// AmlFilesystemArchiveStatus - The status of the archive
type AmlFilesystemArchiveStatus struct {
	// READ-ONLY; Server-defined error code for the archive operation
	ErrorCode *string

	// READ-ONLY; Server-defined error message for the archive operation
	ErrorMessage *string

	// READ-ONLY; The time of the last completed archive operation
	LastCompletionTime *time.Time

	// READ-ONLY; The time the latest archive operation started
	LastStartedTime *time.Time

	// READ-ONLY; The completion percentage of the archive operation
	PercentComplete *int32

	// READ-ONLY; The state of the archive operation
	State *ArchiveStatusType
}

// AmlFilesystemCheckSubnetError - The error details provided when the checkAmlFSSubnets call fails.
type AmlFilesystemCheckSubnetError struct {
	// The error details for the AML file system's subnet.
	FilesystemSubnet *AmlFilesystemCheckSubnetErrorFilesystemSubnet
}

// AmlFilesystemCheckSubnetErrorFilesystemSubnet - The error details for the AML file system's subnet.
type AmlFilesystemCheckSubnetErrorFilesystemSubnet struct {
	// The details of the AML file system subnet check.
	Message *string

	// The status of the AML file system subnet check.
	Status *FilesystemSubnetStatusType
}

// AmlFilesystemClientInfo - AML file system client information
type AmlFilesystemClientInfo struct {
	// READ-ONLY; Container Storage Interface information for the AML file system.
	ContainerStorageInterface *AmlFilesystemContainerStorageInterface

	// READ-ONLY; The version of Lustre running in the AML file system
	LustreVersion *string

	// READ-ONLY; The IPv4 address used by clients to mount the AML file system's Lustre Management Service (MGS).
	MgsAddress *string

	// READ-ONLY; Recommended command to mount the AML file system
	MountCommand *string
}

// AmlFilesystemContainerStorageInterface - AML file system container storage interface information
type AmlFilesystemContainerStorageInterface struct {
	// READ-ONLY; Recommended AKS Persistent Volume for the CSI driver, in Base64 encoded YAML
	PersistentVolume *string

	// READ-ONLY; Recommended AKS Persistent Volume Claim for the CSI driver, in Base64 encoded YAML
	PersistentVolumeClaim *string

	// READ-ONLY; Recommended AKS Storage Class for the CSI driver, in Base64 encoded YAML
	StorageClass *string
}

// AmlFilesystemEncryptionSettings - AML file system encryption settings.
type AmlFilesystemEncryptionSettings struct {
	// Specifies the location of the encryption key in Key Vault.
	KeyEncryptionKey *KeyVaultKeyReference
}

// AmlFilesystemHealth - An indication of AML file system health. Gives more information about health than just that related
// to provisioning.
type AmlFilesystemHealth struct {
	// List of AML file system health states.
	State *AmlFilesystemHealthStateType

	// Server-defined error code for the AML file system health
	StatusCode *string

	// Describes the health state.
	StatusDescription *string
}

// AmlFilesystemHsmSettings - AML file system HSM settings.
type AmlFilesystemHsmSettings struct {
	// REQUIRED; Resource ID of storage container used for hydrating the namespace and archiving from the namespace. The resource
	// provider must have permission to create SAS tokens on the storage account.
	Container *string

	// REQUIRED; Resource ID of storage container used for logging events and errors. Must be a separate container in the same
	// storage account as the hydration and archive container. The resource provider must have
	// permission to create SAS tokens on the storage account.
	LoggingContainer *string

	// Only blobs in the non-logging container that start with this path/prefix get hydrated into the cluster namespace.
	ImportPrefix *string
}

// AmlFilesystemIdentity - Managed Identity properties.
type AmlFilesystemIdentity struct {
	// The type of identity used for the resource.
	Type *AmlFilesystemIdentityType

	// A dictionary where each key is a user assigned identity resource ID, and each key's value is an empty dictionary.
	UserAssignedIdentities map[string]*UserAssignedIdentitiesValue

	// READ-ONLY; The principal ID for the user-assigned identity of the resource.
	PrincipalID *string

	// READ-ONLY; The tenant ID associated with the resource.
	TenantID *string
}

// AmlFilesystemProperties - Properties of the AML file system.
type AmlFilesystemProperties struct {
	// REQUIRED; Subnet used for managing the AML file system and for client-facing operations. This subnet should have at least
	// a /24 subnet mask within the VNET's address space.
	FilesystemSubnet *string

	// REQUIRED; Start time of a 30-minute weekly maintenance window.
	MaintenanceWindow *AmlFilesystemPropertiesMaintenanceWindow

	// REQUIRED; The size of the AML file system, in TiB. This might be rounded up.
	StorageCapacityTiB *float32

	// Specifies encryption settings of the AML file system.
	EncryptionSettings *AmlFilesystemEncryptionSettings

	// Hydration and archive settings and status
	Hsm *AmlFilesystemPropertiesHsm

	// READ-ONLY; Client information for the AML file system.
	ClientInfo *AmlFilesystemClientInfo

	// READ-ONLY; Health of the AML file system.
	Health *AmlFilesystemHealth

	// READ-ONLY; ARM provisioning state.
	ProvisioningState *AmlFilesystemProvisioningStateType

	// READ-ONLY; Throughput provisioned in MB per sec, calculated as storageCapacityTiB * per-unit storage throughput
	ThroughputProvisionedMBps *int32
}

// AmlFilesystemPropertiesHsm - Hydration and archive settings and status
type AmlFilesystemPropertiesHsm struct {
	// Specifies HSM settings of the AML file system.
	Settings *AmlFilesystemHsmSettings

	// READ-ONLY; Archive status
	ArchiveStatus []*AmlFilesystemArchive
}

// AmlFilesystemPropertiesMaintenanceWindow - Start time of a 30-minute weekly maintenance window.
type AmlFilesystemPropertiesMaintenanceWindow struct {
	// Day of the week on which the maintenance window will occur.
	DayOfWeek *MaintenanceDayOfWeekType

	// The time of day (in UTC) to start the maintenance window.
	TimeOfDayUTC *string
}

// AmlFilesystemSubnetInfo - Information required to validate the subnet that will be used in AML file system create
type AmlFilesystemSubnetInfo struct {
	// Subnet used for managing the AML file system and for client-facing operations. This subnet should have at least a /24 subnet
	// mask within the VNET's address space.
	FilesystemSubnet *string

	// Region that the AML file system will be created in.
	Location *string

	// SKU for the resource.
	SKU *SKUName

	// The size of the AML file system, in TiB.
	StorageCapacityTiB *float32
}

// AmlFilesystemUpdate - An AML file system update instance.
type AmlFilesystemUpdate struct {
	// Properties of the AML file system.
	Properties *AmlFilesystemUpdateProperties

	// Resource tags.
	Tags map[string]*string
}

// AmlFilesystemUpdateProperties - Properties of the AML file system.
type AmlFilesystemUpdateProperties struct {
	// Specifies encryption settings of the AML file system.
	EncryptionSettings *AmlFilesystemEncryptionSettings

	// Start time of a 30-minute weekly maintenance window.
	MaintenanceWindow *AmlFilesystemUpdatePropertiesMaintenanceWindow
}

// AmlFilesystemUpdatePropertiesMaintenanceWindow - Start time of a 30-minute weekly maintenance window.
type AmlFilesystemUpdatePropertiesMaintenanceWindow struct {
	// Day of the week on which the maintenance window will occur.
	DayOfWeek *MaintenanceDayOfWeekType

	// The time of day (in UTC) to start the maintenance window.
	TimeOfDayUTC *string
}

// AmlFilesystemsListResult - Result of the request to list AML file systems. It contains a list of AML file systems and a
// URL link to get the next set of results.
type AmlFilesystemsListResult struct {
	// URL to get the next set of AML file system list results, if there are any.
	NextLink *string

	// List of AML file systems.
	Value []*AmlFilesystem
}

// AscOperation - The status of operation.
type AscOperation struct {
	// The end time of the operation.
	EndTime *string

	// The error detail of the operation if any.
	Error *ErrorResponse

	// The operation Id.
	ID *string

	// The operation name.
	Name *string

	// Additional operation-specific properties
	Properties *AscOperationProperties

	// The start time of the operation.
	StartTime *string

	// The status of the operation.
	Status *string
}

// AscOperationProperties - Additional operation-specific output.
type AscOperationProperties struct {
	// Additional operation-specific output.
	Output map[string]any
}

// BlobNfsTarget - Properties pertaining to the BlobNfsTarget.
type BlobNfsTarget struct {
	// Resource ID of the storage container.
	Target *string

	// Identifies the StorageCache usage model to be used for this storage target.
	UsageModel *string

	// Amount of time (in seconds) the cache waits before it checks the back-end storage for file updates.
	VerificationTimer *int32

	// Amount of time (in seconds) the cache waits after the last file change before it copies the changed file to back-end storage.
	WriteBackTimer *int32
}

// Cache - A cache instance. Follows Azure Resource Manager standards: https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/resource-api-reference.md
type Cache struct {
	// The identity of the cache, if configured.
	Identity *CacheIdentity

	// Region name string.
	Location *string

	// Properties of the cache.
	Properties *CacheProperties

	// SKU for the cache.
	SKU *CacheSKU

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Resource ID of the cache.
	ID *string

	// READ-ONLY; Name of cache.
	Name *string

	// READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData

	// READ-ONLY; Type of the cache; Microsoft.StorageCache/Cache
	Type *string
}

// CacheActiveDirectorySettings - Active Directory settings used to join a cache to a domain.
type CacheActiveDirectorySettings struct {
	// REQUIRED; The NetBIOS name to assign to the HPC Cache when it joins the Active Directory domain as a server. Length must
	// 1-15 characters from the class [-0-9a-zA-Z].
	CacheNetBiosName *string

	// REQUIRED; The fully qualified domain name of the Active Directory domain controller.
	DomainName *string

	// REQUIRED; The Active Directory domain's NetBIOS name.
	DomainNetBiosName *string

	// REQUIRED; Primary DNS IP address used to resolve the Active Directory domain controller's fully qualified domain name.
	PrimaryDNSIPAddress *string

	// Active Directory admin credentials used to join the HPC Cache to a domain.
	Credentials *CacheActiveDirectorySettingsCredentials

	// Secondary DNS IP address used to resolve the Active Directory domain controller's fully qualified domain name.
	SecondaryDNSIPAddress *string

	// READ-ONLY; True if the HPC Cache is joined to the Active Directory domain.
	DomainJoined *DomainJoinedType
}

// CacheActiveDirectorySettingsCredentials - Active Directory admin credentials used to join the HPC Cache to a domain.
type CacheActiveDirectorySettingsCredentials struct {
	// REQUIRED; Username of the Active Directory domain administrator. This value is stored encrypted and not returned on response.
	Username *string

	// Plain text password of the Active Directory domain administrator. This value is stored encrypted and not returned on response.
	Password *string
}

// CacheDirectorySettings - Cache Directory Services settings.
type CacheDirectorySettings struct {
	// Specifies settings for joining the HPC Cache to an Active Directory domain.
	ActiveDirectory *CacheActiveDirectorySettings

	// Specifies settings for Extended Groups. Extended Groups allows users to be members of more than 16 groups.
	UsernameDownload *CacheUsernameDownloadSettings
}

// CacheEncryptionSettings - Cache encryption settings.
type CacheEncryptionSettings struct {
	// Specifies the location of the key encryption key in key vault.
	KeyEncryptionKey *KeyVaultKeyReference

	// Specifies whether the service will automatically rotate to the newest version of the key in the key vault.
	RotationToLatestKeyVersionEnabled *bool
}

// CacheHealth - An indication of cache health. Gives more information about health than just that related to provisioning.
type CacheHealth struct {
	// List of cache health states. Down is when the cluster is not responding. Degraded is when its functioning but has some
	// alerts. Transitioning when it is creating or deleting. Unknown will be returned
	// in old api versions when a new value is added in future versions. WaitingForKey is when the create is waiting for the system
	// assigned identity to be given access to the encryption key in the
	// encryption settings.
	State *HealthStateType

	// Describes explanation of state.
	StatusDescription *string

	// READ-ONLY; Outstanding conditions that need to be investigated and resolved.
	Conditions []*Condition
}

// CacheIdentity - Cache identity properties.
type CacheIdentity struct {
	// The type of identity used for the cache
	Type *CacheIdentityType

	// A dictionary where each key is a user assigned identity resource ID, and each key's value is an empty dictionary.
	UserAssignedIdentities map[string]*UserAssignedIdentitiesValue

	// READ-ONLY; The principal ID for the system-assigned identity of the cache.
	PrincipalID *string

	// READ-ONLY; The tenant ID associated with the cache.
	TenantID *string
}

// CacheNetworkSettings - Cache network settings.
type CacheNetworkSettings struct {
	// DNS search domain
	DNSSearchDomain *string

	// DNS servers for the cache to use. It will be set from the network configuration if no value is provided.
	DNSServers []*string

	// The IPv4 maximum transmission unit configured for the subnet.
	Mtu *int32

	// NTP server IP Address or FQDN for the cache to use. The default is time.windows.com.
	NtpServer *string

	// READ-ONLY; Array of additional IP addresses used by this cache.
	UtilityAddresses []*string
}

// CacheProperties - Properties of the cache.
type CacheProperties struct {
	// The size of this Cache, in GB.
	CacheSizeGB *int32

	// Specifies Directory Services settings of the cache.
	DirectoryServicesSettings *CacheDirectorySettings

	// Specifies encryption settings of the cache.
	EncryptionSettings *CacheEncryptionSettings

	// Specifies network settings of the cache.
	NetworkSettings *CacheNetworkSettings

	// Specifies security settings of the cache.
	SecuritySettings *CacheSecuritySettings

	// Subnet used for the cache.
	Subnet *string

	// Upgrade settings of the cache.
	UpgradeSettings *CacheUpgradeSettings

	// Availability zones for resources. This field should only contain a single element in the array.
	Zones []*string

	// READ-ONLY; Health of the cache.
	Health *CacheHealth

	// READ-ONLY; Array of IPv4 addresses that can be used by clients mounting this cache.
	MountAddresses []*string

	// READ-ONLY; Specifies the priming jobs defined in the cache.
	PrimingJobs []*PrimingJob

	// READ-ONLY; ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
	ProvisioningState *ProvisioningStateType

	// READ-ONLY; Specifies the space allocation percentage for each storage target in the cache.
	SpaceAllocation []*StorageTargetSpaceAllocation

	// READ-ONLY; Upgrade status of the cache.
	UpgradeStatus *CacheUpgradeStatus
}

// CacheSKU - SKU for the cache.
type CacheSKU struct {
	// SKU name for this cache.
	Name *string
}

// CacheSecuritySettings - Cache security settings.
type CacheSecuritySettings struct {
	// NFS access policies defined for this cache.
	AccessPolicies []*NfsAccessPolicy
}

// CacheUpgradeSettings - Cache Upgrade Settings.
type CacheUpgradeSettings struct {
	// When upgradeScheduleEnabled is true, this field holds the user-chosen upgrade time. At the user-chosen time, the firmware
	// update will automatically be installed on the cache.
	ScheduledTime *time.Time

	// True if the user chooses to select an installation time between now and firmwareUpdateDeadline. Else the firmware will
	// automatically be installed after firmwareUpdateDeadline if not triggered earlier
	// via the upgrade operation.
	UpgradeScheduleEnabled *bool
}

// CacheUpgradeStatus - Properties describing the software upgrade state of the cache.
type CacheUpgradeStatus struct {
	// READ-ONLY; Version string of the firmware currently installed on this cache.
	CurrentFirmwareVersion *string

	// READ-ONLY; Time at which the pending firmware update will automatically be installed on the cache.
	FirmwareUpdateDeadline *time.Time

	// READ-ONLY; True if there is a firmware update ready to install on this cache. The firmware will automatically be installed
	// after firmwareUpdateDeadline if not triggered earlier via the upgrade operation.
	FirmwareUpdateStatus *FirmwareStatusType

	// READ-ONLY; Time of the last successful firmware update.
	LastFirmwareUpdate *time.Time

	// READ-ONLY; When firmwareUpdateAvailable is true, this field holds the version string for the update.
	PendingFirmwareVersion *string
}

// CacheUsernameDownloadSettings - Settings for Extended Groups username and group download.
type CacheUsernameDownloadSettings struct {
	// Determines if the certificate should be automatically downloaded. This applies to 'caCertificateURI' only if 'requireValidCertificate'
	// is true.
	AutoDownloadCertificate *bool

	// The URI of the CA certificate to validate the LDAP secure connection. This field must be populated when 'requireValidCertificate'
	// is set to true.
	CaCertificateURI *string

	// When present, these are the credentials for the secure LDAP connection.
	Credentials *CacheUsernameDownloadSettingsCredentials

	// Whether or not the LDAP connection should be encrypted.
	EncryptLdapConnection *bool

	// Whether or not Extended Groups is enabled.
	ExtendedGroups *bool

	// The URI of the file containing group information (in /etc/group file format). This field must be populated when 'usernameSource'
	// is set to 'File'.
	GroupFileURI *string

	// The base distinguished name for the LDAP domain.
	LdapBaseDN *string

	// The fully qualified domain name or IP address of the LDAP server to use.
	LdapServer *string

	// Determines if the certificates must be validated by a certificate authority. When true, caCertificateURI must be provided.
	RequireValidCertificate *bool

	// The URI of the file containing user information (in /etc/passwd file format). This field must be populated when 'usernameSource'
	// is set to 'File'.
	UserFileURI *string

	// This setting determines how the cache gets username and group names for clients.
	UsernameSource *UsernameSource

	// READ-ONLY; Indicates whether or not the HPC Cache has performed the username download successfully.
	UsernameDownloaded *UsernameDownloadedType
}

// CacheUsernameDownloadSettingsCredentials - When present, these are the credentials for the secure LDAP connection.
type CacheUsernameDownloadSettingsCredentials struct {
	// The Bind Distinguished Name identity to be used in the secure LDAP connection. This value is stored encrypted and not returned
	// on response.
	BindDn *string

	// The Bind password to be used in the secure LDAP connection. This value is stored encrypted and not returned on response.
	BindPassword *string
}

// CachesListResult - Result of the request to list caches. It contains a list of caches and a URL link to get the next set
// of results.
type CachesListResult struct {
	// URL to get the next set of cache list results, if there are any.
	NextLink *string

	// List of Caches.
	Value []*Cache
}

// ClfsTarget - Properties pertaining to the ClfsTarget
type ClfsTarget struct {
	// Resource ID of storage container.
	Target *string
}

// CloudErrorBody - An error response.
type CloudErrorBody struct {
	// An identifier for the error. Codes are invariant and are intended to be consumed programmatically.
	Code *string

	// A list of additional details about the error.
	Details []*CloudErrorBody

	// A message describing the error, intended to be suitable for display in a user interface.
	Message *string

	// The target of the particular error. For example, the name of the property in error.
	Target *string
}

// Condition - Outstanding conditions that will need to be resolved.
type Condition struct {
	// READ-ONLY; The issue requiring attention.
	Message *string

	// READ-ONLY; The time when the condition was raised.
	Timestamp *time.Time
}

// ErrorResponse - Describes the format of Error response.
type ErrorResponse struct {
	// Error code
	Code *string

	// Error message indicating why the operation failed.
	Message *string
}

// KeyVaultKeyReference - Describes a reference to key vault key.
type KeyVaultKeyReference struct {
	// REQUIRED; The URL referencing a key encryption key in key vault.
	KeyURL *string

	// REQUIRED; Describes a resource Id to source key vault.
	SourceVault *KeyVaultKeyReferenceSourceVault
}

// KeyVaultKeyReferenceSourceVault - Describes a resource Id to source key vault.
type KeyVaultKeyReferenceSourceVault struct {
	// Resource Id.
	ID *string
}

// LogSpecification - Details about operation related to logs.
type LogSpecification struct {
	// Localized display name of the log.
	DisplayName *string

	// The name of the log.
	Name *string
}

// MetricDimension - Specifications of the Dimension of metrics.
type MetricDimension struct {
	// Localized friendly display name of the dimension
	DisplayName *string

	// Internal name of the dimension.
	InternalName *string

	// Name of the dimension
	Name *string

	// To be exported to shoe box.
	ToBeExportedForShoebox *bool
}

// MetricSpecification - Details about operation related to metrics.
type MetricSpecification struct {
	// The type of metric aggregation.
	AggregationType *string

	// Dimensions of the metric
	Dimensions []*MetricDimension

	// The description of the metric.
	DisplayDescription *string

	// Localized display name of the metric.
	DisplayName *string

	// Type of metrics.
	MetricClass *string

	// The name of the metric.
	Name *string

	// Support metric aggregation type.
	SupportedAggregationTypes []*MetricAggregationType

	// The unit that the metric is measured in.
	Unit *string
}

// NamespaceJunction - A namespace junction.
type NamespaceJunction struct {
	// Namespace path on a cache for a Storage Target.
	NamespacePath *string

	// Name of the access policy applied to this junction.
	NfsAccessPolicy *string

	// NFS export where targetPath exists.
	NfsExport *string

	// Path in Storage Target to which namespacePath points.
	TargetPath *string
}

// Nfs3Target - Properties pertaining to the Nfs3Target
type Nfs3Target struct {
	// IP address or host name of an NFSv3 host (e.g., 10.0.44.44).
	Target *string

	// Identifies the StorageCache usage model to be used for this storage target.
	UsageModel *string

	// Amount of time (in seconds) the cache waits before it checks the back-end storage for file updates.
	VerificationTimer *int32

	// Amount of time (in seconds) the cache waits after the last file change before it copies the changed file to back-end storage.
	WriteBackTimer *int32
}

// NfsAccessPolicy - A set of rules describing access policies applied to NFSv3 clients of the cache.
type NfsAccessPolicy struct {
	// REQUIRED; The set of rules describing client accesses allowed under this policy.
	AccessRules []*NfsAccessRule

	// REQUIRED; Name identifying this policy. Access Policy names are not case sensitive.
	Name *string
}

// NfsAccessRule - Rule to place restrictions on portions of the cache namespace being presented to clients.
type NfsAccessRule struct {
	// REQUIRED; Access allowed by this rule.
	Access *NfsAccessRuleAccess

	// REQUIRED; Scope for this rule. The scope and filter determine which clients match the rule.
	Scope *NfsAccessRuleScope

	// GID value that replaces 0 when rootSquash is true. This will use the value of anonymousUID if not provided.
	AnonymousGID *string

	// UID value that replaces 0 when rootSquash is true. 65534 will be used if not provided.
	AnonymousUID *string

	// Filter applied to the scope for this rule. The filter's format depends on its scope. 'default' scope matches all clients
	// and has no filter value. 'network' scope takes a filter in CIDR format (for
	// example, 10.99.1.0/24). 'host' takes an IP address or fully qualified domain name as filter. If a client does not match
	// any filter rule and there is no default rule, access is denied.
	Filter *string

	// Map root accesses to anonymousUID and anonymousGID.
	RootSquash *bool

	// For the default policy, allow access to subdirectories under the root export. If this is set to no, clients can only mount
	// the path '/'. If set to yes, clients can mount a deeper path, like '/a/b'.
	SubmountAccess *bool

	// Allow SUID semantics.
	Suid *bool
}

// PrimingJob - A priming job instance.
type PrimingJob struct {
	// REQUIRED; The priming job name.
	PrimingJobName *string

	// REQUIRED; The URL for the priming manifest file to download. This file must be readable from the HPC Cache. When the file
	// is in Azure blob storage the URL should include a Shared Access Signature (SAS) granting
	// read permissions on the blob.
	PrimingManifestURL *string

	// READ-ONLY; The job details or error information if any.
	PrimingJobDetails *string

	// READ-ONLY; The unique identifier of the priming job.
	PrimingJobID *string

	// READ-ONLY; The current progress of the priming job, as a percentage.
	PrimingJobPercentComplete *float64

	// READ-ONLY; The state of the priming operation.
	PrimingJobState *PrimingJobState

	// READ-ONLY; The status code of the priming job.
	PrimingJobStatus *string
}

// PrimingJobIDParameter - Object containing the priming job ID.
type PrimingJobIDParameter struct {
	// REQUIRED; The unique identifier of the priming job.
	PrimingJobID *string
}

// RequiredAmlFilesystemSubnetsSize - Information about the number of available IP addresses that are required for the AML
// file system.
type RequiredAmlFilesystemSubnetsSize struct {
	// The number of available IP addresses that are required for the AML file system.
	FilesystemSubnetSize *int32
}

// RequiredAmlFilesystemSubnetsSizeInfo - Information required to get the number of available IP addresses a subnet should
// have that will be used in AML file system create
type RequiredAmlFilesystemSubnetsSizeInfo struct {
	// SKU for the resource.
	SKU *SKUName

	// The size of the AML file system, in TiB.
	StorageCapacityTiB *float32
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ResourceSKU - A resource SKU.
type ResourceSKU struct {
	// A list of capabilities of this SKU, such as throughput or ops/sec.
	Capabilities []*ResourceSKUCapabilities

	// The set of locations where the SKU is available.
	LocationInfo []*ResourceSKULocationInfo

	// The name of this SKU.
	Name *string

	// The restrictions preventing this SKU from being used. This is empty if there are no restrictions.
	Restrictions []*Restriction

	// READ-ONLY; The set of locations where the SKU is available. This is the supported and registered Azure Geo Regions (e.g.,
	// West US, East US, Southeast Asia, etc.).
	Locations []*string

	// READ-ONLY; The type of resource the SKU applies to.
	ResourceType *string
}

// ResourceSKUCapabilities - A resource SKU capability.
type ResourceSKUCapabilities struct {
	// Name of a capability, such as ops/sec.
	Name *string

	// Quantity, if the capability is measured by quantity.
	Value *string
}

// ResourceSKULocationInfo - Resource SKU location information.
type ResourceSKULocationInfo struct {
	// Location where this SKU is available.
	Location *string

	// Zones if any.
	Zones []*string
}

// ResourceSKUsResult - The response from the List Cache SKUs operation.
type ResourceSKUsResult struct {
	// The URI to fetch the next page of cache SKUs.
	NextLink *string

	// READ-ONLY; The list of SKUs available for the subscription.
	Value []*ResourceSKU
}

// ResourceUsage - The usage and limit (quota) for a resource.
type ResourceUsage struct {
	// READ-ONLY; The current usage of this resource.
	CurrentValue *int32

	// READ-ONLY; The limit (quota) for this resource.
	Limit *int32

	// READ-ONLY; Naming information for this resource type.
	Name *ResourceUsageName

	// READ-ONLY; Unit that the limit and usages are expressed in, such as 'Count'.
	Unit *string
}

// ResourceUsageName - Naming information for this resource type.
type ResourceUsageName struct {
	// Localized name for this resource type.
	LocalizedValue *string

	// Canonical name for this resource type.
	Value *string
}

// ResourceUsagesListResult - Result of the request to list resource usages. It contains a list of resource usages & limits
// and a URL link to get the next set of results.
type ResourceUsagesListResult struct {
	// READ-ONLY; URL to get the next set of resource usage list results if there are any.
	NextLink *string

	// READ-ONLY; List of usages and limits for resources controlled by the Microsoft.StorageCache resource provider.
	Value []*ResourceUsage
}

// Restriction - The restrictions preventing this SKU from being used.
type Restriction struct {
	// The reason for the restriction. As of now this can be "QuotaId" or "NotAvailableForSubscription". "QuotaId" is set when
	// the SKU has requiredQuotas parameter as the subscription does not belong to that
	// quota. "NotAvailableForSubscription" is related to capacity at the datacenter.
	ReasonCode *ReasonCode

	// READ-ONLY; The type of restrictions. In this version, the only possible value for this is location.
	Type *string

	// READ-ONLY; The value of restrictions. If the restriction type is set to location, then this would be the different locations
	// where the SKU is restricted.
	Values []*string
}

// SKUName - SKU for the resource.
type SKUName struct {
	// SKU name for this resource.
	Name *string
}

// StorageTarget - Type of the Storage Target.
type StorageTarget struct {
	// StorageTarget properties
	Properties *StorageTargetProperties

	// READ-ONLY; Resource ID of the Storage Target.
	ID *string

	// READ-ONLY; Region name string.
	Location *string

	// READ-ONLY; Name of the Storage Target.
	Name *string

	// READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData

	// READ-ONLY; Type of the Storage Target; Microsoft.StorageCache/Cache/StorageTarget
	Type *string
}

// StorageTargetProperties - Properties of the Storage Target.
type StorageTargetProperties struct {
	// REQUIRED; Type of the Storage Target.
	TargetType *StorageTargetType

	// Properties when targetType is blobNfs.
	BlobNfs *BlobNfsTarget

	// Properties when targetType is clfs.
	Clfs *ClfsTarget

	// List of cache namespace junctions to target for namespace associations.
	Junctions []*NamespaceJunction

	// Properties when targetType is nfs3.
	Nfs3 *Nfs3Target

	// Storage target operational state.
	State *OperationalStateType

	// Properties when targetType is unknown.
	Unknown *UnknownTarget

	// READ-ONLY; The percentage of cache space allocated for this storage target
	AllocationPercentage *int32

	// READ-ONLY; ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
	ProvisioningState *ProvisioningStateType
}

// StorageTargetResource - Resource used by a cache.
type StorageTargetResource struct {
	// READ-ONLY; Resource ID of the Storage Target.
	ID *string

	// READ-ONLY; Region name string.
	Location *string

	// READ-ONLY; Name of the Storage Target.
	Name *string

	// READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData

	// READ-ONLY; Type of the Storage Target; Microsoft.StorageCache/Cache/StorageTarget
	Type *string
}

// StorageTargetSpaceAllocation - Storage Target space allocation properties.
type StorageTargetSpaceAllocation struct {
	// The percentage of cache space allocated for this storage target
	AllocationPercentage *int32

	// Name of the storage target.
	Name *string
}

// StorageTargetsResult - A list of Storage Targets.
type StorageTargetsResult struct {
	// The URI to fetch the next page of Storage Targets.
	NextLink *string

	// The list of Storage Targets defined for the cache.
	Value []*StorageTarget
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}

// TrackedResource - The resource model definition for an Azure Resource Manager tracked top level resource which has 'tags'
// and a 'location'
type TrackedResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// UnknownTarget - Properties pertaining to the UnknownTarget
type UnknownTarget struct {
	// Dictionary of string->string pairs containing information about the Storage Target.
	Attributes map[string]*string
}

// UsageModel - A usage model.
type UsageModel struct {
	// Localized information describing this usage model.
	Display *UsageModelDisplay

	// Non-localized keyword name for this usage model.
	ModelName *string

	// The type of Storage Target to which this model is applicable (only nfs3 as of this version).
	TargetType *string
}

// UsageModelDisplay - Localized information describing this usage model.
type UsageModelDisplay struct {
	// String to display for this usage model.
	Description *string
}

// UsageModelsResult - A list of cache usage models.
type UsageModelsResult struct {
	// The URI to fetch the next page of cache usage models.
	NextLink *string

	// The list of usage models available for the subscription.
	Value []*UsageModel
}

type UserAssignedIdentitiesValue struct {
	// READ-ONLY; The client ID of the user-assigned identity.
	ClientID *string

	// READ-ONLY; The principal ID of the user-assigned identity.
	PrincipalID *string
}
