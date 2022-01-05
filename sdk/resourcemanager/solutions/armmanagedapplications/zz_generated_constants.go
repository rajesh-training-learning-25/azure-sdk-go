//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagedapplications

const (
	module  = "armmanagedapplications"
	version = "v0.1.0"
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

// ToPtr returns a *ActionType pointing to the current value.
func (c ActionType) ToPtr() *ActionType {
	return &c
}

// ApplicationArtifactName - The managed application artifact name.
type ApplicationArtifactName string

const (
	ApplicationArtifactNameAuthorizations       ApplicationArtifactName = "Authorizations"
	ApplicationArtifactNameCustomRoleDefinition ApplicationArtifactName = "CustomRoleDefinition"
	ApplicationArtifactNameNotSpecified         ApplicationArtifactName = "NotSpecified"
	ApplicationArtifactNameViewDefinition       ApplicationArtifactName = "ViewDefinition"
)

// PossibleApplicationArtifactNameValues returns the possible values for the ApplicationArtifactName const type.
func PossibleApplicationArtifactNameValues() []ApplicationArtifactName {
	return []ApplicationArtifactName{
		ApplicationArtifactNameAuthorizations,
		ApplicationArtifactNameCustomRoleDefinition,
		ApplicationArtifactNameNotSpecified,
		ApplicationArtifactNameViewDefinition,
	}
}

// ToPtr returns a *ApplicationArtifactName pointing to the current value.
func (c ApplicationArtifactName) ToPtr() *ApplicationArtifactName {
	return &c
}

// ApplicationArtifactType - The managed application artifact type.
type ApplicationArtifactType string

const (
	ApplicationArtifactTypeNotSpecified ApplicationArtifactType = "NotSpecified"
	ApplicationArtifactTypeTemplate     ApplicationArtifactType = "Template"
	ApplicationArtifactTypeCustom       ApplicationArtifactType = "Custom"
)

// PossibleApplicationArtifactTypeValues returns the possible values for the ApplicationArtifactType const type.
func PossibleApplicationArtifactTypeValues() []ApplicationArtifactType {
	return []ApplicationArtifactType{
		ApplicationArtifactTypeNotSpecified,
		ApplicationArtifactTypeTemplate,
		ApplicationArtifactTypeCustom,
	}
}

// ToPtr returns a *ApplicationArtifactType pointing to the current value.
func (c ApplicationArtifactType) ToPtr() *ApplicationArtifactType {
	return &c
}

// ApplicationDefinitionArtifactName - The managed application artifact name.
type ApplicationDefinitionArtifactName string

const (
	ApplicationDefinitionArtifactNameApplicationResourceTemplate ApplicationDefinitionArtifactName = "ApplicationResourceTemplate"
	ApplicationDefinitionArtifactNameCreateUIDefinition          ApplicationDefinitionArtifactName = "CreateUiDefinition"
	ApplicationDefinitionArtifactNameMainTemplateParameters      ApplicationDefinitionArtifactName = "MainTemplateParameters"
	ApplicationDefinitionArtifactNameNotSpecified                ApplicationDefinitionArtifactName = "NotSpecified"
)

// PossibleApplicationDefinitionArtifactNameValues returns the possible values for the ApplicationDefinitionArtifactName const type.
func PossibleApplicationDefinitionArtifactNameValues() []ApplicationDefinitionArtifactName {
	return []ApplicationDefinitionArtifactName{
		ApplicationDefinitionArtifactNameApplicationResourceTemplate,
		ApplicationDefinitionArtifactNameCreateUIDefinition,
		ApplicationDefinitionArtifactNameMainTemplateParameters,
		ApplicationDefinitionArtifactNameNotSpecified,
	}
}

// ToPtr returns a *ApplicationDefinitionArtifactName pointing to the current value.
func (c ApplicationDefinitionArtifactName) ToPtr() *ApplicationDefinitionArtifactName {
	return &c
}

// ApplicationLockLevel - The managed application lock level.
type ApplicationLockLevel string

const (
	ApplicationLockLevelCanNotDelete ApplicationLockLevel = "CanNotDelete"
	ApplicationLockLevelReadOnly     ApplicationLockLevel = "ReadOnly"
	ApplicationLockLevelNone         ApplicationLockLevel = "None"
)

// PossibleApplicationLockLevelValues returns the possible values for the ApplicationLockLevel const type.
func PossibleApplicationLockLevelValues() []ApplicationLockLevel {
	return []ApplicationLockLevel{
		ApplicationLockLevelCanNotDelete,
		ApplicationLockLevelReadOnly,
		ApplicationLockLevelNone,
	}
}

// ToPtr returns a *ApplicationLockLevel pointing to the current value.
func (c ApplicationLockLevel) ToPtr() *ApplicationLockLevel {
	return &c
}

// ApplicationManagementMode - The management mode.
type ApplicationManagementMode string

const (
	ApplicationManagementModeManaged      ApplicationManagementMode = "Managed"
	ApplicationManagementModeNotSpecified ApplicationManagementMode = "NotSpecified"
	ApplicationManagementModeUnmanaged    ApplicationManagementMode = "Unmanaged"
)

// PossibleApplicationManagementModeValues returns the possible values for the ApplicationManagementMode const type.
func PossibleApplicationManagementModeValues() []ApplicationManagementMode {
	return []ApplicationManagementMode{
		ApplicationManagementModeManaged,
		ApplicationManagementModeNotSpecified,
		ApplicationManagementModeUnmanaged,
	}
}

// ToPtr returns a *ApplicationManagementMode pointing to the current value.
func (c ApplicationManagementMode) ToPtr() *ApplicationManagementMode {
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

// DeploymentMode - The deployment mode.
type DeploymentMode string

const (
	DeploymentModeComplete     DeploymentMode = "Complete"
	DeploymentModeIncremental  DeploymentMode = "Incremental"
	DeploymentModeNotSpecified DeploymentMode = "NotSpecified"
)

// PossibleDeploymentModeValues returns the possible values for the DeploymentMode const type.
func PossibleDeploymentModeValues() []DeploymentMode {
	return []DeploymentMode{
		DeploymentModeComplete,
		DeploymentModeIncremental,
		DeploymentModeNotSpecified,
	}
}

// ToPtr returns a *DeploymentMode pointing to the current value.
func (c DeploymentMode) ToPtr() *DeploymentMode {
	return &c
}

// JitApprovalMode - The Jit approval mode.
type JitApprovalMode string

const (
	JitApprovalModeAutoApprove   JitApprovalMode = "AutoApprove"
	JitApprovalModeManualApprove JitApprovalMode = "ManualApprove"
	JitApprovalModeNotSpecified  JitApprovalMode = "NotSpecified"
)

// PossibleJitApprovalModeValues returns the possible values for the JitApprovalMode const type.
func PossibleJitApprovalModeValues() []JitApprovalMode {
	return []JitApprovalMode{
		JitApprovalModeAutoApprove,
		JitApprovalModeManualApprove,
		JitApprovalModeNotSpecified,
	}
}

// ToPtr returns a *JitApprovalMode pointing to the current value.
func (c JitApprovalMode) ToPtr() *JitApprovalMode {
	return &c
}

// JitApproverType - The approver type.
type JitApproverType string

const (
	JitApproverTypeGroup JitApproverType = "group"
	JitApproverTypeUser  JitApproverType = "user"
)

// PossibleJitApproverTypeValues returns the possible values for the JitApproverType const type.
func PossibleJitApproverTypeValues() []JitApproverType {
	return []JitApproverType{
		JitApproverTypeGroup,
		JitApproverTypeUser,
	}
}

// ToPtr returns a *JitApproverType pointing to the current value.
func (c JitApproverType) ToPtr() *JitApproverType {
	return &c
}

// JitRequestState - The JIT request state.
type JitRequestState string

const (
	JitRequestStateApproved     JitRequestState = "Approved"
	JitRequestStateCanceled     JitRequestState = "Canceled"
	JitRequestStateDenied       JitRequestState = "Denied"
	JitRequestStateExpired      JitRequestState = "Expired"
	JitRequestStateFailed       JitRequestState = "Failed"
	JitRequestStateNotSpecified JitRequestState = "NotSpecified"
	JitRequestStatePending      JitRequestState = "Pending"
	JitRequestStateTimeout      JitRequestState = "Timeout"
)

// PossibleJitRequestStateValues returns the possible values for the JitRequestState const type.
func PossibleJitRequestStateValues() []JitRequestState {
	return []JitRequestState{
		JitRequestStateApproved,
		JitRequestStateCanceled,
		JitRequestStateDenied,
		JitRequestStateExpired,
		JitRequestStateFailed,
		JitRequestStateNotSpecified,
		JitRequestStatePending,
		JitRequestStateTimeout,
	}
}

// ToPtr returns a *JitRequestState pointing to the current value.
func (c JitRequestState) ToPtr() *JitRequestState {
	return &c
}

// JitSchedulingType - The JIT request scheduling type.
type JitSchedulingType string

const (
	JitSchedulingTypeNotSpecified JitSchedulingType = "NotSpecified"
	JitSchedulingTypeOnce         JitSchedulingType = "Once"
	JitSchedulingTypeRecurring    JitSchedulingType = "Recurring"
)

// PossibleJitSchedulingTypeValues returns the possible values for the JitSchedulingType const type.
func PossibleJitSchedulingTypeValues() []JitSchedulingType {
	return []JitSchedulingType{
		JitSchedulingTypeNotSpecified,
		JitSchedulingTypeOnce,
		JitSchedulingTypeRecurring,
	}
}

// ToPtr returns a *JitSchedulingType pointing to the current value.
func (c JitSchedulingType) ToPtr() *JitSchedulingType {
	return &c
}

// ManagedServiceIdentityType - Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                       ManagedServiceIdentityType = "None"
	ManagedServiceIdentityTypeSystemAssigned             ManagedServiceIdentityType = "SystemAssigned"
	ManagedServiceIdentityTypeSystemAssignedUserAssigned ManagedServiceIdentityType = "SystemAssigned,UserAssigned"
	ManagedServiceIdentityTypeUserAssigned               ManagedServiceIdentityType = "UserAssigned"
)

// PossibleManagedServiceIdentityTypeValues returns the possible values for the ManagedServiceIdentityType const type.
func PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType {
	return []ManagedServiceIdentityType{
		ManagedServiceIdentityTypeNone,
		ManagedServiceIdentityTypeSystemAssigned,
		ManagedServiceIdentityTypeSystemAssignedUserAssigned,
		ManagedServiceIdentityTypeUserAssigned,
	}
}

// ToPtr returns a *ManagedServiceIdentityType pointing to the current value.
func (c ManagedServiceIdentityType) ToPtr() *ManagedServiceIdentityType {
	return &c
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default value is "user,system"
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

// ToPtr returns a *Origin pointing to the current value.
func (c Origin) ToPtr() *Origin {
	return &c
}

// ProvisioningState - Provisioning status of the managed application.
type ProvisioningState string

const (
	ProvisioningStateAccepted     ProvisioningState = "Accepted"
	ProvisioningStateCanceled     ProvisioningState = "Canceled"
	ProvisioningStateDeleted      ProvisioningState = "Deleted"
	ProvisioningStateDeleting     ProvisioningState = "Deleting"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateNotSpecified ProvisioningState = "NotSpecified"
	ProvisioningStateRunning      ProvisioningState = "Running"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
	ProvisioningStateUpdating     ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateDeleted,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateNotSpecified,
		ProvisioningStateRunning,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// ToPtr returns a *ProvisioningState pointing to the current value.
func (c ProvisioningState) ToPtr() *ProvisioningState {
	return &c
}
