//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfeatures

// ClientGetResponse contains the response from method Client.Get.
type ClientGetResponse struct {
	// Previewed feature information.
	FeatureResult
}

// ClientListAllResponse contains the response from method Client.NewListAllPager.
type ClientListAllResponse struct {
	// List of previewed features.
	FeatureOperationsListResult
}

// ClientListResponse contains the response from method Client.NewListPager.
type ClientListResponse struct {
	// List of previewed features.
	FeatureOperationsListResult
}

// ClientRegisterResponse contains the response from method Client.Register.
type ClientRegisterResponse struct {
	// Previewed feature information.
	FeatureResult
}

// ClientUnregisterResponse contains the response from method Client.Unregister.
type ClientUnregisterResponse struct {
	// Previewed feature information.
	FeatureResult
}

// FeatureClientListOperationsResponse contains the response from method FeatureClient.NewListOperationsPager.
type FeatureClientListOperationsResponse struct {
	// Result of the request to list Microsoft.Features operations. It contains a list of operations and a URL link to get the
	// next set of results.
	OperationListResult
}

// SubscriptionFeatureRegistrationsClientCreateOrUpdateResponse contains the response from method SubscriptionFeatureRegistrationsClient.CreateOrUpdate.
type SubscriptionFeatureRegistrationsClientCreateOrUpdateResponse struct {
	// Subscription feature registration details
	SubscriptionFeatureRegistration
}

// SubscriptionFeatureRegistrationsClientDeleteResponse contains the response from method SubscriptionFeatureRegistrationsClient.Delete.
type SubscriptionFeatureRegistrationsClientDeleteResponse struct {
	// placeholder for future response values
}

// SubscriptionFeatureRegistrationsClientGetResponse contains the response from method SubscriptionFeatureRegistrationsClient.Get.
type SubscriptionFeatureRegistrationsClientGetResponse struct {
	// Subscription feature registration details
	SubscriptionFeatureRegistration
}

// SubscriptionFeatureRegistrationsClientListAllBySubscriptionResponse contains the response from method SubscriptionFeatureRegistrationsClient.NewListAllBySubscriptionPager.
type SubscriptionFeatureRegistrationsClientListAllBySubscriptionResponse struct {
	// The list of subscription feature registrations.
	SubscriptionFeatureRegistrationList
}

// SubscriptionFeatureRegistrationsClientListBySubscriptionResponse contains the response from method SubscriptionFeatureRegistrationsClient.NewListBySubscriptionPager.
type SubscriptionFeatureRegistrationsClientListBySubscriptionResponse struct {
	// The list of subscription feature registrations.
	SubscriptionFeatureRegistrationList
}
