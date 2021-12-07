//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armblockchain

const (
	module  = "armblockchain"
	version = "v0.1.1"
)

// BlockchainMemberProvisioningState - Gets or sets the blockchain member provision state.
type BlockchainMemberProvisioningState string

const (
	BlockchainMemberProvisioningStateDeleting     BlockchainMemberProvisioningState = "Deleting"
	BlockchainMemberProvisioningStateFailed       BlockchainMemberProvisioningState = "Failed"
	BlockchainMemberProvisioningStateNotSpecified BlockchainMemberProvisioningState = "NotSpecified"
	BlockchainMemberProvisioningStateStale        BlockchainMemberProvisioningState = "Stale"
	BlockchainMemberProvisioningStateSucceeded    BlockchainMemberProvisioningState = "Succeeded"
	BlockchainMemberProvisioningStateUpdating     BlockchainMemberProvisioningState = "Updating"
)

// PossibleBlockchainMemberProvisioningStateValues returns the possible values for the BlockchainMemberProvisioningState const type.
func PossibleBlockchainMemberProvisioningStateValues() []BlockchainMemberProvisioningState {
	return []BlockchainMemberProvisioningState{
		BlockchainMemberProvisioningStateDeleting,
		BlockchainMemberProvisioningStateFailed,
		BlockchainMemberProvisioningStateNotSpecified,
		BlockchainMemberProvisioningStateStale,
		BlockchainMemberProvisioningStateSucceeded,
		BlockchainMemberProvisioningStateUpdating,
	}
}

// ToPtr returns a *BlockchainMemberProvisioningState pointing to the current value.
func (c BlockchainMemberProvisioningState) ToPtr() *BlockchainMemberProvisioningState {
	return &c
}

// BlockchainProtocol - Gets or sets the blockchain protocol.
type BlockchainProtocol string

const (
	BlockchainProtocolCorda        BlockchainProtocol = "Corda"
	BlockchainProtocolNotSpecified BlockchainProtocol = "NotSpecified"
	BlockchainProtocolParity       BlockchainProtocol = "Parity"
	BlockchainProtocolQuorum       BlockchainProtocol = "Quorum"
)

// PossibleBlockchainProtocolValues returns the possible values for the BlockchainProtocol const type.
func PossibleBlockchainProtocolValues() []BlockchainProtocol {
	return []BlockchainProtocol{
		BlockchainProtocolCorda,
		BlockchainProtocolNotSpecified,
		BlockchainProtocolParity,
		BlockchainProtocolQuorum,
	}
}

// ToPtr returns a *BlockchainProtocol pointing to the current value.
func (c BlockchainProtocol) ToPtr() *BlockchainProtocol {
	return &c
}

// NameAvailabilityReason - Gets or sets the name availability reason.
type NameAvailabilityReason string

const (
	NameAvailabilityReasonAlreadyExists NameAvailabilityReason = "AlreadyExists"
	NameAvailabilityReasonInvalid       NameAvailabilityReason = "Invalid"
	NameAvailabilityReasonNotSpecified  NameAvailabilityReason = "NotSpecified"
)

// PossibleNameAvailabilityReasonValues returns the possible values for the NameAvailabilityReason const type.
func PossibleNameAvailabilityReasonValues() []NameAvailabilityReason {
	return []NameAvailabilityReason{
		NameAvailabilityReasonAlreadyExists,
		NameAvailabilityReasonInvalid,
		NameAvailabilityReasonNotSpecified,
	}
}

// ToPtr returns a *NameAvailabilityReason pointing to the current value.
func (c NameAvailabilityReason) ToPtr() *NameAvailabilityReason {
	return &c
}

// NodeProvisioningState - Gets or sets the blockchain member provision state.
type NodeProvisioningState string

const (
	NodeProvisioningStateDeleting     NodeProvisioningState = "Deleting"
	NodeProvisioningStateFailed       NodeProvisioningState = "Failed"
	NodeProvisioningStateNotSpecified NodeProvisioningState = "NotSpecified"
	NodeProvisioningStateSucceeded    NodeProvisioningState = "Succeeded"
	NodeProvisioningStateUpdating     NodeProvisioningState = "Updating"
)

// PossibleNodeProvisioningStateValues returns the possible values for the NodeProvisioningState const type.
func PossibleNodeProvisioningStateValues() []NodeProvisioningState {
	return []NodeProvisioningState{
		NodeProvisioningStateDeleting,
		NodeProvisioningStateFailed,
		NodeProvisioningStateNotSpecified,
		NodeProvisioningStateSucceeded,
		NodeProvisioningStateUpdating,
	}
}

// ToPtr returns a *NodeProvisioningState pointing to the current value.
func (c NodeProvisioningState) ToPtr() *NodeProvisioningState {
	return &c
}
