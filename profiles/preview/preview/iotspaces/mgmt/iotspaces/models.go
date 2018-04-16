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

package iotspaces

import (
	original "github.com/Azure/azure-sdk-for-go/services/preview/iotspaces/mgmt/2017-10-01-preview/iotspaces"
	uuid "github.com/satori/go.uuid"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type Client = original.Client
type OperationsClient = original.OperationsClient
type NameUnavailabilityReason = original.NameUnavailabilityReason

const (
	AlreadyExists NameUnavailabilityReason = original.AlreadyExists
	Invalid       NameUnavailabilityReason = original.Invalid
)

type ProvisioningState = original.ProvisioningState

const (
	Canceled     ProvisioningState = original.Canceled
	Deleting     ProvisioningState = original.Deleting
	Failed       ProvisioningState = original.Failed
	Provisioning ProvisioningState = original.Provisioning
	Succeeded    ProvisioningState = original.Succeeded
)

type Sku = original.Sku

const (
	F1 Sku = original.F1
	S1 Sku = original.S1
	S2 Sku = original.S2
	S3 Sku = original.S3
)

type CreateOrUpdateFuture = original.CreateOrUpdateFuture
type DeleteFuture = original.DeleteFuture
type Description = original.Description
type DescriptionListResult = original.DescriptionListResult
type DescriptionListResultIterator = original.DescriptionListResultIterator
type DescriptionListResultPage = original.DescriptionListResultPage
type ErrorDetails = original.ErrorDetails
type NameAvailabilityInfo = original.NameAvailabilityInfo
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationInputs = original.OperationInputs
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type PatchDescription = original.PatchDescription
type Properties = original.Properties
type Resource = original.Resource
type SkuInfo = original.SkuInfo
type StorageContainerProperties = original.StorageContainerProperties
type UpdateFuture = original.UpdateFuture

func New(subscriptionID uuid.UUID) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID uuid.UUID) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewClient(subscriptionID uuid.UUID) Client {
	return original.NewClient(subscriptionID)
}
func NewClientWithBaseURI(baseURI string, subscriptionID uuid.UUID) Client {
	return original.NewClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID uuid.UUID) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID uuid.UUID) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func PossibleNameUnavailabilityReasonValues() []NameUnavailabilityReason {
	return original.PossibleNameUnavailabilityReasonValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleSkuValues() []Sku {
	return original.PossibleSkuValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
