//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armazureadexternalidentities

const (
	moduleName    = "armazureadexternalidentities"
	moduleVersion = "v0.1.0"
)

// B2CResourceSKUName - The name of the SKU for the tenant.
type B2CResourceSKUName string

const (
	// B2CResourceSKUNamePremiumP1 - Azure AD B2C usage is billed to a linked Azure subscription and uses number of authentications
	// based billing.
	B2CResourceSKUNamePremiumP1 B2CResourceSKUName = "PremiumP1"
	// B2CResourceSKUNamePremiumP2 - Azure AD B2C usage is billed to a linked Azure subscription and uses number of authentications
	// based billing.
	B2CResourceSKUNamePremiumP2 B2CResourceSKUName = "PremiumP2"
	// B2CResourceSKUNameStandard - Azure AD B2C usage is billed to a linked Azure subscription and uses a monthly active users
	// (MAU) billing model.
	B2CResourceSKUNameStandard B2CResourceSKUName = "Standard"
)

// PossibleB2CResourceSKUNameValues returns the possible values for the B2CResourceSKUName const type.
func PossibleB2CResourceSKUNameValues() []B2CResourceSKUName {
	return []B2CResourceSKUName{
		B2CResourceSKUNamePremiumP1,
		B2CResourceSKUNamePremiumP2,
		B2CResourceSKUNameStandard,
	}
}

// B2CResourceSKUTier - The tier of the tenant.
type B2CResourceSKUTier string

const (
	// B2CResourceSKUTierA0 - The SKU tier used for all Azure AD B2C tenants.
	B2CResourceSKUTierA0 B2CResourceSKUTier = "A0"
)

// PossibleB2CResourceSKUTierValues returns the possible values for the B2CResourceSKUTier const type.
func PossibleB2CResourceSKUTierValues() []B2CResourceSKUTier {
	return []B2CResourceSKUTier{
		B2CResourceSKUTierA0,
	}
}

// BillingType - The type of billing. Will be MAU for all new customers. If 'Auths', it can be updated to 'MAU'. Cannot be
// changed if value is 'MAU'. Learn more about Azure AD B2C billing at aka.ms/b2cBilling
// [https://aka.ms/b2cbilling].
type BillingType string

const (
	// BillingTypeAuths - Azure AD B2C usage is billed to a linked Azure subscription and uses number of authentications based
	// billing.
	BillingTypeAuths BillingType = "Auths"
	// BillingTypeMAU - Azure AD B2C usage is billed to a linked Azure subscription and uses a monthly active users (MAU) billing
	// model.
	BillingTypeMAU BillingType = "MAU"
)

// PossibleBillingTypeValues returns the possible values for the BillingType const type.
func PossibleBillingTypeValues() []BillingType {
	return []BillingType{
		BillingTypeAuths,
		BillingTypeMAU,
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

// NameAvailabilityReasonType - Describes the reason for the 'nameAvailable' value.
type NameAvailabilityReasonType string

const (
	// NameAvailabilityReasonTypeAlreadyExists - The name is already in use and is therefore unavailable.
	NameAvailabilityReasonTypeAlreadyExists NameAvailabilityReasonType = "AlreadyExists"
	// NameAvailabilityReasonTypeInvalid - The name provided does not match the resource provider’s naming requirements (incorrect
	// length, unsupported characters, etc.).
	NameAvailabilityReasonTypeInvalid NameAvailabilityReasonType = "Invalid"
)

// PossibleNameAvailabilityReasonTypeValues returns the possible values for the NameAvailabilityReasonType const type.
func PossibleNameAvailabilityReasonTypeValues() []NameAvailabilityReasonType {
	return []NameAvailabilityReasonType{
		NameAvailabilityReasonTypeAlreadyExists,
		NameAvailabilityReasonTypeInvalid,
	}
}

// TypeValue - The type of the B2C tenant resource.
type TypeValue string

const (
	// TypeValueMicrosoftAzureActiveDirectoryB2CDirectories - The resource type for Azure AD B2C tenant resource.
	TypeValueMicrosoftAzureActiveDirectoryB2CDirectories TypeValue = "Microsoft.AzureActiveDirectory/b2cDirectories"
)

// PossibleTypeValueValues returns the possible values for the TypeValue const type.
func PossibleTypeValueValues() []TypeValue {
	return []TypeValue{
		TypeValueMicrosoftAzureActiveDirectoryB2CDirectories,
	}
}
