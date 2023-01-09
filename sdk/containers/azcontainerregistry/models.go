//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azcontainerregistry

import (
	"io"
	"time"
)

type acrAccessToken struct {
	// The access token for performing authenticated requests
	AccessToken *string `json:"access_token,omitempty"`
}

// ACRManifests - Manifest attributes
type ACRManifests struct {
	Link *string `json:"link,omitempty"`

	// List of manifests
	Manifests []*ManifestAttributes `json:"manifests,omitempty"`

	// Registry login server name. This is likely to be similar to {registry-name}.azurecr.io.
	RegistryLoginServer *string `json:"registry,omitempty"`

	// Image name
	Repository *string `json:"imageName,omitempty"`
}

type acrRefreshToken struct {
	// The refresh token to be used for generating access tokens
	RefreshToken *string `json:"refresh_token,omitempty"`
}

// ArtifactManifestPlatform - The artifact's platform, consisting of operating system and architecture.
type ArtifactManifestPlatform struct {
	// READ-ONLY; Manifest digest
	Digest *string `json:"digest,omitempty" azure:"ro"`

	// READ-ONLY; CPU architecture
	Architecture *ArtifactArchitecture `json:"architecture,omitempty" azure:"ro"`

	// READ-ONLY; Operating system
	OperatingSystem *ArtifactOperatingSystem `json:"os,omitempty" azure:"ro"`
}

// ArtifactManifestProperties - Manifest attributes details
type ArtifactManifestProperties struct {
	// READ-ONLY; Manifest attributes
	Manifest *ManifestAttributes `json:"manifest,omitempty" azure:"ro"`

	// READ-ONLY; Registry login server name. This is likely to be similar to {registry-name}.azurecr.io.
	RegistryLoginServer *string `json:"registry,omitempty" azure:"ro"`

	// READ-ONLY; Repository name
	RepositoryName *string `json:"imageName,omitempty" azure:"ro"`
}

// ArtifactTagProperties - Tag attributes
type ArtifactTagProperties struct {
	// READ-ONLY; Registry login server name. This is likely to be similar to {registry-name}.azurecr.io.
	RegistryLoginServer *string `json:"registry,omitempty" azure:"ro"`

	// READ-ONLY; Image name
	RepositoryName *string `json:"imageName,omitempty" azure:"ro"`

	// READ-ONLY; List of tag attribute details
	Tag *TagAttributes `json:"tag,omitempty" azure:"ro"`
}

// authenticationClientExchangeAADAccessTokenForACRRefreshTokenOptions contains the optional parameters for the authenticationClient.ExchangeAADAccessTokenForACRRefreshToken
// method.
type authenticationClientExchangeAADAccessTokenForACRRefreshTokenOptions struct {
	// AAD access token, mandatory when granttype is accesstokenrefreshtoken or access_token.
	AccessToken *string
	// AAD refresh token, mandatory when granttype is accesstokenrefreshtoken or refresh_token
	RefreshToken *string
	// AAD tenant associated to the AAD credentials.
	Tenant *string
}

// authenticationClientExchangeACRRefreshTokenForACRAccessTokenOptions contains the optional parameters for the authenticationClient.ExchangeACRRefreshTokenForACRAccessToken
// method.
type authenticationClientExchangeACRRefreshTokenForACRAccessTokenOptions struct {
	// Grant type is expected to be refresh_token
	GrantType *tokenGrantType
}

// BlobClientCancelUploadOptions contains the optional parameters for the BlobClient.CancelUpload method.
type BlobClientCancelUploadOptions struct {
	// placeholder for future optional parameters
}

// BlobClientCheckBlobExistsOptions contains the optional parameters for the BlobClient.CheckBlobExists method.
type BlobClientCheckBlobExistsOptions struct {
	// placeholder for future optional parameters
}

// BlobClientCheckChunkExistsOptions contains the optional parameters for the BlobClient.CheckChunkExists method.
type BlobClientCheckChunkExistsOptions struct {
	// placeholder for future optional parameters
}

// BlobClientCompleteUploadOptions contains the optional parameters for the BlobClient.CompleteUpload method.
type BlobClientCompleteUploadOptions struct {
	// Optional raw data of blob
	BlobData io.ReadSeekCloser
}

// BlobClientDeleteBlobOptions contains the optional parameters for the BlobClient.DeleteBlob method.
type BlobClientDeleteBlobOptions struct {
	// placeholder for future optional parameters
}

// BlobClientGetBlobOptions contains the optional parameters for the BlobClient.GetBlob method.
type BlobClientGetBlobOptions struct {
	// placeholder for future optional parameters
}

// BlobClientGetChunkOptions contains the optional parameters for the BlobClient.GetChunk method.
type BlobClientGetChunkOptions struct {
	// placeholder for future optional parameters
}

// BlobClientGetUploadStatusOptions contains the optional parameters for the BlobClient.GetUploadStatus method.
type BlobClientGetUploadStatusOptions struct {
	// placeholder for future optional parameters
}

// BlobClientMountBlobOptions contains the optional parameters for the BlobClient.MountBlob method.
type BlobClientMountBlobOptions struct {
	// placeholder for future optional parameters
}

// BlobClientStartUploadOptions contains the optional parameters for the BlobClient.StartUpload method.
type BlobClientStartUploadOptions struct {
	// placeholder for future optional parameters
}

// BlobClientUploadChunkOptions contains the optional parameters for the BlobClient.UploadChunk method.
type BlobClientUploadChunkOptions struct {
	// placeholder for future optional parameters
}

// ClientDeleteManifestOptions contains the optional parameters for the Client.DeleteManifest method.
type ClientDeleteManifestOptions struct {
	// placeholder for future optional parameters
}

// ClientDeleteRepositoryOptions contains the optional parameters for the Client.DeleteRepository method.
type ClientDeleteRepositoryOptions struct {
	// placeholder for future optional parameters
}

// ClientDeleteTagOptions contains the optional parameters for the Client.DeleteTag method.
type ClientDeleteTagOptions struct {
	// placeholder for future optional parameters
}

// ClientGetManifestOptions contains the optional parameters for the Client.GetManifest method.
type ClientGetManifestOptions struct {
	// Accept header string delimited by comma. For example, application/vnd.docker.distribution.manifest.v2+json
	Accept *string
}

// ClientGetManifestPropertiesOptions contains the optional parameters for the Client.GetManifestProperties method.
type ClientGetManifestPropertiesOptions struct {
	// placeholder for future optional parameters
}

// ClientGetRepositoryPropertiesOptions contains the optional parameters for the Client.GetRepositoryProperties method.
type ClientGetRepositoryPropertiesOptions struct {
	// placeholder for future optional parameters
}

// ClientGetTagPropertiesOptions contains the optional parameters for the Client.GetTagProperties method.
type ClientGetTagPropertiesOptions struct {
	// placeholder for future optional parameters
}

// ClientListManifestsOptions contains the optional parameters for the Client.NewListManifestsPager method.
type ClientListManifestsOptions struct {
	// Query parameter for the last item in previous query. Result set will include values lexically after last.
	Last *string
	// query parameter for max number of items
	N *int32
	// Sort options for ordering manifests in a collection.
	OrderBy *ArtifactManifestOrderBy
}

// ClientListRepositoriesOptions contains the optional parameters for the Client.NewListRepositoriesPager method.
type ClientListRepositoriesOptions struct {
	// Query parameter for the last item in previous query. Result set will include values lexically after last.
	Last *string
	// query parameter for max number of items
	N *int32
}

// ClientListTagsOptions contains the optional parameters for the Client.NewListTagsPager method.
type ClientListTagsOptions struct {
	// filter by digest
	Digest *string
	// Query parameter for the last item in previous query. Result set will include values lexically after last.
	Last *string
	// query parameter for max number of items
	N *int32
	// Sort options for ordering tags in a collection.
	OrderBy *ArtifactTagOrderBy
}

// ClientUpdateManifestPropertiesOptions contains the optional parameters for the Client.UpdateManifestProperties method.
type ClientUpdateManifestPropertiesOptions struct {
	// Manifest attribute value
	Value *ManifestWriteableProperties
}

// ClientUpdateRepositoryPropertiesOptions contains the optional parameters for the Client.UpdateRepositoryProperties method.
type ClientUpdateRepositoryPropertiesOptions struct {
	// Repository attribute value
	Value *RepositoryWriteableProperties
}

// ClientUpdateTagPropertiesOptions contains the optional parameters for the Client.UpdateTagProperties method.
type ClientUpdateTagPropertiesOptions struct {
	// Tag attribute value
	Value *TagWriteableProperties
}

// ClientUploadManifestOptions contains the optional parameters for the Client.UploadManifest method.
type ClientUploadManifestOptions struct {
	// placeholder for future optional parameters
}

// ContainerRepositoryProperties - Properties of this repository.
type ContainerRepositoryProperties struct {
	// REQUIRED; Writeable properties of the resource
	ChangeableAttributes *RepositoryWriteableProperties `json:"changeableAttributes,omitempty"`

	// READ-ONLY; Image created time
	CreatedOn *time.Time `json:"createdTime,omitempty" azure:"ro"`

	// READ-ONLY; Image last update time
	LastUpdatedOn *time.Time `json:"lastUpdateTime,omitempty" azure:"ro"`

	// READ-ONLY; Number of the manifests
	ManifestCount *int32 `json:"manifestCount,omitempty" azure:"ro"`

	// READ-ONLY; Image name
	Name *string `json:"imageName,omitempty" azure:"ro"`

	// READ-ONLY; Registry login server name. This is likely to be similar to {registry-name}.azurecr.io.
	RegistryLoginServer *string `json:"registry,omitempty" azure:"ro"`

	// READ-ONLY; Number of the tags
	TagCount *int32 `json:"tagCount,omitempty" azure:"ro"`
}

// ManifestAttributes - Manifest details
type ManifestAttributes struct {
	// READ-ONLY; Created time
	CreatedOn *time.Time `json:"createdTime,omitempty" azure:"ro"`

	// READ-ONLY; Manifest
	Digest *string `json:"digest,omitempty" azure:"ro"`

	// READ-ONLY; Last update time
	LastUpdatedOn *time.Time `json:"lastUpdateTime,omitempty" azure:"ro"`

	// Writeable properties of the resource
	ChangeableAttributes *ManifestWriteableProperties `json:"changeableAttributes,omitempty"`

	// READ-ONLY; CPU architecture
	Architecture *ArtifactArchitecture `json:"architecture,omitempty" azure:"ro"`

	// READ-ONLY; Operating system
	OperatingSystem *ArtifactOperatingSystem `json:"os,omitempty" azure:"ro"`

	// READ-ONLY; List of artifacts that are referenced by this manifest list, with information about the platform each supports.
	// This list will be empty if this is a leaf manifest and not a manifest list.
	RelatedArtifacts []*ArtifactManifestPlatform `json:"references,omitempty" azure:"ro"`

	// READ-ONLY; Image size
	Size *int64 `json:"imageSize,omitempty" azure:"ro"`

	// READ-ONLY; List of tags
	Tags []*string `json:"tags,omitempty" azure:"ro"`
}

// ManifestWriteableProperties - Changeable attributes
type ManifestWriteableProperties struct {
	// Delete enabled
	CanDelete *bool `json:"deleteEnabled,omitempty"`

	// List enabled
	CanList *bool `json:"listEnabled,omitempty"`

	// Read enabled
	CanRead *bool `json:"readEnabled,omitempty"`

	// Write enabled
	CanWrite *bool `json:"writeEnabled,omitempty"`
}

// Repositories - List of repositories
type Repositories struct {
	Link *string `json:"link,omitempty"`

	// Repository names
	Names []*string `json:"repositories,omitempty"`
}

// RepositoryWriteableProperties - Changeable attributes for Repository
type RepositoryWriteableProperties struct {
	// Delete enabled
	CanDelete *bool `json:"deleteEnabled,omitempty"`

	// List enabled
	CanList *bool `json:"listEnabled,omitempty"`

	// Read enabled
	CanRead *bool `json:"readEnabled,omitempty"`

	// Write enabled
	CanWrite *bool `json:"writeEnabled,omitempty"`
}

// TagAttributes - Tag attribute details
type TagAttributes struct {
	// REQUIRED; Writeable properties of the resource
	ChangeableAttributes *TagWriteableProperties `json:"changeableAttributes,omitempty"`

	// READ-ONLY; Tag created time
	CreatedOn *time.Time `json:"createdTime,omitempty" azure:"ro"`

	// READ-ONLY; Tag digest
	Digest *string `json:"digest,omitempty" azure:"ro"`

	// READ-ONLY; Tag last update time
	LastUpdatedOn *time.Time `json:"lastUpdateTime,omitempty" azure:"ro"`

	// READ-ONLY; Tag name
	Name *string `json:"name,omitempty" azure:"ro"`
}

// TagList - List of tag details
type TagList struct {
	// REQUIRED; Registry login server name. This is likely to be similar to {registry-name}.azurecr.io.
	RegistryLoginServer *string `json:"registry,omitempty"`

	// REQUIRED; Image name
	Repository *string `json:"imageName,omitempty"`

	// REQUIRED; List of tag attribute details
	Tags []*TagAttributes `json:"tags,omitempty"`
	Link *string          `json:"link,omitempty"`
}

// TagWriteableProperties - Changeable attributes
type TagWriteableProperties struct {
	// Delete enabled
	CanDelete *bool `json:"deleteEnabled,omitempty"`

	// List enabled
	CanList *bool `json:"listEnabled,omitempty"`

	// Read enabled
	CanRead *bool `json:"readEnabled,omitempty"`

	// Write enabled
	CanWrite *bool `json:"writeEnabled,omitempty"`
}
