// +build go1.9

// Copyright 2017 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder
// commit ID: 714052a3359963ba46703fe033cf9dd4838bea11

package subscriptions

import original "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2016-06-01/subscriptions"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type GroupClient = original.GroupClient
type SpendingLimit = original.SpendingLimit

const (
	CurrentPeriodOff SpendingLimit = original.CurrentPeriodOff
	Off              SpendingLimit = original.Off
	On               SpendingLimit = original.On
)

type State = original.State

const (
	Deleted  State = original.Deleted
	Disabled State = original.Disabled
	Enabled  State = original.Enabled
	PastDue  State = original.PastDue
	Warned   State = original.Warned
)

type ListResult = original.ListResult
type Location = original.Location
type LocationListResult = original.LocationListResult
type Policies = original.Policies
type Subscription = original.Subscription
type TenantIDDescription = original.TenantIDDescription
type TenantListResult = original.TenantListResult
type TenantsClient = original.TenantsClient

func New() ManagementClient {
	return original.New()
}
func NewWithBaseURI(baseURI string) ManagementClient {
	return original.NewWithBaseURI(baseURI)
}
func NewGroupClient() GroupClient {
	return original.NewGroupClient()
}
func NewGroupClientWithBaseURI(baseURI string) GroupClient {
	return original.NewGroupClientWithBaseURI(baseURI)
}
func NewTenantsClient() TenantsClient {
	return original.NewTenantsClient()
}
func NewTenantsClientWithBaseURI(baseURI string) TenantsClient {
	return original.NewTenantsClientWithBaseURI(baseURI)
}
func UserAgent() string {
	return original.UserAgent()
}
func Version() string {
	return original.Version()
}
