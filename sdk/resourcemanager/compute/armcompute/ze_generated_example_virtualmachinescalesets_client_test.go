//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/ListVirtualMachineScaleSetsInASubscriptionByLocation.json
func ExampleVirtualMachineScaleSetsClient_ListByLocation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	pager := client.ListByLocation("<location>",
		nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/CreateAScaleSetWithExtensionsSuppressFailuresEnabled.json
func ExampleVirtualMachineScaleSetsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		armcompute.VirtualMachineScaleSet{
			Location: to.StringPtr("<location>"),
			Properties: &armcompute.VirtualMachineScaleSetProperties{
				Overprovision: to.BoolPtr(true),
				UpgradePolicy: &armcompute.UpgradePolicy{
					Mode: armcompute.UpgradeModeManual.ToPtr(),
				},
				VirtualMachineProfile: &armcompute.VirtualMachineScaleSetVMProfile{
					DiagnosticsProfile: &armcompute.DiagnosticsProfile{
						BootDiagnostics: &armcompute.BootDiagnostics{
							Enabled:    to.BoolPtr(true),
							StorageURI: to.StringPtr("<storage-uri>"),
						},
					},
					ExtensionProfile: &armcompute.VirtualMachineScaleSetExtensionProfile{
						Extensions: []*armcompute.VirtualMachineScaleSetExtension{
							{
								Name: to.StringPtr("<name>"),
								Properties: &armcompute.VirtualMachineScaleSetExtensionProperties{
									Type:                    to.StringPtr("<type>"),
									AutoUpgradeMinorVersion: to.BoolPtr(false),
									Publisher:               to.StringPtr("<publisher>"),
									Settings:                map[string]interface{}{},
									SuppressFailures:        to.BoolPtr(true),
									TypeHandlerVersion:      to.StringPtr("<type-handler-version>"),
								},
							}},
					},
					NetworkProfile: &armcompute.VirtualMachineScaleSetNetworkProfile{
						NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetNetworkConfiguration{
							{
								Name: to.StringPtr("<name>"),
								Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
									EnableIPForwarding: to.BoolPtr(true),
									IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
										{
											Name: to.StringPtr("<name>"),
											Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
												Subnet: &armcompute.APIEntityReference{
													ID: to.StringPtr("<id>"),
												},
											},
										}},
									Primary: to.BoolPtr(true),
								},
							}},
					},
					OSProfile: &armcompute.VirtualMachineScaleSetOSProfile{
						AdminPassword:      to.StringPtr("<admin-password>"),
						AdminUsername:      to.StringPtr("<admin-username>"),
						ComputerNamePrefix: to.StringPtr("<computer-name-prefix>"),
					},
					StorageProfile: &armcompute.VirtualMachineScaleSetStorageProfile{
						ImageReference: &armcompute.ImageReference{
							Offer:     to.StringPtr("<offer>"),
							Publisher: to.StringPtr("<publisher>"),
							SKU:       to.StringPtr("<sku>"),
							Version:   to.StringPtr("<version>"),
						},
						OSDisk: &armcompute.VirtualMachineScaleSetOSDisk{
							Caching:      armcompute.CachingTypesReadWrite.ToPtr(),
							CreateOption: armcompute.DiskCreateOptionTypes("FromImage").ToPtr(),
							ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
								StorageAccountType: armcompute.StorageAccountTypes("Standard_LRS").ToPtr(),
							},
						},
					},
				},
			},
			SKU: &armcompute.SKU{
				Name:     to.StringPtr("<name>"),
				Capacity: to.Int64Ptr(3),
				Tier:     to.StringPtr("<tier>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.VirtualMachineScaleSetsClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_Update_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		armcompute.VirtualMachineScaleSetUpdate{
			Tags: map[string]*string{
				"key246": to.StringPtr("aaaaaaaaaaaaaaaaaaaaaaaa"),
			},
			Identity: &armcompute.VirtualMachineScaleSetIdentity{
				Type: armcompute.ResourceIdentityTypeSystemAssigned.ToPtr(),
				UserAssignedIdentities: map[string]*armcompute.VirtualMachineScaleSetIdentityUserAssignedIdentitiesValue{
					"key3951": {},
				},
			},
			Plan: &armcompute.Plan{
				Name:          to.StringPtr("<name>"),
				Product:       to.StringPtr("<product>"),
				PromotionCode: to.StringPtr("<promotion-code>"),
				Publisher:     to.StringPtr("<publisher>"),
			},
			Properties: &armcompute.VirtualMachineScaleSetUpdateProperties{
				AdditionalCapabilities: &armcompute.AdditionalCapabilities{
					HibernationEnabled: to.BoolPtr(true),
					UltraSSDEnabled:    to.BoolPtr(true),
				},
				AutomaticRepairsPolicy: &armcompute.AutomaticRepairsPolicy{
					Enabled:     to.BoolPtr(true),
					GracePeriod: to.StringPtr("<grace-period>"),
				},
				DoNotRunExtensionsOnOverprovisionedVMs: to.BoolPtr(true),
				Overprovision:                          to.BoolPtr(true),
				ProximityPlacementGroup: &armcompute.SubResource{
					ID: to.StringPtr("<id>"),
				},
				ScaleInPolicy: &armcompute.ScaleInPolicy{
					ForceDeletion: to.BoolPtr(true),
					Rules: []*armcompute.VirtualMachineScaleSetScaleInRules{
						armcompute.VirtualMachineScaleSetScaleInRules("OldestVM").ToPtr()},
				},
				SinglePlacementGroup: to.BoolPtr(true),
				UpgradePolicy: &armcompute.UpgradePolicy{
					AutomaticOSUpgradePolicy: &armcompute.AutomaticOSUpgradePolicy{
						DisableAutomaticRollback: to.BoolPtr(true),
						EnableAutomaticOSUpgrade: to.BoolPtr(true),
					},
					Mode: armcompute.UpgradeModeManual.ToPtr(),
					RollingUpgradePolicy: &armcompute.RollingUpgradePolicy{
						EnableCrossZoneUpgrade:              to.BoolPtr(true),
						MaxBatchInstancePercent:             to.Int32Ptr(49),
						MaxUnhealthyInstancePercent:         to.Int32Ptr(81),
						MaxUnhealthyUpgradedInstancePercent: to.Int32Ptr(98),
						PauseTimeBetweenBatches:             to.StringPtr("<pause-time-between-batches>"),
						PrioritizeUnhealthyInstances:        to.BoolPtr(true),
					},
				},
				VirtualMachineProfile: &armcompute.VirtualMachineScaleSetUpdateVMProfile{
					BillingProfile: &armcompute.BillingProfile{
						MaxPrice: to.Float64Ptr(-1),
					},
					DiagnosticsProfile: &armcompute.DiagnosticsProfile{
						BootDiagnostics: &armcompute.BootDiagnostics{
							Enabled:    to.BoolPtr(true),
							StorageURI: to.StringPtr("<storage-uri>"),
						},
					},
					ExtensionProfile: &armcompute.VirtualMachineScaleSetExtensionProfile{
						ExtensionsTimeBudget: to.StringPtr("<extensions-time-budget>"),
						Extensions: []*armcompute.VirtualMachineScaleSetExtension{
							{
								Name: to.StringPtr("<name>"),
								Properties: &armcompute.VirtualMachineScaleSetExtensionProperties{
									Type:                    to.StringPtr("<type>"),
									AutoUpgradeMinorVersion: to.BoolPtr(true),
									EnableAutomaticUpgrade:  to.BoolPtr(true),
									ForceUpdateTag:          to.StringPtr("<force-update-tag>"),
									ProtectedSettings:       map[string]interface{}{},
									ProvisionAfterExtensions: []*string{
										to.StringPtr("aa")},
									Publisher:          to.StringPtr("<publisher>"),
									Settings:           map[string]interface{}{},
									SuppressFailures:   to.BoolPtr(true),
									TypeHandlerVersion: to.StringPtr("<type-handler-version>"),
								},
							}},
					},
					LicenseType: to.StringPtr("<license-type>"),
					NetworkProfile: &armcompute.VirtualMachineScaleSetUpdateNetworkProfile{
						HealthProbe: &armcompute.APIEntityReference{
							ID: to.StringPtr("<id>"),
						},
						NetworkAPIVersion: armcompute.NetworkAPIVersion("2020-11-01").ToPtr(),
						NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetUpdateNetworkConfiguration{
							{
								ID:   to.StringPtr("<id>"),
								Name: to.StringPtr("<name>"),
								Properties: &armcompute.VirtualMachineScaleSetUpdateNetworkConfigurationProperties{
									DeleteOption: armcompute.DeleteOptions("Delete").ToPtr(),
									DNSSettings: &armcompute.VirtualMachineScaleSetNetworkConfigurationDNSSettings{
										DNSServers: []*string{},
									},
									EnableAcceleratedNetworking: to.BoolPtr(true),
									EnableFpga:                  to.BoolPtr(true),
									EnableIPForwarding:          to.BoolPtr(true),
									IPConfigurations: []*armcompute.VirtualMachineScaleSetUpdateIPConfiguration{
										{
											ID:   to.StringPtr("<id>"),
											Name: to.StringPtr("<name>"),
											Properties: &armcompute.VirtualMachineScaleSetUpdateIPConfigurationProperties{
												ApplicationGatewayBackendAddressPools: []*armcompute.SubResource{
													{
														ID: to.StringPtr("<id>"),
													}},
												ApplicationSecurityGroups: []*armcompute.SubResource{
													{
														ID: to.StringPtr("<id>"),
													}},
												LoadBalancerBackendAddressPools: []*armcompute.SubResource{
													{
														ID: to.StringPtr("<id>"),
													}},
												LoadBalancerInboundNatPools: []*armcompute.SubResource{
													{
														ID: to.StringPtr("<id>"),
													}},
												Primary:                 to.BoolPtr(true),
												PrivateIPAddressVersion: armcompute.IPVersion("IPv4").ToPtr(),
												PublicIPAddressConfiguration: &armcompute.VirtualMachineScaleSetUpdatePublicIPAddressConfiguration{
													Name: to.StringPtr("<name>"),
													Properties: &armcompute.VirtualMachineScaleSetUpdatePublicIPAddressConfigurationProperties{
														DeleteOption: armcompute.DeleteOptions("Delete").ToPtr(),
														DNSSettings: &armcompute.VirtualMachineScaleSetPublicIPAddressConfigurationDNSSettings{
															DomainNameLabel: to.StringPtr("<domain-name-label>"),
														},
														IdleTimeoutInMinutes: to.Int32Ptr(3),
													},
												},
												Subnet: &armcompute.APIEntityReference{
													ID: to.StringPtr("<id>"),
												},
											},
										}},
									NetworkSecurityGroup: &armcompute.SubResource{
										ID: to.StringPtr("<id>"),
									},
									Primary: to.BoolPtr(true),
								},
							}},
					},
					OSProfile: &armcompute.VirtualMachineScaleSetUpdateOSProfile{
						CustomData: to.StringPtr("<custom-data>"),
						LinuxConfiguration: &armcompute.LinuxConfiguration{
							DisablePasswordAuthentication: to.BoolPtr(true),
							PatchSettings: &armcompute.LinuxPatchSettings{
								AssessmentMode: armcompute.LinuxPatchAssessmentMode("ImageDefault").ToPtr(),
								PatchMode:      armcompute.LinuxVMGuestPatchMode("ImageDefault").ToPtr(),
							},
							ProvisionVMAgent: to.BoolPtr(true),
							SSH: &armcompute.SSHConfiguration{
								PublicKeys: []*armcompute.SSHPublicKey{
									{
										Path:    to.StringPtr("<path>"),
										KeyData: to.StringPtr("<key-data>"),
									}},
							},
						},
						Secrets: []*armcompute.VaultSecretGroup{
							{
								SourceVault: &armcompute.SubResource{
									ID: to.StringPtr("<id>"),
								},
								VaultCertificates: []*armcompute.VaultCertificate{
									{
										CertificateStore: to.StringPtr("<certificate-store>"),
										CertificateURL:   to.StringPtr("<certificate-url>"),
									}},
							}},
						WindowsConfiguration: &armcompute.WindowsConfiguration{
							AdditionalUnattendContent: []*armcompute.AdditionalUnattendContent{
								{
									ComponentName: to.StringPtr("<component-name>"),
									Content:       to.StringPtr("<content>"),
									PassName:      to.StringPtr("<pass-name>"),
									SettingName:   armcompute.SettingNamesAutoLogon.ToPtr(),
								}},
							EnableAutomaticUpdates: to.BoolPtr(true),
							PatchSettings: &armcompute.PatchSettings{
								AssessmentMode:    armcompute.WindowsPatchAssessmentMode("ImageDefault").ToPtr(),
								EnableHotpatching: to.BoolPtr(true),
								PatchMode:         armcompute.WindowsVMGuestPatchMode("AutomaticByOS").ToPtr(),
							},
							ProvisionVMAgent: to.BoolPtr(true),
							TimeZone:         to.StringPtr("<time-zone>"),
							WinRM: &armcompute.WinRMConfiguration{
								Listeners: []*armcompute.WinRMListener{
									{
										CertificateURL: to.StringPtr("<certificate-url>"),
										Protocol:       armcompute.ProtocolTypesHTTP.ToPtr(),
									}},
							},
						},
					},
					ScheduledEventsProfile: &armcompute.ScheduledEventsProfile{
						TerminateNotificationProfile: &armcompute.TerminateNotificationProfile{
							Enable:           to.BoolPtr(true),
							NotBeforeTimeout: to.StringPtr("<not-before-timeout>"),
						},
					},
					SecurityProfile: &armcompute.SecurityProfile{
						EncryptionAtHost: to.BoolPtr(true),
						SecurityType:     armcompute.SecurityTypes("TrustedLaunch").ToPtr(),
						UefiSettings: &armcompute.UefiSettings{
							SecureBootEnabled: to.BoolPtr(true),
							VTpmEnabled:       to.BoolPtr(true),
						},
					},
					StorageProfile: &armcompute.VirtualMachineScaleSetUpdateStorageProfile{
						DataDisks: []*armcompute.VirtualMachineScaleSetDataDisk{
							{
								Name:              to.StringPtr("<name>"),
								Caching:           armcompute.CachingTypesNone.ToPtr(),
								CreateOption:      armcompute.DiskCreateOptionTypes("Empty").ToPtr(),
								DiskIOPSReadWrite: to.Int64Ptr(28),
								DiskMBpsReadWrite: to.Int64Ptr(15),
								DiskSizeGB:        to.Int32Ptr(1023),
								Lun:               to.Int32Ptr(26),
								ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
									DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
										ID: to.StringPtr("<id>"),
									},
									StorageAccountType: armcompute.StorageAccountTypes("Standard_LRS").ToPtr(),
								},
								WriteAcceleratorEnabled: to.BoolPtr(true),
							}},
						ImageReference: &armcompute.ImageReference{
							ID:                   to.StringPtr("<id>"),
							Offer:                to.StringPtr("<offer>"),
							Publisher:            to.StringPtr("<publisher>"),
							SharedGalleryImageID: to.StringPtr("<shared-gallery-image-id>"),
							SKU:                  to.StringPtr("<sku>"),
							Version:              to.StringPtr("<version>"),
						},
						OSDisk: &armcompute.VirtualMachineScaleSetUpdateOSDisk{
							Caching:    armcompute.CachingTypesReadWrite.ToPtr(),
							DiskSizeGB: to.Int32Ptr(6),
							Image: &armcompute.VirtualHardDisk{
								URI: to.StringPtr("<uri>"),
							},
							ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
								DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
									ID: to.StringPtr("<id>"),
								},
								StorageAccountType: armcompute.StorageAccountTypes("Standard_LRS").ToPtr(),
							},
							VhdContainers: []*string{
								to.StringPtr("aa")},
							WriteAcceleratorEnabled: to.BoolPtr(true),
						},
					},
					UserData: to.StringPtr("<user-data>"),
				},
			},
			SKU: &armcompute.SKU{
				Name:     to.StringPtr("<name>"),
				Capacity: to.Int64Ptr(7),
				Tier:     to.StringPtr("<tier>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.VirtualMachineScaleSetsClientUpdateResult)
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/ForceDeleteVirtualMachineScaleSets.json
func ExampleVirtualMachineScaleSetsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		&armcompute.VirtualMachineScaleSetsClientBeginDeleteOptions{ForceDeletion: to.BoolPtr(true)})
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/GetVirtualMachineScaleSet.json
func ExampleVirtualMachineScaleSetsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		&armcompute.VirtualMachineScaleSetsClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.VirtualMachineScaleSetsClientGetResult)
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_Deallocate_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_BeginDeallocate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDeallocate(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		&armcompute.VirtualMachineScaleSetsClientBeginDeallocateOptions{VMInstanceIDs: &armcompute.VirtualMachineScaleSetVMInstanceIDs{
			InstanceIDs: []*string{
				to.StringPtr("aaaaaaaaaaaaaaaaa")},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_DeleteInstances_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_BeginDeleteInstances() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDeleteInstances(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		armcompute.VirtualMachineScaleSetVMInstanceRequiredIDs{
			InstanceIDs: []*string{
				to.StringPtr("aaaaaaaaaaaaaaaaaaaaaaaaa")},
		},
		&armcompute.VirtualMachineScaleSetsClientBeginDeleteInstancesOptions{ForceDeletion: to.BoolPtr(true)})
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_GetInstanceView_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_GetInstanceView() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	res, err := client.GetInstanceView(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.VirtualMachineScaleSetsClientGetInstanceViewResult)
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_List_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	pager := client.List("<resource-group-name>",
		nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_ListAll_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_ListAll() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	pager := client.ListAll(nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_GetOSUpgradeHistory_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_GetOSUpgradeHistory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	pager := client.GetOSUpgradeHistory("<resource-group-name>",
		"<vm-scale-set-name>",
		nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_PowerOff_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_BeginPowerOff() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginPowerOff(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		&armcompute.VirtualMachineScaleSetsClientBeginPowerOffOptions{SkipShutdown: to.BoolPtr(true),
			VMInstanceIDs: &armcompute.VirtualMachineScaleSetVMInstanceIDs{
				InstanceIDs: []*string{
					to.StringPtr("aaaaaaaaaaaaaaaaa")},
			},
		})
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_Restart_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_BeginRestart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginRestart(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		&armcompute.VirtualMachineScaleSetsClientBeginRestartOptions{VMInstanceIDs: &armcompute.VirtualMachineScaleSetVMInstanceIDs{
			InstanceIDs: []*string{
				to.StringPtr("aaaaaaaaaaaaaaaaa")},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_Start_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_BeginStart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginStart(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		&armcompute.VirtualMachineScaleSetsClientBeginStartOptions{VMInstanceIDs: &armcompute.VirtualMachineScaleSetVMInstanceIDs{
			InstanceIDs: []*string{
				to.StringPtr("aaaaaaaaaaaaaaaaa")},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_Redeploy_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_BeginRedeploy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginRedeploy(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		&armcompute.VirtualMachineScaleSetsClientBeginRedeployOptions{VMInstanceIDs: &armcompute.VirtualMachineScaleSetVMInstanceIDs{
			InstanceIDs: []*string{
				to.StringPtr("aaaaaaaaaaaaaaaaa")},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_PerformMaintenance_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_BeginPerformMaintenance() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginPerformMaintenance(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		&armcompute.VirtualMachineScaleSetsClientBeginPerformMaintenanceOptions{VMInstanceIDs: &armcompute.VirtualMachineScaleSetVMInstanceIDs{
			InstanceIDs: []*string{
				to.StringPtr("aaaaaaaaaaaaaaaaa")},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_UpdateInstances_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_BeginUpdateInstances() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdateInstances(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		armcompute.VirtualMachineScaleSetVMInstanceRequiredIDs{
			InstanceIDs: []*string{
				to.StringPtr("aaaaaaaaaaaaaaaaaaaaaaaaa")},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_Reimage_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_BeginReimage() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginReimage(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		&armcompute.VirtualMachineScaleSetsClientBeginReimageOptions{VMScaleSetReimageInput: &armcompute.VirtualMachineScaleSetReimageParameters{
			TempDisk: to.BoolPtr(true),
			InstanceIDs: []*string{
				to.StringPtr("aaaaaaaaaa")},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_ReimageAll_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_BeginReimageAll() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginReimageAll(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		&armcompute.VirtualMachineScaleSetsClientBeginReimageAllOptions{VMInstanceIDs: &armcompute.VirtualMachineScaleSetVMInstanceIDs{
			InstanceIDs: []*string{
				to.StringPtr("aaaaaaaaaaaaaaaaa")},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_ForceRecoveryServiceFabricPlatformUpdateDomainWalk_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_ForceRecoveryServiceFabricPlatformUpdateDomainWalk() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	res, err := client.ForceRecoveryServiceFabricPlatformUpdateDomainWalk(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		30,
		&armcompute.VirtualMachineScaleSetsClientForceRecoveryServiceFabricPlatformUpdateDomainWalkOptions{Zone: nil,
			PlacementGroupID: nil,
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.VirtualMachineScaleSetsClientForceRecoveryServiceFabricPlatformUpdateDomainWalkResult)
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_ConvertToSinglePlacementGroup_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_ConvertToSinglePlacementGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	_, err = client.ConvertToSinglePlacementGroup(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		armcompute.VMScaleSetConvertToSinglePlacementGroupInput{
			ActivePlacementGroupID: to.StringPtr("<active-placement-group-id>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_SetOrchestrationServiceState_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_BeginSetOrchestrationServiceState() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginSetOrchestrationServiceState(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		armcompute.OrchestrationServiceStateInput{
			Action:      armcompute.OrchestrationServiceStateAction("Resume").ToPtr(),
			ServiceName: armcompute.OrchestrationServiceNames("AutomaticRepairs").ToPtr(),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
