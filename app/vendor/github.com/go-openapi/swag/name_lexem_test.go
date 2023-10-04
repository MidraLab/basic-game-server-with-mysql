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

package swag

import (
	"reflect"
	"testing"
)

func Test_newInitialismNameLexem(t *testing.T) {
	type args struct {
		original          string
		matchedInitialism string
	}
	tests := []struct {
		name string
		args args
		want *initialismNameLexem
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newInitialismNameLexem(tt.args.original, tt.args.matchedInitialism); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newInitialismNameLexem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newCasualNameLexem(t *testing.T) {
	type args struct {
		original string
	}
	tests := []struct {
		name string
		args args
		want *casualNameLexem
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newCasualNameLexem(tt.args.original); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newCasualNameLexem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_initialismNameLexem_GetUnsafeGoName(t *testing.T) {
	type fields struct {
		original          string
		matchedInitialism string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &initialismNameLexem{
				original:          tt.fields.original,
				matchedInitialism: tt.fields.matchedInitialism,
			}
			if got := l.GetUnsafeGoName(); got != tt.want {
				t.Errorf("initialismNameLexem.GetUnsafeGoName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_casualNameLexem_GetUnsafeGoName(t *testing.T) {
	type fields struct {
		original string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &casualNameLexem{
				original: tt.fields.original,
			}
			if got := l.GetUnsafeGoName(); got != tt.want {
				t.Errorf("casualNameLexem.GetUnsafeGoName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_initialismNameLexem_GetOriginal(t *testing.T) {
	type fields struct {
		original          string
		matchedInitialism string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &initialismNameLexem{
				original:          tt.fields.original,
				matchedInitialism: tt.fields.matchedInitialism,
			}
			if got := l.GetOriginal(); got != tt.want {
				t.Errorf("initialismNameLexem.GetOriginal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_casualNameLexem_GetOriginal(t *testing.T) {
	type fields struct {
		original string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &casualNameLexem{
				original: tt.fields.original,
			}
			if got := l.GetOriginal(); got != tt.want {
				t.Errorf("casualNameLexem.GetOriginal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_initialismNameLexem_IsInitialism(t *testing.T) {
	type fields struct {
		original          string
		matchedInitialism string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &initialismNameLexem{
				original:          tt.fields.original,
				matchedInitialism: tt.fields.matchedInitialism,
			}
			if got := l.IsInitialism(); got != tt.want {
				t.Errorf("initialismNameLexem.IsInitialism() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_casualNameLexem_IsInitialism(t *testing.T) {
	type fields struct {
		original string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &casualNameLexem{
				original: tt.fields.original,
			}
			if got := l.IsInitialism(); got != tt.want {
				t.Errorf("casualNameLexem.IsInitialism() = %v, want %v", got, tt.want)
			}
		})
	}
}
