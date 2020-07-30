package datashareapi

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
	"github.com/Azure/azure-sdk-for-go/services/preview/datashare/mgmt/2018-11-01-preview/datashare"
	"github.com/Azure/go-autorest/autorest"
)

// AccountsClientAPI contains the set of methods on the AccountsClient type.
type AccountsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, accountName string, account datashare.Account) (result datashare.AccountsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string) (result datashare.AccountsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string) (result datashare.Account, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, skipToken string) (result datashare.AccountListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, skipToken string) (result datashare.AccountListIterator, err error)
	ListBySubscription(ctx context.Context, skipToken string) (result datashare.AccountListPage, err error)
	ListBySubscriptionComplete(ctx context.Context, skipToken string) (result datashare.AccountListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, accountUpdateParameters datashare.AccountUpdateParameters) (result datashare.Account, err error)
}

var _ AccountsClientAPI = (*datashare.AccountsClient)(nil)

// ConsumerInvitationsClientAPI contains the set of methods on the ConsumerInvitationsClient type.
type ConsumerInvitationsClientAPI interface {
	Get(ctx context.Context, location string, invitationID string) (result datashare.ConsumerInvitation, err error)
	ListInvitations(ctx context.Context, skipToken string) (result datashare.ConsumerInvitationListPage, err error)
	ListInvitationsComplete(ctx context.Context, skipToken string) (result datashare.ConsumerInvitationListIterator, err error)
	RejectInvitation(ctx context.Context, location string, invitation datashare.ConsumerInvitation) (result datashare.ConsumerInvitation, err error)
}

var _ ConsumerInvitationsClientAPI = (*datashare.ConsumerInvitationsClient)(nil)

// DataSetsClientAPI contains the set of methods on the DataSetsClient type.
type DataSetsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, accountName string, shareName string, dataSetName string, dataSet datashare.BasicDataSet) (result datashare.DataSetModel, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, shareName string, dataSetName string) (result datashare.DataSetsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, shareName string, dataSetName string) (result datashare.DataSetModel, err error)
	ListByShare(ctx context.Context, resourceGroupName string, accountName string, shareName string, skipToken string, filter string, orderby string) (result datashare.DataSetListPage, err error)
	ListByShareComplete(ctx context.Context, resourceGroupName string, accountName string, shareName string, skipToken string, filter string, orderby string) (result datashare.DataSetListIterator, err error)
}

var _ DataSetsClientAPI = (*datashare.DataSetsClient)(nil)

// DataSetMappingsClientAPI contains the set of methods on the DataSetMappingsClient type.
type DataSetMappingsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, dataSetMappingName string, dataSetMapping datashare.BasicDataSetMapping) (result datashare.DataSetMappingModel, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, dataSetMappingName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, dataSetMappingName string) (result datashare.DataSetMappingModel, err error)
	ListByShareSubscription(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, skipToken string, filter string, orderby string) (result datashare.DataSetMappingListPage, err error)
	ListByShareSubscriptionComplete(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, skipToken string, filter string, orderby string) (result datashare.DataSetMappingListIterator, err error)
}

var _ DataSetMappingsClientAPI = (*datashare.DataSetMappingsClient)(nil)

// InvitationsClientAPI contains the set of methods on the InvitationsClient type.
type InvitationsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, accountName string, shareName string, invitationName string, invitation datashare.Invitation) (result datashare.Invitation, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, shareName string, invitationName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, shareName string, invitationName string) (result datashare.Invitation, err error)
	ListByShare(ctx context.Context, resourceGroupName string, accountName string, shareName string, skipToken string, filter string, orderby string) (result datashare.InvitationListPage, err error)
	ListByShareComplete(ctx context.Context, resourceGroupName string, accountName string, shareName string, skipToken string, filter string, orderby string) (result datashare.InvitationListIterator, err error)
}

var _ InvitationsClientAPI = (*datashare.InvitationsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result datashare.OperationListPage, err error)
	ListComplete(ctx context.Context) (result datashare.OperationListIterator, err error)
}

var _ OperationsClientAPI = (*datashare.OperationsClient)(nil)

// SharesClientAPI contains the set of methods on the SharesClient type.
type SharesClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, accountName string, shareName string, share datashare.Share) (result datashare.Share, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, shareName string) (result datashare.SharesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, shareName string) (result datashare.Share, err error)
	ListByAccount(ctx context.Context, resourceGroupName string, accountName string, skipToken string, filter string, orderby string) (result datashare.ShareListPage, err error)
	ListByAccountComplete(ctx context.Context, resourceGroupName string, accountName string, skipToken string, filter string, orderby string) (result datashare.ShareListIterator, err error)
	ListSynchronizationDetails(ctx context.Context, resourceGroupName string, accountName string, shareName string, shareSynchronization datashare.ShareSynchronization, skipToken string, filter string, orderby string) (result datashare.SynchronizationDetailsListPage, err error)
	ListSynchronizationDetailsComplete(ctx context.Context, resourceGroupName string, accountName string, shareName string, shareSynchronization datashare.ShareSynchronization, skipToken string, filter string, orderby string) (result datashare.SynchronizationDetailsListIterator, err error)
	ListSynchronizations(ctx context.Context, resourceGroupName string, accountName string, shareName string, skipToken string, filter string, orderby string) (result datashare.ShareSynchronizationListPage, err error)
	ListSynchronizationsComplete(ctx context.Context, resourceGroupName string, accountName string, shareName string, skipToken string, filter string, orderby string) (result datashare.ShareSynchronizationListIterator, err error)
}

var _ SharesClientAPI = (*datashare.SharesClient)(nil)

// ProviderShareSubscriptionsClientAPI contains the set of methods on the ProviderShareSubscriptionsClient type.
type ProviderShareSubscriptionsClientAPI interface {
	GetByShare(ctx context.Context, resourceGroupName string, accountName string, shareName string, providerShareSubscriptionID string) (result datashare.ProviderShareSubscription, err error)
	ListByShare(ctx context.Context, resourceGroupName string, accountName string, shareName string, skipToken string) (result datashare.ProviderShareSubscriptionListPage, err error)
	ListByShareComplete(ctx context.Context, resourceGroupName string, accountName string, shareName string, skipToken string) (result datashare.ProviderShareSubscriptionListIterator, err error)
	Reinstate(ctx context.Context, resourceGroupName string, accountName string, shareName string, providerShareSubscriptionID string) (result datashare.ProviderShareSubscription, err error)
	Revoke(ctx context.Context, resourceGroupName string, accountName string, shareName string, providerShareSubscriptionID string) (result datashare.ProviderShareSubscriptionsRevokeFuture, err error)
}

var _ ProviderShareSubscriptionsClientAPI = (*datashare.ProviderShareSubscriptionsClient)(nil)

// ShareSubscriptionsClientAPI contains the set of methods on the ShareSubscriptionsClient type.
type ShareSubscriptionsClientAPI interface {
	CancelSynchronization(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, shareSubscriptionSynchronization datashare.ShareSubscriptionSynchronization) (result datashare.ShareSubscriptionsCancelSynchronizationFuture, err error)
	Create(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, shareSubscription datashare.ShareSubscription) (result datashare.ShareSubscription, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string) (result datashare.ShareSubscriptionsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string) (result datashare.ShareSubscription, err error)
	ListByAccount(ctx context.Context, resourceGroupName string, accountName string, skipToken string, filter string, orderby string) (result datashare.ShareSubscriptionListPage, err error)
	ListByAccountComplete(ctx context.Context, resourceGroupName string, accountName string, skipToken string, filter string, orderby string) (result datashare.ShareSubscriptionListIterator, err error)
	ListSourceShareSynchronizationSettings(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, skipToken string) (result datashare.SourceShareSynchronizationSettingListPage, err error)
	ListSourceShareSynchronizationSettingsComplete(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, skipToken string) (result datashare.SourceShareSynchronizationSettingListIterator, err error)
	ListSynchronizationDetails(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, shareSubscriptionSynchronization datashare.ShareSubscriptionSynchronization, skipToken string, filter string, orderby string) (result datashare.SynchronizationDetailsListPage, err error)
	ListSynchronizationDetailsComplete(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, shareSubscriptionSynchronization datashare.ShareSubscriptionSynchronization, skipToken string, filter string, orderby string) (result datashare.SynchronizationDetailsListIterator, err error)
	ListSynchronizations(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, skipToken string, filter string, orderby string) (result datashare.ShareSubscriptionSynchronizationListPage, err error)
	ListSynchronizationsComplete(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, skipToken string, filter string, orderby string) (result datashare.ShareSubscriptionSynchronizationListIterator, err error)
	SynchronizeMethod(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, synchronize datashare.Synchronize) (result datashare.ShareSubscriptionsSynchronizeMethodFuture, err error)
}

var _ ShareSubscriptionsClientAPI = (*datashare.ShareSubscriptionsClient)(nil)

// ConsumerSourceDataSetsClientAPI contains the set of methods on the ConsumerSourceDataSetsClient type.
type ConsumerSourceDataSetsClientAPI interface {
	ListByShareSubscription(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, skipToken string) (result datashare.ConsumerSourceDataSetListPage, err error)
	ListByShareSubscriptionComplete(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, skipToken string) (result datashare.ConsumerSourceDataSetListIterator, err error)
}

var _ ConsumerSourceDataSetsClientAPI = (*datashare.ConsumerSourceDataSetsClient)(nil)

// SynchronizationSettingsClientAPI contains the set of methods on the SynchronizationSettingsClient type.
type SynchronizationSettingsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, accountName string, shareName string, synchronizationSettingName string, synchronizationSetting datashare.BasicSynchronizationSetting) (result datashare.SynchronizationSettingModel, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, shareName string, synchronizationSettingName string) (result datashare.SynchronizationSettingsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, shareName string, synchronizationSettingName string) (result datashare.SynchronizationSettingModel, err error)
	ListByShare(ctx context.Context, resourceGroupName string, accountName string, shareName string, skipToken string) (result datashare.SynchronizationSettingListPage, err error)
	ListByShareComplete(ctx context.Context, resourceGroupName string, accountName string, shareName string, skipToken string) (result datashare.SynchronizationSettingListIterator, err error)
}

var _ SynchronizationSettingsClientAPI = (*datashare.SynchronizationSettingsClient)(nil)

// TriggersClientAPI contains the set of methods on the TriggersClient type.
type TriggersClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, triggerName string, trigger datashare.BasicTrigger) (result datashare.TriggersCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, triggerName string) (result datashare.TriggersDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, triggerName string) (result datashare.TriggerModel, err error)
	ListByShareSubscription(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, skipToken string) (result datashare.TriggerListPage, err error)
	ListByShareSubscriptionComplete(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, skipToken string) (result datashare.TriggerListIterator, err error)
}

var _ TriggersClientAPI = (*datashare.TriggersClient)(nil)
