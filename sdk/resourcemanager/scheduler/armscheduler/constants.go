//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armscheduler

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/scheduler/armscheduler"
	moduleVersion = "v1.2.0"
)

type DayOfWeek string

const (
	DayOfWeekFriday    DayOfWeek = "Friday"
	DayOfWeekMonday    DayOfWeek = "Monday"
	DayOfWeekSaturday  DayOfWeek = "Saturday"
	DayOfWeekSunday    DayOfWeek = "Sunday"
	DayOfWeekThursday  DayOfWeek = "Thursday"
	DayOfWeekTuesday   DayOfWeek = "Tuesday"
	DayOfWeekWednesday DayOfWeek = "Wednesday"
)

// PossibleDayOfWeekValues returns the possible values for the DayOfWeek const type.
func PossibleDayOfWeekValues() []DayOfWeek {
	return []DayOfWeek{
		DayOfWeekFriday,
		DayOfWeekMonday,
		DayOfWeekSaturday,
		DayOfWeekSunday,
		DayOfWeekThursday,
		DayOfWeekTuesday,
		DayOfWeekWednesday,
	}
}

// HTTPAuthenticationType - Gets or sets the HTTP authentication type.
type HTTPAuthenticationType string

const (
	HTTPAuthenticationTypeActiveDirectoryOAuth HTTPAuthenticationType = "ActiveDirectoryOAuth"
	HTTPAuthenticationTypeBasic                HTTPAuthenticationType = "Basic"
	HTTPAuthenticationTypeClientCertificate    HTTPAuthenticationType = "ClientCertificate"
	HTTPAuthenticationTypeNotSpecified         HTTPAuthenticationType = "NotSpecified"
)

// PossibleHTTPAuthenticationTypeValues returns the possible values for the HTTPAuthenticationType const type.
func PossibleHTTPAuthenticationTypeValues() []HTTPAuthenticationType {
	return []HTTPAuthenticationType{
		HTTPAuthenticationTypeActiveDirectoryOAuth,
		HTTPAuthenticationTypeBasic,
		HTTPAuthenticationTypeClientCertificate,
		HTTPAuthenticationTypeNotSpecified,
	}
}

// JobActionType - Gets or sets the job action type.
type JobActionType string

const (
	JobActionTypeHTTP            JobActionType = "Http"
	JobActionTypeHTTPS           JobActionType = "Https"
	JobActionTypeServiceBusQueue JobActionType = "ServiceBusQueue"
	JobActionTypeServiceBusTopic JobActionType = "ServiceBusTopic"
	JobActionTypeStorageQueue    JobActionType = "StorageQueue"
)

// PossibleJobActionTypeValues returns the possible values for the JobActionType const type.
func PossibleJobActionTypeValues() []JobActionType {
	return []JobActionType{
		JobActionTypeHTTP,
		JobActionTypeHTTPS,
		JobActionTypeServiceBusQueue,
		JobActionTypeServiceBusTopic,
		JobActionTypeStorageQueue,
	}
}

// JobCollectionState - Gets or sets the state.
type JobCollectionState string

const (
	JobCollectionStateDeleted   JobCollectionState = "Deleted"
	JobCollectionStateDisabled  JobCollectionState = "Disabled"
	JobCollectionStateEnabled   JobCollectionState = "Enabled"
	JobCollectionStateSuspended JobCollectionState = "Suspended"
)

// PossibleJobCollectionStateValues returns the possible values for the JobCollectionState const type.
func PossibleJobCollectionStateValues() []JobCollectionState {
	return []JobCollectionState{
		JobCollectionStateDeleted,
		JobCollectionStateDisabled,
		JobCollectionStateEnabled,
		JobCollectionStateSuspended,
	}
}

// JobExecutionStatus - Gets the job execution status.
type JobExecutionStatus string

const (
	JobExecutionStatusCompleted JobExecutionStatus = "Completed"
	JobExecutionStatusFailed    JobExecutionStatus = "Failed"
	JobExecutionStatusPostponed JobExecutionStatus = "Postponed"
)

// PossibleJobExecutionStatusValues returns the possible values for the JobExecutionStatus const type.
func PossibleJobExecutionStatusValues() []JobExecutionStatus {
	return []JobExecutionStatus{
		JobExecutionStatusCompleted,
		JobExecutionStatusFailed,
		JobExecutionStatusPostponed,
	}
}

// JobHistoryActionName - Gets the job history action name.
type JobHistoryActionName string

const (
	JobHistoryActionNameErrorAction JobHistoryActionName = "ErrorAction"
	JobHistoryActionNameMainAction  JobHistoryActionName = "MainAction"
)

// PossibleJobHistoryActionNameValues returns the possible values for the JobHistoryActionName const type.
func PossibleJobHistoryActionNameValues() []JobHistoryActionName {
	return []JobHistoryActionName{
		JobHistoryActionNameErrorAction,
		JobHistoryActionNameMainAction,
	}
}

// JobScheduleDay - Gets or sets the day. Must be one of monday, tuesday, wednesday, thursday, friday, saturday, sunday.
type JobScheduleDay string

const (
	JobScheduleDayFriday    JobScheduleDay = "Friday"
	JobScheduleDayMonday    JobScheduleDay = "Monday"
	JobScheduleDaySaturday  JobScheduleDay = "Saturday"
	JobScheduleDaySunday    JobScheduleDay = "Sunday"
	JobScheduleDayThursday  JobScheduleDay = "Thursday"
	JobScheduleDayTuesday   JobScheduleDay = "Tuesday"
	JobScheduleDayWednesday JobScheduleDay = "Wednesday"
)

// PossibleJobScheduleDayValues returns the possible values for the JobScheduleDay const type.
func PossibleJobScheduleDayValues() []JobScheduleDay {
	return []JobScheduleDay{
		JobScheduleDayFriday,
		JobScheduleDayMonday,
		JobScheduleDaySaturday,
		JobScheduleDaySunday,
		JobScheduleDayThursday,
		JobScheduleDayTuesday,
		JobScheduleDayWednesday,
	}
}

// JobState - Gets or set the job state.
type JobState string

const (
	JobStateCompleted JobState = "Completed"
	JobStateDisabled  JobState = "Disabled"
	JobStateEnabled   JobState = "Enabled"
	JobStateFaulted   JobState = "Faulted"
)

// PossibleJobStateValues returns the possible values for the JobState const type.
func PossibleJobStateValues() []JobState {
	return []JobState{
		JobStateCompleted,
		JobStateDisabled,
		JobStateEnabled,
		JobStateFaulted,
	}
}

// RecurrenceFrequency - Gets or sets the frequency of recurrence (second, minute, hour, day, week, month).
type RecurrenceFrequency string

const (
	RecurrenceFrequencyDay    RecurrenceFrequency = "Day"
	RecurrenceFrequencyHour   RecurrenceFrequency = "Hour"
	RecurrenceFrequencyMinute RecurrenceFrequency = "Minute"
	RecurrenceFrequencyMonth  RecurrenceFrequency = "Month"
	RecurrenceFrequencyWeek   RecurrenceFrequency = "Week"
)

// PossibleRecurrenceFrequencyValues returns the possible values for the RecurrenceFrequency const type.
func PossibleRecurrenceFrequencyValues() []RecurrenceFrequency {
	return []RecurrenceFrequency{
		RecurrenceFrequencyDay,
		RecurrenceFrequencyHour,
		RecurrenceFrequencyMinute,
		RecurrenceFrequencyMonth,
		RecurrenceFrequencyWeek,
	}
}

// RetryType - Gets or sets the retry strategy to be used.
type RetryType string

const (
	RetryTypeFixed RetryType = "Fixed"
	RetryTypeNone  RetryType = "None"
)

// PossibleRetryTypeValues returns the possible values for the RetryType const type.
func PossibleRetryTypeValues() []RetryType {
	return []RetryType{
		RetryTypeFixed,
		RetryTypeNone,
	}
}

// SKUDefinition - Gets or set the SKU.
type SKUDefinition string

const (
	SKUDefinitionFree       SKUDefinition = "Free"
	SKUDefinitionP10Premium SKUDefinition = "P10Premium"
	SKUDefinitionP20Premium SKUDefinition = "P20Premium"
	SKUDefinitionStandard   SKUDefinition = "Standard"
)

// PossibleSKUDefinitionValues returns the possible values for the SKUDefinition const type.
func PossibleSKUDefinitionValues() []SKUDefinition {
	return []SKUDefinition{
		SKUDefinitionFree,
		SKUDefinitionP10Premium,
		SKUDefinitionP20Premium,
		SKUDefinitionStandard,
	}
}

// ServiceBusAuthenticationType - Gets or sets the authentication type.
type ServiceBusAuthenticationType string

const (
	ServiceBusAuthenticationTypeNotSpecified    ServiceBusAuthenticationType = "NotSpecified"
	ServiceBusAuthenticationTypeSharedAccessKey ServiceBusAuthenticationType = "SharedAccessKey"
)

// PossibleServiceBusAuthenticationTypeValues returns the possible values for the ServiceBusAuthenticationType const type.
func PossibleServiceBusAuthenticationTypeValues() []ServiceBusAuthenticationType {
	return []ServiceBusAuthenticationType{
		ServiceBusAuthenticationTypeNotSpecified,
		ServiceBusAuthenticationTypeSharedAccessKey,
	}
}

// ServiceBusTransportType - Gets or sets the transport type.
type ServiceBusTransportType string

const (
	ServiceBusTransportTypeAMQP         ServiceBusTransportType = "AMQP"
	ServiceBusTransportTypeNetMessaging ServiceBusTransportType = "NetMessaging"
	ServiceBusTransportTypeNotSpecified ServiceBusTransportType = "NotSpecified"
)

// PossibleServiceBusTransportTypeValues returns the possible values for the ServiceBusTransportType const type.
func PossibleServiceBusTransportTypeValues() []ServiceBusTransportType {
	return []ServiceBusTransportType{
		ServiceBusTransportTypeAMQP,
		ServiceBusTransportTypeNetMessaging,
		ServiceBusTransportTypeNotSpecified,
	}
}
