//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecommendationsservice

const (
	moduleName    = "armrecommendationsservice"
	moduleVersion = "v0.1.0"
)

// AccountConfiguration - Account configuration. This can only be set at RecommendationsService Account creation.
type AccountConfiguration string

const (
	AccountConfigurationCapacity AccountConfiguration = "Capacity"
	AccountConfigurationFree     AccountConfiguration = "Free"
)

// PossibleAccountConfigurationValues returns the possible values for the AccountConfiguration const type.
func PossibleAccountConfigurationValues() []AccountConfiguration {
	return []AccountConfiguration{
		AccountConfigurationCapacity,
		AccountConfigurationFree,
	}
}

// ToPtr returns a *AccountConfiguration pointing to the current value.
func (c AccountConfiguration) ToPtr() *AccountConfiguration {
	return &c
}

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

// CheckNameAvailabilityReason - The reason why the given name is not available.
type CheckNameAvailabilityReason string

const (
	CheckNameAvailabilityReasonAlreadyExists CheckNameAvailabilityReason = "AlreadyExists"
	CheckNameAvailabilityReasonInvalid       CheckNameAvailabilityReason = "Invalid"
)

// PossibleCheckNameAvailabilityReasonValues returns the possible values for the CheckNameAvailabilityReason const type.
func PossibleCheckNameAvailabilityReasonValues() []CheckNameAvailabilityReason {
	return []CheckNameAvailabilityReason{
		CheckNameAvailabilityReasonAlreadyExists,
		CheckNameAvailabilityReasonInvalid,
	}
}

// ToPtr returns a *CheckNameAvailabilityReason pointing to the current value.
func (c CheckNameAvailabilityReason) ToPtr() *CheckNameAvailabilityReason {
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

// ModelingFeatures - Modeling features controls the set of supported scenarios\models being computed. This can only be set
// at Modeling creation.
type ModelingFeatures string

const (
	ModelingFeaturesBasic    ModelingFeatures = "Basic"
	ModelingFeaturesPremium  ModelingFeatures = "Premium"
	ModelingFeaturesStandard ModelingFeatures = "Standard"
)

// PossibleModelingFeaturesValues returns the possible values for the ModelingFeatures const type.
func PossibleModelingFeaturesValues() []ModelingFeatures {
	return []ModelingFeatures{
		ModelingFeaturesBasic,
		ModelingFeaturesPremium,
		ModelingFeaturesStandard,
	}
}

// ToPtr returns a *ModelingFeatures pointing to the current value.
func (c ModelingFeatures) ToPtr() *ModelingFeatures {
	return &c
}

// ModelingFrequency - Modeling frequency controls the modeling compute frequency.
type ModelingFrequency string

const (
	ModelingFrequencyHigh   ModelingFrequency = "High"
	ModelingFrequencyLow    ModelingFrequency = "Low"
	ModelingFrequencyMedium ModelingFrequency = "Medium"
)

// PossibleModelingFrequencyValues returns the possible values for the ModelingFrequency const type.
func PossibleModelingFrequencyValues() []ModelingFrequency {
	return []ModelingFrequency{
		ModelingFrequencyHigh,
		ModelingFrequencyLow,
		ModelingFrequencyMedium,
	}
}

// ToPtr returns a *ModelingFrequency pointing to the current value.
func (c ModelingFrequency) ToPtr() *ModelingFrequency {
	return &c
}

// ModelingSize - Modeling size controls the maximum supported input data size.
type ModelingSize string

const (
	ModelingSizeLarge  ModelingSize = "Large"
	ModelingSizeMedium ModelingSize = "Medium"
	ModelingSizeSmall  ModelingSize = "Small"
)

// PossibleModelingSizeValues returns the possible values for the ModelingSize const type.
func PossibleModelingSizeValues() []ModelingSize {
	return []ModelingSize{
		ModelingSizeLarge,
		ModelingSizeMedium,
		ModelingSizeSmall,
	}
}

// ToPtr returns a *ModelingSize pointing to the current value.
func (c ModelingSize) ToPtr() *ModelingSize {
	return &c
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

// ToPtr returns a *Origin pointing to the current value.
func (c Origin) ToPtr() *Origin {
	return &c
}

// PrincipalType - AAD principal type.
type PrincipalType string

const (
	PrincipalTypeApplication PrincipalType = "Application"
	PrincipalTypeUser        PrincipalType = "User"
)

// PossiblePrincipalTypeValues returns the possible values for the PrincipalType const type.
func PossiblePrincipalTypeValues() []PrincipalType {
	return []PrincipalType{
		PrincipalTypeApplication,
		PrincipalTypeUser,
	}
}

// ToPtr returns a *PrincipalType pointing to the current value.
func (c PrincipalType) ToPtr() *PrincipalType {
	return &c
}
