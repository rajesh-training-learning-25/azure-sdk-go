//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azadmin

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type FullBackupOperation.
func (f FullBackupOperation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "azureStorageBlobContainerUri", f.AzureStorageBlobContainerURI)
	populateTimeUnix(objectMap, "endTime", f.EndTime)
	populate(objectMap, "error", f.Error)
	populate(objectMap, "jobId", f.JobID)
	populateTimeUnix(objectMap, "startTime", f.StartTime)
	populate(objectMap, "status", f.Status)
	populate(objectMap, "statusDetails", f.StatusDetails)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type FullBackupOperation.
func (f *FullBackupOperation) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", f, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "azureStorageBlobContainerUri":
			err = unpopulate(val, "AzureStorageBlobContainerURI", &f.AzureStorageBlobContainerURI)
			delete(rawMsg, key)
		case "endTime":
			err = unpopulateTimeUnix(val, "EndTime", &f.EndTime)
			delete(rawMsg, key)
		case "error":
			err = unpopulate(val, "Error", &f.Error)
			delete(rawMsg, key)
		case "jobId":
			err = unpopulate(val, "JobID", &f.JobID)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateTimeUnix(val, "StartTime", &f.StartTime)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &f.Status)
			delete(rawMsg, key)
		case "statusDetails":
			err = unpopulate(val, "StatusDetails", &f.StatusDetails)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", f, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Permission.
func (p Permission) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "actions", p.Actions)
	populate(objectMap, "dataActions", p.DataActions)
	populate(objectMap, "notActions", p.NotActions)
	populate(objectMap, "notDataActions", p.NotDataActions)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Permission.
func (p *Permission) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "actions":
			err = unpopulate(val, "Actions", &p.Actions)
			delete(rawMsg, key)
		case "dataActions":
			err = unpopulate(val, "DataActions", &p.DataActions)
			delete(rawMsg, key)
		case "notActions":
			err = unpopulate(val, "NotActions", &p.NotActions)
			delete(rawMsg, key)
		case "notDataActions":
			err = unpopulate(val, "NotDataActions", &p.NotDataActions)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RestoreOperation.
func (r RestoreOperation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeUnix(objectMap, "endTime", r.EndTime)
	populate(objectMap, "error", r.Error)
	populate(objectMap, "jobId", r.JobID)
	populateTimeUnix(objectMap, "startTime", r.StartTime)
	populate(objectMap, "status", r.Status)
	populate(objectMap, "statusDetails", r.StatusDetails)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RestoreOperation.
func (r *RestoreOperation) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "endTime":
			err = unpopulateTimeUnix(val, "EndTime", &r.EndTime)
			delete(rawMsg, key)
		case "error":
			err = unpopulate(val, "Error", &r.Error)
			delete(rawMsg, key)
		case "jobId":
			err = unpopulate(val, "JobID", &r.JobID)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateTimeUnix(val, "StartTime", &r.StartTime)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &r.Status)
			delete(rawMsg, key)
		case "statusDetails":
			err = unpopulate(val, "StatusDetails", &r.StatusDetails)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RestoreOperationParameters.
func (r RestoreOperationParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "folderToRestore", r.FolderToRestore)
	populate(objectMap, "sasTokenParameters", r.SasTokenParameters)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RestoreOperationParameters.
func (r *RestoreOperationParameters) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "folderToRestore":
			err = unpopulate(val, "FolderToRestore", &r.FolderToRestore)
			delete(rawMsg, key)
		case "sasTokenParameters":
			err = unpopulate(val, "SasTokenParameters", &r.SasTokenParameters)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RoleAssignment.
func (r RoleAssignment) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", r.ID)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "properties", r.Properties)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RoleAssignment.
func (r *RoleAssignment) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &r.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &r.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &r.Properties)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &r.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RoleAssignmentCreateParameters.
func (r RoleAssignmentCreateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", r.Properties)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RoleAssignmentCreateParameters.
func (r *RoleAssignmentCreateParameters) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "properties":
			err = unpopulate(val, "Properties", &r.Properties)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RoleAssignmentFilter.
func (r RoleAssignmentFilter) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "principalId", r.PrincipalID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RoleAssignmentFilter.
func (r *RoleAssignmentFilter) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "principalId":
			err = unpopulate(val, "PrincipalID", &r.PrincipalID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RoleAssignmentListResult.
func (r RoleAssignmentListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", r.NextLink)
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RoleAssignmentListResult.
func (r *RoleAssignmentListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &r.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &r.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RoleAssignmentProperties.
func (r RoleAssignmentProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "principalId", r.PrincipalID)
	populate(objectMap, "roleDefinitionId", r.RoleDefinitionID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RoleAssignmentProperties.
func (r *RoleAssignmentProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "principalId":
			err = unpopulate(val, "PrincipalID", &r.PrincipalID)
			delete(rawMsg, key)
		case "roleDefinitionId":
			err = unpopulate(val, "RoleDefinitionID", &r.RoleDefinitionID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RoleAssignmentPropertiesWithScope.
func (r RoleAssignmentPropertiesWithScope) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "principalId", r.PrincipalID)
	populate(objectMap, "roleDefinitionId", r.RoleDefinitionID)
	populate(objectMap, "scope", r.Scope)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RoleAssignmentPropertiesWithScope.
func (r *RoleAssignmentPropertiesWithScope) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "principalId":
			err = unpopulate(val, "PrincipalID", &r.PrincipalID)
			delete(rawMsg, key)
		case "roleDefinitionId":
			err = unpopulate(val, "RoleDefinitionID", &r.RoleDefinitionID)
			delete(rawMsg, key)
		case "scope":
			err = unpopulate(val, "Scope", &r.Scope)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RoleDefinition.
func (r RoleDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", r.ID)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "properties", r.Properties)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RoleDefinition.
func (r *RoleDefinition) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &r.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &r.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &r.Properties)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &r.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RoleDefinitionCreateParameters.
func (r RoleDefinitionCreateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", r.Properties)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RoleDefinitionCreateParameters.
func (r *RoleDefinitionCreateParameters) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "properties":
			err = unpopulate(val, "Properties", &r.Properties)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RoleDefinitionFilter.
func (r RoleDefinitionFilter) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "roleName", r.RoleName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RoleDefinitionFilter.
func (r *RoleDefinitionFilter) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "roleName":
			err = unpopulate(val, "RoleName", &r.RoleName)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RoleDefinitionListResult.
func (r RoleDefinitionListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", r.NextLink)
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RoleDefinitionListResult.
func (r *RoleDefinitionListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &r.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &r.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RoleDefinitionProperties.
func (r RoleDefinitionProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "assignableScopes", r.AssignableScopes)
	populate(objectMap, "description", r.Description)
	populate(objectMap, "permissions", r.Permissions)
	populate(objectMap, "roleName", r.RoleName)
	populate(objectMap, "type", r.RoleType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RoleDefinitionProperties.
func (r *RoleDefinitionProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "assignableScopes":
			err = unpopulate(val, "AssignableScopes", &r.AssignableScopes)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, "Description", &r.Description)
			delete(rawMsg, key)
		case "permissions":
			err = unpopulate(val, "Permissions", &r.Permissions)
			delete(rawMsg, key)
		case "roleName":
			err = unpopulate(val, "RoleName", &r.RoleName)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "RoleType", &r.RoleType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SASTokenParameter.
func (s SASTokenParameter) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "storageResourceUri", s.StorageResourceURI)
	populate(objectMap, "token", s.Token)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SASTokenParameter.
func (s *SASTokenParameter) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "storageResourceUri":
			err = unpopulate(val, "StorageResourceURI", &s.StorageResourceURI)
			delete(rawMsg, key)
		case "token":
			err = unpopulate(val, "Token", &s.Token)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SelectiveKeyRestoreOperation.
func (s SelectiveKeyRestoreOperation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeUnix(objectMap, "endTime", s.EndTime)
	populate(objectMap, "error", s.Error)
	populate(objectMap, "jobId", s.JobID)
	populateTimeUnix(objectMap, "startTime", s.StartTime)
	populate(objectMap, "status", s.Status)
	populate(objectMap, "statusDetails", s.StatusDetails)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SelectiveKeyRestoreOperation.
func (s *SelectiveKeyRestoreOperation) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "endTime":
			err = unpopulateTimeUnix(val, "EndTime", &s.EndTime)
			delete(rawMsg, key)
		case "error":
			err = unpopulate(val, "Error", &s.Error)
			delete(rawMsg, key)
		case "jobId":
			err = unpopulate(val, "JobID", &s.JobID)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateTimeUnix(val, "StartTime", &s.StartTime)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &s.Status)
			delete(rawMsg, key)
		case "statusDetails":
			err = unpopulate(val, "StatusDetails", &s.StatusDetails)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SelectiveKeyRestoreOperationParameters.
func (s SelectiveKeyRestoreOperationParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "folder", s.Folder)
	populate(objectMap, "sasTokenParameters", s.SasTokenParameters)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SelectiveKeyRestoreOperationParameters.
func (s *SelectiveKeyRestoreOperationParameters) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "folder":
			err = unpopulate(val, "Folder", &s.Folder)
			delete(rawMsg, key)
		case "sasTokenParameters":
			err = unpopulate(val, "SasTokenParameters", &s.SasTokenParameters)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Setting.
func (s Setting) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "name", s.Name)
	populate(objectMap, "type", s.Type)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Setting.
func (s *Setting) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "name":
			err = unpopulate(val, "Name", &s.Name)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &s.Type)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &s.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SettingsListResult.
func (s SettingsListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "settings", s.Settings)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SettingsListResult.
func (s *SettingsListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "settings":
			err = unpopulate(val, "Settings", &s.Settings)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type UpdateSettingRequest.
func (u UpdateSettingRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", u.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type UpdateSettingRequest.
func (u *UpdateSettingRequest) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", u, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "value":
			err = unpopulate(val, "Value", &u.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", u, err)
		}
	}
	return nil
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v interface{}) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
