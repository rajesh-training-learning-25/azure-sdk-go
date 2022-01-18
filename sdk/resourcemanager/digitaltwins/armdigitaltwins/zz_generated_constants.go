//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdigitaltwins

const (
	moduleName    = "armdigitaltwins"
	moduleVersion = "v0.2.0"
)

// AuthenticationType - Specifies the authentication type being used for connecting to the endpoint.
type AuthenticationType string

const (
	AuthenticationTypeIdentityBased AuthenticationType = "IdentityBased"
	AuthenticationTypeKeyBased      AuthenticationType = "KeyBased"
)

// PossibleAuthenticationTypeValues returns the possible values for the AuthenticationType const type.
func PossibleAuthenticationTypeValues() []AuthenticationType {
	return []AuthenticationType{
		AuthenticationTypeIdentityBased,
		AuthenticationTypeKeyBased,
	}
}

// ToPtr returns a *AuthenticationType pointing to the current value.
func (c AuthenticationType) ToPtr() *AuthenticationType {
	return &c
}

// ConnectionPropertiesProvisioningState - The provisioning state.
type ConnectionPropertiesProvisioningState string

const (
	ConnectionPropertiesProvisioningStateApproved     ConnectionPropertiesProvisioningState = "Approved"
	ConnectionPropertiesProvisioningStateDisconnected ConnectionPropertiesProvisioningState = "Disconnected"
	ConnectionPropertiesProvisioningStatePending      ConnectionPropertiesProvisioningState = "Pending"
	ConnectionPropertiesProvisioningStateRejected     ConnectionPropertiesProvisioningState = "Rejected"
)

// PossibleConnectionPropertiesProvisioningStateValues returns the possible values for the ConnectionPropertiesProvisioningState const type.
func PossibleConnectionPropertiesProvisioningStateValues() []ConnectionPropertiesProvisioningState {
	return []ConnectionPropertiesProvisioningState{
		ConnectionPropertiesProvisioningStateApproved,
		ConnectionPropertiesProvisioningStateDisconnected,
		ConnectionPropertiesProvisioningStatePending,
		ConnectionPropertiesProvisioningStateRejected,
	}
}

// ToPtr returns a *ConnectionPropertiesProvisioningState pointing to the current value.
func (c ConnectionPropertiesProvisioningState) ToPtr() *ConnectionPropertiesProvisioningState {
	return &c
}

// DigitalTwinsIdentityType - The type of Managed Identity used by the DigitalTwinsInstance. Only SystemAssigned is supported.
type DigitalTwinsIdentityType string

const (
	DigitalTwinsIdentityTypeNone           DigitalTwinsIdentityType = "None"
	DigitalTwinsIdentityTypeSystemAssigned DigitalTwinsIdentityType = "SystemAssigned"
)

// PossibleDigitalTwinsIdentityTypeValues returns the possible values for the DigitalTwinsIdentityType const type.
func PossibleDigitalTwinsIdentityTypeValues() []DigitalTwinsIdentityType {
	return []DigitalTwinsIdentityType{
		DigitalTwinsIdentityTypeNone,
		DigitalTwinsIdentityTypeSystemAssigned,
	}
}

// ToPtr returns a *DigitalTwinsIdentityType pointing to the current value.
func (c DigitalTwinsIdentityType) ToPtr() *DigitalTwinsIdentityType {
	return &c
}

// EndpointProvisioningState - The provisioning state.
type EndpointProvisioningState string

const (
	EndpointProvisioningStateCanceled     EndpointProvisioningState = "Canceled"
	EndpointProvisioningStateDeleted      EndpointProvisioningState = "Deleted"
	EndpointProvisioningStateDeleting     EndpointProvisioningState = "Deleting"
	EndpointProvisioningStateDisabled     EndpointProvisioningState = "Disabled"
	EndpointProvisioningStateFailed       EndpointProvisioningState = "Failed"
	EndpointProvisioningStateMoving       EndpointProvisioningState = "Moving"
	EndpointProvisioningStateProvisioning EndpointProvisioningState = "Provisioning"
	EndpointProvisioningStateRestoring    EndpointProvisioningState = "Restoring"
	EndpointProvisioningStateSucceeded    EndpointProvisioningState = "Succeeded"
	EndpointProvisioningStateSuspending   EndpointProvisioningState = "Suspending"
	EndpointProvisioningStateWarning      EndpointProvisioningState = "Warning"
)

// PossibleEndpointProvisioningStateValues returns the possible values for the EndpointProvisioningState const type.
func PossibleEndpointProvisioningStateValues() []EndpointProvisioningState {
	return []EndpointProvisioningState{
		EndpointProvisioningStateCanceled,
		EndpointProvisioningStateDeleted,
		EndpointProvisioningStateDeleting,
		EndpointProvisioningStateDisabled,
		EndpointProvisioningStateFailed,
		EndpointProvisioningStateMoving,
		EndpointProvisioningStateProvisioning,
		EndpointProvisioningStateRestoring,
		EndpointProvisioningStateSucceeded,
		EndpointProvisioningStateSuspending,
		EndpointProvisioningStateWarning,
	}
}

// ToPtr returns a *EndpointProvisioningState pointing to the current value.
func (c EndpointProvisioningState) ToPtr() *EndpointProvisioningState {
	return &c
}

// EndpointType - The type of Digital Twins endpoint
type EndpointType string

const (
	EndpointTypeEventGrid  EndpointType = "EventGrid"
	EndpointTypeEventHub   EndpointType = "EventHub"
	EndpointTypeServiceBus EndpointType = "ServiceBus"
)

// PossibleEndpointTypeValues returns the possible values for the EndpointType const type.
func PossibleEndpointTypeValues() []EndpointType {
	return []EndpointType{
		EndpointTypeEventGrid,
		EndpointTypeEventHub,
		EndpointTypeServiceBus,
	}
}

// ToPtr returns a *EndpointType pointing to the current value.
func (c EndpointType) ToPtr() *EndpointType {
	return &c
}

// PrivateLinkServiceConnectionStatus - The status of a private endpoint connection.
type PrivateLinkServiceConnectionStatus string

const (
	PrivateLinkServiceConnectionStatusApproved     PrivateLinkServiceConnectionStatus = "Approved"
	PrivateLinkServiceConnectionStatusDisconnected PrivateLinkServiceConnectionStatus = "Disconnected"
	PrivateLinkServiceConnectionStatusPending      PrivateLinkServiceConnectionStatus = "Pending"
	PrivateLinkServiceConnectionStatusRejected     PrivateLinkServiceConnectionStatus = "Rejected"
)

// PossiblePrivateLinkServiceConnectionStatusValues returns the possible values for the PrivateLinkServiceConnectionStatus const type.
func PossiblePrivateLinkServiceConnectionStatusValues() []PrivateLinkServiceConnectionStatus {
	return []PrivateLinkServiceConnectionStatus{
		PrivateLinkServiceConnectionStatusApproved,
		PrivateLinkServiceConnectionStatusDisconnected,
		PrivateLinkServiceConnectionStatusPending,
		PrivateLinkServiceConnectionStatusRejected,
	}
}

// ToPtr returns a *PrivateLinkServiceConnectionStatus pointing to the current value.
func (c PrivateLinkServiceConnectionStatus) ToPtr() *PrivateLinkServiceConnectionStatus {
	return &c
}

// ProvisioningState - The provisioning state.
type ProvisioningState string

const (
	ProvisioningStateCanceled     ProvisioningState = "Canceled"
	ProvisioningStateDeleted      ProvisioningState = "Deleted"
	ProvisioningStateDeleting     ProvisioningState = "Deleting"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateMoving       ProvisioningState = "Moving"
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	ProvisioningStateRestoring    ProvisioningState = "Restoring"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
	ProvisioningStateSuspending   ProvisioningState = "Suspending"
	ProvisioningStateUpdating     ProvisioningState = "Updating"
	ProvisioningStateWarning      ProvisioningState = "Warning"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCanceled,
		ProvisioningStateDeleted,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateMoving,
		ProvisioningStateProvisioning,
		ProvisioningStateRestoring,
		ProvisioningStateSucceeded,
		ProvisioningStateSuspending,
		ProvisioningStateUpdating,
		ProvisioningStateWarning,
	}
}

// ToPtr returns a *ProvisioningState pointing to the current value.
func (c ProvisioningState) ToPtr() *ProvisioningState {
	return &c
}

// PublicNetworkAccess - Public network access for the DigitalTwinsInstance.
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

// ToPtr returns a *PublicNetworkAccess pointing to the current value.
func (c PublicNetworkAccess) ToPtr() *PublicNetworkAccess {
	return &c
}

// Reason - Message providing the reason why the given name is invalid.
type Reason string

const (
	ReasonAlreadyExists Reason = "AlreadyExists"
	ReasonInvalid       Reason = "Invalid"
)

// PossibleReasonValues returns the possible values for the Reason const type.
func PossibleReasonValues() []Reason {
	return []Reason{
		ReasonAlreadyExists,
		ReasonInvalid,
	}
}

// ToPtr returns a *Reason pointing to the current value.
func (c Reason) ToPtr() *Reason {
	return &c
}
