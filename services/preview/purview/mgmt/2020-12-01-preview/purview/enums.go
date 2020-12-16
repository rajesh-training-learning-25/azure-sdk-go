package purview

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// DefaultAction enumerates the values for default action.
type DefaultAction string

const (
	// Allow ...
	Allow DefaultAction = "Allow"
	// Deny ...
	Deny DefaultAction = "Deny"
)

// PossibleDefaultActionValues returns an array of possible values for the DefaultAction const type.
func PossibleDefaultActionValues() []DefaultAction {
	return []DefaultAction{Allow, Deny}
}

// Name enumerates the values for name.
type Name string

const (
	// Standard ...
	Standard Name = "Standard"
)

// PossibleNameValues returns an array of possible values for the Name const type.
func PossibleNameValues() []Name {
	return []Name{Standard}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Creating ...
	Creating ProvisioningState = "Creating"
	// Deleting ...
	Deleting ProvisioningState = "Deleting"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Moving ...
	Moving ProvisioningState = "Moving"
	// SoftDeleted ...
	SoftDeleted ProvisioningState = "SoftDeleted"
	// SoftDeleting ...
	SoftDeleting ProvisioningState = "SoftDeleting"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
	// Unknown ...
	Unknown ProvisioningState = "Unknown"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Creating, Deleting, Failed, Moving, SoftDeleted, SoftDeleting, Succeeded, Unknown}
}

// Reason enumerates the values for reason.
type Reason string

const (
	// AlreadyExists ...
	AlreadyExists Reason = "AlreadyExists"
	// Invalid ...
	Invalid Reason = "Invalid"
)

// PossibleReasonValues returns an array of possible values for the Reason const type.
func PossibleReasonValues() []Reason {
	return []Reason{AlreadyExists, Invalid}
}

// ScopeType enumerates the values for scope type.
type ScopeType string

const (
	// Subscription ...
	Subscription ScopeType = "Subscription"
	// Tenant ...
	Tenant ScopeType = "Tenant"
)

// PossibleScopeTypeValues returns an array of possible values for the ScopeType const type.
func PossibleScopeTypeValues() []ScopeType {
	return []ScopeType{Subscription, Tenant}
}

// Status enumerates the values for status.
type Status string

const (
	// StatusApproved ...
	StatusApproved Status = "Approved"
	// StatusDisconnected ...
	StatusDisconnected Status = "Disconnected"
	// StatusPending ...
	StatusPending Status = "Pending"
	// StatusRejected ...
	StatusRejected Status = "Rejected"
	// StatusUnknown ...
	StatusUnknown Status = "Unknown"
)

// PossibleStatusValues returns an array of possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{StatusApproved, StatusDisconnected, StatusPending, StatusRejected, StatusUnknown}
}

// Type enumerates the values for type.
type Type string

const (
	// SystemAssigned ...
	SystemAssigned Type = "SystemAssigned"
)

// PossibleTypeValues returns an array of possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{SystemAssigned}
}
