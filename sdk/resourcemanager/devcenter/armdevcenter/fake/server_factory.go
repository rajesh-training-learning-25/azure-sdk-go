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

// ServerFactory is a fake server for instances of the armdevcenter.ClientFactory type.
type ServerFactory struct {
	AttachedNetworksServer               AttachedNetworksServer
	CatalogDevBoxDefinitionsServer       CatalogDevBoxDefinitionsServer
	CatalogsServer                       CatalogsServer
	CheckNameAvailabilityServer          CheckNameAvailabilityServer
	CustomizationTasksServer             CustomizationTasksServer
	DevBoxDefinitionsServer              DevBoxDefinitionsServer
	DevCentersServer                     DevCentersServer
	EnvironmentDefinitionsServer         EnvironmentDefinitionsServer
	EnvironmentTypesServer               EnvironmentTypesServer
	GalleriesServer                      GalleriesServer
	ImageVersionsServer                  ImageVersionsServer
	ImagesServer                         ImagesServer
	NetworkConnectionsServer             NetworkConnectionsServer
	OperationStatusesServer              OperationStatusesServer
	OperationsServer                     OperationsServer
	PoolsServer                          PoolsServer
	ProjectAllowedEnvironmentTypesServer ProjectAllowedEnvironmentTypesServer
	ProjectEnvironmentTypesServer        ProjectEnvironmentTypesServer
	ProjectsServer                       ProjectsServer
	SKUsServer                           SKUsServer
	SchedulesServer                      SchedulesServer
	UsagesServer                         UsagesServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armdevcenter.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armdevcenter.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                    *ServerFactory
	trMu                                   sync.Mutex
	trAttachedNetworksServer               *AttachedNetworksServerTransport
	trCatalogDevBoxDefinitionsServer       *CatalogDevBoxDefinitionsServerTransport
	trCatalogsServer                       *CatalogsServerTransport
	trCheckNameAvailabilityServer          *CheckNameAvailabilityServerTransport
	trCustomizationTasksServer             *CustomizationTasksServerTransport
	trDevBoxDefinitionsServer              *DevBoxDefinitionsServerTransport
	trDevCentersServer                     *DevCentersServerTransport
	trEnvironmentDefinitionsServer         *EnvironmentDefinitionsServerTransport
	trEnvironmentTypesServer               *EnvironmentTypesServerTransport
	trGalleriesServer                      *GalleriesServerTransport
	trImageVersionsServer                  *ImageVersionsServerTransport
	trImagesServer                         *ImagesServerTransport
	trNetworkConnectionsServer             *NetworkConnectionsServerTransport
	trOperationStatusesServer              *OperationStatusesServerTransport
	trOperationsServer                     *OperationsServerTransport
	trPoolsServer                          *PoolsServerTransport
	trProjectAllowedEnvironmentTypesServer *ProjectAllowedEnvironmentTypesServerTransport
	trProjectEnvironmentTypesServer        *ProjectEnvironmentTypesServerTransport
	trProjectsServer                       *ProjectsServerTransport
	trSKUsServer                           *SKUsServerTransport
	trSchedulesServer                      *SchedulesServerTransport
	trUsagesServer                         *UsagesServerTransport
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
	case "AttachedNetworksClient":
		initServer(s, &s.trAttachedNetworksServer, func() *AttachedNetworksServerTransport {
			return NewAttachedNetworksServerTransport(&s.srv.AttachedNetworksServer)
		})
		resp, err = s.trAttachedNetworksServer.Do(req)
	case "CatalogDevBoxDefinitionsClient":
		initServer(s, &s.trCatalogDevBoxDefinitionsServer, func() *CatalogDevBoxDefinitionsServerTransport {
			return NewCatalogDevBoxDefinitionsServerTransport(&s.srv.CatalogDevBoxDefinitionsServer)
		})
		resp, err = s.trCatalogDevBoxDefinitionsServer.Do(req)
	case "CatalogsClient":
		initServer(s, &s.trCatalogsServer, func() *CatalogsServerTransport { return NewCatalogsServerTransport(&s.srv.CatalogsServer) })
		resp, err = s.trCatalogsServer.Do(req)
	case "CheckNameAvailabilityClient":
		initServer(s, &s.trCheckNameAvailabilityServer, func() *CheckNameAvailabilityServerTransport {
			return NewCheckNameAvailabilityServerTransport(&s.srv.CheckNameAvailabilityServer)
		})
		resp, err = s.trCheckNameAvailabilityServer.Do(req)
	case "CustomizationTasksClient":
		initServer(s, &s.trCustomizationTasksServer, func() *CustomizationTasksServerTransport {
			return NewCustomizationTasksServerTransport(&s.srv.CustomizationTasksServer)
		})
		resp, err = s.trCustomizationTasksServer.Do(req)
	case "DevBoxDefinitionsClient":
		initServer(s, &s.trDevBoxDefinitionsServer, func() *DevBoxDefinitionsServerTransport {
			return NewDevBoxDefinitionsServerTransport(&s.srv.DevBoxDefinitionsServer)
		})
		resp, err = s.trDevBoxDefinitionsServer.Do(req)
	case "DevCentersClient":
		initServer(s, &s.trDevCentersServer, func() *DevCentersServerTransport { return NewDevCentersServerTransport(&s.srv.DevCentersServer) })
		resp, err = s.trDevCentersServer.Do(req)
	case "EnvironmentDefinitionsClient":
		initServer(s, &s.trEnvironmentDefinitionsServer, func() *EnvironmentDefinitionsServerTransport {
			return NewEnvironmentDefinitionsServerTransport(&s.srv.EnvironmentDefinitionsServer)
		})
		resp, err = s.trEnvironmentDefinitionsServer.Do(req)
	case "EnvironmentTypesClient":
		initServer(s, &s.trEnvironmentTypesServer, func() *EnvironmentTypesServerTransport {
			return NewEnvironmentTypesServerTransport(&s.srv.EnvironmentTypesServer)
		})
		resp, err = s.trEnvironmentTypesServer.Do(req)
	case "GalleriesClient":
		initServer(s, &s.trGalleriesServer, func() *GalleriesServerTransport { return NewGalleriesServerTransport(&s.srv.GalleriesServer) })
		resp, err = s.trGalleriesServer.Do(req)
	case "ImageVersionsClient":
		initServer(s, &s.trImageVersionsServer, func() *ImageVersionsServerTransport {
			return NewImageVersionsServerTransport(&s.srv.ImageVersionsServer)
		})
		resp, err = s.trImageVersionsServer.Do(req)
	case "ImagesClient":
		initServer(s, &s.trImagesServer, func() *ImagesServerTransport { return NewImagesServerTransport(&s.srv.ImagesServer) })
		resp, err = s.trImagesServer.Do(req)
	case "NetworkConnectionsClient":
		initServer(s, &s.trNetworkConnectionsServer, func() *NetworkConnectionsServerTransport {
			return NewNetworkConnectionsServerTransport(&s.srv.NetworkConnectionsServer)
		})
		resp, err = s.trNetworkConnectionsServer.Do(req)
	case "OperationStatusesClient":
		initServer(s, &s.trOperationStatusesServer, func() *OperationStatusesServerTransport {
			return NewOperationStatusesServerTransport(&s.srv.OperationStatusesServer)
		})
		resp, err = s.trOperationStatusesServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "PoolsClient":
		initServer(s, &s.trPoolsServer, func() *PoolsServerTransport { return NewPoolsServerTransport(&s.srv.PoolsServer) })
		resp, err = s.trPoolsServer.Do(req)
	case "ProjectAllowedEnvironmentTypesClient":
		initServer(s, &s.trProjectAllowedEnvironmentTypesServer, func() *ProjectAllowedEnvironmentTypesServerTransport {
			return NewProjectAllowedEnvironmentTypesServerTransport(&s.srv.ProjectAllowedEnvironmentTypesServer)
		})
		resp, err = s.trProjectAllowedEnvironmentTypesServer.Do(req)
	case "ProjectEnvironmentTypesClient":
		initServer(s, &s.trProjectEnvironmentTypesServer, func() *ProjectEnvironmentTypesServerTransport {
			return NewProjectEnvironmentTypesServerTransport(&s.srv.ProjectEnvironmentTypesServer)
		})
		resp, err = s.trProjectEnvironmentTypesServer.Do(req)
	case "ProjectsClient":
		initServer(s, &s.trProjectsServer, func() *ProjectsServerTransport { return NewProjectsServerTransport(&s.srv.ProjectsServer) })
		resp, err = s.trProjectsServer.Do(req)
	case "SKUsClient":
		initServer(s, &s.trSKUsServer, func() *SKUsServerTransport { return NewSKUsServerTransport(&s.srv.SKUsServer) })
		resp, err = s.trSKUsServer.Do(req)
	case "SchedulesClient":
		initServer(s, &s.trSchedulesServer, func() *SchedulesServerTransport { return NewSchedulesServerTransport(&s.srv.SchedulesServer) })
		resp, err = s.trSchedulesServer.Do(req)
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
