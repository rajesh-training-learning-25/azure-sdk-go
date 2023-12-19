//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azcontainerregistry

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Client contains the methods for the ContainerRegistry group.
// Don't use this type directly, use a constructor function instead.
type Client struct {
	internal *azcore.Client
	endpoint string
}

// DeleteManifest - Delete the manifest identified by name and reference. Note that a manifest can only be deleted by digest.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
//   - name - Name of the image (including the namespace)
//   - digest - Digest of a BLOB
//   - options - ClientDeleteManifestOptions contains the optional parameters for the Client.DeleteManifest method.
func (client *Client) DeleteManifest(ctx context.Context, name string, digest string, options *ClientDeleteManifestOptions) (ClientDeleteManifestResponse, error) {
	var err error
	req, err := client.deleteManifestCreateRequest(ctx, name, digest, options)
	if err != nil {
		return ClientDeleteManifestResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientDeleteManifestResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNotFound) {
		err = runtime.NewResponseError(httpResp)
		return ClientDeleteManifestResponse{}, err
	}
	return ClientDeleteManifestResponse{}, nil
}

// deleteManifestCreateRequest creates the DeleteManifest request.
func (client *Client) deleteManifestCreateRequest(ctx context.Context, name string, digest string, options *ClientDeleteManifestOptions) (*policy.Request, error) {
	urlPath := "/v2/{name}/manifests/{reference}"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if digest == "" {
		return nil, errors.New("parameter digest cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reference}", url.PathEscape(digest))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// DeleteRepository - Delete the repository identified by name
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
//   - name - Name of the image (including the namespace)
//   - options - ClientDeleteRepositoryOptions contains the optional parameters for the Client.DeleteRepository method.
func (client *Client) DeleteRepository(ctx context.Context, name string, options *ClientDeleteRepositoryOptions) (ClientDeleteRepositoryResponse, error) {
	var err error
	req, err := client.deleteRepositoryCreateRequest(ctx, name, options)
	if err != nil {
		return ClientDeleteRepositoryResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientDeleteRepositoryResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNotFound) {
		err = runtime.NewResponseError(httpResp)
		return ClientDeleteRepositoryResponse{}, err
	}
	return ClientDeleteRepositoryResponse{}, nil
}

// deleteRepositoryCreateRequest creates the DeleteRepository request.
func (client *Client) deleteRepositoryCreateRequest(ctx context.Context, name string, options *ClientDeleteRepositoryOptions) (*policy.Request, error) {
	urlPath := "/acr/v1/{name}"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// DeleteTag - Delete tag
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
//   - name - Name of the image (including the namespace)
//   - tag - Tag name
//   - options - ClientDeleteTagOptions contains the optional parameters for the Client.DeleteTag method.
func (client *Client) DeleteTag(ctx context.Context, name string, tag string, options *ClientDeleteTagOptions) (ClientDeleteTagResponse, error) {
	var err error
	req, err := client.deleteTagCreateRequest(ctx, name, tag, options)
	if err != nil {
		return ClientDeleteTagResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientDeleteTagResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNotFound) {
		err = runtime.NewResponseError(httpResp)
		return ClientDeleteTagResponse{}, err
	}
	return ClientDeleteTagResponse{}, nil
}

// deleteTagCreateRequest creates the DeleteTag request.
func (client *Client) deleteTagCreateRequest(ctx context.Context, name string, tag string, options *ClientDeleteTagOptions) (*policy.Request, error) {
	urlPath := "/acr/v1/{name}/_tags/{reference}"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if tag == "" {
		return nil, errors.New("parameter tag cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reference}", url.PathEscape(tag))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetManifest - Get the manifest identified by name and reference where reference can be a tag or digest.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
//   - name - Name of the image (including the namespace)
//   - reference - A tag or a digest, pointing to a specific image
//   - options - ClientGetManifestOptions contains the optional parameters for the Client.GetManifest method.
func (client *Client) GetManifest(ctx context.Context, name string, reference string, options *ClientGetManifestOptions) (ClientGetManifestResponse, error) {
	var err error
	req, err := client.getManifestCreateRequest(ctx, name, reference, options)
	if err != nil {
		return ClientGetManifestResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientGetManifestResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ClientGetManifestResponse{}, err
	}
	resp, err := client.getManifestHandleResponse(httpResp)
	return resp, err
}

// getManifestCreateRequest creates the GetManifest request.
func (client *Client) getManifestCreateRequest(ctx context.Context, name string, reference string, options *ClientGetManifestOptions) (*policy.Request, error) {
	urlPath := "/v2/{name}/manifests/{reference}"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if reference == "" {
		return nil, errors.New("parameter reference cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reference}", url.PathEscape(reference))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	runtime.SkipBodyDownload(req)
	if options != nil && options.Accept != nil {
		req.Raw().Header["accept"] = []string{*options.Accept}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getManifestHandleResponse handles the GetManifest response.
func (client *Client) getManifestHandleResponse(resp *http.Response) (ClientGetManifestResponse, error) {
	result := ClientGetManifestResponse{ManifestData: resp.Body}
	if val := resp.Header.Get("Docker-Content-Digest"); val != "" {
		result.DockerContentDigest = &val
	}
	return result, nil
}

// GetManifestProperties - Get manifest attributes
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
//   - name - Name of the image (including the namespace)
//   - digest - Digest of a BLOB
//   - options - ClientGetManifestPropertiesOptions contains the optional parameters for the Client.GetManifestProperties method.
func (client *Client) GetManifestProperties(ctx context.Context, name string, digest string, options *ClientGetManifestPropertiesOptions) (ClientGetManifestPropertiesResponse, error) {
	var err error
	req, err := client.getManifestPropertiesCreateRequest(ctx, name, digest, options)
	if err != nil {
		return ClientGetManifestPropertiesResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientGetManifestPropertiesResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ClientGetManifestPropertiesResponse{}, err
	}
	resp, err := client.getManifestPropertiesHandleResponse(httpResp)
	return resp, err
}

// getManifestPropertiesCreateRequest creates the GetManifestProperties request.
func (client *Client) getManifestPropertiesCreateRequest(ctx context.Context, name string, digest string, options *ClientGetManifestPropertiesOptions) (*policy.Request, error) {
	urlPath := "/acr/v1/{name}/_manifests/{digest}"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if digest == "" {
		return nil, errors.New("parameter digest cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{digest}", url.PathEscape(digest))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getManifestPropertiesHandleResponse handles the GetManifestProperties response.
func (client *Client) getManifestPropertiesHandleResponse(resp *http.Response) (ClientGetManifestPropertiesResponse, error) {
	result := ClientGetManifestPropertiesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArtifactManifestProperties); err != nil {
		return ClientGetManifestPropertiesResponse{}, err
	}
	return result, nil
}

// GetRepositoryProperties - Get repository attributes
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
//   - name - Name of the image (including the namespace)
//   - options - ClientGetRepositoryPropertiesOptions contains the optional parameters for the Client.GetRepositoryProperties
//     method.
func (client *Client) GetRepositoryProperties(ctx context.Context, name string, options *ClientGetRepositoryPropertiesOptions) (ClientGetRepositoryPropertiesResponse, error) {
	var err error
	req, err := client.getRepositoryPropertiesCreateRequest(ctx, name, options)
	if err != nil {
		return ClientGetRepositoryPropertiesResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientGetRepositoryPropertiesResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ClientGetRepositoryPropertiesResponse{}, err
	}
	resp, err := client.getRepositoryPropertiesHandleResponse(httpResp)
	return resp, err
}

// getRepositoryPropertiesCreateRequest creates the GetRepositoryProperties request.
func (client *Client) getRepositoryPropertiesCreateRequest(ctx context.Context, name string, options *ClientGetRepositoryPropertiesOptions) (*policy.Request, error) {
	urlPath := "/acr/v1/{name}"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getRepositoryPropertiesHandleResponse handles the GetRepositoryProperties response.
func (client *Client) getRepositoryPropertiesHandleResponse(resp *http.Response) (ClientGetRepositoryPropertiesResponse, error) {
	result := ClientGetRepositoryPropertiesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContainerRepositoryProperties); err != nil {
		return ClientGetRepositoryPropertiesResponse{}, err
	}
	return result, nil
}

// GetTagProperties - Get tag attributes by tag
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
//   - name - Name of the image (including the namespace)
//   - tag - Tag name
//   - options - ClientGetTagPropertiesOptions contains the optional parameters for the Client.GetTagProperties method.
func (client *Client) GetTagProperties(ctx context.Context, name string, tag string, options *ClientGetTagPropertiesOptions) (ClientGetTagPropertiesResponse, error) {
	var err error
	req, err := client.getTagPropertiesCreateRequest(ctx, name, tag, options)
	if err != nil {
		return ClientGetTagPropertiesResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientGetTagPropertiesResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ClientGetTagPropertiesResponse{}, err
	}
	resp, err := client.getTagPropertiesHandleResponse(httpResp)
	return resp, err
}

// getTagPropertiesCreateRequest creates the GetTagProperties request.
func (client *Client) getTagPropertiesCreateRequest(ctx context.Context, name string, tag string, options *ClientGetTagPropertiesOptions) (*policy.Request, error) {
	urlPath := "/acr/v1/{name}/_tags/{reference}"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if tag == "" {
		return nil, errors.New("parameter tag cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reference}", url.PathEscape(tag))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getTagPropertiesHandleResponse handles the GetTagProperties response.
func (client *Client) getTagPropertiesHandleResponse(resp *http.Response) (ClientGetTagPropertiesResponse, error) {
	result := ClientGetTagPropertiesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArtifactTagProperties); err != nil {
		return ClientGetTagPropertiesResponse{}, err
	}
	return result, nil
}

// NewListManifestsPager - List manifests of a repository
//
// Generated from API version 2021-07-01
//   - name - Name of the image (including the namespace)
//   - options - ClientListManifestsOptions contains the optional parameters for the Client.NewListManifestsPager method.
func (client *Client) NewListManifestsPager(name string, options *ClientListManifestsOptions) *runtime.Pager[ClientListManifestsResponse] {
	return runtime.NewPager(runtime.PagingHandler[ClientListManifestsResponse]{
		More: func(page ClientListManifestsResponse) bool {
			return page.Link != nil && len(*page.Link) > 0
		},
		Fetcher: func(ctx context.Context, page *ClientListManifestsResponse) (ClientListManifestsResponse, error) {
			nextLink := ""
			if page != nil {
				nextLink = *page.Link
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listManifestsCreateRequest(ctx, name, options)
			}, nil)
			if err != nil {
				return ClientListManifestsResponse{}, err
			}
			return client.listManifestsHandleResponse(resp)
		},
	})
}

// listManifestsCreateRequest creates the ListManifests request.
func (client *Client) listManifestsCreateRequest(ctx context.Context, name string, options *ClientListManifestsOptions) (*policy.Request, error) {
	urlPath := "/acr/v1/{name}/_manifests"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Last != nil {
		reqQP.Set("last", *options.Last)
	}
	if options != nil && options.MaxNum != nil {
		reqQP.Set("n", strconv.FormatInt(int64(*options.MaxNum), 10))
	}
	reqQP.Set("api-version", "2021-07-01")
	if options != nil && options.OrderBy != nil {
		reqQP.Set("orderby", string(*options.OrderBy))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listManifestsHandleResponse handles the ListManifests response.
func (client *Client) listManifestsHandleResponse(resp *http.Response) (ClientListManifestsResponse, error) {
	result := ClientListManifestsResponse{}
	if val := resp.Header.Get("Link"); val != "" {
		val = runtime.JoinPaths(client.endpoint, extractNextLink(val))
		result.Link = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Manifests); err != nil {
		return ClientListManifestsResponse{}, err
	}
	return result, nil
}

// NewListRepositoriesPager - List repositories
//
// Generated from API version 2021-07-01
//   - options - ClientListRepositoriesOptions contains the optional parameters for the Client.NewListRepositoriesPager method.
func (client *Client) NewListRepositoriesPager(options *ClientListRepositoriesOptions) *runtime.Pager[ClientListRepositoriesResponse] {
	return runtime.NewPager(runtime.PagingHandler[ClientListRepositoriesResponse]{
		More: func(page ClientListRepositoriesResponse) bool {
			return page.Link != nil && len(*page.Link) > 0
		},
		Fetcher: func(ctx context.Context, page *ClientListRepositoriesResponse) (ClientListRepositoriesResponse, error) {
			nextLink := ""
			if page != nil {
				nextLink = *page.Link
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listRepositoriesCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return ClientListRepositoriesResponse{}, err
			}
			return client.listRepositoriesHandleResponse(resp)
		},
	})
}

// listRepositoriesCreateRequest creates the ListRepositories request.
func (client *Client) listRepositoriesCreateRequest(ctx context.Context, options *ClientListRepositoriesOptions) (*policy.Request, error) {
	urlPath := "/acr/v1/_catalog"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Last != nil {
		reqQP.Set("last", *options.Last)
	}
	if options != nil && options.MaxNum != nil {
		reqQP.Set("n", strconv.FormatInt(int64(*options.MaxNum), 10))
	}
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listRepositoriesHandleResponse handles the ListRepositories response.
func (client *Client) listRepositoriesHandleResponse(resp *http.Response) (ClientListRepositoriesResponse, error) {
	result := ClientListRepositoriesResponse{}
	if val := resp.Header.Get("Link"); val != "" {
		val = runtime.JoinPaths(client.endpoint, extractNextLink(val))
		result.Link = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Repositories); err != nil {
		return ClientListRepositoriesResponse{}, err
	}
	return result, nil
}

// NewListTagsPager - List tags of a repository
//
// Generated from API version 2021-07-01
//   - name - Name of the image (including the namespace)
//   - options - ClientListTagsOptions contains the optional parameters for the Client.NewListTagsPager method.
func (client *Client) NewListTagsPager(name string, options *ClientListTagsOptions) *runtime.Pager[ClientListTagsResponse] {
	return runtime.NewPager(runtime.PagingHandler[ClientListTagsResponse]{
		More: func(page ClientListTagsResponse) bool {
			return page.Link != nil && len(*page.Link) > 0
		},
		Fetcher: func(ctx context.Context, page *ClientListTagsResponse) (ClientListTagsResponse, error) {
			nextLink := ""
			if page != nil {
				nextLink = *page.Link
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listTagsCreateRequest(ctx, name, options)
			}, nil)
			if err != nil {
				return ClientListTagsResponse{}, err
			}
			return client.listTagsHandleResponse(resp)
		},
	})
}

// listTagsCreateRequest creates the ListTags request.
func (client *Client) listTagsCreateRequest(ctx context.Context, name string, options *ClientListTagsOptions) (*policy.Request, error) {
	urlPath := "/acr/v1/{name}/_tags"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Last != nil {
		reqQP.Set("last", *options.Last)
	}
	if options != nil && options.MaxNum != nil {
		reqQP.Set("n", strconv.FormatInt(int64(*options.MaxNum), 10))
	}
	if options != nil && options.Digest != nil {
		reqQP.Set("digest", *options.Digest)
	}
	reqQP.Set("api-version", "2021-07-01")
	if options != nil && options.OrderBy != nil {
		reqQP.Set("orderby", string(*options.OrderBy))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listTagsHandleResponse handles the ListTags response.
func (client *Client) listTagsHandleResponse(resp *http.Response) (ClientListTagsResponse, error) {
	result := ClientListTagsResponse{}
	if val := resp.Header.Get("Link"); val != "" {
		val = runtime.JoinPaths(client.endpoint, extractNextLink(val))
		result.Link = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.TagList); err != nil {
		return ClientListTagsResponse{}, err
	}
	return result, nil
}

// UpdateManifestProperties - Update properties of a manifest
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
//   - name - Name of the image (including the namespace)
//   - digest - Digest of a BLOB
//   - options - ClientUpdateManifestPropertiesOptions contains the optional parameters for the Client.UpdateManifestProperties
//     method.
func (client *Client) UpdateManifestProperties(ctx context.Context, name string, digest string, options *ClientUpdateManifestPropertiesOptions) (ClientUpdateManifestPropertiesResponse, error) {
	var err error
	req, err := client.updateManifestPropertiesCreateRequest(ctx, name, digest, options)
	if err != nil {
		return ClientUpdateManifestPropertiesResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientUpdateManifestPropertiesResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ClientUpdateManifestPropertiesResponse{}, err
	}
	resp, err := client.updateManifestPropertiesHandleResponse(httpResp)
	return resp, err
}

// updateManifestPropertiesCreateRequest creates the UpdateManifestProperties request.
func (client *Client) updateManifestPropertiesCreateRequest(ctx context.Context, name string, digest string, options *ClientUpdateManifestPropertiesOptions) (*policy.Request, error) {
	urlPath := "/acr/v1/{name}/_manifests/{digest}"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if digest == "" {
		return nil, errors.New("parameter digest cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{digest}", url.PathEscape(digest))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Value != nil {
		if err := runtime.MarshalAsJSON(req, *options.Value); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// updateManifestPropertiesHandleResponse handles the UpdateManifestProperties response.
func (client *Client) updateManifestPropertiesHandleResponse(resp *http.Response) (ClientUpdateManifestPropertiesResponse, error) {
	result := ClientUpdateManifestPropertiesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArtifactManifestProperties); err != nil {
		return ClientUpdateManifestPropertiesResponse{}, err
	}
	return result, nil
}

// UpdateRepositoryProperties - Update the attribute identified by name where reference is the name of the repository.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
//   - name - Name of the image (including the namespace)
//   - options - ClientUpdateRepositoryPropertiesOptions contains the optional parameters for the Client.UpdateRepositoryProperties
//     method.
func (client *Client) UpdateRepositoryProperties(ctx context.Context, name string, options *ClientUpdateRepositoryPropertiesOptions) (ClientUpdateRepositoryPropertiesResponse, error) {
	var err error
	req, err := client.updateRepositoryPropertiesCreateRequest(ctx, name, options)
	if err != nil {
		return ClientUpdateRepositoryPropertiesResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientUpdateRepositoryPropertiesResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ClientUpdateRepositoryPropertiesResponse{}, err
	}
	resp, err := client.updateRepositoryPropertiesHandleResponse(httpResp)
	return resp, err
}

// updateRepositoryPropertiesCreateRequest creates the UpdateRepositoryProperties request.
func (client *Client) updateRepositoryPropertiesCreateRequest(ctx context.Context, name string, options *ClientUpdateRepositoryPropertiesOptions) (*policy.Request, error) {
	urlPath := "/acr/v1/{name}"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Value != nil {
		if err := runtime.MarshalAsJSON(req, *options.Value); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// updateRepositoryPropertiesHandleResponse handles the UpdateRepositoryProperties response.
func (client *Client) updateRepositoryPropertiesHandleResponse(resp *http.Response) (ClientUpdateRepositoryPropertiesResponse, error) {
	result := ClientUpdateRepositoryPropertiesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContainerRepositoryProperties); err != nil {
		return ClientUpdateRepositoryPropertiesResponse{}, err
	}
	return result, nil
}

// UpdateTagProperties - Update tag attributes
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
//   - name - Name of the image (including the namespace)
//   - tag - Tag name
//   - options - ClientUpdateTagPropertiesOptions contains the optional parameters for the Client.UpdateTagProperties method.
func (client *Client) UpdateTagProperties(ctx context.Context, name string, tag string, options *ClientUpdateTagPropertiesOptions) (ClientUpdateTagPropertiesResponse, error) {
	var err error
	req, err := client.updateTagPropertiesCreateRequest(ctx, name, tag, options)
	if err != nil {
		return ClientUpdateTagPropertiesResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientUpdateTagPropertiesResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ClientUpdateTagPropertiesResponse{}, err
	}
	resp, err := client.updateTagPropertiesHandleResponse(httpResp)
	return resp, err
}

// updateTagPropertiesCreateRequest creates the UpdateTagProperties request.
func (client *Client) updateTagPropertiesCreateRequest(ctx context.Context, name string, tag string, options *ClientUpdateTagPropertiesOptions) (*policy.Request, error) {
	urlPath := "/acr/v1/{name}/_tags/{reference}"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if tag == "" {
		return nil, errors.New("parameter tag cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reference}", url.PathEscape(tag))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Value != nil {
		if err := runtime.MarshalAsJSON(req, *options.Value); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// updateTagPropertiesHandleResponse handles the UpdateTagProperties response.
func (client *Client) updateTagPropertiesHandleResponse(resp *http.Response) (ClientUpdateTagPropertiesResponse, error) {
	result := ClientUpdateTagPropertiesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArtifactTagProperties); err != nil {
		return ClientUpdateTagPropertiesResponse{}, err
	}
	return result, nil
}

// UploadManifest - Put the manifest identified by name and reference where reference can be a tag or digest.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
//   - name - Name of the image (including the namespace)
//   - reference - A tag or a digest, pointing to a specific image
//   - contentType - Upload file type
//   - manifestData - Manifest body, can take v1 or v2 values depending on accept header
//   - options - ClientUploadManifestOptions contains the optional parameters for the Client.UploadManifest method.
func (client *Client) UploadManifest(ctx context.Context, name string, reference string, contentType ContentType, manifestData io.ReadSeekCloser, options *ClientUploadManifestOptions) (ClientUploadManifestResponse, error) {
	var err error
	req, err := client.uploadManifestCreateRequest(ctx, name, reference, contentType, manifestData, options)
	if err != nil {
		return ClientUploadManifestResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientUploadManifestResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ClientUploadManifestResponse{}, err
	}
	resp, err := client.uploadManifestHandleResponse(httpResp)
	return resp, err
}

// uploadManifestCreateRequest creates the UploadManifest request.
func (client *Client) uploadManifestCreateRequest(ctx context.Context, name string, reference string, contentType ContentType, manifestData io.ReadSeekCloser, options *ClientUploadManifestOptions) (*policy.Request, error) {
	urlPath := "/v2/{name}/manifests/{reference}"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if reference == "" {
		return nil, errors.New("parameter reference cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reference}", url.PathEscape(reference))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{string(contentType)}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := req.SetBody(manifestData, string(contentType)); err != nil {
		return nil, err
	}
	return req, nil
}

// uploadManifestHandleResponse handles the UploadManifest response.
func (client *Client) uploadManifestHandleResponse(resp *http.Response) (ClientUploadManifestResponse, error) {
	result := ClientUploadManifestResponse{}
	if val := resp.Header.Get("Content-Length"); val != "" {
		contentLength, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return ClientUploadManifestResponse{}, err
		}
		result.ContentLength = &contentLength
	}
	if val := resp.Header.Get("Docker-Content-Digest"); val != "" {
		result.DockerContentDigest = &val
	}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	return result, nil
}
