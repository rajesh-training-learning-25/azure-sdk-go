// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/perf"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azkeys"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azkeys/crypto"
)

type decryptTestOptions struct{}

var decryptTestOpts decryptTestOptions = decryptTestOptions{}

type DecryptTest struct {
	perf.PerfTestOptions
	keyName      string
	client       *azkeys.Client
	cryptoClient *crypto.Client
	encrypAlg    crypto.EncryptionAlg
	cipherText   []byte
}

// NewDecryptTest is called once per process
func NewDecryptTest(ctx context.Context, options perf.PerfTestOptions) (perf.GlobalPerfTest, error) {
	d := &DecryptTest{
		PerfTestOptions: options,
		keyName:         "livekvtestdecryptperfkey",
		encrypAlg:       crypto.EncryptionAlgRSAOAEP256,
	}

	vaultURL, ok := os.LookupEnv("AZURE_KEYVAULT_URL")
	if !ok {
		return nil, fmt.Errorf("the environment variable 'AZURE_KEYVAULT_URL' could not be found")
	}

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		panic(err)
	}

	client, err := azkeys.NewClient(vaultURL, cred, &azkeys.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: options.Transporter,
		},
	})
	if err != nil {
		return nil, err
	}

	resp, err := client.CreateRSAKey(ctx, d.keyName, nil)
	if err != nil {
		return nil, err
	}

	cryptoClient, err := crypto.NewClient(*resp.ID, cred, &crypto.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: options.Transporter,
		},
	})
	if err != nil {
		return nil, err
	}

	d.cryptoClient = cryptoClient
	d.client = client

	b := make([]byte, 32)
	_, err = rand.Read(b)
	if err != nil {
		return nil, err
	}

	result, err := d.cryptoClient.Encrypt(ctx, d.encrypAlg, b, nil)
	if err != nil {
		return nil, err
	}
	d.cipherText = result.Ciphertext
	return d, nil
}

func (gct *DecryptTest) GlobalCleanup(ctx context.Context) error {
	poller, err := gct.client.BeginDeleteKey(ctx, gct.keyName, nil)
	if err != nil {
		return err
	}

	_, err = poller.PollUntilDone(ctx, 500*time.Millisecond)
	if err != nil {
		return err
	}

	_, err = gct.client.PurgeDeletedKey(ctx, gct.keyName, nil)
	return err
}

type DecryptPerfTest struct {
	*DecryptTest
	perf.PerfTestOptions
	cryptoClient *crypto.Client
	alg          crypto.EncryptionAlg
	cipher       []byte
}

// NewPerfTest is called once per goroutine
func (gct *DecryptTest) NewPerfTest(ctx context.Context, options *perf.PerfTestOptions) (perf.PerfTest, error) {
	return &DecryptPerfTest{
		DecryptTest:     gct,
		PerfTestOptions: *options,
		alg:             gct.encrypAlg,
		cipher:          gct.cipherText,
	}, nil
}

func (gcpt *DecryptPerfTest) Run(ctx context.Context) error {
	_, err := gcpt.cryptoClient.Decrypt(ctx, gcpt.alg, gcpt.cipher, nil)
	return err
}

func (*DecryptPerfTest) Cleanup(ctx context.Context) error {
	return nil
}
