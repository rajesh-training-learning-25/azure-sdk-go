//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package visualstudio

import original "github.com/Azure/azure-sdk-for-go/services/preview/visualstudio/mgmt/2014-04-01-preview/visualstudio"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AccountResource = original.AccountResource
type AccountResourceListResult = original.AccountResourceListResult
type AccountResourceRequest = original.AccountResourceRequest
type AccountTagRequest = original.AccountTagRequest
type AccountsClient = original.AccountsClient
type BaseClient = original.BaseClient
type CheckNameAvailabilityParameter = original.CheckNameAvailabilityParameter
type CheckNameAvailabilityResult = original.CheckNameAvailabilityResult
type ExtensionResource = original.ExtensionResource
type ExtensionResourceListResult = original.ExtensionResourceListResult
type ExtensionResourcePlan = original.ExtensionResourcePlan
type ExtensionResourceRequest = original.ExtensionResourceRequest
type ExtensionsClient = original.ExtensionsClient
type Operation = original.Operation
type OperationListResult = original.OperationListResult
type OperationProperties = original.OperationProperties
type OperationsClient = original.OperationsClient
type ProjectResource = original.ProjectResource
type ProjectResourceListResult = original.ProjectResourceListResult
type ProjectsClient = original.ProjectsClient
type ProjectsCreateFuture = original.ProjectsCreateFuture
type Resource = original.Resource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAccountsClient(subscriptionID string) AccountsClient {
	return original.NewAccountsClient(subscriptionID)
}
func NewAccountsClientWithBaseURI(baseURI string, subscriptionID string) AccountsClient {
	return original.NewAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewExtensionsClient(subscriptionID string) ExtensionsClient {
	return original.NewExtensionsClient(subscriptionID)
}
func NewExtensionsClientWithBaseURI(baseURI string, subscriptionID string) ExtensionsClient {
	return original.NewExtensionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewProjectsClient(subscriptionID string) ProjectsClient {
	return original.NewProjectsClient(subscriptionID)
}
func NewProjectsClientWithBaseURI(baseURI string, subscriptionID string) ProjectsClient {
	return original.NewProjectsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
