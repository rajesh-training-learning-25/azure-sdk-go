//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armchangeanalysis

import "time"

// Change - The detected change.
type Change struct {
	// The properties of a change.
	Properties *ChangeProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ChangeList - The list of detected changes.
type ChangeList struct {
	// The URI that can be used to request the next page of changes.
	NextLink *string

	// The list of changes.
	Value []*Change
}

// ChangeProperties - The properties of a change.
type ChangeProperties struct {
	// The type of the change.
	ChangeType *ChangeType

	// The list of identities who might initiated the change. The identity could be user name (email address) or the object ID
	// of the Service Principal.
	InitiatedByList []*string

	// The list of detailed changes at json property level.
	PropertyChanges []*PropertyChange

	// The resource id that the change is attached to.
	ResourceID *string

	// The time when the change is detected.
	TimeStamp *time.Time
}

// ChangesClientListChangesByResourceGroupOptions contains the optional parameters for the ChangesClient.NewListChangesByResourceGroupPager
// method.
type ChangesClientListChangesByResourceGroupOptions struct {
	// A skip token is used to continue retrieving items after an operation returns a partial result. If a previous response contains
	// a nextLink element, the value of the nextLink element will include a
	// skipToken parameter that specifies a starting point to use for subsequent calls.
	SkipToken *string
}

// ChangesClientListChangesBySubscriptionOptions contains the optional parameters for the ChangesClient.NewListChangesBySubscriptionPager
// method.
type ChangesClientListChangesBySubscriptionOptions struct {
	// A skip token is used to continue retrieving items after an operation returns a partial result. If a previous response contains
	// a nextLink element, the value of the nextLink element will include a
	// skipToken parameter that specifies a starting point to use for subsequent calls.
	SkipToken *string
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info any

	// READ-ONLY; The additional info type.
	Type *string
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo

	// READ-ONLY; The error code.
	Code *string

	// READ-ONLY; The error details.
	Details []*ErrorDetail

	// READ-ONLY; The error message.
	Message *string

	// READ-ONLY; The error target.
	Target *string
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.).
type ErrorResponse struct {
	// The error object.
	Error *ErrorDetail
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// A skip token is used to continue retrieving items after an operation returns a partial result. If a previous response contains
	// a nextLink element, the value of the nextLink element will include a
	// skipToken parameter that specifies a starting point to use for subsequent calls.
	SkipToken *string
}

// PropertyChange - Data of a property change.
type PropertyChange struct {
	// The change category.
	ChangeCategory *ChangeCategory

	// The type of the change.
	ChangeType *ChangeType

	// The description of the changed property.
	Description *string

	// The enhanced display name of the json path. E.g., the json path value[0].properties will be translated to something meaningful
	// like slots["Staging"].properties.
	DisplayName *string

	// The boolean indicating whether the oldValue and newValue are masked. The values are masked if it contains sensitive information
	// that the user doesn't have access to.
	IsDataMasked *bool

	// The json path of the changed property.
	JSONPath *string
	Level    *Level

	// The value of the property after the change.
	NewValue *string

	// The value of the property before the change.
	OldValue *string
}

// ProxyResource - The resource model definition for a Azure Resource Manager proxy resource. It will not have tags and a
// location
type ProxyResource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ResourceChangesClientListOptions contains the optional parameters for the ResourceChangesClient.NewListPager method.
type ResourceChangesClientListOptions struct {
	// A skip token is used to continue retrieving items after an operation returns a partial result. If a previous response contains
	// a nextLink element, the value of the nextLink element will include a
	// skipToken parameter that specifies a starting point to use for subsequent calls.
	SkipToken *string
}

// ResourceProviderOperationDefinition - The resource provider operation definition.
type ResourceProviderOperationDefinition struct {
	// The resource provider operation details.
	Display *ResourceProviderOperationDisplay

	// The resource provider operation name.
	Name *string
}

// ResourceProviderOperationDisplay - The resource provider operation details.
type ResourceProviderOperationDisplay struct {
	// Description of the resource provider operation.
	Description *string

	// Name of the resource provider operation.
	Operation *string

	// Name of the resource provider.
	Provider *string

	// Name of the resource type.
	Resource *string
}

// ResourceProviderOperationList - The resource provider operation list.
type ResourceProviderOperationList struct {
	// The URI that can be used to request the next page for list of Azure operations.
	NextLink *string

	// Resource provider operations list.
	Value []*ResourceProviderOperationDefinition
}
