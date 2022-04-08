// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/perf"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
)

type getListSecretsTestOptions struct {
	count int
}

var getListSecretsTestOpts getListSecretsTestOptions = getListSecretsTestOptions{
	count: 100,
}

func registerListSecrets() {
	flag.IntVar(&getListSecretsTestOpts.count, "count", 10, "number of secrets to create")
}

type ListSecretsTest struct {
	perf.PerfTestOptions
	secretName string
	client     *azsecrets.Client
}

// NewListSecretsTest is called once per process
func NewListSecretsTest(ctx context.Context, options perf.PerfTestOptions) (perf.GlobalPerfTest, error) {
	d := &ListSecretsTest{
		PerfTestOptions: options,
		secretName:      "livekvtestlistsecretperfsecret",
	}

	vaultURL, ok := os.LookupEnv("AZURE_KEYVAULT_URL")
	if !ok {
		return nil, fmt.Errorf("the environment variable 'AZURE_KEYVAULT_URL' could not be found")
	}

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		panic(err)
	}

	client, err := azsecret.NewClient(vaultURL, cred, &azsecret.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: options.Transporter,
		},
	})

	for i := 0; i < getListSecretsTestOpts.count; i++ {
		_, err = client.SetSecret(ctx, fmt.Sprintf("%s%d", d.secretName, i), "secret-value", nil)
		if err != nil {
			return nil, err
		}
	}

	d.client = client
	return d, nil
}

func (gct *ListSecretsTest) GlobalCleanup(ctx context.Context) error {
	for i := 0; i < getListSecretsTestOpts.count; i++ {
		poller, err := gct.client.BeginDeleteSecret(ctx, fmt.Sprintf("%s%d", gct.secretName, i), nil)
		if err != nil {
			return err
		}

		_, err = poller.PollUntilDone(ctx, 500*time.Millisecond)
		if err != nil {
			return err
		}

		_, err = gct.client.PurgeDeletedSecret(ctx, gct.secretName, nil)
		if err != nil {
			return err
		}
	}
	return nil
}

type ListSecretsPerfTest struct {
	*ListSecretsTest
	perf.PerfTestOptions
	client     *azsecrets.Client
	secretName string
}

// NewPerfTest is called once per goroutine
func (gct *ListSecretsTest) NewPerfTest(ctx context.Context, options *perf.PerfTestOptions) (perf.PerfTest, error) {
	return &ListSecretsPerfTest{
		ListSecretsTest: gct,
		PerfTestOptions: *options,
		client:          gct.client,
		secretName:      gct.secretName,
	}, nil
}

func (gcpt *ListSecretsPerfTest) Run(ctx context.Context) error {
	pager := gcpt.client.ListPropertiesOfSecrets(nil)
	for pager.More() {
		_, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (*ListSecretsPerfTest) Cleanup(ctx context.Context) error {
	return nil
}
