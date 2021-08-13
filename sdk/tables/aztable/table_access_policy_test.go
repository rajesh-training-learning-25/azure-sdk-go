// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package aztable

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	generated "github.com/Azure/azure-sdk-for-go/sdk/tables/aztable/internal"
)

func TestSetEmptyAccessPolicy(t *testing.T) {
	client, delete := initClientTest(t, "storage", true)
	defer delete()

	_, err := client.SetAccessPolicy(ctx, &generated.TableSetAccessPolicyOptions{})
	require.NoError(t, err)
}

func TestSetAccessPolicy(t *testing.T) {
	client, delete := initClientTest(t, "storage", true)
	defer delete()

	start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	expiration := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	permission := "r"
	id := "1"

	signedIdentifiers := make([]*generated.SignedIdentifier, 0)

	signedIdentifiers = append(signedIdentifiers, &generated.SignedIdentifier{
		AccessPolicy: &generated.AccessPolicy{
			Expiry:     &expiration,
			Start:      &start,
			Permission: &permission,
		},
		ID: &id,
	})

	param := generated.TableSetAccessPolicyOptions{
		TableACL: signedIdentifiers,
	}

	_, err := client.SetAccessPolicy(ctx, &param)
	require.NoError(t, err)
}

func TestSetMultipleAccessPolicies(t *testing.T) {
	client, delete := initClientTest(t, "storage", true)
	defer delete()

	id := "empty"

	signedIdentifiers := make([]*generated.SignedIdentifier, 0)
	signedIdentifiers = append(signedIdentifiers, &generated.SignedIdentifier{
		ID: &id,
	})

	permission2 := "r"
	id2 := "partial"

	signedIdentifiers = append(signedIdentifiers, &generated.SignedIdentifier{
		ID: &id2,
		AccessPolicy: &generated.AccessPolicy{
			Permission: &permission2,
		},
	})

	id3 := "full"
	permission3 := "r"
	start := time.Date(2021, 6, 8, 2, 10, 9, 0, time.UTC)
	expiry := time.Date(2021, 6, 8, 2, 10, 9, 0, time.UTC)

	signedIdentifiers = append(signedIdentifiers, &generated.SignedIdentifier{
		ID: &id3,
		AccessPolicy: &generated.AccessPolicy{
			Start:      &start,
			Expiry:     &expiry,
			Permission: &permission3,
		},
	})

	param := generated.TableSetAccessPolicyOptions{
		TableACL: signedIdentifiers,
	}

	_, err := client.SetAccessPolicy(ctx, &param)
	require.NoError(t, err)

	// Make a Get to assert two access policies
	resp, err := client.GetAccessPolicy(ctx)
	require.NoError(t, err)
	require.Equal(t, len(resp.SignedIdentifiers), 3)
}

func TestSetTooManyAccessPolicies(t *testing.T) {
	client, delete := initClientTest(t, "storage", true)
	defer delete()

	start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	expiration := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	permission := "r"
	id := "1"
	signedIdentifiers := make([]*generated.SignedIdentifier, 0)

	for i := 0; i < 6; i++ {
		expiration = time.Date(2024+i, 1, 1, 0, 0, 0, 0, time.UTC)
		id = strconv.Itoa(i)

		signedIdentifiers = append(signedIdentifiers, &generated.SignedIdentifier{
			AccessPolicy: &generated.AccessPolicy{
				Expiry:     &expiration,
				Start:      &start,
				Permission: &permission,
			},
			ID: &id,
		})

	}

	param := generated.TableSetAccessPolicyOptions{TableACL: signedIdentifiers}

	_, err := client.SetAccessPolicy(ctx, &param)
	require.NotNil(t, err, "Set access policy succeeded but should have failed")
	require.Contains(t, err.Error(), errTooManyAccessPoliciesError.Error())
}

func TestSetNullAccessPolicy(t *testing.T) {
	client, delete := initClientTest(t, "storage", true)
	defer delete()

	id := "null"

	signedIdentifiers := make([]*generated.SignedIdentifier, 0)
	signedIdentifiers = append(signedIdentifiers, &generated.SignedIdentifier{
		ID: &id,
	})

	param := generated.TableSetAccessPolicyOptions{
		TableACL: signedIdentifiers,
	}

	_, err := client.SetAccessPolicy(ctx, &param)
	require.NoError(t, err)

	resp, err := client.GetAccessPolicy(ctx)
	require.NoError(t, err)
	require.Equal(t, len(resp.SignedIdentifiers), 1)
}
