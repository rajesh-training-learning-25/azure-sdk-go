// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package operation

type Client struct{}

func (client *Client) Update(resourceGroupName string, options *ClientUpdateOptions) (ClientUpdateResponse, error) {

	return ClientUpdateResponse{}, nil
}

type ClientUpdateOptions struct{}

type ClientUpdateResponse struct{}

func (client *Client) BeginUpdate(resourceGroupName string, options *BeginClientUpdateOptions) (*runtime.Poller[BeginClientUpdateResponse], error) {

	return &runtime.Poller[BeginClientUpdateResponse]{}, nil
}

type BeginClientUpdateOptions struct{}

type BeginClientUpdateResponse struct{}
