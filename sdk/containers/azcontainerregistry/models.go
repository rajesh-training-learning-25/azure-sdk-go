//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azcontainerregistry

import "time"

type acrAccessToken struct {
	// The access token for performing authenticated requests
	AccessToken *string
}

type acrRefreshToken struct {
	// The refresh token to be used for generating access tokens
	RefreshToken *string
}

// ArtifactManifestPlatform - The artifact's platform, consisting of operating system and architecture.
type ArtifactManifestPlatform struct {
	// READ-ONLY; Manifest digest
	Digest *string

	// READ-ONLY; CPU architecture
	Architecture *ArtifactArchitecture

	// READ-ONLY; Operating system
	OperatingSystem *ArtifactOperatingSystem
}

// ArtifactManifestProperties - Manifest attributes details
type ArtifactManifestProperties struct {
	// READ-ONLY; Manifest attributes
	Manifest *ManifestAttributes

	// READ-ONLY; Registry login server name. This is likely to be similar to {registry-name}.azurecr.io.
	RegistryLoginServer *string

	// READ-ONLY; Repository name
	RepositoryName *string
}

// ArtifactTagProperties - Tag attributes
type ArtifactTagProperties struct {
	// READ-ONLY; Registry login server name. This is likely to be similar to {registry-name}.azurecr.io.
	RegistryLoginServer *string

	// READ-ONLY; Image name
	RepositoryName *string

	// READ-ONLY; List of tag attribute details
	Tag *TagAttributes
}

// ContainerRepositoryProperties - Properties of this repository.
type ContainerRepositoryProperties struct {
	// REQUIRED; Writeable properties of the resource
	ChangeableAttributes *RepositoryWriteableProperties

	// READ-ONLY; Image created time
	CreatedOn *time.Time

	// READ-ONLY; Image last update time
	LastUpdatedOn *time.Time

	// READ-ONLY; Number of the manifests
	ManifestCount *int32

	// READ-ONLY; Image name
	Name *string

	// READ-ONLY; Registry login server name. This is likely to be similar to {registry-name}.azurecr.io.
	RegistryLoginServer *string

	// READ-ONLY; Number of the tags
	TagCount *int32
}

// ManifestAttributes - Manifest details
type ManifestAttributes struct {
	// READ-ONLY; Created time
	CreatedOn *time.Time

	// READ-ONLY; Manifest
	Digest *string

	// READ-ONLY; Last update time
	LastUpdatedOn *time.Time

	// Writeable properties of the resource
	ChangeableAttributes *ManifestWriteableProperties

	// READ-ONLY; CPU architecture
	Architecture *ArtifactArchitecture

	// READ-ONLY; Operating system
	OperatingSystem *ArtifactOperatingSystem

	// READ-ONLY; List of artifacts that are referenced by this manifest list, with information about the platform each supports.
	// This list will be empty if this is a leaf manifest and not a manifest list.
	RelatedArtifacts []*ArtifactManifestPlatform

	// READ-ONLY; Image size
	Size *int64

	// READ-ONLY; List of tags
	Tags []*string
}

// ManifestWriteableProperties - Changeable attributes
type ManifestWriteableProperties struct {
	// Delete enabled
	CanDelete *bool

	// List enabled
	CanList *bool

	// Read enabled
	CanRead *bool

	// Write enabled
	CanWrite *bool
}

// Manifests - Manifest attributes
type Manifests struct {
	// List of manifests
	Attributes []*ManifestAttributes
	Link       *string

	// Registry login server name. This is likely to be similar to {registry-name}.azurecr.io.
	RegistryLoginServer *string

	// Image name
	Repository *string
}

// Repositories - List of repositories
type Repositories struct {
	Link *string

	// Repository names
	Names []*string
}

// RepositoryWriteableProperties - Changeable attributes for Repository
type RepositoryWriteableProperties struct {
	// Delete enabled
	CanDelete *bool

	// List enabled
	CanList *bool

	// Read enabled
	CanRead *bool

	// Write enabled
	CanWrite *bool
}

// TagAttributes - Tag attribute details
type TagAttributes struct {
	// REQUIRED; Writeable properties of the resource
	ChangeableAttributes *TagWriteableProperties

	// READ-ONLY; Tag created time
	CreatedOn *time.Time

	// READ-ONLY; Tag digest
	Digest *string

	// READ-ONLY; Tag last update time
	LastUpdatedOn *time.Time

	// READ-ONLY; Tag name
	Name *string
}

// TagList - List of tag details
type TagList struct {
	// REQUIRED; Registry login server name. This is likely to be similar to {registry-name}.azurecr.io.
	RegistryLoginServer *string

	// REQUIRED; Image name
	Repository *string

	// REQUIRED; List of tag attribute details
	Tags []*TagAttributes
	Link *string
}

// TagWriteableProperties - Changeable attributes
type TagWriteableProperties struct {
	// Delete enabled
	CanDelete *bool

	// List enabled
	CanList *bool

	// Read enabled
	CanRead *bool

	// Write enabled
	CanWrite *bool
}
