// +build go1.9

// Copyright 2018 Microsoft Corporation
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

package managementgroups

import original "github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2018-03-01-preview/management"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type EntitiesClient = original.EntitiesClient
type Client = original.Client
type InheritedPermissions = original.InheritedPermissions

const (
	Delete   InheritedPermissions = original.Delete
	Edit     InheritedPermissions = original.Edit
	Noaccess InheritedPermissions = original.Noaccess
	View     InheritedPermissions = original.View
)

type Permissions = original.Permissions

const (
	PermissionsDelete   Permissions = original.PermissionsDelete
	PermissionsEdit     Permissions = original.PermissionsEdit
	PermissionsNoaccess Permissions = original.PermissionsNoaccess
	PermissionsView     Permissions = original.PermissionsView
)

type Permissions1 = original.Permissions1

const (
	Permissions1Delete   Permissions1 = original.Permissions1Delete
	Permissions1Edit     Permissions1 = original.Permissions1Edit
	Permissions1Noaccess Permissions1 = original.Permissions1Noaccess
	Permissions1View     Permissions1 = original.Permissions1View
)

type ProvisioningState = original.ProvisioningState

const (
	Updating ProvisioningState = original.Updating
)

type Reason = original.Reason

const (
	AlreadyExists Reason = original.AlreadyExists
	Invalid       Reason = original.Invalid
)

type Status = original.Status

const (
	Cancelled                Status = original.Cancelled
	Completed                Status = original.Completed
	Failed                   Status = original.Failed
	NotStarted               Status = original.NotStarted
	NotStartedButGroupsExist Status = original.NotStartedButGroupsExist
	Started                  Status = original.Started
)

type Type = original.Type

const (
	ProvidersMicrosoftManagementmanagementGroup Type = original.ProvidersMicrosoftManagementmanagementGroup
)

type Type1 = original.Type1

const (
	ProvidersMicrosoftManagementmanagementGroups Type1 = original.ProvidersMicrosoftManagementmanagementGroups
	Subscriptions                                Type1 = original.Subscriptions
)

type Type2 = original.Type2

const (
	Type2ProvidersMicrosoftManagementmanagementGroups Type2 = original.Type2ProvidersMicrosoftManagementmanagementGroups
	Type2Subscriptions                                Type2 = original.Type2Subscriptions
)

type CheckNameAvailabilityRequest = original.CheckNameAvailabilityRequest
type CheckNameAvailabilityResult = original.CheckNameAvailabilityResult
type ChildInfo = original.ChildInfo
type CreateManagementGroupChildInfo = original.CreateManagementGroupChildInfo
type CreateManagementGroupDetails = original.CreateManagementGroupDetails
type CreateManagementGroupProperties = original.CreateManagementGroupProperties
type CreateManagementGroupRequest = original.CreateManagementGroupRequest
type CreateOrUpdateFuture = original.CreateOrUpdateFuture
type CreateParentGroupInfo = original.CreateParentGroupInfo
type DeleteFuture = original.DeleteFuture
type Details = original.Details
type EntitiesListAllFuture = original.EntitiesListAllFuture
type EntitiesListFuture = original.EntitiesListFuture
type EntityHierarchyItem = original.EntityHierarchyItem
type EntityHierarchyItemProperties = original.EntityHierarchyItemProperties
type EntityInfo = original.EntityInfo
type EntityInfoProperties = original.EntityInfoProperties
type EntityListResult = original.EntityListResult
type EntityListResultIterator = original.EntityListResultIterator
type EntityListResultPage = original.EntityListResultPage
type EntityParentGroupInfo = original.EntityParentGroupInfo
type ErrorDetails = original.ErrorDetails
type ErrorResponse = original.ErrorResponse
type Info = original.Info
type InfoProperties = original.InfoProperties
type ListResult = original.ListResult
type ListResultIterator = original.ListResultIterator
type ListResultPage = original.ListResultPage
type ManagementGroup = original.ManagementGroup
type Operation = original.Operation
type OperationDisplayProperties = original.OperationDisplayProperties
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationResults = original.OperationResults
type OperationResultsProperties = original.OperationResultsProperties
type ParentGroupInfo = original.ParentGroupInfo
type PatchManagementGroupRequest = original.PatchManagementGroupRequest
type Properties = original.Properties
type SetObject = original.SetObject
type StartTenantBackfillFuture = original.StartTenantBackfillFuture
type TenantBackfillStatusResult = original.TenantBackfillStatusResult
type OperationsClient = original.OperationsClient
type SubscriptionsClient = original.SubscriptionsClient

func New(operationResultID string, skip *int32, top *int32, skiptoken string) BaseClient {
	return original.New(operationResultID, skip, top, skiptoken)
}
func NewWithBaseURI(baseURI string, operationResultID string, skip *int32, top *int32, skiptoken string) BaseClient {
	return original.NewWithBaseURI(baseURI, operationResultID, skip, top, skiptoken)
}
func NewEntitiesClient(operationResultID string, skip *int32, top *int32, skiptoken string) EntitiesClient {
	return original.NewEntitiesClient(operationResultID, skip, top, skiptoken)
}
func NewEntitiesClientWithBaseURI(baseURI string, operationResultID string, skip *int32, top *int32, skiptoken string) EntitiesClient {
	return original.NewEntitiesClientWithBaseURI(baseURI, operationResultID, skip, top, skiptoken)
}
func NewClient(operationResultID string, skip *int32, top *int32, skiptoken string) Client {
	return original.NewClient(operationResultID, skip, top, skiptoken)
}
func NewClientWithBaseURI(baseURI string, operationResultID string, skip *int32, top *int32, skiptoken string) Client {
	return original.NewClientWithBaseURI(baseURI, operationResultID, skip, top, skiptoken)
}
func PossibleInheritedPermissionsValues() []InheritedPermissions {
	return original.PossibleInheritedPermissionsValues()
}
func PossiblePermissionsValues() []Permissions {
	return original.PossiblePermissionsValues()
}
func PossiblePermissions1Values() []Permissions1 {
	return original.PossiblePermissions1Values()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleReasonValues() []Reason {
	return original.PossibleReasonValues()
}
func PossibleStatusValues() []Status {
	return original.PossibleStatusValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func PossibleType1Values() []Type1 {
	return original.PossibleType1Values()
}
func PossibleType2Values() []Type2 {
	return original.PossibleType2Values()
}
func NewOperationsClient(operationResultID string, skip *int32, top *int32, skiptoken string) OperationsClient {
	return original.NewOperationsClient(operationResultID, skip, top, skiptoken)
}
func NewOperationsClientWithBaseURI(baseURI string, operationResultID string, skip *int32, top *int32, skiptoken string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, operationResultID, skip, top, skiptoken)
}
func NewSubscriptionsClient(operationResultID string, skip *int32, top *int32, skiptoken string) SubscriptionsClient {
	return original.NewSubscriptionsClient(operationResultID, skip, top, skiptoken)
}
func NewSubscriptionsClientWithBaseURI(baseURI string, operationResultID string, skip *int32, top *int32, skiptoken string) SubscriptionsClient {
	return original.NewSubscriptionsClientWithBaseURI(baseURI, operationResultID, skip, top, skiptoken)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
