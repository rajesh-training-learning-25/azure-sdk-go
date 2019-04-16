// +build go1.9

// Copyright 2019 Microsoft Corporation
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

package hanaonazure

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/hanaonazure/mgmt/2017-11-03-preview/hanaonazure"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type HanaDatabaseContainersEnum = original.HanaDatabaseContainersEnum

const (
	Multiple HanaDatabaseContainersEnum = original.Multiple
	Single   HanaDatabaseContainersEnum = original.Single
)

type HanaHardwareTypeNamesEnum = original.HanaHardwareTypeNamesEnum

const (
	CiscoUCS HanaHardwareTypeNamesEnum = original.CiscoUCS
	HPE      HanaHardwareTypeNamesEnum = original.HPE
)

type HanaInstancePowerStateEnum = original.HanaInstancePowerStateEnum

const (
	Restarting HanaInstancePowerStateEnum = original.Restarting
	Started    HanaInstancePowerStateEnum = original.Started
	Starting   HanaInstancePowerStateEnum = original.Starting
	Stopped    HanaInstancePowerStateEnum = original.Stopped
	Stopping   HanaInstancePowerStateEnum = original.Stopping
	Unknown    HanaInstancePowerStateEnum = original.Unknown
)

type HanaInstanceSizeNamesEnum = original.HanaInstanceSizeNamesEnum

const (
	S144    HanaInstanceSizeNamesEnum = original.S144
	S144m   HanaInstanceSizeNamesEnum = original.S144m
	S192    HanaInstanceSizeNamesEnum = original.S192
	S192m   HanaInstanceSizeNamesEnum = original.S192m
	S192xm  HanaInstanceSizeNamesEnum = original.S192xm
	S384    HanaInstanceSizeNamesEnum = original.S384
	S384m   HanaInstanceSizeNamesEnum = original.S384m
	S384xm  HanaInstanceSizeNamesEnum = original.S384xm
	S384xxm HanaInstanceSizeNamesEnum = original.S384xxm
	S576m   HanaInstanceSizeNamesEnum = original.S576m
	S576xm  HanaInstanceSizeNamesEnum = original.S576xm
	S72     HanaInstanceSizeNamesEnum = original.S72
	S72m    HanaInstanceSizeNamesEnum = original.S72m
	S768    HanaInstanceSizeNamesEnum = original.S768
	S768m   HanaInstanceSizeNamesEnum = original.S768m
	S768xm  HanaInstanceSizeNamesEnum = original.S768xm
	S96     HanaInstanceSizeNamesEnum = original.S96
	S960m   HanaInstanceSizeNamesEnum = original.S960m
)

type BaseClient = original.BaseClient
type Disk = original.Disk
type Display = original.Display
type ErrorResponse = original.ErrorResponse
type HanaInstance = original.HanaInstance
type HanaInstanceProperties = original.HanaInstanceProperties
type HanaInstancesClient = original.HanaInstancesClient
type HanaInstancesEnableMonitoringFuture = original.HanaInstancesEnableMonitoringFuture
type HanaInstancesListResult = original.HanaInstancesListResult
type HanaInstancesListResultIterator = original.HanaInstancesListResultIterator
type HanaInstancesListResultPage = original.HanaInstancesListResultPage
type HanaInstancesRestartFuture = original.HanaInstancesRestartFuture
type HardwareProfile = original.HardwareProfile
type IPAddress = original.IPAddress
type MonitoringDetails = original.MonitoringDetails
type NetworkProfile = original.NetworkProfile
type OSProfile = original.OSProfile
type Operation = original.Operation
type OperationList = original.OperationList
type OperationsClient = original.OperationsClient
type Resource = original.Resource
type StorageProfile = original.StorageProfile
type Tags = original.Tags

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewHanaInstancesClient(subscriptionID string) HanaInstancesClient {
	return original.NewHanaInstancesClient(subscriptionID)
}
func NewHanaInstancesClientWithBaseURI(baseURI string, subscriptionID string) HanaInstancesClient {
	return original.NewHanaInstancesClientWithBaseURI(baseURI, subscriptionID)
}
func NewHanaInstancesListResultIterator(page HanaInstancesListResultPage) HanaInstancesListResultIterator {
	return original.NewHanaInstancesListResultIterator(page)
}
func NewHanaInstancesListResultPage(getNextPage func(context.Context, HanaInstancesListResult) (HanaInstancesListResult, error)) HanaInstancesListResultPage {
	return original.NewHanaInstancesListResultPage(getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleHanaDatabaseContainersEnumValues() []HanaDatabaseContainersEnum {
	return original.PossibleHanaDatabaseContainersEnumValues()
}
func PossibleHanaHardwareTypeNamesEnumValues() []HanaHardwareTypeNamesEnum {
	return original.PossibleHanaHardwareTypeNamesEnumValues()
}
func PossibleHanaInstancePowerStateEnumValues() []HanaInstancePowerStateEnum {
	return original.PossibleHanaInstancePowerStateEnumValues()
}
func PossibleHanaInstanceSizeNamesEnumValues() []HanaInstanceSizeNamesEnum {
	return original.PossibleHanaInstanceSizeNamesEnumValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
