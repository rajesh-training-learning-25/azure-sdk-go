package search

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

// AdminKeyKind enumerates the values for admin key kind.
type AdminKeyKind string

const (
	// Primary specifies the primary state for admin key kind.
	Primary AdminKeyKind = "primary"
	// Secondary specifies the secondary state for admin key kind.
	Secondary AdminKeyKind = "secondary"
)

// HostingMode enumerates the values for hosting mode.
type HostingMode string

const (
	// Default specifies the default state for hosting mode.
	Default HostingMode = "default"
	// HighDensity specifies the high density state for hosting mode.
	HighDensity HostingMode = "highDensity"
)

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Failed specifies the failed state for provisioning state.
	Failed ProvisioningState = "failed"
	// Provisioning specifies the provisioning state for provisioning state.
	Provisioning ProvisioningState = "provisioning"
	// Succeeded specifies the succeeded state for provisioning state.
	Succeeded ProvisioningState = "succeeded"
)

// ServiceStatus enumerates the values for service status.
type ServiceStatus string

const (
	// ServiceStatusDegraded specifies the service status degraded state for service status.
	ServiceStatusDegraded ServiceStatus = "degraded"
	// ServiceStatusDeleting specifies the service status deleting state for service status.
	ServiceStatusDeleting ServiceStatus = "deleting"
	// ServiceStatusDisabled specifies the service status disabled state for service status.
	ServiceStatusDisabled ServiceStatus = "disabled"
	// ServiceStatusError specifies the service status error state for service status.
	ServiceStatusError ServiceStatus = "error"
	// ServiceStatusProvisioning specifies the service status provisioning state for service status.
	ServiceStatusProvisioning ServiceStatus = "provisioning"
	// ServiceStatusRunning specifies the service status running state for service status.
	ServiceStatusRunning ServiceStatus = "running"
)

// SkuName enumerates the values for sku name.
type SkuName string

const (
	// Basic specifies the basic state for sku name.
	Basic SkuName = "basic"
	// Free specifies the free state for sku name.
	Free SkuName = "free"
	// Standard specifies the standard state for sku name.
	Standard SkuName = "standard"
	// Standard2 specifies the standard 2 state for sku name.
	Standard2 SkuName = "standard2"
	// Standard3 specifies the standard 3 state for sku name.
	Standard3 SkuName = "standard3"
)

// UnavailableNameReason enumerates the values for unavailable name reason.
type UnavailableNameReason string

const (
	// AlreadyExists specifies the already exists state for unavailable name reason.
	AlreadyExists UnavailableNameReason = "AlreadyExists"
	// Invalid specifies the invalid state for unavailable name reason.
	Invalid UnavailableNameReason = "Invalid"
)

// AdminKeyResult is response containing the primary and secondary admin API keys for a given Azure Search service.
type AdminKeyResult struct {
	autorest.Response `json:"-"`
	PrimaryKey        *string `json:"primaryKey,omitempty"`
	SecondaryKey      *string `json:"secondaryKey,omitempty"`
}

// CheckNameAvailabilityInput is input of check name availability API.
type CheckNameAvailabilityInput struct {
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// CheckNameAvailabilityOutput is output of check name availability API.
type CheckNameAvailabilityOutput struct {
	autorest.Response `json:"-"`
	IsNameAvailable   *bool                 `json:"nameAvailable,omitempty"`
	Reason            UnavailableNameReason `json:"reason,omitempty"`
	Message           *string               `json:"message,omitempty"`
}

// CloudError is contains information about an API error.
type CloudError struct {
	Error *CloudErrorBody `json:"error,omitempty"`
}

// CloudErrorBody is describes a particular API error with an error code and a message.
type CloudErrorBody struct {
	Code    *string           `json:"code,omitempty"`
	Message *string           `json:"message,omitempty"`
	Target  *string           `json:"target,omitempty"`
	Details *[]CloudErrorBody `json:"details,omitempty"`
}

// ListQueryKeysResult is response containing the query API keys for a given Azure Search service.
type ListQueryKeysResult struct {
	autorest.Response `json:"-"`
	Value             *[]QueryKey `json:"value,omitempty"`
}

// QueryKey is describes an API key for a given Azure Search service that has permissions for query operations only.
type QueryKey struct {
	autorest.Response `json:"-"`
	Name              *string `json:"name,omitempty"`
	Key               *string `json:"key,omitempty"`
}

// Resource is base type for all Azure resources.
type Resource struct {
	ID       *string             `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Location *string             `json:"location,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
}

// Service is describes an Azure Search service and its current state.
type Service struct {
	autorest.Response  `json:"-"`
	ID                 *string             `json:"id,omitempty"`
	Name               *string             `json:"name,omitempty"`
	Type               *string             `json:"type,omitempty"`
	Location           *string             `json:"location,omitempty"`
	Tags               *map[string]*string `json:"tags,omitempty"`
	*ServiceProperties `json:"properties,omitempty"`
	Sku                *Sku `json:"sku,omitempty"`
}

// ServiceListResult is response containing a list of Azure Search services.
type ServiceListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Service `json:"value,omitempty"`
}

// ServiceProperties is properties of the Search service.
type ServiceProperties struct {
	ReplicaCount      *int32            `json:"replicaCount,omitempty"`
	PartitionCount    *int32            `json:"partitionCount,omitempty"`
	HostingMode       HostingMode       `json:"hostingMode,omitempty"`
	Status            ServiceStatus     `json:"status,omitempty"`
	StatusDetails     *string           `json:"statusDetails,omitempty"`
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
}

// Sku is defines the SKU of an Azure Search Service, which determines price tier and capacity limits.
type Sku struct {
	Name SkuName `json:"name,omitempty"`
}
