package policy

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
// Code generated by Microsoft (R) AutoRest Code Generator 2.2.21.0
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// Mode enumerates the values for mode.
type Mode string

const (
	// All specifies the all state for mode.
	All Mode = "All"
	// Indexed specifies the indexed state for mode.
	Indexed Mode = "Indexed"
	// NotSpecified specifies the not specified state for mode.
	NotSpecified Mode = "NotSpecified"
)

// Type enumerates the values for type.
type Type string

const (
	// TypeBuiltIn specifies the type built in state for type.
	TypeBuiltIn Type = "BuiltIn"
	// TypeCustom specifies the type custom state for type.
	TypeCustom Type = "Custom"
	// TypeNotSpecified specifies the type not specified state for type.
	TypeNotSpecified Type = "NotSpecified"
)

// Assignment is the policy assignment.
type Assignment struct {
	autorest.Response     `json:"-"`
	*AssignmentProperties `json:"properties,omitempty"`
	ID                    *string `json:"id,omitempty"`
	Type                  *string `json:"type,omitempty"`
	Name                  *string `json:"name,omitempty"`
}

// AssignmentListResult is list of policy assignments.
type AssignmentListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Assignment `json:"value,omitempty"`
	NextLink          *string       `json:"nextLink,omitempty"`
}

// AssignmentListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client AssignmentListResult) AssignmentListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// AssignmentProperties is the policy assignment properties.
type AssignmentProperties struct {
	DisplayName        *string                 `json:"displayName,omitempty"`
	PolicyDefinitionID *string                 `json:"policyDefinitionId,omitempty"`
	Scope              *string                 `json:"scope,omitempty"`
	Parameters         *map[string]interface{} `json:"parameters,omitempty"`
	Description        *string                 `json:"description,omitempty"`
}

// Definition is the policy definition.
type Definition struct {
	autorest.Response     `json:"-"`
	*DefinitionProperties `json:"properties,omitempty"`
	ID                    *string `json:"id,omitempty"`
	Name                  *string `json:"name,omitempty"`
}

// DefinitionListResult is list of policy definitions.
type DefinitionListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Definition `json:"value,omitempty"`
	NextLink          *string       `json:"nextLink,omitempty"`
}

// DefinitionListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client DefinitionListResult) DefinitionListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// DefinitionProperties is the policy definition properties.
type DefinitionProperties struct {
	PolicyType  Type                    `json:"policyType,omitempty"`
	Mode        Mode                    `json:"mode,omitempty"`
	DisplayName *string                 `json:"displayName,omitempty"`
	Description *string                 `json:"description,omitempty"`
	PolicyRule  *map[string]interface{} `json:"policyRule,omitempty"`
	Metadata    *map[string]interface{} `json:"metadata,omitempty"`
	Parameters  *map[string]interface{} `json:"parameters,omitempty"`
}
