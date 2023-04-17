//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcognitiveservices

const (
	moduleName    = "armcognitiveservices"
	moduleVersion = "v1.4.1"
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

// CommitmentPlanProvisioningState - Gets the status of the resource at the time the operation was called.
type CommitmentPlanProvisioningState string

const (
	CommitmentPlanProvisioningStateAccepted  CommitmentPlanProvisioningState = "Accepted"
	CommitmentPlanProvisioningStateCanceled  CommitmentPlanProvisioningState = "Canceled"
	CommitmentPlanProvisioningStateCreating  CommitmentPlanProvisioningState = "Creating"
	CommitmentPlanProvisioningStateDeleting  CommitmentPlanProvisioningState = "Deleting"
	CommitmentPlanProvisioningStateFailed    CommitmentPlanProvisioningState = "Failed"
	CommitmentPlanProvisioningStateMoving    CommitmentPlanProvisioningState = "Moving"
	CommitmentPlanProvisioningStateSucceeded CommitmentPlanProvisioningState = "Succeeded"
)

// PossibleCommitmentPlanProvisioningStateValues returns the possible values for the CommitmentPlanProvisioningState const type.
func PossibleCommitmentPlanProvisioningStateValues() []CommitmentPlanProvisioningState {
	return []CommitmentPlanProvisioningState{
		CommitmentPlanProvisioningStateAccepted,
		CommitmentPlanProvisioningStateCanceled,
		CommitmentPlanProvisioningStateCreating,
		CommitmentPlanProvisioningStateDeleting,
		CommitmentPlanProvisioningStateFailed,
		CommitmentPlanProvisioningStateMoving,
		CommitmentPlanProvisioningStateSucceeded,
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

// DeploymentProvisioningState - Gets the status of the resource at the time the operation was called.
type DeploymentProvisioningState string

const (
	DeploymentProvisioningStateAccepted  DeploymentProvisioningState = "Accepted"
	DeploymentProvisioningStateCreating  DeploymentProvisioningState = "Creating"
	DeploymentProvisioningStateDeleting  DeploymentProvisioningState = "Deleting"
	DeploymentProvisioningStateFailed    DeploymentProvisioningState = "Failed"
	DeploymentProvisioningStateMoving    DeploymentProvisioningState = "Moving"
	DeploymentProvisioningStateSucceeded DeploymentProvisioningState = "Succeeded"
)

// PossibleDeploymentProvisioningStateValues returns the possible values for the DeploymentProvisioningState const type.
func PossibleDeploymentProvisioningStateValues() []DeploymentProvisioningState {
	return []DeploymentProvisioningState{
		DeploymentProvisioningStateAccepted,
		DeploymentProvisioningStateCreating,
		DeploymentProvisioningStateDeleting,
		DeploymentProvisioningStateFailed,
		DeploymentProvisioningStateMoving,
		DeploymentProvisioningStateSucceeded,
	}
}

// DeploymentScaleType - Deployment scale type.
type DeploymentScaleType string

const (
	DeploymentScaleTypeManual   DeploymentScaleType = "Manual"
	DeploymentScaleTypeStandard DeploymentScaleType = "Standard"
)

// PossibleDeploymentScaleTypeValues returns the possible values for the DeploymentScaleType const type.
func PossibleDeploymentScaleTypeValues() []DeploymentScaleType {
	return []DeploymentScaleType{
		DeploymentScaleTypeManual,
		DeploymentScaleTypeStandard,
	}
}

// HostingModel - Account hosting model.
type HostingModel string

const (
	HostingModelConnectedContainer    HostingModel = "ConnectedContainer"
	HostingModelDisconnectedContainer HostingModel = "DisconnectedContainer"
	HostingModelWeb                   HostingModel = "Web"
)

// PossibleHostingModelValues returns the possible values for the HostingModel const type.
func PossibleHostingModelValues() []HostingModel {
	return []HostingModel{
		HostingModelConnectedContainer,
		HostingModelDisconnectedContainer,
		HostingModelWeb,
	}
}

// KeyName - key name to generate (Key1|Key2)
type KeyName string

const (
	KeyNameKey1 KeyName = "Key1"
	KeyNameKey2 KeyName = "Key2"
)

// PossibleKeyNameValues returns the possible values for the KeyName const type.
func PossibleKeyNameValues() []KeyName {
	return []KeyName{
		KeyNameKey1,
		KeyNameKey2,
	}
}

// KeySource - Enumerates the possible value of keySource for Encryption
type KeySource string

const (
	KeySourceMicrosoftCognitiveServices KeySource = "Microsoft.CognitiveServices"
	KeySourceMicrosoftKeyVault          KeySource = "Microsoft.KeyVault"
)

// PossibleKeySourceValues returns the possible values for the KeySource const type.
func PossibleKeySourceValues() []KeySource {
	return []KeySource{
		KeySourceMicrosoftCognitiveServices,
		KeySourceMicrosoftKeyVault,
	}
}

// ModelLifecycleStatus - Model lifecycle status.
type ModelLifecycleStatus string

const (
	ModelLifecycleStatusGenerallyAvailable ModelLifecycleStatus = "GenerallyAvailable"
	ModelLifecycleStatusPreview            ModelLifecycleStatus = "Preview"
)

// PossibleModelLifecycleStatusValues returns the possible values for the ModelLifecycleStatus const type.
func PossibleModelLifecycleStatusValues() []ModelLifecycleStatus {
	return []ModelLifecycleStatus{
		ModelLifecycleStatusGenerallyAvailable,
		ModelLifecycleStatusPreview,
	}
}

// NetworkRuleAction - The default action when no rule from ipRules and from virtualNetworkRules match. This is only used
// after the bypass property has been evaluated.
type NetworkRuleAction string

const (
	NetworkRuleActionAllow NetworkRuleAction = "Allow"
	NetworkRuleActionDeny  NetworkRuleAction = "Deny"
)

// PossibleNetworkRuleActionValues returns the possible values for the NetworkRuleAction const type.
func PossibleNetworkRuleActionValues() []NetworkRuleAction {
	return []NetworkRuleAction{
		NetworkRuleActionAllow,
		NetworkRuleActionDeny,
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

// PrivateEndpointConnectionProvisioningState - The current provisioning state.
type PrivateEndpointConnectionProvisioningState string

const (
	PrivateEndpointConnectionProvisioningStateCreating  PrivateEndpointConnectionProvisioningState = "Creating"
	PrivateEndpointConnectionProvisioningStateDeleting  PrivateEndpointConnectionProvisioningState = "Deleting"
	PrivateEndpointConnectionProvisioningStateFailed    PrivateEndpointConnectionProvisioningState = "Failed"
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = "Succeeded"
)

// PossiblePrivateEndpointConnectionProvisioningStateValues returns the possible values for the PrivateEndpointConnectionProvisioningState const type.
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return []PrivateEndpointConnectionProvisioningState{
		PrivateEndpointConnectionProvisioningStateCreating,
		PrivateEndpointConnectionProvisioningStateDeleting,
		PrivateEndpointConnectionProvisioningStateFailed,
		PrivateEndpointConnectionProvisioningStateSucceeded,
	}
}

// PrivateEndpointServiceConnectionStatus - The private endpoint connection status.
type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusApproved PrivateEndpointServiceConnectionStatus = "Approved"
	PrivateEndpointServiceConnectionStatusPending  PrivateEndpointServiceConnectionStatus = "Pending"
	PrivateEndpointServiceConnectionStatusRejected PrivateEndpointServiceConnectionStatus = "Rejected"
)

// PossiblePrivateEndpointServiceConnectionStatusValues returns the possible values for the PrivateEndpointServiceConnectionStatus const type.
func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return []PrivateEndpointServiceConnectionStatus{
		PrivateEndpointServiceConnectionStatusApproved,
		PrivateEndpointServiceConnectionStatusPending,
		PrivateEndpointServiceConnectionStatusRejected,
	}
}

// ProvisioningState - Gets the status of the cognitive services account at the time the operation was called.
type ProvisioningState string

const (
	ProvisioningStateAccepted     ProvisioningState = "Accepted"
	ProvisioningStateCreating     ProvisioningState = "Creating"
	ProvisioningStateDeleting     ProvisioningState = "Deleting"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateMoving       ProvisioningState = "Moving"
	ProvisioningStateResolvingDNS ProvisioningState = "ResolvingDNS"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCreating,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateMoving,
		ProvisioningStateResolvingDNS,
		ProvisioningStateSucceeded,
	}
}

// PublicNetworkAccess - Whether or not public endpoint access is allowed for this account.
type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled  PublicNetworkAccess = "Enabled"
)

// PossiblePublicNetworkAccessValues returns the possible values for the PublicNetworkAccess const type.
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return []PublicNetworkAccess{
		PublicNetworkAccessDisabled,
		PublicNetworkAccessEnabled,
	}
}

// QuotaUsageStatus - Cognitive Services account quota usage status.
type QuotaUsageStatus string

const (
	QuotaUsageStatusBlocked   QuotaUsageStatus = "Blocked"
	QuotaUsageStatusInOverage QuotaUsageStatus = "InOverage"
	QuotaUsageStatusIncluded  QuotaUsageStatus = "Included"
	QuotaUsageStatusUnknown   QuotaUsageStatus = "Unknown"
)

// PossibleQuotaUsageStatusValues returns the possible values for the QuotaUsageStatus const type.
func PossibleQuotaUsageStatusValues() []QuotaUsageStatus {
	return []QuotaUsageStatus{
		QuotaUsageStatusBlocked,
		QuotaUsageStatusInOverage,
		QuotaUsageStatusIncluded,
		QuotaUsageStatusUnknown,
	}
}

// ResourceIdentityType - The identity type.
type ResourceIdentityType string

const (
	ResourceIdentityTypeNone                       ResourceIdentityType = "None"
	ResourceIdentityTypeSystemAssigned             ResourceIdentityType = "SystemAssigned"
	ResourceIdentityTypeUserAssigned               ResourceIdentityType = "UserAssigned"
	ResourceIdentityTypeSystemAssignedUserAssigned ResourceIdentityType = "SystemAssigned, UserAssigned"
)

// PossibleResourceIdentityTypeValues returns the possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{
		ResourceIdentityTypeNone,
		ResourceIdentityTypeSystemAssigned,
		ResourceIdentityTypeUserAssigned,
		ResourceIdentityTypeSystemAssignedUserAssigned,
	}
}

// ResourceSKURestrictionsReasonCode - The reason for restriction.
type ResourceSKURestrictionsReasonCode string

const (
	ResourceSKURestrictionsReasonCodeNotAvailableForSubscription ResourceSKURestrictionsReasonCode = "NotAvailableForSubscription"
	ResourceSKURestrictionsReasonCodeQuotaID                     ResourceSKURestrictionsReasonCode = "QuotaId"
)

// PossibleResourceSKURestrictionsReasonCodeValues returns the possible values for the ResourceSKURestrictionsReasonCode const type.
func PossibleResourceSKURestrictionsReasonCodeValues() []ResourceSKURestrictionsReasonCode {
	return []ResourceSKURestrictionsReasonCode{
		ResourceSKURestrictionsReasonCodeNotAvailableForSubscription,
		ResourceSKURestrictionsReasonCodeQuotaID,
	}
}

// ResourceSKURestrictionsType - The type of restrictions.
type ResourceSKURestrictionsType string

const (
	ResourceSKURestrictionsTypeLocation ResourceSKURestrictionsType = "Location"
	ResourceSKURestrictionsTypeZone     ResourceSKURestrictionsType = "Zone"
)

// PossibleResourceSKURestrictionsTypeValues returns the possible values for the ResourceSKURestrictionsType const type.
func PossibleResourceSKURestrictionsTypeValues() []ResourceSKURestrictionsType {
	return []ResourceSKURestrictionsType{
		ResourceSKURestrictionsTypeLocation,
		ResourceSKURestrictionsTypeZone,
	}
}

// RoutingMethods - Multiregion routing methods.
type RoutingMethods string

const (
	RoutingMethodsPerformance RoutingMethods = "Performance"
	RoutingMethodsPriority    RoutingMethods = "Priority"
	RoutingMethodsWeighted    RoutingMethods = "Weighted"
)

// PossibleRoutingMethodsValues returns the possible values for the RoutingMethods const type.
func PossibleRoutingMethodsValues() []RoutingMethods {
	return []RoutingMethods{
		RoutingMethodsPerformance,
		RoutingMethodsPriority,
		RoutingMethodsWeighted,
	}
}

// SKUTier - This field is required to be implemented by the Resource Provider if the service has more than one tier, but
// is not required on a PUT.
type SKUTier string

const (
	SKUTierBasic      SKUTier = "Basic"
	SKUTierEnterprise SKUTier = "Enterprise"
	SKUTierFree       SKUTier = "Free"
	SKUTierPremium    SKUTier = "Premium"
	SKUTierStandard   SKUTier = "Standard"
)

// PossibleSKUTierValues returns the possible values for the SKUTier const type.
func PossibleSKUTierValues() []SKUTier {
	return []SKUTier{
		SKUTierBasic,
		SKUTierEnterprise,
		SKUTierFree,
		SKUTierPremium,
		SKUTierStandard,
	}
}

// UnitType - The unit of the metric.
type UnitType string

const (
	UnitTypeBytes          UnitType = "Bytes"
	UnitTypeBytesPerSecond UnitType = "BytesPerSecond"
	UnitTypeCount          UnitType = "Count"
	UnitTypeCountPerSecond UnitType = "CountPerSecond"
	UnitTypeMilliseconds   UnitType = "Milliseconds"
	UnitTypePercent        UnitType = "Percent"
	UnitTypeSeconds        UnitType = "Seconds"
)

// PossibleUnitTypeValues returns the possible values for the UnitType const type.
func PossibleUnitTypeValues() []UnitType {
	return []UnitType{
		UnitTypeBytes,
		UnitTypeBytesPerSecond,
		UnitTypeCount,
		UnitTypeCountPerSecond,
		UnitTypeMilliseconds,
		UnitTypePercent,
		UnitTypeSeconds,
	}
}
