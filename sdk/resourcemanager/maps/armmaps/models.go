//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmaps

import "time"

// Account - An Azure resource which represents access to a suite of Maps REST APIs.
type Account struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; The SKU of this account.
	SKU *SKU

	// Managed service identity (system assigned and/or user assigned identities)
	Identity *ManagedServiceIdentity

	// Get or Set Kind property.
	Kind *Kind

	// The map account properties.
	Properties *AccountProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// AccountKeys - The set of keys which can be used to access the Maps REST APIs. Two keys are provided for key rotation without
// interruption.
type AccountKeys struct {
	// READ-ONLY; The primary key for accessing the Maps REST APIs.
	PrimaryKey *string

	// READ-ONLY; The last updated date and time of the primary key.
	PrimaryKeyLastUpdated *string

	// READ-ONLY; The secondary key for accessing the Maps REST APIs.
	SecondaryKey *string

	// READ-ONLY; The last updated date and time of the secondary key.
	SecondaryKeyLastUpdated *string
}

// AccountProperties - Additional Map account properties
type AccountProperties struct {
	// Specifies CORS rules for the Blob service. You can include up to five CorsRule elements in the request. If no CorsRule
	// elements are included in the request body, all CORS rules will be deleted, and
	// CORS will be disabled for the Blob service.
	Cors *CorsRules

	// Allows toggle functionality on Azure Policy to disable Azure Maps local authentication support. This will disable Shared
	// Keys and Shared Access Signature Token authentication from any usage.
	DisableLocalAuth *bool

	// (Optional) Discouraged to include in resource definition. Only needed where it is possible to disable platform (AKA infrastructure)
	// encryption. Azure SQL TDE is an example of this. Values are enabled
	// and disabled.
	Encryption *Encryption

	// The array of associated resources to the Map account. Linked resource in the array cannot individually update, you must
	// update all linked resources in the array together. These resources may be used
	// on operations on the Azure Maps REST API. Access is controlled by the Map Account Managed Identity(s) permissions to those
	// resource(s).
	LinkedResources []*LinkedResource

	// READ-ONLY; The provisioning state of the Map account resource, Account updates can only be performed on terminal states.
	// Terminal states: Succeeded and Failed
	ProvisioningState *string

	// READ-ONLY; A unique identifier for the maps account
	UniqueID *string
}

// AccountSasParameters - Parameters used to create an account Shared Access Signature (SAS) token. The REST API access control
// is provided by Azure Maps Role Based Access (RBAC) identity and access.
type AccountSasParameters struct {
	// REQUIRED; The date time offset of when the token validity expires. For example "2017-05-24T10:42:03.1567373Z". Maximum
	// duration allowed is 24 hours between start and expiry.
	Expiry *string

	// REQUIRED; Required parameter which represents the desired maximum request per second to allowed for the given SAS token.
	// This does not guarantee perfect accuracy in measurements but provides application safe
	// guards of abuse with eventual enforcement.
	MaxRatePerSecond *int32

	// REQUIRED; The principal Id also known as the object Id of a User Assigned Managed Identity currently assigned to the Map
	// Account. To assign a Managed Identity of the account, use operation Create or Update an
	// assign a User Assigned Identity resource Id.
	PrincipalID *string

	// REQUIRED; The Map account key to use for signing. Picking primaryKey or secondaryKey will use the Map account Shared Keys,
	// and using managedIdentity will use the auto-renewed private key to sign the SAS.
	SigningKey *SigningKey

	// REQUIRED; The date time offset of when the token validity begins. For example "2017-05-24T10:42:03.1567373Z". Maximum duration
	// allowed is 24 hours between start and expiry.
	Start *string

	// Optional, allows control of which region locations are permitted access to Azure Maps REST APIs with the SAS token. Example:
	// "eastus", "westus2". Omitting this parameter will allow all region
	// locations to be accessible.
	Regions []*string
}

// AccountSasToken - A new Sas token which can be used to access the Maps REST APIs and is controlled by the specified Managed
// identity permissions on Azure (IAM) Role Based Access Control.
type AccountSasToken struct {
	// READ-ONLY; The shared access signature access token.
	AccountSasToken *string
}

// AccountUpdateParameters - Parameters used to update an existing Maps Account.
type AccountUpdateParameters struct {
	// Managed service identity (system assigned and/or user assigned identities)
	Identity *ManagedServiceIdentity

	// Get or Set Kind property.
	Kind *Kind

	// The map account properties.
	Properties *AccountProperties

	// The SKU of this account.
	SKU *SKU

	// Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this
	// resource (across resource groups). A maximum of 15 tags can be provided for a
	// resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags map[string]*string
}

// Accounts - A list of Maps Accounts.
type Accounts struct {
	// URL client should use to fetch the next page (per server side paging). It's null for now, added for future use.
	NextLink *string

	// READ-ONLY; a Maps Account.
	Value []*Account
}

// CorsRule - Specifies a CORS rule for the Map Account.
type CorsRule struct {
	// REQUIRED; Required if CorsRule element is present. A list of origin domains that will be allowed via CORS, or "*" to allow
	// all domains
	AllowedOrigins []*string
}

// CorsRules - Sets the CORS rules. You can include up to five CorsRule elements in the request.
type CorsRules struct {
	// The list of CORS rules. You can include up to five CorsRule elements in the request.
	CorsRules []*CorsRule
}

// Creator - An Azure resource which represents Maps Creator product and provides ability to manage private location data.
type Creator struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; The Creator resource properties.
	Properties *CreatorProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// CreatorList - A list of Creator resources.
type CreatorList struct {
	// URL client should use to fetch the next page (per server side paging). It's null for now, added for future use.
	NextLink *string

	// READ-ONLY; a Creator account.
	Value []*Creator
}

// CreatorProperties - Creator resource properties
type CreatorProperties struct {
	// REQUIRED; The storage units to be allocated. Integer values from 1 to 100, inclusive.
	StorageUnits *int32

	// READ-ONLY; The state of the resource provisioning, terminal states: Succeeded, Failed, Canceled
	ProvisioningState *string
}

// CreatorUpdateParameters - Parameters used to update an existing Creator resource.
type CreatorUpdateParameters struct {
	// Creator resource properties.
	Properties *CreatorProperties

	// Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this
	// resource (across resource groups). A maximum of 15 tags can be provided for a
	// resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags map[string]*string
}

// CustomerManagedKeyEncryption - All Customer-managed key encryption properties for the resource.
type CustomerManagedKeyEncryption struct {
	// All identity configuration for Customer-managed key settings defining which identity should be used to auth to Key Vault.
	KeyEncryptionKeyIdentity *CustomerManagedKeyEncryptionKeyIdentity

	// key encryption key Url, versioned or non-versioned. Ex: https://contosovault.vault.azure.net/keys/contosokek/562a4bb76b524a1493a6afe8e536ee78
	// or https://contosovault.vault.azure.net/keys/contosokek.
	KeyEncryptionKeyURL *string
}

// CustomerManagedKeyEncryptionKeyIdentity - All identity configuration for Customer-managed key settings defining which identity
// should be used to auth to Key Vault.
type CustomerManagedKeyEncryptionKeyIdentity struct {
	// delegated identity to use for accessing key encryption key Url. Ex: /subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups//providers/Microsoft.ManagedIdentity/userAssignedIdentities/myId.
	// Mutually exclusive with identityType systemAssignedIdentity and userAssignedIdentity - internal use only.
	DelegatedIdentityClientID *string

	// Values can be systemAssignedIdentity or userAssignedIdentity
	IdentityType *IdentityType

	// user assigned identity to use for accessing key encryption key Url. Ex: /subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/
	// /providers/Microsoft.ManagedIdentity/userAssignedIdentities/myId. Mutually exclusive with identityType systemAssignedIdentity
	// and delegatedResourceIdentity.
	UserAssignedIdentityResourceID *string
}

// Dimension of map account, for example API Category, Api Name, Result Type, and Response Code.
type Dimension struct {
	// Display name of dimension.
	DisplayName *string

	// Internal metric name of the dimension.
	InternalMetricName *string

	// Internal name of the dimension.
	InternalName *string

	// Display name of dimension.
	Name *string

	// Source Mdm Namespace of the dimension.
	SourceMdmNamespace *string

	// Flag to indicate exporting to Azure Monitor.
	ToBeExportedToShoebox *bool
}

// Encryption - (Optional) Discouraged to include in resource definition. Only needed where it is possible to disable platform
// (AKA infrastructure) encryption. Azure SQL TDE is an example of this. Values are enabled
// and disabled.
type Encryption struct {
	// All Customer-managed key encryption properties for the resource.
	CustomerManagedKeyEncryption *CustomerManagedKeyEncryption

	// Values are enabled and disabled.
	InfrastructureEncryption *InfrastructureEncryption
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

// KeySpecification - Whether the operation refers to the primary or secondary key.
type KeySpecification struct {
	// REQUIRED; Whether the operation refers to the primary or secondary key.
	KeyType *KeyType
}

// LinkedResource - Linked resource is reference to a resource deployed in an Azure subscription, add the linked resource
// uniqueName value as an optional parameter for operations on Azure Maps Geospatial REST APIs.
type LinkedResource struct {
	// REQUIRED; ARM resource id in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/accounts/{storageName}'.
	ID *string

	// REQUIRED; A provided name which uniquely identifies the linked resource.
	UniqueName *string
}

// ManagedServiceIdentity - Managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentity struct {
	// REQUIRED; Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
	Type *ManagedServiceIdentityType

	// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM
	// resource ids in the form:
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}.
	// The dictionary values can be empty objects ({}) in
	// requests.
	UserAssignedIdentities map[string]*UserAssignedIdentity

	// READ-ONLY; The service principal ID of the system assigned identity. This property will only be provided for a system assigned
	// identity.
	PrincipalID *string

	// READ-ONLY; The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
	TenantID *string
}

// MetricSpecification - Metric specification of operation.
type MetricSpecification struct {
	// Aggregation type could be Average.
	AggregationType *string

	// The category this metric specification belong to, could be Capacity.
	Category *string

	// Dimensions of map account.
	Dimensions []*Dimension

	// Display description of metric specification.
	DisplayDescription *string

	// Display name of metric specification.
	DisplayName *string

	// The property to decide fill gap with zero or not.
	FillGapWithZero *bool

	// Internal metric name.
	InternalMetricName *string

	// Lock aggregation type for metrics.
	LockAggregationType *string

	// Name of metric specification.
	Name *string

	// Account Resource Id.
	ResourceIDDimensionNameOverride *string

	// Source metrics account.
	SourceMdmAccount *string

	// Metrics namespace.
	SourceMdmNamespace *string

	// Allowed aggregation types for metrics.
	SupportedAggregationTypes *string

	// Unit could be Count.
	Unit *string
}

// OperationDetail - Operation detail payload
type OperationDetail struct {
	// Display of the operation
	Display *OperationDisplay

	// Indicates whether the operation is a data action
	IsDataAction *bool

	// Name of the operation
	Name *string

	// Origin of the operation
	Origin *string

	// Properties of the operation
	Properties *OperationProperties
}

// OperationDisplay - Operation display payload
type OperationDisplay struct {
	// Localized friendly description for the operation
	Description *string

	// Localized friendly name for the operation
	Operation *string

	// Resource provider of the operation
	Provider *string

	// Resource of the operation
	Resource *string
}

// OperationProperties - Properties of operation, include metric specifications.
type OperationProperties struct {
	// One property of operation, include metric specifications.
	ServiceSpecification *ServiceSpecification
}

// Operations - The set of operations available for Maps.
type Operations struct {
	// URL client should use to fetch the next page (per server side paging). It's null for now, added for future use.
	NextLink *string

	// READ-ONLY; An operation available for Maps.
	Value []*OperationDetail
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

// SKU - The SKU of the Maps Account.
type SKU struct {
	// REQUIRED; The name of the SKU, in standard format (such as S0).
	Name *Name

	// READ-ONLY; Gets the sku tier. This is based on the SKU name.
	Tier *string
}

// ServiceSpecification - One property of operation, include metric specifications.
type ServiceSpecification struct {
	// Metric specifications of operation.
	MetricSpecifications []*MetricSpecification
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

// UserAssignedIdentity - User assigned identity properties
type UserAssignedIdentity struct {
	// READ-ONLY; The client ID of the assigned identity.
	ClientID *string

	// READ-ONLY; The principal ID of the assigned identity.
	PrincipalID *string
}
