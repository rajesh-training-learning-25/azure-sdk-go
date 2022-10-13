//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package appendblob

import (
	"context"
	"encoding/binary"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/hashing"
	"io"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/internal/base"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/internal/exported"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/internal/generated"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/internal/shared"
)

// ClientOptions contains the optional parameters when creating a Client.
type ClientOptions struct {
	azcore.ClientOptions
}

// Client represents a client to an Azure Storage append blob;
type Client base.CompositeClient[generated.BlobClient, generated.AppendBlobClient]

// NewClient creates an AppendBlobClient with the specified URL, Azure AD credential, and options.
func NewClient(blobURL string, cred azcore.TokenCredential, o *ClientOptions) (*Client, error) {
	authPolicy := runtime.NewBearerTokenPolicy(cred, []string{shared.TokenScope}, nil)
	conOptions := shared.GetClientOptions(o)
	conOptions.PerRetryPolicies = append(conOptions.PerRetryPolicies, authPolicy)
	pl := runtime.NewPipeline(exported.ModuleName, exported.ModuleVersion, runtime.PipelineOptions{}, &conOptions.ClientOptions)

	return (*Client)(base.NewAppendBlobClient(blobURL, pl, nil)), nil
}

// NewClientWithNoCredential creates an AppendBlobClient with the specified URL and options.
func NewClientWithNoCredential(blobURL string, o *ClientOptions) (*Client, error) {
	conOptions := shared.GetClientOptions(o)
	pl := runtime.NewPipeline(exported.ModuleName, exported.ModuleVersion, runtime.PipelineOptions{}, &conOptions.ClientOptions)

	return (*Client)(base.NewAppendBlobClient(blobURL, pl, nil)), nil
}

// NewClientWithSharedKeyCredential creates an AppendBlobClient with the specified URL, shared key, and options.
func NewClientWithSharedKeyCredential(blobURL string, cred *blob.SharedKeyCredential, o *ClientOptions) (*Client, error) {
	authPolicy := exported.NewSharedKeyCredPolicy(cred)
	conOptions := shared.GetClientOptions(o)
	conOptions.PerRetryPolicies = append(conOptions.PerRetryPolicies, authPolicy)
	pl := runtime.NewPipeline(exported.ModuleName, exported.ModuleVersion, runtime.PipelineOptions{}, &conOptions.ClientOptions)

	return (*Client)(base.NewAppendBlobClient(blobURL, pl, cred)), nil
}

// NewClientFromConnectionString creates Client from a connection String
func NewClientFromConnectionString(connectionString, containerName, blobName string, o *ClientOptions) (*Client, error) {
	parsed, err := shared.ParseConnectionString(connectionString)
	if err != nil {
		return nil, err
	}
	parsed.ServiceURL = runtime.JoinPaths(parsed.ServiceURL, containerName, blobName)

	if parsed.AccountKey != "" && parsed.AccountName != "" {
		credential, err := exported.NewSharedKeyCredential(parsed.AccountName, parsed.AccountKey)
		if err != nil {
			return nil, err
		}
		return NewClientWithSharedKeyCredential(parsed.ServiceURL, credential, o)
	}

	return NewClientWithNoCredential(parsed.ServiceURL, o)
}

// BlobClient returns the embedded blob client for this AppendBlob client.
func (ab *Client) BlobClient() *blob.Client {
	innerBlob, _ := base.InnerClients((*base.CompositeClient[generated.BlobClient, generated.AppendBlobClient])(ab))
	return (*blob.Client)(innerBlob)
}

func (ab *Client) sharedKey() *blob.SharedKeyCredential {
	return base.SharedKeyComposite((*base.CompositeClient[generated.BlobClient, generated.AppendBlobClient])(ab))
}

func (ab *Client) generated() *generated.AppendBlobClient {
	_, appendBlob := base.InnerClients((*base.CompositeClient[generated.BlobClient, generated.AppendBlobClient])(ab))
	return appendBlob
}

// URL returns the URL endpoint used by the Client object.
func (ab *Client) URL() string {
	return ab.generated().Endpoint()
}

// WithSnapshot creates a new AppendBlobURL object identical to the source but with the specified snapshot timestamp.
// Pass "" to remove the snapshot returning a URL to the base blob.
func (ab *Client) WithSnapshot(snapshot string) (*Client, error) {
	p, err := blob.ParseURL(ab.URL())
	if err != nil {
		return nil, err
	}
	p.Snapshot = snapshot

	return (*Client)(base.NewAppendBlobClient(p.String(), ab.generated().Pipeline(), ab.sharedKey())), nil
}

// WithVersionID creates a new AppendBlobURL object identical to the source but with the specified version id.
// Pass "" to remove the versionID returning a URL to the base blob.
func (ab *Client) WithVersionID(versionID string) (*Client, error) {
	p, err := blob.ParseURL(ab.URL())
	if err != nil {
		return nil, err
	}
	p.VersionID = versionID

	return (*Client)(base.NewAppendBlobClient(p.String(), ab.generated().Pipeline(), ab.sharedKey())), nil
}

// Create creates a 0-size append blob. Call AppendBlock to append data to an append blob.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/put-blob.
func (ab *Client) Create(ctx context.Context, o *CreateOptions) (CreateResponse, error) {
	opts, httpHeaders, leaseAccessConditions, cpkInfo, cpkScopeInfo, modifiedAccessConditions := o.format()
	resp, err := ab.generated().Create(ctx, 0, opts, httpHeaders, leaseAccessConditions, cpkInfo,
		cpkScopeInfo, modifiedAccessConditions)
	return resp, err
}

// AppendBlock writes a stream to a new block of data to the end of the existing append blob.
// This method panics if the stream is not at position 0.
// Note that the http client closes the body stream after the request is sent to the service.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/append-block.
func (ab *Client) AppendBlock(ctx context.Context, body io.ReadSeekCloser, o *AppendBlockOptions) (AppendBlockResponse, error) {
	count, err := shared.ValidateSeekableStreamAt0AndGetCount(body)
	if err != nil {
		return AppendBlockResponse{}, nil
	}

	appendOptions, appendPositionAccessConditions, cpkInfo, cpkScope, modifiedAccessConditions, leaseAccessConditions := o.format()

	if o.TransactionalContentCRC64 == 0 && o.TransactionalValidationOption != hashing.StorageTransferValidationOptionNone {
		body, err = hashing.NewHashingReadWrapper(body, o.TransactionalValidationOption)

		if o.TransactionalValidationOption&hashing.StorageTransferValidationOptionCRC64 == hashing.StorageTransferValidationOptionCRC64 {
			appendOptions.TransactionalContentCRC64 = make([]byte, 8)
			binary.LittleEndian.PutUint64(appendOptions.TransactionalContentCRC64, (body.(*hashing.ReadWrapper)).CRC64Hash())
		}
	}

	resp, err := ab.generated().AppendBlock(ctx, count, body, appendOptions, leaseAccessConditions, appendPositionAccessConditions, cpkInfo, cpkScope, modifiedAccessConditions)

	return resp, err
}

// AppendBlockFromURL copies a new block of data from source URL to the end of the existing append blob.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/append-block-from-url.
func (ab *Client) AppendBlockFromURL(ctx context.Context, source string, o *AppendBlockFromURLOptions) (AppendBlockFromURLResponse, error) {
	appendBlockFromURLOptions, cpkInfo, cpkScopeInfo, leaseAccessConditions, appendPositionAccessConditions, modifiedAccessConditions, sourceModifiedAccessConditions := o.format()

	// content length should be 0 on * from URL. always. It's a 400 if it isn't.
	resp, err := ab.generated().AppendBlockFromURL(ctx, source, 0, appendBlockFromURLOptions, cpkInfo, cpkScopeInfo,
		leaseAccessConditions, appendPositionAccessConditions, modifiedAccessConditions, sourceModifiedAccessConditions)
	return resp, err
}

// Seal - The purpose of Append Blob Seal is to allow users and applications to seal append blobs, marking them as read only.
// https://docs.microsoft.com/en-us/rest/api/storageservices/append-blob-seal
func (ab *Client) Seal(ctx context.Context, o *SealOptions) (SealResponse, error) {
	leaseAccessConditions, modifiedAccessConditions, positionAccessConditions := o.format()
	resp, err := ab.generated().Seal(ctx, nil, leaseAccessConditions, modifiedAccessConditions, positionAccessConditions)
	return resp, err
}

// Delete marks the specified blob or snapshot for deletion. The blob is later deleted during garbage collection.
// Note that deleting a blob also deletes all its snapshots.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/delete-blob.
func (ab *Client) Delete(ctx context.Context, o *blob.DeleteOptions) (blob.DeleteResponse, error) {
	return ab.BlobClient().Delete(ctx, o)
}

// Undelete restores the contents and metadata of a soft-deleted blob and any associated soft-deleted snapshots.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/undelete-blob.
func (ab *Client) Undelete(ctx context.Context, o *blob.UndeleteOptions) (blob.UndeleteResponse, error) {
	return ab.BlobClient().Undelete(ctx, o)
}

// SetTier operation sets the tier on a blob. The operation is allowed on a page
// blob in a premium storage account and on a block blob in a blob storage account (locally
// redundant storage only). A premium page blob's tier determines the allowed size, IOPS, and
// bandwidth of the blob. A block blob's tier determines Hot/Cool/Archive storage type. This operation
// does not update the blob's ETag.
// For detailed information about block blob level tiering see https://docs.microsoft.com/en-us/azure/storage/blobs/storage-blob-storage-tiers.
func (ab *Client) SetTier(ctx context.Context, tier blob.AccessTier, o *blob.SetTierOptions) (blob.SetTierResponse, error) {
	return ab.BlobClient().SetTier(ctx, tier, o)
}

// GetProperties returns the blob's properties.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/get-blob-properties.
func (ab *Client) GetProperties(ctx context.Context, o *blob.GetPropertiesOptions) (blob.GetPropertiesResponse, error) {
	return ab.BlobClient().GetProperties(ctx, o)
}

// SetHTTPHeaders changes a blob's HTTP headers.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/set-blob-properties.
func (ab *Client) SetHTTPHeaders(ctx context.Context, HTTPHeaders blob.HTTPHeaders, o *blob.SetHTTPHeadersOptions) (blob.SetHTTPHeadersResponse, error) {
	return ab.BlobClient().SetHTTPHeaders(ctx, HTTPHeaders, o)
}

// SetMetadata changes a blob's metadata.
// https://docs.microsoft.com/rest/api/storageservices/set-blob-metadata.
func (ab *Client) SetMetadata(ctx context.Context, metadata map[string]string, o *blob.SetMetadataOptions) (blob.SetMetadataResponse, error) {
	return ab.BlobClient().SetMetadata(ctx, metadata, o)
}

// CreateSnapshot creates a read-only snapshot of a blob.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/snapshot-blob.
func (ab *Client) CreateSnapshot(ctx context.Context, o *blob.CreateSnapshotOptions) (blob.CreateSnapshotResponse, error) {
	return ab.BlobClient().CreateSnapshot(ctx, o)
}

// StartCopyFromURL copies the data at the source URL to a blob.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/copy-blob.
func (ab *Client) StartCopyFromURL(ctx context.Context, copySource string, o *blob.StartCopyFromURLOptions) (blob.StartCopyFromURLResponse, error) {
	return ab.BlobClient().StartCopyFromURL(ctx, copySource, o)
}

// AbortCopyFromURL stops a pending copy that was previously started and leaves a destination blob with 0 length and metadata.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/abort-copy-blob.
func (ab *Client) AbortCopyFromURL(ctx context.Context, copyID string, o *blob.AbortCopyFromURLOptions) (blob.AbortCopyFromURLResponse, error) {
	return ab.BlobClient().AbortCopyFromURL(ctx, copyID, o)
}

// SetTags operation enables users to set tags on a blob or specific blob version, but not snapshot.
// Each call to this operation replaces all existing tags attached to the blob.
// To remove all tags from the blob, call this operation with no tags set.
// https://docs.microsoft.com/en-us/rest/api/storageservices/set-blob-tags
func (ab *Client) SetTags(ctx context.Context, tags map[string]string, o *blob.SetTagsOptions) (blob.SetTagsResponse, error) {
	return ab.BlobClient().SetTags(ctx, tags, o)
}

// GetTags operation enables users to get tags on a blob or specific blob version, or snapshot.
// https://docs.microsoft.com/en-us/rest/api/storageservices/get-blob-tags
func (ab *Client) GetTags(ctx context.Context, o *blob.GetTagsOptions) (blob.GetTagsResponse, error) {
	return ab.BlobClient().GetTags(ctx, o)
}

// CopyFromURL synchronously copies the data at the source URL to a block blob, with sizes up to 256 MB.
// For more information, see https://docs.microsoft.com/en-us/rest/api/storageservices/copy-blob-from-url.
func (ab *Client) CopyFromURL(ctx context.Context, copySource string, o *blob.CopyFromURLOptions) (blob.CopyFromURLResponse, error) {
	return ab.BlobClient().CopyFromURL(ctx, copySource, o)
}

// Concurrent Download Functions -----------------------------------------------------------------------------------------

// DownloadStream reads a range of bytes from a blob. The response also includes the blob's properties and metadata.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/get-blob.
func (ab *Client) DownloadStream(ctx context.Context, o *blob.DownloadStreamOptions) (blob.DownloadStreamResponse, error) {
	return ab.BlobClient().DownloadStream(ctx, o)
}

// DownloadBuffer downloads an Azure blob to a buffer with parallel.
func (ab *Client) DownloadBuffer(ctx context.Context, buffer []byte, o *blob.DownloadBufferOptions) (int64, error) {
	return ab.BlobClient().DownloadBuffer(ctx, shared.NewBytesWriter(buffer), o)
}

// DownloadFile downloads an Azure blob to a local file.
// The file would be truncated if the size doesn't match.
func (ab *Client) DownloadFile(ctx context.Context, file *os.File, o *blob.DownloadFileOptions) (int64, error) {
	return ab.BlobClient().DownloadFile(ctx, file, o)
}
