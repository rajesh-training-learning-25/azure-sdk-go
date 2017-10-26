package account

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
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

// DataLakeStoreAccountState enumerates the values for data lake store account state.
type DataLakeStoreAccountState string

const (
	// Active specifies the active state for data lake store account state.
	Active DataLakeStoreAccountState = "active"
	// Suspended specifies the suspended state for data lake store account state.
	Suspended DataLakeStoreAccountState = "suspended"
)

// DataLakeStoreAccountStatus enumerates the values for data lake store account status.
type DataLakeStoreAccountStatus string

const (
	// Creating specifies the creating state for data lake store account status.
	Creating DataLakeStoreAccountStatus = "Creating"
	// Deleted specifies the deleted state for data lake store account status.
	Deleted DataLakeStoreAccountStatus = "Deleted"
	// Deleting specifies the deleting state for data lake store account status.
	Deleting DataLakeStoreAccountStatus = "Deleting"
	// Failed specifies the failed state for data lake store account status.
	Failed DataLakeStoreAccountStatus = "Failed"
	// Patching specifies the patching state for data lake store account status.
	Patching DataLakeStoreAccountStatus = "Patching"
	// Resuming specifies the resuming state for data lake store account status.
	Resuming DataLakeStoreAccountStatus = "Resuming"
	// Running specifies the running state for data lake store account status.
	Running DataLakeStoreAccountStatus = "Running"
	// Succeeded specifies the succeeded state for data lake store account status.
	Succeeded DataLakeStoreAccountStatus = "Succeeded"
	// Suspending specifies the suspending state for data lake store account status.
	Suspending DataLakeStoreAccountStatus = "Suspending"
)

// EncryptionConfigType enumerates the values for encryption config type.
type EncryptionConfigType string

const (
	// ServiceManaged specifies the service managed state for encryption config type.
	ServiceManaged EncryptionConfigType = "ServiceManaged"
	// UserManaged specifies the user managed state for encryption config type.
	UserManaged EncryptionConfigType = "UserManaged"
)

// EncryptionIdentityType enumerates the values for encryption identity type.
type EncryptionIdentityType string

const (
	// SystemAssigned specifies the system assigned state for encryption identity type.
	SystemAssigned EncryptionIdentityType = "SystemAssigned"
)

// EncryptionProvisioningState enumerates the values for encryption provisioning state.
type EncryptionProvisioningState string

const (
	// EncryptionProvisioningStateCreating specifies the encryption provisioning state creating state for encryption
	// provisioning state.
	EncryptionProvisioningStateCreating EncryptionProvisioningState = "Creating"
	// EncryptionProvisioningStateSucceeded specifies the encryption provisioning state succeeded state for encryption
	// provisioning state.
	EncryptionProvisioningStateSucceeded EncryptionProvisioningState = "Succeeded"
)

// EncryptionState enumerates the values for encryption state.
type EncryptionState string

const (
	// Disabled specifies the disabled state for encryption state.
	Disabled EncryptionState = "Disabled"
	// Enabled specifies the enabled state for encryption state.
	Enabled EncryptionState = "Enabled"
)

// OperationStatus enumerates the values for operation status.
type OperationStatus string

const (
	// OperationStatusFailed specifies the operation status failed state for operation status.
	OperationStatusFailed OperationStatus = "Failed"
	// OperationStatusInProgress specifies the operation status in progress state for operation status.
	OperationStatusInProgress OperationStatus = "InProgress"
	// OperationStatusSucceeded specifies the operation status succeeded state for operation status.
	OperationStatusSucceeded OperationStatus = "Succeeded"
)

// AzureAsyncOperationResult is the response body contains the status of the specified asynchronous operation,
// indicating whether it has succeeded, is in progress, or has failed. Note that this status is distinct from the HTTP
// status code returned for the Get Operation Status operation itself. If the asynchronous operation succeeded, the
// response body includes the HTTP status code for the successful request. If the asynchronous operation failed, the
// response body includes the HTTP status code for the failed request and error information regarding the failure.
type AzureAsyncOperationResult struct {
	Status OperationStatus `json:"status,omitempty"`
	Error  *Error          `json:"error,omitempty"`
}

// DataLakeStoreAccount is data Lake Store account information
type DataLakeStoreAccount struct {
	autorest.Response `json:"-"`
	Location          *string                         `json:"location,omitempty"`
	Name              *string                         `json:"name,omitempty"`
	Type              *string                         `json:"type,omitempty"`
	ID                *string                         `json:"id,omitempty"`
	Identity          *EncryptionIdentity             `json:"identity,omitempty"`
	Tags              *map[string]*string             `json:"tags,omitempty"`
	Properties        *DataLakeStoreAccountProperties `json:"properties,omitempty"`
}

// DataLakeStoreAccountListResult is data Lake Store account list information response.
type DataLakeStoreAccountListResult struct {
	autorest.Response `json:"-"`
	Value             *[]DataLakeStoreAccount `json:"value,omitempty"`
	NextLink          *string                 `json:"nextLink,omitempty"`
	Count             *int64                  `json:"count,omitempty"`
}

// DataLakeStoreAccountListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client DataLakeStoreAccountListResult) DataLakeStoreAccountListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// DataLakeStoreAccountProperties is data Lake Store account properties information
type DataLakeStoreAccountProperties struct {
	ProvisioningState           DataLakeStoreAccountStatus  `json:"provisioningState,omitempty"`
	State                       DataLakeStoreAccountState   `json:"state,omitempty"`
	CreationTime                *date.Time                  `json:"creationTime,omitempty"`
	EncryptionState             EncryptionState             `json:"encryptionState,omitempty"`
	EncryptionProvisioningState EncryptionProvisioningState `json:"encryptionProvisioningState,omitempty"`
	EncryptionConfig            *EncryptionConfig           `json:"encryptionConfig,omitempty"`
	LastModifiedTime            *date.Time                  `json:"lastModifiedTime,omitempty"`
	Endpoint                    *string                     `json:"endpoint,omitempty"`
	DefaultGroup                *string                     `json:"defaultGroup,omitempty"`
}

// DataLakeStoreFirewallRuleListResult is data Lake Store firewall rule list information.
type DataLakeStoreFirewallRuleListResult struct {
	autorest.Response `json:"-"`
	Value             *[]FirewallRule `json:"value,omitempty"`
	NextLink          *string         `json:"nextLink,omitempty"`
	Count             *int64          `json:"count,omitempty"`
}

// DataLakeStoreFirewallRuleListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client DataLakeStoreFirewallRuleListResult) DataLakeStoreFirewallRuleListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// EncryptionConfig is
type EncryptionConfig struct {
	Type             EncryptionConfigType `json:"type,omitempty"`
	KeyVaultMetaInfo *KeyVaultMetaInfo    `json:"keyVaultMetaInfo,omitempty"`
}

// EncryptionIdentity is
type EncryptionIdentity struct {
	Type        EncryptionIdentityType `json:"type,omitempty"`
	PrincipalID *uuid.UUID             `json:"principalId,omitempty"`
	TenantID    *uuid.UUID             `json:"tenantId,omitempty"`
}

// Error is data Lake Store error information
type Error struct {
	Code       *string         `json:"code,omitempty"`
	Message    *string         `json:"message,omitempty"`
	Target     *string         `json:"target,omitempty"`
	Details    *[]ErrorDetails `json:"details,omitempty"`
	InnerError *InnerError     `json:"innerError,omitempty"`
}

// ErrorDetails is data Lake Store error details information
type ErrorDetails struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Target  *string `json:"target,omitempty"`
}

// FirewallRule is data Lake Store firewall rule information
type FirewallRule struct {
	autorest.Response `json:"-"`
	Name              *string                 `json:"name,omitempty"`
	Type              *string                 `json:"type,omitempty"`
	ID                *string                 `json:"id,omitempty"`
	Location          *string                 `json:"location,omitempty"`
	Properties        *FirewallRuleProperties `json:"properties,omitempty"`
}

// FirewallRuleProperties is data Lake Store firewall rule properties information
type FirewallRuleProperties struct {
	StartIPAddress *string `json:"startIpAddress,omitempty"`
	EndIPAddress   *string `json:"endIpAddress,omitempty"`
}

// InnerError is data Lake Store inner error information
type InnerError struct {
	Trace   *string `json:"trace,omitempty"`
	Context *string `json:"context,omitempty"`
}

// KeyVaultMetaInfo is
type KeyVaultMetaInfo struct {
	KeyVaultResourceID   *string `json:"keyVaultResourceId,omitempty"`
	EncryptionKeyName    *string `json:"encryptionKeyName,omitempty"`
	EncryptionKeyVersion *string `json:"encryptionKeyVersion,omitempty"`
}
