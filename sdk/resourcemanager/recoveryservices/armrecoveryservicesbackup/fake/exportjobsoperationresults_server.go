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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v3"
	"net/http"
	"net/url"
	"regexp"
)

// ExportJobsOperationResultsServer is a fake server for instances of the armrecoveryservicesbackup.ExportJobsOperationResultsClient type.
type ExportJobsOperationResultsServer struct {
	// Get is the fake for method ExportJobsOperationResultsClient.Get
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	Get func(ctx context.Context, vaultName string, resourceGroupName string, operationID string, options *armrecoveryservicesbackup.ExportJobsOperationResultsClientGetOptions) (resp azfake.Responder[armrecoveryservicesbackup.ExportJobsOperationResultsClientGetResponse], errResp azfake.ErrorResponder)
}

// NewExportJobsOperationResultsServerTransport creates a new instance of ExportJobsOperationResultsServerTransport with the provided implementation.
// The returned ExportJobsOperationResultsServerTransport instance is connected to an instance of armrecoveryservicesbackup.ExportJobsOperationResultsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewExportJobsOperationResultsServerTransport(srv *ExportJobsOperationResultsServer) *ExportJobsOperationResultsServerTransport {
	return &ExportJobsOperationResultsServerTransport{srv: srv}
}

// ExportJobsOperationResultsServerTransport connects instances of armrecoveryservicesbackup.ExportJobsOperationResultsClient to instances of ExportJobsOperationResultsServer.
// Don't use this type directly, use NewExportJobsOperationResultsServerTransport instead.
type ExportJobsOperationResultsServerTransport struct {
	srv *ExportJobsOperationResultsServer
}

// Do implements the policy.Transporter interface for ExportJobsOperationResultsServerTransport.
func (e *ExportJobsOperationResultsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ExportJobsOperationResultsClient.Get":
		resp, err = e.dispatchGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *ExportJobsOperationResultsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if e.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupJobs/operationResults/(?P<operationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	operationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("operationId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Get(req.Context(), vaultNameParam, resourceGroupNameParam, operationIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusAccepted}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).OperationResultInfoBaseResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
