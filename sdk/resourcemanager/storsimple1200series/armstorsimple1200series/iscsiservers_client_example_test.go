//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armstorsimple1200series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/IscsiServersListByDevice.json
func ExampleIscsiServersClient_NewListByDevicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIscsiServersClient().NewListByDevicePager("HSDK-UGU4PITWNI", "ResourceGroupForSDKTest", "hAzureSDKOperations", nil)
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
		// page.ISCSIServerList = armstorsimple1200series.ISCSIServerList{
		// 	Value: []*armstorsimple1200series.ISCSIServer{
		// 	},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/IscsiServersGet.json
func ExampleIscsiServersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIscsiServersClient().Get(ctx, "HSDK-WSJQERQW3F", "HSDK-WSJQERQW3F", "ResourceGroupForSDKTest", "hAzureSDKOperations", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ISCSIServer = armstorsimple1200series.ISCSIServer{
	// 	Name: to.Ptr("HSDK-WSJQERQW3F"),
	// 	Type: to.Ptr("Microsoft.StorSimple/managers/devices/iscsiServers"),
	// 	ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-WSJQERQW3F/iscsiServers/HSDK-WSJQERQW3F"),
	// 	Properties: &armstorsimple1200series.ISCSIServerProperties{
	// 		BackupScheduleGroupID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-WSJQERQW3F/backupScheduleGroups/Default-HSDK-WSJQERQW3F-BackupScheduleGroup"),
	// 		ChapID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-WSJQERQW3F/chapSettings/ChapSettingForSDK"),
	// 		ReverseChapID: to.Ptr(""),
	// 		StorageDomainID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/storageDomains/Default-HSDK-WSJQERQW3F-StorageDomain"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/IscsiServersCreateOrUpdate.json
func ExampleIscsiServersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIscsiServersClient().BeginCreateOrUpdate(ctx, "HSDK-WSJQERQW3F", "HSDK-WSJQERQW3F", "ResourceGroupForSDKTest", "hAzureSDKOperations", armstorsimple1200series.ISCSIServer{
		Name: to.Ptr("HSDK-WSJQERQW3F"),
		Type: to.Ptr("Microsoft.StorSimple/managers/devices/iscsiServers"),
		ID:   to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-WSJQERQW3F/iscsiServers/HSDK-WSJQERQW3F"),
		Properties: &armstorsimple1200series.ISCSIServerProperties{
			BackupScheduleGroupID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-WSJQERQW3F/backupScheduleGroups/Default-HSDK-WSJQERQW3F-BackupScheduleGroup"),
			ChapID:                to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-WSJQERQW3F/chapSettings/ChapSettingForSDK"),
			ReverseChapID:         to.Ptr(""),
			StorageDomainID:       to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/storageDomains/Default-HSDK-WSJQERQW3F-StorageDomain"),
		},
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
	// res.ISCSIServer = armstorsimple1200series.ISCSIServer{
	// 	Name: to.Ptr("HSDK-WSJQERQW3F"),
	// 	Type: to.Ptr("Microsoft.StorSimple/managers/devices/iscsiServers"),
	// 	ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-WSJQERQW3F/iscsiServers/HSDK-WSJQERQW3F"),
	// 	Properties: &armstorsimple1200series.ISCSIServerProperties{
	// 		BackupScheduleGroupID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-WSJQERQW3F/backupScheduleGroups/Default-HSDK-WSJQERQW3F-BackupScheduleGroup"),
	// 		ChapID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-WSJQERQW3F/chapSettings/ChapSettingForSDK"),
	// 		ReverseChapID: to.Ptr(""),
	// 		StorageDomainID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/storageDomains/Default-HSDK-WSJQERQW3F-StorageDomain"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/IscsiServersDelete.json
func ExampleIscsiServersClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIscsiServersClient().BeginDelete(ctx, "HSDK-UGU4PITWNI", "HSDK-UGU4PITWNI", "ResourceGroupForSDKTest", "hAzureSDKOperations", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/IscsiServersBackupNow.json
func ExampleIscsiServersClient_BeginBackupNow() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIscsiServersClient().BeginBackupNow(ctx, "HSDK-0NZI14MDTF", "HSDK-0NZI14MDTF", "ResourceGroupForSDKTest", "hAzureSDKOperations", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/IscsiServersListMetrics.json
func ExampleIscsiServersClient_NewListMetricsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIscsiServersClient().NewListMetricsPager("HSDK-UGU4PITWNI", "HSDK-UGU4PITWNI", "ResourceGroupForSDKTest", "hAzureSDKOperations", &armstorsimple1200series.IscsiServersClientListMetricsOptions{Filter: to.Ptr("startTime%20ge%20'2018-08-10T18:30:00Z'%20and%20endTime%20le%20'2018-08-11T18:30:00Z'")})
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
		// page.MetricList = armstorsimple1200series.MetricList{
		// 	Value: []*armstorsimple1200series.Metrics{
		// 		{
		// 			Name: &armstorsimple1200series.MetricName{
		// 				LocalizedValue: to.Ptr("Cloud Storage Used"),
		// 				Value: to.Ptr("CloudConsumedStorage"),
		// 			},
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/iscsiServers/metrics"),
		// 			Dimensions: []*armstorsimple1200series.MetricDimension{
		// 				{
		// 					Name: to.Ptr("IscsiServer"),
		// 					Value: to.Ptr("HSDK-UGU4PITWNI"),
		// 			}},
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-11T18:30:00.000Z"); return t}()),
		// 			PrimaryAggregation: to.Ptr(armstorsimple1200series.MetricAggregationTypeAverage),
		// 			ResourceID: to.Ptr("https://pod01-cis2.sea.storsimple.windowsazure.com/managers/4239154091695873374/devices/HSDK-UGU4PITWNI/iscsiservers/HSDK-UGU4PITWNI"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T18:30:00.000Z"); return t}()),
		// 			TimeGrain: to.Ptr("1.00:00:00"),
		// 			Unit: to.Ptr(armstorsimple1200series.MetricUnitBytes),
		// 			Values: []*armstorsimple1200series.MetricData{
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/IscsiServersListMetricDefinition.json
func ExampleIscsiServersClient_NewListMetricDefinitionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIscsiServersClient().NewListMetricDefinitionPager("HSDK-UGU4PITWNI", "HSDK-UGU4PITWNI", "ResourceGroupForSDKTest", "hAzureSDKOperations", nil)
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
		// page.MetricDefinitionList = armstorsimple1200series.MetricDefinitionList{
		// 	Value: []*armstorsimple1200series.MetricDefinition{
		// 		{
		// 			Name: &armstorsimple1200series.MetricName{
		// 				LocalizedValue: to.Ptr("Cloud Storage Used"),
		// 				Value: to.Ptr("CloudConsumedStorage"),
		// 			},
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/iscsiServers/metricsDefinitions"),
		// 			Dimensions: []*armstorsimple1200series.MetricDimension{
		// 				{
		// 					Name: to.Ptr("IscsiServer"),
		// 					Value: to.Ptr("HSDK-UGU4PITWNI"),
		// 			}},
		// 			MetricAvailabilities: []*armstorsimple1200series.MetricAvailablity{
		// 				{
		// 					Retention: to.Ptr("PT6H"),
		// 					TimeGrain: to.Ptr("PT1M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P7D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P1Y"),
		// 					TimeGrain: to.Ptr("P1D"),
		// 			}},
		// 			PrimaryAggregationType: to.Ptr(armstorsimple1200series.MetricAggregationTypeAverage),
		// 			ResourceID: to.Ptr("https://pod01-cis2.sea.storsimple.windowsazure.com/managers/4239154091695873374/devices/HSDK-UGU4PITWNI/iscsiservers/HSDK-UGU4PITWNI"),
		// 			Unit: to.Ptr(armstorsimple1200series.MetricUnitBytes),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/IscsiServersListByManager.json
func ExampleIscsiServersClient_NewListByManagerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIscsiServersClient().NewListByManagerPager("ResourceGroupForSDKTest", "hAzureSDKOperations", nil)
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
		// page.ISCSIServerList = armstorsimple1200series.ISCSIServerList{
		// 	Value: []*armstorsimple1200series.ISCSIServer{
		// 		{
		// 			Name: to.Ptr("HSDK-WSJQERQW3F"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/iscsiServers"),
		// 			ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-WSJQERQW3F/iscsiServers/HSDK-WSJQERQW3F"),
		// 			Properties: &armstorsimple1200series.ISCSIServerProperties{
		// 				BackupScheduleGroupID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-WSJQERQW3F/backupScheduleGroups/Default-HSDK-WSJQERQW3F-BackupScheduleGroup"),
		// 				ChapID: to.Ptr(""),
		// 				ReverseChapID: to.Ptr(""),
		// 				StorageDomainID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/storageDomains/Default-HSDK-WSJQERQW3F-StorageDomain"),
		// 			},
		// 	}},
		// }
	}
}
