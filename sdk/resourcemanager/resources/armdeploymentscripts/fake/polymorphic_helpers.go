//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armdeploymentscripts/v2"
)

func unmarshalDeploymentScriptClassification(rawMsg json.RawMessage) (armdeploymentscripts.DeploymentScriptClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b armdeploymentscripts.DeploymentScriptClassification
	switch m["kind"] {
	case string(armdeploymentscripts.ScriptTypeAzureCLI):
		b = &armdeploymentscripts.AzureCliScript{}
	case string(armdeploymentscripts.ScriptTypeAzurePowerShell):
		b = &armdeploymentscripts.AzurePowerShellScript{}
	default:
		b = &armdeploymentscripts.DeploymentScript{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}
