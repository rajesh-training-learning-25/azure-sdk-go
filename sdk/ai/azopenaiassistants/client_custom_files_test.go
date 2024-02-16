//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azopenaiassistants_test

import (
	"context"
	"io"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/sdk/ai/azopenaiassistants"
)

func TestDownloadFileContent(t *testing.T) {
	args := runThreadArgs{
		Assistant: azopenaiassistants.AssistantCreationBody{
			DeploymentName: &assistantsModel,
			Instructions:   to.Ptr("You are a helpful image generating assistant"),
			Tools: []azopenaiassistants.ToolDefinitionClassification{
				&azopenaiassistants.CodeInterpreterToolDefinition{},
			},
		},
		Thread: azopenaiassistants.CreateAndRunThreadOptions{
			Thread: &azopenaiassistants.AssistantThreadCreationOptions{
				Messages: []azopenaiassistants.ThreadInitializationMessage{
					{
						Role:    to.Ptr(azopenaiassistants.MessageRoleUser),
						Content: to.Ptr("Can you draw an image of two boxes, connected by a line, as a PNG file?"),
					},
				},
			},
		},
	}

	client, messages := mustRunThread(context.Background(), t, args)
	fileFound := false

	for _, m := range messages {
		// the assistants reply should contain a file ID for an image to download
		for _, c := range m.Content {
			switch v := c.(type) {
			case *azopenaiassistants.MessageImageFileContent:
				resp, err := client.GetFileContent(context.Background(), *v.ImageFile.FileID, nil)
				require.NoError(t, err)

				defer func() {
					err := resp.Content.Close()
					require.NoError(t, err)
				}()

				fileBytes, err := io.ReadAll(resp.Content)
				require.NoError(t, err)
				require.NotEmpty(t, fileBytes)
				fileFound = true

				t.Logf("[%s] image file ID: %s, file is %d bytes", *m.Role, *v.ImageFile.FileID, len(fileBytes))
				break
			case *azopenaiassistants.MessageTextContent:
				t.Logf("[%s] %s", *m.Role, *v.Text.Value)
				break
			}
		}
	}

	require.True(t, fileFound)
}
