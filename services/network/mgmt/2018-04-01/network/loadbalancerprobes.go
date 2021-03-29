package network

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// LoadBalancerProbesClient is the network Client
type LoadBalancerProbesClient struct {
	BaseClient
}

// NewLoadBalancerProbesClient creates an instance of the LoadBalancerProbesClient client.
func NewLoadBalancerProbesClient(subscriptionID string) LoadBalancerProbesClient {
	return NewLoadBalancerProbesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLoadBalancerProbesClientWithBaseURI creates an instance of the LoadBalancerProbesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewLoadBalancerProbesClientWithBaseURI(baseURI string, subscriptionID string) LoadBalancerProbesClient {
	return LoadBalancerProbesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets load balancer probe.
// Parameters:
// resourceGroupName - the name of the resource group.
// loadBalancerName - the name of the load balancer.
// probeName - the name of the probe.
func (client LoadBalancerProbesClient) Get(ctx context.Context, resourceGroupName string, loadBalancerName string, probeName string) (result Probe, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LoadBalancerProbesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, loadBalancerName, probeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerProbesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.LoadBalancerProbesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerProbesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client LoadBalancerProbesClient) GetPreparer(ctx context.Context, resourceGroupName string, loadBalancerName string, probeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"loadBalancerName":  autorest.Encode("path", loadBalancerName),
		"probeName":         autorest.Encode("path", probeName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/probes/{probeName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client LoadBalancerProbesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client LoadBalancerProbesClient) GetResponder(resp *http.Response) (result Probe, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all the load balancer probes.
// Parameters:
// resourceGroupName - the name of the resource group.
// loadBalancerName - the name of the load balancer.
func (client LoadBalancerProbesClient) List(ctx context.Context, resourceGroupName string, loadBalancerName string) (result LoadBalancerProbeListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LoadBalancerProbesClient.List")
		defer func() {
			sc := -1
			if result.lbplr.Response.Response != nil {
				sc = result.lbplr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, loadBalancerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerProbesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.lbplr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.LoadBalancerProbesClient", "List", resp, "Failure sending request")
		return
	}

	result.lbplr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerProbesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.lbplr.hasNextLink() && result.lbplr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client LoadBalancerProbesClient) ListPreparer(ctx context.Context, resourceGroupName string, loadBalancerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"loadBalancerName":  autorest.Encode("path", loadBalancerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/probes", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client LoadBalancerProbesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client LoadBalancerProbesClient) ListResponder(resp *http.Response) (result LoadBalancerProbeListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client LoadBalancerProbesClient) listNextResults(ctx context.Context, lastResults LoadBalancerProbeListResult) (result LoadBalancerProbeListResult, err error) {
	req, err := lastResults.loadBalancerProbeListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.LoadBalancerProbesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.LoadBalancerProbesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerProbesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client LoadBalancerProbesClient) ListComplete(ctx context.Context, resourceGroupName string, loadBalancerName string) (result LoadBalancerProbeListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LoadBalancerProbesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, loadBalancerName)
	return
}
