//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package generated

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type KeyListResult.
func (k KeyListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "items", k.Items)
	populate(objectMap, "@nextLink", k.NextLink)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type KeyValue.
func (k KeyValue) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "content_type", k.ContentType)
	populate(objectMap, "etag", k.Etag)
	populate(objectMap, "key", k.Key)
	populate(objectMap, "label", k.Label)
	populateTimeRFC3339(objectMap, "last_modified", k.LastModified)
	populate(objectMap, "locked", k.Locked)
	populate(objectMap, "tags", k.Tags)
	populate(objectMap, "value", k.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type KeyValue.
func (k *KeyValue) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "content_type":
			err = unpopulate(val, &k.ContentType)
			delete(rawMsg, key)
		case "etag":
			err = unpopulate(val, &k.Etag)
			delete(rawMsg, key)
		case "key":
			err = unpopulate(val, &k.Key)
			delete(rawMsg, key)
		case "label":
			err = unpopulate(val, &k.Label)
			delete(rawMsg, key)
		case "last_modified":
			err = unpopulateTimeRFC3339(val, &k.LastModified)
			delete(rawMsg, key)
		case "locked":
			err = unpopulate(val, &k.Locked)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, &k.Tags)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, &k.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type KeyValueListResult.
func (k KeyValueListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "items", k.Items)
	populate(objectMap, "@nextLink", k.NextLink)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type LabelListResult.
func (l LabelListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "items", l.Items)
	populate(objectMap, "@nextLink", l.NextLink)
	return json.Marshal(objectMap)
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}
