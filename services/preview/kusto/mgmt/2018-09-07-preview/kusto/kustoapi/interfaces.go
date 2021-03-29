package kustoapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/kusto/mgmt/2018-09-07-preview/kusto"
)

// ClustersClientAPI contains the set of methods on the ClustersClient type.
type ClustersClientAPI interface {
	CheckNameAvailability(ctx context.Context, location string, clusterName kusto.ClusterCheckNameRequest) (result kusto.CheckNameResult, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, parameters kusto.Cluster) (result kusto.ClustersCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, clusterName string) (result kusto.ClustersDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, clusterName string) (result kusto.Cluster, err error)
	List(ctx context.Context) (result kusto.ClusterListResult, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result kusto.ClusterListResult, err error)
	ListSkus(ctx context.Context) (result kusto.ListSkusResult, err error)
	ListSkusByResource(ctx context.Context, resourceGroupName string, clusterName string) (result kusto.ListResourceSkusResult, err error)
	Start(ctx context.Context, resourceGroupName string, clusterName string) (result kusto.ClustersStartFuture, err error)
	Stop(ctx context.Context, resourceGroupName string, clusterName string) (result kusto.ClustersStopFuture, err error)
	Update(ctx context.Context, resourceGroupName string, clusterName string, parameters kusto.ClusterUpdate) (result kusto.ClustersUpdateFuture, err error)
}

var _ ClustersClientAPI = (*kusto.ClustersClient)(nil)

// DatabasesClientAPI contains the set of methods on the DatabasesClient type.
type DatabasesClientAPI interface {
	AddPrincipals(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, databasePrincipalsToAdd kusto.DatabasePrincipalListRequest) (result kusto.DatabasePrincipalListResult, err error)
	CheckNameAvailability(ctx context.Context, resourceGroupName string, clusterName string, databaseName kusto.DatabaseCheckNameRequest) (result kusto.CheckNameResult, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters kusto.Database) (result kusto.DatabasesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (result kusto.DatabasesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (result kusto.Database, err error)
	ListByCluster(ctx context.Context, resourceGroupName string, clusterName string) (result kusto.DatabaseListResult, err error)
	ListPrincipals(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (result kusto.DatabasePrincipalListResult, err error)
	RemovePrincipals(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, databasePrincipalsToRemove kusto.DatabasePrincipalListRequest) (result kusto.DatabasePrincipalListResult, err error)
	Update(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters kusto.DatabaseUpdate) (result kusto.DatabasesUpdateFuture, err error)
}

var _ DatabasesClientAPI = (*kusto.DatabasesClient)(nil)

// EventHubConnectionsClientAPI contains the set of methods on the EventHubConnectionsClient type.
type EventHubConnectionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, eventHubConnectionName string, parameters kusto.EventHubConnection) (result kusto.EventHubConnectionsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, eventHubConnectionName string) (result kusto.EventHubConnectionsDeleteFuture, err error)
	EventhubConnectionValidation(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters kusto.EventHubConnectionValidation) (result kusto.EventHubConnectionValidationListResult, err error)
	Get(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, eventHubConnectionName string) (result kusto.EventHubConnection, err error)
	ListByDatabase(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (result kusto.EventHubConnectionListResult, err error)
	Update(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, eventHubConnectionName string, parameters kusto.EventHubConnectionUpdate) (result kusto.EventHubConnectionsUpdateFuture, err error)
}

var _ EventHubConnectionsClientAPI = (*kusto.EventHubConnectionsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result kusto.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result kusto.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*kusto.OperationsClient)(nil)
