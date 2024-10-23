//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdnsresolver

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver"
	moduleVersion = "v1.3.0-beta.1"
)

// ActionType - The type of action to take.
type ActionType string

const (
	ActionTypeAlert ActionType = "Alert"
	ActionTypeAllow ActionType = "Allow"
	ActionTypeBlock ActionType = "Block"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeAlert,
		ActionTypeAllow,
		ActionTypeBlock,
	}
}

// BlockResponseCode - The response code for block actions.
type BlockResponseCode string

const (
	BlockResponseCodeSERVFAIL BlockResponseCode = "SERVFAIL"
)

// PossibleBlockResponseCodeValues returns the possible values for the BlockResponseCode const type.
func PossibleBlockResponseCodeValues() []BlockResponseCode {
	return []BlockResponseCode{
		BlockResponseCodeSERVFAIL,
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

// DNSResolverState - The current status of the DNS resolver. This is a read-only property and any attempt to set this value
// will be ignored.
type DNSResolverState string

const (
	DNSResolverStateConnected    DNSResolverState = "Connected"
	DNSResolverStateDisconnected DNSResolverState = "Disconnected"
)

// PossibleDNSResolverStateValues returns the possible values for the DNSResolverState const type.
func PossibleDNSResolverStateValues() []DNSResolverState {
	return []DNSResolverState{
		DNSResolverStateConnected,
		DNSResolverStateDisconnected,
	}
}

// DNSSecurityRuleState - The state of DNS security rule.
type DNSSecurityRuleState string

const (
	DNSSecurityRuleStateDisabled DNSSecurityRuleState = "Disabled"
	DNSSecurityRuleStateEnabled  DNSSecurityRuleState = "Enabled"
)

// PossibleDNSSecurityRuleStateValues returns the possible values for the DNSSecurityRuleState const type.
func PossibleDNSSecurityRuleStateValues() []DNSSecurityRuleState {
	return []DNSSecurityRuleState{
		DNSSecurityRuleStateDisabled,
		DNSSecurityRuleStateEnabled,
	}
}

// ForwardingRuleState - The state of forwarding rule.
type ForwardingRuleState string

const (
	ForwardingRuleStateDisabled ForwardingRuleState = "Disabled"
	ForwardingRuleStateEnabled  ForwardingRuleState = "Enabled"
)

// PossibleForwardingRuleStateValues returns the possible values for the ForwardingRuleState const type.
func PossibleForwardingRuleStateValues() []ForwardingRuleState {
	return []ForwardingRuleState{
		ForwardingRuleStateDisabled,
		ForwardingRuleStateEnabled,
	}
}

// IPAllocationMethod - Private IP address allocation method.
type IPAllocationMethod string

const (
	IPAllocationMethodDynamic IPAllocationMethod = "Dynamic"
	IPAllocationMethodStatic  IPAllocationMethod = "Static"
)

// PossibleIPAllocationMethodValues returns the possible values for the IPAllocationMethod const type.
func PossibleIPAllocationMethodValues() []IPAllocationMethod {
	return []IPAllocationMethod{
		IPAllocationMethodDynamic,
		IPAllocationMethodStatic,
	}
}

// ProvisioningState - The current provisioning state of the resource.
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
