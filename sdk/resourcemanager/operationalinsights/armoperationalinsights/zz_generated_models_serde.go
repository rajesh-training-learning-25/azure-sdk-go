//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoperationalinsights

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type Cluster.
func (c Cluster) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", c.ID)
	populate(objectMap, "identity", c.Identity)
	populate(objectMap, "location", c.Location)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "properties", c.Properties)
	populate(objectMap, "sku", c.SKU)
	populate(objectMap, "tags", c.Tags)
	populate(objectMap, "type", c.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ClusterPatch.
func (c ClusterPatch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "identity", c.Identity)
	populate(objectMap, "properties", c.Properties)
	populate(objectMap, "sku", c.SKU)
	populate(objectMap, "tags", c.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ClusterProperties.
func (c ClusterProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "associatedWorkspaces", c.AssociatedWorkspaces)
	populate(objectMap, "billingType", c.BillingType)
	populate(objectMap, "capacityReservationProperties", c.CapacityReservationProperties)
	populate(objectMap, "clusterId", c.ClusterID)
	populate(objectMap, "createdDate", c.CreatedDate)
	populate(objectMap, "isAvailabilityZonesEnabled", c.IsAvailabilityZonesEnabled)
	populate(objectMap, "isDoubleEncryptionEnabled", c.IsDoubleEncryptionEnabled)
	populate(objectMap, "keyVaultProperties", c.KeyVaultProperties)
	populate(objectMap, "lastModifiedDate", c.LastModifiedDate)
	populate(objectMap, "provisioningState", c.ProvisioningState)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DataExportProperties.
func (d DataExportProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "createdDate", d.CreatedDate)
	populate(objectMap, "dataExportId", d.DataExportID)
	populate(objectMap, "destination", d.Destination)
	populate(objectMap, "enable", d.Enable)
	populate(objectMap, "lastModifiedDate", d.LastModifiedDate)
	populate(objectMap, "tableNames", d.TableNames)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DataSource.
func (d DataSource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", d.Etag)
	populate(objectMap, "id", d.ID)
	populate(objectMap, "kind", d.Kind)
	populate(objectMap, "name", d.Name)
	populate(objectMap, "properties", &d.Properties)
	populate(objectMap, "tags", d.Tags)
	populate(objectMap, "type", d.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Identity.
func (i Identity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "principalId", i.PrincipalID)
	populate(objectMap, "tenantId", i.TenantID)
	populate(objectMap, "type", i.Type)
	populate(objectMap, "userAssignedIdentities", i.UserAssignedIdentities)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type LinkedService.
func (l LinkedService) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", l.ID)
	populate(objectMap, "name", l.Name)
	populate(objectMap, "properties", l.Properties)
	populate(objectMap, "tags", l.Tags)
	populate(objectMap, "type", l.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type LinkedStorageAccountsProperties.
func (l LinkedStorageAccountsProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "dataSourceType", l.DataSourceType)
	populate(objectMap, "storageAccountIds", l.StorageAccountIDs)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type LogAnalyticsQueryPack.
func (l LogAnalyticsQueryPack) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", l.ID)
	populate(objectMap, "location", l.Location)
	populate(objectMap, "name", l.Name)
	populate(objectMap, "properties", l.Properties)
	populate(objectMap, "tags", l.Tags)
	populate(objectMap, "type", l.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type LogAnalyticsQueryPackProperties.
func (l LogAnalyticsQueryPackProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "provisioningState", l.ProvisioningState)
	populate(objectMap, "queryPackId", l.QueryPackID)
	populateTimeRFC3339(objectMap, "timeCreated", l.TimeCreated)
	populateTimeRFC3339(objectMap, "timeModified", l.TimeModified)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LogAnalyticsQueryPackProperties.
func (l *LogAnalyticsQueryPackProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &l.ProvisioningState)
			delete(rawMsg, key)
		case "queryPackId":
			err = unpopulate(val, "QueryPackID", &l.QueryPackID)
			delete(rawMsg, key)
		case "timeCreated":
			err = unpopulateTimeRFC3339(val, "TimeCreated", &l.TimeCreated)
			delete(rawMsg, key)
		case "timeModified":
			err = unpopulateTimeRFC3339(val, "TimeModified", &l.TimeModified)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type LogAnalyticsQueryPackQuery.
func (l LogAnalyticsQueryPackQuery) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", l.ID)
	populate(objectMap, "name", l.Name)
	populate(objectMap, "properties", l.Properties)
	populate(objectMap, "systemData", l.SystemData)
	populate(objectMap, "type", l.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type LogAnalyticsQueryPackQueryProperties.
func (l LogAnalyticsQueryPackQueryProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "author", l.Author)
	populate(objectMap, "body", l.Body)
	populate(objectMap, "description", l.Description)
	populate(objectMap, "displayName", l.DisplayName)
	populate(objectMap, "id", l.ID)
	populate(objectMap, "properties", &l.Properties)
	populate(objectMap, "related", l.Related)
	populate(objectMap, "tags", l.Tags)
	populateTimeRFC3339(objectMap, "timeCreated", l.TimeCreated)
	populateTimeRFC3339(objectMap, "timeModified", l.TimeModified)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LogAnalyticsQueryPackQueryProperties.
func (l *LogAnalyticsQueryPackQueryProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "author":
			err = unpopulate(val, "Author", &l.Author)
			delete(rawMsg, key)
		case "body":
			err = unpopulate(val, "Body", &l.Body)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, "Description", &l.Description)
			delete(rawMsg, key)
		case "displayName":
			err = unpopulate(val, "DisplayName", &l.DisplayName)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &l.ID)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &l.Properties)
			delete(rawMsg, key)
		case "related":
			err = unpopulate(val, "Related", &l.Related)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &l.Tags)
			delete(rawMsg, key)
		case "timeCreated":
			err = unpopulateTimeRFC3339(val, "TimeCreated", &l.TimeCreated)
			delete(rawMsg, key)
		case "timeModified":
			err = unpopulateTimeRFC3339(val, "TimeModified", &l.TimeModified)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type LogAnalyticsQueryPackQueryPropertiesRelated.
func (l LogAnalyticsQueryPackQueryPropertiesRelated) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "categories", l.Categories)
	populate(objectMap, "resourceTypes", l.ResourceTypes)
	populate(objectMap, "solutions", l.Solutions)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type LogAnalyticsQueryPackQuerySearchProperties.
func (l LogAnalyticsQueryPackQuerySearchProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "related", l.Related)
	populate(objectMap, "tags", l.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type LogAnalyticsQueryPackQuerySearchPropertiesRelated.
func (l LogAnalyticsQueryPackQuerySearchPropertiesRelated) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "categories", l.Categories)
	populate(objectMap, "resourceTypes", l.ResourceTypes)
	populate(objectMap, "solutions", l.Solutions)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ManagementGroupProperties.
func (m *ManagementGroupProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "created":
			err = unpopulateTimeRFC3339(val, "Created", &m.Created)
			delete(rawMsg, key)
		case "dataReceived":
			err = unpopulateTimeRFC3339(val, "DataReceived", &m.DataReceived)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &m.ID)
			delete(rawMsg, key)
		case "isGateway":
			err = unpopulate(val, "IsGateway", &m.IsGateway)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &m.Name)
			delete(rawMsg, key)
		case "sku":
			err = unpopulate(val, "SKU", &m.SKU)
			delete(rawMsg, key)
		case "serverCount":
			err = unpopulate(val, "ServerCount", &m.ServerCount)
			delete(rawMsg, key)
		case "version":
			err = unpopulate(val, "Version", &m.Version)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type QueryPacksResource.
func (q QueryPacksResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", q.ID)
	populate(objectMap, "location", q.Location)
	populate(objectMap, "name", q.Name)
	populate(objectMap, "tags", q.Tags)
	populate(objectMap, "type", q.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type RestoredLogs.
func (r RestoredLogs) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "endRestoreTime", r.EndRestoreTime)
	populate(objectMap, "sourceTable", r.SourceTable)
	populateTimeRFC3339(objectMap, "startRestoreTime", r.StartRestoreTime)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RestoredLogs.
func (r *RestoredLogs) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "endRestoreTime":
			err = unpopulateTimeRFC3339(val, "EndRestoreTime", &r.EndRestoreTime)
			delete(rawMsg, key)
		case "sourceTable":
			err = unpopulate(val, "SourceTable", &r.SourceTable)
			delete(rawMsg, key)
		case "startRestoreTime":
			err = unpopulateTimeRFC3339(val, "StartRestoreTime", &r.StartRestoreTime)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SavedSearchProperties.
func (s SavedSearchProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "category", s.Category)
	populate(objectMap, "displayName", s.DisplayName)
	populate(objectMap, "functionAlias", s.FunctionAlias)
	populate(objectMap, "functionParameters", s.FunctionParameters)
	populate(objectMap, "query", s.Query)
	populate(objectMap, "tags", s.Tags)
	populate(objectMap, "version", s.Version)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Schema.
func (s Schema) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "categories", s.Categories)
	populate(objectMap, "columns", s.Columns)
	populate(objectMap, "description", s.Description)
	populate(objectMap, "displayName", s.DisplayName)
	populate(objectMap, "labels", s.Labels)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "restoredLogs", s.RestoredLogs)
	populate(objectMap, "searchResults", s.SearchResults)
	populate(objectMap, "solutions", s.Solutions)
	populate(objectMap, "source", s.Source)
	populate(objectMap, "standardColumns", s.StandardColumns)
	populate(objectMap, "tableSubType", s.TableSubType)
	populate(objectMap, "tableType", s.TableType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SearchMetadata.
func (s *SearchMetadata) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "aggregatedGroupingFields":
			err = unpopulate(val, "AggregatedGroupingFields", &s.AggregatedGroupingFields)
			delete(rawMsg, key)
		case "aggregatedValueField":
			err = unpopulate(val, "AggregatedValueField", &s.AggregatedValueField)
			delete(rawMsg, key)
		case "coreSummaries":
			err = unpopulate(val, "CoreSummaries", &s.CoreSummaries)
			delete(rawMsg, key)
		case "eTag":
			err = unpopulate(val, "ETag", &s.ETag)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &s.ID)
			delete(rawMsg, key)
		case "lastUpdated":
			err = unpopulateTimeRFC3339(val, "LastUpdated", &s.LastUpdated)
			delete(rawMsg, key)
		case "max":
			err = unpopulate(val, "Max", &s.Max)
			delete(rawMsg, key)
		case "requestTime":
			err = unpopulate(val, "RequestTime", &s.RequestTime)
			delete(rawMsg, key)
		case "resultType":
			err = unpopulate(val, "ResultType", &s.ResultType)
			delete(rawMsg, key)
		case "schema":
			err = unpopulate(val, "Schema", &s.Schema)
			delete(rawMsg, key)
		case "requestId":
			err = unpopulate(val, "SearchID", &s.SearchID)
			delete(rawMsg, key)
		case "sort":
			err = unpopulate(val, "Sort", &s.Sort)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateTimeRFC3339(val, "StartTime", &s.StartTime)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &s.Status)
			delete(rawMsg, key)
		case "sum":
			err = unpopulate(val, "Sum", &s.Sum)
			delete(rawMsg, key)
		case "top":
			err = unpopulate(val, "Top", &s.Top)
			delete(rawMsg, key)
		case "total":
			err = unpopulate(val, "Total", &s.Total)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SearchResults.
func (s SearchResults) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "description", s.Description)
	populateTimeRFC3339(objectMap, "endSearchTime", s.EndSearchTime)
	populate(objectMap, "limit", s.Limit)
	populate(objectMap, "query", s.Query)
	populate(objectMap, "sourceTable", s.SourceTable)
	populateTimeRFC3339(objectMap, "startSearchTime", s.StartSearchTime)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SearchResults.
func (s *SearchResults) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "description":
			err = unpopulate(val, "Description", &s.Description)
			delete(rawMsg, key)
		case "endSearchTime":
			err = unpopulateTimeRFC3339(val, "EndSearchTime", &s.EndSearchTime)
			delete(rawMsg, key)
		case "limit":
			err = unpopulate(val, "Limit", &s.Limit)
			delete(rawMsg, key)
		case "query":
			err = unpopulate(val, "Query", &s.Query)
			delete(rawMsg, key)
		case "sourceTable":
			err = unpopulate(val, "SourceTable", &s.SourceTable)
			delete(rawMsg, key)
		case "startSearchTime":
			err = unpopulateTimeRFC3339(val, "StartSearchTime", &s.StartSearchTime)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type StorageInsight.
func (s StorageInsight) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "eTag", s.ETag)
	populate(objectMap, "id", s.ID)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "tags", s.Tags)
	populate(objectMap, "type", s.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type StorageInsightProperties.
func (s StorageInsightProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "containers", s.Containers)
	populate(objectMap, "status", s.Status)
	populate(objectMap, "storageAccount", s.StorageAccount)
	populate(objectMap, "tables", s.Tables)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SystemData.
func (s SystemData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemData.
func (s *SystemData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, "CreatedAt", &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, "CreatedBy", &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, "CreatedByType", &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateTimeRFC3339(val, "LastModifiedAt", &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, "LastModifiedBy", &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, "LastModifiedByType", &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SystemDataAutoGenerated.
func (s SystemDataAutoGenerated) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemDataAutoGenerated.
func (s *SystemDataAutoGenerated) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, "CreatedAt", &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, "CreatedBy", &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, "CreatedByType", &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateTimeRFC3339(val, "LastModifiedAt", &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, "LastModifiedBy", &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, "LastModifiedByType", &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Table.
func (t Table) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", t.ID)
	populate(objectMap, "name", t.Name)
	populate(objectMap, "properties", t.Properties)
	populate(objectMap, "systemData", t.SystemData)
	populate(objectMap, "type", t.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type TagsResource.
func (t TagsResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "tags", t.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type TrackedResource.
func (t TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", t.ID)
	populate(objectMap, "location", t.Location)
	populate(objectMap, "name", t.Name)
	populate(objectMap, "tags", t.Tags)
	populate(objectMap, "type", t.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type UsageMetric.
func (u *UsageMetric) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", u, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "currentValue":
			err = unpopulate(val, "CurrentValue", &u.CurrentValue)
			delete(rawMsg, key)
		case "limit":
			err = unpopulate(val, "Limit", &u.Limit)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &u.Name)
			delete(rawMsg, key)
		case "nextResetTime":
			err = unpopulateTimeRFC3339(val, "NextResetTime", &u.NextResetTime)
			delete(rawMsg, key)
		case "quotaPeriod":
			err = unpopulate(val, "QuotaPeriod", &u.QuotaPeriod)
			delete(rawMsg, key)
		case "unit":
			err = unpopulate(val, "Unit", &u.Unit)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", u, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Workspace.
func (w Workspace) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "eTag", w.ETag)
	populate(objectMap, "id", w.ID)
	populate(objectMap, "location", w.Location)
	populate(objectMap, "name", w.Name)
	populate(objectMap, "properties", w.Properties)
	populate(objectMap, "systemData", w.SystemData)
	populate(objectMap, "tags", w.Tags)
	populate(objectMap, "type", w.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type WorkspaceFeatures.
func (w WorkspaceFeatures) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "clusterResourceId", w.ClusterResourceID)
	populate(objectMap, "disableLocalAuth", w.DisableLocalAuth)
	populate(objectMap, "enableDataExport", w.EnableDataExport)
	populate(objectMap, "enableLogAccessUsingOnlyResourcePermissions", w.EnableLogAccessUsingOnlyResourcePermissions)
	populate(objectMap, "immediatePurgeDataOn30Days", w.ImmediatePurgeDataOn30Days)
	if w.AdditionalProperties != nil {
		for key, val := range w.AdditionalProperties {
			objectMap[key] = val
		}
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type WorkspaceFeatures.
func (w *WorkspaceFeatures) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", w, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "clusterResourceId":
			err = unpopulate(val, "ClusterResourceID", &w.ClusterResourceID)
			delete(rawMsg, key)
		case "disableLocalAuth":
			err = unpopulate(val, "DisableLocalAuth", &w.DisableLocalAuth)
			delete(rawMsg, key)
		case "enableDataExport":
			err = unpopulate(val, "EnableDataExport", &w.EnableDataExport)
			delete(rawMsg, key)
		case "enableLogAccessUsingOnlyResourcePermissions":
			err = unpopulate(val, "EnableLogAccessUsingOnlyResourcePermissions", &w.EnableLogAccessUsingOnlyResourcePermissions)
			delete(rawMsg, key)
		case "immediatePurgeDataOn30Days":
			err = unpopulate(val, "ImmediatePurgeDataOn30Days", &w.ImmediatePurgeDataOn30Days)
			delete(rawMsg, key)
		default:
			if w.AdditionalProperties == nil {
				w.AdditionalProperties = map[string]interface{}{}
			}
			if val != nil {
				var aux interface{}
				err = json.Unmarshal(val, &aux)
				w.AdditionalProperties[key] = aux
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", w, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type WorkspacePatch.
func (w WorkspacePatch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", w.Etag)
	populate(objectMap, "id", w.ID)
	populate(objectMap, "name", w.Name)
	populate(objectMap, "properties", w.Properties)
	populate(objectMap, "tags", w.Tags)
	populate(objectMap, "type", w.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type WorkspaceProperties.
func (w WorkspaceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "createdDate", w.CreatedDate)
	populate(objectMap, "customerId", w.CustomerID)
	populate(objectMap, "defaultDataCollectionRuleResourceId", w.DefaultDataCollectionRuleResourceID)
	populate(objectMap, "features", w.Features)
	populate(objectMap, "forceCmkForQuery", w.ForceCmkForQuery)
	populate(objectMap, "modifiedDate", w.ModifiedDate)
	populate(objectMap, "privateLinkScopedResources", w.PrivateLinkScopedResources)
	populate(objectMap, "provisioningState", w.ProvisioningState)
	populate(objectMap, "publicNetworkAccessForIngestion", w.PublicNetworkAccessForIngestion)
	populate(objectMap, "publicNetworkAccessForQuery", w.PublicNetworkAccessForQuery)
	populate(objectMap, "retentionInDays", w.RetentionInDays)
	populate(objectMap, "sku", w.SKU)
	populate(objectMap, "workspaceCapping", w.WorkspaceCapping)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type WorkspacePurgeBody.
func (w WorkspacePurgeBody) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "filters", w.Filters)
	populate(objectMap, "table", w.Table)
	return json.Marshal(objectMap)
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
