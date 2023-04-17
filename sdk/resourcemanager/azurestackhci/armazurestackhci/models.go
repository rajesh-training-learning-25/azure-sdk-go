//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armazurestackhci

import "time"

// ArcConnectivityProperties - Connectivity related configuration required by arc server.
type ArcConnectivityProperties struct {
	// True indicates ARC connectivity is enabled
	Enabled *bool
}

// ArcIdentityResponse - ArcIdentity details.
type ArcIdentityResponse struct {
	// READ-ONLY; ArcIdentity properties.
	Properties *ArcIdentityResponseProperties
}

type ArcIdentityResponseProperties struct {
	ArcApplicationClientID      *string
	ArcApplicationObjectID      *string
	ArcApplicationTenantID      *string
	ArcServicePrincipalObjectID *string
}

// ArcSetting details.
type ArcSetting struct {
	// ArcSetting properties.
	Properties *ArcSettingProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; System data of ArcSetting resource
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ArcSettingList - List of ArcSetting proxy resources for the HCI cluster.
type ArcSettingList struct {
	// READ-ONLY; Link to the next set of results.
	NextLink *string

	// READ-ONLY; List of ArcSetting proxy resources.
	Value []*ArcSetting
}

// ArcSettingProperties - ArcSetting properties.
type ArcSettingProperties struct {
	// App id of arc AAD identity.
	ArcApplicationClientID *string

	// Object id of arc AAD identity.
	ArcApplicationObjectID *string

	// Tenant id of arc AAD identity.
	ArcApplicationTenantID *string

	// The resource group that hosts the Arc agents, ie. Hybrid Compute Machine resources.
	ArcInstanceResourceGroup *string

	// Object id of arc AAD service principal.
	ArcServicePrincipalObjectID *string

	// contains connectivity related configuration for ARC resources
	ConnectivityProperties any

	// READ-ONLY; Aggregate state of Arc agent across the nodes in this HCI cluster.
	AggregateState *ArcSettingAggregateState

	// READ-ONLY; State of Arc agent in each of the nodes.
	PerNodeDetails []*PerNodeState

	// READ-ONLY; Provisioning state of the ArcSetting proxy resource.
	ProvisioningState *ProvisioningState
}

// ArcSettingsClientBeginCreateIdentityOptions contains the optional parameters for the ArcSettingsClient.BeginCreateIdentity
// method.
type ArcSettingsClientBeginCreateIdentityOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ArcSettingsClientBeginDeleteOptions contains the optional parameters for the ArcSettingsClient.BeginDelete method.
type ArcSettingsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ArcSettingsClientCreateOptions contains the optional parameters for the ArcSettingsClient.Create method.
type ArcSettingsClientCreateOptions struct {
	// placeholder for future optional parameters
}

// ArcSettingsClientGeneratePasswordOptions contains the optional parameters for the ArcSettingsClient.GeneratePassword method.
type ArcSettingsClientGeneratePasswordOptions struct {
	// placeholder for future optional parameters
}

// ArcSettingsClientGetOptions contains the optional parameters for the ArcSettingsClient.Get method.
type ArcSettingsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ArcSettingsClientListByClusterOptions contains the optional parameters for the ArcSettingsClient.NewListByClusterPager
// method.
type ArcSettingsClientListByClusterOptions struct {
	// placeholder for future optional parameters
}

// ArcSettingsClientUpdateOptions contains the optional parameters for the ArcSettingsClient.Update method.
type ArcSettingsClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// ArcSettingsPatch - ArcSetting details to update.
type ArcSettingsPatch struct {
	// ArcSettings properties.
	Properties *ArcSettingsPatchProperties

	// Resource tags.
	Tags map[string]*string
}

// ArcSettingsPatchProperties - ArcSettings properties.
type ArcSettingsPatchProperties struct {
	// contains connectivity related configuration for ARC resources
	ConnectivityProperties any
}

// Cluster details.
type Cluster struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// Cluster properties.
	Properties *ClusterProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; System data of Cluster resource
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ClusterDesiredProperties - Desired properties of the cluster.
type ClusterDesiredProperties struct {
	// Desired level of diagnostic data emitted by the cluster.
	DiagnosticLevel *DiagnosticLevel

	// Desired state of Windows Server Subscription.
	WindowsServerSubscription *WindowsServerSubscription
}

// ClusterIdentityResponse - Cluster Identity details.
type ClusterIdentityResponse struct {
	// READ-ONLY; Cluster identity properties.
	Properties *ClusterIdentityResponseProperties
}

type ClusterIdentityResponseProperties struct {
	AADApplicationObjectID      *string
	AADClientID                 *string
	AADServicePrincipalObjectID *string
	AADTenantID                 *string
}

// ClusterList - List of clusters.
type ClusterList struct {
	// List of clusters.
	Value []*Cluster

	// READ-ONLY; Link to the next set of results.
	NextLink *string
}

// ClusterNode - Cluster node details.
type ClusterNode struct {
	// READ-ONLY; Number of physical cores on the cluster node.
	CoreCount *float32

	// READ-ONLY; Id of the node in the cluster.
	ID *float32

	// READ-ONLY; Manufacturer of the cluster node hardware.
	Manufacturer *string

	// READ-ONLY; Total available memory on the cluster node (in GiB).
	MemoryInGiB *float32

	// READ-ONLY; Model name of the cluster node hardware.
	Model *string

	// READ-ONLY; Name of the cluster node.
	Name *string

	// READ-ONLY; Operating system running on the cluster node.
	OSName *string

	// READ-ONLY; Version of the operating system running on the cluster node.
	OSVersion *string

	// READ-ONLY; Immutable id of the cluster node.
	SerialNumber *string

	// READ-ONLY; State of Windows Server Subscription.
	WindowsServerSubscription *WindowsServerSubscription
}

// ClusterPatch - Cluster details to update.
type ClusterPatch struct {
	// Cluster properties.
	Properties *ClusterPatchProperties

	// Resource tags.
	Tags map[string]*string
}

// ClusterPatchProperties - Cluster properties.
type ClusterPatchProperties struct {
	// App id of cluster AAD identity.
	AADClientID *string

	// Tenant id of cluster AAD identity.
	AADTenantID *string

	// Endpoint configured for management from the Azure portal
	CloudManagementEndpoint *string

	// Desired properties of the cluster.
	DesiredProperties *ClusterDesiredProperties
}

// ClusterProperties - Cluster properties.
type ClusterProperties struct {
	// Object id of cluster AAD identity.
	AADApplicationObjectID *string

	// App id of cluster AAD identity.
	AADClientID *string

	// Id of cluster identity service principal.
	AADServicePrincipalObjectID *string

	// Tenant id of cluster AAD identity.
	AADTenantID *string

	// Endpoint configured for management from the Azure portal.
	CloudManagementEndpoint *string

	// Desired properties of the cluster.
	DesiredProperties *ClusterDesiredProperties

	// READ-ONLY; Type of billing applied to the resource.
	BillingModel *string

	// READ-ONLY; Unique, immutable resource id.
	CloudID *string

	// READ-ONLY; Most recent billing meter timestamp.
	LastBillingTimestamp *time.Time

	// READ-ONLY; Most recent cluster sync timestamp.
	LastSyncTimestamp *time.Time

	// READ-ONLY; Provisioning state.
	ProvisioningState *ProvisioningState

	// READ-ONLY; First cluster sync timestamp.
	RegistrationTimestamp *time.Time

	// READ-ONLY; Properties reported by cluster agent.
	ReportedProperties *ClusterReportedProperties

	// READ-ONLY; Region specific DataPath Endpoint of the cluster.
	ServiceEndpoint *string

	// READ-ONLY; Status of the cluster agent.
	Status *Status

	// READ-ONLY; Number of days remaining in the trial period.
	TrialDaysRemaining *float32
}

// ClusterReportedProperties - Properties reported by cluster agent.
type ClusterReportedProperties struct {
	// Level of diagnostic data emitted by the cluster.
	DiagnosticLevel *DiagnosticLevel

	// READ-ONLY; Unique id generated by the on-prem cluster.
	ClusterID *string

	// READ-ONLY; Name of the on-prem cluster connected to this resource.
	ClusterName *string

	// READ-ONLY; Version of the cluster software.
	ClusterVersion *string

	// READ-ONLY; IMDS attestation status of the cluster.
	ImdsAttestation *ImdsAttestation

	// READ-ONLY; Last time the cluster reported the data.
	LastUpdated *time.Time

	// READ-ONLY; List of nodes reported by the cluster.
	Nodes []*ClusterNode
}

// ClustersClientBeginCreateIdentityOptions contains the optional parameters for the ClustersClient.BeginCreateIdentity method.
type ClustersClientBeginCreateIdentityOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ClustersClientBeginDeleteOptions contains the optional parameters for the ClustersClient.BeginDelete method.
type ClustersClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ClustersClientBeginUploadCertificateOptions contains the optional parameters for the ClustersClient.BeginUploadCertificate
// method.
type ClustersClientBeginUploadCertificateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ClustersClientCreateOptions contains the optional parameters for the ClustersClient.Create method.
type ClustersClientCreateOptions struct {
	// placeholder for future optional parameters
}

// ClustersClientGetOptions contains the optional parameters for the ClustersClient.Get method.
type ClustersClientGetOptions struct {
	// placeholder for future optional parameters
}

// ClustersClientListByResourceGroupOptions contains the optional parameters for the ClustersClient.NewListByResourceGroupPager
// method.
type ClustersClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// ClustersClientListBySubscriptionOptions contains the optional parameters for the ClustersClient.NewListBySubscriptionPager
// method.
type ClustersClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// ClustersClientUpdateOptions contains the optional parameters for the ClustersClient.Update method.
type ClustersClientUpdateOptions struct {
	// placeholder for future optional parameters
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

// Extension - Details of a particular extension in HCI Cluster.
type Extension struct {
	// Describes Machine Extension Properties.
	Properties *ExtensionProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; System data of Extension resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ExtensionList - List of Extensions in HCI cluster.
type ExtensionList struct {
	// READ-ONLY; Link to the next set of results.
	NextLink *string

	// READ-ONLY; List of Extensions in HCI cluster.
	Value []*Extension
}

// ExtensionParameters - Describes the properties of a Machine Extension. This object mirrors the definition in HybridCompute.
type ExtensionParameters struct {
	// Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed,
	// however, the extension will not upgrade minor versions unless redeployed, even
	// with this property set to true.
	AutoUpgradeMinorVersion *bool

	// How the extension handler should be forced to update even if the extension configuration has not changed.
	ForceUpdateTag *string

	// Protected settings (may contain secrets).
	ProtectedSettings any

	// The name of the extension handler publisher.
	Publisher *string

	// Json formatted public settings for the extension.
	Settings any

	// Specifies the type of the extension; an example is "CustomScriptExtension".
	Type *string

	// Specifies the version of the script handler.
	TypeHandlerVersion *string
}

// ExtensionProperties - Status of Arc Extension for a particular node in HCI Cluster.
type ExtensionProperties struct {
	// Parameters specific to this extension type.
	ExtensionParameters *ExtensionParameters

	// READ-ONLY; Aggregate state of Arc Extensions across the nodes in this HCI cluster.
	AggregateState *ExtensionAggregateState

	// READ-ONLY; State of Arc Extension in each of the nodes.
	PerNodeExtensionDetails []*PerNodeExtensionState

	// READ-ONLY; Provisioning state of the Extension proxy resource.
	ProvisioningState *ProvisioningState
}

// ExtensionsClientBeginCreateOptions contains the optional parameters for the ExtensionsClient.BeginCreate method.
type ExtensionsClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ExtensionsClientBeginDeleteOptions contains the optional parameters for the ExtensionsClient.BeginDelete method.
type ExtensionsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ExtensionsClientBeginUpdateOptions contains the optional parameters for the ExtensionsClient.BeginUpdate method.
type ExtensionsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ExtensionsClientGetOptions contains the optional parameters for the ExtensionsClient.Get method.
type ExtensionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ExtensionsClientListByArcSettingOptions contains the optional parameters for the ExtensionsClient.NewListByArcSettingPager
// method.
type ExtensionsClientListByArcSettingOptions struct {
	// placeholder for future optional parameters
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

type PasswordCredential struct {
	EndDateTime   *time.Time
	KeyID         *string
	SecretText    *string
	StartDateTime *time.Time
}

// PerNodeExtensionState - Status of Arc Extension for a particular node in HCI Cluster.
type PerNodeExtensionState struct {
	// READ-ONLY; Fully qualified resource ID for the particular Arc Extension on this node.
	Extension *string

	// READ-ONLY; Name of the node in HCI Cluster.
	Name *string

	// READ-ONLY; State of Arc Extension in this node.
	State *NodeExtensionState
}

// PerNodeState - Status of Arc agent for a particular node in HCI Cluster.
type PerNodeState struct {
	// READ-ONLY; Fully qualified resource ID for the Arc agent of this node.
	ArcInstance *string

	// READ-ONLY; Name of the Node in HCI Cluster
	Name *string

	// READ-ONLY; State of Arc agent in this node.
	State *NodeArcState
}

// ProxyResource - The resource model definition for a Azure Resource Manager proxy resource. It will not have tags and a
// location
type ProxyResource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

type RawCertificateData struct {
	Certificates []*string
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
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

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

type UploadCertificateRequest struct {
	Properties *RawCertificateData
}
