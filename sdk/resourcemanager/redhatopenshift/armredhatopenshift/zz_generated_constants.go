//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armredhatopenshift

const (
	module  = "armredhatopenshift"
	version = "v0.1.0"
)

// ProvisioningState - ProvisioningState represents a provisioning state.
type ProvisioningState string

const (
	ProvisioningStateAdminUpdating ProvisioningState = "AdminUpdating"
	ProvisioningStateCreating      ProvisioningState = "Creating"
	ProvisioningStateDeleting      ProvisioningState = "Deleting"
	ProvisioningStateFailed        ProvisioningState = "Failed"
	ProvisioningStateSucceeded     ProvisioningState = "Succeeded"
	ProvisioningStateUpdating      ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAdminUpdating,
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

// VMSize - VMSize represents a VM size.
type VMSize string

const (
	VMSizeStandardD2SV3 VMSize = "Standard_D2s_v3"
	VMSizeStandardD4SV3 VMSize = "Standard_D4s_v3"
	VMSizeStandardD8SV3 VMSize = "Standard_D8s_v3"
)

// PossibleVMSizeValues returns the possible values for the VMSize const type.
func PossibleVMSizeValues() []VMSize {
	return []VMSize{
		VMSizeStandardD2SV3,
		VMSizeStandardD4SV3,
		VMSizeStandardD8SV3,
	}
}

// ToPtr returns a *VMSize pointing to the current value.
func (c VMSize) ToPtr() *VMSize {
	return &c
}

// Visibility - Visibility represents visibility.
type Visibility string

const (
	VisibilityPrivate Visibility = "Private"
	VisibilityPublic  Visibility = "Public"
)

// PossibleVisibilityValues returns the possible values for the Visibility const type.
func PossibleVisibilityValues() []Visibility {
	return []Visibility{
		VisibilityPrivate,
		VisibilityPublic,
	}
}

// ToPtr returns a *Visibility pointing to the current value.
func (c Visibility) ToPtr() *Visibility {
	return &c
}
