//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package shared

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/internal/generated"
	"io"
	"net/url"
	"strconv"
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

const CountToEnd = 0

// HTTPRange defines a range of bytes within an HTTP resource, starting at offset and
// ending at offset+count. A zero-value HttpRange indicates the entire resource. An HttpRange
// which has an offset but na zero value count indicates from the offset to the resource's end.
type HTTPRange struct {
	Offset int64
	Count  int64
}

func (r HTTPRange) Format() *string {
	if r.Offset == 0 && r.Count == 0 { // Do common case first for performance
		return nil // No specified range
	}
	endOffset := "" // if count == CountToEnd (0)
	if r.Count > 0 {
		endOffset = strconv.FormatInt((r.Offset+r.Count)-1, 10)
	}
	dataRange := fmt.Sprintf("bytes=%v-%s", r.Offset, endOffset)
	return &dataRange
}

func GetSourceRange(offset, count *int64) *string {
	if offset == nil && count == nil {
		return nil
	}
	newOffset := int64(0)
	newCount := int64(CountToEnd)

	if offset != nil {
		newOffset = *offset
	}

	if count != nil {
		newCount = *count
	}

	return (&HTTPRange{Offset: newOffset, Count: newCount}).Format()
}

// CopyOptions returns a zero-value T if opts is nil.
// If opts is not nil, a copy is made and its address returned.
func CopyOptions[T any](opts *T) *T {
	if opts == nil {
		return new(T)
	}
	cp := *opts
	return &cp
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
			ServiceURL: fmt.Sprintf("%v://%v.blob.%v/?%v", defaultScheme, accountName, defaultSuffix, sharedAccessSignature),
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

	if blobEndpoint, ok := connStrMap["BlobEndpoint"]; ok {
		return ParsedConnectionString{
			ServiceURL:  blobEndpoint,
			AccountName: accountName,
			AccountKey:  accountKey,
		}, nil
	}

	return ParsedConnectionString{
		ServiceURL:  fmt.Sprintf("%v://%v.blob.%v", protocol, accountName, suffix),
		AccountName: accountName,
		AccountKey:  accountKey,
	}, nil
}

func SerializeBlobTags(tagsMap map[string]string) *generated.BlobTags {
	if tagsMap == nil {
		return nil
	}
	blobTagSet := make([]*generated.BlobTag, 0)
	for key, val := range tagsMap {
		newKey, newVal := key, val
		blobTagSet = append(blobTagSet, &generated.BlobTag{Key: &newKey, Value: &newVal})
	}
	return &generated.BlobTags{BlobTagSet: blobTagSet}
}
func SerializeBlobTagsToStrPtr(tagsMap map[string]string) *string {
	if tagsMap == nil {
		return nil
	}
	tags := make([]string, 0)
	for key, val := range tagsMap {
		tags = append(tags, url.QueryEscape(key)+"="+url.QueryEscape(val))
	}
	blobTagsString := strings.Join(tags, "&")
	return &blobTagsString
}

func ValidateSeekableStreamAt0AndGetCount(body io.ReadSeeker) (int64, error) {
	if body == nil { // nil body's are "logically" seekable to 0 and are 0 bytes long
		return 0, nil
	}

	err := validateSeekableStreamAt0(body)
	if err != nil {
		return 0, err
	}

	count, err := body.Seek(0, io.SeekEnd)
	if err != nil {
		return 0, errors.New("body stream must be seekable")
	}

	_, err = body.Seek(0, io.SeekStart)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// return an error if body is not a valid seekable stream at 0
func validateSeekableStreamAt0(body io.ReadSeeker) error {
	if body == nil { // nil body's are "logically" seekable to 0
		return nil
	}
	if pos, err := body.Seek(0, io.SeekCurrent); pos != 0 || err != nil {
		// Help detect programmer error
		if err != nil {
			return errors.New("body stream must be seekable")
		}
		return errors.New("body stream must be set to position 0")
	}
	return nil
}

func RangeToString(offset, count int64) string {
	return "bytes=" + strconv.FormatInt(offset, 10) + "-" + strconv.FormatInt(offset+count-1, 10)
}
