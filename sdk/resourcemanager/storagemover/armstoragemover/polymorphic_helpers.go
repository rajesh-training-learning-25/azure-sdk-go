//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragemover

import "encoding/json"

func unmarshalEndpointBasePropertiesClassification(rawMsg json.RawMessage) (EndpointBasePropertiesClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b EndpointBasePropertiesClassification
	switch m["endpointType"] {
	case string(EndpointTypeAzureStorageBlobContainer):
		b = &AzureStorageBlobContainerEndpointProperties{}
	case string(EndpointTypeAzureStorageSmbFileShare):
		b = &AzureStorageSmbFileShareEndpointProperties{}
	case string(EndpointTypeNfsMount):
		b = &NfsMountEndpointProperties{}
	case string(EndpointTypeSmbMount):
		b = &SmbMountEndpointProperties{}
	default:
		b = &EndpointBaseProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalEndpointBaseUpdatePropertiesClassification(rawMsg json.RawMessage) (EndpointBaseUpdatePropertiesClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b EndpointBaseUpdatePropertiesClassification
	switch m["endpointType"] {
	case string(EndpointTypeAzureStorageBlobContainer):
		b = &AzureStorageBlobContainerEndpointUpdateProperties{}
	case string(EndpointTypeAzureStorageSmbFileShare):
		b = &AzureStorageSmbFileShareEndpointUpdateProperties{}
	case string(EndpointTypeNfsMount):
		b = &NfsMountEndpointUpdateProperties{}
	case string(EndpointTypeSmbMount):
		b = &SmbMountEndpointUpdateProperties{}
	default:
		b = &EndpointBaseUpdateProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}
