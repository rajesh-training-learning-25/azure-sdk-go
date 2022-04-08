// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/perf"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
)

type getSecretTestOptions struct{}

var getSecretTestOpts getSecretTestOptions = getSecretTestOptions{}

type GetSecretTest struct {
	perf.PerfTestOptions
	secretName string
	client     *azsecrets.Client
}

// NewGetSecretTest is called once per process
func NewGetSecretTest(ctx context.Context, options perf.PerfTestOptions) (perf.GlobalPerfTest, error) {
	d := &GetSecretTest{
		PerfTestOptions: options,
		secretName:      "livekvtestgetsecretperfsecret",
	}

	vaultURL, ok := os.LookupEnv("AZURE_KEYVAULT_URL")
	if !ok {
		return nil, fmt.Errorf("the environment variable 'AZURE_KEYVAULT_URL' could not be found")
	}

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		panic(err)
	}

	client, err := azsecrets.NewClient(vaultURL, cred, &azsecrets.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: options.Transporter,
		},
	})

	_, err = client.SetSecret(ctx, d.secretName, "secret-value", nil)
	if err != nil {
		return nil, err
	}

	d.client = client
	return d, nil
}

func (gct *GetSecretTest) GlobalCleanup(ctx context.Context) error {
	poller, err := gct.client.BeginDeleteSecret(ctx, gct.secretName, nil)
	if err != nil {
		return err
	}

	_, err = poller.PollUntilDone(ctx, 500*time.Millisecond)
	if err != nil {
		return err
	}

	_, err = gct.client.PurgeDeletedSecret(ctx, gct.secretName, nil)
	return err
}

type GetSecretPerfTest struct {
	*GetSecretTest
	perf.PerfTestOptions
	client     *azsecrets.Client
	secretName string
}

// NewPerfTest is called once per goroutine
func (gct *GetSecretTest) NewPerfTest(ctx context.Context, options *perf.PerfTestOptions) (perf.PerfTest, error) {
	return &GetSecretPerfTest{
		GetSecretTest:   gct,
		PerfTestOptions: *options,
		client:          gct.client,
		secretName:      gct.secretName,
	}, nil
}

func (gcpt *GetSecretPerfTest) Run(ctx context.Context) error {
	_, err := gcpt.client.GetSecret(ctx, gcpt.secretName, nil)
	return err
}

func (*GetSecretPerfTest) Cleanup(ctx context.Context) error {
	return nil
}
