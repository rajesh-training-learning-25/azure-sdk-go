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

// ServerFactory is a fake server for instances of the armsaas.ClientFactory type.
type ServerFactory struct {
	ApplicationsServer      ApplicationsServer
	Server                  Server
	OperationServer         OperationServer
	OperationsServer        OperationsServer
	ResourcesServer         ResourcesServer
	SubscriptionLevelServer SubscriptionLevelServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armsaas.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armsaas.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                       *ServerFactory
	trMu                      sync.Mutex
	trApplicationsServer      *ApplicationsServerTransport
	trServer                  *ServerTransport
	trOperationServer         *OperationServerTransport
	trOperationsServer        *OperationsServerTransport
	trResourcesServer         *ResourcesServerTransport
	trSubscriptionLevelServer *SubscriptionLevelServerTransport
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
	case "ApplicationsClient":
		initServer(s, &s.trApplicationsServer, func() *ApplicationsServerTransport { return NewApplicationsServerTransport(&s.srv.ApplicationsServer) })
		resp, err = s.trApplicationsServer.Do(req)
	case "Client":
		initServer(s, &s.trServer, func() *ServerTransport { return NewServerTransport(&s.srv.Server) })
		resp, err = s.trServer.Do(req)
	case "OperationClient":
		initServer(s, &s.trOperationServer, func() *OperationServerTransport { return NewOperationServerTransport(&s.srv.OperationServer) })
		resp, err = s.trOperationServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "ResourcesClient":
		initServer(s, &s.trResourcesServer, func() *ResourcesServerTransport { return NewResourcesServerTransport(&s.srv.ResourcesServer) })
		resp, err = s.trResourcesServer.Do(req)
	case "SubscriptionLevelClient":
		initServer(s, &s.trSubscriptionLevelServer, func() *SubscriptionLevelServerTransport {
			return NewSubscriptionLevelServerTransport(&s.srv.SubscriptionLevelServer)
		})
		resp, err = s.trSubscriptionLevelServer.Do(req)
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
