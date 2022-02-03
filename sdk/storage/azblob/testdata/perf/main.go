// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package main

import "github.com/Azure/azure-sdk-for-go/sdk/internal/perf"

func main() {
	perf.RegisterArguments(RegisterArguments)
	perf.Run([]perf.NewPerfTest{
		NewDownloadTest,
		NewListTest,
		NewUploadTest,
	})
}
