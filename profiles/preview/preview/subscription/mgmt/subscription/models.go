//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package subscription

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/subscription/mgmt/2019-10-01-preview/subscription"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type OfferType = original.OfferType

const (
	MSAZR0017P OfferType = original.MSAZR0017P
	MSAZR0148P OfferType = original.MSAZR0148P
)

type ProvisioningState = original.ProvisioningState

const (
	Accepted  ProvisioningState = original.Accepted
	Failed    ProvisioningState = original.Failed
	Succeeded ProvisioningState = original.Succeeded
)

type SpendingLimit = original.SpendingLimit

const (
	CurrentPeriodOff SpendingLimit = original.CurrentPeriodOff
	Off              SpendingLimit = original.Off
	On               SpendingLimit = original.On
)

type State = original.State

const (
	Deleted  State = original.Deleted
	Disabled State = original.Disabled
	Enabled  State = original.Enabled
	PastDue  State = original.PastDue
	Warned   State = original.Warned
)

type Workload = original.Workload

const (
	DevTest    Workload = original.DevTest
	Production Workload = original.Production
)

type AdPrincipal = original.AdPrincipal
type BaseClient = original.BaseClient
type CanceledSubscriptionID = original.CanceledSubscriptionID
type Client = original.Client
type CreateAliasFuture = original.CreateAliasFuture
type CreateCspSubscriptionFuture = original.CreateCspSubscriptionFuture
type CreateSubscriptionFuture = original.CreateSubscriptionFuture
type CreateSubscriptionInEnrollmentAccountFuture = original.CreateSubscriptionInEnrollmentAccountFuture
type CreationParameters = original.CreationParameters
type CreationResult = original.CreationResult
type EnabledSubscriptionID = original.EnabledSubscriptionID
type ErrorResponse = original.ErrorResponse
type ErrorResponseBody = original.ErrorResponseBody
type ListResult = original.ListResult
type ListResultIterator = original.ListResultIterator
type ListResultPage = original.ListResultPage
type Location = original.Location
type LocationListResult = original.LocationListResult
type Model = original.Model
type ModernCspSubscriptionCreationParameters = original.ModernCspSubscriptionCreationParameters
type ModernSubscriptionCreationParameters = original.ModernSubscriptionCreationParameters
type Name = original.Name
type Operation = original.Operation
type OperationClient = original.OperationClient
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationsClient = original.OperationsClient
type Policies = original.Policies
type PutAliasListResult = original.PutAliasListResult
type PutAliasRequest = original.PutAliasRequest
type PutAliasRequestProperties = original.PutAliasRequestProperties
type PutAliasResponse = original.PutAliasResponse
type PutAliasResponseProperties = original.PutAliasResponseProperties
type RenamedSubscriptionID = original.RenamedSubscriptionID
type SubscriptionsClient = original.SubscriptionsClient
type TenantIDDescription = original.TenantIDDescription
type TenantListResult = original.TenantListResult
type TenantListResultIterator = original.TenantListResultIterator
type TenantListResultPage = original.TenantListResultPage
type TenantsClient = original.TenantsClient

func New() BaseClient {
	return original.New()
}
func NewClient() Client {
	return original.NewClient()
}
func NewClientWithBaseURI(baseURI string) Client {
	return original.NewClientWithBaseURI(baseURI)
}
func NewListResultIterator(page ListResultPage) ListResultIterator {
	return original.NewListResultIterator(page)
}
func NewListResultPage(cur ListResult, getNextPage func(context.Context, ListResult) (ListResult, error)) ListResultPage {
	return original.NewListResultPage(cur, getNextPage)
}
func NewOperationClient() OperationClient {
	return original.NewOperationClient()
}
func NewOperationClientWithBaseURI(baseURI string) OperationClient {
	return original.NewOperationClientWithBaseURI(baseURI)
}
func NewOperationsClient() OperationsClient {
	return original.NewOperationsClient()
}
func NewOperationsClientWithBaseURI(baseURI string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI)
}
func NewSubscriptionsClient() SubscriptionsClient {
	return original.NewSubscriptionsClient()
}
func NewSubscriptionsClientWithBaseURI(baseURI string) SubscriptionsClient {
	return original.NewSubscriptionsClientWithBaseURI(baseURI)
}
func NewTenantListResultIterator(page TenantListResultPage) TenantListResultIterator {
	return original.NewTenantListResultIterator(page)
}
func NewTenantListResultPage(cur TenantListResult, getNextPage func(context.Context, TenantListResult) (TenantListResult, error)) TenantListResultPage {
	return original.NewTenantListResultPage(cur, getNextPage)
}
func NewTenantsClient() TenantsClient {
	return original.NewTenantsClient()
}
func NewTenantsClientWithBaseURI(baseURI string) TenantsClient {
	return original.NewTenantsClientWithBaseURI(baseURI)
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
func PossibleOfferTypeValues() []OfferType {
	return original.PossibleOfferTypeValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleSpendingLimitValues() []SpendingLimit {
	return original.PossibleSpendingLimitValues()
}
func PossibleStateValues() []State {
	return original.PossibleStateValues()
}
func PossibleWorkloadValues() []Workload {
	return original.PossibleWorkloadValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
