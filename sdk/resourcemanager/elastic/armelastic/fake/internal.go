//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.
// Code generated by @autorest/go. DO NOT EDIT.

package fake

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"net/http"
	"reflect"
	"sync"
)

type nonRetriableError struct {
	error
}

func (nonRetriableError) NonRetriable() {
	// marker method
}

func contains[T comparable](s []T, v T) bool {
	for _, vv := range s {
		if vv == v {
			return true
		}
	}
	return false
}

func getOptional[T any](v T) *T {
	if reflect.ValueOf(v).IsZero() {
		return nil
	}
	return &v
}

func newTracker[T any]() *tracker[T] {
	return &tracker[T]{
		items: map[string]*T{},
	}
}

type tracker[T any] struct {
	items map[string]*T
	mu    sync.Mutex
}

func (p *tracker[T]) get(req *http.Request) *T {
	p.mu.Lock()
	defer p.mu.Unlock()
	if item, ok := p.items[server.SanitizePagerPollerPath(req.URL.Path)]; ok {
		return item
	}
	return nil
}

func (p *tracker[T]) add(req *http.Request, item *T) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.items[server.SanitizePagerPollerPath(req.URL.Path)] = item
}

func (p *tracker[T]) remove(req *http.Request) {
	p.mu.Lock()
	defer p.mu.Unlock()
	delete(p.items, server.SanitizePagerPollerPath(req.URL.Path))
}