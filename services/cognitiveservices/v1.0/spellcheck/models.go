package spellcheck

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go//services/cognitiveservices/v1.0/spellcheck"

// ActionType enumerates the values for action type.
type ActionType string

const (
	// Edit ...
	Edit ActionType = "Edit"
	// Load ...
	Load ActionType = "Load"
)

// PossibleActionTypeValues returns an array of possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{Edit, Load}
}

// ErrorCode enumerates the values for error code.
type ErrorCode string

const (
	// InsufficientAuthorization ...
	InsufficientAuthorization ErrorCode = "InsufficientAuthorization"
	// InvalidAuthorization ...
	InvalidAuthorization ErrorCode = "InvalidAuthorization"
	// InvalidRequest ...
	InvalidRequest ErrorCode = "InvalidRequest"
	// None ...
	None ErrorCode = "None"
	// RateLimitExceeded ...
	RateLimitExceeded ErrorCode = "RateLimitExceeded"
	// ServerError ...
	ServerError ErrorCode = "ServerError"
)

// PossibleErrorCodeValues returns an array of possible values for the ErrorCode const type.
func PossibleErrorCodeValues() []ErrorCode {
	return []ErrorCode{InsufficientAuthorization, InvalidAuthorization, InvalidRequest, None, RateLimitExceeded, ServerError}
}

// ErrorSubCode enumerates the values for error sub code.
type ErrorSubCode string

const (
	// AuthorizationDisabled ...
	AuthorizationDisabled ErrorSubCode = "AuthorizationDisabled"
	// AuthorizationExpired ...
	AuthorizationExpired ErrorSubCode = "AuthorizationExpired"
	// AuthorizationMissing ...
	AuthorizationMissing ErrorSubCode = "AuthorizationMissing"
	// AuthorizationRedundancy ...
	AuthorizationRedundancy ErrorSubCode = "AuthorizationRedundancy"
	// Blocked ...
	Blocked ErrorSubCode = "Blocked"
	// HTTPNotAllowed ...
	HTTPNotAllowed ErrorSubCode = "HttpNotAllowed"
	// NotImplemented ...
	NotImplemented ErrorSubCode = "NotImplemented"
	// ParameterInvalidValue ...
	ParameterInvalidValue ErrorSubCode = "ParameterInvalidValue"
	// ParameterMissing ...
	ParameterMissing ErrorSubCode = "ParameterMissing"
	// ResourceError ...
	ResourceError ErrorSubCode = "ResourceError"
	// UnexpectedError ...
	UnexpectedError ErrorSubCode = "UnexpectedError"
)

// PossibleErrorSubCodeValues returns an array of possible values for the ErrorSubCode const type.
func PossibleErrorSubCodeValues() []ErrorSubCode {
	return []ErrorSubCode{AuthorizationDisabled, AuthorizationExpired, AuthorizationMissing, AuthorizationRedundancy, Blocked, HTTPNotAllowed, NotImplemented, ParameterInvalidValue, ParameterMissing, ResourceError, UnexpectedError}
}

// ErrorType enumerates the values for error type.
type ErrorType string

const (
	// RepeatedToken ...
	RepeatedToken ErrorType = "RepeatedToken"
	// UnknownToken ...
	UnknownToken ErrorType = "UnknownToken"
)

// PossibleErrorTypeValues returns an array of possible values for the ErrorType const type.
func PossibleErrorTypeValues() []ErrorType {
	return []ErrorType{RepeatedToken, UnknownToken}
}

// Type enumerates the values for type.
type Type string

const (
	// TypeAnswer ...
	TypeAnswer Type = "Answer"
	// TypeErrorResponse ...
	TypeErrorResponse Type = "ErrorResponse"
	// TypeIdentifiable ...
	TypeIdentifiable Type = "Identifiable"
	// TypeResponse ...
	TypeResponse Type = "Response"
	// TypeResponseBase ...
	TypeResponseBase Type = "ResponseBase"
	// TypeSpellCheck ...
	TypeSpellCheck Type = "SpellCheck"
)

// PossibleTypeValues returns an array of possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{TypeAnswer, TypeErrorResponse, TypeIdentifiable, TypeResponse, TypeResponseBase, TypeSpellCheck}
}

// BasicAnswer ...
type BasicAnswer interface {
	AsSpellCheck() (*SpellCheck, bool)
	AsAnswer() (*Answer, bool)
}

// Answer ...
type Answer struct {
	// ID - A String identifier.
	ID *string `json:"id,omitempty"`
	// Type - Possible values include: 'TypeResponseBase', 'TypeSpellCheck', 'TypeAnswer', 'TypeResponse', 'TypeIdentifiable', 'TypeErrorResponse'
	Type Type `json:"_type,omitempty"`
}

func unmarshalBasicAnswer(body []byte) (BasicAnswer, error) {
	var m map[string]interface{}
	err := json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}

	switch m["_type"] {
	case string(TypeSpellCheck):
		var sc SpellCheck
		err := json.Unmarshal(body, &sc)
		return sc, err
	default:
		var a Answer
		err := json.Unmarshal(body, &a)
		return a, err
	}
}
func unmarshalBasicAnswerArray(body []byte) ([]BasicAnswer, error) {
	var rawMessages []*json.RawMessage
	err := json.Unmarshal(body, &rawMessages)
	if err != nil {
		return nil, err
	}

	aArray := make([]BasicAnswer, len(rawMessages))

	for index, rawMessage := range rawMessages {
		a, err := unmarshalBasicAnswer(*rawMessage)
		if err != nil {
			return nil, err
		}
		aArray[index] = a
	}
	return aArray, nil
}

// MarshalJSON is the custom marshaler for Answer.
func (a Answer) MarshalJSON() ([]byte, error) {
	a.Type = TypeAnswer
	objectMap := make(map[string]interface{})
	if a.ID != nil {
		objectMap["id"] = a.ID
	}
	if a.Type != "" {
		objectMap["_type"] = a.Type
	}
	return json.Marshal(objectMap)
}

// AsSpellCheck is the BasicResponseBase implementation for Answer.
func (a Answer) AsSpellCheck() (*SpellCheck, bool) {
	return nil, false
}

// AsAnswer is the BasicResponseBase implementation for Answer.
func (a Answer) AsAnswer() (*Answer, bool) {
	return &a, true
}

// AsBasicAnswer is the BasicResponseBase implementation for Answer.
func (a Answer) AsBasicAnswer() (BasicAnswer, bool) {
	return &a, true
}

// AsResponse is the BasicResponseBase implementation for Answer.
func (a Answer) AsResponse() (*Response, bool) {
	return nil, false
}

// AsBasicResponse is the BasicResponseBase implementation for Answer.
func (a Answer) AsBasicResponse() (BasicResponse, bool) {
	return &a, true
}

// AsIdentifiable is the BasicResponseBase implementation for Answer.
func (a Answer) AsIdentifiable() (*Identifiable, bool) {
	return nil, false
}

// AsBasicIdentifiable is the BasicResponseBase implementation for Answer.
func (a Answer) AsBasicIdentifiable() (BasicIdentifiable, bool) {
	return &a, true
}

// AsErrorResponse is the BasicResponseBase implementation for Answer.
func (a Answer) AsErrorResponse() (*ErrorResponse, bool) {
	return nil, false
}

// AsResponseBase is the BasicResponseBase implementation for Answer.
func (a Answer) AsResponseBase() (*ResponseBase, bool) {
	return nil, false
}

// AsBasicResponseBase is the BasicResponseBase implementation for Answer.
func (a Answer) AsBasicResponseBase() (BasicResponseBase, bool) {
	return &a, true
}

// Error defines the error that occurred.
type Error struct {
	// Code - The error code that identifies the category of error. Possible values include: 'None', 'ServerError', 'InvalidRequest', 'RateLimitExceeded', 'InvalidAuthorization', 'InsufficientAuthorization'
	Code ErrorCode `json:"code,omitempty"`
	// SubCode - The error code that further helps to identify the error. Possible values include: 'UnexpectedError', 'ResourceError', 'NotImplemented', 'ParameterMissing', 'ParameterInvalidValue', 'HTTPNotAllowed', 'Blocked', 'AuthorizationMissing', 'AuthorizationRedundancy', 'AuthorizationDisabled', 'AuthorizationExpired'
	SubCode ErrorSubCode `json:"subCode,omitempty"`
	// Message - A description of the error.
	Message *string `json:"message,omitempty"`
	// MoreDetails - A description that provides additional information about the error.
	MoreDetails *string `json:"moreDetails,omitempty"`
	// Parameter - The parameter in the request that caused the error.
	Parameter *string `json:"parameter,omitempty"`
	// Value - The parameter's value in the request that was not valid.
	Value *string `json:"value,omitempty"`
}

// ErrorResponse the top-level response that represents a failed request.
type ErrorResponse struct {
	// Errors - A list of errors that describe the reasons why the request failed.
	Errors *[]Error `json:"errors,omitempty"`
	// ID - A String identifier.
	ID *string `json:"id,omitempty"`
	// Type - Possible values include: 'TypeResponseBase', 'TypeSpellCheck', 'TypeAnswer', 'TypeResponse', 'TypeIdentifiable', 'TypeErrorResponse'
	Type Type `json:"_type,omitempty"`
}

// MarshalJSON is the custom marshaler for ErrorResponse.
func (er ErrorResponse) MarshalJSON() ([]byte, error) {
	er.Type = TypeErrorResponse
	objectMap := make(map[string]interface{})
	if er.Errors != nil {
		objectMap["errors"] = er.Errors
	}
	if er.ID != nil {
		objectMap["id"] = er.ID
	}
	if er.Type != "" {
		objectMap["_type"] = er.Type
	}
	return json.Marshal(objectMap)
}

// AsSpellCheck is the BasicResponseBase implementation for ErrorResponse.
func (er ErrorResponse) AsSpellCheck() (*SpellCheck, bool) {
	return nil, false
}

// AsAnswer is the BasicResponseBase implementation for ErrorResponse.
func (er ErrorResponse) AsAnswer() (*Answer, bool) {
	return nil, false
}

// AsBasicAnswer is the BasicResponseBase implementation for ErrorResponse.
func (er ErrorResponse) AsBasicAnswer() (BasicAnswer, bool) {
	return nil, false
}

// AsResponse is the BasicResponseBase implementation for ErrorResponse.
func (er ErrorResponse) AsResponse() (*Response, bool) {
	return nil, false
}

// AsBasicResponse is the BasicResponseBase implementation for ErrorResponse.
func (er ErrorResponse) AsBasicResponse() (BasicResponse, bool) {
	return &er, true
}

// AsIdentifiable is the BasicResponseBase implementation for ErrorResponse.
func (er ErrorResponse) AsIdentifiable() (*Identifiable, bool) {
	return nil, false
}

// AsBasicIdentifiable is the BasicResponseBase implementation for ErrorResponse.
func (er ErrorResponse) AsBasicIdentifiable() (BasicIdentifiable, bool) {
	return &er, true
}

// AsErrorResponse is the BasicResponseBase implementation for ErrorResponse.
func (er ErrorResponse) AsErrorResponse() (*ErrorResponse, bool) {
	return &er, true
}

// AsResponseBase is the BasicResponseBase implementation for ErrorResponse.
func (er ErrorResponse) AsResponseBase() (*ResponseBase, bool) {
	return nil, false
}

// AsBasicResponseBase is the BasicResponseBase implementation for ErrorResponse.
func (er ErrorResponse) AsBasicResponseBase() (BasicResponseBase, bool) {
	return &er, true
}

// BasicIdentifiable defines the identity of a resource.
type BasicIdentifiable interface {
	AsSpellCheck() (*SpellCheck, bool)
	AsAnswer() (*Answer, bool)
	AsBasicAnswer() (BasicAnswer, bool)
	AsResponse() (*Response, bool)
	AsBasicResponse() (BasicResponse, bool)
	AsErrorResponse() (*ErrorResponse, bool)
	AsIdentifiable() (*Identifiable, bool)
}

// Identifiable defines the identity of a resource.
type Identifiable struct {
	// ID - A String identifier.
	ID *string `json:"id,omitempty"`
	// Type - Possible values include: 'TypeResponseBase', 'TypeSpellCheck', 'TypeAnswer', 'TypeResponse', 'TypeIdentifiable', 'TypeErrorResponse'
	Type Type `json:"_type,omitempty"`
}

func unmarshalBasicIdentifiable(body []byte) (BasicIdentifiable, error) {
	var m map[string]interface{}
	err := json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}

	switch m["_type"] {
	case string(TypeSpellCheck):
		var sc SpellCheck
		err := json.Unmarshal(body, &sc)
		return sc, err
	case string(TypeAnswer):
		var a Answer
		err := json.Unmarshal(body, &a)
		return a, err
	case string(TypeResponse):
		var r Response
		err := json.Unmarshal(body, &r)
		return r, err
	case string(TypeErrorResponse):
		var er ErrorResponse
		err := json.Unmarshal(body, &er)
		return er, err
	default:
		var i Identifiable
		err := json.Unmarshal(body, &i)
		return i, err
	}
}
func unmarshalBasicIdentifiableArray(body []byte) ([]BasicIdentifiable, error) {
	var rawMessages []*json.RawMessage
	err := json.Unmarshal(body, &rawMessages)
	if err != nil {
		return nil, err
	}

	iArray := make([]BasicIdentifiable, len(rawMessages))

	for index, rawMessage := range rawMessages {
		i, err := unmarshalBasicIdentifiable(*rawMessage)
		if err != nil {
			return nil, err
		}
		iArray[index] = i
	}
	return iArray, nil
}

// MarshalJSON is the custom marshaler for Identifiable.
func (i Identifiable) MarshalJSON() ([]byte, error) {
	i.Type = TypeIdentifiable
	objectMap := make(map[string]interface{})
	if i.ID != nil {
		objectMap["id"] = i.ID
	}
	if i.Type != "" {
		objectMap["_type"] = i.Type
	}
	return json.Marshal(objectMap)
}

// AsSpellCheck is the BasicResponseBase implementation for Identifiable.
func (i Identifiable) AsSpellCheck() (*SpellCheck, bool) {
	return nil, false
}

// AsAnswer is the BasicResponseBase implementation for Identifiable.
func (i Identifiable) AsAnswer() (*Answer, bool) {
	return nil, false
}

// AsBasicAnswer is the BasicResponseBase implementation for Identifiable.
func (i Identifiable) AsBasicAnswer() (BasicAnswer, bool) {
	return nil, false
}

// AsResponse is the BasicResponseBase implementation for Identifiable.
func (i Identifiable) AsResponse() (*Response, bool) {
	return nil, false
}

// AsBasicResponse is the BasicResponseBase implementation for Identifiable.
func (i Identifiable) AsBasicResponse() (BasicResponse, bool) {
	return nil, false
}

// AsIdentifiable is the BasicResponseBase implementation for Identifiable.
func (i Identifiable) AsIdentifiable() (*Identifiable, bool) {
	return &i, true
}

// AsBasicIdentifiable is the BasicResponseBase implementation for Identifiable.
func (i Identifiable) AsBasicIdentifiable() (BasicIdentifiable, bool) {
	return &i, true
}

// AsErrorResponse is the BasicResponseBase implementation for Identifiable.
func (i Identifiable) AsErrorResponse() (*ErrorResponse, bool) {
	return nil, false
}

// AsResponseBase is the BasicResponseBase implementation for Identifiable.
func (i Identifiable) AsResponseBase() (*ResponseBase, bool) {
	return nil, false
}

// AsBasicResponseBase is the BasicResponseBase implementation for Identifiable.
func (i Identifiable) AsBasicResponseBase() (BasicResponseBase, bool) {
	return &i, true
}

// BasicResponse defines a response. All schemas that could be returned at the root of a response should inherit from
// this
type BasicResponse interface {
	AsSpellCheck() (*SpellCheck, bool)
	AsAnswer() (*Answer, bool)
	AsBasicAnswer() (BasicAnswer, bool)
	AsErrorResponse() (*ErrorResponse, bool)
	AsResponse() (*Response, bool)
}

// Response defines a response. All schemas that could be returned at the root of a response should inherit
// from this
type Response struct {
	// ID - A String identifier.
	ID *string `json:"id,omitempty"`
	// Type - Possible values include: 'TypeResponseBase', 'TypeSpellCheck', 'TypeAnswer', 'TypeResponse', 'TypeIdentifiable', 'TypeErrorResponse'
	Type Type `json:"_type,omitempty"`
}

func unmarshalBasicResponse(body []byte) (BasicResponse, error) {
	var m map[string]interface{}
	err := json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}

	switch m["_type"] {
	case string(TypeSpellCheck):
		var sc SpellCheck
		err := json.Unmarshal(body, &sc)
		return sc, err
	case string(TypeAnswer):
		var a Answer
		err := json.Unmarshal(body, &a)
		return a, err
	case string(TypeErrorResponse):
		var er ErrorResponse
		err := json.Unmarshal(body, &er)
		return er, err
	default:
		var r Response
		err := json.Unmarshal(body, &r)
		return r, err
	}
}
func unmarshalBasicResponseArray(body []byte) ([]BasicResponse, error) {
	var rawMessages []*json.RawMessage
	err := json.Unmarshal(body, &rawMessages)
	if err != nil {
		return nil, err
	}

	rArray := make([]BasicResponse, len(rawMessages))

	for index, rawMessage := range rawMessages {
		r, err := unmarshalBasicResponse(*rawMessage)
		if err != nil {
			return nil, err
		}
		rArray[index] = r
	}
	return rArray, nil
}

// MarshalJSON is the custom marshaler for Response.
func (r Response) MarshalJSON() ([]byte, error) {
	r.Type = TypeResponse
	objectMap := make(map[string]interface{})
	if r.ID != nil {
		objectMap["id"] = r.ID
	}
	if r.Type != "" {
		objectMap["_type"] = r.Type
	}
	return json.Marshal(objectMap)
}

// AsSpellCheck is the BasicResponseBase implementation for Response.
func (r Response) AsSpellCheck() (*SpellCheck, bool) {
	return nil, false
}

// AsAnswer is the BasicResponseBase implementation for Response.
func (r Response) AsAnswer() (*Answer, bool) {
	return nil, false
}

// AsBasicAnswer is the BasicResponseBase implementation for Response.
func (r Response) AsBasicAnswer() (BasicAnswer, bool) {
	return nil, false
}

// AsResponse is the BasicResponseBase implementation for Response.
func (r Response) AsResponse() (*Response, bool) {
	return &r, true
}

// AsBasicResponse is the BasicResponseBase implementation for Response.
func (r Response) AsBasicResponse() (BasicResponse, bool) {
	return &r, true
}

// AsIdentifiable is the BasicResponseBase implementation for Response.
func (r Response) AsIdentifiable() (*Identifiable, bool) {
	return nil, false
}

// AsBasicIdentifiable is the BasicResponseBase implementation for Response.
func (r Response) AsBasicIdentifiable() (BasicIdentifiable, bool) {
	return &r, true
}

// AsErrorResponse is the BasicResponseBase implementation for Response.
func (r Response) AsErrorResponse() (*ErrorResponse, bool) {
	return nil, false
}

// AsResponseBase is the BasicResponseBase implementation for Response.
func (r Response) AsResponseBase() (*ResponseBase, bool) {
	return nil, false
}

// AsBasicResponseBase is the BasicResponseBase implementation for Response.
func (r Response) AsBasicResponseBase() (BasicResponseBase, bool) {
	return &r, true
}

// BasicResponseBase ...
type BasicResponseBase interface {
	AsSpellCheck() (*SpellCheck, bool)
	AsAnswer() (*Answer, bool)
	AsBasicAnswer() (BasicAnswer, bool)
	AsResponse() (*Response, bool)
	AsBasicResponse() (BasicResponse, bool)
	AsIdentifiable() (*Identifiable, bool)
	AsBasicIdentifiable() (BasicIdentifiable, bool)
	AsErrorResponse() (*ErrorResponse, bool)
	AsResponseBase() (*ResponseBase, bool)
}

// ResponseBase ...
type ResponseBase struct {
	// Type - Possible values include: 'TypeResponseBase', 'TypeSpellCheck', 'TypeAnswer', 'TypeResponse', 'TypeIdentifiable', 'TypeErrorResponse'
	Type Type `json:"_type,omitempty"`
}

func unmarshalBasicResponseBase(body []byte) (BasicResponseBase, error) {
	var m map[string]interface{}
	err := json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}

	switch m["_type"] {
	case string(TypeSpellCheck):
		var sc SpellCheck
		err := json.Unmarshal(body, &sc)
		return sc, err
	case string(TypeAnswer):
		var a Answer
		err := json.Unmarshal(body, &a)
		return a, err
	case string(TypeResponse):
		var r Response
		err := json.Unmarshal(body, &r)
		return r, err
	case string(TypeIdentifiable):
		var i Identifiable
		err := json.Unmarshal(body, &i)
		return i, err
	case string(TypeErrorResponse):
		var er ErrorResponse
		err := json.Unmarshal(body, &er)
		return er, err
	default:
		var rb ResponseBase
		err := json.Unmarshal(body, &rb)
		return rb, err
	}
}
func unmarshalBasicResponseBaseArray(body []byte) ([]BasicResponseBase, error) {
	var rawMessages []*json.RawMessage
	err := json.Unmarshal(body, &rawMessages)
	if err != nil {
		return nil, err
	}

	rbArray := make([]BasicResponseBase, len(rawMessages))

	for index, rawMessage := range rawMessages {
		rb, err := unmarshalBasicResponseBase(*rawMessage)
		if err != nil {
			return nil, err
		}
		rbArray[index] = rb
	}
	return rbArray, nil
}

// MarshalJSON is the custom marshaler for ResponseBase.
func (rb ResponseBase) MarshalJSON() ([]byte, error) {
	rb.Type = TypeResponseBase
	objectMap := make(map[string]interface{})
	if rb.Type != "" {
		objectMap["_type"] = rb.Type
	}
	return json.Marshal(objectMap)
}

// AsSpellCheck is the BasicResponseBase implementation for ResponseBase.
func (rb ResponseBase) AsSpellCheck() (*SpellCheck, bool) {
	return nil, false
}

// AsAnswer is the BasicResponseBase implementation for ResponseBase.
func (rb ResponseBase) AsAnswer() (*Answer, bool) {
	return nil, false
}

// AsBasicAnswer is the BasicResponseBase implementation for ResponseBase.
func (rb ResponseBase) AsBasicAnswer() (BasicAnswer, bool) {
	return nil, false
}

// AsResponse is the BasicResponseBase implementation for ResponseBase.
func (rb ResponseBase) AsResponse() (*Response, bool) {
	return nil, false
}

// AsBasicResponse is the BasicResponseBase implementation for ResponseBase.
func (rb ResponseBase) AsBasicResponse() (BasicResponse, bool) {
	return nil, false
}

// AsIdentifiable is the BasicResponseBase implementation for ResponseBase.
func (rb ResponseBase) AsIdentifiable() (*Identifiable, bool) {
	return nil, false
}

// AsBasicIdentifiable is the BasicResponseBase implementation for ResponseBase.
func (rb ResponseBase) AsBasicIdentifiable() (BasicIdentifiable, bool) {
	return nil, false
}

// AsErrorResponse is the BasicResponseBase implementation for ResponseBase.
func (rb ResponseBase) AsErrorResponse() (*ErrorResponse, bool) {
	return nil, false
}

// AsResponseBase is the BasicResponseBase implementation for ResponseBase.
func (rb ResponseBase) AsResponseBase() (*ResponseBase, bool) {
	return &rb, true
}

// AsBasicResponseBase is the BasicResponseBase implementation for ResponseBase.
func (rb ResponseBase) AsBasicResponseBase() (BasicResponseBase, bool) {
	return &rb, true
}

// SpellCheck ...
type SpellCheck struct {
	autorest.Response `json:"-"`
	FlaggedTokens     *[]SpellingFlaggedToken `json:"flaggedTokens,omitempty"`
	// ID - A String identifier.
	ID *string `json:"id,omitempty"`
	// Type - Possible values include: 'TypeResponseBase', 'TypeSpellCheck', 'TypeAnswer', 'TypeResponse', 'TypeIdentifiable', 'TypeErrorResponse'
	Type Type `json:"_type,omitempty"`
}

// MarshalJSON is the custom marshaler for SpellCheck.
func (sc SpellCheck) MarshalJSON() ([]byte, error) {
	sc.Type = TypeSpellCheck
	objectMap := make(map[string]interface{})
	if sc.FlaggedTokens != nil {
		objectMap["flaggedTokens"] = sc.FlaggedTokens
	}
	if sc.ID != nil {
		objectMap["id"] = sc.ID
	}
	if sc.Type != "" {
		objectMap["_type"] = sc.Type
	}
	return json.Marshal(objectMap)
}

// AsSpellCheck is the BasicResponseBase implementation for SpellCheck.
func (sc SpellCheck) AsSpellCheck() (*SpellCheck, bool) {
	return &sc, true
}

// AsAnswer is the BasicResponseBase implementation for SpellCheck.
func (sc SpellCheck) AsAnswer() (*Answer, bool) {
	return nil, false
}

// AsBasicAnswer is the BasicResponseBase implementation for SpellCheck.
func (sc SpellCheck) AsBasicAnswer() (BasicAnswer, bool) {
	return &sc, true
}

// AsResponse is the BasicResponseBase implementation for SpellCheck.
func (sc SpellCheck) AsResponse() (*Response, bool) {
	return nil, false
}

// AsBasicResponse is the BasicResponseBase implementation for SpellCheck.
func (sc SpellCheck) AsBasicResponse() (BasicResponse, bool) {
	return &sc, true
}

// AsIdentifiable is the BasicResponseBase implementation for SpellCheck.
func (sc SpellCheck) AsIdentifiable() (*Identifiable, bool) {
	return nil, false
}

// AsBasicIdentifiable is the BasicResponseBase implementation for SpellCheck.
func (sc SpellCheck) AsBasicIdentifiable() (BasicIdentifiable, bool) {
	return &sc, true
}

// AsErrorResponse is the BasicResponseBase implementation for SpellCheck.
func (sc SpellCheck) AsErrorResponse() (*ErrorResponse, bool) {
	return nil, false
}

// AsResponseBase is the BasicResponseBase implementation for SpellCheck.
func (sc SpellCheck) AsResponseBase() (*ResponseBase, bool) {
	return nil, false
}

// AsBasicResponseBase is the BasicResponseBase implementation for SpellCheck.
func (sc SpellCheck) AsBasicResponseBase() (BasicResponseBase, bool) {
	return &sc, true
}

// SpellingFlaggedToken ...
type SpellingFlaggedToken struct {
	Offset *int32  `json:"offset,omitempty"`
	Token  *string `json:"token,omitempty"`
	// Type - Possible values include: 'UnknownToken', 'RepeatedToken'
	Type          ErrorType                  `json:"type,omitempty"`
	Suggestions   *[]SpellingTokenSuggestion `json:"suggestions,omitempty"`
	PingURLSuffix *string                    `json:"pingUrlSuffix,omitempty"`
}

// SpellingTokenSuggestion ...
type SpellingTokenSuggestion struct {
	Suggestion    *string  `json:"suggestion,omitempty"`
	Score         *float64 `json:"score,omitempty"`
	PingURLSuffix *string  `json:"pingUrlSuffix,omitempty"`
}
