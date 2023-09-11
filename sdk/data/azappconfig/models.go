//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azappconfig

import "time"

// AddSettingOptions contains the optional parameters for the Client.AddSetting method.
type AddSettingOptions struct {
	// Used to perform an operation only if the targeted resource's etag matches the value provided.
	IfMatch *string
	// Used to perform an operation only if the targeted resource's etag does not match the value provided.
	IfNoneMatch *string
	// The label of the key-value to create.
	Label *string
}

// CheckKeyValueOptions contains the optional parameters for the Client.CheckKeyValue method.
type CheckKeyValueOptions struct {
	// Requests the server to respond with the state of the resource at the specified time.
	AcceptDatetime *string
	// Used to perform an operation only if the targeted resource's etag matches the value provided.
	IfMatch *string
	// Used to perform an operation only if the targeted resource's etag does not match the value provided.
	IfNoneMatch *string
	// The label of the key-value to retrieve.
	Label *string
	// Used to select what fields are present in the returned resource(s).
	Select []KeyValueFields
}

// CheckKeyValuesOptions contains the optional parameters for the Client.CheckKeyValues method.
type CheckKeyValuesOptions struct {
	// Requests the server to respond with the state of the resource at the specified time.
	AcceptDatetime *string
	// Instructs the server to return elements that appear after the element referred to by the specified token.
	After *string
	// A filter used to match keys.
	Key *string
	// A filter used to match labels
	Label *string
	// Used to select what fields are present in the returned resource(s).
	Select []KeyValueFields
}

// CheckKeysOptions contains the optional parameters for the Client.CheckKeys method.
type CheckKeysOptions struct {
	// Requests the server to respond with the state of the resource at the specified time.
	AcceptDatetime *string
	// Instructs the server to return elements that appear after the element referred to by the specified token.
	After *string
	// A filter for the name of the returned keys.
	Name *string
}

// CheckLabelsOptions contains the optional parameters for the Client.CheckLabels method.
type CheckLabelsOptions struct {
	// Requests the server to respond with the state of the resource at the specified time.
	AcceptDatetime *string
	// Instructs the server to return elements that appear after the element referred to by the specified token.
	After *string
	// A filter for the name of the returned labels.
	Name *string
	// Used to select what fields are present in the returned resource(s).
	Select []LabelFields
}

// CheckRevisionsOptions contains the optional parameters for the Client.CheckRevisions method.
type CheckRevisionsOptions struct {
	// Requests the server to respond with the state of the resource at the specified time.
	AcceptDatetime *string
	// Instructs the server to return elements that appear after the element referred to by the specified token.
	After *string
	// A filter used to match keys.
	Key *string
	// A filter used to match labels
	Label *string
	// Used to select what fields are present in the returned resource(s).
	Select []KeyValueFields
}

// DeleteKeyValueOptions contains the optional parameters for the Client.DeleteKeyValue method.
type DeleteKeyValueOptions struct {
	// Used to perform an operation only if the targeted resource's etag matches the value provided.
	IfMatch *string
	// The label of the key-value to delete.
	Label *string
}

// DeleteLockOptions contains the optional parameters for the Client.DeleteLock method.
type DeleteLockOptions struct {
	// Used to perform an operation only if the targeted resource's etag matches the value provided.
	IfMatch *string
	// Used to perform an operation only if the targeted resource's etag does not match the value provided.
	IfNoneMatch *string
	// The label, if any, of the key-value to unlock.
	Label *string
}

// GetKeyValuesOptions contains the optional parameters for the Client.NewGetKeyValuesPager method.
type GetKeyValuesOptions struct {
	// Requests the server to respond with the state of the resource at the specified time.
	AcceptDatetime *string
	// Instructs the server to return elements that appear after the element referred to by the specified token.
	After *string
	// A filter used to match keys.
	Key *string
	// A filter used to match labels
	Label *string
	// Used to select what fields are present in the returned resource(s).
	Select []KeyValueFields
}

// GetKeysOptions contains the optional parameters for the Client.NewGetKeysPager method.
type GetKeysOptions struct {
	// Requests the server to respond with the state of the resource at the specified time.
	AcceptDatetime *string
	// Instructs the server to return elements that appear after the element referred to by the specified token.
	After *string
	// A filter for the name of the returned keys.
	Name *string
}

// GetLabelsOptions contains the optional parameters for the Client.NewGetLabelsPager method.
type GetLabelsOptions struct {
	// Requests the server to respond with the state of the resource at the specified time.
	AcceptDatetime *string
	// Instructs the server to return elements that appear after the element referred to by the specified token.
	After *string
	// A filter for the name of the returned labels.
	Name *string
	// Used to select what fields are present in the returned resource(s).
	Select []LabelFields
}

// GetRevisionsOptions contains the optional parameters for the Client.NewGetRevisionsPager method.
type GetRevisionsOptions struct {
	// Requests the server to respond with the state of the resource at the specified time.
	AcceptDatetime *string
	// Instructs the server to return elements that appear after the element referred to by the specified token.
	After *string
	// A filter used to match keys.
	Key *string
	// A filter used to match labels
	Label *string
	// Used to select what fields are present in the returned resource(s).
	Select []KeyValueFields
}

// GetSettingOptions contains the optional parameters for the Client.GetSetting method.
type GetSettingOptions struct {
	// Requests the server to respond with the state of the resource at the specified time.
	AcceptDatetime *string
	// Used to perform an operation only if the targeted resource's etag matches the value provided.
	IfMatch *string
	// Used to perform an operation only if the targeted resource's etag does not match the value provided.
	IfNoneMatch *string
	// The label of the key-value to retrieve.
	Label *string
	// Used to select what fields are present in the returned resource(s).
	Select []KeyValueFields
}

// SetReadOnlyOptions contains the optional parameters for the Client.SetReadOnly method.
type SetReadOnlyOptions struct {
	// Used to perform an operation only if the targeted resource's etag matches the value provided.
	IfMatch *string
	// Used to perform an operation only if the targeted resource's etag does not match the value provided.
	IfNoneMatch *string
	// The label, if any, of the key-value to lock.
	Label *string
}

// Error - Azure App Configuration error object.
type Error struct {
	// A detailed description of the error.
	Detail *string

	// The name of the parameter that resulted in the error.
	Name *string

	// The HTTP status code that the error maps to.
	Status *int32

	// A brief summary of the error.
	Title *string

	// The type of the error.
	Type *string
}

type Key struct {
	// READ-ONLY
	Name *string
}

// KeyListResult - The result of a list request.
type KeyListResult struct {
	// The collection value.
	Items []*Key

	// The URI that can be used to request the next set of paged results.
	NextLink *string
}

type KeyValue struct {
	ContentType  *string
	ETag         *string
	Key          *string
	Label        *string
	LastModified *time.Time
	IsReadOnly   *bool

	// Dictionary of
	Tags  map[string]*string
	Value *string
}

// KeyValueListResult - The result of a list request.
type KeyValueListResult struct {
	// The collection value.
	Items []*KeyValue

	// The URI that can be used to request the next set of paged results.
	NextLink *string
}

type Label struct {
	// READ-ONLY
	Name *string
}

// LabelListResult - The result of a list request.
type LabelListResult struct {
	// The collection value.
	Items []*Label

	// The URI that can be used to request the next set of paged results.
	NextLink *string
}
