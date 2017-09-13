package resources

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 2.2.18.0
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// DeploymentMode enumerates the values for deployment mode.
type DeploymentMode string

const (
	// Complete specifies the complete state for deployment mode.
	Complete DeploymentMode = "Complete"
	// Incremental specifies the incremental state for deployment mode.
	Incremental DeploymentMode = "Incremental"
)

// ResourceIdentityType enumerates the values for resource identity type.
type ResourceIdentityType string

const (
	// SystemAssigned specifies the system assigned state for resource identity type.
	SystemAssigned ResourceIdentityType = "SystemAssigned"
)

// AliasPathType is the type of the paths for alias.
type AliasPathType struct {
	Path        *string   `json:"path,omitempty"`
	APIVersions *[]string `json:"apiVersions,omitempty"`
}

// AliasType is the alias type.
type AliasType struct {
	Name  *string          `json:"name,omitempty"`
	Paths *[]AliasPathType `json:"paths,omitempty"`
}

// BasicDependency is deployment dependency information.
type BasicDependency struct {
	ID           *string `json:"id,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	ResourceName *string `json:"resourceName,omitempty"`
}

// DebugSetting is
type DebugSetting struct {
	DetailLevel *string `json:"detailLevel,omitempty"`
}

// Dependency is deployment dependency information.
type Dependency struct {
	DependsOn    *[]BasicDependency `json:"dependsOn,omitempty"`
	ID           *string            `json:"id,omitempty"`
	ResourceType *string            `json:"resourceType,omitempty"`
	ResourceName *string            `json:"resourceName,omitempty"`
}

// Deployment is deployment operation parameters.
type Deployment struct {
	Properties *DeploymentProperties `json:"properties,omitempty"`
}

// DeploymentExportResult is the deployment export result.
type DeploymentExportResult struct {
	autorest.Response `json:"-"`
	Template          *map[string]interface{} `json:"template,omitempty"`
}

// DeploymentExtended is deployment information.
type DeploymentExtended struct {
	autorest.Response `json:"-"`
	ID                *string                       `json:"id,omitempty"`
	Name              *string                       `json:"name,omitempty"`
	Properties        *DeploymentPropertiesExtended `json:"properties,omitempty"`
}

// DeploymentExtendedFilter is deployment filter.
type DeploymentExtendedFilter struct {
	ProvisioningState *string `json:"provisioningState,omitempty"`
}

// DeploymentListResult is list of deployments.
type DeploymentListResult struct {
	autorest.Response `json:"-"`
	Value             *[]DeploymentExtended `json:"value,omitempty"`
	NextLink          *string               `json:"nextLink,omitempty"`
}

// DeploymentListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client DeploymentListResult) DeploymentListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// DeploymentOperation is deployment operation information.
type DeploymentOperation struct {
	autorest.Response `json:"-"`
	ID                *string                        `json:"id,omitempty"`
	OperationID       *string                        `json:"operationId,omitempty"`
	Properties        *DeploymentOperationProperties `json:"properties,omitempty"`
}

// DeploymentOperationProperties is deployment operation properties.
type DeploymentOperationProperties struct {
	ProvisioningState *string                 `json:"provisioningState,omitempty"`
	Timestamp         *date.Time              `json:"timestamp,omitempty"`
	ServiceRequestID  *string                 `json:"serviceRequestId,omitempty"`
	StatusCode        *string                 `json:"statusCode,omitempty"`
	StatusMessage     *map[string]interface{} `json:"statusMessage,omitempty"`
	TargetResource    *TargetResource         `json:"targetResource,omitempty"`
	Request           *HTTPMessage            `json:"request,omitempty"`
	Response          *HTTPMessage            `json:"response,omitempty"`
}

// DeploymentOperationsListResult is list of deployment operations.
type DeploymentOperationsListResult struct {
	autorest.Response `json:"-"`
	Value             *[]DeploymentOperation `json:"value,omitempty"`
	NextLink          *string                `json:"nextLink,omitempty"`
}

// DeploymentOperationsListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client DeploymentOperationsListResult) DeploymentOperationsListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// DeploymentProperties is deployment properties.
type DeploymentProperties struct {
	Template       *map[string]interface{} `json:"template,omitempty"`
	TemplateLink   *TemplateLink           `json:"templateLink,omitempty"`
	Parameters     *map[string]interface{} `json:"parameters,omitempty"`
	ParametersLink *ParametersLink         `json:"parametersLink,omitempty"`
	Mode           DeploymentMode          `json:"mode,omitempty"`
	DebugSetting   *DebugSetting           `json:"debugSetting,omitempty"`
}

// DeploymentPropertiesExtended is deployment properties with additional details.
type DeploymentPropertiesExtended struct {
	ProvisioningState *string                 `json:"provisioningState,omitempty"`
	CorrelationID     *string                 `json:"correlationId,omitempty"`
	Timestamp         *date.Time              `json:"timestamp,omitempty"`
	Outputs           *map[string]interface{} `json:"outputs,omitempty"`
	Providers         *[]Provider             `json:"providers,omitempty"`
	Dependencies      *[]Dependency           `json:"dependencies,omitempty"`
	Template          *map[string]interface{} `json:"template,omitempty"`
	TemplateLink      *TemplateLink           `json:"templateLink,omitempty"`
	Parameters        *map[string]interface{} `json:"parameters,omitempty"`
	ParametersLink    *ParametersLink         `json:"parametersLink,omitempty"`
	Mode              DeploymentMode          `json:"mode,omitempty"`
	DebugSetting      *DebugSetting           `json:"debugSetting,omitempty"`
}

// DeploymentValidateResult is information from validate template deployment response.
type DeploymentValidateResult struct {
	autorest.Response `json:"-"`
	Error             *ManagementErrorWithDetails   `json:"error,omitempty"`
	Properties        *DeploymentPropertiesExtended `json:"properties,omitempty"`
}

// ExportTemplateRequest is export resource group template request parameters.
type ExportTemplateRequest struct {
	ResourcesProperty *[]string `json:"resources,omitempty"`
	Options           *string   `json:"options,omitempty"`
}

// GenericResource is resource information.
type GenericResource struct {
	autorest.Response `json:"-"`
	ID                *string                 `json:"id,omitempty"`
	Name              *string                 `json:"name,omitempty"`
	Type              *string                 `json:"type,omitempty"`
	Location          *string                 `json:"location,omitempty"`
	Tags              *map[string]*string     `json:"tags,omitempty"`
	Plan              *Plan                   `json:"plan,omitempty"`
	Properties        *map[string]interface{} `json:"properties,omitempty"`
	Kind              *string                 `json:"kind,omitempty"`
	ManagedBy         *string                 `json:"managedBy,omitempty"`
	Sku               *Sku                    `json:"sku,omitempty"`
	Identity          *Identity               `json:"identity,omitempty"`
}

// GenericResourceFilter is resource filter.
type GenericResourceFilter struct {
	ResourceType *string `json:"resourceType,omitempty"`
	Tagname      *string `json:"tagname,omitempty"`
	Tagvalue     *string `json:"tagvalue,omitempty"`
}

// Group is resource group information.
type Group struct {
	autorest.Response `json:"-"`
	ID                *string             `json:"id,omitempty"`
	Name              *string             `json:"name,omitempty"`
	Properties        *GroupProperties    `json:"properties,omitempty"`
	Location          *string             `json:"location,omitempty"`
	ManagedBy         *string             `json:"managedBy,omitempty"`
	Tags              *map[string]*string `json:"tags,omitempty"`
}

// GroupExportResult is resource group export result.
type GroupExportResult struct {
	autorest.Response `json:"-"`
	Template          *map[string]interface{}     `json:"template,omitempty"`
	Error             *ManagementErrorWithDetails `json:"error,omitempty"`
}

// GroupFilter is resource group filter.
type GroupFilter struct {
	TagName  *string `json:"tagName,omitempty"`
	TagValue *string `json:"tagValue,omitempty"`
}

// GroupListResult is list of resource groups.
type GroupListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Group `json:"value,omitempty"`
	NextLink          *string  `json:"nextLink,omitempty"`
}

// GroupListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client GroupListResult) GroupListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// GroupPatchable is resource group information.
type GroupPatchable struct {
	Name       *string             `json:"name,omitempty"`
	Properties *GroupProperties    `json:"properties,omitempty"`
	ManagedBy  *string             `json:"managedBy,omitempty"`
	Tags       *map[string]*string `json:"tags,omitempty"`
}

// GroupProperties is the resource group properties.
type GroupProperties struct {
	ProvisioningState *string `json:"provisioningState,omitempty"`
}

// HTTPMessage is HTTP message.
type HTTPMessage struct {
	Content *map[string]interface{} `json:"content,omitempty"`
}

// Identity is identity for the resource.
type Identity struct {
	PrincipalID *string              `json:"principalId,omitempty"`
	TenantID    *string              `json:"tenantId,omitempty"`
	Type        ResourceIdentityType `json:"type,omitempty"`
}

// ListResult is list of resource groups.
type ListResult struct {
	autorest.Response `json:"-"`
	Value             *[]GenericResource `json:"value,omitempty"`
	NextLink          *string            `json:"nextLink,omitempty"`
}

// ListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ListResult) ListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// ManagementErrorWithDetails is the detailed error message of resource management.
type ManagementErrorWithDetails struct {
	Code    *string                       `json:"code,omitempty"`
	Message *string                       `json:"message,omitempty"`
	Target  *string                       `json:"target,omitempty"`
	Details *[]ManagementErrorWithDetails `json:"details,omitempty"`
}

// MoveInfo is parameters of move resources.
type MoveInfo struct {
	ResourcesProperty   *[]string `json:"resources,omitempty"`
	TargetResourceGroup *string   `json:"targetResourceGroup,omitempty"`
}

// ParametersLink is entity representing the reference to the deployment paramaters.
type ParametersLink struct {
	URI            *string `json:"uri,omitempty"`
	ContentVersion *string `json:"contentVersion,omitempty"`
}

// Plan is plan for the resource.
type Plan struct {
	Name          *string `json:"name,omitempty"`
	Publisher     *string `json:"publisher,omitempty"`
	Product       *string `json:"product,omitempty"`
	PromotionCode *string `json:"promotionCode,omitempty"`
}

// Provider is resource provider information.
type Provider struct {
	autorest.Response `json:"-"`
	ID                *string                 `json:"id,omitempty"`
	Namespace         *string                 `json:"namespace,omitempty"`
	RegistrationState *string                 `json:"registrationState,omitempty"`
	ResourceTypes     *[]ProviderResourceType `json:"resourceTypes,omitempty"`
}

// ProviderListResult is list of resource providers.
type ProviderListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Provider `json:"value,omitempty"`
	NextLink          *string     `json:"nextLink,omitempty"`
}

// ProviderListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ProviderListResult) ProviderListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// ProviderOperationDisplayProperties is resource provider operation's display properties.
type ProviderOperationDisplayProperties struct {
	Publisher   *string `json:"publisher,omitempty"`
	Provider    *string `json:"provider,omitempty"`
	Resource    *string `json:"resource,omitempty"`
	Operation   *string `json:"operation,omitempty"`
	Description *string `json:"description,omitempty"`
}

// ProviderResourceType is resource type managed by the resource provider.
type ProviderResourceType struct {
	ResourceType *string             `json:"resourceType,omitempty"`
	Locations    *[]string           `json:"locations,omitempty"`
	Aliases      *[]AliasType        `json:"aliases,omitempty"`
	APIVersions  *[]string           `json:"apiVersions,omitempty"`
	Properties   *map[string]*string `json:"properties,omitempty"`
}

// Resource is resource.
type Resource struct {
	ID       *string             `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Location *string             `json:"location,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
}

// Sku is SKU for the resource.
type Sku struct {
	Name     *string `json:"name,omitempty"`
	Tier     *string `json:"tier,omitempty"`
	Size     *string `json:"size,omitempty"`
	Family   *string `json:"family,omitempty"`
	Model    *string `json:"model,omitempty"`
	Capacity *int32  `json:"capacity,omitempty"`
}

// SubResource is sub-resource.
type SubResource struct {
	ID *string `json:"id,omitempty"`
}

// TagCount is tag count.
type TagCount struct {
	Type  *string `json:"type,omitempty"`
	Value *int32  `json:"value,omitempty"`
}

// TagDetails is tag details.
type TagDetails struct {
	autorest.Response `json:"-"`
	ID                *string     `json:"id,omitempty"`
	TagName           *string     `json:"tagName,omitempty"`
	Count             *TagCount   `json:"count,omitempty"`
	Values            *[]TagValue `json:"values,omitempty"`
}

// TagsListResult is list of subscription tags.
type TagsListResult struct {
	autorest.Response `json:"-"`
	Value             *[]TagDetails `json:"value,omitempty"`
	NextLink          *string       `json:"nextLink,omitempty"`
}

// TagsListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client TagsListResult) TagsListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// TagValue is tag information.
type TagValue struct {
	autorest.Response `json:"-"`
	ID                *string   `json:"id,omitempty"`
	TagValue          *string   `json:"tagValue,omitempty"`
	Count             *TagCount `json:"count,omitempty"`
}

// TargetResource is target resource.
type TargetResource struct {
	ID           *string `json:"id,omitempty"`
	ResourceName *string `json:"resourceName,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
}

// TemplateLink is entity representing the reference to the template.
type TemplateLink struct {
	URI            *string `json:"uri,omitempty"`
	ContentVersion *string `json:"contentVersion,omitempty"`
}
