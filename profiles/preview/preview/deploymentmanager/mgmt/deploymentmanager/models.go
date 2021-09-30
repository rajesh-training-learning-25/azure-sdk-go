//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package deploymentmanager

import original "github.com/Azure/azure-sdk-for-go/services/preview/deploymentmanager/mgmt/2019-11-01-preview/deploymentmanager"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type DeploymentMode = original.DeploymentMode

const (
	Complete    DeploymentMode = original.Complete
	Incremental DeploymentMode = original.Incremental
)

type RestAuthLocation = original.RestAuthLocation

const (
	Header RestAuthLocation = original.Header
	Query  RestAuthLocation = original.Query
)

type RestMatchQuantifier = original.RestMatchQuantifier

const (
	All RestMatchQuantifier = original.All
	Any RestMatchQuantifier = original.Any
)

type RestRequestMethod = original.RestRequestMethod

const (
	GET  RestRequestMethod = original.GET
	POST RestRequestMethod = original.POST
)

type StepType = original.StepType

const (
	StepTypeHealthCheck    StepType = original.StepTypeHealthCheck
	StepTypeStepProperties StepType = original.StepTypeStepProperties
	StepTypeWait           StepType = original.StepTypeWait
)

type Type = original.Type

const (
	TypeAuthentication Type = original.TypeAuthentication
	TypeSas            Type = original.TypeSas
)

type TypeBasicHealthCheckStepAttributes = original.TypeBasicHealthCheckStepAttributes

const (
	TypeHealthCheckStepAttributes TypeBasicHealthCheckStepAttributes = original.TypeHealthCheckStepAttributes
	TypeREST                      TypeBasicHealthCheckStepAttributes = original.TypeREST
)

type TypeBasicRestRequestAuthentication = original.TypeBasicRestRequestAuthentication

const (
	TypeAPIKey                    TypeBasicRestRequestAuthentication = original.TypeAPIKey
	TypeRestRequestAuthentication TypeBasicRestRequestAuthentication = original.TypeRestRequestAuthentication
	TypeRolloutIdentity           TypeBasicRestRequestAuthentication = original.TypeRolloutIdentity
)

type APIKeyAuthentication = original.APIKeyAuthentication
type ArtifactSource = original.ArtifactSource
type ArtifactSourceProperties = original.ArtifactSourceProperties
type ArtifactSourcePropertiesModel = original.ArtifactSourcePropertiesModel
type ArtifactSourcesClient = original.ArtifactSourcesClient
type Authentication = original.Authentication
type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type BasicAuthentication = original.BasicAuthentication
type BasicHealthCheckStepAttributes = original.BasicHealthCheckStepAttributes
type BasicRestRequestAuthentication = original.BasicRestRequestAuthentication
type BasicStepProperties = original.BasicStepProperties
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type HealthCheckStepAttributes = original.HealthCheckStepAttributes
type HealthCheckStepProperties = original.HealthCheckStepProperties
type Identity = original.Identity
type ListArtifactSource = original.ListArtifactSource
type ListRollout = original.ListRollout
type ListServiceResource = original.ListServiceResource
type ListServiceTopologyResource = original.ListServiceTopologyResource
type ListServiceUnitResource = original.ListServiceUnitResource
type ListStepResource = original.ListStepResource
type Message = original.Message
type Operation = original.Operation
type OperationDetail = original.OperationDetail
type OperationsClient = original.OperationsClient
type OperationsList = original.OperationsList
type PrePostStep = original.PrePostStep
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type ResourceOperation = original.ResourceOperation
type RestHealthCheck = original.RestHealthCheck
type RestHealthCheckStepAttributes = original.RestHealthCheckStepAttributes
type RestParameters = original.RestParameters
type RestRequest = original.RestRequest
type RestRequestAuthentication = original.RestRequestAuthentication
type RestResponse = original.RestResponse
type RestResponseRegex = original.RestResponseRegex
type Rollout = original.Rollout
type RolloutIdentityAuthentication = original.RolloutIdentityAuthentication
type RolloutOperationInfo = original.RolloutOperationInfo
type RolloutProperties = original.RolloutProperties
type RolloutPropertiesModel = original.RolloutPropertiesModel
type RolloutRequest = original.RolloutRequest
type RolloutRequestProperties = original.RolloutRequestProperties
type RolloutStep = original.RolloutStep
type RolloutsClient = original.RolloutsClient
type RolloutsCreateOrUpdateFuture = original.RolloutsCreateOrUpdateFuture
type SasAuthentication = original.SasAuthentication
type SasProperties = original.SasProperties
type Service = original.Service
type ServiceProperties = original.ServiceProperties
type ServiceResource = original.ServiceResource
type ServiceResourceProperties = original.ServiceResourceProperties
type ServiceTopologiesClient = original.ServiceTopologiesClient
type ServiceTopologyProperties = original.ServiceTopologyProperties
type ServiceTopologyResource = original.ServiceTopologyResource
type ServiceTopologyResourceProperties = original.ServiceTopologyResourceProperties
type ServiceUnit = original.ServiceUnit
type ServiceUnitArtifacts = original.ServiceUnitArtifacts
type ServiceUnitProperties = original.ServiceUnitProperties
type ServiceUnitResource = original.ServiceUnitResource
type ServiceUnitResourceProperties = original.ServiceUnitResourceProperties
type ServiceUnitsClient = original.ServiceUnitsClient
type ServiceUnitsCreateOrUpdateFuture = original.ServiceUnitsCreateOrUpdateFuture
type ServicesClient = original.ServicesClient
type StepGroup = original.StepGroup
type StepOperationInfo = original.StepOperationInfo
type StepProperties = original.StepProperties
type StepResource = original.StepResource
type StepsClient = original.StepsClient
type TrackedResource = original.TrackedResource
type WaitStepAttributes = original.WaitStepAttributes
type WaitStepProperties = original.WaitStepProperties

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewArtifactSourcesClient(subscriptionID string) ArtifactSourcesClient {
	return original.NewArtifactSourcesClient(subscriptionID)
}
func NewArtifactSourcesClientWithBaseURI(baseURI string, subscriptionID string) ArtifactSourcesClient {
	return original.NewArtifactSourcesClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewRolloutsClient(subscriptionID string) RolloutsClient {
	return original.NewRolloutsClient(subscriptionID)
}
func NewRolloutsClientWithBaseURI(baseURI string, subscriptionID string) RolloutsClient {
	return original.NewRolloutsClientWithBaseURI(baseURI, subscriptionID)
}
func NewServiceTopologiesClient(subscriptionID string) ServiceTopologiesClient {
	return original.NewServiceTopologiesClient(subscriptionID)
}
func NewServiceTopologiesClientWithBaseURI(baseURI string, subscriptionID string) ServiceTopologiesClient {
	return original.NewServiceTopologiesClientWithBaseURI(baseURI, subscriptionID)
}
func NewServiceUnitsClient(subscriptionID string) ServiceUnitsClient {
	return original.NewServiceUnitsClient(subscriptionID)
}
func NewServiceUnitsClientWithBaseURI(baseURI string, subscriptionID string) ServiceUnitsClient {
	return original.NewServiceUnitsClientWithBaseURI(baseURI, subscriptionID)
}
func NewServicesClient(subscriptionID string) ServicesClient {
	return original.NewServicesClient(subscriptionID)
}
func NewServicesClientWithBaseURI(baseURI string, subscriptionID string) ServicesClient {
	return original.NewServicesClientWithBaseURI(baseURI, subscriptionID)
}
func NewStepsClient(subscriptionID string) StepsClient {
	return original.NewStepsClient(subscriptionID)
}
func NewStepsClientWithBaseURI(baseURI string, subscriptionID string) StepsClient {
	return original.NewStepsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleDeploymentModeValues() []DeploymentMode {
	return original.PossibleDeploymentModeValues()
}
func PossibleRestAuthLocationValues() []RestAuthLocation {
	return original.PossibleRestAuthLocationValues()
}
func PossibleRestMatchQuantifierValues() []RestMatchQuantifier {
	return original.PossibleRestMatchQuantifierValues()
}
func PossibleRestRequestMethodValues() []RestRequestMethod {
	return original.PossibleRestRequestMethodValues()
}
func PossibleStepTypeValues() []StepType {
	return original.PossibleStepTypeValues()
}
func PossibleTypeBasicHealthCheckStepAttributesValues() []TypeBasicHealthCheckStepAttributes {
	return original.PossibleTypeBasicHealthCheckStepAttributesValues()
}
func PossibleTypeBasicRestRequestAuthenticationValues() []TypeBasicRestRequestAuthentication {
	return original.PossibleTypeBasicRestRequestAuthenticationValues()
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
