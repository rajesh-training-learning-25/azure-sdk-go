//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsqlvirtualmachine

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// AdditionalFeaturesServerConfigurations - Additional SQL Server feature settings.
type AdditionalFeaturesServerConfigurations struct {
	// Enable or disable R services (SQL 2016 onwards).
	IsRServicesEnabled *bool `json:"isRServicesEnabled,omitempty"`
}

// AutoBackupSettings - Configure backups for databases in your SQL virtual machine.
type AutoBackupSettings struct {
	// Backup schedule type.
	BackupScheduleType *BackupScheduleType `json:"backupScheduleType,omitempty"`

	// Include or exclude system databases from auto backup.
	BackupSystemDbs *bool `json:"backupSystemDbs,omitempty"`

	// Enable or disable autobackup on SQL virtual machine.
	Enable *bool `json:"enable,omitempty"`

	// Enable or disable encryption for backup on SQL virtual machine.
	EnableEncryption *bool `json:"enableEncryption,omitempty"`

	// Frequency of full backups. In both cases, full backups begin during the next scheduled time window.
	FullBackupFrequency *FullBackupFrequencyType `json:"fullBackupFrequency,omitempty"`

	// Start time of a given day during which full backups can take place. 0-23 hours.
	FullBackupStartTime *int32 `json:"fullBackupStartTime,omitempty"`

	// Duration of the time window of a given day during which full backups can take place. 1-23 hours.
	FullBackupWindowHours *int32 `json:"fullBackupWindowHours,omitempty"`

	// Frequency of log backups. 5-60 minutes.
	LogBackupFrequency *int32 `json:"logBackupFrequency,omitempty"`

	// Password for encryption on backup.
	Password *string `json:"password,omitempty"`

	// Retention period of backup: 1-30 days.
	RetentionPeriod *int32 `json:"retentionPeriod,omitempty"`

	// Storage account key where backup will be taken to.
	StorageAccessKey *string `json:"storageAccessKey,omitempty"`

	// Storage account url where backup will be taken to.
	StorageAccountURL *string `json:"storageAccountUrl,omitempty"`
}

// AutoPatchingSettings - Set a patching window during which Windows and SQL patches will be applied.
type AutoPatchingSettings struct {
	// Day of week to apply the patch on.
	DayOfWeek *DayOfWeek `json:"dayOfWeek,omitempty"`

	// Enable or disable autopatching on SQL virtual machine.
	Enable *bool `json:"enable,omitempty"`

	// Duration of patching.
	MaintenanceWindowDuration *int32 `json:"maintenanceWindowDuration,omitempty"`

	// Hour of the day when patching is initiated. Local VM time.
	MaintenanceWindowStartingHour *int32 `json:"maintenanceWindowStartingHour,omitempty"`
}

// AvailabilityGroupListener - A SQL Server availability group listener.
type AvailabilityGroupListener struct {
	ProxyResource
	// Resource properties.
	Properties *AvailabilityGroupListenerProperties `json:"properties,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type AvailabilityGroupListener.
func (a AvailabilityGroupListener) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	a.ProxyResource.marshalInternal(objectMap)
	populate(objectMap, "properties", a.Properties)
	return json.Marshal(objectMap)
}

// AvailabilityGroupListenerListResult - A list of availability group listeners.
type AvailabilityGroupListenerListResult struct {
	// READ-ONLY; Link to retrieve next page of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; Array of results.
	Value []*AvailabilityGroupListener `json:"value,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type AvailabilityGroupListenerListResult.
func (a AvailabilityGroupListenerListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", a.NextLink)
	populate(objectMap, "value", a.Value)
	return json.Marshal(objectMap)
}

// AvailabilityGroupListenerProperties - The properties of an availability group listener.
type AvailabilityGroupListenerProperties struct {
	// Name of the availability group.
	AvailabilityGroupName *string `json:"availabilityGroupName,omitempty"`

	// Create a default availability group if it does not exist.
	CreateDefaultAvailabilityGroupIfNotExist *bool `json:"createDefaultAvailabilityGroupIfNotExist,omitempty"`

	// List of load balancer configurations for an availability group listener.
	LoadBalancerConfigurations []*LoadBalancerConfiguration `json:"loadBalancerConfigurations,omitempty"`

	// Listener port.
	Port *int32 `json:"port,omitempty"`

	// READ-ONLY; Provisioning state to track the async operation status.
	ProvisioningState *string `json:"provisioningState,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type AvailabilityGroupListenerProperties.
func (a AvailabilityGroupListenerProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "availabilityGroupName", a.AvailabilityGroupName)
	populate(objectMap, "createDefaultAvailabilityGroupIfNotExist", a.CreateDefaultAvailabilityGroupIfNotExist)
	populate(objectMap, "loadBalancerConfigurations", a.LoadBalancerConfigurations)
	populate(objectMap, "port", a.Port)
	populate(objectMap, "provisioningState", a.ProvisioningState)
	return json.Marshal(objectMap)
}

// AvailabilityGroupListenersBeginCreateOrUpdateOptions contains the optional parameters for the AvailabilityGroupListeners.BeginCreateOrUpdate method.
type AvailabilityGroupListenersBeginCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// AvailabilityGroupListenersBeginDeleteOptions contains the optional parameters for the AvailabilityGroupListeners.BeginDelete method.
type AvailabilityGroupListenersBeginDeleteOptions struct {
	// placeholder for future optional parameters
}

// AvailabilityGroupListenersGetOptions contains the optional parameters for the AvailabilityGroupListeners.Get method.
type AvailabilityGroupListenersGetOptions struct {
	// placeholder for future optional parameters
}

// AvailabilityGroupListenersListByGroupOptions contains the optional parameters for the AvailabilityGroupListeners.ListByGroup method.
type AvailabilityGroupListenersListByGroupOptions struct {
	// placeholder for future optional parameters
}

// KeyVaultCredentialSettings - Configure your SQL virtual machine to be able to connect to the Azure Key Vault service.
type KeyVaultCredentialSettings struct {
	// Azure Key Vault url.
	AzureKeyVaultURL *string `json:"azureKeyVaultUrl,omitempty"`

	// Credential name.
	CredentialName *string `json:"credentialName,omitempty"`

	// Enable or disable key vault credential setting.
	Enable *bool `json:"enable,omitempty"`

	// Service principal name to access key vault.
	ServicePrincipalName *string `json:"servicePrincipalName,omitempty"`

	// Service principal name secret to access key vault.
	ServicePrincipalSecret *string `json:"servicePrincipalSecret,omitempty"`
}

// LoadBalancerConfiguration - A load balancer configuration for an availability group listener.
type LoadBalancerConfiguration struct {
	// Resource id of the load balancer.
	LoadBalancerResourceID *string `json:"loadBalancerResourceId,omitempty"`

	// Private IP address.
	PrivateIPAddress *PrivateIPAddress `json:"privateIpAddress,omitempty"`

	// Probe port.
	ProbePort *int32 `json:"probePort,omitempty"`

	// Resource id of the public IP.
	PublicIPAddressResourceID *string `json:"publicIpAddressResourceId,omitempty"`

	// List of the SQL virtual machine instance resource id's that are enrolled into the availability group listener.
	SQLVirtualMachineInstances []*string `json:"sqlVirtualMachineInstances,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type LoadBalancerConfiguration.
func (l LoadBalancerConfiguration) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "loadBalancerResourceId", l.LoadBalancerResourceID)
	populate(objectMap, "privateIpAddress", l.PrivateIPAddress)
	populate(objectMap, "probePort", l.ProbePort)
	populate(objectMap, "publicIpAddressResourceId", l.PublicIPAddressResourceID)
	populate(objectMap, "sqlVirtualMachineInstances", l.SQLVirtualMachineInstances)
	return json.Marshal(objectMap)
}

// Operation - SQL REST API operation definition.
type Operation struct {
	// READ-ONLY; The localized display information for this particular operation / action.
	Display *OperationDisplay `json:"display,omitempty" azure:"ro"`

	// READ-ONLY; The name of the operation being performed on this particular object.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The intended executor of the operation.
	Origin *OperationOrigin `json:"origin,omitempty" azure:"ro"`

	// READ-ONLY; Additional descriptions for the operation.
	Properties map[string]map[string]interface{} `json:"properties,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type Operation.
func (o Operation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "display", o.Display)
	populate(objectMap, "name", o.Name)
	populate(objectMap, "origin", o.Origin)
	populate(objectMap, "properties", o.Properties)
	return json.Marshal(objectMap)
}

// OperationDisplay - Display metadata associated with the operation.
type OperationDisplay struct {
	// READ-ONLY; The localized friendly description for the operation.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly name for the operation.
	Operation *string `json:"operation,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly form of the resource provider name.
	Provider *string `json:"provider,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly form of the resource type related to this action/operation.
	Resource *string `json:"resource,omitempty" azure:"ro"`
}

// OperationListResult - Result of the request to list SQL operations.
type OperationListResult struct {
	// READ-ONLY; Link to retrieve next page of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; Array of results.
	Value []*Operation `json:"value,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type OperationListResult.
func (o OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// OperationsListOptions contains the optional parameters for the Operations.List method.
type OperationsListOptions struct {
	// placeholder for future optional parameters
}

// PrivateIPAddress - A private IP address bound to the availability group listener.
type PrivateIPAddress struct {
	// Private IP address bound to the availability group listener.
	IPAddress *string `json:"ipAddress,omitempty"`

	// Subnet used to include private IP.
	SubnetResourceID *string `json:"subnetResourceId,omitempty"`
}

// ProxyResource - ARM proxy resource.
type ProxyResource struct {
	Resource
}

func (p ProxyResource) marshalInternal(objectMap map[string]interface{}) {
	p.Resource.marshalInternal(objectMap)
}

// Resource - ARM resource.
type Resource struct {
	// READ-ONLY; Resource ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Resource name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	r.marshalInternal(objectMap)
	return json.Marshal(objectMap)
}

func (r Resource) marshalInternal(objectMap map[string]interface{}) {
	populate(objectMap, "id", r.ID)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "type", r.Type)
}

// ResourceIdentity - Azure Active Directory identity configuration for a resource.
type ResourceIdentity struct {
	// The identity type. Set this to 'SystemAssigned' in order to automatically create and assign an Azure Active Directory principal for the resource.
	Type *IdentityType `json:"type,omitempty"`

	// READ-ONLY; The Azure Active Directory principal id.
	PrincipalID *string `json:"principalId,omitempty" azure:"ro"`

	// READ-ONLY; The Azure Active Directory tenant id.
	TenantID *string `json:"tenantId,omitempty" azure:"ro"`
}

// SQLConnectivityUpdateSettings - Set the access level and network port settings for SQL Server.
type SQLConnectivityUpdateSettings struct {
	// SQL Server connectivity option.
	ConnectivityType *ConnectivityType `json:"connectivityType,omitempty"`

	// SQL Server port.
	Port *int32 `json:"port,omitempty"`

	// SQL Server sysadmin login password.
	SQLAuthUpdatePassword *string `json:"sqlAuthUpdatePassword,omitempty"`

	// SQL Server sysadmin login to create.
	SQLAuthUpdateUserName *string `json:"sqlAuthUpdateUserName,omitempty"`
}

// SQLStorageSettings - Set disk storage settings for SQL Server.
type SQLStorageSettings struct {
	// SQL Server default file path
	DefaultFilePath *string `json:"defaultFilePath,omitempty"`

	// Logical Unit Numbers for the disks.
	Luns []*int32 `json:"luns,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type SQLStorageSettings.
func (s SQLStorageSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "defaultFilePath", s.DefaultFilePath)
	populate(objectMap, "luns", s.Luns)
	return json.Marshal(objectMap)
}

// SQLStorageUpdateSettings - Set disk storage settings for SQL Server.
type SQLStorageUpdateSettings struct {
	// Disk configuration to apply to SQL Server.
	DiskConfigurationType *DiskConfigurationType `json:"diskConfigurationType,omitempty"`

	// Virtual machine disk count.
	DiskCount *int32 `json:"diskCount,omitempty"`

	// Device id of the first disk to be updated.
	StartingDeviceID *int32 `json:"startingDeviceId,omitempty"`
}

// SQLVirtualMachine - A SQL virtual machine.
type SQLVirtualMachine struct {
	TrackedResource
	// Azure Active Directory identity of the server.
	Identity *ResourceIdentity `json:"identity,omitempty"`

	// Resource properties.
	Properties *SQLVirtualMachineProperties `json:"properties,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type SQLVirtualMachine.
func (s SQLVirtualMachine) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	s.TrackedResource.marshalInternal(objectMap)
	populate(objectMap, "identity", s.Identity)
	populate(objectMap, "properties", s.Properties)
	return json.Marshal(objectMap)
}

// SQLVirtualMachineGroup - A SQL virtual machine group.
type SQLVirtualMachineGroup struct {
	TrackedResource
	// Resource properties.
	Properties *SQLVirtualMachineGroupProperties `json:"properties,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type SQLVirtualMachineGroup.
func (s SQLVirtualMachineGroup) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	s.TrackedResource.marshalInternal(objectMap)
	populate(objectMap, "properties", s.Properties)
	return json.Marshal(objectMap)
}

// SQLVirtualMachineGroupListResult - A list of SQL virtual machine groups.
type SQLVirtualMachineGroupListResult struct {
	// READ-ONLY; Link to retrieve next page of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; Array of results.
	Value []*SQLVirtualMachineGroup `json:"value,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type SQLVirtualMachineGroupListResult.
func (s SQLVirtualMachineGroupListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// SQLVirtualMachineGroupProperties - The properties of a SQL virtual machine group.
type SQLVirtualMachineGroupProperties struct {
	// SQL image offer. Examples may include SQL2016-WS2016, SQL2017-WS2016.
	SQLImageOffer *string `json:"sqlImageOffer,omitempty"`

	// SQL image sku.
	SQLImageSKU *SQLVMGroupImageSKU `json:"sqlImageSku,omitempty"`

	// Cluster Active Directory domain profile.
	WsfcDomainProfile *WsfcDomainProfile `json:"wsfcDomainProfile,omitempty"`

	// READ-ONLY; Cluster type.
	ClusterConfiguration *ClusterConfiguration `json:"clusterConfiguration,omitempty" azure:"ro"`

	// READ-ONLY; Type of cluster manager: Windows Server Failover Cluster (WSFC), implied by the scale type of the group and the OS type.
	ClusterManagerType *ClusterManagerType `json:"clusterManagerType,omitempty" azure:"ro"`

	// READ-ONLY; Provisioning state to track the async operation status.
	ProvisioningState *string `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; Scale type.
	ScaleType *ScaleType `json:"scaleType,omitempty" azure:"ro"`
}

// SQLVirtualMachineGroupUpdate - An update to a SQL virtual machine group.
type SQLVirtualMachineGroupUpdate struct {
	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type SQLVirtualMachineGroupUpdate.
func (s SQLVirtualMachineGroupUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "tags", s.Tags)
	return json.Marshal(objectMap)
}

// SQLVirtualMachineGroupsBeginCreateOrUpdateOptions contains the optional parameters for the SQLVirtualMachineGroups.BeginCreateOrUpdate method.
type SQLVirtualMachineGroupsBeginCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// SQLVirtualMachineGroupsBeginDeleteOptions contains the optional parameters for the SQLVirtualMachineGroups.BeginDelete method.
type SQLVirtualMachineGroupsBeginDeleteOptions struct {
	// placeholder for future optional parameters
}

// SQLVirtualMachineGroupsBeginUpdateOptions contains the optional parameters for the SQLVirtualMachineGroups.BeginUpdate method.
type SQLVirtualMachineGroupsBeginUpdateOptions struct {
	// placeholder for future optional parameters
}

// SQLVirtualMachineGroupsGetOptions contains the optional parameters for the SQLVirtualMachineGroups.Get method.
type SQLVirtualMachineGroupsGetOptions struct {
	// placeholder for future optional parameters
}

// SQLVirtualMachineGroupsListByResourceGroupOptions contains the optional parameters for the SQLVirtualMachineGroups.ListByResourceGroup method.
type SQLVirtualMachineGroupsListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// SQLVirtualMachineGroupsListOptions contains the optional parameters for the SQLVirtualMachineGroups.List method.
type SQLVirtualMachineGroupsListOptions struct {
	// placeholder for future optional parameters
}

// SQLVirtualMachineListResult - A list of SQL virtual machines.
type SQLVirtualMachineListResult struct {
	// READ-ONLY; Link to retrieve next page of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; Array of results.
	Value []*SQLVirtualMachine `json:"value,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type SQLVirtualMachineListResult.
func (s SQLVirtualMachineListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// SQLVirtualMachineProperties - The SQL virtual machine properties.
type SQLVirtualMachineProperties struct {
	// Auto backup settings for SQL Server.
	AutoBackupSettings *AutoBackupSettings `json:"autoBackupSettings,omitempty"`

	// Auto patching settings for applying critical security updates to SQL virtual machine.
	AutoPatchingSettings *AutoPatchingSettings `json:"autoPatchingSettings,omitempty"`

	// Key vault credential settings.
	KeyVaultCredentialSettings *KeyVaultCredentialSettings `json:"keyVaultCredentialSettings,omitempty"`

	// SQL image offer. Examples include SQL2016-WS2016, SQL2017-WS2016.
	SQLImageOffer *string `json:"sqlImageOffer,omitempty"`

	// SQL Server edition type.
	SQLImageSKU *SQLImageSKU `json:"sqlImageSku,omitempty"`

	// SQL Server Management type.
	SQLManagement *SQLManagementMode `json:"sqlManagement,omitempty"`

	// SQL Server license type.
	SQLServerLicenseType *SQLServerLicenseType `json:"sqlServerLicenseType,omitempty"`

	// ARM resource id of the SQL virtual machine group this SQL virtual machine is or will be part of.
	SQLVirtualMachineGroupResourceID *string `json:"sqlVirtualMachineGroupResourceId,omitempty"`

	// SQL Server configuration management settings.
	ServerConfigurationsManagementSettings *ServerConfigurationsManagementSettings `json:"serverConfigurationsManagementSettings,omitempty"`

	// Storage Configuration Settings.
	StorageConfigurationSettings *StorageConfigurationSettings `json:"storageConfigurationSettings,omitempty"`

	// ARM Resource id of underlying virtual machine created from SQL marketplace image.
	VirtualMachineResourceID *string `json:"virtualMachineResourceId,omitempty"`

	// Domain credentials for setting up Windows Server Failover Cluster for SQL availability group.
	WsfcDomainCredentials *WsfcDomainCredentials `json:"wsfcDomainCredentials,omitempty"`

	// READ-ONLY; Provisioning state to track the async operation status.
	ProvisioningState *string `json:"provisioningState,omitempty" azure:"ro"`
}

// SQLVirtualMachineUpdate - An update to a SQL virtual machine.
type SQLVirtualMachineUpdate struct {
	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type SQLVirtualMachineUpdate.
func (s SQLVirtualMachineUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "tags", s.Tags)
	return json.Marshal(objectMap)
}

// SQLVirtualMachinesBeginCreateOrUpdateOptions contains the optional parameters for the SQLVirtualMachines.BeginCreateOrUpdate method.
type SQLVirtualMachinesBeginCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// SQLVirtualMachinesBeginDeleteOptions contains the optional parameters for the SQLVirtualMachines.BeginDelete method.
type SQLVirtualMachinesBeginDeleteOptions struct {
	// placeholder for future optional parameters
}

// SQLVirtualMachinesBeginUpdateOptions contains the optional parameters for the SQLVirtualMachines.BeginUpdate method.
type SQLVirtualMachinesBeginUpdateOptions struct {
	// placeholder for future optional parameters
}

// SQLVirtualMachinesGetOptions contains the optional parameters for the SQLVirtualMachines.Get method.
type SQLVirtualMachinesGetOptions struct {
	// The child resources to include in the response.
	Expand *string
}

// SQLVirtualMachinesListByResourceGroupOptions contains the optional parameters for the SQLVirtualMachines.ListByResourceGroup method.
type SQLVirtualMachinesListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// SQLVirtualMachinesListBySQLVMGroupOptions contains the optional parameters for the SQLVirtualMachines.ListBySQLVMGroup method.
type SQLVirtualMachinesListBySQLVMGroupOptions struct {
	// placeholder for future optional parameters
}

// SQLVirtualMachinesListOptions contains the optional parameters for the SQLVirtualMachines.List method.
type SQLVirtualMachinesListOptions struct {
	// placeholder for future optional parameters
}

// SQLWorkloadTypeUpdateSettings - Set workload type to optimize storage for SQL Server.
type SQLWorkloadTypeUpdateSettings struct {
	// SQL Server workload type.
	SQLWorkloadType *SQLWorkloadType `json:"sqlWorkloadType,omitempty"`
}

// ServerConfigurationsManagementSettings - Set the connectivity, storage and workload settings.
type ServerConfigurationsManagementSettings struct {
	// Additional SQL feature settings.
	AdditionalFeaturesServerConfigurations *AdditionalFeaturesServerConfigurations `json:"additionalFeaturesServerConfigurations,omitempty"`

	// SQL connectivity type settings.
	SQLConnectivityUpdateSettings *SQLConnectivityUpdateSettings `json:"sqlConnectivityUpdateSettings,omitempty"`

	// SQL storage update settings.
	SQLStorageUpdateSettings *SQLStorageUpdateSettings `json:"sqlStorageUpdateSettings,omitempty"`

	// SQL workload type settings.
	SQLWorkloadTypeUpdateSettings *SQLWorkloadTypeUpdateSettings `json:"sqlWorkloadTypeUpdateSettings,omitempty"`
}

// StorageConfigurationSettings - Storage Configurations for SQL Data, Log and TempDb.
type StorageConfigurationSettings struct {
	// Disk configuration to apply to SQL Server.
	DiskConfigurationType *DiskConfigurationType `json:"diskConfigurationType,omitempty"`

	// SQL Server Data Storage Settings.
	SQLDataSettings *SQLStorageSettings `json:"sqlDataSettings,omitempty"`

	// SQL Server Log Storage Settings.
	SQLLogSettings *SQLStorageSettings `json:"sqlLogSettings,omitempty"`

	// SQL Server TempDb Storage Settings.
	SQLTempDbSettings *SQLStorageSettings `json:"sqlTempDbSettings,omitempty"`

	// Storage workload type.
	StorageWorkloadType *StorageWorkloadType `json:"storageWorkloadType,omitempty"`
}

// TrackedResource - ARM tracked top level resource.
type TrackedResource struct {
	Resource
	// REQUIRED; Resource location.
	Location *string `json:"location,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type TrackedResource.
func (t TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	t.marshalInternal(objectMap)
	return json.Marshal(objectMap)
}

func (t TrackedResource) marshalInternal(objectMap map[string]interface{}) {
	t.Resource.marshalInternal(objectMap)
	populate(objectMap, "location", t.Location)
	populate(objectMap, "tags", t.Tags)
}

// WsfcDomainCredentials - Domain credentials for setting up Windows Server Failover Cluster for SQL availability group.
type WsfcDomainCredentials struct {
	// Cluster bootstrap account password.
	ClusterBootstrapAccountPassword *string `json:"clusterBootstrapAccountPassword,omitempty"`

	// Cluster operator account password.
	ClusterOperatorAccountPassword *string `json:"clusterOperatorAccountPassword,omitempty"`

	// SQL service account password.
	SQLServiceAccountPassword *string `json:"sqlServiceAccountPassword,omitempty"`
}

// WsfcDomainProfile - Active Directory account details to operate Windows Server Failover Cluster.
type WsfcDomainProfile struct {
	// Account name used for creating cluster (at minimum needs permissions to 'Create Computer Objects' in domain).
	ClusterBootstrapAccount *string `json:"clusterBootstrapAccount,omitempty"`

	// Account name used for operating cluster i.e. will be part of administrators group on all the participating virtual machines in the cluster.
	ClusterOperatorAccount *string `json:"clusterOperatorAccount,omitempty"`

	// Fully qualified name of the domain.
	DomainFqdn *string `json:"domainFqdn,omitempty"`

	// Optional path for fileshare witness.
	FileShareWitnessPath *string `json:"fileShareWitnessPath,omitempty"`

	// Organizational Unit path in which the nodes and cluster will be present.
	OuPath *string `json:"ouPath,omitempty"`

	// Account name under which SQL service will run on all participating SQL virtual machines in the cluster.
	SQLServiceAccount *string `json:"sqlServiceAccount,omitempty"`

	// Primary key of the witness storage account.
	StorageAccountPrimaryKey *string `json:"storageAccountPrimaryKey,omitempty"`

	// Fully qualified ARM resource id of the witness storage account.
	StorageAccountURL *string `json:"storageAccountUrl,omitempty"`
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}
