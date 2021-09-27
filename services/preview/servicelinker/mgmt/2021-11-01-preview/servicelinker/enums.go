package servicelinker

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ActionType enumerates the values for action type.
type ActionType string

const (
	// Internal ...
	Internal ActionType = "Internal"
)

// PossibleActionTypeValues returns an array of possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{Internal}
}

// AuthType enumerates the values for auth type.
type AuthType string

const (
	// Secret ...
	Secret AuthType = "secret"
	// ServicePrincipalCertificate ...
	ServicePrincipalCertificate AuthType = "servicePrincipalCertificate"
	// ServicePrincipalSecret ...
	ServicePrincipalSecret AuthType = "servicePrincipalSecret"
	// SystemAssignedIdentity ...
	SystemAssignedIdentity AuthType = "systemAssignedIdentity"
	// UserAssignedIdentity ...
	UserAssignedIdentity AuthType = "userAssignedIdentity"
)

// PossibleAuthTypeValues returns an array of possible values for the AuthType const type.
func PossibleAuthTypeValues() []AuthType {
	return []AuthType{Secret, ServicePrincipalCertificate, ServicePrincipalSecret, SystemAssignedIdentity, UserAssignedIdentity}
}

// AuthTypeBasicAuthInfoBase enumerates the values for auth type basic auth info base.
type AuthTypeBasicAuthInfoBase string

const (
	// AuthTypeAuthInfoBase ...
	AuthTypeAuthInfoBase AuthTypeBasicAuthInfoBase = "AuthInfoBase"
	// AuthTypeSecret ...
	AuthTypeSecret AuthTypeBasicAuthInfoBase = "secret"
	// AuthTypeServicePrincipalCertificate ...
	AuthTypeServicePrincipalCertificate AuthTypeBasicAuthInfoBase = "servicePrincipalCertificate"
	// AuthTypeServicePrincipalSecret ...
	AuthTypeServicePrincipalSecret AuthTypeBasicAuthInfoBase = "servicePrincipalSecret"
	// AuthTypeSystemAssignedIdentity ...
	AuthTypeSystemAssignedIdentity AuthTypeBasicAuthInfoBase = "systemAssignedIdentity"
	// AuthTypeUserAssignedIdentity ...
	AuthTypeUserAssignedIdentity AuthTypeBasicAuthInfoBase = "userAssignedIdentity"
)

// PossibleAuthTypeBasicAuthInfoBaseValues returns an array of possible values for the AuthTypeBasicAuthInfoBase const type.
func PossibleAuthTypeBasicAuthInfoBaseValues() []AuthTypeBasicAuthInfoBase {
	return []AuthTypeBasicAuthInfoBase{AuthTypeAuthInfoBase, AuthTypeSecret, AuthTypeServicePrincipalCertificate, AuthTypeServicePrincipalSecret, AuthTypeSystemAssignedIdentity, AuthTypeUserAssignedIdentity}
}

// ClientType enumerates the values for client type.
type ClientType string

const (
	// Django ...
	Django ClientType = "django"
	// Dotnet ...
	Dotnet ClientType = "dotnet"
	// Go ...
	Go ClientType = "go"
	// Java ...
	Java ClientType = "java"
	// Nodejs ...
	Nodejs ClientType = "nodejs"
	// None ...
	None ClientType = "none"
	// Php ...
	Php ClientType = "php"
	// Python ...
	Python ClientType = "python"
	// Ruby ...
	Ruby ClientType = "ruby"
	// SpringBoot ...
	SpringBoot ClientType = "springBoot"
)

// PossibleClientTypeValues returns an array of possible values for the ClientType const type.
func PossibleClientTypeValues() []ClientType {
	return []ClientType{Django, Dotnet, Go, Java, Nodejs, None, Php, Python, Ruby, SpringBoot}
}

// CreatedByType enumerates the values for created by type.
type CreatedByType string

const (
	// Application ...
	Application CreatedByType = "Application"
	// Key ...
	Key CreatedByType = "Key"
	// ManagedIdentity ...
	ManagedIdentity CreatedByType = "ManagedIdentity"
	// User ...
	User CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns an array of possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{Application, Key, ManagedIdentity, User}
}

// LinkerStatus enumerates the values for linker status.
type LinkerStatus string

const (
	// Healthy ...
	Healthy LinkerStatus = "Healthy"
	// Nothealthy ...
	Nothealthy LinkerStatus = "Not healthy"
)

// PossibleLinkerStatusValues returns an array of possible values for the LinkerStatus const type.
func PossibleLinkerStatusValues() []LinkerStatus {
	return []LinkerStatus{Healthy, Nothealthy}
}

// Origin enumerates the values for origin.
type Origin string

const (
	// OriginSystem ...
	OriginSystem Origin = "system"
	// OriginUser ...
	OriginUser Origin = "user"
	// OriginUsersystem ...
	OriginUsersystem Origin = "user,system"
)

// PossibleOriginValues returns an array of possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{OriginSystem, OriginUser, OriginUsersystem}
}
