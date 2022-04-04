//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package diagnostics

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2021-05-01-preview/diagnostics"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AccessMode = original.AccessMode

const (
	AccessModeOpen        AccessMode = original.AccessModeOpen
	AccessModePrivateOnly AccessMode = original.AccessModePrivateOnly
)

type CategoryType = original.CategoryType

const (
	CategoryTypeLogs    CategoryType = original.CategoryTypeLogs
	CategoryTypeMetrics CategoryType = original.CategoryTypeMetrics
)

type ComparisonOperationType = original.ComparisonOperationType

const (
	ComparisonOperationTypeEquals             ComparisonOperationType = original.ComparisonOperationTypeEquals
	ComparisonOperationTypeGreaterThan        ComparisonOperationType = original.ComparisonOperationTypeGreaterThan
	ComparisonOperationTypeGreaterThanOrEqual ComparisonOperationType = original.ComparisonOperationTypeGreaterThanOrEqual
	ComparisonOperationTypeLessThan           ComparisonOperationType = original.ComparisonOperationTypeLessThan
	ComparisonOperationTypeLessThanOrEqual    ComparisonOperationType = original.ComparisonOperationTypeLessThanOrEqual
	ComparisonOperationTypeNotEquals          ComparisonOperationType = original.ComparisonOperationTypeNotEquals
)

type CreatedByType = original.CreatedByType

const (
	CreatedByTypeApplication     CreatedByType = original.CreatedByTypeApplication
	CreatedByTypeKey             CreatedByType = original.CreatedByTypeKey
	CreatedByTypeManagedIdentity CreatedByType = original.CreatedByTypeManagedIdentity
	CreatedByTypeUser            CreatedByType = original.CreatedByTypeUser
)

type MetricStatisticType = original.MetricStatisticType

const (
	MetricStatisticTypeAverage MetricStatisticType = original.MetricStatisticTypeAverage
	MetricStatisticTypeCount   MetricStatisticType = original.MetricStatisticTypeCount
	MetricStatisticTypeMax     MetricStatisticType = original.MetricStatisticTypeMax
	MetricStatisticTypeMin     MetricStatisticType = original.MetricStatisticTypeMin
	MetricStatisticTypeSum     MetricStatisticType = original.MetricStatisticTypeSum
)

type PredictiveAutoscalePolicyScaleMode = original.PredictiveAutoscalePolicyScaleMode

const (
	PredictiveAutoscalePolicyScaleModeDisabled     PredictiveAutoscalePolicyScaleMode = original.PredictiveAutoscalePolicyScaleModeDisabled
	PredictiveAutoscalePolicyScaleModeEnabled      PredictiveAutoscalePolicyScaleMode = original.PredictiveAutoscalePolicyScaleModeEnabled
	PredictiveAutoscalePolicyScaleModeForecastOnly PredictiveAutoscalePolicyScaleMode = original.PredictiveAutoscalePolicyScaleModeForecastOnly
)

type PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningState

const (
	PrivateEndpointConnectionProvisioningStateCreating  PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateCreating
	PrivateEndpointConnectionProvisioningStateDeleting  PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateDeleting
	PrivateEndpointConnectionProvisioningStateFailed    PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateFailed
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateSucceeded
)

type PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatus

const (
	PrivateEndpointServiceConnectionStatusApproved PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatusApproved
	PrivateEndpointServiceConnectionStatusPending  PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatusPending
	PrivateEndpointServiceConnectionStatusRejected PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatusRejected
)

type ReceiverStatus = original.ReceiverStatus

const (
	ReceiverStatusDisabled     ReceiverStatus = original.ReceiverStatusDisabled
	ReceiverStatusEnabled      ReceiverStatus = original.ReceiverStatusEnabled
	ReceiverStatusNotSpecified ReceiverStatus = original.ReceiverStatusNotSpecified
)

type RecurrenceFrequency = original.RecurrenceFrequency

const (
	RecurrenceFrequencyDay    RecurrenceFrequency = original.RecurrenceFrequencyDay
	RecurrenceFrequencyHour   RecurrenceFrequency = original.RecurrenceFrequencyHour
	RecurrenceFrequencyMinute RecurrenceFrequency = original.RecurrenceFrequencyMinute
	RecurrenceFrequencyMonth  RecurrenceFrequency = original.RecurrenceFrequencyMonth
	RecurrenceFrequencyNone   RecurrenceFrequency = original.RecurrenceFrequencyNone
	RecurrenceFrequencySecond RecurrenceFrequency = original.RecurrenceFrequencySecond
	RecurrenceFrequencyWeek   RecurrenceFrequency = original.RecurrenceFrequencyWeek
	RecurrenceFrequencyYear   RecurrenceFrequency = original.RecurrenceFrequencyYear
)

type ScaleDirection = original.ScaleDirection

const (
	ScaleDirectionDecrease ScaleDirection = original.ScaleDirectionDecrease
	ScaleDirectionIncrease ScaleDirection = original.ScaleDirectionIncrease
	ScaleDirectionNone     ScaleDirection = original.ScaleDirectionNone
)

type ScaleRuleMetricDimensionOperationType = original.ScaleRuleMetricDimensionOperationType

const (
	ScaleRuleMetricDimensionOperationTypeEquals    ScaleRuleMetricDimensionOperationType = original.ScaleRuleMetricDimensionOperationTypeEquals
	ScaleRuleMetricDimensionOperationTypeNotEquals ScaleRuleMetricDimensionOperationType = original.ScaleRuleMetricDimensionOperationTypeNotEquals
)

type ScaleType = original.ScaleType

const (
	ScaleTypeChangeCount             ScaleType = original.ScaleTypeChangeCount
	ScaleTypeExactCount              ScaleType = original.ScaleTypeExactCount
	ScaleTypePercentChangeCount      ScaleType = original.ScaleTypePercentChangeCount
	ScaleTypeServiceAllowedNextValue ScaleType = original.ScaleTypeServiceAllowedNextValue
)

type TimeAggregationType = original.TimeAggregationType

const (
	TimeAggregationTypeAverage TimeAggregationType = original.TimeAggregationTypeAverage
	TimeAggregationTypeCount   TimeAggregationType = original.TimeAggregationTypeCount
	TimeAggregationTypeLast    TimeAggregationType = original.TimeAggregationTypeLast
	TimeAggregationTypeMaximum TimeAggregationType = original.TimeAggregationTypeMaximum
	TimeAggregationTypeMinimum TimeAggregationType = original.TimeAggregationTypeMinimum
	TimeAggregationTypeTotal   TimeAggregationType = original.TimeAggregationTypeTotal
)

type AccessModeSettings = original.AccessModeSettings
type AccessModeSettingsExclusion = original.AccessModeSettingsExclusion
type ActionDetail = original.ActionDetail
type ActionGroup = original.ActionGroup
type ActionGroupList = original.ActionGroupList
type ActionGroupPatch = original.ActionGroupPatch
type ActionGroupPatchBody = original.ActionGroupPatchBody
type ActionGroupResource = original.ActionGroupResource
type ActionGroupsClient = original.ActionGroupsClient
type ActionGroupsPostTestNotificationsFuture = original.ActionGroupsPostTestNotificationsFuture
type ArmRoleReceiver = original.ArmRoleReceiver
type AutomationRunbookReceiver = original.AutomationRunbookReceiver
type AutoscaleErrorResponse = original.AutoscaleErrorResponse
type AutoscaleErrorResponseError = original.AutoscaleErrorResponseError
type AutoscaleNotification = original.AutoscaleNotification
type AutoscaleProfile = original.AutoscaleProfile
type AutoscaleSetting = original.AutoscaleSetting
type AutoscaleSettingResource = original.AutoscaleSettingResource
type AutoscaleSettingResourceCollection = original.AutoscaleSettingResourceCollection
type AutoscaleSettingResourceCollectionIterator = original.AutoscaleSettingResourceCollectionIterator
type AutoscaleSettingResourceCollectionPage = original.AutoscaleSettingResourceCollectionPage
type AutoscaleSettingResourcePatch = original.AutoscaleSettingResourcePatch
type AutoscaleSettingsClient = original.AutoscaleSettingsClient
type AzureAppPushReceiver = original.AzureAppPushReceiver
type AzureEntityResource = original.AzureEntityResource
type AzureFunctionReceiver = original.AzureFunctionReceiver
type AzureMonitorPrivateLinkScope = original.AzureMonitorPrivateLinkScope
type AzureMonitorPrivateLinkScopeListResult = original.AzureMonitorPrivateLinkScopeListResult
type AzureMonitorPrivateLinkScopeListResultIterator = original.AzureMonitorPrivateLinkScopeListResultIterator
type AzureMonitorPrivateLinkScopeListResultPage = original.AzureMonitorPrivateLinkScopeListResultPage
type AzureMonitorPrivateLinkScopeProperties = original.AzureMonitorPrivateLinkScopeProperties
type AzureResource = original.AzureResource
type BaseClient = original.BaseClient
type Context = original.Context
type DefaultErrorResponse = original.DefaultErrorResponse
type EmailNotification = original.EmailNotification
type EmailReceiver = original.EmailReceiver
type EnableRequest = original.EnableRequest
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorDetail = original.ErrorDetail
type ErrorResponse = original.ErrorResponse
type EventHubReceiver = original.EventHubReceiver
type ItsmReceiver = original.ItsmReceiver
type LogSettings = original.LogSettings
type LogicAppReceiver = original.LogicAppReceiver
type ManagementGroupDiagnosticSettings = original.ManagementGroupDiagnosticSettings
type ManagementGroupDiagnosticSettingsClient = original.ManagementGroupDiagnosticSettingsClient
type ManagementGroupDiagnosticSettingsResource = original.ManagementGroupDiagnosticSettingsResource
type ManagementGroupDiagnosticSettingsResourceCollection = original.ManagementGroupDiagnosticSettingsResourceCollection
type ManagementGroupLogSettings = original.ManagementGroupLogSettings
type MetricSettings = original.MetricSettings
type MetricTrigger = original.MetricTrigger
type NotificationRequestBody = original.NotificationRequestBody
type OperationStatus = original.OperationStatus
type PredictiveAutoscalePolicy = original.PredictiveAutoscalePolicy
type PredictiveMetricClient = original.PredictiveMetricClient
type PredictiveResponse = original.PredictiveResponse
type PredictiveValue = original.PredictiveValue
type PrivateEndpoint = original.PrivateEndpoint
type PrivateEndpointConnection = original.PrivateEndpointConnection
type PrivateEndpointConnectionListResult = original.PrivateEndpointConnectionListResult
type PrivateEndpointConnectionProperties = original.PrivateEndpointConnectionProperties
type PrivateEndpointConnectionsClient = original.PrivateEndpointConnectionsClient
type PrivateEndpointConnectionsCreateOrUpdateFuture = original.PrivateEndpointConnectionsCreateOrUpdateFuture
type PrivateEndpointConnectionsDeleteFuture = original.PrivateEndpointConnectionsDeleteFuture
type PrivateLinkResource = original.PrivateLinkResource
type PrivateLinkResourceListResult = original.PrivateLinkResourceListResult
type PrivateLinkResourceProperties = original.PrivateLinkResourceProperties
type PrivateLinkResourcesClient = original.PrivateLinkResourcesClient
type PrivateLinkScopeOperationStatusClient = original.PrivateLinkScopeOperationStatusClient
type PrivateLinkScopedResourcesClient = original.PrivateLinkScopedResourcesClient
type PrivateLinkScopedResourcesCreateOrUpdateFuture = original.PrivateLinkScopedResourcesCreateOrUpdateFuture
type PrivateLinkScopedResourcesDeleteFuture = original.PrivateLinkScopedResourcesDeleteFuture
type PrivateLinkScopesClient = original.PrivateLinkScopesClient
type PrivateLinkScopesDeleteFuture = original.PrivateLinkScopesDeleteFuture
type PrivateLinkServiceConnectionState = original.PrivateLinkServiceConnectionState
type ProxyResource = original.ProxyResource
type Recurrence = original.Recurrence
type RecurrentSchedule = original.RecurrentSchedule
type Resource = original.Resource
type RetentionPolicy = original.RetentionPolicy
type ScaleAction = original.ScaleAction
type ScaleCapacity = original.ScaleCapacity
type ScaleRule = original.ScaleRule
type ScaleRuleMetricDimension = original.ScaleRuleMetricDimension
type ScopedResource = original.ScopedResource
type ScopedResourceListResult = original.ScopedResourceListResult
type ScopedResourceListResultIterator = original.ScopedResourceListResultIterator
type ScopedResourceListResultPage = original.ScopedResourceListResultPage
type ScopedResourceProperties = original.ScopedResourceProperties
type Settings = original.Settings
type SettingsCategory = original.SettingsCategory
type SettingsCategoryClient = original.SettingsCategoryClient
type SettingsCategoryResource = original.SettingsCategoryResource
type SettingsCategoryResourceCollection = original.SettingsCategoryResourceCollection
type SettingsClient = original.SettingsClient
type SettingsResource = original.SettingsResource
type SettingsResourceCollection = original.SettingsResourceCollection
type SmsReceiver = original.SmsReceiver
type SubscriptionDiagnosticSettings = original.SubscriptionDiagnosticSettings
type SubscriptionDiagnosticSettingsClient = original.SubscriptionDiagnosticSettingsClient
type SubscriptionDiagnosticSettingsResource = original.SubscriptionDiagnosticSettingsResource
type SubscriptionDiagnosticSettingsResourceCollection = original.SubscriptionDiagnosticSettingsResourceCollection
type SubscriptionLogSettings = original.SubscriptionLogSettings
type SystemData = original.SystemData
type TagsResource = original.TagsResource
type TestNotificationDetailsResponse = original.TestNotificationDetailsResponse
type TestNotificationResponse = original.TestNotificationResponse
type TimeWindow = original.TimeWindow
type TrackedResource = original.TrackedResource
type VoiceReceiver = original.VoiceReceiver
type WebhookNotification = original.WebhookNotification
type WebhookReceiver = original.WebhookReceiver

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewActionGroupsClient(subscriptionID string) ActionGroupsClient {
	return original.NewActionGroupsClient(subscriptionID)
}
func NewActionGroupsClientWithBaseURI(baseURI string, subscriptionID string) ActionGroupsClient {
	return original.NewActionGroupsClientWithBaseURI(baseURI, subscriptionID)
}
func NewAutoscaleSettingResourceCollectionIterator(page AutoscaleSettingResourceCollectionPage) AutoscaleSettingResourceCollectionIterator {
	return original.NewAutoscaleSettingResourceCollectionIterator(page)
}
func NewAutoscaleSettingResourceCollectionPage(cur AutoscaleSettingResourceCollection, getNextPage func(context.Context, AutoscaleSettingResourceCollection) (AutoscaleSettingResourceCollection, error)) AutoscaleSettingResourceCollectionPage {
	return original.NewAutoscaleSettingResourceCollectionPage(cur, getNextPage)
}
func NewAutoscaleSettingsClient(subscriptionID string) AutoscaleSettingsClient {
	return original.NewAutoscaleSettingsClient(subscriptionID)
}
func NewAutoscaleSettingsClientWithBaseURI(baseURI string, subscriptionID string) AutoscaleSettingsClient {
	return original.NewAutoscaleSettingsClientWithBaseURI(baseURI, subscriptionID)
}
func NewAzureMonitorPrivateLinkScopeListResultIterator(page AzureMonitorPrivateLinkScopeListResultPage) AzureMonitorPrivateLinkScopeListResultIterator {
	return original.NewAzureMonitorPrivateLinkScopeListResultIterator(page)
}
func NewAzureMonitorPrivateLinkScopeListResultPage(cur AzureMonitorPrivateLinkScopeListResult, getNextPage func(context.Context, AzureMonitorPrivateLinkScopeListResult) (AzureMonitorPrivateLinkScopeListResult, error)) AzureMonitorPrivateLinkScopeListResultPage {
	return original.NewAzureMonitorPrivateLinkScopeListResultPage(cur, getNextPage)
}
func NewManagementGroupDiagnosticSettingsClient(subscriptionID string) ManagementGroupDiagnosticSettingsClient {
	return original.NewManagementGroupDiagnosticSettingsClient(subscriptionID)
}
func NewManagementGroupDiagnosticSettingsClientWithBaseURI(baseURI string, subscriptionID string) ManagementGroupDiagnosticSettingsClient {
	return original.NewManagementGroupDiagnosticSettingsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPredictiveMetricClient(subscriptionID string) PredictiveMetricClient {
	return original.NewPredictiveMetricClient(subscriptionID)
}
func NewPredictiveMetricClientWithBaseURI(baseURI string, subscriptionID string) PredictiveMetricClient {
	return original.NewPredictiveMetricClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateEndpointConnectionsClient(subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClient(subscriptionID)
}
func NewPrivateEndpointConnectionsClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateLinkResourcesClient(subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClient(subscriptionID)
}
func NewPrivateLinkResourcesClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateLinkScopeOperationStatusClient(subscriptionID string) PrivateLinkScopeOperationStatusClient {
	return original.NewPrivateLinkScopeOperationStatusClient(subscriptionID)
}
func NewPrivateLinkScopeOperationStatusClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkScopeOperationStatusClient {
	return original.NewPrivateLinkScopeOperationStatusClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateLinkScopedResourcesClient(subscriptionID string) PrivateLinkScopedResourcesClient {
	return original.NewPrivateLinkScopedResourcesClient(subscriptionID)
}
func NewPrivateLinkScopedResourcesClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkScopedResourcesClient {
	return original.NewPrivateLinkScopedResourcesClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateLinkScopesClient(subscriptionID string) PrivateLinkScopesClient {
	return original.NewPrivateLinkScopesClient(subscriptionID)
}
func NewPrivateLinkScopesClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkScopesClient {
	return original.NewPrivateLinkScopesClientWithBaseURI(baseURI, subscriptionID)
}
func NewScopedResourceListResultIterator(page ScopedResourceListResultPage) ScopedResourceListResultIterator {
	return original.NewScopedResourceListResultIterator(page)
}
func NewScopedResourceListResultPage(cur ScopedResourceListResult, getNextPage func(context.Context, ScopedResourceListResult) (ScopedResourceListResult, error)) ScopedResourceListResultPage {
	return original.NewScopedResourceListResultPage(cur, getNextPage)
}
func NewSettingsCategoryClient(subscriptionID string) SettingsCategoryClient {
	return original.NewSettingsCategoryClient(subscriptionID)
}
func NewSettingsCategoryClientWithBaseURI(baseURI string, subscriptionID string) SettingsCategoryClient {
	return original.NewSettingsCategoryClientWithBaseURI(baseURI, subscriptionID)
}
func NewSettingsClient(subscriptionID string) SettingsClient {
	return original.NewSettingsClient(subscriptionID)
}
func NewSettingsClientWithBaseURI(baseURI string, subscriptionID string) SettingsClient {
	return original.NewSettingsClientWithBaseURI(baseURI, subscriptionID)
}
func NewSubscriptionDiagnosticSettingsClient(subscriptionID string) SubscriptionDiagnosticSettingsClient {
	return original.NewSubscriptionDiagnosticSettingsClient(subscriptionID)
}
func NewSubscriptionDiagnosticSettingsClientWithBaseURI(baseURI string, subscriptionID string) SubscriptionDiagnosticSettingsClient {
	return original.NewSubscriptionDiagnosticSettingsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAccessModeValues() []AccessMode {
	return original.PossibleAccessModeValues()
}
func PossibleCategoryTypeValues() []CategoryType {
	return original.PossibleCategoryTypeValues()
}
func PossibleComparisonOperationTypeValues() []ComparisonOperationType {
	return original.PossibleComparisonOperationTypeValues()
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleMetricStatisticTypeValues() []MetricStatisticType {
	return original.PossibleMetricStatisticTypeValues()
}
func PossiblePredictiveAutoscalePolicyScaleModeValues() []PredictiveAutoscalePolicyScaleMode {
	return original.PossiblePredictiveAutoscalePolicyScaleModeValues()
}
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return original.PossiblePrivateEndpointConnectionProvisioningStateValues()
}
func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return original.PossiblePrivateEndpointServiceConnectionStatusValues()
}
func PossibleReceiverStatusValues() []ReceiverStatus {
	return original.PossibleReceiverStatusValues()
}
func PossibleRecurrenceFrequencyValues() []RecurrenceFrequency {
	return original.PossibleRecurrenceFrequencyValues()
}
func PossibleScaleDirectionValues() []ScaleDirection {
	return original.PossibleScaleDirectionValues()
}
func PossibleScaleRuleMetricDimensionOperationTypeValues() []ScaleRuleMetricDimensionOperationType {
	return original.PossibleScaleRuleMetricDimensionOperationTypeValues()
}
func PossibleScaleTypeValues() []ScaleType {
	return original.PossibleScaleTypeValues()
}
func PossibleTimeAggregationTypeValues() []TimeAggregationType {
	return original.PossibleTimeAggregationTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
