package redisenterpriseapi

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
	"github.com/Azure/azure-sdk-for-go/services/preview/redisenterprise/mgmt/2020-10-01-preview/redisenterprise"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result redisenterprise.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result redisenterprise.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*redisenterprise.OperationsClient)(nil)

// GetClientAPI contains the set of methods on the GetClient type.
type GetClientAPI interface {
	OperationStatusMethod(ctx context.Context, location string, operationID string) (result redisenterprise.OperationStatus, err error)
}

var _ GetClientAPI = (*redisenterprise.GetClient)(nil)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, clusterName string, parameters redisenterprise.Cluster) (result redisenterprise.CreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, clusterName string) (result redisenterprise.DeleteFuture, err error)
	GetMethod(ctx context.Context, resourceGroupName string, clusterName string) (result redisenterprise.Cluster, err error)
	List(ctx context.Context) (result redisenterprise.ClusterListPage, err error)
	ListComplete(ctx context.Context) (result redisenterprise.ClusterListIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result redisenterprise.ClusterListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result redisenterprise.ClusterListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, clusterName string, parameters redisenterprise.ClusterUpdate) (result redisenterprise.UpdateFuture, err error)
}

var _ ClientAPI = (*redisenterprise.Client)(nil)

// DatabasesClientAPI contains the set of methods on the DatabasesClient type.
type DatabasesClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters redisenterprise.Database) (result redisenterprise.DatabasesCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (result redisenterprise.DatabasesDeleteFuture, err error)
	GetMethod(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (result redisenterprise.Database, err error)
	ListByCluster(ctx context.Context, resourceGroupName string, clusterName string) (result redisenterprise.DatabaseListPage, err error)
	ListByClusterComplete(ctx context.Context, resourceGroupName string, clusterName string) (result redisenterprise.DatabaseListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters redisenterprise.DatabaseUpdate) (result redisenterprise.DatabasesUpdateFuture, err error)
}

var _ DatabasesClientAPI = (*redisenterprise.DatabasesClient)(nil)

// DatabaseClientAPI contains the set of methods on the DatabaseClient type.
type DatabaseClientAPI interface {
	Export(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters redisenterprise.ExportClusterParameters) (result redisenterprise.DatabaseExportFuture, err error)
	Import(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters redisenterprise.ImportClusterParameters) (result redisenterprise.DatabaseImportFuture, err error)
	ListKeys(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (result redisenterprise.AccessKeys, err error)
	RegenerateKey(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters redisenterprise.RegenerateKeyParameters) (result redisenterprise.DatabaseRegenerateKeyFuture, err error)
}

var _ DatabaseClientAPI = (*redisenterprise.DatabaseClient)(nil)

// PrivateEndpointConnectionsClientAPI contains the set of methods on the PrivateEndpointConnectionsClient type.
type PrivateEndpointConnectionsClientAPI interface {
	Delete(ctx context.Context, resourceGroupName string, clusterName string, privateEndpointConnectionName string) (result autorest.Response, err error)
	GetMethod(ctx context.Context, resourceGroupName string, clusterName string, privateEndpointConnectionName string) (result redisenterprise.PrivateEndpointConnection, err error)
	List(ctx context.Context, resourceGroupName string, clusterName string) (result redisenterprise.PrivateEndpointConnectionListResult, err error)
	Put(ctx context.Context, resourceGroupName string, clusterName string, privateEndpointConnectionName string, properties redisenterprise.PrivateEndpointConnection) (result redisenterprise.PrivateEndpointConnectionsPutFuture, err error)
}

var _ PrivateEndpointConnectionsClientAPI = (*redisenterprise.PrivateEndpointConnectionsClient)(nil)

// PrivateLinkResourcesClientAPI contains the set of methods on the PrivateLinkResourcesClient type.
type PrivateLinkResourcesClientAPI interface {
	ListByRedisEnterpriseCache(ctx context.Context, resourceGroupName string, clusterName string) (result redisenterprise.PrivateLinkResourceListResult, err error)
}

var _ PrivateLinkResourcesClientAPI = (*redisenterprise.PrivateLinkResourcesClient)(nil)
