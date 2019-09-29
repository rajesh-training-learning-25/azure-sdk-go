package logicapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2019-05-01/logic"
	"github.com/Azure/go-autorest/autorest"
)

// WorkflowsClientAPI contains the set of methods on the WorkflowsClient type.
type WorkflowsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workflowName string, workflow logic.Workflow) (result logic.Workflow, err error)
	Delete(ctx context.Context, resourceGroupName string, workflowName string) (result autorest.Response, err error)
	Disable(ctx context.Context, resourceGroupName string, workflowName string) (result autorest.Response, err error)
	Enable(ctx context.Context, resourceGroupName string, workflowName string) (result autorest.Response, err error)
	GenerateUpgradedDefinition(ctx context.Context, resourceGroupName string, workflowName string, parameters logic.GenerateUpgradedDefinitionParameters) (result logic.SetObject, err error)
	Get(ctx context.Context, resourceGroupName string, workflowName string) (result logic.Workflow, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, top *int32, filter string) (result logic.WorkflowListResultPage, err error)
	ListBySubscription(ctx context.Context, top *int32, filter string) (result logic.WorkflowListResultPage, err error)
	ListCallbackURL(ctx context.Context, resourceGroupName string, workflowName string, listCallbackURL logic.GetCallbackURLParameters) (result logic.WorkflowTriggerCallbackURL, err error)
	ListSwagger(ctx context.Context, resourceGroupName string, workflowName string) (result logic.SetObject, err error)
	Move(ctx context.Context, resourceGroupName string, workflowName string, move logic.Workflow) (result logic.WorkflowsMoveFuture, err error)
	RegenerateAccessKey(ctx context.Context, resourceGroupName string, workflowName string, keyType logic.RegenerateActionParameter) (result autorest.Response, err error)
	Update(ctx context.Context, resourceGroupName string, workflowName string, workflow logic.Workflow) (result logic.Workflow, err error)
	ValidateByLocation(ctx context.Context, resourceGroupName string, location string, workflowName string) (result autorest.Response, err error)
	ValidateByResourceGroup(ctx context.Context, resourceGroupName string, workflowName string, validate logic.Workflow) (result autorest.Response, err error)
}

var _ WorkflowsClientAPI = (*logic.WorkflowsClient)(nil)

// WorkflowVersionsClientAPI contains the set of methods on the WorkflowVersionsClient type.
type WorkflowVersionsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, workflowName string, versionID string) (result logic.WorkflowVersion, err error)
	List(ctx context.Context, resourceGroupName string, workflowName string, top *int32) (result logic.WorkflowVersionListResultPage, err error)
}

var _ WorkflowVersionsClientAPI = (*logic.WorkflowVersionsClient)(nil)

// WorkflowTriggersClientAPI contains the set of methods on the WorkflowTriggersClient type.
type WorkflowTriggersClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, workflowName string, triggerName string) (result logic.WorkflowTrigger, err error)
	GetSchemaJSON(ctx context.Context, resourceGroupName string, workflowName string, triggerName string) (result logic.JSONSchema, err error)
	List(ctx context.Context, resourceGroupName string, workflowName string, top *int32, filter string) (result logic.WorkflowTriggerListResultPage, err error)
	ListCallbackURL(ctx context.Context, resourceGroupName string, workflowName string, triggerName string) (result logic.WorkflowTriggerCallbackURL, err error)
	Reset(ctx context.Context, resourceGroupName string, workflowName string, triggerName string) (result autorest.Response, err error)
	Run(ctx context.Context, resourceGroupName string, workflowName string, triggerName string) (result logic.SetObject, err error)
	SetState(ctx context.Context, resourceGroupName string, workflowName string, triggerName string, setState logic.SetTriggerStateActionDefinition) (result autorest.Response, err error)
}

var _ WorkflowTriggersClientAPI = (*logic.WorkflowTriggersClient)(nil)

// WorkflowVersionTriggersClientAPI contains the set of methods on the WorkflowVersionTriggersClient type.
type WorkflowVersionTriggersClientAPI interface {
	ListCallbackURL(ctx context.Context, resourceGroupName string, workflowName string, versionID string, triggerName string, parameters *logic.GetCallbackURLParameters) (result logic.WorkflowTriggerCallbackURL, err error)
}

var _ WorkflowVersionTriggersClientAPI = (*logic.WorkflowVersionTriggersClient)(nil)

// WorkflowTriggerHistoriesClientAPI contains the set of methods on the WorkflowTriggerHistoriesClient type.
type WorkflowTriggerHistoriesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, workflowName string, triggerName string, historyName string) (result logic.WorkflowTriggerHistory, err error)
	List(ctx context.Context, resourceGroupName string, workflowName string, triggerName string, top *int32, filter string) (result logic.WorkflowTriggerHistoryListResultPage, err error)
	Resubmit(ctx context.Context, resourceGroupName string, workflowName string, triggerName string, historyName string) (result autorest.Response, err error)
}

var _ WorkflowTriggerHistoriesClientAPI = (*logic.WorkflowTriggerHistoriesClient)(nil)

// WorkflowRunsClientAPI contains the set of methods on the WorkflowRunsClient type.
type WorkflowRunsClientAPI interface {
	Cancel(ctx context.Context, resourceGroupName string, workflowName string, runName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workflowName string, runName string) (result logic.WorkflowRun, err error)
	List(ctx context.Context, resourceGroupName string, workflowName string, top *int32, filter string) (result logic.WorkflowRunListResultPage, err error)
}

var _ WorkflowRunsClientAPI = (*logic.WorkflowRunsClient)(nil)

// WorkflowRunActionsClientAPI contains the set of methods on the WorkflowRunActionsClient type.
type WorkflowRunActionsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string) (result logic.WorkflowRunAction, err error)
	List(ctx context.Context, resourceGroupName string, workflowName string, runName string, top *int32, filter string) (result logic.WorkflowRunActionListResultPage, err error)
	ListExpressionTraces(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string) (result logic.ExpressionTraces, err error)
}

var _ WorkflowRunActionsClientAPI = (*logic.WorkflowRunActionsClient)(nil)

// WorkflowRunActionRepetitionsClientAPI contains the set of methods on the WorkflowRunActionRepetitionsClient type.
type WorkflowRunActionRepetitionsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string, repetitionName string) (result logic.WorkflowRunActionRepetitionDefinition, err error)
	List(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string) (result logic.WorkflowRunActionRepetitionDefinitionCollection, err error)
	ListExpressionTraces(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string, repetitionName string) (result logic.ExpressionTraces, err error)
}

var _ WorkflowRunActionRepetitionsClientAPI = (*logic.WorkflowRunActionRepetitionsClient)(nil)

// WorkflowRunActionRepetitionsRequestHistoriesClientAPI contains the set of methods on the WorkflowRunActionRepetitionsRequestHistoriesClient type.
type WorkflowRunActionRepetitionsRequestHistoriesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string, repetitionName string, requestHistoryName string) (result logic.RequestHistory, err error)
	List(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string, repetitionName string) (result logic.RequestHistoryListResultPage, err error)
}

var _ WorkflowRunActionRepetitionsRequestHistoriesClientAPI = (*logic.WorkflowRunActionRepetitionsRequestHistoriesClient)(nil)

// WorkflowRunActionRequestHistoriesClientAPI contains the set of methods on the WorkflowRunActionRequestHistoriesClient type.
type WorkflowRunActionRequestHistoriesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string, requestHistoryName string) (result logic.RequestHistory, err error)
	List(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string) (result logic.RequestHistoryListResultPage, err error)
}

var _ WorkflowRunActionRequestHistoriesClientAPI = (*logic.WorkflowRunActionRequestHistoriesClient)(nil)

// WorkflowRunActionScopeRepetitionsClientAPI contains the set of methods on the WorkflowRunActionScopeRepetitionsClient type.
type WorkflowRunActionScopeRepetitionsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string, repetitionName string) (result logic.WorkflowRunActionRepetitionDefinition, err error)
	List(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string) (result logic.WorkflowRunActionRepetitionDefinitionCollection, err error)
}

var _ WorkflowRunActionScopeRepetitionsClientAPI = (*logic.WorkflowRunActionScopeRepetitionsClient)(nil)

// WorkflowRunOperationsClientAPI contains the set of methods on the WorkflowRunOperationsClient type.
type WorkflowRunOperationsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, workflowName string, runName string, operationID string) (result logic.WorkflowRun, err error)
}

var _ WorkflowRunOperationsClientAPI = (*logic.WorkflowRunOperationsClient)(nil)

// IntegrationAccountsClientAPI contains the set of methods on the IntegrationAccountsClient type.
type IntegrationAccountsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, integrationAccountName string, integrationAccount logic.IntegrationAccount) (result logic.IntegrationAccount, err error)
	Delete(ctx context.Context, resourceGroupName string, integrationAccountName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, integrationAccountName string) (result logic.IntegrationAccount, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, top *int32) (result logic.IntegrationAccountListResultPage, err error)
	ListBySubscription(ctx context.Context, top *int32) (result logic.IntegrationAccountListResultPage, err error)
	ListCallbackURL(ctx context.Context, resourceGroupName string, integrationAccountName string, parameters logic.GetCallbackURLParameters) (result logic.CallbackURL, err error)
	ListKeyVaultKeys(ctx context.Context, resourceGroupName string, integrationAccountName string, listKeyVaultKeys logic.ListKeyVaultKeysDefinition) (result logic.KeyVaultKeyCollection, err error)
	LogTrackingEvents(ctx context.Context, resourceGroupName string, integrationAccountName string, logTrackingEvents logic.TrackingEventsDefinition) (result autorest.Response, err error)
	RegenerateAccessKey(ctx context.Context, resourceGroupName string, integrationAccountName string, regenerateAccessKey logic.RegenerateActionParameter) (result logic.IntegrationAccount, err error)
	Update(ctx context.Context, resourceGroupName string, integrationAccountName string, integrationAccount logic.IntegrationAccount) (result logic.IntegrationAccount, err error)
}

var _ IntegrationAccountsClientAPI = (*logic.IntegrationAccountsClient)(nil)

// IntegrationAccountAssembliesClientAPI contains the set of methods on the IntegrationAccountAssembliesClient type.
type IntegrationAccountAssembliesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, integrationAccountName string, assemblyArtifactName string, assemblyArtifact logic.AssemblyDefinition) (result logic.AssemblyDefinition, err error)
	Delete(ctx context.Context, resourceGroupName string, integrationAccountName string, assemblyArtifactName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, integrationAccountName string, assemblyArtifactName string) (result logic.AssemblyDefinition, err error)
	List(ctx context.Context, resourceGroupName string, integrationAccountName string) (result logic.AssemblyCollection, err error)
	ListContentCallbackURL(ctx context.Context, resourceGroupName string, integrationAccountName string, assemblyArtifactName string) (result logic.WorkflowTriggerCallbackURL, err error)
}

var _ IntegrationAccountAssembliesClientAPI = (*logic.IntegrationAccountAssembliesClient)(nil)

// IntegrationAccountBatchConfigurationsClientAPI contains the set of methods on the IntegrationAccountBatchConfigurationsClient type.
type IntegrationAccountBatchConfigurationsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, integrationAccountName string, batchConfigurationName string, batchConfiguration logic.BatchConfiguration) (result logic.BatchConfiguration, err error)
	Delete(ctx context.Context, resourceGroupName string, integrationAccountName string, batchConfigurationName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, integrationAccountName string, batchConfigurationName string) (result logic.BatchConfiguration, err error)
	List(ctx context.Context, resourceGroupName string, integrationAccountName string) (result logic.BatchConfigurationCollection, err error)
}

var _ IntegrationAccountBatchConfigurationsClientAPI = (*logic.IntegrationAccountBatchConfigurationsClient)(nil)

// IntegrationAccountSchemasClientAPI contains the set of methods on the IntegrationAccountSchemasClient type.
type IntegrationAccountSchemasClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, integrationAccountName string, schemaName string, schema logic.IntegrationAccountSchema) (result logic.IntegrationAccountSchema, err error)
	Delete(ctx context.Context, resourceGroupName string, integrationAccountName string, schemaName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, integrationAccountName string, schemaName string) (result logic.IntegrationAccountSchema, err error)
	List(ctx context.Context, resourceGroupName string, integrationAccountName string, top *int32, filter string) (result logic.IntegrationAccountSchemaListResultPage, err error)
	ListContentCallbackURL(ctx context.Context, resourceGroupName string, integrationAccountName string, schemaName string, listContentCallbackURL logic.GetCallbackURLParameters) (result logic.WorkflowTriggerCallbackURL, err error)
}

var _ IntegrationAccountSchemasClientAPI = (*logic.IntegrationAccountSchemasClient)(nil)

// IntegrationAccountMapsClientAPI contains the set of methods on the IntegrationAccountMapsClient type.
type IntegrationAccountMapsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, integrationAccountName string, mapName string, mapParameter logic.IntegrationAccountMap) (result logic.IntegrationAccountMap, err error)
	Delete(ctx context.Context, resourceGroupName string, integrationAccountName string, mapName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, integrationAccountName string, mapName string) (result logic.IntegrationAccountMap, err error)
	List(ctx context.Context, resourceGroupName string, integrationAccountName string, top *int32, filter string) (result logic.IntegrationAccountMapListResultPage, err error)
	ListContentCallbackURL(ctx context.Context, resourceGroupName string, integrationAccountName string, mapName string, listContentCallbackURL logic.GetCallbackURLParameters) (result logic.WorkflowTriggerCallbackURL, err error)
}

var _ IntegrationAccountMapsClientAPI = (*logic.IntegrationAccountMapsClient)(nil)

// IntegrationAccountPartnersClientAPI contains the set of methods on the IntegrationAccountPartnersClient type.
type IntegrationAccountPartnersClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, integrationAccountName string, partnerName string, partner logic.IntegrationAccountPartner) (result logic.IntegrationAccountPartner, err error)
	Delete(ctx context.Context, resourceGroupName string, integrationAccountName string, partnerName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, integrationAccountName string, partnerName string) (result logic.IntegrationAccountPartner, err error)
	List(ctx context.Context, resourceGroupName string, integrationAccountName string, top *int32, filter string) (result logic.IntegrationAccountPartnerListResultPage, err error)
	ListContentCallbackURL(ctx context.Context, resourceGroupName string, integrationAccountName string, partnerName string, listContentCallbackURL logic.GetCallbackURLParameters) (result logic.WorkflowTriggerCallbackURL, err error)
}

var _ IntegrationAccountPartnersClientAPI = (*logic.IntegrationAccountPartnersClient)(nil)

// IntegrationAccountAgreementsClientAPI contains the set of methods on the IntegrationAccountAgreementsClient type.
type IntegrationAccountAgreementsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, integrationAccountName string, agreementName string, agreement logic.IntegrationAccountAgreement) (result logic.IntegrationAccountAgreement, err error)
	Delete(ctx context.Context, resourceGroupName string, integrationAccountName string, agreementName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, integrationAccountName string, agreementName string) (result logic.IntegrationAccountAgreement, err error)
	List(ctx context.Context, resourceGroupName string, integrationAccountName string, top *int32, filter string) (result logic.IntegrationAccountAgreementListResultPage, err error)
	ListContentCallbackURL(ctx context.Context, resourceGroupName string, integrationAccountName string, agreementName string, listContentCallbackURL logic.GetCallbackURLParameters) (result logic.WorkflowTriggerCallbackURL, err error)
}

var _ IntegrationAccountAgreementsClientAPI = (*logic.IntegrationAccountAgreementsClient)(nil)

// IntegrationAccountCertificatesClientAPI contains the set of methods on the IntegrationAccountCertificatesClient type.
type IntegrationAccountCertificatesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, integrationAccountName string, certificateName string, certificate logic.IntegrationAccountCertificate) (result logic.IntegrationAccountCertificate, err error)
	Delete(ctx context.Context, resourceGroupName string, integrationAccountName string, certificateName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, integrationAccountName string, certificateName string) (result logic.IntegrationAccountCertificate, err error)
	List(ctx context.Context, resourceGroupName string, integrationAccountName string, top *int32) (result logic.IntegrationAccountCertificateListResultPage, err error)
}

var _ IntegrationAccountCertificatesClientAPI = (*logic.IntegrationAccountCertificatesClient)(nil)

// IntegrationAccountSessionsClientAPI contains the set of methods on the IntegrationAccountSessionsClient type.
type IntegrationAccountSessionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, integrationAccountName string, sessionName string, session logic.IntegrationAccountSession) (result logic.IntegrationAccountSession, err error)
	Delete(ctx context.Context, resourceGroupName string, integrationAccountName string, sessionName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, integrationAccountName string, sessionName string) (result logic.IntegrationAccountSession, err error)
	List(ctx context.Context, resourceGroupName string, integrationAccountName string, top *int32, filter string) (result logic.IntegrationAccountSessionListResultPage, err error)
}

var _ IntegrationAccountSessionsClientAPI = (*logic.IntegrationAccountSessionsClient)(nil)

// IntegrationServiceEnvironmentsClientAPI contains the set of methods on the IntegrationServiceEnvironmentsClient type.
type IntegrationServiceEnvironmentsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, integrationServiceEnvironment logic.IntegrationServiceEnvironment) (result logic.IntegrationServiceEnvironmentsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string) (result logic.IntegrationServiceEnvironment, err error)
	ListByResourceGroup(ctx context.Context, resourceGroup string, top *int32) (result logic.IntegrationServiceEnvironmentListResultPage, err error)
	ListBySubscription(ctx context.Context, top *int32) (result logic.IntegrationServiceEnvironmentListResultPage, err error)
	Restart(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string) (result autorest.Response, err error)
	Update(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, integrationServiceEnvironment logic.IntegrationServiceEnvironment) (result logic.IntegrationServiceEnvironmentsUpdateFuture, err error)
}

var _ IntegrationServiceEnvironmentsClientAPI = (*logic.IntegrationServiceEnvironmentsClient)(nil)

// IntegrationServiceEnvironmentSkusClientAPI contains the set of methods on the IntegrationServiceEnvironmentSkusClient type.
type IntegrationServiceEnvironmentSkusClientAPI interface {
	List(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string) (result logic.IntegrationServiceEnvironmentSkuListPage, err error)
}

var _ IntegrationServiceEnvironmentSkusClientAPI = (*logic.IntegrationServiceEnvironmentSkusClient)(nil)

// IntegrationServiceEnvironmentNetworkHealthClientAPI contains the set of methods on the IntegrationServiceEnvironmentNetworkHealthClient type.
type IntegrationServiceEnvironmentNetworkHealthClientAPI interface {
	Get(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string) (result logic.SetIntegrationServiceEnvironmentSubnetNetworkHealth, err error)
}

var _ IntegrationServiceEnvironmentNetworkHealthClientAPI = (*logic.IntegrationServiceEnvironmentNetworkHealthClient)(nil)

// IntegrationServiceEnvironmentManagedApisClientAPI contains the set of methods on the IntegrationServiceEnvironmentManagedApisClient type.
type IntegrationServiceEnvironmentManagedApisClientAPI interface {
	Delete(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, APIName string) (result logic.IntegrationServiceEnvironmentManagedApisDeleteFuture, err error)
	Get(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, APIName string) (result logic.ManagedAPI, err error)
	List(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string) (result logic.ManagedAPIListResultPage, err error)
	Put(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, APIName string) (result logic.IntegrationServiceEnvironmentManagedApisPutFuture, err error)
}

var _ IntegrationServiceEnvironmentManagedApisClientAPI = (*logic.IntegrationServiceEnvironmentManagedApisClient)(nil)

// IntegrationServiceEnvironmentManagedAPIOperationsClientAPI contains the set of methods on the IntegrationServiceEnvironmentManagedAPIOperationsClient type.
type IntegrationServiceEnvironmentManagedAPIOperationsClientAPI interface {
	List(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, APIName string) (result logic.APIOperationListResultPage, err error)
}

var _ IntegrationServiceEnvironmentManagedAPIOperationsClientAPI = (*logic.IntegrationServiceEnvironmentManagedAPIOperationsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result logic.OperationListResultPage, err error)
}

var _ OperationsClientAPI = (*logic.OperationsClient)(nil)
