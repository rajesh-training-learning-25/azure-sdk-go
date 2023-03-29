//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/watchlists/GetWatchlists.json
func ExampleWatchlistsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWatchlistsClient().NewListPager("myRg", "myWorkspace", &armsecurityinsights.WatchlistsClientListOptions{SkipToken: nil})
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
		// page.WatchlistList = armsecurityinsights.WatchlistList{
		// 	Value: []*armsecurityinsights.Watchlist{
		// 		{
		// 			Name: to.Ptr("highValueAsset"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/Watchlists"),
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalIinsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/watchlists/highValueAsset"),
		// 			Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		// 			Properties: &armsecurityinsights.WatchlistProperties{
		// 				Description: to.Ptr("Watchlist from CSV content"),
		// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-28T00:26:54.7746089+00:00"); return t}()),
		// 				CreatedBy: &armsecurityinsights.UserInfo{
		// 					Name: to.Ptr("john doe"),
		// 					Email: to.Ptr("john@contoso.com"),
		// 					ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
		// 				},
		// 				DefaultDuration: to.Ptr("P1279DT12H30M5S"),
		// 				DisplayName: to.Ptr("High Value Assets Watchlist"),
		// 				IsDeleted: to.Ptr(false),
		// 				ItemsSearchKey: to.Ptr("header1"),
		// 				Labels: []*string{
		// 					to.Ptr("Tag1"),
		// 					to.Ptr("Tag2")},
		// 					Provider: to.Ptr("Microsoft"),
		// 					Source: to.Ptr("watchlist.csv"),
		// 					SourceType: to.Ptr(armsecurityinsights.SourceTypeLocalFile),
		// 					TenantID: to.Ptr("f686d426-8d16-42db-81b7-ab578e110ccd"),
		// 					Updated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-28T00:26:57+00:00"); return t}()),
		// 					UpdatedBy: &armsecurityinsights.UserInfo{
		// 						Name: to.Ptr("john doe"),
		// 						Email: to.Ptr("john@contoso.com"),
		// 						ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
		// 					},
		// 					WatchlistAlias: to.Ptr("highValueAsset"),
		// 					WatchlistID: to.Ptr("76d5a51f-ba1f-4038-9d22-59fda38dc017"),
		// 					WatchlistType: to.Ptr("watchlist"),
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/watchlists/GetWatchlistByAlias.json
func ExampleWatchlistsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWatchlistsClient().Get(ctx, "myRg", "myWorkspace", "highValueAsset", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Watchlist = armsecurityinsights.Watchlist{
	// 	Name: to.Ptr("highValueAsset"),
	// 	Type: to.Ptr("Microsoft.SecurityInsights/Watchlists"),
	// 	ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalIinsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/watchlists/highValueAsset"),
	// 	Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
	// 	Properties: &armsecurityinsights.WatchlistProperties{
	// 		Description: to.Ptr("Watchlist from CSV content"),
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-28T00:26:54.7746089+00:00"); return t}()),
	// 		CreatedBy: &armsecurityinsights.UserInfo{
	// 			Name: to.Ptr("john doe"),
	// 			Email: to.Ptr("john@contoso.com"),
	// 			ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
	// 		},
	// 		DefaultDuration: to.Ptr("P1279DT12H30M5S"),
	// 		DisplayName: to.Ptr("High Value Assets Watchlist"),
	// 		IsDeleted: to.Ptr(false),
	// 		ItemsSearchKey: to.Ptr("header1"),
	// 		Labels: []*string{
	// 			to.Ptr("Tag1"),
	// 			to.Ptr("Tag2")},
	// 			Provider: to.Ptr("Microsoft"),
	// 			Source: to.Ptr("watchlist.csv"),
	// 			SourceType: to.Ptr(armsecurityinsights.SourceTypeLocalFile),
	// 			TenantID: to.Ptr("f686d426-8d16-42db-81b7-ab578e110ccd"),
	// 			Updated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-28T00:26:57+00:00"); return t}()),
	// 			UpdatedBy: &armsecurityinsights.UserInfo{
	// 				Name: to.Ptr("john doe"),
	// 				Email: to.Ptr("john@contoso.com"),
	// 				ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
	// 			},
	// 			WatchlistAlias: to.Ptr("highValueAsset"),
	// 			WatchlistID: to.Ptr("76d5a51f-ba1f-4038-9d22-59fda38dc017"),
	// 			WatchlistType: to.Ptr("watchlist"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/watchlists/DeleteWatchlist.json
func ExampleWatchlistsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewWatchlistsClient().Delete(ctx, "myRg", "myWorkspace", "highValueAsset", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/watchlists/CreateWatchlistAndWatchlistItems.json
func ExampleWatchlistsClient_CreateOrUpdate_createsOrUpdatesAWatchlistAndBulkCreatesWatchlistItems() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWatchlistsClient().CreateOrUpdate(ctx, "myRg", "myWorkspace", "highValueAsset", armsecurityinsights.Watchlist{
		Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		Properties: &armsecurityinsights.WatchlistProperties{
			Description:         to.Ptr("Watchlist from CSV content"),
			ContentType:         to.Ptr("text/csv"),
			DisplayName:         to.Ptr("High Value Assets Watchlist"),
			ItemsSearchKey:      to.Ptr("header1"),
			NumberOfLinesToSkip: to.Ptr[int32](1),
			Provider:            to.Ptr("Microsoft"),
			RawContent:          to.Ptr("This line will be skipped\nheader1,header2\nvalue1,value2"),
			Source:              to.Ptr("watchlist.csv"),
			SourceType:          to.Ptr(armsecurityinsights.SourceTypeLocalFile),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Watchlist = armsecurityinsights.Watchlist{
	// 	Name: to.Ptr("highValueAsset"),
	// 	Type: to.Ptr("Microsoft.SecurityInsights/Watchlists"),
	// 	ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalIinsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/watchlists/highValueAsset"),
	// 	Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
	// 	Properties: &armsecurityinsights.WatchlistProperties{
	// 		Description: to.Ptr("Watchlist from CSV content"),
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-28T00:26:54.7746089+00:00"); return t}()),
	// 		CreatedBy: &armsecurityinsights.UserInfo{
	// 			Name: to.Ptr("john doe"),
	// 			Email: to.Ptr("john@contoso.com"),
	// 			ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
	// 		},
	// 		DisplayName: to.Ptr("High Value Assets Watchlist"),
	// 		IsDeleted: to.Ptr(false),
	// 		ItemsSearchKey: to.Ptr("header1"),
	// 		Provider: to.Ptr("Microsoft"),
	// 		Source: to.Ptr("watchlist.csv"),
	// 		SourceType: to.Ptr(armsecurityinsights.SourceTypeLocalFile),
	// 		TenantID: to.Ptr("f686d426-8d16-42db-81b7-ab578e110ccd"),
	// 		Updated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-28T00:26:57+00:00"); return t}()),
	// 		UpdatedBy: &armsecurityinsights.UserInfo{
	// 			Name: to.Ptr("john doe"),
	// 			Email: to.Ptr("john@contoso.com"),
	// 			ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
	// 		},
	// 		WatchlistAlias: to.Ptr("highValueAsset"),
	// 		WatchlistID: to.Ptr("76d5a51f-ba1f-4038-9d22-59fda38dc017"),
	// 		WatchlistType: to.Ptr("watchlist"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/watchlists/CreateWatchlist.json
func ExampleWatchlistsClient_CreateOrUpdate_createsOrUpdatesAWatchlist() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWatchlistsClient().CreateOrUpdate(ctx, "myRg", "myWorkspace", "highValueAsset", armsecurityinsights.Watchlist{
		Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		Properties: &armsecurityinsights.WatchlistProperties{
			Description:    to.Ptr("Watchlist from CSV content"),
			DisplayName:    to.Ptr("High Value Assets Watchlist"),
			ItemsSearchKey: to.Ptr("header1"),
			Provider:       to.Ptr("Microsoft"),
			Source:         to.Ptr("watchlist.csv"),
			SourceType:     to.Ptr(armsecurityinsights.SourceTypeLocalFile),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Watchlist = armsecurityinsights.Watchlist{
	// 	Name: to.Ptr("highValueAsset"),
	// 	Type: to.Ptr("Microsoft.SecurityInsights/Watchlists"),
	// 	ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalIinsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/watchlists/highValueAsset"),
	// 	Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
	// 	Properties: &armsecurityinsights.WatchlistProperties{
	// 		Description: to.Ptr("Watchlist from CSV content"),
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-28T00:26:54.7746089+00:00"); return t}()),
	// 		CreatedBy: &armsecurityinsights.UserInfo{
	// 			Name: to.Ptr("john doe"),
	// 			Email: to.Ptr("john@contoso.com"),
	// 			ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
	// 		},
	// 		DisplayName: to.Ptr("High Value Assets Watchlist"),
	// 		IsDeleted: to.Ptr(false),
	// 		ItemsSearchKey: to.Ptr("header1"),
	// 		Provider: to.Ptr("Microsoft"),
	// 		Source: to.Ptr("watchlist.csv"),
	// 		SourceType: to.Ptr(armsecurityinsights.SourceTypeLocalFile),
	// 		TenantID: to.Ptr("f686d426-8d16-42db-81b7-ab578e110ccd"),
	// 		Updated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-28T00:26:57+00:00"); return t}()),
	// 		UpdatedBy: &armsecurityinsights.UserInfo{
	// 			Name: to.Ptr("john doe"),
	// 			Email: to.Ptr("john@contoso.com"),
	// 			ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
	// 		},
	// 		WatchlistAlias: to.Ptr("highValueAsset"),
	// 		WatchlistID: to.Ptr("76d5a51f-ba1f-4038-9d22-59fda38dc017"),
	// 		WatchlistType: to.Ptr("watchlist"),
	// 	},
	// }
}
