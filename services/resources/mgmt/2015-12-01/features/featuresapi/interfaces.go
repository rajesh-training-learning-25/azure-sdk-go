package featuresapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2015-12-01/features"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	ListOperations(ctx context.Context) (result features.OperationListResultPage, err error)
	ListOperationsComplete(ctx context.Context) (result features.OperationListResultIterator, err error)
}

var _ BaseClientAPI = (*features.BaseClient)(nil)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	Get(ctx context.Context, resourceProviderNamespace string, featureName string) (result features.Result, err error)
	List(ctx context.Context, resourceProviderNamespace string) (result features.OperationsListResultPage, err error)
	ListComplete(ctx context.Context, resourceProviderNamespace string) (result features.OperationsListResultIterator, err error)
	ListAll(ctx context.Context) (result features.OperationsListResultPage, err error)
	ListAllComplete(ctx context.Context) (result features.OperationsListResultIterator, err error)
	Register(ctx context.Context, resourceProviderNamespace string, featureName string) (result features.Result, err error)
	Unregister(ctx context.Context, resourceProviderNamespace string, featureName string) (result features.Result, err error)
}

var _ ClientAPI = (*features.Client)(nil)
