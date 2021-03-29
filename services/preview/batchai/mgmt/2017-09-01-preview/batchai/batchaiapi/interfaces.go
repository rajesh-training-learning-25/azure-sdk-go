package batchaiapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/batchai/mgmt/2017-09-01-preview/batchai"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result batchai.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result batchai.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*batchai.OperationsClient)(nil)

// ClustersClientAPI contains the set of methods on the ClustersClient type.
type ClustersClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, clusterName string, parameters batchai.ClusterCreateParameters) (result batchai.ClustersCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, clusterName string) (result batchai.ClustersDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, clusterName string) (result batchai.Cluster, err error)
	List(ctx context.Context, filter string, selectParameter string, maxResults *int32) (result batchai.ClusterListResultPage, err error)
	ListComplete(ctx context.Context, filter string, selectParameter string, maxResults *int32) (result batchai.ClusterListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, selectParameter string, maxResults *int32) (result batchai.ClusterListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string, selectParameter string, maxResults *int32) (result batchai.ClusterListResultIterator, err error)
	ListRemoteLoginInformation(ctx context.Context, resourceGroupName string, clusterName string) (result batchai.RemoteLoginInformationListResultPage, err error)
	ListRemoteLoginInformationComplete(ctx context.Context, resourceGroupName string, clusterName string) (result batchai.RemoteLoginInformationListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, clusterName string, parameters batchai.ClusterUpdateParameters) (result batchai.Cluster, err error)
}

var _ ClustersClientAPI = (*batchai.ClustersClient)(nil)

// JobsClientAPI contains the set of methods on the JobsClient type.
type JobsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, jobName string, parameters batchai.JobCreateParameters) (result batchai.JobsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, jobName string) (result batchai.JobsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, jobName string) (result batchai.Job, err error)
	List(ctx context.Context, filter string, selectParameter string, maxResults *int32) (result batchai.JobListResultPage, err error)
	ListComplete(ctx context.Context, filter string, selectParameter string, maxResults *int32) (result batchai.JobListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, selectParameter string, maxResults *int32) (result batchai.JobListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string, selectParameter string, maxResults *int32) (result batchai.JobListResultIterator, err error)
	ListOutputFiles(ctx context.Context, resourceGroupName string, jobName string, outputdirectoryid string, linkexpiryinminutes *int32, maxResults *int32) (result batchai.FileListResultPage, err error)
	ListOutputFilesComplete(ctx context.Context, resourceGroupName string, jobName string, outputdirectoryid string, linkexpiryinminutes *int32, maxResults *int32) (result batchai.FileListResultIterator, err error)
	ListRemoteLoginInformation(ctx context.Context, resourceGroupName string, jobName string) (result batchai.RemoteLoginInformationListResultPage, err error)
	ListRemoteLoginInformationComplete(ctx context.Context, resourceGroupName string, jobName string) (result batchai.RemoteLoginInformationListResultIterator, err error)
	Terminate(ctx context.Context, resourceGroupName string, jobName string) (result batchai.JobsTerminateFuture, err error)
}

var _ JobsClientAPI = (*batchai.JobsClient)(nil)

// FileServersClientAPI contains the set of methods on the FileServersClient type.
type FileServersClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, fileServerName string, parameters batchai.FileServerCreateParameters) (result batchai.FileServersCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, fileServerName string) (result batchai.FileServersDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, fileServerName string) (result batchai.FileServer, err error)
	List(ctx context.Context, filter string, selectParameter string, maxResults *int32) (result batchai.FileServerListResultPage, err error)
	ListComplete(ctx context.Context, filter string, selectParameter string, maxResults *int32) (result batchai.FileServerListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, selectParameter string, maxResults *int32) (result batchai.FileServerListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string, selectParameter string, maxResults *int32) (result batchai.FileServerListResultIterator, err error)
}

var _ FileServersClientAPI = (*batchai.FileServersClient)(nil)
