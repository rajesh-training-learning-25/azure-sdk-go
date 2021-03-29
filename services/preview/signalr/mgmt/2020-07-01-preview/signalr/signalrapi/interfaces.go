package signalrapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/signalr/mgmt/2020-07-01-preview/signalr"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result signalr.OperationListPage, err error)
	ListComplete(ctx context.Context) (result signalr.OperationListIterator, err error)
}

var _ OperationsClientAPI = (*signalr.OperationsClient)(nil)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	CheckNameAvailability(ctx context.Context, location string, parameters *signalr.NameAvailabilityParameters) (result signalr.NameAvailability, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters *signalr.ResourceType) (result signalr.CreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string) (result signalr.DeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string) (result signalr.ResourceType, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result signalr.ResourceListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result signalr.ResourceListIterator, err error)
	ListBySubscription(ctx context.Context) (result signalr.ResourceListPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result signalr.ResourceListIterator, err error)
	ListKeys(ctx context.Context, resourceGroupName string, resourceName string) (result signalr.Keys, err error)
	RegenerateKey(ctx context.Context, resourceGroupName string, resourceName string, parameters *signalr.RegenerateKeyParameters) (result signalr.RegenerateKeyFuture, err error)
	Restart(ctx context.Context, resourceGroupName string, resourceName string) (result signalr.RestartFuture, err error)
	Update(ctx context.Context, resourceGroupName string, resourceName string, parameters *signalr.ResourceType) (result signalr.UpdateFuture, err error)
}

var _ ClientAPI = (*signalr.Client)(nil)

// PrivateEndpointConnectionsClientAPI contains the set of methods on the PrivateEndpointConnectionsClient type.
type PrivateEndpointConnectionsClientAPI interface {
	Delete(ctx context.Context, privateEndpointConnectionName string, resourceGroupName string, resourceName string) (result signalr.PrivateEndpointConnectionsDeleteFuture, err error)
	Get(ctx context.Context, privateEndpointConnectionName string, resourceGroupName string, resourceName string) (result signalr.PrivateEndpointConnection, err error)
	Update(ctx context.Context, privateEndpointConnectionName string, resourceGroupName string, resourceName string, parameters *signalr.PrivateEndpointConnection) (result signalr.PrivateEndpointConnection, err error)
}

var _ PrivateEndpointConnectionsClientAPI = (*signalr.PrivateEndpointConnectionsClient)(nil)

// PrivateLinkResourcesClientAPI contains the set of methods on the PrivateLinkResourcesClient type.
type PrivateLinkResourcesClientAPI interface {
	List(ctx context.Context, resourceGroupName string, resourceName string) (result signalr.PrivateLinkResourceListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, resourceName string) (result signalr.PrivateLinkResourceListIterator, err error)
}

var _ PrivateLinkResourcesClientAPI = (*signalr.PrivateLinkResourcesClient)(nil)

// UsagesClientAPI contains the set of methods on the UsagesClient type.
type UsagesClientAPI interface {
	List(ctx context.Context, location string) (result signalr.UsageListPage, err error)
	ListComplete(ctx context.Context, location string) (result signalr.UsageListIterator, err error)
}

var _ UsagesClientAPI = (*signalr.UsagesClient)(nil)
