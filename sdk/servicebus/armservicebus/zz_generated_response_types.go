//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicebus

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// DisasterRecoveryConfigsBreakPairingResponse contains the response from method DisasterRecoveryConfigs.BreakPairing.
type DisasterRecoveryConfigsBreakPairingResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DisasterRecoveryConfigsCheckNameAvailabilityResponse contains the response from method DisasterRecoveryConfigs.CheckNameAvailability.
type DisasterRecoveryConfigsCheckNameAvailabilityResponse struct {
	DisasterRecoveryConfigsCheckNameAvailabilityResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DisasterRecoveryConfigsCheckNameAvailabilityResult contains the result from method DisasterRecoveryConfigs.CheckNameAvailability.
type DisasterRecoveryConfigsCheckNameAvailabilityResult struct {
	CheckNameAvailabilityResult
}

// DisasterRecoveryConfigsCreateOrUpdateResponse contains the response from method DisasterRecoveryConfigs.CreateOrUpdate.
type DisasterRecoveryConfigsCreateOrUpdateResponse struct {
	DisasterRecoveryConfigsCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DisasterRecoveryConfigsCreateOrUpdateResult contains the result from method DisasterRecoveryConfigs.CreateOrUpdate.
type DisasterRecoveryConfigsCreateOrUpdateResult struct {
	ArmDisasterRecovery
}

// DisasterRecoveryConfigsDeleteResponse contains the response from method DisasterRecoveryConfigs.Delete.
type DisasterRecoveryConfigsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DisasterRecoveryConfigsFailOverResponse contains the response from method DisasterRecoveryConfigs.FailOver.
type DisasterRecoveryConfigsFailOverResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DisasterRecoveryConfigsGetAuthorizationRuleResponse contains the response from method DisasterRecoveryConfigs.GetAuthorizationRule.
type DisasterRecoveryConfigsGetAuthorizationRuleResponse struct {
	DisasterRecoveryConfigsGetAuthorizationRuleResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DisasterRecoveryConfigsGetAuthorizationRuleResult contains the result from method DisasterRecoveryConfigs.GetAuthorizationRule.
type DisasterRecoveryConfigsGetAuthorizationRuleResult struct {
	SBAuthorizationRule
}

// DisasterRecoveryConfigsGetResponse contains the response from method DisasterRecoveryConfigs.Get.
type DisasterRecoveryConfigsGetResponse struct {
	DisasterRecoveryConfigsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DisasterRecoveryConfigsGetResult contains the result from method DisasterRecoveryConfigs.Get.
type DisasterRecoveryConfigsGetResult struct {
	ArmDisasterRecovery
}

// DisasterRecoveryConfigsListAuthorizationRulesResponse contains the response from method DisasterRecoveryConfigs.ListAuthorizationRules.
type DisasterRecoveryConfigsListAuthorizationRulesResponse struct {
	DisasterRecoveryConfigsListAuthorizationRulesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DisasterRecoveryConfigsListAuthorizationRulesResult contains the result from method DisasterRecoveryConfigs.ListAuthorizationRules.
type DisasterRecoveryConfigsListAuthorizationRulesResult struct {
	SBAuthorizationRuleListResult
}

// DisasterRecoveryConfigsListKeysResponse contains the response from method DisasterRecoveryConfigs.ListKeys.
type DisasterRecoveryConfigsListKeysResponse struct {
	DisasterRecoveryConfigsListKeysResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DisasterRecoveryConfigsListKeysResult contains the result from method DisasterRecoveryConfigs.ListKeys.
type DisasterRecoveryConfigsListKeysResult struct {
	AccessKeys
}

// DisasterRecoveryConfigsListResponse contains the response from method DisasterRecoveryConfigs.List.
type DisasterRecoveryConfigsListResponse struct {
	DisasterRecoveryConfigsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DisasterRecoveryConfigsListResult contains the result from method DisasterRecoveryConfigs.List.
type DisasterRecoveryConfigsListResult struct {
	ArmDisasterRecoveryListResult
}

// EventHubsListByNamespaceResponse contains the response from method EventHubs.ListByNamespace.
type EventHubsListByNamespaceResponse struct {
	EventHubsListByNamespaceResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EventHubsListByNamespaceResult contains the result from method EventHubs.ListByNamespace.
type EventHubsListByNamespaceResult struct {
	EventHubListResult
}

// MigrationConfigsCompleteMigrationResponse contains the response from method MigrationConfigs.CompleteMigration.
type MigrationConfigsCompleteMigrationResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MigrationConfigsCreateAndStartMigrationPollerResponse contains the response from method MigrationConfigs.CreateAndStartMigration.
type MigrationConfigsCreateAndStartMigrationPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *MigrationConfigsCreateAndStartMigrationPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l MigrationConfigsCreateAndStartMigrationPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (MigrationConfigsCreateAndStartMigrationResponse, error) {
	respType := MigrationConfigsCreateAndStartMigrationResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.MigrationConfigProperties)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a MigrationConfigsCreateAndStartMigrationPollerResponse from the provided client and resume token.
func (l *MigrationConfigsCreateAndStartMigrationPollerResponse) Resume(ctx context.Context, client *MigrationConfigsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("MigrationConfigsClient.CreateAndStartMigration", token, client.pl, client.createAndStartMigrationHandleError)
	if err != nil {
		return err
	}
	poller := &MigrationConfigsCreateAndStartMigrationPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// MigrationConfigsCreateAndStartMigrationResponse contains the response from method MigrationConfigs.CreateAndStartMigration.
type MigrationConfigsCreateAndStartMigrationResponse struct {
	MigrationConfigsCreateAndStartMigrationResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MigrationConfigsCreateAndStartMigrationResult contains the result from method MigrationConfigs.CreateAndStartMigration.
type MigrationConfigsCreateAndStartMigrationResult struct {
	MigrationConfigProperties
}

// MigrationConfigsDeleteResponse contains the response from method MigrationConfigs.Delete.
type MigrationConfigsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MigrationConfigsGetResponse contains the response from method MigrationConfigs.Get.
type MigrationConfigsGetResponse struct {
	MigrationConfigsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MigrationConfigsGetResult contains the result from method MigrationConfigs.Get.
type MigrationConfigsGetResult struct {
	MigrationConfigProperties
}

// MigrationConfigsListResponse contains the response from method MigrationConfigs.List.
type MigrationConfigsListResponse struct {
	MigrationConfigsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MigrationConfigsListResult contains the result from method MigrationConfigs.List.
type MigrationConfigsListResult struct {
	MigrationConfigListResult
}

// MigrationConfigsRevertResponse contains the response from method MigrationConfigs.Revert.
type MigrationConfigsRevertResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesCheckNameAvailabilityResponse contains the response from method Namespaces.CheckNameAvailability.
type NamespacesCheckNameAvailabilityResponse struct {
	NamespacesCheckNameAvailabilityResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesCheckNameAvailabilityResult contains the result from method Namespaces.CheckNameAvailability.
type NamespacesCheckNameAvailabilityResult struct {
	CheckNameAvailabilityResult
}

// NamespacesCreateOrUpdateAuthorizationRuleResponse contains the response from method Namespaces.CreateOrUpdateAuthorizationRule.
type NamespacesCreateOrUpdateAuthorizationRuleResponse struct {
	NamespacesCreateOrUpdateAuthorizationRuleResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesCreateOrUpdateAuthorizationRuleResult contains the result from method Namespaces.CreateOrUpdateAuthorizationRule.
type NamespacesCreateOrUpdateAuthorizationRuleResult struct {
	SBAuthorizationRule
}

// NamespacesCreateOrUpdateIPFilterRuleResponse contains the response from method Namespaces.CreateOrUpdateIPFilterRule.
type NamespacesCreateOrUpdateIPFilterRuleResponse struct {
	NamespacesCreateOrUpdateIPFilterRuleResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesCreateOrUpdateIPFilterRuleResult contains the result from method Namespaces.CreateOrUpdateIPFilterRule.
type NamespacesCreateOrUpdateIPFilterRuleResult struct {
	IPFilterRule
}

// NamespacesCreateOrUpdateNetworkRuleSetResponse contains the response from method Namespaces.CreateOrUpdateNetworkRuleSet.
type NamespacesCreateOrUpdateNetworkRuleSetResponse struct {
	NamespacesCreateOrUpdateNetworkRuleSetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesCreateOrUpdateNetworkRuleSetResult contains the result from method Namespaces.CreateOrUpdateNetworkRuleSet.
type NamespacesCreateOrUpdateNetworkRuleSetResult struct {
	NetworkRuleSet
}

// NamespacesCreateOrUpdatePollerResponse contains the response from method Namespaces.CreateOrUpdate.
type NamespacesCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *NamespacesCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l NamespacesCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (NamespacesCreateOrUpdateResponse, error) {
	respType := NamespacesCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.SBNamespace)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a NamespacesCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *NamespacesCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *NamespacesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("NamespacesClient.CreateOrUpdate", token, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return err
	}
	poller := &NamespacesCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// NamespacesCreateOrUpdateResponse contains the response from method Namespaces.CreateOrUpdate.
type NamespacesCreateOrUpdateResponse struct {
	NamespacesCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesCreateOrUpdateResult contains the result from method Namespaces.CreateOrUpdate.
type NamespacesCreateOrUpdateResult struct {
	SBNamespace
}

// NamespacesCreateOrUpdateVirtualNetworkRuleResponse contains the response from method Namespaces.CreateOrUpdateVirtualNetworkRule.
type NamespacesCreateOrUpdateVirtualNetworkRuleResponse struct {
	NamespacesCreateOrUpdateVirtualNetworkRuleResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesCreateOrUpdateVirtualNetworkRuleResult contains the result from method Namespaces.CreateOrUpdateVirtualNetworkRule.
type NamespacesCreateOrUpdateVirtualNetworkRuleResult struct {
	VirtualNetworkRule
}

// NamespacesDeleteAuthorizationRuleResponse contains the response from method Namespaces.DeleteAuthorizationRule.
type NamespacesDeleteAuthorizationRuleResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesDeleteIPFilterRuleResponse contains the response from method Namespaces.DeleteIPFilterRule.
type NamespacesDeleteIPFilterRuleResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesDeletePollerResponse contains the response from method Namespaces.Delete.
type NamespacesDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *NamespacesDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l NamespacesDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (NamespacesDeleteResponse, error) {
	respType := NamespacesDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a NamespacesDeletePollerResponse from the provided client and resume token.
func (l *NamespacesDeletePollerResponse) Resume(ctx context.Context, client *NamespacesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("NamespacesClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &NamespacesDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// NamespacesDeleteResponse contains the response from method Namespaces.Delete.
type NamespacesDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesDeleteVirtualNetworkRuleResponse contains the response from method Namespaces.DeleteVirtualNetworkRule.
type NamespacesDeleteVirtualNetworkRuleResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesGetAuthorizationRuleResponse contains the response from method Namespaces.GetAuthorizationRule.
type NamespacesGetAuthorizationRuleResponse struct {
	NamespacesGetAuthorizationRuleResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesGetAuthorizationRuleResult contains the result from method Namespaces.GetAuthorizationRule.
type NamespacesGetAuthorizationRuleResult struct {
	SBAuthorizationRule
}

// NamespacesGetIPFilterRuleResponse contains the response from method Namespaces.GetIPFilterRule.
type NamespacesGetIPFilterRuleResponse struct {
	NamespacesGetIPFilterRuleResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesGetIPFilterRuleResult contains the result from method Namespaces.GetIPFilterRule.
type NamespacesGetIPFilterRuleResult struct {
	IPFilterRule
}

// NamespacesGetNetworkRuleSetResponse contains the response from method Namespaces.GetNetworkRuleSet.
type NamespacesGetNetworkRuleSetResponse struct {
	NamespacesGetNetworkRuleSetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesGetNetworkRuleSetResult contains the result from method Namespaces.GetNetworkRuleSet.
type NamespacesGetNetworkRuleSetResult struct {
	NetworkRuleSet
}

// NamespacesGetResponse contains the response from method Namespaces.Get.
type NamespacesGetResponse struct {
	NamespacesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesGetResult contains the result from method Namespaces.Get.
type NamespacesGetResult struct {
	SBNamespace
}

// NamespacesGetVirtualNetworkRuleResponse contains the response from method Namespaces.GetVirtualNetworkRule.
type NamespacesGetVirtualNetworkRuleResponse struct {
	NamespacesGetVirtualNetworkRuleResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesGetVirtualNetworkRuleResult contains the result from method Namespaces.GetVirtualNetworkRule.
type NamespacesGetVirtualNetworkRuleResult struct {
	VirtualNetworkRule
}

// NamespacesListAuthorizationRulesResponse contains the response from method Namespaces.ListAuthorizationRules.
type NamespacesListAuthorizationRulesResponse struct {
	NamespacesListAuthorizationRulesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesListAuthorizationRulesResult contains the result from method Namespaces.ListAuthorizationRules.
type NamespacesListAuthorizationRulesResult struct {
	SBAuthorizationRuleListResult
}

// NamespacesListByResourceGroupResponse contains the response from method Namespaces.ListByResourceGroup.
type NamespacesListByResourceGroupResponse struct {
	NamespacesListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesListByResourceGroupResult contains the result from method Namespaces.ListByResourceGroup.
type NamespacesListByResourceGroupResult struct {
	SBNamespaceListResult
}

// NamespacesListIPFilterRulesResponse contains the response from method Namespaces.ListIPFilterRules.
type NamespacesListIPFilterRulesResponse struct {
	NamespacesListIPFilterRulesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesListIPFilterRulesResult contains the result from method Namespaces.ListIPFilterRules.
type NamespacesListIPFilterRulesResult struct {
	IPFilterRuleListResult
}

// NamespacesListKeysResponse contains the response from method Namespaces.ListKeys.
type NamespacesListKeysResponse struct {
	NamespacesListKeysResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesListKeysResult contains the result from method Namespaces.ListKeys.
type NamespacesListKeysResult struct {
	AccessKeys
}

// NamespacesListNetworkRuleSetsResponse contains the response from method Namespaces.ListNetworkRuleSets.
type NamespacesListNetworkRuleSetsResponse struct {
	NamespacesListNetworkRuleSetsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesListNetworkRuleSetsResult contains the result from method Namespaces.ListNetworkRuleSets.
type NamespacesListNetworkRuleSetsResult struct {
	NetworkRuleSetListResult
}

// NamespacesListResponse contains the response from method Namespaces.List.
type NamespacesListResponse struct {
	NamespacesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesListResult contains the result from method Namespaces.List.
type NamespacesListResult struct {
	SBNamespaceListResult
}

// NamespacesListVirtualNetworkRulesResponse contains the response from method Namespaces.ListVirtualNetworkRules.
type NamespacesListVirtualNetworkRulesResponse struct {
	NamespacesListVirtualNetworkRulesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesListVirtualNetworkRulesResult contains the result from method Namespaces.ListVirtualNetworkRules.
type NamespacesListVirtualNetworkRulesResult struct {
	VirtualNetworkRuleListResult
}

// NamespacesMigrateResponse contains the response from method Namespaces.Migrate.
type NamespacesMigrateResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesRegenerateKeysResponse contains the response from method Namespaces.RegenerateKeys.
type NamespacesRegenerateKeysResponse struct {
	NamespacesRegenerateKeysResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesRegenerateKeysResult contains the result from method Namespaces.RegenerateKeys.
type NamespacesRegenerateKeysResult struct {
	AccessKeys
}

// NamespacesUpdateResponse contains the response from method Namespaces.Update.
type NamespacesUpdateResponse struct {
	NamespacesUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesUpdateResult contains the result from method Namespaces.Update.
type NamespacesUpdateResult struct {
	SBNamespace
}

// OperationsListResponse contains the response from method Operations.List.
type OperationsListResponse struct {
	OperationsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsListResult contains the result from method Operations.List.
type OperationsListResult struct {
	OperationListResult
}

// PremiumMessagingRegionsListResponse contains the response from method PremiumMessagingRegions.List.
type PremiumMessagingRegionsListResponse struct {
	PremiumMessagingRegionsListResultEnvelope
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PremiumMessagingRegionsListResultEnvelope contains the result from method PremiumMessagingRegions.List.
type PremiumMessagingRegionsListResultEnvelope struct {
	PremiumMessagingRegionsListResult
}

// PrivateEndpointConnectionsCreateOrUpdateResponse contains the response from method PrivateEndpointConnections.CreateOrUpdate.
type PrivateEndpointConnectionsCreateOrUpdateResponse struct {
	PrivateEndpointConnectionsCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsCreateOrUpdateResult contains the result from method PrivateEndpointConnections.CreateOrUpdate.
type PrivateEndpointConnectionsCreateOrUpdateResult struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsDeletePollerResponse contains the response from method PrivateEndpointConnections.Delete.
type PrivateEndpointConnectionsDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *PrivateEndpointConnectionsDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l PrivateEndpointConnectionsDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (PrivateEndpointConnectionsDeleteResponse, error) {
	respType := PrivateEndpointConnectionsDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a PrivateEndpointConnectionsDeletePollerResponse from the provided client and resume token.
func (l *PrivateEndpointConnectionsDeletePollerResponse) Resume(ctx context.Context, client *PrivateEndpointConnectionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("PrivateEndpointConnectionsClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &PrivateEndpointConnectionsDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// PrivateEndpointConnectionsDeleteResponse contains the response from method PrivateEndpointConnections.Delete.
type PrivateEndpointConnectionsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsGetResponse contains the response from method PrivateEndpointConnections.Get.
type PrivateEndpointConnectionsGetResponse struct {
	PrivateEndpointConnectionsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsGetResult contains the result from method PrivateEndpointConnections.Get.
type PrivateEndpointConnectionsGetResult struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsListResponse contains the response from method PrivateEndpointConnections.List.
type PrivateEndpointConnectionsListResponse struct {
	PrivateEndpointConnectionsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsListResult contains the result from method PrivateEndpointConnections.List.
type PrivateEndpointConnectionsListResult struct {
	PrivateEndpointConnectionListResult
}

// PrivateLinkResourcesGetResponse contains the response from method PrivateLinkResources.Get.
type PrivateLinkResourcesGetResponse struct {
	PrivateLinkResourcesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateLinkResourcesGetResult contains the result from method PrivateLinkResources.Get.
type PrivateLinkResourcesGetResult struct {
	PrivateLinkResourcesListResult
}

// QueuesCreateOrUpdateAuthorizationRuleResponse contains the response from method Queues.CreateOrUpdateAuthorizationRule.
type QueuesCreateOrUpdateAuthorizationRuleResponse struct {
	QueuesCreateOrUpdateAuthorizationRuleResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// QueuesCreateOrUpdateAuthorizationRuleResult contains the result from method Queues.CreateOrUpdateAuthorizationRule.
type QueuesCreateOrUpdateAuthorizationRuleResult struct {
	SBAuthorizationRule
}

// QueuesCreateOrUpdateResponse contains the response from method Queues.CreateOrUpdate.
type QueuesCreateOrUpdateResponse struct {
	QueuesCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// QueuesCreateOrUpdateResult contains the result from method Queues.CreateOrUpdate.
type QueuesCreateOrUpdateResult struct {
	SBQueue
}

// QueuesDeleteAuthorizationRuleResponse contains the response from method Queues.DeleteAuthorizationRule.
type QueuesDeleteAuthorizationRuleResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// QueuesDeleteResponse contains the response from method Queues.Delete.
type QueuesDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// QueuesGetAuthorizationRuleResponse contains the response from method Queues.GetAuthorizationRule.
type QueuesGetAuthorizationRuleResponse struct {
	QueuesGetAuthorizationRuleResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// QueuesGetAuthorizationRuleResult contains the result from method Queues.GetAuthorizationRule.
type QueuesGetAuthorizationRuleResult struct {
	SBAuthorizationRule
}

// QueuesGetResponse contains the response from method Queues.Get.
type QueuesGetResponse struct {
	QueuesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// QueuesGetResult contains the result from method Queues.Get.
type QueuesGetResult struct {
	SBQueue
}

// QueuesListAuthorizationRulesResponse contains the response from method Queues.ListAuthorizationRules.
type QueuesListAuthorizationRulesResponse struct {
	QueuesListAuthorizationRulesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// QueuesListAuthorizationRulesResult contains the result from method Queues.ListAuthorizationRules.
type QueuesListAuthorizationRulesResult struct {
	SBAuthorizationRuleListResult
}

// QueuesListByNamespaceResponse contains the response from method Queues.ListByNamespace.
type QueuesListByNamespaceResponse struct {
	QueuesListByNamespaceResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// QueuesListByNamespaceResult contains the result from method Queues.ListByNamespace.
type QueuesListByNamespaceResult struct {
	SBQueueListResult
}

// QueuesListKeysResponse contains the response from method Queues.ListKeys.
type QueuesListKeysResponse struct {
	QueuesListKeysResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// QueuesListKeysResult contains the result from method Queues.ListKeys.
type QueuesListKeysResult struct {
	AccessKeys
}

// QueuesRegenerateKeysResponse contains the response from method Queues.RegenerateKeys.
type QueuesRegenerateKeysResponse struct {
	QueuesRegenerateKeysResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// QueuesRegenerateKeysResult contains the result from method Queues.RegenerateKeys.
type QueuesRegenerateKeysResult struct {
	AccessKeys
}

// RegionsListBySKUResponse contains the response from method Regions.ListBySKU.
type RegionsListBySKUResponse struct {
	RegionsListBySKUResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RegionsListBySKUResult contains the result from method Regions.ListBySKU.
type RegionsListBySKUResult struct {
	PremiumMessagingRegionsListResult
}

// RulesCreateOrUpdateResponse contains the response from method Rules.CreateOrUpdate.
type RulesCreateOrUpdateResponse struct {
	RulesCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RulesCreateOrUpdateResult contains the result from method Rules.CreateOrUpdate.
type RulesCreateOrUpdateResult struct {
	Rule
}

// RulesDeleteResponse contains the response from method Rules.Delete.
type RulesDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RulesGetResponse contains the response from method Rules.Get.
type RulesGetResponse struct {
	RulesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RulesGetResult contains the result from method Rules.Get.
type RulesGetResult struct {
	Rule
}

// RulesListBySubscriptionsResponse contains the response from method Rules.ListBySubscriptions.
type RulesListBySubscriptionsResponse struct {
	RulesListBySubscriptionsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RulesListBySubscriptionsResult contains the result from method Rules.ListBySubscriptions.
type RulesListBySubscriptionsResult struct {
	RuleListResult
}

// SubscriptionsCreateOrUpdateResponse contains the response from method Subscriptions.CreateOrUpdate.
type SubscriptionsCreateOrUpdateResponse struct {
	SubscriptionsCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionsCreateOrUpdateResult contains the result from method Subscriptions.CreateOrUpdate.
type SubscriptionsCreateOrUpdateResult struct {
	SBSubscription
}

// SubscriptionsDeleteResponse contains the response from method Subscriptions.Delete.
type SubscriptionsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionsGetResponse contains the response from method Subscriptions.Get.
type SubscriptionsGetResponse struct {
	SubscriptionsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionsGetResult contains the result from method Subscriptions.Get.
type SubscriptionsGetResult struct {
	SBSubscription
}

// SubscriptionsListByTopicResponse contains the response from method Subscriptions.ListByTopic.
type SubscriptionsListByTopicResponse struct {
	SubscriptionsListByTopicResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionsListByTopicResult contains the result from method Subscriptions.ListByTopic.
type SubscriptionsListByTopicResult struct {
	SBSubscriptionListResult
}

// TopicsCreateOrUpdateAuthorizationRuleResponse contains the response from method Topics.CreateOrUpdateAuthorizationRule.
type TopicsCreateOrUpdateAuthorizationRuleResponse struct {
	TopicsCreateOrUpdateAuthorizationRuleResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TopicsCreateOrUpdateAuthorizationRuleResult contains the result from method Topics.CreateOrUpdateAuthorizationRule.
type TopicsCreateOrUpdateAuthorizationRuleResult struct {
	SBAuthorizationRule
}

// TopicsCreateOrUpdateResponse contains the response from method Topics.CreateOrUpdate.
type TopicsCreateOrUpdateResponse struct {
	TopicsCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TopicsCreateOrUpdateResult contains the result from method Topics.CreateOrUpdate.
type TopicsCreateOrUpdateResult struct {
	SBTopic
}

// TopicsDeleteAuthorizationRuleResponse contains the response from method Topics.DeleteAuthorizationRule.
type TopicsDeleteAuthorizationRuleResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TopicsDeleteResponse contains the response from method Topics.Delete.
type TopicsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TopicsGetAuthorizationRuleResponse contains the response from method Topics.GetAuthorizationRule.
type TopicsGetAuthorizationRuleResponse struct {
	TopicsGetAuthorizationRuleResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TopicsGetAuthorizationRuleResult contains the result from method Topics.GetAuthorizationRule.
type TopicsGetAuthorizationRuleResult struct {
	SBAuthorizationRule
}

// TopicsGetResponse contains the response from method Topics.Get.
type TopicsGetResponse struct {
	TopicsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TopicsGetResult contains the result from method Topics.Get.
type TopicsGetResult struct {
	SBTopic
}

// TopicsListAuthorizationRulesResponse contains the response from method Topics.ListAuthorizationRules.
type TopicsListAuthorizationRulesResponse struct {
	TopicsListAuthorizationRulesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TopicsListAuthorizationRulesResult contains the result from method Topics.ListAuthorizationRules.
type TopicsListAuthorizationRulesResult struct {
	SBAuthorizationRuleListResult
}

// TopicsListByNamespaceResponse contains the response from method Topics.ListByNamespace.
type TopicsListByNamespaceResponse struct {
	TopicsListByNamespaceResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TopicsListByNamespaceResult contains the result from method Topics.ListByNamespace.
type TopicsListByNamespaceResult struct {
	SBTopicListResult
}

// TopicsListKeysResponse contains the response from method Topics.ListKeys.
type TopicsListKeysResponse struct {
	TopicsListKeysResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TopicsListKeysResult contains the result from method Topics.ListKeys.
type TopicsListKeysResult struct {
	AccessKeys
}

// TopicsRegenerateKeysResponse contains the response from method Topics.RegenerateKeys.
type TopicsRegenerateKeysResponse struct {
	TopicsRegenerateKeysResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TopicsRegenerateKeysResult contains the result from method Topics.RegenerateKeys.
type TopicsRegenerateKeysResult struct {
	AccessKeys
}
