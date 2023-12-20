//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armgraphservices

import "time"

// AccountPatchResource - Account patch properties
type AccountPatchResource struct {
	// List of key value pairs that describe the resource. This will overwrite the existing tags.
	Tags map[string]*string
}

// AccountResource - Account details
type AccountResource struct {
	// REQUIRED; Property bag from billing account
	Properties *AccountResourceProperties

	// Location of the resource.
	Location *string

	// resource tags.
	Tags map[string]*string

	// READ-ONLY; Azure resource ID.
	ID *string

	// READ-ONLY; Azure resource name.
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *AccountResourceSystemData

	// READ-ONLY; Azure resource type.
	Type *string
}

// AccountResourceList - The list of accounts.
type AccountResourceList struct {
	// The link to the next page of items
	NextLink *string

	// The list of recommendations.
	Value []*AccountResource
}

// AccountResourceProperties - Property bag from billing account
type AccountResourceProperties struct {
	// REQUIRED; Customer owned application ID
	AppID *string

	// READ-ONLY; Billing Plan Id
	BillingPlanID *string

	// READ-ONLY; Provisioning state.
	ProvisioningState *ProvisioningState
}

// AccountResourceSystemData - Metadata pertaining to creation and last modification of the resource.
type AccountResourceSystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation
}
