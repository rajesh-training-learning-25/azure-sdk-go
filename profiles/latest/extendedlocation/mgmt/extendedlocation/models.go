//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package extendedlocation

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/extendedlocation/mgmt/2021-08-15/extendedlocation"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type CreatedByType = original.CreatedByType

const (
	CreatedByTypeApplication     CreatedByType = original.CreatedByTypeApplication
	CreatedByTypeKey             CreatedByType = original.CreatedByTypeKey
	CreatedByTypeManagedIdentity CreatedByType = original.CreatedByTypeManagedIdentity
	CreatedByTypeUser            CreatedByType = original.CreatedByTypeUser
)

type HostType = original.HostType

const (
	HostTypeKubernetes HostType = original.HostTypeKubernetes
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	ResourceIdentityTypeNone           ResourceIdentityType = original.ResourceIdentityTypeNone
	ResourceIdentityTypeSystemAssigned ResourceIdentityType = original.ResourceIdentityTypeSystemAssigned
)

type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type CustomLocation = original.CustomLocation
type CustomLocationListResult = original.CustomLocationListResult
type CustomLocationListResultIterator = original.CustomLocationListResultIterator
type CustomLocationListResultPage = original.CustomLocationListResultPage
type CustomLocationOperation = original.CustomLocationOperation
type CustomLocationOperationValueDisplay = original.CustomLocationOperationValueDisplay
type CustomLocationOperationsList = original.CustomLocationOperationsList
type CustomLocationOperationsListIterator = original.CustomLocationOperationsListIterator
type CustomLocationOperationsListPage = original.CustomLocationOperationsListPage
type CustomLocationProperties = original.CustomLocationProperties
type CustomLocationPropertiesAuthentication = original.CustomLocationPropertiesAuthentication
type CustomLocationsClient = original.CustomLocationsClient
type CustomLocationsCreateOrUpdateFuture = original.CustomLocationsCreateOrUpdateFuture
type CustomLocationsDeleteFuture = original.CustomLocationsDeleteFuture
type EnabledResourceType = original.EnabledResourceType
type EnabledResourceTypeProperties = original.EnabledResourceTypeProperties
type EnabledResourceTypePropertiesTypesMetadataItem = original.EnabledResourceTypePropertiesTypesMetadataItem
type EnabledResourceTypesListResult = original.EnabledResourceTypesListResult
type EnabledResourceTypesListResultIterator = original.EnabledResourceTypesListResultIterator
type EnabledResourceTypesListResultPage = original.EnabledResourceTypesListResultPage
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorDetail = original.ErrorDetail
type ErrorResponse = original.ErrorResponse
type Identity = original.Identity
type PatchableCustomLocations = original.PatchableCustomLocations
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type SystemData = original.SystemData
type TrackedResource = original.TrackedResource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewCustomLocationListResultIterator(page CustomLocationListResultPage) CustomLocationListResultIterator {
	return original.NewCustomLocationListResultIterator(page)
}
func NewCustomLocationListResultPage(cur CustomLocationListResult, getNextPage func(context.Context, CustomLocationListResult) (CustomLocationListResult, error)) CustomLocationListResultPage {
	return original.NewCustomLocationListResultPage(cur, getNextPage)
}
func NewCustomLocationOperationsListIterator(page CustomLocationOperationsListPage) CustomLocationOperationsListIterator {
	return original.NewCustomLocationOperationsListIterator(page)
}
func NewCustomLocationOperationsListPage(cur CustomLocationOperationsList, getNextPage func(context.Context, CustomLocationOperationsList) (CustomLocationOperationsList, error)) CustomLocationOperationsListPage {
	return original.NewCustomLocationOperationsListPage(cur, getNextPage)
}
func NewCustomLocationsClient(subscriptionID string) CustomLocationsClient {
	return original.NewCustomLocationsClient(subscriptionID)
}
func NewCustomLocationsClientWithBaseURI(baseURI string, subscriptionID string) CustomLocationsClient {
	return original.NewCustomLocationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewEnabledResourceTypesListResultIterator(page EnabledResourceTypesListResultPage) EnabledResourceTypesListResultIterator {
	return original.NewEnabledResourceTypesListResultIterator(page)
}
func NewEnabledResourceTypesListResultPage(cur EnabledResourceTypesListResult, getNextPage func(context.Context, EnabledResourceTypesListResult) (EnabledResourceTypesListResult, error)) EnabledResourceTypesListResultPage {
	return original.NewEnabledResourceTypesListResultPage(cur, getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleHostTypeValues() []HostType {
	return original.PossibleHostTypeValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
