package storage

import (
	"fmt"
	"net/http"
	"net/url"
)

// BlobServiceClient contains operations for Microsoft Azure Blob Storage
// Service.
type BlobServiceClient struct {
	client Client
	auth   authentication
}

// ListContainersParameters defines the set of customizable parameters to make a
// List Containers call.
//
// See https://msdn.microsoft.com/en-us/library/azure/dd179352.aspx
type ListContainersParameters struct {
	Prefix     string
	Marker     string
	Include    string
	MaxResults uint
	Timeout    uint
}

// GetContainerReference returns a Container object for the specified container name.
func (b BlobServiceClient) GetContainerReference(name string) Container {
	return Container{
		bsc:  &b,
		Name: name,
	}
}

// ListContainers returns the list of containers in a storage account along with
// pagination token and other response details.
//
// See https://msdn.microsoft.com/en-us/library/azure/dd179352.aspx
func (b BlobServiceClient) ListContainers(params ListContainersParameters) (*ContainerListResponse, error) {
	q := mergeParams(params.getParameters(), url.Values{"comp": {"list"}})
	uri := b.client.getEndpoint(blobServiceName, "", q)
	headers := b.client.getStandardHeaders()

	var out ContainerListResponse
	resp, err := b.client.exec(http.MethodGet, uri, headers, nil, b.auth)
	if err != nil {
		return nil, err
	}
	defer resp.body.Close()
	err = xmlUnmarshal(resp.body, &out)

	// assign our client to the newly created Container objects
	for i := range out.Containers {
		out.Containers[i].bsc = &b
	}
	return &out, err
}

func (p ListContainersParameters) getParameters() url.Values {
	out := url.Values{}

	if p.Prefix != "" {
		out.Set("prefix", p.Prefix)
	}
	if p.Marker != "" {
		out.Set("marker", p.Marker)
	}
	if p.Include != "" {
		out.Set("include", p.Include)
	}
	if p.MaxResults != 0 {
		out.Set("maxresults", fmt.Sprintf("%v", p.MaxResults))
	}
	if p.Timeout != 0 {
		out.Set("timeout", fmt.Sprintf("%v", p.Timeout))
	}

	return out
}
