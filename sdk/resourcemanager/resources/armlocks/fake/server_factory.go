//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// ServerFactory is a fake server for instances of the armlocks.ClientFactory type.
type ServerFactory struct {
	AuthorizationOperationsServer AuthorizationOperationsServer
	ManagementLocksServer         ManagementLocksServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armlocks.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armlocks.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                             *ServerFactory
	trMu                            sync.Mutex
	trAuthorizationOperationsServer *AuthorizationOperationsServerTransport
	trManagementLocksServer         *ManagementLocksServerTransport
}

// Do implements the policy.Transporter interface for ServerFactoryTransport.
func (s *ServerFactoryTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	client := method[:strings.Index(method, ".")]
	var resp *http.Response
	var err error

	switch client {
	case "AuthorizationOperationsClient":
		initServer(s, &s.trAuthorizationOperationsServer, func() *AuthorizationOperationsServerTransport {
			return NewAuthorizationOperationsServerTransport(&s.srv.AuthorizationOperationsServer)
		})
		resp, err = s.trAuthorizationOperationsServer.Do(req)
	case "ManagementLocksClient":
		initServer(s, &s.trManagementLocksServer, func() *ManagementLocksServerTransport {
			return NewManagementLocksServerTransport(&s.srv.ManagementLocksServer)
		})
		resp, err = s.trManagementLocksServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func initServer[T any](s *ServerFactoryTransport, dst **T, src func() *T) {
	s.trMu.Lock()
	if *dst == nil {
		*dst = src()
	}
	s.trMu.Unlock()
}
