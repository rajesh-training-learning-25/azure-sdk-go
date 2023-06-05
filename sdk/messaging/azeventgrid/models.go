//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azeventgrid

import "time"

// AcknowledgeOptions - Array of lock token strings for the corresponding received Cloud Events to be acknowledged.
type AcknowledgeOptions struct {
	// REQUIRED; String array of lock tokens.
	LockTokens []*string `json:"lockTokens,omitempty"`
}

// AcknowledgeResult - The result of the Acknowledge operation.
type AcknowledgeResult struct {
	// REQUIRED; Array of LockToken values for failed cloud events. Each LockToken includes the lock token value along with the
	// related error information (namely, the error code and description).
	FailedLockTokens []*FailedLockToken `json:"failedLockTokens,omitempty"`

	// REQUIRED; Array of lock tokens values for the successfully acknowledged cloud events.
	SucceededLockTokens []*string `json:"succeededLockTokens,omitempty"`
}

// AzureCoreFoundationsError - The error object.
type AzureCoreFoundationsError struct {
	// REQUIRED; One of a server-defined set of error codes.
	Code *string `json:"code,omitempty"`

	// REQUIRED; An array of details about specific errors that led to this reported error.
	Details []*AzureCoreFoundationsError `json:"details,omitempty"`

	// REQUIRED; A human-readable representation of the error.
	Message *string `json:"message,omitempty"`

	// An object containing more specific information than the current object about the error.
	Innererror *AzureCoreFoundationsInnerError `json:"innererror,omitempty"`

	// The target of the error.
	Target *string `json:"target,omitempty"`
}

// AzureCoreFoundationsErrorResponse - A response containing error details.
type AzureCoreFoundationsErrorResponse struct {
	// REQUIRED; The error object.
	Error *AzureCoreFoundationsError `json:"error,omitempty"`
}

// AzureCoreFoundationsInnerError - An object containing more specific information about the error. As per Microsoft One API
// guidelines -
// https://github.com/Microsoft/api-guidelines/blob/vNext/Guidelines.md#7102-error-condition-responses.
type AzureCoreFoundationsInnerError struct {
	// REQUIRED; One of a server-defined set of error codes.
	Code *string `json:"code,omitempty"`

	// Inner error.
	Innererror *AzureCoreFoundationsInnerError `json:"innererror,omitempty"`
}

// BrokerProperties - Properties of the Event Broker operation.
type BrokerProperties struct {
	// REQUIRED; The attempt count for delivering the event.
	DeliveryCount *int32 `json:"deliveryCount,omitempty"`

	// REQUIRED; The token used to lock the event.
	LockToken *string `json:"lockToken,omitempty"`
}

// AcknowledgeCloudEventsOptions contains the optional parameters for the Client.AcknowledgeCloudEvents method.
type AcknowledgeCloudEventsOptions struct {
	// placeholder for future optional parameters
}

// PublishCloudEventsOptions contains the optional parameters for the Client.PublishCloudEvents method.
type PublishCloudEventsOptions struct {
	// placeholder for future optional parameters
}

// ReceiveCloudEventsOptions contains the optional parameters for the Client.ReceiveCloudEvents method.
type ReceiveCloudEventsOptions struct {
	// Max Events count to be received. Minimum value is 1, while maximum value is 100 events. If not specified, the default value
	// is 1.
	MaxEvents *int32
	// Max wait time value for receive operation in Seconds. It is the time in seconds that the server approximately waits for
	// the availability of an event and responds to the request. If an event is
	// available, the broker responds immediately to the client. Minimum value is 10 seconds, while maximum value is 120 seconds.
	// If not specified, the default value is 60 seconds.
	MaxWaitTime *int32
}

// RejectCloudEventsOptions contains the optional parameters for the Client.RejectCloudEvents method.
type RejectCloudEventsOptions struct {
	// placeholder for future optional parameters
}

// ReleaseCloudEventsOptions contains the optional parameters for the Client.ReleaseCloudEvents method.
type ReleaseCloudEventsOptions struct {
	// placeholder for future optional parameters
}

// CloudEvent - Properties of an event published to an Azure Messaging EventGrid Namespace topic using the CloudEvent 1.0
// Schema.
type CloudEvent struct {
	// REQUIRED; An identifier for the event. The combination of id and source must be unique for each distinct event.
	ID *string `json:"id,omitempty"`

	// REQUIRED; Identifies the context in which an event happened. The combination of id and source must be unique for each distinct
	// event.
	Source *string `json:"source,omitempty"`

	// REQUIRED; The version of the CloudEvents specification which the event uses.
	SpecVersion *string `json:"specversion,omitempty"`

	// REQUIRED; Type of event related to the originating occurrence.
	Type *string `json:"type,omitempty"`

	// Event data specific to the event type.
	Data any `json:"data,omitempty"`

	// Event data specific to the event type, encoded as a base64 string.
	DataBase64 []byte `json:"data_base64,omitempty"`

	// Content type of data value.
	DataContentType *string `json:"datacontenttype,omitempty"`

	// Identifies the schema that data adheres to.
	DataSchema *string `json:"dataschema,omitempty"`

	// This describes the subject of the event in the context of the event producer (identified by source).
	Subject *string `json:"subject,omitempty"`

	// The time (in UTC) the event was generated, in RFC3339 format.
	Time *time.Time `json:"time,omitempty"`
}

// FailedLockToken - Failed LockToken information.
type FailedLockToken struct {
	// REQUIRED; Error code related to the token. Example of such error codes are BadToken: which indicates the Token is not formatted
	// correctly, TokenLost: which indicates that token is not found, and
	// InternalServerError: For any internal server errors.
	ErrorCode *string `json:"errorCode,omitempty"`

	// REQUIRED; Description of the token error.
	ErrorDescription *string `json:"errorDescription,omitempty"`

	// REQUIRED; LockToken value
	LockToken *string `json:"lockToken,omitempty"`
}

// ReceiveDetails - Receive operation details per Cloud Event.
type ReceiveDetails struct {
	// REQUIRED; The Event Broker details.
	BrokerProperties *BrokerProperties `json:"brokerProperties,omitempty"`

	// REQUIRED; Cloud Event details.
	Event *CloudEvent `json:"event,omitempty"`
}

// ReceiveResult - Details of the Receive operation response.
type ReceiveResult struct {
	// REQUIRED; Array of receive responses, one per cloud event.
	Value []*ReceiveDetails `json:"value,omitempty"`
}

// RejectOptions - Array of lock token strings for the corresponding received Cloud Events to be rejected.
type RejectOptions struct {
	// REQUIRED; String array of lock tokens.
	LockTokens []*string `json:"lockTokens,omitempty"`
}

// RejectResult - The result of the Reject operation.
type RejectResult struct {
	// REQUIRED; Array of LockToken values for failed cloud events. Each LockToken includes the lock token value along with the
	// related error information (namely, the error code and description).
	FailedLockTokens []*FailedLockToken `json:"failedLockTokens,omitempty"`

	// REQUIRED; Array of lock tokens values for the successfully rejected cloud events.
	SucceededLockTokens []*string `json:"succeededLockTokens,omitempty"`
}

// ReleaseOptions - Array of lock token strings for the corresponding received Cloud Events to be released.
type ReleaseOptions struct {
	// REQUIRED; String array of lock tokens.
	LockTokens []*string `json:"lockTokens,omitempty"`
}

// ReleaseResult - The result of the Release operation.
type ReleaseResult struct {
	// REQUIRED; Array of LockToken values for failed cloud events. Each LockToken includes the lock token value along with the
	// related error information (namely, the error code and description).
	FailedLockTokens []*FailedLockToken `json:"failedLockTokens,omitempty"`

	// REQUIRED; Array of lock tokens values for the successfully released cloud events.
	SucceededLockTokens []*string `json:"succeededLockTokens,omitempty"`
}
