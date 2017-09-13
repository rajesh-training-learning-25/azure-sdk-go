package powerbiembedded

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
)

// AccessKeyName enumerates the values for access key name.
type AccessKeyName string

const (
	// Key1 specifies the key 1 state for access key name.
	Key1 AccessKeyName = "key1"
	// Key2 specifies the key 2 state for access key name.
	Key2 AccessKeyName = "key2"
)

// CheckNameReason enumerates the values for check name reason.
type CheckNameReason string

const (
	// Invalid specifies the invalid state for check name reason.
	Invalid CheckNameReason = "Invalid"
	// Unavailable specifies the unavailable state for check name reason.
	Unavailable CheckNameReason = "Unavailable"
)

// AzureSku is
type AzureSku struct {
	Name *string `json:"name,omitempty"`
	Tier *string `json:"tier,omitempty"`
}

// CheckNameRequest is
type CheckNameRequest struct {
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// CheckNameResponse is
type CheckNameResponse struct {
	autorest.Response `json:"-"`
	NameAvailable     *bool           `json:"nameAvailable,omitempty"`
	Reason            CheckNameReason `json:"reason,omitempty"`
	Message           *string         `json:"message,omitempty"`
}

// CreateWorkspaceCollectionRequest is
type CreateWorkspaceCollectionRequest struct {
	Location *string             `json:"location,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
	Sku      *AzureSku           `json:"sku,omitempty"`
}

// Display is
type Display struct {
	Provider    *string `json:"provider,omitempty"`
	Resource    *string `json:"resource,omitempty"`
	Operation   *string `json:"operation,omitempty"`
	Description *string `json:"description,omitempty"`
	Origin      *string `json:"origin,omitempty"`
}

// Error is
type Error struct {
	Code    *string        `json:"code,omitempty"`
	Message *string        `json:"message,omitempty"`
	Target  *string        `json:"target,omitempty"`
	Details *[]ErrorDetail `json:"details,omitempty"`
}

// ErrorDetail is
type ErrorDetail struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Target  *string `json:"target,omitempty"`
}

// MigrateWorkspaceCollectionRequest is
type MigrateWorkspaceCollectionRequest struct {
	TargetResourceGroup *string   `json:"targetResourceGroup,omitempty"`
	Resources           *[]string `json:"resources,omitempty"`
}

// Operation is
type Operation struct {
	Name    *string  `json:"name,omitempty"`
	Display *Display `json:"display,omitempty"`
}

// OperationList is
type OperationList struct {
	autorest.Response `json:"-"`
	Value             *[]Operation `json:"value,omitempty"`
}

// UpdateWorkspaceCollectionRequest is
type UpdateWorkspaceCollectionRequest struct {
	Tags *map[string]*string `json:"tags,omitempty"`
	Sku  *AzureSku           `json:"sku,omitempty"`
}

// Workspace is
type Workspace struct {
	ID         *string                 `json:"id,omitempty"`
	Name       *string                 `json:"name,omitempty"`
	Type       *string                 `json:"type,omitempty"`
	Properties *map[string]interface{} `json:"properties,omitempty"`
}

// WorkspaceCollection is
type WorkspaceCollection struct {
	autorest.Response `json:"-"`
	ID                *string                 `json:"id,omitempty"`
	Name              *string                 `json:"name,omitempty"`
	Type              *string                 `json:"type,omitempty"`
	Location          *string                 `json:"location,omitempty"`
	Tags              *map[string]*string     `json:"tags,omitempty"`
	Sku               *AzureSku               `json:"sku,omitempty"`
	Properties        *map[string]interface{} `json:"properties,omitempty"`
}

// WorkspaceCollectionAccessKey is
type WorkspaceCollectionAccessKey struct {
	KeyName AccessKeyName `json:"keyName,omitempty"`
}

// WorkspaceCollectionAccessKeys is
type WorkspaceCollectionAccessKeys struct {
	autorest.Response `json:"-"`
	Key1              *string `json:"key1,omitempty"`
	Key2              *string `json:"key2,omitempty"`
}

// WorkspaceCollectionList is
type WorkspaceCollectionList struct {
	autorest.Response `json:"-"`
	Value             *[]WorkspaceCollection `json:"value,omitempty"`
}

// WorkspaceList is
type WorkspaceList struct {
	autorest.Response `json:"-"`
	Value             *[]Workspace `json:"value,omitempty"`
}
