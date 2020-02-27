// +build go1.9

// Copyright 2020 Microsoft Corporation
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

package resources

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-05-01/resources"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type DeploymentMode = original.DeploymentMode

const (
	Complete    DeploymentMode = original.Complete
	Incremental DeploymentMode = original.Incremental
)

type OnErrorDeploymentType = original.OnErrorDeploymentType

const (
	LastSuccessful     OnErrorDeploymentType = original.LastSuccessful
	SpecificDeployment OnErrorDeploymentType = original.SpecificDeployment
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	None                       ResourceIdentityType = original.None
	SystemAssigned             ResourceIdentityType = original.SystemAssigned
	SystemAssignedUserAssigned ResourceIdentityType = original.SystemAssignedUserAssigned
	UserAssigned               ResourceIdentityType = original.UserAssigned
)

type AliasPathType = original.AliasPathType
type AliasType = original.AliasType
type BaseClient = original.BaseClient
type BasicDependency = original.BasicDependency
type Client = original.Client
type CloudError = original.CloudError
type CreateOrUpdateByIDFuture = original.CreateOrUpdateByIDFuture
type CreateOrUpdateFuture = original.CreateOrUpdateFuture
type DebugSetting = original.DebugSetting
type DeleteByIDFuture = original.DeleteByIDFuture
type DeleteFuture = original.DeleteFuture
type Dependency = original.Dependency
type Deployment = original.Deployment
type DeploymentExportResult = original.DeploymentExportResult
type DeploymentExtended = original.DeploymentExtended
type DeploymentExtendedFilter = original.DeploymentExtendedFilter
type DeploymentListResult = original.DeploymentListResult
type DeploymentListResultIterator = original.DeploymentListResultIterator
type DeploymentListResultPage = original.DeploymentListResultPage
type DeploymentOperation = original.DeploymentOperation
type DeploymentOperationProperties = original.DeploymentOperationProperties
type DeploymentOperationsClient = original.DeploymentOperationsClient
type DeploymentOperationsListResult = original.DeploymentOperationsListResult
type DeploymentOperationsListResultIterator = original.DeploymentOperationsListResultIterator
type DeploymentOperationsListResultPage = original.DeploymentOperationsListResultPage
type DeploymentProperties = original.DeploymentProperties
type DeploymentPropertiesExtended = original.DeploymentPropertiesExtended
type DeploymentValidateResult = original.DeploymentValidateResult
type DeploymentsClient = original.DeploymentsClient
type DeploymentsCreateOrUpdateAtManagementGroupScopeFuture = original.DeploymentsCreateOrUpdateAtManagementGroupScopeFuture
type DeploymentsCreateOrUpdateAtSubscriptionScopeFuture = original.DeploymentsCreateOrUpdateAtSubscriptionScopeFuture
type DeploymentsCreateOrUpdateFuture = original.DeploymentsCreateOrUpdateFuture
type DeploymentsDeleteAtManagementGroupScopeFuture = original.DeploymentsDeleteAtManagementGroupScopeFuture
type DeploymentsDeleteAtSubscriptionScopeFuture = original.DeploymentsDeleteAtSubscriptionScopeFuture
type DeploymentsDeleteFuture = original.DeploymentsDeleteFuture
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorResponse = original.ErrorResponse
type ExportTemplateRequest = original.ExportTemplateRequest
type GenericResource = original.GenericResource
type GenericResourceExpanded = original.GenericResourceExpanded
type GenericResourceFilter = original.GenericResourceFilter
type Group = original.Group
type GroupExportResult = original.GroupExportResult
type GroupFilter = original.GroupFilter
type GroupListResult = original.GroupListResult
type GroupListResultIterator = original.GroupListResultIterator
type GroupListResultPage = original.GroupListResultPage
type GroupPatchable = original.GroupPatchable
type GroupProperties = original.GroupProperties
type GroupsClient = original.GroupsClient
type GroupsDeleteFuture = original.GroupsDeleteFuture
type HTTPMessage = original.HTTPMessage
type Identity = original.Identity
type IdentityUserAssignedIdentitiesValue = original.IdentityUserAssignedIdentitiesValue
type ListResult = original.ListResult
type ListResultIterator = original.ListResultIterator
type ListResultPage = original.ListResultPage
type ManagementErrorWithDetails = original.ManagementErrorWithDetails
type MoveInfo = original.MoveInfo
type MoveResourcesFuture = original.MoveResourcesFuture
type OnErrorDeployment = original.OnErrorDeployment
type OnErrorDeploymentExtended = original.OnErrorDeploymentExtended
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type ParametersLink = original.ParametersLink
type Plan = original.Plan
type Provider = original.Provider
type ProviderListResult = original.ProviderListResult
type ProviderListResultIterator = original.ProviderListResultIterator
type ProviderListResultPage = original.ProviderListResultPage
type ProviderOperationDisplayProperties = original.ProviderOperationDisplayProperties
type ProviderResourceType = original.ProviderResourceType
type ProvidersClient = original.ProvidersClient
type Resource = original.Resource
type Sku = original.Sku
type SubResource = original.SubResource
type TagCount = original.TagCount
type TagDetails = original.TagDetails
type TagValue = original.TagValue
type TagsClient = original.TagsClient
type TagsListResult = original.TagsListResult
type TagsListResultIterator = original.TagsListResultIterator
type TagsListResultPage = original.TagsListResultPage
type TargetResource = original.TargetResource
type TemplateHashResult = original.TemplateHashResult
type TemplateLink = original.TemplateLink
type UpdateByIDFuture = original.UpdateByIDFuture
type UpdateFuture = original.UpdateFuture
type ValidateMoveResourcesFuture = original.ValidateMoveResourcesFuture

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewClient(subscriptionID string) Client {
	return original.NewClient(subscriptionID)
}
func NewClientWithBaseURI(baseURI string, subscriptionID string) Client {
	return original.NewClientWithBaseURI(baseURI, subscriptionID)
}
func NewDeploymentListResultIterator(page DeploymentListResultPage) DeploymentListResultIterator {
	return original.NewDeploymentListResultIterator(page)
}
func NewDeploymentListResultPage(getNextPage func(context.Context, DeploymentListResult) (DeploymentListResult, error)) DeploymentListResultPage {
	return original.NewDeploymentListResultPage(getNextPage)
}
func NewDeploymentOperationsClient(subscriptionID string) DeploymentOperationsClient {
	return original.NewDeploymentOperationsClient(subscriptionID)
}
func NewDeploymentOperationsClientWithBaseURI(baseURI string, subscriptionID string) DeploymentOperationsClient {
	return original.NewDeploymentOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDeploymentOperationsListResultIterator(page DeploymentOperationsListResultPage) DeploymentOperationsListResultIterator {
	return original.NewDeploymentOperationsListResultIterator(page)
}
func NewDeploymentOperationsListResultPage(getNextPage func(context.Context, DeploymentOperationsListResult) (DeploymentOperationsListResult, error)) DeploymentOperationsListResultPage {
	return original.NewDeploymentOperationsListResultPage(getNextPage)
}
func NewDeploymentsClient(subscriptionID string) DeploymentsClient {
	return original.NewDeploymentsClient(subscriptionID)
}
func NewDeploymentsClientWithBaseURI(baseURI string, subscriptionID string) DeploymentsClient {
	return original.NewDeploymentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewGroupListResultIterator(page GroupListResultPage) GroupListResultIterator {
	return original.NewGroupListResultIterator(page)
}
func NewGroupListResultPage(getNextPage func(context.Context, GroupListResult) (GroupListResult, error)) GroupListResultPage {
	return original.NewGroupListResultPage(getNextPage)
}
func NewGroupsClient(subscriptionID string) GroupsClient {
	return original.NewGroupsClient(subscriptionID)
}
func NewGroupsClientWithBaseURI(baseURI string, subscriptionID string) GroupsClient {
	return original.NewGroupsClientWithBaseURI(baseURI, subscriptionID)
}
func NewListResultIterator(page ListResultPage) ListResultIterator {
	return original.NewListResultIterator(page)
}
func NewListResultPage(getNextPage func(context.Context, ListResult) (ListResult, error)) ListResultPage {
	return original.NewListResultPage(getNextPage)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewProviderListResultIterator(page ProviderListResultPage) ProviderListResultIterator {
	return original.NewProviderListResultIterator(page)
}
func NewProviderListResultPage(getNextPage func(context.Context, ProviderListResult) (ProviderListResult, error)) ProviderListResultPage {
	return original.NewProviderListResultPage(getNextPage)
}
func NewProvidersClient(subscriptionID string) ProvidersClient {
	return original.NewProvidersClient(subscriptionID)
}
func NewProvidersClientWithBaseURI(baseURI string, subscriptionID string) ProvidersClient {
	return original.NewProvidersClientWithBaseURI(baseURI, subscriptionID)
}
func NewTagsClient(subscriptionID string) TagsClient {
	return original.NewTagsClient(subscriptionID)
}
func NewTagsClientWithBaseURI(baseURI string, subscriptionID string) TagsClient {
	return original.NewTagsClientWithBaseURI(baseURI, subscriptionID)
}
func NewTagsListResultIterator(page TagsListResultPage) TagsListResultIterator {
	return original.NewTagsListResultIterator(page)
}
func NewTagsListResultPage(getNextPage func(context.Context, TagsListResult) (TagsListResult, error)) TagsListResultPage {
	return original.NewTagsListResultPage(getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleDeploymentModeValues() []DeploymentMode {
	return original.PossibleDeploymentModeValues()
}
func PossibleOnErrorDeploymentTypeValues() []OnErrorDeploymentType {
	return original.PossibleOnErrorDeploymentTypeValues()
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
