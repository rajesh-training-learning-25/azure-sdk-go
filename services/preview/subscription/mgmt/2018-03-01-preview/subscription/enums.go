package subscription

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// OfferType enumerates the values for offer type.
type OfferType string

const (
	// MSAZR0017P ...
	MSAZR0017P OfferType = "MS-AZR-0017P"
	// MSAZR0148P ...
	MSAZR0148P OfferType = "MS-AZR-0148P"
)

// PossibleOfferTypeValues returns an array of possible values for the OfferType const type.
func PossibleOfferTypeValues() []OfferType {
	return []OfferType{MSAZR0017P, MSAZR0148P}
}

// SpendingLimit enumerates the values for spending limit.
type SpendingLimit string

const (
	// CurrentPeriodOff ...
	CurrentPeriodOff SpendingLimit = "CurrentPeriodOff"
	// Off ...
	Off SpendingLimit = "Off"
	// On ...
	On SpendingLimit = "On"
)

// PossibleSpendingLimitValues returns an array of possible values for the SpendingLimit const type.
func PossibleSpendingLimitValues() []SpendingLimit {
	return []SpendingLimit{CurrentPeriodOff, Off, On}
}

// State enumerates the values for state.
type State string

const (
	// Deleted ...
	Deleted State = "Deleted"
	// Disabled ...
	Disabled State = "Disabled"
	// Enabled ...
	Enabled State = "Enabled"
	// PastDue ...
	PastDue State = "PastDue"
	// Warned ...
	Warned State = "Warned"
)

// PossibleStateValues returns an array of possible values for the State const type.
func PossibleStateValues() []State {
	return []State{Deleted, Disabled, Enabled, PastDue, Warned}
}
