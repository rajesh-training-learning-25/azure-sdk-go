package accountapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/datalake/analytics/mgmt/2015-10-01-preview/account"
	"github.com/Azure/go-autorest/autorest"
)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	AddDataLakeStoreAccount(ctx context.Context, resourceGroupName string, accountName string, dataLakeStoreAccountName string, parameters account.AddDataLakeStoreParameters) (result autorest.Response, err error)
	AddStorageAccount(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, parameters account.AddStorageAccountParameters) (result autorest.Response, err error)
	Create(ctx context.Context, resourceGroupName string, name string, parameters account.DataLakeAnalyticsAccount) (result account.CreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string) (result account.DeleteFuture, err error)
	DeleteDataLakeStoreAccount(ctx context.Context, resourceGroupName string, accountName string, dataLakeStoreAccountName string) (result autorest.Response, err error)
	DeleteStorageAccount(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string) (result account.DataLakeAnalyticsAccount, err error)
	GetDataLakeStoreAccount(ctx context.Context, resourceGroupName string, accountName string, dataLakeStoreAccountName string) (result account.DataLakeStoreAccountInfo, err error)
	GetStorageAccount(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string) (result account.StorageAccountInfo, err error)
	GetStorageContainer(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, containerName string) (result account.BlobContainer, err error)
	List(ctx context.Context, filter string, top *int32, skip *int32, expand string, selectParameter string, orderby string, count *bool, search string, formatParameter string) (result account.DataLakeAnalyticsAccountListResultPage, err error)
	ListComplete(ctx context.Context, filter string, top *int32, skip *int32, expand string, selectParameter string, orderby string, count *bool, search string, formatParameter string) (result account.DataLakeAnalyticsAccountListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, top *int32, skip *int32, expand string, selectParameter string, orderby string, count *bool, search string, formatParameter string) (result account.DataLakeAnalyticsAccountListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string, top *int32, skip *int32, expand string, selectParameter string, orderby string, count *bool, search string, formatParameter string) (result account.DataLakeAnalyticsAccountListResultIterator, err error)
	ListDataLakeStoreAccounts(ctx context.Context, resourceGroupName string, accountName string, filter string, top *int32, skip *int32, expand string, selectParameter string, orderby string, count *bool, search string, formatParameter string) (result account.DataLakeAnalyticsAccountListDataLakeStoreResultPage, err error)
	ListDataLakeStoreAccountsComplete(ctx context.Context, resourceGroupName string, accountName string, filter string, top *int32, skip *int32, expand string, selectParameter string, orderby string, count *bool, search string, formatParameter string) (result account.DataLakeAnalyticsAccountListDataLakeStoreResultIterator, err error)
	ListSasTokens(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, containerName string) (result account.ListSasTokensResultPage, err error)
	ListSasTokensComplete(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, containerName string) (result account.ListSasTokensResultIterator, err error)
	ListStorageAccounts(ctx context.Context, resourceGroupName string, accountName string, filter string, top *int32, skip *int32, expand string, selectParameter string, orderby string, count *bool, search string, formatParameter string) (result account.DataLakeAnalyticsAccountListStorageAccountsResultPage, err error)
	ListStorageAccountsComplete(ctx context.Context, resourceGroupName string, accountName string, filter string, top *int32, skip *int32, expand string, selectParameter string, orderby string, count *bool, search string, formatParameter string) (result account.DataLakeAnalyticsAccountListStorageAccountsResultIterator, err error)
	ListStorageContainers(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string) (result account.ListBlobContainersResultPage, err error)
	ListStorageContainersComplete(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string) (result account.ListBlobContainersResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, name string, parameters account.DataLakeAnalyticsAccount) (result account.UpdateFuture, err error)
	UpdateStorageAccount(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, parameters account.AddStorageAccountParameters) (result autorest.Response, err error)
}

var _ ClientAPI = (*account.Client)(nil)
