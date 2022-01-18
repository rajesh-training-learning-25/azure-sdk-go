//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhealthcareapis

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// FhirDestinationsClient contains the methods for the FhirDestinations group.
// Don't use this type directly, use NewFhirDestinationsClient() instead.
type FhirDestinationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewFhirDestinationsClient creates a new instance of FhirDestinationsClient with the specified values.
// subscriptionID - The subscription identifier.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewFhirDestinationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *FhirDestinationsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &FhirDestinationsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// ListByIotConnector - Lists all FHIR destinations for the given IoT Connector
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the service instance.
// workspaceName - The name of workspace resource.
// iotConnectorName - The name of IoT Connector resource.
// options - FhirDestinationsClientListByIotConnectorOptions contains the optional parameters for the FhirDestinationsClient.ListByIotConnector
// method.
func (client *FhirDestinationsClient) ListByIotConnector(resourceGroupName string, workspaceName string, iotConnectorName string, options *FhirDestinationsClientListByIotConnectorOptions) *FhirDestinationsClientListByIotConnectorPager {
	return &FhirDestinationsClientListByIotConnectorPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByIotConnectorCreateRequest(ctx, resourceGroupName, workspaceName, iotConnectorName, options)
		},
		advancer: func(ctx context.Context, resp FhirDestinationsClientListByIotConnectorResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.IotFhirDestinationCollection.NextLink)
		},
	}
}

// listByIotConnectorCreateRequest creates the ListByIotConnector request.
func (client *FhirDestinationsClient) listByIotConnectorCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, iotConnectorName string, options *FhirDestinationsClientListByIotConnectorOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/iotconnectors/{iotConnectorName}/fhirdestinations"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if iotConnectorName == "" {
		return nil, errors.New("parameter iotConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iotConnectorName}", url.PathEscape(iotConnectorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByIotConnectorHandleResponse handles the ListByIotConnector response.
func (client *FhirDestinationsClient) listByIotConnectorHandleResponse(resp *http.Response) (FhirDestinationsClientListByIotConnectorResponse, error) {
	result := FhirDestinationsClientListByIotConnectorResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IotFhirDestinationCollection); err != nil {
		return FhirDestinationsClientListByIotConnectorResponse{}, err
	}
	return result, nil
}
