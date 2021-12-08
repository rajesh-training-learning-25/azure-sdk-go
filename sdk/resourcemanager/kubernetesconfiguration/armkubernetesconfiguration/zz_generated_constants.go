//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkubernetesconfiguration

const (
	module  = "armkubernetesconfiguration"
	version = "v0.1.0"
)

// ClusterTypes - Cluster types
type ClusterTypes string

const (
	ClusterTypesConnectedClusters ClusterTypes = "connectedClusters"
	ClusterTypesManagedClusters   ClusterTypes = "managedClusters"
)

// PossibleClusterTypesValues returns the possible values for the ClusterTypes const type.
func PossibleClusterTypesValues() []ClusterTypes {
	return []ClusterTypes{
		ClusterTypesConnectedClusters,
		ClusterTypesManagedClusters,
	}
}

// ToPtr returns a *ClusterTypes pointing to the current value.
func (c ClusterTypes) ToPtr() *ClusterTypes {
	return &c
}

// ComplianceStateType - The compliance state of the configuration.
type ComplianceStateType string

const (
	ComplianceStateTypeCompliant    ComplianceStateType = "Compliant"
	ComplianceStateTypeFailed       ComplianceStateType = "Failed"
	ComplianceStateTypeInstalled    ComplianceStateType = "Installed"
	ComplianceStateTypeNoncompliant ComplianceStateType = "Noncompliant"
	ComplianceStateTypePending      ComplianceStateType = "Pending"
)

// PossibleComplianceStateTypeValues returns the possible values for the ComplianceStateType const type.
func PossibleComplianceStateTypeValues() []ComplianceStateType {
	return []ComplianceStateType{
		ComplianceStateTypeCompliant,
		ComplianceStateTypeFailed,
		ComplianceStateTypeInstalled,
		ComplianceStateTypeNoncompliant,
		ComplianceStateTypePending,
	}
}

// ToPtr returns a *ComplianceStateType pointing to the current value.
func (c ComplianceStateType) ToPtr() *ComplianceStateType {
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

type Enum0 string

const (
	Enum0MicrosoftContainerService Enum0 = "Microsoft.ContainerService"
	Enum0MicrosoftKubernetes       Enum0 = "Microsoft.Kubernetes"
)

// PossibleEnum0Values returns the possible values for the Enum0 const type.
func PossibleEnum0Values() []Enum0 {
	return []Enum0{
		Enum0MicrosoftContainerService,
		Enum0MicrosoftKubernetes,
	}
}

// ToPtr returns a *Enum0 pointing to the current value.
func (c Enum0) ToPtr() *Enum0 {
	return &c
}

type Enum1 string

const (
	Enum1ConnectedClusters Enum1 = "connectedClusters"
	Enum1ManagedClusters   Enum1 = "managedClusters"
)

// PossibleEnum1Values returns the possible values for the Enum1 const type.
func PossibleEnum1Values() []Enum1 {
	return []Enum1{
		Enum1ConnectedClusters,
		Enum1ManagedClusters,
	}
}

// ToPtr returns a *Enum1 pointing to the current value.
func (c Enum1) ToPtr() *Enum1 {
	return &c
}

// FluxComplianceState - Compliance state of the cluster object.
type FluxComplianceState string

const (
	FluxComplianceStateCompliant    FluxComplianceState = "Compliant"
	FluxComplianceStateNonCompliant FluxComplianceState = "Non-Compliant"
	FluxComplianceStatePending      FluxComplianceState = "Pending"
	FluxComplianceStateSuspended    FluxComplianceState = "Suspended"
	FluxComplianceStateUnknown      FluxComplianceState = "Unknown"
)

// PossibleFluxComplianceStateValues returns the possible values for the FluxComplianceState const type.
func PossibleFluxComplianceStateValues() []FluxComplianceState {
	return []FluxComplianceState{
		FluxComplianceStateCompliant,
		FluxComplianceStateNonCompliant,
		FluxComplianceStatePending,
		FluxComplianceStateSuspended,
		FluxComplianceStateUnknown,
	}
}

// ToPtr returns a *FluxComplianceState pointing to the current value.
func (c FluxComplianceState) ToPtr() *FluxComplianceState {
	return &c
}

// KustomizationValidationType - Specify whether to validate the Kubernetes objects referenced in the Kustomization before applying them to the cluster.
type KustomizationValidationType string

const (
	KustomizationValidationTypeClient KustomizationValidationType = "client"
	KustomizationValidationTypeNone   KustomizationValidationType = "none"
	KustomizationValidationTypeServer KustomizationValidationType = "server"
)

// PossibleKustomizationValidationTypeValues returns the possible values for the KustomizationValidationType const type.
func PossibleKustomizationValidationTypeValues() []KustomizationValidationType {
	return []KustomizationValidationType{
		KustomizationValidationTypeClient,
		KustomizationValidationTypeNone,
		KustomizationValidationTypeServer,
	}
}

// ToPtr returns a *KustomizationValidationType pointing to the current value.
func (c KustomizationValidationType) ToPtr() *KustomizationValidationType {
	return &c
}

// LevelType - Level of the status.
type LevelType string

const (
	LevelTypeError       LevelType = "Error"
	LevelTypeInformation LevelType = "Information"
	LevelTypeWarning     LevelType = "Warning"
)

// PossibleLevelTypeValues returns the possible values for the LevelType const type.
func PossibleLevelTypeValues() []LevelType {
	return []LevelType{
		LevelTypeError,
		LevelTypeInformation,
		LevelTypeWarning,
	}
}

// ToPtr returns a *LevelType pointing to the current value.
func (c LevelType) ToPtr() *LevelType {
	return &c
}

// MessageLevelType - Level of the message.
type MessageLevelType string

const (
	MessageLevelTypeError       MessageLevelType = "Error"
	MessageLevelTypeInformation MessageLevelType = "Information"
	MessageLevelTypeWarning     MessageLevelType = "Warning"
)

// PossibleMessageLevelTypeValues returns the possible values for the MessageLevelType const type.
func PossibleMessageLevelTypeValues() []MessageLevelType {
	return []MessageLevelType{
		MessageLevelTypeError,
		MessageLevelTypeInformation,
		MessageLevelTypeWarning,
	}
}

// ToPtr returns a *MessageLevelType pointing to the current value.
func (c MessageLevelType) ToPtr() *MessageLevelType {
	return &c
}

// OperatorScopeType - Scope at which the operator will be installed.
type OperatorScopeType string

const (
	OperatorScopeTypeCluster   OperatorScopeType = "cluster"
	OperatorScopeTypeNamespace OperatorScopeType = "namespace"
)

// PossibleOperatorScopeTypeValues returns the possible values for the OperatorScopeType const type.
func PossibleOperatorScopeTypeValues() []OperatorScopeType {
	return []OperatorScopeType{
		OperatorScopeTypeCluster,
		OperatorScopeTypeNamespace,
	}
}

// ToPtr returns a *OperatorScopeType pointing to the current value.
func (c OperatorScopeType) ToPtr() *OperatorScopeType {
	return &c
}

// OperatorType - Type of the operator
type OperatorType string

const (
	OperatorTypeFlux OperatorType = "Flux"
)

// PossibleOperatorTypeValues returns the possible values for the OperatorType const type.
func PossibleOperatorTypeValues() []OperatorType {
	return []OperatorType{
		OperatorTypeFlux,
	}
}

// ToPtr returns a *OperatorType pointing to the current value.
func (c OperatorType) ToPtr() *OperatorType {
	return &c
}

// ProvisioningState - The provisioning state of the resource.
type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// ToPtr returns a *ProvisioningState pointing to the current value.
func (c ProvisioningState) ToPtr() *ProvisioningState {
	return &c
}

// ProvisioningStateType - The provisioning state of the resource provider.
type ProvisioningStateType string

const (
	ProvisioningStateTypeAccepted  ProvisioningStateType = "Accepted"
	ProvisioningStateTypeDeleting  ProvisioningStateType = "Deleting"
	ProvisioningStateTypeFailed    ProvisioningStateType = "Failed"
	ProvisioningStateTypeRunning   ProvisioningStateType = "Running"
	ProvisioningStateTypeSucceeded ProvisioningStateType = "Succeeded"
)

// PossibleProvisioningStateTypeValues returns the possible values for the ProvisioningStateType const type.
func PossibleProvisioningStateTypeValues() []ProvisioningStateType {
	return []ProvisioningStateType{
		ProvisioningStateTypeAccepted,
		ProvisioningStateTypeDeleting,
		ProvisioningStateTypeFailed,
		ProvisioningStateTypeRunning,
		ProvisioningStateTypeSucceeded,
	}
}

// ToPtr returns a *ProvisioningStateType pointing to the current value.
func (c ProvisioningStateType) ToPtr() *ProvisioningStateType {
	return &c
}

// ScopeType - Scope at which the configuration will be installed.
type ScopeType string

const (
	ScopeTypeCluster   ScopeType = "cluster"
	ScopeTypeNamespace ScopeType = "namespace"
)

// PossibleScopeTypeValues returns the possible values for the ScopeType const type.
func PossibleScopeTypeValues() []ScopeType {
	return []ScopeType{
		ScopeTypeCluster,
		ScopeTypeNamespace,
	}
}

// ToPtr returns a *ScopeType pointing to the current value.
func (c ScopeType) ToPtr() *ScopeType {
	return &c
}

// SourceKindType - Source Kind to pull the configuration data from.
type SourceKindType string

const (
	SourceKindTypeBucket        SourceKindType = "Bucket"
	SourceKindTypeGitRepository SourceKindType = "GitRepository"
)

// PossibleSourceKindTypeValues returns the possible values for the SourceKindType const type.
func PossibleSourceKindTypeValues() []SourceKindType {
	return []SourceKindType{
		SourceKindTypeBucket,
		SourceKindTypeGitRepository,
	}
}

// ToPtr returns a *SourceKindType pointing to the current value.
func (c SourceKindType) ToPtr() *SourceKindType {
	return &c
}
