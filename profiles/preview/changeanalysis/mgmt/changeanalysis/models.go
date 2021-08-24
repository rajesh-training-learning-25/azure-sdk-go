//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package changeanalysis

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/changeanalysis/mgmt/2021-04-01/changeanalysis"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ChangeCategory = original.ChangeCategory

const (
	System ChangeCategory = original.System
	User   ChangeCategory = original.User
)

type ChangeType = original.ChangeType

const (
	Add    ChangeType = original.Add
	Remove ChangeType = original.Remove
	Update ChangeType = original.Update
)

type Level = original.Level

const (
	Important Level = original.Important
	Noisy     Level = original.Noisy
	Normal    Level = original.Normal
)

type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type Change = original.Change
type ChangeList = original.ChangeList
type ChangeListIterator = original.ChangeListIterator
type ChangeListPage = original.ChangeListPage
type ChangeProperties = original.ChangeProperties
type ChangesClient = original.ChangesClient
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorDetail = original.ErrorDetail
type ErrorResponse = original.ErrorResponse
type OperationsClient = original.OperationsClient
type PropertyChange = original.PropertyChange
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type ResourceChangesClient = original.ResourceChangesClient
type ResourceProviderOperationDefinition = original.ResourceProviderOperationDefinition
type ResourceProviderOperationDisplay = original.ResourceProviderOperationDisplay
type ResourceProviderOperationList = original.ResourceProviderOperationList
type ResourceProviderOperationListIterator = original.ResourceProviderOperationListIterator
type ResourceProviderOperationListPage = original.ResourceProviderOperationListPage
type TrackedResource = original.TrackedResource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewChangeListIterator(page ChangeListPage) ChangeListIterator {
	return original.NewChangeListIterator(page)
}
func NewChangeListPage(cur ChangeList, getNextPage func(context.Context, ChangeList) (ChangeList, error)) ChangeListPage {
	return original.NewChangeListPage(cur, getNextPage)
}
func NewChangesClient(subscriptionID string) ChangesClient {
	return original.NewChangesClient(subscriptionID)
}
func NewChangesClientWithBaseURI(baseURI string, subscriptionID string) ChangesClient {
	return original.NewChangesClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewResourceChangesClient(subscriptionID string) ResourceChangesClient {
	return original.NewResourceChangesClient(subscriptionID)
}
func NewResourceChangesClientWithBaseURI(baseURI string, subscriptionID string) ResourceChangesClient {
	return original.NewResourceChangesClientWithBaseURI(baseURI, subscriptionID)
}
func NewResourceProviderOperationListIterator(page ResourceProviderOperationListPage) ResourceProviderOperationListIterator {
	return original.NewResourceProviderOperationListIterator(page)
}
func NewResourceProviderOperationListPage(cur ResourceProviderOperationList, getNextPage func(context.Context, ResourceProviderOperationList) (ResourceProviderOperationList, error)) ResourceProviderOperationListPage {
	return original.NewResourceProviderOperationListPage(cur, getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleChangeCategoryValues() []ChangeCategory {
	return original.PossibleChangeCategoryValues()
}
func PossibleChangeTypeValues() []ChangeType {
	return original.PossibleChangeTypeValues()
}
func PossibleLevelValues() []Level {
	return original.PossibleLevelValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
