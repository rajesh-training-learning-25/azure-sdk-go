//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconsumption

const (
	module  = "armconsumption"
	version = "v0.1.1"
)

// BillingFrequency - The billing frequency.
type BillingFrequency string

const (
	BillingFrequencyMonth   BillingFrequency = "Month"
	BillingFrequencyQuarter BillingFrequency = "Quarter"
	BillingFrequencyYear    BillingFrequency = "Year"
)

// PossibleBillingFrequencyValues returns the possible values for the BillingFrequency const type.
func PossibleBillingFrequencyValues() []BillingFrequency {
	return []BillingFrequency{
		BillingFrequencyMonth,
		BillingFrequencyQuarter,
		BillingFrequencyYear,
	}
}

// ToPtr returns a *BillingFrequency pointing to the current value.
func (c BillingFrequency) ToPtr() *BillingFrequency {
	return &c
}

// BudgetOperatorType - The operator to use for comparison.
type BudgetOperatorType string

const (
	BudgetOperatorTypeIn BudgetOperatorType = "In"
)

// PossibleBudgetOperatorTypeValues returns the possible values for the BudgetOperatorType const type.
func PossibleBudgetOperatorTypeValues() []BudgetOperatorType {
	return []BudgetOperatorType{
		BudgetOperatorTypeIn,
	}
}

// ToPtr returns a *BudgetOperatorType pointing to the current value.
func (c BudgetOperatorType) ToPtr() *BudgetOperatorType {
	return &c
}

// CategoryType - The category of the budget, whether the budget tracks cost or usage.
type CategoryType string

const (
	CategoryTypeCost CategoryType = "Cost"
)

// PossibleCategoryTypeValues returns the possible values for the CategoryType const type.
func PossibleCategoryTypeValues() []CategoryType {
	return []CategoryType{
		CategoryTypeCost,
	}
}

// ToPtr returns a *CategoryType pointing to the current value.
func (c CategoryType) ToPtr() *CategoryType {
	return &c
}

// ChargeSummaryKind - Specifies the kind of charge summary.
type ChargeSummaryKind string

const (
	ChargeSummaryKindLegacy ChargeSummaryKind = "legacy"
	ChargeSummaryKindModern ChargeSummaryKind = "modern"
)

// PossibleChargeSummaryKindValues returns the possible values for the ChargeSummaryKind const type.
func PossibleChargeSummaryKindValues() []ChargeSummaryKind {
	return []ChargeSummaryKind{
		ChargeSummaryKindLegacy,
		ChargeSummaryKindModern,
	}
}

// ToPtr returns a *ChargeSummaryKind pointing to the current value.
func (c ChargeSummaryKind) ToPtr() *ChargeSummaryKind {
	return &c
}

type Datagrain string

const (
	// DatagrainDailyGrain - Daily grain of data
	DatagrainDailyGrain Datagrain = "daily"
	// DatagrainMonthlyGrain - Monthly grain of data
	DatagrainMonthlyGrain Datagrain = "monthly"
)

// PossibleDatagrainValues returns the possible values for the Datagrain const type.
func PossibleDatagrainValues() []Datagrain {
	return []Datagrain{
		DatagrainDailyGrain,
		DatagrainMonthlyGrain,
	}
}

// ToPtr returns a *Datagrain pointing to the current value.
func (c Datagrain) ToPtr() *Datagrain {
	return &c
}

// EventType - Identifies the type of the event.
type EventType string

const (
	EventTypeNewCredit            EventType = "NewCredit"
	EventTypePendingAdjustments   EventType = "PendingAdjustments"
	EventTypePendingCharges       EventType = "PendingCharges"
	EventTypePendingExpiredCredit EventType = "PendingExpiredCredit"
	EventTypePendingNewCredit     EventType = "PendingNewCredit"
	EventTypeSettledCharges       EventType = "SettledCharges"
	EventTypeUnKnown              EventType = "UnKnown"
)

// PossibleEventTypeValues returns the possible values for the EventType const type.
func PossibleEventTypeValues() []EventType {
	return []EventType{
		EventTypeNewCredit,
		EventTypePendingAdjustments,
		EventTypePendingCharges,
		EventTypePendingExpiredCredit,
		EventTypePendingNewCredit,
		EventTypeSettledCharges,
		EventTypeUnKnown,
	}
}

// ToPtr returns a *EventType pointing to the current value.
func (c EventType) ToPtr() *EventType {
	return &c
}

type LookBackPeriod string

const (
	// LookBackPeriodLast07Days - Use 7 days of data for recommendations
	LookBackPeriodLast07Days LookBackPeriod = "Last7Days"
	// LookBackPeriodLast30Days - Use 30 days of data for recommendations
	LookBackPeriodLast30Days LookBackPeriod = "Last30Days"
	// LookBackPeriodLast60Days - Use 60 days of data for recommendations
	LookBackPeriodLast60Days LookBackPeriod = "Last60Days"
)

// PossibleLookBackPeriodValues returns the possible values for the LookBackPeriod const type.
func PossibleLookBackPeriodValues() []LookBackPeriod {
	return []LookBackPeriod{
		LookBackPeriodLast07Days,
		LookBackPeriodLast30Days,
		LookBackPeriodLast60Days,
	}
}

// ToPtr returns a *LookBackPeriod pointing to the current value.
func (c LookBackPeriod) ToPtr() *LookBackPeriod {
	return &c
}

// LotSource - The source of the lot.
type LotSource string

const (
	LotSourceConsumptionCommitment LotSource = "ConsumptionCommitment"
	LotSourcePromotionalCredit     LotSource = "PromotionalCredit"
	LotSourcePurchasedCredit       LotSource = "PurchasedCredit"
)

// PossibleLotSourceValues returns the possible values for the LotSource const type.
func PossibleLotSourceValues() []LotSource {
	return []LotSource{
		LotSourceConsumptionCommitment,
		LotSourcePromotionalCredit,
		LotSourcePurchasedCredit,
	}
}

// ToPtr returns a *LotSource pointing to the current value.
func (c LotSource) ToPtr() *LotSource {
	return &c
}

type Metrictype string

const (
	// MetrictypeActualCostMetricType - Actual cost data.
	MetrictypeActualCostMetricType Metrictype = "actualcost"
	// MetrictypeAmortizedCostMetricType - Amortized cost data.
	MetrictypeAmortizedCostMetricType Metrictype = "amortizedcost"
	// MetrictypeUsageMetricType - Usage data.
	MetrictypeUsageMetricType Metrictype = "usage"
)

// PossibleMetrictypeValues returns the possible values for the Metrictype const type.
func PossibleMetrictypeValues() []Metrictype {
	return []Metrictype{
		MetrictypeActualCostMetricType,
		MetrictypeAmortizedCostMetricType,
		MetrictypeUsageMetricType,
	}
}

// ToPtr returns a *Metrictype pointing to the current value.
func (c Metrictype) ToPtr() *Metrictype {
	return &c
}

// OperatorType - The comparison operator.
type OperatorType string

const (
	OperatorTypeEqualTo              OperatorType = "EqualTo"
	OperatorTypeGreaterThan          OperatorType = "GreaterThan"
	OperatorTypeGreaterThanOrEqualTo OperatorType = "GreaterThanOrEqualTo"
)

// PossibleOperatorTypeValues returns the possible values for the OperatorType const type.
func PossibleOperatorTypeValues() []OperatorType {
	return []OperatorType{
		OperatorTypeEqualTo,
		OperatorTypeGreaterThan,
		OperatorTypeGreaterThanOrEqualTo,
	}
}

// ToPtr returns a *OperatorType pointing to the current value.
func (c OperatorType) ToPtr() *OperatorType {
	return &c
}

// PricingModelType - Identifier that indicates how the meter is priced.
type PricingModelType string

const (
	PricingModelTypeOnDemand    PricingModelType = "On Demand"
	PricingModelTypeReservation PricingModelType = "Reservation"
	PricingModelTypeSpot        PricingModelType = "Spot"
)

// PossiblePricingModelTypeValues returns the possible values for the PricingModelType const type.
func PossiblePricingModelTypeValues() []PricingModelType {
	return []PricingModelType{
		PricingModelTypeOnDemand,
		PricingModelTypeReservation,
		PricingModelTypeSpot,
	}
}

// ToPtr returns a *PricingModelType pointing to the current value.
func (c PricingModelType) ToPtr() *PricingModelType {
	return &c
}

// ReservationRecommendationKind - Specifies the kind of reservation recommendation.
type ReservationRecommendationKind string

const (
	ReservationRecommendationKindLegacy ReservationRecommendationKind = "legacy"
	ReservationRecommendationKindModern ReservationRecommendationKind = "modern"
)

// PossibleReservationRecommendationKindValues returns the possible values for the ReservationRecommendationKind const type.
func PossibleReservationRecommendationKindValues() []ReservationRecommendationKind {
	return []ReservationRecommendationKind{
		ReservationRecommendationKindLegacy,
		ReservationRecommendationKindModern,
	}
}

// ToPtr returns a *ReservationRecommendationKind pointing to the current value.
func (c ReservationRecommendationKind) ToPtr() *ReservationRecommendationKind {
	return &c
}

type Scope string

const (
	ScopeShared Scope = "Shared"
	ScopeSingle Scope = "Single"
)

// PossibleScopeValues returns the possible values for the Scope const type.
func PossibleScopeValues() []Scope {
	return []Scope{
		ScopeShared,
		ScopeSingle,
	}
}

// ToPtr returns a *Scope pointing to the current value.
func (c Scope) ToPtr() *Scope {
	return &c
}

// Status - The status of the lot.
type Status string

const (
	StatusActive   Status = "Active"
	StatusCanceled Status = "Canceled"
	StatusComplete Status = "Complete"
	StatusExpired  Status = "Expired"
	StatusInactive Status = "Inactive"
	StatusNone     Status = "None"
)

// PossibleStatusValues returns the possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{
		StatusActive,
		StatusCanceled,
		StatusComplete,
		StatusExpired,
		StatusInactive,
		StatusNone,
	}
}

// ToPtr returns a *Status pointing to the current value.
func (c Status) ToPtr() *Status {
	return &c
}

type Term string

const (
	// TermP1Y - 1 year reservation term
	TermP1Y Term = "P1Y"
	// TermP3Y - 3 year reservation term
	TermP3Y Term = "P3Y"
)

// PossibleTermValues returns the possible values for the Term const type.
func PossibleTermValues() []Term {
	return []Term{
		TermP1Y,
		TermP3Y,
	}
}

// ToPtr returns a *Term pointing to the current value.
func (c Term) ToPtr() *Term {
	return &c
}

// ThresholdType - The type of threshold
type ThresholdType string

const (
	ThresholdTypeActual ThresholdType = "Actual"
)

// PossibleThresholdTypeValues returns the possible values for the ThresholdType const type.
func PossibleThresholdTypeValues() []ThresholdType {
	return []ThresholdType{
		ThresholdTypeActual,
	}
}

// ToPtr returns a *ThresholdType pointing to the current value.
func (c ThresholdType) ToPtr() *ThresholdType {
	return &c
}

// TimeGrainType - The time covered by a budget. Tracking of the amount will be reset based on the time grain. BillingMonth, BillingQuarter, and BillingAnnual
// are only supported by WD customers
type TimeGrainType string

const (
	TimeGrainTypeAnnually       TimeGrainType = "Annually"
	TimeGrainTypeBillingAnnual  TimeGrainType = "BillingAnnual"
	TimeGrainTypeBillingMonth   TimeGrainType = "BillingMonth"
	TimeGrainTypeBillingQuarter TimeGrainType = "BillingQuarter"
	TimeGrainTypeMonthly        TimeGrainType = "Monthly"
	TimeGrainTypeQuarterly      TimeGrainType = "Quarterly"
)

// PossibleTimeGrainTypeValues returns the possible values for the TimeGrainType const type.
func PossibleTimeGrainTypeValues() []TimeGrainType {
	return []TimeGrainType{
		TimeGrainTypeAnnually,
		TimeGrainTypeBillingAnnual,
		TimeGrainTypeBillingMonth,
		TimeGrainTypeBillingQuarter,
		TimeGrainTypeMonthly,
		TimeGrainTypeQuarterly,
	}
}

// ToPtr returns a *TimeGrainType pointing to the current value.
func (c TimeGrainType) ToPtr() *TimeGrainType {
	return &c
}

// UsageDetailsKind - Specifies the kind of usage details.
type UsageDetailsKind string

const (
	UsageDetailsKindLegacy UsageDetailsKind = "legacy"
	UsageDetailsKindModern UsageDetailsKind = "modern"
)

// PossibleUsageDetailsKindValues returns the possible values for the UsageDetailsKind const type.
func PossibleUsageDetailsKindValues() []UsageDetailsKind {
	return []UsageDetailsKind{
		UsageDetailsKindLegacy,
		UsageDetailsKindModern,
	}
}

// ToPtr returns a *UsageDetailsKind pointing to the current value.
func (c UsageDetailsKind) ToPtr() *UsageDetailsKind {
	return &c
}
