//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// Contains common helpers for TESTS ONLY
package testcommon

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/recording"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blockblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/service"
	"github.com/stretchr/testify/require"
)

type TestAccountType string

const (
	TestAccountDefault    TestAccountType = ""
	TestAccountSecondary  TestAccountType = "SECONDARY_"
	TestAccountPremium    TestAccountType = "PREMIUM_"
	TestAccountSoftDelete TestAccountType = "SOFT_DELETE_"
	TestAccountDatalake   TestAccountType = "DATALAKE_"
	TestAccountImmutable  TestAccountType = "IMMUTABLE_"
)

const (
	DefaultEndpointSuffix       = "core.windows.net/"
	DefaultBlobEndpointSuffix   = "blob.core.windows.net/"
	AccountNameEnvVar           = "AZURE_STORAGE_ACCOUNT_NAME"
	AccountKeyEnvVar            = "AZURE_STORAGE_ACCOUNT_KEY"
	DefaultEndpointSuffixEnvVar = "AZURE_STORAGE_ENDPOINT_SUFFIX"
)

const (
	FakeStorageAccount = "fakestorage"
	FakeStorageURL     = "https://fakestorage.blob.core.windows.net"
)

var (
	BlobContentType        = "my_type"
	BlobContentDisposition = "my_disposition"
	BlobCacheControl       = "control"
	BlobContentLanguage    = "my_language"
	BlobContentEncoding    = "my_encoding"
)

var BasicHeaders = blob.HTTPHeaders{
	BlobContentType:        &BlobContentType,
	BlobContentDisposition: &BlobContentDisposition,
	BlobCacheControl:       &BlobCacheControl,
	BlobContentMD5:         nil,
	BlobContentLanguage:    &BlobContentLanguage,
	BlobContentEncoding:    &BlobContentEncoding,
}

var BasicMetadata = map[string]string{"Foo": "bar"}

var BasicBlobTagsMap = map[string]string{
	"azure": "blob",
	"blob":  "sdk",
	"sdk":   "go",
}

var SpecialCharBlobTagsMap = map[string]string{
	"+-./:=_ ":        "firsttag",
	"tag2":            "+-./:=_",
	"+-./:=_1":        "+-./:=_",
	"Microsoft Azure": "Azure Storage",
	"Storage+SDK":     "SDK/GO",
	"GO ":             ".Net",
}

func setClientOptions(t *testing.T, opts *azcore.ClientOptions) {
	opts.Logging.AllowedHeaders = append(opts.Logging.AllowedHeaders, "X-Request-Mismatch", "X-Request-Mismatch-Error")

	transport, err := recording.NewRecordingHTTPClient(t, nil)
	require.NoError(t, err)
	opts.Transport = transport
}

func GetClient(t *testing.T, accountType TestAccountType, options *azblob.ClientOptions) (*azblob.Client, error) {
	if options == nil {
		options = &azblob.ClientOptions{}
	}

	setClientOptions(t, &options.ClientOptions)

	cred, err := GetGenericSharedKeyCredential(accountType)
	if err != nil {
		return nil, err
	}

	client, err := azblob.NewClientWithSharedKeyCredential("https://"+cred.AccountName()+".blob.core.windows.net/", cred, options)

	return client, err
}

func GetServiceClient(t *testing.T, accountType TestAccountType, options *service.ClientOptions) (*service.Client, error) {
	if options == nil {
		options = &service.ClientOptions{}
	}

	setClientOptions(t, &options.ClientOptions)

	cred, err := GetGenericSharedKeyCredential(accountType)
	if err != nil {
		return nil, err
	}

	serviceClient, err := service.NewClientWithSharedKeyCredential("https://"+cred.AccountName()+".blob.core.windows.net/", cred, options)

	return serviceClient, err
}

func GetGenericAccountInfo(accountType TestAccountType) (string, string) {
	if recording.GetRecordMode() == recording.PlaybackMode {
		return FakeStorageAccount, "ZmFrZQ=="
	}
	accountNameEnvVar := string(accountType) + AccountNameEnvVar
	accountKeyEnvVar := string(accountType) + AccountKeyEnvVar
	accountName, _ := GetRequiredEnv(accountNameEnvVar)
	accountKey, _ := GetRequiredEnv(accountKeyEnvVar)
	return accountName, accountKey
}

func GetGenericSharedKeyCredential(accountType TestAccountType) (*azblob.SharedKeyCredential, error) {
	accountName, accountKey := GetGenericAccountInfo(accountType)
	if accountName == "" || accountKey == "" {
		return nil, errors.New(string(accountType) + AccountNameEnvVar + " and/or " + string(accountType) + AccountKeyEnvVar + " environment variables not specified.")
	}
	return azblob.NewSharedKeyCredential(accountName, accountKey)
}

func GetGenericConnectionString(accountType TestAccountType) (*string, error) {
	accountName, accountKey := GetGenericAccountInfo(accountType)
	if accountName == "" || accountKey == "" {
		return nil, errors.New(string(accountType) + AccountNameEnvVar + " and/or " + string(accountType) + AccountKeyEnvVar + " environment variables not specified.")
	}
	connectionString := fmt.Sprintf("DefaultEndpointsProtocol=https;AccountName=%s;AccountKey=%s;EndpointSuffix=core.windows.net/",
		accountName, accountKey)
	return &connectionString, nil
}

func GetServiceClientFromConnectionString(t *testing.T, accountType TestAccountType, options *service.ClientOptions) (*service.Client, error) {
	if options == nil {
		options = &service.ClientOptions{}
	}

	setClientOptions(t, &options.ClientOptions)

	cred, err := GetGenericConnectionString(accountType)
	if err != nil {
		return nil, err
	}

	serviceClient, err := service.NewClientFromConnectionString(*cred, options)

	return serviceClient, err
}

func GetContainerClient(containerName string, s *service.Client) *container.Client {
	return s.NewContainerClient(containerName)
}

func CreateNewContainer(ctx context.Context, _require *require.Assertions, containerName string, serviceClient *service.Client) *container.Client {
	containerClient := GetContainerClient(containerName, serviceClient)

	_, err := containerClient.Create(ctx, nil)
	_require.Nil(err)
	// _require.Equal(cResp.RawResponse.StatusCode, 201)
	return containerClient
}

func DeleteContainer(ctx context.Context, _require *require.Assertions, containerClient *container.Client) {
	_, err := containerClient.Delete(ctx, nil)
	_require.Nil(err)
}

func GetBlobClient(blockBlobName string, containerClient *container.Client) *blob.Client {
	return containerClient.NewBlobClient(blockBlobName)
}

func CreateNewBlobs(ctx context.Context, _require *require.Assertions, blobNames []string, containerClient *container.Client) {
	for _, blobName := range blobNames {
		CreateNewBlockBlob(ctx, _require, blobName, containerClient)
	}
}

func GetBlockBlobClient(blockBlobName string, containerClient *container.Client) *blockblob.Client {
	return containerClient.NewBlockBlobClient(blockBlobName)
}

func CreateNewBlockBlob(ctx context.Context, _require *require.Assertions, blockBlobName string, containerClient *container.Client) *blockblob.Client {
	bbClient := GetBlockBlobClient(blockBlobName, containerClient)

	_, err := bbClient.Upload(ctx, streaming.NopCloser(strings.NewReader(BlockBlobDefaultData)), nil)
	_require.Nil(err)
	// _require.Equal(cResp.RawResponse.StatusCode, 201)
	return bbClient
}

func CreateNewBlockBlobWithCPK(ctx context.Context, _require *require.Assertions, blockBlobName string, containerClient *container.Client, cpkInfo *blob.CpkInfo, cpkScopeInfo *blob.CpkScopeInfo) (bbClient *blockblob.Client) {
	bbClient = GetBlockBlobClient(blockBlobName, containerClient)

	uploadBlockBlobOptions := blockblob.UploadOptions{
		CpkInfo:      cpkInfo,
		CpkScopeInfo: cpkScopeInfo,
	}
	cResp, err := bbClient.Upload(ctx, streaming.NopCloser(strings.NewReader(BlockBlobDefaultData)), &uploadBlockBlobOptions)
	_require.Nil(err)
	// _require.Equal(cResp.RawResponse.StatusCode, 201)
	_require.Equal(*cResp.IsServerEncrypted, true)
	if cpkInfo != nil {
		_require.EqualValues(cResp.EncryptionKeySHA256, cpkInfo.EncryptionKeySHA256)
	}
	if cpkScopeInfo != nil {
		_require.EqualValues(cResp.EncryptionScope, cpkScopeInfo.EncryptionScope)
	}
	return
}

// Some tests require setting service properties. It can take up to 30 seconds for the new properties to be reflected across all FEs.
// We will enable the necessary property and try to run the test implementation. If it fails with an error that should be due to
// those changes not being reflected yet, we will wait 30 seconds and try the test again. If it fails this time for any reason,
// we fail the test. It is the responsibility of the testImplFunc to determine which error string indicates the test should be retried.
// There can only be one such string. All errors that cannot be due to this detail should be asserted and not returned as an error string.
func RunTestRequiringServiceProperties(ctx context.Context, _require *require.Assertions, svcClient *service.Client, code string,
	enableServicePropertyFunc func(context.Context, *require.Assertions, *service.Client),
	testImplFunc func(context.Context, *require.Assertions, *service.Client) error,
	disableServicePropertyFunc func(context.Context, *require.Assertions, *service.Client)) {

	enableServicePropertyFunc(ctx, _require, svcClient)
	defer disableServicePropertyFunc(ctx, _require, svcClient)

	err := testImplFunc(ctx, _require, svcClient)
	// We cannot assume that the error indicative of slow update will necessarily be a StorageError. As in ListBlobs.
	if err != nil && err.Error() == code {
		time.Sleep(time.Second * 30)
		err = testImplFunc(ctx, _require, svcClient)
		_require.Nil(err)
	}
}

func EnableSoftDelete(ctx context.Context, _require *require.Assertions, client *service.Client) {
	days := int32(1)
	_, err := client.SetProperties(ctx, &service.SetPropertiesOptions{
		DeleteRetentionPolicy: &service.RetentionPolicy{Enabled: to.Ptr(true), Days: &days}})
	_require.Nil(err)
}

func DisableSoftDelete(ctx context.Context, _require *require.Assertions, client *service.Client) {
	_, err := client.SetProperties(ctx, &service.SetPropertiesOptions{DeleteRetentionPolicy: &service.RetentionPolicy{Enabled: to.Ptr(false)}})
	_require.Nil(err)
}

func ListBlobsCount(ctx context.Context, _require *require.Assertions, listPager *runtime.Pager[container.ListBlobsFlatResponse], ctr int) {
	found := make([]*container.BlobItem, 0)
	for listPager.More() {
		resp, err := listPager.NextPage(ctx)
		_require.Nil(err)
		if err != nil {
			break
		}
		found = append(found, resp.Segment.BlobItems...)
	}
	_require.Len(found, ctr)
}
