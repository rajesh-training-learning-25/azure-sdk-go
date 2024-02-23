//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armworkloadssapvirtualinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadssapvirtualinstance/armworkloadssapvirtualinstance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1e318cbfd2e239db54c80af5e6aea7fdf658851/specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/examples/sapcentralinstances/SAPCentralInstances_Get.json
func ExampleSAPCentralInstancesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadssapvirtualinstance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSAPCentralInstancesClient().Get(ctx, "test-rg", "X00", "centralServer", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SAPCentralServerInstance = armworkloadssapvirtualinstance.SAPCentralServerInstance{
	// 	Name: to.Ptr("centralServer"),
	// 	Type: to.Ptr("Microsoft.Workloads/sapVirtualInstances/centralInstances"),
	// 	ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Workloads/sapVirtualInstances/X00/centralInstances/centralServer"),
	// 	SystemData: &armworkloadssapvirtualinstance.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@xyz.com"),
	// 		CreatedByType: to.Ptr(armworkloadssapvirtualinstance.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@xyz.com"),
	// 		LastModifiedByType: to.Ptr(armworkloadssapvirtualinstance.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armworkloadssapvirtualinstance.SAPCentralServerProperties{
	// 		EnqueueReplicationServerProperties: &armworkloadssapvirtualinstance.EnqueueReplicationServerProperties{
	// 			ErsVersion: to.Ptr(armworkloadssapvirtualinstance.EnqueueReplicationServerTypeEnqueueReplicator1),
	// 			Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateHealthy),
	// 			Hostname: to.Ptr("vh-ers1"),
	// 			InstanceNo: to.Ptr("00"),
	// 			IPAddress: to.Ptr("10.0.0.5"),
	// 			KernelPatch: to.Ptr("patch 300"),
	// 			KernelVersion: to.Ptr("777"),
	// 		},
	// 		EnqueueServerProperties: &armworkloadssapvirtualinstance.EnqueueServerProperties{
	// 			Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateUnknown),
	// 			Hostname: to.Ptr("vh-ascs1"),
	// 			IPAddress: to.Ptr("10.0.0.5"),
	// 			Port: to.Ptr[int64](3600),
	// 		},
	// 		GatewayServerProperties: &armworkloadssapvirtualinstance.GatewayServerProperties{
	// 			Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateDegraded),
	// 			Port: to.Ptr[int64](3300),
	// 		},
	// 		Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateUnknown),
	// 		InstanceNo: to.Ptr("00"),
	// 		KernelPatch: to.Ptr("patch 300"),
	// 		KernelVersion: to.Ptr("777"),
	// 		MessageServerProperties: &armworkloadssapvirtualinstance.MessageServerProperties{
	// 			Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateUnhealthy),
	// 			Hostname: to.Ptr("vh-ascs1"),
	// 			HTTPPort: to.Ptr[int64](8100),
	// 			HTTPSPort: to.Ptr[int64](44400),
	// 			InternalMsPort: to.Ptr[int64](3900),
	// 			IPAddress: to.Ptr("10.0.0.5"),
	// 			MSPort: to.Ptr[int64](3600),
	// 		},
	// 		ProvisioningState: to.Ptr(armworkloadssapvirtualinstance.SapVirtualInstanceProvisioningStateSucceeded),
	// 		Status: to.Ptr(armworkloadssapvirtualinstance.SAPVirtualInstanceStatusRunning),
	// 		Subnet: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/SAP-E2ETest-rg/providers/Microsoft.Network/virtualNetworks/loop-test-vnet/subnets/loopsubnet"),
	// 		VMDetails: []*armworkloadssapvirtualinstance.CentralServerVMDetails{
	// 			{
	// 				Type: to.Ptr(armworkloadssapvirtualinstance.CentralServerVirtualMachineTypePrimary),
	// 				VirtualMachineID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Compute/virtualMachines/cs-vm"),
	// 		}},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1e318cbfd2e239db54c80af5e6aea7fdf658851/specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/examples/sapcentralinstances/SAPCentralInstances_Create_HA_AvSet.json
func ExampleSAPCentralInstancesClient_BeginCreate_createSapCentralInstancesForHaSystemWithAvailabilitySet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadssapvirtualinstance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSAPCentralInstancesClient().BeginCreate(ctx, "test-rg", "X00", "centralServer", armworkloadssapvirtualinstance.SAPCentralServerInstance{
		Location:   to.Ptr("eastus"),
		Tags:       map[string]*string{},
		Properties: &armworkloadssapvirtualinstance.SAPCentralServerProperties{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SAPCentralServerInstance = armworkloadssapvirtualinstance.SAPCentralServerInstance{
	// 	Name: to.Ptr("centralServer"),
	// 	Type: to.Ptr("Microsoft.Workloads/sapVirtualInstances/centralInstances"),
	// 	ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Workloads/sapVirtualInstances/X00/centralInstances/centralServer"),
	// 	SystemData: &armworkloadssapvirtualinstance.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@xyz.com"),
	// 		CreatedByType: to.Ptr(armworkloadssapvirtualinstance.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@xyz.com"),
	// 		LastModifiedByType: to.Ptr(armworkloadssapvirtualinstance.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armworkloadssapvirtualinstance.SAPCentralServerProperties{
	// 		EnqueueReplicationServerProperties: &armworkloadssapvirtualinstance.EnqueueReplicationServerProperties{
	// 			ErsVersion: to.Ptr(armworkloadssapvirtualinstance.EnqueueReplicationServerTypeEnqueueReplicator1),
	// 			Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateHealthy),
	// 			Hostname: to.Ptr("vh-ers1"),
	// 			InstanceNo: to.Ptr("00"),
	// 			IPAddress: to.Ptr("10.0.0.5"),
	// 			KernelPatch: to.Ptr("patch 300"),
	// 			KernelVersion: to.Ptr("777"),
	// 		},
	// 		EnqueueServerProperties: &armworkloadssapvirtualinstance.EnqueueServerProperties{
	// 			Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateHealthy),
	// 			Hostname: to.Ptr("vh-ascs1"),
	// 			IPAddress: to.Ptr("10.0.0.5"),
	// 			Port: to.Ptr[int64](3600),
	// 		},
	// 		GatewayServerProperties: &armworkloadssapvirtualinstance.GatewayServerProperties{
	// 			Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateHealthy),
	// 			Port: to.Ptr[int64](3300),
	// 		},
	// 		Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateHealthy),
	// 		InstanceNo: to.Ptr("00"),
	// 		KernelPatch: to.Ptr("patch 300"),
	// 		KernelVersion: to.Ptr("777"),
	// 		LoadBalancerDetails: &armworkloadssapvirtualinstance.LoadBalancerDetails{
	// 			ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Network/loadBalancers/cs-ASCS-loadBalancer"),
	// 		},
	// 		MessageServerProperties: &armworkloadssapvirtualinstance.MessageServerProperties{
	// 			Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateHealthy),
	// 			Hostname: to.Ptr("vh-ascs1"),
	// 			HTTPPort: to.Ptr[int64](8100),
	// 			HTTPSPort: to.Ptr[int64](44400),
	// 			InternalMsPort: to.Ptr[int64](3900),
	// 			IPAddress: to.Ptr("10.0.0.5"),
	// 			MSPort: to.Ptr[int64](3600),
	// 		},
	// 		ProvisioningState: to.Ptr(armworkloadssapvirtualinstance.SapVirtualInstanceProvisioningStateSucceeded),
	// 		Status: to.Ptr(armworkloadssapvirtualinstance.SAPVirtualInstanceStatusRunning),
	// 		Subnet: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/SAP-E2ETest-rg/providers/Microsoft.Network/virtualNetworks/loop-test-vnet/subnets/loopsubnet"),
	// 		VMDetails: []*armworkloadssapvirtualinstance.CentralServerVMDetails{
	// 			{
	// 				Type: to.Ptr(armworkloadssapvirtualinstance.CentralServerVirtualMachineTypePrimary),
	// 				StorageDetails: []*armworkloadssapvirtualinstance.StorageInformation{
	// 					{
	// 						ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Storage/storageAccounts/nfsstorageaccount"),
	// 				}},
	// 				VirtualMachineID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Compute/virtualMachines/cs-vm1"),
	// 			},
	// 			{
	// 				Type: to.Ptr(armworkloadssapvirtualinstance.CentralServerVirtualMachineTypePrimary),
	// 				StorageDetails: []*armworkloadssapvirtualinstance.StorageInformation{
	// 					{
	// 						ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Storage/storageAccounts/nfsstorageaccount"),
	// 				}},
	// 				VirtualMachineID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Compute/virtualMachines/cs-vm2"),
	// 		}},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1e318cbfd2e239db54c80af5e6aea7fdf658851/specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/examples/sapcentralinstances/SAPCentralInstances_Create.json
func ExampleSAPCentralInstancesClient_BeginCreate_sapCentralInstancesCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadssapvirtualinstance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSAPCentralInstancesClient().BeginCreate(ctx, "test-rg", "X00", "centralServer", armworkloadssapvirtualinstance.SAPCentralServerInstance{
		Location:   to.Ptr("eastus"),
		Tags:       map[string]*string{},
		Properties: &armworkloadssapvirtualinstance.SAPCentralServerProperties{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SAPCentralServerInstance = armworkloadssapvirtualinstance.SAPCentralServerInstance{
	// 	Name: to.Ptr("centralServer"),
	// 	Type: to.Ptr("Microsoft.Workloads/sapVirtualInstances/centralInstances"),
	// 	ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Workloads/sapVirtualInstances/X00/centralInstances/centralServer"),
	// 	SystemData: &armworkloadssapvirtualinstance.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@xyz.com"),
	// 		CreatedByType: to.Ptr(armworkloadssapvirtualinstance.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@xyz.com"),
	// 		LastModifiedByType: to.Ptr(armworkloadssapvirtualinstance.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armworkloadssapvirtualinstance.SAPCentralServerProperties{
	// 		EnqueueReplicationServerProperties: &armworkloadssapvirtualinstance.EnqueueReplicationServerProperties{
	// 			ErsVersion: to.Ptr(armworkloadssapvirtualinstance.EnqueueReplicationServerTypeEnqueueReplicator1),
	// 			Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateHealthy),
	// 			Hostname: to.Ptr("vh-ers1"),
	// 			InstanceNo: to.Ptr("00"),
	// 			IPAddress: to.Ptr("10.0.0.5"),
	// 			KernelPatch: to.Ptr("patch 300"),
	// 			KernelVersion: to.Ptr("777"),
	// 		},
	// 		EnqueueServerProperties: &armworkloadssapvirtualinstance.EnqueueServerProperties{
	// 			Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateHealthy),
	// 			Hostname: to.Ptr("vh-ascs1"),
	// 			IPAddress: to.Ptr("10.0.0.5"),
	// 			Port: to.Ptr[int64](3600),
	// 		},
	// 		GatewayServerProperties: &armworkloadssapvirtualinstance.GatewayServerProperties{
	// 			Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateHealthy),
	// 			Port: to.Ptr[int64](3300),
	// 		},
	// 		Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateHealthy),
	// 		InstanceNo: to.Ptr("00"),
	// 		KernelPatch: to.Ptr("patch 300"),
	// 		KernelVersion: to.Ptr("777"),
	// 		MessageServerProperties: &armworkloadssapvirtualinstance.MessageServerProperties{
	// 			Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateHealthy),
	// 			Hostname: to.Ptr("vh-ascs1"),
	// 			HTTPPort: to.Ptr[int64](8100),
	// 			HTTPSPort: to.Ptr[int64](44400),
	// 			InternalMsPort: to.Ptr[int64](3900),
	// 			IPAddress: to.Ptr("10.0.0.5"),
	// 			MSPort: to.Ptr[int64](3600),
	// 		},
	// 		ProvisioningState: to.Ptr(armworkloadssapvirtualinstance.SapVirtualInstanceProvisioningStateSucceeded),
	// 		Status: to.Ptr(armworkloadssapvirtualinstance.SAPVirtualInstanceStatusRunning),
	// 		Subnet: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/SAP-E2ETest-rg/providers/Microsoft.Network/virtualNetworks/loop-test-vnet/subnets/loopsubnet"),
	// 		VMDetails: []*armworkloadssapvirtualinstance.CentralServerVMDetails{
	// 			{
	// 				Type: to.Ptr(armworkloadssapvirtualinstance.CentralServerVirtualMachineTypePrimary),
	// 				VirtualMachineID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Compute/virtualMachines/cs-vm"),
	// 		}},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1e318cbfd2e239db54c80af5e6aea7fdf658851/specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/examples/sapcentralinstances/SAPCentralInstances_Update.json
func ExampleSAPCentralInstancesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadssapvirtualinstance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSAPCentralInstancesClient().Update(ctx, "test-rg", "X00", "centralServer", armworkloadssapvirtualinstance.UpdateSAPCentralInstanceRequest{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SAPCentralServerInstance = armworkloadssapvirtualinstance.SAPCentralServerInstance{
	// 	Name: to.Ptr("centralServer"),
	// 	Type: to.Ptr("Microsoft.Workloads/sapVirtualInstances/centralInstances"),
	// 	ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Workloads/sapVirtualInstances/X00/centralInstances/centralServer"),
	// 	SystemData: &armworkloadssapvirtualinstance.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@xyz.com"),
	// 		CreatedByType: to.Ptr(armworkloadssapvirtualinstance.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@xyz.com"),
	// 		LastModifiedByType: to.Ptr(armworkloadssapvirtualinstance.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armworkloadssapvirtualinstance.SAPCentralServerProperties{
	// 		EnqueueReplicationServerProperties: &armworkloadssapvirtualinstance.EnqueueReplicationServerProperties{
	// 			ErsVersion: to.Ptr(armworkloadssapvirtualinstance.EnqueueReplicationServerTypeEnqueueReplicator1),
	// 			Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateHealthy),
	// 			Hostname: to.Ptr("vh-ers1"),
	// 			InstanceNo: to.Ptr("00"),
	// 			IPAddress: to.Ptr("10.0.0.5"),
	// 			KernelPatch: to.Ptr("patch 300"),
	// 			KernelVersion: to.Ptr("777"),
	// 		},
	// 		EnqueueServerProperties: &armworkloadssapvirtualinstance.EnqueueServerProperties{
	// 			Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateHealthy),
	// 			Hostname: to.Ptr("vh-ascs1"),
	// 			IPAddress: to.Ptr("10.0.0.5"),
	// 			Port: to.Ptr[int64](3600),
	// 		},
	// 		GatewayServerProperties: &armworkloadssapvirtualinstance.GatewayServerProperties{
	// 			Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateHealthy),
	// 			Port: to.Ptr[int64](3300),
	// 		},
	// 		Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateHealthy),
	// 		InstanceNo: to.Ptr("00"),
	// 		KernelPatch: to.Ptr("patch 300"),
	// 		KernelVersion: to.Ptr("777"),
	// 		MessageServerProperties: &armworkloadssapvirtualinstance.MessageServerProperties{
	// 			Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateHealthy),
	// 			Hostname: to.Ptr("vh-ascs1"),
	// 			HTTPPort: to.Ptr[int64](8100),
	// 			HTTPSPort: to.Ptr[int64](44400),
	// 			InternalMsPort: to.Ptr[int64](3900),
	// 			IPAddress: to.Ptr("10.0.0.5"),
	// 			MSPort: to.Ptr[int64](3600),
	// 		},
	// 		ProvisioningState: to.Ptr(armworkloadssapvirtualinstance.SapVirtualInstanceProvisioningStateSucceeded),
	// 		Status: to.Ptr(armworkloadssapvirtualinstance.SAPVirtualInstanceStatusRunning),
	// 		Subnet: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/SAP-E2ETest-rg/providers/Microsoft.Network/virtualNetworks/loop-test-vnet/subnets/loopsubnet"),
	// 		VMDetails: []*armworkloadssapvirtualinstance.CentralServerVMDetails{
	// 			{
	// 				Type: to.Ptr(armworkloadssapvirtualinstance.CentralServerVirtualMachineTypePrimary),
	// 				VirtualMachineID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Compute/virtualMachines/cs-vm"),
	// 		}},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1e318cbfd2e239db54c80af5e6aea7fdf658851/specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/examples/sapcentralinstances/SAPCentralInstances_Delete.json
func ExampleSAPCentralInstancesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadssapvirtualinstance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSAPCentralInstancesClient().BeginDelete(ctx, "test-rg", "X00", "centralServer", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1e318cbfd2e239db54c80af5e6aea7fdf658851/specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/examples/sapcentralinstances/SAPCentralInstances_List.json
func ExampleSAPCentralInstancesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadssapvirtualinstance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSAPCentralInstancesClient().NewListPager("test-rg", "X00", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.SAPCentralInstanceList = armworkloadssapvirtualinstance.SAPCentralInstanceList{
		// 	Value: []*armworkloadssapvirtualinstance.SAPCentralServerInstance{
		// 		{
		// 			Name: to.Ptr("centralServer"),
		// 			Type: to.Ptr("Microsoft.Workloads/sapVirtualInstances/centralInstances"),
		// 			ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Workloads/sapVirtualInstances/X00/centralInstances/centralServer"),
		// 			SystemData: &armworkloadssapvirtualinstance.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
		// 				CreatedBy: to.Ptr("user@xyz.com"),
		// 				CreatedByType: to.Ptr(armworkloadssapvirtualinstance.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user@xyz.com"),
		// 				LastModifiedByType: to.Ptr(armworkloadssapvirtualinstance.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armworkloadssapvirtualinstance.SAPCentralServerProperties{
		// 				EnqueueReplicationServerProperties: &armworkloadssapvirtualinstance.EnqueueReplicationServerProperties{
		// 					ErsVersion: to.Ptr(armworkloadssapvirtualinstance.EnqueueReplicationServerTypeEnqueueReplicator1),
		// 					Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateHealthy),
		// 					Hostname: to.Ptr("vh-ers1"),
		// 					InstanceNo: to.Ptr("00"),
		// 					IPAddress: to.Ptr("10.0.0.5"),
		// 					KernelPatch: to.Ptr("patch 300"),
		// 					KernelVersion: to.Ptr("777"),
		// 				},
		// 				EnqueueServerProperties: &armworkloadssapvirtualinstance.EnqueueServerProperties{
		// 					Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateHealthy),
		// 					Hostname: to.Ptr("vh-ascs1"),
		// 					IPAddress: to.Ptr("10.0.0.5"),
		// 					Port: to.Ptr[int64](3600),
		// 				},
		// 				GatewayServerProperties: &armworkloadssapvirtualinstance.GatewayServerProperties{
		// 					Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateHealthy),
		// 					Port: to.Ptr[int64](3300),
		// 				},
		// 				Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateHealthy),
		// 				InstanceNo: to.Ptr("00"),
		// 				KernelPatch: to.Ptr("patch 300"),
		// 				KernelVersion: to.Ptr("777"),
		// 				MessageServerProperties: &armworkloadssapvirtualinstance.MessageServerProperties{
		// 					Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateHealthy),
		// 					Hostname: to.Ptr("vh-ascs1"),
		// 					HTTPPort: to.Ptr[int64](8100),
		// 					HTTPSPort: to.Ptr[int64](44400),
		// 					InternalMsPort: to.Ptr[int64](3900),
		// 					IPAddress: to.Ptr("10.0.0.5"),
		// 					MSPort: to.Ptr[int64](3600),
		// 				},
		// 				ProvisioningState: to.Ptr(armworkloadssapvirtualinstance.SapVirtualInstanceProvisioningStateSucceeded),
		// 				Status: to.Ptr(armworkloadssapvirtualinstance.SAPVirtualInstanceStatusRunning),
		// 				Subnet: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/SAP-E2ETest-rg/providers/Microsoft.Network/virtualNetworks/loop-test-vnet/subnets/loopsubnet"),
		// 				VMDetails: []*armworkloadssapvirtualinstance.CentralServerVMDetails{
		// 					{
		// 						Type: to.Ptr(armworkloadssapvirtualinstance.CentralServerVirtualMachineTypePrimary),
		// 						VirtualMachineID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Compute/virtualMachines/cs-vm"),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1e318cbfd2e239db54c80af5e6aea7fdf658851/specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/examples/sapcentralinstances/SAPCentralInstances_StartInstance.json
func ExampleSAPCentralInstancesClient_BeginStartInstance_startTheSapCentralServicesInstance() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadssapvirtualinstance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSAPCentralInstancesClient().BeginStartInstance(ctx, "test-rg", "X00", "centralServer", &armworkloadssapvirtualinstance.SAPCentralInstancesClientBeginStartInstanceOptions{Body: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationStatusResult = armworkloadssapvirtualinstance.OperationStatusResult{
	// 	Name: to.Ptr("centralServer"),
	// 	EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T03:38:07.000Z"); return t}()),
	// 	ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Workloads/sapVirtualInstances/X00/centralInstances/centralServer"),
	// 	StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T03:36:07.000Z"); return t}()),
	// 	Status: to.Ptr("Succeeded"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1e318cbfd2e239db54c80af5e6aea7fdf658851/specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/examples/sapcentralinstances/SAPCentralInstances_StartInstance_WithInfraOperations.json
func ExampleSAPCentralInstancesClient_BeginStartInstance_startTheVirtualMachineSAndTheSapCentralServicesInstanceOnIt() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadssapvirtualinstance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSAPCentralInstancesClient().BeginStartInstance(ctx, "test-rg", "X00", "centralServer", &armworkloadssapvirtualinstance.SAPCentralInstancesClientBeginStartInstanceOptions{Body: &armworkloadssapvirtualinstance.StartRequest{
		StartVM: to.Ptr(true),
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationStatusResult = armworkloadssapvirtualinstance.OperationStatusResult{
	// 	Name: to.Ptr("centralServer"),
	// 	EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T03:38:07.000Z"); return t}()),
	// 	ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Workloads/sapVirtualInstances/X00/centralInstances/centralServer"),
	// 	StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T03:36:07.000Z"); return t}()),
	// 	Status: to.Ptr("Succeeded"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1e318cbfd2e239db54c80af5e6aea7fdf658851/specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/examples/sapcentralinstances/SAPCentralInstances_StopInstance_WithInfraOperations.json
func ExampleSAPCentralInstancesClient_BeginStopInstance_stopTheSapCentralServicesInstanceAndItsUnderlyingVirtualMachineS() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadssapvirtualinstance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSAPCentralInstancesClient().BeginStopInstance(ctx, "test-rg", "X00", "centralServer", &armworkloadssapvirtualinstance.SAPCentralInstancesClientBeginStopInstanceOptions{Body: &armworkloadssapvirtualinstance.StopRequest{
		DeallocateVM: to.Ptr(true),
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationStatusResult = armworkloadssapvirtualinstance.OperationStatusResult{
	// 	Name: to.Ptr("centralServer"),
	// 	EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T03:38:07.000Z"); return t}()),
	// 	ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Workloads/sapVirtualInstances/X00/centralInstances/centralServer"),
	// 	StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T03:36:07.000Z"); return t}()),
	// 	Status: to.Ptr("Succeeded"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1e318cbfd2e239db54c80af5e6aea7fdf658851/specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/examples/sapcentralinstances/SAPCentralInstances_StopInstance.json
func ExampleSAPCentralInstancesClient_BeginStopInstance_stopTheSapCentralServicesInstance() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadssapvirtualinstance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSAPCentralInstancesClient().BeginStopInstance(ctx, "test-rg", "X00", "centralServer", &armworkloadssapvirtualinstance.SAPCentralInstancesClientBeginStopInstanceOptions{Body: &armworkloadssapvirtualinstance.StopRequest{
		SoftStopTimeoutSeconds: to.Ptr[int64](1200),
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationStatusResult = armworkloadssapvirtualinstance.OperationStatusResult{
	// 	Name: to.Ptr("centralServer"),
	// 	EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T03:38:07.000Z"); return t}()),
	// 	ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Workloads/sapVirtualInstances/X00/centralInstances/centralServer"),
	// 	StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T03:36:07.000Z"); return t}()),
	// 	Status: to.Ptr("Succeeded"),
	// }
}
