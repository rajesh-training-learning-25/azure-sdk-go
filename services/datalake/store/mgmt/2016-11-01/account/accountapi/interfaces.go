package accountapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/datalake/store/mgmt/2016-11-01/account"
	"github.com/Azure/go-autorest/autorest"
)

// AccountsClientAPI contains the set of methods on the AccountsClient type.
type AccountsClientAPI interface {
	CheckNameAvailability(ctx context.Context, location string, parameters account.CheckNameAvailabilityParameters) (result account.NameAvailabilityInformation, err error)
	Create(ctx context.Context, resourceGroupName string, accountName string, parameters account.CreateDataLakeStoreAccountParameters) (result account.AccountsCreateFutureType, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string) (result account.AccountsDeleteFutureType, err error)
	EnableKeyVault(ctx context.Context, resourceGroupName string, accountName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string) (result account.DataLakeStoreAccount, err error)
	List(ctx context.Context, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result account.DataLakeStoreAccountListResultPage, err error)
	ListComplete(ctx context.Context, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result account.DataLakeStoreAccountListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result account.DataLakeStoreAccountListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result account.DataLakeStoreAccountListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, parameters account.UpdateDataLakeStoreAccountParameters) (result account.AccountsUpdateFutureType, err error)
}

var _ AccountsClientAPI = (*account.AccountsClient)(nil)

// FirewallRulesClientAPI contains the set of methods on the FirewallRulesClient type.
type FirewallRulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, firewallRuleName string, parameters account.CreateOrUpdateFirewallRuleParameters) (result account.FirewallRule, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, firewallRuleName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, firewallRuleName string) (result account.FirewallRule, err error)
	ListByAccount(ctx context.Context, resourceGroupName string, accountName string) (result account.FirewallRuleListResultPage, err error)
	ListByAccountComplete(ctx context.Context, resourceGroupName string, accountName string) (result account.FirewallRuleListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, firewallRuleName string, parameters *account.UpdateFirewallRuleParameters) (result account.FirewallRule, err error)
}

var _ FirewallRulesClientAPI = (*account.FirewallRulesClient)(nil)

// VirtualNetworkRulesClientAPI contains the set of methods on the VirtualNetworkRulesClient type.
type VirtualNetworkRulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, virtualNetworkRuleName string, parameters account.CreateOrUpdateVirtualNetworkRuleParameters) (result account.VirtualNetworkRule, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, virtualNetworkRuleName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, virtualNetworkRuleName string) (result account.VirtualNetworkRule, err error)
	ListByAccount(ctx context.Context, resourceGroupName string, accountName string) (result account.VirtualNetworkRuleListResultPage, err error)
	ListByAccountComplete(ctx context.Context, resourceGroupName string, accountName string) (result account.VirtualNetworkRuleListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, virtualNetworkRuleName string, parameters *account.UpdateVirtualNetworkRuleParameters) (result account.VirtualNetworkRule, err error)
}

var _ VirtualNetworkRulesClientAPI = (*account.VirtualNetworkRulesClient)(nil)

// TrustedIDProvidersClientAPI contains the set of methods on the TrustedIDProvidersClient type.
type TrustedIDProvidersClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, trustedIDProviderName string, parameters account.CreateOrUpdateTrustedIDProviderParameters) (result account.TrustedIDProvider, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, trustedIDProviderName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, trustedIDProviderName string) (result account.TrustedIDProvider, err error)
	ListByAccount(ctx context.Context, resourceGroupName string, accountName string) (result account.TrustedIDProviderListResultPage, err error)
	ListByAccountComplete(ctx context.Context, resourceGroupName string, accountName string) (result account.TrustedIDProviderListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, trustedIDProviderName string, parameters *account.UpdateTrustedIDProviderParameters) (result account.TrustedIDProvider, err error)
}

var _ TrustedIDProvidersClientAPI = (*account.TrustedIDProvidersClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result account.OperationListResult, err error)
}

var _ OperationsClientAPI = (*account.OperationsClient)(nil)

// LocationsClientAPI contains the set of methods on the LocationsClient type.
type LocationsClientAPI interface {
	GetCapability(ctx context.Context, location string) (result account.CapabilityInformation, err error)
	GetUsage(ctx context.Context, location string) (result account.UsageListResult, err error)
}

var _ LocationsClientAPI = (*account.LocationsClient)(nil)
