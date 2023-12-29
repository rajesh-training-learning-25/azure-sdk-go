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

// ServerFactory is a fake server for instances of the armstorage.ClientFactory type.
type ServerFactory struct {
	AccountsServer                   AccountsServer
	BlobContainersServer             BlobContainersServer
	BlobInventoryPoliciesServer      BlobInventoryPoliciesServer
	BlobServicesServer               BlobServicesServer
	DeletedAccountsServer            DeletedAccountsServer
	EncryptionScopesServer           EncryptionScopesServer
	FileServicesServer               FileServicesServer
	FileSharesServer                 FileSharesServer
	LocalUsersServer                 LocalUsersServer
	ManagementPoliciesServer         ManagementPoliciesServer
	ObjectReplicationPoliciesServer  ObjectReplicationPoliciesServer
	OperationsServer                 OperationsServer
	PrivateEndpointConnectionsServer PrivateEndpointConnectionsServer
	PrivateLinkResourcesServer       PrivateLinkResourcesServer
	QueueServer                      QueueServer
	QueueServicesServer              QueueServicesServer
	SKUsServer                       SKUsServer
	TableServer                      TableServer
	TableServicesServer              TableServicesServer
	UsagesServer                     UsagesServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armstorage.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armstorage.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                *ServerFactory
	trMu                               sync.Mutex
	trAccountsServer                   *AccountsServerTransport
	trBlobContainersServer             *BlobContainersServerTransport
	trBlobInventoryPoliciesServer      *BlobInventoryPoliciesServerTransport
	trBlobServicesServer               *BlobServicesServerTransport
	trDeletedAccountsServer            *DeletedAccountsServerTransport
	trEncryptionScopesServer           *EncryptionScopesServerTransport
	trFileServicesServer               *FileServicesServerTransport
	trFileSharesServer                 *FileSharesServerTransport
	trLocalUsersServer                 *LocalUsersServerTransport
	trManagementPoliciesServer         *ManagementPoliciesServerTransport
	trObjectReplicationPoliciesServer  *ObjectReplicationPoliciesServerTransport
	trOperationsServer                 *OperationsServerTransport
	trPrivateEndpointConnectionsServer *PrivateEndpointConnectionsServerTransport
	trPrivateLinkResourcesServer       *PrivateLinkResourcesServerTransport
	trQueueServer                      *QueueServerTransport
	trQueueServicesServer              *QueueServicesServerTransport
	trSKUsServer                       *SKUsServerTransport
	trTableServer                      *TableServerTransport
	trTableServicesServer              *TableServicesServerTransport
	trUsagesServer                     *UsagesServerTransport
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
	case "AccountsClient":
		initServer(s, &s.trAccountsServer, func() *AccountsServerTransport { return NewAccountsServerTransport(&s.srv.AccountsServer) })
		resp, err = s.trAccountsServer.Do(req)
	case "BlobContainersClient":
		initServer(s, &s.trBlobContainersServer, func() *BlobContainersServerTransport {
			return NewBlobContainersServerTransport(&s.srv.BlobContainersServer)
		})
		resp, err = s.trBlobContainersServer.Do(req)
	case "BlobInventoryPoliciesClient":
		initServer(s, &s.trBlobInventoryPoliciesServer, func() *BlobInventoryPoliciesServerTransport {
			return NewBlobInventoryPoliciesServerTransport(&s.srv.BlobInventoryPoliciesServer)
		})
		resp, err = s.trBlobInventoryPoliciesServer.Do(req)
	case "BlobServicesClient":
		initServer(s, &s.trBlobServicesServer, func() *BlobServicesServerTransport { return NewBlobServicesServerTransport(&s.srv.BlobServicesServer) })
		resp, err = s.trBlobServicesServer.Do(req)
	case "DeletedAccountsClient":
		initServer(s, &s.trDeletedAccountsServer, func() *DeletedAccountsServerTransport {
			return NewDeletedAccountsServerTransport(&s.srv.DeletedAccountsServer)
		})
		resp, err = s.trDeletedAccountsServer.Do(req)
	case "EncryptionScopesClient":
		initServer(s, &s.trEncryptionScopesServer, func() *EncryptionScopesServerTransport {
			return NewEncryptionScopesServerTransport(&s.srv.EncryptionScopesServer)
		})
		resp, err = s.trEncryptionScopesServer.Do(req)
	case "FileServicesClient":
		initServer(s, &s.trFileServicesServer, func() *FileServicesServerTransport { return NewFileServicesServerTransport(&s.srv.FileServicesServer) })
		resp, err = s.trFileServicesServer.Do(req)
	case "FileSharesClient":
		initServer(s, &s.trFileSharesServer, func() *FileSharesServerTransport { return NewFileSharesServerTransport(&s.srv.FileSharesServer) })
		resp, err = s.trFileSharesServer.Do(req)
	case "LocalUsersClient":
		initServer(s, &s.trLocalUsersServer, func() *LocalUsersServerTransport { return NewLocalUsersServerTransport(&s.srv.LocalUsersServer) })
		resp, err = s.trLocalUsersServer.Do(req)
	case "ManagementPoliciesClient":
		initServer(s, &s.trManagementPoliciesServer, func() *ManagementPoliciesServerTransport {
			return NewManagementPoliciesServerTransport(&s.srv.ManagementPoliciesServer)
		})
		resp, err = s.trManagementPoliciesServer.Do(req)
	case "ObjectReplicationPoliciesClient":
		initServer(s, &s.trObjectReplicationPoliciesServer, func() *ObjectReplicationPoliciesServerTransport {
			return NewObjectReplicationPoliciesServerTransport(&s.srv.ObjectReplicationPoliciesServer)
		})
		resp, err = s.trObjectReplicationPoliciesServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "PrivateEndpointConnectionsClient":
		initServer(s, &s.trPrivateEndpointConnectionsServer, func() *PrivateEndpointConnectionsServerTransport {
			return NewPrivateEndpointConnectionsServerTransport(&s.srv.PrivateEndpointConnectionsServer)
		})
		resp, err = s.trPrivateEndpointConnectionsServer.Do(req)
	case "PrivateLinkResourcesClient":
		initServer(s, &s.trPrivateLinkResourcesServer, func() *PrivateLinkResourcesServerTransport {
			return NewPrivateLinkResourcesServerTransport(&s.srv.PrivateLinkResourcesServer)
		})
		resp, err = s.trPrivateLinkResourcesServer.Do(req)
	case "QueueClient":
		initServer(s, &s.trQueueServer, func() *QueueServerTransport { return NewQueueServerTransport(&s.srv.QueueServer) })
		resp, err = s.trQueueServer.Do(req)
	case "QueueServicesClient":
		initServer(s, &s.trQueueServicesServer, func() *QueueServicesServerTransport {
			return NewQueueServicesServerTransport(&s.srv.QueueServicesServer)
		})
		resp, err = s.trQueueServicesServer.Do(req)
	case "SKUsClient":
		initServer(s, &s.trSKUsServer, func() *SKUsServerTransport { return NewSKUsServerTransport(&s.srv.SKUsServer) })
		resp, err = s.trSKUsServer.Do(req)
	case "TableClient":
		initServer(s, &s.trTableServer, func() *TableServerTransport { return NewTableServerTransport(&s.srv.TableServer) })
		resp, err = s.trTableServer.Do(req)
	case "TableServicesClient":
		initServer(s, &s.trTableServicesServer, func() *TableServicesServerTransport {
			return NewTableServicesServerTransport(&s.srv.TableServicesServer)
		})
		resp, err = s.trTableServicesServer.Do(req)
	case "UsagesClient":
		initServer(s, &s.trUsagesServer, func() *UsagesServerTransport { return NewUsagesServerTransport(&s.srv.UsagesServer) })
		resp, err = s.trUsagesServer.Do(req)
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
