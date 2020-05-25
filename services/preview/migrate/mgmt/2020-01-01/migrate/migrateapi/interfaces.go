package migrateapi

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
	"github.com/Azure/azure-sdk-for-go/services/preview/migrate/mgmt/2020-01-01/migrate"
	"github.com/Azure/go-autorest/autorest"
)

// HyperVClusterClientAPI contains the set of methods on the HyperVClusterClient type.
type HyperVClusterClientAPI interface {
	GetAllClustersInSite(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string, filter string) (result migrate.HyperVClusterCollectionPage, err error)
	GetAllClustersInSiteComplete(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string, filter string) (result migrate.HyperVClusterCollectionIterator, err error)
	GetCluster(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, clusterName string, APIVersion string) (result migrate.HyperVCluster, err error)
	PutCluster(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, clusterName string, body migrate.HyperVCluster, APIVersion string) (result autorest.Response, err error)
}

var _ HyperVClusterClientAPI = (*migrate.HyperVClusterClient)(nil)

// HyperVHostClientAPI contains the set of methods on the HyperVHostClient type.
type HyperVHostClientAPI interface {
	GetAllHostsInSite(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string, filter string) (result migrate.HyperVHostCollectionPage, err error)
	GetAllHostsInSiteComplete(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string, filter string) (result migrate.HyperVHostCollectionIterator, err error)
	GetHost(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, hostName string, APIVersion string) (result migrate.HyperVHost, err error)
	PutHost(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, hostName string, body migrate.HyperVHost, APIVersion string) (result autorest.Response, err error)
}

var _ HyperVHostClientAPI = (*migrate.HyperVHostClient)(nil)

// HyperVJobsClientAPI contains the set of methods on the HyperVJobsClient type.
type HyperVJobsClientAPI interface {
	GetAllJobsInSite(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string) (result migrate.HyperVJobCollectionPage, err error)
	GetAllJobsInSiteComplete(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string) (result migrate.HyperVJobCollectionIterator, err error)
	GetJob(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, jobName string, APIVersion string) (result migrate.HyperVJob, err error)
}

var _ HyperVJobsClientAPI = (*migrate.HyperVJobsClient)(nil)

// HyperVMachinesClientAPI contains the set of methods on the HyperVMachinesClient type.
type HyperVMachinesClientAPI interface {
	GetAllMachinesInSite(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string, filter string, top *int32, continuationToken string, totalRecordCount *int32) (result migrate.HyperVMachineCollectionPage, err error)
	GetAllMachinesInSiteComplete(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string, filter string, top *int32, continuationToken string, totalRecordCount *int32) (result migrate.HyperVMachineCollectionIterator, err error)
	GetMachine(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, machineName string, APIVersion string) (result migrate.HyperVMachine, err error)
}

var _ HyperVMachinesClientAPI = (*migrate.HyperVMachinesClient)(nil)

// HyperVOperationsStatusClientAPI contains the set of methods on the HyperVOperationsStatusClient type.
type HyperVOperationsStatusClientAPI interface {
	GetOperationStatus(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, operationStatusName string, APIVersion string) (result migrate.OperationStatus, err error)
}

var _ HyperVOperationsStatusClientAPI = (*migrate.HyperVOperationsStatusClient)(nil)

// HyperVRunAsAccountsClientAPI contains the set of methods on the HyperVRunAsAccountsClient type.
type HyperVRunAsAccountsClientAPI interface {
	GetAllRunAsAccountsInSite(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string) (result migrate.HyperVRunAsAccountCollectionPage, err error)
	GetAllRunAsAccountsInSiteComplete(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string) (result migrate.HyperVRunAsAccountCollectionIterator, err error)
	GetRunAsAccount(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, accountName string, APIVersion string) (result migrate.HyperVRunAsAccount, err error)
}

var _ HyperVRunAsAccountsClientAPI = (*migrate.HyperVRunAsAccountsClient)(nil)

// HyperVSitesClientAPI contains the set of methods on the HyperVSitesClient type.
type HyperVSitesClientAPI interface {
	DeleteSite(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string) (result autorest.Response, err error)
	GetSite(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string) (result migrate.HyperVSite, err error)
	GetSiteHealthSummary(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string) (result migrate.SiteHealthSummaryCollectionPage, err error)
	GetSiteHealthSummaryComplete(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string) (result migrate.SiteHealthSummaryCollectionIterator, err error)
	GetSiteUsage(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string) (result migrate.HyperVSiteUsage, err error)
	PatchSite(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, body migrate.HyperVSite, APIVersion string) (result migrate.HyperVSite, err error)
	PutSite(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, body migrate.HyperVSite, APIVersion string) (result migrate.HyperVSite, err error)
	RefreshSite(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string) (result autorest.Response, err error)
}

var _ HyperVSitesClientAPI = (*migrate.HyperVSitesClient)(nil)

// JobsClientAPI contains the set of methods on the JobsClient type.
type JobsClientAPI interface {
	GetAllJobsInSite(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string) (result migrate.VMwareJobCollectionPage, err error)
	GetAllJobsInSiteComplete(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string) (result migrate.VMwareJobCollectionIterator, err error)
	GetJob(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, jobName string, APIVersion string) (result migrate.VMwareJob, err error)
}

var _ JobsClientAPI = (*migrate.JobsClient)(nil)

// MachinesClientAPI contains the set of methods on the MachinesClient type.
type MachinesClientAPI interface {
	GetAllMachinesInSite(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string, filter string, top *int32, continuationToken string, totalRecordCount *int32) (result migrate.VMwareMachineCollectionPage, err error)
	GetAllMachinesInSiteComplete(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string, filter string, top *int32, continuationToken string, totalRecordCount *int32) (result migrate.VMwareMachineCollectionIterator, err error)
	GetMachine(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, machineName string, APIVersion string) (result migrate.VMwareMachine, err error)
	StartMachine(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, machineName string, APIVersion string) (result autorest.Response, err error)
	StopMachine(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, machineName string, APIVersion string) (result autorest.Response, err error)
}

var _ MachinesClientAPI = (*migrate.MachinesClient)(nil)

// RunAsAccountsClientAPI contains the set of methods on the RunAsAccountsClient type.
type RunAsAccountsClientAPI interface {
	GetAllRunAsAccountsInSite(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string) (result migrate.VMwareRunAsAccountCollectionPage, err error)
	GetAllRunAsAccountsInSiteComplete(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string) (result migrate.VMwareRunAsAccountCollectionIterator, err error)
	GetRunAsAccount(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, accountName string, APIVersion string) (result migrate.VMwareRunAsAccount, err error)
}

var _ RunAsAccountsClientAPI = (*migrate.RunAsAccountsClient)(nil)

// SitesClientAPI contains the set of methods on the SitesClient type.
type SitesClientAPI interface {
	DeleteSite(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string) (result autorest.Response, err error)
	GetSite(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string) (result migrate.VMwareSite, err error)
	GetSiteHealthSummary(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string) (result migrate.SiteHealthSummaryCollectionPage, err error)
	GetSiteHealthSummaryComplete(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string) (result migrate.SiteHealthSummaryCollectionIterator, err error)
	GetSiteUsage(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string) (result migrate.VMwareSiteUsage, err error)
	PatchSite(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, body migrate.VMwareSite, APIVersion string) (result migrate.VMwareSite, err error)
	PutSite(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, body migrate.VMwareSite, APIVersion string) (result migrate.VMwareSite, err error)
	RefreshSite(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string) (result autorest.Response, err error)
}

var _ SitesClientAPI = (*migrate.SitesClient)(nil)

// VCenterClientAPI contains the set of methods on the VCenterClient type.
type VCenterClientAPI interface {
	DeleteVCenter(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, vcenterName string, APIVersion string) (result autorest.Response, err error)
	GetAllVCentersInSite(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string, filter string) (result migrate.VCenterCollectionPage, err error)
	GetAllVCentersInSiteComplete(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string, filter string) (result migrate.VCenterCollectionIterator, err error)
	GetVCenter(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, vcenterName string, APIVersion string) (result migrate.VCenter, err error)
	PutVCenter(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, vcenterName string, body migrate.VCenter, APIVersion string) (result autorest.Response, err error)
}

var _ VCenterClientAPI = (*migrate.VCenterClient)(nil)

// VMwareOperationsStatusClientAPI contains the set of methods on the VMwareOperationsStatusClient type.
type VMwareOperationsStatusClientAPI interface {
	GetOperationStatus(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, operationStatusName string, APIVersion string) (result migrate.OperationStatus, err error)
}

var _ VMwareOperationsStatusClientAPI = (*migrate.VMwareOperationsStatusClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context, APIVersion string) (result migrate.OperationResultListPage, err error)
	ListComplete(ctx context.Context, APIVersion string) (result migrate.OperationResultListIterator, err error)
}

var _ OperationsClientAPI = (*migrate.OperationsClient)(nil)
