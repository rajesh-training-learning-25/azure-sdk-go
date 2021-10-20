//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventhub

const (
	module  = "armeventhub"
	version = "v0.1.0"
)

type AccessRights string

const (
	AccessRightsListen AccessRights = "Listen"
	AccessRightsManage AccessRights = "Manage"
	AccessRightsSend   AccessRights = "Send"
)

// PossibleAccessRightsValues returns the possible values for the AccessRights const type.
func PossibleAccessRightsValues() []AccessRights {
	return []AccessRights{
		AccessRightsListen,
		AccessRightsManage,
		AccessRightsSend,
	}
}

// ToPtr returns a *AccessRights pointing to the current value.
func (c AccessRights) ToPtr() *AccessRights {
	return &c
}

// ClusterSKUName - Name of this SKU.
type ClusterSKUName string

const (
	ClusterSKUNameDedicated ClusterSKUName = "Dedicated"
)

// PossibleClusterSKUNameValues returns the possible values for the ClusterSKUName const type.
func PossibleClusterSKUNameValues() []ClusterSKUName {
	return []ClusterSKUName{
		ClusterSKUNameDedicated,
	}
}

// ToPtr returns a *ClusterSKUName pointing to the current value.
func (c ClusterSKUName) ToPtr() *ClusterSKUName {
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

// DefaultAction - Default Action for Network Rule Set
type DefaultAction string

const (
	DefaultActionAllow DefaultAction = "Allow"
	DefaultActionDeny  DefaultAction = "Deny"
)

// PossibleDefaultActionValues returns the possible values for the DefaultAction const type.
func PossibleDefaultActionValues() []DefaultAction {
	return []DefaultAction{
		DefaultActionAllow,
		DefaultActionDeny,
	}
}

// ToPtr returns a *DefaultAction pointing to the current value.
func (c DefaultAction) ToPtr() *DefaultAction {
	return &c
}

// EncodingCaptureDescription - Enumerates the possible values for the encoding format of capture description. Note: 'AvroDeflate' will be deprecated in
// New API Version
type EncodingCaptureDescription string

const (
	EncodingCaptureDescriptionAvro        EncodingCaptureDescription = "Avro"
	EncodingCaptureDescriptionAvroDeflate EncodingCaptureDescription = "AvroDeflate"
)

// PossibleEncodingCaptureDescriptionValues returns the possible values for the EncodingCaptureDescription const type.
func PossibleEncodingCaptureDescriptionValues() []EncodingCaptureDescription {
	return []EncodingCaptureDescription{
		EncodingCaptureDescriptionAvro,
		EncodingCaptureDescriptionAvroDeflate,
	}
}

// ToPtr returns a *EncodingCaptureDescription pointing to the current value.
func (c EncodingCaptureDescription) ToPtr() *EncodingCaptureDescription {
	return &c
}

// EndPointProvisioningState - Provisioning state of the Private Endpoint Connection.
type EndPointProvisioningState string

const (
	EndPointProvisioningStateCanceled  EndPointProvisioningState = "Canceled"
	EndPointProvisioningStateCreating  EndPointProvisioningState = "Creating"
	EndPointProvisioningStateDeleting  EndPointProvisioningState = "Deleting"
	EndPointProvisioningStateFailed    EndPointProvisioningState = "Failed"
	EndPointProvisioningStateSucceeded EndPointProvisioningState = "Succeeded"
	EndPointProvisioningStateUpdating  EndPointProvisioningState = "Updating"
)

// PossibleEndPointProvisioningStateValues returns the possible values for the EndPointProvisioningState const type.
func PossibleEndPointProvisioningStateValues() []EndPointProvisioningState {
	return []EndPointProvisioningState{
		EndPointProvisioningStateCanceled,
		EndPointProvisioningStateCreating,
		EndPointProvisioningStateDeleting,
		EndPointProvisioningStateFailed,
		EndPointProvisioningStateSucceeded,
		EndPointProvisioningStateUpdating,
	}
}

// ToPtr returns a *EndPointProvisioningState pointing to the current value.
func (c EndPointProvisioningState) ToPtr() *EndPointProvisioningState {
	return &c
}

// EntityStatus - Enumerates the possible values for the status of the Event Hub.
type EntityStatus string

const (
	EntityStatusActive          EntityStatus = "Active"
	EntityStatusDisabled        EntityStatus = "Disabled"
	EntityStatusRestoring       EntityStatus = "Restoring"
	EntityStatusSendDisabled    EntityStatus = "SendDisabled"
	EntityStatusReceiveDisabled EntityStatus = "ReceiveDisabled"
	EntityStatusCreating        EntityStatus = "Creating"
	EntityStatusDeleting        EntityStatus = "Deleting"
	EntityStatusRenaming        EntityStatus = "Renaming"
	EntityStatusUnknown         EntityStatus = "Unknown"
)

// PossibleEntityStatusValues returns the possible values for the EntityStatus const type.
func PossibleEntityStatusValues() []EntityStatus {
	return []EntityStatus{
		EntityStatusActive,
		EntityStatusDisabled,
		EntityStatusRestoring,
		EntityStatusSendDisabled,
		EntityStatusReceiveDisabled,
		EntityStatusCreating,
		EntityStatusDeleting,
		EntityStatusRenaming,
		EntityStatusUnknown,
	}
}

// ToPtr returns a *EntityStatus pointing to the current value.
func (c EntityStatus) ToPtr() *EntityStatus {
	return &c
}

// KeyType - The access key to regenerate.
type KeyType string

const (
	KeyTypePrimaryKey   KeyType = "PrimaryKey"
	KeyTypeSecondaryKey KeyType = "SecondaryKey"
)

// PossibleKeyTypeValues returns the possible values for the KeyType const type.
func PossibleKeyTypeValues() []KeyType {
	return []KeyType{
		KeyTypePrimaryKey,
		KeyTypeSecondaryKey,
	}
}

// ToPtr returns a *KeyType pointing to the current value.
func (c KeyType) ToPtr() *KeyType {
	return &c
}

// ManagedServiceIdentityType - Type of managed service identity.
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeSystemAssigned             ManagedServiceIdentityType = "SystemAssigned"
	ManagedServiceIdentityTypeUserAssigned               ManagedServiceIdentityType = "UserAssigned"
	ManagedServiceIdentityTypeSystemAssignedUserAssigned ManagedServiceIdentityType = "SystemAssigned, UserAssigned"
	ManagedServiceIdentityTypeNone                       ManagedServiceIdentityType = "None"
)

// PossibleManagedServiceIdentityTypeValues returns the possible values for the ManagedServiceIdentityType const type.
func PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType {
	return []ManagedServiceIdentityType{
		ManagedServiceIdentityTypeSystemAssigned,
		ManagedServiceIdentityTypeUserAssigned,
		ManagedServiceIdentityTypeSystemAssignedUserAssigned,
		ManagedServiceIdentityTypeNone,
	}
}

// ToPtr returns a *ManagedServiceIdentityType pointing to the current value.
func (c ManagedServiceIdentityType) ToPtr() *ManagedServiceIdentityType {
	return &c
}

// NetworkRuleIPAction - The IP Filter Action
type NetworkRuleIPAction string

const (
	NetworkRuleIPActionAllow NetworkRuleIPAction = "Allow"
)

// PossibleNetworkRuleIPActionValues returns the possible values for the NetworkRuleIPAction const type.
func PossibleNetworkRuleIPActionValues() []NetworkRuleIPAction {
	return []NetworkRuleIPAction{
		NetworkRuleIPActionAllow,
	}
}

// ToPtr returns a *NetworkRuleIPAction pointing to the current value.
func (c NetworkRuleIPAction) ToPtr() *NetworkRuleIPAction {
	return &c
}

// PrivateLinkConnectionStatus - Status of the connection.
type PrivateLinkConnectionStatus string

const (
	PrivateLinkConnectionStatusApproved     PrivateLinkConnectionStatus = "Approved"
	PrivateLinkConnectionStatusDisconnected PrivateLinkConnectionStatus = "Disconnected"
	PrivateLinkConnectionStatusPending      PrivateLinkConnectionStatus = "Pending"
	PrivateLinkConnectionStatusRejected     PrivateLinkConnectionStatus = "Rejected"
)

// PossiblePrivateLinkConnectionStatusValues returns the possible values for the PrivateLinkConnectionStatus const type.
func PossiblePrivateLinkConnectionStatusValues() []PrivateLinkConnectionStatus {
	return []PrivateLinkConnectionStatus{
		PrivateLinkConnectionStatusApproved,
		PrivateLinkConnectionStatusDisconnected,
		PrivateLinkConnectionStatusPending,
		PrivateLinkConnectionStatusRejected,
	}
}

// ToPtr returns a *PrivateLinkConnectionStatus pointing to the current value.
func (c PrivateLinkConnectionStatus) ToPtr() *PrivateLinkConnectionStatus {
	return &c
}

// ProvisioningStateDR - Provisioning state of the Alias(Disaster Recovery configuration) - possible values 'Accepted' or 'Succeeded' or 'Failed'
type ProvisioningStateDR string

const (
	ProvisioningStateDRAccepted  ProvisioningStateDR = "Accepted"
	ProvisioningStateDRSucceeded ProvisioningStateDR = "Succeeded"
	ProvisioningStateDRFailed    ProvisioningStateDR = "Failed"
)

// PossibleProvisioningStateDRValues returns the possible values for the ProvisioningStateDR const type.
func PossibleProvisioningStateDRValues() []ProvisioningStateDR {
	return []ProvisioningStateDR{
		ProvisioningStateDRAccepted,
		ProvisioningStateDRSucceeded,
		ProvisioningStateDRFailed,
	}
}

// ToPtr returns a *ProvisioningStateDR pointing to the current value.
func (c ProvisioningStateDR) ToPtr() *ProvisioningStateDR {
	return &c
}

// PublicNetworkAccessFlag - This determines if traffic is allowed over public network. By default it is enabled.
type PublicNetworkAccessFlag string

const (
	PublicNetworkAccessFlagDisabled PublicNetworkAccessFlag = "Disabled"
	PublicNetworkAccessFlagEnabled  PublicNetworkAccessFlag = "Enabled"
)

// PossiblePublicNetworkAccessFlagValues returns the possible values for the PublicNetworkAccessFlag const type.
func PossiblePublicNetworkAccessFlagValues() []PublicNetworkAccessFlag {
	return []PublicNetworkAccessFlag{
		PublicNetworkAccessFlagDisabled,
		PublicNetworkAccessFlagEnabled,
	}
}

// ToPtr returns a *PublicNetworkAccessFlag pointing to the current value.
func (c PublicNetworkAccessFlag) ToPtr() *PublicNetworkAccessFlag {
	return &c
}

// RoleDisasterRecovery - role of namespace in GEO DR - possible values 'Primary' or 'PrimaryNotReplicating' or 'Secondary'
type RoleDisasterRecovery string

const (
	RoleDisasterRecoveryPrimary               RoleDisasterRecovery = "Primary"
	RoleDisasterRecoveryPrimaryNotReplicating RoleDisasterRecovery = "PrimaryNotReplicating"
	RoleDisasterRecoverySecondary             RoleDisasterRecovery = "Secondary"
)

// PossibleRoleDisasterRecoveryValues returns the possible values for the RoleDisasterRecovery const type.
func PossibleRoleDisasterRecoveryValues() []RoleDisasterRecovery {
	return []RoleDisasterRecovery{
		RoleDisasterRecoveryPrimary,
		RoleDisasterRecoveryPrimaryNotReplicating,
		RoleDisasterRecoverySecondary,
	}
}

// ToPtr returns a *RoleDisasterRecovery pointing to the current value.
func (c RoleDisasterRecovery) ToPtr() *RoleDisasterRecovery {
	return &c
}

// SKUName - Name of this SKU.
type SKUName string

const (
	SKUNameBasic    SKUName = "Basic"
	SKUNamePremium  SKUName = "Premium"
	SKUNameStandard SKUName = "Standard"
)

// PossibleSKUNameValues returns the possible values for the SKUName const type.
func PossibleSKUNameValues() []SKUName {
	return []SKUName{
		SKUNameBasic,
		SKUNamePremium,
		SKUNameStandard,
	}
}

// ToPtr returns a *SKUName pointing to the current value.
func (c SKUName) ToPtr() *SKUName {
	return &c
}

// SKUTier - The billing tier of this particular SKU.
type SKUTier string

const (
	SKUTierBasic    SKUTier = "Basic"
	SKUTierPremium  SKUTier = "Premium"
	SKUTierStandard SKUTier = "Standard"
)

// PossibleSKUTierValues returns the possible values for the SKUTier const type.
func PossibleSKUTierValues() []SKUTier {
	return []SKUTier{
		SKUTierBasic,
		SKUTierPremium,
		SKUTierStandard,
	}
}

// ToPtr returns a *SKUTier pointing to the current value.
func (c SKUTier) ToPtr() *SKUTier {
	return &c
}

// UnavailableReason - Specifies the reason for the unavailability of the service.
type UnavailableReason string

const (
	UnavailableReasonNone                                  UnavailableReason = "None"
	UnavailableReasonInvalidName                           UnavailableReason = "InvalidName"
	UnavailableReasonSubscriptionIsDisabled                UnavailableReason = "SubscriptionIsDisabled"
	UnavailableReasonNameInUse                             UnavailableReason = "NameInUse"
	UnavailableReasonNameInLockdown                        UnavailableReason = "NameInLockdown"
	UnavailableReasonTooManyNamespaceInCurrentSubscription UnavailableReason = "TooManyNamespaceInCurrentSubscription"
)

// PossibleUnavailableReasonValues returns the possible values for the UnavailableReason const type.
func PossibleUnavailableReasonValues() []UnavailableReason {
	return []UnavailableReason{
		UnavailableReasonNone,
		UnavailableReasonInvalidName,
		UnavailableReasonSubscriptionIsDisabled,
		UnavailableReasonNameInUse,
		UnavailableReasonNameInLockdown,
		UnavailableReasonTooManyNamespaceInCurrentSubscription,
	}
}

// ToPtr returns a *UnavailableReason pointing to the current value.
func (c UnavailableReason) ToPtr() *UnavailableReason {
	return &c
}
