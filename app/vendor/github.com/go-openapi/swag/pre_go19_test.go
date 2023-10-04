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

//go:build !go1.9
// +build !go1.9

package swag

import (
	"reflect"
	"sync"
	"testing"
)

func Test_newIndexOfInitialisms(t *testing.T) {
	tests := []struct {
		name string
		want *indexOfInitialisms
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newIndexOfInitialisms(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newIndexOfInitialisms() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_indexOfInitialisms_load(t *testing.T) {
	type fields struct {
		sortMutex *sync.Mutex
		index     *sync.Map
	}
	type args struct {
		initial map[string]bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *indexOfInitialisms
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &indexOfInitialisms{
				sortMutex: tt.fields.sortMutex,
				index:     tt.fields.index,
			}
			if got := m.load(tt.args.initial); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("indexOfInitialisms.load() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_indexOfInitialisms_isInitialism(t *testing.T) {
	type fields struct {
		sortMutex *sync.Mutex
		index     *sync.Map
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &indexOfInitialisms{
				sortMutex: tt.fields.sortMutex,
				index:     tt.fields.index,
			}
			if got := m.isInitialism(tt.args.key); got != tt.want {
				t.Errorf("indexOfInitialisms.isInitialism() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_indexOfInitialisms_add(t *testing.T) {
	type fields struct {
		sortMutex *sync.Mutex
		index     *sync.Map
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *indexOfInitialisms
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &indexOfInitialisms{
				sortMutex: tt.fields.sortMutex,
				index:     tt.fields.index,
			}
			if got := m.add(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("indexOfInitialisms.add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_indexOfInitialisms_sorted(t *testing.T) {
	type fields struct {
		sortMutex *sync.Mutex
		index     *sync.Map
	}
	tests := []struct {
		name       string
		fields     fields
		wantResult []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &indexOfInitialisms{
				sortMutex: tt.fields.sortMutex,
				index:     tt.fields.index,
			}
			if gotResult := m.sorted(); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("indexOfInitialisms.sorted() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
