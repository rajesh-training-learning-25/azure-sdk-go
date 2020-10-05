// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azidentity

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/pkg/browser"
)

// InteractiveBrowserCredentialOptions can be used when providing additional credential information, such as a client secret.
// Also use these options to modify the default pipeline behavior through the TokenCredentialOptions.
type InteractiveBrowserCredentialOptions struct {
	// The Azure Active Directory tenant (directory) ID of the service principal.
	TenantID *string
	// The client (application) ID of the service principal.
	ClientID *string
	// The client secret that was generated for the App Registration used to authenticate the client. Only applies for web apps.
	ClientSecret *string
	// The redirect URI used to request the authorization code. Must be the same URI that is configured for the App Registration.
	RedirectURI *string
	// Options allows configuring the management of the requests sent to Azure Active Directory.
	Options *TokenCredentialOptions
}

// InteractiveBrowserCredential enables authentication to Azure Active Directory using an interactive browser to log in.
type InteractiveBrowserCredential struct {
	client *aadIdentityClient
	// Gets the Azure Active Directory tenant (directory) ID of the service principal.
	tenantID string
	// Gets the client (application) ID of the service principal.
	clientID string
	// Gets the client secret that was generated for the App Registration used to authenticate the client.
	clientSecret string
	// The redirect URI used to request the authorization code. Must be the same URI that is configured for the App Registration.
	redirectURI string
}

// NewInteractiveBrowserCredential constructs a new InteractiveBrowserCredential with the details needed to authenticate against Azure Active Directory through an interactive browser window.
// options: allow to configure the management of the requests sent to Azure Active Directory, leave as nil for default behavior.
func NewInteractiveBrowserCredential(options *InteractiveBrowserCredentialOptions) (*InteractiveBrowserCredential, error) {
	var credentialOptions *TokenCredentialOptions
	tenantID := "organizations"
	clientID := developerSignOnClientID
	redirectURI := ""
	var clientSecret string
	if options != nil {
		if options.Options != nil {
			credentialOptions = options.Options
		}
		if options.TenantID != nil {
			tenantID = *options.TenantID
		}
		if options.ClientID != nil {
			clientID = *options.ClientID
		}
		if options.ClientSecret != nil {
			clientSecret = *options.ClientSecret
		}
		if options.RedirectURI != nil {
			redirectURI = *options.RedirectURI
		}
	}
	c, err := newAADIdentityClient(credentialOptions)
	if err != nil {
		return nil, err
	}
	return &InteractiveBrowserCredential{tenantID: tenantID, clientID: clientID, clientSecret: clientSecret, redirectURI: redirectURI, client: c}, nil
}

// GetToken obtains a token from Azure Active Directory using an interactive browser to authenticate.
// ctx: Context used to control the request lifetime.
// opts: TokenRequestOptions contains the list of scopes for which the token will have access.
// Returns an AccessToken which can be used to authenticate service client calls.
func (c *InteractiveBrowserCredential) GetToken(ctx context.Context, opts azcore.TokenRequestOptions) (*azcore.AccessToken, error) {
	tk, err := c.client.authenticateInteractiveBrowser(ctx, c.tenantID, c.clientID, c.clientSecret, c.redirectURI, opts.Scopes)
	if err != nil {
		addGetTokenFailureLogs("Interactive Browser Credential", err, true)
		return nil, err
	}
	logGetTokenSuccess(c, opts)
	return tk, nil
}

// AuthenticationPolicy implements the azcore.Credential interface on InteractiveBrowserCredential and calls the Bearer Token policy
// to get the bearer token.
func (c *InteractiveBrowserCredential) AuthenticationPolicy(options azcore.AuthenticationPolicyOptions) azcore.Policy {
	return newBearerTokenPolicy(c, options)
}

var _ azcore.TokenCredential = (*InteractiveBrowserCredential)(nil)

var authCodeReceiver = func(authorityHost string, tenantID string, clientID string, redirectURI string, scopes []string) (*interactiveConfig, error) {
	return interactiveBrowserLogin(authorityHost, tenantID, clientID, redirectURI, scopes)
}

// interactiveBrowserLogin opens an interactive browser with the specified tenant and client IDs provided then returns the authorization code
// received or an error.
func interactiveBrowserLogin(authorityHost string, tenantID string, clientID string, redirectURL string, scopes []string) (*interactiveConfig, error) {
	const authURLFormat = "%s/%s/oauth2/v2.0/authorize?response_type=code&response_mode=query&client_id=%s&redirect_uri=%s&state=%s&scope=%s&prompt=select_account"
	if tenantID == "" {
		tenantID = "organizations"
	}
	if clientID == "" {
		clientID = developerSignOnClientID
	}
	state := func() string {
		rand.Seed(time.Now().Unix())
		// generate a 20-char random alpha-numeric string
		const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		buff := make([]byte, 20)
		for i := range buff {
			buff[i] = charset[rand.Intn(len(charset))]
		}
		return string(buff)
	}()
	// start local redirect server so login can call us back
	rs := newServer()
	if redirectURL == "" {
		redirectURL = rs.Start(state)
	}
	defer rs.Stop()
	authURL := fmt.Sprintf(authURLFormat, authorityHost, tenantID, clientID, redirectURL, state, strings.Join(scopes, " "))
	fmt.Println(authURL)
	// open browser window so user can select credentials
	err := browser.OpenURL(authURL)
	if err != nil {
		return nil, err
	}
	// now wait until the logic calls us back
	rs.WaitForCallback()

	authCode, err := rs.AuthorizationCode()
	if err != nil {
		return nil, err
	}
	return &interactiveConfig{
		authCode:    authCode,
		redirectURI: redirectURL,
	}, nil
}
