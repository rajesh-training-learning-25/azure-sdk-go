//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
	"net/http"
	"regexp"
)

// TaskRunsServer is a fake server for instances of the armcontainerregistry.TaskRunsClient type.
type TaskRunsServer struct {
	// BeginCreate is the fake for method TaskRunsClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, registryName string, taskRunName string, taskRun armcontainerregistry.TaskRun, options *armcontainerregistry.TaskRunsClientBeginCreateOptions) (resp azfake.PollerResponder[armcontainerregistry.TaskRunsClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method TaskRunsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, registryName string, taskRunName string, options *armcontainerregistry.TaskRunsClientBeginDeleteOptions) (resp azfake.PollerResponder[armcontainerregistry.TaskRunsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method TaskRunsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, registryName string, taskRunName string, options *armcontainerregistry.TaskRunsClientGetOptions) (resp azfake.Responder[armcontainerregistry.TaskRunsClientGetResponse], errResp azfake.ErrorResponder)

	// GetDetails is the fake for method TaskRunsClient.GetDetails
	// HTTP status codes to indicate success: http.StatusOK
	GetDetails func(ctx context.Context, resourceGroupName string, registryName string, taskRunName string, options *armcontainerregistry.TaskRunsClientGetDetailsOptions) (resp azfake.Responder[armcontainerregistry.TaskRunsClientGetDetailsResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method TaskRunsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, registryName string, options *armcontainerregistry.TaskRunsClientListOptions) (resp azfake.PagerResponder[armcontainerregistry.TaskRunsClientListResponse])

	// BeginUpdate is the fake for method TaskRunsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginUpdate func(ctx context.Context, resourceGroupName string, registryName string, taskRunName string, updateParameters armcontainerregistry.TaskRunUpdateParameters, options *armcontainerregistry.TaskRunsClientBeginUpdateOptions) (resp azfake.PollerResponder[armcontainerregistry.TaskRunsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewTaskRunsServerTransport creates a new instance of TaskRunsServerTransport with the provided implementation.
// The returned TaskRunsServerTransport instance is connected to an instance of armcontainerregistry.TaskRunsClient by way of the
// undefined.Transporter field.
func NewTaskRunsServerTransport(srv *TaskRunsServer) *TaskRunsServerTransport {
	return &TaskRunsServerTransport{srv: srv}
}

// TaskRunsServerTransport connects instances of armcontainerregistry.TaskRunsClient to instances of TaskRunsServer.
// Don't use this type directly, use NewTaskRunsServerTransport instead.
type TaskRunsServerTransport struct {
	srv          *TaskRunsServer
	beginCreate  *azfake.PollerResponder[armcontainerregistry.TaskRunsClientCreateResponse]
	beginDelete  *azfake.PollerResponder[armcontainerregistry.TaskRunsClientDeleteResponse]
	newListPager *azfake.PagerResponder[armcontainerregistry.TaskRunsClientListResponse]
	beginUpdate  *azfake.PollerResponder[armcontainerregistry.TaskRunsClientUpdateResponse]
}

// Do implements the policy.Transporter interface for TaskRunsServerTransport.
func (t *TaskRunsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "TaskRunsClient.BeginCreate":
		resp, err = t.dispatchBeginCreate(req)
	case "TaskRunsClient.BeginDelete":
		resp, err = t.dispatchBeginDelete(req)
	case "TaskRunsClient.Get":
		resp, err = t.dispatchGet(req)
	case "TaskRunsClient.GetDetails":
		resp, err = t.dispatchGetDetails(req)
	case "TaskRunsClient.NewListPager":
		resp, err = t.dispatchNewListPager(req)
	case "TaskRunsClient.BeginUpdate":
		resp, err = t.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (t *TaskRunsServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if t.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("method BeginCreate not implemented")}
	}
	if t.beginCreate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.ContainerRegistry/registries/(?P<registryName>[a-zA-Z0-9-_]+)/taskRuns/(?P<taskRunName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcontainerregistry.TaskRun](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := t.srv.BeginCreate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("registryName")], matches[regex.SubexpIndex("taskRunName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		t.beginCreate = &respr
	}

	resp, err := server.PollerResponderNext(t.beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(t.beginCreate) {
		t.beginCreate = nil
	}

	return resp, nil
}

func (t *TaskRunsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if t.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("method BeginDelete not implemented")}
	}
	if t.beginDelete == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.ContainerRegistry/registries/(?P<registryName>[a-zA-Z0-9-_]+)/taskRuns/(?P<taskRunName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := t.srv.BeginDelete(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("registryName")], matches[regex.SubexpIndex("taskRunName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		t.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(t.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(t.beginDelete) {
		t.beginDelete = nil
	}

	return resp, nil
}

func (t *TaskRunsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if t.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.ContainerRegistry/registries/(?P<registryName>[a-zA-Z0-9-_]+)/taskRuns/(?P<taskRunName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := t.srv.Get(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("registryName")], matches[regex.SubexpIndex("taskRunName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).TaskRun, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TaskRunsServerTransport) dispatchGetDetails(req *http.Request) (*http.Response, error) {
	if t.srv.GetDetails == nil {
		return nil, &nonRetriableError{errors.New("method GetDetails not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.ContainerRegistry/registries/(?P<registryName>[a-zA-Z0-9-_]+)/taskRuns/(?P<taskRunName>[a-zA-Z0-9-_]+)/listDetails"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := t.srv.GetDetails(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("registryName")], matches[regex.SubexpIndex("taskRunName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).TaskRun, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TaskRunsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if t.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListPager not implemented")}
	}
	if t.newListPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.ContainerRegistry/registries/(?P<registryName>[a-zA-Z0-9-_]+)/taskRuns"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := t.srv.NewListPager(matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("registryName")], nil)
		t.newListPager = &resp
		server.PagerResponderInjectNextLinks(t.newListPager, req, func(page *armcontainerregistry.TaskRunsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(t.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(t.newListPager) {
		t.newListPager = nil
	}
	return resp, nil
}

func (t *TaskRunsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if t.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("method BeginUpdate not implemented")}
	}
	if t.beginUpdate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.ContainerRegistry/registries/(?P<registryName>[a-zA-Z0-9-_]+)/taskRuns/(?P<taskRunName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcontainerregistry.TaskRunUpdateParameters](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := t.srv.BeginUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("registryName")], matches[regex.SubexpIndex("taskRunName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		t.beginUpdate = &respr
	}

	resp, err := server.PollerResponderNext(t.beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(t.beginUpdate) {
		t.beginUpdate = nil
	}

	return resp, nil
}
