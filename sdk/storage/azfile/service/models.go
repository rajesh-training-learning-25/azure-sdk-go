//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package service

import (
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azfile/internal/exported"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azfile/internal/generated"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azfile/share"
)

// SharedKeyCredential contains an account's name and its primary or secondary key.
type SharedKeyCredential = exported.SharedKeyCredential

// CreateShareOptions contains the optional parameters for the share.Client.Create method.
type CreateShareOptions = share.CreateOptions

// DeleteShareOptions contains the optional parameters for the share.Client.Delete method.
type DeleteShareOptions = share.DeleteOptions

// RestoreShareOptions contains the optional parameters for the share.Client.Restore method.
type RestoreShareOptions = share.RestoreOptions

// ---------------------------------------------------------------------------------------------------------------------

// GetPropertiesOptions provides set of options for Client.GetProperties
type GetPropertiesOptions struct {
	// placeholder for future options
}

func (o *GetPropertiesOptions) format() *generated.ServiceClientGetPropertiesOptions {
	return nil
}

// ---------------------------------------------------------------------------------------------------------------------

// SetPropertiesOptions provides set of options for Client.SetProperties
type SetPropertiesOptions struct {
	// The set of CORS rules.
	CORS []*ShareCORSRule

	// A summary of request statistics grouped by API in hourly aggregates for files.
	HourMetrics *ShareMetrics

	// A summary of request statistics grouped by API in minute aggregates for files.
	MinuteMetrics *ShareMetrics

	// Protocol settings
	Protocol *ShareProtocolSettings `xml:"ProtocolSettings"`
}

// ShareServiceProperties - Storage service properties.
type ShareServiceProperties = generated.ShareServiceProperties

// ShareCORSRule - CORS is an HTTP feature that enables a web application running under one domain to access resources in
// another domain. Web browsers implement a security restriction known as same-origin policy that
// prevents a web page from calling APIs in a different domain; CORS provides a secure way to allow one domain (the origin
// domain) to call APIs in another domain.
type ShareCORSRule = generated.ShareCORSRule

// ShareMetrics - Storage Analytics metrics for file service.
type ShareMetrics = generated.ShareMetrics

// ShareRetentionPolicy - The retention policy.
type ShareRetentionPolicy = generated.ShareRetentionPolicy

// ShareProtocolSettings - Protocol settings
type ShareProtocolSettings = generated.ShareProtocolSettings

// ShareSMBSettings - Settings for SMB protocol.
type ShareSMBSettings = generated.ShareSMBSettings

// SMBMultichannel - Settings for SMB multichannel
type SMBMultichannel = generated.SMBMultichannel

// ---------------------------------------------------------------------------------------------------------------------

// ListSharesOptions contains the optional parameters for the Client.NewListSharesPager method.
type ListSharesOptions struct {
	// Include this parameter to specify one or more datasets to include in the responseBody.
	Include []ListSharesIncludeType
	// A string value that identifies the portion of the list to be returned with the next list operation. The operation returns
	// a marker value within the responseBody body if the list returned was not complete.
	// The marker value may then be used in a subsequent call to request the next set of list items. The marker value is opaque
	// to the client.
	Marker *string
	// Specifies the maximum number of entries to return. If the request does not specify maxresults, or specifies a value greater
	// than 5,000, the server will return up to 5,000 items.
	MaxResults *int32
	// Filters the results to return only entries whose name begins with the specified prefix.
	Prefix *string
}

// Share - A listed Azure Storage share item.
type Share = generated.Share

// ShareProperties - Properties of a share.
type ShareProperties = generated.ShareProperties
