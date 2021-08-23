//go:build go1.13
// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package internal

import (
	"net/http"
	"time"
)

// ServiceGetPropertiesResponse contains the response from method Service.GetProperties.
type ServiceGetPropertiesResponse struct {
	ServiceGetPropertiesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServiceGetPropertiesResult contains the result from method Service.GetProperties.
type ServiceGetPropertiesResult struct {
	TableServiceProperties
	// ClientRequestID contains the information returned from the x-ms-client-request-id header response.
	ClientRequestID *string `xml:"ClientRequestID"`

	// RequestID contains the information returned from the x-ms-request-id header response.
	RequestID *string `xml:"RequestID"`

	// Version contains the information returned from the x-ms-version header response.
	Version *string `xml:"Version"`
}

// ServiceGetStatisticsResponse contains the response from method Service.GetStatistics.
type ServiceGetStatisticsResponse struct {
	ServiceGetStatisticsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServiceGetStatisticsResult contains the result from method Service.GetStatistics.
type ServiceGetStatisticsResult struct {
	TableServiceStats
	// ClientRequestID contains the information returned from the x-ms-client-request-id header response.
	ClientRequestID *string `xml:"ClientRequestID"`

	// Date contains the information returned from the Date header response.
	Date *time.Time `xml:"Date"`

	// RequestID contains the information returned from the x-ms-request-id header response.
	RequestID *string `xml:"RequestID"`

	// Version contains the information returned from the x-ms-version header response.
	Version *string `xml:"Version"`
}

// ServiceSetPropertiesResponse contains the response from method Service.SetProperties.
type ServiceSetPropertiesResponse struct {
	ServiceSetPropertiesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServiceSetPropertiesResult contains the result from method Service.SetProperties.
type ServiceSetPropertiesResult struct {
	// ClientRequestID contains the information returned from the x-ms-client-request-id header response.
	ClientRequestID *string

	// RequestID contains the information returned from the x-ms-request-id header response.
	RequestID *string

	// Version contains the information returned from the x-ms-version header response.
	Version *string
}

// TableCreateResponse contains the response from method Table.Create.
type TableCreateResponse struct {
	TableCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TableCreateResult contains the result from method Table.Create.
type TableCreateResult struct {
	TableResponse
	// ClientRequestID contains the information returned from the x-ms-client-request-id header response.
	ClientRequestID *string

	// Date contains the information returned from the Date header response.
	Date *time.Time

	// PreferenceApplied contains the information returned from the Preference-Applied header response.
	PreferenceApplied *string

	// RequestID contains the information returned from the x-ms-request-id header response.
	RequestID *string

	// Version contains the information returned from the x-ms-version header response.
	Version *string
}

// TableDeleteEntityResponse contains the response from method Table.DeleteEntity.
type TableDeleteEntityResponse struct {
	TableDeleteEntityResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TableDeleteEntityResult contains the result from method Table.DeleteEntity.
type TableDeleteEntityResult struct {
	// ClientRequestID contains the information returned from the x-ms-client-request-id header response.
	ClientRequestID *string

	// Date contains the information returned from the Date header response.
	Date *time.Time

	// RequestID contains the information returned from the x-ms-request-id header response.
	RequestID *string

	// Version contains the information returned from the x-ms-version header response.
	Version *string
}

// TableDeleteResponse contains the response from method Table.Delete.
type TableDeleteResponse struct {
	TableDeleteResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TableDeleteResult contains the result from method Table.Delete.
type TableDeleteResult struct {
	// ClientRequestID contains the information returned from the x-ms-client-request-id header response.
	ClientRequestID *string

	// Date contains the information returned from the Date header response.
	Date *time.Time

	// RequestID contains the information returned from the x-ms-request-id header response.
	RequestID *string

	// Version contains the information returned from the x-ms-version header response.
	Version *string
}

// TableGetAccessPolicyResponse contains the response from method Table.GetAccessPolicy.
type TableGetAccessPolicyResponse struct {
	TableGetAccessPolicyResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TableGetAccessPolicyResult contains the result from method Table.GetAccessPolicy.
type TableGetAccessPolicyResult struct {
	// ClientRequestID contains the information returned from the x-ms-client-request-id header response.
	ClientRequestID *string `xml:"ClientRequestID"`

	// Date contains the information returned from the Date header response.
	Date *time.Time `xml:"Date"`

	// RequestID contains the information returned from the x-ms-request-id header response.
	RequestID *string `xml:"RequestID"`

	// A collection of signed identifiers.
	SignedIdentifiers []*SignedIdentifier `xml:"SignedIdentifier"`

	// Version contains the information returned from the x-ms-version header response.
	Version *string `xml:"Version"`
}

// TableInsertEntityResponse contains the response from method Table.InsertEntity.
type TableInsertEntityResponse struct {
	TableInsertEntityResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TableInsertEntityResult contains the result from method Table.InsertEntity.
type TableInsertEntityResult struct {
	// ClientRequestID contains the information returned from the x-ms-client-request-id header response.
	ClientRequestID *string

	// ContentType contains the information returned from the Content-Type header response.
	ContentType *string

	// Date contains the information returned from the Date header response.
	Date *time.Time

	// ETag contains the information returned from the ETag header response.
	ETag *string

	// PreferenceApplied contains the information returned from the Preference-Applied header response.
	PreferenceApplied *string

	// RequestID contains the information returned from the x-ms-request-id header response.
	RequestID *string

	// The other properties of the table entity.
	Value map[string]interface{}

	// Version contains the information returned from the x-ms-version header response.
	Version *string
}

// TableMergeEntityResponse contains the response from method Table.MergeEntity.
type TableMergeEntityResponse struct {
	TableMergeEntityResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TableMergeEntityResult contains the result from method Table.MergeEntity.
type TableMergeEntityResult struct {
	// ClientRequestID contains the information returned from the x-ms-client-request-id header response.
	ClientRequestID *string

	// Date contains the information returned from the Date header response.
	Date *time.Time

	// ETag contains the information returned from the ETag header response.
	ETag *string

	// RequestID contains the information returned from the x-ms-request-id header response.
	RequestID *string

	// Version contains the information returned from the x-ms-version header response.
	Version *string
}

// TableQueryEntitiesResponse contains the response from method Table.QueryEntities.
type TableQueryEntitiesResponse struct {
	TableQueryEntitiesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TableQueryEntitiesResult contains the result from method Table.QueryEntities.
type TableQueryEntitiesResult struct {
	TableEntityQueryResponse
	// ClientRequestID contains the information returned from the x-ms-client-request-id header response.
	ClientRequestID *string

	// Date contains the information returned from the Date header response.
	Date *time.Time

	// RequestID contains the information returned from the x-ms-request-id header response.
	RequestID *string

	// Version contains the information returned from the x-ms-version header response.
	Version *string

	// XMSContinuationNextPartitionKey contains the information returned from the x-ms-continuation-NextPartitionKey header response.
	XMSContinuationNextPartitionKey *string

	// XMSContinuationNextRowKey contains the information returned from the x-ms-continuation-NextRowKey header response.
	XMSContinuationNextRowKey *string
}

// TableQueryEntityWithPartitionAndRowKeyResponse contains the response from method Table.QueryEntityWithPartitionAndRowKey.
type TableQueryEntityWithPartitionAndRowKeyResponse struct {
	TableQueryEntityWithPartitionAndRowKeyResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TableQueryEntityWithPartitionAndRowKeyResult contains the result from method Table.QueryEntityWithPartitionAndRowKey.
type TableQueryEntityWithPartitionAndRowKeyResult struct {
	// ClientRequestID contains the information returned from the x-ms-client-request-id header response.
	ClientRequestID *string

	// Date contains the information returned from the Date header response.
	Date *time.Time

	// ETag contains the information returned from the ETag header response.
	ETag *string

	// RequestID contains the information returned from the x-ms-request-id header response.
	RequestID *string

	// The other properties of the table entity.
	Value map[string]interface{}

	// Version contains the information returned from the x-ms-version header response.
	Version *string

	// XMSContinuationNextPartitionKey contains the information returned from the x-ms-continuation-NextPartitionKey header response.
	XMSContinuationNextPartitionKey *string

	// XMSContinuationNextRowKey contains the information returned from the x-ms-continuation-NextRowKey header response.
	XMSContinuationNextRowKey *string
}

// TableQueryResponseEnvelope contains the response from method Table.Query.
type TableQueryResponseEnvelope struct {
	TableQueryResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TableQueryResult contains the result from method Table.Query.
type TableQueryResult struct {
	TableQueryResponse
	// ClientRequestID contains the information returned from the x-ms-client-request-id header response.
	ClientRequestID *string

	// Date contains the information returned from the Date header response.
	Date *time.Time

	// RequestID contains the information returned from the x-ms-request-id header response.
	RequestID *string

	// Version contains the information returned from the x-ms-version header response.
	Version *string

	// XMSContinuationNextTableName contains the information returned from the x-ms-continuation-NextTableName header response.
	XMSContinuationNextTableName *string
}

// TableSetAccessPolicyResponse contains the response from method Table.SetAccessPolicy.
type TableSetAccessPolicyResponse struct {
	TableSetAccessPolicyResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TableSetAccessPolicyResult contains the result from method Table.SetAccessPolicy.
type TableSetAccessPolicyResult struct {
	// ClientRequestID contains the information returned from the x-ms-client-request-id header response.
	ClientRequestID *string

	// Date contains the information returned from the Date header response.
	Date *time.Time

	// RequestID contains the information returned from the x-ms-request-id header response.
	RequestID *string

	// Version contains the information returned from the x-ms-version header response.
	Version *string
}

// TableUpdateEntityResponse contains the response from method Table.UpdateEntity.
type TableUpdateEntityResponse struct {
	TableUpdateEntityResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TableUpdateEntityResult contains the result from method Table.UpdateEntity.
type TableUpdateEntityResult struct {
	// ClientRequestID contains the information returned from the x-ms-client-request-id header response.
	ClientRequestID *string

	// Date contains the information returned from the Date header response.
	Date *time.Time

	// ETag contains the information returned from the ETag header response.
	ETag *string

	// RequestID contains the information returned from the x-ms-request-id header response.
	RequestID *string

	// Version contains the information returned from the x-ms-version header response.
	Version *string
}
