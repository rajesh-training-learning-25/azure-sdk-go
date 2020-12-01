// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkeyvault

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// An identity that have access to the key vault. All identities in the array must use the same tenant ID as the key vault's tenant ID.
type AccessPolicyEntry struct {
	// Application ID of the client making request on behalf of a principal
	ApplicationID *string `json:"applicationId,omitempty"`

	// The object ID of a user, service principal or security group in the Azure Active Directory tenant for the vault. The object ID must be unique for the
	// list of access policies.
	ObjectID *string `json:"objectId,omitempty"`

	// Permissions the identity has for keys, secrets and certificates.
	Permissions *Permissions `json:"permissions,omitempty"`

	// The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
	TenantID *string `json:"tenantId,omitempty"`
}

// The object attributes managed by the Azure Key Vault service.
type Attributes struct {
	// READ-ONLY; Creation time in seconds since 1970-01-01T00:00:00Z.
	Created *int64 `json:"created,omitempty" azure:"ro"`

	// Determines whether or not the object is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// Expiry date in seconds since 1970-01-01T00:00:00Z.
	Expires *int64 `json:"exp,omitempty"`

	// Not before date in seconds since 1970-01-01T00:00:00Z.
	NotBefore *int64 `json:"nbf,omitempty"`

	// READ-ONLY; The deletion recovery level currently in effect for the object. If it contains 'Purgeable', then the object can be permanently deleted by
	// a privileged user; otherwise, only the system can purge the
	// object at the end of the retention interval.
	RecoveryLevel *DeletionRecoveryLevel `json:"recoveryLevel,omitempty" azure:"ro"`

	// READ-ONLY; Last updated time in seconds since 1970-01-01T00:00:00Z.
	Updated *int64 `json:"updated,omitempty" azure:"ro"`
}

// The CheckNameAvailability operation response.
type CheckNameAvailabilityResult struct {
	// READ-ONLY; An error message explaining the Reason value in more detail.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; A boolean value that indicates whether the name is available for you to use. If true, the name is available. If false, the name has already
	// been taken or is invalid and cannot be used.
	NameAvailable *bool `json:"nameAvailable,omitempty" azure:"ro"`

	// READ-ONLY; The reason that a vault name could not be used. The Reason element is only returned if NameAvailable is false.
	Reason *Reason `json:"reason,omitempty" azure:"ro"`
}

// CheckNameAvailabilityResultResponse is the response envelope for operations that return a CheckNameAvailabilityResult type.
type CheckNameAvailabilityResultResponse struct {
	// The CheckNameAvailability operation response.
	CheckNameAvailabilityResult *CheckNameAvailabilityResult

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// An error response from Key Vault resource provider
type CloudError struct {
	// An error response from Key Vault resource provider
	InnerError *CloudErrorBody `json:"error,omitempty"`
}

// Error implements the error interface for type CloudError.
func (e CloudError) Error() string {
	msg := ""
	if e.InnerError != nil {
		msg += fmt.Sprintf("InnerError: %v\n", *e.InnerError)
	}
	if msg == "" {
		msg = "missing error info"
	}
	return msg
}

// An error response from Key Vault resource provider
type CloudErrorBody struct {
	// Error code. This is a mnemonic that can be consumed programmatically.
	Code *string `json:"code,omitempty"`

	// User friendly error message. The message is typically localized and may vary with service version.
	Message *string `json:"message,omitempty"`
}

// Deleted vault information with extended details.
type DeletedVault struct {
	// READ-ONLY; The resource ID for the deleted key vault.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the key vault.
	Name *string `json:"name,omitempty" azure:"ro"`

	// Properties of the vault
	Properties *DeletedVaultProperties `json:"properties,omitempty"`

	// READ-ONLY; The resource type of the key vault.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// List of vaults
type DeletedVaultListResult struct {
	// The URL to get the next set of deleted vaults.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of deleted vaults.
	Value *[]DeletedVault `json:"value,omitempty"`
}

// DeletedVaultListResultResponse is the response envelope for operations that return a DeletedVaultListResult type.
type DeletedVaultListResultResponse struct {
	// List of vaults
	DeletedVaultListResult *DeletedVaultListResult

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// Properties of the deleted vault.
type DeletedVaultProperties struct {
	// READ-ONLY; The deleted date.
	DeletionDate *time.Time `json:"deletionDate,omitempty" azure:"ro"`

	// READ-ONLY; The location of the original vault.
	Location *string `json:"location,omitempty" azure:"ro"`

	// READ-ONLY; The scheduled purged date.
	ScheduledPurgeDate *time.Time `json:"scheduledPurgeDate,omitempty" azure:"ro"`

	// READ-ONLY; Tags of the original vault.
	Tags *map[string]string `json:"tags,omitempty" azure:"ro"`

	// READ-ONLY; The resource id of the original vault.
	VaultID *string `json:"vaultId,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type DeletedVaultProperties.
func (d DeletedVaultProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if d.DeletionDate != nil {
		objectMap["deletionDate"] = (*timeRFC3339)(d.DeletionDate)
	}
	if d.Location != nil {
		objectMap["location"] = d.Location
	}
	if d.ScheduledPurgeDate != nil {
		objectMap["scheduledPurgeDate"] = (*timeRFC3339)(d.ScheduledPurgeDate)
	}
	if d.Tags != nil {
		objectMap["tags"] = d.Tags
	}
	if d.VaultID != nil {
		objectMap["vaultId"] = d.VaultID
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DeletedVaultProperties.
func (d *DeletedVaultProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "deletionDate":
			if val != nil {
				var aux timeRFC3339
				err = json.Unmarshal(*val, &aux)
				d.DeletionDate = (*time.Time)(&aux)
			}
			delete(rawMsg, key)
		case "location":
			if val != nil {
				err = json.Unmarshal(*val, &d.Location)
			}
			delete(rawMsg, key)
		case "scheduledPurgeDate":
			if val != nil {
				var aux timeRFC3339
				err = json.Unmarshal(*val, &aux)
				d.ScheduledPurgeDate = (*time.Time)(&aux)
			}
			delete(rawMsg, key)
		case "tags":
			if val != nil {
				err = json.Unmarshal(*val, &d.Tags)
			}
			delete(rawMsg, key)
		case "vaultId":
			if val != nil {
				err = json.Unmarshal(*val, &d.VaultID)
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// DeletedVaultResponse is the response envelope for operations that return a DeletedVault type.
type DeletedVaultResponse struct {
	// Deleted vault information with extended details.
	DeletedVault *DeletedVault

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HTTPPollerResponse contains the asynchronous HTTP response from the call to the service endpoint.
type HTTPPollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received
	PollUntilDone func(ctx context.Context, frequency time.Duration) (*http.Response, error)

	// Poller contains an initialized poller.
	Poller HTTPPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// A rule governing the accessibility of a vault from a specific ip address or ip range.
type IPRule struct {
	// An IPv4 address range in CIDR notation, such as '124.56.78.91' (simple IP address) or '124.56.78.0/24' (all addresses that start with 124.56.78).
	Value *string `json:"value,omitempty"`
}

// The key resource.
type Key struct {
	Resource
	// The properties of the key.
	Properties *KeyProperties `json:"properties,omitempty"`
}

// The attributes of the key.
type KeyAttributes struct {
	Attributes
}

// The parameters used to create a key.
type KeyCreateParameters struct {
	// The properties of the key to be created.
	Properties *KeyProperties `json:"properties,omitempty"`

	// The tags that will be assigned to the key.
	Tags *map[string]string `json:"tags,omitempty"`
}

// The page of keys.
type KeyListResult struct {
	// The URL to get the next page of keys.
	NextLink *string `json:"nextLink,omitempty"`

	// The key resources.
	Value *[]Key `json:"value,omitempty"`
}

// KeyListResultResponse is the response envelope for operations that return a KeyListResult type.
type KeyListResultResponse struct {
	// The page of keys.
	KeyListResult *KeyListResult

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// The properties of the key.
type KeyProperties struct {
	// The attributes of the key.
	Attributes *Attributes `json:"attributes,omitempty"`

	// The elliptic curve name. For valid values, see JsonWebKeyCurveName.
	CurveName *JSONWebKeyCurveName   `json:"curveName,omitempty"`
	KeyOps    *[]JSONWebKeyOperation `json:"keyOps,omitempty"`

	// The key size in bits. For example: 2048, 3072, or 4096 for RSA.
	KeySize *int32 `json:"keySize,omitempty"`

	// READ-ONLY; The URI to retrieve the current version of the key.
	KeyURI *string `json:"keyUri,omitempty" azure:"ro"`

	// READ-ONLY; The URI to retrieve the specific version of the key.
	KeyURIWithVersion *string `json:"keyUriWithVersion,omitempty" azure:"ro"`

	// The type of the key. For valid values, see JsonWebKeyType.
	Kty *JSONWebKeyType `json:"kty,omitempty"`
}

// KeyResponse is the response envelope for operations that return a Key type.
type KeyResponse struct {
	// The key resource.
	Key *Key

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// KeysCreateIfNotExistOptions contains the optional parameters for the Keys.CreateIfNotExist method.
type KeysCreateIfNotExistOptions struct {
	// placeholder for future optional parameters
}

// KeysGetOptions contains the optional parameters for the Keys.Get method.
type KeysGetOptions struct {
	// placeholder for future optional parameters
}

// KeysGetVersionOptions contains the optional parameters for the Keys.GetVersion method.
type KeysGetVersionOptions struct {
	// placeholder for future optional parameters
}

// KeysListOptions contains the optional parameters for the Keys.List method.
type KeysListOptions struct {
	// placeholder for future optional parameters
}

// KeysListVersionsOptions contains the optional parameters for the Keys.ListVersions method.
type KeysListVersionsOptions struct {
	// placeholder for future optional parameters
}

// Log specification of operation.
type LogSpecification struct {
	// Blob duration of specification.
	BlobDuration *string `json:"blobDuration,omitempty"`

	// Display name of log specification.
	DisplayName *string `json:"displayName,omitempty"`

	// Name of log specification.
	Name *string `json:"name,omitempty"`
}

// A set of rules governing the network accessibility of a vault.
type NetworkRuleSet struct {
	// Tells what traffic can bypass network rules. This can be 'AzureServices' or 'None'. If not specified the default is 'AzureServices'.
	Bypass *NetworkRuleBypassOptions `json:"bypass,omitempty"`

	// The default action when no rule from ipRules and from virtualNetworkRules match. This is only used after the bypass property has been evaluated.
	DefaultAction *NetworkRuleAction `json:"defaultAction,omitempty"`

	// The list of IP address rules.
	IPRules *[]IPRule `json:"ipRules,omitempty"`

	// The list of virtual network rules.
	VirtualNetworkRules *[]VirtualNetworkRule `json:"virtualNetworkRules,omitempty"`
}

// Key Vault REST API operation definition.
type Operation struct {
	// Display metadata associated with the operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`

	// Properties of operation, include metric specifications.
	OperationProperties *OperationProperties `json:"properties,omitempty"`

	// The origin of operations.
	Origin *string `json:"origin,omitempty"`
}

// Display metadata associated with the operation.
type OperationDisplay struct {
	// Description of operation.
	Description *string `json:"description,omitempty"`

	// Type of operation: get, read, delete, etc.
	Operation *string `json:"operation,omitempty"`

	// Service provider: Microsoft Key Vault.
	Provider *string `json:"provider,omitempty"`

	// Resource on which the operation is performed etc.
	Resource *string `json:"resource,omitempty"`
}

// Result of the request to list Storage operations. It contains a list of operations and a URL link to get the next set of results.
type OperationListResult struct {
	// The URL to get the next set of operations.
	NextLink *string `json:"nextLink,omitempty"`

	// List of Storage operations supported by the Storage resource provider.
	Value *[]Operation `json:"value,omitempty"`
}

// OperationListResultResponse is the response envelope for operations that return a OperationListResult type.
type OperationListResultResponse struct {
	// Result of the request to list Storage operations. It contains a list of operations and a URL link to get the next set of results.
	OperationListResult *OperationListResult

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// Properties of operation, include metric specifications.
type OperationProperties struct {
	// One property of operation, include metric specifications.
	ServiceSpecification *ServiceSpecification `json:"serviceSpecification,omitempty"`
}

// OperationsListOptions contains the optional parameters for the Operations.List method.
type OperationsListOptions struct {
	// placeholder for future optional parameters
}

// Permissions the identity has for keys, secrets, certificates and storage.
type Permissions struct {
	// Permissions to certificates
	Certificates *[]CertificatePermissions `json:"certificates,omitempty"`

	// Permissions to keys
	Keys *[]KeyPermissions `json:"keys,omitempty"`

	// Permissions to secrets
	Secrets *[]SecretPermissions `json:"secrets,omitempty"`

	// Permissions to storage accounts
	Storage *[]StoragePermissions `json:"storage,omitempty"`
}

// Private endpoint object properties.
type PrivateEndpoint struct {
	// READ-ONLY; Full identifier of the private endpoint resource.
	ID *string `json:"id,omitempty" azure:"ro"`
}

// Private endpoint connection resource.
type PrivateEndpointConnection struct {
	Resource
	// Resource properties.
	Properties *PrivateEndpointConnectionProperties `json:"properties,omitempty"`
}

// Private endpoint connection item.
type PrivateEndpointConnectionItem struct {
	// Private endpoint connection properties.
	Properties *PrivateEndpointConnectionProperties `json:"properties,omitempty"`
}

// PrivateEndpointConnectionPollerResponse is the response envelope for operations that asynchronously return a PrivateEndpointConnection type.
type PrivateEndpointConnectionPollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received
	PollUntilDone func(ctx context.Context, frequency time.Duration) (PrivateEndpointConnectionResponse, error)

	// Poller contains an initialized poller.
	Poller PrivateEndpointConnectionPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// Properties of the private endpoint connection resource.
type PrivateEndpointConnectionProperties struct {
	// Properties of the private endpoint object.
	PrivateEndpoint *PrivateEndpoint `json:"privateEndpoint,omitempty"`

	// Approval state of the private link connection.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState `json:"privateLinkServiceConnectionState,omitempty"`

	// READ-ONLY; Provisioning state of the private endpoint connection.
	ProvisioningState *PrivateEndpointConnectionProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// PrivateEndpointConnectionResponse is the response envelope for operations that return a PrivateEndpointConnection type.
type PrivateEndpointConnectionResponse struct {
	// AzureAsyncOperation contains the information returned from the Azure-AsyncOperation header response.
	AzureAsyncOperation *string

	// Private endpoint connection resource.
	PrivateEndpointConnection *PrivateEndpointConnection

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// RetryAfter contains the information returned from the Retry-After header response.
	RetryAfter *int32
}

// PrivateEndpointConnectionsBeginDeleteOptions contains the optional parameters for the PrivateEndpointConnections.BeginDelete method.
type PrivateEndpointConnectionsBeginDeleteOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsGetOptions contains the optional parameters for the PrivateEndpointConnections.Get method.
type PrivateEndpointConnectionsGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsPutOptions contains the optional parameters for the PrivateEndpointConnections.Put method.
type PrivateEndpointConnectionsPutOptions struct {
	// placeholder for future optional parameters
}

// A private link resource
type PrivateLinkResource struct {
	Resource
	// Resource properties.
	Properties *PrivateLinkResourceProperties `json:"properties,omitempty"`
}

// A list of private link resources
type PrivateLinkResourceListResult struct {
	// Array of private link resources
	Value *[]PrivateLinkResource `json:"value,omitempty"`
}

// PrivateLinkResourceListResultResponse is the response envelope for operations that return a PrivateLinkResourceListResult type.
type PrivateLinkResourceListResultResponse struct {
	// A list of private link resources
	PrivateLinkResourceListResult *PrivateLinkResourceListResult

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// Properties of a private link resource.
type PrivateLinkResourceProperties struct {
	// READ-ONLY; Group identifier of private link resource.
	GroupID *string `json:"groupId,omitempty" azure:"ro"`

	// READ-ONLY; Required member names of private link resource.
	RequiredMembers *[]string `json:"requiredMembers,omitempty" azure:"ro"`

	// Required DNS zone names of the the private link resource.
	RequiredZoneNames *[]string `json:"requiredZoneNames,omitempty"`
}

// PrivateLinkResourcesListByVaultOptions contains the optional parameters for the PrivateLinkResources.ListByVault method.
type PrivateLinkResourcesListByVaultOptions struct {
	// placeholder for future optional parameters
}

// An object that represents the approval state of the private link connection.
type PrivateLinkServiceConnectionState struct {
	// A message indicating if changes on the service provider require any updates on the consumer.
	ActionRequired *string `json:"actionRequired,omitempty"`

	// The reason for approval or rejection.
	Description *string `json:"description,omitempty"`

	// Indicates whether the connection has been approved, rejected or removed by the key vault owner.
	Status *PrivateEndpointServiceConnectionStatus `json:"status,omitempty"`
}

// Key Vault resource
type Resource struct {
	// READ-ONLY; Fully qualified identifier of the key vault resource.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Azure location of the key vault resource.
	Location *string `json:"location,omitempty" azure:"ro"`

	// READ-ONLY; Name of the key vault resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Tags assigned to the key vault resource.
	Tags *map[string]string `json:"tags,omitempty" azure:"ro"`

	// READ-ONLY; Resource type of the key vault resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// List of vault resources.
type ResourceListResult struct {
	// The URL to get the next set of vault resources.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of vault resources.
	Value *[]Resource `json:"value,omitempty"`
}

// ResourceListResultResponse is the response envelope for operations that return a ResourceListResult type.
type ResourceListResultResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// List of vault resources.
	ResourceListResult *ResourceListResult
}

// SKU details
type SKU struct {
	// SKU family name
	Family *SKUFamily `json:"family,omitempty"`

	// SKU name to specify whether the key vault is a standard vault or a premium vault.
	Name *SKUName `json:"name,omitempty"`
}

// One property of operation, include log specifications.
type ServiceSpecification struct {
	// Log specifications of operation.
	LogSpecifications *[]LogSpecification `json:"logSpecifications,omitempty"`
}

// Resource information with extended details.
type Vault struct {
	// READ-ONLY; Fully qualified identifier of the key vault resource.
	ID *string `json:"id,omitempty" azure:"ro"`

	// Azure location of the key vault resource.
	Location *string `json:"location,omitempty"`

	// READ-ONLY; Name of the key vault resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// Properties of the vault
	Properties *VaultProperties `json:"properties,omitempty"`

	// Tags assigned to the key vault resource.
	Tags *map[string]string `json:"tags,omitempty"`

	// READ-ONLY; Resource type of the key vault resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// Parameters for updating the access policy in a vault
type VaultAccessPolicyParameters struct {
	// READ-ONLY; The resource id of the access policy.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The resource type of the access policy.
	Location *string `json:"location,omitempty" azure:"ro"`

	// READ-ONLY; The resource name of the access policy.
	Name *string `json:"name,omitempty" azure:"ro"`

	// Properties of the access policy
	Properties *VaultAccessPolicyProperties `json:"properties,omitempty"`

	// READ-ONLY; The resource name of the access policy.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// VaultAccessPolicyParametersResponse is the response envelope for operations that return a VaultAccessPolicyParameters type.
type VaultAccessPolicyParametersResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Parameters for updating the access policy in a vault
	VaultAccessPolicyParameters *VaultAccessPolicyParameters
}

// Properties of the vault access policy
type VaultAccessPolicyProperties struct {
	// An array of 0 to 16 identities that have access to the key vault. All identities in the array must use the same tenant ID as the key vault's tenant ID.
	AccessPolicies *[]AccessPolicyEntry `json:"accessPolicies,omitempty"`
}

// The parameters used to check the availability of the vault name.
type VaultCheckNameAvailabilityParameters struct {
	// The vault name.
	Name *string `json:"name,omitempty"`

	// The type of resource, Microsoft.KeyVault/vaults
	Type *string `json:"type,omitempty"`
}

// Parameters for creating or updating a vault
type VaultCreateOrUpdateParameters struct {
	// The supported Azure location where the key vault should be created.
	Location *string `json:"location,omitempty"`

	// Properties of the vault
	Properties *VaultProperties `json:"properties,omitempty"`

	// The tags that will be assigned to the key vault.
	Tags *map[string]string `json:"tags,omitempty"`
}

// List of vaults
type VaultListResult struct {
	// The URL to get the next set of vaults.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of vaults.
	Value *[]Vault `json:"value,omitempty"`
}

// VaultListResultResponse is the response envelope for operations that return a VaultListResult type.
type VaultListResultResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// List of vaults
	VaultListResult *VaultListResult
}

// Parameters for creating or updating a vault
type VaultPatchParameters struct {
	// Properties of the vault
	Properties *VaultPatchProperties `json:"properties,omitempty"`

	// The tags that will be assigned to the key vault.
	Tags *map[string]string `json:"tags,omitempty"`
}

// Properties of the vault
type VaultPatchProperties struct {
	// An array of 0 to 16 identities that have access to the key vault. All identities in the array must use the same tenant ID as the key vault's tenant ID.
	AccessPolicies *[]AccessPolicyEntry `json:"accessPolicies,omitempty"`

	// The vault's create mode to indicate whether the vault need to be recovered or not.
	CreateMode *CreateMode `json:"createMode,omitempty"`

	// Property specifying whether protection against purge is enabled for this vault. Setting this property to true activates protection against purge for
	// this vault and its content - only the Key Vault
	// service may initiate a hard, irrecoverable deletion. The setting is effective only if soft delete is also enabled. Enabling this functionality is irreversible
	// - that is, the property does not accept
	// false as its value.
	EnablePurgeProtection *bool `json:"enablePurgeProtection,omitempty"`

	// Property that controls how data actions are authorized. When true, the key vault will use Role Based Access Control (RBAC) for authorization of data
	// actions, and the access policies specified in vault
	// properties will be ignored (warning: this is a preview feature). When false, the key vault will use the access policies specified in vault properties,
	// and any policy stored on Azure Resource Manager
	// will be ignored. If null or not specified, the value of this property will not change.
	EnableRbacAuthorization *bool `json:"enableRbacAuthorization,omitempty"`

	// Property to specify whether the 'soft delete' functionality is enabled for this key vault. Once set to true, it cannot be reverted to false.
	EnableSoftDelete *bool `json:"enableSoftDelete,omitempty"`

	// Property to specify whether Azure Virtual Machines are permitted to retrieve certificates stored as secrets from the key vault.
	EnabledForDeployment *bool `json:"enabledForDeployment,omitempty"`

	// Property to specify whether Azure Disk Encryption is permitted to retrieve secrets from the vault and unwrap keys.
	EnabledForDiskEncryption *bool `json:"enabledForDiskEncryption,omitempty"`

	// Property to specify whether Azure Resource Manager is permitted to retrieve secrets from the key vault.
	EnabledForTemplateDeployment *bool `json:"enabledForTemplateDeployment,omitempty"`

	// A collection of rules governing the accessibility of the vault from specific network locations.
	NetworkACLs *NetworkRuleSet `json:"networkAcls,omitempty"`

	// SKU details
	SKU *SKU `json:"sku,omitempty"`

	// softDelete data retention days. It accepts >=7 and <=90.
	SoftDeleteRetentionInDays *int32 `json:"softDeleteRetentionInDays,omitempty"`

	// The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
	TenantID *string `json:"tenantId,omitempty"`
}

// VaultPollerResponse is the response envelope for operations that asynchronously return a Vault type.
type VaultPollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received
	PollUntilDone func(ctx context.Context, frequency time.Duration) (VaultResponse, error)

	// Poller contains an initialized poller.
	Poller VaultPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// Properties of the vault
type VaultProperties struct {
	// An array of 0 to 1024 identities that have access to the key vault. All identities in the array must use the same tenant ID as the key vault's tenant
	// ID. When createMode is set to recover, access
	// policies are not required. Otherwise, access policies are required.
	AccessPolicies *[]AccessPolicyEntry `json:"accessPolicies,omitempty"`

	// The vault's create mode to indicate whether the vault need to be recovered or not.
	CreateMode *CreateMode `json:"createMode,omitempty"`

	// Property specifying whether protection against purge is enabled for this vault. Setting this property to true activates protection against purge for
	// this vault and its content - only the Key Vault
	// service may initiate a hard, irrecoverable deletion. The setting is effective only if soft delete is also enabled. Enabling this functionality is irreversible
	// - that is, the property does not accept
	// false as its value.
	EnablePurgeProtection *bool `json:"enablePurgeProtection,omitempty"`

	// Property that controls how data actions are authorized. When true, the key vault will use Role Based Access Control (RBAC) for authorization of data
	// actions, and the access policies specified in vault
	// properties will be ignored (warning: this is a preview feature). When false, the key vault will use the access policies specified in vault properties,
	// and any policy stored on Azure Resource Manager
	// will be ignored. If null or not specified, the vault is created with the default value of false. Note that management actions are always authorized with
	// RBAC.
	EnableRbacAuthorization *bool `json:"enableRbacAuthorization,omitempty"`

	// Property to specify whether the 'soft delete' functionality is enabled for this key vault. If it's not set to any value(true or false) when creating
	// new key vault, it will be set to true by default.
	// Once set to true, it cannot be reverted to false.
	EnableSoftDelete *bool `json:"enableSoftDelete,omitempty"`

	// Property to specify whether Azure Virtual Machines are permitted to retrieve certificates stored as secrets from the key vault.
	EnabledForDeployment *bool `json:"enabledForDeployment,omitempty"`

	// Property to specify whether Azure Disk Encryption is permitted to retrieve secrets from the vault and unwrap keys.
	EnabledForDiskEncryption *bool `json:"enabledForDiskEncryption,omitempty"`

	// Property to specify whether Azure Resource Manager is permitted to retrieve secrets from the key vault.
	EnabledForTemplateDeployment *bool `json:"enabledForTemplateDeployment,omitempty"`

	// Rules governing the accessibility of the key vault from specific network locations.
	NetworkACLs *NetworkRuleSet `json:"networkAcls,omitempty"`

	// READ-ONLY; List of private endpoint connections associated with the key vault.
	PrivateEndpointConnections *[]PrivateEndpointConnectionItem `json:"privateEndpointConnections,omitempty" azure:"ro"`

	// SKU details
	SKU *SKU `json:"sku,omitempty"`

	// softDelete data retention days. It accepts >=7 and <=90.
	SoftDeleteRetentionInDays *int32 `json:"softDeleteRetentionInDays,omitempty"`

	// The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
	TenantID *string `json:"tenantId,omitempty"`

	// The URI of the vault for performing operations on keys and secrets.
	VaultURI *string `json:"vaultUri,omitempty"`
}

// VaultResponse is the response envelope for operations that return a Vault type.
type VaultResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Resource information with extended details.
	Vault *Vault
}

// VaultsBeginCreateOrUpdateOptions contains the optional parameters for the Vaults.BeginCreateOrUpdate method.
type VaultsBeginCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// VaultsBeginPurgeDeletedOptions contains the optional parameters for the Vaults.BeginPurgeDeleted method.
type VaultsBeginPurgeDeletedOptions struct {
	// placeholder for future optional parameters
}

// VaultsCheckNameAvailabilityOptions contains the optional parameters for the Vaults.CheckNameAvailability method.
type VaultsCheckNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// VaultsDeleteOptions contains the optional parameters for the Vaults.Delete method.
type VaultsDeleteOptions struct {
	// placeholder for future optional parameters
}

// VaultsGetDeletedOptions contains the optional parameters for the Vaults.GetDeleted method.
type VaultsGetDeletedOptions struct {
	// placeholder for future optional parameters
}

// VaultsGetOptions contains the optional parameters for the Vaults.Get method.
type VaultsGetOptions struct {
	// placeholder for future optional parameters
}

// VaultsListByResourceGroupOptions contains the optional parameters for the Vaults.ListByResourceGroup method.
type VaultsListByResourceGroupOptions struct {
	// Maximum number of results to return.
	Top *int32
}

// VaultsListBySubscriptionOptions contains the optional parameters for the Vaults.ListBySubscription method.
type VaultsListBySubscriptionOptions struct {
	// Maximum number of results to return.
	Top *int32
}

// VaultsListDeletedOptions contains the optional parameters for the Vaults.ListDeleted method.
type VaultsListDeletedOptions struct {
	// placeholder for future optional parameters
}

// VaultsListOptions contains the optional parameters for the Vaults.List method.
type VaultsListOptions struct {
	// Maximum number of results to return.
	Top *int32
}

// VaultsUpdateAccessPolicyOptions contains the optional parameters for the Vaults.UpdateAccessPolicy method.
type VaultsUpdateAccessPolicyOptions struct {
	// placeholder for future optional parameters
}

// VaultsUpdateOptions contains the optional parameters for the Vaults.Update method.
type VaultsUpdateOptions struct {
	// placeholder for future optional parameters
}

// A rule governing the accessibility of a vault from a specific virtual network.
type VirtualNetworkRule struct {
	// Full resource id of a vnet subnet, such as '/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/subnet1'.
	ID *string `json:"id,omitempty"`
}
