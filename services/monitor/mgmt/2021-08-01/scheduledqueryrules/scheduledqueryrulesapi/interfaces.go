package scheduledqueryrulesapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/monitor/mgmt/2021-08-01/scheduledqueryrules"
	"github.com/Azure/go-autorest/autorest"
)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, ruleName string, parameters scheduledqueryrules.ResourceType) (result scheduledqueryrules.ResourceType, err error)
	Delete(ctx context.Context, resourceGroupName string, ruleName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, ruleName string) (result scheduledqueryrules.ResourceType, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result scheduledqueryrules.ResourceCollectionPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result scheduledqueryrules.ResourceCollectionIterator, err error)
	ListBySubscription(ctx context.Context) (result scheduledqueryrules.ResourceCollectionPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result scheduledqueryrules.ResourceCollectionIterator, err error)
	Update(ctx context.Context, resourceGroupName string, ruleName string, parameters scheduledqueryrules.ResourcePatch) (result scheduledqueryrules.ResourceType, err error)
}

var _ ClientAPI = (*scheduledqueryrules.Client)(nil)
