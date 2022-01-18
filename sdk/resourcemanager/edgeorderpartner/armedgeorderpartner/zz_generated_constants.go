//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armedgeorderpartner

const (
	moduleName    = "armedgeorderpartner"
	moduleVersion = "v0.2.0"
)

// ActionType - Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

// ToPtr returns a *ActionType pointing to the current value.
func (c ActionType) ToPtr() *ActionType {
	return &c
}

// ManageLinkOperation - Operation to be performed - Link, Unlink, Relink
type ManageLinkOperation string

const (
	// ManageLinkOperationLink - Link.
	ManageLinkOperationLink ManageLinkOperation = "Link"
	// ManageLinkOperationUnlink - Unlink.
	ManageLinkOperationUnlink ManageLinkOperation = "Unlink"
	// ManageLinkOperationRelink - Relink.
	ManageLinkOperationRelink ManageLinkOperation = "Relink"
)

// PossibleManageLinkOperationValues returns the possible values for the ManageLinkOperation const type.
func PossibleManageLinkOperationValues() []ManageLinkOperation {
	return []ManageLinkOperation{
		ManageLinkOperationLink,
		ManageLinkOperationUnlink,
		ManageLinkOperationRelink,
	}
}

// ToPtr returns a *ManageLinkOperation pointing to the current value.
func (c ManageLinkOperation) ToPtr() *ManageLinkOperation {
	return &c
}

// OrderItemType - Order item type - purchase or rental
type OrderItemType string

const (
	// OrderItemTypePurchase - Purchase OrderItem.
	OrderItemTypePurchase OrderItemType = "Purchase"
	// OrderItemTypeRental - Rental OrderItem.
	OrderItemTypeRental OrderItemType = "Rental"
)

// PossibleOrderItemTypeValues returns the possible values for the OrderItemType const type.
func PossibleOrderItemTypeValues() []OrderItemType {
	return []OrderItemType{
		OrderItemTypePurchase,
		OrderItemTypeRental,
	}
}

// ToPtr returns a *OrderItemType pointing to the current value.
func (c OrderItemType) ToPtr() *OrderItemType {
	return &c
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	OriginSystem     Origin = "system"
	OriginUser       Origin = "user"
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// ToPtr returns a *Origin pointing to the current value.
func (c Origin) ToPtr() *Origin {
	return &c
}

// StageName - Stage name
type StageName string

const (
	// StageNameDeviceOrdered - An order has been created.
	StageNameDeviceOrdered StageName = "DeviceOrdered"
	// StageNameDevicePrepared - A device has been prepared for the order.
	StageNameDevicePrepared StageName = "DevicePrepared"
	// StageNamePickedUp - Device has been picked up from user and in transit to Azure datacenter.
	StageNamePickedUp StageName = "PickedUp"
	// StageNameAtAzureDC - Device has been received at Azure datacenter from the user.
	StageNameAtAzureDC StageName = "AtAzureDC"
	// StageNameDataCopy - Data copy from the device at Azure datacenter.
	StageNameDataCopy StageName = "DataCopy"
	// StageNameCompleted - Order has completed.
	StageNameCompleted StageName = "Completed"
	// StageNameCompletedWithErrors - Order has completed with errors.
	StageNameCompletedWithErrors StageName = "CompletedWithErrors"
	// StageNameCancelled - Order has been cancelled.
	StageNameCancelled StageName = "Cancelled"
	// StageNameAborted - Order has been aborted.
	StageNameAborted StageName = "Aborted"
	StageNameCurrent StageName = "Current"
	// StageNameCompletedWithWarnings - Order has completed with warnings.
	StageNameCompletedWithWarnings StageName = "CompletedWithWarnings"
	// StageNameReadyToDispatchFromAzureDC - Device is ready to be handed to customer from Azure DC.
	StageNameReadyToDispatchFromAzureDC StageName = "ReadyToDispatchFromAzureDC"
	// StageNameReadyToReceiveAtAzureDC - Device can be dropped off at Azure DC.
	StageNameReadyToReceiveAtAzureDC StageName = "ReadyToReceiveAtAzureDC"
	// StageNamePlaced - Currently in draft mode and can still be cancelled
	StageNamePlaced StageName = "Placed"
	// StageNameInReview - Order is currently in draft mode and can still be cancelled
	StageNameInReview StageName = "InReview"
	// StageNameConfirmed - Order is confirmed
	StageNameConfirmed StageName = "Confirmed"
	// StageNameReadyForDispatch - Order is ready for dispatch
	StageNameReadyForDispatch StageName = "ReadyForDispatch"
	// StageNameShipped - Order is in transit to customer
	StageNameShipped StageName = "Shipped"
	// StageNameDelivered - Order is delivered to customer
	StageNameDelivered StageName = "Delivered"
	// StageNameInUse - Order is in use at customer site
	StageNameInUse StageName = "InUse"
)

// PossibleStageNameValues returns the possible values for the StageName const type.
func PossibleStageNameValues() []StageName {
	return []StageName{
		StageNameDeviceOrdered,
		StageNameDevicePrepared,
		StageNamePickedUp,
		StageNameAtAzureDC,
		StageNameDataCopy,
		StageNameCompleted,
		StageNameCompletedWithErrors,
		StageNameCancelled,
		StageNameAborted,
		StageNameCurrent,
		StageNameCompletedWithWarnings,
		StageNameReadyToDispatchFromAzureDC,
		StageNameReadyToReceiveAtAzureDC,
		StageNamePlaced,
		StageNameInReview,
		StageNameConfirmed,
		StageNameReadyForDispatch,
		StageNameShipped,
		StageNameDelivered,
		StageNameInUse,
	}
}

// ToPtr returns a *StageName pointing to the current value.
func (c StageName) ToPtr() *StageName {
	return &c
}

// StageStatus - Stage status.
type StageStatus string

const (
	// StageStatusNone - No status available yet.
	StageStatusNone StageStatus = "None"
	// StageStatusInProgress - Stage is in progress.
	StageStatusInProgress StageStatus = "InProgress"
	// StageStatusSucceeded - Stage has succeeded.
	StageStatusSucceeded StageStatus = "Succeeded"
	// StageStatusFailed - Stage has failed.
	StageStatusFailed StageStatus = "Failed"
	// StageStatusCancelled - Stage has been cancelled.
	StageStatusCancelled StageStatus = "Cancelled"
	// StageStatusCancelling - Stage is cancelling.
	StageStatusCancelling StageStatus = "Cancelling"
)

// PossibleStageStatusValues returns the possible values for the StageStatus const type.
func PossibleStageStatusValues() []StageStatus {
	return []StageStatus{
		StageStatusNone,
		StageStatusInProgress,
		StageStatusSucceeded,
		StageStatusFailed,
		StageStatusCancelled,
		StageStatusCancelling,
	}
}

// ToPtr returns a *StageStatus pointing to the current value.
func (c StageStatus) ToPtr() *StageStatus {
	return &c
}
