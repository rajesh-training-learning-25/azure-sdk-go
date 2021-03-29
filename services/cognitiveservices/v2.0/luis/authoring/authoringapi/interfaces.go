package authoringapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v2.0/luis/authoring"
	"github.com/gofrs/uuid"
)

// FeaturesClientAPI contains the set of methods on the FeaturesClient type.
type FeaturesClientAPI interface {
	AddPhraseList(ctx context.Context, appID uuid.UUID, versionID string, phraselistCreateObject authoring.PhraselistCreateObject) (result authoring.Int32, err error)
	DeletePhraseList(ctx context.Context, appID uuid.UUID, versionID string, phraselistID int32) (result authoring.OperationStatus, err error)
	GetPhraseList(ctx context.Context, appID uuid.UUID, versionID string, phraselistID int32) (result authoring.PhraseListFeatureInfo, err error)
	List(ctx context.Context, appID uuid.UUID, versionID string, skip *int32, take *int32) (result authoring.FeaturesResponseObject, err error)
	ListApplicationVersionPatternFeatures(ctx context.Context, appID uuid.UUID, versionID string, skip *int32, take *int32) (result authoring.ListPatternFeatureInfo, err error)
	ListPhraseLists(ctx context.Context, appID uuid.UUID, versionID string, skip *int32, take *int32) (result authoring.ListPhraseListFeatureInfo, err error)
	UpdatePhraseList(ctx context.Context, appID uuid.UUID, versionID string, phraselistID int32, phraselistUpdateObject *authoring.PhraselistUpdateObject) (result authoring.OperationStatus, err error)
}

var _ FeaturesClientAPI = (*authoring.FeaturesClient)(nil)

// ExamplesClientAPI contains the set of methods on the ExamplesClient type.
type ExamplesClientAPI interface {
	Add(ctx context.Context, appID uuid.UUID, versionID string, exampleLabelObject authoring.ExampleLabelObject) (result authoring.LabelExampleResponse, err error)
	Batch(ctx context.Context, appID uuid.UUID, versionID string, exampleLabelObjectArray []authoring.ExampleLabelObject) (result authoring.ListBatchLabelExample, err error)
	Delete(ctx context.Context, appID uuid.UUID, versionID string, exampleID int32) (result authoring.OperationStatus, err error)
	List(ctx context.Context, appID uuid.UUID, versionID string, skip *int32, take *int32) (result authoring.ListLabeledUtterance, err error)
}

var _ ExamplesClientAPI = (*authoring.ExamplesClient)(nil)

// ModelClientAPI contains the set of methods on the ModelClient type.
type ModelClientAPI interface {
	AddClosedList(ctx context.Context, appID uuid.UUID, versionID string, closedListModelCreateObject authoring.ClosedListModelCreateObject) (result authoring.UUID, err error)
	AddCompositeEntity(ctx context.Context, appID uuid.UUID, versionID string, compositeModelCreateObject authoring.CompositeEntityModel) (result authoring.UUID, err error)
	AddCompositeEntityChild(ctx context.Context, appID uuid.UUID, versionID string, cEntityID uuid.UUID, compositeChildModelCreateObject authoring.CompositeChildModelCreateObject) (result authoring.UUID, err error)
	AddCustomPrebuiltDomain(ctx context.Context, appID uuid.UUID, versionID string, prebuiltDomainObject authoring.PrebuiltDomainCreateBaseObject) (result authoring.ListUUID, err error)
	AddCustomPrebuiltEntity(ctx context.Context, appID uuid.UUID, versionID string, prebuiltDomainModelCreateObject authoring.PrebuiltDomainModelCreateObject) (result authoring.UUID, err error)
	AddCustomPrebuiltIntent(ctx context.Context, appID uuid.UUID, versionID string, prebuiltDomainModelCreateObject authoring.PrebuiltDomainModelCreateObject) (result authoring.UUID, err error)
	AddEntity(ctx context.Context, appID uuid.UUID, versionID string, modelCreateObject authoring.ModelCreateObject) (result authoring.UUID, err error)
	AddExplicitListItem(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID, item authoring.ExplicitListItemCreateObject) (result authoring.Int32, err error)
	AddHierarchicalEntity(ctx context.Context, appID uuid.UUID, versionID string, hierarchicalModelCreateObject authoring.HierarchicalEntityModel) (result authoring.UUID, err error)
	AddHierarchicalEntityChild(ctx context.Context, appID uuid.UUID, versionID string, hEntityID uuid.UUID, hierarchicalChildModelCreateObject authoring.HierarchicalChildModelCreateObject) (result authoring.UUID, err error)
	AddIntent(ctx context.Context, appID uuid.UUID, versionID string, intentCreateObject authoring.ModelCreateObject) (result authoring.UUID, err error)
	AddPrebuilt(ctx context.Context, appID uuid.UUID, versionID string, prebuiltExtractorNames []string) (result authoring.ListPrebuiltEntityExtractor, err error)
	AddSubList(ctx context.Context, appID uuid.UUID, versionID string, clEntityID uuid.UUID, wordListCreateObject authoring.WordListObject) (result authoring.Int64, err error)
	CreateClosedListEntityRole(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID, entityRoleCreateObject authoring.EntityRoleCreateObject) (result authoring.UUID, err error)
	CreateCompositeEntityRole(ctx context.Context, appID uuid.UUID, versionID string, cEntityID uuid.UUID, entityRoleCreateObject authoring.EntityRoleCreateObject) (result authoring.UUID, err error)
	CreateCustomPrebuiltEntityRole(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID, entityRoleCreateObject authoring.EntityRoleCreateObject) (result authoring.UUID, err error)
	CreateEntityRole(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID, entityRoleCreateObject authoring.EntityRoleCreateObject) (result authoring.UUID, err error)
	CreateHierarchicalEntityRole(ctx context.Context, appID uuid.UUID, versionID string, hEntityID uuid.UUID, entityRoleCreateObject authoring.EntityRoleCreateObject) (result authoring.UUID, err error)
	CreatePatternAnyEntityModel(ctx context.Context, appID uuid.UUID, versionID string, extractorCreateObject authoring.PatternAnyModelCreateObject) (result authoring.UUID, err error)
	CreatePatternAnyEntityRole(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID, entityRoleCreateObject authoring.EntityRoleCreateObject) (result authoring.UUID, err error)
	CreatePrebuiltEntityRole(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID, entityRoleCreateObject authoring.EntityRoleCreateObject) (result authoring.UUID, err error)
	CreateRegexEntityModel(ctx context.Context, appID uuid.UUID, versionID string, regexEntityExtractorCreateObj authoring.RegexModelCreateObject) (result authoring.UUID, err error)
	CreateRegexEntityRole(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID, entityRoleCreateObject authoring.EntityRoleCreateObject) (result authoring.UUID, err error)
	DeleteClosedList(ctx context.Context, appID uuid.UUID, versionID string, clEntityID uuid.UUID) (result authoring.OperationStatus, err error)
	DeleteClosedListEntityRole(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID, roleID uuid.UUID) (result authoring.OperationStatus, err error)
	DeleteCompositeEntity(ctx context.Context, appID uuid.UUID, versionID string, cEntityID uuid.UUID) (result authoring.OperationStatus, err error)
	DeleteCompositeEntityChild(ctx context.Context, appID uuid.UUID, versionID string, cEntityID uuid.UUID, cChildID uuid.UUID) (result authoring.OperationStatus, err error)
	DeleteCompositeEntityRole(ctx context.Context, appID uuid.UUID, versionID string, cEntityID uuid.UUID, roleID uuid.UUID) (result authoring.OperationStatus, err error)
	DeleteCustomEntityRole(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID, roleID uuid.UUID) (result authoring.OperationStatus, err error)
	DeleteCustomPrebuiltDomain(ctx context.Context, appID uuid.UUID, versionID string, domainName string) (result authoring.OperationStatus, err error)
	DeleteEntity(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID) (result authoring.OperationStatus, err error)
	DeleteEntityRole(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID, roleID uuid.UUID) (result authoring.OperationStatus, err error)
	DeleteExplicitListItem(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID, itemID int64) (result authoring.OperationStatus, err error)
	DeleteHierarchicalEntity(ctx context.Context, appID uuid.UUID, versionID string, hEntityID uuid.UUID) (result authoring.OperationStatus, err error)
	DeleteHierarchicalEntityChild(ctx context.Context, appID uuid.UUID, versionID string, hEntityID uuid.UUID, hChildID uuid.UUID) (result authoring.OperationStatus, err error)
	DeleteHierarchicalEntityRole(ctx context.Context, appID uuid.UUID, versionID string, hEntityID uuid.UUID, roleID uuid.UUID) (result authoring.OperationStatus, err error)
	DeleteIntent(ctx context.Context, appID uuid.UUID, versionID string, intentID uuid.UUID, deleteUtterances *bool) (result authoring.OperationStatus, err error)
	DeletePatternAnyEntityModel(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID) (result authoring.OperationStatus, err error)
	DeletePatternAnyEntityRole(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID, roleID uuid.UUID) (result authoring.OperationStatus, err error)
	DeletePrebuilt(ctx context.Context, appID uuid.UUID, versionID string, prebuiltID uuid.UUID) (result authoring.OperationStatus, err error)
	DeletePrebuiltEntityRole(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID, roleID uuid.UUID) (result authoring.OperationStatus, err error)
	DeleteRegexEntityModel(ctx context.Context, appID uuid.UUID, versionID string, regexEntityID uuid.UUID) (result authoring.OperationStatus, err error)
	DeleteRegexEntityRole(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID, roleID uuid.UUID) (result authoring.OperationStatus, err error)
	DeleteSubList(ctx context.Context, appID uuid.UUID, versionID string, clEntityID uuid.UUID, subListID int64) (result authoring.OperationStatus, err error)
	ExamplesMethod(ctx context.Context, appID uuid.UUID, versionID string, modelID string, skip *int32, take *int32) (result authoring.ListLabelTextObject, err error)
	GetClosedList(ctx context.Context, appID uuid.UUID, versionID string, clEntityID uuid.UUID) (result authoring.ClosedListEntityExtractor, err error)
	GetClosedListEntityRole(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID, roleID uuid.UUID) (result authoring.EntityRole, err error)
	GetCompositeEntity(ctx context.Context, appID uuid.UUID, versionID string, cEntityID uuid.UUID) (result authoring.CompositeEntityExtractor, err error)
	GetCompositeEntityRole(ctx context.Context, appID uuid.UUID, versionID string, cEntityID uuid.UUID, roleID uuid.UUID) (result authoring.EntityRole, err error)
	GetCustomEntityRole(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID, roleID uuid.UUID) (result authoring.EntityRole, err error)
	GetEntity(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID) (result authoring.EntityExtractor, err error)
	GetEntityRole(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID, roleID uuid.UUID) (result authoring.EntityRole, err error)
	GetExplicitList(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID) (result authoring.ListExplicitListItem, err error)
	GetExplicitListItem(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID, itemID int64) (result authoring.ExplicitListItem, err error)
	GetHierarchicalEntity(ctx context.Context, appID uuid.UUID, versionID string, hEntityID uuid.UUID) (result authoring.HierarchicalEntityExtractor, err error)
	GetHierarchicalEntityChild(ctx context.Context, appID uuid.UUID, versionID string, hEntityID uuid.UUID, hChildID uuid.UUID) (result authoring.HierarchicalChildEntity, err error)
	GetHierarchicalEntityRole(ctx context.Context, appID uuid.UUID, versionID string, hEntityID uuid.UUID, roleID uuid.UUID) (result authoring.EntityRole, err error)
	GetIntent(ctx context.Context, appID uuid.UUID, versionID string, intentID uuid.UUID) (result authoring.IntentClassifier, err error)
	GetPatternAnyEntityInfo(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID) (result authoring.PatternAnyEntityExtractor, err error)
	GetPatternAnyEntityRole(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID, roleID uuid.UUID) (result authoring.EntityRole, err error)
	GetPrebuilt(ctx context.Context, appID uuid.UUID, versionID string, prebuiltID uuid.UUID) (result authoring.PrebuiltEntityExtractor, err error)
	GetPrebuiltEntityRole(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID, roleID uuid.UUID) (result authoring.EntityRole, err error)
	GetRegexEntityEntityInfo(ctx context.Context, appID uuid.UUID, versionID string, regexEntityID uuid.UUID) (result authoring.RegexEntityExtractor, err error)
	GetRegexEntityRole(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID, roleID uuid.UUID) (result authoring.EntityRole, err error)
	ListClosedListEntityRoles(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID) (result authoring.ListEntityRole, err error)
	ListClosedLists(ctx context.Context, appID uuid.UUID, versionID string, skip *int32, take *int32) (result authoring.ListClosedListEntityExtractor, err error)
	ListCompositeEntities(ctx context.Context, appID uuid.UUID, versionID string, skip *int32, take *int32) (result authoring.ListCompositeEntityExtractor, err error)
	ListCompositeEntityRoles(ctx context.Context, appID uuid.UUID, versionID string, cEntityID uuid.UUID) (result authoring.ListEntityRole, err error)
	ListCustomPrebuiltEntities(ctx context.Context, appID uuid.UUID, versionID string) (result authoring.ListEntityExtractor, err error)
	ListCustomPrebuiltEntityRoles(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID) (result authoring.ListEntityRole, err error)
	ListCustomPrebuiltIntents(ctx context.Context, appID uuid.UUID, versionID string) (result authoring.ListIntentClassifier, err error)
	ListCustomPrebuiltModels(ctx context.Context, appID uuid.UUID, versionID string) (result authoring.ListCustomPrebuiltModel, err error)
	ListEntities(ctx context.Context, appID uuid.UUID, versionID string, skip *int32, take *int32) (result authoring.ListEntityExtractor, err error)
	ListEntityRoles(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID) (result authoring.ListEntityRole, err error)
	ListEntitySuggestions(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID, take *int32) (result authoring.ListEntitiesSuggestionExample, err error)
	ListHierarchicalEntities(ctx context.Context, appID uuid.UUID, versionID string, skip *int32, take *int32) (result authoring.ListHierarchicalEntityExtractor, err error)
	ListHierarchicalEntityRoles(ctx context.Context, appID uuid.UUID, versionID string, hEntityID uuid.UUID) (result authoring.ListEntityRole, err error)
	ListIntents(ctx context.Context, appID uuid.UUID, versionID string, skip *int32, take *int32) (result authoring.ListIntentClassifier, err error)
	ListIntentSuggestions(ctx context.Context, appID uuid.UUID, versionID string, intentID uuid.UUID, take *int32) (result authoring.ListIntentsSuggestionExample, err error)
	ListModels(ctx context.Context, appID uuid.UUID, versionID string, skip *int32, take *int32) (result authoring.ListModelInfoResponse, err error)
	ListPatternAnyEntityInfos(ctx context.Context, appID uuid.UUID, versionID string, skip *int32, take *int32) (result authoring.ListPatternAnyEntityExtractor, err error)
	ListPatternAnyEntityRoles(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID) (result authoring.ListEntityRole, err error)
	ListPrebuiltEntities(ctx context.Context, appID uuid.UUID, versionID string) (result authoring.ListAvailablePrebuiltEntityModel, err error)
	ListPrebuiltEntityRoles(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID) (result authoring.ListEntityRole, err error)
	ListPrebuilts(ctx context.Context, appID uuid.UUID, versionID string, skip *int32, take *int32) (result authoring.ListPrebuiltEntityExtractor, err error)
	ListRegexEntityInfos(ctx context.Context, appID uuid.UUID, versionID string, skip *int32, take *int32) (result authoring.ListRegexEntityExtractor, err error)
	ListRegexEntityRoles(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID) (result authoring.ListEntityRole, err error)
	PatchClosedList(ctx context.Context, appID uuid.UUID, versionID string, clEntityID uuid.UUID, closedListModelPatchObject authoring.ClosedListModelPatchObject) (result authoring.OperationStatus, err error)
	UpdateClosedList(ctx context.Context, appID uuid.UUID, versionID string, clEntityID uuid.UUID, closedListModelUpdateObject authoring.ClosedListModelUpdateObject) (result authoring.OperationStatus, err error)
	UpdateClosedListEntityRole(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID, roleID uuid.UUID, entityRoleUpdateObject authoring.EntityRoleUpdateObject) (result authoring.OperationStatus, err error)
	UpdateCompositeEntity(ctx context.Context, appID uuid.UUID, versionID string, cEntityID uuid.UUID, compositeModelUpdateObject authoring.CompositeEntityModel) (result authoring.OperationStatus, err error)
	UpdateCompositeEntityRole(ctx context.Context, appID uuid.UUID, versionID string, cEntityID uuid.UUID, roleID uuid.UUID, entityRoleUpdateObject authoring.EntityRoleUpdateObject) (result authoring.OperationStatus, err error)
	UpdateCustomPrebuiltEntityRole(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID, roleID uuid.UUID, entityRoleUpdateObject authoring.EntityRoleUpdateObject) (result authoring.OperationStatus, err error)
	UpdateEntity(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID, modelUpdateObject authoring.ModelUpdateObject) (result authoring.OperationStatus, err error)
	UpdateEntityRole(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID, roleID uuid.UUID, entityRoleUpdateObject authoring.EntityRoleUpdateObject) (result authoring.OperationStatus, err error)
	UpdateExplicitListItem(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID, itemID int64, item authoring.ExplicitListItemUpdateObject) (result authoring.OperationStatus, err error)
	UpdateHierarchicalEntity(ctx context.Context, appID uuid.UUID, versionID string, hEntityID uuid.UUID, hierarchicalModelUpdateObject authoring.HierarchicalEntityModel) (result authoring.OperationStatus, err error)
	UpdateHierarchicalEntityChild(ctx context.Context, appID uuid.UUID, versionID string, hEntityID uuid.UUID, hChildID uuid.UUID, hierarchicalChildModelUpdateObject authoring.HierarchicalChildModelUpdateObject) (result authoring.OperationStatus, err error)
	UpdateHierarchicalEntityRole(ctx context.Context, appID uuid.UUID, versionID string, hEntityID uuid.UUID, roleID uuid.UUID, entityRoleUpdateObject authoring.EntityRoleUpdateObject) (result authoring.OperationStatus, err error)
	UpdateIntent(ctx context.Context, appID uuid.UUID, versionID string, intentID uuid.UUID, modelUpdateObject authoring.ModelUpdateObject) (result authoring.OperationStatus, err error)
	UpdatePatternAnyEntityModel(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID, patternAnyUpdateObject authoring.PatternAnyModelUpdateObject) (result authoring.OperationStatus, err error)
	UpdatePatternAnyEntityRole(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID, roleID uuid.UUID, entityRoleUpdateObject authoring.EntityRoleUpdateObject) (result authoring.OperationStatus, err error)
	UpdatePrebuiltEntityRole(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID, roleID uuid.UUID, entityRoleUpdateObject authoring.EntityRoleUpdateObject) (result authoring.OperationStatus, err error)
	UpdateRegexEntityModel(ctx context.Context, appID uuid.UUID, versionID string, regexEntityID uuid.UUID, regexEntityUpdateObject authoring.RegexModelUpdateObject) (result authoring.OperationStatus, err error)
	UpdateRegexEntityRole(ctx context.Context, appID uuid.UUID, versionID string, entityID uuid.UUID, roleID uuid.UUID, entityRoleUpdateObject authoring.EntityRoleUpdateObject) (result authoring.OperationStatus, err error)
	UpdateSubList(ctx context.Context, appID uuid.UUID, versionID string, clEntityID uuid.UUID, subListID int64, wordListBaseUpdateObject authoring.WordListBaseUpdateObject) (result authoring.OperationStatus, err error)
}

var _ ModelClientAPI = (*authoring.ModelClient)(nil)

// AppsClientAPI contains the set of methods on the AppsClient type.
type AppsClientAPI interface {
	Add(ctx context.Context, applicationCreateObject authoring.ApplicationCreateObject) (result authoring.UUID, err error)
	AddCustomPrebuiltDomain(ctx context.Context, prebuiltDomainCreateObject authoring.PrebuiltDomainCreateObject) (result authoring.UUID, err error)
	Delete(ctx context.Context, appID uuid.UUID, force *bool) (result authoring.OperationStatus, err error)
	DownloadQueryLogs(ctx context.Context, appID uuid.UUID) (result authoring.ReadCloser, err error)
	Get(ctx context.Context, appID uuid.UUID) (result authoring.ApplicationInfoResponse, err error)
	GetPublishSettings(ctx context.Context, appID uuid.UUID) (result authoring.PublishSettings, err error)
	GetSettings(ctx context.Context, appID uuid.UUID) (result authoring.ApplicationSettings, err error)
	Import(ctx context.Context, luisApp authoring.LuisApp, appName string) (result authoring.UUID, err error)
	List(ctx context.Context, skip *int32, take *int32) (result authoring.ListApplicationInfoResponse, err error)
	ListAvailableCustomPrebuiltDomains(ctx context.Context) (result authoring.ListPrebuiltDomain, err error)
	ListAvailableCustomPrebuiltDomainsForCulture(ctx context.Context, culture string) (result authoring.ListPrebuiltDomain, err error)
	ListCortanaEndpoints(ctx context.Context) (result authoring.PersonalAssistantsResponse, err error)
	ListDomains(ctx context.Context) (result authoring.ListString, err error)
	ListEndpoints(ctx context.Context, appID uuid.UUID) (result authoring.SetString, err error)
	ListSupportedCultures(ctx context.Context) (result authoring.ListAvailableCulture, err error)
	ListUsageScenarios(ctx context.Context) (result authoring.ListString, err error)
	PackagePublishedApplicationAsGzip(ctx context.Context, appID uuid.UUID, slotName string) (result authoring.ReadCloser, err error)
	PackageTrainedApplicationAsGzip(ctx context.Context, appID uuid.UUID, versionID string) (result authoring.ReadCloser, err error)
	Publish(ctx context.Context, appID uuid.UUID, applicationPublishObject authoring.ApplicationPublishObject) (result authoring.ProductionOrStagingEndpointInfo, err error)
	Update(ctx context.Context, appID uuid.UUID, applicationUpdateObject authoring.ApplicationUpdateObject) (result authoring.OperationStatus, err error)
	UpdatePublishSettings(ctx context.Context, appID uuid.UUID, publishSettingUpdateObject authoring.PublishSettingUpdateObject) (result authoring.OperationStatus, err error)
	UpdateSettings(ctx context.Context, appID uuid.UUID, applicationSettingUpdateObject authoring.ApplicationSettingUpdateObject) (result authoring.OperationStatus, err error)
}

var _ AppsClientAPI = (*authoring.AppsClient)(nil)

// VersionsClientAPI contains the set of methods on the VersionsClient type.
type VersionsClientAPI interface {
	Clone(ctx context.Context, appID uuid.UUID, versionID string, versionCloneObject authoring.TaskUpdateObject) (result authoring.String, err error)
	Delete(ctx context.Context, appID uuid.UUID, versionID string) (result authoring.OperationStatus, err error)
	DeleteUnlabelledUtterance(ctx context.Context, appID uuid.UUID, versionID string, utterance string) (result authoring.OperationStatus, err error)
	Export(ctx context.Context, appID uuid.UUID, versionID string) (result authoring.LuisApp, err error)
	Get(ctx context.Context, appID uuid.UUID, versionID string) (result authoring.VersionInfo, err error)
	Import(ctx context.Context, appID uuid.UUID, luisApp authoring.LuisApp, versionID string) (result authoring.String, err error)
	List(ctx context.Context, appID uuid.UUID, skip *int32, take *int32) (result authoring.ListVersionInfo, err error)
	Update(ctx context.Context, appID uuid.UUID, versionID string, versionUpdateObject authoring.TaskUpdateObject) (result authoring.OperationStatus, err error)
}

var _ VersionsClientAPI = (*authoring.VersionsClient)(nil)

// TrainClientAPI contains the set of methods on the TrainClient type.
type TrainClientAPI interface {
	GetStatus(ctx context.Context, appID uuid.UUID, versionID string) (result authoring.ListModelTrainingInfo, err error)
	TrainVersion(ctx context.Context, appID uuid.UUID, versionID string) (result authoring.EnqueueTrainingResponse, err error)
}

var _ TrainClientAPI = (*authoring.TrainClient)(nil)

// PermissionsClientAPI contains the set of methods on the PermissionsClient type.
type PermissionsClientAPI interface {
	Add(ctx context.Context, appID uuid.UUID, userToAdd authoring.UserCollaborator) (result authoring.OperationStatus, err error)
	Delete(ctx context.Context, appID uuid.UUID, userToDelete authoring.UserCollaborator) (result authoring.OperationStatus, err error)
	List(ctx context.Context, appID uuid.UUID) (result authoring.UserAccessList, err error)
	Update(ctx context.Context, appID uuid.UUID, collaborators authoring.CollaboratorsArray) (result authoring.OperationStatus, err error)
}

var _ PermissionsClientAPI = (*authoring.PermissionsClient)(nil)

// PatternClientAPI contains the set of methods on the PatternClient type.
type PatternClientAPI interface {
	AddPattern(ctx context.Context, appID uuid.UUID, versionID string, pattern authoring.PatternRuleCreateObject) (result authoring.PatternRuleInfo, err error)
	BatchAddPatterns(ctx context.Context, appID uuid.UUID, versionID string, patterns []authoring.PatternRuleCreateObject) (result authoring.ListPatternRuleInfo, err error)
	DeletePattern(ctx context.Context, appID uuid.UUID, versionID string, patternID uuid.UUID) (result authoring.OperationStatus, err error)
	DeletePatterns(ctx context.Context, appID uuid.UUID, versionID string, patternIds []uuid.UUID) (result authoring.OperationStatus, err error)
	ListIntentPatterns(ctx context.Context, appID uuid.UUID, versionID string, intentID uuid.UUID, skip *int32, take *int32) (result authoring.ListPatternRuleInfo, err error)
	ListPatterns(ctx context.Context, appID uuid.UUID, versionID string, skip *int32, take *int32) (result authoring.ListPatternRuleInfo, err error)
	UpdatePattern(ctx context.Context, appID uuid.UUID, versionID string, patternID uuid.UUID, pattern authoring.PatternRuleUpdateObject) (result authoring.PatternRuleInfo, err error)
	UpdatePatterns(ctx context.Context, appID uuid.UUID, versionID string, patterns []authoring.PatternRuleUpdateObject) (result authoring.ListPatternRuleInfo, err error)
}

var _ PatternClientAPI = (*authoring.PatternClient)(nil)

// SettingsClientAPI contains the set of methods on the SettingsClient type.
type SettingsClientAPI interface {
	List(ctx context.Context, appID uuid.UUID, versionID string) (result authoring.ListAppVersionSettingObject, err error)
	Update(ctx context.Context, appID uuid.UUID, versionID string, listOfAppVersionSettingObject []authoring.AppVersionSettingObject) (result authoring.OperationStatus, err error)
}

var _ SettingsClientAPI = (*authoring.SettingsClient)(nil)

// AzureAccountsClientAPI contains the set of methods on the AzureAccountsClient type.
type AzureAccountsClientAPI interface {
	AssignToApp(ctx context.Context, appID uuid.UUID, azureAccountInfoObject *authoring.AzureAccountInfoObject) (result authoring.OperationStatus, err error)
	GetAssigned(ctx context.Context, appID uuid.UUID) (result authoring.ListAzureAccountInfoObject, err error)
	ListUserLUISAccounts(ctx context.Context) (result authoring.ListAzureAccountInfoObject, err error)
	RemoveFromApp(ctx context.Context, appID uuid.UUID, azureAccountInfoObject *authoring.AzureAccountInfoObject) (result authoring.OperationStatus, err error)
}

var _ AzureAccountsClientAPI = (*authoring.AzureAccountsClient)(nil)
