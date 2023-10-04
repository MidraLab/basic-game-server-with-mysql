// Copyright 2015 go-swagger maintainers
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package spec

import (
	"reflect"
	"sync"
	"testing"
)

func Test_simpleCache_ShallowClone(t *testing.T) {
	type fields struct {
		lock  sync.RWMutex
		store map[string]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   ResolutionCache
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &simpleCache{
				lock:  tt.fields.lock,
				store: tt.fields.store,
			}
			if got := s.ShallowClone(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("simpleCache.ShallowClone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_simpleCache_Get(t *testing.T) {
	type fields struct {
		lock  sync.RWMutex
		store map[string]interface{}
	}
	type args struct {
		uri string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &simpleCache{
				lock:  tt.fields.lock,
				store: tt.fields.store,
			}
			got, got1 := s.Get(tt.args.uri)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("simpleCache.Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("simpleCache.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_simpleCache_Set(t *testing.T) {
	type fields struct {
		lock  sync.RWMutex
		store map[string]interface{}
	}
	type args struct {
		uri  string
		data interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &simpleCache{
				lock:  tt.fields.lock,
				store: tt.fields.store,
			}
			s.Set(tt.args.uri, tt.args.data)
		})
	}
}

func Test_initResolutionCache(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initResolutionCache()
		})
	}
}

func Test_defaultResolutionCache(t *testing.T) {
	tests := []struct {
		name string
		want *simpleCache
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := defaultResolutionCache(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("defaultResolutionCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cacheOrDefault(t *testing.T) {
	type args struct {
		cache ResolutionCache
	}
	tests := []struct {
		name string
		args args
		want ResolutionCache
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cacheOrDefault(tt.args.cache); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cacheOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
