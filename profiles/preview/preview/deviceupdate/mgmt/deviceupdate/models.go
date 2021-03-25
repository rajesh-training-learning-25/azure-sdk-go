// +build go1.9

// Copyright 2021 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package deviceupdate

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/deviceupdate/mgmt/2020-03-01-preview/deviceupdate"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ActionType = original.ActionType

const (
	Internal ActionType = original.Internal
)

type Origin = original.Origin

const (
	System     Origin = original.System
	User       Origin = original.User
	Usersystem Origin = original.Usersystem
)

type ProvisioningState = original.ProvisioningState

const (
	Accepted  ProvisioningState = original.Accepted
	Canceled  ProvisioningState = original.Canceled
	Creating  ProvisioningState = original.Creating
	Deleted   ProvisioningState = original.Deleted
	Failed    ProvisioningState = original.Failed
	Succeeded ProvisioningState = original.Succeeded
)

type Account = original.Account
type AccountList = original.AccountList
type AccountListIterator = original.AccountListIterator
type AccountListPage = original.AccountListPage
type AccountProperties = original.AccountProperties
type AccountUpdate = original.AccountUpdate
type AccountsClient = original.AccountsClient
type AccountsCreateFuture = original.AccountsCreateFuture
type AccountsDeleteFuture = original.AccountsDeleteFuture
type AccountsUpdateFuture = original.AccountsUpdateFuture
type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorDefinition = original.ErrorDefinition
type ErrorResponse = original.ErrorResponse
type Instance = original.Instance
type InstanceList = original.InstanceList
type InstanceListIterator = original.InstanceListIterator
type InstanceListPage = original.InstanceListPage
type InstanceProperties = original.InstanceProperties
type InstancesClient = original.InstancesClient
type InstancesCreateFuture = original.InstancesCreateFuture
type InstancesDeleteFuture = original.InstancesDeleteFuture
type IotHubSettings = original.IotHubSettings
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type TagUpdate = original.TagUpdate
type TrackedResource = original.TrackedResource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAccountListIterator(page AccountListPage) AccountListIterator {
	return original.NewAccountListIterator(page)
}
func NewAccountListPage(cur AccountList, getNextPage func(context.Context, AccountList) (AccountList, error)) AccountListPage {
	return original.NewAccountListPage(cur, getNextPage)
}
func NewAccountsClient(subscriptionID string) AccountsClient {
	return original.NewAccountsClient(subscriptionID)
}
func NewAccountsClientWithBaseURI(baseURI string, subscriptionID string) AccountsClient {
	return original.NewAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewInstanceListIterator(page InstanceListPage) InstanceListIterator {
	return original.NewInstanceListIterator(page)
}
func NewInstanceListPage(cur InstanceList, getNextPage func(context.Context, InstanceList) (InstanceList, error)) InstanceListPage {
	return original.NewInstanceListPage(cur, getNextPage)
}
func NewInstancesClient(subscriptionID string) InstancesClient {
	return original.NewInstancesClient(subscriptionID)
}
func NewInstancesClientWithBaseURI(baseURI string, subscriptionID string) InstancesClient {
	return original.NewInstancesClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleActionTypeValues() []ActionType {
	return original.PossibleActionTypeValues()
}
func PossibleOriginValues() []Origin {
	return original.PossibleOriginValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
