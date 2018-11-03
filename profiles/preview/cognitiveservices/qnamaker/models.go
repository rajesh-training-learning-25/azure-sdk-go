// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package qnamaker

import original "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v4.0/qnamaker"

type BaseClient = original.BaseClient
type Code = original.Code

const (
	BadArgument       Code = original.BadArgument
	EndpointKeysError Code = original.EndpointKeysError
	ExtractionFailure Code = original.ExtractionFailure
	Forbidden         Code = original.Forbidden
	KbNotFound        Code = original.KbNotFound
	NotFound          Code = original.NotFound
	OperationNotFound Code = original.OperationNotFound
	QnaRuntimeError   Code = original.QnaRuntimeError
	QuotaExceeded     Code = original.QuotaExceeded
	ServiceError      Code = original.ServiceError
	SKULimitExceeded  Code = original.SKULimitExceeded
	Unauthorized      Code = original.Unauthorized
	Unspecified       Code = original.Unspecified
	ValidationFailure Code = original.ValidationFailure
)

type OperationState = original.OperationState

const (
	Failed     OperationState = original.Failed
	NotStarted OperationState = original.NotStarted
	Running    OperationState = original.Running
	Succeeded  OperationState = original.Succeeded
)

type AlterationsDTO = original.AlterationsDTO
type CreateKbDTO = original.CreateKbDTO
type CreateKbInputDTO = original.CreateKbInputDTO
type DeleteKbContentsDTO = original.DeleteKbContentsDTO
type EndpointKeysDTO = original.EndpointKeysDTO
type Error = original.Error
type ErrorResponse = original.ErrorResponse
type ErrorResponseError = original.ErrorResponseError
type FileDTO = original.FileDTO
type InnerErrorModel = original.InnerErrorModel
type KnowledgebaseDTO = original.KnowledgebaseDTO
type KnowledgebasesDTO = original.KnowledgebasesDTO
type MetadataDTO = original.MetadataDTO
type Operation = original.Operation
type QnADocumentsDTO = original.QnADocumentsDTO
type QnADTO = original.QnADTO
type ReplaceKbDTO = original.ReplaceKbDTO
type UpdateKbContentsDTO = original.UpdateKbContentsDTO
type UpdateKbOperationDTO = original.UpdateKbOperationDTO
type UpdateKbOperationDTOAdd = original.UpdateKbOperationDTOAdd
type UpdateKbOperationDTODelete = original.UpdateKbOperationDTODelete
type UpdateKbOperationDTOUpdate = original.UpdateKbOperationDTOUpdate
type UpdateMetadataDTO = original.UpdateMetadataDTO
type UpdateQnaDTO = original.UpdateQnaDTO
type UpdateQnaDTOMetadata = original.UpdateQnaDTOMetadata
type UpdateQnaDTOQuestions = original.UpdateQnaDTOQuestions
type UpdateQuestionsDTO = original.UpdateQuestionsDTO
type WordAlterationsDTO = original.WordAlterationsDTO

func New(endpoint string) BaseClient {
	return original.New(endpoint)
}
func NewWithoutDefaults(endpoint string) BaseClient {
	return original.NewWithoutDefaults(endpoint)
}
func PossibleCodeValues() []Code {
	return original.PossibleCodeValues()
}
func PossibleOperationStateValues() []OperationState {
	return original.PossibleOperationStateValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
