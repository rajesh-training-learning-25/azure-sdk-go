// Package features implements the Azure ARM Features service API version 2015-12-01.
//
// Azure Feature Exposure Control (AFEC) provides a mechanism for the resource providers to control feature exposure to
// users. Resource providers typically use this mechanism to provide public/private preview for new features prior to
// making them generally available. Users need to explicitly register for AFEC features to get access to such
// functionality.
package features

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

const (
	// DefaultBaseURI is the default URI used for the service Features
	DefaultBaseURI = "https://management.azure.com"
)

// ManagementClient is the base client for Features.
type ManagementClient struct {
	autorest.Client
	BaseURI        string
	SubscriptionID string
}

// New creates an instance of the ManagementClient client.
func New(subscriptionID string) ManagementClient {
	return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWithBaseURI creates an instance of the ManagementClient client.
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return ManagementClient{
		Client:         autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:        baseURI,
		SubscriptionID: subscriptionID,
	}
}

// NewWithAuthFile creates an instance of the ManagementClient client.
func NewWithAuthFile() (c ManagementClient, err error) {
	authentication, err := auth.GetClientSetup(DefaultBaseURI)
	if err != nil {
		return
	}
	c = NewWithBaseURI(authentication.BaseURI, authentication.SubscriptionID)
	c.Authorizer = authentication
	return
}
