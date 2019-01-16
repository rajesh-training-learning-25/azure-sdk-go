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

package cognitiveservices

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/mgmt/2017-04-18/cognitiveservices"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type KeyName = original.KeyName

const (
	Key1 KeyName = original.Key1
	Key2 KeyName = original.Key2
)

type Kind = original.Kind

const (
	BingAutosuggestv7      Kind = original.BingAutosuggestv7
	BingCustomSearch       Kind = original.BingCustomSearch
	BingSearchv7           Kind = original.BingSearchv7
	BingSpeech             Kind = original.BingSpeech
	BingSpellCheckv7       Kind = original.BingSpellCheckv7
	ComputerVision         Kind = original.ComputerVision
	ContentModerator       Kind = original.ContentModerator
	CustomSpeech           Kind = original.CustomSpeech
	CustomVisionPrediction Kind = original.CustomVisionPrediction
	CustomVisionTraining   Kind = original.CustomVisionTraining
	Emotion                Kind = original.Emotion
	Face                   Kind = original.Face
	LUIS                   Kind = original.LUIS
	QnAMaker               Kind = original.QnAMaker
	SpeakerRecognition     Kind = original.SpeakerRecognition
	SpeechTranslation      Kind = original.SpeechTranslation
	TextAnalytics          Kind = original.TextAnalytics
	TextTranslation        Kind = original.TextTranslation
	WebLM                  Kind = original.WebLM
)

type ProvisioningState = original.ProvisioningState

const (
	Creating     ProvisioningState = original.Creating
	Deleting     ProvisioningState = original.Deleting
	Failed       ProvisioningState = original.Failed
	Moving       ProvisioningState = original.Moving
	ResolvingDNS ProvisioningState = original.ResolvingDNS
	Succeeded    ProvisioningState = original.Succeeded
)

type QuotaUsageStatus = original.QuotaUsageStatus

const (
	Blocked   QuotaUsageStatus = original.Blocked
	Included  QuotaUsageStatus = original.Included
	InOverage QuotaUsageStatus = original.InOverage
	Unknown   QuotaUsageStatus = original.Unknown
)

type ResourceSkuRestrictionsReasonCode = original.ResourceSkuRestrictionsReasonCode

const (
	NotAvailableForSubscription ResourceSkuRestrictionsReasonCode = original.NotAvailableForSubscription
	QuotaID                     ResourceSkuRestrictionsReasonCode = original.QuotaID
)

type ResourceSkuRestrictionsType = original.ResourceSkuRestrictionsType

const (
	Location ResourceSkuRestrictionsType = original.Location
	Zone     ResourceSkuRestrictionsType = original.Zone
)

type SkuName = original.SkuName

const (
	F0 SkuName = original.F0
	P0 SkuName = original.P0
	P1 SkuName = original.P1
	P2 SkuName = original.P2
	S0 SkuName = original.S0
	S1 SkuName = original.S1
	S2 SkuName = original.S2
	S3 SkuName = original.S3
	S4 SkuName = original.S4
	S5 SkuName = original.S5
	S6 SkuName = original.S6
)

type SkuTier = original.SkuTier

const (
	Free     SkuTier = original.Free
	Premium  SkuTier = original.Premium
	Standard SkuTier = original.Standard
)

type UnitType = original.UnitType

const (
	Bytes          UnitType = original.Bytes
	BytesPerSecond UnitType = original.BytesPerSecond
	Count          UnitType = original.Count
	CountPerSecond UnitType = original.CountPerSecond
	Milliseconds   UnitType = original.Milliseconds
	Percent        UnitType = original.Percent
	Seconds        UnitType = original.Seconds
)

type Account = original.Account
type AccountCreateParameters = original.AccountCreateParameters
type AccountEnumerateSkusResult = original.AccountEnumerateSkusResult
type AccountKeys = original.AccountKeys
type AccountListResult = original.AccountListResult
type AccountListResultIterator = original.AccountListResultIterator
type AccountListResultPage = original.AccountListResultPage
type AccountProperties = original.AccountProperties
type AccountUpdateParameters = original.AccountUpdateParameters
type AccountsClient = original.AccountsClient
type BaseClient = original.BaseClient
type CheckSkuAvailabilityClient = original.CheckSkuAvailabilityClient
type CheckSkuAvailabilityParameter = original.CheckSkuAvailabilityParameter
type CheckSkuAvailabilityResult = original.CheckSkuAvailabilityResult
type CheckSkuAvailabilityResultList = original.CheckSkuAvailabilityResultList
type Error = original.Error
type ErrorBody = original.ErrorBody
type MetricName = original.MetricName
type OperationDisplayInfo = original.OperationDisplayInfo
type OperationEntity = original.OperationEntity
type OperationEntityListResult = original.OperationEntityListResult
type OperationEntityListResultIterator = original.OperationEntityListResultIterator
type OperationEntityListResultPage = original.OperationEntityListResultPage
type OperationsClient = original.OperationsClient
type RegenerateKeyParameters = original.RegenerateKeyParameters
type ResourceAndSku = original.ResourceAndSku
type ResourceSku = original.ResourceSku
type ResourceSkuRestrictionInfo = original.ResourceSkuRestrictionInfo
type ResourceSkuRestrictions = original.ResourceSkuRestrictions
type ResourceSkusClient = original.ResourceSkusClient
type ResourceSkusResult = original.ResourceSkusResult
type ResourceSkusResultIterator = original.ResourceSkusResultIterator
type ResourceSkusResultPage = original.ResourceSkusResultPage
type Sku = original.Sku
type Usage = original.Usage
type UsagesResult = original.UsagesResult

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAccountListResultIterator(page AccountListResultPage) AccountListResultIterator {
	return original.NewAccountListResultIterator(page)
}
func NewAccountListResultPage(getNextPage func(context.Context, AccountListResult) (AccountListResult, error)) AccountListResultPage {
	return original.NewAccountListResultPage(getNextPage)
}
func NewAccountsClient(subscriptionID string) AccountsClient {
	return original.NewAccountsClient(subscriptionID)
}
func NewAccountsClientWithBaseURI(baseURI string, subscriptionID string) AccountsClient {
	return original.NewAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewCheckSkuAvailabilityClient(subscriptionID string) CheckSkuAvailabilityClient {
	return original.NewCheckSkuAvailabilityClient(subscriptionID)
}
func NewCheckSkuAvailabilityClientWithBaseURI(baseURI string, subscriptionID string) CheckSkuAvailabilityClient {
	return original.NewCheckSkuAvailabilityClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationEntityListResultIterator(page OperationEntityListResultPage) OperationEntityListResultIterator {
	return original.NewOperationEntityListResultIterator(page)
}
func NewOperationEntityListResultPage(getNextPage func(context.Context, OperationEntityListResult) (OperationEntityListResult, error)) OperationEntityListResultPage {
	return original.NewOperationEntityListResultPage(getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
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
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleKeyNameValues() []KeyName {
	return original.PossibleKeyNameValues()
}
func PossibleKindValues() []Kind {
	return original.PossibleKindValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleQuotaUsageStatusValues() []QuotaUsageStatus {
	return original.PossibleQuotaUsageStatusValues()
}
func PossibleResourceSkuRestrictionsReasonCodeValues() []ResourceSkuRestrictionsReasonCode {
	return original.PossibleResourceSkuRestrictionsReasonCodeValues()
}
func PossibleResourceSkuRestrictionsTypeValues() []ResourceSkuRestrictionsType {
	return original.PossibleResourceSkuRestrictionsTypeValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}
func PossibleUnitTypeValues() []UnitType {
	return original.PossibleUnitTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
