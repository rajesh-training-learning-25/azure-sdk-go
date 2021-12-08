//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicessiterecovery

import "encoding/json"

func unmarshalResourceCertificateDetailsClassification(rawMsg json.RawMessage) (ResourceCertificateDetailsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ResourceCertificateDetailsClassification
	switch m["authType"] {
	case "AccessControlService":
		b = &ResourceCertificateAndAcsDetails{}
	case "AzureActiveDirectory":
		b = &ResourceCertificateAndAADDetails{}
	default:
		b = &ResourceCertificateDetails{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalResourceCertificateDetailsClassificationArray(rawMsg json.RawMessage) ([]ResourceCertificateDetailsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ResourceCertificateDetailsClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalResourceCertificateDetailsClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalResourceCertificateDetailsClassificationMap(rawMsg json.RawMessage) (map[string]ResourceCertificateDetailsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]ResourceCertificateDetailsClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalResourceCertificateDetailsClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}
