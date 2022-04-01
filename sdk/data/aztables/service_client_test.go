// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package aztables

import (
	"fmt"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/recording"
	"github.com/stretchr/testify/require"
)

func TestServiceErrorsServiceClient(t *testing.T) {
	for _, service := range services {
		t.Run(fmt.Sprintf("%v_%v", t.Name(), service), func(t *testing.T) {
			service, delete := initServiceTest(t, service)
			defer delete()

			tableName, err := createRandomName(t, tableNamePrefix)
			require.NoError(t, err)

			_, err = service.CreateTable(ctx, tableName, nil)
			require.NoError(t, err)

			// Create a duplicate table to produce an error
			_, err = service.CreateTable(ctx, tableName, nil)
			require.Error(t, err)

			_, err = service.DeleteTable(ctx, tableName, nil)
			require.NoError(t, err)
		})
	}
}

func TestCreateTableFromService(t *testing.T) {
	for _, service := range services {
		t.Run(fmt.Sprintf("%v_%v", t.Name(), service), func(t *testing.T) {
			service, delete := initServiceTest(t, service)
			defer delete()

			tableName, err := createRandomName(t, tableNamePrefix)
			require.NoError(t, err)

			_, err = service.CreateTable(ctx, tableName, nil)
			deleteTable := func() {
				_, err := service.DeleteTable(ctx, tableName, nil)
				if err != nil {
					fmt.Printf("Error cleaning up test. %v\n", err.Error())
				}
			}
			defer deleteTable()

			require.NoError(t, err)
			// require.Equal(t, *resp.TableResponse.TableName, tableName)
		})
	}
}

func TestQueryTable(t *testing.T) {
	for _, svc := range services {
		t.Run(fmt.Sprintf("%v_%v", t.Name(), svc), func(t *testing.T) {
			service, delete := initServiceTest(t, svc)
			defer delete()

			tableCount := 5
			tableNames := make([]string, tableCount)
			prefix1 := "zzza"
			prefix2 := "zzzb"

			defer clearAllTables(service) //nolint
			//create 10 tables with our exected prefix and 1 with a different prefix
			for i := 0; i < tableCount; i++ {
				if i < (tableCount - 1) {
					name := fmt.Sprintf("%v%v", prefix1, i)
					tableNames[i] = name
				} else {
					name := fmt.Sprintf("%v%v", prefix2, i)
					tableNames[i] = name
				}
				_, err := service.CreateTable(ctx, tableNames[i], nil)
				require.NoError(t, err)
			}

			// Query for tables with no pagination. The filter should exclude one table from the results
			filter := fmt.Sprintf("TableName ge '%s' and TableName lt '%s'", prefix1, prefix2)
			pager := service.ListTables(&ListTablesOptions{Filter: &filter})

			resultCount := 0
			for pager.More() {
				resp, err := pager.NextPage(ctx)
				require.NoError(t, err)
				resultCount += len(resp.Tables)
			}

			require.Equal(t, resultCount, tableCount-1)

			// Query for tables with pagination
			top := int32(2)
			pager = service.ListTables(&ListTablesOptions{Filter: &filter, Top: &top})

			resultCount = 0
			pageCount := 0
			for pager.More() {
				resp, err := pager.NextPage(ctx)
				require.NoError(t, err)
				require.LessOrEqual(t, len(resp.Tables), 2)
				resultCount += len(resp.Tables)
				pageCount++
				fmt.Printf("pageCount: %d\tresultCount: %d\n", pageCount, resultCount)
			}

			require.Equal(t, resultCount, tableCount-1)
			if svc == "storage" {
				require.Equal(t, pageCount, int(top))
			}
		})
	}
}

func TestListTables(t *testing.T) {
	for _, service := range services {
		t.Run(fmt.Sprintf("%v_%v", t.Name(), service), func(t *testing.T) {
			service, delete := initServiceTest(t, service)
			defer delete()
			tableName, err := createRandomName(t, tableNamePrefix)
			require.NoError(t, err)

			err = clearAllTables(service)
			require.NoError(t, err)

			for i := 0; i < 5; i++ {
				_, err := service.CreateTable(ctx, fmt.Sprintf("%v%v", tableName, i), nil)
				require.NoError(t, err)
			}

			count := 0
			pager := service.ListTables(nil)
			for pager.More() {
				resp, err := pager.NextPage(ctx)
				require.NoError(t, err)
				count += len(resp.Tables)
			}

			require.Equal(t, 5, count)

			deleteTable := func() {
				for i := 0; i < 5; i++ {
					_, err := service.DeleteTable(ctx, fmt.Sprintf("%v%v", tableName, i), nil)
					if err != nil {
						fmt.Printf("Error cleaning up test. %v\n", err.Error())
					}
				}
			}
			defer deleteTable()

		})
	}
}

// This functionality is only available on storage accounts
func TestGetStatistics(t *testing.T) {
	var cred *SharedKeyCredential
	var err error

	err = recording.Start(t, pathToPackage, nil)
	require.NoError(t, err)
	stop := func() {
		err = recording.Stop(t, nil)
		require.NoError(t, err)
	}
	defer stop()

	accountName := recording.GetEnvVariable("TABLES_STORAGE_ACCOUNT_NAME", "fakeaccount")
	accountKey := recording.GetEnvVariable("TABLES_PRIMARY_STORAGE_ACCOUNT_KEY", "fakeAccountKey")

	if recording.GetRecordMode() == "playback" {
		cred, err = NewSharedKeyCredential("fakeaccount", "fakeAccountKey==")
	} else {
		cred, err = NewSharedKeyCredential(accountName, accountKey)
	}

	serviceURL := storageURI(accountName + "-secondary")
	service, err := createServiceClientForRecording(t, serviceURL, *cred)
	require.NoError(t, err)

	resp, err := service.GetStatistics(ctx, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, resp.GeoReplication.LastSyncTime)
	require.NotNil(t, resp.GeoReplication.Status)
}

// Functionality is only available on storage accounts
func TestGetProperties(t *testing.T) {
	service, delete := initServiceTest(t, "storage")
	defer delete()

	resp, err := service.GetProperties(ctx, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
}

// Logging is only available on storage accounts
func TestSetLogging(t *testing.T) {
	service, delete := initServiceTest(t, "storage")
	defer delete()

	getResp, err := service.GetProperties(ctx, nil)
	require.NoError(t, err)

	getResp.Properties.Logging = &Logging{
		Read:    to.Ptr(true),
		Write:   to.Ptr(true),
		Delete:  to.Ptr(true),
		Version: to.Ptr("1.0"),
		RetentionPolicy: &RetentionPolicy{
			Enabled: to.Ptr(true),
			Days:    to.Ptr(int32(5)),
		},
	}

	resp, err := service.SetProperties(ctx, getResp.Properties, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	recording.Sleep(time.Second * 45)

	received, err := service.GetProperties(ctx, nil)
	require.NoError(t, err)

	require.Equal(t, *getResp.Properties.Logging.Read, *received.Properties.Logging.Read)
	require.Equal(t, *getResp.Properties.Logging.Write, *received.Properties.Logging.Write)
	require.Equal(t, *getResp.Properties.Logging.Delete, *received.Properties.Logging.Delete)
	require.Equal(t, *getResp.Properties.Logging.RetentionPolicy.Enabled, *received.Properties.Logging.RetentionPolicy.Enabled)
	require.Equal(t, *getResp.Properties.Logging.RetentionPolicy.Days, *received.Properties.Logging.RetentionPolicy.Days)
}

func TestSetHoursMetrics(t *testing.T) {
	service, delete := initServiceTest(t, "storage")
	defer delete()

	getResp, err := service.GetProperties(ctx, nil)
	require.NoError(t, err)

	getResp.Properties.HourMetrics = &Metrics{
		Enabled:     to.Ptr(true),
		IncludeAPIs: to.Ptr(true),
		RetentionPolicy: &RetentionPolicy{
			Enabled: to.Ptr(true),
			Days:    to.Ptr(int32(5)),
		},
		Version: to.Ptr("1.0"),
	}

	resp, err := service.SetProperties(ctx, getResp.Properties, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	recording.Sleep(time.Second * 45)

	received, err := service.GetProperties(ctx, nil)
	require.NoError(t, err)

	require.Equal(t, *getResp.Properties.HourMetrics.Enabled, *received.Properties.HourMetrics.Enabled)
	require.Equal(t, *getResp.Properties.HourMetrics.IncludeAPIs, *received.Properties.HourMetrics.IncludeAPIs)
	require.Equal(t, *getResp.Properties.HourMetrics.RetentionPolicy.Days, *received.Properties.HourMetrics.RetentionPolicy.Days)
	require.Equal(t, *getResp.Properties.HourMetrics.RetentionPolicy.Enabled, *received.Properties.HourMetrics.RetentionPolicy.Enabled)
}

func TestSetMinuteMetrics(t *testing.T) {
	service, delete := initServiceTest(t, "storage")
	defer delete()

	getResp, err := service.GetProperties(ctx, nil)
	require.NoError(t, err)

	getResp.Properties.MinuteMetrics = &Metrics{
		Enabled:     to.Ptr(true),
		IncludeAPIs: to.Ptr(true),
		RetentionPolicy: &RetentionPolicy{
			Enabled: to.Ptr(true),
			Days:    to.Ptr(int32(5)),
		},
		Version: to.Ptr("1.0"),
	}

	resp, err := service.SetProperties(ctx, getResp.Properties, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	recording.Sleep(time.Second * 45)

	received, err := service.GetProperties(ctx, nil)
	require.NoError(t, err)

	require.Equal(t, *getResp.Properties.MinuteMetrics.Enabled, *received.Properties.MinuteMetrics.Enabled)
	require.Equal(t, *getResp.Properties.MinuteMetrics.IncludeAPIs, *received.Properties.MinuteMetrics.IncludeAPIs)
	require.Equal(t, *getResp.Properties.MinuteMetrics.RetentionPolicy.Days, *received.Properties.MinuteMetrics.RetentionPolicy.Days)
	require.Equal(t, *getResp.Properties.MinuteMetrics.RetentionPolicy.Enabled, *received.Properties.MinuteMetrics.RetentionPolicy.Enabled)
}

func TestSetCors(t *testing.T) {
	service, delete := initServiceTest(t, "storage")
	defer delete()

	getResp, err := service.GetProperties(ctx, nil)
	require.NoError(t, err)

	getResp.Properties.Cors = []*CorsRule{
		&CorsRule{
			AllowedHeaders:  to.Ptr("x-ms-meta-data*"),
			AllowedMethods:  to.Ptr("PUT"),
			AllowedOrigins:  to.Ptr("www.xyz.com"),
			ExposedHeaders:  to.Ptr("x-ms-meta-source*"),
			MaxAgeInSeconds: to.Ptr(int32(500)),
		},
	}

	resp, err := service.SetProperties(ctx, getResp.Properties, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	recording.Sleep(time.Second * 45)

	received, err := service.GetProperties(ctx, nil)
	require.NoError(t, err)

	require.Equal(t, *getResp.Properties.Cors[0].AllowedHeaders, *received.Properties.Cors[0].AllowedHeaders)
	require.Equal(t, *getResp.Properties.Cors[0].AllowedMethods, *received.Properties.Cors[0].AllowedMethods)
	require.Equal(t, *getResp.Properties.Cors[0].AllowedOrigins, *received.Properties.Cors[0].AllowedOrigins)
	require.Equal(t, *getResp.Properties.Cors[0].ExposedHeaders, *received.Properties.Cors[0].ExposedHeaders)
	require.Equal(t, *getResp.Properties.Cors[0].MaxAgeInSeconds, *received.Properties.Cors[0].MaxAgeInSeconds)
}

func TestSetTooManyCors(t *testing.T) {
	service, delete := initServiceTest(t, "storage")
	defer delete()

	corsRules1 := CorsRule{
		AllowedHeaders:  to.Ptr("x-ms-meta-data*"),
		AllowedMethods:  to.Ptr("PUT"),
		AllowedOrigins:  to.Ptr("www.xyz.com"),
		ExposedHeaders:  to.Ptr("x-ms-meta-source*"),
		MaxAgeInSeconds: to.Ptr(int32(500)),
	}
	props := ServiceProperties{Cors: make([]*CorsRule, 0)}
	for i := 0; i < 6; i++ {
		props.Cors = append(props.Cors, &corsRules1)
	}

	_, err := service.SetProperties(ctx, props, nil)
	require.Error(t, err)
}

func TestRetentionTooLong(t *testing.T) {
	service, delete := initServiceTest(t, "storage")
	defer delete()

	metrics := Metrics{
		Enabled:     to.Ptr(true),
		IncludeAPIs: to.Ptr(true),
		RetentionPolicy: &RetentionPolicy{
			Enabled: to.Ptr(true),
			Days:    to.Ptr(int32(366)),
		},
		Version: to.Ptr("1.0"),
	}
	props := ServiceProperties{MinuteMetrics: &metrics}

	_, err := service.SetProperties(ctx, props, nil)
	require.Error(t, err)
}

func TestGetAccountSASToken(t *testing.T) {
	cred, err := NewSharedKeyCredential("myAccountName", "daaaaaaaaaabbbbbbbbbbcccccccccccccccccccdddddddddddddddddddeeeeeeeeeeefffffffffffggggg==")
	require.NoError(t, err)
	service, err := NewServiceClientWithSharedKey("https://myAccountName.table.core.windows.net", cred, nil)
	require.NoError(t, err)

	resources := AccountSASResourceTypes{Service: true}
	perms := AccountSASPermissions{Read: true}
	start := time.Date(2021, time.September, 8, 14, 30, 0, 0, time.UTC)
	end := start.AddDate(0, 0, 1)

	sas, err := service.GetAccountSASToken(resources, perms, start, end)
	require.NoError(t, err)
	require.Equal(t, "https://myAccountName.table.core.windows.net/?se=2021-09-09T14%3A30%3A00Z&sig=m%2F%2FxhMvxidHaswzZRpyuiHykqnTppPi%2BQ9S5xHMksIQ%3D&sp=r&spr=https&srt=s&ss=t&st=2021-09-08T14%3A30%3A00Z&sv=2019-02-02", sas)
}

func TestGetAccountSASTokenError(t *testing.T) {
	cred := NewFakeCredential("fakeaccount", "fakekey")
	service, err := NewServiceClient("https://myAccountName.table.core.windows.net", cred, nil)
	require.NoError(t, err)

	resources := AccountSASResourceTypes{Service: true}
	perms := AccountSASPermissions{Read: true}

	_, err = service.GetAccountSASToken(resources, perms, time.Now(), time.Now().Add(time.Hour))
	require.Error(t, err)
}
