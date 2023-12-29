//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

import "io"

// AccountClientCreateOrUpdateResponse contains the response from method AccountClient.CreateOrUpdate.
type AccountClientCreateOrUpdateResponse struct {
	// Definition of the automation account type.
	Account
}

// AccountClientDeleteResponse contains the response from method AccountClient.Delete.
type AccountClientDeleteResponse struct {
	// placeholder for future response values
}

// AccountClientGetResponse contains the response from method AccountClient.Get.
type AccountClientGetResponse struct {
	// Definition of the automation account type.
	Account
}

// AccountClientListByResourceGroupResponse contains the response from method AccountClient.NewListByResourceGroupPager.
type AccountClientListByResourceGroupResponse struct {
	// The response model for the list account operation.
	AccountListResult
}

// AccountClientListResponse contains the response from method AccountClient.NewListPager.
type AccountClientListResponse struct {
	// The response model for the list account operation.
	AccountListResult
}

// AccountClientUpdateResponse contains the response from method AccountClient.Update.
type AccountClientUpdateResponse struct {
	// Definition of the automation account type.
	Account
}

// ActivityClientGetResponse contains the response from method ActivityClient.Get.
type ActivityClientGetResponse struct {
	// Definition of the activity.
	Activity
}

// ActivityClientListByModuleResponse contains the response from method ActivityClient.NewListByModulePager.
type ActivityClientListByModuleResponse struct {
	// The response model for the list activity operation.
	ActivityListResult
}

// AgentRegistrationInformationClientGetResponse contains the response from method AgentRegistrationInformationClient.Get.
type AgentRegistrationInformationClientGetResponse struct {
	// Definition of the agent registration information type.
	AgentRegistration
}

// AgentRegistrationInformationClientRegenerateKeyResponse contains the response from method AgentRegistrationInformationClient.RegenerateKey.
type AgentRegistrationInformationClientRegenerateKeyResponse struct {
	// Definition of the agent registration information type.
	AgentRegistration
}

// CertificateClientCreateOrUpdateResponse contains the response from method CertificateClient.CreateOrUpdate.
type CertificateClientCreateOrUpdateResponse struct {
	// Definition of the certificate.
	Certificate
}

// CertificateClientDeleteResponse contains the response from method CertificateClient.Delete.
type CertificateClientDeleteResponse struct {
	// placeholder for future response values
}

// CertificateClientGetResponse contains the response from method CertificateClient.Get.
type CertificateClientGetResponse struct {
	// Definition of the certificate.
	Certificate
}

// CertificateClientListByAutomationAccountResponse contains the response from method CertificateClient.NewListByAutomationAccountPager.
type CertificateClientListByAutomationAccountResponse struct {
	// The response model for the list certificate operation.
	CertificateListResult
}

// CertificateClientUpdateResponse contains the response from method CertificateClient.Update.
type CertificateClientUpdateResponse struct {
	// Definition of the certificate.
	Certificate
}

// ClientConvertGraphRunbookContentResponse contains the response from method Client.ConvertGraphRunbookContent.
type ClientConvertGraphRunbookContentResponse struct {
	// Graphical Runbook Content
	GraphicalRunbookContent
}

// ConnectionClientCreateOrUpdateResponse contains the response from method ConnectionClient.CreateOrUpdate.
type ConnectionClientCreateOrUpdateResponse struct {
	// Definition of the connection.
	Connection
}

// ConnectionClientDeleteResponse contains the response from method ConnectionClient.Delete.
type ConnectionClientDeleteResponse struct {
	// placeholder for future response values
}

// ConnectionClientGetResponse contains the response from method ConnectionClient.Get.
type ConnectionClientGetResponse struct {
	// Definition of the connection.
	Connection
}

// ConnectionClientListByAutomationAccountResponse contains the response from method ConnectionClient.NewListByAutomationAccountPager.
type ConnectionClientListByAutomationAccountResponse struct {
	// The response model for the list connection operation.
	ConnectionListResult
}

// ConnectionClientUpdateResponse contains the response from method ConnectionClient.Update.
type ConnectionClientUpdateResponse struct {
	// Definition of the connection.
	Connection
}

// ConnectionTypeClientCreateOrUpdateResponse contains the response from method ConnectionTypeClient.CreateOrUpdate.
type ConnectionTypeClientCreateOrUpdateResponse struct {
	// Definition of the connection type.
	ConnectionType
}

// ConnectionTypeClientDeleteResponse contains the response from method ConnectionTypeClient.Delete.
type ConnectionTypeClientDeleteResponse struct {
	// placeholder for future response values
}

// ConnectionTypeClientGetResponse contains the response from method ConnectionTypeClient.Get.
type ConnectionTypeClientGetResponse struct {
	// Definition of the connection type.
	ConnectionType
}

// ConnectionTypeClientListByAutomationAccountResponse contains the response from method ConnectionTypeClient.NewListByAutomationAccountPager.
type ConnectionTypeClientListByAutomationAccountResponse struct {
	// The response model for the list connection type operation.
	ConnectionTypeListResult
}

// CredentialClientCreateOrUpdateResponse contains the response from method CredentialClient.CreateOrUpdate.
type CredentialClientCreateOrUpdateResponse struct {
	// Definition of the credential.
	Credential
}

// CredentialClientDeleteResponse contains the response from method CredentialClient.Delete.
type CredentialClientDeleteResponse struct {
	// placeholder for future response values
}

// CredentialClientGetResponse contains the response from method CredentialClient.Get.
type CredentialClientGetResponse struct {
	// Definition of the credential.
	Credential
}

// CredentialClientListByAutomationAccountResponse contains the response from method CredentialClient.NewListByAutomationAccountPager.
type CredentialClientListByAutomationAccountResponse struct {
	// The response model for the list credential operation.
	CredentialListResult
}

// CredentialClientUpdateResponse contains the response from method CredentialClient.Update.
type CredentialClientUpdateResponse struct {
	// Definition of the credential.
	Credential
}

// DeletedAutomationAccountsClientListBySubscriptionResponse contains the response from method DeletedAutomationAccountsClient.ListBySubscription.
type DeletedAutomationAccountsClientListBySubscriptionResponse struct {
	// The response model for the list deleted automation account.
	DeletedAutomationAccountListResult
}

// DscCompilationJobClientCreateResponse contains the response from method DscCompilationJobClient.BeginCreate.
type DscCompilationJobClientCreateResponse struct {
	// Definition of the Dsc Compilation job.
	DscCompilationJob
}

// DscCompilationJobClientGetResponse contains the response from method DscCompilationJobClient.Get.
type DscCompilationJobClientGetResponse struct {
	// Definition of the Dsc Compilation job.
	DscCompilationJob
}

// DscCompilationJobClientGetStreamResponse contains the response from method DscCompilationJobClient.GetStream.
type DscCompilationJobClientGetStreamResponse struct {
	// Definition of the job stream.
	JobStream
}

// DscCompilationJobClientListByAutomationAccountResponse contains the response from method DscCompilationJobClient.NewListByAutomationAccountPager.
type DscCompilationJobClientListByAutomationAccountResponse struct {
	// The response model for the list job operation.
	DscCompilationJobListResult
}

// DscCompilationJobStreamClientListByJobResponse contains the response from method DscCompilationJobStreamClient.ListByJob.
type DscCompilationJobStreamClientListByJobResponse struct {
	// The response model for the list job stream operation.
	JobStreamListResult
}

// DscConfigurationClientCreateOrUpdateWithJSONResponse contains the response from method DscConfigurationClient.CreateOrUpdateWithJSON.
type DscConfigurationClientCreateOrUpdateWithJSONResponse struct {
	// Definition of the configuration type.
	DscConfiguration
}

// DscConfigurationClientCreateOrUpdateWithTextResponse contains the response from method DscConfigurationClient.CreateOrUpdateWithText.
type DscConfigurationClientCreateOrUpdateWithTextResponse struct {
	// Definition of the configuration type.
	DscConfiguration
}

// DscConfigurationClientDeleteResponse contains the response from method DscConfigurationClient.Delete.
type DscConfigurationClientDeleteResponse struct {
	// placeholder for future response values
}

// DscConfigurationClientGetContentResponse contains the response from method DscConfigurationClient.GetContent.
type DscConfigurationClientGetContentResponse struct {
	Value *string
}

// DscConfigurationClientGetResponse contains the response from method DscConfigurationClient.Get.
type DscConfigurationClientGetResponse struct {
	// Definition of the configuration type.
	DscConfiguration
}

// DscConfigurationClientListByAutomationAccountResponse contains the response from method DscConfigurationClient.NewListByAutomationAccountPager.
type DscConfigurationClientListByAutomationAccountResponse struct {
	// The response model for the list configuration operation.
	DscConfigurationListResult
}

// DscConfigurationClientUpdateWithJSONResponse contains the response from method DscConfigurationClient.UpdateWithJSON.
type DscConfigurationClientUpdateWithJSONResponse struct {
	// Definition of the configuration type.
	DscConfiguration
}

// DscConfigurationClientUpdateWithTextResponse contains the response from method DscConfigurationClient.UpdateWithText.
type DscConfigurationClientUpdateWithTextResponse struct {
	// Definition of the configuration type.
	DscConfiguration
}

// DscNodeClientDeleteResponse contains the response from method DscNodeClient.Delete.
type DscNodeClientDeleteResponse struct {
	// placeholder for future response values
}

// DscNodeClientGetResponse contains the response from method DscNodeClient.Get.
type DscNodeClientGetResponse struct {
	// Definition of a DscNode
	DscNode
}

// DscNodeClientListByAutomationAccountResponse contains the response from method DscNodeClient.NewListByAutomationAccountPager.
type DscNodeClientListByAutomationAccountResponse struct {
	// The response model for the list dsc nodes operation.
	DscNodeListResult
}

// DscNodeClientUpdateResponse contains the response from method DscNodeClient.Update.
type DscNodeClientUpdateResponse struct {
	// Definition of a DscNode
	DscNode
}

// DscNodeConfigurationClientCreateOrUpdateResponse contains the response from method DscNodeConfigurationClient.BeginCreateOrUpdate.
type DscNodeConfigurationClientCreateOrUpdateResponse struct {
	// Definition of the dsc node configuration.
	DscNodeConfiguration
}

// DscNodeConfigurationClientDeleteResponse contains the response from method DscNodeConfigurationClient.Delete.
type DscNodeConfigurationClientDeleteResponse struct {
	// placeholder for future response values
}

// DscNodeConfigurationClientGetResponse contains the response from method DscNodeConfigurationClient.Get.
type DscNodeConfigurationClientGetResponse struct {
	// Definition of the dsc node configuration.
	DscNodeConfiguration
}

// DscNodeConfigurationClientListByAutomationAccountResponse contains the response from method DscNodeConfigurationClient.NewListByAutomationAccountPager.
type DscNodeConfigurationClientListByAutomationAccountResponse struct {
	// The response model for the list job operation.
	DscNodeConfigurationListResult
}

// FieldsClientListByTypeResponse contains the response from method FieldsClient.NewListByTypePager.
type FieldsClientListByTypeResponse struct {
	// The response model for the list fields operation.
	TypeFieldListResult
}

// HybridRunbookWorkerGroupClientCreateResponse contains the response from method HybridRunbookWorkerGroupClient.Create.
type HybridRunbookWorkerGroupClientCreateResponse struct {
	// Definition of hybrid runbook worker group.
	HybridRunbookWorkerGroup
}

// HybridRunbookWorkerGroupClientDeleteResponse contains the response from method HybridRunbookWorkerGroupClient.Delete.
type HybridRunbookWorkerGroupClientDeleteResponse struct {
	// placeholder for future response values
}

// HybridRunbookWorkerGroupClientGetResponse contains the response from method HybridRunbookWorkerGroupClient.Get.
type HybridRunbookWorkerGroupClientGetResponse struct {
	// Definition of hybrid runbook worker group.
	HybridRunbookWorkerGroup
}

// HybridRunbookWorkerGroupClientListByAutomationAccountResponse contains the response from method HybridRunbookWorkerGroupClient.NewListByAutomationAccountPager.
type HybridRunbookWorkerGroupClientListByAutomationAccountResponse struct {
	// The response model for the list hybrid runbook worker groups.
	HybridRunbookWorkerGroupsListResult
}

// HybridRunbookWorkerGroupClientUpdateResponse contains the response from method HybridRunbookWorkerGroupClient.Update.
type HybridRunbookWorkerGroupClientUpdateResponse struct {
	// Definition of hybrid runbook worker group.
	HybridRunbookWorkerGroup
}

// HybridRunbookWorkersClientCreateResponse contains the response from method HybridRunbookWorkersClient.Create.
type HybridRunbookWorkersClientCreateResponse struct {
	// Definition of hybrid runbook worker.
	HybridRunbookWorker
}

// HybridRunbookWorkersClientDeleteResponse contains the response from method HybridRunbookWorkersClient.Delete.
type HybridRunbookWorkersClientDeleteResponse struct {
	// placeholder for future response values
}

// HybridRunbookWorkersClientGetResponse contains the response from method HybridRunbookWorkersClient.Get.
type HybridRunbookWorkersClientGetResponse struct {
	// Definition of hybrid runbook worker.
	HybridRunbookWorker
}

// HybridRunbookWorkersClientListByHybridRunbookWorkerGroupResponse contains the response from method HybridRunbookWorkersClient.NewListByHybridRunbookWorkerGroupPager.
type HybridRunbookWorkersClientListByHybridRunbookWorkerGroupResponse struct {
	// The response model for the list hybrid runbook workers.
	HybridRunbookWorkersListResult
}

// HybridRunbookWorkersClientMoveResponse contains the response from method HybridRunbookWorkersClient.Move.
type HybridRunbookWorkersClientMoveResponse struct {
	// placeholder for future response values
}

// JobClientCreateResponse contains the response from method JobClient.Create.
type JobClientCreateResponse struct {
	// Definition of the job.
	Job
}

// JobClientGetOutputResponse contains the response from method JobClient.GetOutput.
type JobClientGetOutputResponse struct {
	Value *string
}

// JobClientGetResponse contains the response from method JobClient.Get.
type JobClientGetResponse struct {
	// Definition of the job.
	Job
}

// JobClientGetRunbookContentResponse contains the response from method JobClient.GetRunbookContent.
type JobClientGetRunbookContentResponse struct {
	Value *string
}

// JobClientListByAutomationAccountResponse contains the response from method JobClient.NewListByAutomationAccountPager.
type JobClientListByAutomationAccountResponse struct {
	// The response model for the list job operation.
	JobListResultV2
}

// JobClientResumeResponse contains the response from method JobClient.Resume.
type JobClientResumeResponse struct {
	// placeholder for future response values
}

// JobClientStopResponse contains the response from method JobClient.Stop.
type JobClientStopResponse struct {
	// placeholder for future response values
}

// JobClientSuspendResponse contains the response from method JobClient.Suspend.
type JobClientSuspendResponse struct {
	// placeholder for future response values
}

// JobScheduleClientCreateResponse contains the response from method JobScheduleClient.Create.
type JobScheduleClientCreateResponse struct {
	// Definition of the job schedule.
	JobSchedule
}

// JobScheduleClientDeleteResponse contains the response from method JobScheduleClient.Delete.
type JobScheduleClientDeleteResponse struct {
	// placeholder for future response values
}

// JobScheduleClientGetResponse contains the response from method JobScheduleClient.Get.
type JobScheduleClientGetResponse struct {
	// Definition of the job schedule.
	JobSchedule
}

// JobScheduleClientListByAutomationAccountResponse contains the response from method JobScheduleClient.NewListByAutomationAccountPager.
type JobScheduleClientListByAutomationAccountResponse struct {
	// The response model for the list job schedule operation.
	JobScheduleListResult
}

// JobStreamClientGetResponse contains the response from method JobStreamClient.Get.
type JobStreamClientGetResponse struct {
	// Definition of the job stream.
	JobStream
}

// JobStreamClientListByJobResponse contains the response from method JobStreamClient.NewListByJobPager.
type JobStreamClientListByJobResponse struct {
	// The response model for the list job stream operation.
	JobStreamListResult
}

// KeysClientListByAutomationAccountResponse contains the response from method KeysClient.ListByAutomationAccount.
type KeysClientListByAutomationAccountResponse struct {
	KeyListResult
}

// LinkedWorkspaceClientGetResponse contains the response from method LinkedWorkspaceClient.Get.
type LinkedWorkspaceClientGetResponse struct {
	// Definition of the linked workspace.
	LinkedWorkspace
}

// ModuleClientCreateOrUpdateResponse contains the response from method ModuleClient.CreateOrUpdate.
type ModuleClientCreateOrUpdateResponse struct {
	// Definition of the module type.
	Module
}

// ModuleClientDeleteResponse contains the response from method ModuleClient.Delete.
type ModuleClientDeleteResponse struct {
	// placeholder for future response values
}

// ModuleClientGetResponse contains the response from method ModuleClient.Get.
type ModuleClientGetResponse struct {
	// Definition of the module type.
	Module
}

// ModuleClientListByAutomationAccountResponse contains the response from method ModuleClient.NewListByAutomationAccountPager.
type ModuleClientListByAutomationAccountResponse struct {
	// The response model for the list module operation.
	ModuleListResult
}

// ModuleClientUpdateResponse contains the response from method ModuleClient.Update.
type ModuleClientUpdateResponse struct {
	// Definition of the module type.
	Module
}

// NodeCountInformationClientGetResponse contains the response from method NodeCountInformationClient.Get.
type NodeCountInformationClientGetResponse struct {
	// Gets the count of nodes by count type
	NodeCounts
}

// NodeReportsClientGetContentResponse contains the response from method NodeReportsClient.GetContent.
type NodeReportsClientGetContentResponse struct {
	// Anything
	Interface any
}

// NodeReportsClientGetResponse contains the response from method NodeReportsClient.Get.
type NodeReportsClientGetResponse struct {
	// Definition of the dsc node report type.
	DscNodeReport
}

// NodeReportsClientListByNodeResponse contains the response from method NodeReportsClient.NewListByNodePager.
type NodeReportsClientListByNodeResponse struct {
	// The response model for the list dsc nodes operation.
	DscNodeReportListResult
}

// ObjectDataTypesClientListFieldsByModuleAndTypeResponse contains the response from method ObjectDataTypesClient.NewListFieldsByModuleAndTypePager.
type ObjectDataTypesClientListFieldsByModuleAndTypeResponse struct {
	// The response model for the list fields operation.
	TypeFieldListResult
}

// ObjectDataTypesClientListFieldsByTypeResponse contains the response from method ObjectDataTypesClient.NewListFieldsByTypePager.
type ObjectDataTypesClientListFieldsByTypeResponse struct {
	// The response model for the list fields operation.
	TypeFieldListResult
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// The response model for the list of Automation operations
	OperationListResult
}

// PrivateEndpointConnectionsClientCreateOrUpdateResponse contains the response from method PrivateEndpointConnectionsClient.BeginCreateOrUpdate.
type PrivateEndpointConnectionsClientCreateOrUpdateResponse struct {
	// A private endpoint connection
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.BeginDelete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	// A private endpoint connection
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListByAutomationAccountResponse contains the response from method PrivateEndpointConnectionsClient.NewListByAutomationAccountPager.
type PrivateEndpointConnectionsClientListByAutomationAccountResponse struct {
	// A list of private endpoint connections
	PrivateEndpointConnectionListResult
}

// PrivateLinkResourcesClientAutomationResponse contains the response from method PrivateLinkResourcesClient.NewAutomationPager.
type PrivateLinkResourcesClientAutomationResponse struct {
	// A list of private link resources
	PrivateLinkResourceListResult
}

// Python2PackageClientCreateOrUpdateResponse contains the response from method Python2PackageClient.CreateOrUpdate.
type Python2PackageClientCreateOrUpdateResponse struct {
	// Definition of the module type.
	Module
}

// Python2PackageClientDeleteResponse contains the response from method Python2PackageClient.Delete.
type Python2PackageClientDeleteResponse struct {
	// placeholder for future response values
}

// Python2PackageClientGetResponse contains the response from method Python2PackageClient.Get.
type Python2PackageClientGetResponse struct {
	// Definition of the module type.
	Module
}

// Python2PackageClientListByAutomationAccountResponse contains the response from method Python2PackageClient.NewListByAutomationAccountPager.
type Python2PackageClientListByAutomationAccountResponse struct {
	// The response model for the list module operation.
	ModuleListResult
}

// Python2PackageClientUpdateResponse contains the response from method Python2PackageClient.Update.
type Python2PackageClientUpdateResponse struct {
	// Definition of the module type.
	Module
}

// RunbookClientCreateOrUpdateResponse contains the response from method RunbookClient.CreateOrUpdate.
type RunbookClientCreateOrUpdateResponse struct {
	// Definition of the runbook type.
	Runbook
}

// RunbookClientDeleteResponse contains the response from method RunbookClient.Delete.
type RunbookClientDeleteResponse struct {
	// placeholder for future response values
}

// RunbookClientGetContentResponse contains the response from method RunbookClient.GetContent.
type RunbookClientGetContentResponse struct {
	// placeholder for future response values
}

// RunbookClientGetResponse contains the response from method RunbookClient.Get.
type RunbookClientGetResponse struct {
	// Definition of the runbook type.
	Runbook
}

// RunbookClientListByAutomationAccountResponse contains the response from method RunbookClient.NewListByAutomationAccountPager.
type RunbookClientListByAutomationAccountResponse struct {
	// The response model for the list runbook operation.
	RunbookListResult
}

// RunbookClientPublishResponse contains the response from method RunbookClient.BeginPublish.
type RunbookClientPublishResponse struct {
	// placeholder for future response values
}

// RunbookClientUpdateResponse contains the response from method RunbookClient.Update.
type RunbookClientUpdateResponse struct {
	// Definition of the runbook type.
	Runbook
}

// RunbookDraftClientGetContentResponse contains the response from method RunbookDraftClient.GetContent.
type RunbookDraftClientGetContentResponse struct {
	// placeholder for future response values
}

// RunbookDraftClientGetResponse contains the response from method RunbookDraftClient.Get.
type RunbookDraftClientGetResponse struct {
	RunbookDraft
}

// RunbookDraftClientReplaceContentResponse contains the response from method RunbookDraftClient.BeginReplaceContent.
type RunbookDraftClientReplaceContentResponse struct {
	// Body contains the streaming response.
	Body io.ReadCloser
}

// RunbookDraftClientUndoEditResponse contains the response from method RunbookDraftClient.UndoEdit.
type RunbookDraftClientUndoEditResponse struct {
	// The response model for the undo edit runbook operation.
	RunbookDraftUndoEditResult
}

// ScheduleClientCreateOrUpdateResponse contains the response from method ScheduleClient.CreateOrUpdate.
type ScheduleClientCreateOrUpdateResponse struct {
	// Definition of the schedule.
	Schedule
}

// ScheduleClientDeleteResponse contains the response from method ScheduleClient.Delete.
type ScheduleClientDeleteResponse struct {
	// placeholder for future response values
}

// ScheduleClientGetResponse contains the response from method ScheduleClient.Get.
type ScheduleClientGetResponse struct {
	// Definition of the schedule.
	Schedule
}

// ScheduleClientListByAutomationAccountResponse contains the response from method ScheduleClient.NewListByAutomationAccountPager.
type ScheduleClientListByAutomationAccountResponse struct {
	// The response model for the list schedule operation.
	ScheduleListResult
}

// ScheduleClientUpdateResponse contains the response from method ScheduleClient.Update.
type ScheduleClientUpdateResponse struct {
	// Definition of the schedule.
	Schedule
}

// SoftwareUpdateConfigurationMachineRunsClientGetByIDResponse contains the response from method SoftwareUpdateConfigurationMachineRunsClient.GetByID.
type SoftwareUpdateConfigurationMachineRunsClientGetByIDResponse struct {
	// Software update configuration machine run model.
	SoftwareUpdateConfigurationMachineRun
}

// SoftwareUpdateConfigurationMachineRunsClientListResponse contains the response from method SoftwareUpdateConfigurationMachineRunsClient.List.
type SoftwareUpdateConfigurationMachineRunsClientListResponse struct {
	// result of listing all software update configuration machine runs
	SoftwareUpdateConfigurationMachineRunListResult
}

// SoftwareUpdateConfigurationRunsClientGetByIDResponse contains the response from method SoftwareUpdateConfigurationRunsClient.GetByID.
type SoftwareUpdateConfigurationRunsClientGetByIDResponse struct {
	// Software update configuration Run properties.
	SoftwareUpdateConfigurationRun
}

// SoftwareUpdateConfigurationRunsClientListResponse contains the response from method SoftwareUpdateConfigurationRunsClient.List.
type SoftwareUpdateConfigurationRunsClientListResponse struct {
	// result of listing all software update configuration runs
	SoftwareUpdateConfigurationRunListResult
}

// SoftwareUpdateConfigurationsClientCreateResponse contains the response from method SoftwareUpdateConfigurationsClient.Create.
type SoftwareUpdateConfigurationsClientCreateResponse struct {
	// Software update configuration properties.
	SoftwareUpdateConfiguration
}

// SoftwareUpdateConfigurationsClientDeleteResponse contains the response from method SoftwareUpdateConfigurationsClient.Delete.
type SoftwareUpdateConfigurationsClientDeleteResponse struct {
	// placeholder for future response values
}

// SoftwareUpdateConfigurationsClientGetByNameResponse contains the response from method SoftwareUpdateConfigurationsClient.GetByName.
type SoftwareUpdateConfigurationsClientGetByNameResponse struct {
	// Software update configuration properties.
	SoftwareUpdateConfiguration
}

// SoftwareUpdateConfigurationsClientListResponse contains the response from method SoftwareUpdateConfigurationsClient.List.
type SoftwareUpdateConfigurationsClientListResponse struct {
	// result of listing all software update configuration
	SoftwareUpdateConfigurationListResult
}

// SourceControlClientCreateOrUpdateResponse contains the response from method SourceControlClient.CreateOrUpdate.
type SourceControlClientCreateOrUpdateResponse struct {
	// Definition of the source control.
	SourceControl
}

// SourceControlClientDeleteResponse contains the response from method SourceControlClient.Delete.
type SourceControlClientDeleteResponse struct {
	// placeholder for future response values
}

// SourceControlClientGetResponse contains the response from method SourceControlClient.Get.
type SourceControlClientGetResponse struct {
	// Definition of the source control.
	SourceControl
}

// SourceControlClientListByAutomationAccountResponse contains the response from method SourceControlClient.NewListByAutomationAccountPager.
type SourceControlClientListByAutomationAccountResponse struct {
	// The response model for the list source controls operation.
	SourceControlListResult
}

// SourceControlClientUpdateResponse contains the response from method SourceControlClient.Update.
type SourceControlClientUpdateResponse struct {
	// Definition of the source control.
	SourceControl
}

// SourceControlSyncJobClientCreateResponse contains the response from method SourceControlSyncJobClient.Create.
type SourceControlSyncJobClientCreateResponse struct {
	// Definition of the source control sync job.
	SourceControlSyncJob
}

// SourceControlSyncJobClientGetResponse contains the response from method SourceControlSyncJobClient.Get.
type SourceControlSyncJobClientGetResponse struct {
	// Definition of the source control sync job.
	SourceControlSyncJobByID
}

// SourceControlSyncJobClientListByAutomationAccountResponse contains the response from method SourceControlSyncJobClient.NewListByAutomationAccountPager.
type SourceControlSyncJobClientListByAutomationAccountResponse struct {
	// The response model for the list source control sync jobs operation.
	SourceControlSyncJobListResult
}

// SourceControlSyncJobStreamsClientGetResponse contains the response from method SourceControlSyncJobStreamsClient.Get.
type SourceControlSyncJobStreamsClientGetResponse struct {
	// Definition of the source control sync job stream by id.
	SourceControlSyncJobStreamByID
}

// SourceControlSyncJobStreamsClientListBySyncJobResponse contains the response from method SourceControlSyncJobStreamsClient.NewListBySyncJobPager.
type SourceControlSyncJobStreamsClientListBySyncJobResponse struct {
	// The response model for the list source control sync job streams operation.
	SourceControlSyncJobStreamsListBySyncJob
}

// StatisticsClientListByAutomationAccountResponse contains the response from method StatisticsClient.NewListByAutomationAccountPager.
type StatisticsClientListByAutomationAccountResponse struct {
	// The response model for the list statistics operation.
	StatisticsListResult
}

// TestJobClientCreateResponse contains the response from method TestJobClient.Create.
type TestJobClientCreateResponse struct {
	// Definition of the test job.
	TestJob
}

// TestJobClientGetResponse contains the response from method TestJobClient.Get.
type TestJobClientGetResponse struct {
	// Definition of the test job.
	TestJob
}

// TestJobClientResumeResponse contains the response from method TestJobClient.Resume.
type TestJobClientResumeResponse struct {
	// placeholder for future response values
}

// TestJobClientStopResponse contains the response from method TestJobClient.Stop.
type TestJobClientStopResponse struct {
	// placeholder for future response values
}

// TestJobClientSuspendResponse contains the response from method TestJobClient.Suspend.
type TestJobClientSuspendResponse struct {
	// placeholder for future response values
}

// TestJobStreamsClientGetResponse contains the response from method TestJobStreamsClient.Get.
type TestJobStreamsClientGetResponse struct {
	// Definition of the job stream.
	JobStream
}

// TestJobStreamsClientListByTestJobResponse contains the response from method TestJobStreamsClient.NewListByTestJobPager.
type TestJobStreamsClientListByTestJobResponse struct {
	// The response model for the list job stream operation.
	JobStreamListResult
}

// UsagesClientListByAutomationAccountResponse contains the response from method UsagesClient.NewListByAutomationAccountPager.
type UsagesClientListByAutomationAccountResponse struct {
	// The response model for the get usage operation.
	UsageListResult
}

// VariableClientCreateOrUpdateResponse contains the response from method VariableClient.CreateOrUpdate.
type VariableClientCreateOrUpdateResponse struct {
	// Definition of the variable.
	Variable
}

// VariableClientDeleteResponse contains the response from method VariableClient.Delete.
type VariableClientDeleteResponse struct {
	// placeholder for future response values
}

// VariableClientGetResponse contains the response from method VariableClient.Get.
type VariableClientGetResponse struct {
	// Definition of the variable.
	Variable
}

// VariableClientListByAutomationAccountResponse contains the response from method VariableClient.NewListByAutomationAccountPager.
type VariableClientListByAutomationAccountResponse struct {
	// The response model for the list variables operation.
	VariableListResult
}

// VariableClientUpdateResponse contains the response from method VariableClient.Update.
type VariableClientUpdateResponse struct {
	// Definition of the variable.
	Variable
}

// WatcherClientCreateOrUpdateResponse contains the response from method WatcherClient.CreateOrUpdate.
type WatcherClientCreateOrUpdateResponse struct {
	// Definition of the watcher type.
	Watcher
}

// WatcherClientDeleteResponse contains the response from method WatcherClient.Delete.
type WatcherClientDeleteResponse struct {
	// placeholder for future response values
}

// WatcherClientGetResponse contains the response from method WatcherClient.Get.
type WatcherClientGetResponse struct {
	// Definition of the watcher type.
	Watcher
}

// WatcherClientListByAutomationAccountResponse contains the response from method WatcherClient.NewListByAutomationAccountPager.
type WatcherClientListByAutomationAccountResponse struct {
	// The response model for the list watcher operation.
	WatcherListResult
}

// WatcherClientStartResponse contains the response from method WatcherClient.Start.
type WatcherClientStartResponse struct {
	// placeholder for future response values
}

// WatcherClientStopResponse contains the response from method WatcherClient.Stop.
type WatcherClientStopResponse struct {
	// placeholder for future response values
}

// WatcherClientUpdateResponse contains the response from method WatcherClient.Update.
type WatcherClientUpdateResponse struct {
	// Definition of the watcher type.
	Watcher
}

// WebhookClientCreateOrUpdateResponse contains the response from method WebhookClient.CreateOrUpdate.
type WebhookClientCreateOrUpdateResponse struct {
	// Definition of the webhook type.
	Webhook
}

// WebhookClientDeleteResponse contains the response from method WebhookClient.Delete.
type WebhookClientDeleteResponse struct {
	// placeholder for future response values
}

// WebhookClientGenerateURIResponse contains the response from method WebhookClient.GenerateURI.
type WebhookClientGenerateURIResponse struct {
	Value *string
}

// WebhookClientGetResponse contains the response from method WebhookClient.Get.
type WebhookClientGetResponse struct {
	// Definition of the webhook type.
	Webhook
}

// WebhookClientListByAutomationAccountResponse contains the response from method WebhookClient.NewListByAutomationAccountPager.
type WebhookClientListByAutomationAccountResponse struct {
	// The response model for the list webhook operation.
	WebhookListResult
}

// WebhookClientUpdateResponse contains the response from method WebhookClient.Update.
type WebhookClientUpdateResponse struct {
	// Definition of the webhook type.
	Webhook
}
