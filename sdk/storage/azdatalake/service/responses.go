//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package service

import (
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/service"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azdatalake/filesystem"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azdatalake/internal/generated_blob"
)

// CreateFilesystemResponse contains the response fields for the CreateFilesystem operation.
type CreateFilesystemResponse = filesystem.CreateResponse

// DeleteFilesystemResponse contains the response fields for the DeleteFilesystem operation.
type DeleteFilesystemResponse = filesystem.DeleteResponse

// SetPropertiesResponse contains the response fields for the SetProperties operation.
type SetPropertiesResponse = service.SetPropertiesResponse

// GetPropertiesResponse contains the response fields for the GetProperties operation.
type GetPropertiesResponse = service.GetPropertiesResponse

// ListFilesystemsResponse contains the response fields for the ListFilesystems operation.
type ListFilesystemsResponse = generated_blob.ServiceClientListFileSystemsSegmentResponse
