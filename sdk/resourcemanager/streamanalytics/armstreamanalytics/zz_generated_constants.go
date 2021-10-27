//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstreamanalytics

const (
	module  = "armstreamanalytics"
	version = "v0.1.1"
)

// AuthenticationMode - Authentication Mode. Valid modes are ConnectionString, Msi and 'UserToken'.
type AuthenticationMode string

const (
	AuthenticationModeConnectionString AuthenticationMode = "ConnectionString"
	AuthenticationModeMsi              AuthenticationMode = "Msi"
	AuthenticationModeUserToken        AuthenticationMode = "UserToken"
)

// PossibleAuthenticationModeValues returns the possible values for the AuthenticationMode const type.
func PossibleAuthenticationModeValues() []AuthenticationMode {
	return []AuthenticationMode{
		AuthenticationModeConnectionString,
		AuthenticationModeMsi,
		AuthenticationModeUserToken,
	}
}

// ToPtr returns a *AuthenticationMode pointing to the current value.
func (c AuthenticationMode) ToPtr() *AuthenticationMode {
	return &c
}

// ClusterProvisioningState - The status of the cluster provisioning. The three terminal states are: Succeeded, Failed and Canceled
type ClusterProvisioningState string

const (
	// ClusterProvisioningStateCanceled - The cluster provisioning was canceled.
	ClusterProvisioningStateCanceled ClusterProvisioningState = "Canceled"
	// ClusterProvisioningStateFailed - The cluster provisioning failed.
	ClusterProvisioningStateFailed ClusterProvisioningState = "Failed"
	// ClusterProvisioningStateInProgress - The cluster provisioning was inprogress.
	ClusterProvisioningStateInProgress ClusterProvisioningState = "InProgress"
	// ClusterProvisioningStateSucceeded - The cluster provisioning succeeded.
	ClusterProvisioningStateSucceeded ClusterProvisioningState = "Succeeded"
)

// PossibleClusterProvisioningStateValues returns the possible values for the ClusterProvisioningState const type.
func PossibleClusterProvisioningStateValues() []ClusterProvisioningState {
	return []ClusterProvisioningState{
		ClusterProvisioningStateCanceled,
		ClusterProvisioningStateFailed,
		ClusterProvisioningStateInProgress,
		ClusterProvisioningStateSucceeded,
	}
}

// ToPtr returns a *ClusterProvisioningState pointing to the current value.
func (c ClusterProvisioningState) ToPtr() *ClusterProvisioningState {
	return &c
}

// ClusterSKUName - Specifies the SKU name of the cluster. Required on PUT (CreateOrUpdate) requests.
type ClusterSKUName string

const (
	// ClusterSKUNameDefault - The default SKU.
	ClusterSKUNameDefault ClusterSKUName = "Default"
)

// PossibleClusterSKUNameValues returns the possible values for the ClusterSKUName const type.
func PossibleClusterSKUNameValues() []ClusterSKUName {
	return []ClusterSKUName{
		ClusterSKUNameDefault,
	}
}

// ToPtr returns a *ClusterSKUName pointing to the current value.
func (c ClusterSKUName) ToPtr() *ClusterSKUName {
	return &c
}

// CompatibilityLevel - Controls certain runtime behaviors of the streaming job.
type CompatibilityLevel string

const (
	CompatibilityLevelOne0 CompatibilityLevel = "1.0"
)

// PossibleCompatibilityLevelValues returns the possible values for the CompatibilityLevel const type.
func PossibleCompatibilityLevelValues() []CompatibilityLevel {
	return []CompatibilityLevel{
		CompatibilityLevelOne0,
	}
}

// ToPtr returns a *CompatibilityLevel pointing to the current value.
func (c CompatibilityLevel) ToPtr() *CompatibilityLevel {
	return &c
}

// ContentStoragePolicy - Valid values are JobStorageAccount and SystemAccount. If set to JobStorageAccount, this requires the user to also specify jobStorageAccount
// property. .
type ContentStoragePolicy string

const (
	ContentStoragePolicyJobStorageAccount ContentStoragePolicy = "JobStorageAccount"
	ContentStoragePolicySystemAccount     ContentStoragePolicy = "SystemAccount"
)

// PossibleContentStoragePolicyValues returns the possible values for the ContentStoragePolicy const type.
func PossibleContentStoragePolicyValues() []ContentStoragePolicy {
	return []ContentStoragePolicy{
		ContentStoragePolicyJobStorageAccount,
		ContentStoragePolicySystemAccount,
	}
}

// ToPtr returns a *ContentStoragePolicy pointing to the current value.
func (c ContentStoragePolicy) ToPtr() *ContentStoragePolicy {
	return &c
}

// Encoding - Specifies the encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output.
type Encoding string

const (
	EncodingUTF8 Encoding = "UTF8"
)

// PossibleEncodingValues returns the possible values for the Encoding const type.
func PossibleEncodingValues() []Encoding {
	return []Encoding{
		EncodingUTF8,
	}
}

// ToPtr returns a *Encoding pointing to the current value.
func (c Encoding) ToPtr() *Encoding {
	return &c
}

// EventSerializationType - Indicates the type of serialization that the input or output uses. Required on PUT (CreateOrReplace) requests.
type EventSerializationType string

const (
	EventSerializationTypeAvro      EventSerializationType = "Avro"
	EventSerializationTypeCSV       EventSerializationType = "Csv"
	EventSerializationTypeCustomClr EventSerializationType = "CustomClr"
	EventSerializationTypeJSON      EventSerializationType = "Json"
	EventSerializationTypeParquet   EventSerializationType = "Parquet"
)

// PossibleEventSerializationTypeValues returns the possible values for the EventSerializationType const type.
func PossibleEventSerializationTypeValues() []EventSerializationType {
	return []EventSerializationType{
		EventSerializationTypeAvro,
		EventSerializationTypeCSV,
		EventSerializationTypeCustomClr,
		EventSerializationTypeJSON,
		EventSerializationTypeParquet,
	}
}

// ToPtr returns a *EventSerializationType pointing to the current value.
func (c EventSerializationType) ToPtr() *EventSerializationType {
	return &c
}

// EventsOutOfOrderPolicy - Indicates the policy to apply to events that arrive out of order in the input event stream.
type EventsOutOfOrderPolicy string

const (
	EventsOutOfOrderPolicyAdjust EventsOutOfOrderPolicy = "Adjust"
	EventsOutOfOrderPolicyDrop   EventsOutOfOrderPolicy = "Drop"
)

// PossibleEventsOutOfOrderPolicyValues returns the possible values for the EventsOutOfOrderPolicy const type.
func PossibleEventsOutOfOrderPolicyValues() []EventsOutOfOrderPolicy {
	return []EventsOutOfOrderPolicy{
		EventsOutOfOrderPolicyAdjust,
		EventsOutOfOrderPolicyDrop,
	}
}

// ToPtr returns a *EventsOutOfOrderPolicy pointing to the current value.
func (c EventsOutOfOrderPolicy) ToPtr() *EventsOutOfOrderPolicy {
	return &c
}

// JSONOutputSerializationFormat - Specifies the format of the JSON the output will be written in. The currently supported values are 'lineSeparated' indicating
// the output will be formatted by having each JSON object separated by a new
// line and 'array' indicating the output will be formatted as an array of JSON objects.
type JSONOutputSerializationFormat string

const (
	JSONOutputSerializationFormatArray         JSONOutputSerializationFormat = "Array"
	JSONOutputSerializationFormatLineSeparated JSONOutputSerializationFormat = "LineSeparated"
)

// PossibleJSONOutputSerializationFormatValues returns the possible values for the JSONOutputSerializationFormat const type.
func PossibleJSONOutputSerializationFormatValues() []JSONOutputSerializationFormat {
	return []JSONOutputSerializationFormat{
		JSONOutputSerializationFormatArray,
		JSONOutputSerializationFormatLineSeparated,
	}
}

// ToPtr returns a *JSONOutputSerializationFormat pointing to the current value.
func (c JSONOutputSerializationFormat) ToPtr() *JSONOutputSerializationFormat {
	return &c
}

// JobState - The current execution state of the streaming job.
type JobState string

const (
	// JobStateCreated - The job is currently in the Created state.
	JobStateCreated JobState = "Created"
	// JobStateDegraded - The job is currently in the Degraded state.
	JobStateDegraded JobState = "Degraded"
	// JobStateDeleting - The job is currently in the Deleting state.
	JobStateDeleting JobState = "Deleting"
	// JobStateFailed - The job is currently in the Failed state.
	JobStateFailed JobState = "Failed"
	// JobStateRestarting - The job is currently in the Restarting state.
	JobStateRestarting JobState = "Restarting"
	// JobStateRunning - The job is currently in the Running state.
	JobStateRunning JobState = "Running"
	// JobStateScaling - The job is currently in the Scaling state.
	JobStateScaling JobState = "Scaling"
	// JobStateStarting - The job is currently in the Starting state.
	JobStateStarting JobState = "Starting"
	// JobStateStopped - The job is currently in the Stopped state.
	JobStateStopped JobState = "Stopped"
	// JobStateStopping - The job is currently in the Stopping state.
	JobStateStopping JobState = "Stopping"
)

// PossibleJobStateValues returns the possible values for the JobState const type.
func PossibleJobStateValues() []JobState {
	return []JobState{
		JobStateCreated,
		JobStateDegraded,
		JobStateDeleting,
		JobStateFailed,
		JobStateRestarting,
		JobStateRunning,
		JobStateScaling,
		JobStateStarting,
		JobStateStopped,
		JobStateStopping,
	}
}

// ToPtr returns a *JobState pointing to the current value.
func (c JobState) ToPtr() *JobState {
	return &c
}

// JobType - Describes the type of the job. Valid modes are Cloud and 'Edge'.
type JobType string

const (
	JobTypeCloud JobType = "Cloud"
	JobTypeEdge  JobType = "Edge"
)

// PossibleJobTypeValues returns the possible values for the JobType const type.
func PossibleJobTypeValues() []JobType {
	return []JobType{
		JobTypeCloud,
		JobTypeEdge,
	}
}

// ToPtr returns a *JobType pointing to the current value.
func (c JobType) ToPtr() *JobType {
	return &c
}

// OutputErrorPolicy - Indicates the policy to apply to events that arrive at the output and cannot be written to the external storage due to being malformed
// (missing column values, column values of wrong type or size).
type OutputErrorPolicy string

const (
	OutputErrorPolicyDrop OutputErrorPolicy = "Drop"
	OutputErrorPolicyStop OutputErrorPolicy = "Stop"
)

// PossibleOutputErrorPolicyValues returns the possible values for the OutputErrorPolicy const type.
func PossibleOutputErrorPolicyValues() []OutputErrorPolicy {
	return []OutputErrorPolicy{
		OutputErrorPolicyDrop,
		OutputErrorPolicyStop,
	}
}

// ToPtr returns a *OutputErrorPolicy pointing to the current value.
func (c OutputErrorPolicy) ToPtr() *OutputErrorPolicy {
	return &c
}

// OutputStartMode - Value may be JobStartTime, CustomTime, or LastOutputEventTime to indicate whether the starting point of the output event stream should
// start whenever the job is started, start at a custom user time
// stamp specified via the outputStartTime property, or start from the last event output time.
type OutputStartMode string

const (
	OutputStartModeCustomTime          OutputStartMode = "CustomTime"
	OutputStartModeJobStartTime        OutputStartMode = "JobStartTime"
	OutputStartModeLastOutputEventTime OutputStartMode = "LastOutputEventTime"
)

// PossibleOutputStartModeValues returns the possible values for the OutputStartMode const type.
func PossibleOutputStartModeValues() []OutputStartMode {
	return []OutputStartMode{
		OutputStartModeCustomTime,
		OutputStartModeJobStartTime,
		OutputStartModeLastOutputEventTime,
	}
}

// ToPtr returns a *OutputStartMode pointing to the current value.
func (c OutputStartMode) ToPtr() *OutputStartMode {
	return &c
}

// QueryTestingResultStatus - The status of the query testing request.
type QueryTestingResultStatus string

const (
	// QueryTestingResultStatusCompilerError - The query testing operation failed due to a compiler error.
	QueryTestingResultStatusCompilerError QueryTestingResultStatus = "CompilerError"
	// QueryTestingResultStatusRuntimeError - The query testing operation failed due to a runtime error.
	QueryTestingResultStatusRuntimeError QueryTestingResultStatus = "RuntimeError"
	// QueryTestingResultStatusStarted - The query testing operation was initiated.
	QueryTestingResultStatusStarted QueryTestingResultStatus = "Started"
	// QueryTestingResultStatusSuccess - The query testing operation succeeded.
	QueryTestingResultStatusSuccess QueryTestingResultStatus = "Success"
	// QueryTestingResultStatusTimeout - The query testing operation failed due to a timeout.
	QueryTestingResultStatusTimeout QueryTestingResultStatus = "Timeout"
	// QueryTestingResultStatusUnknownError - The query testing operation failed due to an unknown error .
	QueryTestingResultStatusUnknownError QueryTestingResultStatus = "UnknownError"
)

// PossibleQueryTestingResultStatusValues returns the possible values for the QueryTestingResultStatus const type.
func PossibleQueryTestingResultStatusValues() []QueryTestingResultStatus {
	return []QueryTestingResultStatus{
		QueryTestingResultStatusCompilerError,
		QueryTestingResultStatusRuntimeError,
		QueryTestingResultStatusStarted,
		QueryTestingResultStatusSuccess,
		QueryTestingResultStatusTimeout,
		QueryTestingResultStatusUnknownError,
	}
}

// ToPtr returns a *QueryTestingResultStatus pointing to the current value.
func (c QueryTestingResultStatus) ToPtr() *QueryTestingResultStatus {
	return &c
}

// SampleInputResultStatus - The status of the sample input request.
type SampleInputResultStatus string

const (
	// SampleInputResultStatusErrorConnectingToInput - The sample input operation failed to connect to the input.
	SampleInputResultStatusErrorConnectingToInput SampleInputResultStatus = "ErrorConnectingToInput"
	// SampleInputResultStatusNoEventsFoundInRange - The sample input operation found no events in the range.
	SampleInputResultStatusNoEventsFoundInRange SampleInputResultStatus = "NoEventsFoundInRange"
	// SampleInputResultStatusReadAllEventsInRange - The sample input operation successfully read all the events in the range.
	SampleInputResultStatusReadAllEventsInRange SampleInputResultStatus = "ReadAllEventsInRange"
)

// PossibleSampleInputResultStatusValues returns the possible values for the SampleInputResultStatus const type.
func PossibleSampleInputResultStatusValues() []SampleInputResultStatus {
	return []SampleInputResultStatus{
		SampleInputResultStatusErrorConnectingToInput,
		SampleInputResultStatusNoEventsFoundInRange,
		SampleInputResultStatusReadAllEventsInRange,
	}
}

// ToPtr returns a *SampleInputResultStatus pointing to the current value.
func (c SampleInputResultStatus) ToPtr() *SampleInputResultStatus {
	return &c
}

// StreamingJobSKUName - The name of the SKU. Required on PUT (CreateOrReplace) requests.
type StreamingJobSKUName string

const (
	StreamingJobSKUNameStandard StreamingJobSKUName = "Standard"
)

// PossibleStreamingJobSKUNameValues returns the possible values for the StreamingJobSKUName const type.
func PossibleStreamingJobSKUNameValues() []StreamingJobSKUName {
	return []StreamingJobSKUName{
		StreamingJobSKUNameStandard,
	}
}

// ToPtr returns a *StreamingJobSKUName pointing to the current value.
func (c StreamingJobSKUName) ToPtr() *StreamingJobSKUName {
	return &c
}

// TestDatasourceResultStatus - The status of the test input or output request.
type TestDatasourceResultStatus string

const (
	// TestDatasourceResultStatusTestFailed - The test datasource operation failed.
	TestDatasourceResultStatusTestFailed TestDatasourceResultStatus = "TestFailed"
	// TestDatasourceResultStatusTestSucceeded - The test datasource operation succeeded.
	TestDatasourceResultStatusTestSucceeded TestDatasourceResultStatus = "TestSucceeded"
)

// PossibleTestDatasourceResultStatusValues returns the possible values for the TestDatasourceResultStatus const type.
func PossibleTestDatasourceResultStatusValues() []TestDatasourceResultStatus {
	return []TestDatasourceResultStatus{
		TestDatasourceResultStatusTestFailed,
		TestDatasourceResultStatusTestSucceeded,
	}
}

// ToPtr returns a *TestDatasourceResultStatus pointing to the current value.
func (c TestDatasourceResultStatus) ToPtr() *TestDatasourceResultStatus {
	return &c
}
