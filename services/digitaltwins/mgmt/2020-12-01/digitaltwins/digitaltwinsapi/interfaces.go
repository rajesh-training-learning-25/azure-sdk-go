package digitaltwinsapi

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
	"github.com/Azure/azure-sdk-for-go/services/digitaltwins/mgmt/2020-12-01/digitaltwins"
)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	CheckNameAvailability(ctx context.Context, location string, digitalTwinsInstanceCheckName digitaltwins.CheckNameRequest) (result digitaltwins.CheckNameResult, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, digitalTwinsCreate digitaltwins.Description) (result digitaltwins.CreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string) (result digitaltwins.DeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string) (result digitaltwins.Description, err error)
	List(ctx context.Context) (result digitaltwins.DescriptionListResultPage, err error)
	ListComplete(ctx context.Context) (result digitaltwins.DescriptionListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result digitaltwins.DescriptionListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result digitaltwins.DescriptionListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, resourceName string, digitalTwinsPatchDescription digitaltwins.PatchDescription) (result digitaltwins.UpdateFuture, err error)
}

var _ ClientAPI = (*digitaltwins.Client)(nil)

// EndpointClientAPI contains the set of methods on the EndpointClient type.
type EndpointClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, endpointName string, endpointDescription digitaltwins.EndpointResource) (result digitaltwins.EndpointCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string, endpointName string) (result digitaltwins.EndpointDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string, endpointName string) (result digitaltwins.EndpointResource, err error)
	List(ctx context.Context, resourceGroupName string, resourceName string) (result digitaltwins.EndpointResourceListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, resourceName string) (result digitaltwins.EndpointResourceListResultIterator, err error)
}

var _ EndpointClientAPI = (*digitaltwins.EndpointClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result digitaltwins.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result digitaltwins.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*digitaltwins.OperationsClient)(nil)

// PrivateLinkResourcesClientAPI contains the set of methods on the PrivateLinkResourcesClient type.
type PrivateLinkResourcesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, resourceName string, resourceID string) (result digitaltwins.GroupIDInformation, err error)
	List(ctx context.Context, resourceGroupName string, resourceName string) (result digitaltwins.GroupIDInformationResponse, err error)
}

var _ PrivateLinkResourcesClientAPI = (*digitaltwins.PrivateLinkResourcesClient)(nil)

// PrivateEndpointConnectionsClientAPI contains the set of methods on the PrivateEndpointConnectionsClient type.
type PrivateEndpointConnectionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, privateEndpointConnection digitaltwins.PrivateEndpointConnection) (result digitaltwins.PrivateEndpointConnectionsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string) (result digitaltwins.PrivateEndpointConnectionsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string) (result digitaltwins.PrivateEndpointConnection, err error)
	List(ctx context.Context, resourceGroupName string, resourceName string) (result digitaltwins.PrivateEndpointConnectionsResponse, err error)
}

var _ PrivateEndpointConnectionsClientAPI = (*digitaltwins.PrivateEndpointConnectionsClient)(nil)
