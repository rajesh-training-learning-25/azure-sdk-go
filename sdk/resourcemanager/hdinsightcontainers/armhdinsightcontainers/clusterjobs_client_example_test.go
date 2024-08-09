//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhdinsightcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsightcontainers/armhdinsightcontainers"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/aec83a5f0ed56da4fd16fa027b9fa27edfa8988b/specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2024-05-01-preview/examples/RunClusterJob.json
func ExampleClusterJobsClient_BeginRunJob() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsightcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClusterJobsClient().BeginRunJob(ctx, "hiloResourcegroup", "clusterpool1", "cluster1", armhdinsightcontainers.ClusterJob{
		Properties: &armhdinsightcontainers.FlinkJobProperties{
			JobType:    to.Ptr(armhdinsightcontainers.JobTypeFlinkJob),
			Action:     to.Ptr(armhdinsightcontainers.ActionSTART),
			EntryClass: to.Ptr("com.microsoft.hilo.flink.job.streaming.SleepJob"),
			FlinkConfiguration: map[string]*string{
				"parallelism":         to.Ptr("1"),
				"savepoint.directory": to.Ptr("abfs://flinkjob@hilosa.dfs.core.windows.net/savepoint"),
			},
			JarName:         to.Ptr("flink-sleep-job-0.0.1-SNAPSHOT.jar"),
			JobJarDirectory: to.Ptr("abfs://flinkjob@hilosa.dfs.core.windows.net/jars"),
			JobName:         to.Ptr("flink-job-name"),
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
	// res.ClusterJob = armhdinsightcontainers.ClusterJob{
	// 	Properties: &armhdinsightcontainers.FlinkJobProperties{
	// 		JobType: to.Ptr(armhdinsightcontainers.JobTypeFlinkJob),
	// 		Action: to.Ptr(armhdinsightcontainers.ActionSTART),
	// 		EntryClass: to.Ptr("com.microsoft.hilo.flink.job.streaming.SleepJob"),
	// 		FlinkConfiguration: map[string]*string{
	// 			"parallelism": to.Ptr("1"),
	// 			"savepoint.directory": to.Ptr("abfs://flinkjob@hilosa.dfs.core.windows.net/savepoint"),
	// 		},
	// 		JarName: to.Ptr("flink-sleep-job-0.0.1-SNAPSHOT.jar"),
	// 		JobJarDirectory: to.Ptr("abfs://flinkjob@hilosa.dfs.core.windows.net/jars"),
	// 		JobName: to.Ptr("flink-job-name"),
	// 		RunID: to.Ptr("job-15a-4322-b32c-ea541845e911"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/aec83a5f0ed56da4fd16fa027b9fa27edfa8988b/specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2024-05-01-preview/examples/ListClusterJobs.json
func ExampleClusterJobsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsightcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClusterJobsClient().NewListPager("hiloResourcegroup", "clusterPool1", "cluster1", &armhdinsightcontainers.ClusterJobsClientListOptions{Filter: nil})
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
		// page.ClusterJobList = armhdinsightcontainers.ClusterJobList{
		// 	Value: []*armhdinsightcontainers.ClusterJob{
		// 		{
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hiloResourcegroup/providers/Microsoft.HDInsight/clusterPools/clusterpool1/clusters/cluster1/jobs/flink-job-1"),
		// 			Properties: &armhdinsightcontainers.FlinkJobProperties{
		// 				JobType: to.Ptr(armhdinsightcontainers.JobTypeFlinkJob),
		// 				ActionResult: to.Ptr("SUCCESS"),
		// 				EntryClass: to.Ptr("com.microsoft.hilo.flink.job.streaming.ExampleJob"),
		// 				FlinkConfiguration: map[string]*string{
		// 					"parallelism": to.Ptr("1"),
		// 					"savepoint.directory": to.Ptr("savepoint-directory"),
		// 				},
		// 				JarName: to.Ptr("job.jar"),
		// 				JobID: to.Ptr("362b911137dfefc2e55784666f4d4253"),
		// 				JobJarDirectory: to.Ptr("jobJarDirectory1"),
		// 				JobName: to.Ptr("flink-job-1"),
		// 				JobOutput: to.Ptr("job-output"),
		// 				RunID: to.Ptr("job-15a-4322-b32c-ea541845e911"),
		// 				Status: to.Ptr("RUNNING"),
		// 			},
		// 	}},
		// }
	}
}
