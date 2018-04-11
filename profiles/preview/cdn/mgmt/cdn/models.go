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

package cdn

import original "github.com/Azure/azure-sdk-for-go/services/cdn/mgmt/2017-10-12/cdn"

type OriginsClient = original.OriginsClient

func NewOriginsClient(subscriptionID string, resourceGroupName string) OriginsClient {
	return original.NewOriginsClient(subscriptionID, resourceGroupName)
}
func NewOriginsClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string) OriginsClient {
	return original.NewOriginsClientWithBaseURI(baseURI, subscriptionID, resourceGroupName)
}

type OperationsClient = original.OperationsClient

func NewOperationsClient(subscriptionID string, resourceGroupName string) OperationsClient {
	return original.NewOperationsClient(subscriptionID, resourceGroupName)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID, resourceGroupName)
}

type ResourceUsageClient = original.ResourceUsageClient

func NewResourceUsageClient(subscriptionID string, resourceGroupName string) ResourceUsageClient {
	return original.NewResourceUsageClient(subscriptionID, resourceGroupName)
}
func NewResourceUsageClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string) ResourceUsageClient {
	return original.NewResourceUsageClientWithBaseURI(baseURI, subscriptionID, resourceGroupName)
}

type CacheBehavior = original.CacheBehavior

const (
	BypassCache  CacheBehavior = original.BypassCache
	Override     CacheBehavior = original.Override
	SetIfMissing CacheBehavior = original.SetIfMissing
)

func PossibleCacheBehaviorValues() []CacheBehavior {
	return original.PossibleCacheBehaviorValues()
}

type CustomDomainResourceState = original.CustomDomainResourceState

const (
	Active   CustomDomainResourceState = original.Active
	Creating CustomDomainResourceState = original.Creating
	Deleting CustomDomainResourceState = original.Deleting
)

func PossibleCustomDomainResourceStateValues() []CustomDomainResourceState {
	return original.PossibleCustomDomainResourceStateValues()
}

type CustomHTTPSProvisioningState = original.CustomHTTPSProvisioningState

const (
	Disabled  CustomHTTPSProvisioningState = original.Disabled
	Disabling CustomHTTPSProvisioningState = original.Disabling
	Enabled   CustomHTTPSProvisioningState = original.Enabled
	Enabling  CustomHTTPSProvisioningState = original.Enabling
	Failed    CustomHTTPSProvisioningState = original.Failed
)

func PossibleCustomHTTPSProvisioningStateValues() []CustomHTTPSProvisioningState {
	return original.PossibleCustomHTTPSProvisioningStateValues()
}

type CustomHTTPSProvisioningSubstate = original.CustomHTTPSProvisioningSubstate

const (
	CertificateDeleted                            CustomHTTPSProvisioningSubstate = original.CertificateDeleted
	CertificateDeployed                           CustomHTTPSProvisioningSubstate = original.CertificateDeployed
	DeletingCertificate                           CustomHTTPSProvisioningSubstate = original.DeletingCertificate
	DeployingCertificate                          CustomHTTPSProvisioningSubstate = original.DeployingCertificate
	DomainControlValidationRequestApproved        CustomHTTPSProvisioningSubstate = original.DomainControlValidationRequestApproved
	DomainControlValidationRequestRejected        CustomHTTPSProvisioningSubstate = original.DomainControlValidationRequestRejected
	DomainControlValidationRequestTimedOut        CustomHTTPSProvisioningSubstate = original.DomainControlValidationRequestTimedOut
	IssuingCertificate                            CustomHTTPSProvisioningSubstate = original.IssuingCertificate
	PendingDomainControlValidationREquestApproval CustomHTTPSProvisioningSubstate = original.PendingDomainControlValidationREquestApproval
	SubmittingDomainControlValidationRequest      CustomHTTPSProvisioningSubstate = original.SubmittingDomainControlValidationRequest
)

func PossibleCustomHTTPSProvisioningSubstateValues() []CustomHTTPSProvisioningSubstate {
	return original.PossibleCustomHTTPSProvisioningSubstateValues()
}

type EndpointResourceState = original.EndpointResourceState

const (
	EndpointResourceStateCreating EndpointResourceState = original.EndpointResourceStateCreating
	EndpointResourceStateDeleting EndpointResourceState = original.EndpointResourceStateDeleting
	EndpointResourceStateRunning  EndpointResourceState = original.EndpointResourceStateRunning
	EndpointResourceStateStarting EndpointResourceState = original.EndpointResourceStateStarting
	EndpointResourceStateStopped  EndpointResourceState = original.EndpointResourceStateStopped
	EndpointResourceStateStopping EndpointResourceState = original.EndpointResourceStateStopping
)

func PossibleEndpointResourceStateValues() []EndpointResourceState {
	return original.PossibleEndpointResourceStateValues()
}

type GeoFilterActions = original.GeoFilterActions

const (
	Allow GeoFilterActions = original.Allow
	Block GeoFilterActions = original.Block
)

func PossibleGeoFilterActionsValues() []GeoFilterActions {
	return original.PossibleGeoFilterActionsValues()
}

type MatchType = original.MatchType

const (
	Literal  MatchType = original.Literal
	Wildcard MatchType = original.Wildcard
)

func PossibleMatchTypeValues() []MatchType {
	return original.PossibleMatchTypeValues()
}

type Name = original.Name

const (
	NameCacheExpiration    Name = original.NameCacheExpiration
	NameDeliveryRuleAction Name = original.NameDeliveryRuleAction
)

func PossibleNameValues() []Name {
	return original.PossibleNameValues()
}

type NameBasicDeliveryRuleCondition = original.NameBasicDeliveryRuleCondition

const (
	NameDeliveryRuleCondition NameBasicDeliveryRuleCondition = original.NameDeliveryRuleCondition
	NameURLFileExtension      NameBasicDeliveryRuleCondition = original.NameURLFileExtension
	NameURLPath               NameBasicDeliveryRuleCondition = original.NameURLPath
)

func PossibleNameBasicDeliveryRuleConditionValues() []NameBasicDeliveryRuleCondition {
	return original.PossibleNameBasicDeliveryRuleConditionValues()
}

type OptimizationType = original.OptimizationType

const (
	DynamicSiteAcceleration     OptimizationType = original.DynamicSiteAcceleration
	GeneralMediaStreaming       OptimizationType = original.GeneralMediaStreaming
	GeneralWebDelivery          OptimizationType = original.GeneralWebDelivery
	LargeFileDownload           OptimizationType = original.LargeFileDownload
	VideoOnDemandMediaStreaming OptimizationType = original.VideoOnDemandMediaStreaming
)

func PossibleOptimizationTypeValues() []OptimizationType {
	return original.PossibleOptimizationTypeValues()
}

type OriginResourceState = original.OriginResourceState

const (
	OriginResourceStateActive   OriginResourceState = original.OriginResourceStateActive
	OriginResourceStateCreating OriginResourceState = original.OriginResourceStateCreating
	OriginResourceStateDeleting OriginResourceState = original.OriginResourceStateDeleting
)

func PossibleOriginResourceStateValues() []OriginResourceState {
	return original.PossibleOriginResourceStateValues()
}

type ProfileResourceState = original.ProfileResourceState

const (
	ProfileResourceStateActive   ProfileResourceState = original.ProfileResourceStateActive
	ProfileResourceStateCreating ProfileResourceState = original.ProfileResourceStateCreating
	ProfileResourceStateDeleting ProfileResourceState = original.ProfileResourceStateDeleting
	ProfileResourceStateDisabled ProfileResourceState = original.ProfileResourceStateDisabled
)

func PossibleProfileResourceStateValues() []ProfileResourceState {
	return original.PossibleProfileResourceStateValues()
}

type QueryStringCachingBehavior = original.QueryStringCachingBehavior

const (
	BypassCaching     QueryStringCachingBehavior = original.BypassCaching
	IgnoreQueryString QueryStringCachingBehavior = original.IgnoreQueryString
	NotSet            QueryStringCachingBehavior = original.NotSet
	UseQueryString    QueryStringCachingBehavior = original.UseQueryString
)

func PossibleQueryStringCachingBehaviorValues() []QueryStringCachingBehavior {
	return original.PossibleQueryStringCachingBehaviorValues()
}

type ResourceType = original.ResourceType

const (
	MicrosoftCdnProfilesEndpoints ResourceType = original.MicrosoftCdnProfilesEndpoints
)

func PossibleResourceTypeValues() []ResourceType {
	return original.PossibleResourceTypeValues()
}

type SkuName = original.SkuName

const (
	CustomVerizon    SkuName = original.CustomVerizon
	PremiumVerizon   SkuName = original.PremiumVerizon
	StandardAkamai   SkuName = original.StandardAkamai
	StandardChinaCdn SkuName = original.StandardChinaCdn
	StandardVerizon  SkuName = original.StandardVerizon
)

func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}

type CacheExpirationActionParameters = original.CacheExpirationActionParameters
type CheckNameAvailabilityInput = original.CheckNameAvailabilityInput
type CheckNameAvailabilityOutput = original.CheckNameAvailabilityOutput
type CidrIPAddress = original.CidrIPAddress
type CustomDomain = original.CustomDomain
type CustomDomainListResult = original.CustomDomainListResult
type CustomDomainListResultIterator = original.CustomDomainListResultIterator
type CustomDomainListResultPage = original.CustomDomainListResultPage
type CustomDomainParameters = original.CustomDomainParameters
type CustomDomainProperties = original.CustomDomainProperties
type CustomDomainPropertiesParameters = original.CustomDomainPropertiesParameters
type CustomDomainsCreateFuture = original.CustomDomainsCreateFuture
type CustomDomainsDeleteFuture = original.CustomDomainsDeleteFuture
type DeepCreatedOrigin = original.DeepCreatedOrigin
type DeepCreatedOriginProperties = original.DeepCreatedOriginProperties
type DeliveryRule = original.DeliveryRule
type BasicDeliveryRuleAction = original.BasicDeliveryRuleAction
type DeliveryRuleAction = original.DeliveryRuleAction
type DeliveryRuleCacheExpirationAction = original.DeliveryRuleCacheExpirationAction
type BasicDeliveryRuleCondition = original.BasicDeliveryRuleCondition
type DeliveryRuleCondition = original.DeliveryRuleCondition
type DeliveryRuleURLFileExtensionCondition = original.DeliveryRuleURLFileExtensionCondition
type DeliveryRuleURLPathCondition = original.DeliveryRuleURLPathCondition
type EdgeNode = original.EdgeNode
type EdgeNodeProperties = original.EdgeNodeProperties
type EdgenodeResult = original.EdgenodeResult
type EdgenodeResultIterator = original.EdgenodeResultIterator
type EdgenodeResultPage = original.EdgenodeResultPage
type Endpoint = original.Endpoint
type EndpointListResult = original.EndpointListResult
type EndpointListResultIterator = original.EndpointListResultIterator
type EndpointListResultPage = original.EndpointListResultPage
type EndpointProperties = original.EndpointProperties
type EndpointPropertiesUpdateParameters = original.EndpointPropertiesUpdateParameters
type EndpointPropertiesUpdateParametersDeliveryPolicy = original.EndpointPropertiesUpdateParametersDeliveryPolicy
type EndpointsCreateFuture = original.EndpointsCreateFuture
type EndpointsDeleteFuture = original.EndpointsDeleteFuture
type EndpointsLoadContentFuture = original.EndpointsLoadContentFuture
type EndpointsPurgeContentFuture = original.EndpointsPurgeContentFuture
type EndpointsStartFuture = original.EndpointsStartFuture
type EndpointsStopFuture = original.EndpointsStopFuture
type EndpointsUpdateFuture = original.EndpointsUpdateFuture
type EndpointUpdateParameters = original.EndpointUpdateParameters
type ErrorResponse = original.ErrorResponse
type GeoFilter = original.GeoFilter
type IPAddressGroup = original.IPAddressGroup
type LoadParameters = original.LoadParameters
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationsListResult = original.OperationsListResult
type OperationsListResultIterator = original.OperationsListResultIterator
type OperationsListResultPage = original.OperationsListResultPage
type Origin = original.Origin
type OriginListResult = original.OriginListResult
type OriginListResultIterator = original.OriginListResultIterator
type OriginListResultPage = original.OriginListResultPage
type OriginProperties = original.OriginProperties
type OriginPropertiesParameters = original.OriginPropertiesParameters
type OriginsUpdateFuture = original.OriginsUpdateFuture
type OriginUpdateParameters = original.OriginUpdateParameters
type Profile = original.Profile
type ProfileListResult = original.ProfileListResult
type ProfileListResultIterator = original.ProfileListResultIterator
type ProfileListResultPage = original.ProfileListResultPage
type ProfileProperties = original.ProfileProperties
type ProfilesCreateFuture = original.ProfilesCreateFuture
type ProfilesDeleteFuture = original.ProfilesDeleteFuture
type ProfilesUpdateFuture = original.ProfilesUpdateFuture
type ProfileUpdateParameters = original.ProfileUpdateParameters
type ProxyResource = original.ProxyResource
type PurgeParameters = original.PurgeParameters
type Resource = original.Resource
type ResourceUsage = original.ResourceUsage
type ResourceUsageListResult = original.ResourceUsageListResult
type ResourceUsageListResultIterator = original.ResourceUsageListResultIterator
type ResourceUsageListResultPage = original.ResourceUsageListResultPage
type Sku = original.Sku
type SsoURI = original.SsoURI
type SupportedOptimizationTypesListResult = original.SupportedOptimizationTypesListResult
type TrackedResource = original.TrackedResource
type URLFileExtensionConditionParameters = original.URLFileExtensionConditionParameters
type URLPathConditionParameters = original.URLPathConditionParameters
type ValidateCustomDomainInput = original.ValidateCustomDomainInput
type ValidateCustomDomainOutput = original.ValidateCustomDomainOutput
type ValidateProbeInput = original.ValidateProbeInput
type ValidateProbeOutput = original.ValidateProbeOutput
type ProfilesClient = original.ProfilesClient

func NewProfilesClient(subscriptionID string, resourceGroupName string) ProfilesClient {
	return original.NewProfilesClient(subscriptionID, resourceGroupName)
}
func NewProfilesClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string) ProfilesClient {
	return original.NewProfilesClientWithBaseURI(baseURI, subscriptionID, resourceGroupName)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}

type EndpointsClient = original.EndpointsClient

func NewEndpointsClient(subscriptionID string, resourceGroupName string) EndpointsClient {
	return original.NewEndpointsClient(subscriptionID, resourceGroupName)
}
func NewEndpointsClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string) EndpointsClient {
	return original.NewEndpointsClientWithBaseURI(baseURI, subscriptionID, resourceGroupName)
}

type CustomDomainsClient = original.CustomDomainsClient

func NewCustomDomainsClient(subscriptionID string, resourceGroupName string) CustomDomainsClient {
	return original.NewCustomDomainsClient(subscriptionID, resourceGroupName)
}
func NewCustomDomainsClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string) CustomDomainsClient {
	return original.NewCustomDomainsClientWithBaseURI(baseURI, subscriptionID, resourceGroupName)
}

type EdgeNodesClient = original.EdgeNodesClient

func NewEdgeNodesClient(subscriptionID string, resourceGroupName string) EdgeNodesClient {
	return original.NewEdgeNodesClient(subscriptionID, resourceGroupName)
}
func NewEdgeNodesClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string) EdgeNodesClient {
	return original.NewEdgeNodesClientWithBaseURI(baseURI, subscriptionID, resourceGroupName)
}

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient

func New(subscriptionID string, resourceGroupName string) BaseClient {
	return original.New(subscriptionID, resourceGroupName)
}
func NewWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID, resourceGroupName)
}
