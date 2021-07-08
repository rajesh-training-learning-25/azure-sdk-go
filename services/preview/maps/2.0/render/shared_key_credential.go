// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package render

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

type SharedKeyCredential struct {
	SubscriptionKey string
}

func (creds SharedKeyCredential) AuthenticationPolicy(options azcore.AuthenticationPolicyOptions) azcore.Policy {
	return azcore.PolicyFunc(func(req *azcore.Request) (*azcore.Response, error) {
		q := req.URL.Query()
		q.Add("subscription-key", creds.SubscriptionKey)
		req.URL.RawQuery = q.Encode()
		return req.Next()
	})
}
