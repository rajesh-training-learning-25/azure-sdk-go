// +build go1.9

// Copyright 2018 Microsoft Corporation
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

package policyinsights

import original "github.com/Azure/azure-sdk-for-go/services/policyinsights/mgmt/2017-12-12-preview/policyinsights"

type PolicyStatesClient = original.PolicyStatesClient
type PolicyStatesResource = original.PolicyStatesResource

const (
	Default PolicyStatesResource = original.Default
	Latest  PolicyStatesResource = original.Latest
)

type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationsListResults = original.OperationsListResults
type PolicyAssignmentSummary = original.PolicyAssignmentSummary
type PolicyDefinitionSummary = original.PolicyDefinitionSummary
type PolicyEvent = original.PolicyEvent
type PolicyEventsQueryResults = original.PolicyEventsQueryResults
type PolicyState = original.PolicyState
type PolicyStatesQueryResults = original.PolicyStatesQueryResults
type QueryFailure = original.QueryFailure
type QueryFailureError = original.QueryFailureError
type String = original.String
type SummarizeResults = original.SummarizeResults
type Summary = original.Summary
type SummaryResults = original.SummaryResults
type PolicyEventsClient = original.PolicyEventsClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type OperationsClient = original.OperationsClient

func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
func New() BaseClient {
	return original.New()
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
func NewOperationsClient() OperationsClient {
	return original.NewOperationsClient()
}
func NewOperationsClientWithBaseURI(baseURI string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI)
}
func NewPolicyStatesClient() PolicyStatesClient {
	return original.NewPolicyStatesClient()
}
func NewPolicyStatesClientWithBaseURI(baseURI string) PolicyStatesClient {
	return original.NewPolicyStatesClientWithBaseURI(baseURI)
}
func PossiblePolicyStatesResourceValues() []PolicyStatesResource {
	return original.PossiblePolicyStatesResourceValues()
}
func NewPolicyEventsClient() PolicyEventsClient {
	return original.NewPolicyEventsClient()
}
func NewPolicyEventsClientWithBaseURI(baseURI string) PolicyEventsClient {
	return original.NewPolicyEventsClientWithBaseURI(baseURI)
}
