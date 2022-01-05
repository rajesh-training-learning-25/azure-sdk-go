//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdeviceprovisioningservices

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// DpsCertificateClient contains the methods for the DpsCertificate group.
// Don't use this type directly, use NewDpsCertificateClient() instead.
type DpsCertificateClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewDpsCertificateClient creates a new instance of DpsCertificateClient with the specified values.
func NewDpsCertificateClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *DpsCertificateClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &DpsCertificateClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CreateOrUpdate - Add new certificate or update an existing certificate.
// If the operation fails it returns the *ErrorDetails error type.
func (client *DpsCertificateClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, provisioningServiceName string, certificateName string, certificateDescription CertificateBodyDescription, options *DpsCertificateCreateOrUpdateOptions) (DpsCertificateCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, provisioningServiceName, certificateName, certificateDescription, options)
	if err != nil {
		return DpsCertificateCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DpsCertificateCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DpsCertificateCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DpsCertificateClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, provisioningServiceName string, certificateName string, certificateDescription CertificateBodyDescription, options *DpsCertificateCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}/certificates/{certificateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if provisioningServiceName == "" {
		return nil, errors.New("parameter provisioningServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{provisioningServiceName}", url.PathEscape(provisioningServiceName))
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, certificateDescription)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DpsCertificateClient) createOrUpdateHandleResponse(resp *http.Response) (DpsCertificateCreateOrUpdateResponse, error) {
	result := DpsCertificateCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CertificateResponse); err != nil {
		return DpsCertificateCreateOrUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *DpsCertificateClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorDetails{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Deletes the specified certificate associated with the Provisioning Service
// If the operation fails it returns the *ErrorDetails error type.
func (client *DpsCertificateClient) Delete(ctx context.Context, resourceGroupName string, ifMatch string, provisioningServiceName string, certificateName string, options *DpsCertificateDeleteOptions) (DpsCertificateDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, ifMatch, provisioningServiceName, certificateName, options)
	if err != nil {
		return DpsCertificateDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DpsCertificateDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return DpsCertificateDeleteResponse{}, client.deleteHandleError(resp)
	}
	return DpsCertificateDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DpsCertificateClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, ifMatch string, provisioningServiceName string, certificateName string, options *DpsCertificateDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}/certificates/{certificateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if provisioningServiceName == "" {
		return nil, errors.New("parameter provisioningServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{provisioningServiceName}", url.PathEscape(provisioningServiceName))
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.CertificateName1 != nil {
		reqQP.Set("certificate.name", *options.CertificateName1)
	}
	if options != nil && options.CertificateRawBytes != nil {
		reqQP.Set("certificate.rawBytes", base64.StdEncoding.EncodeToString(options.CertificateRawBytes))
	}
	if options != nil && options.CertificateIsVerified != nil {
		reqQP.Set("certificate.isVerified", strconv.FormatBool(*options.CertificateIsVerified))
	}
	if options != nil && options.CertificatePurpose != nil {
		reqQP.Set("certificate.purpose", string(*options.CertificatePurpose))
	}
	if options != nil && options.CertificateCreated != nil {
		reqQP.Set("certificate.created", options.CertificateCreated.Format(time.RFC3339Nano))
	}
	if options != nil && options.CertificateLastUpdated != nil {
		reqQP.Set("certificate.lastUpdated", options.CertificateLastUpdated.Format(time.RFC3339Nano))
	}
	if options != nil && options.CertificateHasPrivateKey != nil {
		reqQP.Set("certificate.hasPrivateKey", strconv.FormatBool(*options.CertificateHasPrivateKey))
	}
	if options != nil && options.CertificateNonce != nil {
		reqQP.Set("certificate.nonce", *options.CertificateNonce)
	}
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("If-Match", ifMatch)
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *DpsCertificateClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorDetails{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GenerateVerificationCode - Generate verification code for Proof of Possession.
// If the operation fails it returns the *ErrorDetails error type.
func (client *DpsCertificateClient) GenerateVerificationCode(ctx context.Context, certificateName string, ifMatch string, resourceGroupName string, provisioningServiceName string, options *DpsCertificateGenerateVerificationCodeOptions) (DpsCertificateGenerateVerificationCodeResponse, error) {
	req, err := client.generateVerificationCodeCreateRequest(ctx, certificateName, ifMatch, resourceGroupName, provisioningServiceName, options)
	if err != nil {
		return DpsCertificateGenerateVerificationCodeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DpsCertificateGenerateVerificationCodeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DpsCertificateGenerateVerificationCodeResponse{}, client.generateVerificationCodeHandleError(resp)
	}
	return client.generateVerificationCodeHandleResponse(resp)
}

// generateVerificationCodeCreateRequest creates the GenerateVerificationCode request.
func (client *DpsCertificateClient) generateVerificationCodeCreateRequest(ctx context.Context, certificateName string, ifMatch string, resourceGroupName string, provisioningServiceName string, options *DpsCertificateGenerateVerificationCodeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}/certificates/{certificateName}/generateVerificationCode"
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if provisioningServiceName == "" {
		return nil, errors.New("parameter provisioningServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{provisioningServiceName}", url.PathEscape(provisioningServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.CertificateName1 != nil {
		reqQP.Set("certificate.name", *options.CertificateName1)
	}
	if options != nil && options.CertificateRawBytes != nil {
		reqQP.Set("certificate.rawBytes", base64.StdEncoding.EncodeToString(options.CertificateRawBytes))
	}
	if options != nil && options.CertificateIsVerified != nil {
		reqQP.Set("certificate.isVerified", strconv.FormatBool(*options.CertificateIsVerified))
	}
	if options != nil && options.CertificatePurpose != nil {
		reqQP.Set("certificate.purpose", string(*options.CertificatePurpose))
	}
	if options != nil && options.CertificateCreated != nil {
		reqQP.Set("certificate.created", options.CertificateCreated.Format(time.RFC3339Nano))
	}
	if options != nil && options.CertificateLastUpdated != nil {
		reqQP.Set("certificate.lastUpdated", options.CertificateLastUpdated.Format(time.RFC3339Nano))
	}
	if options != nil && options.CertificateHasPrivateKey != nil {
		reqQP.Set("certificate.hasPrivateKey", strconv.FormatBool(*options.CertificateHasPrivateKey))
	}
	if options != nil && options.CertificateNonce != nil {
		reqQP.Set("certificate.nonce", *options.CertificateNonce)
	}
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("If-Match", ifMatch)
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// generateVerificationCodeHandleResponse handles the GenerateVerificationCode response.
func (client *DpsCertificateClient) generateVerificationCodeHandleResponse(resp *http.Response) (DpsCertificateGenerateVerificationCodeResponse, error) {
	result := DpsCertificateGenerateVerificationCodeResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VerificationCodeResponse); err != nil {
		return DpsCertificateGenerateVerificationCodeResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// generateVerificationCodeHandleError handles the GenerateVerificationCode error response.
func (client *DpsCertificateClient) generateVerificationCodeHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorDetails{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Get the certificate from the provisioning service.
// If the operation fails it returns the *ErrorDetails error type.
func (client *DpsCertificateClient) Get(ctx context.Context, certificateName string, resourceGroupName string, provisioningServiceName string, options *DpsCertificateGetOptions) (DpsCertificateGetResponse, error) {
	req, err := client.getCreateRequest(ctx, certificateName, resourceGroupName, provisioningServiceName, options)
	if err != nil {
		return DpsCertificateGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DpsCertificateGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DpsCertificateGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DpsCertificateClient) getCreateRequest(ctx context.Context, certificateName string, resourceGroupName string, provisioningServiceName string, options *DpsCertificateGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}/certificates/{certificateName}"
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if provisioningServiceName == "" {
		return nil, errors.New("parameter provisioningServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{provisioningServiceName}", url.PathEscape(provisioningServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DpsCertificateClient) getHandleResponse(resp *http.Response) (DpsCertificateGetResponse, error) {
	result := DpsCertificateGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CertificateResponse); err != nil {
		return DpsCertificateGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *DpsCertificateClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorDetails{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Get all the certificates tied to the provisioning service.
// If the operation fails it returns the *ErrorDetails error type.
func (client *DpsCertificateClient) List(ctx context.Context, resourceGroupName string, provisioningServiceName string, options *DpsCertificateListOptions) (DpsCertificateListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, provisioningServiceName, options)
	if err != nil {
		return DpsCertificateListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DpsCertificateListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DpsCertificateListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *DpsCertificateClient) listCreateRequest(ctx context.Context, resourceGroupName string, provisioningServiceName string, options *DpsCertificateListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}/certificates"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if provisioningServiceName == "" {
		return nil, errors.New("parameter provisioningServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{provisioningServiceName}", url.PathEscape(provisioningServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DpsCertificateClient) listHandleResponse(resp *http.Response) (DpsCertificateListResponse, error) {
	result := DpsCertificateListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CertificateListDescription); err != nil {
		return DpsCertificateListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *DpsCertificateClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorDetails{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// VerifyCertificate - Verifies the certificate's private key possession by providing the leaf cert issued by the verifying pre uploaded certificate.
// If the operation fails it returns the *ErrorDetails error type.
func (client *DpsCertificateClient) VerifyCertificate(ctx context.Context, certificateName string, ifMatch string, resourceGroupName string, provisioningServiceName string, request VerificationCodeRequest, options *DpsCertificateVerifyCertificateOptions) (DpsCertificateVerifyCertificateResponse, error) {
	req, err := client.verifyCertificateCreateRequest(ctx, certificateName, ifMatch, resourceGroupName, provisioningServiceName, request, options)
	if err != nil {
		return DpsCertificateVerifyCertificateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DpsCertificateVerifyCertificateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DpsCertificateVerifyCertificateResponse{}, client.verifyCertificateHandleError(resp)
	}
	return client.verifyCertificateHandleResponse(resp)
}

// verifyCertificateCreateRequest creates the VerifyCertificate request.
func (client *DpsCertificateClient) verifyCertificateCreateRequest(ctx context.Context, certificateName string, ifMatch string, resourceGroupName string, provisioningServiceName string, request VerificationCodeRequest, options *DpsCertificateVerifyCertificateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}/certificates/{certificateName}/verify"
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if provisioningServiceName == "" {
		return nil, errors.New("parameter provisioningServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{provisioningServiceName}", url.PathEscape(provisioningServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.CertificateName1 != nil {
		reqQP.Set("certificate.name", *options.CertificateName1)
	}
	if options != nil && options.CertificateRawBytes != nil {
		reqQP.Set("certificate.rawBytes", base64.StdEncoding.EncodeToString(options.CertificateRawBytes))
	}
	if options != nil && options.CertificateIsVerified != nil {
		reqQP.Set("certificate.isVerified", strconv.FormatBool(*options.CertificateIsVerified))
	}
	if options != nil && options.CertificatePurpose != nil {
		reqQP.Set("certificate.purpose", string(*options.CertificatePurpose))
	}
	if options != nil && options.CertificateCreated != nil {
		reqQP.Set("certificate.created", options.CertificateCreated.Format(time.RFC3339Nano))
	}
	if options != nil && options.CertificateLastUpdated != nil {
		reqQP.Set("certificate.lastUpdated", options.CertificateLastUpdated.Format(time.RFC3339Nano))
	}
	if options != nil && options.CertificateHasPrivateKey != nil {
		reqQP.Set("certificate.hasPrivateKey", strconv.FormatBool(*options.CertificateHasPrivateKey))
	}
	if options != nil && options.CertificateNonce != nil {
		reqQP.Set("certificate.nonce", *options.CertificateNonce)
	}
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("If-Match", ifMatch)
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, request)
}

// verifyCertificateHandleResponse handles the VerifyCertificate response.
func (client *DpsCertificateClient) verifyCertificateHandleResponse(resp *http.Response) (DpsCertificateVerifyCertificateResponse, error) {
	result := DpsCertificateVerifyCertificateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CertificateResponse); err != nil {
		return DpsCertificateVerifyCertificateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// verifyCertificateHandleError handles the VerifyCertificate error response.
func (client *DpsCertificateClient) verifyCertificateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorDetails{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
