//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package shared

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/uuid"
	"net"
	"strings"
)

const (
	TokenScope = "https://storage.azure.com/.default"
)

const (
	HeaderAuthorization     = "Authorization"
	HeaderXmsDate           = "x-ms-date"
	HeaderContentLength     = "Content-Length"
	HeaderContentEncoding   = "Content-Encoding"
	HeaderContentLanguage   = "Content-Language"
	HeaderContentType       = "Content-Type"
	HeaderContentMD5        = "Content-MD5"
	HeaderIfModifiedSince   = "If-Modified-Since"
	HeaderIfMatch           = "If-Match"
	HeaderIfNoneMatch       = "If-None-Match"
	HeaderIfUnmodifiedSince = "If-Unmodified-Since"
	HeaderRange             = "Range"
)

const StorageAnalyticsVersion = "1.0"

const (
	// DefaultFilePermissionString is a constant for all intents and purposes.
	// Inherit inherits permissions from the parent folder (default when creating files/folders)
	DefaultFilePermissionString = "inherit"

	// DefaultCurrentTimeString sets creation/last write times to now
	DefaultCurrentTimeString = "now"

	// DefaultPreserveString preserves old permissions on the file/folder (default when updating properties)
	DefaultPreserveString = "preserve"

	// FileAttributesNone is defaults for file attributes when creating file.
	// This attribute is valid only when used alone.
	FileAttributesNone = "None"

	// FileAttributesDirectory is defaults for file attributes when creating directory.
	// The attribute that identifies a directory
	FileAttributesDirectory = "Directory"
)

func GetClientOptions[T any](o *T) *T {
	if o == nil {
		return new(T)
	}
	return o
}

var errConnectionString = errors.New("connection string is either blank or malformed. The expected connection string " +
	"should contain key value pairs separated by semicolons. For example 'DefaultEndpointsProtocol=https;AccountName=<accountName>;" +
	"AccountKey=<accountKey>;EndpointSuffix=core.windows.net'")

type ParsedConnectionString struct {
	ServiceURL  string
	AccountName string
	AccountKey  string
}

func ParseConnectionString(connectionString string) (ParsedConnectionString, error) {
	const (
		defaultScheme = "https"
		defaultSuffix = "core.windows.net"
	)

	connStrMap := make(map[string]string)
	connectionString = strings.TrimRight(connectionString, ";")

	splitString := strings.Split(connectionString, ";")
	if len(splitString) == 0 {
		return ParsedConnectionString{}, errConnectionString
	}
	for _, stringPart := range splitString {
		parts := strings.SplitN(stringPart, "=", 2)
		if len(parts) != 2 {
			return ParsedConnectionString{}, errConnectionString
		}
		connStrMap[parts[0]] = parts[1]
	}

	accountName, ok := connStrMap["AccountName"]
	if !ok {
		return ParsedConnectionString{}, errors.New("connection string missing AccountName")
	}

	accountKey, ok := connStrMap["AccountKey"]
	if !ok {
		sharedAccessSignature, ok := connStrMap["SharedAccessSignature"]
		if !ok {
			return ParsedConnectionString{}, errors.New("connection string missing AccountKey and SharedAccessSignature")
		}
		return ParsedConnectionString{
			ServiceURL: fmt.Sprintf("%v://%v.file.%v/?%v", defaultScheme, accountName, defaultSuffix, sharedAccessSignature),
		}, nil
	}

	protocol, ok := connStrMap["DefaultEndpointsProtocol"]
	if !ok {
		protocol = defaultScheme
	}

	suffix, ok := connStrMap["EndpointSuffix"]
	if !ok {
		suffix = defaultSuffix
	}

	if fileEndpoint, ok := connStrMap["FileEndpoint"]; ok {
		return ParsedConnectionString{
			ServiceURL:  fileEndpoint,
			AccountName: accountName,
			AccountKey:  accountKey,
		}, nil
	}

	return ParsedConnectionString{
		ServiceURL:  fmt.Sprintf("%v://%v.file.%v", protocol, accountName, suffix),
		AccountName: accountName,
		AccountKey:  accountKey,
	}, nil
}

// IsIPEndpointStyle checks if URL's host is IP, in this case the storage account endpoint will be composed as:
// http(s)://IP(:port)/storageaccount/share(||container||etc)/...
// As url's Host property, host could be both host or host:port
func IsIPEndpointStyle(host string) bool {
	if host == "" {
		return false
	}
	if h, _, err := net.SplitHostPort(host); err == nil {
		host = h
	}
	// For IPv6, there could be case where SplitHostPort fails for cannot finding port.
	// In this case, eliminate the '[' and ']' in the URL.
	// For details about IPv6 URL, please refer to https://tools.ietf.org/html/rfc2732
	if host[0] == '[' && host[len(host)-1] == ']' {
		host = host[1 : len(host)-1]
	}
	return net.ParseIP(host) != nil
}

func GenerateLeaseID(leaseID *string) (*string, error) {
	if leaseID == nil {
		generatedUuid, err := uuid.New()
		if err != nil {
			return nil, err
		}
		leaseID = to.Ptr(generatedUuid.String())
	}
	return leaseID, nil
}
