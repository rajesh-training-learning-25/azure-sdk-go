//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package internal

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

const headerAuthorization = "Authorization"
const bearerHeader = "Bearer "

type KeyVaultChallengePolicy struct {
	// mainResource is the resource to be retrieved using the tenant specified in the credential
	mainResource *ExpiringResource
	cred         azcore.TokenCredential
	transport    policy.Transporter
	scope        *string
	tenantID     *string
}

func NewKeyVaultChallengePolicy(cred azcore.TokenCredential, t policy.Transporter) *KeyVaultChallengePolicy {
	return &KeyVaultChallengePolicy{
		cred:         cred,
		transport:    t,
		mainResource: NewExpiringResource(acquire),
	}
}

func (k *KeyVaultChallengePolicy) Do(req *policy.Request) (*http.Response, error) {
	as := acquiringResourceState{
		p:   k,
		req: req,
	}

	if k.scope == nil || k.tenantID == nil {
		// First request, get both to get the token
		challengeReq, err := k.getChallengeRequest(*req)
		if err != nil {
			return nil, err
		}

		// challengeResp, err := k.transport.Do(challengeReq)
		challengeResp, err := challengeReq.Next()
		if err != nil {
			return nil, err
		}

		err = k.findScopeAndTenant(challengeResp)
		if err != nil {
			return nil, err
		}
	}

	tk, err := k.mainResource.GetResource(as)
	if err != nil {
		return nil, err
	}

	if token, ok := tk.(*azcore.AccessToken); ok {
		req.Raw().Header.Set(
			headerAuthorization,
			fmt.Sprintf("%s %s", bearerHeader, token.Token),
		)
	}

	// try the request
	resp, requestErr := req.Next()
	if requestErr != nil {
		return nil, requestErr
	}

	// If it fails and has a 401, try it with a new token
	if resp.StatusCode == 401 {
		// Force a new token
		k.mainResource.Reset()

		// Check for a new auth policy
		err := k.findScopeAndTenant(resp)

		// Error parsing challenge, doomed to fail. Return
		if err != nil {
			return resp, err
		}

		tk, err := k.mainResource.GetResource(as)
		if err != nil {
			return resp, err
		}

		if token, ok := tk.(*azcore.AccessToken); ok {
			req.Raw().Header.Set(
				headerAuthorization,
				fmt.Sprintf("%s %s", bearerHeader, token.Token),
			)
		} else {
			// tk is not an azcore.AccessToken type, something went wrong and we should return the 401 and accompanying error
			return resp, requestErr
		}

		resp, err = http.DefaultClient.Do(req.Raw())
		if err != nil {
			// A second request failed, return error
			return nil, err
		}
	}

	return resp, err
}

// parses Tenant ID from auth challenge
// https://login.microsoftonline.com/00000000-0000-0000-0000-000000000000
func parseTenant(url string) *string {
	parts := strings.Split(url, "/")
	tenant := parts[3]
	return &tenant
}

// sets the k.scope and k.tenantID from the WWW-Authenticate header
func (k *KeyVaultChallengePolicy) findScopeAndTenant(resp *http.Response) error {
	authHeader := resp.Header.Get("WWW-Authenticate")
	if authHeader == "" {
		return errors.New("response has no WWW-Authenticate header for challenge authentication")
	}

	// Strip down to auth and resource
	// Format is "Bearer authorization=\"<site>\" resource=\"<site>\"" OR
	// "Bearer authorization=\"<site>\" scope=\"<site>\""
	authHeader = strings.ReplaceAll(authHeader, "Bearer ", "")

	parts := strings.Split(authHeader, " ")

	vals := make([]string, 0)
	for _, part := range parts {
		subParts := strings.Split(part, "=")
		if len(subParts) != 2 {
			k.scope = nil
			k.tenantID = nil
			return fmt.Errorf("could not understand WWW-Authenticate header: %s", authHeader)
		}
		url := subParts[1]

		url = strings.ReplaceAll(url, "\"", "")
		vals = append(vals, url)
	}

	if !strings.HasSuffix(vals[1], "/.default") {
		vals[1] += "/.default"
	}

	k.tenantID = parseTenant(vals[0])
	k.scope = &vals[1]

	return nil
}

func (k KeyVaultChallengePolicy) getChallengeRequest(orig policy.Request) (*policy.Request, error) {
	req, err := runtime.NewRequest(orig.Raw().Context(), orig.Raw().Method, orig.Raw().URL.String())
	if err != nil {
		return nil, err
	}

	req.Raw().Header = orig.Raw().Header
	req.Raw().Header.Set("Content-Length", "0")

	copied := orig.Clone(orig.Raw().Context())
	copied.Raw().Body = req.Body()

	return copied, err
}

type acquiringResourceState struct {
	req *policy.Request
	p   *KeyVaultChallengePolicy
}

// acquire acquires or updates the resource; only one
// thread/goroutine at a time ever calls this function
func acquire(state interface{}) (newResource interface{}, newExpiration time.Time, err error) {
	s := state.(acquiringResourceState)
	tk, err := s.p.cred.GetToken(
		s.req.Raw().Context(),
		policy.TokenRequestOptions{
			Scopes:   []string{*s.p.scope},
			TenantID: *s.p.scope,
		},
	)
	if err != nil {
		return nil, time.Time{}, err
	}
	return tk, tk.ExpiresOn, nil
}
