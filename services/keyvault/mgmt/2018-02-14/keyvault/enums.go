package keyvault

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AccessPolicyUpdateKind enumerates the values for access policy update kind.
type AccessPolicyUpdateKind string

const (
	// Add ...
	Add AccessPolicyUpdateKind = "add"
	// Remove ...
	Remove AccessPolicyUpdateKind = "remove"
	// Replace ...
	Replace AccessPolicyUpdateKind = "replace"
)

// PossibleAccessPolicyUpdateKindValues returns an array of possible values for the AccessPolicyUpdateKind const type.
func PossibleAccessPolicyUpdateKindValues() []AccessPolicyUpdateKind {
	return []AccessPolicyUpdateKind{Add, Remove, Replace}
}

// CertificatePermissions enumerates the values for certificate permissions.
type CertificatePermissions string

const (
	// Backup ...
	Backup CertificatePermissions = "backup"
	// Create ...
	Create CertificatePermissions = "create"
	// Delete ...
	Delete CertificatePermissions = "delete"
	// Deleteissuers ...
	Deleteissuers CertificatePermissions = "deleteissuers"
	// Get ...
	Get CertificatePermissions = "get"
	// Getissuers ...
	Getissuers CertificatePermissions = "getissuers"
	// Import ...
	Import CertificatePermissions = "import"
	// List ...
	List CertificatePermissions = "list"
	// Listissuers ...
	Listissuers CertificatePermissions = "listissuers"
	// Managecontacts ...
	Managecontacts CertificatePermissions = "managecontacts"
	// Manageissuers ...
	Manageissuers CertificatePermissions = "manageissuers"
	// Purge ...
	Purge CertificatePermissions = "purge"
	// Recover ...
	Recover CertificatePermissions = "recover"
	// Restore ...
	Restore CertificatePermissions = "restore"
	// Setissuers ...
	Setissuers CertificatePermissions = "setissuers"
	// Update ...
	Update CertificatePermissions = "update"
)

// PossibleCertificatePermissionsValues returns an array of possible values for the CertificatePermissions const type.
func PossibleCertificatePermissionsValues() []CertificatePermissions {
	return []CertificatePermissions{Backup, Create, Delete, Deleteissuers, Get, Getissuers, Import, List, Listissuers, Managecontacts, Manageissuers, Purge, Recover, Restore, Setissuers, Update}
}

// CreateMode enumerates the values for create mode.
type CreateMode string

const (
	// CreateModeDefault ...
	CreateModeDefault CreateMode = "default"
	// CreateModeRecover ...
	CreateModeRecover CreateMode = "recover"
)

// PossibleCreateModeValues returns an array of possible values for the CreateMode const type.
func PossibleCreateModeValues() []CreateMode {
	return []CreateMode{CreateModeDefault, CreateModeRecover}
}

// KeyPermissions enumerates the values for key permissions.
type KeyPermissions string

const (
	// KeyPermissionsBackup ...
	KeyPermissionsBackup KeyPermissions = "backup"
	// KeyPermissionsCreate ...
	KeyPermissionsCreate KeyPermissions = "create"
	// KeyPermissionsDecrypt ...
	KeyPermissionsDecrypt KeyPermissions = "decrypt"
	// KeyPermissionsDelete ...
	KeyPermissionsDelete KeyPermissions = "delete"
	// KeyPermissionsEncrypt ...
	KeyPermissionsEncrypt KeyPermissions = "encrypt"
	// KeyPermissionsGet ...
	KeyPermissionsGet KeyPermissions = "get"
	// KeyPermissionsImport ...
	KeyPermissionsImport KeyPermissions = "import"
	// KeyPermissionsList ...
	KeyPermissionsList KeyPermissions = "list"
	// KeyPermissionsPurge ...
	KeyPermissionsPurge KeyPermissions = "purge"
	// KeyPermissionsRecover ...
	KeyPermissionsRecover KeyPermissions = "recover"
	// KeyPermissionsRestore ...
	KeyPermissionsRestore KeyPermissions = "restore"
	// KeyPermissionsSign ...
	KeyPermissionsSign KeyPermissions = "sign"
	// KeyPermissionsUnwrapKey ...
	KeyPermissionsUnwrapKey KeyPermissions = "unwrapKey"
	// KeyPermissionsUpdate ...
	KeyPermissionsUpdate KeyPermissions = "update"
	// KeyPermissionsVerify ...
	KeyPermissionsVerify KeyPermissions = "verify"
	// KeyPermissionsWrapKey ...
	KeyPermissionsWrapKey KeyPermissions = "wrapKey"
)

// PossibleKeyPermissionsValues returns an array of possible values for the KeyPermissions const type.
func PossibleKeyPermissionsValues() []KeyPermissions {
	return []KeyPermissions{KeyPermissionsBackup, KeyPermissionsCreate, KeyPermissionsDecrypt, KeyPermissionsDelete, KeyPermissionsEncrypt, KeyPermissionsGet, KeyPermissionsImport, KeyPermissionsList, KeyPermissionsPurge, KeyPermissionsRecover, KeyPermissionsRestore, KeyPermissionsSign, KeyPermissionsUnwrapKey, KeyPermissionsUpdate, KeyPermissionsVerify, KeyPermissionsWrapKey}
}

// NetworkRuleAction enumerates the values for network rule action.
type NetworkRuleAction string

const (
	// Allow ...
	Allow NetworkRuleAction = "Allow"
	// Deny ...
	Deny NetworkRuleAction = "Deny"
)

// PossibleNetworkRuleActionValues returns an array of possible values for the NetworkRuleAction const type.
func PossibleNetworkRuleActionValues() []NetworkRuleAction {
	return []NetworkRuleAction{Allow, Deny}
}

// NetworkRuleBypassOptions enumerates the values for network rule bypass options.
type NetworkRuleBypassOptions string

const (
	// AzureServices ...
	AzureServices NetworkRuleBypassOptions = "AzureServices"
	// None ...
	None NetworkRuleBypassOptions = "None"
)

// PossibleNetworkRuleBypassOptionsValues returns an array of possible values for the NetworkRuleBypassOptions const type.
func PossibleNetworkRuleBypassOptionsValues() []NetworkRuleBypassOptions {
	return []NetworkRuleBypassOptions{AzureServices, None}
}

// PrivateEndpointConnectionProvisioningState enumerates the values for private endpoint connection
// provisioning state.
type PrivateEndpointConnectionProvisioningState string

const (
	// Creating ...
	Creating PrivateEndpointConnectionProvisioningState = "Creating"
	// Deleting ...
	Deleting PrivateEndpointConnectionProvisioningState = "Deleting"
	// Disconnected ...
	Disconnected PrivateEndpointConnectionProvisioningState = "Disconnected"
	// Failed ...
	Failed PrivateEndpointConnectionProvisioningState = "Failed"
	// Succeeded ...
	Succeeded PrivateEndpointConnectionProvisioningState = "Succeeded"
	// Updating ...
	Updating PrivateEndpointConnectionProvisioningState = "Updating"
)

// PossiblePrivateEndpointConnectionProvisioningStateValues returns an array of possible values for the PrivateEndpointConnectionProvisioningState const type.
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return []PrivateEndpointConnectionProvisioningState{Creating, Deleting, Disconnected, Failed, Succeeded, Updating}
}

// PrivateEndpointServiceConnectionStatus enumerates the values for private endpoint service connection status.
type PrivateEndpointServiceConnectionStatus string

const (
	// PrivateEndpointServiceConnectionStatusApproved ...
	PrivateEndpointServiceConnectionStatusApproved PrivateEndpointServiceConnectionStatus = "Approved"
	// PrivateEndpointServiceConnectionStatusDisconnected ...
	PrivateEndpointServiceConnectionStatusDisconnected PrivateEndpointServiceConnectionStatus = "Disconnected"
	// PrivateEndpointServiceConnectionStatusPending ...
	PrivateEndpointServiceConnectionStatusPending PrivateEndpointServiceConnectionStatus = "Pending"
	// PrivateEndpointServiceConnectionStatusRejected ...
	PrivateEndpointServiceConnectionStatusRejected PrivateEndpointServiceConnectionStatus = "Rejected"
)

// PossiblePrivateEndpointServiceConnectionStatusValues returns an array of possible values for the PrivateEndpointServiceConnectionStatus const type.
func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return []PrivateEndpointServiceConnectionStatus{PrivateEndpointServiceConnectionStatusApproved, PrivateEndpointServiceConnectionStatusDisconnected, PrivateEndpointServiceConnectionStatusPending, PrivateEndpointServiceConnectionStatusRejected}
}

// Reason enumerates the values for reason.
type Reason string

const (
	// AccountNameInvalid ...
	AccountNameInvalid Reason = "AccountNameInvalid"
	// AlreadyExists ...
	AlreadyExists Reason = "AlreadyExists"
)

// PossibleReasonValues returns an array of possible values for the Reason const type.
func PossibleReasonValues() []Reason {
	return []Reason{AccountNameInvalid, AlreadyExists}
}

// SecretPermissions enumerates the values for secret permissions.
type SecretPermissions string

const (
	// SecretPermissionsBackup ...
	SecretPermissionsBackup SecretPermissions = "backup"
	// SecretPermissionsDelete ...
	SecretPermissionsDelete SecretPermissions = "delete"
	// SecretPermissionsGet ...
	SecretPermissionsGet SecretPermissions = "get"
	// SecretPermissionsList ...
	SecretPermissionsList SecretPermissions = "list"
	// SecretPermissionsPurge ...
	SecretPermissionsPurge SecretPermissions = "purge"
	// SecretPermissionsRecover ...
	SecretPermissionsRecover SecretPermissions = "recover"
	// SecretPermissionsRestore ...
	SecretPermissionsRestore SecretPermissions = "restore"
	// SecretPermissionsSet ...
	SecretPermissionsSet SecretPermissions = "set"
)

// PossibleSecretPermissionsValues returns an array of possible values for the SecretPermissions const type.
func PossibleSecretPermissionsValues() []SecretPermissions {
	return []SecretPermissions{SecretPermissionsBackup, SecretPermissionsDelete, SecretPermissionsGet, SecretPermissionsList, SecretPermissionsPurge, SecretPermissionsRecover, SecretPermissionsRestore, SecretPermissionsSet}
}

// SkuName enumerates the values for sku name.
type SkuName string

const (
	// Premium ...
	Premium SkuName = "premium"
	// Standard ...
	Standard SkuName = "standard"
)

// PossibleSkuNameValues returns an array of possible values for the SkuName const type.
func PossibleSkuNameValues() []SkuName {
	return []SkuName{Premium, Standard}
}

// StoragePermissions enumerates the values for storage permissions.
type StoragePermissions string

const (
	// StoragePermissionsBackup ...
	StoragePermissionsBackup StoragePermissions = "backup"
	// StoragePermissionsDelete ...
	StoragePermissionsDelete StoragePermissions = "delete"
	// StoragePermissionsDeletesas ...
	StoragePermissionsDeletesas StoragePermissions = "deletesas"
	// StoragePermissionsGet ...
	StoragePermissionsGet StoragePermissions = "get"
	// StoragePermissionsGetsas ...
	StoragePermissionsGetsas StoragePermissions = "getsas"
	// StoragePermissionsList ...
	StoragePermissionsList StoragePermissions = "list"
	// StoragePermissionsListsas ...
	StoragePermissionsListsas StoragePermissions = "listsas"
	// StoragePermissionsPurge ...
	StoragePermissionsPurge StoragePermissions = "purge"
	// StoragePermissionsRecover ...
	StoragePermissionsRecover StoragePermissions = "recover"
	// StoragePermissionsRegeneratekey ...
	StoragePermissionsRegeneratekey StoragePermissions = "regeneratekey"
	// StoragePermissionsRestore ...
	StoragePermissionsRestore StoragePermissions = "restore"
	// StoragePermissionsSet ...
	StoragePermissionsSet StoragePermissions = "set"
	// StoragePermissionsSetsas ...
	StoragePermissionsSetsas StoragePermissions = "setsas"
	// StoragePermissionsUpdate ...
	StoragePermissionsUpdate StoragePermissions = "update"
)

// PossibleStoragePermissionsValues returns an array of possible values for the StoragePermissions const type.
func PossibleStoragePermissionsValues() []StoragePermissions {
	return []StoragePermissions{StoragePermissionsBackup, StoragePermissionsDelete, StoragePermissionsDeletesas, StoragePermissionsGet, StoragePermissionsGetsas, StoragePermissionsList, StoragePermissionsListsas, StoragePermissionsPurge, StoragePermissionsRecover, StoragePermissionsRegeneratekey, StoragePermissionsRestore, StoragePermissionsSet, StoragePermissionsSetsas, StoragePermissionsUpdate}
}
