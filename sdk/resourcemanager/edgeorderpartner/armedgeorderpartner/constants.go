//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armedgeorderpartner

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeorderpartner/armedgeorderpartner"
	moduleVersion = "v0.7.0"
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

// ManageLinkOperation - Operation to be performed - Link, Unlink, Relink
type ManageLinkOperation string

const (
	// ManageLinkOperationLink - Link.
	ManageLinkOperationLink ManageLinkOperation = "Link"
	// ManageLinkOperationRelink - Relink.
	ManageLinkOperationRelink ManageLinkOperation = "Relink"
	// ManageLinkOperationUnlink - Unlink.
	ManageLinkOperationUnlink ManageLinkOperation = "Unlink"
)

// PossibleManageLinkOperationValues returns the possible values for the ManageLinkOperation const type.
func PossibleManageLinkOperationValues() []ManageLinkOperation {
	return []ManageLinkOperation{
		ManageLinkOperationLink,
		ManageLinkOperationRelink,
		ManageLinkOperationUnlink,
	}
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

// StageName - Stage name
type StageName string

const (
	// StageNameAborted - Order has been aborted.
	StageNameAborted StageName = "Aborted"
	// StageNameAtAzureDC - Device has been received at Azure datacenter from the user.
	StageNameAtAzureDC StageName = "AtAzureDC"
	// StageNameCancelled - Order has been cancelled.
	StageNameCancelled StageName = "Cancelled"
	// StageNameCompleted - Order has completed.
	StageNameCompleted StageName = "Completed"
	// StageNameCompletedWithErrors - Order has completed with errors.
	StageNameCompletedWithErrors StageName = "CompletedWithErrors"
	// StageNameCompletedWithWarnings - Order has completed with warnings.
	StageNameCompletedWithWarnings StageName = "CompletedWithWarnings"
	// StageNameConfirmed - Order is confirmed
	StageNameConfirmed StageName = "Confirmed"
	StageNameCurrent   StageName = "Current"
	// StageNameDataCopy - Data copy from the device at Azure datacenter.
	StageNameDataCopy StageName = "DataCopy"
	// StageNameDelivered - Order is delivered to customer
	StageNameDelivered StageName = "Delivered"
	// StageNameDeviceOrdered - An order has been created.
	StageNameDeviceOrdered StageName = "DeviceOrdered"
	// StageNameDevicePrepared - A device has been prepared for the order.
	StageNameDevicePrepared StageName = "DevicePrepared"
	// StageNameInReview - Order is currently in draft mode and can still be cancelled
	StageNameInReview StageName = "InReview"
	// StageNameInUse - Order is in use at customer site
	StageNameInUse StageName = "InUse"
	// StageNamePickedUp - Device has been picked up from user and in transit to Azure datacenter.
	StageNamePickedUp StageName = "PickedUp"
	// StageNamePlaced - Currently in draft mode and can still be cancelled
	StageNamePlaced StageName = "Placed"
	// StageNameReadyForDispatch - Order is ready for dispatch
	StageNameReadyForDispatch StageName = "ReadyForDispatch"
	// StageNameReadyToDispatchFromAzureDC - Device is ready to be handed to customer from Azure DC.
	StageNameReadyToDispatchFromAzureDC StageName = "ReadyToDispatchFromAzureDC"
	// StageNameReadyToReceiveAtAzureDC - Device can be dropped off at Azure DC.
	StageNameReadyToReceiveAtAzureDC StageName = "ReadyToReceiveAtAzureDC"
	// StageNameShipped - Order is in transit to customer
	StageNameShipped StageName = "Shipped"
)

// PossibleStageNameValues returns the possible values for the StageName const type.
func PossibleStageNameValues() []StageName {
	return []StageName{
		StageNameAborted,
		StageNameAtAzureDC,
		StageNameCancelled,
		StageNameCompleted,
		StageNameCompletedWithErrors,
		StageNameCompletedWithWarnings,
		StageNameConfirmed,
		StageNameCurrent,
		StageNameDataCopy,
		StageNameDelivered,
		StageNameDeviceOrdered,
		StageNameDevicePrepared,
		StageNameInReview,
		StageNameInUse,
		StageNamePickedUp,
		StageNamePlaced,
		StageNameReadyForDispatch,
		StageNameReadyToDispatchFromAzureDC,
		StageNameReadyToReceiveAtAzureDC,
		StageNameShipped,
	}
}

// StageStatus - Stage status.
type StageStatus string

const (
	// StageStatusCancelled - Stage has been cancelled.
	StageStatusCancelled StageStatus = "Cancelled"
	// StageStatusCancelling - Stage is cancelling.
	StageStatusCancelling StageStatus = "Cancelling"
	// StageStatusFailed - Stage has failed.
	StageStatusFailed StageStatus = "Failed"
	// StageStatusInProgress - Stage is in progress.
	StageStatusInProgress StageStatus = "InProgress"
	// StageStatusNone - No status available yet.
	StageStatusNone StageStatus = "None"
	// StageStatusSucceeded - Stage has succeeded.
	StageStatusSucceeded StageStatus = "Succeeded"
)

// PossibleStageStatusValues returns the possible values for the StageStatus const type.
func PossibleStageStatusValues() []StageStatus {
	return []StageStatus{
		StageStatusCancelled,
		StageStatusCancelling,
		StageStatusFailed,
		StageStatusInProgress,
		StageStatusNone,
		StageStatusSucceeded,
	}
}
