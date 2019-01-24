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

package securityinsight

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/securityinsight/mgmt/2017-08-01-preview/securityinsight"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AlertTriggerOperator = original.AlertTriggerOperator

const (
	Eq AlertTriggerOperator = original.Eq
	Gt AlertTriggerOperator = original.Gt
	Lt AlertTriggerOperator = original.Lt
	Ne AlertTriggerOperator = original.Ne
)

type Severity = original.Severity

const (
	High          Severity = original.High
	Informational Severity = original.Informational
	Low           Severity = original.Low
	Medium        Severity = original.Medium
)

type BaseClient = original.BaseClient
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationsClient = original.OperationsClient
type OperationsList = original.OperationsList
type OperationsListIterator = original.OperationsListIterator
type OperationsListPage = original.OperationsListPage
type Resource = original.Resource
type ScheduledAlertRule = original.ScheduledAlertRule
type ScheduledAlertRuleProperties = original.ScheduledAlertRuleProperties
type ScheduledAlertRulesClient = original.ScheduledAlertRulesClient
type ScheduledAlertRulesList = original.ScheduledAlertRulesList
type ScheduledAlertRulesListIterator = original.ScheduledAlertRulesListIterator
type ScheduledAlertRulesListPage = original.ScheduledAlertRulesListPage

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsListIterator(page OperationsListPage) OperationsListIterator {
	return original.NewOperationsListIterator(page)
}
func NewOperationsListPage(getNextPage func(context.Context, OperationsList) (OperationsList, error)) OperationsListPage {
	return original.NewOperationsListPage(getNextPage)
}
func NewScheduledAlertRulesClient(subscriptionID string) ScheduledAlertRulesClient {
	return original.NewScheduledAlertRulesClient(subscriptionID)
}
func NewScheduledAlertRulesClientWithBaseURI(baseURI string, subscriptionID string) ScheduledAlertRulesClient {
	return original.NewScheduledAlertRulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewScheduledAlertRulesListIterator(page ScheduledAlertRulesListPage) ScheduledAlertRulesListIterator {
	return original.NewScheduledAlertRulesListIterator(page)
}
func NewScheduledAlertRulesListPage(getNextPage func(context.Context, ScheduledAlertRulesList) (ScheduledAlertRulesList, error)) ScheduledAlertRulesListPage {
	return original.NewScheduledAlertRulesListPage(getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAlertTriggerOperatorValues() []AlertTriggerOperator {
	return original.PossibleAlertTriggerOperatorValues()
}
func PossibleSeverityValues() []Severity {
	return original.PossibleSeverityValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
