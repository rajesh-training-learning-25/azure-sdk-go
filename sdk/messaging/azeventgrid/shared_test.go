//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azeventgrid_test

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventgrid"
	"github.com/stretchr/testify/require"
)

type testVars struct {
	Key          string
	Endpoint     string
	Topic        string
	Subscription string

	KeyLogPath string
}

func loadEnv() (testVars, error) {
	var missing []string

	get := func(n string) string {
		if v := os.Getenv(n); v == "" {
			missing = append(missing, n)
		}

		return os.Getenv(n)
	}

	tv := testVars{
		Key:          get("EVENTGRID_KEY"),
		Endpoint:     get("EVENTGRID_ENDPOINT"),
		Topic:        get("EVENTGRID_TOPIC"),
		Subscription: get("EVENTGRID_SUBSCRIPTION"),
	}

	if len(missing) > 0 {
		return testVars{}, fmt.Errorf("Missing env variables: %s", strings.Join(missing, ","))
	}

	// Setting this variable will cause the test clients to dump out the pre-master-key
	// for your HTTP connection. This allows you decrypt a packet capture from wireshark.
	//
	// If you want to do this just set SSLKEYLOGFILE_TEST env var to a path on disk and
	// Go will write out the key.
	tv.KeyLogPath = os.Getenv("SSLKEYLOGFILE_TEST")
	return tv, nil
}

type clientWrapper struct {
	*azeventgrid.Client
	TestVars   testVars
	keyLogFile *os.File
}

func (c clientWrapper) cleanup() {
	if c.keyLogFile != nil {
		c.keyLogFile.Close()
	}
}

func newClientWrapperForTest(t *testing.T) clientWrapper {
	cw, err := newClientWrapper()

	if err != nil {
		fmt.Printf("WARNING: Not running live test: %s\n", err)
		t.SkipNow()
	}

	return cw
}

func newClientWrapper() (clientWrapper, error) {
	vars, err := loadEnv()

	if err != nil {
		return clientWrapper{}, err
	}

	var opts *azeventgrid.ClientOptions
	var keyLogWriter *os.File

	if vars.KeyLogPath != "" {
		tmpKeyLogWriter, err := os.OpenFile(vars.KeyLogPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)

		if err != nil {
			panic(err)
		}

		keyLogWriter = tmpKeyLogWriter

		tp := http.DefaultTransport.(*http.Transport).Clone()
		tp.TLSClientConfig = &tls.Config{
			KeyLogWriter: keyLogWriter,
		}

		opts = &azeventgrid.ClientOptions{
			ClientOptions: azcore.ClientOptions{
				Transport: &http.Client{Transport: tp},
			},
		}
	}

	c, err := azeventgrid.NewClientWithSharedKeyCredential(vars.Endpoint, vars.Key, opts)

	if err != nil {
		panic(err)
	}

	return clientWrapper{
		Client:     c,
		TestVars:   vars,
		keyLogFile: keyLogWriter,
	}, nil
}

func requireEqualCloudEvent(t *testing.T, expected *azeventgrid.CloudEvent, actual *azeventgrid.CloudEvent) {
	t.Helper()

	require.NotEmpty(t, actual.ID, "ID is not empty")
	require.NotEmpty(t, actual.SpecVersion, "SpecVersion is not empty")

	expected.ID = actual.ID

	if expected.SpecVersion == nil {
		expected.SpecVersion = actual.SpecVersion
	}

	require.Equal(t, actual, expected)
}
