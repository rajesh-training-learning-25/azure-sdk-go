package containerinstance

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

// ContainersClient is the client for the Containers methods of the Containerinstance service.
type ContainersClient struct {
	BaseClient
}

// NewContainersClient creates an instance of the ContainersClient client.
func NewContainersClient(subscriptionID string) ContainersClient {
	return NewContainersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewContainersClientWithBaseURI creates an instance of the ContainersClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewContainersClientWithBaseURI(baseURI string, subscriptionID string) ContainersClient {
	return ContainersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Attach attach to the output stream of a specific container instance in a specified resource group and container
// group.
// Parameters:
// resourceGroupName - the name of the resource group.
// containerGroupName - the name of the container group.
// containerName - the name of the container instance.
func (client ContainersClient) Attach(ctx context.Context, resourceGroupName string, containerGroupName string, containerName string) (result ContainerAttachResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ContainersClient.Attach")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.AttachPreparer(ctx, resourceGroupName, containerGroupName, containerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.ContainersClient", "Attach", nil, "Failure preparing request")
		return
	}

	resp, err := client.AttachSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerinstance.ContainersClient", "Attach", resp, "Failure sending request")
		return
	}

	result, err = client.AttachResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.ContainersClient", "Attach", resp, "Failure responding to request")
		return
	}

	return
}

// AttachPreparer prepares the Attach request.
func (client ContainersClient) AttachPreparer(ctx context.Context, resourceGroupName string, containerGroupName string, containerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerGroupName": autorest.Encode("path", containerGroupName),
		"containerName":      autorest.Encode("path", containerName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}/containers/{containerName}/attach", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AttachSender sends the Attach request. The method will close the
// http.Response Body if it receives an error.
func (client ContainersClient) AttachSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// AttachResponder handles the response to the Attach request. The method always
// closes the http.Response Body.
func (client ContainersClient) AttachResponder(resp *http.Response) (result ContainerAttachResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ExecuteCommand executes a command for a specific container instance in a specified resource group and container
// group.
// Parameters:
// resourceGroupName - the name of the resource group.
// containerGroupName - the name of the container group.
// containerName - the name of the container instance.
// containerExecRequest - the request for the exec command.
func (client ContainersClient) ExecuteCommand(ctx context.Context, resourceGroupName string, containerGroupName string, containerName string, containerExecRequest ContainerExecRequest) (result ContainerExecResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ContainersClient.ExecuteCommand")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ExecuteCommandPreparer(ctx, resourceGroupName, containerGroupName, containerName, containerExecRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.ContainersClient", "ExecuteCommand", nil, "Failure preparing request")
		return
	}

	resp, err := client.ExecuteCommandSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerinstance.ContainersClient", "ExecuteCommand", resp, "Failure sending request")
		return
	}

	result, err = client.ExecuteCommandResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.ContainersClient", "ExecuteCommand", resp, "Failure responding to request")
		return
	}

	return
}

// ExecuteCommandPreparer prepares the ExecuteCommand request.
func (client ContainersClient) ExecuteCommandPreparer(ctx context.Context, resourceGroupName string, containerGroupName string, containerName string, containerExecRequest ContainerExecRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerGroupName": autorest.Encode("path", containerGroupName),
		"containerName":      autorest.Encode("path", containerName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}/containers/{containerName}/exec", pathParameters),
		autorest.WithJSON(containerExecRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ExecuteCommandSender sends the ExecuteCommand request. The method will close the
// http.Response Body if it receives an error.
func (client ContainersClient) ExecuteCommandSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ExecuteCommandResponder handles the response to the ExecuteCommand request. The method always
// closes the http.Response Body.
func (client ContainersClient) ExecuteCommandResponder(resp *http.Response) (result ContainerExecResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListLogs get the logs for a specified container instance in a specified resource group and container group.
// Parameters:
// resourceGroupName - the name of the resource group.
// containerGroupName - the name of the container group.
// containerName - the name of the container instance.
// tail - the number of lines to show from the tail of the container instance log. If not provided, all
// available logs are shown up to 4mb.
// timestamps - if true, adds a timestamp at the beginning of every line of log output. If not provided,
// defaults to false.
func (client ContainersClient) ListLogs(ctx context.Context, resourceGroupName string, containerGroupName string, containerName string, tail *int32, timestamps *bool) (result Logs, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ContainersClient.ListLogs")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListLogsPreparer(ctx, resourceGroupName, containerGroupName, containerName, tail, timestamps)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.ContainersClient", "ListLogs", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListLogsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerinstance.ContainersClient", "ListLogs", resp, "Failure sending request")
		return
	}

	result, err = client.ListLogsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.ContainersClient", "ListLogs", resp, "Failure responding to request")
		return
	}

	return
}

// ListLogsPreparer prepares the ListLogs request.
func (client ContainersClient) ListLogsPreparer(ctx context.Context, resourceGroupName string, containerGroupName string, containerName string, tail *int32, timestamps *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerGroupName": autorest.Encode("path", containerGroupName),
		"containerName":      autorest.Encode("path", containerName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if tail != nil {
		queryParameters["tail"] = autorest.Encode("query", *tail)
	}
	if timestamps != nil {
		queryParameters["timestamps"] = autorest.Encode("query", *timestamps)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}/containers/{containerName}/logs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListLogsSender sends the ListLogs request. The method will close the
// http.Response Body if it receives an error.
func (client ContainersClient) ListLogsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListLogsResponder handles the response to the ListLogs request. The method always
// closes the http.Response Body.
func (client ContainersClient) ListLogsResponder(resp *http.Response) (result Logs, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
