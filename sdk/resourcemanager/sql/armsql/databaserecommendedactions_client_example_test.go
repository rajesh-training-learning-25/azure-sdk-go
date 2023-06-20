//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fba7ffa9cee6453e2a3cf8c857074a323252a12d/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DatabaseRecommendedActionList.json
func ExampleDatabaseRecommendedActionsClient_ListByDatabaseAdvisor() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDatabaseRecommendedActionsClient().ListByDatabaseAdvisor(ctx, "workloadinsight-demos", "misosisvr", "IndexAdvisor_test_3", "CreateIndex", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RecommendedActionArray = []*armsql.RecommendedAction{
	// 	{
	// 		Name: to.Ptr("IR_[CRM]_[DataPoints]_4821CD2F9510D98184BB"),
	// 		Type: to.Ptr("Microsoft.Sql/servers/databases/advisors/recommendedActions"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workloadinsight-demos/providers/Microsoft.Sql/servers/misosisvr/databases/IndexAdvisor_test_3/advisors/CreateIndex/recommendedActions/IR_[CRM]_[DataPoints]_4821CD2F9510D98184BB"),
	// 		Kind: to.Ptr(""),
	// 		Location: to.Ptr("East Asia"),
	// 		Properties: &armsql.RecommendedActionProperties{
	// 			ErrorDetails: &armsql.RecommendedActionErrorInfo{
	// 			},
	// 			EstimatedImpact: []*armsql.RecommendedActionImpactRecord{
	// 				{
	// 					AbsoluteValue: to.Ptr[float64](1440),
	// 					DimensionName: to.Ptr("ActionDuration"),
	// 					Unit: to.Ptr("Seconds"),
	// 				},
	// 				{
	// 					AbsoluteValue: to.Ptr[float64](209.3125),
	// 					DimensionName: to.Ptr("SpaceChange"),
	// 					Unit: to.Ptr("Megabytes"),
	// 			}},
	// 			ImplementationDetails: &armsql.RecommendedActionImplementationInfo{
	// 				Method: to.Ptr(armsql.ImplementationMethodTSQL),
	// 				Script: to.Ptr("CREATE NONCLUSTERED INDEX [nci_wi_DataPoints_B892614093BAC56295EF6018BD4CB51B] ON [CRM].[DataPoints] ([Name],[Money],[Power]) INCLUDE ([Hour], [System], [LastChanged]) WITH (ONLINE = ON)"),
	// 			},
	// 			IsArchivedAction: to.Ptr(false),
	// 			IsExecutableAction: to.Ptr(true),
	// 			IsRevertableAction: to.Ptr(true),
	// 			LastRefresh: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:04Z"); return t}()),
	// 			ObservedImpact: []*armsql.RecommendedActionImpactRecord{
	// 			},
	// 			RecommendationReason: to.Ptr(""),
	// 			Score: to.Ptr[int32](1),
	// 			State: &armsql.RecommendedActionStateInfo{
	// 				CurrentValue: to.Ptr(armsql.RecommendedActionCurrentStateActive),
	// 				LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:05Z"); return t}()),
	// 			},
	// 			TimeSeries: []*armsql.RecommendedActionMetricInfo{
	// 			},
	// 			ValidSince: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:04Z"); return t}()),
	// 			Details: map[string]any{
	// 				"schema": "[CRM]",
	// 				"includedColumns": "[Hour], [System], [LastChanged]",
	// 				"indexColumns": "[Name],[Money],[Power]",
	// 				"indexName": "nci_wi_DataPoints_B892614093BAC56295EF6018BD4CB51B",
	// 				"indexType": "NONCLUSTERED",
	// 				"table": "[DataPoints]",
	// 			},
	// 		},
	// 	},
	// 	{
	// 		Name: to.Ptr("IR_[dbo]_[DataPoints]_F5D2F347AA22DB46E4CC"),
	// 		Type: to.Ptr("Microsoft.Sql/servers/databases/advisors/recommendedActions"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workloadinsight-demos/providers/Microsoft.Sql/servers/misosisvr/databases/IndexAdvisor_test_3/advisors/CreateIndex/recommendedActions/IR_[dbo]_[DataPoints]_F5D2F347AA22DB46E4CC"),
	// 		Kind: to.Ptr(""),
	// 		Location: to.Ptr("East Asia"),
	// 		Properties: &armsql.RecommendedActionProperties{
	// 			ErrorDetails: &armsql.RecommendedActionErrorInfo{
	// 			},
	// 			EstimatedImpact: []*armsql.RecommendedActionImpactRecord{
	// 				{
	// 					AbsoluteValue: to.Ptr[float64](5040),
	// 					DimensionName: to.Ptr("ActionDuration"),
	// 					Unit: to.Ptr("Seconds"),
	// 				},
	// 				{
	// 					AbsoluteValue: to.Ptr[float64](120),
	// 					DimensionName: to.Ptr("SpaceChange"),
	// 					Unit: to.Ptr("Megabytes"),
	// 			}},
	// 			ExecuteActionDuration: to.Ptr("PT1M"),
	// 			ExecuteActionInitiatedBy: to.Ptr(armsql.RecommendedActionInitiatedByUser),
	// 			ExecuteActionInitiatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:05Z"); return t}()),
	// 			ExecuteActionStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:05Z"); return t}()),
	// 			ImplementationDetails: &armsql.RecommendedActionImplementationInfo{
	// 				Method: to.Ptr(armsql.ImplementationMethodTSQL),
	// 				Script: to.Ptr("DROP INDEX [nci_wi_DataPoints_609E4B7D6A3813990ED44B28B340C8FC] ON [dbo].[DataPoints]"),
	// 			},
	// 			IsArchivedAction: to.Ptr(false),
	// 			IsExecutableAction: to.Ptr(true),
	// 			IsRevertableAction: to.Ptr(true),
	// 			LastRefresh: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:05Z"); return t}()),
	// 			ObservedImpact: []*armsql.RecommendedActionImpactRecord{
	// 				{
	// 					ChangeValueAbsolute: to.Ptr[float64](-12.7),
	// 					ChangeValueRelative: to.Ptr[float64](-0.9),
	// 					DimensionName: to.Ptr("AffectedQueriesCpuUtilization"),
	// 					Unit: to.Ptr("CpuCores"),
	// 				},
	// 				{
	// 					ChangeValueAbsolute: to.Ptr[float64](-12.7),
	// 					ChangeValueRelative: to.Ptr[float64](-0.3175),
	// 					DimensionName: to.Ptr("CpuUtilization"),
	// 					Unit: to.Ptr("CpuCores"),
	// 				},
	// 				{
	// 					AbsoluteValue: to.Ptr[float64](12),
	// 					DimensionName: to.Ptr("QueriesWithImprovedPerformance"),
	// 					Unit: to.Ptr("Count"),
	// 				},
	// 				{
	// 					AbsoluteValue: to.Ptr[float64](1),
	// 					DimensionName: to.Ptr("QueriesWithRegressedPerformance"),
	// 					Unit: to.Ptr("Count"),
	// 				},
	// 				{
	// 					AbsoluteValue: to.Ptr[float64](130.742187),
	// 					DimensionName: to.Ptr("SpaceChange"),
	// 					Unit: to.Ptr("Megabytes"),
	// 				},
	// 				{
	// 					AbsoluteValue: to.Ptr[float64](0),
	// 					DimensionName: to.Ptr("VerificationProgress"),
	// 					Unit: to.Ptr("Percent"),
	// 			}},
	// 			RecommendationReason: to.Ptr(""),
	// 			Score: to.Ptr[int32](3),
	// 			State: &armsql.RecommendedActionStateInfo{
	// 				ActionInitiatedBy: to.Ptr(armsql.RecommendedActionInitiatedByUser),
	// 				CurrentValue: to.Ptr(armsql.RecommendedActionCurrentStateSuccess),
	// 				LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:05Z"); return t}()),
	// 			},
	// 			TimeSeries: []*armsql.RecommendedActionMetricInfo{
	// 			},
	// 			ValidSince: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:05Z"); return t}()),
	// 			Details: map[string]any{
	// 				"schema": "[dbo]",
	// 				"includedColumns": "[Power],[Pineapple]",
	// 				"indexActionDuration": "00:01:00",
	// 				"indexActionStartTime": "2017-03-01T14:38:05.337",
	// 				"indexColumns": "[Name],[Money]",
	// 				"indexName": "nci_wi_DataPoints_609E4B7D6A3813990ED44B28B340C8FC",
	// 				"indexType": "NONCLUSTERED",
	// 				"table": "[DataPoints]",
	// 			},
	// 		},
	// 	},
	// 	{
	// 		Name: to.Ptr("IR_[dbo]_[Employees]_560E15A98D14CA09BDFB"),
	// 		Type: to.Ptr("Microsoft.Sql/servers/databases/advisors/recommendedActions"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workloadinsight-demos/providers/Microsoft.Sql/servers/misosisvr/databases/IndexAdvisor_test_3/advisors/CreateIndex/recommendedActions/IR_[dbo]_[Employees]_560E15A98D14CA09BDFB"),
	// 		Kind: to.Ptr(""),
	// 		Location: to.Ptr("East Asia"),
	// 		Properties: &armsql.RecommendedActionProperties{
	// 			ErrorDetails: &armsql.RecommendedActionErrorInfo{
	// 			},
	// 			EstimatedImpact: []*armsql.RecommendedActionImpactRecord{
	// 				{
	// 					AbsoluteValue: to.Ptr[float64](17),
	// 					DimensionName: to.Ptr("ActionDuration"),
	// 					Unit: to.Ptr("Seconds"),
	// 				},
	// 				{
	// 					AbsoluteValue: to.Ptr[float64](128),
	// 					DimensionName: to.Ptr("SpaceChange"),
	// 					Unit: to.Ptr("Megabytes"),
	// 			}},
	// 			ImplementationDetails: &armsql.RecommendedActionImplementationInfo{
	// 				Method: to.Ptr(armsql.ImplementationMethodTSQL),
	// 				Script: to.Ptr("CREATE NONCLUSTERED INDEX [nci_wi_Employees_8C18C2AF4267DC77793040782641CCDE] ON [dbo].[Employees] ([City], [State]) INCLUDE ([Postal]) WITH (ONLINE = ON)"),
	// 			},
	// 			IsArchivedAction: to.Ptr(false),
	// 			IsExecutableAction: to.Ptr(true),
	// 			IsRevertableAction: to.Ptr(true),
	// 			LastRefresh: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:05Z"); return t}()),
	// 			ObservedImpact: []*armsql.RecommendedActionImpactRecord{
	// 			},
	// 			RecommendationReason: to.Ptr(""),
	// 			Score: to.Ptr[int32](3),
	// 			State: &armsql.RecommendedActionStateInfo{
	// 				CurrentValue: to.Ptr(armsql.RecommendedActionCurrentStateActive),
	// 				LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:05Z"); return t}()),
	// 			},
	// 			TimeSeries: []*armsql.RecommendedActionMetricInfo{
	// 			},
	// 			ValidSince: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:05Z"); return t}()),
	// 			Details: map[string]any{
	// 				"schema": "[dbo]",
	// 				"includedColumns": "[Postal]",
	// 				"indexColumns": "[City], [State]",
	// 				"indexName": "nci_wi_Employees_8C18C2AF4267DC77793040782641CCDE",
	// 				"indexType": "NONCLUSTERED",
	// 				"table": "[Employees]",
	// 			},
	// 		},
	// }}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fba7ffa9cee6453e2a3cf8c857074a323252a12d/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DatabaseRecommendedActionGet.json
func ExampleDatabaseRecommendedActionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDatabaseRecommendedActionsClient().Get(ctx, "workloadinsight-demos", "misosisvr", "IndexAdvisor_test_3", "CreateIndex", "IR_[CRM]_[DataPoints]_4821CD2F9510D98184BB", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RecommendedAction = armsql.RecommendedAction{
	// 	Name: to.Ptr("IR_[CRM]_[DataPoints]_4821CD2F9510D98184BB"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/advisors/recommendedActions"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workloadinsight-demos/providers/Microsoft.Sql/servers/misosisvr/databases/IndexAdvisor_test_3/advisors/CreateIndex/recommendedActions/IR_[CRM]_[DataPoints]_4821CD2F9510D98184BB"),
	// 	Kind: to.Ptr(""),
	// 	Location: to.Ptr("East Asia"),
	// 	Properties: &armsql.RecommendedActionProperties{
	// 		ErrorDetails: &armsql.RecommendedActionErrorInfo{
	// 		},
	// 		EstimatedImpact: []*armsql.RecommendedActionImpactRecord{
	// 			{
	// 				AbsoluteValue: to.Ptr[float64](1440),
	// 				DimensionName: to.Ptr("ActionDuration"),
	// 				Unit: to.Ptr("Seconds"),
	// 			},
	// 			{
	// 				AbsoluteValue: to.Ptr[float64](209.3125),
	// 				DimensionName: to.Ptr("SpaceChange"),
	// 				Unit: to.Ptr("Megabytes"),
	// 		}},
	// 		ImplementationDetails: &armsql.RecommendedActionImplementationInfo{
	// 			Method: to.Ptr(armsql.ImplementationMethodTSQL),
	// 			Script: to.Ptr("CREATE NONCLUSTERED INDEX [nci_wi_DataPoints_B892614093BAC56295EF6018BD4CB51B] ON [CRM].[DataPoints] ([Name],[Money],[Power]) INCLUDE ([Hour], [System], [LastChanged]) WITH (ONLINE = ON)"),
	// 		},
	// 		IsArchivedAction: to.Ptr(false),
	// 		IsExecutableAction: to.Ptr(true),
	// 		IsRevertableAction: to.Ptr(true),
	// 		LastRefresh: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:04Z"); return t}()),
	// 		ObservedImpact: []*armsql.RecommendedActionImpactRecord{
	// 		},
	// 		RecommendationReason: to.Ptr(""),
	// 		Score: to.Ptr[int32](1),
	// 		State: &armsql.RecommendedActionStateInfo{
	// 			CurrentValue: to.Ptr(armsql.RecommendedActionCurrentStateActive),
	// 			LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:05Z"); return t}()),
	// 		},
	// 		TimeSeries: []*armsql.RecommendedActionMetricInfo{
	// 		},
	// 		ValidSince: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:04Z"); return t}()),
	// 		Details: map[string]any{
	// 			"schema": "[CRM]",
	// 			"includedColumns": "[Hour], [System], [LastChanged]",
	// 			"indexColumns": "[Name],[Money],[Power]",
	// 			"indexName": "nci_wi_DataPoints_B892614093BAC56295EF6018BD4CB51B",
	// 			"indexType": "NONCLUSTERED",
	// 			"table": "[DataPoints]",
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fba7ffa9cee6453e2a3cf8c857074a323252a12d/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DatabaseRecommendedActionUpdate.json
func ExampleDatabaseRecommendedActionsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDatabaseRecommendedActionsClient().Update(ctx, "workloadinsight-demos", "misosisvr", "IndexAdvisor_test_3", "CreateIndex", "IR_[CRM]_[DataPoints]_4821CD2F9510D98184BB", armsql.RecommendedAction{
		Properties: &armsql.RecommendedActionProperties{
			State: &armsql.RecommendedActionStateInfo{
				CurrentValue: to.Ptr(armsql.RecommendedActionCurrentStatePending),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RecommendedAction = armsql.RecommendedAction{
	// 	Name: to.Ptr("IR_[CRM]_[DataPoints]_4821CD2F9510D98184BB"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/advisors/recommendedActions"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workloadinsight-demos/providers/Microsoft.Sql/servers/misosisvr/databases/IndexAdvisor_test_3/advisors/CreateIndex/recommendedActions/IR_[CRM]_[DataPoints]_4821CD2F9510D98184BB"),
	// 	Kind: to.Ptr(""),
	// 	Location: to.Ptr("East Asia"),
	// 	Properties: &armsql.RecommendedActionProperties{
	// 		ErrorDetails: &armsql.RecommendedActionErrorInfo{
	// 		},
	// 		EstimatedImpact: []*armsql.RecommendedActionImpactRecord{
	// 			{
	// 				AbsoluteValue: to.Ptr[float64](1440),
	// 				DimensionName: to.Ptr("ActionDuration"),
	// 				Unit: to.Ptr("Seconds"),
	// 			},
	// 			{
	// 				AbsoluteValue: to.Ptr[float64](209.3125),
	// 				DimensionName: to.Ptr("SpaceChange"),
	// 				Unit: to.Ptr("Megabytes"),
	// 		}},
	// 		ExecuteActionInitiatedBy: to.Ptr(armsql.RecommendedActionInitiatedByUser),
	// 		ExecuteActionInitiatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-20T15:11:15Z"); return t}()),
	// 		ImplementationDetails: &armsql.RecommendedActionImplementationInfo{
	// 			Method: to.Ptr(armsql.ImplementationMethodTSQL),
	// 			Script: to.Ptr("CREATE NONCLUSTERED INDEX [nci_wi_DataPoints_B892614093BAC56295EF6018BD4CB51B] ON [CRM].[DataPoints] ([Name],[Money],[Power]) INCLUDE ([Hour], [System], [LastChanged]) WITH (ONLINE = ON)"),
	// 		},
	// 		IsArchivedAction: to.Ptr(false),
	// 		IsExecutableAction: to.Ptr(true),
	// 		IsRevertableAction: to.Ptr(true),
	// 		LastRefresh: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:04Z"); return t}()),
	// 		ObservedImpact: []*armsql.RecommendedActionImpactRecord{
	// 		},
	// 		RecommendationReason: to.Ptr(""),
	// 		Score: to.Ptr[int32](1),
	// 		State: &armsql.RecommendedActionStateInfo{
	// 			ActionInitiatedBy: to.Ptr(armsql.RecommendedActionInitiatedByUser),
	// 			CurrentValue: to.Ptr(armsql.RecommendedActionCurrentStatePending),
	// 			LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-20T15:11:15Z"); return t}()),
	// 		},
	// 		TimeSeries: []*armsql.RecommendedActionMetricInfo{
	// 		},
	// 		ValidSince: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:04Z"); return t}()),
	// 		Details: map[string]any{
	// 			"schema": "[CRM]",
	// 			"includedColumns": "[Hour], [System], [LastChanged]",
	// 			"indexColumns": "[Name],[Money],[Power]",
	// 			"indexName": "nci_wi_DataPoints_B892614093BAC56295EF6018BD4CB51B",
	// 			"indexType": "NONCLUSTERED",
	// 			"table": "[DataPoints]",
	// 		},
	// 	},
	// }
}
