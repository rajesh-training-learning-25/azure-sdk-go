package securityinsightapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/securityinsight/mgmt/2022-01-01-preview/securityinsight"
	"github.com/Azure/go-autorest/autorest"
)

// AlertRulesClientAPI contains the set of methods on the AlertRulesClient type.
type AlertRulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, ruleID string, alertRule securityinsight.BasicAlertRule) (result securityinsight.AlertRuleModel, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, ruleID string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, ruleID string) (result securityinsight.AlertRuleModel, err error)
	List(ctx context.Context, resourceGroupName string, workspaceName string) (result securityinsight.AlertRulesListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, workspaceName string) (result securityinsight.AlertRulesListIterator, err error)
}

var _ AlertRulesClientAPI = (*securityinsight.AlertRulesClient)(nil)

// ActionsClientAPI contains the set of methods on the ActionsClient type.
type ActionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, ruleID string, actionID string, action securityinsight.ActionRequest) (result securityinsight.ActionResponse, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, ruleID string, actionID string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, ruleID string, actionID string) (result securityinsight.ActionResponse, err error)
	ListByAlertRule(ctx context.Context, resourceGroupName string, workspaceName string, ruleID string) (result securityinsight.ActionsListPage, err error)
	ListByAlertRuleComplete(ctx context.Context, resourceGroupName string, workspaceName string, ruleID string) (result securityinsight.ActionsListIterator, err error)
}

var _ ActionsClientAPI = (*securityinsight.ActionsClient)(nil)

// AlertRuleTemplatesClientAPI contains the set of methods on the AlertRuleTemplatesClient type.
type AlertRuleTemplatesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, workspaceName string, alertRuleTemplateID string) (result securityinsight.AlertRuleTemplateModel, err error)
	List(ctx context.Context, resourceGroupName string, workspaceName string) (result securityinsight.AlertRuleTemplatesListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, workspaceName string) (result securityinsight.AlertRuleTemplatesListIterator, err error)
}

var _ AlertRuleTemplatesClientAPI = (*securityinsight.AlertRuleTemplatesClient)(nil)

// AutomationRulesClientAPI contains the set of methods on the AutomationRulesClient type.
type AutomationRulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, automationRuleID string, automationRuleToUpsert *securityinsight.AutomationRule) (result securityinsight.AutomationRule, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, automationRuleID string) (result securityinsight.SetObject, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, automationRuleID string) (result securityinsight.AutomationRule, err error)
	List(ctx context.Context, resourceGroupName string, workspaceName string) (result securityinsight.AutomationRulesListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, workspaceName string) (result securityinsight.AutomationRulesListIterator, err error)
}

var _ AutomationRulesClientAPI = (*securityinsight.AutomationRulesClient)(nil)

// IncidentsClientAPI contains the set of methods on the IncidentsClient type.
type IncidentsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, incident securityinsight.Incident) (result securityinsight.Incident, err error)
	CreateTeam(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, teamProperties securityinsight.TeamProperties) (result securityinsight.TeamInformation, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string) (result securityinsight.Incident, err error)
	List(ctx context.Context, resourceGroupName string, workspaceName string, filter string, orderby string, top *int32, skipToken string) (result securityinsight.IncidentListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, workspaceName string, filter string, orderby string, top *int32, skipToken string) (result securityinsight.IncidentListIterator, err error)
	ListAlerts(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string) (result securityinsight.IncidentAlertList, err error)
	ListBookmarks(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string) (result securityinsight.IncidentBookmarkList, err error)
	ListEntities(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string) (result securityinsight.IncidentEntitiesResponse, err error)
	RunPlaybook(ctx context.Context, resourceGroupName string, workspaceName string, incidentIdentifier string, requestBody *securityinsight.ManualTriggerRequestBody) (result securityinsight.SetObject, err error)
}

var _ IncidentsClientAPI = (*securityinsight.IncidentsClient)(nil)

// BookmarksClientAPI contains the set of methods on the BookmarksClient type.
type BookmarksClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, bookmarkID string, bookmark securityinsight.Bookmark) (result securityinsight.Bookmark, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, bookmarkID string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, bookmarkID string) (result securityinsight.Bookmark, err error)
	List(ctx context.Context, resourceGroupName string, workspaceName string) (result securityinsight.BookmarkListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, workspaceName string) (result securityinsight.BookmarkListIterator, err error)
}

var _ BookmarksClientAPI = (*securityinsight.BookmarksClient)(nil)

// BookmarkRelationsClientAPI contains the set of methods on the BookmarkRelationsClient type.
type BookmarkRelationsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, bookmarkID string, relationName string, relation securityinsight.Relation) (result securityinsight.Relation, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, bookmarkID string, relationName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, bookmarkID string, relationName string) (result securityinsight.Relation, err error)
	List(ctx context.Context, resourceGroupName string, workspaceName string, bookmarkID string, filter string, orderby string, top *int32, skipToken string) (result securityinsight.RelationListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, workspaceName string, bookmarkID string, filter string, orderby string, top *int32, skipToken string) (result securityinsight.RelationListIterator, err error)
}

var _ BookmarkRelationsClientAPI = (*securityinsight.BookmarkRelationsClient)(nil)

// BookmarkClientAPI contains the set of methods on the BookmarkClient type.
type BookmarkClientAPI interface {
	Expand(ctx context.Context, resourceGroupName string, workspaceName string, bookmarkID string, parameters securityinsight.BookmarkExpandParameters) (result securityinsight.BookmarkExpandResponse, err error)
}

var _ BookmarkClientAPI = (*securityinsight.BookmarkClient)(nil)

// IPGeodataClientAPI contains the set of methods on the IPGeodataClient type.
type IPGeodataClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, IPAddress string) (result securityinsight.EnrichmentIPGeodata, err error)
}

var _ IPGeodataClientAPI = (*securityinsight.IPGeodataClient)(nil)

// DomainWhoisClientAPI contains the set of methods on the DomainWhoisClient type.
type DomainWhoisClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, domain string) (result securityinsight.EnrichmentDomainWhois, err error)
}

var _ DomainWhoisClientAPI = (*securityinsight.DomainWhoisClient)(nil)

// EntitiesClientAPI contains the set of methods on the EntitiesClient type.
type EntitiesClientAPI interface {
	Expand(ctx context.Context, resourceGroupName string, workspaceName string, entityID string, parameters securityinsight.EntityExpandParameters) (result securityinsight.EntityExpandResponse, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, entityID string) (result securityinsight.EntityModel, err error)
	GetInsights(ctx context.Context, resourceGroupName string, workspaceName string, entityID string, parameters securityinsight.EntityGetInsightsParameters) (result securityinsight.EntityGetInsightsResponse, err error)
	List(ctx context.Context, resourceGroupName string, workspaceName string) (result securityinsight.EntityListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, workspaceName string) (result securityinsight.EntityListIterator, err error)
	Queries(ctx context.Context, resourceGroupName string, workspaceName string, entityID string) (result securityinsight.GetQueriesResponse, err error)
}

var _ EntitiesClientAPI = (*securityinsight.EntitiesClient)(nil)

// EntitiesGetTimelineClientAPI contains the set of methods on the EntitiesGetTimelineClient type.
type EntitiesGetTimelineClientAPI interface {
	List(ctx context.Context, resourceGroupName string, workspaceName string, entityID string, parameters securityinsight.EntityTimelineParameters) (result securityinsight.EntityTimelineResponse, err error)
}

var _ EntitiesGetTimelineClientAPI = (*securityinsight.EntitiesGetTimelineClient)(nil)

// EntitiesRelationsClientAPI contains the set of methods on the EntitiesRelationsClient type.
type EntitiesRelationsClientAPI interface {
	List(ctx context.Context, resourceGroupName string, workspaceName string, entityID string, filter string, orderby string, top *int32, skipToken string) (result securityinsight.RelationListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, workspaceName string, entityID string, filter string, orderby string, top *int32, skipToken string) (result securityinsight.RelationListIterator, err error)
}

var _ EntitiesRelationsClientAPI = (*securityinsight.EntitiesRelationsClient)(nil)

// EntityRelationsClientAPI contains the set of methods on the EntityRelationsClient type.
type EntityRelationsClientAPI interface {
	GetRelation(ctx context.Context, resourceGroupName string, workspaceName string, entityID string, relationName string) (result securityinsight.Relation, err error)
}

var _ EntityRelationsClientAPI = (*securityinsight.EntityRelationsClient)(nil)

// EntityQueriesClientAPI contains the set of methods on the EntityQueriesClient type.
type EntityQueriesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, entityQueryID string, entityQuery securityinsight.BasicCustomEntityQuery) (result securityinsight.EntityQueryModel, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, entityQueryID string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, entityQueryID string) (result securityinsight.EntityQueryModel, err error)
	List(ctx context.Context, resourceGroupName string, workspaceName string, kind string) (result securityinsight.EntityQueryListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, workspaceName string, kind string) (result securityinsight.EntityQueryListIterator, err error)
}

var _ EntityQueriesClientAPI = (*securityinsight.EntityQueriesClient)(nil)

// EntityQueryTemplatesClientAPI contains the set of methods on the EntityQueryTemplatesClient type.
type EntityQueryTemplatesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, workspaceName string, entityQueryTemplateID string) (result securityinsight.EntityQueryTemplateModel, err error)
	List(ctx context.Context, resourceGroupName string, workspaceName string, kind string) (result securityinsight.EntityQueryTemplateListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, workspaceName string, kind string) (result securityinsight.EntityQueryTemplateListIterator, err error)
}

var _ EntityQueryTemplatesClientAPI = (*securityinsight.EntityQueryTemplatesClient)(nil)

// IncidentCommentsClientAPI contains the set of methods on the IncidentCommentsClient type.
type IncidentCommentsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, incidentCommentID string, incidentComment securityinsight.IncidentComment) (result securityinsight.IncidentComment, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, incidentCommentID string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, incidentCommentID string) (result securityinsight.IncidentComment, err error)
	List(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, filter string, orderby string, top *int32, skipToken string) (result securityinsight.IncidentCommentListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, filter string, orderby string, top *int32, skipToken string) (result securityinsight.IncidentCommentListIterator, err error)
}

var _ IncidentCommentsClientAPI = (*securityinsight.IncidentCommentsClient)(nil)

// IncidentRelationsClientAPI contains the set of methods on the IncidentRelationsClient type.
type IncidentRelationsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, relationName string, relation securityinsight.Relation) (result securityinsight.Relation, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, relationName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, relationName string) (result securityinsight.Relation, err error)
	List(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, filter string, orderby string, top *int32, skipToken string) (result securityinsight.RelationListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, filter string, orderby string, top *int32, skipToken string) (result securityinsight.RelationListIterator, err error)
}

var _ IncidentRelationsClientAPI = (*securityinsight.IncidentRelationsClient)(nil)

// MetadataClientAPI contains the set of methods on the MetadataClient type.
type MetadataClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, workspaceName string, metadataName string, metadata securityinsight.MetadataModel) (result securityinsight.MetadataModel, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, metadataName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, metadataName string) (result securityinsight.MetadataModel, err error)
	List(ctx context.Context, resourceGroupName string, workspaceName string, filter string, orderby string, top *int32, skip *int32) (result securityinsight.MetadataListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, workspaceName string, filter string, orderby string, top *int32, skip *int32) (result securityinsight.MetadataListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, workspaceName string, metadataName string, metadataPatch securityinsight.MetadataPatch) (result securityinsight.MetadataModel, err error)
}

var _ MetadataClientAPI = (*securityinsight.MetadataClient)(nil)

// OfficeConsentsClientAPI contains the set of methods on the OfficeConsentsClient type.
type OfficeConsentsClientAPI interface {
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, consentID string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, consentID string) (result securityinsight.OfficeConsent, err error)
	List(ctx context.Context, resourceGroupName string, workspaceName string) (result securityinsight.OfficeConsentListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, workspaceName string) (result securityinsight.OfficeConsentListIterator, err error)
}

var _ OfficeConsentsClientAPI = (*securityinsight.OfficeConsentsClient)(nil)

// SentinelOnboardingStatesClientAPI contains the set of methods on the SentinelOnboardingStatesClient type.
type SentinelOnboardingStatesClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, workspaceName string, sentinelOnboardingStateName string, sentinelOnboardingStateParameter *securityinsight.SentinelOnboardingState) (result securityinsight.SentinelOnboardingState, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, sentinelOnboardingStateName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, sentinelOnboardingStateName string) (result securityinsight.SentinelOnboardingState, err error)
	List(ctx context.Context, resourceGroupName string, workspaceName string) (result securityinsight.SentinelOnboardingStatesList, err error)
}

var _ SentinelOnboardingStatesClientAPI = (*securityinsight.SentinelOnboardingStatesClient)(nil)

// ProductSettingsClientAPI contains the set of methods on the ProductSettingsClient type.
type ProductSettingsClientAPI interface {
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, settingsName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, settingsName string) (result securityinsight.SettingsModel, err error)
	List(ctx context.Context, resourceGroupName string, workspaceName string) (result securityinsight.SettingList, err error)
	Update(ctx context.Context, resourceGroupName string, workspaceName string, settingsName string, settings securityinsight.BasicSettings) (result securityinsight.SettingsModel, err error)
}

var _ ProductSettingsClientAPI = (*securityinsight.ProductSettingsClient)(nil)

// SourceControlClientAPI contains the set of methods on the SourceControlClient type.
type SourceControlClientAPI interface {
	ListRepositories(ctx context.Context, resourceGroupName string, workspaceName string, repoType securityinsight.RepoType) (result securityinsight.RepoListPage, err error)
	ListRepositoriesComplete(ctx context.Context, resourceGroupName string, workspaceName string, repoType securityinsight.RepoType) (result securityinsight.RepoListIterator, err error)
}

var _ SourceControlClientAPI = (*securityinsight.SourceControlClient)(nil)

// SourceControlsClientAPI contains the set of methods on the SourceControlsClient type.
type SourceControlsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, workspaceName string, sourceControlID string, sourceControl securityinsight.SourceControl) (result securityinsight.SourceControl, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, sourceControlID string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, sourceControlID string) (result securityinsight.SourceControl, err error)
	List(ctx context.Context, resourceGroupName string, workspaceName string) (result securityinsight.SourceControlListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, workspaceName string) (result securityinsight.SourceControlListIterator, err error)
}

var _ SourceControlsClientAPI = (*securityinsight.SourceControlsClient)(nil)

// ThreatIntelligenceIndicatorClientAPI contains the set of methods on the ThreatIntelligenceIndicatorClient type.
type ThreatIntelligenceIndicatorClientAPI interface {
	AppendTags(ctx context.Context, resourceGroupName string, workspaceName string, name string, threatIntelligenceAppendTags securityinsight.ThreatIntelligenceAppendTags) (result autorest.Response, err error)
	Create(ctx context.Context, resourceGroupName string, workspaceName string, name string, threatIntelligenceProperties securityinsight.ThreatIntelligenceIndicatorModel) (result securityinsight.ThreatIntelligenceInformationModel, err error)
	CreateIndicator(ctx context.Context, resourceGroupName string, workspaceName string, threatIntelligenceProperties securityinsight.ThreatIntelligenceIndicatorModel) (result securityinsight.ThreatIntelligenceInformationModel, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, name string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, name string) (result securityinsight.ThreatIntelligenceInformationModel, err error)
	QueryIndicators(ctx context.Context, resourceGroupName string, workspaceName string, threatIntelligenceFilteringCriteria securityinsight.ThreatIntelligenceFilteringCriteria) (result securityinsight.ThreatIntelligenceInformationListPage, err error)
	QueryIndicatorsComplete(ctx context.Context, resourceGroupName string, workspaceName string, threatIntelligenceFilteringCriteria securityinsight.ThreatIntelligenceFilteringCriteria) (result securityinsight.ThreatIntelligenceInformationListIterator, err error)
	ReplaceTags(ctx context.Context, resourceGroupName string, workspaceName string, name string, threatIntelligenceReplaceTags securityinsight.ThreatIntelligenceIndicatorModel) (result securityinsight.ThreatIntelligenceInformationModel, err error)
}

var _ ThreatIntelligenceIndicatorClientAPI = (*securityinsight.ThreatIntelligenceIndicatorClient)(nil)

// ThreatIntelligenceIndicatorsClientAPI contains the set of methods on the ThreatIntelligenceIndicatorsClient type.
type ThreatIntelligenceIndicatorsClientAPI interface {
	List(ctx context.Context, resourceGroupName string, workspaceName string, filter string, orderby string, top *int32, skipToken string) (result securityinsight.ThreatIntelligenceInformationListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, workspaceName string, filter string, orderby string, top *int32, skipToken string) (result securityinsight.ThreatIntelligenceInformationListIterator, err error)
}

var _ ThreatIntelligenceIndicatorsClientAPI = (*securityinsight.ThreatIntelligenceIndicatorsClient)(nil)

// ThreatIntelligenceIndicatorMetricsClientAPI contains the set of methods on the ThreatIntelligenceIndicatorMetricsClient type.
type ThreatIntelligenceIndicatorMetricsClientAPI interface {
	List(ctx context.Context, resourceGroupName string, workspaceName string) (result securityinsight.ThreatIntelligenceMetricsList, err error)
}

var _ ThreatIntelligenceIndicatorMetricsClientAPI = (*securityinsight.ThreatIntelligenceIndicatorMetricsClient)(nil)

// WatchlistsClientAPI contains the set of methods on the WatchlistsClient type.
type WatchlistsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, watchlistAlias string, watchlist securityinsight.Watchlist) (result securityinsight.Watchlist, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, watchlistAlias string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, watchlistAlias string) (result securityinsight.Watchlist, err error)
	List(ctx context.Context, resourceGroupName string, workspaceName string, skipToken string) (result securityinsight.WatchlistListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, workspaceName string, skipToken string) (result securityinsight.WatchlistListIterator, err error)
}

var _ WatchlistsClientAPI = (*securityinsight.WatchlistsClient)(nil)

// WatchlistItemsClientAPI contains the set of methods on the WatchlistItemsClient type.
type WatchlistItemsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, watchlistAlias string, watchlistItemID string, watchlistItem securityinsight.WatchlistItem) (result securityinsight.WatchlistItem, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, watchlistAlias string, watchlistItemID string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, watchlistAlias string, watchlistItemID string) (result securityinsight.WatchlistItem, err error)
	List(ctx context.Context, resourceGroupName string, workspaceName string, watchlistAlias string, skipToken string) (result securityinsight.WatchlistItemListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, workspaceName string, watchlistAlias string, skipToken string) (result securityinsight.WatchlistItemListIterator, err error)
}

var _ WatchlistItemsClientAPI = (*securityinsight.WatchlistItemsClient)(nil)

// DataConnectorsClientAPI contains the set of methods on the DataConnectorsClient type.
type DataConnectorsClientAPI interface {
	Connect(ctx context.Context, resourceGroupName string, workspaceName string, dataConnectorID string, connectBody securityinsight.DataConnectorConnectBody) (result autorest.Response, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, dataConnectorID string, dataConnector securityinsight.BasicDataConnector) (result securityinsight.DataConnectorModel, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, dataConnectorID string) (result autorest.Response, err error)
	Disconnect(ctx context.Context, resourceGroupName string, workspaceName string, dataConnectorID string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, dataConnectorID string) (result securityinsight.DataConnectorModel, err error)
	List(ctx context.Context, resourceGroupName string, workspaceName string) (result securityinsight.DataConnectorListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, workspaceName string) (result securityinsight.DataConnectorListIterator, err error)
}

var _ DataConnectorsClientAPI = (*securityinsight.DataConnectorsClient)(nil)

// DataConnectorsCheckRequirementsClientAPI contains the set of methods on the DataConnectorsCheckRequirementsClient type.
type DataConnectorsCheckRequirementsClientAPI interface {
	Post(ctx context.Context, resourceGroupName string, workspaceName string, dataConnectorsCheckRequirements securityinsight.BasicDataConnectorsCheckRequirements) (result securityinsight.DataConnectorRequirementsState, err error)
}

var _ DataConnectorsCheckRequirementsClientAPI = (*securityinsight.DataConnectorsCheckRequirementsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result securityinsight.OperationsListPage, err error)
	ListComplete(ctx context.Context) (result securityinsight.OperationsListIterator, err error)
}

var _ OperationsClientAPI = (*securityinsight.OperationsClient)(nil)
