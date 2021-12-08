//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicessiterecovery

const (
	module  = "armrecoveryservicessiterecovery"
	version = "v0.1.0"
)

// AuthType - Specifies the authentication type.
type AuthType string

const (
	AuthTypeAAD                  AuthType = "AAD"
	AuthTypeACS                  AuthType = "ACS"
	AuthTypeAccessControlService AuthType = "AccessControlService"
	AuthTypeAzureActiveDirectory AuthType = "AzureActiveDirectory"
	AuthTypeInvalid              AuthType = "Invalid"
)

// PossibleAuthTypeValues returns the possible values for the AuthType const type.
func PossibleAuthTypeValues() []AuthType {
	return []AuthType{
		AuthTypeAAD,
		AuthTypeACS,
		AuthTypeAccessControlService,
		AuthTypeAzureActiveDirectory,
		AuthTypeInvalid,
	}
}

// ToPtr returns a *AuthType pointing to the current value.
func (c AuthType) ToPtr() *AuthType {
	return &c
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

// ToPtr returns a *CreatedByType pointing to the current value.
func (c CreatedByType) ToPtr() *CreatedByType {
	return &c
}

// InfrastructureEncryptionState - Enabling/Disabling the Double Encryption state
type InfrastructureEncryptionState string

const (
	InfrastructureEncryptionStateDisabled InfrastructureEncryptionState = "Disabled"
	InfrastructureEncryptionStateEnabled  InfrastructureEncryptionState = "Enabled"
)

// PossibleInfrastructureEncryptionStateValues returns the possible values for the InfrastructureEncryptionState const type.
func PossibleInfrastructureEncryptionStateValues() []InfrastructureEncryptionState {
	return []InfrastructureEncryptionState{
		InfrastructureEncryptionStateDisabled,
		InfrastructureEncryptionStateEnabled,
	}
}

// ToPtr returns a *InfrastructureEncryptionState pointing to the current value.
func (c InfrastructureEncryptionState) ToPtr() *InfrastructureEncryptionState {
	return &c
}

// PrivateEndpointConnectionStatus - Gets or sets the status.
type PrivateEndpointConnectionStatus string

const (
	PrivateEndpointConnectionStatusApproved     PrivateEndpointConnectionStatus = "Approved"
	PrivateEndpointConnectionStatusDisconnected PrivateEndpointConnectionStatus = "Disconnected"
	PrivateEndpointConnectionStatusPending      PrivateEndpointConnectionStatus = "Pending"
	PrivateEndpointConnectionStatusRejected     PrivateEndpointConnectionStatus = "Rejected"
)

// PossiblePrivateEndpointConnectionStatusValues returns the possible values for the PrivateEndpointConnectionStatus const type.
func PossiblePrivateEndpointConnectionStatusValues() []PrivateEndpointConnectionStatus {
	return []PrivateEndpointConnectionStatus{
		PrivateEndpointConnectionStatusApproved,
		PrivateEndpointConnectionStatusDisconnected,
		PrivateEndpointConnectionStatusPending,
		PrivateEndpointConnectionStatusRejected,
	}
}

// ToPtr returns a *PrivateEndpointConnectionStatus pointing to the current value.
func (c PrivateEndpointConnectionStatus) ToPtr() *PrivateEndpointConnectionStatus {
	return &c
}

// ProvisioningState - Gets or sets provisioning state of the private endpoint connection.
type ProvisioningState string

const (
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStatePending   ProvisioningState = "Pending"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStatePending,
		ProvisioningStateSucceeded,
	}
}

// ToPtr returns a *ProvisioningState pointing to the current value.
func (c ProvisioningState) ToPtr() *ProvisioningState {
	return &c
}

// ResourceIdentityType - The type of managed identity used. The type 'SystemAssigned, UserAssigned' includes both an implicitly created identity and a
// set of user-assigned identities. The type 'None' will remove any
// identities.
type ResourceIdentityType string

const (
	ResourceIdentityTypeNone                       ResourceIdentityType = "None"
	ResourceIdentityTypeSystemAssigned             ResourceIdentityType = "SystemAssigned"
	ResourceIdentityTypeSystemAssignedUserAssigned ResourceIdentityType = "SystemAssigned, UserAssigned"
	ResourceIdentityTypeUserAssigned               ResourceIdentityType = "UserAssigned"
)

// PossibleResourceIdentityTypeValues returns the possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{
		ResourceIdentityTypeNone,
		ResourceIdentityTypeSystemAssigned,
		ResourceIdentityTypeSystemAssignedUserAssigned,
		ResourceIdentityTypeUserAssigned,
	}
}

// ToPtr returns a *ResourceIdentityType pointing to the current value.
func (c ResourceIdentityType) ToPtr() *ResourceIdentityType {
	return &c
}

// ResourceMoveState - The State of the Resource after the move operation
type ResourceMoveState string

const (
	ResourceMoveStateCommitFailed    ResourceMoveState = "CommitFailed"
	ResourceMoveStateCommitTimedout  ResourceMoveState = "CommitTimedout"
	ResourceMoveStateCriticalFailure ResourceMoveState = "CriticalFailure"
	ResourceMoveStateFailure         ResourceMoveState = "Failure"
	ResourceMoveStateInProgress      ResourceMoveState = "InProgress"
	ResourceMoveStateMoveSucceeded   ResourceMoveState = "MoveSucceeded"
	ResourceMoveStatePartialSuccess  ResourceMoveState = "PartialSuccess"
	ResourceMoveStatePrepareFailed   ResourceMoveState = "PrepareFailed"
	ResourceMoveStatePrepareTimedout ResourceMoveState = "PrepareTimedout"
	ResourceMoveStateUnknown         ResourceMoveState = "Unknown"
)

// PossibleResourceMoveStateValues returns the possible values for the ResourceMoveState const type.
func PossibleResourceMoveStateValues() []ResourceMoveState {
	return []ResourceMoveState{
		ResourceMoveStateCommitFailed,
		ResourceMoveStateCommitTimedout,
		ResourceMoveStateCriticalFailure,
		ResourceMoveStateFailure,
		ResourceMoveStateInProgress,
		ResourceMoveStateMoveSucceeded,
		ResourceMoveStatePartialSuccess,
		ResourceMoveStatePrepareFailed,
		ResourceMoveStatePrepareTimedout,
		ResourceMoveStateUnknown,
	}
}

// ToPtr returns a *ResourceMoveState pointing to the current value.
func (c ResourceMoveState) ToPtr() *ResourceMoveState {
	return &c
}

// SKUName - The Sku name.
type SKUName string

const (
	SKUNameRS0      SKUName = "RS0"
	SKUNameStandard SKUName = "Standard"
)

// PossibleSKUNameValues returns the possible values for the SKUName const type.
func PossibleSKUNameValues() []SKUName {
	return []SKUName{
		SKUNameRS0,
		SKUNameStandard,
	}
}

// ToPtr returns a *SKUName pointing to the current value.
func (c SKUName) ToPtr() *SKUName {
	return &c
}

// TriggerType - The way the vault upgrade was triggered.
type TriggerType string

const (
	TriggerTypeForcedUpgrade TriggerType = "ForcedUpgrade"
	TriggerTypeUserTriggered TriggerType = "UserTriggered"
)

// PossibleTriggerTypeValues returns the possible values for the TriggerType const type.
func PossibleTriggerTypeValues() []TriggerType {
	return []TriggerType{
		TriggerTypeForcedUpgrade,
		TriggerTypeUserTriggered,
	}
}

// ToPtr returns a *TriggerType pointing to the current value.
func (c TriggerType) ToPtr() *TriggerType {
	return &c
}

// UsagesUnit - Unit of the usage.
type UsagesUnit string

const (
	UsagesUnitBytes          UsagesUnit = "Bytes"
	UsagesUnitBytesPerSecond UsagesUnit = "BytesPerSecond"
	UsagesUnitCount          UsagesUnit = "Count"
	UsagesUnitCountPerSecond UsagesUnit = "CountPerSecond"
	UsagesUnitPercent        UsagesUnit = "Percent"
	UsagesUnitSeconds        UsagesUnit = "Seconds"
)

// PossibleUsagesUnitValues returns the possible values for the UsagesUnit const type.
func PossibleUsagesUnitValues() []UsagesUnit {
	return []UsagesUnit{
		UsagesUnitBytes,
		UsagesUnitBytesPerSecond,
		UsagesUnitCount,
		UsagesUnitCountPerSecond,
		UsagesUnitPercent,
		UsagesUnitSeconds,
	}
}

// ToPtr returns a *UsagesUnit pointing to the current value.
func (c UsagesUnit) ToPtr() *UsagesUnit {
	return &c
}

// VaultPrivateEndpointState - Private endpoint state for backup.
type VaultPrivateEndpointState string

const (
	VaultPrivateEndpointStateEnabled VaultPrivateEndpointState = "Enabled"
	VaultPrivateEndpointStateNone    VaultPrivateEndpointState = "None"
)

// PossibleVaultPrivateEndpointStateValues returns the possible values for the VaultPrivateEndpointState const type.
func PossibleVaultPrivateEndpointStateValues() []VaultPrivateEndpointState {
	return []VaultPrivateEndpointState{
		VaultPrivateEndpointStateEnabled,
		VaultPrivateEndpointStateNone,
	}
}

// ToPtr returns a *VaultPrivateEndpointState pointing to the current value.
func (c VaultPrivateEndpointState) ToPtr() *VaultPrivateEndpointState {
	return &c
}

// VaultUpgradeState - Status of the vault upgrade operation.
type VaultUpgradeState string

const (
	VaultUpgradeStateFailed     VaultUpgradeState = "Failed"
	VaultUpgradeStateInProgress VaultUpgradeState = "InProgress"
	VaultUpgradeStateUnknown    VaultUpgradeState = "Unknown"
	VaultUpgradeStateUpgraded   VaultUpgradeState = "Upgraded"
)

// PossibleVaultUpgradeStateValues returns the possible values for the VaultUpgradeState const type.
func PossibleVaultUpgradeStateValues() []VaultUpgradeState {
	return []VaultUpgradeState{
		VaultUpgradeStateFailed,
		VaultUpgradeStateInProgress,
		VaultUpgradeStateUnknown,
		VaultUpgradeStateUpgraded,
	}
}

// ToPtr returns a *VaultUpgradeState pointing to the current value.
func (c VaultUpgradeState) ToPtr() *VaultUpgradeState {
	return &c
}
