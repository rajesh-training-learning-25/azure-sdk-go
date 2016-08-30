// Package osimage provides a client for Operating System Images.
package osimage

import (
	"encoding/xml"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/management"
)

const (
	azureImageListURL    = "services/images"
	azureImageShareURL   = "services/images/%s/share?permission=%s"
	errInvalidImage      = "Can not find image %s in specified subscription, please specify another image name."
	errParamNotSpecified = "Parameter %s is not specified."
)

// NewClient is used to instantiate a new OSImageClient from an Azure client.
func NewClient(client management.Client) OSImageClient {
	return OSImageClient{client: client}
}

func (c OSImageClient) ListOSImages() (ListOSImagesResponse, error) {
	var l ListOSImagesResponse

	response, err := c.client.SendAzureGetRequest(azureImageListURL)
	if err != nil {
		return l, err
	}

	err = xml.Unmarshal(response, &l)
	return l, err
}

func (c OSImageClient) ShareImage(image string, permission ImagePermission) (management.OperationID, error) {
	url := fmt.Sprintf(azureImageShareURL, image, permission)
	return c.client.SendAzurePutRequest(url, "", []byte{})
}
