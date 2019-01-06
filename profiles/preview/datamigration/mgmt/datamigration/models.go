// +build go1.9

// Copyright 2019 Microsoft Corporation
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

package datamigration

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/datamigration/mgmt/2018-04-19/datamigration"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AuthenticationType = original.AuthenticationType

const (
	ActiveDirectoryIntegrated AuthenticationType = original.ActiveDirectoryIntegrated
	ActiveDirectoryPassword   AuthenticationType = original.ActiveDirectoryPassword
	None                      AuthenticationType = original.None
	SQLAuthentication         AuthenticationType = original.SQLAuthentication
	WindowsAuthentication     AuthenticationType = original.WindowsAuthentication
)

type DatabaseCompatLevel = original.DatabaseCompatLevel

const (
	CompatLevel100 DatabaseCompatLevel = original.CompatLevel100
	CompatLevel110 DatabaseCompatLevel = original.CompatLevel110
	CompatLevel120 DatabaseCompatLevel = original.CompatLevel120
	CompatLevel130 DatabaseCompatLevel = original.CompatLevel130
	CompatLevel140 DatabaseCompatLevel = original.CompatLevel140
	CompatLevel80  DatabaseCompatLevel = original.CompatLevel80
	CompatLevel90  DatabaseCompatLevel = original.CompatLevel90
)

type DatabaseFileType = original.DatabaseFileType

const (
	Filestream   DatabaseFileType = original.Filestream
	Fulltext     DatabaseFileType = original.Fulltext
	Log          DatabaseFileType = original.Log
	NotSupported DatabaseFileType = original.NotSupported
	Rows         DatabaseFileType = original.Rows
)

type DatabaseMigrationStage = original.DatabaseMigrationStage

const (
	DatabaseMigrationStageBackup     DatabaseMigrationStage = original.DatabaseMigrationStageBackup
	DatabaseMigrationStageCompleted  DatabaseMigrationStage = original.DatabaseMigrationStageCompleted
	DatabaseMigrationStageFileCopy   DatabaseMigrationStage = original.DatabaseMigrationStageFileCopy
	DatabaseMigrationStageInitialize DatabaseMigrationStage = original.DatabaseMigrationStageInitialize
	DatabaseMigrationStageNone       DatabaseMigrationStage = original.DatabaseMigrationStageNone
	DatabaseMigrationStageRestore    DatabaseMigrationStage = original.DatabaseMigrationStageRestore
)

type DatabaseState = original.DatabaseState

const (
	Copying          DatabaseState = original.Copying
	Emergency        DatabaseState = original.Emergency
	Offline          DatabaseState = original.Offline
	OfflineSecondary DatabaseState = original.OfflineSecondary
	Online           DatabaseState = original.Online
	Recovering       DatabaseState = original.Recovering
	RecoveryPending  DatabaseState = original.RecoveryPending
	Restoring        DatabaseState = original.Restoring
	Suspect          DatabaseState = original.Suspect
)

type ErrorType = original.ErrorType

const (
	ErrorTypeDefault ErrorType = original.ErrorTypeDefault
	ErrorTypeError   ErrorType = original.ErrorTypeError
	ErrorTypeWarning ErrorType = original.ErrorTypeWarning
)

type LoginMigrationStage = original.LoginMigrationStage

const (
	LoginMigrationStageAssignRoleMembership       LoginMigrationStage = original.LoginMigrationStageAssignRoleMembership
	LoginMigrationStageAssignRoleOwnership        LoginMigrationStage = original.LoginMigrationStageAssignRoleOwnership
	LoginMigrationStageCompleted                  LoginMigrationStage = original.LoginMigrationStageCompleted
	LoginMigrationStageEstablishObjectPermissions LoginMigrationStage = original.LoginMigrationStageEstablishObjectPermissions
	LoginMigrationStageEstablishServerPermissions LoginMigrationStage = original.LoginMigrationStageEstablishServerPermissions
	LoginMigrationStageEstablishUserMapping       LoginMigrationStage = original.LoginMigrationStageEstablishUserMapping
	LoginMigrationStageInitialize                 LoginMigrationStage = original.LoginMigrationStageInitialize
	LoginMigrationStageLoginMigration             LoginMigrationStage = original.LoginMigrationStageLoginMigration
	LoginMigrationStageNone                       LoginMigrationStage = original.LoginMigrationStageNone
)

type LoginType = original.LoginType

const (
	AsymmetricKey LoginType = original.AsymmetricKey
	Certificate   LoginType = original.Certificate
	ExternalGroup LoginType = original.ExternalGroup
	ExternalUser  LoginType = original.ExternalUser
	SQLLogin      LoginType = original.SQLLogin
	WindowsGroup  LoginType = original.WindowsGroup
	WindowsUser   LoginType = original.WindowsUser
)

type MigrationState = original.MigrationState

const (
	MigrationStateCompleted  MigrationState = original.MigrationStateCompleted
	MigrationStateFailed     MigrationState = original.MigrationStateFailed
	MigrationStateInProgress MigrationState = original.MigrationStateInProgress
	MigrationStateNone       MigrationState = original.MigrationStateNone
	MigrationStateSkipped    MigrationState = original.MigrationStateSkipped
	MigrationStateStopped    MigrationState = original.MigrationStateStopped
	MigrationStateWarning    MigrationState = original.MigrationStateWarning
)

type MigrationStatus = original.MigrationStatus

const (
	MigrationStatusCompleted               MigrationStatus = original.MigrationStatusCompleted
	MigrationStatusCompletedWithWarnings   MigrationStatus = original.MigrationStatusCompletedWithWarnings
	MigrationStatusConfigured              MigrationStatus = original.MigrationStatusConfigured
	MigrationStatusConnecting              MigrationStatus = original.MigrationStatusConnecting
	MigrationStatusDefault                 MigrationStatus = original.MigrationStatusDefault
	MigrationStatusError                   MigrationStatus = original.MigrationStatusError
	MigrationStatusRunning                 MigrationStatus = original.MigrationStatusRunning
	MigrationStatusSelectLogins            MigrationStatus = original.MigrationStatusSelectLogins
	MigrationStatusSourceAndTargetSelected MigrationStatus = original.MigrationStatusSourceAndTargetSelected
	MigrationStatusStopped                 MigrationStatus = original.MigrationStatusStopped
)

type NameCheckFailureReason = original.NameCheckFailureReason

const (
	AlreadyExists NameCheckFailureReason = original.AlreadyExists
	Invalid       NameCheckFailureReason = original.Invalid
)

type ObjectType = original.ObjectType

const (
	Function         ObjectType = original.Function
	StoredProcedures ObjectType = original.StoredProcedures
	Table            ObjectType = original.Table
	User             ObjectType = original.User
	View             ObjectType = original.View
)

type ProjectProvisioningState = original.ProjectProvisioningState

const (
	Deleting  ProjectProvisioningState = original.Deleting
	Succeeded ProjectProvisioningState = original.Succeeded
)

type ProjectSourcePlatform = original.ProjectSourcePlatform

const (
	SQL     ProjectSourcePlatform = original.SQL
	Unknown ProjectSourcePlatform = original.Unknown
)

type ProjectTargetPlatform = original.ProjectTargetPlatform

const (
	ProjectTargetPlatformSQLDB   ProjectTargetPlatform = original.ProjectTargetPlatformSQLDB
	ProjectTargetPlatformUnknown ProjectTargetPlatform = original.ProjectTargetPlatformUnknown
)

type ResourceSkuCapacityScaleType = original.ResourceSkuCapacityScaleType

const (
	ResourceSkuCapacityScaleTypeAutomatic ResourceSkuCapacityScaleType = original.ResourceSkuCapacityScaleTypeAutomatic
	ResourceSkuCapacityScaleTypeManual    ResourceSkuCapacityScaleType = original.ResourceSkuCapacityScaleTypeManual
	ResourceSkuCapacityScaleTypeNone      ResourceSkuCapacityScaleType = original.ResourceSkuCapacityScaleTypeNone
)

type ResourceSkuRestrictionsReasonCode = original.ResourceSkuRestrictionsReasonCode

const (
	NotAvailableForSubscription ResourceSkuRestrictionsReasonCode = original.NotAvailableForSubscription
	QuotaID                     ResourceSkuRestrictionsReasonCode = original.QuotaID
)

type ResourceSkuRestrictionsType = original.ResourceSkuRestrictionsType

const (
	Location ResourceSkuRestrictionsType = original.Location
)

type ResultType = original.ResultType

const (
	ResultTypeDatabaseLevelOutput                    ResultType = original.ResultTypeDatabaseLevelOutput
	ResultTypeErrorOutput                            ResultType = original.ResultTypeErrorOutput
	ResultTypeMigrateSQLServerSQLDbTaskOutput        ResultType = original.ResultTypeMigrateSQLServerSQLDbTaskOutput
	ResultTypeMigrationDatabaseLevelValidationOutput ResultType = original.ResultTypeMigrationDatabaseLevelValidationOutput
	ResultTypeMigrationLevelOutput                   ResultType = original.ResultTypeMigrationLevelOutput
	ResultTypeMigrationValidationOutput              ResultType = original.ResultTypeMigrationValidationOutput
	ResultTypeTableLevelOutput                       ResultType = original.ResultTypeTableLevelOutput
)

type ResultTypeBasicConnectToSourceSQLServerTaskOutput = original.ResultTypeBasicConnectToSourceSQLServerTaskOutput

const (
	ResultTypeBasicConnectToSourceSQLServerTaskOutputResultTypeAgentJobLevelOutput                ResultTypeBasicConnectToSourceSQLServerTaskOutput = original.ResultTypeBasicConnectToSourceSQLServerTaskOutputResultTypeAgentJobLevelOutput
	ResultTypeBasicConnectToSourceSQLServerTaskOutputResultTypeConnectToSourceSQLServerTaskOutput ResultTypeBasicConnectToSourceSQLServerTaskOutput = original.ResultTypeBasicConnectToSourceSQLServerTaskOutputResultTypeConnectToSourceSQLServerTaskOutput
	ResultTypeBasicConnectToSourceSQLServerTaskOutputResultTypeDatabaseLevelOutput                ResultTypeBasicConnectToSourceSQLServerTaskOutput = original.ResultTypeBasicConnectToSourceSQLServerTaskOutputResultTypeDatabaseLevelOutput
	ResultTypeBasicConnectToSourceSQLServerTaskOutputResultTypeLoginLevelOutput                   ResultTypeBasicConnectToSourceSQLServerTaskOutput = original.ResultTypeBasicConnectToSourceSQLServerTaskOutputResultTypeLoginLevelOutput
	ResultTypeBasicConnectToSourceSQLServerTaskOutputResultTypeTaskLevelOutput                    ResultTypeBasicConnectToSourceSQLServerTaskOutput = original.ResultTypeBasicConnectToSourceSQLServerTaskOutputResultTypeTaskLevelOutput
)

type ServerLevelPermissionsGroup = original.ServerLevelPermissionsGroup

const (
	Default                         ServerLevelPermissionsGroup = original.Default
	MigrationFromSQLServerToAzureDB ServerLevelPermissionsGroup = original.MigrationFromSQLServerToAzureDB
)

type ServiceProvisioningState = original.ServiceProvisioningState

const (
	ServiceProvisioningStateAccepted      ServiceProvisioningState = original.ServiceProvisioningStateAccepted
	ServiceProvisioningStateDeleting      ServiceProvisioningState = original.ServiceProvisioningStateDeleting
	ServiceProvisioningStateDeploying     ServiceProvisioningState = original.ServiceProvisioningStateDeploying
	ServiceProvisioningStateFailed        ServiceProvisioningState = original.ServiceProvisioningStateFailed
	ServiceProvisioningStateFailedToStart ServiceProvisioningState = original.ServiceProvisioningStateFailedToStart
	ServiceProvisioningStateFailedToStop  ServiceProvisioningState = original.ServiceProvisioningStateFailedToStop
	ServiceProvisioningStateStarting      ServiceProvisioningState = original.ServiceProvisioningStateStarting
	ServiceProvisioningStateStopped       ServiceProvisioningState = original.ServiceProvisioningStateStopped
	ServiceProvisioningStateStopping      ServiceProvisioningState = original.ServiceProvisioningStateStopping
	ServiceProvisioningStateSucceeded     ServiceProvisioningState = original.ServiceProvisioningStateSucceeded
)

type ServiceScalability = original.ServiceScalability

const (
	ServiceScalabilityAutomatic ServiceScalability = original.ServiceScalabilityAutomatic
	ServiceScalabilityManual    ServiceScalability = original.ServiceScalabilityManual
	ServiceScalabilityNone      ServiceScalability = original.ServiceScalabilityNone
)

type Severity = original.Severity

const (
	SeverityError   Severity = original.SeverityError
	SeverityMessage Severity = original.SeverityMessage
	SeverityWarning Severity = original.SeverityWarning
)

type TaskState = original.TaskState

const (
	TaskStateCanceled              TaskState = original.TaskStateCanceled
	TaskStateFailed                TaskState = original.TaskStateFailed
	TaskStateFailedInputValidation TaskState = original.TaskStateFailedInputValidation
	TaskStateFaulted               TaskState = original.TaskStateFaulted
	TaskStateQueued                TaskState = original.TaskStateQueued
	TaskStateRunning               TaskState = original.TaskStateRunning
	TaskStateSucceeded             TaskState = original.TaskStateSucceeded
	TaskStateUnknown               TaskState = original.TaskStateUnknown
)

type TaskType = original.TaskType

const (
	TaskTypeConnectToSourceSQLServer TaskType = original.TaskTypeConnectToSourceSQLServer
	TaskTypeConnectToTargetSQLDb     TaskType = original.TaskTypeConnectToTargetSQLDb
	TaskTypeGetUserTablesSQL         TaskType = original.TaskTypeGetUserTablesSQL
	TaskTypeMigrateSQLServerSQLDb    TaskType = original.TaskTypeMigrateSQLServerSQLDb
	TaskTypeUnknown                  TaskType = original.TaskTypeUnknown
)

type Type = original.Type

const (
	TypeSQLConnectionInfo Type = original.TypeSQLConnectionInfo
	TypeUnknown           Type = original.TypeUnknown
)

type UpdateActionType = original.UpdateActionType

const (
	AddedOnTarget   UpdateActionType = original.AddedOnTarget
	ChangedOnTarget UpdateActionType = original.ChangedOnTarget
	DeletedOnTarget UpdateActionType = original.DeletedOnTarget
)

type ValidationStatus = original.ValidationStatus

const (
	ValidationStatusCompleted           ValidationStatus = original.ValidationStatusCompleted
	ValidationStatusCompletedWithIssues ValidationStatus = original.ValidationStatusCompletedWithIssues
	ValidationStatusDefault             ValidationStatus = original.ValidationStatusDefault
	ValidationStatusFailed              ValidationStatus = original.ValidationStatusFailed
	ValidationStatusInitialized         ValidationStatus = original.ValidationStatusInitialized
	ValidationStatusInProgress          ValidationStatus = original.ValidationStatusInProgress
	ValidationStatusNotStarted          ValidationStatus = original.ValidationStatusNotStarted
	ValidationStatusStopped             ValidationStatus = original.ValidationStatusStopped
)

type APIError = original.APIError
type AvailableServiceSku = original.AvailableServiceSku
type AvailableServiceSkuCapacity = original.AvailableServiceSkuCapacity
type AvailableServiceSkuSku = original.AvailableServiceSkuSku
type BaseClient = original.BaseClient
type BasicConnectToSourceSQLServerTaskOutput = original.BasicConnectToSourceSQLServerTaskOutput
type BasicConnectionInfo = original.BasicConnectionInfo
type BasicMigrateSQLServerSQLDbTaskOutput = original.BasicMigrateSQLServerSQLDbTaskOutput
type BasicProjectTaskProperties = original.BasicProjectTaskProperties
type BlobShare = original.BlobShare
type ConnectToSourceSQLServerTaskInput = original.ConnectToSourceSQLServerTaskInput
type ConnectToSourceSQLServerTaskOutput = original.ConnectToSourceSQLServerTaskOutput
type ConnectToSourceSQLServerTaskOutputAgentJobLevel = original.ConnectToSourceSQLServerTaskOutputAgentJobLevel
type ConnectToSourceSQLServerTaskOutputDatabaseLevel = original.ConnectToSourceSQLServerTaskOutputDatabaseLevel
type ConnectToSourceSQLServerTaskOutputLoginLevel = original.ConnectToSourceSQLServerTaskOutputLoginLevel
type ConnectToSourceSQLServerTaskOutputTaskLevel = original.ConnectToSourceSQLServerTaskOutputTaskLevel
type ConnectToSourceSQLServerTaskProperties = original.ConnectToSourceSQLServerTaskProperties
type ConnectToTargetSQLDbTaskInput = original.ConnectToTargetSQLDbTaskInput
type ConnectToTargetSQLDbTaskOutput = original.ConnectToTargetSQLDbTaskOutput
type ConnectToTargetSQLDbTaskProperties = original.ConnectToTargetSQLDbTaskProperties
type ConnectionInfo = original.ConnectionInfo
type DataIntegrityValidationResult = original.DataIntegrityValidationResult
type DataItemMigrationSummaryResult = original.DataItemMigrationSummaryResult
type Database = original.Database
type DatabaseFileInfo = original.DatabaseFileInfo
type DatabaseFileInput = original.DatabaseFileInput
type DatabaseInfo = original.DatabaseInfo
type DatabaseObjectName = original.DatabaseObjectName
type DatabaseSummaryResult = original.DatabaseSummaryResult
type DatabaseTable = original.DatabaseTable
type Error = original.Error
type ExecutionStatistics = original.ExecutionStatistics
type FileShare = original.FileShare
type GetUserTablesSQLTaskInput = original.GetUserTablesSQLTaskInput
type GetUserTablesSQLTaskOutput = original.GetUserTablesSQLTaskOutput
type GetUserTablesSQLTaskProperties = original.GetUserTablesSQLTaskProperties
type MigrateSQLServerSQLDbDatabaseInput = original.MigrateSQLServerSQLDbDatabaseInput
type MigrateSQLServerSQLDbTaskInput = original.MigrateSQLServerSQLDbTaskInput
type MigrateSQLServerSQLDbTaskOutput = original.MigrateSQLServerSQLDbTaskOutput
type MigrateSQLServerSQLDbTaskOutputDatabaseLevel = original.MigrateSQLServerSQLDbTaskOutputDatabaseLevel
type MigrateSQLServerSQLDbTaskOutputDatabaseLevelValidationResult = original.MigrateSQLServerSQLDbTaskOutputDatabaseLevelValidationResult
type MigrateSQLServerSQLDbTaskOutputError = original.MigrateSQLServerSQLDbTaskOutputError
type MigrateSQLServerSQLDbTaskOutputMigrationLevel = original.MigrateSQLServerSQLDbTaskOutputMigrationLevel
type MigrateSQLServerSQLDbTaskOutputTableLevel = original.MigrateSQLServerSQLDbTaskOutputTableLevel
type MigrateSQLServerSQLDbTaskOutputValidationResult = original.MigrateSQLServerSQLDbTaskOutputValidationResult
type MigrateSQLServerSQLDbTaskProperties = original.MigrateSQLServerSQLDbTaskProperties
type MigrateSQLServerSQLServerDatabaseInput = original.MigrateSQLServerSQLServerDatabaseInput
type MigrationEligibilityInfo = original.MigrationEligibilityInfo
type MigrationReportResult = original.MigrationReportResult
type MigrationTableMetadata = original.MigrationTableMetadata
type MigrationValidationDatabaseSummaryResult = original.MigrationValidationDatabaseSummaryResult
type MigrationValidationOptions = original.MigrationValidationOptions
type NameAvailabilityRequest = original.NameAvailabilityRequest
type NameAvailabilityResponse = original.NameAvailabilityResponse
type ODataError = original.ODataError
type OperationsClient = original.OperationsClient
type Project = original.Project
type ProjectList = original.ProjectList
type ProjectListIterator = original.ProjectListIterator
type ProjectListPage = original.ProjectListPage
type ProjectMetadata = original.ProjectMetadata
type ProjectProperties = original.ProjectProperties
type ProjectTask = original.ProjectTask
type ProjectTaskProperties = original.ProjectTaskProperties
type ProjectsClient = original.ProjectsClient
type QueryAnalysisValidationResult = original.QueryAnalysisValidationResult
type QueryExecutionResult = original.QueryExecutionResult
type Quota = original.Quota
type QuotaList = original.QuotaList
type QuotaListIterator = original.QuotaListIterator
type QuotaListPage = original.QuotaListPage
type QuotaName = original.QuotaName
type ReportableException = original.ReportableException
type Resource = original.Resource
type ResourceSku = original.ResourceSku
type ResourceSkuCapabilities = original.ResourceSkuCapabilities
type ResourceSkuCapacity = original.ResourceSkuCapacity
type ResourceSkuCosts = original.ResourceSkuCosts
type ResourceSkuRestrictions = original.ResourceSkuRestrictions
type ResourceSkusClient = original.ResourceSkusClient
type ResourceSkusResult = original.ResourceSkusResult
type ResourceSkusResultIterator = original.ResourceSkusResultIterator
type ResourceSkusResultPage = original.ResourceSkusResultPage
type SQLConnectionInfo = original.SQLConnectionInfo
type SQLMigrationTaskInput = original.SQLMigrationTaskInput
type SchemaComparisonValidationResult = original.SchemaComparisonValidationResult
type SchemaComparisonValidationResultType = original.SchemaComparisonValidationResultType
type Service = original.Service
type ServiceList = original.ServiceList
type ServiceListIterator = original.ServiceListIterator
type ServiceListPage = original.ServiceListPage
type ServiceOperation = original.ServiceOperation
type ServiceOperationDisplay = original.ServiceOperationDisplay
type ServiceOperationList = original.ServiceOperationList
type ServiceOperationListIterator = original.ServiceOperationListIterator
type ServiceOperationListPage = original.ServiceOperationListPage
type ServiceProperties = original.ServiceProperties
type ServiceSku = original.ServiceSku
type ServiceSkuList = original.ServiceSkuList
type ServiceSkuListIterator = original.ServiceSkuListIterator
type ServiceSkuListPage = original.ServiceSkuListPage
type ServiceStatusResponse = original.ServiceStatusResponse
type ServicesClient = original.ServicesClient
type ServicesCreateOrUpdateFuture = original.ServicesCreateOrUpdateFuture
type ServicesDeleteFuture = original.ServicesDeleteFuture
type ServicesStartFuture = original.ServicesStartFuture
type ServicesStopFuture = original.ServicesStopFuture
type ServicesUpdateFuture = original.ServicesUpdateFuture
type StartMigrationScenarioServerRoleResult = original.StartMigrationScenarioServerRoleResult
type TaskList = original.TaskList
type TaskListIterator = original.TaskListIterator
type TaskListPage = original.TaskListPage
type TasksClient = original.TasksClient
type TrackedResource = original.TrackedResource
type UsagesClient = original.UsagesClient
type ValidationError = original.ValidationError
type WaitStatistics = original.WaitStatistics

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewProjectListIterator(page ProjectListPage) ProjectListIterator {
	return original.NewProjectListIterator(page)
}
func NewProjectListPage(getNextPage func(context.Context, ProjectList) (ProjectList, error)) ProjectListPage {
	return original.NewProjectListPage(getNextPage)
}
func NewProjectsClient(subscriptionID string) ProjectsClient {
	return original.NewProjectsClient(subscriptionID)
}
func NewProjectsClientWithBaseURI(baseURI string, subscriptionID string) ProjectsClient {
	return original.NewProjectsClientWithBaseURI(baseURI, subscriptionID)
}
func NewQuotaListIterator(page QuotaListPage) QuotaListIterator {
	return original.NewQuotaListIterator(page)
}
func NewQuotaListPage(getNextPage func(context.Context, QuotaList) (QuotaList, error)) QuotaListPage {
	return original.NewQuotaListPage(getNextPage)
}
func NewResourceSkusClient(subscriptionID string) ResourceSkusClient {
	return original.NewResourceSkusClient(subscriptionID)
}
func NewResourceSkusClientWithBaseURI(baseURI string, subscriptionID string) ResourceSkusClient {
	return original.NewResourceSkusClientWithBaseURI(baseURI, subscriptionID)
}
func NewResourceSkusResultIterator(page ResourceSkusResultPage) ResourceSkusResultIterator {
	return original.NewResourceSkusResultIterator(page)
}
func NewResourceSkusResultPage(getNextPage func(context.Context, ResourceSkusResult) (ResourceSkusResult, error)) ResourceSkusResultPage {
	return original.NewResourceSkusResultPage(getNextPage)
}
func NewServiceListIterator(page ServiceListPage) ServiceListIterator {
	return original.NewServiceListIterator(page)
}
func NewServiceListPage(getNextPage func(context.Context, ServiceList) (ServiceList, error)) ServiceListPage {
	return original.NewServiceListPage(getNextPage)
}
func NewServiceOperationListIterator(page ServiceOperationListPage) ServiceOperationListIterator {
	return original.NewServiceOperationListIterator(page)
}
func NewServiceOperationListPage(getNextPage func(context.Context, ServiceOperationList) (ServiceOperationList, error)) ServiceOperationListPage {
	return original.NewServiceOperationListPage(getNextPage)
}
func NewServiceSkuListIterator(page ServiceSkuListPage) ServiceSkuListIterator {
	return original.NewServiceSkuListIterator(page)
}
func NewServiceSkuListPage(getNextPage func(context.Context, ServiceSkuList) (ServiceSkuList, error)) ServiceSkuListPage {
	return original.NewServiceSkuListPage(getNextPage)
}
func NewServicesClient(subscriptionID string) ServicesClient {
	return original.NewServicesClient(subscriptionID)
}
func NewServicesClientWithBaseURI(baseURI string, subscriptionID string) ServicesClient {
	return original.NewServicesClientWithBaseURI(baseURI, subscriptionID)
}
func NewTaskListIterator(page TaskListPage) TaskListIterator {
	return original.NewTaskListIterator(page)
}
func NewTaskListPage(getNextPage func(context.Context, TaskList) (TaskList, error)) TaskListPage {
	return original.NewTaskListPage(getNextPage)
}
func NewTasksClient(subscriptionID string) TasksClient {
	return original.NewTasksClient(subscriptionID)
}
func NewTasksClientWithBaseURI(baseURI string, subscriptionID string) TasksClient {
	return original.NewTasksClientWithBaseURI(baseURI, subscriptionID)
}
func NewUsagesClient(subscriptionID string) UsagesClient {
	return original.NewUsagesClient(subscriptionID)
}
func NewUsagesClientWithBaseURI(baseURI string, subscriptionID string) UsagesClient {
	return original.NewUsagesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAuthenticationTypeValues() []AuthenticationType {
	return original.PossibleAuthenticationTypeValues()
}
func PossibleDatabaseCompatLevelValues() []DatabaseCompatLevel {
	return original.PossibleDatabaseCompatLevelValues()
}
func PossibleDatabaseFileTypeValues() []DatabaseFileType {
	return original.PossibleDatabaseFileTypeValues()
}
func PossibleDatabaseMigrationStageValues() []DatabaseMigrationStage {
	return original.PossibleDatabaseMigrationStageValues()
}
func PossibleDatabaseStateValues() []DatabaseState {
	return original.PossibleDatabaseStateValues()
}
func PossibleErrorTypeValues() []ErrorType {
	return original.PossibleErrorTypeValues()
}
func PossibleLoginMigrationStageValues() []LoginMigrationStage {
	return original.PossibleLoginMigrationStageValues()
}
func PossibleLoginTypeValues() []LoginType {
	return original.PossibleLoginTypeValues()
}
func PossibleMigrationStateValues() []MigrationState {
	return original.PossibleMigrationStateValues()
}
func PossibleMigrationStatusValues() []MigrationStatus {
	return original.PossibleMigrationStatusValues()
}
func PossibleNameCheckFailureReasonValues() []NameCheckFailureReason {
	return original.PossibleNameCheckFailureReasonValues()
}
func PossibleObjectTypeValues() []ObjectType {
	return original.PossibleObjectTypeValues()
}
func PossibleProjectProvisioningStateValues() []ProjectProvisioningState {
	return original.PossibleProjectProvisioningStateValues()
}
func PossibleProjectSourcePlatformValues() []ProjectSourcePlatform {
	return original.PossibleProjectSourcePlatformValues()
}
func PossibleProjectTargetPlatformValues() []ProjectTargetPlatform {
	return original.PossibleProjectTargetPlatformValues()
}
func PossibleResourceSkuCapacityScaleTypeValues() []ResourceSkuCapacityScaleType {
	return original.PossibleResourceSkuCapacityScaleTypeValues()
}
func PossibleResourceSkuRestrictionsReasonCodeValues() []ResourceSkuRestrictionsReasonCode {
	return original.PossibleResourceSkuRestrictionsReasonCodeValues()
}
func PossibleResourceSkuRestrictionsTypeValues() []ResourceSkuRestrictionsType {
	return original.PossibleResourceSkuRestrictionsTypeValues()
}
func PossibleResultTypeBasicConnectToSourceSQLServerTaskOutputValues() []ResultTypeBasicConnectToSourceSQLServerTaskOutput {
	return original.PossibleResultTypeBasicConnectToSourceSQLServerTaskOutputValues()
}
func PossibleResultTypeValues() []ResultType {
	return original.PossibleResultTypeValues()
}
func PossibleServerLevelPermissionsGroupValues() []ServerLevelPermissionsGroup {
	return original.PossibleServerLevelPermissionsGroupValues()
}
func PossibleServiceProvisioningStateValues() []ServiceProvisioningState {
	return original.PossibleServiceProvisioningStateValues()
}
func PossibleServiceScalabilityValues() []ServiceScalability {
	return original.PossibleServiceScalabilityValues()
}
func PossibleSeverityValues() []Severity {
	return original.PossibleSeverityValues()
}
func PossibleTaskStateValues() []TaskState {
	return original.PossibleTaskStateValues()
}
func PossibleTaskTypeValues() []TaskType {
	return original.PossibleTaskTypeValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func PossibleUpdateActionTypeValues() []UpdateActionType {
	return original.PossibleUpdateActionTypeValues()
}
func PossibleValidationStatusValues() []ValidationStatus {
	return original.PossibleValidationStatusValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
