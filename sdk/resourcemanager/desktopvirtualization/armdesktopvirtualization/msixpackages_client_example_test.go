//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdesktopvirtualization_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/MsixPackage_Get.json
func ExampleMSIXPackagesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdesktopvirtualization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMSIXPackagesClient().Get(ctx, "resourceGroup1", "hostpool1", "packagefullname", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MSIXPackage = armdesktopvirtualization.MSIXPackage{
	// 	Name: to.Ptr("hostpool1/MsixPackageFullName"),
	// 	Type: to.Ptr("Microsoft.DesktopVirtualization/hostpools/msixpackages"),
	// 	ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourcegroups/resourcegroup1/providers/Microsoft.DesktopVirtualization/hostpools/hostpool1/msixpackages/msixPackageFullName"),
	// 	Properties: &armdesktopvirtualization.MSIXPackageProperties{
	// 		DisplayName: to.Ptr("dis"),
	// 		ImagePath: to.Ptr("imagepath"),
	// 		IsActive: to.Ptr(false),
	// 		IsRegularRegistration: to.Ptr(false),
	// 		LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2008-09-22T14:01:54.9571247Z"); return t}()),
	// 		PackageApplications: []*armdesktopvirtualization.MsixPackageApplications{
	// 			{
	// 				Description: to.Ptr("desc"),
	// 				AppID: to.Ptr("Application_Id"),
	// 				AppUserModelID: to.Ptr("Application_ModelID"),
	// 				FriendlyName: to.Ptr("fri"),
	// 				IconImageName: to.Ptr("Apptile"),
	// 				RawIcon: []byte("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
	// 				RawPNG: []byte("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
	// 		}},
	// 		PackageDependencies: []*armdesktopvirtualization.MsixPackageDependencies{
	// 			{
	// 				DependencyName: to.Ptr("MsixPackage_Dependency_Name"),
	// 				MinVersion: to.Ptr("packageDep_version"),
	// 				Publisher: to.Ptr("MsixPackage_Dependency_Publisher"),
	// 		}},
	// 		PackageFamilyName: to.Ptr("MsixPackage_FamilyName"),
	// 		PackageName: to.Ptr("MsixPackage_Name"),
	// 		PackageRelativePath: to.Ptr("MsixPackage_RelativePackageRoot"),
	// 		Version: to.Ptr("version"),
	// 	},
	// 	SystemData: &armdesktopvirtualization.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/MsixPackage_Create.json
func ExampleMSIXPackagesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdesktopvirtualization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMSIXPackagesClient().CreateOrUpdate(ctx, "resourceGroup1", "hostpool1", "msixpackagefullname", armdesktopvirtualization.MSIXPackage{
		Properties: &armdesktopvirtualization.MSIXPackageProperties{
			DisplayName:           to.Ptr("displayname"),
			ImagePath:             to.Ptr("imagepath"),
			IsActive:              to.Ptr(false),
			IsRegularRegistration: to.Ptr(false),
			LastUpdated:           to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2008-09-22T14:01:54.9571247Z"); return t }()),
			PackageApplications: []*armdesktopvirtualization.MsixPackageApplications{
				{
					Description:    to.Ptr("application-desc"),
					AppID:          to.Ptr("ApplicationId"),
					AppUserModelID: to.Ptr("AppUserModelId"),
					FriendlyName:   to.Ptr("friendlyname"),
					IconImageName:  to.Ptr("Apptile"),
					RawIcon:        []byte("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
					RawPNG:         []byte("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
				}},
			PackageDependencies: []*armdesktopvirtualization.MsixPackageDependencies{
				{
					DependencyName: to.Ptr("MsixTest_Dependency_Name"),
					MinVersion:     to.Ptr("version"),
					Publisher:      to.Ptr("PublishedName"),
				}},
			PackageFamilyName:   to.Ptr("MsixPackage_FamilyName"),
			PackageName:         to.Ptr("MsixPackage_name"),
			PackageRelativePath: to.Ptr("packagerelativepath"),
			Version:             to.Ptr("version"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MSIXPackage = armdesktopvirtualization.MSIXPackage{
	// 	Name: to.Ptr("hostpool1/MsixPackageFullName"),
	// 	Type: to.Ptr("Microsoft.DesktopVirtualization/hostpools/msixpackages"),
	// 	ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourcegroups/resourcegroup1/providers/Microsoft.DesktopVirtualization/hostpools/hostpool1/msixpackages/msixPackageFullName"),
	// 	Properties: &armdesktopvirtualization.MSIXPackageProperties{
	// 		DisplayName: to.Ptr("dis"),
	// 		ImagePath: to.Ptr("imagepath"),
	// 		IsActive: to.Ptr(false),
	// 		IsRegularRegistration: to.Ptr(false),
	// 		LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2008-09-22T14:01:54.9571247Z"); return t}()),
	// 		PackageApplications: []*armdesktopvirtualization.MsixPackageApplications{
	// 			{
	// 				Description: to.Ptr("desc"),
	// 				AppID: to.Ptr("Application_Id"),
	// 				AppUserModelID: to.Ptr("Application_ModelID"),
	// 				FriendlyName: to.Ptr("fri"),
	// 				IconImageName: to.Ptr("Apptile"),
	// 				RawIcon: []byte("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
	// 				RawPNG: []byte("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
	// 		}},
	// 		PackageDependencies: []*armdesktopvirtualization.MsixPackageDependencies{
	// 			{
	// 				DependencyName: to.Ptr("MsixPackage_Dependency_Name"),
	// 				MinVersion: to.Ptr("packageDep_version"),
	// 				Publisher: to.Ptr("MsixPackage_Dependency_Publisher"),
	// 		}},
	// 		PackageFamilyName: to.Ptr("MsixPackage_FamilyName"),
	// 		PackageName: to.Ptr("MsixPackage_Name"),
	// 		PackageRelativePath: to.Ptr("MsixPackage_RelativePackageRoot"),
	// 		Version: to.Ptr("version"),
	// 	},
	// 	SystemData: &armdesktopvirtualization.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/MsixPackage_Delete.json
func ExampleMSIXPackagesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdesktopvirtualization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewMSIXPackagesClient().Delete(ctx, "resourceGroup1", "hostpool1", "packagefullname", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/MsixPackage_Update.json
func ExampleMSIXPackagesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdesktopvirtualization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMSIXPackagesClient().Update(ctx, "resourceGroup1", "hostpool1", "msixpackagefullname", &armdesktopvirtualization.MSIXPackagesClientUpdateOptions{MsixPackage: &armdesktopvirtualization.MSIXPackagePatch{
		Properties: &armdesktopvirtualization.MSIXPackagePatchProperties{
			DisplayName:           to.Ptr("displayname"),
			IsActive:              to.Ptr(true),
			IsRegularRegistration: to.Ptr(false),
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MSIXPackage = armdesktopvirtualization.MSIXPackage{
	// 	Name: to.Ptr("hostpool1/MsixPackageFullName"),
	// 	Type: to.Ptr("Microsoft.DesktopVirtualization/hostpools/msixpackages"),
	// 	ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourcegroups/resourcegroup1/providers/Microsoft.DesktopVirtualization/hostpools/hostpool1/msixpackages/msixPackageFullName"),
	// 	Properties: &armdesktopvirtualization.MSIXPackageProperties{
	// 		DisplayName: to.Ptr("dis"),
	// 		ImagePath: to.Ptr("imagepath"),
	// 		IsActive: to.Ptr(true),
	// 		IsRegularRegistration: to.Ptr(false),
	// 		LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2008-09-22T14:01:54.9571247Z"); return t}()),
	// 		PackageApplications: []*armdesktopvirtualization.MsixPackageApplications{
	// 			{
	// 				Description: to.Ptr("desc"),
	// 				AppID: to.Ptr("Application_Id"),
	// 				AppUserModelID: to.Ptr("Application_ModelID"),
	// 				FriendlyName: to.Ptr("fri"),
	// 				IconImageName: to.Ptr("Apptile"),
	// 				RawIcon: []byte("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
	// 				RawPNG: []byte("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
	// 		}},
	// 		PackageDependencies: []*armdesktopvirtualization.MsixPackageDependencies{
	// 			{
	// 				DependencyName: to.Ptr("MsixPackage_Dependency_Name"),
	// 				MinVersion: to.Ptr("packageDep_version"),
	// 				Publisher: to.Ptr("MsixPackage_Dependency_Publisher"),
	// 		}},
	// 		PackageFamilyName: to.Ptr("MsixPackage_FamilyName"),
	// 		PackageName: to.Ptr("MsixPackage_Name"),
	// 		PackageRelativePath: to.Ptr("MsixPackage_RelativePackageRoot"),
	// 		Version: to.Ptr("version"),
	// 	},
	// 	SystemData: &armdesktopvirtualization.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/MsixPackage_List.json
func ExampleMSIXPackagesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdesktopvirtualization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMSIXPackagesClient().NewListPager("resourceGroup1", "hostpool1", &armdesktopvirtualization.MSIXPackagesClientListOptions{PageSize: to.Ptr[int32](10),
		IsDescending: to.Ptr(true),
		InitialSkip:  to.Ptr[int32](0),
	})
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
		// page.MSIXPackageList = armdesktopvirtualization.MSIXPackageList{
		// 	Value: []*armdesktopvirtualization.MSIXPackage{
		// 		{
		// 			Name: to.Ptr("hostpool1/MsixPackageFullName"),
		// 			Type: to.Ptr("Microsoft.DesktopVirtualization/hostpools/msixpackages"),
		// 			ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourcegroups/resourcegroup1/providers/Microsoft.DesktopVirtualization/hostpools/hostpool1/msixpackages/msixPackageFullName"),
		// 			Properties: &armdesktopvirtualization.MSIXPackageProperties{
		// 				DisplayName: to.Ptr("dis"),
		// 				ImagePath: to.Ptr("imagepath"),
		// 				IsActive: to.Ptr(false),
		// 				IsRegularRegistration: to.Ptr(false),
		// 				LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2008-09-22T14:01:54.9571247Z"); return t}()),
		// 				PackageApplications: []*armdesktopvirtualization.MsixPackageApplications{
		// 					{
		// 						Description: to.Ptr("desc"),
		// 						AppID: to.Ptr("Application_Id"),
		// 						AppUserModelID: to.Ptr("Application_ModelID"),
		// 						FriendlyName: to.Ptr("fri"),
		// 						IconImageName: to.Ptr("Apptile"),
		// 						RawIcon: []byte("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
		// 						RawPNG: []byte("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
		// 				}},
		// 				PackageDependencies: []*armdesktopvirtualization.MsixPackageDependencies{
		// 					{
		// 						DependencyName: to.Ptr("MsixPackage_Dependency_Name"),
		// 						MinVersion: to.Ptr("packageDep_version"),
		// 						Publisher: to.Ptr("MsixPackage_Dependency_Publisher"),
		// 				}},
		// 				PackageFamilyName: to.Ptr("MsixPackage_FamilyName"),
		// 				PackageName: to.Ptr("MsixPackage_Name"),
		// 				PackageRelativePath: to.Ptr("MsixPackage_RelativePackageRoot"),
		// 				Version: to.Ptr("version"),
		// 			},
		// 			SystemData: &armdesktopvirtualization.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("hostpool1/MsixPackageFullName2"),
		// 			Type: to.Ptr("Microsoft.DesktopVirtualization/hostpools/msixpackages"),
		// 			ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourcegroups/resourcegroup1/providers/Microsoft.DesktopVirtualization/hostpools/hostpool1/msixpackages/msixPackageFullName2"),
		// 			Properties: &armdesktopvirtualization.MSIXPackageProperties{
		// 				DisplayName: to.Ptr("dis2"),
		// 				ImagePath: to.Ptr("imagepath2"),
		// 				IsActive: to.Ptr(false),
		// 				IsRegularRegistration: to.Ptr(false),
		// 				LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2008-09-22T14:01:54.9571247Z"); return t}()),
		// 				PackageApplications: []*armdesktopvirtualization.MsixPackageApplications{
		// 					{
		// 						Description: to.Ptr("desc2"),
		// 						AppID: to.Ptr("Application_Id2"),
		// 						AppUserModelID: to.Ptr("Application_ModelID2"),
		// 						FriendlyName: to.Ptr("fri2"),
		// 						IconImageName: to.Ptr("Apptile2"),
		// 						RawIcon: []byte("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
		// 						RawPNG: []byte("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
		// 				}},
		// 				PackageDependencies: []*armdesktopvirtualization.MsixPackageDependencies{
		// 					{
		// 						DependencyName: to.Ptr("MsixPackage_Dependency_Name2"),
		// 						MinVersion: to.Ptr("packageDep_version2"),
		// 						Publisher: to.Ptr("MsixPackage_Dependency_Publisher2"),
		// 				}},
		// 				PackageFamilyName: to.Ptr("MsixPackage_FamilyName2"),
		// 				PackageName: to.Ptr("MsixPackage_Name2"),
		// 				PackageRelativePath: to.Ptr("MsixPackage_RelativePackageRoot2"),
		// 				Version: to.Ptr("version2"),
		// 			},
		// 			SystemData: &armdesktopvirtualization.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}
