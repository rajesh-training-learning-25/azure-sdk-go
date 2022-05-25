package alertsmanagement

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ActionRuleStatus enumerates the values for action rule status.
type ActionRuleStatus string

const (
	// Disabled ...
	Disabled ActionRuleStatus = "Disabled"
	// Enabled ...
	Enabled ActionRuleStatus = "Enabled"
)

// PossibleActionRuleStatusValues returns an array of possible values for the ActionRuleStatus const type.
func PossibleActionRuleStatusValues() []ActionRuleStatus {
	return []ActionRuleStatus{Disabled, Enabled}
}

// AlertModificationEvent enumerates the values for alert modification event.
type AlertModificationEvent string

const (
	// ActionRuleSuppressed ...
	ActionRuleSuppressed AlertModificationEvent = "ActionRuleSuppressed"
	// ActionRuleTriggered ...
	ActionRuleTriggered AlertModificationEvent = "ActionRuleTriggered"
	// ActionsFailed ...
	ActionsFailed AlertModificationEvent = "ActionsFailed"
	// ActionsSuppressed ...
	ActionsSuppressed AlertModificationEvent = "ActionsSuppressed"
	// ActionsTriggered ...
	ActionsTriggered AlertModificationEvent = "ActionsTriggered"
	// AlertCreated ...
	AlertCreated AlertModificationEvent = "AlertCreated"
	// MonitorConditionChange ...
	MonitorConditionChange AlertModificationEvent = "MonitorConditionChange"
	// SeverityChange ...
	SeverityChange AlertModificationEvent = "SeverityChange"
	// StateChange ...
	StateChange AlertModificationEvent = "StateChange"
)

// PossibleAlertModificationEventValues returns an array of possible values for the AlertModificationEvent const type.
func PossibleAlertModificationEventValues() []AlertModificationEvent {
	return []AlertModificationEvent{ActionRuleSuppressed, ActionRuleTriggered, ActionsFailed, ActionsSuppressed, ActionsTriggered, AlertCreated, MonitorConditionChange, SeverityChange, StateChange}
}

// AlertRuleState enumerates the values for alert rule state.
type AlertRuleState string

const (
	// AlertRuleStateDisabled ...
	AlertRuleStateDisabled AlertRuleState = "Disabled"
	// AlertRuleStateEnabled ...
	AlertRuleStateEnabled AlertRuleState = "Enabled"
)

// PossibleAlertRuleStateValues returns an array of possible values for the AlertRuleState const type.
func PossibleAlertRuleStateValues() []AlertRuleState {
	return []AlertRuleState{AlertRuleStateDisabled, AlertRuleStateEnabled}
}

// AlertsSortByFields enumerates the values for alerts sort by fields.
type AlertsSortByFields string

const (
	// AlertsSortByFieldsAlertState ...
	AlertsSortByFieldsAlertState AlertsSortByFields = "alertState"
	// AlertsSortByFieldsLastModifiedDateTime ...
	AlertsSortByFieldsLastModifiedDateTime AlertsSortByFields = "lastModifiedDateTime"
	// AlertsSortByFieldsMonitorCondition ...
	AlertsSortByFieldsMonitorCondition AlertsSortByFields = "monitorCondition"
	// AlertsSortByFieldsName ...
	AlertsSortByFieldsName AlertsSortByFields = "name"
	// AlertsSortByFieldsSeverity ...
	AlertsSortByFieldsSeverity AlertsSortByFields = "severity"
	// AlertsSortByFieldsStartDateTime ...
	AlertsSortByFieldsStartDateTime AlertsSortByFields = "startDateTime"
	// AlertsSortByFieldsTargetResource ...
	AlertsSortByFieldsTargetResource AlertsSortByFields = "targetResource"
	// AlertsSortByFieldsTargetResourceGroup ...
	AlertsSortByFieldsTargetResourceGroup AlertsSortByFields = "targetResourceGroup"
	// AlertsSortByFieldsTargetResourceName ...
	AlertsSortByFieldsTargetResourceName AlertsSortByFields = "targetResourceName"
	// AlertsSortByFieldsTargetResourceType ...
	AlertsSortByFieldsTargetResourceType AlertsSortByFields = "targetResourceType"
)

// PossibleAlertsSortByFieldsValues returns an array of possible values for the AlertsSortByFields const type.
func PossibleAlertsSortByFieldsValues() []AlertsSortByFields {
	return []AlertsSortByFields{AlertsSortByFieldsAlertState, AlertsSortByFieldsLastModifiedDateTime, AlertsSortByFieldsMonitorCondition, AlertsSortByFieldsName, AlertsSortByFieldsSeverity, AlertsSortByFieldsStartDateTime, AlertsSortByFieldsTargetResource, AlertsSortByFieldsTargetResourceGroup, AlertsSortByFieldsTargetResourceName, AlertsSortByFieldsTargetResourceType}
}

// AlertsSummaryGroupByFields enumerates the values for alerts summary group by fields.
type AlertsSummaryGroupByFields string

const (
	// AlertsSummaryGroupByFieldsAlertRule ...
	AlertsSummaryGroupByFieldsAlertRule AlertsSummaryGroupByFields = "alertRule"
	// AlertsSummaryGroupByFieldsAlertState ...
	AlertsSummaryGroupByFieldsAlertState AlertsSummaryGroupByFields = "alertState"
	// AlertsSummaryGroupByFieldsMonitorCondition ...
	AlertsSummaryGroupByFieldsMonitorCondition AlertsSummaryGroupByFields = "monitorCondition"
	// AlertsSummaryGroupByFieldsMonitorService ...
	AlertsSummaryGroupByFieldsMonitorService AlertsSummaryGroupByFields = "monitorService"
	// AlertsSummaryGroupByFieldsSeverity ...
	AlertsSummaryGroupByFieldsSeverity AlertsSummaryGroupByFields = "severity"
	// AlertsSummaryGroupByFieldsSignalType ...
	AlertsSummaryGroupByFieldsSignalType AlertsSummaryGroupByFields = "signalType"
)

// PossibleAlertsSummaryGroupByFieldsValues returns an array of possible values for the AlertsSummaryGroupByFields const type.
func PossibleAlertsSummaryGroupByFieldsValues() []AlertsSummaryGroupByFields {
	return []AlertsSummaryGroupByFields{AlertsSummaryGroupByFieldsAlertRule, AlertsSummaryGroupByFieldsAlertState, AlertsSummaryGroupByFieldsMonitorCondition, AlertsSummaryGroupByFieldsMonitorService, AlertsSummaryGroupByFieldsSeverity, AlertsSummaryGroupByFieldsSignalType}
}

// AlertState enumerates the values for alert state.
type AlertState string

const (
	// AlertStateAcknowledged ...
	AlertStateAcknowledged AlertState = "Acknowledged"
	// AlertStateClosed ...
	AlertStateClosed AlertState = "Closed"
	// AlertStateNew ...
	AlertStateNew AlertState = "New"
)

// PossibleAlertStateValues returns an array of possible values for the AlertState const type.
func PossibleAlertStateValues() []AlertState {
	return []AlertState{AlertStateAcknowledged, AlertStateClosed, AlertStateNew}
}

// MetadataIdentifier enumerates the values for metadata identifier.
type MetadataIdentifier string

const (
	// MetadataIdentifierAlertsMetaDataProperties ...
	MetadataIdentifierAlertsMetaDataProperties MetadataIdentifier = "alertsMetaDataProperties"
	// MetadataIdentifierMonitorServiceList ...
	MetadataIdentifierMonitorServiceList MetadataIdentifier = "MonitorServiceList"
)

// PossibleMetadataIdentifierValues returns an array of possible values for the MetadataIdentifier const type.
func PossibleMetadataIdentifierValues() []MetadataIdentifier {
	return []MetadataIdentifier{MetadataIdentifierAlertsMetaDataProperties, MetadataIdentifierMonitorServiceList}
}

// MonitorCondition enumerates the values for monitor condition.
type MonitorCondition string

const (
	// Fired ...
	Fired MonitorCondition = "Fired"
	// Resolved ...
	Resolved MonitorCondition = "Resolved"
)

// PossibleMonitorConditionValues returns an array of possible values for the MonitorCondition const type.
func PossibleMonitorConditionValues() []MonitorCondition {
	return []MonitorCondition{Fired, Resolved}
}

// MonitorService enumerates the values for monitor service.
type MonitorService string

const (
	// ActivityLogAdministrative ...
	ActivityLogAdministrative MonitorService = "ActivityLog Administrative"
	// ActivityLogAutoscale ...
	ActivityLogAutoscale MonitorService = "ActivityLog Autoscale"
	// ActivityLogPolicy ...
	ActivityLogPolicy MonitorService = "ActivityLog Policy"
	// ActivityLogRecommendation ...
	ActivityLogRecommendation MonitorService = "ActivityLog Recommendation"
	// ActivityLogSecurity ...
	ActivityLogSecurity MonitorService = "ActivityLog Security"
	// ApplicationInsights ...
	ApplicationInsights MonitorService = "Application Insights"
	// LogAnalytics ...
	LogAnalytics MonitorService = "Log Analytics"
	// Nagios ...
	Nagios MonitorService = "Nagios"
	// Platform ...
	Platform MonitorService = "Platform"
	// SCOM ...
	SCOM MonitorService = "SCOM"
	// ServiceHealth ...
	ServiceHealth MonitorService = "ServiceHealth"
	// SmartDetector ...
	SmartDetector MonitorService = "SmartDetector"
	// VMInsights ...
	VMInsights MonitorService = "VM Insights"
	// Zabbix ...
	Zabbix MonitorService = "Zabbix"
)

// PossibleMonitorServiceValues returns an array of possible values for the MonitorService const type.
func PossibleMonitorServiceValues() []MonitorService {
	return []MonitorService{ActivityLogAdministrative, ActivityLogAutoscale, ActivityLogPolicy, ActivityLogRecommendation, ActivityLogSecurity, ApplicationInsights, LogAnalytics, Nagios, Platform, SCOM, ServiceHealth, SmartDetector, VMInsights, Zabbix}
}

// Operator enumerates the values for operator.
type Operator string

const (
	// Contains ...
	Contains Operator = "Contains"
	// DoesNotContain ...
	DoesNotContain Operator = "DoesNotContain"
	// Equals ...
	Equals Operator = "Equals"
	// NotEquals ...
	NotEquals Operator = "NotEquals"
)

// PossibleOperatorValues returns an array of possible values for the Operator const type.
func PossibleOperatorValues() []Operator {
	return []Operator{Contains, DoesNotContain, Equals, NotEquals}
}

// ScopeType enumerates the values for scope type.
type ScopeType string

const (
	// ScopeTypeResource ...
	ScopeTypeResource ScopeType = "Resource"
	// ScopeTypeResourceGroup ...
	ScopeTypeResourceGroup ScopeType = "ResourceGroup"
	// ScopeTypeSubscription ...
	ScopeTypeSubscription ScopeType = "Subscription"
)

// PossibleScopeTypeValues returns an array of possible values for the ScopeType const type.
func PossibleScopeTypeValues() []ScopeType {
	return []ScopeType{ScopeTypeResource, ScopeTypeResourceGroup, ScopeTypeSubscription}
}

// Severity enumerates the values for severity.
type Severity string

const (
	// Sev0 ...
	Sev0 Severity = "Sev0"
	// Sev1 ...
	Sev1 Severity = "Sev1"
	// Sev2 ...
	Sev2 Severity = "Sev2"
	// Sev3 ...
	Sev3 Severity = "Sev3"
	// Sev4 ...
	Sev4 Severity = "Sev4"
)

// PossibleSeverityValues returns an array of possible values for the Severity const type.
func PossibleSeverityValues() []Severity {
	return []Severity{Sev0, Sev1, Sev2, Sev3, Sev4}
}

// SignalType enumerates the values for signal type.
type SignalType string

const (
	// Log ...
	Log SignalType = "Log"
	// Metric ...
	Metric SignalType = "Metric"
	// Unknown ...
	Unknown SignalType = "Unknown"
)

// PossibleSignalTypeValues returns an array of possible values for the SignalType const type.
func PossibleSignalTypeValues() []SignalType {
	return []SignalType{Log, Metric, Unknown}
}

// SmartGroupModificationEvent enumerates the values for smart group modification event.
type SmartGroupModificationEvent string

const (
	// SmartGroupModificationEventAlertAdded ...
	SmartGroupModificationEventAlertAdded SmartGroupModificationEvent = "AlertAdded"
	// SmartGroupModificationEventAlertRemoved ...
	SmartGroupModificationEventAlertRemoved SmartGroupModificationEvent = "AlertRemoved"
	// SmartGroupModificationEventSmartGroupCreated ...
	SmartGroupModificationEventSmartGroupCreated SmartGroupModificationEvent = "SmartGroupCreated"
	// SmartGroupModificationEventStateChange ...
	SmartGroupModificationEventStateChange SmartGroupModificationEvent = "StateChange"
)

// PossibleSmartGroupModificationEventValues returns an array of possible values for the SmartGroupModificationEvent const type.
func PossibleSmartGroupModificationEventValues() []SmartGroupModificationEvent {
	return []SmartGroupModificationEvent{SmartGroupModificationEventAlertAdded, SmartGroupModificationEventAlertRemoved, SmartGroupModificationEventSmartGroupCreated, SmartGroupModificationEventStateChange}
}

// SmartGroupsSortByFields enumerates the values for smart groups sort by fields.
type SmartGroupsSortByFields string

const (
	// SmartGroupsSortByFieldsAlertsCount ...
	SmartGroupsSortByFieldsAlertsCount SmartGroupsSortByFields = "alertsCount"
	// SmartGroupsSortByFieldsLastModifiedDateTime ...
	SmartGroupsSortByFieldsLastModifiedDateTime SmartGroupsSortByFields = "lastModifiedDateTime"
	// SmartGroupsSortByFieldsSeverity ...
	SmartGroupsSortByFieldsSeverity SmartGroupsSortByFields = "severity"
	// SmartGroupsSortByFieldsStartDateTime ...
	SmartGroupsSortByFieldsStartDateTime SmartGroupsSortByFields = "startDateTime"
	// SmartGroupsSortByFieldsState ...
	SmartGroupsSortByFieldsState SmartGroupsSortByFields = "state"
)

// PossibleSmartGroupsSortByFieldsValues returns an array of possible values for the SmartGroupsSortByFields const type.
func PossibleSmartGroupsSortByFieldsValues() []SmartGroupsSortByFields {
	return []SmartGroupsSortByFields{SmartGroupsSortByFieldsAlertsCount, SmartGroupsSortByFieldsLastModifiedDateTime, SmartGroupsSortByFieldsSeverity, SmartGroupsSortByFieldsStartDateTime, SmartGroupsSortByFieldsState}
}

// SortOrder enumerates the values for sort order.
type SortOrder string

const (
	// Asc ...
	Asc SortOrder = "asc"
	// Desc ...
	Desc SortOrder = "desc"
)

// PossibleSortOrderValues returns an array of possible values for the SortOrder const type.
func PossibleSortOrderValues() []SortOrder {
	return []SortOrder{Asc, Desc}
}

// State enumerates the values for state.
type State string

const (
	// StateAcknowledged ...
	StateAcknowledged State = "Acknowledged"
	// StateClosed ...
	StateClosed State = "Closed"
	// StateNew ...
	StateNew State = "New"
)

// PossibleStateValues returns an array of possible values for the State const type.
func PossibleStateValues() []State {
	return []State{StateAcknowledged, StateClosed, StateNew}
}

// SuppressionType enumerates the values for suppression type.
type SuppressionType string

const (
	// Always ...
	Always SuppressionType = "Always"
	// Daily ...
	Daily SuppressionType = "Daily"
	// Monthly ...
	Monthly SuppressionType = "Monthly"
	// Once ...
	Once SuppressionType = "Once"
	// Weekly ...
	Weekly SuppressionType = "Weekly"
)

// PossibleSuppressionTypeValues returns an array of possible values for the SuppressionType const type.
func PossibleSuppressionTypeValues() []SuppressionType {
	return []SuppressionType{Always, Daily, Monthly, Once, Weekly}
}

// TimeRange enumerates the values for time range.
type TimeRange string

const (
	// Oned ...
	Oned TimeRange = "1d"
	// Oneh ...
	Oneh TimeRange = "1h"
	// Sevend ...
	Sevend TimeRange = "7d"
	// ThreeZerod ...
	ThreeZerod TimeRange = "30d"
)

// PossibleTimeRangeValues returns an array of possible values for the TimeRange const type.
func PossibleTimeRangeValues() []TimeRange {
	return []TimeRange{Oned, Oneh, Sevend, ThreeZerod}
}

// Type enumerates the values for type.
type Type string

const (
	// TypeActionGroup ...
	TypeActionGroup Type = "ActionGroup"
	// TypeActionRuleProperties ...
	TypeActionRuleProperties Type = "ActionRuleProperties"
	// TypeDiagnostics ...
	TypeDiagnostics Type = "Diagnostics"
	// TypeSuppression ...
	TypeSuppression Type = "Suppression"
)

// PossibleTypeValues returns an array of possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{TypeActionGroup, TypeActionRuleProperties, TypeDiagnostics, TypeSuppression}
}
