package authorization

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
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// ClassicAdministrator is classic Administrators
type ClassicAdministrator struct {
	ID         *string                         `json:"id,omitempty"`
	Name       *string                         `json:"name,omitempty"`
	Type       *string                         `json:"type,omitempty"`
	Properties *ClassicAdministratorProperties `json:"properties,omitempty"`
}

// ClassicAdministratorListResult is classicAdministrator list result information.
type ClassicAdministratorListResult struct {
	autorest.Response `json:"-"`
	Value             *[]ClassicAdministrator `json:"value,omitempty"`
	NextLink          *string                 `json:"nextLink,omitempty"`
}

// ClassicAdministratorListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ClassicAdministratorListResult) ClassicAdministratorListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// ClassicAdministratorProperties is classic Administrator properties.
type ClassicAdministratorProperties struct {
	EmailAddress *string `json:"emailAddress,omitempty"`
	Role         *string `json:"role,omitempty"`
}

// Permission is role definition permissions.
type Permission struct {
	Actions    *[]string `json:"actions,omitempty"`
	NotActions *[]string `json:"notActions,omitempty"`
}

// PermissionGetResult is permissions information.
type PermissionGetResult struct {
	autorest.Response `json:"-"`
	Value             *[]Permission `json:"value,omitempty"`
	NextLink          *string       `json:"nextLink,omitempty"`
}

// PermissionGetResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client PermissionGetResult) PermissionGetResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// ProviderOperation is operation
type ProviderOperation struct {
	Name        *string                 `json:"name,omitempty"`
	DisplayName *string                 `json:"displayName,omitempty"`
	Description *string                 `json:"description,omitempty"`
	Origin      *string                 `json:"origin,omitempty"`
	Properties  *map[string]interface{} `json:"properties,omitempty"`
}

// ProviderOperationsMetadata is provider Operations metadata
type ProviderOperationsMetadata struct {
	autorest.Response `json:"-"`
	ID                *string              `json:"id,omitempty"`
	Name              *string              `json:"name,omitempty"`
	Type              *string              `json:"type,omitempty"`
	DisplayName       *string              `json:"displayName,omitempty"`
	ResourceTypes     *[]ResourceType      `json:"resourceTypes,omitempty"`
	Operations        *[]ProviderOperation `json:"operations,omitempty"`
}

// ProviderOperationsMetadataListResult is provider operations metadata list
type ProviderOperationsMetadataListResult struct {
	autorest.Response `json:"-"`
	Value             *[]ProviderOperationsMetadata `json:"value,omitempty"`
	NextLink          *string                       `json:"nextLink,omitempty"`
}

// ProviderOperationsMetadataListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ProviderOperationsMetadataListResult) ProviderOperationsMetadataListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// ResourceType is resource Type
type ResourceType struct {
	Name        *string              `json:"name,omitempty"`
	DisplayName *string              `json:"displayName,omitempty"`
	Operations  *[]ProviderOperation `json:"operations,omitempty"`
}

// RoleAssignment is role Assignments
type RoleAssignment struct {
	autorest.Response `json:"-"`
	ID                *string                            `json:"id,omitempty"`
	Name              *string                            `json:"name,omitempty"`
	Type              *string                            `json:"type,omitempty"`
	Properties        *RoleAssignmentPropertiesWithScope `json:"properties,omitempty"`
}

// RoleAssignmentCreateParameters is role assignment create parameters.
type RoleAssignmentCreateParameters struct {
	Properties *RoleAssignmentProperties `json:"properties,omitempty"`
}

// RoleAssignmentFilter is role Assignments filter
type RoleAssignmentFilter struct {
	PrincipalID *string `json:"principalId,omitempty"`
}

// RoleAssignmentListResult is role assignment list operation result.
type RoleAssignmentListResult struct {
	autorest.Response `json:"-"`
	Value             *[]RoleAssignment `json:"value,omitempty"`
	NextLink          *string           `json:"nextLink,omitempty"`
}

// RoleAssignmentListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client RoleAssignmentListResult) RoleAssignmentListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// RoleAssignmentProperties is role assignment properties.
type RoleAssignmentProperties struct {
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty"`
	PrincipalID      *string `json:"principalId,omitempty"`
}

// RoleAssignmentPropertiesWithScope is role assignment properties with scope.
type RoleAssignmentPropertiesWithScope struct {
	Scope            *string `json:"scope,omitempty"`
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty"`
	PrincipalID      *string `json:"principalId,omitempty"`
}

// RoleDefinition is role definition.
type RoleDefinition struct {
	autorest.Response `json:"-"`
	ID                *string                   `json:"id,omitempty"`
	Name              *string                   `json:"name,omitempty"`
	Type              *string                   `json:"type,omitempty"`
	Properties        *RoleDefinitionProperties `json:"properties,omitempty"`
}

// RoleDefinitionFilter is role Definitions filter
type RoleDefinitionFilter struct {
	RoleName *string `json:"roleName,omitempty"`
}

// RoleDefinitionListResult is role definition list operation result.
type RoleDefinitionListResult struct {
	autorest.Response `json:"-"`
	Value             *[]RoleDefinition `json:"value,omitempty"`
	NextLink          *string           `json:"nextLink,omitempty"`
}

// RoleDefinitionListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client RoleDefinitionListResult) RoleDefinitionListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// RoleDefinitionProperties is role definition properties.
type RoleDefinitionProperties struct {
	RoleName         *string       `json:"roleName,omitempty"`
	Description      *string       `json:"description,omitempty"`
	Type             *string       `json:"type,omitempty"`
	Permissions      *[]Permission `json:"permissions,omitempty"`
	AssignableScopes *[]string     `json:"assignableScopes,omitempty"`
}
