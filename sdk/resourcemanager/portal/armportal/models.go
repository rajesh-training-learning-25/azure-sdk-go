//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armportal

// Configuration - Tenant configuration.
type Configuration struct {
	// Tenant configuration properties.
	Properties *ConfigurationProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ConfigurationList - List of tenant configurations.
type ConfigurationList struct {
	// The URL to use for getting the next set of results.
	NextLink *string

	// The array of tenant configurations.
	Value []*Configuration
}

// ConfigurationProperties - Tenant configuration properties.
type ConfigurationProperties struct {
	// When flag is set to true Markdown tile will require external storage configuration (URI). The inline content configuration
	// will be prohibited.
	EnforcePrivateMarkdownStorage *bool
}

// Dashboard - The shared dashboard resource definition.
type Dashboard struct {
	// REQUIRED; Resource location
	Location *string

	// The shared dashboard properties.
	Properties *DashboardProperties

	// Resource tags
	Tags map[string]*string

	// READ-ONLY; Resource Id
	ID *string

	// READ-ONLY; Resource name
	Name *string

	// READ-ONLY; Resource type
	Type *string
}

// DashboardLens - A dashboard lens.
type DashboardLens struct {
	// REQUIRED; The lens order.
	Order *int32

	// REQUIRED; The dashboard parts.
	Parts []*DashboardParts

	// The dashboard len's metadata.
	Metadata map[string]any
}

// DashboardListResult - List of dashboards.
type DashboardListResult struct {
	// The URL to use for getting the next set of results.
	NextLink *string

	// The array of custom resource provider manifests.
	Value []*Dashboard
}

// DashboardPartMetadata - A dashboard part metadata.
type DashboardPartMetadata struct {
	// REQUIRED; The type of dashboard part.
	Type *string

	// OPTIONAL; Contains additional key/value pairs not defined in the schema.
	AdditionalProperties map[string]any
}

// GetDashboardPartMetadata implements the DashboardPartMetadataClassification interface for type DashboardPartMetadata.
func (d *DashboardPartMetadata) GetDashboardPartMetadata() *DashboardPartMetadata { return d }

// DashboardParts - A dashboard part.
type DashboardParts struct {
	// REQUIRED; The dashboard's part position.
	Position *DashboardPartsPosition

	// The dashboard part's metadata.
	Metadata DashboardPartMetadataClassification
}

// DashboardPartsPosition - The dashboard's part position.
type DashboardPartsPosition struct {
	// REQUIRED; The dashboard's part column span.
	ColSpan *int32

	// REQUIRED; The dashboard's part row span.
	RowSpan *int32

	// REQUIRED; The dashboard's part x coordinate.
	X *int32

	// REQUIRED; The dashboard's part y coordinate.
	Y *int32

	// The dashboard part's metadata.
	Metadata map[string]any
}

// DashboardProperties - The shared dashboard properties.
type DashboardProperties struct {
	// The dashboard lenses.
	Lenses []*DashboardLens

	// The dashboard metadata.
	Metadata map[string]any
}

// ErrorDefinition - Error definition.
type ErrorDefinition struct {
	// READ-ONLY; Service specific error code which serves as the substatus for the HTTP error code.
	Code *int32

	// READ-ONLY; Internal error details.
	Details []*ErrorDefinition

	// READ-ONLY; Description of the error.
	Message *string
}

// ErrorResponse - Error response.
type ErrorResponse struct {
	// The error details.
	Error *ErrorDefinition
}

// MarkdownPartMetadata - Markdown part metadata.
type MarkdownPartMetadata struct {
	// REQUIRED; The type of dashboard part.
	Type *string

	// OPTIONAL; Contains additional key/value pairs not defined in the schema.
	AdditionalProperties map[string]any

	// Input to dashboard part.
	Inputs []any

	// Markdown part settings.
	Settings *MarkdownPartMetadataSettings
}

// GetDashboardPartMetadata implements the DashboardPartMetadataClassification interface for type MarkdownPartMetadata.
func (m *MarkdownPartMetadata) GetDashboardPartMetadata() *DashboardPartMetadata {
	return &DashboardPartMetadata{
		AdditionalProperties: m.AdditionalProperties,
		Type:                 m.Type,
	}
}

// MarkdownPartMetadataSettings - Markdown part settings.
type MarkdownPartMetadataSettings struct {
	// The content of markdown part.
	Content *MarkdownPartMetadataSettingsContent
}

// MarkdownPartMetadataSettingsContent - The content of markdown part.
type MarkdownPartMetadataSettingsContent struct {
	// The setting of the content of markdown part.
	Settings *MarkdownPartMetadataSettingsContentSettings
}

// MarkdownPartMetadataSettingsContentSettings - The setting of the content of markdown part.
type MarkdownPartMetadataSettingsContentSettings struct {
	// The content of the markdown part.
	Content *string

	// The source of the content of the markdown part.
	MarkdownSource *int32

	// The uri of markdown content.
	MarkdownURI *string

	// The subtitle of the markdown part.
	Subtitle *string

	// The title of the markdown part.
	Title *string
}

// PatchableDashboard - The shared dashboard resource definition.
type PatchableDashboard struct {
	// The shared dashboard properties.
	Properties *DashboardProperties

	// Resource tags
	Tags map[string]*string
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

// ResourceProviderOperation - Supported operations of this resource provider.
type ResourceProviderOperation struct {
	// Display metadata associated with the operation.
	Display *ResourceProviderOperationDisplay

	// Indicates whether the operation applies to data-plane.
	IsDataAction *string

	// Operation name, in format of {provider}/{resource}/{operation}
	Name *string
}

// ResourceProviderOperationDisplay - Display metadata associated with the operation.
type ResourceProviderOperationDisplay struct {
	// Description of this operation.
	Description *string

	// Type of operation: get, read, delete, etc.
	Operation *string

	// Resource provider: Microsoft Custom Providers.
	Provider *string

	// Resource on which the operation is performed.
	Resource *string
}

// ResourceProviderOperationList - Results of the request to list operations.
type ResourceProviderOperationList struct {
	// The URL to use for getting the next set of results.
	NextLink *string

	// List of operations supported by this resource provider.
	Value []*ResourceProviderOperation
}

// Violation information.
type Violation struct {
	// READ-ONLY; Error message.
	ErrorMessage *string

	// READ-ONLY; Id of the item that violates tenant configuration.
	ID *string

	// READ-ONLY; Id of the user who owns violated item.
	UserID *string
}

// ViolationsList - List of list of items that violate tenant's configuration.
type ViolationsList struct {
	// The URL to use for getting the next set of results.
	NextLink *string

	// The array of violations.
	Value []*Violation
}
