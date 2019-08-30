// +build go1.9

// Copyright 2019 Microsoft Corporation
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

package managednetwork

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/managednetwork/mgmt/2019-06-01-preview/managednetwork"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type Kind = original.Kind

const (
	Connectivity Kind = original.Connectivity
)

type ProvisioningState = original.ProvisioningState

const (
	Deleting  ProvisioningState = original.Deleting
	Failed    ProvisioningState = original.Failed
	Succeeded ProvisioningState = original.Succeeded
	Updating  ProvisioningState = original.Updating
)

type Type = original.Type

const (
	HubAndSpokeTopology Type = original.HubAndSpokeTopology
	MeshTopology        Type = original.MeshTopology
)

type BaseClient = original.BaseClient
type ConnectivityCollection = original.ConnectivityCollection
type ErrorResponse = original.ErrorResponse
type Group = original.Group
type GroupListResult = original.GroupListResult
type GroupListResultIterator = original.GroupListResultIterator
type GroupListResultPage = original.GroupListResultPage
type GroupProperties = original.GroupProperties
type GroupsClient = original.GroupsClient
type GroupsCreateOrUpdateFuture = original.GroupsCreateOrUpdateFuture
type GroupsDeleteFuture = original.GroupsDeleteFuture
type HubAndSpokePeeringPolicyProperties = original.HubAndSpokePeeringPolicyProperties
type ListResult = original.ListResult
type ListResultIterator = original.ListResultIterator
type ListResultPage = original.ListResultPage
type ManagedNetwork = original.ManagedNetwork
type ManagedNetworksClient = original.ManagedNetworksClient
type ManagedNetworksDeleteFutureType = original.ManagedNetworksDeleteFutureType
type ManagedNetworksUpdateFutureType = original.ManagedNetworksUpdateFutureType
type MeshPeeringPolicyProperties = original.MeshPeeringPolicyProperties
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type PeeringPoliciesClient = original.PeeringPoliciesClient
type PeeringPoliciesCreateOrUpdateFuture = original.PeeringPoliciesCreateOrUpdateFuture
type PeeringPoliciesDeleteFuture = original.PeeringPoliciesDeleteFuture
type PeeringPolicy = original.PeeringPolicy
type PeeringPolicyListResult = original.PeeringPolicyListResult
type PeeringPolicyListResultIterator = original.PeeringPolicyListResultIterator
type PeeringPolicyListResultPage = original.PeeringPolicyListResultPage
type PeeringPolicyProperties = original.PeeringPolicyProperties
type Properties = original.Properties
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type ResourceID = original.ResourceID
type ResourceProperties = original.ResourceProperties
type Scope = original.Scope
type ScopeAssignment = original.ScopeAssignment
type ScopeAssignmentListResult = original.ScopeAssignmentListResult
type ScopeAssignmentListResultIterator = original.ScopeAssignmentListResultIterator
type ScopeAssignmentListResultPage = original.ScopeAssignmentListResultPage
type ScopeAssignmentProperties = original.ScopeAssignmentProperties
type ScopeAssignmentsClient = original.ScopeAssignmentsClient
type TrackedResource = original.TrackedResource
type Update = original.Update

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
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
func NewManagedNetworksClient(subscriptionID string) ManagedNetworksClient {
	return original.NewManagedNetworksClient(subscriptionID)
}
func NewManagedNetworksClientWithBaseURI(baseURI string, subscriptionID string) ManagedNetworksClient {
	return original.NewManagedNetworksClientWithBaseURI(baseURI, subscriptionID)
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
func NewPeeringPoliciesClient(subscriptionID string) PeeringPoliciesClient {
	return original.NewPeeringPoliciesClient(subscriptionID)
}
func NewPeeringPoliciesClientWithBaseURI(baseURI string, subscriptionID string) PeeringPoliciesClient {
	return original.NewPeeringPoliciesClientWithBaseURI(baseURI, subscriptionID)
}
func NewPeeringPolicyListResultIterator(page PeeringPolicyListResultPage) PeeringPolicyListResultIterator {
	return original.NewPeeringPolicyListResultIterator(page)
}
func NewPeeringPolicyListResultPage(getNextPage func(context.Context, PeeringPolicyListResult) (PeeringPolicyListResult, error)) PeeringPolicyListResultPage {
	return original.NewPeeringPolicyListResultPage(getNextPage)
}
func NewScopeAssignmentListResultIterator(page ScopeAssignmentListResultPage) ScopeAssignmentListResultIterator {
	return original.NewScopeAssignmentListResultIterator(page)
}
func NewScopeAssignmentListResultPage(getNextPage func(context.Context, ScopeAssignmentListResult) (ScopeAssignmentListResult, error)) ScopeAssignmentListResultPage {
	return original.NewScopeAssignmentListResultPage(getNextPage)
}
func NewScopeAssignmentsClient(subscriptionID string) ScopeAssignmentsClient {
	return original.NewScopeAssignmentsClient(subscriptionID)
}
func NewScopeAssignmentsClientWithBaseURI(baseURI string, subscriptionID string) ScopeAssignmentsClient {
	return original.NewScopeAssignmentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleKindValues() []Kind {
	return original.PossibleKindValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
