//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecuritydevops

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securitydevops/armsecuritydevops"
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

type ActionableRemediationState string

const (
	ActionableRemediationStateDisabled ActionableRemediationState = "Disabled"
	ActionableRemediationStateEnabled  ActionableRemediationState = "Enabled"
	ActionableRemediationStateNone     ActionableRemediationState = "None"
)

// PossibleActionableRemediationStateValues returns the possible values for the ActionableRemediationState const type.
func PossibleActionableRemediationStateValues() []ActionableRemediationState {
	return []ActionableRemediationState{
		ActionableRemediationStateDisabled,
		ActionableRemediationStateEnabled,
		ActionableRemediationStateNone,
	}
}

type AutoDiscovery string

const (
	AutoDiscoveryDisabled AutoDiscovery = "Disabled"
	AutoDiscoveryEnabled  AutoDiscovery = "Enabled"
)

// PossibleAutoDiscoveryValues returns the possible values for the AutoDiscovery const type.
func PossibleAutoDiscoveryValues() []AutoDiscovery {
	return []AutoDiscovery{
		AutoDiscoveryDisabled,
		AutoDiscoveryEnabled,
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

type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCanceled,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
	}
}

type RuleCategory string

const (
	RuleCategoryArtifacts    RuleCategory = "Artifacts"
	RuleCategoryCode         RuleCategory = "Code"
	RuleCategoryContainers   RuleCategory = "Containers"
	RuleCategoryDependencies RuleCategory = "Dependencies"
	RuleCategoryIaC          RuleCategory = "IaC"
	RuleCategorySecrets      RuleCategory = "Secrets"
)

// PossibleRuleCategoryValues returns the possible values for the RuleCategory const type.
func PossibleRuleCategoryValues() []RuleCategory {
	return []RuleCategory{
		RuleCategoryArtifacts,
		RuleCategoryCode,
		RuleCategoryContainers,
		RuleCategoryDependencies,
		RuleCategoryIaC,
		RuleCategorySecrets,
	}
}
