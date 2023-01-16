//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armloadtesting

// LoadTestsClientCreateOrUpdateResponse contains the response from method LoadTestsClient.CreateOrUpdate.
type LoadTestsClientCreateOrUpdateResponse struct {
	LoadTestResource
}

// LoadTestsClientDeleteResponse contains the response from method LoadTestsClient.Delete.
type LoadTestsClientDeleteResponse struct {
	// placeholder for future response values
}

// LoadTestsClientGetResponse contains the response from method LoadTestsClient.Get.
type LoadTestsClientGetResponse struct {
	LoadTestResource
}

// LoadTestsClientListByResourceGroupResponse contains the response from method LoadTestsClient.ListByResourceGroup.
type LoadTestsClientListByResourceGroupResponse struct {
	LoadTestResourcePageList
}

// LoadTestsClientListBySubscriptionResponse contains the response from method LoadTestsClient.ListBySubscription.
type LoadTestsClientListBySubscriptionResponse struct {
	LoadTestResourcePageList
}

// LoadTestsClientListOutboundNetworkDependenciesEndpointsResponse contains the response from method LoadTestsClient.ListOutboundNetworkDependenciesEndpoints.
type LoadTestsClientListOutboundNetworkDependenciesEndpointsResponse struct {
	OutboundEnvironmentEndpointCollection
}

// LoadTestsClientUpdateResponse contains the response from method LoadTestsClient.Update.
type LoadTestsClientUpdateResponse struct {
	LoadTestResource
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}

// QuotasClientCheckAvailabilityResponse contains the response from method QuotasClient.CheckAvailability.
type QuotasClientCheckAvailabilityResponse struct {
	CheckQuotaAvailabilityResponse
}

// QuotasClientGetResponse contains the response from method QuotasClient.Get.
type QuotasClientGetResponse struct {
	QuotaResource
}

// QuotasClientListResponse contains the response from method QuotasClient.List.
type QuotasClientListResponse struct {
	QuotaResourceList
}
