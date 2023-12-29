//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcostmanagement

// AlertsClientDismissOptions contains the optional parameters for the AlertsClient.Dismiss method.
type AlertsClientDismissOptions struct {
	// placeholder for future optional parameters
}

// AlertsClientGetOptions contains the optional parameters for the AlertsClient.Get method.
type AlertsClientGetOptions struct {
	// placeholder for future optional parameters
}

// AlertsClientListExternalOptions contains the optional parameters for the AlertsClient.ListExternal method.
type AlertsClientListExternalOptions struct {
	// placeholder for future optional parameters
}

// AlertsClientListOptions contains the optional parameters for the AlertsClient.List method.
type AlertsClientListOptions struct {
	// placeholder for future optional parameters
}

// BenefitRecommendationsClientListOptions contains the optional parameters for the BenefitRecommendationsClient.NewListPager
// method.
type BenefitRecommendationsClientListOptions struct {
	// May be used to expand the properties by: properties/usage, properties/allRecommendationDetails
	Expand *string

	// Can be used to filter benefitRecommendations by: properties/scope with allowed values ['Single', 'Shared'] and default
	// value 'Shared'; and properties/lookBackPeriod with allowed values ['Last7Days',
	// 'Last30Days', 'Last60Days'] and default value 'Last60Days'; properties/term with allowed values ['P1Y', 'P3Y'] and default
	// value 'P3Y'; properties/subscriptionId; properties/resourceGroup
	Filter *string

	// May be used to order the recommendations by: properties/armSkuName. For the savings plan, the results are in order by default.
	// There is no need to use this clause.
	Orderby *string
}

// BenefitUtilizationSummariesClientListByBillingAccountIDOptions contains the optional parameters for the BenefitUtilizationSummariesClient.NewListByBillingAccountIDPager
// method.
type BenefitUtilizationSummariesClientListByBillingAccountIDOptions struct {
	// Supports filtering by properties/benefitId, properties/benefitOrderId and properties/usageDate.
	Filter *string

	// Grain.
	GrainParameter *GrainParameter
}

// BenefitUtilizationSummariesClientListByBillingProfileIDOptions contains the optional parameters for the BenefitUtilizationSummariesClient.NewListByBillingProfileIDPager
// method.
type BenefitUtilizationSummariesClientListByBillingProfileIDOptions struct {
	// Supports filtering by properties/benefitId, properties/benefitOrderId and properties/usageDate.
	Filter *string

	// Grain.
	GrainParameter *GrainParameter
}

// BenefitUtilizationSummariesClientListBySavingsPlanIDOptions contains the optional parameters for the BenefitUtilizationSummariesClient.NewListBySavingsPlanIDPager
// method.
type BenefitUtilizationSummariesClientListBySavingsPlanIDOptions struct {
	// Supports filtering by properties/usageDate.
	Filter *string

	// Grain.
	GrainParameter *GrainParameter
}

// BenefitUtilizationSummariesClientListBySavingsPlanOrderOptions contains the optional parameters for the BenefitUtilizationSummariesClient.NewListBySavingsPlanOrderPager
// method.
type BenefitUtilizationSummariesClientListBySavingsPlanOrderOptions struct {
	// Supports filtering by properties/usageDate.
	Filter *string

	// Grain.
	GrainParameter *GrainParameter
}

// DimensionsClientByExternalCloudProviderTypeOptions contains the optional parameters for the DimensionsClient.NewByExternalCloudProviderTypePager
// method.
type DimensionsClientByExternalCloudProviderTypeOptions struct {
	// May be used to expand the properties/data within a dimension category. By default, data is not included when listing dimensions.
	Expand *string

	// May be used to filter dimensions by properties/category, properties/usageStart, properties/usageEnd. Supported operators
	// are 'eq','lt', 'gt', 'le', 'ge'.
	Filter *string

	// Skiptoken is only used if a previous operation returned a partial result. If a previous response contains a nextLink element,
	// the value of the nextLink element will include a skiptoken parameter that
	// specifies a starting point to use for subsequent calls.
	Skiptoken *string

	// May be used to limit the number of results to the most recent N dimension data.
	Top *int32
}

// DimensionsClientListOptions contains the optional parameters for the DimensionsClient.NewListPager method.
type DimensionsClientListOptions struct {
	// May be used to expand the properties/data within a dimension category. By default, data is not included when listing dimensions.
	Expand *string

	// May be used to filter dimensions by properties/category, properties/usageStart, properties/usageEnd. Supported operators
	// are 'eq','lt', 'gt', 'le', 'ge'.
	Filter *string

	// Skiptoken is only used if a previous operation returned a partial result. If a previous response contains a nextLink element,
	// the value of the nextLink element will include a skiptoken parameter that
	// specifies a starting point to use for subsequent calls.
	Skiptoken *string

	// May be used to limit the number of results to the most recent N dimension data.
	Top *int32
}

// ExportsClientCreateOrUpdateOptions contains the optional parameters for the ExportsClient.CreateOrUpdate method.
type ExportsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ExportsClientDeleteOptions contains the optional parameters for the ExportsClient.Delete method.
type ExportsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ExportsClientExecuteOptions contains the optional parameters for the ExportsClient.Execute method.
type ExportsClientExecuteOptions struct {
	// placeholder for future optional parameters
}

// ExportsClientGetExecutionHistoryOptions contains the optional parameters for the ExportsClient.GetExecutionHistory method.
type ExportsClientGetExecutionHistoryOptions struct {
	// placeholder for future optional parameters
}

// ExportsClientGetOptions contains the optional parameters for the ExportsClient.Get method.
type ExportsClientGetOptions struct {
	// May be used to expand the properties within an export. Currently only 'runHistory' is supported and will return information
	// for the last 10 runs of the export.
	Expand *string
}

// ExportsClientListOptions contains the optional parameters for the ExportsClient.List method.
type ExportsClientListOptions struct {
	// May be used to expand the properties within an export. Currently only 'runHistory' is supported and will return information
	// for the last run of each export.
	Expand *string
}

// ForecastClientExternalCloudProviderUsageOptions contains the optional parameters for the ForecastClient.ExternalCloudProviderUsage
// method.
type ForecastClientExternalCloudProviderUsageOptions struct {
	// May be used to filter forecasts by properties/usageDate (Utc time), properties/chargeType or properties/grain. The filter
	// supports 'eq', 'lt', 'gt', 'le', 'ge', and 'and'. It does not currently
	// support 'ne', 'or', or 'not'.
	Filter *string
}

// ForecastClientUsageOptions contains the optional parameters for the ForecastClient.Usage method.
type ForecastClientUsageOptions struct {
	// May be used to filter forecasts by properties/usageDate (Utc time), properties/chargeType or properties/grain. The filter
	// supports 'eq', 'lt', 'gt', 'le', 'ge', and 'and'. It does not currently
	// support 'ne', 'or', or 'not'.
	Filter *string
}

// GenerateCostDetailsReportClientBeginCreateOperationOptions contains the optional parameters for the GenerateCostDetailsReportClient.BeginCreateOperation
// method.
type GenerateCostDetailsReportClientBeginCreateOperationOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// GenerateCostDetailsReportClientBeginGetOperationResultsOptions contains the optional parameters for the GenerateCostDetailsReportClient.BeginGetOperationResults
// method.
type GenerateCostDetailsReportClientBeginGetOperationResultsOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// GenerateDetailedCostReportClientBeginCreateOperationOptions contains the optional parameters for the GenerateDetailedCostReportClient.BeginCreateOperation
// method.
type GenerateDetailedCostReportClientBeginCreateOperationOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// GenerateDetailedCostReportOperationResultsClientBeginGetOptions contains the optional parameters for the GenerateDetailedCostReportOperationResultsClient.BeginGet
// method.
type GenerateDetailedCostReportOperationResultsClientBeginGetOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// GenerateDetailedCostReportOperationStatusClientGetOptions contains the optional parameters for the GenerateDetailedCostReportOperationStatusClient.Get
// method.
type GenerateDetailedCostReportOperationStatusClientGetOptions struct {
	// placeholder for future optional parameters
}

// GenerateReservationDetailsReportClientBeginByBillingAccountIDOptions contains the optional parameters for the GenerateReservationDetailsReportClient.BeginByBillingAccountID
// method.
type GenerateReservationDetailsReportClientBeginByBillingAccountIDOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// GenerateReservationDetailsReportClientBeginByBillingProfileIDOptions contains the optional parameters for the GenerateReservationDetailsReportClient.BeginByBillingProfileID
// method.
type GenerateReservationDetailsReportClientBeginByBillingProfileIDOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// PriceSheetClientBeginDownloadByBillingProfileOptions contains the optional parameters for the PriceSheetClient.BeginDownloadByBillingProfile
// method.
type PriceSheetClientBeginDownloadByBillingProfileOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PriceSheetClientBeginDownloadOptions contains the optional parameters for the PriceSheetClient.BeginDownload method.
type PriceSheetClientBeginDownloadOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// QueryClientUsageByExternalCloudProviderTypeOptions contains the optional parameters for the QueryClient.UsageByExternalCloudProviderType
// method.
type QueryClientUsageByExternalCloudProviderTypeOptions struct {
	// placeholder for future optional parameters
}

// QueryClientUsageOptions contains the optional parameters for the QueryClient.Usage method.
type QueryClientUsageOptions struct {
	// placeholder for future optional parameters
}

// ScheduledActionsClientCheckNameAvailabilityByScopeOptions contains the optional parameters for the ScheduledActionsClient.CheckNameAvailabilityByScope
// method.
type ScheduledActionsClientCheckNameAvailabilityByScopeOptions struct {
	// placeholder for future optional parameters
}

// ScheduledActionsClientCheckNameAvailabilityOptions contains the optional parameters for the ScheduledActionsClient.CheckNameAvailability
// method.
type ScheduledActionsClientCheckNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// ScheduledActionsClientCreateOrUpdateByScopeOptions contains the optional parameters for the ScheduledActionsClient.CreateOrUpdateByScope
// method.
type ScheduledActionsClientCreateOrUpdateByScopeOptions struct {
	// ETag of the Entity. Not required when creating an entity. Optional when updating an entity and can be specified to achieve
	// optimistic concurrency.
	IfMatch *string
}

// ScheduledActionsClientCreateOrUpdateOptions contains the optional parameters for the ScheduledActionsClient.CreateOrUpdate
// method.
type ScheduledActionsClientCreateOrUpdateOptions struct {
	// ETag of the Entity. Not required when creating an entity. Optional when updating an entity and can be specified to achieve
	// optimistic concurrency.
	IfMatch *string
}

// ScheduledActionsClientDeleteByScopeOptions contains the optional parameters for the ScheduledActionsClient.DeleteByScope
// method.
type ScheduledActionsClientDeleteByScopeOptions struct {
	// placeholder for future optional parameters
}

// ScheduledActionsClientDeleteOptions contains the optional parameters for the ScheduledActionsClient.Delete method.
type ScheduledActionsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ScheduledActionsClientGetByScopeOptions contains the optional parameters for the ScheduledActionsClient.GetByScope method.
type ScheduledActionsClientGetByScopeOptions struct {
	// placeholder for future optional parameters
}

// ScheduledActionsClientGetOptions contains the optional parameters for the ScheduledActionsClient.Get method.
type ScheduledActionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ScheduledActionsClientListByScopeOptions contains the optional parameters for the ScheduledActionsClient.NewListByScopePager
// method.
type ScheduledActionsClientListByScopeOptions struct {
	// May be used to filter scheduled actions by properties/viewId. Supported operator is 'eq'.
	Filter *string
}

// ScheduledActionsClientListOptions contains the optional parameters for the ScheduledActionsClient.NewListPager method.
type ScheduledActionsClientListOptions struct {
	// May be used to filter scheduled actions by properties/viewId. Supported operator is 'eq'.
	Filter *string
}

// ScheduledActionsClientRunByScopeOptions contains the optional parameters for the ScheduledActionsClient.RunByScope method.
type ScheduledActionsClientRunByScopeOptions struct {
	// placeholder for future optional parameters
}

// ScheduledActionsClientRunOptions contains the optional parameters for the ScheduledActionsClient.Run method.
type ScheduledActionsClientRunOptions struct {
	// placeholder for future optional parameters
}

// ViewsClientCreateOrUpdateByScopeOptions contains the optional parameters for the ViewsClient.CreateOrUpdateByScope method.
type ViewsClientCreateOrUpdateByScopeOptions struct {
	// placeholder for future optional parameters
}

// ViewsClientCreateOrUpdateOptions contains the optional parameters for the ViewsClient.CreateOrUpdate method.
type ViewsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ViewsClientDeleteByScopeOptions contains the optional parameters for the ViewsClient.DeleteByScope method.
type ViewsClientDeleteByScopeOptions struct {
	// placeholder for future optional parameters
}

// ViewsClientDeleteOptions contains the optional parameters for the ViewsClient.Delete method.
type ViewsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ViewsClientGetByScopeOptions contains the optional parameters for the ViewsClient.GetByScope method.
type ViewsClientGetByScopeOptions struct {
	// placeholder for future optional parameters
}

// ViewsClientGetOptions contains the optional parameters for the ViewsClient.Get method.
type ViewsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ViewsClientListByScopeOptions contains the optional parameters for the ViewsClient.NewListByScopePager method.
type ViewsClientListByScopeOptions struct {
	// placeholder for future optional parameters
}

// ViewsClientListOptions contains the optional parameters for the ViewsClient.NewListPager method.
type ViewsClientListOptions struct {
	// placeholder for future optional parameters
}
