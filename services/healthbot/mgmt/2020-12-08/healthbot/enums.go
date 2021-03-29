package healthbot

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// IdentityType enumerates the values for identity type.
type IdentityType string

const (
	// Application ...
	Application IdentityType = "Application"
	// Key ...
	Key IdentityType = "Key"
	// ManagedIdentity ...
	ManagedIdentity IdentityType = "ManagedIdentity"
	// User ...
	User IdentityType = "User"
)

// PossibleIdentityTypeValues returns an array of possible values for the IdentityType const type.
func PossibleIdentityTypeValues() []IdentityType {
	return []IdentityType{Application, Key, ManagedIdentity, User}
}

// SkuName enumerates the values for sku name.
type SkuName string

const (
	// C0 ...
	C0 SkuName = "C0"
	// F0 ...
	F0 SkuName = "F0"
	// S1 ...
	S1 SkuName = "S1"
)

// PossibleSkuNameValues returns an array of possible values for the SkuName const type.
func PossibleSkuNameValues() []SkuName {
	return []SkuName{C0, F0, S1}
}
